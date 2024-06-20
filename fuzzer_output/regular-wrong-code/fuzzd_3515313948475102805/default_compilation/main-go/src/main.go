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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	return (((_dafny.SetOf(Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uakfefor")).Cardinality())) == 0) || (false)
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	return ((Companion_D10_.Create_DC28_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("syud")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yfngghe")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-859))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()), _dafny.IntOfInt64(313))).Cardinality()), _dafny.IntOfInt64(420))).Cardinality()))).Dtor_cf38()).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Set, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.CodePoint {
	var _source0 D1 = Companion_D1_.Create_DC3_(false, !(false), !(false))
	_ = _source0
	if _source0.Is_DC3() {
		var _1___mcc_h0 bool = _source0.Get_().(D1_DC3).Cf2
		_ = _1___mcc_h0
		var _2___mcc_h1 bool = _source0.Get_().(D1_DC3).Cf3
		_ = _2___mcc_h1
		var _3___mcc_h2 bool = _source0.Get_().(D1_DC3).Cf4
		_ = _3___mcc_h2
		var _4_cf4 bool = _3___mcc_h2
		_ = _4_cf4
		var _5_cf3 bool = _2___mcc_h1
		_ = _5_cf3
		var _6_cf2 bool = _1___mcc_h0
		_ = _6_cf2
		return _dafny.CodePoint('y')
	} else {
		var _7___mcc_h3 _dafny.CodePoint = _source0.Get_().(D1_DC2).Cf1
		_ = _7___mcc_h3
		var _8_cf1 _dafny.CodePoint = _7___mcc_h3
		_ = _8_cf1
		return _8_cf1
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
		if !(true) {
			return false
		}
		return false
	})()), _dafny.CodePoint('t'))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(true), false)).Intersection(_dafny.SetOf(true, false))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("sua")
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC5_(!((true) && (!(!(true)))), _dafny.IntOfInt64(989), _dafny.IntOfInt64(149))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source1 D7 = Companion_D7_.Create_DC21_(_dafny.MultiSetOf(true), (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()), true)
	_ = _source1
	if _source1.Is_DC21() {
		var _9___mcc_h0 _dafny.MultiSet = _source1.Get_().(D7_DC21).Cf30
		_ = _9___mcc_h0
		var _10___mcc_h1 _dafny.Int = _source1.Get_().(D7_DC21).Cf31
		_ = _10___mcc_h1
		var _11___mcc_h2 bool = _source1.Get_().(D7_DC21).Cf32
		_ = _11___mcc_h2
		var _12_cf32 bool = _11___mcc_h2
		_ = _12_cf32
		var _13_cf31 _dafny.Int = _10___mcc_h1
		_ = _13_cf31
		var _14_cf30 _dafny.MultiSet = _9___mcc_h0
		_ = _14_cf30
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(935))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_15_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _dafny.UnicodeSeqOfUtf8Bytes("qgxrugvjl")))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xpq")))
	} else if _source1.Is_DC22() {
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("dyaxwe"), _dafny.UnicodeSeqOfUtf8Bytes("ppj"), _dafny.UnicodeSeqOfUtf8Bytes("pwwoqsahs"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-594))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_16_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), _dafny.UnicodeSeqOfUtf8Bytes("b"))
	} else if _source1.Is_DC23() {
		var _17___mcc_h3 _dafny.MultiSet = _source1.Get_().(D7_DC23).Cf33
		_ = _17___mcc_h3
		var _18_cf33 _dafny.MultiSet = _17___mcc_h3
		_ = _18_cf33
		return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ybamd")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qndkpncb"))))
	} else {
		var _19___mcc_h4 *C0 = _source1.Get_().(D7_DC20).Cf29
		_ = _19___mcc_h4
		var _20_cf29 *C0 = _19___mcc_h4
		_ = _20_cf29
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sca"))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false, false), _dafny.SeqOf(true, true)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(698)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return (Companion_D11_.Create_DC31_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(794))).Cardinality(), true))).Dtor_cf43()
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(922), false)
	})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-337), true))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(260)), _dafny.SetOf(!(!(!(true)))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-886)), _dafny.SetOf(false, true, true)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(272))).Cardinality()), _dafny.IntOfInt64(332), _dafny.IntOfInt64(-582), _dafny.IntOfInt64(968), _dafny.IntOfInt64(104)), _dafny.SetOf(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(74)), _dafny.SetOf(false))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf(true)).IsDisjointFrom(_dafny.SetOf(!(!(true)), true)) {
		return func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-156), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_21_i0 _dafny.Int) _dafny.Int {
				return ((Companion_D12_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(199), _dafny.CodePoint('f')))).Dtor_cf46()).Cardinality()
			}))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-761))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _22_v0 _dafny.Int
				_22_v0 = interface{}(_compr_0).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(-156), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_21_i0 _dafny.Int) _dafny.Int {
					return ((Companion_D12_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(199), _dafny.CodePoint('f')))).Dtor_cf46()).Cardinality()
				}))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-761))).Contains(_22_v0) {
					_coll0.Add((_22_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.IntOfInt64(950))
				}
			}
			return _coll0.ToMap()
		}()
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cgj"), _dafny.UnicodeSeqOfUtf8Bytes("fdhfk"))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-889))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_23_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(734), (func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_24_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(940)
			}))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _25_v1 _dafny.Int
				_25_v1 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_24_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(940)
				})), _25_v1) {
					_coll1.Add((_25_v1).Minus(_dafny.IntOfInt64(678)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-429))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg8 _dafny.Int) interface{} {
							return coer8(arg8)
						}
					}(func(_26_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))).Cardinality()))))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SetOf(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Map, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int, bool) {
	var r0 bool = false
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 bool = false
	_ = r3
	var _27_v0 _dafny.Map
	_ = _27_v0
	_27_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
	_27_v0 = (_27_v0).Update(p2, _dafny.IntOfInt64(776))
	r2 = p3
	var _28_v1 _dafny.MultiSet
	_ = _28_v1
	_28_v1 = _dafny.MultiSetOf(p2, p2, p2, !(true), p2)
	var _29_v2 _dafny.MultiSet
	_ = _29_v2
	_29_v2 = _dafny.MultiSetOf((_28_v1).Cardinality(), p3, p3)
	var _30_v3 *C2
	_ = _30_v3
	var _nw0 *C2 = New_C2_()
	_ = _nw0
	_nw0.Ctor__(_dafny.IntOfInt64(-881))
	_30_v3 = _nw0
	var _31_v4 _dafny.Map
	_ = _31_v4
	_31_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _30_v3)
	var _32_v5 _dafny.Sequence
	_ = _32_v5
	_32_v5 = _dafny.SeqOf(p3)
	var _33_v6 _dafny.Sequence
	_ = _33_v6
	_33_v6 = _dafny.SeqOf(_30_v3)
	var _34_v7 _dafny.Sequence
	_ = _34_v7
	_34_v7 = _dafny.SeqOf((func() _dafny.Int {
		if (_29_v2).Contains(p3) {
			return (_29_v2).Multiplicity(p3)
		}
		return (_dafny.Zero).Minus(p3)
	})(), Companion_Default___.SafeModuloInt((_31_v4).Cardinality(), p3), _dafny.IntOfUint32((_32_v5).Cardinality()), Companion_Default___.SafeModuloInt((_30_v3).F3(), _dafny.IntOfUint32((_33_v6).Cardinality())))
	_32_v5 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(446))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_35_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aymdpyag")).Cardinality())
	}))
	var _36_v8 _dafny.Sequence
	_ = _36_v8
	_36_v8 = _dafny.UnicodeSeqOfUtf8Bytes("wpx")
	_36_v8 = _dafny.Companion_Sequence_.Concatenate(_36_v8, _36_v8)
	var _37_v9 D7
	_ = _37_v9
	_37_v9 = Companion_D7_.Create_DC21_(_28_v1, p3, p2)
	var _pat_let_tv0 = p2
	_ = _pat_let_tv0
	var _pat_let_tv1 = p2
	_ = _pat_let_tv1
	var _pat_let_tv2 = p2
	_ = _pat_let_tv2
	if func(_source2 D7) bool {
		if _source2.Is_DC21() {
			var _38___mcc_h0 _dafny.MultiSet = _source2.Get_().(D7_DC21).Cf30
			_ = _38___mcc_h0
			var _39___mcc_h1 _dafny.Int = _source2.Get_().(D7_DC21).Cf31
			_ = _39___mcc_h1
			var _40___mcc_h2 bool = _source2.Get_().(D7_DC21).Cf32
			_ = _40___mcc_h2
			var _41_cf32 bool = _40___mcc_h2
			_ = _41_cf32
			var _42_cf31 _dafny.Int = _39___mcc_h1
			_ = _42_cf31
			var _43_cf30 _dafny.MultiSet = _38___mcc_h0
			_ = _43_cf30
			return _41_cf32
		} else if _source2.Is_DC22() {
			return _pat_let_tv0
		} else if _source2.Is_DC23() {
			var _44___mcc_h3 _dafny.MultiSet = _source2.Get_().(D7_DC23).Cf33
			_ = _44___mcc_h3
			var _45_cf33 _dafny.MultiSet = _44___mcc_h3
			_ = _45_cf33
			return _pat_let_tv1
		} else {
			var _46___mcc_h4 *C0 = _source2.Get_().(D7_DC20).Cf29
			_ = _46___mcc_h4
			var _47_cf29 *C0 = _46___mcc_h4
			_ = _47_cf29
			return _pat_let_tv2
		}
	}(_37_v9) {
		var _48_v10 _dafny.Array
		_ = _48_v10
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_49_p2 bool) func(_dafny.Int) bool {
				return func(_50_i1 _dafny.Int) bool {
					return _49_p2
				}
			})(p2)
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
		_48_v10 = _nw1
		_48_v10 = _48_v10
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_48_v10), 0))
		_ = _index0
		(_48_v10).ArraySet1(true, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_48_v10), 0))
		_ = _index1
		var _rhs0 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_30_v3).F3(), (_dafny.Zero).Minus(p3)))
		_ = _rhs0
		var _rhs1 bool = p2
		_ = _rhs1
		var _lhs0 _dafny.Array = _48_v10
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_48_v10), 0))
		_ = _lhs1
		r2 = _rhs0
		(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		var _51_v11 _dafny.CodePoint
		_ = _51_v11
		_51_v11 = _dafny.CodePoint('n')
		var _52_v12 _dafny.Set
		_ = _52_v12
		_52_v12 = _dafny.SetOf(p2)
		var _53_v13 _dafny.Map
		_ = _53_v13
		_53_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _52_v12)
		r2 = (((Companion_Default___.Fm23((_30_v3).F3(), _51_v11, (_48_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_48_v10), 0))).Int()).(bool), (_30_v3).F3(), globalState)).Merge(_53_v13)).Update(p2, _52_v12)).Cardinality()
		r2 = ((_30_v3).F3()).Plus((func() _dafny.Int {
			if (_48_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_48_v10), 0))).Int()).(bool) {
				return (_27_v0).Cardinality()
			}
			return (_30_v3).F3()
		})())
		r2 = (_30_v3).F3()
	} else {
		r2 = _dafny.IntOfInt64(-257)
		var _54_v14 _dafny.Sequence
		_ = _54_v14
		_54_v14 = _dafny.SeqOf(p2)
		r1 = (_54_v14).Select((Companion_Default___.SafeIndex(((_30_v3).F3()).Plus((_30_v3).F3()), _dafny.IntOfUint32((_54_v14).Cardinality()))).Uint32()).(bool)
		var _55_v15 _dafny.Array
		_ = _55_v15
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_1
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) D2 = (func(_56_p2 bool, _57_v8 _dafny.Sequence) func(_dafny.Int) D2 {
				return func(_58_i2 _dafny.Int) D2 {
					return (func() D2 {
						if _56_p2 {
							return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("mw")))
						}
						return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC4_(_57_v8))))
					})()
				}
			})(p2, _36_v8)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw2 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw2).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw2).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_55_v15 = _nw2
		var _59_v16 D2
		_ = _59_v16
		_59_v16 = Companion_D2_.Create_DC4_(_36_v8)
		var _60_v17 D2
		_ = _60_v17
		_60_v17 = Companion_D2_.Create_DC6_(_59_v16)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_55_v15), 0))
		_ = _index2
		(_55_v15).ArraySet1(_60_v17, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_55_v15), 0))
		_ = _index3
		(_55_v15).ArraySet1((func() D2 {
			if p2 {
				return _60_v17
			}
			return Companion_D2_.Create_DC6_(_59_v16)
		})(), (_index3).Int())
		var _61_v18 _dafny.Set
		_ = _61_v18
		_61_v18 = _dafny.SetOf(p2, p2, p2)
		var _62_v19 _dafny.Map
		_ = _62_v19
		_62_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_61_v18).Union(_dafny.SetOf(p2, p2, p2, true)), Companion_Default___.Fm12(globalState))
		_62_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(741))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_63_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))
		var _64_v20 _dafny.Array
		_ = _64_v20
		var _nwElement0_0 _dafny.Sequence = _dafny.SeqOf(p3)
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(2))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_36_v8).Cardinality()))), 1)
		_64_v20 = _nw3
		var _65_v21 *C0
		_ = _65_v21
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__(_64_v20)
		_65_v21 = _nw4
		var _66_v22 _dafny.Map
		_ = _66_v22
		_66_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v3, (func() *C0 {
			if true {
				return _65_v21
			}
			return _65_v21
		})())
		_66_v22 = (_66_v22).Update(_30_v3, _65_v21)
	}
	var _67_v23 _dafny.Array
	_ = _67_v23
	var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
	_ = _nw5
	_67_v23 = _nw5
	var _68_v24 _dafny.Map
	_ = _68_v24
	_68_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_67_v23), 0))
	_ = _index4
	(_67_v23).ArraySet1(_dafny.SeqOf(_68_v24, _68_v24, _68_v24, _68_v24, Companion_Default___.Fm19(p2, (_30_v3).Fm18(Companion_D2_.Create_DC5_(p2, (_30_v3).F3(), (_30_v3).F3()), globalState), globalState)), (_index4).Int())
	var _69_v25 _dafny.Array
	_ = _69_v25
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_2
	var _nw6 _dafny.Array
	_ = _nw6
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw6 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Int = (func(_70_v3 *C2) func(_dafny.Int) _dafny.Int {
			return func(_71_i4 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_71_i4, (_70_v3).F3())
			}
		})(_30_v3)
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
	_69_v25 = _nw6
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_69_v25), 0))
	_ = _index5
	(_69_v25).ArraySet1(p3, (_index5).Int())
	var _72_v26 _dafny.Sequence
	_ = _72_v26
	_72_v26 = _dafny.SeqOf((_68_v24).Merge(_68_v24), _68_v24, _68_v24)
	var _73_v27 *C1
	_ = _73_v27
	var _nw7 *C1 = New_C1_()
	_ = _nw7
	_nw7.Ctor__()
	_73_v27 = _nw7
	var _74_v28 _dafny.MultiSet
	_ = _74_v28
	_74_v28 = _dafny.MultiSetOf(_73_v27)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_67_v23), 0))
	_ = _index6
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_69_v25), 0))
	_ = _index7
	var _rhs2 _dafny.Sequence = _72_v26
	_ = _rhs2
	var _rhs3 _dafny.Int = ((func() _dafny.Map {
		if p2 {
			return _68_v24
		}
		return _68_v24
	})()).Cardinality()
	_ = _rhs3
	var _rhs4 bool = ((_74_v28).IsProperSubsetOf(_74_v28)) || (p2)
	_ = _rhs4
	var _lhs2 _dafny.Array = _67_v23
	_ = _lhs2
	var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_67_v23), 0))
	_ = _lhs3
	var _lhs4 _dafny.Array = _69_v25
	_ = _lhs4
	var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_69_v25), 0))
	_ = _lhs5
	(_lhs2).ArraySet1(_rhs2, (_lhs3).Int())
	(_lhs4).ArraySet1(_rhs3, (_lhs5).Int())
	r1 = _rhs4
	var _75_v29 _dafny.Sequence
	_ = _75_v29
	_75_v29 = _dafny.SeqOf(p2)
	r0 = ((_75_v29).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_75_v29).Cardinality()))).Uint32()).(bool)) && (p2)
	r1 = false
	r2 = (_dafny.Zero).Minus((_30_v3).F3())
	r3 = p2
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _76_v0 bool
	_ = _76_v0
	_76_v0 = false
	var _77_v1 _dafny.Int
	_ = _77_v1
	_77_v1 = _dafny.IntOfInt64(-894)
	var _78_globalState *GlobalState
	_ = _78_globalState
	var _nw8 *GlobalState = New_GlobalState_()
	_ = _nw8
	_nw8.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _77_v1), false)
	_78_globalState = _nw8
	var _79_v2 _dafny.Set
	_ = _79_v2
	_79_v2 = _dafny.SetOf(_76_v0)
	var _80_v3 D0
	_ = _80_v3
	_80_v3 = Companion_D0_.Create_DC0_((_79_v2).Cardinality())
	var _hi0 _dafny.Int = ((_80_v3).Dtor_cf0()).Minus(_77_v1)
	_ = _hi0
	for _81_i0 := _77_v1; _81_i0.Cmp(_hi0) < 0; _81_i0 = _81_i0.Plus(_dafny.One) {
		var _82_v4 _dafny.Array
		_ = _82_v4
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw9
		_82_v4 = _nw9
		var _83_v5 _dafny.Array
		_ = _83_v5
		var _nwElement0_1 _dafny.Int = _81_i0
		_ = _nwElement0_1
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_1, 0)
		_83_v5 = _nw10
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))
		_ = _index8
		(_82_v4).ArraySet1(_83_v5, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))
		_ = _index9
		(_82_v4).ArraySet1(_83_v5, (_index9).Int())
		var _84_v6 _dafny.Sequence
		_ = _84_v6
		_84_v6 = _dafny.SeqOf(_dafny.IntOfInt64(662))
		var _85_v7 _dafny.Sequence
		_ = _85_v7
		_85_v7 = _dafny.UnicodeSeqOfUtf8Bytes("t")
		var _86_v8 _dafny.MultiSet
		_ = _86_v8
		_86_v8 = _dafny.MultiSetOf(_76_v0)
		var _87_v9 _dafny.Map
		_ = _87_v9
		_87_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_i0).Times((func() _dafny.Int {
			if (_86_v8).Contains(_76_v0) {
				return (_86_v8).Multiplicity(_76_v0)
			}
			return _dafny.IntOfInt64(545)
		})()), _76_v0)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))
		_ = _index10
		var _rhs5 bool = Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_84_v6).Cardinality()), _81_i0), Companion_Default___.Fm1((_79_v2).Cardinality(), _81_i0, _dafny.IntOfUint32((_85_v7).Cardinality()), false, _78_globalState), _dafny.CodePoint('j'), _78_globalState)
		_ = _rhs5
		var _rhs6 bool = _76_v0
		_ = _rhs6
		var _rhs7 _dafny.Int = (_dafny.Zero).Minus((_87_v9).Cardinality())
		_ = _rhs7
		var _rhs8 _dafny.Array = _dafny.ArrayCastTo((_82_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))).Int()))
		_ = _rhs8
		var _lhs6 _dafny.Array = _82_v4
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))
		_ = _lhs7
		_76_v0 = _rhs5
		_76_v0 = _rhs6
		_77_v1 = _rhs7
		(_lhs6).ArraySet1(_rhs8, (_lhs7).Int())
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_82_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))).Int()))
		_ = _arr0
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_dafny.ArrayCastTo((_82_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))).Int()))), 0))
		_ = _index11
		(_arr0).ArraySet1((_81_i0).Plus((_dafny.Zero).Minus(_81_i0)), (_index11).Int())
		var _88_v10 _dafny.Sequence
		_ = _88_v10
		_88_v10 = _dafny.SeqOf(_76_v0, true)
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_82_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))).Int()))
		_ = _arr1
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_dafny.ArrayCastTo((_82_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_82_v4), 0))).Int()))), 0))
		_ = _index12
		(_arr1).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_88_v10).Cardinality()), _81_i0, (_81_i0).Times(_81_i0), (_81_i0).Cmp(_dafny.IntOfInt64(-356)) != 0, _78_globalState), (_index12).Int())
		var _89_v11 _dafny.Array
		_ = _89_v11
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_3
		var _nw11 _dafny.Array
		_ = _nw11
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw11 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_90_i1 _dafny.Int) bool {
				return true
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw11 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw11).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw11).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_89_v11 = _nw11
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_89_v11), 0))
		_ = _index13
		(_89_v11).ArraySet1(_76_v0, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_89_v11), 0))
		_ = _index14
		(_89_v11).ArraySet1(((_86_v8).Update(_76_v0, Companion_Default___.Abs((_dafny.Zero).Minus((_dafny.SetOf(_76_v0, _76_v0, _76_v0, !(_76_v0))).Cardinality())))).IsSubsetOf((_86_v8).Intersection(_86_v8)), (_index14).Int())
	}
	var _91_v12 _dafny.Array
	_ = _91_v12
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(2)
	_ = _len0_4
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = func(_92_i3 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeModuloInt(_92_i3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw12 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw12).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw12).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_91_v12 = _nw12
	for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_91_v12), 0))); ; {
		_guard_loop_0, _ok2 := _iter2()
		if !_ok2 {
			break
		}
		var _93_i2 _dafny.Int
		_93_i2 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_93_i2).Sign() != -1) && ((_93_i2).Cmp(_dafny.ArrayLen((_91_v12), 0)) < 0)) {
			(_91_v12).ArraySet1(Companion_Default___.SafeDivisionInt(_93_i2, _77_v1), (_93_i2).Int())
		}
	}
	var _94_v13 _dafny.MultiSet
	_ = _94_v13
	_94_v13 = _dafny.MultiSetOf(_77_v1, _77_v1)
	var _95_v14 D0
	_ = _95_v14
	_95_v14 = Companion_D0_.Create_DC1_()
	var _source3 D0 = (func() D0 {
		if ((_94_v13).Cardinality()).Cmp(_77_v1) >= 0 {
			return Companion_Default___.Fm2(_79_v2, _76_v0, _76_v0, _76_v0, _78_globalState)
		}
		return _95_v14
	})()
	_ = _source3
	if _source3.Is_DC1() {
		var _96_v15 _dafny.CodePoint
		_ = _96_v15
		_96_v15 = _dafny.CodePoint('l')
		var _97_v16 _dafny.Map
		_ = _97_v16
		_97_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _96_v15)
		var _98_v18 _dafny.Map
		_ = _98_v18
		_98_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, _96_v15)
		var _99_v19 bool
		_ = _99_v19
		var _100_v20 bool
		_ = _100_v20
		var _101_v21 _dafny.Int
		_ = _101_v21
		var _102_v22 bool
		_ = _102_v22
		var _out0 bool
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.Int
		_ = _out2
		var _out3 bool
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(_97_v16, func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter3 := _dafny.Iterate((_98_v18).Keys().Elements()); ; {
				_compr_2, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _103_v17 _dafny.MultiSet
				_103_v17 = interface{}(_compr_2).(_dafny.MultiSet)
				if (_98_v18).Contains(_103_v17) {
					_coll2.Add(_103_v17, _dafny.SetOf(_76_v0, _76_v0, _76_v0))
				}
			}
			return _coll2.ToMap()
		}(), false, _77_v1, _78_globalState)
		_99_v19 = _out0
		_100_v20 = _out1
		_101_v21 = _out2
		_102_v22 = _out3
		var _104_v23 _dafny.Map
		_ = _104_v23
		_104_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v15, _76_v0)
		var _105_v24 _dafny.Map
		_ = _105_v24
		_105_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v22, _101_v21)
		var _106_v25 _dafny.Map
		_ = _106_v25
		_106_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _77_v1)
		var _107_v26 _dafny.Sequence
		_ = _107_v26
		_107_v26 = _dafny.UnicodeSeqOfUtf8Bytes("ojlppbr")
		var _108_v27 _dafny.Map
		_ = _108_v27
		_108_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v26, _100_v20)
		var _109_v28 _dafny.MultiSet
		_ = _109_v28
		_109_v28 = _dafny.MultiSetOf(true)
		var _rhs9 _dafny.Int = (_101_v21).Minus(_dafny.IntOfInt64(151))
		_ = _rhs9
		var _rhs10 _dafny.Int = _101_v21
		_ = _rhs10
		var _rhs11 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_78_globalState), _99_v19)).Merge((_104_v23).Update(_96_v15, _76_v0))
		_ = _rhs11
		var _rhs12 _dafny.Map = ((_105_v24).Update(Companion_Default___.Fm0((func() _dafny.Int {
			if (_106_v25).Contains(_101_v21) {
				return (_106_v25).Get(_101_v21).(_dafny.Int)
			}
			return (_dafny.MultiSetOf(_101_v21)).Cardinality()
		})(), _77_v1, (Companion_D1_.Create_DC2_(_96_v15)).Dtor_cf1(), _78_globalState), _77_v1)).Update((func() bool {
			if (_108_v27).Contains(_107_v26) {
				return (_108_v27).Get(_107_v26).(bool)
			}
			return _102_v22
		})(), Companion_Default___.Fm1((_109_v28).Cardinality(), _77_v1, _dafny.IntOfInt64(2), _100_v20, _78_globalState))
		_ = _rhs12
		var _rhs13 bool = _100_v20
		_ = _rhs13
		var _lhs8 *GlobalState = _78_globalState
		_ = _lhs8
		_101_v21 = _rhs9
		_101_v21 = _rhs10
		_104_v23 = _rhs11
		_lhs8.F0 = _rhs12
		_99_v19 = _rhs13
		var _110_v29 _dafny.Sequence
		_ = _110_v29
		_110_v29 = _dafny.SeqOf(_dafny.IntOfInt64(-209))
		var _111_v30 _dafny.Sequence
		_ = _111_v30
		_111_v30 = _dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-91)), _101_v21), _dafny.Companion_Sequence_.Concatenate(_110_v29, _dafny.SeqOf(_dafny.IntOfInt64(333))))
		_101_v21 = _dafny.IntOfUint32((_111_v30).Cardinality())
		if !(_102_v22) {
			_77_v1 = _77_v1
			var _112_v31 _dafny.MultiSet
			_ = _112_v31
			_112_v31 = _dafny.MultiSetOf(_107_v26, _dafny.UnicodeSeqOfUtf8Bytes("t"))
			var _113_v32 _dafny.Sequence
			_ = _113_v32
			_113_v32 = _dafny.SeqOf(_76_v0, !(_102_v22))
			var _114_v33 _dafny.Map
			_ = _114_v33
			_114_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v26, (_106_v25).Merge((_106_v25).Update(((_112_v31).Update(_107_v26, Companion_Default___.Abs(_dafny.IntOfUint32((_113_v32).Cardinality())))).Cardinality(), (_110_v29).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_110_v29).Cardinality()))).Uint32()).(_dafny.Int))))
			_114_v33 = (_114_v33).Update(_107_v26, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _77_v1))
			_96_v15 = _96_v15
			var _115_v34 _dafny.MultiSet
			_ = _115_v34
			_115_v34 = _dafny.MultiSetOf(_110_v29, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_116_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(136)
			})))
			var _117_v35 _dafny.Map
			_ = _117_v35
			_117_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, Companion_Default___.Fm5((_dafny.Zero).Minus(_dafny.IntOfInt64(-664)), _102_v22, _99_v19, _dafny.IntOfInt64(-128), _78_globalState))
			var _118_v36 bool
			_ = _118_v36
			var _119_v37 bool
			_ = _119_v37
			var _120_v38 _dafny.Int
			_ = _120_v38
			var _121_v39 bool
			_ = _121_v39
			var _out4 bool
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 bool
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(Companion_Default___.Fm4(Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("th")).Cardinality()), (_115_v34).Cardinality(), _101_v21, Companion_Default___.Fm0((_94_v13).Cardinality(), _101_v21, _96_v15, _78_globalState), _78_globalState), _99_v19, _78_globalState), _117_v35, (_100_v20) == (_99_v19), _77_v1, _78_globalState)
			_118_v36 = _out4
			_119_v37 = _out5
			_120_v38 = _out6
			_121_v39 = _out7
			_77_v1 = Companion_Default___.SafeModuloInt(_101_v21, _dafny.IntOfInt64(-822))
		} else {
			_76_v0 = !(Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt(_77_v1, _77_v1), _101_v21, _96_v15, _78_globalState))
			var _122_v40 _dafny.Array
			_ = _122_v40
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_5
			var _nw13 _dafny.Array
			_ = _nw13
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw13 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_123_v0 bool) func(_dafny.Int) bool {
					return func(_124_i5 _dafny.Int) bool {
						return (_123_v0) == (false)
					}
				})(_76_v0)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw13 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw13).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw13).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_122_v40 = _nw13
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_91_v12), 0))
			_ = _index15
			(_91_v12).ArraySet1(_77_v1, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_91_v12), 0))
			_ = _index16
			var _rhs14 bool = _76_v0
			_ = _rhs14
			var _rhs15 _dafny.Array = _122_v40
			_ = _rhs15
			var _rhs16 bool = !(_76_v0) || (_102_v22)
			_ = _rhs16
			var _rhs17 _dafny.Int = _dafny.Zero
			_ = _rhs17
			var _rhs18 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_125_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_126_i6 _dafny.Int) _dafny.CodePoint {
					return _125_v15
				}
			})(_96_v15))), (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_127_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_128_i6 _dafny.Int) _dafny.CodePoint {
					return _127_v15
				}
			})(_96_v15)))).Cardinality()))).Uint32(), _96_v15)
			_ = _rhs18
			var _lhs9 _dafny.Array = _91_v12
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_91_v12), 0))
			_ = _lhs10
			_100_v20 = _rhs14
			_122_v40 = _rhs15
			_102_v22 = _rhs16
			(_lhs9).ArraySet1(_rhs17, (_lhs10).Int())
			_107_v26 = _rhs18
			var _rhs19 _dafny.Int = (_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int)
			_ = _rhs19
			var _rhs20 _dafny.MultiSet = _109_v28
			_ = _rhs20
			var _rhs21 bool = true
			_ = _rhs21
			_101_v21 = _rhs19
			_109_v28 = _rhs20
			_99_v19 = _rhs21
			var _129_v41 _dafny.Array
			_ = _129_v41
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw14
			_129_v41 = _nw14
			_129_v41 = _129_v41
			_101_v21 = (_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int)
		}
	} else {
		var _130___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
		_ = _130___mcc_h0
		var _131_cf0 _dafny.Int = _130___mcc_h0
		_ = _131_cf0
		if (_79_v2).IsSubsetOf(_79_v2) {
			var _132_v42 _dafny.CodePoint
			_ = _132_v42
			_132_v42 = _dafny.CodePoint('x')
			_76_v0 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-201))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_133_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_134_i7 _dafny.Int) _dafny.CodePoint {
					return _133_v42
				}
			})(_132_v42))), _132_v42)
			var _135_v43 *C1
			_ = _135_v43
			var _nw15 *C1 = New_C1_()
			_ = _nw15
			_nw15.Ctor__()
			_135_v43 = _nw15
			var _136_v44 *C1
			_ = _136_v44
			var _nw16 *C1 = New_C1_()
			_ = _nw16
			_nw16.Ctor__()
			_136_v44 = _nw16
			_131_cf0 = _131_cf0
			_76_v0 = _76_v0
		} else {
			var _137_v45 *C1
			_ = _137_v45
			var _nw17 *C1 = New_C1_()
			_ = _nw17
			_nw17.Ctor__()
			_137_v45 = _nw17
			var _138_v46 _dafny.Sequence
			_ = _138_v46
			_138_v46 = _dafny.SeqOf(_137_v45)
			var _139_v47 _dafny.Sequence
			_ = _139_v47
			_139_v47 = _dafny.SeqOf(_131_cf0)
			var _140_v48 _dafny.Array
			_ = _140_v48
			var _nwElement0_2 *C1 = _137_v45
			_ = _nwElement0_2
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_2, 0)
			(_nw18).ArraySet1(_137_v45, 1)
			(_nw18).ArraySet1(_137_v45, 2)
			(_nw18).ArraySet1(_137_v45, 3)
			(_nw18).ArraySet1(_137_v45, 4)
			(_nw18).ArraySet1(_137_v45, 5)
			(_nw18).ArraySet1(_137_v45, 6)
			(_nw18).ArraySet1(_137_v45, 7)
			(_nw18).ArraySet1(_137_v45, 8)
			(_nw18).ArraySet1((_138_v46).Select((Companion_Default___.SafeIndex((_139_v47).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_139_v47).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_138_v46).Cardinality()))).Uint32()).(*C1), 9)
			(_nw18).ArraySet1(_137_v45, 10)
			(_nw18).ArraySet1(_137_v45, 11)
			_140_v48 = _nw18
			var _141_v49 _dafny.Map
			_ = _141_v49
			_141_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v48, _91_v12)
			_91_v12 = (func() _dafny.Array {
				if (_141_v49).Contains(_140_v48) {
					return (_141_v49).Get(_140_v48).(_dafny.Array)
				}
				return _91_v12
			})()
			var _142_v50 _dafny.Array
			_ = _142_v50
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_6
			var _nw19 _dafny.Array
			_ = _nw19
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw19 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Sequence = (func(_143_v47 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_144_i8 _dafny.Int) _dafny.Sequence {
						return _143_v47
					}
				})(_139_v47)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw19 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw19).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw19).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_142_v50 = _nw19
			var _145_v51 *C0
			_ = _145_v51
			var _nw20 *C0 = New_C0_()
			_ = _nw20
			_nw20.Ctor__(_142_v50)
			_145_v51 = _nw20
			var _146_v52 D1
			_ = _146_v52
			_146_v52 = Companion_D1_.Create_DC3_(_76_v0, _76_v0, _76_v0)
			_146_v52 = _146_v52
			var _147_v53 _dafny.Sequence
			_ = _147_v53
			_147_v53 = _dafny.UnicodeSeqOfUtf8Bytes("ppn")
			var _148_v54 _dafny.CodePoint
			_ = _148_v54
			_148_v54 = _dafny.CodePoint('m')
			_147_v53 = _dafny.Companion_Sequence_.Update(_147_v53, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_131_cf0, _dafny.IntOfInt64(10)), _dafny.IntOfUint32((_147_v53).Cardinality()))).Uint32(), _148_v54)
			var _149_v55 T1
			_ = _149_v55
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__(_142_v50)
			_149_v55 = _nw21
			_149_v55 = _149_v55
		}
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_91_v12), 0))
		_ = _index17
		(_91_v12).ArraySet1(_131_cf0, (_index17).Int())
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_91_v12), 0))
		_ = _index18
		(_91_v12).ArraySet1(_131_cf0, (_index18).Int())
		var _150_v56 *C1
		_ = _150_v56
		var _nw22 *C1 = New_C1_()
		_ = _nw22
		_nw22.Ctor__()
		_150_v56 = _nw22
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_91_v12), 0))
		_ = _index19
		(_91_v12).ArraySet1(_77_v1, (_index19).Int())
	}
	_77_v1 = _dafny.IntOfInt64(882)
	var _151_v57 _dafny.Set
	_ = _151_v57
	_151_v57 = _dafny.SetOf(_79_v2)
	var _152_v58 D2
	_ = _152_v58
	_152_v58 = Companion_D2_.Create_DC5_(_76_v0, _77_v1, (_151_v57).Cardinality())
	var _153_v59 D2
	_ = _153_v59
	_153_v59 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(_152_v58))
	var _source4 D2 = _153_v59
	_ = _source4
	if _source4.Is_DC5() {
		var _154___mcc_h1 bool = _source4.Get_().(D2_DC5).Cf6
		_ = _154___mcc_h1
		var _155___mcc_h2 _dafny.Int = _source4.Get_().(D2_DC5).Cf7
		_ = _155___mcc_h2
		var _156___mcc_h3 _dafny.Int = _source4.Get_().(D2_DC5).Cf8
		_ = _156___mcc_h3
		var _157_cf8 _dafny.Int = _156___mcc_h3
		_ = _157_cf8
		var _158_cf7 _dafny.Int = _155___mcc_h2
		_ = _158_cf7
		var _159_cf6 bool = _154___mcc_h1
		_ = _159_cf6
		var _160_v60 _dafny.Map
		_ = _160_v60
		_160_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_cf8, _157_cf8)
		var _161_v61 _dafny.MultiSet
		_ = _161_v61
		_161_v61 = _dafny.MultiSetOf(_159_cf6)
		var _162_v62 D2
		_ = _162_v62
		_162_v62 = Companion_D2_.Create_DC5_((!(_76_v0)) || (_159_cf6), (func() _dafny.Int {
			if (_160_v60).Contains(Companion_Default___.Fm1(_157_cf8, _157_cf8, _157_cf8, _76_v0, _78_globalState)) {
				return (_160_v60).Get(Companion_Default___.Fm1(_157_cf8, _157_cf8, _157_cf8, _76_v0, _78_globalState)).(_dafny.Int)
			}
			return _158_cf7
		})(), (_161_v61).Cardinality())
		var _source5 D2 = _162_v62
		_ = _source5
		if _source5.Is_DC5() {
			var _163___mcc_h6 bool = _source5.Get_().(D2_DC5).Cf6
			_ = _163___mcc_h6
			var _164___mcc_h7 _dafny.Int = _source5.Get_().(D2_DC5).Cf7
			_ = _164___mcc_h7
			var _165___mcc_h8 _dafny.Int = _source5.Get_().(D2_DC5).Cf8
			_ = _165___mcc_h8
			var _166_cf8 _dafny.Int = _165___mcc_h8
			_ = _166_cf8
			var _167_cf7 _dafny.Int = _164___mcc_h7
			_ = _167_cf7
			var _168_cf6 bool = _163___mcc_h6
			_ = _168_cf6
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_91_v12), 0))
			_ = _index20
			(_91_v12).ArraySet1(Companion_Default___.SafeDivisionInt(_167_cf7, _166_cf8), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_91_v12), 0))
			_ = _index21
			(_91_v12).ArraySet1(_166_cf8, (_index21).Int())
			var _169_v63 _dafny.Set
			_ = _169_v63
			_169_v63 = _dafny.SetOf(_dafny.IntOfInt64(42))
			var _170_v64 _dafny.Sequence
			_ = _170_v64
			_170_v64 = _dafny.UnicodeSeqOfUtf8Bytes("btxl")
			var _171_v65 _dafny.Array
			_ = _171_v65
			var _nwElement0_3 bool = _76_v0
			_ = _nwElement0_3
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_3, 0)
			(_nw23).ArraySet1((_94_v13).IsDisjointFrom(_94_v13), 1)
			(_nw23).ArraySet1(_168_cf6, 2)
			(_nw23).ArraySet1(_76_v0, 3)
			(_nw23).ArraySet1(_76_v0, 4)
			(_nw23).ArraySet1((_169_v63).IsSubsetOf(_169_v63), 5)
			(_nw23).ArraySet1(false, 6)
			(_nw23).ArraySet1(true, 7)
			(_nw23).ArraySet1(_168_cf6, 8)
			(_nw23).ArraySet1(_dafny.Companion_Sequence_.Equal(_170_v64, _170_v64), 9)
			(_nw23).ArraySet1(_76_v0, 10)
			(_nw23).ArraySet1(_76_v0, 11)
			(_nw23).ArraySet1((_162_v62).Dtor_cf6(), 12)
			_171_v65 = _nw23
			_171_v65 = _171_v65
			var _172_v66 _dafny.Array
			_ = _172_v66
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw24
			_172_v66 = _nw24
			var _173_v67 _dafny.Sequence
			_ = _173_v67
			_173_v67 = _dafny.SeqOf(_172_v66, _172_v66, _172_v66)
			var _174_v68 _dafny.Set
			_ = _174_v68
			_174_v68 = _dafny.SetOf((_173_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_170_v64).Cardinality()), _dafny.IntOfUint32((_173_v67).Cardinality()))).Uint32()).(_dafny.Array), _172_v66, _172_v66, _172_v66, _172_v66)
			_174_v68 = _174_v68
			_76_v0 = _76_v0
		} else if _source5.Is_DC4() {
			var _175___mcc_h9 _dafny.Sequence = _source5.Get_().(D2_DC4).Cf5
			_ = _175___mcc_h9
			var _176_cf5 _dafny.Sequence = _175___mcc_h9
			_ = _176_cf5
			var _177_v69 _dafny.Array
			_ = _177_v69
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw25
			_177_v69 = _nw25
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_177_v69), 0))
			_ = _index22
			(_177_v69).ArraySet1(!(!(_159_cf6)), (_index22).Int())
			var _178_v70 *C1
			_ = _178_v70
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__()
			_178_v70 = _nw26
			var _179_v71 _dafny.Map
			_ = _179_v71
			_179_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v70, _177_v69)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_177_v69), 0))
			_ = _index23
			(_177_v69).ArraySet1(!(_179_v71).Contains(_178_v70), (_index23).Int())
			var _180_v72 _dafny.Map
			_ = _180_v72
			_180_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v12, _dafny.IntOfInt64(666))
			_180_v72 = (_180_v72).Update(_91_v12, _158_cf7)
			var _181_v73 _dafny.Map
			_ = _181_v73
			_181_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_cf6, _159_cf6)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_177_v69), 0))
			_ = _index24
			(_177_v69).ArraySet1((func() bool {
				if ((_181_v73).Cardinality()).Cmp(_158_cf7) > 0 {
					return _159_cf6
				}
				return _159_cf6
			})(), (_index24).Int())
			var _182_v74 _dafny.Array
			_ = _182_v74
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_7
			var _nw27 _dafny.Array
			_ = _nw27
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw27 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.MultiSet = (func(_183_cf5 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
					return func(_184_i9 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(_183_cf5)).Difference(_dafny.MultiSetOf(_183_cf5))
					}
				})(_176_cf5)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw27 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw27).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw27).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_182_v74 = _nw27
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_182_v74), 0))
			_ = _index25
			(_182_v74).ArraySet1(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-309))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_185_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), _176_cf5, _176_cf5), (_index25).Int())
			var _186_v75 _dafny.Sequence
			_ = _186_v75
			_186_v75 = _dafny.SeqOf(_159_cf6)
			var _187_v76 _dafny.Sequence
			_ = _187_v76
			_187_v76 = _dafny.SeqOf(_186_v75, Companion_Default___.Fm15(Companion_Default___.Fm1(_77_v1, _158_cf7, _158_cf7, (_177_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_177_v69), 0))).Int()).(bool), _78_globalState), _78_globalState), _186_v75, _186_v75)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_182_v74), 0))
			_ = _index26
			(_182_v74).ArraySet1(Companion_Default___.Fm14(_dafny.Companion_Sequence_.IsProperPrefixOf(_176_cf5, _176_cf5), _187_v76, false, _78_globalState), (_index26).Int())
		} else {
			var _188___mcc_h10 D2 = _source5.Get_().(D2_DC6).Cf9
			_ = _188___mcc_h10
			var _189_cf9 D2 = _188___mcc_h10
			_ = _189_cf9
			var _190_v77 _dafny.Array
			_ = _190_v77
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_8
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = (func(_191_cf6 bool) func(_dafny.Int) bool {
					return func(_192_i11 _dafny.Int) bool {
						return _191_cf6
					}
				})(_159_cf6)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw28 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw28).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw28).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_190_v77 = _nw28
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_190_v77), 0))
			_ = _index27
			(_190_v77).ArraySet1(_76_v0, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_190_v77), 0))
			_ = _index28
			(_190_v77).ArraySet1(!((_158_cf7).Cmp(_158_cf7) <= 0), (_index28).Int())
			_76_v0 = (_159_cf6) == (_76_v0)
			var _193_v78 *C1
			_ = _193_v78
			var _nw29 *C1 = New_C1_()
			_ = _nw29
			_nw29.Ctor__()
			_193_v78 = _nw29
			var _194_v79 _dafny.Sequence
			_ = _194_v79
			_194_v79 = _dafny.SeqOf(_77_v1)
			var _195_v80 _dafny.Map
			_ = _195_v80
			_195_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v12, _dafny.IntOfInt64(909)), Companion_Default___.Fm1(_157_cf8, Companion_Default___.Fm1((_dafny.MultiSetOf(_193_v78)).Cardinality(), _77_v1, _158_cf7, _159_cf6, _78_globalState), _dafny.IntOfUint32((_194_v79).Cardinality()), _159_cf6, _78_globalState))
			var _196_v81 _dafny.CodePoint
			_ = _196_v81
			_196_v81 = _dafny.CodePoint('t')
			var _197_v82 _dafny.Sequence
			_ = _197_v82
			_197_v82 = _dafny.SeqOf(_76_v0, _159_cf6, _76_v0)
			var _198_v83 _dafny.Map
			_ = _198_v83
			_198_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_196_v81, _dafny.IntOfUint32((_197_v82).Cardinality()))
			var _199_v84 _dafny.Sequence
			_ = _199_v84
			_199_v84 = _dafny.UnicodeSeqOfUtf8Bytes("gx")
			_158_cf7 = (func() _dafny.Int {
				if (_195_v80).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v12, (_198_v83).Cardinality())).Update(_91_v12, _dafny.IntOfUint32((_199_v84).Cardinality()))) {
					return (_195_v80).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v12, (_198_v83).Cardinality())).Update(_91_v12, _dafny.IntOfUint32((_199_v84).Cardinality()))).(_dafny.Int)
				}
				return _157_cf8
			})()
			var _200_v85 _dafny.Sequence
			_ = _200_v85
			_200_v85 = _dafny.SeqOf(_199_v84, _dafny.UnicodeSeqOfUtf8Bytes("enxbdbrux"))
			var _201_v86 _dafny.Array
			_ = _201_v86
			var _nwElement0_4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_200_v85, _200_v85)
			_ = _nwElement0_4
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_4, 0)
			(_nw30).ArraySet1(_200_v85, 1)
			(_nw30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(725))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_202_v84 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_203_i12 _dafny.Int) _dafny.Sequence {
					return _202_v84
				}
			})(_199_v84))), _200_v85), 2)
			(_nw30).ArraySet1((func() _dafny.Sequence {
				if !(false) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}((func(_204_v84 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_205_i13 _dafny.Int) _dafny.Sequence {
							return _204_v84
						}
					})(_199_v84)))
				}
				return _200_v85
			})(), 3)
			(_nw30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_200_v85, _200_v85), 4)
			(_nw30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_200_v85, _dafny.Companion_Sequence_.Update(_200_v85, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.IntOfUint32((_200_v85).Cardinality()))).Uint32(), _199_v84)), 5)
			(_nw30).ArraySet1(_200_v85, 6)
			(_nw30).ArraySet1(_dafny.SeqOf(_199_v84, _199_v84), 7)
			(_nw30).ArraySet1(_200_v85, 8)
			_201_v86 = _nw30
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_201_v86), 0))
			_ = _index29
			(_201_v86).ArraySet1(_200_v85, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_91_v12), 0))
			_ = _index30
			(_91_v12).ArraySet1((_158_cf7).Plus(_157_cf8), (_index30).Int())
			var _206_v87 _dafny.Array
			_ = _206_v87
			var _nw31 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
			_ = _nw31
			_206_v87 = _nw31
			var _207_v88 _dafny.Array
			_ = _207_v88
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_9
			var _nw32 _dafny.Array
			_ = _nw32
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw32 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = (func(_208_cf8 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_209_i14 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}((func(_210_cf8 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_211_i15 _dafny.Int) _dafny.Int {
								return _210_cf8
							}
						})(_208_cf8)))
					}
				})(_157_cf8)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw32 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw32).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw32).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_207_v88 = _nw32
			var _212_v89 *C0
			_ = _212_v89
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__(_207_v88)
			_212_v89 = _nw33
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_206_v87), 0))
			_ = _index31
			(_206_v87).ArraySet1(_212_v89, (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_91_v12), 0))
			_ = _index32
			(_91_v12).ArraySet1(_77_v1, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_201_v86), 0))
			_ = _index33
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_91_v12), 0))
			_ = _index34
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_206_v87), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_91_v12), 0))
			_ = _index36
			var _rhs22 _dafny.Sequence = _200_v85
			_ = _rhs22
			var _rhs23 _dafny.Int = Companion_Default___.SafeDivisionInt(_157_cf8, Companion_Default___.Fm1(_157_cf8, _77_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_213_v81 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_214_i16 _dafny.Int) _dafny.CodePoint {
					return _213_v81
				}
			})(_196_v81)))).Cardinality()), _76_v0, _78_globalState))
			_ = _rhs23
			var _rhs24 *C0 = _212_v89
			_ = _rhs24
			var _rhs25 bool = (_190_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_190_v77), 0))).Int()).(bool)
			_ = _rhs25
			var _rhs26 _dafny.Int = (_dafny.Zero).Minus(_158_cf7)
			_ = _rhs26
			var _lhs11 _dafny.Array = _201_v86
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_201_v86), 0))
			_ = _lhs12
			var _lhs13 _dafny.Array = _91_v12
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_91_v12), 0))
			_ = _lhs14
			var _lhs15 _dafny.Array = _206_v87
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_206_v87), 0))
			_ = _lhs16
			var _lhs17 _dafny.Array = _91_v12
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_91_v12), 0))
			_ = _lhs18
			(_lhs11).ArraySet1(_rhs22, (_lhs12).Int())
			(_lhs13).ArraySet1(_rhs23, (_lhs14).Int())
			(_lhs15).ArraySet1(_rhs24, (_lhs16).Int())
			_159_cf6 = _rhs25
			(_lhs17).ArraySet1(_rhs26, (_lhs18).Int())
		}
		var _215_v90 _dafny.Array
		_ = _215_v90
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_10
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Sequence = (func(_216_v1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_217_i17 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_216_v1)
				}
			})(_77_v1)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw34 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw34).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw34).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_215_v90 = _nw34
		var _218_v91 *C0
		_ = _218_v91
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__(_215_v90)
		_218_v91 = _nw35
		var _219_v92 _dafny.Sequence
		_ = _219_v92
		_219_v92 = _dafny.SeqOf(_218_v91)
		if (_94_v13).Equals(Companion_Default___.Fm16((_dafny.SetOf(_158_cf7, _158_cf7, _dafny.IntOfUint32((_219_v92).Cardinality()), _157_cf8)).Cardinality(), _78_globalState)) {
			var _220_v93 _dafny.Map
			_ = _220_v93
			_220_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_cf8, (_79_v2).IsSubsetOf(_79_v2))
			var _221_v94 *C1
			_ = _221_v94
			var _nw36 *C1 = New_C1_()
			_ = _nw36
			_nw36.Ctor__()
			_221_v94 = _nw36
			var _222_v95 _dafny.Map
			_ = _222_v95
			_222_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_221_v94, _157_cf8)
			var _223_v96 _dafny.Map
			_ = _223_v96
			_223_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v95, _79_v2)
			var _224_v97 _dafny.Sequence
			_ = _224_v97
			_224_v97 = _dafny.SeqOf(_77_v1)
			var _225_v98 _dafny.Map
			_ = _225_v98
			_225_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_cf8, _224_v97)
			var _226_v99 T1
			_ = _226_v99
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__(_218_v91.F2)
			_226_v99 = _nw37
			var _227_v100 D3
			_ = _227_v100
			_227_v100 = Companion_D3_.Create_DC8_(_159_cf6, Companion_D2_.Create_DC6_(_152_v58), _76_v0, _79_v2, _221_v94)
			var _228_v101 _dafny.Array
			_ = _228_v101
			var _nwElement0_5 _dafny.Set = (_79_v2).Intersection(_79_v2)
			_ = _nwElement0_5
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(24))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_5, 0)
			(_nw38).ArraySet1((_79_v2).Intersection(_79_v2), 1)
			(_nw38).ArraySet1((func() _dafny.Set {
				if (_223_v96).Contains((_222_v95).Update(_221_v94, _157_cf8)) {
					return (_223_v96).Get((_222_v95).Update(_221_v94, _157_cf8)).(_dafny.Set)
				}
				return _dafny.SetOf(_76_v0)
			})(), 2)
			(_nw38).ArraySet1(_79_v2, 3)
			(_nw38).ArraySet1(_79_v2, 4)
			(_nw38).ArraySet1(_dafny.SetOf(_159_cf6), 5)
			(_nw38).ArraySet1(_dafny.SetOf(_159_cf6, _76_v0, _76_v0), 6)
			(_nw38).ArraySet1(_dafny.SetOf(_76_v0, false, _159_cf6, _159_cf6), 7)
			(_nw38).ArraySet1((_79_v2).Union(Companion_Default___.Fm5(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_225_v98).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(736)), _226_v99)).Cardinality()) {
					return (_225_v98).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(736)), _226_v99)).Cardinality()).(_dafny.Sequence)
				}
				return _224_v97
			})()).Cardinality()), _76_v0, _159_cf6, _158_cf7, _78_globalState)), 8)
			(_nw38).ArraySet1((func() _dafny.Set {
				if _159_cf6 {
					return _79_v2
				}
				return _79_v2
			})(), 9)
			(_nw38).ArraySet1(_dafny.SetOf(_76_v0, false), 10)
			(_nw38).ArraySet1(_79_v2, 11)
			(_nw38).ArraySet1(_79_v2, 12)
			(_nw38).ArraySet1((_79_v2).Difference(_79_v2), 13)
			(_nw38).ArraySet1((func() _dafny.Set {
				if _76_v0 {
					return Companion_Default___.Fm5(_77_v1, _159_cf6, _76_v0, Companion_Default___.Fm1(_157_cf8, _158_cf7, _dafny.IntOfInt64(361), _159_cf6, _78_globalState), _78_globalState)
				}
				return _79_v2
			})(), 14)
			(_nw38).ArraySet1(_79_v2, 15)
			(_nw38).ArraySet1(((_227_v100).Dtor_cf14()).Union(_79_v2), 16)
			(_nw38).ArraySet1((_79_v2).Union(_79_v2), 17)
			(_nw38).ArraySet1(_79_v2, 18)
			(_nw38).ArraySet1(_79_v2, 19)
			(_nw38).ArraySet1(_79_v2, 20)
			(_nw38).ArraySet1(_dafny.SetOf(true, _76_v0, _159_cf6, _76_v0), 21)
			(_nw38).ArraySet1((_79_v2).Intersection(_79_v2), 22)
			(_nw38).ArraySet1(_79_v2, 23)
			_228_v101 = _nw38
			var _229_v102 _dafny.Sequence
			_ = _229_v102
			_229_v102 = _dafny.UnicodeSeqOfUtf8Bytes("cr")
			var _230_v103 _dafny.Map
			_ = _230_v103
			_230_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _76_v0)
			var _231_v104 D1
			_ = _231_v104
			_231_v104 = Companion_D1_.Create_DC3_(_76_v0, _76_v0, _159_cf6)
			var _232_v105 _dafny.Map
			_ = _232_v105
			_232_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v104, _76_v0)
			var _233_v106 _dafny.Sequence
			_ = _233_v106
			_233_v106 = _dafny.SeqOf(_232_v105)
			var _rhs27 _dafny.Map = _220_v93
			_ = _rhs27
			var _rhs28 _dafny.Int = _158_cf7
			_ = _rhs28
			var _rhs29 _dafny.Array = _228_v101
			_ = _rhs29
			var _rhs30 _dafny.Sequence = (func() _dafny.Sequence {
				if _76_v0 {
					return _229_v102
				}
				return _229_v102
			})()
			_ = _rhs30
			var _rhs31 bool = (func() bool {
				if (_230_v103).Contains(_76_v0) {
					return (_230_v103).Get(_76_v0).(bool)
				}
				return (_218_v91).Fm7(_233_v106, _76_v0, _157_cf8, _78_globalState)
			})()
			_ = _rhs31
			_220_v93 = _rhs27
			_77_v1 = _rhs28
			_228_v101 = _rhs29
			_229_v102 = _rhs30
			_159_cf6 = _rhs31
			_159_cf6 = _76_v0
			var _234_v107 _dafny.Array
			_ = _234_v107
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_11
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_235_cf6 bool) func(_dafny.Int) bool {
					return func(_236_i18 _dafny.Int) bool {
						return _235_cf6
					}
				})(_159_cf6)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw39 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw39).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw39).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_234_v107 = _nw39
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_234_v107), 0))
			_ = _index37
			(_234_v107).ArraySet1((_94_v13).IsSubsetOf(_94_v13), (_index37).Int())
			var _237_v108 _dafny.Sequence
			_ = _237_v108
			_237_v108 = _dafny.SeqOf(_159_cf6)
			var _238_v109 _dafny.CodePoint
			_ = _238_v109
			_238_v109 = _dafny.CodePoint('y')
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_234_v107), 0))
			_ = _index38
			var _rhs32 bool = !(_76_v0)
			_ = _rhs32
			var _rhs33 bool = Companion_Default___.Fm0(_dafny.IntOfUint32((_237_v108).Cardinality()), _157_cf8, _238_v109, _78_globalState)
			_ = _rhs33
			var _rhs34 D0 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_158_cf7))
			_ = _rhs34
			var _lhs19 _dafny.Array = _234_v107
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_234_v107), 0))
			_ = _lhs20
			(_lhs19).ArraySet1(_rhs32, (_lhs20).Int())
			_159_cf6 = _rhs33
			_80_v3 = _rhs34
			_221_v94 = (func() *C1 {
				if (_234_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_234_v107), 0))).Int()).(bool) {
					return (func() *C1 {
						if _159_cf6 {
							return _221_v94
						}
						return _221_v94
					})()
				}
				return _221_v94
			})()
			var _rhs35 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(698), Companion_Default___.Fm1(_157_cf8, (Companion_D3_.Create_DC9_(_158_cf7)).Dtor_cf16(), _157_cf8, _76_v0, _78_globalState)))
			_ = _rhs35
			var _rhs36 _dafny.Int = _77_v1
			_ = _rhs36
			_157_cf8 = _rhs35
			_77_v1 = _rhs36
		} else {
			var _239_v110 _dafny.Map
			_ = _239_v110
			_239_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_77_v1), _79_v2)
			var _240_v111 _dafny.Sequence
			_ = _240_v111
			_240_v111 = _dafny.SeqOf(_79_v2, (func() _dafny.Set {
				if (_239_v110).Contains((_79_v2).Cardinality()) {
					return (_239_v110).Get((_79_v2).Cardinality()).(_dafny.Set)
				}
				return _79_v2
			})(), Companion_Default___.Fm5(_77_v1, _76_v0, _159_cf6, _77_v1, _78_globalState))
			_240_v111 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(137))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_241_v2 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_242_i19 _dafny.Int) _dafny.Set {
					return _241_v2
				}
			})(_79_v2))), _240_v111), _240_v111)
			var _243_v112 _dafny.Array
			_ = _243_v112
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw40
			_243_v112 = _nw40
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_243_v112), 0))
			_ = _index39
			(_243_v112).ArraySet1(_76_v0, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_243_v112), 0))
			_ = _index40
			(_243_v112).ArraySet1((_158_cf7).Cmp(_157_cf8) > 0, (_index40).Int())
			var _244_v113 _dafny.Sequence
			_ = _244_v113
			_244_v113 = _dafny.SeqOf((_158_cf7).Cmp(_158_cf7) < 0, (_243_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_243_v112), 0))).Int()).(bool))
			_244_v113 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_159_cf6, _76_v0, true), _244_v113)
			var _245_v114 T0
			_ = _245_v114
			var _nw41 *C2 = New_C2_()
			_ = _nw41
			_nw41.Ctor__(_77_v1)
			_245_v114 = _nw41
			var _246_v115 _dafny.Map
			_ = _246_v115
			_246_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _245_v114)
			var _247_v116 _dafny.Sequence
			_ = _247_v116
			_247_v116 = _dafny.SeqOf((_246_v115).Merge(_246_v115))
			var _248_v117 _dafny.Sequence
			_ = _248_v117
			_248_v117 = _dafny.SeqOf(_247_v116, _247_v116)
			_247_v116 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_248_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-164), _dafny.IntOfUint32((_248_v117).Cardinality()))).Uint32()).(_dafny.Sequence), _247_v116), _247_v116)
			_243_v112 = _243_v112
		}
		var _249_v118 _dafny.Map
		_ = _249_v118
		_249_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_cf6, (func() bool {
			if _76_v0 {
				return _76_v0
			}
			return _159_cf6
		})())
		_249_v118 = (_249_v118).Update(_159_cf6, _76_v0)
		_76_v0 = _76_v0
	} else if _source4.Is_DC4() {
		var _250___mcc_h4 _dafny.Sequence = _source4.Get_().(D2_DC4).Cf5
		_ = _250___mcc_h4
		var _251_cf5 _dafny.Sequence = _250___mcc_h4
		_ = _251_cf5
		var _252_v119 _dafny.Array
		_ = _252_v119
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_12
		var _nw42 _dafny.Array
		_ = _nw42
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw42 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) D3 = (func(_253_v2 _dafny.Set) func(_dafny.Int) D3 {
				return func(_254_i20 _dafny.Int) D3 {
					return Companion_D3_.Create_DC7_(_253_v2)
				}
			})(_79_v2)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw42 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw42).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw42).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_252_v119 = _nw42
		var _255_v120 D3
		_ = _255_v120
		_255_v120 = Companion_D3_.Create_DC7_(_79_v2)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_252_v119), 0))
		_ = _index41
		(_252_v119).ArraySet1(_255_v120, (_index41).Int())
		var _256_v121 _dafny.Array
		_ = _256_v121
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
		_ = _nw43
		_256_v121 = _nw43
		var _257_v122 _dafny.Array
		_ = _257_v122
		var _nwElement0_6 _dafny.Array = _256_v121
		_ = _nwElement0_6
		var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(7))
		_ = _nw44
		(_nw44).ArraySet1(_nwElement0_6, 0)
		(_nw44).ArraySet1(_256_v121, 1)
		(_nw44).ArraySet1(_256_v121, 2)
		(_nw44).ArraySet1(_256_v121, 3)
		(_nw44).ArraySet1(_256_v121, 4)
		(_nw44).ArraySet1(_256_v121, 5)
		(_nw44).ArraySet1(_256_v121, 6)
		_257_v122 = _nw44
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_257_v122), 0))
		_ = _index42
		(_257_v122).ArraySet1(_256_v121, (_index42).Int())
		var _258_v123 _dafny.Map
		_ = _258_v123
		_258_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_256_v121, _256_v121, _256_v121), _91_v12)
		var _259_v124 _dafny.Set
		_ = _259_v124
		_259_v124 = _dafny.SetOf(_256_v121)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_252_v119), 0))
		_ = _index43
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_257_v122), 0))
		_ = _index44
		var _rhs37 _dafny.Array = (func() _dafny.Array {
			if (_258_v123).Contains((_259_v124).Difference(_259_v124)) {
				return (_258_v123).Get((_259_v124).Difference(_259_v124)).(_dafny.Array)
			}
			return _91_v12
		})()
		_ = _rhs37
		var _rhs38 D3 = _255_v120
		_ = _rhs38
		var _rhs39 _dafny.Array = _256_v121
		_ = _rhs39
		var _rhs40 _dafny.Int = ((_77_v1).Times(_77_v1)).Times(_77_v1)
		_ = _rhs40
		var _rhs41 bool = _76_v0
		_ = _rhs41
		var _lhs21 _dafny.Array = _252_v119
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_252_v119), 0))
		_ = _lhs22
		var _lhs23 _dafny.Array = _257_v122
		_ = _lhs23
		var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_257_v122), 0))
		_ = _lhs24
		_91_v12 = _rhs37
		(_lhs21).ArraySet1(_rhs38, (_lhs22).Int())
		(_lhs23).ArraySet1(_rhs39, (_lhs24).Int())
		_77_v1 = _rhs40
		_76_v0 = _rhs41
		var _260_v125 _dafny.Sequence
		_ = _260_v125
		_260_v125 = _dafny.SeqOf((_77_v1).Cmp(_77_v1) == 0, _76_v0)
		_260_v125 = _dafny.SeqOf(_76_v0, _76_v0, _76_v0)
		_76_v0 = _76_v0
		_77_v1 = (_dafny.Zero).Minus((((_dafny.Zero).Minus(_77_v1)).Minus(Companion_Default___.Fm1(_77_v1, _77_v1, _77_v1, _76_v0, _78_globalState))).Plus(_77_v1))
	} else {
		var _261___mcc_h5 D2 = _source4.Get_().(D2_DC6).Cf9
		_ = _261___mcc_h5
		var _262_cf9 D2 = _261___mcc_h5
		_ = _262_cf9
		var _263_v126 _dafny.Array
		_ = _263_v126
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_13
		var _nw45 _dafny.Array
		_ = _nw45
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw45 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Sequence = (func(_264_v0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_265_i21 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(_264_v0)), _dafny.SeqOf(_264_v0))
				}
			})(_76_v0)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw45 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw45).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw45).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_263_v126 = _nw45
		var _266_v127 _dafny.Sequence
		_ = _266_v127
		_266_v127 = _dafny.SeqOf(_76_v0)
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_263_v126), 0))
		_ = _index45
		(_263_v126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_266_v127, _266_v127), (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_263_v126), 0))
		_ = _index46
		(_263_v126).ArraySet1(_266_v127, (_index46).Int())
		_76_v0 = (_dafny.IntOfInt64(-732)).Cmp(_77_v1) < 0
		var _267_v128 _dafny.CodePoint
		_ = _267_v128
		_267_v128 = _dafny.CodePoint('v')
		var _268_v129 _dafny.Map
		_ = _268_v129
		_268_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _267_v128)
		var _269_v130 _dafny.Map
		_ = _269_v130
		_269_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, _79_v2)
		var _270_v131 bool
		_ = _270_v131
		var _271_v132 bool
		_ = _271_v132
		var _272_v133 _dafny.Int
		_ = _272_v133
		var _273_v134 bool
		_ = _273_v134
		var _out8 bool
		_ = _out8
		var _out9 bool
		_ = _out9
		var _out10 _dafny.Int
		_ = _out10
		var _out11 bool
		_ = _out11
		_out8, _out9, _out10, _out11 = Companion_Default___.M0(_268_v129, _269_v130, _76_v0, Companion_Default___.SafeModuloInt(_77_v1, _77_v1), _78_globalState)
		_270_v131 = _out8
		_271_v132 = _out9
		_272_v133 = _out10
		_273_v134 = _out11
		var _274_v135 _dafny.Sequence
		_ = _274_v135
		_274_v135 = _dafny.SeqOf(_272_v133, _77_v1)
		var _275_v136 _dafny.Array
		_ = _275_v136
		var _nwElement0_7 _dafny.Sequence = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))
		_ = _nwElement0_7
		var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
		_ = _nw46
		(_nw46).ArraySet1(_nwElement0_7, 0)
		(_nw46).ArraySet1(_274_v135, 1)
		(_nw46).ArraySet1(_274_v135, 2)
		(_nw46).ArraySet1(_274_v135, 3)
		(_nw46).ArraySet1(_274_v135, 4)
		(_nw46).ArraySet1(_274_v135, 5)
		(_nw46).ArraySet1(_dafny.Companion_Sequence_.Update(_274_v135, (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_274_v135).Cardinality()))).Uint32(), _272_v133), 6)
		_275_v136 = _nw46
		var _276_v137 T1
		_ = _276_v137
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__(_275_v136)
		_276_v137 = _nw47
		var _277_v138 D4
		_ = _277_v138
		_277_v138 = Companion_D4_.Create_DC10_(_276_v137)
		var _278_v139 _dafny.MultiSet
		_ = _278_v139
		_278_v139 = _dafny.MultiSetOf((_277_v138).Dtor_cf17(), _276_v137, _276_v137, _276_v137, _276_v137)
		_278_v139 = (((_278_v139).Update(_276_v137, Companion_Default___.Abs(_272_v133))).Union(_278_v139)).Update(_276_v137, Companion_Default___.Abs(Companion_Default___.SafeDivisionInt(_272_v133, _dafny.IntOfInt64(-895))))
	}
	var _source6 D0 = Companion_Default___.Fm2(_79_v2, !(_76_v0), _76_v0, false, _78_globalState)
	_ = _source6
	if _source6.Is_DC1() {
		_76_v0 = false
		var _279_v140 _dafny.Array
		_ = _279_v140
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw48
		_279_v140 = _nw48
		var _280_v141 _dafny.Sequence
		_ = _280_v141
		_280_v141 = _dafny.SeqOf(_279_v140)
		var _281_v142 _dafny.Array
		_ = _281_v142
		var _nwElement0_8 _dafny.Array = _279_v140
		_ = _nwElement0_8
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_8, 0)
		(_nw49).ArraySet1(_279_v140, 1)
		(_nw49).ArraySet1(_279_v140, 2)
		(_nw49).ArraySet1(_279_v140, 3)
		(_nw49).ArraySet1(_279_v140, 4)
		(_nw49).ArraySet1((_280_v141).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_77_v1), _dafny.IntOfUint32((_280_v141).Cardinality()))).Uint32()).(_dafny.Array), 5)
		(_nw49).ArraySet1(_279_v140, 6)
		(_nw49).ArraySet1(_279_v140, 7)
		(_nw49).ArraySet1(_279_v140, 8)
		(_nw49).ArraySet1(_279_v140, 9)
		(_nw49).ArraySet1(_279_v140, 10)
		(_nw49).ArraySet1(_279_v140, 11)
		(_nw49).ArraySet1(_279_v140, 12)
		(_nw49).ArraySet1(_279_v140, 13)
		_281_v142 = _nw49
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_91_v12), 0))
		_ = _index47
		(_91_v12).ArraySet1(_77_v1, (_index47).Int())
		var _282_v143 _dafny.CodePoint
		_ = _282_v143
		_282_v143 = _dafny.CodePoint('s')
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_91_v12), 0))
		_ = _index48
		var _rhs42 _dafny.Array = (func() _dafny.Array {
			if Companion_Default___.Fm0(_77_v1, _77_v1, _282_v143, _78_globalState) {
				return _281_v142
			}
			return _281_v142
		})()
		_ = _rhs42
		var _rhs43 _dafny.Int = _77_v1
		_ = _rhs43
		var _rhs44 _dafny.Int = Companion_Default___.SafeDivisionInt(_77_v1, _77_v1)
		_ = _rhs44
		var _lhs25 _dafny.Array = _91_v12
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_91_v12), 0))
		_ = _lhs26
		_281_v142 = _rhs42
		_77_v1 = _rhs43
		(_lhs25).ArraySet1(_rhs44, (_lhs26).Int())
		var _283_v144 _dafny.Map
		_ = _283_v144
		_283_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _282_v143)
		var _284_v145 _dafny.Sequence
		_ = _284_v145
		_284_v145 = _dafny.UnicodeSeqOfUtf8Bytes("cmrorbj")
		var _285_v146 _dafny.Map
		_ = _285_v146
		_285_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_284_v145).Cardinality()), _79_v2)
		var _286_v147 _dafny.Map
		_ = _286_v147
		_286_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int))).Update(_dafny.IntOfUint32((_284_v145).Cardinality()), Companion_Default___.Abs((_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int))), (func() _dafny.Set {
			if (_285_v146).Contains(_dafny.IntOfUint32((Companion_Default___.Fm12(_78_globalState)).Cardinality())) {
				return (_285_v146).Get(_dafny.IntOfUint32((Companion_Default___.Fm12(_78_globalState)).Cardinality())).(_dafny.Set)
			}
			return _79_v2
		})())
		var _287_v148 bool
		_ = _287_v148
		var _288_v149 bool
		_ = _288_v149
		var _289_v150 _dafny.Int
		_ = _289_v150
		var _290_v151 bool
		_ = _290_v151
		var _out12 bool
		_ = _out12
		var _out13 bool
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		var _out15 bool
		_ = _out15
		_out12, _out13, _out14, _out15 = Companion_Default___.M0(_283_v144, _286_v147, (func() bool {
			if _76_v0 {
				return _76_v0
			}
			return _76_v0
		})(), _77_v1, _78_globalState)
		_287_v148 = _out12
		_288_v149 = _out13
		_289_v150 = _out14
		_290_v151 = _out15
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_279_v140), 0))
		_ = _index49
		(_279_v140).ArraySet1(!(_290_v151) || (_76_v0), (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_279_v140), 0))
		_ = _index50
		(_279_v140).ArraySet1(((_76_v0) == (_288_v149)) == (_290_v151), (_index50).Int())
	} else {
		var _291___mcc_h11 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
		_ = _291___mcc_h11
		var _292_cf0 _dafny.Int = _291___mcc_h11
		_ = _292_cf0
		if !((_76_v0) == (_76_v0)) || (_76_v0) {
			var _293_v152 _dafny.Map
			_ = _293_v152
			_293_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_76_v0), !(_76_v0))
			var _294_v153 _dafny.Sequence
			_ = _294_v153
			_294_v153 = _dafny.UnicodeSeqOfUtf8Bytes("ysdj")
			_77_v1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (func() bool {
					if (_293_v152).Contains(_76_v0) {
						return (_293_v152).Get(_76_v0).(bool)
					}
					return true
				})() {
					return _dafny.Companion_Sequence_.Concatenate(_294_v153, _dafny.UnicodeSeqOfUtf8Bytes("m"))
				}
				return _294_v153
			})()).Cardinality())
			var _295_v154 _dafny.Map
			_ = _295_v154
			_295_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_79_v2).Union(_79_v2), _292_cf0)
			_295_v154 = ((_295_v154).Merge(_295_v154)).Merge(_295_v154)
			var _296_v155 _dafny.Sequence
			_ = _296_v155
			_296_v155 = _dafny.SeqOf(_77_v1, _77_v1)
			_296_v155 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_296_v155, _296_v155), _296_v155)
			var _297_v156 _dafny.Array
			_ = _297_v156
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw50
			_297_v156 = _nw50
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_297_v156), 0))
			_ = _index51
			(_297_v156).ArraySet1(_76_v0, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_297_v156), 0))
			_ = _index52
			(_297_v156).ArraySet1((_79_v2).IsSubsetOf(_dafny.SetOf(_76_v0)), (_index52).Int())
			var _298_v157 *C2
			_ = _298_v157
			var _nw51 *C2 = New_C2_()
			_ = _nw51
			_nw51.Ctor__(_292_cf0)
			_298_v157 = _nw51
			_298_v157 = _298_v157
		} else {
			var _299_v158 _dafny.Sequence
			_ = _299_v158
			_299_v158 = _dafny.SeqOf(_292_cf0, _292_cf0)
			var _300_v159 _dafny.MultiSet
			_ = _300_v159
			_300_v159 = _dafny.MultiSetOf(_76_v0)
			var _301_v160 _dafny.CodePoint
			_ = _301_v160
			_301_v160 = _dafny.CodePoint('f')
			var _302_v161 _dafny.Map
			_ = _302_v161
			_302_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_300_v159).Contains(_76_v0) {
					return (_300_v159).Multiplicity(_76_v0)
				}
				return _77_v1
			})(), _301_v160)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))
			_ = _index53
			(_91_v12).ArraySet1((_302_v161).Cardinality(), (_index53).Int())
			var _303_v162 _dafny.Sequence
			_ = _303_v162
			_303_v162 = _dafny.UnicodeSeqOfUtf8Bytes("icr")
			var _304_v163 _dafny.Set
			_ = _304_v163
			_304_v163 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_303_v162, (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_303_v162).Cardinality()))).Uint32(), _301_v160))
			var _305_v164 _dafny.Sequence
			_ = _305_v164
			_305_v164 = _dafny.SeqOf(_303_v162)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))
			_ = _index54
			var _rhs45 bool = (_304_v163).Equals((_304_v163).Intersection(_304_v163))
			_ = _rhs45
			var _rhs46 _dafny.Sequence = _299_v158
			_ = _rhs46
			var _rhs47 _dafny.Int = _77_v1
			_ = _rhs47
			var _rhs48 _dafny.Sequence = (_305_v164).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_305_v164).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs48
			var _lhs27 _dafny.Array = _91_v12
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))
			_ = _lhs28
			_76_v0 = _rhs45
			_299_v158 = _rhs46
			(_lhs27).ArraySet1(_rhs47, (_lhs28).Int())
			_303_v162 = _rhs48
			_76_v0 = _76_v0
			_77_v1 = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_292_cf0, _dafny.IntOfUint32((_303_v162).Cardinality())))).Minus((_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int))
			_292_cf0 = (_91_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))).Int()).(_dafny.Int)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_91_v12), 0))
			_ = _index55
			(_91_v12).ArraySet1((_292_cf0).Plus(_dafny.IntOfInt64(18)), (_index55).Int())
		}
		var _306_v165 _dafny.CodePoint
		_ = _306_v165
		_306_v165 = _dafny.CodePoint('l')
		var _307_v166 _dafny.Map
		_ = _307_v166
		_307_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _306_v165)
		var _308_v167 _dafny.Sequence
		_ = _308_v167
		_308_v167 = _dafny.UnicodeSeqOfUtf8Bytes("me")
		var _309_v168 _dafny.Map
		_ = _309_v168
		_309_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_292_cf0, _77_v1)).Update(_dafny.IntOfUint32((_308_v167).Cardinality()), Companion_Default___.Abs((_94_v13).Cardinality())), _dafny.SetOf(_76_v0))
		var _310_v169 _dafny.Set
		_ = _310_v169
		_310_v169 = _dafny.SetOf(_77_v1)
		var _311_v170 _dafny.Map
		_ = _311_v170
		_311_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(866), _76_v0)
		var _312_v172 bool
		_ = _312_v172
		var _313_v173 bool
		_ = _313_v173
		var _314_v174 _dafny.Int
		_ = _314_v174
		var _315_v175 bool
		_ = _315_v175
		var _out16 bool
		_ = _out16
		var _out17 bool
		_ = _out17
		var _out18 _dafny.Int
		_ = _out18
		var _out19 bool
		_ = _out19
		_out16, _out17, _out18, _out19 = Companion_Default___.M0(_307_v166, _309_v168, (_310_v169).IsDisjointFrom(func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter4 := _dafny.Iterate((_311_v170).Keys().Elements()); ; {
				_compr_3, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _316_v171 _dafny.Int
				_316_v171 = interface{}(_compr_3).(_dafny.Int)
				if (_311_v170).Contains(_316_v171) {
					_coll3.Add((_316_v171).Minus(_dafny.IntOfInt64(304)))
				}
			}
			return _coll3.ToSet()
		}()), Companion_Default___.SafeDivisionInt(_77_v1, _77_v1), _78_globalState)
		_312_v172 = _out16
		_313_v173 = _out17
		_314_v174 = _out18
		_315_v175 = _out19
		_95_v14 = _95_v14
		_77_v1 = (_77_v1).Plus(_292_cf0)
	}
	var _317_v176 _dafny.CodePoint
	_ = _317_v176
	_317_v176 = _dafny.CodePoint('q')
	var _318_v177 _dafny.Map
	_ = _318_v177
	_318_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_317_v176, _76_v0)
	var _319_v178 D2
	_ = _319_v178
	_319_v178 = Companion_D2_.Create_DC5_(!((func() bool {
		if (_318_v177).Contains(_317_v176) {
			return (_318_v177).Get(_317_v176).(bool)
		}
		return false
	})()), _77_v1, _77_v1)
	var _source7 D2 = _319_v178
	_ = _source7
	if _source7.Is_DC5() {
		var _320___mcc_h12 bool = _source7.Get_().(D2_DC5).Cf6
		_ = _320___mcc_h12
		var _321___mcc_h13 _dafny.Int = _source7.Get_().(D2_DC5).Cf7
		_ = _321___mcc_h13
		var _322___mcc_h14 _dafny.Int = _source7.Get_().(D2_DC5).Cf8
		_ = _322___mcc_h14
		var _323_cf8 _dafny.Int = _322___mcc_h14
		_ = _323_cf8
		var _324_cf7 _dafny.Int = _321___mcc_h13
		_ = _324_cf7
		var _325_cf6 bool = _320___mcc_h12
		_ = _325_cf6
		var _pat_let_tv3 = _324_cf7
		_ = _pat_let_tv3
		var _source8 D0 = func(_pat_let0_0 D0) D0 {
			return func(_326_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Int) D0 {
					return func(_327_dt__update_hcf0_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_327_dt__update_hcf0_h0)
					}(_pat_let1_0)
				}(_pat_let_tv3)
			}(_pat_let0_0)
		}(_80_v3)
		_ = _source8
		if _source8.Is_DC1() {
			var _328_v179 *C1
			_ = _328_v179
			var _nw52 *C1 = New_C1_()
			_ = _nw52
			_nw52.Ctor__()
			_328_v179 = _nw52
			_317_v176 = Companion_Default___.Fm3(_78_globalState)
			var _329_v180 D0
			_ = _329_v180
			var _330_v181 _dafny.Int
			_ = _330_v181
			var _out20 D0
			_ = _out20
			var _out21 _dafny.Int
			_ = _out21
			_out20, _out21 = (_328_v179).M1(_78_globalState)
			_329_v180 = _out20
			_330_v181 = _out21
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_91_v12), 0))
			_ = _index56
			(_91_v12).ArraySet1(_324_cf7, (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_91_v12), 0))
			_ = _index57
			(_91_v12).ArraySet1((_dafny.Zero).Minus(_77_v1), (_index57).Int())
		} else {
			var _331___mcc_h17 _dafny.Int = _source8.Get_().(D0_DC0).Cf0
			_ = _331___mcc_h17
			var _332_cf0 _dafny.Int = _331___mcc_h17
			_ = _332_cf0
			var _333_v182 _dafny.Array
			_ = _333_v182
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw53
			_333_v182 = _nw53
			var _334_v183 *C0
			_ = _334_v183
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__(_333_v182)
			_334_v183 = _nw54
			var _335_v184 _dafny.Array
			_ = _335_v184
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_14
			var _nw55 _dafny.Array
			_ = _nw55
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw55 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_336_v1 _dafny.Int, _337_cf8 _dafny.Int) func(_dafny.Int) bool {
					return func(_338_i22 _dafny.Int) bool {
						return (_dafny.SetOf(_337_cf8)).Contains(_336_v1)
					}
				})(_77_v1, _323_cf8)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw55 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw55).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw55).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_335_v184 = _nw55
			var _339_v185 _dafny.Sequence
			_ = _339_v185
			_339_v185 = _dafny.SeqOf(_335_v184, _335_v184)
			var _340_v186 _dafny.Sequence
			_ = _340_v186
			_340_v186 = _dafny.UnicodeSeqOfUtf8Bytes("vnhtq")
			_335_v184 = (_339_v185).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_340_v186).Cardinality()), _dafny.IntOfUint32((_339_v185).Cardinality()))).Uint32()).(_dafny.Array)
			_77_v1 = (func() _dafny.Int {
				if _76_v0 {
					return _332_cf0
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qixb")).Cardinality())
			})()
			var _341_v187 _dafny.Map
			_ = _341_v187
			_341_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12(_78_globalState), _317_v176)
			var _rhs49 _dafny.Map = _341_v187
			_ = _rhs49
			var _rhs50 _dafny.Int = (_dafny.Zero).Minus(_77_v1)
			_ = _rhs50
			_341_v187 = _rhs49
			_77_v1 = _rhs50
		}
		var _342_v188 *C1
		_ = _342_v188
		var _nw56 *C1 = New_C1_()
		_ = _nw56
		_nw56.Ctor__()
		_342_v188 = _nw56
		var _343_v189 D3
		_ = _343_v189
		_343_v189 = Companion_D3_.Create_DC8_(_325_cf6, _153_v59, _76_v0, _79_v2, _342_v188)
		var _344_v190 _dafny.Map
		_ = _344_v190
		_344_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_317_v176, _343_v189)
		var _345_v191 D5
		_ = _345_v191
		_345_v191 = Companion_D5_.Create_DC13_(_344_v190)
		var _346_v192 D3
		_ = _346_v192
		_346_v192 = Companion_D3_.Create_DC9_(((_345_v191).Dtor_cf19()).Cardinality())
		var _347_v193 _dafny.Sequence
		_ = _347_v193
		_347_v193 = _dafny.UnicodeSeqOfUtf8Bytes("ka")
		var _348_v194 _dafny.Sequence
		_ = _348_v194
		_348_v194 = _dafny.SeqOf(_dafny.IntOfUint32((_347_v193).Cardinality()))
		var _349_v195 _dafny.MultiSet
		_ = _349_v195
		_349_v195 = _dafny.MultiSetOf(_348_v194, _348_v194, _348_v194, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(829))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_350_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_351_i23 _dafny.Int) _dafny.Int {
				return _350_v1
			}
		})(_77_v1))))
		var _352_v196 _dafny.Map
		_ = _352_v196
		_352_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _349_v195)
		var _rhs51 _dafny.Int = (_346_v192).Dtor_cf16()
		_ = _rhs51
		var _rhs52 _dafny.Int = _77_v1
		_ = _rhs52
		var _rhs53 bool = (((func() _dafny.MultiSet {
			if (_352_v196).Contains(_324_cf7) {
				return (_352_v196).Get(_324_cf7).(_dafny.MultiSet)
			}
			return _349_v195
		})()).Update(_348_v194, Companion_Default___.Abs(_77_v1))).IsProperSubsetOf(_349_v195)
		_ = _rhs53
		_324_cf7 = _rhs51
		_77_v1 = _rhs52
		_325_cf6 = _rhs53
		var _353_v197 _dafny.Array
		_ = _353_v197
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
		_ = _nw57
		_353_v197 = _nw57
		_353_v197 = _353_v197
		var _354_v198 _dafny.Array
		_ = _354_v198
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_15
		var _nw58 _dafny.Array
		_ = _nw58
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw58 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.Sequence = (func(_355_v194 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_356_i24 _dafny.Int) _dafny.Sequence {
					return _355_v194
				}
			})(_348_v194)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw58 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw58).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw58).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_354_v198 = _nw58
		var _357_v199 T1
		_ = _357_v199
		var _nw59 *C0 = New_C0_()
		_ = _nw59
		_nw59.Ctor__(_354_v198)
		_357_v199 = _nw59
		var _358_v200 _dafny.Map
		_ = _358_v200
		_358_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_cf7, _357_v199)
		var _359_v201 _dafny.Set
		_ = _359_v201
		_359_v201 = _dafny.SetOf(_358_v200)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_91_v12), 0))
		_ = _index58
		(_91_v12).ArraySet1(_323_cf8, (_index58).Int())
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_91_v12), 0))
		_ = _index59
		var _rhs54 _dafny.Set = (_359_v201).Union((func() _dafny.Set {
			if _76_v0 {
				return _359_v201
			}
			return _359_v201
		})())
		_ = _rhs54
		var _rhs55 bool = _76_v0
		_ = _rhs55
		var _rhs56 T1 = _357_v199
		_ = _rhs56
		var _rhs57 _dafny.Int = _324_cf7
		_ = _rhs57
		var _lhs29 _dafny.Array = _91_v12
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_91_v12), 0))
		_ = _lhs30
		_359_v201 = _rhs54
		_325_cf6 = _rhs55
		_357_v199 = _rhs56
		(_lhs29).ArraySet1(_rhs57, (_lhs30).Int())
	} else if _source7.Is_DC4() {
		var _360___mcc_h15 _dafny.Sequence = _source7.Get_().(D2_DC4).Cf5
		_ = _360___mcc_h15
		var _361_cf5 _dafny.Sequence = _360___mcc_h15
		_ = _361_cf5
		if _76_v0 {
			_91_v12 = _91_v12
			var _362_v202 _dafny.Array
			_ = _362_v202
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_16
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = func(_363_i25 _dafny.Int) bool {
					return true
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw60 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw60).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw60).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_362_v202 = _nw60
			var _364_v203 D1
			_ = _364_v203
			_364_v203 = Companion_D1_.Create_DC3_(_76_v0, !(_76_v0), _76_v0)
			var _365_v204 _dafny.Map
			_ = _365_v204
			_365_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_364_v203).Dtor_cf4(), _dafny.IntOfInt64(-932))
			var _366_v205 _dafny.Array
			_ = _366_v205
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
			_ = _nw61
			_366_v205 = _nw61
			var _367_v206 *C0
			_ = _367_v206
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__(_366_v205)
			_367_v206 = _nw62
			var _368_v207 _dafny.MultiSet
			_ = _368_v207
			_368_v207 = _dafny.MultiSetOf(_367_v206)
			var _369_v208 _dafny.Map
			_ = _369_v208
			_369_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_365_v204).Contains(_76_v0) {
					return (_365_v204).Get(_76_v0).(_dafny.Int)
				}
				return (_368_v207).Cardinality()
			})(), !(_76_v0))
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))
			_ = _index60
			(_362_v202).ArraySet1((func() bool {
				if (_369_v208).Contains(_dafny.IntOfInt64(204)) {
					return (_369_v208).Get(_dafny.IntOfInt64(204)).(bool)
				}
				return _76_v0
			})(), (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))
			_ = _index61
			(_362_v202).ArraySet1(_76_v0, (_index61).Int())
			_153_v59 = Companion_D2_.Create_DC6_(_152_v58)
			var _370_v209 _dafny.Map
			_ = _370_v209
			_370_v209 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_94_v13).Intersection(_94_v13), _77_v1)
			var _371_v210 _dafny.Sequence
			_ = _371_v210
			_371_v210 = _dafny.SeqOf(_361_cf5)
			var _372_v211 _dafny.Array
			_ = _372_v211
			var _nwElement0_9 _dafny.Sequence = _371_v210
			_ = _nwElement0_9
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.One)
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_9, 0)
			_372_v211 = _nw63
			var _373_v212 D6
			_ = _373_v212
			_373_v212 = Companion_D6_.Create_DC17_(_371_v210)
			var _374_v213 _dafny.Sequence
			_ = _374_v213
			_374_v213 = _dafny.SeqOf((_373_v212).Dtor_cf23())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_372_v211), 0))
			_ = _index62
			(_372_v211).ArraySet1((_374_v213).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_374_v213).Cardinality()))).Uint32()).(_dafny.Sequence), (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))
			_ = _index63
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_372_v211), 0))
			_ = _index64
			var _rhs58 bool = (_362_v202).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))).Int()).(bool)
			_ = _rhs58
			var _rhs59 bool = (_77_v1).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm15((_365_v204).Cardinality(), _78_globalState), (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((Companion_Default___.Fm15((_365_v204).Cardinality(), _78_globalState)).Cardinality()))).Uint32(), _76_v0)).Cardinality()), _77_v1)) <= 0
			_ = _rhs59
			var _rhs60 _dafny.Map = _370_v209
			_ = _rhs60
			var _rhs61 _dafny.Sequence = _371_v210
			_ = _rhs61
			var _lhs31 _dafny.Array = _362_v202
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _372_v211
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_372_v211), 0))
			_ = _lhs34
			_76_v0 = _rhs58
			(_lhs31).ArraySet1(_rhs59, (_lhs32).Int())
			_370_v209 = _rhs60
			(_lhs33).ArraySet1(_rhs61, (_lhs34).Int())
			var _375_v214 _dafny.Map
			_ = _375_v214
			_375_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_362_v202).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))).Int()).(bool), _317_v176)
			var _376_v215 bool
			_ = _376_v215
			var _377_v216 bool
			_ = _377_v216
			var _378_v217 _dafny.Int
			_ = _378_v217
			var _379_v218 bool
			_ = _379_v218
			var _out22 bool
			_ = _out22
			var _out23 bool
			_ = _out23
			var _out24 _dafny.Int
			_ = _out24
			var _out25 bool
			_ = _out25
			_out22, _out23, _out24, _out25 = Companion_Default___.M0(_375_v214, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, Companion_Default___.Fm5(_77_v1, false, (_362_v202).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))).Int()).(bool), _77_v1, _78_globalState)), (_362_v202).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_362_v202), 0))).Int()).(bool), _77_v1, _78_globalState)
			_376_v215 = _out22
			_377_v216 = _out23
			_378_v217 = _out24
			_379_v218 = _out25
		} else {
			_80_v3 = _80_v3
			var _380_v219 _dafny.Array
			_ = _380_v219
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw64
			_380_v219 = _nw64
			var _381_v220 T0
			_ = _381_v220
			var _nw65 *C2 = New_C2_()
			_ = _nw65
			_nw65.Ctor__(_dafny.IntOfInt64(53))
			_381_v220 = _nw65
			var _382_v221 D6
			_ = _382_v221
			_382_v221 = Companion_D6_.Create_DC18_(_77_v1, _381_v220, Companion_D5_.Create_DC15_(_76_v0), _77_v1)
			var _383_v222 D6
			_ = _383_v222
			_383_v222 = Companion_D6_.Create_DC19_(_382_v221)
			var _384_v223 _dafny.Set
			_ = _384_v223
			_384_v223 = _dafny.SetOf(_383_v222)
			var _385_v224 _dafny.Map
			_ = _385_v224
			_385_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v223, _76_v0)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))
			_ = _index65
			(_380_v219).ArraySet1((_77_v1).Cmp((_385_v224).Cardinality()) < 0, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))
			_ = _index66
			(_380_v219).ArraySet1(_76_v0, (_index66).Int())
			_77_v1 = _dafny.IntOfInt64(742)
			var _386_v225 _dafny.Sequence
			_ = _386_v225
			_386_v225 = _dafny.SeqOf(_361_cf5)
			_386_v225 = _386_v225
			var _387_v226 _dafny.Map
			_ = _387_v226
			_387_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, true)
			var _388_v227 _dafny.Map
			_ = _388_v227
			_388_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _387_v226)
			var _389_v228 _dafny.Sequence
			_ = _389_v228
			_389_v228 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), _387_v226)
			var _390_v229 *C1
			_ = _390_v229
			var _nw66 *C1 = New_C1_()
			_ = _nw66
			_nw66.Ctor__()
			_390_v229 = _nw66
			var _391_v230 _dafny.Sequence
			_ = _391_v230
			_391_v230 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), _390_v229))
			var _392_v231 _dafny.Array
			_ = _392_v231
			var _nwElement0_10 _dafny.Map = (func() _dafny.Map {
				if (_388_v227).Contains(_77_v1) {
					return (_388_v227).Get(_77_v1).(_dafny.Map)
				}
				return _387_v226
			})()
			_ = _nwElement0_10
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_10, 0)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), 1)
			(_nw67).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool))).Merge(Companion_Default___.Fm19((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), _77_v1, _78_globalState)), 2)
			(_nw67).ArraySet1(_387_v226, 3)
			(_nw67).ArraySet1(((_387_v226).Update((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool))), 4)
			(_nw67).ArraySet1((_389_v228).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_389_v228).Cardinality()))).Uint32()).(_dafny.Map), 5)
			(_nw67).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), _76_v0)).Merge(_387_v226), 6)
			(_nw67).ArraySet1(_387_v226, 7)
			(_nw67).ArraySet1(Companion_Default___.Fm19(!(false), _dafny.IntOfInt64(747), _78_globalState), 8)
			(_nw67).ArraySet1(_387_v226, 9)
			(_nw67).ArraySet1((_387_v226).Merge(_387_v226), 10)
			(_nw67).ArraySet1(_387_v226, 11)
			(_nw67).ArraySet1(_387_v226, 12)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), 13)
			(_nw67).ArraySet1(_387_v226, 14)
			(_nw67).ArraySet1((_389_v228).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_391_v230).Cardinality()), _dafny.IntOfUint32((_389_v228).Cardinality()))).Uint32()).(_dafny.Map), 15)
			(_nw67).ArraySet1(_387_v226, 16)
			(_nw67).ArraySet1((_387_v226).Merge(_387_v226), 17)
			(_nw67).ArraySet1((_389_v228).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_389_v228).Cardinality()))).Uint32()).(_dafny.Map), 18)
			(_nw67).ArraySet1(_387_v226, 19)
			(_nw67).ArraySet1(_387_v226, 20)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), false), 21)
			(_nw67).ArraySet1(Companion_Default___.Fm19(_76_v0, _dafny.IntOfUint32((_361_cf5).Cardinality()), _78_globalState), 22)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _76_v0), 23)
			(_nw67).ArraySet1((_387_v226).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() bool {
				if (_318_v177).Contains(_317_v176) {
					return (_318_v177).Get(_317_v176).(bool)
				}
				return (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)
			})())), 24)
			(_nw67).ArraySet1(_387_v226, 25)
			(_nw67).ArraySet1((_387_v226).Merge(_387_v226), 26)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool)), 27)
			(_nw67).ArraySet1((_387_v226).Merge((_387_v226).Update((_380_v219).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_380_v219), 0))).Int()).(bool), _76_v0)), 28)
			_392_v231 = _nw67
			_392_v231 = _392_v231
		}
		var _393_v232 _dafny.Map
		_ = _393_v232
		_393_v232 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _76_v0)
		var _394_v233 T0
		_ = _394_v233
		var _nw68 *C2 = New_C2_()
		_ = _nw68
		_nw68.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_361_cf5).Cardinality()))))
		_394_v233 = _nw68
		var _395_v234 _dafny.Array
		_ = _395_v234
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_17
		var _nw69 _dafny.Array
		_ = _nw69
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw69 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) bool = (func(_396_v0 bool) func(_dafny.Int) bool {
				return func(_397_i26 _dafny.Int) bool {
					return !(_396_v0) || (_396_v0)
				}
			})(_76_v0)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw69 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw69).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw69).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_395_v234 = _nw69
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))
		_ = _index67
		(_395_v234).ArraySet1(true, (_index67).Int())
		var _398_v235 _dafny.Array
		_ = _398_v235
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
		_ = _nw70
		_398_v235 = _nw70
		var _399_v236 *C0
		_ = _399_v236
		var _nw71 *C0 = New_C0_()
		_ = _nw71
		_nw71.Ctor__(_398_v235)
		_399_v236 = _nw71
		var _400_v237 _dafny.Set
		_ = _400_v237
		_400_v237 = _dafny.SetOf(_399_v236)
		var _401_v238 _dafny.Sequence
		_ = _401_v238
		_401_v238 = _dafny.SeqOf(_94_v13)
		var _402_v239 _dafny.Sequence
		_ = _402_v239
		_402_v239 = _dafny.SeqOf(_77_v1)
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))
		_ = _index68
		var _rhs62 _dafny.Int = (((_400_v237).Difference(_400_v237)).Cardinality()).Plus(_dafny.IntOfUint32((_361_cf5).Cardinality()))
		_ = _rhs62
		var _rhs63 _dafny.Map = Companion_Default___.Fm19(!(((_401_v238).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_401_v238).Cardinality()))).Uint32()).(_dafny.MultiSet)).Equals(_94_v13)), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_402_v239).Cardinality()), (Companion_Default___.Fm20(_77_v1, _76_v0, _76_v0, _77_v1, _78_globalState)).Cardinality()), _78_globalState)
		_ = _rhs63
		var _rhs64 T0 = _394_v233
		_ = _rhs64
		var _rhs65 bool = _76_v0
		_ = _rhs65
		var _rhs66 bool = _76_v0
		_ = _rhs66
		var _lhs35 _dafny.Array = _395_v234
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))
		_ = _lhs36
		_77_v1 = _rhs62
		_393_v232 = _rhs63
		_394_v233 = _rhs64
		(_lhs35).ArraySet1(_rhs65, (_lhs36).Int())
		_76_v0 = _rhs66
		var _403_v240 _dafny.Array
		_ = _403_v240
		var _nw72 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
		_ = _nw72
		_403_v240 = _nw72
		var _404_v241 D7
		_ = _404_v241
		_404_v241 = Companion_D7_.Create_DC20_(_399_v236)
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_403_v240), 0))
		_ = _index69
		(_403_v240).ArraySet1((_404_v241).Dtor_cf29(), (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_403_v240), 0))
		_ = _index70
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(_398_v235)
		(_403_v240).ArraySet1(_nw73, (_index70).Int())
		var _405_v242 _dafny.Array
		_ = _405_v242
		var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw74
		_405_v242 = _nw74
		var _406_v243 _dafny.Map
		_ = _406_v243
		_406_v243 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_393_v232).Contains((_395_v234).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))).Int()).(bool)) {
				return (_393_v232).Get((_395_v234).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))).Int()).(bool)).(bool)
			}
			return false
		})(), _395_v234)
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_405_v242), 0))
		_ = _index71
		(_405_v242).ArraySet1((_406_v243).Merge(_406_v243), (_index71).Int())
		var _407_v244 _dafny.Map
		_ = _407_v244
		_407_v244 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_395_v234).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_395_v234), 0))).Int()).(bool), _361_cf5)
		var _408_v245 _dafny.Array
		_ = _408_v245
		var _nw75 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(10))
		_ = _nw75
		_408_v245 = _nw75
		var _409_v246 D8
		_ = _409_v246
		_409_v246 = Companion_D8_.Create_DC24_(_408_v245)
		var _410_v247 _dafny.Map
		_ = _410_v247
		_410_v247 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _408_v245)
		var _pat_let_tv4 = _410_v247
		_ = _pat_let_tv4
		var _pat_let_tv5 = _76_v0
		_ = _pat_let_tv5
		var _pat_let_tv6 = _410_v247
		_ = _pat_let_tv6
		var _pat_let_tv7 = _76_v0
		_ = _pat_let_tv7
		var _pat_let_tv8 = _408_v245
		_ = _pat_let_tv8
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_405_v242), 0))
		_ = _index72
		var _rhs67 _dafny.Map = _406_v243
		_ = _rhs67
		var _rhs68 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _361_cf5)
		_ = _rhs68
		var _rhs69 _dafny.Array = (func(_pat_let2_0 D8) D8 {
			return func(_411_dt__update__tmp_h1 D8) D8 {
				return func(_pat_let3_0 _dafny.Array) D8 {
					return func(_412_dt__update_hcf34_h0 _dafny.Array) D8 {
						return Companion_D8_.Create_DC24_(_412_dt__update_hcf34_h0)
					}(_pat_let3_0)
				}((func() _dafny.Array {
					if (_pat_let_tv4).Contains(_pat_let_tv5) {
						return (_pat_let_tv6).Get(_pat_let_tv7).(_dafny.Array)
					}
					return _pat_let_tv8
				})())
			}(_pat_let2_0)
		}(_409_v246)).Dtor_cf34()
		_ = _rhs69
		var _rhs70 _dafny.Int = (_77_v1).Plus(_77_v1)
		_ = _rhs70
		var _rhs71 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_402_v239, _402_v239)
		_ = _rhs71
		var _lhs37 _dafny.Array = _405_v242
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_405_v242), 0))
		_ = _lhs38
		(_lhs37).ArraySet1(_rhs67, (_lhs38).Int())
		_407_v244 = _rhs68
		_408_v245 = _rhs69
		_77_v1 = _rhs70
		_402_v239 = _rhs71
	} else {
		var _413___mcc_h16 D2 = _source7.Get_().(D2_DC6).Cf9
		_ = _413___mcc_h16
		var _414_cf9 D2 = _413___mcc_h16
		_ = _414_cf9
		var _415_v248 *C2
		_ = _415_v248
		var _nw76 *C2 = New_C2_()
		_ = _nw76
		_nw76.Ctor__(_77_v1)
		_415_v248 = _nw76
		var _416_v249 bool
		_ = _416_v249
		var _417_v250 bool
		_ = _417_v250
		var _418_v251 _dafny.Int
		_ = _418_v251
		var _419_v252 bool
		_ = _419_v252
		var _out26 bool
		_ = _out26
		var _out27 bool
		_ = _out27
		var _out28 _dafny.Int
		_ = _out28
		var _out29 bool
		_ = _out29
		_out26, _out27, _out28, _out29 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _317_v176), Companion_Default___.Fm21(_77_v1, Companion_Default___.Fm1(_77_v1, (_dafny.Zero).Minus(_77_v1), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _415_v248)).Cardinality(), _76_v0, _78_globalState), _78_globalState), _76_v0, _77_v1, _78_globalState)
		_416_v249 = _out26
		_417_v250 = _out27
		_418_v251 = _out28
		_419_v252 = _out29
		var _420_v253 *C1
		_ = _420_v253
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__()
		_420_v253 = _nw77
		var _421_v254 D3
		_ = _421_v254
		_421_v254 = Companion_D3_.Create_DC8_(_419_v252, _153_v59, true, _dafny.SetOf(_76_v0), _420_v253)
		var _422_v255 _dafny.Sequence
		_ = _422_v255
		_422_v255 = _dafny.SeqOf(_421_v254)
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_91_v12), 0))
		_ = _index73
		(_91_v12).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_421_v254, _421_v254, _421_v254, _421_v254), _422_v255)).Cardinality())), (_index73).Int())
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_91_v12), 0))
		_ = _index74
		(_91_v12).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if _76_v0 {
				return _dafny.IntOfInt64(624)
			}
			return (_415_v248).Fm18(_319_v178, _78_globalState)
		})())), (_index74).Int())
		var _423_v256 *C2
		_ = _423_v256
		var _nw78 *C2 = New_C2_()
		_ = _nw78
		_nw78.Ctor__(((_dafny.MultiSetOf(_417_v250)).Cardinality()).Plus(_77_v1))
		_423_v256 = _nw78
		var _424_v257 _dafny.Map
		_ = _424_v257
		_424_v257 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _417_v250)
		var _425_v258 _dafny.Sequence
		_ = _425_v258
		_425_v258 = _dafny.SeqOf(_76_v0, (_76_v0) && ((func() bool {
			if (_424_v257).Contains(_419_v252) {
				return (_424_v257).Get(_419_v252).(bool)
			}
			return _76_v0
		})()), _417_v250)
		_425_v258 = _425_v258
	}
	if _76_v0 {
		var _426_v259 _dafny.Map
		_ = _426_v259
		_426_v259 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _317_v176)
		var _427_v260 _dafny.Map
		_ = _427_v260
		_427_v260 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, _dafny.SetOf(_76_v0))
		var _428_v261 _dafny.Sequence
		_ = _428_v261
		_428_v261 = _dafny.UnicodeSeqOfUtf8Bytes("klwl")
		var _429_v262 bool
		_ = _429_v262
		var _430_v263 bool
		_ = _430_v263
		var _431_v264 _dafny.Int
		_ = _431_v264
		var _432_v265 bool
		_ = _432_v265
		var _out30 bool
		_ = _out30
		var _out31 bool
		_ = _out31
		var _out32 _dafny.Int
		_ = _out32
		var _out33 bool
		_ = _out33
		_out30, _out31, _out32, _out33 = Companion_Default___.M0((_426_v259).Merge(_426_v259), _427_v260, _76_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_428_v261, _428_v261)).Cardinality()), _78_globalState)
		_429_v262 = _out30
		_430_v263 = _out31
		_431_v264 = _out32
		_432_v265 = _out33
		var _433_v266 _dafny.Map
		_ = _433_v266
		_433_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _dafny.IntOfInt64(717))
		var _rhs72 bool = _430_v263
		_ = _rhs72
		var _rhs73 _dafny.Int = (_433_v266).Cardinality()
		_ = _rhs73
		var _rhs74 _dafny.Int = _77_v1
		_ = _rhs74
		var _rhs75 _dafny.Int = _77_v1
		_ = _rhs75
		_432_v265 = _rhs72
		_431_v264 = _rhs73
		_431_v264 = _rhs74
		_77_v1 = _rhs75
		var _434_v267 T0
		_ = _434_v267
		var _nw79 *C2 = New_C2_()
		_ = _nw79
		_nw79.Ctor__(_77_v1)
		_434_v267 = _nw79
		var _435_v268 _dafny.Sequence
		_ = _435_v268
		_435_v268 = _dafny.SeqOf(_434_v267)
		var _436_v269 *C2
		_ = _436_v269
		var _nw80 *C2 = New_C2_()
		_ = _nw80
		_nw80.Ctor__((_dafny.IntOfUint32((_435_v268).Cardinality())).Times(_77_v1))
		_436_v269 = _nw80
		_436_v269 = _436_v269
		_428_v261 = _428_v261
		if (_319_v178).Dtor_cf6() {
			var _437_v270 *C2
			_ = _437_v270
			var _nw81 *C2 = New_C2_()
			_ = _nw81
			_nw81.Ctor__(_77_v1)
			_437_v270 = _nw81
			var _438_v271 _dafny.Sequence
			_ = _438_v271
			_438_v271 = _dafny.SeqOf(_429_v262, _432_v265, _429_v262)
			var _439_v272 _dafny.MultiSet
			_ = _439_v272
			_439_v272 = _dafny.MultiSetOf(_429_v262)
			var _440_v273 _dafny.Map
			_ = _440_v273
			_440_v273 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v272, _431_v264)
			_438_v271 = _dafny.SeqOf(_430_v263, !((_77_v1).Cmp((_dafny.Zero).Minus((_440_v273).Cardinality())) == 0), _430_v263)
			_429_v262 = false
			var _441_v274 *C2
			_ = _441_v274
			var _nw82 *C2 = New_C2_()
			_ = _nw82
			_nw82.Ctor__(_431_v264)
			_441_v274 = _nw82
			var _442_v276 _dafny.Map
			_ = _442_v276
			_442_v276 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v13, true)
			var _443_v277 bool
			_ = _443_v277
			var _444_v278 bool
			_ = _444_v278
			var _445_v279 _dafny.Int
			_ = _445_v279
			var _446_v280 bool
			_ = _446_v280
			var _out34 bool
			_ = _out34
			var _out35 bool
			_ = _out35
			var _out36 _dafny.Int
			_ = _out36
			var _out37 bool
			_ = _out37
			_out34, _out35, _out36, _out37 = Companion_Default___.M0(_426_v259, func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter5 := _dafny.Iterate((_442_v276).Keys().Elements()); ; {
					_compr_4, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _447_v275 _dafny.MultiSet
					_447_v275 = interface{}(_compr_4).(_dafny.MultiSet)
					if (_442_v276).Contains(_447_v275) {
						_coll4.Add(_447_v275, _dafny.SetOf(_432_v265))
					}
				}
				return _coll4.ToMap()
			}(), _432_v265, (_441_v274).Fm18(_319_v178, _78_globalState), _78_globalState)
			_443_v277 = _out34
			_444_v278 = _out35
			_445_v279 = _out36
			_446_v280 = _out37
		} else {
			var _448_v281 _dafny.Sequence
			_ = _448_v281
			_448_v281 = _dafny.SeqOf(Companion_Default___.SafeModuloInt((_436_v269).F3(), (_436_v269).F3()))
			_448_v281 = _448_v281
			var _449_v282 _dafny.Array
			_ = _449_v282
			var _nw83 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw83
			_449_v282 = _nw83
			_449_v282 = _449_v282
			var _450_v283 bool
			_ = _450_v283
			var _451_v284 bool
			_ = _451_v284
			var _452_v285 _dafny.Int
			_ = _452_v285
			var _453_v286 bool
			_ = _453_v286
			var _out38 bool
			_ = _out38
			var _out39 bool
			_ = _out39
			var _out40 _dafny.Int
			_ = _out40
			var _out41 bool
			_ = _out41
			_out38, _out39, _out40, _out41 = Companion_Default___.M0((Companion_Default___.Fm4(_431_v264, _430_v263, _78_globalState)).Merge(_426_v259), _427_v260, _430_v263, (_436_v269).F3(), _78_globalState)
			_450_v283 = _out38
			_451_v284 = _out39
			_452_v285 = _out40
			_453_v286 = _out41
			var _454_v287 _dafny.Array
			_ = _454_v287
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw84
			_454_v287 = _nw84
			var _455_v288 _dafny.Map
			_ = _455_v288
			_455_v288 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v269).F3(), _454_v287)
			var _456_v289 _dafny.Array
			_ = _456_v289
			var _nwElement0_11 _dafny.Array = _454_v287
			_ = _nwElement0_11
			var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(17))
			_ = _nw85
			(_nw85).ArraySet1(_nwElement0_11, 0)
			(_nw85).ArraySet1(_454_v287, 1)
			(_nw85).ArraySet1(_454_v287, 2)
			(_nw85).ArraySet1(_454_v287, 3)
			(_nw85).ArraySet1(_454_v287, 4)
			(_nw85).ArraySet1(_454_v287, 5)
			(_nw85).ArraySet1(_454_v287, 6)
			(_nw85).ArraySet1(_454_v287, 7)
			(_nw85).ArraySet1(_454_v287, 8)
			(_nw85).ArraySet1(_454_v287, 9)
			(_nw85).ArraySet1(_454_v287, 10)
			(_nw85).ArraySet1(_454_v287, 11)
			(_nw85).ArraySet1(_454_v287, 12)
			(_nw85).ArraySet1((func() _dafny.Array {
				if (_455_v288).Contains((_436_v269).F3()) {
					return (_455_v288).Get((_436_v269).F3()).(_dafny.Array)
				}
				return _454_v287
			})(), 13)
			(_nw85).ArraySet1(_454_v287, 14)
			(_nw85).ArraySet1(_454_v287, 15)
			(_nw85).ArraySet1(_454_v287, 16)
			_456_v289 = _nw85
			var _457_v290 _dafny.Map
			_ = _457_v290
			_457_v290 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v289, _431_v264)
			var _458_v291 D9
			_ = _458_v291
			_458_v291 = Companion_D9_.Create_DC26_(_448_v281)
			_457_v290 = (_457_v290).Update(_456_v289, (_dafny.Zero).Minus((_dafny.Zero).Minus((_452_v285).Times((_dafny.MultiSetFromSeq((_458_v291).Dtor_cf37())).Cardinality()))))
			var _459_v292 _dafny.Map
			_ = _459_v292
			_459_v292 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_430_v263, _428_v261)
			var _460_v293 _dafny.Map
			_ = _460_v293
			_460_v293 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v267, (_dafny.IntOfInt64(883)).Cmp(_dafny.IntOfInt64(812)) <= 0)
			var _rhs76 bool = (_436_v269).Fm8(_448_v281, false, (_436_v269).F3(), _78_globalState)
			_ = _rhs76
			var _rhs77 _dafny.Map = (_459_v292).Merge((_459_v292).Merge(_459_v292))
			_ = _rhs77
			var _rhs78 _dafny.Map = (_460_v293).Merge((_460_v293).Merge(_460_v293))
			_ = _rhs78
			var _rhs79 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm13(_78_globalState)).Dtor_cf7())
			_ = _rhs79
			_451_v284 = _rhs76
			_459_v292 = _rhs77
			_460_v293 = _rhs78
			_452_v285 = _rhs79
		}
	} else {
		var _461_v294 _dafny.Sequence
		_ = _461_v294
		_461_v294 = _dafny.UnicodeSeqOfUtf8Bytes("sqshqoby")
		var _462_v295 _dafny.Array
		_ = _462_v295
		var _nwElement0_12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_461_v294, _461_v294)
		_ = _nwElement0_12
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(23))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_12, 0)
		(_nw86).ArraySet1(_461_v294, 1)
		(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_461_v294, _dafny.UnicodeSeqOfUtf8Bytes("kjnhobbkk")), 2)
		(_nw86).ArraySet1(_dafny.Companion_Sequence_.Update(_461_v294, (Companion_Default___.SafeIndex((_79_v2).Cardinality(), _dafny.IntOfUint32((_461_v294).Cardinality()))).Uint32(), _317_v176), 3)
		(_nw86).ArraySet1(_461_v294, 4)
		(_nw86).ArraySet1(_461_v294, 5)
		(_nw86).ArraySet1(_461_v294, 6)
		(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_463_v294 _dafny.Sequence, _464_v1 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
			return func(_465_i27 _dafny.Int) _dafny.CodePoint {
				return (_463_v294).Select((Companion_Default___.SafeIndex(_464_v1, _dafny.IntOfUint32((_463_v294).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_461_v294, _77_v1))), 7)
		(_nw86).ArraySet1(_461_v294, 8)
		(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_461_v294, _461_v294), 9)
		(_nw86).ArraySet1(_461_v294, 10)
		(_nw86).ArraySet1(_461_v294, 11)
		(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12(_78_globalState), _461_v294), 12)
		(_nw86).ArraySet1(_461_v294, 13)
		(_nw86).ArraySet1(_461_v294, 14)
		(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_466_v176 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_467_i28 _dafny.Int) _dafny.CodePoint {
				return _466_v176
			}
		})(_317_v176))), 15)
		(_nw86).ArraySet1((func() _dafny.Sequence {
			if _76_v0 {
				return _461_v294
			}
			return _461_v294
		})(), 16)
		(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ifdyyjm"), 17)
		(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ibqnypvji"), 18)
		(_nw86).ArraySet1(_461_v294, 19)
		(_nw86).ArraySet1(_461_v294, 20)
		(_nw86).ArraySet1(_461_v294, 21)
		(_nw86).ArraySet1(_461_v294, 22)
		_462_v295 = _nw86
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_462_v295), 0))
		_ = _index75
		(_462_v295).ArraySet1(_461_v294, (_index75).Int())
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_462_v295), 0))
		_ = _index76
		(_462_v295).ArraySet1(Companion_Default___.Fm12(_78_globalState), (_index76).Int())
		_77_v1 = (func() _dafny.Int {
			if _76_v0 {
				return _77_v1
			}
			return _dafny.IntOfInt64(-181)
		})()
		var _468_v296 _dafny.Sequence
		_ = _468_v296
		_468_v296 = _dafny.SeqOf(_dafny.IntOfInt64(119), _dafny.IntOfInt64(411))
		var _469_v297 _dafny.Map
		_ = _469_v297
		_469_v297 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_461_v294, _dafny.Companion_Sequence_.Update(_468_v296, (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_468_v296).Cardinality()))).Uint32(), _dafny.IntOfInt64(567)))
		_469_v297 = (_469_v297).Update((_462_v295).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_462_v295), 0))).Int()).(_dafny.Sequence), _468_v296)
		var _470_v298 _dafny.Array
		_ = _470_v298
		var _nwElement0_13 bool = _76_v0
		_ = _nwElement0_13
		var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(9))
		_ = _nw87
		(_nw87).ArraySet1(_nwElement0_13, 0)
		(_nw87).ArraySet1(_76_v0, 1)
		(_nw87).ArraySet1(!(!(_76_v0)), 2)
		(_nw87).ArraySet1(_76_v0, 3)
		(_nw87).ArraySet1(_76_v0, 4)
		(_nw87).ArraySet1(_76_v0, 5)
		(_nw87).ArraySet1(!(_76_v0), 6)
		(_nw87).ArraySet1(_76_v0, 7)
		(_nw87).ArraySet1(false, 8)
		_470_v298 = _nw87
		var _471_v299 _dafny.Map
		_ = _471_v299
		_471_v299 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _76_v0)
		var _472_v300 _dafny.Sequence
		_ = _472_v300
		_472_v300 = _dafny.SeqOf(_471_v299, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, false))
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_470_v298), 0))
		_ = _index77
		(_470_v298).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_471_v299), _472_v300), (_index77).Int())
		var _473_v301 D4
		_ = _473_v301
		_473_v301 = Companion_D4_.Create_DC11_(_462_v295)
		var _474_v302 _dafny.Map
		_ = _474_v302
		_474_v302 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _76_v0)
		var _475_v303 _dafny.Sequence
		_ = _475_v303
		_475_v303 = _dafny.SeqOf((func() bool {
			if (_474_v302).Contains(_76_v0) {
				return (_474_v302).Get(_76_v0).(bool)
			}
			return _76_v0
		})(), _76_v0)
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_470_v298), 0))
		_ = _index78
		var _rhs80 _dafny.Set = _79_v2
		_ = _rhs80
		var _rhs81 bool = (_475_v303).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_475_v303).Cardinality()))).Uint32()).(bool)
		_ = _rhs81
		var _rhs82 bool = _76_v0
		_ = _rhs82
		var _rhs83 D4 = Companion_D4_.Create_DC11_(_462_v295)
		_ = _rhs83
		var _rhs84 bool = _76_v0
		_ = _rhs84
		var _lhs39 _dafny.Array = _470_v298
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_470_v298), 0))
		_ = _lhs40
		_79_v2 = _rhs80
		_76_v0 = _rhs81
		(_lhs39).ArraySet1(_rhs82, (_lhs40).Int())
		_473_v301 = _rhs83
		_76_v0 = _rhs84
		_91_v12 = _91_v12
	}
	var _476_v304 _dafny.Sequence
	_ = _476_v304
	_476_v304 = _dafny.UnicodeSeqOfUtf8Bytes("wmlrc")
	var _477_v305 _dafny.Sequence
	_ = _477_v305
	_477_v305 = _dafny.SeqOf(false, _76_v0, _76_v0)
	var _478_v306 *C2
	_ = _478_v306
	var _nw88 *C2 = New_C2_()
	_ = _nw88
	_nw88.Ctor__(_77_v1)
	_478_v306 = _nw88
	var _479_v307 _dafny.Set
	_ = _479_v307
	_479_v307 = _dafny.SetOf(_478_v306, _478_v306)
	var _480_v308 _dafny.Map
	_ = _480_v308
	_480_v308 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v0, _76_v0)
	var _481_v309 _dafny.Sequence
	_ = _481_v309
	_481_v309 = _dafny.SeqOf(_dafny.IntOfInt64(899), _77_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dlokosuc")).Cardinality()))
	var _482_v310 _dafny.Set
	_ = _482_v310
	_482_v310 = _dafny.SetOf((_481_v309).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_481_v309).Cardinality()))).Uint32()).(_dafny.Int), _77_v1)
	var _483_v311 _dafny.Array
	_ = _483_v311
	var _nwElement0_14 bool = _76_v0
	_ = _nwElement0_14
	var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(28))
	_ = _nw89
	(_nw89).ArraySet1(_nwElement0_14, 0)
	(_nw89).ArraySet1(_76_v0, 1)
	(_nw89).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("obq"), _476_v304), 2)
	(_nw89).ArraySet1(_76_v0, 3)
	(_nw89).ArraySet1(_76_v0, 4)
	(_nw89).ArraySet1((func() bool {
		if _76_v0 {
			return _76_v0
		}
		return _76_v0
	})(), 5)
	(_nw89).ArraySet1(_76_v0, 6)
	(_nw89).ArraySet1((func() bool {
		if _76_v0 {
			return _76_v0
		}
		return _76_v0
	})(), 7)
	(_nw89).ArraySet1(_76_v0, 8)
	(_nw89).ArraySet1(_76_v0, 9)
	(_nw89).ArraySet1(((_dafny.MultiSetFromSeq(_477_v305)).Update(_76_v0, Companion_Default___.Abs(_77_v1))).IsDisjointFrom(_dafny.MultiSetOf(Companion_Default___.Fm0(_77_v1, (_479_v307).Cardinality(), _317_v176, _78_globalState), (func() bool {
		if (_480_v308).Contains(_76_v0) {
			return (_480_v308).Get(_76_v0).(bool)
		}
		return !(!(_76_v0))
	})(), _76_v0)), 10)
	(_nw89).ArraySet1((_79_v2).IsDisjointFrom(_79_v2), 11)
	(_nw89).ArraySet1(_76_v0, 12)
	(_nw89).ArraySet1(_76_v0, 13)
	(_nw89).ArraySet1(_76_v0, 14)
	(_nw89).ArraySet1((_482_v310).IsProperSubsetOf(_482_v310), 15)
	(_nw89).ArraySet1(_76_v0, 16)
	(_nw89).ArraySet1((_77_v1).Cmp((_478_v306).F3()) < 0, 17)
	(_nw89).ArraySet1(_76_v0, 18)
	(_nw89).ArraySet1((func() bool {
		if true {
			return !(_76_v0)
		}
		return _76_v0
	})(), 19)
	(_nw89).ArraySet1(_76_v0, 20)
	(_nw89).ArraySet1(true, 21)
	(_nw89).ArraySet1(_76_v0, 22)
	(_nw89).ArraySet1(_76_v0, 23)
	(_nw89).ArraySet1(_76_v0, 24)
	(_nw89).ArraySet1(_76_v0, 25)
	(_nw89).ArraySet1(_76_v0, 26)
	(_nw89).ArraySet1((_478_v306).Fm7(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}((func(_484_v0 bool) func(_dafny.Int) _dafny.Map {
		return func(_485_i29 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_484_v0, !(_484_v0), _484_v0), _484_v0)
		}
	})(_76_v0))), (Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}((func(_486_v0 bool) func(_dafny.Int) _dafny.Map {
		return func(_487_i29 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_486_v0, !(_486_v0), _486_v0), _486_v0)
		}
	})(_76_v0)))).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_76_v0, _76_v0, _76_v0), _76_v0)), _76_v0, _dafny.IntOfUint32((_477_v305).Cardinality()), _78_globalState), 27)
	_483_v311 = _nw89
	_483_v311 = _483_v311
	var _488_v312 _dafny.Sequence
	_ = _488_v312
	_488_v312 = _dafny.SeqOf(_476_v304, _dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}((func(_489_v176 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_490_i30 _dafny.Int) _dafny.CodePoint {
			return _489_v176
		}
	})(_317_v176))))
	var _491_v313 D6
	_ = _491_v313
	_491_v313 = Companion_D6_.Create_DC17_(_dafny.Companion_Sequence_.Concatenate(_488_v312, _488_v312))
	_491_v313 = _491_v313
	var _492_v314 _dafny.Array
	_ = _492_v314
	var _nw90 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(7))
	_ = _nw90
	_492_v314 = _nw90
	var _493_v315 _dafny.Array
	_ = _493_v315
	var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
	_ = _len0_18
	var _nw91 _dafny.Array
	_ = _nw91
	if _len0_18.Cmp(_dafny.Zero) == 0 {
		_nw91 = _dafny.NewArray(_len0_18)
	} else {
		var _init18 func(_dafny.Int) _dafny.Sequence = (func(_494_v309 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_495_i31 _dafny.Int) _dafny.Sequence {
				return _494_v309
			}
		})(_481_v309)
		_ = _init18
		var _element0_18 = _init18(_dafny.Zero)
		_ = _element0_18
		_nw91 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
		(_nw91).ArraySet1(_element0_18, 0)
		var _nativeLen0_18 = (_len0_18).Int()
		_ = _nativeLen0_18
		for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
			(_nw91).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
		}
	}
	_493_v315 = _nw91
	var _496_v316 T1
	_ = _496_v316
	var _nw92 *C0 = New_C0_()
	_ = _nw92
	_nw92.Ctor__(_493_v315)
	_496_v316 = _nw92
	var _497_v317 _dafny.Map
	_ = _497_v317
	_497_v317 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, _496_v316)
	var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_492_v314), 0))
	_ = _index79
	(_492_v314).ArraySet1(Companion_D4_.Create_DC10_((func() T1 {
		if (_497_v317).Contains(_77_v1) {
			return (_497_v317).Get(_77_v1).(T1)
		}
		return _496_v316
	})()), (_index79).Int())
	var _498_v318 _dafny.Array
	_ = _498_v318
	var _nw93 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
	_ = _nw93
	_498_v318 = _nw93
	var _499_v319 _dafny.Map
	_ = _499_v319
	_499_v319 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_498_v318, _91_v12)
	var _500_v320 D4
	_ = _500_v320
	_500_v320 = Companion_D4_.Create_DC10_(_496_v316)
	var _501_v321 _dafny.Map
	_ = _501_v321
	_501_v321 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_481_v309, (_478_v306).F3())
	var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_492_v314), 0))
	_ = _index80
	var _rhs85 _dafny.Int = _dafny.IntOfUint32((_476_v304).Cardinality())
	_ = _rhs85
	var _rhs86 _dafny.CodePoint = _317_v176
	_ = _rhs86
	var _rhs87 D4 = _500_v320
	_ = _rhs87
	var _rhs88 bool = !(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter6 := _dafny.Iterate((_501_v321).Keys().Elements()); ; {
			_compr_5, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _502_v322 _dafny.Sequence
			_502_v322 = interface{}(_compr_5).(_dafny.Sequence)
			if (_501_v321).Contains(_502_v322) {
				_coll5.Add(_502_v322)
			}
		}
		return _coll5.ToSet()
	}()).Contains(_481_v309)
	_ = _rhs88
	var _rhs89 _dafny.Map = _499_v319
	_ = _rhs89
	var _lhs41 _dafny.Array = _492_v314
	_ = _lhs41
	var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_492_v314), 0))
	_ = _lhs42
	_77_v1 = _rhs85
	_317_v176 = _rhs86
	(_lhs41).ArraySet1(_rhs87, (_lhs42).Int())
	_76_v0 = _rhs88
	_499_v319 = _rhs89
	var _503_v324 _dafny.Sequence
	_ = _503_v324
	_503_v324 = _dafny.SeqOf(func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter7 := _dafny.Iterate((_318_v177).Keys().Elements()); ; {
			_compr_6, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _504_v323 _dafny.CodePoint
			_504_v323 = interface{}(_compr_6).(_dafny.CodePoint)
			if (_318_v177).Contains(_504_v323) {
				_coll6.Add(_504_v323)
			}
		}
		return _coll6.ToSet()
	}())
	var _rhs90 _dafny.Array = _483_v311
	_ = _rhs90
	var _rhs91 _dafny.Sequence = _503_v324
	_ = _rhs91
	_483_v311 = _rhs90
	_503_v324 = _rhs91
	_77_v1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_478_v306).F3(), _77_v1), (_dafny.IntOfInt64(542)).Minus((_478_v306).F3())))
	var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_483_v311), 0))
	_ = _index81
	(_483_v311).ArraySet1((_77_v1).Cmp(_77_v1) != 0, (_index81).Int())
	var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_483_v311), 0))
	_ = _index82
	(_483_v311).ArraySet1(_76_v0, (_index82).Int())
	var _505_v325 _dafny.MultiSet
	_ = _505_v325
	_505_v325 = _dafny.MultiSetOf((_483_v311).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_483_v311), 0))).Int()).(bool))
	var _506_v326 _dafny.Map
	_ = _506_v326
	_506_v326 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463))
	var _507_v328 _dafny.Sequence
	_ = _507_v328
	_507_v328 = _dafny.SeqOf(_506_v326, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_478_v306).F3(), (_478_v306).F3()), _506_v326, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-381), _77_v1), func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter8 := _dafny.Iterate((_482_v310).Elements()); ; {
			_compr_7, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _508_v327 _dafny.Int
			_508_v327 = interface{}(_compr_7).(_dafny.Int)
			if (_482_v310).Contains(_508_v327) {
				_coll7.Add((_508_v327).Minus((_478_v306).F3()), (_478_v306).F3())
			}
		}
		return _coll7.ToMap()
	}())
	var _509_v329 D2
	_ = _509_v329
	_509_v329 = Companion_D2_.Create_DC4_(_476_v304)
	var _510_v330 _dafny.Array
	_ = _510_v330
	var _nwElement0_15 _dafny.Sequence = (_509_v329).Dtor_cf5()
	_ = _nwElement0_15
	var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(29))
	_ = _nw94
	(_nw94).ArraySet1(_nwElement0_15, 0)
	(_nw94).ArraySet1(_476_v304, 1)
	(_nw94).ArraySet1(_476_v304, 2)
	(_nw94).ArraySet1(_476_v304, 3)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _476_v304), 4)
	(_nw94).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pajtyxdav"), 5)
	(_nw94).ArraySet1(_476_v304, 6)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _dafny.UnicodeSeqOfUtf8Bytes("btgoc")), 7)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _dafny.UnicodeSeqOfUtf8Bytes("wvmwjuth")), 8)
	(_nw94).ArraySet1(_476_v304, 9)
	(_nw94).ArraySet1(_476_v304, 10)
	(_nw94).ArraySet1(_476_v304, 11)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _476_v304), 12)
	(_nw94).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pnb"), 13)
	(_nw94).ArraySet1(_476_v304, 14)
	(_nw94).ArraySet1(_476_v304, 15)
	(_nw94).ArraySet1((func() _dafny.Sequence {
		if (_483_v311).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_483_v311), 0))).Int()).(bool) {
			return _476_v304
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(673))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_511_v176 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_512_i32 _dafny.Int) _dafny.CodePoint {
				return _511_v176
			}
		})(_317_v176)))
	})(), 16)
	(_nw94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-614))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_513_i33 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), 17)
	(_nw94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_514_v176 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_515_i34 _dafny.Int) _dafny.CodePoint {
			return _514_v176
		}
	})(_317_v176))), 18)
	(_nw94).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jpv"), 19)
	(_nw94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(91))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer30(arg30)
		}
	}((func(_516_v176 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_517_i35 _dafny.Int) _dafny.CodePoint {
			return _516_v176
		}
	})(_317_v176))), 20)
	(_nw94).ArraySet1(_476_v304, 21)
	(_nw94).ArraySet1(_476_v304, 22)
	(_nw94).ArraySet1(_476_v304, 23)
	(_nw94).ArraySet1(_476_v304, 24)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _476_v304), 25)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Update(_476_v304, (Companion_Default___.SafeIndex((func() _dafny.Int {
		if (_94_v13).Contains(_77_v1) {
			return (_94_v13).Multiplicity(_77_v1)
		}
		return (_478_v306).F3()
	})(), _dafny.IntOfUint32((_476_v304).Cardinality()))).Uint32(), _317_v176), 26)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_476_v304, _476_v304), 27)
	(_nw94).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("mdd"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mdd")).Cardinality()))).Uint32(), _317_v176), 28)
	_510_v330 = _nw94
	var _518_v331 _dafny.Array
	_ = _518_v331
	var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
	_ = _nw95
	_518_v331 = _nw95
	var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_518_v331), 0))
	_ = _index83
	(_518_v331).ArraySet1(_483_v311, (_index83).Int())
	var _519_v333 *C0
	_ = _519_v333
	var _nw96 *C0 = New_C0_()
	_ = _nw96
	_nw96.Ctor__(_493_v315)
	_519_v333 = _nw96
	var _520_v334 _dafny.Map
	_ = _520_v334
	_520_v334 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v333, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v1, Companion_Default___.Fm1(_77_v1, _77_v1, _dafny.IntOfInt64(336), _76_v0, _78_globalState)))
	var _521_v335 _dafny.Map
	_ = _521_v335
	_521_v335 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_483_v311).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_483_v311), 0))).Int()).(bool), (func() _dafny.Map {
		if (_520_v334).Contains(_519_v333) {
			return (_520_v334).Get(_519_v333).(_dafny.Map)
		}
		return _506_v326
	})())
	var _522_v336 _dafny.Sequence
	_ = _522_v336
	_522_v336 = _dafny.SeqOf(_91_v12)
	var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_518_v331), 0))
	_ = _index84
	var _rhs92 _dafny.MultiSet = _505_v325
	_ = _rhs92
	var _rhs93 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_506_v326, _506_v326, Companion_Default___.Fm22(_77_v1, _76_v0, _78_globalState), _506_v326, func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(645), _dafny.IntOfInt64(744))); ; {
			_compr_8, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _523_v332 _dafny.Int
			_523_v332 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(645)).Cmp(_523_v332) <= 0) && ((_523_v332).Cmp(_dafny.IntOfInt64(744)) < 0) {
				_coll8.Add((_523_v332).Minus((_478_v306).F3()), _77_v1)
			}
		}
		return _coll8.ToMap()
	}()), _dafny.SeqOf(_506_v326, (func() _dafny.Map {
		if (_521_v335).Contains(_76_v0) {
			return (_521_v335).Get(_76_v0).(_dafny.Map)
		}
		return _506_v326
	})()))
	_ = _rhs93
	var _rhs94 _dafny.Array = _510_v330
	_ = _rhs94
	var _rhs95 _dafny.Array = (_522_v336).Select((Companion_Default___.SafeIndex(_77_v1, _dafny.IntOfUint32((_522_v336).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _rhs95
	var _rhs96 _dafny.Array = _483_v311
	_ = _rhs96
	var _lhs43 _dafny.Array = _518_v331
	_ = _lhs43
	var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_518_v331), 0))
	_ = _lhs44
	_505_v325 = _rhs92
	_507_v328 = _rhs93
	_510_v330 = _rhs94
	_91_v12 = _rhs95
	(_lhs43).ArraySet1(_rhs96, (_lhs44).Int())
	_77_v1 = Companion_Default___.SafeModuloInt(_77_v1, _77_v1)
	_dafny.Print(_76_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_globalState.F0).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v2).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v12).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v13).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-1), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v57).Equals(_dafny.SetOf(_dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v58).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v58).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v58).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_153_v59).Dtor_cf9()).Dtor_cf9()).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_153_v59).Dtor_cf9()).Dtor_cf9()).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_153_v59).Dtor_cf9()).Dtor_cf9()).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_317_v176)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v177).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_319_v178).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_319_v178).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_319_v178).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_476_v304.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_477_v305, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_478_v306).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_479_v307).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_480_v308).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_481_v309, _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_482_v310).Equals(_dafny.SetOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v311).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_488_v312, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wmlrc"), _dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_491_v313).Dtor_cf23(), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wmlrc"), _dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')), _dafny.UnicodeSeqOfUtf8Bytes("wmlrc"), _dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_493_v315).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_497_v317).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_499_v319).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_501_v321).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8)), _dafny.IntOfInt64(-181))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_503_v324, _dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_505_v325).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_506_v326).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_507_v328, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.One).UpdateUnsafe(_dafny.IntOfInt64(734), _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(826), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(827), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(828), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(829), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(830), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(831), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(832), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(833), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(834), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(835), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(836), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(837), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(838), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(839), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(840), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(841), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(842), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(843), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(844), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(845), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(846), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(847), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(848), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(849), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(850), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(851), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(852), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(853), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(854), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(855), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(856), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(857), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(858), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(859), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(860), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(861), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(862), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(863), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(864), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(865), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(866), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(867), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(868), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(869), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(870), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(871), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(872), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(873), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(874), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(875), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(876), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(877), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(878), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(879), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(880), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(881), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(882), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(883), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(884), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(885), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(886), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(887), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(888), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(889), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(890), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(891), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(892), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(893), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(894), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(895), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(896), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(897), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(898), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(899), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(900), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(901), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(902), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(903), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(904), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(905), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(906), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(907), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(908), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(909), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(910), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(911), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(912), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(913), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(914), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(915), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(916), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(917), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(918), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(919), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(920), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(921), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(922), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(923), _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(924), _dafny.Zero), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(-463)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_509_v329).Dtor_cf5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_510_v330).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_510_v330).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_510_v330).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_510_v330).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('q'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_510_v330).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_518_v331).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_519_v333.F2).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(-181), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_520_v334).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_521_v335).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_522_v336).Cardinality()))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
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

