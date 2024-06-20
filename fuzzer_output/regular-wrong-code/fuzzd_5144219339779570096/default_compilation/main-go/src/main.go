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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	var _source0 D8 = Companion_D8_.Create_DC23_()
	_ = _source0
	if _source0.Is_DC23() {
		return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(25))).Cardinality()) < 0
	} else if _source0.Is_DC24() {
		var _1___mcc_h0 bool = _source0.Get_().(D8_DC24).Cf45
		_ = _1___mcc_h0
		var _2_cf45 bool = _1___mcc_h0
		_ = _2_cf45
		return true
	} else if _source0.Is_DC22() {
		var _3___mcc_h1 _dafny.Sequence = _source0.Get_().(D8_DC22).Cf44
		_ = _3___mcc_h1
		var _4_cf44 _dafny.Sequence = _3___mcc_h1
		_ = _4_cf44
		return _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(true, false), false)
	} else {
		var _5___mcc_h2 D8 = _source0.Get_().(D8_DC25).Cf46
		_ = _5___mcc_h2
		var _6_cf46 D8 = _5___mcc_h2
		_ = _6_cf46
		return true
	}
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(112), !(true))).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-891), _dafny.IntOfInt64(290))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(636)).Minus((_dafny.SetOf(_dafny.IntOfInt64(-371), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qywhj")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(616))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(929))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(503))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, globalState *GlobalState) D1 {
	if _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(false, true, false, true, false), _dafny.SeqOf(false)) {
		return Companion_D1_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("dudmm"))
	} else {
		return Companion_D1_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("o"))
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(868), _dafny.IntOfInt64(482)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-324))), _dafny.SeqOf(_dafny.IntOfInt64(822))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, p1 D1, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _dafny.UnicodeSeqOfUtf8Bytes("rtqhqq")))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D7 = Companion_D7_.Create_DC21_(Companion_D7_.Create_DC20_(true, !(false), _dafny.IntOfInt64(-30), true))
	_ = _source1
	if _source1.Is_DC20() {
		var _8___mcc_h0 bool = _source1.Get_().(D7_DC20).Cf39
		_ = _8___mcc_h0
		var _9___mcc_h1 bool = _source1.Get_().(D7_DC20).Cf40
		_ = _9___mcc_h1
		var _10___mcc_h2 _dafny.Int = _source1.Get_().(D7_DC20).Cf41
		_ = _10___mcc_h2
		var _11___mcc_h3 bool = _source1.Get_().(D7_DC20).Cf42
		_ = _11___mcc_h3
		var _12_cf42 bool = _11___mcc_h3
		_ = _12_cf42
		var _13_cf41 _dafny.Int = _10___mcc_h2
		_ = _13_cf41
		var _14_cf40 bool = _9___mcc_h1
		_ = _14_cf40
		var _15_cf39 bool = _8___mcc_h0
		_ = _15_cf39
		return _dafny.CodePoint('b')
	} else if _source1.Is_DC19() {
		var _16___mcc_h4 _dafny.MultiSet = _source1.Get_().(D7_DC19).Cf38
		_ = _16___mcc_h4
		var _17_cf38 _dafny.MultiSet = _16___mcc_h4
		_ = _17_cf38
		return _dafny.CodePoint('w')
	} else {
		var _18___mcc_h5 D7 = _source1.Get_().(D7_DC21).Cf43
		_ = _18___mcc_h5
		var _19_cf43 D7 = _18___mcc_h5
		_ = _19_cf43
		return _dafny.CodePoint('k')
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(702)))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, true)).Intersection(_dafny.SetOf(!(!(!(true)))))).Union((_dafny.SetOf(false, false)).Union(_dafny.SetOf(true, true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D1 {
	var _source2 D1 = Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality()), _dafny.IntOfInt64(-624)))
	_ = _source2
	if _source2.Is_DC3() {
		var _20___mcc_h0 _dafny.Int = _source2.Get_().(D1_DC3).Cf3
		_ = _20___mcc_h0
		var _21___mcc_h1 _dafny.Int = _source2.Get_().(D1_DC3).Cf4
		_ = _21___mcc_h1
		var _22_cf4 _dafny.Int = _21___mcc_h1
		_ = _22_cf4
		var _23_cf3 _dafny.Int = _20___mcc_h0
		_ = _23_cf3
		if !(false) {
			return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_(_22_cf4, true, _22_cf4, true))
		} else {
			return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_22_cf4, _23_cf3))
		}
	} else if _source2.Is_DC4() {
		var _24___mcc_h2 bool = _source2.Get_().(D1_DC4).Cf5
		_ = _24___mcc_h2
		var _25___mcc_h3 _dafny.Int = _source2.Get_().(D1_DC4).Cf6
		_ = _25___mcc_h3
		var _26_cf6 _dafny.Int = _25___mcc_h3
		_ = _26_cf6
		var _27_cf5 bool = _24___mcc_h2
		_ = _27_cf5
		return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_(_26_cf6, _27_cf5, _26_cf6, _27_cf5)))
	} else if _source2.Is_DC5() {
		var _28___mcc_h4 _dafny.Int = _source2.Get_().(D1_DC5).Cf7
		_ = _28___mcc_h4
		var _29___mcc_h5 bool = _source2.Get_().(D1_DC5).Cf8
		_ = _29___mcc_h5
		var _30___mcc_h6 _dafny.Int = _source2.Get_().(D1_DC5).Cf9
		_ = _30___mcc_h6
		var _31___mcc_h7 bool = _source2.Get_().(D1_DC5).Cf10
		_ = _31___mcc_h7
		var _32_cf10 bool = _31___mcc_h7
		_ = _32_cf10
		var _33_cf9 _dafny.Int = _30___mcc_h6
		_ = _33_cf9
		var _34_cf8 bool = _29___mcc_h5
		_ = _34_cf8
		var _35_cf7 _dafny.Int = _28___mcc_h4
		_ = _35_cf7
		return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_((_dafny.Zero).Minus(_35_cf7), _32_cf10, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xexlenl")).Cardinality()), _34_cf8))))
	} else if _source2.Is_DC2() {
		var _36___mcc_h8 _dafny.Sequence = _source2.Get_().(D1_DC2).Cf2
		_ = _36___mcc_h8
		var _37_cf2 _dafny.Sequence = _36___mcc_h8
		_ = _37_cf2
		return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC2_(_37_cf2)))
	} else {
		var _38___mcc_h9 D1 = _source2.Get_().(D1_DC6).Cf11
		_ = _38___mcc_h9
		var _39_cf11 D1 = _38___mcc_h9
		_ = _39_cf11
		if false {
			return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("bno")))
		} else {
			return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC4_(false, _dafny.IntOfInt64(755)))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((func() bool {
		if true {
			return true
		}
		return false
	})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(true))), _dafny.SeqOf(true, false, !(false), false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("dm")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("yx")
	})(), _dafny.UnicodeSeqOfUtf8Bytes("badgldb"))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(22), _dafny.IntOfInt64(-190), _dafny.IntOfInt64(-212))).Union((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('i'))).Cardinality()), _dafny.IntOfInt64(-603))).Cardinality()))).Difference(_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(297), true)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-80)), _dafny.IntOfInt64(377), _dafny.IntOfInt64(-598))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 D2, globalState *GlobalState) _dafny.MultiSet {
	var _source3 D8 = Companion_D8_.Create_DC23_()
	_ = _source3
	if _source3.Is_DC23() {
		return _dafny.MultiSetOf(true)
	} else if _source3.Is_DC24() {
		var _40___mcc_h0 bool = _source3.Get_().(D8_DC24).Cf45
		_ = _40___mcc_h0
		var _41_cf45 bool = _40___mcc_h0
		_ = _41_cf45
		return _dafny.MultiSetOf(_41_cf45, false, _41_cf45, _41_cf45)
	} else if _source3.Is_DC22() {
		var _42___mcc_h1 _dafny.Sequence = _source3.Get_().(D8_DC22).Cf44
		_ = _42___mcc_h1
		var _43_cf44 _dafny.Sequence = _42___mcc_h1
		_ = _43_cf44
		return _dafny.MultiSetOf((Companion_D1_.Create_DC4_(true, _dafny.IntOfInt64(371))).Dtor_cf5(), false)
	} else {
		var _44___mcc_h2 D8 = _source3.Get_().(D8_DC25).Cf46
		_ = _44___mcc_h2
		var _45_cf46 D8 = _44___mcc_h2
		_ = _45_cf46
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Difference(_dafny.MultiSetOf(false))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Sequence, _dafny.Array, _dafny.Int) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var _46_v0 _dafny.CodePoint
	_ = _46_v0
	_46_v0 = _dafny.CodePoint('r')
	var _47_v1 D2
	_ = _47_v1
	_47_v1 = Companion_D2_.Create_DC10_(_dafny.IntOfInt64(-185), _46_v0)
	var _source4 D2 = _47_v1
	_ = _source4
	if _source4.Is_DC8() {
		var _48___mcc_h0 _dafny.Sequence = _source4.Get_().(D2_DC8).Cf13
		_ = _48___mcc_h0
		var _49_cf13 _dafny.Sequence = _48___mcc_h0
		_ = _49_cf13
		var _50_v2 bool
		_ = _50_v2
		_50_v2 = true
		var _51_v3 _dafny.Sequence
		_ = _51_v3
		_51_v3 = _dafny.SeqOf(_50_v2)
		var _52_v4 *C0
		_ = _52_v4
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_51_v3, _51_v3))
		_52_v4 = _nw0
		_52_v4 = _52_v4
		var _53_v5 _dafny.Sequence
		_ = _53_v5
		_53_v5 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_51_v3, _dafny.SeqOf(_50_v2)))
		var _54_v6 _dafny.Int
		_ = _54_v6
		_54_v6 = _dafny.IntOfInt64(978)
		_53_v5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_53_v5, (Companion_Default___.SafeIndex(Companion_Default___.Fm2(_dafny.MultiSetOf(_46_v0, _dafny.CodePoint('e')), _54_v6, _54_v6, false, globalState), _dafny.IntOfUint32((_53_v5).Cardinality()))).Uint32(), _dafny.SeqOf((_52_v4).F29(), _50_v2)), _53_v5)
		_52_v4 = _52_v4
		(globalState).F27 = _54_v6
	} else if _source4.Is_DC9() {
		var _55___mcc_h1 bool = _source4.Get_().(D2_DC9).Cf14
		_ = _55___mcc_h1
		var _56___mcc_h2 *C0 = _source4.Get_().(D2_DC9).Cf15
		_ = _56___mcc_h2
		var _57_cf15 *C0 = _56___mcc_h2
		_ = _57_cf15
		var _58_cf14 bool = _55___mcc_h1
		_ = _58_cf14
		var _59_v8 _dafny.Array
		_ = _59_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_60_cf15 *C0, _61_cf14 bool) func(_dafny.Int) _dafny.Int {
				return func(_62_i0 _dafny.Int) _dafny.Int {
					return (_62_i0).Minus((func() _dafny.Map {
						var _coll0 = _dafny.NewMapBuilder()
						_ = _coll0
						for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), true)).Keys().Elements()); ; {
							_compr_0, _ok0 := _iter0()
							if !_ok0 {
								break
							}
							var _63_v7 _dafny.Int
							_63_v7 = interface{}(_compr_0).(_dafny.Int)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), true)).Contains(_63_v7) {
								_coll0.Add(Companion_Default___.SafeDivisionInt(_63_v7, (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(284), _dafny.IntOfInt64(-907))).Cardinality())), (_60_cf15).F29())).Cardinality())).Cardinality()), _61_cf14)
							}
						}
						return _coll0.ToMap()
					}()).Cardinality())
				}
			})(_57_cf15, _58_cf14)
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
		_59_v8 = _nw1
		var _64_v9 D4
		_ = _64_v9
		_64_v9 = Companion_D4_.Create_DC12_(_59_v8)
		(globalState).F28 = (_64_v9).Dtor_cf19()
		if (_57_cf15).F29() {
			var _65_v10 _dafny.Int
			_ = _65_v10
			_65_v10 = _dafny.IntOfInt64(581)
			var _66_v11 _dafny.Sequence
			_ = _66_v11
			_66_v11 = _dafny.SeqOf(_65_v10)
			var _67_v12 _dafny.Array
			_ = _67_v12
			var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw2
			_67_v12 = _nw2
			var _68_v13 _dafny.Map
			_ = _68_v13
			_68_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v12, _57_cf15)
			_57_cf15 = (func() *C0 {
				if !((_65_v10).Cmp((_66_v11).Select((Companion_Default___.SafeIndex((_68_v13).Cardinality(), _dafny.IntOfUint32((_66_v11).Cardinality()))).Uint32()).(_dafny.Int)) >= 0) {
					return _57_cf15
				}
				return _57_cf15
			})()
			(globalState).F13 = (_66_v11).Select((Companion_Default___.SafeIndex(_65_v10, _dafny.IntOfUint32((_66_v11).Cardinality()))).Uint32()).(_dafny.Int)
			var _69_v14 _dafny.Array
			_ = _69_v14
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(5))
			_ = _nw3
			_69_v14 = _nw3
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_69_v14), 0))
			_ = _index0
			(_69_v14).ArraySet1(_47_v1, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_69_v14), 0))
			_ = _index1
			(_69_v14).ArraySet1(_47_v1, (_index1).Int())
			var _70_v15 _dafny.Map
			_ = _70_v15
			_70_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_cf15).F29(), (_57_cf15).F29())
			_70_v15 = (_70_v15).Update(Companion_Default___.Fm0(p0, Companion_Default___.Fm0(p0, false, globalState), globalState), _58_cf14)
			(globalState).F13 = _dafny.IntOfInt64(249)
		} else {
			var _71_v16 *C0
			_ = _71_v16
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(!((_57_cf15).F29()))
			_71_v16 = _nw4
			(globalState).F14 = (func() bool {
				if false {
					return true
				}
				return true
			})()
			var _72_v17 _dafny.Set
			_ = _72_v17
			_72_v17 = _dafny.SetOf(_58_cf14, (_57_cf15).F29())
			(globalState).F13 = (_72_v17).Cardinality()
			(globalState).F14 = (_57_cf15).F29()
			var _73_v18 _dafny.Int
			_ = _73_v18
			_73_v18 = _dafny.IntOfInt64(927)
			var _74_v19 _dafny.Sequence
			_ = _74_v19
			_74_v19 = _dafny.SeqOf(_73_v18, _73_v18, _73_v18)
			var _75_v20 _dafny.Set
			_ = _75_v20
			_75_v20 = _dafny.SetOf(_74_v19, _74_v19)
			(globalState).F14 = ((_75_v20).Cardinality()).Cmp(_73_v18) == 0
		}
		var _76_v21 _dafny.Int
		_ = _76_v21
		_76_v21 = _dafny.IntOfInt64(82)
		var _77_v22 _dafny.Map
		_ = _77_v22
		_77_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v21, (func() _dafny.Int {
			if (_57_cf15).F29() {
				return (_dafny.Zero).Minus(_76_v21)
			}
			return _76_v21
		})())
		var _78_v23 _dafny.Sequence
		_ = _78_v23
		_78_v23 = _dafny.SeqOf(_76_v21)
		_77_v22 = (_77_v22).Update(_76_v21, (_76_v21).Minus(_dafny.IntOfUint32((_78_v23).Cardinality())))
		var _79_v24 *C0
		_ = _79_v24
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(_58_cf14)
		_79_v24 = _nw5
	} else if _source4.Is_DC10() {
		var _80___mcc_h3 _dafny.Int = _source4.Get_().(D2_DC10).Cf16
		_ = _80___mcc_h3
		var _81___mcc_h4 _dafny.CodePoint = _source4.Get_().(D2_DC10).Cf17
		_ = _81___mcc_h4
		var _82_cf17 _dafny.CodePoint = _81___mcc_h4
		_ = _82_cf17
		var _83_cf16 _dafny.Int = _80___mcc_h3
		_ = _83_cf16
		var _84_v25 bool
		_ = _84_v25
		_84_v25 = false
		var _85_v26 _dafny.Array
		_ = _85_v26
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_1
		var _nw6 _dafny.Array
		_ = _nw6
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw6 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_86_v25 bool) func(_dafny.Int) bool {
				return func(_87_i1 _dafny.Int) bool {
					return _86_v25
				}
			})(_84_v25)
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
		_85_v26 = _nw6
		var _88_v27 _dafny.Sequence
		_ = _88_v27
		_88_v27 = _dafny.SeqOf(_84_v25, _84_v25)
		var _89_v28 _dafny.Sequence
		_ = _89_v28
		_89_v28 = _dafny.SeqOf(_88_v27)
		var _90_v29 _dafny.Map
		_ = _90_v29
		_90_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v28, _85_v26)
		var _91_v30 D5
		_ = _91_v30
		_91_v30 = Companion_D5_.Create_DC15_(p0, _84_v25, _84_v25, _84_v25, _85_v26)
		var _92_v31 _dafny.Array
		_ = _92_v31
		var _nwElement0_0 _dafny.Array = (func() _dafny.Array {
			if _84_v25 {
				return _85_v26
			}
			return _85_v26
		})()
		_ = _nwElement0_0
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_0, 0)
		(_nw7).ArraySet1(_85_v26, 1)
		(_nw7).ArraySet1(_85_v26, 2)
		(_nw7).ArraySet1(_85_v26, 3)
		(_nw7).ArraySet1(_85_v26, 4)
		(_nw7).ArraySet1(_85_v26, 5)
		(_nw7).ArraySet1(_85_v26, 6)
		(_nw7).ArraySet1(_85_v26, 7)
		(_nw7).ArraySet1(_85_v26, 8)
		(_nw7).ArraySet1(_85_v26, 9)
		(_nw7).ArraySet1(_85_v26, 10)
		(_nw7).ArraySet1(_85_v26, 11)
		(_nw7).ArraySet1(_85_v26, 12)
		(_nw7).ArraySet1(_85_v26, 13)
		(_nw7).ArraySet1(_85_v26, 14)
		(_nw7).ArraySet1((func() _dafny.Array {
			if (_90_v29).Contains(_89_v28) {
				return (_90_v29).Get(_89_v28).(_dafny.Array)
			}
			return _85_v26
		})(), 15)
		(_nw7).ArraySet1(_85_v26, 16)
		(_nw7).ArraySet1(_85_v26, 17)
		(_nw7).ArraySet1((_91_v30).Dtor_cf29(), 18)
		(_nw7).ArraySet1(_85_v26, 19)
		(_nw7).ArraySet1(_85_v26, 20)
		(_nw7).ArraySet1(_85_v26, 21)
		_92_v31 = _nw7
		var _93_v32 D0
		_ = _93_v32
		_93_v32 = Companion_D0_.Create_DC0_(_83_cf16, _83_cf16)
		_92_v31 = (func() _dafny.Array {
			if (_dafny.IntOfInt64(-176)).Cmp((func(_pat_let0_0 D0) D0 {
				return func(_94_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 _dafny.Int) D0 {
						return func(_95_dt__update_hcf0_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_95_dt__update_hcf0_h0, (_94_dt__update__tmp_h0).Dtor_cf1())
						}(_pat_let1_0)
					}(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfraklcw")).Cardinality()))
				}(_pat_let0_0)
			}(_93_v32)).Dtor_cf0()) < 0 {
				return _92_v31
			}
			return _92_v31
		})()
		(globalState).F14 = _84_v25
		var _96_v33 _dafny.Sequence
		_ = _96_v33
		_96_v33 = _dafny.SeqOf(_83_cf16)
		(globalState).F27 = (_96_v33).Select((Companion_Default___.SafeIndex(_83_cf16, _dafny.IntOfUint32((_96_v33).Cardinality()))).Uint32()).(_dafny.Int)
		var _97_v34 *C0
		_ = _97_v34
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__(_84_v25)
		_97_v34 = _nw8
		var _98_v35 _dafny.Map
		_ = _98_v35
		_98_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_cf16).Cmp(_83_cf16) != 0, _97_v34)
		var _rhs0 _dafny.Map = (_98_v35).Update((_97_v34).F29(), _97_v34)
		_ = _rhs0
		var _rhs1 _dafny.Int = Companion_Default___.Fm2(_dafny.MultiSetOf(_dafny.CodePoint('y'), _82_cf17, _82_cf17, Companion_Default___.Fm7(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_83_cf16)), globalState)), (func() _dafny.Int {
			if (_97_v34).F29() {
				return _83_cf16
			}
			return _83_cf16
		})(), (_dafny.Zero).Minus(_83_cf16), true, globalState)
		_ = _rhs1
		var _rhs2 bool = (_97_v34).F29()
		_ = _rhs2
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		_98_v35 = _rhs0
		_83_cf16 = _rhs1
		_lhs0.F14 = _rhs2
	} else {
		var _99___mcc_h5 _dafny.Set = _source4.Get_().(D2_DC7).Cf12
		_ = _99___mcc_h5
		var _100_cf12 _dafny.Set = _99___mcc_h5
		_ = _100_cf12
		var _101_v36 bool
		_ = _101_v36
		_101_v36 = true
		var _102_v37 *C0
		_ = _102_v37
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__((func() bool {
			if _101_v36 {
				return _101_v36
			}
			return _101_v36
		})())
		_102_v37 = _nw9
		var _103_v38 _dafny.MultiSet
		_ = _103_v38
		_103_v38 = _dafny.MultiSetOf((_102_v37).F29())
		var _104_v39 *C0
		_ = _104_v39
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__((_103_v38).Contains(_101_v36))
		_104_v39 = _nw10
		var _105_v40 *C0
		_ = _105_v40
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__((_102_v37).F29())
		_105_v40 = _nw11
		var _106_v41 _dafny.Int
		_ = _106_v41
		_106_v41 = _dafny.IntOfInt64(556)
		var _107_v42 D1
		_ = _107_v42
		_107_v42 = Companion_D1_.Create_DC4_((_105_v40).F29(), _106_v41)
		_101_v36 = (_107_v42).Dtor_cf5()
	}
	var _108_v43 bool
	_ = _108_v43
	_108_v43 = true
	var _109_i2 _dafny.Int
	_ = _109_i2
	_109_i2 = _dafny.Zero
	{
		for !(_108_v43) {
			{
				if (_109_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_109_i2 = (_109_i2).Plus(_dafny.One)
				(globalState).F14 = _108_v43
				(globalState).F14 = (func() bool {
					if !(_108_v43) {
						return (func() bool {
							if _108_v43 {
								return _108_v43
							}
							return _108_v43
						})()
					}
					return _108_v43
				})()
				var _110_v44 *C0
				_ = _110_v44
				var _nw12 *C0 = New_C0_()
				_ = _nw12
				_nw12.Ctor__(false)
				_110_v44 = _nw12
				var _111_v45 _dafny.MultiSet
				_ = _111_v45
				_111_v45 = _dafny.MultiSetOf(false)
				var _112_v46 _dafny.Sequence
				_ = _112_v46
				_112_v46 = _dafny.SeqOf((_110_v44).F29())
				var _113_v47 _dafny.Array
				_ = _113_v47
				var _nwElement0_1 _dafny.MultiSet = _111_v45
				_ = _nwElement0_1
				var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(25))
				_ = _nw13
				(_nw13).ArraySet1(_nwElement0_1, 0)
				(_nw13).ArraySet1(_111_v45, 1)
				(_nw13).ArraySet1(_111_v45, 2)
				(_nw13).ArraySet1(_dafny.MultiSetOf((_110_v44).F29()), 3)
				(_nw13).ArraySet1(_111_v45, 4)
				(_nw13).ArraySet1(_111_v45, 5)
				(_nw13).ArraySet1(_111_v45, 6)
				(_nw13).ArraySet1(_111_v45, 7)
				(_nw13).ArraySet1((Companion_D7_.Create_DC19_(_dafny.MultiSetFromSeq(_112_v46))).Dtor_cf38(), 8)
				(_nw13).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm0(p0, _108_v43, globalState)), 9)
				(_nw13).ArraySet1((_111_v45).Update((_110_v44).F29(), Companion_Default___.Abs(_dafny.IntOfInt64(518))), 10)
				(_nw13).ArraySet1(_dafny.MultiSetFromSeq(_112_v46), 11)
				(_nw13).ArraySet1(_111_v45, 12)
				(_nw13).ArraySet1(_111_v45, 13)
				(_nw13).ArraySet1(_111_v45, 14)
				(_nw13).ArraySet1(_111_v45, 15)
				(_nw13).ArraySet1(_dafny.MultiSetOf(!(_108_v43), (_110_v44).F29(), !(_108_v43), _108_v43), 16)
				(_nw13).ArraySet1(_111_v45, 17)
				(_nw13).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)), 18)
				(_nw13).ArraySet1(_dafny.MultiSetOf((_110_v44).F29(), _108_v43), 19)
				(_nw13).ArraySet1(_111_v45, 20)
				(_nw13).ArraySet1(_111_v45, 21)
				(_nw13).ArraySet1(_111_v45, 22)
				(_nw13).ArraySet1(_111_v45, 23)
				(_nw13).ArraySet1(_111_v45, 24)
				_113_v47 = _nw13
				var _114_v48 _dafny.Int
				_ = _114_v48
				_114_v48 = _dafny.IntOfInt64(2)
				var _115_v49 _dafny.Sequence
				_ = _115_v49
				_115_v49 = _dafny.SeqOf(_114_v48, _dafny.IntOfInt64(-565), _114_v48)
				var _116_v50 _dafny.Sequence
				_ = _116_v50
				_116_v50 = _dafny.SeqOf(_115_v49)
				var _117_v51 _dafny.Map
				_ = _117_v51
				_117_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v47, (_116_v50).Select((Companion_Default___.SafeIndex(_114_v48, _dafny.IntOfUint32((_116_v50).Cardinality()))).Uint32()).(_dafny.Sequence))
				_117_v51 = (_117_v51).Update(_113_v47, _115_v49)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _118_v52 _dafny.Int
	_ = _118_v52
	_118_v52 = _dafny.IntOfInt64(366)
	var _119_v53 _dafny.Set
	_ = _119_v53
	_119_v53 = _dafny.SetOf(_118_v52)
	var _120_v54 _dafny.Set
	_ = _120_v54
	_120_v54 = _dafny.SetOf((_119_v53).Cardinality())
	(globalState).F14 = ((_119_v53).Equals(_119_v53)) == (_108_v43)
	var _121_v55 _dafny.Map
	_ = _121_v55
	_121_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v43, _46_v0)
	_121_v55 = (_121_v55).Update(!(_108_v43), _46_v0)
	var _122_i3 _dafny.Int
	_ = _122_i3
	_122_i3 = _dafny.Zero
	{
		for _108_v43 {
			{
				if (_122_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_122_i3 = (_122_i3).Plus(_dafny.One)
				if _108_v43 {
					var _123_v56 _dafny.Map
					_ = _123_v56
					_123_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v43, (Companion_Default___.Fm9(_118_v52, globalState)).Cardinality())
					var _124_v57 *C0
					_ = _124_v57
					var _nw14 *C0 = New_C0_()
					_ = _nw14
					_nw14.Ctor__(_108_v43)
					_124_v57 = _nw14
					var _125_v58 _dafny.Set
					_ = _125_v58
					_125_v58 = _dafny.SetOf(_124_v57, _124_v57, _124_v57)
					_123_v56 = (_123_v56).Update((func() bool {
						if _108_v43 {
							return false
						}
						return _108_v43
					})(), (_125_v58).Cardinality())
					(globalState).F14 = (_124_v57).F29()
					var _126_v59 _dafny.Set
					_ = _126_v59
					_126_v59 = _dafny.SetOf((_124_v57).F29(), _108_v43)
					var _127_v60 _dafny.Sequence
					_ = _127_v60
					_127_v60 = _dafny.SeqOf(_118_v52)
					var _128_v61 *C0
					_ = _128_v61
					var _nw15 *C0 = New_C0_()
					_ = _nw15
					_nw15.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((Companion_Default___.Fm13((_124_v57).F29(), _126_v59, _127_v60, _118_v52, globalState)).Cardinality(), _118_v52), _127_v60))
					_128_v61 = _nw15
					var _129_v62 _dafny.Sequence
					_ = _129_v62
					_129_v62 = _dafny.SeqOf(_108_v43, (_128_v61).F29())
					var _130_v63 D4
					_ = _130_v63
					_130_v63 = Companion_D4_.Create_DC13_(_128_v61, _dafny.CodePoint('w'), _46_v0, _129_v62)
					var _131_v64 _dafny.Map
					_ = _131_v64
					_131_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v52, _128_v61)
					var _132_v65 _dafny.Array
					_ = _132_v65
					var _nwElement0_2 *C0 = (_130_v63).Dtor_cf20()
					_ = _nwElement0_2
					var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
					_ = _nw16
					(_nw16).ArraySet1(_nwElement0_2, 0)
					(_nw16).ArraySet1(_128_v61, 1)
					(_nw16).ArraySet1(_124_v57, 2)
					(_nw16).ArraySet1(_128_v61, 3)
					(_nw16).ArraySet1(_124_v57, 4)
					(_nw16).ArraySet1(_124_v57, 5)
					(_nw16).ArraySet1(_124_v57, 6)
					(_nw16).ArraySet1((func() *C0 {
						if (_128_v61).F29() {
							return _124_v57
						}
						return (func() *C0 {
							if (_131_v64).Contains(_118_v52) {
								return (_131_v64).Get(_118_v52).(*C0)
							}
							return _128_v61
						})()
					})(), 7)
					_132_v65 = _nw16
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_132_v65), 0))
					_ = _index2
					(_132_v65).ArraySet1(_124_v57, (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_132_v65), 0))
					_ = _index3
					(_132_v65).ArraySet1(_124_v57, (_index3).Int())
					r2 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(932))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg2 _dafny.Int) interface{} {
							return coer2(arg2)
						}
					}((func(_133_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_134_i4 _dafny.Int) _dafny.CodePoint {
							return _133_v0
						}
					})(_46_v0)))).Cardinality())).Times(_118_v52)
				} else {
					var _135_v66 _dafny.Array
					_ = _135_v66
					var _len0_2 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_2
					var _nw17 _dafny.Array
					_ = _nw17
					if _len0_2.Cmp(_dafny.Zero) == 0 {
						_nw17 = _dafny.NewArray(_len0_2)
					} else {
						var _init2 func(_dafny.Int) _dafny.Sequence = (func(_136_v0 _dafny.CodePoint, _137_v1 D2) func(_dafny.Int) _dafny.Sequence {
							return func(_138_i5 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(383))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg3 _dafny.Int) interface{} {
										return coer3(arg3)
									}
								}((func(_139_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_140_i6 _dafny.Int) _dafny.CodePoint {
										return _139_v0
									}
								})(_136_v0))), (Companion_Default___.SafeIndex((_137_v1).Dtor_cf16(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(383))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg4 _dafny.Int) interface{} {
										return coer4(arg4)
									}
								}((func(_141_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_142_i6 _dafny.Int) _dafny.CodePoint {
										return _141_v0
									}
								})(_136_v0)))).Cardinality()))).Uint32(), _136_v0)
							}
						})(_46_v0, _47_v1)
						_ = _init2
						var _element0_2 = _init2(_dafny.Zero)
						_ = _element0_2
						_nw17 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
						(_nw17).ArraySet1(_element0_2, 0)
						var _nativeLen0_2 = (_len0_2).Int()
						_ = _nativeLen0_2
						for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
							(_nw17).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
						}
					}
					_135_v66 = _nw17
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_135_v66), 0))
					_ = _index4
					(_135_v66).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("txf"), (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_135_v66), 0))
					_ = _index5
					(_135_v66).ArraySet1(p0, (_index5).Int())
					var _143_v67 _dafny.MultiSet
					_ = _143_v67
					_143_v67 = _dafny.MultiSetOf(_46_v0, (p0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_118_v52), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint))
					(globalState).F25 = Companion_Default___.Fm2(_143_v67, _118_v52, _118_v52, false, globalState)
					var _144_v68 *C0
					_ = _144_v68
					var _nw18 *C0 = New_C0_()
					_ = _nw18
					_nw18.Ctor__(_108_v43)
					_144_v68 = _nw18
					var _rhs3 *C0 = _144_v68
					_ = _rhs3
					var _rhs4 bool = _108_v43
					_ = _rhs4
					_144_v68 = _rhs3
					_108_v43 = _rhs4
					var _145_v69 _dafny.Array
					_ = _145_v69
					var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
					_ = _nw19
					_145_v69 = _nw19
					var _146_v70 _dafny.Map
					_ = _146_v70
					_146_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v69, _144_v68)
					var _147_v71 _dafny.Sequence
					_ = _147_v71
					_147_v71 = _dafny.SeqOf(_144_v68, (func() *C0 {
						if (_146_v70).Contains(_145_v69) {
							return (_146_v70).Get(_145_v69).(*C0)
						}
						return _144_v68
					})())
					var _148_v72 _dafny.Map
					_ = _148_v72
					_148_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v52, _144_v68)
					var _149_v73 D8
					_ = _149_v73
					_149_v73 = Companion_D8_.Create_DC22_(_147_v71)
					var _150_v74 _dafny.Map
					_ = _150_v74
					_150_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v43, _147_v71)
					var _151_v75 _dafny.Sequence
					_ = _151_v75
					_151_v75 = _dafny.SeqOf(!(!((_144_v68).F29())), true)
					var _152_v76 _dafny.Array
					_ = _152_v76
					var _nwElement0_3 _dafny.Sequence = _147_v71
					_ = _nwElement0_3
					var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(26))
					_ = _nw20
					(_nw20).ArraySet1(_nwElement0_3, 0)
					(_nw20).ArraySet1(_dafny.SeqOf(_144_v68, _144_v68), 1)
					(_nw20).ArraySet1(_147_v71, 2)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_144_v68), _dafny.Companion_Sequence_.Update(_147_v71, (Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_147_v71).Cardinality()))).Uint32(), _144_v68)), 3)
					(_nw20).ArraySet1(_dafny.SeqOf(_144_v68, _144_v68, _144_v68), 4)
					(_nw20).ArraySet1(_dafny.SeqOf(_144_v68, _144_v68, (func() *C0 {
						if (_148_v72).Contains(_118_v52) {
							return (_148_v72).Get(_118_v52).(*C0)
						}
						return _144_v68
					})(), _144_v68, _144_v68), 5)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_147_v71, _147_v71), 6)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_147_v71, (_149_v73).Dtor_cf44()), 7)
					(_nw20).ArraySet1(_dafny.SeqOf(_144_v68), 8)
					(_nw20).ArraySet1(_147_v71, 9)
					(_nw20).ArraySet1(_147_v71, 10)
					(_nw20).ArraySet1((func() _dafny.Sequence {
						if (_144_v68).F29() {
							return _147_v71
						}
						return _dafny.Companion_Sequence_.Update(_147_v71, (Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_147_v71).Cardinality()))).Uint32(), _144_v68)
					})(), 11)
					(_nw20).ArraySet1(_147_v71, 12)
					(_nw20).ArraySet1(_147_v71, 13)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_147_v71, _147_v71), 14)
					(_nw20).ArraySet1(_147_v71, 15)
					(_nw20).ArraySet1(_147_v71, 16)
					(_nw20).ArraySet1(_147_v71, 17)
					(_nw20).ArraySet1((func() _dafny.Sequence {
						if (_150_v74).Contains(_108_v43) {
							return (_150_v74).Get(_108_v43).(_dafny.Sequence)
						}
						return _147_v71
					})(), 18)
					(_nw20).ArraySet1((func() _dafny.Sequence {
						if (_144_v68).F29() {
							return _147_v71
						}
						return _147_v71
					})(), 19)
					(_nw20).ArraySet1(_147_v71, 20)
					(_nw20).ArraySet1(_147_v71, 21)
					(_nw20).ArraySet1(_147_v71, 22)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_147_v71, _147_v71), 23)
					(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_147_v71, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_135_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_135_v66), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_147_v71).Cardinality()))).Uint32(), (_147_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_151_v75).Cardinality()), _dafny.IntOfUint32((_147_v71).Cardinality()))).Uint32()).(*C0)), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_144_v68), (Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_dafny.SeqOf(_144_v68)).Cardinality()))).Uint32(), _144_v68)), 24)
					(_nw20).ArraySet1(_147_v71, 25)
					_152_v76 = _nw20
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_152_v76), 0))
					_ = _index6
					(_152_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_147_v71, _147_v71), (_index6).Int())
					var _153_v77 _dafny.Map
					_ = _153_v77
					_153_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v68, _108_v43)
					var _154_v78 _dafny.Set
					_ = _154_v78
					_154_v78 = _dafny.SetOf(!((_144_v68).F29()), (func() bool {
						if (_153_v77).Contains(_144_v68) {
							return (_153_v77).Get(_144_v68).(bool)
						}
						return !(false)
					})(), (_144_v68).F29(), (_144_v68).F29(), (_144_v68).F29())
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_152_v76), 0))
					_ = _index7
					var _rhs5 _dafny.Sequence = _147_v71
					_ = _rhs5
					var _rhs6 bool = ((_144_v68).F29()) && (!((_118_v52).Cmp((_154_v78).Cardinality()) != 0))
					_ = _rhs6
					var _rhs7 bool = false
					_ = _rhs7
					var _lhs1 _dafny.Array = _152_v76
					_ = _lhs1
					var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_152_v76), 0))
					_ = _lhs2
					var _lhs3 *GlobalState = globalState
					_ = _lhs3
					(_lhs1).ArraySet1(_rhs5, (_lhs2).Int())
					_lhs3.F14 = _rhs6
					_108_v43 = _rhs7
					_144_v68 = _144_v68
				}
				var _155_v79 *C0
				_ = _155_v79
				var _nw21 *C0 = New_C0_()
				_ = _nw21
				_nw21.Ctor__(_108_v43)
				_155_v79 = _nw21
				var _156_v80 _dafny.Sequence
				_ = _156_v80
				_156_v80 = _dafny.SeqOf(_155_v79)
				var _157_v81 D8
				_ = _157_v81
				_157_v81 = Companion_D8_.Create_DC22_(_156_v80)
				var _158_v82 _dafny.Array
				_ = _158_v82
				var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw22
				_158_v82 = _nw22
				var _159_v83 D5
				_ = _159_v83
				_159_v83 = Companion_D5_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(418))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}((func(_160_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_161_i7 _dafny.Int) _dafny.CodePoint {
						return _160_v0
					}
				})(_46_v0))), _108_v43, false, (_155_v79).F29(), _158_v82)
				var _162_v84 _dafny.Map
				_ = _162_v84
				_162_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v82, _159_v83)
				var _rhs8 _dafny.CodePoint = _46_v0
				_ = _rhs8
				var _rhs9 _dafny.Sequence = p0
				_ = _rhs9
				var _rhs10 _dafny.Int = _118_v52
				_ = _rhs10
				var _rhs11 D8 = _157_v81
				_ = _rhs11
				var _rhs12 _dafny.Int = ((_162_v84).Merge(_162_v84)).Cardinality()
				_ = _rhs12
				var _lhs4 *GlobalState = globalState
				_ = _lhs4
				var _lhs5 *GlobalState = globalState
				_ = _lhs5
				_lhs4.F2 = _rhs8
				r0 = _rhs9
				_118_v52 = _rhs10
				_157_v81 = _rhs11
				_lhs5.F17 = _rhs12
				var _163_v85 _dafny.MultiSet
				_ = _163_v85
				_163_v85 = _dafny.MultiSetOf(_108_v43)
				var _164_v86 _dafny.Map
				_ = _164_v86
				_164_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_155_v79).F29(), _118_v52)
				var _165_v87 _dafny.Map
				_ = _165_v87
				_165_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_118_v52), _118_v52)
				var _166_v88 _dafny.Sequence
				_ = _166_v88
				_166_v88 = _dafny.SeqOf(_118_v52)
				var _167_v89 _dafny.Sequence
				_ = _167_v89
				_167_v89 = _dafny.SeqOf(_108_v43, false)
				var _168_v90 _dafny.Array
				_ = _168_v90
				var _nwElement0_4 _dafny.Int = (_164_v86).Cardinality()
				_ = _nwElement0_4
				var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(11))
				_ = _nw23
				(_nw23).ArraySet1(_nwElement0_4, 0)
				(_nw23).ArraySet1((_dafny.Zero).Minus(_118_v52), 1)
				(_nw23).ArraySet1(_118_v52, 2)
				(_nw23).ArraySet1(_118_v52, 3)
				(_nw23).ArraySet1((_165_v87).Cardinality(), 4)
				(_nw23).ArraySet1(_118_v52, 5)
				(_nw23).ArraySet1(_dafny.IntOfUint32((_166_v88).Cardinality()), 6)
				(_nw23).ArraySet1(_dafny.IntOfUint32((_167_v89).Cardinality()), 7)
				(_nw23).ArraySet1(_118_v52, 8)
				(_nw23).ArraySet1(_118_v52, 9)
				(_nw23).ArraySet1((_dafny.Zero).Minus(_118_v52), 10)
				_168_v90 = _nw23
				var _169_v91 _dafny.Map
				_ = _169_v91
				_169_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v85).Cardinality(), _168_v90)
				_169_v91 = (_169_v91).Update((_118_v52).Times(_dafny.IntOfInt64(941)), _168_v90)
				var _170_v92 *C0
				_ = _170_v92
				var _nw24 *C0 = New_C0_()
				_ = _nw24
				_nw24.Ctor__((_155_v79).F29())
				_170_v92 = _nw24
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _hi0 _dafny.Int = Companion_Default___.SafeModuloInt((_120_v54).Cardinality(), _118_v52)
	_ = _hi0
	for _171_i8 := _118_v52; _171_i8.Cmp(_hi0) < 0; _171_i8 = _171_i8.Plus(_dafny.One) {
		if _108_v43 {
			var _172_v93 *C0
			_ = _172_v93
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__(Companion_Default___.Fm0(p0, _108_v43, globalState))
			_172_v93 = _nw25
			var _173_v94 _dafny.Sequence
			_ = _173_v94
			_173_v94 = _dafny.SeqOf(_172_v93, _172_v93, _172_v93)
			var _174_v95 _dafny.Array
			_ = _174_v95
			var _nwElement0_5 *C0 = _172_v93
			_ = _nwElement0_5
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(15))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_5, 0)
			(_nw26).ArraySet1(_172_v93, 1)
			(_nw26).ArraySet1(_172_v93, 2)
			(_nw26).ArraySet1(_172_v93, 3)
			(_nw26).ArraySet1(_172_v93, 4)
			(_nw26).ArraySet1(_172_v93, 5)
			(_nw26).ArraySet1(_172_v93, 6)
			(_nw26).ArraySet1(_172_v93, 7)
			(_nw26).ArraySet1(_172_v93, 8)
			(_nw26).ArraySet1((_173_v94).Select((Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_173_v94).Cardinality()))).Uint32()).(*C0), 9)
			(_nw26).ArraySet1(_172_v93, 10)
			(_nw26).ArraySet1(_172_v93, 11)
			(_nw26).ArraySet1(_172_v93, 12)
			(_nw26).ArraySet1(_172_v93, 13)
			(_nw26).ArraySet1(_172_v93, 14)
			_174_v95 = _nw26
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_174_v95), 0))
			_ = _index8
			(_174_v95).ArraySet1(_172_v93, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_174_v95), 0))
			_ = _index9
			(_174_v95).ArraySet1(_172_v93, (_index9).Int())
			var _175_v96 _dafny.Sequence
			_ = _175_v96
			_175_v96 = _dafny.SeqOf(_108_v43, Companion_Default___.Fm0(p0, _108_v43, globalState))
			var _176_v97 _dafny.MultiSet
			_ = _176_v97
			_176_v97 = _dafny.MultiSetOf(_118_v52)
			var _177_v98 _dafny.Map
			_ = _177_v98
			_177_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v52, (_172_v93).F29())
			var _178_v99 _dafny.Set
			_ = _178_v99
			_178_v99 = _dafny.SetOf((func() bool {
				if (_177_v98).Contains(_171_i8) {
					return (_177_v98).Get(_171_i8).(bool)
				}
				return !(true)
			})())
			var _179_v100 _dafny.Array
			_ = _179_v100
			var _nwElement0_6 bool = _108_v43
			_ = _nwElement0_6
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_6, 0)
			(_nw27).ArraySet1(_108_v43, 1)
			(_nw27).ArraySet1(_108_v43, 2)
			(_nw27).ArraySet1(false, 3)
			(_nw27).ArraySet1(_108_v43, 4)
			(_nw27).ArraySet1(_108_v43, 5)
			(_nw27).ArraySet1(_108_v43, 6)
			(_nw27).ArraySet1((_172_v93).F29(), 7)
			(_nw27).ArraySet1(true, 8)
			(_nw27).ArraySet1((_172_v93).F29(), 9)
			(_nw27).ArraySet1((_172_v93).F29(), 10)
			(_nw27).ArraySet1(_108_v43, 11)
			(_nw27).ArraySet1((_172_v93).F29(), 12)
			(_nw27).ArraySet1((_172_v93).F29(), 13)
			(_nw27).ArraySet1(_108_v43, 14)
			(_nw27).ArraySet1(_108_v43, 15)
			(_nw27).ArraySet1((_172_v93).F29(), 16)
			(_nw27).ArraySet1((_172_v93).F29(), 17)
			(_nw27).ArraySet1(!(_108_v43), 18)
			(_nw27).ArraySet1(true, 19)
			(_nw27).ArraySet1((_172_v93).F29(), 20)
			(_nw27).ArraySet1(true, 21)
			(_nw27).ArraySet1(_108_v43, 22)
			(_nw27).ArraySet1(true, 23)
			(_nw27).ArraySet1((_172_v93).F29(), 24)
			_179_v100 = _nw27
			var _180_v101 _dafny.Map
			_ = _180_v101
			_180_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v100, true)
			var _181_v102 _dafny.Array
			_ = _181_v102
			var _nwElement0_7 bool = _108_v43
			_ = _nwElement0_7
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(28))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_7, 0)
			(_nw28).ArraySet1((_175_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_173_v94).Cardinality()), _dafny.IntOfUint32((_175_v96).Cardinality()))).Uint32()).(bool), 1)
			(_nw28).ArraySet1((false) && (false), 2)
			(_nw28).ArraySet1((_172_v93).F29(), 3)
			(_nw28).ArraySet1(!(_dafny.MultiSetOf(_118_v52)).Equals(_176_v97), 4)
			(_nw28).ArraySet1((true) == (Companion_Default___.Fm0(_dafny.UnicodeSeqOfUtf8Bytes("shqrkmdpp"), true, globalState)), 5)
			(_nw28).ArraySet1(_108_v43, 6)
			(_nw28).ArraySet1(_108_v43, 7)
			(_nw28).ArraySet1((_172_v93).F29(), 8)
			(_nw28).ArraySet1((_172_v93).F29(), 9)
			(_nw28).ArraySet1(_108_v43, 10)
			(_nw28).ArraySet1(((_178_v99).Cardinality()).Cmp(_118_v52) < 0, 11)
			(_nw28).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Cmp(_dafny.IntOfInt64(-454)) < 0, 12)
			(_nw28).ArraySet1(_108_v43, 13)
			(_nw28).ArraySet1(Companion_Default___.Fm0(p0, !(false), globalState), 14)
			(_nw28).ArraySet1((_172_v93).F29(), 15)
			(_nw28).ArraySet1(_108_v43, 16)
			(_nw28).ArraySet1(!_dafny.Companion_Sequence_.Contains(_175_v96, !((func() bool {
				if (_180_v101).Contains(_179_v100) {
					return (_180_v101).Get(_179_v100).(bool)
				}
				return _108_v43
			})())), 17)
			(_nw28).ArraySet1(((_172_v93).F29()) || (_108_v43), 18)
			(_nw28).ArraySet1(_108_v43, 19)
			(_nw28).ArraySet1(_108_v43, 20)
			(_nw28).ArraySet1((_171_i8).Cmp((_dafny.Zero).Minus(_118_v52)) > 0, 21)
			(_nw28).ArraySet1(true, 22)
			(_nw28).ArraySet1((_172_v93).F29(), 23)
			(_nw28).ArraySet1(_108_v43, 24)
			(_nw28).ArraySet1(_108_v43, 25)
			(_nw28).ArraySet1((_172_v93).F29(), 26)
			(_nw28).ArraySet1((Companion_D6_.Create_DC18_(p0, _dafny.IntOfUint32((_175_v96).Cardinality()), _108_v43, _dafny.IntOfInt64(834))).Dtor_cf36(), 27)
			_181_v102 = _nw28
			_181_v102 = _179_v100
			var _182_v103 _dafny.Map
			_ = _182_v103
			_182_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(p0, _171_i8, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(848))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_183_v52 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_184_i9 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _183_v52)
				}
			})(_118_v52))))
			var _185_v104 _dafny.Map
			_ = _185_v104
			_185_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_172_v93).F29(), _171_i8)
			var _186_v105 _dafny.Sequence
			_ = _186_v105
			_186_v105 = _dafny.SeqOf(_185_v104)
			_182_v103 = (_182_v103).Update(_46_v0, _186_v105)
			var _187_v106 _dafny.MultiSet
			_ = _187_v106
			_187_v106 = _dafny.MultiSetOf((_172_v93).F29())
			var _188_v107 _dafny.Sequence
			_ = _188_v107
			_188_v107 = _dafny.SeqOf((func() _dafny.Int {
				if (_187_v106).Contains(!((_172_v93).F29())) {
					return (_187_v106).Multiplicity(!((_172_v93).F29()))
				}
				return _171_i8
			})(), _118_v52)
			var _189_v108 _dafny.Array
			_ = _189_v108
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw29
			_189_v108 = _nw29
			var _190_v109 _dafny.Set
			_ = _190_v109
			_190_v109 = _dafny.SetOf(_189_v108, _189_v108)
			(globalState).F27 = (_188_v107).Select((Companion_Default___.SafeIndex(((_190_v109).Intersection(_190_v109)).Cardinality(), _dafny.IntOfUint32((_188_v107).Cardinality()))).Uint32()).(_dafny.Int)
			var _191_v110 *C0
			_ = _191_v110
			var _nw30 *C0 = New_C0_()
			_ = _nw30
			_nw30.Ctor__(true)
			_191_v110 = _nw30
		} else {
			var _192_v111 _dafny.MultiSet
			_ = _192_v111
			_192_v111 = _dafny.MultiSetOf(true, _108_v43)
			(globalState).F17 = Companion_Default___.SafeDivisionInt(_171_i8, Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_192_v111).Contains(_108_v43) {
					return (_192_v111).Multiplicity(_108_v43)
				}
				return _118_v52
			})(), _dafny.IntOfInt64(892)))
			(globalState).F14 = _108_v43
			var _193_v112 _dafny.Sequence
			_ = _193_v112
			_193_v112 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kvspt"))
			var _194_v113 _dafny.Map
			_ = _194_v113
			_194_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_119_v53).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("yjrw"))
			var _195_v114 _dafny.Array
			_ = _195_v114
			var _nwElement0_8 _dafny.Sequence = (_193_v112).Select((Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_193_v112).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _nwElement0_8
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(20))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_8, 0)
			(_nw31).ArraySet1(p0, 1)
			(_nw31).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_196_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_197_i10 _dafny.Int) _dafny.CodePoint {
					return _196_v0
				}
			})(_46_v0))), 2)
			(_nw31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iattnj"), 3)
			(_nw31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 4)
			(_nw31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("amks"), 5)
			(_nw31).ArraySet1(p0, 6)
			(_nw31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 7)
			(_nw31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uut"), 8)
			(_nw31).ArraySet1(p0, 9)
			(_nw31).ArraySet1(p0, 10)
			(_nw31).ArraySet1(p0, 11)
			(_nw31).ArraySet1(p0, 12)
			(_nw31).ArraySet1(p0, 13)
			(_nw31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rivncvdvu"), 14)
			(_nw31).ArraySet1((_193_v112).Select((Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_193_v112).Cardinality()))).Uint32()).(_dafny.Sequence), 15)
			(_nw31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pukcc"), 16)
			(_nw31).ArraySet1(p0, 17)
			(_nw31).ArraySet1((func() _dafny.Sequence {
				if (_194_v113).Contains(_dafny.IntOfUint32((p0).Cardinality())) {
					return (_194_v113).Get(_dafny.IntOfUint32((p0).Cardinality())).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("mpehwcpgi")
			})(), 18)
			(_nw31).ArraySet1(p0, 19)
			_195_v114 = _nw31
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_195_v114), 0))
			_ = _index10
			(_195_v114).ArraySet1(p0, (_index10).Int())
			var _198_v115 _dafny.Map
			_ = _198_v115
			_198_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v43, _dafny.UnicodeSeqOfUtf8Bytes("cugx"))
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_195_v114), 0))
			_ = _index11
			(_195_v114).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_198_v115).Contains(_108_v43) {
					return (_198_v115).Get(_108_v43).(_dafny.Sequence)
				}
				return p0
			})(), p0)), (_index11).Int())
			(globalState).F17 = ((_118_v52).Minus(_171_i8)).Plus(_118_v52)
			var _199_v116 D2
			_ = _199_v116
			_199_v116 = Companion_D2_.Create_DC8_(p0)
			var _200_v117 _dafny.Sequence
			_ = _200_v117
			_200_v117 = _dafny.SeqOf(Companion_Default___.Fm14(_199_v116, globalState), _192_v111, _192_v111, _192_v111, _192_v111)
			var _201_v118 _dafny.Map
			_ = _201_v118
			_201_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v117).Select((Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_200_v117).Cardinality()))).Uint32()).(_dafny.MultiSet), _108_v43)
			_201_v118 = _201_v118
		}
		_118_v52 = _dafny.IntOfUint32((p0).Cardinality())
		var _202_v119 _dafny.Sequence
		_ = _202_v119
		_202_v119 = _dafny.SeqOf(_118_v52)
		_202_v119 = _202_v119
		var _203_v120 _dafny.Map
		_ = _203_v120
		_203_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v43, _171_i8)
		_203_v120 = _203_v120
	}
	r0 = p0
	var _204_v121 _dafny.MultiSet
	_ = _204_v121
	_204_v121 = _dafny.MultiSetOf(false)
	var _205_v122 _dafny.Sequence
	_ = _205_v122
	_205_v122 = _dafny.SeqOf((func() _dafny.Int {
		if (_204_v121).Contains(_108_v43) {
			return (_204_v121).Multiplicity(_108_v43)
		}
		return _118_v52
	})())
	var _206_v123 _dafny.Sequence
	_ = _206_v123
	_206_v123 = _dafny.SeqOf(_108_v43)
	var _207_v124 _dafny.Map
	_ = _207_v124
	_207_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v52, _dafny.IntOfUint32((_206_v123).Cardinality()))
	var _208_v125 _dafny.Set
	_ = _208_v125
	_208_v125 = _dafny.SetOf(_207_v124)
	var _209_v126 _dafny.MultiSet
	_ = _209_v126
	_209_v126 = _dafny.MultiSetOf(_118_v52, _118_v52)
	var _210_v127 _dafny.Map
	_ = _210_v127
	_210_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v52, _108_v43)
	var _211_v128 _dafny.Map
	_ = _211_v128
	_211_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v127, false)
	var _nwElement0_9 _dafny.Int = _dafny.IntOfUint32((_205_v122).Cardinality())
	_ = _nwElement0_9
	var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(25))
	_ = _nw32
	(_nw32).ArraySet1(_nwElement0_9, 0)
	(_nw32).ArraySet1(_dafny.IntOfUint32((_206_v123).Cardinality()), 1)
	(_nw32).ArraySet1(_118_v52, 2)
	(_nw32).ArraySet1((_208_v125).Cardinality(), 3)
	(_nw32).ArraySet1(_118_v52, 4)
	(_nw32).ArraySet1(_118_v52, 5)
	(_nw32).ArraySet1(_118_v52, 6)
	(_nw32).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_118_v52), _dafny.IntOfUint32((p0).Cardinality())), 7)
	(_nw32).ArraySet1((_dafny.Zero).Minus(_118_v52), 8)
	(_nw32).ArraySet1(_118_v52, 9)
	(_nw32).ArraySet1((func() _dafny.Int {
		if _108_v43 {
			return _dafny.IntOfInt64(721)
		}
		return _118_v52
	})(), 10)
	(_nw32).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(_118_v52, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _dafny.CodePoint('b'))).Cardinality()), 11)
	(_nw32).ArraySet1(_dafny.IntOfInt64(580), 12)
	(_nw32).ArraySet1((Companion_Default___.Fm9(_118_v52, globalState)).Cardinality(), 13)
	(_nw32).ArraySet1(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_118_v52, _118_v52))).Union(_209_v126)).Cardinality(), 14)
	(_nw32).ArraySet1(_118_v52, 15)
	(_nw32).ArraySet1((_118_v52).Plus(_118_v52), 16)
	(_nw32).ArraySet1((func() _dafny.Int {
		if (_204_v121).Contains(_108_v43) {
			return (_204_v121).Multiplicity(_108_v43)
		}
		return _118_v52
	})(), 17)
	(_nw32).ArraySet1(_118_v52, 18)
	(_nw32).ArraySet1(_118_v52, 19)
	(_nw32).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((p0).Cardinality()), _118_v52), 20)
	(_nw32).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm12(_118_v52, _dafny.IntOfInt64(-213), _108_v43, globalState)).Cardinality())).Times(_118_v52), 21)
	(_nw32).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 22)
	(_nw32).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_118_v52), (_211_v128).Cardinality()), 23)
	(_nw32).ArraySet1((_dafny.IntOfInt64(-200)).Minus((_dafny.Zero).Minus(_118_v52)), 24)
	r1 = _nw32
	var _212_v129 D7
	_ = _212_v129
	_212_v129 = Companion_D7_.Create_DC19_(_204_v121)
	var _213_v130 D7
	_ = _213_v130
	_213_v130 = Companion_D7_.Create_DC21_(_212_v129)
	var _pat_let_tv0 = _118_v52
	_ = _pat_let_tv0
	var _pat_let_tv1 = _209_v126
	_ = _pat_let_tv1
	r2 = (_dafny.Zero).Minus(func(_source5 D7) _dafny.Int {
		if _source5.Is_DC20() {
			var _214___mcc_h6 bool = _source5.Get_().(D7_DC20).Cf39
			_ = _214___mcc_h6
			var _215___mcc_h7 bool = _source5.Get_().(D7_DC20).Cf40
			_ = _215___mcc_h7
			var _216___mcc_h8 _dafny.Int = _source5.Get_().(D7_DC20).Cf41
			_ = _216___mcc_h8
			var _217___mcc_h9 bool = _source5.Get_().(D7_DC20).Cf42
			_ = _217___mcc_h9
			var _218_cf42 bool = _217___mcc_h9
			_ = _218_cf42
			var _219_cf41 _dafny.Int = _216___mcc_h8
			_ = _219_cf41
			var _220_cf40 bool = _215___mcc_h7
			_ = _220_cf40
			var _221_cf39 bool = _214___mcc_h6
			_ = _221_cf39
			return _219_cf41
		} else if _source5.Is_DC19() {
			var _222___mcc_h10 _dafny.MultiSet = _source5.Get_().(D7_DC19).Cf38
			_ = _222___mcc_h10
			var _223_cf38 _dafny.MultiSet = _222___mcc_h10
			_ = _223_cf38
			return _pat_let_tv0
		} else {
			var _224___mcc_h11 D7 = _source5.Get_().(D7_DC21).Cf43
			_ = _224___mcc_h11
			var _225_cf43 D7 = _224___mcc_h11
			_ = _225_cf43
			return (_pat_let_tv1).Cardinality()
		}
	}(_213_v130))
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _226_v0 bool
	_ = _226_v0
	_226_v0 = true
	var _227_v1 _dafny.MultiSet
	_ = _227_v1
	_227_v1 = _dafny.MultiSetOf(_226_v0, _226_v0)
	var _228_v2 _dafny.Sequence
	_ = _228_v2
	_228_v2 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_226_v0)), _227_v1)
	var _229_v3 _dafny.Sequence
	_ = _229_v3
	_229_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ujuh")
	var _230_v4 _dafny.Sequence
	_ = _230_v4
	_230_v4 = _dafny.SeqOf(_226_v0, _226_v0, _226_v0)
	var _231_v5 _dafny.Map
	_ = _231_v5
	_231_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_230_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_229_v3).Cardinality()), _dafny.IntOfUint32((_230_v4).Cardinality()))).Uint32()).(bool), _230_v4)
	var _232_v6 _dafny.Array
	_ = _232_v6
	var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
	_ = _nw33
	_232_v6 = _nw33
	var _233_v7 _dafny.Int
	_ = _233_v7
	_233_v7 = _dafny.IntOfInt64(773)
	var _234_v8 _dafny.Set
	_ = _234_v8
	_234_v8 = _dafny.SetOf(_226_v0)
	var _235_v9 _dafny.Array
	_ = _235_v9
	var _nwElement0_10 _dafny.Int = _233_v7
	_ = _nwElement0_10
	var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(20))
	_ = _nw34
	(_nw34).ArraySet1(_nwElement0_10, 0)
	(_nw34).ArraySet1(_dafny.IntOfInt64(-741), 1)
	(_nw34).ArraySet1(_233_v7, 2)
	(_nw34).ArraySet1(_233_v7, 3)
	(_nw34).ArraySet1(_dafny.IntOfInt64(-617), 4)
	(_nw34).ArraySet1(_233_v7, 5)
	(_nw34).ArraySet1((_234_v8).Cardinality(), 6)
	(_nw34).ArraySet1(_233_v7, 7)
	(_nw34).ArraySet1(_dafny.IntOfInt64(744), 8)
	(_nw34).ArraySet1(_233_v7, 9)
	(_nw34).ArraySet1(_233_v7, 10)
	(_nw34).ArraySet1(_dafny.IntOfInt64(462), 11)
	(_nw34).ArraySet1(_233_v7, 12)
	(_nw34).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvc")).Cardinality()), 13)
	(_nw34).ArraySet1(_233_v7, 14)
	(_nw34).ArraySet1(_233_v7, 15)
	(_nw34).ArraySet1(_233_v7, 16)
	(_nw34).ArraySet1(_dafny.IntOfInt64(726), 17)
	(_nw34).ArraySet1(_233_v7, 18)
	(_nw34).ArraySet1(_233_v7, 19)
	_235_v9 = _nw34
	var _236_v10 _dafny.MultiSet
	_ = _236_v10
	_236_v10 = _dafny.MultiSetOf(_235_v9)
	var _237_v12 _dafny.Map
	_ = _237_v12
	_237_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v0, !(_226_v0))
	var _238_v13 _dafny.Sequence
	_ = _238_v13
	_238_v13 = _dafny.SeqOf(_237_v12, _237_v12)
	var _239_v14 _dafny.Array
	_ = _239_v14
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_3
	var _nw35 _dafny.Array
	_ = _nw35
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw35 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) bool = (func(_240_v0 bool) func(_dafny.Int) bool {
			return func(_241_i0 _dafny.Int) bool {
				return _240_v0
			}
		})(_226_v0)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw35 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw35).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw35).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_239_v14 = _nw35
	var _242_v15 _dafny.Map
	_ = _242_v15
	_242_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), _226_v0)
	var _243_v16 _dafny.CodePoint
	_ = _243_v16
	_243_v16 = _dafny.CodePoint('a')
	var _244_v17 _dafny.Set
	_ = _244_v17
	_244_v17 = _dafny.SetOf(_243_v16, _243_v16)
	var _245_v18 _dafny.Map
	_ = _245_v18
	_245_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v15, _244_v17)
	var _246_globalState *GlobalState
	_ = _246_globalState
	var _nw36 *GlobalState = New_GlobalState_()
	_ = _nw36
	_nw36.Ctor__(_228_v2, _dafny.IntOfInt64(-624), _dafny.CodePoint('l'), _dafny.IntOfInt64(951), true, _dafny.IntOfInt64(868), _229_v3, _231_v5, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nwvoes"), _229_v3), _232_v6, _dafny.IntOfInt64(113), (_227_v1).Update(_226_v0, Companion_Default___.Abs(_233_v7)), _dafny.IntOfInt64(664), _dafny.IntOfInt64(310), false, _dafny.IntOfInt64(536), _dafny.IntOfInt64(495), _dafny.IntOfInt64(-530), (_236_v10).Difference(_dafny.MultiSetOf(_235_v9, _235_v9, _235_v9)), func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_238_v13).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _247_v11 _dafny.Map
			_247_v11 = interface{}(_compr_1).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_238_v13, _247_v11) {
				_coll1.Add(_247_v11, _233_v7)
			}
		}
		return _coll1.ToMap()
	}(), _239_v14, _dafny.IntOfInt64(294), true, _dafny.IntOfInt64(119), (_245_v18).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v15, _244_v17)), _dafny.IntOfInt64(-643), _227_v1, _dafny.IntOfInt64(239), _235_v9)
	_246_globalState = _nw36
	var _248_i1 _dafny.Int
	_ = _248_i1
	_248_i1 = _dafny.Zero
	{
		for !(!(_226_v0) || (_226_v0)) {
			{
				if (_248_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_248_i1 = (_248_i1).Plus(_dafny.One)
				var _249_v19 *C0
				_ = _249_v19
				var _nw37 *C0 = New_C0_()
				_ = _nw37
				_nw37.Ctor__(((_242_v15).Cardinality()).Cmp(_dafny.IntOfInt64(579)) <= 0)
				_249_v19 = _nw37
				var _250_v20 _dafny.Sequence
				_ = _250_v20
				_250_v20 = _dafny.SeqOf(_229_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-469))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}((func(_251_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_252_i2 _dafny.Int) _dafny.CodePoint {
						return _251_v16
					}
				})(_243_v16))))
				_226_v0 = Companion_Default___.Fm0((_250_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.IntOfUint32((_250_v20).Cardinality()))).Uint32()).(_dafny.Sequence), _226_v0, _246_globalState)
				_235_v9 = _235_v9
				(_246_globalState).F27 = _233_v7
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _253_v21 _dafny.Sequence
	_ = _253_v21
	var _254_v22 _dafny.Array
	_ = _254_v22
	var _255_v23 _dafny.Int
	_ = _255_v23
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 _dafny.Array
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	_out0, _out1, _out2 = Companion_Default___.M0(_229_v3, _246_globalState)
	_253_v21 = _out0
	_254_v22 = _out1
	_255_v23 = _out2
	var _hi1 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}((func(_256_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_257_i4 _dafny.Int) _dafny.CodePoint {
			return _256_v16
		}
	})(_243_v16)))).Cardinality())
	_ = _hi1
	for _258_i3 := _dafny.IntOfUint32((_253_v21).Cardinality()); _258_i3.Cmp(_hi1) < 0; _258_i3 = _258_i3.Plus(_dafny.One) {
		_237_v12 = (_237_v12).Update((_234_v8).IsProperSubsetOf(_234_v8), _226_v0)
		var _source6 D0 = Companion_Default___.Fm1(_246_globalState)
		_ = _source6
		if _source6.Is_DC0() {
			var _259___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
			_ = _259___mcc_h0
			var _260___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC0).Cf1
			_ = _260___mcc_h1
			var _261_cf1 _dafny.Int = _260___mcc_h1
			_ = _261_cf1
			var _262_cf0 _dafny.Int = _259___mcc_h0
			_ = _262_cf0
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_254_v22), 0))
			_ = _index12
			(_254_v22).ArraySet1(Companion_Default___.SafeDivisionInt(_255_v23, _255_v23), (_index12).Int())
			var _263_v24 *C0
			_ = _263_v24
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(_226_v0)
			_263_v24 = _nw38
			var _264_v25 _dafny.Array
			_ = _264_v25
			var _nwElement0_11 *C0 = (func() *C0 {
				if _226_v0 {
					return _263_v24
				}
				return _263_v24
			})()
			_ = _nwElement0_11
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.One)
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_11, 0)
			_264_v25 = _nw39
			var _265_v26 _dafny.Map
			_ = _265_v26
			_265_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v24).F29(), _233_v7)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_254_v22), 0))
			_ = _index13
			var _rhs13 _dafny.Int = (func() _dafny.Int {
				if (_226_v0) && ((_263_v24).F29()) {
					return (func() _dafny.Int {
						if (_265_v26).Contains((_263_v24).F29()) {
							return (_265_v26).Get((_263_v24).F29()).(_dafny.Int)
						}
						return _261_cf1
					})()
				}
				return _255_v23
			})()
			_ = _rhs13
			var _rhs14 _dafny.Int = _258_i3
			_ = _rhs14
			var _rhs15 _dafny.Array = _264_v25
			_ = _rhs15
			var _lhs6 _dafny.Array = _254_v22
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_254_v22), 0))
			_ = _lhs7
			(_lhs6).ArraySet1(_rhs13, (_lhs7).Int())
			_262_cf0 = _rhs14
			_264_v25 = _rhs15
			var _266_v27 *C0
			_ = _266_v27
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(_226_v0)
			_266_v27 = _nw40
			var _267_v28 _dafny.Array
			_ = _267_v28
			var _nwElement0_12 _dafny.MultiSet = _236_v10
			_ = _nwElement0_12
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(18))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_12, 0)
			(_nw41).ArraySet1((func() _dafny.MultiSet {
				if false {
					return _236_v10
				}
				return _236_v10
			})(), 1)
			(_nw41).ArraySet1(_dafny.MultiSetOf(_235_v9, _235_v9, _235_v9), 2)
			(_nw41).ArraySet1(_dafny.MultiSetOf(_235_v9, _235_v9), 3)
			(_nw41).ArraySet1(_236_v10, 4)
			(_nw41).ArraySet1(_236_v10, 5)
			(_nw41).ArraySet1(_236_v10, 6)
			(_nw41).ArraySet1(_236_v10, 7)
			(_nw41).ArraySet1(_236_v10, 8)
			(_nw41).ArraySet1(_236_v10, 9)
			(_nw41).ArraySet1(_236_v10, 10)
			(_nw41).ArraySet1((_236_v10).Update(_235_v9, Companion_Default___.Abs(_dafny.IntOfInt64(437))), 11)
			(_nw41).ArraySet1(_236_v10, 12)
			(_nw41).ArraySet1((_236_v10).Difference(_236_v10), 13)
			(_nw41).ArraySet1(_236_v10, 14)
			(_nw41).ArraySet1(_236_v10, 15)
			(_nw41).ArraySet1(_236_v10, 16)
			(_nw41).ArraySet1((_236_v10).Union(_dafny.MultiSetOf(_235_v9)), 17)
			_267_v28 = _nw41
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_267_v28), 0))
			_ = _index14
			(_267_v28).ArraySet1(_dafny.MultiSetOf(_235_v9, _235_v9, _235_v9), (_index14).Int())
			var _268_v29 _dafny.Array
			_ = _268_v29
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(17))
			_ = _nw42
			_268_v29 = _nw42
			var _269_v30 D0
			_ = _269_v30
			_269_v30 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(781), (_dafny.Zero).Minus(_261_cf1))
			var _pat_let_tv2 = _261_cf1
			_ = _pat_let_tv2
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_268_v29), 0))
			_ = _index15
			(_268_v29).ArraySet1(func(_pat_let2_0 D0) D0 {
				return func(_270_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let3_0 _dafny.Int) D0 {
						return func(_271_dt__update_hcf0_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_271_dt__update_hcf0_h0, (_270_dt__update__tmp_h0).Dtor_cf1())
						}(_pat_let3_0)
					}(_pat_let_tv2)
				}(_pat_let2_0)
			}(_269_v30), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_267_v28), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_268_v29), 0))
			_ = _index17
			var _rhs16 _dafny.MultiSet = _dafny.MultiSetOf(_235_v9, _235_v9, _235_v9, _235_v9, _235_v9)
			_ = _rhs16
			var _rhs17 D0 = Companion_D0_.Create_DC0_(_262_cf0, _262_cf0)
			_ = _rhs17
			var _rhs18 bool = (_266_v27).F29()
			_ = _rhs18
			var _rhs19 *C0 = _263_v24
			_ = _rhs19
			var _rhs20 _dafny.Int = ((_dafny.Zero).Minus(_261_cf1)).Plus((func() _dafny.Int {
				if (_266_v27).F29() {
					return _261_cf1
				}
				return _262_cf0
			})())
			_ = _rhs20
			var _lhs8 _dafny.Array = _267_v28
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_267_v28), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _268_v29
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_268_v29), 0))
			_ = _lhs11
			var _lhs12 *GlobalState = _246_globalState
			_ = _lhs12
			var _lhs13 *GlobalState = _246_globalState
			_ = _lhs13
			(_lhs8).ArraySet1(_rhs16, (_lhs9).Int())
			(_lhs10).ArraySet1(_rhs17, (_lhs11).Int())
			_lhs12.F14 = _rhs18
			_263_v24 = _rhs19
			_lhs13.F13 = _rhs20
			var _272_v31 _dafny.Map
			_ = _272_v31
			_272_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _266_v27)
			_272_v31 = (_272_v31).Update(!((_266_v27).F29()), _263_v24)
		} else {
			(_246_globalState).F27 = (_233_v7).Plus(_dafny.IntOfInt64(316))
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_239_v14), 0))
			_ = _index18
			(_239_v14).ArraySet1((_226_v0) == (_226_v0), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_239_v14), 0))
			_ = _index19
			(_239_v14).ArraySet1(!(_226_v0), (_index19).Int())
			_255_v23 = _255_v23
			var _273_v32 _dafny.Array
			_ = _273_v32
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_4
			var _nw43 _dafny.Array
			_ = _nw43
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw43 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Map = (func(_274_v0 bool, _275_i3 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_276_i5 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_274_v0, _dafny.IntOfInt64(889))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_274_v0, _275_i3))
					}
				})(_226_v0, _258_i3)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw43 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw43).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw43).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_273_v32 = _nw43
			var _277_v33 _dafny.Map
			_ = _277_v33
			_277_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool), _dafny.IntOfInt64(878))
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_273_v32), 0))
			_ = _index20
			(_273_v32).ArraySet1((func() _dafny.Map {
				if (func() bool {
					if (_237_v12).Contains((_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool)) {
						return (_237_v12).Get((_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool)).(bool)
					}
					return _226_v0
				})() {
					return _277_v33
				}
				return _277_v33
			})(), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_273_v32), 0))
			_ = _index21
			(_273_v32).ArraySet1(_277_v33, (_index21).Int())
		}
		var _278_v34 _dafny.Array
		_ = _278_v34
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
		_ = _nw44
		_278_v34 = _nw44
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_278_v34), 0))
		_ = _index22
		(_278_v34).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_230_v4, _230_v4), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_278_v34), 0))
		_ = _index23
		(_278_v34).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_230_v4, _230_v4), (_index23).Int())
		var _279_v35 *C0
		_ = _279_v35
		var _nw45 *C0 = New_C0_()
		_ = _nw45
		_nw45.Ctor__(_226_v0)
		_279_v35 = _nw45
	}
	var _280_v36 _dafny.Array
	_ = _280_v36
	var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(21))
	_ = _nw46
	_280_v36 = _nw46
	var _281_v37 _dafny.MultiSet
	_ = _281_v37
	_281_v37 = _dafny.MultiSetOf(_239_v14, _239_v14, _239_v14, _239_v14, _239_v14)
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_280_v36), 0))
	_ = _index24
	(_280_v36).ArraySet1(_281_v37, (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_280_v36), 0))
	_ = _index25
	(_280_v36).ArraySet1(((_281_v37).Difference(_281_v37)).Intersection((_281_v37).Intersection(_281_v37)), (_index25).Int())
	if (_255_v23).Cmp(_255_v23) <= 0 {
		_243_v16 = _243_v16
		var _282_v38 _dafny.Sequence
		_ = _282_v38
		var _283_v39 _dafny.Array
		_ = _283_v39
		var _284_v40 _dafny.Int
		_ = _284_v40
		var _out3 _dafny.Sequence
		_ = _out3
		var _out4 _dafny.Array
		_ = _out4
		var _out5 _dafny.Int
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.M0(_229_v3, _246_globalState)
		_282_v38 = _out3
		_283_v39 = _out4
		_284_v40 = _out5
		var _285_v41 _dafny.Sequence
		_ = _285_v41
		var _286_v42 _dafny.Array
		_ = _286_v42
		var _287_v43 _dafny.Int
		_ = _287_v43
		var _out6 _dafny.Sequence
		_ = _out6
		var _out7 _dafny.Array
		_ = _out7
		var _out8 _dafny.Int
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.M0(_dafny.UnicodeSeqOfUtf8Bytes("konvcdp"), _246_globalState)
		_285_v41 = _out6
		_286_v42 = _out7
		_287_v43 = _out8
		if _226_v0 {
			(_246_globalState).F13 = _233_v7
			var _288_v44 *C0
			_ = _288_v44
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(true)
			_288_v44 = _nw47
			var _289_v45 _dafny.Map
			_ = _289_v45
			_289_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(309), _284_v40)
			var _290_v46 *C0
			_ = _290_v46
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__((_289_v45).Equals(_289_v45))
			_290_v46 = _nw48
			_229_v3 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			(_246_globalState).F14 = (func() bool {
				if (_242_v15).Contains(_287_v43) {
					return (_242_v15).Get(_287_v43).(bool)
				}
				return ((_dafny.Zero).Minus((_dafny.Zero).Minus(_284_v40))).Cmp(_284_v40) <= 0
			})()
		} else {
			(_246_globalState).F27 = _dafny.IntOfInt64(157)
			var _291_v47 _dafny.Map
			_ = _291_v47
			_291_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v16, _239_v14)
			var _rhs21 _dafny.Map = ((_291_v47).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v16, _239_v14))).Update(_243_v16, _239_v14)
			_ = _rhs21
			var _rhs22 bool = Companion_Default___.Fm0(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-390))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_292_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_293_i6 _dafny.Int) _dafny.CodePoint {
					return _292_v16
				}
			})(_243_v16))), !(_242_v15).Equals(_242_v15), _246_globalState)
			_ = _rhs22
			var _rhs23 _dafny.MultiSet = _227_v1
			_ = _rhs23
			var _lhs14 *GlobalState = _246_globalState
			_ = _lhs14
			_291_v47 = _rhs21
			_lhs14.F14 = _rhs22
			_227_v1 = _rhs23
			var _294_v48 _dafny.Sequence
			_ = _294_v48
			_294_v48 = _dafny.SeqOf(_255_v23)
			var _295_v49 _dafny.Map
			_ = _295_v49
			_295_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_284_v40, _dafny.IntOfInt64(663))
			var _296_v50 _dafny.Sequence
			_ = _296_v50
			_296_v50 = _dafny.SeqOf(_dafny.IntOfInt64(745), (_294_v48).Select((Companion_Default___.SafeIndex((_295_v49).Cardinality(), _dafny.IntOfUint32((_294_v48).Cardinality()))).Uint32()).(_dafny.Int))
			(_246_globalState).F13 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(57), (_287_v43).Plus(_dafny.IntOfUint32((_296_v50).Cardinality())))
			var _297_v51 *C0
			_ = _297_v51
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_294_v48, _296_v50))
			_297_v51 = _nw49
			var _298_v52 _dafny.MultiSet
			_ = _298_v52
			_298_v52 = _dafny.MultiSetOf(_243_v16, _243_v16)
			var _299_v53 _dafny.Set
			_ = _299_v53
			_299_v53 = _dafny.SetOf(Companion_Default___.Fm2(_298_v52, _255_v23, _233_v7, _226_v0, _246_globalState))
			(_246_globalState).F25 = (((_242_v15).Update((_299_v53).Cardinality(), (_297_v51).F29())).Update(_dafny.IntOfInt64(886), _226_v0)).Cardinality()
		}
		var _300_v54 *C0
		_ = _300_v54
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__(_226_v0)
		_300_v54 = _nw50
	} else {
		(_246_globalState).F2 = _243_v16
		(_246_globalState).F25 = _255_v23
		_226_v0 = (_233_v7).Cmp(_255_v23) >= 0
		(_246_globalState).F27 = (_dafny.Zero).Minus(_233_v7)
		_237_v12 = (_237_v12).Update(_226_v0, (func() bool {
			if Companion_Default___.Fm0(_229_v3, _226_v0, _246_globalState) {
				return _226_v0
			}
			return (func() bool {
				if (_242_v15).Contains(_255_v23) {
					return (_242_v15).Get(_255_v23).(bool)
				}
				return _226_v0
			})()
		})())
	}
	var _301_v55 *C0
	_ = _301_v55
	var _nw51 *C0 = New_C0_()
	_ = _nw51
	_nw51.Ctor__(_dafny.Companion_Sequence_.Contains(_253_v21, _243_v16))
	_301_v55 = _nw51
	(_246_globalState).F27 = _255_v23
	if (_301_v55).F29() {
		(_246_globalState).F13 = _255_v23
		(_246_globalState).F14 = (_301_v55).F29()
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_239_v14), 0))
		_ = _index26
		(_239_v14).ArraySet1(_226_v0, (_index26).Int())
		var _302_v56 D1
		_ = _302_v56
		_302_v56 = Companion_D1_.Create_DC2_(_253_v21)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_239_v14), 0))
		_ = _index27
		(_239_v14).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_229_v3, (Companion_Default___.SafeIndex(_255_v23, _dafny.IntOfUint32((_229_v3).Cardinality()))).Uint32(), _243_v16), (_302_v56).Dtor_cf2()), _229_v3), (_index27).Int())
		(_246_globalState).F14 = !((_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool))
		var _303_v57 _dafny.Set
		_ = _303_v57
		_303_v57 = _dafny.SetOf(_255_v23)
		var _304_v58 _dafny.Sequence
		_ = _304_v58
		_304_v58 = _dafny.SeqOf(_233_v7, (_dafny.Zero).Minus(_dafny.IntOfInt64(-473)), _233_v7)
		var _305_v59 _dafny.Sequence
		_ = _305_v59
		_305_v59 = _dafny.SeqOf(_303_v57, _dafny.SetOf((_dafny.Zero).Minus((_304_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.IntOfUint32((_304_v58).Cardinality()))).Uint32()).(_dafny.Int))))
		var _306_v60 _dafny.MultiSet
		_ = _306_v60
		_306_v60 = _dafny.MultiSetOf(_303_v57, (_305_v59).Select((Companion_Default___.SafeIndex(_255_v23, _dafny.IntOfUint32((_305_v59).Cardinality()))).Uint32()).(_dafny.Set))
		_306_v60 = _306_v60
	} else {
		var _307_v61 D1
		_ = _307_v61
		_307_v61 = Companion_D1_.Create_DC4_((_301_v55).F29(), _233_v7)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))
		_ = _index28
		(_235_v9).ArraySet1((_307_v61).Dtor_cf6(), (_index28).Int())
		var _308_v62 _dafny.Map
		_ = _308_v62
		_308_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v55).F29(), _237_v12)
		var _309_v63 _dafny.Map
		_ = _309_v63
		_309_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_308_v62).Cardinality(), _243_v16)
		var _310_v64 _dafny.Map
		_ = _310_v64
		_310_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_309_v63).Cardinality(), _dafny.SeqOf(_226_v0, (_301_v55).F29(), _226_v0, (_301_v55).F29(), (_301_v55).F29()))
		var _311_v65 _dafny.Map
		_ = _311_v65
		_311_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_310_v64).Merge(_310_v64), _255_v23)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))
		_ = _index29
		(_235_v9).ArraySet1((func() _dafny.Int {
			if (_311_v65).Contains(_310_v64) {
				return (_311_v65).Get(_310_v64).(_dafny.Int)
			}
			return _255_v23
		})(), (_index29).Int())
		var _312_v66 _dafny.Sequence
		_ = _312_v66
		_312_v66 = _dafny.SeqOf(_dafny.IntOfInt64(238), (_235_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))).Int()).(_dafny.Int))
		var _313_v67 _dafny.Sequence
		_ = _313_v67
		_313_v67 = _dafny.SeqOf(_255_v23, (_312_v66).Select((Companion_Default___.SafeIndex(_233_v7, _dafny.IntOfUint32((_312_v66).Cardinality()))).Uint32()).(_dafny.Int))
		var _rhs24 _dafny.Sequence = _312_v66
		_ = _rhs24
		var _rhs25 bool = (_301_v55).F29()
		_ = _rhs25
		var _lhs15 *GlobalState = _246_globalState
		_ = _lhs15
		_312_v66 = _rhs24
		_lhs15.F14 = _rhs25
		var _314_v68 D0
		_ = _314_v68
		_314_v68 = Companion_D0_.Create_DC1_()
		_314_v68 = Companion_Default___.Fm3(_233_v7, _246_globalState)
		if _226_v0 {
			_226_v0 = false
			(_246_globalState).F17 = _233_v7
			var _315_v69 _dafny.Sequence
			_ = _315_v69
			var _316_v70 _dafny.Array
			_ = _316_v70
			var _317_v71 _dafny.Int
			_ = _317_v71
			var _out9 _dafny.Sequence
			_ = _out9
			var _out10 _dafny.Array
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out9, _out10, _out11 = Companion_Default___.M0(_253_v21, _246_globalState)
			_315_v69 = _out9
			_316_v70 = _out10
			_317_v71 = _out11
			var _318_v72 _dafny.Map
			_ = _318_v72
			_318_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v55).F29(), _dafny.IntOfUint32((_dafny.SeqOf(_255_v23, _317_v71)).Cardinality()))
			var _rhs26 _dafny.Int = _317_v71
			_ = _rhs26
			var _rhs27 bool = Companion_Default___.Fm0(_253_v21, !(_226_v0), _246_globalState)
			_ = _rhs27
			var _rhs28 bool = _226_v0
			_ = _rhs28
			var _rhs29 _dafny.Int = ((_235_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))).Int()).(_dafny.Int)).Plus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_318_v72).Contains((_301_v55).F29()) {
					return (_318_v72).Get((_301_v55).F29()).(_dafny.Int)
				}
				return _255_v23
			})(), (_235_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))).Int()).(_dafny.Int)))
			_ = _rhs29
			var _lhs16 *GlobalState = _246_globalState
			_ = _lhs16
			var _lhs17 *GlobalState = _246_globalState
			_ = _lhs17
			var _lhs18 *GlobalState = _246_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _246_globalState
			_ = _lhs19
			_lhs16.F25 = _rhs26
			_lhs17.F14 = _rhs27
			_lhs18.F14 = _rhs28
			_lhs19.F25 = _rhs29
			(_246_globalState).F14 = ((func() bool {
				if !(Companion_Default___.Fm0(_253_v21, (_301_v55).F29(), _246_globalState)) {
					return (_301_v55).F29()
				}
				return (_301_v55).F29()
			})()) && ((_301_v55).F29())
		} else {
			var _319_v73 D1
			_ = _319_v73
			_319_v73 = Companion_D1_.Create_DC2_(_253_v21)
			_319_v73 = Companion_Default___.Fm4((_dafny.IntOfInt64(509)).Cmp(_233_v7) >= 0, (_301_v55).F29(), _246_globalState)
			var _320_v74 _dafny.Set
			_ = _320_v74
			_320_v74 = _dafny.SetOf(_253_v21)
			var _321_v75 _dafny.Sequence
			_ = _321_v75
			_321_v75 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-172))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_322_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(57))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_323_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})))
			var _rhs30 bool = (_320_v74).IsProperSubsetOf((_320_v74).Difference(func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_321_v75).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _324_v76 _dafny.Sequence
					_324_v76 = interface{}(_compr_2).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_321_v75, _324_v76) {
						_coll2.Add(_324_v76)
					}
				}
				return _coll2.ToSet()
			}()))
			_ = _rhs30
			var _rhs31 bool = (_230_v4).Select((Companion_Default___.SafeIndex((_235_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_230_v4).Cardinality()))).Uint32()).(bool)
			_ = _rhs31
			var _rhs32 D0 = _314_v68
			_ = _rhs32
			_226_v0 = _rhs30
			_226_v0 = _rhs31
			_314_v68 = _rhs32
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))
			_ = _index30
			(_235_v9).ArraySet1((_dafny.Zero).Minus((_255_v23).Minus(((_dafny.Zero).Minus(_255_v23)).Minus((_235_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_235_v9), 0))).Int()).(_dafny.Int)))), (_index30).Int())
			(_246_globalState).F17 = (_dafny.Zero).Minus((_313_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_312_v66).Cardinality()), _dafny.IntOfUint32((_313_v67).Cardinality()))).Uint32()).(_dafny.Int))
			(_246_globalState).F14 = false
		}
		var _325_v77 _dafny.Sequence
		_ = _325_v77
		var _326_v78 _dafny.Array
		_ = _326_v78
		var _327_v79 _dafny.Int
		_ = _327_v79
		var _out12 _dafny.Sequence
		_ = _out12
		var _out13 _dafny.Array
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		_out12, _out13, _out14 = Companion_Default___.M0(_dafny.UnicodeSeqOfUtf8Bytes("arg"), _246_globalState)
		_325_v77 = _out12
		_326_v78 = _out13
		_327_v79 = _out14
	}
	if (_233_v7).Cmp(_255_v23) <= 0 {
		var _328_v80 _dafny.Sequence
		_ = _328_v80
		var _329_v81 _dafny.Array
		_ = _329_v81
		var _330_v82 _dafny.Int
		_ = _330_v82
		var _out15 _dafny.Sequence
		_ = _out15
		var _out16 _dafny.Array
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		_out15, _out16, _out17 = Companion_Default___.M0(_229_v3, _246_globalState)
		_328_v80 = _out15
		_329_v81 = _out16
		_330_v82 = _out17
		var _331_v83 _dafny.Sequence
		_ = _331_v83
		var _332_v84 _dafny.Array
		_ = _332_v84
		var _333_v85 _dafny.Int
		_ = _333_v85
		var _out18 _dafny.Sequence
		_ = _out18
		var _out19 _dafny.Array
		_ = _out19
		var _out20 _dafny.Int
		_ = _out20
		_out18, _out19, _out20 = Companion_Default___.M0(_229_v3, _246_globalState)
		_331_v83 = _out18
		_332_v84 = _out19
		_333_v85 = _out20
		(_246_globalState).F14 = (_301_v55).F29()
		var _334_v86 _dafny.Array
		_ = _334_v86
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
		_ = _nw52
		_334_v86 = _nw52
		_334_v86 = _334_v86
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_239_v14), 0))
		_ = _index31
		(_239_v14).ArraySet1((_301_v55).F29(), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_239_v14), 0))
		_ = _index32
		var _rhs33 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_330_v82))
		_ = _rhs33
		var _rhs34 bool = _226_v0
		_ = _rhs34
		var _lhs20 *GlobalState = _246_globalState
		_ = _lhs20
		var _lhs21 _dafny.Array = _239_v14
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_239_v14), 0))
		_ = _lhs22
		_lhs20.F17 = _rhs33
		(_lhs21).ArraySet1(_rhs34, (_lhs22).Int())
	} else {
		var _335_v87 _dafny.MultiSet
		_ = _335_v87
		_335_v87 = _dafny.MultiSetOf(_243_v16, _243_v16, _243_v16, _243_v16, _243_v16)
		var _336_v88 _dafny.Set
		_ = _336_v88
		_336_v88 = _dafny.SetOf(_255_v23, Companion_Default___.Fm2(_335_v87, _233_v7, _dafny.IntOfInt64(331), true, _246_globalState), (func() _dafny.Int {
			if _226_v0 {
				return (_242_v15).Cardinality()
			}
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-400))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_337_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))).Cardinality())
		})(), _255_v23)
		_336_v88 = ((_336_v88).Difference(_336_v88)).Intersection(func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((Companion_Default___.Fm5(_226_v0, _255_v23, (_301_v55).F29(), _246_globalState)).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _338_v89 _dafny.Int
				_338_v89 = interface{}(_compr_3).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm5(_226_v0, _255_v23, (_301_v55).F29(), _246_globalState), _338_v89) {
					_coll3.Add((_338_v89).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khn")).Cardinality())))
				}
			}
			return _coll3.ToSet()
		}())
		var _339_v90 *C0
		_ = _339_v90
		var _nw53 *C0 = New_C0_()
		_ = _nw53
		_nw53.Ctor__((_301_v55).F29())
		_339_v90 = _nw53
		if (_301_v55).F29() {
			var _340_v91 _dafny.Sequence
			_ = _340_v91
			_340_v91 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_253_v21, _229_v3), _229_v3, _dafny.UnicodeSeqOfUtf8Bytes("dhyao"))
			_340_v91 = Companion_Default___.Fm6(_253_v21, Companion_D1_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("yuhsgwf")), (_dafny.Zero).Minus(_255_v23), (_233_v7).Cmp(_255_v23) >= 0, _246_globalState)
			(_246_globalState).F17 = _255_v23
			var _341_v92 _dafny.Map
			_ = _341_v92
			_341_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_233_v7)), _339_v90)
			(_246_globalState).F14 = !((_341_v92).Merge((_341_v92).Update((_dafny.Zero).Minus(_255_v23), _339_v90))).Contains(_dafny.IntOfUint32((_dafny.SeqOf((_301_v55).F29())).Cardinality()))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_235_v9), 0))
			_ = _index33
			(_235_v9).ArraySet1((_dafny.Zero).Minus(_255_v23), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_235_v9), 0))
			_ = _index34
			(_235_v9).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_255_v23, _233_v7), _255_v23), (_index34).Int())
			var _342_v93 _dafny.Array
			_ = _342_v93
			var _nwElement0_13 _dafny.MultiSet = _335_v87
			_ = _nwElement0_13
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(3))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_13, 0)
			(_nw54).ArraySet1(_335_v87, 1)
			(_nw54).ArraySet1(_335_v87, 2)
			_342_v93 = _nw54
			var _343_v94 _dafny.Set
			_ = _343_v94
			_343_v94 = _dafny.SetOf(_336_v88)
			var _344_v95 D2
			_ = _344_v95
			_344_v95 = Companion_D2_.Create_DC7_(_343_v94)
			var _345_v96 _dafny.Map
			_ = _345_v96
			_345_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v93, ((_344_v95).Dtor_cf12()).IsDisjointFrom(_343_v94))
			(_246_globalState).F14 = (func() bool {
				if (_345_v96).Contains(_342_v93) {
					return (_345_v96).Get(_342_v93).(bool)
				}
				return (_301_v55).F29()
			})()
		} else {
			var _346_v97 _dafny.Map
			_ = _346_v97
			_346_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v90, _339_v90)
			_339_v90 = (func() *C0 {
				if (_346_v97).Contains(_301_v55) {
					return (_346_v97).Get(_301_v55).(*C0)
				}
				return _339_v90
			})()
			var _347_v98 *C0
			_ = _347_v98
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__((_301_v55).F29())
			_347_v98 = _nw55
			var _348_v99 _dafny.Sequence
			_ = _348_v99
			var _349_v100 _dafny.Array
			_ = _349_v100
			var _350_v101 _dafny.Int
			_ = _350_v101
			var _out21 _dafny.Sequence
			_ = _out21
			var _out22 _dafny.Array
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			_out21, _out22, _out23 = Companion_Default___.M0(_dafny.UnicodeSeqOfUtf8Bytes("dkmvx"), _246_globalState)
			_348_v99 = _out21
			_349_v100 = _out22
			_350_v101 = _out23
			(_246_globalState).F14 = (_230_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_233_v7), _dafny.IntOfUint32((_230_v4).Cardinality()))).Uint32()).(bool)
			var _351_v102 _dafny.Array
			_ = _351_v102
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_5
			var _nw56 _dafny.Array
			_ = _nw56
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw56 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_352_v3 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_353_i10 _dafny.Int) _dafny.CodePoint {
						return (_352_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.IntOfUint32((_352_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_229_v3)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw56 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw56).ArraySet1CodePoint(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw56).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_351_v102 = _nw56
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_351_v102), 0))
			_ = _index35
			(_351_v102).ArraySet1CodePoint(_243_v16, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_351_v102), 0))
			_ = _index36
			var _rhs35 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_230_v4, _230_v4)
			_ = _rhs35
			var _rhs36 _dafny.CodePoint = Companion_Default___.Fm7((func() _dafny.Sequence {
				if (_301_v55).F29() {
					return _229_v3
				}
				return _253_v21
			})(), (_dafny.Zero).Minus(_233_v7), _246_globalState)
			_ = _rhs36
			var _lhs23 _dafny.Array = _351_v102
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_351_v102), 0))
			_ = _lhs24
			_230_v4 = _rhs35
			(_lhs23).ArraySet1CodePoint(_rhs36, (_lhs24).Int())
		}
		(_246_globalState).F13 = ((_dafny.Zero).Minus((_255_v23).Times(_dafny.IntOfInt64(-758)))).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_354_v7 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_355_i11 _dafny.Int) _dafny.Int {
				return _354_v7
			}
		})(_233_v7)))).Cardinality()))
		var _356_v103 _dafny.MultiSet
		_ = _356_v103
		_356_v103 = _dafny.MultiSetOf(_233_v7, _dafny.IntOfInt64(-464), _233_v7, _255_v23)
		if (_356_v103).IsProperSubsetOf(_356_v103) {
			_233_v7 = _233_v7
			(_246_globalState).F14 = (_301_v55).F29()
			(_246_globalState).F17 = _255_v23
			var _357_v104 _dafny.Sequence
			_ = _357_v104
			_357_v104 = _dafny.SeqOf(_233_v7)
			var _358_v105 _dafny.Map
			_ = _358_v105
			_358_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_357_v104, (_301_v55).F29())
			var _359_v106 _dafny.Sequence
			_ = _359_v106
			_359_v106 = _dafny.SeqOf(Companion_Default___.Fm2(_dafny.MultiSetFromSeq(_229_v3), (_358_v105).Cardinality(), _255_v23, _226_v0, _246_globalState))
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_235_v9), 0))
			_ = _index37
			(_235_v9).ArraySet1(Companion_Default___.SafeModuloInt(_255_v23, _dafny.IntOfUint32((_359_v106).Cardinality())), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_235_v9), 0))
			_ = _index38
			(_235_v9).ArraySet1(_255_v23, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_235_v9), 0))
			_ = _index39
			(_235_v9).ArraySet1(_dafny.IntOfUint32((_359_v106).Cardinality()), (_index39).Int())
		} else {
			var _360_v107 _dafny.Map
			_ = _360_v107
			_360_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC10_(_233_v7, _243_v16), _dafny.MultiSetOf((_339_v90).F29(), _226_v0, (_339_v90).F29(), (_301_v55).F29(), (_301_v55).F29()))
			var _361_v108 D2
			_ = _361_v108
			_361_v108 = Companion_D2_.Create_DC10_(_255_v23, _243_v16)
			_226_v0 = !((func() _dafny.MultiSet {
				if (_360_v107).Contains(_361_v108) {
					return (_360_v107).Get(_361_v108).(_dafny.MultiSet)
				}
				return _227_v1
			})()).Contains((_301_v55).F29())
			_253_v21 = _229_v3
			var _362_v109 _dafny.Array
			_ = _362_v109
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
			_ = _nw57
			_362_v109 = _nw57
			_362_v109 = _362_v109
			var _363_v110 _dafny.Array
			_ = _363_v110
			var _nwElement0_14 _dafny.Array = _235_v9
			_ = _nwElement0_14
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.One)
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_14, 0)
			_363_v110 = _nw58
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_363_v110), 0))
			_ = _index40
			(_363_v110).ArraySet1(_254_v22, (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_239_v14), 0))
			_ = _index41
			(_239_v14).ArraySet1(_226_v0, (_index41).Int())
			var _364_v111 _dafny.Map
			_ = _364_v111
			_364_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v55).F29(), (_227_v1).Cardinality())
			var _365_v112 _dafny.Map
			_ = _365_v112
			_365_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v7, _364_v111)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_363_v110), 0))
			_ = _index42
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_239_v14), 0))
			_ = _index43
			var _rhs37 _dafny.Array = _235_v9
			_ = _rhs37
			var _rhs38 bool = Companion_Default___.Fm0(_229_v3, ((func() _dafny.Map {
				if (_365_v112).Contains(_255_v23) {
					return (_365_v112).Get(_255_v23).(_dafny.Map)
				}
				return _364_v111
			})()).Contains((_339_v90).F29()), _246_globalState)
			_ = _rhs38
			var _rhs39 bool = (_301_v55).F29()
			_ = _rhs39
			var _lhs25 _dafny.Array = _363_v110
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_363_v110), 0))
			_ = _lhs26
			var _lhs27 _dafny.Array = _239_v14
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_239_v14), 0))
			_ = _lhs28
			var _lhs29 *GlobalState = _246_globalState
			_ = _lhs29
			(_lhs25).ArraySet1(_rhs37, (_lhs26).Int())
			(_lhs27).ArraySet1(_rhs38, (_lhs28).Int())
			_lhs29.F14 = _rhs39
			var _366_v113 _dafny.Sequence
			_ = _366_v113
			var _367_v114 _dafny.Array
			_ = _367_v114
			var _368_v115 _dafny.Int
			_ = _368_v115
			var _out24 _dafny.Sequence
			_ = _out24
			var _out25 _dafny.Array
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			_out24, _out25, _out26 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_229_v3, _253_v21), _246_globalState)
			_366_v113 = _out24
			_367_v114 = _out25
			_368_v115 = _out26
		}
	}
	if !((_255_v23).Cmp(_255_v23) <= 0) {
		var _369_v116 *C0
		_ = _369_v116
		var _nw59 *C0 = New_C0_()
		_ = _nw59
		_nw59.Ctor__(true)
		_369_v116 = _nw59
		_226_v0 = _226_v0
		(_246_globalState).F14 = (_301_v55).F29()
		var _370_v117 _dafny.Set
		_ = _370_v117
		_370_v117 = _dafny.SetOf(_233_v7, _233_v7)
		var _371_v118 _dafny.Set
		_ = _371_v118
		_371_v118 = _dafny.SetOf(_370_v117)
		var _372_v119 D2
		_ = _372_v119
		_372_v119 = Companion_D2_.Create_DC7_(_371_v118)
		var _pat_let_tv3 = _371_v118
		_ = _pat_let_tv3
		var _pat_let_tv4 = _371_v118
		_ = _pat_let_tv4
		_372_v119 = func(_pat_let4_0 D2) D2 {
			return func(_373_dt__update__tmp_h1 D2) D2 {
				return func(_pat_let5_0 _dafny.Set) D2 {
					return func(_374_dt__update_hcf12_h0 _dafny.Set) D2 {
						return Companion_D2_.Create_DC7_(_374_dt__update_hcf12_h0)
					}(_pat_let5_0)
				}((_pat_let_tv3).Difference(_pat_let_tv4))
			}(_pat_let4_0)
		}(Companion_D2_.Create_DC7_(_371_v118))
		var _375_v120 _dafny.Map
		_ = _375_v120
		_375_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v16, (_369_v116).F29())
		var _rhs40 _dafny.Int = _255_v23
		_ = _rhs40
		var _rhs41 _dafny.Map = _375_v120
		_ = _rhs41
		var _rhs42 _dafny.Int = _233_v7
		_ = _rhs42
		var _lhs30 *GlobalState = _246_globalState
		_ = _lhs30
		var _lhs31 *GlobalState = _246_globalState
		_ = _lhs31
		_lhs30.F17 = _rhs40
		_375_v120 = _rhs41
		_lhs31.F27 = _rhs42
	} else {
		var _376_v121 _dafny.Sequence
		_ = _376_v121
		var _377_v122 _dafny.Array
		_ = _377_v122
		var _378_v123 _dafny.Int
		_ = _378_v123
		var _out27 _dafny.Sequence
		_ = _out27
		var _out28 _dafny.Array
		_ = _out28
		var _out29 _dafny.Int
		_ = _out29
		_out27, _out28, _out29 = Companion_Default___.M0(_dafny.UnicodeSeqOfUtf8Bytes("yqs"), _246_globalState)
		_376_v121 = _out27
		_377_v122 = _out28
		_378_v123 = _out29
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_239_v14), 0))
		_ = _index44
		(_239_v14).ArraySet1((_301_v55).F29(), (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_239_v14), 0))
		_ = _index45
		(_239_v14).ArraySet1((_301_v55).F29(), (_index45).Int())
		var _379_v124 *C0
		_ = _379_v124
		var _nw60 *C0 = New_C0_()
		_ = _nw60
		_nw60.Ctor__(_226_v0)
		_379_v124 = _nw60
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_239_v14), 0))
		_ = _index46
		(_239_v14).ArraySet1(_226_v0, (_index46).Int())
		var _380_v125 _dafny.Map
		_ = _380_v125
		_380_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v0, _376_v121)
		var _381_v126 _dafny.MultiSet
		_ = _381_v126
		_381_v126 = _dafny.MultiSetOf(_255_v23, _233_v7, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_380_v125).Contains((_301_v55).F29()) {
				return (_380_v125).Get((_301_v55).F29()).(_dafny.Sequence)
			}
			return _253_v21
		})()).Cardinality()), Companion_Default___.SafeModuloInt(_233_v7, _378_v123))
		var _382_v127 _dafny.Map
		_ = _382_v127
		_382_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v23, _dafny.MultiSetOf(_378_v123, _255_v23, _233_v7, _233_v7, _dafny.IntOfInt64(101)))
		_381_v126 = (func() _dafny.MultiSet {
			if !((_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool)) {
				return _381_v126
			}
			return (func() _dafny.MultiSet {
				if (_379_v124).F29() {
					return _381_v126
				}
				return (func() _dafny.MultiSet {
					if (_382_v127).Contains(_378_v123) {
						return (_382_v127).Get(_378_v123).(_dafny.MultiSet)
					}
					return _381_v126
				})()
			})()
		})()
	}
	var _383_v128 _dafny.Sequence
	_ = _383_v128
	var _384_v129 _dafny.Array
	_ = _384_v129
	var _385_v130 _dafny.Int
	_ = _385_v130
	var _out30 _dafny.Sequence
	_ = _out30
	var _out31 _dafny.Array
	_ = _out31
	var _out32 _dafny.Int
	_ = _out32
	_out30, _out31, _out32 = Companion_Default___.M0(_253_v21, _246_globalState)
	_383_v128 = _out30
	_384_v129 = _out31
	_385_v130 = _out32
	_301_v55 = (func() *C0 {
		if false {
			return _301_v55
		}
		return _301_v55
	})()
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_239_v14), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _386_i12 _dafny.Int
		_386_i12 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_386_i12).Sign() != -1) && ((_386_i12).Cmp(_dafny.ArrayLen((_239_v14), 0)) < 0)) {
			(_239_v14).ArraySet1((_301_v55).F29(), (_386_i12).Int())
		}
	}
	if _226_v0 {
		var _387_v131 _dafny.Array
		_ = _387_v131
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw61
		_387_v131 = _nw61
		_387_v131 = (func() _dafny.Array {
			if _226_v0 {
				return _387_v131
			}
			return _387_v131
		})()
		(_246_globalState).F13 = Companion_Default___.SafeDivisionInt(_385_v130, Companion_Default___.SafeModuloInt(_255_v23, _233_v7))
		(_246_globalState).F27 = _385_v130
		var _388_v132 _dafny.Sequence
		_ = _388_v132
		var _389_v133 _dafny.Array
		_ = _389_v133
		var _390_v134 _dafny.Int
		_ = _390_v134
		var _out33 _dafny.Sequence
		_ = _out33
		var _out34 _dafny.Array
		_ = _out34
		var _out35 _dafny.Int
		_ = _out35
		_out33, _out34, _out35 = Companion_Default___.M0(_253_v21, _246_globalState)
		_388_v132 = _out33
		_389_v133 = _out34
		_390_v134 = _out35
		var _391_v135 _dafny.Map
		_ = _391_v135
		_391_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_385_v130, (_301_v55).F29())).Cardinality(), (_dafny.Zero).Minus(_233_v7))
		var _392_v136 _dafny.Set
		_ = _392_v136
		_392_v136 = _dafny.SetOf(_233_v7, (_dafny.Zero).Minus((_391_v135).Cardinality()))
		var _393_v137 _dafny.Map
		_ = _393_v137
		_393_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v23, (_392_v136).Cardinality())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_235_v9), 0))
		_ = _index47
		(_235_v9).ArraySet1((func() _dafny.Int {
			if (_393_v137).Contains(_255_v23) {
				return (_393_v137).Get(_255_v23).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-935)
		})(), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_235_v9), 0))
		_ = _index48
		var _rhs43 bool = _226_v0
		_ = _rhs43
		var _rhs44 _dafny.Int = _dafny.IntOfInt64(-171)
		_ = _rhs44
		var _rhs45 _dafny.Int = Companion_Default___.SafeDivisionInt(_255_v23, _390_v134)
		_ = _rhs45
		var _lhs32 _dafny.Array = _235_v9
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_235_v9), 0))
		_ = _lhs33
		_226_v0 = _rhs43
		_385_v130 = _rhs44
		(_lhs32).ArraySet1(_rhs45, (_lhs33).Int())
	} else {
		var _394_v138 _dafny.Map
		_ = _394_v138
		_394_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v16, _226_v0)
		var _395_v139 _dafny.Sequence
		_ = _395_v139
		_395_v139 = _dafny.SeqOf(_233_v7, (_dafny.SetOf(_dafny.IntOfInt64(395), _233_v7)).Cardinality(), (_394_v138).Cardinality(), (_237_v12).Cardinality(), _255_v23)
		var _396_v140 _dafny.Set
		_ = _396_v140
		_396_v140 = _dafny.SetOf(_395_v139)
		var _397_v141 _dafny.MultiSet
		_ = _397_v141
		_397_v141 = _dafny.MultiSetOf(_395_v139, _395_v139)
		var _398_v143 _dafny.Array
		_ = _398_v143
		var _nwElement0_15 _dafny.Set = _396_v140
		_ = _nwElement0_15
		var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(13))
		_ = _nw62
		(_nw62).ArraySet1(_nwElement0_15, 0)
		(_nw62).ArraySet1(_396_v140, 1)
		(_nw62).ArraySet1((_396_v140).Intersection(_396_v140), 2)
		(_nw62).ArraySet1(Companion_Default___.Fm8(_233_v7, true, (_301_v55).F29(), (_301_v55).F29(), _246_globalState), 3)
		(_nw62).ArraySet1((_396_v140).Union(_396_v140), 4)
		(_nw62).ArraySet1(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter5 := _dafny.Iterate((_397_v141).Elements()); ; {
				_compr_4, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _399_v142 _dafny.Sequence
				_399_v142 = interface{}(_compr_4).(_dafny.Sequence)
				if (_397_v141).Contains(_399_v142) {
					_coll4.Add(_399_v142)
				}
			}
			return _coll4.ToSet()
		}(), 5)
		(_nw62).ArraySet1(_396_v140, 6)
		(_nw62).ArraySet1(_396_v140, 7)
		(_nw62).ArraySet1(_396_v140, 8)
		(_nw62).ArraySet1(_396_v140, 9)
		(_nw62).ArraySet1(_dafny.SetOf(_395_v139, _395_v139), 10)
		(_nw62).ArraySet1((_396_v140).Union(_dafny.SetOf(_395_v139, _395_v139)), 11)
		(_nw62).ArraySet1(_396_v140, 12)
		_398_v143 = _nw62
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_398_v143), 0))
		_ = _index49
		(_398_v143).ArraySet1(_396_v140, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_398_v143), 0))
		_ = _index50
		var _rhs46 bool = !(_226_v0)
		_ = _rhs46
		var _rhs47 _dafny.Set = (_396_v140).Intersection(Companion_Default___.Fm8(_255_v23, false, (_301_v55).F29(), _226_v0, _246_globalState))
		_ = _rhs47
		var _lhs34 *GlobalState = _246_globalState
		_ = _lhs34
		var _lhs35 _dafny.Array = _398_v143
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_398_v143), 0))
		_ = _lhs36
		_lhs34.F14 = _rhs46
		(_lhs35).ArraySet1(_rhs47, (_lhs36).Int())
		_255_v23 = _255_v23
		if (_230_v4).Select((Companion_Default___.SafeIndex(_385_v130, _dafny.IntOfUint32((_230_v4).Cardinality()))).Uint32()).(bool) {
			var _400_v144 D1
			_ = _400_v144
			_400_v144 = Companion_D1_.Create_DC4_((_301_v55).F29(), _385_v130)
			_226_v0 = (_400_v144).Dtor_cf5()
			var _401_v145 _dafny.Set
			_ = _401_v145
			_401_v145 = _234_v8
			var _402_v146 _dafny.Sequence
			_ = _402_v146
			_402_v146 = _dafny.SeqOf(_dafny.SetOf((_301_v55).F29(), (_301_v55).F29()), Companion_Default___.Fm9(_255_v23, _246_globalState))
			var _403_v147 _dafny.Array
			_ = _403_v147
			var _nwElement0_16 _dafny.Set = _234_v8
			_ = _nwElement0_16
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(18))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_16, 0)
			(_nw63).ArraySet1((func() _dafny.Set {
				if (_301_v55).F29() {
					return _dafny.SetOf(_226_v0, (_400_v144).Dtor_cf5())
				}
				return _234_v8
			})(), 1)
			(_nw63).ArraySet1((func() _dafny.Set {
				if _226_v0 {
					return _234_v8
				}
				return _234_v8
			})(), 2)
			(_nw63).ArraySet1(Companion_Default___.Fm9(_255_v23, _246_globalState), 3)
			(_nw63).ArraySet1((_401_v145), 4)
			(_nw63).ArraySet1(_234_v8, 5)
			(_nw63).ArraySet1(_dafny.SetOf(_226_v0, (_301_v55).F29(), (_301_v55).F29(), (_301_v55).F29()), 6)
			(_nw63).ArraySet1((_234_v8).Intersection(_234_v8), 7)
			(_nw63).ArraySet1((_dafny.SetOf((_301_v55).F29())).Intersection(_dafny.SetOf((_301_v55).F29(), (_301_v55).F29())), 8)
			(_nw63).ArraySet1(_234_v8, 9)
			(_nw63).ArraySet1((_402_v146).Select((Companion_Default___.SafeIndex(_233_v7, _dafny.IntOfUint32((_402_v146).Cardinality()))).Uint32()).(_dafny.Set), 10)
			(_nw63).ArraySet1((_dafny.SetOf(_226_v0, !((_301_v55).F29()))).Difference(Companion_Default___.Fm9(_dafny.IntOfInt64(629), _246_globalState)), 11)
			(_nw63).ArraySet1(_234_v8, 12)
			(_nw63).ArraySet1((_234_v8).Difference(_234_v8), 13)
			(_nw63).ArraySet1(_234_v8, 14)
			(_nw63).ArraySet1(_234_v8, 15)
			(_nw63).ArraySet1(_234_v8, 16)
			(_nw63).ArraySet1((_234_v8).Union(_234_v8), 17)
			_403_v147 = _nw63
			var _404_v148 D1
			_ = _404_v148
			_404_v148 = Companion_D1_.Create_DC5_(_233_v7, (_301_v55).F29(), _385_v130, (_301_v55).F29())
			var _405_v149 D1
			_ = _405_v149
			_405_v149 = Companion_D1_.Create_DC6_(_404_v148)
			var _406_v150 D2
			_ = _406_v150
			_406_v150 = Companion_D2_.Create_DC9_(_226_v0, _301_v55)
			var _rhs48 _dafny.Array = _403_v147
			_ = _rhs48
			var _rhs49 bool = !((_237_v12).Merge(_237_v12)).Contains(true)
			_ = _rhs49
			var _rhs50 D1 = Companion_Default___.Fm10(_226_v0, _dafny.IntOfInt64(620), (_301_v55).F29(), _246_globalState)
			_ = _rhs50
			var _rhs51 bool = (_406_v150).Dtor_cf14()
			_ = _rhs51
			var _lhs37 *GlobalState = _246_globalState
			_ = _lhs37
			var _lhs38 *GlobalState = _246_globalState
			_ = _lhs38
			_403_v147 = _rhs48
			_lhs37.F14 = _rhs49
			_405_v149 = _rhs50
			_lhs38.F14 = _rhs51
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_235_v9), 0))
			_ = _index51
			(_235_v9).ArraySet1(_255_v23, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_235_v9), 0))
			_ = _index52
			(_235_v9).ArraySet1(_385_v130, (_index52).Int())
			(_246_globalState).F14 = (_301_v55).F29()
			_226_v0 = (_233_v7).Cmp(_233_v7) <= 0
		} else {
			var _407_v151 _dafny.Array
			_ = _407_v151
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw64
			_407_v151 = _nw64
			var _408_v152 _dafny.MultiSet
			_ = _408_v152
			_408_v152 = _dafny.MultiSetOf(_255_v23)
			var _409_v153 D2
			_ = _409_v153
			_409_v153 = Companion_D2_.Create_DC9_((_301_v55).F29(), _301_v55)
			var _410_v154 D2
			_ = _410_v154
			_410_v154 = Companion_D2_.Create_DC9_((_301_v55).F29(), (_409_v153).Dtor_cf15())
			var _411_v155 _dafny.Map
			_ = _411_v155
			_411_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_Default___.Fm11(_385_v130, _408_v152, (_301_v55).F29(), _246_globalState)).Dtor_cf5()), _409_v153)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_407_v151), 0))
			_ = _index53
			(_407_v151).ArraySet1(_411_v155, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_407_v151), 0))
			_ = _index54
			(_407_v151).ArraySet1(_411_v155, (_index54).Int())
			_253_v21 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_253_v21, (Companion_Default___.SafeIndex(_255_v23, _dafny.IntOfUint32((_253_v21).Cardinality()))).Uint32(), _243_v16), _383_v128)
			var _412_v156 *C0
			_ = _412_v156
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__(_226_v0)
			_412_v156 = _nw65
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_254_v22), 0))
			_ = _index55
			(_254_v22).ArraySet1(_233_v7, (_index55).Int())
			var _413_v157 _dafny.Array
			_ = _413_v157
			var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw66
			_413_v157 = _nw66
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_413_v157), 0))
			_ = _index56
			(_413_v157).ArraySet1(_239_v14, (_index56).Int())
			var _414_v158 _dafny.Set
			_ = _414_v158
			_414_v158 = _dafny.SetOf(_233_v7, (_237_v12).Cardinality(), _dafny.IntOfInt64(-263), (_dafny.Zero).Minus((_dafny.SetOf(_395_v139, _395_v139)).Cardinality()), (_dafny.SetOf((_dafny.Zero).Minus(_255_v23))).Cardinality())
			var _415_v159 _dafny.Map
			_ = _415_v159
			_415_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_412_v156).F29(), _254_v22)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_254_v22), 0))
			_ = _index57
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_413_v157), 0))
			_ = _index58
			var _rhs52 _dafny.Int = (_dafny.Zero).Minus((_385_v130).Times((_414_v158).Cardinality()))
			_ = _rhs52
			var _rhs53 _dafny.Array = (func() _dafny.Array {
				if (_415_v159).Contains((_dafny.IntOfUint32((_253_v21).Cardinality())).Cmp(_233_v7) != 0) {
					return (_415_v159).Get((_dafny.IntOfUint32((_253_v21).Cardinality())).Cmp(_233_v7) != 0).(_dafny.Array)
				}
				return _254_v22
			})()
			_ = _rhs53
			var _rhs54 D2 = _409_v153
			_ = _rhs54
			var _rhs55 _dafny.Array = _239_v14
			_ = _rhs55
			var _rhs56 bool = Companion_Default___.Fm0(_229_v3, (_301_v55).F29(), _246_globalState)
			_ = _rhs56
			var _lhs39 _dafny.Array = _254_v22
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_254_v22), 0))
			_ = _lhs40
			var _lhs41 *GlobalState = _246_globalState
			_ = _lhs41
			var _lhs42 _dafny.Array = _413_v157
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_413_v157), 0))
			_ = _lhs43
			var _lhs44 *GlobalState = _246_globalState
			_ = _lhs44
			(_lhs39).ArraySet1(_rhs52, (_lhs40).Int())
			_lhs41.F28 = _rhs53
			_410_v154 = _rhs54
			(_lhs42).ArraySet1(_rhs55, (_lhs43).Int())
			_lhs44.F14 = _rhs56
			(_246_globalState).F17 = Companion_Default___.SafeDivisionInt(_233_v7, _233_v7)
		}
		var _416_v160 _dafny.Set
		_ = _416_v160
		_416_v160 = _dafny.SetOf(_255_v23)
		var _417_v161 D1
		_ = _417_v161
		_417_v161 = Companion_D1_.Create_DC5_(_233_v7, _226_v0, Companion_Default___.SafeDivisionInt((_416_v160).Cardinality(), _255_v23), _226_v0)
		var _source7 D1 = _417_v161
		_ = _source7
		if _source7.Is_DC3() {
			var _418___mcc_h2 _dafny.Int = _source7.Get_().(D1_DC3).Cf3
			_ = _418___mcc_h2
			var _419___mcc_h3 _dafny.Int = _source7.Get_().(D1_DC3).Cf4
			_ = _419___mcc_h3
			var _420_cf4 _dafny.Int = _419___mcc_h3
			_ = _420_cf4
			var _421_cf3 _dafny.Int = _418___mcc_h2
			_ = _421_cf3
			(_246_globalState).F14 = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_253_v21, _253_v21))
			var _422_v162 _dafny.Sequence
			_ = _422_v162
			_422_v162 = _dafny.SeqOf(_229_v3, _dafny.UnicodeSeqOfUtf8Bytes("msee"))
			var _423_v163 _dafny.Set
			_ = _423_v163
			_423_v163 = _dafny.SetOf(_384_v129)
			var _424_v164 D4
			_ = _424_v164
			_424_v164 = Companion_D4_.Create_DC12_(_235_v9)
			var _rhs57 _dafny.Int = (_255_v23).Minus(_dafny.IntOfUint32(((_422_v162).Select((Companion_Default___.SafeIndex(_255_v23, _dafny.IntOfUint32((_422_v162).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			_ = _rhs57
			var _rhs58 bool = (_423_v163).IsProperSubsetOf((_423_v163).Difference(_dafny.SetOf((_424_v164).Dtor_cf19(), _235_v9, _384_v129)))
			_ = _rhs58
			var _lhs45 *GlobalState = _246_globalState
			_ = _lhs45
			var _lhs46 *GlobalState = _246_globalState
			_ = _lhs46
			_lhs45.F13 = _rhs57
			_lhs46.F14 = _rhs58
			var _425_v165 _dafny.MultiSet
			_ = _425_v165
			_425_v165 = _dafny.MultiSetOf(_230_v4)
			_425_v165 = ((_425_v165).Update(_230_v4, Companion_Default___.Abs(_255_v23))).Difference(_425_v165)
			var _rhs59 bool = (_301_v55).F29()
			_ = _rhs59
			var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_230_v4, _dafny.Companion_Sequence_.Concatenate(_230_v4, _230_v4))
			_ = _rhs60
			var _lhs47 *GlobalState = _246_globalState
			_ = _lhs47
			_lhs47.F14 = _rhs59
			_230_v4 = _rhs60
		} else if _source7.Is_DC4() {
			var _426___mcc_h4 bool = _source7.Get_().(D1_DC4).Cf5
			_ = _426___mcc_h4
			var _427___mcc_h5 _dafny.Int = _source7.Get_().(D1_DC4).Cf6
			_ = _427___mcc_h5
			var _428_cf6 _dafny.Int = _427___mcc_h5
			_ = _428_cf6
			var _429_cf5 bool = _426___mcc_h4
			_ = _429_cf5
			var _430_v166 _dafny.Sequence
			_ = _430_v166
			_430_v166 = _dafny.SeqOf(_229_v3)
			(_246_globalState).F14 = !(!(!(!_dafny.Companion_Sequence_.Contains(_430_v166, _383_v128))))
			(_246_globalState).F27 = (_395_v139).Select((Companion_Default___.SafeIndex(_255_v23, _dafny.IntOfUint32((_395_v139).Cardinality()))).Uint32()).(_dafny.Int)
			var _431_v167 *C0
			_ = _431_v167
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__(_429_cf5)
			_431_v167 = _nw67
			var _432_v168 _dafny.Array
			_ = _432_v168
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw68
			_432_v168 = _nw68
			var _433_v169 D5
			_ = _433_v169
			_433_v169 = Companion_D5_.Create_DC14_(_432_v168)
			var _434_v170 _dafny.Array
			_ = _434_v170
			var _nwElement0_17 _dafny.Array = _432_v168
			_ = _nwElement0_17
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(28))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_17, 0)
			(_nw69).ArraySet1(_432_v168, 1)
			(_nw69).ArraySet1(_432_v168, 2)
			(_nw69).ArraySet1(_432_v168, 3)
			(_nw69).ArraySet1(_432_v168, 4)
			(_nw69).ArraySet1(_432_v168, 5)
			(_nw69).ArraySet1(_432_v168, 6)
			(_nw69).ArraySet1(_432_v168, 7)
			(_nw69).ArraySet1(_432_v168, 8)
			(_nw69).ArraySet1(_432_v168, 9)
			(_nw69).ArraySet1(_432_v168, 10)
			(_nw69).ArraySet1(_432_v168, 11)
			(_nw69).ArraySet1(_432_v168, 12)
			(_nw69).ArraySet1(_432_v168, 13)
			(_nw69).ArraySet1(_432_v168, 14)
			(_nw69).ArraySet1(_432_v168, 15)
			(_nw69).ArraySet1(_432_v168, 16)
			(_nw69).ArraySet1(_432_v168, 17)
			(_nw69).ArraySet1(_432_v168, 18)
			(_nw69).ArraySet1((func() _dafny.Array {
				if (_301_v55).F29() {
					return _432_v168
				}
				return _432_v168
			})(), 19)
			(_nw69).ArraySet1(_432_v168, 20)
			(_nw69).ArraySet1(_432_v168, 21)
			(_nw69).ArraySet1(_432_v168, 22)
			(_nw69).ArraySet1((_433_v169).Dtor_cf24(), 23)
			(_nw69).ArraySet1(_432_v168, 24)
			(_nw69).ArraySet1(_432_v168, 25)
			(_nw69).ArraySet1(_432_v168, 26)
			(_nw69).ArraySet1(_432_v168, 27)
			_434_v170 = _nw69
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_434_v170), 0))
			_ = _index59
			(_434_v170).ArraySet1(_432_v168, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_434_v170), 0))
			_ = _index60
			(_434_v170).ArraySet1((func() _dafny.Array {
				if (_301_v55).F29() {
					return _432_v168
				}
				return _432_v168
			})(), (_index60).Int())
		} else if _source7.Is_DC5() {
			var _435___mcc_h6 _dafny.Int = _source7.Get_().(D1_DC5).Cf7
			_ = _435___mcc_h6
			var _436___mcc_h7 bool = _source7.Get_().(D1_DC5).Cf8
			_ = _436___mcc_h7
			var _437___mcc_h8 _dafny.Int = _source7.Get_().(D1_DC5).Cf9
			_ = _437___mcc_h8
			var _438___mcc_h9 bool = _source7.Get_().(D1_DC5).Cf10
			_ = _438___mcc_h9
			var _439_cf10 bool = _438___mcc_h9
			_ = _439_cf10
			var _440_cf9 _dafny.Int = _437___mcc_h8
			_ = _440_cf9
			var _441_cf8 bool = _436___mcc_h7
			_ = _441_cf8
			var _442_cf7 _dafny.Int = _435___mcc_h6
			_ = _442_cf7
			var _443_v171 _dafny.Array
			_ = _443_v171
			var _nwElement0_18 *C0 = _301_v55
			_ = _nwElement0_18
			var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.One)
			_ = _nw70
			(_nw70).ArraySet1(_nwElement0_18, 0)
			_443_v171 = _nw70
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_443_v171), 0))
			_ = _index61
			(_443_v171).ArraySet1(_301_v55, (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_443_v171), 0))
			_ = _index62
			(_443_v171).ArraySet1(_301_v55, (_index62).Int())
			var _444_v172 D1
			_ = _444_v172
			_444_v172 = Companion_D1_.Create_DC2_(_253_v21)
			var _445_v173 _dafny.Sequence
			_ = _445_v173
			var _446_v174 _dafny.Array
			_ = _446_v174
			var _447_v175 _dafny.Int
			_ = _447_v175
			var _out36 _dafny.Sequence
			_ = _out36
			var _out37 _dafny.Array
			_ = _out37
			var _out38 _dafny.Int
			_ = _out38
			_out36, _out37, _out38 = Companion_Default___.M0((_444_v172).Dtor_cf2(), _246_globalState)
			_445_v173 = _out36
			_446_v174 = _out37
			_447_v175 = _out38
			(_246_globalState).F17 = (_395_v139).Select((Companion_Default___.SafeIndex(_233_v7, _dafny.IntOfUint32((_395_v139).Cardinality()))).Uint32()).(_dafny.Int)
			var _448_v176 _dafny.MultiSet
			_ = _448_v176
			_448_v176 = _dafny.MultiSetOf(_243_v16, _243_v16)
			var _449_v177 _dafny.MultiSet
			_ = _449_v177
			_449_v177 = _dafny.MultiSetOf(Companion_Default___.Fm2((_448_v176).Update(_243_v16, Companion_Default___.Abs(_233_v7)), _440_cf9, _447_v175, (_301_v55).F29(), _246_globalState), _dafny.IntOfUint32((_395_v139).Cardinality()), _440_cf9)
			_449_v177 = _449_v177
		} else if _source7.Is_DC2() {
			var _450___mcc_h10 _dafny.Sequence = _source7.Get_().(D1_DC2).Cf2
			_ = _450___mcc_h10
			var _451_cf2 _dafny.Sequence = _450___mcc_h10
			_ = _451_cf2
			var _452_v178 _dafny.Sequence
			_ = _452_v178
			var _453_v179 _dafny.Array
			_ = _453_v179
			var _454_v180 _dafny.Int
			_ = _454_v180
			var _out39 _dafny.Sequence
			_ = _out39
			var _out40 _dafny.Array
			_ = _out40
			var _out41 _dafny.Int
			_ = _out41
			_out39, _out40, _out41 = Companion_Default___.M0(_dafny.Companion_Sequence_.Update(_451_cf2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-918), _dafny.IntOfUint32((_451_cf2).Cardinality()))).Uint32(), (_383_v128).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-862), _dafny.IntOfUint32((_383_v128).Cardinality()))).Uint32()).(_dafny.CodePoint)), _246_globalState)
			_452_v178 = _out39
			_453_v179 = _out40
			_454_v180 = _out41
			(_246_globalState).F11 = ((_dafny.MultiSetOf((_301_v55).F29())).Union((_227_v1).Update(_226_v0, Companion_Default___.Abs(_454_v180)))).Intersection(_227_v1)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_239_v14), 0))
			_ = _index63
			(_239_v14).ArraySet1((_301_v55).F29(), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_239_v14), 0))
			_ = _index64
			(_239_v14).ArraySet1(!(_226_v0), (_index64).Int())
			var _455_v181 _dafny.MultiSet
			_ = _455_v181
			_455_v181 = _dafny.MultiSetOf(_dafny.IntOfInt64(-929), _233_v7)
			var _456_v182 D6
			_ = _456_v182
			_456_v182 = Companion_D6_.Create_DC16_(_455_v181)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_239_v14), 0))
			_ = _index65
			var _rhs61 *C0 = _301_v55
			_ = _rhs61
			var _rhs62 _dafny.Sequence = (func() _dafny.Sequence {
				if (_239_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_239_v14), 0))).Int()).(bool) {
					return _dafny.Companion_Sequence_.Update(_230_v4, (Companion_Default___.SafeIndex(_385_v130, _dafny.IntOfUint32((_230_v4).Cardinality()))).Uint32(), true)
				}
				return _230_v4
			})()
			_ = _rhs62
			var _rhs63 bool = !((_455_v181).IsSubsetOf((_456_v182).Dtor_cf30()))
			_ = _rhs63
			var _rhs64 _dafny.Sequence = _230_v4
			_ = _rhs64
			var _lhs48 _dafny.Array = _239_v14
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_239_v14), 0))
			_ = _lhs49
			_301_v55 = _rhs61
			_230_v4 = _rhs62
			(_lhs48).ArraySet1(_rhs63, (_lhs49).Int())
			_230_v4 = _rhs64
		} else {
			var _457___mcc_h11 D1 = _source7.Get_().(D1_DC6).Cf11
			_ = _457___mcc_h11
			var _458_cf11 D1 = _457___mcc_h11
			_ = _458_cf11
			var _459_v183 _dafny.Sequence
			_ = _459_v183
			var _460_v184 _dafny.Array
			_ = _460_v184
			var _461_v185 _dafny.Int
			_ = _461_v185
			var _out42 _dafny.Sequence
			_ = _out42
			var _out43 _dafny.Array
			_ = _out43
			var _out44 _dafny.Int
			_ = _out44
			_out42, _out43, _out44 = Companion_Default___.M0(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ynbc"), (Companion_Default___.SafeIndex(_233_v7, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ynbc")).Cardinality()))).Uint32(), _243_v16), _246_globalState)
			_459_v183 = _out42
			_460_v184 = _out43
			_461_v185 = _out44
			_242_v15 = _242_v15
			var _462_v186 _dafny.Sequence
			_ = _462_v186
			var _463_v187 _dafny.Array
			_ = _463_v187
			var _464_v188 _dafny.Int
			_ = _464_v188
			var _out45 _dafny.Sequence
			_ = _out45
			var _out46 _dafny.Array
			_ = _out46
			var _out47 _dafny.Int
			_ = _out47
			_out45, _out46, _out47 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_253_v21, _253_v21), _246_globalState)
			_462_v186 = _out45
			_463_v187 = _out46
			_464_v188 = _out47
			var _465_v189 _dafny.Map
			_ = _465_v189
			_465_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_461_v185, _384_v129)
			(_246_globalState).F14 = Companion_Default___.Fm0(_dafny.Companion_Sequence_.Concatenate(_459_v183, Companion_Default___.Fm12(_233_v7, (_465_v189).Cardinality(), (_301_v55).F29(), _246_globalState)), (_301_v55).F29(), _246_globalState)
		}
		(_246_globalState).F27 = (_dafny.Zero).Minus((_385_v130).Times(_dafny.IntOfUint32((_395_v139).Cardinality())))
	}
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_239_v14), 0))); ; {
		_guard_loop_1, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _466_i13 _dafny.Int
		_466_i13 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_466_i13).Sign() != -1) && ((_466_i13).Cmp(_dafny.ArrayLen((_239_v14), 0)) < 0)) {
			(_239_v14).ArraySet1((_301_v55).F29(), (_466_i13).Int())
		}
	}
	_226_v0 = (_301_v55).F29()
	_dafny.Print(_226_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v1).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_228_v2, _dafny.SeqOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_229_v3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_230_v4, _dafny.SeqOf(true, true, true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(true, true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_233_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v8).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false).UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_238_v13, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v14).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_243_v16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v17).Equals(_dafny.SetOf(_dafny.CodePoint('a'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_245_v18).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), true), _dafny.SetOf(_dafny.CodePoint('a')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_246_globalState).F0(), _dafny.SeqOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F7()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(true, true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F11).Equals(_dafny.MultiSetOf(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F18()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F19()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.IntOfInt64(773))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F24()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), true), _dafny.SetOf(_dafny.CodePoint('a')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F25)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F26()).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_globalState.F27)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F28).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_248_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_253_v21.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v22).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_255_v23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_280_v36).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v37).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v55).F29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_383_v128.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v129).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_385_v130)
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
	Cf1 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC1:
		{
			_, ok := other.Get_().(D0_DC1)
			return ok
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
	Cf3 _dafny.Int
	Cf4 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Int, Cf4 _dafny.Int) D1 {
	return D1{D1_DC3{Cf3, Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf5 bool
	Cf6 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 bool, Cf6 _dafny.Int) D1 {
	return D1{D1_DC4{Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf7  _dafny.Int
	Cf8  bool
	Cf9  _dafny.Int
	Cf10 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 _dafny.Int, Cf8 bool, Cf9 _dafny.Int, Cf10 bool) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC2 struct {
	Cf2 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC6 struct {
	Cf11 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf11 D1) D1 {
	return D1{D1_DC6{Cf11}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf11() D1 {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + data.Cf2.VerbatimString(true) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
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
	Cf13 _dafny.Sequence
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Sequence) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf14 bool
	Cf15 *C0
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf14 bool, Cf15 *C0) D2 {
	return D2{D2_DC9{Cf14, Cf15}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC10 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.CodePoint
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf16 _dafny.Int, Cf17 _dafny.CodePoint) D2 {
	return D2{D2_DC10{Cf16, Cf17}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

type D2_DC7 struct {
	Cf12 _dafny.Set
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.Set) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.EmptySeq)
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() bool {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf15() *C0 {
	return _this.Get_().(D2_DC9).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D2_DC10).Cf16
}

func (_this D2) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D2_DC10).Cf17
}

func (_this D2) Dtor_cf12() _dafny.Set {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + data.Cf13.VerbatimString(true) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
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
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D3_DC11 struct {
	Cf18 _dafny.Set
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf18 _dafny.Set) D3 {
	return D3{D3_DC11{Cf18}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D3) Dtor_cf18() _dafny.Set {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D4_DC13 struct {
	Cf20 *C0
	Cf21 _dafny.CodePoint
	Cf22 _dafny.CodePoint
	Cf23 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf20 *C0, Cf21 _dafny.CodePoint, Cf22 _dafny.CodePoint, Cf23 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC12 struct {
	Cf19 _dafny.Array
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf19 _dafny.Array) D4 {
	return D4{D4_DC12{Cf19}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_((*C0)(nil), _dafny.CodePoint('D'), _dafny.CodePoint('D'), _dafny.EmptySeq)
}

func (_this D4) Dtor_cf20() *C0 {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf21() _dafny.CodePoint {
	return _this.Get_().(D4_DC13).Cf21
}

func (_this D4) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf23
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf19 == data2.Cf19
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

type D5_DC15 struct {
	Cf25 _dafny.Sequence
	Cf26 bool
	Cf27 bool
	Cf28 bool
	Cf29 _dafny.Array
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf25 _dafny.Sequence, Cf26 bool, Cf27 bool, Cf28 bool, Cf29 _dafny.Array) D5 {
	return D5{D5_DC15{Cf25, Cf26, Cf27, Cf28, Cf29}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf24 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf24 _dafny.Array) D5 {
	return D5{D5_DC14{Cf24}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(_dafny.EmptySeq, false, false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D5) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D5_DC15).Cf25
}

func (_this D5) Dtor_cf26() bool {
	return _this.Get_().(D5_DC15).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC15).Cf27
}

func (_this D5) Dtor_cf28() bool {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf29
}

func (_this D5) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf24
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + data.Cf25.VerbatimString(true) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf24 == data2.Cf24
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
	Cf31 _dafny.CodePoint
	Cf32 bool
	Cf33 *C0
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf31 _dafny.CodePoint, Cf32 bool, Cf33 *C0) D6 {
	return D6{D6_DC17{Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf34 _dafny.Sequence
	Cf35 _dafny.Int
	Cf36 bool
	Cf37 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf34 _dafny.Sequence, Cf35 _dafny.Int, Cf36 bool, Cf37 _dafny.Int) D6 {
	return D6{D6_DC18{Cf34, Cf35, Cf36, Cf37}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC16 struct {
	Cf30 _dafny.MultiSet
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf30 _dafny.MultiSet) D6 {
	return D6{D6_DC16{Cf30}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.CodePoint('D'), false, (*C0)(nil))
}

func (_this D6) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC17).Cf32
}

func (_this D6) Dtor_cf33() *C0 {
	return _this.Get_().(D6_DC17).Cf33
}

func (_this D6) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D6_DC18).Cf34
}

func (_this D6) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf35
}

func (_this D6) Dtor_cf36() bool {
	return _this.Get_().(D6_DC18).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf37
}

func (_this D6) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D6_DC16).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + data.Cf34.VerbatimString(true) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf30) + ")"
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
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf34.Equals(data2.Cf34) && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D7_DC20 struct {
	Cf39 bool
	Cf40 bool
	Cf41 _dafny.Int
	Cf42 bool
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf39 bool, Cf40 bool, Cf41 _dafny.Int, Cf42 bool) D7 {
	return D7{D7_DC20{Cf39, Cf40, Cf41, Cf42}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC19 struct {
	Cf38 _dafny.MultiSet
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf38 _dafny.MultiSet) D7 {
	return D7{D7_DC19{Cf38}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC21 struct {
	Cf43 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf43 D7) D7 {
	return D7{D7_DC21{Cf43}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(false, false, _dafny.Zero, false)
}

func (_this D7) Dtor_cf39() bool {
	return _this.Get_().(D7_DC20).Cf39
}

func (_this D7) Dtor_cf40() bool {
	return _this.Get_().(D7_DC20).Cf40
}

func (_this D7) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf41
}

func (_this D7) Dtor_cf42() bool {
	return _this.Get_().(D7_DC20).Cf42
}

func (_this D7) Dtor_cf38() _dafny.MultiSet {
	return _this.Get_().(D7_DC19).Cf38
}

func (_this D7) Dtor_cf43() D7 {
	return _this.Get_().(D7_DC21).Cf43
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40 && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf38.Equals(data2.Cf38)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf43.Equals(data2.Cf43)
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
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_() D8 {
	return D8{D8_DC23{}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf45 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf45 bool) D8 {
	return D8{D8_DC24{Cf45}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf44 _dafny.Sequence
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf44 _dafny.Sequence) D8 {
	return D8{D8_DC22{Cf44}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC25 struct {
	Cf46 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf46 D8) D8 {
	return D8{D8_DC25{Cf46}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_()
}

func (_this D8) Dtor_cf45() bool {
	return _this.Get_().(D8_DC24).Cf45
}

func (_this D8) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D8_DC22).Cf44
}

func (_this D8) Dtor_cf46() D8 {
	return _this.Get_().(D8_DC25).Cf46
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf46) + ")"
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
			_, ok := other.Get_().(D8_DC23)
			return ok
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf45 == data2.Cf45
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf44.Equals(data2.Cf44)
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf46.Equals(data2.Cf46)
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
	F2   _dafny.CodePoint
	F11  _dafny.MultiSet
	F13  _dafny.Int
	F14  bool
	F17  _dafny.Int
	F25  _dafny.Int
	F27  _dafny.Int
	F28  _dafny.Array
	_f0  _dafny.Sequence
	_f1  _dafny.Int
	_f3  _dafny.Int
	_f4  bool
	_f5  _dafny.Int
	_f6  _dafny.Sequence
	_f7  _dafny.Map
	_f8  _dafny.Sequence
	_f9  _dafny.Array
	_f10 _dafny.Int
	_f12 _dafny.Int
	_f15 _dafny.Int
	_f16 _dafny.Int
	_f18 _dafny.MultiSet
	_f19 _dafny.Map
	_f20 _dafny.Array
	_f21 _dafny.Int
	_f22 bool
	_f23 _dafny.Int
	_f24 _dafny.Map
	_f26 _dafny.MultiSet
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.CodePoint('D')
	_this.F11 = _dafny.EmptyMultiSet
	_this.F13 = _dafny.Zero
	_this.F14 = false
	_this.F17 = _dafny.Zero
	_this.F25 = _dafny.Zero
	_this.F27 = _dafny.Zero
	_this.F28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.EmptySeq
	_this._f1 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.EmptySeq
	_this._f7 = _dafny.EmptyMap
	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.Zero
	_this._f18 = _dafny.EmptyMultiSet
	_this._f19 = _dafny.EmptyMap
	_this._f20 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.Zero
	_this._f22 = false
	_this._f23 = _dafny.Zero
	_this._f24 = _dafny.EmptyMap
	_this._f26 = _dafny.EmptyMultiSet
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 _dafny.Int, f2 _dafny.CodePoint, f3 _dafny.Int, f4 bool, f5 _dafny.Int, f6 _dafny.Sequence, f7 _dafny.Map, f8 _dafny.Sequence, f9 _dafny.Array, f10 _dafny.Int, f11 _dafny.MultiSet, f12 _dafny.Int, f13 _dafny.Int, f14 bool, f15 _dafny.Int, f16 _dafny.Int, f17 _dafny.Int, f18 _dafny.MultiSet, f19 _dafny.Map, f20 _dafny.Array, f21 _dafny.Int, f22 bool, f23 _dafny.Int, f24 _dafny.Map, f25 _dafny.Int, f26 _dafny.MultiSet, f27 _dafny.Int, f28 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this).F25 = f25
		(_this)._f26 = f26
		(_this).F27 = f27
		(_this).F28 = f28
	}
}
func (_this *GlobalState) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Map {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Array {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F18() _dafny.MultiSet {
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
func (_this *GlobalState) F21() _dafny.Int {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() bool {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Int {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() _dafny.Map {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F26() _dafny.MultiSet {
	{
		return _this._f26
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f29 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f29 = false
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

func (_this *C0) Ctor__(f29 bool) {
	{
		(_this)._f29 = f29
	}
}
func (_this *C0) F29() bool {
	{
		return _this._f29
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
