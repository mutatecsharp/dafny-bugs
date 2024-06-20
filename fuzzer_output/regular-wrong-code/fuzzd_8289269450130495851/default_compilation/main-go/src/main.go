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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.IntOfInt64(834)).Cmp((_dafny.IntOfInt64(-650)).Times(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))) != 0
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((func(_source0 _dafny.Sequence) _dafny.Set {
		var _0___mcc_h0 _dafny.Sequence = _source0
		_ = _0___mcc_h0
		var _1_cf0 _dafny.Sequence = _0___mcc_h0
		_ = _1_cf0
		return (_dafny.SetOf(_1_cf0)).Union(_dafny.SetOf(_1_cf0, _1_cf0, _1_cf0))
	}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(597))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vpftgh"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(639))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})), _dafny.UnicodeSeqOfUtf8Bytes("jimxsttfl")))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Sequence {
	var _source1 _dafny.Set = _dafny.SetOf(_dafny.SeqOf(true), _dafny.SeqOf(true))
	_ = _source1
	var _4___mcc_h0 _dafny.Set = _source1
	_ = _4___mcc_h0
	var _5_cf62 _dafny.Set = _4___mcc_h0
	_ = _5_cf62
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(875)
	}))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if false {
			return _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ouixhpwjj")).Cardinality()), ((Companion_D9_.Create_DC20_(_dafny.IntOfUint32((_dafny.SeqOf(false, true, false, false, false)).Cardinality()), true, false, false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(315))).Cardinality(), _dafny.IntOfInt64(106)))).Dtor_cf38()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-680), _dafny.SeqOf(_dafny.CodePoint('f')))).Cardinality())
		}
		return (Companion_D9_.Create_DC17_(func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(786), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-631), _dafny.IntOfInt64(-19))).Cardinality())).Keys().Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _7_v0 _dafny.Int
				_7_v0 = interface{}(_compr_0).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(786), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-631), _dafny.IntOfInt64(-19))).Cardinality())).Contains(_7_v0) {
					_coll0.Add(Companion_Default___.SafeDivisionInt(_7_v0, _dafny.IntOfInt64(389)))
				}
			}
			return _coll0.ToSet()
		}())).Dtor_cf31()
	})()).Union((_dafny.SetOf(_dafny.IntOfInt64(531), _dafny.IntOfUint32((_dafny.SeqOf(false, true, false)).Cardinality()), (_dafny.SetOf(!(false))).Cardinality())).Intersection(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-219))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_8_i0 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("suq")).Cardinality()), !(false))).Cardinality()
		}))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _9_v1 _dafny.Int
			_9_v1 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-219))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_8_i0 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("suq")).Cardinality()), !(false))).Cardinality()
			})), _9_v1) {
				_coll1.Add((_9_v1).Plus(_dafny.IntOfInt64(199)))
			}
		}
		return _coll1.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if true {
		if true {
			return _dafny.CodePoint('f')
		} else {
			return _dafny.CodePoint('r')
		}
	} else {
		return _dafny.CodePoint('t')
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false, true, false, false), (Companion_D7_.Create_DC13_(_dafny.SeqOf(true, true))).Dtor_cf23()), (func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(false, true)
		}
		return (Companion_D7_.Create_DC13_(_dafny.SeqOf(false))).Dtor_cf23()
	})())
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source2 D18 = Companion_D18_.Create_DC46_(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_ = _source2
	if _source2.Is_DC47() {
		var _10___mcc_h0 _dafny.Int = _source2.Get_().(D18_DC47).Cf86
		_ = _10___mcc_h0
		var _11___mcc_h1 _dafny.CodePoint = _source2.Get_().(D18_DC47).Cf87
		_ = _11___mcc_h1
		var _12_cf87 _dafny.CodePoint = _11___mcc_h1
		_ = _12_cf87
		var _13_cf86 _dafny.Int = _10___mcc_h0
		_ = _13_cf86
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.UnicodeSeqOfUtf8Bytes("obdympjc"))
	} else {
		var _14___mcc_h2 _dafny.Set = _source2.Get_().(D18_DC46).Cf85
		_ = _14___mcc_h2
		var _15_cf85 _dafny.Set = _14___mcc_h2
		_ = _15_cf85
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})), _dafny.UnicodeSeqOfUtf8Bytes("rw"))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-721)).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(_dafny.IntOfInt64(771), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !((Companion_D16_.Create_DC40_(true, _dafny.IntOfInt64(499), (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()))).Dtor_cf74())), _dafny.SeqOf(true, true, false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(767)), _dafny.SeqOf(_dafny.IntOfInt64(676))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(false, true)).Cardinality()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), false)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("otnqi")).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	var _source3 D15 = Companion_D15_.Create_DC36_(_dafny.CodePoint('c'), _dafny.IntOfInt64(604))
	_ = _source3
	if _source3.Is_DC36() {
		var _17___mcc_h0 _dafny.CodePoint = _source3.Get_().(D15_DC36).Cf64
		_ = _17___mcc_h0
		var _18___mcc_h1 _dafny.Int = _source3.Get_().(D15_DC36).Cf65
		_ = _18___mcc_h1
		var _19_cf65 _dafny.Int = _18___mcc_h1
		_ = _19_cf65
		var _20_cf64 _dafny.CodePoint = _17___mcc_h0
		_ = _20_cf64
		return _dafny.SetOf(false, !(true), !(true), !(true))
	} else if _source3.Is_DC37() {
		var _21___mcc_h2 bool = _source3.Get_().(D15_DC37).Cf66
		_ = _21___mcc_h2
		var _22___mcc_h3 _dafny.Int = _source3.Get_().(D15_DC37).Cf67
		_ = _22___mcc_h3
		var _23___mcc_h4 _dafny.Array = _source3.Get_().(D15_DC37).Cf68
		_ = _23___mcc_h4
		var _24___mcc_h5 bool = _source3.Get_().(D15_DC37).Cf69
		_ = _24___mcc_h5
		var _25_cf69 bool = _24___mcc_h5
		_ = _25_cf69
		var _26_cf68 _dafny.Array = _23___mcc_h4
		_ = _26_cf68
		var _27_cf67 _dafny.Int = _22___mcc_h3
		_ = _27_cf67
		var _28_cf66 bool = _21___mcc_h2
		_ = _28_cf66
		return (_dafny.SetOf(_28_cf66)).Intersection(_dafny.SetOf(_25_cf69))
	} else if _source3.Is_DC38() {
		var _29___mcc_h6 _dafny.Int = _source3.Get_().(D15_DC38).Cf70
		_ = _29___mcc_h6
		var _30___mcc_h7 bool = _source3.Get_().(D15_DC38).Cf71
		_ = _30___mcc_h7
		var _31___mcc_h8 _dafny.Int = _source3.Get_().(D15_DC38).Cf72
		_ = _31___mcc_h8
		var _32_cf72 _dafny.Int = _31___mcc_h8
		_ = _32_cf72
		var _33_cf71 bool = _30___mcc_h7
		_ = _33_cf71
		var _34_cf70 _dafny.Int = _29___mcc_h6
		_ = _34_cf70
		return (Companion_D24_.Create_DC58_(_dafny.SetOf(_33_cf71, _33_cf71))).Dtor_cf109()
	} else {
		var _35___mcc_h9 _dafny.Sequence = _source3.Get_().(D15_DC35).Cf63
		_ = _35___mcc_h9
		var _36_cf63 _dafny.Sequence = _35___mcc_h9
		_ = _36_cf63
		return _dafny.SetOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(936))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_37_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()))), _dafny.SeqOf(_dafny.IntOfInt64(-501)))).Union(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(878))))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Set {
	if (_dafny.IntOfInt64(896)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality())).Cardinality())) >= 0 {
		return _dafny.SetOf(_dafny.CodePoint('t'))
	} else {
		return _dafny.SetOf(_dafny.CodePoint('s'), _dafny.CodePoint('l'), _dafny.CodePoint('h'), _dafny.CodePoint('k'), _dafny.CodePoint('g'))
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(true, false), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(false)
		}
		return _dafny.SeqOf(!(true))
	})())
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
	})())
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.IntOfInt64(295))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(856)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("skaroir")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(false, !(true))), _dafny.SeqOf(!(true), false))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source4 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.IntOfInt64(520), _dafny.IntOfInt64(-680))
	_ = _source4
	var _38___mcc_h0 _dafny.MultiSet = _source4
	_ = _38___mcc_h0
	var _39_cf18 _dafny.MultiSet = _38___mcc_h0
	_ = _39_cf18
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(264))).Union(_dafny.MultiSetOf(((Companion_D9_.Create_DC17_(_dafny.SetOf(_dafny.IntOfInt64(743)))).Dtor_cf31()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_40_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(true)
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()))).Cardinality(), true))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, p2 bool, globalState *GlobalState) bool {
	return (func() bool {
		if true {
			return true
		}
		return false
	})()
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(229), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(978), true), false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(138), _dafny.IntOfInt64(704)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-387), _dafny.IntOfInt64(884)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	var _source5 D9 = Companion_D9_.Create_DC18_(!(true))
	_ = _source5
	if _source5.Is_DC18() {
		var _41___mcc_h0 bool = _source5.Get_().(D9_DC18).Cf32
		_ = _41___mcc_h0
		var _42_cf32 bool = _41___mcc_h0
		_ = _42_cf32
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rpb"), _dafny.UnicodeSeqOfUtf8Bytes("ypsbjt"))
	} else if _source5.Is_DC19() {
		var _43___mcc_h1 bool = _source5.Get_().(D9_DC19).Cf33
		_ = _43___mcc_h1
		var _44_cf33 bool = _43___mcc_h1
		_ = _44_cf33
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xa"), _dafny.UnicodeSeqOfUtf8Bytes("vxt"))
	} else if _source5.Is_DC20() {
		var _45___mcc_h2 _dafny.Int = _source5.Get_().(D9_DC20).Cf34
		_ = _45___mcc_h2
		var _46___mcc_h3 bool = _source5.Get_().(D9_DC20).Cf35
		_ = _46___mcc_h3
		var _47___mcc_h4 bool = _source5.Get_().(D9_DC20).Cf36
		_ = _47___mcc_h4
		var _48___mcc_h5 bool = _source5.Get_().(D9_DC20).Cf37
		_ = _48___mcc_h5
		var _49___mcc_h6 _dafny.Map = _source5.Get_().(D9_DC20).Cf38
		_ = _49___mcc_h6
		var _50_cf38 _dafny.Map = _49___mcc_h6
		_ = _50_cf38
		var _51_cf37 bool = _48___mcc_h5
		_ = _51_cf37
		var _52_cf36 bool = _47___mcc_h4
		_ = _52_cf36
		var _53_cf35 bool = _46___mcc_h3
		_ = _53_cf35
		var _54_cf34 _dafny.Int = _45___mcc_h2
		_ = _54_cf34
		if _51_cf37 {
			return _dafny.UnicodeSeqOfUtf8Bytes("eg")
		} else {
			return _dafny.UnicodeSeqOfUtf8Bytes("uj")
		}
	} else {
		var _55___mcc_h7 _dafny.Set = _source5.Get_().(D9_DC17).Cf31
		_ = _55___mcc_h7
		var _56_cf31 _dafny.Set = _55___mcc_h7
		_ = _56_cf31
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_57_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(865))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_58_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("qaed")), (_dafny.MultiSetOf(_dafny.IntOfInt64(672))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(573), _dafny.IntOfInt64(-43), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()), _dafny.IntOfInt64(950))).Cardinality()), _dafny.IntOfInt64(615), _dafny.IntOfInt64(-460))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false, true)), _dafny.SeqOf(true, !(false)))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, globalState *GlobalState) D2 {
	var _source6 D20 = Companion_D20_.Create_DC51_()
	_ = _source6
	if _source6.Is_DC51() {
		if !(true) {
			return Companion_D2_.Create_DC4_((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), _dafny.IntOfInt64(608), _dafny.IntOfInt64(-857), _dafny.CodePoint('f'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(199))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_59_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})))
		} else {
			return Companion_D2_.Create_DC4_(_dafny.IntOfInt64(478), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("crgitqdfk")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hlk")).Cardinality()), _dafny.CodePoint('k'), _dafny.UnicodeSeqOfUtf8Bytes("ovgnevf"))
		}
	} else {
		var _60___mcc_h0 _dafny.MultiSet = _source6.Get_().(D20_DC50).Cf92
		_ = _60___mcc_h0
		var _61_cf92 _dafny.MultiSet = _60___mcc_h0
		_ = _61_cf92
		return Companion_D2_.Create_DC4_((_dafny.Zero).Minus(_dafny.IntOfInt64(-718)), _dafny.IntOfInt64(896), _dafny.IntOfInt64(-841), _dafny.CodePoint('u'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-641))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_62_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_63_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('g')
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) D11 {
	var _source7 _dafny.Set = func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _64_v0 _dafny.Sequence
				_64_v0 = interface{}(_compr_3).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Contains(_64_v0) {
					_coll3.Add(_64_v0, !(false))
				}
			}
			return _coll3.ToMap()
		}()).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _65_v1 _dafny.Sequence
			_65_v1 = interface{}(_compr_2).(_dafny.Sequence)
			if (func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Keys().Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _64_v0 _dafny.Sequence
					_64_v0 = interface{}(_compr_4).(_dafny.Sequence)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Contains(_64_v0) {
						_coll4.Add(_64_v0, !(false))
					}
				}
				return _coll4.ToMap()
			}()).Contains(_65_v1) {
				_coll2.Add(_65_v1)
			}
		}
		return _coll2.ToSet()
	}()
	_ = _source7
	var _66___mcc_h0 _dafny.Set = _source7
	_ = _66___mcc_h0
	var _67_cf62 _dafny.Set = _66___mcc_h0
	_ = _67_cf62
	return Companion_D11_.Create_DC24_((_dafny.MultiSetOf(false, false)).Cardinality(), Companion_D4_.Create_DC8_(_dafny.CodePoint('n'), _dafny.IntOfInt64(190), false))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ofqxguj"), _dafny.SeqOf(true, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vh"), _dafny.SeqOf(!(!(false)))))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if false {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)).Cardinality()
		}
		return _dafny.IntOfInt64(213)
	})(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(661), _dafny.IntOfInt64(501)))
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(545))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_68_i0 _dafny.Int) _dafny.Int {
		return (func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('p'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('d'))).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _69_v0 _dafny.CodePoint
				_69_v0 = interface{}(_compr_5).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('p'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('d')), _69_v0) {
					_coll5.Add(_69_v0)
				}
			}
			return _coll5.ToSet()
		}()).Cardinality()
	})), _dafny.SeqOf(_dafny.IntOfInt64(676), _dafny.IntOfInt64(-823))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(313))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_70_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-438)
	})), _dafny.SeqOf(_dafny.IntOfInt64(631))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-592))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_71_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(false, true, false)
	}))
}
func (_static *CompanionStruct_Default___) Fm47(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.SetOf(false)).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D6_.Create_DC12_(false, _dafny.SeqOf(_dafny.SeqOf(true, true), _dafny.SeqOf(true)), _dafny.UnicodeSeqOfUtf8Bytes("kvltui")), Companion_D6_.Create_DC12_(false, _dafny.SeqOf(_dafny.SeqOf(false)), (Companion_D6_.Create_DC12_(false, _dafny.SeqOf(_dafny.SeqOf(!(true), true)), _dafny.UnicodeSeqOfUtf8Bytes("dyruvhv"))).Dtor_cf22()))).Cardinality()))).Cardinality())), !(true)))
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D1_.Create_DC1_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(876), _dafny.IntOfInt64(396))).Cardinality(), true))).Dtor_cf1()
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Int, globalState *GlobalState) D4 {
	if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(587)))).IsProperSubsetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-927)))) {
		return Companion_D4_.Create_DC9_(_dafny.IntOfUint32((_dafny.SeqOf(false, true, true)).Cardinality()), !(false))
	} else {
		return Companion_D4_.Create_DC9_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okgwkrflp")).Cardinality()), true)
	}
}
func (_static *CompanionStruct_Default___) Fm50(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-469), (_dafny.MultiSetOf(_dafny.CodePoint('u'))).Cardinality()), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-827)))))).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(256))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_72_i0 _dafny.Int) _dafny.Int {
		return ((Companion_D24_.Create_DC58_(_dafny.SetOf(false, true))).Dtor_cf109()).Cardinality()
	})), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(880))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(942), (_dafny.MultiSetOf(_dafny.IntOfInt64(-878), _dafny.IntOfInt64(888))).Cardinality()), _dafny.SeqOf((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(996), _dafny.IntOfInt64(410))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _73_v0 _dafny.Int
			_73_v0 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(996)).Cmp(_73_v0) <= 0) && ((_73_v0).Cmp(_dafny.IntOfInt64(410)) < 0) {
				_coll6.Add(Companion_Default___.SafeDivisionInt(_73_v0, _dafny.IntOfInt64(72)), false)
			}
		}
		return _coll6.ToMap()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Sequence, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), !(false))), (func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('u'))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _74_v0 _dafny.CodePoint
			_74_v0 = interface{}(_compr_7).(_dafny.CodePoint)
			if (_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('u'))).Contains(_74_v0) {
				_coll7.Add(_74_v0, true)
			}
		}
		return _coll7.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), false)).Keys().Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _75_v1 _dafny.CodePoint
			_75_v1 = interface{}(_compr_8).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), false)).Contains(_75_v1) {
				_coll8.Add(_75_v1, true)
			}
		}
		return _coll8.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC18_(false)
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, true), _dafny.SeqOf(false)), (_dafny.SetOf(false)).IsProperSubsetOf(_dafny.SetOf(true, false, true, true, false)))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-113))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_76_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()), _dafny.IntOfInt64(80))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('n'))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(524), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-541))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_77_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality())))).Cardinality())))), (_dafny.SetOf(_dafny.IntOfInt64(868), _dafny.IntOfInt64(-924))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, p1 D2, p2 bool, p3 _dafny.Int, globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC40_(!(!(true)), ((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('s'))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _78_v0 _dafny.CodePoint
			_78_v0 = interface{}(_compr_9).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('s'))).Contains(_78_v0) {
				_coll9.Add(_78_v0, false)
			}
		}
		return _coll9.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), true))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), _dafny.IntOfInt64(651))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm56(p0 D4, p1 bool, p2 D15, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_79_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), _dafny.UnicodeSeqOfUtf8Bytes("vgydj"), _dafny.UnicodeSeqOfUtf8Bytes("mxhrte"))).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("fqymdm")))).Intersection(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("hydj")))
}
func (_static *CompanionStruct_Default___) Fm57(p0 bool, p1 bool, globalState *GlobalState) D15 {
	return Companion_D15_.Create_DC35_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(364), _dafny.IntOfInt64(789))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _80_v0 _dafny.Int
			_80_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(364)).Cmp(_80_v0) <= 0) && ((_80_v0).Cmp(_dafny.IntOfInt64(789)) < 0) {
				_coll10.Add((_80_v0).Times(_dafny.IntOfInt64(171)), _dafny.MultiSetOf(false))
			}
		}
		return _coll10.ToMap()
	}()).Cardinality()), _dafny.CodePoint('q'))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var _pat_let_tv0 = p1
	_ = _pat_let_tv0
	var _source8 D13 = func(_pat_let0_0 D13) D13 {
		return func(_81_dt__update__tmp_h0 D13) D13 {
			return func(_pat_let1_0 _dafny.Int) D13 {
				return func(_82_dt__update_hcf56_h0 _dafny.Int) D13 {
					return Companion_D13_.Create_DC31_(_82_dt__update_hcf56_h0)
				}(_pat_let1_0)
			}((_dafny.Zero).Minus((_dafny.Zero).Minus(_pat_let_tv0)))
		}(_pat_let0_0)
	}(Companion_D13_.Create_DC31_(p2))
	_ = _source8
	if _source8.Is_DC31() {
		var _83___mcc_h0 _dafny.Int = _source8.Get_().(D13_DC31).Cf56
		_ = _83___mcc_h0
		var _84_cf56 _dafny.Int = _83___mcc_h0
		_ = _84_cf56
		var _85_v0 _dafny.Array
		_ = _85_v0
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = func(_86_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wqcmjr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}(func(_87_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})))
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
		_85_v0 = _nw0
		_85_v0 = _85_v0
		var _88_v1 D4
		_ = _88_v1
		_88_v1 = Companion_D4_.Create_DC9_(_dafny.IntOfInt64(-389), p0)
		_88_v1 = (func() D4 {
			if p0 {
				return _88_v1
			}
			return _88_v1
		})()
		if Companion_Default___.Fm0((p3).Minus(_84_cf56), globalState) {
			var _89_v2 _dafny.Set
			_ = _89_v2
			_89_v2 = _dafny.SetOf(p2, p2)
			var _90_v3 _dafny.Sequence
			_ = _90_v3
			_90_v3 = _dafny.UnicodeSeqOfUtf8Bytes("hqrd")
			var _91_v4 _dafny.Map
			_ = _91_v4
			_91_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _90_v3)
			var _92_v5 _dafny.MultiSet
			_ = _92_v5
			_92_v5 = _dafny.MultiSetOf(p0, p0, p0, p0, p0)
			var _93_v6 D10
			_ = _93_v6
			_93_v6 = Companion_D10_.Create_DC22_((_91_v4).Cardinality(), _92_v5)
			var _94_v7 _dafny.Sequence
			_ = _94_v7
			_94_v7 = _dafny.SeqOf(p0)
			var _95_v8 _dafny.Map
			_ = _95_v8
			_95_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _85_v0)
			var _96_v9 *C1
			_ = _96_v9
			var _nw1 *C1 = New_C1_()
			_ = _nw1
			_nw1.Ctor__((_89_v2).IsDisjointFrom(Companion_Default___.Fm10(globalState)), (_93_v6).Dtor_cf41(), Companion_Default___.Fm16(p3, globalState), _dafny.IntOfUint32((_94_v7).Cardinality()), (func() _dafny.Array {
				if (_95_v8).Contains(_dafny.IntOfInt64(471)) {
					return (_95_v8).Get(_dafny.IntOfInt64(471)).(_dafny.Array)
				}
				return _85_v0
			})())
			_96_v9 = _nw1
			var _97_v10 *C0
			_ = _97_v10
			var _nw2 *C0 = New_C0_()
			_ = _nw2
			_nw2.Ctor__(_84_cf56, _dafny.Companion_Sequence_.Concatenate(_90_v3, _90_v3))
			_97_v10 = _nw2
			_97_v10 = _97_v10
			var _98_v12 T0
			_ = _98_v12
			var _nw3 *C3 = New_C3_()
			_ = _nw3
			_nw3.Ctor__(p1, _85_v0)
			_98_v12 = _nw3
			var _99_v13 _dafny.MultiSet
			_ = _99_v13
			_99_v13 = _dafny.MultiSetOf(_98_v12, _98_v12)
			var _100_v14 _dafny.MultiSet
			_ = _100_v14
			_100_v14 = _dafny.MultiSetOf(_dafny.IntOfUint32((_90_v3).Cardinality()), (_99_v13).Cardinality())
			var _101_v15 _dafny.Sequence
			_ = _101_v15
			_101_v15 = _dafny.SeqOf((_97_v10).F36(), p2)
			var _102_v17 _dafny.Map
			_ = _102_v17
			_102_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v9.F35, p0)
			(globalState).F3 = (Companion_Default___.SafeModuloInt(((func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_100_v14).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _103_v11 _dafny.Int
					_103_v11 = interface{}(_compr_11).(_dafny.Int)
					if (_100_v14).Contains(_103_v11) {
						_coll11.Add(Companion_Default___.SafeDivisionInt(_103_v11, (_97_v10).F36()), _dafny.IntOfUint32(((_97_v10).F37()).Cardinality()))
					}
				}
				return _coll11.ToMap()
			}()).Update((_101_v15).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p3), _dafny.IntOfUint32((_101_v15).Cardinality()))).Uint32()).(_dafny.Int), p2)).Cardinality(), (func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(719), _dafny.IntOfInt64(727))); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _104_v16 _dafny.Int
					_104_v16 = interface{}(_compr_12).(_dafny.Int)
					if ((_dafny.IntOfInt64(719)).Cmp(_104_v16) <= 0) && ((_104_v16).Cmp(_dafny.IntOfInt64(727)) < 0) {
						_coll12.Add((_104_v16).Times(_dafny.IntOfInt64(105)))
					}
				}
				return _coll12.ToSet()
			}()).Cardinality())).Times(((_102_v17).Cardinality()).Plus(_84_cf56))
			var _105_v18 *C6
			_ = _105_v18
			var _nw4 *C6 = New_C6_()
			_ = _nw4
			_nw4.Ctor__((_96_v9).Fm3(_92_v5, globalState), _dafny.IntOfInt64(-792), (_92_v5).Difference(Companion_Default___.Fm32(globalState)), _101_v15, p2, _85_v0, (_93_v6).Dtor_cf41(), _dafny.IntOfUint32((Companion_Default___.Fm8(globalState)).Cardinality()))
			_105_v18 = _nw4
			var _106_v19 _dafny.Array
			_ = _106_v19
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw5
			_106_v19 = _nw5
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_106_v19), 0))
			_ = _index0
			(_106_v19).ArraySet1((_100_v14).Cardinality(), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_106_v19), 0))
			_ = _index1
			(_106_v19).ArraySet1((func() _dafny.Int {
				if true {
					return (_97_v10).F36()
				}
				return (func() _dafny.Int {
					if (_92_v5).Contains((_105_v18).F33()) {
						return (_92_v5).Multiplicity((_105_v18).F33())
					}
					return _105_v18.F34
				})()
			})(), (_index1).Int())
		} else {
			_84_cf56 = Companion_Default___.SafeDivisionInt(p2, p1)
			var _107_v20 _dafny.CodePoint
			_ = _107_v20
			_107_v20 = _dafny.CodePoint('s')
			var _108_v21 _dafny.Set
			_ = _108_v21
			_108_v21 = _dafny.SetOf(_107_v20)
			var _109_v24 _dafny.Array
			_ = _109_v24
			var _nwElement0_0 _dafny.Set = _108_v21
			_ = _nwElement0_0
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(19))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_0, 0)
			(_nw6).ArraySet1(_108_v21, 1)
			(_nw6).ArraySet1(_108_v21, 2)
			(_nw6).ArraySet1(_108_v21, 3)
			(_nw6).ArraySet1(Companion_Default___.Fm19(globalState), 4)
			(_nw6).ArraySet1(_dafny.SetOf(_107_v20), 5)
			(_nw6).ArraySet1(_108_v21, 6)
			(_nw6).ArraySet1(_108_v21, 7)
			(_nw6).ArraySet1(_108_v21, 8)
			(_nw6).ArraySet1(_108_v21, 9)
			(_nw6).ArraySet1(_108_v21, 10)
			(_nw6).ArraySet1((func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v20, _107_v20)).Keys().Elements()); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _110_v22 _dafny.CodePoint
						_110_v22 = interface{}(_compr_14).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v20, _107_v20)).Contains(_110_v22) {
							_coll14.Add(_110_v22, false)
						}
					}
					return _coll14.ToMap()
				}()).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _111_v23 _dafny.CodePoint
					_111_v23 = interface{}(_compr_13).(_dafny.CodePoint)
					if (func() _dafny.Map {
						var _coll15 = _dafny.NewMapBuilder()
						_ = _coll15
						for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v20, _107_v20)).Keys().Elements()); ; {
							_compr_15, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _110_v22 _dafny.CodePoint
							_110_v22 = interface{}(_compr_15).(_dafny.CodePoint)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v20, _107_v20)).Contains(_110_v22) {
								_coll15.Add(_110_v22, false)
							}
						}
						return _coll15.ToMap()
					}()).Contains(_111_v23) {
						_coll13.Add(_111_v23)
					}
				}
				return _coll13.ToSet()
			}()).Union(_108_v21), 11)
			(_nw6).ArraySet1((_108_v21).Difference(_108_v21), 12)
			(_nw6).ArraySet1(_108_v21, 13)
			(_nw6).ArraySet1(_dafny.SetOf(_107_v20), 14)
			(_nw6).ArraySet1(_108_v21, 15)
			(_nw6).ArraySet1(_108_v21, 16)
			(_nw6).ArraySet1((_108_v21).Union(_108_v21), 17)
			(_nw6).ArraySet1(_108_v21, 18)
			_109_v24 = _nw6
			var _112_v25 _dafny.Sequence
			_ = _112_v25
			_112_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khlhpemq")).Cardinality()))
			var _113_v26 _dafny.Array
			_ = _113_v26
			var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
			_ = _nw7
			_113_v26 = _nw7
			var _114_v27 _dafny.Map
			_ = _114_v27
			_114_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v26, p0)
			var _115_v28 _dafny.Array
			_ = _115_v28
			var _nwElement0_1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_112_v25, _dafny.SeqOf(p3, (_114_v27).Cardinality()))
			_ = _nwElement0_1
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(7))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_1, 0)
			(_nw8).ArraySet1(_112_v25, 1)
			(_nw8).ArraySet1(_dafny.SeqOf(p1, p3, _dafny.IntOfInt64(7)), 2)
			(_nw8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_112_v25, _dafny.Companion_Sequence_.Update(_112_v25, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_112_v25).Cardinality()))).Uint32(), _dafny.IntOfInt64(-307))), 3)
			(_nw8).ArraySet1(_112_v25, 4)
			(_nw8).ArraySet1(_112_v25, 5)
			(_nw8).ArraySet1(_112_v25, 6)
			_115_v28 = _nw8
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_115_v28), 0))
			_ = _index2
			(_115_v28).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(754))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_116_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_117_i2 _dafny.Int) _dafny.Int {
					return _116_p2
				}
			})(p2))), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_115_v28), 0))
			_ = _index3
			var _rhs0 _dafny.Array = _109_v24
			_ = _rhs0
			var _rhs1 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _dafny.Companion_Sequence_.Concatenate(_112_v25, _112_v25)
				}
				return _112_v25
			})()
			_ = _rhs1
			var _lhs0 _dafny.Array = _115_v28
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_115_v28), 0))
			_ = _lhs1
			_109_v24 = _rhs0
			(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
			var _118_v29 _dafny.Map
			_ = _118_v29
			_118_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_cf56, _dafny.IntOfInt64(690))
			var _119_v31 _dafny.Sequence
			_ = _119_v31
			_119_v31 = _dafny.UnicodeSeqOfUtf8Bytes("adhmxyo")
			var _120_v32 _dafny.Array
			_ = _120_v32
			var _nwElement0_2 _dafny.Map = _118_v29
			_ = _nwElement0_2
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(27))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_2, 0)
			(_nw9).ArraySet1(_118_v29, 1)
			(_nw9).ArraySet1(_118_v29, 2)
			(_nw9).ArraySet1(_118_v29, 3)
			(_nw9).ArraySet1(_118_v29, 4)
			(_nw9).ArraySet1(_118_v29, 5)
			(_nw9).ArraySet1(_118_v29, 6)
			(_nw9).ArraySet1(_118_v29, 7)
			(_nw9).ArraySet1(func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-415), _dafny.IntOfInt64(552))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _121_v30 _dafny.Int
					_121_v30 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-415)).Cmp(_121_v30) <= 0) && ((_121_v30).Cmp(_dafny.IntOfInt64(552)) < 0) {
						_coll16.Add(Companion_Default___.SafeDivisionInt(_121_v30, _84_cf56), p3)
					}
				}
				return _coll16.ToMap()
			}(), 8)
			(_nw9).ArraySet1(_118_v29, 9)
			(_nw9).ArraySet1(_118_v29, 10)
			(_nw9).ArraySet1(_118_v29, 11)
			(_nw9).ArraySet1(_118_v29, 12)
			(_nw9).ArraySet1(_118_v29, 13)
			(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _84_cf56), 14)
			(_nw9).ArraySet1(_118_v29, 15)
			(_nw9).ArraySet1(_118_v29, 16)
			(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _84_cf56), 17)
			(_nw9).ArraySet1(_118_v29, 18)
			(_nw9).ArraySet1(_118_v29, 19)
			(_nw9).ArraySet1(_118_v29, 20)
			(_nw9).ArraySet1(_118_v29, 21)
			(_nw9).ArraySet1(_118_v29, 22)
			(_nw9).ArraySet1(_118_v29, 23)
			(_nw9).ArraySet1(_118_v29, 24)
			(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_119_v31).Cardinality())), 25)
			(_nw9).ArraySet1(_118_v29, 26)
			_120_v32 = _nw9
			var _122_v33 _dafny.Map
			_ = _122_v33
			_122_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_D16_.Create_DC39_(_120_v32))
			var _123_v34 _dafny.MultiSet
			_ = _123_v34
			_123_v34 = _dafny.MultiSetOf(_122_v33)
			_123_v34 = _123_v34
			var _124_v35 _dafny.MultiSet
			_ = _124_v35
			_124_v35 = _dafny.MultiSetOf(_107_v20)
			var _125_v36 _dafny.Array
			_ = _125_v36
			var _nwElement0_3 _dafny.Int = (_dafny.Zero).Minus(_84_cf56)
			_ = _nwElement0_3
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(10))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_3, 0)
			(_nw10).ArraySet1(_84_cf56, 1)
			(_nw10).ArraySet1(_84_cf56, 2)
			(_nw10).ArraySet1(_84_cf56, 3)
			(_nw10).ArraySet1(Companion_Default___.SafeDivisionInt((_124_v35).Cardinality(), _84_cf56), 4)
			(_nw10).ArraySet1(p3, 5)
			(_nw10).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eiavcl")).Cardinality()), 6)
			(_nw10).ArraySet1(_84_cf56, 7)
			(_nw10).ArraySet1(p2, 8)
			(_nw10).ArraySet1(p3, 9)
			_125_v36 = _nw10
			_125_v36 = _125_v36
			var _126_v37 _dafny.MultiSet
			_ = _126_v37
			_126_v37 = _dafny.MultiSetOf(p0)
			var _127_v38 *C1
			_ = _127_v38
			var _nw11 *C1 = New_C1_()
			_ = _nw11
			_nw11.Ctor__(p0, _126_v37, _112_v25, _84_cf56, _85_v0)
			_127_v38 = _nw11
		}
		var _128_v39 _dafny.MultiSet
		_ = _128_v39
		_128_v39 = _dafny.MultiSetOf(p0)
		var _129_v40 *C2
		_ = _129_v40
		var _nw12 *C2 = New_C2_()
		_ = _nw12
		_nw12.Ctor__(_128_v39, (func() _dafny.Int {
			if p0 {
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p1, globalState), p0)).Update(Companion_Default___.Fm0(p3, globalState), p0)).Cardinality()
			}
			return _dafny.IntOfInt64(207)
		})())
		_129_v40 = _nw12
		_129_v40 = _129_v40
	} else if _source8.Is_DC32() {
		var _130___mcc_h1 _dafny.Int = _source8.Get_().(D13_DC32).Cf57
		_ = _130___mcc_h1
		var _131___mcc_h2 _dafny.Int = _source8.Get_().(D13_DC32).Cf58
		_ = _131___mcc_h2
		var _132___mcc_h3 _dafny.Sequence = _source8.Get_().(D13_DC32).Cf59
		_ = _132___mcc_h3
		var _133___mcc_h4 bool = _source8.Get_().(D13_DC32).Cf60
		_ = _133___mcc_h4
		var _134_cf60 bool = _133___mcc_h4
		_ = _134_cf60
		var _135_cf59 _dafny.Sequence = _132___mcc_h3
		_ = _135_cf59
		var _136_cf58 _dafny.Int = _131___mcc_h2
		_ = _136_cf58
		var _137_cf57 _dafny.Int = _130___mcc_h1
		_ = _137_cf57
		var _138_v41 _dafny.Sequence
		_ = _138_v41
		_138_v41 = _dafny.UnicodeSeqOfUtf8Bytes("jc")
		var _139_v42 _dafny.Array
		_ = _139_v42
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_1
		var _nw13 _dafny.Array
		_ = _nw13
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw13 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_140_cf57 _dafny.Int, _141_p0 bool) func(_dafny.Int) bool {
				return func(_142_i3 _dafny.Int) bool {
					return (_140_cf57).Cmp((Companion_D4_.Create_DC9_(_140_cf57, _141_p0)).Dtor_cf16()) == 0
				}
			})(_137_cf57, p0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw13 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw13).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw13).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_139_v42 = _nw13
		var _rhs2 _dafny.Int = _136_cf58
		_ = _rhs2
		var _rhs3 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fkc")
		_ = _rhs3
		var _rhs4 bool = false
		_ = _rhs4
		var _rhs5 _dafny.Array = _139_v42
		_ = _rhs5
		var _rhs6 _dafny.Int = (func() _dafny.Int {
			if false {
				return p3
			}
			return _136_cf58
		})()
		_ = _rhs6
		_136_cf58 = _rhs2
		_138_v41 = _rhs3
		_134_cf60 = _rhs4
		_139_v42 = _rhs5
		_136_cf58 = _rhs6
		(globalState).F15 = _136_cf58
		if !(_134_cf60) || (p0) {
			var _143_v43 _dafny.MultiSet
			_ = _143_v43
			_143_v43 = _dafny.MultiSetOf(p0)
			var _144_v44 _dafny.Array
			_ = _144_v44
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw14
			_144_v44 = _nw14
			var _145_v45 _dafny.Map
			_ = _145_v45
			_145_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_cf60, _144_v44)
			var _146_v46 T2
			_ = _146_v46
			var _nw15 *C6 = New_C6_()
			_ = _nw15
			_nw15.Ctor__(true, p2, _143_v43, _135_cf59, p3, (func() _dafny.Array {
				if (_145_v45).Contains(!(!(_134_cf60))) {
					return (_145_v45).Get(!(!(_134_cf60))).(_dafny.Array)
				}
				return _144_v44
			})(), _143_v43, p2)
			_146_v46 = _nw15
			var _147_v47 _dafny.Array
			_ = _147_v47
			var _nwElement0_4 T2 = _146_v46
			_ = _nwElement0_4
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(13))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_4, 0)
			(_nw16).ArraySet1(_146_v46, 1)
			(_nw16).ArraySet1(_146_v46, 2)
			(_nw16).ArraySet1(_146_v46, 3)
			(_nw16).ArraySet1(_146_v46, 4)
			(_nw16).ArraySet1(_146_v46, 5)
			(_nw16).ArraySet1(_146_v46, 6)
			(_nw16).ArraySet1(_146_v46, 7)
			(_nw16).ArraySet1(_146_v46, 8)
			(_nw16).ArraySet1(_146_v46, 9)
			(_nw16).ArraySet1(_146_v46, 10)
			(_nw16).ArraySet1(_146_v46, 11)
			(_nw16).ArraySet1(_146_v46, 12)
			_147_v47 = _nw16
			_147_v47 = _147_v47
			var _148_v48 *C0
			_ = _148_v48
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__(_137_cf57, _138_v41)
			_148_v48 = _nw17
			var _149_v49 _dafny.Sequence
			_ = _149_v49
			_149_v49 = _dafny.SeqOf(_148_v48, _148_v48)
			var _150_v50 _dafny.Sequence
			_ = _150_v50
			_150_v50 = _dafny.SeqOf(_149_v49)
			var _151_v51 _dafny.CodePoint
			_ = _151_v51
			_151_v51 = _dafny.CodePoint('w')
			var _152_v52 _dafny.Array
			_ = _152_v52
			var _nwElement0_5 _dafny.Int = (_146_v46).F32()
			_ = _nwElement0_5
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(10))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_5, 0)
			(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_150_v50, _150_v50)).Cardinality()), 1)
			(_nw18).ArraySet1((_148_v48).F36(), 2)
			(_nw18).ArraySet1((_148_v48).F36(), 3)
			(_nw18).ArraySet1(_137_cf57, 4)
			(_nw18).ArraySet1(_dafny.IntOfInt64(913), 5)
			(_nw18).ArraySet1((Companion_Default___.Fm56(Companion_D4_.Create_DC8_(_151_v51, _dafny.IntOfInt64(470), p0), p0, Companion_Default___.Fm57(p0, p0, globalState), globalState)).Cardinality(), 6)
			(_nw18).ArraySet1((_146_v46).F32(), 7)
			(_nw18).ArraySet1(_136_cf58, 8)
			(_nw18).ArraySet1(_dafny.IntOfInt64(543), 9)
			_152_v52 = _nw18
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))
			_ = _index4
			(_152_v52).ArraySet1(_137_cf57, (_index4).Int())
			var _153_v53 _dafny.Map
			_ = _153_v53
			_153_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _137_cf57)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))
			_ = _index5
			var _rhs7 _dafny.Sequence = Companion_Default___.Fm8(globalState)
			_ = _rhs7
			var _rhs8 _dafny.Int = (Companion_D9_.Create_DC20_(_136_cf58, !(_134_cf60), p0, p0, _153_v53)).Dtor_cf34()
			_ = _rhs8
			var _rhs9 bool = false
			_ = _rhs9
			var _rhs10 bool = !(p0)
			_ = _rhs10
			var _lhs2 _dafny.Array = _152_v52
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))
			_ = _lhs3
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			_138_v41 = _rhs7
			(_lhs2).ArraySet1(_rhs8, (_lhs3).Int())
			_lhs4.F26 = _rhs9
			_134_cf60 = _rhs10
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_139_v42), 0))
			_ = _index6
			(_139_v42).ArraySet1(p0, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_139_v42), 0))
			_ = _index7
			(_139_v42).ArraySet1(p0, (_index7).Int())
			var _154_v54 _dafny.MultiSet
			_ = _154_v54
			_154_v54 = _dafny.MultiSetOf(_dafny.IntOfUint32((_135_cf59).Cardinality()), (p2).Minus(p1), Companion_Default___.SafeDivisionInt((_152_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))).Int()).(_dafny.Int), (_148_v48).F36()), _dafny.IntOfInt64(-492))
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))
			_ = _index8
			var _rhs11 _dafny.MultiSet = (_154_v54).Union(_154_v54)
			_ = _rhs11
			var _rhs12 _dafny.Int = _dafny.IntOfUint32(((_148_v48).F37()).Cardinality())
			_ = _rhs12
			var _lhs5 _dafny.Array = _152_v52
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_152_v52), 0))
			_ = _lhs6
			_154_v54 = _rhs11
			(_lhs5).ArraySet1(_rhs12, (_lhs6).Int())
			(globalState).F26 = (_139_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_139_v42), 0))).Int()).(bool)
		} else {
			var _155_v55 _dafny.Map
			_ = _155_v55
			_155_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_138_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.IntOfUint32((_138_v41).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _156_v56 _dafny.Map
			_ = _156_v56
			_156_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _155_v55)
			var _157_v57 _dafny.Sequence
			_ = _157_v57
			_157_v57 = _dafny.SeqOf((func() _dafny.Map {
				if (_156_v56).Contains(_136_cf58) {
					return (_156_v56).Get(_136_cf58).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(678))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_158_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()), _dafny.CodePoint('t'))
			})())
			var _159_v58 D15
			_ = _159_v58
			_159_v58 = Companion_D15_.Create_DC35_(_157_v57)
			_159_v58 = (func() D15 {
				if p0 {
					return _159_v58
				}
				return Companion_D15_.Create_DC35_(_157_v57)
			})()
			(globalState).F18 = Companion_Default___.SafeModuloInt((_dafny.SetOf(_134_cf60)).Cardinality(), _dafny.IntOfInt64(-822))
			var _160_v59 D13
			_ = _160_v59
			_160_v59 = Companion_D13_.Create_DC32_(_dafny.IntOfInt64(14), p3, _135_cf59, _134_cf60)
			var _161_v60 _dafny.Sequence
			_ = _161_v60
			_161_v60 = _dafny.SeqOf(_160_v59)
			_161_v60 = _dafny.Companion_Sequence_.Concatenate(_161_v60, _dafny.SeqOf(_160_v59, _160_v59))
			var _162_v61 _dafny.CodePoint
			_ = _162_v61
			_162_v61 = _dafny.CodePoint('o')
			var _163_v62 _dafny.Array
			_ = _163_v62
			var _nwElement0_6 _dafny.CodePoint = _162_v61
			_ = _nwElement0_6
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(12))
			_ = _nw19
			(_nw19).ArraySet1CodePoint(_nwElement0_6, 0)
			(_nw19).ArraySet1CodePoint(_162_v61, 1)
			(_nw19).ArraySet1CodePoint(_dafny.CodePoint('d'), 2)
			(_nw19).ArraySet1CodePoint(_dafny.CodePoint('u'), 3)
			(_nw19).ArraySet1CodePoint(_162_v61, 4)
			(_nw19).ArraySet1CodePoint(_dafny.CodePoint('p'), 5)
			(_nw19).ArraySet1CodePoint(_162_v61, 6)
			(_nw19).ArraySet1CodePoint(_162_v61, 7)
			(_nw19).ArraySet1CodePoint(_162_v61, 8)
			(_nw19).ArraySet1CodePoint(_162_v61, 9)
			(_nw19).ArraySet1CodePoint(_162_v61, 10)
			(_nw19).ArraySet1CodePoint(_162_v61, 11)
			_163_v62 = _nw19
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_163_v62), 0))
			_ = _index9
			(_163_v62).ArraySet1CodePoint(_162_v61, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_163_v62), 0))
			_ = _index10
			(_163_v62).ArraySet1CodePoint(_162_v61, (_index10).Int())
			(globalState).F21 = (_dafny.Zero).Minus(p3)
		}
		r0 = p0
	} else if _source8.Is_DC30() {
		var _164___mcc_h5 *C1 = _source8.Get_().(D13_DC30).Cf55
		_ = _164___mcc_h5
		var _165_cf55 *C1 = _164___mcc_h5
		_ = _165_cf55
		var _166_v63 _dafny.Sequence
		_ = _166_v63
		_166_v63 = _dafny.SeqOf(p3)
		if !(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_166_v63, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3, p3, (_dafny.Zero).Minus(p3)), _166_v63)))) {
			var _167_v64 _dafny.Sequence
			_ = _167_v64
			_167_v64 = _dafny.SeqOf(p0, _165_cf55.F35)
			var _168_v65 _dafny.Sequence
			_ = _168_v65
			_168_v65 = _dafny.UnicodeSeqOfUtf8Bytes("qmttqthg")
			var _169_v66 _dafny.Array
			_ = _169_v66
			var _nwElement0_7 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-212))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_170_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))
			_ = _nwElement0_7
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(27))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_7, 0)
			(_nw20).ArraySet1(_168_v65, 1)
			(_nw20).ArraySet1(_168_v65, 2)
			(_nw20).ArraySet1(_168_v65, 3)
			(_nw20).ArraySet1(_168_v65, 4)
			(_nw20).ArraySet1(_168_v65, 5)
			(_nw20).ArraySet1(_168_v65, 6)
			(_nw20).ArraySet1(_168_v65, 7)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cfdjji"), 8)
			(_nw20).ArraySet1(_168_v65, 9)
			(_nw20).ArraySet1(_168_v65, 10)
			(_nw20).ArraySet1(_168_v65, 11)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("eet"), 12)
			(_nw20).ArraySet1(_168_v65, 13)
			(_nw20).ArraySet1(_168_v65, 14)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tqv"), 15)
			(_nw20).ArraySet1(_168_v65, 16)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fwkfera"), 17)
			(_nw20).ArraySet1(_168_v65, 18)
			(_nw20).ArraySet1(_168_v65, 19)
			(_nw20).ArraySet1(_168_v65, 20)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("y"), 21)
			(_nw20).ArraySet1(_168_v65, 22)
			(_nw20).ArraySet1(_168_v65, 23)
			(_nw20).ArraySet1(_168_v65, 24)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ldwxsm"), 25)
			(_nw20).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fytgb"), 26)
			_169_v66 = _nw20
			var _171_v67 _dafny.MultiSet
			_ = _171_v67
			_171_v67 = _dafny.MultiSetOf(_165_cf55.F35, p0, p0, p0)
			var _172_v68 *C6
			_ = _172_v68
			var _nw21 *C6 = New_C6_()
			_ = _nw21
			_nw21.Ctor__((_165_cf55).Fm3(_dafny.MultiSetOf(p0), globalState), p2, _dafny.MultiSetFromSeq(_167_v64), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(_dafny.IntOfInt64(511), globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm16(_dafny.IntOfInt64(511), globalState)).Cardinality()))).Uint32(), p2), (p1).Minus(p2), _169_v66, (_171_v67).Update(true, Companion_Default___.Abs((_171_v67).Cardinality())), _dafny.IntOfUint32((_168_v65).Cardinality()))
			_172_v68 = _nw21
			var _173_v69 *C4
			_ = _173_v69
			var _nw22 *C4 = New_C4_()
			_ = _nw22
			_nw22.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_166_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-21))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_174_v65 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_175_i6 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_174_v65).Cardinality()))
				}
			})(_168_v65))))).Cardinality()), _169_v66)
			_173_v69 = _nw22
			_173_v69 = _173_v69
			var _176_v70 _dafny.Sequence
			_ = _176_v70
			_176_v70 = _168_v65
			var _177_v71 _dafny.Set
			_ = _177_v71
			_177_v71 = _dafny.SetOf((_165_cf55).Fm11(false, _172_v68.F34, _176_v70, p3, globalState))
			var _178_v72 _dafny.Map
			_ = _178_v72
			_178_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, (_177_v71).IsSubsetOf(_177_v71))
			_178_v72 = (_178_v72).Update((_172_v68).F33(), !(_165_cf55.F35))
			var _179_v73 _dafny.CodePoint
			_ = _179_v73
			_179_v73 = _dafny.CodePoint('n')
			_179_v73 = _179_v73
			var _180_v74 _dafny.Map
			_ = _180_v74
			_180_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_168_v65).Cardinality()), (_dafny.IntOfInt64(926)).Times(p1))
			_180_v74 = (_180_v74).Update(((_180_v74).Update((_dafny.Zero).Minus(p3), _dafny.IntOfInt64(93))).Cardinality(), _172_v68.F34)
		} else {
			var _181_v75 _dafny.MultiSet
			_ = _181_v75
			_181_v75 = _dafny.MultiSetOf(_165_cf55.F35)
			var _182_v76 _dafny.Array
			_ = _182_v76
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_2
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = func(_183_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("mte")
				}
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw23 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw23).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw23).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_182_v76 = _nw23
			var _184_v77 T2
			_ = _184_v77
			var _nw24 *C6 = New_C6_()
			_ = _nw24
			_nw24.Ctor__(_165_cf55.F35, p3, _181_v75, _dafny.SeqOf(p1), _dafny.IntOfUint32((_166_v63).Cardinality()), _182_v76, _181_v75, p3)
			_184_v77 = _nw24
			var _185_v78 _dafny.Sequence
			_ = _185_v78
			_185_v78 = _dafny.SeqOf(_165_cf55.F35)
			var _186_v79 _dafny.Map
			_ = _186_v79
			_186_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v77, _dafny.IntOfUint32((_185_v78).Cardinality()))
			var _187_v80 _dafny.Map
			_ = _187_v80
			_187_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, _186_v79)
			var _188_v81 _dafny.MultiSet
			_ = _188_v81
			_188_v81 = _dafny.MultiSetOf(_186_v79, (func() _dafny.Map {
				if (_187_v80).Contains(p0) {
					return (_187_v80).Get(p0).(_dafny.Map)
				}
				return _186_v79
			})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v77, p3))
			var _189_v82 T1
			_ = _189_v82
			var _nw25 *C1 = New_C1_()
			_ = _nw25
			_nw25.Ctor__(_165_cf55.F35, _dafny.MultiSetOf(p0), _166_v63, (_188_v81).Cardinality(), _182_v76)
			_189_v82 = _nw25
			var _190_v83 _dafny.CodePoint
			_ = _190_v83
			_190_v83 = _dafny.CodePoint('u')
			var _rhs13 bool = true
			_ = _rhs13
			var _rhs14 bool = Companion_Default___.Fm0((_189_v82).Fm4(_dafny.SetOf(_dafny.SeqOf(_190_v83)), p0, _185_v78, globalState), globalState)
			_ = _rhs14
			var _rhs15 T1 = _189_v82
			_ = _rhs15
			var _lhs7 *C1 = _165_cf55
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			_lhs7.F35 = _rhs13
			_lhs8.F26 = _rhs14
			_189_v82 = _rhs15
			var _191_v84 _dafny.Array
			_ = _191_v84
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw26
			_191_v84 = _nw26
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_191_v84), 0))
			_ = _index11
			(_191_v84).ArraySet1((p2).Cmp((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(446), _dafny.IntOfInt64(851))); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _192_v85 _dafny.Int
					_192_v85 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(446)).Cmp(_192_v85) <= 0) && ((_192_v85).Cmp(_dafny.IntOfInt64(851)) < 0) {
						_coll17.Add((_192_v85).Minus(_dafny.IntOfInt64(531)), p0)
					}
				}
				return _coll17.ToMap()
			}()).Cardinality()) >= 0, (_index11).Int())
			var _193_v86 _dafny.Map
			_ = _193_v86
			_193_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, _dafny.IntOfInt64(-355))
			var _194_v87 _dafny.Sequence
			_ = _194_v87
			_194_v87 = _dafny.UnicodeSeqOfUtf8Bytes("ostrgckj")
			var _195_v88 _dafny.Sequence
			_ = _195_v88
			_195_v88 = _194_v87
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_191_v84), 0))
			_ = _index12
			(_191_v84).ArraySet1(!((_165_cf55).Fm11(true, ((_193_v86).Cardinality()).Minus((_189_v82).F27()), _195_v88, (_184_v77).F32(), globalState)), (_index12).Int())
			r0 = p0
			var _196_v89 _dafny.Array
			_ = _196_v89
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_3
			var _nw27 _dafny.Array
			_ = _nw27
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw27 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.MultiSet = (func(_197_v77 T2, _198_v82 T1, _199_v75 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_200_i8 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf((_197_v77).F31(), (_198_v82).F29(), _199_v75, (_197_v77).F31())
					}
				})(_184_v77, _189_v82, _181_v75)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw27 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw27).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw27).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_196_v89 = _nw27
			var _201_v90 _dafny.MultiSet
			_ = _201_v90
			_201_v90 = _dafny.MultiSetOf(_181_v75)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_196_v89), 0))
			_ = _index13
			(_196_v89).ArraySet1(_201_v90, (_index13).Int())
			var _202_v91 _dafny.Map
			_ = _202_v91
			_202_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_194_v87).Cardinality()), _201_v90)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_196_v89), 0))
			_ = _index14
			(_196_v89).ArraySet1(((func() _dafny.MultiSet {
				if (_202_v91).Contains(p3) {
					return (_202_v91).Get(p3).(_dafny.MultiSet)
				}
				return _201_v90
			})()).Intersection(_201_v90), (_index14).Int())
			(globalState).F3 = ((_193_v86).Cardinality()).Minus(_dafny.IntOfInt64(904))
		}
		var _203_v92 _dafny.Array
		_ = _203_v92
		var _nw28 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw28
		_203_v92 = _nw28
		var _204_v93 _dafny.MultiSet
		_ = _204_v93
		_204_v93 = _dafny.MultiSetOf(true, !(_165_cf55.F35))
		var _205_v94 _dafny.Sequence
		_ = _205_v94
		_205_v94 = _dafny.SeqOf(_204_v93, _204_v93)
		var _206_v95 *C2
		_ = _206_v95
		var _nw29 *C2 = New_C2_()
		_ = _nw29
		_nw29.Ctor__((_205_v94).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_205_v94).Cardinality()))).Uint32()).(_dafny.MultiSet), p3)
		_206_v95 = _nw29
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_203_v92), 0))
		_ = _index15
		(_203_v92).ArraySet1(_206_v95, (_index15).Int())
		var _207_v96 _dafny.Sequence
		_ = _207_v96
		_207_v96 = _dafny.UnicodeSeqOfUtf8Bytes("tytaticn")
		var _208_v97 _dafny.Map
		_ = _208_v97
		_208_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, p1)
		var _209_v98 _dafny.CodePoint
		_ = _209_v98
		_209_v98 = _dafny.CodePoint('l')
		var _210_v99 _dafny.Set
		_ = _210_v99
		_210_v99 = _dafny.SetOf(_dafny.SeqOf(_209_v98))
		var _211_v100 _dafny.Sequence
		_ = _211_v100
		_211_v100 = _dafny.SeqOf(Companion_Default___.Fm0(p1, globalState))
		var _212_v101 _dafny.Array
		_ = _212_v101
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
		_ = _nw30
		_212_v101 = _nw30
		var _213_v102 *C3
		_ = _213_v102
		var _nw31 *C3 = New_C3_()
		_ = _nw31
		_nw31.Ctor__(p1, _212_v101)
		_213_v102 = _nw31
		var _214_v103 _dafny.Map
		_ = _214_v103
		_214_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_208_v97).Contains(false) {
				return (_208_v97).Get(false).(_dafny.Int)
			}
			return (_204_v93).Cardinality()
		})(), _213_v102)
		var _215_v104 _dafny.MultiSet
		_ = _215_v104
		_215_v104 = _dafny.MultiSetOf((_dafny.IntOfUint32((_207_v96).Cardinality())).Times((_208_v97).Cardinality()), p2, p1, (_165_cf55).Fm4(_210_v99, false, _211_v100, globalState), (_214_v103).Cardinality())
		var _216_v105 _dafny.Set
		_ = _216_v105
		_216_v105 = _dafny.SetOf(p0, _165_cf55.F35)
		var _217_v106 _dafny.Set
		_ = _217_v106
		_217_v106 = _dafny.SetOf(p1)
		var _218_v107 _dafny.Map
		_ = _218_v107
		_218_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, p0)
		var _219_v108 D10
		_ = _219_v108
		_219_v108 = Companion_D10_.Create_DC22_(_dafny.IntOfUint32((_207_v96).Cardinality()), _204_v93)
		var _220_v109 _dafny.MultiSet
		_ = _220_v109
		_220_v109 = _dafny.MultiSetOf(_216_v105, _216_v105)
		var _221_v110 _dafny.Map
		_ = _221_v110
		_221_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
		var _222_v111 _dafny.Array
		_ = _222_v111
		var _nwElement0_8 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, p1)
		_ = _nwElement0_8
		var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(21))
		_ = _nw32
		(_nw32).ArraySet1(_nwElement0_8, 0)
		(_nw32).ArraySet1(p1, 1)
		(_nw32).ArraySet1(p3, 2)
		(_nw32).ArraySet1(p1, 3)
		(_nw32).ArraySet1(p2, 4)
		(_nw32).ArraySet1((Companion_Default___.Fm1(globalState)).Times((_216_v105).Cardinality()), 5)
		(_nw32).ArraySet1(p3, 6)
		(_nw32).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (_217_v106).Cardinality()), 7)
		(_nw32).ArraySet1(p2, 8)
		(_nw32).ArraySet1((_218_v107).Cardinality(), 9)
		(_nw32).ArraySet1(p1, 10)
		(_nw32).ArraySet1(_dafny.IntOfInt64(307), 11)
		(_nw32).ArraySet1(p1, 12)
		(_nw32).ArraySet1(p2, 13)
		(_nw32).ArraySet1((_219_v108).Dtor_cf40(), 14)
		(_nw32).ArraySet1(p1, 15)
		(_nw32).ArraySet1(((_220_v109).Update(_216_v105, Companion_Default___.Abs((func() _dafny.Int {
			if (_204_v93).Contains(p0) {
				return (_204_v93).Multiplicity(p0)
			}
			return p1
		})()))).Cardinality(), 16)
		(_nw32).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm1(globalState)), 17)
		(_nw32).ArraySet1((_dafny.Zero).Minus((p2).Plus(p1)), 18)
		(_nw32).ArraySet1(p2, 19)
		(_nw32).ArraySet1((p2).Minus((_221_v110).Cardinality()), 20)
		_222_v111 = _nw32
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))
		_ = _index16
		(_222_v111).ArraySet1(Companion_Default___.SafeModuloInt((_208_v97).Cardinality(), p2), (_index16).Int())
		var _223_v112 _dafny.Array
		_ = _223_v112
		var _nw33 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
		_ = _nw33
		_223_v112 = _nw33
		var _224_v113 _dafny.Sequence
		_ = _224_v113
		_224_v113 = _dafny.SeqOf(_223_v112)
		var _225_v114 _dafny.Map
		_ = _225_v114
		_225_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _207_v96)
		var _226_v115 D22
		_ = _226_v115
		_226_v115 = Companion_D22_.Create_DC53_(_224_v113)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_203_v92), 0))
		_ = _index17
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))
		_ = _index18
		var _rhs16 *C2 = _206_v95
		_ = _rhs16
		var _rhs17 _dafny.MultiSet = ((_215_v104).Intersection(_215_v104)).Difference(_215_v104)
		_ = _rhs17
		var _rhs18 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if !(_165_cf55.F35) || (p0) {
				return _166_v63
			}
			return _dafny.Companion_Sequence_.Concatenate(_166_v63, _166_v63)
		})()).Cardinality())
		_ = _rhs18
		var _rhs19 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _207_v96)).Merge((_225_v114).Merge(_225_v114))).Cardinality()
		_ = _rhs19
		var _rhs20 _dafny.Sequence = (_226_v115).Dtor_cf94()
		_ = _rhs20
		var _lhs9 _dafny.Array = _203_v92
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_203_v92), 0))
		_ = _lhs10
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		var _lhs12 _dafny.Array = _222_v111
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))
		_ = _lhs13
		(_lhs9).ArraySet1(_rhs16, (_lhs10).Int())
		_215_v104 = _rhs17
		_lhs11.F3 = _rhs18
		(_lhs12).ArraySet1(_rhs19, (_lhs13).Int())
		_224_v113 = _rhs20
		var _227_v116 *C6
		_ = _227_v116
		var _nw34 *C6 = New_C6_()
		_ = _nw34
		_nw34.Ctor__(!(_165_cf55.F35), _dafny.IntOfUint32((_207_v96).Cardinality()), _204_v93, Companion_Default___.Fm16(p1, globalState), (_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int), _212_v101, _204_v93, p3)
		_227_v116 = _nw34
		var _228_v117 _dafny.Map
		_ = _228_v117
		_228_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v116, (_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int))
		var _229_v118 _dafny.Map
		_ = _229_v118
		_229_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v111, _228_v117)
		_229_v118 = (_229_v118).Update(_222_v111, (_228_v117).Update(_227_v116, (_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int)))
		var _230_v119 D4
		_ = _230_v119
		_230_v119 = Companion_D4_.Create_DC9_(_dafny.IntOfInt64(-28), p0)
		var _231_v120 _dafny.Array
		_ = _231_v120
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw35
		_231_v120 = _nw35
		var _232_v121 _dafny.Array
		_ = _232_v121
		_232_v121 = _231_v120
		var _source9 _dafny.Array = (func() _dafny.Array {
			if !(((_230_v119).Dtor_cf16()).Cmp((_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int)) >= 0) {
				return _232_v121
			}
			return _232_v121
		})()
		_ = _source9
		var _233___mcc_h7 _dafny.Array = _source9
		_ = _233___mcc_h7
		var _234_cf93 _dafny.Array = _233___mcc_h7
		_ = _234_cf93
		var _235_v122 _dafny.Array
		_ = _235_v122
		var _len0_4 _dafny.Int = _dafny.One
		_ = _len0_4
		var _nw36 _dafny.Array
		_ = _nw36
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw36 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Map = (func(_236_p0 bool, _237_v98 _dafny.CodePoint, _238_v116 *C6) func(_dafny.Int) _dafny.Map {
				return func(_239_i9 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_p0, Companion_D4_.Create_DC8_(_237_v98, (_dafny.Zero).Minus(_238_v116.F34), _236_p0))
				}
			})(p0, _209_v98, _227_v116)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw36 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw36).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw36).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_235_v122 = _nw36
		var _240_v123 D4
		_ = _240_v123
		_240_v123 = Companion_D4_.Create_DC8_(_209_v98, (_218_v107).Cardinality(), (_227_v116).F33())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_235_v122), 0))
		_ = _index19
		(_235_v122).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _240_v123), (_index19).Int())
		var _241_v124 _dafny.Map
		_ = _241_v124
		_241_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55.F35, _240_v123)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_235_v122), 0))
		_ = _index20
		(_235_v122).ArraySet1(((_241_v124).Update(_165_cf55.F35, _240_v123)).Update(!(p0) || (false), _240_v123), (_index20).Int())
		var _242_v125 *C0
		_ = _242_v125
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__(p1, (func() _dafny.Sequence {
			if (_225_v114).Contains(!(_165_cf55.F35)) {
				return (_225_v114).Get(!(_165_cf55.F35)).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_243_v98 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_244_i10 _dafny.Int) _dafny.CodePoint {
					return _243_v98
				}
			})(_209_v98)))
		})())
		_242_v125 = _nw37
		var _245_v126 _dafny.Sequence
		_ = _245_v126
		_245_v126 = _dafny.SeqOf((_203_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_203_v92), 0))).Int()).(*C2), _206_v95)
		var _246_v127 _dafny.Map
		_ = _246_v127
		_246_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kqkukel")).Cardinality()))
		var _247_v128 *C5
		_ = _247_v128
		var _nw38 *C5 = New_C5_()
		_ = _nw38
		_nw38.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_246_v127, p1), _204_v93, _166_v63, _dafny.IntOfUint32((_207_v96).Cardinality()), _212_v101)
		_247_v128 = _nw38
		var _248_v129 _dafny.MultiSet
		_ = _248_v129
		_248_v129 = _dafny.MultiSetOf(_247_v128)
		var _nw39 *C1 = New_C1_()
		_ = _nw39
		_nw39.Ctor__(_dafny.Companion_Sequence_.Equal(_245_v126, _245_v126), _204_v93, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-554), (_225_v114).Cardinality(), (_248_v129).Cardinality()), _dafny.SeqOf(p2)), _dafny.IntOfUint32((_166_v63).Cardinality()), _212_v101)
		_165_cf55 = _nw39
		var _249_v130 _dafny.Array
		_ = _249_v130
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_5
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Map = (func(_250_v98 _dafny.CodePoint, _251_v63 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
				return func(_252_i11 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v98, _251_v63)
				}
			})(_209_v98, _166_v63)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw40).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw40).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_249_v130 = _nw40
		var _253_v131 D2
		_ = _253_v131
		_253_v131 = Companion_D2_.Create_DC4_(p2, (_222_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_222_v111), 0))).Int()).(_dafny.Int), (_242_v125).F36(), _209_v98, _207_v96)
		var _254_v132 _dafny.Map
		_ = _254_v132
		_254_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_253_v131).Dtor_cf8(), Companion_Default___.Fm43((_227_v116).F33(), p0, !(p0), globalState))
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_249_v130), 0))
		_ = _index21
		(_249_v130).ArraySet1(_254_v132, (_index21).Int())
		var _255_v133 _dafny.Map
		_ = _255_v133
		_255_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_cf55, _254_v132)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_249_v130), 0))
		_ = _index22
		(_249_v130).ArraySet1((func() _dafny.Map {
			if (_255_v133).Contains(_165_cf55) {
				return (_255_v133).Get(_165_cf55).(_dafny.Map)
			}
			return _254_v132
		})(), (_index22).Int())
	} else {
		var _256___mcc_h6 D13 = _source8.Get_().(D13_DC33).Cf61
		_ = _256___mcc_h6
		var _257_cf61 D13 = _256___mcc_h6
		_ = _257_cf61
		(globalState).F26 = p0
		var _258_v134 _dafny.Array
		_ = _258_v134
		var _nw41 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
		_ = _nw41
		_258_v134 = _nw41
		var _259_v135 _dafny.Set
		_ = _259_v135
		_259_v135 = _dafny.SetOf(_258_v134, _258_v134, _258_v134)
		if !(p0) || ((_259_v135).IsProperSubsetOf(_259_v135)) {
			(globalState).F26 = !(false)
			var _260_v136 *C0
			_ = _260_v136
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__(p1, Companion_Default___.Fm8(globalState))
			_260_v136 = _nw42
			var _261_v137 D6
			_ = _261_v137
			_261_v137 = Companion_D6_.Create_DC11_(_260_v136)
			var _262_v138 D19
			_ = _262_v138
			_262_v138 = Companion_D19_.Create_DC49_(p0, p0, _dafny.IntOfInt64(807))
			var _263_v139 _dafny.Map
			_ = _263_v139
			_263_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v137, _262_v138)
			_263_v139 = (_263_v139).Update(_261_v137, _262_v138)
			var _264_v140 _dafny.Array
			_ = _264_v140
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw43
			_264_v140 = _nw43
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_264_v140), 0))
			_ = _index23
			(_264_v140).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_260_v136).F37(), (_260_v136).F37())).Cardinality()), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_264_v140), 0))
			_ = _index24
			(_264_v140).ArraySet1(p1, (_index24).Int())
			var _265_v141 _dafny.Map
			_ = _265_v141
			_265_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_260_v136).F36()).Times(p3), p0)
			_265_v141 = (_265_v141).Update(p2, p0)
			var _266_v142 _dafny.Sequence
			_ = _266_v142
			_266_v142 = _dafny.UnicodeSeqOfUtf8Bytes("dsshj")
			var _267_v143 _dafny.Set
			_ = _267_v143
			_267_v143 = _dafny.SetOf((_260_v136).F36(), p2)
			var _268_v144 _dafny.Map
			_ = _268_v144
			_268_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_267_v143).Intersection(_267_v143))
			var _269_v145 _dafny.CodePoint
			_ = _269_v145
			_269_v145 = _dafny.CodePoint('e')
			var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_260_v136).F37(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_260_v136).F37()).Cardinality()))).Uint32(), _269_v145), (_260_v136).F37())
			_ = _rhs21
			var _rhs22 _dafny.Map = _268_v144
			_ = _rhs22
			_266_v142 = _rhs21
			_268_v144 = _rhs22
		} else {
			var _270_v146 _dafny.Sequence
			_ = _270_v146
			_270_v146 = _dafny.SeqOf((p1).Cmp(p2) > 0, p0, p0, p0, p0)
			var _271_v147 D4
			_ = _271_v147
			_271_v147 = Companion_D4_.Create_DC9_(_dafny.IntOfInt64(534), p0)
			var _rhs23 _dafny.Sequence = _dafny.SeqOf(p0, p0, p0)
			_ = _rhs23
			var _rhs24 bool = ((func() D4 {
				if p0 {
					return _271_v147
				}
				return Companion_Default___.Fm49(p1, globalState)
			})()).Dtor_cf17()
			_ = _rhs24
			var _rhs25 _dafny.Int = p3
			_ = _rhs25
			var _lhs14 *GlobalState = globalState
			_ = _lhs14
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			_270_v146 = _rhs23
			_lhs14.F26 = _rhs24
			_lhs15.F3 = _rhs25
			var _272_v148 D10
			_ = _272_v148
			_272_v148 = Companion_D10_.Create_DC22_(p3, _dafny.MultiSetFromSeq(_270_v146))
			var _273_v149 _dafny.Sequence
			_ = _273_v149
			_273_v149 = _dafny.SeqOf(p3)
			var _274_v150 _dafny.Array
			_ = _274_v150
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw44
			_274_v150 = _nw44
			var _275_v151 *C1
			_ = _275_v151
			var _nw45 *C1 = New_C1_()
			_ = _nw45
			_nw45.Ctor__(true, (_272_v148).Dtor_cf41(), _273_v149, p2, _274_v150)
			_275_v151 = _nw45
			var _276_v152 _dafny.Map
			_ = _276_v152
			_276_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _275_v151)
			var _277_v153 _dafny.Sequence
			_ = _277_v153
			_277_v153 = _dafny.UnicodeSeqOfUtf8Bytes("f")
			_276_v152 = (_276_v152).Update(_dafny.IntOfUint32((_277_v153).Cardinality()), _275_v151)
			var _278_v154 _dafny.MultiSet
			_ = _278_v154
			_278_v154 = _dafny.MultiSetOf(_275_v151.F35, p0, _275_v151.F35, true)
			var _279_v155 _dafny.Map
			_ = _279_v155
			_279_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v154, p2)
			var _280_v156 _dafny.Map
			_ = _280_v156
			_280_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v155, !(p0))
			_280_v156 = (_280_v156).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v154, _dafny.IntOfInt64(231))).Merge(_279_v155), p0)
			var _281_v157 *C3
			_ = _281_v157
			var _nw46 *C3 = New_C3_()
			_ = _nw46
			_nw46.Ctor__(_dafny.IntOfInt64(-596), _274_v150)
			_281_v157 = _nw46
			var _282_v158 _dafny.Array
			_ = _282_v158
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_6
			var _nw47 _dafny.Array
			_ = _nw47
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw47 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Map = (func(_283_p0 bool, _284_p3 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_285_i12 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_p0, _284_p3)
					}
				})(p0, p3)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw47 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw47).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw47).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_282_v158 = _nw47
			var _rhs26 *C3 = _281_v157
			_ = _rhs26
			var _rhs27 _dafny.Array = _282_v158
			_ = _rhs27
			_281_v157 = _rhs26
			_282_v158 = _rhs27
			var _286_v159 _dafny.Set
			_ = _286_v159
			_286_v159 = _dafny.SetOf(_dafny.IntOfUint32((_270_v146).Cardinality()), p3, (_dafny.SetOf((_278_v154).Update(p0, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}(func(_287_i13 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			}))).Cardinality()))))).Cardinality())
			var _288_v160 _dafny.CodePoint
			_ = _288_v160
			_288_v160 = _dafny.CodePoint('r')
			var _289_v161 _dafny.Map
			_ = _289_v161
			_289_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_288_v160, _286_v159)
			_286_v159 = (func() _dafny.Set {
				if (_289_v161).Contains(_288_v160) {
					return (_289_v161).Get(_288_v160).(_dafny.Set)
				}
				return _286_v159
			})()
		}
		var _290_v162 _dafny.MultiSet
		_ = _290_v162
		_290_v162 = _dafny.MultiSetOf(p0, p0, p0, false)
		var _291_v163 _dafny.Map
		_ = _291_v163
		_291_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v162, p0)
		var _292_v164 _dafny.Sequence
		_ = _292_v164
		_292_v164 = _dafny.SeqOf(p0, p0)
		var _293_v165 _dafny.Set
		_ = _293_v165
		_293_v165 = _dafny.SetOf(p3)
		var _294_v166 _dafny.Map
		_ = _294_v166
		_294_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
		var _295_v167 _dafny.Array
		_ = _295_v167
		var _nwElement0_9 _dafny.Int = p1
		_ = _nwElement0_9
		var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(29))
		_ = _nw48
		(_nw48).ArraySet1(_nwElement0_9, 0)
		(_nw48).ArraySet1(p2, 1)
		(_nw48).ArraySet1(p3, 2)
		(_nw48).ArraySet1(p2, 3)
		(_nw48).ArraySet1(p2, 4)
		(_nw48).ArraySet1(p3, 5)
		(_nw48).ArraySet1(p3, 6)
		(_nw48).ArraySet1((_291_v163).Cardinality(), 7)
		(_nw48).ArraySet1(p1, 8)
		(_nw48).ArraySet1(_dafny.IntOfUint32((_292_v164).Cardinality()), 9)
		(_nw48).ArraySet1(p1, 10)
		(_nw48).ArraySet1(_dafny.IntOfInt64(37), 11)
		(_nw48).ArraySet1(p1, 12)
		(_nw48).ArraySet1(p3, 13)
		(_nw48).ArraySet1(p3, 14)
		(_nw48).ArraySet1(p1, 15)
		(_nw48).ArraySet1(p1, 16)
		(_nw48).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-917)), 17)
		(_nw48).ArraySet1((_293_v165).Cardinality(), 18)
		(_nw48).ArraySet1(p2, 19)
		(_nw48).ArraySet1(p2, 20)
		(_nw48).ArraySet1((_293_v165).Cardinality(), 21)
		(_nw48).ArraySet1(_dafny.IntOfInt64(-608), 22)
		(_nw48).ArraySet1(_dafny.IntOfInt64(877), 23)
		(_nw48).ArraySet1(p3, 24)
		(_nw48).ArraySet1((_294_v166).Cardinality(), 25)
		(_nw48).ArraySet1(p3, 26)
		(_nw48).ArraySet1(p3, 27)
		(_nw48).ArraySet1(p2, 28)
		_295_v167 = _nw48
		var _296_v168 _dafny.Set
		_ = _296_v168
		_296_v168 = _dafny.SetOf(p0, p0)
		var _297_v169 _dafny.Map
		_ = _297_v169
		_297_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v168, (_dafny.SetOf(true)).Cardinality())
		var _298_v170 _dafny.Map
		_ = _298_v170
		_298_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_295_v167, (_297_v169).Cardinality())
		var _299_v171 D23
		_ = _299_v171
		_299_v171 = Companion_D23_.Create_DC56_(_298_v170)
		(globalState).F18 = ((((_299_v171).Dtor_cf104()).Merge(_298_v170)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_295_v167, p3))).Cardinality()
		var _300_v172 _dafny.Map
		_ = _300_v172
		_300_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_292_v164).Cardinality()), p1)
		var _301_v173 _dafny.Map
		_ = _301_v173
		_301_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v172, p3)
		var _302_v174 _dafny.Sequence
		_ = _302_v174
		_302_v174 = _dafny.UnicodeSeqOfUtf8Bytes("b")
		var _303_v175 _dafny.Sequence
		_ = _303_v175
		_303_v175 = _dafny.SeqOf(_dafny.IntOfInt64(590), _dafny.IntOfUint32((_302_v174).Cardinality()), p1)
		var _304_v176 _dafny.Array
		_ = _304_v176
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw49
		_304_v176 = _nw49
		var _305_v177 _dafny.Sequence
		_ = _305_v177
		_305_v177 = _dafny.SeqOf(_304_v176)
		var _306_v178 *C3
		_ = _306_v178
		var _nw50 *C3 = New_C3_()
		_ = _nw50
		_nw50.Ctor__(p1, _304_v176)
		_306_v178 = _nw50
		var _307_v179 _dafny.Sequence
		_ = _307_v179
		_307_v179 = _dafny.SeqOf(_306_v178)
		var _308_v180 *C5
		_ = _308_v180
		var _nw51 *C5 = New_C5_()
		_ = _nw51
		_nw51.Ctor__((_301_v173).Merge(_301_v173), (_dafny.MultiSetOf(p0, p0)).Union(_290_v162), _303_v175, p2, (_305_v177).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_307_v179).Cardinality()), _dafny.IntOfUint32((_305_v177).Cardinality()))).Uint32()).(_dafny.Array))
		_308_v180 = _nw51
	}
	var _309_v181 _dafny.Sequence
	_ = _309_v181
	_309_v181 = _dafny.UnicodeSeqOfUtf8Bytes("ym")
	var _310_v182 _dafny.CodePoint
	_ = _310_v182
	_310_v182 = _dafny.CodePoint('l')
	var _311_v183 _dafny.MultiSet
	_ = _311_v183
	_311_v183 = _dafny.MultiSetOf(_309_v181, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("w"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))).Uint32(), _310_v182), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("w"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))).Uint32(), _310_v182)).Cardinality()))).Uint32(), _310_v182))
	var _312_v184 _dafny.Sequence
	_ = _312_v184
	_312_v184 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fuf"))
	var _313_i14 _dafny.Int
	_ = _313_i14
	_313_i14 = _dafny.Zero
	{
		for (_dafny.MultiSetFromSeq(_312_v184)).IsProperSubsetOf(_311_v183) {
			{
				if (_313_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_313_i14 = (_313_i14).Plus(_dafny.One)
				var _314_v186 _dafny.Set
				_ = _314_v186
				_314_v186 = _dafny.SetOf(p3, p1)
				var _315_v187 _dafny.Sequence
				_ = _315_v187
				_315_v187 = _dafny.SeqOf(func() _dafny.Set {
					var _coll18 = _dafny.NewBuilder()
					_ = _coll18
					for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-577), _dafny.IntOfInt64(565))); ; {
						_compr_18, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _316_v185 _dafny.Int
						_316_v185 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(-577)).Cmp(_316_v185) <= 0) && ((_316_v185).Cmp(_dafny.IntOfInt64(565)) < 0) {
							_coll18.Add((_316_v185).Plus(p3))
						}
					}
					return _coll18.ToSet()
				}(), _314_v186)
				var _317_v188 D11
				_ = _317_v188
				_317_v188 = Companion_D11_.Create_DC25_(Companion_D11_.Create_DC23_(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_315_v187, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_315_v187).Cardinality()))).Uint32(), _314_v186), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_315_v187, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_315_v187).Cardinality()))).Uint32(), _314_v186)).Cardinality()))).Uint32(), _314_v186))))
				var _318_v189 D4
				_ = _318_v189
				_318_v189 = Companion_D4_.Create_DC8_(_310_v182, _dafny.IntOfInt64(-919), p0)
				var _319_v190 D11
				_ = _319_v190
				_319_v190 = Companion_D11_.Create_DC24_(_dafny.IntOfInt64(731), _318_v189)
				var _pat_let_tv1 = _319_v190
				_ = _pat_let_tv1
				var _pat_let_tv2 = _319_v190
				_ = _pat_let_tv2
				var _320_v191 _dafny.Array
				_ = _320_v191
				var _nwElement0_10 D11 = _317_v188
				_ = _nwElement0_10
				var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(18))
				_ = _nw52
				(_nw52).ArraySet1(_nwElement0_10, 0)
				(_nw52).ArraySet1(_317_v188, 1)
				(_nw52).ArraySet1(Companion_D11_.Create_DC25_(_319_v190), 2)
				(_nw52).ArraySet1(Companion_Default___.Fm39(globalState), 3)
				(_nw52).ArraySet1(_317_v188, 4)
				(_nw52).ArraySet1(_317_v188, 5)
				(_nw52).ArraySet1(_317_v188, 6)
				(_nw52).ArraySet1(Companion_D11_.Create_DC25_(_319_v190), 7)
				(_nw52).ArraySet1(_317_v188, 8)
				(_nw52).ArraySet1(_317_v188, 9)
				(_nw52).ArraySet1(func(_pat_let2_0 D11) D11 {
					return func(_321_dt__update__tmp_h1 D11) D11 {
						return func(_pat_let3_0 D11) D11 {
							return func(_322_dt__update_hcf45_h0 D11) D11 {
								return Companion_D11_.Create_DC25_(_322_dt__update_hcf45_h0)
							}(_pat_let3_0)
						}(_pat_let_tv1)
					}(_pat_let2_0)
				}(_317_v188), 10)
				(_nw52).ArraySet1(_317_v188, 11)
				(_nw52).ArraySet1(_317_v188, 12)
				(_nw52).ArraySet1(func(_pat_let4_0 D11) D11 {
					return func(_323_dt__update__tmp_h2 D11) D11 {
						return func(_pat_let5_0 D11) D11 {
							return func(_324_dt__update_hcf45_h1 D11) D11 {
								return Companion_D11_.Create_DC25_(_324_dt__update_hcf45_h1)
							}(_pat_let5_0)
						}(_pat_let_tv2)
					}(_pat_let4_0)
				}(_317_v188), 13)
				(_nw52).ArraySet1(_317_v188, 14)
				(_nw52).ArraySet1(_317_v188, 15)
				(_nw52).ArraySet1(_317_v188, 16)
				(_nw52).ArraySet1(Companion_Default___.Fm39(globalState), 17)
				_320_v191 = _nw52
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_320_v191), 0))
				_ = _index25
				(_320_v191).ArraySet1(Companion_Default___.Fm39(globalState), (_index25).Int())
				var _pat_let_tv3 = _319_v190
				_ = _pat_let_tv3
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_320_v191), 0))
				_ = _index26
				(_320_v191).ArraySet1(func(_pat_let6_0 D11) D11 {
					return func(_325_dt__update__tmp_h3 D11) D11 {
						return func(_pat_let7_0 D11) D11 {
							return func(_326_dt__update_hcf45_h2 D11) D11 {
								return Companion_D11_.Create_DC25_(_326_dt__update_hcf45_h2)
							}(_pat_let7_0)
						}(_pat_let_tv3)
					}(_pat_let6_0)
				}(Companion_D11_.Create_DC25_(_319_v190)), (_index26).Int())
				r0 = !(!(p0)) || (((_dafny.Zero).Minus(p2)).Cmp(p1) < 0)
				var _327_v192 _dafny.MultiSet
				_ = _327_v192
				_327_v192 = _dafny.MultiSetOf(!(false), p0)
				var _328_v193 bool
				_ = _328_v193
				_328_v193 = p0
				var _329_v194 _dafny.Array
				_ = _329_v194
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_7
				var _nw53 _dafny.Array
				_ = _nw53
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw53 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Sequence = (func(_330_v181 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_331_i15 _dafny.Int) _dafny.Sequence {
							return _330_v181
						}
					})(_309_v181)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw53 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw53).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw53).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_329_v194 = _nw53
				var _332_v195 *C1
				_ = _332_v195
				var _nw54 *C1 = New_C1_()
				_ = _nw54
				_nw54.Ctor__((p3).Cmp(p1) <= 0, _327_v192, Companion_Default___.Fm43(p0, _328_v193, p0, globalState), Companion_Default___.SafeModuloInt(p2, p3), _329_v194)
				_332_v195 = _nw54
				(globalState).F18 = (_dafny.Zero).Minus(((_327_v192).Cardinality()).Minus((_dafny.IntOfInt64(450)).Minus(p1)))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _333_v196 _dafny.Array
	_ = _333_v196
	var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
	_ = _nw55
	_333_v196 = _nw55
	for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_333_v196), 0))); ; {
		_guard_loop_0, _ok19 := _iter19()
		if !_ok19 {
			break
		}
		var _334_i16 _dafny.Int
		_334_i16 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_334_i16).Sign() != -1) && ((_334_i16).Cmp(_dafny.ArrayLen((_333_v196), 0)) < 0)) {
			(_333_v196).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)), (_334_i16).Int())
		}
	}
	var _335_i17 _dafny.Int
	_ = _335_i17
	_335_i17 = _dafny.Zero
	{
		for p0 {
			{
				if (_335_i17).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_335_i17 = (_335_i17).Plus(_dafny.One)
				(globalState).F24 = p2
				r0 = ((_dafny.IntOfUint32((_309_v181).Cardinality())).Plus(p1)).Cmp(((_dafny.Zero).Minus(p3)).Times(_dafny.IntOfUint32((_309_v181).Cardinality()))) <= 0
				var _336_v197 _dafny.Sequence
				_ = _336_v197
				_336_v197 = _dafny.SeqOf(p0, p0, p0, p0, p0)
				var _337_v198 _dafny.Map
				_ = _337_v198
				_337_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_336_v197).Cardinality()), p0)
				r0 = (p0) || (!((func() bool {
					if (_337_v198).Contains(p1) {
						return (_337_v198).Get(p1).(bool)
					}
					return p0
				})()))
				var _338_v200 D9
				_ = _338_v200
				_338_v200 = Companion_D9_.Create_DC20_(_dafny.IntOfInt64(89), p0, p0, p0, func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter20 := _dafny.Iterate((_337_v198).Keys().Elements()); ; {
						_compr_19, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _339_v199 _dafny.Int
						_339_v199 = interface{}(_compr_19).(_dafny.Int)
						if (_337_v198).Contains(_339_v199) {
							_coll19.Add((_339_v199).Times(_dafny.IntOfInt64(-886)), p3)
						}
					}
					return _coll19.ToMap()
				}())
				var _340_v201 _dafny.Map
				_ = _340_v201
				_340_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_338_v200).Dtor_cf34(), p3)
				var _341_v202 _dafny.Map
				_ = _341_v202
				_341_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_340_v201, _dafny.IntOfUint32((_309_v181).Cardinality()))
				var _342_v203 _dafny.MultiSet
				_ = _342_v203
				_342_v203 = _dafny.MultiSetOf(false)
				var _343_v204 _dafny.Sequence
				_ = _343_v204
				_343_v204 = _dafny.SeqOf(p1)
				var _344_v205 _dafny.Map
				_ = _344_v205
				_344_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
				var _345_v206 _dafny.Array
				_ = _345_v206
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_8
				var _nw56 _dafny.Array
				_ = _nw56
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw56 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Sequence = (func(_346_v181 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_347_i18 _dafny.Int) _dafny.Sequence {
							return _346_v181
						}
					})(_309_v181)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw56 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw56).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw56).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_345_v206 = _nw56
				var _348_v207 _dafny.Map
				_ = _348_v207
				_348_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_344_v205).Contains(p0) {
						return (_344_v205).Get(p0).(bool)
					}
					return p0
				})(), _345_v206)
				var _349_v208 *C5
				_ = _349_v208
				var _nw57 *C5 = New_C5_()
				_ = _nw57
				_nw57.Ctor__(_341_v202, _342_v203, _343_v204, p3, (func() _dafny.Array {
					if (_348_v207).Contains(p0) {
						return (_348_v207).Get(p0).(_dafny.Array)
					}
					return _345_v206
				})())
				_349_v208 = _nw57
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _350_v209 _dafny.Set
	_ = _350_v209
	_350_v209 = _dafny.SetOf(_310_v182)
	var _351_v210 _dafny.MultiSet
	_ = _351_v210
	_351_v210 = _dafny.MultiSetOf(_dafny.IntOfInt64(365))
	var _352_v211 _dafny.MultiSet
	_ = _352_v211
	_352_v211 = _dafny.MultiSetOf((_350_v209).Cardinality(), p1, p3, (_351_v210).Cardinality())
	_351_v210 = _351_v210
	var _353_v212 _dafny.Array
	_ = _353_v212
	var _nw58 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
	_ = _nw58
	_353_v212 = _nw58
	var _354_v213 _dafny.MultiSet
	_ = _354_v213
	_354_v213 = _dafny.MultiSetOf(_353_v212, _353_v212)
	var _hi0 _dafny.Int = (_354_v213).Cardinality()
	_ = _hi0
	for _355_i19 := _dafny.IntOfUint32((Companion_Default___.Fm14(p0, p0, p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("owyumxl")).Cardinality()), globalState)).Cardinality()); _355_i19.Cmp(_hi0) < 0; _355_i19 = _355_i19.Plus(_dafny.One) {
		var _356_v214 _dafny.Array
		_ = _356_v214
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_9
		var _nw59 _dafny.Array
		_ = _nw59
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw59 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) D13 = (func(_357_v181 _dafny.Sequence, _358_p0 bool, _359_i19 _dafny.Int) func(_dafny.Int) D13 {
				return func(_360_i20 _dafny.Int) D13 {
					return Companion_D13_.Create_DC32_(_dafny.IntOfUint32((_357_v181).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_358_p0)).Cardinality())), _dafny.SeqOf(_359_i19), _358_p0)
				}
			})(_309_v181, p0, _355_i19)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw59 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw59).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw59).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_356_v214 = _nw59
		var _361_v215 _dafny.Sequence
		_ = _361_v215
		_361_v215 = _dafny.SeqOf(p1, p3, _355_i19, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-99))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_362_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_363_i21 _dafny.Int) _dafny.CodePoint {
				return _362_v182
			}
		})(_310_v182)))).Cardinality()), p3)
		var _364_v216 bool
		_ = _364_v216
		_364_v216 = false
		var _365_v217 D13
		_ = _365_v217
		_365_v217 = Companion_D13_.Create_DC32_(p3, p3, _361_v215, (_364_v216))
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_356_v214), 0))
		_ = _index27
		(_356_v214).ArraySet1(_365_v217, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_356_v214), 0))
		_ = _index28
		(_356_v214).ArraySet1(_365_v217, (_index28).Int())
		(globalState).F26 = p0
		_309_v181 = _dafny.Companion_Sequence_.Concatenate(_309_v181, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(809))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_366_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_367_i22 _dafny.Int) _dafny.CodePoint {
				return _366_v182
			}
		})(_310_v182))))
		var _368_v218 _dafny.Array
		_ = _368_v218
		var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw60
		_368_v218 = _nw60
		var _369_v219 _dafny.Array
		_ = _369_v219
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw61
		_369_v219 = _nw61
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_368_v218), 0))
		_ = _index29
		(_368_v218).ArraySet1(_369_v219, (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_368_v218), 0))
		_ = _index30
		(_368_v218).ArraySet1(_369_v219, (_index30).Int())
	}
	r0 = p0
	var _370_v220 _dafny.Map
	_ = _370_v220
	_370_v220 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
	var _371_v221 _dafny.Map
	_ = _371_v221
	_371_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
	var _372_v222 _dafny.Map
	_ = _372_v222
	_372_v222 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_370_v220, (_371_v221).Cardinality())
	var _373_v223 _dafny.Sequence
	_ = _373_v223
	_373_v223 = _dafny.SeqOf(p1, p3)
	var _374_v224 _dafny.Array
	_ = _374_v224
	var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
	_ = _nw62
	_374_v224 = _nw62
	var _375_v225 *C5
	_ = _375_v225
	var _nw63 *C5 = New_C5_()
	_ = _nw63
	_nw63.Ctor__(_372_v222, _dafny.MultiSetOf(p0, p0), _373_v223, p2, _374_v224)
	_375_v225 = _nw63
	var _376_v226 D22
	_ = _376_v226
	_376_v226 = Companion_D22_.Create_DC54_(_310_v182, _375_v225, p0, _dafny.IntOfUint32((_309_v181).Cardinality()))
	r1 = (_376_v226).Dtor_cf98()
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _377_v0 _dafny.Array
	_ = _377_v0
	var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
	_ = _nw64
	_377_v0 = _nw64
	var _378_v1 _dafny.Sequence
	_ = _378_v1
	_378_v1 = _dafny.UnicodeSeqOfUtf8Bytes("yn")
	var _379_v2 _dafny.MultiSet
	_ = _379_v2
	_379_v2 = _dafny.MultiSetOf(_378_v1)
	var _380_v3 _dafny.CodePoint
	_ = _380_v3
	_380_v3 = _dafny.CodePoint('g')
	var _381_v4 bool
	_ = _381_v4
	_381_v4 = false
	var _382_v5 _dafny.Set
	_ = _382_v5
	_382_v5 = _dafny.SetOf(_381_v4)
	var _383_v6 _dafny.Int
	_ = _383_v6
	_383_v6 = _dafny.IntOfInt64(491)
	var _384_v7 _dafny.Set
	_ = _384_v7
	_384_v7 = _dafny.SetOf((_382_v5).Cardinality(), _383_v6, _383_v6)
	var _385_v8 _dafny.Map
	_ = _385_v8
	_385_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v3, (_384_v7).Cardinality())
	var _386_v9 _dafny.Sequence
	_ = _386_v9
	_386_v9 = _dafny.SeqOf(_383_v6)
	var _387_v10 _dafny.Array
	_ = _387_v10
	var _nwElement0_11 _dafny.Sequence = _dafny.SeqOf((_385_v8).Cardinality())
	_ = _nwElement0_11
	var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(3))
	_ = _nw65
	(_nw65).ArraySet1(_nwElement0_11, 0)
	(_nw65).ArraySet1(_386_v9, 1)
	(_nw65).ArraySet1(_386_v9, 2)
	_387_v10 = _nw65
	var _388_v11 _dafny.Array
	_ = _388_v11
	var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
	_ = _nw66
	_388_v11 = _nw66
	var _389_v12 _dafny.Map
	_ = _389_v12
	_389_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v1, _388_v11)
	var _390_v13 _dafny.Set
	_ = _390_v13
	_390_v13 = _dafny.SetOf(_378_v1)
	var _391_globalState *GlobalState
	_ = _391_globalState
	var _nw67 *GlobalState = New_GlobalState_()
	_ = _nw67
	_nw67.Ctor__(_dafny.IntOfInt64(367), _dafny.UnicodeSeqOfUtf8Bytes("thcvtep"), _377_v0, _dafny.IntOfInt64(880), _dafny.IntOfInt64(597), _dafny.IntOfInt64(980), false, true, (_379_v2).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("kseko"))), false, _387_v10, _dafny.IntOfInt64(683), _dafny.IntOfInt64(646), _389_v12, false, _dafny.IntOfInt64(368), false, (_390_v13).Intersection(_390_v13), _dafny.IntOfInt64(-783), _dafny.IntOfInt64(756), true, _dafny.IntOfInt64(853), func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(979), _dafny.IntOfInt64(540))); ; {
			_compr_20, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _392_v14 _dafny.Int
			_392_v14 = interface{}(_compr_20).(_dafny.Int)
			if ((_dafny.IntOfInt64(979)).Cmp(_392_v14) <= 0) && ((_392_v14).Cmp(_dafny.IntOfInt64(540)) < 0) {
				_coll20.Add((_392_v14).Minus((_384_v7).Cardinality()), _383_v6)
			}
		}
		return _coll20.ToMap()
	}(), _388_v11, _dafny.IntOfInt64(539), _dafny.IntOfInt64(-892), true)
	_391_globalState = _nw67
	var _rhs28 bool = _381_v4
	_ = _rhs28
	var _rhs29 _dafny.Int = _383_v6
	_ = _rhs29
	var _lhs16 *GlobalState = _391_globalState
	_ = _lhs16
	_381_v4 = _rhs28
	_lhs16.F0 = _rhs29
	if Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_383_v6, _383_v6), _391_globalState) {
		if !(_381_v4) || (_381_v4) {
			_381_v4 = _381_v4
			_378_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_378_v1, _378_v1), (Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_378_v1, _378_v1)).Cardinality()))).Uint32(), _380_v3), _dafny.Companion_Sequence_.Concatenate(_378_v1, _378_v1))
			var _393_v15 _dafny.Map
			_ = _393_v15
			_393_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, _383_v6)
			_393_v15 = (_393_v15).Update(_381_v4, _dafny.IntOfUint32((_378_v1).Cardinality()))
			(_391_globalState).F26 = _381_v4
			var _394_v16 _dafny.MultiSet
			_ = _394_v16
			_394_v16 = _dafny.MultiSetOf(_381_v4)
			var _395_v17 _dafny.Sequence
			_ = _395_v17
			_395_v17 = _dafny.SeqOf((_dafny.MultiSetOf(_381_v4, _381_v4)).Intersection(_394_v16), _394_v16, (_dafny.MultiSetOf(_381_v4, _381_v4, _381_v4)).Union(_394_v16))
			_395_v17 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-117))).Uint32(), func(coer29 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_396_v4 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_397_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_396_v4)
				}
			})(_381_v4)))
		} else {
			(_391_globalState).F26 = _381_v4
			(_391_globalState).F18 = Companion_Default___.Fm1(_391_globalState)
			var _398_v18 _dafny.Array
			_ = _398_v18
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
			_ = _nw68
			_398_v18 = _nw68
			var _rhs30 _dafny.Int = _383_v6
			_ = _rhs30
			var _rhs31 _dafny.Array = _398_v18
			_ = _rhs31
			var _lhs17 *GlobalState = _391_globalState
			_ = _lhs17
			_lhs17.F24 = _rhs30
			_398_v18 = _rhs31
			var _399_v19 bool
			_ = _399_v19
			var _400_v20 _dafny.Int
			_ = _400_v20
			var _out0 bool
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			_out0, _out1 = Companion_Default___.M0((_381_v4) == (_381_v4), _383_v6, _383_v6, _383_v6, _391_globalState)
			_399_v19 = _out0
			_400_v20 = _out1
			_399_v19 = false
		}
		var _401_v21 _dafny.Sequence
		_ = _401_v21
		_401_v21 = _dafny.SeqOf(_377_v0, _377_v0, _377_v0)
		_401_v21 = _dafny.Companion_Sequence_.Concatenate(_401_v21, _dafny.Companion_Sequence_.Concatenate(_401_v21, _401_v21))
		if !(_381_v4) {
			(_391_globalState).F0 = _383_v6
			var _402_v22 _dafny.Map
			_ = _402_v22
			_402_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _381_v4)
			_402_v22 = (_402_v22).Update(_383_v6, true)
			var _403_v23 _dafny.Array
			_ = _403_v23
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw69
			_403_v23 = _nw69
			var _404_v24 *C4
			_ = _404_v24
			var _nw70 *C4 = New_C4_()
			_ = _nw70
			_nw70.Ctor__(_383_v6, _403_v23)
			_404_v24 = _nw70
			var _405_v25 _dafny.MultiSet
			_ = _405_v25
			_405_v25 = _dafny.MultiSetOf(_383_v6)
			var _406_v26 *C0
			_ = _406_v26
			var _nw71 *C0 = New_C0_()
			_ = _nw71
			_nw71.Ctor__((_dafny.Zero).Minus((func() _dafny.Int {
				if (_405_v25).Contains(_dafny.IntOfUint32((_386_v9).Cardinality())) {
					return (_405_v25).Multiplicity(_dafny.IntOfUint32((_386_v9).Cardinality()))
				}
				return _383_v6
			})()), _378_v1)
			_406_v26 = _nw71
			_383_v6 = (_406_v26).F36()
		} else {
			_388_v11 = _388_v11
			var _407_v27 _dafny.Map
			_ = _407_v27
			_407_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(205), _381_v4)
			(_391_globalState).F26 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_386_v9, _386_v9), (Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_386_v9, _386_v9)).Cardinality()))).Uint32(), (_407_v27).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-955))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_408_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_409_i1 _dafny.Int) _dafny.Int {
					return _408_v6
				}
			})(_383_v6))))
			(_391_globalState).F15 = (_383_v6).Minus(_dafny.IntOfInt64(-306))
			var _410_v28 D16
			_ = _410_v28
			_410_v28 = Companion_D16_.Create_DC40_(true, (_dafny.Zero).Minus(_383_v6), _dafny.IntOfInt64(855))
			var _411_v29 _dafny.Map
			_ = _411_v29
			_411_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _410_v28)
			(_391_globalState).F21 = (_dafny.Zero).Minus((_383_v6).Minus((_411_v29).Cardinality()))
			var _412_v30 _dafny.Map
			_ = _412_v30
			_412_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _dafny.IntOfInt64(189))
			_412_v30 = _412_v30
		}
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_388_v11), 0))
		_ = _index31
		(_388_v11).ArraySet1(_381_v4, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_377_v0), 0))
		_ = _index32
		(_377_v0).ArraySet1(_383_v6, (_index32).Int())
		var _413_v31 _dafny.Sequence
		_ = _413_v31
		_413_v31 = _dafny.SeqOf(_386_v9, _386_v9)
		var _414_v32 D9
		_ = _414_v32
		_414_v32 = Companion_D9_.Create_DC17_(_384_v7)
		var _415_v33 _dafny.Sequence
		_ = _415_v33
		_415_v33 = _dafny.SeqOf(_414_v32)
		var _416_v34 _dafny.Map
		_ = _416_v34
		_416_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v33, _377_v0)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_388_v11), 0))
		_ = _index33
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_377_v0), 0))
		_ = _index34
		var _rhs32 bool = !_dafny.Companion_Sequence_.Equal(_413_v31, _dafny.SeqOf(_386_v9))
		_ = _rhs32
		var _rhs33 bool = !(_381_v4)
		_ = _rhs33
		var _rhs34 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_383_v6, _383_v6))
		_ = _rhs34
		var _rhs35 bool = !(_381_v4) || ((_381_v4) || (false))
		_ = _rhs35
		var _rhs36 bool = !((_416_v34).Merge(_416_v34)).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-536))).Uint32(), func(coer31 func(_dafny.Int) D9) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_417_v6 _dafny.Int) func(_dafny.Int) D9 {
			return func(_418_i2 _dafny.Int) D9 {
				return Companion_D9_.Create_DC17_(_dafny.SetOf(_417_v6))
			}
		})(_383_v6))))
		_ = _rhs36
		var _lhs18 *GlobalState = _391_globalState
		_ = _lhs18
		var _lhs19 _dafny.Array = _388_v11
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_388_v11), 0))
		_ = _lhs20
		var _lhs21 _dafny.Array = _377_v0
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_377_v0), 0))
		_ = _lhs22
		_lhs18.F26 = _rhs32
		(_lhs19).ArraySet1(_rhs33, (_lhs20).Int())
		(_lhs21).ArraySet1(_rhs34, (_lhs22).Int())
		_381_v4 = _rhs35
		_381_v4 = _rhs36
		var _419_v35 _dafny.MultiSet
		_ = _419_v35
		_419_v35 = _dafny.MultiSetOf((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), _381_v4, ((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)) == (_381_v4))
		_419_v35 = _419_v35
	} else {
		var _420_v36 _dafny.Sequence
		_ = _420_v36
		_420_v36 = _dafny.SeqOf((_383_v6).Cmp(_383_v6) > 0)
		_420_v36 = _dafny.Companion_Sequence_.Concatenate(_420_v36, _420_v36)
		var _421_v37 bool
		_ = _421_v37
		var _422_v38 _dafny.Int
		_ = _422_v38
		var _out2 bool
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out2, _out3 = Companion_Default___.M0(_381_v4, _dafny.IntOfUint32((_386_v9).Cardinality()), _383_v6, _383_v6, _391_globalState)
		_421_v37 = _out2
		_422_v38 = _out3
		if _381_v4 {
			var _423_v39 _dafny.Map
			_ = _423_v39
			_423_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _381_v4)
			var _424_v40 bool
			_ = _424_v40
			var _425_v41 _dafny.Int
			_ = _425_v41
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out4, _out5 = Companion_Default___.M0(_421_v37, (_422_v38).Minus(_383_v6), _422_v38, (_dafny.MultiSetOf(_421_v37, _421_v37, (func() bool {
				if (_423_v39).Contains(_422_v38) {
					return (_423_v39).Get(_422_v38).(bool)
				}
				return _381_v4
			})())).Cardinality(), _391_globalState)
			_424_v40 = _out4
			_425_v41 = _out5
			_380_v3 = Companion_Default___.Fm12(Companion_Default___.SafeModuloInt(_383_v6, _383_v6), _424_v40, _391_globalState)
			var _426_v42 _dafny.Map
			_ = _426_v42
			_426_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v3, _424_v40)
			var _427_v43 D2
			_ = _427_v43
			_427_v43 = Companion_D2_.Create_DC3_(_426_v42)
			var _428_v44 _dafny.Map
			_ = _428_v44
			_428_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_427_v43, _383_v6)
			_428_v44 = (_428_v44).Update(_427_v43, (_425_v41).Minus(_383_v6))
			(_391_globalState).F26 = _381_v4
			var _429_v45 _dafny.MultiSet
			_ = _429_v45
			_429_v45 = _dafny.MultiSetOf(false, _381_v4, !(_381_v4))
			var _430_v46 _dafny.Array
			_ = _430_v46
			var _nwElement0_12 _dafny.Sequence = _378_v1
			_ = _nwElement0_12
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(10))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_12, 0)
			(_nw72).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("w"), 1)
			(_nw72).ArraySet1(_378_v1, 2)
			(_nw72).ArraySet1(_378_v1, 3)
			(_nw72).ArraySet1(_378_v1, 4)
			(_nw72).ArraySet1(Companion_Default___.Fm8(_391_globalState), 5)
			(_nw72).ArraySet1(_378_v1, 6)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_422_v38, _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3), 7)
			(_nw72).ArraySet1(_378_v1, 8)
			(_nw72).ArraySet1(_378_v1, 9)
			_430_v46 = _nw72
			var _431_v47 *C1
			_ = _431_v47
			var _nw73 *C1 = New_C1_()
			_ = _nw73
			_nw73.Ctor__(_424_v40, _429_v45, _386_v9, _383_v6, _430_v46)
			_431_v47 = _nw73
			_431_v47 = _431_v47
		} else {
			_381_v4 = (func() bool {
				if _421_v37 {
					return _381_v4
				}
				return _381_v4
			})()
			var _432_v48 D13
			_ = _432_v48
			_432_v48 = Companion_D13_.Create_DC32_((_dafny.Zero).Minus(_383_v6), _422_v38, _386_v9, _421_v37)
			var _433_v49 bool
			_ = _433_v49
			var _434_v50 _dafny.Int
			_ = _434_v50
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out6, _out7 = Companion_Default___.M0(true, (_dafny.IntOfUint32(((_432_v48).Dtor_cf59()).Cardinality())).Times(_422_v38), ((_dafny.Zero).Minus(_dafny.IntOfInt64(-481))).Times(_422_v38), Companion_Default___.SafeModuloInt(Companion_Default___.Fm1(_391_globalState), _383_v6), _391_globalState)
			_433_v49 = _out6
			_434_v50 = _out7
			(_391_globalState).F26 = true
			var _435_v51 _dafny.Map
			_ = _435_v51
			_435_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-243), _422_v38)
			var _436_v52 _dafny.Map
			_ = _436_v52
			_436_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_435_v51, _434_v50)
			var _437_v53 _dafny.MultiSet
			_ = _437_v53
			_437_v53 = _dafny.MultiSetOf(_433_v49)
			var _438_v54 _dafny.Array
			_ = _438_v54
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_10
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Sequence = (func(_439_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_440_i3 _dafny.Int) _dafny.Sequence {
						return _439_v1
					}
				})(_378_v1)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw74 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw74).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw74).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_438_v54 = _nw74
			var _441_v55 *C5
			_ = _441_v55
			var _nw75 *C5 = New_C5_()
			_ = _nw75
			_nw75.Ctor__((_436_v52).Merge(Companion_Default___.Fm54(_381_v4, _433_v49, _391_globalState)), _437_v53, _386_v9, _434_v50, _438_v54)
			_441_v55 = _nw75
			var _442_v56 _dafny.Map
			_ = _442_v56
			_442_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v50, _dafny.UnicodeSeqOfUtf8Bytes("ijuo"))
			var _443_v57 *C0
			_ = _443_v57
			var _nw76 *C0 = New_C0_()
			_ = _nw76
			_nw76.Ctor__(_dafny.IntOfInt64(448), (func() _dafny.Sequence {
				if (_442_v56).Contains(_dafny.IntOfInt64(-214)) {
					return (_442_v56).Get(_dafny.IntOfInt64(-214)).(_dafny.Sequence)
				}
				return _378_v1
			})())
			_443_v57 = _nw76
			_443_v57 = _443_v57
		}
		var _444_v58 _dafny.MultiSet
		_ = _444_v58
		_444_v58 = _dafny.MultiSetOf(_420_v36, _420_v36, _420_v36, _420_v36, _420_v36)
		var _445_v60 _dafny.Set
		_ = _445_v60
		_445_v60 = (_dafny.SetOf(_420_v36, _dafny.SeqOf(_421_v37, _381_v4))).Difference(func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter22 := _dafny.Iterate(((_444_v58).Update(_dafny.SeqOf(true), Companion_Default___.Abs(_422_v38))).Elements()); ; {
				_compr_21, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _446_v59 _dafny.Sequence
				_446_v59 = interface{}(_compr_21).(_dafny.Sequence)
				if ((_444_v58).Update(_dafny.SeqOf(true), Companion_Default___.Abs(_422_v38))).Contains(_446_v59) {
					_coll21.Add(_446_v59)
				}
			}
			return _coll21.ToSet()
		}())
		var _source10 _dafny.Set = _445_v60
		_ = _source10
		var _447___mcc_h0 _dafny.Set = _source10
		_ = _447___mcc_h0
		var _448_cf62 _dafny.Set = _447___mcc_h0
		_ = _448_cf62
		(_391_globalState).F26 = !(!(!(_381_v4))) || (_381_v4)
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_11
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = (func(_449_v38 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_450_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_450_i4, _449_v38)
				}
			})(_422_v38)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw77 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw77).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw77).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_377_v0 = _nw77
		_378_v1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(675))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_451_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_452_i5 _dafny.Int) _dafny.CodePoint {
				return _451_v3
			}
		})(_380_v3)))
		_380_v3 = _380_v3
		(_391_globalState).F18 = _383_v6
	}
	var _453_v61 D11
	_ = _453_v61
	_453_v61 = Companion_D11_.Create_DC25_(Companion_D11_.Create_DC23_(_dafny.MultiSetOf(_384_v7)))
	var _454_v62 _dafny.MultiSet
	_ = _454_v62
	_454_v62 = _dafny.MultiSetOf(_384_v7)
	var _455_v63 D11
	_ = _455_v63
	_455_v63 = Companion_D11_.Create_DC23_(_454_v62)
	var _pat_let_tv4 = _455_v63
	_ = _pat_let_tv4
	var _source11 D11 = func(_pat_let8_0 D11) D11 {
		return func(_456_dt__update__tmp_h0 D11) D11 {
			return func(_pat_let9_0 D11) D11 {
				return func(_457_dt__update_hcf45_h0 D11) D11 {
					return Companion_D11_.Create_DC25_(_457_dt__update_hcf45_h0)
				}(_pat_let9_0)
			}(_pat_let_tv4)
		}(_pat_let8_0)
	}(_453_v61)
	_ = _source11
	if _source11.Is_DC24() {
		var _458___mcc_h1 _dafny.Int = _source11.Get_().(D11_DC24).Cf43
		_ = _458___mcc_h1
		var _459___mcc_h2 D4 = _source11.Get_().(D11_DC24).Cf44
		_ = _459___mcc_h2
		var _460_cf44 D4 = _459___mcc_h2
		_ = _460_cf44
		var _461_cf43 _dafny.Int = _458___mcc_h1
		_ = _461_cf43
		var _462_v64 _dafny.Map
		_ = _462_v64
		_462_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_378_v1).Cardinality()), _461_cf43)
		var _463_v65 _dafny.MultiSet
		_ = _463_v65
		_463_v65 = _dafny.MultiSetOf(_381_v4)
		var _464_v66 T2
		_ = _464_v66
		var _nw78 *C2 = New_C2_()
		_ = _nw78
		_nw78.Ctor__(_463_v65, _383_v6)
		_464_v66 = _nw78
		var _465_v67 _dafny.Set
		_ = _465_v67
		_465_v67 = _dafny.SetOf(_464_v66, _464_v66, _464_v66)
		_462_v64 = (_462_v64).Update((_386_v9).Select((Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int), (_465_v67).Cardinality())
		var _466_v68 _dafny.Array
		_ = _466_v68
		var _467_v69 _dafny.Int
		_ = _467_v69
		var _out8 _dafny.Array
		_ = _out8
		var _out9 _dafny.Int
		_ = _out9
		_out8, _out9 = (_464_v66).M2(_381_v4, _381_v4, (_463_v65).Cardinality(), (_dafny.Zero).Minus(_461_cf43), _391_globalState)
		_466_v68 = _out8
		_467_v69 = _out9
		var _468_v70 _dafny.Array
		_ = _468_v70
		var _nwElement0_13 _dafny.Sequence = _378_v1
		_ = _nwElement0_13
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(22))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_13, 0)
		(_nw79).ArraySet1(_378_v1, 1)
		(_nw79).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-447))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}((func(_469_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_470_i6 _dafny.Int) _dafny.CodePoint {
				return _469_v3
			}
		})(_380_v3))), 2)
		(_nw79).ArraySet1(_378_v1, 3)
		(_nw79).ArraySet1(_378_v1, 4)
		(_nw79).ArraySet1(_378_v1, 5)
		(_nw79).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ukgaklleg"), 6)
		(_nw79).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hl"), 7)
		(_nw79).ArraySet1(_378_v1, 8)
		(_nw79).ArraySet1(_378_v1, 9)
		(_nw79).ArraySet1(_378_v1, 10)
		(_nw79).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wkbli"), 11)
		(_nw79).ArraySet1(_378_v1, 12)
		(_nw79).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("usfy"), 13)
		(_nw79).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yakla"), 14)
		(_nw79).ArraySet1(_378_v1, 15)
		(_nw79).ArraySet1(_378_v1, 16)
		(_nw79).ArraySet1(_378_v1, 17)
		(_nw79).ArraySet1(_378_v1, 18)
		(_nw79).ArraySet1(_378_v1, 19)
		(_nw79).ArraySet1(_378_v1, 20)
		(_nw79).ArraySet1(_378_v1, 21)
		_468_v70 = _nw79
		var _471_v71 _dafny.Map
		_ = _471_v71
		_471_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_391_globalState), _468_v70)
		var _472_v72 *C4
		_ = _472_v72
		var _nw80 *C4 = New_C4_()
		_ = _nw80
		_nw80.Ctor__(Companion_Default___.Fm1(_391_globalState), (func() _dafny.Array {
			if (_471_v71).Contains(_467_v69) {
				return (_471_v71).Get(_467_v69).(_dafny.Array)
			}
			return _468_v70
		})())
		_472_v72 = _nw80
		var _rhs37 bool = _381_v4
		_ = _rhs37
		var _rhs38 bool = _381_v4
		_ = _rhs38
		var _rhs39 _dafny.Int = _467_v69
		_ = _rhs39
		var _rhs40 bool = _381_v4
		_ = _rhs40
		var _lhs23 *GlobalState = _391_globalState
		_ = _lhs23
		var _lhs24 *GlobalState = _391_globalState
		_ = _lhs24
		var _lhs25 *GlobalState = _391_globalState
		_ = _lhs25
		_381_v4 = _rhs37
		_lhs23.F26 = _rhs38
		_lhs24.F15 = _rhs39
		_lhs25.F26 = _rhs40
	} else if _source11.Is_DC23() {
		var _473___mcc_h3 _dafny.MultiSet = _source11.Get_().(D11_DC23).Cf42
		_ = _473___mcc_h3
		var _474_cf42 _dafny.MultiSet = _473___mcc_h3
		_ = _474_cf42
		var _475_v73 bool
		_ = _475_v73
		var _476_v74 _dafny.Int
		_ = _476_v74
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out10, _out11 = Companion_Default___.M0((_383_v6).Cmp(_383_v6) >= 0, (_dafny.IntOfInt64(84)).Minus(_383_v6), _383_v6, (func() _dafny.Int {
			if _381_v4 {
				return _383_v6
			}
			return _dafny.IntOfUint32((_378_v1).Cardinality())
		})(), _391_globalState)
		_475_v73 = _out10
		_476_v74 = _out11
		if _381_v4 {
			var _477_v75 bool
			_ = _477_v75
			var _478_v76 _dafny.Int
			_ = _478_v76
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			_out12, _out13 = Companion_Default___.M0(_381_v4, _383_v6, _476_v74, _383_v6, _391_globalState)
			_477_v75 = _out12
			_478_v76 = _out13
			var _479_v77 _dafny.MultiSet
			_ = _479_v77
			_479_v77 = _dafny.MultiSetOf(_381_v4, _477_v75)
			var _480_v78 D13
			_ = _480_v78
			_480_v78 = Companion_D13_.Create_DC32_(_383_v6, (_dafny.Zero).Minus((_386_v9).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_476_v74), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int)), _386_v9, _381_v4)
			var _481_v79 _dafny.Map
			_ = _481_v79
			_481_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_378_v1).Cardinality()), _477_v75)
			var _482_v80 _dafny.Map
			_ = _482_v80
			_482_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_477_v75, true)
			var _483_v81 _dafny.Map
			_ = _483_v81
			_483_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_481_v79).Contains((_482_v80).Cardinality()) {
					return (_481_v79).Get((_482_v80).Cardinality()).(bool)
				}
				return _381_v4
			})(), _478_v76)
			var _484_v82 _dafny.Array
			_ = _484_v82
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw81
			_484_v82 = _nw81
			var _485_v83 D17
			_ = _485_v83
			_485_v83 = Companion_D17_.Create_DC43_(_484_v82)
			var _486_v84 T2
			_ = _486_v84
			var _nw82 *C6 = New_C6_()
			_ = _nw82
			_nw82.Ctor__(_475_v73, Companion_Default___.Fm1(_391_globalState), _479_v77, (_480_v78).Dtor_cf59(), (_476_v74).Minus((_483_v81).Cardinality()), (_485_v83).Dtor_cf82(), _479_v77, _478_v76)
			_486_v84 = _nw82
			_486_v84 = (func() T2 {
				if _477_v75 {
					return _486_v84
				}
				return _486_v84
			})()
			var _487_v85 _dafny.Array
			_ = _487_v85
			var _nwElement0_14 _dafny.Map = _481_v79
			_ = _nwElement0_14
			var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.One)
			_ = _nw83
			(_nw83).ArraySet1(_nwElement0_14, 0)
			_487_v85 = _nw83
			var _488_v86 _dafny.Array
			_ = _488_v86
			_488_v86 = _487_v85
			var _489_v87 _dafny.Map
			_ = _489_v87
			_489_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_383_v6), (_487_v85))
			_489_v87 = (_489_v87).Update(_478_v76, _487_v85)
			var _490_v88 *C1
			_ = _490_v88
			var _nw84 *C1 = New_C1_()
			_ = _nw84
			_nw84.Ctor__(true, (_479_v77).Update(_475_v73, Companion_Default___.Abs(_dafny.IntOfUint32((_378_v1).Cardinality()))), _386_v9, _dafny.IntOfInt64(805), _484_v82)
			_490_v88 = _nw84
			var _491_v89 _dafny.MultiSet
			_ = _491_v89
			_491_v89 = _dafny.MultiSetOf(_490_v88)
			var _492_v90 *C6
			_ = _492_v90
			var _nw85 *C6 = New_C6_()
			_ = _nw85
			_nw85.Ctor__(((_486_v84).F32()).Cmp((_491_v89).Cardinality()) == 0, Companion_Default___.SafeModuloInt((_486_v84).F32(), _dafny.IntOfUint32((_378_v1).Cardinality())), (_486_v84).F31(), _386_v9, Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_383_v6), _dafny.IntOfInt64(-584)), _484_v82, (_486_v84).F31(), _383_v6)
			_492_v90 = _nw85
			var _493_v91 _dafny.Map
			_ = _493_v91
			_493_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_494_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_495_i7 _dafny.Int) _dafny.CodePoint {
					return _494_v3
				}
			})(_380_v3))), (func() *C6 {
				if (_492_v90).F33() {
					return _492_v90
				}
				return _492_v90
			})())
			_492_v90 = (func() *C6 {
				if (_493_v91).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bcvfglbs"), _378_v1)) {
					return (_493_v91).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bcvfglbs"), _378_v1)).(*C6)
				}
				return _492_v90
			})()
			var _496_v92 _dafny.Sequence
			_ = _496_v92
			_496_v92 = _378_v1
			var _497_v93 bool
			_ = _497_v93
			var _498_v94 _dafny.Array
			_ = _498_v94
			var _499_v95 _dafny.Map
			_ = _499_v95
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Array
			_ = _out15
			var _out16 _dafny.Map
			_ = _out16
			_out14, _out15, _out16 = (_490_v88).M5((_490_v88).Fm11(!((_492_v90).F33()), _dafny.IntOfUint32((_378_v1).Cardinality()), _496_v92, _dafny.IntOfUint32((_378_v1).Cardinality()), _391_globalState), Companion_Default___.Fm0(_383_v6, _391_globalState), _391_globalState)
			_497_v93 = _out14
			_498_v94 = _out15
			_499_v95 = _out16
		} else {
			(_391_globalState).F26 = (_384_v7).IsDisjointFrom(_384_v7)
			var _500_v96 D10
			_ = _500_v96
			_500_v96 = Companion_D10_.Create_DC21_(_386_v9)
			var _501_v97 _dafny.Map
			_ = _501_v97
			_501_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _381_v4 {
					return _381_v4
				}
				return _381_v4
			})(), _500_v96)
			_501_v97 = (_501_v97).Update(true, _500_v96)
			var _502_v98 _dafny.Sequence
			_ = _502_v98
			_502_v98 = _378_v1
			var _503_v99 _dafny.Sequence
			_ = _503_v99
			_503_v99 = _dafny.SeqOf(Companion_Default___.Fm14(_381_v4, false, _475_v73, _383_v6, _391_globalState))
			var _504_v100 _dafny.Map
			_ = _504_v100
			_504_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_476_v74, true)
			var _505_v101 _dafny.Array
			_ = _505_v101
			var _nwElement0_15 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fgvfq")
			_ = _nwElement0_15
			var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(27))
			_ = _nw86
			(_nw86).ArraySet1(_nwElement0_15, 0)
			(_nw86).ArraySet1(_378_v1, 1)
			(_nw86).ArraySet1(_378_v1, 2)
			(_nw86).ArraySet1(_378_v1, 3)
			(_nw86).ArraySet1(_378_v1, 4)
			(_nw86).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_506_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_507_i8 _dafny.Int) _dafny.CodePoint {
					return _506_v3
				}
			})(_380_v3))), (Companion_Default___.SafeIndex(_476_v74, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_508_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_509_i8 _dafny.Int) _dafny.CodePoint {
					return _508_v3
				}
			})(_380_v3)))).Cardinality()))).Uint32(), _380_v3), 5)
			(_nw86).ArraySet1(_dafny.Companion_Sequence_.Update((_502_v98), (Companion_Default___.SafeIndex(Companion_Default___.Fm1(_391_globalState), _dafny.IntOfUint32((_502_v98).Cardinality()))).Uint32(), _380_v3), 6)
			(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("a"), 7)
			(_nw86).ArraySet1(_378_v1, 8)
			(_nw86).ArraySet1(_378_v1, 9)
			(_nw86).ArraySet1(Companion_Default___.Fm31(_475_v73, _381_v4, _391_globalState), 10)
			(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_510_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_511_i9 _dafny.Int) _dafny.CodePoint {
					return _510_v3
				}
			})(_380_v3))), 11)
			(_nw86).ArraySet1((_503_v99).Select((Companion_Default___.SafeIndex((_504_v100).Cardinality(), _dafny.IntOfUint32((_503_v99).Cardinality()))).Uint32()).(_dafny.Sequence), 12)
			(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qgwx"), 13)
			(_nw86).ArraySet1(_378_v1, 14)
			(_nw86).ArraySet1(_378_v1, 15)
			(_nw86).ArraySet1(_378_v1, 16)
			(_nw86).ArraySet1(_378_v1, 17)
			(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kx"), 18)
			(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_512_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_513_i10 _dafny.Int) _dafny.CodePoint {
					return _512_v3
				}
			})(_380_v3))), 19)
			(_nw86).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("olpvqovqh"), (Companion_Default___.SafeIndex(_476_v74, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("olpvqovqh")).Cardinality()))).Uint32(), _380_v3), 20)
			(_nw86).ArraySet1(_378_v1, 21)
			(_nw86).ArraySet1(_378_v1, 22)
			(_nw86).ArraySet1(_378_v1, 23)
			(_nw86).ArraySet1(_378_v1, 24)
			(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}((func(_514_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_515_i11 _dafny.Int) _dafny.CodePoint {
					return _514_v3
				}
			})(_380_v3))), 25)
			(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sfnpbmx"), 26)
			_505_v101 = _nw86
			var _516_v102 *C3
			_ = _516_v102
			var _nw87 *C3 = New_C3_()
			_ = _nw87
			_nw87.Ctor__(_476_v74, _505_v101)
			_516_v102 = _nw87
			var _517_v103 D13
			_ = _517_v103
			_517_v103 = Companion_D13_.Create_DC32_(_383_v6, _476_v74, _386_v9, _475_v73)
			var _518_v104 _dafny.Map
			_ = _518_v104
			_518_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_517_v103).Dtor_cf59(), _516_v102)
			_516_v102 = (func() *C3 {
				if (_518_v104).Contains(_386_v9) {
					return (_518_v104).Get(_386_v9).(*C3)
				}
				return _516_v102
			})()
			var _519_v105 _dafny.MultiSet
			_ = _519_v105
			_519_v105 = _dafny.MultiSetOf(_dafny.IntOfInt64(201), _476_v74)
			var _520_v106 T0
			_ = _520_v106
			var _nw88 *C3 = New_C3_()
			_ = _nw88
			_nw88.Ctor__(_383_v6, _505_v101)
			_520_v106 = _nw88
			var _521_v107 _dafny.Sequence
			_ = _521_v107
			_521_v107 = _dafny.SeqOf(_520_v106, _520_v106)
			var _522_v108 D9
			_ = _522_v108
			_522_v108 = Companion_D9_.Create_DC18_(true)
			var _523_v109 _dafny.Sequence
			_ = _523_v109
			_523_v109 = _dafny.SeqOf((_522_v108).Dtor_cf32())
			var _524_v110 D16
			_ = _524_v110
			_524_v110 = Companion_D16_.Create_DC40_(_475_v73, (_dafny.Zero).Minus((func() _dafny.Int {
				if (_519_v105).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_521_v107).Cardinality()), _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3)).Cardinality())) {
					return (_519_v105).Multiplicity(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_521_v107).Cardinality()), _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3)).Cardinality()))
				}
				return _383_v6
			})()), _dafny.IntOfUint32((_523_v109).Cardinality()))
			var _525_v111 _dafny.MultiSet
			_ = _525_v111
			_525_v111 = _dafny.MultiSetOf(_381_v4, (_524_v110).Dtor_cf74())
			var _526_v112 *C6
			_ = _526_v112
			var _nw89 *C6 = New_C6_()
			_ = _nw89
			_nw89.Ctor__(_475_v73, _dafny.IntOfInt64(549), _525_v111, _386_v9, _476_v74, _505_v101, _525_v111, (_520_v106).F27())
			_526_v112 = _nw89
			var _527_v113 _dafny.Map
			_ = _527_v113
			_527_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v3, _526_v112)
			_527_v113 = (_527_v113).Update(_380_v3, _526_v112)
			_385_v8 = (_385_v8).Update(_380_v3, _526_v112.F34)
		}
		var _528_v114 bool
		_ = _528_v114
		var _529_v115 _dafny.Int
		_ = _529_v115
		var _out17 bool
		_ = _out17
		var _out18 _dafny.Int
		_ = _out18
		_out17, _out18 = Companion_Default___.M0(_475_v73, _476_v74, _476_v74, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(391))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_530_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_531_i12 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_530_v1).Cardinality())
			}
		})(_378_v1)))).Cardinality()), _391_globalState)
		_528_v114 = _out17
		_529_v115 = _out18
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_388_v11), 0))
		_ = _index35
		(_388_v11).ArraySet1(_475_v73, (_index35).Int())
		var _532_v116 _dafny.MultiSet
		_ = _532_v116
		_532_v116 = _dafny.MultiSetOf(_529_v115, _529_v115)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_388_v11), 0))
		_ = _index36
		var _rhs41 bool = _528_v114
		_ = _rhs41
		var _rhs42 bool = _475_v73
		_ = _rhs42
		var _rhs43 bool = !(!(!(_528_v114) || (_528_v114)) || ((func() bool {
			if Companion_Default___.Fm0((_532_v116).Cardinality(), _391_globalState) {
				return _528_v114
			}
			return !(_528_v114)
		})()))
		_ = _rhs43
		var _lhs26 _dafny.Array = _388_v11
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_388_v11), 0))
		_ = _lhs27
		var _lhs28 *GlobalState = _391_globalState
		_ = _lhs28
		(_lhs26).ArraySet1(_rhs41, (_lhs27).Int())
		_lhs28.F26 = _rhs42
		_528_v114 = _rhs43
	} else {
		var _533___mcc_h4 D11 = _source11.Get_().(D11_DC25).Cf45
		_ = _533___mcc_h4
		var _534_cf45 D11 = _533___mcc_h4
		_ = _534_cf45
		if _381_v4 {
			var _535_v117 _dafny.Map
			_ = _535_v117
			_535_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _dafny.UnicodeSeqOfUtf8Bytes("jrjwq"))
			var _536_v118 _dafny.Map
			_ = _536_v118
			_536_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, _381_v4)
			var _537_v119 _dafny.Array
			_ = _537_v119
			var _nwElement0_16 _dafny.Sequence = _378_v1
			_ = _nwElement0_16
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(27))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_16, 0)
			(_nw90).ArraySet1(_378_v1, 1)
			(_nw90).ArraySet1(_378_v1, 2)
			(_nw90).ArraySet1(_378_v1, 3)
			(_nw90).ArraySet1(_378_v1, 4)
			(_nw90).ArraySet1(_378_v1, 5)
			(_nw90).ArraySet1(_378_v1, 6)
			(_nw90).ArraySet1(_378_v1, 7)
			(_nw90).ArraySet1(_378_v1, 8)
			(_nw90).ArraySet1(_378_v1, 9)
			(_nw90).ArraySet1(_378_v1, 10)
			(_nw90).ArraySet1(_378_v1, 11)
			(_nw90).ArraySet1(_378_v1, 12)
			(_nw90).ArraySet1(_378_v1, 13)
			(_nw90).ArraySet1(_378_v1, 14)
			(_nw90).ArraySet1(_378_v1, 15)
			(_nw90).ArraySet1(_378_v1, 16)
			(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(43))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_538_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_539_i13 _dafny.Int) _dafny.CodePoint {
					return _538_v3
				}
			})(_380_v3))), 17)
			(_nw90).ArraySet1(Companion_Default___.Fm37(_383_v6, _383_v6, _391_globalState), 18)
			(_nw90).ArraySet1(_378_v1, 19)
			(_nw90).ArraySet1(_378_v1, 20)
			(_nw90).ArraySet1(_378_v1, 21)
			(_nw90).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("by"), 22)
			(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_540_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_541_i14 _dafny.Int) _dafny.CodePoint {
					return _540_v3
				}
			})(_380_v3))), 23)
			(_nw90).ArraySet1(_378_v1, 24)
			(_nw90).ArraySet1((func() _dafny.Sequence {
				if (_535_v117).Contains((_536_v118).Cardinality()) {
					return (_535_v117).Get((_536_v118).Cardinality()).(_dafny.Sequence)
				}
				return _378_v1
			})(), 25)
			(_nw90).ArraySet1(_378_v1, 26)
			_537_v119 = _nw90
			var _542_v120 *C4
			_ = _542_v120
			var _nw91 *C4 = New_C4_()
			_ = _nw91
			_nw91.Ctor__(_dafny.IntOfUint32((_386_v9).Cardinality()), _537_v119)
			_542_v120 = _nw91
			var _543_v121 _dafny.MultiSet
			_ = _543_v121
			_543_v121 = _dafny.MultiSetOf(_542_v120)
			var _544_v122 _dafny.Map
			_ = _544_v122
			_544_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _386_v9)
			var _545_v123 _dafny.Sequence
			_ = _545_v123
			_545_v123 = _dafny.SeqOf(_381_v4)
			(_391_globalState).F21 = (((_543_v121).Cardinality()).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_544_v122).Contains(_dafny.IntOfUint32((_545_v123).Cardinality())) {
					return (_544_v122).Get(_dafny.IntOfUint32((_545_v123).Cardinality())).(_dafny.Sequence)
				}
				return _386_v9
			})()).Cardinality()))).Times((_383_v6).Minus(_383_v6))
			var _546_v124 _dafny.MultiSet
			_ = _546_v124
			_546_v124 = _dafny.MultiSetOf(_388_v11)
			var _rhs44 _dafny.MultiSet = (_546_v124).Union(_546_v124)
			_ = _rhs44
			var _rhs45 bool = ((_dafny.IntOfUint32((_378_v1).Cardinality())).Plus((_dafny.Zero).Minus(_383_v6))).Cmp(_383_v6) >= 0
			_ = _rhs45
			var _lhs29 *GlobalState = _391_globalState
			_ = _lhs29
			_546_v124 = _rhs44
			_lhs29.F26 = _rhs45
			_378_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_378_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_547_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_548_i15 _dafny.Int) _dafny.CodePoint {
					return _547_v3
				}
			})(_380_v3))), (Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_549_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_550_i15 _dafny.Int) _dafny.CodePoint {
					return _549_v3
				}
			})(_380_v3)))).Cardinality()))).Uint32(), _380_v3)), _378_v1)
			(_391_globalState).F26 = (_381_v4) == (_381_v4)
			var _551_v125 *C0
			_ = _551_v125
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__(_dafny.IntOfInt64(745), _378_v1)
			_551_v125 = _nw92
			var _552_v126 _dafny.Int
			_ = _552_v126
			var _out19 _dafny.Int
			_ = _out19
			_out19 = (_542_v120).M11(_551_v125, (_381_v4) && (_381_v4), _391_globalState)
			_552_v126 = _out19
		} else {
			(_391_globalState).F26 = _381_v4
			(_391_globalState).F26 = _381_v4
			_381_v4 = (_379_v2).IsDisjointFrom((_379_v2).Union(_379_v2))
			var _553_v127 _dafny.Array
			_ = _553_v127
			var _nwElement0_17 _dafny.Sequence = _378_v1
			_ = _nwElement0_17
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(24))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_17, 0)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(232))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}((func(_554_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_555_i16 _dafny.Int) _dafny.CodePoint {
					return _554_v3
				}
			})(_380_v3))), 1)
			(_nw93).ArraySet1(_378_v1, 2)
			(_nw93).ArraySet1(_378_v1, 3)
			(_nw93).ArraySet1(_378_v1, 4)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(195))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_556_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), 5)
			(_nw93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wfnbai"), 6)
			(_nw93).ArraySet1(_378_v1, 7)
			(_nw93).ArraySet1(_378_v1, 8)
			(_nw93).ArraySet1(_378_v1, 9)
			(_nw93).ArraySet1(_378_v1, 10)
			(_nw93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("apfximgjv"), 11)
			(_nw93).ArraySet1(_378_v1, 12)
			(_nw93).ArraySet1(_378_v1, 13)
			(_nw93).ArraySet1(_378_v1, 14)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(983))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_557_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_558_i18 _dafny.Int) _dafny.CodePoint {
					return _557_v3
				}
			})(_380_v3))), 15)
			(_nw93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gadjbunu"), 16)
			(_nw93).ArraySet1(_378_v1, 17)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}((func(_559_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_560_i19 _dafny.Int) _dafny.CodePoint {
					return _559_v3
				}
			})(_380_v3))), 18)
			(_nw93).ArraySet1(_378_v1, 19)
			(_nw93).ArraySet1(_378_v1, 20)
			(_nw93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gwlwggmnv"), 21)
			(_nw93).ArraySet1(_378_v1, 22)
			(_nw93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ukpupoan"), 23)
			_553_v127 = _nw93
			var _561_v128 _dafny.MultiSet
			_ = _561_v128
			_561_v128 = _dafny.MultiSetOf(_381_v4)
			var _562_v129 T2
			_ = _562_v129
			var _nw94 *C6 = New_C6_()
			_ = _nw94
			_nw94.Ctor__(_381_v4, _383_v6, _dafny.MultiSetOf(_381_v4, _381_v4), _386_v9, (_384_v7).Cardinality(), _553_v127, _561_v128, _383_v6)
			_562_v129 = _nw94
			var _563_v130 _dafny.Set
			_ = _563_v130
			_563_v130 = _dafny.SetOf(_562_v129)
			(_391_globalState).F21 = (func() _dafny.Int {
				if (_563_v130).IsDisjointFrom(_563_v130) {
					return (_383_v6).Plus(_383_v6)
				}
				return _dafny.IntOfUint32((_378_v1).Cardinality())
			})()
			var _564_v131 _dafny.Map
			_ = _564_v131
			_564_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, (_562_v129).F32())
			(_391_globalState).F15 = (func() _dafny.Int {
				if (_564_v131).Contains(_377_v0) {
					return (_564_v131).Get(_377_v0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(Companion_Default___.Fm1(_391_globalState))
			})()
		}
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))
		_ = _index37
		(_388_v11).ArraySet1((_381_v4) || (Companion_Default___.Fm0(_383_v6, _391_globalState)), (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))
		_ = _index38
		(_388_v11).ArraySet1(!(_381_v4) || (_381_v4), (_index38).Int())
		var _565_v132 _dafny.MultiSet
		_ = _565_v132
		_565_v132 = _dafny.MultiSetOf(_381_v4, (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool))
		var _566_v133 *C2
		_ = _566_v133
		var _nw95 *C2 = New_C2_()
		_ = _nw95
		_nw95.Ctor__((_565_v132).Update(Companion_Default___.Fm0((_dafny.Zero).Minus(_383_v6), _391_globalState), Companion_Default___.Abs(_dafny.IntOfInt64(869))), _383_v6)
		_566_v133 = _nw95
		var _567_v134 _dafny.Map
		_ = _567_v134
		_567_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-78), _566_v133)
		var _568_v136 _dafny.Sequence
		_ = _568_v136
		_568_v136 = _dafny.SeqOf(func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(987), _dafny.IntOfInt64(775))); ; {
				_compr_22, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _569_v135 _dafny.Int
				_569_v135 = interface{}(_compr_22).(_dafny.Int)
				if ((_dafny.IntOfInt64(987)).Cmp(_569_v135) <= 0) && ((_569_v135).Cmp(_dafny.IntOfInt64(775)) < 0) {
					_coll22.Add((_569_v135).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), _383_v6)).Cardinality()), (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool))
				}
			}
			return _coll22.ToMap()
		}())
		var _570_v137 _dafny.Sequence
		_ = _570_v137
		_570_v137 = _dafny.SeqOf(_566_v133, _566_v133)
		_566_v133 = (func() *C2 {
			if (_567_v134).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_568_v136, _568_v136)).Cardinality())) {
				return (_567_v134).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_568_v136, _568_v136)).Cardinality())).(*C2)
			}
			return (_570_v137).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm48((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), _381_v4, _383_v6, _391_globalState)).Cardinality(), _dafny.IntOfUint32((_570_v137).Cardinality()))).Uint32()).(*C2)
		})()
		var _571_v138 _dafny.Array
		_ = _571_v138
		var _nwElement0_18 _dafny.Sequence = _378_v1
		_ = _nwElement0_18
		var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(3))
		_ = _nw96
		(_nw96).ArraySet1(_nwElement0_18, 0)
		(_nw96).ArraySet1(Companion_Default___.Fm37(_383_v6, _383_v6, _391_globalState), 1)
		(_nw96).ArraySet1(_378_v1, 2)
		_571_v138 = _nw96
		var _572_v139 D17
		_ = _572_v139
		_572_v139 = Companion_D17_.Create_DC43_(_571_v138)
		var _source12 D17 = _572_v139
		_ = _source12
		if _source12.Is_DC44() {
			var _573___mcc_h5 _dafny.Int = _source12.Get_().(D17_DC44).Cf83
			_ = _573___mcc_h5
			var _574_cf83 _dafny.Int = _573___mcc_h5
			_ = _574_cf83
			var _575_v140 _dafny.Sequence
			_ = _575_v140
			_575_v140 = _dafny.SeqOf(_381_v4)
			var _576_v141 _dafny.Array
			_ = _576_v141
			var _nwElement0_19 bool = (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)
			_ = _nwElement0_19
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(23))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_19, 0)
			(_nw97).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 1)
			(_nw97).ArraySet1(false, 2)
			(_nw97).ArraySet1(_381_v4, 3)
			(_nw97).ArraySet1(_381_v4, 4)
			(_nw97).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 5)
			(_nw97).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 6)
			(_nw97).ArraySet1(_381_v4, 7)
			(_nw97).ArraySet1((func() bool {
				if false {
					return (_566_v133).Fm5(_391_globalState)
				}
				return _381_v4
			})(), 8)
			(_nw97).ArraySet1(true, 9)
			(_nw97).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 10)
			(_nw97).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_575_v140, _575_v140), 11)
			(_nw97).ArraySet1(!(_381_v4) || ((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)), 12)
			(_nw97).ArraySet1(_381_v4, 13)
			(_nw97).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 14)
			(_nw97).ArraySet1((Companion_Default___.Fm1(_391_globalState)).Cmp(_383_v6) == 0, 15)
			(_nw97).ArraySet1((_566_v133).Fm21(_391_globalState), 16)
			(_nw97).ArraySet1(true, 17)
			(_nw97).ArraySet1(false, 18)
			(_nw97).ArraySet1((_381_v4) && ((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)), 19)
			(_nw97).ArraySet1(_381_v4, 20)
			(_nw97).ArraySet1(_381_v4, 21)
			(_nw97).ArraySet1((_565_v132).IsProperSubsetOf(_dafny.MultiSetFromSeq(_575_v140)), 22)
			_576_v141 = _nw97
			var _577_v142 _dafny.Set
			_ = _577_v142
			_577_v142 = _dafny.SetOf(_576_v141, _576_v141)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))
			_ = _index39
			var _rhs46 _dafny.Array = _576_v141
			_ = _rhs46
			var _rhs47 _dafny.Sequence = _378_v1
			_ = _rhs47
			var _rhs48 bool = ((_577_v142).Difference(_577_v142)).IsProperSubsetOf(_dafny.SetOf(_576_v141))
			_ = _rhs48
			var _lhs30 _dafny.Array = _388_v11
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))
			_ = _lhs31
			_576_v141 = _rhs46
			_378_v1 = _rhs47
			(_lhs30).ArraySet1(_rhs48, (_lhs31).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))
			_ = _index40
			(_388_v11).ArraySet1(!((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)), (_index40).Int())
			var _578_v143 _dafny.Map
			_ = _578_v143
			_578_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, _377_v0)
			var _579_v144 D15
			_ = _579_v144
			_579_v144 = Companion_D15_.Create_DC37_((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), _574_cf83, _377_v0, (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool))
			var _580_v145 _dafny.Array
			_ = _580_v145
			var _nwElement0_20 _dafny.Array = _377_v0
			_ = _nwElement0_20
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(18))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_20, 0)
			(_nw98).ArraySet1(_377_v0, 1)
			(_nw98).ArraySet1(_377_v0, 2)
			(_nw98).ArraySet1((func() _dafny.Array {
				if (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool) {
					return _377_v0
				}
				return (func() _dafny.Array {
					if (_578_v143).Contains((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)) {
						return (_578_v143).Get((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)).(_dafny.Array)
					}
					return _377_v0
				})()
			})(), 3)
			(_nw98).ArraySet1(_377_v0, 4)
			(_nw98).ArraySet1((_579_v144).Dtor_cf68(), 5)
			(_nw98).ArraySet1(_377_v0, 6)
			(_nw98).ArraySet1(_377_v0, 7)
			(_nw98).ArraySet1(_377_v0, 8)
			(_nw98).ArraySet1(_377_v0, 9)
			(_nw98).ArraySet1(_377_v0, 10)
			(_nw98).ArraySet1(_377_v0, 11)
			(_nw98).ArraySet1(_377_v0, 12)
			(_nw98).ArraySet1(_377_v0, 13)
			(_nw98).ArraySet1(_377_v0, 14)
			(_nw98).ArraySet1(_377_v0, 15)
			(_nw98).ArraySet1(_377_v0, 16)
			(_nw98).ArraySet1(_377_v0, 17)
			_580_v145 = _nw98
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_580_v145), 0))
			_ = _index41
			(_580_v145).ArraySet1(_377_v0, (_index41).Int())
			var _581_v146 D12
			_ = _581_v146
			_581_v146 = Companion_D12_.Create_DC28_(_574_cf83, _383_v6, _381_v4)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_580_v145), 0))
			_ = _index42
			var _nwElement0_21 _dafny.Int = _dafny.IntOfInt64(-444)
			_ = _nwElement0_21
			var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(19))
			_ = _nw99
			(_nw99).ArraySet1(_nwElement0_21, 0)
			(_nw99).ArraySet1(_574_cf83, 1)
			(_nw99).ArraySet1((_574_cf83).Times(_574_cf83), 2)
			(_nw99).ArraySet1(_dafny.IntOfUint32((_575_v140).Cardinality()), 3)
			(_nw99).ArraySet1(_dafny.IntOfInt64(963), 4)
			(_nw99).ArraySet1(_383_v6, 5)
			(_nw99).ArraySet1(_383_v6, 6)
			(_nw99).ArraySet1((_383_v6).Times(_383_v6), 7)
			(_nw99).ArraySet1((_383_v6).Minus(_383_v6), 8)
			(_nw99).ArraySet1(_dafny.IntOfInt64(909), 9)
			(_nw99).ArraySet1(_383_v6, 10)
			(_nw99).ArraySet1(_383_v6, 11)
			(_nw99).ArraySet1(((_581_v146).Dtor_cf52()).Times(_383_v6), 12)
			(_nw99).ArraySet1((_574_cf83).Times(_dafny.IntOfUint32((Companion_Default___.Fm9((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), _391_globalState)).Cardinality())), 13)
			(_nw99).ArraySet1(_dafny.IntOfUint32((_378_v1).Cardinality()), 14)
			(_nw99).ArraySet1(_383_v6, 15)
			(_nw99).ArraySet1(_574_cf83, 16)
			(_nw99).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm1(_391_globalState)), 17)
			(_nw99).ArraySet1(_383_v6, 18)
			(_580_v145).ArraySet1(_nw99, (_index42).Int())
			(_391_globalState).F15 = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_381_v4), _dafny.SeqOf(true))).Cardinality())).Times((_383_v6).Plus(_574_cf83)))
		} else if _source12.Is_DC43() {
			var _582___mcc_h6 _dafny.Array = _source12.Get_().(D17_DC43).Cf82
			_ = _582___mcc_h6
			var _583_cf82 _dafny.Array = _582___mcc_h6
			_ = _583_cf82
			_381_v4 = _381_v4
			var _584_v147 _dafny.Sequence
			_ = _584_v147
			_584_v147 = _dafny.SeqOf(_386_v9)
			var _585_v148 _dafny.Set
			_ = _585_v148
			_585_v148 = _dafny.SetOf((_584_v147).Select((Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_584_v147).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm43(_381_v4, _381_v4, _381_v4, _391_globalState))
			(_391_globalState).F15 = (_585_v148).Cardinality()
			var _586_v149 _dafny.Map
			_ = _586_v149
			_586_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_386_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int))
			var _587_v150 _dafny.Map
			_ = _587_v150
			_587_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_586_v149, (_dafny.IntOfInt64(423)).Cmp(_383_v6) <= 0)
			var _588_v151 *C1
			_ = _588_v151
			var _nw100 *C1 = New_C1_()
			_ = _nw100
			_nw100.Ctor__(false, _565_v132, _386_v9, _383_v6, _583_cf82)
			_588_v151 = _nw100
			var _589_v152 _dafny.Map
			_ = _589_v152
			_589_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_588_v151, _585_v148)
			_587_v150 = (_587_v150).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, ((func() _dafny.Set {
				if (_589_v152).Contains(_588_v151) {
					return (_589_v152).Get(_588_v151).(_dafny.Set)
				}
				return _585_v148
			})()).Cardinality()), !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)), _dafny.SeqOf(_588_v151.F35, false, _588_v151.F35, _381_v4, (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool))))
			var _590_v153 _dafny.Map
			_ = _590_v153
			_590_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(664), _382_v5)
			var _591_v154 D11
			_ = _591_v154
			_591_v154 = Companion_D11_.Create_DC23_(_dafny.MultiSetOf(_384_v7))
			var _rhs49 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.SetOf(_588_v151.F35, (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)))
			_ = _rhs49
			var _rhs50 D11 = _591_v154
			_ = _rhs50
			_590_v153 = _rhs49
			_591_v154 = _rhs50
		} else {
			var _592___mcc_h7 D17 = _source12.Get_().(D17_DC45).Cf84
			_ = _592___mcc_h7
			var _593_cf84 D17 = _592___mcc_h7
			_ = _593_cf84
			var _594_v155 _dafny.Array
			_ = _594_v155
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_12
			var _nw101 _dafny.Array
			_ = _nw101
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw101 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Sequence = func(_595_i20 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(false)
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw101 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw101).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw101).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_594_v155 = _nw101
			var _596_v156 _dafny.Map
			_ = _596_v156
			_596_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v155, _386_v9)
			_386_v9 = (func() _dafny.Sequence {
				if (_596_v156).Contains(_594_v155) {
					return (_596_v156).Get(_594_v155).(_dafny.Sequence)
				}
				return _386_v9
			})()
			(_391_globalState).F26 = (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)
			var _597_v157 *C3
			_ = _597_v157
			var _nw102 *C3 = New_C3_()
			_ = _nw102
			_nw102.Ctor__(_dafny.IntOfInt64(-958), _571_v138)
			_597_v157 = _nw102
			var _598_v158 D15
			_ = _598_v158
			_598_v158 = Companion_D15_.Create_DC36_(_380_v3, _383_v6)
			_598_v158 = _598_v158
		}
	}
	var _599_v159 _dafny.Map
	_ = _599_v159
	_599_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_383_v6), _383_v6)
	var _600_v160 _dafny.Map
	_ = _600_v160
	_600_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_599_v159, Companion_Default___.Fm1(_391_globalState))
	var _601_v161 _dafny.MultiSet
	_ = _601_v161
	_601_v161 = _dafny.MultiSetOf(_381_v4)
	var _602_v162 _dafny.Array
	_ = _602_v162
	var _nwElement0_22 _dafny.Sequence = _378_v1
	_ = _nwElement0_22
	var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
	_ = _nw103
	(_nw103).ArraySet1(_nwElement0_22, 0)
	(_nw103).ArraySet1(_378_v1, 1)
	(_nw103).ArraySet1(_378_v1, 2)
	(_nw103).ArraySet1(_378_v1, 3)
	(_nw103).ArraySet1(_378_v1, 4)
	_602_v162 = _nw103
	var _603_v163 *C5
	_ = _603_v163
	var _nw104 *C5 = New_C5_()
	_ = _nw104
	_nw104.Ctor__(_600_v160, (_601_v161).Update(_381_v4, Companion_Default___.Abs((_382_v5).Cardinality())), (Companion_D10_.Create_DC21_(_386_v9)).Dtor_cf39(), _383_v6, _602_v162)
	_603_v163 = _nw104
	var _hi1 _dafny.Int = _383_v6
	_ = _hi1
	for _604_i21 := _383_v6; _604_i21.Cmp(_hi1) < 0; _604_i21 = _604_i21.Plus(_dafny.One) {
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_377_v0), 0))
		_ = _index43
		(_377_v0).ArraySet1(_dafny.IntOfInt64(-775), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_377_v0), 0))
		_ = _index44
		(_377_v0).ArraySet1(_604_i21, (_index44).Int())
		var _605_v164 _dafny.Int
		_ = _605_v164
		var _606_v165 _dafny.Int
		_ = _606_v165
		var _607_v166 _dafny.Int
		_ = _607_v166
		var _out20 _dafny.Int
		_ = _out20
		var _out21 _dafny.Int
		_ = _out21
		var _out22 _dafny.Int
		_ = _out22
		_out20, _out21, _out22 = (_603_v163).M1(_391_globalState)
		_605_v164 = _out20
		_606_v165 = _out21
		_607_v166 = _out22
		var _608_v167 _dafny.Sequence
		_ = _608_v167
		_608_v167 = _dafny.SeqOf(_dafny.MultiSetOf(_381_v4, _381_v4))
		var _609_v168 D19
		_ = _609_v168
		_609_v168 = Companion_D19_.Create_DC49_(_381_v4, _381_v4, _604_i21)
		var _610_v169 _dafny.Sequence
		_ = _610_v169
		_610_v169 = _dafny.SeqOf(_602_v162)
		var _pat_let_tv5 = _381_v4
		_ = _pat_let_tv5
		var _611_v170 *C6
		_ = _611_v170
		var _nw105 *C6 = New_C6_()
		_ = _nw105
		_nw105.Ctor__(_381_v4, (_dafny.IntOfInt64(615)).Plus(_383_v6), (_608_v167).Select((Companion_Default___.SafeIndex(_604_i21, _dafny.IntOfUint32((_608_v167).Cardinality()))).Uint32()).(_dafny.MultiSet), _dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (func(_pat_let10_0 D19) D19 {
			return func(_612_dt__update__tmp_h2 D19) D19 {
				return func(_pat_let11_0 bool) D19 {
					return func(_613_dt__update_hcf89_h0 bool) D19 {
						return Companion_D19_.Create_DC49_(_613_dt__update_hcf89_h0, (_612_dt__update__tmp_h2).Dtor_cf90(), (_612_dt__update__tmp_h2).Dtor_cf91())
					}(_pat_let11_0)
				}(_pat_let_tv5)
			}(_pat_let10_0)
		}(_609_v168)).Dtor_cf91()), _383_v6, (_610_v169).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.IntOfUint32((_610_v169).Cardinality()))).Uint32()).(_dafny.Array), _601_v161, _606_v165)
		_611_v170 = _nw105
		_611_v170 = _611_v170
		var _614_v171 _dafny.MultiSet
		_ = _614_v171
		_614_v171 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ugxwp")).Cardinality()))
		(_391_globalState).F26 = (((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_386_v9).Cardinality()))).Cmp((_dafny.IntOfUint32((_386_v9).Cardinality())).Plus((_614_v171).Cardinality())) >= 0
	}
	var _615_v172 D4
	_ = _615_v172
	_615_v172 = Companion_D4_.Create_DC9_(_383_v6, false)
	var _616_v173 _dafny.Map
	_ = _616_v173
	_616_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, (_615_v172).Dtor_cf17())
	_616_v173 = (_616_v173).Update(_381_v4, _381_v4)
	var _617_v174 *C3
	_ = _617_v174
	var _nw106 *C3 = New_C3_()
	_ = _nw106
	_nw106.Ctor__(_dafny.IntOfInt64(923), _602_v162)
	_617_v174 = _nw106
	var _618_v175 _dafny.Map
	_ = _618_v175
	_618_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_378_v1).Cardinality()), _617_v174)
	var _619_v176 _dafny.Map
	_ = _619_v176
	_619_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_618_v175, _383_v6)
	_619_v176 = (_619_v176).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_383_v6), _617_v174), _383_v6)
	var _620_v177 _dafny.MultiSet
	_ = _620_v177
	_620_v177 = _dafny.MultiSetOf(_383_v6, _383_v6)
	if ((_603_v163).Fm4(_dafny.SetOf(_378_v1, _378_v1, _dafny.SeqOf(Companion_Default___.Fm26(_380_v3, (_620_v177).Cardinality(), _383_v6, _391_globalState))), _381_v4, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(!(_381_v4)), (Companion_Default___.SafeIndex((_dafny.SetOf(_381_v4)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(!(_381_v4))).Cardinality()))).Uint32(), _381_v4), _391_globalState)).Cmp(_383_v6) < 0 {
		(_391_globalState).F26 = true
		_378_v1 = _dafny.Companion_Sequence_.Concatenate(_378_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_621_i22 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))
		(_391_globalState).F0 = _dafny.IntOfInt64(522)
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_388_v11), 0))
		_ = _index45
		(_388_v11).ArraySet1(!(_381_v4), (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_388_v11), 0))
		_ = _index46
		(_388_v11).ArraySet1(_381_v4, (_index46).Int())
		_620_v177 = _620_v177
	} else {
		var _622_v178 _dafny.Sequence
		_ = _622_v178
		_622_v178 = _dafny.SeqOf(_378_v1)
		_378_v1 = (_622_v178).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_383_v6, _383_v6), _dafny.IntOfUint32((_622_v178).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _623_v179 _dafny.Array
		_ = _623_v179
		var _624_v180 bool
		_ = _624_v180
		var _625_v181 bool
		_ = _625_v181
		var _out23 _dafny.Array
		_ = _out23
		var _out24 bool
		_ = _out24
		var _out25 bool
		_ = _out25
		_out23, _out24, _out25 = (_603_v163).M9(_391_globalState)
		_623_v179 = _out23
		_624_v180 = _out24
		_625_v181 = _out25
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_377_v0), 0))
		_ = _index47
		(_377_v0).ArraySet1(_dafny.IntOfUint32((_378_v1).Cardinality()), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_377_v0), 0))
		_ = _index48
		(_377_v0).ArraySet1(_383_v6, (_index48).Int())
		_380_v3 = (func() _dafny.CodePoint {
			if _625_v181 {
				return _380_v3
			}
			return _380_v3
		})()
		var _626_v182 _dafny.Int
		_ = _626_v182
		var _627_v183 _dafny.Int
		_ = _627_v183
		var _628_v184 _dafny.Int
		_ = _628_v184
		var _out26 _dafny.Int
		_ = _out26
		var _out27 _dafny.Int
		_ = _out27
		var _out28 _dafny.Int
		_ = _out28
		_out26, _out27, _out28 = (_603_v163).M1(_391_globalState)
		_626_v182 = _out26
		_627_v183 = _out27
		_628_v184 = _out28
	}
	_599_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_383_v6)).Cardinality()), (_dafny.Zero).Minus(_383_v6)), _383_v6)
	var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))
	_ = _index49
	(_377_v0).ArraySet1(_383_v6, (_index49).Int())
	var _629_v185 _dafny.Sequence
	_ = _629_v185
	_629_v185 = _dafny.SeqOf(_381_v4)
	var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))
	_ = _index50
	(_377_v0).ArraySet1((_dafny.Zero).Minus(((_603_v163).Fm4(_390_v13, _381_v4, _629_v185, _391_globalState)).Plus(Companion_Default___.SafeDivisionInt((_386_v9).Select((Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int), _383_v6))), (_index50).Int())
	var _rhs51 bool = (func() bool {
		if _381_v4 {
			return _381_v4
		}
		return _381_v4
	})()
	_ = _rhs51
	var _rhs52 _dafny.Int = _383_v6
	_ = _rhs52
	var _rhs53 bool = (func() bool {
		if (func() bool {
			if _381_v4 {
				return _381_v4
			}
			return _381_v4
		})() {
			return _381_v4
		}
		return _381_v4
	})()
	_ = _rhs53
	var _rhs54 _dafny.Int = _dafny.IntOfInt64(451)
	_ = _rhs54
	var _lhs32 *GlobalState = _391_globalState
	_ = _lhs32
	var _lhs33 *GlobalState = _391_globalState
	_ = _lhs33
	_lhs32.F26 = _rhs51
	_lhs33.F18 = _rhs52
	_381_v4 = _rhs53
	_383_v6 = _rhs54
	_388_v11 = _388_v11
	(_391_globalState).F26 = ((_616_v173).Cardinality()).Cmp((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)) != 0
	var _hi2 _dafny.Int = (_383_v6).Minus((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))
	_ = _hi2
	for _630_i23 := (_617_v174).Fm45(_391_globalState); _630_i23.Cmp(_hi2) < 0; _630_i23 = _630_i23.Plus(_dafny.One) {
		var _631_v186 D16
		_ = _631_v186
		_631_v186 = Companion_D16_.Create_DC40_(_381_v4, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), ((_385_v8).Update(_380_v3, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Cardinality())
		var _632_v187 _dafny.Map
		_ = _632_v187
		_632_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(419), _380_v3)
		var _633_v188 D4
		_ = _633_v188
		_633_v188 = Companion_D4_.Create_DC8_((func() _dafny.CodePoint {
			if (_632_v187).Contains((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)) {
				return (_632_v187).Get((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).(_dafny.CodePoint)
			}
			return _380_v3
		})(), _dafny.IntOfUint32((_378_v1).Cardinality()), _381_v4)
		var _pat_let_tv6 = _380_v3
		_ = _pat_let_tv6
		var _634_v189 _dafny.Map
		_ = _634_v189
		_634_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, func(_pat_let12_0 D4) D4 {
			return func(_635_dt__update__tmp_h3 D4) D4 {
				return func(_pat_let13_0 _dafny.CodePoint) D4 {
					return func(_636_dt__update_hcf13_h0 _dafny.CodePoint) D4 {
						return Companion_D4_.Create_DC8_(_636_dt__update_hcf13_h0, (_635_dt__update__tmp_h3).Dtor_cf14(), (_635_dt__update__tmp_h3).Dtor_cf15())
					}(_pat_let13_0)
				}(_pat_let_tv6)
			}(_pat_let12_0)
		}(_633_v188))
		var _637_v190 _dafny.MultiSet
		_ = _637_v190
		_637_v190 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, _633_v188)).Update(_381_v4, _633_v188), _634_v189)
		var _638_v191 D2
		_ = _638_v191
		_638_v191 = Companion_D2_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _381_v4))
		var _639_v192 D17
		_ = _639_v192
		_639_v192 = Companion_D17_.Create_DC43_(_602_v162)
		var _640_v193 D17
		_ = _640_v193
		_640_v193 = Companion_D17_.Create_DC45_(_639_v192)
		var _641_v194 _dafny.Sequence
		_ = _641_v194
		_641_v194 = _dafny.SeqOf(_640_v193, Companion_D17_.Create_DC45_(_639_v192), _640_v193)
		var _642_v195 _dafny.Map
		_ = _642_v195
		_642_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_641_v194, _381_v4)
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))
		_ = _index51
		var _rhs55 bool = !(!(((_603_v163).Fm4(_390_v13, false, _629_v185, _391_globalState)).Cmp(_630_i23) > 0))
		_ = _rhs55
		var _rhs56 D16 = (func() D16 {
			if (!(true)) || (_381_v4) {
				return Companion_Default___.Fm55(_383_v6, _638_v191, _381_v4, (_dafny.SetOf(_381_v4, _381_v4)).Cardinality(), _391_globalState)
			}
			return Companion_D16_.Create_DC40_(_381_v4, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _630_i23)
		})()
		_ = _rhs56
		var _rhs57 _dafny.MultiSet = _637_v190
		_ = _rhs57
		var _rhs58 _dafny.Int = (_642_v195).Cardinality()
		_ = _rhs58
		var _rhs59 bool = _381_v4
		_ = _rhs59
		var _lhs34 *GlobalState = _391_globalState
		_ = _lhs34
		var _lhs35 _dafny.Array = _377_v0
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))
		_ = _lhs36
		_lhs34.F26 = _rhs55
		_631_v186 = _rhs56
		_637_v190 = _rhs57
		(_lhs35).ArraySet1(_rhs58, (_lhs36).Int())
		_381_v4 = _rhs59
		(_391_globalState).F21 = _630_i23
		var _643_v196 D2
		_ = _643_v196
		_643_v196 = Companion_D2_.Create_DC4_(_383_v6, _dafny.IntOfUint32((_378_v1).Cardinality()), (_617_v174).Fm45(_391_globalState), _380_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg50 _dafny.Int) interface{} {
				return coer50(arg50)
			}
		}(func(_644_i24 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})))
		_380_v3 = (func() _dafny.CodePoint {
			if ((_620_v177).Update(_630_i23, Companion_Default___.Abs(_630_i23))).Equals(_620_v177) {
				return _380_v3
			}
			return (_643_v196).Dtor_cf8()
		})()
		if _381_v4 {
			var _645_v197 _dafny.Array
			_ = _645_v197
			var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw107
			_645_v197 = _nw107
			_645_v197 = _645_v197
			var _646_v198 D15
			_ = _646_v198
			_646_v198 = Companion_D15_.Create_DC37_(_381_v4, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _377_v0, true)
			var _647_v199 _dafny.Map
			_ = _647_v199
			_647_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_i23, false)
			var _648_v200 _dafny.Map
			_ = _648_v200
			_648_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_381_v4), _383_v6)
			var _649_v201 T2
			_ = _649_v201
			var _nw108 *C2 = New_C2_()
			_ = _nw108
			_nw108.Ctor__(_601_v161, (_648_v200).Cardinality())
			_649_v201 = _nw108
			var _650_v202 _dafny.Array
			_ = _650_v202
			var _nwElement0_23 D15 = _646_v198
			_ = _nwElement0_23
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_23, 0)
			(_nw109).ArraySet1(_646_v198, 1)
			(_nw109).ArraySet1(_646_v198, 2)
			(_nw109).ArraySet1(Companion_D15_.Create_DC37_(_381_v4, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _377_v0, _381_v4), 3)
			(_nw109).ArraySet1(_646_v198, 4)
			(_nw109).ArraySet1(Companion_D15_.Create_DC37_(_381_v4, (_647_v199).Cardinality(), _377_v0, _381_v4), 5)
			(_nw109).ArraySet1(_646_v198, 6)
			(_nw109).ArraySet1(_646_v198, 7)
			(_nw109).ArraySet1(_646_v198, 8)
			(_nw109).ArraySet1((func() D15 {
				if false {
					return _646_v198
				}
				return _646_v198
			})(), 9)
			(_nw109).ArraySet1(_646_v198, 10)
			(_nw109).ArraySet1(Companion_D15_.Create_DC37_(!((_603_v163).Fm3(_601_v161, _391_globalState)), _dafny.IntOfUint32((_dafny.SeqOf(_649_v201)).Cardinality()), _377_v0, _381_v4), 11)
			(_nw109).ArraySet1(_646_v198, 12)
			_650_v202 = _nw109
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_602_v162), 0))
			_ = _index52
			(_602_v162).ArraySet1(_378_v1, (_index52).Int())
			var _651_v203 _dafny.Sequence
			_ = _651_v203
			_651_v203 = _dafny.SeqOf(_382_v5)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_602_v162), 0))
			_ = _index53
			var _rhs60 _dafny.Int = Companion_Default___.SafeModuloInt((_383_v6).Plus(_630_i23), _630_i23)
			_ = _rhs60
			var _rhs61 bool = (Companion_Default___.SafeDivisionInt((_649_v201).F32(), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Cmp((_386_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_651_v203).Cardinality()), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int)) >= 0
			_ = _rhs61
			var _rhs62 _dafny.Array = _650_v202
			_ = _rhs62
			var _rhs63 _dafny.Sequence = _378_v1
			_ = _rhs63
			var _rhs64 _dafny.Int = _383_v6
			_ = _rhs64
			var _lhs37 *GlobalState = _391_globalState
			_ = _lhs37
			var _lhs38 *GlobalState = _391_globalState
			_ = _lhs38
			var _lhs39 _dafny.Array = _602_v162
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_602_v162), 0))
			_ = _lhs40
			var _lhs41 *GlobalState = _391_globalState
			_ = _lhs41
			_lhs37.F18 = _rhs60
			_lhs38.F26 = _rhs61
			_650_v202 = _rhs62
			(_lhs39).ArraySet1(_rhs63, (_lhs40).Int())
			_lhs41.F3 = _rhs64
			var _652_v204 _dafny.Set
			_ = _652_v204
			_652_v204 = _dafny.SetOf(_380_v3)
			var _653_v205 _dafny.Map
			_ = _653_v205
			_653_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v4, _652_v204)
			var _654_v206 _dafny.Map
			_ = _654_v206
			_654_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_616_v173, (func() _dafny.Set {
				if (_653_v205).Contains(_381_v4) {
					return (_653_v205).Get(_381_v4).(_dafny.Set)
				}
				return _652_v204
			})())
			_381_v4 = (func() bool {
				if _dafny.Companion_Sequence_.Equal(_629_v185, _629_v185) {
					return (func() bool {
						if (_647_v199).Contains((_654_v206).Cardinality()) {
							return (_647_v199).Get((_654_v206).Cardinality()).(bool)
						}
						return _381_v4
					})()
				}
				return _381_v4
			})()
			_383_v6 = _dafny.IntOfInt64(897)
			(_391_globalState).F0 = _383_v6
		} else {
			var _655_v207 _dafny.Array
			_ = _655_v207
			var _nwElement0_24 _dafny.Array = _388_v11
			_ = _nwElement0_24
			var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(16))
			_ = _nw110
			(_nw110).ArraySet1(_nwElement0_24, 0)
			(_nw110).ArraySet1(_388_v11, 1)
			(_nw110).ArraySet1(_388_v11, 2)
			(_nw110).ArraySet1(_388_v11, 3)
			(_nw110).ArraySet1(_388_v11, 4)
			(_nw110).ArraySet1(_388_v11, 5)
			(_nw110).ArraySet1(_388_v11, 6)
			(_nw110).ArraySet1(_388_v11, 7)
			(_nw110).ArraySet1(_388_v11, 8)
			(_nw110).ArraySet1(_388_v11, 9)
			(_nw110).ArraySet1(_388_v11, 10)
			(_nw110).ArraySet1(_388_v11, 11)
			(_nw110).ArraySet1(_388_v11, 12)
			(_nw110).ArraySet1(_388_v11, 13)
			(_nw110).ArraySet1(_388_v11, 14)
			(_nw110).ArraySet1(_388_v11, 15)
			_655_v207 = _nw110
			var _656_v209 _dafny.Map
			_ = _656_v209
			_656_v209 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v3, _381_v4)
			var _657_v210 D2
			_ = _657_v210
			_657_v210 = Companion_D2_.Create_DC3_(func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate((_656_v209).Keys().Elements()); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _658_v208 _dafny.CodePoint
					_658_v208 = interface{}(_compr_23).(_dafny.CodePoint)
					if (_656_v209).Contains(_658_v208) {
						_coll23.Add(_658_v208, _381_v4)
					}
				}
				return _coll23.ToMap()
			}())
			var _659_v211 D2
			_ = _659_v211
			_659_v211 = Companion_D2_.Create_DC5_(_657_v210)
			var _660_v212 D2
			_ = _660_v212
			_660_v212 = Companion_D2_.Create_DC5_(_659_v211)
			(_391_globalState).F26 = (func() bool {
				if (_616_v173).Contains(Companion_Default___.Fm0(_630_i23, _391_globalState)) {
					return (_616_v173).Get(Companion_Default___.Fm0(_630_i23, _391_globalState)).(bool)
				}
				return (func() bool {
					if _381_v4 {
						return (Companion_D16_.Create_DC42_((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _655_v207, (_617_v174).Fm46(_383_v6, _380_v3, _391_globalState), _660_v212)).Dtor_cf80()
					}
					return _381_v4
				})()
			})()
			var _661_v213 bool
			_ = _661_v213
			var _662_v214 _dafny.Int
			_ = _662_v214
			var _out29 bool
			_ = _out29
			var _out30 _dafny.Int
			_ = _out30
			_out29, _out30 = Companion_Default___.M0(((_dafny.Zero).Minus((_617_v174).Fm45(_391_globalState))).Cmp((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)) > 0, _dafny.IntOfUint32((_378_v1).Cardinality()), Companion_Default___.SafeDivisionInt(_630_i23, _383_v6), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _391_globalState)
			_661_v213 = _out29
			_662_v214 = _out30
			var _663_v215 _dafny.Array
			_ = _663_v215
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_13
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Map = (func(_664_v214 _dafny.Int, _665_v4 bool) func(_dafny.Int) _dafny.Map {
					return func(_666_i25 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_664_v214, !(_665_v4))
					}
				})(_662_v214, _381_v4)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw111 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw111).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw111).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_663_v215 = _nw111
			var _667_v216 _dafny.Sequence
			_ = _667_v216
			_667_v216 = _dafny.SeqOf(_663_v215, _663_v215)
			var _668_v217 _dafny.Map
			_ = _668_v217
			_668_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if false {
					return _dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_662_v214, _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3)
				}
				return _378_v1
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_630_i23), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if false {
					return _dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_662_v214, _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3)
				}
				return _378_v1
			})()).Cardinality()))).Uint32(), _380_v3)).Cardinality()), (_667_v216).Select((Companion_Default___.SafeIndex(_662_v214, _dafny.IntOfUint32((_667_v216).Cardinality()))).Uint32()).(_dafny.Array))
			_668_v217 = (_668_v217).Update((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _663_v215)
			var _669_v218 _dafny.Sequence
			_ = _669_v218
			_669_v218 = _dafny.SeqOf(_602_v162, _602_v162, _602_v162)
			var _670_v219 _dafny.Map
			_ = _670_v219
			_670_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_669_v218).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.IntOfUint32((_669_v218).Cardinality()))).Uint32()).(_dafny.Array))
			var _671_v220 *C5
			_ = _671_v220
			var _nw112 *C5 = New_C5_()
			_ = _nw112
			_nw112.Ctor__((_603_v163).F38(), _601_v161, _386_v9, (_379_v2).Cardinality(), (func() _dafny.Array {
				if _381_v4 {
					return (func() _dafny.Array {
						if (_670_v219).Contains(_dafny.IntOfUint32((_629_v185).Cardinality())) {
							return (_670_v219).Get(_dafny.IntOfUint32((_629_v185).Cardinality())).(_dafny.Array)
						}
						return _602_v162
					})()
				}
				return _602_v162
			})())
			_671_v220 = _nw112
			var _672_v221 _dafny.Sequence
			_ = _672_v221
			_672_v221 = _dafny.SeqOf((func() _dafny.Sequence {
				if _381_v4 {
					return _378_v1
				}
				return _378_v1
			})(), _378_v1, _378_v1)
			_672_v221 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_672_v221, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.IntOfUint32((_672_v221).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_673_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_674_i26 _dafny.Int) _dafny.CodePoint {
					return _673_v3
				}
			})(_380_v3)))), (Companion_Default___.SafeIndex(_662_v214, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_672_v221, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.IntOfUint32((_672_v221).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_675_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_676_i26 _dafny.Int) _dafny.CodePoint {
					return _675_v3
				}
			})(_380_v3))))).Cardinality()))).Uint32(), Companion_Default___.Fm14(!(_661_v213), _381_v4, _381_v4, _383_v6, _391_globalState))
		}
	}
	if _381_v4 {
		var _rhs65 _dafny.Int = (func() _dafny.Int {
			if false {
				return (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
			}
			return Companion_Default___.SafeDivisionInt((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_616_v173).Cardinality())
		})()
		_ = _rhs65
		var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_386_v9, _386_v9)
		_ = _rhs66
		var _lhs42 *GlobalState = _391_globalState
		_ = _lhs42
		_lhs42.F21 = _rhs65
		_386_v9 = _rhs66
		var _677_v222 D9
		_ = _677_v222
		_677_v222 = Companion_D9_.Create_DC20_(Companion_Default___.Fm1(_391_globalState), _381_v4, ((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Cmp((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)) < 0, _381_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_380_v3, Companion_Default___.Fm12((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _381_v4, _391_globalState))).Cardinality()), _dafny.IntOfInt64(870)))
		var _source13 D9 = _677_v222
		_ = _source13
		if _source13.Is_DC18() {
			var _678___mcc_h8 bool = _source13.Get_().(D9_DC18).Cf32
			_ = _678___mcc_h8
			var _679_cf32 bool = _678___mcc_h8
			_ = _679_cf32
			_383_v6 = (_617_v174).Fm45(_391_globalState)
			var _680_v223 _dafny.Sequence
			_ = _680_v223
			_680_v223 = _dafny.SeqOf(((_601_v161).Update(_679_cf32, Companion_Default___.Abs((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)))).Union(_601_v161))
			_680_v223 = _dafny.Companion_Sequence_.Update(_680_v223, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_382_v5).Cardinality(), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_680_v223).Cardinality()))).Uint32(), Companion_Default___.Fm32(_391_globalState))
			_386_v9 = Companion_Default___.Fm9(_679_cf32, _391_globalState)
			var _681_v224 _dafny.Map
			_ = _681_v224
			_681_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_cf32, _601_v161)
			var _682_v225 D11
			_ = _682_v225
			_682_v225 = Companion_D11_.Create_DC23_(_454_v62)
			var _683_v226 _dafny.Map
			_ = _683_v226
			_683_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_681_v224).Contains(_679_cf32), _682_v225)
			_683_v226 = (_683_v226).Update((Companion_Default___.Fm1(_391_globalState)).Cmp((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)) == 0, _682_v225)
		} else if _source13.Is_DC19() {
			var _684___mcc_h9 bool = _source13.Get_().(D9_DC19).Cf33
			_ = _684___mcc_h9
			var _685_cf33 bool = _684___mcc_h9
			_ = _685_cf33
			var _686_v227 _dafny.Map
			_ = _686_v227
			_686_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_629_v185, _602_v162)
			var _687_v228 D7
			_ = _687_v228
			_687_v228 = Companion_D7_.Create_DC13_(_629_v185)
			var _688_v229 _dafny.MultiSet
			_ = _688_v229
			_688_v229 = _dafny.MultiSetOf(_599_v159)
			var _689_v230 *C6
			_ = _689_v230
			var _nw113 *C6 = New_C6_()
			_ = _nw113
			_nw113.Ctor__(_381_v4, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _601_v161, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm9(_685_cf33, _391_globalState), _dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (func() _dafny.Array {
				if (_686_v227).Contains((_687_v228).Dtor_cf23()) {
					return (_686_v227).Get((_687_v228).Dtor_cf23()).(_dafny.Array)
				}
				return _602_v162
			})(), _601_v161, (func() _dafny.Int {
				if (_688_v229).Contains(_599_v159) {
					return (_688_v229).Multiplicity(_599_v159)
				}
				return (_dafny.Zero).Minus((_382_v5).Cardinality())
			})())
			_689_v230 = _nw113
			_689_v230 = _689_v230
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_602_v162), 0))
			_ = _index54
			(_602_v162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_378_v1, _378_v1), (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_602_v162), 0))
			_ = _index55
			(_602_v162).ArraySet1(_378_v1, (_index55).Int())
			var _690_v231 D17
			_ = _690_v231
			_690_v231 = Companion_D17_.Create_DC45_(Companion_D17_.Create_DC44_(_dafny.IntOfUint32((_378_v1).Cardinality())))
			var _691_v232 D17
			_ = _691_v232
			_691_v232 = Companion_D17_.Create_DC45_(_690_v231)
			var _692_v233 D12
			_ = _692_v233
			_692_v233 = Companion_D12_.Create_DC27_(true, _620_v177, _383_v6, _377_v0)
			var _693_v234 _dafny.Map
			_ = _693_v234
			_693_v234 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v232, (func() _dafny.Array {
				if (_689_v230).F33() {
					return (_692_v233).Dtor_cf50()
				}
				return _377_v0
			})())
			_693_v234 = (_693_v234).Update(_691_v232, _377_v0)
			var _694_v235 D17
			_ = _694_v235
			_694_v235 = Companion_D17_.Create_DC44_(_383_v6)
			var _695_v236 _dafny.Array
			_ = _695_v236
			var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw114
			_695_v236 = _nw114
			var _696_v237 *C4
			_ = _696_v237
			var _nw115 *C4 = New_C4_()
			_ = _nw115
			_nw115.Ctor__((_694_v235).Dtor_cf83(), _695_v236)
			_696_v237 = _nw115
			(_391_globalState).F26 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_696_v237, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Cardinality()).Minus(_383_v6)).Cmp(_dafny.IntOfInt64(744)) > 0
		} else if _source13.Is_DC20() {
			var _697___mcc_h10 _dafny.Int = _source13.Get_().(D9_DC20).Cf34
			_ = _697___mcc_h10
			var _698___mcc_h11 bool = _source13.Get_().(D9_DC20).Cf35
			_ = _698___mcc_h11
			var _699___mcc_h12 bool = _source13.Get_().(D9_DC20).Cf36
			_ = _699___mcc_h12
			var _700___mcc_h13 bool = _source13.Get_().(D9_DC20).Cf37
			_ = _700___mcc_h13
			var _701___mcc_h14 _dafny.Map = _source13.Get_().(D9_DC20).Cf38
			_ = _701___mcc_h14
			var _702_cf38 _dafny.Map = _701___mcc_h14
			_ = _702_cf38
			var _703_cf37 bool = _700___mcc_h13
			_ = _703_cf37
			var _704_cf36 bool = _699___mcc_h12
			_ = _704_cf36
			var _705_cf35 bool = _698___mcc_h11
			_ = _705_cf35
			var _706_cf34 _dafny.Int = _697___mcc_h10
			_ = _706_cf34
			(_391_globalState).F26 = false
			_381_v4 = (_384_v7).IsProperSubsetOf(_384_v7)
			var _707_v238 _dafny.Map
			_ = _707_v238
			_707_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_388_v11, _383_v6)
			var _708_v239 _dafny.Map
			_ = _708_v239
			_708_v239 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_381_v4), _378_v1)
			var _709_v240 _dafny.Sequence
			_ = _709_v240
			_709_v240 = _dafny.SeqOf(_707_v238, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_388_v11, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_708_v239).Contains(_381_v4) {
					return (_708_v239).Get(_381_v4).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_378_v1, (Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_378_v1).Cardinality()))).Uint32(), _380_v3)
			})()).Cardinality())))
			var _710_v241 _dafny.Map
			_ = _710_v241
			_710_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _707_v238)
			var _711_v242 _dafny.Set
			_ = _711_v242
			_711_v242 = _dafny.SetOf(_380_v3, _380_v3)
			(_391_globalState).F26 = ((_709_v240).Select((Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_709_v240).Cardinality()))).Uint32()).(_dafny.Map)).Equals((func() _dafny.Map {
				if (_710_v241).Contains((_711_v242).Cardinality()) {
					return (_710_v241).Get((_711_v242).Cardinality()).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_388_v11, (func() _dafny.Int {
					if (_601_v161).Contains(_703_cf37) {
						return (_601_v161).Multiplicity(_703_cf37)
					}
					return (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
				})())
			})())
			(_391_globalState).F18 = ((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Plus((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))
		} else {
			var _712___mcc_h15 _dafny.Set = _source13.Get_().(D9_DC17).Cf31
			_ = _712___mcc_h15
			var _713_cf31 _dafny.Set = _712___mcc_h15
			_ = _713_cf31
			(_603_v163).M10(false, _dafny.Companion_Sequence_.Concatenate(_386_v9, _386_v9), _391_globalState)
			var _714_v243 _dafny.Set
			_ = _714_v243
			_714_v243 = _dafny.SetOf(_377_v0, _377_v0)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_388_v11), 0))
			_ = _index56
			(_388_v11).ArraySet1((_714_v243).IsDisjointFrom(_714_v243), (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_388_v11), 0))
			_ = _index57
			var _rhs67 bool = _381_v4
			_ = _rhs67
			var _rhs68 bool = _381_v4
			_ = _rhs68
			var _lhs43 _dafny.Array = _388_v11
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_388_v11), 0))
			_ = _lhs44
			var _lhs45 *GlobalState = _391_globalState
			_ = _lhs45
			(_lhs43).ArraySet1(_rhs67, (_lhs44).Int())
			_lhs45.F26 = _rhs68
			_377_v0 = _377_v0
			var _715_v244 _dafny.Array
			_ = _715_v244
			var _nwElement0_25 bool = true
			_ = _nwElement0_25
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(6))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_25, 0)
			(_nw116).ArraySet1(!((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)), 1)
			(_nw116).ArraySet1(false, 2)
			(_nw116).ArraySet1((_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool), 3)
			(_nw116).ArraySet1(_381_v4, 4)
			(_nw116).ArraySet1((_603_v163).Fm3(_601_v161, _391_globalState), 5)
			_715_v244 = _nw116
			var _716_v245 *C4
			_ = _716_v245
			var _nw117 *C4 = New_C4_()
			_ = _nw117
			_nw117.Ctor__(_dafny.IntOfInt64(231), _602_v162)
			_716_v245 = _nw117
			var _717_v246 _dafny.Map
			_ = _717_v246
			_717_v246 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_715_v244, _716_v245)
			_717_v246 = (_717_v246).Update(_715_v244, _716_v245)
		}
		var _718_v247 _dafny.Set
		_ = _718_v247
		_718_v247 = _dafny.SetOf(_382_v5, _382_v5)
		var _719_v249 _dafny.MultiSet
		_ = _719_v249
		_719_v249 = _dafny.MultiSetOf(_dafny.SetOf(_381_v4), _382_v5, _382_v5)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_388_v11), 0))
		_ = _index58
		(_388_v11).ArraySet1((func() _dafny.Set {
			var _coll24 = _dafny.NewBuilder()
			_ = _coll24
			for _iter25 := _dafny.Iterate((_718_v247).Elements()); ; {
				_compr_24, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _720_v248 _dafny.Set
				_720_v248 = interface{}(_compr_24).(_dafny.Set)
				if (_718_v247).Contains(_720_v248) {
					_coll24.Add(_720_v248)
				}
			}
			return _coll24.ToSet()
		}()).IsDisjointFrom(func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter26 := _dafny.Iterate((_719_v249).Elements()); ; {
				_compr_25, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _721_v250 _dafny.Set
				_721_v250 = interface{}(_compr_25).(_dafny.Set)
				if (_719_v249).Contains(_721_v250) {
					_coll25.Add(_721_v250)
				}
			}
			return _coll25.ToSet()
		}()), (_index58).Int())
		var _722_v251 _dafny.Set
		_ = _722_v251
		_722_v251 = _dafny.SetOf(_617_v174, _617_v174)
		var _723_v252 _dafny.Sequence
		_ = _723_v252
		_723_v252 = _dafny.SeqOf(_722_v251)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_388_v11), 0))
		_ = _index59
		(_388_v11).ArraySet1(!(!_dafny.Companion_Sequence_.Equal(_723_v252, _723_v252)), (_index59).Int())
		(_391_globalState).F26 = (_388_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_388_v11), 0))).Int()).(bool)
		var _724_v253 _dafny.Array
		_ = _724_v253
		var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw118
		_724_v253 = _nw118
		var _725_v254 _dafny.Array
		_ = _725_v254
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_14
		var _nw119 _dafny.Array
		_ = _nw119
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw119 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Map = (func(_726_v0 _dafny.Array, _727_v4 bool) func(_dafny.Int) _dafny.Map {
				return func(_728_i27 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_726_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_726_v0), 0))).Int()).(_dafny.Int)), _727_v4)
				}
			})(_377_v0, _381_v4)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw119 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw119).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw119).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_725_v254 = _nw119
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_724_v253), 0))
		_ = _index60
		(_724_v253).ArraySet1(_725_v254, (_index60).Int())
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_724_v253), 0))
		_ = _index61
		(_724_v253).ArraySet1(_725_v254, (_index61).Int())
	} else {
		_377_v0 = _377_v0
		var _729_v255 _dafny.Set
		_ = _729_v255
		_729_v255 = _dafny.SetOf(Companion_Default___.Fm26(_380_v3, _dafny.IntOfUint32((_386_v9).Cardinality()), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _391_globalState), _380_v3)
		var _730_v256 T2
		_ = _730_v256
		var _nw120 *C2 = New_C2_()
		_ = _nw120
		_nw120.Ctor__(_dafny.MultiSetOf(false), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))
		_730_v256 = _nw120
		var _731_v257 _dafny.MultiSet
		_ = _731_v257
		_731_v257 = _dafny.MultiSetOf(_730_v256, _730_v256)
		var _732_v258 T0
		_ = _732_v258
		var _nw121 *C4 = New_C4_()
		_ = _nw121
		_nw121.Ctor__(_383_v6, _602_v162)
		_732_v258 = _nw121
		var _nwElement0_26 _dafny.Int = Companion_Default___.SafeModuloInt((_729_v255).Cardinality(), Companion_Default___.Fm1(_391_globalState))
		_ = _nwElement0_26
		var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(19))
		_ = _nw122
		(_nw122).ArraySet1(_nwElement0_26, 0)
		(_nw122).ArraySet1(_383_v6, 1)
		(_nw122).ArraySet1(_383_v6, 2)
		(_nw122).ArraySet1(((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Minus(_383_v6), 3)
		(_nw122).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_386_v9, (Companion_Default___.SafeIndex((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32(), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Cardinality())).Plus(_383_v6), 4)
		(_nw122).ArraySet1(_dafny.IntOfInt64(802), 5)
		(_nw122).ArraySet1((func() _dafny.Int {
			if _381_v4 {
				return (_731_v257).Cardinality()
			}
			return _383_v6
		})(), 6)
		(_nw122).ArraySet1(((_603_v163).Fm4(_dafny.SetOf(_378_v1, _378_v1), _381_v4, _629_v185, _391_globalState)).Minus(_dafny.IntOfInt64(-311)), 7)
		(_nw122).ArraySet1(_383_v6, 8)
		(_nw122).ArraySet1(_383_v6, 9)
		(_nw122).ArraySet1(Companion_Default___.SafeModuloInt((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_629_v185).Cardinality())), 10)
		(_nw122).ArraySet1(_dafny.IntOfInt64(956), 11)
		(_nw122).ArraySet1((func() _dafny.Int {
			if _381_v4 {
				return (_730_v256).F32()
			}
			return (_386_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_629_v185).Cardinality()), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int)
		})(), 12)
		(_nw122).ArraySet1(((_601_v161).Cardinality()).Times((_730_v256).F32()), 13)
		(_nw122).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_383_v6, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))), 14)
		(_nw122).ArraySet1((_730_v256).F32(), 15)
		(_nw122).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_v258, func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter27 := _dafny.Iterate((_378_v1).Elements()); ; {
				_compr_26, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _733_v259 _dafny.CodePoint
				_733_v259 = interface{}(_compr_26).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_378_v1, _733_v259) {
					_coll26.Add(_733_v259)
				}
			}
			return _coll26.ToSet()
		}())).Cardinality(), 16)
		(_nw122).ArraySet1((_603_v163).Fm4(_390_v13, _381_v4, _629_v185, _391_globalState), 17)
		(_nw122).ArraySet1((_732_v258).F27(), 18)
		_377_v0 = _nw122
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_388_v11), 0))
		_ = _index62
		(_388_v11).ArraySet1((_381_v4) || (!(_381_v4)), (_index62).Int())
		var _734_v260 _dafny.Map
		_ = _734_v260
		_734_v260 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7)
		var _735_v263 _dafny.Array
		_ = _735_v263
		var _nwElement0_27 _dafny.Map = _734_v260
		_ = _nwElement0_27
		var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(29))
		_ = _nw123
		(_nw123).ArraySet1(_nwElement0_27, 0)
		(_nw123).ArraySet1(_734_v260, 1)
		(_nw123).ArraySet1(_734_v260, 2)
		(_nw123).ArraySet1(_734_v260, 3)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7), 4)
		(_nw123).ArraySet1((func() _dafny.Map {
			if _381_v4 {
				return _734_v260
			}
			return _734_v260
		})(), 5)
		(_nw123).ArraySet1(_734_v260, 6)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 7)
		(_nw123).ArraySet1(_734_v260, 8)
		(_nw123).ArraySet1(_734_v260, 9)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 10)
		(_nw123).ArraySet1((_734_v260).Update(_377_v0, _384_v7), 11)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 12)
		(_nw123).ArraySet1(_734_v260, 13)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, func() _dafny.Set {
			var _coll27 = _dafny.NewBuilder()
			_ = _coll27
			for _iter28 := _dafny.Iterate((func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter29 := _dafny.Iterate((_dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Elements()); ; {
					_compr_28, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _736_v261 _dafny.Int
					_736_v261 = interface{}(_compr_28).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)), _736_v261) {
						_coll28.Add((_736_v261).Times((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)), _381_v4)
					}
				}
				return _coll28.ToMap()
			}()).Keys().Elements()); ; {
				_compr_27, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _737_v262 _dafny.Int
				_737_v262 = interface{}(_compr_27).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate((_dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))).Elements()); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _736_v261 _dafny.Int
						_736_v261 = interface{}(_compr_29).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)), _736_v261) {
							_coll29.Add((_736_v261).Times((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)), _381_v4)
						}
					}
					return _coll29.ToMap()
				}()).Contains(_737_v262) {
					_coll27.Add((_737_v262).Minus(_dafny.IntOfInt64(458)))
				}
			}
			return _coll27.ToSet()
		}()), 14)
		(_nw123).ArraySet1(_734_v260, 15)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 16)
		(_nw123).ArraySet1((_734_v260).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7)), 17)
		(_nw123).ArraySet1(_734_v260, 18)
		(_nw123).ArraySet1(_734_v260, 19)
		(_nw123).ArraySet1(_734_v260, 20)
		(_nw123).ArraySet1(_734_v260, 21)
		(_nw123).ArraySet1(_734_v260, 22)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7), 23)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 24)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7), 25)
		(_nw123).ArraySet1((_734_v260).Merge(_734_v260), 26)
		(_nw123).ArraySet1(_734_v260, 27)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _384_v7), 28)
		_735_v263 = _nw123
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_735_v263), 0))
		_ = _index63
		(_735_v263).ArraySet1(_734_v260, (_index63).Int())
		var _738_v264 D2
		_ = _738_v264
		_738_v264 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus(_383_v6), (_732_v258).F27(), ((_730_v256).F32()).Times((_732_v258).F27()), _380_v3, _378_v1)
		var _739_v265 D1
		_ = _739_v265
		_739_v265 = Companion_D1_.Create_DC2_(_380_v3, (_730_v256).F32())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_388_v11), 0))
		_ = _index64
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_735_v263), 0))
		_ = _index65
		var _rhs69 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_732_v258).F27()), Companion_Default___.SafeModuloInt((_730_v256).F32(), (_730_v256).F32()))
		_ = _rhs69
		var _rhs70 bool = _381_v4
		_ = _rhs70
		var _rhs71 _dafny.Int = ((_739_v265).Dtor_cf3()).Times((_732_v258).F27())
		_ = _rhs71
		var _rhs72 _dafny.Map = (func() _dafny.Map {
			if (_381_v4) || (_381_v4) {
				return _734_v260
			}
			return _734_v260
		})()
		_ = _rhs72
		var _rhs73 D2 = _738_v264
		_ = _rhs73
		var _lhs46 *GlobalState = _391_globalState
		_ = _lhs46
		var _lhs47 _dafny.Array = _388_v11
		_ = _lhs47
		var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_388_v11), 0))
		_ = _lhs48
		var _lhs49 *GlobalState = _391_globalState
		_ = _lhs49
		var _lhs50 _dafny.Array = _735_v263
		_ = _lhs50
		var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_735_v263), 0))
		_ = _lhs51
		_lhs46.F15 = _rhs69
		(_lhs47).ArraySet1(_rhs70, (_lhs48).Int())
		_lhs49.F15 = _rhs71
		(_lhs50).ArraySet1(_rhs72, (_lhs51).Int())
		_738_v264 = _rhs73
		var _740_v266 *C2
		_ = _740_v266
		var _nw124 *C2 = New_C2_()
		_ = _nw124
		_nw124.Ctor__((_730_v256).F31(), _dafny.IntOfInt64(-102))
		_740_v266 = _nw124
		_740_v266 = _740_v266
		(_391_globalState).F26 = (func() bool {
			if _381_v4 {
				return true
			}
			return !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_741_v256 T2) func(_dafny.Int) _dafny.Int {
				return func(_742_i28 _dafny.Int) _dafny.Int {
					return (_741_v256).F32()
				}
			})(_730_v256))), (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))
		})()
	}
	var _743_i29 _dafny.Int
	_ = _743_i29
	_743_i29 = _dafny.Zero
	{
		for !(_381_v4) {
			{
				if (_743_i29).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_743_i29 = (_743_i29).Plus(_dafny.One)
				var _744_v267 D8
				_ = _744_v267
				_744_v267 = Companion_D8_.Create_DC16_((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), (_383_v6).Times(_383_v6))
				_744_v267 = _744_v267
				(_391_globalState).F0 = (_dafny.Zero).Minus((_dafny.IntOfUint32((_378_v1).Cardinality())).Minus(((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Plus((_382_v5).Cardinality())))
				if _381_v4 {
					(_391_globalState).F21 = Companion_Default___.SafeModuloInt(((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_386_v9).Cardinality())), (_620_v177).Cardinality())
					var _745_v268 _dafny.Sequence
					_ = _745_v268
					_745_v268 = _dafny.SeqOf(_386_v9, _386_v9)
					_599_v159 = (_599_v159).Update((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_745_v268).Cardinality()))
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))
					_ = _index66
					(_377_v0).ArraySet1((_386_v9).Select((Companion_Default___.SafeIndex((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int), (_index66).Int())
					var _746_v269 _dafny.Array
					_ = _746_v269
					var _len0_15 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_15
					var _nw125 _dafny.Array
					_ = _nw125
					if _len0_15.Cmp(_dafny.Zero) == 0 {
						_nw125 = _dafny.NewArray(_len0_15)
					} else {
						var _init15 func(_dafny.Int) _dafny.Map = (func(_747_v4 bool) func(_dafny.Int) _dafny.Map {
							return func(_748_i30 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_747_v4, _747_v4), _747_v4)
							}
						})(_381_v4)
						_ = _init15
						var _element0_15 = _init15(_dafny.Zero)
						_ = _element0_15
						_nw125 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
						(_nw125).ArraySet1(_element0_15, 0)
						var _nativeLen0_15 = (_len0_15).Int()
						_ = _nativeLen0_15
						for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
							(_nw125).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
						}
					}
					_746_v269 = _nw125
					var _749_v270 _dafny.Map
					_ = _749_v270
					_749_v270 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_381_v4, _381_v4), _381_v4)
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_746_v269), 0))
					_ = _index67
					(_746_v269).ArraySet1(_749_v270, (_index67).Int())
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_746_v269), 0))
					_ = _index68
					(_746_v269).ArraySet1(_749_v270, (_index68).Int())
					(_603_v163).M10(!(_381_v4) || (_381_v4), _386_v9, _391_globalState)
				} else {
					_381_v4 = _381_v4
					(_391_globalState).F26 = _381_v4
					var _750_v271 D12
					_ = _750_v271
					_750_v271 = Companion_D12_.Create_DC27_(!(_381_v4), (_620_v177).Update(_dafny.IntOfInt64(489), Companion_Default___.Abs(_383_v6)), _383_v6, _377_v0)
					var _751_v272 _dafny.Array
					_ = _751_v272
					var _nwElement0_28 D12 = _750_v271
					_ = _nwElement0_28
					var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(3))
					_ = _nw126
					(_nw126).ArraySet1(_nwElement0_28, 0)
					(_nw126).ArraySet1(_750_v271, 1)
					(_nw126).ArraySet1(_750_v271, 2)
					_751_v272 = _nw126
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_751_v272), 0))
					_ = _index69
					(_751_v272).ArraySet1(_750_v271, (_index69).Int())
					var _752_v274 _dafny.Map
					_ = _752_v274
					_752_v274 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v3, _381_v4)
					var _753_v275 _dafny.Sequence
					_ = _753_v275
					_753_v275 = _dafny.SeqOf(func() _dafny.Map {
						var _coll30 = _dafny.NewMapBuilder()
						_ = _coll30
						for _iter31 := _dafny.Iterate((_752_v274).Keys().Elements()); ; {
							_compr_30, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _754_v273 _dafny.CodePoint
							_754_v273 = interface{}(_compr_30).(_dafny.CodePoint)
							if (_752_v274).Contains(_754_v273) {
								_coll30.Add(_754_v273, (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int))
							}
						}
						return _coll30.ToMap()
					}())
					var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_751_v272), 0))
					_ = _index70
					var _rhs74 D12 = _750_v271
					_ = _rhs74
					var _rhs75 _dafny.Int = (_dafny.IntOfInt64(166)).Times(((_386_v9).Select((Companion_Default___.SafeIndex(_383_v6, _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)))
					_ = _rhs75
					var _rhs76 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfInt64(-560)).Plus(_dafny.IntOfUint32((_753_v275).Cardinality())))), _383_v6)
					_ = _rhs76
					var _rhs77 _dafny.Int = (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
					_ = _rhs77
					var _rhs78 bool = !(true) || (!(_381_v4))
					_ = _rhs78
					var _lhs52 _dafny.Array = _751_v272
					_ = _lhs52
					var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_751_v272), 0))
					_ = _lhs53
					var _lhs54 *GlobalState = _391_globalState
					_ = _lhs54
					var _lhs55 *GlobalState = _391_globalState
					_ = _lhs55
					var _lhs56 *GlobalState = _391_globalState
					_ = _lhs56
					var _lhs57 *GlobalState = _391_globalState
					_ = _lhs57
					(_lhs52).ArraySet1(_rhs74, (_lhs53).Int())
					_lhs54.F0 = _rhs75
					_lhs55.F0 = _rhs76
					_lhs56.F21 = _rhs77
					_lhs57.F26 = _rhs78
					var _755_v276 _dafny.Map
					_ = _755_v276
					_755_v276 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v6, _381_v4)
					_755_v276 = _755_v276
					var _rhs79 bool = _381_v4
					_ = _rhs79
					var _rhs80 bool = _381_v4
					_ = _rhs80
					var _rhs81 _dafny.Int = (_dafny.IntOfInt64(139)).Plus(_383_v6)
					_ = _rhs81
					var _rhs82 _dafny.Sequence = Companion_Default___.Fm41((_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int), _383_v6, _381_v4, _391_globalState)
					_ = _rhs82
					var _rhs83 _dafny.Sequence = _386_v9
					_ = _rhs83
					var _lhs58 *GlobalState = _391_globalState
					_ = _lhs58
					var _lhs59 *GlobalState = _391_globalState
					_ = _lhs59
					var _lhs60 *GlobalState = _391_globalState
					_ = _lhs60
					_lhs58.F26 = _rhs79
					_lhs59.F26 = _rhs80
					_lhs60.F0 = _rhs81
					_629_v185 = _rhs82
					_386_v9 = _rhs83
				}
				var _rhs84 _dafny.Int = (func() _dafny.Int {
					if (Companion_D6_.Create_DC12_(_381_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(721))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}((func(_756_v4 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_757_i31 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_756_v4, false)
						}
					})(_381_v4))), _378_v1)).Dtor_cf20() {
						return (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
					}
					return (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
				})()
				_ = _rhs84
				var _rhs85 _dafny.Int = (_377_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_377_v0), 0))).Int()).(_dafny.Int)
				_ = _rhs85
				var _rhs86 _dafny.Map = _616_v173
				_ = _rhs86
				var _lhs61 *GlobalState = _391_globalState
				_ = _lhs61
				var _lhs62 *GlobalState = _391_globalState
				_ = _lhs62
				_lhs61.F24 = _rhs84
				_lhs62.F18 = _rhs85
				_616_v173 = _rhs86
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	_dafny.Print((_377_v0).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v0).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_378_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_379_v2).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("yn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_380_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_381_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_382_v5).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_383_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_384_v7).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_386_v9, _dafny.SeqOf(_dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_387_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_387_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_387_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_388_v11).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_388_v11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_388_v11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v12).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_390_v13).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("yn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState.F8).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("yn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_391_globalState).F10()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_391_globalState).F10()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_391_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState.F13).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F17()).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("yn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F22()).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F23()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_391_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_391_globalState).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_391_globalState.F26)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_453_v61).Dtor_cf45()).Dtor_cf42()).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(491)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_454_v62).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(491)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_455_v63).Dtor_cf42()).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(491)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_599_v159).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(-491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_600_v160).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(491), _dafny.IntOfInt64(-491)), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_601_v161).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_602_v162).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_602_v162).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_602_v162).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_602_v162).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_602_v162).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_603_v163).F38()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(491), _dafny.IntOfInt64(-491)), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_615_v172).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_615_v172).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_616_v173).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_618_v175).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_619_v176).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_620_v177).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-491), _dafny.IntOfInt64(-491))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_629_v185, _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_743_i29)
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

func (CompanionStruct_D0_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
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

type D1_DC2 struct {
	Cf2 _dafny.CodePoint
	Cf3 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.CodePoint, Cf3 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2, Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Map
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Map) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf1() _dafny.Map {
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
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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

type D2_DC4 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.Int
	Cf7 _dafny.Int
	Cf8 _dafny.CodePoint
	Cf9 _dafny.Sequence
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Int, Cf6 _dafny.Int, Cf7 _dafny.Int, Cf8 _dafny.CodePoint, Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC4{Cf5, Cf6, Cf7, Cf8, Cf9}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC3 struct {
	Cf4 _dafny.Map
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf4 _dafny.Map) D2 {
	return D2{D2_DC3{Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC5 struct {
	Cf10 D2
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf10 D2) D2 {
	return D2{D2_DC5{Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.CodePoint('D'), _dafny.EmptySeq)
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.CodePoint {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf9
}

func (_this D2) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) Dtor_cf10() D2 {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + data.Cf9.VerbatimString(true) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9.Equals(data2.Cf9)
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D3_DC6 struct {
	Cf11 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf11 bool) D3 {
	return D3{D3_DC6{Cf11}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() bool {
	return false
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC6).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf11 == data2.Cf11
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

type D4_DC8 struct {
	Cf13 _dafny.CodePoint
	Cf14 _dafny.Int
	Cf15 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf13 _dafny.CodePoint, Cf14 _dafny.Int, Cf15 bool) D4 {
	return D4{D4_DC8{Cf13, Cf14, Cf15}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC9 struct {
	Cf16 _dafny.Int
	Cf17 bool
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf16 _dafny.Int, Cf17 bool) D4 {
	return D4{D4_DC9{Cf16, Cf17}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC7 struct {
	Cf12 _dafny.Array
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf12 _dafny.Array) D4 {
	return D4{D4_DC7{Cf12}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_(_dafny.CodePoint('D'), _dafny.Zero, false)
}

func (_this D4) Dtor_cf13() _dafny.CodePoint {
	return _this.Get_().(D4_DC8).Cf13
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC8).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC8).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC9).Cf17
}

func (_this D4) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D4_DC7).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf12 == data2.Cf12
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

type D5_DC10 struct {
	Cf18 _dafny.MultiSet
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf18 _dafny.MultiSet) D5 {
	return D5{D5_DC10{Cf18}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D5_DC10).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
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

type D6_DC12 struct {
	Cf20 bool
	Cf21 _dafny.Sequence
	Cf22 _dafny.Sequence
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf20 bool, Cf21 _dafny.Sequence, Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC12{Cf20, Cf21, Cf22}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC11 struct {
	Cf19 *C0
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf19 *C0) D6 {
	return D6{D6_DC11{Cf19}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC12_(false, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D6) Dtor_cf20() bool {
	return _this.Get_().(D6_DC12).Cf20
}

func (_this D6) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D6_DC12).Cf21
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC12).Cf22
}

func (_this D6) Dtor_cf19() *C0 {
	return _this.Get_().(D6_DC11).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + data.Cf22.VerbatimString(true) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22.Equals(data2.Cf22)
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf19 == data2.Cf19
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

type D7_DC14 struct {
	Cf24 _dafny.Int
	Cf25 T0
	Cf26 _dafny.Sequence
	Cf27 _dafny.MultiSet
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf24 _dafny.Int, Cf25 T0, Cf26 _dafny.Sequence, Cf27 _dafny.MultiSet) D7 {
	return D7{D7_DC14{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC13 struct {
	Cf23 _dafny.Sequence
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf23 _dafny.Sequence) D7 {
	return D7{D7_DC13{Cf23}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC14_(_dafny.Zero, (T0)(nil), _dafny.EmptySeq, _dafny.EmptyMultiSet)
}

func (_this D7) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf24
}

func (_this D7) Dtor_cf25() T0 {
	return _this.Get_().(D7_DC14).Cf25
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC14).Cf26
}

func (_this D7) Dtor_cf27() _dafny.MultiSet {
	return _this.Get_().(D7_DC14).Cf27
}

func (_this D7) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D7_DC13).Cf23
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && _dafny.AreEqual(data1.Cf25, data2.Cf25) && data1.Cf26.Equals(data2.Cf26) && data1.Cf27.Equals(data2.Cf27)
		}
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
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

type D8_DC16 struct {
	Cf29 _dafny.Int
	Cf30 _dafny.Int
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf29 _dafny.Int, Cf30 _dafny.Int) D8 {
	return D8{D8_DC16{Cf29, Cf30}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC15 struct {
	Cf28 _dafny.Map
}

func (D8_DC15) isD8() {}

func (CompanionStruct_D8_) Create_DC15_(Cf28 _dafny.Map) D8 {
	return D8{D8_DC15{Cf28}}
}

func (_this D8) Is_DC15() bool {
	_, ok := _this.Get_().(D8_DC15)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC16_(_dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D8_DC16).Cf29
}

func (_this D8) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D8_DC16).Cf30
}

func (_this D8) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D8_DC15).Cf28
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC15:
		{
			return "D8.DC15" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D8_DC15:
		{
			data2, ok := other.Get_().(D8_DC15)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D9_DC18 struct {
	Cf32 bool
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf32 bool) D9 {
	return D9{D9_DC18{Cf32}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

type D9_DC19 struct {
	Cf33 bool
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf33 bool) D9 {
	return D9{D9_DC19{Cf33}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC20 struct {
	Cf34 _dafny.Int
	Cf35 bool
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.Map
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf34 _dafny.Int, Cf35 bool, Cf36 bool, Cf37 bool, Cf38 _dafny.Map) D9 {
	return D9{D9_DC20{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC17 struct {
	Cf31 _dafny.Set
}

func (D9_DC17) isD9() {}

func (CompanionStruct_D9_) Create_DC17_(Cf31 _dafny.Set) D9 {
	return D9{D9_DC17{Cf31}}
}

func (_this D9) Is_DC17() bool {
	_, ok := _this.Get_().(D9_DC17)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC18_(false)
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC18).Cf32
}

func (_this D9) Dtor_cf33() bool {
	return _this.Get_().(D9_DC19).Cf33
}

func (_this D9) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf34
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC20).Cf35
}

func (_this D9) Dtor_cf36() bool {
	return _this.Get_().(D9_DC20).Cf36
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC20).Cf37
}

func (_this D9) Dtor_cf38() _dafny.Map {
	return _this.Get_().(D9_DC20).Cf38
}

func (_this D9) Dtor_cf31() _dafny.Set {
	return _this.Get_().(D9_DC17).Cf31
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC17:
		{
			return "D9.DC17" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf33 == data2.Cf33
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Equals(data2.Cf38)
		}
	case D9_DC17:
		{
			data2, ok := other.Get_().(D9_DC17)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D10_DC22 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.MultiSet
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf40 _dafny.Int, Cf41 _dafny.MultiSet) D10 {
	return D10{D10_DC22{Cf40, Cf41}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC21 struct {
	Cf39 _dafny.Sequence
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf39 _dafny.Sequence) D10 {
	return D10{D10_DC21{Cf39}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(_dafny.Zero, _dafny.EmptyMultiSet)
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf40
}

func (_this D10) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D10_DC22).Cf41
}

func (_this D10) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D10_DC21).Cf39
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC22:
		{
			data2, ok := other.Get_().(D10_DC22)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Equals(data2.Cf41)
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D11_DC24 struct {
	Cf43 _dafny.Int
	Cf44 D4
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf43 _dafny.Int, Cf44 D4) D11 {
	return D11{D11_DC24{Cf43, Cf44}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC23 struct {
	Cf42 _dafny.MultiSet
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf42 _dafny.MultiSet) D11 {
	return D11{D11_DC23{Cf42}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

type D11_DC25 struct {
	Cf45 D11
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf45 D11) D11 {
	return D11{D11_DC25{Cf45}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC24_(_dafny.Zero, Companion_D4_.Default())
}

func (_this D11) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D11_DC24).Cf43
}

func (_this D11) Dtor_cf44() D4 {
	return _this.Get_().(D11_DC24).Cf44
}

func (_this D11) Dtor_cf42() _dafny.MultiSet {
	return _this.Get_().(D11_DC23).Cf42
}

func (_this D11) Dtor_cf45() D11 {
	return _this.Get_().(D11_DC25).Cf45
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Equals(data2.Cf44)
		}
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
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

type D12_DC27 struct {
	Cf47 bool
	Cf48 _dafny.MultiSet
	Cf49 _dafny.Int
	Cf50 _dafny.Array
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf47 bool, Cf48 _dafny.MultiSet, Cf49 _dafny.Int, Cf50 _dafny.Array) D12 {
	return D12{D12_DC27{Cf47, Cf48, Cf49, Cf50}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

type D12_DC28 struct {
	Cf51 _dafny.Int
	Cf52 _dafny.Int
	Cf53 bool
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf51 _dafny.Int, Cf52 _dafny.Int, Cf53 bool) D12 {
	return D12{D12_DC28{Cf51, Cf52, Cf53}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC26 struct {
	Cf46 *C2
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf46 *C2) D12 {
	return D12{D12_DC26{Cf46}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

type D12_DC29 struct {
	Cf54 D12
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf54 D12) D12 {
	return D12{D12_DC29{Cf54}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC27_(false, _dafny.EmptyMultiSet, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D12) Dtor_cf47() bool {
	return _this.Get_().(D12_DC27).Cf47
}

func (_this D12) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D12_DC27).Cf48
}

func (_this D12) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf49
}

func (_this D12) Dtor_cf50() _dafny.Array {
	return _this.Get_().(D12_DC27).Cf50
}

func (_this D12) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf51
}

func (_this D12) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf52
}

func (_this D12) Dtor_cf53() bool {
	return _this.Get_().(D12_DC28).Cf53
}

func (_this D12) Dtor_cf46() *C2 {
	return _this.Get_().(D12_DC26).Cf46
}

func (_this D12) Dtor_cf54() D12 {
	return _this.Get_().(D12_DC29).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf46) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC27:
		{
			data2, ok := other.Get_().(D12_DC27)
			return ok && data1.Cf47 == data2.Cf47 && data1.Cf48.Equals(data2.Cf48) && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50 == data2.Cf50
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53 == data2.Cf53
		}
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf46 == data2.Cf46
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D13_DC31 struct {
	Cf56 _dafny.Int
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf56 _dafny.Int) D13 {
	return D13{D13_DC31{Cf56}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

type D13_DC32 struct {
	Cf57 _dafny.Int
	Cf58 _dafny.Int
	Cf59 _dafny.Sequence
	Cf60 bool
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf57 _dafny.Int, Cf58 _dafny.Int, Cf59 _dafny.Sequence, Cf60 bool) D13 {
	return D13{D13_DC32{Cf57, Cf58, Cf59, Cf60}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

type D13_DC30 struct {
	Cf55 *C1
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf55 *C1) D13 {
	return D13{D13_DC30{Cf55}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

type D13_DC33 struct {
	Cf61 D13
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf61 D13) D13 {
	return D13{D13_DC33{Cf61}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC31_(_dafny.Zero)
}

func (_this D13) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D13_DC31).Cf56
}

func (_this D13) Dtor_cf57() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf57
}

func (_this D13) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf58
}

func (_this D13) Dtor_cf59() _dafny.Sequence {
	return _this.Get_().(D13_DC32).Cf59
}

func (_this D13) Dtor_cf60() bool {
	return _this.Get_().(D13_DC32).Cf60
}

func (_this D13) Dtor_cf55() *C1 {
	return _this.Get_().(D13_DC30).Cf55
}

func (_this D13) Dtor_cf61() D13 {
	return _this.Get_().(D13_DC33).Cf61
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0
		}
	case D13_DC32:
		{
			data2, ok := other.Get_().(D13_DC32)
			return ok && data1.Cf57.Cmp(data2.Cf57) == 0 && data1.Cf58.Cmp(data2.Cf58) == 0 && data1.Cf59.Equals(data2.Cf59) && data1.Cf60 == data2.Cf60
		}
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf55 == data2.Cf55
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D14_DC34 struct {
	Cf62 _dafny.Set
}

func (D14_DC34) isD14() {}

func (CompanionStruct_D14_) Create_DC34_(Cf62 _dafny.Set) D14 {
	return D14{D14_DC34{Cf62}}
}

func (_this D14) Is_DC34() bool {
	_, ok := _this.Get_().(D14_DC34)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D14) Dtor_cf62() _dafny.Set {
	return _this.Get_().(D14_DC34).Cf62
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC34:
		{
			return "D14.DC34" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC34:
		{
			data2, ok := other.Get_().(D14_DC34)
			return ok && data1.Cf62.Equals(data2.Cf62)
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

type D15_DC36 struct {
	Cf64 _dafny.CodePoint
	Cf65 _dafny.Int
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf64 _dafny.CodePoint, Cf65 _dafny.Int) D15 {
	return D15{D15_DC36{Cf64, Cf65}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

type D15_DC37 struct {
	Cf66 bool
	Cf67 _dafny.Int
	Cf68 _dafny.Array
	Cf69 bool
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf66 bool, Cf67 _dafny.Int, Cf68 _dafny.Array, Cf69 bool) D15 {
	return D15{D15_DC37{Cf66, Cf67, Cf68, Cf69}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

type D15_DC38 struct {
	Cf70 _dafny.Int
	Cf71 bool
	Cf72 _dafny.Int
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf70 _dafny.Int, Cf71 bool, Cf72 _dafny.Int) D15 {
	return D15{D15_DC38{Cf70, Cf71, Cf72}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC35 struct {
	Cf63 _dafny.Sequence
}

func (D15_DC35) isD15() {}

func (CompanionStruct_D15_) Create_DC35_(Cf63 _dafny.Sequence) D15 {
	return D15{D15_DC35{Cf63}}
}

func (_this D15) Is_DC35() bool {
	_, ok := _this.Get_().(D15_DC35)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC36_(_dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D15) Dtor_cf64() _dafny.CodePoint {
	return _this.Get_().(D15_DC36).Cf64
}

func (_this D15) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D15_DC36).Cf65
}

func (_this D15) Dtor_cf66() bool {
	return _this.Get_().(D15_DC37).Cf66
}

func (_this D15) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D15_DC37).Cf67
}

func (_this D15) Dtor_cf68() _dafny.Array {
	return _this.Get_().(D15_DC37).Cf68
}

func (_this D15) Dtor_cf69() bool {
	return _this.Get_().(D15_DC37).Cf69
}

func (_this D15) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D15_DC38).Cf70
}

func (_this D15) Dtor_cf71() bool {
	return _this.Get_().(D15_DC38).Cf71
}

func (_this D15) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D15_DC38).Cf72
}

func (_this D15) Dtor_cf63() _dafny.Sequence {
	return _this.Get_().(D15_DC35).Cf63
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D15_DC35:
		{
			return "D15.DC35" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf64 == data2.Cf64 && data1.Cf65.Cmp(data2.Cf65) == 0
		}
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf66 == data2.Cf66 && data1.Cf67.Cmp(data2.Cf67) == 0 && data1.Cf68 == data2.Cf68 && data1.Cf69 == data2.Cf69
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71 == data2.Cf71 && data1.Cf72.Cmp(data2.Cf72) == 0
		}
	case D15_DC35:
		{
			data2, ok := other.Get_().(D15_DC35)
			return ok && data1.Cf63.Equals(data2.Cf63)
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
	Cf74 bool
	Cf75 _dafny.Int
	Cf76 _dafny.Int
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf74 bool, Cf75 _dafny.Int, Cf76 _dafny.Int) D16 {
	return D16{D16_DC40{Cf74, Cf75, Cf76}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC41 struct {
	Cf77 _dafny.Map
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf77 _dafny.Map) D16 {
	return D16{D16_DC41{Cf77}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC42 struct {
	Cf78 _dafny.Int
	Cf79 _dafny.Array
	Cf80 bool
	Cf81 D2
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf78 _dafny.Int, Cf79 _dafny.Array, Cf80 bool, Cf81 D2) D16 {
	return D16{D16_DC42{Cf78, Cf79, Cf80, Cf81}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC39 struct {
	Cf73 _dafny.Array
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf73 _dafny.Array) D16 {
	return D16{D16_DC39{Cf73}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC40_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf74() bool {
	return _this.Get_().(D16_DC40).Cf74
}

func (_this D16) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf75
}

func (_this D16) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf76
}

func (_this D16) Dtor_cf77() _dafny.Map {
	return _this.Get_().(D16_DC41).Cf77
}

func (_this D16) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf78
}

func (_this D16) Dtor_cf79() _dafny.Array {
	return _this.Get_().(D16_DC42).Cf79
}

func (_this D16) Dtor_cf80() bool {
	return _this.Get_().(D16_DC42).Cf80
}

func (_this D16) Dtor_cf81() D2 {
	return _this.Get_().(D16_DC42).Cf81
}

func (_this D16) Dtor_cf73() _dafny.Array {
	return _this.Get_().(D16_DC39).Cf73
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf77) + ")"
		}
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf73) + ")"
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
			return ok && data1.Cf74 == data2.Cf74 && data1.Cf75.Cmp(data2.Cf75) == 0 && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf77.Equals(data2.Cf77)
		}
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79 == data2.Cf79 && data1.Cf80 == data2.Cf80 && data1.Cf81.Equals(data2.Cf81)
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf73 == data2.Cf73
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

// Definition of datatype D17
type D17 struct {
	Data_D17_
}

func (_this D17) Get_() Data_D17_ {
	return _this.Data_D17_
}

type Data_D17_ interface {
	isD17()
}

type CompanionStruct_D17_ struct {
}

var Companion_D17_ = CompanionStruct_D17_{}

type D17_DC44 struct {
	Cf83 _dafny.Int
}

func (D17_DC44) isD17() {}

func (CompanionStruct_D17_) Create_DC44_(Cf83 _dafny.Int) D17 {
	return D17{D17_DC44{Cf83}}
}

func (_this D17) Is_DC44() bool {
	_, ok := _this.Get_().(D17_DC44)
	return ok
}

type D17_DC43 struct {
	Cf82 _dafny.Array
}

func (D17_DC43) isD17() {}

func (CompanionStruct_D17_) Create_DC43_(Cf82 _dafny.Array) D17 {
	return D17{D17_DC43{Cf82}}
}

func (_this D17) Is_DC43() bool {
	_, ok := _this.Get_().(D17_DC43)
	return ok
}

type D17_DC45 struct {
	Cf84 D17
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_(Cf84 D17) D17 {
	return D17{D17_DC45{Cf84}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC44_(_dafny.Zero)
}

func (_this D17) Dtor_cf83() _dafny.Int {
	return _this.Get_().(D17_DC44).Cf83
}

func (_this D17) Dtor_cf82() _dafny.Array {
	return _this.Get_().(D17_DC43).Cf82
}

func (_this D17) Dtor_cf84() D17 {
	return _this.Get_().(D17_DC45).Cf84
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC44:
		{
			return "D17.DC44" + "(" + _dafny.String(data.Cf83) + ")"
		}
	case D17_DC43:
		{
			return "D17.DC43" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D17_DC45:
		{
			return "D17.DC45" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC44:
		{
			data2, ok := other.Get_().(D17_DC44)
			return ok && data1.Cf83.Cmp(data2.Cf83) == 0
		}
	case D17_DC43:
		{
			data2, ok := other.Get_().(D17_DC43)
			return ok && data1.Cf82 == data2.Cf82
		}
	case D17_DC45:
		{
			data2, ok := other.Get_().(D17_DC45)
			return ok && data1.Cf84.Equals(data2.Cf84)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D17) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D17)
	return ok && _this.Equals(typed)
}

func Type_D17_() _dafny.TypeDescriptor {
	return type_D17_{}
}

type type_D17_ struct {
}

func (_this type_D17_) Default() interface{} {
	return Companion_D17_.Default()
}

func (_this type_D17_) String() string {
	return "main.D17"
}
func (_this D17) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D17{}

// End of datatype D17

// Definition of datatype D18
type D18 struct {
	Data_D18_
}

func (_this D18) Get_() Data_D18_ {
	return _this.Data_D18_
}

type Data_D18_ interface {
	isD18()
}

type CompanionStruct_D18_ struct {
}

var Companion_D18_ = CompanionStruct_D18_{}

type D18_DC47 struct {
	Cf86 _dafny.Int
	Cf87 _dafny.CodePoint
}

func (D18_DC47) isD18() {}

func (CompanionStruct_D18_) Create_DC47_(Cf86 _dafny.Int, Cf87 _dafny.CodePoint) D18 {
	return D18{D18_DC47{Cf86, Cf87}}
}

func (_this D18) Is_DC47() bool {
	_, ok := _this.Get_().(D18_DC47)
	return ok
}

type D18_DC46 struct {
	Cf85 _dafny.Set
}

func (D18_DC46) isD18() {}

func (CompanionStruct_D18_) Create_DC46_(Cf85 _dafny.Set) D18 {
	return D18{D18_DC46{Cf85}}
}

func (_this D18) Is_DC46() bool {
	_, ok := _this.Get_().(D18_DC46)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC47_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D18) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D18_DC47).Cf86
}

func (_this D18) Dtor_cf87() _dafny.CodePoint {
	return _this.Get_().(D18_DC47).Cf87
}

func (_this D18) Dtor_cf85() _dafny.Set {
	return _this.Get_().(D18_DC46).Cf85
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC47:
		{
			return "D18.DC47" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ")"
		}
	case D18_DC46:
		{
			return "D18.DC46" + "(" + _dafny.String(data.Cf85) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC47:
		{
			data2, ok := other.Get_().(D18_DC47)
			return ok && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87 == data2.Cf87
		}
	case D18_DC46:
		{
			data2, ok := other.Get_().(D18_DC46)
			return ok && data1.Cf85.Equals(data2.Cf85)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D18) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D18)
	return ok && _this.Equals(typed)
}

func Type_D18_() _dafny.TypeDescriptor {
	return type_D18_{}
}

type type_D18_ struct {
}

func (_this type_D18_) Default() interface{} {
	return Companion_D18_.Default()
}

func (_this type_D18_) String() string {
	return "main.D18"
}
func (_this D18) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D18{}

// End of datatype D18

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC49 struct {
	Cf89 bool
	Cf90 bool
	Cf91 _dafny.Int
}

func (D19_DC49) isD19() {}

func (CompanionStruct_D19_) Create_DC49_(Cf89 bool, Cf90 bool, Cf91 _dafny.Int) D19 {
	return D19{D19_DC49{Cf89, Cf90, Cf91}}
}

func (_this D19) Is_DC49() bool {
	_, ok := _this.Get_().(D19_DC49)
	return ok
}

type D19_DC48 struct {
	Cf88 _dafny.Set
}

func (D19_DC48) isD19() {}

func (CompanionStruct_D19_) Create_DC48_(Cf88 _dafny.Set) D19 {
	return D19{D19_DC48{Cf88}}
}

func (_this D19) Is_DC48() bool {
	_, ok := _this.Get_().(D19_DC48)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC49_(false, false, _dafny.Zero)
}

func (_this D19) Dtor_cf89() bool {
	return _this.Get_().(D19_DC49).Cf89
}

func (_this D19) Dtor_cf90() bool {
	return _this.Get_().(D19_DC49).Cf90
}

func (_this D19) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D19_DC49).Cf91
}

func (_this D19) Dtor_cf88() _dafny.Set {
	return _this.Get_().(D19_DC48).Cf88
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC49:
		{
			return "D19.DC49" + "(" + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D19_DC48:
		{
			return "D19.DC48" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC49:
		{
			data2, ok := other.Get_().(D19_DC49)
			return ok && data1.Cf89 == data2.Cf89 && data1.Cf90 == data2.Cf90 && data1.Cf91.Cmp(data2.Cf91) == 0
		}
	case D19_DC48:
		{
			data2, ok := other.Get_().(D19_DC48)
			return ok && data1.Cf88.Equals(data2.Cf88)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC51 struct {
}

func (D20_DC51) isD20() {}

func (CompanionStruct_D20_) Create_DC51_() D20 {
	return D20{D20_DC51{}}
}

func (_this D20) Is_DC51() bool {
	_, ok := _this.Get_().(D20_DC51)
	return ok
}

type D20_DC50 struct {
	Cf92 _dafny.MultiSet
}

func (D20_DC50) isD20() {}

func (CompanionStruct_D20_) Create_DC50_(Cf92 _dafny.MultiSet) D20 {
	return D20{D20_DC50{Cf92}}
}

func (_this D20) Is_DC50() bool {
	_, ok := _this.Get_().(D20_DC50)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC51_()
}

func (_this D20) Dtor_cf92() _dafny.MultiSet {
	return _this.Get_().(D20_DC50).Cf92
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC51:
		{
			return "D20.DC51"
		}
	case D20_DC50:
		{
			return "D20.DC50" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC51:
		{
			_, ok := other.Get_().(D20_DC51)
			return ok
		}
	case D20_DC50:
		{
			data2, ok := other.Get_().(D20_DC50)
			return ok && data1.Cf92.Equals(data2.Cf92)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC52 struct {
	Cf93 _dafny.Array
}

func (D21_DC52) isD21() {}

func (CompanionStruct_D21_) Create_DC52_(Cf93 _dafny.Array) D21 {
	return D21{D21_DC52{Cf93}}
}

func (_this D21) Is_DC52() bool {
	_, ok := _this.Get_().(D21_DC52)
	return ok
}

func (CompanionStruct_D21_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D21) Dtor_cf93() _dafny.Array {
	return _this.Get_().(D21_DC52).Cf93
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC52:
		{
			return "D21.DC52" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC52:
		{
			data2, ok := other.Get_().(D21_DC52)
			return ok && data1.Cf93 == data2.Cf93
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC54 struct {
	Cf95 _dafny.CodePoint
	Cf96 *C5
	Cf97 bool
	Cf98 _dafny.Int
}

func (D22_DC54) isD22() {}

func (CompanionStruct_D22_) Create_DC54_(Cf95 _dafny.CodePoint, Cf96 *C5, Cf97 bool, Cf98 _dafny.Int) D22 {
	return D22{D22_DC54{Cf95, Cf96, Cf97, Cf98}}
}

func (_this D22) Is_DC54() bool {
	_, ok := _this.Get_().(D22_DC54)
	return ok
}

type D22_DC55 struct {
	Cf99  _dafny.Sequence
	Cf100 *C4
	Cf101 _dafny.Sequence
	Cf102 _dafny.Sequence
	Cf103 _dafny.CodePoint
}

func (D22_DC55) isD22() {}

func (CompanionStruct_D22_) Create_DC55_(Cf99 _dafny.Sequence, Cf100 *C4, Cf101 _dafny.Sequence, Cf102 _dafny.Sequence, Cf103 _dafny.CodePoint) D22 {
	return D22{D22_DC55{Cf99, Cf100, Cf101, Cf102, Cf103}}
}

func (_this D22) Is_DC55() bool {
	_, ok := _this.Get_().(D22_DC55)
	return ok
}

type D22_DC53 struct {
	Cf94 _dafny.Sequence
}

func (D22_DC53) isD22() {}

func (CompanionStruct_D22_) Create_DC53_(Cf94 _dafny.Sequence) D22 {
	return D22{D22_DC53{Cf94}}
}

func (_this D22) Is_DC53() bool {
	_, ok := _this.Get_().(D22_DC53)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC54_(_dafny.CodePoint('D'), (*C5)(nil), false, _dafny.Zero)
}

func (_this D22) Dtor_cf95() _dafny.CodePoint {
	return _this.Get_().(D22_DC54).Cf95
}

func (_this D22) Dtor_cf96() *C5 {
	return _this.Get_().(D22_DC54).Cf96
}

func (_this D22) Dtor_cf97() bool {
	return _this.Get_().(D22_DC54).Cf97
}

func (_this D22) Dtor_cf98() _dafny.Int {
	return _this.Get_().(D22_DC54).Cf98
}

func (_this D22) Dtor_cf99() _dafny.Sequence {
	return _this.Get_().(D22_DC55).Cf99
}

func (_this D22) Dtor_cf100() *C4 {
	return _this.Get_().(D22_DC55).Cf100
}

func (_this D22) Dtor_cf101() _dafny.Sequence {
	return _this.Get_().(D22_DC55).Cf101
}

func (_this D22) Dtor_cf102() _dafny.Sequence {
	return _this.Get_().(D22_DC55).Cf102
}

func (_this D22) Dtor_cf103() _dafny.CodePoint {
	return _this.Get_().(D22_DC55).Cf103
}

func (_this D22) Dtor_cf94() _dafny.Sequence {
	return _this.Get_().(D22_DC53).Cf94
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC54:
		{
			return "D22.DC54" + "(" + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ", " + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ")"
		}
	case D22_DC55:
		{
			return "D22.DC55" + "(" + _dafny.String(data.Cf99) + ", " + _dafny.String(data.Cf100) + ", " + data.Cf101.VerbatimString(true) + ", " + data.Cf102.VerbatimString(true) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D22_DC53:
		{
			return "D22.DC53" + "(" + _dafny.String(data.Cf94) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC54:
		{
			data2, ok := other.Get_().(D22_DC54)
			return ok && data1.Cf95 == data2.Cf95 && data1.Cf96 == data2.Cf96 && data1.Cf97 == data2.Cf97 && data1.Cf98.Cmp(data2.Cf98) == 0
		}
	case D22_DC55:
		{
			data2, ok := other.Get_().(D22_DC55)
			return ok && data1.Cf99.Equals(data2.Cf99) && data1.Cf100 == data2.Cf100 && data1.Cf101.Equals(data2.Cf101) && data1.Cf102.Equals(data2.Cf102) && data1.Cf103 == data2.Cf103
		}
	case D22_DC53:
		{
			data2, ok := other.Get_().(D22_DC53)
			return ok && data1.Cf94.Equals(data2.Cf94)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC57 struct {
	Cf105 bool
	Cf106 _dafny.Array
	Cf107 _dafny.Int
	Cf108 _dafny.Int
}

func (D23_DC57) isD23() {}

func (CompanionStruct_D23_) Create_DC57_(Cf105 bool, Cf106 _dafny.Array, Cf107 _dafny.Int, Cf108 _dafny.Int) D23 {
	return D23{D23_DC57{Cf105, Cf106, Cf107, Cf108}}
}

func (_this D23) Is_DC57() bool {
	_, ok := _this.Get_().(D23_DC57)
	return ok
}

type D23_DC56 struct {
	Cf104 _dafny.Map
}

func (D23_DC56) isD23() {}

func (CompanionStruct_D23_) Create_DC56_(Cf104 _dafny.Map) D23 {
	return D23{D23_DC56{Cf104}}
}

func (_this D23) Is_DC56() bool {
	_, ok := _this.Get_().(D23_DC56)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC57_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, _dafny.Zero)
}

func (_this D23) Dtor_cf105() bool {
	return _this.Get_().(D23_DC57).Cf105
}

func (_this D23) Dtor_cf106() _dafny.Array {
	return _this.Get_().(D23_DC57).Cf106
}

func (_this D23) Dtor_cf107() _dafny.Int {
	return _this.Get_().(D23_DC57).Cf107
}

func (_this D23) Dtor_cf108() _dafny.Int {
	return _this.Get_().(D23_DC57).Cf108
}

func (_this D23) Dtor_cf104() _dafny.Map {
	return _this.Get_().(D23_DC56).Cf104
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC57:
		{
			return "D23.DC57" + "(" + _dafny.String(data.Cf105) + ", " + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ", " + _dafny.String(data.Cf108) + ")"
		}
	case D23_DC56:
		{
			return "D23.DC56" + "(" + _dafny.String(data.Cf104) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC57:
		{
			data2, ok := other.Get_().(D23_DC57)
			return ok && data1.Cf105 == data2.Cf105 && data1.Cf106 == data2.Cf106 && data1.Cf107.Cmp(data2.Cf107) == 0 && data1.Cf108.Cmp(data2.Cf108) == 0
		}
	case D23_DC56:
		{
			data2, ok := other.Get_().(D23_DC56)
			return ok && data1.Cf104.Equals(data2.Cf104)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of datatype D24
type D24 struct {
	Data_D24_
}

func (_this D24) Get_() Data_D24_ {
	return _this.Data_D24_
}

type Data_D24_ interface {
	isD24()
}

type CompanionStruct_D24_ struct {
}

var Companion_D24_ = CompanionStruct_D24_{}

type D24_DC59 struct {
}

func (D24_DC59) isD24() {}

func (CompanionStruct_D24_) Create_DC59_() D24 {
	return D24{D24_DC59{}}
}

func (_this D24) Is_DC59() bool {
	_, ok := _this.Get_().(D24_DC59)
	return ok
}

type D24_DC58 struct {
	Cf109 _dafny.Set
}

func (D24_DC58) isD24() {}

func (CompanionStruct_D24_) Create_DC58_(Cf109 _dafny.Set) D24 {
	return D24{D24_DC58{Cf109}}
}

func (_this D24) Is_DC58() bool {
	_, ok := _this.Get_().(D24_DC58)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC59_()
}

func (_this D24) Dtor_cf109() _dafny.Set {
	return _this.Get_().(D24_DC58).Cf109
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC59:
		{
			return "D24.DC59"
		}
	case D24_DC58:
		{
			return "D24.DC58" + "(" + _dafny.String(data.Cf109) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC59:
		{
			_, ok := other.Get_().(D24_DC59)
			return ok
		}
	case D24_DC58:
		{
			data2, ok := other.Get_().(D24_DC58)
			return ok && data1.Cf109.Equals(data2.Cf109)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D24) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D24)
	return ok && _this.Equals(typed)
}

func Type_D24_() _dafny.TypeDescriptor {
	return type_D24_{}
}

type type_D24_ struct {
}

func (_this type_D24_) Default() interface{} {
	return Companion_D24_.Default()
}

func (_this type_D24_) String() string {
	return "main.D24"
}
func (_this D24) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D24{}

// End of datatype D24

// Definition of trait T0
type T0 interface {
	String() string
	Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map
	F27() _dafny.Int
	F28() _dafny.Array
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
	Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map
	F27() _dafny.Int
	F28() _dafny.Array
	Fm3(p0 _dafny.MultiSet, globalState *GlobalState) bool
	Fm4(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int
	M1(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int)
	F29() _dafny.MultiSet
	F30() _dafny.Sequence
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

// Definition of trait T2
type T2 interface {
	String() string
	Fm5(globalState *GlobalState) bool
	Fm6(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence
	M2(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int)
	M3(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState)
	F31() _dafny.MultiSet
	F32() _dafny.Int
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F3   _dafny.Int
	F8   _dafny.MultiSet
	F13  _dafny.Map
	F15  _dafny.Int
	F18  _dafny.Int
	F21  _dafny.Int
	F24  _dafny.Int
	F26  bool
	_f1  _dafny.Sequence
	_f2  _dafny.Array
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f6  bool
	_f7  bool
	_f9  bool
	_f10 _dafny.Array
	_f11 _dafny.Int
	_f12 _dafny.Int
	_f14 bool
	_f16 bool
	_f17 _dafny.Set
	_f19 _dafny.Int
	_f20 bool
	_f22 _dafny.Map
	_f23 _dafny.Array
	_f25 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F3 = _dafny.Zero
	_this.F8 = _dafny.EmptyMultiSet
	_this.F13 = _dafny.EmptyMap
	_this.F15 = _dafny.Zero
	_this.F18 = _dafny.Zero
	_this.F21 = _dafny.Zero
	_this.F24 = _dafny.Zero
	_this.F26 = false
	_this._f1 = _dafny.EmptySeq
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f7 = false
	_this._f9 = false
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f14 = false
	_this._f16 = false
	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.Zero
	_this._f20 = false
	_this._f22 = _dafny.EmptyMap
	_this._f23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f25 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 _dafny.Array, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 bool, f8 _dafny.MultiSet, f9 bool, f10 _dafny.Array, f11 _dafny.Int, f12 _dafny.Int, f13 _dafny.Map, f14 bool, f15 _dafny.Int, f16 bool, f17 _dafny.Set, f18 _dafny.Int, f19 _dafny.Int, f20 bool, f21 _dafny.Int, f22 _dafny.Map, f23 _dafny.Array, f24 _dafny.Int, f25 _dafny.Int, f26 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this).F24 = f24
		(_this)._f25 = f25
		(_this).F26 = f26
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
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
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Array {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Set {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F19() _dafny.Int {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() bool {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F22() _dafny.Map {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Array {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F25() _dafny.Int {
	{
		return _this._f25
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f36 _dafny.Int
	_f37 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f36 = _dafny.Zero
	_this._f37 = _dafny.EmptySeq
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

func (_this *C0) Ctor__(f36 _dafny.Int, f37 _dafny.Sequence) {
	{
		(_this)._f36 = f36
		(_this)._f37 = f37
	}
}
func (_this *C0) F36() _dafny.Int {
	{
		return _this._f36
	}
}
func (_this *C0) F37() _dafny.Sequence {
	{
		return _this._f37
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f29 _dafny.MultiSet
	_f30 _dafny.Sequence
	_f27 _dafny.Int
	_f28 _dafny.Array
	F35  bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f29 = _dafny.EmptyMultiSet
	_this._f30 = _dafny.EmptySeq
	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F35 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F29() _dafny.MultiSet {
	return _this._f29
}
func (_this *C1) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C1) F27() _dafny.Int {
	return _this._f27
}
func (_this *C1) F28() _dafny.Array {
	return _this._f28
}
func (_this *C1) Ctor__(f35 bool, f29 _dafny.MultiSet, f30 _dafny.Sequence, f27 _dafny.Int, f28 _dafny.Array) {
	{
		(_this).F35 = f35
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C1) Fm3(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}(func(_758_i0 _dafny.Int) _dafny.Int {
			return (_this).F27()
		})), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('l'))).Cardinality())), _dafny.IntOfUint32(((_this).F30()).Cardinality()))))
	}
}
func (_this *C1) Fm4(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27())
	}
}
func (_this *C1) Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F27())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tcfvt")).Cardinality()), _this.F35)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("istggpsd")).Cardinality())))
	}
}
func (_this *C1) Fm11(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F35
	}
}
func (_this *C1) M1(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _759_v0 _dafny.Array
		_ = _759_v0
		var _nw127 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw127
		_759_v0 = _nw127
		_759_v0 = _759_v0
		var _760_v1 _dafny.Sequence
		_ = _760_v1
		_760_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ltinafpk")
		var _761_v2 _dafny.CodePoint
		_ = _761_v2
		_761_v2 = _dafny.CodePoint('j')
		var _762_v3 _dafny.Set
		_ = _762_v3
		_762_v3 = _dafny.SetOf(_760_v1, _dafny.SeqOf(_761_v2), _760_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}((func(_763_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_764_i0 _dafny.Int) _dafny.CodePoint {
				return _763_v2
			}
		})(_761_v2))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F27()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}((func(_765_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_766_i0 _dafny.Int) _dafny.CodePoint {
				return _765_v2
			}
		})(_761_v2)))).Cardinality()))).Uint32(), _761_v2), _760_v1)
		var _767_v4 _dafny.Map
		_ = _767_v4
		_767_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _this.F35)
		var _768_v5 _dafny.Sequence
		_ = _768_v5
		_768_v5 = _dafny.SeqOf(_this.F35)
		var _769_v6 _dafny.Sequence
		_ = _769_v6
		_769_v6 = _dafny.SeqOf(_this.F35, _this.F35, (func() bool {
			if (_767_v4).Contains(_this.F35) {
				return (_767_v4).Get(_this.F35).(bool)
			}
			return _this.F35
		})(), _this.F35, (_768_v5).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_768_v5).Cardinality()))).Uint32()).(bool))
		var _rhs87 _dafny.Sequence = _760_v1
		_ = _rhs87
		var _rhs88 bool = !(((_this).F27()).Cmp((_this).F27()) <= 0)
		_ = _rhs88
		var _rhs89 bool = _this.F35
		_ = _rhs89
		var _rhs90 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_760_v1).Cardinality()), (_this).F27())
		_ = _rhs90
		var _rhs91 bool = ((_this).Fm4(_762_v3, _this.F35, _768_v5, globalState)).Cmp((_this).F27()) > 0
		_ = _rhs91
		var _lhs63 *GlobalState = globalState
		_ = _lhs63
		var _lhs64 *C1 = _this
		_ = _lhs64
		var _lhs65 *GlobalState = globalState
		_ = _lhs65
		_760_v1 = _rhs87
		_lhs63.F26 = _rhs88
		_lhs64.F35 = _rhs89
		r0 = _rhs90
		_lhs65.F26 = _rhs91
		var _770_v7 _dafny.Array
		_ = _770_v7
		var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
		_ = _nw128
		_770_v7 = _nw128
		var _771_v8 _dafny.Set
		_ = _771_v8
		_771_v8 = _dafny.SetOf(_759_v0, _759_v0)
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))
		_ = _index71
		(_759_v0).ArraySet1((_dafny.SetOf(_759_v0)).IsDisjointFrom(_771_v8), (_index71).Int())
		var _772_v9 _dafny.Array
		_ = _772_v9
		var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw129
		_772_v9 = _nw129
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))
		_ = _index72
		(_772_v9).ArraySet1((_this).F27(), (_index72).Int())
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))
		_ = _index73
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))
		_ = _index74
		var _rhs92 _dafny.Array = _770_v7
		_ = _rhs92
		var _rhs93 bool = (func() bool {
			if _this.F35 {
				return _this.F35
			}
			return (func() bool {
				if _this.F35 {
					return _this.F35
				}
				return _this.F35
			})()
		})()
		_ = _rhs93
		var _rhs94 _dafny.CodePoint = _761_v2
		_ = _rhs94
		var _rhs95 _dafny.Int = _dafny.IntOfInt64(-114)
		_ = _rhs95
		var _rhs96 _dafny.Int = (_this).F27()
		_ = _rhs96
		var _lhs66 _dafny.Array = _759_v0
		_ = _lhs66
		var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))
		_ = _lhs67
		var _lhs68 _dafny.Array = _772_v9
		_ = _lhs68
		var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))
		_ = _lhs69
		_770_v7 = _rhs92
		(_lhs66).ArraySet1(_rhs93, (_lhs67).Int())
		_761_v2 = _rhs94
		(_lhs68).ArraySet1(_rhs95, (_lhs69).Int())
		r2 = _rhs96
		var _773_v10 _dafny.Set
		_ = _773_v10
		_773_v10 = _dafny.SetOf((_this).F27())
		var _774_v11 _dafny.Map
		_ = _774_v11
		_774_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), _773_v10)
		var _775_v12 _dafny.Map
		_ = _775_v12
		_775_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_759_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))).Int()).(bool), _dafny.IntOfInt64(145))
		_774_v11 = (_774_v11).Update(((_this).F27()).Plus((_775_v12).Cardinality()), _dafny.SetOf((_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), (_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), (_this).F27()))
		var _776_v13 _dafny.Map
		_ = _776_v13
		_776_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _760_v1)
		var _777_v14 bool
		_ = _777_v14
		_777_v14 = (_759_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))).Int()).(bool)
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))
		_ = _index75
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))
		_ = _index76
		var _rhs97 _dafny.Int = (_this).F27()
		_ = _rhs97
		var _rhs98 _dafny.CodePoint = Companion_Default___.Fm12(Companion_Default___.SafeDivisionInt((_this).F27(), (_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int)), (_759_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))).Int()).(bool), globalState)
		_ = _rhs98
		var _rhs99 _dafny.Map = _776_v13
		_ = _rhs99
		var _rhs100 bool = _dafny.Companion_Sequence_.Equal(_760_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}((func(_778_v1 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
			return func(_779_i1 _dafny.Int) _dafny.CodePoint {
				return (_778_v1).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_778_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_760_v1))), _760_v1))
		_ = _rhs100
		var _rhs101 bool = (_777_v14)
		_ = _rhs101
		var _lhs70 _dafny.Array = _772_v9
		_ = _lhs70
		var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))
		_ = _lhs71
		var _lhs72 *C1 = _this
		_ = _lhs72
		var _lhs73 _dafny.Array = _759_v0
		_ = _lhs73
		var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_759_v0), 0))
		_ = _lhs74
		(_lhs70).ArraySet1(_rhs97, (_lhs71).Int())
		_761_v2 = _rhs98
		_776_v13 = _rhs99
		_lhs72.F35 = _rhs100
		(_lhs73).ArraySet1(_rhs101, (_lhs74).Int())
		var _780_v15 *C0
		_ = _780_v15
		var _nw130 *C0 = New_C0_()
		_ = _nw130
		_nw130.Ctor__((_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), _760_v1)
		_780_v15 = _nw130
		r0 = ((_dafny.Zero).Minus(((_this).F30()).Select((Companion_Default___.SafeIndex((_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int))).Times(_dafny.IntOfUint32((_760_v1).Cardinality()))
		r1 = (_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int)
		r2 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_772_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_772_v9), 0))).Int()).(_dafny.Int), (_780_v15).F36()))
		return r0, r1, r2
	}
}
func (_this *C1) M5(p0 bool, p1 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _781_v0 _dafny.Sequence
		_ = _781_v0
		_781_v0 = _dafny.SeqOf(p1, p0)
		var _782_v1 _dafny.Set
		_ = _782_v1
		_782_v1 = _dafny.SetOf((_this).F27(), (_this).F27())
		(globalState).F21 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_781_v0, _781_v0)).Cardinality()), (_782_v1).Cardinality())
		var _783_i0 _dafny.Int
		_ = _783_i0
		_783_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_783_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_783_i0 = (_783_i0).Plus(_dafny.One)
					if true {
						r0 = !(p0)
						var _784_v2 _dafny.Sequence
						_ = _784_v2
						_784_v2 = _dafny.UnicodeSeqOfUtf8Bytes("pbfgu")
						_784_v2 = _784_v2
						var _785_v3 _dafny.Map
						_ = _785_v3
						_785_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, p1)
						var _786_v4 _dafny.Sequence
						_ = _786_v4
						_786_v4 = _784_v2
						_785_v3 = (_785_v3).Update(Companion_Default___.Fm0((_this).F27(), globalState), !((_this).Fm11(p1, (_this).F27(), _786_v4, (_this).F27(), globalState)))
						(globalState).F26 = (func() bool {
							if !(_dafny.Companion_Sequence_.IsPrefixOf((_this).F30(), (_this).F30())) {
								return _this.F35
							}
							return _this.F35
						})()
						(globalState).F24 = ((_this).F27()).Times((_this).F27())
					} else {
						var _787_v5 _dafny.Sequence
						_ = _787_v5
						_787_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ddutqmb")
						var _788_v6 _dafny.CodePoint
						_ = _788_v6
						_788_v6 = _dafny.CodePoint('m')
						var _789_v7 *C0
						_ = _789_v7
						var _nw131 *C0 = New_C0_()
						_ = _nw131
						_nw131.Ctor__(_dafny.IntOfInt64(-849), _dafny.Companion_Sequence_.Update(_787_v5, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_787_v5).Cardinality()))).Uint32(), _788_v6))
						_789_v7 = _nw131
						(globalState).F3 = ((_this).F27()).Minus(Companion_Default___.Fm1(globalState))
						_787_v5 = _dafny.UnicodeSeqOfUtf8Bytes("o")
						var _790_v8 *C0
						_ = _790_v8
						var _nw132 *C0 = New_C0_()
						_ = _nw132
						_nw132.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(735), (_dafny.Zero).Minus((_this).F27())), (_789_v7).F37())
						_790_v8 = _nw132
						_781_v0 = Companion_Default___.Fm13((_this).F27(), (func() _dafny.Int {
							if p0 {
								return (_dafny.Zero).Minus((_790_v8).F36())
							}
							return (_789_v7).F36()
						})(), p0, _dafny.IntOfInt64(-249), globalState)
					}
					var _791_v9 *C0
					_ = _791_v9
					var _nw133 *C0 = New_C0_()
					_ = _nw133
					_nw133.Ctor__((_this).F27(), _dafny.UnicodeSeqOfUtf8Bytes("fyxj"))
					_791_v9 = _nw133
					var _792_v10 _dafny.Map
					_ = _792_v10
					_792_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F27())
					(globalState).F21 = Companion_Default___.SafeModuloInt((_this).F27(), (_dafny.IntOfInt64(301)).Plus((func() _dafny.Int {
						if (_792_v10).Contains(p0) {
							return (_792_v10).Get(p0).(_dafny.Int)
						}
						return (_this).F27()
					})()))
					var _793_v11 _dafny.Sequence
					_ = _793_v11
					_793_v11 = _dafny.UnicodeSeqOfUtf8Bytes("yxypbryxn")
					var _794_v12 _dafny.Array
					_ = _794_v12
					var _nwElement0_29 _dafny.Int = ((_791_v9).F36()).Minus((_this).F27())
					_ = _nwElement0_29
					var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(6))
					_ = _nw134
					(_nw134).ArraySet1(_nwElement0_29, 0)
					(_nw134).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(373))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg59 _dafny.Int) interface{} {
							return coer59(arg59)
						}
					}((func(_795_v9 *C0) func(_dafny.Int) _dafny.Int {
						return func(_796_i1 _dafny.Int) _dafny.Int {
							return (_795_v9).F36()
						}
					})(_791_v9))), (_this).F30())).Cardinality()), 1)
					(_nw134).ArraySet1(_dafny.IntOfUint32((_781_v0).Cardinality()), 2)
					(_nw134).ArraySet1(_dafny.IntOfInt64(-645), 3)
					(_nw134).ArraySet1((_this).F27(), 4)
					(_nw134).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F27(), (_791_v9).F36())), 5)
					_794_v12 = _nw134
					var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_794_v12), 0))
					_ = _index77
					(_794_v12).ArraySet1((_this).F27(), (_index77).Int())
					var _797_v13 _dafny.Sequence
					_ = _797_v13
					_797_v13 = _dafny.SeqOf(_793_v11)
					var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_794_v12), 0))
					_ = _index78
					var _rhs102 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_791_v9).F37(), _dafny.Companion_Sequence_.Concatenate((_791_v9).F37(), Companion_Default___.Fm14(_this.F35, p0, (_781_v0).Select((Companion_Default___.SafeIndex((_791_v9).F36(), _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32()).(bool), (_this).F27(), globalState)))
					_ = _rhs102
					var _rhs103 _dafny.Int = (_this).Fm4((func() _dafny.Set {
						var _coll31 = _dafny.NewBuilder()
						_ = _coll31
						for _iter32 := _dafny.Iterate((_797_v13).Elements()); ; {
							_compr_31, _ok32 := _iter32()
							if !_ok32 {
								break
							}
							var _798_v14 _dafny.Sequence
							_798_v14 = interface{}(_compr_31).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_797_v13, _798_v14) {
								_coll31.Add(_798_v14)
							}
						}
						return _coll31.ToSet()
					}()).Intersection(func() _dafny.Set {
						var _coll32 = _dafny.NewBuilder()
						_ = _coll32
						for _iter33 := _dafny.Iterate((_797_v13).Elements()); ; {
							_compr_32, _ok33 := _iter33()
							if !_ok33 {
								break
							}
							var _799_v15 _dafny.Sequence
							_799_v15 = interface{}(_compr_32).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_797_v13, _799_v15) {
								_coll32.Add(_799_v15)
							}
						}
						return _coll32.ToSet()
					}()), false, _781_v0, globalState)
					_ = _rhs103
					var _rhs104 _dafny.Int = (_this).F27()
					_ = _rhs104
					var _rhs105 bool = !(true) || (!(Companion_Default___.Fm15((func() _dafny.Set {
						var _coll33 = _dafny.NewBuilder()
						_ = _coll33
						for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(993), _dafny.IntOfInt64(216))); ; {
							_compr_33, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _800_v16 _dafny.Int
							_800_v16 = interface{}(_compr_33).(_dafny.Int)
							if ((_dafny.IntOfInt64(993)).Cmp(_800_v16) <= 0) && ((_800_v16).Cmp(_dafny.IntOfInt64(216)) < 0) {
								_coll33.Add((_800_v16).Minus((_791_v9).F36()))
							}
						}
						return _coll33.ToSet()
					}()).Cardinality(), (_791_v9).F36(), globalState)).Contains((_this).F27()))
					_ = _rhs105
					var _rhs106 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm16((_this).F27(), globalState)).Cardinality())
					_ = _rhs106
					var _lhs75 _dafny.Array = _794_v12
					_ = _lhs75
					var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_794_v12), 0))
					_ = _lhs76
					var _lhs77 *GlobalState = globalState
					_ = _lhs77
					var _lhs78 *C1 = _this
					_ = _lhs78
					var _lhs79 *GlobalState = globalState
					_ = _lhs79
					_793_v11 = _rhs102
					(_lhs75).ArraySet1(_rhs103, (_lhs76).Int())
					_lhs77.F18 = _rhs104
					_lhs78.F35 = _rhs105
					_lhs79.F0 = _rhs106
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _801_v17 _dafny.Sequence
		_ = _801_v17
		_801_v17 = _dafny.UnicodeSeqOfUtf8Bytes("elhuc")
		var _802_v18 _dafny.Sequence
		_ = _802_v18
		_802_v18 = _801_v17
		var _803_v19 _dafny.Map
		_ = _803_v19
		_803_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfUint32(((_this).F30()).Cardinality()), globalState), (_this).Fm11(p1, (_this).F27(), _802_v18, (_this).F27(), globalState))
		(globalState).F18 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_782_v1).Cardinality(), Companion_Default___.SafeModuloInt((_803_v19).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p0)).Cardinality())))
		var _hi3 _dafny.Int = (_this).F27()
		_ = _hi3
		for _804_i2 := _dafny.IntOfInt64(36); _804_i2.Cmp(_hi3) < 0; _804_i2 = _804_i2.Plus(_dafny.One) {
			if p1 {
				var _805_v20 _dafny.Array
				_ = _805_v20
				var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_801_v17, _801_v17)
				_ = _nwElement0_30
				var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(4))
				_ = _nw135
				(_nw135).ArraySet1(_nwElement0_30, 0)
				(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_806_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				})), 1)
				(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(743))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_807_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})), 2)
				(_nw135).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("teeic"), 3)
				_805_v20 = _nw135
				_805_v20 = (_this).F28()
				var _808_v21 _dafny.Array
				_ = _808_v21
				var _nw136 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw136
				_808_v21 = _nw136
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_808_v21), 0))
				_ = _index79
				(_808_v21).ArraySet1(p0, (_index79).Int())
				var _809_v22 D4
				_ = _809_v22
				_809_v22 = Companion_D4_.Create_DC7_(_808_v21)
				var _810_v23 _dafny.Sequence
				_ = _810_v23
				_810_v23 = _dafny.SeqOf(_808_v21, (_809_v22).Dtor_cf12(), _808_v21)
				var _811_v24 _dafny.CodePoint
				_ = _811_v24
				_811_v24 = _dafny.CodePoint('t')
				var _812_v25 _dafny.Array
				_ = _812_v25
				var _nwElement0_31 _dafny.Int = (func() _dafny.Int {
					if _this.F35 {
						return _804_i2
					}
					return _dafny.IntOfUint32((Companion_Default___.Fm14(p1, p1, p0, (_dafny.Zero).Minus((_this).F27()), globalState)).Cardinality())
				})()
				_ = _nwElement0_31
				var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(25))
				_ = _nw137
				(_nw137).ArraySet1(_nwElement0_31, 0)
				(_nw137).ArraySet1((_dafny.Zero).Minus((_this).F27()), 1)
				(_nw137).ArraySet1((_this).F27(), 2)
				(_nw137).ArraySet1(((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), 3)
				(_nw137).ArraySet1(_804_i2, 4)
				(_nw137).ArraySet1(_dafny.IntOfUint32((_781_v0).Cardinality()), 5)
				(_nw137).ArraySet1((_dafny.SetOf(_804_i2, _804_i2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(351))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_813_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))).Cardinality()))).Cardinality(), 6)
				(_nw137).ArraySet1(_804_i2, 7)
				(_nw137).ArraySet1(_804_i2, 8)
				(_nw137).ArraySet1((_this).F27(), 9)
				(_nw137).ArraySet1(_dafny.IntOfInt64(-750), 10)
				(_nw137).ArraySet1((_this).F27(), 11)
				(_nw137).ArraySet1((((_this).F29()).Update(false, Companion_Default___.Abs(_804_i2))).Cardinality(), 12)
				(_nw137).ArraySet1((Companion_Default___.Fm17(p1, _this.F35, p0, globalState)).Cardinality(), 13)
				(_nw137).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F27()), p0)).Cardinality(), 14)
				(_nw137).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_810_v23, _810_v23)).Cardinality()), 15)
				(_nw137).ArraySet1((func() _dafny.Int {
					if _this.F35 {
						return _dafny.IntOfInt64(180)
					}
					return _804_i2
				})(), 16)
				(_nw137).ArraySet1(_804_i2, 17)
				(_nw137).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_801_v17, _dafny.UnicodeSeqOfUtf8Bytes("lylquqbbn")), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_814_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-837))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}(func(_815_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))).Cardinality())
				}))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_801_v17, _dafny.UnicodeSeqOfUtf8Bytes("lylquqbbn"))).Cardinality()))).Uint32(), _811_v24)).Cardinality())), 18)
				(_nw137).ArraySet1(((_this).F27()).Minus(_dafny.IntOfUint32(((_this).F30()).Cardinality())), 19)
				(_nw137).ArraySet1((_dafny.Zero).Minus((_this).F27()), 20)
				(_nw137).ArraySet1(((_this).F29()).Cardinality(), 21)
				(_nw137).ArraySet1(_804_i2, 22)
				(_nw137).ArraySet1(_dafny.IntOfInt64(669), 23)
				(_nw137).ArraySet1((_this).F27(), 24)
				_812_v25 = _nw137
				var _816_v26 _dafny.Map
				_ = _816_v26
				_816_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_781_v0, _781_v0), _this.F35)
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_808_v21), 0))
				_ = _index80
				var _rhs107 _dafny.Int = (_816_v26).Cardinality()
				_ = _rhs107
				var _rhs108 bool = p1
				_ = _rhs108
				var _rhs109 _dafny.Array = _812_v25
				_ = _rhs109
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 _dafny.Array = _808_v21
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_808_v21), 0))
				_ = _lhs82
				_lhs80.F0 = _rhs107
				(_lhs81).ArraySet1(_rhs108, (_lhs82).Int())
				_812_v25 = _rhs109
				_811_v24 = _811_v24
				var _817_v27 bool
				_ = _817_v27
				var _818_v28 _dafny.Int
				_ = _818_v28
				var _out31 bool
				_ = _out31
				var _out32 _dafny.Int
				_ = _out32
				_out31, _out32 = Companion_Default___.M0(!(!(!(true))), _dafny.IntOfInt64(470), (_this).F27(), _dafny.IntOfInt64(253), globalState)
				_817_v27 = _out31
				_818_v28 = _out32
				var _819_v29 _dafny.Map
				_ = _819_v29
				_819_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if p0 {
						return _817_v27
					}
					return p1
				})(), _802_v18)
				_819_v29 = (_819_v29).Update(_this.F35, _802_v18)
			} else {
				r0 = p1
				var _820_v30 _dafny.CodePoint
				_ = _820_v30
				_820_v30 = _dafny.CodePoint('w')
				var _821_v31 _dafny.MultiSet
				_ = _821_v31
				_821_v31 = _dafny.MultiSetOf((_this).F27(), (_dafny.Zero).Minus(_804_i2), (_this).F27())
				var _822_v32 _dafny.Map
				_ = _822_v32
				_822_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_821_v31, _804_i2)
				var _rhs110 _dafny.CodePoint = _dafny.CodePoint('e')
				_ = _rhs110
				var _rhs111 bool = _dafny.Companion_Sequence_.IsPrefixOf(_801_v17, _801_v17)
				_ = _rhs111
				var _rhs112 bool = (func() bool {
					if (_781_v0).Select((Companion_Default___.SafeIndex(_804_i2, _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32()).(bool) {
						return !(((_dafny.Zero).Minus((Companion_D1_.Create_DC2_(_820_v30, _dafny.IntOfInt64(168))).Dtor_cf3())).Cmp((_this).F27()) == 0)
					}
					return p0
				})()
				_ = _rhs112
				var _rhs113 _dafny.Int = (func() _dafny.Int {
					if _dafny.Companion_Sequence_.IsProperPrefixOf(_801_v17, _801_v17) {
						return _804_i2
					}
					return (_this).F27()
				})()
				_ = _rhs113
				var _rhs114 _dafny.Map = _822_v32
				_ = _rhs114
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				_820_v30 = _rhs110
				_lhs83.F26 = _rhs111
				_lhs84.F26 = _rhs112
				_lhs85.F18 = _rhs113
				_822_v32 = _rhs114
				(globalState).F26 = (p1) || (p0)
				(globalState).F21 = Companion_Default___.SafeModuloInt((_this).F27(), (func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(750), _dafny.IntOfInt64(86))); ; {
						_compr_34, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _823_v33 _dafny.Int
						_823_v33 = interface{}(_compr_34).(_dafny.Int)
						if ((_dafny.IntOfInt64(750)).Cmp(_823_v33) <= 0) && ((_823_v33).Cmp(_dafny.IntOfInt64(86)) < 0) {
							_coll34.Add((_823_v33).Minus(_804_i2), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _804_i2))
						}
					}
					return _coll34.ToMap()
				}()).Cardinality())
				(globalState).F15 = (_dafny.IntOfUint32((_801_v17).Cardinality())).Plus((_804_i2).Times(_dafny.IntOfInt64(-170)))
			}
			var _824_v34 _dafny.CodePoint
			_ = _824_v34
			_824_v34 = _dafny.CodePoint('c')
			var _825_v35 _dafny.Map
			_ = _825_v35
			_825_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v34, p1)
			var _826_v36 _dafny.Array
			_ = _826_v36
			var _nwElement0_32 bool = !(p0)
			_ = _nwElement0_32
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(4))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_32, 0)
			(_nw138).ArraySet1((func() bool {
				if (_825_v35).Contains(_824_v34) {
					return (_825_v35).Get(_824_v34).(bool)
				}
				return !(_this.F35)
			})(), 1)
			(_nw138).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihcoyjmp")).Cardinality())).Cmp((_this).F27()) <= 0, 2)
			(_nw138).ArraySet1(false, 3)
			_826_v36 = _nw138
			r1 = _826_v36
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_826_v36), 0))
			_ = _index81
			(_826_v36).ArraySet1(_this.F35, (_index81).Int())
			var _827_v37 _dafny.MultiSet
			_ = _827_v37
			_827_v37 = _dafny.MultiSetOf(Companion_Default___.Fm16((_this).F27(), globalState), (_this).F30())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_826_v36), 0))
			_ = _index82
			(_826_v36).ArraySet1((Companion_Default___.Fm18(globalState)).IsSubsetOf(_827_v37), (_index82).Int())
			var _828_v38 _dafny.Map
			_ = _828_v38
			_828_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_826_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_826_v36), 0))).Int()).(bool), (_this).F27())
			var _829_v39 _dafny.Map
			_ = _829_v39
			_829_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_dafny.Zero).Minus((_804_i2).Minus((func() _dafny.Int {
				if (_828_v38).Contains(_this.F35) {
					return (_828_v38).Get(_this.F35).(_dafny.Int)
				}
				return (_this).F27()
			})())))
			_829_v39 = (_829_v39).Update((_this).F28(), (_this).F27())
		}
		if (((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(Companion_Default___.SafeModuloInt((_this).F27(), _dafny.IntOfInt64(604))) <= 0 {
			var _830_v40 D4
			_ = _830_v40
			_830_v40 = Companion_D4_.Create_DC9_((_this).F27(), p0)
			var _831_v41 _dafny.Set
			_ = _831_v41
			_831_v41 = _dafny.SetOf((_830_v40).Dtor_cf17())
			var _832_v42 _dafny.Sequence
			_ = _832_v42
			_832_v42 = _dafny.SeqOf(_dafny.SetOf(p1))
			var _833_v43 _dafny.MultiSet
			_ = _833_v43
			_833_v43 = _dafny.MultiSetOf((_this).F27())
			var _834_v44 _dafny.MultiSet
			_ = _834_v44
			_834_v44 = _dafny.MultiSetOf((_this).F27(), (_782_v1).Cardinality(), (_833_v43).Cardinality(), (_this).F27(), (_this).F27())
			var _835_v45 _dafny.MultiSet
			_ = _835_v45
			_835_v45 = _833_v43
			var _836_v46 _dafny.Set
			_ = _836_v46
			_836_v46 = _dafny.SetOf((_835_v45))
			var _rhs115 _dafny.Set = _dafny.SetOf(p0, (_dafny.IntOfInt64(314)).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_832_v42).Cardinality()))) < 0, !((_836_v46).IsDisjointFrom(_836_v46)), Companion_Default___.Fm0((_831_v41).Cardinality(), globalState), p0)
			_ = _rhs115
			var _rhs116 _dafny.Int = (_this).F27()
			_ = _rhs116
			var _lhs86 *GlobalState = globalState
			_ = _lhs86
			_831_v41 = _rhs115
			_lhs86.F3 = _rhs116
			var _837_v47 _dafny.Array
			_ = _837_v47
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw139
			_837_v47 = _nw139
			_837_v47 = _837_v47
			var _838_v48 _dafny.Array
			_ = _838_v48
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw140
			_838_v48 = _nw140
			r1 = _838_v48
			var _839_v49 _dafny.Array
			_ = _839_v49
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(27))
			_ = _nw141
			_839_v49 = _nw141
			var _840_v50 _dafny.CodePoint
			_ = _840_v50
			_840_v50 = _dafny.CodePoint('v')
			var _841_v51 _dafny.Map
			_ = _841_v51
			_841_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v50, p1)
			var _842_v52 D2
			_ = _842_v52
			_842_v52 = Companion_D2_.Create_DC3_(_841_v51)
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_839_v49), 0))
			_ = _index83
			(_839_v49).ArraySet1(_842_v52, (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_839_v49), 0))
			_ = _index84
			(_839_v49).ArraySet1(_842_v52, (_index84).Int())
			r0 = Companion_Default___.Fm0((func() _dafny.Int {
				if ((_this).F29()).Contains(p1) {
					return ((_this).F29()).Multiplicity(p1)
				}
				return (_this).F27()
			})(), globalState)
		} else {
			var _843_v53 _dafny.Array
			_ = _843_v53
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw142
			_843_v53 = _nw142
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_843_v53), 0))
			_ = _index85
			(_843_v53).ArraySet1((_781_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32()).(bool), (_index85).Int())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_843_v53), 0))
			_ = _index86
			(_843_v53).ArraySet1(false, (_index86).Int())
			if ((_this).F27()).Cmp((_this).F27()) == 0 {
				var _844_v54 _dafny.Map
				_ = _844_v54
				_844_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.IntOfInt64(962))
				var _845_v55 _dafny.CodePoint
				_ = _845_v55
				_845_v55 = _dafny.CodePoint('o')
				var _846_v56 _dafny.Set
				_ = _846_v56
				_846_v56 = _dafny.SetOf(_801_v17, _dafny.Companion_Sequence_.Update(_801_v17, (Companion_Default___.SafeIndex((_844_v54).Cardinality(), _dafny.IntOfUint32((_801_v17).Cardinality()))).Uint32(), _845_v55))
				var _847_v57 _dafny.Array
				_ = _847_v57
				var _nwElement0_33 _dafny.Int = (_this).F27()
				_ = _nwElement0_33
				var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(8))
				_ = _nw143
				(_nw143).ArraySet1(_nwElement0_33, 0)
				(_nw143).ArraySet1(_dafny.IntOfUint32(((_this).F30()).Cardinality()), 1)
				(_nw143).ArraySet1(_dafny.IntOfInt64(818), 2)
				(_nw143).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F27())), (_this).F27()), 3)
				(_nw143).ArraySet1((_dafny.Zero).Minus((_this).F27()), 4)
				(_nw143).ArraySet1((_dafny.Zero).Minus((_this).F27()), 5)
				(_nw143).ArraySet1((_this).Fm4(_846_v56, p0, _781_v0, globalState), 6)
				(_nw143).ArraySet1((_dafny.Zero).Minus(((_this).F27()).Times((_this).F27())), 7)
				_847_v57 = _nw143
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))
				_ = _index87
				(_847_v57).ArraySet1((_dafny.Zero).Minus((_this).F27()), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))
				_ = _index88
				(_847_v57).ArraySet1(Companion_Default___.SafeModuloInt((_this).F27(), (_this).F27()), (_index88).Int())
				(globalState).F26 = !(_this.F35)
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))
				_ = _index89
				var _rhs117 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(((_this).F27()).Cmp((_this).F27()) > 0, p1)).Cardinality())
				_ = _rhs117
				var _rhs118 _dafny.Int = (((_this).F30()).Select((Companion_Default___.SafeIndex((_847_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)).Times((_847_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))).Int()).(_dafny.Int))
				_ = _rhs118
				var _rhs119 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("vvpnx")
				_ = _rhs119
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				var _lhs88 _dafny.Array = _847_v57
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))
				_ = _lhs89
				_lhs87.F24 = _rhs117
				(_lhs88).ArraySet1(_rhs118, (_lhs89).Int())
				_801_v17 = _rhs119
				(globalState).F3 = (_847_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))).Int()).(_dafny.Int)
				var _848_v58 _dafny.Map
				_ = _848_v58
				_848_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _803_v19)
				var _rhs120 bool = !(((_dafny.IntOfUint32((_801_v17).Cardinality())).Plus((_847_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_847_v57), 0))).Int()).(_dafny.Int))).Cmp(Companion_Default___.Fm1(globalState)) >= 0)
				_ = _rhs120
				var _rhs121 _dafny.Map = ((func() _dafny.Map {
					if (_848_v58).Contains((_this).F27()) {
						return (_848_v58).Get((_this).F27()).(_dafny.Map)
					}
					return _803_v19
				})()).Merge(_803_v19)
				_ = _rhs121
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				_lhs90.F26 = _rhs120
				_803_v19 = _rhs121
			} else {
				var _849_v59 _dafny.Set
				_ = _849_v59
				_849_v59 = _dafny.SetOf(p1)
				var _850_v60 _dafny.Map
				_ = _850_v60
				_850_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.SetOf(p0)).Difference(_849_v59))
				_850_v60 = _850_v60
				(globalState).F24 = Companion_Default___.SafeModuloInt((_this).F27(), (_this).F27())
				r1 = _843_v53
				var _851_v61 _dafny.Array
				_ = _851_v61
				var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(3))
				_ = _nw144
				_851_v61 = _nw144
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_851_v61), 0))
				_ = _index90
				(_851_v61).ArraySet1(_782_v1, (_index90).Int())
				var _852_v62 _dafny.Map
				_ = _852_v62
				_852_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_851_v61), 0))
				_ = _index91
				(_851_v61).ArraySet1(func() _dafny.Set {
					var _coll35 = _dafny.NewBuilder()
					_ = _coll35
					for _iter36 := _dafny.Iterate((_852_v62).Keys().Elements()); ; {
						_compr_35, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _853_v63 _dafny.Int
						_853_v63 = interface{}(_compr_35).(_dafny.Int)
						if (_852_v62).Contains(_853_v63) {
							_coll35.Add((_853_v63).Minus((_dafny.IntOfInt64(782)).Times(_dafny.IntOfInt64(566))))
						}
					}
					return _coll35.ToSet()
				}(), (_index91).Int())
				var _854_v64 _dafny.Map
				_ = _854_v64
				_854_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, (_this).F27())
				var _855_v65 _dafny.Sequence
				_ = _855_v65
				_855_v65 = _dafny.SeqOf(_854_v64, (_854_v64).Merge(_854_v64), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _dafny.IntOfUint32((_801_v17).Cardinality())), _854_v64)
				_855_v65 = _dafny.SeqOf(_854_v64)
			}
			var _856_v66 _dafny.CodePoint
			_ = _856_v66
			_856_v66 = _dafny.CodePoint('y')
			var _857_v67 *C0
			_ = _857_v67
			var _nw145 *C0 = New_C0_()
			_ = _nw145
			_nw145.Ctor__((_this).F27(), _dafny.Companion_Sequence_.Update(_801_v17, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_801_v17).Cardinality()))).Uint32(), _856_v66))
			_857_v67 = _nw145
			var _858_v68 D6
			_ = _858_v68
			_858_v68 = Companion_D6_.Create_DC11_(_857_v67)
			var _859_v69 _dafny.Array
			_ = _859_v69
			var _nwElement0_34 *C0 = _857_v67
			_ = _nwElement0_34
			var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(23))
			_ = _nw146
			(_nw146).ArraySet1(_nwElement0_34, 0)
			(_nw146).ArraySet1(_857_v67, 1)
			(_nw146).ArraySet1(_857_v67, 2)
			(_nw146).ArraySet1(_857_v67, 3)
			(_nw146).ArraySet1(_857_v67, 4)
			(_nw146).ArraySet1(_857_v67, 5)
			(_nw146).ArraySet1(_857_v67, 6)
			(_nw146).ArraySet1(_857_v67, 7)
			(_nw146).ArraySet1(_857_v67, 8)
			(_nw146).ArraySet1(_857_v67, 9)
			(_nw146).ArraySet1(_857_v67, 10)
			(_nw146).ArraySet1(_857_v67, 11)
			(_nw146).ArraySet1(_857_v67, 12)
			(_nw146).ArraySet1(_857_v67, 13)
			(_nw146).ArraySet1(_857_v67, 14)
			(_nw146).ArraySet1(_857_v67, 15)
			(_nw146).ArraySet1(_857_v67, 16)
			(_nw146).ArraySet1(_857_v67, 17)
			(_nw146).ArraySet1(_857_v67, 18)
			(_nw146).ArraySet1(_857_v67, 19)
			(_nw146).ArraySet1(_857_v67, 20)
			(_nw146).ArraySet1(_857_v67, 21)
			(_nw146).ArraySet1((_858_v68).Dtor_cf19(), 22)
			_859_v69 = _nw146
			var _rhs122 _dafny.Array = _859_v69
			_ = _rhs122
			var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_857_v67).F37(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_857_v67).F36()), _dafny.IntOfUint32(((_857_v67).F37()).Cardinality()))).Uint32(), _dafny.CodePoint('b')), (_857_v67).F37())
			_ = _rhs123
			var _rhs124 _dafny.Int = (_857_v67).F36()
			_ = _rhs124
			var _lhs91 *GlobalState = globalState
			_ = _lhs91
			_859_v69 = _rhs122
			_801_v17 = _rhs123
			_lhs91.F24 = _rhs124
			var _860_v70 _dafny.Map
			_ = _860_v70
			_860_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_857_v67).F37(), (_dafny.Zero).Minus((_this).F27()))
			_860_v70 = (_860_v70).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(935))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_861_v66 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_862_i8 _dafny.Int) _dafny.CodePoint {
					return _861_v66
				}
			})(_856_v66))), (_857_v67).F36())
			var _863_v71 _dafny.Map
			_ = _863_v71
			_863_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_843_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_843_v53), 0))).Int()).(bool), (_857_v67).F36())
			_863_v71 = (_863_v71).Update((_843_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_843_v53), 0))).Int()).(bool), (_dafny.IntOfInt64(-461)).Times((_803_v19).Cardinality()))
		}
		if (_781_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32()).(bool) {
			(globalState).F26 = p1
			(globalState).F3 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p0 {
					return (Companion_D7_.Create_DC13_(_dafny.SeqOf(_this.F35, true))).Dtor_cf23()
				}
				return _781_v0
			})()).Cardinality())
			(globalState).F0 = (_this).F27()
			if _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm14(_this.F35, _this.F35, p1, (_this).F27(), globalState), _dafny.Companion_Sequence_.Concatenate(_801_v17, _dafny.UnicodeSeqOfUtf8Bytes("qk"))) {
				(globalState).F15 = (_this).F27()
				var _864_v72 *C0
				_ = _864_v72
				var _nw147 *C0 = New_C0_()
				_ = _nw147
				_nw147.Ctor__((_this).F27(), _801_v17)
				_864_v72 = _nw147
				(globalState).F26 = _this.F35
				var _865_v73 _dafny.Array
				_ = _865_v73
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw148
				_865_v73 = _nw148
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_865_v73), 0))
				_ = _index92
				(_865_v73).ArraySet1(p0, (_index92).Int())
				var _866_v74 D6
				_ = _866_v74
				_866_v74 = Companion_D6_.Create_DC11_(_864_v72)
				var _867_v75 _dafny.Map
				_ = _867_v75
				_867_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _866_v74)
				var _868_v76 _dafny.Map
				_ = _868_v76
				_868_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_867_v75).Cardinality()).Times((_this).F27()), true)
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_865_v73), 0))
				_ = _index93
				(_865_v73).ArraySet1((func() bool {
					if (_868_v76).Contains(_dafny.IntOfInt64(341)) {
						return (_868_v76).Get(_dafny.IntOfInt64(341)).(bool)
					}
					return _this.F35
				})(), (_index93).Int())
				var _869_v77 _dafny.Set
				_ = _869_v77
				_869_v77 = _dafny.SetOf(_dafny.MultiSetFromSeq(_781_v0))
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_865_v73), 0))
				_ = _index94
				var _rhs125 bool = p0
				_ = _rhs125
				var _rhs126 bool = !((_869_v77).IsSubsetOf(_869_v77))
				_ = _rhs126
				var _rhs127 _dafny.Int = ((_this).F27()).Times((_864_v72).F36())
				_ = _rhs127
				var _lhs92 _dafny.Array = _865_v73
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_865_v73), 0))
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				(_lhs92).ArraySet1(_rhs125, (_lhs93).Int())
				_lhs94.F26 = _rhs126
				_lhs95.F24 = _rhs127
			} else {
				var _870_v78 D8
				_ = _870_v78
				_870_v78 = Companion_D8_.Create_DC15_(_803_v19)
				var _871_v79 _dafny.Map
				_ = _871_v79
				_871_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_870_v78).Dtor_cf28(), p0)
				_781_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func() bool {
					if (_871_v79).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, p0)) {
						return (_871_v79).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, p0)).(bool)
					}
					return p0
				})()), _dafny.Companion_Sequence_.Update(_781_v0, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32(), p1))
				var _872_v80 _dafny.CodePoint
				_ = _872_v80
				_872_v80 = _dafny.CodePoint('u')
				var _873_v81 _dafny.MultiSet
				_ = _873_v81
				_873_v81 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if p0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-284))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg66 _dafny.Int) interface{} {
								return coer66(arg66)
							}
						}(func(_874_i9 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						}))
					}
					return _dafny.Companion_Sequence_.Update(_801_v17, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_801_v17).Cardinality()))).Uint32(), _872_v80)
				})())
				(globalState).F3 = (func() _dafny.Int {
					if (_873_v81).Contains(_dafny.Companion_Sequence_.Concatenate(_801_v17, _801_v17)) {
						return (_873_v81).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_801_v17, _801_v17))
					}
					return (_dafny.Zero).Minus((func() _dafny.Int {
						if !(_this.F35) {
							return (_this).F27()
						}
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg67 _dafny.Int) interface{} {
								return coer67(arg67)
							}
						}((func(_875_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
							return func(_876_i10 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_875_v0).Cardinality())
							}
						})(_781_v0)))).Cardinality())
					})())
				})()
				(globalState).F18 = (_this).F27()
				(_this).F35 = !(_this.F35)
				var _877_v82 _dafny.MultiSet
				_ = _877_v82
				_877_v82 = _dafny.MultiSetOf((_this).F27(), _dafny.IntOfUint32((_801_v17).Cardinality()), (_this).F27())
				_877_v82 = _877_v82
			}
			var _878_v83 _dafny.Array
			_ = _878_v83
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw149
			_878_v83 = _nw149
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_878_v83), 0))
			_ = _index95
			(_878_v83).ArraySet1(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter37 := _dafny.Iterate((Companion_Default___.Fm19(globalState)).Elements()); ; {
					_compr_36, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _879_v84 _dafny.CodePoint
					_879_v84 = interface{}(_compr_36).(_dafny.CodePoint)
					if (Companion_Default___.Fm19(globalState)).Contains(_879_v84) {
						_coll36.Add(_879_v84, (_this).F27())
					}
				}
				return _coll36.ToMap()
			}(), (_index95).Int())
			var _880_v85 _dafny.CodePoint
			_ = _880_v85
			_880_v85 = _dafny.CodePoint('t')
			var _881_v86 _dafny.Map
			_ = _881_v86
			_881_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_880_v85, (_this).F27())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_878_v83), 0))
			_ = _index96
			(_878_v83).ArraySet1(_881_v86, (_index96).Int())
		} else {
			(globalState).F0 = (_this).F27()
			(globalState).F18 = (_this).F27()
			var _882_v87 _dafny.Map
			_ = _882_v87
			_882_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())
			var _883_v88 _dafny.CodePoint
			_ = _883_v88
			_883_v88 = _dafny.CodePoint('k')
			var _884_v89 D4
			_ = _884_v89
			_884_v89 = Companion_D4_.Create_DC8_(_883_v88, _dafny.IntOfInt64(57), p1)
			_882_v87 = (_882_v87).Update((_884_v89).Dtor_cf14(), (_this).F27())
			var _885_v90 _dafny.Array
			_ = _885_v90
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw150
			_885_v90 = _nw150
			var _886_v91 _dafny.Sequence
			_ = _886_v91
			_886_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(208))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}(func(_887_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()))
			var _rhs128 _dafny.Array = _885_v90
			_ = _rhs128
			var _rhs129 _dafny.Int = (_this).F27()
			_ = _rhs129
			var _rhs130 _dafny.Sequence = (_this).F30()
			_ = _rhs130
			var _rhs131 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F27(), _dafny.IntOfUint32((_801_v17).Cardinality()))
			_ = _rhs131
			var _rhs132 _dafny.Int = (_this).F27()
			_ = _rhs132
			var _lhs96 *GlobalState = globalState
			_ = _lhs96
			var _lhs97 *GlobalState = globalState
			_ = _lhs97
			var _lhs98 *GlobalState = globalState
			_ = _lhs98
			_885_v90 = _rhs128
			_lhs96.F21 = _rhs129
			_886_v91 = _rhs130
			_lhs97.F0 = _rhs131
			_lhs98.F0 = _rhs132
			_801_v17 = _dafny.Companion_Sequence_.Concatenate(_801_v17, _dafny.Companion_Sequence_.Concatenate(_801_v17, Companion_Default___.Fm14(_this.F35, _this.F35, p1, (_this).F27(), globalState)))
		}
		r0 = p0
		var _888_v92 _dafny.Map
		_ = _888_v92
		_888_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F27()), (_this).F27())
		var _889_v93 _dafny.Map
		_ = _889_v93
		_889_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(globalState), _888_v92)
		var _890_v94 _dafny.CodePoint
		_ = _890_v94
		_890_v94 = _dafny.CodePoint('r')
		var _891_v95 _dafny.Map
		_ = _891_v95
		_891_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _890_v94)
		var _nwElement0_35 bool = _this.F35
		_ = _nwElement0_35
		var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(21))
		_ = _nw151
		(_nw151).ArraySet1(_nwElement0_35, 0)
		(_nw151).ArraySet1(p1, 1)
		(_nw151).ArraySet1((p1) && (p0), 2)
		(_nw151).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_781_v0, _781_v0), 3)
		(_nw151).ArraySet1(p0, 4)
		(_nw151).ArraySet1(p0, 5)
		(_nw151).ArraySet1(p0, 6)
		(_nw151).ArraySet1(p1, 7)
		(_nw151).ArraySet1(((_this).F27()).Cmp((_this).F27()) >= 0, 8)
		(_nw151).ArraySet1(p0, 9)
		(_nw151).ArraySet1(p0, 10)
		(_nw151).ArraySet1(p1, 11)
		(_nw151).ArraySet1((_dafny.IntOfUint32(((_this).F30()).Cardinality())).Cmp((_this).F27()) >= 0, 12)
		(_nw151).ArraySet1((_781_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_781_v0).Cardinality()))).Uint32()).(bool), 13)
		(_nw151).ArraySet1(p1, 14)
		(_nw151).ArraySet1((_this).Fm11(p1, (_889_v93).Cardinality(), _802_v18, (_891_v95).Cardinality(), globalState), 15)
		(_nw151).ArraySet1((_this).Fm11(p1, (_this).F27(), _802_v18, (_this).F27(), globalState), 16)
		(_nw151).ArraySet1((_this).Fm11(p0, _dafny.IntOfUint32((_801_v17).Cardinality()), _802_v18, (_this).F27(), globalState), 17)
		(_nw151).ArraySet1(p1, 18)
		(_nw151).ArraySet1((_this).Fm11(p1, (_this).F27(), _802_v18, _dafny.IntOfUint32((_801_v17).Cardinality()), globalState), 19)
		(_nw151).ArraySet1(_this.F35, 20)
		r1 = _nw151
		var _892_v96 _dafny.MultiSet
		_ = _892_v96
		_892_v96 = _dafny.MultiSetOf((_this).F27())
		var _893_v97 _dafny.Map
		_ = _893_v97
		_893_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if p1 {
				return _801_v17
			}
			return _801_v17
		})()).Cardinality()), !((_892_v96).IsProperSubsetOf(_892_v96)))
		r2 = _893_v97
		return r0, r1, r2
	}
}
func (_this *C1) M6(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _894_v0 _dafny.Array
		_ = _894_v0
		var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw152
		_894_v0 = _nw152
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_894_v0), 0))); ; {
			_guard_loop_1, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _895_i0 _dafny.Int
			_895_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_895_i0).Sign() != -1) && ((_895_i0).Cmp(_dafny.ArrayLen((_894_v0), 0)) < 0)) {
				(_894_v0).ArraySet1((_895_i0).Plus((_dafny.SetOf(p0)).Cardinality()), (_895_i0).Int())
			}
		}
		(globalState).F24 = ((_this).F30()).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)
		var _896_v1 _dafny.Array
		_ = _896_v1
		var _nwElement0_36 bool = _this.F35
		_ = _nwElement0_36
		var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(4))
		_ = _nw153
		(_nw153).ArraySet1(_nwElement0_36, 0)
		(_nw153).ArraySet1(Companion_Default___.Fm0((_this).F27(), globalState), 1)
		(_nw153).ArraySet1(_this.F35, 2)
		(_nw153).ArraySet1(_this.F35, 3)
		_896_v1 = _nw153
		var _897_v2 D4
		_ = _897_v2
		_897_v2 = Companion_D4_.Create_DC7_(_896_v1)
		var _pat_let_tv7 = _896_v1
		_ = _pat_let_tv7
		var _898_v3 _dafny.Map
		_ = _898_v3
		_898_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let14_0 D4) D4 {
			return func(_899_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let15_0 _dafny.Array) D4 {
					return func(_900_dt__update_hcf12_h0 _dafny.Array) D4 {
						return Companion_D4_.Create_DC7_(_900_dt__update_hcf12_h0)
					}(_pat_let15_0)
				}(_pat_let_tv7)
			}(_pat_let14_0)
		}(_897_v2), _dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F35, false, _this.F35, _this.F35, _this.F35)))
		var _901_v4 _dafny.Map
		_ = _901_v4
		_901_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, (_898_v3).Cardinality())
		var _902_v5 _dafny.Sequence
		_ = _902_v5
		_902_v5 = _dafny.UnicodeSeqOfUtf8Bytes("it")
		var _903_v6 _dafny.Map
		_ = _903_v6
		_903_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_902_v5, (_this).F27())
		var _904_v7 _dafny.Map
		_ = _904_v7
		_904_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm16((func() _dafny.Int {
			if (_901_v4).Contains(_this.F35) {
				return (_901_v4).Get(_this.F35).(_dafny.Int)
			}
			return p0
		})(), globalState), (_this).F30()), Companion_Default___.SafeDivisionInt((_this).F27(), (func() _dafny.Int {
			if (_903_v6).Contains(_902_v5) {
				return (_903_v6).Get(_902_v5).(_dafny.Int)
			}
			return (_this).F27()
		})()))
		_904_v7 = (_904_v7).Update(_dafny.Companion_Sequence_.Concatenate((_this).F30(), (_this).F30()), (_this).F27())
		_902_v5 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14(true, _this.F35, _this.F35, _dafny.IntOfUint32((_902_v5).Cardinality()), globalState), Companion_Default___.Fm14(_this.F35, _this.F35, _this.F35, _dafny.IntOfUint32((Companion_Default___.Fm14(_this.F35, _this.F35, _this.F35, (_this).F27(), globalState)).Cardinality()), globalState))
		if !((func() bool {
			if _this.F35 {
				return _this.F35
			}
			return _this.F35
		})()) {
			var _905_v8 _dafny.CodePoint
			_ = _905_v8
			_905_v8 = _dafny.CodePoint('w')
			var _906_v9 D2
			_ = _906_v9
			_906_v9 = Companion_D2_.Create_DC4_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, false)).Cardinality(), _this.F35)).Cardinality(), ((_this).F27()).Minus(p0), (_dafny.Zero).Minus(p0), _905_v8, _dafny.Companion_Sequence_.Concatenate(_902_v5, _902_v5))
			var _907_v10 _dafny.Map
			_ = _907_v10
			_907_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v8, _this.F35)
			var _908_v11 D2
			_ = _908_v11
			_908_v11 = Companion_D2_.Create_DC3_(_907_v10)
			var _pat_let_tv8 = _907_v10
			_ = _pat_let_tv8
			var _909_v12 D2
			_ = _909_v12
			_909_v12 = Companion_D2_.Create_DC3_((func(_pat_let16_0 D2) D2 {
				return func(_910_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let17_0 _dafny.Map) D2 {
						return func(_911_dt__update_hcf4_h0 _dafny.Map) D2 {
							return Companion_D2_.Create_DC3_(_911_dt__update_hcf4_h0)
						}(_pat_let17_0)
					}(_pat_let_tv8)
				}(_pat_let16_0)
			}(_908_v11)).Dtor_cf4())
			var _912_v13 D2
			_ = _912_v13
			_912_v13 = Companion_D2_.Create_DC5_(_909_v12)
			var _rhs133 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_902_v5, _902_v5)
			_ = _rhs133
			var _rhs134 D2 = _906_v9
			_ = _rhs134
			var _rhs135 _dafny.Int = (_this).F27()
			_ = _rhs135
			var _rhs136 D2 = _912_v13
			_ = _rhs136
			var _rhs137 bool = false
			_ = _rhs137
			var _lhs99 *GlobalState = globalState
			_ = _lhs99
			var _lhs100 *GlobalState = globalState
			_ = _lhs100
			_902_v5 = _rhs133
			_906_v9 = _rhs134
			_lhs99.F21 = _rhs135
			_912_v13 = _rhs136
			_lhs100.F26 = _rhs137
			var _913_v14 _dafny.Set
			_ = _913_v14
			_913_v14 = _dafny.SetOf(_this.F35, false)
			var _914_v15 _dafny.Sequence
			_ = _914_v15
			_914_v15 = _dafny.SeqOf(_913_v14)
			var _rhs138 bool = _this.F35
			_ = _rhs138
			var _rhs139 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_914_v15, _dafny.SeqOf(_913_v14)), _914_v15), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_914_v15, _dafny.SeqOf(_913_v14)), _914_v15)).Cardinality()))).Uint32(), _913_v14)
			_ = _rhs139
			var _lhs101 *C1 = _this
			_ = _lhs101
			_lhs101.F35 = _rhs138
			_914_v15 = _rhs139
			(globalState).F26 = _this.F35
			var _915_v16 _dafny.Sequence
			_ = _915_v16
			_915_v16 = _dafny.SeqOf(Companion_Default___.Fm14(_this.F35, _this.F35, _this.F35, p0, globalState), _902_v5)
			var _916_v17 D6
			_ = _916_v17
			_916_v17 = Companion_D6_.Create_DC12_(_this.F35, Companion_Default___.Fm20(_this.F35, globalState), (_915_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_915_v16).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _source14 D6 = _916_v17
			_ = _source14
			if _source14.Is_DC12() {
				var _917___mcc_h0 bool = _source14.Get_().(D6_DC12).Cf20
				_ = _917___mcc_h0
				var _918___mcc_h1 _dafny.Sequence = _source14.Get_().(D6_DC12).Cf21
				_ = _918___mcc_h1
				var _919___mcc_h2 _dafny.Sequence = _source14.Get_().(D6_DC12).Cf22
				_ = _919___mcc_h2
				var _920_cf22 _dafny.Sequence = _919___mcc_h2
				_ = _920_cf22
				var _921_cf21 _dafny.Sequence = _918___mcc_h1
				_ = _921_cf21
				var _922_cf20 bool = _917___mcc_h0
				_ = _922_cf20
				var _923_v18 _dafny.Array
				_ = _923_v18
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_16
				var _nw154 _dafny.Array
				_ = _nw154
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw154 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Set = (func(_924_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_925_i1 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(_924_p0)).Difference((Companion_D9_.Create_DC17_(_dafny.SetOf(((_this).F29()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vlmvh")).Cardinality())))).Dtor_cf31())
						}
					})(p0)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw154 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw154).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw154).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_923_v18 = _nw154
				var _rhs140 _dafny.Array = (func() _dafny.Array {
					if _this.F35 {
						return _923_v18
					}
					return _923_v18
				})()
				_ = _rhs140
				var _rhs141 _dafny.CodePoint = _905_v8
				_ = _rhs141
				_923_v18 = _rhs140
				_905_v8 = _rhs141
				var _926_v19 _dafny.Set
				_ = _926_v19
				_926_v19 = _dafny.SetOf(_905_v8)
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_894_v0), 0))
				_ = _index97
				(_894_v0).ArraySet1((_926_v19).Cardinality(), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_894_v0), 0))
				_ = _index98
				(_894_v0).ArraySet1((p0).Minus((p0).Minus((_this).F27())), (_index98).Int())
				var _927_v20 _dafny.MultiSet
				_ = _927_v20
				_927_v20 = _dafny.MultiSetOf((_this).F27(), p0, (_894_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_894_v0), 0))).Int()).(_dafny.Int), (_this).F27(), (_this).F27())
				var _928_v21 D8
				_ = _928_v21
				_928_v21 = Companion_D8_.Create_DC16_((_this).F27(), (_this).F27())
				var _929_v22 _dafny.Map
				_ = _929_v22
				_929_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm13(p0, _dafny.IntOfInt64(612), !(_922_cf20), _dafny.IntOfUint32((_920_cf22).Cardinality()), globalState)).Cardinality()), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), _928_v21)
				var _930_v23 _dafny.Map
				_ = _930_v23
				_930_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v8, (_this).F27())).Cardinality(), (_this).F27())
				var _931_v24 bool
				_ = _931_v24
				var _932_v25 _dafny.Int
				_ = _932_v25
				var _out33 bool
				_ = _out33
				var _out34 _dafny.Int
				_ = _out34
				_out33, _out34 = Companion_Default___.M0((_927_v20).IsDisjointFrom(_dafny.MultiSetOf((_894_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_894_v0), 0))).Int()).(_dafny.Int))), (_929_v22).Cardinality(), (_this).F27(), (func() _dafny.Int {
					if (_930_v23).Contains((_this).F27()) {
						return (_930_v23).Get((_this).F27()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_902_v5).Cardinality())
				})(), globalState)
				_931_v24 = _out33
				_932_v25 = _out34
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_896_v1), 0))
				_ = _index99
				(_896_v1).ArraySet1(_922_cf20, (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_896_v1), 0))
				_ = _index100
				(_896_v1).ArraySet1(_this.F35, (_index100).Int())
			} else {
				var _933___mcc_h3 *C0 = _source14.Get_().(D6_DC11).Cf19
				_ = _933___mcc_h3
				var _934_cf19 *C0 = _933___mcc_h3
				_ = _934_cf19
				(globalState).F26 = _this.F35
				var _935_v26 _dafny.MultiSet
				_ = _935_v26
				_935_v26 = _dafny.MultiSetOf((_934_cf19).F36(), _dafny.IntOfInt64(-805))
				var _936_v27 _dafny.Sequence
				_ = _936_v27
				_936_v27 = _dafny.SeqOf(_this.F35, _this.F35)
				var _937_v28 D1
				_ = _937_v28
				_937_v28 = Companion_D1_.Create_DC2_(_905_v8, (_dafny.Zero).Minus(_dafny.IntOfUint32((_936_v27).Cardinality())))
				var _938_v29 _dafny.Set
				_ = _938_v29
				_938_v29 = _dafny.SetOf((_935_v26).Cardinality(), (_937_v28).Dtor_cf3())
				var _939_v30 bool
				_ = _939_v30
				var _940_v31 _dafny.Array
				_ = _940_v31
				var _941_v32 _dafny.Map
				_ = _941_v32
				var _out35 bool
				_ = _out35
				var _out36 _dafny.Array
				_ = _out36
				var _out37 _dafny.Map
				_ = _out37
				_out35, _out36, _out37 = (_this).M5(_this.F35, (_938_v29).IsDisjointFrom(_938_v29), globalState)
				_939_v30 = _out35
				_940_v31 = _out36
				_941_v32 = _out37
				(globalState).F24 = (_this).F27()
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_894_v0), 0))
				_ = _index101
				(_894_v0).ArraySet1(((_this).F27()).Plus(p0), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_894_v0), 0))
				_ = _index102
				(_894_v0).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_902_v5).Cardinality()), p0), (_index102).Int())
			}
			var _942_v33 _dafny.Array
			_ = _942_v33
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw155
			_942_v33 = _nw155
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_942_v33), 0))
			_ = _index103
			(_942_v33).ArraySet1(_896_v1, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_942_v33), 0))
			_ = _index104
			(_942_v33).ArraySet1(_896_v1, (_index104).Int())
		} else {
			var _943_v34 _dafny.Map
			_ = _943_v34
			_943_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(767), _this.F35)
			(globalState).F21 = ((p0).Minus((_dafny.Zero).Minus((_943_v34).Cardinality()))).Plus((_this).F27())
			(globalState).F26 = (p0).Cmp(Companion_Default___.Fm1(globalState)) == 0
			var _944_v35 *C0
			_ = _944_v35
			var _nw156 *C0 = New_C0_()
			_ = _nw156
			_nw156.Ctor__((_this).F27(), _902_v5)
			_944_v35 = _nw156
			var _945_v36 _dafny.Map
			_ = _945_v36
			_945_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(364), _944_v35)
			var _946_v37 D6
			_ = _946_v37
			_946_v37 = Companion_D6_.Create_DC12_(_this.F35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_947_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_this.F35)
			})), (_944_v35).F37())
			var _948_v38 D6
			_ = _948_v38
			_948_v38 = Companion_D6_.Create_DC11_((func() *C0 {
				if (_945_v36).Contains(_dafny.IntOfUint32(((_946_v37).Dtor_cf22()).Cardinality())) {
					return (_945_v36).Get(_dafny.IntOfUint32(((_946_v37).Dtor_cf22()).Cardinality())).(*C0)
				}
				return _944_v35
			})())
			_948_v38 = _948_v38
			_943_v34 = (_943_v34).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), true))
			(globalState).F3 = (_944_v35).F36()
		}
		var _949_v39 _dafny.Map
		_ = _949_v39
		_949_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _902_v5)
		var _950_v40 *C0
		_ = _950_v40
		var _nw157 *C0 = New_C0_()
		_ = _nw157
		_nw157.Ctor__((_this).F27(), (func() _dafny.Sequence {
			if (_949_v39).Contains(true) {
				return (_949_v39).Get(true).(_dafny.Sequence)
			}
			return _902_v5
		})())
		_950_v40 = _nw157
		var _951_v41 _dafny.Map
		_ = _951_v41
		_951_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35, _this.F35)
		var _952_v42 D8
		_ = _952_v42
		_952_v42 = Companion_D8_.Create_DC15_(_951_v41)
		var _pat_let_tv9 = p0
		_ = _pat_let_tv9
		r0 = func(_source15 D8) _dafny.Int {
			if _source15.Is_DC16() {
				var _953___mcc_h4 _dafny.Int = _source15.Get_().(D8_DC16).Cf29
				_ = _953___mcc_h4
				var _954___mcc_h5 _dafny.Int = _source15.Get_().(D8_DC16).Cf30
				_ = _954___mcc_h5
				var _955_cf30 _dafny.Int = _954___mcc_h5
				_ = _955_cf30
				var _956_cf29 _dafny.Int = _953___mcc_h4
				_ = _956_cf29
				return _956_cf29
			} else {
				var _957___mcc_h6 _dafny.Map = _source15.Get_().(D8_DC15).Cf28
				_ = _957___mcc_h6
				var _958_cf28 _dafny.Map = _957___mcc_h6
				_ = _958_cf28
				return _pat_let_tv9
			}
		}(_952_v42)
		return r0
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f31 _dafny.MultiSet
	_f32 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f31 = _dafny.EmptyMultiSet
	_this._f32 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_}
}

var _ T2 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F31() _dafny.MultiSet {
	return _this._f31
}
func (_this *C2) F32() _dafny.Int {
	return _this._f32
}
func (_this *C2) Ctor__(f31 _dafny.MultiSet, f32 _dafny.Int) {
	{
		(_this)._f31 = f31
		(_this)._f32 = f32
	}
}
func (_this *C2) Fm5(globalState *GlobalState) bool {
	{
		return !(func(_source16 D9) bool {
			if _source16.Is_DC18() {
				var _959___mcc_h0 bool = _source16.Get_().(D9_DC18).Cf32
				_ = _959___mcc_h0
				var _960_cf32 bool = _959___mcc_h0
				_ = _960_cf32
				return (_960_cf32) || (_960_cf32)
			} else if _source16.Is_DC19() {
				var _961___mcc_h1 bool = _source16.Get_().(D9_DC19).Cf33
				_ = _961___mcc_h1
				var _962_cf33 bool = _961___mcc_h1
				_ = _962_cf33
				return (Companion_D9_.Create_DC20_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okdff")).Cardinality()), _962_cf33, _962_cf33, _962_cf33, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_962_cf33, _962_cf33)).Cardinality()), (_dafny.SetOf(_962_cf33)).Cardinality()))).Dtor_cf36()
			} else if _source16.Is_DC20() {
				var _963___mcc_h2 _dafny.Int = _source16.Get_().(D9_DC20).Cf34
				_ = _963___mcc_h2
				var _964___mcc_h3 bool = _source16.Get_().(D9_DC20).Cf35
				_ = _964___mcc_h3
				var _965___mcc_h4 bool = _source16.Get_().(D9_DC20).Cf36
				_ = _965___mcc_h4
				var _966___mcc_h5 bool = _source16.Get_().(D9_DC20).Cf37
				_ = _966___mcc_h5
				var _967___mcc_h6 _dafny.Map = _source16.Get_().(D9_DC20).Cf38
				_ = _967___mcc_h6
				var _968_cf38 _dafny.Map = _967___mcc_h6
				_ = _968_cf38
				var _969_cf37 bool = _966___mcc_h5
				_ = _969_cf37
				var _970_cf36 bool = _965___mcc_h4
				_ = _970_cf36
				var _971_cf35 bool = _964___mcc_h3
				_ = _971_cf35
				var _972_cf34 _dafny.Int = _963___mcc_h2
				_ = _972_cf34
				return !(_970_cf36)
			} else {
				var _973___mcc_h7 _dafny.Set = _source16.Get_().(D9_DC17).Cf31
				_ = _973___mcc_h7
				var _974_cf31 _dafny.Set = _973___mcc_h7
				_ = _974_cf31
				return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-888))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}(func(_975_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})), _dafny.UnicodeSeqOfUtf8Bytes("im"))
			}
		}(Companion_D9_.Create_DC20_((_this).F32(), !(!(true)), true, false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-15), (_dafny.MultiSetOf((_this).F32(), (_this).F32())).Cardinality()))))
	}
}
func (_this *C2) Fm6(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-875))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg71 _dafny.Int) interface{} {
				return coer71(arg71)
			}
		}(func(_976_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(true)
		})), _dafny.SeqOf(_dafny.SeqOf(true, !(true)))), _dafny.SeqOf(_dafny.SeqOf(!(false), true)))
	}
}
func (_this *C2) Fm21(globalState *GlobalState) bool {
	{
		return !((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), true))).Cardinality()).Cmp(((_this).F31()).Cardinality()) != 0)
	}
}
func (_this *C2) Fm22(p0 _dafny.Map, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var _source17 D1 = Companion_D1_.Create_DC1_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F32()), !(true)))
		_ = _source17
		if _source17.Is_DC2() {
			var _977___mcc_h0 _dafny.CodePoint = _source17.Get_().(D1_DC2).Cf2
			_ = _977___mcc_h0
			var _978___mcc_h1 _dafny.Int = _source17.Get_().(D1_DC2).Cf3
			_ = _978___mcc_h1
			var _979_cf3 _dafny.Int = _978___mcc_h1
			_ = _979_cf3
			var _980_cf2 _dafny.CodePoint = _977___mcc_h0
			_ = _980_cf2
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rdr"), _dafny.UnicodeSeqOfUtf8Bytes("gxqk"))
		} else {
			var _981___mcc_h2 _dafny.Map = _source17.Get_().(D1_DC1).Cf1
			_ = _981___mcc_h2
			var _982_cf1 _dafny.Map = _981___mcc_h2
			_ = _982_cf1
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kgvy"), _dafny.UnicodeSeqOfUtf8Bytes("tp"))
		}
	}
}
func (_this *C2) M2(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _983_v0 _dafny.CodePoint
		_ = _983_v0
		_983_v0 = _dafny.CodePoint('i')
		_983_v0 = _983_v0
		var _984_v1 _dafny.Array
		_ = _984_v1
		var _nw158 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw158
		_984_v1 = _nw158
		var _985_v2 _dafny.Map
		_ = _985_v2
		_985_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
		var _986_v3 _dafny.Map
		_ = _986_v3
		_986_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_984_v1, _985_v2)
		_986_v3 = (_986_v3).Update(_984_v1, _985_v2)
		if p1 {
			(globalState).F26 = Companion_Default___.Fm0(p2, globalState)
			var _987_v4 _dafny.MultiSet
			_ = _987_v4
			_987_v4 = _dafny.MultiSetOf(_984_v1, _984_v1)
			if (_987_v4).IsSubsetOf(_987_v4) {
				(globalState).F24 = _dafny.IntOfInt64(308)
				var _988_v5 _dafny.Map
				_ = _988_v5
				_988_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F31())
				var _989_v6 _dafny.Map
				_ = _989_v6
				_989_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))
				_ = _index105
				(_984_v1).ArraySet1(((_this).F31()).IsDisjointFrom((func() _dafny.MultiSet {
					if (_988_v5).Contains((_989_v6).Cardinality()) {
						return (_988_v5).Get((_989_v6).Cardinality()).(_dafny.MultiSet)
					}
					return (_this).F31()
				})()), (_index105).Int())
				var _990_v7 _dafny.Sequence
				_ = _990_v7
				_990_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ojlonqco")
				var _991_v8 D4
				_ = _991_v8
				_991_v8 = Companion_D4_.Create_DC8_(_983_v0, _dafny.IntOfUint32((_990_v7).Cardinality()), p1)
				var _992_v9 _dafny.Set
				_ = _992_v9
				_992_v9 = _dafny.SetOf(_991_v8)
				var _993_v10 _dafny.Map
				_ = _993_v10
				_993_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((Companion_Default___.Fm23((_989_v6).Cardinality(), !(true), globalState)).Cardinality()))
				var _994_v11 _dafny.Sequence
				_ = _994_v11
				_994_v11 = _dafny.SeqOf((_dafny.IntOfUint32((_990_v7).Cardinality())).Minus((_992_v9).Cardinality()), (_dafny.Zero).Minus(p3), p3, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_990_v7).Cardinality()), (_993_v10).Cardinality()))
				var _995_v12 _dafny.MultiSet
				_ = _995_v12
				_995_v12 = _dafny.MultiSetOf(_993_v10)
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))
				_ = _index106
				var _rhs142 bool = !_dafny.Companion_Sequence_.Equal(_990_v7, _990_v7)
				_ = _rhs142
				var _rhs143 _dafny.Sequence = _dafny.SeqOf(p3, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if p1 {
						return _dafny.SeqOf(p2)
					}
					return _994_v11
				})()).Cardinality()))
				_ = _rhs143
				var _rhs144 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_995_v12).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_dafny.Zero).Minus(p3))) {
						return (_995_v12).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_dafny.Zero).Minus(p3)))
					}
					return Companion_Default___.SafeDivisionInt((_this).F32(), p3)
				})())
				_ = _rhs144
				var _rhs145 _dafny.Array = _984_v1
				_ = _rhs145
				var _rhs146 bool = (_this).Fm21(globalState)
				_ = _rhs146
				var _lhs102 _dafny.Array = _984_v1
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))
				_ = _lhs103
				var _lhs104 *GlobalState = globalState
				_ = _lhs104
				var _lhs105 *GlobalState = globalState
				_ = _lhs105
				(_lhs102).ArraySet1(_rhs142, (_lhs103).Int())
				_994_v11 = _rhs143
				_lhs104.F15 = _rhs144
				_984_v1 = _rhs145
				_lhs105.F26 = _rhs146
				var _996_v13 _dafny.Sequence
				_ = _996_v13
				_996_v13 = _990_v7
				var _997_v14 _dafny.Array
				_ = _997_v14
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw159
				_997_v14 = _nw159
				var _998_v15 bool
				_ = _998_v15
				_998_v15 = (_984_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))).Int()).(bool)
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_997_v14), 0))
				_ = _index107
				(_997_v14).ArraySet1(_998_v15, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_997_v14), 0))
				_ = _index108
				var _rhs147 bool = p1
				_ = _rhs147
				var _rhs148 _dafny.Sequence = _996_v13
				_ = _rhs148
				var _rhs149 bool = p1
				_ = _rhs149
				var _rhs150 bool = _998_v15
				_ = _rhs150
				var _rhs151 bool = p0
				_ = _rhs151
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				var _lhs108 _dafny.Array = _997_v14
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_997_v14), 0))
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				_lhs106.F26 = _rhs147
				_996_v13 = _rhs148
				_lhs107.F26 = _rhs149
				(_lhs108).ArraySet1(_rhs150, (_lhs109).Int())
				_lhs110.F26 = _rhs151
				var _999_v16 _dafny.Map
				_ = _999_v16
				_999_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if false {
						return _990_v7
					}
					return _990_v7
				})(), ((_this).F31()).Difference((_this).F31()))
				_999_v16 = _999_v16
				var _rhs152 bool = (_984_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))).Int()).(bool)
				_ = _rhs152
				var _rhs153 _dafny.Sequence = (_this).Fm22((_993_v10).Merge(Companion_Default___.Fm24(p3, globalState)), false, (func() _dafny.Int {
					if (_993_v10).Contains((_984_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))).Int()).(bool)) {
						return (_993_v10).Get((_984_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_984_v1), 0))).Int()).(bool)).(_dafny.Int)
					}
					return ((_this).F31()).Cardinality()
				})(), globalState)
				_ = _rhs153
				var _rhs154 bool = !(p1)
				_ = _rhs154
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_lhs111.F26 = _rhs152
				_990_v7 = _rhs153
				_lhs112.F26 = _rhs154
			} else {
				var _1000_v17 _dafny.Sequence
				_ = _1000_v17
				_1000_v17 = _dafny.UnicodeSeqOfUtf8Bytes("vkvarnrk")
				var _1001_v18 *C0
				_ = _1001_v18
				var _nw160 *C0 = New_C0_()
				_ = _nw160
				_nw160.Ctor__((_dafny.Zero).Minus(p3), _1000_v17)
				_1001_v18 = _nw160
				var _1002_v19 D6
				_ = _1002_v19
				_1002_v19 = Companion_D6_.Create_DC11_(_1001_v18)
				var _1003_v20 _dafny.Map
				_ = _1003_v20
				_1003_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1002_v19, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1000_v17, _1000_v17), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_983_v0)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1000_v17, _1000_v17)).Cardinality()))).Uint32(), _983_v0))
				_1003_v20 = (_1003_v20).Update(_1002_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gvq"), _1000_v17))
				var _1004_v21 _dafny.Map
				_ = _1004_v21
				_1004_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1005_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1006_i0 _dafny.Int) _dafny.CodePoint {
						return _1005_v0
					}
				})(_983_v0))), p0)
				var _1007_v22 *C0
				_ = _1007_v22
				var _nw161 *C0 = New_C0_()
				_ = _nw161
				_nw161.Ctor__(((_1004_v21).Cardinality()).Plus((_this).F32()), (_1001_v18).F37())
				_1007_v22 = _nw161
				var _1008_v23 _dafny.Sequence
				_ = _1008_v23
				_1008_v23 = _dafny.SeqOf(p0)
				_1008_v23 = _dafny.SeqOf(false, p1, _dafny.Companion_Sequence_.IsPrefixOf((_1001_v18).F37(), _dafny.UnicodeSeqOfUtf8Bytes("oxvupb")), p0, p1)
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_984_v1), 0))
				_ = _index109
				(_984_v1).ArraySet1(p1, (_index109).Int())
				var _1009_v24 _dafny.Set
				_ = _1009_v24
				_1009_v24 = _dafny.SetOf(_984_v1)
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_984_v1), 0))
				_ = _index110
				(_984_v1).ArraySet1((_dafny.SetOf(_984_v1)).IsDisjointFrom(_1009_v24), (_index110).Int())
				(globalState).F24 = _dafny.IntOfInt64(315)
			}
			var _1010_v25 _dafny.Sequence
			_ = _1010_v25
			_1010_v25 = _dafny.SeqOf(_984_v1)
			var _1011_v26 _dafny.Sequence
			_ = _1011_v26
			_1011_v26 = _dafny.UnicodeSeqOfUtf8Bytes("pbbsuc")
			var _1012_v27 _dafny.MultiSet
			_ = _1012_v27
			_1012_v27 = _dafny.MultiSetOf(_1011_v26, _1011_v26)
			var _rhs155 bool = ((_this).F32()).Cmp((_dafny.IntOfUint32((_1010_v25).Cardinality())).Minus(p3)) > 0
			_ = _rhs155
			var _rhs156 _dafny.MultiSet = _1012_v27
			_ = _rhs156
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			_lhs113.F26 = _rhs155
			_lhs114.F8 = _rhs156
			(globalState).F26 = !(p0)
			(globalState).F26 = ((_dafny.Zero).Minus(p3)).Cmp((_this).F32()) != 0
		} else {
			(globalState).F26 = (p2).Cmp(p2) >= 0
			(globalState).F26 = p1
			var _1013_v28 _dafny.Array
			_ = _1013_v28
			var _nwElement0_37 _dafny.Int = p2
			_ = _nwElement0_37
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_37, 0)
			(_nw162).ArraySet1(Companion_Default___.Fm1(globalState), 1)
			_1013_v28 = _nw162
			_1013_v28 = _1013_v28
			var _1014_v29 _dafny.Array
			_ = _1014_v29
			var _nwElement0_38 _dafny.Array = _984_v1
			_ = _nwElement0_38
			var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(6))
			_ = _nw163
			(_nw163).ArraySet1(_nwElement0_38, 0)
			(_nw163).ArraySet1(_984_v1, 1)
			(_nw163).ArraySet1(_984_v1, 2)
			(_nw163).ArraySet1(_984_v1, 3)
			(_nw163).ArraySet1(_984_v1, 4)
			(_nw163).ArraySet1(_984_v1, 5)
			_1014_v29 = _nw163
			_1014_v29 = _1014_v29
			var _1015_v30 _dafny.Array
			_ = _1015_v30
			var _nw164 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(13))
			_ = _nw164
			_1015_v30 = _nw164
			var _1016_v31 D8
			_ = _1016_v31
			_1016_v31 = Companion_D8_.Create_DC16_((_this).F32(), p3)
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_1015_v30), 0))
			_ = _index111
			(_1015_v30).ArraySet1(_1016_v31, (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_1015_v30), 0))
			_ = _index112
			var _rhs157 _dafny.Array = _1014_v29
			_ = _rhs157
			var _rhs158 D8 = _1016_v31
			_ = _rhs158
			var _rhs159 _dafny.Int = (_dafny.Zero).Minus(((_dafny.Zero).Minus(Companion_Default___.Fm1(globalState))).Plus((_dafny.Zero).Minus(p3)))
			_ = _rhs159
			var _lhs115 _dafny.Array = _1015_v30
			_ = _lhs115
			var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_1015_v30), 0))
			_ = _lhs116
			_1014_v29 = _rhs157
			(_lhs115).ArraySet1(_rhs158, (_lhs116).Int())
			r1 = _rhs159
		}
		var _1017_v32 _dafny.Map
		_ = _1017_v32
		_1017_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(globalState), (_this).F32())
		var _1018_v33 _dafny.Sequence
		_ = _1018_v33
		_1018_v33 = _dafny.UnicodeSeqOfUtf8Bytes("efdnwcf")
		var _1019_v34 D2
		_ = _1019_v34
		_1019_v34 = Companion_D2_.Create_DC4_((_1017_v32).Cardinality(), _dafny.IntOfInt64(343), p3, _983_v0, _1018_v33)
		var _1020_v35 _dafny.MultiSet
		_ = _1020_v35
		_1020_v35 = _dafny.MultiSetOf(Companion_D2_.Create_DC5_(_1019_v34))
		var _1021_v36 D2
		_ = _1021_v36
		_1021_v36 = Companion_D2_.Create_DC5_(_1019_v34)
		var _1022_v37 _dafny.Map
		_ = _1022_v37
		_1022_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, false)
		var _1023_v38 _dafny.Sequence
		_ = _1023_v38
		_1023_v38 = _dafny.SeqOf((func() bool {
			if (_1022_v37).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality()) {
				return (_1022_v37).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality()).(bool)
			}
			return p1
		})(), true)
		var _1024_v39 _dafny.Sequence
		_ = _1024_v39
		_1024_v39 = _dafny.SeqOf(p2)
		var _1025_v40 D10
		_ = _1025_v40
		_1025_v40 = Companion_D10_.Create_DC21_(_1024_v39)
		var _1026_v41 _dafny.Array
		_ = _1026_v41
		var _nwElement0_39 _dafny.Sequence = _1018_v33
		_ = _nwElement0_39
		var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(9))
		_ = _nw165
		(_nw165).ArraySet1(_nwElement0_39, 0)
		(_nw165).ArraySet1(_1018_v33, 1)
		(_nw165).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("n"), 2)
		(_nw165).ArraySet1(_1018_v33, 3)
		(_nw165).ArraySet1(_1018_v33, 4)
		(_nw165).ArraySet1(_1018_v33, 5)
		(_nw165).ArraySet1(_1018_v33, 6)
		(_nw165).ArraySet1(_1018_v33, 7)
		(_nw165).ArraySet1(_1018_v33, 8)
		_1026_v41 = _nw165
		var _1027_v42 *C1
		_ = _1027_v42
		var _nw166 *C1 = New_C1_()
		_ = _nw166
		_nw166.Ctor__((_dafny.MultiSetOf(_1021_v36)).IsSubsetOf(_1020_v35), (_dafny.MultiSetFromSeq(_dafny.SeqOf(p0))).Difference(_dafny.MultiSetFromSeq(_1023_v38)), (_1025_v40).Dtor_cf39(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-160), (_this).F32()), (func() _dafny.Array {
			if p0 {
				return _1026_v41
			}
			return _1026_v41
		})())
		_1027_v42 = _nw166
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_984_v1), 0))
		_ = _index113
		(_984_v1).ArraySet1(!(_dafny.SetOf(p0)).Contains(p0), (_index113).Int())
		var _pat_let_tv10 = _1027_v42
		_ = _pat_let_tv10
		var _pat_let_tv11 = _1027_v42
		_ = _pat_let_tv11
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_984_v1), 0))
		_ = _index114
		(_984_v1).ArraySet1(func(_source18 D10) bool {
			if _source18.Is_DC22() {
				var _1028___mcc_h0 _dafny.Int = _source18.Get_().(D10_DC22).Cf40
				_ = _1028___mcc_h0
				var _1029___mcc_h1 _dafny.MultiSet = _source18.Get_().(D10_DC22).Cf41
				_ = _1029___mcc_h1
				var _1030_cf41 _dafny.MultiSet = _1029___mcc_h1
				_ = _1030_cf41
				var _1031_cf40 _dafny.Int = _1028___mcc_h0
				_ = _1031_cf40
				return _pat_let_tv10.F35
			} else {
				var _1032___mcc_h2 _dafny.Sequence = _source18.Get_().(D10_DC21).Cf39
				_ = _1032___mcc_h2
				var _1033_cf39 _dafny.Sequence = _1032___mcc_h2
				_ = _1033_cf39
				return _pat_let_tv11.F35
			}
		}(_1025_v40), (_index114).Int())
		var _1034_v43 _dafny.Array
		_ = _1034_v43
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
		_ = _nw167
		_1034_v43 = _nw167
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_1034_v43), 0))
		_ = _index115
		(_1034_v43).ArraySet1(_1023_v38, (_index115).Int())
		var _1035_v44 _dafny.Set
		_ = _1035_v44
		_1035_v44 = _dafny.SetOf(true)
		var _1036_v45 _dafny.Sequence
		_ = _1036_v45
		_1036_v45 = _dafny.SeqOf(_1035_v44)
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_1034_v43), 0))
		_ = _index116
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_984_v1), 0))
		_ = _index117
		var _rhs160 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1023_v38, _1023_v38)
		_ = _rhs160
		var _rhs161 bool = (p1) == (!(false))
		_ = _rhs161
		var _rhs162 bool = !((_1035_v44).IsProperSubsetOf((_1036_v45).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F32()), _dafny.IntOfUint32((_1036_v45).Cardinality()))).Uint32()).(_dafny.Set)))
		_ = _rhs162
		var _lhs117 _dafny.Array = _1034_v43
		_ = _lhs117
		var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_1034_v43), 0))
		_ = _lhs118
		var _lhs119 _dafny.Array = _984_v1
		_ = _lhs119
		var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_984_v1), 0))
		_ = _lhs120
		var _lhs121 *GlobalState = globalState
		_ = _lhs121
		(_lhs117).ArraySet1(_rhs160, (_lhs118).Int())
		(_lhs119).ArraySet1(_rhs161, (_lhs120).Int())
		_lhs121.F26 = _rhs162
		r0 = _984_v1
		r1 = p2
		return r0, r1
	}
}
func (_this *C2) M3(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) {
	{
		var _1037_v0 _dafny.Array
		_ = _1037_v0
		var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw168
		_1037_v0 = _nw168
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
		_ = _index118
		(_1037_v0).ArraySet1((func() _dafny.Int {
			if p0 {
				return _dafny.IntOfUint32((p2).Cardinality())
			}
			return _dafny.IntOfUint32((p2).Cardinality())
		})(), (_index118).Int())
		var _1038_v1 _dafny.Map
		_ = _1038_v1
		_1038_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1039_v2 D8
		_ = _1039_v2
		_1039_v2 = Companion_D8_.Create_DC15_(_1038_v1)
		var _pat_let_tv12 = p3
		_ = _pat_let_tv12
		var _pat_let_tv13 = p3
		_ = _pat_let_tv13
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
		_ = _index119
		(_1037_v0).ArraySet1(func(_source19 D8) _dafny.Int {
			if _source19.Is_DC16() {
				var _1040___mcc_h0 _dafny.Int = _source19.Get_().(D8_DC16).Cf29
				_ = _1040___mcc_h0
				var _1041___mcc_h1 _dafny.Int = _source19.Get_().(D8_DC16).Cf30
				_ = _1041___mcc_h1
				var _1042_cf30 _dafny.Int = _1041___mcc_h1
				_ = _1042_cf30
				var _1043_cf29 _dafny.Int = _1040___mcc_h0
				_ = _1043_cf29
				return (_dafny.MultiSetOf(_dafny.SetOf(true), _dafny.SetOf(_pat_let_tv12, _pat_let_tv13))).Cardinality()
			} else {
				var _1044___mcc_h2 _dafny.Map = _source19.Get_().(D8_DC15).Cf28
				_ = _1044___mcc_h2
				var _1045_cf28 _dafny.Map = _1044___mcc_h2
				_ = _1045_cf28
				return _dafny.IntOfInt64(-118)
			}
		}(_1039_v2), (_index119).Int())
		if false {
			var _1046_v3 _dafny.Sequence
			_ = _1046_v3
			_1046_v3 = _dafny.SeqOf((_this).F32())
			var _1047_v4 _dafny.Map
			_ = _1047_v4
			_1047_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			var _1048_v5 _dafny.Map
			_ = _1048_v5
			_1048_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(680))
			var _1049_v6 _dafny.Sequence
			_ = _1049_v6
			_1049_v6 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(250))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg73 _dafny.Int) interface{} {
					return coer73(arg73)
				}
			}((func(_1050_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1051_i1 _dafny.Int) _dafny.CodePoint {
					return _1050_p1
				}
			})(p1))))
			var _1052_v7 _dafny.Array
			_ = _1052_v7
			var _nwElement0_40 _dafny.Sequence = p2
			_ = _nwElement0_40
			var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(9))
			_ = _nw169
			(_nw169).ArraySet1(_nwElement0_40, 0)
			(_nw169).ArraySet1(p2, 1)
			(_nw169).ArraySet1((func() _dafny.Sequence {
				if (_1047_v4).Contains(p1) {
					return (_1047_v4).Get(p1).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("qnpxuk")
			})(), 2)
			(_nw169).ArraySet1(p2, 3)
			(_nw169).ArraySet1((_this).Fm22(_1048_v5, p3, (_this).F32(), globalState), 4)
			(_nw169).ArraySet1((_1049_v6).Select((Companion_Default___.SafeIndex((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1049_v6).Cardinality()))).Uint32()).(_dafny.Sequence), 5)
			(_nw169).ArraySet1(p2, 6)
			(_nw169).ArraySet1(p2, 7)
			(_nw169).ArraySet1(p2, 8)
			_1052_v7 = _nw169
			var _1053_v8 *C1
			_ = _1053_v8
			var _nw170 *C1 = New_C1_()
			_ = _nw170
			_nw170.Ctor__(p0, (_this).F31(), _dafny.Companion_Sequence_.Concatenate(_1046_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}(func(_1054_i0 _dafny.Int) _dafny.Int {
				return (_this).F32()
			}))), _dafny.IntOfInt64(798), _1052_v7)
			_1053_v8 = _nw170
			var _1055_v9 _dafny.Array
			_ = _1055_v9
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_17
			var _nw171 _dafny.Array
			_ = _nw171
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw171 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = (func(_1056_v8 *C1) func(_dafny.Int) bool {
					return func(_1057_i2 _dafny.Int) bool {
						return _1056_v8.F35
					}
				})(_1053_v8)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw171 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw171).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw171).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_1055_v9 = _nw171
			var _1058_v10 _dafny.Map
			_ = _1058_v10
			_1058_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_v9, p3)
			var _1059_v11 *C0
			_ = _1059_v11
			var _nw172 *C0 = New_C0_()
			_ = _nw172
			_nw172.Ctor__((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), p2)
			_1059_v11 = _nw172
			var _1060_v12 _dafny.Map
			_ = _1060_v12
			_1060_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1053_v8.F35, _1059_v11)
			var _1061_v13 _dafny.Sequence
			_ = _1061_v13
			_1061_v13 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(186))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg75 _dafny.Int) interface{} {
					return coer75(arg75)
				}
			}((func(_1062_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1063_i4 _dafny.Int) _dafny.CodePoint {
					return _1062_p1
				}
			})(p1)))
			var _1064_v14 _dafny.Array
			_ = _1064_v14
			var _nwElement0_41 *C0 = _1059_v11
			_ = _nwElement0_41
			var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(22))
			_ = _nw173
			(_nw173).ArraySet1(_nwElement0_41, 0)
			(_nw173).ArraySet1(_1059_v11, 1)
			(_nw173).ArraySet1(_1059_v11, 2)
			(_nw173).ArraySet1(_1059_v11, 3)
			(_nw173).ArraySet1(_1059_v11, 4)
			(_nw173).ArraySet1(_1059_v11, 5)
			(_nw173).ArraySet1(_1059_v11, 6)
			(_nw173).ArraySet1(_1059_v11, 7)
			(_nw173).ArraySet1(_1059_v11, 8)
			(_nw173).ArraySet1(_1059_v11, 9)
			(_nw173).ArraySet1(_1059_v11, 10)
			(_nw173).ArraySet1(_1059_v11, 11)
			(_nw173).ArraySet1(_1059_v11, 12)
			(_nw173).ArraySet1(_1059_v11, 13)
			(_nw173).ArraySet1(_1059_v11, 14)
			(_nw173).ArraySet1(_1059_v11, 15)
			(_nw173).ArraySet1(_1059_v11, 16)
			(_nw173).ArraySet1(_1059_v11, 17)
			(_nw173).ArraySet1(_1059_v11, 18)
			(_nw173).ArraySet1(_1059_v11, 19)
			(_nw173).ArraySet1(_1059_v11, 20)
			(_nw173).ArraySet1((func() *C0 {
				if (_1060_v12).Contains((_1053_v8).Fm11(p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-262))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1065_v11 *C0) func(_dafny.Int) _dafny.Int {
					return func(_1066_i3 _dafny.Int) _dafny.Int {
						return (_1065_v11).F36()
					}
				})(_1059_v11)))).Cardinality()), _1061_v13, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), globalState)) {
					return (_1060_v12).Get((_1053_v8).Fm11(p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-262))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg77 _dafny.Int) interface{} {
							return coer77(arg77)
						}
					}((func(_1067_v11 *C0) func(_dafny.Int) _dafny.Int {
						return func(_1068_i3 _dafny.Int) _dafny.Int {
							return (_1067_v11).F36()
						}
					})(_1059_v11)))).Cardinality()), _1061_v13, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), globalState)).(*C0)
				}
				return _1059_v11
			})(), 21)
			_1064_v14 = _nw173
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index120
			(_1064_v14).ArraySet1(_1059_v11, (_index120).Int())
			var _1069_v15 _dafny.Sequence
			_ = _1069_v15
			_1069_v15 = _dafny.UnicodeSeqOfUtf8Bytes("vhmyrad")
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index121
			var _rhs163 _dafny.Map = _1058_v10
			_ = _rhs163
			var _rhs164 *C0 = _1059_v11
			_ = _rhs164
			var _rhs165 _dafny.Sequence = p2
			_ = _rhs165
			var _lhs122 _dafny.Array = _1064_v14
			_ = _lhs122
			var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1064_v14), 0))
			_ = _lhs123
			_1058_v10 = _rhs163
			(_lhs122).ArraySet1(_rhs164, (_lhs123).Int())
			_1069_v15 = _rhs165
			(globalState).F26 = ((_1059_v11).F36()).Cmp((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)) != 0
			(globalState).F26 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}((func(_1070_v11 *C0) func(_dafny.Int) _dafny.Int {
				return func(_1071_i5 _dafny.Int) _dafny.Int {
					return (_1070_v11).F36()
				}
			})(_1059_v11))), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-20), _dafny.IntOfUint32((_1046_v3).Cardinality())))
			_1069_v15 = _1069_v15
		} else {
			_1037_v0 = _1037_v0
			var _1072_v16 _dafny.Sequence
			_ = _1072_v16
			_1072_v16 = _dafny.SeqOf((_this).F32(), (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm25(p0, globalState)).Cardinality())), (_this).F32())
			var _1073_v17 _dafny.Array
			_ = _1073_v17
			var _nwElement0_42 _dafny.Sequence = p2
			_ = _nwElement0_42
			var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(2))
			_ = _nw174
			(_nw174).ArraySet1(_nwElement0_42, 0)
			(_nw174).ArraySet1(p2, 1)
			_1073_v17 = _nw174
			var _1074_v18 *C1
			_ = _1074_v18
			var _nw175 *C1 = New_C1_()
			_ = _nw175
			_nw175.Ctor__(p3, (_this).F31(), _1072_v16, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1073_v17)
			_1074_v18 = _nw175
			var _1075_v19 D4
			_ = _1075_v19
			_1075_v19 = Companion_D4_.Create_DC8_(_dafny.CodePoint('x'), (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), p0)
			var _1076_v20 _dafny.Sequence
			_ = _1076_v20
			_1076_v20 = p2
			var _1077_v21 _dafny.Map
			_ = _1077_v21
			_1077_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1075_v19, _1076_v20)
			if !(_1077_v21).Contains(_1075_v19) {
				var _1078_v22 *C1
				_ = _1078_v22
				var _nw176 *C1 = New_C1_()
				_ = _nw176
				_nw176.Ctor__(((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-977)) == 0, _dafny.MultiSetOf(p3, !(p0)), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1072_v16, _1072_v16), (Companion_Default___.SafeIndex((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1072_v16, _1072_v16)).Cardinality()))).Uint32(), _dafny.IntOfUint32((p2).Cardinality())), (_this).F32(), _1073_v17)
				_1078_v22 = _nw176
				var _1079_v23 _dafny.Set
				_ = _1079_v23
				_1079_v23 = _dafny.SetOf(p3)
				var _1080_v24 _dafny.MultiSet
				_ = _1080_v24
				_1080_v24 = _dafny.MultiSetOf((_1079_v23).Cardinality(), (_this).F32(), _dafny.IntOfInt64(419))
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
				_ = _index122
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
				_ = _index123
				var _rhs166 _dafny.Int = (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)
				_ = _rhs166
				var _rhs167 _dafny.Int = ((func() _dafny.Int {
					if p3 {
						return (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)
					}
					return (func() _dafny.Set {
						var _coll37 = _dafny.NewBuilder()
						_ = _coll37
						for _iter39 := _dafny.Iterate((func() _dafny.Set {
							var _coll38 = _dafny.NewBuilder()
							_ = _coll38
							for _iter40 := _dafny.Iterate((_1080_v24).Elements()); ; {
								_compr_38, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _1081_v25 _dafny.Int
								_1081_v25 = interface{}(_compr_38).(_dafny.Int)
								if (_1080_v24).Contains(_1081_v25) {
									_coll38.Add((_1081_v25).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
										var _coll39 = _dafny.NewBuilder()
										_ = _coll39
										for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-421), _dafny.IntOfInt64(239))); ; {
											_compr_39, _ok41 := _iter41()
											if !_ok41 {
												break
											}
											var _1082_v26 _dafny.Int
											_1082_v26 = interface{}(_compr_39).(_dafny.Int)
											if ((_dafny.IntOfInt64(-421)).Cmp(_1082_v26) <= 0) && ((_1082_v26).Cmp(_dafny.IntOfInt64(239)) < 0) {
												_coll39.Add((_1082_v26).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("pdfc"))).Cardinality()))
											}
										}
										return _coll39.ToSet()
									}()).Cardinality(), _dafny.IntOfInt64(543))).Cardinality()))))
								}
							}
							return _coll38.ToSet()
						}()).Elements()); ; {
							_compr_37, _ok39 := _iter39()
							if !_ok39 {
								break
							}
							var _1083_v27 _dafny.Int
							_1083_v27 = interface{}(_compr_37).(_dafny.Int)
							if (func() _dafny.Set {
								var _coll40 = _dafny.NewBuilder()
								_ = _coll40
								for _iter42 := _dafny.Iterate((_1080_v24).Elements()); ; {
									_compr_40, _ok42 := _iter42()
									if !_ok42 {
										break
									}
									var _1084_v25 _dafny.Int
									_1084_v25 = interface{}(_compr_40).(_dafny.Int)
									if (_1080_v24).Contains(_1084_v25) {
										_coll40.Add((_1084_v25).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
											var _coll41 = _dafny.NewBuilder()
											_ = _coll41
											for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-421), _dafny.IntOfInt64(239))); ; {
												_compr_41, _ok43 := _iter43()
												if !_ok43 {
													break
												}
												var _1085_v26 _dafny.Int
												_1085_v26 = interface{}(_compr_41).(_dafny.Int)
												if ((_dafny.IntOfInt64(-421)).Cmp(_1085_v26) <= 0) && ((_1085_v26).Cmp(_dafny.IntOfInt64(239)) < 0) {
													_coll41.Add((_1085_v26).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("pdfc"))).Cardinality()))
												}
											}
											return _coll41.ToSet()
										}()).Cardinality(), _dafny.IntOfInt64(543))).Cardinality()))))
									}
								}
								return _coll40.ToSet()
							}()).Contains(_1083_v27) {
								_coll37.Add(Companion_Default___.SafeModuloInt(_1083_v27, _dafny.IntOfInt64(891)))
							}
						}
						return _coll37.ToSet()
					}()).Cardinality()
				})()).Plus((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))
				_ = _rhs167
				var _rhs168 _dafny.Int = _dafny.IntOfInt64(135)
				_ = _rhs168
				var _rhs169 _dafny.Int = (_1078_v22).Fm4(_dafny.SetOf(p2), p0, _dafny.SeqOf(p0, _1078_v22.F35), globalState)
				_ = _rhs169
				var _rhs170 _dafny.Int = (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)
				_ = _rhs170
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				var _lhs125 _dafny.Array = _1037_v0
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
				_ = _lhs126
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				var _lhs129 _dafny.Array = _1037_v0
				_ = _lhs129
				var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
				_ = _lhs130
				_lhs124.F3 = _rhs166
				(_lhs125).ArraySet1(_rhs167, (_lhs126).Int())
				_lhs127.F18 = _rhs168
				_lhs128.F24 = _rhs169
				(_lhs129).ArraySet1(_rhs170, (_lhs130).Int())
				var _1086_v28 _dafny.Array
				_ = _1086_v28
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw177
				_1086_v28 = _nw177
				var _1087_v29 _dafny.Sequence
				_ = _1087_v29
				_1087_v29 = _dafny.SeqOf(_1074_v18)
				var _1088_v30 _dafny.Array
				_ = _1088_v30
				var _nwElement0_43 *C1 = _1078_v22
				_ = _nwElement0_43
				var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(24))
				_ = _nw178
				(_nw178).ArraySet1(_nwElement0_43, 0)
				(_nw178).ArraySet1(_1074_v18, 1)
				(_nw178).ArraySet1(_1074_v18, 2)
				(_nw178).ArraySet1(_1074_v18, 3)
				(_nw178).ArraySet1(_1078_v22, 4)
				(_nw178).ArraySet1(_1078_v22, 5)
				(_nw178).ArraySet1(_1078_v22, 6)
				(_nw178).ArraySet1(_1078_v22, 7)
				(_nw178).ArraySet1(_1074_v18, 8)
				(_nw178).ArraySet1(_1078_v22, 9)
				(_nw178).ArraySet1(_1078_v22, 10)
				(_nw178).ArraySet1(_1078_v22, 11)
				(_nw178).ArraySet1(_1078_v22, 12)
				(_nw178).ArraySet1(_1074_v18, 13)
				(_nw178).ArraySet1(_1078_v22, 14)
				(_nw178).ArraySet1(_1074_v18, 15)
				(_nw178).ArraySet1(_1074_v18, 16)
				(_nw178).ArraySet1((_1087_v29).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_1087_v29).Cardinality()))).Uint32()).(*C1), 17)
				(_nw178).ArraySet1(_1074_v18, 18)
				(_nw178).ArraySet1(_1078_v22, 19)
				(_nw178).ArraySet1(_1078_v22, 20)
				(_nw178).ArraySet1(_1074_v18, 21)
				(_nw178).ArraySet1(_1078_v22, 22)
				(_nw178).ArraySet1(_1074_v18, 23)
				_1088_v30 = _nw178
				var _1089_v31 _dafny.Set
				_ = _1089_v31
				_1089_v31 = _dafny.SetOf(_1088_v30)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1086_v28), 0))
				_ = _index124
				(_1086_v28).ArraySet1((_1089_v31).Difference(_1089_v31), (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1086_v28), 0))
				_ = _index125
				(_1086_v28).ArraySet1((func() _dafny.Set {
					if _1074_v18.F35 {
						return _dafny.SetOf(_1088_v30)
					}
					return _dafny.SetOf(_1088_v30)
				})(), (_index125).Int())
				var _1090_v32 _dafny.Sequence
				_ = _1090_v32
				_1090_v32 = _dafny.SeqOf(!(p3), ((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-645)) <= 0, p0, _1078_v22.F35, (_this).Fm5(globalState))
				_1090_v32 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1074_v18.F35, _1074_v18.F35), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p3), (Companion_Default___.SafeIndex((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality()))).Uint32(), p0)), _1090_v32)
				var _1091_v33 _dafny.Array
				_ = _1091_v33
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_18
				var _nw179 _dafny.Array
				_ = _nw179
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw179 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) bool = (func(_1092_p3 bool) func(_dafny.Int) bool {
						return func(_1093_i6 _dafny.Int) bool {
							return _1092_p3
						}
					})(p3)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw179 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw179).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw179).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_1091_v33 = _nw179
				var _1094_v34 D4
				_ = _1094_v34
				_1094_v34 = Companion_D4_.Create_DC7_(_1091_v33)
				var _1095_v35 _dafny.Map
				_ = _1095_v35
				_1095_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _1091_v33)
				var _1096_v36 _dafny.Array
				_ = _1096_v36
				var _nwElement0_44 _dafny.Array = (func() _dafny.Array {
					if false {
						return _1091_v33
					}
					return _1091_v33
				})()
				_ = _nwElement0_44
				var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(6))
				_ = _nw180
				(_nw180).ArraySet1(_nwElement0_44, 0)
				(_nw180).ArraySet1(_1091_v33, 1)
				(_nw180).ArraySet1(_1091_v33, 2)
				(_nw180).ArraySet1((_1094_v34).Dtor_cf12(), 3)
				(_nw180).ArraySet1((func() _dafny.Array {
					if (_1095_v35).Contains((_this).F32()) {
						return (_1095_v35).Get((_this).F32()).(_dafny.Array)
					}
					return _1091_v33
				})(), 4)
				(_nw180).ArraySet1(_1091_v33, 5)
				_1096_v36 = _nw180
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1096_v36), 0))
				_ = _index126
				(_1096_v36).ArraySet1(_1091_v33, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1096_v36), 0))
				_ = _index127
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw181
				(_1096_v36).ArraySet1(_nw181, (_index127).Int())
			} else {
				var _1097_v37 T1
				_ = _1097_v37
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__(false, (_this).F31(), _1072_v16, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1073_v17)
				_1097_v37 = _nw182
				var _1098_v38 _dafny.Map
				_ = _1098_v38
				_1098_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1097_v37, p0)
				_1098_v38 = (_1098_v38).Update(_1097_v37, (((_1097_v37).F30()).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32(((_1097_v37).F30()).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-231)) > 0)
				var _1099_v39 _dafny.Map
				_ = _1099_v39
				_1099_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), p3)
				var _1100_v40 *C1
				_ = _1100_v40
				var _nw183 *C1 = New_C1_()
				_ = _nw183
				_nw183.Ctor__(p0, (_this).F31(), _dafny.SeqOf((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), (Companion_D2_.Create_DC4_((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_1099_v39).Cardinality()), (_this).F32(), p1, p2)).Dtor_cf6()), ((_dafny.MultiSetOf(p0)).Cardinality()).Times((_1097_v37).F27()), (_1097_v37).F28())
				_1100_v40 = _nw183
				var _1101_v41 _dafny.Sequence
				_ = _1101_v41
				_1101_v41 = _dafny.SeqOf(p0)
				var _1102_v42 _dafny.Map
				_ = _1102_v42
				_1102_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1074_v18.F35, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))
				var _1103_v43 *C0
				_ = _1103_v43
				var _nw184 *C0 = New_C0_()
				_ = _nw184
				_nw184.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1101_v41).Cardinality()), _dafny.IntOfUint32(((_this).Fm22(_1102_v42, _1074_v18.F35, Companion_Default___.Fm1(globalState), globalState)).Cardinality()))).Cardinality()), p2)
				_1103_v43 = _nw184
				(globalState).F26 = !((_this).Fm5(globalState)) || (!((_1074_v18.F35) || (_1074_v18.F35)))
				(_1074_v18).F35 = p3
			}
			var _1104_v44 _dafny.Map
			_ = _1104_v44
			_1104_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm26(p1, (_this).F32(), (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), globalState), (_this).F32())
			(globalState).F18 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1104_v44).Contains(p1) {
					return (_1104_v44).Get(p1).(_dafny.Int)
				}
				return (_this).F32()
			})(), (_this).F32())
			var _1105_v45 *C0
			_ = _1105_v45
			var _nw185 *C0 = New_C0_()
			_ = _nw185
			_nw185.Ctor__(_dafny.IntOfUint32((_1072_v16).Cardinality()), p2)
			_1105_v45 = _nw185
		}
		var _hi4 _dafny.Int = _dafny.IntOfInt64(220)
		_ = _hi4
		for _1106_i7 := (_dafny.Zero).Minus((_this).F32()); _1106_i7.Cmp(_hi4) < 0; _1106_i7 = _1106_i7.Plus(_dafny.One) {
			if (_1106_i7).Cmp(_1106_i7) < 0 {
				var _1107_v46 *C0
				_ = _1107_v46
				var _nw186 *C0 = New_C0_()
				_ = _nw186
				_nw186.Ctor__((_this).F32(), p2)
				_1107_v46 = _nw186
				var _1108_v47 _dafny.Set
				_ = _1108_v47
				_1108_v47 = _dafny.SetOf(_1107_v46, _1107_v46, _1107_v46)
				(globalState).F26 = (_1108_v47).IsProperSubsetOf(_1108_v47)
				var _1109_v48 _dafny.CodePoint
				_ = _1109_v48
				_1109_v48 = _dafny.CodePoint('k')
				var _1110_v49 _dafny.Sequence
				_ = _1110_v49
				_1110_v49 = _dafny.SeqOf(p3)
				var _rhs171 bool = !(Companion_Default___.Fm0((func() _dafny.Int {
					if p3 {
						return _dafny.IntOfUint32((_1110_v49).Cardinality())
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg79 _dafny.Int) interface{} {
							return coer79(arg79)
						}
					}(func(_1111_i8 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality())
				})(), globalState))
				_ = _rhs171
				var _rhs172 _dafny.Int = ((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)).Minus(Companion_Default___.SafeModuloInt((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), (_this).F32()))
				_ = _rhs172
				var _rhs173 _dafny.Int = ((_1107_v46).F36()).Minus(_1106_i7)
				_ = _rhs173
				var _rhs174 _dafny.CodePoint = p1
				_ = _rhs174
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				_lhs131.F26 = _rhs171
				_lhs132.F15 = _rhs172
				_lhs133.F18 = _rhs173
				_1109_v48 = _rhs174
				var _1112_v50 _dafny.Map
				_ = _1112_v50
				_1112_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1107_v46).F36(), p0)
				var _1113_v51 _dafny.Array
				_ = _1113_v51
				var _nwElement0_45 _dafny.Sequence = (_1107_v46).F37()
				_ = _nwElement0_45
				var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(23))
				_ = _nw187
				(_nw187).ArraySet1(_nwElement0_45, 0)
				(_nw187).ArraySet1((_1107_v46).F37(), 1)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex((_1107_v46).F36(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1112_v50).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex((_1107_v46).F36(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), p1), 2)
				(_nw187).ArraySet1((_1107_v46).F37(), 3)
				(_nw187).ArraySet1((_this).Fm22(Companion_Default___.Fm24((_this).F32(), globalState), p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngpnnpqao")).Cardinality()), globalState), 4)
				(_nw187).ArraySet1(p2, 5)
				(_nw187).ArraySet1(p2, 6)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, p2), 7)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vkphmjmvi"), p2), 8)
				(_nw187).ArraySet1(p2, 9)
				(_nw187).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tmodjtc"), 10)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("baincv"), (_1107_v46).F37()), 11)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hehm"), _dafny.UnicodeSeqOfUtf8Bytes("fgcyfxm")), 12)
				(_nw187).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, p2), 13)
				(_nw187).ArraySet1((_1107_v46).F37(), 14)
				(_nw187).ArraySet1(p2, 15)
				(_nw187).ArraySet1((_1107_v46).F37(), 16)
				(_nw187).ArraySet1((_1107_v46).F37(), 17)
				(_nw187).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bwkywvb"), 18)
				(_nw187).ArraySet1((_1107_v46).F37(), 19)
				(_nw187).ArraySet1(p2, 20)
				(_nw187).ArraySet1(p2, 21)
				(_nw187).ArraySet1(p2, 22)
				_1113_v51 = _nw187
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_1113_v51), 0))
				_ = _index128
				(_1113_v51).ArraySet1((_1107_v46).F37(), (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_1113_v51), 0))
				_ = _index129
				(_1113_v51).ArraySet1((_1107_v46).F37(), (_index129).Int())
				var _1114_v53 _dafny.MultiSet
				_ = _1114_v53
				_1114_v53 = _dafny.MultiSetOf(func() _dafny.Set {
					var _coll42 = _dafny.NewBuilder()
					_ = _coll42
					for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(274), _dafny.IntOfInt64(326))); ; {
						_compr_42, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _1115_v52 _dafny.Int
						_1115_v52 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(274)).Cmp(_1115_v52) <= 0) && ((_1115_v52).Cmp(_dafny.IntOfInt64(326)) < 0) {
							_coll42.Add((_1115_v52).Times((_1038_v1).Cardinality()))
						}
					}
					return _coll42.ToSet()
				}(), _dafny.SetOf((func() _dafny.Int {
					if ((_this).F31()).Contains(p3) {
						return ((_this).F31()).Multiplicity(p3)
					}
					return (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)
				})(), (_1107_v46).F36(), (_1107_v46).F36(), (_1107_v46).F36(), (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)))
				var _1116_v54 D11
				_ = _1116_v54
				_1116_v54 = Companion_D11_.Create_DC23_(_1114_v53)
				_1112_v50 = (_1112_v50).Update(((_1107_v46).F36()).Plus(((_1116_v54).Dtor_cf42()).Cardinality()), !(p0))
				var _1117_v55 _dafny.Array
				_ = _1117_v55
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw188
				_1117_v55 = _nw188
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1117_v55), 0))
				_ = _index130
				(_1117_v55).ArraySet1(!(true), (_index130).Int())
				var _1118_v56 D2
				_ = _1118_v56
				_1118_v56 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)), (_1107_v46).F36(), _dafny.IntOfInt64(986), _dafny.CodePoint('y'), (_1107_v46).F37())
				var _1119_v57 _dafny.Map
				_ = _1119_v57
				_1119_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v56, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)).Cardinality())
				var _1120_v58 _dafny.Map
				_ = _1120_v58
				_1120_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1119_v57, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1106_i7))
				var _1121_v59 bool
				_ = _1121_v59
				_1121_v59 = true
				var _1122_v60 _dafny.Map
				_ = _1122_v60
				_1122_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1121_v59)
				var _1123_v61 _dafny.Sequence
				_ = _1123_v61
				_1123_v61 = _dafny.SeqOf(_1107_v46)
				var _1124_v62 _dafny.Map
				_ = _1124_v62
				_1124_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1122_v60, _dafny.SeqOf(_1107_v46, (_1123_v61).Select((Companion_Default___.SafeIndex((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1123_v61).Cardinality()))).Uint32()).(*C0)))
				var _1125_v63 _dafny.Sequence
				_ = _1125_v63
				_1125_v63 = _dafny.SeqOf(_1124_v62)
				var _1126_v64 _dafny.Sequence
				_ = _1126_v64
				_1126_v64 = _dafny.SeqOf((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))
				var _1127_v65 _dafny.Map
				_ = _1127_v65
				_1127_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1125_v63).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_1125_v63).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1126_v64).Cardinality())))
				var _1128_v66 _dafny.Map
				_ = _1128_v66
				_1128_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F32())
				var _1129_v67 _dafny.Array
				_ = _1129_v67
				var _nwElement0_46 _dafny.Map = _1120_v58
				_ = _nwElement0_46
				var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(12))
				_ = _nw189
				(_nw189).ArraySet1(_nwElement0_46, 0)
				(_nw189).ArraySet1(_1120_v58, 1)
				(_nw189).ArraySet1((_1120_v58).Merge(_1120_v58), 2)
				(_nw189).ArraySet1(_1120_v58, 3)
				(_nw189).ArraySet1(_1120_v58, 4)
				(_nw189).ArraySet1(_1120_v58, 5)
				(_nw189).ArraySet1((_1120_v58).Merge(_1120_v58), 6)
				(_nw189).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v56, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))).Update(p0, _1106_i7))).Update(_1119_v57, (func() _dafny.Map {
					if (_1127_v65).Contains((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)) {
						return (_1127_v65).Get((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)).(_dafny.Map)
					}
					return _1128_v66
				})()), 7)
				(_nw189).ArraySet1(_1120_v58, 8)
				(_nw189).ArraySet1(_1120_v58, 9)
				(_nw189).ArraySet1(_1120_v58, 10)
				(_nw189).ArraySet1((_1120_v58).Merge(_1120_v58), 11)
				_1129_v67 = _nw189
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1129_v67), 0))
				_ = _index131
				(_1129_v67).ArraySet1(_1120_v58, (_index131).Int())
				var _1130_v68 D4
				_ = _1130_v68
				_1130_v68 = Companion_D4_.Create_DC8_(p1, _dafny.IntOfInt64(802), false)
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_1113_v51), 0))
				_ = _index132
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1117_v55), 0))
				_ = _index133
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1129_v67), 0))
				_ = _index134
				var _rhs175 _dafny.Int = (Companion_D11_.Create_DC24_((func() _dafny.Int {
					if (_1128_v66).Contains(p0) {
						return (_1128_v66).Get(p0).(_dafny.Int)
					}
					return (_1107_v46).F36()
				})(), _1130_v68)).Dtor_cf43()
				_ = _rhs175
				var _rhs176 _dafny.Sequence = (_1107_v46).F37()
				_ = _rhs176
				var _rhs177 bool = true
				_ = _rhs177
				var _rhs178 _dafny.Int = (_1107_v46).F36()
				_ = _rhs178
				var _rhs179 _dafny.Map = _1120_v58
				_ = _rhs179
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 _dafny.Array = _1113_v51
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_1113_v51), 0))
				_ = _lhs136
				var _lhs137 _dafny.Array = _1117_v55
				_ = _lhs137
				var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1117_v55), 0))
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				var _lhs140 _dafny.Array = _1129_v67
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1129_v67), 0))
				_ = _lhs141
				_lhs134.F24 = _rhs175
				(_lhs135).ArraySet1(_rhs176, (_lhs136).Int())
				(_lhs137).ArraySet1(_rhs177, (_lhs138).Int())
				_lhs139.F21 = _rhs178
				(_lhs140).ArraySet1(_rhs179, (_lhs141).Int())
			} else {
				var _1131_v70 _dafny.Sequence
				_ = _1131_v70
				_1131_v70 = _dafny.SeqOf((_this).F32())
				(globalState).F24 = (func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter45 := _dafny.Iterate((_1131_v70).Elements()); ; {
						_compr_43, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _1132_v69 _dafny.Int
						_1132_v69 = interface{}(_compr_43).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1131_v70, _1132_v69) {
							_coll43.Add(Companion_Default___.SafeDivisionInt(_1132_v69, _1106_i7), p1)
						}
					}
					return _coll43.ToMap()
				}()).Cardinality()
				var _1133_v71 _dafny.Map
				_ = _1133_v71
				_1133_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _1134_v72 D10
				_ = _1134_v72
				_1134_v72 = Companion_D10_.Create_DC22_(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1133_v71).Contains(p3) {
						return (_1133_v71).Get(p3).(_dafny.Sequence)
					}
					return p2
				})()).Cardinality()), _dafny.MultiSetOf(p0))
				(_this).M7(_1038_v1, (_1134_v72).Dtor_cf40(), globalState)
				(globalState).F26 = !(((_this).F32()).Cmp(Companion_Default___.Fm1(globalState)) < 0)
				var _1135_v73 _dafny.Array
				_ = _1135_v73
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_19
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = (func(_1136_p0 bool) func(_dafny.Int) bool {
						return func(_1137_i9 _dafny.Int) bool {
							return !(!(true)) || (_1136_p0)
						}
					})(p0)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw190 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw190).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw190).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_1135_v73 = _nw190
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_1135_v73), 0))
				_ = _index135
				(_1135_v73).ArraySet1(p3, (_index135).Int())
				var _1138_v74 _dafny.Map
				_ = _1138_v74
				_1138_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))
				var _1139_v75 _dafny.Sequence
				_ = _1139_v75
				_1139_v75 = _dafny.SeqOf(p0)
				var _1140_v76 _dafny.Sequence
				_ = _1140_v76
				_1140_v76 = _dafny.SeqOf(_1139_v75, _1139_v75)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_1135_v73), 0))
				_ = _index136
				(_1135_v73).ArraySet1((Companion_Default___.Fm27(_dafny.IntOfUint32((_1140_v76).Cardinality()), globalState)).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_1138_v74).Cardinality()))), (_index136).Int())
				var _1141_v77 _dafny.Map
				_ = _1141_v77
				_1141_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), p3)
				var _1142_v78 D1
				_ = _1142_v78
				_1142_v78 = Companion_D1_.Create_DC1_(_1141_v77)
				var _1143_v79 _dafny.Sequence
				_ = _1143_v79
				_1143_v79 = _dafny.UnicodeSeqOfUtf8Bytes("tqpsjmtb")
				var _1144_v80 _dafny.Sequence
				_ = _1144_v80
				_1144_v80 = _dafny.SeqOf(_1135_v73, _1135_v73)
				var _1145_v81 _dafny.MultiSet
				_ = _1145_v81
				_1145_v81 = _dafny.MultiSetOf(_1106_i7, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1106_i7, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), (_this).F32())
				var _pat_let_tv14 = _1037_v0
				_ = _pat_let_tv14
				var _pat_let_tv15 = _1037_v0
				_ = _pat_let_tv15
				var _pat_let_tv16 = _1135_v73
				_ = _pat_let_tv16
				var _pat_let_tv17 = _1135_v73
				_ = _pat_let_tv17
				var _pat_let_tv18 = _1141_v77
				_ = _pat_let_tv18
				var _rhs180 _dafny.Array = (_1144_v80).Select((Companion_Default___.SafeIndex((_1145_v81).Cardinality(), _dafny.IntOfUint32((_1144_v80).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs180
				var _rhs181 D1 = (func() D1 {
					if p0 {
						return func(_pat_let18_0 D1) D1 {
							return func(_1146_dt__update__tmp_h0 D1) D1 {
								return func(_pat_let19_0 _dafny.Map) D1 {
									return func(_1147_dt__update_hcf1_h0 _dafny.Map) D1 {
										return Companion_D1_.Create_DC1_(_1147_dt__update_hcf1_h0)
									}(_pat_let19_0)
								}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_pat_let_tv14), 0))).Int()).(_dafny.Int), (_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(bool)))
							}(_pat_let18_0)
						}(_1142_v78)
					}
					return func(_pat_let20_0 D1) D1 {
						return func(_1148_dt__update__tmp_h1 D1) D1 {
							return func(_pat_let21_0 _dafny.Map) D1 {
								return func(_1149_dt__update_hcf1_h1 _dafny.Map) D1 {
									return Companion_D1_.Create_DC1_(_1149_dt__update_hcf1_h1)
								}(_pat_let21_0)
							}(_pat_let_tv18)
						}(_pat_let20_0)
					}(Companion_Default___.Fm28(_dafny.CodePoint('k'), (_dafny.Zero).Minus((_dafny.Zero).Minus(_1106_i7)), (_1135_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_1135_v73), 0))).Int()).(bool), (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), globalState))
				})()
				_ = _rhs181
				var _rhs182 _dafny.Sequence = _1143_v79
				_ = _rhs182
				_1135_v73 = _rhs180
				_1142_v78 = _rhs181
				_1143_v79 = _rhs182
			}
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
			_ = _index137
			(_1037_v0).ArraySet1(Companion_Default___.SafeModuloInt(_1106_i7, _dafny.IntOfInt64(408)), (_index137).Int())
			(globalState).F26 = p3
			if p0 {
				var _1150_v82 _dafny.Sequence
				_ = _1150_v82
				_1150_v82 = _dafny.SeqOf(_1106_i7)
				(globalState).F18 = ((_this).F32()).Plus((_1150_v82).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_1150_v82).Cardinality()))).Uint32()).(_dafny.Int))
				(globalState).F3 = _1106_i7
				var _1151_v83 _dafny.Map
				_ = _1151_v83
				_1151_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1106_i7)
				var _1152_v84 _dafny.Set
				_ = _1152_v84
				_1152_v84 = _dafny.SetOf(p3)
				(globalState).F0 = Companion_Default___.SafeModuloInt((_this).F32(), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).Fm22(_1151_v83, p0, (_1152_v84).Cardinality(), globalState)).Cardinality())), _1106_i7))
				(globalState).F26 = !(p3) || (true)
				var _1153_v85 *C0
				_ = _1153_v85
				var _nw191 *C0 = New_C0_()
				_ = _nw191
				_nw191.Ctor__(Companion_Default___.SafeDivisionInt((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1106_i7), _dafny.Companion_Sequence_.Concatenate((_this).Fm22(_1151_v83, true, (_this).F32(), globalState), (_this).Fm22(_1151_v83, p0, _dafny.IntOfInt64(306), globalState)))
				_1153_v85 = _nw191
			} else {
				(globalState).F26 = p0
				var _1154_v86 _dafny.Array
				_ = _1154_v86
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw192
				_1154_v86 = _nw192
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1154_v86), 0))
				_ = _index138
				(_1154_v86).ArraySet1(p0, (_index138).Int())
				var _1155_v87 _dafny.Sequence
				_ = _1155_v87
				_1155_v87 = _dafny.SeqOf(!(p0))
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1154_v86), 0))
				_ = _index139
				(_1154_v86).ArraySet1((_1155_v87).Select((Companion_Default___.SafeIndex(_1106_i7, _dafny.IntOfUint32((_1155_v87).Cardinality()))).Uint32()).(bool), (_index139).Int())
				var _1156_v88 _dafny.MultiSet
				_ = _1156_v88
				_1156_v88 = _dafny.MultiSetOf(p2)
				(globalState).F8 = _1156_v88
				(globalState).F26 = p0
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1154_v86), 0))
				_ = _index140
				(_1154_v86).ArraySet1(p0, (_index140).Int())
			}
		}
		var _1157_v89 _dafny.CodePoint
		_ = _1157_v89
		_1157_v89 = _dafny.CodePoint('g')
		_1157_v89 = p1
		var _ingredients0 = _dafny.NewBuilder()
		_ = _ingredients0
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1037_v0), 0))); ; {
			_guard_loop_2, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _1158_i10 _dafny.Int
			_1158_i10 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1158_i10).Sign() != -1) && ((_1158_i10).Cmp(_dafny.ArrayLen((_1037_v0), 0)) < 0)) {
				_ingredients0.Add(_dafny.TupleOf(_1037_v0, (_1158_i10).Int(), (_1158_i10).Times((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int))))
			}
		}
		for _iter47 := _dafny.Iterate(_ingredients0); ; {
			_tup0, _ok47 := _iter47()
			if !_ok47 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
		}
		if !(p0) || (!(p3)) {
			(globalState).F15 = (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int)
			var _1159_v90 _dafny.Sequence
			_ = _1159_v90
			_1159_v90 = _dafny.SeqOf(_dafny.IntOfUint32((p2).Cardinality()), (_this).F32())
			var _1160_v91 _dafny.Array
			_ = _1160_v91
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw193
			_1160_v91 = _nw193
			var _1161_v92 *C1
			_ = _1161_v92
			var _nw194 *C1 = New_C1_()
			_ = _nw194
			_nw194.Ctor__(p3, (_this).F31(), _1159_v90, (_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), _1160_v91)
			_1161_v92 = _nw194
			var _1162_v93 _dafny.Array
			_ = _1162_v93
			var _nwElement0_47 *C1 = _1161_v92
			_ = _nwElement0_47
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(5))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_47, 0)
			(_nw195).ArraySet1(_1161_v92, 1)
			(_nw195).ArraySet1(_1161_v92, 2)
			(_nw195).ArraySet1(_1161_v92, 3)
			(_nw195).ArraySet1(_1161_v92, 4)
			_1162_v93 = _nw195
			_1162_v93 = _1162_v93
			var _1163_v94 _dafny.Map
			_ = _1163_v94
			_1163_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm1(globalState)).Plus((_this).F32()), !(p0))
			_1163_v94 = (_1163_v94).Update((_1163_v94).Cardinality(), !(p0))
			var _1164_v95 _dafny.MultiSet
			_ = _1164_v95
			_1164_v95 = _dafny.MultiSetOf((_this).F32())
			var _1165_v96 D4
			_ = _1165_v96
			_1165_v96 = Companion_D4_.Create_DC9_((_1164_v95).Cardinality(), p3)
			_1165_v96 = _1165_v96
		} else {
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))
			_ = _index141
			(_1037_v0).ArraySet1((_1037_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1037_v0), 0))).Int()).(_dafny.Int), (_index141).Int())
			(globalState).F26 = p0
			var _1166_v97 _dafny.Array
			_ = _1166_v97
			var _nw196 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw196
			_1166_v97 = _nw196
			_1166_v97 = _1166_v97
			var _1167_v98 _dafny.Array
			_ = _1167_v98
			var _nw197 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw197
			_1167_v98 = _nw197
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1167_v98), 0))
			_ = _index142
			(_1167_v98).ArraySet1(false, (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1167_v98), 0))
			_ = _index143
			(_1167_v98).ArraySet1(p0, (_index143).Int())
			var _1168_v99 _dafny.Set
			_ = _1168_v99
			_1168_v99 = _dafny.SetOf(p0, (_this).Fm5(globalState))
			(globalState).F26 = (_1168_v99).IsSubsetOf(_1168_v99)
		}
	}
}
func (_this *C2) M7(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) {
	{
		(globalState).F3 = (p1).Times(p1)
		if (p1).Cmp(p1) != 0 {
			var _1169_v0 _dafny.Array
			_ = _1169_v0
			var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
			_ = _nw198
			_1169_v0 = _nw198
			var _1170_v1 _dafny.CodePoint
			_ = _1170_v1
			_1170_v1 = _dafny.CodePoint('d')
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1169_v0), 0))
			_ = _index144
			(_1169_v0).ArraySet1CodePoint(_1170_v1, (_index144).Int())
			var _1171_v2 bool
			_ = _1171_v2
			_1171_v2 = true
			var _1172_v3 _dafny.Set
			_ = _1172_v3
			_1172_v3 = _dafny.SetOf(_1171_v2)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1169_v0), 0))
			_ = _index145
			var _rhs183 _dafny.Int = Companion_Default___.SafeModuloInt(((func() _dafny.Set {
				if _1171_v2 {
					return _1172_v3
				}
				return _1172_v3
			})()).Cardinality(), (_this).F32())
			_ = _rhs183
			var _rhs184 bool = !(_1171_v2)
			_ = _rhs184
			var _rhs185 _dafny.CodePoint = _1170_v1
			_ = _rhs185
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			var _lhs143 *GlobalState = globalState
			_ = _lhs143
			var _lhs144 _dafny.Array = _1169_v0
			_ = _lhs144
			var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1169_v0), 0))
			_ = _lhs145
			_lhs142.F21 = _rhs183
			_lhs143.F26 = _rhs184
			(_lhs144).ArraySet1CodePoint(_rhs185, (_lhs145).Int())
			(globalState).F24 = (_this).F32()
			var _1173_v4 _dafny.Array
			_ = _1173_v4
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_20
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Int = func(_1174_i0 _dafny.Int) _dafny.Int {
					return (_1174_i0).Minus((_this).F32())
				}
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw199 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw199).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw199).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_1173_v4 = _nw199
			var _rhs186 bool = (Companion_Default___.Fm29(_1171_v2, true, _1171_v2, globalState))
			_ = _rhs186
			var _rhs187 bool = _1171_v2
			_ = _rhs187
			var _rhs188 _dafny.Array = _1173_v4
			_ = _rhs188
			var _lhs146 *GlobalState = globalState
			_ = _lhs146
			var _lhs147 *GlobalState = globalState
			_ = _lhs147
			_lhs146.F26 = _rhs186
			_lhs147.F26 = _rhs187
			_1173_v4 = _rhs188
			var _1175_v5 _dafny.Array
			_ = _1175_v5
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw200
			_1175_v5 = _nw200
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_1175_v5), 0))
			_ = _index146
			(_1175_v5).ArraySet1((_this).Fm21(globalState), (_index146).Int())
			var _1176_v6 _dafny.Sequence
			_ = _1176_v6
			_1176_v6 = _dafny.UnicodeSeqOfUtf8Bytes("wxwtlr")
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_1175_v5), 0))
			_ = _index147
			(_1175_v5).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1176_v6, _1176_v6), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1173_v4), 0))
			_ = _index148
			(_1173_v4).ArraySet1(_dafny.IntOfInt64(-949), (_index148).Int())
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1173_v4), 0))
			_ = _index149
			(_1173_v4).ArraySet1(Companion_Default___.Fm1(globalState), (_index149).Int())
		} else {
			var _1177_v7 bool
			_ = _1177_v7
			_1177_v7 = false
			var _1178_v8 bool
			_ = _1178_v8
			var _1179_v9 _dafny.Int
			_ = _1179_v9
			var _out38 bool
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			_out38, _out39 = Companion_Default___.M0(!(_1177_v7), (_dafny.Zero).Minus((p1).Minus(p1)), (_this).F32(), (p1).Plus(p1), globalState)
			_1178_v8 = _out38
			_1179_v9 = _out39
			var _1180_v10 _dafny.Sequence
			_ = _1180_v10
			_1180_v10 = _dafny.UnicodeSeqOfUtf8Bytes("lmvkrjqv")
			var _1181_v11 *C0
			_ = _1181_v11
			var _nw201 *C0 = New_C0_()
			_ = _nw201
			_nw201.Ctor__((_1179_v9).Times(Companion_Default___.Fm1(globalState)), _1180_v10)
			_1181_v11 = _nw201
			var _1182_v12 _dafny.Set
			_ = _1182_v12
			_1182_v12 = _dafny.SetOf((_dafny.Zero).Minus(p1))
			var _1183_v13 _dafny.MultiSet
			_ = _1183_v13
			_1183_v13 = _dafny.MultiSetOf((_this).F32())
			var _1184_v14 _dafny.Map
			_ = _1184_v14
			_1184_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1177_v7, _1183_v13)
			var _1185_v15 _dafny.Set
			_ = _1185_v15
			_1185_v15 = _dafny.SetOf(_1182_v12, _dafny.SetOf(_1179_v9, (_1181_v11).F36(), (_dafny.Zero).Minus((_1184_v14).Cardinality()), (_1181_v11).F36()), _dafny.SetOf((_this).F32(), _dafny.IntOfInt64(322)), _1182_v12, (_1182_v12).Difference(_1182_v12))
			var _1186_v16 _dafny.Sequence
			_ = _1186_v16
			_1186_v16 = _dafny.SeqOf(_1178_v8, _1177_v7)
			var _rhs189 _dafny.Set = _1185_v15
			_ = _rhs189
			var _rhs190 bool = _1177_v7
			_ = _rhs190
			var _rhs191 _dafny.Int = (_1179_v9).Minus(_dafny.IntOfUint32((_1186_v16).Cardinality()))
			_ = _rhs191
			var _lhs148 *GlobalState = globalState
			_ = _lhs148
			_1185_v15 = _rhs189
			_1178_v8 = _rhs190
			_lhs148.F21 = _rhs191
			var _rhs192 _dafny.Sequence = _1186_v16
			_ = _rhs192
			var _rhs193 _dafny.Sequence = (func() _dafny.Sequence {
				if !(_1177_v7) {
					return _dafny.UnicodeSeqOfUtf8Bytes("coh")
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("vbd")
			})()
			_ = _rhs193
			_1186_v16 = _rhs192
			_1180_v10 = _rhs193
			if (func() bool {
				if !((_1186_v16).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_1186_v16).Cardinality()))).Uint32()).(bool)) {
					return _1177_v7
				}
				return _1177_v7
			})() {
				(globalState).F26 = _1178_v8
				var _1187_v17 _dafny.Map
				_ = _1187_v17
				_1187_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1179_v9, (_this).F32())
				_1187_v17 = (_1187_v17).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1180_v10, (_1181_v11).F37())).Cardinality()), Companion_Default___.Fm1(globalState))
				var _1188_v18 _dafny.Sequence
				_ = _1188_v18
				_1188_v18 = _dafny.SeqOf((_1181_v11).F36())
				var _1189_v19 _dafny.Map
				_ = _1189_v19
				_1189_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1178_v8, _1188_v18)
				var _1190_v20 _dafny.Map
				_ = _1190_v20
				_1190_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1178_v8), (_1189_v19).Cardinality())
				var _rhs194 _dafny.Int = ((_1190_v20).Update(_1178_v8, p1)).Cardinality()
				_ = _rhs194
				var _rhs195 bool = !(_1178_v8)
				_ = _rhs195
				var _rhs196 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1179_v9))
				_ = _rhs196
				var _rhs197 _dafny.Int = ((_1179_v9).Plus(p1)).Times((_1181_v11).F36())
				_ = _rhs197
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				var _lhs150 *GlobalState = globalState
				_ = _lhs150
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				_lhs149.F21 = _rhs194
				_lhs150.F26 = _rhs195
				_lhs151.F21 = _rhs196
				_lhs152.F18 = _rhs197
				var _1191_v21 _dafny.CodePoint
				_ = _1191_v21
				_1191_v21 = _dafny.CodePoint('g')
				var _1192_v22 *C0
				_ = _1192_v22
				var _nw202 *C0 = New_C0_()
				_ = _nw202
				_nw202.Ctor__(Companion_Default___.Fm1(globalState), _dafny.Companion_Sequence_.Update((_1181_v11).F37(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.IntOfUint32(((_1181_v11).F37()).Cardinality()))).Uint32(), _1191_v21))
				_1192_v22 = _nw202
				(globalState).F26 = (func() bool {
					if _1178_v8 {
						return _1177_v7
					}
					return _1177_v7
				})()
			} else {
				_1178_v8 = _1178_v8
				_1180_v10 = (_1181_v11).F37()
				(globalState).F18 = _1179_v9
				(globalState).F26 = !(((_this).F31()).Union((_this).F31())).Equals((_this).F31())
				var _1193_v24 _dafny.Sequence
				_ = _1193_v24
				_1193_v24 = _dafny.SeqOf(func() _dafny.Map {
					var _coll44 = _dafny.NewMapBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(454), _dafny.IntOfInt64(314))); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1194_v23 _dafny.Int
						_1194_v23 = interface{}(_compr_44).(_dafny.Int)
						if ((_dafny.IntOfInt64(454)).Cmp(_1194_v23) <= 0) && ((_1194_v23).Cmp(_dafny.IntOfInt64(314)) < 0) {
							_coll44.Add((_1194_v23).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1177_v7, _1179_v9)).Cardinality()), (_1181_v11).F36())
						}
					}
					return _coll44.ToMap()
				}())
				var _1195_v25 _dafny.Map
				_ = _1195_v25
				_1195_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1181_v11, _1179_v9)
				var _1196_v26 _dafny.Map
				_ = _1196_v26
				_1196_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1181_v11).F36(), (func() _dafny.Int {
					if (_1195_v25).Contains(_1181_v11) {
						return (_1195_v25).Get(_1181_v11).(_dafny.Int)
					}
					return _dafny.IntOfUint32(((_1181_v11).F37()).Cardinality())
				})())
				var _1197_v27 _dafny.Sequence
				_ = _1197_v27
				_1197_v27 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1181_v11).F36(), _dafny.IntOfInt64(63)), ((_1193_v24).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_1193_v24).Cardinality()))).Uint32()).(_dafny.Map)).Update(_dafny.IntOfInt64(371), _1179_v9), Companion_Default___.Fm30(globalState), _1196_v26, _1196_v26)
				var _1198_v28 _dafny.Map
				_ = _1198_v28
				_1198_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1197_v27).Select((Companion_Default___.SafeIndex((_1181_v11).F36(), _dafny.IntOfUint32((_1197_v27).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), !(!(_1178_v8)) || ((_this).Fm5(globalState)))
				_1198_v28 = _1198_v28
			}
		}
		var _1199_v29 _dafny.Array
		_ = _1199_v29
		var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw203
		_1199_v29 = _nw203
		for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1199_v29), 0))); ; {
			_guard_loop_3, _ok49 := _iter49()
			if !_ok49 {
				break
			}
			var _1200_i1 _dafny.Int
			_1200_i1 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1200_i1).Sign() != -1) && ((_1200_i1).Cmp(_dafny.ArrayLen((_1199_v29), 0)) < 0)) {
				(_1199_v29).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ic"), _dafny.UnicodeSeqOfUtf8Bytes("hksuqeycu")) {
						return _dafny.CodePoint('u')
					}
					return _dafny.CodePoint('r')
				})(), (_1200_i1).Int())
			}
		}
		var _1201_v30 bool
		_ = _1201_v30
		_1201_v30 = true
		(globalState).F26 = _1201_v30
		var _1202_v31 _dafny.MultiSet
		_ = _1202_v31
		_1202_v31 = _dafny.MultiSetOf(_1201_v30)
		_1202_v31 = (_this).F31()
		var _1203_v32 _dafny.Sequence
		_ = _1203_v32
		_1203_v32 = _dafny.SeqOf(false)
		var _1204_v33 _dafny.MultiSet
		_ = _1204_v33
		_1204_v33 = _dafny.MultiSetOf(_1203_v32)
		var _1205_v34 _dafny.Array
		_ = _1205_v34
		var _nwElement0_48 _dafny.Int = (_1204_v33).Cardinality()
		_ = _nwElement0_48
		var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.One)
		_ = _nw204
		(_nw204).ArraySet1(_nwElement0_48, 0)
		_1205_v34 = _nw204
		_1205_v34 = _1205_v34
	}
}
func (_this *C2) M8(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _1206_v0 _dafny.Array
		_ = _1206_v0
		var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw205
		_1206_v0 = _nw205
		var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))
		_ = _index150
		(_1206_v0).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(155))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg80 _dafny.Int) interface{} {
				return coer80(arg80)
			}
		}(func(_1207_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})), (_index150).Int())
		var _1208_v1 _dafny.Sequence
		_ = _1208_v1
		_1208_v1 = _dafny.UnicodeSeqOfUtf8Bytes("rwtdknsgw")
		var _1209_v2 bool
		_ = _1209_v2
		_1209_v2 = false
		var _1210_v3 _dafny.Map
		_ = _1210_v3
		_1210_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1208_v1).Cardinality()), _1209_v2)
		var _1211_v4 _dafny.Map
		_ = _1211_v4
		_1211_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_1210_v3).Contains(p1) {
				return (_1210_v3).Get(p1).(bool)
			}
			return _1209_v2
		})(), _1208_v1)
		var _1212_v5 _dafny.Set
		_ = _1212_v5
		_1212_v5 = _dafny.SetOf(!(_1209_v2), _1209_v2)
		var _1213_v6 _dafny.Map
		_ = _1213_v6
		_1213_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1212_v5, !(_1209_v2))
		var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1206_v0), 0))
		_ = _index151
		(_1206_v0).ArraySet1((func() _dafny.Sequence {
			if (_1211_v4).Contains((func() bool {
				if (_1213_v6).Contains(_1212_v5) {
					return (_1213_v6).Get(_1212_v5).(bool)
				}
				return _1209_v2
			})()) {
				return (_1211_v4).Get((func() bool {
					if (_1213_v6).Contains(_1212_v5) {
						return (_1213_v6).Get(_1212_v5).(bool)
					}
					return _1209_v2
				})()).(_dafny.Sequence)
			}
			return _1208_v1
		})(), (_index151).Int())
		var _1214_v7 _dafny.MultiSet
		_ = _1214_v7
		_1214_v7 = _dafny.MultiSetOf(_1209_v2)
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))
		_ = _index152
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1206_v0), 0))
		_ = _index153
		var _rhs198 _dafny.Int = ((_this).F32()).Minus(_dafny.IntOfInt64(978))
		_ = _rhs198
		var _rhs199 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1215_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))
		_ = _rhs199
		var _rhs200 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1216_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))
		_ = _rhs200
		var _rhs201 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("hucvcbfcp")
		_ = _rhs201
		var _rhs202 _dafny.MultiSet = _1214_v7
		_ = _rhs202
		var _lhs153 *GlobalState = globalState
		_ = _lhs153
		var _lhs154 _dafny.Array = _1206_v0
		_ = _lhs154
		var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))
		_ = _lhs155
		var _lhs156 _dafny.Array = _1206_v0
		_ = _lhs156
		var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1206_v0), 0))
		_ = _lhs157
		_lhs153.F18 = _rhs198
		(_lhs154).ArraySet1(_rhs199, (_lhs155).Int())
		(_lhs156).ArraySet1(_rhs200, (_lhs157).Int())
		_1208_v1 = _rhs201
		_1214_v7 = _rhs202
		var _1217_v8 _dafny.Set
		_ = _1217_v8
		_1217_v8 = _dafny.SetOf((_this).F32(), (_this).F32())
		var _1218_v9 _dafny.Sequence
		_ = _1218_v9
		_1218_v9 = _dafny.SeqOf(_1217_v8, _1217_v8)
		var _1219_v10 _dafny.MultiSet
		_ = _1219_v10
		_1219_v10 = _dafny.MultiSetOf(p1, p1, p1)
		var _1220_v11 _dafny.Map
		_ = _1220_v11
		_1220_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v2, (_this).F32())
		var _1221_v12 D8
		_ = _1221_v12
		_1221_v12 = Companion_D8_.Create_DC16_((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1218_v9).Cardinality()))), (func() _dafny.Int {
			if (_1219_v10).Contains((_1220_v11).Cardinality()) {
				return (_1219_v10).Multiplicity((_1220_v11).Cardinality())
			}
			return (_this).F32()
		})())
		var _pat_let_tv19 = _1209_v2
		_ = _pat_let_tv19
		(globalState).F26 = func(_source20 D8) bool {
			if _source20.Is_DC16() {
				var _1222___mcc_h0 _dafny.Int = _source20.Get_().(D8_DC16).Cf29
				_ = _1222___mcc_h0
				var _1223___mcc_h1 _dafny.Int = _source20.Get_().(D8_DC16).Cf30
				_ = _1223___mcc_h1
				var _1224_cf30 _dafny.Int = _1223___mcc_h1
				_ = _1224_cf30
				var _1225_cf29 _dafny.Int = _1222___mcc_h0
				_ = _1225_cf29
				return _pat_let_tv19
			} else {
				var _1226___mcc_h2 _dafny.Map = _source20.Get_().(D8_DC15).Cf28
				_ = _1226___mcc_h2
				var _1227_cf28 _dafny.Map = _1226___mcc_h2
				_ = _1227_cf28
				return true
			}
		}(_1221_v12)
		var _1228_v13 _dafny.Map
		_ = _1228_v13
		_1228_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1209_v2), _1209_v2)
		var _1229_v14 _dafny.Sequence
		_ = _1229_v14
		_1229_v14 = _dafny.SeqOf(p1)
		var _1230_v15 _dafny.Map
		_ = _1230_v15
		_1230_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1229_v14), false)
		var _1231_v16 D8
		_ = _1231_v16
		_1231_v16 = Companion_D8_.Create_DC15_(_1228_v13)
		var _1232_v17 _dafny.Array
		_ = _1232_v17
		var _nwElement0_49 _dafny.Map = _1228_v13
		_ = _nwElement0_49
		var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(19))
		_ = _nw206
		(_nw206).ArraySet1(_nwElement0_49, 0)
		(_nw206).ArraySet1((_1228_v13).Update(_1209_v2, (func() bool {
			if (_1230_v15).Contains(_dafny.MultiSetFromSeq(_1229_v14)) {
				return (_1230_v15).Get(_dafny.MultiSetFromSeq(_1229_v14)).(bool)
			}
			return !(_1209_v2)
		})()), 1)
		(_nw206).ArraySet1(_1228_v13, 2)
		(_nw206).ArraySet1(_1228_v13, 3)
		(_nw206).ArraySet1((_1228_v13).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v2, !(_1209_v2))), 4)
		(_nw206).ArraySet1(_1228_v13, 5)
		(_nw206).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1209_v2), 6)
		(_nw206).ArraySet1(_1228_v13, 7)
		(_nw206).ArraySet1((_1228_v13).Update(_1209_v2, _1209_v2), 8)
		(_nw206).ArraySet1(_1228_v13, 9)
		(_nw206).ArraySet1(_1228_v13, 10)
		(_nw206).ArraySet1((_1228_v13).Merge(Companion_Default___.Fm23(p1, _1209_v2, globalState)), 11)
		(_nw206).ArraySet1((_1231_v16).Dtor_cf28(), 12)
		(_nw206).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v2, _1209_v2)).Merge(_1228_v13), 13)
		(_nw206).ArraySet1(_1228_v13, 14)
		(_nw206).ArraySet1(_1228_v13, 15)
		(_nw206).ArraySet1((_1228_v13).Merge(_1228_v13), 16)
		(_nw206).ArraySet1(_1228_v13, 17)
		(_nw206).ArraySet1(_1228_v13, 18)
		_1232_v17 = _nw206
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1232_v17), 0))); ; {
			_guard_loop_4, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _1233_i3 _dafny.Int
			_1233_i3 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1233_i3).Sign() != -1) && ((_1233_i3).Cmp(_dafny.ArrayLen((_1232_v17), 0)) < 0)) {
				(_1232_v17).ArraySet1((_1228_v13).Merge(_1228_v13), (_1233_i3).Int())
			}
		}
		var _1234_i4 _dafny.Int
		_ = _1234_i4
		_1234_i4 = _dafny.Zero
		{
			for ((((_1214_v7).Update(_1209_v2, Companion_Default___.Abs(p1))).Intersection((_this).F31())).Cardinality()).Cmp(p1) == 0 {
				{
					if (_1234_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1234_i4 = (_1234_i4).Plus(_dafny.One)
					var _1235_v18 *C0
					_ = _1235_v18
					var _nw207 *C0 = New_C0_()
					_ = _nw207
					_nw207.Ctor__((_this).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(521))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg83 _dafny.Int) interface{} {
							return coer83(arg83)
						}
					}(func(_1236_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					})))
					_1235_v18 = _nw207
					var _nw208 *C0 = New_C0_()
					_ = _nw208
					_nw208.Ctor__(Companion_Default___.SafeDivisionInt((_1217_v8).Cardinality(), p1), _dafny.UnicodeSeqOfUtf8Bytes("lwjbo"))
					_1235_v18 = _nw208
					var _1237_v19 _dafny.Array
					_ = _1237_v19
					var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
					_ = _nw209
					_1237_v19 = _nw209
					var _1238_v20 _dafny.Map
					_ = _1238_v20
					_1238_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v19, (_dafny.MultiSetOf(_1214_v7, (_this).F31())).Cardinality())
					var _1239_v21 _dafny.CodePoint
					_ = _1239_v21
					_1239_v21 = _dafny.CodePoint('g')
					var _1240_v22 D2
					_ = _1240_v22
					_1240_v22 = Companion_D2_.Create_DC4_((_this).F32(), _dafny.IntOfUint32(((_1235_v18).F37()).Cardinality()), (func() _dafny.Int {
						if (_1219_v10).Contains((func() _dafny.Int {
							if (_1238_v20).Contains(_1237_v19) {
								return (_1238_v20).Get(_1237_v19).(_dafny.Int)
							}
							return p1
						})()) {
							return (_1219_v10).Multiplicity((func() _dafny.Int {
								if (_1238_v20).Contains(_1237_v19) {
									return (_1238_v20).Get(_1237_v19).(_dafny.Int)
								}
								return p1
							})())
						}
						return (_this).F32()
					})(), _1239_v21, _1208_v1)
					var _source21 D2 = _1240_v22
					_ = _source21
					if _source21.Is_DC4() {
						var _1241___mcc_h3 _dafny.Int = _source21.Get_().(D2_DC4).Cf5
						_ = _1241___mcc_h3
						var _1242___mcc_h4 _dafny.Int = _source21.Get_().(D2_DC4).Cf6
						_ = _1242___mcc_h4
						var _1243___mcc_h5 _dafny.Int = _source21.Get_().(D2_DC4).Cf7
						_ = _1243___mcc_h5
						var _1244___mcc_h6 _dafny.CodePoint = _source21.Get_().(D2_DC4).Cf8
						_ = _1244___mcc_h6
						var _1245___mcc_h7 _dafny.Sequence = _source21.Get_().(D2_DC4).Cf9
						_ = _1245___mcc_h7
						var _1246_cf9 _dafny.Sequence = _1245___mcc_h7
						_ = _1246_cf9
						var _1247_cf8 _dafny.CodePoint = _1244___mcc_h6
						_ = _1247_cf8
						var _1248_cf7 _dafny.Int = _1243___mcc_h5
						_ = _1248_cf7
						var _1249_cf6 _dafny.Int = _1242___mcc_h4
						_ = _1249_cf6
						var _1250_cf5 _dafny.Int = _1241___mcc_h3
						_ = _1250_cf5
						var _1251_v23 _dafny.Sequence
						_ = _1251_v23
						_1251_v23 = _dafny.SeqOf((_1235_v18).F37())
						_1249_cf6 = (func() _dafny.Int {
							if _1209_v2 {
								return ((_dafny.MultiSetFromSeq(_1251_v23)).Cardinality()).Plus((_1235_v18).F36())
							}
							return (_1235_v18).F36()
						})()
						_1249_cf6 = ((_dafny.SetOf(false, _1209_v2)).Difference((_dafny.SetOf(_1209_v2)).Difference(_1212_v5))).Cardinality()
						var _1252_v24 bool
						_ = _1252_v24
						_1252_v24 = _1209_v2
						var _1253_v25 _dafny.Map
						_ = _1253_v25
						_1253_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v19, _1209_v2)
						var _1254_v26 _dafny.Sequence
						_ = _1254_v26
						_1254_v26 = _dafny.SeqOf((p1).Cmp(_1249_cf6) >= 0, (_1209_v2), (func() bool {
							if (_1253_v25).Contains(_1237_v19) {
								return (_1253_v25).Get(_1237_v19).(bool)
							}
							return _1209_v2
						})(), _1209_v2, _1209_v2)
						var _1255_v27 _dafny.MultiSet
						_ = _1255_v27
						_1255_v27 = _dafny.MultiSetOf(_1237_v19)
						var _1256_v28 _dafny.Array
						_ = _1256_v28
						var _len0_21 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_21
						var _nw210 _dafny.Array
						_ = _nw210
						if _len0_21.Cmp(_dafny.Zero) == 0 {
							_nw210 = _dafny.NewArray(_len0_21)
						} else {
							var _init21 func(_dafny.Int) _dafny.MultiSet = (func(_1257_v10 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
								return func(_1258_i6 _dafny.Int) _dafny.MultiSet {
									return _1257_v10
								}
							})(_1219_v10)
							_ = _init21
							var _element0_21 = _init21(_dafny.Zero)
							_ = _element0_21
							_nw210 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
							(_nw210).ArraySet1(_element0_21, 0)
							var _nativeLen0_21 = (_len0_21).Int()
							_ = _nativeLen0_21
							for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
								(_nw210).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
							}
						}
						_1256_v28 = _nw210
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1256_v28), 0))
						_ = _index154
						(_1256_v28).ArraySet1((func() _dafny.MultiSet {
							if _1209_v2 {
								return _1219_v10
							}
							return _1219_v10
						})(), (_index154).Int())
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1256_v28), 0))
						_ = _index155
						var _rhs203 bool = !(_1209_v2)
						_ = _rhs203
						var _rhs204 _dafny.Sequence = _1254_v26
						_ = _rhs204
						var _rhs205 _dafny.MultiSet = (_1255_v27).Difference(((_1255_v27).Update(_1237_v19, Companion_Default___.Abs((_1235_v18).F36()))).Union(_1255_v27))
						_ = _rhs205
						var _rhs206 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1229_v14, _1229_v14))).Difference((_1219_v10).Intersection(_1219_v10))
						_ = _rhs206
						var _rhs207 _dafny.Int = (_dafny.Zero).Minus(((_dafny.IntOfInt64(163)).Times(_dafny.IntOfInt64(-469))).Minus(_dafny.IntOfInt64(883)))
						_ = _rhs207
						var _lhs158 *GlobalState = globalState
						_ = _lhs158
						var _lhs159 _dafny.Array = _1256_v28
						_ = _lhs159
						var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1256_v28), 0))
						_ = _lhs160
						var _lhs161 *GlobalState = globalState
						_ = _lhs161
						_lhs158.F26 = _rhs203
						_1254_v26 = _rhs204
						_1255_v27 = _rhs205
						(_lhs159).ArraySet1(_rhs206, (_lhs160).Int())
						_lhs161.F24 = _rhs207
						var _1259_v29 *C0
						_ = _1259_v29
						var _nw211 *C0 = New_C0_()
						_ = _nw211
						_nw211.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jlhbrjm")).Cardinality()), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if _1209_v2 {
								return _1208_v1
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("fjfolha")
						})(), (Companion_Default___.SafeIndex((_1235_v18).F36(), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if _1209_v2 {
								return _1208_v1
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("fjfolha")
						})()).Cardinality()))).Uint32(), ((_1206_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_1235_v18).F36(), _dafny.IntOfUint32(((_1206_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)))
						_1259_v29 = _nw211
					} else if _source21.Is_DC3() {
						var _1260___mcc_h8 _dafny.Map = _source21.Get_().(D2_DC3).Cf4
						_ = _1260___mcc_h8
						var _1261_cf4 _dafny.Map = _1260___mcc_h8
						_ = _1261_cf4
						(globalState).F18 = p1
						_1239_v21 = (func() _dafny.CodePoint {
							if _1209_v2 {
								return (func() _dafny.CodePoint {
									if _1209_v2 {
										return _1239_v21
									}
									return _1239_v21
								})()
							}
							return _1239_v21
						})()
						var _1262_v30 _dafny.Map
						_ = _1262_v30
						_1262_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v2, _1228_v13)
						(globalState).F24 = (_1262_v30).Cardinality()
						var _1263_v31 bool
						_ = _1263_v31
						var _1264_v32 _dafny.Int
						_ = _1264_v32
						var _out40 bool
						_ = _out40
						var _out41 _dafny.Int
						_ = _out41
						_out40, _out41 = Companion_Default___.M0(_1209_v2, _dafny.IntOfInt64(935), p1, _dafny.IntOfInt64(299), globalState)
						_1263_v31 = _out40
						_1264_v32 = _out41
					} else {
						var _1265___mcc_h9 D2 = _source21.Get_().(D2_DC5).Cf10
						_ = _1265___mcc_h9
						var _1266_cf10 D2 = _1265___mcc_h9
						_ = _1266_cf10
						_1229_v14 = _1229_v14
						(globalState).F26 = !(_1209_v2)
						var _rhs208 _dafny.Int = (_1235_v18).F36()
						_ = _rhs208
						var _rhs209 bool = !(_1209_v2) || (_1209_v2)
						_ = _rhs209
						var _rhs210 bool = !(((_1235_v18).F36()).Cmp(((_1235_v18).F36()).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v2, true)).Cardinality())) < 0)
						_ = _rhs210
						var _lhs162 *GlobalState = globalState
						_ = _lhs162
						var _lhs163 *GlobalState = globalState
						_ = _lhs163
						_lhs162.F24 = _rhs208
						_1209_v2 = _rhs209
						_lhs163.F26 = _rhs210
						(globalState).F26 = _1209_v2
					}
					var _1267_v33 _dafny.MultiSet
					_ = _1267_v33
					_1267_v33 = _dafny.MultiSetOf((_1206_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1206_v0), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("mlmcu"), (_1235_v18).F37())
					(globalState).F8 = _1267_v33
					var _1268_v34 _dafny.MultiSet
					_ = _1268_v34
					_1268_v34 = _dafny.MultiSetOf(_1217_v8, _1217_v8)
					var _1269_v35 D11
					_ = _1269_v35
					_1269_v35 = Companion_D11_.Create_DC23_(_1268_v34)
					var _source22 D11 = _1269_v35
					_ = _source22
					if _source22.Is_DC24() {
						var _1270___mcc_h10 _dafny.Int = _source22.Get_().(D11_DC24).Cf43
						_ = _1270___mcc_h10
						var _1271___mcc_h11 D4 = _source22.Get_().(D11_DC24).Cf44
						_ = _1271___mcc_h11
						var _1272_cf44 D4 = _1271___mcc_h11
						_ = _1272_cf44
						var _1273_cf43 _dafny.Int = _1270___mcc_h10
						_ = _1273_cf43
						(globalState).F26 = _1209_v2
						var _1274_v36 _dafny.Sequence
						_ = _1274_v36
						_1274_v36 = _dafny.SeqOf(_1209_v2)
						(globalState).F3 = (_dafny.IntOfUint32((_1274_v36).Cardinality())).Times((_1235_v18).F36())
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1237_v19), 0))
						_ = _index156
						(_1237_v19).ArraySet1(_1209_v2, (_index156).Int())
						var _1275_v37 D4
						_ = _1275_v37
						_1275_v37 = Companion_D4_.Create_DC9_(p1, _1209_v2)
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1237_v19), 0))
						_ = _index157
						(_1237_v19).ArraySet1((func() bool {
							if (_1228_v13).Contains(!((_1275_v37).Dtor_cf17()) || (true)) {
								return (_1228_v13).Get(!((_1275_v37).Dtor_cf17()) || (true)).(bool)
							}
							return !(_1209_v2)
						})(), (_index157).Int())
						var _1276_v38 _dafny.Array
						_ = _1276_v38
						var _len0_22 _dafny.Int = _dafny.IntOfInt64(29)
						_ = _len0_22
						var _nw212 _dafny.Array
						_ = _nw212
						if _len0_22.Cmp(_dafny.Zero) == 0 {
							_nw212 = _dafny.NewArray(_len0_22)
						} else {
							var _init22 func(_dafny.Int) _dafny.Int = (func(_1277_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1278_i7 _dafny.Int) _dafny.Int {
									return (_1278_i7).Plus(_1277_p1)
								}
							})(p1)
							_ = _init22
							var _element0_22 = _init22(_dafny.Zero)
							_ = _element0_22
							_nw212 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
							(_nw212).ArraySet1(_element0_22, 0)
							var _nativeLen0_22 = (_len0_22).Int()
							_ = _nativeLen0_22
							for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
								(_nw212).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
							}
						}
						_1276_v38 = _nw212
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1276_v38), 0))
						_ = _index158
						(_1276_v38).ArraySet1(_1273_cf43, (_index158).Int())
						var _1279_v39 _dafny.Array
						_ = _1279_v39
						var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
						_ = _nw213
						_1279_v39 = _nw213
						var _1280_v40 _dafny.Array
						_ = _1280_v40
						var _nwElement0_50 *C0 = _1235_v18
						_ = _nwElement0_50
						var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(10))
						_ = _nw214
						(_nw214).ArraySet1(_nwElement0_50, 0)
						(_nw214).ArraySet1(_1235_v18, 1)
						(_nw214).ArraySet1(_1235_v18, 2)
						(_nw214).ArraySet1(_1235_v18, 3)
						(_nw214).ArraySet1(_1235_v18, 4)
						(_nw214).ArraySet1(_1235_v18, 5)
						(_nw214).ArraySet1(_1235_v18, 6)
						(_nw214).ArraySet1(_1235_v18, 7)
						(_nw214).ArraySet1(_1235_v18, 8)
						(_nw214).ArraySet1(_1235_v18, 9)
						_1280_v40 = _nw214
						var _1281_v41 _dafny.Sequence
						_ = _1281_v41
						_1281_v41 = _dafny.SeqOf(_1280_v40, _1280_v40)
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1279_v39), 0))
						_ = _index159
						(_1279_v39).ArraySet1((_1281_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.IntOfUint32((_1281_v41).Cardinality()))).Uint32()).(_dafny.Array), (_index159).Int())
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1276_v38), 0))
						_ = _index160
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1279_v39), 0))
						_ = _index161
						var _rhs211 _dafny.Set = _1217_v8
						_ = _rhs211
						var _rhs212 _dafny.Int = Companion_Default___.SafeModuloInt(_1273_cf43, _1273_cf43)
						_ = _rhs212
						var _rhs213 bool = ((func() bool {
							if (_1237_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1237_v19), 0))).Int()).(bool) {
								return (_1237_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1237_v19), 0))).Int()).(bool)
							}
							return (_1237_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1237_v19), 0))).Int()).(bool)
						})()) == (_1209_v2)
						_ = _rhs213
						var _rhs214 _dafny.Int = (_this).F32()
						_ = _rhs214
						var _rhs215 _dafny.Array = _1280_v40
						_ = _rhs215
						var _lhs164 _dafny.Array = _1276_v38
						_ = _lhs164
						var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1276_v38), 0))
						_ = _lhs165
						var _lhs166 *GlobalState = globalState
						_ = _lhs166
						var _lhs167 *GlobalState = globalState
						_ = _lhs167
						var _lhs168 _dafny.Array = _1279_v39
						_ = _lhs168
						var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1279_v39), 0))
						_ = _lhs169
						_1217_v8 = _rhs211
						(_lhs164).ArraySet1(_rhs212, (_lhs165).Int())
						_lhs166.F26 = _rhs213
						_lhs167.F15 = _rhs214
						(_lhs168).ArraySet1(_rhs215, (_lhs169).Int())
					} else if _source22.Is_DC23() {
						var _1282___mcc_h12 _dafny.MultiSet = _source22.Get_().(D11_DC23).Cf42
						_ = _1282___mcc_h12
						var _1283_cf42 _dafny.MultiSet = _1282___mcc_h12
						_ = _1283_cf42
						_1208_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg84 _dafny.Int) interface{} {
								return coer84(arg84)
							}
						}((func(_1284_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1285_i8 _dafny.Int) _dafny.CodePoint {
								return _1284_v21
							}
						})(_1239_v21))), (_1235_v18).F37())
						var _1286_v42 *C1
						_ = _1286_v42
						var _nw215 *C1 = New_C1_()
						_ = _nw215
						_nw215.Ctor__(_1209_v2, (_this).F31(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(9))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg85 _dafny.Int) interface{} {
								return coer85(arg85)
							}
						}((func(_1287_v18 *C0) func(_dafny.Int) _dafny.Int {
							return func(_1288_i9 _dafny.Int) _dafny.Int {
								return (_1287_v18).F36()
							}
						})(_1235_v18))), _1229_v14), (_this).F32(), _1206_v0)
						_1286_v42 = _nw215
						var _1289_v43 _dafny.Sequence
						_ = _1289_v43
						_1289_v43 = _dafny.SeqOf(_1286_v42.F35)
						(globalState).F26 = ((func() _dafny.Int {
							if true {
								return (_1235_v18).F36()
							}
							return (_1235_v18).F36()
						})()).Cmp(_dafny.IntOfUint32((_1289_v43).Cardinality())) < 0
						var _1290_v44 _dafny.Sequence
						_ = _1290_v44
						_1290_v44 = _dafny.SeqOf(_1289_v43)
						var _1291_v45 D6
						_ = _1291_v45
						_1291_v45 = Companion_D6_.Create_DC12_(!(_1286_v42.F35), _1290_v44, (func() _dafny.Sequence {
							if (_1211_v4).Contains(_1286_v42.F35) {
								return (_1211_v4).Get(_1286_v42.F35).(_dafny.Sequence)
							}
							return _1208_v1
						})())
						(_1286_v42).F35 = _dafny.Companion_Sequence_.Contains((_1291_v45).Dtor_cf22(), _1239_v21)
					} else {
						var _1292___mcc_h13 D11 = _source22.Get_().(D11_DC25).Cf45
						_ = _1292___mcc_h13
						var _1293_cf45 D11 = _1292___mcc_h13
						_ = _1293_cf45
						(globalState).F26 = ((_1214_v7).Union((_1214_v7).Update(_1209_v2, Companion_Default___.Abs((_1235_v18).F36())))).IsDisjointFrom(_1214_v7)
						(globalState).F0 = (_dafny.Zero).Minus((_1235_v18).F36())
						var _1294_v46 *C1
						_ = _1294_v46
						var _nw216 *C1 = New_C1_()
						_ = _nw216
						_nw216.Ctor__(_1209_v2, (_this).F31(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-32))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg86 _dafny.Int) interface{} {
								return coer86(arg86)
							}
						}((func(_1295_v18 *C0) func(_dafny.Int) _dafny.Int {
							return func(_1296_i10 _dafny.Int) _dafny.Int {
								return (_1295_v18).F36()
							}
						})(_1235_v18))), (_1212_v5).Cardinality(), _1206_v0)
						_1294_v46 = _nw216
						var _1297_v47 _dafny.Set
						_ = _1297_v47
						_1297_v47 = _dafny.SetOf(_1294_v46)
						var _1298_v48 _dafny.Array
						_ = _1298_v48
						var _nwElement0_51 _dafny.Int = (_this).F32()
						_ = _nwElement0_51
						var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(18))
						_ = _nw217
						(_nw217).ArraySet1(_nwElement0_51, 0)
						(_nw217).ArraySet1((_1235_v18).F36(), 1)
						(_nw217).ArraySet1((_this).F32(), 2)
						(_nw217).ArraySet1((_1235_v18).F36(), 3)
						(_nw217).ArraySet1(((_1238_v20).Update(_1237_v19, p1)).Cardinality(), 4)
						(_nw217).ArraySet1((_1235_v18).F36(), 5)
						(_nw217).ArraySet1(p1, 6)
						(_nw217).ArraySet1(p1, 7)
						(_nw217).ArraySet1((_1235_v18).F36(), 8)
						(_nw217).ArraySet1(((_1235_v18).F36()).Minus((_dafny.Zero).Minus(p1)), 9)
						(_nw217).ArraySet1(_dafny.IntOfInt64(204), 10)
						(_nw217).ArraySet1(p1, 11)
						(_nw217).ArraySet1((_1235_v18).F36(), 12)
						(_nw217).ArraySet1(_dafny.IntOfUint32((_1208_v1).Cardinality()), 13)
						(_nw217).ArraySet1((_dafny.Zero).Minus((_1297_v47).Cardinality()), 14)
						(_nw217).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1208_v1, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1208_v1).Cardinality()))).Uint32(), _1239_v21)).Cardinality())), 15)
						(_nw217).ArraySet1((_1235_v18).F36(), 16)
						(_nw217).ArraySet1((_1235_v18).F36(), 17)
						_1298_v48 = _nw217
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1298_v48), 0))
						_ = _index162
						(_1298_v48).ArraySet1((_dafny.Zero).Minus((_1235_v18).F36()), (_index162).Int())
						var _pat_let_tv20 = _1239_v21
						_ = _pat_let_tv20
						var _pat_let_tv21 = _1235_v18
						_ = _pat_let_tv21
						var _pat_let_tv22 = globalState
						_ = _pat_let_tv22
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1298_v48), 0))
						_ = _index163
						(_1298_v48).ArraySet1((func(_pat_let22_0 D2) D2 {
							return func(_1299_dt__update__tmp_h1 D2) D2 {
								return func(_pat_let23_0 _dafny.CodePoint) D2 {
									return func(_1300_dt__update_hcf8_h0 _dafny.CodePoint) D2 {
										return Companion_D2_.Create_DC4_((_1299_dt__update__tmp_h1).Dtor_cf5(), (_1299_dt__update__tmp_h1).Dtor_cf6(), (_1299_dt__update__tmp_h1).Dtor_cf7(), _1300_dt__update_hcf8_h0, (_1299_dt__update__tmp_h1).Dtor_cf9())
									}(_pat_let23_0)
								}(Companion_Default___.Fm26(_pat_let_tv20, _dafny.IntOfInt64(677), (_pat_let_tv21).F36(), _pat_let_tv22))
							}(_pat_let22_0)
						}(_1240_v22)).Dtor_cf7(), (_index163).Int())
						var _1301_v49 _dafny.Sequence
						_ = _1301_v49
						_1301_v49 = _dafny.SeqOf((func() _dafny.Array {
							if (func() bool {
								if (_1228_v13).Contains(_1209_v2) {
									return (_1228_v13).Get(_1209_v2).(bool)
								}
								return _1209_v2
							})() {
								return _1298_v48
							}
							return _1298_v48
						})())
						var _1302_v50 D10
						_ = _1302_v50
						_1302_v50 = Companion_D10_.Create_DC21_(_1229_v14)
						_1298_v48 = (_1301_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1229_v14, (_1302_v50).Dtor_cf39())).Cardinality()), _dafny.IntOfUint32((_1301_v49).Cardinality()))).Uint32()).(_dafny.Array)
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1303_v51 _dafny.MultiSet
		_ = _1303_v51
		_1303_v51 = _dafny.MultiSetOf(_dafny.SetOf((_this).F32()))
		_1303_v51 = _1303_v51
		_1209_v2 = _1209_v2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f27 _dafny.Int
	_f28 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F27() _dafny.Int {
	return _this._f27
}
func (_this *C3) F28() _dafny.Array {
	return _this._f28
}
func (_this *C3) Ctor__(f27 _dafny.Int, f28 _dafny.Array) {
	{
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C3) Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F27())
	}
}
func (_this *C3) Fm45(globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_this).F27())).Minus((_this).F27())
	}
}
func (_this *C3) Fm46(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return ((!(false)) || (false)) == (!((func() bool {
			if (Companion_D16_.Create_DC40_(false, _dafny.IntOfInt64(418), (_dafny.Zero).Minus((_this).F27()))).Dtor_cf74() {
				return true
			}
			return true
		})()))
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f27 _dafny.Int
	_f28 _dafny.Array
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F27() _dafny.Int {
	return _this._f27
}
func (_this *C4) F28() _dafny.Array {
	return _this._f28
}
func (_this *C4) Ctor__(f27 _dafny.Int, f28 _dafny.Array) {
	{
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C4) Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		if !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false, true), false) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F27())
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F27())
		}
	}
}
func (_this *C4) Fm34(p0 _dafny.Set, p1 D2, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F27()).Minus((_this).F27())
	}
}
func (_this *C4) Fm35(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(true)
	}
}
func (_this *C4) M11(p0 *C0, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1304_v0 _dafny.Map
		_ = _1304_v0
		_1304_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _1305_v1 _dafny.Sequence
		_ = _1305_v1
		_1305_v1 = _dafny.SeqOf(p1, p1)
		var _1306_v2 _dafny.Sequence
		_ = _1306_v2
		_1306_v2 = _dafny.SeqOf(_1305_v1)
		var _1307_v3 D6
		_ = _1307_v3
		_1307_v3 = Companion_D6_.Create_DC12_(p1, _1306_v2, (p0).F37())
		var _1308_v4 _dafny.CodePoint
		_ = _1308_v4
		_1308_v4 = _dafny.CodePoint('k')
		(globalState).F18 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (func() bool {
				if (_1304_v0).Contains(p1) {
					return (_1304_v0).Get(p1).(bool)
				}
				return p1
			})() {
				return (_1307_v3).Dtor_cf22()
			}
			return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((p0).F37(), (p0).F37()), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((p0).F37(), (p0).F37())).Cardinality()))).Uint32(), _1308_v4)
		})()).Cardinality())
		var _1309_i0 _dafny.Int
		_ = _1309_i0
		_1309_i0 = _dafny.Zero
		{
			for !(!((p1) && (p1)) || (p1)) {
				{
					if (_1309_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1309_i0 = (_1309_i0).Plus(_dafny.One)
					var _1310_v5 _dafny.Sequence
					_ = _1310_v5
					_1310_v5 = _dafny.SeqOf((_dafny.SetOf(p1)).Cardinality(), _dafny.IntOfInt64(277))
					var _1311_v6 _dafny.Set
					_ = _1311_v6
					_1311_v6 = _dafny.SetOf(_dafny.IntOfInt64(-76))
					var _1312_v7 D2
					_ = _1312_v7
					_1312_v7 = Companion_D2_.Create_DC4_((_this).F27(), (p0).F36(), (p0).F36(), _1308_v4, _dafny.UnicodeSeqOfUtf8Bytes("njlrvlaga"))
					var _1313_v8 *C1
					_ = _1313_v8
					var _nw218 *C1 = New_C1_()
					_ = _nw218
					_nw218.Ctor__(p1, _dafny.MultiSetFromSeq(_1305_v1), _dafny.Companion_Sequence_.Concatenate(_1310_v5, _1310_v5), Companion_Default___.SafeDivisionInt((_this).Fm34(_1311_v6, _1312_v7, globalState), (_1311_v6).Cardinality()), (_this).F28())
					_1313_v8 = _nw218
					var _1314_v9 _dafny.Array
					_ = _1314_v9
					var _len0_23 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_23
					var _nw219 _dafny.Array
					_ = _nw219
					if _len0_23.Cmp(_dafny.Zero) == 0 {
						_nw219 = _dafny.NewArray(_len0_23)
					} else {
						var _init23 func(_dafny.Int) _dafny.Sequence = (func(_1315_p1 bool, _1316_v8 *C1) func(_dafny.Int) _dafny.Sequence {
							return func(_1317_i1 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(_1315_p1, _1316_v8.F35)
							}
						})(p1, _1313_v8)
						_ = _init23
						var _element0_23 = _init23(_dafny.Zero)
						_ = _element0_23
						_nw219 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
						(_nw219).ArraySet1(_element0_23, 0)
						var _nativeLen0_23 = (_len0_23).Int()
						_ = _nativeLen0_23
						for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
							(_nw219).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
						}
					}
					_1314_v9 = _nw219
					var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))
					_ = _index164
					(_1314_v9).ArraySet1((_1306_v2).Select((Companion_Default___.SafeIndex((_1311_v6).Cardinality(), _dafny.IntOfUint32((_1306_v2).Cardinality()))).Uint32()).(_dafny.Sequence), (_index164).Int())
					var _1318_v10 _dafny.MultiSet
					_ = _1318_v10
					_1318_v10 = _dafny.MultiSetOf(p1)
					var _pat_let_tv23 = _1313_v8
					_ = _pat_let_tv23
					var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))
					_ = _index165
					var _rhs216 *C1 = (func(_pat_let24_0 D13) D13 {
						return func(_1319_dt__update__tmp_h0 D13) D13 {
							return func(_pat_let25_0 *C1) D13 {
								return func(_1320_dt__update_hcf55_h0 *C1) D13 {
									return Companion_D13_.Create_DC30_(_1320_dt__update_hcf55_h0)
								}(_pat_let25_0)
							}(_pat_let_tv23)
						}(_pat_let24_0)
					}(Companion_D13_.Create_DC30_(_1313_v8))).Dtor_cf55()
					_ = _rhs216
					var _rhs217 bool = true
					_ = _rhs217
					var _rhs218 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1305_v1, (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_1305_v1).Cardinality()))).Uint32(), _1313_v8.F35), (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1305_v1, (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_1305_v1).Cardinality()))).Uint32(), _1313_v8.F35)).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((p0).F36(), _dafny.IntOfInt64(688)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1305_v1, (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_1305_v1).Cardinality()))).Uint32(), _1313_v8.F35), (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1305_v1, (Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_1305_v1).Cardinality()))).Uint32(), _1313_v8.F35)).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.SeqOf(p1, (_1313_v8).Fm3(_1318_v10, globalState)), _1305_v1, _1305_v1), _1306_v2))
					_ = _rhs218
					var _lhs170 *GlobalState = globalState
					_ = _lhs170
					var _lhs171 _dafny.Array = _1314_v9
					_ = _lhs171
					var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))
					_ = _lhs172
					_1313_v8 = _rhs216
					_lhs170.F26 = _rhs217
					(_lhs171).ArraySet1(_rhs218, (_lhs172).Int())
					(globalState).F3 = (p0).F36()
					(globalState).F26 = (p1) && (!(_1313_v8.F35) || ((_1305_v1).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1305_v1).Cardinality()))).Uint32()).(bool)))
					if p1 {
						var _1321_v11 _dafny.Sequence
						_ = _1321_v11
						_1321_v11 = _dafny.SeqOf(_1312_v7, Companion_Default___.Fm36(!(p1), globalState))
						var _pat_let_tv24 = p0
						_ = _pat_let_tv24
						var _pat_let_tv25 = globalState
						_ = _pat_let_tv25
						_1321_v11 = _dafny.SeqOf(func(_pat_let26_0 D2) D2 {
							return func(_1322_dt__update__tmp_h1 D2) D2 {
								return func(_pat_let27_0 _dafny.Int) D2 {
									return func(_1323_dt__update_hcf6_h0 _dafny.Int) D2 {
										return func(_pat_let28_0 _dafny.Int) D2 {
											return func(_1324_dt__update_hcf5_h0 _dafny.Int) D2 {
												return Companion_D2_.Create_DC4_(_1324_dt__update_hcf5_h0, _1323_dt__update_hcf6_h0, (_1322_dt__update__tmp_h1).Dtor_cf7(), (_1322_dt__update__tmp_h1).Dtor_cf8(), (_1322_dt__update__tmp_h1).Dtor_cf9())
											}(_pat_let28_0)
										}(Companion_Default___.Fm1(_pat_let_tv25))
									}(_pat_let27_0)
								}((_pat_let_tv24).F36())
							}(_pat_let26_0)
						}(_1312_v7), _1312_v7, _1312_v7)
						var _1325_v12 _dafny.Set
						_ = _1325_v12
						_1325_v12 = _dafny.SetOf(_1313_v8.F35)
						var _1326_v13 T2
						_ = _1326_v13
						var _nw220 *C2 = New_C2_()
						_ = _nw220
						_nw220.Ctor__(_1318_v10, (p0).F36())
						_1326_v13 = _nw220
						var _1327_v14 _dafny.Map
						_ = _1327_v14
						_1327_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1313_v8, _1326_v13)
						var _1328_v15 _dafny.Sequence
						_ = _1328_v15
						_1328_v15 = _dafny.SeqOf((p0).F37())
						var _1329_v16 _dafny.Array
						_ = _1329_v16
						var _nwElement0_52 _dafny.Int = _dafny.IntOfUint32(((p0).F37()).Cardinality())
						_ = _nwElement0_52
						var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(11))
						_ = _nw221
						(_nw221).ArraySet1(_nwElement0_52, 0)
						(_nw221).ArraySet1((p0).F36(), 1)
						(_nw221).ArraySet1(_dafny.IntOfUint32(((_1314_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))).Int()).(_dafny.Sequence)).Cardinality()), 2)
						(_nw221).ArraySet1(_dafny.IntOfUint32(((p0).F37()).Cardinality()), 3)
						(_nw221).ArraySet1((_this).F27(), 4)
						(_nw221).ArraySet1((p0).F36(), 5)
						(_nw221).ArraySet1((_1326_v13).F32(), 6)
						(_nw221).ArraySet1((p0).F36(), 7)
						(_nw221).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(557))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg87 _dafny.Int) interface{} {
								return coer87(arg87)
							}
						}((func(_1330_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1331_i2 _dafny.Int) _dafny.CodePoint {
								return _1330_v4
							}
						})(_1308_v4)))).Cardinality()), 8)
						(_nw221).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg88 _dafny.Int) interface{} {
								return coer88(arg88)
							}
						}((func(_1332_p0 *C0) func(_dafny.Int) _dafny.Int {
							return func(_1333_i3 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg89 _dafny.Int) interface{} {
										return coer89(arg89)
									}
								}((func(_1334_p0 *C0) func(_dafny.Int) _dafny.CodePoint {
									return func(_1335_i4 _dafny.Int) _dafny.CodePoint {
										return ((_1334_p0).F37()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_1334_p0).F37()).Cardinality()))).Uint32()).(_dafny.CodePoint)
									}
								})(_1332_p0)))).Cardinality())
							}
						})(p0)))).Cardinality()), 9)
						(_nw221).ArraySet1((p0).F36(), 10)
						_1329_v16 = _nw221
						var _1336_v17 _dafny.Map
						_ = _1336_v17
						_1336_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1305_v1, _1329_v16)
						var _1337_v18 _dafny.Sequence
						_ = _1337_v18
						_1337_v18 = _dafny.SeqOf((func() _dafny.Array {
							if (_1336_v17).Contains((_1314_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))).Int()).(_dafny.Sequence)) {
								return (_1336_v17).Get((_1314_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1314_v9), 0))).Int()).(_dafny.Sequence)).(_dafny.Array)
							}
							return _1329_v16
						})(), _1329_v16, _1329_v16)
						var _1338_v19 _dafny.Map
						_ = _1338_v19
						_1338_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).F36(), (_1326_v13).F32())
						var _1339_v20 _dafny.Map
						_ = _1339_v20
						_1339_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1313_v8.F35, (_1338_v19).Cardinality())
						var _1340_v21 _dafny.MultiSet
						_ = _1340_v21
						_1340_v21 = _dafny.MultiSetOf((_1337_v18).Select((Companion_Default___.SafeIndex((_1339_v20).Cardinality(), _dafny.IntOfUint32((_1337_v18).Cardinality()))).Uint32()).(_dafny.Array), _1329_v16)
						var _1341_v22 _dafny.MultiSet
						_ = _1341_v22
						_1341_v22 = _dafny.MultiSetOf(_dafny.MultiSetOf(_1329_v16, _1329_v16, _1329_v16), _1340_v21, _1340_v21)
						var _1342_v23 _dafny.Map
						_ = _1342_v23
						_1342_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm37((p0).F36(), (_this).Fm34(_1311_v6, Companion_D2_.Create_DC4_((_1325_v12).Cardinality(), (_1327_v14).Cardinality(), _dafny.IntOfInt64(662), _1308_v4, (_1328_v15).Select((Companion_Default___.SafeIndex((p0).F36(), _dafny.IntOfUint32((_1328_v15).Cardinality()))).Uint32()).(_dafny.Sequence)), globalState), globalState), (func() _dafny.Int {
							if (_1341_v22).Contains(_dafny.MultiSetOf(_1329_v16)) {
								return (_1341_v22).Multiplicity(_dafny.MultiSetOf(_1329_v16))
							}
							return (_dafny.Zero).Minus((_this).F27())
						})())
						(globalState).F24 = (_dafny.Zero).Minus((func() _dafny.Int {
							if (_1342_v23).Contains((p0).F37()) {
								return (_1342_v23).Get((p0).F37()).(_dafny.Int)
							}
							return (p0).F36()
						})())
						var _1343_v24 _dafny.Sequence
						_ = _1343_v24
						_1343_v24 = _dafny.UnicodeSeqOfUtf8Bytes("tqatqn")
						_1343_v24 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((p0).F37(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jqpon"), (Companion_Default___.SafeIndex((_1326_v13).F32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jqpon")).Cardinality()))).Uint32(), _1308_v4)), _dafny.UnicodeSeqOfUtf8Bytes("qq"))
						r0 = _dafny.IntOfInt64(695)
						(globalState).F26 = _1313_v8.F35
					} else {
						var _1344_v25 _dafny.Sequence
						_ = _1344_v25
						_1344_v25 = _dafny.UnicodeSeqOfUtf8Bytes("kjg")
						var _1345_v26 _dafny.Sequence
						_ = _1345_v26
						_1345_v26 = _dafny.SeqOf(_dafny.SeqOf(_1308_v4, _dafny.CodePoint('a')))
						var _1346_v27 _dafny.Map
						_ = _1346_v27
						_1346_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1308_v4, _dafny.IntOfInt64(985))
						var _1347_v28 _dafny.MultiSet
						_ = _1347_v28
						_1347_v28 = _dafny.MultiSetOf((_1346_v27).Cardinality())
						var _1348_v29 D9
						_ = _1348_v29
						_1348_v29 = Companion_D9_.Create_DC19_(p1)
						var _1349_v30 _dafny.Array
						_ = _1349_v30
						var _nwElement0_53 bool = _1313_v8.F35
						_ = _nwElement0_53
						var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(17))
						_ = _nw222
						(_nw222).ArraySet1(_nwElement0_53, 0)
						(_nw222).ArraySet1((func() bool {
							if p1 {
								return false
							}
							return _1313_v8.F35
						})(), 1)
						(_nw222).ArraySet1(true, 2)
						(_nw222).ArraySet1(false, 3)
						(_nw222).ArraySet1(_1313_v8.F35, 4)
						(_nw222).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1345_v26, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1345_v26).Cardinality()))).Uint32(), (p0).F37()), _dafny.SeqOf(_1308_v4)), 5)
						(_nw222).ArraySet1((func() bool {
							if p1 {
								return (_this).Fm35(_1308_v4, Companion_Default___.Fm1(globalState), globalState)
							}
							return _1313_v8.F35
						})(), 6)
						(_nw222).ArraySet1(true, 7)
						(_nw222).ArraySet1(((p0).F36()).Cmp((p0).F36()) != 0, 8)
						(_nw222).ArraySet1((_1347_v28).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1310_v5)), 9)
						(_nw222).ArraySet1(((p0).F36()).Cmp((p0).F36()) >= 0, 10)
						(_nw222).ArraySet1((_1318_v10).Equals(_dafny.MultiSetOf(_1313_v8.F35, true)), 11)
						(_nw222).ArraySet1(!(_1313_v8.F35), 12)
						(_nw222).ArraySet1(p1, 13)
						(_nw222).ArraySet1(_1313_v8.F35, 14)
						(_nw222).ArraySet1((_1348_v29).Dtor_cf33(), 15)
						(_nw222).ArraySet1(false, 16)
						_1349_v30 = _nw222
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1349_v30), 0))
						_ = _index166
						(_1349_v30).ArraySet1(!(p1), (_index166).Int())
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1349_v30), 0))
						_ = _index167
						var _rhs219 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("gpkm")
						_ = _rhs219
						var _rhs220 bool = !(!(p1) || (_1313_v8.F35))
						_ = _rhs220
						var _lhs173 _dafny.Array = _1349_v30
						_ = _lhs173
						var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1349_v30), 0))
						_ = _lhs174
						_1344_v25 = _rhs219
						(_lhs173).ArraySet1(_rhs220, (_lhs174).Int())
						var _1350_v31 *C0
						_ = _1350_v31
						var _nw223 *C0 = New_C0_()
						_ = _nw223
						_nw223.Ctor__(_dafny.IntOfUint32((_1344_v25).Cardinality()), (p0).F37())
						_1350_v31 = _nw223
						(globalState).F0 = Companion_Default___.Fm1(globalState)
						var _1351_v32 D4
						_ = _1351_v32
						_1351_v32 = Companion_D4_.Create_DC8_(_1308_v4, _dafny.IntOfInt64(444), p1)
						var _1352_v33 D11
						_ = _1352_v33
						_1352_v33 = Companion_D11_.Create_DC24_((_1350_v31).F36(), _1351_v32)
						_1352_v33 = _1352_v33
						var _1353_v34 D13
						_ = _1353_v34
						_1353_v34 = Companion_D13_.Create_DC31_((p0).F36())
						var _1354_v35 D13
						_ = _1354_v35
						_1354_v35 = Companion_D13_.Create_DC33_(_1353_v34)
						_1354_v35 = _1354_v35
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		(globalState).F3 = (_this).F27()
		(globalState).F21 = ((p0).F36()).Minus(_dafny.IntOfInt64(-794))
		var _rhs221 bool = !(p1)
		_ = _rhs221
		var _rhs222 bool = p1
		_ = _rhs222
		var _rhs223 bool = !(p1)
		_ = _rhs223
		var _lhs175 *GlobalState = globalState
		_ = _lhs175
		var _lhs176 *GlobalState = globalState
		_ = _lhs176
		var _lhs177 *GlobalState = globalState
		_ = _lhs177
		_lhs175.F26 = _rhs221
		_lhs176.F26 = _rhs222
		_lhs177.F26 = _rhs223
		var _1355_v36 _dafny.Map
		_ = _1355_v36
		_1355_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F27())
		var _1356_v37 _dafny.Set
		_ = _1356_v37
		_1356_v37 = _dafny.SetOf(_1355_v36, (_1355_v36).Update(true, _dafny.IntOfInt64(677)), (_1355_v36).Update(Companion_Default___.Fm0(_dafny.IntOfUint32((_1305_v1).Cardinality()), globalState), (_this).F27()), _1355_v36)
		var _hi5 _dafny.Int = ((_1356_v37).Difference(_1356_v37)).Cardinality()
		_ = _hi5
		for _1357_i5 := Companion_Default___.SafeModuloInt((_this).F27(), (_this).F27()); _1357_i5.Cmp(_hi5) < 0; _1357_i5 = _1357_i5.Plus(_dafny.One) {
			if !(((p0).F36()).Cmp(_dafny.IntOfInt64(431)) == 0) {
				var _1358_v38 _dafny.Map
				_ = _1358_v38
				_1358_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).F36(), _1305_v1)
				_1358_v38 = (_1358_v38).Update(((_this).F27()).Minus(_1357_i5), _dafny.SeqOf(p1))
				var _1359_v39 _dafny.Map
				_ = _1359_v39
				_1359_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).F36(), Companion_Default___.Fm0((p0).F36(), globalState))
				var _1360_v40 _dafny.Map
				_ = _1360_v40
				_1360_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1359_v39)
				var _1361_v41 _dafny.Array
				_ = _1361_v41
				var _nwElement0_54 _dafny.Map = _1360_v40
				_ = _nwElement0_54
				var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(4))
				_ = _nw224
				(_nw224).ArraySet1(_nwElement0_54, 0)
				(_nw224).ArraySet1((_1360_v40).Merge(_1360_v40), 1)
				(_nw224).ArraySet1(_1360_v40, 2)
				(_nw224).ArraySet1((_1360_v40).Update(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1357_i5), false)), 3)
				_1361_v41 = _nw224
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1361_v41), 0))
				_ = _index168
				(_1361_v41).ArraySet1(_1360_v40, (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1361_v41), 0))
				_ = _index169
				(_1361_v41).ArraySet1(((_1360_v40).Merge(_1360_v40)).Merge(_1360_v40), (_index169).Int())
				var _1362_v42 _dafny.MultiSet
				_ = _1362_v42
				_1362_v42 = _dafny.MultiSetOf(_dafny.IntOfInt64(422))
				var _1363_v43 *C0
				_ = _1363_v43
				var _nw225 *C0 = New_C0_()
				_ = _nw225
				_nw225.Ctor__((func() _dafny.Int {
					if p1 {
						return (_dafny.Zero).Minus((p0).F36())
					}
					return (_1362_v42).Cardinality()
				})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg90 _dafny.Int) interface{} {
						return coer90(arg90)
					}
				}(func(_1364_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})))
				_1363_v43 = _nw225
				var _1365_v44 _dafny.Array
				_ = _1365_v44
				var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw226
				_1365_v44 = _nw226
				var _nw227 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw227
				_1365_v44 = _nw227
				var _1366_v45 _dafny.Array
				_ = _1366_v45
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
				_ = _nw228
				_1366_v45 = _nw228
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1366_v45), 0))
				_ = _index170
				(_1366_v45).ArraySet1CodePoint(_dafny.CodePoint('c'), (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1366_v45), 0))
				_ = _index171
				(_1366_v45).ArraySet1CodePoint(Companion_Default___.Fm38(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), globalState), (_index171).Int())
			} else {
				var _1367_v46 _dafny.Array
				_ = _1367_v46
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_24
				var _nw229 _dafny.Array
				_ = _nw229
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw229 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) D4 = (func(_1368_v4 _dafny.CodePoint, _1369_i5 _dafny.Int, _1370_p1 bool) func(_dafny.Int) D4 {
						return func(_1371_i7 _dafny.Int) D4 {
							return Companion_D4_.Create_DC8_(_1368_v4, _1369_i5, _1370_p1)
						}
					})(_1308_v4, _1357_i5, p1)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw229 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw229).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw229).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1367_v46 = _nw229
				var _1372_v47 _dafny.Array
				_ = _1372_v47
				var _nw230 _dafny.Array = _dafny.NewArrayWithValue(Companion_D10_.Default(), _dafny.IntOfInt64(26))
				_ = _nw230
				_1372_v47 = _nw230
				var _1373_v48 _dafny.Sequence
				_ = _1373_v48
				_1373_v48 = _dafny.SeqOf(_dafny.IntOfInt64(320))
				var _1374_v49 D10
				_ = _1374_v49
				_1374_v49 = Companion_D10_.Create_DC21_(_1373_v48)
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1372_v47), 0))
				_ = _index172
				(_1372_v47).ArraySet1(_1374_v49, (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1372_v47), 0))
				_ = _index173
				var _rhs224 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(691), (p0).F36())).Minus((p0).F36())
				_ = _rhs224
				var _rhs225 bool = p1
				_ = _rhs225
				var _rhs226 _dafny.Map = _1304_v0
				_ = _rhs226
				var _rhs227 _dafny.Array = _1367_v46
				_ = _rhs227
				var _rhs228 D10 = _1374_v49
				_ = _rhs228
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 _dafny.Array = _1372_v47
				_ = _lhs180
				var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1372_v47), 0))
				_ = _lhs181
				_lhs178.F21 = _rhs224
				_lhs179.F26 = _rhs225
				_1304_v0 = _rhs226
				_1367_v46 = _rhs227
				(_lhs180).ArraySet1(_rhs228, (_lhs181).Int())
				var _1375_v51 _dafny.Set
				_ = _1375_v51
				_1375_v51 = _dafny.SetOf((p0).F36())
				var _1376_v52 _dafny.Map
				_ = _1376_v52
				_1376_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				var _1377_v53 D2
				_ = _1377_v53
				_1377_v53 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus((_this).F27()), (_1376_v52).Cardinality(), _dafny.IntOfInt64(-286), _1308_v4, (p0).F37())
				var _1378_v54 _dafny.Map
				_ = _1378_v54
				_1378_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).F36(), p1)
				var _1379_v55 _dafny.MultiSet
				_ = _1379_v55
				_1379_v55 = _dafny.MultiSetOf(Companion_Default___.Fm0((_this).F27(), globalState))
				var _1380_v56 _dafny.Array
				_ = _1380_v56
				var _nwElement0_55 _dafny.Int = Companion_Default___.SafeDivisionInt(_1357_i5, _1357_i5)
				_ = _nwElement0_55
				var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(28))
				_ = _nw231
				(_nw231).ArraySet1(_nwElement0_55, 0)
				(_nw231).ArraySet1((_dafny.Zero).Minus(_1357_i5), 1)
				(_nw231).ArraySet1((_1357_i5).Times((p0).F36()), 2)
				(_nw231).ArraySet1(_1357_i5, 3)
				(_nw231).ArraySet1((_1373_v48).Select((Companion_Default___.SafeIndex(_1357_i5, _dafny.IntOfUint32((_1373_v48).Cardinality()))).Uint32()).(_dafny.Int), 4)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_1305_v1).Cardinality()), 5)
				(_nw231).ArraySet1((p0).F36(), 6)
				(_nw231).ArraySet1(Companion_Default___.SafeDivisionInt((p0).F36(), (func() _dafny.Set {
					var _coll45 = _dafny.NewBuilder()
					_ = _coll45
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(865), _dafny.IntOfInt64(146))); ; {
						_compr_45, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1381_v50 _dafny.Int
						_1381_v50 = interface{}(_compr_45).(_dafny.Int)
						if ((_dafny.IntOfInt64(865)).Cmp(_1381_v50) <= 0) && ((_1381_v50).Cmp(_dafny.IntOfInt64(146)) < 0) {
							_coll45.Add((_1381_v50).Times((p0).F36()))
						}
					}
					return _coll45.ToSet()
				}()).Cardinality()), 7)
				(_nw231).ArraySet1((_dafny.Zero).Minus((_this).F27()), 8)
				(_nw231).ArraySet1((p0).F36(), 9)
				(_nw231).ArraySet1((_this).F27(), 10)
				(_nw231).ArraySet1((p0).F36(), 11)
				(_nw231).ArraySet1(((_this).Fm34(_1375_v51, _1377_v53, globalState)).Minus((_dafny.Zero).Minus((_1378_v54).Cardinality())), 12)
				(_nw231).ArraySet1((p0).F36(), 13)
				(_nw231).ArraySet1((func() _dafny.Int {
					if (_1355_v36).Contains(p1) {
						return (_1355_v36).Get(p1).(_dafny.Int)
					}
					return (_this).F27()
				})(), 14)
				(_nw231).ArraySet1((func() _dafny.Int {
					if (_1379_v55).Contains(p1) {
						return (_1379_v55).Multiplicity(p1)
					}
					return (p0).F36()
				})(), 15)
				(_nw231).ArraySet1((_dafny.Zero).Minus((p0).F36()), 16)
				(_nw231).ArraySet1(_1357_i5, 17)
				(_nw231).ArraySet1((_this).F27(), 18)
				(_nw231).ArraySet1((p0).F36(), 19)
				(_nw231).ArraySet1((p0).F36(), 20)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}((func(_1382_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1383_i8 _dafny.Int) _dafny.CodePoint {
						return _1382_v4
					}
				})(_1308_v4)))).Cardinality()), 21)
				(_nw231).ArraySet1(((p0).F36()).Plus((p0).F36()), 22)
				(_nw231).ArraySet1(_1357_i5, 23)
				(_nw231).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(594), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(884))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}(func(_1384_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality())), 24)
				(_nw231).ArraySet1((_this).F27(), 25)
				(_nw231).ArraySet1(_1357_i5, 26)
				(_nw231).ArraySet1((p0).F36(), 27)
				_1380_v56 = _nw231
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1380_v56), 0))
				_ = _index174
				(_1380_v56).ArraySet1(Companion_Default___.SafeDivisionInt((p0).F36(), _dafny.IntOfUint32((_1373_v48).Cardinality())), (_index174).Int())
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1380_v56), 0))
				_ = _index175
				(_1380_v56).ArraySet1(_1357_i5, (_index175).Int())
				(globalState).F3 = _1357_i5
				var _1385_v57 _dafny.Array
				_ = _1385_v57
				var _nw232 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw232
				_1385_v57 = _nw232
				_1385_v57 = (func() _dafny.Array {
					if (_dafny.MultiSetOf(p1, p1, !(Companion_Default___.Fm0(_dafny.IntOfInt64(736), globalState)))).IsSubsetOf(_1379_v55) {
						return _1385_v57
					}
					return _1385_v57
				})()
				var _1386_v58 _dafny.Sequence
				_ = _1386_v58
				_1386_v58 = _dafny.UnicodeSeqOfUtf8Bytes("efdl")
				_1386_v58 = Companion_Default___.Fm37((_this).F27(), _1357_i5, globalState)
			}
			(globalState).F26 = !(p1)
			var _1387_v59 _dafny.Array
			_ = _1387_v59
			var _nw233 _dafny.Array = _dafny.NewArrayWithValue(Companion_D13_.Default(), _dafny.IntOfInt64(20))
			_ = _nw233
			_1387_v59 = _nw233
			var _1388_v60 _dafny.MultiSet
			_ = _1388_v60
			_1388_v60 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1305_v1).Cardinality())), (_this).F27())
			var _1389_v61 _dafny.MultiSet
			_ = _1389_v61
			_1389_v61 = _dafny.MultiSetOf(p1, p1)
			var _1390_v62 D13
			_ = _1390_v62
			_1390_v62 = Companion_D13_.Create_DC32_((func() _dafny.Int {
				if (_1388_v60).Contains((p0).F36()) {
					return (_1388_v60).Multiplicity((p0).F36())
				}
				return _1357_i5
			})(), (_1389_v61).Cardinality(), _dafny.SeqOf(_1357_i5, (_dafny.Zero).Minus((_1304_v0).Cardinality())), p1)
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1387_v59), 0))
			_ = _index176
			(_1387_v59).ArraySet1((func() D13 {
				if p1 {
					return _1390_v62
				}
				return _1390_v62
			})(), (_index176).Int())
			var _1391_v63 _dafny.Set
			_ = _1391_v63
			_1391_v63 = _dafny.SetOf(_1305_v1)
			var _1392_v64 _dafny.Map
			_ = _1392_v64
			_1392_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1391_v63)
			var _1393_v65 _dafny.Set
			_ = _1393_v65
			_1393_v65 = _dafny.SetOf(_1305_v1)
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1387_v59), 0))
			_ = _index177
			var _rhs229 _dafny.Int = ((p0).F36()).Minus((p0).F36())
			_ = _rhs229
			var _rhs230 _dafny.Int = (p0).F36()
			_ = _rhs230
			var _rhs231 bool = (_1393_v65).IsSubsetOf((func() _dafny.Set {
				if (_1392_v64).Contains((p0).F36()) {
					return (_1392_v64).Get((p0).F36()).(_dafny.Set)
				}
				return _dafny.SetOf(_1305_v1)
			})())
			_ = _rhs231
			var _rhs232 D13 = _1390_v62
			_ = _rhs232
			var _lhs182 *GlobalState = globalState
			_ = _lhs182
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			var _lhs184 *GlobalState = globalState
			_ = _lhs184
			var _lhs185 _dafny.Array = _1387_v59
			_ = _lhs185
			var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1387_v59), 0))
			_ = _lhs186
			_lhs182.F0 = _rhs229
			_lhs183.F21 = _rhs230
			_lhs184.F26 = _rhs231
			(_lhs185).ArraySet1(_rhs232, (_lhs186).Int())
			_1308_v4 = _1308_v4
		}
		r0 = (_this).F27()
		return r0
	}
}
func (_this *C4) M12(p0 bool, p1 _dafny.Map, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Array, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 bool = false
		_ = r3
		if !(p2) {
			var _1394_v0 _dafny.Sequence
			_ = _1394_v0
			_1394_v0 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			var _1395_v1 _dafny.Sequence
			_ = _1395_v1
			_1395_v1 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qppgnht"), _1394_v0), Companion_Default___.Fm37((_this).F27(), _dafny.IntOfUint32((_1394_v0).Cardinality()), globalState))
			var _1396_v2 _dafny.Map
			_ = _1396_v2
			_1396_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
			var _1397_v3 _dafny.MultiSet
			_ = _1397_v3
			_1397_v3 = _dafny.MultiSetOf(p3)
			var _1398_v4 _dafny.Sequence
			_ = _1398_v4
			_1398_v4 = _dafny.SeqOf(p0, p2, false, p0)
			var _1399_v5 _dafny.Sequence
			_ = _1399_v5
			_1399_v5 = _dafny.SeqOf(_1398_v4)
			var _1400_v6 D6
			_ = _1400_v6
			_1400_v6 = Companion_D6_.Create_DC12_(p0, _1399_v5, _1394_v0)
			var _rhs233 _dafny.Sequence = _1395_v1
			_ = _rhs233
			var _rhs234 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1400_v6).Dtor_cf20())
			_ = _rhs234
			var _rhs235 _dafny.MultiSet = _1397_v3
			_ = _rhs235
			var _rhs236 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F27(), (Companion_D4_.Create_DC8_(p3, (_dafny.SetOf(p2)).Cardinality(), p2)).Dtor_cf14())
			_ = _rhs236
			var _lhs187 *GlobalState = globalState
			_ = _lhs187
			_1395_v1 = _rhs233
			_1396_v2 = _rhs234
			_1397_v3 = _rhs235
			_lhs187.F0 = _rhs236
			_1394_v0 = _dafny.Companion_Sequence_.Concatenate(_1394_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("m"), _1394_v0))
			_1395_v1 = _1395_v1
			(globalState).F24 = (func() _dafny.Int {
				if !(p0) {
					return (_this).F27()
				}
				return (_this).F27()
			})()
			(globalState).F21 = (_this).F27()
		} else {
			var _1401_v7 _dafny.Array
			_ = _1401_v7
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_25
			var _nw234 _dafny.Array
			_ = _nw234
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw234 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_1402_p0 bool) func(_dafny.Int) bool {
					return func(_1403_i0 _dafny.Int) bool {
						return _1402_p0
					}
				})(p0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw234 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw234).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw234).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1401_v7 = _nw234
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
			_ = _index178
			(_1401_v7).ArraySet1((true) && (p0), (_index178).Int())
			var _1404_v8 _dafny.Map
			_ = _1404_v8
			_1404_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
			_ = _index179
			(_1401_v7).ArraySet1(!(!((func() bool {
				if p2 {
					return p0
				}
				return (_this).Fm35(Companion_Default___.Fm38(_1404_v8, globalState), (_this).F27(), globalState)
			})()) || (false)), (_index179).Int())
			var _1405_v11 _dafny.Sequence
			_ = _1405_v11
			_1405_v11 = _dafny.UnicodeSeqOfUtf8Bytes("g")
			var _1406_v12 D2
			_ = _1406_v12
			_1406_v12 = Companion_D2_.Create_DC4_(_dafny.IntOfUint32((_1405_v11).Cardinality()), (_this).F27(), (_this).F27(), p3, _1405_v11)
			r3 = (((_this).F27()).Plus((_this).F27())).Cmp((_this).Fm34(func() _dafny.Set {
				var _coll46 = _dafny.NewBuilder()
				_ = _coll46
				for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(234), _dafny.IntOfInt64(-710))); ; {
					_compr_46, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _1407_v9 _dafny.Int
					_1407_v9 = interface{}(_compr_46).(_dafny.Int)
					if ((_dafny.IntOfInt64(234)).Cmp(_1407_v9) <= 0) && ((_1407_v9).Cmp(_dafny.IntOfInt64(-710)) < 0) {
						_coll46.Add(Companion_Default___.SafeModuloInt(_1407_v9, (func() _dafny.Map {
							var _coll47 = _dafny.NewMapBuilder()
							_ = _coll47
							for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-807), _dafny.IntOfInt64(827))); ; {
								_compr_47, _ok53 := _iter53()
								if !_ok53 {
									break
								}
								var _1408_v10 _dafny.Int
								_1408_v10 = interface{}(_compr_47).(_dafny.Int)
								if ((_dafny.IntOfInt64(-807)).Cmp(_1408_v10) <= 0) && ((_1408_v10).Cmp(_dafny.IntOfInt64(827)) < 0) {
									_coll47.Add(Companion_Default___.SafeDivisionInt(_1408_v10, (_this).F27()), p3)
								}
							}
							return _coll47.ToMap()
						}()).Cardinality()))
					}
				}
				return _coll46.ToSet()
			}(), _1406_v12, globalState)) == 0
			if p2 {
				var _1409_v13 D11
				_ = _1409_v13
				_1409_v13 = Companion_D11_.Create_DC25_(Companion_Default___.Fm39(globalState))
				var _1410_v14 _dafny.Array
				_ = _1410_v14
				var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw235
				_1410_v14 = _nw235
				var _1411_v15 _dafny.Sequence
				_ = _1411_v15
				_1411_v15 = _dafny.SeqOf(p0, p2)
				var _1412_v16 _dafny.Map
				_ = _1412_v16
				_1412_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v11, _1411_v15)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1410_v14), 0))
				_ = _index180
				(_1410_v14).ArraySet1((_1412_v16).Update(_1405_v11, _1411_v15), (_index180).Int())
				var _1413_v17 _dafny.Map
				_ = _1413_v17
				_1413_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F27())
				var _1414_v18 D11
				_ = _1414_v18
				_1414_v18 = Companion_D11_.Create_DC24_(_dafny.IntOfInt64(166), Companion_D4_.Create_DC8_(p3, (_1413_v17).Cardinality(), p2))
				var _1415_v19 _dafny.MultiSet
				_ = _1415_v19
				_1415_v19 = _dafny.MultiSetOf(p0)
				var _1416_v20 *C2
				_ = _1416_v20
				var _nw236 *C2 = New_C2_()
				_ = _nw236
				_nw236.Ctor__(_1415_v19, (_dafny.IntOfUint32((_1405_v11).Cardinality())).Minus((_this).F27()))
				_1416_v20 = _nw236
				var _1417_v21 _dafny.MultiSet
				_ = _1417_v21
				_1417_v21 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(437))).Cardinality())
				var _1418_v22 _dafny.Set
				_ = _1418_v22
				_1418_v22 = _dafny.SetOf((func() _dafny.Int {
					if (_1417_v21).Contains(_dafny.IntOfUint32((_1405_v11).Cardinality())) {
						return (_1417_v21).Multiplicity(_dafny.IntOfUint32((_1405_v11).Cardinality()))
					}
					return (_this).F27()
				})(), (_this).F27())
				var _1419_v23 _dafny.MultiSet
				_ = _1419_v23
				_1419_v23 = _dafny.MultiSetOf(_1418_v22, _1418_v22)
				var _1420_v24 D11
				_ = _1420_v24
				_1420_v24 = Companion_D11_.Create_DC23_(_1419_v23)
				var _1421_v25 D11
				_ = _1421_v25
				_1421_v25 = Companion_D11_.Create_DC25_(_1420_v24)
				var _1422_v26 D12
				_ = _1422_v26
				_1422_v26 = Companion_D12_.Create_DC26_(_1416_v20)
				var _1423_v27 _dafny.Sequence
				_ = _1423_v27
				_1423_v27 = _dafny.SeqOf((_1422_v26).Dtor_cf46(), _1416_v20, _1416_v20)
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1410_v14), 0))
				_ = _index181
				var _rhs237 D11 = Companion_D11_.Create_DC25_(_1421_v25)
				_ = _rhs237
				var _rhs238 _dafny.Map = Companion_Default___.Fm40(_dafny.IntOfInt64(154), p0, !((_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool)) || (p2), Companion_Default___.Fm41((_this).F27(), (_this).F27(), p2, globalState), globalState)
				_ = _rhs238
				var _rhs239 D11 = Companion_Default___.Fm39(globalState)
				_ = _rhs239
				var _rhs240 _dafny.Int = _dafny.IntOfInt64(-619)
				_ = _rhs240
				var _rhs241 *C2 = (_1423_v27).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1423_v27).Cardinality()))).Uint32()).(*C2)
				_ = _rhs241
				var _lhs188 _dafny.Array = _1410_v14
				_ = _lhs188
				var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1410_v14), 0))
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				_1409_v13 = _rhs237
				(_lhs188).ArraySet1(_rhs238, (_lhs189).Int())
				_1414_v18 = _rhs239
				_lhs190.F3 = _rhs240
				_1416_v20 = _rhs241
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
				_ = _index182
				(_1401_v7).ArraySet1(p2, (_index182).Int())
				var _1424_v28 _dafny.Sequence
				_ = _1424_v28
				_1424_v28 = _dafny.SeqOf(_1405_v11, _1405_v11)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
				_ = _index183
				var _rhs242 _dafny.Sequence = (_1424_v28).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1424_v28).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs242
				var _rhs243 _dafny.Int = (_this).F27()
				_ = _rhs243
				var _rhs244 _dafny.Int = _dafny.IntOfInt64(-720)
				_ = _rhs244
				var _rhs245 bool = !((_dafny.IntOfInt64(848)).Cmp((_this).F27()) < 0)
				_ = _rhs245
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				var _lhs193 _dafny.Array = _1401_v7
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
				_ = _lhs194
				_1405_v11 = _rhs242
				_lhs191.F15 = _rhs243
				_lhs192.F0 = _rhs244
				(_lhs193).ArraySet1(_rhs245, (_lhs194).Int())
				var _1425_v29 *C0
				_ = _1425_v29
				var _nw237 *C0 = New_C0_()
				_ = _nw237
				_nw237.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1411_v15, _1411_v15)).Cardinality()), _1405_v11)
				_1425_v29 = _nw237
				var _1426_v30 _dafny.Map
				_ = _1426_v30
				_1426_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(754), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(478), true)).Cardinality())
				var _1427_v31 _dafny.Map
				_ = _1427_v31
				_1427_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1425_v29).F36(), p2)
				var _1428_v32 D9
				_ = _1428_v32
				_1428_v32 = Companion_D9_.Create_DC20_(_dafny.IntOfUint32((_1411_v15).Cardinality()), (func() bool {
					if (_1427_v31).Contains((_1425_v29).F36()) {
						return (_1427_v31).Get((_1425_v29).F36()).(bool)
					}
					return (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool)
				})(), p0, (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool), _1426_v30)
				var _1429_v34 _dafny.Sequence
				_ = _1429_v34
				_1429_v34 = _dafny.SeqOf((_1425_v29).F36(), Companion_Default___.Fm1(globalState))
				var _1430_v37 _dafny.Sequence
				_ = _1430_v37
				_1430_v37 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
					var _coll48 = _dafny.NewBuilder()
					_ = _coll48
					for _iter54 := _dafny.Iterate((_1429_v34).Elements()); ; {
						_compr_48, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1431_v36 _dafny.Int
						_1431_v36 = interface{}(_compr_48).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1429_v34, _1431_v36) {
							_coll48.Add(Companion_Default___.SafeDivisionInt(_1431_v36, _dafny.IntOfInt64(29)))
						}
					}
					return _coll48.ToSet()
				}()).Cardinality(), (_1425_v29).F36()), _1426_v30)
				var _1432_v38 _dafny.Map
				_ = _1432_v38
				_1432_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1411_v15).Cardinality()), p3)
				var _1433_v39 _dafny.Sequence
				_ = _1433_v39
				_1433_v39 = _dafny.SeqOf(_1432_v38)
				var _1434_v40 D15
				_ = _1434_v40
				_1434_v40 = Companion_D15_.Create_DC35_(_1433_v39)
				var _1435_v41 _dafny.Array
				_ = _1435_v41
				var _nwElement0_56 _dafny.Map = (_1426_v30).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1425_v29).F37()).Cardinality()), (_1425_v29).F36()))
				_ = _nwElement0_56
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(29))
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_56, 0)
				(_nw238).ArraySet1((_1428_v32).Dtor_cf38(), 1)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27()), 2)
				(_nw238).ArraySet1((_1426_v30).Merge(func() _dafny.Map {
					var _coll49 = _dafny.NewMapBuilder()
					_ = _coll49
					for _iter55 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1429_v34).Cardinality()), (_1425_v29).F36())).Keys().Elements()); ; {
						_compr_49, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1436_v33 _dafny.Int
						_1436_v33 = interface{}(_compr_49).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1429_v34).Cardinality()), (_1425_v29).F36())).Contains(_1436_v33) {
							_coll49.Add((_1436_v33).Plus((_1425_v29).F36()), _dafny.IntOfInt64(319))
						}
					}
					return _coll49.ToMap()
				}()), 3)
				(_nw238).ArraySet1(_1426_v30, 4)
				(_nw238).ArraySet1((func() _dafny.Map {
					if false {
						return _1426_v30
					}
					return _1426_v30
				})(), 5)
				(_nw238).ArraySet1(_1426_v30, 6)
				(_nw238).ArraySet1((_1426_v30).Merge(func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(32), _dafny.IntOfInt64(657))); ; {
						_compr_50, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1437_v35 _dafny.Int
						_1437_v35 = interface{}(_compr_50).(_dafny.Int)
						if ((_dafny.IntOfInt64(32)).Cmp(_1437_v35) <= 0) && ((_1437_v35).Cmp(_dafny.IntOfInt64(657)) < 0) {
							_coll50.Add(Companion_Default___.SafeModuloInt(_1437_v35, (_1425_v29).F36()), (_1425_v29).F36())
						}
					}
					return _coll50.ToMap()
				}()), 7)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27()), 8)
				(_nw238).ArraySet1(_1426_v30, 9)
				(_nw238).ArraySet1(_1426_v30, 10)
				(_nw238).ArraySet1(_1426_v30, 11)
				(_nw238).ArraySet1(Companion_Default___.Fm42((func() _dafny.Int {
					if (_1415_v19).Contains(p2) {
						return (_1415_v19).Multiplicity(p2)
					}
					return (_this).F27()
				})(), true, (_1425_v29).F36(), globalState), 12)
				(_nw238).ArraySet1(_1426_v30, 13)
				(_nw238).ArraySet1(_1426_v30, 14)
				(_nw238).ArraySet1(_1426_v30, 15)
				(_nw238).ArraySet1((_1426_v30).Merge(_1426_v30), 16)
				(_nw238).ArraySet1(_1426_v30, 17)
				(_nw238).ArraySet1(_1426_v30, 18)
				(_nw238).ArraySet1((_1430_v37).Select((Companion_Default___.SafeIndex(((_dafny.MultiSetFromSeq((_1434_v40).Dtor_cf63())).Update(_1432_v38, Companion_Default___.Abs((_1425_v29).F36()))).Cardinality(), _dafny.IntOfUint32((_1430_v37).Cardinality()))).Uint32()).(_dafny.Map), 19)
				(_nw238).ArraySet1(_1426_v30, 20)
				(_nw238).ArraySet1(_1426_v30, 21)
				(_nw238).ArraySet1((_1426_v30).Merge((_1426_v30).Update((_this).Fm34(_1418_v22, _1406_v12, globalState), (_1425_v29).F36())), 22)
				(_nw238).ArraySet1(_1426_v30, 23)
				(_nw238).ArraySet1((_1426_v30).Merge(((_1428_v32).Dtor_cf38()).Update((_1425_v29).F36(), (_this).F27())), 24)
				(_nw238).ArraySet1(_1426_v30, 25)
				(_nw238).ArraySet1(_1426_v30, 26)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1411_v15).Cardinality()), (_1425_v29).F36()), 27)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_1425_v29).F36()), 28)
				_1435_v41 = _nw238
				var _1438_v42 D16
				_ = _1438_v42
				_1438_v42 = Companion_D16_.Create_DC39_(_1435_v41)
				var _1439_v43 bool
				_ = _1439_v43
				_1439_v43 = (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool)
				var _1440_v44 D9
				_ = _1440_v44
				_1440_v44 = Companion_D9_.Create_DC18_(true)
				var _rhs246 _dafny.Int = (_1425_v29).F36()
				_ = _rhs246
				var _rhs247 _dafny.Array = (_1438_v42).Dtor_cf73()
				_ = _rhs247
				var _rhs248 bool = true
				_ = _rhs248
				var _rhs249 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm43(p2, p2, p0, globalState), _1429_v34)
				_ = _rhs249
				var _rhs250 bool = !(false) || ((_1440_v44).Dtor_cf32())
				_ = _rhs250
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				var _lhs196 *GlobalState = globalState
				_ = _lhs196
				_lhs195.F21 = _rhs246
				_1435_v41 = _rhs247
				_lhs196.F26 = _rhs248
				_1429_v34 = _rhs249
				r1 = _rhs250
			} else {
				var _1441_v45 D9
				_ = _1441_v45
				_1441_v45 = Companion_D9_.Create_DC19_(p0)
				_1441_v45 = (func() D9 {
					if (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool) {
						return _1441_v45
					}
					return _1441_v45
				})()
				var _1442_v46 *C0
				_ = _1442_v46
				var _nw239 *C0 = New_C0_()
				_ = _nw239
				_nw239.Ctor__((_this).F27(), _1405_v11)
				_1442_v46 = _nw239
				var _1443_v47 _dafny.Array
				_ = _1443_v47
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_26
				var _nw240 _dafny.Array
				_ = _nw240
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw240 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Int = (func(_1444_v46 *C0) func(_dafny.Int) _dafny.Int {
						return func(_1445_i1 _dafny.Int) _dafny.Int {
							return (_1445_i1).Minus((_1444_v46).F36())
						}
					})(_1442_v46)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw240 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw240).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw240).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1443_v47 = _nw240
				var _1446_v48 _dafny.Map
				_ = _1446_v48
				_1446_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p2)
				var _1447_v49 _dafny.Map
				_ = _1447_v49
				_1447_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v47, (_1446_v48).Cardinality())
				var _1448_v50 _dafny.Map
				_ = _1448_v50
				_1448_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v49, p0)
				(globalState).F21 = (func() _dafny.Int {
					if (func() bool {
						if (_1448_v50).Contains(_1447_v49) {
							return (_1448_v50).Get(_1447_v49).(bool)
						}
						return p2
					})() {
						return (_this).F27()
					}
					return _dafny.IntOfInt64(729)
				})()
				(globalState).F26 = (_dafny.IntOfUint32((_1405_v11).Cardinality())).Cmp((_1442_v46).F36()) < 0
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))
				_ = _index184
				(_1401_v7).ArraySet1(p0, (_index184).Int())
			}
			r3 = (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool)
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index185
			((_this).F28()).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("naxcrpkh"), (_index185).Int())
			var _1449_v51 _dafny.MultiSet
			_ = _1449_v51
			_1449_v51 = _dafny.MultiSetOf((_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool))
			var _1450_v52 D15
			_ = _1450_v52
			_1450_v52 = Companion_D15_.Create_DC38_((_this).F27(), true, _dafny.IntOfInt64(245))
			var _1451_v53 bool
			_ = _1451_v53
			_1451_v53 = (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool)
			var _1452_v55 D16
			_ = _1452_v55
			_1452_v55 = Companion_D16_.Create_DC40_(!((_1449_v51).IsProperSubsetOf(_1449_v51)), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).Fm34(func() _dafny.Set {
				var _coll51 = _dafny.NewBuilder()
				_ = _coll51
				for _iter57 := _dafny.Iterate((Companion_Default___.Fm43((_1450_v52).Dtor_cf71(), _1451_v53, (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool), globalState)).Elements()); ; {
					_compr_51, _ok57 := _iter57()
					if !_ok57 {
						break
					}
					var _1453_v54 _dafny.Int
					_1453_v54 = interface{}(_compr_51).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm43((_1450_v52).Dtor_cf71(), _1451_v53, (_1401_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1401_v7), 0))).Int()).(bool), globalState), _1453_v54) {
						_coll51.Add((_1453_v54).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dp")).Cardinality()), false)).Cardinality()))
					}
				}
				return _coll51.ToSet()
			}(), _1406_v12, globalState), (_dafny.Zero).Minus((_this).F27()))), (_this).F27())
			var _1454_v56 _dafny.Array
			_ = _1454_v56
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_27
			var _nw241 _dafny.Array
			_ = _nw241
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw241 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Sequence = (func(_1455_p0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1456_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(false, _1455_p0)), _dafny.SeqOf(_dafny.SeqOf(_1455_p0)))
					}
				})(p0)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw241 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw241).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw241).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1454_v56 = _nw241
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_1454_v56), 0))
			_ = _index186
			(_1454_v56).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}((func(_1457_p0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_1458_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_1457_p0)
				}
			})(p0))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg94 _dafny.Int) interface{} {
					return coer94(arg94)
				}
			}((func(_1459_v7 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
				return func(_1460_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_1459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1459_v7), 0))).Int()).(bool))
				}
			})(_1401_v7)))), (_index186).Int())
			var _1461_v57 D4
			_ = _1461_v57
			_1461_v57 = Companion_D4_.Create_DC7_(_1401_v7)
			var _pat_let_tv26 = p2
			_ = _pat_let_tv26
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index187
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_1454_v56), 0))
			_ = _index188
			var _rhs251 _dafny.Sequence = _1405_v11
			_ = _rhs251
			var _rhs252 D16 = func(_pat_let29_0 D16) D16 {
				return func(_1462_dt__update__tmp_h1 D16) D16 {
					return func(_pat_let30_0 bool) D16 {
						return func(_1463_dt__update_hcf74_h0 bool) D16 {
							return func(_pat_let31_0 _dafny.Int) D16 {
								return func(_1464_dt__update_hcf76_h0 _dafny.Int) D16 {
									return Companion_D16_.Create_DC40_(_1463_dt__update_hcf74_h0, (_1462_dt__update__tmp_h1).Dtor_cf75(), _1464_dt__update_hcf76_h0)
								}(_pat_let31_0)
							}(Companion_Default___.SafeDivisionInt((_this).F27(), _dafny.IntOfInt64(527)))
						}(_pat_let30_0)
					}(_pat_let_tv26)
				}(_pat_let29_0)
			}(_1452_v55)
			_ = _rhs252
			var _rhs253 _dafny.Sequence = Companion_Default___.Fm44(globalState)
			_ = _rhs253
			var _rhs254 _dafny.Int = (Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27())).Plus((_dafny.IntOfInt64(155)).Plus((_this).F27()))
			_ = _rhs254
			var _rhs255 _dafny.Array = (_1461_v57).Dtor_cf12()
			_ = _rhs255
			var _lhs197 _dafny.Array = (_this).F28()
			_ = _lhs197
			var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _lhs198
			var _lhs199 _dafny.Array = _1454_v56
			_ = _lhs199
			var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_1454_v56), 0))
			_ = _lhs200
			var _lhs201 *GlobalState = globalState
			_ = _lhs201
			(_lhs197).ArraySet1(_rhs251, (_lhs198).Int())
			_1452_v55 = _rhs252
			(_lhs199).ArraySet1(_rhs253, (_lhs200).Int())
			_lhs201.F15 = _rhs254
			_1401_v7 = _rhs255
		}
		var _1465_v58 _dafny.Map
		_ = _1465_v58
		_1465_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1466_v59 _dafny.Array
		_ = _1466_v59
		var _nwElement0_57 _dafny.CodePoint = Companion_Default___.Fm38(_1465_v58, globalState)
		_ = _nwElement0_57
		var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(25))
		_ = _nw242
		(_nw242).ArraySet1CodePoint(_nwElement0_57, 0)
		(_nw242).ArraySet1CodePoint(p3, 1)
		(_nw242).ArraySet1CodePoint(Companion_Default___.Fm38(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), globalState), 2)
		(_nw242).ArraySet1CodePoint(p3, 3)
		(_nw242).ArraySet1CodePoint(p3, 4)
		(_nw242).ArraySet1CodePoint(p3, 5)
		(_nw242).ArraySet1CodePoint(p3, 6)
		(_nw242).ArraySet1CodePoint(p3, 7)
		(_nw242).ArraySet1CodePoint(p3, 8)
		(_nw242).ArraySet1CodePoint(_dafny.CodePoint('c'), 9)
		(_nw242).ArraySet1CodePoint(p3, 10)
		(_nw242).ArraySet1CodePoint(p3, 11)
		(_nw242).ArraySet1CodePoint(_dafny.CodePoint('g'), 12)
		(_nw242).ArraySet1CodePoint(p3, 13)
		(_nw242).ArraySet1CodePoint(p3, 14)
		(_nw242).ArraySet1CodePoint(p3, 15)
		(_nw242).ArraySet1CodePoint(p3, 16)
		(_nw242).ArraySet1CodePoint(_dafny.CodePoint('b'), 17)
		(_nw242).ArraySet1CodePoint(p3, 18)
		(_nw242).ArraySet1CodePoint(p3, 19)
		(_nw242).ArraySet1CodePoint(p3, 20)
		(_nw242).ArraySet1CodePoint(p3, 21)
		(_nw242).ArraySet1CodePoint(p3, 22)
		(_nw242).ArraySet1CodePoint(Companion_Default___.Fm38(_1465_v58, globalState), 23)
		(_nw242).ArraySet1CodePoint(p3, 24)
		_1466_v59 = _nw242
		for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1466_v59), 0))); ; {
			_guard_loop_5, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _1467_i5 _dafny.Int
			_1467_i5 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1467_i5).Sign() != -1) && ((_1467_i5).Cmp(_dafny.ArrayLen((_1466_v59), 0)) < 0)) {
				(_1466_v59).ArraySet1CodePoint(_dafny.CodePoint('y'), (_1467_i5).Int())
			}
		}
		var _1468_v60 _dafny.Sequence
		_ = _1468_v60
		_1468_v60 = _dafny.SeqOf(p0, p0)
		var _1469_v61 T0
		_ = _1469_v61
		var _nw243 *C3 = New_C3_()
		_ = _nw243
		_nw243.Ctor__((_dafny.IntOfUint32((_1468_v60).Cardinality())).Plus((_this).F27()), (_this).F28())
		_1469_v61 = _nw243
		_1469_v61 = _1469_v61
		var _1470_v62 D12
		_ = _1470_v62
		_1470_v62 = Companion_D12_.Create_DC28_((_this).F27(), (_dafny.Zero).Minus((_1469_v61).F27()), p0)
		var _1471_i6 _dafny.Int
		_ = _1471_i6
		_1471_i6 = _dafny.Zero
		{
			for Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1468_v60, Companion_Default___.Fm41((_1469_v61).F27(), (_1469_v61).F27(), (_1470_v62).Dtor_cf53(), globalState)), (Companion_Default___.SafeIndex((_1469_v61).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1468_v60, Companion_Default___.Fm41((_1469_v61).F27(), (_1469_v61).F27(), (_1470_v62).Dtor_cf53(), globalState))).Cardinality()))).Uint32(), p0)).Cardinality()), globalState) {
				{
					if (_1471_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1471_i6 = (_1471_i6).Plus(_dafny.One)
					if p2 {
						var _1472_v63 *C3
						_ = _1472_v63
						var _nw244 *C3 = New_C3_()
						_ = _nw244
						_nw244.Ctor__(_dafny.IntOfInt64(342), (_1469_v61).F28())
						_1472_v63 = _nw244
						var _1473_v64 _dafny.Sequence
						_ = _1473_v64
						_1473_v64 = _dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus((_1469_v61).F27())))
						_1473_v64 = _1473_v64
						var _1474_v65 _dafny.Sequence
						_ = _1474_v65
						_1474_v65 = _dafny.UnicodeSeqOfUtf8Bytes("r")
						var _1475_v66 _dafny.Set
						_ = _1475_v66
						_1475_v66 = _dafny.SetOf((_1469_v61).F27(), Companion_Default___.Fm1(globalState), (_this).F27(), (_1469_v61).F27())
						var _1476_v67 _dafny.Set
						_ = _1476_v67
						_1476_v67 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_1468_v60, (Companion_Default___.SafeIndex((_1475_v66).Cardinality(), _dafny.IntOfUint32((_1468_v60).Cardinality()))).Uint32(), p2))
						var _1477_v68 _dafny.Array
						_ = _1477_v68
						var _nwElement0_58 _dafny.Int = (_1469_v61).F27()
						_ = _nwElement0_58
						var _nw245 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(10))
						_ = _nw245
						(_nw245).ArraySet1(_nwElement0_58, 0)
						(_nw245).ArraySet1((_1469_v61).F27(), 1)
						(_nw245).ArraySet1((_this).F27(), 2)
						(_nw245).ArraySet1((_1469_v61).F27(), 3)
						(_nw245).ArraySet1((_1469_v61).F27(), 4)
						(_nw245).ArraySet1((_1469_v61).F27(), 5)
						(_nw245).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1474_v65).Cardinality()), (_this).F27()), 6)
						(_nw245).ArraySet1(((_1469_v61).F27()).Plus((_1476_v67).Cardinality()), 7)
						(_nw245).ArraySet1((Companion_Default___.Fm47(globalState)).Cardinality(), 8)
						(_nw245).ArraySet1((_1469_v61).F27(), 9)
						_1477_v68 = _nw245
						var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1477_v68), 0))
						_ = _index189
						(_1477_v68).ArraySet1(Companion_Default___.SafeDivisionInt((_1469_v61).F27(), _dafny.IntOfInt64(595)), (_index189).Int())
						var _1478_v69 _dafny.Array
						_ = _1478_v69
						var _nw246 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
						_ = _nw246
						_1478_v69 = _nw246
						var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1478_v69), 0))
						_ = _index190
						(_1478_v69).ArraySet1(!(p0) || (p0), (_index190).Int())
						var _1479_v70 _dafny.Map
						_ = _1479_v70
						_1479_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1469_v61).F27())
						var _1480_v71 _dafny.Map
						_ = _1480_v71
						_1480_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1479_v70).Cardinality())
						var _1481_v72 _dafny.Map
						_ = _1481_v72
						_1481_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1469_v61).F27(), p2)
						var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1477_v68), 0))
						_ = _index191
						var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1478_v69), 0))
						_ = _index192
						var _rhs256 _dafny.Int = (_1469_v61).F27()
						_ = _rhs256
						var _rhs257 bool = ((_dafny.Zero).Minus((_this).F27())).Cmp((((_1479_v70).Update(false, (_1469_v61).F27())).Update(p0, (_1481_v72).Cardinality())).Cardinality()) > 0
						_ = _rhs257
						var _rhs258 bool = p0
						_ = _rhs258
						var _rhs259 bool = Companion_Default___.Fm0((func() _dafny.Int {
							if p0 {
								return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nngcbxfxb")).Cardinality())
							}
							return (_1469_v61).F27()
						})(), globalState)
						_ = _rhs259
						var _rhs260 bool = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus((_1469_v61).F27()))).Cardinality()).Cmp((_1469_v61).F27()) > 0
						_ = _rhs260
						var _lhs202 _dafny.Array = _1477_v68
						_ = _lhs202
						var _lhs203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1477_v68), 0))
						_ = _lhs203
						var _lhs204 *GlobalState = globalState
						_ = _lhs204
						var _lhs205 *GlobalState = globalState
						_ = _lhs205
						var _lhs206 _dafny.Array = _1478_v69
						_ = _lhs206
						var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1478_v69), 0))
						_ = _lhs207
						(_lhs202).ArraySet1(_rhs256, (_lhs203).Int())
						_lhs204.F26 = _rhs257
						r1 = _rhs258
						_lhs205.F26 = _rhs259
						(_lhs206).ArraySet1(_rhs260, (_lhs207).Int())
						_1477_v68 = _1477_v68
						(globalState).F0 = Companion_Default___.SafeModuloInt((_1477_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1477_v68), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(477))
					} else {
						var _1482_v73 _dafny.Map
						_ = _1482_v73
						_1482_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm1(globalState), (_this).F27())), p0)
						_1482_v73 = (Companion_Default___.Fm48(p2, p2, ((_dafny.MultiSetOf((_1469_v61).F27(), (_dafny.Zero).Minus((_1469_v61).F27()))).Update((_1469_v61).F27(), Companion_Default___.Abs(_dafny.IntOfInt64(-216)))).Cardinality(), globalState)).Merge(_1482_v73)
						var _1483_v74 _dafny.Sequence
						_ = _1483_v74
						_1483_v74 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(33))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg95 _dafny.Int) interface{} {
								return coer95(arg95)
							}
						}((func(_1484_v60 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
							return func(_1485_i7 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_1484_v60).Cardinality())
							}
						})(_1468_v60))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(942))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg96 _dafny.Int) interface{} {
								return coer96(arg96)
							}
						}((func(_1486_v61 T0) func(_dafny.Int) _dafny.Int {
							return func(_1487_i8 _dafny.Int) _dafny.Int {
								return (_1486_v61).F27()
							}
						})(_1469_v61))))
						var _rhs261 _dafny.Int = (func() _dafny.Int {
							if p0 {
								return Companion_Default___.Fm1(globalState)
							}
							return (_dafny.MultiSetFromSeq((_1483_v74).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1483_v74).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()
						})()
						_ = _rhs261
						var _rhs262 bool = p2
						_ = _rhs262
						var _lhs208 *GlobalState = globalState
						_ = _lhs208
						var _lhs209 *GlobalState = globalState
						_ = _lhs209
						_lhs208.F15 = _rhs261
						_lhs209.F26 = _rhs262
						var _1488_v75 _dafny.Sequence
						_ = _1488_v75
						_1488_v75 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cscw"))
						var _1489_v76 _dafny.Sequence
						_ = _1489_v76
						_1489_v76 = _dafny.SeqOf((_1469_v61).F27(), _dafny.IntOfInt64(892), (_1469_v61).F27())
						var _1490_v77 _dafny.MultiSet
						_ = _1490_v77
						_1490_v77 = _dafny.MultiSetOf(p2)
						var _1491_v78 _dafny.Set
						_ = _1491_v78
						_1491_v78 = _dafny.SetOf((_this).F27())
						var _1492_v79 _dafny.Map
						_ = _1492_v79
						_1492_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1469_v61).F27())
						var _rhs263 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg97 _dafny.Int) interface{} {
								return coer97(arg97)
							}
						}(func(_1493_i9 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qwi"), _dafny.UnicodeSeqOfUtf8Bytes("aqpplff"))
						}))
						_ = _rhs263
						var _rhs264 bool = ((_1490_v77).Union(_1490_v77)).IsDisjointFrom((_1490_v77).Update(p0, Companion_Default___.Abs((_1469_v61).F27())))
						_ = _rhs264
						var _rhs265 _dafny.Sequence = _1489_v76
						_ = _rhs265
						var _rhs266 bool = !((_1491_v78).IsDisjointFrom(_1491_v78))
						_ = _rhs266
						var _rhs267 _dafny.Int = ((_1492_v79).Cardinality()).Minus((_this).F27())
						_ = _rhs267
						var _lhs210 *GlobalState = globalState
						_ = _lhs210
						_1488_v75 = _rhs263
						r1 = _rhs264
						_1489_v76 = _rhs265
						r1 = _rhs266
						_lhs210.F0 = _rhs267
						(globalState).F18 = ((_dafny.Zero).Minus((_dafny.IntOfInt64(866)).Minus((_dafny.Zero).Minus(Companion_Default___.Fm1(globalState))))).Times((_1469_v61).F27())
						var _1494_v80 bool
						_ = _1494_v80
						var _1495_v81 _dafny.Int
						_ = _1495_v81
						var _out42 bool
						_ = _out42
						var _out43 _dafny.Int
						_ = _out43
						_out42, _out43 = Companion_Default___.M0(p0, (_1469_v61).F27(), (_1469_v61).F27(), (_1469_v61).F27(), globalState)
						_1494_v80 = _out42
						_1495_v81 = _out43
					}
					var _1496_v82 _dafny.Sequence
					_ = _1496_v82
					_1496_v82 = _dafny.SeqOf((_1469_v61).F27(), (_1469_v61).F27())
					var _1497_v83 _dafny.Sequence
					_ = _1497_v83
					_1497_v83 = _dafny.SeqOf(_1496_v82)
					(globalState).F3 = (_dafny.IntOfInt64(125)).Times(((_1469_v61).F27()).Times(_dafny.IntOfUint32((_1497_v83).Cardinality())))
					var _1498_v84 _dafny.Map
					_ = _1498_v84
					_1498_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), true)
					var _1499_v85 _dafny.Sequence
					_ = _1499_v85
					_1499_v85 = _dafny.SeqOf(_1498_v84, _1498_v84)
					var _1500_v86 _dafny.Sequence
					_ = _1500_v86
					_1500_v86 = _dafny.UnicodeSeqOfUtf8Bytes("cyisvag")
					var _1501_v87 _dafny.MultiSet
					_ = _1501_v87
					_1501_v87 = _dafny.MultiSetOf(p2)
					var _1502_v88 _dafny.MultiSet
					_ = _1502_v88
					_1502_v88 = _dafny.MultiSetOf((_this).F27(), _dafny.IntOfUint32((_1468_v60).Cardinality()))
					var _1503_v89 _dafny.Array
					_ = _1503_v89
					var _nwElement0_59 _dafny.Int = (_1469_v61).F27()
					_ = _nwElement0_59
					var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(27))
					_ = _nw247
					(_nw247).ArraySet1(_nwElement0_59, 0)
					(_nw247).ArraySet1((((_1499_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.IntOfUint32((_1499_v85).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Times((_1469_v61).F27()), 1)
					(_nw247).ArraySet1((_1469_v61).F27(), 2)
					(_nw247).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm41((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)).Cardinality(), (_this).F27(), p0, globalState)).Cardinality()), (_dafny.Zero).Minus((_this).F27())), 3)
					(_nw247).ArraySet1((_1469_v61).F27(), 4)
					(_nw247).ArraySet1((_1469_v61).F27(), 5)
					(_nw247).ArraySet1(_dafny.IntOfInt64(-541), 6)
					(_nw247).ArraySet1((_1469_v61).F27(), 7)
					(_nw247).ArraySet1((_1469_v61).F27(), 8)
					(_nw247).ArraySet1((_this).F27(), 9)
					(_nw247).ArraySet1((_1469_v61).F27(), 10)
					(_nw247).ArraySet1((_this).F27(), 11)
					(_nw247).ArraySet1((_this).F27(), 12)
					(_nw247).ArraySet1((_1469_v61).F27(), 13)
					(_nw247).ArraySet1((_1469_v61).F27(), 14)
					(_nw247).ArraySet1((_this).F27(), 15)
					(_nw247).ArraySet1((_this).F27(), 16)
					(_nw247).ArraySet1((Companion_Default___.Fm1(globalState)).Times(_dafny.IntOfUint32((_1500_v86).Cardinality())), 17)
					(_nw247).ArraySet1((_dafny.SetOf((_1468_v60).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_1468_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.IntOfUint32((_1468_v60).Cardinality()))).Uint32()).(bool), p0)).Cardinality(), _dafny.IntOfUint32((_1468_v60).Cardinality()))).Uint32()).(bool), p2)).Cardinality(), 18)
					(_nw247).ArraySet1(((_1469_v61).F27()).Plus(_dafny.IntOfUint32((_1496_v82).Cardinality())), 19)
					(_nw247).ArraySet1((_this).F27(), 20)
					(_nw247).ArraySet1((_1469_v61).F27(), 21)
					(_nw247).ArraySet1((_dafny.Zero).Minus((_1469_v61).F27()), 22)
					(_nw247).ArraySet1(((_this).F27()).Times((_1469_v61).F27()), 23)
					(_nw247).ArraySet1(_dafny.IntOfInt64(-743), 24)
					(_nw247).ArraySet1((func() _dafny.Int {
						if (_1501_v87).Contains(p2) {
							return (_1501_v87).Multiplicity(p2)
						}
						return (func() _dafny.Int {
							if (_1502_v88).Contains((_1469_v61).F27()) {
								return (_1502_v88).Multiplicity((_1469_v61).F27())
							}
							return (_1501_v87).Cardinality()
						})()
					})(), 25)
					(_nw247).ArraySet1((_1469_v61).F27(), 26)
					_1503_v89 = _nw247
					var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1503_v89), 0))
					_ = _index193
					(_1503_v89).ArraySet1((_this).F27(), (_index193).Int())
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1503_v89), 0))
					_ = _index194
					(_1503_v89).ArraySet1(_dafny.IntOfUint32((_1496_v82).Cardinality()), (_index194).Int())
					var _1504_v90 _dafny.Set
					_ = _1504_v90
					_1504_v90 = _dafny.SetOf(p2, p0, p0)
					(globalState).F0 = Companion_Default___.SafeDivisionInt((_1469_v61).F27(), (func() _dafny.Int {
						if (_1502_v88).Contains((_1503_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1503_v89), 0))).Int()).(_dafny.Int)) {
							return (_1502_v88).Multiplicity((_1503_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1503_v89), 0))).Int()).(_dafny.Int))
						}
						return (_1504_v90).Cardinality()
					})())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1505_v91 _dafny.Sequence
		_ = _1505_v91
		_1505_v91 = _dafny.UnicodeSeqOfUtf8Bytes("cd")
		var _1506_v92 D4
		_ = _1506_v92
		_1506_v92 = Companion_D4_.Create_DC8_(p3, _dafny.IntOfUint32((_1505_v91).Cardinality()), p0)
		(globalState).F24 = (_1506_v92).Dtor_cf14()
		var _1507_v93 _dafny.Array
		_ = _1507_v93
		var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
		_ = _nw248
		_1507_v93 = _nw248
		var _1508_v94 D16
		_ = _1508_v94
		_1508_v94 = Companion_D16_.Create_DC39_(_1507_v93)
		var _1509_v95 _dafny.Array
		_ = _1509_v95
		var _nw249 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw249
		_1509_v95 = _nw249
		var _1510_v96 _dafny.Map
		_ = _1510_v96
		_1510_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1508_v94, _1509_v95)
		_1510_v96 = _1510_v96
		r0 = _1468_v60
		r1 = (Companion_Default___.SafeDivisionInt((_this).F27(), (_dafny.Zero).Minus((_1469_v61).F27()))).Cmp(_dafny.IntOfInt64(987)) < 0
		var _1511_v97 _dafny.Array
		_ = _1511_v97
		var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
		_ = _nw250
		_1511_v97 = _nw250
		var _1512_v98 _dafny.Sequence
		_ = _1512_v98
		_1512_v98 = _dafny.SeqOf(_1511_v97, _1511_v97, _1511_v97, _1511_v97)
		var _1513_v99 _dafny.Sequence
		_ = _1513_v99
		_1513_v99 = _dafny.SeqOf(_dafny.IntOfUint32((_1505_v91).Cardinality()))
		r2 = (_1512_v98).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1513_v99).Cardinality()), _dafny.IntOfUint32((_1512_v98).Cardinality()))).Uint32()).(_dafny.Array)
		var _1514_v100 _dafny.MultiSet
		_ = _1514_v100
		_1514_v100 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1505_v91).Cardinality()))
		var _1515_v101 _dafny.MultiSet
		_ = _1515_v101
		_1515_v101 = _1514_v100
		var _pat_let_tv27 = p2
		_ = _pat_let_tv27
		r3 = func(_source23 _dafny.MultiSet) bool {
			var _1516___mcc_h0 _dafny.MultiSet = _source23
			_ = _1516___mcc_h0
			var _1517_cf18 _dafny.MultiSet = _1516___mcc_h0
			_ = _1517_cf18
			return _pat_let_tv27
		}(_1515_v101)
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f29 _dafny.MultiSet
	_f30 _dafny.Sequence
	_f27 _dafny.Int
	_f28 _dafny.Array
	_f38 _dafny.Map
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f29 = _dafny.EmptyMultiSet
	_this._f30 = _dafny.EmptySeq
	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f38 = _dafny.EmptyMap
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

func (_this *C5) F29() _dafny.MultiSet {
	return _this._f29
}
func (_this *C5) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C5) F27() _dafny.Int {
	return _this._f27
}
func (_this *C5) F28() _dafny.Array {
	return _this._f28
}
func (_this *C5) Ctor__(f38 _dafny.Map, f29 _dafny.MultiSet, f30 _dafny.Sequence, f27 _dafny.Int, f28 _dafny.Array) {
	{
		(_this)._f38 = f38
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C5) Fm3(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !(false) || (true)
	}
}
func (_this *C5) Fm4(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		if ((_this).F29()).Contains(true) {
			return ((_this).F29()).Multiplicity(true)
		} else {
			return (_dafny.Zero).Minus(((_this).F27()).Times((_this).F27()))
		}
	}
}
func (_this *C5) Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		var _source24 D8 = Companion_D8_.Create_DC16_((_this).F27(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(205))).Cardinality()))
		_ = _source24
		if _source24.Is_DC16() {
			var _1518___mcc_h0 _dafny.Int = _source24.Get_().(D8_DC16).Cf29
			_ = _1518___mcc_h0
			var _1519___mcc_h1 _dafny.Int = _source24.Get_().(D8_DC16).Cf30
			_ = _1519___mcc_h1
			var _1520_cf30 _dafny.Int = _1519___mcc_h1
			_ = _1520_cf30
			var _1521_cf29 _dafny.Int = _1518___mcc_h0
			_ = _1521_cf29
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1521_cf29)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_this).F27()))
		} else {
			var _1522___mcc_h2 _dafny.Map = _source24.Get_().(D8_DC15).Cf28
			_ = _1522___mcc_h2
			var _1523_cf28 _dafny.Map = _1522___mcc_h2
			_ = _1523_cf28
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F27())
		}
	}
}
func (_this *C5) M1(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1524_v0 bool
		_ = _1524_v0
		_1524_v0 = true
		var _1525_v1 _dafny.Sequence
		_ = _1525_v1
		_1525_v1 = _dafny.SeqOf(!((((_this).F29()).Update(!(_1524_v0), Companion_Default___.Abs((_this).F27()))).IsDisjointFrom((_this).F29())))
		if (_1525_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(globalState), _dafny.IntOfUint32((_1525_v1).Cardinality()))).Uint32()).(bool) {
			var _1526_v2 _dafny.MultiSet
			_ = _1526_v2
			_1526_v2 = _dafny.MultiSetOf((_this).F27())
			var _1527_v3 _dafny.Map
			_ = _1527_v3
			_1527_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1526_v2).Cardinality(), (_this).F27())
			var _1528_v4 D9
			_ = _1528_v4
			_1528_v4 = Companion_D9_.Create_DC20_((_this).F27(), _1524_v0, _1524_v0, _1524_v0, _1527_v3)
			var _pat_let_tv28 = _1524_v0
			_ = _pat_let_tv28
			var _pat_let_tv29 = _1524_v0
			_ = _pat_let_tv29
			var _pat_let_tv30 = _1524_v0
			_ = _pat_let_tv30
			var _pat_let_tv31 = _1524_v0
			_ = _pat_let_tv31
			var _source25 D9 = func(_pat_let32_0 D9) D9 {
				return func(_1531_dt__update__tmp_h1 D9) D9 {
					return func(_pat_let35_0 bool) D9 {
						return func(_1532_dt__update_hcf37_h0 bool) D9 {
							return func(_pat_let36_0 _dafny.Int) D9 {
								return func(_1533_dt__update_hcf34_h0 _dafny.Int) D9 {
									return func(_pat_let37_0 bool) D9 {
										return func(_1534_dt__update_hcf35_h0 bool) D9 {
											return Companion_D9_.Create_DC20_(_1533_dt__update_hcf34_h0, _1534_dt__update_hcf35_h0, (_1531_dt__update__tmp_h1).Dtor_cf36(), _1532_dt__update_hcf37_h0, (_1531_dt__update__tmp_h1).Dtor_cf38())
										}(_pat_let37_0)
									}((func() bool {
										if _pat_let_tv29 {
											return _pat_let_tv30
										}
										return _pat_let_tv31
									})())
								}(_pat_let36_0)
							}((_this).F27())
						}(_pat_let35_0)
					}(!(false))
				}(_pat_let32_0)
			}(func(_pat_let33_0 D9) D9 {
				return func(_1529_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let34_0 bool) D9 {
						return func(_1530_dt__update_hcf36_h0 bool) D9 {
							return Companion_D9_.Create_DC20_((_1529_dt__update__tmp_h0).Dtor_cf34(), (_1529_dt__update__tmp_h0).Dtor_cf35(), _1530_dt__update_hcf36_h0, (_1529_dt__update__tmp_h0).Dtor_cf37(), (_1529_dt__update__tmp_h0).Dtor_cf38())
						}(_pat_let34_0)
					}(_pat_let_tv28)
				}(_pat_let33_0)
			}(_1528_v4))
			_ = _source25
			if _source25.Is_DC18() {
				var _1535___mcc_h0 bool = _source25.Get_().(D9_DC18).Cf32
				_ = _1535___mcc_h0
				var _1536_cf32 bool = _1535___mcc_h0
				_ = _1536_cf32
				var _1537_v5 _dafny.Sequence
				_ = _1537_v5
				_1537_v5 = _dafny.UnicodeSeqOfUtf8Bytes("fq")
				var _1538_v6 *C0
				_ = _1538_v6
				var _nw251 *C0 = New_C0_()
				_ = _nw251
				_nw251.Ctor__((_this).F27(), _1537_v5)
				_1538_v6 = _nw251
				var _1539_v7 D6
				_ = _1539_v7
				_1539_v7 = Companion_D6_.Create_DC11_(_1538_v6)
				var _1540_v8 _dafny.Map
				_ = _1540_v8
				_1540_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1536_cf32, _1539_v7)
				var _1541_v9 _dafny.Map
				_ = _1541_v9
				_1541_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1540_v8)
				var _rhs268 _dafny.Int = (_this).F27()
				_ = _rhs268
				var _rhs269 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F27(), ((_this).F27()).Plus((_this).F27()))
				_ = _rhs269
				var _rhs270 bool = ((_this).F27()).Cmp((_this).F27()) > 0
				_ = _rhs270
				var _rhs271 bool = (_1541_v9).Contains(((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs271
				var _rhs272 _dafny.Int = (_1538_v6).F36()
				_ = _rhs272
				var _lhs211 *GlobalState = globalState
				_ = _lhs211
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				_lhs211.F18 = _rhs268
				_lhs212.F0 = _rhs269
				_lhs213.F26 = _rhs270
				_1524_v0 = _rhs271
				_lhs214.F18 = _rhs272
				var _1542_v10 _dafny.Array
				_ = _1542_v10
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_28
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) bool = (func(_1543_v0 bool) func(_dafny.Int) bool {
						return func(_1544_i0 _dafny.Int) bool {
							return _1543_v0
						}
					})(_1524_v0)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw252 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw252).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw252).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1542_v10 = _nw252
				var _1545_v11 _dafny.Set
				_ = _1545_v11
				_1545_v11 = _dafny.SetOf((_1538_v6).F36())
				var _1546_v12 _dafny.MultiSet
				_ = _1546_v12
				_1546_v12 = _dafny.MultiSetOf(_1545_v11)
				var _1547_v13 D11
				_ = _1547_v13
				_1547_v13 = Companion_D11_.Create_DC23_(_1546_v12)
				var _1548_v14 _dafny.Map
				_ = _1548_v14
				_1548_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1547_v13, _1547_v13, _1547_v13), _1542_v10)
				var _1549_v15 _dafny.Set
				_ = _1549_v15
				_1549_v15 = _dafny.SetOf(_1547_v13)
				_1542_v10 = (func() _dafny.Array {
					if (_1548_v14).Contains((_1549_v15).Union(_1549_v15)) {
						return (_1548_v14).Get((_1549_v15).Union(_1549_v15)).(_dafny.Array)
					}
					return _1542_v10
				})()
				_1527_v3 = _1527_v3
				_1542_v10 = _1542_v10
			} else if _source25.Is_DC19() {
				var _1550___mcc_h1 bool = _source25.Get_().(D9_DC19).Cf33
				_ = _1550___mcc_h1
				var _1551_cf33 bool = _1550___mcc_h1
				_ = _1551_cf33
				var _1552_v16 _dafny.CodePoint
				_ = _1552_v16
				_1552_v16 = _dafny.CodePoint('h')
				var _1553_v17 _dafny.Map
				_ = _1553_v17
				_1553_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v16, (_this).F27())
				var _1554_v19 _dafny.MultiSet
				_ = _1554_v19
				_1554_v19 = _dafny.MultiSetOf(_1553_v17, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v16, (_this).F27()), _1553_v17, func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter59 := _dafny.Iterate((_dafny.MultiSetOf(_1552_v16, _dafny.CodePoint('o'), _1552_v16)).Elements()); ; {
						_compr_52, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1555_v18 _dafny.CodePoint
						_1555_v18 = interface{}(_compr_52).(_dafny.CodePoint)
						if (_dafny.MultiSetOf(_1552_v16, _dafny.CodePoint('o'), _1552_v16)).Contains(_1555_v18) {
							_coll52.Add(_1555_v18, (func() _dafny.Int {
								if (_1527_v3).Contains((_this).F27()) {
									return (_1527_v3).Get((_this).F27()).(_dafny.Int)
								}
								return (_this).F27()
							})())
						}
					}
					return _coll52.ToMap()
				}())
				var _1556_v20 _dafny.Map
				_ = _1556_v20
				_1556_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, !(_1524_v0))
				var _1557_v21 _dafny.Sequence
				_ = _1557_v21
				_1557_v21 = _dafny.SeqOf(_1556_v20)
				var _1558_v22 _dafny.Set
				_ = _1558_v22
				_1558_v22 = _dafny.SetOf((_this).F27(), (_this).F27(), (_this).F27(), (_this).F27(), _dafny.IntOfUint32((Companion_Default___.Fm31(_1524_v0, _1551_cf33, globalState)).Cardinality()))
				var _1559_v23 _dafny.Map
				_ = _1559_v23
				_1559_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1525_v1).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1525_v1).Cardinality()))).Uint32()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v16, (_this).F27())).Cardinality())
				var _1560_v24 _dafny.Sequence
				_ = _1560_v24
				_1560_v24 = _dafny.SeqOf(_1552_v16)
				var _1561_v25 _dafny.Array
				_ = _1561_v25
				var _nwElement0_60 _dafny.Int = (func() _dafny.Int {
					if (_1554_v19).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v16, (_dafny.Zero).Minus((_this).F27()))) {
						return (_1554_v19).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v16, (_dafny.Zero).Minus((_this).F27())))
					}
					return (_this).F27()
				})()
				_ = _nwElement0_60
				var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(19))
				_ = _nw253
				(_nw253).ArraySet1(_nwElement0_60, 0)
				(_nw253).ArraySet1((_this).F27(), 1)
				(_nw253).ArraySet1((_this).F27(), 2)
				(_nw253).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1551_cf33, _1551_cf33)).Cardinality()), 3)
				(_nw253).ArraySet1((func() _dafny.Int {
					if _1524_v0 {
						return (_this).F27()
					}
					return (_dafny.Zero).Minus((_this).F27())
				})(), 4)
				(_nw253).ArraySet1(_dafny.IntOfInt64(19), 5)
				(_nw253).ArraySet1(((_1557_v21).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1557_v21).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 6)
				(_nw253).ArraySet1((Companion_D10_.Create_DC22_((_1556_v20).Cardinality(), (_this).F29())).Dtor_cf40(), 7)
				(_nw253).ArraySet1((_this).F27(), 8)
				(_nw253).ArraySet1((_this).F27(), 9)
				(_nw253).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(967), (_this).F27()), 10)
				(_nw253).ArraySet1((_this).F27(), 11)
				(_nw253).ArraySet1((_this).F27(), 12)
				(_nw253).ArraySet1((((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_dafny.IntOfInt64(-925)), 13)
				(_nw253).ArraySet1((_this).F27(), 14)
				(_nw253).ArraySet1((func() _dafny.Int {
					if _1551_cf33 {
						return (_this).F27()
					}
					return (_1558_v22).Cardinality()
				})(), 15)
				(_nw253).ArraySet1((_1559_v23).Cardinality(), 16)
				(_nw253).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1562_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1563_i1 _dafny.Int) _dafny.CodePoint {
						return _1562_v16
					}
				})(_1552_v16))), _1560_v24)).Cardinality())), 17)
				(_nw253).ArraySet1((_this).F27(), 18)
				_1561_v25 = _nw253
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1561_v25), 0))
				_ = _index195
				(_1561_v25).ArraySet1(((_this).F27()).Minus((_this).F27()), (_index195).Int())
				var _1564_v27 _dafny.Map
				_ = _1564_v27
				_1564_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F29()).IsProperSubsetOf((_this).F29()), func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(593), _dafny.IntOfInt64(-287))); ; {
						_compr_53, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _1565_v26 _dafny.Int
						_1565_v26 = interface{}(_compr_53).(_dafny.Int)
						if ((_dafny.IntOfInt64(593)).Cmp(_1565_v26) <= 0) && ((_1565_v26).Cmp(_dafny.IntOfInt64(-287)) < 0) {
							_coll53.Add(Companion_Default___.SafeDivisionInt(_1565_v26, (_this).F27()), _1524_v0)
						}
					}
					return _coll53.ToMap()
				}())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1561_v25), 0))
				_ = _index196
				var _rhs273 _dafny.Int = (_this).F27()
				_ = _rhs273
				var _rhs274 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1560_v24, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1560_v24).Cardinality()))).Uint32(), _1552_v16), _1560_v24)).Cardinality())
				_ = _rhs274
				var _rhs275 _dafny.Int = (_1564_v27).Cardinality()
				_ = _rhs275
				var _rhs276 bool = (func() bool {
					if (_1556_v20).Contains(_1524_v0) {
						return (_1556_v20).Get(_1524_v0).(bool)
					}
					return true
				})()
				_ = _rhs276
				var _lhs215 *GlobalState = globalState
				_ = _lhs215
				var _lhs216 _dafny.Array = _1561_v25
				_ = _lhs216
				var _lhs217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1561_v25), 0))
				_ = _lhs217
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				r0 = _rhs273
				_lhs215.F24 = _rhs274
				(_lhs216).ArraySet1(_rhs275, (_lhs217).Int())
				_lhs218.F26 = _rhs276
				var _1566_v28 *C1
				_ = _1566_v28
				var _nw254 *C1 = New_C1_()
				_ = _nw254
				_nw254.Ctor__(_1551_cf33, (_this).F29(), (_this).F30(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F27())), (_this).F28())
				_1566_v28 = _nw254
				var _1567_v29 _dafny.Sequence
				_ = _1567_v29
				_1567_v29 = _dafny.SeqOf(_1560_v24, _1560_v24)
				var _1568_v30 _dafny.Map
				_ = _1568_v30
				_1568_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1551_cf33, _dafny.Companion_Sequence_.Update(_1567_v29, (Companion_Default___.SafeIndex((_1561_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1561_v25), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1567_v29).Cardinality()))).Uint32(), _dafny.SeqOf(_1552_v16, _1552_v16, _dafny.CodePoint('u'))))
				(_1566_v28).F35 = _dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
					if (_1568_v30).Contains(_1524_v0) {
						return (_1568_v30).Get(_1524_v0).(_dafny.Sequence)
					}
					return _1567_v29
				})(), _dafny.Companion_Sequence_.Concatenate(_1567_v29, _1567_v29))
				var _1569_v31 _dafny.Array
				_ = _1569_v31
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw255
				_1569_v31 = _nw255
				var _1570_v32 _dafny.Map
				_ = _1570_v32
				_1570_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1569_v31, false)
				_1524_v0 = !(_1570_v32).Contains(_1569_v31)
			} else if _source25.Is_DC20() {
				var _1571___mcc_h2 _dafny.Int = _source25.Get_().(D9_DC20).Cf34
				_ = _1571___mcc_h2
				var _1572___mcc_h3 bool = _source25.Get_().(D9_DC20).Cf35
				_ = _1572___mcc_h3
				var _1573___mcc_h4 bool = _source25.Get_().(D9_DC20).Cf36
				_ = _1573___mcc_h4
				var _1574___mcc_h5 bool = _source25.Get_().(D9_DC20).Cf37
				_ = _1574___mcc_h5
				var _1575___mcc_h6 _dafny.Map = _source25.Get_().(D9_DC20).Cf38
				_ = _1575___mcc_h6
				var _1576_cf38 _dafny.Map = _1575___mcc_h6
				_ = _1576_cf38
				var _1577_cf37 bool = _1574___mcc_h5
				_ = _1577_cf37
				var _1578_cf36 bool = _1573___mcc_h4
				_ = _1578_cf36
				var _1579_cf35 bool = _1572___mcc_h3
				_ = _1579_cf35
				var _1580_cf34 _dafny.Int = _1571___mcc_h2
				_ = _1580_cf34
				var _1581_v33 _dafny.Map
				_ = _1581_v33
				_1581_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1578_cf36, _1524_v0)
				var _1582_v34 _dafny.Map
				_ = _1582_v34
				_1582_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1581_v33).Contains(_1524_v0) {
						return (_1581_v33).Get(_1524_v0).(bool)
					}
					return _1577_cf37
				})(), (_this).F27())
				var _1583_v35 _dafny.Array
				_ = _1583_v35
				var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw256
				_1583_v35 = _nw256
				var _1584_v36 _dafny.Map
				_ = _1584_v36
				_1584_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1582_v34, _1583_v35)
				var _1585_v37 _dafny.Array
				_ = _1585_v37
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(17))
				_ = _nw257
				_1585_v37 = _nw257
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1585_v37), 0))
				_ = _index197
				(_1585_v37).ArraySet1((_this).F29(), (_index197).Int())
				var _1586_v38 _dafny.Set
				_ = _1586_v38
				_1586_v38 = _dafny.SetOf(_1580_cf34, _1580_cf34, (_this).F27(), (_this).F27(), _1580_cf34)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1585_v37), 0))
				_ = _index198
				var _rhs277 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1525_v1, _1525_v1), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1525_v1, _1525_v1)).Cardinality()))).Uint32(), _1577_cf37)
				_ = _rhs277
				var _rhs278 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1525_v1, _1525_v1)).Cardinality())).Cmp(_1580_cf34) >= 0
				_ = _rhs278
				var _rhs279 _dafny.Map = ((_1584_v36).Update(_1582_v34, _1583_v35)).Merge(_1584_v36)
				_ = _rhs279
				var _rhs280 bool = (_1586_v38).IsSubsetOf(_1586_v38)
				_ = _rhs280
				var _rhs281 _dafny.MultiSet = (((_this).F29()).Difference(Companion_Default___.Fm32(globalState))).Union(_dafny.MultiSetOf(false))
				_ = _rhs281
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				var _lhs220 _dafny.Array = _1585_v37
				_ = _lhs220
				var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1585_v37), 0))
				_ = _lhs221
				_1525_v1 = _rhs277
				_lhs219.F26 = _rhs278
				_1584_v36 = _rhs279
				_1577_cf37 = _rhs280
				(_lhs220).ArraySet1(_rhs281, (_lhs221).Int())
				var _1587_v39 _dafny.Set
				_ = _1587_v39
				_1587_v39 = _dafny.SetOf(_1578_cf36, _1579_cf35, _1579_cf35)
				_1579_cf35 = !((_1587_v39).IsProperSubsetOf(_1587_v39))
				var _1588_v40 _dafny.CodePoint
				_ = _1588_v40
				_1588_v40 = _dafny.CodePoint('y')
				var _1589_v41 D4
				_ = _1589_v41
				_1589_v41 = Companion_D4_.Create_DC9_(_1580_cf34, _1579_cf35)
				var _1590_v42 _dafny.Set
				_ = _1590_v42
				_1590_v42 = _dafny.SetOf(_1589_v41, _1589_v41)
				var _1591_v43 D1
				_ = _1591_v43
				_1591_v43 = Companion_D1_.Create_DC2_(_1588_v40, (_1590_v42).Cardinality())
				(globalState).F18 = (func() _dafny.Int {
					if ((_1585_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1585_v37), 0))).Int()).(_dafny.MultiSet)).Contains(_1578_cf36) {
						return ((_1585_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1585_v37), 0))).Int()).(_dafny.MultiSet)).Multiplicity(_1578_cf36)
					}
					return (_1591_v43).Dtor_cf3()
				})()
				(globalState).F18 = (_1580_cf34).Minus((_this).F27())
			} else {
				var _1592___mcc_h7 _dafny.Set = _source25.Get_().(D9_DC17).Cf31
				_ = _1592___mcc_h7
				var _1593_cf31 _dafny.Set = _1592___mcc_h7
				_ = _1593_cf31
				var _1594_v44 *C0
				_ = _1594_v44
				var _nw258 *C0 = New_C0_()
				_ = _nw258
				_nw258.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F30()).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("tidvowswu"))
				_1594_v44 = _nw258
				var _1595_v45 D6
				_ = _1595_v45
				_1595_v45 = Companion_D6_.Create_DC11_(_1594_v44)
				var _pat_let_tv32 = _1594_v44
				_ = _pat_let_tv32
				var _1596_v46 _dafny.Array
				_ = _1596_v46
				var _nwElement0_61 D6 = _1595_v45
				_ = _nwElement0_61
				var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(4))
				_ = _nw259
				(_nw259).ArraySet1(_nwElement0_61, 0)
				(_nw259).ArraySet1(_1595_v45, 1)
				(_nw259).ArraySet1(Companion_D6_.Create_DC11_(_1594_v44), 2)
				(_nw259).ArraySet1(func(_pat_let38_0 D6) D6 {
					return func(_1597_dt__update__tmp_h2 D6) D6 {
						return func(_pat_let39_0 *C0) D6 {
							return func(_1598_dt__update_hcf19_h0 *C0) D6 {
								return Companion_D6_.Create_DC11_(_1598_dt__update_hcf19_h0)
							}(_pat_let39_0)
						}(_pat_let_tv32)
					}(_pat_let38_0)
				}(Companion_D6_.Create_DC11_(_1594_v44)), 3)
				_1596_v46 = _nw259
				var _pat_let_tv33 = _1594_v44
				_ = _pat_let_tv33
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1596_v46), 0))
				_ = _index199
				(_1596_v46).ArraySet1(func(_pat_let40_0 D6) D6 {
					return func(_1599_dt__update__tmp_h3 D6) D6 {
						return func(_pat_let41_0 *C0) D6 {
							return func(_1600_dt__update_hcf19_h1 *C0) D6 {
								return Companion_D6_.Create_DC11_(_1600_dt__update_hcf19_h1)
							}(_pat_let41_0)
						}(_pat_let_tv33)
					}(_pat_let40_0)
				}(_1595_v45), (_index199).Int())
				var _pat_let_tv34 = _1594_v44
				_ = _pat_let_tv34
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1596_v46), 0))
				_ = _index200
				(_1596_v46).ArraySet1(func(_pat_let42_0 D6) D6 {
					return func(_1601_dt__update__tmp_h4 D6) D6 {
						return func(_pat_let43_0 *C0) D6 {
							return func(_1602_dt__update_hcf19_h2 *C0) D6 {
								return Companion_D6_.Create_DC11_(_1602_dt__update_hcf19_h2)
							}(_pat_let43_0)
						}(_pat_let_tv34)
					}(_pat_let42_0)
				}(_1595_v45), (_index200).Int())
				(globalState).F26 = !(_1524_v0)
				(globalState).F0 = ((_this).F27()).Times((_this).F27())
				var _1603_v47 *C0
				_ = _1603_v47
				var _nw260 *C0 = New_C0_()
				_ = _nw260
				_nw260.Ctor__((_1594_v44).F36(), (func() _dafny.Sequence {
					if !(_1524_v0) {
						return (_1594_v44).F37()
					}
					return (_1594_v44).F37()
				})())
				_1603_v47 = _nw260
			}
			if _1524_v0 {
				var _1604_v48 _dafny.Sequence
				_ = _1604_v48
				_1604_v48 = _dafny.UnicodeSeqOfUtf8Bytes("aljehft")
				var _1605_v49 *C0
				_ = _1605_v49
				var _nw261 *C0 = New_C0_()
				_ = _nw261
				_nw261.Ctor__(_dafny.IntOfUint32((_1604_v48).Cardinality()), _1604_v48)
				_1605_v49 = _nw261
				var _1606_v50 D6
				_ = _1606_v50
				_1606_v50 = Companion_D6_.Create_DC11_(_1605_v49)
				var _pat_let_tv35 = _1605_v49
				_ = _pat_let_tv35
				_1606_v50 = func(_pat_let44_0 D6) D6 {
					return func(_1607_dt__update__tmp_h5 D6) D6 {
						return func(_pat_let45_0 *C0) D6 {
							return func(_1608_dt__update_hcf19_h3 *C0) D6 {
								return Companion_D6_.Create_DC11_(_1608_dt__update_hcf19_h3)
							}(_pat_let45_0)
						}(_pat_let_tv35)
					}(_pat_let44_0)
				}(_1606_v50)
				(globalState).F26 = !(_1524_v0)
				var _1609_v51 _dafny.Array
				_ = _1609_v51
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_29
				var _nw262 _dafny.Array
				_ = _nw262
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw262 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = (func(_1610_v49 *C0) func(_dafny.Int) _dafny.Int {
						return func(_1611_i2 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1611_i2, (_1610_v49).F36())
						}
					})(_1605_v49)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw262 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw262).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw262).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1609_v51 = _nw262
				var _1612_v52 _dafny.Map
				_ = _1612_v52
				_1612_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1528_v4).Dtor_cf34(), _1609_v51)
				var _1613_v53 _dafny.Sequence
				_ = _1613_v53
				_1613_v53 = _dafny.SeqOf(_1609_v51, _1609_v51, _1609_v51, _1609_v51, _1609_v51)
				_1612_v52 = (_1612_v52).Update((_1605_v49).F36(), (_1613_v53).Select((Companion_Default___.SafeIndex((_1605_v49).F36(), _dafny.IntOfUint32((_1613_v53).Cardinality()))).Uint32()).(_dafny.Array))
				var _1614_v54 _dafny.CodePoint
				_ = _1614_v54
				_1614_v54 = _dafny.CodePoint('r')
				var _1615_v55 D4
				_ = _1615_v55
				_1615_v55 = Companion_D4_.Create_DC8_(_1614_v54, (_this).F27(), false)
				var _1616_v56 _dafny.Map
				_ = _1616_v56
				_1616_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, (_this).F27())
				var _1617_v57 _dafny.Map
				_ = _1617_v57
				_1617_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_1615_v55).Dtor_cf15(), _1524_v0)).IsDisjointFrom((_this).F29()), (_1616_v56).Cardinality())
				_1616_v56 = (_1616_v56).Update(_1524_v0, (_dafny.Zero).Minus((_this).F27()))
				(globalState).F26 = (_this).Fm3((_this).F29(), globalState)
			} else {
				var _1618_v58 _dafny.Array
				_ = _1618_v58
				var _nw263 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw263
				_1618_v58 = _nw263
				var _1619_v59 _dafny.Sequence
				_ = _1619_v59
				_1619_v59 = _dafny.SeqOf(_1618_v58)
				var _1620_v60 *C0
				_ = _1620_v60
				var _nw264 *C0 = New_C0_()
				_ = _nw264
				_nw264.Ctor__(_dafny.IntOfUint32((_1619_v59).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("mwnfm"))
				_1620_v60 = _nw264
				var _1621_v61 _dafny.CodePoint
				_ = _1621_v61
				_1621_v61 = _dafny.CodePoint('k')
				_1621_v61 = _1621_v61
				_1524_v0 = _1524_v0
				var _1622_v62 _dafny.Set
				_ = _1622_v62
				_1622_v62 = _dafny.SetOf(_1524_v0)
				var _1623_v63 _dafny.Set
				_ = _1623_v63
				_1623_v63 = _dafny.SetOf(_1622_v62)
				(globalState).F26 = !(!(((_1623_v63).Cardinality()).Cmp((_1620_v60).F36()) > 0))
				(globalState).F21 = (_1620_v60).F36()
			}
			var _1624_v64 D10
			_ = _1624_v64
			_1624_v64 = Companion_D10_.Create_DC22_(((_this).F27()).Minus(((_this).F29()).Cardinality()), ((_this).F29()).Union((_this).F29()))
			_1624_v64 = _1624_v64
			var _1625_v65 _dafny.Map
			_ = _1625_v65
			_1625_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1524_v0)
			_1625_v65 = ((_1625_v65).Merge(_1625_v65)).Merge((_1625_v65).Merge(_1625_v65))
			var _1626_v66 _dafny.Sequence
			_ = _1626_v66
			_1626_v66 = _dafny.UnicodeSeqOfUtf8Bytes("lxncklgcj")
			var _1627_v67 D6
			_ = _1627_v67
			_1627_v67 = Companion_D6_.Create_DC12_(true, _dafny.SeqOf(_1525_v1, _1525_v1, _1525_v1), _1626_v66)
			var _1628_v68 _dafny.Map
			_ = _1628_v68
			_1628_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1627_v67, true)
			var _1629_v69 _dafny.Array
			_ = _1629_v69
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_30
			var _nw265 _dafny.Array
			_ = _nw265
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw265 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) bool = func(_1630_i3 _dafny.Int) bool {
					return false
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw265 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw265).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw265).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1629_v69 = _nw265
			var _1631_v70 _dafny.Map
			_ = _1631_v70
			_1631_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1524_v0)
			var _rhs282 _dafny.Int = (_this).F27()
			_ = _rhs282
			var _rhs283 _dafny.Map = _1628_v68
			_ = _rhs283
			var _rhs284 _dafny.Array = _1629_v69
			_ = _rhs284
			var _rhs285 bool = (!(_1631_v70).Contains(_1524_v0)) || (_1524_v0)
			_ = _rhs285
			var _rhs286 bool = _1524_v0
			_ = _rhs286
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			var _lhs223 *GlobalState = globalState
			_ = _lhs223
			r2 = _rhs282
			_1628_v68 = _rhs283
			_1629_v69 = _rhs284
			_lhs222.F26 = _rhs285
			_lhs223.F26 = _rhs286
		} else {
			(globalState).F26 = _1524_v0
			_1525_v1 = _1525_v1
			var _1632_v71 _dafny.Map
			_ = _1632_v71
			_1632_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F27(), globalState), _1524_v0)
			var _1633_v72 _dafny.Sequence
			_ = _1633_v72
			_1633_v72 = _dafny.UnicodeSeqOfUtf8Bytes("ukkh")
			var _1634_v73 *C0
			_ = _1634_v73
			var _nw266 *C0 = New_C0_()
			_ = _nw266
			_nw266.Ctor__((_1632_v71).Cardinality(), _1633_v72)
			_1634_v73 = _nw266
			var _1635_v74 _dafny.Map
			_ = _1635_v74
			_1635_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1634_v73).F36(), _1634_v73)
			var _1636_v75 _dafny.Array
			_ = _1636_v75
			var _nwElement0_62 *C0 = _1634_v73
			_ = _nwElement0_62
			var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(22))
			_ = _nw267
			(_nw267).ArraySet1(_nwElement0_62, 0)
			(_nw267).ArraySet1(_1634_v73, 1)
			(_nw267).ArraySet1(_1634_v73, 2)
			(_nw267).ArraySet1(_1634_v73, 3)
			(_nw267).ArraySet1((func() *C0 {
				if (_1635_v74).Contains((_dafny.Zero).Minus((_this).F27())) {
					return (_1635_v74).Get((_dafny.Zero).Minus((_this).F27())).(*C0)
				}
				return _1634_v73
			})(), 4)
			(_nw267).ArraySet1(_1634_v73, 5)
			(_nw267).ArraySet1(_1634_v73, 6)
			(_nw267).ArraySet1(_1634_v73, 7)
			(_nw267).ArraySet1(_1634_v73, 8)
			(_nw267).ArraySet1(_1634_v73, 9)
			(_nw267).ArraySet1(_1634_v73, 10)
			(_nw267).ArraySet1(_1634_v73, 11)
			(_nw267).ArraySet1(_1634_v73, 12)
			(_nw267).ArraySet1(_1634_v73, 13)
			(_nw267).ArraySet1(_1634_v73, 14)
			(_nw267).ArraySet1(_1634_v73, 15)
			(_nw267).ArraySet1(_1634_v73, 16)
			(_nw267).ArraySet1(_1634_v73, 17)
			(_nw267).ArraySet1(_1634_v73, 18)
			(_nw267).ArraySet1(_1634_v73, 19)
			(_nw267).ArraySet1(_1634_v73, 20)
			(_nw267).ArraySet1(_1634_v73, 21)
			_1636_v75 = _nw267
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1636_v75), 0))
			_ = _index201
			(_1636_v75).ArraySet1(_1634_v73, (_index201).Int())
			var _1637_v76 D2
			_ = _1637_v76
			_1637_v76 = Companion_D2_.Create_DC4_((_this).F27(), (_dafny.SetOf(_1524_v0)).Cardinality(), (_this).F27(), _dafny.CodePoint('s'), _dafny.UnicodeSeqOfUtf8Bytes("iga"))
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1636_v75), 0))
			_ = _index202
			var _rhs287 _dafny.Int = ((_this).F27()).Plus((_dafny.Zero).Minus((_1634_v73).F36()))
			_ = _rhs287
			var _rhs288 *C0 = _1634_v73
			_ = _rhs288
			var _rhs289 _dafny.Int = _dafny.IntOfUint32(((_1637_v76).Dtor_cf9()).Cardinality())
			_ = _rhs289
			var _lhs224 *GlobalState = globalState
			_ = _lhs224
			var _lhs225 _dafny.Array = _1636_v75
			_ = _lhs225
			var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1636_v75), 0))
			_ = _lhs226
			var _lhs227 *GlobalState = globalState
			_ = _lhs227
			_lhs224.F15 = _rhs287
			(_lhs225).ArraySet1(_rhs288, (_lhs226).Int())
			_lhs227.F3 = _rhs289
			var _1638_v77 _dafny.CodePoint
			_ = _1638_v77
			_1638_v77 = _dafny.CodePoint('e')
			_1638_v77 = _1638_v77
			(globalState).F21 = (_this).F27()
		}
		var _1639_v78 _dafny.Sequence
		_ = _1639_v78
		_1639_v78 = _dafny.UnicodeSeqOfUtf8Bytes("odaravso")
		var _1640_v79 _dafny.Map
		_ = _1640_v79
		_1640_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F27()), _1524_v0)
		var _1641_v80 *C2
		_ = _1641_v80
		var _nw268 *C2 = New_C2_()
		_ = _nw268
		_nw268.Ctor__((_this).F29(), (_this).F27())
		_1641_v80 = _nw268
		var _1642_v81 D12
		_ = _1642_v81
		_1642_v81 = Companion_D12_.Create_DC26_(_1641_v80)
		var _1643_v82 _dafny.Set
		_ = _1643_v82
		_1643_v82 = _dafny.SetOf(_1641_v80, _1641_v80, _1641_v80, (_1642_v81).Dtor_cf46(), _1641_v80)
		var _1644_v83 _dafny.MultiSet
		_ = _1644_v83
		_1644_v83 = _dafny.MultiSetOf((_this).F27(), Companion_Default___.Fm1(globalState))
		var _1645_v84 _dafny.Array
		_ = _1645_v84
		var _nwElement0_63 bool = (_1525_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yspuqcihq")).Cardinality()), _dafny.IntOfUint32((_1525_v1).Cardinality()))).Uint32()).(bool)
		_ = _nwElement0_63
		var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(13))
		_ = _nw269
		(_nw269).ArraySet1(_nwElement0_63, 0)
		(_nw269).ArraySet1(_1524_v0, 1)
		(_nw269).ArraySet1(_1524_v0, 2)
		(_nw269).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1639_v78, _1639_v78), 3)
		(_nw269).ArraySet1(_1524_v0, 4)
		(_nw269).ArraySet1(_1524_v0, 5)
		(_nw269).ArraySet1(_1524_v0, 6)
		(_nw269).ArraySet1(false, 7)
		(_nw269).ArraySet1(((_this).F27()).Cmp((_this).F27()) < 0, 8)
		(_nw269).ArraySet1(((_this).F27()).Cmp((_this).F27()) < 0, 9)
		(_nw269).ArraySet1(_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm33(_1640_v79, (_1643_v82).Cardinality(), _dafny.IntOfUint32((_1525_v1).Cardinality()), (_1644_v83).Cardinality(), globalState), _1525_v1), 10)
		(_nw269).ArraySet1(_1524_v0, 11)
		(_nw269).ArraySet1(true, 12)
		_1645_v84 = _nw269
		for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1645_v84), 0))); ; {
			_guard_loop_6, _ok61 := _iter61()
			if !_ok61 {
				break
			}
			var _1646_i4 _dafny.Int
			_1646_i4 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1646_i4).Sign() != -1) && ((_1646_i4).Cmp(_dafny.ArrayLen((_1645_v84), 0)) < 0)) {
				(_1645_v84).ArraySet1(!((_dafny.IntOfInt64(175)).Cmp(_dafny.IntOfInt64(871)) > 0), (_1646_i4).Int())
			}
		}
		var _1647_v85 _dafny.Map
		_ = _1647_v85
		_1647_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1645_v84)
		var _1648_v86 _dafny.Map
		_ = _1648_v86
		_1648_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, (_this).F27())
		_1647_v85 = (_1647_v85).Update((func() _dafny.Int {
			if (_1648_v86).Contains(Companion_Default___.Fm0((_this).F27(), globalState)) {
				return (_1648_v86).Get(Companion_Default___.Fm0((_this).F27(), globalState)).(_dafny.Int)
			}
			return (_1640_v79).Cardinality()
		})(), _1645_v84)
		var _1649_v87 _dafny.CodePoint
		_ = _1649_v87
		_1649_v87 = _dafny.CodePoint('k')
		var _1650_v88 _dafny.Map
		_ = _1650_v88
		_1650_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, _1649_v87)
		_1650_v88 = (_1650_v88).Update(_1524_v0, _1649_v87)
		var _hi6 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _hi6
		for _1651_i5 := (_this).F27(); _1651_i5.Cmp(_hi6) < 0; _1651_i5 = _1651_i5.Plus(_dafny.One) {
			(globalState).F18 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-374))
			(globalState).F26 = (_dafny.IntOfInt64(420)).Cmp(_dafny.IntOfInt64(871)) != 0
			if (func() bool {
				if (_this).Fm3((_this).F29(), globalState) {
					return !((_dafny.MultiSetOf(_1524_v0)).IsProperSubsetOf((_dafny.MultiSetOf(_1524_v0, _1524_v0)).Update(false, Companion_Default___.Abs(_1651_i5))))
				}
				return _1524_v0
			})() {
				_1524_v0 = !(true)
				var _rhs290 bool = _1524_v0
				_ = _rhs290
				var _rhs291 _dafny.Sequence = _1639_v78
				_ = _rhs291
				_1524_v0 = _rhs290
				_1639_v78 = _rhs291
				var _1652_v89 _dafny.Array
				_ = _1652_v89
				var _nw270 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw270
				_1652_v89 = _nw270
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_1652_v89), 0))
				_ = _index203
				(_1652_v89).ArraySet1(_1651_i5, (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))
				_ = _index204
				(_1645_v84).ArraySet1(_1524_v0, (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_1652_v89), 0))
				_ = _index205
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))
				_ = _index206
				var _rhs292 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F27(), _dafny.IntOfInt64(316))
				_ = _rhs292
				var _rhs293 bool = (_1651_i5).Cmp((_this).F27()) == 0
				_ = _rhs293
				var _rhs294 _dafny.Array = _1645_v84
				_ = _rhs294
				var _rhs295 _dafny.Int = Companion_Default___.SafeModuloInt(_1651_i5, _dafny.IntOfUint32((_dafny.SeqOf((_this).F27(), (_this).F27())).Cardinality()))
				_ = _rhs295
				var _rhs296 bool = (_dafny.MultiSetOf((_this).F27(), (_this).F27(), _1651_i5)).IsProperSubsetOf(_dafny.MultiSetOf((_this).F27(), _1651_i5))
				_ = _rhs296
				var _lhs228 _dafny.Array = _1652_v89
				_ = _lhs228
				var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_1652_v89), 0))
				_ = _lhs229
				var _lhs230 _dafny.Array = _1645_v84
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))
				_ = _lhs231
				(_lhs228).ArraySet1(_rhs292, (_lhs229).Int())
				_1524_v0 = _rhs293
				_1645_v84 = _rhs294
				r2 = _rhs295
				(_lhs230).ArraySet1(_rhs296, (_lhs231).Int())
				var _1653_v90 _dafny.MultiSet
				_ = _1653_v90
				_1653_v90 = _1644_v83
				var _1654_v91 _dafny.Map
				_ = _1654_v91
				_1654_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1653_v90, (_1645_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))).Int()).(bool))
				var _1655_v92 _dafny.Sequence
				_ = _1655_v92
				_1655_v92 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1653_v90, (_1645_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))).Int()).(bool)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1653_v90, (_1645_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))).Int()).(bool)))
				var _1656_v93 _dafny.Set
				_ = _1656_v93
				_1656_v93 = _dafny.SetOf((_1645_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1645_v84), 0))).Int()).(bool))
				_1654_v91 = (_1655_v92).Select((Companion_Default___.SafeIndex(((_1656_v93).Cardinality()).Plus(_1651_i5), _dafny.IntOfUint32((_1655_v92).Cardinality()))).Uint32()).(_dafny.Map)
				(globalState).F3 = (_this).F27()
			} else {
				(globalState).F0 = _1651_i5
				var _1657_v94 _dafny.Array
				_ = _1657_v94
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_31
				var _nw271 _dafny.Array
				_ = _nw271
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw271 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Int = (func(_1658_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1659_i6 _dafny.Int) _dafny.Int {
							return (_1659_i6).Minus(_dafny.IntOfUint32((_1658_v1).Cardinality()))
						}
					})(_1525_v1)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw271 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw271).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw271).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1657_v94 = _nw271
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1657_v94), 0))
				_ = _index207
				(_1657_v94).ArraySet1(_dafny.IntOfInt64(311), (_index207).Int())
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1657_v94), 0))
				_ = _index208
				(_1657_v94).ArraySet1(_1651_i5, (_index208).Int())
				var _1660_v95 _dafny.Sequence
				_ = _1660_v95
				_1660_v95 = _dafny.UnicodeSeqOfUtf8Bytes("hfpe")
				var _1661_v96 *C0
				_ = _1661_v96
				var _nw272 *C0 = New_C0_()
				_ = _nw272
				_nw272.Ctor__(((_this).F29()).Cardinality(), (_1660_v95))
				_1661_v96 = _nw272
				var _1662_v97 _dafny.Map
				_ = _1662_v97
				_1662_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, _1661_v96)
				var _1663_v98 *C2
				_ = _1663_v98
				var _nw273 *C2 = New_C2_()
				_ = _nw273
				_nw273.Ctor__((_this).F29(), (_1662_v97).Cardinality())
				_1663_v98 = _nw273
				_1639_v78 = (_1661_v96).F37()
				var _1664_v99 _dafny.Map
				_ = _1664_v99
				_1664_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1657_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1657_v94), 0))).Int()).(_dafny.Int), (_1657_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1657_v94), 0))).Int()).(_dafny.Int))
				var _1665_v100 _dafny.Set
				_ = _1665_v100
				_1665_v100 = _dafny.SetOf((_1641_v80).Fm21(globalState), _1524_v0)
				var _1666_v101 T0
				_ = _1666_v101
				var _nw274 *C3 = New_C3_()
				_ = _nw274
				_nw274.Ctor__((_1665_v100).Cardinality(), (_this).F28())
				_1666_v101 = _nw274
				var _1667_v102 _dafny.Set
				_ = _1667_v102
				_1667_v102 = _dafny.SetOf(_1666_v101, _1666_v101)
				var _1668_v103 _dafny.Map
				_ = _1668_v103
				_1668_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1651_i5, _1648_v86)
				_1664_v99 = (_1664_v99).Update(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1651_i5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v0, (_1667_v102).Cardinality()))).Merge(_1668_v103)).Cardinality(), _1651_i5)
			}
			var _1669_v104 D4
			_ = _1669_v104
			_1669_v104 = Companion_D4_.Create_DC7_(_1645_v84)
			_1669_v104 = _1669_v104
		}
		var _hi7 _dafny.Int = (_this).F27()
		_ = _hi7
		for _1670_i7 := (_dafny.Zero).Minus((_this).F27()); _1670_i7.Cmp(_hi7) < 0; _1670_i7 = _1670_i7.Plus(_dafny.One) {
			var _1671_v105 _dafny.Array
			_ = _1671_v105
			var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw275
			_1671_v105 = _nw275
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1671_v105), 0))
			_ = _index209
			(_1671_v105).ArraySet1((_1640_v79).Merge(_1640_v79), (_index209).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1671_v105), 0))
			_ = _index210
			(_1671_v105).ArraySet1(_1640_v79, (_index210).Int())
			var _1672_v106 _dafny.Array
			_ = _1672_v106
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_32
			var _nw276 _dafny.Array
			_ = _nw276
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw276 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1673_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1674_i8 _dafny.Int) _dafny.Sequence {
						return (func() _dafny.Sequence {
							if _1673_v0 {
								return _dafny.SeqOf((_this).F30())
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-529))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg99 _dafny.Int) interface{} {
									return coer99(arg99)
								}
							}(func(_1675_i9 _dafny.Int) _dafny.Sequence {
								return (_this).F30()
							}))
						})()
					}
				})(_1524_v0)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw276 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw276).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw276).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1672_v106 = _nw276
			var _1676_v107 _dafny.Sequence
			_ = _1676_v107
			_1676_v107 = _dafny.SeqOf((_this).F30(), (_this).F30())
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1672_v106), 0))
			_ = _index211
			(_1672_v106).ArraySet1(_1676_v107, (_index211).Int())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1672_v106), 0))
			_ = _index212
			(_1672_v106).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1676_v107, _1676_v107), (_index212).Int())
			var _1677_v108 _dafny.Sequence
			_ = _1677_v108
			_1677_v108 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1670_i7, _1524_v0))
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1671_v105), 0))
			_ = _index213
			(_1671_v105).ArraySet1((_1677_v108).Select((Companion_Default___.SafeIndex(_1670_i7, _dafny.IntOfUint32((_1677_v108).Cardinality()))).Uint32()).(_dafny.Map), (_index213).Int())
			r0 = (_this).F27()
		}
		var _1678_v109 _dafny.Map
		_ = _1678_v109
		_1678_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).Plus((_this).F27()), _dafny.IntOfInt64(53))
		r0 = (func() _dafny.Int {
			if (_1678_v109).Contains((_this).F27()) {
				return (_1678_v109).Get((_this).F27()).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_this).F27())
		})()
		var _1679_v110 D9
		_ = _1679_v110
		_1679_v110 = Companion_D9_.Create_DC20_((_this).F27(), true, _1524_v0, false, _1678_v109)
		var _1680_v111 _dafny.Array
		_ = _1680_v111
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_33
		var _nw277 _dafny.Array
		_ = _nw277
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw277 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.Int = func(_1681_i10 _dafny.Int) _dafny.Int {
				return (_1681_i10).Times((_dafny.Zero).Minus((_this).F27()))
			}
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw277 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw277).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw277).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1680_v111 = _nw277
		var _1682_v112 D12
		_ = _1682_v112
		_1682_v112 = Companion_D12_.Create_DC27_(_1524_v0, _1644_v83, _dafny.IntOfUint32((_1639_v78).Cardinality()), _1680_v111)
		var _pat_let_tv36 = _1524_v0
		_ = _pat_let_tv36
		var _1683_v113 _dafny.Sequence
		_ = _1683_v113
		_1683_v113 = _dafny.SeqOf(func(_pat_let46_0 D12) D12 {
			return func(_1684_dt__update__tmp_h6 D12) D12 {
				return func(_pat_let47_0 bool) D12 {
					return func(_1685_dt__update_hcf47_h0 bool) D12 {
						return Companion_D12_.Create_DC27_(_1685_dt__update_hcf47_h0, (_1684_dt__update__tmp_h6).Dtor_cf48(), (_1684_dt__update__tmp_h6).Dtor_cf49(), (_1684_dt__update__tmp_h6).Dtor_cf50())
					}(_pat_let47_0)
				}(_pat_let_tv36)
			}(_pat_let46_0)
		}(Companion_D12_.Create_DC27_(_1524_v0, _1644_v83, _dafny.IntOfInt64(468), _1680_v111)), _1682_v112)
		r1 = Companion_Default___.SafeDivisionInt(((_this).F30()).Select((Companion_Default___.SafeIndex((_1679_v110).Dtor_cf34(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.SafeDivisionInt((_this).F27(), ((_dafny.MultiSetFromSeq(_1683_v113)).Update(_1682_v112, Companion_Default___.Abs((_this).F27()))).Cardinality()))
		r2 = (func() _dafny.Int {
			if _1524_v0 {
				return (_this).F27()
			}
			return _dafny.IntOfUint32(((_this).F30()).Cardinality())
		})()
		return r0, r1, r2
	}
}
func (_this *C5) M9(globalState *GlobalState) (_dafny.Array, bool, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _1686_v0 bool
		_ = _1686_v0
		_1686_v0 = false
		if (((_this).F27()).Plus(((_this).Fm2(_1686_v0, _1686_v0, _1686_v0, globalState)).Cardinality())).Cmp((_this).F27()) >= 0 {
			var _1687_v1 _dafny.Sequence
			_ = _1687_v1
			_1687_v1 = _dafny.UnicodeSeqOfUtf8Bytes("a")
			var _1688_v2 bool
			_ = _1688_v2
			var _1689_v3 _dafny.Int
			_ = _1689_v3
			var _out44 bool
			_ = _out44
			var _out45 _dafny.Int
			_ = _out45
			_out44, _out45 = Companion_Default___.M0(true, _dafny.IntOfUint32((_1687_v1).Cardinality()), ((_this).F27()).Times((_this).F27()), (_this).F27(), globalState)
			_1688_v2 = _out44
			_1689_v3 = _out45
			var _1690_v4 *C2
			_ = _1690_v4
			var _nw278 *C2 = New_C2_()
			_ = _nw278
			_nw278.Ctor__((_this).F29(), (_this).F27())
			_1690_v4 = _nw278
			var _1691_v5 D12
			_ = _1691_v5
			_1691_v5 = Companion_D12_.Create_DC29_(Companion_D12_.Create_DC26_(_1690_v4))
			var _1692_v6 _dafny.Map
			_ = _1692_v6
			_1692_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v5, _1688_v2)
			var _1693_v7 _dafny.Map
			_ = _1693_v7
			_1693_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1686_v0)
			_1692_v6 = (_1692_v6).Update(_1691_v5, (func() bool {
				if (_1693_v7).Contains(_1688_v2) {
					return (_1693_v7).Get(_1688_v2).(bool)
				}
				return _1686_v0
			})())
			var _1694_v8 _dafny.Array
			_ = _1694_v8
			var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw279
			_1694_v8 = _nw279
			var _1695_v9 _dafny.Sequence
			_ = _1695_v9
			_1695_v9 = _dafny.SeqOf(_1694_v8)
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_1695_v9, _1695_v9) {
				var _1696_v10 _dafny.Array
				_ = _1696_v10
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_34
				var _nw280 _dafny.Array
				_ = _nw280
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw280 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1697_v0 bool, _1698_v2 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_1699_i0 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_1697_v0, _1697_v0, _1698_v2)
						}
					})(_1686_v0, _1688_v2)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw280 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw280).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw280).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1696_v10 = _nw280
				var _1700_v11 _dafny.Sequence
				_ = _1700_v11
				_1700_v11 = _dafny.SeqOf(!((_1690_v4).Fm21(globalState)))
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1696_v10), 0))
				_ = _index214
				(_1696_v10).ArraySet1(_1700_v11, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1696_v10), 0))
				_ = _index215
				(_1696_v10).ArraySet1(_1700_v11, (_index215).Int())
				var _1701_v12 *C3
				_ = _1701_v12
				var _nw281 *C3 = New_C3_()
				_ = _nw281
				_nw281.Ctor__((_this).F27(), (_this).F28())
				_1701_v12 = _nw281
				var _1702_v13 _dafny.Map
				_ = _1702_v13
				_1702_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if false {
						return !(_1686_v0)
					}
					return _1686_v0
				})(), Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27()))
				_1702_v13 = _1702_v13
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_1694_v8), 0))
				_ = _index216
				(_1694_v8).ArraySet1(_dafny.IntOfUint32((_1700_v11).Cardinality()), (_index216).Int())
				var _1703_v14 _dafny.Array
				_ = _1703_v14
				var _nwElement0_64 bool = _1686_v0
				_ = _nwElement0_64
				var _nw282 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.One)
				_ = _nw282
				(_nw282).ArraySet1(_nwElement0_64, 0)
				_1703_v14 = _nw282
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_1694_v8), 0))
				_ = _index217
				var _rhs297 bool = _1686_v0
				_ = _rhs297
				var _rhs298 _dafny.Array = _1703_v14
				_ = _rhs298
				var _rhs299 _dafny.Int = _1689_v3
				_ = _rhs299
				var _rhs300 bool = ((_dafny.IntOfInt64(520)).Cmp(_dafny.IntOfUint32((_1687_v1).Cardinality())) >= 0) || (_dafny.Companion_Sequence_.IsPrefixOf(_1700_v11, (_1696_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1696_v10), 0))).Int()).(_dafny.Sequence)))
				_ = _rhs300
				var _lhs232 *GlobalState = globalState
				_ = _lhs232
				var _lhs233 _dafny.Array = _1694_v8
				_ = _lhs233
				var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_1694_v8), 0))
				_ = _lhs234
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				_lhs232.F26 = _rhs297
				r0 = _rhs298
				(_lhs233).ArraySet1(_rhs299, (_lhs234).Int())
				_lhs235.F26 = _rhs300
				var _1704_v15 _dafny.Set
				_ = _1704_v15
				_1704_v15 = _dafny.SetOf(_1703_v14)
				var _1705_v16 D12
				_ = _1705_v16
				_1705_v16 = Companion_D12_.Create_DC28_(_dafny.IntOfInt64(507), (_1704_v15).Cardinality(), _1688_v2)
				var _1706_v17 _dafny.Map
				_ = _1706_v17
				_1706_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v16, _1687_v1)
				_1706_v17 = (_1706_v17).Update(_1705_v16, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fcnwjnotv"), _1687_v1))
			} else {
				var _1707_v18 _dafny.Array
				_ = _1707_v18
				var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
				_ = _nw283
				_1707_v18 = _nw283
				_1707_v18 = _1707_v18
				(globalState).F24 = (_this).F27()
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw284
				r0 = _nw284
				r2 = _1686_v0
				r1 = _1688_v2
			}
			_1693_v7 = (_1693_v7).Update(_1686_v0, _1688_v2)
			var _1708_v19 _dafny.Sequence
			_ = _1708_v19
			_1708_v19 = _dafny.SeqOf(_1688_v2, true, false, _1688_v2)
			if (_1708_v19).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1708_v19).Cardinality()))).Uint32()).(bool) {
				var _1709_v20 _dafny.CodePoint
				_ = _1709_v20
				_1709_v20 = _dafny.CodePoint('x')
				var _1710_v21 _dafny.Array
				_ = _1710_v21
				var _nwElement0_65 _dafny.CodePoint = _1709_v20
				_ = _nwElement0_65
				var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(13))
				_ = _nw285
				(_nw285).ArraySet1CodePoint(_nwElement0_65, 0)
				(_nw285).ArraySet1CodePoint(_1709_v20, 1)
				(_nw285).ArraySet1CodePoint(_1709_v20, 2)
				(_nw285).ArraySet1CodePoint(_1709_v20, 3)
				(_nw285).ArraySet1CodePoint((_1687_v1).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1687_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), 4)
				(_nw285).ArraySet1CodePoint(_1709_v20, 5)
				(_nw285).ArraySet1CodePoint(_1709_v20, 6)
				(_nw285).ArraySet1CodePoint(_1709_v20, 7)
				(_nw285).ArraySet1CodePoint(_1709_v20, 8)
				(_nw285).ArraySet1CodePoint(Companion_Default___.Fm38(_1693_v7, globalState), 9)
				(_nw285).ArraySet1CodePoint(_1709_v20, 10)
				(_nw285).ArraySet1CodePoint(_1709_v20, 11)
				(_nw285).ArraySet1CodePoint((Companion_D1_.Create_DC2_(_1709_v20, _1689_v3)).Dtor_cf2(), 12)
				_1710_v21 = _nw285
				var _1711_v22 _dafny.Map
				_ = _1711_v22
				_1711_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D13_.Create_DC31_(_1689_v3), _1710_v21)
				var _1712_v23 D13
				_ = _1712_v23
				_1712_v23 = Companion_D13_.Create_DC31_(_1689_v3)
				var _1713_v24 _dafny.Array
				_ = _1713_v24
				var _nwElement0_66 _dafny.Array = _1710_v21
				_ = _nwElement0_66
				var _nw286 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(12))
				_ = _nw286
				(_nw286).ArraySet1(_nwElement0_66, 0)
				(_nw286).ArraySet1(_1710_v21, 1)
				(_nw286).ArraySet1(_1710_v21, 2)
				(_nw286).ArraySet1(_1710_v21, 3)
				(_nw286).ArraySet1(_1710_v21, 4)
				(_nw286).ArraySet1((func() _dafny.Array {
					if (_1711_v22).Contains(_1712_v23) {
						return (_1711_v22).Get(_1712_v23).(_dafny.Array)
					}
					return _1710_v21
				})(), 5)
				(_nw286).ArraySet1(_1710_v21, 6)
				(_nw286).ArraySet1(_1710_v21, 7)
				(_nw286).ArraySet1(_1710_v21, 8)
				(_nw286).ArraySet1(_1710_v21, 9)
				(_nw286).ArraySet1(_1710_v21, 10)
				(_nw286).ArraySet1(_1710_v21, 11)
				_1713_v24 = _nw286
				var _1714_v25 _dafny.Sequence
				_ = _1714_v25
				_1714_v25 = _dafny.SeqOf(_dafny.SeqOf(false, _1688_v2, _1686_v0), _1708_v19, _1708_v19)
				var _1715_v26 D6
				_ = _1715_v26
				_1715_v26 = Companion_D6_.Create_DC12_(_1688_v2, _1714_v25, _1687_v1)
				var _1716_v27 _dafny.Sequence
				_ = _1716_v27
				_1716_v27 = _dafny.SeqOf(_1715_v26)
				var _1717_v28 _dafny.Set
				_ = _1717_v28
				_1717_v28 = _dafny.SetOf(_1687_v1, _1687_v1, _1687_v1)
				var _rhs301 _dafny.Int = ((_this).Fm4(_1717_v28, _1688_v2, _dafny.SeqOf(_1686_v0), globalState)).Times((_1689_v3).Plus(_1689_v3))
				_ = _rhs301
				var _rhs302 _dafny.Array = _1713_v24
				_ = _rhs302
				var _rhs303 bool = _1688_v2
				_ = _rhs303
				var _rhs304 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer100 func(_dafny.Int) D6) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1718_v26 D6) func(_dafny.Int) D6 {
					return func(_1719_i1 _dafny.Int) D6 {
						return _1718_v26
					}
				})(_1715_v26)))
				_ = _rhs304
				var _rhs305 _dafny.Int = _dafny.IntOfUint32((_1687_v1).Cardinality())
				_ = _rhs305
				var _lhs236 *GlobalState = globalState
				_ = _lhs236
				var _lhs237 *GlobalState = globalState
				_ = _lhs237
				_lhs236.F0 = _rhs301
				_1713_v24 = _rhs302
				_1686_v0 = _rhs303
				_1716_v27 = _rhs304
				_lhs237.F21 = _rhs305
				var _1720_v29 _dafny.MultiSet
				_ = _1720_v29
				_1720_v29 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("wtdcmrxwf"))
				(globalState).F3 = ((_1720_v29).Update(_1687_v1, Companion_Default___.Abs((_this).F27()))).Cardinality()
				var _1721_v30 *C3
				_ = _1721_v30
				var _nw287 *C3 = New_C3_()
				_ = _nw287
				_nw287.Ctor__(_dafny.IntOfInt64(918), (_this).F28())
				_1721_v30 = _nw287
				var _1722_v31 T0
				_ = _1722_v31
				var _nw288 *C3 = New_C3_()
				_ = _nw288
				_nw288.Ctor__(_1689_v3, (_this).F28())
				_1722_v31 = _nw288
				var _1723_v32 _dafny.Map
				_ = _1723_v32
				_1723_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1722_v31, _dafny.CodePoint('k'))
				_1723_v32 = _1723_v32
				var _1724_v33 *C3
				_ = _1724_v33
				var _nw289 *C3 = New_C3_()
				_ = _nw289
				_nw289.Ctor__((_this).F27(), (_1722_v31).F28())
				_1724_v33 = _nw289
			} else {
				(globalState).F21 = _1689_v3
				var _1725_v34 _dafny.Sequence
				_ = _1725_v34
				_1725_v34 = _dafny.SeqOf(_1708_v19)
				var _1726_v35 D6
				_ = _1726_v35
				_1726_v35 = Companion_D6_.Create_DC12_(_1686_v0, _1725_v34, _1687_v1)
				(globalState).F26 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((_1726_v35).Dtor_cf21(), _1725_v34), _1725_v34)
				var _1727_v36 *C3
				_ = _1727_v36
				var _nw290 *C3 = New_C3_()
				_ = _nw290
				_nw290.Ctor__((_dafny.Zero).Minus(_1689_v3), (_this).F28())
				_1727_v36 = _nw290
				var _1728_v37 _dafny.CodePoint
				_ = _1728_v37
				_1728_v37 = _dafny.CodePoint('k')
				var _rhs306 bool = _1686_v0
				_ = _rhs306
				var _rhs307 _dafny.CodePoint = _1728_v37
				_ = _rhs307
				var _rhs308 _dafny.Int = (_this).F27()
				_ = _rhs308
				var _lhs238 *GlobalState = globalState
				_ = _lhs238
				_1688_v2 = _rhs306
				_1728_v37 = _rhs307
				_lhs238.F0 = _rhs308
				var _1729_v38 _dafny.Map
				_ = _1729_v38
				_1729_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1728_v37, _1686_v0)
				var _1730_v39 _dafny.Map
				_ = _1730_v39
				_1730_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1689_v3, _1728_v37)
				var _1731_v40 _dafny.Sequence
				_ = _1731_v40
				_1731_v40 = _dafny.SeqOf(((_this).F29()).Update(_1688_v2, Companion_Default___.Abs((_this).F27())))
				_1729_v38 = (_1729_v38).Update((func() _dafny.CodePoint {
					if (_1730_v39).Contains(_1689_v3) {
						return (_1730_v39).Get(_1689_v3).(_dafny.CodePoint)
					}
					return _1728_v37
				})(), ((_1731_v40).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1731_v40).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsDisjointFrom((_this).F29()))
			}
		} else {
			var _1732_v41 _dafny.Array
			_ = _1732_v41
			var _nwElement0_67 bool = _1686_v0
			_ = _nwElement0_67
			var _nw291 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(23))
			_ = _nw291
			(_nw291).ArraySet1(_nwElement0_67, 0)
			(_nw291).ArraySet1(_1686_v0, 1)
			(_nw291).ArraySet1(_1686_v0, 2)
			(_nw291).ArraySet1(_1686_v0, 3)
			(_nw291).ArraySet1(_1686_v0, 4)
			(_nw291).ArraySet1(_1686_v0, 5)
			(_nw291).ArraySet1(!(_1686_v0), 6)
			(_nw291).ArraySet1(_1686_v0, 7)
			(_nw291).ArraySet1(!(!(_1686_v0)), 8)
			(_nw291).ArraySet1(_1686_v0, 9)
			(_nw291).ArraySet1(_1686_v0, 10)
			(_nw291).ArraySet1(_1686_v0, 11)
			(_nw291).ArraySet1(_1686_v0, 12)
			(_nw291).ArraySet1(_1686_v0, 13)
			(_nw291).ArraySet1(!(_1686_v0), 14)
			(_nw291).ArraySet1(_1686_v0, 15)
			(_nw291).ArraySet1(_1686_v0, 16)
			(_nw291).ArraySet1(true, 17)
			(_nw291).ArraySet1(_1686_v0, 18)
			(_nw291).ArraySet1(true, 19)
			(_nw291).ArraySet1(_1686_v0, 20)
			(_nw291).ArraySet1(_1686_v0, 21)
			(_nw291).ArraySet1(_1686_v0, 22)
			_1732_v41 = _nw291
			var _1733_v42 _dafny.MultiSet
			_ = _1733_v42
			_1733_v42 = _dafny.MultiSetOf(_1732_v41, _1732_v41, _1732_v41)
			var _1734_v43 *C4
			_ = _1734_v43
			var _nw292 *C4 = New_C4_()
			_ = _nw292
			_nw292.Ctor__(((_1733_v42).Union(_1733_v42)).Cardinality(), (_this).F28())
			_1734_v43 = _nw292
			(globalState).F0 = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wibmqbfun")).Cardinality())).Times((_this).F27()), (_this).F27())
			(globalState).F21 = (_this).F27()
			if _1686_v0 {
				var _1735_v44 *C3
				_ = _1735_v44
				var _nw293 *C3 = New_C3_()
				_ = _nw293
				_nw293.Ctor__((_this).F27(), (func(_pat_let48_0 D17) D17 {
					return func(_1736_dt__update__tmp_h0 D17) D17 {
						return func(_pat_let49_0 _dafny.Array) D17 {
							return func(_1737_dt__update_hcf82_h0 _dafny.Array) D17 {
								return Companion_D17_.Create_DC43_(_1737_dt__update_hcf82_h0)
							}(_pat_let49_0)
						}((_this).F28())
					}(_pat_let48_0)
				}(Companion_D17_.Create_DC43_((_this).F28()))).Dtor_cf82())
				_1735_v44 = _nw293
				(globalState).F26 = _1686_v0
				_1686_v0 = _1686_v0
				var _1738_v45 _dafny.Map
				_ = _1738_v45
				_1738_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1686_v0)
				var _1739_v46 _dafny.Sequence
				_ = _1739_v46
				_1739_v46 = _dafny.UnicodeSeqOfUtf8Bytes("ikcash")
				_1738_v45 = (_1738_v45).Update(_1686_v0, (_dafny.IntOfUint32((_1739_v46).Cardinality())).Cmp((_this).F27()) > 0)
				var _1740_v47 _dafny.Array
				_ = _1740_v47
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_35
				var _nw294 _dafny.Array
				_ = _nw294
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw294 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) bool = (func(_1741_v0 bool) func(_dafny.Int) bool {
						return func(_1742_i2 _dafny.Int) bool {
							return _1741_v0
						}
					})(_1686_v0)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw294 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw294).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw294).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1740_v47 = _nw294
				var _1743_v48 _dafny.Sequence
				_ = _1743_v48
				_1743_v48 = _dafny.SeqOf(_1740_v47)
				r2 = !_dafny.Companion_Sequence_.Equal(_1743_v48, _1743_v48)
			} else {
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))
				_ = _index218
				(_1732_v41).ArraySet1(_1686_v0, (_index218).Int())
				var _1744_v49 _dafny.Array
				_ = _1744_v49
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_36
				var _nw295 _dafny.Array
				_ = _nw295
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw295 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = func(_1745_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1745_i3, (_this).F27())
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw295 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw295).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw295).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1744_v49 = _nw295
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1744_v49), 0))
				_ = _index219
				(_1744_v49).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm0((_dafny.Zero).Minus((_this).F27()), globalState) {
						return (_this).F27()
					}
					return _dafny.IntOfInt64(-75)
				})(), (_index219).Int())
				var _1746_v50 D17
				_ = _1746_v50
				_1746_v50 = Companion_D17_.Create_DC44_((_this).F27())
				var _1747_v51 D17
				_ = _1747_v51
				_1747_v51 = Companion_D17_.Create_DC45_(_1746_v50)
				var _1748_v52 _dafny.Map
				_ = _1748_v52
				_1748_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1686_v0)
				var _1749_v53 _dafny.Set
				_ = _1749_v53
				_1749_v53 = _dafny.SetOf(_1748_v52, _1748_v52, _1748_v52)
				var _1750_v54 _dafny.Sequence
				_ = _1750_v54
				_1750_v54 = _dafny.SeqOf(_1749_v53)
				var _1751_v55 D18
				_ = _1751_v55
				_1751_v55 = Companion_D18_.Create_DC46_((_1750_v54).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1750_v54).Cardinality()))).Uint32()).(_dafny.Set))
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))
				_ = _index220
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1744_v49), 0))
				_ = _index221
				var _rhs309 bool = !_dafny.Companion_Sequence_.Contains((_this).F30(), (_dafny.Zero).Minus((_this).F27()))
				_ = _rhs309
				var _rhs310 _dafny.Int = (_this).F27()
				_ = _rhs310
				var _rhs311 D17 = _1747_v51
				_ = _rhs311
				var _rhs312 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F27(), (((_1751_v55).Dtor_cf85()).Union(_dafny.SetOf(_1748_v52, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, false), _1748_v52))).Cardinality())
				_ = _rhs312
				var _rhs313 _dafny.Int = (_this).F27()
				_ = _rhs313
				var _lhs239 _dafny.Array = _1732_v41
				_ = _lhs239
				var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))
				_ = _lhs240
				var _lhs241 _dafny.Array = _1744_v49
				_ = _lhs241
				var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1744_v49), 0))
				_ = _lhs242
				var _lhs243 *GlobalState = globalState
				_ = _lhs243
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				(_lhs239).ArraySet1(_rhs309, (_lhs240).Int())
				(_lhs241).ArraySet1(_rhs310, (_lhs242).Int())
				_1747_v51 = _rhs311
				_lhs243.F3 = _rhs312
				_lhs244.F15 = _rhs313
				var _1752_v56 *C2
				_ = _1752_v56
				var _nw296 *C2 = New_C2_()
				_ = _nw296
				_nw296.Ctor__(_dafny.MultiSetOf((_1732_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))).Int()).(bool), _1686_v0, !((_1732_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))).Int()).(bool))), _dafny.IntOfInt64(268))
				_1752_v56 = _nw296
				var _1753_v57 _dafny.Array
				_ = _1753_v57
				var _nw297 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw297
				_1753_v57 = _nw297
				var _1754_v58 _dafny.Map
				_ = _1754_v58
				_1754_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1744_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1744_v49), 0))).Int()).(_dafny.Int), _1753_v57)
				_1754_v58 = _1754_v58
				var _1755_v59 _dafny.MultiSet
				_ = _1755_v59
				_1755_v59 = _dafny.MultiSetOf((_this).F27())
				var _1756_v60 _dafny.Set
				_ = _1756_v60
				_1756_v60 = _dafny.SetOf(_1755_v59, _1755_v59)
				var _1757_v61 _dafny.Map
				_ = _1757_v61
				_1757_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D17_.Create_DC44_((_this).F27()), _1756_v60)
				var _1758_v62 _dafny.Sequence
				_ = _1758_v62
				_1758_v62 = _dafny.SeqOf(_1686_v0)
				var _1759_v63 _dafny.Map
				_ = _1759_v63
				_1759_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1758_v62, (_1754_v58).Cardinality())
				var _1760_v64 D17
				_ = _1760_v64
				_1760_v64 = Companion_D17_.Create_DC44_((_1759_v63).Cardinality())
				_1757_v61 = (_1757_v61).Update((func() D17 {
					if _1686_v0 {
						return func(_pat_let50_0 D17) D17 {
							return func(_1761_dt__update__tmp_h1 D17) D17 {
								return func(_pat_let51_0 _dafny.Int) D17 {
									return func(_1762_dt__update_hcf83_h0 _dafny.Int) D17 {
										return Companion_D17_.Create_DC44_(_1762_dt__update_hcf83_h0)
									}(_pat_let51_0)
								}(_dafny.IntOfUint32(((_this).F30()).Cardinality()))
							}(_pat_let50_0)
						}(_1760_v64)
					}
					return _1760_v64
				})(), _1756_v60)
				var _1763_v65 _dafny.Set
				_ = _1763_v65
				_1763_v65 = _dafny.SetOf(_1744_v49)
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))
				_ = _index222
				var _rhs314 bool = ((_1763_v65).Difference(_dafny.SetOf(_1744_v49))).IsSubsetOf(_1763_v65)
				_ = _rhs314
				var _rhs315 bool = !(_1686_v0)
				_ = _rhs315
				var _rhs316 _dafny.Array = _1744_v49
				_ = _rhs316
				var _rhs317 bool = (_1732_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))).Int()).(bool)
				_ = _rhs317
				var _lhs245 _dafny.Array = _1732_v41
				_ = _lhs245
				var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_1732_v41), 0))
				_ = _lhs246
				var _lhs247 *GlobalState = globalState
				_ = _lhs247
				var _lhs248 *GlobalState = globalState
				_ = _lhs248
				(_lhs245).ArraySet1(_rhs314, (_lhs246).Int())
				_lhs247.F26 = _rhs315
				_1744_v49 = _rhs316
				_lhs248.F26 = _rhs317
			}
			var _1764_v66 _dafny.Set
			_ = _1764_v66
			_1764_v66 = _dafny.SetOf((_this).F27())
			var _1765_v67 D9
			_ = _1765_v67
			_1765_v67 = Companion_D9_.Create_DC17_(_1764_v66)
			var _1766_v68 _dafny.Map
			_ = _1766_v68
			_1766_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1765_v67, (_this).F27())
			var _1767_v69 _dafny.Sequence
			_ = _1767_v69
			_1767_v69 = _dafny.UnicodeSeqOfUtf8Bytes("heeuww")
			var _1768_v70 *C0
			_ = _1768_v70
			var _nw298 *C0 = New_C0_()
			_ = _nw298
			_nw298.Ctor__((_1766_v68).Cardinality(), _1767_v69)
			_1768_v70 = _nw298
			var _1769_v71 _dafny.Sequence
			_ = _1769_v71
			_1769_v71 = _dafny.SeqOf(_1768_v70)
			var _1770_v72 _dafny.Int
			_ = _1770_v72
			var _out46 _dafny.Int
			_ = _out46
			_out46 = (_1734_v43).M11((_1769_v71).Select((Companion_Default___.SafeIndex((_1768_v70).F36(), _dafny.IntOfUint32((_1769_v71).Cardinality()))).Uint32()).(*C0), _1686_v0, globalState)
			_1770_v72 = _out46
		}
		var _1771_v73 _dafny.CodePoint
		_ = _1771_v73
		_1771_v73 = _dafny.CodePoint('g')
		var _1772_v74 _dafny.Sequence
		_ = _1772_v74
		_1772_v74 = _dafny.SeqOf(_1771_v73)
		var _1773_v75 _dafny.MultiSet
		_ = _1773_v75
		_1773_v75 = _dafny.MultiSetOf((_this).F27(), (_this).F27())
		var _1774_v76 _dafny.Set
		_ = _1774_v76
		_1774_v76 = _dafny.SetOf(_1772_v74, _dafny.Companion_Sequence_.Update(_1772_v74, (Companion_Default___.SafeIndex((_1773_v75).Cardinality(), _dafny.IntOfUint32((_1772_v74).Cardinality()))).Uint32(), _1771_v73), _1772_v74, _1772_v74)
		var _1775_v77 _dafny.Map
		_ = _1775_v77
		_1775_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1771_v73, _1686_v0)
		var _1776_v78 _dafny.Map
		_ = _1776_v78
		_1776_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), false)
		var _1777_v79 _dafny.Sequence
		_ = _1777_v79
		_1777_v79 = _dafny.SeqOf((func() bool {
			if (_1776_v78).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1772_v74, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1772_v74).Cardinality()))).Uint32(), _1771_v73)).Cardinality())) {
				return (_1776_v78).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1772_v74, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1772_v74).Cardinality()))).Uint32(), _1771_v73)).Cardinality())).(bool)
			}
			return _1686_v0
		})())
		var _1778_v80 D15
		_ = _1778_v80
		_1778_v80 = Companion_D15_.Create_DC36_((func() _dafny.CodePoint {
			if _1686_v0 {
				return _1771_v73
			}
			return _1771_v73
		})(), (_this).Fm4(_1774_v76, (func() bool {
			if (_1775_v77).Contains(_1771_v73) {
				return (_1775_v77).Get(_1771_v73).(bool)
			}
			return _1686_v0
		})(), _1777_v79, globalState))
		_1778_v80 = _1778_v80
		_1772_v74 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm31(_1686_v0, _1686_v0, globalState), _1772_v74)
		var _1779_i4 _dafny.Int
		_ = _1779_i4
		_1779_i4 = _dafny.Zero
		{
			for _1686_v0 {
				{
					if (_1779_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1779_i4 = (_1779_i4).Plus(_dafny.One)
					var _1780_v81 *C1
					_ = _1780_v81
					var _nw299 *C1 = New_C1_()
					_ = _nw299
					_nw299.Ctor__(((_this).F27()).Cmp(_dafny.IntOfUint32(((_this).F30()).Cardinality())) >= 0, (_this).F29(), (_this).F30(), (_this).F27(), (_this).F28())
					_1780_v81 = _nw299
					var _1781_v82 _dafny.Sequence
					_ = _1781_v82
					_1781_v82 = _dafny.SeqOf(_1780_v81)
					_1780_v81 = (func() *C1 {
						if _1686_v0 {
							return (_1781_v82).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F30()).Cardinality())), _dafny.IntOfUint32((_1781_v82).Cardinality()))).Uint32()).(*C1)
						}
						return _1780_v81
					})()
					_1772_v74 = _1772_v74
					(globalState).F21 = (_this).F27()
					var _1782_v83 _dafny.MultiSet
					_ = _1782_v83
					_1782_v83 = _dafny.MultiSetOf(_1771_v73)
					var _1783_v84 _dafny.Map
					_ = _1783_v84
					_1783_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1686_v0) && (_1780_v81.F35), _1782_v83)
					_1783_v84 = (_1783_v84).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_1777_v79, _1777_v79), _1782_v83)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1784_v85 *C3
		_ = _1784_v85
		var _nw300 *C3 = New_C3_()
		_ = _nw300
		_nw300.Ctor__((_this).F27(), (_this).F28())
		_1784_v85 = _nw300
		if _1686_v0 {
			_1686_v0 = Companion_Default___.Fm0(_dafny.IntOfInt64(860), globalState)
			var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index223
			((_this).F28()).ArraySet1(_1772_v74, (_index223).Int())
			var _1785_v86 _dafny.Array
			_ = _1785_v86
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_37
			var _nw301 _dafny.Array
			_ = _nw301
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw301 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) bool = (func(_1786_v0 bool) func(_dafny.Int) bool {
					return func(_1787_i5 _dafny.Int) bool {
						return _1786_v0
					}
				})(_1686_v0)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw301 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw301).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw301).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1785_v86 = _nw301
			var _1788_v87 _dafny.Set
			_ = _1788_v87
			_1788_v87 = _dafny.SetOf((_this).F27())
			var _1789_v88 D4
			_ = _1789_v88
			_1789_v88 = Companion_D4_.Create_DC8_(_1771_v73, (_1788_v87).Cardinality(), _1686_v0)
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))
			_ = _index224
			(_1785_v86).ArraySet1((_1777_v79).Select((Companion_Default___.SafeIndex((_1789_v88).Dtor_cf14(), _dafny.IntOfUint32((_1777_v79).Cardinality()))).Uint32()).(bool), (_index224).Int())
			var _1790_v89 D6
			_ = _1790_v89
			_1790_v89 = Companion_D6_.Create_DC12_(Companion_Default___.Fm0((_this).F27(), globalState), Companion_Default___.Fm44(globalState), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm31(_1686_v0, _1686_v0, globalState), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((Companion_Default___.Fm31(_1686_v0, _1686_v0, globalState)).Cardinality()))).Uint32(), Companion_Default___.Fm38(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1686_v0), globalState)))
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index225
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))
			_ = _index226
			var _rhs318 _dafny.Sequence = (_1790_v89).Dtor_cf22()
			_ = _rhs318
			var _rhs319 bool = (_1686_v0) == (false)
			_ = _rhs319
			var _lhs249 _dafny.Array = (_this).F28()
			_ = _lhs249
			var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _lhs250
			var _lhs251 _dafny.Array = _1785_v86
			_ = _lhs251
			var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))
			_ = _lhs252
			(_lhs249).ArraySet1(_rhs318, (_lhs250).Int())
			(_lhs251).ArraySet1(_rhs319, (_lhs252).Int())
			var _1791_v90 _dafny.Array
			_ = _1791_v90
			var _len0_38 _dafny.Int = _dafny.One
			_ = _len0_38
			var _nw302 _dafny.Array
			_ = _nw302
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw302 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Map = (func(_1792_v73 _dafny.CodePoint, _1793_v0 bool) func(_dafny.Int) _dafny.Map {
					return func(_1794_i6 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer101 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
							return func(arg101 _dafny.Int) interface{} {
								return coer101(arg101)
							}
						}((func(_1795_v73 _dafny.CodePoint) func(_dafny.Int) D1 {
							return func(_1796_i7 _dafny.Int) D1 {
								return Companion_D1_.Create_DC2_(_1795_v73, (_this).F27())
							}
						})(_1792_v73))), _1793_v0)
					}
				})(_1771_v73, _1686_v0)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw302 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw302).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw302).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1791_v90 = _nw302
			var _1797_v91 _dafny.Map
			_ = _1797_v91
			_1797_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(-418))
			var _1798_v92 _dafny.Map
			_ = _1798_v92
			_1798_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1785_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))).Int()).(bool), _dafny.SetOf(_1686_v0))
			var _1799_v93 D1
			_ = _1799_v93
			_1799_v93 = Companion_D1_.Create_DC2_(_1771_v73, (func() _dafny.Int {
				if (_1797_v91).Contains(_dafny.IntOfUint32((_1777_v79).Cardinality())) {
					return (_1797_v91).Get(_dafny.IntOfUint32((_1777_v79).Cardinality())).(_dafny.Int)
				}
				return (_1798_v92).Cardinality()
			})())
			var _1800_v94 _dafny.Sequence
			_ = _1800_v94
			_1800_v94 = _dafny.SeqOf(_1799_v93)
			var _1801_v95 _dafny.Map
			_ = _1801_v95
			_1801_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1800_v94, (_1785_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))).Int()).(bool))
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1791_v90), 0))
			_ = _index227
			(_1791_v90).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1800_v94, (_1784_v85).Fm46((_this).F27(), _1771_v73, globalState))).Merge(_1801_v95), (_index227).Int())
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1791_v90), 0))
			_ = _index228
			(_1791_v90).ArraySet1(_1801_v95, (_index228).Int())
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))
			_ = _index229
			(_1785_v86).ArraySet1((_1785_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))).Int()).(bool), (_index229).Int())
			var _1802_v96 _dafny.Sequence
			_ = _1802_v96
			_1802_v96 = _dafny.SeqOf((_this).F30())
			var _1803_v97 _dafny.Map
			_ = _1803_v97
			_1803_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1771_v73)
			var _1804_v98 _dafny.Map
			_ = _1804_v98
			_1804_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1802_v96).Cardinality())).Cmp((_1803_v97).Cardinality()) != 0, (_this).F27())
			(globalState).F21 = (func() _dafny.Int {
				if (_1804_v98).Contains((_1785_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))).Int()).(bool)) {
					return (_1804_v98).Get((_1785_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1785_v86), 0))).Int()).(bool)).(_dafny.Int)
				}
				return (_this).F27()
			})()
		} else {
			_1686_v0 = _1686_v0
			var _1805_v99 _dafny.Set
			_ = _1805_v99
			_1805_v99 = _dafny.SetOf(_1771_v73, _1771_v73)
			var _1806_v100 D19
			_ = _1806_v100
			_1806_v100 = Companion_D19_.Create_DC48_(_1805_v99)
			(globalState).F3 = (_dafny.IntOfInt64(363)).Minus(((func() _dafny.Set {
				var _coll54 = _dafny.NewBuilder()
				_ = _coll54
				for _iter62 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm37((_this).F27(), ((_1806_v100).Dtor_cf88()).Cardinality(), globalState), _1686_v0)).Keys().Elements()); ; {
					_compr_54, _ok62 := _iter62()
					if !_ok62 {
						break
					}
					var _1807_v101 _dafny.Sequence
					_1807_v101 = interface{}(_compr_54).(_dafny.Sequence)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm37((_this).F27(), ((_1806_v100).Dtor_cf88()).Cardinality(), globalState), _1686_v0)).Contains(_1807_v101) {
						_coll54.Add(_1807_v101)
					}
				}
				return _coll54.ToSet()
			}()).Union(_1774_v76)).Cardinality())
			var _1808_v102 _dafny.Map
			_ = _1808_v102
			_1808_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1686_v0)
			_1808_v102 = (_1808_v102).Update(((_this).F27()).Cmp(_dafny.IntOfUint32((_1772_v74).Cardinality())) == 0, _1686_v0)
			var _1809_v103 _dafny.Map
			_ = _1809_v103
			_1809_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())
			var _1810_v104 T0
			_ = _1810_v104
			var _nw303 *C3 = New_C3_()
			_ = _nw303
			_nw303.Ctor__((_this).F27(), (_this).F28())
			_1810_v104 = _nw303
			var _1811_v105 _dafny.Array
			_ = _1811_v105
			var _nwElement0_68 _dafny.Int = (_this).F27()
			_ = _nwElement0_68
			var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(29))
			_ = _nw304
			(_nw304).ArraySet1(_nwElement0_68, 0)
			(_nw304).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1809_v103).Contains((_this).F27()) {
					return (_1809_v103).Get((_this).F27()).(_dafny.Int)
				}
				return (Companion_Default___.Fm48(_1686_v0, _1686_v0, (_dafny.SetOf(_1810_v104)).Cardinality(), globalState)).Cardinality()
			})(), (_this).F27())).Cardinality(), 1)
			(_nw304).ArraySet1(_dafny.IntOfInt64(477), 2)
			(_nw304).ArraySet1((_1810_v104).F27(), 3)
			(_nw304).ArraySet1((_this).F27(), 4)
			(_nw304).ArraySet1((_1810_v104).F27(), 5)
			(_nw304).ArraySet1(_dafny.IntOfUint32((_1772_v74).Cardinality()), 6)
			(_nw304).ArraySet1((_this).F27(), 7)
			(_nw304).ArraySet1((_this).F27(), 8)
			(_nw304).ArraySet1((_this).F27(), 9)
			(_nw304).ArraySet1((_1810_v104).F27(), 10)
			(_nw304).ArraySet1((_dafny.Zero).Minus((_this).F27()), 11)
			(_nw304).ArraySet1((_1776_v78).Cardinality(), 12)
			(_nw304).ArraySet1((_this).F27(), 13)
			(_nw304).ArraySet1((_this).F27(), 14)
			(_nw304).ArraySet1((_this).F27(), 15)
			(_nw304).ArraySet1(_dafny.IntOfUint32((_1772_v74).Cardinality()), 16)
			(_nw304).ArraySet1((_1810_v104).F27(), 17)
			(_nw304).ArraySet1((_this).F27(), 18)
			(_nw304).ArraySet1((_1810_v104).F27(), 19)
			(_nw304).ArraySet1(_dafny.IntOfInt64(386), 20)
			(_nw304).ArraySet1((_1810_v104).F27(), 21)
			(_nw304).ArraySet1((_1810_v104).F27(), 22)
			(_nw304).ArraySet1((_this).F27(), 23)
			(_nw304).ArraySet1((_this).F27(), 24)
			(_nw304).ArraySet1((_this).Fm4(_1774_v76, false, _1777_v79, globalState), 25)
			(_nw304).ArraySet1((_1810_v104).F27(), 26)
			(_nw304).ArraySet1((_1810_v104).F27(), 27)
			(_nw304).ArraySet1(_dafny.IntOfInt64(235), 28)
			_1811_v105 = _nw304
			var _1812_v106 _dafny.Set
			_ = _1812_v106
			_1812_v106 = _dafny.SetOf(_1811_v105, _1811_v105, _1811_v105, _1811_v105)
			var _1813_v107 _dafny.Map
			_ = _1813_v107
			_1813_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm3((_this).F29(), globalState)), _1809_v103)
			var _1814_v108 _dafny.Array
			_ = _1814_v108
			var _nwElement0_69 _dafny.Map = _1809_v103
			_ = _nwElement0_69
			var _nw305 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(17))
			_ = _nw305
			(_nw305).ArraySet1(_nwElement0_69, 0)
			(_nw305).ArraySet1(_1809_v103, 1)
			(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27()), 2)
			(_nw305).ArraySet1((_1809_v103).Merge(_1809_v103), 3)
			(_nw305).ArraySet1(_1809_v103, 4)
			(_nw305).ArraySet1((_1809_v103).Update((_this).F27(), (_this).F27()), 5)
			(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27()), 6)
			(_nw305).ArraySet1(_1809_v103, 7)
			(_nw305).ArraySet1(_1809_v103, 8)
			(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-591), (_1812_v106).Cardinality()), 9)
			(_nw305).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_1810_v104).F27())).Update(_dafny.IntOfUint32((_1777_v79).Cardinality()), (_this).F27()), 10)
			(_nw305).ArraySet1(_1809_v103, 11)
			(_nw305).ArraySet1((func() _dafny.Map {
				if _1686_v0 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1810_v104).F27(), (_this).F27())
				}
				return _1809_v103
			})(), 12)
			(_nw305).ArraySet1(_1809_v103, 13)
			(_nw305).ArraySet1((func() _dafny.Map {
				if (_1813_v107).Contains((func() bool {
					if (_1808_v102).Contains((_1777_v79).Select((Companion_Default___.SafeIndex((_1810_v104).F27(), _dafny.IntOfUint32((_1777_v79).Cardinality()))).Uint32()).(bool)) {
						return (_1808_v102).Get((_1777_v79).Select((Companion_Default___.SafeIndex((_1810_v104).F27(), _dafny.IntOfUint32((_1777_v79).Cardinality()))).Uint32()).(bool)).(bool)
					}
					return _1686_v0
				})()) {
					return (_1813_v107).Get((func() bool {
						if (_1808_v102).Contains((_1777_v79).Select((Companion_Default___.SafeIndex((_1810_v104).F27(), _dafny.IntOfUint32((_1777_v79).Cardinality()))).Uint32()).(bool)) {
							return (_1808_v102).Get((_1777_v79).Select((Companion_Default___.SafeIndex((_1810_v104).F27(), _dafny.IntOfUint32((_1777_v79).Cardinality()))).Uint32()).(bool)).(bool)
						}
						return _1686_v0
					})()).(_dafny.Map)
				}
				return _1809_v103
			})(), 14)
			(_nw305).ArraySet1(_1809_v103, 15)
			(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.IntOfUint32((_1772_v74).Cardinality())), 16)
			_1814_v108 = _nw305
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1814_v108), 0))
			_ = _index230
			(_1814_v108).ArraySet1((_1809_v103).Merge(_1809_v103), (_index230).Int())
			var _1815_v110 _dafny.Map
			_ = _1815_v110
			_1815_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1810_v104).F27(), _1771_v73)
			var _1816_v111 _dafny.Map
			_ = _1816_v111
			_1816_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1810_v104).F27(), _1815_v110)
			var _1817_v112 _dafny.Sequence
			_ = _1817_v112
			_1817_v112 = _dafny.SeqOf(func() _dafny.Map {
				var _coll55 = _dafny.NewMapBuilder()
				_ = _coll55
				for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(124), _dafny.IntOfInt64(382))); ; {
					_compr_55, _ok63 := _iter63()
					if !_ok63 {
						break
					}
					var _1818_v109 _dafny.Int
					_1818_v109 = interface{}(_compr_55).(_dafny.Int)
					if ((_dafny.IntOfInt64(124)).Cmp(_1818_v109) <= 0) && ((_1818_v109).Cmp(_dafny.IntOfInt64(382)) < 0) {
						_coll55.Add(Companion_Default___.SafeModuloInt(_1818_v109, (_1810_v104).F27()), _dafny.CodePoint('t'))
					}
				}
				return _coll55.ToMap()
			}(), (_1815_v110).Update(_dafny.IntOfUint32((_dafny.SeqOf(!(_1686_v0))).Cardinality()), _1771_v73), _1815_v110, (func() _dafny.Map {
				if (_1816_v111).Contains((_this).F27()) {
					return (_1816_v111).Get((_this).F27()).(_dafny.Map)
				}
				return _1815_v110
			})())
			var _1819_v113 D15
			_ = _1819_v113
			_1819_v113 = Companion_D15_.Create_DC35_(_1817_v112)
			var _1820_v114 _dafny.Sequence
			_ = _1820_v114
			_1820_v114 = _dafny.SeqOf(Companion_D15_.Create_DC35_(_1817_v112), _1819_v113)
			var _1821_v115 _dafny.Sequence
			_ = _1821_v115
			_1821_v115 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer102 func(_dafny.Int) D15) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1822_v104 T0, _1823_v73 _dafny.CodePoint) func(_dafny.Int) D15 {
				return func(_1824_i8 _dafny.Int) D15 {
					return Companion_D15_.Create_DC35_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1822_v104).F27(), _1823_v73)))
				}
			})(_1810_v104, _1771_v73))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-918), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer103 func(_dafny.Int) D15) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}((func(_1825_v104 T0, _1826_v73 _dafny.CodePoint) func(_dafny.Int) D15 {
				return func(_1827_i8 _dafny.Int) D15 {
					return Companion_D15_.Create_DC35_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1825_v104).F27(), _1826_v73)))
				}
			})(_1810_v104, _1771_v73)))).Cardinality()))).Uint32(), _1819_v113), _dafny.Companion_Sequence_.Concatenate(_1820_v114, _1820_v114), _dafny.Companion_Sequence_.Concatenate(_1820_v114, _1820_v114))
			var _1828_v116 *C0
			_ = _1828_v116
			var _nw306 *C0 = New_C0_()
			_ = _nw306
			_nw306.Ctor__((_this).F27(), _1772_v74)
			_1828_v116 = _nw306
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1814_v108), 0))
			_ = _index231
			var _rhs320 _dafny.Map = ((_1809_v103).Merge(_1809_v103)).Merge(_1809_v103)
			_ = _rhs320
			var _rhs321 bool = _1686_v0
			_ = _rhs321
			var _rhs322 _dafny.Sequence = _1821_v115
			_ = _rhs322
			var _rhs323 *C0 = _1828_v116
			_ = _rhs323
			var _lhs253 _dafny.Array = _1814_v108
			_ = _lhs253
			var _lhs254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1814_v108), 0))
			_ = _lhs254
			(_lhs253).ArraySet1(_rhs320, (_lhs254).Int())
			r2 = _rhs321
			_1821_v115 = _rhs322
			_1828_v116 = _rhs323
			var _1829_v117 _dafny.Sequence
			_ = _1829_v117
			_1829_v117 = _dafny.SeqOf(_1811_v105)
			_1829_v117 = _dafny.Companion_Sequence_.Concatenate(_1829_v117, _1829_v117)
		}
		var _1830_v118 D9
		_ = _1830_v118
		_1830_v118 = Companion_D9_.Create_DC19_(!(_1686_v0))
		var _1831_v119 _dafny.Set
		_ = _1831_v119
		_1831_v119 = _dafny.SetOf(true, true, _1686_v0, _1686_v0, _1686_v0)
		var _nwElement0_70 bool = !((_1830_v118).Dtor_cf33()) || (Companion_Default___.Fm0((_this).F27(), globalState))
		_ = _nwElement0_70
		var _nw307 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(3))
		_ = _nw307
		(_nw307).ArraySet1(_nwElement0_70, 0)
		(_nw307).ArraySet1(_1686_v0, 1)
		(_nw307).ArraySet1(!((_1831_v119).IsProperSubsetOf(_1831_v119)), 2)
		r0 = _nw307
		var _1832_v120 D7
		_ = _1832_v120
		_1832_v120 = Companion_D7_.Create_DC13_(_1777_v79)
		var _pat_let_tv37 = _1686_v0
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1686_v0
		_ = _pat_let_tv38
		r1 = func(_source26 D7) bool {
			if _source26.Is_DC14() {
				var _1833___mcc_h0 _dafny.Int = _source26.Get_().(D7_DC14).Cf24
				_ = _1833___mcc_h0
				var _1834___mcc_h1 T0 = _source26.Get_().(D7_DC14).Cf25
				_ = _1834___mcc_h1
				var _1835___mcc_h2 _dafny.Sequence = _source26.Get_().(D7_DC14).Cf26
				_ = _1835___mcc_h2
				var _1836___mcc_h3 _dafny.MultiSet = _source26.Get_().(D7_DC14).Cf27
				_ = _1836___mcc_h3
				var _1837_cf27 _dafny.MultiSet = _1836___mcc_h3
				_ = _1837_cf27
				var _1838_cf26 _dafny.Sequence = _1835___mcc_h2
				_ = _1838_cf26
				var _1839_cf25 T0 = _1834___mcc_h1
				_ = _1839_cf25
				var _1840_cf24 _dafny.Int = _1833___mcc_h0
				_ = _1840_cf24
				return _pat_let_tv37
			} else {
				var _1841___mcc_h4 _dafny.Sequence = _source26.Get_().(D7_DC13).Cf23
				_ = _1841___mcc_h4
				var _1842_cf23 _dafny.Sequence = _1841___mcc_h4
				_ = _1842_cf23
				return _pat_let_tv38
			}
		}(_1832_v120)
		r2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_this).F27(), (_this).F27()), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F30(), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), _1686_v0)).Cardinality())), (Companion_Default___.SafeIndex(((_dafny.MultiSetFromSeq((_this).F30())).Update((_this).F27(), Companion_Default___.Abs((_this).F27()))).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F30(), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), _1686_v0)).Cardinality()))).Cardinality()))).Uint32(), _dafny.IntOfInt64(669)))
		return r0, r1, r2
	}
}
func (_this *C5) M10(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) {
	{
		var _1843_v0 _dafny.CodePoint
		_ = _1843_v0
		_1843_v0 = _dafny.CodePoint('i')
		var _1844_v1 _dafny.Sequence
		_ = _1844_v1
		_1844_v1 = _dafny.UnicodeSeqOfUtf8Bytes("xggvd")
		var _1845_v2 D2
		_ = _1845_v2
		_1845_v2 = Companion_D2_.Create_DC4_((_this).F27(), _dafny.IntOfInt64(521), (_dafny.Zero).Minus((_this).F27()), _1843_v0, _1844_v1)
		var _1846_v3 _dafny.Set
		_ = _1846_v3
		_1846_v3 = _dafny.SetOf(_1845_v2)
		var _hi8 _dafny.Int = (func() _dafny.Int {
			if p0 {
				return (_1846_v3).Cardinality()
			}
			return (_this).F27()
		})()
		_ = _hi8
		for _1847_i0 := (func() _dafny.Int {
			if p0 {
				return (_this).F27()
			}
			return (_this).F27()
		})(); _1847_i0.Cmp(_hi8) < 0; _1847_i0 = _1847_i0.Plus(_dafny.One) {
			var _1848_v4 _dafny.Map
			_ = _1848_v4
			_1848_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p0)
			var _1849_v5 *C2
			_ = _1849_v5
			var _nw308 *C2 = New_C2_()
			_ = _nw308
			_nw308.Ctor__(_dafny.MultiSetOf(p0), ((_1848_v4).Merge(_1848_v4)).Cardinality())
			_1849_v5 = _nw308
			(globalState).F26 = p0
			var _1850_v6 _dafny.Sequence
			_ = _1850_v6
			_1850_v6 = _dafny.SeqOf(p0)
			_1850_v6 = _1850_v6
			var _1851_v7 _dafny.Array
			_ = _1851_v7
			var _nw309 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
			_ = _nw309
			_1851_v7 = _nw309
			var _1852_v8 _dafny.Map
			_ = _1852_v8
			_1852_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1850_v6, _1851_v7)
			var _1853_v9 _dafny.Map
			_ = _1853_v9
			_1853_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1850_v6, (func() _dafny.Array {
				if (_1852_v8).Contains(_1850_v6) {
					return (_1852_v8).Get(_1850_v6).(_dafny.Array)
				}
				return _1851_v7
			})())
			_1853_v9 = (_1853_v9).Update(_dafny.SeqOf(!(p0), p0), _1851_v7)
		}
		var _1854_v10 _dafny.Array
		_ = _1854_v10
		var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
		_ = _nw310
		_1854_v10 = _nw310
		var _1855_v11 D16
		_ = _1855_v11
		_1855_v11 = Companion_D16_.Create_DC39_(_1854_v10)
		var _pat_let_tv39 = _1854_v10
		_ = _pat_let_tv39
		var _pat_let_tv40 = _1854_v10
		_ = _pat_let_tv40
		_1855_v11 = func(_pat_let52_0 D16) D16 {
			return func(_1858_dt__update__tmp_h1 D16) D16 {
				return func(_pat_let55_0 _dafny.Array) D16 {
					return func(_1859_dt__update_hcf73_h1 _dafny.Array) D16 {
						return Companion_D16_.Create_DC39_(_1859_dt__update_hcf73_h1)
					}(_pat_let55_0)
				}(_pat_let_tv40)
			}(_pat_let52_0)
		}(func(_pat_let53_0 D16) D16 {
			return func(_1856_dt__update__tmp_h0 D16) D16 {
				return func(_pat_let54_0 _dafny.Array) D16 {
					return func(_1857_dt__update_hcf73_h0 _dafny.Array) D16 {
						return Companion_D16_.Create_DC39_(_1857_dt__update_hcf73_h0)
					}(_pat_let54_0)
				}(_pat_let_tv39)
			}(_pat_let53_0)
		}(_1855_v11))
		var _1860_v12 _dafny.Sequence
		_ = _1860_v12
		_1860_v12 = _dafny.SeqOf(((_this).F27()).Cmp(_dafny.IntOfInt64(386)) <= 0)
		var _rhs324 _dafny.Sequence = _dafny.SeqOf(p0, p0)
		_ = _rhs324
		var _rhs325 _dafny.Int = Companion_Default___.Fm1(globalState)
		_ = _rhs325
		var _rhs326 _dafny.Int = (_this).F27()
		_ = _rhs326
		var _lhs255 *GlobalState = globalState
		_ = _lhs255
		var _lhs256 *GlobalState = globalState
		_ = _lhs256
		_1860_v12 = _rhs324
		_lhs255.F24 = _rhs325
		_lhs256.F0 = _rhs326
		var _1861_v13 _dafny.Map
		_ = _1861_v13
		_1861_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(425), (_dafny.Zero).Minus((_this).F27()))
		var _1862_v14 _dafny.Map
		_ = _1862_v14
		_1862_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1861_v13).Contains(((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_1861_v13).Get(((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_this).F27())
		})(), false)
		_1862_v14 = (_1862_v14).Update((Companion_Default___.Fm48(p0, p0, _dafny.IntOfInt64(-276), globalState)).Cardinality(), p0)
		_1843_v0 = _1843_v0
		var _1863_v15 _dafny.Map
		_ = _1863_v15
		_1863_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm0((_this).F27(), globalState)), true)
		var _1864_v16 D8
		_ = _1864_v16
		_1864_v16 = Companion_D8_.Create_DC15_((_1863_v15).Merge(_1863_v15))
		var _source27 D8 = _1864_v16
		_ = _source27
		if _source27.Is_DC16() {
			var _1865___mcc_h0 _dafny.Int = _source27.Get_().(D8_DC16).Cf29
			_ = _1865___mcc_h0
			var _1866___mcc_h1 _dafny.Int = _source27.Get_().(D8_DC16).Cf30
			_ = _1866___mcc_h1
			var _1867_cf30 _dafny.Int = _1866___mcc_h1
			_ = _1867_cf30
			var _1868_cf29 _dafny.Int = _1865___mcc_h0
			_ = _1868_cf29
			var _1869_v17 _dafny.Map
			_ = _1869_v17
			_1869_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("jiyhtfc"))
			var _1870_v18 _dafny.Array
			_ = _1870_v18
			var _nwElement0_71 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("rnsowybbf")
			_ = _nwElement0_71
			var _nw311 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(20))
			_ = _nw311
			(_nw311).ArraySet1(_nwElement0_71, 0)
			(_nw311).ArraySet1(_1844_v1, 1)
			(_nw311).ArraySet1(_1844_v1, 2)
			(_nw311).ArraySet1(Companion_Default___.Fm31(p0, false, globalState), 3)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1844_v1, _1844_v1), 4)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm37((_this).F27(), _dafny.IntOfInt64(422), globalState), _1844_v1), 5)
			(_nw311).ArraySet1(_1844_v1, 6)
			(_nw311).ArraySet1(_1844_v1, 7)
			(_nw311).ArraySet1(_1844_v1, 8)
			(_nw311).ArraySet1(_1844_v1, 9)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yo"), _1844_v1), (Companion_Default___.SafeIndex(_1867_cf30, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yo"), _1844_v1)).Cardinality()))).Uint32(), _1843_v0), 10)
			(_nw311).ArraySet1(_1844_v1, 11)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1844_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-646))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1871_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1872_i1 _dafny.Int) _dafny.CodePoint {
					return _1871_v0
				}
			})(_1843_v0)))), 12)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-14))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1873_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1874_i2 _dafny.Int) _dafny.CodePoint {
					return _1873_v0
				}
			})(_1843_v0))), _1844_v1), 13)
			(_nw311).ArraySet1(_1844_v1, 14)
			(_nw311).ArraySet1(_1844_v1, 15)
			(_nw311).ArraySet1((func() _dafny.Sequence {
				if (_1869_v17).Contains(true) {
					return (_1869_v17).Get(true).(_dafny.Sequence)
				}
				return _1844_v1
			})(), 16)
			(_nw311).ArraySet1(_1844_v1, 17)
			(_nw311).ArraySet1(_1844_v1, 18)
			(_nw311).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ggdic"), 19)
			_1870_v18 = _nw311
			var _1875_v19 _dafny.Map
			_ = _1875_v19
			_1875_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1844_v1).Cardinality())).Minus(_dafny.IntOfUint32((_1844_v1).Cardinality())), (_this).F28())
			_1870_v18 = (func() _dafny.Array {
				if (_1875_v19).Contains(_1868_cf29) {
					return (_1875_v19).Get(_1868_cf29).(_dafny.Array)
				}
				return (_this).F28()
			})()
			var _1876_v20 _dafny.Array
			_ = _1876_v20
			var _nw312 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(16))
			_ = _nw312
			_1876_v20 = _nw312
			var _1877_v21 _dafny.Map
			_ = _1877_v21
			_1877_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1876_v20, _1843_v0)
			var _1878_v22 _dafny.Map
			_ = _1878_v22
			_1878_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1868_cf29, _1843_v0)
			var _rhs327 _dafny.Int = ((_this).F27()).Times((p1).Select((Companion_Default___.SafeIndex(_1868_cf29, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs327
			var _rhs328 bool = p0
			_ = _rhs328
			var _rhs329 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1876_v20, (func() _dafny.CodePoint {
				if (_1878_v22).Contains((_this).F27()) {
					return (_1878_v22).Get((_this).F27()).(_dafny.CodePoint)
				}
				return _1843_v0
			})())
			_ = _rhs329
			var _rhs330 _dafny.Int = _dafny.IntOfInt64(-672)
			_ = _rhs330
			var _lhs257 *GlobalState = globalState
			_ = _lhs257
			var _lhs258 *GlobalState = globalState
			_ = _lhs258
			var _lhs259 *GlobalState = globalState
			_ = _lhs259
			_lhs257.F0 = _rhs327
			_lhs258.F26 = _rhs328
			_1877_v21 = _rhs329
			_lhs259.F18 = _rhs330
			var _1879_v23 _dafny.Sequence
			_ = _1879_v23
			_1879_v23 = _dafny.SeqOf(_1844_v1)
			var _1880_v24 T0
			_ = _1880_v24
			var _nw313 *C4 = New_C4_()
			_ = _nw313
			_nw313.Ctor__((_dafny.IntOfUint32((_1844_v1).Cardinality())).Minus(_1868_cf29), (_this).F28())
			_1880_v24 = _nw313
			var _1881_v25 _dafny.Array
			_ = _1881_v25
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_39
			var _nw314 _dafny.Array
			_ = _nw314
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw314 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Int = (func(_1882_cf29 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1883_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1883_i3, _1882_cf29)
					}
				})(_1868_cf29)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw314 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw314).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw314).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1881_v25 = _nw314
			var _1884_v26 _dafny.MultiSet
			_ = _1884_v26
			_1884_v26 = _dafny.MultiSetOf((_dafny.IntOfInt64(737)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1844_v1, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1844_v1).Cardinality()))).Uint32(), _1843_v0)).Cardinality())) < 0, false, p0, (func() bool {
				if p0 {
					return p0
				}
				return p0
			})())
			var _1885_v27 _dafny.Map
			_ = _1885_v27
			_1885_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F27())
			var _rhs331 _dafny.Sequence = _1879_v23
			_ = _rhs331
			var _rhs332 T0 = _1880_v24
			_ = _rhs332
			var _rhs333 _dafny.Array = (func() _dafny.Array {
				if p0 {
					return _1881_v25
				}
				return _1881_v25
			})()
			_ = _rhs333
			var _rhs334 _dafny.Int = (_1868_cf29).Minus(((_1885_v27).Cardinality()).Plus((_this).F27()))
			_ = _rhs334
			var _rhs335 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm41(_1868_cf29, (_1862_v14).Cardinality(), (_this).Fm3(_dafny.MultiSetOf(p0), globalState), globalState), _1860_v12), (Companion_Default___.SafeIndex(_1868_cf29, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm41(_1868_cf29, (_1862_v14).Cardinality(), (_this).Fm3(_dafny.MultiSetOf(p0), globalState), globalState), _1860_v12)).Cardinality()))).Uint32(), p0))
			_ = _rhs335
			var _lhs260 *GlobalState = globalState
			_ = _lhs260
			_1879_v23 = _rhs331
			_1880_v24 = _rhs332
			_1881_v25 = _rhs333
			_lhs260.F0 = _rhs334
			_1884_v26 = _rhs335
			var _1886_v28 *C0
			_ = _1886_v28
			var _nw315 *C0 = New_C0_()
			_ = _nw315
			_nw315.Ctor__(_dafny.IntOfUint32((_1844_v1).Cardinality()), _1844_v1)
			_1886_v28 = _nw315
			var _1887_v29 D6
			_ = _1887_v29
			_1887_v29 = Companion_D6_.Create_DC11_(_1886_v28)
			var _source28 D6 = _1887_v29
			_ = _source28
			if _source28.Is_DC12() {
				var _1888___mcc_h3 bool = _source28.Get_().(D6_DC12).Cf20
				_ = _1888___mcc_h3
				var _1889___mcc_h4 _dafny.Sequence = _source28.Get_().(D6_DC12).Cf21
				_ = _1889___mcc_h4
				var _1890___mcc_h5 _dafny.Sequence = _source28.Get_().(D6_DC12).Cf22
				_ = _1890___mcc_h5
				var _1891_cf22 _dafny.Sequence = _1890___mcc_h5
				_ = _1891_cf22
				var _1892_cf21 _dafny.Sequence = _1889___mcc_h4
				_ = _1892_cf21
				var _1893_cf20 bool = _1888___mcc_h3
				_ = _1893_cf20
				var _1894_v30 *C0
				_ = _1894_v30
				var _nw316 *C0 = New_C0_()
				_ = _nw316
				_nw316.Ctor__((_1880_v24).F27(), (_1886_v28).F37())
				_1894_v30 = _nw316
				_1868_cf29 = (_1886_v28).F36()
				var _1895_v31 D4
				_ = _1895_v31
				_1895_v31 = Companion_D4_.Create_DC8_(_1843_v0, _1867_cf30, false)
				var _1896_v32 _dafny.Set
				_ = _1896_v32
				_1896_v32 = _dafny.SetOf((_1894_v30).F36(), (_1894_v30).F36(), _1868_cf29, _1868_cf29, (_1880_v24).F27())
				var _1897_v33 _dafny.Array
				_ = _1897_v33
				var _nwElement0_72 bool = p0
				_ = _nwElement0_72
				var _nw317 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(15))
				_ = _nw317
				(_nw317).ArraySet1(_nwElement0_72, 0)
				(_nw317).ArraySet1(p0, 1)
				(_nw317).ArraySet1((_this).Fm3((_this).F29(), globalState), 2)
				(_nw317).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1860_v12, _1860_v12), 3)
				(_nw317).ArraySet1(_1893_cf20, 4)
				(_nw317).ArraySet1(p0, 5)
				(_nw317).ArraySet1(p0, 6)
				(_nw317).ArraySet1((_1896_v32).Contains(_1867_cf30), 7)
				(_nw317).ArraySet1(_1893_cf20, 8)
				(_nw317).ArraySet1((Companion_Default___.Fm49((_this).F27(), globalState)).Dtor_cf17(), 9)
				(_nw317).ArraySet1(_1893_cf20, 10)
				(_nw317).ArraySet1((_1868_cf29).Cmp((_1880_v24).F27()) >= 0, 11)
				(_nw317).ArraySet1(p0, 12)
				(_nw317).ArraySet1(false, 13)
				(_nw317).ArraySet1(_1893_cf20, 14)
				_1897_v33 = _nw317
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1897_v33), 0))
				_ = _index232
				(_1897_v33).ArraySet1(false, (_index232).Int())
				var _pat_let_tv41 = _1894_v30
				_ = _pat_let_tv41
				var _pat_let_tv42 = _1868_cf29
				_ = _pat_let_tv42
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1897_v33), 0))
				_ = _index233
				var _rhs336 _dafny.Int = _1867_cf30
				_ = _rhs336
				var _rhs337 D4 = func(_pat_let56_0 D4) D4 {
					return func(_1898_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let57_0 bool) D4 {
							return func(_1899_dt__update_hcf15_h0 bool) D4 {
								return Companion_D4_.Create_DC8_((_1898_dt__update__tmp_h2).Dtor_cf13(), (_1898_dt__update__tmp_h2).Dtor_cf14(), _1899_dt__update_hcf15_h0)
							}(_pat_let57_0)
						}(((_pat_let_tv41).F36()).Cmp(_pat_let_tv42) == 0)
					}(_pat_let56_0)
				}(Companion_D4_.Create_DC8_(_1843_v0, (_1886_v28).F36(), _1893_cf20))
				_ = _rhs337
				var _rhs338 bool = !(((_1886_v28).F36()).Cmp(_dafny.IntOfUint32((_1860_v12).Cardinality())) > 0)
				_ = _rhs338
				var _lhs261 *GlobalState = globalState
				_ = _lhs261
				var _lhs262 _dafny.Array = _1897_v33
				_ = _lhs262
				var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1897_v33), 0))
				_ = _lhs263
				_lhs261.F0 = _rhs336
				_1895_v31 = _rhs337
				(_lhs262).ArraySet1(_rhs338, (_lhs263).Int())
				(globalState).F24 = (_1894_v30).F36()
			} else {
				var _1900___mcc_h6 *C0 = _source28.Get_().(D6_DC11).Cf19
				_ = _1900___mcc_h6
				var _1901_cf19 *C0 = _1900___mcc_h6
				_ = _1901_cf19
				var _rhs339 _dafny.Array = _1881_v25
				_ = _rhs339
				var _rhs340 bool = (func() bool {
					if (_1862_v14).Contains((_1880_v24).F27()) {
						return (_1862_v14).Get((_1880_v24).F27()).(bool)
					}
					return p0
				})()
				_ = _rhs340
				var _rhs341 _dafny.Int = ((_1880_v24).F27()).Times(_1868_cf29)
				_ = _rhs341
				var _rhs342 bool = (func() bool {
					if p0 {
						return !(p0)
					}
					return false
				})()
				_ = _rhs342
				var _rhs343 _dafny.Map = _1861_v13
				_ = _rhs343
				var _lhs264 *GlobalState = globalState
				_ = _lhs264
				var _lhs265 *GlobalState = globalState
				_ = _lhs265
				var _lhs266 *GlobalState = globalState
				_ = _lhs266
				_1881_v25 = _rhs339
				_lhs264.F26 = _rhs340
				_lhs265.F24 = _rhs341
				_lhs266.F26 = _rhs342
				_1861_v13 = _rhs343
				_1863_v15 = (_1863_v15).Update(p0, p0)
				var _1902_v34 *C1
				_ = _1902_v34
				var _nw318 *C1 = New_C1_()
				_ = _nw318
				_nw318.Ctor__(p0, (func() _dafny.MultiSet {
					if p0 {
						return _1884_v26
					}
					return (_this).F29()
				})(), (_this).F30(), ((_1880_v24).F27()).Times(_1868_cf29), (_1880_v24).F28())
				_1902_v34 = _nw318
				(globalState).F0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1844_v1).Cardinality()), (_1886_v28).F36())
			}
		} else {
			var _1903___mcc_h2 _dafny.Map = _source27.Get_().(D8_DC15).Cf28
			_ = _1903___mcc_h2
			var _1904_cf28 _dafny.Map = _1903___mcc_h2
			_ = _1904_cf28
			var _1905_v35 _dafny.Array
			_ = _1905_v35
			var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw319
			_1905_v35 = _nw319
			var _1906_v36 *C0
			_ = _1906_v36
			var _nw320 *C0 = New_C0_()
			_ = _nw320
			_nw320.Ctor__((_this).F27(), _1844_v1)
			_1906_v36 = _nw320
			var _1907_v37 _dafny.Array
			_ = _1907_v37
			var _nw321 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw321
			_1907_v37 = _nw321
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1907_v37), 0))
			_ = _index234
			(_1907_v37).ArraySet1(p0, (_index234).Int())
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1907_v37), 0))
			_ = _index235
			var _rhs344 _dafny.Array = _1905_v35
			_ = _rhs344
			var _rhs345 *C0 = _1906_v36
			_ = _rhs345
			var _rhs346 bool = !(p0)
			_ = _rhs346
			var _lhs267 _dafny.Array = _1907_v37
			_ = _lhs267
			var _lhs268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1907_v37), 0))
			_ = _lhs268
			_1905_v35 = _rhs344
			_1906_v36 = _rhs345
			(_lhs267).ArraySet1(_rhs346, (_lhs268).Int())
			var _1908_v38 _dafny.MultiSet
			_ = _1908_v38
			_1908_v38 = _dafny.MultiSetOf((_1907_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1907_v37), 0))).Int()).(bool))
			var _1909_v39 _dafny.Array
			_ = _1909_v39
			var _nw322 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw322
			_1909_v39 = _nw322
			var _1910_v40 *C3
			_ = _1910_v40
			var _nw323 *C3 = New_C3_()
			_ = _nw323
			_nw323.Ctor__((_1906_v36).F36(), (_this).F28())
			_1910_v40 = _nw323
			var _1911_v41 _dafny.Sequence
			_ = _1911_v41
			_1911_v41 = _dafny.SeqOf(_1910_v40, _1910_v40)
			var _1912_v42 _dafny.Array
			_ = _1912_v42
			var _nwElement0_73 _dafny.Int = (_1906_v36).F36()
			_ = _nwElement0_73
			var _nw324 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(26))
			_ = _nw324
			(_nw324).ArraySet1(_nwElement0_73, 0)
			(_nw324).ArraySet1((_1906_v36).F36(), 1)
			(_nw324).ArraySet1((_this).F27(), 2)
			(_nw324).ArraySet1((_this).F27(), 3)
			(_nw324).ArraySet1(_dafny.IntOfUint32((_1911_v41).Cardinality()), 4)
			(_nw324).ArraySet1((_1863_v15).Cardinality(), 5)
			(_nw324).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("gtwh"), (Companion_Default___.SafeIndex((_1906_v36).F36(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gtwh")).Cardinality()))).Uint32(), _1843_v0)).Cardinality()), 6)
			(_nw324).ArraySet1((_this).F27(), 7)
			(_nw324).ArraySet1((_1906_v36).F36(), 8)
			(_nw324).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qrfkfxp")).Cardinality()), 9)
			(_nw324).ArraySet1(Companion_Default___.Fm1(globalState), 10)
			(_nw324).ArraySet1((_this).F27(), 11)
			(_nw324).ArraySet1((_1906_v36).F36(), 12)
			(_nw324).ArraySet1(((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), 13)
			(_nw324).ArraySet1((_1906_v36).F36(), 14)
			(_nw324).ArraySet1(_dafny.IntOfInt64(943), 15)
			(_nw324).ArraySet1((_1906_v36).F36(), 16)
			(_nw324).ArraySet1((_this).F27(), 17)
			(_nw324).ArraySet1(_dafny.IntOfInt64(719), 18)
			(_nw324).ArraySet1((_this).F27(), 19)
			(_nw324).ArraySet1(_dafny.IntOfInt64(306), 20)
			(_nw324).ArraySet1((_this).F27(), 21)
			(_nw324).ArraySet1(_dafny.IntOfInt64(947), 22)
			(_nw324).ArraySet1(_dafny.IntOfInt64(-688), 23)
			(_nw324).ArraySet1((_this).F27(), 24)
			(_nw324).ArraySet1((_this).F27(), 25)
			_1912_v42 = _nw324
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1909_v39), 0))
			_ = _index236
			(_1909_v39).ArraySet1(_1912_v42, (_index236).Int())
			var _1913_v43 _dafny.MultiSet
			_ = _1913_v43
			_1913_v43 = _dafny.MultiSetOf(_1843_v0)
			var _1914_v44 *C2
			_ = _1914_v44
			var _nw325 *C2 = New_C2_()
			_ = _nw325
			_nw325.Ctor__(_1908_v38, (_this).F27())
			_1914_v44 = _nw325
			var _1915_v45 _dafny.Map
			_ = _1915_v45
			_1915_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1914_v44, (_1907_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1907_v37), 0))).Int()).(bool))
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1909_v39), 0))
			_ = _index237
			var _rhs347 _dafny.Int = ((func() _dafny.Int {
				if (_1913_v43).Contains(_1843_v0) {
					return (_1913_v43).Multiplicity(_1843_v0)
				}
				return (_1915_v45).Cardinality()
			})()).Minus((_1906_v36).F36())
			_ = _rhs347
			var _rhs348 _dafny.MultiSet = _1908_v38
			_ = _rhs348
			var _rhs349 _dafny.Array = _1912_v42
			_ = _rhs349
			var _rhs350 bool = p0
			_ = _rhs350
			var _lhs269 *GlobalState = globalState
			_ = _lhs269
			var _lhs270 _dafny.Array = _1909_v39
			_ = _lhs270
			var _lhs271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1909_v39), 0))
			_ = _lhs271
			var _lhs272 *GlobalState = globalState
			_ = _lhs272
			_lhs269.F15 = _rhs347
			_1908_v38 = _rhs348
			(_lhs270).ArraySet1(_rhs349, (_lhs271).Int())
			_lhs272.F26 = _rhs350
			var _rhs351 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F27(), (_this).F27())
			_ = _rhs351
			var _rhs352 _dafny.Array = _1907_v37
			_ = _rhs352
			var _lhs273 *GlobalState = globalState
			_ = _lhs273
			_lhs273.F24 = _rhs351
			_1907_v37 = _rhs352
			_1862_v14 = Companion_Default___.Fm48(!(true), false, (_dafny.Zero).Minus((_1906_v36).F36()), globalState)
		}
	}
}
func (_this *C5) F38() _dafny.Map {
	{
		return _this._f38
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f29 _dafny.MultiSet
	_f30 _dafny.Sequence
	_f27 _dafny.Int
	_f28 _dafny.Array
	_f31 _dafny.MultiSet
	_f32 _dafny.Int
	F34  _dafny.Int
	_f33 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f29 = _dafny.EmptyMultiSet
	_this._f30 = _dafny.EmptySeq
	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f31 = _dafny.EmptyMultiSet
	_this._f32 = _dafny.Zero
	_this.F34 = _dafny.Zero
	_this._f33 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C6{}
var _ T2 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F29() _dafny.MultiSet {
	return _this._f29
}
func (_this *C6) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C6) F27() _dafny.Int {
	return _this._f27
}
func (_this *C6) F28() _dafny.Array {
	return _this._f28
}
func (_this *C6) F31() _dafny.MultiSet {
	return _this._f31
}
func (_this *C6) F32() _dafny.Int {
	return _this._f32
}
func (_this *C6) Ctor__(f33 bool, f34 _dafny.Int, f29 _dafny.MultiSet, f30 _dafny.Sequence, f27 _dafny.Int, f28 _dafny.Array, f31 _dafny.MultiSet, f32 _dafny.Int) {
	{
		(_this)._f33 = f33
		(_this).F34 = f34
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f27 = f27
		(_this)._f28 = f28
		(_this)._f31 = f31
		(_this)._f32 = f32
	}
}
func (_this *C6) Fm3(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (_this).F33()
	}
}
func (_this *C6) Fm4(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F33(), false, (_this).F33(), false)).Cardinality()), (_this).F33())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rio")).Cardinality()))).Cardinality()
	}
}
func (_this *C6) Fm2(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_this).F27(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spkmp")).Cardinality()), _dafny.IntOfInt64(4))).IsDisjointFrom(func() _dafny.Set {
			var _coll56 = _dafny.NewBuilder()
			_ = _coll56
			for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-947), _dafny.IntOfInt64(-397))); ; {
				_compr_56, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1916_v0 _dafny.Int
				_1916_v0 = interface{}(_compr_56).(_dafny.Int)
				if ((_dafny.IntOfInt64(-947)).Cmp(_1916_v0) <= 0) && ((_1916_v0).Cmp(_dafny.IntOfInt64(-397)) < 0) {
					_coll56.Add((_1916_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F33(), (_this).F33())).Cardinality())))
				}
			}
			return _coll56.ToSet()
		}()), (_this).F32())
	}
}
func (_this *C6) Fm5(globalState *GlobalState) bool {
	{
		return (_this).F33()
	}
}
func (_this *C6) Fm6(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-678))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}(func(_1917_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((_this).F33(), (_this).F33())
		}))
	}
}
func (_this *C6) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return ((Companion_D1_.Create_DC1_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, !((_this).F33())))).Dtor_cf1()).Cardinality()
	}
}
func (_this *C6) M1(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1918_v0 _dafny.CodePoint
		_ = _1918_v0
		_1918_v0 = _dafny.CodePoint('f')
		var _1919_v1 _dafny.Map
		_ = _1919_v1
		_1919_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v0, (_this).F33())
		var _1920_v2 _dafny.Sequence
		_ = _1920_v2
		_1920_v2 = _dafny.SeqOf(_1919_v1, _1919_v1)
		var _1921_v3 D2
		_ = _1921_v3
		_1921_v3 = Companion_D2_.Create_DC3_(_1919_v1)
		var _1922_v6 _dafny.Set
		_ = _1922_v6
		_1922_v6 = _dafny.SetOf(_1918_v0)
		var _1923_v7 _dafny.Array
		_ = _1923_v7
		var _nwElement0_74 _dafny.Map = _1919_v1
		_ = _nwElement0_74
		var _nw326 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(22))
		_ = _nw326
		(_nw326).ArraySet1(_nwElement0_74, 0)
		(_nw326).ArraySet1((_1919_v1).Update(_1918_v0, !((_this).Fm3((_this).F29(), globalState))), 1)
		(_nw326).ArraySet1(_1919_v1, 2)
		(_nw326).ArraySet1((_1920_v2).Select((Companion_Default___.SafeIndex(_this.F34, _dafny.IntOfUint32((_1920_v2).Cardinality()))).Uint32()).(_dafny.Map), 3)
		(_nw326).ArraySet1((_1921_v3).Dtor_cf4(), 4)
		(_nw326).ArraySet1(_1919_v1, 5)
		(_nw326).ArraySet1(_1919_v1, 6)
		(_nw326).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v0, (_this).F33()), 7)
		(_nw326).ArraySet1((_1919_v1).Merge(func() _dafny.Map {
			var _coll57 = _dafny.NewMapBuilder()
			_ = _coll57
			for _iter65 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v0, (_this).F33())).Keys().Elements()); ; {
				_compr_57, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1924_v4 _dafny.CodePoint
				_1924_v4 = interface{}(_compr_57).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v0, (_this).F33())).Contains(_1924_v4) {
					_coll57.Add(_1924_v4, (_this).F33())
				}
			}
			return _coll57.ToMap()
		}()), 8)
		(_nw326).ArraySet1(_1919_v1, 9)
		(_nw326).ArraySet1(_1919_v1, 10)
		(_nw326).ArraySet1(_1919_v1, 11)
		(_nw326).ArraySet1((_1921_v3).Dtor_cf4(), 12)
		(_nw326).ArraySet1(_1919_v1, 13)
		(_nw326).ArraySet1(_1919_v1, 14)
		(_nw326).ArraySet1((_1919_v1).Merge(_1919_v1), 15)
		(_nw326).ArraySet1((_1920_v2).Select((Companion_Default___.SafeIndex(_this.F34, _dafny.IntOfUint32((_1920_v2).Cardinality()))).Uint32()).(_dafny.Map), 16)
		(_nw326).ArraySet1((func() _dafny.Map {
			if (_this).F33() {
				return _1919_v1
			}
			return _1919_v1
		})(), 17)
		(_nw326).ArraySet1((_1919_v1).Merge(_1919_v1), 18)
		(_nw326).ArraySet1(_1919_v1, 19)
		(_nw326).ArraySet1((func() _dafny.Map {
			var _coll58 = _dafny.NewMapBuilder()
			_ = _coll58
			for _iter66 := _dafny.Iterate((_1922_v6).Elements()); ; {
				_compr_58, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1925_v5 _dafny.CodePoint
				_1925_v5 = interface{}(_compr_58).(_dafny.CodePoint)
				if (_1922_v6).Contains(_1925_v5) {
					_coll58.Add(_1925_v5, (_this).F33())
				}
			}
			return _coll58.ToMap()
		}()).Merge(_1919_v1), 20)
		(_nw326).ArraySet1(_1919_v1, 21)
		_1923_v7 = _nw326
		for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1923_v7), 0))); ; {
			_guard_loop_7, _ok67 := _iter67()
			if !_ok67 {
				break
			}
			var _1926_i0 _dafny.Int
			_1926_i0 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1926_i0).Sign() != -1) && ((_1926_i0).Cmp(_dafny.ArrayLen((_1923_v7), 0)) < 0)) {
				(_1923_v7).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v0, (_this).F33()), (_1926_i0).Int())
			}
		}
		var _1927_v8 _dafny.Map
		_ = _1927_v8
		_1927_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true) || (!(!((_this).F33()))), (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(110), _dafny.IntOfUint32((Companion_Default___.Fm8(globalState)).Cardinality()))).Cardinality())).Cmp((_this).F27()) > 0)
		_1927_v8 = (_1927_v8).Update((_this).F33(), (_this).F33())
		(globalState).F26 = (_this).F33()
		(globalState).F18 = ((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F30()).Cardinality()), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)
		_1918_v0 = _1918_v0
		var _1928_v9 _dafny.Set
		_ = _1928_v9
		_1928_v9 = _dafny.SetOf((_this).F33())
		var _1929_v10 _dafny.Map
		_ = _1929_v10
		_1929_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), _1928_v9)
		var _hi9 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm9((_this).F33(), globalState), (Companion_Default___.SafeIndex(_this.F34, _dafny.IntOfUint32((Companion_Default___.Fm9((_this).F33(), globalState)).Cardinality()))).Uint32(), Companion_Default___.Fm1(globalState))).Cardinality())
		_ = _hi9
		for _1930_i1 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg107 _dafny.Int) interface{} {
				return coer107(arg107)
			}
		}((func(_1955_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1956_i2 _dafny.Int) _dafny.CodePoint {
				return _1955_v0
			}
		})(_1918_v0))), (Companion_Default___.SafeIndex((_1929_v10).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg108 _dafny.Int) interface{} {
				return coer108(arg108)
			}
		}((func(_1957_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1958_i2 _dafny.Int) _dafny.CodePoint {
				return _1957_v0
			}
		})(_1918_v0)))).Cardinality()))).Uint32(), _1918_v0)).Cardinality()); _1930_i1.Cmp(_hi9) < 0; _1930_i1 = _1930_i1.Plus(_dafny.One) {
			var _1931_v11 _dafny.Map
			_ = _1931_v11
			_1931_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F33())
			var _1932_v12 _dafny.MultiSet
			_ = _1932_v12
			_1932_v12 = _dafny.MultiSetOf(_1931_v11)
			var _1933_v13 _dafny.Sequence
			_ = _1933_v13
			_1933_v13 = _dafny.UnicodeSeqOfUtf8Bytes("yhq")
			var _1934_v14 _dafny.Map
			_ = _1934_v14
			_1934_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1932_v12).Cardinality(), _1933_v13)
			if (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1934_v14).Contains((_this).F27()) {
					return (_1934_v14).Get((_this).F27()).(_dafny.Sequence)
				}
				return _1933_v13
			})()).Cardinality()), (_this).F32())).Cmp((_this).F27()) == 0 {
				var _1935_v15 _dafny.Sequence
				_ = _1935_v15
				_1935_v15 = _dafny.SeqOf((func() bool {
					if (_1931_v11).Contains(_dafny.IntOfUint32(((_this).F30()).Cardinality())) {
						return (_1931_v11).Get(_dafny.IntOfUint32(((_this).F30()).Cardinality())).(bool)
					}
					return (_this).F33()
				})())
				var _1936_v16 _dafny.Array
				_ = _1936_v16
				var _nwElement0_75 bool = (_this).F33()
				_ = _nwElement0_75
				var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(28))
				_ = _nw327
				(_nw327).ArraySet1(_nwElement0_75, 0)
				(_nw327).ArraySet1((_this).F33(), 1)
				(_nw327).ArraySet1((_this).F33(), 2)
				(_nw327).ArraySet1((_this).F33(), 3)
				(_nw327).ArraySet1((_this).F33(), 4)
				(_nw327).ArraySet1((_this).F33(), 5)
				(_nw327).ArraySet1((_this).F33(), 6)
				(_nw327).ArraySet1((_this).F33(), 7)
				(_nw327).ArraySet1(!((_this).F33()), 8)
				(_nw327).ArraySet1((_this).F33(), 9)
				(_nw327).ArraySet1((_this).F33(), 10)
				(_nw327).ArraySet1((_this).F33(), 11)
				(_nw327).ArraySet1((_this).F33(), 12)
				(_nw327).ArraySet1(!((_this).F33()), 13)
				(_nw327).ArraySet1(true, 14)
				(_nw327).ArraySet1((_this).F33(), 15)
				(_nw327).ArraySet1((_this).F33(), 16)
				(_nw327).ArraySet1((_1935_v15).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1935_v15).Cardinality()))).Uint32()).(bool), 17)
				(_nw327).ArraySet1((_this).F33(), 18)
				(_nw327).ArraySet1((_this).F33(), 19)
				(_nw327).ArraySet1((_this).F33(), 20)
				(_nw327).ArraySet1((_this).F33(), 21)
				(_nw327).ArraySet1((_this).F33(), 22)
				(_nw327).ArraySet1(!((_this).F33()), 23)
				(_nw327).ArraySet1((_this).F33(), 24)
				(_nw327).ArraySet1((_this).F33(), 25)
				(_nw327).ArraySet1((_this).F33(), 26)
				(_nw327).ArraySet1((_this).F33(), 27)
				_1936_v16 = _nw327
				var _1937_v17 _dafny.Map
				_ = _1937_v17
				_1937_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F32()).Times((_1931_v11).Cardinality()), _1936_v16)
				_1937_v17 = (_1937_v17).Update((_this).F32(), _1936_v16)
				(globalState).F26 = (func() bool {
					if (_1927_v8).Contains(((_1928_v9).Cardinality()).Cmp((_this).F27()) >= 0) {
						return (_1927_v8).Get(((_1928_v9).Cardinality()).Cmp((_this).F27()) >= 0).(bool)
					}
					return (_this).F33()
				})()
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))
				_ = _index238
				(_1936_v16).ArraySet1(!((_this).F33()) || ((_this).F33()), (_index238).Int())
				var _1938_v18 _dafny.Array
				_ = _1938_v18
				var _nw328 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw328
				_1938_v18 = _nw328
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1938_v18), 0))
				_ = _index239
				(_1938_v18).ArraySet1((_this).F27(), (_index239).Int())
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))
				_ = _index240
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1938_v18), 0))
				_ = _index241
				var _rhs353 bool = (_this).Fm3(_dafny.MultiSetOf((_this).F33()), globalState)
				_ = _rhs353
				var _rhs354 _dafny.Int = _dafny.IntOfInt64(-998)
				_ = _rhs354
				var _rhs355 _dafny.Array = _1936_v16
				_ = _rhs355
				var _lhs274 _dafny.Array = _1936_v16
				_ = _lhs274
				var _lhs275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))
				_ = _lhs275
				var _lhs276 _dafny.Array = _1938_v18
				_ = _lhs276
				var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1938_v18), 0))
				_ = _lhs277
				(_lhs274).ArraySet1(_rhs353, (_lhs275).Int())
				(_lhs276).ArraySet1(_rhs354, (_lhs277).Int())
				_1936_v16 = _rhs355
				var _1939_v19 _dafny.MultiSet
				_ = _1939_v19
				_1939_v19 = _dafny.MultiSetOf((_dafny.IntOfUint32((_1935_v15).Cardinality())).Minus(_1930_i1), (_dafny.Zero).Minus(_1930_i1), (_this).F27(), _dafny.IntOfUint32((_1933_v13).Cardinality()), ((_1938_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1938_v18), 0))).Int()).(_dafny.Int)).Times(_1930_i1))
				_1939_v19 = _1939_v19
				var _1940_v20 _dafny.Map
				_ = _1940_v20
				_1940_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1936_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))).Int()).(bool), (_this).Fm7(globalState))
				var _1941_v22 _dafny.Set
				_ = _1941_v22
				_1941_v22 = _dafny.SetOf((func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(215), _dafny.IntOfInt64(925))); ; {
						_compr_59, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1942_v21 _dafny.Int
						_1942_v21 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(215)).Cmp(_1942_v21) <= 0) && ((_1942_v21).Cmp(_dafny.IntOfInt64(925)) < 0) {
							_coll59.Add(Companion_Default___.SafeDivisionInt(_1942_v21, (_dafny.SetOf(_this.F34)).Cardinality()), (_this).F33())
						}
					}
					return _coll59.ToMap()
				}()).Cardinality())
				var _1943_v23 _dafny.Map
				_ = _1943_v23
				_1943_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1940_v20, ((func() _dafny.Set {
					if (_1936_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))).Int()).(bool) {
						return _1941_v22
					}
					return Companion_Default___.Fm10(globalState)
				})()).Cardinality())
				_1943_v23 = (func() _dafny.Map {
					if (_1936_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1936_v16), 0))).Int()).(bool) {
						return _1943_v23
					}
					return _1943_v23
				})()
			} else {
				r2 = (_this.F34).Plus((_1930_i1).Minus((_this).F32()))
				r0 = (_this).F32()
				var _1944_v24 _dafny.Map
				_ = _1944_v24
				_1944_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1928_v9).Cardinality(), _1927_v8)
				var _1945_v25 _dafny.Sequence
				_ = _1945_v25
				_1945_v25 = _dafny.SeqOf(_1927_v8, _1927_v8, (func() _dafny.Map {
					if (_1944_v24).Contains(_1930_i1) {
						return (_1944_v24).Get(_1930_i1).(_dafny.Map)
					}
					return _1927_v8
				})())
				var _1946_v26 _dafny.Sequence
				_ = _1946_v26
				_1946_v26 = _dafny.SeqOf((_1945_v25).Select((Companion_Default___.SafeIndex(_1930_i1, _dafny.IntOfUint32((_1945_v25).Cardinality()))).Uint32()).(_dafny.Map), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F33()), (_1927_v8).Merge(_1927_v8), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F33())).Merge(_1927_v8), _1927_v8)
				_1946_v26 = _1946_v26
				var _1947_v27 _dafny.Map
				_ = _1947_v27
				_1947_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1919_v1).Contains(_1918_v0) {
						return (_1919_v1).Get(_1918_v0).(bool)
					}
					return (_this).F33()
				})(), (_this).F27())
				var _1948_v28 T1
				_ = _1948_v28
				var _nw329 *C1 = New_C1_()
				_ = _nw329
				_nw329.Ctor__(false, (_this).F29(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex(_1930_i1, _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), (_1947_v27).Cardinality()), (Companion_Default___.SafeIndex(_this.F34, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex(_1930_i1, _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), (_1947_v27).Cardinality())).Cardinality()))).Uint32(), (_this).F27()), (_this).F27(), (_this).F28())
				_1948_v28 = _nw329
				var _1949_v29 _dafny.Map
				_ = _1949_v29
				_1949_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1933_v13, _1948_v28)
				_1949_v29 = (_1949_v29).Update(_dafny.UnicodeSeqOfUtf8Bytes("nl"), _1948_v28)
				var _1950_v30 _dafny.Sequence
				_ = _1950_v30
				_1950_v30 = _dafny.SeqOf(!((_this).F33()), (_this).F33(), (_this).F33(), (_this).F33(), !(!((_this).F33())))
				var _1951_v31 _dafny.Array
				_ = _1951_v31
				var _nwElement0_76 _dafny.Sequence = _1950_v30
				_ = _nwElement0_76
				var _nw330 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.One)
				_ = _nw330
				(_nw330).ArraySet1(_nwElement0_76, 0)
				_1951_v31 = _nw330
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1951_v31), 0))
				_ = _index242
				(_1951_v31).ArraySet1(_1950_v30, (_index242).Int())
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1951_v31), 0))
				_ = _index243
				(_1951_v31).ArraySet1(_dafny.Companion_Sequence_.Update(_1950_v30, (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1950_v30).Cardinality()))).Uint32(), (_this).F33()), (_index243).Int())
			}
			var _1952_v32 _dafny.Array
			_ = _1952_v32
			var _nw331 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw331
			_1952_v32 = _nw331
			var _1953_v33 _dafny.Sequence
			_ = _1953_v33
			_1953_v33 = _dafny.SeqOf((_this).F33())
			var _1954_v34 T2
			_ = _1954_v34
			var _nw332 *C2 = New_C2_()
			_ = _nw332
			_nw332.Ctor__(_dafny.MultiSetFromSeq(_1953_v33), (_this).F32())
			_1954_v34 = _nw332
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1952_v32), 0))
			_ = _index244
			(_1952_v32).ArraySet1((func() T2 {
				if (_this).F33() {
					return _1954_v34
				}
				return _1954_v34
			})(), (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1952_v32), 0))
			_ = _index245
			(_1952_v32).ArraySet1(_1954_v34, (_index245).Int())
			var _rhs356 _dafny.Int = _1930_i1
			_ = _rhs356
			var _rhs357 bool = (_this).F33()
			_ = _rhs357
			var _lhs278 *GlobalState = globalState
			_ = _lhs278
			var _lhs279 *GlobalState = globalState
			_ = _lhs279
			_lhs278.F15 = _rhs356
			_lhs279.F26 = _rhs357
			_1933_v13 = _1933_v13
		}
		var _1959_v35 _dafny.Sequence
		_ = _1959_v35
		_1959_v35 = _dafny.UnicodeSeqOfUtf8Bytes("asxhdikym")
		r0 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm1(globalState), (_dafny.IntOfUint32((_1959_v35).Cardinality())).Times(_this.F34))
		r1 = (_this).F27()
		r2 = (_this).F27()
		return r0, r1, r2
	}
}
func (_this *C6) M2(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1960_v0 _dafny.Sequence
		_ = _1960_v0
		_1960_v0 = _dafny.SeqOf(p1)
		var _1961_v1 _dafny.Sequence
		_ = _1961_v1
		_1961_v1 = _dafny.SeqOf(_1960_v0)
		var _1962_v2 _dafny.Sequence
		_ = _1962_v2
		_1962_v2 = _dafny.UnicodeSeqOfUtf8Bytes("mmsbclr")
		var _pat_let_tv43 = _1961_v1
		_ = _pat_let_tv43
		var _1963_v3 bool
		_ = _1963_v3
		var _1964_v4 _dafny.Int
		_ = _1964_v4
		var _out47 bool
		_ = _out47
		var _out48 _dafny.Int
		_ = _out48
		_out47, _out48 = Companion_Default___.M0((_this).Fm5(globalState), ((_this).F27()).Plus(_dafny.IntOfUint32(((func(_pat_let58_0 D6) D6 {
			return func(_1965_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let59_0 _dafny.Sequence) D6 {
					return func(_1966_dt__update_hcf21_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC12_((_1965_dt__update__tmp_h0).Dtor_cf20(), _1966_dt__update_hcf21_h0, (_1965_dt__update__tmp_h0).Dtor_cf22())
					}(_pat_let59_0)
				}(_pat_let_tv43)
			}(_pat_let58_0)
		}(Companion_D6_.Create_DC12_((_this).F33(), _1961_v1, _1962_v2))).Dtor_cf21()).Cardinality())), (_dafny.Zero).Minus(_this.F34), _this.F34, globalState)
		_1963_v3 = _out47
		_1964_v4 = _out48
		var _1967_i0 _dafny.Int
		_ = _1967_i0
		_1967_i0 = _dafny.Zero
		{
			for ((_this).F29()).IsProperSubsetOf((_this).F29()) {
				{
					if (_1967_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1967_i0 = (_1967_i0).Plus(_dafny.One)
					var _1968_v5 _dafny.Map
					_ = _1968_v5
					_1968_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1963_v3, ((_this).F27()).Cmp((_this).F32()) < 0)
					_1968_v5 = (_1968_v5).Merge(_1968_v5)
					var _1969_v6 _dafny.Map
					_ = _1969_v6
					_1969_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, !(_1963_v3))
					(globalState).F15 = ((((_this).Fm2((_this).F33(), _1963_v3, p1, globalState)).Cardinality()).Times((_1969_v6).Cardinality())).Plus((_1969_v6).Cardinality())
					var _1970_v7 _dafny.Set
					_ = _1970_v7
					_1970_v7 = _dafny.SetOf((_this).Fm5(globalState))
					(globalState).F26 = ((_1970_v7).Intersection(_1970_v7)).Contains((_dafny.IntOfUint32((_1960_v0).Cardinality())).Cmp((_this).F32()) != 0)
					var _1971_v8 _dafny.Array
					_ = _1971_v8
					var _nwElement0_77 bool = _1963_v3
					_ = _nwElement0_77
					var _nw333 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(24))
					_ = _nw333
					(_nw333).ArraySet1(_nwElement0_77, 0)
					(_nw333).ArraySet1((_this).F33(), 1)
					(_nw333).ArraySet1(!(p0), 2)
					(_nw333).ArraySet1(false, 3)
					(_nw333).ArraySet1((_this).F33(), 4)
					(_nw333).ArraySet1(p1, 5)
					(_nw333).ArraySet1((_this).F33(), 6)
					(_nw333).ArraySet1(p1, 7)
					(_nw333).ArraySet1(_1963_v3, 8)
					(_nw333).ArraySet1(_1963_v3, 9)
					(_nw333).ArraySet1(p1, 10)
					(_nw333).ArraySet1(_1963_v3, 11)
					(_nw333).ArraySet1(!(p0), 12)
					(_nw333).ArraySet1(Companion_Default___.Fm0((_this).F27(), globalState), 13)
					(_nw333).ArraySet1((_this).F33(), 14)
					(_nw333).ArraySet1(true, 15)
					(_nw333).ArraySet1(p1, 16)
					(_nw333).ArraySet1(p1, 17)
					(_nw333).ArraySet1((_1960_v0).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1960_v0).Cardinality()))).Uint32()).(bool), 18)
					(_nw333).ArraySet1(_1963_v3, 19)
					(_nw333).ArraySet1(p0, 20)
					(_nw333).ArraySet1(p1, 21)
					(_nw333).ArraySet1(p1, 22)
					(_nw333).ArraySet1((_this).F33(), 23)
					_1971_v8 = _nw333
					var _1972_v9 _dafny.Map
					_ = _1972_v9
					_1972_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _1971_v8)
					_1972_v9 = (_1972_v9).Update(p1, _1971_v8)
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		_1963_v3 = (_this).F33()
		(globalState).F15 = ((_this).F32()).Plus(_1964_v4)
		var _hi10 _dafny.Int = (_this).F32()
		_ = _hi10
		for _1973_i1 := _dafny.IntOfInt64(395); _1973_i1.Cmp(_hi10) < 0; _1973_i1 = _1973_i1.Plus(_dafny.One) {
			(globalState).F18 = (_dafny.Zero).Minus(_this.F34)
			var _1974_v10 _dafny.CodePoint
			_ = _1974_v10
			_1974_v10 = _dafny.CodePoint('t')
			var _1975_v11 _dafny.Set
			_ = _1975_v11
			_1975_v11 = _dafny.SetOf(_1974_v10)
			if (_1975_v11).Contains(_1974_v10) {
				var _1976_v12 _dafny.Map
				_ = _1976_v12
				_1976_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, _1963_v3)
				var _1977_v13 D1
				_ = _1977_v13
				_1977_v13 = Companion_D1_.Create_DC1_(_1976_v12)
				var _1978_v14 _dafny.Map
				_ = _1978_v14
				_1978_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1977_v13, p0)
				_1978_v14 = (_1978_v14).Update(_1977_v13, !(_1963_v3))
				(globalState).F0 = _1973_i1
				var _1979_v15 _dafny.Set
				_ = _1979_v15
				var _1980_v16 _dafny.Int
				_ = _1980_v16
				var _1981_v17 _dafny.Array
				_ = _1981_v17
				var _out49 _dafny.Set
				_ = _out49
				var _out50 _dafny.Int
				_ = _out50
				var _out51 _dafny.Array
				_ = _out51
				_out49, _out50, _out51 = (_this).M4(Companion_Default___.Fm0((_this).F27(), globalState), _this.F34, p1, p1, globalState)
				_1979_v15 = _out49
				_1980_v16 = _out50
				_1981_v17 = _out51
				_1962_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}(func(_1982_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})), (func() _dafny.Sequence {
					if p1 {
						return _1962_v2
					}
					return _1962_v2
				})())
				var _1983_v18 *C2
				_ = _1983_v18
				var _nw334 *C2 = New_C2_()
				_ = _nw334
				_nw334.Ctor__((_this).F31(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1960_v0).Cardinality())))
				_1983_v18 = _nw334
			} else {
				var _1984_v19 _dafny.Array
				_ = _1984_v19
				var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw335
				_1984_v19 = _nw335
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))
				_ = _index246
				(_1984_v19).ArraySet1((_this).F30(), (_index246).Int())
				var _1985_v20 _dafny.Map
				_ = _1985_v20
				_1985_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))
				_ = _index247
				var _rhs358 _dafny.Sequence = Companion_Default___.Fm14(((_this).F32()).Cmp((_this).F27()) <= 0, (func() bool {
					if (_1985_v20).Contains((_this).F33()) {
						return (_1985_v20).Get((_this).F33()).(bool)
					}
					return p1
				})(), p1, (_this).F27(), globalState)
				_ = _rhs358
				var _rhs359 _dafny.Int = p3
				_ = _rhs359
				var _rhs360 _dafny.Int = (_this).F32()
				_ = _rhs360
				var _rhs361 bool = p0
				_ = _rhs361
				var _rhs362 _dafny.Sequence = (_this).F30()
				_ = _rhs362
				var _lhs280 *GlobalState = globalState
				_ = _lhs280
				var _lhs281 *GlobalState = globalState
				_ = _lhs281
				var _lhs282 _dafny.Array = _1984_v19
				_ = _lhs282
				var _lhs283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))
				_ = _lhs283
				_1962_v2 = _rhs358
				_lhs280.F21 = _rhs359
				_lhs281.F0 = _rhs360
				_1963_v3 = _rhs361
				(_lhs282).ArraySet1(_rhs362, (_lhs283).Int())
				var _1986_v21 _dafny.Array
				_ = _1986_v21
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_40
				var _nw336 _dafny.Array
				_ = _nw336
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw336 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Set = (func(_1987_p2 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1988_i3 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_1987_p2)
						}
					})(p2)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw336 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw336).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw336).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1986_v21 = _nw336
				var _1989_v22 _dafny.Array
				_ = _1989_v22
				var _nw337 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
				_ = _nw337
				_1989_v22 = _nw337
				var _1990_v23 _dafny.Map
				_ = _1990_v23
				_1990_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1986_v21, _1989_v22)
				_1990_v23 = (_1990_v23).Update(_1986_v21, _1989_v22)
				var _1991_v24 _dafny.Map
				_ = _1991_v24
				_1991_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F27(), globalState), (_this).F28())
				var _1992_v25 *C1
				_ = _1992_v25
				var _nw338 *C1 = New_C1_()
				_ = _nw338
				_nw338.Ctor__(p1, (_dafny.MultiSetFromSeq(Companion_Default___.Fm13(p2, (_this).F32(), p0, _dafny.IntOfInt64(421), globalState))).Union(_dafny.MultiSetOf(p1)), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_1984_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_1984_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), p3), (Companion_Default___.SafeIndex(Companion_Default___.Fm1(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1984_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((_1984_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1984_v19), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), p3)).Cardinality()))).Uint32(), Companion_Default___.Fm1(globalState)), p3, (func() _dafny.Array {
					if (_1991_v24).Contains((_this).F33()) {
						return (_1991_v24).Get((_this).F33()).(_dafny.Array)
					}
					return (_this).F28()
				})())
				_1992_v25 = _nw338
				_1992_v25 = _1992_v25
				var _1993_v26 _dafny.Array
				_ = _1993_v26
				var _nwElement0_78 bool = p1
				_ = _nwElement0_78
				var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(2))
				_ = _nw339
				(_nw339).ArraySet1(_nwElement0_78, 0)
				(_nw339).ArraySet1(false, 1)
				_1993_v26 = _nw339
				var _1994_v27 _dafny.MultiSet
				_ = _1994_v27
				_1994_v27 = _dafny.MultiSetOf(_1993_v26)
				var _1995_v28 bool
				_ = _1995_v28
				var _1996_v29 _dafny.Array
				_ = _1996_v29
				var _1997_v30 _dafny.Map
				_ = _1997_v30
				var _out52 bool
				_ = _out52
				var _out53 _dafny.Array
				_ = _out53
				var _out54 _dafny.Map
				_ = _out54
				_out52, _out53, _out54 = (_1992_v25).M5((_1994_v27).IsDisjointFrom(_dafny.MultiSetOf(_1993_v26)), p1, globalState)
				_1995_v28 = _out52
				_1996_v29 = _out53
				_1997_v30 = _out54
				var _1998_v32 _dafny.Set
				_ = _1998_v32
				_1998_v32 = _dafny.SetOf(_1973_i1)
				var _1999_v33 _dafny.Map
				_ = _1999_v33
				_1999_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1973_i1, (_1998_v32).Cardinality())
				var _2000_v34 _dafny.MultiSet
				_ = _2000_v34
				_2000_v34 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1973_i1, ((_this).F31()).Cardinality()), _1999_v33)
				var _2001_v35 D20
				_ = _2001_v35
				_2001_v35 = Companion_D20_.Create_DC50_(_2000_v34)
				var _2002_v36 T0
				_ = _2002_v36
				var _nw340 *C5 = New_C5_()
				_ = _nw340
				_nw340.Ctor__(func() _dafny.Map {
					var _coll60 = _dafny.NewMapBuilder()
					_ = _coll60
					for _iter69 := _dafny.Iterate(((_2001_v35).Dtor_cf92()).Elements()); ; {
						_compr_60, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _2003_v31 _dafny.Map
						_2003_v31 = interface{}(_compr_60).(_dafny.Map)
						if ((_2001_v35).Dtor_cf92()).Contains(_2003_v31) {
							_coll60.Add(_2003_v31, _dafny.IntOfInt64(4))
						}
					}
					return _coll60.ToMap()
				}(), _dafny.MultiSetFromSeq(_dafny.SeqOf(p1, false, _1992_v25.F35)), (_this).F30(), _this.F34, (_this).F28())
				_2002_v36 = _nw340
				var _2004_v37 _dafny.Map
				_ = _2004_v37
				_2004_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2002_v36, p3)
				var _2005_v38 D10
				_ = _2005_v38
				_2005_v38 = Companion_D10_.Create_DC22_((func() _dafny.Int {
					if (_2004_v37).Contains(_2002_v36) {
						return (_2004_v37).Get(_2002_v36).(_dafny.Int)
					}
					return _dafny.IntOfInt64(596)
				})(), (_this).F31())
				var _pat_let_tv44 = _1960_v0
				_ = _pat_let_tv44
				var _pat_let_tv45 = _1960_v0
				_ = _pat_let_tv45
				_2005_v38 = func(_pat_let60_0 D10) D10 {
					return func(_2006_dt__update__tmp_h1 D10) D10 {
						return func(_pat_let61_0 _dafny.MultiSet) D10 {
							return func(_2007_dt__update_hcf41_h0 _dafny.MultiSet) D10 {
								return Companion_D10_.Create_DC22_((_2006_dt__update__tmp_h1).Dtor_cf40(), _2007_dt__update_hcf41_h0)
							}(_pat_let61_0)
						}(((_this).F31()).Intersection(_dafny.MultiSetOf((_pat_let_tv44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-773), _dafny.IntOfUint32((_pat_let_tv45).Cardinality()))).Uint32()).(bool))))
					}(_pat_let60_0)
				}(_2005_v38)
			}
			(globalState).F26 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1960_v0, _1960_v0), _dafny.Companion_Sequence_.Concatenate(_1960_v0, _dafny.SeqOf(!(p0))))
			var _2008_v39 _dafny.Set
			_ = _2008_v39
			_2008_v39 = _dafny.SetOf((_this).F30(), (_this).F30())
			(globalState).F26 = !((func() _dafny.Set {
				if p1 {
					return _2008_v39
				}
				return _2008_v39
			})()).Contains((_this).F30())
		}
		var _2009_v40 _dafny.Map
		_ = _2009_v40
		_2009_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(550), (_dafny.MultiSetFromSeq((_this).F30())).Cardinality())
		_1963_v3 = !(((_2009_v40).Cardinality()).Cmp((_this).F32()) <= 0) || ((_this).F33())
		var _2010_v41 _dafny.Array
		_ = _2010_v41
		var _nw341 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw341
		_2010_v41 = _nw341
		r0 = (func() _dafny.Array {
			if ((_this).F32()).Cmp(_this.F34) > 0 {
				return _2010_v41
			}
			return _2010_v41
		})()
		r1 = ((_dafny.MultiSetOf((_this).Fm3((_this).F31(), globalState), _1963_v3, (_this).F33())).Cardinality()).Plus(((_this).F31()).Cardinality())
		return r0, r1
	}
}
func (_this *C6) M3(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) {
	{
		var _2011_v0 _dafny.Array
		_ = _2011_v0
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_41
		var _nw342 _dafny.Array
		_ = _nw342
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw342 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) bool = func(_2012_i0 _dafny.Int) bool {
				return (_this).F33()
			}
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw342 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw342).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw342).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		_2011_v0 = _nw342
		_2011_v0 = _2011_v0
		var _2013_v1 _dafny.Sequence
		_ = _2013_v1
		_2013_v1 = _dafny.SeqOf(_2011_v0, _2011_v0, _2011_v0)
		var _2014_v2 _dafny.Map
		_ = _2014_v2
		_2014_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() bool {
			if p0 {
				return p3
			}
			return p0
		})())
		var _2015_v3 _dafny.Map
		_ = _2015_v3
		_2015_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F33())
		var _2016_v4 _dafny.MultiSet
		_ = _2016_v4
		_2016_v4 = _dafny.MultiSetOf((_this).F27())
		var _rhs363 _dafny.Sequence = _2013_v1
		_ = _rhs363
		var _rhs364 bool = (_this).Fm3((_this).F31(), globalState)
		_ = _rhs364
		var _rhs365 _dafny.Map = (func() _dafny.Map {
			if (_this).F33() {
				return (_2014_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yskuy"), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yskuy")).Cardinality()))).Uint32(), p1), p3))
			}
			return (_2014_v2).Merge(_2014_v2)
		})()
		_ = _rhs365
		var _rhs366 bool = !(!((_this).F33())) || ((func() bool {
			if (_2015_v3).Contains(p0) {
				return (_2015_v3).Get(p0).(bool)
			}
			return p3
		})())
		_ = _rhs366
		var _rhs367 bool = (_2016_v4).IsSubsetOf(_2016_v4)
		_ = _rhs367
		var _lhs284 *GlobalState = globalState
		_ = _lhs284
		var _lhs285 *GlobalState = globalState
		_ = _lhs285
		var _lhs286 *GlobalState = globalState
		_ = _lhs286
		_2013_v1 = _rhs363
		_lhs284.F26 = _rhs364
		_2014_v2 = _rhs365
		_lhs285.F26 = _rhs366
		_lhs286.F26 = _rhs367
		var _2017_i1 _dafny.Int
		_ = _2017_i1
		_2017_i1 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("msled"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg110 _dafny.Int) interface{} {
					return coer110(arg110)
				}
			}((func(_2027_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2028_i2 _dafny.Int) _dafny.CodePoint {
					return _2027_p1
				}
			})(p1)))) {
				{
					if (_2017_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_2017_i1 = (_2017_i1).Plus(_dafny.One)
					(globalState).F0 = (func() _dafny.Set {
						var _coll61 = _dafny.NewBuilder()
						_ = _coll61
						for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-80), _dafny.IntOfInt64(763))); ; {
							_compr_61, _ok70 := _iter70()
							if !_ok70 {
								break
							}
							var _2018_v5 _dafny.Int
							_2018_v5 = interface{}(_compr_61).(_dafny.Int)
							if ((_dafny.IntOfInt64(-80)).Cmp(_2018_v5) <= 0) && ((_2018_v5).Cmp(_dafny.IntOfInt64(763)) < 0) {
								_coll61.Add(Companion_Default___.SafeModuloInt(_2018_v5, (_this).F32()))
							}
						}
						return _coll61.ToSet()
					}()).Cardinality()
					var _2019_v6 _dafny.Map
					_ = _2019_v6
					_2019_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, (p2).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.CodePoint)))
					var _2020_v7 _dafny.Map
					_ = _2020_v7
					_2020_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, _dafny.CodePoint('j'))
					var _2021_v8 *C2
					_ = _2021_v8
					var _nw343 *C2 = New_C2_()
					_ = _nw343
					_nw343.Ctor__((_this).F31(), ((func() _dafny.Map {
						if (_2019_v6).Contains(p1) {
							return (_2019_v6).Get(p1).(_dafny.Map)
						}
						return _2020_v7
					})()).Cardinality())
					_2021_v8 = _nw343
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_2011_v0), 0))
					_ = _index248
					(_2011_v0).ArraySet1(p3, (_index248).Int())
					var _2022_v9 T1
					_ = _2022_v9
					var _nw344 *C1 = New_C1_()
					_ = _nw344
					_nw344.Ctor__(p3, ((_this).F31()).Update(p3, Companion_Default___.Abs(Companion_Default___.Fm1(globalState))), _dafny.SeqOf(_this.F34), _this.F34, (_this).F28())
					_2022_v9 = _nw344
					var _2023_v10 _dafny.Map
					_ = _2023_v10
					_2023_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2022_v9, _this.F34)
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_2011_v0), 0))
					_ = _index249
					(_2011_v0).ArraySet1(((_this.F34).Cmp((_2023_v10).Cardinality()) < 0) || (p0), (_index249).Int())
					var _2024_v11 _dafny.Set
					_ = _2024_v11
					_2024_v11 = _dafny.SetOf(p2, p2)
					var _2025_v12 _dafny.Sequence
					_ = _2025_v12
					_2025_v12 = _dafny.SeqOf(false)
					var _2026_v13 _dafny.Map
					_ = _2026_v13
					_2026_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2022_v9).Fm4(_2024_v11, (_this).F33(), _2025_v12, globalState), p0)
					_2026_v13 = (_2026_v13).Update((_this).F32(), p3)
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _2029_v14 _dafny.Set
		_ = _2029_v14
		_2029_v14 = _dafny.SetOf(_dafny.SeqOf(_this.F34, _this.F34), (_this).F30(), (_this).F30(), (_this).F30(), (_this).F30())
		(globalState).F26 = (_2029_v14).IsProperSubsetOf(Companion_Default___.Fm50(globalState))
		var _hi11 _dafny.Int = _this.F34
		_ = _hi11
		for _2030_i3 := (_dafny.Zero).Minus((_this).F32()); _2030_i3.Cmp(_hi11) < 0; _2030_i3 = _2030_i3.Plus(_dafny.One) {
			var _2031_v15 _dafny.Map
			_ = _2031_v15
			_2031_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			var _2032_v16 _dafny.Map
			_ = _2032_v16
			_2032_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2031_v15).Contains(p0) {
					return (_2031_v15).Get(p0).(_dafny.Sequence)
				}
				return p2
			})()).Cardinality()), _2030_i3)
			var _2033_v18 _dafny.Array
			_ = _2033_v18
			var _nwElement0_79 _dafny.Map = _2032_v16
			_ = _nwElement0_79
			var _nw345 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(14))
			_ = _nw345
			(_nw345).ArraySet1(_nwElement0_79, 0)
			(_nw345).ArraySet1(Companion_Default___.Fm15((_this).F27(), _2030_i3, globalState), 1)
			(_nw345).ArraySet1(func() _dafny.Map {
				var _coll62 = _dafny.NewMapBuilder()
				_ = _coll62
				for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-310), _dafny.IntOfInt64(360))); ; {
					_compr_62, _ok71 := _iter71()
					if !_ok71 {
						break
					}
					var _2034_v17 _dafny.Int
					_2034_v17 = interface{}(_compr_62).(_dafny.Int)
					if ((_dafny.IntOfInt64(-310)).Cmp(_2034_v17) <= 0) && ((_2034_v17).Cmp(_dafny.IntOfInt64(360)) < 0) {
						_coll62.Add(Companion_Default___.SafeModuloInt(_2034_v17, (_this).F27()), _dafny.IntOfInt64(-274))
					}
				}
				return _coll62.ToMap()
			}(), 2)
			(_nw345).ArraySet1(_2032_v16, 3)
			(_nw345).ArraySet1(Companion_Default___.Fm30(globalState), 4)
			(_nw345).ArraySet1(_2032_v16, 5)
			(_nw345).ArraySet1(_2032_v16, 6)
			(_nw345).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2030_i3, _2030_i3), 7)
			(_nw345).ArraySet1(_2032_v16, 8)
			(_nw345).ArraySet1(_2032_v16, 9)
			(_nw345).ArraySet1(_2032_v16, 10)
			(_nw345).ArraySet1(_2032_v16, 11)
			(_nw345).ArraySet1(_2032_v16, 12)
			(_nw345).ArraySet1(Companion_Default___.Fm42((_this).F32(), p3, _2030_i3, globalState), 13)
			_2033_v18 = _nw345
			var _2035_v19 _dafny.MultiSet
			_ = _2035_v19
			_2035_v19 = _dafny.MultiSetOf(_2033_v18, _2033_v18, _2033_v18)
			(globalState).F0 = (func() _dafny.Int {
				if (_2035_v19).Contains(_2033_v18) {
					return (_2035_v19).Multiplicity(_2033_v18)
				}
				return (_this.F34).Minus((_this).F32())
			})()
			var _2036_v20 _dafny.Sequence
			_ = _2036_v20
			_2036_v20 = _dafny.SeqOf(Companion_D18_.Create_DC47_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex(_2030_i3, _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), _2030_i3)).Cardinality()), p1))
			(globalState).F3 = _dafny.IntOfUint32((Companion_Default___.Fm51(_dafny.Companion_Sequence_.Concatenate(_2036_v20, _dafny.SeqOf(Companion_D18_.Create_DC47_((_this).F27(), p1))), !(_2015_v3).Equals(_2015_v3), (_2031_v15).Merge(_2031_v15), globalState)).Cardinality())
			(globalState).F0 = (func() _dafny.Int {
				if p0 {
					return Companion_Default___.SafeDivisionInt(_2030_i3, _this.F34)
				}
				return _dafny.IntOfInt64(861)
			})()
			var _2037_v21 _dafny.Array
			_ = _2037_v21
			var _nw346 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw346
			_2037_v21 = _nw346
			var _2038_v22 _dafny.Sequence
			_ = _2038_v22
			_2038_v22 = _dafny.SeqOf((_this).F33(), p3, p0)
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_2037_v21), 0))
			_ = _index250
			(_2037_v21).ArraySet1((_2030_i3).Plus((_this).Fm4(_dafny.SetOf(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg111 _dafny.Int) interface{} {
					return coer111(arg111)
				}
			}((func(_2039_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2040_i4 _dafny.Int) _dafny.CodePoint {
					return _2039_p1
				}
			})(p1)))), (_this).F33(), _2038_v22, globalState)), (_index250).Int())
			var _2041_v23 _dafny.Sequence
			_ = _2041_v23
			_2041_v23 = _dafny.UnicodeSeqOfUtf8Bytes("jbd")
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_2037_v21), 0))
			_ = _index251
			var _rhs368 _dafny.Int = (_this).F32()
			_ = _rhs368
			var _rhs369 _dafny.Array = _2033_v18
			_ = _rhs369
			var _rhs370 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2041_v23, _2041_v23)
			_ = _rhs370
			var _lhs287 _dafny.Array = _2037_v21
			_ = _lhs287
			var _lhs288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_2037_v21), 0))
			_ = _lhs288
			(_lhs287).ArraySet1(_rhs368, (_lhs288).Int())
			_2033_v18 = _rhs369
			_2041_v23 = _rhs370
		}
		var _2042_v24 _dafny.Map
		_ = _2042_v24
		_2042_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(76), p3)
		var _2043_v26 _dafny.Map
		_ = _2043_v26
		_2043_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F27())
		var _2044_v28 D1
		_ = _2044_v28
		_2044_v28 = Companion_D1_.Create_DC1_(_2042_v24)
		var _2045_v29 _dafny.Sequence
		_ = _2045_v29
		_2045_v29 = _dafny.SeqOf((_this).F33(), p0)
		var _2046_v30 _dafny.Sequence
		_ = _2046_v30
		_2046_v30 = _dafny.SeqOf(_2042_v24)
		var _2047_v32 _dafny.Set
		_ = _2047_v32
		_2047_v32 = _dafny.SetOf(((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus((_2015_v3).Cardinality()), _this.F34, (_this).F32(), _this.F34)
		var _2048_v34 _dafny.Array
		_ = _2048_v34
		var _nwElement0_80 _dafny.Map = _2042_v24
		_ = _nwElement0_80
		var _nw347 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(19))
		_ = _nw347
		(_nw347).ArraySet1(_nwElement0_80, 0)
		(_nw347).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F33()), 1)
		(_nw347).ArraySet1(_2042_v24, 2)
		(_nw347).ArraySet1((func() _dafny.Map {
			var _coll63 = _dafny.NewMapBuilder()
			_ = _coll63
			for _iter72 := _dafny.Iterate((_2043_v26).Keys().Elements()); ; {
				_compr_63, _ok72 := _iter72()
				if !_ok72 {
					break
				}
				var _2049_v25 _dafny.Int
				_2049_v25 = interface{}(_compr_63).(_dafny.Int)
				if (_2043_v26).Contains(_2049_v25) {
					_coll63.Add((_2049_v25).Minus(_dafny.IntOfUint32(((_this).F30()).Cardinality())), (_this).F33())
				}
			}
			return _coll63.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll64 = _dafny.NewMapBuilder()
			_ = _coll64
			for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(432), _dafny.IntOfInt64(-747))); ; {
				_compr_64, _ok73 := _iter73()
				if !_ok73 {
					break
				}
				var _2050_v27 _dafny.Int
				_2050_v27 = interface{}(_compr_64).(_dafny.Int)
				if ((_dafny.IntOfInt64(432)).Cmp(_2050_v27) <= 0) && ((_2050_v27).Cmp(_dafny.IntOfInt64(-747)) < 0) {
					_coll64.Add(Companion_Default___.SafeDivisionInt(_2050_v27, (_this).F32()), p3)
				}
			}
			return _coll64.ToMap()
		}()), 3)
		(_nw347).ArraySet1(_2042_v24, 4)
		(_nw347).ArraySet1(((_2044_v28).Dtor_cf1()).Update((_2043_v26).Cardinality(), p3), 5)
		(_nw347).ArraySet1(_2042_v24, 6)
		(_nw347).ArraySet1((_2042_v24).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2045_v29).Cardinality()), p3)).Cardinality(), p3)), 7)
		(_nw347).ArraySet1((_2046_v30).Select((Companion_Default___.SafeIndex((_2015_v3).Cardinality(), _dafny.IntOfUint32((_2046_v30).Cardinality()))).Uint32()).(_dafny.Map), 8)
		(_nw347).ArraySet1(_2042_v24, 9)
		(_nw347).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, false), 10)
		(_nw347).ArraySet1(Companion_Default___.Fm48(p3, p3, _this.F34, globalState), 11)
		(_nw347).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, !(p3)), 12)
		(_nw347).ArraySet1(func() _dafny.Map {
			var _coll65 = _dafny.NewMapBuilder()
			_ = _coll65
			for _iter74 := _dafny.Iterate((_2047_v32).Elements()); ; {
				_compr_65, _ok74 := _iter74()
				if !_ok74 {
					break
				}
				var _2051_v31 _dafny.Int
				_2051_v31 = interface{}(_compr_65).(_dafny.Int)
				if (_2047_v32).Contains(_2051_v31) {
					_coll65.Add(Companion_Default___.SafeDivisionInt(_2051_v31, (Companion_D13_.Create_DC32_((_this).F32(), (_this).F32(), (_this).F30(), p3)).Dtor_cf57()), (_this).F33())
				}
			}
			return _coll65.ToMap()
		}(), 13)
		(_nw347).ArraySet1((_2042_v24).Merge(Companion_Default___.Fm48(p0, p3, (_this).F27(), globalState)), 14)
		(_nw347).ArraySet1(_2042_v24, 15)
		(_nw347).ArraySet1((_2042_v24).Merge(func() _dafny.Map {
			var _coll66 = _dafny.NewMapBuilder()
			_ = _coll66
			for _iter75 := _dafny.Iterate(((_this).F30()).Elements()); ; {
				_compr_66, _ok75 := _iter75()
				if !_ok75 {
					break
				}
				var _2052_v33 _dafny.Int
				_2052_v33 = interface{}(_compr_66).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains((_this).F30(), _2052_v33) {
					_coll66.Add((_2052_v33).Plus((_2015_v3).Cardinality()), (_this).F33())
				}
			}
			return _coll66.ToMap()
		}()), 16)
		(_nw347).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), p3)).Merge(_2042_v24), 17)
		(_nw347).ArraySet1(_2042_v24, 18)
		_2048_v34 = _nw347
		_2048_v34 = _2048_v34
	}
}
func (_this *C6) M4(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _2053_v0 _dafny.Map
		_ = _2053_v0
		_2053_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F32())
		var _2054_v1 _dafny.Sequence
		_ = _2054_v1
		_2054_v1 = _dafny.SeqOf((_this).F33(), p3)
		var _2055_v2 _dafny.Map
		_ = _2055_v2
		_2055_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2053_v0, _dafny.IntOfUint32((_2054_v1).Cardinality()))
		var _2056_v3 _dafny.Sequence
		_ = _2056_v3
		_2056_v3 = _dafny.SeqOf((_this).F30(), (_this).F30(), (_this).F30())
		var _2057_v4 _dafny.Sequence
		_ = _2057_v4
		_2057_v4 = _dafny.UnicodeSeqOfUtf8Bytes("pfggacs")
		var _2058_v5 _dafny.CodePoint
		_ = _2058_v5
		_2058_v5 = _dafny.CodePoint('p')
		var _2059_v6 *C5
		_ = _2059_v6
		var _nw348 *C5 = New_C5_()
		_ = _nw348
		_nw348.Ctor__(_2055_v2, (_this).F29(), (_2056_v3).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_2056_v3).Cardinality()))).Uint32()).(_dafny.Sequence), (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2057_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32(), _2058_v5)).Cardinality())).Minus((_this).F27()), (_this).F28())
		_2059_v6 = _nw348
		var _2060_v7 *C3
		_ = _2060_v7
		var _nw349 *C3 = New_C3_()
		_ = _nw349
		_nw349.Ctor__(_dafny.IntOfInt64(90), (_this).F28())
		_2060_v7 = _nw349
		var _2061_v9 _dafny.Set
		_ = _2061_v9
		_2061_v9 = _dafny.SetOf(_this.F34, _this.F34)
		var _2062_v10 D11
		_ = _2062_v10
		_2062_v10 = Companion_D11_.Create_DC23_(_dafny.MultiSetOf(func() _dafny.Set {
			var _coll67 = _dafny.NewBuilder()
			_ = _coll67
			for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-381), _dafny.IntOfInt64(691))); ; {
				_compr_67, _ok76 := _iter76()
				if !_ok76 {
					break
				}
				var _2063_v8 _dafny.Int
				_2063_v8 = interface{}(_compr_67).(_dafny.Int)
				if ((_dafny.IntOfInt64(-381)).Cmp(_2063_v8) <= 0) && ((_2063_v8).Cmp(_dafny.IntOfInt64(691)) < 0) {
					_coll67.Add(Companion_Default___.SafeModuloInt(_2063_v8, p1))
				}
			}
			return _coll67.ToSet()
		}(), _2061_v9, _2061_v9))
		var _source29 D11 = (func() D11 {
			if p0 {
				return _2062_v10
			}
			return _2062_v10
		})()
		_ = _source29
		if _source29.Is_DC24() {
			var _2064___mcc_h0 _dafny.Int = _source29.Get_().(D11_DC24).Cf43
			_ = _2064___mcc_h0
			var _2065___mcc_h1 D4 = _source29.Get_().(D11_DC24).Cf44
			_ = _2065___mcc_h1
			var _2066_cf44 D4 = _2065___mcc_h1
			_ = _2066_cf44
			var _2067_cf43 _dafny.Int = _2064___mcc_h0
			_ = _2067_cf43
			var _2068_v11 _dafny.Array
			_ = _2068_v11
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_42
			var _nw350 _dafny.Array
			_ = _nw350
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw350 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Int = func(_2069_i0 _dafny.Int) _dafny.Int {
					return (_2069_i0).Times((_this).F32())
				}
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw350 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw350).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw350).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_2068_v11 = _nw350
			_2068_v11 = _2068_v11
			var _2070_v12 _dafny.Sequence
			_ = _2070_v12
			_2070_v12 = _dafny.SeqOf(_2054_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F33()), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F33())).Cardinality()))).Uint32(), p0))
			var _2071_v13 D6
			_ = _2071_v13
			_2071_v13 = Companion_D6_.Create_DC12_(false, _2070_v12, Companion_Default___.Fm31(p3, p2, globalState))
			_2071_v13 = _2071_v13
			var _2072_v14 _dafny.Array
			_ = _2072_v14
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_43
			var _nw351 _dafny.Array
			_ = _nw351
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw351 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Map = (func(_2073_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_2074_i1 _dafny.Int) _dafny.Map {
						return _2073_v0
					}
				})(_2053_v0)
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw351 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw351).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw351).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_2072_v14 = _nw351
			var _source30 D16 = Companion_D16_.Create_DC39_(_2072_v14)
			_ = _source30
			if _source30.Is_DC40() {
				var _2075___mcc_h4 bool = _source30.Get_().(D16_DC40).Cf74
				_ = _2075___mcc_h4
				var _2076___mcc_h5 _dafny.Int = _source30.Get_().(D16_DC40).Cf75
				_ = _2076___mcc_h5
				var _2077___mcc_h6 _dafny.Int = _source30.Get_().(D16_DC40).Cf76
				_ = _2077___mcc_h6
				var _2078_cf76 _dafny.Int = _2077___mcc_h6
				_ = _2078_cf76
				var _2079_cf75 _dafny.Int = _2076___mcc_h5
				_ = _2079_cf75
				var _2080_cf74 bool = _2075___mcc_h4
				_ = _2080_cf74
				var _2081_v15 _dafny.Sequence
				_ = _2081_v15
				_2081_v15 = _dafny.SeqOf(_2057_v4)
				_2057_v4 = (_2081_v15).Select((Companion_Default___.SafeIndex(_2079_cf75, _dafny.IntOfUint32((_2081_v15).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _2082_v16 _dafny.Sequence
				_ = _2082_v16
				_2082_v16 = _dafny.SeqOf(_2067_cf43)
				var _2083_v17 _dafny.Sequence
				_ = _2083_v17
				_2083_v17 = _dafny.SeqOf(_2053_v0)
				var _rhs371 _dafny.Sequence = (_this).F30()
				_ = _rhs371
				var _rhs372 bool = (func() bool {
					if (_this).F33() {
						return p0
					}
					return Companion_Default___.Fm0(_2078_cf76, globalState)
				})()
				_ = _rhs372
				var _rhs373 _dafny.Sequence = _2083_v17
				_ = _rhs373
				var _lhs289 *GlobalState = globalState
				_ = _lhs289
				_2082_v16 = _rhs371
				_lhs289.F26 = _rhs372
				_2083_v17 = _rhs373
				_2080_cf74 = !(((_this).F27()).Cmp((_2079_cf75).Times(_this.F34)) <= 0)
				_2068_v11 = _2068_v11
			} else if _source30.Is_DC41() {
				var _2084___mcc_h7 _dafny.Map = _source30.Get_().(D16_DC41).Cf77
				_ = _2084___mcc_h7
				var _2085_cf77 _dafny.Map = _2084___mcc_h7
				_ = _2085_cf77
				var _2086_v19 _dafny.MultiSet
				_ = _2086_v19
				_2086_v19 = _dafny.MultiSetOf(_2061_v9, func() _dafny.Set {
					var _coll68 = _dafny.NewBuilder()
					_ = _coll68
					for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(799), _dafny.IntOfInt64(419))); ; {
						_compr_68, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _2087_v18 _dafny.Int
						_2087_v18 = interface{}(_compr_68).(_dafny.Int)
						if ((_dafny.IntOfInt64(799)).Cmp(_2087_v18) <= 0) && ((_2087_v18).Cmp(_dafny.IntOfInt64(419)) < 0) {
							_coll68.Add(Companion_Default___.SafeModuloInt(_2087_v18, _2067_cf43))
						}
					}
					return _coll68.ToSet()
				}())
				var _2088_v20 D12
				_ = _2088_v20
				_2088_v20 = Companion_D12_.Create_DC28_(((_this).F32()).Times((_dafny.Zero).Minus((_2086_v19).Cardinality())), ((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-352), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int), (_this).F33())
				_2088_v20 = _2088_v20
				var _2089_v21 _dafny.Array
				_ = _2089_v21
				var _nw352 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw352
				_2089_v21 = _nw352
				var _2090_v22 T1
				_ = _2090_v22
				var _nw353 *C1 = New_C1_()
				_ = _nw353
				_nw353.Ctor__(p0, _dafny.MultiSetOf(p2), _dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), _dafny.IntOfUint32((_2054_v1).Cardinality())), (_dafny.Zero).Minus((_dafny.Zero).Minus(_2067_cf43)), (_this).F28())
				_2090_v22 = _nw353
				var _2091_v23 _dafny.Array
				_ = _2091_v23
				var _nwElement0_81 T1 = _2090_v22
				_ = _nwElement0_81
				var _nw354 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(18))
				_ = _nw354
				(_nw354).ArraySet1(_nwElement0_81, 0)
				(_nw354).ArraySet1(_2090_v22, 1)
				(_nw354).ArraySet1(_2090_v22, 2)
				(_nw354).ArraySet1(_2090_v22, 3)
				(_nw354).ArraySet1(_2090_v22, 4)
				(_nw354).ArraySet1(_2090_v22, 5)
				(_nw354).ArraySet1(_2090_v22, 6)
				(_nw354).ArraySet1(_2090_v22, 7)
				(_nw354).ArraySet1(_2090_v22, 8)
				(_nw354).ArraySet1(_2090_v22, 9)
				(_nw354).ArraySet1(_2090_v22, 10)
				(_nw354).ArraySet1(_2090_v22, 11)
				(_nw354).ArraySet1(_2090_v22, 12)
				(_nw354).ArraySet1(_2090_v22, 13)
				(_nw354).ArraySet1(_2090_v22, 14)
				(_nw354).ArraySet1(_2090_v22, 15)
				(_nw354).ArraySet1(_2090_v22, 16)
				(_nw354).ArraySet1(_2090_v22, 17)
				_2091_v23 = _nw354
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_2089_v21), 0))
				_ = _index252
				(_2089_v21).ArraySet1(_2091_v23, (_index252).Int())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_2089_v21), 0))
				_ = _index253
				(_2089_v21).ArraySet1(_2091_v23, (_index253).Int())
				(globalState).F24 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_2092_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2093_i2 _dafny.Int) _dafny.CodePoint {
						return _2092_v5
					}
				})(_2058_v5))), _2057_v4)).Cardinality())).Plus(_dafny.IntOfUint32((_2054_v1).Cardinality()))
				(globalState).F21 = Companion_Default___.SafeModuloInt((_this).F32(), (_this).F27())
			} else if _source30.Is_DC42() {
				var _2094___mcc_h8 _dafny.Int = _source30.Get_().(D16_DC42).Cf78
				_ = _2094___mcc_h8
				var _2095___mcc_h9 _dafny.Array = _source30.Get_().(D16_DC42).Cf79
				_ = _2095___mcc_h9
				var _2096___mcc_h10 bool = _source30.Get_().(D16_DC42).Cf80
				_ = _2096___mcc_h10
				var _2097___mcc_h11 D2 = _source30.Get_().(D16_DC42).Cf81
				_ = _2097___mcc_h11
				var _2098_cf81 D2 = _2097___mcc_h11
				_ = _2098_cf81
				var _2099_cf80 bool = _2096___mcc_h10
				_ = _2099_cf80
				var _2100_cf79 _dafny.Array = _2095___mcc_h9
				_ = _2100_cf79
				var _2101_cf78 _dafny.Int = _2094___mcc_h8
				_ = _2101_cf78
				(globalState).F3 = (_this).F27()
				var _2102_v24 _dafny.Array
				_ = _2102_v24
				var _nwElement0_82 bool = _2099_cf80
				_ = _nwElement0_82
				var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(4))
				_ = _nw355
				(_nw355).ArraySet1(_nwElement0_82, 0)
				(_nw355).ArraySet1(true, 1)
				(_nw355).ArraySet1((_this).F33(), 2)
				(_nw355).ArraySet1(((_this).F32()).Cmp(p1) > 0, 3)
				_2102_v24 = _nw355
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_2102_v24), 0))
				_ = _index254
				(_2102_v24).ArraySet1(p0, (_index254).Int())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_2102_v24), 0))
				_ = _index255
				(_2102_v24).ArraySet1((_2054_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2057_v4, _2057_v4)).Cardinality()), _dafny.IntOfUint32((_2054_v1).Cardinality()))).Uint32()).(bool), (_index255).Int())
				var _2103_v25 *C2
				_ = _2103_v25
				var _nw356 *C2 = New_C2_()
				_ = _nw356
				_nw356.Ctor__((_this).F29(), (_dafny.Zero).Minus((_dafny.IntOfInt64(753)).Times(p1)))
				_2103_v25 = _nw356
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2068_v11), 0))
				_ = _index256
				(_2068_v11).ArraySet1(_2067_cf43, (_index256).Int())
				var _2104_v26 T0
				_ = _2104_v26
				var _nw357 *C3 = New_C3_()
				_ = _nw357
				_nw357.Ctor__((_this).F27(), (_this).F28())
				_2104_v26 = _nw357
				var _2105_v27 D16
				_ = _2105_v27
				_2105_v27 = Companion_D16_.Create_DC40_((_2054_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lpyjyi")).Cardinality()), _dafny.IntOfUint32((_2054_v1).Cardinality()))).Uint32()).(bool), p1, (_this).F27())
				var _2106_v28 _dafny.Map
				_ = _2106_v28
				_2106_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2104_v26, (_2105_v27).Dtor_cf76())
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2068_v11), 0))
				_ = _index257
				var _rhs374 *C2 = _2103_v25
				_ = _rhs374
				var _rhs375 _dafny.Int = _dafny.IntOfInt64(446)
				_ = _rhs375
				var _rhs376 _dafny.Map = (func() _dafny.Map {
					if (_this).F33() {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2104_v26, ((_this).F29()).Cardinality())
					}
					return (func() _dafny.Map {
						if (_this).F33() {
							return _2106_v28
						}
						return _2106_v28
					})()
				})()
				_ = _rhs376
				var _lhs290 _dafny.Array = _2068_v11
				_ = _lhs290
				var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2068_v11), 0))
				_ = _lhs291
				_2103_v25 = _rhs374
				(_lhs290).ArraySet1(_rhs375, (_lhs291).Int())
				_2106_v28 = _rhs376
				(globalState).F21 = (_2068_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2068_v11), 0))).Int()).(_dafny.Int)
			} else {
				var _2107___mcc_h12 _dafny.Array = _source30.Get_().(D16_DC39).Cf73
				_ = _2107___mcc_h12
				var _2108_cf73 _dafny.Array = _2107___mcc_h12
				_ = _2108_cf73
				var _rhs377 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(556), p1)
				_ = _rhs377
				var _rhs378 bool = (_2054_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F30()).Cardinality()), _dafny.IntOfUint32((_2054_v1).Cardinality()))).Uint32()).(bool)
				_ = _rhs378
				var _lhs292 *GlobalState = globalState
				_ = _lhs292
				var _lhs293 *GlobalState = globalState
				_ = _lhs293
				_lhs292.F15 = _rhs377
				_lhs293.F26 = _rhs378
				var _2109_v29 _dafny.Map
				_ = _2109_v29
				_2109_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.SetOf(p3)).Cardinality())
				var _2110_v30 _dafny.Map
				_ = _2110_v30
				_2110_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2109_v29)
				var _rhs379 bool = p2
				_ = _rhs379
				var _rhs380 _dafny.Sequence = (func() _dafny.Sequence {
					if !(p3) {
						return _2057_v4
					}
					return _2057_v4
				})()
				_ = _rhs380
				var _rhs381 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2057_v4, _dafny.UnicodeSeqOfUtf8Bytes("nq"))).Cardinality())).Plus((_this).F32())
				_ = _rhs381
				var _rhs382 _dafny.Int = (func() _dafny.Int {
					if p0 {
						return _this.F34
					}
					return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2057_v4).Cardinality()), _dafny.IntOfInt64(760)))
				})()
				_ = _rhs382
				var _rhs383 _dafny.Int = ((_2110_v30).Merge(_2110_v30)).Cardinality()
				_ = _rhs383
				var _lhs294 *GlobalState = globalState
				_ = _lhs294
				var _lhs295 *GlobalState = globalState
				_ = _lhs295
				var _lhs296 *GlobalState = globalState
				_ = _lhs296
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				_lhs294.F26 = _rhs379
				_2057_v4 = _rhs380
				_lhs295.F18 = _rhs381
				_lhs296.F21 = _rhs382
				_lhs297.F21 = _rhs383
				var _2111_v31 _dafny.Array
				_ = _2111_v31
				var _nwElement0_83 _dafny.Sequence = _2057_v4
				_ = _nwElement0_83
				var _nw358 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(2))
				_ = _nw358
				(_nw358).ArraySet1(_nwElement0_83, 0)
				(_nw358).ArraySet1((func() _dafny.Sequence {
					if p3 {
						return _2057_v4
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("yggtvlhp")
				})(), 1)
				_2111_v31 = _nw358
				var _rhs384 _dafny.CodePoint = (_2057_v4).Select((Companion_Default___.SafeIndex(((_this).F31()).Cardinality(), _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs384
				var _rhs385 _dafny.Array = (_this).F28()
				_ = _rhs385
				var _rhs386 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2054_v1).Cardinality()), _2067_cf43)
				_ = _rhs386
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				_2058_v5 = _rhs384
				_2111_v31 = _rhs385
				_lhs298.F24 = _rhs386
				(globalState).F0 = (_this).F27()
			}
			r0 = _2061_v9
		} else if _source29.Is_DC23() {
			var _2112___mcc_h2 _dafny.MultiSet = _source29.Get_().(D11_DC23).Cf42
			_ = _2112___mcc_h2
			var _2113_cf42 _dafny.MultiSet = _2112___mcc_h2
			_ = _2113_cf42
			(globalState).F26 = (_this).Fm3(((_this).F31()).Union(_dafny.MultiSetOf(p2)), globalState)
			(globalState).F26 = (_this).F33()
			var _2114_v32 D9
			_ = _2114_v32
			_2114_v32 = Companion_D9_.Create_DC19_(p0)
			var _2115_v33 _dafny.MultiSet
			_ = _2115_v33
			_2115_v33 = _dafny.MultiSetOf(_2114_v32)
			var _2116_v34 *C1
			_ = _2116_v34
			var _nw359 *C1 = New_C1_()
			_ = _nw359
			_nw359.Ctor__(p2, (_this).F29(), _dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_2115_v33).Contains(_2114_v32) {
					return (_2115_v33).Multiplicity(_2114_v32)
				}
				return _this.F34
			})(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), (_this).F27()), _dafny.IntOfUint32((Companion_Default___.Fm25(p2, globalState)).Cardinality()), (_this).F28())
			_2116_v34 = _nw359
			var _2117_v35 _dafny.Array
			_ = _2117_v35
			var _nw360 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw360
			_2117_v35 = _nw360
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_2117_v35), 0))
			_ = _index258
			(_2117_v35).ArraySet1(((_this).F33()) || ((_this).F33()), (_index258).Int())
			var _2118_v36 bool
			_ = _2118_v36
			_2118_v36 = _2116_v34.F35
			var _2119_v37 _dafny.Map
			_ = _2119_v37
			_2119_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2058_v5)
			var _2120_v38 _dafny.Map
			_ = _2120_v38
			_2120_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2119_v37, (_this).F33())
			var _2121_v39 _dafny.Map
			_ = _2121_v39
			_2121_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2118_v36, (func() bool {
				if (_2120_v38).Contains(_2119_v37) {
					return (_2120_v38).Get(_2119_v37).(bool)
				}
				return p3
			})())
			var _2122_v40 _dafny.Map
			_ = _2122_v40
			_2122_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2121_v39)
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_2117_v35), 0))
			_ = _index259
			(_2117_v35).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg113 _dafny.Int) interface{} {
					return coer113(arg113)
				}
			}((func(_2123_v39 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_2124_i3 _dafny.Int) _dafny.Map {
					return _2123_v39
				}
			})(_2121_v39))), (func() _dafny.Map {
				if (_2122_v40).Contains(p2) {
					return (_2122_v40).Get(p2).(_dafny.Map)
				}
				return _2121_v39
			})()), (_index259).Int())
		} else {
			var _2125___mcc_h3 D11 = _source29.Get_().(D11_DC25).Cf45
			_ = _2125___mcc_h3
			var _2126_cf45 D11 = _2125___mcc_h3
			_ = _2126_cf45
			var _2127_v41 _dafny.MultiSet
			_ = _2127_v41
			_2127_v41 = _dafny.MultiSetOf(_dafny.IntOfInt64(284))
			var _2128_v42 T0
			_ = _2128_v42
			var _nw361 *C5 = New_C5_()
			_ = _nw361
			_nw361.Ctor__((_2059_v6).F38(), (_this).F31(), (_this).F30(), (_this).F32(), (_this).F28())
			_2128_v42 = _nw361
			var _2129_v43 _dafny.Map
			_ = _2129_v43
			_2129_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2128_v42, true)
			var _2130_v44 _dafny.Array
			_ = _2130_v44
			var _nwElement0_84 bool = p3
			_ = _nwElement0_84
			var _nw362 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(20))
			_ = _nw362
			(_nw362).ArraySet1(_nwElement0_84, 0)
			(_nw362).ArraySet1((_this).F33(), 1)
			(_nw362).ArraySet1(p3, 2)
			(_nw362).ArraySet1(true, 3)
			(_nw362).ArraySet1(p3, 4)
			(_nw362).ArraySet1(p2, 5)
			(_nw362).ArraySet1(p0, 6)
			(_nw362).ArraySet1(!(p0), 7)
			(_nw362).ArraySet1(p2, 8)
			(_nw362).ArraySet1((func() bool {
				if false {
					return (_this).F33()
				}
				return !(p0)
			})(), 9)
			(_nw362).ArraySet1(!(!(p0)) || (p0), 10)
			(_nw362).ArraySet1(!(!(p0)), 11)
			(_nw362).ArraySet1(p0, 12)
			(_nw362).ArraySet1(!(p0), 13)
			(_nw362).ArraySet1((_2127_v41).Equals((_2127_v41).Update(p1, Companion_Default___.Abs((_2129_v43).Cardinality()))), 14)
			(_nw362).ArraySet1(p3, 15)
			(_nw362).ArraySet1(p2, 16)
			(_nw362).ArraySet1(p0, 17)
			(_nw362).ArraySet1(true, 18)
			(_nw362).ArraySet1(p0, 19)
			_2130_v44 = _nw362
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_2130_v44), 0))
			_ = _index260
			(_2130_v44).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_2054_v1, _2054_v1), (_index260).Int())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_2130_v44), 0))
			_ = _index261
			(_2130_v44).ArraySet1(Companion_Default___.Fm0((_this).F32(), globalState), (_index261).Int())
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_2130_v44), 0))
			_ = _index262
			(_2130_v44).ArraySet1(p0, (_index262).Int())
			var _2131_v45 *C1
			_ = _2131_v45
			var _nw363 *C1 = New_C1_()
			_ = _nw363
			_nw363.Ctor__(p0, ((_this).F31()).Intersection((_this).F29()), (_this).F30(), (_2128_v42).F27(), (_this).F28())
			_2131_v45 = _nw363
			var _2132_v46 _dafny.Set
			_ = _2132_v46
			_2132_v46 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ogu"), _dafny.Companion_Sequence_.Update(_2057_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32(), _2058_v5))
			var _rhs387 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, (_this).F27())
			_ = _rhs387
			var _rhs388 *C1 = _2131_v45
			_ = _rhs388
			var _rhs389 _dafny.Set = _2132_v46
			_ = _rhs389
			var _lhs299 *C6 = _this
			_ = _lhs299
			_lhs299.F34 = _rhs387
			_2131_v45 = _rhs388
			_2132_v46 = _rhs389
			var _2133_v47 _dafny.Array
			_ = _2133_v47
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_44
			var _nw364 _dafny.Array
			_ = _nw364
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw364 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) _dafny.Int = func(_2134_i4 _dafny.Int) _dafny.Int {
					return (_2134_i4).Plus(_this.F34)
				}
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw364 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw364).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw364).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_2133_v47 = _nw364
			var _2135_v48 _dafny.MultiSet
			_ = _2135_v48
			_2135_v48 = _dafny.MultiSetOf(_2133_v47, _2133_v47, _2133_v47, _2133_v47)
			_2135_v48 = _2135_v48
		}
		var _2136_v49 D4
		_ = _2136_v49
		_2136_v49 = Companion_D4_.Create_DC8_(_2058_v5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kybibih")).Cardinality()), p2)
		var _2137_v50 D11
		_ = _2137_v50
		_2137_v50 = Companion_D11_.Create_DC24_((_dafny.Zero).Minus((_this).F32()), (func() D4 {
			if (_this).F33() {
				return _2136_v49
			}
			return _2136_v49
		})())
		var _source31 D11 = _2137_v50
		_ = _source31
		if _source31.Is_DC24() {
			var _2138___mcc_h13 _dafny.Int = _source31.Get_().(D11_DC24).Cf43
			_ = _2138___mcc_h13
			var _2139___mcc_h14 D4 = _source31.Get_().(D11_DC24).Cf44
			_ = _2139___mcc_h14
			var _2140_cf44 D4 = _2139___mcc_h14
			_ = _2140_cf44
			var _2141_cf43 _dafny.Int = _2138___mcc_h13
			_ = _2141_cf43
			var _2142_v51 _dafny.Map
			_ = _2142_v51
			_2142_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !_dafny.Companion_Sequence_.Equal((_this).F30(), _dafny.SeqOf(_this.F34, _2141_cf43)))
			_2142_v51 = (_2142_v51).Update(p2, (_this).F33())
			if p3 {
				var _2143_v52 _dafny.Array
				_ = _2143_v52
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_45
				var _nw365 _dafny.Array
				_ = _nw365
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw365 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.Int = func(_2144_i5 _dafny.Int) _dafny.Int {
						return (_2144_i5).Minus((_this).F32())
					}
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw365 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw365).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw365).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_2143_v52 = _nw365
				var _2145_v53 T2
				_ = _2145_v53
				var _nw366 *C2 = New_C2_()
				_ = _nw366
				_nw366.Ctor__((_this).F29(), p1)
				_2145_v53 = _nw366
				var _2146_v54 _dafny.Sequence
				_ = _2146_v54
				_2146_v54 = _dafny.SeqOf(_2145_v53)
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))
				_ = _index263
				(_2143_v52).ArraySet1((_dafny.IntOfUint32((_2146_v54).Cardinality())).Times((_this).F27()), (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))
				_ = _index264
				(_2143_v52).ArraySet1(((_this).F32()).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2057_v4, _2057_v4)).Cardinality())), (_index264).Int())
				var _2147_v55 _dafny.Set
				_ = _2147_v55
				_2147_v55 = _dafny.SetOf((_this).F33(), false)
				var _2148_v56 _dafny.Array
				_ = _2148_v56
				var _nwElement0_85 bool = true
				_ = _nwElement0_85
				var _nw367 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(10))
				_ = _nw367
				(_nw367).ArraySet1(_nwElement0_85, 0)
				(_nw367).ArraySet1((_this).F33(), 1)
				(_nw367).ArraySet1(p2, 2)
				(_nw367).ArraySet1(false, 3)
				(_nw367).ArraySet1((_this).F33(), 4)
				(_nw367).ArraySet1(p2, 5)
				(_nw367).ArraySet1(p0, 6)
				(_nw367).ArraySet1(p2, 7)
				(_nw367).ArraySet1(true, 8)
				(_nw367).ArraySet1(p2, 9)
				_2148_v56 = _nw367
				var _2149_v57 _dafny.Array
				_ = _2149_v57
				var _nwElement0_86 _dafny.Array = _2148_v56
				_ = _nwElement0_86
				var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(8))
				_ = _nw368
				(_nw368).ArraySet1(_nwElement0_86, 0)
				(_nw368).ArraySet1(_2148_v56, 1)
				(_nw368).ArraySet1(_2148_v56, 2)
				(_nw368).ArraySet1(_2148_v56, 3)
				(_nw368).ArraySet1(_2148_v56, 4)
				(_nw368).ArraySet1(_2148_v56, 5)
				(_nw368).ArraySet1(_2148_v56, 6)
				(_nw368).ArraySet1(_2148_v56, 7)
				_2149_v57 = _nw368
				var _2150_v58 D2
				_ = _2150_v58
				_2150_v58 = Companion_D2_.Create_DC4_((_2145_v53).F32(), _dafny.IntOfInt64(982), _this.F34, _2058_v5, _2057_v4)
				var _2151_v59 D16
				_ = _2151_v59
				_2151_v59 = Companion_D16_.Create_DC42_(p1, _2149_v57, p3, Companion_D2_.Create_DC5_(_2150_v58))
				var _2152_v60 _dafny.MultiSet
				_ = _2152_v60
				_2152_v60 = _dafny.MultiSetOf((_dafny.IntOfInt64(399)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2147_v55, p3)).Cardinality()), _this.F34, _this.F34, (_2151_v59).Dtor_cf78(), (_2145_v53).F32())
				var _2153_v61 _dafny.Sequence
				_ = _2153_v61
				_2153_v61 = _dafny.SeqOf(_2061_v9)
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))
				_ = _index265
				(_2143_v52).ArraySet1((func() _dafny.Int {
					if (_2152_v60).Contains((p1).Plus((_2143_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))).Int()).(_dafny.Int))) {
						return (_2152_v60).Multiplicity((p1).Plus((_2143_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))).Int()).(_dafny.Int)))
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2153_v61, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2153_v61).Cardinality()))).Uint32(), _2061_v9)).Cardinality())
				})(), (_index265).Int())
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))
				_ = _index266
				var _rhs390 bool = p0
				_ = _rhs390
				var _rhs391 _dafny.Int = (_2145_v53).F32()
				_ = _rhs391
				var _lhs300 *GlobalState = globalState
				_ = _lhs300
				var _lhs301 _dafny.Array = _2143_v52
				_ = _lhs301
				var _lhs302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_2143_v52), 0))
				_ = _lhs302
				_lhs300.F26 = _rhs390
				(_lhs301).ArraySet1(_rhs391, (_lhs302).Int())
				(globalState).F26 = true
				var _2154_v62 D12
				_ = _2154_v62
				_2154_v62 = Companion_D12_.Create_DC27_(p0, _2152_v60, (func() _dafny.Int {
					if (_2152_v60).Contains((_2061_v9).Cardinality()) {
						return (_2152_v60).Multiplicity((_2061_v9).Cardinality())
					}
					return p1
				})(), _2143_v52)
				var _2155_v63 _dafny.Sequence
				_ = _2155_v63
				_2155_v63 = _dafny.SeqOf(_2154_v62)
				(globalState).F24 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2155_v63, _dafny.SeqOf(_2154_v62))).Cardinality())).Minus(((_dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2057_v4, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32(), _dafny.CodePoint('k'))).Cardinality()))).Intersection(_2061_v9)).Cardinality())
			} else {
				(globalState).F0 = Companion_Default___.SafeModuloInt(_2141_cf43, (_this).F32())
				var _2156_v64 _dafny.Sequence
				_ = _2156_v64
				_2156_v64 = _dafny.SeqOf((_this).F27())
				var _2157_v65 D9
				_ = _2157_v65
				_2157_v65 = Companion_D9_.Create_DC18_(p0)
				var _rhs392 _dafny.Sequence = Companion_Default___.Fm9((_this).F33(), globalState)
				_ = _rhs392
				var _rhs393 bool = !(p3)
				_ = _rhs393
				var _rhs394 D9 = Companion_Default___.Fm52((_this).F33(), Companion_Default___.Fm0((_this).F32(), globalState), false, _dafny.IntOfInt64(914), globalState)
				_ = _rhs394
				var _rhs395 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2054_v1, _2054_v1)
				_ = _rhs395
				var _rhs396 _dafny.Int = _dafny.IntOfInt64(844)
				_ = _rhs396
				var _lhs303 *GlobalState = globalState
				_ = _lhs303
				var _lhs304 *GlobalState = globalState
				_ = _lhs304
				_2156_v64 = _rhs392
				_lhs303.F26 = _rhs393
				_2157_v65 = _rhs394
				_2054_v1 = _rhs395
				_lhs304.F21 = _rhs396
				(globalState).F18 = _this.F34
				(globalState).F3 = _this.F34
				var _2158_v66 _dafny.Array
				_ = _2158_v66
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_46
				var _nw369 _dafny.Array
				_ = _nw369
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw369 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) bool = (func(_2159_cf43 _dafny.Int, _2160_p1 _dafny.Int, _2161_v0 _dafny.Map) func(_dafny.Int) bool {
						return func(_2162_i6 _dafny.Int) bool {
							return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2159_cf43, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34, _2160_p1))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2159_cf43, _2161_v0))
						}
					})(_2141_cf43, p1, _2053_v0)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw369 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw369).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw369).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_2158_v66 = _nw369
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_2158_v66), 0))
				_ = _index267
				(_2158_v66).ArraySet1(p3, (_index267).Int())
				var _2163_v67 _dafny.Map
				_ = _2163_v67
				_2163_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2054_v1).Cardinality()), _2058_v5)
				var _2164_v68 _dafny.Array
				_ = _2164_v68
				var _nwElement0_87 _dafny.Int = ((_2163_v67).Cardinality()).Times((_2060_v7).Fm45(globalState))
				_ = _nwElement0_87
				var _nw370 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(5))
				_ = _nw370
				(_nw370).ArraySet1(_nwElement0_87, 0)
				(_nw370).ArraySet1(_dafny.IntOfInt64(638), 1)
				(_nw370).ArraySet1((p1).Plus(_2141_cf43), 2)
				(_nw370).ArraySet1((_2141_cf43).Times(p1), 3)
				(_nw370).ArraySet1(_2141_cf43, 4)
				_2164_v68 = _nw370
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_2164_v68), 0))
				_ = _index268
				(_2164_v68).ArraySet1((_this).F32(), (_index268).Int())
				var _2165_v69 D2
				_ = _2165_v69
				_2165_v69 = Companion_D2_.Create_DC4_((_this).Fm7(globalState), (_this).F32(), _2141_cf43, _2058_v5, Companion_Default___.Fm37((_2061_v9).Cardinality(), _this.F34, globalState))
				var _2166_v70 _dafny.Map
				_ = _2166_v70
				_2166_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2165_v69, p3)
				var _2167_v71 _dafny.Sequence
				_ = _2167_v71
				_2167_v71 = _dafny.SeqOf(_2166_v70)
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_2158_v66), 0))
				_ = _index269
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_2164_v68), 0))
				_ = _index270
				var _rhs397 bool = ((((Companion_D8_.Create_DC15_(_2142_v51)).Dtor_cf28()).Cardinality()).Plus(p1)).Cmp((_2142_v51).Cardinality()) <= 0
				_ = _rhs397
				var _rhs398 bool = !(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_2167_v71, _2167_v71))).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm36(p3, globalState), p2))
				_ = _rhs398
				var _rhs399 _dafny.Int = (_this).Fm7(globalState)
				_ = _rhs399
				var _rhs400 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(870), _2141_cf43)
				_ = _rhs400
				var _lhs305 *GlobalState = globalState
				_ = _lhs305
				var _lhs306 _dafny.Array = _2158_v66
				_ = _lhs306
				var _lhs307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_2158_v66), 0))
				_ = _lhs307
				var _lhs308 _dafny.Array = _2164_v68
				_ = _lhs308
				var _lhs309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_2164_v68), 0))
				_ = _lhs309
				var _lhs310 *GlobalState = globalState
				_ = _lhs310
				_lhs305.F26 = _rhs397
				(_lhs306).ArraySet1(_rhs398, (_lhs307).Int())
				(_lhs308).ArraySet1(_rhs399, (_lhs309).Int())
				_lhs310.F3 = _rhs400
			}
			var _2168_v72 _dafny.Array
			_ = _2168_v72
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_47
			var _nw371 _dafny.Array
			_ = _nw371
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw371 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) _dafny.Sequence = func(_2169_i7 _dafny.Int) _dafny.Sequence {
					return (_this).F30()
				}
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw371 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw371).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw371).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_2168_v72 = _nw371
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_2168_v72), 0))
			_ = _index271
			(_2168_v72).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-680))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg114 _dafny.Int) interface{} {
					return coer114(arg114)
				}
			}(func(_2170_i8 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F32()))
			})), (_index271).Int())
			var _2171_v73 _dafny.Map
			_ = _2171_v73
			_2171_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-289), true)
			var _2172_v74 _dafny.Array
			_ = _2172_v74
			var _nwElement0_88 bool = (_this).F33()
			_ = _nwElement0_88
			var _nw372 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(22))
			_ = _nw372
			(_nw372).ArraySet1(_nwElement0_88, 0)
			(_nw372).ArraySet1(p0, 1)
			(_nw372).ArraySet1(!(!(p0)) || (p0), 2)
			(_nw372).ArraySet1((p3) == (p0), 3)
			(_nw372).ArraySet1((_this).F33(), 4)
			(_nw372).ArraySet1(p3, 5)
			(_nw372).ArraySet1((p0) && (p0), 6)
			(_nw372).ArraySet1(p3, 7)
			(_nw372).ArraySet1(p3, 8)
			(_nw372).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(p1, (_this).F27()), Companion_Default___.Fm16((_this).Fm7(globalState), globalState)), 9)
			(_nw372).ArraySet1((_this).F33(), 10)
			(_nw372).ArraySet1(((_2171_v73).Cardinality()).Cmp(_this.F34) < 0, 11)
			(_nw372).ArraySet1(((_this).F31()).IsSubsetOf(Companion_Default___.Fm32(globalState)), 12)
			(_nw372).ArraySet1(!(p2), 13)
			(_nw372).ArraySet1((_2061_v9).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(525))), 14)
			(_nw372).ArraySet1(p3, 15)
			(_nw372).ArraySet1(Companion_Default___.Fm0((_this).Fm7(globalState), globalState), 16)
			(_nw372).ArraySet1(p3, 17)
			(_nw372).ArraySet1((_this).Fm5(globalState), 18)
			(_nw372).ArraySet1((p1).Cmp((_this).F32()) < 0, 19)
			(_nw372).ArraySet1(!(p3), 20)
			(_nw372).ArraySet1(p0, 21)
			_2172_v74 = _nw372
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2172_v74), 0))
			_ = _index272
			(_2172_v74).ArraySet1(false, (_index272).Int())
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_2168_v72), 0))
			_ = _index273
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2172_v74), 0))
			_ = _index274
			var _rhs401 bool = !(p0) || (p3)
			_ = _rhs401
			var _rhs402 _dafny.Int = p1
			_ = _rhs402
			var _rhs403 _dafny.Sequence = _dafny.SeqOf(_this.F34)
			_ = _rhs403
			var _rhs404 bool = (_this).F33()
			_ = _rhs404
			var _rhs405 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2054_v1, _dafny.SeqOf(p2)), _2054_v1)
			_ = _rhs405
			var _lhs311 *GlobalState = globalState
			_ = _lhs311
			var _lhs312 *GlobalState = globalState
			_ = _lhs312
			var _lhs313 _dafny.Array = _2168_v72
			_ = _lhs313
			var _lhs314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_2168_v72), 0))
			_ = _lhs314
			var _lhs315 _dafny.Array = _2172_v74
			_ = _lhs315
			var _lhs316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2172_v74), 0))
			_ = _lhs316
			_lhs311.F26 = _rhs401
			_lhs312.F24 = _rhs402
			(_lhs313).ArraySet1(_rhs403, (_lhs314).Int())
			(_lhs315).ArraySet1(_rhs404, (_lhs316).Int())
			_2054_v1 = _rhs405
			(globalState).F3 = _2141_cf43
		} else if _source31.Is_DC23() {
			var _2173___mcc_h15 _dafny.MultiSet = _source31.Get_().(D11_DC23).Cf42
			_ = _2173___mcc_h15
			var _2174_cf42 _dafny.MultiSet = _2173___mcc_h15
			_ = _2174_cf42
			var _2175_v75 D15
			_ = _2175_v75
			_2175_v75 = Companion_D15_.Create_DC36_(_2058_v5, (_this).F32())
			var _source32 D15 = _2175_v75
			_ = _source32
			if _source32.Is_DC36() {
				var _2176___mcc_h17 _dafny.CodePoint = _source32.Get_().(D15_DC36).Cf64
				_ = _2176___mcc_h17
				var _2177___mcc_h18 _dafny.Int = _source32.Get_().(D15_DC36).Cf65
				_ = _2177___mcc_h18
				var _2178_cf65 _dafny.Int = _2177___mcc_h18
				_ = _2178_cf65
				var _2179_cf64 _dafny.CodePoint = _2176___mcc_h17
				_ = _2179_cf64
				var _2180_v76 D16
				_ = _2180_v76
				_2180_v76 = Companion_D16_.Create_DC40_(_dafny.Companion_Sequence_.IsProperPrefixOf(_2057_v4, _dafny.Companion_Sequence_.Update(_2057_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32(), _dafny.CodePoint('g'))), _this.F34, (_dafny.Zero).Minus((func() _dafny.Int {
					if !(p0) {
						return (_dafny.Zero).Minus((_this).F27())
					}
					return (_this).F27()
				})()))
				var _2181_v77 _dafny.Array
				_ = _2181_v77
				var _nw373 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw373
				_2181_v77 = _nw373
				var _2182_v78 _dafny.Sequence
				_ = _2182_v78
				_2182_v78 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2054_v1))
				var _2183_v79 _dafny.Map
				_ = _2183_v79
				_2183_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2061_v9)
				var _2184_v80 _dafny.Map
				_ = _2184_v80
				_2184_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), p0)
				var _pat_let_tv46 = _2182_v78
				_ = _pat_let_tv46
				var _pat_let_tv47 = _2182_v78
				_ = _pat_let_tv47
				var _rhs406 D16 = func(_pat_let62_0 D16) D16 {
					return func(_2185_dt__update__tmp_h0 D16) D16 {
						return func(_pat_let63_0 _dafny.Int) D16 {
							return func(_2186_dt__update_hcf75_h0 _dafny.Int) D16 {
								return Companion_D16_.Create_DC40_((_2185_dt__update__tmp_h0).Dtor_cf74(), _2186_dt__update_hcf75_h0, (_2185_dt__update__tmp_h0).Dtor_cf76())
							}(_pat_let63_0)
						}(((_pat_let_tv46).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_pat_let_tv47).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
					}(_pat_let62_0)
				}(Companion_D16_.Create_DC40_(p3, _2178_cf65, p1))
				_ = _rhs406
				var _rhs407 _dafny.Int = ((func() _dafny.Set {
					if (_2183_v79).Contains(((_2184_v80).Cardinality()).Times(_this.F34)) {
						return (_2183_v79).Get(((_2184_v80).Cardinality()).Times(_this.F34)).(_dafny.Set)
					}
					return _2061_v9
				})()).Cardinality()
				_ = _rhs407
				var _rhs408 _dafny.Array = (func() _dafny.Array {
					if (_this).F33() {
						return _2181_v77
					}
					return _2181_v77
				})()
				_ = _rhs408
				var _lhs317 *GlobalState = globalState
				_ = _lhs317
				_2180_v76 = _rhs406
				_lhs317.F15 = _rhs407
				_2181_v77 = _rhs408
				(globalState).F0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2057_v4, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2057_v4).Cardinality()), _dafny.IntOfUint32((_2057_v4).Cardinality()))).Uint32(), _2058_v5)).Cardinality())
				var _2187_v81 T0
				_ = _2187_v81
				var _nw374 *C4 = New_C4_()
				_ = _nw374
				_nw374.Ctor__(p1, (_this).F28())
				_2187_v81 = _nw374
				var _2188_v82 _dafny.Sequence
				_ = _2188_v82
				_2188_v82 = _2057_v4
				var _2189_v83 _dafny.MultiSet
				_ = _2189_v83
				_2189_v83 = _dafny.MultiSetOf(_2057_v4)
				var _2190_v84 D7
				_ = _2190_v84
				_2190_v84 = Companion_D7_.Create_DC14_((_this).F27(), _2187_v81, _2188_v82, _2189_v83)
				var _2191_v85 _dafny.Map
				_ = _2191_v85
				_2191_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2190_v84)
				_2191_v85 = (_2191_v85).Update(_dafny.IntOfInt64(165), _2190_v84)
				var _2192_v86 _dafny.Map
				_ = _2192_v86
				_2192_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2058_v5, p1)
				_2192_v86 = (_2192_v86).Update(_dafny.CodePoint('n'), (_this).F27())
			} else if _source32.Is_DC37() {
				var _2193___mcc_h19 bool = _source32.Get_().(D15_DC37).Cf66
				_ = _2193___mcc_h19
				var _2194___mcc_h20 _dafny.Int = _source32.Get_().(D15_DC37).Cf67
				_ = _2194___mcc_h20
				var _2195___mcc_h21 _dafny.Array = _source32.Get_().(D15_DC37).Cf68
				_ = _2195___mcc_h21
				var _2196___mcc_h22 bool = _source32.Get_().(D15_DC37).Cf69
				_ = _2196___mcc_h22
				var _2197_cf69 bool = _2196___mcc_h22
				_ = _2197_cf69
				var _2198_cf68 _dafny.Array = _2195___mcc_h21
				_ = _2198_cf68
				var _2199_cf67 _dafny.Int = _2194___mcc_h20
				_ = _2199_cf67
				var _2200_cf66 bool = _2193___mcc_h19
				_ = _2200_cf66
				var _2201_v87 _dafny.Array
				_ = _2201_v87
				var _nw375 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw375
				_2201_v87 = _nw375
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_2201_v87), 0))
				_ = _index275
				(_2201_v87).ArraySet1(_2197_cf69, (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_2201_v87), 0))
				_ = _index276
				(_2201_v87).ArraySet1((_dafny.IntOfInt64(-729)).Cmp(((_this).F30()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}((func(_2202_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2203_i9 _dafny.Int) _dafny.CodePoint {
						return _2202_v5
					}
				})(_2058_v5)))).Cardinality()), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.Int)) != 0, (_index276).Int())
				var _2204_v88 _dafny.Set
				_ = _2204_v88
				_2204_v88 = _dafny.SetOf(_2200_cf66, (_2201_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_2201_v87), 0))).Int()).(bool))
				var _2205_v89 _dafny.MultiSet
				_ = _2205_v89
				_2205_v89 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if _2197_cf69 {
						return _2054_v1
					}
					return _2054_v1
				})(), _2054_v1, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm41(_2199_cf67, (_2204_v88).Cardinality(), p0, globalState), _2054_v1))
				_2205_v89 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm44(globalState), _dafny.SeqOf(_2054_v1)))
				var _2206_v90 _dafny.Array
				_ = _2206_v90
				var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
				_ = _nw376
				_2206_v90 = _nw376
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2206_v90), 0))
				_ = _index277
				(_2206_v90).ArraySet1CodePoint(_2058_v5, (_index277).Int())
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2206_v90), 0))
				_ = _index278
				(_2206_v90).ArraySet1CodePoint(_2058_v5, (_index278).Int())
				var _2207_v91 _dafny.Sequence
				_ = _2207_v91
				_2207_v91 = _dafny.SeqOf((_this).F27())
				_2207_v91 = (_this).F30()
			} else if _source32.Is_DC38() {
				var _2208___mcc_h23 _dafny.Int = _source32.Get_().(D15_DC38).Cf70
				_ = _2208___mcc_h23
				var _2209___mcc_h24 bool = _source32.Get_().(D15_DC38).Cf71
				_ = _2209___mcc_h24
				var _2210___mcc_h25 _dafny.Int = _source32.Get_().(D15_DC38).Cf72
				_ = _2210___mcc_h25
				var _2211_cf72 _dafny.Int = _2210___mcc_h25
				_ = _2211_cf72
				var _2212_cf71 bool = _2209___mcc_h24
				_ = _2212_cf71
				var _2213_cf70 _dafny.Int = _2208___mcc_h23
				_ = _2213_cf70
				var _2214_v92 _dafny.Map
				_ = _2214_v92
				_2214_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2061_v9).Intersection(_2061_v9))
				(globalState).F3 = ((func() _dafny.Set {
					if (_2214_v92).Contains(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg116 _dafny.Int) interface{} {
							return coer116(arg116)
						}
					}((func(_2215_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2216_i10 _dafny.Int) _dafny.CodePoint {
							return _2215_v5
						}
					})(_2058_v5))), _2057_v4)) {
						return (_2214_v92).Get(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg117 _dafny.Int) interface{} {
								return coer117(arg117)
							}
						}((func(_2217_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2218_i10 _dafny.Int) _dafny.CodePoint {
								return _2217_v5
							}
						})(_2058_v5))), _2057_v4)).(_dafny.Set)
					}
					return _2061_v9
				})()).Cardinality()
				var _2219_v93 _dafny.Array
				_ = _2219_v93
				var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw377
				_2219_v93 = _nw377
				_2219_v93 = _2219_v93
				(globalState).F15 = p1
				var _2220_v94 _dafny.Array
				_ = _2220_v94
				var _nwElement0_89 _dafny.Int = _2211_cf72
				_ = _nwElement0_89
				var _nw378 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(3))
				_ = _nw378
				(_nw378).ArraySet1(_nwElement0_89, 0)
				(_nw378).ArraySet1(_2211_cf72, 1)
				(_nw378).ArraySet1((_this).F32(), 2)
				_2220_v94 = _nw378
				var _2221_v95 _dafny.Map
				_ = _2221_v95
				_2221_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2220_v94, (_this).F29())
				var _2222_v96 *C5
				_ = _2222_v96
				var _nw379 *C5 = New_C5_()
				_ = _nw379
				_nw379.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2053_v0, _dafny.IntOfUint32(((_this).F30()).Cardinality())), (_this).F29(), _dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex((((func() _dafny.MultiSet {
					if (_2221_v95).Contains(_2220_v94) {
						return (_2221_v95).Get(_2220_v94).(_dafny.MultiSet)
					}
					return (_this).F29()
				})()).Update(false, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F30(), (Companion_Default___.SafeIndex(_this.F34, _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), _2211_cf72)).Cardinality())))).Cardinality(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32(), _this.F34), (func() _dafny.Int {
					if p2 {
						return p1
					}
					return (func() _dafny.Int {
						if ((_this).F29()).Contains(p2) {
							return ((_this).F29()).Multiplicity(p2)
						}
						return p1
					})()
				})(), (_this).F28())
				_2222_v96 = _nw379
			} else {
				var _2223___mcc_h26 _dafny.Sequence = _source32.Get_().(D15_DC35).Cf63
				_ = _2223___mcc_h26
				var _2224_cf63 _dafny.Sequence = _2223___mcc_h26
				_ = _2224_cf63
				var _2225_v97 D8
				_ = _2225_v97
				_2225_v97 = Companion_D8_.Create_DC16_(_this.F34, (_this).F27())
				var _2226_v98 _dafny.Map
				_ = _2226_v98
				_2226_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2225_v97, (_2053_v0).Cardinality())
				(globalState).F0 = (_2226_v98).Cardinality()
				r0 = ((_2061_v9).Union(_2061_v9)).Union(_2061_v9)
				var _2227_v99 *C2
				_ = _2227_v99
				var _nw380 *C2 = New_C2_()
				_ = _nw380
				_nw380.Ctor__(_dafny.MultiSetOf(p0), _this.F34)
				_2227_v99 = _nw380
				var _2228_v100 _dafny.Map
				_ = _2228_v100
				_2228_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2054_v1, false)
				_2228_v100 = ((func() _dafny.Map {
					if p3 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2054_v1, p0)
					}
					return _2228_v100
				})()).Merge(Companion_Default___.Fm53(_2058_v5, globalState))
			}
			var _2229_v101 _dafny.Array
			_ = _2229_v101
			var _nw381 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
			_ = _nw381
			_2229_v101 = _nw381
			var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_2229_v101), 0))
			_ = _index279
			(_2229_v101).ArraySet1((_this).F31(), (_index279).Int())
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_2229_v101), 0))
			_ = _index280
			(_2229_v101).ArraySet1(((_this).F29()).Union((_this).F29()), (_index280).Int())
			(globalState).F26 = (_this).F33()
			var _2230_v102 _dafny.MultiSet
			_ = _2230_v102
			_2230_v102 = _dafny.MultiSetOf((_this).F32())
			var _2231_v103 *C2
			_ = _2231_v103
			var _nw382 *C2 = New_C2_()
			_ = _nw382
			_nw382.Ctor__(_dafny.MultiSetOf(p0, (_this).F33(), (_this).F33(), true, p3), (_2230_v102).Cardinality())
			_2231_v103 = _nw382
		} else {
			var _2232___mcc_h16 D11 = _source31.Get_().(D11_DC25).Cf45
			_ = _2232___mcc_h16
			var _2233_cf45 D11 = _2232___mcc_h16
			_ = _2233_cf45
			var _2234_v104 *C4
			_ = _2234_v104
			var _nw383 *C4 = New_C4_()
			_ = _nw383
			_nw383.Ctor__(_this.F34, (_this).F28())
			_2234_v104 = _nw383
			var _2235_v106 _dafny.MultiSet
			_ = _2235_v106
			_2235_v106 = _dafny.MultiSetOf((_this).F27(), p1, p1)
			var _2236_v107 _dafny.Array
			_ = _2236_v107
			var _nwElement0_90 bool = !((_this).F33())
			_ = _nwElement0_90
			var _nw384 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(27))
			_ = _nw384
			(_nw384).ArraySet1(_nwElement0_90, 0)
			(_nw384).ArraySet1(p0, 1)
			(_nw384).ArraySet1(p0, 2)
			(_nw384).ArraySet1(p0, 3)
			(_nw384).ArraySet1(p0, 4)
			(_nw384).ArraySet1(!_dafny.Companion_Sequence_.Equal(_2054_v1, _2054_v1), 5)
			(_nw384).ArraySet1(p2, 6)
			(_nw384).ArraySet1((_dafny.IntOfUint32((_2057_v4).Cardinality())).Cmp((_dafny.Zero).Minus((_this).F27())) < 0, 7)
			(_nw384).ArraySet1(!(false) || (p3), 8)
			(_nw384).ArraySet1((_this.F34).Cmp(p1) <= 0, 9)
			(_nw384).ArraySet1(!((_this).F33()) || (p2), 10)
			(_nw384).ArraySet1(!(p0) || (p3), 11)
			(_nw384).ArraySet1((func() _dafny.Map {
				var _coll69 = _dafny.NewMapBuilder()
				_ = _coll69
				for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(371), _dafny.IntOfInt64(979))); ; {
					_compr_69, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _2237_v105 _dafny.Int
					_2237_v105 = interface{}(_compr_69).(_dafny.Int)
					if ((_dafny.IntOfInt64(371)).Cmp(_2237_v105) <= 0) && ((_2237_v105).Cmp(_dafny.IntOfInt64(979)) < 0) {
						_coll69.Add((_2237_v105).Minus(_dafny.IntOfInt64(566)), (_this).F32())
					}
				}
				return _coll69.ToMap()
			}()).Contains((_this).F32()), 12)
			(_nw384).ArraySet1((_this).F33(), 13)
			(_nw384).ArraySet1((_2235_v106).IsSubsetOf(_2235_v106), 14)
			(_nw384).ArraySet1(p2, 15)
			(_nw384).ArraySet1(p3, 16)
			(_nw384).ArraySet1(!((_2059_v6).Fm3((_this).F31(), globalState)), 17)
			(_nw384).ArraySet1((_2054_v1).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_2054_v1).Cardinality()))).Uint32()).(bool), 18)
			(_nw384).ArraySet1(p0, 19)
			(_nw384).ArraySet1((_this).F33(), 20)
			(_nw384).ArraySet1(p2, 21)
			(_nw384).ArraySet1(p0, 22)
			(_nw384).ArraySet1(p3, 23)
			(_nw384).ArraySet1(p0, 24)
			(_nw384).ArraySet1(p2, 25)
			(_nw384).ArraySet1(p3, 26)
			_2236_v107 = _nw384
			var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_2236_v107), 0))
			_ = _index281
			(_2236_v107).ArraySet1((_this).F33(), (_index281).Int())
			var _2238_v108 _dafny.Map
			_ = _2238_v108
			_2238_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).Cmp(p1) <= 0, !((_2234_v104).Fm35(_2058_v5, (_this).F27(), globalState)))
			var _2239_v109 _dafny.Map
			_ = _2239_v109
			_2239_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _2057_v4)
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_2236_v107), 0))
			_ = _index282
			var _rhs409 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_2239_v109).Contains(_this.F34) {
					return (_2239_v109).Get(_this.F34).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_2240_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2241_i11 _dafny.Int) _dafny.CodePoint {
						return _2240_v5
					}
				})(_2058_v5)))
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg119 _dafny.Int) interface{} {
					return coer119(arg119)
				}
			}((func(_2242_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2243_i12 _dafny.Int) _dafny.CodePoint {
					return _2242_v5
				}
			})(_2058_v5)))), _2058_v5)
			_ = _rhs409
			var _rhs410 _dafny.Map = _2238_v108
			_ = _rhs410
			var _lhs318 _dafny.Array = _2236_v107
			_ = _lhs318
			var _lhs319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_2236_v107), 0))
			_ = _lhs319
			(_lhs318).ArraySet1(_rhs409, (_lhs319).Int())
			_2238_v108 = _rhs410
			var _2244_v110 _dafny.Array
			_ = _2244_v110
			var _nw385 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw385
			_2244_v110 = _nw385
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_2244_v110), 0))
			_ = _index283
			(_2244_v110).ArraySet1((_this).F32(), (_index283).Int())
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_2244_v110), 0))
			_ = _index284
			(_2244_v110).ArraySet1(Companion_Default___.Fm1(globalState), (_index284).Int())
			_2057_v4 = Companion_Default___.Fm37((func() _dafny.Int {
				if (_2053_v0).Contains(_dafny.IntOfInt64(333)) {
					return (_2053_v0).Get(_dafny.IntOfInt64(333)).(_dafny.Int)
				}
				return (_this).F32()
			})(), _dafny.IntOfInt64(474), globalState)
		}
		var _hi12 _dafny.Int = (_this).F27()
		_ = _hi12
		for _2245_i13 := (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality())).Plus(_this.F34); _2245_i13.Cmp(_hi12) < 0; _2245_i13 = _2245_i13.Plus(_dafny.One) {
			r1 = (_2245_i13).Times(_2245_i13)
			var _2246_v111 *C3
			_ = _2246_v111
			var _nw386 *C3 = New_C3_()
			_ = _nw386
			_nw386.Ctor__((_this).F32(), (_this).F28())
			_2246_v111 = _nw386
			(globalState).F26 = !(p3) || ((_this).F33())
			var _2247_v112 _dafny.Array
			_ = _2247_v112
			var _len0_48 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_48
			var _nw387 _dafny.Array
			_ = _nw387
			if _len0_48.Cmp(_dafny.Zero) == 0 {
				_nw387 = _dafny.NewArray(_len0_48)
			} else {
				var _init48 func(_dafny.Int) _dafny.Map = (func(_2248_p0 bool) func(_dafny.Int) _dafny.Map {
					return func(_2249_i14 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2248_p0, true)
					}
				})(p0)
				_ = _init48
				var _element0_48 = _init48(_dafny.Zero)
				_ = _element0_48
				_nw387 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
				(_nw387).ArraySet1(_element0_48, 0)
				var _nativeLen0_48 = (_len0_48).Int()
				_ = _nativeLen0_48
				for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
					(_nw387).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
				}
			}
			_2247_v112 = _nw387
			var _2250_v113 _dafny.Set
			_ = _2250_v113
			_2250_v113 = _dafny.SetOf(_2057_v4, Companion_Default___.Fm8(globalState))
			var _rhs411 _dafny.Int = Companion_Default___.SafeDivisionInt(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(831), (_this).F27(), (_this).F27(), _2245_i13))).Cardinality()).Minus(_dafny.IntOfUint32((_2057_v4).Cardinality())), p1)
			_ = _rhs411
			var _rhs412 _dafny.Map = Companion_Default___.Fm15((_this).Fm4(_2250_v113, !(p3), _2054_v1, globalState), _2245_i13, globalState)
			_ = _rhs412
			var _rhs413 _dafny.Array = _2247_v112
			_ = _rhs413
			var _rhs414 _dafny.Int = p1
			_ = _rhs414
			var _lhs320 *GlobalState = globalState
			_ = _lhs320
			var _lhs321 *GlobalState = globalState
			_ = _lhs321
			_lhs320.F18 = _rhs411
			_2053_v0 = _rhs412
			_2247_v112 = _rhs413
			_lhs321.F15 = _rhs414
		}
		var _2251_v114 _dafny.Set
		_ = _2251_v114
		_2251_v114 = _dafny.SetOf((_this).F31())
		(globalState).F21 = ((_this.F34).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hgatgwca")).Cardinality()))).Times((_2251_v114).Cardinality())
		var _2252_v115 _dafny.MultiSet
		_ = _2252_v115
		_2252_v115 = _dafny.MultiSetOf(_this.F34, p1)
		var _2253_v116 D4
		_ = _2253_v116
		_2253_v116 = Companion_D4_.Create_DC9_(_this.F34, p3)
		var _2254_v117 _dafny.Map
		_ = _2254_v117
		_2254_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2252_v115).Cardinality(), _2253_v116)
		r0 = func() _dafny.Set {
			var _coll70 = _dafny.NewBuilder()
			_ = _coll70
			for _iter79 := _dafny.Iterate((_2254_v117).Keys().Elements()); ; {
				_compr_70, _ok79 := _iter79()
				if !_ok79 {
					break
				}
				var _2255_v118 _dafny.Int
				_2255_v118 = interface{}(_compr_70).(_dafny.Int)
				if (_2254_v117).Contains(_2255_v118) {
					_coll70.Add((_2255_v118).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-421), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), (_dafny.MultiSetOf(_dafny.IntOfInt64(280))).Cardinality())).Cardinality())))
				}
			}
			return _coll70.ToSet()
		}()
		r1 = _dafny.IntOfInt64(245)
		var _nw388 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw388
		r2 = _nw388
		return r0, r1, r2
	}
}
func (_this *C6) F33() bool {
	{
		return _this._f33
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