type D1_DC3 struct {
	Cf2 bool
	Cf3 bool
	Cf4 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 bool, Cf3 bool, Cf4 bool) D1 {
	return D1{D1_DC3{Cf2, Cf3, Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf1 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf1 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf1}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, false, false)
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf1() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
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

type D2_DC5 struct {
	Cf6 bool
	Cf7 _dafny.Int
	Cf8 _dafny.Int
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 bool, Cf7 _dafny.Int, Cf8 _dafny.Int) D2 {
	return D2{D2_DC5{Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC4 struct {
	Cf5 _dafny.Sequence
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Sequence) D2 {
	return D2{D2_DC4{Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC6 struct {
	Cf9 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 D2) D2 {
	return D2{D2_DC6{Cf9}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf9() D2 {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + data.Cf5.VerbatimString(true) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D3_DC8 struct {
	Cf11 bool
	Cf12 D2
	Cf13 bool
	Cf14 _dafny.Set
	Cf15 *C1
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf11 bool, Cf12 D2, Cf13 bool, Cf14 _dafny.Set, Cf15 *C1) D3 {
	return D3{D3_DC8{Cf11, Cf12, Cf13, Cf14, Cf15}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf16 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Int) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf10 _dafny.Set
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf10 _dafny.Set) D3 {
	return D3{D3_DC7{Cf10}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(false, Companion_D2_.Default(), false, _dafny.EmptySet, (*C1)(nil))
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC8).Cf11
}

func (_this D3) Dtor_cf12() D2 {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Set {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf15() *C1 {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) Dtor_cf10() _dafny.Set {
	return _this.Get_().(D3_DC7).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Equals(data2.Cf12) && data1.Cf13 == data2.Cf13 && data1.Cf14.Equals(data2.Cf14) && data1.Cf15 == data2.Cf15
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
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
	Cf18 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf18 _dafny.Array) D4 {
	return D4{D4_DC11{Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

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

type D4_DC10 struct {
	Cf17 T1
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf17 T1) D4 {
	return D4{D4_DC10{Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf18() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf17() T1 {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf18 == data2.Cf18
		}
	case D4_DC12:
		{
			_, ok := other.Get_().(D4_DC12)
			return ok
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && _dafny.AreEqual(data1.Cf17, data2.Cf17)
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

type D5_DC14 struct {
	Cf20 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 bool) D5 {
	return D5{D5_DC14{Cf20}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC15 struct {
	Cf21 bool
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf21 bool) D5 {
	return D5{D5_DC15{Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC13 struct {
	Cf19 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf19 _dafny.Map) D5 {
	return D5{D5_DC13{Cf19}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC16 struct {
	Cf22 D5
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf22 D5) D5 {
	return D5{D5_DC16{Cf22}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(false)
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf22() D5 {
	return _this.Get_().(D5_DC16).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf20 == data2.Cf20
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf21 == data2.Cf21
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D6_DC18 struct {
	Cf24 _dafny.Int
	Cf25 T0
	Cf26 D5
	Cf27 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf24 _dafny.Int, Cf25 T0, Cf26 D5, Cf27 _dafny.Int) D6 {
	return D6{D6_DC18{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC17 struct {
	Cf23 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf23 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf23}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC19 struct {
	Cf28 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf28 D6) D6 {
	return D6{D6_DC19{Cf28}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(_dafny.Zero, (T0)(nil), Companion_D5_.Default(), _dafny.Zero)
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf24
}

func (_this D6) Dtor_cf25() T0 {
	return _this.Get_().(D6_DC18).Cf25
}

func (_this D6) Dtor_cf26() D5 {
	return _this.Get_().(D6_DC18).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf23
}

func (_this D6) Dtor_cf28() D6 {
	return _this.Get_().(D6_DC19).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && _dafny.AreEqual(data1.Cf25, data2.Cf25) && data1.Cf26.Equals(data2.Cf26) && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf30 _dafny.MultiSet
	Cf31 _dafny.Int
	Cf32 bool
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 _dafny.MultiSet, Cf31 _dafny.Int, Cf32 bool) D7 {
	return D7{D7_DC21{Cf30, Cf31, Cf32}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC22 struct {
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_() D7 {
	return D7{D7_DC22{}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC23 struct {
	Cf33 _dafny.MultiSet
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf33 _dafny.MultiSet) D7 {
	return D7{D7_DC23{Cf33}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC20 struct {
	Cf29 *C0
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf29 *C0) D7 {
	return D7{D7_DC20{Cf29}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(_dafny.EmptyMultiSet, _dafny.Zero, false)
}

func (_this D7) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf31
}

func (_this D7) Dtor_cf32() bool {
	return _this.Get_().(D7_DC21).Cf32
}

func (_this D7) Dtor_cf33() _dafny.MultiSet {
	return _this.Get_().(D7_DC23).Cf33
}

func (_this D7) Dtor_cf29() *C0 {
	return _this.Get_().(D7_DC20).Cf29
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22"
		}
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30.Equals(data2.Cf30) && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D7_DC22:
		{
			_, ok := other.Get_().(D7_DC22)
			return ok
		}
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf33.Equals(data2.Cf33)
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf29 == data2.Cf29
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

type D8_DC25 struct {
	Cf35 _dafny.Map
	Cf36 _dafny.CodePoint
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf35 _dafny.Map, Cf36 _dafny.CodePoint) D8 {
	return D8{D8_DC25{Cf35, Cf36}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC24 struct {
	Cf34 _dafny.Array
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf34 _dafny.Array) D8 {
	return D8{D8_DC24{Cf34}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC25_(_dafny.EmptyMap, _dafny.CodePoint('D'))
}

func (_this D8) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D8_DC25).Cf35
}

func (_this D8) Dtor_cf36() _dafny.CodePoint {
	return _this.Get_().(D8_DC25).Cf36
}

func (_this D8) Dtor_cf34() _dafny.Array {
	return _this.Get_().(D8_DC24).Cf34
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf35.Equals(data2.Cf35) && data1.Cf36 == data2.Cf36
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf34 == data2.Cf34
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

type D9_DC27 struct {
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_() D9 {
	return D9{D9_DC27{}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC26 struct {
	Cf37 _dafny.Sequence
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf37 _dafny.Sequence) D9 {
	return D9{D9_DC26{Cf37}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC27_()
}

func (_this D9) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D9_DC26).Cf37
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC27:
		{
			return "D9.DC27"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC27:
		{
			_, ok := other.Get_().(D9_DC27)
			return ok
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf37.Equals(data2.Cf37)
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

// Definition of datatype D10
type D10 struct {
	Data_D10_
}

func (_this D10) Get_() Data_D10_ {
	return _this.Data_D10_
}

type Data_D10_ interface {
	isD10()
}

type CompanionStruct_D10_ struct {
}

var Companion_D10_ = CompanionStruct_D10_{}

type D10_DC29 struct {
	Cf39 _dafny.Int
	Cf40 _dafny.Int
	Cf41 bool
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf39 _dafny.Int, Cf40 _dafny.Int, Cf41 bool) D10 {
	return D10{D10_DC29{Cf39, Cf40, Cf41}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf38 _dafny.MultiSet
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf38 _dafny.MultiSet) D10 {
	return D10{D10_DC28{Cf38}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC30 struct {
	Cf42 D10
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf42 D10) D10 {
	return D10{D10_DC30{Cf42}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D10) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf39
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf40
}

func (_this D10) Dtor_cf41() bool {
	return _this.Get_().(D10_DC29).Cf41
}

func (_this D10) Dtor_cf38() _dafny.MultiSet {
	return _this.Get_().(D10_DC28).Cf38
}

func (_this D10) Dtor_cf42() D10 {
	return _this.Get_().(D10_DC30).Cf42
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf39.Cmp(data2.Cf39) == 0 && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41 == data2.Cf41
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf38.Equals(data2.Cf38)
		}
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D10) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D10)
	return ok && _this.Equals(typed)
}

func Type_D10_() _dafny.TypeDescriptor {
	return type_D10_{}
}

type type_D10_ struct {
}

func (_this type_D10_) Default() interface{} {
	return Companion_D10_.Default()
}

func (_this type_D10_) String() string {
	return "main.D10"
}
func (_this D10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D10{}

// End of datatype D10

// Definition of datatype D11
type D11 struct {
	Data_D11_
}

func (_this D11) Get_() Data_D11_ {
	return _this.Data_D11_
}

type Data_D11_ interface {
	isD11()
}

type CompanionStruct_D11_ struct {
}

var Companion_D11_ = CompanionStruct_D11_{}

type D11_DC32 struct {
	Cf44 bool
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf44 bool) D11 {
	return D11{D11_DC32{Cf44}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC31 struct {
	Cf43 _dafny.Map
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf43 _dafny.Map) D11 {
	return D11{D11_DC31{Cf43}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC33 struct {
	Cf45 D11
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf45 D11) D11 {
	return D11{D11_DC33{Cf45}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC32_(false)
}

func (_this D11) Dtor_cf44() bool {
	return _this.Get_().(D11_DC32).Cf44
}

func (_this D11) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D11_DC31).Cf43
}

func (_this D11) Dtor_cf45() D11 {
	return _this.Get_().(D11_DC33).Cf45
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf44 == data2.Cf44
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D11) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D11)
	return ok && _this.Equals(typed)
}

func Type_D11_() _dafny.TypeDescriptor {
	return type_D11_{}
}

type type_D11_ struct {
}

func (_this type_D11_) Default() interface{} {
	return Companion_D11_.Default()
}

func (_this type_D11_) String() string {
	return "main.D11"
}
func (_this D11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D11{}

// End of datatype D11

// Definition of datatype D12
type D12 struct {
	Data_D12_
}

func (_this D12) Get_() Data_D12_ {
	return _this.Data_D12_
}

type Data_D12_ interface {
	isD12()
}

type CompanionStruct_D12_ struct {
}

var Companion_D12_ = CompanionStruct_D12_{}

type D12_DC35 struct {
	Cf47 _dafny.Int
	Cf48 _dafny.CodePoint
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf47 _dafny.Int, Cf48 _dafny.CodePoint) D12 {
	return D12{D12_DC35{Cf47, Cf48}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_() D12 {
	return D12{D12_DC36{}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC34 struct {
	Cf46 _dafny.Map
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf46 _dafny.Map) D12 {
	return D12{D12_DC34{Cf46}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC35_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D12) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D12_DC35).Cf47
}

func (_this D12) Dtor_cf48() _dafny.CodePoint {
	return _this.Get_().(D12_DC35).Cf48
}

func (_this D12) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D12_DC34).Cf46
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48
		}
	case D12_DC36:
		{
			_, ok := other.Get_().(D12_DC36)
			return ok
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf46.Equals(data2.Cf46)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D12) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D12)
	return ok && _this.Equals(typed)
}

func Type_D12_() _dafny.TypeDescriptor {
	return type_D12_{}
}

type type_D12_ struct {
}

func (_this type_D12_) Default() interface{} {
	return Companion_D12_.Default()
}

func (_this type_D12_) String() string {
	return "main.D12"
}
func (_this D12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D12{}

// End of datatype D12

// Definition of trait T0
type T0 interface {
	String() string
	Fm7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
	Fm8(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
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

// Definition of trait T1
type T1 interface {
	String() string
	Fm7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
	Fm8(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
	Fm9(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Set
	Fm10(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool
}
type CompanionStruct_T1_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T1_ = CompanionStruct_T1_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T1_) CastTo_(x interface{}) T1 {
	var t T1
	t, _ = x.(T1)
	return t
}

// End of trait T1

// Definition of class GlobalState
type GlobalState struct {
	F0  _dafny.Map
	_f1 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptyMap
	_this._f1 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F2 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f2 _dafny.Array) {
	{
		(_this).F2 = f2
	}
}
func (_this *C0) Fm9(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngubj")).Cardinality())).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-318), false)).Cardinality()))
	}
}
func (_this *C0) Fm10(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C0) Fm7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(((_dafny.IntOfInt64(-806)).Plus(_dafny.IntOfInt64(781))).Cmp(_dafny.IntOfInt64(522)) != 0)
	}
}
func (_this *C0) Fm8(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(685)), _dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality())), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmlbrpfx")).Cardinality())).Times(_dafny.IntOfUint32(((Companion_D2_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("d"))).Dtor_cf5()).Cardinality())))
	}
}
func (_this *C0) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	{
		if (true) && (false) {
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))).Union(_dafny.MultiSetOf(!(!(true))))
		} else {
			return _dafny.MultiSetOf(true)
		}
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	dummy byte
}

func New_C1_() *C1 {
	_this := C1{}

	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm6(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return (!((func() bool {
			if !(false) {
				return false
			}
			return true
		})())) || (false)
	}
}
func (_this *C1) M1(globalState *GlobalState) (D0, _dafny.Int) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _524_v0 _dafny.Int
		_ = _524_v0
		_524_v0 = _dafny.IntOfInt64(786)
		var _525_v1 _dafny.MultiSet
		_ = _525_v1
		_525_v1 = _dafny.MultiSetOf(_524_v0)
		var _526_i0 _dafny.Int
		_ = _526_i0
		_526_i0 = _dafny.Zero
		{
			for ((_dafny.MultiSetOf(_524_v0, _524_v0)).Union(_525_v1)).IsSubsetOf(_525_v1) {
				{
					if (_526_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_526_i0 = (_526_i0).Plus(_dafny.One)
					var _527_v2 _dafny.Array
					_ = _527_v2
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_19
					var _nw97 _dafny.Array
					_ = _nw97
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw97 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) _dafny.Sequence = (func(_528_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
							return func(_529_i1 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(_528_v0, _528_v0)
							}
						})(_524_v0)
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw97 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw97).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw97).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_527_v2 = _nw97
					var _530_v3 *C0
					_ = _530_v3
					var _nw98 *C0 = New_C0_()
					_ = _nw98
					_nw98.Ctor__(_527_v2)
					_530_v3 = _nw98
					var _531_v4 *C0
					_ = _531_v4
					var _nw99 *C0 = New_C0_()
					_ = _nw99
					_nw99.Ctor__(_530_v3.F2)
					_531_v4 = _nw99
					r1 = _524_v0
					var _532_v5 bool
					_ = _532_v5
					_532_v5 = true
					var _533_v6 _dafny.Array
					_ = _533_v6
					var _nwElement0_16 _dafny.Int = _524_v0
					_ = _nwElement0_16
					var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(6))
					_ = _nw100
					(_nw100).ArraySet1(_nwElement0_16, 0)
					(_nw100).ArraySet1(_524_v0, 1)
					(_nw100).ArraySet1((_dafny.Zero).Minus(_524_v0), 2)
					(_nw100).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm12(globalState)).Cardinality()), 3)
					(_nw100).ArraySet1((_524_v0).Minus((_dafny.SetOf(_532_v5, _532_v5, _532_v5)).Cardinality()), 4)
					(_nw100).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_524_v0), _524_v0)), 5)
					_533_v6 = _nw100
					_533_v6 = _533_v6
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _534_v7 bool
		_ = _534_v7
		_534_v7 = true
		var _535_v8 _dafny.Sequence
		_ = _535_v8
		_535_v8 = _dafny.SeqOf(_534_v7)
		var _536_v9 _dafny.Sequence
		_ = _536_v9
		_536_v9 = _dafny.SeqOf(_535_v8, _535_v8, _535_v8, _535_v8, _535_v8)
		var _537_v10 _dafny.Map
		_ = _537_v10
		_537_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v7, _534_v7)
		if (_534_v7) && (_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.Zero).Minus(_524_v0), _dafny.IntOfInt64(-736), _524_v0, _dafny.IntOfUint32((_536_v9).Cardinality()), (_537_v10).Cardinality()), (Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_524_v0), _dafny.IntOfInt64(-736), _524_v0, _dafny.IntOfUint32((_536_v9).Cardinality()), (_537_v10).Cardinality())).Cardinality()))).Uint32(), _524_v0), _dafny.SeqOf(_524_v0))) {
			var _538_v11 _dafny.Array
			_ = _538_v11
			var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw101
			_538_v11 = _nw101
			var _539_v12 _dafny.Sequence
			_ = _539_v12
			_539_v12 = _dafny.UnicodeSeqOfUtf8Bytes("eelvyw")
			var _540_v13 _dafny.Sequence
			_ = _540_v13
			_540_v13 = _dafny.SeqOf(_539_v12, _539_v12)
			var _541_v14 _dafny.Map
			_ = _541_v14
			_541_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_540_v13).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_540_v13).Cardinality()))).Uint32()).(_dafny.Sequence), _538_v11)
			_538_v11 = (func() _dafny.Array {
				if false {
					return (func() _dafny.Array {
						if (_541_v14).Contains(_539_v12) {
							return (_541_v14).Get(_539_v12).(_dafny.Array)
						}
						return _538_v11
					})()
				}
				return _538_v11
			})()
			r1 = (_524_v0).Plus((_524_v0).Times(_524_v0))
			var _542_v15 _dafny.Array
			_ = _542_v15
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw102
			_542_v15 = _nw102
			var _543_v16 *C0
			_ = _543_v16
			var _nw103 *C0 = New_C0_()
			_ = _nw103
			_nw103.Ctor__(_542_v15)
			_543_v16 = _nw103
			var _544_v17 _dafny.Array
			_ = _544_v17
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_20
			var _nw104 _dafny.Array
			_ = _nw104
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw104 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Map = (func(_545_v0 _dafny.Int, _546_v7 bool) func(_dafny.Int) _dafny.Map {
					return func(_547_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_v0, _546_v7)
					}
				})(_524_v0, _534_v7)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw104 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw104).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw104).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_544_v17 = _nw104
			_544_v17 = _544_v17
			var _548_v18 D2
			_ = _548_v18
			_548_v18 = Companion_D2_.Create_DC6_(Companion_Default___.Fm13(globalState))
			var _source9 D2 = _548_v18
			_ = _source9
			if _source9.Is_DC5() {
				var _549___mcc_h0 bool = _source9.Get_().(D2_DC5).Cf6
				_ = _549___mcc_h0
				var _550___mcc_h1 _dafny.Int = _source9.Get_().(D2_DC5).Cf7
				_ = _550___mcc_h1
				var _551___mcc_h2 _dafny.Int = _source9.Get_().(D2_DC5).Cf8
				_ = _551___mcc_h2
				var _552_cf8 _dafny.Int = _551___mcc_h2
				_ = _552_cf8
				var _553_cf7 _dafny.Int = _550___mcc_h1
				_ = _553_cf7
				var _554_cf6 bool = _549___mcc_h0
				_ = _554_cf6
				_538_v11 = _538_v11
				var _555_v19 _dafny.Array
				_ = _555_v19
				var _nw105 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw105
				_555_v19 = _nw105
				_555_v19 = _555_v19
				var _556_v20 _dafny.CodePoint
				_ = _556_v20
				_556_v20 = _dafny.CodePoint('i')
				var _557_v21 _dafny.Map
				_ = _557_v21
				_557_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_cf6, _556_v20)
				var _558_v22 _dafny.Set
				_ = _558_v22
				_558_v22 = _dafny.SetOf(_534_v7)
				var _559_v23 _dafny.Map
				_ = _559_v23
				_559_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_525_v1, _558_v22)
				var _560_v24 _dafny.Map
				_ = _560_v24
				_560_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_535_v8).Cardinality()), _524_v0)
				var _561_v25 bool
				_ = _561_v25
				var _562_v26 bool
				_ = _562_v26
				var _563_v27 _dafny.Int
				_ = _563_v27
				var _564_v28 bool
				_ = _564_v28
				var _out42 bool
				_ = _out42
				var _out43 bool
				_ = _out43
				var _out44 _dafny.Int
				_ = _out44
				var _out45 bool
				_ = _out45
				_out42, _out43, _out44, _out45 = Companion_Default___.M0((func() _dafny.Map {
					if _554_cf6 {
						return _557_v21
					}
					return _557_v21
				})(), _559_v23, _534_v7, (_560_v24).Cardinality(), globalState)
				_561_v25 = _out42
				_562_v26 = _out43
				_563_v27 = _out44
				_564_v28 = _out45
				_534_v7 = Companion_Default___.Fm0(_553_cf7, (_552_cf8).Plus(_553_cf7), _556_v20, globalState)
			} else if _source9.Is_DC4() {
				var _565___mcc_h3 _dafny.Sequence = _source9.Get_().(D2_DC4).Cf5
				_ = _565___mcc_h3
				var _566_cf5 _dafny.Sequence = _565___mcc_h3
				_ = _566_cf5
				var _567_v29 _dafny.Set
				_ = _567_v29
				_567_v29 = _dafny.SetOf((_dafny.Zero).Minus(_524_v0))
				var _568_v30 _dafny.Sequence
				_ = _568_v30
				_568_v30 = _dafny.SeqOf(_567_v29, _567_v29)
				var _569_v31 _dafny.Sequence
				_ = _569_v31
				_569_v31 = _dafny.SeqOf(_524_v0, _524_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-389))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_570_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_571_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqOf(_570_v0, _570_v0)).Cardinality())
					}
				})(_524_v0)))).Cardinality()), _524_v0, (_567_v29).Cardinality())
				var _572_v32 _dafny.CodePoint
				_ = _572_v32
				_572_v32 = _dafny.CodePoint('p')
				var _573_v33 _dafny.Map
				_ = _573_v33
				_573_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(_572_v32, _534_v7, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(504))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_574_v12 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_575_i5 _dafny.Int) _dafny.CodePoint {
						return (_574_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_574_v12).Cardinality()), _dafny.IntOfUint32((_574_v12).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_539_v12)))).Cardinality()), _534_v7, globalState), _524_v0)
				var _576_v34 _dafny.Array
				_ = _576_v34
				var _nwElement0_17 bool = (func() bool {
					if _534_v7 {
						return _534_v7
					}
					return _534_v7
				})()
				_ = _nwElement0_17
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(17))
				_ = _nw106
				(_nw106).ArraySet1(_nwElement0_17, 0)
				(_nw106).ArraySet1(((_568_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_569_v31).Cardinality()), _dafny.IntOfUint32((_568_v30).Cardinality()))).Uint32()).(_dafny.Set)).IsSubsetOf(_567_v29), 1)
				(_nw106).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_577_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_578_i4 _dafny.Int) _dafny.Int {
						return _577_v0
					}
				})(_524_v0)))).Cardinality())).Cmp(_524_v0) <= 0, 2)
				(_nw106).ArraySet1(!(_534_v7), 3)
				(_nw106).ArraySet1(!(!(_dafny.MultiSetOf(_572_v32)).Contains(_572_v32)), 4)
				(_nw106).ArraySet1(_534_v7, 5)
				(_nw106).ArraySet1(_534_v7, 6)
				(_nw106).ArraySet1(true, 7)
				(_nw106).ArraySet1(!(_537_v10).Contains(_534_v7), 8)
				(_nw106).ArraySet1(_534_v7, 9)
				(_nw106).ArraySet1(_534_v7, 10)
				(_nw106).ArraySet1(_534_v7, 11)
				(_nw106).ArraySet1(_534_v7, 12)
				(_nw106).ArraySet1(_534_v7, 13)
				(_nw106).ArraySet1((_525_v1).IsProperSubsetOf(_525_v1), 14)
				(_nw106).ArraySet1(false, 15)
				(_nw106).ArraySet1(!(((_543_v16).Fm9(_534_v7, _573_v33, false, globalState)).Equals(_567_v29)), 16)
				_576_v34 = _nw106
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))
				_ = _index85
				(_576_v34).ArraySet1(_534_v7, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))
				_ = _index86
				(_576_v34).ArraySet1(_534_v7, (_index86).Int())
				var _579_v35 _dafny.Array
				_ = _579_v35
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw107
				_579_v35 = _nw107
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_579_v35), 0))
				_ = _index87
				(_579_v35).ArraySet1(Companion_Default___.SafeModuloInt(_524_v0, _524_v0), (_index87).Int())
				var _580_v36 _dafny.Map
				_ = _580_v36
				_580_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v0, (_524_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_539_v12).Cardinality()))))
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_579_v35), 0))
				_ = _index88
				(_579_v35).ArraySet1(_dafny.IntOfInt64(702), (_index88).Int())
				var _581_v37 _dafny.Map
				_ = _581_v37
				_581_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_524_v0, _524_v0, _572_v32, globalState), _566_cf5)
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_579_v35), 0))
				_ = _index89
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_579_v35), 0))
				_ = _index90
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))
				_ = _index91
				var _rhs97 _dafny.Int = (_dafny.IntOfUint32((_566_cf5).Cardinality())).Minus(_524_v0)
				_ = _rhs97
				var _rhs98 _dafny.Map = ((_580_v36).Merge(_580_v36)).Merge(_580_v36)
				_ = _rhs98
				var _rhs99 _dafny.Int = _524_v0
				_ = _rhs99
				var _rhs100 bool = (func() bool {
					if _534_v7 {
						return (_534_v7) || (_534_v7)
					}
					return !((_567_v29).IsProperSubsetOf(_567_v29))
				})()
				_ = _rhs100
				var _rhs101 bool = (func() bool {
					if _534_v7 {
						return ((_576_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))).Int()).(bool)) && (_534_v7)
					}
					return (_543_v16).Fm10(_534_v7, _524_v0, (_543_v16).Fm8(_569_v31, _534_v7, _524_v0, globalState), (_581_v37).Cardinality(), globalState)
				})()
				_ = _rhs101
				var _lhs45 _dafny.Array = _579_v35
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_579_v35), 0))
				_ = _lhs46
				var _lhs47 _dafny.Array = _579_v35
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_579_v35), 0))
				_ = _lhs48
				var _lhs49 _dafny.Array = _576_v34
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))
				_ = _lhs50
				(_lhs45).ArraySet1(_rhs97, (_lhs46).Int())
				_580_v36 = _rhs98
				(_lhs47).ArraySet1(_rhs99, (_lhs48).Int())
				_534_v7 = _rhs100
				(_lhs49).ArraySet1(_rhs101, (_lhs50).Int())
				var _582_v38 *C0
				_ = _582_v38
				var _nw108 *C0 = New_C0_()
				_ = _nw108
				_nw108.Ctor__(_542_v15)
				_582_v38 = _nw108
				var _583_v39 _dafny.Array
				_ = _583_v39
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_21
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.MultiSet = (func(_584_v0 _dafny.Int, _585_v8 _dafny.Sequence, _586_cf5 _dafny.Sequence, _587_v7 bool, _588_v35 _dafny.Array, _589_v1 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_590_i6 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_584_v0, (_dafny.MultiSetFromSeq(_585_v8)).Cardinality(), _dafny.IntOfUint32((_586_cf5).Cardinality()), (Companion_D2_.Create_DC5_(_587_v7, _dafny.IntOfInt64(468), (_588_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_588_v35), 0))).Int()).(_dafny.Int))).Dtor_cf8(), (_588_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_588_v35), 0))).Int()).(_dafny.Int))).Intersection(_589_v1)
						}
					})(_524_v0, _535_v8, _566_cf5, _534_v7, _579_v35, _525_v1)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw109 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw109).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw109).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_583_v39 = _nw109
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_583_v39), 0))
				_ = _index92
				(_583_v39).ArraySet1(_dafny.MultiSetOf((_579_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_579_v35), 0))).Int()).(_dafny.Int)), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_583_v39), 0))
				_ = _index93
				(_583_v39).ArraySet1((func() _dafny.MultiSet {
					if (_576_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_576_v34), 0))).Int()).(bool) {
						return _525_v1
					}
					return (_525_v1).Union(_525_v1)
				})(), (_index93).Int())
			} else {
				var _591___mcc_h4 D2 = _source9.Get_().(D2_DC6).Cf9
				_ = _591___mcc_h4
				var _592_cf9 D2 = _591___mcc_h4
				_ = _592_cf9
				var _593_v40 _dafny.Map
				_ = _593_v40
				_593_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_524_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_524_v0))), _524_v0)
				_593_v40 = (_593_v40).Update(_524_v0, _524_v0)
				var _rhs102 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_539_v12, _539_v12), _dafny.Companion_Sequence_.Concatenate(_539_v12, _539_v12))
				_ = _rhs102
				var _rhs103 _dafny.Int = _524_v0
				_ = _rhs103
				_539_v12 = _rhs102
				_524_v0 = _rhs103
				var _594_v41 _dafny.Sequence
				_ = _594_v41
				_594_v41 = _dafny.SeqOf((_537_v10).Cardinality(), _dafny.IntOfInt64(288))
				var _595_v42 _dafny.Array
				_ = _595_v42
				var _nwElement0_18 _dafny.Int = _dafny.IntOfInt64(858)
				_ = _nwElement0_18
				var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(7))
				_ = _nw110
				(_nw110).ArraySet1(_nwElement0_18, 0)
				(_nw110).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_594_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.IntOfUint32((_594_v41).Cardinality()))).Uint32(), _524_v0)).Cardinality()), 1)
				(_nw110).ArraySet1(Companion_Default___.Fm1(_524_v0, _524_v0, _dafny.IntOfInt64(574), _534_v7, globalState), 2)
				(_nw110).ArraySet1(_524_v0, 3)
				(_nw110).ArraySet1(_dafny.IntOfUint32((_539_v12).Cardinality()), 4)
				(_nw110).ArraySet1(_524_v0, 5)
				(_nw110).ArraySet1(_dafny.IntOfInt64(351), 6)
				_595_v42 = _nw110
				var _596_v43 _dafny.CodePoint
				_ = _596_v43
				_596_v43 = _dafny.CodePoint('v')
				var _597_v44 _dafny.Map
				_ = _597_v44
				_597_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v42, _596_v43)
				_597_v44 = (_597_v44).Update(_595_v42, _dafny.CodePoint('u'))
				_524_v0 = _524_v0
			}
		} else {
			var _598_v45 _dafny.Array
			_ = _598_v45
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
			_ = _nw111
			_598_v45 = _nw111
			var _599_v46 *C0
			_ = _599_v46
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__(_598_v45)
			_599_v46 = _nw112
			_534_v7 = _534_v7
			var _600_v47 D0
			_ = _600_v47
			_600_v47 = Companion_D0_.Create_DC0_(_524_v0)
			var _source10 D0 = _600_v47
			_ = _source10
			if _source10.Is_DC1() {
				var _601_v48 _dafny.Array
				_ = _601_v48
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_22
				var _nw113 _dafny.Array
				_ = _nw113
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw113 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Int = (func(_602_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_603_i7 _dafny.Int) _dafny.Int {
							return (_603_i7).Times(_602_v0)
						}
					})(_524_v0)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw113 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw113).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw113).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_601_v48 = _nw113
				var _604_v49 _dafny.Map
				_ = _604_v49
				_604_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v7, _601_v48)
				var _605_v50 _dafny.CodePoint
				_ = _605_v50
				_605_v50 = _dafny.CodePoint('d')
				r1 = Companion_Default___.SafeDivisionInt(((_604_v49).Merge((_604_v49).Update(Companion_Default___.Fm0(_524_v0, _524_v0, _605_v50, globalState), _601_v48))).Cardinality(), _524_v0)
				_534_v7 = ((_dafny.Zero).Minus(_524_v0)).Cmp(_524_v0) > 0
				var _606_v51 _dafny.Array
				_ = _606_v51
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw114
				_606_v51 = _nw114
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_606_v51), 0))
				_ = _index94
				(_606_v51).ArraySet1((_535_v8).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_535_v8).Cardinality()))).Uint32()).(bool), (_index94).Int())
				var _607_v52 _dafny.Sequence
				_ = _607_v52
				_607_v52 = _dafny.UnicodeSeqOfUtf8Bytes("llor")
				var _608_v53 D2
				_ = _608_v53
				_608_v53 = Companion_D2_.Create_DC4_(_607_v52)
				var _609_v54 _dafny.Sequence
				_ = _609_v54
				_609_v54 = _dafny.SeqOf(_607_v52)
				var _pat_let_tv9 = _609_v54
				_ = _pat_let_tv9
				var _pat_let_tv10 = _524_v0
				_ = _pat_let_tv10
				var _pat_let_tv11 = _609_v54
				_ = _pat_let_tv11
				var _610_v55 _dafny.Sequence
				_ = _610_v55
				_610_v55 = _dafny.SeqOf(_dafny.IntOfUint32(((func(_pat_let4_0 D2) D2 {
					return func(_611_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let5_0 _dafny.Sequence) D2 {
							return func(_612_dt__update_hcf5_h0 _dafny.Sequence) D2 {
								return Companion_D2_.Create_DC4_(_612_dt__update_hcf5_h0)
							}(_pat_let5_0)
						}((_pat_let_tv9).Select((Companion_Default___.SafeIndex(_pat_let_tv10, _dafny.IntOfUint32((_pat_let_tv11).Cardinality()))).Uint32()).(_dafny.Sequence))
					}(_pat_let4_0)
				}(_608_v53)).Dtor_cf5()).Cardinality()))
				var _613_v56 _dafny.Set
				_ = _613_v56
				_613_v56 = _dafny.SetOf(_534_v7)
				var _614_v57 _dafny.MultiSet
				_ = _614_v57
				_614_v57 = _dafny.MultiSetOf(_534_v7)
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_606_v51), 0))
				_ = _index95
				var _rhs104 bool = (_dafny.IntOfUint32((_610_v55).Cardinality())).Cmp((_dafny.Zero).Minus(((_613_v56).Cardinality()).Times(_524_v0))) != 0
				_ = _rhs104
				var _rhs105 bool = (!(_534_v7)) || (((_599_v46).Fm11(_524_v0, _524_v0, _535_v8, globalState)).IsProperSubsetOf(_614_v57))
				_ = _rhs105
				var _rhs106 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(15), _dafny.IntOfInt64(389))
				_ = _rhs106
				var _lhs51 _dafny.Array = _606_v51
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_606_v51), 0))
				_ = _lhs52
				(_lhs51).ArraySet1(_rhs104, (_lhs52).Int())
				_534_v7 = _rhs105
				r1 = _rhs106
				_524_v0 = _524_v0
			} else {
				var _615___mcc_h5 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
				_ = _615___mcc_h5
				var _616_cf0 _dafny.Int = _615___mcc_h5
				_ = _616_cf0
				var _617_v58 _dafny.Sequence
				_ = _617_v58
				_617_v58 = _dafny.UnicodeSeqOfUtf8Bytes("sbodes")
				var _618_v59 D2
				_ = _618_v59
				_618_v59 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC4_(_617_v58))
				var _619_v60 _dafny.Map
				_ = _619_v60
				_619_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}(func(_620_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()), _618_v59)
				var _621_v61 D2
				_ = _621_v61
				_621_v61 = Companion_D2_.Create_DC6_((func() D2 {
					if (_619_v60).Contains(_524_v0) {
						return (_619_v60).Get(_524_v0).(D2)
					}
					return _618_v59
				})())
				var _622_v62 _dafny.Array
				_ = _622_v62
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
				_ = _nw115
				_622_v62 = _nw115
				var _623_v63 _dafny.Map
				_ = _623_v63
				_623_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v0, _616_cf0)
				var _624_v64 _dafny.Sequence
				_ = _624_v64
				_624_v64 = _dafny.SeqOf((_dafny.Zero).Minus((func() _dafny.Int {
					if (_623_v63).Contains(_dafny.IntOfInt64(295)) {
						return (_623_v63).Get(_dafny.IntOfInt64(295)).(_dafny.Int)
					}
					return _616_cf0
				})()), _616_cf0, Companion_Default___.Fm1(_524_v0, _524_v0, _dafny.IntOfInt64(-203), _534_v7, globalState), _524_v0)
				var _625_v65 _dafny.Map
				_ = _625_v65
				_625_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_624_v64)).IsProperSubsetOf(_525_v1), _524_v0)
				var _rhs107 _dafny.Int = (_625_v65).Cardinality()
				_ = _rhs107
				var _rhs108 D2 = Companion_D2_.Create_DC6_(_618_v59)
				_ = _rhs108
				var _rhs109 _dafny.Array = _622_v62
				_ = _rhs109
				var _rhs110 bool = ((_534_v7) && (_534_v7)) == (!((_dafny.SetOf(_534_v7)).IsSubsetOf(_dafny.SetOf(false, _534_v7))))
				_ = _rhs110
				r1 = _rhs107
				_621_v61 = _rhs108
				_622_v62 = _rhs109
				_534_v7 = _rhs110
				var _626_v66 _dafny.Array
				_ = _626_v66
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw116
				_626_v66 = _nw116
				var _627_v67 _dafny.Set
				_ = _627_v67
				_627_v67 = _dafny.SetOf(_534_v7, (_535_v8).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_616_cf0), _dafny.IntOfUint32((_535_v8).Cardinality()))).Uint32()).(bool), _534_v7, _534_v7)
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_626_v66), 0))
				_ = _index96
				(_626_v66).ArraySet1((_627_v67).Cardinality(), (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_626_v66), 0))
				_ = _index97
				(_626_v66).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}(func(_628_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality())).Minus((_616_cf0).Plus(_616_cf0)), (_index97).Int())
				var _629_v68 *C0
				_ = _629_v68
				var _nw117 *C0 = New_C0_()
				_ = _nw117
				_nw117.Ctor__(_599_v46.F2)
				_629_v68 = _nw117
				var _630_v69 D0
				_ = _630_v69
				_630_v69 = Companion_D0_.Create_DC1_()
				var _631_v70 _dafny.Array
				_ = _631_v70
				var _nwElement0_19 bool = _534_v7
				_ = _nwElement0_19
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(11))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_19, 0)
				(_nw118).ArraySet1(_534_v7, 1)
				(_nw118).ArraySet1(_534_v7, 2)
				(_nw118).ArraySet1(_534_v7, 3)
				(_nw118).ArraySet1(_534_v7, 4)
				(_nw118).ArraySet1(!(_534_v7), 5)
				(_nw118).ArraySet1(_534_v7, 6)
				(_nw118).ArraySet1(_534_v7, 7)
				(_nw118).ArraySet1(_534_v7, 8)
				(_nw118).ArraySet1(_534_v7, 9)
				(_nw118).ArraySet1(_534_v7, 10)
				_631_v70 = _nw118
				var _632_v71 _dafny.Map
				_ = _632_v71
				_632_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_v69, _631_v70)
				_632_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_v69, _631_v70)
			}
			var _633_v72 _dafny.CodePoint
			_ = _633_v72
			_633_v72 = _dafny.CodePoint('v')
			_633_v72 = _633_v72
			var _634_v73 *C0
			_ = _634_v73
			var _nw119 *C0 = New_C0_()
			_ = _nw119
			_nw119.Ctor__(_598_v45)
			_634_v73 = _nw119
		}
		var _635_v74 _dafny.Array
		_ = _635_v74
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_23
		var _nw120 _dafny.Array
		_ = _nw120
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw120 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) _dafny.Int = (func(_636_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_637_i10 _dafny.Int) _dafny.Int {
					return (_637_i10).Minus(_636_v0)
				}
			})(_524_v0)
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw120 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw120).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw120).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_635_v74 = _nw120
		var _638_v75 _dafny.Map
		_ = _638_v75
		_638_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v7, _524_v0)
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))
		_ = _index98
		(_635_v74).ArraySet1(Companion_Default___.Fm1((_dafny.Zero).Minus(_524_v0), (_638_v75).Cardinality(), _dafny.IntOfInt64(846), _534_v7, globalState), (_index98).Int())
		var _639_v76 _dafny.Sequence
		_ = _639_v76
		_639_v76 = _dafny.UnicodeSeqOfUtf8Bytes("irlyr")
		var _640_v77 _dafny.CodePoint
		_ = _640_v77
		_640_v77 = _dafny.CodePoint('r')
		var _641_v78 D1
		_ = _641_v78
		_641_v78 = Companion_D1_.Create_DC2_(_640_v77)
		var _pat_let_tv12 = _524_v0
		_ = _pat_let_tv12
		var _pat_let_tv13 = _524_v0
		_ = _pat_let_tv13
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))
		_ = _index99
		var _rhs111 _dafny.Int = _524_v0
		_ = _rhs111
		var _rhs112 _dafny.Map = (func() _dafny.Map {
			if _534_v7 {
				return _638_v75
			}
			return (_638_v75).Merge(_638_v75)
		})()
		_ = _rhs112
		var _rhs113 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_639_v76).Cardinality()))
		_ = _rhs113
		var _rhs114 _dafny.Int = func(_source11 D1) _dafny.Int {
			if _source11.Is_DC3() {
				var _642___mcc_h6 bool = _source11.Get_().(D1_DC3).Cf2
				_ = _642___mcc_h6
				var _643___mcc_h7 bool = _source11.Get_().(D1_DC3).Cf3
				_ = _643___mcc_h7
				var _644___mcc_h8 bool = _source11.Get_().(D1_DC3).Cf4
				_ = _644___mcc_h8
				var _645_cf4 bool = _644___mcc_h8
				_ = _645_cf4
				var _646_cf3 bool = _643___mcc_h7
				_ = _646_cf3
				var _647_cf2 bool = _642___mcc_h6
				_ = _647_cf2
				return _pat_let_tv12
			} else {
				var _648___mcc_h9 _dafny.CodePoint = _source11.Get_().(D1_DC2).Cf1
				_ = _648___mcc_h9
				var _649_cf1 _dafny.CodePoint = _648___mcc_h9
				_ = _649_cf1
				return _pat_let_tv13
			}
		}(_641_v78)
		_ = _rhs114
		var _lhs53 _dafny.Array = _635_v74
		_ = _lhs53
		var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))
		_ = _lhs54
		var _lhs55 *GlobalState = globalState
		_ = _lhs55
		(_lhs53).ArraySet1(_rhs111, (_lhs54).Int())
		_lhs55.F0 = _rhs112
		r1 = _rhs113
		r1 = _rhs114
		_524_v0 = _dafny.IntOfInt64(599)
		var _650_v79 _dafny.Array
		_ = _650_v79
		var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
		_ = _nw121
		_650_v79 = _nw121
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_650_v79), 0))); ; {
			_guard_loop_1, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _651_i11 _dafny.Int
			_651_i11 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_651_i11).Sign() != -1) && ((_651_i11).Cmp(_dafny.ArrayLen((_650_v79), 0)) < 0)) {
				(_650_v79).ArraySet1CodePoint(_640_v77, (_651_i11).Int())
			}
		}
		var _652_i12 _dafny.Int
		_ = _652_i12
		_652_i12 = _dafny.Zero
		{
			for _534_v7 {
				{
					if (_652_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_652_i12 = (_652_i12).Plus(_dafny.One)
					var _653_v80 _dafny.Array
					_ = _653_v80
					var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
					_ = _nw122
					_653_v80 = _nw122
					var _654_v81 _dafny.Array
					_ = _654_v81
					var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
					_ = _nw123
					_654_v81 = _nw123
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_653_v80), 0))
					_ = _index100
					(_653_v80).ArraySet1(_654_v81, (_index100).Int())
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_653_v80), 0))
					_ = _index101
					var _rhs115 _dafny.Array = _654_v81
					_ = _rhs115
					var _rhs116 _dafny.Int = _524_v0
					_ = _rhs116
					var _rhs117 bool = _534_v7
					_ = _rhs117
					var _lhs56 _dafny.Array = _653_v80
					_ = _lhs56
					var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_653_v80), 0))
					_ = _lhs57
					(_lhs56).ArraySet1(_rhs115, (_lhs57).Int())
					_524_v0 = _rhs116
					_534_v7 = _rhs117
					var _655_v82 _dafny.Map
					_ = _655_v82
					_655_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v7, (_639_v76).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_639_v76).Cardinality()))).Uint32()).(_dafny.CodePoint))
					var _656_v83 _dafny.Set
					_ = _656_v83
					_656_v83 = _dafny.SetOf(_534_v7, _534_v7, _534_v7, _534_v7)
					var _657_v84 _dafny.Map
					_ = _657_v84
					_657_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_524_v0), _656_v83)
					var _658_v85 _dafny.Map
					_ = _658_v85
					_658_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v0, _524_v0)
					var _659_v86 _dafny.Sequence
					_ = _659_v86
					_659_v86 = _dafny.SeqOf(_dafny.IntOfInt64(-204))
					var _660_v87 _dafny.Sequence
					_ = _660_v87
					_660_v87 = _dafny.SeqOf((_658_v85).Cardinality(), (_659_v86).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_659_v86).Cardinality()))).Uint32()).(_dafny.Int))
					var _661_v88 _dafny.MultiSet
					_ = _661_v88
					_661_v88 = _dafny.MultiSetOf(_534_v7)
					var _662_v89 _dafny.Array
					_ = _662_v89
					var _nwElement0_20 _dafny.Sequence = _660_v87
					_ = _nwElement0_20
					var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(13))
					_ = _nw124
					(_nw124).ArraySet1(_nwElement0_20, 0)
					(_nw124).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_659_v86).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbnakhyf")).Cardinality())), 1)
					(_nw124).ArraySet1(_659_v86, 2)
					(_nw124).ArraySet1(_659_v86, 3)
					(_nw124).ArraySet1(_660_v87, 4)
					(_nw124).ArraySet1(_659_v86, 5)
					(_nw124).ArraySet1(_660_v87, 6)
					(_nw124).ArraySet1(_dafny.Companion_Sequence_.Update(_659_v86, (Companion_Default___.SafeIndex((_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_659_v86).Cardinality()))).Uint32(), (_661_v88).Cardinality()), 7)
					(_nw124).ArraySet1(_659_v86, 8)
					(_nw124).ArraySet1(_659_v86, 9)
					(_nw124).ArraySet1(_660_v87, 10)
					(_nw124).ArraySet1(_659_v86, 11)
					(_nw124).ArraySet1(_659_v86, 12)
					_662_v89 = _nw124
					var _663_v90 *C0
					_ = _663_v90
					var _nw125 *C0 = New_C0_()
					_ = _nw125
					_nw125.Ctor__(_662_v89)
					_663_v90 = _nw125
					var _664_v91 _dafny.Map
					_ = _664_v91
					_664_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v90, _534_v7)
					var _665_v92 bool
					_ = _665_v92
					var _666_v93 bool
					_ = _666_v93
					var _667_v94 _dafny.Int
					_ = _667_v94
					var _668_v95 bool
					_ = _668_v95
					var _out46 bool
					_ = _out46
					var _out47 bool
					_ = _out47
					var _out48 _dafny.Int
					_ = _out48
					var _out49 bool
					_ = _out49
					_out46, _out47, _out48, _out49 = Companion_Default___.M0(_655_v82, _657_v84, (func() bool {
						if (_664_v91).Contains(_663_v90) {
							return (_664_v91).Get(_663_v90).(bool)
						}
						return !(_534_v7)
					})(), (_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), globalState)
					_665_v92 = _out46
					_666_v93 = _out47
					_667_v94 = _out48
					_668_v95 = _out49
					_666_v93 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(48))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg36 _dafny.Int) interface{} {
							return coer36(arg36)
						}
					}(func(_669_i13 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('u')
					})), _639_v76)
					var _670_v96 _dafny.Map
					_ = _670_v96
					_670_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), _534_v7)
					var _671_v97 _dafny.Map
					_ = _671_v97
					_671_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_659_v86, _dafny.SeqOf(_dafny.IntOfUint32((_659_v86).Cardinality()), (_670_v96).Cardinality(), _dafny.IntOfInt64(878), _667_v94, _dafny.IntOfInt64(951))), (_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int))
					_671_v97 = (_671_v97).Update(_dafny.Companion_Sequence_.Concatenate(_659_v86, _660_v87), Companion_Default___.Fm1((_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), (_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), (_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), _534_v7, globalState))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _672_v98 D0
		_ = _672_v98
		_672_v98 = Companion_D0_.Create_DC1_()
		r0 = _672_v98
		var _673_v99 _dafny.MultiSet
		_ = _673_v99
		_673_v99 = _dafny.MultiSetOf(_534_v7)
		var _674_v100 _dafny.Sequence
		_ = _674_v100
		_674_v100 = _dafny.SeqOf(Companion_Default___.Fm1(_524_v0, _524_v0, (_635_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_635_v74), 0))).Int()).(_dafny.Int), _534_v7, globalState), _524_v0)
		var _675_v101 _dafny.Map
		_ = _675_v101
		_675_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_673_v99).Cardinality(), _674_v100)
		r1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_675_v101).Contains((_674_v100).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_674_v100).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_675_v101).Get((_674_v100).Select((Companion_Default___.SafeIndex(_524_v0, _dafny.IntOfUint32((_674_v100).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Sequence)
			}
			return _674_v100
		})()).Cardinality())
		return r0, r1
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f3 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f3 = _dafny.Zero
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f3 _dafny.Int) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C2) Fm7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_this).F3()).Cmp(_dafny.IntOfInt64(327)) == 0
	}
}
func (_this *C2) Fm8(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C2) Fm17(p0 bool, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F3()).Plus((_dafny.Zero).Minus((_this).F3()))
	}
}
func (_this *C2) Fm18(p0 D2, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F3()
	}
}
func (_this *C2) F3() _dafny.Int {
	{
		return _this._f3
	}
}

// End of class C2
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
