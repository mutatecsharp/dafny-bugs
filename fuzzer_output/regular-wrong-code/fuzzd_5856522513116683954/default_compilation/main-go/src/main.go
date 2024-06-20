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
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.CodePoint, p1 D0, p2 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.MultiSetOf(false)).Equals(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(false))))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) bool {
	return ((_dafny.SetOf(false, true)).Cardinality()).Cmp(_dafny.IntOfInt64(743)) < 0
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("gyeq"), _dafny.UnicodeSeqOfUtf8Bytes("gvla")), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(192), _dafny.IntOfInt64(920)))).Equals(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(606))).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfInt64(-374)).Plus(_dafny.IntOfInt64(980)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(252))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()), ((_dafny.SetOf(true)).Cardinality()).Times(_dafny.IntOfInt64(99)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Int {
	return (func(_source0 D8) _dafny.Map {
		if _source0.Is_DC23() {
			var _1___mcc_h0 bool = _source0.Get_().(D8_DC23).Cf35
			_ = _1___mcc_h0
			var _2___mcc_h1 bool = _source0.Get_().(D8_DC23).Cf36
			_ = _2___mcc_h1
			var _3___mcc_h2 _dafny.Int = _source0.Get_().(D8_DC23).Cf37
			_ = _3___mcc_h2
			var _4_cf37 _dafny.Int = _3___mcc_h2
			_ = _4_cf37
			var _5_cf36 bool = _2___mcc_h1
			_ = _5_cf36
			var _6_cf35 bool = _1___mcc_h0
			_ = _6_cf35
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_cf36, _5_cf36)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _6_cf35))
		} else if _source0.Is_DC24() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
		} else {
			var _7___mcc_h3 _dafny.Map = _source0.Get_().(D8_DC22).Cf34
			_ = _7___mcc_h3
			var _8_cf34 _dafny.Map = _7___mcc_h3
			_ = _8_cf34
			return (Companion_D16_.Create_DC39_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Dtor_cf57()
		}
	}(Companion_D8_.Create_DC24_())).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_((Companion_D8_.Create_DC23_(true, true, (_dafny.SetOf(_dafny.IntOfInt64(800))).Cardinality())).Dtor_cf36())
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ftpspqpu")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(740)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-788))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(673))))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.MultiSet {
	var _source1 D8 = Companion_D8_.Create_DC24_()
	_ = _source1
	if _source1.Is_DC23() {
		var _9___mcc_h0 bool = _source1.Get_().(D8_DC23).Cf35
		_ = _9___mcc_h0
		var _10___mcc_h1 bool = _source1.Get_().(D8_DC23).Cf36
		_ = _10___mcc_h1
		var _11___mcc_h2 _dafny.Int = _source1.Get_().(D8_DC23).Cf37
		_ = _11___mcc_h2
		var _12_cf37 _dafny.Int = _11___mcc_h2
		_ = _12_cf37
		var _13_cf36 bool = _10___mcc_h1
		_ = _13_cf36
		var _14_cf35 bool = _9___mcc_h0
		_ = _14_cf35
		return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(676))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_15_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(263)
		})))
	} else if _source1.Is_DC24() {
		return _dafny.MultiSetOf(_dafny.One)
	} else {
		var _16___mcc_h3 _dafny.Map = _source1.Get_().(D8_DC22).Cf34
		_ = _16___mcc_h3
		var _17_cf34 _dafny.Map = _16___mcc_h3
		_ = _17_cf34
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(905)))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) D0 {
	var _source2 D7 = Companion_D7_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(946), true))
	_ = _source2
	if _source2.Is_DC19() {
		var _18___mcc_h0 D1 = _source2.Get_().(D7_DC19).Cf27
		_ = _18___mcc_h0
		var _19___mcc_h1 _dafny.Int = _source2.Get_().(D7_DC19).Cf28
		_ = _19___mcc_h1
		var _20___mcc_h2 bool = _source2.Get_().(D7_DC19).Cf29
		_ = _20___mcc_h2
		var _21___mcc_h3 _dafny.Array = _source2.Get_().(D7_DC19).Cf30
		_ = _21___mcc_h3
		var _22_cf30 _dafny.Array = _21___mcc_h3
		_ = _22_cf30
		var _23_cf29 bool = _20___mcc_h2
		_ = _23_cf29
		var _24_cf28 _dafny.Int = _19___mcc_h1
		_ = _24_cf28
		var _25_cf27 D1 = _18___mcc_h0
		_ = _25_cf27
		return Companion_D0_.Create_DC1_(_23_cf29, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("whpnil")), _dafny.UnicodeSeqOfUtf8Bytes("xdqneud"))
	} else if _source2.Is_DC20() {
		var _26___mcc_h4 _dafny.Sequence = _source2.Get_().(D7_DC20).Cf31
		_ = _26___mcc_h4
		var _27___mcc_h5 bool = _source2.Get_().(D7_DC20).Cf32
		_ = _27___mcc_h5
		var _28_cf32 bool = _27___mcc_h5
		_ = _28_cf32
		var _29_cf31 _dafny.Sequence = _26___mcc_h4
		_ = _29_cf31
		if _28_cf32 {
			return Companion_D0_.Create_DC1_(_28_cf32, _dafny.SeqOf(_29_cf31), (Companion_D7_.Create_DC20_(_29_cf31, _28_cf32)).Dtor_cf31())
		} else {
			return Companion_D0_.Create_DC1_(true, _dafny.SeqOf(_29_cf31, _29_cf31, _29_cf31, _29_cf31, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-869))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))), _29_cf31)
		}
	} else if _source2.Is_DC18() {
		var _31___mcc_h6 _dafny.Map = _source2.Get_().(D7_DC18).Cf26
		_ = _31___mcc_h6
		var _32_cf26 _dafny.Map = _31___mcc_h6
		_ = _32_cf26
		return Companion_D0_.Create_DC1_(true, _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_33_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))), _dafny.UnicodeSeqOfUtf8Bytes("narkygd"))
	} else {
		var _34___mcc_h7 D7 = _source2.Get_().(D7_DC21).Cf33
		_ = _34___mcc_h7
		var _35_cf33 D7 = _34___mcc_h7
		_ = _35_cf33
		return Companion_D0_.Create_DC1_(!(false), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_36_i2 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("u")
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_37_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(256))).Uint32(), func(coer6 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_38_i0 _dafny.Int) D0 {
				return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwq")).Cardinality()))
			})), _dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(694))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(559))).Uint32(), func(coer7 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_39_i1 _dafny.Int) D0 {
				return Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
			})))
		}
		return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(865))).Uint32(), func(coer8 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_40_i2 _dafny.Int) D0 {
			return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnmlywcha")).Cardinality()))
		})))
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-414))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_41_i3 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(Companion_D0_.Create_DC0_((_dafny.SetOf(false, false)).Cardinality()), Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgl")).Cardinality())))
	})))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	return !(false) || (!(!(true)))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(583))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_42_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), _dafny.UnicodeSeqOfUtf8Bytes("h"))
}
func (_static *CompanionStruct_Default___) Fm22(p0 D2, p1 D0, p2 bool, globalState *GlobalState) _dafny.Int {
	return ((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('m'))), Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('p')))), _dafny.SeqOf(Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('a'))))))).Difference(_dafny.MultiSetOf(Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('f'))), Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('k'))), Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('p'))), Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('e'), _dafny.CodePoint('t'), _dafny.CodePoint('j'))), Companion_D14_.Create_DC35_(_dafny.MultiSetOf(_dafny.CodePoint('b'), _dafny.CodePoint('w')))))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source3 D9 = Companion_D9_.Create_DC27_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-433), _dafny.IntOfInt64(600))).Cardinality()), false)
	_ = _source3
	if _source3.Is_DC26() {
		var _43___mcc_h0 bool = _source3.Get_().(D9_DC26).Cf39
		_ = _43___mcc_h0
		var _44___mcc_h1 _dafny.CodePoint = _source3.Get_().(D9_DC26).Cf40
		_ = _44___mcc_h1
		var _45___mcc_h2 bool = _source3.Get_().(D9_DC26).Cf41
		_ = _45___mcc_h2
		var _46_cf41 bool = _45___mcc_h2
		_ = _46_cf41
		var _47_cf40 _dafny.CodePoint = _44___mcc_h1
		_ = _47_cf40
		var _48_cf39 bool = _43___mcc_h0
		_ = _48_cf39
		return _47_cf40
	} else if _source3.Is_DC27() {
		var _49___mcc_h3 bool = _source3.Get_().(D9_DC27).Cf42
		_ = _49___mcc_h3
		var _50___mcc_h4 _dafny.Int = _source3.Get_().(D9_DC27).Cf43
		_ = _50___mcc_h4
		var _51___mcc_h5 bool = _source3.Get_().(D9_DC27).Cf44
		_ = _51___mcc_h5
		var _52_cf44 bool = _51___mcc_h5
		_ = _52_cf44
		var _53_cf43 _dafny.Int = _50___mcc_h4
		_ = _53_cf43
		var _54_cf42 bool = _49___mcc_h3
		_ = _54_cf42
		return _dafny.CodePoint('a')
	} else {
		var _55___mcc_h6 _dafny.Map = _source3.Get_().(D9_DC25).Cf38
		_ = _55___mcc_h6
		var _56_cf38 _dafny.Map = _55___mcc_h6
		_ = _56_cf38
		return _dafny.CodePoint('e')
	}
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D9_.Create_DC26_(true, _dafny.CodePoint('a'), true)).Dtor_cf40(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
	})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-739), _dafny.IntOfInt64(362), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(600), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(612))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_57_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(801)
	})))).Cardinality(), _dafny.IntOfInt64(708), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(true), !(true))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hbr")).Cardinality())))).Intersection(((Companion_D12_.Create_DC31_(_dafny.MultiSetOf(_dafny.IntOfInt64(108)))).Dtor_cf48()).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(417))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 bool, globalState *GlobalState) D1 {
	var _source4 D4 = Companion_D4_.Create_DC12_(_dafny.IntOfInt64(642))
	_ = _source4
	if _source4.Is_DC11() {
		return Companion_D1_.Create_DC5_(!(false), _dafny.UnicodeSeqOfUtf8Bytes("hknqq"), _dafny.SeqOf(!(false), false))
	} else if _source4.Is_DC12() {
		var _58___mcc_h0 _dafny.Int = _source4.Get_().(D4_DC12).Cf18
		_ = _58___mcc_h0
		var _59_cf18 _dafny.Int = _58___mcc_h0
		_ = _59_cf18
		return Companion_D1_.Create_DC5_(true, _dafny.UnicodeSeqOfUtf8Bytes("tdv"), _dafny.SeqOf(true))
	} else {
		var _60___mcc_h1 _dafny.Set = _source4.Get_().(D4_DC10).Cf17
		_ = _60___mcc_h1
		var _61_cf17 _dafny.Set = _60___mcc_h1
		_ = _61_cf17
		return Companion_D1_.Create_DC5_(true, _dafny.UnicodeSeqOfUtf8Bytes("spg"), _dafny.SeqOf(false))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(true) || (false)))
}
func (_static *CompanionStruct_Default___) Fm28(globalState *GlobalState) D0 {
	if (_dafny.IntOfInt64(497)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(984))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_62_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()) > 0 {
		if !(false) {
			return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(true))
		} else {
			return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_(!(false), _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_63_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(958))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_64_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))), _dafny.UnicodeSeqOfUtf8Bytes("fwyar"))))
		}
	} else {
		return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(296)))
	}
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfptdctot")).Cardinality())).Cmp(_dafny.IntOfInt64(855)) == 0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(272)))), (_dafny.IntOfInt64(987)).Times(_dafny.IntOfUint32((_dafny.SeqOf(false, true, true, !(false), true)).Cardinality())), !(!(true) || (false)))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(93)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_65_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(753), _dafny.IntOfInt64(220))).Cardinality())
	})))
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-89))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_66_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	})), _dafny.UnicodeSeqOfUtf8Bytes("mijsmoem"))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D2_.Create_DC7_(!(true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(430))).Cardinality()))), _dafny.IntOfInt64(710), true)).Dtor_cf13())).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(990), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), false))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(false), true, false, true)).Union((_dafny.SetOf(false, !(false), !(true), true)).Difference((_dafny.SetOf(false))))
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.MultiSet {
	if (_dafny.SetOf(_dafny.IntOfInt64(14), (_dafny.Zero).Minus((_dafny.MultiSetOf(false, false, !(false))).Cardinality()))).IsSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("amok")).Cardinality()))) {
		return (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-406), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality())))).Difference(_dafny.MultiSetOf(_dafny.SeqOf((func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(299), _dafny.IntOfInt64(729))); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _67_v0 _dafny.Int
				_67_v0 = interface{}(_compr_0).(_dafny.Int)
				if ((_dafny.IntOfInt64(299)).Cmp(_67_v0) <= 0) && ((_67_v0).Cmp(_dafny.IntOfInt64(729)) < 0) {
					_coll0.Add((_67_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), false)).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mwcb")).Cardinality()))).Cardinality())).Cardinality())), _dafny.CodePoint('u'))
				}
			}
			return _coll0.ToMap()
		}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_68_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(962)
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(589))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_69_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(159)
		})), _dafny.SeqOf((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_70_i2 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()
			}))).Cardinality()))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _71_v1 _dafny.Int
				_71_v1 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}(func(_70_i2 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()
				}))).Cardinality())), _71_v1) {
					_coll1.Add(Companion_Default___.SafeDivisionInt(_71_v1, _dafny.IntOfInt64(672)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(228), (func() _dafny.Set {
						var _coll2 = _dafny.NewBuilder()
						_ = _coll2
						for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(228), _dafny.IntOfInt64(681))); ; {
							_compr_2, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _72_v2 _dafny.Int
							_72_v2 = interface{}(_compr_2).(_dafny.Int)
							if ((_dafny.IntOfInt64(228)).Cmp(_72_v2) <= 0) && ((_72_v2).Cmp(_dafny.IntOfInt64(681)) < 0) {
								_coll2.Add(Companion_Default___.SafeModuloInt(_72_v2, _dafny.IntOfInt64(70)))
							}
						}
						return _coll2.ToSet()
					}()).Cardinality())).Cardinality()))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality())))
	} else {
		return (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("juf")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}(func(_73_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(919), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uise")).Cardinality()), _dafny.IntOfInt64(23))).Cardinality()))).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(739), _dafny.IntOfInt64(490), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(738))).Cardinality())))).Union(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(487))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_74_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC12_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-151), (_dafny.SetOf(_dafny.IntOfInt64(788))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(-655)))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _dafny.IntOfInt64(238))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_75_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-886)
	})))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(true), !(false))).Union(_dafny.MultiSetOf(false, true, false))).Intersection((_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 D4, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC20_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-531))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_76_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_77_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))), !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(979))).Cardinality()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))) < 0))
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 bool, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC26_(!(_dafny.SetOf(true, true)).Contains(!(true)), _dafny.CodePoint('s'), !(!(!(true))))
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(Companion_D0_.Create_DC1_(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-783))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_78_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("jtdyx")
	})), _dafny.UnicodeSeqOfUtf8Bytes("pj")), Companion_D0_.Create_DC1_(!(false), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("woy"), _dafny.UnicodeSeqOfUtf8Bytes("dwuebhwk")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(321))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_79_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))), Companion_D0_.Create_DC1_(false, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ue")), _dafny.UnicodeSeqOfUtf8Bytes("wo")))).Intersection(_dafny.MultiSetOf(Companion_D0_.Create_DC1_(!(!(true)), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wg"), _dafny.UnicodeSeqOfUtf8Bytes("jwmnfd")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(475))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_80_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))), Companion_D0_.Create_DC1_(false, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jmo")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}(func(_81_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))), Companion_D0_.Create_DC1_(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer30(arg30)
		}
	}(func(_82_i4 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("boc")
	})), _dafny.UnicodeSeqOfUtf8Bytes("cihhuxg"))))
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false)), _dafny.MultiSetOf(false, false), _dafny.MultiSetOf(!(!(true))))).Intersection(_dafny.SetOf(_dafny.MultiSetOf(!(true))))).Intersection((_dafny.SetOf(_dafny.MultiSetOf(true, true, true))).Union(_dafny.SetOf(_dafny.MultiSetOf(true, false, true, false), _dafny.MultiSetOf(true))))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	if (func() bool {
		if false {
			return false
		}
		return false
	})() {
		return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), !(!(true))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	} else {
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))))
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _83_v0 _dafny.Array
	_ = _83_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
	_ = _nw0
	_83_v0 = _nw0
	var _84_v1 _dafny.Array
	_ = _84_v1
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
	_ = _nw1
	_84_v1 = _nw1
	var _85_v2 _dafny.Sequence
	_ = _85_v2
	_85_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fno")
	var _86_v3 _dafny.Int
	_ = _86_v3
	_86_v3 = _dafny.IntOfInt64(636)
	var _87_v4 _dafny.CodePoint
	_ = _87_v4
	_87_v4 = _dafny.CodePoint('a')
	var _88_v5 _dafny.Set
	_ = _88_v5
	_88_v5 = _dafny.SetOf((_85_v2).Select((Companion_Default___.SafeIndex(_86_v3, _dafny.IntOfUint32((_85_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), _87_v4)
	var _89_v6 _dafny.Map
	_ = _89_v6
	_89_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v5, _86_v3)
	var _90_v7 _dafny.Sequence
	_ = _90_v7
	_90_v7 = _dafny.SeqOf(_89_v6, _89_v6)
	var _91_v8 bool
	_ = _91_v8
	_91_v8 = false
	var _92_v9 _dafny.MultiSet
	_ = _92_v9
	_92_v9 = _dafny.MultiSetOf(_91_v8)
	var _93_v10 _dafny.Map
	_ = _93_v10
	_93_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v8, _91_v8)
	var _94_v11 _dafny.MultiSet
	_ = _94_v11
	_94_v11 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v8, _91_v8), _93_v10, _93_v10)
	var _95_v12 _dafny.Map
	_ = _95_v12
	_95_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, _92_v9)
	var _96_v13 _dafny.Set
	_ = _96_v13
	_96_v13 = _dafny.SetOf(_91_v8)
	var _97_globalState *GlobalState
	_ = _97_globalState
	var _nw2 *GlobalState = New_GlobalState_()
	_ = _nw2
	_nw2.Ctor__(_dafny.CodePoint('b'), _83_v0, _84_v1, true, (_90_v7).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
		if (_92_v9).Contains(true) {
			return (_92_v9).Multiplicity(true)
		}
		return (_94_v11).Cardinality()
	})(), _dafny.IntOfUint32((_90_v7).Cardinality()))).Uint32()).(_dafny.Map), true, _dafny.IntOfInt64(-209), (_95_v12).Merge(_95_v12), _dafny.Companion_Sequence_.Concatenate(_85_v2, _85_v2), true, true, _96_v13, _85_v2, false)
	_97_globalState = _nw2
	_83_v0 = _83_v0
	var _98_v14 _dafny.Array
	_ = _98_v14
	var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
	_ = _nw3
	_98_v14 = _nw3
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _index0
	(_98_v14).ArraySet1(_86_v3, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _index1
	(_98_v14).ArraySet1(_86_v3, (_index1).Int())
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
	_ = _index2
	(_83_v0).ArraySet1(((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v2, _85_v2)).Cardinality())) != 0, (_index2).Int())
	var _99_v16 _dafny.Sequence
	_ = _99_v16
	_99_v16 = _dafny.SeqOf((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))
	var _100_v17 _dafny.Sequence
	_ = _100_v17
	_100_v17 = _dafny.SeqOf((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_99_v16).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _101_v15 _dafny.Int
			_101_v15 = interface{}(_compr_3).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_99_v16, _101_v15) {
				_coll3.Add((_101_v15).Times((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)), _91_v8)
			}
		}
		return _coll3.ToMap()
	}()).Cardinality())
	var _102_v18 _dafny.MultiSet
	_ = _102_v18
	_102_v18 = _dafny.MultiSetOf(_99_v16)
	var _103_v19 _dafny.Map
	_ = _103_v19
	_103_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(964), _86_v3)
	var _104_v20 _dafny.MultiSet
	_ = _104_v20
	_104_v20 = _dafny.MultiSetOf((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (_92_v9).Cardinality(), (func() _dafny.Int {
		if (_103_v19).Contains(_86_v3) {
			return (_103_v19).Get(_86_v3).(_dafny.Int)
		}
		return (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)
	})())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _index3
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
	_ = _index4
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _index5
	var _rhs0 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)), _86_v3)
	_ = _rhs0
	var _rhs1 bool = ((_104_v20).Union(_104_v20)).IsSubsetOf(_dafny.MultiSetOf((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
		if (_102_v18).Contains(_100_v17) {
			return (_102_v18).Multiplicity(_100_v17)
		}
		return _86_v3
	})(), _86_v3, _dafny.IntOfInt64(-296)))
	_ = _rhs1
	var _rhs2 bool = _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("qahncvdqm"), (_85_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_93_v10).Cardinality()), _dafny.IntOfUint32((_85_v2).Cardinality()))).Uint32()).(_dafny.CodePoint))
	_ = _rhs2
	var _rhs3 bool = ((_dafny.IntOfInt64(145)).Plus(_dafny.IntOfInt64(482))).Cmp(Companion_Default___.SafeModuloInt(_86_v3, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))) < 0
	_ = _rhs3
	var _rhs4 _dafny.Int = (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)
	_ = _rhs4
	var _lhs0 _dafny.Array = _98_v14
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _lhs1
	var _lhs2 *GlobalState = _97_globalState
	_ = _lhs2
	var _lhs3 _dafny.Array = _83_v0
	_ = _lhs3
	var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
	_ = _lhs4
	var _lhs5 _dafny.Array = _98_v14
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
	_ = _lhs6
	(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
	_91_v8 = _rhs1
	_lhs2.F9 = _rhs2
	(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
	(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
	var _105_v21 _dafny.Map
	_ = _105_v21
	_105_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v8, _91_v8)).Update((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)), _98_v14)
	_98_v14 = (func() _dafny.Array {
		if (_105_v21).Contains(_93_v10) {
			return (_105_v21).Get(_93_v10).(_dafny.Array)
		}
		return _98_v14
	})()
	_86_v3 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-954))
	var _106_v22 *C7
	_ = _106_v22
	var _nw4 *C7 = New_C7_()
	_ = _nw4
	_nw4.Ctor__()
	_106_v22 = _nw4
	var _107_v23 *C5
	_ = _107_v23
	var _nw5 *C5 = New_C5_()
	_ = _nw5
	_nw5.Ctor__()
	_107_v23 = _nw5
	_91_v8 = (func() bool {
		if !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_107_v23), _107_v23) {
			return _91_v8
		}
		return (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)
	})()
	var _hi0 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm20((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _91_v8, _91_v8, _86_v3, _97_globalState)).Cardinality())
	_ = _hi0
	for _108_i0 := ((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)).Times(_86_v3); _108_i0.Cmp(_hi0) < 0; _108_i0 = _108_i0.Plus(_dafny.One) {
		var _109_v24 _dafny.Array
		_ = _109_v24
		var _nwElement0_0 _dafny.CodePoint = _87_v4
		_ = _nwElement0_0
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(5))
		_ = _nw6
		(_nw6).ArraySet1CodePoint(_nwElement0_0, 0)
		(_nw6).ArraySet1CodePoint(_87_v4, 1)
		(_nw6).ArraySet1CodePoint(_87_v4, 2)
		(_nw6).ArraySet1CodePoint(_87_v4, 3)
		(_nw6).ArraySet1CodePoint(_87_v4, 4)
		_109_v24 = _nw6
		var _110_v25 _dafny.Map
		_ = _110_v25
		_110_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v24, _86_v3)
		_110_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v24, _108_i0)
		_93_v10 = (_93_v10).Update(_91_v8, (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
		var _111_v26 *C1
		_ = _111_v26
		var _nw7 *C1 = New_C1_()
		_ = _nw7
		_nw7.Ctor__()
		_111_v26 = _nw7
		_85_v2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_112_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_113_i1 _dafny.Int) _dafny.CodePoint {
				return _112_v4
			}
		})(_87_v4)))
	}
	var _114_v27 _dafny.Map
	_ = _114_v27
	_114_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer32(arg32)
		}
	}((func(_115_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_116_i2 _dafny.Int) _dafny.CodePoint {
			return _115_v4
		}
	})(_87_v4)))).Cardinality()))
	var _117_v28 D9
	_ = _117_v28
	_117_v28 = Companion_D9_.Create_DC25_(_114_v27)
	var _118_v29 _dafny.Map
	_ = _118_v29
	_118_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, _117_v28)
	var _rhs5 _dafny.Int = (_118_v29).Cardinality()
	_ = _rhs5
	var _rhs6 _dafny.Sequence = _85_v2
	_ = _rhs6
	_86_v3 = _rhs5
	_85_v2 = _rhs6
	var _119_v30 _dafny.Map
	_ = _119_v30
	_119_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _85_v2)
	var _hi1 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_119_v30).Contains(_dafny.IntOfInt64(288)) {
			return (_119_v30).Get(_dafny.IntOfInt64(288)).(_dafny.Sequence)
		}
		return _85_v2
	})()).Cardinality())
	_ = _hi1
	for _120_i3 := (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int); _120_i3.Cmp(_hi1) < 0; _120_i3 = _120_i3.Plus(_dafny.One) {
		var _121_v31 _dafny.MultiSet
		_ = _121_v31
		_121_v31 = _dafny.MultiSetOf(_87_v4, _87_v4)
		var _122_v32 D14
		_ = _122_v32
		_122_v32 = Companion_D14_.Create_DC35_(_121_v31)
		var _123_v33 _dafny.Set
		_ = _123_v33
		_123_v33 = _dafny.SetOf(_dafny.IntOfUint32((_85_v2).Cardinality()))
		var _124_v34 _dafny.Sequence
		_ = _124_v34
		_124_v34 = _dafny.SeqOf(Companion_Default___.Fm19((_123_v33).Cardinality(), (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _97_globalState))
		if !((_121_v31).IsSubsetOf((_122_v32).Dtor_cf53())) || (_dafny.Companion_Sequence_.IsPrefixOf(_124_v34, _124_v34)) {
			(_97_globalState).F5 = (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)
			_86_v3 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(85), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_85_v2).Cardinality()), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)))
			var _125_v35 _dafny.Set
			_ = _125_v35
			_125_v35 = _dafny.SetOf(_85_v2)
			var _126_v36 T0
			_ = _126_v36
			var _nw8 *C2 = New_C2_()
			_ = _nw8
			_nw8.Ctor__()
			_126_v36 = _nw8
			var _127_v37 D4
			_ = _127_v37
			_127_v37 = Companion_D4_.Create_DC12_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-645), (_92_v9).Cardinality()))
			var _128_v38 _dafny.Map
			_ = _128_v38
			_128_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(44), _125_v35)
			var _129_v39 _dafny.Map
			_ = _129_v39
			_129_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _125_v35)
			var _rhs7 _dafny.Set = ((func() _dafny.Set {
				if (_128_v38).Contains((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)) {
					return (_128_v38).Get((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)).(_dafny.Set)
				}
				return _125_v35
			})()).Intersection((_125_v35).Union((func() _dafny.Set {
				if (_129_v39).Contains(_91_v8) {
					return (_129_v39).Get(_91_v8).(_dafny.Set)
				}
				return _125_v35
			})()))
			_ = _rhs7
			var _rhs8 T0 = _126_v36
			_ = _rhs8
			var _rhs9 D4 = _127_v37
			_ = _rhs9
			var _rhs10 bool = _91_v8
			_ = _rhs10
			var _lhs7 *GlobalState = _97_globalState
			_ = _lhs7
			_125_v35 = _rhs7
			_126_v36 = _rhs8
			_127_v37 = _rhs9
			_lhs7.F5 = _rhs10
			var _130_v40 *C1
			_ = _130_v40
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__()
			_130_v40 = _nw9
			var _131_v41 _dafny.Map
			_ = _131_v41
			_131_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))
			var _132_v42 _dafny.Map
			_ = _132_v42
			_132_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_100_v17).Select((Companion_Default___.SafeIndex((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_100_v17).Cardinality()))).Uint32()).(_dafny.Int), _130_v40)
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
			_ = _index6
			var _rhs11 *C1 = (func() *C1 {
				if (_132_v42).Contains(_120_i3) {
					return (_132_v42).Get(_120_i3).(*C1)
				}
				return _130_v40
			})()
			_ = _rhs11
			var _rhs12 bool = (_dafny.MultiSetOf(_86_v3)).IsProperSubsetOf(_dafny.MultiSetOf(_120_i3, (_99_v16).Select((Companion_Default___.SafeIndex(_120_i3, _dafny.IntOfUint32((_99_v16).Cardinality()))).Uint32()).(_dafny.Int)))
			_ = _rhs12
			var _rhs13 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_124_v34, _124_v34)
			_ = _rhs13
			var _rhs14 _dafny.Map = _103_v19
			_ = _rhs14
			var _rhs15 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_93_v10).Contains((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)) {
					return (_93_v10).Get((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)).(bool)
				}
				return _91_v8
			})(), _dafny.IntOfUint32((_85_v2).Cardinality()))).Merge((_131_v41).Merge(_131_v41))
			_ = _rhs15
			var _lhs8 _dafny.Array = _83_v0
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
			_ = _lhs9
			_130_v40 = _rhs11
			(_lhs8).ArraySet1(_rhs12, (_lhs9).Int())
			_124_v34 = _rhs13
			_103_v19 = _rhs14
			_131_v41 = _rhs15
			var _133_v43 *C6
			_ = _133_v43
			var _nw10 *C6 = New_C6_()
			_ = _nw10
			_nw10.Ctor__()
			_133_v43 = _nw10
		} else {
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
			_ = _index7
			(_98_v14).ArraySet1((((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)).Times(_86_v3)).Plus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_120_i3, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)))), (_index7).Int())
			var _134_v44 _dafny.Map
			_ = _134_v44
			_134_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v14, (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
			var _135_v45 _dafny.Array
			_ = _135_v45
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
			_ = _nw11
			_135_v45 = _nw11
			var _136_v46 D1
			_ = _136_v46
			_136_v46 = Companion_D1_.Create_DC5_((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _85_v2, _124_v34)
			var _137_v47 _dafny.MultiSet
			_ = _137_v47
			_137_v47 = _dafny.MultiSetOf(Companion_D1_.Create_DC5_((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(700))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_138_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_139_i4 _dafny.Int) _dafny.CodePoint {
					return _138_v4
				}
			})(_87_v4))), _124_v34), _136_v46)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_135_v45), 0))
			_ = _index8
			(_135_v45).ArraySet1(_137_v47, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_135_v45), 0))
			_ = _index9
			var _rhs16 _dafny.Map = _134_v44
			_ = _rhs16
			var _rhs17 _dafny.MultiSet = _dafny.MultiSetOf(_136_v46, _136_v46, Companion_Default___.Fm26(true, _91_v8, _97_globalState))
			_ = _rhs17
			var _rhs18 D9 = _117_v28
			_ = _rhs18
			var _lhs10 _dafny.Array = _135_v45
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_135_v45), 0))
			_ = _lhs11
			_134_v44 = _rhs16
			(_lhs10).ArraySet1(_rhs17, (_lhs11).Int())
			_117_v28 = _rhs18
			(_97_globalState).F5 = !(_91_v8)
			var _140_v48 _dafny.Map
			_ = _140_v48
			_140_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _91_v8)
			var _141_v49 D7
			_ = _141_v49
			_141_v49 = Companion_D7_.Create_DC18_(_140_v48)
			var _pat_let_tv0 = _91_v8
			_ = _pat_let_tv0
			_141_v49 = func(_pat_let0_0 D7) D7 {
				return func(_142_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let1_0 _dafny.Map) D7 {
						return func(_143_dt__update_hcf26_h0 _dafny.Map) D7 {
							return Companion_D7_.Create_DC18_(_143_dt__update_hcf26_h0)
						}(_pat_let1_0)
					}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(506), _pat_let_tv0))
				}(_pat_let0_0)
			}(_141_v49)
			(_97_globalState).F13 = Companion_Default___.Fm6(_97_globalState)
		}
		var _144_v50 T1
		_ = _144_v50
		var _nw12 *C3 = New_C3_()
		_ = _nw12
		_nw12.Ctor__(!((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)), _120_i3)
		_144_v50 = _nw12
		var _145_v51 _dafny.Sequence
		_ = _145_v51
		_145_v51 = _dafny.SeqOf(_dafny.MultiSetOf(Companion_Default___.Fm11(_91_v8, _97_globalState), _dafny.IntOfUint32((_124_v34).Cardinality())), _104_v20)
		(_107_v23).M3(_86_v3, _144_v50, (_144_v50).Fm0(_86_v3, _87_v4, _97_globalState), (_145_v51).Select((Companion_Default___.SafeIndex((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_145_v51).Cardinality()))).Uint32()).(_dafny.MultiSet), _97_globalState)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
		_ = _index10
		(_83_v0).ArraySet1((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
		_ = _index11
		(_98_v14).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_86_v3)), (_index11).Int())
	}
	var _146_v52 _dafny.Map
	_ = _146_v52
	_146_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
	var _147_v53 D7
	_ = _147_v53
	_147_v53 = Companion_D7_.Create_DC18_(_146_v52)
	var _148_v54 _dafny.Sequence
	_ = _148_v54
	_148_v54 = _dafny.SeqOf((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
	var _149_v55 _dafny.Sequence
	_ = _149_v55
	_149_v55 = _dafny.SeqOf(_85_v2, _85_v2)
	var _150_v56 D0
	_ = _150_v56
	_150_v56 = Companion_D0_.Create_DC1_(_91_v8, _149_v55, _85_v2)
	var _151_v57 D0
	_ = _151_v57
	_151_v57 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(_150_v56))
	var _pat_let_tv1 = _147_v53
	_ = _pat_let_tv1
	var _pat_let_tv2 = _147_v53
	_ = _pat_let_tv2
	var _pat_let_tv3 = _146_v52
	_ = _pat_let_tv3
	var _pat_let_tv4 = _146_v52
	_ = _pat_let_tv4
	var _pat_let_tv5 = _147_v53
	_ = _pat_let_tv5
	var _rhs19 _dafny.Map = _93_v10
	_ = _rhs19
	var _rhs20 _dafny.Int = (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)
	_ = _rhs20
	var _rhs21 bool = (_148_v54).Select((Companion_Default___.SafeIndex(_86_v3, _dafny.IntOfUint32((_148_v54).Cardinality()))).Uint32()).(bool)
	_ = _rhs21
	var _rhs22 _dafny.Sequence = _85_v2
	_ = _rhs22
	var _rhs23 D7 = func(_source5 D0) D7 {
		if _source5.Is_DC1() {
			var _152___mcc_h0 bool = _source5.Get_().(D0_DC1).Cf1
			_ = _152___mcc_h0
			var _153___mcc_h1 _dafny.Sequence = _source5.Get_().(D0_DC1).Cf2
			_ = _153___mcc_h1
			var _154___mcc_h2 _dafny.Sequence = _source5.Get_().(D0_DC1).Cf3
			_ = _154___mcc_h2
			var _155_cf3 _dafny.Sequence = _154___mcc_h2
			_ = _155_cf3
			var _156_cf2 _dafny.Sequence = _153___mcc_h1
			_ = _156_cf2
			var _157_cf1 bool = _152___mcc_h0
			_ = _157_cf1
			return _pat_let_tv1
		} else if _source5.Is_DC2() {
			var _158___mcc_h3 bool = _source5.Get_().(D0_DC2).Cf4
			_ = _158___mcc_h3
			var _159_cf4 bool = _158___mcc_h3
			_ = _159_cf4
			return _pat_let_tv2
		} else if _source5.Is_DC0() {
			var _160___mcc_h4 _dafny.Int = _source5.Get_().(D0_DC0).Cf0
			_ = _160___mcc_h4
			var _161_cf0 _dafny.Int = _160___mcc_h4
			_ = _161_cf0
			var _162_dt__update__tmp_h1 D7 = Companion_D7_.Create_DC18_(_pat_let_tv3)
			_ = _162_dt__update__tmp_h1
			var _163_dt__update_hcf26_h1 _dafny.Map = _pat_let_tv4
			_ = _163_dt__update_hcf26_h1
			return Companion_D7_.Create_DC18_(_163_dt__update_hcf26_h1)
		} else {
			var _164___mcc_h5 D0 = _source5.Get_().(D0_DC3).Cf5
			_ = _164___mcc_h5
			var _165_cf5 D0 = _164___mcc_h5
			_ = _165_cf5
			return Companion_D7_.Create_DC18_((_pat_let_tv5).Dtor_cf26())
		}
	}(_151_v57)
	_ = _rhs23
	var _lhs12 *GlobalState = _97_globalState
	_ = _lhs12
	_93_v10 = _rhs19
	_86_v3 = _rhs20
	_lhs12.F13 = _rhs21
	_85_v2 = _rhs22
	_147_v53 = _rhs23
	var _166_v58 _dafny.Map
	_ = _166_v58
	_166_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v8, _86_v3)
	var _167_v59 D9
	_ = _167_v59
	_167_v59 = Companion_D9_.Create_DC26_((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _87_v4, _91_v8)
	var _168_v60 _dafny.Array
	_ = _168_v60
	var _nwElement0_1 _dafny.CodePoint = _87_v4
	_ = _nwElement0_1
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(29))
	_ = _nw13
	(_nw13).ArraySet1CodePoint(_nwElement0_1, 0)
	(_nw13).ArraySet1CodePoint(_dafny.CodePoint('y'), 1)
	(_nw13).ArraySet1CodePoint(_87_v4, 2)
	(_nw13).ArraySet1CodePoint(_87_v4, 3)
	(_nw13).ArraySet1CodePoint(Companion_Default___.Fm23((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v4, _93_v10), false, (func() _dafny.Int {
		if (_166_v58).Contains(_91_v8) {
			return (_166_v58).Get(_91_v8).(_dafny.Int)
		}
		return (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)
	})(), _97_globalState), 4)
	(_nw13).ArraySet1CodePoint(_87_v4, 5)
	(_nw13).ArraySet1CodePoint(_87_v4, 6)
	(_nw13).ArraySet1CodePoint(_87_v4, 7)
	(_nw13).ArraySet1CodePoint(_87_v4, 8)
	(_nw13).ArraySet1CodePoint(_87_v4, 9)
	(_nw13).ArraySet1CodePoint(_87_v4, 10)
	(_nw13).ArraySet1CodePoint(_87_v4, 11)
	(_nw13).ArraySet1CodePoint(_87_v4, 12)
	(_nw13).ArraySet1CodePoint(_87_v4, 13)
	(_nw13).ArraySet1CodePoint((_167_v59).Dtor_cf40(), 14)
	(_nw13).ArraySet1CodePoint(_87_v4, 15)
	(_nw13).ArraySet1CodePoint(_87_v4, 16)
	(_nw13).ArraySet1CodePoint(_87_v4, 17)
	(_nw13).ArraySet1CodePoint(_87_v4, 18)
	(_nw13).ArraySet1CodePoint(_87_v4, 19)
	(_nw13).ArraySet1CodePoint(_87_v4, 20)
	(_nw13).ArraySet1CodePoint(_87_v4, 21)
	(_nw13).ArraySet1CodePoint(_87_v4, 22)
	(_nw13).ArraySet1CodePoint(_87_v4, 23)
	(_nw13).ArraySet1CodePoint(_87_v4, 24)
	(_nw13).ArraySet1CodePoint(_87_v4, 25)
	(_nw13).ArraySet1CodePoint(_87_v4, 26)
	(_nw13).ArraySet1CodePoint(_87_v4, 27)
	(_nw13).ArraySet1CodePoint(_87_v4, 28)
	_168_v60 = _nw13
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_168_v60), 0))
	_ = _index12
	(_168_v60).ArraySet1CodePoint(_87_v4, (_index12).Int())
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_168_v60), 0))
	_ = _index13
	(_168_v60).ArraySet1CodePoint(_87_v4, (_index13).Int())
	var _169_v61 T1
	_ = _169_v61
	var _nw14 *C5 = New_C5_()
	_ = _nw14
	_nw14.Ctor__()
	_169_v61 = _nw14
	var _170_v62 _dafny.Map
	_ = _170_v62
	_170_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _169_v61)
	var _171_v63 _dafny.Set
	_ = _171_v63
	_171_v63 = _dafny.SetOf(_170_v62, _170_v62)
	var _172_i5 _dafny.Int
	_ = _172_i5
	_172_i5 = _dafny.Zero
	{
		for ((func() _dafny.Set {
			if (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool) {
				return _171_v63
			}
			return _171_v63
		})()).IsDisjointFrom(_171_v63) {
			{
				if (_172_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_172_i5 = (_172_i5).Plus(_dafny.One)
				_86_v3 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(357), _86_v3, _86_v3, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))).Cardinality()))))).Plus((Companion_D0_.Create_DC0_((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))).Dtor_cf0())
				(_97_globalState).F13 = (func() bool {
					if (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool) {
						return _91_v8
					}
					return _91_v8
				})()
				var _173_v64 _dafny.Map
				_ = _173_v64
				_173_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, (func() _dafny.Array {
					if (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool) {
						return _98_v14
					}
					return _98_v14
				})())
				var _174_v65 D11
				_ = _174_v65
				_174_v65 = Companion_D11_.Create_DC30_(false)
				var _175_v66 _dafny.Map
				_ = _175_v66
				_175_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(465))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}((func(_176_v60 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_177_i6 _dafny.Int) _dafny.CodePoint {
						return (_176_v60).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_176_v60), 0))).Int())
					}
				})(_168_v60)))).Cardinality())), _174_v65)
				var _178_v67 D0
				_ = _178_v67
				_178_v67 = Companion_D0_.Create_DC2_((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
				_ = _index14
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
				_ = _index15
				var _rhs24 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if _91_v8 {
						return (_175_v66).Cardinality()
					}
					return _86_v3
				})(), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))
				_ = _rhs24
				var _rhs25 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _98_v14)
				_ = _rhs25
				var _rhs26 T1 = _169_v61
				_ = _rhs26
				var _rhs27 _dafny.Int = (_107_v23).Fm5(_178_v67, _97_globalState)
				_ = _rhs27
				var _rhs28 bool = !(false)
				_ = _rhs28
				var _lhs13 _dafny.Array = _98_v14
				_ = _lhs13
				var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
				_ = _lhs14
				var _lhs15 _dafny.Array = _98_v14
				_ = _lhs15
				var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))
				_ = _lhs16
				var _lhs17 *GlobalState = _97_globalState
				_ = _lhs17
				(_lhs13).ArraySet1(_rhs24, (_lhs14).Int())
				_173_v64 = _rhs25
				_169_v61 = _rhs26
				(_lhs15).ArraySet1(_rhs27, (_lhs16).Int())
				_lhs17.F13 = _rhs28
				var _179_v68 *C1
				_ = _179_v68
				var _nw15 *C1 = New_C1_()
				_ = _nw15
				_nw15.Ctor__()
				_179_v68 = _nw15
				var _180_v69 _dafny.Set
				_ = _180_v69
				_180_v69 = _dafny.SetOf(_179_v68, _179_v68)
				(_97_globalState).F3 = ((_180_v69).Intersection((_dafny.SetOf(_179_v68)))).Equals((func() _dafny.Set {
					if Companion_Default___.Fm6(_97_globalState) {
						return _180_v69
					}
					return _180_v69
				})())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if _91_v8 {
		(_97_globalState).F13 = !(Companion_Default___.Fm2((func() _dafny.CodePoint {
			if _91_v8 {
				return _dafny.CodePoint('d')
			}
			return _87_v4
		})(), _151_v57, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _97_globalState))
		var _181_v70 _dafny.Set
		_ = _181_v70
		_181_v70 = _dafny.SetOf((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))
		var _182_v71 _dafny.Map
		_ = _182_v71
		_182_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)) == (true)), (_181_v70).Union(_181_v70))
		var _183_v72 D4
		_ = _183_v72
		_183_v72 = Companion_D4_.Create_DC10_(_181_v70)
		_182_v71 = (_182_v71).Update((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), (_183_v72).Dtor_cf17())
		(_97_globalState).F13 = _91_v8
		var _184_v73 _dafny.Map
		_ = _184_v73
		_184_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v14, _85_v2)
		var _185_v74 _dafny.Array
		_ = _185_v74
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw16
		_185_v74 = _nw16
		var _186_v75 _dafny.Array
		_ = _186_v75
		var _nw17 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw17
		_186_v75 = _nw17
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_185_v74), 0))
		_ = _index16
		(_185_v74).ArraySet1(_186_v75, (_index16).Int())
		var _187_v76 _dafny.Sequence
		_ = _187_v76
		_187_v76 = _dafny.SeqOf(_83_v0, _83_v0, _83_v0, _83_v0, _83_v0)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_185_v74), 0))
		_ = _index17
		var _rhs29 _dafny.Map = ((_184_v73).Merge(_184_v73)).Merge((func() _dafny.Map {
			if false {
				return (_184_v73).Update(_98_v14, _85_v2)
			}
			return _184_v73
		})())
		_ = _rhs29
		var _rhs30 bool = false
		_ = _rhs30
		var _rhs31 _dafny.Array = _186_v75
		_ = _rhs31
		var _rhs32 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_148_v54, _148_v54), _dafny.Companion_Sequence_.Contains(_148_v54, _91_v8))
		_ = _rhs32
		var _rhs33 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_187_v76, _187_v76), _187_v76)
		_ = _rhs33
		var _lhs18 *GlobalState = _97_globalState
		_ = _lhs18
		var _lhs19 _dafny.Array = _185_v74
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_185_v74), 0))
		_ = _lhs20
		var _lhs21 *GlobalState = _97_globalState
		_ = _lhs21
		_184_v73 = _rhs29
		_lhs18.F9 = _rhs30
		(_lhs19).ArraySet1(_rhs31, (_lhs20).Int())
		_lhs21.F13 = _rhs32
		_187_v76 = _rhs33
		(_107_v23).M3((_86_v3).Plus(_86_v3), _169_v61, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (Companion_Default___.Fm16(_97_globalState)).Difference(_104_v20), _97_globalState)
	} else {
		_148_v54 = _dafny.SeqOf((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
		var _188_v77 _dafny.Map
		_ = _188_v77
		_188_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _169_v61)
		var _189_v78 *C2
		_ = _189_v78
		var _nw18 *C2 = New_C2_()
		_ = _nw18
		_nw18.Ctor__()
		_189_v78 = _nw18
		var _190_v79 _dafny.Map
		_ = _190_v79
		_190_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _189_v78)
		(_107_v23).M3(Companion_Default___.SafeModuloInt((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _86_v3), (func() T1 {
			if (_188_v77).Contains(_86_v3) {
				return (_188_v77).Get(_86_v3).(T1)
			}
			return _169_v61
		})(), (_dafny.MultiSetOf(_190_v79)).Cardinality(), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), (_146_v52).Cardinality())).Cardinality(), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_85_v2).Cardinality()), (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)), _97_globalState)
		var _191_v81 _dafny.Sequence
		_ = _191_v81
		_191_v81 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, true), _146_v52, func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-331), _dafny.IntOfInt64(-333))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _192_v80 _dafny.Int
				_192_v80 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(-331)).Cmp(_192_v80) <= 0) && ((_192_v80).Cmp(_dafny.IntOfInt64(-333)) < 0) {
					_coll4.Add(Companion_Default___.SafeModuloInt(_192_v80, (_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int)), (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
				}
			}
			return _coll4.ToMap()
		}(), _146_v52, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_99_v16).Select((Companion_Default___.SafeIndex((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_99_v16).Cardinality()))).Uint32()).(_dafny.Int), (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)))
		_86_v3 = (_dafny.Zero).Minus((func() _dafny.Int {
			if !(_91_v8) {
				return _dafny.IntOfUint32((_191_v81).Cardinality())
			}
			return (_96_v13).Cardinality()
		})())
		var _193_v82 _dafny.Map
		_ = _193_v82
		_193_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v8, _91_v8))
		var _194_v83 _dafny.Sequence
		_ = _194_v83
		_194_v83 = _dafny.SeqOf((func() _dafny.Map {
			if (_193_v82).Contains(_86_v3) {
				return (_193_v82).Get(_86_v3).(_dafny.Map)
			}
			return _93_v10
		})(), _93_v10)
		var _195_v84 _dafny.Sequence
		_ = _195_v84
		_195_v84 = _dafny.SeqOf(_194_v83, _194_v83)
		var _196_v85 _dafny.Array
		_ = _196_v85
		var _nwElement0_2 _dafny.MultiSet = (_94_v11).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool)), Companion_Default___.Abs(((_146_v52).Update((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _91_v8)).Cardinality()))
		_ = _nwElement0_2
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(16))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_2, 0)
		(_nw19).ArraySet1(_94_v11, 1)
		(_nw19).ArraySet1(_dafny.MultiSetOf(_93_v10), 2)
		(_nw19).ArraySet1(_94_v11, 3)
		(_nw19).ArraySet1(Companion_Default___.Fm44((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int), _86_v3, _91_v8, _97_globalState), 4)
		(_nw19).ArraySet1((_94_v11).Difference(_94_v11), 5)
		(_nw19).ArraySet1(_94_v11, 6)
		(_nw19).ArraySet1(_94_v11, 7)
		(_nw19).ArraySet1(_94_v11, 8)
		(_nw19).ArraySet1(_dafny.MultiSetFromSeq((_195_v84).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_98_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_98_v14), 0))).Int()).(_dafny.Int))), _dafny.IntOfUint32((_195_v84).Cardinality()))).Uint32()).(_dafny.Sequence)), 9)
		(_nw19).ArraySet1(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), false), _93_v10, _93_v10, _93_v10), 10)
		(_nw19).ArraySet1((func() _dafny.MultiSet {
			if (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool) {
				return _94_v11
			}
			return _94_v11
		})(), 11)
		(_nw19).ArraySet1(_94_v11, 12)
		(_nw19).ArraySet1(_94_v11, 13)
		(_nw19).ArraySet1((_94_v11).Update(_93_v10, Companion_Default___.Abs(_86_v3)), 14)
		(_nw19).ArraySet1(_94_v11, 15)
		_196_v85 = _nw19
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_196_v85), 0))
		_ = _index18
		(_196_v85).ArraySet1((_94_v11).Union(_94_v11), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_196_v85), 0))
		_ = _index19
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
		_ = _index20
		var _rhs34 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool) {
				return _86_v3
			}
			return _86_v3
		})(), _86_v3)
		_ = _rhs34
		var _rhs35 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool), _91_v8))
		_ = _rhs35
		var _rhs36 bool = !((_83_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))).Int()).(bool))
		_ = _rhs36
		var _rhs37 bool = _91_v8
		_ = _rhs37
		var _lhs22 _dafny.Array = _196_v85
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_196_v85), 0))
		_ = _lhs23
		var _lhs24 _dafny.Array = _83_v0
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_83_v0), 0))
		_ = _lhs25
		var _lhs26 *GlobalState = _97_globalState
		_ = _lhs26
		_86_v3 = _rhs34
		(_lhs22).ArraySet1(_rhs35, (_lhs23).Int())
		(_lhs24).ArraySet1(_rhs36, (_lhs25).Int())
		_lhs26.F9 = _rhs37
		_86_v3 = _86_v3
	}
	var _197_v86 T0
	_ = _197_v86
	var _nw20 *C2 = New_C2_()
	_ = _nw20
	_nw20.Ctor__()
	_197_v86 = _nw20
	_197_v86 = _197_v86
	var _198_v87 bool
	_ = _198_v87
	var _199_v88 _dafny.Int
	_ = _199_v88
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = (_106_v22).M0(_87_v4, _97_globalState)
	_198_v87 = _out0
	_199_v88 = _out1
	_dafny.Print((_83_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_86_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_v5).Equals(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('a'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_89_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('a')), _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_90_v7, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('a')), _dafny.IntOfInt64(636)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('a')), _dafny.IntOfInt64(636)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_91_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_92_v9).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v11).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_95_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(636), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v13).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_97_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState.F4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('a')), _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_97_globalState).F7()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(636), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_97_globalState).F11()).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_globalState).F12().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v14).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_99_v16, _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_100_v17, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v18).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(964), _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v20).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636), _dafny.IntOfInt64(636), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v21).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(90))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_117_v28).Dtor_cf38()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(90))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v29).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(954), Companion_D9_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(90))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(636), _dafny.UnicodeSeqOfUtf8Bytes("fno"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v52).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_v53).Dtor_cf26()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_148_v54, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_149_v55, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fno"), _dafny.UnicodeSeqOfUtf8Bytes("fno"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v56).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_150_v56).Dtor_cf2(), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fno"), _dafny.UnicodeSeqOfUtf8Bytes("fno"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v56).Dtor_cf3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v57).Dtor_cf5()).Dtor_cf5()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((((_151_v57).Dtor_cf5()).Dtor_cf5()).Dtor_cf2(), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fno"), _dafny.UnicodeSeqOfUtf8Bytes("fno"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v57).Dtor_cf5()).Dtor_cf5()).Dtor_cf3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v58).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v59).Dtor_cf39())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v59).Dtor_cf40())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v59).Dtor_cf41())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(23)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(24)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(25)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(26)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(27)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v60).ArrayGet1CodePoint((_dafny.IntOfInt64(28)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v62).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v63).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_172_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_198_v87)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_199_v88)
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
	Cf2 _dafny.Sequence
	Cf3 _dafny.Sequence
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Sequence, Cf3 _dafny.Sequence) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool) D0 {
	return D0{D0_DC2{Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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

type D0_DC3 struct {
	Cf5 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf5 D0) D0 {
	return D0{D0_DC3{Cf5}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf5() D0 {
	return _this.Get_().(D0_DC3).Cf5
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + data.Cf3.VerbatimString(true) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Equals(data2.Cf3)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf7 bool
	Cf8 _dafny.Sequence
	Cf9 _dafny.Sequence
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 bool, Cf8 _dafny.Sequence, Cf9 _dafny.Sequence) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf6 _dafny.Map
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Map) D1 {
	return D1{D1_DC4{Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + data.Cf8.VerbatimString(true) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Equals(data2.Cf9)
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf12 _dafny.Map
	Cf13 _dafny.Int
	Cf14 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool, Cf12 _dafny.Map, Cf13 _dafny.Int, Cf14 bool) D2 {
	return D2{D2_DC7{Cf11, Cf12, Cf13, Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf10 _dafny.MultiSet
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.MultiSet) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf15 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf15 D2) D2 {
	return D2{D2_DC8{Cf15}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, _dafny.EmptyMap, _dafny.Zero, false)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() bool {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf10() _dafny.MultiSet {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf15() D2 {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Equals(data2.Cf12) && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf16 _dafny.Set
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Set) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D3) Dtor_cf16() _dafny.Set {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf16.Equals(data2.Cf16)
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
	Cf18 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf18 _dafny.Int) D4 {
	return D4{D4_DC12{Cf18}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf17 _dafny.Set
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf17 _dafny.Set) D4 {
	return D4{D4_DC10{Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_()
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D4_DC10).Cf17
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
			return "D4.DC12" + "(" + _dafny.String(data.Cf18) + ")"
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
			_, ok := other.Get_().(D4_DC11)
			return ok
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
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

type D5_DC14 struct {
	Cf20 bool
	Cf21 bool
	Cf22 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 bool, Cf21 bool, Cf22 _dafny.Array) D5 {
	return D5{D5_DC14{Cf20, Cf21, Cf22}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf19 _dafny.Array
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf19 _dafny.Array) D5 {
	return D5{D5_DC13{Cf19}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC14).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf19 == data2.Cf19
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

type D6_DC16 struct {
	Cf24 _dafny.Int
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf24 _dafny.Int) D6 {
	return D6{D6_DC16{Cf24}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf23 _dafny.MultiSet
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.MultiSet) D6 {
	return D6{D6_DC15{Cf23}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC17 struct {
	Cf25 D6
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf25 D6) D6 {
	return D6{D6_DC17{Cf25}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.Zero)
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) Dtor_cf23() _dafny.MultiSet {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) Dtor_cf25() D6 {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D7_DC19 struct {
	Cf27 D1
	Cf28 _dafny.Int
	Cf29 bool
	Cf30 _dafny.Array
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf27 D1, Cf28 _dafny.Int, Cf29 bool, Cf30 _dafny.Array) D7 {
	return D7{D7_DC19{Cf27, Cf28, Cf29, Cf30}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf31 _dafny.Sequence
	Cf32 bool
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf31 _dafny.Sequence, Cf32 bool) D7 {
	return D7{D7_DC20{Cf31, Cf32}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf26 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf26 _dafny.Map) D7 {
	return D7{D7_DC18{Cf26}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC21 struct {
	Cf33 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf33 D7) D7 {
	return D7{D7_DC21{Cf33}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(Companion_D1_.Default(), _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D7) Dtor_cf27() D1 {
	return _this.Get_().(D7_DC19).Cf27
}

func (_this D7) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf28
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC19).Cf29
}

func (_this D7) Dtor_cf30() _dafny.Array {
	return _this.Get_().(D7_DC19).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) Dtor_cf32() bool {
	return _this.Get_().(D7_DC20).Cf32
}

func (_this D7) Dtor_cf26() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf26
}

func (_this D7) Dtor_cf33() D7 {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + data.Cf31.VerbatimString(true) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf27.Equals(data2.Cf27) && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf31.Equals(data2.Cf31) && data1.Cf32 == data2.Cf32
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf33.Equals(data2.Cf33)
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
	Cf35 bool
	Cf36 bool
	Cf37 _dafny.Int
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf35 bool, Cf36 bool, Cf37 _dafny.Int) D8 {
	return D8{D8_DC23{Cf35, Cf36, Cf37}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_() D8 {
	return D8{D8_DC24{}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf34 _dafny.Map
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf34 _dafny.Map) D8 {
	return D8{D8_DC22{Cf34}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC23).Cf35
}

func (_this D8) Dtor_cf36() bool {
	return _this.Get_().(D8_DC23).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf37
}

func (_this D8) Dtor_cf34() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf34
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D8_DC24:
		{
			_, ok := other.Get_().(D8_DC24)
			return ok
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D9_DC26 struct {
	Cf39 bool
	Cf40 _dafny.CodePoint
	Cf41 bool
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf39 bool, Cf40 _dafny.CodePoint, Cf41 bool) D9 {
	return D9{D9_DC26{Cf39, Cf40, Cf41}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf42 bool
	Cf43 _dafny.Int
	Cf44 bool
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf42 bool, Cf43 _dafny.Int, Cf44 bool) D9 {
	return D9{D9_DC27{Cf42, Cf43, Cf44}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC25 struct {
	Cf38 _dafny.Map
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf38 _dafny.Map) D9 {
	return D9{D9_DC25{Cf38}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_(false, _dafny.CodePoint('D'), false)
}

func (_this D9) Dtor_cf39() bool {
	return _this.Get_().(D9_DC26).Cf39
}

func (_this D9) Dtor_cf40() _dafny.CodePoint {
	return _this.Get_().(D9_DC26).Cf40
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC26).Cf41
}

func (_this D9) Dtor_cf42() bool {
	return _this.Get_().(D9_DC27).Cf42
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC27).Cf43
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC27).Cf44
}

func (_this D9) Dtor_cf38() _dafny.Map {
	return _this.Get_().(D9_DC25).Cf38
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44 == data2.Cf44
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
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

type D10_DC28 struct {
	Cf45 _dafny.Array
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf45 _dafny.Array) D10 {
	return D10{D10_DC28{Cf45}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf45() _dafny.Array {
	return _this.Get_().(D10_DC28).Cf45
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf45 == data2.Cf45
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

type D11_DC30 struct {
	Cf47 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf47 bool) D11 {
	return D11{D11_DC30{Cf47}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC29 struct {
	Cf46 _dafny.MultiSet
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf46 _dafny.MultiSet) D11 {
	return D11{D11_DC29{Cf46}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(false)
}

func (_this D11) Dtor_cf47() bool {
	return _this.Get_().(D11_DC30).Cf47
}

func (_this D11) Dtor_cf46() _dafny.MultiSet {
	return _this.Get_().(D11_DC29).Cf46
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf47 == data2.Cf47
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D12_DC32 struct {
	Cf49 bool
	Cf50 bool
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf49 bool, Cf50 bool) D12 {
	return D12{D12_DC32{Cf49, Cf50}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC31 struct {
	Cf48 _dafny.MultiSet
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf48 _dafny.MultiSet) D12 {
	return D12{D12_DC31{Cf48}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(false, false)
}

func (_this D12) Dtor_cf49() bool {
	return _this.Get_().(D12_DC32).Cf49
}

func (_this D12) Dtor_cf50() bool {
	return _this.Get_().(D12_DC32).Cf50
}

func (_this D12) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D12_DC31).Cf48
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

// Definition of datatype D13
type D13 struct {
	Data_D13_
}

func (_this D13) Get_() Data_D13_ {
	return _this.Data_D13_
}

type Data_D13_ interface {
	isD13()
}

type CompanionStruct_D13_ struct {
}

var Companion_D13_ = CompanionStruct_D13_{}

type D13_DC34 struct {
	Cf52 _dafny.Int
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf52 _dafny.Int) D13 {
	return D13{D13_DC34{Cf52}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

type D13_DC33 struct {
	Cf51 _dafny.Array
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf51 _dafny.Array) D13 {
	return D13{D13_DC33{Cf51}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC34_(_dafny.Zero)
}

func (_this D13) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D13_DC34).Cf52
}

func (_this D13) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D13_DC33).Cf51
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf51 == data2.Cf51
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D13) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D13)
	return ok && _this.Equals(typed)
}

func Type_D13_() _dafny.TypeDescriptor {
	return type_D13_{}
}

type type_D13_ struct {
}

func (_this type_D13_) Default() interface{} {
	return Companion_D13_.Default()
}

func (_this type_D13_) String() string {
	return "main.D13"
}
func (_this D13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D13{}

// End of datatype D13

// Definition of datatype D14
type D14 struct {
	Data_D14_
}

func (_this D14) Get_() Data_D14_ {
	return _this.Data_D14_
}

type Data_D14_ interface {
	isD14()
}

type CompanionStruct_D14_ struct {
}

var Companion_D14_ = CompanionStruct_D14_{}

type D14_DC36 struct {
	Cf54 _dafny.Int
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf54 _dafny.Int) D14 {
	return D14{D14_DC36{Cf54}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC35 struct {
	Cf53 _dafny.MultiSet
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf53 _dafny.MultiSet) D14 {
	return D14{D14_DC35{Cf53}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC37 struct {
	Cf55 D14
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf55 D14) D14 {
	return D14{D14_DC37{Cf55}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC36_(_dafny.Zero)
}

func (_this D14) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf54
}

func (_this D14) Dtor_cf53() _dafny.MultiSet {
	return _this.Get_().(D14_DC35).Cf53
}

func (_this D14) Dtor_cf55() D14 {
	return _this.Get_().(D14_DC37).Cf55
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf54) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf53.Equals(data2.Cf53)
		}
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D14) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D14)
	return ok && _this.Equals(typed)
}

func Type_D14_() _dafny.TypeDescriptor {
	return type_D14_{}
}

type type_D14_ struct {
}

func (_this type_D14_) Default() interface{} {
	return Companion_D14_.Default()
}

func (_this type_D14_) String() string {
	return "main.D14"
}
func (_this D14) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D14{}

// End of datatype D14

// Definition of datatype D15
type D15 struct {
	Data_D15_
}

func (_this D15) Get_() Data_D15_ {
	return _this.Data_D15_
}

type Data_D15_ interface {
	isD15()
}

type CompanionStruct_D15_ struct {
}

var Companion_D15_ = CompanionStruct_D15_{}

type D15_DC38 struct {
	Cf56 _dafny.Set
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf56 _dafny.Set) D15 {
	return D15{D15_DC38{Cf56}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D15) Dtor_cf56() _dafny.Set {
	return _this.Get_().(D15_DC38).Cf56
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf56.Equals(data2.Cf56)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D15) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D15)
	return ok && _this.Equals(typed)
}

func Type_D15_() _dafny.TypeDescriptor {
	return type_D15_{}
}

type type_D15_ struct {
}

func (_this type_D15_) Default() interface{} {
	return Companion_D15_.Default()
}

func (_this type_D15_) String() string {
	return "main.D15"
}
func (_this D15) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D15{}

// End of datatype D15

// Definition of datatype D16
type D16 struct {
	Data_D16_
}

func (_this D16) Get_() Data_D16_ {
	return _this.Data_D16_
}

type Data_D16_ interface {
	isD16()
}

type CompanionStruct_D16_ struct {
}

var Companion_D16_ = CompanionStruct_D16_{}

type D16_DC40 struct {
	Cf58 bool
	Cf59 bool
	Cf60 _dafny.Int
	Cf61 _dafny.Int
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf58 bool, Cf59 bool, Cf60 _dafny.Int, Cf61 _dafny.Int) D16 {
	return D16{D16_DC40{Cf58, Cf59, Cf60, Cf61}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC39 struct {
	Cf57 _dafny.Map
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf57 _dafny.Map) D16 {
	return D16{D16_DC39{Cf57}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC40_(false, false, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf58() bool {
	return _this.Get_().(D16_DC40).Cf58
}

func (_this D16) Dtor_cf59() bool {
	return _this.Get_().(D16_DC40).Cf59
}

func (_this D16) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf60
}

func (_this D16) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf61
}

func (_this D16) Dtor_cf57() _dafny.Map {
	return _this.Get_().(D16_DC39).Cf57
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59 && data1.Cf60.Cmp(data2.Cf60) == 0 && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D16) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D16)
	return ok && _this.Equals(typed)
}

func Type_D16_() _dafny.TypeDescriptor {
	return type_D16_{}
}

type type_D16_ struct {
}

func (_this type_D16_) Default() interface{} {
	return Companion_D16_.Default()
}

func (_this type_D16_) String() string {
	return "main.D16"
}
func (_this D16) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D16{}

// End of datatype D16

// Definition of trait T0
type T0 interface {
	String() string
	Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence
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
	Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence
	M0(p0 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int)
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
	F2   _dafny.Array
	F3   bool
	F4   _dafny.Map
	F5   bool
	F9   bool
	F13  bool
	_f0  _dafny.CodePoint
	_f1  _dafny.Array
	_f6  _dafny.Int
	_f7  _dafny.Map
	_f8  _dafny.Sequence
	_f10 bool
	_f11 _dafny.Set
	_f12 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F3 = false
	_this.F4 = _dafny.EmptyMap
	_this.F5 = false
	_this.F9 = false
	_this.F13 = false
	_this._f0 = _dafny.CodePoint('D')
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptyMap
	_this._f8 = _dafny.EmptySeq
	_this._f10 = false
	_this._f11 = _dafny.EmptySet
	_this._f12 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.CodePoint, f1 _dafny.Array, f2 _dafny.Array, f3 bool, f4 _dafny.Map, f5 bool, f6 _dafny.Int, f7 _dafny.Map, f8 _dafny.Sequence, f9 bool, f10 bool, f11 _dafny.Set, f12 _dafny.Sequence, f13 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
	}
}
func (_this *GlobalState) F0() _dafny.CodePoint {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F6() _dafny.Int {
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
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Set {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f18 _dafny.MultiSet
	_f19 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f18 = _dafny.EmptyMultiSet
	_this._f19 = false
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

func (_this *C0) Ctor__(f18 _dafny.MultiSet, f19 bool) {
	{
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C0) Fm21(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		if (_this).F19() {
			return _dafny.SeqOf(!((_this).F19()))
		} else {
			return _dafny.SeqOf(false, (_this).F19())
		}
	}
}
func (_this *C0) F18() _dafny.MultiSet {
	{
		return _this._f18
	}
}
func (_this *C0) F19() bool {
	{
		return _this._f19
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
func (_this *C1) M8(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _200_v0 bool
		_ = _200_v0
		_200_v0 = false
		var _201_v1 _dafny.Int
		_ = _201_v1
		_201_v1 = _dafny.IntOfInt64(-189)
		var _202_v2 D0
		_ = _202_v2
		_202_v2 = Companion_D0_.Create_DC2_(_200_v0)
		var _203_v3 _dafny.Sequence
		_ = _203_v3
		_203_v3 = _dafny.SeqOf(_202_v2, Companion_D0_.Create_DC2_(_200_v0), _202_v2, _202_v2, _202_v2)
		(globalState).F13 = (func() bool {
			if _200_v0 {
				return false
			}
			return ((_dafny.Zero).Minus(_201_v1)).Cmp(_dafny.IntOfUint32((_203_v3).Cardinality())) <= 0
		})()
		var _204_v4 _dafny.Map
		_ = _204_v4
		_204_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, _200_v0)
		_204_v4 = (_204_v4).Update((_201_v1).Cmp(_dafny.IntOfInt64(-972)) >= 0, _200_v0)
		if !(_200_v0) {
			var _205_v5 _dafny.Array
			_ = _205_v5
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_0
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = func(_206_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_206_i0, _dafny.IntOfInt64(-167))
				}
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw21 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw21).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw21).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_205_v5 = _nw21
			var _207_v6 _dafny.Map
			_ = _207_v6
			_207_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _205_v5)
			var _208_v7 _dafny.Map
			_ = _208_v7
			_208_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_201_v1, _200_v0, globalState), (func() _dafny.Array {
				if (_207_v6).Contains(_201_v1) {
					return (_207_v6).Get(_201_v1).(_dafny.Array)
				}
				return _205_v5
			})())
			var _209_v8 D1
			_ = _209_v8
			_209_v8 = Companion_D1_.Create_DC4_(_208_v7)
			_208_v7 = (_209_v8).Dtor_cf6()
			var _210_v9 _dafny.Sequence
			_ = _210_v9
			_210_v9 = _dafny.UnicodeSeqOfUtf8Bytes("a")
			_210_v9 = _210_v9
			var _211_v10 _dafny.Array
			_ = _211_v10
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_1
			var _nw22 _dafny.Array
			_ = _nw22
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw22 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_212_v0 bool, _213_v4 _dafny.Map) func(_dafny.Int) bool {
					return func(_214_i1 _dafny.Int) bool {
						return ((func() bool {
							if (_213_v4).Contains(_212_v0) {
								return (_213_v4).Get(_212_v0).(bool)
							}
							return true
						})()) == (_212_v0)
					}
				})(_200_v0, _204_v4)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw22 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw22).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw22).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_211_v10 = _nw22
			var _215_v11 _dafny.Sequence
			_ = _215_v11
			_215_v11 = _dafny.SeqOf(_200_v0)
			var _216_v12 _dafny.Sequence
			_ = _216_v12
			_216_v12 = _dafny.SeqOf(_200_v0, _200_v0, (_215_v11).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_215_v11).Cardinality()))).Uint32()).(bool))
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_211_v10), 0))
			_ = _index21
			(_211_v10).ArraySet1((_216_v12).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_216_v12).Cardinality()))).Uint32()).(bool), (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_205_v5), 0))
			_ = _index22
			(_205_v5).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm20(_dafny.IntOfUint32((_210_v9).Cardinality()), _200_v0, _200_v0, _201_v1, globalState)).Cardinality()))).Minus(_201_v1), (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_211_v10), 0))
			_ = _index23
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_205_v5), 0))
			_ = _index24
			var _rhs38 bool = !(false)
			_ = _rhs38
			var _rhs39 _dafny.Int = ((_201_v1).Plus(_201_v1)).Minus((_201_v1).Plus(_201_v1))
			_ = _rhs39
			var _rhs40 bool = (((_dafny.Zero).Minus(_dafny.IntOfInt64(-89))).Times(_201_v1)).Cmp(Companion_Default___.SafeDivisionInt(_201_v1, _201_v1)) != 0
			_ = _rhs40
			var _lhs27 _dafny.Array = _211_v10
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_211_v10), 0))
			_ = _lhs28
			var _lhs29 _dafny.Array = _205_v5
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_205_v5), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			(_lhs27).ArraySet1(_rhs38, (_lhs28).Int())
			(_lhs29).ArraySet1(_rhs39, (_lhs30).Int())
			_lhs31.F3 = _rhs40
			(globalState).F3 = true
			var _217_v13 _dafny.MultiSet
			_ = _217_v13
			_217_v13 = _dafny.MultiSetOf(!(_200_v0), _200_v0, (_211_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_211_v10), 0))).Int()).(bool))
			r0 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_217_v13).Contains(_200_v0) {
					return (_217_v13).Multiplicity(_200_v0)
				}
				return (_205_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_205_v5), 0))).Int()).(_dafny.Int)
			})(), (_205_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_205_v5), 0))).Int()).(_dafny.Int))
		} else {
			var _218_v14 _dafny.Array
			_ = _218_v14
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw23
			_218_v14 = _nw23
			var _219_v15 _dafny.Sequence
			_ = _219_v15
			_219_v15 = _dafny.SeqOf(_201_v1)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
			_ = _index25
			(_218_v14).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_219_v15, (Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_219_v15).Cardinality()))).Uint32(), _201_v1)).Cardinality()), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
			_ = _index26
			(_218_v14).ArraySet1((_dafny.Zero).Minus(_201_v1), (_index26).Int())
			var _220_v16 _dafny.Array
			_ = _220_v16
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw24
			_220_v16 = _nw24
			var _221_v17 _dafny.Array
			_ = _221_v17
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
			_ = _nw25
			_221_v17 = _nw25
			var _222_v18 _dafny.Sequence
			_ = _222_v18
			_222_v18 = _dafny.SeqOf(_221_v17, _221_v17, _221_v17, _221_v17, _221_v17)
			var _223_v19 _dafny.Array
			_ = _223_v19
			var _nwElement0_3 _dafny.Array = _221_v17
			_ = _nwElement0_3
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_3, 0)
			(_nw26).ArraySet1(_221_v17, 1)
			(_nw26).ArraySet1(_221_v17, 2)
			(_nw26).ArraySet1(_221_v17, 3)
			(_nw26).ArraySet1(_221_v17, 4)
			(_nw26).ArraySet1(_221_v17, 5)
			(_nw26).ArraySet1(_221_v17, 6)
			(_nw26).ArraySet1((_222_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-11), _dafny.IntOfUint32((_222_v18).Cardinality()))).Uint32()).(_dafny.Array), 7)
			(_nw26).ArraySet1(_221_v17, 8)
			(_nw26).ArraySet1(_221_v17, 9)
			(_nw26).ArraySet1(_221_v17, 10)
			(_nw26).ArraySet1((func() _dafny.Array {
				if _200_v0 {
					return _221_v17
				}
				return _221_v17
			})(), 11)
			(_nw26).ArraySet1(_221_v17, 12)
			(_nw26).ArraySet1(_221_v17, 13)
			(_nw26).ArraySet1(_221_v17, 14)
			(_nw26).ArraySet1(_221_v17, 15)
			(_nw26).ArraySet1(_221_v17, 16)
			(_nw26).ArraySet1(_221_v17, 17)
			(_nw26).ArraySet1(_221_v17, 18)
			(_nw26).ArraySet1(_221_v17, 19)
			(_nw26).ArraySet1(_221_v17, 20)
			(_nw26).ArraySet1((func() _dafny.Array {
				if _200_v0 {
					return _221_v17
				}
				return _221_v17
			})(), 21)
			(_nw26).ArraySet1((func() _dafny.Array {
				if _200_v0 {
					return _221_v17
				}
				return _221_v17
			})(), 22)
			(_nw26).ArraySet1(_221_v17, 23)
			(_nw26).ArraySet1(_221_v17, 24)
			(_nw26).ArraySet1(_221_v17, 25)
			(_nw26).ArraySet1(_221_v17, 26)
			(_nw26).ArraySet1(_221_v17, 27)
			_223_v19 = _nw26
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_223_v19), 0))
			_ = _index27
			(_223_v19).ArraySet1(_221_v17, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_223_v19), 0))
			_ = _index28
			var _rhs41 _dafny.Array = _220_v16
			_ = _rhs41
			var _rhs42 _dafny.Array = _221_v17
			_ = _rhs42
			var _rhs43 bool = _200_v0
			_ = _rhs43
			var _lhs32 _dafny.Array = _223_v19
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_223_v19), 0))
			_ = _lhs33
			_220_v16 = _rhs41
			(_lhs32).ArraySet1(_rhs42, (_lhs33).Int())
			_200_v0 = _rhs43
			var _224_v20 _dafny.Sequence
			_ = _224_v20
			_224_v20 = _dafny.UnicodeSeqOfUtf8Bytes("ddwvkjgl")
			if !_dafny.Companion_Sequence_.Equal(_224_v20, _224_v20) {
				var _225_v21 _dafny.MultiSet
				_ = _225_v21
				_225_v21 = _dafny.MultiSetOf(_201_v1, _201_v1)
				var _226_v22 _dafny.Map
				_ = _226_v22
				_226_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), (_225_v21).Intersection(_225_v21))
				_226_v22 = (_226_v22).Update(((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)).Plus(_201_v1), (_dafny.MultiSetFromSeq(_219_v15)).Intersection((_225_v21).Update((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(_201_v1))))
				var _227_v23 _dafny.Sequence
				_ = _227_v23
				_227_v23 = _dafny.SeqOf(_224_v20, _224_v20)
				var _228_v24 _dafny.Sequence
				_ = _228_v24
				_228_v24 = _dafny.SeqOf((_227_v23).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_227_v23).Cardinality()))).Uint32()).(_dafny.Sequence), _224_v20)
				var _229_v25 _dafny.Sequence
				_ = _229_v25
				_229_v25 = _dafny.SeqOf(_200_v0)
				var _230_v26 D1
				_ = _230_v26
				_230_v26 = Companion_D1_.Create_DC5_(_200_v0, _224_v20, _229_v25)
				var _231_v27 _dafny.Map
				_ = _231_v27
				_231_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))
				var _232_v28 _dafny.Map
				_ = _232_v28
				_232_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v4, _200_v0)
				var _233_v29 _dafny.Set
				_ = _233_v29
				_233_v29 = _dafny.SetOf(_200_v0)
				var _234_v30 _dafny.MultiSet
				_ = _234_v30
				_234_v30 = _dafny.MultiSetOf(_233_v29)
				var _235_v31 _dafny.Array
				_ = _235_v31
				var _nwElement0_4 bool = _200_v0
				_ = _nwElement0_4
				var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
				_ = _nw27
				(_nw27).ArraySet1(_nwElement0_4, 0)
				(_nw27).ArraySet1((false) && (_200_v0), 1)
				(_nw27).ArraySet1(true, 2)
				(_nw27).ArraySet1(_200_v0, 3)
				(_nw27).ArraySet1(_200_v0, 4)
				(_nw27).ArraySet1(!(_200_v0), 5)
				(_nw27).ArraySet1(_200_v0, 6)
				(_nw27).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_224_v20, _dafny.UnicodeSeqOfUtf8Bytes("rexhpd")), 7)
				(_nw27).ArraySet1(_200_v0, 8)
				(_nw27).ArraySet1(_200_v0, 9)
				(_nw27).ArraySet1(!(_200_v0), 10)
				(_nw27).ArraySet1((_201_v1).Cmp(_dafny.IntOfUint32(((_227_v23).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_225_v21).Contains((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)) {
						return (_225_v21).Multiplicity((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))
					}
					return (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_227_v23).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) > 0, 11)
				(_nw27).ArraySet1(!(_200_v0), 12)
				(_nw27).ArraySet1(true, 13)
				(_nw27).ArraySet1((_230_v26).Dtor_cf7(), 14)
				(_nw27).ArraySet1(_200_v0, 15)
				(_nw27).ArraySet1(((_231_v27).Update((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))).Equals(_231_v27), 16)
				(_nw27).ArraySet1(_200_v0, 17)
				(_nw27).ArraySet1((func() bool {
					if (func() bool {
						if (_232_v28).Contains(_204_v4) {
							return (_232_v28).Get(_204_v4).(bool)
						}
						return _200_v0
					})() {
						return true
					}
					return !(!(_200_v0))
				})(), 18)
				(_nw27).ArraySet1(true, 19)
				(_nw27).ArraySet1((_225_v21).IsProperSubsetOf(_225_v21), 20)
				(_nw27).ArraySet1(_200_v0, 21)
				(_nw27).ArraySet1(_200_v0, 22)
				(_nw27).ArraySet1(((func() _dafny.Int {
					if (_234_v30).Contains(_233_v29) {
						return (_234_v30).Multiplicity(_233_v29)
					}
					return (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)
				})()).Cmp(_201_v1) == 0, 23)
				(_nw27).ArraySet1(_200_v0, 24)
				(_nw27).ArraySet1(_dafny.Companion_Sequence_.Equal(_229_v25, _229_v25), 25)
				(_nw27).ArraySet1(_200_v0, 26)
				_235_v31 = _nw27
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_235_v31), 0))
				_ = _index29
				(_235_v31).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_224_v20, _224_v20), (_index29).Int())
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_235_v31), 0))
				_ = _index30
				(_235_v31).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("h"), _dafny.UnicodeSeqOfUtf8Bytes("kfkrjpla")), (_index30).Int())
				var _236_v32 D0
				_ = _236_v32
				_236_v32 = Companion_D0_.Create_DC1_(_200_v0, _227_v23, _224_v20)
				var _237_v33 _dafny.MultiSet
				_ = _237_v33
				_237_v33 = _dafny.MultiSetOf(_236_v32, _236_v32, _236_v32, _236_v32)
				var _238_v34 *C0
				_ = _238_v34
				var _nw28 *C0 = New_C0_()
				_ = _nw28
				_nw28.Ctor__((Companion_D2_.Create_DC6_(_237_v33)).Dtor_cf10(), !(!(false)))
				_238_v34 = _nw28
				_235_v31 = _235_v31
				var _239_v35 _dafny.Array
				_ = _239_v35
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_2
				var _nw29 _dafny.Array
				_ = _nw29
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw29 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) _dafny.Sequence = (func(_240_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_241_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_240_v20, _240_v20)
						}
					})(_224_v20)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw29 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw29).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw29).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_239_v35 = _nw29
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_239_v35), 0))
				_ = _index31
				(_239_v35).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mcihcac"), (_index31).Int())
				var _242_v36 _dafny.Map
				_ = _242_v36
				_242_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_238_v34).F19(), (func() _dafny.Int {
					if (_238_v34).F19() {
						return (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)
					}
					return (_dafny.SetOf(_dafny.IntOfInt64(584), (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))).Cardinality()
				})())
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_239_v35), 0))
				_ = _index32
				var _rhs44 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_224_v20, _dafny.Companion_Sequence_.Concatenate(_224_v20, _224_v20))
				_ = _rhs44
				var _rhs45 _dafny.Int = (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)
				_ = _rhs45
				var _rhs46 _dafny.Map = _242_v36
				_ = _rhs46
				var _lhs34 _dafny.Array = _239_v35
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_239_v35), 0))
				_ = _lhs35
				(_lhs34).ArraySet1(_rhs44, (_lhs35).Int())
				r0 = _rhs45
				_242_v36 = _rhs46
			} else {
				(globalState).F3 = (_202_v2).Dtor_cf4()
				var _243_v37 _dafny.CodePoint
				_ = _243_v37
				_243_v37 = _dafny.CodePoint('s')
				var _244_v38 _dafny.Array
				_ = _244_v38
				var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw30
				_244_v38 = _nw30
				var _245_v39 _dafny.MultiSet
				_ = _245_v39
				_245_v39 = _dafny.MultiSetOf(_201_v1, _201_v1, (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))
				var _246_v40 _dafny.Sequence
				_ = _246_v40
				_246_v40 = _dafny.SeqOf(_224_v20)
				var _247_v41 D0
				_ = _247_v41
				_247_v41 = Companion_D0_.Create_DC2_(_200_v0)
				var _248_v42 D0
				_ = _248_v42
				_248_v42 = Companion_D0_.Create_DC3_(_247_v41)
				var _249_v43 _dafny.Map
				_ = _249_v43
				_249_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_(_200_v0, _246_v40, _224_v20)).Dtor_cf1(), _248_v42)
				var _250_v44 D2
				_ = _250_v44
				_250_v44 = Companion_D2_.Create_DC7_(_200_v0, _249_v43, _201_v1, _200_v0)
				var _251_v45 _dafny.Map
				_ = _251_v45
				_251_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v20, _244_v38)
				var _252_v46 _dafny.Set
				_ = _252_v46
				_252_v46 = _dafny.SetOf((_dafny.Zero).Minus((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int)))
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
				_ = _index33
				var _rhs47 _dafny.Int = Companion_Default___.Fm22(_250_v44, _248_v42, _200_v0, globalState)
				_ = _rhs47
				var _rhs48 _dafny.Int = Companion_Default___.Fm22(_250_v44, _248_v42, _200_v0, globalState)
				_ = _rhs48
				var _rhs49 _dafny.CodePoint = Companion_Default___.Fm23(_200_v0, Companion_Default___.Fm24(globalState), _200_v0, (_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), globalState)
				_ = _rhs49
				var _rhs50 _dafny.Array = (func() _dafny.Array {
					if (_251_v45).Contains(_dafny.Companion_Sequence_.Concatenate(_224_v20, _224_v20)) {
						return (_251_v45).Get(_dafny.Companion_Sequence_.Concatenate(_224_v20, _224_v20)).(_dafny.Array)
					}
					return _244_v38
				})()
				_ = _rhs50
				var _rhs51 _dafny.MultiSet = (_245_v39).Update((_dafny.IntOfUint32((_224_v20).Cardinality())).Times(_201_v1), Companion_Default___.Abs(((_252_v46).Difference(_252_v46)).Cardinality()))
				_ = _rhs51
				var _lhs36 _dafny.Array = _218_v14
				_ = _lhs36
				var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
				_ = _lhs37
				(_lhs36).ArraySet1(_rhs47, (_lhs37).Int())
				r0 = _rhs48
				_243_v37 = _rhs49
				_244_v38 = _rhs50
				_245_v39 = _rhs51
				var _253_v47 D0
				_ = _253_v47
				_253_v47 = Companion_D0_.Create_DC1_(_200_v0, _246_v40, Companion_Default___.Fm20(_201_v1, !(_200_v0), _200_v0, _dafny.IntOfUint32((_224_v20).Cardinality()), globalState))
				var _254_v48 *C0
				_ = _254_v48
				var _nw31 *C0 = New_C0_()
				_ = _nw31
				_nw31.Ctor__(_dafny.MultiSetFromSeq(_dafny.SeqOf(_253_v47, _253_v47)), _200_v0)
				_254_v48 = _nw31
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
				_ = _index34
				(_218_v14).ArraySet1((_dafny.Zero).Minus((_201_v1).Plus(Companion_Default___.SafeDivisionInt((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_201_v1)))), (_index34).Int())
				var _255_v49 *C0
				_ = _255_v49
				var _nw32 *C0 = New_C0_()
				_ = _nw32
				_nw32.Ctor__(((_254_v48).F18()).Update(_253_v47, Companion_Default___.Abs((_252_v46).Cardinality())), !((_254_v48).F19()))
				_255_v49 = _nw32
			}
			(globalState).F5 = _200_v0
			var _256_v50 D0
			_ = _256_v50
			_256_v50 = Companion_D0_.Create_DC0_((_218_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))).Int()).(_dafny.Int))
			var _257_v51 D0
			_ = _257_v51
			_257_v51 = Companion_D0_.Create_DC3_(_256_v50)
			var _258_v52 _dafny.Map
			_ = _258_v52
			_258_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, _257_v51)
			var _259_v53 D2
			_ = _259_v53
			_259_v53 = Companion_D2_.Create_DC7_(_200_v0, _258_v52, _201_v1, false)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_218_v14), 0))
			_ = _index35
			(_218_v14).ArraySet1(Companion_Default___.Fm22(_259_v53, Companion_D0_.Create_DC3_(_256_v50), !(_200_v0), globalState), (_index35).Int())
		}
		var _260_v54 _dafny.Map
		_ = _260_v54
		_260_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _201_v1)
		var _261_v55 _dafny.Sequence
		_ = _261_v55
		_261_v55 = _dafny.UnicodeSeqOfUtf8Bytes("tlaqlophn")
		_260_v54 = (_260_v54).Update(_dafny.IntOfUint32((_261_v55).Cardinality()), _201_v1)
		var _262_v56 _dafny.Map
		_ = _262_v56
		_262_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, _201_v1)
		var _263_v57 _dafny.CodePoint
		_ = _263_v57
		_263_v57 = _dafny.CodePoint('g')
		var _264_v58 _dafny.MultiSet
		_ = _264_v58
		_264_v58 = _dafny.MultiSetOf(_263_v57)
		r0 = (_201_v1).Plus((func() _dafny.Int {
			if (_262_v56).Contains(_200_v0) {
				return (_262_v56).Get(_200_v0).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_264_v58).Contains(_263_v57) {
					return (_264_v58).Multiplicity(_263_v57)
				}
				return _201_v1
			})()
		})())
		var _265_i3 _dafny.Int
		_ = _265_i3
		_265_i3 = _dafny.Zero
		{
			for _200_v0 {
				{
					if (_265_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_265_i3 = (_265_i3).Plus(_dafny.One)
					if (Companion_Default___.Fm19(_201_v1, _200_v0, globalState)) && (_200_v0) {
						var _266_v59 _dafny.Array
						_ = _266_v59
						var _len0_3 _dafny.Int = _dafny.IntOfInt64(20)
						_ = _len0_3
						var _nw33 _dafny.Array
						_ = _nw33
						if _len0_3.Cmp(_dafny.Zero) == 0 {
							_nw33 = _dafny.NewArray(_len0_3)
						} else {
							var _init3 func(_dafny.Int) _dafny.Int = (func(_267_v0 bool) func(_dafny.Int) _dafny.Int {
								return func(_268_i4 _dafny.Int) _dafny.Int {
									return (_268_i4).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_267_v0)).Cardinality()))
								}
							})(_200_v0)
							_ = _init3
							var _element0_3 = _init3(_dafny.Zero)
							_ = _element0_3
							_nw33 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
							(_nw33).ArraySet1(_element0_3, 0)
							var _nativeLen0_3 = (_len0_3).Int()
							_ = _nativeLen0_3
							for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
								(_nw33).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
							}
						}
						_266_v59 = _nw33
						var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _index36
						(_266_v59).ArraySet1((_dafny.Zero).Minus(_201_v1), (_index36).Int())
						var _269_v60 _dafny.Map
						_ = _269_v60
						_269_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _200_v0)
						var _270_v61 D0
						_ = _270_v61
						_270_v61 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_269_v60).Cardinality()))
						var _271_v62 D0
						_ = _271_v62
						_271_v62 = Companion_D0_.Create_DC3_(_270_v61)
						var _272_v63 _dafny.Map
						_ = _272_v63
						_272_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, _271_v62)
						var _273_v64 D2
						_ = _273_v64
						_273_v64 = Companion_D2_.Create_DC7_(_200_v0, _272_v63, _201_v1, _200_v0)
						var _274_v65 _dafny.Sequence
						_ = _274_v65
						_274_v65 = _dafny.SeqOf(_200_v0)
						var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _index37
						var _rhs52 _dafny.Int = Companion_Default___.SafeModuloInt(_201_v1, _dafny.IntOfUint32((Companion_Default___.Fm20((_262_v56).Cardinality(), _200_v0, _200_v0, Companion_Default___.Fm22(_273_v64, _271_v62, _200_v0, globalState), globalState)).Cardinality()))
						_ = _rhs52
						var _rhs53 bool = !(_200_v0) || (_200_v0)
						_ = _rhs53
						var _rhs54 _dafny.Int = _dafny.IntOfUint32((_274_v65).Cardinality())
						_ = _rhs54
						var _lhs38 _dafny.Array = _266_v59
						_ = _lhs38
						var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _lhs39
						var _lhs40 *GlobalState = globalState
						_ = _lhs40
						(_lhs38).ArraySet1(_rhs52, (_lhs39).Int())
						_lhs40.F5 = _rhs53
						_201_v1 = _rhs54
						var _275_v66 _dafny.MultiSet
						_ = _275_v66
						_275_v66 = _dafny.MultiSetOf((_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int))
						var _276_v67 _dafny.Sequence
						_ = _276_v67
						_276_v67 = _dafny.SeqOf((_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int))).Cardinality()))
						var _rhs55 _dafny.MultiSet = ((Companion_Default___.Fm25(_200_v0, _276_v67, globalState)).Intersection(_275_v66)).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(-802), _201_v1, _201_v1)).Union(_275_v66))
						_ = _rhs55
						var _rhs56 bool = _200_v0
						_ = _rhs56
						var _lhs41 *GlobalState = globalState
						_ = _lhs41
						_275_v66 = _rhs55
						_lhs41.F13 = _rhs56
						var _277_v68 D1
						_ = _277_v68
						_277_v68 = Companion_D1_.Create_DC5_(_200_v0, _261_v55, _dafny.Companion_Sequence_.Update(_274_v65, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_201_v1), _dafny.IntOfUint32((_274_v65).Cardinality()))).Uint32(), (_274_v65).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_274_v65).Cardinality()))).Uint32()).(bool)))
						var _pat_let_tv6 = _261_v55
						_ = _pat_let_tv6
						var _278_v69 _dafny.Array
						_ = _278_v69
						var _nwElement0_5 D1 = Companion_Default___.Fm26(_200_v0, _200_v0, globalState)
						_ = _nwElement0_5
						var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(5))
						_ = _nw34
						(_nw34).ArraySet1(_nwElement0_5, 0)
						(_nw34).ArraySet1(_277_v68, 1)
						(_nw34).ArraySet1(_277_v68, 2)
						(_nw34).ArraySet1(_277_v68, 3)
						(_nw34).ArraySet1(func(_pat_let2_0 D1) D1 {
							return func(_279_dt__update__tmp_h0 D1) D1 {
								return func(_pat_let3_0 _dafny.Sequence) D1 {
									return func(_280_dt__update_hcf8_h0 _dafny.Sequence) D1 {
										return Companion_D1_.Create_DC5_((_279_dt__update__tmp_h0).Dtor_cf7(), _280_dt__update_hcf8_h0, (_279_dt__update__tmp_h0).Dtor_cf9())
									}(_pat_let3_0)
								}(_pat_let_tv6)
							}(_pat_let2_0)
						}(Companion_Default___.Fm26(_200_v0, _200_v0, globalState)), 4)
						_278_v69 = _nw34
						var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_278_v69), 0))
						_ = _index38
						(_278_v69).ArraySet1(_277_v68, (_index38).Int())
						var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_278_v69), 0))
						_ = _index39
						(_278_v69).ArraySet1(_277_v68, (_index39).Int())
						var _281_v70 _dafny.Set
						_ = _281_v70
						_281_v70 = _dafny.SetOf(_dafny.SeqOf((_275_v66).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-342))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}((func(_282_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_283_i5 _dafny.Int) _dafny.Int {
								return _282_v1
							}
						})(_201_v1))))
						var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _index40
						var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _index41
						var _rhs57 _dafny.Int = Companion_Default___.SafeDivisionInt((_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm20((_dafny.Zero).Minus(_dafny.IntOfUint32((_276_v67).Cardinality())), !(_200_v0), true, (_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int), globalState), _261_v55)).Cardinality()))
						_ = _rhs57
						var _rhs58 bool = (_dafny.SetOf(_276_v67)).IsProperSubsetOf(_281_v70)
						_ = _rhs58
						var _rhs59 _dafny.Int = ((_dafny.Zero).Minus((func() _dafny.Int {
							if _200_v0 {
								return (_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int)
							}
							return _dafny.IntOfInt64(273)
						})())).Plus((_201_v1).Minus(_201_v1))
						_ = _rhs59
						var _rhs60 _dafny.Map = _204_v4
						_ = _rhs60
						var _lhs42 _dafny.Array = _266_v59
						_ = _lhs42
						var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _lhs43
						var _lhs44 _dafny.Array = _266_v59
						_ = _lhs44
						var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))
						_ = _lhs45
						(_lhs42).ArraySet1(_rhs57, (_lhs43).Int())
						_200_v0 = _rhs58
						(_lhs44).ArraySet1(_rhs59, (_lhs45).Int())
						_204_v4 = _rhs60
						r0 = (_276_v67).Select((Companion_Default___.SafeIndex((_266_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_266_v59), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_276_v67).Cardinality()))).Uint32()).(_dafny.Int)
					} else {
						var _284_v71 _dafny.MultiSet
						_ = _284_v71
						_284_v71 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, _263_v57)).Cardinality())
						var _285_v72 _dafny.Sequence
						_ = _285_v72
						_285_v72 = _dafny.SeqOf(_dafny.IntOfUint32((_261_v55).Cardinality()), _201_v1)
						var _286_v73 _dafny.MultiSet
						_ = _286_v73
						_286_v73 = _dafny.MultiSetOf(_285_v72)
						var _287_v74 _dafny.Sequence
						_ = _287_v74
						_287_v74 = _dafny.SeqOf(_260_v54, _260_v54, _260_v54, _260_v54)
						var _288_v75 _dafny.Array
						_ = _288_v75
						var _nwElement0_6 bool = ((_284_v71).Update(_201_v1, Companion_Default___.Abs(_201_v1))).IsSubsetOf(_284_v71)
						_ = _nwElement0_6
						var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(14))
						_ = _nw35
						(_nw35).ArraySet1(_nwElement0_6, 0)
						(_nw35).ArraySet1(_200_v0, 1)
						(_nw35).ArraySet1(false, 2)
						(_nw35).ArraySet1(_200_v0, 3)
						(_nw35).ArraySet1((_200_v0) || (_200_v0), 4)
						(_nw35).ArraySet1(((_286_v73).Cardinality()).Cmp(_201_v1) != 0, 5)
						(_nw35).ArraySet1(_200_v0, 6)
						(_nw35).ArraySet1(_200_v0, 7)
						(_nw35).ArraySet1(_200_v0, 8)
						(_nw35).ArraySet1(!(true) || (false), 9)
						(_nw35).ArraySet1((false) || (_200_v0), 10)
						(_nw35).ArraySet1(!_dafny.Companion_Sequence_.Equal(_287_v74, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _201_v1))), 11)
						(_nw35).ArraySet1(_200_v0, 12)
						(_nw35).ArraySet1(_200_v0, 13)
						_288_v75 = _nw35
						var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))
						_ = _index42
						(_288_v75).ArraySet1(_200_v0, (_index42).Int())
						var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))
						_ = _index43
						(_288_v75).ArraySet1((_dafny.Companion_Sequence_.IsPrefixOf(_261_v55, _261_v55)) || (_200_v0), (_index43).Int())
						var _289_v76 D0
						_ = _289_v76
						_289_v76 = Companion_D0_.Create_DC1_((_288_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg36 _dafny.Int) interface{} {
								return coer36(arg36)
							}
						}((func(_290_v55 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_291_i6 _dafny.Int) _dafny.Sequence {
								return _290_v55
							}
						})(_261_v55))), _261_v55)
						var _292_v77 *C0
						_ = _292_v77
						var _nw36 *C0 = New_C0_()
						_ = _nw36
						_nw36.Ctor__(_dafny.MultiSetOf(_289_v76, _289_v76, _289_v76), _200_v0)
						_292_v77 = _nw36
						var _293_v78 _dafny.Array
						_ = _293_v78
						var _nwElement0_7 *C0 = _292_v77
						_ = _nwElement0_7
						var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(17))
						_ = _nw37
						(_nw37).ArraySet1(_nwElement0_7, 0)
						(_nw37).ArraySet1(_292_v77, 1)
						(_nw37).ArraySet1(_292_v77, 2)
						(_nw37).ArraySet1(_292_v77, 3)
						(_nw37).ArraySet1(_292_v77, 4)
						(_nw37).ArraySet1(_292_v77, 5)
						(_nw37).ArraySet1(_292_v77, 6)
						(_nw37).ArraySet1(_292_v77, 7)
						(_nw37).ArraySet1(_292_v77, 8)
						(_nw37).ArraySet1(_292_v77, 9)
						(_nw37).ArraySet1(_292_v77, 10)
						(_nw37).ArraySet1(_292_v77, 11)
						(_nw37).ArraySet1(_292_v77, 12)
						(_nw37).ArraySet1(_292_v77, 13)
						(_nw37).ArraySet1(_292_v77, 14)
						(_nw37).ArraySet1(_292_v77, 15)
						(_nw37).ArraySet1(_292_v77, 16)
						_293_v78 = _nw37
						var _294_v79 _dafny.Map
						_ = _294_v79
						_294_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v0, (func() _dafny.Array {
							if (_288_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))).Int()).(bool) {
								return _293_v78
							}
							return _293_v78
						})())
						_294_v79 = (_294_v79).Update(false, _293_v78)
						_261_v55 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("phrjej"), _261_v55)
						var _295_v80 _dafny.Map
						_ = _295_v80
						_295_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_292_v77).F19(), _262_v56)
						var _296_v81 _dafny.Sequence
						_ = _296_v81
						_296_v81 = _dafny.SeqOf(true)
						_262_v56 = ((func() _dafny.Map {
							if _200_v0 {
								return _262_v56
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_292_v77).F19(), _201_v1)
						})()).Merge((func() _dafny.Map {
							if (_295_v80).Contains((_296_v81).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_296_v81).Cardinality()))).Uint32()).(bool)) {
								return (_295_v80).Get((_296_v81).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_296_v81).Cardinality()))).Uint32()).(bool)).(_dafny.Map)
							}
							return _262_v56
						})())
						var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))
						_ = _index44
						(_288_v75).ArraySet1((func() bool {
							if (_204_v4).Contains((_288_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))).Int()).(bool)) {
								return (_204_v4).Get((_288_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_288_v75), 0))).Int()).(bool)).(bool)
							}
							return !((_292_v77).F19())
						})(), (_index44).Int())
					}
					var _297_v83 _dafny.MultiSet
					_ = _297_v83
					_297_v83 = _dafny.MultiSetOf(_201_v1)
					var _298_v84 _dafny.Sequence
					_ = _298_v84
					_298_v84 = _dafny.SeqOf(_201_v1, (func() _dafny.Map {
						var _coll5 = _dafny.NewMapBuilder()
						_ = _coll5
						for _iter5 := _dafny.Iterate((_297_v83).Elements()); ; {
							_compr_5, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _299_v82 _dafny.Int
							_299_v82 = interface{}(_compr_5).(_dafny.Int)
							if (_297_v83).Contains(_299_v82) {
								_coll5.Add(Companion_Default___.SafeModuloInt(_299_v82, (func() _dafny.Int {
									if (_262_v56).Contains(_200_v0) {
										return (_262_v56).Get(_200_v0).(_dafny.Int)
									}
									return _dafny.IntOfUint32((_261_v55).Cardinality())
								})()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_201_v1)).Cardinality(), _200_v0)).Cardinality())
							}
						}
						return _coll5.ToMap()
					}()).Cardinality(), _201_v1)
					r0 = (_298_v84).Select((Companion_Default___.SafeIndex((_298_v84).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_298_v84).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_298_v84).Cardinality()))).Uint32()).(_dafny.Int)
					var _300_v86 _dafny.Map
					_ = _300_v86
					_300_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _261_v55)
					var _301_v88 _dafny.Array
					_ = _301_v88
					var _nwElement0_8 _dafny.Map = (func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(955), _dafny.IntOfInt64(953))); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _302_v85 _dafny.Int
							_302_v85 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(955)).Cmp(_302_v85) <= 0) && ((_302_v85).Cmp(_dafny.IntOfInt64(953)) < 0) {
								_coll6.Add(Companion_Default___.SafeDivisionInt(_302_v85, _dafny.IntOfInt64(-853)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(346))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg37 _dafny.Int) interface{} {
										return coer37(arg37)
									}
								}(func(_303_i7 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('i')
								})))
							}
						}
						return _coll6.ToMap()
					}()).Merge(_300_v86)
					_ = _nwElement0_8
					var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(11))
					_ = _nw38
					(_nw38).ArraySet1(_nwElement0_8, 0)
					(_nw38).ArraySet1(_300_v86, 1)
					(_nw38).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _261_v55)).Merge(_300_v86), 2)
					(_nw38).ArraySet1(((_300_v86).Update(_dafny.IntOfInt64(418), Companion_Default___.Fm20((_260_v54).Cardinality(), _200_v0, _200_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gj")).Cardinality()), globalState))).Merge(_300_v86), 3)
					(_nw38).ArraySet1((_300_v86).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _261_v55)), 4)
					(_nw38).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v1, _261_v55), 5)
					(_nw38).ArraySet1(_300_v86, 6)
					(_nw38).ArraySet1(_300_v86, 7)
					(_nw38).ArraySet1((_300_v86).Merge(_300_v86), 8)
					(_nw38).ArraySet1(_300_v86, 9)
					(_nw38).ArraySet1((func() _dafny.Map {
						var _coll7 = _dafny.NewMapBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(81), _dafny.IntOfInt64(133))); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _304_v87 _dafny.Int
							_304_v87 = interface{}(_compr_7).(_dafny.Int)
							if ((_dafny.IntOfInt64(81)).Cmp(_304_v87) <= 0) && ((_304_v87).Cmp(_dafny.IntOfInt64(133)) < 0) {
								_coll7.Add((_304_v87).Minus((_264_v58).Cardinality()), _261_v55)
							}
						}
						return _coll7.ToMap()
					}()).Update(_201_v1, _261_v55), 10)
					_301_v88 = _nw38
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_301_v88), 0))
					_ = _index45
					(_301_v88).ArraySet1((_300_v86).Merge(_300_v86), (_index45).Int())
					var _305_v89 _dafny.Sequence
					_ = _305_v89
					_305_v89 = _dafny.SeqOf(_300_v86)
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_301_v88), 0))
					_ = _index46
					(_301_v88).ArraySet1(((_300_v86).Merge((_305_v89).Select((Companion_Default___.SafeIndex(_201_v1, _dafny.IntOfUint32((_305_v89).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_300_v86), (_index46).Int())
					_201_v1 = _201_v1
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _306_v90 _dafny.Sequence
		_ = _306_v90
		_306_v90 = _dafny.SeqOf(_201_v1, _201_v1, _201_v1)
		var _307_v91 D1
		_ = _307_v91
		_307_v91 = Companion_D1_.Create_DC5_(_200_v0, _261_v55, Companion_Default___.Fm27(_dafny.IntOfUint32((_306_v90).Cardinality()), _201_v1, _201_v1, globalState))
		var _pat_let_tv7 = _201_v1
		_ = _pat_let_tv7
		var _pat_let_tv8 = _201_v1
		_ = _pat_let_tv8
		r0 = func(_source6 D1) _dafny.Int {
			if _source6.Is_DC5() {
				var _308___mcc_h0 bool = _source6.Get_().(D1_DC5).Cf7
				_ = _308___mcc_h0
				var _309___mcc_h1 _dafny.Sequence = _source6.Get_().(D1_DC5).Cf8
				_ = _309___mcc_h1
				var _310___mcc_h2 _dafny.Sequence = _source6.Get_().(D1_DC5).Cf9
				_ = _310___mcc_h2
				var _311_cf9 _dafny.Sequence = _310___mcc_h2
				_ = _311_cf9
				var _312_cf8 _dafny.Sequence = _309___mcc_h1
				_ = _312_cf8
				var _313_cf7 bool = _308___mcc_h0
				_ = _313_cf7
				return _pat_let_tv7
			} else {
				var _314___mcc_h3 _dafny.Map = _source6.Get_().(D1_DC4).Cf6
				_ = _314___mcc_h3
				var _315_cf6 _dafny.Map = _314___mcc_h3
				_ = _315_cf6
				return _pat_let_tv8
			}
		}(_307_v91)
		return r0
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	dummy byte
}

func New_C2_() *C2 {
	_this := C2{}

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

func (_this *C2) Ctor__() {
	{
	}
}
func (_this *C2) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(226)
	}
}
func (_this *C2) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("fbtqkyvh")
	}
}
func (_this *C2) Fm14(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(354))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_316_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-772)
		})), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(944))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_317_i0 _dafny.Int) _dafny.Int {
			return (Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Dtor_cf0()
		})))
	}
}
func (_this *C2) M6(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		(globalState).F13 = !(p2) || (!(p2) || (p2))
		r1 = Companion_Default___.SafeModuloInt((Companion_Default___.Fm15(p2, p2, p2, globalState)).Cardinality(), (func() _dafny.Int {
			if p2 {
				return (_this).Fm0(_dafny.IntOfUint32((p0).Cardinality()), _dafny.CodePoint('p'), globalState)
			}
			return p3
		})())
		var _318_v0 _dafny.Set
		_ = _318_v0
		_318_v0 = _dafny.SetOf(p3)
		var _hi2 _dafny.Int = ((_318_v0).Intersection(_318_v0)).Cardinality()
		_ = _hi2
		for _319_i0 := p3; _319_i0.Cmp(_hi2) < 0; _319_i0 = _319_i0.Plus(_dafny.One) {
			var _320_v1 _dafny.Array
			_ = _320_v1
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw39
			_320_v1 = _nw39
			_320_v1 = _320_v1
			r1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(p3, _319_i0), p3)
			r1 = _319_i0
			var _321_v2 _dafny.Array
			_ = _321_v2
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw40
			_321_v2 = _nw40
			var _322_v3 _dafny.Array
			_ = _322_v3
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw41
			_322_v3 = _nw41
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_321_v2), 0))
			_ = _index47
			(_321_v2).ArraySet1((func() _dafny.Array {
				if p2 {
					return _322_v3
				}
				return _322_v3
			})(), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_321_v2), 0))
			_ = _index48
			(_321_v2).ArraySet1(_322_v3, (_index48).Int())
		}
		var _323_v4 _dafny.Map
		_ = _323_v4
		_323_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _324_v5 _dafny.Sequence
		_ = _324_v5
		_324_v5 = _dafny.SeqOf((_this).Fm14(_dafny.SetOf(p3, (_323_v4).Cardinality()), !(p2), p1, globalState), false, true)
		var _325_v6 _dafny.CodePoint
		_ = _325_v6
		_325_v6 = _dafny.CodePoint('a')
		if (_324_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mkjhatikf"), p0), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mkjhatikf"), p0)).Cardinality()))).Uint32(), _325_v6)).Cardinality()), _dafny.IntOfUint32((_324_v5).Cardinality()))).Uint32()).(bool) {
			var _326_v7 _dafny.Map
			_ = _326_v7
			_326_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
			var _327_v8 _dafny.Array
			_ = _327_v8
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw42
			_327_v8 = _nw42
			var _328_v9 _dafny.Map
			_ = _328_v9
			_328_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(-712))
			var _329_v10 _dafny.Map
			_ = _329_v10
			_329_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v8, (_328_v9).Cardinality())
			var _330_v11 _dafny.Array
			_ = _330_v11
			var _nwElement0_9 _dafny.Int = _dafny.IntOfInt64(-737)
			_ = _nwElement0_9
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(12))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_9, 0)
			(_nw43).ArraySet1((func() _dafny.Int {
				if (_326_v7).Contains(p2) {
					return (_326_v7).Get(p2).(_dafny.Int)
				}
				return p3
			})(), 1)
			(_nw43).ArraySet1(p3, 2)
			(_nw43).ArraySet1(p3, 3)
			(_nw43).ArraySet1(p3, 4)
			(_nw43).ArraySet1(p3, 5)
			(_nw43).ArraySet1(p3, 6)
			(_nw43).ArraySet1(p3, 7)
			(_nw43).ArraySet1(p3, 8)
			(_nw43).ArraySet1(p3, 9)
			(_nw43).ArraySet1((_329_v10).Cardinality(), 10)
			(_nw43).ArraySet1(p3, 11)
			_330_v11 = _nw43
			var _331_v12 _dafny.Map
			_ = _331_v12
			_331_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _327_v8)
			var _332_v13 _dafny.Map
			_ = _332_v13
			_332_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _325_v6)
			var _333_v15 _dafny.MultiSet
			_ = _333_v15
			_333_v15 = _dafny.MultiSetOf((func() _dafny.Array {
				if (_331_v12).Contains((_this).Fm14(func() _dafny.Set {
					var _coll8 = _dafny.NewBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((_332_v13).Keys().Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _334_v14 _dafny.Int
						_334_v14 = interface{}(_compr_8).(_dafny.Int)
						if (_332_v13).Contains(_334_v14) {
							_coll8.Add((_334_v14).Times(_dafny.IntOfInt64(290)))
						}
					}
					return _coll8.ToSet()
				}(), p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(900))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_335_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_336_i1 _dafny.Int) _dafny.CodePoint {
						return _335_v6
					}
				})(_325_v6))), globalState)) {
					return (_331_v12).Get((_this).Fm14(func() _dafny.Set {
						var _coll9 = _dafny.NewBuilder()
						_ = _coll9
						for _iter9 := _dafny.Iterate((_332_v13).Keys().Elements()); ; {
							_compr_9, _ok9 := _iter9()
							if !_ok9 {
								break
							}
							var _337_v14 _dafny.Int
							_337_v14 = interface{}(_compr_9).(_dafny.Int)
							if (_332_v13).Contains(_337_v14) {
								_coll9.Add((_337_v14).Times(_dafny.IntOfInt64(290)))
							}
						}
						return _coll9.ToSet()
					}(), p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(900))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg41 _dafny.Int) interface{} {
							return coer41(arg41)
						}
					}((func(_338_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_339_i1 _dafny.Int) _dafny.CodePoint {
							return _338_v6
						}
					})(_325_v6))), globalState)).(_dafny.Array)
				}
				return _330_v11
			})(), _330_v11, _327_v8, _330_v11, _330_v11)
			var _340_v16 _dafny.MultiSet
			_ = _340_v16
			_340_v16 = _dafny.MultiSetOf(p3)
			var _rhs61 _dafny.MultiSet = (_333_v15).Union(_333_v15)
			_ = _rhs61
			var _rhs62 _dafny.MultiSet = (_340_v16).Difference(Companion_Default___.Fm16(globalState))
			_ = _rhs62
			_333_v15 = _rhs61
			_340_v16 = _rhs62
			var _341_v17 D0
			_ = _341_v17
			_341_v17 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_this).Fm0(_dafny.IntOfInt64(144), _dafny.CodePoint('w'), globalState)))
			var _342_v18 _dafny.Map
			_ = _342_v18
			_342_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfInt64(698)), _341_v17)
			var _343_v20 _dafny.Map
			_ = _343_v20
			_343_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			_342_v18 = (_342_v18).Merge(func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate((_343_v20).Keys().Elements()); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _344_v19 _dafny.Int
					_344_v19 = interface{}(_compr_10).(_dafny.Int)
					if (_343_v20).Contains(_344_v19) {
						_coll10.Add(Companion_Default___.SafeDivisionInt(_344_v19, p3), _341_v17)
					}
				}
				return _coll10.ToMap()
			}())
			(globalState).F9 = p2
			var _345_v21 _dafny.Array
			_ = _345_v21
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(17))
			_ = _nw44
			_345_v21 = _nw44
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_345_v21), 0))
			_ = _index49
			(_345_v21).ArraySet1(Companion_Default___.Fm17(p2, globalState), (_index49).Int())
			var _346_v22 _dafny.Map
			_ = _346_v22
			_346_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _318_v0)
			var _347_v23 _dafny.Sequence
			_ = _347_v23
			_347_v23 = _dafny.SeqOf(p3)
			var _348_v25 _dafny.Sequence
			_ = _348_v25
			_348_v25 = _dafny.SeqOf((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Set {
				if (_346_v22).Contains(p2) {
					return (_346_v22).Get(p2).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll11 = _dafny.NewBuilder()
					_ = _coll11
					for _iter11 := _dafny.Iterate((_347_v23).Elements()); ; {
						_compr_11, _ok11 := _iter11()
						if !_ok11 {
							break
						}
						var _349_v24 _dafny.Int
						_349_v24 = interface{}(_compr_11).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_347_v23, _349_v24) {
							_coll11.Add((_349_v24).Plus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
						}
					}
					return _coll11.ToSet()
				}()
			})()).Cardinality(), p3), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-726))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_350_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_351_i2 _dafny.Int) _dafny.CodePoint {
					return _350_v6
				}
			})(_325_v6))), p1)
			var _352_v26 _dafny.Map
			_ = _352_v26
			_352_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(p3, _325_v6, globalState), _348_v25)
			var _353_v27 D0
			_ = _353_v27
			_353_v27 = Companion_D0_.Create_DC1_(p2, (func() _dafny.Sequence {
				if (_352_v26).Contains(p3) {
					return (_352_v26).Get(p3).(_dafny.Sequence)
				}
				return _dafny.SeqOf(p1)
			})(), _dafny.Companion_Sequence_.Concatenate(p0, p0))
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_345_v21), 0))
			_ = _index50
			(_345_v21).ArraySet1(_353_v27, (_index50).Int())
			var _354_v28 _dafny.Array
			_ = _354_v28
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_4
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_355_p2 bool) func(_dafny.Int) bool {
					return func(_356_i3 _dafny.Int) bool {
						return _355_p2
					}
				})(p2)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw45 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw45).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw45).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_354_v28 = _nw45
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_354_v28), 0))
			_ = _index51
			(_354_v28).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(65))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_357_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), p1), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_354_v28), 0))
			_ = _index52
			(_354_v28).ArraySet1(!((p2) || (p2)) || ((p3).Cmp(_dafny.IntOfUint32((_324_v5).Cardinality())) < 0), (_index52).Int())
		} else {
			var _rhs63 _dafny.Int = p3
			_ = _rhs63
			var _rhs64 bool = false
			_ = _rhs64
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			r1 = _rhs63
			_lhs46.F3 = _rhs64
			(globalState).F3 = p2
			r1 = p3
			var _358_v29 D0
			_ = _358_v29
			_358_v29 = Companion_D0_.Create_DC0_(p3)
			var _359_v30 _dafny.Sequence
			_ = _359_v30
			_359_v30 = _dafny.SeqOf(_358_v29)
			var _360_v31 _dafny.Sequence
			_ = _360_v31
			_360_v31 = _dafny.SeqOf(_359_v30)
			var _361_v32 _dafny.MultiSet
			_ = _361_v32
			_361_v32 = _dafny.MultiSetOf(p3, p3)
			var _362_v33 _dafny.Sequence
			_ = _362_v33
			_362_v33 = _dafny.SeqOf(_361_v32)
			_360_v31 = Companion_Default___.Fm18(_325_v6, ((_362_v33).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_362_v33).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), globalState)
			r1 = p3
		}
		var _363_v34 _dafny.MultiSet
		_ = _363_v34
		_363_v34 = _dafny.MultiSetOf(_318_v0, _318_v0)
		var _hi3 _dafny.Int = ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (func() _dafny.Int {
			if (_363_v34).Contains(_318_v0) {
				return (_363_v34).Multiplicity(_318_v0)
			}
			return p3
		})())).Cardinality())).Times((_dafny.Zero).Minus(p3))
		_ = _hi3
		for _364_i5 := p3; _364_i5.Cmp(_hi3) < 0; _364_i5 = _364_i5.Plus(_dafny.One) {
			_325_v6 = _325_v6
			_325_v6 = _dafny.CodePoint('c')
			var _365_v35 _dafny.Sequence
			_ = _365_v35
			_365_v35 = _dafny.SeqOf(_364_i5)
			var _366_v36 _dafny.Map
			_ = _366_v36
			_366_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_i5, _dafny.IntOfUint32((_365_v35).Cardinality()))
			var _367_v37 _dafny.Map
			_ = _367_v37
			_367_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_i5, (_366_v36).Cardinality())
			var _368_v38 _dafny.Sequence
			_ = _368_v38
			_368_v38 = _dafny.SeqOf(p0, p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(29))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_369_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_370_i6 _dafny.Int) _dafny.CodePoint {
					return _369_v6
				}
			})(_325_v6))), (_this).Fm1(_367_v37, globalState))
			var _371_v39 D0
			_ = _371_v39
			_371_v39 = Companion_D0_.Create_DC1_(p2, _368_v38, _dafny.UnicodeSeqOfUtf8Bytes("unnkkr"))
			var _372_v40 _dafny.MultiSet
			_ = _372_v40
			_372_v40 = _dafny.MultiSetOf(_371_v39)
			var _373_v41 *C0
			_ = _373_v41
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__((_372_v40).Update(_371_v39, Companion_Default___.Abs(_364_i5)), p2)
			_373_v41 = _nw46
			var _374_v42 _dafny.Sequence
			_ = _374_v42
			_374_v42 = _dafny.UnicodeSeqOfUtf8Bytes("qumwtaj")
			_374_v42 = _dafny.UnicodeSeqOfUtf8Bytes("juj")
		}
		var _hi4 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p3), p3)
		_ = _hi4
		for _375_i7 := p3; _375_i7.Cmp(_hi4) < 0; _375_i7 = _375_i7.Plus(_dafny.One) {
			var _376_v43 _dafny.Sequence
			_ = _376_v43
			_376_v43 = _dafny.SeqOf(p3, _375_i7)
			var _377_v44 _dafny.Array
			_ = _377_v44
			var _nwElement0_10 bool = p2
			_ = _nwElement0_10
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_10, 0)
			(_nw47).ArraySet1(!(p2), 1)
			(_nw47).ArraySet1(p2, 2)
			(_nw47).ArraySet1(p2, 3)
			_377_v44 = _nw47
			var _378_v45 _dafny.Map
			_ = _378_v45
			_378_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_376_v43).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p3), _dafny.IntOfUint32((_376_v43).Cardinality()))).Uint32()).(_dafny.Int), _377_v44)
			var _379_v46 _dafny.Map
			_ = _379_v46
			_379_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_376_v43).Cardinality()), _324_v5)
			_378_v45 = (_378_v45).Update(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_379_v46).Contains(_375_i7) {
					return (_379_v46).Get(_375_i7).(_dafny.Sequence)
				}
				return _324_v5
			})()).Cardinality()), _377_v44)
			_325_v6 = _325_v6
			r1 = (_this).Fm0(_375_i7, _325_v6, globalState)
			var _380_v47 D1
			_ = _380_v47
			_380_v47 = Companion_D1_.Create_DC5_(p2, p0, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2, p2), _324_v5))
			_380_v47 = Companion_D1_.Create_DC5_(p2, _dafny.Companion_Sequence_.Concatenate(p0, p0), _dafny.Companion_Sequence_.Concatenate(_324_v5, _324_v5))
		}
		var _381_v48 _dafny.Map
		_ = _381_v48
		_381_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, !(p2) || (true))
		var _382_v49 _dafny.Sequence
		_ = _382_v49
		_382_v49 = _dafny.SeqOf(p3)
		r0 = !((func() bool {
			if (_381_v48).Contains((_382_v49).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_382_v49).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_381_v48).Get((_382_v49).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_382_v49).Cardinality()))).Uint32()).(_dafny.Int)).(bool)
			}
			return !(p2)
		})())
		r1 = p3
		r2 = p2
		return r0, r1, r2
	}
}
func (_this *C2) M7(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 bool = false
		_ = r2
		var _383_v0 _dafny.CodePoint
		_ = _383_v0
		_383_v0 = _dafny.CodePoint('b')
		var _384_v1 _dafny.MultiSet
		_ = _384_v1
		_384_v1 = _dafny.MultiSetOf(_383_v0, _dafny.CodePoint('b'), _383_v0, _383_v0)
		var _385_v2 _dafny.Map
		_ = _385_v2
		_385_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_384_v1).Cardinality())
		var _386_v3 _dafny.MultiSet
		_ = _386_v3
		_386_v3 = _dafny.MultiSetOf(_385_v2)
		var _387_v4 _dafny.MultiSet
		_ = _387_v4
		_387_v4 = _dafny.MultiSetOf(!(p3), p3)
		var _388_v5 _dafny.Sequence
		_ = _388_v5
		_388_v5 = _dafny.SeqOf(true)
		var _389_v6 _dafny.Array
		_ = _389_v6
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw48
		_389_v6 = _nw48
		var _390_v7 _dafny.Map
		_ = _390_v7
		_390_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_388_v5).Cardinality()), _389_v6)
		var _391_v8 _dafny.Map
		_ = _391_v8
		_391_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cardinality(), p2)
		var _392_v9 _dafny.Sequence
		_ = _392_v9
		_392_v9 = _dafny.UnicodeSeqOfUtf8Bytes("uwerfkjaf")
		var _393_v10 _dafny.Map
		_ = _393_v10
		_393_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_392_v9).Cardinality()))), p3)
		var _394_v11 _dafny.Array
		_ = _394_v11
		var _nwElement0_11 _dafny.Int = (_dafny.Zero).Minus(p1)
		_ = _nwElement0_11
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(25))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_11, 0)
		(_nw49).ArraySet1((p2).Minus(p2), 1)
		(_nw49).ArraySet1((func() _dafny.Int {
			if p3 {
				return _dafny.IntOfInt64(435)
			}
			return p1
		})(), 2)
		(_nw49).ArraySet1(p2, 3)
		(_nw49).ArraySet1((p1).Plus(_dafny.IntOfInt64(766)), 4)
		(_nw49).ArraySet1((_386_v3).Cardinality(), 5)
		(_nw49).ArraySet1(p2, 6)
		(_nw49).ArraySet1(Companion_Default___.SafeModuloInt((_387_v4).Cardinality(), p2), 7)
		(_nw49).ArraySet1((_dafny.IntOfUint32((_388_v5).Cardinality())).Minus((_390_v7).Cardinality()), 8)
		(_nw49).ArraySet1((_dafny.Zero).Minus(p1), 9)
		(_nw49).ArraySet1((func() _dafny.Int {
			if (_391_v8).Contains((_dafny.Zero).Minus(p1)) {
				return (_391_v8).Get((_dafny.Zero).Minus(p1)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(970)
		})(), 10)
		(_nw49).ArraySet1(p1, 11)
		(_nw49).ArraySet1((p2).Times(p1), 12)
		(_nw49).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_392_v9, _392_v9)).Cardinality()), 13)
		(_nw49).ArraySet1(_dafny.IntOfInt64(-260), 14)
		(_nw49).ArraySet1(p2, 15)
		(_nw49).ArraySet1(p2, 16)
		(_nw49).ArraySet1((func() _dafny.Int {
			if (_385_v2).Contains(p3) {
				return (_385_v2).Get(p3).(_dafny.Int)
			}
			return p1
		})(), 17)
		(_nw49).ArraySet1(p2, 18)
		(_nw49).ArraySet1(_dafny.IntOfInt64(-272), 19)
		(_nw49).ArraySet1(p2, 20)
		(_nw49).ArraySet1(_dafny.IntOfInt64(156), 21)
		(_nw49).ArraySet1(p1, 22)
		(_nw49).ArraySet1((_393_v10).Cardinality(), 23)
		(_nw49).ArraySet1(p2, 24)
		_394_v11 = _nw49
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_5
		var _nw50 _dafny.Array
		_ = _nw50
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw50 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = (func(_395_p3 bool, _396_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_397_i0 _dafny.Int) _dafny.Int {
					return (_397_i0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_p3, _396_p1)).Cardinality())
				}
			})(p3, p1)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw50 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw50).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw50).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_394_v11 = _nw50
		var _398_i1 _dafny.Int
		_ = _398_i1
		_398_i1 = _dafny.Zero
		{
			for p3 {
				{
					if (_398_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_398_i1 = (_398_i1).Plus(_dafny.One)
					var _399_v12 _dafny.Set
					_ = _399_v12
					_399_v12 = _dafny.SetOf(p1)
					var _400_v13 _dafny.Map
					_ = _400_v13
					_400_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v12, _392_v9)
					var _401_v14 _dafny.Set
					_ = _401_v14
					_401_v14 = _dafny.SetOf((func() _dafny.Int {
						if false {
							return p1
						}
						return _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_400_v13).Contains(_399_v12) {
								return (_400_v13).Get(_399_v12).(_dafny.Sequence)
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("j")
						})()).Cardinality())
					})(), p1)
					_399_v12 = _399_v12
					var _402_v15 _dafny.Map
					_ = _402_v15
					_402_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm28(globalState))
					var _403_v16 D2
					_ = _403_v16
					_403_v16 = Companion_D2_.Create_DC7_(p3, (_402_v15).Merge(_402_v15), p1, p3)
					var _source7 D2 = _403_v16
					_ = _source7
					if _source7.Is_DC7() {
						var _404___mcc_h0 bool = _source7.Get_().(D2_DC7).Cf11
						_ = _404___mcc_h0
						var _405___mcc_h1 _dafny.Map = _source7.Get_().(D2_DC7).Cf12
						_ = _405___mcc_h1
						var _406___mcc_h2 _dafny.Int = _source7.Get_().(D2_DC7).Cf13
						_ = _406___mcc_h2
						var _407___mcc_h3 bool = _source7.Get_().(D2_DC7).Cf14
						_ = _407___mcc_h3
						var _408_cf14 bool = _407___mcc_h3
						_ = _408_cf14
						var _409_cf13 _dafny.Int = _406___mcc_h2
						_ = _409_cf13
						var _410_cf12 _dafny.Map = _405___mcc_h1
						_ = _410_cf12
						var _411_cf11 bool = _404___mcc_h0
						_ = _411_cf11
						_409_cf13 = _dafny.IntOfInt64(655)
						var _412_v17 _dafny.Sequence
						_ = _412_v17
						_412_v17 = _dafny.SeqOf(_392_v9, _392_v9)
						var _413_v18 D0
						_ = _413_v18
						_413_v18 = Companion_D0_.Create_DC1_(p3, _412_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(574))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_414_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_415_i2 _dafny.Int) _dafny.CodePoint {
								return _414_v0
							}
						})(_383_v0))))
						var _416_v19 _dafny.MultiSet
						_ = _416_v19
						_416_v19 = _dafny.MultiSetOf(_413_v18, _413_v18)
						var _417_v20 *C0
						_ = _417_v20
						var _nw51 *C0 = New_C0_()
						_ = _nw51
						_nw51.Ctor__(_416_v19, _411_cf11)
						_417_v20 = _nw51
						var _418_v21 _dafny.Map
						_ = _418_v21
						_418_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_cf11, (_417_v20).F19())
						(globalState).F5 = (func() bool {
							if (_418_v21).Contains(p3) {
								return (_418_v21).Get(p3).(bool)
							}
							return _411_cf11
						})()
						var _419_v22 _dafny.Map
						_ = _419_v22
						_419_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
						var _420_v23 D0
						_ = _420_v23
						_420_v23 = Companion_D0_.Create_DC2_(true)
						var _421_v24 _dafny.Array
						_ = _421_v24
						var _nwElement0_12 bool = (_dafny.IntOfUint32((_392_v9).Cardinality())).Cmp(_409_cf13) != 0
						_ = _nwElement0_12
						var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(26))
						_ = _nw52
						(_nw52).ArraySet1(_nwElement0_12, 0)
						(_nw52).ArraySet1(true, 1)
						(_nw52).ArraySet1(((_417_v20).F19()) == (false), 2)
						(_nw52).ArraySet1(_411_cf11, 3)
						(_nw52).ArraySet1((_393_v10).Equals(_393_v10), 4)
						(_nw52).ArraySet1(!((_417_v20).F19()), 5)
						(_nw52).ArraySet1(_408_cf14, 6)
						(_nw52).ArraySet1(_408_cf14, 7)
						(_nw52).ArraySet1((_409_cf13).Cmp((p0).Cardinality()) < 0, 8)
						(_nw52).ArraySet1(!(_408_cf14), 9)
						(_nw52).ArraySet1(p3, 10)
						(_nw52).ArraySet1(_411_cf11, 11)
						(_nw52).ArraySet1((p0).IsDisjointFrom((func() _dafny.Set {
							if (_419_v22).Contains(_408_cf14) {
								return (_419_v22).Get(_408_cf14).(_dafny.Set)
							}
							return p0
						})()), 12)
						(_nw52).ArraySet1(p3, 13)
						(_nw52).ArraySet1((_417_v20).F19(), 14)
						(_nw52).ArraySet1((_411_cf11) == (_411_cf11), 15)
						(_nw52).ArraySet1((_417_v20).F19(), 16)
						(_nw52).ArraySet1((!((_417_v20).F19())) == (_408_cf14), 17)
						(_nw52).ArraySet1(_408_cf14, 18)
						(_nw52).ArraySet1(p3, 19)
						(_nw52).ArraySet1(Companion_Default___.Fm19(_409_cf13, _411_cf11, globalState), 20)
						(_nw52).ArraySet1((_420_v23).Dtor_cf4(), 21)
						(_nw52).ArraySet1(!(p3), 22)
						(_nw52).ArraySet1(false, 23)
						(_nw52).ArraySet1(!((_399_v12).IsSubsetOf(_401_v14)), 24)
						(_nw52).ArraySet1((_417_v20).F19(), 25)
						_421_v24 = _nw52
						var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_421_v24), 0))
						_ = _index53
						(_421_v24).ArraySet1((_417_v20).F19(), (_index53).Int())
						var _422_v25 _dafny.Sequence
						_ = _422_v25
						_422_v25 = _dafny.SeqOf(p2, _dafny.IntOfUint32((_392_v9).Cardinality()), p1, p2, _409_cf13)
						var _423_v26 _dafny.Map
						_ = _423_v26
						_423_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(990), _392_v9)
						var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_421_v24), 0))
						_ = _index54
						(_421_v24).ArraySet1((_this).Fm14(_399_v12, _dafny.Companion_Sequence_.Contains(_422_v25, _dafny.IntOfUint32((_392_v9).Cardinality())), (func() _dafny.Sequence {
							if (_423_v26).Contains(p1) {
								return (_423_v26).Get(p1).(_dafny.Sequence)
							}
							return _392_v9
						})(), globalState), (_index54).Int())
					} else if _source7.Is_DC6() {
						var _424___mcc_h4 _dafny.MultiSet = _source7.Get_().(D2_DC6).Cf10
						_ = _424___mcc_h4
						var _425_cf10 _dafny.MultiSet = _424___mcc_h4
						_ = _425_cf10
						var _426_v27 _dafny.Sequence
						_ = _426_v27
						_426_v27 = _dafny.SeqOf(_dafny.IntOfInt64(-738))
						var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_389_v6), 0))
						_ = _index55
						(_389_v6).ArraySet1((p2).Times((_426_v27).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_426_v27).Cardinality()))).Uint32()).(_dafny.Int)), (_index55).Int())
						var _427_v28 D0
						_ = _427_v28
						_427_v28 = Companion_D0_.Create_DC0_(p1)
						var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_389_v6), 0))
						_ = _index56
						var _rhs65 _dafny.Sequence = _426_v27
						_ = _rhs65
						var _rhs66 bool = (p2).Cmp(p1) != 0
						_ = _rhs66
						var _rhs67 _dafny.Int = Companion_Default___.Fm22(_403_v16, Companion_D0_.Create_DC3_(_427_v28), p3, globalState)
						_ = _rhs67
						var _lhs47 _dafny.Array = _389_v6
						_ = _lhs47
						var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_389_v6), 0))
						_ = _lhs48
						_426_v27 = _rhs65
						r2 = _rhs66
						(_lhs47).ArraySet1(_rhs67, (_lhs48).Int())
						var _428_v29 bool
						_ = _428_v29
						var _429_v30 _dafny.Int
						_ = _429_v30
						var _430_v31 bool
						_ = _430_v31
						var _out2 bool
						_ = _out2
						var _out3 _dafny.Int
						_ = _out3
						var _out4 bool
						_ = _out4
						_out2, _out3, _out4 = (_this).M6(_392_v9, _392_v9, p3, p2, globalState)
						_428_v29 = _out2
						_429_v30 = _out3
						_430_v31 = _out4
						var _431_v32 D0
						_ = _431_v32
						_431_v32 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int)))
						var _pat_let_tv9 = _427_v28
						_ = _pat_let_tv9
						var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_389_v6), 0))
						_ = _index57
						(_389_v6).ArraySet1((Companion_Default___.Fm22(_403_v16, func(_pat_let4_0 D0) D0 {
							return func(_432_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let5_0 D0) D0 {
									return func(_433_dt__update_hcf5_h0 D0) D0 {
										return Companion_D0_.Create_DC3_(_433_dt__update_hcf5_h0)
									}(_pat_let5_0)
								}(_pat_let_tv9)
							}(_pat_let4_0)
						}(_431_v32), _428_v29, globalState)).Minus(_dafny.IntOfUint32((_388_v5).Cardinality())), (_index57).Int())
						var _434_v33 _dafny.Map
						_ = _434_v33
						_434_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_430_v31, _430_v31)
						var _435_v34 _dafny.Map
						_ = _435_v34
						_435_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v0, _434_v33)
						_383_v0 = Companion_Default___.Fm23(_430_v31, _435_v34, _428_v29, _dafny.IntOfInt64(-87), globalState)
					} else {
						var _436___mcc_h5 D2 = _source7.Get_().(D2_DC8).Cf15
						_ = _436___mcc_h5
						var _437_cf15 D2 = _436___mcc_h5
						_ = _437_cf15
						var _438_v35 *C1
						_ = _438_v35
						var _nw53 *C1 = New_C1_()
						_ = _nw53
						_nw53.Ctor__()
						_438_v35 = _nw53
						var _439_v36 _dafny.Map
						_ = _439_v36
						_439_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _383_v0)
						_439_v36 = (_439_v36).Update((p1).Minus(p1), _383_v0)
						(globalState).F9 = p3
						var _440_v37 _dafny.Array
						_ = _440_v37
						var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
						_ = _nw54
						_440_v37 = _nw54
						var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_440_v37), 0))
						_ = _index58
						(_440_v37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_392_v9, _392_v9), (_index58).Int())
						var _441_v38 _dafny.Map
						_ = _441_v38
						_441_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _392_v9)
						var _442_v39 _dafny.Sequence
						_ = _442_v39
						_442_v39 = _dafny.SeqOf(_dafny.SeqOf(_383_v0, _383_v0), _392_v9, (func() _dafny.Sequence {
							if (_441_v38).Contains(p3) {
								return (_441_v38).Get(p3).(_dafny.Sequence)
							}
							return _392_v9
						})(), _392_v9, _dafny.SeqOf(_383_v0))
						var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_440_v37), 0))
						_ = _index59
						(_440_v37).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_442_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).Fm0(p2, _383_v0, globalState)), _dafny.IntOfUint32((_442_v39).Cardinality()))).Uint32()).(_dafny.Sequence), _392_v9), (_index59).Int())
					}
					var _443_v40 _dafny.Sequence
					_ = _443_v40
					_443_v40 = _dafny.SeqOf(_392_v9)
					var _444_v41 D0
					_ = _444_v41
					_444_v41 = Companion_D0_.Create_DC1_(p3, _443_v40, _392_v9)
					var _445_v42 _dafny.MultiSet
					_ = _445_v42
					_445_v42 = _dafny.MultiSetOf(_444_v41, _444_v41)
					var _446_v43 *C0
					_ = _446_v43
					var _nw55 *C0 = New_C0_()
					_ = _nw55
					_nw55.Ctor__(_445_v42, p3)
					_446_v43 = _nw55
					var _447_v44 _dafny.Sequence
					_ = _447_v44
					_447_v44 = _dafny.SeqOf(p1, p2)
					var _448_v45 _dafny.Map
					_ = _448_v45
					_448_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_447_v44, p1)
					_448_v45 = (_448_v45).Update(_447_v44, p2)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _449_v46 _dafny.Int
		_ = _449_v46
		_449_v46 = _dafny.IntOfInt64(-694)
		_449_v46 = _dafny.IntOfInt64(278)
		var _450_v47 _dafny.Set
		_ = _450_v47
		_450_v47 = _dafny.SetOf(p1)
		var _source8 D0 = Companion_D0_.Create_DC0_((_450_v47).Cardinality())
		_ = _source8
		if _source8.Is_DC1() {
			var _451___mcc_h6 bool = _source8.Get_().(D0_DC1).Cf1
			_ = _451___mcc_h6
			var _452___mcc_h7 _dafny.Sequence = _source8.Get_().(D0_DC1).Cf2
			_ = _452___mcc_h7
			var _453___mcc_h8 _dafny.Sequence = _source8.Get_().(D0_DC1).Cf3
			_ = _453___mcc_h8
			var _454_cf3 _dafny.Sequence = _453___mcc_h8
			_ = _454_cf3
			var _455_cf2 _dafny.Sequence = _452___mcc_h7
			_ = _455_cf2
			var _456_cf1 bool = _451___mcc_h6
			_ = _456_cf1
			var _457_v48 _dafny.Array
			_ = _457_v48
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw56
			_457_v48 = _nw56
			var _458_v49 _dafny.Sequence
			_ = _458_v49
			_458_v49 = _dafny.SeqOf(p2, p1)
			var _459_v50 _dafny.Map
			_ = _459_v50
			_459_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(_456_cf1, _454_cf3, _388_v5), _458_v49)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_457_v48), 0))
			_ = _index60
			(_457_v48).ArraySet1(_459_v50, (_index60).Int())
			var _460_v51 D0
			_ = _460_v51
			_460_v51 = Companion_D0_.Create_DC1_(p3, _455_cf2, _454_cf3)
			var _461_v52 _dafny.Map
			_ = _461_v52
			_461_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC3_(_460_v51))
			var _462_v53 D2
			_ = _462_v53
			_462_v53 = Companion_D2_.Create_DC7_(p3, _461_v52, _dafny.IntOfUint32((_dafny.SeqOf(_456_cf1)).Cardinality()), _456_cf1)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _index61
			(_394_v11).ArraySet1(((_462_v53).Dtor_cf13()).Times(_449_v46), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_457_v48), 0))
			_ = _index62
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _index63
			var _rhs68 _dafny.Map = _459_v50
			_ = _rhs68
			var _rhs69 _dafny.Sequence = _454_cf3
			_ = _rhs69
			var _rhs70 _dafny.Int = _449_v46
			_ = _rhs70
			var _rhs71 bool = p3
			_ = _rhs71
			var _rhs72 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(((_450_v47).Cardinality()).Minus(p1), p2))
			_ = _rhs72
			var _lhs49 _dafny.Array = _457_v48
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_457_v48), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = _394_v11
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _lhs52
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			(_lhs49).ArraySet1(_rhs68, (_lhs50).Int())
			_454_cf3 = _rhs69
			(_lhs51).ArraySet1(_rhs70, (_lhs52).Int())
			_lhs53.F13 = _rhs71
			_449_v46 = _rhs72
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _index64
			(_394_v11).ArraySet1(Companion_Default___.SafeDivisionInt(_449_v46, Companion_Default___.SafeModuloInt(p1, p1)), (_index64).Int())
			var _463_v54 _dafny.Map
			_ = _463_v54
			_463_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_cf1, _389_v6))
			var _464_v55 _dafny.Map
			_ = _464_v55
			_464_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _389_v6)
			var _465_v56 D1
			_ = _465_v56
			_465_v56 = Companion_D1_.Create_DC4_((func() _dafny.Map {
				if (_463_v54).Contains(_456_cf1) {
					return (_463_v54).Get(_456_cf1).(_dafny.Map)
				}
				return _464_v55
			})())
			var _466_v57 _dafny.Map
			_ = _466_v57
			_466_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v56, (_462_v53).Dtor_cf13())
			_466_v57 = (_466_v57).Update(_465_v56, _dafny.IntOfUint32((_455_cf2).Cardinality()))
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _index65
			var _rhs73 _dafny.Map = (_385_v2).Merge(_385_v2)
			_ = _rhs73
			var _rhs74 _dafny.Int = _dafny.IntOfInt64(-205)
			_ = _rhs74
			var _rhs75 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_454_cf3, _392_v9)
			_ = _rhs75
			var _rhs76 _dafny.Sequence = (func() _dafny.Sequence {
				if _456_cf1 {
					return _388_v5
				}
				return _dafny.Companion_Sequence_.Concatenate(_388_v5, _388_v5)
			})()
			_ = _rhs76
			var _lhs54 _dafny.Array = _394_v11
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_394_v11), 0))
			_ = _lhs55
			_385_v2 = _rhs73
			(_lhs54).ArraySet1(_rhs74, (_lhs55).Int())
			_454_cf3 = _rhs75
			r0 = _rhs76
		} else if _source8.Is_DC2() {
			var _467___mcc_h9 bool = _source8.Get_().(D0_DC2).Cf4
			_ = _467___mcc_h9
			var _468_cf4 bool = _467___mcc_h9
			_ = _468_cf4
			var _469_v58 _dafny.Array
			_ = _469_v58
			var _nwElement0_13 bool = _468_cf4
			_ = _nwElement0_13
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(2))
			_ = _nw57
			(_nw57).ArraySet1(_nwElement0_13, 0)
			(_nw57).ArraySet1(Companion_Default___.Fm19(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-373))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}((func(_470_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_471_i3 _dafny.Int) _dafny.CodePoint {
					return _470_v0
				}
			})(_383_v0)))).Cardinality()), _468_cf4, globalState), 1)
			_469_v58 = _nw57
			var _472_v59 _dafny.Sequence
			_ = _472_v59
			_472_v59 = _dafny.SeqOf(_392_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_473_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_474_i4 _dafny.Int) _dafny.CodePoint {
					return _473_v0
				}
			})(_383_v0))), _dafny.UnicodeSeqOfUtf8Bytes("t"))
			var _475_v60 _dafny.Sequence
			_ = _475_v60
			_475_v60 = _dafny.SeqOf(_dafny.IntOfUint32((_472_v59).Cardinality()))
			var _476_v61 _dafny.Map
			_ = _476_v61
			_476_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_v58, _dafny.Companion_Sequence_.Concatenate(_475_v60, _475_v60))
			_449_v46 = (_dafny.Zero).Minus((_476_v61).Cardinality())
			var _477_v62 _dafny.Array
			_ = _477_v62
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw58
			_477_v62 = _nw58
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))
			_ = _index66
			(_477_v62).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_392_v9, _392_v9), (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))
			_ = _index67
			(_389_v6).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_cf4, p3)).Cardinality(), (p0).Cardinality()), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))
			_ = _index68
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))
			_ = _index69
			var _rhs77 _dafny.Array = _394_v11
			_ = _rhs77
			var _rhs78 _dafny.Sequence = _392_v9
			_ = _rhs78
			var _rhs79 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_475_v60).Cardinality()), (_385_v2).Cardinality())).Minus(p2)))
			_ = _rhs79
			var _rhs80 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("dwq")
			_ = _rhs80
			var _lhs56 _dafny.Array = _477_v62
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))
			_ = _lhs57
			var _lhs58 _dafny.Array = _389_v6
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))
			_ = _lhs59
			_389_v6 = _rhs77
			(_lhs56).ArraySet1(_rhs78, (_lhs57).Int())
			(_lhs58).ArraySet1(_rhs79, (_lhs59).Int())
			_392_v9 = _rhs80
			var _478_v63 D0
			_ = _478_v63
			_478_v63 = Companion_D0_.Create_DC0_((p1).Plus(p2))
			var _source9 D0 = _478_v63
			_ = _source9
			if _source9.Is_DC1() {
				var _479___mcc_h12 bool = _source9.Get_().(D0_DC1).Cf1
				_ = _479___mcc_h12
				var _480___mcc_h13 _dafny.Sequence = _source9.Get_().(D0_DC1).Cf2
				_ = _480___mcc_h13
				var _481___mcc_h14 _dafny.Sequence = _source9.Get_().(D0_DC1).Cf3
				_ = _481___mcc_h14
				var _482_cf3 _dafny.Sequence = _481___mcc_h14
				_ = _482_cf3
				var _483_cf2 _dafny.Sequence = _480___mcc_h13
				_ = _483_cf2
				var _484_cf1 bool = _479___mcc_h12
				_ = _484_cf1
				(globalState).F9 = _484_cf1
				var _485_v64 D0
				_ = _485_v64
				_485_v64 = Companion_D0_.Create_DC1_(_484_cf1, _472_v59, (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence))
				var _486_v65 D0
				_ = _486_v65
				_486_v65 = Companion_D0_.Create_DC3_(_485_v64)
				var _487_v66 D0
				_ = _487_v66
				_487_v66 = Companion_D0_.Create_DC3_(_485_v64)
				var _488_v67 D2
				_ = _488_v67
				_488_v67 = Companion_D2_.Create_DC7_(_468_cf4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_cf1, _487_v66), _dafny.IntOfInt64(779), _484_cf1)
				_449_v46 = Companion_Default___.Fm22(_488_v67, _487_v66, (_388_v5).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_449_v46)).Cardinality(), _dafny.IntOfUint32((_388_v5).Cardinality()))).Uint32()).(bool), globalState)
				var _489_v68 _dafny.MultiSet
				_ = _489_v68
				_489_v68 = _dafny.MultiSetOf(_394_v11, _394_v11)
				var _490_v69 _dafny.Sequence
				_ = _490_v69
				_490_v69 = _dafny.SeqOf((_489_v68).Update(_394_v11, Companion_Default___.Abs(_449_v46)))
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))
				_ = _index70
				(_389_v6).ArraySet1((_dafny.Zero).Minus(((_490_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.IntOfUint32((_490_v69).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()), (_index70).Int())
				_469_v58 = _469_v58
			} else if _source9.Is_DC2() {
				var _491___mcc_h15 bool = _source9.Get_().(D0_DC2).Cf4
				_ = _491___mcc_h15
				var _492_cf4 bool = _491___mcc_h15
				_ = _492_cf4
				var _493_v70 D0
				_ = _493_v70
				_493_v70 = Companion_D0_.Create_DC0_(p1)
				var _494_v71 D0
				_ = _494_v71
				_494_v71 = Companion_D0_.Create_DC3_(_493_v70)
				var _495_v72 _dafny.Map
				_ = _495_v72
				_495_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _494_v71)
				var _496_v73 D2
				_ = _496_v73
				_496_v73 = Companion_D2_.Create_DC7_(p3, _495_v72, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(804))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_497_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_498_i5 _dafny.Int) _dafny.Sequence {
						return _497_v5
					}
				})(_388_v5)))).Cardinality())), p3)
				var _499_v74 _dafny.Set
				_ = _499_v74
				_499_v74 = _dafny.SetOf(_383_v0)
				var _pat_let_tv10 = p3
				_ = _pat_let_tv10
				var _pat_let_tv11 = _495_v72
				_ = _pat_let_tv11
				var _pat_let_tv12 = _388_v5
				_ = _pat_let_tv12
				var _pat_let_tv13 = _492_cf4
				_ = _pat_let_tv13
				var _pat_let_tv14 = p3
				_ = _pat_let_tv14
				var _500_v77 _dafny.Array
				_ = _500_v77
				var _nwElement0_14 D2 = _496_v73
				_ = _nwElement0_14
				var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(26))
				_ = _nw59
				(_nw59).ArraySet1(_nwElement0_14, 0)
				(_nw59).ArraySet1(_496_v73, 1)
				(_nw59).ArraySet1(_496_v73, 2)
				(_nw59).ArraySet1(func(_pat_let6_0 D2) D2 {
					return func(_501_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let7_0 bool) D2 {
							return func(_502_dt__update_hcf11_h0 bool) D2 {
								return func(_pat_let8_0 _dafny.Map) D2 {
									return func(_503_dt__update_hcf12_h0 _dafny.Map) D2 {
										return Companion_D2_.Create_DC7_(_502_dt__update_hcf11_h0, _503_dt__update_hcf12_h0, (_501_dt__update__tmp_h1).Dtor_cf13(), (_501_dt__update__tmp_h1).Dtor_cf14())
									}(_pat_let8_0)
								}(_pat_let_tv11)
							}(_pat_let7_0)
						}(_pat_let_tv10)
					}(_pat_let6_0)
				}(_496_v73), 3)
				(_nw59).ArraySet1(_496_v73, 4)
				(_nw59).ArraySet1(_496_v73, 5)
				(_nw59).ArraySet1(func(_pat_let9_0 D2) D2 {
					return func(_504_dt__update__tmp_h2 D2) D2 {
						return func(_pat_let10_0 _dafny.Int) D2 {
							return func(_505_dt__update_hcf13_h0 _dafny.Int) D2 {
								return func(_pat_let11_0 bool) D2 {
									return func(_506_dt__update_hcf14_h0 bool) D2 {
										return func(_pat_let12_0 bool) D2 {
											return func(_507_dt__update_hcf11_h1 bool) D2 {
												return Companion_D2_.Create_DC7_(_507_dt__update_hcf11_h1, (_504_dt__update__tmp_h2).Dtor_cf12(), _505_dt__update_hcf13_h0, _506_dt__update_hcf14_h0)
											}(_pat_let12_0)
										}(_pat_let_tv14)
									}(_pat_let11_0)
								}(_pat_let_tv13)
							}(_pat_let10_0)
						}(_dafny.IntOfUint32((_pat_let_tv12).Cardinality()))
					}(_pat_let9_0)
				}(_496_v73), 6)
				(_nw59).ArraySet1(_496_v73, 7)
				(_nw59).ArraySet1(_496_v73, 8)
				(_nw59).ArraySet1(_496_v73, 9)
				(_nw59).ArraySet1(_496_v73, 10)
				(_nw59).ArraySet1(_496_v73, 11)
				(_nw59).ArraySet1(_496_v73, 12)
				(_nw59).ArraySet1(_496_v73, 13)
				(_nw59).ArraySet1(_496_v73, 14)
				(_nw59).ArraySet1(_496_v73, 15)
				(_nw59).ArraySet1(Companion_D2_.Create_DC7_(p3, (_495_v72).Update(p3, _494_v71), (_499_v74).Cardinality(), _492_cf4), 16)
				(_nw59).ArraySet1(Companion_Default___.Fm29((_385_v2).Cardinality(), (_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int), (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence), p1, globalState), 17)
				(_nw59).ArraySet1(Companion_D2_.Create_DC7_((_388_v5).Select((Companion_Default___.SafeIndex(_449_v46, _dafny.IntOfUint32((_388_v5).Cardinality()))).Uint32()).(bool), _495_v72, (_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int), p3), 18)
				(_nw59).ArraySet1(_496_v73, 19)
				(_nw59).ArraySet1(_496_v73, 20)
				(_nw59).ArraySet1(_496_v73, 21)
				(_nw59).ArraySet1(_496_v73, 22)
				(_nw59).ArraySet1((func() D2 {
					if _492_cf4 {
						return _496_v73
					}
					return Companion_D2_.Create_DC7_((_this).Fm14(func() _dafny.Set {
						var _coll12 = _dafny.NewBuilder()
						_ = _coll12
						for _iter12 := _dafny.Iterate((func() _dafny.Map {
							var _coll13 = _dafny.NewMapBuilder()
							_ = _coll13
							for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(288), _dafny.IntOfInt64(-460))); ; {
								_compr_13, _ok13 := _iter13()
								if !_ok13 {
									break
								}
								var _508_v75 _dafny.Int
								_508_v75 = interface{}(_compr_13).(_dafny.Int)
								if ((_dafny.IntOfInt64(288)).Cmp(_508_v75) <= 0) && ((_508_v75).Cmp(_dafny.IntOfInt64(-460)) < 0) {
									_coll13.Add((_508_v75).Minus(p1), _492_cf4)
								}
							}
							return _coll13.ToMap()
						}()).Keys().Elements()); ; {
							_compr_12, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _509_v76 _dafny.Int
							_509_v76 = interface{}(_compr_12).(_dafny.Int)
							if (func() _dafny.Map {
								var _coll14 = _dafny.NewMapBuilder()
								_ = _coll14
								for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(288), _dafny.IntOfInt64(-460))); ; {
									_compr_14, _ok14 := _iter14()
									if !_ok14 {
										break
									}
									var _508_v75 _dafny.Int
									_508_v75 = interface{}(_compr_14).(_dafny.Int)
									if ((_dafny.IntOfInt64(288)).Cmp(_508_v75) <= 0) && ((_508_v75).Cmp(_dafny.IntOfInt64(-460)) < 0) {
										_coll14.Add((_508_v75).Minus(p1), _492_cf4)
									}
								}
								return _coll14.ToMap()
							}()).Contains(_509_v76) {
								_coll12.Add((_509_v76).Minus(_dafny.IntOfInt64(-977)))
							}
						}
						return _coll12.ToSet()
					}(), false, _392_v9, globalState), _495_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(globalState), _492_cf4)).Cardinality(), p3)
				})(), 23)
				(_nw59).ArraySet1(_496_v73, 24)
				(_nw59).ArraySet1(_496_v73, 25)
				_500_v77 = _nw59
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_500_v77), 0))
				_ = _index71
				(_500_v77).ArraySet1(Companion_D2_.Create_DC7_(p3, _495_v72, (_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int), p3), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_500_v77), 0))
				_ = _index72
				(_500_v77).ArraySet1(Companion_Default___.Fm29((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_388_v5, _388_v5)).Cardinality())), (p2).Plus((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int)), _392_v9, Companion_Default___.SafeDivisionInt((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int), (_496_v73).Dtor_cf13()), globalState), (_index72).Int())
				var _510_v78 _dafny.Map
				_ = _510_v78
				_510_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gortewya"), !(_492_cf4))
				var _511_v79 D0
				_ = _511_v79
				_511_v79 = Companion_D0_.Create_DC1_(_492_cf4, _472_v59, (_472_v59).Select((Companion_Default___.SafeIndex(_449_v46, _dafny.IntOfUint32((_472_v59).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _512_v80 _dafny.Set
				_ = _512_v80
				_512_v80 = _dafny.SetOf(p3)
				_510_v78 = (_510_v78).Update((_511_v79).Dtor_cf3(), (_512_v80).IsSubsetOf(p0))
				_394_v11 = _394_v11
				var _513_v81 _dafny.MultiSet
				_ = _513_v81
				_513_v81 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_(_492_cf4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_514_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_515_i6 _dafny.Int) _dafny.Sequence {
						return _514_v9
					}
				})(_392_v9))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_516_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_517_i7 _dafny.Int) _dafny.CodePoint {
						return _516_v0
					}
				})(_383_v0)))))
				var _518_v82 _dafny.Sequence
				_ = _518_v82
				_518_v82 = _dafny.SeqOf(_513_v81, _513_v81)
				var _519_v83 _dafny.MultiSet
				_ = _519_v83
				_519_v83 = _dafny.MultiSetOf((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int))
				var _520_v84 _dafny.Sequence
				_ = _520_v84
				_520_v84 = _dafny.SeqOf((_518_v82).Select((Companion_Default___.SafeIndex((_519_v83).Cardinality(), _dafny.IntOfUint32((_518_v82).Cardinality()))).Uint32()).(_dafny.MultiSet), _513_v81)
				var _521_v85 _dafny.Map
				_ = _521_v85
				_521_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _518_v82)
				_521_v85 = (_521_v85).Update(((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int)).Plus(_449_v46), _518_v82)
			} else if _source9.Is_DC0() {
				var _522___mcc_h16 _dafny.Int = _source9.Get_().(D0_DC0).Cf0
				_ = _522___mcc_h16
				var _523_cf0 _dafny.Int = _522___mcc_h16
				_ = _523_cf0
				var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw60
				_394_v11 = _nw60
				var _rhs81 _dafny.CodePoint = _383_v0
				_ = _rhs81
				var _rhs82 _dafny.MultiSet = (_dafny.MultiSetOf(_383_v0)).Difference((_384_v1).Union(_384_v1))
				_ = _rhs82
				var _rhs83 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence), (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence))).Cardinality())
				_ = _rhs83
				var _rhs84 _dafny.Int = (_dafny.Zero).Minus(_449_v46)
				_ = _rhs84
				var _rhs85 _dafny.Int = (_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int)
				_ = _rhs85
				_383_v0 = _rhs81
				_384_v1 = _rhs82
				_449_v46 = _rhs83
				_449_v46 = _rhs84
				_449_v46 = _rhs85
				var _524_v86 _dafny.Map
				_ = _524_v86
				_524_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_393_v10).Cardinality(), _472_v59)
				var _525_v87 D0
				_ = _525_v87
				_525_v87 = Companion_D0_.Create_DC1_(p3, _dafny.SeqOf(_392_v9), (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v46, (_391_v8).Cardinality()), globalState))
				var _526_v88 _dafny.MultiSet
				_ = _526_v88
				_526_v88 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_(p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(638))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_527_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_528_i8 _dafny.Int) _dafny.Sequence {
						return _527_v9
					}
				})(_392_v9))), _dafny.UnicodeSeqOfUtf8Bytes("cv")), Companion_D0_.Create_DC1_(_468_cf4, _472_v59, (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence)), Companion_Default___.Fm17(p3, globalState), Companion_D0_.Create_DC1_(p3, (func() _dafny.Sequence {
					if (_524_v86).Contains(_449_v46) {
						return (_524_v86).Get(_449_v46).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg52 _dafny.Int) interface{} {
							return coer52(arg52)
						}
					}((func(_529_v62 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
						return func(_530_i9 _dafny.Int) _dafny.Sequence {
							return (_529_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_529_v62), 0))).Int()).(_dafny.Sequence)
						}
					})(_477_v62)))
				})(), _392_v9), _525_v87)
				var _531_v89 *C0
				_ = _531_v89
				var _nw61 *C0 = New_C0_()
				_ = _nw61
				_nw61.Ctor__(_526_v88, _468_cf4)
				_531_v89 = _nw61
				var _532_v90 _dafny.Sequence
				_ = _532_v90
				_532_v90 = _dafny.SeqOf(Companion_Default___.Fm17((_531_v89).F19(), globalState))
				var _533_v91 *C0
				_ = _533_v91
				var _nw62 *C0 = New_C0_()
				_ = _nw62
				_nw62.Ctor__((_dafny.MultiSetFromSeq(_532_v90)).Update(Companion_Default___.Fm17((_531_v89).F19(), globalState), Companion_Default___.Abs(_449_v46)), _468_cf4)
				_533_v91 = _nw62
			} else {
				var _534___mcc_h17 D0 = _source9.Get_().(D0_DC3).Cf5
				_ = _534___mcc_h17
				var _535_cf5 D0 = _534___mcc_h17
				_ = _535_cf5
				var _536_v92 _dafny.Map
				_ = _536_v92
				_536_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm19(_dafny.IntOfUint32((_475_v60).Cardinality()), p3, globalState))
				var _537_v93 D0
				_ = _537_v93
				_537_v93 = Companion_D0_.Create_DC1_(_468_cf4, _472_v59, _392_v9)
				var _538_v94 _dafny.Map
				_ = _538_v94
				_538_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_cf4, Companion_D0_.Create_DC3_(_537_v93))
				var _539_v95 D2
				_ = _539_v95
				_539_v95 = Companion_D2_.Create_DC7_((func() bool {
					if (_393_v10).Contains((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}((func(_540_v6 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_541_i10 _dafny.Int) _dafny.Int {
							return (_540_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_540_v6), 0))).Int()).(_dafny.Int)
						}
					})(_389_v6))))).Cardinality()) {
						return (_393_v10).Get((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}((func(_542_v6 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_543_i10 _dafny.Int) _dafny.Int {
								return (_542_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_542_v6), 0))).Int()).(_dafny.Int)
							}
						})(_389_v6))))).Cardinality()).(bool)
					}
					return p3
				})(), _538_v94, _449_v46, !(_468_cf4))
				var _544_v96 _dafny.Sequence
				_ = _544_v96
				_544_v96 = _dafny.SeqOf(_539_v95)
				var _545_v97 _dafny.Map
				_ = _545_v97
				_545_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _469_v58)
				var _546_v98 _dafny.Map
				_ = _546_v98
				_546_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_536_v92).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_475_v60).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_536_v92).Cardinality())).Cardinality()))).Uint32(), _dafny.IntOfUint32((_544_v96).Cardinality())), (_545_v97).Cardinality())
				var _547_v99 _dafny.Array
				_ = _547_v99
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw63
				_547_v99 = _nw63
				var _548_v100 _dafny.Array
				_ = _548_v100
				var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw64
				_548_v100 = _nw64
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_547_v99), 0))
				_ = _index73
				(_547_v99).ArraySet1(_548_v100, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_547_v99), 0))
				_ = _index74
				var _rhs86 _dafny.Map = _546_v98
				_ = _rhs86
				var _rhs87 _dafny.Array = _548_v100
				_ = _rhs87
				var _lhs60 _dafny.Array = _547_v99
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_547_v99), 0))
				_ = _lhs61
				_546_v98 = _rhs86
				(_lhs60).ArraySet1(_rhs87, (_lhs61).Int())
				var _549_v101 _dafny.Map
				_ = _549_v101
				_549_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence), _475_v60)
				var _550_v102 _dafny.Array
				_ = _550_v102
				var _nwElement0_15 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_475_v60, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_475_v60).Cardinality()))).Uint32(), p1)
				_ = _nwElement0_15
				var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(10))
				_ = _nw65
				(_nw65).ArraySet1(_nwElement0_15, 0)
				(_nw65).ArraySet1(_475_v60, 1)
				(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_475_v60, _475_v60), 2)
				(_nw65).ArraySet1((func() _dafny.Sequence {
					if false {
						return _475_v60
					}
					return Companion_Default___.Fm30(globalState)
				})(), 3)
				(_nw65).ArraySet1((func() _dafny.Sequence {
					if (_549_v101).Contains(Companion_Default___.Fm20(p1, p3, _468_cf4, _449_v46, globalState)) {
						return (_549_v101).Get(Companion_Default___.Fm20(p1, p3, _468_cf4, _449_v46, globalState)).(_dafny.Sequence)
					}
					return _475_v60
				})(), 4)
				(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-514))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_551_v60 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_552_i11 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_551_v60).Cardinality())
					}
				})(_475_v60))), 5)
				(_nw65).ArraySet1(_475_v60, 6)
				(_nw65).ArraySet1(_475_v60, 7)
				(_nw65).ArraySet1(_475_v60, 8)
				(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_389_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v6), 0))).Int()).(_dafny.Int)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(224))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_553_i12 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-161)
				}))), 9)
				_550_v102 = _nw65
				(globalState).F2 = _550_v102
				var _554_v103 D1
				_ = _554_v103
				_554_v103 = Companion_D1_.Create_DC5_(!(p3), (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence), _388_v5)
				var _rhs88 _dafny.Int = (((_dafny.Zero).Minus(_dafny.IntOfUint32((_392_v9).Cardinality()))).Plus(p2)).Times(p1)
				_ = _rhs88
				var _rhs89 bool = (_554_v103).Dtor_cf7()
				_ = _rhs89
				var _rhs90 bool = _468_cf4
				_ = _rhs90
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_449_v46 = _rhs88
				_lhs62.F13 = _rhs89
				_lhs63.F5 = _rhs90
				var _555_v104 _dafny.Map
				_ = _555_v104
				_555_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _394_v11)
				var _556_v105 D1
				_ = _556_v105
				_556_v105 = Companion_D1_.Create_DC4_(_555_v104)
				var _557_v106 _dafny.Map
				_ = _557_v106
				_557_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_475_v60, _556_v105)
				var _558_v107 _dafny.Map
				_ = _558_v107
				_558_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_557_v106, p3)
				_558_v107 = (_558_v107).Update(_557_v106, p3)
			}
			var _559_v108 _dafny.Array
			_ = _559_v108
			var _nw66 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
			_ = _nw66
			_559_v108 = _nw66
			var _560_v109 D0
			_ = _560_v109
			_560_v109 = Companion_D0_.Create_DC1_(_468_cf4, _472_v59, (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence))
			var _561_v110 D0
			_ = _561_v110
			_561_v110 = Companion_D0_.Create_DC1_(false, (_560_v109).Dtor_cf2(), (_477_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_477_v62), 0))).Int()).(_dafny.Sequence))
			var _562_v111 _dafny.MultiSet
			_ = _562_v111
			_562_v111 = _dafny.MultiSetOf(_561_v110)
			var _563_v112 *C0
			_ = _563_v112
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__(_562_v111, (_560_v109).Dtor_cf1())
			_563_v112 = _nw67
			var _564_v113 _dafny.Map
			_ = _564_v113
			_564_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_559_v108, _563_v112)
			var _565_v114 _dafny.Map
			_ = _565_v114
			_565_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v113, p1)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_469_v58), 0))
			_ = _index75
			(_469_v58).ArraySet1(((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).Fm0((func() _dafny.Int {
				if (_565_v114).Contains(_564_v113) {
					return (_565_v114).Get(_564_v113).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-858)
			})(), _383_v0, globalState)))).Cmp(p1) != 0, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_469_v58), 0))
			_ = _index76
			(_469_v58).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ooskv"), Companion_Default___.Fm20(_449_v46, _468_cf4, (_563_v112).F19(), _449_v46, globalState)), _dafny.UnicodeSeqOfUtf8Bytes("xfq")), (_index76).Int())
		} else if _source8.Is_DC0() {
			var _566___mcc_h10 _dafny.Int = _source8.Get_().(D0_DC0).Cf0
			_ = _566___mcc_h10
			var _567_cf0 _dafny.Int = _566___mcc_h10
			_ = _567_cf0
			if p3 {
				var _568_v115 _dafny.Sequence
				_ = _568_v115
				_568_v115 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_569_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_570_i13 _dafny.Int) _dafny.CodePoint {
						return _569_v0
					}
				})(_383_v0))), _dafny.Companion_Sequence_.Update(_392_v9, (Companion_Default___.SafeIndex(_567_cf0, _dafny.IntOfUint32((_392_v9).Cardinality()))).Uint32(), _383_v0))
				var _571_v116 D0
				_ = _571_v116
				_571_v116 = Companion_D0_.Create_DC1_(p3, _568_v115, _392_v9)
				var _572_v117 D0
				_ = _572_v117
				_572_v117 = Companion_D0_.Create_DC3_(_571_v116)
				var _573_v118 D0
				_ = _573_v118
				_573_v118 = Companion_D0_.Create_DC3_(_571_v116)
				var _574_v119 D2
				_ = _574_v119
				_574_v119 = Companion_D2_.Create_DC7_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19((_450_v47).Cardinality(), p3, globalState), _573_v118), p2, p3)
				_392_v9 = (_this).Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v46, (_574_v119).Dtor_cf13())).Merge(Companion_Default___.Fm31(globalState)), globalState)
				var _575_v120 D0
				_ = _575_v120
				_575_v120 = Companion_D0_.Create_DC1_(p3, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_576_i14 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("uert")
				})), (Companion_Default___.SafeIndex(_449_v46, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}(func(_577_i14 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("uert")
				}))).Cardinality()))).Uint32(), _392_v9), _392_v9)
				var _578_v121 _dafny.MultiSet
				_ = _578_v121
				_578_v121 = _dafny.MultiSetOf(_575_v120)
				var _579_v122 _dafny.Array
				_ = _579_v122
				var _nw68 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw68
				_579_v122 = _nw68
				var _580_v123 _dafny.Map
				_ = _580_v123
				_580_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_579_v122, (_387_v4).Cardinality())
				var _581_v124 *C0
				_ = _581_v124
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__((_578_v121).Update(Companion_D0_.Create_DC1_(p3, _568_v115, _dafny.UnicodeSeqOfUtf8Bytes("xbt")), Companion_Default___.Abs((_580_v123).Cardinality())), p3)
				_581_v124 = _nw69
				var _582_v125 bool
				_ = _582_v125
				var _583_v126 _dafny.Int
				_ = _583_v126
				var _584_v127 bool
				_ = _584_v127
				var _out5 bool
				_ = _out5
				var _out6 _dafny.Int
				_ = _out6
				var _out7 bool
				_ = _out7
				_out5, _out6, _out7 = (_this).M6(_392_v9, _392_v9, p3, _dafny.IntOfInt64(163), globalState)
				_582_v125 = _out5
				_583_v126 = _out6
				_584_v127 = _out7
				var _585_v128 bool
				_ = _585_v128
				var _586_v129 _dafny.Int
				_ = _586_v129
				var _587_v130 bool
				_ = _587_v130
				var _out8 bool
				_ = _out8
				var _out9 _dafny.Int
				_ = _out9
				var _out10 bool
				_ = _out10
				_out8, _out9, _out10 = (_this).M6(_392_v9, _392_v9, _582_v125, p2, globalState)
				_585_v128 = _out8
				_586_v129 = _out9
				_587_v130 = _out10
				_583_v126 = (Companion_Default___.SafeModuloInt(_567_cf0, (_dafny.Zero).Minus(_567_cf0))).Plus(_567_cf0)
			} else {
				_449_v46 = p1
				var _588_v131 _dafny.Array
				_ = _588_v131
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
				_ = _nw70
				_588_v131 = _nw70
				var _rhs91 _dafny.Array = _588_v131
				_ = _rhs91
				var _rhs92 bool = p3
				_ = _rhs92
				var _lhs64 *GlobalState = globalState
				_ = _lhs64
				_588_v131 = _rhs91
				_lhs64.F13 = _rhs92
				(globalState).F3 = p3
				var _589_v132 D1
				_ = _589_v132
				_589_v132 = Companion_D1_.Create_DC5_(p3, _392_v9, _388_v5)
				var _590_v133 _dafny.Map
				_ = _590_v133
				_590_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_589_v132, _392_v9)
				_590_v133 = _590_v133
				var _591_v134 _dafny.Set
				_ = _591_v134
				_591_v134 = p0
				var _592_v135 _dafny.Sequence
				_ = _592_v135
				_592_v135 = _dafny.SeqOf(_392_v9)
				var _593_v136 D0
				_ = _593_v136
				_593_v136 = Companion_D0_.Create_DC1_(p3, _592_v135, _392_v9)
				var _594_v137 _dafny.MultiSet
				_ = _594_v137
				_594_v137 = _dafny.MultiSetOf(_593_v136)
				var _595_v138 *C0
				_ = _595_v138
				var _nw71 *C0 = New_C0_()
				_ = _nw71
				_nw71.Ctor__(_594_v137, (func() bool {
					if (_393_v10).Contains((_387_v4).Cardinality()) {
						return (_393_v10).Get((_387_v4).Cardinality()).(bool)
					}
					return p3
				})())
				_595_v138 = _nw71
				var _596_v139 _dafny.Sequence
				_ = _596_v139
				_596_v139 = _dafny.SeqOf(_595_v138)
				var _rhs93 _dafny.Set = (p0).Union(p0)
				_ = _rhs93
				var _rhs94 _dafny.Array = _389_v6
				_ = _rhs94
				var _rhs95 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_596_v139, _596_v139)
				_ = _rhs95
				_591_v134 = _rhs93
				_394_v11 = _rhs94
				_596_v139 = _rhs95
			}
			var _597_v140 _dafny.Map
			_ = _597_v140
			_597_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("astkcoj"), p3)
			var _598_v142 _dafny.Map
			_ = _598_v142
			_598_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("v"), _dafny.IntOfUint32((_388_v5).Cardinality()))
			_597_v140 = func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate(((_598_v142).Update(_dafny.UnicodeSeqOfUtf8Bytes("rjqeb"), _dafny.IntOfUint32((_392_v9).Cardinality()))).Keys().Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _599_v141 _dafny.Sequence
					_599_v141 = interface{}(_compr_15).(_dafny.Sequence)
					if ((_598_v142).Update(_dafny.UnicodeSeqOfUtf8Bytes("rjqeb"), _dafny.IntOfUint32((_392_v9).Cardinality()))).Contains(_599_v141) {
						_coll15.Add(_599_v141, p3)
					}
				}
				return _coll15.ToMap()
			}()
			var _600_v143 _dafny.Map
			_ = _600_v143
			_600_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _601_v144 _dafny.Set
			_ = _601_v144
			_601_v144 = _dafny.SetOf((func() bool {
				if (_600_v143).Contains((_this).Fm14(_dafny.SetOf(p1), p3, _392_v9, globalState)) {
					return (_600_v143).Get((_this).Fm14(_dafny.SetOf(p1), p3, _392_v9, globalState)).(bool)
				}
				return p3
			})(), p3)
			_601_v144 = p0
			var _602_v145 _dafny.Sequence
			_ = _602_v145
			_602_v145 = _dafny.SeqOf(_392_v9, _392_v9)
			var _603_v146 D0
			_ = _603_v146
			_603_v146 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32(((Companion_D0_.Create_DC1_(!(p3), _dafny.Companion_Sequence_.Update(_602_v145, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_602_v145).Cardinality()))).Uint32(), _392_v9), _392_v9)).Dtor_cf3()).Cardinality()))
			var _604_v147 D0
			_ = _604_v147
			_604_v147 = Companion_D0_.Create_DC3_(_603_v146)
			_604_v147 = _604_v147
		} else {
			var _605___mcc_h11 D0 = _source8.Get_().(D0_DC3).Cf5
			_ = _605___mcc_h11
			var _606_cf5 D0 = _605___mcc_h11
			_ = _606_cf5
			r2 = _dafny.Companion_Sequence_.IsProperPrefixOf(_392_v9, _dafny.Companion_Sequence_.Concatenate(_392_v9, _dafny.UnicodeSeqOfUtf8Bytes("ienvw")))
			_393_v10 = (_393_v10).Update(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(-748)), (true) && (p3))
			var _607_v148 _dafny.Array
			_ = _607_v148
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw72
			_607_v148 = _nw72
			var _608_v149 _dafny.Map
			_ = _608_v149
			_608_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _392_v9)
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_607_v148), 0))
			_ = _index77
			(_607_v148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_392_v9, (func() _dafny.Sequence {
				if (_608_v149).Contains(true) {
					return (_608_v149).Get(true).(_dafny.Sequence)
				}
				return _392_v9
			})()), (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_607_v148), 0))
			_ = _index78
			(_607_v148).ArraySet1(Companion_Default___.Fm20(p2, p3, _dafny.Companion_Sequence_.IsProperPrefixOf(_392_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}((func(_609_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_610_i15 _dafny.Int) _dafny.CodePoint {
					return _609_v0
				}
			})(_383_v0)))), p1, globalState), (_index78).Int())
			(globalState).F3 = p3
		}
		var _611_v150 _dafny.Array
		_ = _611_v150
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_6
		var _nw73 _dafny.Array
		_ = _nw73
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw73 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) bool = (func(_612_p3 bool, _613_v9 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_614_i16 _dafny.Int) bool {
					return (Companion_D0_.Create_DC1_(_612_p3, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("usbo")), _613_v9)).Dtor_cf1()
				}
			})(p3, _392_v9)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw73 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw73).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw73).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_611_v150 = _nw73
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))
		_ = _index79
		(_611_v150).ArraySet1(p3, (_index79).Int())
		var _615_v151 _dafny.Map
		_ = _615_v151
		_615_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
		var _616_v152 _dafny.Map
		_ = _616_v152
		_616_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v151, p3)
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_611_v150), 0))
		_ = _index80
		(_611_v150).ArraySet1((_616_v152).Contains(_615_v151), (_index80).Int())
		var _617_v153 _dafny.MultiSet
		_ = _617_v153
		_617_v153 = _dafny.MultiSetOf(_dafny.IntOfInt64(-21), (_dafny.Zero).Minus((_dafny.Zero).Minus(_449_v46)))
		var _618_v154 _dafny.Sequence
		_ = _618_v154
		_618_v154 = _dafny.SeqOf(_617_v153)
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))
		_ = _index81
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_611_v150), 0))
		_ = _index82
		var _rhs96 bool = p3
		_ = _rhs96
		var _rhs97 bool = (_this).Fm14(_450_v47, p3, _392_v9, globalState)
		_ = _rhs97
		var _rhs98 bool = !_dafny.Companion_Sequence_.Equal(_618_v154, (func() _dafny.Sequence {
			if p3 {
				return _dafny.SeqOf(_dafny.MultiSetOf((_391_v8).Cardinality(), p2, p1), _617_v153)
			}
			return _618_v154
		})())
		_ = _rhs98
		var _rhs99 _dafny.Int = _449_v46
		_ = _rhs99
		var _lhs65 _dafny.Array = _611_v150
		_ = _lhs65
		var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))
		_ = _lhs66
		var _lhs67 _dafny.Array = _611_v150
		_ = _lhs67
		var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_611_v150), 0))
		_ = _lhs68
		var _lhs69 *GlobalState = globalState
		_ = _lhs69
		(_lhs65).ArraySet1(_rhs96, (_lhs66).Int())
		(_lhs67).ArraySet1(_rhs97, (_lhs68).Int())
		_lhs69.F5 = _rhs98
		_449_v46 = _rhs99
		var _619_v155 D0
		_ = _619_v155
		_619_v155 = Companion_D0_.Create_DC0_(p2)
		var _620_v156 D0
		_ = _620_v156
		_620_v156 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(_619_v155))
		var _621_v157 _dafny.Map
		_ = _621_v157
		_621_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _620_v156)
		var _622_v158 _dafny.Set
		_ = _622_v158
		_622_v158 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), Companion_Default___.Fm22(Companion_D2_.Create_DC7_((_611_v150).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))).Int()).(bool), _621_v157, (_dafny.MultiSetOf(p3, false)).Cardinality(), p3), Companion_Default___.Fm28(globalState), (_611_v150).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))).Int()).(bool), globalState)), _391_v8)
		var _623_v159 _dafny.Sequence
		_ = _623_v159
		_623_v159 = _dafny.SeqOf(_392_v9, _392_v9)
		var _624_v160 D0
		_ = _624_v160
		_624_v160 = Companion_D0_.Create_DC1_((_611_v150).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))).Int()).(bool), _623_v159, _392_v9)
		var _625_v161 D2
		_ = _625_v161
		_625_v161 = Companion_D2_.Create_DC6_(_dafny.MultiSetOf(_624_v160, Companion_D0_.Create_DC1_((_611_v150).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))).Int()).(bool), _623_v159, _dafny.UnicodeSeqOfUtf8Bytes("bfwqt")), _624_v160))
		var _pat_let_tv15 = _622_v158
		_ = _pat_let_tv15
		var _pat_let_tv16 = _622_v158
		_ = _pat_let_tv16
		var _pat_let_tv17 = _622_v158
		_ = _pat_let_tv17
		_622_v158 = func(_source10 D2) _dafny.Set {
			if _source10.Is_DC7() {
				var _626___mcc_h18 bool = _source10.Get_().(D2_DC7).Cf11
				_ = _626___mcc_h18
				var _627___mcc_h19 _dafny.Map = _source10.Get_().(D2_DC7).Cf12
				_ = _627___mcc_h19
				var _628___mcc_h20 _dafny.Int = _source10.Get_().(D2_DC7).Cf13
				_ = _628___mcc_h20
				var _629___mcc_h21 bool = _source10.Get_().(D2_DC7).Cf14
				_ = _629___mcc_h21
				var _630_cf14 bool = _629___mcc_h21
				_ = _630_cf14
				var _631_cf13 _dafny.Int = _628___mcc_h20
				_ = _631_cf13
				var _632_cf12 _dafny.Map = _627___mcc_h19
				_ = _632_cf12
				var _633_cf11 bool = _626___mcc_h18
				_ = _633_cf11
				return _pat_let_tv15
			} else if _source10.Is_DC6() {
				var _634___mcc_h22 _dafny.MultiSet = _source10.Get_().(D2_DC6).Cf10
				_ = _634___mcc_h22
				var _635_cf10 _dafny.MultiSet = _634___mcc_h22
				_ = _635_cf10
				return _pat_let_tv16
			} else {
				var _636___mcc_h23 D2 = _source10.Get_().(D2_DC8).Cf15
				_ = _636___mcc_h23
				var _637_cf15 D2 = _636___mcc_h23
				_ = _637_cf15
				return _pat_let_tv17
			}
		}(_625_v161)
		r0 = _dafny.Companion_Sequence_.Concatenate(_388_v5, _dafny.Companion_Sequence_.Concatenate(_388_v5, _388_v5))
		r1 = _388_v5
		r2 = (_611_v150).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_611_v150), 0))).Int()).(bool)
		return r0, r1, r2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f16 bool
	_f17 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f16 = false
	_this._f17 = _dafny.Zero
	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__(f16 bool, f17 _dafny.Int) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C3) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_this).F17())
	}
}
func (_this *C3) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("vcy")
	}
}
func (_this *C3) Fm13(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.SeqOf((_this).F16(), (_this).F16()), _dafny.SeqOf((_this).F16()), _dafny.SeqOf((_this).F16()))).Intersection(_dafny.SetOf(_dafny.SeqOf((_this).F16()), _dafny.SeqOf((_this).F16(), (_this).F16(), (_this).F16(), (_this).F16(), (_this).F16())))).IsDisjointFrom((func() _dafny.Set {
			var _coll16 = _dafny.NewBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}(func(_638_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf((_this).F16())
			}))).Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _639_v0 _dafny.Sequence
				_639_v0 = interface{}(_compr_16).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_638_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F16())
				})), _639_v0) {
					_coll16.Add(_639_v0)
				}
			}
			return _coll16.ToSet()
		}()).Intersection(_dafny.SetOf(_dafny.SeqOf((_this).F16()))))
	}
}
func (_this *C3) M0(p0 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _640_v0 *C2
		_ = _640_v0
		var _nw74 *C2 = New_C2_()
		_ = _nw74
		_nw74.Ctor__()
		_640_v0 = _nw74
		var _641_v1 _dafny.Array
		_ = _641_v1
		var _nw75 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw75
		_641_v1 = _nw75
		_641_v1 = _641_v1
		r1 = (_this).F17()
		(globalState).F3 = ((_this).F16()) || ((_this).F16())
		var _642_v2 _dafny.Array
		_ = _642_v2
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw76
		_642_v2 = _nw76
		var _643_v3 _dafny.Map
		_ = _643_v3
		_643_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_642_v2, (_this).F16())
		var _644_v4 _dafny.Sequence
		_ = _644_v4
		_644_v4 = _dafny.UnicodeSeqOfUtf8Bytes("eehtvwmi")
		var _645_v5 _dafny.Sequence
		_ = _645_v5
		_645_v5 = _dafny.SeqOf((_this).F16())
		var _646_v6 D1
		_ = _646_v6
		_646_v6 = Companion_D1_.Create_DC5_((_this).F16(), _dafny.UnicodeSeqOfUtf8Bytes("dr"), _645_v5)
		var _647_v7 _dafny.Map
		_ = _647_v7
		_647_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v3, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_this).F16() {
				return _644_v4
			}
			return (_646_v6).Dtor_cf8()
		})()).Cardinality()))
		_647_v7 = (_647_v7).Update(_643_v3, (_this).Fm0((_this).F17(), p0, globalState))
		var _648_v8 *C1
		_ = _648_v8
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__()
		_648_v8 = _nw77
		r0 = (_this).F16()
		var _649_v9 _dafny.Sequence
		_ = _649_v9
		_649_v9 = _dafny.SeqOf(_644_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(828))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg63 _dafny.Int) interface{} {
				return coer63(arg63)
			}
		}((func(_650_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_651_i0 _dafny.Int) _dafny.CodePoint {
				return _650_p0
			}
		})(p0))))
		var _652_v10 D0
		_ = _652_v10
		_652_v10 = Companion_D0_.Create_DC1_(true, _649_v9, _644_v4)
		var _653_v11 _dafny.MultiSet
		_ = _653_v11
		_653_v11 = _dafny.MultiSetOf((func(_pat_let13_0 D0) D0 {
			return func(_654_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let14_0 bool) D0 {
					return func(_655_dt__update_hcf1_h0 bool) D0 {
						return Companion_D0_.Create_DC1_(_655_dt__update_hcf1_h0, (_654_dt__update__tmp_h0).Dtor_cf2(), (_654_dt__update__tmp_h0).Dtor_cf3())
					}(_pat_let14_0)
				}((_this).F16())
			}(_pat_let13_0)
		}(_652_v10)).Dtor_cf3())
		r1 = ((_this).F17()).Plus((_653_v11).Cardinality())
		return r0, r1
	}
}
func (_this *C3) F16() bool {
	{
		return _this._f16
	}
}
func (_this *C3) F17() _dafny.Int {
	{
		return _this._f17
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f14 _dafny.Int
	_f15 _dafny.Map
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.EmptyMap
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f14 _dafny.Int, f15 _dafny.Map) {
	{
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C4) Fm7(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_this).F14()))).Dtor_cf0()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_dafny.SetOf((_this).F14())).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _656_v0 _dafny.Int
				_656_v0 = interface{}(_compr_17).(_dafny.Int)
				if (_dafny.SetOf((_this).F14())).Contains(_656_v0) {
					_coll17.Add(Companion_Default___.SafeModuloInt(_656_v0, _dafny.IntOfInt64(-958)), false)
				}
			}
			return _coll17.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(832), (_this).F14()), _dafny.SeqOf((_this).F14())))
	}
}
func (_this *C4) Fm8(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return !((_dafny.SetOf((_this).F14())).IsDisjointFrom(_dafny.SetOf((_this).F14(), (func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_dafny.SetOf((_this).F14())).Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _657_v0 _dafny.Int
				_657_v0 = interface{}(_compr_18).(_dafny.Int)
				if (_dafny.SetOf((_this).F14())).Contains(_657_v0) {
					_coll18.Add(Companion_Default___.SafeDivisionInt(_657_v0, (_this).F14()), true)
				}
			}
			return _coll18.ToMap()
		}()).Cardinality(), (_dafny.Zero).Minus((_this).F14()), (_dafny.MultiSetOf((_this).F14())).Cardinality()))) || (_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("axmcm"), _dafny.UnicodeSeqOfUtf8Bytes("elfoqgqbn")))
	}
}
func (_this *C4) M5(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _658_v0 bool
		_ = _658_v0
		_658_v0 = false
		var _659_i0 _dafny.Int
		_ = _659_i0
		_659_i0 = _dafny.Zero
		{
			for _658_v0 {
				{
					if (_659_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_659_i0 = (_659_i0).Plus(_dafny.One)
					var _660_v1 D0
					_ = _660_v1
					_660_v1 = Companion_D0_.Create_DC0_(Companion_Default___.SafeDivisionInt((_this).F14(), (_this).F14()))
					var _661_v2 _dafny.Map
					_ = _661_v2
					_661_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), !(_658_v0))
					var _662_v3 _dafny.Sequence
					_ = _662_v3
					_662_v3 = _dafny.SeqOf((func() bool {
						if (_661_v2).Contains((_this).F14()) {
							return (_661_v2).Get((_this).F14()).(bool)
						}
						return _658_v0
					})(), _658_v0)
					var _pat_let_tv18 = _662_v3
					_ = _pat_let_tv18
					_660_v1 = func(_pat_let15_0 D0) D0 {
						return func(_663_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let16_0 _dafny.Int) D0 {
								return func(_664_dt__update_hcf0_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC0_(_664_dt__update_hcf0_h0)
								}(_pat_let16_0)
							}(_dafny.IntOfUint32((_pat_let_tv18).Cardinality()))
						}(_pat_let15_0)
					}(_660_v1)
					var _665_v4 _dafny.Int
					_ = _665_v4
					_665_v4 = _dafny.IntOfInt64(665)
					_665_v4 = (_this).F14()
					var _666_v5 _dafny.Map
					_ = _666_v5
					_666_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, _658_v0)
					if (func() bool {
						if !((_dafny.SetOf((_this).F14())).IsDisjointFrom(_dafny.SetOf((_dafny.Zero).Minus(_665_v4), _665_v4, (_666_v5).Cardinality(), (_this).F14()))) {
							return _658_v0
						}
						return _658_v0
					})() {
						var _667_v6 _dafny.Map
						_ = _667_v6
						_667_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, _665_v4)
						var _668_v7 _dafny.MultiSet
						_ = _668_v7
						_668_v7 = _dafny.MultiSetOf((_this).F14(), _665_v4)
						var _669_v8 _dafny.Map
						_ = _669_v8
						_669_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_658_v0, _658_v0, _658_v0), (_this).F14())
						var _670_v9 _dafny.Map
						_ = _670_v9
						_670_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, _669_v8)
						var _671_v10 _dafny.Array
						_ = _671_v10
						var _nwElement0_16 _dafny.Int = (func() _dafny.Int {
							if (_667_v6).Contains(_658_v0) {
								return (_667_v6).Get(_658_v0).(_dafny.Int)
							}
							return (_this).F14()
						})()
						_ = _nwElement0_16
						var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(26))
						_ = _nw78
						(_nw78).ArraySet1(_nwElement0_16, 0)
						(_nw78).ArraySet1((_this).F14(), 1)
						(_nw78).ArraySet1(((_668_v7).Cardinality()).Minus((_this).F14()), 2)
						(_nw78).ArraySet1((_this).F14(), 3)
						(_nw78).ArraySet1((_660_v1).Dtor_cf0(), 4)
						(_nw78).ArraySet1((_this).F14(), 5)
						(_nw78).ArraySet1(_dafny.IntOfInt64(321), 6)
						(_nw78).ArraySet1((_this).F14(), 7)
						(_nw78).ArraySet1(((func() _dafny.Map {
							if (_670_v9).Contains(false) {
								return (_670_v9).Get(false).(_dafny.Map)
							}
							return _669_v8
						})()).Cardinality(), 8)
						(_nw78).ArraySet1((_this).F14(), 9)
						(_nw78).ArraySet1(Companion_Default___.SafeModuloInt(_665_v4, _dafny.IntOfInt64(841)), 10)
						(_nw78).ArraySet1((_this).F14(), 11)
						(_nw78).ArraySet1((_this).F14(), 12)
						(_nw78).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm9(_dafny.IntOfInt64(-178), _662_v3, p0, globalState), _662_v3)).Cardinality()), 13)
						(_nw78).ArraySet1(_665_v4, 14)
						(_nw78).ArraySet1((_this).F14(), 15)
						(_nw78).ArraySet1((Companion_Default___.Fm10(_658_v0, _dafny.IntOfInt64(569), (_this).Fm8(_665_v4, _658_v0, _658_v0, globalState), !(_658_v0), globalState)).Cardinality(), 16)
						(_nw78).ArraySet1(((_668_v7).Cardinality()).Minus((_dafny.Zero).Minus(Companion_Default___.Fm11(_658_v0, globalState))), 17)
						(_nw78).ArraySet1(_665_v4, 18)
						(_nw78).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_665_v4), _665_v4), 19)
						(_nw78).ArraySet1(_665_v4, 20)
						(_nw78).ArraySet1((_this).F14(), 21)
						(_nw78).ArraySet1((_dafny.Zero).Minus((_665_v4).Minus(_665_v4)), 22)
						(_nw78).ArraySet1(((_this).F14()).Plus((_this).F14()), 23)
						(_nw78).ArraySet1((_dafny.Zero).Minus(_665_v4), 24)
						(_nw78).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sdnbh")).Cardinality()), 25)
						_671_v10 = _nw78
						var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_671_v10), 0))
						_ = _index83
						(_671_v10).ArraySet1(_665_v4, (_index83).Int())
						var _672_v11 _dafny.MultiSet
						_ = _672_v11
						_672_v11 = _dafny.MultiSetOf(Companion_Default___.Fm12(_658_v0, _658_v0, (_this).F14(), globalState))
						var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_671_v10), 0))
						_ = _index84
						(_671_v10).ArraySet1((func() _dafny.Set {
							var _coll19 = _dafny.NewBuilder()
							_ = _coll19
							for _iter19 := _dafny.Iterate((_672_v11).Elements()); ; {
								_compr_19, _ok19 := _iter19()
								if !_ok19 {
									break
								}
								var _673_v12 D0
								_673_v12 = interface{}(_compr_19).(D0)
								if (_672_v11).Contains(_673_v12) {
									_coll19.Add(_673_v12)
								}
							}
							return _coll19.ToSet()
						}()).Cardinality(), (_index84).Int())
						(globalState).F13 = _658_v0
						(globalState).F5 = _658_v0
						var _674_v13 *C2
						_ = _674_v13
						var _nw79 *C2 = New_C2_()
						_ = _nw79
						_nw79.Ctor__()
						_674_v13 = _nw79
						var _675_v14 _dafny.Set
						_ = _675_v14
						_675_v14 = _dafny.SetOf(_658_v0, (_667_v6).Equals(Companion_Default___.Fm15(_658_v0, _658_v0, _658_v0, globalState)))
						var _676_v15 D1
						_ = _676_v15
						_676_v15 = Companion_D1_.Create_DC5_(_658_v0, p0, _dafny.SeqOf(true))
						var _677_v16 _dafny.Array
						_ = _677_v16
						var _len0_7 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_7
						var _nw80 _dafny.Array
						_ = _nw80
						if _len0_7.Cmp(_dafny.Zero) == 0 {
							_nw80 = _dafny.NewArray(_len0_7)
						} else {
							var _init7 func(_dafny.Int) bool = (func(_678_v0 bool) func(_dafny.Int) bool {
								return func(_679_i1 _dafny.Int) bool {
									return _678_v0
								}
							})(_658_v0)
							_ = _init7
							var _element0_7 = _init7(_dafny.Zero)
							_ = _element0_7
							_nw80 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
							(_nw80).ArraySet1(_element0_7, 0)
							var _nativeLen0_7 = (_len0_7).Int()
							_ = _nativeLen0_7
							for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
								(_nw80).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
							}
						}
						_677_v16 = _nw80
						var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_677_v16), 0))
						_ = _index85
						(_677_v16).ArraySet1((_662_v3).Select((Companion_Default___.SafeIndex(_665_v4, _dafny.IntOfUint32((_662_v3).Cardinality()))).Uint32()).(bool), (_index85).Int())
						var _680_v17 _dafny.Set
						_ = _680_v17
						_680_v17 = _dafny.SetOf((_this).F14(), _665_v4, (_this).F14(), (_dafny.Zero).Minus((_this).F14()))
						var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_677_v16), 0))
						_ = _index86
						var _rhs100 _dafny.Set = _675_v14
						_ = _rhs100
						var _rhs101 D1 = _676_v15
						_ = _rhs101
						var _rhs102 bool = (_674_v13).Fm14(_680_v17, _658_v0, p0, globalState)
						_ = _rhs102
						var _lhs70 _dafny.Array = _677_v16
						_ = _lhs70
						var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_677_v16), 0))
						_ = _lhs71
						_675_v14 = _rhs100
						_676_v15 = _rhs101
						(_lhs70).ArraySet1(_rhs102, (_lhs71).Int())
					} else {
						var _681_v18 *C2
						_ = _681_v18
						var _nw81 *C2 = New_C2_()
						_ = _nw81
						_nw81.Ctor__()
						_681_v18 = _nw81
						_681_v18 = _681_v18
						var _682_v19 _dafny.Array
						_ = _682_v19
						var _nw82 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw82
						_682_v19 = _nw82
						var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))
						_ = _index87
						(_682_v19).ArraySet1(_658_v0, (_index87).Int())
						var _683_v20 _dafny.Sequence
						_ = _683_v20
						_683_v20 = _dafny.UnicodeSeqOfUtf8Bytes("rtdk")
						var _684_v21 _dafny.Sequence
						_ = _684_v21
						_684_v21 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_682_v19, _682_v19)).Cardinality()), (_this).F14(), (_this).F14(), _665_v4, _dafny.IntOfUint32((_662_v3).Cardinality()))
						var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))
						_ = _index88
						var _rhs103 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_684_v21, _684_v21)
						_ = _rhs103
						var _rhs104 _dafny.Sequence = p0
						_ = _rhs104
						var _lhs72 _dafny.Array = _682_v19
						_ = _lhs72
						var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))
						_ = _lhs73
						(_lhs72).ArraySet1(_rhs103, (_lhs73).Int())
						_683_v20 = _rhs104
						var _685_v23 _dafny.MultiSet
						_ = _685_v23
						_685_v23 = _dafny.MultiSetOf((func() _dafny.Set {
							var _coll20 = _dafny.NewBuilder()
							_ = _coll20
							for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(720), _dafny.IntOfInt64(-698))); ; {
								_compr_20, _ok20 := _iter20()
								if !_ok20 {
									break
								}
								var _686_v22 _dafny.Int
								_686_v22 = interface{}(_compr_20).(_dafny.Int)
								if ((_dafny.IntOfInt64(720)).Cmp(_686_v22) <= 0) && ((_686_v22).Cmp(_dafny.IntOfInt64(-698)) < 0) {
									_coll20.Add(Companion_Default___.SafeDivisionInt(_686_v22, (_this).F14()))
								}
							}
							return _coll20.ToSet()
						}()).Cardinality())
						_685_v23 = _685_v23
						var _687_v24 _dafny.Sequence
						_ = _687_v24
						_687_v24 = _dafny.SeqOf(_683_v20)
						(globalState).F13 = !(_dafny.Companion_Sequence_.Equal(_687_v24, _687_v24))
						var _688_v25 _dafny.Map
						_ = _688_v25
						_688_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_682_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))).Int()).(bool), (_682_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))).Int()).(bool), (_682_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_682_v19), 0))).Int()).(bool)), _682_v19)
						_682_v19 = (func() _dafny.Array {
							if (_688_v25).Contains(_662_v3) {
								return (_688_v25).Get(_662_v3).(_dafny.Array)
							}
							return _682_v19
						})()
					}
					var _689_v26 _dafny.Sequence
					_ = _689_v26
					_689_v26 = _dafny.UnicodeSeqOfUtf8Bytes("y")
					var _690_v27 _dafny.Set
					_ = _690_v27
					_690_v27 = _dafny.SetOf((_dafny.Zero).Minus(_665_v4), (_this).F14())
					var _691_v28 _dafny.Map
					_ = _691_v28
					_691_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, _665_v4)
					var _692_v29 _dafny.Map
					_ = _692_v29
					_692_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_690_v27, (_691_v28).Cardinality())
					var _rhs105 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_689_v26, _689_v26), _689_v26)
					_ = _rhs105
					var _rhs106 _dafny.Map = (_692_v29).Merge((_692_v29).Merge(func() _dafny.Map {
						var _coll21 = _dafny.NewMapBuilder()
						_ = _coll21
						for _iter21 := _dafny.Iterate((_692_v29).Keys().Elements()); ; {
							_compr_21, _ok21 := _iter21()
							if !_ok21 {
								break
							}
							var _693_v30 _dafny.Set
							_693_v30 = interface{}(_compr_21).(_dafny.Set)
							if (_692_v29).Contains(_693_v30) {
								_coll21.Add(_693_v30, _dafny.IntOfInt64(615))
							}
						}
						return _coll21.ToMap()
					}()))
					_ = _rhs106
					_689_v26 = _rhs105
					_692_v29 = _rhs106
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _694_v31 _dafny.Array
		_ = _694_v31
		var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw83
		_694_v31 = _nw83
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_694_v31), 0))); ; {
			_guard_loop_0, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _695_i2 _dafny.Int
			_695_i2 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_695_i2).Sign() != -1) && ((_695_i2).Cmp(_dafny.ArrayLen((_694_v31), 0)) < 0)) {
				(_694_v31).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(213), _dafny.IntOfInt64(746))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(469), (_this).F14())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), (_this).F14()))), (_695_i2).Int())
			}
		}
		var _hi5 _dafny.Int = (_this).F14()
		_ = _hi5
		for _696_i3 := (Companion_Default___.Fm32(true, _658_v0, p0, _658_v0, globalState)).Cardinality(); _696_i3.Cmp(_hi5) < 0; _696_i3 = _696_i3.Plus(_dafny.One) {
			var _697_v32 _dafny.Sequence
			_ = _697_v32
			_697_v32 = _dafny.SeqOf(_658_v0, _dafny.Companion_Sequence_.IsPrefixOf(p0, p0))
			if (_697_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.IntOfUint32((_697_v32).Cardinality()))).Uint32()).(bool) {
				var _698_v33 _dafny.Int
				_ = _698_v33
				_698_v33 = _dafny.IntOfInt64(412)
				_698_v33 = _dafny.IntOfInt64(491)
				var _699_v34 _dafny.Array
				_ = _699_v34
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw84
				_699_v34 = _nw84
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_699_v34), 0))
				_ = _index89
				(_699_v34).ArraySet1(_696_i3, (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_699_v34), 0))
				_ = _index90
				(_699_v34).ArraySet1(_696_i3, (_index90).Int())
				var _700_v35 _dafny.Set
				_ = _700_v35
				_700_v35 = _dafny.SetOf(Companion_Default___.Fm11(false, globalState))
				var _701_v36 D4
				_ = _701_v36
				_701_v36 = Companion_D4_.Create_DC10_(_700_v35)
				var _702_v37 _dafny.Map
				_ = _702_v37
				_702_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_700_v35).IsDisjointFrom((_701_v36).Dtor_cf17())), _696_i3)
				_702_v37 = (_702_v37).Update(_658_v0, _698_v33)
				var _703_v38 *C2
				_ = _703_v38
				var _nw85 *C2 = New_C2_()
				_ = _nw85
				_nw85.Ctor__()
				_703_v38 = _nw85
				_698_v33 = _696_i3
			} else {
				var _704_v39 _dafny.Int
				_ = _704_v39
				_704_v39 = _dafny.IntOfInt64(528)
				_704_v39 = (_dafny.Zero).Minus((((_this).F14()).Minus(_704_v39)).Times(_696_i3))
				(globalState).F5 = ((_this).F14()).Cmp(_704_v39) >= 0
				var _705_v40 _dafny.CodePoint
				_ = _705_v40
				_705_v40 = _dafny.CodePoint('j')
				var _706_v41 _dafny.MultiSet
				_ = _706_v41
				_706_v41 = _dafny.MultiSetOf(_705_v40)
				_704_v39 = ((_706_v41).Union(_dafny.MultiSetFromSeq(p0))).Cardinality()
				var _707_v42 _dafny.Array
				_ = _707_v42
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_8
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) bool = (func(_708_v0 bool) func(_dafny.Int) bool {
						return func(_709_i4 _dafny.Int) bool {
							return _708_v0
						}
					})(_658_v0)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw86 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw86).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw86).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_707_v42 = _nw86
				var _710_v43 D5
				_ = _710_v43
				_710_v43 = Companion_D5_.Create_DC13_(_707_v42)
				_707_v42 = (_710_v43).Dtor_cf19()
				var _711_v44 D0
				_ = _711_v44
				_711_v44 = Companion_D0_.Create_DC2_(_658_v0)
				var _712_v45 _dafny.Sequence
				_ = _712_v45
				_712_v45 = _dafny.SeqOf(p0, p0)
				var _713_v46 D0
				_ = _713_v46
				_713_v46 = Companion_D0_.Create_DC1_(Companion_Default___.Fm19(_696_i3, (_711_v44).Dtor_cf4(), globalState), _712_v45, p0)
				var _714_v47 _dafny.MultiSet
				_ = _714_v47
				_714_v47 = _dafny.MultiSetOf(_713_v46)
				var _715_v48 _dafny.Sequence
				_ = _715_v48
				_715_v48 = _dafny.SeqOf(_714_v47)
				var _716_v49 *C0
				_ = _716_v49
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__((_715_v48).Select((Companion_Default___.SafeIndex(_704_v39, _dafny.IntOfUint32((_715_v48).Cardinality()))).Uint32()).(_dafny.MultiSet), false)
				_716_v49 = _nw87
			}
			var _717_v50 _dafny.Array
			_ = _717_v50
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw88
			_717_v50 = _nw88
			var _718_v51 _dafny.Array
			_ = _718_v51
			var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw89
			_718_v51 = _nw89
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_717_v50), 0))
			_ = _index91
			(_717_v50).ArraySet1(_718_v51, (_index91).Int())
			var _719_v53 _dafny.Map
			_ = _719_v53
			_719_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter23 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)).Keys().Elements()); ; {
					_compr_22, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _720_v52 _dafny.Int
					_720_v52 = interface{}(_compr_22).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)).Contains(_720_v52) {
						_coll22.Add(Companion_Default___.SafeDivisionInt(_720_v52, _dafny.IntOfInt64(275)))
					}
				}
				return _coll22.ToSet()
			}()).Cardinality(), _dafny.IntOfInt64(332))
			var _721_v54 _dafny.MultiSet
			_ = _721_v54
			_721_v54 = _dafny.MultiSetOf(_696_i3, (_this).F14(), (_719_v53).Cardinality(), _696_i3)
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_717_v50), 0))
			_ = _index92
			(_717_v50).ArraySet1((func() _dafny.Array {
				if (_721_v54).IsSubsetOf(_dafny.MultiSetOf(_696_i3)) {
					return _718_v51
				}
				return _718_v51
			})(), (_index92).Int())
			_719_v53 = (_719_v53).Update(((_this).F14()).Minus((_this).F14()), Companion_Default___.SafeModuloInt(_696_i3, (_this).F14()))
			var _722_v55 _dafny.Int
			_ = _722_v55
			_722_v55 = _dafny.IntOfInt64(977)
			var _723_v56 _dafny.CodePoint
			_ = _723_v56
			_723_v56 = _dafny.CodePoint('x')
			_722_v55 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_dafny.IntOfInt64(902)).Plus(_696_i3), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _723_v56)).Cardinality())
		}
		var _724_v57 _dafny.CodePoint
		_ = _724_v57
		_724_v57 = _dafny.CodePoint('a')
		var _725_v58 _dafny.MultiSet
		_ = _725_v58
		_725_v58 = _dafny.MultiSetOf(_724_v57)
		var _726_v59 _dafny.Sequence
		_ = _726_v59
		_726_v59 = _dafny.SeqOf(_658_v0)
		var _727_v60 _dafny.Int
		_ = _727_v60
		_727_v60 = _dafny.IntOfInt64(976)
		var _728_v61 _dafny.Sequence
		_ = _728_v61
		_728_v61 = _dafny.SeqOf((_this).F14(), (_this).F14())
		var _729_v62 _dafny.Map
		_ = _729_v62
		_729_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, _727_v60)
		var _rhs107 bool = !((_dafny.IntOfInt64(-758)).Cmp(Companion_Default___.Fm11(!(_658_v0), globalState)) == 0)
		_ = _rhs107
		var _rhs108 _dafny.MultiSet = ((_725_v58).Intersection(_dafny.MultiSetOf(_724_v57))).Difference(_725_v58)
		_ = _rhs108
		var _rhs109 _dafny.Sequence = Companion_Default___.Fm27((_this).F14(), _727_v60, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_728_v61).Cardinality()), _dafny.IntOfUint32((_728_v61).Cardinality())), globalState)
		_ = _rhs109
		var _rhs110 bool = _658_v0
		_ = _rhs110
		var _rhs111 _dafny.Int = ((_729_v62).Cardinality()).Times((_this).F14())
		_ = _rhs111
		var _lhs74 *GlobalState = globalState
		_ = _lhs74
		var _lhs75 *GlobalState = globalState
		_ = _lhs75
		_lhs74.F13 = _rhs107
		_725_v58 = _rhs108
		_726_v59 = _rhs109
		_lhs75.F13 = _rhs110
		_727_v60 = _rhs111
		var _730_v63 D0
		_ = _730_v63
		_730_v63 = Companion_D0_.Create_DC1_(_658_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}((func(_731_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
			return func(_732_i5 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_733_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_734_i6 _dafny.Int) _dafny.CodePoint {
						return _733_v57
					}
				})(_731_v57)))
			}
		})(_724_v57))), p0)
		var _735_v64 _dafny.MultiSet
		_ = _735_v64
		_735_v64 = _dafny.MultiSetOf(_730_v63, _730_v63, Companion_D0_.Create_DC1_(_658_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}((func(_736_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_737_i7 _dafny.Int) _dafny.Sequence {
				return _736_p0
			}
		})(p0))), p0))
		var _738_v65 D2
		_ = _738_v65
		_738_v65 = Companion_D2_.Create_DC6_(_735_v64)
		var _pat_let_tv19 = _735_v64
		_ = _pat_let_tv19
		var _pat_let_tv20 = _735_v64
		_ = _pat_let_tv20
		var _source11 D2 = func(_pat_let17_0 D2) D2 {
			return func(_739_dt__update__tmp_h1 D2) D2 {
				return func(_pat_let18_0 _dafny.MultiSet) D2 {
					return func(_740_dt__update_hcf10_h0 _dafny.MultiSet) D2 {
						return Companion_D2_.Create_DC6_(_740_dt__update_hcf10_h0)
					}(_pat_let18_0)
				}((_pat_let_tv19).Union(_pat_let_tv20))
			}(_pat_let17_0)
		}(_738_v65)
		_ = _source11
		if _source11.Is_DC7() {
			var _741___mcc_h0 bool = _source11.Get_().(D2_DC7).Cf11
			_ = _741___mcc_h0
			var _742___mcc_h1 _dafny.Map = _source11.Get_().(D2_DC7).Cf12
			_ = _742___mcc_h1
			var _743___mcc_h2 _dafny.Int = _source11.Get_().(D2_DC7).Cf13
			_ = _743___mcc_h2
			var _744___mcc_h3 bool = _source11.Get_().(D2_DC7).Cf14
			_ = _744___mcc_h3
			var _745_cf14 bool = _744___mcc_h3
			_ = _745_cf14
			var _746_cf13 _dafny.Int = _743___mcc_h2
			_ = _746_cf13
			var _747_cf12 _dafny.Map = _742___mcc_h1
			_ = _747_cf12
			var _748_cf11 bool = _741___mcc_h0
			_ = _748_cf11
			var _749_v66 _dafny.Sequence
			_ = _749_v66
			_749_v66 = _dafny.SeqOf(_730_v63)
			var _750_v67 *C0
			_ = _750_v67
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__((_735_v64).Union((_dafny.MultiSetFromSeq(_749_v66)).Update(_730_v63, Companion_Default___.Abs(Companion_Default___.Fm11(_748_cf11, globalState)))), _748_cf11)
			_750_v67 = _nw90
			_727_v60 = (_this).F14()
			var _751_v68 _dafny.Map
			_ = _751_v68
			_751_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_746_cf13, Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_727_v60), _dafny.IntOfInt64(-627)))
			_751_v68 = (_751_v68).Update(_727_v60, _727_v60)
			var _752_v69 D2
			_ = _752_v69
			_752_v69 = Companion_D2_.Create_DC7_(_658_v0, _747_cf12, _727_v60, _658_v0)
			var _753_v70 D0
			_ = _753_v70
			_753_v70 = Companion_D0_.Create_DC2_(!((_750_v67).F19()))
			var _754_v71 D0
			_ = _754_v71
			_754_v71 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(_753_v70))
			var _755_v72 _dafny.Sequence
			_ = _755_v72
			_755_v72 = _dafny.SeqOf((_this).Fm7(p0, _745_cf14, _748_cf11, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_cf14, _748_cf11), globalState))
			var _756_v73 _dafny.Array
			_ = _756_v73
			var _nwElement0_17 _dafny.Int = (Companion_Default___.Fm10(_745_cf14, _746_cf13, _658_v0, (_750_v67).F19(), globalState)).Cardinality()
			_ = _nwElement0_17
			var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(20))
			_ = _nw91
			(_nw91).ArraySet1(_nwElement0_17, 0)
			(_nw91).ArraySet1(_727_v60, 1)
			(_nw91).ArraySet1(_727_v60, 2)
			(_nw91).ArraySet1(_727_v60, 3)
			(_nw91).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _745_cf14 {
					return _746_cf13
				}
				return _727_v60
			})()), 4)
			(_nw91).ArraySet1(_746_cf13, 5)
			(_nw91).ArraySet1((_746_cf13).Times(_dafny.IntOfUint32((_728_v61).Cardinality())), 6)
			(_nw91).ArraySet1(Companion_Default___.Fm22(_752_v69, _754_v71, false, globalState), 7)
			(_nw91).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(405))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_757_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_758_i8 _dafny.Int) _dafny.CodePoint {
					return _757_v57
				}
			})(_724_v57))))).Cardinality()), 8)
			(_nw91).ArraySet1(Companion_Default___.SafeModuloInt((_this).F14(), _746_cf13), 9)
			(_nw91).ArraySet1(_dafny.IntOfUint32((_755_v72).Cardinality()), 10)
			(_nw91).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sxb")).Cardinality())).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_726_v59).Cardinality()), (_750_v67).F19())).Cardinality()), 11)
			(_nw91).ArraySet1(_727_v60, 12)
			(_nw91).ArraySet1((_727_v60).Minus((_this).F14()), 13)
			(_nw91).ArraySet1(_727_v60, 14)
			(_nw91).ArraySet1(_dafny.IntOfInt64(884), 15)
			(_nw91).ArraySet1((_751_v68).Cardinality(), 16)
			(_nw91).ArraySet1(_dafny.IntOfInt64(945), 17)
			(_nw91).ArraySet1(Companion_Default___.SafeDivisionInt(_746_cf13, _727_v60), 18)
			(_nw91).ArraySet1((_this).F14(), 19)
			_756_v73 = _nw91
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_756_v73), 0))
			_ = _index93
			(_756_v73).ArraySet1((_this).F14(), (_index93).Int())
			var _759_v74 D1
			_ = _759_v74
			_759_v74 = Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_cf14, _756_v73))
			var _760_v75 _dafny.Map
			_ = _760_v75
			_760_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_cf14, _756_v73)
			var _761_v76 D6
			_ = _761_v76
			_761_v76 = Companion_D6_.Create_DC15_(_dafny.MultiSetOf(_748_cf11, (_750_v67).F19()))
			var _762_v77 D4
			_ = _762_v77
			_762_v77 = Companion_D4_.Create_DC12_(((_761_v76).Dtor_cf23()).Cardinality())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_756_v73), 0))
			_ = _index94
			var _rhs112 _dafny.Int = Companion_Default___.SafeDivisionInt((_727_v60).Minus(_dafny.IntOfInt64(42)), (_727_v60).Times((_this).F14()))
			_ = _rhs112
			var _rhs113 bool = true
			_ = _rhs113
			var _rhs114 D1 = Companion_D1_.Create_DC4_(_760_v75)
			_ = _rhs114
			var _rhs115 _dafny.Int = (_dafny.Zero).Minus((_762_v77).Dtor_cf18())
			_ = _rhs115
			var _lhs76 _dafny.Array = _756_v73
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_756_v73), 0))
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			(_lhs76).ArraySet1(_rhs112, (_lhs77).Int())
			_lhs78.F13 = _rhs113
			_759_v74 = _rhs114
			_727_v60 = _rhs115
		} else if _source11.Is_DC6() {
			var _763___mcc_h4 _dafny.MultiSet = _source11.Get_().(D2_DC6).Cf10
			_ = _763___mcc_h4
			var _764_cf10 _dafny.MultiSet = _763___mcc_h4
			_ = _764_cf10
			var _765_v78 _dafny.Array
			_ = _765_v78
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw92
			_765_v78 = _nw92
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_765_v78), 0))
			_ = _index95
			(_765_v78).ArraySet1(false, (_index95).Int())
			var _766_v79 D1
			_ = _766_v79
			_766_v79 = Companion_D1_.Create_DC5_(_658_v0, p0, _726_v59)
			var _767_v80 _dafny.Sequence
			_ = _767_v80
			_767_v80 = _dafny.SeqOf(_766_v79)
			var _768_v81 _dafny.MultiSet
			_ = _768_v81
			_768_v81 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Contains(_767_v80, _766_v79), _658_v0)
			var _769_v82 _dafny.Array
			_ = _769_v82
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_9
			var _nw93 _dafny.Array
			_ = _nw93
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw93 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_770_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_771_i9 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_771_i9, _dafny.IntOfUint32((_770_p0).Cardinality()))
					}
				})(p0)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw93 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw93).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw93).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_769_v82 = _nw93
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_769_v82), 0))
			_ = _index96
			(_769_v82).ArraySet1((_this).F14(), (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_765_v78), 0))
			_ = _index97
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_769_v82), 0))
			_ = _index98
			var _rhs116 bool = _658_v0
			_ = _rhs116
			var _rhs117 _dafny.MultiSet = (_768_v81).Difference(_768_v81)
			_ = _rhs117
			var _rhs118 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(661))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}(func(_772_i10 _dafny.Int) _dafny.Int {
				return (_this).F14()
			})), _728_v61)
			_ = _rhs118
			var _rhs119 _dafny.Int = (_this).F14()
			_ = _rhs119
			var _rhs120 bool = !(_658_v0)
			_ = _rhs120
			var _lhs79 _dafny.Array = _765_v78
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_765_v78), 0))
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			var _lhs82 _dafny.Array = _769_v82
			_ = _lhs82
			var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_769_v82), 0))
			_ = _lhs83
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			(_lhs79).ArraySet1(_rhs116, (_lhs80).Int())
			_768_v81 = _rhs117
			_lhs81.F13 = _rhs118
			(_lhs82).ArraySet1(_rhs119, (_lhs83).Int())
			_lhs84.F13 = _rhs120
			(globalState).F5 = (_765_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_765_v78), 0))).Int()).(bool)
			var _773_v83 _dafny.Sequence
			_ = _773_v83
			_773_v83 = _dafny.UnicodeSeqOfUtf8Bytes("umgbbbrln")
			_773_v83 = _dafny.Companion_Sequence_.Concatenate(p0, _773_v83)
			var _774_v84 D6
			_ = _774_v84
			_774_v84 = Companion_D6_.Create_DC16_((_this).F14())
			(globalState).F5 = ((_774_v84).Dtor_cf24()).Cmp(_727_v60) >= 0
		} else {
			var _775___mcc_h5 D2 = _source11.Get_().(D2_DC8).Cf15
			_ = _775___mcc_h5
			var _776_cf15 D2 = _775___mcc_h5
			_ = _776_cf15
			_724_v57 = (func() _dafny.CodePoint {
				if !((_727_v60).Cmp(Companion_Default___.Fm11(_658_v0, globalState)) >= 0) {
					return _724_v57
				}
				return _724_v57
			})()
			(globalState).F13 = !(_658_v0)
			var _777_v85 _dafny.Array
			_ = _777_v85
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw94
			_777_v85 = _nw94
			var _778_v86 _dafny.Array
			_ = _778_v86
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_10
			var _nw95 _dafny.Array
			_ = _nw95
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw95 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) bool = func(_779_i11 _dafny.Int) bool {
					return true
				}
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw95 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw95).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw95).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_778_v86 = _nw95
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_777_v85), 0))
			_ = _index99
			(_777_v85).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_778_v86, _658_v0), (_index99).Int())
			var _780_v87 _dafny.Array
			_ = _780_v87
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_11
			var _nw96 _dafny.Array
			_ = _nw96
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw96 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Set = (func(_781_v0 bool) func(_dafny.Int) _dafny.Set {
					return func(_782_i12 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_781_v0)).Intersection(_dafny.SetOf(_781_v0))
					}
				})(_658_v0)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw96 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw96).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw96).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_780_v87 = _nw96
			var _783_v88 _dafny.Set
			_ = _783_v88
			_783_v88 = _dafny.SetOf(_658_v0, _658_v0, _658_v0, _658_v0, _658_v0)
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_780_v87), 0))
			_ = _index100
			(_780_v87).ArraySet1(_783_v88, (_index100).Int())
			var _784_v89 _dafny.MultiSet
			_ = _784_v89
			_784_v89 = _dafny.MultiSetOf(_658_v0)
			var _785_v90 _dafny.Map
			_ = _785_v90
			_785_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_778_v86, (_726_v59).Select((Companion_Default___.SafeIndex((_this).F14(), _dafny.IntOfUint32((_726_v59).Cardinality()))).Uint32()).(bool))
			var _786_v91 _dafny.MultiSet
			_ = _786_v91
			_786_v91 = _dafny.MultiSetOf(_727_v60)
			var _787_v92 _dafny.MultiSet
			_ = _787_v92
			_787_v92 = _dafny.MultiSetOf((_786_v91).Cardinality())
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_777_v85), 0))
			_ = _index101
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_780_v87), 0))
			_ = _index102
			var _rhs121 _dafny.Sequence = (func() _dafny.Sequence {
				if _658_v0 {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_728_v61, _728_v61), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_728_v61, _728_v61)).Cardinality()))).Uint32(), (_this).F14())
				}
				return _728_v61
			})()
			_ = _rhs121
			var _rhs122 _dafny.Map = (_785_v90).Merge(_785_v90)
			_ = _rhs122
			var _rhs123 _dafny.Set = ((Companion_Default___.Fm33(_658_v0, _658_v0, (_726_v59).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)).Cardinality(), _dafny.IntOfUint32((_726_v59).Cardinality()))).Uint32()).(bool), globalState)).Difference(_783_v88)).Union((_783_v88).Union(_dafny.SetOf(false)))
			_ = _rhs123
			var _rhs124 _dafny.MultiSet = _dafny.MultiSetFromSeq(_726_v59)
			_ = _rhs124
			var _rhs125 _dafny.Array = (func() _dafny.Array {
				if ((_786_v91).Update(_727_v60, Companion_Default___.Abs(_727_v60))).IsSubsetOf(_787_v92) {
					return _778_v86
				}
				return _778_v86
			})()
			_ = _rhs125
			var _lhs85 _dafny.Array = _777_v85
			_ = _lhs85
			var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_777_v85), 0))
			_ = _lhs86
			var _lhs87 _dafny.Array = _780_v87
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_780_v87), 0))
			_ = _lhs88
			_728_v61 = _rhs121
			(_lhs85).ArraySet1(_rhs122, (_lhs86).Int())
			(_lhs87).ArraySet1(_rhs123, (_lhs88).Int())
			_784_v89 = _rhs124
			_778_v86 = _rhs125
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_778_v86), 0))
			_ = _index103
			(_778_v86).ArraySet1(_658_v0, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_778_v86), 0))
			_ = _index104
			(_778_v86).ArraySet1(_658_v0, (_index104).Int())
		}
		var _788_i13 _dafny.Int
		_ = _788_i13
		_788_i13 = _dafny.Zero
		{
			for !(!(_658_v0) || (!(_658_v0) || (_658_v0))) {
				{
					if (_788_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_788_i13 = (_788_i13).Plus(_dafny.One)
					var _789_v93 _dafny.Sequence
					_ = _789_v93
					_789_v93 = _dafny.UnicodeSeqOfUtf8Bytes("hglx")
					_789_v93 = _dafny.Companion_Sequence_.Update(_789_v93, (Companion_Default___.SafeIndex(((_dafny.Zero).Minus((_this).F14())).Plus((_this).F14()), _dafny.IntOfUint32((_789_v93).Cardinality()))).Uint32(), _724_v57)
					_727_v60 = _dafny.IntOfUint32((p0).Cardinality())
					var _790_v94 D1
					_ = _790_v94
					_790_v94 = Companion_D1_.Create_DC5_((_this).Fm8((_this).F14(), _658_v0, _658_v0, globalState), p0, _726_v59)
					var _791_v95 *C0
					_ = _791_v95
					var _nw97 *C0 = New_C0_()
					_ = _nw97
					_nw97.Ctor__(_735_v64, (_790_v94).Dtor_cf7())
					_791_v95 = _nw97
					_726_v59 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_726_v59, _dafny.SeqOf(_658_v0, _658_v0)), _726_v59)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
	}
}
func (_this *C4) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *C4) F15() _dafny.Map {
	{
		return _this._f15
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	dummy byte
}

func New_C5_() *C5 {
	_this := C5{}

	return &_this
}

type CompanionStruct_C5_ struct {
}

var Companion_C5_ = CompanionStruct_C5_{}

func (_this *C5) Equals(other *C5) bool {
	return _this == other
}

func (_this *C5) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C5)
	return ok && _this.Equals(other)
}

func (*C5) String() string {
	return "_module.C5"
}

func Type_C5_() _dafny.TypeDescriptor {
	return type_C5_{}
}

type type_C5_ struct {
}

func (_this type_C5_) Default() interface{} {
	return (*C5)(nil)
}

func (_this type_C5_) String() string {
	return "main.C5"
}
func (_this *C5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__() {
	{
	}
}
func (_this *C5) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (((func() _dafny.Set {
			var _coll23 = _dafny.NewBuilder()
			_ = _coll23
			for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(947), _dafny.IntOfInt64(745))); ; {
				_compr_23, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _792_v0 _dafny.Int
				_792_v0 = interface{}(_compr_23).(_dafny.Int)
				if ((_dafny.IntOfInt64(947)).Cmp(_792_v0) <= 0) && ((_792_v0).Cmp(_dafny.IntOfInt64(745)) < 0) {
					_coll23.Add(Companion_Default___.SafeModuloInt(_792_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
				}
			}
			return _coll23.ToSet()
		}()).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Plus(_dafny.IntOfInt64(274))
	}
}
func (_this *C5) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_793_i0 _dafny.Int) _dafny.CodePoint {
			return (func() _dafny.CodePoint {
				if false {
					return _dafny.CodePoint('e')
				}
				return _dafny.CodePoint('c')
			})()
		}))
	}
}
func (_this *C5) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(825), (_dafny.SetOf(false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(764)), _dafny.SeqOf(_dafny.IntOfInt64(-875), _dafny.IntOfInt64(578))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(856)), _dafny.SeqOf(_dafny.IntOfInt64(-773), (func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-249), _dafny.IntOfInt64(694))); ; {
				_compr_24, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _794_v0 _dafny.Int
				_794_v0 = interface{}(_compr_24).(_dafny.Int)
				if ((_dafny.IntOfInt64(-249)).Cmp(_794_v0) <= 0) && ((_794_v0).Cmp(_dafny.IntOfInt64(694)) < 0) {
					_coll24.Add((_794_v0).Times((func() _dafny.Map {
						var _coll25 = _dafny.NewMapBuilder()
						_ = _coll25
						for _iter26 := _dafny.Iterate((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(636))).Cardinality())).Elements()); ; {
							_compr_25, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _795_v1 _dafny.Int
							_795_v1 = interface{}(_compr_25).(_dafny.Int)
							if (_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(636))).Cardinality())).Contains(_795_v1) {
								_coll25.Add(Companion_Default___.SafeDivisionInt(_795_v1, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfInt64(-483))
							}
						}
						return _coll25.ToMap()
					}()).Cardinality()), _dafny.IntOfInt64(365))
				}
			}
			return _coll24.ToMap()
		}()).Cardinality())))
	}
}
func (_this *C5) Fm5(p0 D0, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(732)
	}
}
func (_this *C5) M0(p0 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _796_v0 bool
		_ = _796_v0
		_796_v0 = true
		var _797_v1 _dafny.Sequence
		_ = _797_v1
		_797_v1 = _dafny.SeqOf(Companion_Default___.Fm6(globalState), _796_v0, _796_v0)
		var _798_v2 _dafny.Int
		_ = _798_v2
		_798_v2 = _dafny.IntOfInt64(-350)
		var _799_v3 _dafny.Sequence
		_ = _799_v3
		_799_v3 = _dafny.SeqOf(_797_v1, _dafny.Companion_Sequence_.Update(_797_v1, (Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_797_v1).Cardinality()))).Uint32(), _796_v0))
		var _800_v4 _dafny.Sequence
		_ = _800_v4
		_800_v4 = _dafny.UnicodeSeqOfUtf8Bytes("dnhvpnw")
		var _801_v5 _dafny.Sequence
		_ = _801_v5
		_801_v5 = _dafny.SeqOf(_800_v4, _800_v4)
		var _802_i0 _dafny.Int
		_ = _802_i0
		_802_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_797_v1, _dafny.Companion_Sequence_.Concatenate((_799_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_801_v5).Cardinality()), _dafny.IntOfUint32((_799_v3).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(_796_v0))) {
				{
					if (_802_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_802_i0 = (_802_i0).Plus(_dafny.One)
					var _803_v6 *C2
					_ = _803_v6
					var _nw98 *C2 = New_C2_()
					_ = _nw98
					_nw98.Ctor__()
					_803_v6 = _nw98
					_798_v2 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_798_v2), _798_v2)
					if _796_v0 {
						_798_v2 = _798_v2
						r1 = _798_v2
						var _804_v8 _dafny.MultiSet
						_ = _804_v8
						_804_v8 = _dafny.MultiSetOf(_798_v2)
						var _805_v9 _dafny.Map
						_ = _805_v9
						_805_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_798_v2, _798_v2)
						var _806_v10 _dafny.Sequence
						_ = _806_v10
						_806_v10 = _dafny.SeqOf(_798_v2, (_798_v2).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg70 _dafny.Int) interface{} {
								return coer70(arg70)
							}
						}((func(_807_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_808_i1 _dafny.Int) _dafny.Map {
								return func() _dafny.Map {
									var _coll26 = _dafny.NewMapBuilder()
									_ = _coll26
									for _iter27 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-454), _807_v2, _dafny.IntOfInt64(576), _807_v2, _807_v2)).Elements()); ; {
										_compr_26, _ok27 := _iter27()
										if !_ok27 {
											break
										}
										var _809_v7 _dafny.Int
										_809_v7 = interface{}(_compr_26).(_dafny.Int)
										if (_dafny.MultiSetOf(_dafny.IntOfInt64(-454), _807_v2, _dafny.IntOfInt64(576), _807_v2, _807_v2)).Contains(_809_v7) {
											_coll26.Add(Companion_Default___.SafeDivisionInt(_809_v7, _807_v2), _807_v2)
										}
									}
									return _coll26.ToMap()
								}()
							}
						})(_798_v2))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_804_v8).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg71 _dafny.Int) interface{} {
								return coer71(arg71)
							}
						}((func(_810_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_811_i1 _dafny.Int) _dafny.Map {
								return func() _dafny.Map {
									var _coll27 = _dafny.NewMapBuilder()
									_ = _coll27
									for _iter28 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-454), _810_v2, _dafny.IntOfInt64(576), _810_v2, _810_v2)).Elements()); ; {
										_compr_27, _ok28 := _iter28()
										if !_ok28 {
											break
										}
										var _812_v7 _dafny.Int
										_812_v7 = interface{}(_compr_27).(_dafny.Int)
										if (_dafny.MultiSetOf(_dafny.IntOfInt64(-454), _810_v2, _dafny.IntOfInt64(576), _810_v2, _810_v2)).Contains(_812_v7) {
											_coll27.Add(Companion_Default___.SafeDivisionInt(_812_v7, _810_v2), _810_v2)
										}
									}
									return _coll27.ToMap()
								}()
							}
						})(_798_v2)))).Cardinality()))).Uint32(), _805_v9)).Cardinality())), _dafny.IntOfUint32((_800_v4).Cardinality()))
						var _rhs126 bool = _796_v0
						_ = _rhs126
						var _rhs127 _dafny.Int = (_798_v2).Times(_798_v2)
						_ = _rhs127
						var _rhs128 _dafny.Sequence = _806_v10
						_ = _rhs128
						r0 = _rhs126
						_798_v2 = _rhs127
						_806_v10 = _rhs128
						var _813_v11 _dafny.CodePoint
						_ = _813_v11
						_813_v11 = _dafny.CodePoint('b')
						_813_v11 = _813_v11
						var _814_v12 _dafny.Array
						_ = _814_v12
						var _len0_12 _dafny.Int = _dafny.IntOfInt64(20)
						_ = _len0_12
						var _nw99 _dafny.Array
						_ = _nw99
						if _len0_12.Cmp(_dafny.Zero) == 0 {
							_nw99 = _dafny.NewArray(_len0_12)
						} else {
							var _init12 func(_dafny.Int) _dafny.Int = (func(_815_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_816_i2 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_816_i2, _815_v2)
								}
							})(_798_v2)
							_ = _init12
							var _element0_12 = _init12(_dafny.Zero)
							_ = _element0_12
							_nw99 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
							(_nw99).ArraySet1(_element0_12, 0)
							var _nativeLen0_12 = (_len0_12).Int()
							_ = _nativeLen0_12
							for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
								(_nw99).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
							}
						}
						_814_v12 = _nw99
						var _817_v13 _dafny.Map
						_ = _817_v13
						_817_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_814_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(805))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg72 _dafny.Int) interface{} {
								return coer72(arg72)
							}
						}((func(_818_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_819_i3 _dafny.Int) _dafny.CodePoint {
								return _818_p0
							}
						})(p0))))
						var _820_v14 D0
						_ = _820_v14
						_820_v14 = Companion_D0_.Create_DC1_(false, _801_v5, (func() _dafny.Sequence {
							if (_817_v13).Contains(_814_v12) {
								return (_817_v13).Get(_814_v12).(_dafny.Sequence)
							}
							return _800_v4
						})())
						var _821_v15 _dafny.MultiSet
						_ = _821_v15
						_821_v15 = _dafny.MultiSetOf(_820_v14, _820_v14)
						var _822_v16 D2
						_ = _822_v16
						_822_v16 = Companion_D2_.Create_DC6_(_821_v15)
						var _823_v17 *C0
						_ = _823_v17
						var _nw100 *C0 = New_C0_()
						_ = _nw100
						_nw100.Ctor__((_822_v16).Dtor_cf10(), _796_v0)
						_823_v17 = _nw100
					} else {
						_796_v0 = (_796_v0) || (true)
						(globalState).F13 = _796_v0
						(globalState).F13 = _796_v0
						_797_v1 = _797_v1
						(globalState).F9 = !(_796_v0)
					}
					var _824_v18 _dafny.Map
					_ = _824_v18
					_824_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-708), _796_v0)
					var _825_v20 _dafny.Set
					_ = _825_v20
					_825_v20 = _dafny.SetOf(_796_v0)
					var _826_v21 _dafny.Map
					_ = _826_v21
					_826_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_825_v20).Cardinality(), _798_v2)
					var _827_v22 D7
					_ = _827_v22
					_827_v22 = Companion_D7_.Create_DC18_(func() _dafny.Map {
						var _coll28 = _dafny.NewMapBuilder()
						_ = _coll28
						for _iter29 := _dafny.Iterate((_826_v21).Keys().Elements()); ; {
							_compr_28, _ok29 := _iter29()
							if !_ok29 {
								break
							}
							var _828_v19 _dafny.Int
							_828_v19 = interface{}(_compr_28).(_dafny.Int)
							if (_826_v21).Contains(_828_v19) {
								_coll28.Add((_828_v19).Times(_dafny.IntOfInt64(523)), true)
							}
						}
						return _coll28.ToMap()
					}())
					_824_v18 = (_827_v22).Dtor_cf26()
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if _796_v0 {
			var _829_v23 _dafny.Array
			_ = _829_v23
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_13
			var _nw101 _dafny.Array
			_ = _nw101
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw101 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.CodePoint = (func(_830_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_831_i4 _dafny.Int) _dafny.CodePoint {
						return _830_p0
					}
				})(p0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw101 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw101).ArraySet1CodePoint(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw101).ArraySet1CodePoint(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_829_v23 = _nw101
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_829_v23), 0))
			_ = _index105
			(_829_v23).ArraySet1CodePoint(p0, (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_829_v23), 0))
			_ = _index106
			(_829_v23).ArraySet1CodePoint(p0, (_index106).Int())
			var _832_v24 D0
			_ = _832_v24
			_832_v24 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_(_796_v0, _801_v5, _800_v4))
			var _833_v25 _dafny.Map
			_ = _833_v25
			_833_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v0, _832_v24)
			var _834_v26 D2
			_ = _834_v26
			_834_v26 = Companion_D2_.Create_DC7_(_796_v0, _833_v25, _798_v2, _796_v0)
			var _835_v27 D0
			_ = _835_v27
			_835_v27 = Companion_D0_.Create_DC1_(true, _801_v5, _800_v4)
			_798_v2 = ((_798_v2).Minus(Companion_Default___.Fm22(_834_v26, Companion_D0_.Create_DC3_(_835_v27), _796_v0, globalState))).Plus(Companion_Default___.Fm22(Companion_D2_.Create_DC7_(_796_v0, _833_v25, _dafny.IntOfInt64(168), _796_v0), _832_v24, _796_v0, globalState))
			var _836_v28 _dafny.Array
			_ = _836_v28
			var _nwElement0_18 bool = _796_v0
			_ = _nwElement0_18
			var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(27))
			_ = _nw102
			(_nw102).ArraySet1(_nwElement0_18, 0)
			(_nw102).ArraySet1(_796_v0, 1)
			(_nw102).ArraySet1(_796_v0, 2)
			(_nw102).ArraySet1(_796_v0, 3)
			(_nw102).ArraySet1(!(_796_v0), 4)
			(_nw102).ArraySet1(!(_796_v0), 5)
			(_nw102).ArraySet1(_796_v0, 6)
			(_nw102).ArraySet1(_796_v0, 7)
			(_nw102).ArraySet1(_796_v0, 8)
			(_nw102).ArraySet1(_796_v0, 9)
			(_nw102).ArraySet1(_796_v0, 10)
			(_nw102).ArraySet1(_796_v0, 11)
			(_nw102).ArraySet1(_796_v0, 12)
			(_nw102).ArraySet1(!(_796_v0), 13)
			(_nw102).ArraySet1(!(_796_v0), 14)
			(_nw102).ArraySet1(_796_v0, 15)
			(_nw102).ArraySet1(_796_v0, 16)
			(_nw102).ArraySet1(_796_v0, 17)
			(_nw102).ArraySet1(_796_v0, 18)
			(_nw102).ArraySet1(_796_v0, 19)
			(_nw102).ArraySet1(_796_v0, 20)
			(_nw102).ArraySet1(false, 21)
			(_nw102).ArraySet1(true, 22)
			(_nw102).ArraySet1(_796_v0, 23)
			(_nw102).ArraySet1(_796_v0, 24)
			(_nw102).ArraySet1(_796_v0, 25)
			(_nw102).ArraySet1(_796_v0, 26)
			_836_v28 = _nw102
			var _837_v29 _dafny.Array
			_ = _837_v29
			var _nwElement0_19 _dafny.Array = _836_v28
			_ = _nwElement0_19
			var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(16))
			_ = _nw103
			(_nw103).ArraySet1(_nwElement0_19, 0)
			(_nw103).ArraySet1(_836_v28, 1)
			(_nw103).ArraySet1(_836_v28, 2)
			(_nw103).ArraySet1(_836_v28, 3)
			(_nw103).ArraySet1(_836_v28, 4)
			(_nw103).ArraySet1(_836_v28, 5)
			(_nw103).ArraySet1(_836_v28, 6)
			(_nw103).ArraySet1(_836_v28, 7)
			(_nw103).ArraySet1(_836_v28, 8)
			(_nw103).ArraySet1(_836_v28, 9)
			(_nw103).ArraySet1(_836_v28, 10)
			(_nw103).ArraySet1(_836_v28, 11)
			(_nw103).ArraySet1(_836_v28, 12)
			(_nw103).ArraySet1(_836_v28, 13)
			(_nw103).ArraySet1(_836_v28, 14)
			(_nw103).ArraySet1(_836_v28, 15)
			_837_v29 = _nw103
			_837_v29 = _837_v29
			var _838_v30 T1
			_ = _838_v30
			var _nw104 *C3 = New_C3_()
			_ = _nw104
			_nw104.Ctor__(_796_v0, _798_v2)
			_838_v30 = _nw104
			var _839_v31 _dafny.Set
			_ = _839_v31
			_839_v31 = _dafny.SetOf(_796_v0)
			var _840_v32 _dafny.MultiSet
			_ = _840_v32
			_840_v32 = _dafny.MultiSetOf(_798_v2)
			(_this).M3((_798_v2).Plus(_798_v2), _838_v30, ((_839_v31).Union(_dafny.SetOf(_796_v0, _796_v0, _796_v0))).Cardinality(), (_840_v32).Union(_840_v32), globalState)
			var _841_v33 _dafny.Set
			_ = _841_v33
			_841_v33 = _839_v31
			var _842_v34 _dafny.Set
			_ = _842_v34
			_842_v34 = _dafny.SetOf(_841_v33)
			var _843_v35 _dafny.Map
			_ = _843_v35
			_843_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v34, _838_v30)
			(globalState).F3 = !(!(_843_v35).Equals(_843_v35))
		} else {
			var _844_v36 _dafny.MultiSet
			_ = _844_v36
			_844_v36 = _dafny.MultiSetOf(_796_v0)
			var _845_v37 _dafny.Sequence
			_ = _845_v37
			_845_v37 = _dafny.SeqOf((func() _dafny.Int {
				if (_844_v36).Contains(_796_v0) {
					return (_844_v36).Multiplicity(_796_v0)
				}
				return _798_v2
			})(), _798_v2, _798_v2, _798_v2)
			_798_v2 = Companion_Default___.SafeModuloInt(_798_v2, ((func() _dafny.Int {
				if (_844_v36).Contains(_796_v0) {
					return (_844_v36).Multiplicity(_796_v0)
				}
				return _dafny.IntOfUint32((_845_v37).Cardinality())
			})()).Plus(Companion_Default___.Fm11(Companion_Default___.Fm6(globalState), globalState)))
			r0 = _796_v0
			var _846_v38 _dafny.CodePoint
			_ = _846_v38
			_846_v38 = _dafny.CodePoint('e')
			var _847_v39 _dafny.Array
			_ = _847_v39
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_14
			var _nw105 _dafny.Array
			_ = _nw105
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw105 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_848_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_849_i5 _dafny.Int) _dafny.Int {
						return (_849_i5).Minus(_848_v2)
					}
				})(_798_v2)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw105 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw105).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw105).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_847_v39 = _nw105
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_847_v39), 0))
			_ = _index107
			(_847_v39).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_800_v4, _800_v4)).Cardinality()), (_index107).Int())
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_847_v39), 0))
			_ = _index108
			(_847_v39).ArraySet1((_845_v37).Select((Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_845_v37).Cardinality()))).Uint32()).(_dafny.Int), (_index108).Int())
			var _850_v40 D0
			_ = _850_v40
			_850_v40 = Companion_D0_.Create_DC2_(_796_v0)
			var _851_v41 _dafny.Sequence
			_ = _851_v41
			_851_v41 = _dafny.SeqOf(_850_v40)
			var _852_v42 _dafny.Map
			_ = _852_v42
			_852_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_798_v2, !(false))
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_847_v39), 0))
			_ = _index109
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_847_v39), 0))
			_ = _index110
			var _rhs129 _dafny.CodePoint = _846_v38
			_ = _rhs129
			var _rhs130 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ebekj"), (_801_v5).Select((Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_801_v5).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality())
			_ = _rhs130
			var _rhs131 _dafny.Int = _dafny.IntOfUint32((_797_v1).Cardinality())
			_ = _rhs131
			var _rhs132 _dafny.Int = (_dafny.IntOfInt64(-52)).Minus((_852_v42).Cardinality())
			_ = _rhs132
			var _rhs133 _dafny.Sequence = _851_v41
			_ = _rhs133
			var _lhs89 _dafny.Array = _847_v39
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_847_v39), 0))
			_ = _lhs90
			var _lhs91 _dafny.Array = _847_v39
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_847_v39), 0))
			_ = _lhs92
			_846_v38 = _rhs129
			(_lhs89).ArraySet1(_rhs130, (_lhs90).Int())
			(_lhs91).ArraySet1(_rhs131, (_lhs92).Int())
			r1 = _rhs132
			_851_v41 = _rhs133
			r1 = _798_v2
			var _853_v43 *C2
			_ = _853_v43
			var _nw106 *C2 = New_C2_()
			_ = _nw106
			_nw106.Ctor__()
			_853_v43 = _nw106
		}
		if (_796_v0) && (_796_v0) {
			var _854_v44 _dafny.Array
			_ = _854_v44
			var _nwElement0_20 _dafny.CodePoint = p0
			_ = _nwElement0_20
			var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(19))
			_ = _nw107
			(_nw107).ArraySet1CodePoint(_nwElement0_20, 0)
			(_nw107).ArraySet1CodePoint(p0, 1)
			(_nw107).ArraySet1CodePoint(p0, 2)
			(_nw107).ArraySet1CodePoint(p0, 3)
			(_nw107).ArraySet1CodePoint(_dafny.CodePoint('c'), 4)
			(_nw107).ArraySet1CodePoint(p0, 5)
			(_nw107).ArraySet1CodePoint(_dafny.CodePoint('b'), 6)
			(_nw107).ArraySet1CodePoint(p0, 7)
			(_nw107).ArraySet1CodePoint(p0, 8)
			(_nw107).ArraySet1CodePoint(p0, 9)
			(_nw107).ArraySet1CodePoint(p0, 10)
			(_nw107).ArraySet1CodePoint(p0, 11)
			(_nw107).ArraySet1CodePoint(_dafny.CodePoint('e'), 12)
			(_nw107).ArraySet1CodePoint(p0, 13)
			(_nw107).ArraySet1CodePoint(p0, 14)
			(_nw107).ArraySet1CodePoint(p0, 15)
			(_nw107).ArraySet1CodePoint(_dafny.CodePoint('q'), 16)
			(_nw107).ArraySet1CodePoint(p0, 17)
			(_nw107).ArraySet1CodePoint(p0, 18)
			_854_v44 = _nw107
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_854_v44), 0))
			_ = _index111
			(_854_v44).ArraySet1CodePoint(p0, (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_854_v44), 0))
			_ = _index112
			(_854_v44).ArraySet1CodePoint(p0, (_index112).Int())
			var _855_v45 _dafny.Array
			_ = _855_v45
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_15
			var _nw108 _dafny.Array
			_ = _nw108
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw108 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = (func(_856_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_857_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_857_i6, _856_v2)
					}
				})(_798_v2)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw108 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw108).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw108).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_855_v45 = _nw108
			var _858_v46 _dafny.Set
			_ = _858_v46
			_858_v46 = _dafny.SetOf(_796_v0, Companion_Default___.Fm6(globalState), _796_v0, !(_796_v0), _796_v0)
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_855_v45), 0))
			_ = _index113
			(_855_v45).ArraySet1((_858_v46).Cardinality(), (_index113).Int())
			var _859_v47 _dafny.Set
			_ = _859_v47
			_859_v47 = _dafny.SetOf(_798_v2)
			var _860_v48 _dafny.Map
			_ = _860_v48
			_860_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_798_v2, (_859_v47).Cardinality())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_855_v45), 0))
			_ = _index114
			(_855_v45).ArraySet1(((_860_v48).Cardinality()).Times(_798_v2), (_index114).Int())
			var _861_v49 _dafny.Array
			_ = _861_v49
			var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
			_ = _nw109
			_861_v49 = _nw109
			var _862_v50 _dafny.Set
			_ = _862_v50
			_862_v50 = _dafny.SetOf(_855_v45)
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_861_v49), 0))
			_ = _index115
			(_861_v49).ArraySet1(_862_v50, (_index115).Int())
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_861_v49), 0))
			_ = _index116
			(_861_v49).ArraySet1(_862_v50, (_index116).Int())
			_798_v2 = (func() _dafny.Int {
				if (_860_v48).Contains((_dafny.Zero).Minus((_855_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_855_v45), 0))).Int()).(_dafny.Int))) {
					return (_860_v48).Get((_dafny.Zero).Minus((_855_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_855_v45), 0))).Int()).(_dafny.Int))).(_dafny.Int)
				}
				return (_798_v2).Minus(_798_v2)
			})()
			r1 = (_855_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_855_v45), 0))).Int()).(_dafny.Int)
		} else {
			(globalState).F9 = (func() bool {
				if _796_v0 {
					return _796_v0
				}
				return _796_v0
			})()
			var _863_v51 _dafny.Set
			_ = _863_v51
			_863_v51 = _dafny.SetOf(!(false), _796_v0, _796_v0)
			var _864_v52 _dafny.Map
			_ = _864_v52
			_864_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_863_v51).Cardinality(), _798_v2)
			_864_v52 = (_864_v52).Update((_dafny.IntOfInt64(151)).Times(_dafny.IntOfInt64(147)), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_798_v2, _798_v2)))
			var _865_v53 _dafny.Sequence
			_ = _865_v53
			_865_v53 = _dafny.SeqOf(_dafny.IntOfUint32((_800_v4).Cardinality()))
			var _866_v54 _dafny.MultiSet
			_ = _866_v54
			_866_v54 = _dafny.MultiSetOf(_865_v53, _dafny.SeqOf(_798_v2, _798_v2, _798_v2, _798_v2))
			var _867_v55 D1
			_ = _867_v55
			_867_v55 = Companion_D1_.Create_DC5_(_796_v0, _800_v4, _797_v1)
			var _868_v56 _dafny.Array
			_ = _868_v56
			var _nwElement0_21 _dafny.Sequence = _800_v4
			_ = _nwElement0_21
			var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(24))
			_ = _nw110
			(_nw110).ArraySet1(_nwElement0_21, 0)
			(_nw110).ArraySet1(_800_v4, 1)
			(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-339))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg73 _dafny.Int) interface{} {
					return coer73(arg73)
				}
			}((func(_869_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_870_i7 _dafny.Int) _dafny.CodePoint {
					return _869_p0
				}
			})(p0))), 2)
			(_nw110).ArraySet1(_800_v4, 3)
			(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(932))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}((func(_871_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_872_i8 _dafny.Int) _dafny.CodePoint {
					return _871_p0
				}
			})(p0))), 4)
			(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-895))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg75 _dafny.Int) interface{} {
					return coer75(arg75)
				}
			}((func(_873_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_874_i9 _dafny.Int) _dafny.CodePoint {
					return _873_p0
				}
			})(p0))), 5)
			(_nw110).ArraySet1(_800_v4, 6)
			(_nw110).ArraySet1(_800_v4, 7)
			(_nw110).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rl"), 8)
			(_nw110).ArraySet1(_800_v4, 9)
			(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-2))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg76 _dafny.Int) interface{} {
					return coer76(arg76)
				}
			}((func(_875_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_876_i10 _dafny.Int) _dafny.CodePoint {
					return _875_p0
				}
			})(p0))), 10)
			(_nw110).ArraySet1(_800_v4, 11)
			(_nw110).ArraySet1(Companion_Default___.Fm20(_dafny.IntOfUint32((_865_v53).Cardinality()), _796_v0, _796_v0, _798_v2, globalState), 12)
			(_nw110).ArraySet1(_800_v4, 13)
			(_nw110).ArraySet1(_800_v4, 14)
			(_nw110).ArraySet1(_800_v4, 15)
			(_nw110).ArraySet1(_800_v4, 16)
			(_nw110).ArraySet1(_800_v4, 17)
			(_nw110).ArraySet1(_800_v4, 18)
			(_nw110).ArraySet1(_800_v4, 19)
			(_nw110).ArraySet1(_800_v4, 20)
			(_nw110).ArraySet1(_800_v4, 21)
			(_nw110).ArraySet1(_800_v4, 22)
			(_nw110).ArraySet1(_800_v4, 23)
			_868_v56 = _nw110
			var _877_v57 D7
			_ = _877_v57
			_877_v57 = Companion_D7_.Create_DC19_(_867_v55, _798_v2, _796_v0, _868_v56)
			var _878_v58 D0
			_ = _878_v58
			_878_v58 = Companion_D0_.Create_DC2_(false)
			var _879_v59 _dafny.Sequence
			_ = _879_v59
			_879_v59 = _dafny.SeqOf(_865_v53, _865_v53, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_865_v53, (Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_865_v53).Cardinality()))).Uint32(), _dafny.IntOfInt64(927)), (Companion_Default___.SafeIndex((_877_v57).Dtor_cf28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_865_v53, (Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_865_v53).Cardinality()))).Uint32(), _dafny.IntOfInt64(927))).Cardinality()))).Uint32(), (_this).Fm5(_878_v58, globalState)), _dafny.SeqOf(_798_v2), _865_v53)
			if ((_866_v54).Intersection(Companion_Default___.Fm34(globalState))).IsDisjointFrom(_dafny.MultiSetFromSeq(_879_v59)) {
				var _880_v60 _dafny.Map
				_ = _880_v60
				_880_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v0, _798_v2)
				_880_v60 = (_880_v60).Update(true, _798_v2)
				var _881_v61 D0
				_ = _881_v61
				_881_v61 = Companion_D0_.Create_DC1_(_796_v0, _801_v5, _800_v4)
				var _882_v62 _dafny.MultiSet
				_ = _882_v62
				_882_v62 = _dafny.MultiSetOf(_881_v61)
				var _883_v63 *C0
				_ = _883_v63
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__(_882_v62, _796_v0)
				_883_v63 = _nw111
				var _884_v64 _dafny.Array
				_ = _884_v64
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw112
				_884_v64 = _nw112
				var _885_v65 bool
				_ = _885_v65
				var _886_v66 _dafny.Int
				_ = _886_v66
				var _887_v67 _dafny.Int
				_ = _887_v67
				var _out11 bool
				_ = _out11
				var _out12 _dafny.Int
				_ = _out12
				var _out13 _dafny.Int
				_ = _out13
				_out11, _out12, _out13 = (_this).M4(_865_v53, _884_v64, globalState)
				_885_v65 = _out11
				_886_v66 = _out12
				_887_v67 = _out13
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_884_v64), 0))
				_ = _index117
				(_884_v64).ArraySet1((func() bool {
					if _885_v65 {
						return _796_v0
					}
					return _796_v0
				})(), (_index117).Int())
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_884_v64), 0))
				_ = _index118
				(_884_v64).ArraySet1(_885_v65, (_index118).Int())
				var _888_v68 _dafny.Array
				_ = _888_v68
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw113
				_888_v68 = _nw113
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_888_v68), 0))
				_ = _index119
				(_888_v68).ArraySet1(_884_v64, (_index119).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_888_v68), 0))
				_ = _index120
				(_888_v68).ArraySet1(_884_v64, (_index120).Int())
			} else {
				var _889_v69 _dafny.MultiSet
				_ = _889_v69
				_889_v69 = _dafny.MultiSetOf(_798_v2)
				(globalState).F5 = !((_dafny.MultiSetFromSeq(_865_v53)).IsSubsetOf(_889_v69))
				(globalState).F13 = ((_dafny.Zero).Minus(_798_v2)).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-799))) <= 0
				var _890_v70 _dafny.Array
				_ = _890_v70
				var _nwElement0_22 _dafny.Int = (func() _dafny.Int {
					if _796_v0 {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_800_v4, (Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_800_v4).Cardinality()))).Uint32(), p0)).Cardinality())
					}
					return _798_v2
				})()
				_ = _nwElement0_22
				var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(12))
				_ = _nw114
				(_nw114).ArraySet1(_nwElement0_22, 0)
				(_nw114).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_797_v1, _dafny.SeqOf(_796_v0))).Cardinality()), 1)
				(_nw114).ArraySet1((_798_v2).Plus(_798_v2), 2)
				(_nw114).ArraySet1(_798_v2, 3)
				(_nw114).ArraySet1(_798_v2, 4)
				(_nw114).ArraySet1((_dafny.Zero).Minus(_798_v2), 5)
				(_nw114).ArraySet1((_dafny.Zero).Minus(_798_v2), 6)
				(_nw114).ArraySet1(Companion_Default___.SafeModuloInt(_798_v2, _798_v2), 7)
				(_nw114).ArraySet1(_798_v2, 8)
				(_nw114).ArraySet1(_798_v2, 9)
				(_nw114).ArraySet1((_798_v2).Plus(_798_v2), 10)
				(_nw114).ArraySet1(_dafny.IntOfUint32((_800_v4).Cardinality()), 11)
				_890_v70 = _nw114
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_890_v70), 0))
				_ = _index121
				(_890_v70).ArraySet1((_798_v2).Plus(_798_v2), (_index121).Int())
				var _891_v71 _dafny.Array
				_ = _891_v71
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw115
				_891_v71 = _nw115
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_891_v71), 0))
				_ = _index122
				(_891_v71).ArraySet1(!(!(!(_796_v0))), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_890_v70), 0))
				_ = _index123
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_891_v71), 0))
				_ = _index124
				var _rhs134 _dafny.Int = _798_v2
				_ = _rhs134
				var _rhs135 bool = _796_v0
				_ = _rhs135
				var _rhs136 bool = _796_v0
				_ = _rhs136
				var _lhs93 _dafny.Array = _890_v70
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_890_v70), 0))
				_ = _lhs94
				var _lhs95 _dafny.Array = _891_v71
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_891_v71), 0))
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				(_lhs93).ArraySet1(_rhs134, (_lhs94).Int())
				(_lhs95).ArraySet1(_rhs135, (_lhs96).Int())
				_lhs97.F9 = _rhs136
				var _892_v72 _dafny.Array
				_ = _892_v72
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_16
				var _nw116 _dafny.Array
				_ = _nw116
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw116 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Sequence = (func(_893_v53 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_894_i11 _dafny.Int) _dafny.Sequence {
							return _893_v53
						}
					})(_865_v53)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw116 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw116).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw116).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_892_v72 = _nw116
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_892_v72), 0))
				_ = _index125
				(_892_v72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_865_v53, _865_v53), (_index125).Int())
				var _895_v73 _dafny.CodePoint
				_ = _895_v73
				_895_v73 = _dafny.CodePoint('o')
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_892_v72), 0))
				_ = _index126
				var _rhs137 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_865_v53, (func() _dafny.Sequence {
					if true {
						return _865_v53
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(735))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg77 _dafny.Int) interface{} {
							return coer77(arg77)
						}
					}((func(_896_v71 _dafny.Array, _897_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_898_i12 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_896_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_896_v71), 0))).Int()).(bool), _897_v2)).Cardinality()
						}
					})(_891_v71, _798_v2)))
				})()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_798_v2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_865_v53, (func() _dafny.Sequence {
					if true {
						return _865_v53
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(735))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg78 _dafny.Int) interface{} {
							return coer78(arg78)
						}
					}((func(_899_v71 _dafny.Array, _900_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_901_i12 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_899_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_899_v71), 0))).Int()).(bool), _900_v2)).Cardinality()
						}
					})(_891_v71, _798_v2)))
				})())).Cardinality()))).Uint32(), (_890_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_890_v70), 0))).Int()).(_dafny.Int))
				_ = _rhs137
				var _rhs138 _dafny.Int = (_890_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_890_v70), 0))).Int()).(_dafny.Int)
				_ = _rhs138
				var _rhs139 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_800_v4, (Companion_Default___.SafeIndex(_798_v2, _dafny.IntOfUint32((_800_v4).Cardinality()))).Uint32(), p0), _800_v4), _dafny.UnicodeSeqOfUtf8Bytes("wje"))
				_ = _rhs139
				var _rhs140 _dafny.CodePoint = (_800_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_800_v4).Cardinality()), _dafny.IntOfUint32((_800_v4).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs140
				var _lhs98 _dafny.Array = _892_v72
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_892_v72), 0))
				_ = _lhs99
				(_lhs98).ArraySet1(_rhs137, (_lhs99).Int())
				_798_v2 = _rhs138
				_800_v4 = _rhs139
				_895_v73 = _rhs140
				_895_v73 = p0
			}
			var _902_v74 D0
			_ = _902_v74
			_902_v74 = Companion_D0_.Create_DC1_(_796_v0, _801_v5, _800_v4)
			var _903_v75 _dafny.MultiSet
			_ = _903_v75
			_903_v75 = _dafny.MultiSetOf(_902_v74, _902_v74, _902_v74, _902_v74, _902_v74)
			var _904_v76 *C0
			_ = _904_v76
			var _nw117 *C0 = New_C0_()
			_ = _nw117
			_nw117.Ctor__(_903_v75, !(_796_v0))
			_904_v76 = _nw117
			var _905_v77 _dafny.CodePoint
			_ = _905_v77
			_905_v77 = _dafny.CodePoint('s')
			_905_v77 = _905_v77
		}
		var _906_v78 _dafny.Set
		_ = _906_v78
		_906_v78 = _dafny.SetOf(_796_v0, _796_v0, _796_v0)
		_906_v78 = (_906_v78).Intersection(_906_v78)
		var _907_v79 D0
		_ = _907_v79
		_907_v79 = Companion_D0_.Create_DC2_(_796_v0)
		var _908_v80 *C4
		_ = _908_v80
		var _nw118 *C4 = New_C4_()
		_ = _nw118
		_nw118.Ctor__(_798_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), (_this).Fm5(_907_v79, globalState))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _798_v2)))
		_908_v80 = _nw118
		var _909_i13 _dafny.Int
		_ = _909_i13
		_909_i13 = _dafny.Zero
		{
			for _796_v0 {
				{
					if (_909_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_909_i13 = (_909_i13).Plus(_dafny.One)
					var _910_v81 _dafny.MultiSet
					_ = _910_v81
					_910_v81 = _dafny.MultiSetOf((_908_v80).F14(), ((_908_v80).F14()).Times(_dafny.IntOfUint32((_797_v1).Cardinality())))
					_910_v81 = _910_v81
					_798_v2 = (_dafny.Zero).Minus(Companion_Default___.Fm11(_796_v0, globalState))
					_798_v2 = (_908_v80).F14()
					(globalState).F5 = _796_v0
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = (_796_v0) || (false)
		r1 = (_dafny.Zero).Minus((_798_v2).Times((((_908_v80).F15()).Merge((_908_v80).F15())).Cardinality()))
		return r0, r1
	}
}
func (_this *C5) M3(p0 _dafny.Int, p1 T1, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) {
	{
		var _911_v0 _dafny.Array
		_ = _911_v0
		var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
		_ = _nw119
		_911_v0 = _nw119
		for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_911_v0), 0))); ; {
			_guard_loop_1, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _912_i0 _dafny.Int
			_912_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_912_i0).Sign() != -1) && ((_912_i0).Cmp(_dafny.ArrayLen((_911_v0), 0)) < 0)) {
				(_911_v0).ArraySet1CodePoint((func() _dafny.CodePoint {
					if true {
						return _dafny.CodePoint('y')
					}
					return _dafny.CodePoint('a')
				})(), (_912_i0).Int())
			}
		}
		var _913_v1 _dafny.Array
		_ = _913_v1
		var _nw120 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw120
		_913_v1 = _nw120
		var _914_v2 bool
		_ = _914_v2
		_914_v2 = true
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
		_ = _index127
		(_913_v1).ArraySet1(!(!(_914_v2)), (_index127).Int())
		var _915_v3 _dafny.Int
		_ = _915_v3
		_915_v3 = _dafny.IntOfInt64(-355)
		var _916_v4 _dafny.Sequence
		_ = _916_v4
		_916_v4 = _dafny.UnicodeSeqOfUtf8Bytes("hlo")
		var _917_v5 _dafny.Map
		_ = _917_v5
		_917_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(504), true)
		var _918_v6 D7
		_ = _918_v6
		_918_v6 = Companion_D7_.Create_DC18_(_917_v5)
		var _919_v7 D7
		_ = _919_v7
		_919_v7 = Companion_D7_.Create_DC21_(_918_v6)
		var _920_v8 _dafny.Map
		_ = _920_v8
		_920_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v3, _919_v7)
		var _921_v9 _dafny.Set
		_ = _921_v9
		_921_v9 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, _920_v8)).Cardinality())
		var _922_v10 _dafny.Set
		_ = _922_v10
		_922_v10 = _dafny.SetOf(_914_v2)
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
		_ = _index128
		var _rhs141 bool = !(!(_921_v9).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg79 _dafny.Int) interface{} {
				return coer79(arg79)
			}
		}((func(_923_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_924_i1 _dafny.Int) _dafny.Int {
				return _923_p0
			}
		})(p0)))).Cardinality())))
		_ = _rhs141
		var _rhs142 _dafny.Int = ((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))).Times(((func() _dafny.Set {
			if _914_v2 {
				return _922_v10
			}
			return _922_v10
		})()).Cardinality())
		_ = _rhs142
		var _rhs143 bool = _914_v2
		_ = _rhs143
		var _rhs144 _dafny.Sequence = _916_v4
		_ = _rhs144
		var _lhs100 _dafny.Array = _913_v1
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
		_ = _lhs101
		(_lhs100).ArraySet1(_rhs141, (_lhs101).Int())
		_915_v3 = _rhs142
		_914_v2 = _rhs143
		_916_v4 = _rhs144
		_915_v3 = (func() _dafny.Int {
			if (p3).Contains(p2) {
				return (p3).Multiplicity(p2)
			}
			return Companion_Default___.SafeDivisionInt(p0, _915_v3)
		})()
		(globalState).F3 = _914_v2
		(globalState).F9 = !((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)) || (_914_v2)
		var _925_v11 D0
		_ = _925_v11
		_925_v11 = Companion_D0_.Create_DC2_(_914_v2)
		var _926_v12 D0
		_ = _926_v12
		_926_v12 = Companion_D0_.Create_DC3_(_925_v11)
		var _source12 D0 = _926_v12
		_ = _source12
		if _source12.Is_DC1() {
			var _927___mcc_h0 bool = _source12.Get_().(D0_DC1).Cf1
			_ = _927___mcc_h0
			var _928___mcc_h1 _dafny.Sequence = _source12.Get_().(D0_DC1).Cf2
			_ = _928___mcc_h1
			var _929___mcc_h2 _dafny.Sequence = _source12.Get_().(D0_DC1).Cf3
			_ = _929___mcc_h2
			var _930_cf3 _dafny.Sequence = _929___mcc_h2
			_ = _930_cf3
			var _931_cf2 _dafny.Sequence = _928___mcc_h1
			_ = _931_cf2
			var _932_cf1 bool = _927___mcc_h0
			_ = _932_cf1
			_915_v3 = Companion_Default___.SafeDivisionInt(_915_v3, (_dafny.Zero).Minus(p2))
			_915_v3 = ((_dafny.Zero).Minus(p0)).Times((p2).Times(p0))
			var _933_v13 D4
			_ = _933_v13
			_933_v13 = Companion_D4_.Create_DC11_()
			var _source13 D4 = _933_v13
			_ = _source13
			if _source13.Is_DC11() {
				var _934_v14 _dafny.Array
				_ = _934_v14
				var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw121
				_934_v14 = _nw121
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_934_v14), 0))
				_ = _index129
				(_934_v14).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lklrsh")).Cardinality()), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_934_v14), 0))
				_ = _index130
				(_934_v14).ArraySet1(p0, (_index130).Int())
				var _935_v15 _dafny.CodePoint
				_ = _935_v15
				_935_v15 = _dafny.CodePoint('n')
				var _936_v16 _dafny.MultiSet
				_ = _936_v16
				_936_v16 = _dafny.MultiSetOf(_935_v15)
				_915_v3 = (_936_v16).Cardinality()
				var _937_v17 _dafny.Map
				_ = _937_v17
				_937_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_932_cf1, _933_v13)
				var _938_v18 *C3
				_ = _938_v18
				var _nw122 *C3 = New_C3_()
				_ = _nw122
				_nw122.Ctor__((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (func() _dafny.Int {
					if (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool) {
						return p2
					}
					return (_937_v17).Cardinality()
				})())
				_938_v18 = _nw122
				_921_v9 = Companion_Default___.Fm10((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), p2, _932_cf1, (p2).Cmp(p2) >= 0, globalState)
			} else if _source13.Is_DC12() {
				var _939___mcc_h6 _dafny.Int = _source13.Get_().(D4_DC12).Cf18
				_ = _939___mcc_h6
				var _940_cf18 _dafny.Int = _939___mcc_h6
				_ = _940_cf18
				_921_v9 = _921_v9
				var _941_v19 _dafny.Array
				_ = _941_v19
				var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw123
				_941_v19 = _nw123
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_941_v19), 0))
				_ = _index131
				(_941_v19).ArraySet1((_940_cf18).Minus(_dafny.IntOfInt64(-548)), (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_941_v19), 0))
				_ = _index132
				(_941_v19).ArraySet1((_dafny.Zero).Minus(p0), (_index132).Int())
				(globalState).F13 = (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)
				var _942_v20 _dafny.Sequence
				_ = _942_v20
				_942_v20 = _dafny.SeqOf(_932_cf1)
				var _943_v21 _dafny.MultiSet
				_ = _943_v21
				_943_v21 = _dafny.MultiSetOf(_911_v0, _911_v0, _911_v0)
				var _944_v22 _dafny.Map
				_ = _944_v22
				_944_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v20, (_943_v21).Union(_943_v21))
				_944_v22 = (_944_v22).Update(_942_v20, (func() _dafny.MultiSet {
					if (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool) {
						return _943_v21
					}
					return _943_v21
				})())
			} else {
				var _945___mcc_h7 _dafny.Set = _source13.Get_().(D4_DC10).Cf17
				_ = _945___mcc_h7
				var _946_cf17 _dafny.Set = _945___mcc_h7
				_ = _946_cf17
				var _947_v23 _dafny.CodePoint
				_ = _947_v23
				_947_v23 = _dafny.CodePoint('e')
				var _948_v25 _dafny.Sequence
				_ = _948_v25
				_948_v25 = _dafny.SeqOf(_946_cf17)
				_947_v23 = (_930_cf3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool) {
						return (func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter31 := _dafny.Iterate((_dafny.MultiSetFromSeq(_948_v25)).Elements()); ; {
								_compr_29, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _949_v24 _dafny.Set
								_949_v24 = interface{}(_compr_29).(_dafny.Set)
								if (_dafny.MultiSetFromSeq(_948_v25)).Contains(_949_v24) {
									_coll29.Add(_949_v24, _915_v3)
								}
							}
							return _coll29.ToMap()
						}()).Cardinality()
					}
					return p0
				})(), _dafny.IntOfUint32((_930_cf3).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _950_v26 _dafny.Map
				_ = _950_v26
				_950_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_932_cf1), _dafny.IntOfInt64(-362))
				var _951_v27 _dafny.Map
				_ = _951_v27
				_951_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _952_v28 _dafny.Array
				_ = _952_v28
				var _nwElement0_23 _dafny.Int = p0
				_ = _nwElement0_23
				var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(11))
				_ = _nw124
				(_nw124).ArraySet1(_nwElement0_23, 0)
				(_nw124).ArraySet1((_950_v26).Cardinality(), 1)
				(_nw124).ArraySet1((_dafny.Zero).Minus(p2), 2)
				(_nw124).ArraySet1(_dafny.IntOfInt64(578), 3)
				(_nw124).ArraySet1(p0, 4)
				(_nw124).ArraySet1(p0, 5)
				(_nw124).ArraySet1(p2, 6)
				(_nw124).ArraySet1(_915_v3, 7)
				(_nw124).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v3, _915_v3)).Merge(_951_v27)).Cardinality(), 8)
				(_nw124).ArraySet1((_dafny.Zero).Minus(_915_v3), 9)
				(_nw124).ArraySet1(Companion_Default___.SafeModuloInt(p0, p2), 10)
				_952_v28 = _nw124
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_952_v28), 0))
				_ = _index133
				(_952_v28).ArraySet1(p0, (_index133).Int())
				var _953_v29 D0
				_ = _953_v29
				_953_v29 = Companion_D0_.Create_DC0_(p2)
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_952_v28), 0))
				_ = _index134
				(_952_v28).ArraySet1((_dafny.IntOfInt64(-951)).Times(_915_v3), (_index134).Int())
				var _954_v30 _dafny.MultiSet
				_ = _954_v30
				_954_v30 = _dafny.MultiSetOf(_913_v1)
				var _955_v31 _dafny.Sequence
				_ = _955_v31
				_955_v31 = _dafny.SeqOf(!(_914_v2))
				var _956_v32 _dafny.Map
				_ = _956_v32
				_956_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_955_v31, _915_v3)
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_952_v28), 0))
				_ = _index135
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_952_v28), 0))
				_ = _index136
				var _rhs145 bool = Companion_Default___.Fm19((_dafny.Zero).Minus(_915_v3), _914_v2, globalState)
				_ = _rhs145
				var _rhs146 _dafny.Int = p2
				_ = _rhs146
				var _rhs147 D0 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(p0))
				_ = _rhs147
				var _rhs148 _dafny.Int = (func() _dafny.Int {
					if (_956_v32).Contains(_dafny.Companion_Sequence_.Concatenate(_955_v31, _955_v31)) {
						return (_956_v32).Get(_dafny.Companion_Sequence_.Concatenate(_955_v31, _955_v31)).(_dafny.Int)
					}
					return (_951_v27).Cardinality()
				})()
				_ = _rhs148
				var _rhs149 _dafny.MultiSet = _954_v30
				_ = _rhs149
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				var _lhs103 _dafny.Array = _952_v28
				_ = _lhs103
				var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_952_v28), 0))
				_ = _lhs104
				var _lhs105 _dafny.Array = _952_v28
				_ = _lhs105
				var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_952_v28), 0))
				_ = _lhs106
				_lhs102.F3 = _rhs145
				(_lhs103).ArraySet1(_rhs146, (_lhs104).Int())
				_953_v29 = _rhs147
				(_lhs105).ArraySet1(_rhs148, (_lhs106).Int())
				_954_v30 = _rhs149
				var _957_v33 D0
				_ = _957_v33
				_957_v33 = Companion_D0_.Create_DC2_(_932_cf1)
				(globalState).F9 = ((_this).Fm5(_957_v33, globalState)).Cmp((_946_cf17).Cardinality()) > 0
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_952_v28), 0))
				_ = _index137
				(_952_v28).ArraySet1(Companion_Default___.Fm11((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), globalState), (_index137).Int())
			}
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
			_ = _index138
			(_913_v1).ArraySet1((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (_index138).Int())
		} else if _source12.Is_DC2() {
			var _958___mcc_h3 bool = _source12.Get_().(D0_DC2).Cf4
			_ = _958___mcc_h3
			var _959_cf4 bool = _958___mcc_h3
			_ = _959_cf4
			var _960_v35 _dafny.Sequence
			_ = _960_v35
			_960_v35 = _dafny.SeqOf(_915_v3, p2)
			var _961_v36 _dafny.Map
			_ = _961_v36
			_961_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_960_v35).Cardinality()), p0)
			var _962_v37 *C4
			_ = _962_v37
			var _nw125 *C4 = New_C4_()
			_ = _nw125
			_nw125.Ctor__(_dafny.IntOfInt64(13), func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter32 := _dafny.Iterate(((p1).Fm1(_961_v36, globalState)).Elements()); ; {
					_compr_30, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _963_v34 _dafny.CodePoint
					_963_v34 = interface{}(_compr_30).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains((p1).Fm1(_961_v36, globalState), _963_v34) {
						_coll30.Add(_963_v34, (_dafny.Zero).Minus(p2))
					}
				}
				return _coll30.ToMap()
			}())
			_962_v37 = _nw125
			_959_cf4 = !((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool))
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
			_ = _index139
			(_913_v1).ArraySet1(_914_v2, (_index139).Int())
			_915_v3 = (_962_v37).F14()
		} else if _source12.Is_DC0() {
			var _964___mcc_h4 _dafny.Int = _source12.Get_().(D0_DC0).Cf0
			_ = _964___mcc_h4
			var _965_cf0 _dafny.Int = _964___mcc_h4
			_ = _965_cf0
			var _966_v38 _dafny.CodePoint
			_ = _966_v38
			_966_v38 = _dafny.CodePoint('j')
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_911_v0), 0))
			_ = _index140
			(_911_v0).ArraySet1CodePoint(_966_v38, (_index140).Int())
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_911_v0), 0))
			_ = _index141
			(_911_v0).ArraySet1CodePoint(_966_v38, (_index141).Int())
			(globalState).F3 = (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)
			var _source14 D4 = Companion_D4_.Create_DC10_((func() _dafny.Set {
				var _coll31 = _dafny.NewBuilder()
				_ = _coll31
				for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(21), _dafny.IntOfInt64(-521))); ; {
					_compr_31, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _967_v39 _dafny.Int
					_967_v39 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(21)).Cmp(_967_v39) <= 0) && ((_967_v39).Cmp(_dafny.IntOfInt64(-521)) < 0) {
						_coll31.Add(Companion_Default___.SafeDivisionInt(_967_v39, _915_v3))
					}
				}
				return _coll31.ToSet()
			}()).Intersection(_921_v9))
			_ = _source14
			if _source14.Is_DC11() {
				var _968_v40 _dafny.Sequence
				_ = _968_v40
				_968_v40 = _dafny.SeqOf((_dafny.Zero).Minus(p0), p0, (_922_v10).Cardinality())
				var _969_v41 _dafny.Map
				_ = _969_v41
				_969_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, _968_v40)
				var _970_v42 D6
				_ = _970_v42
				_970_v42 = Companion_D6_.Create_DC16_(_dafny.IntOfUint32((_968_v40).Cardinality()))
				var _971_v43 D0
				_ = _971_v43
				_971_v43 = Companion_D0_.Create_DC0_(p0)
				var _972_v44 _dafny.Array
				_ = _972_v44
				var _nwElement0_24 _dafny.Sequence = _968_v40
				_ = _nwElement0_24
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(21))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_24, 0)
				(_nw126).ArraySet1((func() _dafny.Sequence {
					if (_969_v41).Contains(_914_v2) {
						return (_969_v41).Get(_914_v2).(_dafny.Sequence)
					}
					return _968_v40
				})(), 1)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}(func(_973_i2 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vcqodo")).Cardinality()))
				})), _968_v40), 2)
				(_nw126).ArraySet1(_968_v40, 3)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Update(_968_v40, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_968_v40).Cardinality()))).Uint32(), (_dafny.MultiSetOf(_914_v2, (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), _914_v2)).Cardinality()), 4)
				(_nw126).ArraySet1(_968_v40, 5)
				(_nw126).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-667))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}(func(_974_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(8)
				})), 6)
				(_nw126).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(268))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg82 _dafny.Int) interface{} {
						return coer82(arg82)
					}
				}((func(_975_cf0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_976_i4 _dafny.Int) _dafny.Int {
						return _975_cf0
					}
				})(_965_cf0))), 7)
				(_nw126).ArraySet1(_968_v40, 8)
				(_nw126).ArraySet1(_968_v40, 9)
				(_nw126).ArraySet1(_968_v40, 10)
				(_nw126).ArraySet1(_968_v40, 11)
				(_nw126).ArraySet1(_968_v40, 12)
				(_nw126).ArraySet1(_dafny.SeqOf((_970_v42).Dtor_cf24(), p2, _965_cf0), 13)
				(_nw126).ArraySet1(_968_v40, 14)
				(_nw126).ArraySet1((func() _dafny.Sequence {
					if _914_v2 {
						return _968_v40
					}
					return _968_v40
				})(), 15)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_977_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_978_i5 _dafny.Int) _dafny.Int {
						return _977_v3
					}
				})(_915_v3))), _968_v40), 16)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, _965_cf0, (_971_v43).Dtor_cf0()), _968_v40), 17)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_968_v40, _968_v40), 18)
				(_nw126).ArraySet1(_968_v40, 19)
				(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm30(globalState), _dafny.Companion_Sequence_.Update(_968_v40, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_968_v40).Cardinality()))).Uint32(), _dafny.IntOfUint32((Companion_Default___.Fm20(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ybsadm")).Cardinality()), _914_v2, _914_v2, _915_v3, globalState)).Cardinality()))), 20)
				_972_v44 = _nw126
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_972_v44), 0))
				_ = _index142
				(_972_v44).ArraySet1(_968_v40, (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_972_v44), 0))
				_ = _index143
				(_972_v44).ArraySet1(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(547), p0)).Cardinality(), _965_cf0, _915_v3), (_index143).Int())
				var _979_v46 _dafny.MultiSet
				_ = _979_v46
				_979_v46 = _dafny.MultiSetOf(_966_v38)
				var _980_v47 _dafny.Sequence
				_ = _980_v47
				_980_v47 = _dafny.SeqOf(!(!((func() bool {
					if (_917_v5).Contains(p2) {
						return (_917_v5).Get(p2).(bool)
					}
					return (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)
				})())), _914_v2)
				var _rhs150 bool = !_dafny.Companion_Sequence_.Contains(_916_v4, Companion_Default___.Fm23(!(!(_914_v2)), func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter34 := _dafny.Iterate(((_979_v46).Update((_911_v0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_911_v0), 0))).Int()), Companion_Default___.Abs(p2))).Elements()); ; {
						_compr_32, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _981_v45 _dafny.CodePoint
						_981_v45 = interface{}(_compr_32).(_dafny.CodePoint)
						if ((_979_v46).Update((_911_v0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_911_v0), 0))).Int()), Companion_Default___.Abs(p2))).Contains(_981_v45) {
							_coll32.Add(_981_v45, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)))
						}
					}
					return _coll32.ToMap()
				}(), _914_v2, p2, globalState))
				_ = _rhs150
				var _rhs151 _dafny.Int = (_dafny.IntOfUint32((_980_v47).Cardinality())).Plus(_915_v3)
				_ = _rhs151
				var _rhs152 _dafny.Int = _dafny.IntOfInt64(831)
				_ = _rhs152
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				_lhs107.F9 = _rhs150
				_915_v3 = _rhs151
				_915_v3 = _rhs152
				(globalState).F13 = (_914_v2) && (!((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)))
				var _982_v48 D7
				_ = _982_v48
				_982_v48 = Companion_D7_.Create_DC18_(_917_v5)
				var _983_v49 _dafny.Map
				_ = _983_v49
				_983_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_982_v48, p1)
				_915_v3 = (_dafny.SetOf(!(_983_v49).Contains(_982_v48), false)).Cardinality()
			} else if _source14.Is_DC12() {
				var _984___mcc_h8 _dafny.Int = _source14.Get_().(D4_DC12).Cf18
				_ = _984___mcc_h8
				var _985_cf18 _dafny.Int = _984___mcc_h8
				_ = _985_cf18
				(globalState).F3 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_916_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_986_v0 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_987_i6 _dafny.Int) _dafny.CodePoint {
						return (_986_v0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_986_v0), 0))).Int())
					}
				})(_911_v0)))), _916_v4)
				_919_v7 = Companion_D7_.Create_DC21_(_918_v6)
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
				_ = _index144
				(_913_v1).ArraySet1((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (_index144).Int())
				var _988_v50 D4
				_ = _988_v50
				_988_v50 = Companion_D4_.Create_DC11_()
				var _989_v51 _dafny.Array
				_ = _989_v51
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw127
				_989_v51 = _nw127
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_989_v51), 0))
				_ = _index145
				(_989_v51).ArraySet1(_965_cf0, (_index145).Int())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_989_v51), 0))
				_ = _index146
				var _rhs153 bool = !((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool))
				_ = _rhs153
				var _rhs154 D4 = _988_v50
				_ = _rhs154
				var _rhs155 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(957))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_990_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_991_i7 _dafny.Int) _dafny.CodePoint {
						return _990_v38
					}
				})(_966_v38))), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_916_v4, (Companion_Default___.SafeIndex(_985_cf18, _dafny.IntOfUint32((_916_v4).Cardinality()))).Uint32(), _dafny.CodePoint('u')), _dafny.UnicodeSeqOfUtf8Bytes("qwrmofy")))).Cardinality())
				_ = _rhs155
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				var _lhs109 _dafny.Array = _989_v51
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_989_v51), 0))
				_ = _lhs110
				_lhs108.F9 = _rhs153
				_988_v50 = _rhs154
				(_lhs109).ArraySet1(_rhs155, (_lhs110).Int())
			} else {
				var _992___mcc_h9 _dafny.Set = _source14.Get_().(D4_DC10).Cf17
				_ = _992___mcc_h9
				var _993_cf17 _dafny.Set = _992___mcc_h9
				_ = _993_cf17
				var _994_v52 _dafny.Array
				_ = _994_v52
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw128
				_994_v52 = _nw128
				_994_v52 = _994_v52
				var _995_v53 _dafny.Map
				_ = _995_v53
				_995_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v3, _916_v4)
				var _996_v54 _dafny.Map
				_ = _996_v54
				_996_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v1, _dafny.IntOfInt64(198))
				var _997_v55 _dafny.Map
				_ = _997_v55
				_997_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_996_v54).Cardinality(), p2)
				_995_v53 = (_995_v53).Update((_dafny.Zero).Minus(p0), (p1).Fm1(_997_v55, globalState))
				var _998_v56 _dafny.Array
				_ = _998_v56
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_17
				var _nw129 _dafny.Array
				_ = _nw129
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw129 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) D2 = (func(_999_v2 bool, _1000_v12 D0, _1001_v4 _dafny.Sequence, _1002_v38 _dafny.CodePoint) func(_dafny.Int) D2 {
						return func(_1003_i8 _dafny.Int) D2 {
							return Companion_D2_.Create_DC8_(Companion_D2_.Create_DC7_(_999_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v2, _1000_v12), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1001_v4, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1001_v4).Cardinality()), _dafny.IntOfUint32((_1001_v4).Cardinality()))).Uint32(), _1002_v38)).Cardinality()), _999_v2))
						}
					})(_914_v2, _926_v12, _916_v4, _966_v38)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw129 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw129).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw129).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_998_v56 = _nw129
				var _1004_v57 _dafny.Map
				_ = _1004_v57
				_1004_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_914_v2) == ((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)), _998_v56)
				_1004_v57 = (_1004_v57).Update((func() bool {
					if _914_v2 {
						return _914_v2
					}
					return _914_v2
				})(), _998_v56)
				var _1005_v58 _dafny.Array
				_ = _1005_v58
				var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw130
				_1005_v58 = _nw130
				var _1006_v59 _dafny.Map
				_ = _1006_v59
				_1006_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), _1005_v58)
				_1005_v58 = (func() _dafny.Array {
					if (_1006_v59).Contains((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)) {
						return (_1006_v59).Get((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)).(_dafny.Array)
					}
					return _1005_v58
				})()
			}
			_965_cf0 = (_dafny.Zero).Minus(p0)
		} else {
			var _1007___mcc_h5 D0 = _source12.Get_().(D0_DC3).Cf5
			_ = _1007___mcc_h5
			var _1008_cf5 D0 = _1007___mcc_h5
			_ = _1008_cf5
			var _1009_v60 _dafny.CodePoint
			_ = _1009_v60
			_1009_v60 = _dafny.CodePoint('t')
			var _1010_v61 _dafny.Map
			_ = _1010_v61
			_1010_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1009_v60, _916_v4)
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
			_ = _index147
			(_913_v1).ArraySet1((_1010_v61).Contains(_dafny.CodePoint('s')), (_index147).Int())
			(globalState).F9 = _914_v2
			var _1011_v62 _dafny.Sequence
			_ = _1011_v62
			_1011_v62 = _dafny.SeqOf(_916_v4)
			var _1012_v63 _dafny.Map
			_ = _1012_v63
			_1012_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v3, _916_v4)
			var _1013_v64 D0
			_ = _1013_v64
			_1013_v64 = Companion_D0_.Create_DC1_((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), _1011_v62, (func() _dafny.Sequence {
				if (_1012_v63).Contains(p0) {
					return (_1012_v63).Get(p0).(_dafny.Sequence)
				}
				return _916_v4
			})())
			var _1014_v65 _dafny.MultiSet
			_ = _1014_v65
			_1014_v65 = _dafny.MultiSetOf(_1013_v64)
			var _1015_v66 D2
			_ = _1015_v66
			_1015_v66 = Companion_D2_.Create_DC6_(_1014_v65)
			var _pat_let_tv21 = _1014_v65
			_ = _pat_let_tv21
			var _1016_v67 D2
			_ = _1016_v67
			_1016_v67 = Companion_D2_.Create_DC6_((func(_pat_let19_0 D2) D2 {
				return func(_1017_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let20_0 _dafny.MultiSet) D2 {
						return func(_1018_dt__update_hcf10_h0 _dafny.MultiSet) D2 {
							return Companion_D2_.Create_DC6_(_1018_dt__update_hcf10_h0)
						}(_pat_let20_0)
					}(_pat_let_tv21)
				}(_pat_let19_0)
			}(_1015_v66)).Dtor_cf10())
			var _source15 D2 = _1015_v66
			_ = _source15
			if _source15.Is_DC7() {
				var _1019___mcc_h10 bool = _source15.Get_().(D2_DC7).Cf11
				_ = _1019___mcc_h10
				var _1020___mcc_h11 _dafny.Map = _source15.Get_().(D2_DC7).Cf12
				_ = _1020___mcc_h11
				var _1021___mcc_h12 _dafny.Int = _source15.Get_().(D2_DC7).Cf13
				_ = _1021___mcc_h12
				var _1022___mcc_h13 bool = _source15.Get_().(D2_DC7).Cf14
				_ = _1022___mcc_h13
				var _1023_cf14 bool = _1022___mcc_h13
				_ = _1023_cf14
				var _1024_cf13 _dafny.Int = _1021___mcc_h12
				_ = _1024_cf13
				var _1025_cf12 _dafny.Map = _1020___mcc_h11
				_ = _1025_cf12
				var _1026_cf11 bool = _1019___mcc_h10
				_ = _1026_cf11
				var _1027_v68 _dafny.Array
				_ = _1027_v68
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_18
				var _nw131 _dafny.Array
				_ = _nw131
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw131 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Int = (func(_1028_cf13 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1029_i9 _dafny.Int) _dafny.Int {
							return (_1029_i9).Minus(_1028_cf13)
						}
					})(_1024_cf13)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw131 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw131).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw131).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_1027_v68 = _nw131
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_1027_v68), 0))
				_ = _index148
				(_1027_v68).ArraySet1(_1024_cf13, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1027_v68), 0))
				_ = _index149
				(_1027_v68).ArraySet1(p2, (_index149).Int())
				var _1030_v69 _dafny.Map
				_ = _1030_v69
				_1030_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v1, _1023_cf14)
				var _1031_v70 _dafny.Map
				_ = _1031_v70
				_1031_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1026_cf11, _1024_cf13)
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_1027_v68), 0))
				_ = _index150
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1027_v68), 0))
				_ = _index151
				var _rhs156 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v1, _914_v2)).Merge(_1030_v69)).Cardinality(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-701), (func() _dafny.Int {
					if (_1031_v70).Contains(_914_v2) {
						return (_1031_v70).Get(_914_v2).(_dafny.Int)
					}
					return _dafny.IntOfInt64(214)
				})())))
				_ = _rhs156
				var _rhs157 bool = (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)
				_ = _rhs157
				var _rhs158 _dafny.Int = _dafny.IntOfUint32((_916_v4).Cardinality())
				_ = _rhs158
				var _rhs159 _dafny.Int = _915_v3
				_ = _rhs159
				var _lhs111 _dafny.Array = _1027_v68
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_1027_v68), 0))
				_ = _lhs112
				var _lhs113 _dafny.Array = _1027_v68
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1027_v68), 0))
				_ = _lhs114
				_1024_cf13 = _rhs156
				_914_v2 = _rhs157
				(_lhs111).ArraySet1(_rhs158, (_lhs112).Int())
				(_lhs113).ArraySet1(_rhs159, (_lhs114).Int())
				_913_v1 = _913_v1
				var _1032_v71 _dafny.Sequence
				_ = _1032_v71
				_1032_v71 = _dafny.SeqOf(_1026_cf11)
				var _1033_v72 _dafny.Sequence
				_ = _1033_v72
				_1033_v72 = _dafny.SeqOf(_1032_v71)
				var _1034_v73 D2
				_ = _1034_v73
				_1034_v73 = Companion_D2_.Create_DC7_((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), _1025_cf12, Companion_Default___.Fm11(_1023_cf14, globalState), _1026_cf11)
				var _pat_let_tv22 = _1025_cf12
				_ = _pat_let_tv22
				_1024_cf13 = (_dafny.IntOfInt64(143)).Plus(_dafny.IntOfUint32(((_1033_v72).Select((Companion_Default___.SafeIndex((func(_pat_let21_0 D2) D2 {
					return func(_1035_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let22_0 _dafny.Map) D2 {
							return func(_1036_dt__update_hcf12_h0 _dafny.Map) D2 {
								return Companion_D2_.Create_DC7_((_1035_dt__update__tmp_h1).Dtor_cf11(), _1036_dt__update_hcf12_h0, (_1035_dt__update__tmp_h1).Dtor_cf13(), (_1035_dt__update__tmp_h1).Dtor_cf14())
							}(_pat_let22_0)
						}(_pat_let_tv22)
					}(_pat_let21_0)
				}(_1034_v73)).Dtor_cf13(), _dafny.IntOfUint32((_1033_v72).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _1037_v74 *C1
				_ = _1037_v74
				var _nw132 *C1 = New_C1_()
				_ = _nw132
				_nw132.Ctor__()
				_1037_v74 = _nw132
			} else if _source15.Is_DC6() {
				var _1038___mcc_h14 _dafny.MultiSet = _source15.Get_().(D2_DC6).Cf10
				_ = _1038___mcc_h14
				var _1039_cf10 _dafny.MultiSet = _1038___mcc_h14
				_ = _1039_cf10
				var _1040_v75 *C1
				_ = _1040_v75
				var _nw133 *C1 = New_C1_()
				_ = _nw133
				_nw133.Ctor__()
				_1040_v75 = _nw133
				(globalState).F13 = true
				var _1041_v76 D2
				_ = _1041_v76
				_1041_v76 = Companion_D2_.Create_DC7_(_914_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, Companion_D0_.Create_DC3_(_925_v11)), p0, _914_v2)
				var _1042_v77 D2
				_ = _1042_v77
				_1042_v77 = Companion_D2_.Create_DC8_(_1041_v76)
				var _1043_v78 _dafny.Map
				_ = _1043_v78
				_1043_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1042_v77, p2)
				var _1044_v79 _dafny.Map
				_ = _1044_v79
				_1044_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, _915_v3)
				var _1045_v80 _dafny.Map
				_ = _1045_v80
				_1045_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool))
				_1043_v78 = (_1043_v78).Update(Companion_D2_.Create_DC8_(_1041_v76), (func() _dafny.Int {
					if (_1044_v79).Contains((func() bool {
						if (_1045_v80).Contains((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)) {
							return (_1045_v80).Get((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)).(bool)
						}
						return _914_v2
					})()) {
						return (_1044_v79).Get((func() bool {
							if (_1045_v80).Contains((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)) {
								return (_1045_v80).Get((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)).(bool)
							}
							return _914_v2
						})()).(_dafny.Int)
					}
					return p2
				})())
				(globalState).F9 = _914_v2
			} else {
				var _1046___mcc_h15 D2 = _source15.Get_().(D2_DC8).Cf15
				_ = _1046___mcc_h15
				var _1047_cf15 D2 = _1046___mcc_h15
				_ = _1047_cf15
				var _1048_v81 _dafny.Map
				_ = _1048_v81
				_1048_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, _926_v12)
				var _1049_v82 D2
				_ = _1049_v82
				_1049_v82 = Companion_D2_.Create_DC7_(false, _1048_v81, (_dafny.Zero).Minus(p0), _914_v2)
				var _1050_v83 _dafny.Map
				_ = _1050_v83
				_1050_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1009_v60, _914_v2)
				_915_v3 = (_dafny.Zero).Minus(((Companion_Default___.Fm22(_1049_v82, _926_v12, (func() bool {
					if (_1050_v83).Contains(_1009_v60) {
						return (_1050_v83).Get(_1009_v60).(bool)
					}
					return _914_v2
				})(), globalState)).Times(p2)).Plus(p2))
				var _1051_v84 _dafny.Sequence
				_ = _1051_v84
				_1051_v84 = _dafny.SeqOf(_914_v2, _914_v2, _914_v2)
				var _1052_v85 _dafny.Map
				_ = _1052_v85
				_1052_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('k'))
				var _1053_v86 _dafny.Map
				_ = _1053_v86
				_1053_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool))
				var _1054_v87 _dafny.Map
				_ = _1054_v87
				_1054_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v2, !((func() bool {
					if (_1053_v86).Contains((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)) {
						return (_1053_v86).Get((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)).(bool)
					}
					return Companion_Default___.Fm6(globalState)
				})()))
				var _1055_v88 _dafny.Sequence
				_ = _1055_v88
				_1055_v88 = _dafny.SeqOf(p0, _dafny.IntOfInt64(-878), _915_v3, p0, (_1054_v87).Cardinality())
				var _1056_v89 _dafny.Array
				_ = _1056_v89
				var _nwElement0_25 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1051_v84).Cardinality()))
				_ = _nwElement0_25
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_25, 0)
				(_nw134).ArraySet1((_915_v3).Minus(_dafny.IntOfInt64(-518)), 1)
				(_nw134).ArraySet1(_915_v3, 2)
				(_nw134).ArraySet1(_915_v3, 3)
				(_nw134).ArraySet1((_1052_v85).Cardinality(), 4)
				(_nw134).ArraySet1(_915_v3, 5)
				(_nw134).ArraySet1((_1055_v88).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1055_v88).Cardinality()))).Uint32()).(_dafny.Int), 6)
				(_nw134).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvsvmpmou")).Cardinality()), 7)
				(_nw134).ArraySet1(p0, 8)
				(_nw134).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-64)), 9)
				(_nw134).ArraySet1(p0, 10)
				(_nw134).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p2, _915_v3))), 11)
				(_nw134).ArraySet1(Companion_Default___.SafeModuloInt(p0, p2), 12)
				_1056_v89 = _nw134
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
				_ = _index152
				var _rhs160 bool = true
				_ = _rhs160
				var _rhs161 _dafny.Array = _1056_v89
				_ = _rhs161
				var _rhs162 bool = _914_v2
				_ = _rhs162
				var _rhs163 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}((func(_1057_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1058_i10 _dafny.Int) _dafny.CodePoint {
						return _1057_v60
					}
				})(_1009_v60))), _916_v4), _1009_v60)
				_ = _rhs163
				var _rhs164 _dafny.Int = p2
				_ = _rhs164
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				var _lhs117 _dafny.Array = _913_v1
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))
				_ = _lhs118
				_lhs115.F3 = _rhs160
				_1056_v89 = _rhs161
				_lhs116.F5 = _rhs162
				(_lhs117).ArraySet1(_rhs163, (_lhs118).Int())
				_915_v3 = _rhs164
				(globalState).F3 = (_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool)
				var _1059_v90 _dafny.Array
				_ = _1059_v90
				var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw135
				_1059_v90 = _nw135
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1059_v90), 0))
				_ = _index153
				(_1059_v90).ArraySet1(_1056_v89, (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1059_v90), 0))
				_ = _index154
				(_1059_v90).ArraySet1(_1056_v89, (_index154).Int())
			}
			var _1060_v91 _dafny.Array
			_ = _1060_v91
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw136
			_1060_v91 = _nw136
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_1060_v91), 0))
			_ = _index155
			(_1060_v91).ArraySet1(Companion_Default___.Fm11((_913_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_913_v1), 0))).Int()).(bool), globalState), (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_1060_v91), 0))
			_ = _index156
			(_1060_v91).ArraySet1(p0, (_index156).Int())
		}
	}
}
func (_this *C5) M4(p0 _dafny.Sequence, p1 _dafny.Array, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1061_v0 _dafny.Sequence
		_ = _1061_v0
		_1061_v0 = _dafny.UnicodeSeqOfUtf8Bytes("scsqu")
		_1061_v0 = _1061_v0
		var _1062_v1 bool
		_ = _1062_v1
		_1062_v1 = false
		r2 = (_this).Fm5(Companion_D0_.Create_DC2_(_1062_v1), globalState)
		var _1063_v2 _dafny.Int
		_ = _1063_v2
		_1063_v2 = _dafny.IntOfInt64(931)
		var _hi6 _dafny.Int = _1063_v2
		_ = _hi6
		for _1064_i0 := _dafny.IntOfUint32((_1061_v0).Cardinality()); _1064_i0.Cmp(_hi6) < 0; _1064_i0 = _1064_i0.Plus(_dafny.One) {
			r2 = (_1063_v2).Times(_1064_i0)
			var _1065_v3 _dafny.Map
			_ = _1065_v3
			_1065_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v2, Companion_Default___.Fm6(globalState))
			var _1066_v4 _dafny.Array
			_ = _1066_v4
			var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw137
			_1066_v4 = _nw137
			var _1067_v5 _dafny.Array
			_ = _1067_v5
			var _nwElement0_26 bool = _1062_v1
			_ = _nwElement0_26
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(14))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_26, 0)
			(_nw138).ArraySet1(!(_1062_v1), 1)
			(_nw138).ArraySet1(_1062_v1, 2)
			(_nw138).ArraySet1((_1062_v1) && (_1062_v1), 3)
			(_nw138).ArraySet1((_1063_v2).Cmp(_1063_v2) <= 0, 4)
			(_nw138).ArraySet1(false, 5)
			(_nw138).ArraySet1((_1063_v2).Cmp(_dafny.IntOfInt64(807)) < 0, 6)
			(_nw138).ArraySet1(_1062_v1, 7)
			(_nw138).ArraySet1(_1062_v1, 8)
			(_nw138).ArraySet1(_1062_v1, 9)
			(_nw138).ArraySet1(((func() bool {
				if (_1065_v3).Contains(_dafny.IntOfInt64(350)) {
					return (_1065_v3).Get(_dafny.IntOfInt64(350)).(bool)
				}
				return _1062_v1
			})()) == (_1062_v1), 10)
			(_nw138).ArraySet1(_1062_v1, 11)
			(_nw138).ArraySet1(false, 12)
			(_nw138).ArraySet1(false, 13)
			_1067_v5 = _nw138
			var _rhs165 _dafny.Map = (func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(791), _dafny.IntOfInt64(-974))); ; {
					_compr_33, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _1068_v6 _dafny.Int
					_1068_v6 = interface{}(_compr_33).(_dafny.Int)
					if ((_dafny.IntOfInt64(791)).Cmp(_1068_v6) <= 0) && ((_1068_v6).Cmp(_dafny.IntOfInt64(-974)) < 0) {
						_coll33.Add((_1068_v6).Minus(_1063_v2), _1062_v1)
					}
				}
				return _coll33.ToMap()
			}()).Update(_1064_i0, _1062_v1)
			_ = _rhs165
			var _rhs166 _dafny.Int = _1064_i0
			_ = _rhs166
			var _rhs167 _dafny.Array = _1066_v4
			_ = _rhs167
			var _rhs168 _dafny.Array = (func() _dafny.Array {
				if !(!(!(Companion_Default___.Fm6(globalState)))) {
					return _1067_v5
				}
				return p1
			})()
			_ = _rhs168
			_1065_v3 = _rhs165
			r1 = _rhs166
			_1066_v4 = _rhs167
			_1067_v5 = _rhs168
			r2 = _1064_i0
			var _1069_v7 _dafny.Array
			_ = _1069_v7
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw139
			_1069_v7 = _nw139
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_1069_v7), 0))
			_ = _index157
			(_1069_v7).ArraySet1(_1063_v2, (_index157).Int())
			var _1070_v8 _dafny.Sequence
			_ = _1070_v8
			_1070_v8 = _dafny.SeqOf(_1062_v1)
			var _1071_v9 D1
			_ = _1071_v9
			_1071_v9 = Companion_D1_.Create_DC5_(_1062_v1, _1061_v0, _1070_v8)
			var _1072_v10 _dafny.Sequence
			_ = _1072_v10
			_1072_v10 = _dafny.SeqOf(_1062_v1, !((_1071_v9).Dtor_cf7()))
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_1069_v7), 0))
			_ = _index158
			(_1069_v7).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1062_v1, _1062_v1, Companion_Default___.Fm6(globalState)), _1070_v8), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1063_v2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1062_v1, _1062_v1, Companion_Default___.Fm6(globalState)), _1070_v8)).Cardinality()))).Uint32(), _1062_v1), (Companion_Default___.SafeIndex(_1063_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1062_v1, _1062_v1, Companion_Default___.Fm6(globalState)), _1070_v8), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1063_v2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1062_v1, _1062_v1, Companion_Default___.Fm6(globalState)), _1070_v8)).Cardinality()))).Uint32(), _1062_v1)).Cardinality()))).Uint32(), _1062_v1), _1072_v10)).Cardinality()), (_index158).Int())
		}
		var _1073_i1 _dafny.Int
		_ = _1073_i1
		_1073_i1 = _dafny.Zero
		{
			for !(Companion_Default___.Fm6(globalState)) {
				{
					if (_1073_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1073_i1 = (_1073_i1).Plus(_dafny.One)
					var _1074_v11 _dafny.Sequence
					_ = _1074_v11
					_1074_v11 = _dafny.SeqOf(_1062_v1)
					var _1075_v12 _dafny.CodePoint
					_ = _1075_v12
					_1075_v12 = _dafny.CodePoint('c')
					var _source16 D4 = Companion_Default___.Fm35(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1074_v11, _1074_v11), (Companion_Default___.SafeIndex(_1063_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1074_v11, _1074_v11)).Cardinality()))).Uint32(), _1062_v1), !(_1062_v1) || (_1062_v1), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(703))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg87 _dafny.Int) interface{} {
							return coer87(arg87)
						}
					}((func(_1076_v0 _dafny.Sequence, _1077_v2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
						return func(_1078_i2 _dafny.Int) _dafny.CodePoint {
							return (_1076_v0).Select((Companion_Default___.SafeIndex(_1077_v2, _dafny.IntOfUint32((_1076_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
						}
					})(_1061_v0, _1063_v2))), (Companion_Default___.SafeIndex(_1063_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(703))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg88 _dafny.Int) interface{} {
							return coer88(arg88)
						}
					}((func(_1079_v0 _dafny.Sequence, _1080_v2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
						return func(_1081_i2 _dafny.Int) _dafny.CodePoint {
							return (_1079_v0).Select((Companion_Default___.SafeIndex(_1080_v2, _dafny.IntOfUint32((_1079_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
						}
					})(_1061_v0, _1063_v2)))).Cardinality()))).Uint32(), _1075_v12), globalState)
					_ = _source16
					if _source16.Is_DC11() {
						_1061_v0 = _1061_v0
						var _1082_v13 _dafny.MultiSet
						_ = _1082_v13
						_1082_v13 = _dafny.MultiSetOf(_1074_v11, _dafny.Companion_Sequence_.Concatenate(_1074_v11, _1074_v11), _1074_v11)
						var _rhs169 _dafny.Int = (p0).Select((Companion_Default___.SafeIndex(_1063_v2, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
						_ = _rhs169
						var _rhs170 _dafny.MultiSet = (_dafny.MultiSetOf(_1074_v11, _1074_v11, _1074_v11)).Difference(_1082_v13)
						_ = _rhs170
						var _rhs171 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-497))
						_ = _rhs171
						r2 = _rhs169
						_1082_v13 = _rhs170
						_1063_v2 = _rhs171
						var _1083_v14 _dafny.MultiSet
						_ = _1083_v14
						_1083_v14 = _dafny.MultiSetOf(_1061_v0)
						var _rhs172 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
							if Companion_Default___.Fm6(globalState) {
								return _1063_v2
							}
							return (_dafny.Zero).Minus(_1063_v2)
						})(), _1063_v2)
						_ = _rhs172
						var _rhs173 _dafny.MultiSet = _dafny.MultiSetOf(_1061_v0)
						_ = _rhs173
						var _rhs174 _dafny.Int = _1063_v2
						_ = _rhs174
						var _rhs175 bool = Companion_Default___.Fm6(globalState)
						_ = _rhs175
						var _rhs176 bool = !(!(_1062_v1))
						_ = _rhs176
						var _lhs119 *GlobalState = globalState
						_ = _lhs119
						var _lhs120 *GlobalState = globalState
						_ = _lhs120
						r2 = _rhs172
						_1083_v14 = _rhs173
						r1 = _rhs174
						_lhs119.F3 = _rhs175
						_lhs120.F3 = _rhs176
						var _1084_v15 _dafny.Map
						_ = _1084_v15
						_1084_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v1, _1062_v1)
						var _1085_v16 _dafny.Map
						_ = _1085_v16
						_1085_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1075_v12, _1084_v15)
						var _1086_v17 _dafny.Array
						_ = _1086_v17
						var _nwElement0_27 _dafny.CodePoint = _1075_v12
						_ = _nwElement0_27
						var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(19))
						_ = _nw140
						(_nw140).ArraySet1CodePoint(_nwElement0_27, 0)
						(_nw140).ArraySet1CodePoint(_1075_v12, 1)
						(_nw140).ArraySet1CodePoint(_1075_v12, 2)
						(_nw140).ArraySet1CodePoint(_1075_v12, 3)
						(_nw140).ArraySet1CodePoint(_1075_v12, 4)
						(_nw140).ArraySet1CodePoint(_1075_v12, 5)
						(_nw140).ArraySet1CodePoint(_1075_v12, 6)
						(_nw140).ArraySet1CodePoint(_1075_v12, 7)
						(_nw140).ArraySet1CodePoint(_1075_v12, 8)
						(_nw140).ArraySet1CodePoint(_1075_v12, 9)
						(_nw140).ArraySet1CodePoint(_1075_v12, 10)
						(_nw140).ArraySet1CodePoint(_dafny.CodePoint('a'), 11)
						(_nw140).ArraySet1CodePoint(_1075_v12, 12)
						(_nw140).ArraySet1CodePoint(_1075_v12, 13)
						(_nw140).ArraySet1CodePoint(_1075_v12, 14)
						(_nw140).ArraySet1CodePoint(Companion_Default___.Fm23(_1062_v1, _1085_v16, _1062_v1, _1063_v2, globalState), 15)
						(_nw140).ArraySet1CodePoint(_1075_v12, 16)
						(_nw140).ArraySet1CodePoint(_1075_v12, 17)
						(_nw140).ArraySet1CodePoint(_dafny.CodePoint('o'), 18)
						_1086_v17 = _nw140
						var _1087_v20 _dafny.Set
						_ = _1087_v20
						_1087_v20 = _dafny.SetOf(_1063_v2)
						var _1088_v21 _dafny.Map
						_ = _1088_v21
						_1088_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1075_v12, func() _dafny.Map {
							var _coll34 = _dafny.NewMapBuilder()
							_ = _coll34
							for _iter36 := _dafny.Iterate((_1087_v20).Elements()); ; {
								_compr_34, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _1089_v19 _dafny.Int
								_1089_v19 = interface{}(_compr_34).(_dafny.Int)
								if (_1087_v20).Contains(_1089_v19) {
									_coll34.Add(Companion_Default___.SafeDivisionInt(_1089_v19, _1063_v2), _1062_v1)
								}
							}
							return _coll34.ToMap()
						}())
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_1086_v17), 0))
						_ = _index159
						(_1086_v17).ArraySet1CodePoint(Companion_Default___.Fm23(_1062_v1, func() _dafny.Map {
							var _coll35 = _dafny.NewMapBuilder()
							_ = _coll35
							for _iter37 := _dafny.Iterate((_1088_v21).Keys().Elements()); ; {
								_compr_35, _ok37 := _iter37()
								if !_ok37 {
									break
								}
								var _1090_v18 _dafny.CodePoint
								_1090_v18 = interface{}(_compr_35).(_dafny.CodePoint)
								if (_1088_v21).Contains(_1090_v18) {
									_coll35.Add(_1090_v18, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v1, _1062_v1)).Update(_1062_v1, _1062_v1))
								}
							}
							return _coll35.ToMap()
						}(), _1062_v1, _1063_v2, globalState), (_index159).Int())
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_1086_v17), 0))
						_ = _index160
						(_1086_v17).ArraySet1CodePoint(_1075_v12, (_index160).Int())
					} else if _source16.Is_DC12() {
						var _1091___mcc_h0 _dafny.Int = _source16.Get_().(D4_DC12).Cf18
						_ = _1091___mcc_h0
						var _1092_cf18 _dafny.Int = _1091___mcc_h0
						_ = _1092_cf18
						(globalState).F3 = _1062_v1
						r2 = Companion_Default___.Fm11(_1062_v1, globalState)
						(globalState).F13 = true
						_1062_v1 = _1062_v1
					} else {
						var _1093___mcc_h1 _dafny.Set = _source16.Get_().(D4_DC10).Cf17
						_ = _1093___mcc_h1
						var _1094_cf17 _dafny.Set = _1093___mcc_h1
						_ = _1094_cf17
						var _1095_v22 _dafny.Map
						_ = _1095_v22
						_1095_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1094_cf17).Intersection(_1094_cf17), _dafny.UnicodeSeqOfUtf8Bytes("bgddjiqf"))
						_1095_v22 = (_1095_v22).Update(_1094_cf17, _dafny.Companion_Sequence_.Concatenate(_1061_v0, _1061_v0))
						var _1096_v23 _dafny.Set
						_ = _1096_v23
						_1096_v23 = _dafny.SetOf(_1062_v1)
						var _1097_v24 _dafny.Map
						_ = _1097_v24
						_1097_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v1, _1074_v11)).Cardinality(), (_1063_v2).Times((_1096_v23).Cardinality()))
						_1097_v24 = _1097_v24
						_1062_v1 = false
						r2 = _1063_v2
					}
					var _1098_v25 _dafny.Array
					_ = _1098_v25
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(3)
					_ = _len0_19
					var _nw141 _dafny.Array
					_ = _nw141
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw141 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) _dafny.Int = (func(_1099_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1100_i3 _dafny.Int) _dafny.Int {
								return (_1100_i3).Plus(_1099_v2)
							}
						})(_1063_v2)
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw141 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw141).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw141).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_1098_v25 = _nw141
					var _1101_v26 _dafny.Map
					_ = _1101_v26
					_1101_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(_1063_v2, _1075_v12, globalState), false)).Cardinality())
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1098_v25), 0))
					_ = _index161
					(_1098_v25).ArraySet1((_1101_v26).Cardinality(), (_index161).Int())
					var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1098_v25), 0))
					_ = _index162
					(_1098_v25).ArraySet1(_dafny.IntOfUint32((_1061_v0).Cardinality()), (_index162).Int())
					var _1102_v27 _dafny.MultiSet
					_ = _1102_v27
					_1102_v27 = _dafny.MultiSetOf(_dafny.IntOfInt64(468))
					_1102_v27 = _dafny.MultiSetOf((_1098_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1098_v25), 0))).Int()).(_dafny.Int), (_1098_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1098_v25), 0))).Int()).(_dafny.Int))
					_1075_v12 = _dafny.CodePoint('a')
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		r2 = _1063_v2
		var _1103_v28 _dafny.Set
		_ = _1103_v28
		_1103_v28 = _dafny.SetOf((_1063_v2).Cmp(_1063_v2) == 0)
		_1103_v28 = Companion_Default___.Fm33(_1062_v1, _1062_v1, _1062_v1, globalState)
		r0 = false
		var _1104_v29 D0
		_ = _1104_v29
		_1104_v29 = Companion_D0_.Create_DC0_(_1063_v2)
		var _1105_v30 D0
		_ = _1105_v30
		_1105_v30 = Companion_D0_.Create_DC3_(_1104_v29)
		var _1106_v31 _dafny.Map
		_ = _1106_v31
		_1106_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v1, _1105_v30)
		var _1107_v32 D2
		_ = _1107_v32
		_1107_v32 = Companion_D2_.Create_DC7_(_1062_v1, _1106_v31, _1063_v2, false)
		r1 = Companion_Default___.Fm22(_1107_v32, Companion_D0_.Create_DC3_(_1104_v29), (_dafny.IntOfUint32((p0).Cardinality())).Cmp(_dafny.IntOfInt64(229)) == 0, globalState)
		r2 = _1063_v2
		return r0, r1, r2
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	dummy byte
}

func New_C6_() *C6 {
	_this := C6{}

	return &_this
}

type CompanionStruct_C6_ struct {
}

var Companion_C6_ = CompanionStruct_C6_{}

func (_this *C6) Equals(other *C6) bool {
	return _this == other
}

func (_this *C6) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C6)
	return ok && _this.Equals(other)
}

func (*C6) String() string {
	return "_module.C6"
}

func Type_C6_() _dafny.TypeDescriptor {
	return type_C6_{}
}

type type_C6_ struct {
}

func (_this type_C6_) Default() interface{} {
	return (*C6)(nil)
}

func (_this type_C6_) String() string {
	return "main.C6"
}
func (_this *C6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__() {
	{
	}
}
func (_this *C6) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-869))), (Companion_D0_.Create_DC0_(_dafny.IntOfInt64(211))).Dtor_cf0())).Times(_dafny.IntOfInt64(-444))
	}
}
func (_this *C6) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tiwttktph"), _dafny.UnicodeSeqOfUtf8Bytes("cd")), _dafny.UnicodeSeqOfUtf8Bytes("aprlnhebx"))
	}
}
func (_this *C6) Fm3(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C6) M0(p0 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1108_v0 bool
		_ = _1108_v0
		_1108_v0 = true
		var _1109_v1 _dafny.Int
		_ = _1109_v1
		_1109_v1 = _dafny.IntOfInt64(619)
		var _1110_v2 T1
		_ = _1110_v2
		var _nw142 *C3 = New_C3_()
		_ = _nw142
		_nw142.Ctor__(_1108_v0, _1109_v1)
		_1110_v2 = _nw142
		var _1111_v3 _dafny.Set
		_ = _1111_v3
		_1111_v3 = _dafny.SetOf(_1110_v2, _1110_v2)
		var _1112_v4 _dafny.Set
		_ = _1112_v4
		_1112_v4 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(381))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}((func(_1113_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1114_i0 _dafny.Int) _dafny.CodePoint {
				return _1113_p0
			}
		})(p0)))).Cardinality()), (_1111_v3).Cardinality())
		if (func() bool {
			if _1108_v0 {
				return (_dafny.SetOf(_1109_v1)).IsDisjointFrom(_1112_v4)
			}
			return _1108_v0
		})() {
			var _1115_v5 _dafny.Sequence
			_ = _1115_v5
			_1115_v5 = _dafny.SeqOf(_1108_v0, _1108_v0, _1108_v0)
			var _1116_v6 _dafny.Map
			_ = _1116_v6
			var _1117_v7 _dafny.MultiSet
			_ = _1117_v7
			var _1118_v8 bool
			_ = _1118_v8
			var _1119_v9 _dafny.Array
			_ = _1119_v9
			var _out14 _dafny.Map
			_ = _out14
			var _out15 _dafny.MultiSet
			_ = _out15
			var _out16 bool
			_ = _out16
			var _out17 _dafny.Array
			_ = _out17
			_out14, _out15, _out16, _out17 = (_this).M2(_1115_v5, _1108_v0, globalState)
			_1116_v6 = _out14
			_1117_v7 = _out15
			_1118_v8 = _out16
			_1119_v9 = _out17
			var _1120_v10 _dafny.Array
			_ = _1120_v10
			var _nw143 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw143
			_1120_v10 = _nw143
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))
			_ = _index163
			(_1120_v10).ArraySet1(_1108_v0, (_index163).Int())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))
			_ = _index164
			var _rhs177 bool = (_1109_v1).Cmp(_1109_v1) > 0
			_ = _rhs177
			var _rhs178 _dafny.Int = (_1109_v1).Times(_1109_v1)
			_ = _rhs178
			var _rhs179 bool = Companion_Default___.Fm19(_1109_v1, _1118_v8, globalState)
			_ = _rhs179
			var _lhs121 *GlobalState = globalState
			_ = _lhs121
			var _lhs122 _dafny.Array = _1120_v10
			_ = _lhs122
			var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))
			_ = _lhs123
			_lhs121.F5 = _rhs177
			_1109_v1 = _rhs178
			(_lhs122).ArraySet1(_rhs179, (_lhs123).Int())
			var _1121_v11 _dafny.Sequence
			_ = _1121_v11
			_1121_v11 = _dafny.UnicodeSeqOfUtf8Bytes("mcaxr")
			var _1122_v12 _dafny.Sequence
			_ = _1122_v12
			_1122_v12 = _dafny.SeqOf(_1121_v11)
			var _1123_v13 D0
			_ = _1123_v13
			_1123_v13 = Companion_D0_.Create_DC1_(_1118_v8, _1122_v12, _1121_v11)
			var _1124_v14 D2
			_ = _1124_v14
			_1124_v14 = Companion_D2_.Create_DC6_(_dafny.MultiSetOf(_1123_v13))
			var _source17 D2 = _1124_v14
			_ = _source17
			if _source17.Is_DC7() {
				var _1125___mcc_h0 bool = _source17.Get_().(D2_DC7).Cf11
				_ = _1125___mcc_h0
				var _1126___mcc_h1 _dafny.Map = _source17.Get_().(D2_DC7).Cf12
				_ = _1126___mcc_h1
				var _1127___mcc_h2 _dafny.Int = _source17.Get_().(D2_DC7).Cf13
				_ = _1127___mcc_h2
				var _1128___mcc_h3 bool = _source17.Get_().(D2_DC7).Cf14
				_ = _1128___mcc_h3
				var _1129_cf14 bool = _1128___mcc_h3
				_ = _1129_cf14
				var _1130_cf13 _dafny.Int = _1127___mcc_h2
				_ = _1130_cf13
				var _1131_cf12 _dafny.Map = _1126___mcc_h1
				_ = _1131_cf12
				var _1132_cf11 bool = _1125___mcc_h0
				_ = _1132_cf11
				var _1133_v15 _dafny.Array
				_ = _1133_v15
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_20
				var _nw144 _dafny.Array
				_ = _nw144
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw144 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Sequence = (func(_1134_p0 _dafny.CodePoint, _1135_v11 _dafny.Sequence, _1136_cf14 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_1137_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg90 _dafny.Int) interface{} {
									return coer90(arg90)
								}
							}((func(_1138_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1139_i2 _dafny.Int) _dafny.CodePoint {
									return _1138_p0
								}
							})(_1134_p0))), (Companion_D7_.Create_DC20_(_1135_v11, _1136_cf14)).Dtor_cf31())
						}
					})(p0, _1121_v11, _1129_cf14)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw144 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw144).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw144).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_1133_v15 = _nw144
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1133_v15), 0))
				_ = _index165
				(_1133_v15).ArraySet1(_1121_v11, (_index165).Int())
				var _1140_v16 _dafny.Map
				_ = _1140_v16
				_1140_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_cf11, _1130_cf13)
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1133_v15), 0))
				_ = _index166
				(_1133_v15).ArraySet1(Companion_Default___.Fm20(_1130_cf13, _1108_v0, ((func() _dafny.Int {
					if (_1140_v16).Contains(_1132_cf11) {
						return (_1140_v16).Get(_1132_cf11).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("prnfmcqd")).Cardinality())
				})()).Cmp(_1109_v1) > 0, _1109_v1, globalState), (_index166).Int())
				r1 = _1130_cf13
				var _1141_v17 _dafny.Sequence
				_ = _1141_v17
				_1141_v17 = _dafny.SeqOf(_1115_v5)
				var _1142_v18 _dafny.Sequence
				_ = _1142_v18
				_1142_v18 = _dafny.SeqOf(_1109_v1, _dafny.IntOfUint32((_1141_v17).Cardinality()), _1130_cf13, _1109_v1)
				var _1143_v19 _dafny.Map
				_ = _1143_v19
				_1143_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, _1121_v11)
				_1130_cf13 = ((_this).Fm0(_1109_v1, p0, globalState)).Times((_1142_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1143_v19).Contains(_1130_cf13) {
						return (_1143_v19).Get(_1130_cf13).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("uqbs")
				})()).Cardinality()), _dafny.IntOfUint32((_1142_v18).Cardinality()))).Uint32()).(_dafny.Int))
				_1129_cf14 = !((_1109_v1).Cmp(_dafny.IntOfInt64(939)) != 0)
			} else if _source17.Is_DC6() {
				var _1144___mcc_h4 _dafny.MultiSet = _source17.Get_().(D2_DC6).Cf10
				_ = _1144___mcc_h4
				var _1145_cf10 _dafny.MultiSet = _1144___mcc_h4
				_ = _1145_cf10
				var _1146_v20 *C4
				_ = _1146_v20
				var _nw145 *C4 = New_C4_()
				_ = _nw145
				_nw145.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(459), _1109_v1), Companion_Default___.Fm36(globalState))
				_1146_v20 = _nw145
				var _1147_v21 _dafny.Map
				_ = _1147_v21
				_1147_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1120_v10, _1118_v8)
				_1147_v21 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1120_v10, (_1120_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))).Int()).(bool))).Merge(_1147_v21)).Merge(_1147_v21)
				_1109_v1 = _dafny.IntOfInt64(999)
				var _1148_v22 _dafny.Map
				_ = _1148_v22
				_1148_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v8, !_dafny.Companion_Sequence_.Contains(_1115_v5, _1108_v0))
				_1148_v22 = (_1148_v22).Update(_1118_v8, !(_1108_v0))
			} else {
				var _1149___mcc_h5 D2 = _source17.Get_().(D2_DC8).Cf15
				_ = _1149___mcc_h5
				var _1150_cf15 D2 = _1149___mcc_h5
				_ = _1150_cf15
				r1 = Companion_Default___.SafeModuloInt(_1109_v1, _1109_v1)
				(globalState).F9 = (_1120_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))).Int()).(bool)
				(globalState).F13 = ((_1112_v4).IsSubsetOf(_1112_v4)) == (!(!(_1118_v8) || (_1118_v8)))
				_1118_v8 = ((func() _dafny.Int {
					if _1118_v8 {
						return _dafny.IntOfInt64(850)
					}
					return _1109_v1
				})()).Cmp(_1109_v1) <= 0
			}
			var _1151_v23 *C2
			_ = _1151_v23
			var _nw146 *C2 = New_C2_()
			_ = _nw146
			_nw146.Ctor__()
			_1151_v23 = _nw146
			var _1152_v24 _dafny.MultiSet
			_ = _1152_v24
			_1152_v24 = _dafny.MultiSetOf(_1151_v23)
			(globalState).F9 = !((_1152_v24).IsDisjointFrom((_dafny.MultiSetOf(_1151_v23)).Update(_1151_v23, Companion_Default___.Abs(_dafny.IntOfUint32((_1115_v5).Cardinality())))))
			if (_1120_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))).Int()).(bool) {
				_1109_v1 = _1109_v1
				var _1153_v25 _dafny.MultiSet
				_ = _1153_v25
				_1153_v25 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_(true, _1122_v12, _1121_v11))
				var _1154_v26 *C0
				_ = _1154_v26
				var _nw147 *C0 = New_C0_()
				_ = _nw147
				_nw147.Ctor__((_1153_v25).Intersection(_1153_v25), _1118_v8)
				_1154_v26 = _nw147
				var _1155_v27 _dafny.Map
				_ = _1155_v27
				var _1156_v28 _dafny.MultiSet
				_ = _1156_v28
				var _1157_v29 bool
				_ = _1157_v29
				var _1158_v30 _dafny.Array
				_ = _1158_v30
				var _out18 _dafny.Map
				_ = _out18
				var _out19 _dafny.MultiSet
				_ = _out19
				var _out20 bool
				_ = _out20
				var _out21 _dafny.Array
				_ = _out21
				_out18, _out19, _out20, _out21 = (_this).M2(_1115_v5, !(!(_1108_v0)), globalState)
				_1155_v27 = _out18
				_1156_v28 = _out19
				_1157_v29 = _out20
				_1158_v30 = _out21
				var _1159_v31 D5
				_ = _1159_v31
				_1159_v31 = Companion_D5_.Create_DC13_(_1120_v10)
				var _1160_v32 _dafny.MultiSet
				_ = _1160_v32
				_1160_v32 = _dafny.MultiSetOf(_1159_v31)
				_1108_v0 = (_1160_v32).IsProperSubsetOf(_1160_v32)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))
				_ = _index167
				(_1120_v10).ArraySet1((_1112_v4).IsDisjointFrom(_1112_v4), (_index167).Int())
			} else {
				var _1161_v33 bool
				_ = _1161_v33
				var _1162_v34 _dafny.Int
				_ = _1162_v34
				var _1163_v35 bool
				_ = _1163_v35
				var _out22 bool
				_ = _out22
				var _out23 _dafny.Int
				_ = _out23
				var _out24 bool
				_ = _out24
				_out22, _out23, _out24 = (_1151_v23).M6(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-864))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}(func(_1164_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), _1121_v11, _1118_v8, _1109_v1, globalState)
				_1161_v33 = _out22
				_1162_v34 = _out23
				_1163_v35 = _out24
				var _1165_v36 *C2
				_ = _1165_v36
				var _nw148 *C2 = New_C2_()
				_ = _nw148
				_nw148.Ctor__()
				_1165_v36 = _nw148
				var _1166_v37 _dafny.Map
				_ = _1166_v37
				_1166_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1162_v34, (_1120_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))).Int()).(bool))
				var _1167_v38 _dafny.Sequence
				_ = _1167_v38
				_1167_v38 = _dafny.SeqOf(_dafny.IntOfInt64(-119), _1162_v34)
				var _1168_v39 _dafny.Map
				_ = _1168_v39
				_1168_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v35, (_1167_v38).Select((Companion_Default___.SafeIndex(_1162_v34, _dafny.IntOfUint32((_1167_v38).Cardinality()))).Uint32()).(_dafny.Int))
				var _1169_v40 _dafny.Array
				_ = _1169_v40
				var _nwElement0_28 _dafny.Int = _1109_v1
				_ = _nwElement0_28
				var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(25))
				_ = _nw149
				(_nw149).ArraySet1(_nwElement0_28, 0)
				(_nw149).ArraySet1(_1109_v1, 1)
				(_nw149).ArraySet1(_1162_v34, 2)
				(_nw149).ArraySet1(_1109_v1, 3)
				(_nw149).ArraySet1(_1109_v1, 4)
				(_nw149).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1108_v0, _1161_v33)).Cardinality())), 5)
				(_nw149).ArraySet1((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}((func(_1170_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1171_i4 _dafny.Int) _dafny.CodePoint {
						return _1170_p0
					}
				})(p0))), _1121_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg93 _dafny.Int) interface{} {
						return coer93(arg93)
					}
				}((func(_1172_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1173_i5 _dafny.Int) _dafny.CodePoint {
						return _1172_p0
					}
				})(p0))))).Cardinality(), 6)
				(_nw149).ArraySet1(_1162_v34, 7)
				(_nw149).ArraySet1(_1109_v1, 8)
				(_nw149).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v8, p0)).Cardinality(), 9)
				(_nw149).ArraySet1(_1162_v34, 10)
				(_nw149).ArraySet1(_1109_v1, 11)
				(_nw149).ArraySet1(_1109_v1, 12)
				(_nw149).ArraySet1(_dafny.IntOfUint32((_1121_v11).Cardinality()), 13)
				(_nw149).ArraySet1(_dafny.IntOfInt64(36), 14)
				(_nw149).ArraySet1((_1165_v36).Fm0(_1109_v1, p0, globalState), 15)
				(_nw149).ArraySet1((_1167_v38).Select((Companion_Default___.SafeIndex((_1168_v39).Cardinality(), _dafny.IntOfUint32((_1167_v38).Cardinality()))).Uint32()).(_dafny.Int), 16)
				(_nw149).ArraySet1(_1109_v1, 17)
				(_nw149).ArraySet1(_1162_v34, 18)
				(_nw149).ArraySet1(_1109_v1, 19)
				(_nw149).ArraySet1(_1109_v1, 20)
				(_nw149).ArraySet1(_1109_v1, 21)
				(_nw149).ArraySet1(_1162_v34, 22)
				(_nw149).ArraySet1(_1109_v1, 23)
				(_nw149).ArraySet1(_1109_v1, 24)
				_1169_v40 = _nw149
				var _1174_v41 _dafny.Map
				_ = _1174_v41
				_1174_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1169_v40, (_dafny.Zero).Minus(_1109_v1))
				var _1175_v42 D8
				_ = _1175_v42
				_1175_v42 = Companion_D8_.Create_DC22_(_1174_v41)
				var _1176_v43 _dafny.Set
				_ = _1176_v43
				_1176_v43 = _dafny.SetOf(_1108_v0, _1163_v35, _1108_v0)
				_1166_v37 = (_1166_v37).Update((_1167_v38).Select((Companion_Default___.SafeIndex(((_1175_v42).Dtor_cf34()).Cardinality(), _dafny.IntOfUint32((_1167_v38).Cardinality()))).Uint32()).(_dafny.Int), (_1176_v43).IsProperSubsetOf(_1176_v43))
				_1121_v11 = _dafny.Companion_Sequence_.Update((_1110_v2).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), _1162_v34), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.IntOfUint32(((_1110_v2).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), _1162_v34), globalState)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if !(_1118_v8) {
						return p0
					}
					return p0
				})())
				_1161_v33 = (_1120_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1120_v10), 0))).Int()).(bool)
			}
		} else {
			var _source18 D7 = Companion_D7_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, true))
			_ = _source18
			if _source18.Is_DC19() {
				var _1177___mcc_h6 D1 = _source18.Get_().(D7_DC19).Cf27
				_ = _1177___mcc_h6
				var _1178___mcc_h7 _dafny.Int = _source18.Get_().(D7_DC19).Cf28
				_ = _1178___mcc_h7
				var _1179___mcc_h8 bool = _source18.Get_().(D7_DC19).Cf29
				_ = _1179___mcc_h8
				var _1180___mcc_h9 _dafny.Array = _source18.Get_().(D7_DC19).Cf30
				_ = _1180___mcc_h9
				var _1181_cf30 _dafny.Array = _1180___mcc_h9
				_ = _1181_cf30
				var _1182_cf29 bool = _1179___mcc_h8
				_ = _1182_cf29
				var _1183_cf28 _dafny.Int = _1178___mcc_h7
				_ = _1183_cf28
				var _1184_cf27 D1 = _1177___mcc_h6
				_ = _1184_cf27
				var _1185_v44 _dafny.Sequence
				_ = _1185_v44
				_1185_v44 = _dafny.SeqOf(true)
				var _1186_v45 _dafny.Map
				_ = _1186_v45
				var _1187_v46 _dafny.MultiSet
				_ = _1187_v46
				var _1188_v47 bool
				_ = _1188_v47
				var _1189_v48 _dafny.Array
				_ = _1189_v48
				var _out25 _dafny.Map
				_ = _out25
				var _out26 _dafny.MultiSet
				_ = _out26
				var _out27 bool
				_ = _out27
				var _out28 _dafny.Array
				_ = _out28
				_out25, _out26, _out27, _out28 = (_this).M2(_dafny.Companion_Sequence_.Concatenate(_1185_v44, _1185_v44), (_1183_cf28).Cmp(_1109_v1) < 0, globalState)
				_1186_v45 = _out25
				_1187_v46 = _out26
				_1188_v47 = _out27
				_1189_v48 = _out28
				_1109_v1 = _1183_cf28
				var _1190_v49 _dafny.Map
				_ = _1190_v49
				_1190_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() bool {
					if _1182_cf29 {
						return _1188_v47
					}
					return _1188_v47
				})())
				var _1191_v50 _dafny.Sequence
				_ = _1191_v50
				_1191_v50 = _dafny.SeqOf(_1190_v49)
				var _1192_v51 _dafny.Sequence
				_ = _1192_v51
				_1192_v51 = _dafny.UnicodeSeqOfUtf8Bytes("tvnpxdv")
				var _1193_v52 _dafny.Sequence
				_ = _1193_v52
				_1193_v52 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("anlbx"), _1192_v51)
				var _1194_v53 D0
				_ = _1194_v53
				_1194_v53 = Companion_D0_.Create_DC1_(_1188_v47, _dafny.Companion_Sequence_.Update(_1193_v52, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.IntOfUint32((_1193_v52).Cardinality()))).Uint32(), _1192_v51), _1192_v51)
				var _1195_v54 _dafny.MultiSet
				_ = _1195_v54
				_1195_v54 = _dafny.MultiSetOf(_1194_v53, _1194_v53)
				var _1196_v55 *C0
				_ = _1196_v55
				var _nw150 *C0 = New_C0_()
				_ = _nw150
				_nw150.Ctor__(_1195_v54, _1108_v0)
				_1196_v55 = _nw150
				var _1197_v56 _dafny.Map
				_ = _1197_v56
				_1197_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1196_v55, _1182_cf29)
				var _1198_v57 _dafny.Map
				_ = _1198_v57
				_1198_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1182_cf29), ((_1197_v56).Update(_1196_v55, (_1196_v55).F19())).Cardinality())
				_1190_v49 = (_1190_v49).Merge((_1191_v50).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1198_v57).Contains(_1182_cf29) {
						return (_1198_v57).Get(_1182_cf29).(_dafny.Int)
					}
					return _1183_cf28
				})(), _dafny.IntOfUint32((_1191_v50).Cardinality()))).Uint32()).(_dafny.Map))
				var _1199_v58 T0
				_ = _1199_v58
				var _nw151 *C2 = New_C2_()
				_ = _nw151
				_nw151.Ctor__()
				_1199_v58 = _nw151
				_1199_v58 = _1199_v58
			} else if _source18.Is_DC20() {
				var _1200___mcc_h10 _dafny.Sequence = _source18.Get_().(D7_DC20).Cf31
				_ = _1200___mcc_h10
				var _1201___mcc_h11 bool = _source18.Get_().(D7_DC20).Cf32
				_ = _1201___mcc_h11
				var _1202_cf32 bool = _1201___mcc_h11
				_ = _1202_cf32
				var _1203_cf31 _dafny.Sequence = _1200___mcc_h10
				_ = _1203_cf31
				var _1204_v59 _dafny.Map
				_ = _1204_v59
				_1204_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, _dafny.IntOfInt64(207))
				var _1205_v60 _dafny.Sequence
				_ = _1205_v60
				_1205_v60 = _dafny.SeqOf(_1203_cf31)
				var _1206_v61 D0
				_ = _1206_v61
				_1206_v61 = Companion_D0_.Create_DC1_(_1108_v0, _1205_v60, _1203_cf31)
				var _1207_v62 _dafny.Map
				_ = _1207_v62
				_1207_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19((_dafny.SetOf(_1202_cf32, _1108_v0, _1202_cf32, true)).Cardinality(), _1108_v0, globalState), Companion_D0_.Create_DC3_(_1206_v61))
				var _1208_v63 D2
				_ = _1208_v63
				_1208_v63 = Companion_D2_.Create_DC7_(_1108_v0, _1207_v62, _dafny.IntOfUint32((_1203_cf31).Cardinality()), true)
				var _1209_v64 _dafny.Map
				_ = _1209_v64
				_1209_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, _dafny.CodePoint('t'))
				var _1210_v65 _dafny.MultiSet
				_ = _1210_v65
				_1210_v65 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_v0, _1202_cf32))
				var _1211_v67 _dafny.Sequence
				_ = _1211_v67
				_1211_v67 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1210_v65).Cardinality(), _1202_cf32), func() _dafny.Map {
					var _coll36 = _dafny.NewMapBuilder()
					_ = _coll36
					for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(713), _dafny.IntOfInt64(736))); ; {
						_compr_36, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _1212_v66 _dafny.Int
						_1212_v66 = interface{}(_compr_36).(_dafny.Int)
						if ((_dafny.IntOfInt64(713)).Cmp(_1212_v66) <= 0) && ((_1212_v66).Cmp(_dafny.IntOfInt64(736)) < 0) {
							_coll36.Add(Companion_Default___.SafeModuloInt(_1212_v66, _1109_v1), _1202_cf32)
						}
					}
					return _coll36.ToMap()
				}())
				var _1213_v68 _dafny.Set
				_ = _1213_v68
				_1213_v68 = _dafny.SetOf(_1108_v0)
				var _1214_v69 _dafny.Array
				_ = _1214_v69
				var _nwElement0_29 _dafny.Int = _1109_v1
				_ = _nwElement0_29
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(20))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_29, 0)
				(_nw152).ArraySet1((_1109_v1).Minus(_1109_v1), 1)
				(_nw152).ArraySet1(((_dafny.SetOf(_1202_cf32)).Cardinality()).Minus(((_1204_v59).Update(_1109_v1, _1109_v1)).Cardinality()), 2)
				(_nw152).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1203_cf31).Cardinality()), _1109_v1), 3)
				(_nw152).ArraySet1(_1109_v1, 4)
				(_nw152).ArraySet1(_1109_v1, 5)
				(_nw152).ArraySet1(Companion_Default___.SafeModuloInt(_1109_v1, _1109_v1), 6)
				(_nw152).ArraySet1(_1109_v1, 7)
				(_nw152).ArraySet1((_dafny.Zero).Minus(_1109_v1), 8)
				(_nw152).ArraySet1(Companion_Default___.Fm22(_1208_v63, Companion_D0_.Create_DC3_(_1206_v61), _1202_cf32, globalState), 9)
				(_nw152).ArraySet1((_1109_v1).Minus(_1109_v1), 10)
				(_nw152).ArraySet1(_1109_v1, 11)
				(_nw152).ArraySet1(_1109_v1, 12)
				(_nw152).ArraySet1(_1109_v1, 13)
				(_nw152).ArraySet1(_dafny.IntOfInt64(183), 14)
				(_nw152).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1109_v1, (Companion_Default___.Fm37(!(_1202_cf32), _1109_v1, _1109_v1, (func() _dafny.CodePoint {
					if (_1209_v64).Contains(_1109_v1) {
						return (_1209_v64).Get(_1109_v1).(_dafny.CodePoint)
					}
					return p0
				})(), globalState)).Cardinality())), 15)
				(_nw152).ArraySet1(_1109_v1, 16)
				(_nw152).ArraySet1((_dafny.IntOfUint32((_1211_v67).Cardinality())).Minus(_dafny.IntOfUint32((_1203_cf31).Cardinality())), 17)
				(_nw152).ArraySet1((_1213_v68).Cardinality(), 18)
				(_nw152).ArraySet1((_dafny.IntOfInt64(342)).Plus(_dafny.IntOfInt64(138)), 19)
				_1214_v69 = _nw152
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1214_v69), 0))
				_ = _index168
				(_1214_v69).ArraySet1((_1204_v59).Cardinality(), (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1214_v69), 0))
				_ = _index169
				(_1214_v69).ArraySet1(_dafny.IntOfInt64(674), (_index169).Int())
				var _1215_v70 _dafny.Array
				_ = _1215_v70
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw153
				_1215_v70 = _nw153
				_1215_v70 = _1215_v70
				_1109_v1 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1202_cf32, Companion_Default___.Fm6(globalState), _1202_cf32, !(_1108_v0), _1108_v0)).Cardinality()))
				r1 = _1109_v1
			} else if _source18.Is_DC18() {
				var _1216___mcc_h12 _dafny.Map = _source18.Get_().(D7_DC18).Cf26
				_ = _1216___mcc_h12
				var _1217_cf26 _dafny.Map = _1216___mcc_h12
				_ = _1217_cf26
				var _1218_v71 _dafny.Array
				_ = _1218_v71
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_21
				var _nw154 _dafny.Array
				_ = _nw154
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw154 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Set = (func(_1219_v4 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1220_i6 _dafny.Int) _dafny.Set {
							return (func() _dafny.Set {
								if true {
									return _1219_v4
								}
								return _1219_v4
							})()
						}
					})(_1112_v4)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw154 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw154).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw154).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_1218_v71 = _nw154
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_1218_v71), 0))
				_ = _index170
				(_1218_v71).ArraySet1((func() _dafny.Set {
					if _1108_v0 {
						return _1112_v4
					}
					return _1112_v4
				})(), (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_1218_v71), 0))
				_ = _index171
				(_1218_v71).ArraySet1(_1112_v4, (_index171).Int())
				var _1221_v72 _dafny.Array
				_ = _1221_v72
				var _len0_22 _dafny.Int = _dafny.One
				_ = _len0_22
				var _nw155 _dafny.Array
				_ = _nw155
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw155 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Int = (func(_1222_v0 bool) func(_dafny.Int) _dafny.Int {
						return func(_1223_i7 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1223_i7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1222_v0, _1222_v0)).Cardinality())
						}
					})(_1108_v0)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw155 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw155).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw155).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1221_v72 = _nw155
				var _1224_v73 _dafny.Sequence
				_ = _1224_v73
				_1224_v73 = _dafny.UnicodeSeqOfUtf8Bytes("taqaliqve")
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1221_v72), 0))
				_ = _index172
				(_1221_v72).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1224_v73).Cardinality()), _1109_v1)), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1221_v72), 0))
				_ = _index173
				(_1221_v72).ArraySet1(_1109_v1, (_index173).Int())
				var _1225_v74 _dafny.Array
				_ = _1225_v74
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
				_ = _nw156
				_1225_v74 = _nw156
				var _1226_v75 _dafny.Map
				_ = _1226_v75
				_1226_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1221_v72, _dafny.IntOfUint32((_1224_v73).Cardinality()))
				var _1227_v76 D8
				_ = _1227_v76
				_1227_v76 = Companion_D8_.Create_DC22_((_1226_v75).Update(_1221_v72, (_1221_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1221_v72), 0))).Int()).(_dafny.Int)))
				var _1228_v77 _dafny.Array
				_ = _1228_v77
				var _nwElement0_30 bool = _1108_v0
				_ = _nwElement0_30
				var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(3))
				_ = _nw157
				(_nw157).ArraySet1(_nwElement0_30, 0)
				(_nw157).ArraySet1(_1108_v0, 1)
				(_nw157).ArraySet1(_1108_v0, 2)
				_1228_v77 = _nw157
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1228_v77), 0))
				_ = _index174
				(_1228_v77).ArraySet1(_1108_v0, (_index174).Int())
				var _1229_v78 _dafny.Set
				_ = _1229_v78
				_1229_v78 = _dafny.SetOf(_1108_v0, !(_1108_v0), _1108_v0, _1108_v0, _1108_v0)
				var _pat_let_tv23 = _1226_v75
				_ = _pat_let_tv23
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1221_v72), 0))
				_ = _index175
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1228_v77), 0))
				_ = _index176
				var _rhs180 _dafny.Array = _1225_v74
				_ = _rhs180
				var _rhs181 bool = !(_1108_v0)
				_ = _rhs181
				var _rhs182 _dafny.Int = ((_this).Fm0(_1109_v1, p0, globalState)).Minus(_1109_v1)
				_ = _rhs182
				var _rhs183 D8 = func(_pat_let23_0 D8) D8 {
					return func(_1230_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let24_0 _dafny.Map) D8 {
							return func(_1231_dt__update_hcf34_h0 _dafny.Map) D8 {
								return Companion_D8_.Create_DC22_(_1231_dt__update_hcf34_h0)
							}(_pat_let24_0)
						}(_pat_let_tv23)
					}(_pat_let23_0)
				}(_1227_v76)
				_ = _rhs183
				var _rhs184 bool = (Companion_Default___.Fm33(_1108_v0, _1108_v0, _1108_v0, globalState)).IsDisjointFrom((_1229_v78).Intersection(_1229_v78))
				_ = _rhs184
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				var _lhs125 _dafny.Array = _1221_v72
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1221_v72), 0))
				_ = _lhs126
				var _lhs127 _dafny.Array = _1228_v77
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1228_v77), 0))
				_ = _lhs128
				_1225_v74 = _rhs180
				_lhs124.F3 = _rhs181
				(_lhs125).ArraySet1(_rhs182, (_lhs126).Int())
				_1227_v76 = _rhs183
				(_lhs127).ArraySet1(_rhs184, (_lhs128).Int())
				var _1232_v79 D5
				_ = _1232_v79
				_1232_v79 = Companion_D5_.Create_DC13_(_1228_v77)
				_1228_v77 = (_1232_v79).Dtor_cf19()
			} else {
				var _1233___mcc_h13 D7 = _source18.Get_().(D7_DC21).Cf33
				_ = _1233___mcc_h13
				var _1234_cf33 D7 = _1233___mcc_h13
				_ = _1234_cf33
				var _1235_v80 _dafny.Sequence
				_ = _1235_v80
				_1235_v80 = _dafny.UnicodeSeqOfUtf8Bytes("pdxrtkj")
				r1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1235_v80).Cardinality()), Companion_Default___.SafeModuloInt(_1109_v1, _1109_v1))
				_1109_v1 = (_dafny.Zero).Minus(_1109_v1)
				var _1236_v81 _dafny.Array
				_ = _1236_v81
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw158
				_1236_v81 = _nw158
				var _1237_v82 _dafny.Array
				_ = _1237_v82
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_23
				var _nw159 _dafny.Array
				_ = _nw159
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw159 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) bool = (func(_1238_v0 bool) func(_dafny.Int) bool {
						return func(_1239_i8 _dafny.Int) bool {
							return _1238_v0
						}
					})(_1108_v0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw159 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw159).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw159).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1237_v82 = _nw159
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1236_v81), 0))
				_ = _index177
				(_1236_v81).ArraySet1(_1237_v82, (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1236_v81), 0))
				_ = _index178
				(_1236_v81).ArraySet1(_1237_v82, (_index178).Int())
				var _1240_v83 D4
				_ = _1240_v83
				_1240_v83 = Companion_D4_.Create_DC12_(_1109_v1)
				var _1241_v84 _dafny.Map
				_ = _1241_v84
				_1241_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1240_v83, (func() _dafny.Int {
					if _1108_v0 {
						return (_dafny.Zero).Minus((_1110_v2).Fm0(_1109_v1, p0, globalState))
					}
					return _1109_v1
				})())
				r1 = (func() _dafny.Int {
					if (_1241_v84).Contains(Companion_Default___.Fm35(_dafny.SeqOf(true), _1108_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg94 _dafny.Int) interface{} {
							return coer94(arg94)
						}
					}((func(_1242_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1243_i9 _dafny.Int) _dafny.CodePoint {
							return _1242_p0
						}
					})(p0))), globalState)) {
						return (_1241_v84).Get(Companion_Default___.Fm35(_dafny.SeqOf(true), _1108_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg95 _dafny.Int) interface{} {
								return coer95(arg95)
							}
						}((func(_1244_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1245_i9 _dafny.Int) _dafny.CodePoint {
								return _1244_p0
							}
						})(p0))), globalState)).(_dafny.Int)
					}
					return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(279), _1109_v1)
				})()
			}
			var _1246_v85 _dafny.Array
			_ = _1246_v85
			var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw160
			_1246_v85 = _nw160
			_1246_v85 = _1246_v85
			var _1247_v86 _dafny.Map
			_ = _1247_v86
			_1247_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, (false) == (_1108_v0))
			var _1248_v87 _dafny.Sequence
			_ = _1248_v87
			_1248_v87 = _dafny.UnicodeSeqOfUtf8Bytes("lgyduwq")
			var _rhs185 _dafny.Map = _1247_v86
			_ = _rhs185
			var _rhs186 _dafny.Int = _dafny.IntOfUint32((_1248_v87).Cardinality())
			_ = _rhs186
			var _rhs187 _dafny.Array = _1246_v85
			_ = _rhs187
			_1247_v86 = _rhs185
			r1 = _rhs186
			_1246_v85 = _rhs187
			var _1249_v88 _dafny.Map
			_ = _1249_v88
			_1249_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1109_v1).Cmp(_1109_v1) <= 0, _1108_v0)
			_1249_v88 = (_1249_v88).Update(_1108_v0, false)
			if _1108_v0 {
				var _1250_v89 _dafny.Sequence
				_ = _1250_v89
				_1250_v89 = _dafny.SeqOf(_1108_v0, _1108_v0, _1108_v0, _1108_v0, true)
				r1 = _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1250_v89).Cardinality())))).Cardinality())
				var _1251_v90 D1
				_ = _1251_v90
				_1251_v90 = Companion_D1_.Create_DC5_(_1108_v0, _1248_v87, _1250_v89)
				var _1252_v91 _dafny.Set
				_ = _1252_v91
				_1252_v91 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(760), _1251_v90))
				var _1253_v92 _dafny.Map
				_ = _1253_v92
				_1253_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v1, _1251_v90)
				_1252_v91 = (_1252_v91).Union(_dafny.SetOf((_1253_v92).Update(_1109_v1, _1251_v90)))
				var _1254_v93 *C3
				_ = _1254_v93
				var _nw161 *C3 = New_C3_()
				_ = _nw161
				_nw161.Ctor__((_dafny.SetOf(_dafny.IntOfUint32((_1248_v87).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kenqbixrf")).Cardinality()))).IsProperSubsetOf(_1112_v4), _1109_v1)
				_1254_v93 = _nw161
				_1250_v89 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1250_v89, _1250_v89), _1250_v89)
				var _1255_v94 _dafny.Sequence
				_ = _1255_v94
				_1255_v94 = _dafny.SeqOf(_1109_v1)
				var _1256_v95 _dafny.Array
				_ = _1256_v95
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(29))
				_ = _nw162
				_1256_v95 = _nw162
				var _1257_v96 _dafny.Set
				_ = _1257_v96
				_1257_v96 = _dafny.SetOf((Companion_D5_.Create_DC14_((_1254_v93).Fm13((_1254_v93).F17(), _1255_v94, (_1254_v93).F17(), _1108_v0, globalState), _1108_v0, _1256_v95)).Dtor_cf20(), Companion_Default___.Fm19((_1254_v93).F17(), (_1254_v93).F16(), globalState))
				var _1258_v97 _dafny.Sequence
				_ = _1258_v97
				_1258_v97 = _dafny.SeqOf(_1257_v96, _1257_v96, _1257_v96, _1257_v96)
				_1257_v96 = (func() _dafny.Set {
					if ((_1254_v93).F17()).Cmp(_dafny.IntOfInt64(-723)) != 0 {
						return (_1257_v96).Union(_1257_v96)
					}
					return ((_1258_v97).Select((Companion_Default___.SafeIndex(_1109_v1, _dafny.IntOfUint32((_1258_v97).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_dafny.SetOf((_1254_v93).F16()))
				})()
			} else {
				(globalState).F13 = _1108_v0
				r1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1109_v1))
				var _1259_v98 _dafny.Sequence
				_ = _1259_v98
				_1259_v98 = _dafny.SeqOf(_1108_v0, _1108_v0)
				r0 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1259_v98, _1259_v98), _dafny.Companion_Sequence_.Concatenate(_1259_v98, _1259_v98))
				r1 = _1109_v1
				var _rhs188 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("jmukv")
				_ = _rhs188
				var _rhs189 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1248_v87, (Companion_Default___.SafeIndex(_1109_v1, _dafny.IntOfUint32((_1248_v87).Cardinality()))).Uint32(), _dafny.CodePoint('c')), (Companion_Default___.SafeIndex(((_dafny.Zero).Minus(_1109_v1)).Times(_1109_v1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1248_v87, (Companion_Default___.SafeIndex(_1109_v1, _dafny.IntOfUint32((_1248_v87).Cardinality()))).Uint32(), _dafny.CodePoint('c'))).Cardinality()))).Uint32(), p0)
				_ = _rhs189
				_1248_v87 = _rhs188
				_1248_v87 = _rhs189
			}
		}
		var _1260_v99 D0
		_ = _1260_v99
		_1260_v99 = Companion_D0_.Create_DC0_(_1109_v1)
		var _1261_v100 D0
		_ = _1261_v100
		_1261_v100 = Companion_D0_.Create_DC3_(_1260_v99)
		var _1262_v101 _dafny.Map
		_ = _1262_v101
		_1262_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_v0, _1261_v100)
		var _1263_v102 D7
		_ = _1263_v102
		_1263_v102 = Companion_D7_.Create_DC20_(_dafny.UnicodeSeqOfUtf8Bytes("khrd"), _1108_v0)
		var _1264_v103 D2
		_ = _1264_v103
		_1264_v103 = Companion_D2_.Create_DC7_(_1108_v0, _1262_v101, _1109_v1, (_1263_v102).Dtor_cf32())
		var _1265_v104 _dafny.Array
		_ = _1265_v104
		var _nwElement0_31 bool = _1108_v0
		_ = _nwElement0_31
		var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(19))
		_ = _nw163
		(_nw163).ArraySet1(_nwElement0_31, 0)
		(_nw163).ArraySet1(!(_1108_v0), 1)
		(_nw163).ArraySet1(_1108_v0, 2)
		(_nw163).ArraySet1(_1108_v0, 3)
		(_nw163).ArraySet1(false, 4)
		(_nw163).ArraySet1(_1108_v0, 5)
		(_nw163).ArraySet1(_1108_v0, 6)
		(_nw163).ArraySet1(_1108_v0, 7)
		(_nw163).ArraySet1(!(_1108_v0), 8)
		(_nw163).ArraySet1(false, 9)
		(_nw163).ArraySet1(_1108_v0, 10)
		(_nw163).ArraySet1(_1108_v0, 11)
		(_nw163).ArraySet1((_this).Fm3(true, _1108_v0, _dafny.IntOfInt64(248), globalState), 12)
		(_nw163).ArraySet1(_1108_v0, 13)
		(_nw163).ArraySet1(_1108_v0, 14)
		(_nw163).ArraySet1((_1264_v103).Dtor_cf11(), 15)
		(_nw163).ArraySet1(_1108_v0, 16)
		(_nw163).ArraySet1(_1108_v0, 17)
		(_nw163).ArraySet1(_1108_v0, 18)
		_1265_v104 = _nw163
		var _1266_v105 _dafny.Array
		_ = _1266_v105
		var _nwElement0_32 _dafny.Array = _1265_v104
		_ = _nwElement0_32
		var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(9))
		_ = _nw164
		(_nw164).ArraySet1(_nwElement0_32, 0)
		(_nw164).ArraySet1(_1265_v104, 1)
		(_nw164).ArraySet1(_1265_v104, 2)
		(_nw164).ArraySet1(_1265_v104, 3)
		(_nw164).ArraySet1(_1265_v104, 4)
		(_nw164).ArraySet1(_1265_v104, 5)
		(_nw164).ArraySet1(_1265_v104, 6)
		(_nw164).ArraySet1(_1265_v104, 7)
		(_nw164).ArraySet1(_1265_v104, 8)
		_1266_v105 = _nw164
		var _1267_v106 _dafny.Sequence
		_ = _1267_v106
		_1267_v106 = _dafny.SeqOf(_1265_v104, _1265_v104)
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1266_v105), 0))
		_ = _index179
		(_1266_v105).ArraySet1((_1267_v106).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-808), _dafny.IntOfUint32((_1267_v106).Cardinality()))).Uint32()).(_dafny.Array), (_index179).Int())
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1266_v105), 0))
		_ = _index180
		(_1266_v105).ArraySet1(_1265_v104, (_index180).Int())
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
		_ = _index181
		(_1265_v104).ArraySet1(_1108_v0, (_index181).Int())
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
		_ = _index182
		(_1265_v104).ArraySet1(!(_1108_v0) || ((func() bool {
			if _1108_v0 {
				return !(_1108_v0)
			}
			return _1108_v0
		})()), (_index182).Int())
		var _1268_v107 *C3
		_ = _1268_v107
		var _nw165 *C3 = New_C3_()
		_ = _nw165
		_nw165.Ctor__((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), _1109_v1)
		_1268_v107 = _nw165
		var _1269_i10 _dafny.Int
		_ = _1269_i10
		_1269_i10 = _dafny.Zero
		{
			for (_1268_v107).F16() {
				{
					if (_1269_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1269_i10 = (_1269_i10).Plus(_dafny.One)
					var _1270_v108 _dafny.Sequence
					_ = _1270_v108
					_1270_v108 = _dafny.UnicodeSeqOfUtf8Bytes("pjglis")
					_1270_v108 = (func() _dafny.Sequence {
						if _1108_v0 {
							return _1270_v108
						}
						return _1270_v108
					})()
					var _1271_v109 _dafny.Sequence
					_ = _1271_v109
					_1271_v109 = _dafny.SeqOf(_1270_v108)
					var _1272_v110 _dafny.Sequence
					_ = _1272_v110
					_1272_v110 = _dafny.SeqOf(Companion_D0_.Create_DC1_((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), _1271_v109, _1270_v108))
					var _1273_v111 D0
					_ = _1273_v111
					_1273_v111 = Companion_D0_.Create_DC1_((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), _1271_v109, _1270_v108)
					var _1274_v112 _dafny.MultiSet
					_ = _1274_v112
					_1274_v112 = _dafny.MultiSetOf(_1273_v111, _1273_v111, _1273_v111)
					var _pat_let_tv24 = _1274_v112
					_ = _pat_let_tv24
					var _source19 D2 = func(_pat_let25_0 D2) D2 {
						return func(_1275_dt__update__tmp_h1 D2) D2 {
							return func(_pat_let26_0 D2) D2 {
								return func(_1276_dt__update_hcf15_h0 D2) D2 {
									return Companion_D2_.Create_DC8_(_1276_dt__update_hcf15_h0)
								}(_pat_let26_0)
							}(Companion_D2_.Create_DC6_(_pat_let_tv24))
						}(_pat_let25_0)
					}(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC6_(_dafny.MultiSetFromSeq(_1272_v110))))
					_ = _source19
					if _source19.Is_DC7() {
						var _1277___mcc_h14 bool = _source19.Get_().(D2_DC7).Cf11
						_ = _1277___mcc_h14
						var _1278___mcc_h15 _dafny.Map = _source19.Get_().(D2_DC7).Cf12
						_ = _1278___mcc_h15
						var _1279___mcc_h16 _dafny.Int = _source19.Get_().(D2_DC7).Cf13
						_ = _1279___mcc_h16
						var _1280___mcc_h17 bool = _source19.Get_().(D2_DC7).Cf14
						_ = _1280___mcc_h17
						var _1281_cf14 bool = _1280___mcc_h17
						_ = _1281_cf14
						var _1282_cf13 _dafny.Int = _1279___mcc_h16
						_ = _1282_cf13
						var _1283_cf12 _dafny.Map = _1278___mcc_h15
						_ = _1283_cf12
						var _1284_cf11 bool = _1277___mcc_h14
						_ = _1284_cf11
						var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _index183
						(_1265_v104).ArraySet1(true, (_index183).Int())
						var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _index184
						(_1265_v104).ArraySet1(_1281_cf14, (_index184).Int())
						var _1285_v113 _dafny.Map
						_ = _1285_v113
						_1285_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1268_v107).F17())
						var _1286_v114 D9
						_ = _1286_v114
						_1286_v114 = Companion_D9_.Create_DC25_(_1285_v113)
						var _1287_v115 *C4
						_ = _1287_v115
						var _nw166 *C4 = New_C4_()
						_ = _nw166
						_nw166.Ctor__(_1109_v1, (_1286_v114).Dtor_cf38())
						_1287_v115 = _nw166
						var _1288_v116 _dafny.Sequence
						_ = _1288_v116
						_1288_v116 = _dafny.SeqOf(_1281_cf14, true)
						var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _index185
						(_1265_v104).ArraySet1((_1288_v116).Select((Companion_Default___.SafeIndex((_1268_v107).F17(), _dafny.IntOfUint32((_1288_v116).Cardinality()))).Uint32()).(bool), (_index185).Int())
					} else if _source19.Is_DC6() {
						var _1289___mcc_h18 _dafny.MultiSet = _source19.Get_().(D2_DC6).Cf10
						_ = _1289___mcc_h18
						var _1290_cf10 _dafny.MultiSet = _1289___mcc_h18
						_ = _1290_cf10
						var _1291_v117 D9
						_ = _1291_v117
						_1291_v117 = Companion_D9_.Create_DC27_((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), (_1268_v107).F17(), (_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool))
						var _1292_v118 _dafny.MultiSet
						_ = _1292_v118
						_1292_v118 = _dafny.MultiSetOf((_1268_v107).F16())
						var _rhs190 _dafny.Int = ((func() _dafny.MultiSet {
							if (_1268_v107).F16() {
								return _1292_v118
							}
							return _1292_v118
						})()).Cardinality()
						_ = _rhs190
						var _rhs191 D9 = _1291_v117
						_ = _rhs191
						_1109_v1 = _rhs190
						_1291_v117 = _rhs191
						var _1293_v119 _dafny.MultiSet
						_ = _1293_v119
						_1293_v119 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm30(globalState)).Cardinality()))).Cardinality()))
						var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _index186
						(_1265_v104).ArraySet1((_1293_v119).IsSubsetOf(_dafny.MultiSetOf((_1268_v107).F17())), (_index186).Int())
						_1270_v108 = _dafny.Companion_Sequence_.Concatenate(_1270_v108, _dafny.Companion_Sequence_.Concatenate(_1270_v108, _1270_v108))
						var _1294_v120 _dafny.Sequence
						_ = _1294_v120
						_1294_v120 = _dafny.SeqOf(_1268_v107)
						var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _index187
						var _rhs192 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1294_v120, _1294_v120)
						_ = _rhs192
						var _rhs193 _dafny.Int = (_1268_v107).F17()
						_ = _rhs193
						var _lhs129 _dafny.Array = _1265_v104
						_ = _lhs129
						var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))
						_ = _lhs130
						(_lhs129).ArraySet1(_rhs192, (_lhs130).Int())
						r1 = _rhs193
					} else {
						var _1295___mcc_h19 D2 = _source19.Get_().(D2_DC8).Cf15
						_ = _1295___mcc_h19
						var _1296_cf15 D2 = _1295___mcc_h19
						_ = _1296_cf15
						var _1297_v121 D5
						_ = _1297_v121
						_1297_v121 = Companion_D5_.Create_DC13_(_dafny.ArrayCastTo((_1266_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1266_v105), 0))).Int())))
						var _1298_v122 D5
						_ = _1298_v122
						_1298_v122 = Companion_D5_.Create_DC13_((_1297_v121).Dtor_cf19())
						var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1266_v105), 0))
						_ = _index188
						(_1266_v105).ArraySet1((_1297_v121).Dtor_cf19(), (_index188).Int())
						(globalState).F5 = ((_1268_v107).F17()).Cmp((_1268_v107).F17()) <= 0
						var _1299_v123 _dafny.Map
						_ = _1299_v123
						_1299_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1268_v107).F17(), (_1268_v107).F17())
						r1 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1109_v1), _dafny.IntOfUint32(((_this).Fm1(_1299_v123, globalState)).Cardinality()))).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_1270_v108).Cardinality())))
						var _1300_v124 _dafny.Array
						_ = _1300_v124
						_1300_v124 = _1266_v105
						_1266_v105 = (_1300_v124)
					}
					var _1301_v125 _dafny.Set
					_ = _1301_v125
					_1301_v125 = _dafny.SetOf((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), Companion_Default___.Fm19((_1268_v107).F17(), _1108_v0, globalState))
					_1301_v125 = _1301_v125
					var _1302_v126 _dafny.Array
					_ = _1302_v126
					var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
					_ = _nw167
					_1302_v126 = _nw167
					var _1303_v127 _dafny.Map
					_ = _1303_v127
					_1303_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1302_v126, (_1268_v107).F17())
					var _1304_v128 D8
					_ = _1304_v128
					_1304_v128 = Companion_D8_.Create_DC22_((_1303_v127).Update(_1302_v126, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg96 _dafny.Int) interface{} {
							return coer96(arg96)
						}
					}(func(_1305_i11 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					}))).Cardinality())))
					_1304_v128 = _1304_v128
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1306_v129 _dafny.Array
		_ = _1306_v129
		var _nw168 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
		_ = _nw168
		_1306_v129 = _nw168
		var _1307_v130 T0
		_ = _1307_v130
		var _nw169 *C2 = New_C2_()
		_ = _nw169
		_nw169.Ctor__()
		_1307_v130 = _nw169
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1306_v129), 0))
		_ = _index189
		(_1306_v129).ArraySet1(_1307_v130, (_index189).Int())
		var _1308_v131 _dafny.Array
		_ = _1308_v131
		var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
		_ = _nw170
		_1308_v131 = _nw170
		var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1308_v131), 0))
		_ = _index190
		(_1308_v131).ArraySet1(_1112_v4, (_index190).Int())
		var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1306_v129), 0))
		_ = _index191
		var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1308_v131), 0))
		_ = _index192
		var _rhs194 T0 = _1307_v130
		_ = _rhs194
		var _rhs195 _dafny.Set = _1112_v4
		_ = _rhs195
		var _lhs131 _dafny.Array = _1306_v129
		_ = _lhs131
		var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1306_v129), 0))
		_ = _lhs132
		var _lhs133 _dafny.Array = _1308_v131
		_ = _lhs133
		var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1308_v131), 0))
		_ = _lhs134
		(_lhs131).ArraySet1(_rhs194, (_lhs132).Int())
		(_lhs133).ArraySet1(_rhs195, (_lhs134).Int())
		var _1309_v132 _dafny.Sequence
		_ = _1309_v132
		_1309_v132 = _dafny.SeqOf(!(Companion_Default___.Fm6(globalState)), (_1268_v107).F16(), (func() bool {
			if (_1268_v107).F16() {
				return (_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool)
			}
			return _1108_v0
		})())
		r0 = (_1309_v132).Select((Companion_Default___.SafeIndex((_1268_v107).F17(), _dafny.IntOfUint32((_1309_v132).Cardinality()))).Uint32()).(bool)
		var _1310_v133 _dafny.Map
		_ = _1310_v133
		_1310_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1309_v132)
		r1 = ((func() _dafny.Map {
			if (_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool) {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1265_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1265_v104), 0))).Int()).(bool), _1309_v132)
			}
			return _1310_v133
		})()).Cardinality()
		return r0, r1
	}
}
func (_this *C6) M2(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Map, _dafny.MultiSet, bool, _dafny.Array) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r3
		var _1311_v0 _dafny.Sequence
		_ = _1311_v0
		_1311_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ypgvmngw")
		var _1312_v1 _dafny.MultiSet
		_ = _1312_v1
		_1312_v1 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1311_v0).Cardinality()))
		var _1313_v2 _dafny.Map
		_ = _1313_v2
		_1313_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1312_v1)
		var _1314_v3 _dafny.Map
		_ = _1314_v3
		_1314_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1313_v2, _1312_v1)
		var _1315_v4 _dafny.Int
		_ = _1315_v4
		_1315_v4 = _dafny.IntOfInt64(-81)
		var _1316_v5 _dafny.Sequence
		_ = _1316_v5
		_1316_v5 = _dafny.SeqOf(_1315_v4)
		_1314_v3 = (_1314_v3).Update(_1313_v2, _dafny.MultiSetFromSeq((func() _dafny.Sequence {
			if p1 {
				return _1316_v5
			}
			return _1316_v5
		})()))
		_1315_v4 = Companion_Default___.Fm11((p1) && (true), globalState)
		var _1317_v6 _dafny.Array
		_ = _1317_v6
		var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
		_ = _nw171
		_1317_v6 = _nw171
		var _1318_v7 _dafny.Map
		_ = _1318_v7
		_1318_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg97 _dafny.Int) interface{} {
				return coer97(arg97)
			}
		}(func(_1319_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})))
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1317_v6), 0))
		_ = _index193
		(_1317_v6).ArraySet1((_1318_v7).Merge(_1318_v7), (_index193).Int())
		var _1320_v8 _dafny.CodePoint
		_ = _1320_v8
		_1320_v8 = _dafny.CodePoint('l')
		var _1321_v9 _dafny.Map
		_ = _1321_v9
		_1321_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1316_v5, _1315_v4)
		var _1322_v10 _dafny.Map
		_ = _1322_v10
		_1322_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1315_v4, (_1312_v1).Cardinality())
		var _1323_v11 _dafny.Map
		_ = _1323_v11
		_1323_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1322_v10).Cardinality(), _1315_v4)
		var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1317_v6), 0))
		_ = _index194
		var _rhs196 _dafny.Int = (_dafny.Zero).Minus(_1315_v4)
		_ = _rhs196
		var _rhs197 _dafny.Map = (_1318_v7).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1311_v0))
		_ = _rhs197
		var _rhs198 _dafny.CodePoint = _1320_v8
		_ = _rhs198
		var _rhs199 _dafny.Map = _1321_v9
		_ = _rhs199
		var _rhs200 _dafny.Sequence = (_this).Fm1(_1323_v11, globalState)
		_ = _rhs200
		var _lhs135 _dafny.Array = _1317_v6
		_ = _lhs135
		var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1317_v6), 0))
		_ = _lhs136
		_1315_v4 = _rhs196
		(_lhs135).ArraySet1(_rhs197, (_lhs136).Int())
		_1320_v8 = _rhs198
		_1321_v9 = _rhs199
		_1311_v0 = _rhs200
		_1320_v8 = _1320_v8
		var _1324_v12 _dafny.Set
		_ = _1324_v12
		_1324_v12 = _dafny.SetOf(_1315_v4, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cdtw")).Cardinality()))
		_1324_v12 = _1324_v12
		if p1 {
			_1311_v0 = _dafny.Companion_Sequence_.Concatenate(_1311_v0, _1311_v0)
			var _1325_v13 _dafny.Map
			_ = _1325_v13
			_1325_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1315_v4, p1)
			_1325_v13 = ((_1325_v13).Merge(_1325_v13)).Merge(_1325_v13)
			_1311_v0 = _1311_v0
			var _1326_v14 _dafny.Sequence
			_ = _1326_v14
			_1326_v14 = _dafny.SeqOf(_1316_v5, _1316_v5)
			var _1327_v15 _dafny.Sequence
			_ = _1327_v15
			_1327_v15 = _dafny.SeqOf(_dafny.SeqOf(_1316_v5), _dafny.Companion_Sequence_.Concatenate(_1326_v14, _1326_v14), Companion_Default___.Fm38(globalState))
			var _1328_v16 D0
			_ = _1328_v16
			_1328_v16 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((p0).Cardinality()))
			var _1329_v17 D0
			_ = _1329_v17
			_1329_v17 = Companion_D0_.Create_DC3_(_1328_v16)
			var _1330_v18 _dafny.Map
			_ = _1330_v18
			_1330_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1329_v17)
			_1326_v14 = (_1327_v15).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm22(Companion_D2_.Create_DC7_(p1, _1330_v18, _1315_v4, p1), _1329_v17, p1, globalState), _dafny.IntOfUint32((_1327_v15).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _1331_v19 _dafny.Map
			_ = _1331_v19
			_1331_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
			var _1332_v20 _dafny.Array
			_ = _1332_v20
			var _nwElement0_33 bool = p1
			_ = _nwElement0_33
			var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(17))
			_ = _nw172
			(_nw172).ArraySet1(_nwElement0_33, 0)
			(_nw172).ArraySet1((func() bool {
				if (_1331_v19).Contains(p1) {
					return (_1331_v19).Get(p1).(bool)
				}
				return p1
			})(), 1)
			(_nw172).ArraySet1(p1, 2)
			(_nw172).ArraySet1(p1, 3)
			(_nw172).ArraySet1(true, 4)
			(_nw172).ArraySet1(true, 5)
			(_nw172).ArraySet1(!(p1), 6)
			(_nw172).ArraySet1(true, 7)
			(_nw172).ArraySet1((_1324_v12).IsProperSubsetOf(_1324_v12), 8)
			(_nw172).ArraySet1(p1, 9)
			(_nw172).ArraySet1(p1, 10)
			(_nw172).ArraySet1(!((p1) == (p1)), 11)
			(_nw172).ArraySet1(p1, 12)
			(_nw172).ArraySet1(!(Companion_Default___.Fm19(_1315_v4, p1, globalState)) || (p1), 13)
			(_nw172).ArraySet1(p1, 14)
			(_nw172).ArraySet1(((_dafny.Zero).Minus(_1315_v4)).Cmp(_1315_v4) < 0, 15)
			(_nw172).ArraySet1(!(p1), 16)
			_1332_v20 = _nw172
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1332_v20), 0))
			_ = _index195
			(_1332_v20).ArraySet1(false, (_index195).Int())
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1332_v20), 0))
			_ = _index196
			(_1332_v20).ArraySet1(p1, (_index196).Int())
		} else {
			var _1333_v21 _dafny.Map
			_ = _1333_v21
			_1333_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1320_v8, _1315_v4)
			var _1334_v22 *C4
			_ = _1334_v22
			var _nw173 *C4 = New_C4_()
			_ = _nw173
			_nw173.Ctor__((_1315_v4).Times(_1315_v4), _1333_v21)
			_1334_v22 = _nw173
			var _1335_v23 _dafny.Map
			_ = _1335_v23
			_1335_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_1334_v22).F14())), p0)
			(globalState).F5 = !_dafny.Companion_Sequence_.Equal(p0, (func() _dafny.Sequence {
				if (_1335_v23).Contains((_1334_v22).F14()) {
					return (_1335_v23).Get((_1334_v22).F14()).(_dafny.Sequence)
				}
				return _dafny.SeqOf(p1)
			})())
			(globalState).F13 = p1
			var _1336_v24 D4
			_ = _1336_v24
			_1336_v24 = Companion_D4_.Create_DC11_()
			var _source20 D4 = _1336_v24
			_ = _source20
			if _source20.Is_DC11() {
				var _1337_v25 D9
				_ = _1337_v25
				_1337_v25 = Companion_D9_.Create_DC26_(p1, _1320_v8, (p0).Select((Companion_Default___.SafeIndex((_1334_v22).F14(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(bool))
				var _1338_v26 _dafny.MultiSet
				_ = _1338_v26
				_1338_v26 = _dafny.MultiSetOf(p1, (_1337_v25).Dtor_cf41())
				_1338_v26 = (_1338_v26).Union((_1338_v26).Intersection(Companion_Default___.Fm39(p1, globalState)))
				var _1339_v27 _dafny.Set
				_ = _1339_v27
				_1339_v27 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_1311_v0, _1311_v0))
				_1339_v27 = _1339_v27
				var _1340_v28 _dafny.Array
				_ = _1340_v28
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw174
				_1340_v28 = _nw174
				var _1341_v29 _dafny.Sequence
				_ = _1341_v29
				_1341_v29 = _dafny.SeqOf(_1340_v28, _1340_v28)
				var _1342_v30 _dafny.Set
				_ = _1342_v30
				_1342_v30 = _dafny.SetOf(p1, p1)
				var _1343_v31 _dafny.Map
				_ = _1343_v31
				_1343_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1342_v30)
				var _1344_v32 _dafny.Map
				_ = _1344_v32
				_1344_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _1345_v33 D8
				_ = _1345_v33
				_1345_v33 = Companion_D8_.Create_DC23_(p1, p1, (_1344_v32).Cardinality())
				var _1346_v34 _dafny.Map
				_ = _1346_v34
				_1346_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1342_v30).Cardinality())
				var _1347_v35 _dafny.Array
				_ = _1347_v35
				var _nwElement0_34 bool = p1
				_ = _nwElement0_34
				var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(19))
				_ = _nw175
				(_nw175).ArraySet1(_nwElement0_34, 0)
				(_nw175).ArraySet1(!(((_dafny.Zero).Minus((_dafny.Zero).Minus((_1334_v22).F14()))).Cmp(_dafny.IntOfUint32((_1341_v29).Cardinality())) <= 0), 1)
				(_nw175).ArraySet1(p1, 2)
				(_nw175).ArraySet1(!(p1), 3)
				(_nw175).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1316_v5, (Companion_Default___.SafeIndex((_1334_v22).F14(), _dafny.IntOfUint32((_1316_v5).Cardinality()))).Uint32(), (_1343_v31).Cardinality()), _dafny.SeqOf(_1315_v4, (_1334_v22).F14())), 4)
				(_nw175).ArraySet1((_1345_v33).Dtor_cf36(), 5)
				(_nw175).ArraySet1((_1346_v34).Contains(p1), 6)
				(_nw175).ArraySet1((_1315_v4).Cmp(_1315_v4) <= 0, 7)
				(_nw175).ArraySet1(p1, 8)
				(_nw175).ArraySet1(p1, 9)
				(_nw175).ArraySet1(!(p1) || (true), 10)
				(_nw175).ArraySet1(p1, 11)
				(_nw175).ArraySet1(p1, 12)
				(_nw175).ArraySet1(!(p1), 13)
				(_nw175).ArraySet1(!(p1), 14)
				(_nw175).ArraySet1(!(p1), 15)
				(_nw175).ArraySet1((p1) || (p1), 16)
				(_nw175).ArraySet1(false, 17)
				(_nw175).ArraySet1(p1, 18)
				_1347_v35 = _nw175
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_1347_v35), 0))
				_ = _index197
				(_1347_v35).ArraySet1(p1, (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_1347_v35), 0))
				_ = _index198
				(_1347_v35).ArraySet1(Companion_Default___.Fm19((_1334_v22).F14(), p1, globalState), (_index198).Int())
				var _1348_v36 _dafny.Map
				_ = _1348_v36
				_1348_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1316_v5)).Cardinality(), _1341_v29)
				var _rhs201 bool = _dafny.Companion_Sequence_.Equal(_1341_v29, (func() _dafny.Sequence {
					if (_1348_v36).Contains((_1334_v22).F14()) {
						return (_1348_v36).Get((_1334_v22).F14()).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_1340_v28)
				})())
				_ = _rhs201
				var _rhs202 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1311_v0, _1311_v0)
				_ = _rhs202
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				_lhs137.F3 = _rhs201
				_1311_v0 = _rhs202
			} else if _source20.Is_DC12() {
				var _1349___mcc_h0 _dafny.Int = _source20.Get_().(D4_DC12).Cf18
				_ = _1349___mcc_h0
				var _1350_cf18 _dafny.Int = _1349___mcc_h0
				_ = _1350_cf18
				var _1351_v37 _dafny.Array
				_ = _1351_v37
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw176
				_1351_v37 = _nw176
				var _1352_v38 _dafny.Sequence
				_ = _1352_v38
				_1352_v38 = _dafny.SeqOf(_1351_v37, _1351_v37, _1351_v37)
				_1315_v4 = _dafny.IntOfUint32((_1352_v38).Cardinality())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1351_v37), 0))
				_ = _index199
				(_1351_v37).ArraySet1(_1350_cf18, (_index199).Int())
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1351_v37), 0))
				_ = _index200
				(_1351_v37).ArraySet1(Companion_Default___.Fm22(Companion_Default___.Fm29((_1316_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-759), _dafny.IntOfUint32((_1316_v5).Cardinality()))).Uint32()).(_dafny.Int), _1350_cf18, _dafny.UnicodeSeqOfUtf8Bytes("kkiqvf"), _1350_cf18, globalState), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_1350_cf18)), p1, globalState), (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1351_v37), 0))
				_ = _index201
				(_1351_v37).ArraySet1(Companion_Default___.Fm11(p1, globalState), (_index201).Int())
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1351_v37), 0))
				_ = _index202
				(_1351_v37).ArraySet1(_1350_cf18, (_index202).Int())
			} else {
				var _1353___mcc_h1 _dafny.Set = _source20.Get_().(D4_DC10).Cf17
				_ = _1353___mcc_h1
				var _1354_cf17 _dafny.Set = _1353___mcc_h1
				_ = _1354_cf17
				(globalState).F3 = p1
				var _1355_v39 _dafny.Array
				_ = _1355_v39
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_24
				var _nw177 _dafny.Array
				_ = _nw177
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw177 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = (func(_1356_cf17 _dafny.Set) func(_dafny.Int) _dafny.Int {
						return func(_1357_i1 _dafny.Int) _dafny.Int {
							return (_1357_i1).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg98 _dafny.Int) interface{} {
									return coer98(arg98)
								}
							}((func(_1358_cf17 _dafny.Set) func(_dafny.Int) _dafny.Int {
								return func(_1359_i2 _dafny.Int) _dafny.Int {
									return (_1358_cf17).Cardinality()
								}
							})(_1356_cf17)))).Cardinality()))
						}
					})(_1354_cf17)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw177 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw177).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw177).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1355_v39 = _nw177
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1355_v39), 0))
				_ = _index203
				(_1355_v39).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if p1 {
						return (_1334_v22).F14()
					}
					return (_1334_v22).F14()
				})()), (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1355_v39), 0))
				_ = _index204
				(_1355_v39).ArraySet1((_1334_v22).F14(), (_index204).Int())
				var _1360_v40 _dafny.Array
				_ = _1360_v40
				var _nwElement0_35 bool = !(p1)
				_ = _nwElement0_35
				var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(3))
				_ = _nw178
				(_nw178).ArraySet1(_nwElement0_35, 0)
				(_nw178).ArraySet1(_dafny.Companion_Sequence_.Contains(_1311_v0, _1320_v8), 1)
				(_nw178).ArraySet1(p1, 2)
				_1360_v40 = _nw178
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))
				_ = _index205
				(_1360_v40).ArraySet1(p1, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1355_v39), 0))
				_ = _index206
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1355_v39), 0))
				_ = _index207
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))
				_ = _index208
				var _rhs203 _dafny.Int = (_1334_v22).F14()
				_ = _rhs203
				var _rhs204 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1318_v7).Contains(p1) {
						return (_1318_v7).Get(p1).(_dafny.Sequence)
					}
					return (_this).Fm1(_1323_v11, globalState)
				})(), _dafny.UnicodeSeqOfUtf8Bytes("wucl"))).Cardinality())).Plus((_dafny.MultiSetOf(false)).Cardinality())
				_ = _rhs204
				var _rhs205 bool = p1
				_ = _rhs205
				var _rhs206 bool = (_1315_v4).Cmp(_1315_v4) == 0
				_ = _rhs206
				var _lhs138 _dafny.Array = _1355_v39
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1355_v39), 0))
				_ = _lhs139
				var _lhs140 _dafny.Array = _1355_v39
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1355_v39), 0))
				_ = _lhs141
				var _lhs142 _dafny.Array = _1360_v40
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))
				_ = _lhs143
				(_lhs138).ArraySet1(_rhs203, (_lhs139).Int())
				(_lhs140).ArraySet1(_rhs204, (_lhs141).Int())
				(_lhs142).ArraySet1(_rhs205, (_lhs143).Int())
				r2 = _rhs206
				var _1361_v41 _dafny.Map
				_ = _1361_v41
				_1361_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1360_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))).Int()).(bool), p1)
				var _1362_v42 _dafny.Map
				_ = _1362_v42
				_1362_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1320_v8, _1361_v41)
				var _1363_v43 *C4
				_ = _1363_v43
				var _nw179 *C4 = New_C4_()
				_ = _nw179
				_nw179.Ctor__((_1334_v22).F14(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23((_1360_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))).Int()).(bool), _1362_v42, (_1360_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))).Int()).(bool), (_1334_v22).F14(), globalState), _dafny.IntOfInt64(848)))
				_1363_v43 = _nw179
				(globalState).F13 = (_1360_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1360_v40), 0))).Int()).(bool)
			}
			(_1334_v22).M5((func() _dafny.Sequence {
				if true {
					return _1311_v0
				}
				return _1311_v0
			})(), globalState)
		}
		var _1364_v44 _dafny.Sequence
		_ = _1364_v44
		_1364_v44 = _dafny.SeqOf(_1311_v0)
		var _1365_v45 D0
		_ = _1365_v45
		_1365_v45 = Companion_D0_.Create_DC1_(p1, _1364_v44, _1311_v0)
		var _1366_v46 _dafny.Map
		_ = _1366_v46
		_1366_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1365_v45).Dtor_cf3(), _1323_v11)
		r0 = (func() _dafny.Map {
			if (_1366_v46).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1367_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1368_i3 _dafny.Int) _dafny.CodePoint {
					return _1367_v8
				}
			})(_1320_v8)))) {
				return (_1366_v46).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1369_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1370_i3 _dafny.Int) _dafny.CodePoint {
						return _1369_v8
					}
				})(_1320_v8)))).(_dafny.Map)
			}
			return (Companion_Default___.Fm31(globalState)).Merge(_1323_v11)
		})()
		r1 = Companion_Default___.Fm25(p1, Companion_Default___.Fm30(globalState), globalState)
		r2 = p1
		var _1371_v47 _dafny.Array
		_ = _1371_v47
		var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(16))
		_ = _nw180
		_1371_v47 = _nw180
		r3 = _1371_v47
		return r0, r1, r2, r3
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	dummy byte
}

func New_C7_() *C7 {
	_this := C7{}

	return &_this
}

type CompanionStruct_C7_ struct {
}

var Companion_C7_ = CompanionStruct_C7_{}

func (_this *C7) Equals(other *C7) bool {
	return _this == other
}

func (_this *C7) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C7)
	return ok && _this.Equals(other)
}

func (*C7) String() string {
	return "_module.C7"
}

func Type_C7_() _dafny.TypeDescriptor {
	return type_C7_{}
}

type type_C7_ struct {
}

func (_this type_C7_) Default() interface{} {
	return (*C7)(nil)
}

func (_this type_C7_) String() string {
	return "main.C7"
}
func (_this *C7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) Ctor__() {
	{
	}
}
func (_this *C7) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true, false)).Cardinality()))))).Dtor_cf0()))
	}
}
func (_this *C7) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return (Companion_D0_.Create_DC1_(false, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cfv"), _dafny.UnicodeSeqOfUtf8Bytes("lo")), _dafny.UnicodeSeqOfUtf8Bytes("lohm"))).Dtor_cf3()
	}
}
func (_this *C7) M0(p0 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1372_v0 _dafny.Int
		_ = _1372_v0
		_1372_v0 = _dafny.IntOfInt64(57)
		var _1373_v1 _dafny.Map
		_ = _1373_v1
		_1373_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v0, _1372_v0)
		_1373_v1 = (_1373_v1).Update(_1372_v0, _dafny.IntOfInt64(31))
		var _1374_v2 bool
		_ = _1374_v2
		_1374_v2 = false
		var _1375_v3 D0
		_ = _1375_v3
		_1375_v3 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(664))
		var _1376_v4 D0
		_ = _1376_v4
		_1376_v4 = Companion_D0_.Create_DC3_(_1375_v3)
		var _pat_let_tv25 = _1375_v3
		_ = _pat_let_tv25
		(globalState).F13 = Companion_Default___.Fm2(p0, (func() D0 {
			if _1374_v2 {
				return func(_pat_let27_0 D0) D0 {
					return func(_1377_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let28_0 D0) D0 {
							return func(_1378_dt__update_hcf5_h0 D0) D0 {
								return Companion_D0_.Create_DC3_(_1378_dt__update_hcf5_h0)
							}(_pat_let28_0)
						}(_pat_let_tv25)
					}(_pat_let27_0)
				}(_1376_v4)
			}
			return _1376_v4
		})(), Companion_Default___.SafeDivisionInt(_1372_v0, (_dafny.Zero).Minus(_1372_v0)), globalState)
		_1372_v0 = (_this).Fm0((_this).Fm0(_1372_v0, p0, globalState), p0, globalState)
		var _1379_v5 _dafny.Sequence
		_ = _1379_v5
		_1379_v5 = _dafny.UnicodeSeqOfUtf8Bytes("iqufbbgna")
		var _1380_v6 _dafny.Sequence
		_ = _1380_v6
		_1380_v6 = _dafny.SeqOf(_1379_v5, _1379_v5)
		var _1381_v7 D0
		_ = _1381_v7
		_1381_v7 = Companion_D0_.Create_DC1_(_1374_v2, _1380_v6, (func() _dafny.Sequence {
			if _1374_v2 {
				return _1379_v5
			}
			return _1379_v5
		})())
		var _source21 D0 = _1381_v7
		_ = _source21
		if _source21.Is_DC1() {
			var _1382___mcc_h0 bool = _source21.Get_().(D0_DC1).Cf1
			_ = _1382___mcc_h0
			var _1383___mcc_h1 _dafny.Sequence = _source21.Get_().(D0_DC1).Cf2
			_ = _1383___mcc_h1
			var _1384___mcc_h2 _dafny.Sequence = _source21.Get_().(D0_DC1).Cf3
			_ = _1384___mcc_h2
			var _1385_cf3 _dafny.Sequence = _1384___mcc_h2
			_ = _1385_cf3
			var _1386_cf2 _dafny.Sequence = _1383___mcc_h1
			_ = _1386_cf2
			var _1387_cf1 bool = _1382___mcc_h0
			_ = _1387_cf1
			var _1388_v8 *C5
			_ = _1388_v8
			var _nw181 *C5 = New_C5_()
			_ = _nw181
			_nw181.Ctor__()
			_1388_v8 = _nw181
			var _1389_v9 T1
			_ = _1389_v9
			var _nw182 *C5 = New_C5_()
			_ = _nw182
			_nw182.Ctor__()
			_1389_v9 = _nw182
			var _1390_v10 _dafny.Array
			_ = _1390_v10
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_25
			var _nw183 _dafny.Array
			_ = _nw183
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw183 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_1391_v2 bool) func(_dafny.Int) bool {
					return func(_1392_i0 _dafny.Int) bool {
						return (Companion_D0_.Create_DC2_(_1391_v2)).Dtor_cf4()
					}
				})(_1374_v2)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw183 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw183).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw183).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1390_v10 = _nw183
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))
			_ = _index209
			(_1390_v10).ArraySet1(_1374_v2, (_index209).Int())
			var _1393_v11 _dafny.Sequence
			_ = _1393_v11
			_1393_v11 = _dafny.SeqOf(_1387_cf1, _1387_cf1, _1374_v2, _1374_v2)
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))
			_ = _index210
			var _rhs207 bool = _1387_cf1
			_ = _rhs207
			var _rhs208 T1 = (func() T1 {
				if _1387_cf1 {
					return _1389_v9
				}
				return (func() T1 {
					if _1387_cf1 {
						return _1389_v9
					}
					return _1389_v9
				})()
			})()
			_ = _rhs208
			var _rhs209 bool = _1374_v2
			_ = _rhs209
			var _rhs210 bool = (_1374_v2) && ((_dafny.IntOfUint32((_1393_v11).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_1372_v0, _dafny.IntOfInt64(463), _1372_v0)).Cardinality())) < 0)
			_ = _rhs210
			var _lhs144 *GlobalState = globalState
			_ = _lhs144
			var _lhs145 _dafny.Array = _1390_v10
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))
			_ = _lhs146
			_1374_v2 = _rhs207
			_1389_v9 = _rhs208
			_lhs144.F3 = _rhs209
			(_lhs145).ArraySet1(_rhs210, (_lhs146).Int())
			if _1374_v2 {
				_1393_v11 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm27((_1372_v0).Minus(_1372_v0), (_1372_v0).Times(_1372_v0), (_1372_v0).Plus(_1372_v0), globalState), (Companion_Default___.SafeIndex(_1372_v0, _dafny.IntOfUint32((Companion_Default___.Fm27((_1372_v0).Minus(_1372_v0), (_1372_v0).Times(_1372_v0), (_1372_v0).Plus(_1372_v0), globalState)).Cardinality()))).Uint32(), (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))).Int()).(bool))
				var _1394_v12 _dafny.CodePoint
				_ = _1394_v12
				_1394_v12 = _dafny.CodePoint('o')
				_1394_v12 = p0
				(globalState).F3 = !(_1374_v2)
				var _1395_v13 *C5
				_ = _1395_v13
				var _nw184 *C5 = New_C5_()
				_ = _nw184
				_nw184.Ctor__()
				_1395_v13 = _nw184
				var _1396_v14 _dafny.Set
				_ = _1396_v14
				_1396_v14 = _dafny.SetOf(_1372_v0, _1372_v0)
				var _1397_v15 _dafny.Sequence
				_ = _1397_v15
				_1397_v15 = _dafny.SeqOf((_1396_v14).Cardinality(), _dafny.IntOfInt64(83))
				_1397_v15 = _dafny.Companion_Sequence_.Concatenate(_1397_v15, _1397_v15)
			} else {
				_1372_v0 = _1372_v0
				r1 = _1372_v0
				var _1398_v16 _dafny.Map
				_ = _1398_v16
				_1398_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))).Int()).(bool), _1376_v4)
				var _1399_v17 _dafny.Map
				_ = _1399_v17
				_1399_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v0, _1374_v2)
				var _1400_v18 T0
				_ = _1400_v18
				var _nw185 *C2 = New_C2_()
				_ = _nw185
				_nw185.Ctor__()
				_1400_v18 = _nw185
				var _1401_v19 _dafny.Map
				_ = _1401_v19
				_1401_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1399_v17).Cardinality(), _1400_v18)
				var _1402_v20 D2
				_ = _1402_v20
				_1402_v20 = Companion_D2_.Create_DC7_(true, _1398_v16, ((_1401_v19).Update(_1372_v0, _1400_v18)).Cardinality(), _1387_cf1)
				var _1403_v21 D2
				_ = _1403_v21
				_1403_v21 = Companion_D2_.Create_DC7_((_1402_v20).Dtor_cf14(), _1398_v16, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vwxlge")).Cardinality()), _1374_v2)
				_1372_v0 = Companion_Default___.Fm22(_1403_v21, _1376_v4, true, globalState)
				var _1404_v22 _dafny.Set
				_ = _1404_v22
				_1404_v22 = _dafny.SetOf(!(_1387_cf1), _1374_v2, _1374_v2)
				var _1405_v23 _dafny.Set
				_ = _1405_v23
				_1405_v23 = _dafny.SetOf(_1404_v22)
				var _1406_v24 D6
				_ = _1406_v24
				_1406_v24 = Companion_D6_.Create_DC16_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-888))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1407_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1408_i1 _dafny.Int) _dafny.CodePoint {
						return _1407_p0
					}
				})(p0)))).Cardinality()))
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))
				_ = _index211
				var _rhs211 bool = (_1405_v23).IsDisjointFrom(_dafny.SetOf(Companion_Default___.Fm33(_1387_cf1, _1387_cf1, (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))).Int()).(bool), globalState), _1404_v22, _1404_v22, _1404_v22, _1404_v22))
				_ = _rhs211
				var _rhs212 _dafny.Int = (_1372_v0).Times(_1372_v0)
				_ = _rhs212
				var _rhs213 bool = _1387_cf1
				_ = _rhs213
				var _rhs214 _dafny.Int = (_1406_v24).Dtor_cf24()
				_ = _rhs214
				var _rhs215 bool = (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))).Int()).(bool)
				_ = _rhs215
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				var _lhs148 _dafny.Array = _1390_v10
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1390_v10), 0))
				_ = _lhs149
				var _lhs150 *GlobalState = globalState
				_ = _lhs150
				_lhs147.F5 = _rhs211
				r1 = _rhs212
				(_lhs148).ArraySet1(_rhs213, (_lhs149).Int())
				r1 = _rhs214
				_lhs150.F3 = _rhs215
				r1 = (_1372_v0).Plus(_1372_v0)
			}
			var _1409_v25 _dafny.Array
			_ = _1409_v25
			var _nw186 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw186
			_1409_v25 = _nw186
			r1 = (_dafny.SetOf(_1409_v25, _1409_v25, _1409_v25, _1409_v25)).Cardinality()
		} else if _source21.Is_DC2() {
			var _1410___mcc_h3 bool = _source21.Get_().(D0_DC2).Cf4
			_ = _1410___mcc_h3
			var _1411_cf4 bool = _1410___mcc_h3
			_ = _1411_cf4
			var _1412_v26 _dafny.Array
			_ = _1412_v26
			var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw187
			_1412_v26 = _nw187
			var _1413_v27 _dafny.Map
			_ = _1413_v27
			_1413_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1412_v26, _dafny.IntOfInt64(636))
			var _1414_v28 D8
			_ = _1414_v28
			_1414_v28 = Companion_D8_.Create_DC22_(_1413_v27)
			var _1415_v29 _dafny.Sequence
			_ = _1415_v29
			_1415_v29 = _dafny.SeqOf(_1372_v0)
			var _pat_let_tv26 = _1412_v26
			_ = _pat_let_tv26
			var _pat_let_tv27 = _1372_v0
			_ = _pat_let_tv27
			var _rhs216 D8 = func(_pat_let29_0 D8) D8 {
				return func(_1416_dt__update__tmp_h1 D8) D8 {
					return func(_pat_let30_0 _dafny.Map) D8 {
						return func(_1417_dt__update_hcf34_h0 _dafny.Map) D8 {
							return Companion_D8_.Create_DC22_(_1417_dt__update_hcf34_h0)
						}(_pat_let30_0)
					}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv26, _pat_let_tv27))
				}(_pat_let29_0)
			}(_1414_v28)
			_ = _rhs216
			var _rhs217 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if _1374_v2 {
					return _1372_v0
				}
				return _1372_v0
			})())
			_ = _rhs217
			var _rhs218 _dafny.Int = (_dafny.IntOfUint32(((func() _dafny.Sequence {
				if _1411_cf4 {
					return _1415_v29
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg102 _dafny.Int) interface{} {
						return coer102(arg102)
					}
				}((func(_1418_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1419_i2 _dafny.Int) _dafny.Int {
						return _1418_v0
					}
				})(_1372_v0)))
			})()).Cardinality())).Times((_dafny.Zero).Minus(_1372_v0))
			_ = _rhs218
			_1414_v28 = _rhs216
			r1 = _rhs217
			r1 = _rhs218
			var _1420_v30 _dafny.Map
			_ = _1420_v30
			_1420_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, (_dafny.SetOf(_1411_cf4, _1374_v2, Companion_Default___.Fm6(globalState))).Cardinality())
			_1373_v1 = (_1373_v1).Update((_1420_v30).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-311), _1372_v0)))
			var _1421_v32 *C4
			_ = _1421_v32
			var _nw188 *C4 = New_C4_()
			_ = _nw188
			_nw188.Ctor__(_1372_v0, func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter39 := _dafny.Iterate((_dafny.MultiSetOf(p0, p0)).Elements()); ; {
					_compr_37, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _1422_v31 _dafny.CodePoint
					_1422_v31 = interface{}(_compr_37).(_dafny.CodePoint)
					if (_dafny.MultiSetOf(p0, p0)).Contains(_1422_v31) {
						_coll37.Add(_1422_v31, _1372_v0)
					}
				}
				return _coll37.ToMap()
			}())
			_1421_v32 = _nw188
			var _1423_v33 *C1
			_ = _1423_v33
			var _nw189 *C1 = New_C1_()
			_ = _nw189
			_nw189.Ctor__()
			_1423_v33 = _nw189
			var _nw190 *C1 = New_C1_()
			_ = _nw190
			_nw190.Ctor__()
			_1423_v33 = _nw190
		} else if _source21.Is_DC0() {
			var _1424___mcc_h4 _dafny.Int = _source21.Get_().(D0_DC0).Cf0
			_ = _1424___mcc_h4
			var _1425_cf0 _dafny.Int = _1424___mcc_h4
			_ = _1425_cf0
			var _1426_v34 _dafny.Map
			_ = _1426_v34
			_1426_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if _1374_v2 {
					return _1374_v2
				}
				return _1374_v2
			})()), _1374_v2)
			_1426_v34 = (_1426_v34).Update(_1374_v2, _1374_v2)
			(globalState).F9 = !(_1374_v2)
			_1425_cf0 = _1372_v0
			var _1427_v35 _dafny.Map
			_ = _1427_v35
			_1427_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1379_v5, (func() bool {
				if _1374_v2 {
					return !(_1374_v2)
				}
				return _1374_v2
			})())
			_1427_v35 = (_1427_v35).Update(_1379_v5, !((Companion_Default___.Fm11(_1374_v2, globalState)).Cmp(_1372_v0) == 0))
		} else {
			var _1428___mcc_h5 D0 = _source21.Get_().(D0_DC3).Cf5
			_ = _1428___mcc_h5
			var _1429_cf5 D0 = _1428___mcc_h5
			_ = _1429_cf5
			var _1430_v36 _dafny.MultiSet
			_ = _1430_v36
			_1430_v36 = _dafny.MultiSetOf(_dafny.SeqOf(p0))
			var _1431_v37 _dafny.Map
			_ = _1431_v37
			_1431_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if _1374_v2 {
					return _1430_v36
				}
				return _1430_v36
			})(), _1372_v0)
			var _1432_v38 _dafny.Set
			_ = _1432_v38
			_1432_v38 = _dafny.SetOf(!(_1374_v2))
			var _1433_v39 _dafny.Sequence
			_ = _1433_v39
			_1433_v39 = _dafny.SeqOf(_dafny.SetOf(_1374_v2), _1432_v38, _dafny.SetOf(false, _1374_v2))
			_1431_v37 = (_1431_v37).Update(_1430_v36, ((_1433_v39).Select((Companion_Default___.SafeIndex(_1372_v0, _dafny.IntOfUint32((_1433_v39).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
			var _1434_v40 D4
			_ = _1434_v40
			_1434_v40 = Companion_D4_.Create_DC12_(_1372_v0)
			var _source22 D4 = _1434_v40
			_ = _source22
			if _source22.Is_DC11() {
				r0 = (_1374_v2) || (_1374_v2)
				var _1435_v41 _dafny.Map
				_ = _1435_v41
				_1435_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, _1374_v2)
				var _1436_v42 _dafny.Set
				_ = _1436_v42
				_1436_v42 = _dafny.SetOf(_1372_v0, _1372_v0, _1372_v0, (_dafny.IntOfUint32((_1379_v5).Cardinality())).Plus((_1435_v41).Cardinality()))
				_1436_v42 = _1436_v42
				var _1437_v43 _dafny.Array
				_ = _1437_v43
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_26
				var _nw191 _dafny.Array
				_ = _nw191
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw191 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = (func(_1438_v42 _dafny.Set, _1439_v0 _dafny.Int) func(_dafny.Int) bool {
						return func(_1440_i3 _dafny.Int) bool {
							return (_dafny.MultiSetOf(_dafny.SetOf(_1439_v0))).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1438_v42, _1438_v42)))
						}
					})(_1436_v42, _1372_v0)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw191 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw191).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw191).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1437_v43 = _nw191
				var _1441_v44 D8
				_ = _1441_v44
				_1441_v44 = Companion_D8_.Create_DC23_(_1374_v2, true, _1372_v0)
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_1437_v43), 0))
				_ = _index212
				(_1437_v43).ArraySet1((_1441_v44).Dtor_cf35(), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_1437_v43), 0))
				_ = _index213
				(_1437_v43).ArraySet1(_1374_v2, (_index213).Int())
				var _rhs219 bool = (_1437_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_1437_v43), 0))).Int()).(bool)
				_ = _rhs219
				var _rhs220 _dafny.Sequence = _1379_v5
				_ = _rhs220
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				_lhs151.F9 = _rhs219
				_1379_v5 = _rhs220
			} else if _source22.Is_DC12() {
				var _1442___mcc_h6 _dafny.Int = _source22.Get_().(D4_DC12).Cf18
				_ = _1442___mcc_h6
				var _1443_cf18 _dafny.Int = _1442___mcc_h6
				_ = _1443_cf18
				var _1444_v45 _dafny.Set
				_ = _1444_v45
				_1444_v45 = _dafny.SetOf(_1443_cf18, (_1430_v36).Cardinality())
				var _1445_v46 _dafny.MultiSet
				_ = _1445_v46
				_1445_v46 = _dafny.MultiSetOf(_1444_v45)
				var _rhs221 _dafny.Int = (func() _dafny.Int {
					if _1374_v2 {
						return _1443_cf18
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(600))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg103 _dafny.Int) interface{} {
							return coer103(arg103)
						}
					}((func(_1446_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1447_i4 _dafny.Int) _dafny.CodePoint {
							return _1446_p0
						}
					})(p0)))).Cardinality())
				})()
				_ = _rhs221
				var _rhs222 bool = (_1445_v46).IsDisjointFrom(_1445_v46)
				_ = _rhs222
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				_1443_cf18 = _rhs221
				_lhs152.F13 = _rhs222
				var _1448_v47 *C1
				_ = _1448_v47
				var _nw192 *C1 = New_C1_()
				_ = _nw192
				_nw192.Ctor__()
				_1448_v47 = _nw192
				var _1449_v48 _dafny.Map
				_ = _1449_v48
				_1449_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v0, _1443_cf18), _1444_v45)
				var _1450_v49 _dafny.Array
				_ = _1450_v49
				var _nwElement0_36 bool = _1374_v2
				_ = _nwElement0_36
				var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(19))
				_ = _nw193
				(_nw193).ArraySet1(_nwElement0_36, 0)
				(_nw193).ArraySet1(!(_1374_v2), 1)
				(_nw193).ArraySet1(_1374_v2, 2)
				(_nw193).ArraySet1((_1372_v0).Cmp(_1443_cf18) <= 0, 3)
				(_nw193).ArraySet1(_1374_v2, 4)
				(_nw193).ArraySet1(_1374_v2, 5)
				(_nw193).ArraySet1(_1374_v2, 6)
				(_nw193).ArraySet1(_1374_v2, 7)
				(_nw193).ArraySet1(_1374_v2, 8)
				(_nw193).ArraySet1(_1374_v2, 9)
				(_nw193).ArraySet1((_1449_v48).Contains(_1373_v1), 10)
				(_nw193).ArraySet1(_1374_v2, 11)
				(_nw193).ArraySet1(_1374_v2, 12)
				(_nw193).ArraySet1((_1372_v0).Cmp(_1372_v0) < 0, 13)
				(_nw193).ArraySet1(_1374_v2, 14)
				(_nw193).ArraySet1(_1374_v2, 15)
				(_nw193).ArraySet1((_1374_v2) && (_1374_v2), 16)
				(_nw193).ArraySet1(_1374_v2, 17)
				(_nw193).ArraySet1(_1374_v2, 18)
				_1450_v49 = _nw193
				var _1451_v50 _dafny.MultiSet
				_ = _1451_v50
				_1451_v50 = _dafny.MultiSetOf(_1374_v2)
				var _1452_v51 _dafny.Map
				_ = _1452_v51
				_1452_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_cf18, _1448_v47)
				var _1453_v52 _dafny.Map
				_ = _1453_v52
				_1453_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, _dafny.IntOfInt64(815))
				var _rhs223 bool = (!(_1374_v2)) || ((((_1451_v50).Update(_1374_v2, Companion_Default___.Abs(_dafny.IntOfInt64(-436)))).Cardinality()).Cmp((_dafny.Zero).Minus(_1372_v0)) == 0)
				_ = _rhs223
				var _rhs224 *C1 = (func() *C1 {
					if (_1452_v51).Contains(_1443_cf18) {
						return (_1452_v51).Get(_1443_cf18).(*C1)
					}
					return _1448_v47
				})()
				_ = _rhs224
				var _rhs225 _dafny.Array = _1450_v49
				_ = _rhs225
				var _rhs226 _dafny.Int = (_1443_cf18).Plus(_dafny.IntOfUint32((_1379_v5).Cardinality()))
				_ = _rhs226
				var _rhs227 _dafny.Int = (_this).Fm0((func() _dafny.Int {
					if (_1453_v52).Contains(_1374_v2) {
						return (_1453_v52).Get(_1374_v2).(_dafny.Int)
					}
					return _1443_cf18
				})(), (func() _dafny.CodePoint {
					if !(_1374_v2) {
						return p0
					}
					return p0
				})(), globalState)
				_ = _rhs227
				var _lhs153 *GlobalState = globalState
				_ = _lhs153
				_lhs153.F13 = _rhs223
				_1448_v47 = _rhs224
				_1450_v49 = _rhs225
				r1 = _rhs226
				r1 = _rhs227
				_1374_v2 = _1374_v2
				_1379_v5 = _dafny.UnicodeSeqOfUtf8Bytes("uf")
			} else {
				var _1454___mcc_h7 _dafny.Set = _source22.Get_().(D4_DC10).Cf17
				_ = _1454___mcc_h7
				var _1455_cf17 _dafny.Set = _1454___mcc_h7
				_ = _1455_cf17
				var _1456_v53 _dafny.Array
				_ = _1456_v53
				var _nw194 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw194
				_1456_v53 = _nw194
				var _1457_v54 _dafny.MultiSet
				_ = _1457_v54
				_1457_v54 = _dafny.MultiSetOf(_1374_v2)
				var _1458_v55 D6
				_ = _1458_v55
				_1458_v55 = Companion_D6_.Create_DC15_(_1457_v54)
				var _1459_v56 _dafny.Map
				_ = _1459_v56
				_1459_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, (func() bool {
					if _1374_v2 {
						return !(!(_1374_v2))
					}
					return true
				})())
				var _rhs228 _dafny.Array = _1456_v53
				_ = _rhs228
				var _rhs229 D6 = _1458_v55
				_ = _rhs229
				var _rhs230 _dafny.Int = _1372_v0
				_ = _rhs230
				var _rhs231 bool = (func() bool {
					if (_1459_v56).Contains(_1374_v2) {
						return (_1459_v56).Get(_1374_v2).(bool)
					}
					return _1374_v2
				})()
				_ = _rhs231
				var _rhs232 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm11((_1372_v0).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lecq")).Cardinality())) == 0, globalState))
				_ = _rhs232
				var _lhs154 *GlobalState = globalState
				_ = _lhs154
				_1456_v53 = _rhs228
				_1458_v55 = _rhs229
				r1 = _rhs230
				_lhs154.F9 = _rhs231
				_1372_v0 = _rhs232
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1456_v53), 0))
				_ = _index214
				(_1456_v53).ArraySet1(_1374_v2, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1456_v53), 0))
				_ = _index215
				(_1456_v53).ArraySet1(!(Companion_Default___.Fm6(globalState)) || (_1374_v2), (_index215).Int())
				var _1460_v57 _dafny.Sequence
				_ = _1460_v57
				_1460_v57 = _dafny.SeqOf(_1374_v2, _1374_v2)
				var _1461_v58 _dafny.MultiSet
				_ = _1461_v58
				_1461_v58 = _dafny.MultiSetOf(p0)
				var _1462_v59 _dafny.MultiSet
				_ = _1462_v59
				_1462_v59 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.IntOfUint32((_1460_v57).Cardinality())).Times((_1461_v58).Cardinality())), (_1372_v0).Minus(_1372_v0), _1372_v0)
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1456_v53), 0))
				_ = _index216
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1456_v53), 0))
				_ = _index217
				var _rhs233 bool = _1374_v2
				_ = _rhs233
				var _rhs234 bool = (_1372_v0).Cmp(_1372_v0) < 0
				_ = _rhs234
				var _rhs235 _dafny.MultiSet = (_1462_v59).Difference(_1462_v59)
				_ = _rhs235
				var _lhs155 _dafny.Array = _1456_v53
				_ = _lhs155
				var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1456_v53), 0))
				_ = _lhs156
				var _lhs157 _dafny.Array = _1456_v53
				_ = _lhs157
				var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1456_v53), 0))
				_ = _lhs158
				(_lhs155).ArraySet1(_rhs233, (_lhs156).Int())
				(_lhs157).ArraySet1(_rhs234, (_lhs158).Int())
				_1462_v59 = _rhs235
				var _1463_v60 _dafny.Map
				_ = _1463_v60
				_1463_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1376_v4)
				var _1464_v61 D2
				_ = _1464_v61
				_1464_v61 = Companion_D2_.Create_DC7_((_1456_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1456_v53), 0))).Int()).(bool), _1463_v60, _dafny.IntOfInt64(229), _1374_v2)
				var _1465_v62 _dafny.Array
				_ = _1465_v62
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_27
				var _nw195 _dafny.Array
				_ = _nw195
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw195 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Int = (func(_1466_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1467_i5 _dafny.Int) _dafny.Int {
							return (_1467_i5).Plus(_1466_v0)
						}
					})(_1372_v0)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw195 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw195).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw195).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1465_v62 = _nw195
				var _1468_v63 _dafny.Map
				_ = _1468_v63
				_1468_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_v62, _1372_v0)
				var _pat_let_tv28 = _1375_v3
				_ = _pat_let_tv28
				var _rhs236 _dafny.MultiSet = ((_1462_v59).Intersection(_1462_v59)).Update(Companion_Default___.Fm22(_1464_v61, _1376_v4, (_1456_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1456_v53), 0))).Int()).(bool), globalState), Companion_Default___.Abs((_dafny.Zero).Minus(_1372_v0)))
				_ = _rhs236
				var _rhs237 _dafny.Int = Companion_Default___.Fm22(_1464_v61, func(_pat_let31_0 D0) D0 {
					return func(_1469_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let32_0 D0) D0 {
							return func(_1470_dt__update_hcf5_h1 D0) D0 {
								return Companion_D0_.Create_DC3_(_1470_dt__update_hcf5_h1)
							}(_pat_let32_0)
						}(_pat_let_tv28)
					}(_pat_let31_0)
				}(_1376_v4), (_1460_v57).Select((Companion_Default___.SafeIndex(_1372_v0, _dafny.IntOfUint32((_1460_v57).Cardinality()))).Uint32()).(bool), globalState)
				_ = _rhs237
				var _rhs238 bool = !((_1468_v63).Contains(_1465_v62))
				_ = _rhs238
				var _rhs239 bool = Companion_Default___.Fm19(_1372_v0, !(_1374_v2) || (_1374_v2), globalState)
				_ = _rhs239
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				_1462_v59 = _rhs236
				_1372_v0 = _rhs237
				_lhs159.F5 = _rhs238
				_lhs160.F13 = _rhs239
				var _1471_v64 D8
				_ = _1471_v64
				_1471_v64 = Companion_D8_.Create_DC24_()
				_1471_v64 = _1471_v64
			}
			var _1472_v65 *C6
			_ = _1472_v65
			var _nw196 *C6 = New_C6_()
			_ = _nw196
			_nw196.Ctor__()
			_1472_v65 = _nw196
			var _1473_v66 _dafny.Sequence
			_ = _1473_v66
			_1473_v66 = _dafny.SeqOf(_1472_v65)
			var _1474_v67 D11
			_ = _1474_v67
			_1474_v67 = Companion_D11_.Create_DC29_(_dafny.MultiSetFromSeq(_1473_v66))
			var _1475_v68 _dafny.Array
			_ = _1475_v68
			var _nwElement0_37 _dafny.Int = Companion_Default___.SafeModuloInt(_1372_v0, _dafny.IntOfInt64(940))
			_ = _nwElement0_37
			var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(6))
			_ = _nw197
			(_nw197).ArraySet1(_nwElement0_37, 0)
			(_nw197).ArraySet1(_1372_v0, 1)
			(_nw197).ArraySet1((_this).Fm0(_1372_v0, p0, globalState), 2)
			(_nw197).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(587), _1372_v0), 3)
			(_nw197).ArraySet1(_1372_v0, 4)
			(_nw197).ArraySet1((_1372_v0).Plus(((_1474_v67).Dtor_cf46()).Cardinality()), 5)
			_1475_v68 = _nw197
			_1475_v68 = _1475_v68
			var _1476_v69 _dafny.Sequence
			_ = _1476_v69
			_1476_v69 = _dafny.SeqOf(_1372_v0)
			var _1477_v70 _dafny.Map
			_ = _1477_v70
			_1477_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1476_v69).Cardinality()), _1379_v5)
			_1477_v70 = (_1477_v70).Merge(_1477_v70)
		}
		var _1478_v71 _dafny.Map
		_ = _1478_v71
		_1478_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, _1374_v2)
		var _1479_v72 _dafny.Map
		_ = _1479_v72
		_1479_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, _1478_v71)
		var _1480_v73 _dafny.Sequence
		_ = _1480_v73
		_1480_v73 = _dafny.SeqOf(_1372_v0, _1372_v0)
		var _1481_v74 _dafny.Sequence
		_ = _1481_v74
		_1481_v74 = _dafny.SeqOf(_1374_v2)
		var _1482_v75 _dafny.Array
		_ = _1482_v75
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_28
		var _nw198 _dafny.Array
		_ = _nw198
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw198 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) _dafny.Sequence = (func(_1483_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1484_i6 _dafny.Int) _dafny.Sequence {
					return _1483_v5
				}
			})(_1379_v5)
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw198 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw198).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw198).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_1482_v75 = _nw198
		var _1485_v76 D7
		_ = _1485_v76
		_1485_v76 = Companion_D7_.Create_DC19_(Companion_D1_.Create_DC5_(_1374_v2, _1379_v5, _1481_v74), _dafny.IntOfUint32((_1379_v5).Cardinality()), _1374_v2, _1482_v75)
		var _1486_v77 _dafny.Sequence
		_ = _1486_v77
		_1486_v77 = _dafny.SeqOf(_1480_v73)
		var _1487_v78 _dafny.Sequence
		_ = _1487_v78
		_1487_v78 = _dafny.SeqOf((_1486_v77).Select((Companion_Default___.SafeIndex(_1372_v0, _dafny.IntOfUint32((_1486_v77).Cardinality()))).Uint32()).(_dafny.Sequence), _1480_v73)
		var _1488_v79 _dafny.Array
		_ = _1488_v79
		var _nwElement0_38 _dafny.Int = _1372_v0
		_ = _nwElement0_38
		var _nw199 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(15))
		_ = _nw199
		(_nw199).ArraySet1(_nwElement0_38, 0)
		(_nw199).ArraySet1((_dafny.Zero).Minus(_1372_v0), 1)
		(_nw199).ArraySet1(_1372_v0, 2)
		(_nw199).ArraySet1(_1372_v0, 3)
		(_nw199).ArraySet1((((func() _dafny.Map {
			if (_1479_v72).Contains(_1374_v2) {
				return (_1479_v72).Get(_1374_v2).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v2, _1374_v2)
		})()).Merge(_1478_v71)).Cardinality(), 4)
		(_nw199).ArraySet1(_1372_v0, 5)
		(_nw199).ArraySet1(_1372_v0, 6)
		(_nw199).ArraySet1((_dafny.Zero).Minus(_1372_v0), 7)
		(_nw199).ArraySet1(_1372_v0, 8)
		(_nw199).ArraySet1(_dafny.IntOfUint32((_1480_v73).Cardinality()), 9)
		(_nw199).ArraySet1((func() _dafny.Int {
			if !(_1374_v2) {
				return (_1485_v76).Dtor_cf28()
			}
			return _dafny.IntOfUint32((_1486_v77).Cardinality())
		})(), 10)
		(_nw199).ArraySet1(_1372_v0, 11)
		(_nw199).ArraySet1(_1372_v0, 12)
		(_nw199).ArraySet1((_1372_v0).Times(_1372_v0), 13)
		(_nw199).ArraySet1(_1372_v0, 14)
		_1488_v79 = _nw199
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))
		_ = _index218
		(_1488_v79).ArraySet1((_dafny.IntOfInt64(460)).Plus(_1372_v0), (_index218).Int())
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))
		_ = _index219
		(_1488_v79).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1372_v0)), (_index219).Int())
		var _1489_i7 _dafny.Int
		_ = _1489_i7
		_1489_i7 = _dafny.Zero
		{
			for _1374_v2 {
				{
					if (_1489_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1489_i7 = (_1489_i7).Plus(_dafny.One)
					var _1490_v80 _dafny.Map
					_ = _1490_v80
					_1490_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1379_v5).Cardinality()))
					var _1491_v81 *C4
					_ = _1491_v81
					var _nw200 *C4 = New_C4_()
					_ = _nw200
					_nw200.Ctor__((_1488_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))).Int()).(_dafny.Int), _1490_v80)
					_1491_v81 = _nw200
					var _1492_v82 _dafny.Map
					_ = _1492_v82
					_1492_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1491_v81)
					_1492_v82 = ((_1492_v82).Merge(_1492_v82)).Merge((_1492_v82).Merge(_1492_v82))
					(globalState).F13 = _1374_v2
					var _1493_v83 _dafny.Map
					_ = _1493_v83
					_1493_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1374_v2), _1376_v4)
					var _1494_v86 _dafny.MultiSet
					_ = _1494_v86
					_1494_v86 = _dafny.MultiSetOf(p0, p0, p0)
					var _1495_v87 D9
					_ = _1495_v87
					_1495_v87 = Companion_D9_.Create_DC27_(!(_1374_v2), _dafny.IntOfUint32((_1379_v5).Cardinality()), _1374_v2)
					var _1496_v88 D2
					_ = _1496_v88
					_1496_v88 = Companion_D2_.Create_DC7_(_1374_v2, _1493_v83, (func() _dafny.Map {
						var _coll38 = _dafny.NewMapBuilder()
						_ = _coll38
						for _iter40 := _dafny.Iterate((func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter41 := _dafny.Iterate((_1494_v86).Elements()); ; {
								_compr_39, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _1497_v85 _dafny.CodePoint
								_1497_v85 = interface{}(_compr_39).(_dafny.CodePoint)
								if (_1494_v86).Contains(_1497_v85) {
									_coll39.Add(_1497_v85, (_1491_v81).F14())
								}
							}
							return _coll39.ToMap()
						}()).Keys().Elements()); ; {
							_compr_38, _ok40 := _iter40()
							if !_ok40 {
								break
							}
							var _1498_v84 _dafny.CodePoint
							_1498_v84 = interface{}(_compr_38).(_dafny.CodePoint)
							if (func() _dafny.Map {
								var _coll40 = _dafny.NewMapBuilder()
								_ = _coll40
								for _iter42 := _dafny.Iterate((_1494_v86).Elements()); ; {
									_compr_40, _ok42 := _iter42()
									if !_ok42 {
										break
									}
									var _1497_v85 _dafny.CodePoint
									_1497_v85 = interface{}(_compr_40).(_dafny.CodePoint)
									if (_1494_v86).Contains(_1497_v85) {
										_coll40.Add(_1497_v85, (_1491_v81).F14())
									}
								}
								return _coll40.ToMap()
							}()).Contains(_1498_v84) {
								_coll38.Add(_1498_v84, _dafny.SetOf(_1374_v2))
							}
						}
						return _coll38.ToMap()
					}()).Cardinality(), !((_1495_v87).Dtor_cf42()))
					r1 = Companion_Default___.Fm22(_1496_v88, _1376_v4, _1374_v2, globalState)
					var _1499_v89 _dafny.Array
					_ = _1499_v89
					var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
					_ = _nw201
					_1499_v89 = _nw201
					var _1500_v90 *C1
					_ = _1500_v90
					var _nw202 *C1 = New_C1_()
					_ = _nw202
					_nw202.Ctor__()
					_1500_v90 = _nw202
					var _1501_v91 _dafny.Array
					_ = _1501_v91
					var _nwElement0_39 *C1 = _1500_v90
					_ = _nwElement0_39
					var _nw203 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(8))
					_ = _nw203
					(_nw203).ArraySet1(_nwElement0_39, 0)
					(_nw203).ArraySet1(_1500_v90, 1)
					(_nw203).ArraySet1(_1500_v90, 2)
					(_nw203).ArraySet1(_1500_v90, 3)
					(_nw203).ArraySet1(_1500_v90, 4)
					(_nw203).ArraySet1(_1500_v90, 5)
					(_nw203).ArraySet1(_1500_v90, 6)
					(_nw203).ArraySet1(_1500_v90, 7)
					_1501_v91 = _nw203
					var _1502_v92 _dafny.Map
					_ = _1502_v92
					_1502_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1501_v91, (_1491_v81).F14())
					var _rhs240 _dafny.Array = _1499_v89
					_ = _rhs240
					var _rhs241 _dafny.Int = (func() _dafny.Int {
						if (_1502_v92).Contains(_1501_v91) {
							return (_1502_v92).Get(_1501_v91).(_dafny.Int)
						}
						return (_1372_v0).Times(_dafny.IntOfInt64(448))
					})()
					_ = _rhs241
					_1499_v89 = _rhs240
					_1372_v0 = _rhs241
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1503_v93 _dafny.MultiSet
		_ = _1503_v93
		_1503_v93 = _dafny.MultiSetOf(_1374_v2)
		var _1504_v94 _dafny.MultiSet
		_ = _1504_v94
		_1504_v94 = _dafny.MultiSetOf((_1488_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))).Int()).(_dafny.Int), (_1503_v93).Cardinality())
		r0 = (_1504_v94).IsSubsetOf((_dafny.MultiSetOf(_1372_v0, (_1488_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))).Int()).(_dafny.Int))).Update(_1372_v0, Companion_Default___.Abs((_1488_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))).Int()).(_dafny.Int))))
		r1 = (_1488_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1488_v79), 0))).Int()).(_dafny.Int)
		return r0, r1
	}
}
func (_this *C7) M1(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) (_dafny.Set, bool, _dafny.Int, bool) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _1505_v0 _dafny.MultiSet
		_ = _1505_v0
		_1505_v0 = _dafny.MultiSetOf(p0, p0, true)
		var _1506_v1 D6
		_ = _1506_v1
		_1506_v1 = Companion_D6_.Create_DC16_((func() _dafny.Int {
			if p0 {
				return (_1505_v0).Cardinality()
			}
			return _dafny.IntOfInt64(489)
		})())
		var _source23 D6 = _1506_v1
		_ = _source23
		if _source23.Is_DC16() {
			var _1507___mcc_h0 _dafny.Int = _source23.Get_().(D6_DC16).Cf24
			_ = _1507___mcc_h0
			var _1508_cf24 _dafny.Int = _1507___mcc_h0
			_ = _1508_cf24
			var _1509_v2 D0
			_ = _1509_v2
			_1509_v2 = Companion_D0_.Create_DC0_(_1508_cf24)
			var _1510_v3 D0
			_ = _1510_v3
			_1510_v3 = Companion_D0_.Create_DC3_(_1509_v2)
			var _1511_v4 _dafny.Map
			_ = _1511_v4
			_1511_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1510_v3)
			var _1512_v5 D2
			_ = _1512_v5
			_1512_v5 = Companion_D2_.Create_DC7_(p0, _1511_v4, _1508_cf24, p0)
			var _1513_v6 _dafny.Sequence
			_ = _1513_v6
			_1513_v6 = _dafny.SeqOf(true)
			var _1514_v7 _dafny.CodePoint
			_ = _1514_v7
			_1514_v7 = _dafny.CodePoint('f')
			var _1515_v8 _dafny.Array
			_ = _1515_v8
			var _nwElement0_40 bool = p0
			_ = _nwElement0_40
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(22))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_40, 0)
			(_nw204).ArraySet1(p0, 1)
			(_nw204).ArraySet1(Companion_Default___.Fm2(_1514_v7, _1510_v3, _1508_cf24, globalState), 2)
			(_nw204).ArraySet1(p0, 3)
			(_nw204).ArraySet1(p0, 4)
			(_nw204).ArraySet1(p0, 5)
			(_nw204).ArraySet1(p0, 6)
			(_nw204).ArraySet1(p0, 7)
			(_nw204).ArraySet1(p0, 8)
			(_nw204).ArraySet1(!(p0), 9)
			(_nw204).ArraySet1(p0, 10)
			(_nw204).ArraySet1(p0, 11)
			(_nw204).ArraySet1(true, 12)
			(_nw204).ArraySet1(!(false), 13)
			(_nw204).ArraySet1(p0, 14)
			(_nw204).ArraySet1(p0, 15)
			(_nw204).ArraySet1(p0, 16)
			(_nw204).ArraySet1(p0, 17)
			(_nw204).ArraySet1(!(p0), 18)
			(_nw204).ArraySet1(p0, 19)
			(_nw204).ArraySet1(p0, 20)
			(_nw204).ArraySet1(p0, 21)
			_1515_v8 = _nw204
			var _1516_v9 _dafny.Sequence
			_ = _1516_v9
			_1516_v9 = _dafny.SeqOf(_1515_v8, _1515_v8)
			var _1517_v10 _dafny.Map
			_ = _1517_v10
			_1517_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-271))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1518_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1519_i0 _dafny.Int) _dafny.CodePoint {
					return _1518_v7
				}
			})(_1514_v7))), p0)
			var _1520_v12 _dafny.Sequence
			_ = _1520_v12
			_1520_v12 = _dafny.SeqOf(_dafny.IntOfInt64(601), _1508_cf24)
			var _1521_v13 _dafny.Map
			_ = _1521_v13
			_1521_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1508_cf24, _1508_cf24)
			var _1522_v14 _dafny.Map
			_ = _1522_v14
			_1522_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1523_v15 _dafny.Set
			_ = _1523_v15
			_1523_v15 = _dafny.SetOf(true)
			var _1524_v16 _dafny.Array
			_ = _1524_v16
			var _nwElement0_41 _dafny.Int = Companion_Default___.Fm22(_1512_v5, _1510_v3, p0, globalState)
			_ = _nwElement0_41
			var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(24))
			_ = _nw205
			(_nw205).ArraySet1(_nwElement0_41, 0)
			(_nw205).ArraySet1((_1508_cf24).Plus((_dafny.Zero).Minus(_1508_cf24)), 1)
			(_nw205).ArraySet1(Companion_Default___.SafeModuloInt(_1508_cf24, (_dafny.Zero).Minus(_1508_cf24)), 2)
			(_nw205).ArraySet1(_1508_cf24, 3)
			(_nw205).ArraySet1(Companion_Default___.Fm22(Companion_D2_.Create_DC7_(p0, _1511_v4, _1508_cf24, p0), _1510_v3, p0, globalState), 4)
			(_nw205).ArraySet1(_1508_cf24, 5)
			(_nw205).ArraySet1(_1508_cf24, 6)
			(_nw205).ArraySet1(Companion_Default___.Fm11(p0, globalState), 7)
			(_nw205).ArraySet1(_1508_cf24, 8)
			(_nw205).ArraySet1(_dafny.IntOfUint32((_1513_v6).Cardinality()), 9)
			(_nw205).ArraySet1(_1508_cf24, 10)
			(_nw205).ArraySet1(_dafny.IntOfUint32((_1516_v9).Cardinality()), 11)
			(_nw205).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_1513_v6).Cardinality()), (func() _dafny.Set {
				var _coll41 = _dafny.NewBuilder()
				_ = _coll41
				for _iter43 := _dafny.Iterate((_1517_v10).Keys().Elements()); ; {
					_compr_41, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _1525_v11 _dafny.Sequence
					_1525_v11 = interface{}(_compr_41).(_dafny.Sequence)
					if (_1517_v10).Contains(_1525_v11) {
						_coll41.Add(_1525_v11)
					}
				}
				return _coll41.ToSet()
			}()).Cardinality()), _1520_v12)).Cardinality()), 12)
			(_nw205).ArraySet1(_1508_cf24, 13)
			(_nw205).ArraySet1(((_1521_v13).Cardinality()).Times((_1522_v14).Cardinality()), 14)
			(_nw205).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((p1).Cardinality()))).Cardinality(), 15)
			(_nw205).ArraySet1(_1508_cf24, 16)
			(_nw205).ArraySet1((_1523_v15).Cardinality(), 17)
			(_nw205).ArraySet1(_1508_cf24, 18)
			(_nw205).ArraySet1(_1508_cf24, 19)
			(_nw205).ArraySet1((_1508_cf24).Minus(_1508_cf24), 20)
			(_nw205).ArraySet1(Companion_Default___.Fm22(_1512_v5, _1510_v3, p0, globalState), 21)
			(_nw205).ArraySet1(((_1505_v0).Cardinality()).Times(_1508_cf24), 22)
			(_nw205).ArraySet1(_1508_cf24, 23)
			_1524_v16 = _nw205
			var _rhs242 _dafny.Int = (Companion_Default___.SafeDivisionInt(_1508_cf24, _1508_cf24)).Times(_1508_cf24)
			_ = _rhs242
			var _rhs243 _dafny.Array = _1524_v16
			_ = _rhs243
			_1508_cf24 = _rhs242
			_1524_v16 = _rhs243
			var _1526_v17 *C5
			_ = _1526_v17
			var _nw206 *C5 = New_C5_()
			_ = _nw206
			_nw206.Ctor__()
			_1526_v17 = _nw206
			var _1527_v18 _dafny.Map
			_ = _1527_v18
			_1527_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1526_v17, _1524_v16)
			_1527_v18 = (_1527_v18).Update(_1526_v17, _1524_v16)
			var _1528_v19 _dafny.Sequence
			_ = _1528_v19
			_1528_v19 = _dafny.SeqOf(p1, p1, p1, p1)
			_1508_cf24 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(54), _dafny.IntOfUint32(((_1528_v19).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1505_v0).Contains(p0) {
					return (_1505_v0).Multiplicity(p0)
				}
				return _1508_cf24
			})(), _dafny.IntOfUint32((_1528_v19).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			var _1529_v20 D8
			_ = _1529_v20
			_1529_v20 = Companion_D8_.Create_DC24_()
			_1529_v20 = _1529_v20
		} else if _source23.Is_DC15() {
			var _1530___mcc_h1 _dafny.MultiSet = _source23.Get_().(D6_DC15).Cf23
			_ = _1530___mcc_h1
			var _1531_cf23 _dafny.MultiSet = _1530___mcc_h1
			_ = _1531_cf23
			var _1532_v21 _dafny.Int
			_ = _1532_v21
			_1532_v21 = _dafny.IntOfInt64(-864)
			(globalState).F3 = (_1532_v21).Cmp(_1532_v21) >= 0
			r2 = _1532_v21
			if (func() bool {
				if p0 {
					return true
				}
				return (_1532_v21).Cmp(_1532_v21) == 0
			})() {
				_1532_v21 = _1532_v21
				(globalState).F3 = (_1532_v21).Cmp(_1532_v21) >= 0
				var _1533_v22 _dafny.MultiSet
				_ = _1533_v22
				_1533_v22 = _dafny.MultiSetOf(_1532_v21, _1532_v21, _1532_v21)
				var _1534_v23 D12
				_ = _1534_v23
				_1534_v23 = Companion_D12_.Create_DC31_(_1533_v22)
				var _1535_v24 _dafny.Array
				_ = _1535_v24
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw207
				_1535_v24 = _nw207
				var _1536_v25 _dafny.MultiSet
				_ = _1536_v25
				_1536_v25 = _dafny.MultiSetOf(_1535_v24, _1535_v24)
				var _1537_v26 _dafny.Map
				_ = _1537_v26
				_1537_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1535_v24, _1535_v24)
				var _rhs244 bool = p0
				_ = _rhs244
				var _rhs245 _dafny.MultiSet = (_1534_v23).Dtor_cf48()
				_ = _rhs245
				var _rhs246 bool = p0
				_ = _rhs246
				var _rhs247 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1536_v25).Contains((func() _dafny.Array {
						if p0 {
							return (func() _dafny.Array {
								if (_1537_v26).Contains(_1535_v24) {
									return (_1537_v26).Get(_1535_v24).(_dafny.Array)
								}
								return _1535_v24
							})()
						}
						return _1535_v24
					})()) {
						return (_1536_v25).Multiplicity((func() _dafny.Array {
							if p0 {
								return (func() _dafny.Array {
									if (_1537_v26).Contains(_1535_v24) {
										return (_1537_v26).Get(_1535_v24).(_dafny.Array)
									}
									return _1535_v24
								})()
							}
							return _1535_v24
						})())
					}
					return _1532_v21
				})())
				_ = _rhs247
				var _rhs248 _dafny.Int = Companion_Default___.SafeDivisionInt(_1532_v21, _1532_v21)
				_ = _rhs248
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				_lhs161.F3 = _rhs244
				_1533_v22 = _rhs245
				_lhs162.F3 = _rhs246
				_1532_v21 = _rhs247
				_1532_v21 = _rhs248
				r2 = _1532_v21
				var _1538_v27 _dafny.Array
				_ = _1538_v27
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw208
				_1538_v27 = _nw208
				_1538_v27 = _1538_v27
			} else {
				(globalState).F9 = (!(p0)) || ((p0) || (Companion_Default___.Fm19(_dafny.IntOfUint32((p1).Cardinality()), p0, globalState)))
				var _1539_v28 _dafny.Map
				_ = _1539_v28
				_1539_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1532_v21).Times((_dafny.Zero).Minus(_1532_v21)), _1532_v21)
				_1539_v28 = _1539_v28
				(globalState).F13 = (_1532_v21).Cmp(_dafny.IntOfInt64(303)) > 0
				r2 = _1532_v21
				var _1540_v29 _dafny.Array
				_ = _1540_v29
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw209
				_1540_v29 = _nw209
				var _1541_v30 _dafny.Array
				_ = _1541_v30
				var _nwElement0_42 _dafny.Array = _1540_v29
				_ = _nwElement0_42
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(17))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_42, 0)
				(_nw210).ArraySet1(_1540_v29, 1)
				(_nw210).ArraySet1(_1540_v29, 2)
				(_nw210).ArraySet1(_1540_v29, 3)
				(_nw210).ArraySet1(_1540_v29, 4)
				(_nw210).ArraySet1(_1540_v29, 5)
				(_nw210).ArraySet1(_1540_v29, 6)
				(_nw210).ArraySet1(_1540_v29, 7)
				(_nw210).ArraySet1(_1540_v29, 8)
				(_nw210).ArraySet1(_1540_v29, 9)
				(_nw210).ArraySet1(_1540_v29, 10)
				(_nw210).ArraySet1(_1540_v29, 11)
				(_nw210).ArraySet1(_1540_v29, 12)
				(_nw210).ArraySet1(_1540_v29, 13)
				(_nw210).ArraySet1(_1540_v29, 14)
				(_nw210).ArraySet1(_1540_v29, 15)
				(_nw210).ArraySet1(_1540_v29, 16)
				_1541_v30 = _nw210
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_1541_v30), 0))
				_ = _index220
				(_1541_v30).ArraySet1(_1540_v29, (_index220).Int())
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_1541_v30), 0))
				_ = _index221
				(_1541_v30).ArraySet1(_1540_v29, (_index221).Int())
			}
			var _1542_v31 _dafny.Sequence
			_ = _1542_v31
			_1542_v31 = _dafny.SeqOf(p0)
			var _1543_v32 _dafny.Sequence
			_ = _1543_v32
			_1543_v32 = _dafny.SeqOf(p1)
			var _1544_v33 D0
			_ = _1544_v33
			_1544_v33 = Companion_D0_.Create_DC1_(!(p0), _1543_v32, p1)
			var _1545_v34 D0
			_ = _1545_v34
			_1545_v34 = Companion_D0_.Create_DC3_(_1544_v33)
			var _1546_v35 _dafny.Map
			_ = _1546_v35
			_1546_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1542_v31).Select((Companion_Default___.SafeIndex(_1532_v21, _dafny.IntOfUint32((_1542_v31).Cardinality()))).Uint32()).(bool), _1545_v34)
			var _1547_v36 D2
			_ = _1547_v36
			_1547_v36 = Companion_D2_.Create_DC7_(false, _1546_v35, _1532_v21, p0)
			var _1548_v37 _dafny.CodePoint
			_ = _1548_v37
			_1548_v37 = _dafny.CodePoint('u')
			var _1549_v38 _dafny.MultiSet
			_ = _1549_v38
			_1549_v38 = _dafny.MultiSetOf(_1548_v37)
			var _1550_v39 _dafny.MultiSet
			_ = _1550_v39
			_1550_v39 = _dafny.MultiSetOf(_dafny.MultiSetOf(!(p0)))
			var _1551_v40 _dafny.Sequence
			_ = _1551_v40
			_1551_v40 = _dafny.SeqOf(_1532_v21)
			var _1552_v41 _dafny.Map
			_ = _1552_v41
			_1552_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _1553_v42 _dafny.Sequence
			_ = _1553_v42
			_1553_v42 = _dafny.SeqOf(_1552_v41)
			var _1554_v43 _dafny.Array
			_ = _1554_v43
			var _nwElement0_43 _dafny.Int = _1532_v21
			_ = _nwElement0_43
			var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(27))
			_ = _nw211
			(_nw211).ArraySet1(_nwElement0_43, 0)
			(_nw211).ArraySet1((_1532_v21).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}(func(_1555_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality())), 1)
			(_nw211).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm22(_1547_v36, _1545_v34, p0, globalState)), 2)
			(_nw211).ArraySet1(_1532_v21, 3)
			(_nw211).ArraySet1(_1532_v21, 4)
			(_nw211).ArraySet1((func() _dafny.Int {
				if (_1549_v38).Contains(_1548_v37) {
					return (_1549_v38).Multiplicity(_1548_v37)
				}
				return _1532_v21
			})(), 5)
			(_nw211).ArraySet1(_1532_v21, 6)
			(_nw211).ArraySet1(_1532_v21, 7)
			(_nw211).ArraySet1(_1532_v21, 8)
			(_nw211).ArraySet1((func() _dafny.Int {
				if (_1550_v39).Contains(_1531_cf23) {
					return (_1550_v39).Multiplicity(_1531_cf23)
				}
				return _1532_v21
			})(), 9)
			(_nw211).ArraySet1(_1532_v21, 10)
			(_nw211).ArraySet1(_1532_v21, 11)
			(_nw211).ArraySet1(_1532_v21, 12)
			(_nw211).ArraySet1((_1551_v40).Select((Companion_Default___.SafeIndex(_1532_v21, _dafny.IntOfUint32((_1551_v40).Cardinality()))).Uint32()).(_dafny.Int), 13)
			(_nw211).ArraySet1((func() _dafny.Int {
				if p0 {
					return _1532_v21
				}
				return _1532_v21
			})(), 14)
			(_nw211).ArraySet1(_dafny.IntOfInt64(542), 15)
			(_nw211).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("da"))).Cardinality()), 16)
			(_nw211).ArraySet1((_dafny.Zero).Minus((_1532_v21).Plus(_dafny.IntOfInt64(471))), 17)
			(_nw211).ArraySet1(_1532_v21, 18)
			(_nw211).ArraySet1((func() _dafny.Int {
				if p0 {
					return _1532_v21
				}
				return _1532_v21
			})(), 19)
			(_nw211).ArraySet1(_1532_v21, 20)
			(_nw211).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1553_v42).Cardinality()), (_dafny.Zero).Minus(_1532_v21)), 21)
			(_nw211).ArraySet1((_dafny.Zero).Minus((_1532_v21).Plus(_1532_v21)), 22)
			(_nw211).ArraySet1(_1532_v21, 23)
			(_nw211).ArraySet1((_1532_v21).Plus(_1532_v21), 24)
			(_nw211).ArraySet1(_1532_v21, 25)
			(_nw211).ArraySet1(Companion_Default___.SafeDivisionInt(_1532_v21, (_dafny.Zero).Minus(_1532_v21)), 26)
			_1554_v43 = _nw211
			var _1556_v44 T0
			_ = _1556_v44
			var _nw212 *C2 = New_C2_()
			_ = _nw212
			_nw212.Ctor__()
			_1556_v44 = _nw212
			var _1557_v45 _dafny.Array
			_ = _1557_v45
			var _nw213 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
			_ = _nw213
			_1557_v45 = _nw213
			var _1558_v46 D0
			_ = _1558_v46
			_1558_v46 = Companion_D0_.Create_DC1_(p0, _1543_v32, p1)
			var _1559_v47 _dafny.Sequence
			_ = _1559_v47
			_1559_v47 = _dafny.SeqOf(_1558_v46)
			var _1560_v48 *C0
			_ = _1560_v48
			var _nw214 *C0 = New_C0_()
			_ = _nw214
			_nw214.Ctor__((_dafny.MultiSetFromSeq(_1559_v47)).Update(_1558_v46, Companion_Default___.Abs(_1532_v21)), p0)
			_1560_v48 = _nw214
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_1557_v45), 0))
			_ = _index222
			(_1557_v45).ArraySet1(_1560_v48, (_index222).Int())
			var _1561_v49 _dafny.Set
			_ = _1561_v49
			_1561_v49 = _dafny.SetOf((_1560_v48).F19())
			var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_1557_v45), 0))
			_ = _index223
			var _rhs249 bool = Companion_Default___.Fm19((func() _dafny.Int {
				if p0 {
					return _1532_v21
				}
				return (_1561_v49).Cardinality()
			})(), (_1560_v48).F19(), globalState)
			_ = _rhs249
			var _rhs250 _dafny.Array = (Companion_D13_.Create_DC33_(_1554_v43)).Dtor_cf51()
			_ = _rhs250
			var _rhs251 T0 = _1556_v44
			_ = _rhs251
			var _rhs252 *C0 = _1560_v48
			_ = _rhs252
			var _rhs253 bool = p0
			_ = _rhs253
			var _lhs163 *GlobalState = globalState
			_ = _lhs163
			var _lhs164 _dafny.Array = _1557_v45
			_ = _lhs164
			var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_1557_v45), 0))
			_ = _lhs165
			var _lhs166 *GlobalState = globalState
			_ = _lhs166
			_lhs163.F3 = _rhs249
			_1554_v43 = _rhs250
			_1556_v44 = _rhs251
			(_lhs164).ArraySet1(_rhs252, (_lhs165).Int())
			_lhs166.F5 = _rhs253
		} else {
			var _1562___mcc_h2 D6 = _source23.Get_().(D6_DC17).Cf25
			_ = _1562___mcc_h2
			var _1563_cf25 D6 = _1562___mcc_h2
			_ = _1563_cf25
			var _1564_v50 _dafny.Int
			_ = _1564_v50
			_1564_v50 = _dafny.IntOfInt64(377)
			var _1565_v51 _dafny.Set
			_ = _1565_v51
			_1565_v51 = _dafny.SetOf(p0, p0)
			var _1566_v52 _dafny.Map
			_ = _1566_v52
			_1566_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(p0))
			var _1567_v53 _dafny.Map
			_ = _1567_v53
			_1567_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1564_v50)
			var _1568_v54 _dafny.Set
			_ = _1568_v54
			_1568_v54 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1564_v50), _1567_v53, _1567_v53, _1567_v53)
			var _1569_v55 _dafny.Map
			_ = _1569_v55
			_1569_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1564_v50, p0)
			var _1570_v56 _dafny.Sequence
			_ = _1570_v56
			_1570_v56 = _dafny.SeqOf(_1564_v50, _1564_v50)
			var _1571_v57 _dafny.Array
			_ = _1571_v57
			var _nwElement0_44 _dafny.Int = (_1564_v50).Minus(_1564_v50)
			_ = _nwElement0_44
			var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(15))
			_ = _nw215
			(_nw215).ArraySet1(_nwElement0_44, 0)
			(_nw215).ArraySet1(_1564_v50, 1)
			(_nw215).ArraySet1((_1564_v50).Times(_1564_v50), 2)
			(_nw215).ArraySet1(_1564_v50, 3)
			(_nw215).ArraySet1(((_1565_v51).Difference(Companion_Default___.Fm33(p0, p0, !(p0), globalState))).Cardinality(), 4)
			(_nw215).ArraySet1((_1566_v52).Cardinality(), 5)
			(_nw215).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())), 6)
			(_nw215).ArraySet1(Companion_Default___.Fm11(p0, globalState), 7)
			(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()), 8)
			(_nw215).ArraySet1(Companion_Default___.SafeDivisionInt(_1564_v50, _1564_v50), 9)
			(_nw215).ArraySet1((_dafny.IntOfInt64(124)).Times(_1564_v50), 10)
			(_nw215).ArraySet1(((_1568_v54).Union(_1568_v54)).Cardinality(), 11)
			(_nw215).ArraySet1(((_1569_v55).Merge(_1569_v55)).Cardinality(), 12)
			(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1570_v56, _1570_v56, _1570_v56, _1570_v56)).Cardinality()), 13)
			(_nw215).ArraySet1(_1564_v50, 14)
			_1571_v57 = _nw215
			_1571_v57 = _1571_v57
			r3 = p0
			_1564_v50 = (_dafny.Zero).Minus(_1564_v50)
			var _1572_v58 _dafny.Sequence
			_ = _1572_v58
			_1572_v58 = _dafny.UnicodeSeqOfUtf8Bytes("vcvnthiyb")
			_1572_v58 = p1
		}
		if p0 {
			var _1573_v59 _dafny.Array
			_ = _1573_v59
			var _nw216 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw216
			_1573_v59 = _nw216
			var _1574_v60 _dafny.Set
			_ = _1574_v60
			_1574_v60 = _dafny.SetOf(_1573_v59, _1573_v59, _1573_v59, _1573_v59, _1573_v59)
			_1574_v60 = _1574_v60
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _index224
			(_1573_v59).ArraySet1(false, (_index224).Int())
			var _1575_v61 _dafny.Int
			_ = _1575_v61
			_1575_v61 = _dafny.IntOfInt64(330)
			var _1576_v62 _dafny.Map
			_ = _1576_v62
			_1576_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-713))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}((func(_1577_v61 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1578_i2 _dafny.Int) _dafny.Int {
					return _1577_v61
				}
			})(_1575_v61)))).Cardinality()), _1575_v61)
			var _1579_v63 _dafny.Array
			_ = _1579_v63
			var _nwElement0_45 _dafny.Int = (_dafny.Zero).Minus(_1575_v61)
			_ = _nwElement0_45
			var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(24))
			_ = _nw217
			(_nw217).ArraySet1(_nwElement0_45, 0)
			(_nw217).ArraySet1(_1575_v61, 1)
			(_nw217).ArraySet1(_1575_v61, 2)
			(_nw217).ArraySet1(_1575_v61, 3)
			(_nw217).ArraySet1((_dafny.Zero).Minus(_1575_v61), 4)
			(_nw217).ArraySet1(_1575_v61, 5)
			(_nw217).ArraySet1(_1575_v61, 6)
			(_nw217).ArraySet1(_1575_v61, 7)
			(_nw217).ArraySet1(_dafny.IntOfInt64(440), 8)
			(_nw217).ArraySet1((func() _dafny.Int {
				if (_1576_v62).Contains(_1575_v61) {
					return (_1576_v62).Get(_1575_v61).(_dafny.Int)
				}
				return _1575_v61
			})(), 9)
			(_nw217).ArraySet1(_1575_v61, 10)
			(_nw217).ArraySet1(_1575_v61, 11)
			(_nw217).ArraySet1(_1575_v61, 12)
			(_nw217).ArraySet1(_1575_v61, 13)
			(_nw217).ArraySet1(_1575_v61, 14)
			(_nw217).ArraySet1(_dafny.IntOfInt64(688), 15)
			(_nw217).ArraySet1(_1575_v61, 16)
			(_nw217).ArraySet1(_1575_v61, 17)
			(_nw217).ArraySet1(_1575_v61, 18)
			(_nw217).ArraySet1(Companion_Default___.Fm11(p0, globalState), 19)
			(_nw217).ArraySet1((func() _dafny.Int {
				if (_1505_v0).Contains(p0) {
					return (_1505_v0).Multiplicity(p0)
				}
				return _1575_v61
			})(), 20)
			(_nw217).ArraySet1(_1575_v61, 21)
			(_nw217).ArraySet1(_1575_v61, 22)
			(_nw217).ArraySet1((_1574_v60).Cardinality(), 23)
			_1579_v63 = _nw217
			var _1580_v64 _dafny.Sequence
			_ = _1580_v64
			_1580_v64 = _dafny.SeqOf(_1579_v63)
			var _1581_v65 _dafny.Map
			_ = _1581_v65
			_1581_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1582_v66 _dafny.Sequence
			_ = _1582_v66
			_1582_v66 = _dafny.SeqOf(p0, p0, p0)
			var _1583_v67 _dafny.Array
			_ = _1583_v67
			var _nwElement0_46 _dafny.Int = (_1575_v61).Plus(_dafny.IntOfInt64(233))
			_ = _nwElement0_46
			var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(14))
			_ = _nw218
			(_nw218).ArraySet1(_nwElement0_46, 0)
			(_nw218).ArraySet1((_1575_v61).Times(_dafny.IntOfUint32((_1580_v64).Cardinality())), 1)
			(_nw218).ArraySet1(_1575_v61, 2)
			(_nw218).ArraySet1(_1575_v61, 3)
			(_nw218).ArraySet1(_1575_v61, 4)
			(_nw218).ArraySet1(_1575_v61, 5)
			(_nw218).ArraySet1(_1575_v61, 6)
			(_nw218).ArraySet1(_1575_v61, 7)
			(_nw218).ArraySet1(_1575_v61, 8)
			(_nw218).ArraySet1(_1575_v61, 9)
			(_nw218).ArraySet1((_1581_v65).Cardinality(), 10)
			(_nw218).ArraySet1((_1575_v61).Plus(_1575_v61), 11)
			(_nw218).ArraySet1((_1575_v61).Times(_dafny.IntOfUint32((_1582_v66).Cardinality())), 12)
			(_nw218).ArraySet1(_1575_v61, 13)
			_1583_v67 = _nw218
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))
			_ = _index225
			(_1583_v67).ArraySet1(_1575_v61, (_index225).Int())
			var _1584_v68 _dafny.Sequence
			_ = _1584_v68
			_1584_v68 = _dafny.SeqOf(_1575_v61, _1575_v61, _1575_v61, _1575_v61)
			var _1585_v69 *C5
			_ = _1585_v69
			var _nw219 *C5 = New_C5_()
			_ = _nw219
			_nw219.Ctor__()
			_1585_v69 = _nw219
			var _1586_v70 _dafny.Map
			_ = _1586_v70
			_1586_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1585_v69)
			var _1587_v71 _dafny.Sequence
			_ = _1587_v71
			_1587_v71 = _dafny.SeqOf((_1586_v70).Update(false, _1585_v69), _1586_v70, _1586_v70, _1586_v70)
			var _1588_v72 _dafny.Sequence
			_ = _1588_v72
			_1588_v72 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}(func(_1589_i3 _dafny.Int) _dafny.Int {
				return (Companion_D4_.Create_DC12_(_dafny.IntOfInt64(725))).Dtor_cf18()
			})), _dafny.Companion_Sequence_.Update(_1584_v68, (Companion_Default___.SafeIndex(_1575_v61, _dafny.IntOfUint32((_1584_v68).Cardinality()))).Uint32(), _dafny.IntOfInt64(173)), _1584_v68, (func() _dafny.Sequence {
				if (_1582_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1587_v71).Cardinality()), _dafny.IntOfUint32((_1582_v66).Cardinality()))).Uint32()).(bool) {
					return _1584_v68
				}
				return _1584_v68
			})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(178))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}((func(_1590_v62 _dafny.Map) func(_dafny.Int) _dafny.Int {
				return func(_1591_i4 _dafny.Int) _dafny.Int {
					return (_1590_v62).Cardinality()
				}
			})(_1576_v62))), Companion_Default___.Fm30(globalState)))
			var _1592_v73 T0
			_ = _1592_v73
			var _nw220 *C2 = New_C2_()
			_ = _nw220
			_nw220.Ctor__()
			_1592_v73 = _nw220
			var _1593_v74 _dafny.Map
			_ = _1593_v74
			_1593_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1592_v73, _1584_v68)
			var _1594_v75 _dafny.Map
			_ = _1594_v75
			_1594_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1583_v67, (_1593_v74).Cardinality())
			var _1595_v76 _dafny.MultiSet
			_ = _1595_v76
			_1595_v76 = _dafny.MultiSetOf(_1575_v61)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _index226
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))
			_ = _index227
			var _rhs254 _dafny.Int = _1575_v61
			_ = _rhs254
			var _rhs255 bool = p0
			_ = _rhs255
			var _rhs256 _dafny.Int = (func() _dafny.Int {
				if (_1594_v75).Contains(_1583_v67) {
					return (_1594_v75).Get(_1583_v67).(_dafny.Int)
				}
				return ((_1595_v76).Cardinality()).Times(_dafny.IntOfInt64(623))
			})()
			_ = _rhs256
			var _rhs257 _dafny.Int = (_1581_v65).Cardinality()
			_ = _rhs257
			var _rhs258 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}((func(_1596_v68 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1597_i5 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_1596_v68, _1596_v68)
				}
			})(_1584_v68)))
			_ = _rhs258
			var _lhs167 _dafny.Array = _1573_v59
			_ = _lhs167
			var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _lhs168
			var _lhs169 _dafny.Array = _1583_v67
			_ = _lhs169
			var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))
			_ = _lhs170
			r2 = _rhs254
			(_lhs167).ArraySet1(_rhs255, (_lhs168).Int())
			r2 = _rhs256
			(_lhs169).ArraySet1(_rhs257, (_lhs170).Int())
			_1588_v72 = _rhs258
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _index228
			(_1573_v59).ArraySet1(!(p0), (_index228).Int())
			if !(false) {
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
				_ = _index229
				(_1573_v59).ArraySet1((_1582_v66).Select((Companion_Default___.SafeIndex(_1575_v61, _dafny.IntOfUint32((_1582_v66).Cardinality()))).Uint32()).(bool), (_index229).Int())
				_1575_v61 = _dafny.IntOfInt64(653)
				r1 = p0
				var _1598_v77 _dafny.Array
				_ = _1598_v77
				var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw221
				_1598_v77 = _nw221
				var _1599_v78 _dafny.Sequence
				_ = _1599_v78
				_1599_v78 = _dafny.SeqOf(_1598_v77, _1598_v77, _1598_v77)
				var _1600_v79 _dafny.Array
				_ = _1600_v79
				var _nwElement0_47 _dafny.Array = _1598_v77
				_ = _nwElement0_47
				var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(27))
				_ = _nw222
				(_nw222).ArraySet1(_nwElement0_47, 0)
				(_nw222).ArraySet1(_1598_v77, 1)
				(_nw222).ArraySet1(_1598_v77, 2)
				(_nw222).ArraySet1(_1598_v77, 3)
				(_nw222).ArraySet1(_1598_v77, 4)
				(_nw222).ArraySet1(_1598_v77, 5)
				(_nw222).ArraySet1(_1598_v77, 6)
				(_nw222).ArraySet1(_1598_v77, 7)
				(_nw222).ArraySet1(_1598_v77, 8)
				(_nw222).ArraySet1(_1598_v77, 9)
				(_nw222).ArraySet1(_1598_v77, 10)
				(_nw222).ArraySet1(_1598_v77, 11)
				(_nw222).ArraySet1(_1598_v77, 12)
				(_nw222).ArraySet1(_1598_v77, 13)
				(_nw222).ArraySet1(_1598_v77, 14)
				(_nw222).ArraySet1(_1598_v77, 15)
				(_nw222).ArraySet1(_1598_v77, 16)
				(_nw222).ArraySet1(_1598_v77, 17)
				(_nw222).ArraySet1(_1598_v77, 18)
				(_nw222).ArraySet1((_1599_v78).Select((Companion_Default___.SafeIndex((_1583_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1599_v78).Cardinality()))).Uint32()).(_dafny.Array), 19)
				(_nw222).ArraySet1(_1598_v77, 20)
				(_nw222).ArraySet1(_1598_v77, 21)
				(_nw222).ArraySet1(_1598_v77, 22)
				(_nw222).ArraySet1(_1598_v77, 23)
				(_nw222).ArraySet1(_1598_v77, 24)
				(_nw222).ArraySet1(_1598_v77, 25)
				(_nw222).ArraySet1(_1598_v77, 26)
				_1600_v79 = _nw222
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1600_v79), 0))
				_ = _index230
				(_1600_v79).ArraySet1(_1598_v77, (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1600_v79), 0))
				_ = _index231
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
				_ = _index232
				var _rhs259 _dafny.Array = _1598_v77
				_ = _rhs259
				var _rhs260 bool = (func() bool {
					if !(p0) {
						return ((_1505_v0).Update(p0, Companion_Default___.Abs(_1575_v61))).IsSubsetOf(_1505_v0)
					}
					return p0
				})()
				_ = _rhs260
				var _lhs171 _dafny.Array = _1600_v79
				_ = _lhs171
				var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1600_v79), 0))
				_ = _lhs172
				var _lhs173 _dafny.Array = _1573_v59
				_ = _lhs173
				var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
				_ = _lhs174
				(_lhs171).ArraySet1(_rhs259, (_lhs172).Int())
				(_lhs173).ArraySet1(_rhs260, (_lhs174).Int())
				r2 = _1575_v61
			} else {
				var _1601_v80 _dafny.Set
				_ = _1601_v80
				_1601_v80 = _dafny.SetOf(p0)
				r0 = ((_1601_v80).Intersection(Companion_Default___.Fm33(p0, (_1573_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))).Int()).(bool), (_1573_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))).Int()).(bool), globalState))).Union(_1601_v80)
				var _1602_v81 D0
				_ = _1602_v81
				_1602_v81 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_1582_v66).Cardinality()))
				var _1603_v82 _dafny.Array
				_ = _1603_v82
				var _nwElement0_48 D0 = _1602_v81
				_ = _nwElement0_48
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(4))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_48, 0)
				(_nw223).ArraySet1(_1602_v81, 1)
				(_nw223).ArraySet1(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((p1).Cardinality())), 2)
				(_nw223).ArraySet1(_1602_v81, 3)
				_1603_v82 = _nw223
				var _1604_v83 D7
				_ = _1604_v83
				_1604_v83 = Companion_D7_.Create_DC20_(p1, Companion_Default___.Fm2(_dafny.CodePoint('u'), Companion_Default___.Fm28(globalState), (_1583_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))).Int()).(_dafny.Int), globalState))
				var _1605_v84 _dafny.Sequence
				_ = _1605_v84
				_1605_v84 = _dafny.SeqOf(p1)
				var _1606_v85 _dafny.Map
				_ = _1606_v85
				_1606_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bjx"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(416))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}((func(_1607_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1608_i6 _dafny.Int) _dafny.Sequence {
						return _1607_p1
					}
				})(p1))), _1605_v84))
				var _1609_v86 _dafny.Set
				_ = _1609_v86
				_1609_v86 = _dafny.SetOf(_1575_v61, _1575_v61)
				var _1610_v87 _dafny.Map
				_ = _1610_v87
				_1610_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1609_v86)
				var _1611_v88 D4
				_ = _1611_v88
				_1611_v88 = Companion_D4_.Create_DC10_((func() _dafny.Set {
					if (_1610_v87).Contains(p1) {
						return (_1610_v87).Get(p1).(_dafny.Set)
					}
					return _1609_v86
				})())
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))
				_ = _index233
				var _rhs261 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1606_v85).Contains(p1) {
						return (_1606_v85).Get(p1).(_dafny.Sequence)
					}
					return _1605_v84
				})()).Cardinality())
				_ = _rhs261
				var _rhs262 _dafny.Array = _1603_v82
				_ = _rhs262
				var _rhs263 _dafny.Int = (_1575_v61).Times((_dafny.Zero).Minus(_1575_v61))
				_ = _rhs263
				var _rhs264 D7 = Companion_Default___.Fm40(_1575_v61, _1611_v88, globalState)
				_ = _rhs264
				var _rhs265 _dafny.Int = (_1575_v61).Plus(_1575_v61)
				_ = _rhs265
				var _lhs175 _dafny.Array = _1583_v67
				_ = _lhs175
				var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))
				_ = _lhs176
				(_lhs175).ArraySet1(_rhs261, (_lhs176).Int())
				_1603_v82 = _rhs262
				_1575_v61 = _rhs263
				_1604_v83 = _rhs264
				r2 = _rhs265
				var _1612_v89 *C3
				_ = _1612_v89
				var _nw224 *C3 = New_C3_()
				_ = _nw224
				_nw224.Ctor__((_1575_v61).Cmp(_1575_v61) <= 0, (_1583_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_1583_v67), 0))).Int()).(_dafny.Int))
				_1612_v89 = _nw224
				var _1613_v90 _dafny.Set
				_ = _1613_v90
				_1613_v90 = _dafny.SetOf(_1604_v83)
				var _1614_v91 _dafny.Sequence
				_ = _1614_v91
				_1614_v91 = _dafny.SeqOf(Companion_D7_.Create_DC20_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(965))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}(func(_1615_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				})), p0))
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
				_ = _index234
				(_1573_v59).ArraySet1((func() bool {
					if (_1573_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))).Int()).(bool) {
						return (_1613_v90).IsDisjointFrom(func() _dafny.Set {
							var _coll42 = _dafny.NewBuilder()
							_ = _coll42
							for _iter44 := _dafny.Iterate((_1614_v91).Elements()); ; {
								_compr_42, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _1616_v92 D7
								_1616_v92 = interface{}(_compr_42).(D7)
								if _dafny.Companion_Sequence_.Contains(_1614_v91, _1616_v92) {
									_coll42.Add(_1616_v92)
								}
							}
							return _coll42.ToSet()
						}())
					}
					return p0
				})(), (_index234).Int())
				var _1617_v93 _dafny.Array
				_ = _1617_v93
				var _nw225 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw225
				_1617_v93 = _nw225
				_1617_v93 = _1617_v93
			}
			var _1618_v94 D13
			_ = _1618_v94
			_1618_v94 = Companion_D13_.Create_DC33_(_1579_v63)
			var _1619_v95 _dafny.Array
			_ = _1619_v95
			var _nwElement0_49 _dafny.Array = _1583_v67
			_ = _nwElement0_49
			var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(11))
			_ = _nw226
			(_nw226).ArraySet1(_nwElement0_49, 0)
			(_nw226).ArraySet1(_1583_v67, 1)
			(_nw226).ArraySet1(_1583_v67, 2)
			(_nw226).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1583_v67
				}
				return _1579_v63
			})(), 3)
			(_nw226).ArraySet1((_1580_v64).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-168)), _dafny.IntOfUint32((_1580_v64).Cardinality()))).Uint32()).(_dafny.Array), 4)
			(_nw226).ArraySet1((_1618_v94).Dtor_cf51(), 5)
			(_nw226).ArraySet1(_1579_v63, 6)
			(_nw226).ArraySet1(_1579_v63, 7)
			(_nw226).ArraySet1(_1579_v63, 8)
			(_nw226).ArraySet1(_1579_v63, 9)
			(_nw226).ArraySet1(_1579_v63, 10)
			_1619_v95 = _nw226
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1619_v95), 0))
			_ = _index235
			(_1619_v95).ArraySet1(_1583_v67, (_index235).Int())
			var _1620_v96 D5
			_ = _1620_v96
			_1620_v96 = Companion_D5_.Create_DC13_(_1573_v59)
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1619_v95), 0))
			_ = _index236
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _index237
			var _rhs266 bool = !(p0)
			_ = _rhs266
			var _rhs267 _dafny.Array = _1583_v67
			_ = _rhs267
			var _rhs268 bool = p0
			_ = _rhs268
			var _rhs269 bool = (func() bool {
				if p0 {
					return p0
				}
				return p0
			})()
			_ = _rhs269
			var _rhs270 _dafny.Array = (_1620_v96).Dtor_cf19()
			_ = _rhs270
			var _lhs177 *GlobalState = globalState
			_ = _lhs177
			var _lhs178 _dafny.Array = _1619_v95
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1619_v95), 0))
			_ = _lhs179
			var _lhs180 *GlobalState = globalState
			_ = _lhs180
			var _lhs181 _dafny.Array = _1573_v59
			_ = _lhs181
			var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1573_v59), 0))
			_ = _lhs182
			_lhs177.F13 = _rhs266
			(_lhs178).ArraySet1(_rhs267, (_lhs179).Int())
			_lhs180.F9 = _rhs268
			(_lhs181).ArraySet1(_rhs269, (_lhs182).Int())
			_1573_v59 = _rhs270
		} else {
			var _1621_v97 _dafny.Set
			_ = _1621_v97
			_1621_v97 = _dafny.SetOf(!(false), p0)
			r0 = (_1621_v97).Intersection(_dafny.SetOf(p0, p0, p0))
			var _1622_v98 _dafny.Array
			_ = _1622_v98
			var _nw227 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw227
			_1622_v98 = _nw227
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1622_v98), 0))
			_ = _index238
			(_1622_v98).ArraySet1(true, (_index238).Int())
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1622_v98), 0))
			_ = _index239
			(_1622_v98).ArraySet1(p0, (_index239).Int())
			var _1623_v99 _dafny.Int
			_ = _1623_v99
			_1623_v99 = _dafny.IntOfInt64(564)
			r3 = (_dafny.IntOfInt64(836)).Cmp(Companion_Default___.SafeDivisionInt(_1623_v99, _1623_v99)) <= 0
			(globalState).F9 = !(p0)
			var _1624_v100 _dafny.Sequence
			_ = _1624_v100
			_1624_v100 = _dafny.UnicodeSeqOfUtf8Bytes("bspwfx")
			_1624_v100 = _dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Concatenate(p1, _1624_v100))
		}
		(globalState).F9 = Companion_Default___.Fm19(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gb")).Cardinality()), !(!(p0)), globalState)
		if p0 {
			var _1625_v101 _dafny.CodePoint
			_ = _1625_v101
			_1625_v101 = _dafny.CodePoint('s')
			_1625_v101 = _1625_v101
			var _1626_v102 _dafny.Sequence
			_ = _1626_v102
			_1626_v102 = _dafny.UnicodeSeqOfUtf8Bytes("xmcve")
			_1626_v102 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-191))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg112 _dafny.Int) interface{} {
					return coer112(arg112)
				}
			}((func(_1627_v101 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1628_i8 _dafny.Int) _dafny.CodePoint {
					return _1627_v101
				}
			})(_1625_v101)))
			var _1629_v103 _dafny.Sequence
			_ = _1629_v103
			_1629_v103 = _dafny.SeqOf(p0)
			var _1630_v104 _dafny.Set
			_ = _1630_v104
			_1630_v104 = _dafny.SetOf(_1629_v103, _1629_v103)
			r2 = ((_1630_v104).Union(_1630_v104)).Cardinality()
			var _1631_v105 _dafny.Int
			_ = _1631_v105
			_1631_v105 = _dafny.IntOfInt64(252)
			r2 = (_dafny.Zero).Minus(_1631_v105)
			var _1632_v106 D9
			_ = _1632_v106
			_1632_v106 = Companion_D9_.Create_DC26_(p0, _1625_v101, p0)
			var _1633_v107 _dafny.Map
			_ = _1633_v107
			_1633_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1634_v108 _dafny.Map
			_ = _1634_v108
			_1634_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1625_v101, _1633_v107)
			var _pat_let_tv29 = p0
			_ = _pat_let_tv29
			var _pat_let_tv30 = p0
			_ = _pat_let_tv30
			var _pat_let_tv31 = _1634_v108
			_ = _pat_let_tv31
			var _pat_let_tv32 = p0
			_ = _pat_let_tv32
			var _pat_let_tv33 = _1631_v105
			_ = _pat_let_tv33
			var _pat_let_tv34 = globalState
			_ = _pat_let_tv34
			var _1635_v109 _dafny.Array
			_ = _1635_v109
			var _nwElement0_50 _dafny.CodePoint = _dafny.CodePoint('b')
			_ = _nwElement0_50
			var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(9))
			_ = _nw228
			(_nw228).ArraySet1CodePoint(_nwElement0_50, 0)
			(_nw228).ArraySet1CodePoint(_1625_v101, 1)
			(_nw228).ArraySet1CodePoint(_dafny.CodePoint('u'), 2)
			(_nw228).ArraySet1CodePoint(_1625_v101, 3)
			(_nw228).ArraySet1CodePoint(_1625_v101, 4)
			(_nw228).ArraySet1CodePoint((func(_pat_let33_0 D9) D9 {
				return func(_1636_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let34_0 bool) D9 {
						return func(_1637_dt__update_hcf41_h0 bool) D9 {
							return func(_pat_let35_0 _dafny.CodePoint) D9 {
								return func(_1638_dt__update_hcf40_h0 _dafny.CodePoint) D9 {
									return Companion_D9_.Create_DC26_((_1636_dt__update__tmp_h0).Dtor_cf39(), _1638_dt__update_hcf40_h0, _1637_dt__update_hcf41_h0)
								}(_pat_let35_0)
							}(Companion_Default___.Fm23(_pat_let_tv30, _pat_let_tv31, _pat_let_tv32, _pat_let_tv33, _pat_let_tv34))
						}(_pat_let34_0)
					}(_pat_let_tv29)
				}(_pat_let33_0)
			}(_1632_v106)).Dtor_cf40(), 5)
			(_nw228).ArraySet1CodePoint(_1625_v101, 6)
			(_nw228).ArraySet1CodePoint(_dafny.CodePoint('s'), 7)
			(_nw228).ArraySet1CodePoint(_1625_v101, 8)
			_1635_v109 = _nw228
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1635_v109), 0))
			_ = _index240
			(_1635_v109).ArraySet1CodePoint(_dafny.CodePoint('w'), (_index240).Int())
			var _1639_v110 _dafny.Set
			_ = _1639_v110
			_1639_v110 = _dafny.SetOf(true)
			var _1640_v111 D7
			_ = _1640_v111
			_1640_v111 = Companion_D7_.Create_DC20_(_1626_v102, p0)
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1635_v109), 0))
			_ = _index241
			var _rhs271 bool = (Companion_Default___.Fm41(p0, true, globalState)).Dtor_cf39()
			_ = _rhs271
			var _rhs272 _dafny.CodePoint = Companion_Default___.Fm23((_1639_v110).IsProperSubsetOf(_dafny.SetOf(p0, p0)), _1634_v108, (_1640_v111).Dtor_cf32(), _1631_v105, globalState)
			_ = _rhs272
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			var _lhs184 _dafny.Array = _1635_v109
			_ = _lhs184
			var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1635_v109), 0))
			_ = _lhs185
			_lhs183.F5 = _rhs271
			(_lhs184).ArraySet1CodePoint(_rhs272, (_lhs185).Int())
		} else {
			var _1641_v112 _dafny.Int
			_ = _1641_v112
			_1641_v112 = _dafny.IntOfInt64(579)
			r2 = _1641_v112
			if p0 {
				var _1642_v113 _dafny.Array
				_ = _1642_v113
				var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw229
				_1642_v113 = _nw229
				var _1643_v114 _dafny.Map
				_ = _1643_v114
				_1643_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1642_v113, p0)
				_1643_v114 = _1643_v114
				var _1644_v115 _dafny.Array
				_ = _1644_v115
				var _nw230 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw230
				_1644_v115 = _nw230
				var _1645_v116 _dafny.Sequence
				_ = _1645_v116
				_1645_v116 = _dafny.SeqOf(p0)
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1644_v115), 0))
				_ = _index242
				(_1644_v115).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1645_v116, p0), (_index242).Int())
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1644_v115), 0))
				_ = _index243
				var _rhs273 bool = p0
				_ = _rhs273
				var _rhs274 _dafny.Int = _1641_v112
				_ = _rhs274
				var _rhs275 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((p1).Cardinality()), _1641_v112)
				_ = _rhs275
				var _lhs186 _dafny.Array = _1644_v115
				_ = _lhs186
				var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1644_v115), 0))
				_ = _lhs187
				(_lhs186).ArraySet1(_rhs273, (_lhs187).Int())
				r2 = _rhs274
				r2 = _rhs275
				var _1646_v117 _dafny.Array
				_ = _1646_v117
				var _nw231 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
				_ = _nw231
				_1646_v117 = _nw231
				var _1647_v118 *C2
				_ = _1647_v118
				var _nw232 *C2 = New_C2_()
				_ = _nw232
				_nw232.Ctor__()
				_1647_v118 = _nw232
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1646_v117), 0))
				_ = _index244
				(_1646_v117).ArraySet1(_1647_v118, (_index244).Int())
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1646_v117), 0))
				_ = _index245
				(_1646_v117).ArraySet1(_1647_v118, (_index245).Int())
				var _1648_v119 _dafny.Map
				_ = _1648_v119
				_1648_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1644_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1644_v115), 0))).Int()).(bool)), _1642_v113)
				var _1649_v120 _dafny.Set
				_ = _1649_v120
				_1649_v120 = _dafny.SetOf((_1641_v112).Minus(_1641_v112), _dafny.IntOfInt64(38), (_1648_v119).Cardinality())
				_1649_v120 = _1649_v120
				(globalState).F9 = (_1644_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1644_v115), 0))).Int()).(bool)
			} else {
				var _1650_v121 _dafny.CodePoint
				_ = _1650_v121
				_1650_v121 = _dafny.CodePoint('x')
				var _1651_v122 _dafny.Map
				_ = _1651_v122
				_1651_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1650_v121, _1641_v112)
				var _1652_v123 *C4
				_ = _1652_v123
				var _nw233 *C4 = New_C4_()
				_ = _nw233
				_nw233.Ctor__(_1641_v112, (_1651_v122).Merge(_1651_v122))
				_1652_v123 = _nw233
				var _1653_v124 _dafny.Map
				_ = _1653_v124
				_1653_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _1641_v112)
				var _1654_v125 _dafny.MultiSet
				_ = _1654_v125
				_1654_v125 = _dafny.MultiSetOf(_dafny.CodePoint('o'), _1650_v121)
				var _1655_v126 _dafny.Array
				_ = _1655_v126
				var _nwElement0_51 _dafny.Int = (_1652_v123).F14()
				_ = _nwElement0_51
				var _nw234 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(17))
				_ = _nw234
				(_nw234).ArraySet1(_nwElement0_51, 0)
				(_nw234).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).Fm1(_1653_v124, globalState), (Companion_Default___.SafeIndex(_1641_v112, _dafny.IntOfUint32(((_this).Fm1(_1653_v124, globalState)).Cardinality()))).Uint32(), _1650_v121)).Cardinality()), 1)
				(_nw234).ArraySet1((_1652_v123).F14(), 2)
				(_nw234).ArraySet1((_1652_v123).F14(), 3)
				(_nw234).ArraySet1(((_1654_v125).Update(_1650_v121, Companion_Default___.Abs((_1652_v123).F14()))).Cardinality(), 4)
				(_nw234).ArraySet1((_1652_v123).F14(), 5)
				(_nw234).ArraySet1(_1641_v112, 6)
				(_nw234).ArraySet1(_1641_v112, 7)
				(_nw234).ArraySet1((_1652_v123).F14(), 8)
				(_nw234).ArraySet1((_1652_v123).F14(), 9)
				(_nw234).ArraySet1(_dafny.IntOfInt64(-354), 10)
				(_nw234).ArraySet1((_1652_v123).F14(), 11)
				(_nw234).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(148))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_1656_v123 *C4) func(_dafny.Int) _dafny.Int {
					return func(_1657_i9 _dafny.Int) _dafny.Int {
						return (_1656_v123).F14()
					}
				})(_1652_v123)))).Cardinality()), 12)
				(_nw234).ArraySet1(_1641_v112, 13)
				(_nw234).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 14)
				(_nw234).ArraySet1(_1641_v112, 15)
				(_nw234).ArraySet1((_1652_v123).F14(), 16)
				_1655_v126 = _nw234
				var _1658_v127 _dafny.Sequence
				_ = _1658_v127
				_1658_v127 = _dafny.SeqOf(_1655_v126)
				var _1659_v128 _dafny.Set
				_ = _1659_v128
				_1659_v128 = _dafny.SetOf(!(p0))
				var _1660_v129 _dafny.Map
				_ = _1660_v129
				_1660_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p0, true), (_1659_v128).IsDisjointFrom(_1659_v128))
				var _rhs276 _dafny.Int = Companion_Default___.SafeModuloInt((_1652_v123).F14(), (_1652_v123).F14())
				_ = _rhs276
				var _rhs277 bool = (func() bool {
					if (_1660_v129).Contains(_1505_v0) {
						return (_1660_v129).Get(_1505_v0).(bool)
					}
					return p0
				})()
				_ = _rhs277
				var _rhs278 _dafny.Sequence = _1658_v127
				_ = _rhs278
				var _lhs188 *GlobalState = globalState
				_ = _lhs188
				r2 = _rhs276
				_lhs188.F9 = _rhs277
				_1658_v127 = _rhs278
				var _1661_v130 _dafny.Map
				_ = _1661_v130
				_1661_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_1661_v130 = _1661_v130
				var _1662_v131 _dafny.Sequence
				_ = _1662_v131
				_1662_v131 = _dafny.SeqOf(p0, p0)
				var _1663_v132 *C0
				_ = _1663_v132
				var _nw235 *C0 = New_C0_()
				_ = _nw235
				_nw235.Ctor__(Companion_Default___.Fm42(p0, _dafny.IntOfUint32((_1662_v131).Cardinality()), (_1652_v123).F15(), _dafny.SeqOf(p0), globalState), (!(p0)) && (true))
				_1663_v132 = _nw235
				_1663_v132 = _1663_v132
				r1 = p0
			}
			var _1664_v133 _dafny.Array
			_ = _1664_v133
			var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
			_ = _nw236
			_1664_v133 = _nw236
			var _1665_v134 _dafny.Map
			_ = _1665_v134
			_1665_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1666_v135 _dafny.Sequence
			_ = _1666_v135
			_1666_v135 = _dafny.SeqOf(_1665_v134)
			var _1667_v136 _dafny.MultiSet
			_ = _1667_v136
			_1667_v136 = _dafny.MultiSetOf(_1665_v134, (_1666_v135).Select((Companion_Default___.SafeIndex(_1641_v112, _dafny.IntOfUint32((_1666_v135).Cardinality()))).Uint32()).(_dafny.Map))
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_1664_v133), 0))
			_ = _index246
			(_1664_v133).ArraySet1(_1667_v136, (_index246).Int())
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_1664_v133), 0))
			_ = _index247
			(_1664_v133).ArraySet1((_1667_v136).Union(_dafny.MultiSetOf(_1665_v134, _1665_v134)), (_index247).Int())
			var _1668_v137 *C1
			_ = _1668_v137
			var _nw237 *C1 = New_C1_()
			_ = _nw237
			_nw237.Ctor__()
			_1668_v137 = _nw237
			var _1669_v138 _dafny.Array
			_ = _1669_v138
			var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw238
			_1669_v138 = _nw238
			_1669_v138 = _1669_v138
		}
		var _1670_v139 _dafny.Array
		_ = _1670_v139
		var _nw239 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw239
		_1670_v139 = _nw239
		var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))
		_ = _index248
		(_1670_v139).ArraySet1(!(p0), (_index248).Int())
		var _1671_v140 _dafny.Map
		_ = _1671_v140
		_1671_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1672_v141 _dafny.Int
		_ = _1672_v141
		_1672_v141 = _dafny.IntOfInt64(18)
		var _1673_v142 _dafny.Sequence
		_ = _1673_v142
		_1673_v142 = _dafny.SeqOf(p0, p0, p0, false, true)
		var _1674_v143 D0
		_ = _1674_v143
		_1674_v143 = Companion_D0_.Create_DC2_(!((_1673_v142).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1673_v142).Cardinality()))).Uint32()).(bool)))
		var _1675_v144 D0
		_ = _1675_v144
		_1675_v144 = Companion_D0_.Create_DC3_(_1674_v143)
		var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))
		_ = _index249
		(_1670_v139).ArraySet1((func() bool {
			if (_1671_v140).Contains((func() bool {
				if p0 {
					return p0
				}
				return p0
			})()) {
				return (_1671_v140).Get((func() bool {
					if p0 {
						return p0
					}
					return p0
				})()).(bool)
			}
			return Companion_Default___.Fm2((p1).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint), _1675_v144, _1672_v141, globalState)
		})(), (_index249).Int())
		if (_dafny.IntOfInt64(210)).Cmp(Companion_Default___.SafeDivisionInt(_1672_v141, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1673_v142, (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1673_v142).Cardinality()))).Uint32(), p0)).Cardinality()))) >= 0 {
			var _1676_v145 _dafny.Sequence
			_ = _1676_v145
			_1676_v145 = _dafny.SeqOf(_1670_v139, _1670_v139)
			var _1677_v146 _dafny.Map
			_ = _1677_v146
			_1677_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), _1670_v139)
			var _1678_v147 _dafny.Array
			_ = _1678_v147
			var _nwElement0_52 _dafny.Array = _1670_v139
			_ = _nwElement0_52
			var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(26))
			_ = _nw240
			(_nw240).ArraySet1(_nwElement0_52, 0)
			(_nw240).ArraySet1(_1670_v139, 1)
			(_nw240).ArraySet1(_1670_v139, 2)
			(_nw240).ArraySet1(_1670_v139, 3)
			(_nw240).ArraySet1(_1670_v139, 4)
			(_nw240).ArraySet1(_1670_v139, 5)
			(_nw240).ArraySet1(_1670_v139, 6)
			(_nw240).ArraySet1(_1670_v139, 7)
			(_nw240).ArraySet1(_1670_v139, 8)
			(_nw240).ArraySet1(_1670_v139, 9)
			(_nw240).ArraySet1(_1670_v139, 10)
			(_nw240).ArraySet1(_1670_v139, 11)
			(_nw240).ArraySet1(_1670_v139, 12)
			(_nw240).ArraySet1(_1670_v139, 13)
			(_nw240).ArraySet1((_1676_v145).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1676_v145).Cardinality()))).Uint32()).(_dafny.Array), 14)
			(_nw240).ArraySet1(_1670_v139, 15)
			(_nw240).ArraySet1(_1670_v139, 16)
			(_nw240).ArraySet1(_1670_v139, 17)
			(_nw240).ArraySet1(_1670_v139, 18)
			(_nw240).ArraySet1(_1670_v139, 19)
			(_nw240).ArraySet1(_1670_v139, 20)
			(_nw240).ArraySet1(_1670_v139, 21)
			(_nw240).ArraySet1(_1670_v139, 22)
			(_nw240).ArraySet1(_1670_v139, 23)
			(_nw240).ArraySet1((func() _dafny.Array {
				if (_1677_v146).Contains((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)) {
					return (_1677_v146).Get((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _1670_v139
			})(), 24)
			(_nw240).ArraySet1(_1670_v139, 25)
			_1678_v147 = _nw240
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))
			_ = _index250
			(_1678_v147).ArraySet1(_1670_v139, (_index250).Int())
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))
			_ = _index251
			(_1678_v147).ArraySet1((Companion_D5_.Create_DC13_(_1670_v139)).Dtor_cf19(), (_index251).Int())
			var _1679_v148 _dafny.Sequence
			_ = _1679_v148
			_1679_v148 = _dafny.SeqOf(_1672_v141)
			var _1680_v149 _dafny.Sequence
			_ = _1680_v149
			_1680_v149 = _dafny.SeqOf(_1676_v145)
			var _1681_v150 _dafny.Array
			_ = _1681_v150
			var _nwElement0_53 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1676_v145, (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1676_v145).Cardinality()))).Uint32(), _dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int())))
			_ = _nwElement0_53
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(24))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_53, 0)
			(_nw241).ArraySet1(_1676_v145, 1)
			(_nw241).ArraySet1(_1676_v145, 2)
			(_nw241).ArraySet1(_dafny.SeqOf(_dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int()))), 3)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1676_v145, _1676_v145), 4)
			(_nw241).ArraySet1(_1676_v145, 5)
			(_nw241).ArraySet1(_1676_v145, 6)
			(_nw241).ArraySet1(_1676_v145, 7)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Update(_1676_v145, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1679_v148).Cardinality())), _dafny.IntOfUint32((_1676_v145).Cardinality()))).Uint32(), _1670_v139), 8)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1676_v145, _1676_v145), 9)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1676_v145, _1676_v145), 10)
			(_nw241).ArraySet1(_1676_v145, 11)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int()))), _dafny.SeqOf(_dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int())), _1670_v139)), 12)
			(_nw241).ArraySet1(_1676_v145, 13)
			(_nw241).ArraySet1(_dafny.SeqOf(_dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int())), _dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int())), _1670_v139), 14)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1676_v145, _1676_v145), 15)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1680_v149).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1680_v149).Cardinality()))).Uint32()).(_dafny.Sequence), _1676_v145), 16)
			(_nw241).ArraySet1(_dafny.SeqOf(_dafny.ArrayCastTo((_1678_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1678_v147), 0))).Int()))), 17)
			(_nw241).ArraySet1(_1676_v145, 18)
			(_nw241).ArraySet1(_dafny.SeqOf(_1670_v139, _1670_v139), 19)
			(_nw241).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1670_v139, _1670_v139), (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_dafny.SeqOf(_1670_v139, _1670_v139)).Cardinality()))).Uint32(), _1670_v139), 20)
			(_nw241).ArraySet1(_1676_v145, 21)
			(_nw241).ArraySet1(_1676_v145, 22)
			(_nw241).ArraySet1(_dafny.SeqOf(_1670_v139, _1670_v139), 23)
			_1681_v150 = _nw241
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1681_v150), 0))
			_ = _index252
			(_1681_v150).ArraySet1(_1676_v145, (_index252).Int())
			var _1682_v151 _dafny.Map
			_ = _1682_v151
			_1682_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(791))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg114 _dafny.Int) interface{} {
					return coer114(arg114)
				}
			}(func(_1683_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})))
			var _1684_v152 _dafny.Sequence
			_ = _1684_v152
			_1684_v152 = _dafny.SeqOf(p1, p1)
			var _1685_v153 _dafny.CodePoint
			_ = _1685_v153
			_1685_v153 = _dafny.CodePoint('e')
			var _1686_v154 D0
			_ = _1686_v154
			_1686_v154 = Companion_D0_.Create_DC1_(false, _dafny.Companion_Sequence_.Update(_1684_v152, (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1684_v152).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _dafny.CodePoint('f'))), _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1685_v153))
			var _pat_let_tv35 = _1670_v139
			_ = _pat_let_tv35
			var _pat_let_tv36 = _1670_v139
			_ = _pat_let_tv36
			var _pat_let_tv37 = _1684_v152
			_ = _pat_let_tv37
			var _1687_v155 _dafny.Array
			_ = _1687_v155
			var _nwElement0_54 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg115 _dafny.Int) interface{} {
					return coer115(arg115)
				}
			}((func(_1688_p1 _dafny.Sequence, _1689_v141 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_1690_i10 _dafny.Int) _dafny.CodePoint {
					return (_1688_p1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1689_v141), _dafny.IntOfUint32((_1688_p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(p1, _1672_v141)))
			_ = _nwElement0_54
			var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(26))
			_ = _nw242
			(_nw242).ArraySet1(_nwElement0_54, 0)
			(_nw242).ArraySet1((func() _dafny.Sequence {
				if (_1682_v151).Contains(p0) {
					return (_1682_v151).Get(p0).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("sfabqetxt")
			})(), 1)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 2)
			(_nw242).ArraySet1(p1, 3)
			(_nw242).ArraySet1(p1, 4)
			(_nw242).ArraySet1(p1, 5)
			(_nw242).ArraySet1(p1, 6)
			(_nw242).ArraySet1(p1, 7)
			(_nw242).ArraySet1(p1, 8)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 9)
			(_nw242).ArraySet1(p1, 10)
			(_nw242).ArraySet1(p1, 11)
			(_nw242).ArraySet1(p1, 12)
			(_nw242).ArraySet1(p1, 13)
			(_nw242).ArraySet1(p1, 14)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("uatwgx")), 15)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, (func(_pat_let36_0 D0) D0 {
				return func(_1691_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let37_0 bool) D0 {
						return func(_1692_dt__update_hcf1_h0 bool) D0 {
							return func(_pat_let38_0 _dafny.Sequence) D0 {
								return func(_1693_dt__update_hcf2_h0 _dafny.Sequence) D0 {
									return Companion_D0_.Create_DC1_(_1692_dt__update_hcf1_h0, _1693_dt__update_hcf2_h0, (_1691_dt__update__tmp_h1).Dtor_cf3())
								}(_pat_let38_0)
							}(_pat_let_tv37)
						}(_pat_let37_0)
					}((_pat_let_tv36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_pat_let_tv35), 0))).Int()).(bool))
				}(_pat_let36_0)
			}(_1686_v154)).Dtor_cf3()), 16)
			(_nw242).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg116 _dafny.Int) interface{} {
					return coer116(arg116)
				}
			}((func(_1694_v153 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1695_i12 _dafny.Int) _dafny.CodePoint {
					return _1694_v153
				}
			})(_1685_v153))), 17)
			(_nw242).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wocuatbi"), 18)
			(_nw242).ArraySet1(p1, 19)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 20)
			(_nw242).ArraySet1((Companion_D1_.Create_DC5_(p0, p1, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), false), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(984))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg117 _dafny.Int) interface{} {
					return coer117(arg117)
				}
			}((func(_1696_v153 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1697_i13 _dafny.Int) _dafny.CodePoint {
					return _1696_v153
				}
			})(_1685_v153)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), false)).Cardinality()))).Uint32(), (_1673_v142).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1673_v142).Cardinality()))).Uint32()).(bool)))).Dtor_cf8(), 21)
			(_nw242).ArraySet1(p1, 22)
			(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 23)
			(_nw242).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jp"), 24)
			(_nw242).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("scmvf"), 25)
			_1687_v155 = _nw242
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1687_v155), 0))
			_ = _index253
			(_1687_v155).ArraySet1(p1, (_index253).Int())
			var _1698_v156 _dafny.MultiSet
			_ = _1698_v156
			_1698_v156 = _dafny.MultiSetOf(_1672_v141)
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1681_v150), 0))
			_ = _index254
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1687_v155), 0))
			_ = _index255
			var _rhs279 bool = (_dafny.MultiSetOf(_1672_v141, _1672_v141, _1672_v141, _dafny.IntOfUint32((p1).Cardinality()))).IsDisjointFrom(_1698_v156)
			_ = _rhs279
			var _rhs280 bool = (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
			_ = _rhs280
			var _rhs281 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1676_v145, _1676_v145)
			_ = _rhs281
			var _rhs282 _dafny.Sequence = p1
			_ = _rhs282
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			var _lhs191 _dafny.Array = _1681_v150
			_ = _lhs191
			var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1681_v150), 0))
			_ = _lhs192
			var _lhs193 _dafny.Array = _1687_v155
			_ = _lhs193
			var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1687_v155), 0))
			_ = _lhs194
			_lhs189.F9 = _rhs279
			_lhs190.F13 = _rhs280
			(_lhs191).ArraySet1(_rhs281, (_lhs192).Int())
			(_lhs193).ArraySet1(_rhs282, (_lhs194).Int())
			var _1699_v157 _dafny.Set
			_ = _1699_v157
			_1699_v157 = _dafny.SetOf((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool))
			var _1700_v158 _dafny.Sequence
			_ = _1700_v158
			_1700_v158 = _dafny.SeqOf(_1699_v157)
			var _1701_v159 _dafny.MultiSet
			_ = _1701_v159
			_1701_v159 = _dafny.MultiSetOf(_1699_v157)
			_1673_v142 = (func() _dafny.Sequence {
				if (_dafny.MultiSetFromSeq(_1700_v158)).IsProperSubsetOf(_1701_v159) {
					return _1673_v142
				}
				return _1673_v142
			})()
			var _1702_v160 _dafny.Array
			_ = _1702_v160
			var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw243
			_1702_v160 = _nw243
			_1702_v160 = _1702_v160
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1687_v155), 0))
			_ = _index256
			(_1687_v155).ArraySet1(p1, (_index256).Int())
		} else {
			r2 = (_dafny.Zero).Minus((_1672_v141).Minus(_1672_v141))
			if (func() bool {
				if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(540))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}(func(_1703_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality())).Cmp(_1672_v141) <= 0 {
					return (func() bool {
						if (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool) {
							return (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
						}
						return (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
					})()
				}
				return (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
			})() {
				var _1704_v161 _dafny.Sequence
				_ = _1704_v161
				_1704_v161 = _dafny.SeqOf(p1)
				var _1705_v162 D0
				_ = _1705_v162
				_1705_v162 = Companion_D0_.Create_DC1_((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), _1704_v161, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}(func(_1706_i15 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				})))
				var _1707_v163 _dafny.MultiSet
				_ = _1707_v163
				_1707_v163 = _dafny.MultiSetOf(_1705_v162)
				var _1708_v164 *C0
				_ = _1708_v164
				var _nw244 *C0 = New_C0_()
				_ = _nw244
				_nw244.Ctor__(_1707_v163, p0)
				_1708_v164 = _nw244
				var _1709_v165 _dafny.Sequence
				_ = _1709_v165
				_1709_v165 = _dafny.UnicodeSeqOfUtf8Bytes("hrsxsmked")
				var _1710_v166 _dafny.Map
				_ = _1710_v166
				_1710_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v141, false)
				var _1711_v167 _dafny.Map
				_ = _1711_v167
				_1711_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v141, _dafny.IntOfUint32((_1673_v142).Cardinality()))
				var _1712_v168 _dafny.Set
				_ = _1712_v168
				_1712_v168 = _dafny.SetOf(Companion_Default___.SafeDivisionInt(_1672_v141, (_1710_v166).Cardinality()), (_1711_v167).Cardinality(), _1672_v141)
				var _1713_v169 D1
				_ = _1713_v169
				_1713_v169 = Companion_D1_.Create_DC5_((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), _1709_v165, _1673_v142)
				var _rhs283 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1709_v165, (_1713_v169).Dtor_cf8())
				_ = _rhs283
				var _rhs284 bool = (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
				_ = _rhs284
				var _rhs285 _dafny.Set = _1712_v168
				_ = _rhs285
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				_1709_v165 = _rhs283
				_lhs195.F9 = _rhs284
				_1712_v168 = _rhs285
				var _1714_v170 _dafny.Sequence
				_ = _1714_v170
				_1714_v170 = _dafny.SeqOf(_1672_v141, (func() _dafny.Int {
					if (_1708_v164).F19() {
						return _1672_v141
					}
					return _1672_v141
				})(), _1672_v141)
				_1714_v170 = _dafny.SeqOf((_1671_v140).Cardinality(), _1672_v141, _1672_v141, (_1671_v140).Cardinality())
				_1672_v141 = _1672_v141
				var _1715_v171 T1
				_ = _1715_v171
				var _nw245 *C5 = New_C5_()
				_ = _nw245
				_nw245.Ctor__()
				_1715_v171 = _nw245
				var _1716_v172 _dafny.Set
				_ = _1716_v172
				_1716_v172 = _dafny.SetOf(_1715_v171)
				var _1717_v173 D13
				_ = _1717_v173
				_1717_v173 = Companion_D13_.Create_DC34_(Companion_Default___.SafeModuloInt(_1672_v141, (_1716_v172).Cardinality()))
				_1717_v173 = _1717_v173
			} else {
				var _1718_v174 *C1
				_ = _1718_v174
				var _nw246 *C1 = New_C1_()
				_ = _nw246
				_nw246.Ctor__()
				_1718_v174 = _nw246
				var _1719_v175 *C6
				_ = _1719_v175
				var _nw247 *C6 = New_C6_()
				_ = _nw247
				_nw247.Ctor__()
				_1719_v175 = _nw247
				var _1720_v176 _dafny.Sequence
				_ = _1720_v176
				_1720_v176 = _dafny.SeqOf(_1672_v141)
				var _1721_v177 _dafny.Map
				_ = _1721_v177
				_1721_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1720_v176)
				var _1722_v178 _dafny.Sequence
				_ = _1722_v178
				_1722_v178 = _dafny.SeqOf(_1721_v177)
				var _1723_v179 D11
				_ = _1723_v179
				_1723_v179 = Companion_D11_.Create_DC29_((_dafny.MultiSetOf(_1719_v175)).Update(_1719_v175, Companion_Default___.Abs(((_1722_v178).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1722_v178).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())))
				_1723_v179 = _1723_v179
				(globalState).F9 = p0
				var _1724_v180 _dafny.Map
				_ = _1724_v180
				_1724_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(836))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg120 _dafny.Int) interface{} {
						return coer120(arg120)
					}
				}(func(_1725_i16 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), _1672_v141)
				var _1726_v181 T1
				_ = _1726_v181
				var _nw248 *C3 = New_C3_()
				_ = _nw248
				_nw248.Ctor__((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), (_1505_v0).Cardinality())
				_1726_v181 = _nw248
				var _1727_v182 _dafny.MultiSet
				_ = _1727_v182
				_1727_v182 = _dafny.MultiSetOf(_1726_v181)
				_1724_v180 = (_1724_v180).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg121 _dafny.Int) interface{} {
						return coer121(arg121)
					}
				}(func(_1728_i17 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), (_1727_v182).Cardinality())
				var _1729_v183 _dafny.Array
				_ = _1729_v183
				var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
				_ = _nw249
				_1729_v183 = _nw249
				var _1730_v184 _dafny.CodePoint
				_ = _1730_v184
				_1730_v184 = _dafny.CodePoint('b')
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1729_v183), 0))
				_ = _index257
				(_1729_v183).ArraySet1CodePoint(_1730_v184, (_index257).Int())
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1729_v183), 0))
				_ = _index258
				(_1729_v183).ArraySet1CodePoint((p1).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_1673_v142).Cardinality())).Minus(_1672_v141), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index258).Int())
			}
			var _1731_v185 _dafny.Array
			_ = _1731_v185
			var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
			_ = _nw250
			_1731_v185 = _nw250
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_1731_v185), 0))
			_ = _index259
			(_1731_v185).ArraySet1CodePoint(_dafny.CodePoint('c'), (_index259).Int())
			var _1732_v186 _dafny.Map
			_ = _1732_v186
			_1732_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint), _1671_v140)
			var _1733_v187 _dafny.Sequence
			_ = _1733_v187
			_1733_v187 = _dafny.SeqOf(_1672_v141, _1672_v141)
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))
			_ = _index260
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_1731_v185), 0))
			_ = _index261
			var _rhs286 bool = (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool)
			_ = _rhs286
			var _rhs287 bool = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("wkpov"), p1)
			_ = _rhs287
			var _rhs288 _dafny.CodePoint = Companion_Default___.Fm23(_dafny.Companion_Sequence_.IsPrefixOf(p1, p1), _1732_v186, (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), (_1733_v187).Select((Companion_Default___.SafeIndex(_1672_v141, _dafny.IntOfUint32((_1733_v187).Cardinality()))).Uint32()).(_dafny.Int), globalState)
			_ = _rhs288
			var _rhs289 _dafny.Int = _dafny.IntOfInt64(986)
			_ = _rhs289
			var _rhs290 _dafny.Int = _1672_v141
			_ = _rhs290
			var _lhs196 *GlobalState = globalState
			_ = _lhs196
			var _lhs197 _dafny.Array = _1670_v139
			_ = _lhs197
			var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))
			_ = _lhs198
			var _lhs199 _dafny.Array = _1731_v185
			_ = _lhs199
			var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_1731_v185), 0))
			_ = _lhs200
			_lhs196.F3 = _rhs286
			(_lhs197).ArraySet1(_rhs287, (_lhs198).Int())
			(_lhs199).ArraySet1CodePoint(_rhs288, (_lhs200).Int())
			r2 = _rhs289
			r2 = _rhs290
			(globalState).F5 = true
			var _1734_v188 _dafny.Set
			_ = _1734_v188
			_1734_v188 = _dafny.SetOf(_1505_v0, _1505_v0)
			var _1735_v189 _dafny.MultiSet
			_ = _1735_v189
			_1735_v189 = _dafny.MultiSetOf(_dafny.IntOfInt64(759))
			var _1736_v190 _dafny.Map
			_ = _1736_v190
			_1736_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1734_v188).IsDisjointFrom(Companion_Default___.Fm43((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), globalState)), _1735_v189)
			var _1737_v191 _dafny.Map
			_ = _1737_v191
			_1737_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1670_v139, _dafny.SeqOf(false, p0))
			_1736_v190 = (_1736_v190).Update(Companion_Default___.Fm2(Companion_Default___.Fm23((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), _1732_v186, (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool), (_1737_v191).Cardinality(), globalState), Companion_D0_.Create_DC3_(_1674_v143), _1672_v141, globalState), _1735_v189)
		}
		var _1738_v192 _dafny.Set
		_ = _1738_v192
		_1738_v192 = _dafny.SetOf(Companion_Default___.Fm6(globalState), (_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool))
		r0 = _1738_v192
		r1 = !((_1670_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1670_v139), 0))).Int()).(bool))
		r2 = _1672_v141
		r3 = p0
		return r0, r1, r2, r3
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
