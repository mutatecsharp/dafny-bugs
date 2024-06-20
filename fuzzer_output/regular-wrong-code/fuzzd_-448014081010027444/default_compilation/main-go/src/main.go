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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(!(true), true, !(true))))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	if false {
		return true
	} else {
		return (!(!(!(!(!(!(!(false)))))))) && (true)
	}
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(true, true, false)).Cardinality()), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(463), false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(308), (_dafny.SetOf(_dafny.IntOfInt64(930), _dafny.IntOfInt64(305))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kie")).Cardinality()), _dafny.IntOfInt64(941), (_dafny.MultiSetOf(Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(true))).Cardinality(), true, true))).Cardinality())).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-918), true)))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true, true)).Difference(_dafny.MultiSetOf(false))).Intersection(_dafny.MultiSetOf(true))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.Int {
	if false {
		return _dafny.IntOfInt64(-885)
	} else {
		return _dafny.IntOfUint32((_dafny.SeqOf(true, !(false), true, false)).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return (((Companion_D6_.Create_DC18_(_dafny.SetOf(false))).Dtor_cf27()).Union(_dafny.SetOf(false, true))).Difference((_dafny.SetOf(false)).Intersection(_dafny.SetOf(!(true), !(!(false)))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source0 D2 = Companion_D2_.Create_DC7_(true, _dafny.IntOfInt64(462))
	_ = _source0
	if _source0.Is_DC7() {
		var _0___mcc_h0 bool = _source0.Get_().(D2_DC7).Cf11
		_ = _0___mcc_h0
		var _1___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC7).Cf12
		_ = _1___mcc_h1
		var _2_cf12 _dafny.Int = _1___mcc_h1
		_ = _2_cf12
		var _3_cf11 bool = _0___mcc_h0
		_ = _3_cf11
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(127), true, _3_cf11), _3_cf11)
	} else if _source0.Is_DC6() {
		var _4___mcc_h2 _dafny.Map = _source0.Get_().(D2_DC6).Cf10
		_ = _4___mcc_h2
		var _5_cf10 _dafny.Map = _4___mcc_h2
		_ = _5_cf10
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wdogpmi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))).Cardinality()))).Keys().Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _7_v0 _dafny.Int
				_7_v0 = interface{}(_compr_0).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wdogpmi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg1 _dafny.Int) interface{} {
						return coer1(arg1)
					}
				}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality()))).Contains(_7_v0) {
					_coll0.Add(Companion_Default___.SafeModuloInt(_7_v0, _dafny.IntOfInt64(857)), true)
				}
			}
			return _coll0.ToMap()
		}()).Cardinality()), _dafny.IntOfInt64(969))).Cardinality(), false, !(false)), false)
	} else {
		var _8___mcc_h3 D2 = _source0.Get_().(D2_DC8).Cf13
		_ = _8___mcc_h3
		var _9_cf13 D2 = _8___mcc_h3
		_ = _9_cf13
		return func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SetOf(Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ogqm")).Cardinality()), true, true), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-484), !(false), true))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _10_v1 D1
				_10_v1 = interface{}(_compr_1).(D1)
				if (_dafny.SetOf(Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ogqm")).Cardinality()), true, true), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-484), !(false), true))).Contains(_10_v1) {
					_coll1.Add(_10_v1, true)
				}
			}
			return _coll1.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(589), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(645))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-720))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-525))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_12_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(44)
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(false, !(!(!(true))))).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mqtumvw")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mrikrvux"), (Companion_D4_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("ur"))).Dtor_cf20()), _dafny.UnicodeSeqOfUtf8Bytes("hytv"))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('t')
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D9_.Create_DC24_(_dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(878)))))).Dtor_cf38()
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('k'), _dafny.CodePoint('x'), _dafny.CodePoint('n'), _dafny.CodePoint('x'), _dafny.CodePoint('l'))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("urylnfh"), _dafny.UnicodeSeqOfUtf8Bytes("mimbr")))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-208), _dafny.IntOfInt64(303))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _13_v0 _dafny.Int
			_13_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-208)).Cmp(_13_v0) <= 0) && ((_13_v0).Cmp(_dafny.IntOfInt64(303)) < 0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_13_v0, _dafny.IntOfInt64(869)))
			}
		}
		return _coll2.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 bool = false
	_ = r2
	var _14_v0 _dafny.MultiSet
	_ = _14_v0
	_14_v0 = _dafny.MultiSetOf(true)
	var _15_v1 _dafny.Sequence
	_ = _15_v1
	_15_v1 = _dafny.SeqOf(_14_v0)
	if (((_15_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_15_v1).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_14_v0)).IsProperSubsetOf((_dafny.MultiSetOf(p2, !(!(p2)), p2, false)).Intersection(_14_v0)) {
		if !(_dafny.MultiSetOf(p1)).Contains(Companion_Default___.SafeDivisionInt(p1, p1)) {
			var _16_v2 _dafny.Set
			_ = _16_v2
			_16_v2 = _dafny.SetOf(!(p0), p0)
			r1 = Companion_Default___.SafeDivisionInt((p1).Plus(p1), (_16_v2).Cardinality())
			var _17_v3 _dafny.Array
			_ = _17_v3
			var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw0
			_17_v3 = _nw0
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_17_v3), 0))
			_ = _index0
			(_17_v3).ArraySet1(p1, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_17_v3), 0))
			_ = _index1
			(_17_v3).ArraySet1(p1, (_index1).Int())
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_17_v3), 0))
			_ = _index2
			(_17_v3).ArraySet1(p1, (_index2).Int())
			var _18_v4 _dafny.Array
			_ = _18_v4
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_0
			var _nw1 _dafny.Array
			_ = _nw1
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw1 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) bool = (func(_19_p0 bool) func(_dafny.Int) bool {
					return func(_20_i0 _dafny.Int) bool {
						return _19_p0
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
			_18_v4 = _nw1
			var _21_v5 *C0
			_ = _21_v5
			var _nw2 *C0 = New_C0_()
			_ = _nw2
			_nw2.Ctor__(_18_v4)
			_21_v5 = _nw2
			r2 = p2
		} else {
			var _22_v6 _dafny.Sequence
			_ = _22_v6
			_22_v6 = _dafny.SeqOf(p0, p0)
			var _23_v7 _dafny.Sequence
			_ = _23_v7
			_23_v7 = _dafny.SeqOf((_dafny.MultiSetFromSeq(_22_v6)).Cardinality())
			var _24_v8 D3
			_ = _24_v8
			_24_v8 = Companion_D3_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(835))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}((func(_25_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_26_i1 _dafny.Int) _dafny.Int {
					return _25_p1
				}
			})(p1))), p2)
			var _27_v9 _dafny.Sequence
			_ = _27_v9
			_27_v9 = _dafny.UnicodeSeqOfUtf8Bytes("dwnjkq")
			var _28_v10 _dafny.Sequence
			_ = _28_v10
			_28_v10 = _dafny.SeqOf(_23_v7)
			var _29_v11 _dafny.CodePoint
			_ = _29_v11
			_29_v11 = _dafny.CodePoint('r')
			var _30_v12 _dafny.Array
			_ = _30_v12
			var _nwElement0_0 _dafny.Sequence = _23_v7
			_ = _nwElement0_0
			var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(18))
			_ = _nw3
			(_nw3).ArraySet1(_nwElement0_0, 0)
			(_nw3).ArraySet1((_24_v8).Dtor_cf17(), 1)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-640))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_31_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_32_i2 _dafny.Int) _dafny.Int {
					return _31_p1
				}
			})(p1))), _23_v7), 2)
			(_nw3).ArraySet1(_23_v7, 3)
			(_nw3).ArraySet1(_dafny.SeqOf(p1, p1, p1, p1, p1), 4)
			(_nw3).ArraySet1(_dafny.SeqOf(p1, p1, p1, p1), 5)
			(_nw3).ArraySet1(_23_v7, 6)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_33_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_34_i3 _dafny.Int) _dafny.Int {
					return _33_p1
				}
			})(p1))), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_35_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_36_i3 _dafny.Int) _dafny.Int {
					return _35_p1
				}
			})(p1)))).Cardinality()))).Uint32(), p1), 7)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Update(_23_v7, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_23_v7).Cardinality()))).Uint32(), _dafny.IntOfUint32((_27_v9).Cardinality())), 8)
			(_nw3).ArraySet1((_28_v10).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_28_v10).Cardinality()))).Uint32()).(_dafny.Sequence), 9)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Update(_23_v7, (Companion_Default___.SafeIndex(Companion_Default___.Fm4(globalState), _dafny.IntOfUint32((_23_v7).Cardinality()))).Uint32(), _dafny.IntOfInt64(-246)), 10)
			(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(374))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_37_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(882)
			})), 11)
			(_nw3).ArraySet1(_23_v7, 12)
			(_nw3).ArraySet1(_23_v7, 13)
			(_nw3).ArraySet1(_23_v7, 14)
			(_nw3).ArraySet1(_23_v7, 15)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_23_v7, _23_v7), 16)
			(_nw3).ArraySet1(Companion_Default___.Fm7(p2, _29_v11, _dafny.IntOfInt64(349), globalState), 17)
			_30_v12 = _nw3
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_30_v12), 0))
			_ = _index3
			(_30_v12).ArraySet1(_23_v7, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_30_v12), 0))
			_ = _index4
			(_30_v12).ArraySet1(_23_v7, (_index4).Int())
			var _38_v13 D3
			_ = _38_v13
			_38_v13 = Companion_D3_.Create_DC9_((_30_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_30_v12), 0))).Int()).(_dafny.Sequence))
			_38_v13 = _38_v13
			var _39_v14 _dafny.Array
			_ = _39_v14
			var _nwElement0_1 _dafny.CodePoint = Companion_Default___.Fm9(p2, globalState)
			_ = _nwElement0_1
			var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
			_ = _nw4
			(_nw4).ArraySet1CodePoint(_nwElement0_1, 0)
			(_nw4).ArraySet1CodePoint(_29_v11, 1)
			(_nw4).ArraySet1CodePoint(_29_v11, 2)
			(_nw4).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p2 {
					return _29_v11
				}
				return _dafny.CodePoint('v')
			})(), 3)
			(_nw4).ArraySet1CodePoint(_29_v11, 4)
			(_nw4).ArraySet1CodePoint(_29_v11, 5)
			(_nw4).ArraySet1CodePoint(_29_v11, 6)
			(_nw4).ArraySet1CodePoint(_29_v11, 7)
			_39_v14 = _nw4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_39_v14), 0))
			_ = _index5
			(_39_v14).ArraySet1CodePoint(_29_v11, (_index5).Int())
			var _40_v15 _dafny.Array
			_ = _40_v15
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_41_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_42_i5 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_42_i5, _41_p1)
					}
				})(p1)
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
			_40_v15 = _nw5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_39_v14), 0))
			_ = _index6
			var _rhs0 bool = p0
			_ = _rhs0
			var _rhs1 _dafny.CodePoint = _29_v11
			_ = _rhs1
			var _rhs2 _dafny.Array = _40_v15
			_ = _rhs2
			var _lhs0 _dafny.Array = _39_v14
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_39_v14), 0))
			_ = _lhs1
			r2 = _rhs0
			(_lhs0).ArraySet1CodePoint(_rhs1, (_lhs1).Int())
			_40_v15 = _rhs2
			var _43_v16 _dafny.Array
			_ = _43_v16
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.MultiSet = (func(_44_v0 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_45_i6 _dafny.Int) _dafny.MultiSet {
						return _44_v0
					}
				})(_14_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw6 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw6).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw6).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_43_v16 = _nw6
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_43_v16), 0))
			_ = _index7
			(_43_v16).ArraySet1(_14_v0, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_43_v16), 0))
			_ = _index8
			(_43_v16).ArraySet1(_14_v0, (_index8).Int())
			var _46_v17 _dafny.Array
			_ = _46_v17
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw7
			_46_v17 = _nw7
			var _47_v18 *C0
			_ = _47_v18
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__(_46_v17)
			_47_v18 = _nw8
		}
		var _48_v19 _dafny.Map
		_ = _48_v19
		_48_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), p0)
		var _49_v20 _dafny.Sequence
		_ = _49_v20
		_49_v20 = _dafny.UnicodeSeqOfUtf8Bytes("fxdguiwrc")
		var _50_v21 D4
		_ = _50_v21
		_50_v21 = Companion_D4_.Create_DC13_(_49_v20)
		var _51_v22 _dafny.Sequence
		_ = _51_v22
		_51_v22 = _dafny.SeqOf(p0, false, p0)
		var _52_v23 _dafny.Map
		_ = _52_v23
		_52_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v22, p0)
		var _53_v24 D3
		_ = _53_v24
		_53_v24 = Companion_D3_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_54_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_55_i7 _dafny.Int) _dafny.Int {
				return _54_p1
			}
		})(p1))), p0)
		var _56_v25 _dafny.Array
		_ = _56_v25
		var _nwElement0_2 bool = false
		_ = _nwElement0_2
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(29))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_2, 0)
		(_nw9).ArraySet1(p0, 1)
		(_nw9).ArraySet1(Companion_Default___.Fm1(_48_v19, (_50_v21).Dtor_cf20(), _dafny.UnicodeSeqOfUtf8Bytes("lenrsoqb"), globalState), 2)
		(_nw9).ArraySet1(p2, 3)
		(_nw9).ArraySet1(p0, 4)
		(_nw9).ArraySet1(p0, 5)
		(_nw9).ArraySet1(p0, 6)
		(_nw9).ArraySet1(p2, 7)
		(_nw9).ArraySet1(p0, 8)
		(_nw9).ArraySet1(p0, 9)
		(_nw9).ArraySet1(p2, 10)
		(_nw9).ArraySet1(!(!(p2)), 11)
		(_nw9).ArraySet1(true, 12)
		(_nw9).ArraySet1(!(p0), 13)
		(_nw9).ArraySet1(false, 14)
		(_nw9).ArraySet1(p2, 15)
		(_nw9).ArraySet1(p2, 16)
		(_nw9).ArraySet1(p0, 17)
		(_nw9).ArraySet1((func() bool {
			if (_52_v23).Contains(_51_v22) {
				return (_52_v23).Get(_51_v22).(bool)
			}
			return !(p0)
		})(), 18)
		(_nw9).ArraySet1(true, 19)
		(_nw9).ArraySet1(p2, 20)
		(_nw9).ArraySet1(!(p0), 21)
		(_nw9).ArraySet1((_53_v24).Dtor_cf18(), 22)
		(_nw9).ArraySet1(p0, 23)
		(_nw9).ArraySet1(p2, 24)
		(_nw9).ArraySet1(p0, 25)
		(_nw9).ArraySet1(p2, 26)
		(_nw9).ArraySet1(!(p2), 27)
		(_nw9).ArraySet1(p0, 28)
		_56_v25 = _nw9
		var _57_v26 *C0
		_ = _57_v26
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__(_56_v25)
		_57_v26 = _nw10
		var _58_v27 _dafny.Array
		_ = _58_v27
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
		_ = _nw11
		_58_v27 = _nw11
		var _59_v28 _dafny.Array
		_ = _59_v28
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_3
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_60_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_61_i8 _dafny.Int) _dafny.Int {
					return (_61_i8).Plus(_60_p1)
				}
			})(p1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw12 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw12).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw12).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_59_v28 = _nw12
		var _62_v29 _dafny.Sequence
		_ = _62_v29
		_62_v29 = _dafny.SeqOf(_59_v28)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_58_v27), 0))
		_ = _index9
		(_58_v27).ArraySet1(_62_v29, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_58_v27), 0))
		_ = _index10
		(_58_v27).ArraySet1(_62_v29, (_index10).Int())
		var _63_v30 *C0
		_ = _63_v30
		var _nw13 *C0 = New_C0_()
		_ = _nw13
		_nw13.Ctor__((_57_v26).F15())
		_63_v30 = _nw13
		(globalState).F0 = Companion_Default___.SafeModuloInt(p1, (_dafny.Zero).Minus(p1))
	} else {
		if !(p0) {
			var _64_v31 _dafny.Array
			_ = _64_v31
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw14
			_64_v31 = _nw14
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_64_v31), 0))
			_ = _index11
			(_64_v31).ArraySet1(p2, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_64_v31), 0))
			_ = _index12
			(_64_v31).ArraySet1(p2, (_index12).Int())
			var _65_v32 *C0
			_ = _65_v32
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_64_v31)
			_65_v32 = _nw15
			_65_v32 = _65_v32
			var _66_v33 *C0
			_ = _66_v33
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__((_65_v32).F15())
			_66_v33 = _nw16
			var _67_v34 _dafny.Set
			_ = _67_v34
			_67_v34 = _dafny.SetOf(p0, p0, p2, p2, (_64_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_64_v31), 0))).Int()).(bool))
			var _68_v35 _dafny.MultiSet
			_ = _68_v35
			_68_v35 = _dafny.MultiSetOf(_67_v34, _67_v34, _67_v34)
			var _69_v36 D5
			_ = _69_v36
			_69_v36 = Companion_D5_.Create_DC15_(_68_v35)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_64_v31), 0))
			_ = _index13
			(_64_v31).ArraySet1(((_69_v36).Dtor_cf23()).IsDisjointFrom(_dafny.MultiSetOf(_67_v34)), (_index13).Int())
			var _70_v37 _dafny.Array
			_ = _70_v37
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
			_ = _nw17
			_70_v37 = _nw17
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_70_v37), 0))
			_ = _index14
			(_70_v37).ArraySet1(Companion_Default___.Fm11(globalState), (_index14).Int())
			var _71_v38 _dafny.Set
			_ = _71_v38
			_71_v38 = _dafny.SetOf(_dafny.CodePoint('f'))
			var _72_v39 _dafny.CodePoint
			_ = _72_v39
			_72_v39 = _dafny.CodePoint('v')
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_70_v37), 0))
			_ = _index15
			(_70_v37).ArraySet1(((_71_v38).Intersection(_dafny.SetOf(_dafny.CodePoint('y'), _72_v39, _72_v39))).Union(_71_v38), (_index15).Int())
		} else {
			var _73_v40 D5
			_ = _73_v40
			_73_v40 = Companion_D5_.Create_DC17_(p2)
			var _pat_let_tv0 = p0
			_ = _pat_let_tv0
			_73_v40 = (func() D5 {
				if true {
					return func(_pat_let0_0 D5) D5 {
						return func(_74_dt__update__tmp_h0 D5) D5 {
							return func(_pat_let1_0 bool) D5 {
								return func(_75_dt__update_hcf26_h0 bool) D5 {
									return Companion_D5_.Create_DC17_(_75_dt__update_hcf26_h0)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_73_v40)
				}
				return _73_v40
			})()
			var _76_v41 _dafny.Array
			_ = _76_v41
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw18
			_76_v41 = _nw18
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_76_v41), 0))
			_ = _index16
			(_76_v41).ArraySet1(p1, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_76_v41), 0))
			_ = _index17
			(_76_v41).ArraySet1(((_dafny.Zero).Minus(p1)).Minus(_dafny.IntOfInt64(-577)), (_index17).Int())
			var _77_v42 _dafny.Array
			_ = _77_v42
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
			_ = _nw19
			_77_v42 = _nw19
			var _78_v43 _dafny.Sequence
			_ = _78_v43
			_78_v43 = _dafny.UnicodeSeqOfUtf8Bytes("ctnu")
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_77_v42), 0))
			_ = _index18
			(_77_v42).ArraySet1(_78_v43, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_77_v42), 0))
			_ = _index19
			(_77_v42).ArraySet1(_78_v43, (_index19).Int())
			var _79_v44 _dafny.Array
			_ = _79_v44
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_4
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_80_p0 bool) func(_dafny.Int) bool {
					return func(_81_i9 _dafny.Int) bool {
						return _80_p0
					}
				})(p0)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw20 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw20).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw20).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_79_v44 = _nw20
			var _82_v45 *C0
			_ = _82_v45
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__(_79_v44)
			_82_v45 = _nw21
			var _83_v46 _dafny.CodePoint
			_ = _83_v46
			_83_v46 = _dafny.CodePoint('r')
			var _rhs3 _dafny.Sequence = Companion_Default___.Fm8(!(p0) || (p0), _83_v46, _83_v46, globalState)
			_ = _rhs3
			var _rhs4 _dafny.Int = p1
			_ = _rhs4
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			_lhs2.F5 = _rhs3
			_lhs3.F0 = _rhs4
		}
		var _84_v47 _dafny.Map
		_ = _84_v47
		_84_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
		var _85_v48 _dafny.Array
		_ = _85_v48
		var _nwElement0_3 bool = p0
		_ = _nwElement0_3
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(4))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_3, 0)
		(_nw22).ArraySet1(p2, 1)
		(_nw22).ArraySet1(p2, 2)
		(_nw22).ArraySet1((func() bool {
			if (_84_v47).Contains(p1) {
				return (_84_v47).Get(p1).(bool)
			}
			return p2
		})(), 3)
		_85_v48 = _nw22
		var _86_v49 *C0
		_ = _86_v49
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__(_85_v48)
		_86_v49 = _nw23
		var _87_v50 _dafny.Map
		_ = _87_v50
		_87_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_84_v47).Merge(_84_v47))
		var _88_v51 _dafny.MultiSet
		_ = _88_v51
		_88_v51 = _dafny.MultiSetOf(_86_v49, _86_v49)
		_87_v50 = (_87_v50).Update(((_88_v51).Cardinality()).Times(p1), (func() _dafny.Map {
			if (_87_v50).Contains(_dafny.IntOfInt64(70)) {
				return (_87_v50).Get(_dafny.IntOfInt64(70)).(_dafny.Map)
			}
			return _84_v47
		})())
		var _89_v52 *C0
		_ = _89_v52
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__((_86_v49).F15())
		_89_v52 = _nw24
		_89_v52 = _89_v52
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen(((_89_v52).F15()), 0))
		_ = _index20
		((_89_v52).F15()).ArraySet1(p2, (_index20).Int())
		var _90_v53 _dafny.Sequence
		_ = _90_v53
		_90_v53 = _dafny.UnicodeSeqOfUtf8Bytes("uuolsr")
		var _91_v54 _dafny.MultiSet
		_ = _91_v54
		_91_v54 = _dafny.MultiSetOf(_90_v53)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen(((_89_v52).F15()), 0))
		_ = _index21
		((_89_v52).F15()).ArraySet1(!(_91_v54).Equals(_91_v54), (_index21).Int())
	}
	if true {
		r1 = Companion_Default___.Fm4(globalState)
		var _source1 D4 = Companion_Default___.Fm12(globalState)
		_ = _source1
		if _source1.Is_DC14() {
			var _92___mcc_h0 *C0 = _source1.Get_().(D4_DC14).Cf21
			_ = _92___mcc_h0
			var _93___mcc_h1 *C0 = _source1.Get_().(D4_DC14).Cf22
			_ = _93___mcc_h1
			var _94_cf22 *C0 = _93___mcc_h1
			_ = _94_cf22
			var _95_cf21 *C0 = _92___mcc_h0
			_ = _95_cf21
			r2 = p2
			r2 = p2
			var _96_v55 D1
			_ = _96_v55
			_96_v55 = Companion_D1_.Create_DC1_((_dafny.Zero).Minus(p1))
			var _97_v56 D2
			_ = _97_v56
			_97_v56 = Companion_D2_.Create_DC7_(p2, (_96_v55).Dtor_cf1())
			_97_v56 = _97_v56
			var _98_v57 _dafny.Map
			_ = _98_v57
			_98_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _99_v58 _dafny.MultiSet
			_ = _99_v58
			_99_v58 = _dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(521)))
			var _100_v59 _dafny.Set
			_ = _100_v59
			_100_v59 = _dafny.SetOf(p1, _dafny.IntOfInt64(-743))
			_98_v57 = (_98_v57).Update(true, (_99_v58).IsDisjointFrom(_dafny.MultiSetOf(_100_v59, _100_v59)))
		} else {
			var _101___mcc_h2 _dafny.Sequence = _source1.Get_().(D4_DC13).Cf20
			_ = _101___mcc_h2
			var _102_cf20 _dafny.Sequence = _101___mcc_h2
			_ = _102_cf20
			var _103_v60 _dafny.Map
			_ = _103_v60
			_103_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(316), ((_dafny.MultiSetOf(false)).Update(p0, Companion_Default___.Abs(p1))).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(884))))).Cardinality())
			var _104_v61 _dafny.CodePoint
			_ = _104_v61
			_104_v61 = _dafny.CodePoint('d')
			var _105_v62 _dafny.Map
			_ = _105_v62
			_105_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jyww"), p0)
			var _106_v63 _dafny.Sequence
			_ = _106_v63
			_106_v63 = _dafny.SeqOf(!(p0), p2, Companion_Default___.Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_103_v60).Cardinality(), p0), _dafny.Companion_Sequence_.Update(_102_cf20, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_102_cf20).Cardinality()))).Uint32(), _104_v61), _dafny.UnicodeSeqOfUtf8Bytes("rswrhg"), globalState), (func() bool {
				if (_105_v62).Contains(_dafny.Companion_Sequence_.Update(_102_cf20, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_102_cf20).Cardinality()))).Uint32(), _104_v61)) {
					return (_105_v62).Get(_dafny.Companion_Sequence_.Update(_102_cf20, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_102_cf20).Cardinality()))).Uint32(), _104_v61)).(bool)
				}
				return p0
			})())
			var _107_v64 _dafny.Map
			_ = _107_v64
			_107_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _106_v63)
			_107_v64 = _107_v64
			var _108_v65 _dafny.Array
			_ = _108_v65
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw25
			_108_v65 = _nw25
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_108_v65), 0))
			_ = _index22
			(_108_v65).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2), (_index22).Int())
			var _109_v66 _dafny.Map
			_ = _109_v66
			_109_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (func() bool {
				if true {
					return p0
				}
				return p2
			})())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_108_v65), 0))
			_ = _index23
			(_108_v65).ArraySet1(_109_v66, (_index23).Int())
			var _110_v67 _dafny.Array
			_ = _110_v67
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw26
			_110_v67 = _nw26
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_110_v67), 0))
			_ = _index24
			(_110_v67).ArraySet1(p0, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_110_v67), 0))
			_ = _index25
			(_110_v67).ArraySet1((_dafny.IntOfInt64(989)).Cmp(p1) != 0, (_index25).Int())
			var _111_v68 _dafny.Map
			_ = _111_v68
			_111_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_110_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_110_v67), 0))).Int()).(bool), p1)
			var _112_v69 D2
			_ = _112_v69
			_112_v69 = Companion_D2_.Create_DC6_(_111_v68)
			var _113_v70 _dafny.Sequence
			_ = _113_v70
			_113_v70 = _dafny.SeqOf(_112_v69, _112_v69, Companion_D2_.Create_DC6_(_111_v68))
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_110_v67), 0))
			_ = _index26
			(_110_v67).ArraySet1(_dafny.Companion_Sequence_.Contains(_113_v70, _112_v69), (_index26).Int())
		}
		var _114_v71 _dafny.Array
		_ = _114_v71
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_5
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_115_p0 bool) func(_dafny.Int) bool {
				return func(_116_i10 _dafny.Int) bool {
					return _115_p0
				}
			})(p0)
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
		_114_v71 = _nw27
		var _117_v72 *C0
		_ = _117_v72
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__(_114_v71)
		_117_v72 = _nw28
		if p0 {
			r0 = p1
			var _118_v73 _dafny.Set
			_ = _118_v73
			_118_v73 = _dafny.SetOf(true, !(true))
			var _119_v74 D6
			_ = _119_v74
			_119_v74 = Companion_D6_.Create_DC18_(_118_v73)
			var _rhs5 _dafny.Int = (_dafny.Zero).Minus(p1)
			_ = _rhs5
			var _rhs6 _dafny.Set = (_119_v74).Dtor_cf27()
			_ = _rhs6
			r1 = _rhs5
			_118_v73 = _rhs6
			var _120_v75 _dafny.Array
			_ = _120_v75
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw29
			_120_v75 = _nw29
			var _121_v76 D7
			_ = _121_v76
			_121_v76 = Companion_D7_.Create_DC20_(_120_v75)
			var _122_v77 _dafny.Array
			_ = _122_v77
			var _nwElement0_4 _dafny.Array = _120_v75
			_ = _nwElement0_4
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(18))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_4, 0)
			(_nw30).ArraySet1(_120_v75, 1)
			(_nw30).ArraySet1(_120_v75, 2)
			(_nw30).ArraySet1(_120_v75, 3)
			(_nw30).ArraySet1(_120_v75, 4)
			(_nw30).ArraySet1(_120_v75, 5)
			(_nw30).ArraySet1(_120_v75, 6)
			(_nw30).ArraySet1(_120_v75, 7)
			(_nw30).ArraySet1(_120_v75, 8)
			(_nw30).ArraySet1(_120_v75, 9)
			(_nw30).ArraySet1((_121_v76).Dtor_cf31(), 10)
			(_nw30).ArraySet1((_121_v76).Dtor_cf31(), 11)
			(_nw30).ArraySet1(_120_v75, 12)
			(_nw30).ArraySet1(_120_v75, 13)
			(_nw30).ArraySet1(_120_v75, 14)
			(_nw30).ArraySet1(_120_v75, 15)
			(_nw30).ArraySet1(_120_v75, 16)
			(_nw30).ArraySet1(_120_v75, 17)
			_122_v77 = _nw30
			var _123_v78 _dafny.Map
			_ = _123_v78
			_123_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _122_v77)
			_123_v78 = (_123_v78).Update(p0, _122_v77)
			(globalState).F0 = p1
			var _124_v79 _dafny.Sequence
			_ = _124_v79
			_124_v79 = _dafny.UnicodeSeqOfUtf8Bytes("qcyapl")
			r2 = !((func() bool {
				if !_dafny.Companion_Sequence_.Equal(_124_v79, _124_v79) {
					return !(p2)
				}
				return false
			})())
		} else {
			var _125_v80 _dafny.Map
			_ = _125_v80
			_125_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _126_v81 _dafny.Sequence
			_ = _126_v81
			_126_v81 = _dafny.SeqOf((_125_v80).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm0(p0, globalState)).Cardinality()), p1)
			var _127_v82 D3
			_ = _127_v82
			_127_v82 = Companion_D3_.Create_DC11_(_126_v81, p2)
			_127_v82 = _127_v82
			r1 = p1
			var _128_v83 _dafny.Array
			_ = _128_v83
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_6
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_129_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_130_i11 _dafny.Int) _dafny.Int {
						return (_130_i11).Times(_129_p1)
					}
				})(p1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw31 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw31).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw31).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_128_v83 = _nw31
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_128_v83), 0))
			_ = _index27
			(_128_v83).ArraySet1((_dafny.Zero).Minus(p1), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_128_v83), 0))
			_ = _index28
			(_128_v83).ArraySet1(((func() _dafny.Int {
				if (_14_v0).Contains(true) {
					return (_14_v0).Multiplicity(true)
				}
				return p1
			})()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eskgngpo")).Cardinality())), (_index28).Int())
			var _131_v84 *C0
			_ = _131_v84
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__((_117_v72).F15())
			_131_v84 = _nw32
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen(((_117_v72).F15()), 0))
			_ = _index29
			((_117_v72).F15()).ArraySet1(p0, (_index29).Int())
			var _132_v85 _dafny.Set
			_ = _132_v85
			_132_v85 = _dafny.SetOf((_128_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_128_v83), 0))).Int()).(_dafny.Int), p1)
			var _133_v86 _dafny.Sequence
			_ = _133_v86
			_133_v86 = _dafny.SeqOf(p0)
			var _134_v87 _dafny.CodePoint
			_ = _134_v87
			_134_v87 = _dafny.CodePoint('f')
			var _135_v88 _dafny.Map
			_ = _135_v88
			_135_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_133_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_126_v81).Cardinality()), _dafny.IntOfUint32((_133_v86).Cardinality()))).Uint32()).(bool), _134_v87)
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen(((_117_v72).F15()), 0))
			_ = _index30
			((_117_v72).F15()).ArraySet1(((_dafny.SetOf(p1, _dafny.IntOfUint32((_dafny.SeqOf(_135_v88)).Cardinality()))).Difference(_132_v85)).IsProperSubsetOf(_132_v85), (_index30).Int())
		}
		var _136_v89 _dafny.Map
		_ = _136_v89
		_136_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (_117_v72).F15())
		_136_v89 = (_136_v89).Update(p2, _114_v71)
	} else {
		var _137_v90 _dafny.Sequence
		_ = _137_v90
		_137_v90 = _dafny.SeqOf((p1).Times(p1), p1)
		var _138_v91 _dafny.Array
		_ = _138_v91
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw33
		_138_v91 = _nw33
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v91), 0))
		_ = _index31
		(_138_v91).ArraySet1(p0, (_index31).Int())
		var _139_v92 _dafny.Sequence
		_ = _139_v92
		_139_v92 = _dafny.UnicodeSeqOfUtf8Bytes("j")
		var _140_v93 _dafny.CodePoint
		_ = _140_v93
		_140_v93 = _dafny.CodePoint('p')
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v91), 0))
		_ = _index32
		var _rhs7 _dafny.Sequence = _137_v90
		_ = _rhs7
		var _rhs8 bool = ((p1).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_139_v92, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_139_v92).Cardinality()))).Uint32(), _140_v93)).Cardinality()))).Cmp((_137_v90).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_137_v90).Cardinality()))).Uint32()).(_dafny.Int)) < 0
		_ = _rhs8
		var _lhs4 _dafny.Array = _138_v91
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v91), 0))
		_ = _lhs5
		_137_v90 = _rhs7
		(_lhs4).ArraySet1(_rhs8, (_lhs5).Int())
		var _141_v94 *C0
		_ = _141_v94
		var _nw34 *C0 = New_C0_()
		_ = _nw34
		_nw34.Ctor__(_138_v91)
		_141_v94 = _nw34
		var _142_v95 _dafny.Sequence
		_ = _142_v95
		_142_v95 = _dafny.SeqOf(_141_v94)
		var _143_v96 _dafny.MultiSet
		_ = _143_v96
		_143_v96 = _dafny.MultiSetOf(_141_v94, _141_v94, (_142_v95).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-526), _dafny.IntOfUint32((_142_v95).Cardinality()))).Uint32()).(*C0), _141_v94)
		_143_v96 = _143_v96
		var _144_v97 _dafny.Array
		_ = _144_v97
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
		_ = _nw35
		_144_v97 = _nw35
		var _145_v98 _dafny.Map
		_ = _145_v98
		_145_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), p1)
		var _146_v99 _dafny.MultiSet
		_ = _146_v99
		_146_v99 = _dafny.MultiSetOf(_145_v98)
		var _147_v100 _dafny.MultiSet
		_ = _147_v100
		_147_v100 = _146_v99
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_144_v97), 0))
		_ = _index33
		(_144_v97).ArraySet1((_147_v100), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_144_v97), 0))
		_ = _index34
		(_144_v97).ArraySet1(_146_v99, (_index34).Int())
		var _148_v101 _dafny.Map
		_ = _148_v101
		_148_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _149_v102 _dafny.Map
		_ = _149_v102
		_149_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v101, (_138_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v91), 0))).Int()).(bool))
		_149_v102 = (_149_v102).Update(_148_v101, (_dafny.IntOfUint32((_139_v92).Cardinality())).Cmp(p1) == 0)
		r0 = p1
	}
	var _150_v103 _dafny.MultiSet
	_ = _150_v103
	_150_v103 = _dafny.MultiSetOf(p1)
	var _151_v105 _dafny.CodePoint
	_ = _151_v105
	_151_v105 = _dafny.CodePoint('a')
	var _152_v106 _dafny.Sequence
	_ = _152_v106
	_152_v106 = _dafny.SeqOf(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v105, p2)).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _153_v104 _dafny.CodePoint
			_153_v104 = interface{}(_compr_3).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v105, p2)).Contains(_153_v104) {
				_coll3.Add(_153_v104, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spyhcfqae")).Cardinality()))
			}
		}
		return _coll3.ToMap()
	}())
	var _154_v107 _dafny.Map
	_ = _154_v107
	_154_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v105, _dafny.IntOfInt64(844))
	var _155_v108 _dafny.Set
	_ = _155_v108
	_155_v108 = _dafny.SetOf(p1, (_dafny.Zero).Minus(p1), (_150_v103).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_152_v106, (Companion_Default___.SafeIndex(Companion_Default___.Fm4(globalState), _dafny.IntOfUint32((_152_v106).Cardinality()))).Uint32(), _154_v107)).Cardinality()), p1)
	var _156_v109 _dafny.Sequence
	_ = _156_v109
	_156_v109 = _dafny.SeqOf(Companion_Default___.Fm13(_dafny.IntOfInt64(503), globalState))
	var _157_v110 _dafny.Set
	_ = _157_v110
	_157_v110 = _dafny.SetOf(p0)
	var _158_v111 _dafny.MultiSet
	_ = _158_v111
	_158_v111 = _dafny.MultiSetOf(_157_v110, _dafny.SetOf(true, p2), _157_v110, _157_v110)
	var _159_v112 D5
	_ = _159_v112
	_159_v112 = Companion_D5_.Create_DC15_(_158_v111)
	var _160_v113 _dafny.Map
	_ = _160_v113
	_160_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.SetOf(p1), _155_v108, (_156_v109).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_156_v109).Cardinality()))).Uint32()).(_dafny.Set), _155_v108, _155_v108), _159_v112)
	_160_v113 = (_160_v113).Update(_156_v109, _159_v112)
	var _161_v114 _dafny.Array
	_ = _161_v114
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
	_ = _len0_7
	var _nw36 _dafny.Array
	_ = _nw36
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw36 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) bool = (func(_162_p2 bool) func(_dafny.Int) bool {
			return func(_163_i13 _dafny.Int) bool {
				return _162_p2
			}
		})(p2)
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
	_161_v114 = _nw36
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_161_v114), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _164_i12 _dafny.Int
		_164_i12 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_164_i12).Sign() != -1) && ((_164_i12).Cmp(_dafny.ArrayLen((_161_v114), 0)) < 0)) {
			(_161_v114).ArraySet1((func() bool {
				if p0 {
					return !(p2)
				}
				return (p1).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ldv")).Cardinality())) >= 0
			})(), (_164_i12).Int())
		}
	}
	var _165_i14 _dafny.Int
	_ = _165_i14
	_165_i14 = _dafny.Zero
	{
		for (_14_v0).IsProperSubsetOf(((_14_v0).Update(p2, Companion_Default___.Abs((_dafny.Zero).Minus(p1)))).Update(false, Companion_Default___.Abs(p1))) {
			{
				if (_165_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_165_i14 = (_165_i14).Plus(_dafny.One)
				r2 = (func() bool {
					if p2 {
						return p0
					}
					return !((func() bool {
						if p0 {
							return p2
						}
						return p2
					})())
				})()
				r2 = (func() bool {
					if p2 {
						return false
					}
					return p2
				})()
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _index35
				(_161_v114).ArraySet1(!(p0), (_index35).Int())
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _index36
				(_161_v114).ArraySet1(false, (_index36).Int())
				var _166_v115 D5
				_ = _166_v115
				_166_v115 = Companion_D5_.Create_DC16_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)).Cardinality()).Times(p1), p1)
				var _167_v116 _dafny.Sequence
				_ = _167_v116
				_167_v116 = _dafny.UnicodeSeqOfUtf8Bytes("gsrufvirc")
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _index37
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _index38
				var _rhs9 D5 = func(_pat_let2_0 D5) D5 {
					return func(_168_dt__update__tmp_h1 D5) D5 {
						return func(_pat_let3_0 _dafny.Int) D5 {
							return func(_169_dt__update_hcf24_h0 _dafny.Int) D5 {
								return Companion_D5_.Create_DC16_(_169_dt__update_hcf24_h0, (_168_dt__update__tmp_h1).Dtor_cf25())
							}(_pat_let3_0)
						}(_dafny.IntOfInt64(590))
					}(_pat_let2_0)
				}(Companion_D5_.Create_DC16_(p1, p1))
				_ = _rhs9
				var _rhs10 _dafny.Int = _dafny.IntOfUint32((_167_v116).Cardinality())
				_ = _rhs10
				var _rhs11 bool = !(p2)
				_ = _rhs11
				var _rhs12 bool = (_161_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))).Int()).(bool)
				_ = _rhs12
				var _lhs6 *GlobalState = globalState
				_ = _lhs6
				var _lhs7 _dafny.Array = _161_v114
				_ = _lhs7
				var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _lhs8
				var _lhs9 _dafny.Array = _161_v114
				_ = _lhs9
				var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_161_v114), 0))
				_ = _lhs10
				_166_v115 = _rhs9
				_lhs6.F0 = _rhs10
				(_lhs7).ArraySet1(_rhs11, (_lhs8).Int())
				(_lhs9).ArraySet1(_rhs12, (_lhs10).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _170_v117 *C0
	_ = _170_v117
	var _nw37 *C0 = New_C0_()
	_ = _nw37
	_nw37.Ctor__(_161_v114)
	_170_v117 = _nw37
	r0 = p1
	r1 = (_dafny.Zero).Minus((_dafny.IntOfInt64(937)).Plus(Companion_Default___.Fm4(globalState)))
	r2 = !(p0)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _171_v0 bool
	_ = _171_v0
	_171_v0 = true
	var _172_v1 _dafny.Sequence
	_ = _172_v1
	_172_v1 = _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer")
	var _173_v2 _dafny.Array
	_ = _173_v2
	var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
	_ = _nw38
	_173_v2 = _nw38
	var _174_globalState *GlobalState
	_ = _174_globalState
	var _nw39 *GlobalState = New_GlobalState_()
	_ = _nw39
	_nw39.Ctor__(_dafny.IntOfInt64(-982), _dafny.SeqOf(_171_v0), false, _dafny.Companion_Sequence_.Concatenate(_172_v1, _172_v1), _dafny.IntOfInt64(745), _172_v1, _dafny.IntOfInt64(292), _dafny.CodePoint('x'), _dafny.IntOfInt64(420), _dafny.Companion_Sequence_.Concatenate(_172_v1, _172_v1), _dafny.IntOfInt64(835), false, _dafny.IntOfInt64(-89), _dafny.IntOfInt64(488), _173_v2)
	_174_globalState = _nw39
	var _175_v3 _dafny.Sequence
	_ = _175_v3
	_175_v3 = _dafny.SeqOf(_171_v0, _171_v0)
	_175_v3 = _dafny.Companion_Sequence_.Concatenate(_175_v3, Companion_Default___.Fm0(_171_v0, _174_globalState))
	var _176_v4 _dafny.Sequence
	_ = _176_v4
	_176_v4 = _dafny.SeqOf(_175_v3)
	var _177_v5 _dafny.Map
	_ = _177_v5
	_177_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((true), _171_v0)
	var _178_v6 _dafny.Map
	_ = _178_v6
	_178_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal((_176_v4).Select((Companion_Default___.SafeIndex((_177_v5).Cardinality(), _dafny.IntOfUint32((_176_v4).Cardinality()))).Uint32()).(_dafny.Sequence), _175_v3), _172_v1)
	_178_v6 = (_178_v6).Update(Companion_Default___.Fm1(Companion_Default___.Fm2(_171_v0, _171_v0, _174_globalState), _dafny.UnicodeSeqOfUtf8Bytes("xqfstui"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_179_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	})), _174_globalState), _dafny.Companion_Sequence_.Concatenate(_172_v1, _172_v1))
	var _180_i1 _dafny.Int
	_ = _180_i1
	_180_i1 = _dafny.Zero
	{
		for (func() bool {
			if _171_v0 {
				return _171_v0
			}
			return _171_v0
		})() {
			{
				if (_180_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_180_i1 = (_180_i1).Plus(_dafny.One)
				var _181_v7 _dafny.Int
				_ = _181_v7
				_181_v7 = _dafny.IntOfInt64(795)
				var _182_v8 _dafny.Int
				_ = _182_v8
				var _183_v9 _dafny.Int
				_ = _183_v9
				var _184_v10 bool
				_ = _184_v10
				var _out0 _dafny.Int
				_ = _out0
				var _out1 _dafny.Int
				_ = _out1
				var _out2 bool
				_ = _out2
				_out0, _out1, _out2 = Companion_Default___.M0(!(_171_v0), (_dafny.Zero).Minus(_181_v7), _171_v0, _174_globalState)
				_182_v8 = _out0
				_183_v9 = _out1
				_184_v10 = _out2
				var _185_v11 _dafny.Array
				_ = _185_v11
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw40
				_185_v11 = _nw40
				var _186_v12 _dafny.Array
				_ = _186_v12
				var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw41
				_186_v12 = _nw41
				var _nwElement0_5 _dafny.Array = _186_v12
				_ = _nwElement0_5
				var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(27))
				_ = _nw42
				(_nw42).ArraySet1(_nwElement0_5, 0)
				(_nw42).ArraySet1(_186_v12, 1)
				(_nw42).ArraySet1(_186_v12, 2)
				(_nw42).ArraySet1(_186_v12, 3)
				(_nw42).ArraySet1(_186_v12, 4)
				(_nw42).ArraySet1(_186_v12, 5)
				(_nw42).ArraySet1(_186_v12, 6)
				(_nw42).ArraySet1(_186_v12, 7)
				(_nw42).ArraySet1(_186_v12, 8)
				(_nw42).ArraySet1(_186_v12, 9)
				(_nw42).ArraySet1(_186_v12, 10)
				(_nw42).ArraySet1(_186_v12, 11)
				(_nw42).ArraySet1(_186_v12, 12)
				(_nw42).ArraySet1(_186_v12, 13)
				(_nw42).ArraySet1(_186_v12, 14)
				(_nw42).ArraySet1(_186_v12, 15)
				(_nw42).ArraySet1(_186_v12, 16)
				(_nw42).ArraySet1(_186_v12, 17)
				(_nw42).ArraySet1(_186_v12, 18)
				(_nw42).ArraySet1(_186_v12, 19)
				(_nw42).ArraySet1(_186_v12, 20)
				(_nw42).ArraySet1(_186_v12, 21)
				(_nw42).ArraySet1(_186_v12, 22)
				(_nw42).ArraySet1(_186_v12, 23)
				(_nw42).ArraySet1(_186_v12, 24)
				(_nw42).ArraySet1(_186_v12, 25)
				(_nw42).ArraySet1(_186_v12, 26)
				_185_v11 = _nw42
				var _187_v13 _dafny.Array
				_ = _187_v13
				var _nw43 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw43
				_187_v13 = _nw43
				var _188_v14 bool
				_ = _188_v14
				_188_v14 = _171_v0
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_187_v13), 0))
				_ = _index39
				(_187_v13).ArraySet1(_188_v14, (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_187_v13), 0))
				_ = _index40
				(_187_v13).ArraySet1(_188_v14, (_index40).Int())
				var _rhs13 _dafny.Int = _181_v7
				_ = _rhs13
				var _rhs14 bool = _184_v10
				_ = _rhs14
				_181_v7 = _rhs13
				_184_v10 = _rhs14
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _189_v15 _dafny.Int
	_ = _189_v15
	_189_v15 = _dafny.IntOfInt64(261)
	(_174_globalState).F0 = (_189_v15).Times(_189_v15)
	var _190_v16 bool
	_ = _190_v16
	_190_v16 = !(_171_v0)
	_190_v16 = _190_v16
	(_174_globalState).F0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
		if _171_v0 {
			return _189_v15
		}
		return _189_v15
	})())), (_189_v15).Minus(_189_v15))
	var _191_i2 _dafny.Int
	_ = _191_i2
	_191_i2 = _dafny.Zero
	{
		for _171_v0 {
			{
				if (_191_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_191_i2 = (_191_i2).Plus(_dafny.One)
				var _192_v17 _dafny.Array
				_ = _192_v17
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_8
				var _nw44 _dafny.Array
				_ = _nw44
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw44 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) bool = (func(_193_v0 bool) func(_dafny.Int) bool {
						return func(_194_i3 _dafny.Int) bool {
							return _193_v0
						}
					})(_171_v0)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw44 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw44).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw44).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_192_v17 = _nw44
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_192_v17), 0))
				_ = _index41
				(_192_v17).ArraySet1(_171_v0, (_index41).Int())
				var _195_v18 _dafny.Array
				_ = _195_v18
				var _nwElement0_6 bool = _190_v16
				_ = _nwElement0_6
				var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
				_ = _nw45
				(_nw45).ArraySet1(_nwElement0_6, 0)
				_195_v18 = _nw45
				var _196_v19 _dafny.Map
				_ = _196_v19
				_196_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v18, _189_v15)
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_192_v17), 0))
				_ = _index42
				(_192_v17).ArraySet1((_196_v19).Equals(_196_v19), (_index42).Int())
				var _197_v20 _dafny.CodePoint
				_ = _197_v20
				_197_v20 = _dafny.CodePoint('b')
				var _rhs15 _dafny.Int = _189_v15
				_ = _rhs15
				var _rhs16 _dafny.Int = Companion_Default___.SafeModuloInt(_189_v15, _189_v15)
				_ = _rhs16
				var _rhs17 bool = !((_192_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_192_v17), 0))).Int()).(bool))
				_ = _rhs17
				var _rhs18 _dafny.Sequence = _172_v1
				_ = _rhs18
				var _rhs19 _dafny.CodePoint = _dafny.CodePoint('c')
				_ = _rhs19
				var _lhs11 *GlobalState = _174_globalState
				_ = _lhs11
				var _lhs12 *GlobalState = _174_globalState
				_ = _lhs12
				_lhs11.F0 = _rhs15
				_189_v15 = _rhs16
				_171_v0 = _rhs17
				_lhs12.F5 = _rhs18
				_197_v20 = _rhs19
				var _198_v21 _dafny.Set
				_ = _198_v21
				_198_v21 = _dafny.SetOf(_dafny.SeqOf((_175_v3).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm3(_189_v15, _189_v15, _189_v15, _174_globalState)).Cardinality(), _dafny.IntOfUint32((_175_v3).Cardinality()))).Uint32()).(bool)), _175_v3)
				var _199_v22 _dafny.Map
				_ = _199_v22
				_199_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_198_v21).Cardinality(), _dafny.IntOfInt64(426)), _189_v15)
				_199_v22 = ((_199_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _189_v15))).Merge(_199_v22)
				var _200_v23 _dafny.MultiSet
				_ = _200_v23
				_200_v23 = _dafny.MultiSetOf(_189_v15)
				var _201_v24 _dafny.Map
				_ = _201_v24
				_201_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v23).Update(_189_v15, Companion_Default___.Abs(_189_v15)), _189_v15)
				_175_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_175_v3, (Companion_Default___.SafeIndex((_201_v24).Cardinality(), _dafny.IntOfUint32((_175_v3).Cardinality()))).Uint32(), _171_v0), _175_v3)
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _202_v25 _dafny.Array
	_ = _202_v25
	var _nw46 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw46
	_202_v25 = _nw46
	var _203_v26 _dafny.Set
	_ = _203_v26
	_203_v26 = _dafny.SetOf(_202_v25, _202_v25)
	var _204_v27 _dafny.Map
	_ = _204_v27
	_204_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_174_globalState), _203_v26)
	var _205_v28 _dafny.Array
	_ = _205_v28
	var _nwElement0_7 _dafny.Set = _203_v26
	_ = _nwElement0_7
	var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(25))
	_ = _nw47
	(_nw47).ArraySet1(_nwElement0_7, 0)
	(_nw47).ArraySet1(_dafny.SetOf(_202_v25, _202_v25), 1)
	(_nw47).ArraySet1(_203_v26, 2)
	(_nw47).ArraySet1(_dafny.SetOf(_202_v25), 3)
	(_nw47).ArraySet1((_203_v26).Union(_203_v26), 4)
	(_nw47).ArraySet1(_dafny.SetOf(_202_v25), 5)
	(_nw47).ArraySet1(_203_v26, 6)
	(_nw47).ArraySet1(_203_v26, 7)
	(_nw47).ArraySet1(_dafny.SetOf(_202_v25, _202_v25, _202_v25), 8)
	(_nw47).ArraySet1(_203_v26, 9)
	(_nw47).ArraySet1((_203_v26).Intersection(_203_v26), 10)
	(_nw47).ArraySet1(_203_v26, 11)
	(_nw47).ArraySet1(_203_v26, 12)
	(_nw47).ArraySet1(_203_v26, 13)
	(_nw47).ArraySet1(_dafny.SetOf(_202_v25), 14)
	(_nw47).ArraySet1(_203_v26, 15)
	(_nw47).ArraySet1(_203_v26, 16)
	(_nw47).ArraySet1(_203_v26, 17)
	(_nw47).ArraySet1((_dafny.SetOf(_202_v25, _202_v25)).Union(_203_v26), 18)
	(_nw47).ArraySet1(_203_v26, 19)
	(_nw47).ArraySet1(_203_v26, 20)
	(_nw47).ArraySet1(((func() _dafny.Set {
		if (_204_v27).Contains(_189_v15) {
			return (_204_v27).Get(_189_v15).(_dafny.Set)
		}
		return _203_v26
	})()).Union(_203_v26), 21)
	(_nw47).ArraySet1(_203_v26, 22)
	(_nw47).ArraySet1((func() _dafny.Set {
		if (_175_v3).Select((Companion_Default___.SafeIndex(_189_v15, _dafny.IntOfUint32((_175_v3).Cardinality()))).Uint32()).(bool) {
			return _203_v26
		}
		return _203_v26
	})(), 23)
	(_nw47).ArraySet1(_203_v26, 24)
	_205_v28 = _nw47
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_205_v28), 0))
	_ = _index43
	(_205_v28).ArraySet1(_203_v26, (_index43).Int())
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_205_v28), 0))
	_ = _index44
	(_205_v28).ArraySet1(_203_v26, (_index44).Int())
	var _source2 bool = _190_v16
	_ = _source2
	var _206___mcc_h0 bool = _source2
	_ = _206___mcc_h0
	var _207_cf0 bool = _206___mcc_h0
	_ = _207_cf0
	_178_v6 = (_178_v6).Update((!(_171_v0)) == (_171_v0), _dafny.Companion_Sequence_.Concatenate(_172_v1, _172_v1))
	var _208_v29 _dafny.Set
	_ = _208_v29
	_208_v29 = _dafny.SetOf(_207_cf0)
	var _209_v30 _dafny.Array
	_ = _209_v30
	var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
	_ = _nw48
	_209_v30 = _nw48
	var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))
	_ = _index45
	(_209_v30).ArraySet1(_189_v15, (_index45).Int())
	var _210_v31 _dafny.Set
	_ = _210_v31
	_210_v31 = _dafny.SetOf((_dafny.Zero).Minus(_189_v15))
	var _211_v32 _dafny.MultiSet
	_ = _211_v32
	_211_v32 = _dafny.MultiSetOf(_172_v1)
	var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))
	_ = _index46
	var _rhs20 _dafny.Set = Companion_Default___.Fm5(_189_v15, ((_210_v31).Cardinality()).Cmp(_189_v15) < 0, _dafny.Companion_Sequence_.Concatenate(_172_v1, _172_v1), (_211_v32).Union(_211_v32), _174_globalState)
	_ = _rhs20
	var _rhs21 bool = _171_v0
	_ = _rhs21
	var _rhs22 _dafny.Int = (_189_v15).Plus(Companion_Default___.Fm4(_174_globalState))
	_ = _rhs22
	var _lhs13 _dafny.Array = _209_v30
	_ = _lhs13
	var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))
	_ = _lhs14
	_208_v29 = _rhs20
	_207_cf0 = _rhs21
	(_lhs13).ArraySet1(_rhs22, (_lhs14).Int())
	var _212_v33 _dafny.Map
	_ = _212_v33
	_212_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int), _171_v0)
	var _213_v34 _dafny.CodePoint
	_ = _213_v34
	_213_v34 = _dafny.CodePoint('f')
	var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))
	_ = _index47
	var _rhs23 _dafny.Int = _dafny.IntOfInt64(373)
	_ = _rhs23
	var _rhs24 bool = ((_dafny.SetOf(_207_cf0)).Intersection(_208_v29)).IsProperSubsetOf(_208_v29)
	_ = _rhs24
	var _rhs25 bool = Companion_Default___.Fm1(_212_v33, _dafny.UnicodeSeqOfUtf8Bytes("mtng"), _172_v1, _174_globalState)
	_ = _rhs25
	var _rhs26 _dafny.Int = (_dafny.Zero).Minus(((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int)).Plus((func() _dafny.Int {
		if _207_cf0 {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_172_v1, (Companion_Default___.SafeIndex((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_172_v1).Cardinality()))).Uint32(), _213_v34)).Cardinality())
		}
		return (_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int)
	})()))
	_ = _rhs26
	var _rhs27 bool = (_189_v15).Cmp(_dafny.IntOfUint32((_172_v1).Cardinality())) >= 0
	_ = _rhs27
	var _lhs15 *GlobalState = _174_globalState
	_ = _lhs15
	var _lhs16 _dafny.Array = _209_v30
	_ = _lhs16
	var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))
	_ = _lhs17
	_lhs15.F0 = _rhs23
	_207_cf0 = _rhs24
	_171_v0 = _rhs25
	(_lhs16).ArraySet1(_rhs26, (_lhs17).Int())
	_207_cf0 = _rhs27
	if !(false) {
		var _214_v35 *C0
		_ = _214_v35
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__(_202_v25)
		_214_v35 = _nw49
		_171_v0 = (Companion_Default___.Fm4(_174_globalState)).Cmp((_dafny.IntOfUint32((_175_v3).Cardinality())).Times((_dafny.Zero).Minus((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int)))) >= 0
		var _215_v36 _dafny.Int
		_ = _215_v36
		var _216_v37 _dafny.Int
		_ = _216_v37
		var _217_v38 bool
		_ = _217_v38
		var _out3 _dafny.Int
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		var _out5 bool
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.M0(_207_cf0, (_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int), _207_cf0, _174_globalState)
		_215_v36 = _out3
		_216_v37 = _out4
		_217_v38 = _out5
		_177_v5 = (_177_v5).Merge((_177_v5).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_207_cf0), _171_v0)))
		var _218_v39 _dafny.Int
		_ = _218_v39
		var _219_v40 _dafny.Int
		_ = _219_v40
		var _220_v41 bool
		_ = _220_v41
		var _out6 _dafny.Int
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		var _out8 bool
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.M0((Companion_Default___.Fm4(_174_globalState)).Cmp(_215_v36) > 0, (_dafny.Zero).Minus((_216_v37).Plus((Companion_D1_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_221_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_222_i4 _dafny.Int) _dafny.CodePoint {
				return _221_v34
			}
		})(_213_v34)))).Cardinality()))).Dtor_cf1())), _207_cf0, _174_globalState)
		_218_v39 = _out6
		_219_v40 = _out7
		_220_v41 = _out8
	} else {
		var _223_v42 _dafny.Map
		_ = _223_v42
		_223_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_212_v33).Cardinality(), (_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int))
		var _224_v43 _dafny.Map
		_ = _224_v43
		_224_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_cf0, (_223_v42).Cardinality())
		var _225_v44 _dafny.Map
		_ = _225_v44
		_225_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _224_v43)
		var _226_v45 _dafny.Sequence
		_ = _226_v45
		_226_v45 = _dafny.SeqOf((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int), _189_v15, (_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int), ((func() _dafny.Map {
			if (_225_v44).Contains(_189_v15) {
				return (_225_v44).Get(_189_v15).(_dafny.Map)
			}
			return (_224_v43).Update(false, _189_v15)
		})()).Cardinality(), (_210_v31).Cardinality())
		var _227_v46 _dafny.Int
		_ = _227_v46
		var _228_v47 _dafny.Int
		_ = _228_v47
		var _229_v48 bool
		_ = _229_v48
		var _out9 _dafny.Int
		_ = _out9
		var _out10 _dafny.Int
		_ = _out10
		var _out11 bool
		_ = _out11
		_out9, _out10, _out11 = Companion_Default___.M0(_171_v0, (func() _dafny.Int {
			if _207_cf0 {
				return _dafny.IntOfUint32((_226_v45).Cardinality())
			}
			return (_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int)
		})(), !(((_177_v5).Cardinality()).Cmp((_209_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_209_v30), 0))).Int()).(_dafny.Int)) != 0), _174_globalState)
		_227_v46 = _out9
		_228_v47 = _out10
		_229_v48 = _out11
		_178_v6 = (_178_v6).Update(false, _dafny.UnicodeSeqOfUtf8Bytes("yhphth"))
		var _230_v49 _dafny.Int
		_ = _230_v49
		var _231_v50 _dafny.Int
		_ = _231_v50
		var _232_v51 bool
		_ = _232_v51
		var _out12 _dafny.Int
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		var _out14 bool
		_ = _out14
		_out12, _out13, _out14 = Companion_Default___.M0((true) && (_207_cf0), _dafny.IntOfInt64(825), Companion_Default___.Fm1(Companion_Default___.Fm2(_229_v48, _171_v0, _174_globalState), _172_v1, _172_v1, _174_globalState), _174_globalState)
		_230_v49 = _out12
		_231_v50 = _out13
		_232_v51 = _out14
		(_174_globalState).F5 = _172_v1
		(_174_globalState).F5 = _dafny.Companion_Sequence_.Concatenate(_172_v1, _dafny.Companion_Sequence_.Concatenate(_172_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(72))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_233_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_234_i5 _dafny.Int) _dafny.CodePoint {
				return _233_v34
			}
		})(_213_v34)))))
	}
	var _235_v52 *C0
	_ = _235_v52
	var _nw50 *C0 = New_C0_()
	_ = _nw50
	_nw50.Ctor__(_202_v25)
	_235_v52 = _nw50
	var _236_v53 _dafny.Set
	_ = _236_v53
	_236_v53 = _dafny.SetOf(_171_v0)
	var _237_v54 _dafny.MultiSet
	_ = _237_v54
	_237_v54 = _dafny.MultiSetOf(_172_v1, _172_v1)
	var _238_v55 D1
	_ = _238_v55
	_238_v55 = Companion_D1_.Create_DC4_((Companion_Default___.Fm5(_dafny.IntOfInt64(867), !(false), _172_v1, _237_v54, _174_globalState)).IsSubsetOf(_236_v53), _171_v0, _202_v25, _171_v0)
	var _source3 D1 = _238_v55
	_ = _source3
	if _source3.Is_DC2() {
		var _239___mcc_h1 _dafny.Int = _source3.Get_().(D1_DC2).Cf2
		_ = _239___mcc_h1
		var _240___mcc_h2 bool = _source3.Get_().(D1_DC2).Cf3
		_ = _240___mcc_h2
		var _241___mcc_h3 bool = _source3.Get_().(D1_DC2).Cf4
		_ = _241___mcc_h3
		var _242_cf4 bool = _241___mcc_h3
		_ = _242_cf4
		var _243_cf3 bool = _240___mcc_h2
		_ = _243_cf3
		var _244_cf2 _dafny.Int = _239___mcc_h1
		_ = _244_cf2
		var _245_v56 D1
		_ = _245_v56
		_245_v56 = Companion_D1_.Create_DC2_(_244_cf2, _171_v0, _171_v0)
		var _246_v57 _dafny.Map
		_ = _246_v57
		_246_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v56, _243_cf3)
		var _247_v58 _dafny.MultiSet
		_ = _247_v58
		_247_v58 = _dafny.MultiSetOf(_246_v57, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_(_244_cf2, false, _190_v16), _243_cf3))
		var _248_v59 _dafny.Sequence
		_ = _248_v59
		_248_v59 = _dafny.SeqOf(_dafny.IntOfInt64(357), (func() _dafny.Int {
			if (_247_v58).Contains(Companion_Default___.Fm6(_dafny.UnicodeSeqOfUtf8Bytes("yhshk"), _dafny.IntOfUint32((_175_v3).Cardinality()), _174_globalState)) {
				return (_247_v58).Multiplicity(Companion_Default___.Fm6(_dafny.UnicodeSeqOfUtf8Bytes("yhshk"), _dafny.IntOfUint32((_175_v3).Cardinality()), _174_globalState))
			}
			return _244_cf2
		})(), _244_cf2, _244_cf2, _189_v15)
		var _249_v60 _dafny.CodePoint
		_ = _249_v60
		_249_v60 = _dafny.CodePoint('d')
		var _250_v61 _dafny.Array
		_ = _250_v61
		var _nw51 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
		_ = _nw51
		_250_v61 = _nw51
		var _251_v62 _dafny.Map
		_ = _251_v62
		_251_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _250_v61)
		var _252_v63 _dafny.Array
		_ = _252_v63
		var _nwElement0_8 _dafny.Sequence = _248_v59
		_ = _nwElement0_8
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(19))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_8, 0)
		(_nw52).ArraySet1(_248_v59, 1)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_189_v15), _248_v59), 2)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_253_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_254_i6 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v15, true)).Cardinality()
			}
		})(_189_v15))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-511))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_255_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_256_i7 _dafny.Int) _dafny.Int {
				return _255_v15
			}
		})(_189_v15)))), 3)
		(_nw52).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_257_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_258_i8 _dafny.Int) _dafny.Int {
				return _257_v15
			}
		})(_189_v15))), 4)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59), 5)
		(_nw52).ArraySet1(_248_v59, 6)
		(_nw52).ArraySet1(_248_v59, 7)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59), 8)
		(_nw52).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(637))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_259_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_260_i9 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_259_v15)
			}
		})(_189_v15))), 9)
		(_nw52).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-277))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_261_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_262_i10 _dafny.Int) _dafny.Int {
				return _261_cf2
			}
		})(_244_cf2))), 10)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59), 11)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59), 12)
		(_nw52).ArraySet1(_248_v59, 13)
		(_nw52).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if true {
				return _248_v59
			}
			return Companion_Default___.Fm7(_243_cf3, _249_v60, _244_cf2, _174_globalState)
		})(), (Companion_Default___.SafeIndex(((_251_v62).Update(_189_v15, _250_v61)).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if true {
				return _248_v59
			}
			return Companion_Default___.Fm7(_243_cf3, _249_v60, _244_cf2, _174_globalState)
		})()).Cardinality()))).Uint32(), _244_cf2), 14)
		(_nw52).ArraySet1(_248_v59, 15)
		(_nw52).ArraySet1(_248_v59, 16)
		(_nw52).ArraySet1(_248_v59, 17)
		(_nw52).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_263_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_264_i11 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_263_v3).Cardinality())
			}
		})(_175_v3))), 18)
		_252_v63 = _nw52
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_252_v63), 0))
		_ = _index48
		(_252_v63).ArraySet1(_248_v59, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_252_v63), 0))
		_ = _index49
		(_252_v63).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59), (Companion_Default___.SafeIndex((_244_cf2).Times(_dafny.IntOfInt64(25)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_248_v59, _248_v59)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_172_v1).Cardinality())), (_index49).Int())
		var _rhs28 bool = !(!(_243_cf3))
		_ = _rhs28
		var _rhs29 *C0 = _235_v52
		_ = _rhs29
		var _rhs30 bool = _171_v0
		_ = _rhs30
		var _rhs31 _dafny.Int = _244_cf2
		_ = _rhs31
		var _rhs32 _dafny.Int = _189_v15
		_ = _rhs32
		_243_cf3 = _rhs28
		_235_v52 = _rhs29
		_171_v0 = _rhs30
		_189_v15 = _rhs31
		_189_v15 = _rhs32
		var _265_v64 _dafny.Int
		_ = _265_v64
		var _266_v65 _dafny.Int
		_ = _266_v65
		var _267_v66 bool
		_ = _267_v66
		var _out15 _dafny.Int
		_ = _out15
		var _out16 _dafny.Int
		_ = _out16
		var _out17 bool
		_ = _out17
		_out15, _out16, _out17 = Companion_Default___.M0((_175_v3).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_244_cf2), _dafny.IntOfUint32((_175_v3).Cardinality()))).Uint32()).(bool), _189_v15, _171_v0, _174_globalState)
		_265_v64 = _out15
		_266_v65 = _out16
		_267_v66 = _out17
		var _268_v67 _dafny.Map
		_ = _268_v67
		_268_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v66, _175_v3)
		_189_v15 = (((_268_v67).Merge(_268_v67)).Cardinality()).Times(_dafny.IntOfInt64(214))
	} else if _source3.Is_DC3() {
		(_174_globalState).F0 = (_dafny.Zero).Minus(_189_v15)
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen(((_235_v52).F15()), 0))
		_ = _index50
		((_235_v52).F15()).ArraySet1(_171_v0, (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen(((_235_v52).F15()), 0))
		_ = _index51
		((_235_v52).F15()).ArraySet1(_171_v0, (_index51).Int())
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen(((_235_v52).F15()), 0))
		_ = _index52
		((_235_v52).F15()).ArraySet1(_171_v0, (_index52).Int())
		var _269_v69 _dafny.Array
		_ = _269_v69
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_9
		var _nw53 _dafny.Array
		_ = _nw53
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw53 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Int = (func(_270_v15 _dafny.Int, _271_v52 *C0) func(_dafny.Int) _dafny.Int {
				return func(_272_i12 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_272_i12, (func() _dafny.Map {
						var _coll4 = _dafny.NewMapBuilder()
						_ = _coll4
						for _iter5 := _dafny.Iterate((_dafny.SeqOf(_270_v15)).Elements()); ; {
							_compr_4, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _273_v68 _dafny.Int
							_273_v68 = interface{}(_compr_4).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_270_v15), _273_v68) {
								_coll4.Add(Companion_Default___.SafeDivisionInt(_273_v68, _270_v15), ((_271_v52).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen(((_271_v52).F15()), 0))).Int()).(bool))
							}
						}
						return _coll4.ToMap()
					}()).Cardinality())
				}
			})(_189_v15, _235_v52)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw53 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw53).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw53).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_269_v69 = _nw53
		_269_v69 = _269_v69
	} else if _source3.Is_DC4() {
		var _274___mcc_h4 bool = _source3.Get_().(D1_DC4).Cf5
		_ = _274___mcc_h4
		var _275___mcc_h5 bool = _source3.Get_().(D1_DC4).Cf6
		_ = _275___mcc_h5
		var _276___mcc_h6 _dafny.Array = _source3.Get_().(D1_DC4).Cf7
		_ = _276___mcc_h6
		var _277___mcc_h7 bool = _source3.Get_().(D1_DC4).Cf8
		_ = _277___mcc_h7
		var _278_cf8 bool = _277___mcc_h7
		_ = _278_cf8
		var _279_cf7 _dafny.Array = _276___mcc_h6
		_ = _279_cf7
		var _280_cf6 bool = _275___mcc_h5
		_ = _280_cf6
		var _281_cf5 bool = _274___mcc_h4
		_ = _281_cf5
		var _282_v70 _dafny.Array
		_ = _282_v70
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
		_ = _nw54
		_282_v70 = _nw54
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_282_v70), 0))
		_ = _index53
		(_282_v70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cccpixh"), (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_282_v70), 0))
		_ = _index54
		(_282_v70).ArraySet1(_172_v1, (_index54).Int())
		if _278_cf8 {
			(_174_globalState).F0 = _189_v15
			var _283_v71 _dafny.Array
			_ = _283_v71
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw55
			_283_v71 = _nw55
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_283_v71), 0))
			_ = _index55
			(_283_v71).ArraySet1(Companion_Default___.Fm4(_174_globalState), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_283_v71), 0))
			_ = _index56
			(_283_v71).ArraySet1(Companion_Default___.Fm4(_174_globalState), (_index56).Int())
			var _284_v72 _dafny.MultiSet
			_ = _284_v72
			_284_v72 = _dafny.MultiSetOf(true, _281_cf5)
			var _rhs33 bool = !(!(((_283_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_283_v71), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(550))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_285_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_286_i13 _dafny.Int) _dafny.Int {
					return _285_v15
				}
			})(_189_v15)))).Cardinality())) == 0) || ((_238_v55).Dtor_cf5()))
			_ = _rhs33
			var _rhs34 bool = (_284_v72).IsProperSubsetOf(_284_v72)
			_ = _rhs34
			var _rhs35 _dafny.Sequence = (_282_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_282_v70), 0))).Int()).(_dafny.Sequence)
			_ = _rhs35
			_171_v0 = _rhs33
			_280_cf6 = _rhs34
			_172_v1 = _rhs35
			var _287_v73 _dafny.Map
			_ = _287_v73
			_287_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_283_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_283_v71), 0))).Int()).(_dafny.Int), _280_cf6)
			var _288_v74 _dafny.CodePoint
			_ = _288_v74
			_288_v74 = _dafny.CodePoint('b')
			_278_cf8 = Companion_Default___.Fm1(_287_v73, _172_v1, Companion_Default___.Fm8(_278_cf8, _288_v74, Companion_Default___.Fm9(_281_cf5, _174_globalState), _174_globalState), _174_globalState)
			var _289_v75 _dafny.Array
			_ = _289_v75
			var _nw56 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw56
			_289_v75 = _nw56
			var _290_v76 _dafny.Sequence
			_ = _290_v76
			_290_v76 = _dafny.SeqOf(_289_v75, _289_v75, _289_v75, _289_v75, _289_v75)
			_290_v76 = _290_v76
		} else {
			(_174_globalState).F0 = _189_v15
			(_174_globalState).F0 = (_189_v15).Plus(_189_v15)
			var _291_v77 _dafny.Int
			_ = _291_v77
			var _292_v78 _dafny.Int
			_ = _292_v78
			var _293_v79 bool
			_ = _293_v79
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			var _out20 bool
			_ = _out20
			_out18, _out19, _out20 = Companion_Default___.M0((_189_v15).Cmp(_189_v15) != 0, (_dafny.Zero).Minus(_189_v15), _278_cf8, _174_globalState)
			_291_v77 = _out18
			_292_v78 = _out19
			_293_v79 = _out20
			var _294_v80 *C0
			_ = _294_v80
			var _nw57 *C0 = New_C0_()
			_ = _nw57
			_nw57.Ctor__((_235_v52).F15())
			_294_v80 = _nw57
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen(((_235_v52).F15()), 0))
			_ = _index57
			((_235_v52).F15()).ArraySet1(_171_v0, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen(((_235_v52).F15()), 0))
			_ = _index58
			((_235_v52).F15()).ArraySet1(_281_cf5, (_index58).Int())
		}
		var _295_v81 _dafny.Map
		_ = _295_v81
		_295_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _278_cf8)
		_295_v81 = (_295_v81).Update(_189_v15, _171_v0)
		(_174_globalState).F5 = _172_v1
	} else if _source3.Is_DC1() {
		var _296___mcc_h8 _dafny.Int = _source3.Get_().(D1_DC1).Cf1
		_ = _296___mcc_h8
		var _297_cf1 _dafny.Int = _296___mcc_h8
		_ = _297_cf1
		(_174_globalState).F0 = Companion_Default___.SafeDivisionInt(_297_cf1, _297_cf1)
		var _298_v82 _dafny.Map
		_ = _298_v82
		_298_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _171_v0)
		var _299_v83 _dafny.Map
		_ = _299_v83
		_299_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, (_298_v82).Cardinality())
		var _300_v84 D1
		_ = _300_v84
		_300_v84 = Companion_D1_.Create_DC1_(_297_cf1)
		(_174_globalState).F0 = (func() _dafny.Int {
			if true {
				return _297_cf1
			}
			return ((_299_v83).Update((_dafny.Zero).Minus((_300_v84).Dtor_cf1()), _297_cf1)).Cardinality()
		})()
		_171_v0 = _171_v0
		(_174_globalState).F0 = _189_v15
	} else {
		var _301___mcc_h9 D1 = _source3.Get_().(D1_DC5).Cf9
		_ = _301___mcc_h9
		var _302_cf9 D1 = _301___mcc_h9
		_ = _302_cf9
		var _303_v85 _dafny.Map
		_ = _303_v85
		_303_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, _189_v15)
		(_174_globalState).F0 = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt(_189_v15, _189_v15)).Minus((func() _dafny.Int {
			if (_303_v85).Contains(_189_v15) {
				return (_303_v85).Get(_189_v15).(_dafny.Int)
			}
			return _189_v15
		})()))
		var _304_v86 *C0
		_ = _304_v86
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__(_202_v25)
		_304_v86 = _nw58
		var _305_v87 _dafny.Map
		_ = _305_v87
		_305_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_177_v5).Contains(_171_v0) {
				return (_177_v5).Get(_171_v0).(bool)
			}
			return _171_v0
		})(), _189_v15)
		var _306_v88 D2
		_ = _306_v88
		_306_v88 = Companion_D2_.Create_DC6_(_305_v87)
		_305_v87 = (((_306_v88).Dtor_cf10()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v0, _189_v15))).Merge(_305_v87)
		_171_v0 = _171_v0
	}
	var _307_v89 _dafny.Map
	_ = _307_v89
	_307_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v15, false)
	var _308_v90 _dafny.Map
	_ = _308_v90
	_308_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_307_v89, _172_v1, _dafny.UnicodeSeqOfUtf8Bytes("vvidmhm"), _174_globalState), Companion_Default___.Fm7(!(_171_v0), _dafny.CodePoint('p'), _189_v15, _174_globalState))
	var _309_v91 _dafny.Sequence
	_ = _309_v91
	_309_v91 = _dafny.SeqOf((_236_v53).Cardinality(), Companion_Default___.Fm4(_174_globalState), _189_v15)
	_308_v90 = (_308_v90).Update(_171_v0, _309_v91)
	var _310_v92 _dafny.Map
	_ = _310_v92
	_310_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xndmxbthy"), _172_v1)
	var _311_v93 _dafny.Map
	_ = _311_v93
	_311_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v92, _171_v0)
	_311_v93 = (_311_v93).Update(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter6 := _dafny.Iterate((_237_v54).Elements()); ; {
			_compr_5, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _312_v94 _dafny.Sequence
			_312_v94 = interface{}(_compr_5).(_dafny.Sequence)
			if (_237_v54).Contains(_312_v94) {
				_coll5.Add(_312_v94, _172_v1)
			}
		}
		return _coll5.ToMap()
	}(), (func() bool {
		if _171_v0 {
			return _171_v0
		}
		return _171_v0
	})())
	var _313_v95 _dafny.CodePoint
	_ = _313_v95
	_313_v95 = _dafny.CodePoint('p')
	var _314_v96 _dafny.MultiSet
	_ = _314_v96
	_314_v96 = _dafny.MultiSetOf(_171_v0, _171_v0, _171_v0)
	var _315_v97 _dafny.Sequence
	_ = _315_v97
	_315_v97 = _dafny.SeqOf(_314_v96)
	var _316_v98 _dafny.Map
	_ = _316_v98
	_316_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_313_v95, _171_v0, _dafny.IntOfUint32((_315_v97).Cardinality()), _174_globalState), (_235_v52).F15())
	var _317_v99 _dafny.Map
	_ = _317_v99
	_317_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v0, _189_v15)
	var _318_v100 D2
	_ = _318_v100
	_318_v100 = Companion_D2_.Create_DC6_(_317_v99)
	var _319_v101 _dafny.Sequence
	_ = _319_v101
	_319_v101 = _dafny.SeqOf(_318_v100, _318_v100)
	_316_v98 = (_316_v98).Update(_319_v101, (_235_v52).F15())
	var _320_v102 _dafny.Int
	_ = _320_v102
	var _321_v103 _dafny.Int
	_ = _321_v103
	var _322_v104 bool
	_ = _322_v104
	var _out21 _dafny.Int
	_ = _out21
	var _out22 _dafny.Int
	_ = _out22
	var _out23 bool
	_ = _out23
	_out21, _out22, _out23 = Companion_Default___.M0(_171_v0, _dafny.IntOfUint32((_175_v3).Cardinality()), (func() bool {
		if _171_v0 {
			return _171_v0
		}
		return false
	})(), _174_globalState)
	_320_v102 = _out21
	_321_v103 = _out22
	_322_v104 = _out23
	_322_v104 = !(_171_v0)
	_dafny.Print(_171_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_172_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_174_globalState).F1(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_globalState.F5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F9().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_175_v3, _dafny.SeqOf(true, false, true, true, false, false, true, false, true, true, true, true, false, false, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_176_v4, _dafny.SeqOf(_dafny.SeqOf(true, true, true, true, false, false, true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqerfyhqxeqer")).UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqerfyhqxeqer"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_180_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_189_v15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v16))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_191_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_202_v25).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v26).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_204_v27).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_205_v28).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_235_v52).F15()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v53).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v54).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer"), _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v55).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v55).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v55).Dtor_cf7()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v55).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_307_v89).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v90).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(589), _dafny.IntOfInt64(-1), _dafny.IntOfInt64(-720), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(44), _dafny.IntOfInt64(-2), _dafny.IntOfInt64(7))).UpdateUnsafe(false, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(4), _dafny.Zero))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_309_v91, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(4), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v92).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xndmxbthy"), _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_311_v93).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xndmxbthy"), _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer")), false).UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer"), _dafny.UnicodeSeqOfUtf8Bytes("fyhqxeqer")), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_313_v95)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_314_v96).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_315_v97, _dafny.SeqOf(_dafny.MultiSetOf(false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_316_v98).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_317_v99).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_318_v100).Dtor_cf10()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_319_v101, _dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)), Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_320_v102)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_321_v103)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_322_v104)
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

func (CompanionStruct_D0_) Default() bool {
	return false
}

func (_this D0) Dtor_cf0() bool {
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

type D1_DC2 struct {
	Cf2 _dafny.Int
	Cf3 bool
	Cf4 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 bool, Cf4 bool) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

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

type D1_DC4 struct {
	Cf5 bool
	Cf6 bool
	Cf7 _dafny.Array
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 bool, Cf6 bool, Cf7 _dafny.Array, Cf8 bool) D1 {
	return D1{D1_DC4{Cf5, Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Int
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Int) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
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
	return Companion_D1_.Create_DC2_(_dafny.Zero, false, false)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf9() D1 {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
		}
	case D1_DC3:
		{
			_, ok := other.Get_().(D1_DC3)
			return ok
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
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
	Cf11 bool
	Cf12 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool, Cf12 _dafny.Int) D2 {
	return D2{D2_DC7{Cf11, Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf10 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.Map) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf13 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 D2) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, _dafny.Zero)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf10() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf13() D2 {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
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
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
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

type D3_DC10 struct {
	Cf15 _dafny.Int
	Cf16 _dafny.Sequence
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf15 _dafny.Int, Cf16 _dafny.Sequence) D3 {
	return D3{D3_DC10{Cf15, Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf17 _dafny.Sequence
	Cf18 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf17 _dafny.Sequence, Cf18 bool) D3 {
	return D3{D3_DC11{Cf17, Cf18}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

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

type D3_DC12 struct {
	Cf19 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf19 D3) D3 {
	return D3{D3_DC12{Cf19}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero, _dafny.EmptySeq)
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D3_DC11).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) Dtor_cf19() D3 {
	return _this.Get_().(D3_DC12).Cf19
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Equals(data2.Cf16)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17) && data1.Cf18 == data2.Cf18
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf14.Equals(data2.Cf14)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
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
	Cf21 *C0
	Cf22 *C0
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 *C0, Cf22 *C0) D4 {
	return D4{D4_DC14{Cf21, Cf22}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf20 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf20 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_((*C0)(nil), (*C0)(nil))
}

func (_this D4) Dtor_cf21() *C0 {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() *C0 {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + data.Cf20.VerbatimString(true) + ")"
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
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D5_DC16 struct {
	Cf24 _dafny.Int
	Cf25 _dafny.Int
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf24 _dafny.Int, Cf25 _dafny.Int) D5 {
	return D5{D5_DC16{Cf24, Cf25}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC17 struct {
	Cf26 bool
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf26 bool) D5 {
	return D5{D5_DC17{Cf26}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC15 struct {
	Cf23 _dafny.MultiSet
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf23 _dafny.MultiSet) D5 {
	return D5{D5_DC15{Cf23}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_(_dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf24
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf25
}

func (_this D5) Dtor_cf26() bool {
	return _this.Get_().(D5_DC17).Cf26
}

func (_this D5) Dtor_cf23() _dafny.MultiSet {
	return _this.Get_().(D5_DC15).Cf23
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf26 == data2.Cf26
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D6_DC19 struct {
	Cf28 _dafny.CodePoint
	Cf29 bool
	Cf30 _dafny.Map
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf28 _dafny.CodePoint, Cf29 bool, Cf30 _dafny.Map) D6 {
	return D6{D6_DC19{Cf28, Cf29, Cf30}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC18 struct {
	Cf27 _dafny.Set
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf27 _dafny.Set) D6 {
	return D6{D6_DC18{Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC19_(_dafny.CodePoint('D'), false, _dafny.EmptyMap)
}

func (_this D6) Dtor_cf28() _dafny.CodePoint {
	return _this.Get_().(D6_DC19).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC19).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D6_DC19).Cf30
}

func (_this D6) Dtor_cf27() _dafny.Set {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30.Equals(data2.Cf30)
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
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

type D7_DC21 struct {
	Cf32 _dafny.Int
	Cf33 bool
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf32 _dafny.Int, Cf33 bool) D7 {
	return D7{D7_DC21{Cf32, Cf33}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC22 struct {
	Cf34 _dafny.Int
	Cf35 bool
	Cf36 _dafny.MultiSet
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf34 _dafny.Int, Cf35 bool, Cf36 _dafny.MultiSet) D7 {
	return D7{D7_DC22{Cf34, Cf35, Cf36}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC20 struct {
	Cf31 _dafny.Array
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf31 _dafny.Array) D7 {
	return D7{D7_DC20{Cf31}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(_dafny.Zero, false)
}

func (_this D7) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf32
}

func (_this D7) Dtor_cf33() bool {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf34
}

func (_this D7) Dtor_cf35() bool {
	return _this.Get_().(D7_DC22).Cf35
}

func (_this D7) Dtor_cf36() _dafny.MultiSet {
	return _this.Get_().(D7_DC22).Cf36
}

func (_this D7) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33 == data2.Cf33
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36)
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf31 == data2.Cf31
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

type D8_DC23 struct {
	Cf37 _dafny.MultiSet
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf37 _dafny.MultiSet) D8 {
	return D8{D8_DC23{Cf37}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D8) Dtor_cf37() _dafny.MultiSet {
	return _this.Get_().(D8_DC23).Cf37
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf37.Equals(data2.Cf37)
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

// Definition of datatype D9
type D9 struct {
	Data_D9_
}

func (_this D9) Get_() Data_D9_ {
	return _this.Data_D9_
}

type Data_D9_ interface {
	isD9()
}

type CompanionStruct_D9_ struct {
}

var Companion_D9_ = CompanionStruct_D9_{}

type D9_DC25 struct {
	Cf39 _dafny.Int
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf39 _dafny.Int) D9 {
	return D9{D9_DC25{Cf39}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf40 bool
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf40 bool) D9 {
	return D9{D9_DC26{Cf40}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf41 bool
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf41 bool) D9 {
	return D9{D9_DC27{Cf41}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC24 struct {
	Cf38 _dafny.Sequence
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf38 _dafny.Sequence) D9 {
	return D9{D9_DC24{Cf38}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(_dafny.Zero)
}

func (_this D9) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf39
}

func (_this D9) Dtor_cf40() bool {
	return _this.Get_().(D9_DC26).Cf40
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC27).Cf41
}

func (_this D9) Dtor_cf38() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf38
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf40 == data2.Cf40
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf41 == data2.Cf41
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf38.Equals(data2.Cf38)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D9) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D9)
	return ok && _this.Equals(typed)
}

func Type_D9_() _dafny.TypeDescriptor {
	return type_D9_{}
}

type type_D9_ struct {
}

func (_this type_D9_) Default() interface{} {
	return Companion_D9_.Default()
}

func (_this type_D9_) String() string {
	return "main.D9"
}
func (_this D9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D9{}

// End of datatype D9

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F5   _dafny.Sequence
	F14  _dafny.Array
	_f1  _dafny.Sequence
	_f2  bool
	_f3  _dafny.Sequence
	_f4  _dafny.Int
	_f6  _dafny.Int
	_f7  _dafny.CodePoint
	_f8  _dafny.Int
	_f9  _dafny.Sequence
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.Int
	_f13 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F5 = _dafny.EmptySeq
	_this.F14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.EmptySeq
	_this._f2 = false
	_this._f3 = _dafny.EmptySeq
	_this._f4 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.CodePoint('D')
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.EmptySeq
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 bool, f3 _dafny.Sequence, f4 _dafny.Int, f5 _dafny.Sequence, f6 _dafny.Int, f7 _dafny.CodePoint, f8 _dafny.Int, f9 _dafny.Sequence, f10 _dafny.Int, f11 bool, f12 _dafny.Int, f13 _dafny.Int, f14 _dafny.Array) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.CodePoint {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Sequence {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f15 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f15 _dafny.Array) {
	{
		(_this)._f15 = f15
	}
}
func (_this *C0) F15() _dafny.Array {
	{
		return _this._f15
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
