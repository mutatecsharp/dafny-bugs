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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32(((Companion_D9_.Create_DC27_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("itjb"), false)).Cardinality(), (_dafny.SetOf(true, !(false))).Cardinality()))).Dtor_cf40()).Cardinality()))).Plus((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-111))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-111)), _0_v0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_0_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("exuvfvup")).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	var _source0 D15 = Companion_D15_.Create_DC46_()
	_ = _source0
	if _source0.Is_DC44() {
		var _1___mcc_h0 _dafny.Int = _source0.Get_().(D15_DC44).Cf71
		_ = _1___mcc_h0
		var _2___mcc_h1 bool = _source0.Get_().(D15_DC44).Cf72
		_ = _2___mcc_h1
		var _3___mcc_h2 _dafny.Array = _source0.Get_().(D15_DC44).Cf73
		_ = _3___mcc_h2
		var _4___mcc_h3 _dafny.Array = _source0.Get_().(D15_DC44).Cf74
		_ = _4___mcc_h3
		var _5___mcc_h4 _dafny.Int = _source0.Get_().(D15_DC44).Cf75
		_ = _5___mcc_h4
		var _6_cf75 _dafny.Int = _5___mcc_h4
		_ = _6_cf75
		var _7_cf74 _dafny.Array = _4___mcc_h3
		_ = _7_cf74
		var _8_cf73 _dafny.Array = _3___mcc_h2
		_ = _8_cf73
		var _9_cf72 bool = _2___mcc_h1
		_ = _9_cf72
		var _10_cf71 _dafny.Int = _1___mcc_h0
		_ = _10_cf71
		return _9_cf72
	} else if _source0.Is_DC45() {
		var _11___mcc_h5 _dafny.Int = _source0.Get_().(D15_DC45).Cf76
		_ = _11___mcc_h5
		var _12___mcc_h6 _dafny.Int = _source0.Get_().(D15_DC45).Cf77
		_ = _12___mcc_h6
		var _13___mcc_h7 bool = _source0.Get_().(D15_DC45).Cf78
		_ = _13___mcc_h7
		var _14___mcc_h8 bool = _source0.Get_().(D15_DC45).Cf79
		_ = _14___mcc_h8
		var _15___mcc_h9 _dafny.Array = _source0.Get_().(D15_DC45).Cf80
		_ = _15___mcc_h9
		var _16_cf80 _dafny.Array = _15___mcc_h9
		_ = _16_cf80
		var _17_cf79 bool = _14___mcc_h8
		_ = _17_cf79
		var _18_cf78 bool = _13___mcc_h7
		_ = _18_cf78
		var _19_cf77 _dafny.Int = _12___mcc_h6
		_ = _19_cf77
		var _20_cf76 _dafny.Int = _11___mcc_h5
		_ = _20_cf76
		return _18_cf78
	} else if _source0.Is_DC46() {
		return !(true)
	} else {
		var _21___mcc_h10 _dafny.Map = _source0.Get_().(D15_DC43).Cf70
		_ = _21___mcc_h10
		var _22_cf70 _dafny.Map = _21___mcc_h10
		_ = _22_cf70
		return (!(true)) && (!(true))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("hgxt")
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, globalState *GlobalState) D2 {
	if !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(759))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_23_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(!(false), !(true), false, false, false)).Cardinality()
	})), _dafny.SeqOf(_dafny.IntOfInt64(6))) {
		return Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("txpfrb"), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-952), _dafny.IntOfUint32((_dafny.SeqOf(true, false, false, true, false)).Cardinality()))).Cardinality()), true, _dafny.IntOfInt64(-911))
	} else {
		return Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("asarcrdjh"), (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-842)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(545))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_24_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-674)
		})), _dafny.SeqOf(_dafny.IntOfInt64(46)))).Cardinality(), true, _dafny.IntOfInt64(585))
	}
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(444), ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wuuxa")).Cardinality())))).Times(_dafny.IntOfInt64(-128)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate((Companion_D9_.Create_DC28_(_dafny.SeqOf(false), false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(348), _dafny.IntOfInt64(715))).Cardinality(), false, !(true))).Dtor_cf41(), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wqkswi")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(12)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(84))).Cardinality()))).Difference(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("awlgsvyam"))).Cardinality())).Cardinality())))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _25_v0 _dafny.Map
			_25_v0 = interface{}(_compr_1).(_dafny.Map)
			if ((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(12)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(84))).Cardinality()))).Difference(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("awlgsvyam"))).Cardinality())).Cardinality())))).Contains(_25_v0) {
				_coll1.Add(_25_v0)
			}
		}
		return _coll1.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(681))).Difference(_dafny.SetOf(_dafny.IntOfInt64(340)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("nc"), (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(957), _dafny.IntOfInt64(-55))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _26_v0 _dafny.Int
			_26_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(957)).Cmp(_26_v0) <= 0) && ((_26_v0).Cmp(_dafny.IntOfInt64(-55)) < 0) {
				_coll2.Add((_26_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _dafny.IntOfInt64(-164))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality(), true, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()).Times((_dafny.MultiSetOf(!(!((Companion_D9_.Create_DC29_(_dafny.IntOfUint32(((Companion_D9_.Create_DC27_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(849))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_27_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(388)
	})))).Dtor_cf40()).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqe")).Cardinality()), _dafny.IntOfInt64(37))).Cardinality(), true)).Dtor_cf48())))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.CodePoint, globalState *GlobalState) bool {
	return ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.SetOf(true, true, false)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("enqyyh")).Cardinality()))))).Cmp((_dafny.IntOfInt64(-519)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aabgm")).Cardinality()))) > 0
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 bool, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fnjmfaxg"), _dafny.UnicodeSeqOfUtf8Bytes("ms")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_28_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_29_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D5 {
	if (true) && (false) {
		return Companion_D5_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D1_.Create_DC3_(_dafny.SeqOf(true))))
	} else {
		return Companion_D5_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D1_.Create_DC3_(_dafny.SeqOf(false))))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true)
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(!(!(true)))), _dafny.SeqOf(_dafny.SeqOf(true))), _dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(true, false, false), _dafny.SeqOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(false)
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	var _source1 D20 = Companion_D20_.Create_DC62_()
	_ = _source1
	if _source1.Is_DC62() {
		return func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vkghxcks")).Cardinality())), false)).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _30_v0 _dafny.Sequence
				_30_v0 = interface{}(_compr_3).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vkghxcks")).Cardinality())), false)).Contains(_30_v0) {
					_coll3.Add(_30_v0)
				}
			}
			return _coll3.ToSet()
		}()
	} else if _source1.Is_DC63() {
		var _31___mcc_h0 _dafny.Int = _source1.Get_().(D20_DC63).Cf112
		_ = _31___mcc_h0
		var _32_cf112 _dafny.Int = _31___mcc_h0
		_ = _32_cf112
		return _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-605))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_33_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(710))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_34_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))).Cardinality())
		})), _dafny.SeqOf(_dafny.IntOfInt64(605)))
	} else if _source1.Is_DC61() {
		var _35___mcc_h1 _dafny.Sequence = _source1.Get_().(D20_DC61).Cf111
		_ = _35___mcc_h1
		var _36_cf111 _dafny.Sequence = _35___mcc_h1
		_ = _36_cf111
		return (Companion_D21_.Create_DC65_(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(747), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-938), !(true))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-416))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_37_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-799)
			})))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _38_v1 _dafny.Sequence
				_38_v1 = interface{}(_compr_4).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(747), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-938), !(true))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-416))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_37_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-799)
				}))), _38_v1) {
					_coll4.Add(_38_v1)
				}
			}
			return _coll4.ToSet()
		}())).Dtor_cf114()
	} else {
		var _39___mcc_h2 D20 = _source1.Get_().(D20_DC64).Cf113
		_ = _39___mcc_h2
		var _40_cf113 D20 = _39___mcc_h2
		_ = _40_cf113
		return (_dafny.SetOf(_dafny.SeqOf((_dafny.SetOf(true, true, true)).Cardinality()))).Union(_dafny.SetOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_41_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(530)
		}))).Cardinality())), _dafny.SeqOf((func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_42_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality())))).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _43_v2 _dafny.Int
				_43_v2 = interface{}(_compr_5).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_42_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()))), _43_v2) {
					_coll5.Add((_43_v2).Minus(_dafny.IntOfInt64(367)), _dafny.IntOfInt64(-245))
				}
			}
			return _coll5.ToMap()
		}()).Cardinality()))).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqoca")).Cardinality())) < 0 {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), (_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xvqqeji")).Cardinality()))).Cardinality())).Cardinality())).Merge(func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('j'), _dafny.CodePoint('l'))).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _44_v0 _dafny.CodePoint
				_44_v0 = interface{}(_compr_6).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('j'), _dafny.CodePoint('l')), _44_v0) {
					_coll6.Add(_44_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), (func() _dafny.Set {
						var _coll7 = _dafny.NewBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Elements()); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _45_v1 _dafny.Set
							_45_v1 = interface{}(_compr_7).(_dafny.Set)
							if (_dafny.MultiSetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Contains(_45_v1) {
								_coll7.Add(_45_v1)
							}
						}
						return _coll7.ToSet()
					}()).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vpk")).Cardinality()), (func() _dafny.Map {
						var _coll8 = _dafny.NewMapBuilder()
						_ = _coll8
						for _iter8 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(!(false), !(false))).Cardinality())).Elements()); ; {
							_compr_8, _ok8 := _iter8()
							if !_ok8 {
								break
							}
							var _46_v2 _dafny.Int
							_46_v2 = interface{}(_compr_8).(_dafny.Int)
							if (_dafny.MultiSetOf((_dafny.MultiSetOf(!(false), !(false))).Cardinality())).Contains(_46_v2) {
								_coll8.Add((_46_v2).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()), false)
							}
						}
						return _coll8.ToMap()
					}()).Cardinality())).Cardinality()))).Cardinality())))
				}
			}
			return _coll6.ToMap()
		}())
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfInt64(334))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfInt64(642)))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(453), _dafny.IntOfInt64(892)), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_47_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nbv")).Cardinality())))), ((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, !(false), false))).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uibebmq")).Cardinality())), _dafny.IntOfUint32(((Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("vmqtwapqp"), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sdv")).Cardinality()), false, (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true, !(!(!(false))))).Cardinality()))).Cardinality()))).Cardinality())).Dtor_cf14()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_(true, false))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_D6_.Create_DC21_(_dafny.SetOf(!(false))), Companion_D6_.Create_DC21_(_dafny.SetOf(!(true), true)), Companion_D6_.Create_DC21_(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!((_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-252)))).IsProperSubsetOf((Companion_D17_.Create_DC52_(false, false, _dafny.IntOfInt64(-932), _dafny.MultiSetOf(_dafny.IntOfInt64(162)))).Dtor_cf91()))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	return (false) && (!(false))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("exnq"), _dafny.UnicodeSeqOfUtf8Bytes("gpfxxmfqm")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D1_.Create_DC3_(_dafny.SeqOf(true, true, false))).Dtor_cf7()
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Int {
	return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-821))).Minus((_dafny.IntOfInt64(986)).Minus(_dafny.IntOfInt64(174)))
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) D6 {
	var _source2 D21 = Companion_D21_.Create_DC65_(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(146))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_49_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(452)
	}))))
	_ = _source2
	if _source2.Is_DC66() {
		var _50___mcc_h0 _dafny.Int = _source2.Get_().(D21_DC66).Cf115
		_ = _50___mcc_h0
		var _51_cf115 _dafny.Int = _50___mcc_h0
		_ = _51_cf115
		return Companion_D6_.Create_DC22_(_51_cf115, (_dafny.MultiSetOf(_dafny.CodePoint('q'))).Cardinality(), true)
	} else if _source2.Is_DC67() {
		var _52___mcc_h1 _dafny.Int = _source2.Get_().(D21_DC67).Cf116
		_ = _52___mcc_h1
		var _53_cf116 _dafny.Int = _52___mcc_h1
		_ = _53_cf116
		return Companion_D6_.Create_DC22_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_cf116, _53_cf116)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D16_.Create_DC48_(), _53_cf116)).Cardinality(), false)
	} else {
		var _54___mcc_h2 _dafny.Set = _source2.Get_().(D21_DC65).Cf114
		_ = _54___mcc_h2
		var _55_cf114 _dafny.Set = _54___mcc_h2
		_ = _55_cf114
		return Companion_D6_.Create_DC22_((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))).Cardinality(), _dafny.IntOfInt64(548), false)
	}
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) D9 {
	if !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false, false, true), false) {
		return Companion_D9_.Create_DC27_(_dafny.SeqOf(_dafny.IntOfInt64(-945)))
	} else {
		return Companion_D9_.Create_DC27_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('t'))).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(false)) || (false), !(!(true) || (true))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, globalState *GlobalState) D0 {
	if (_dafny.SetOf(_dafny.IntOfInt64(30))).IsDisjointFrom(_dafny.SetOf((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qfejidul")).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_56_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), true)).Cardinality())) {
		return Companion_D0_.Create_DC2_(_dafny.MultiSetOf(false, !(!(true)), true, !(true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_57_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), true)
	} else {
		return Companion_D0_.Create_DC2_(_dafny.MultiSetOf(true, true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_58_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), false)
	}
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true)), Companion_D4_.Create_DC15_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rddv")).Cardinality()), (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-719)), _dafny.IntOfInt64(-79))).Cardinality()))).Cardinality(), true)).Merge(func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-602), _dafny.IntOfInt64(331))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _59_v0 _dafny.Int
			_59_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(-602)).Cmp(_59_v0) <= 0) && ((_59_v0).Cmp(_dafny.IntOfInt64(331)) < 0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_59_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eqcdui")).Cardinality())), true)
			}
		}
		return _coll9.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source3 D5 = Companion_D5_.Create_DC18_()
	_ = _source3
	if _source3.Is_DC17() {
		var _60___mcc_h0 _dafny.Set = _source3.Get_().(D5_DC17).Cf27
		_ = _60___mcc_h0
		var _61_cf27 _dafny.Set = _60___mcc_h0
		_ = _61_cf27
		return (Companion_D10_.Create_DC32_(_dafny.CodePoint('a'), !(!(true)))).Dtor_cf52()
	} else if _source3.Is_DC18() {
		return _dafny.CodePoint('g')
	} else if _source3.Is_DC19() {
		var _62___mcc_h1 _dafny.Int = _source3.Get_().(D5_DC19).Cf28
		_ = _62___mcc_h1
		var _63___mcc_h2 _dafny.Int = _source3.Get_().(D5_DC19).Cf29
		_ = _63___mcc_h2
		var _64___mcc_h3 _dafny.CodePoint = _source3.Get_().(D5_DC19).Cf30
		_ = _64___mcc_h3
		var _65_cf30 _dafny.CodePoint = _64___mcc_h3
		_ = _65_cf30
		var _66_cf29 _dafny.Int = _63___mcc_h2
		_ = _66_cf29
		var _67_cf28 _dafny.Int = _62___mcc_h1
		_ = _67_cf28
		return _65_cf30
	} else if _source3.Is_DC16() {
		var _68___mcc_h4 _dafny.Map = _source3.Get_().(D5_DC16).Cf26
		_ = _68___mcc_h4
		var _69_cf26 _dafny.Map = _68___mcc_h4
		_ = _69_cf26
		return _dafny.CodePoint('y')
	} else {
		var _70___mcc_h5 D5 = _source3.Get_().(D5_DC20).Cf31
		_ = _70___mcc_h5
		var _71_cf31 D5 = _70___mcc_h5
		_ = _71_cf31
		return _dafny.CodePoint('o')
	}
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) D1 {
	if !((_dafny.MultiSetOf(_dafny.IntOfInt64(564), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)).Cardinality()))), _dafny.IntOfInt64(-373))).Cardinality())).Cardinality())).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.IntOfInt64(776)))) {
		return Companion_D1_.Create_DC3_(_dafny.SeqOf(true, true))
	} else {
		return Companion_D1_.Create_DC3_(_dafny.SeqOf(true, false))
	}
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.MultiSetOf(!(false), true)).IsProperSubsetOf(_dafny.MultiSetOf(true)) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_72_i0 _dafny.Int) _dafny.Int {
			return (Companion_D19_.Create_DC59_(false, _dafny.IntOfInt64(810))).Dtor_cf109()
		})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))
	} else {
		return _dafny.SeqOf((_dafny.SetOf(!(false))).Cardinality(), _dafny.IntOfInt64(181))
	}
}
func (_static *CompanionStruct_Default___) Fm45(globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-250))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_73_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			}))).Cardinality())), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(992), (Companion_D19_.Create_DC59_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Dtor_cf109(), _dafny.IntOfInt64(175), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-566))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_74_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("druvqxil")).Cardinality())
			})))).Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _75_v1 _dafny.Sequence
				_75_v1 = interface{}(_compr_11).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-250))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}(func(_73_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality())), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(992), (Companion_D19_.Create_DC59_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Dtor_cf109(), _dafny.IntOfInt64(175), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-566))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_74_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("druvqxil")).Cardinality())
				}))), _75_v1) {
					_coll11.Add(_75_v1, _dafny.CodePoint('s'))
				}
			}
			return _coll11.ToMap()
		}()).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _76_v0 _dafny.Sequence
			_76_v0 = interface{}(_compr_10).(_dafny.Sequence)
			if (func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-250))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_73_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality())), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(992), (Companion_D19_.Create_DC59_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Dtor_cf109(), _dafny.IntOfInt64(175), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-566))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_74_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("druvqxil")).Cardinality())
				})))).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _75_v1 _dafny.Sequence
					_75_v1 = interface{}(_compr_12).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-250))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_73_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality())), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(992), (Companion_D19_.Create_DC59_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Dtor_cf109(), _dafny.IntOfInt64(175), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-566))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}(func(_74_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("druvqxil")).Cardinality())
					}))), _75_v1) {
						_coll12.Add(_75_v1, _dafny.CodePoint('s'))
					}
				}
				return _coll12.ToMap()
			}()).Contains(_76_v0) {
				_coll10.Add(_76_v0, _dafny.IntOfInt64(590))
			}
		}
		return _coll10.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_77_i2 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jqmycnx")).Cardinality()), false)).Cardinality()
		})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), false))).Keys().Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _78_v2 _dafny.Sequence
			_78_v2 = interface{}(_compr_13).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_77_i2 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jqmycnx")).Cardinality()), false)).Cardinality()
			})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), false))).Contains(_78_v2) {
				_coll13.Add(_78_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fjboup")).Cardinality()))
			}
		}
		return _coll13.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(513), _dafny.IntOfInt64(779), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("yncuqfdbc"))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ofyfibxcp")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(162), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("te")).Cardinality()), _dafny.IntOfInt64(286)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dv")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.CodePoint, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality()), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()), true)))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll14 = _dafny.NewBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, false)).Cardinality(), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("edl"), true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(118), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("tah"), true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(591))).Cardinality(), Companion_D11_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(327))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_79_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), false))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-837), Companion_D11_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_80_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), true)), func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-456), _dafny.IntOfInt64(4))); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _81_v0 _dafny.Int
				_81_v0 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(-456)).Cmp(_81_v0) <= 0) && ((_81_v0).Cmp(_dafny.IntOfInt64(4)) < 0) {
					_coll15.Add((_81_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), true)).Cardinality())).Cardinality())), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("rh"), true))
				}
			}
			return _coll15.ToMap()
		}()))).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _82_v1 _dafny.Map
			_82_v1 = interface{}(_compr_14).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, false)).Cardinality(), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("edl"), true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(118), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("tah"), true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(591))).Cardinality(), Companion_D11_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(327))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_79_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), false))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-837), Companion_D11_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_80_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), true)), func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-456), _dafny.IntOfInt64(4))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _81_v0 _dafny.Int
					_81_v0 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-456)).Cmp(_81_v0) <= 0) && ((_81_v0).Cmp(_dafny.IntOfInt64(4)) < 0) {
						_coll16.Add((_81_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), true)).Cardinality())).Cardinality())), Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("rh"), true))
					}
				}
				return _coll16.ToMap()
			}())), _82_v1) {
				_coll14.Add(_82_v1)
			}
		}
		return _coll14.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm48(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("krnlu")).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ojsoe")).Cardinality()))).Merge((func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_dafny.MultiSetOf(_dafny.SeqOf((func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-922))).Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _83_v1 _dafny.Int
				_83_v1 = interface{}(_compr_18).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-922)), _83_v1) {
					_coll18.Add((_83_v1).Plus(_dafny.IntOfInt64(631)), false)
				}
			}
			return _coll18.ToMap()
		}()).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(743)))).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfInt64(482), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhpvmy")).Cardinality()), (_dafny.SetOf(false, true)).Cardinality())).Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _84_v0 _dafny.Int
			_84_v0 = interface{}(_compr_17).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_dafny.MultiSetOf(_dafny.SeqOf((func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-922))).Elements()); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _83_v1 _dafny.Int
					_83_v1 = interface{}(_compr_19).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-922)), _83_v1) {
						_coll19.Add((_83_v1).Plus(_dafny.IntOfInt64(631)), false)
					}
				}
				return _coll19.ToMap()
			}()).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(743)))).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfInt64(482), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhpvmy")).Cardinality()), (_dafny.SetOf(false, true)).Cardinality())).Contains(_84_v0) {
				_coll17.Add((_84_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))
			}
		}
		return _coll17.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(393), _dafny.IntOfInt64(-765))).Cardinality(), (_dafny.MultiSetOf(!(!(false)), true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.CodePoint, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC32_((func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('c')
		}
		return _dafny.CodePoint('j')
	})(), !(!(true) || (false)))
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC21_((_dafny.SetOf(true, !(true))).Intersection(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm51(globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC35_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(251))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_85_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg34 _dafny.Int) interface{} {
			return coer34(arg34)
		}
	}(func(_86_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))), (Companion_D1_.Create_DC4_(false)).Dtor_cf8())
}
func (_static *CompanionStruct_Default___) Fm52(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg35 _dafny.Int) interface{} {
			return coer35(arg35)
		}
	}(func(_87_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})))
}
func (_static *CompanionStruct_Default___) Fm53(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D10_.Create_DC33_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(!(true), false)).Cardinality()))).Dtor_cf54(), Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_(!(!(false)), false)))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.Int, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) D15 {
	if false {
		return Companion_D15_.Create_DC43_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(142), false))
	} else {
		return Companion_D15_.Create_DC43_(func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-229), _dafny.IntOfInt64(503))); ; {
				_compr_20, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _88_v0 _dafny.Int
				_88_v0 = interface{}(_compr_20).(_dafny.Int)
				if ((_dafny.IntOfInt64(-229)).Cmp(_88_v0) <= 0) && ((_88_v0).Cmp(_dafny.IntOfInt64(503)) < 0) {
					_coll20.Add((_88_v0).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gslmjy")).Cardinality()))), true)
				}
			}
			return _coll20.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm55(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sltsbvpru"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bslnuaicl"), _dafny.UnicodeSeqOfUtf8Bytes("hl")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ekjt"), _dafny.UnicodeSeqOfUtf8Bytes("qhpp")))
}
func (_static *CompanionStruct_Default___) Fm56(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D21_.Create_DC65_(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg36 _dafny.Int) interface{} {
			return coer36(arg36)
		}
	}(func(_89_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-225)
	})))))).Cardinality()), true)), _dafny.SeqOf(func() _dafny.Map {
		var _coll21 = _dafny.NewMapBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-338))).Elements()); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _90_v0 _dafny.Int
			_90_v0 = interface{}(_compr_21).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-338)), _90_v0) {
				_coll21.Add((_90_v0).Times((_dafny.SetOf(_dafny.CodePoint('v'), _dafny.CodePoint('q'))).Cardinality()), false)
			}
		}
		return _coll21.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm57(p0 bool, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll22 = _dafny.NewBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_91_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Keys().Elements()); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _92_v0 _dafny.CodePoint
			_92_v0 = interface{}(_compr_22).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}(func(_91_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Contains(_92_v0) {
				_coll22.Add(_92_v0)
			}
		}
		return _coll22.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) M10(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) {
	if (func() bool {
		if !(p2) {
			return p3
		}
		return p2
	})() {
		var _93_v0 _dafny.Array
		_ = _93_v0
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw0
		_93_v0 = _nw0
		var _94_v1 _dafny.Sequence
		_ = _94_v1
		_94_v1 = _dafny.UnicodeSeqOfUtf8Bytes("brux")
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))
		_ = _index0
		(_93_v0).ArraySet1(_dafny.IntOfUint32((_94_v1).Cardinality()), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))
		_ = _index1
		(_93_v0).ArraySet1(p1, (_index1).Int())
		var _95_v2 _dafny.CodePoint
		_ = _95_v2
		_95_v2 = _dafny.CodePoint('s')
		var _96_v3 _dafny.Array
		_ = _96_v3
		var _nwElement0_0 bool = p2
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		_96_v3 = _nw1
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_96_v3), 0))
		_ = _index2
		(_96_v3).ArraySet1(p3, (_index2).Int())
		var _97_v4 _dafny.Sequence
		_ = _97_v4
		_97_v4 = _dafny.SeqOf(p3)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_96_v3), 0))
		_ = _index3
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))
		_ = _index4
		var _rhs0 _dafny.CodePoint = p0
		_ = _rhs0
		var _rhs1 bool = p3
		_ = _rhs1
		var _rhs2 bool = ((_93_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0((_93_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))).Int()).(_dafny.Int), p3, (_97_v4).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm36(globalState), _dafny.IntOfUint32((_97_v4).Cardinality()))).Uint32()).(bool), globalState), (_93_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))).Int()).(_dafny.Int))) >= 0
		_ = _rhs2
		var _rhs3 _dafny.Int = Companion_Default___.SafeDivisionInt((_93_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))).Int()).(_dafny.Int), p1)
		_ = _rhs3
		var _lhs0 _dafny.Array = _96_v3
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_96_v3), 0))
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 _dafny.Array = _93_v0
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))
		_ = _lhs4
		_95_v2 = _rhs0
		(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		_lhs2.F18 = _rhs2
		(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw2
		(globalState).F17 = _nw2
		var _98_v5 *C2
		_ = _98_v5
		var _nw3 *C2 = New_C2_()
		_ = _nw3
		_nw3.Ctor__((_93_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_93_v0), 0))).Int()).(_dafny.Int), _dafny.SetOf(_93_v0, _93_v0), (_96_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_96_v3), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_99_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_100_i0 _dafny.Int) _dafny.CodePoint {
				return _99_v2
			}
		})(_95_v2))))
		_98_v5 = _nw3
		var _101_v6 _dafny.Map
		_ = _101_v6
		_101_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v5, _96_v3)
		var _102_v7 T0
		_ = _102_v7
		var _nw4 *C1 = New_C1_()
		_ = _nw4
		_nw4.Ctor__(p2)
		_102_v7 = _nw4
		var _103_v8 *C0
		_ = _103_v8
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(false)
		_103_v8 = _nw5
		var _104_v9 _dafny.Array
		_ = _104_v9
		var _nwElement0_1 *C0 = _103_v8
		_ = _nwElement0_1
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(28))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_1, 0)
		(_nw6).ArraySet1(_103_v8, 1)
		(_nw6).ArraySet1(_103_v8, 2)
		(_nw6).ArraySet1(_103_v8, 3)
		(_nw6).ArraySet1(_103_v8, 4)
		(_nw6).ArraySet1(_103_v8, 5)
		(_nw6).ArraySet1(_103_v8, 6)
		(_nw6).ArraySet1(_103_v8, 7)
		(_nw6).ArraySet1(_103_v8, 8)
		(_nw6).ArraySet1(_103_v8, 9)
		(_nw6).ArraySet1(_103_v8, 10)
		(_nw6).ArraySet1(_103_v8, 11)
		(_nw6).ArraySet1(_103_v8, 12)
		(_nw6).ArraySet1(_103_v8, 13)
		(_nw6).ArraySet1(_103_v8, 14)
		(_nw6).ArraySet1(_103_v8, 15)
		(_nw6).ArraySet1(_103_v8, 16)
		(_nw6).ArraySet1(_103_v8, 17)
		(_nw6).ArraySet1(_103_v8, 18)
		(_nw6).ArraySet1(_103_v8, 19)
		(_nw6).ArraySet1(_103_v8, 20)
		(_nw6).ArraySet1(_103_v8, 21)
		(_nw6).ArraySet1(_103_v8, 22)
		(_nw6).ArraySet1(_103_v8, 23)
		(_nw6).ArraySet1(_103_v8, 24)
		(_nw6).ArraySet1(_103_v8, 25)
		(_nw6).ArraySet1(_103_v8, 26)
		(_nw6).ArraySet1(_103_v8, 27)
		_104_v9 = _nw6
		var _105_v10 _dafny.Map
		_ = _105_v10
		_105_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v9, _96_v3)
		var _106_v11 _dafny.Map
		_ = _106_v11
		_106_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v7, (func() _dafny.Array {
			if (_105_v10).Contains(_104_v9) {
				return (_105_v10).Get(_104_v9).(_dafny.Array)
			}
			return _96_v3
		})())
		_101_v6 = (_101_v6).Update(_98_v5, (func() _dafny.Array {
			if (_106_v11).Contains(_102_v7) {
				return (_106_v11).Get(_102_v7).(_dafny.Array)
			}
			return _96_v3
		})())
		var _107_v12 _dafny.Map
		_ = _107_v12
		_107_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v8.F37, p3)
		_107_v12 = (_107_v12).Update(p2, p3)
	} else {
		(globalState).F18 = p2
		var _108_v13 _dafny.Array
		_ = _108_v13
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw7
		_108_v13 = _nw7
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_108_v13), 0))
		_ = _index5
		(_108_v13).ArraySet1((true) || (p3), (_index5).Int())
		var _109_v14 _dafny.Map
		_ = _109_v14
		_109_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
		var _110_v15 _dafny.Sequence
		_ = _110_v15
		_110_v15 = _dafny.UnicodeSeqOfUtf8Bytes("ofqjjeh")
		var _111_v16 _dafny.Map
		_ = _111_v16
		_111_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(864))
		var _112_v17 _dafny.Map
		_ = _112_v17
		_112_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_110_v15).Cardinality())), _dafny.IntOfInt64(-368)), (_111_v16).Contains(p1))
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_108_v13), 0))
		_ = _index6
		var _rhs4 bool = (!(_109_v14).Contains(Companion_Default___.Fm1(Companion_Default___.Fm42(p3, false, globalState), p1, p1, p3, globalState))) == ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oa")).Cardinality())).Cmp(p1) >= 0)
		_ = _rhs4
		var _rhs5 bool = (func() bool {
			if (_112_v17).Contains(p1) {
				return (_112_v17).Get(p1).(bool)
			}
			return p2
		})()
		_ = _rhs5
		var _lhs5 _dafny.Array = _108_v13
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_108_v13), 0))
		_ = _lhs6
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
		_lhs7.F15 = _rhs5
		var _113_v18 _dafny.Set
		_ = _113_v18
		_113_v18 = _dafny.SetOf(p0, p0, Companion_Default___.Fm42(Companion_Default___.Fm19(p0, globalState), p3, globalState))
		_113_v18 = Companion_Default___.Fm57(!(p2), globalState)
		var _114_v19 *C4
		_ = _114_v19
		var _nw8 *C4 = New_C4_()
		_ = _nw8
		_nw8.Ctor__((_108_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_108_v13), 0))).Int()).(bool), p1)
		_114_v19 = _nw8
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_108_v13), 0))
		_ = _index7
		(_108_v13).ArraySet1((_114_v19).F33(), (_index7).Int())
	}
	var _115_v20 _dafny.Set
	_ = _115_v20
	_115_v20 = _dafny.SetOf(p1, p1, _dafny.IntOfInt64(906))
	var _116_v21 _dafny.Sequence
	_ = _116_v21
	_116_v21 = _dafny.SeqOf((_115_v20).Cardinality())
	var _117_v22 _dafny.Set
	_ = _117_v22
	_117_v22 = _dafny.SetOf(!(p3))
	var _118_v23 _dafny.MultiSet
	_ = _118_v23
	_118_v23 = _dafny.MultiSetOf(_117_v22, _dafny.SetOf(p2))
	var _119_v24 _dafny.Array
	_ = _119_v24
	var _nwElement0_2 bool = _dafny.Companion_Sequence_.IsPrefixOf(_116_v21, _116_v21)
	_ = _nwElement0_2
	var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
	_ = _nw9
	(_nw9).ArraySet1(_nwElement0_2, 0)
	(_nw9).ArraySet1(false, 1)
	(_nw9).ArraySet1(p2, 2)
	(_nw9).ArraySet1(!(p3) || (p2), 3)
	(_nw9).ArraySet1(p2, 4)
	(_nw9).ArraySet1(p2, 5)
	(_nw9).ArraySet1(false, 6)
	(_nw9).ArraySet1((_118_v23).IsProperSubsetOf(_118_v23), 7)
	(_nw9).ArraySet1(p2, 8)
	(_nw9).ArraySet1(p3, 9)
	(_nw9).ArraySet1(!(_117_v22).Contains(p3), 10)
	(_nw9).ArraySet1(Companion_Default___.Fm1(p0, p1, _dafny.IntOfInt64(371), p2, globalState), 11)
	_119_v24 = _nw9
	for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_119_v24), 0))); ; {
		_guard_loop_0, _ok23 := _iter23()
		if !_ok23 {
			break
		}
		var _120_i1 _dafny.Int
		_120_i1 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_120_i1).Sign() != -1) && ((_120_i1).Cmp(_dafny.ArrayLen((_119_v24), 0)) < 0)) {
			(_119_v24).ArraySet1(!(p3), (_120_i1).Int())
		}
	}
	(globalState).F15 = (func() bool {
		if p3 {
			return p3
		}
		return p3
	})()
	var _121_v25 _dafny.Array
	_ = _121_v25
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
	_ = _nw10
	_121_v25 = _nw10
	_121_v25 = _121_v25
	for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_121_v25), 0))); ; {
		_guard_loop_1, _ok24 := _iter24()
		if !_ok24 {
			break
		}
		var _122_i2 _dafny.Int
		_122_i2 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_122_i2).Sign() != -1) && ((_122_i2).Cmp(_dafny.ArrayLen((_121_v25), 0)) < 0)) {
			(_121_v25).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mp"), (_122_i2).Int())
		}
	}
	var _123_v26 _dafny.Sequence
	_ = _123_v26
	_123_v26 = _dafny.SeqOf(p3, p3, p3, false, p2)
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_119_v24), 0))
	_ = _index8
	(_119_v24).ArraySet1(Companion_Default___.Fm33(p2, (_117_v22).Cardinality(), _dafny.IntOfUint32((_123_v26).Cardinality()), globalState), (_index8).Int())
	var _124_v27 T0
	_ = _124_v27
	var _nw11 *C6 = New_C6_()
	_ = _nw11
	_nw11.Ctor__(p2, _dafny.UnicodeSeqOfUtf8Bytes("bydkuf"))
	_124_v27 = _nw11
	var _125_v28 _dafny.Map
	_ = _125_v28
	_125_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _124_v27)
	var _126_v29 _dafny.Map
	_ = _126_v29
	_126_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T0 {
		if (_125_v28).Contains(p1) {
			return (_125_v28).Get(p1).(T0)
		}
		return _124_v27
	})(), p0)
	var _127_v30 _dafny.Sequence
	_ = _127_v30
	_127_v30 = _dafny.UnicodeSeqOfUtf8Bytes("juplsu")
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_119_v24), 0))
	_ = _index9
	var _rhs6 _dafny.Int = p1
	_ = _rhs6
	var _rhs7 bool = !_dafny.Companion_Sequence_.Contains(_127_v30, (func() _dafny.CodePoint {
		if (_126_v29).Contains(_124_v27) {
			return (_126_v29).Get(_124_v27).(_dafny.CodePoint)
		}
		return p0
	})())
	_ = _rhs7
	var _rhs8 _dafny.Int = p1
	_ = _rhs8
	var _lhs8 *GlobalState = globalState
	_ = _lhs8
	var _lhs9 _dafny.Array = _119_v24
	_ = _lhs9
	var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_119_v24), 0))
	_ = _lhs10
	var _lhs11 *GlobalState = globalState
	_ = _lhs11
	_lhs8.F21 = _rhs6
	(_lhs9).ArraySet1(_rhs7, (_lhs10).Int())
	_lhs11.F14 = _rhs8
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _128_v0 bool
	_ = _128_v0
	_128_v0 = true
	var _129_v1 _dafny.Set
	_ = _129_v1
	_129_v1 = _dafny.SetOf(_128_v0, _128_v0)
	var _130_v2 _dafny.Sequence
	_ = _130_v2
	_130_v2 = _dafny.SeqOf(_dafny.SetOf(_128_v0), _dafny.SetOf(_128_v0), _129_v1, _dafny.SetOf(_128_v0), _129_v1)
	var _131_v3 _dafny.Int
	_ = _131_v3
	_131_v3 = _dafny.IntOfInt64(-208)
	var _132_v4 _dafny.Set
	_ = _132_v4
	_132_v4 = _dafny.SetOf(_131_v3)
	var _133_v6 _dafny.MultiSet
	_ = _133_v6
	_133_v6 = _dafny.MultiSetOf(_dafny.IntOfInt64(137))
	var _134_v7 _dafny.Sequence
	_ = _134_v7
	_134_v7 = _dafny.SeqOf((_132_v4).Cardinality(), ((func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter25 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-496))).Uint32(), func(coer40 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_135_v3 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
			return func(_136_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_135_v3)
			}
		})(_131_v3)))).Elements()); ; {
			_compr_23, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _137_v5 _dafny.MultiSet
			_137_v5 = interface{}(_compr_23).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-496))).Uint32(), func(coer41 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_138_v3 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
				return func(_136_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_138_v3)
				}
			})(_131_v3))), _137_v5) {
				_coll23.Add(_137_v5, _dafny.SeqOf(_dafny.CodePoint('r'), _dafny.CodePoint('s')))
			}
		}
		return _coll23.ToMap()
	}()).Update(_133_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg42 _dafny.Int) interface{} {
			return coer42(arg42)
		}
	}(func(_139_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})))).Cardinality())
	var _140_v8 _dafny.Sequence
	_ = _140_v8
	_140_v8 = _dafny.UnicodeSeqOfUtf8Bytes("rqigbg")
	var _141_v9 _dafny.Array
	_ = _141_v9
	var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
	_ = _nw12
	_141_v9 = _nw12
	var _142_v10 _dafny.Map
	_ = _142_v10
	_142_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _141_v9)
	var _143_v11 _dafny.Map
	_ = _143_v11
	_143_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v3, (_140_v8).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_140_v8).Cardinality()))).Uint32()).(_dafny.CodePoint))
	var _144_v12 _dafny.CodePoint
	_ = _144_v12
	_144_v12 = _dafny.CodePoint('l')
	var _145_v13 _dafny.Map
	_ = _145_v13
	_145_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
		if (_143_v11).Contains(_131_v3) {
			return (_143_v11).Get(_131_v3).(_dafny.CodePoint)
		}
		return _144_v12
	})(), _128_v0)
	var _146_globalState *GlobalState
	_ = _146_globalState
	var _nw13 *GlobalState = New_GlobalState_()
	_ = _nw13
	_nw13.Ctor__(true, _dafny.IntOfInt64(724), _dafny.IntOfInt64(130), _130_v2, true, _dafny.MultiSetOf(_131_v3), _dafny.IntOfInt64(288), _134_v7, _dafny.IntOfInt64(643), _dafny.IntOfInt64(276), _dafny.IntOfInt64(647), _dafny.Companion_Sequence_.Concatenate(_140_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-875))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg43 _dafny.Int) interface{} {
			return coer43(arg43)
		}
	}(func(_147_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))), _dafny.IntOfInt64(741), false, _dafny.IntOfInt64(70), true, _142_v10, _141_v9, false, _dafny.IntOfInt64(723), _140_v8, _dafny.IntOfInt64(-758), _145_v13, _dafny.IntOfInt64(863), true, _140_v8, _dafny.IntOfInt64(637), false, true)
	_146_globalState = _nw13
	if !(_128_v0) {
		var _148_v14 D0
		_ = _148_v14
		_148_v14 = Companion_D0_.Create_DC0_((func() bool {
			if _128_v0 {
				return _128_v0
			}
			return false
		})())
		var _source4 D0 = _148_v14
		_ = _source4
		if _source4.Is_DC0() {
			var _149___mcc_h0 bool = _source4.Get_().(D0_DC0).Cf0
			_ = _149___mcc_h0
			var _150_cf0 bool = _149___mcc_h0
			_ = _150_cf0
			(_146_globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_140_v8, _140_v8)
			var _151_v15 *C4
			_ = _151_v15
			var _nw14 *C4 = New_C4_()
			_ = _nw14
			_nw14.Ctor__(!(_128_v0), _131_v3)
			_151_v15 = _nw14
			var _152_v16 D6
			_ = _152_v16
			_152_v16 = Companion_D6_.Create_DC21_(_129_v1)
			_152_v16 = Companion_D6_.Create_DC21_(_129_v1)
			var _153_v17 _dafny.Array
			_ = _153_v17
			var _nw15 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
			_ = _nw15
			_153_v17 = _nw15
			_153_v17 = _153_v17
		} else if _source4.Is_DC1() {
			var _154___mcc_h1 _dafny.Array = _source4.Get_().(D0_DC1).Cf1
			_ = _154___mcc_h1
			var _155___mcc_h2 _dafny.Int = _source4.Get_().(D0_DC1).Cf2
			_ = _155___mcc_h2
			var _156___mcc_h3 _dafny.Int = _source4.Get_().(D0_DC1).Cf3
			_ = _156___mcc_h3
			var _157_cf3 _dafny.Int = _156___mcc_h3
			_ = _157_cf3
			var _158_cf2 _dafny.Int = _155___mcc_h2
			_ = _158_cf2
			var _159_cf1 _dafny.Array = _154___mcc_h1
			_ = _159_cf1
			_157_cf3 = Companion_Default___.SafeDivisionInt(_157_cf3, _157_cf3)
			var _160_v18 T1
			_ = _160_v18
			var _nw16 *C6 = New_C6_()
			_ = _nw16
			_nw16.Ctor__(_128_v0, _140_v8)
			_160_v18 = _nw16
			var _161_v19 _dafny.MultiSet
			_ = _161_v19
			_161_v19 = _dafny.MultiSetOf(_160_v18, _160_v18, _160_v18, _160_v18, _160_v18)
			var _162_v20 T0
			_ = _162_v20
			var _nw17 *C1 = New_C1_()
			_ = _nw17
			_nw17.Ctor__((_dafny.MultiSetOf(_160_v18)).IsProperSubsetOf(_161_v19))
			_162_v20 = _nw17
			var _163_v21 _dafny.Sequence
			_ = _163_v21
			_163_v21 = _dafny.SeqOf(!(_129_v1).Contains(_128_v0), _128_v0, (_160_v18).F29(), false)
			var _164_v22 _dafny.Map
			_ = _164_v22
			_164_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_165_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_166_i3 _dafny.Int) _dafny.CodePoint {
					return _165_v12
				}
			})(_144_v12)))).Cardinality()), !((_160_v18).F29()))
			var _167_v23 D15
			_ = _167_v23
			_167_v23 = Companion_D15_.Create_DC43_(_164_v22)
			var _pat_let_tv0 = _164_v22
			_ = _pat_let_tv0
			var _168_v24 _dafny.Array
			_ = _168_v24
			var _nwElement0_3 D15 = Companion_D15_.Create_DC43_(_164_v22)
			_ = _nwElement0_3
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_3, 0)
			(_nw18).ArraySet1(Companion_D15_.Create_DC43_(_164_v22), 1)
			(_nw18).ArraySet1(_167_v23, 2)
			(_nw18).ArraySet1(_167_v23, 3)
			(_nw18).ArraySet1(_167_v23, 4)
			(_nw18).ArraySet1(Companion_Default___.Fm54(!((_160_v18).F29()), (Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("oegckwcwk"), _131_v3, _128_v0, _158_cf2)).Dtor_cf15(), _132_v4, (_dafny.Zero).Minus(_158_cf2), _146_globalState), 5)
			(_nw18).ArraySet1(Companion_Default___.Fm54(_128_v0, _131_v3, _132_v4, _dafny.IntOfInt64(627), _146_globalState), 6)
			(_nw18).ArraySet1(_167_v23, 7)
			(_nw18).ArraySet1(_167_v23, 8)
			(_nw18).ArraySet1(func(_pat_let0_0 D15) D15 {
				return func(_169_dt__update__tmp_h0 D15) D15 {
					return func(_pat_let1_0 _dafny.Map) D15 {
						return func(_170_dt__update_hcf70_h0 _dafny.Map) D15 {
							return Companion_D15_.Create_DC43_(_170_dt__update_hcf70_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_167_v23), 9)
			(_nw18).ArraySet1(_167_v23, 10)
			(_nw18).ArraySet1(_167_v23, 11)
			(_nw18).ArraySet1(_167_v23, 12)
			_168_v24 = _nw18
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_168_v24), 0))
			_ = _index10
			(_168_v24).ArraySet1(_167_v23, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_168_v24), 0))
			_ = _index11
			var _rhs9 T0 = _162_v20
			_ = _rhs9
			var _rhs10 _dafny.Sequence = (func() _dafny.Sequence {
				if _128_v0 {
					return _163_v21
				}
				return _163_v21
			})()
			_ = _rhs10
			var _rhs11 _dafny.MultiSet = ((func() _dafny.MultiSet {
				if _128_v0 {
					return _133_v6
				}
				return _133_v6
			})()).Difference(_133_v6)
			_ = _rhs11
			var _rhs12 D15 = _167_v23
			_ = _rhs12
			var _rhs13 _dafny.Int = (_158_cf2).Times((_129_v1).Cardinality())
			_ = _rhs13
			var _lhs12 _dafny.Array = _168_v24
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_168_v24), 0))
			_ = _lhs13
			var _lhs14 *GlobalState = _146_globalState
			_ = _lhs14
			_162_v20 = _rhs9
			_163_v21 = _rhs10
			_133_v6 = _rhs11
			(_lhs12).ArraySet1(_rhs12, (_lhs13).Int())
			_lhs14.F14 = _rhs13
			(_146_globalState).F14 = (_dafny.Zero).Minus((_158_cf2).Minus(_131_v3))
			var _171_v25 _dafny.Array
			_ = _171_v25
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw19
			_171_v25 = _nw19
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_171_v25), 0))
			_ = _index12
			(_171_v25).ArraySet1(_141_v9, (_index12).Int())
			var _172_v26 _dafny.Map
			_ = _172_v26
			_172_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _dafny.IntOfInt64(-824))
			var _173_v27 _dafny.Map
			_ = _173_v27
			_173_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_160_v18).F29(), _140_v8)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_171_v25), 0))
			_ = _index13
			var _rhs14 _dafny.Array = _141_v9
			_ = _rhs14
			var _rhs15 _dafny.Int = ((func() _dafny.Int {
				if (_172_v26).Contains(_128_v0) {
					return (_172_v26).Get(_128_v0).(_dafny.Int)
				}
				return _131_v3
			})()).Minus((_dafny.Zero).Minus(Companion_Default___.Fm36(_146_globalState)))
			_ = _rhs15
			var _rhs16 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_140_v8, (func() _dafny.Sequence {
				if (_173_v27).Contains(true) {
					return (_173_v27).Get(true).(_dafny.Sequence)
				}
				return _160_v18.F30()
			})())).Cardinality())
			_ = _rhs16
			var _lhs15 _dafny.Array = _171_v25
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_171_v25), 0))
			_ = _lhs16
			var _lhs17 *GlobalState = _146_globalState
			_ = _lhs17
			(_lhs15).ArraySet1(_rhs14, (_lhs16).Int())
			_131_v3 = _rhs15
			_lhs17.F21 = _rhs16
		} else {
			var _174___mcc_h4 _dafny.MultiSet = _source4.Get_().(D0_DC2).Cf4
			_ = _174___mcc_h4
			var _175___mcc_h5 _dafny.Sequence = _source4.Get_().(D0_DC2).Cf5
			_ = _175___mcc_h5
			var _176___mcc_h6 bool = _source4.Get_().(D0_DC2).Cf6
			_ = _176___mcc_h6
			var _177_cf6 bool = _176___mcc_h6
			_ = _177_cf6
			var _178_cf5 _dafny.Sequence = _175___mcc_h5
			_ = _178_cf5
			var _179_cf4 _dafny.MultiSet = _174___mcc_h4
			_ = _179_cf4
			(_146_globalState).F13 = ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdwnsce")).Cardinality())).Times(_131_v3)).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm44(_177_cf6, _131_v3, !(false), _177_cf6, _146_globalState)).Cardinality())) == 0
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_141_v9), 0))
			_ = _index14
			(_141_v9).ArraySet1((_dafny.Zero).Minus(_131_v3), (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_141_v9), 0))
			_ = _index15
			(_141_v9).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_131_v3, _dafny.IntOfUint32((_134_v7).Cardinality())))), (_index15).Int())
			var _180_v28 _dafny.Array
			_ = _180_v28
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw20
			_180_v28 = _nw20
			var _181_v29 _dafny.Map
			_ = _181_v29
			_181_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v28, false)
			_181_v29 = (_181_v29).Merge(_181_v29)
			(_146_globalState).F1 = _131_v3
		}
		var _182_v30 _dafny.Array
		_ = _182_v30
		var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw21
		_182_v30 = _nw21
		var _183_v31 _dafny.Array
		_ = _183_v31
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
		_ = _nw22
		_183_v31 = _nw22
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_182_v30), 0))
		_ = _index16
		(_182_v30).ArraySet1(_183_v31, (_index16).Int())
		var _184_v32 _dafny.Array
		_ = _184_v32
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_0
		var _nw23 _dafny.Array
		_ = _nw23
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw23 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Map = (func(_185_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_186_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v3, !(false))
				}
			})(_131_v3)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw23 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw23).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw23).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_184_v32 = _nw23
		var _187_v33 _dafny.Map
		_ = _187_v33
		_187_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v3, _128_v0)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_184_v32), 0))
		_ = _index17
		(_184_v32).ArraySet1(_187_v33, (_index17).Int())
		var _188_v34 _dafny.Map
		_ = _188_v34
		_188_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, (_dafny.Zero).Minus(_131_v3))
		var _189_v35 _dafny.Sequence
		_ = _189_v35
		_189_v35 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _131_v3))
		var _190_v36 _dafny.Sequence
		_ = _190_v36
		_190_v36 = _dafny.SeqOf(_188_v34, _188_v34, (_189_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.IntOfUint32((_189_v35).Cardinality()))).Uint32()).(_dafny.Map))
		var _191_v37 _dafny.Map
		_ = _191_v37
		_191_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _140_v8)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_182_v30), 0))
		_ = _index18
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_184_v32), 0))
		_ = _index19
		var _rhs17 _dafny.Array = _183_v31
		_ = _rhs17
		var _rhs18 bool = true
		_ = _rhs18
		var _rhs19 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v3, _128_v0)
		_ = _rhs19
		var _rhs20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_190_v36, _190_v36)
		_ = _rhs20
		var _rhs21 _dafny.Sequence = (func() _dafny.Sequence {
			if (_191_v37).Contains(_128_v0) {
				return (_191_v37).Get(_128_v0).(_dafny.Sequence)
			}
			return _140_v8
		})()
		_ = _rhs21
		var _lhs18 _dafny.Array = _182_v30
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_182_v30), 0))
		_ = _lhs19
		var _lhs20 *GlobalState = _146_globalState
		_ = _lhs20
		var _lhs21 _dafny.Array = _184_v32
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_184_v32), 0))
		_ = _lhs22
		var _lhs23 *GlobalState = _146_globalState
		_ = _lhs23
		(_lhs18).ArraySet1(_rhs17, (_lhs19).Int())
		_lhs20.F13 = _rhs18
		(_lhs21).ArraySet1(_rhs19, (_lhs22).Int())
		_190_v36 = _rhs20
		_lhs23.F11 = _rhs21
		var _192_v38 *C6
		_ = _192_v38
		var _nw24 *C6 = New_C6_()
		_ = _nw24
		_nw24.Ctor__(_128_v0, _dafny.Companion_Sequence_.Concatenate(_140_v8, _dafny.UnicodeSeqOfUtf8Bytes("onhi")))
		_192_v38 = _nw24
		var _nw25 *C6 = New_C6_()
		_ = _nw25
		_nw25.Ctor__(!(!(_128_v0)), _dafny.UnicodeSeqOfUtf8Bytes("aidqxkrn"))
		_192_v38 = _nw25
		(_146_globalState).F19 = _131_v3
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_141_v9), 0))
		_ = _index20
		(_141_v9).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_134_v7, _134_v7)).Cardinality()), (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_141_v9), 0))
		_ = _index21
		(_141_v9).ArraySet1(_131_v3, (_index21).Int())
	} else {
		var _193_v39 _dafny.Sequence
		_ = _193_v39
		_193_v39 = _dafny.SeqOf(_128_v0)
		var _194_v40 *C2
		_ = _194_v40
		var _nw26 *C2 = New_C2_()
		_ = _nw26
		_nw26.Ctor__(_131_v3, _dafny.SetOf(_141_v9), !((_193_v39).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_193_v39).Cardinality()))).Uint32()).(bool)), _140_v8)
		_194_v40 = _nw26
		var _195_v41 _dafny.MultiSet
		_ = _195_v41
		_195_v41 = _dafny.MultiSetOf(_140_v8, _140_v8)
		_195_v41 = Companion_Default___.Fm55(_146_globalState)
		var _196_v42 *C0
		_ = _196_v42
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__(_128_v0)
		_196_v42 = _nw27
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_141_v9), 0))
		_ = _index22
		(_141_v9).ArraySet1((_dafny.MultiSetOf(_128_v0, _196_v42.F37, _196_v42.F37)).Cardinality(), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_141_v9), 0))
		_ = _index23
		var _rhs22 _dafny.MultiSet = (Companion_Default___.Fm27((_133_v6).Cardinality(), _196_v42.F37, _146_globalState)).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-916), _131_v3))
		_ = _rhs22
		var _rhs23 _dafny.Int = (_dafny.IntOfUint32((Companion_Default___.Fm56(_146_globalState)).Cardinality())).Minus((_194_v40).F38())
		_ = _rhs23
		var _rhs24 _dafny.Int = (_194_v40).F38()
		_ = _rhs24
		var _rhs25 bool = (func() bool {
			if !(_196_v42.F37) {
				return true
			}
			return _128_v0
		})()
		_ = _rhs25
		var _lhs24 *GlobalState = _146_globalState
		_ = _lhs24
		var _lhs25 _dafny.Array = _141_v9
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_141_v9), 0))
		_ = _lhs26
		var _lhs27 *GlobalState = _146_globalState
		_ = _lhs27
		_133_v6 = _rhs22
		_lhs24.F2 = _rhs23
		(_lhs25).ArraySet1(_rhs24, (_lhs26).Int())
		_lhs27.F15 = _rhs25
		var _197_v43 _dafny.Array
		_ = _197_v43
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw28
		_197_v43 = _nw28
		var _198_v44 _dafny.Map
		_ = _198_v44
		_198_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _197_v43)
		var _199_v45 _dafny.Sequence
		_ = _199_v45
		var _200_v46 bool
		_ = _200_v46
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 bool
		_ = _out1
		_out0, _out1 = (_194_v40).M3((_dafny.SetOf(_196_v42.F37, true)).IsDisjointFrom(_129_v1), (func() _dafny.Array {
			if (_198_v44).Contains(_128_v0) {
				return (_198_v44).Get(_128_v0).(_dafny.Array)
			}
			return _197_v43
		})(), _131_v3, _146_globalState)
		_199_v45 = _out0
		_200_v46 = _out1
	}
	var _201_v47 D2
	_ = _201_v47
	_201_v47 = Companion_D2_.Create_DC8_(_144_v12, !(Companion_Default___.Fm1(_144_v12, _dafny.IntOfInt64(615), _dafny.IntOfUint32((_134_v7).Cardinality()), _128_v0, _146_globalState)))
	var _202_v48 _dafny.Sequence
	_ = _202_v48
	_202_v48 = _dafny.SeqOf(_201_v47, _201_v47, _201_v47, _201_v47, _201_v47)
	var _203_v49 _dafny.Array
	_ = _203_v49
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_1
	var _nw29 _dafny.Array
	_ = _nw29
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw29 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = func(_204_i5 _dafny.Int) bool {
			return true
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw29 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw29).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw29).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_203_v49 = _nw29
	var _205_v50 D9
	_ = _205_v50
	_205_v50 = Companion_D9_.Create_DC28_(Companion_Default___.Fm35(_202_v48, _dafny.IntOfInt64(-251), _146_globalState), _128_v0, _131_v3, (_131_v3).Cmp((Companion_D0_.Create_DC1_(_203_v49, _dafny.IntOfUint32((_140_v8).Cardinality()), _131_v3)).Dtor_cf3()) == 0, _dafny.Companion_Sequence_.IsProperPrefixOf(_134_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg45 _dafny.Int) interface{} {
			return coer45(arg45)
		}
	}((func(_206_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_207_i6 _dafny.Int) _dafny.Int {
			return _206_v3
		}
	})(_131_v3)))))
	var _source5 D9 = _205_v50
	_ = _source5
	if _source5.Is_DC28() {
		var _208___mcc_h7 _dafny.Sequence = _source5.Get_().(D9_DC28).Cf41
		_ = _208___mcc_h7
		var _209___mcc_h8 bool = _source5.Get_().(D9_DC28).Cf42
		_ = _209___mcc_h8
		var _210___mcc_h9 _dafny.Int = _source5.Get_().(D9_DC28).Cf43
		_ = _210___mcc_h9
		var _211___mcc_h10 bool = _source5.Get_().(D9_DC28).Cf44
		_ = _211___mcc_h10
		var _212___mcc_h11 bool = _source5.Get_().(D9_DC28).Cf45
		_ = _212___mcc_h11
		var _213_cf45 bool = _212___mcc_h11
		_ = _213_cf45
		var _214_cf44 bool = _211___mcc_h10
		_ = _214_cf44
		var _215_cf43 _dafny.Int = _210___mcc_h9
		_ = _215_cf43
		var _216_cf42 bool = _209___mcc_h8
		_ = _216_cf42
		var _217_cf41 _dafny.Sequence = _208___mcc_h7
		_ = _217_cf41
		var _218_v51 _dafny.Array
		_ = _218_v51
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
		_ = _nw30
		_218_v51 = _nw30
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_218_v51), 0))
		_ = _index24
		(_218_v51).ArraySet1(_133_v6, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_218_v51), 0))
		_ = _index25
		(_218_v51).ArraySet1(Companion_Default___.Fm27(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_140_v8, _140_v8)).Cardinality()), true, _146_globalState), (_index25).Int())
		(_146_globalState).F13 = _214_cf44
		_131_v3 = _215_cf43
		var _219_v52 T1
		_ = _219_v52
		var _nw31 *C7 = New_C7_()
		_ = _nw31
		_nw31.Ctor__(_213_cf45, _213_cf45, _213_cf45, _dafny.UnicodeSeqOfUtf8Bytes("fnfnj"))
		_219_v52 = _nw31
		var _220_v53 _dafny.Set
		_ = _220_v53
		_220_v53 = _dafny.SetOf(_219_v52, _219_v52, _219_v52, _219_v52, _219_v52)
		var _221_v54 _dafny.Set
		_ = _221_v54
		_221_v54 = _dafny.SetOf(_220_v53, _220_v53, _220_v53, _220_v53, _220_v53)
		var _222_v55 _dafny.Map
		_ = _222_v55
		_222_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _221_v54)
		var _223_v56 D19
		_ = _223_v56
		_223_v56 = Companion_D19_.Create_DC56_(_221_v54)
		var _rhs26 _dafny.Set = (func() _dafny.Set {
			if _214_cf44 {
				return (func() _dafny.Set {
					if (_222_v55).Contains(_213_cf45) {
						return (_222_v55).Get(_213_cf45).(_dafny.Set)
					}
					return _dafny.SetOf(_dafny.SetOf(_219_v52))
				})()
			}
			return (_223_v56).Dtor_cf99()
		})()
		_ = _rhs26
		var _rhs27 _dafny.Int = (((_dafny.SetOf(true, _128_v0, _216_cf42)).Difference(_dafny.SetOf(_216_cf42, _128_v0, _214_cf44))).Cardinality()).Times(_215_cf43)
		_ = _rhs27
		var _lhs28 *GlobalState = _146_globalState
		_ = _lhs28
		_221_v54 = _rhs26
		_lhs28.F14 = _rhs27
	} else if _source5.Is_DC29() {
		var _224___mcc_h12 _dafny.Int = _source5.Get_().(D9_DC29).Cf46
		_ = _224___mcc_h12
		var _225___mcc_h13 _dafny.Int = _source5.Get_().(D9_DC29).Cf47
		_ = _225___mcc_h13
		var _226___mcc_h14 bool = _source5.Get_().(D9_DC29).Cf48
		_ = _226___mcc_h14
		var _227_cf48 bool = _226___mcc_h14
		_ = _227_cf48
		var _228_cf47 _dafny.Int = _225___mcc_h13
		_ = _228_cf47
		var _229_cf46 _dafny.Int = _224___mcc_h12
		_ = _229_cf46
		Companion_Default___.M10(_dafny.CodePoint('x'), _229_cf46, (func() bool {
			if true {
				return _227_cf48
			}
			return _227_cf48
		})(), (_131_v3).Cmp(_dafny.IntOfInt64(459)) <= 0, _146_globalState)
		var _230_v57 *C7
		_ = _230_v57
		var _nw32 *C7 = New_C7_()
		_ = _nw32
		_nw32.Ctor__(_128_v0, _128_v0, _128_v0, _140_v8)
		_230_v57 = _nw32
		var _231_v58 _dafny.Set
		_ = _231_v58
		_231_v58 = _dafny.SetOf(_230_v57, _230_v57)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_141_v9), 0))
		_ = _index26
		(_141_v9).ArraySet1(Companion_Default___.Fm0((_231_v58).Cardinality(), (_230_v57).F31(), _227_cf48, _146_globalState), (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_141_v9), 0))
		_ = _index27
		(_141_v9).ArraySet1(_131_v3, (_index27).Int())
		var _232_v59 _dafny.Map
		_ = _232_v59
		_232_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v7, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wegckpd")).Cardinality()))
		var _233_v60 D11
		_ = _233_v60
		_233_v60 = Companion_D11_.Create_DC34_(_232_v59)
		_233_v60 = _233_v60
		(_146_globalState).F21 = (_141_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_141_v9), 0))).Int()).(_dafny.Int)
	} else {
		var _234___mcc_h15 _dafny.Sequence = _source5.Get_().(D9_DC27).Cf40
		_ = _234___mcc_h15
		var _235_cf40 _dafny.Sequence = _234___mcc_h15
		_ = _235_cf40
		if _128_v0 {
			var _236_v61 _dafny.Array
			_ = _236_v61
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_2
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_237_v8 _dafny.Sequence, _238_v3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_239_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_237_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg46 _dafny.Int) interface{} {
								return coer46(arg46)
							}
						}((func(_240_v8 _dafny.Sequence, _241_v3 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
							return func(_242_i8 _dafny.Int) _dafny.CodePoint {
								return (_240_v8).Select((Companion_Default___.SafeIndex(_241_v3, _dafny.IntOfUint32((_240_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
							}
						})(_237_v8, _238_v3))))
					}
				})(_140_v8, _131_v3)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw33 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw33).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw33).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_236_v61 = _nw33
			var _243_v62 *C6
			_ = _243_v62
			var _nw34 *C6 = New_C6_()
			_ = _nw34
			_nw34.Ctor__(_128_v0, _140_v8)
			_243_v62 = _nw34
			var _244_v63 _dafny.Map
			_ = _244_v63
			_244_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v62, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_245_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_246_i9 _dafny.Int) _dafny.CodePoint {
					return _245_v12
				}
			})(_144_v12))))
			var _247_v64 _dafny.Sequence
			_ = _247_v64
			_247_v64 = _dafny.SeqOf(_243_v62, _243_v62)
			var _248_v65 _dafny.Sequence
			_ = _248_v65
			_248_v65 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_244_v63).Contains((_247_v64).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_247_v64).Cardinality()))).Uint32()).(*C6)) {
					return (_244_v63).Get((_247_v64).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_247_v64).Cardinality()))).Uint32()).(*C6)).(_dafny.Sequence)
				}
				return _140_v8
			})(), _140_v8, _140_v8, _dafny.UnicodeSeqOfUtf8Bytes("elrvfpxmc"))
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_236_v61), 0))
			_ = _index28
			(_236_v61).ArraySet1(_248_v65, (_index28).Int())
			var _249_v66 *C0
			_ = _249_v66
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_128_v0)
			_249_v66 = _nw35
			var _250_v67 _dafny.Sequence
			_ = _250_v67
			_250_v67 = _dafny.SeqOf(_249_v66, _249_v66, _249_v66, _249_v66, _249_v66)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_236_v61), 0))
			_ = _index29
			var _rhs28 _dafny.Int = _131_v3
			_ = _rhs28
			var _rhs29 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_140_v8, _140_v8), _140_v8)
			_ = _rhs29
			var _rhs30 _dafny.Sequence = _248_v65
			_ = _rhs30
			var _rhs31 *C0 = (_250_v67).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_250_v67).Cardinality()))).Uint32()).(*C0)
			_ = _rhs31
			var _lhs29 *GlobalState = _146_globalState
			_ = _lhs29
			var _lhs30 _dafny.Array = _236_v61
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_236_v61), 0))
			_ = _lhs31
			_131_v3 = _rhs28
			_lhs29.F20 = _rhs29
			(_lhs30).ArraySet1(_rhs30, (_lhs31).Int())
			_249_v66 = _rhs31
			var _251_v68 D0
			_ = _251_v68
			_251_v68 = Companion_D0_.Create_DC1_(_203_v49, _dafny.IntOfInt64(-974), _131_v3)
			_251_v68 = _251_v68
			var _252_v69 _dafny.Map
			_ = _252_v69
			_252_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_140_v8).Cardinality()))
			var _253_v70 *C3
			_ = _253_v70
			var _nw36 *C3 = New_C3_()
			_ = _nw36
			_nw36.Ctor__((Companion_Default___.Fm9(_146_globalState)).Cmp(_131_v3) <= 0, (_dafny.IntOfInt64(579)).Minus((func() _dafny.Int {
				if (_252_v69).Contains(_249_v66.F37) {
					return (_252_v69).Get(_249_v66.F37).(_dafny.Int)
				}
				return _131_v3
			})()), _249_v66.F37, _dafny.Companion_Sequence_.Concatenate(_140_v8, _dafny.UnicodeSeqOfUtf8Bytes("ogivdrrsy")))
			_253_v70 = _nw36
			(_146_globalState).F2 = _dafny.IntOfInt64(750)
			var _254_v71 _dafny.Array
			_ = _254_v71
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw37
			_254_v71 = _nw37
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_254_v71), 0))
			_ = _index30
			(_254_v71).ArraySet1(_140_v8, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_254_v71), 0))
			_ = _index31
			(_254_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_140_v8, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ohllvco"), (Companion_Default___.SafeIndex((_253_v70).F36(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ohllvco")).Cardinality()))).Uint32(), _144_v12)), (_index31).Int())
		} else {
			(_146_globalState).F7 = _235_cf40
			var _255_v72 _dafny.Map
			_ = _255_v72
			_255_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _144_v12)
			var _256_v73 _dafny.Map
			_ = _256_v73
			_256_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_140_v8).Cardinality()), _128_v0)
			var _257_v74 T0
			_ = _257_v74
			var _nw38 *C6 = New_C6_()
			_ = _nw38
			_nw38.Ctor__(_128_v0, _140_v8)
			_257_v74 = _nw38
			var _258_v75 _dafny.Map
			_ = _258_v75
			_258_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _257_v74)
			var _259_v76 _dafny.Array
			_ = _259_v76
			var _nwElement0_4 _dafny.CodePoint = _144_v12
			_ = _nwElement0_4
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(25))
			_ = _nw39
			(_nw39).ArraySet1CodePoint(_nwElement0_4, 0)
			(_nw39).ArraySet1CodePoint(_144_v12, 1)
			(_nw39).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_255_v72).Contains(_128_v0) {
					return (_255_v72).Get(_128_v0).(_dafny.CodePoint)
				}
				return _144_v12
			})(), 2)
			(_nw39).ArraySet1CodePoint(_144_v12, 3)
			(_nw39).ArraySet1CodePoint(_144_v12, 4)
			(_nw39).ArraySet1CodePoint(_144_v12, 5)
			(_nw39).ArraySet1CodePoint(_144_v12, 6)
			(_nw39).ArraySet1CodePoint(Companion_Default___.Fm42(_128_v0, _128_v0, _146_globalState), 7)
			(_nw39).ArraySet1CodePoint(_144_v12, 8)
			(_nw39).ArraySet1CodePoint(Companion_Default___.Fm42((func() bool {
				if (_256_v73).Contains((_258_v75).Cardinality()) {
					return (_256_v73).Get((_258_v75).Cardinality()).(bool)
				}
				return false
			})(), false, _146_globalState), 9)
			(_nw39).ArraySet1CodePoint(_144_v12, 10)
			(_nw39).ArraySet1CodePoint(_144_v12, 11)
			(_nw39).ArraySet1CodePoint(_144_v12, 12)
			(_nw39).ArraySet1CodePoint(_dafny.CodePoint('x'), 13)
			(_nw39).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _128_v0 {
					return _144_v12
				}
				return _144_v12
			})(), 14)
			(_nw39).ArraySet1CodePoint(_144_v12, 15)
			(_nw39).ArraySet1CodePoint(_144_v12, 16)
			(_nw39).ArraySet1CodePoint(_144_v12, 17)
			(_nw39).ArraySet1CodePoint(Companion_Default___.Fm42(_128_v0, _128_v0, _146_globalState), 18)
			(_nw39).ArraySet1CodePoint(_144_v12, 19)
			(_nw39).ArraySet1CodePoint(_144_v12, 20)
			(_nw39).ArraySet1CodePoint(Companion_Default___.Fm42(_128_v0, _128_v0, _146_globalState), 21)
			(_nw39).ArraySet1CodePoint(_144_v12, 22)
			(_nw39).ArraySet1CodePoint(_144_v12, 23)
			(_nw39).ArraySet1CodePoint(_144_v12, 24)
			_259_v76 = _nw39
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_259_v76), 0))
			_ = _index32
			(_259_v76).ArraySet1CodePoint(Companion_Default___.Fm42(false, _128_v0, _146_globalState), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_203_v49), 0))
			_ = _index33
			(_203_v49).ArraySet1(_128_v0, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_259_v76), 0))
			_ = _index34
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_203_v49), 0))
			_ = _index35
			var _rhs32 bool = (_131_v3).Cmp(_131_v3) != 0
			_ = _rhs32
			var _rhs33 _dafny.CodePoint = _dafny.CodePoint('j')
			_ = _rhs33
			var _rhs34 _dafny.Int = (_dafny.Zero).Minus(_131_v3)
			_ = _rhs34
			var _rhs35 bool = _128_v0
			_ = _rhs35
			var _rhs36 bool = _128_v0
			_ = _rhs36
			var _lhs32 *GlobalState = _146_globalState
			_ = _lhs32
			var _lhs33 _dafny.Array = _259_v76
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_259_v76), 0))
			_ = _lhs34
			var _lhs35 *GlobalState = _146_globalState
			_ = _lhs35
			var _lhs36 _dafny.Array = _203_v49
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_203_v49), 0))
			_ = _lhs37
			var _lhs38 *GlobalState = _146_globalState
			_ = _lhs38
			_lhs32.F15 = _rhs32
			(_lhs33).ArraySet1CodePoint(_rhs33, (_lhs34).Int())
			_lhs35.F2 = _rhs34
			(_lhs36).ArraySet1(_rhs35, (_lhs37).Int())
			_lhs38.F15 = _rhs36
			var _260_v77 _dafny.Array
			_ = _260_v77
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw40
			_260_v77 = _nw40
			var _261_v78 _dafny.Map
			_ = _261_v78
			_261_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_235_cf40, (Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_235_cf40).Cardinality()))).Uint32(), _131_v3)).Cardinality()), _260_v77)
			_261_v78 = (_261_v78).Update(_dafny.IntOfInt64(777), _260_v77)
			var _262_v79 *C5
			_ = _262_v79
			var _nw41 *C5 = New_C5_()
			_ = _nw41
			_nw41.Ctor__(_128_v0, _140_v8)
			_262_v79 = _nw41
			var _263_v80 _dafny.MultiSet
			_ = _263_v80
			_263_v80 = _dafny.MultiSetOf((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool))
			(_146_globalState).F13 = (func() bool {
				if (_145_v13).Contains(_144_v12) {
					return (_145_v13).Get(_144_v12).(bool)
				}
				return ((_263_v80).Update(_128_v0, Companion_Default___.Abs(_131_v3))).IsSubsetOf(_dafny.MultiSetOf(_128_v0, _128_v0, _128_v0))
			})()
		}
		Companion_Default___.M10(_144_v12, _131_v3, (_dafny.IntOfInt64(-976)).Cmp(Companion_Default___.Fm0(_131_v3, _128_v0, _128_v0, _146_globalState)) >= 0, _128_v0, _146_globalState)
		var _264_v81 _dafny.Map
		_ = _264_v81
		_264_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _dafny.IntOfInt64(241))
		var _265_v82 _dafny.Sequence
		_ = _265_v82
		_265_v82 = _dafny.SeqOf(_203_v49, _203_v49, _203_v49)
		var _rhs37 _dafny.Array = (_265_v82).Select((Companion_Default___.SafeIndex(_131_v3, _dafny.IntOfUint32((_265_v82).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs37
		var _rhs38 _dafny.CodePoint = _144_v12
		_ = _rhs38
		var _rhs39 _dafny.Map = _264_v81
		_ = _rhs39
		_203_v49 = _rhs37
		_144_v12 = _rhs38
		_264_v81 = _rhs39
		Companion_Default___.M10(_144_v12, (_131_v3).Minus((_dafny.Zero).Minus(_131_v3)), !(_132_v4).Contains(_131_v3), _128_v0, _146_globalState)
	}
	(_146_globalState).F14 = (_131_v3).Times((_131_v3).Times(_dafny.IntOfInt64(302)))
	if (_131_v3).Cmp(_131_v3) == 0 {
		_131_v3 = _131_v3
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_141_v9), 0))
		_ = _index36
		(_141_v9).ArraySet1(_131_v3, (_index36).Int())
		var _266_v83 _dafny.Map
		_ = _266_v83
		_266_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v3, _dafny.IntOfInt64(121))
		var _267_v84 D10
		_ = _267_v84
		_267_v84 = Companion_D10_.Create_DC32_(_dafny.CodePoint('k'), _128_v0)
		var _pat_let_tv1 = _128_v0
		_ = _pat_let_tv1
		var _268_v85 _dafny.Array
		_ = _268_v85
		var _nwElement0_5 D10 = _267_v84
		_ = _nwElement0_5
		var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(12))
		_ = _nw42
		(_nw42).ArraySet1(_nwElement0_5, 0)
		(_nw42).ArraySet1(_267_v84, 1)
		(_nw42).ArraySet1(_267_v84, 2)
		(_nw42).ArraySet1(Companion_D10_.Create_DC32_(_144_v12, _128_v0), 3)
		(_nw42).ArraySet1(func(_pat_let2_0 D10) D10 {
			return func(_269_dt__update__tmp_h1 D10) D10 {
				return func(_pat_let3_0 bool) D10 {
					return func(_270_dt__update_hcf53_h0 bool) D10 {
						return Companion_D10_.Create_DC32_((_269_dt__update__tmp_h1).Dtor_cf52(), _270_dt__update_hcf53_h0)
					}(_pat_let3_0)
				}(_pat_let_tv1)
			}(_pat_let2_0)
		}(_267_v84), 4)
		(_nw42).ArraySet1(_267_v84, 5)
		(_nw42).ArraySet1(Companion_D10_.Create_DC32_(_144_v12, _128_v0), 6)
		(_nw42).ArraySet1(_267_v84, 7)
		(_nw42).ArraySet1(Companion_D10_.Create_DC32_(_144_v12, _128_v0), 8)
		(_nw42).ArraySet1(_267_v84, 9)
		(_nw42).ArraySet1(Companion_D10_.Create_DC32_(_144_v12, _128_v0), 10)
		(_nw42).ArraySet1(_267_v84, 11)
		_268_v85 = _nw42
		var _271_v86 *C4
		_ = _271_v86
		var _nw43 *C4 = New_C4_()
		_ = _nw43
		_nw43.Ctor__(!(_128_v0), (Companion_D15_.Create_DC44_(_131_v3, _128_v0, _268_v85, _141_v9, (_dafny.Zero).Minus(_131_v3))).Dtor_cf71())
		_271_v86 = _nw43
		var _272_v87 _dafny.Map
		_ = _272_v87
		_272_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_271_v86, _128_v0)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_141_v9), 0))
		_ = _index37
		(_141_v9).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
			if (_266_v83).Contains(_131_v3) {
				return (_266_v83).Get(_131_v3).(_dafny.Int)
			}
			return (_129_v1).Cardinality()
		})())).Minus((_272_v87).Cardinality()), (_index37).Int())
		var _273_v88 _dafny.Array
		_ = _273_v88
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw44
		_273_v88 = _nw44
		_273_v88 = _273_v88
		var _274_v89 _dafny.Sequence
		_ = _274_v89
		var _275_v90 bool
		_ = _275_v90
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 bool
		_ = _out3
		_out2, _out3 = (_271_v86).M2(_128_v0, _146_globalState)
		_274_v89 = _out2
		_275_v90 = _out3
		var _276_v91 *C6
		_ = _276_v91
		var _nw45 *C6 = New_C6_()
		_ = _nw45
		_nw45.Ctor__((_271_v86).F33(), _140_v8)
		_276_v91 = _nw45
		_276_v91 = _276_v91
	} else {
		(_146_globalState).F15 = (_128_v0) == (_128_v0)
		var _277_v92 _dafny.Sequence
		_ = _277_v92
		_277_v92 = _dafny.SeqOf((_131_v3).Cmp(_dafny.IntOfInt64(-811)) != 0)
		_131_v3 = _dafny.IntOfUint32((_277_v92).Cardinality())
		_144_v12 = _144_v12
		if _128_v0 {
			(_146_globalState).F1 = (func() _dafny.Int {
				if (_133_v6).Contains(Companion_Default___.SafeModuloInt(_131_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(394))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_278_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_279_i10 _dafny.Int) _dafny.CodePoint {
						return _278_v12
					}
				})(_144_v12)))).Cardinality()))) {
					return (_133_v6).Multiplicity(Companion_Default___.SafeModuloInt(_131_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(394))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_280_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_281_i10 _dafny.Int) _dafny.CodePoint {
							return _280_v12
						}
					})(_144_v12)))).Cardinality())))
				}
				return _dafny.IntOfInt64(660)
			})()
			var _282_v93 *C5
			_ = _282_v93
			var _nw46 *C5 = New_C5_()
			_ = _nw46
			_nw46.Ctor__(_128_v0, _140_v8)
			_282_v93 = _nw46
			var _283_v94 *C7
			_ = _283_v94
			var _nw47 *C7 = New_C7_()
			_ = _nw47
			_nw47.Ctor__(_128_v0, _128_v0, _128_v0, _140_v8)
			_283_v94 = _nw47
			(_146_globalState).F7 = _134_v7
			var _rhs40 _dafny.Int = (_131_v3).Plus((func() _dafny.Int {
				if (_133_v6).Contains(_131_v3) {
					return (_133_v6).Multiplicity(_131_v3)
				}
				return _131_v3
			})())
			_ = _rhs40
			var _rhs41 _dafny.Int = _131_v3
			_ = _rhs41
			var _lhs39 *GlobalState = _146_globalState
			_ = _lhs39
			var _lhs40 *GlobalState = _146_globalState
			_ = _lhs40
			_lhs39.F19 = _rhs40
			_lhs40.F14 = _rhs41
		} else {
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))
			_ = _index38
			(_203_v49).ArraySet1(Companion_Default___.Fm6(_131_v3, _146_globalState), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))
			_ = _index39
			(_203_v49).ArraySet1(!(_128_v0) || (_128_v0), (_index39).Int())
			var _284_v95 _dafny.Array
			_ = _284_v95
			var _nwElement0_6 bool = _128_v0
			_ = _nwElement0_6
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(5))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_6, 0)
			(_nw48).ArraySet1((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), 1)
			(_nw48).ArraySet1((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), 2)
			(_nw48).ArraySet1(true, 3)
			(_nw48).ArraySet1(!(_128_v0), 4)
			_284_v95 = _nw48
			var _285_v96 _dafny.Sequence
			_ = _285_v96
			_285_v96 = _dafny.SeqOf(_284_v95, _284_v95)
			(_146_globalState).F28 = _dafny.Companion_Sequence_.Contains(_285_v96, _284_v95)
			var _286_v97 _dafny.Map
			_ = _286_v97
			_286_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(825), _128_v0)
			Companion_Default___.M10((func() _dafny.CodePoint {
				if Companion_Default___.Fm19(_144_v12, _146_globalState) {
					return _144_v12
				}
				return _144_v12
			})(), _131_v3, Companion_Default___.Fm33((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), _131_v3, (_286_v97).Cardinality(), _146_globalState), (Companion_Default___.Fm36(_146_globalState)).Cmp(_131_v3) < 0, _146_globalState)
			var _287_v98 _dafny.Set
			_ = _287_v98
			_287_v98 = _dafny.SetOf(_141_v9)
			var _288_v99 *C2
			_ = _288_v99
			var _nw49 *C2 = New_C2_()
			_ = _nw49
			_nw49.Ctor__(_131_v3, _287_v98, _128_v0, _140_v8)
			_288_v99 = _nw49
			var _289_v100 _dafny.Set
			_ = _289_v100
			_289_v100 = _dafny.SetOf(_288_v99, _288_v99)
			var _290_v101 _dafny.Map
			_ = _290_v101
			_290_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool) {
					return _140_v8
				}
				return Companion_Default___.Fm20(_128_v0, (_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), _277_v92, (_289_v100).Cardinality(), _146_globalState)
			})(), _134_v7)
			_290_v101 = (_290_v101).Update(_dafny.Companion_Sequence_.Concatenate(_140_v8, _140_v8), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm44((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), (func() _dafny.Int {
				if (_133_v6).Contains((_288_v99).F38()) {
					return (_133_v6).Multiplicity((_288_v99).F38())
				}
				return _131_v3
			})(), _128_v0, _128_v0, _146_globalState), _134_v7))
			var _291_v102 _dafny.Array
			_ = _291_v102
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_3
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = func(_292_i12 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("ughfu")
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw50 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw50).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw50).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_291_v102 = _nw50
			var _293_v103 _dafny.Sequence
			_ = _293_v103
			var _294_v104 bool
			_ = _294_v104
			var _out4 _dafny.Sequence
			_ = _out4
			var _out5 bool
			_ = _out5
			_out4, _out5 = (_288_v99).M3(Companion_Default___.Fm6(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_295_v8 _dafny.Sequence, _296_v99 *C2) func(_dafny.Int) _dafny.CodePoint {
				return func(_297_i11 _dafny.Int) _dafny.CodePoint {
					return (_295_v8).Select((Companion_Default___.SafeIndex((_296_v99).F38(), _dafny.IntOfUint32((_295_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_140_v8, _288_v99)))).Cardinality()), _146_globalState), _291_v102, (_131_v3).Plus((_288_v99).F38()), _146_globalState)
			_293_v103 = _out4
			_294_v104 = _out5
		}
		(_146_globalState).F13 = _128_v0
	}
	var _298_v105 _dafny.Array
	_ = _298_v105
	var _nw51 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
	_ = _nw51
	_298_v105 = _nw51
	var _299_v106 T1
	_ = _299_v106
	var _nw52 *C7 = New_C7_()
	_ = _nw52
	_nw52.Ctor__(_128_v0, _128_v0, _128_v0, _140_v8)
	_299_v106 = _nw52
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_298_v105), 0))
	_ = _index40
	(_298_v105).ArraySet1(_299_v106, (_index40).Int())
	var _300_v107 _dafny.Sequence
	_ = _300_v107
	_300_v107 = _dafny.SeqOf(_299_v106)
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_298_v105), 0))
	_ = _index41
	(_298_v105).ArraySet1((_300_v107).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm9(_146_globalState), _dafny.IntOfUint32((_300_v107).Cardinality()))).Uint32()).(T1), (_index41).Int())
	_203_v49 = _203_v49
	var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))
	_ = _index42
	(_203_v49).ArraySet1(_128_v0, (_index42).Int())
	var _301_v108 _dafny.Sequence
	_ = _301_v108
	_301_v108 = _dafny.SeqOf(_128_v0, (_299_v106).F29(), _128_v0)
	var _302_v109 _dafny.Map
	_ = _302_v109
	_302_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_144_v12, _dafny.IntOfUint32((_301_v108).Cardinality()), _131_v3, true, _146_globalState), (_299_v106).F29())
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))
	_ = _index43
	(_203_v49).ArraySet1(!((func() bool {
		if (_302_v109).Contains((_299_v106).F29()) {
			return (_302_v109).Get((_299_v106).F29()).(bool)
		}
		return (_299_v106).F29()
	})()), (_index43).Int())
	(_146_globalState).F28 = (_299_v106).F29()
	var _303_v110 _dafny.Array
	_ = _303_v110
	var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
	_ = _nw53
	_303_v110 = _nw53
	var _304_v111 _dafny.Map
	_ = _304_v111
	_304_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _131_v3)
	var _305_v112 *C5
	_ = _305_v112
	var _nw54 *C5 = New_C5_()
	_ = _nw54
	_nw54.Ctor__(((_dafny.Zero).Minus(((_304_v111).Update(_128_v0, _131_v3)).Cardinality())).Cmp(_131_v3) <= 0, _140_v8)
	_305_v112 = _nw54
	var _306_v114 D17
	_ = _306_v114
	_306_v114 = Companion_D17_.Create_DC52_((_299_v106).F29(), (_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), _131_v3, _133_v6)
	var _pat_let_tv2 = _131_v3
	_ = _pat_let_tv2
	var _pat_let_tv3 = _299_v106
	_ = _pat_let_tv3
	var _pat_let_tv4 = _131_v3
	_ = _pat_let_tv4
	var _pat_let_tv5 = _304_v111
	_ = _pat_let_tv5
	var _pat_let_tv6 = _131_v3
	_ = _pat_let_tv6
	var _pat_let_tv7 = _301_v108
	_ = _pat_let_tv7
	var _pat_let_tv8 = _301_v108
	_ = _pat_let_tv8
	var _rhs42 _dafny.Array = _303_v110
	_ = _rhs42
	var _rhs43 bool = !(((_132_v4).Intersection(func() _dafny.Set {
		var _coll24 = _dafny.NewBuilder()
		_ = _coll24
		for _iter26 := _dafny.Iterate((_134_v7).Elements()); ; {
			_compr_24, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _307_v113 _dafny.Int
			_307_v113 = interface{}(_compr_24).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_134_v7, _307_v113) {
				_coll24.Add((_307_v113).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jm")).Cardinality())))
			}
		}
		return _coll24.ToSet()
	}())).IsSubsetOf(_dafny.SetOf(_131_v3, (_dafny.SetOf(true)).Cardinality())))
	_ = _rhs43
	var _rhs44 bool = (_299_v106).F29()
	_ = _rhs44
	var _rhs45 _dafny.Int = func(_source6 D17) _dafny.Int {
		if _source6.Is_DC50() {
			var _308___mcc_h16 bool = _source6.Get_().(D17_DC50).Cf83
			_ = _308___mcc_h16
			var _309___mcc_h17 bool = _source6.Get_().(D17_DC50).Cf84
			_ = _309___mcc_h17
			var _310___mcc_h18 T1 = _source6.Get_().(D17_DC50).Cf85
			_ = _310___mcc_h18
			var _311_cf85 T1 = _310___mcc_h18
			_ = _311_cf85
			var _312_cf84 bool = _309___mcc_h17
			_ = _312_cf84
			var _313_cf83 bool = _308___mcc_h16
			_ = _313_cf83
			return Companion_Default___.SafeDivisionInt(_pat_let_tv2, _dafny.IntOfUint32((_pat_let_tv3.F30()).Cardinality()))
		} else if _source6.Is_DC51() {
			var _314___mcc_h19 _dafny.Int = _source6.Get_().(D17_DC51).Cf86
			_ = _314___mcc_h19
			var _315___mcc_h20 _dafny.Array = _source6.Get_().(D17_DC51).Cf87
			_ = _315___mcc_h20
			var _316_cf87 _dafny.Array = _315___mcc_h20
			_ = _316_cf87
			var _317_cf86 _dafny.Int = _314___mcc_h19
			_ = _317_cf86
			return _317_cf86
		} else if _source6.Is_DC52() {
			var _318___mcc_h21 bool = _source6.Get_().(D17_DC52).Cf88
			_ = _318___mcc_h21
			var _319___mcc_h22 bool = _source6.Get_().(D17_DC52).Cf89
			_ = _319___mcc_h22
			var _320___mcc_h23 _dafny.Int = _source6.Get_().(D17_DC52).Cf90
			_ = _320___mcc_h23
			var _321___mcc_h24 _dafny.MultiSet = _source6.Get_().(D17_DC52).Cf91
			_ = _321___mcc_h24
			var _322_cf91 _dafny.MultiSet = _321___mcc_h24
			_ = _322_cf91
			var _323_cf90 _dafny.Int = _320___mcc_h23
			_ = _323_cf90
			var _324_cf89 bool = _319___mcc_h22
			_ = _324_cf89
			var _325_cf88 bool = _318___mcc_h21
			_ = _325_cf88
			return (_pat_let_tv4).Times((_pat_let_tv5).Cardinality())
		} else if _source6.Is_DC49() {
			var _326___mcc_h25 _dafny.Set = _source6.Get_().(D17_DC49).Cf82
			_ = _326___mcc_h25
			var _327_cf82 _dafny.Set = _326___mcc_h25
			_ = _327_cf82
			return (_dafny.Zero).Minus(_pat_let_tv6)
		} else {
			var _328___mcc_h26 D17 = _source6.Get_().(D17_DC53).Cf92
			_ = _328___mcc_h26
			var _329_cf92 D17 = _328___mcc_h26
			_ = _329_cf92
			return (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if true {
					return _pat_let_tv7
				}
				return _pat_let_tv8
			})())).Cardinality()
		}
	}(_306_v114)
	_ = _rhs45
	var _rhs46 *C5 = (func() *C5 {
		if !(_128_v0) || ((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool)) {
			return (func() *C5 {
				if (_299_v106).F29() {
					return _305_v112
				}
				return _305_v112
			})()
		}
		return _305_v112
	})()
	_ = _rhs46
	var _lhs41 *GlobalState = _146_globalState
	_ = _lhs41
	var _lhs42 *GlobalState = _146_globalState
	_ = _lhs42
	_303_v110 = _rhs42
	_128_v0 = _rhs43
	_lhs41.F18 = _rhs44
	_lhs42.F19 = _rhs45
	_305_v112 = _rhs46
	var _330_v115 _dafny.Array
	_ = _330_v115
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_4
	var _nw55 _dafny.Array
	_ = _nw55
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw55 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Sequence = (func(_331_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_332_i13 _dafny.Int) _dafny.Sequence {
				return _331_v8
			}
		})(_140_v8)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw55 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw55).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw55).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_330_v115 = _nw55
	var _333_v116 *C7
	_ = _333_v116
	var _nw56 *C7 = New_C7_()
	_ = _nw56
	_nw56.Ctor__((_299_v106).F29(), (_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool), (_299_v106).F29(), _299_v106.F30())
	_333_v116 = _nw56
	var _334_v117 D20
	_ = _334_v117
	_334_v117 = Companion_D20_.Create_DC61_(_dafny.SeqOf(_333_v116))
	var _335_v118 _dafny.Sequence
	_ = _335_v118
	var _336_v119 bool
	_ = _336_v119
	var _out6 _dafny.Sequence
	_ = _out6
	var _out7 bool
	_ = _out7
	_out6, _out7 = (_305_v112).M3((_299_v106).F29(), _330_v115, (_dafny.IntOfUint32(((_334_v117).Dtor_cf111()).Cardinality())).Minus(Companion_Default___.Fm36(_146_globalState)), _146_globalState)
	_335_v118 = _out6
	_336_v119 = _out7
	var _337_i14 _dafny.Int
	_ = _337_i14
	_337_i14 = _dafny.Zero
	{
		for ((_305_v112).Fm4(_146_globalState)).Cmp((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_333_v116).F31(), _131_v3))).Cardinality()) >= 0 {
			{
				if (_337_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_337_i14 = (_337_i14).Plus(_dafny.One)
				(_146_globalState).F17 = _141_v9
				_299_v106 = _299_v106
				var _338_v120 _dafny.Array
				_ = _338_v120
				var _nw57 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw57
				_338_v120 = _nw57
				_338_v120 = _338_v120
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_141_v9), 0))
				_ = _index44
				(_141_v9).ArraySet1((_131_v3).Times(_131_v3), (_index44).Int())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_141_v9), 0))
				_ = _index45
				var _rhs47 _dafny.CodePoint = _dafny.CodePoint('k')
				_ = _rhs47
				var _rhs48 _dafny.Int = _131_v3
				_ = _rhs48
				var _lhs43 _dafny.Array = _141_v9
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_141_v9), 0))
				_ = _lhs44
				_144_v12 = _rhs47
				(_lhs43).ArraySet1(_rhs48, (_lhs44).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _339_v121 *C4
	_ = _339_v121
	var _nw58 *C4 = New_C4_()
	_ = _nw58
	_nw58.Ctor__(_336_v119, _131_v3)
	_339_v121 = _nw58
	_339_v121 = _339_v121
	var _340_v122 D18
	_ = _340_v122
	_340_v122 = Companion_D18_.Create_DC55_(_144_v12, _128_v0, (_339_v121).F34(), _141_v9, _336_v119)
	var _source7 D18 = _340_v122
	_ = _source7
	if _source7.Is_DC55() {
		var _341___mcc_h27 _dafny.CodePoint = _source7.Get_().(D18_DC55).Cf94
		_ = _341___mcc_h27
		var _342___mcc_h28 bool = _source7.Get_().(D18_DC55).Cf95
		_ = _342___mcc_h28
		var _343___mcc_h29 _dafny.Int = _source7.Get_().(D18_DC55).Cf96
		_ = _343___mcc_h29
		var _344___mcc_h30 _dafny.Array = _source7.Get_().(D18_DC55).Cf97
		_ = _344___mcc_h30
		var _345___mcc_h31 bool = _source7.Get_().(D18_DC55).Cf98
		_ = _345___mcc_h31
		var _346_cf98 bool = _345___mcc_h31
		_ = _346_cf98
		var _347_cf97 _dafny.Array = _344___mcc_h30
		_ = _347_cf97
		var _348_cf96 _dafny.Int = _343___mcc_h29
		_ = _348_cf96
		var _349_cf95 bool = _342___mcc_h28
		_ = _349_cf95
		var _350_cf94 _dafny.CodePoint = _341___mcc_h27
		_ = _350_cf94
		_128_v0 = (_339_v121).F33()
		var _351_v123 _dafny.Map
		_ = _351_v123
		_351_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v7, (_128_v0) == ((_339_v121).F33()))
		(_146_globalState).F15 = (func() bool {
			if (_351_v123).Contains(_134_v7) {
				return (_351_v123).Get(_134_v7).(bool)
			}
			return (_333_v116).F31()
		})()
		var _352_v124 _dafny.Map
		_ = _352_v124
		_352_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v49, _301_v108)
		var _353_v125 D6
		_ = _353_v125
		_353_v125 = Companion_D6_.Create_DC21_(_129_v1)
		var _354_v126 _dafny.MultiSet
		_ = _354_v126
		_354_v126 = _dafny.MultiSetOf(_353_v125)
		var _355_v127 _dafny.Sequence
		_ = _355_v127
		_355_v127 = _dafny.SeqOf(Companion_D6_.Create_DC21_(_dafny.SetOf(_336_v119)), _353_v125)
		var _356_v128 _dafny.Sequence
		_ = _356_v128
		_356_v128 = _dafny.SeqOf(_354_v126, _dafny.MultiSetFromSeq(_355_v127), _dafny.MultiSetOf(_353_v125, _353_v125), _354_v126, _354_v126)
		var _rhs49 _dafny.Map = _352_v124
		_ = _rhs49
		var _rhs50 _dafny.Int = Companion_Default___.SafeModuloInt((_339_v121).F34(), _dafny.IntOfUint32((_134_v7).Cardinality()))
		_ = _rhs50
		var _rhs51 bool = (_354_v126).Equals((_356_v128).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.IntOfUint32((_356_v128).Cardinality()))).Uint32()).(_dafny.MultiSet))
		_ = _rhs51
		var _lhs45 *GlobalState = _146_globalState
		_ = _lhs45
		var _lhs46 *GlobalState = _146_globalState
		_ = _lhs46
		_352_v124 = _rhs49
		_lhs45.F2 = _rhs50
		_lhs46.F15 = _rhs51
		var _rhs52 _dafny.Int = _131_v3
		_ = _rhs52
		var _rhs53 _dafny.Int = _dafny.IntOfInt64(734)
		_ = _rhs53
		var _rhs54 bool = ((_339_v121).F34()).Cmp(_dafny.IntOfInt64(-143)) >= 0
		_ = _rhs54
		var _rhs55 _dafny.Set = _132_v4
		_ = _rhs55
		var _lhs47 *GlobalState = _146_globalState
		_ = _lhs47
		var _lhs48 *GlobalState = _146_globalState
		_ = _lhs48
		_131_v3 = _rhs52
		_lhs47.F2 = _rhs53
		_lhs48.F28 = _rhs54
		_132_v4 = _rhs55
	} else {
		var _357___mcc_h32 _dafny.Sequence = _source7.Get_().(D18_DC54).Cf93
		_ = _357___mcc_h32
		var _358_cf93 _dafny.Sequence = _357___mcc_h32
		_ = _358_cf93
		var _359_v129 D0
		_ = _359_v129
		_359_v129 = Companion_D0_.Create_DC0_((_299_v106).F29())
		var _rhs56 D0 = _359_v129
		_ = _rhs56
		var _rhs57 _dafny.Array = _203_v49
		_ = _rhs57
		_359_v129 = _rhs56
		_203_v49 = _rhs57
		var _360_v130 _dafny.Array
		_ = _360_v130
		var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
		_ = _nw59
		_360_v130 = _nw59
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_360_v130), 0))
		_ = _index46
		(_360_v130).ArraySet1(_203_v49, (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_360_v130), 0))
		_ = _index47
		(_360_v130).ArraySet1(_203_v49, (_index47).Int())
		var _361_v131 _dafny.Array
		_ = _361_v131
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_5
		var _nw60 _dafny.Array
		_ = _nw60
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw60 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = (func(_362_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_363_i15 _dafny.Int) _dafny.Sequence {
					return _362_v7
				}
			})(_134_v7)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw60 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw60).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw60).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_361_v131 = _nw60
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_361_v131), 0))
		_ = _index48
		(_361_v131).ArraySet1(_134_v7, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_361_v131), 0))
		_ = _index49
		(_361_v131).ArraySet1(_134_v7, (_index49).Int())
		_304_v111 = (_304_v111).Update((_333_v116).F32(), (_339_v121).F34())
	}
	var _364_i16 _dafny.Int
	_ = _364_i16
	_364_i16 = _dafny.Zero
	{
		for !((_333_v116).F31()) || ((_203_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_203_v49), 0))).Int()).(bool)) {
			{
				if (_364_i16).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_364_i16 = (_364_i16).Plus(_dafny.One)
				var _365_v132 _dafny.Map
				_ = _365_v132
				_365_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_301_v108).Cardinality()), _301_v108)
				var _366_v133 _dafny.Map
				_ = _366_v133
				_366_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v6, (_365_v132).Merge(_365_v132))
				_366_v133 = (_366_v133).Update((_133_v6).Update((_339_v121).F34(), Companion_Default___.Abs(_131_v3)), _365_v132)
				var _367_v134 D0
				_ = _367_v134
				_367_v134 = Companion_D0_.Create_DC1_(_203_v49, (_339_v121).F34(), _dafny.IntOfUint32((_299_v106.F30()).Cardinality()))
				_203_v49 = ((func() D0 {
					if (_299_v106).F29() {
						return _367_v134
					}
					return _367_v134
				})()).Dtor_cf1()
				_131_v3 = (_dafny.Zero).Minus(((_339_v121).F34()).Times((_129_v1).Cardinality()))
				(_146_globalState).F19 = _131_v3
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	(_146_globalState).F2 = (_dafny.Zero).Minus((_339_v121).F34())
	var _368_v135 _dafny.MultiSet
	_ = _368_v135
	_368_v135 = _dafny.MultiSetOf(_140_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-956))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg51 _dafny.Int) interface{} {
			return coer51(arg51)
		}
	}((func(_369_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_370_i17 _dafny.Int) _dafny.CodePoint {
			return _369_v12
		}
	})(_144_v12))))
	var _371_v136 bool
	_ = _371_v136
	var _372_v137 bool
	_ = _372_v137
	var _out8 bool
	_ = _out8
	var _out9 bool
	_ = _out9
	_out8, _out9 = (_339_v121).M6(Companion_Default___.Fm10((_339_v121).F34(), (_368_v135).Cardinality(), _146_globalState), _335_v118, _140_v8, Companion_D0_.Create_DC0_(!(!(_128_v0))), _146_globalState)
	_371_v136 = _out8
	_372_v137 = _out9
	_dafny.Print(_128_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v1).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_130_v2, _dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v4).Equals(_dafny.SetOf(_dafny.IntOfInt64(-208))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v6).Equals(_dafny.MultiSetOf(_dafny.Zero, _dafny.IntOfInt64(-1), _dafny.IntOfInt64(10), _dafny.IntOfInt64(9), _dafny.IntOfInt64(-916), _dafny.IntOfInt64(-208))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_134_v7, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_140_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-208), _dafny.CodePoint('r'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_144_v12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_146_globalState.F3, _dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_146_globalState).F5()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-208))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_146_globalState.F7, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F11.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState.F16).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState.F17).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState.F17).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState.F17).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F20.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_146_globalState).F22()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F25().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_globalState).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_globalState.F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v47).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v47).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_202_v48, _dafny.SeqOf(Companion_D2_.Create_DC8_(_dafny.CodePoint('l'), true), Companion_D2_.Create_DC8_(_dafny.CodePoint('l'), true), Companion_D2_.Create_DC8_(_dafny.CodePoint('l'), true), Companion_D2_.Create_DC8_(_dafny.CodePoint('l'), true), Companion_D2_.Create_DC8_(_dafny.CodePoint('l'), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v49).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_205_v50).Dtor_cf41(), _dafny.SeqOf(true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v50).Dtor_cf42())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v50).Dtor_cf43())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v50).Dtor_cf44())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v50).Dtor_cf45())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_298_v105).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).F29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_298_v105).ArrayGet1((_dafny.IntOfInt64(11)).Int())).F30().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_299_v106).F29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_299_v106.F30().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_300_v107).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_301_v108, _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_302_v109).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_304_v111).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-208))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v114).Dtor_cf88())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v114).Dtor_cf89())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v114).Dtor_cf90())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_306_v114).Dtor_cf91()).Equals(_dafny.MultiSetOf(_dafny.Zero, _dafny.IntOfInt64(-1), _dafny.IntOfInt64(10), _dafny.IntOfInt64(9), _dafny.IntOfInt64(-916), _dafny.IntOfInt64(-208))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v115).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_333_v116).F31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_333_v116).F32())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_334_v117).Dtor_cf111()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_335_v118.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_336_v119)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_337_i14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_339_v121).F33())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_339_v121).F34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v122).Dtor_cf94())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v122).Dtor_cf95())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v122).Dtor_cf96())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_340_v122).Dtor_cf97()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_340_v122).Dtor_cf97()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_340_v122).Dtor_cf97()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v122).Dtor_cf98())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_364_i16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_368_v135).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("rqigbg"), _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_371_v136)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_372_v137)
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

type D0_DC1 struct {
	Cf1 _dafny.Array
	Cf2 _dafny.Int
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array, Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.MultiSet
	Cf5 _dafny.Sequence
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.MultiSet, Cf5 _dafny.Sequence, Cf6 bool) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(false)
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.MultiSet {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + data.Cf5.VerbatimString(true) + ", " + _dafny.String(data.Cf6) + ")"
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
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6
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

type D1_DC4 struct {
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 bool) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf7 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
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
	return Companion_D1_.Create_DC4_(false)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf9() D1 {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8 == data2.Cf8
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool) D2 {
	return D2{D2_DC7{Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf12 _dafny.CodePoint
	Cf13 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf12 _dafny.CodePoint, Cf13 bool) D2 {
	return D2{D2_DC8{Cf12, Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf14 _dafny.Sequence
	Cf15 _dafny.Int
	Cf16 bool
	Cf17 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf14 _dafny.Sequence, Cf15 _dafny.Int, Cf16 bool, Cf17 _dafny.Int) D2 {
	return D2{D2_DC9{Cf14, Cf15, Cf16, Cf17}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
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

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D2_DC8).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf15
}

func (_this D2) Dtor_cf16() bool {
	return _this.Get_().(D2_DC9).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf17
}

func (_this D2) Dtor_cf10() _dafny.MultiSet {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + data.Cf14.VerbatimString(true) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf14.Equals(data2.Cf14) && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
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

type D3_DC11 struct {
	Cf19 bool
	Cf20 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 bool, Cf20 bool) D3 {
	return D3{D3_DC11{Cf19, Cf20}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf18 _dafny.MultiSet
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf18 _dafny.MultiSet) D3 {
	return D3{D3_DC10{Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC12 struct {
	Cf21 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf21 D3) D3 {
	return D3{D3_DC12{Cf21}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(false, false)
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf21() D3 {
	return _this.Get_().(D3_DC12).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf23 _dafny.Int) D4 {
	return D4{D4_DC14{Cf23}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf24 _dafny.Int
	Cf25 _dafny.Int
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf24 _dafny.Int, Cf25 _dafny.Int) D4 {
	return D4{D4_DC15{Cf24, Cf25}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC13 struct {
	Cf22 _dafny.Map
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 _dafny.Map) D4 {
	return D4{D4_DC13{Cf22}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(_dafny.Zero)
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf25
}

func (_this D4) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D5_DC17 struct {
	Cf27 _dafny.Set
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf27 _dafny.Set) D5 {
	return D5{D5_DC17{Cf27}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_() D5 {
	return D5{D5_DC18{}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf28 _dafny.Int
	Cf29 _dafny.Int
	Cf30 _dafny.CodePoint
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf28 _dafny.Int, Cf29 _dafny.Int, Cf30 _dafny.CodePoint) D5 {
	return D5{D5_DC19{Cf28, Cf29, Cf30}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC16 struct {
	Cf26 _dafny.Map
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf26 _dafny.Map) D5 {
	return D5{D5_DC16{Cf26}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC20 struct {
	Cf31 D5
}

func (D5_DC20) isD5() {}

func (CompanionStruct_D5_) Create_DC20_(Cf31 D5) D5 {
	return D5{D5_DC20{Cf31}}
}

func (_this D5) Is_DC20() bool {
	_, ok := _this.Get_().(D5_DC20)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.EmptySet)
}

func (_this D5) Dtor_cf27() _dafny.Set {
	return _this.Get_().(D5_DC17).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D5_DC19).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D5_DC19).Cf29
}

func (_this D5) Dtor_cf30() _dafny.CodePoint {
	return _this.Get_().(D5_DC19).Cf30
}

func (_this D5) Dtor_cf26() _dafny.Map {
	return _this.Get_().(D5_DC16).Cf26
}

func (_this D5) Dtor_cf31() D5 {
	return _this.Get_().(D5_DC20).Cf31
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D5_DC20:
		{
			return "D5.DC20" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf27.Equals(data2.Cf27)
		}
	case D5_DC18:
		{
			_, ok := other.Get_().(D5_DC18)
			return ok
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf26.Equals(data2.Cf26)
		}
	case D5_DC20:
		{
			data2, ok := other.Get_().(D5_DC20)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D6_DC22 struct {
	Cf33 _dafny.Int
	Cf34 _dafny.Int
	Cf35 bool
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_(Cf33 _dafny.Int, Cf34 _dafny.Int, Cf35 bool) D6 {
	return D6{D6_DC22{Cf33, Cf34, Cf35}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC23 struct {
	Cf36 bool
	Cf37 _dafny.Int
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf36 bool, Cf37 _dafny.Int) D6 {
	return D6{D6_DC23{Cf36, Cf37}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

type D6_DC21 struct {
	Cf32 _dafny.Set
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf32 _dafny.Set) D6 {
	return D6{D6_DC21{Cf32}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC22_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D6) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf33
}

func (_this D6) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf34
}

func (_this D6) Dtor_cf35() bool {
	return _this.Get_().(D6_DC22).Cf35
}

func (_this D6) Dtor_cf36() bool {
	return _this.Get_().(D6_DC23).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf37
}

func (_this D6) Dtor_cf32() _dafny.Set {
	return _this.Get_().(D6_DC21).Cf32
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC22:
		{
			return "D6.DC22" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC22:
		{
			data2, ok := other.Get_().(D6_DC22)
			return ok && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D7_DC24 struct {
	Cf38 _dafny.Map
}

func (D7_DC24) isD7() {}

func (CompanionStruct_D7_) Create_DC24_(Cf38 _dafny.Map) D7 {
	return D7{D7_DC24{Cf38}}
}

func (_this D7) Is_DC24() bool {
	_, ok := _this.Get_().(D7_DC24)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf38() _dafny.Map {
	return _this.Get_().(D7_DC24).Cf38
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC24:
		{
			return "D7.DC24" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC24:
		{
			data2, ok := other.Get_().(D7_DC24)
			return ok && data1.Cf38.Equals(data2.Cf38)
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

type D8_DC26 struct {
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_() D8 {
	return D8{D8_DC26{}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC25 struct {
	Cf39 _dafny.Sequence
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf39 _dafny.Sequence) D8 {
	return D8{D8_DC25{Cf39}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC26_()
}

func (_this D8) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D8_DC25).Cf39
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC26:
		{
			_, ok := other.Get_().(D8_DC26)
			return ok
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D9_DC28 struct {
	Cf41 _dafny.Sequence
	Cf42 bool
	Cf43 _dafny.Int
	Cf44 bool
	Cf45 bool
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf41 _dafny.Sequence, Cf42 bool, Cf43 _dafny.Int, Cf44 bool, Cf45 bool) D9 {
	return D9{D9_DC28{Cf41, Cf42, Cf43, Cf44, Cf45}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

type D9_DC29 struct {
	Cf46 _dafny.Int
	Cf47 _dafny.Int
	Cf48 bool
}

func (D9_DC29) isD9() {}

func (CompanionStruct_D9_) Create_DC29_(Cf46 _dafny.Int, Cf47 _dafny.Int, Cf48 bool) D9 {
	return D9{D9_DC29{Cf46, Cf47, Cf48}}
}

func (_this D9) Is_DC29() bool {
	_, ok := _this.Get_().(D9_DC29)
	return ok
}

type D9_DC27 struct {
	Cf40 _dafny.Sequence
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf40 _dafny.Sequence) D9 {
	return D9{D9_DC27{Cf40}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC28_(_dafny.EmptySeq, false, _dafny.Zero, false, false)
}

func (_this D9) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D9_DC28).Cf41
}

func (_this D9) Dtor_cf42() bool {
	return _this.Get_().(D9_DC28).Cf42
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf43
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC28).Cf44
}

func (_this D9) Dtor_cf45() bool {
	return _this.Get_().(D9_DC28).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D9_DC29).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC29).Cf47
}

func (_this D9) Dtor_cf48() bool {
	return _this.Get_().(D9_DC29).Cf48
}

func (_this D9) Dtor_cf40() _dafny.Sequence {
	return _this.Get_().(D9_DC27).Cf40
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D9_DC29:
		{
			return "D9.DC29" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf41.Equals(data2.Cf41) && data1.Cf42 == data2.Cf42 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45
		}
	case D9_DC29:
		{
			data2, ok := other.Get_().(D9_DC29)
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D10_DC31 struct {
	Cf50 bool
	Cf51 bool
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf50 bool, Cf51 bool) D10 {
	return D10{D10_DC31{Cf50, Cf51}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

type D10_DC32 struct {
	Cf52 _dafny.CodePoint
	Cf53 bool
}

func (D10_DC32) isD10() {}

func (CompanionStruct_D10_) Create_DC32_(Cf52 _dafny.CodePoint, Cf53 bool) D10 {
	return D10{D10_DC32{Cf52, Cf53}}
}

func (_this D10) Is_DC32() bool {
	_, ok := _this.Get_().(D10_DC32)
	return ok
}

type D10_DC33 struct {
	Cf54 bool
	Cf55 _dafny.Int
}

func (D10_DC33) isD10() {}

func (CompanionStruct_D10_) Create_DC33_(Cf54 bool, Cf55 _dafny.Int) D10 {
	return D10{D10_DC33{Cf54, Cf55}}
}

func (_this D10) Is_DC33() bool {
	_, ok := _this.Get_().(D10_DC33)
	return ok
}

type D10_DC30 struct {
	Cf49 T0
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf49 T0) D10 {
	return D10{D10_DC30{Cf49}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC31_(false, false)
}

func (_this D10) Dtor_cf50() bool {
	return _this.Get_().(D10_DC31).Cf50
}

func (_this D10) Dtor_cf51() bool {
	return _this.Get_().(D10_DC31).Cf51
}

func (_this D10) Dtor_cf52() _dafny.CodePoint {
	return _this.Get_().(D10_DC32).Cf52
}

func (_this D10) Dtor_cf53() bool {
	return _this.Get_().(D10_DC32).Cf53
}

func (_this D10) Dtor_cf54() bool {
	return _this.Get_().(D10_DC33).Cf54
}

func (_this D10) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D10_DC33).Cf55
}

func (_this D10) Dtor_cf49() T0 {
	return _this.Get_().(D10_DC30).Cf49
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC31:
		{
			return "D10.DC31" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ")"
		}
	case D10_DC32:
		{
			return "D10.DC32" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D10_DC33:
		{
			return "D10.DC33" + "(" + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf50 == data2.Cf50 && data1.Cf51 == data2.Cf51
		}
	case D10_DC32:
		{
			data2, ok := other.Get_().(D10_DC32)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53
		}
	case D10_DC33:
		{
			data2, ok := other.Get_().(D10_DC33)
			return ok && data1.Cf54 == data2.Cf54 && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && _dafny.AreEqual(data1.Cf49, data2.Cf49)
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

type D11_DC35 struct {
	Cf57 _dafny.Sequence
	Cf58 bool
}

func (D11_DC35) isD11() {}

func (CompanionStruct_D11_) Create_DC35_(Cf57 _dafny.Sequence, Cf58 bool) D11 {
	return D11{D11_DC35{Cf57, Cf58}}
}

func (_this D11) Is_DC35() bool {
	_, ok := _this.Get_().(D11_DC35)
	return ok
}

type D11_DC36 struct {
	Cf59 D6
	Cf60 _dafny.Set
	Cf61 bool
}

func (D11_DC36) isD11() {}

func (CompanionStruct_D11_) Create_DC36_(Cf59 D6, Cf60 _dafny.Set, Cf61 bool) D11 {
	return D11{D11_DC36{Cf59, Cf60, Cf61}}
}

func (_this D11) Is_DC36() bool {
	_, ok := _this.Get_().(D11_DC36)
	return ok
}

type D11_DC37 struct {
	Cf62 _dafny.Int
	Cf63 bool
}

func (D11_DC37) isD11() {}

func (CompanionStruct_D11_) Create_DC37_(Cf62 _dafny.Int, Cf63 bool) D11 {
	return D11{D11_DC37{Cf62, Cf63}}
}

func (_this D11) Is_DC37() bool {
	_, ok := _this.Get_().(D11_DC37)
	return ok
}

type D11_DC34 struct {
	Cf56 _dafny.Map
}

func (D11_DC34) isD11() {}

func (CompanionStruct_D11_) Create_DC34_(Cf56 _dafny.Map) D11 {
	return D11{D11_DC34{Cf56}}
}

func (_this D11) Is_DC34() bool {
	_, ok := _this.Get_().(D11_DC34)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC35_(_dafny.EmptySeq, false)
}

func (_this D11) Dtor_cf57() _dafny.Sequence {
	return _this.Get_().(D11_DC35).Cf57
}

func (_this D11) Dtor_cf58() bool {
	return _this.Get_().(D11_DC35).Cf58
}

func (_this D11) Dtor_cf59() D6 {
	return _this.Get_().(D11_DC36).Cf59
}

func (_this D11) Dtor_cf60() _dafny.Set {
	return _this.Get_().(D11_DC36).Cf60
}

func (_this D11) Dtor_cf61() bool {
	return _this.Get_().(D11_DC36).Cf61
}

func (_this D11) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D11_DC37).Cf62
}

func (_this D11) Dtor_cf63() bool {
	return _this.Get_().(D11_DC37).Cf63
}

func (_this D11) Dtor_cf56() _dafny.Map {
	return _this.Get_().(D11_DC34).Cf56
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC35:
		{
			return "D11.DC35" + "(" + data.Cf57.VerbatimString(true) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D11_DC36:
		{
			return "D11.DC36" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D11_DC37:
		{
			return "D11.DC37" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D11_DC34:
		{
			return "D11.DC34" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC35:
		{
			data2, ok := other.Get_().(D11_DC35)
			return ok && data1.Cf57.Equals(data2.Cf57) && data1.Cf58 == data2.Cf58
		}
	case D11_DC36:
		{
			data2, ok := other.Get_().(D11_DC36)
			return ok && data1.Cf59.Equals(data2.Cf59) && data1.Cf60.Equals(data2.Cf60) && data1.Cf61 == data2.Cf61
		}
	case D11_DC37:
		{
			data2, ok := other.Get_().(D11_DC37)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63
		}
	case D11_DC34:
		{
			data2, ok := other.Get_().(D11_DC34)
			return ok && data1.Cf56.Equals(data2.Cf56)
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

type D12_DC39 struct {
	Cf65 _dafny.Int
	Cf66 bool
	Cf67 _dafny.Int
}

func (D12_DC39) isD12() {}

func (CompanionStruct_D12_) Create_DC39_(Cf65 _dafny.Int, Cf66 bool, Cf67 _dafny.Int) D12 {
	return D12{D12_DC39{Cf65, Cf66, Cf67}}
}

func (_this D12) Is_DC39() bool {
	_, ok := _this.Get_().(D12_DC39)
	return ok
}

type D12_DC38 struct {
	Cf64 _dafny.Array
}

func (D12_DC38) isD12() {}

func (CompanionStruct_D12_) Create_DC38_(Cf64 _dafny.Array) D12 {
	return D12{D12_DC38{Cf64}}
}

func (_this D12) Is_DC38() bool {
	_, ok := _this.Get_().(D12_DC38)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC39_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D12) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D12_DC39).Cf65
}

func (_this D12) Dtor_cf66() bool {
	return _this.Get_().(D12_DC39).Cf66
}

func (_this D12) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D12_DC39).Cf67
}

func (_this D12) Dtor_cf64() _dafny.Array {
	return _this.Get_().(D12_DC38).Cf64
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC39:
		{
			return "D12.DC39" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D12_DC38:
		{
			return "D12.DC38" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC39:
		{
			data2, ok := other.Get_().(D12_DC39)
			return ok && data1.Cf65.Cmp(data2.Cf65) == 0 && data1.Cf66 == data2.Cf66 && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D12_DC38:
		{
			data2, ok := other.Get_().(D12_DC38)
			return ok && data1.Cf64 == data2.Cf64
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

type D13_DC41 struct {
}

func (D13_DC41) isD13() {}

func (CompanionStruct_D13_) Create_DC41_() D13 {
	return D13{D13_DC41{}}
}

func (_this D13) Is_DC41() bool {
	_, ok := _this.Get_().(D13_DC41)
	return ok
}

type D13_DC40 struct {
	Cf68 *C4
}

func (D13_DC40) isD13() {}

func (CompanionStruct_D13_) Create_DC40_(Cf68 *C4) D13 {
	return D13{D13_DC40{Cf68}}
}

func (_this D13) Is_DC40() bool {
	_, ok := _this.Get_().(D13_DC40)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC41_()
}

func (_this D13) Dtor_cf68() *C4 {
	return _this.Get_().(D13_DC40).Cf68
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC41:
		{
			return "D13.DC41"
		}
	case D13_DC40:
		{
			return "D13.DC40" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC41:
		{
			_, ok := other.Get_().(D13_DC41)
			return ok
		}
	case D13_DC40:
		{
			data2, ok := other.Get_().(D13_DC40)
			return ok && data1.Cf68 == data2.Cf68
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

type D14_DC42 struct {
	Cf69 _dafny.Map
}

func (D14_DC42) isD14() {}

func (CompanionStruct_D14_) Create_DC42_(Cf69 _dafny.Map) D14 {
	return D14{D14_DC42{Cf69}}
}

func (_this D14) Is_DC42() bool {
	_, ok := _this.Get_().(D14_DC42)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D14) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D14_DC42).Cf69
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC42:
		{
			return "D14.DC42" + "(" + _dafny.String(data.Cf69) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC42:
		{
			data2, ok := other.Get_().(D14_DC42)
			return ok && data1.Cf69.Equals(data2.Cf69)
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

type D15_DC44 struct {
	Cf71 _dafny.Int
	Cf72 bool
	Cf73 _dafny.Array
	Cf74 _dafny.Array
	Cf75 _dafny.Int
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf71 _dafny.Int, Cf72 bool, Cf73 _dafny.Array, Cf74 _dafny.Array, Cf75 _dafny.Int) D15 {
	return D15{D15_DC44{Cf71, Cf72, Cf73, Cf74, Cf75}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

type D15_DC45 struct {
	Cf76 _dafny.Int
	Cf77 _dafny.Int
	Cf78 bool
	Cf79 bool
	Cf80 _dafny.Array
}

func (D15_DC45) isD15() {}

func (CompanionStruct_D15_) Create_DC45_(Cf76 _dafny.Int, Cf77 _dafny.Int, Cf78 bool, Cf79 bool, Cf80 _dafny.Array) D15 {
	return D15{D15_DC45{Cf76, Cf77, Cf78, Cf79, Cf80}}
}

func (_this D15) Is_DC45() bool {
	_, ok := _this.Get_().(D15_DC45)
	return ok
}

type D15_DC46 struct {
}

func (D15_DC46) isD15() {}

func (CompanionStruct_D15_) Create_DC46_() D15 {
	return D15{D15_DC46{}}
}

func (_this D15) Is_DC46() bool {
	_, ok := _this.Get_().(D15_DC46)
	return ok
}

type D15_DC43 struct {
	Cf70 _dafny.Map
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf70 _dafny.Map) D15 {
	return D15{D15_DC43{Cf70}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC44_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D15) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D15_DC44).Cf71
}

func (_this D15) Dtor_cf72() bool {
	return _this.Get_().(D15_DC44).Cf72
}

func (_this D15) Dtor_cf73() _dafny.Array {
	return _this.Get_().(D15_DC44).Cf73
}

func (_this D15) Dtor_cf74() _dafny.Array {
	return _this.Get_().(D15_DC44).Cf74
}

func (_this D15) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D15_DC44).Cf75
}

func (_this D15) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D15_DC45).Cf76
}

func (_this D15) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D15_DC45).Cf77
}

func (_this D15) Dtor_cf78() bool {
	return _this.Get_().(D15_DC45).Cf78
}

func (_this D15) Dtor_cf79() bool {
	return _this.Get_().(D15_DC45).Cf79
}

func (_this D15) Dtor_cf80() _dafny.Array {
	return _this.Get_().(D15_DC45).Cf80
}

func (_this D15) Dtor_cf70() _dafny.Map {
	return _this.Get_().(D15_DC43).Cf70
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ")"
		}
	case D15_DC45:
		{
			return "D15.DC45" + "(" + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ")"
		}
	case D15_DC46:
		{
			return "D15.DC46"
		}
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
			return ok && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72 == data2.Cf72 && data1.Cf73 == data2.Cf73 && data1.Cf74 == data2.Cf74 && data1.Cf75.Cmp(data2.Cf75) == 0
		}
	case D15_DC45:
		{
			data2, ok := other.Get_().(D15_DC45)
			return ok && data1.Cf76.Cmp(data2.Cf76) == 0 && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78 == data2.Cf78 && data1.Cf79 == data2.Cf79 && data1.Cf80 == data2.Cf80
		}
	case D15_DC46:
		{
			_, ok := other.Get_().(D15_DC46)
			return ok
		}
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf70.Equals(data2.Cf70)
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

type D16_DC48 struct {
}

func (D16_DC48) isD16() {}

func (CompanionStruct_D16_) Create_DC48_() D16 {
	return D16{D16_DC48{}}
}

func (_this D16) Is_DC48() bool {
	_, ok := _this.Get_().(D16_DC48)
	return ok
}

type D16_DC47 struct {
	Cf81 _dafny.Set
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf81 _dafny.Set) D16 {
	return D16{D16_DC47{Cf81}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC48_()
}

func (_this D16) Dtor_cf81() _dafny.Set {
	return _this.Get_().(D16_DC47).Cf81
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC48:
		{
			return "D16.DC48"
		}
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf81) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC48:
		{
			_, ok := other.Get_().(D16_DC48)
			return ok
		}
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf81.Equals(data2.Cf81)
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

type D17_DC50 struct {
	Cf83 bool
	Cf84 bool
	Cf85 T1
}

func (D17_DC50) isD17() {}

func (CompanionStruct_D17_) Create_DC50_(Cf83 bool, Cf84 bool, Cf85 T1) D17 {
	return D17{D17_DC50{Cf83, Cf84, Cf85}}
}

func (_this D17) Is_DC50() bool {
	_, ok := _this.Get_().(D17_DC50)
	return ok
}

type D17_DC51 struct {
	Cf86 _dafny.Int
	Cf87 _dafny.Array
}

func (D17_DC51) isD17() {}

func (CompanionStruct_D17_) Create_DC51_(Cf86 _dafny.Int, Cf87 _dafny.Array) D17 {
	return D17{D17_DC51{Cf86, Cf87}}
}

func (_this D17) Is_DC51() bool {
	_, ok := _this.Get_().(D17_DC51)
	return ok
}

type D17_DC52 struct {
	Cf88 bool
	Cf89 bool
	Cf90 _dafny.Int
	Cf91 _dafny.MultiSet
}

func (D17_DC52) isD17() {}

func (CompanionStruct_D17_) Create_DC52_(Cf88 bool, Cf89 bool, Cf90 _dafny.Int, Cf91 _dafny.MultiSet) D17 {
	return D17{D17_DC52{Cf88, Cf89, Cf90, Cf91}}
}

func (_this D17) Is_DC52() bool {
	_, ok := _this.Get_().(D17_DC52)
	return ok
}

type D17_DC49 struct {
	Cf82 _dafny.Set
}

func (D17_DC49) isD17() {}

func (CompanionStruct_D17_) Create_DC49_(Cf82 _dafny.Set) D17 {
	return D17{D17_DC49{Cf82}}
}

func (_this D17) Is_DC49() bool {
	_, ok := _this.Get_().(D17_DC49)
	return ok
}

type D17_DC53 struct {
	Cf92 D17
}

func (D17_DC53) isD17() {}

func (CompanionStruct_D17_) Create_DC53_(Cf92 D17) D17 {
	return D17{D17_DC53{Cf92}}
}

func (_this D17) Is_DC53() bool {
	_, ok := _this.Get_().(D17_DC53)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC50_(false, false, (T1)(nil))
}

func (_this D17) Dtor_cf83() bool {
	return _this.Get_().(D17_DC50).Cf83
}

func (_this D17) Dtor_cf84() bool {
	return _this.Get_().(D17_DC50).Cf84
}

func (_this D17) Dtor_cf85() T1 {
	return _this.Get_().(D17_DC50).Cf85
}

func (_this D17) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D17_DC51).Cf86
}

func (_this D17) Dtor_cf87() _dafny.Array {
	return _this.Get_().(D17_DC51).Cf87
}

func (_this D17) Dtor_cf88() bool {
	return _this.Get_().(D17_DC52).Cf88
}

func (_this D17) Dtor_cf89() bool {
	return _this.Get_().(D17_DC52).Cf89
}

func (_this D17) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D17_DC52).Cf90
}

func (_this D17) Dtor_cf91() _dafny.MultiSet {
	return _this.Get_().(D17_DC52).Cf91
}

func (_this D17) Dtor_cf82() _dafny.Set {
	return _this.Get_().(D17_DC49).Cf82
}

func (_this D17) Dtor_cf92() D17 {
	return _this.Get_().(D17_DC53).Cf92
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC50:
		{
			return "D17.DC50" + "(" + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D17_DC51:
		{
			return "D17.DC51" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ")"
		}
	case D17_DC52:
		{
			return "D17.DC52" + "(" + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D17_DC49:
		{
			return "D17.DC49" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D17_DC53:
		{
			return "D17.DC53" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC50:
		{
			data2, ok := other.Get_().(D17_DC50)
			return ok && data1.Cf83 == data2.Cf83 && data1.Cf84 == data2.Cf84 && _dafny.AreEqual(data1.Cf85, data2.Cf85)
		}
	case D17_DC51:
		{
			data2, ok := other.Get_().(D17_DC51)
			return ok && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87 == data2.Cf87
		}
	case D17_DC52:
		{
			data2, ok := other.Get_().(D17_DC52)
			return ok && data1.Cf88 == data2.Cf88 && data1.Cf89 == data2.Cf89 && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91.Equals(data2.Cf91)
		}
	case D17_DC49:
		{
			data2, ok := other.Get_().(D17_DC49)
			return ok && data1.Cf82.Equals(data2.Cf82)
		}
	case D17_DC53:
		{
			data2, ok := other.Get_().(D17_DC53)
			return ok && data1.Cf92.Equals(data2.Cf92)
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

type D18_DC55 struct {
	Cf94 _dafny.CodePoint
	Cf95 bool
	Cf96 _dafny.Int
	Cf97 _dafny.Array
	Cf98 bool
}

func (D18_DC55) isD18() {}

func (CompanionStruct_D18_) Create_DC55_(Cf94 _dafny.CodePoint, Cf95 bool, Cf96 _dafny.Int, Cf97 _dafny.Array, Cf98 bool) D18 {
	return D18{D18_DC55{Cf94, Cf95, Cf96, Cf97, Cf98}}
}

func (_this D18) Is_DC55() bool {
	_, ok := _this.Get_().(D18_DC55)
	return ok
}

type D18_DC54 struct {
	Cf93 _dafny.Sequence
}

func (D18_DC54) isD18() {}

func (CompanionStruct_D18_) Create_DC54_(Cf93 _dafny.Sequence) D18 {
	return D18{D18_DC54{Cf93}}
}

func (_this D18) Is_DC54() bool {
	_, ok := _this.Get_().(D18_DC54)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC55_(_dafny.CodePoint('D'), false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D18) Dtor_cf94() _dafny.CodePoint {
	return _this.Get_().(D18_DC55).Cf94
}

func (_this D18) Dtor_cf95() bool {
	return _this.Get_().(D18_DC55).Cf95
}

func (_this D18) Dtor_cf96() _dafny.Int {
	return _this.Get_().(D18_DC55).Cf96
}

func (_this D18) Dtor_cf97() _dafny.Array {
	return _this.Get_().(D18_DC55).Cf97
}

func (_this D18) Dtor_cf98() bool {
	return _this.Get_().(D18_DC55).Cf98
}

func (_this D18) Dtor_cf93() _dafny.Sequence {
	return _this.Get_().(D18_DC54).Cf93
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC55:
		{
			return "D18.DC55" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ", " + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ")"
		}
	case D18_DC54:
		{
			return "D18.DC54" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC55:
		{
			data2, ok := other.Get_().(D18_DC55)
			return ok && data1.Cf94 == data2.Cf94 && data1.Cf95 == data2.Cf95 && data1.Cf96.Cmp(data2.Cf96) == 0 && data1.Cf97 == data2.Cf97 && data1.Cf98 == data2.Cf98
		}
	case D18_DC54:
		{
			data2, ok := other.Get_().(D18_DC54)
			return ok && data1.Cf93.Equals(data2.Cf93)
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

type D19_DC57 struct {
	Cf100 _dafny.Array
	Cf101 _dafny.Int
	Cf102 T0
	Cf103 *C3
	Cf104 _dafny.Int
}

func (D19_DC57) isD19() {}

func (CompanionStruct_D19_) Create_DC57_(Cf100 _dafny.Array, Cf101 _dafny.Int, Cf102 T0, Cf103 *C3, Cf104 _dafny.Int) D19 {
	return D19{D19_DC57{Cf100, Cf101, Cf102, Cf103, Cf104}}
}

func (_this D19) Is_DC57() bool {
	_, ok := _this.Get_().(D19_DC57)
	return ok
}

type D19_DC58 struct {
	Cf105 bool
	Cf106 bool
	Cf107 bool
}

func (D19_DC58) isD19() {}

func (CompanionStruct_D19_) Create_DC58_(Cf105 bool, Cf106 bool, Cf107 bool) D19 {
	return D19{D19_DC58{Cf105, Cf106, Cf107}}
}

func (_this D19) Is_DC58() bool {
	_, ok := _this.Get_().(D19_DC58)
	return ok
}

type D19_DC59 struct {
	Cf108 bool
	Cf109 _dafny.Int
}

func (D19_DC59) isD19() {}

func (CompanionStruct_D19_) Create_DC59_(Cf108 bool, Cf109 _dafny.Int) D19 {
	return D19{D19_DC59{Cf108, Cf109}}
}

func (_this D19) Is_DC59() bool {
	_, ok := _this.Get_().(D19_DC59)
	return ok
}

type D19_DC56 struct {
	Cf99 _dafny.Set
}

func (D19_DC56) isD19() {}

func (CompanionStruct_D19_) Create_DC56_(Cf99 _dafny.Set) D19 {
	return D19{D19_DC56{Cf99}}
}

func (_this D19) Is_DC56() bool {
	_, ok := _this.Get_().(D19_DC56)
	return ok
}

type D19_DC60 struct {
	Cf110 D19
}

func (D19_DC60) isD19() {}

func (CompanionStruct_D19_) Create_DC60_(Cf110 D19) D19 {
	return D19{D19_DC60{Cf110}}
}

func (_this D19) Is_DC60() bool {
	_, ok := _this.Get_().(D19_DC60)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC57_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, (T0)(nil), (*C3)(nil), _dafny.Zero)
}

func (_this D19) Dtor_cf100() _dafny.Array {
	return _this.Get_().(D19_DC57).Cf100
}

func (_this D19) Dtor_cf101() _dafny.Int {
	return _this.Get_().(D19_DC57).Cf101
}

func (_this D19) Dtor_cf102() T0 {
	return _this.Get_().(D19_DC57).Cf102
}

func (_this D19) Dtor_cf103() *C3 {
	return _this.Get_().(D19_DC57).Cf103
}

func (_this D19) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D19_DC57).Cf104
}

func (_this D19) Dtor_cf105() bool {
	return _this.Get_().(D19_DC58).Cf105
}

func (_this D19) Dtor_cf106() bool {
	return _this.Get_().(D19_DC58).Cf106
}

func (_this D19) Dtor_cf107() bool {
	return _this.Get_().(D19_DC58).Cf107
}

func (_this D19) Dtor_cf108() bool {
	return _this.Get_().(D19_DC59).Cf108
}

func (_this D19) Dtor_cf109() _dafny.Int {
	return _this.Get_().(D19_DC59).Cf109
}

func (_this D19) Dtor_cf99() _dafny.Set {
	return _this.Get_().(D19_DC56).Cf99
}

func (_this D19) Dtor_cf110() D19 {
	return _this.Get_().(D19_DC60).Cf110
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC57:
		{
			return "D19.DC57" + "(" + _dafny.String(data.Cf100) + ", " + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ", " + _dafny.String(data.Cf104) + ")"
		}
	case D19_DC58:
		{
			return "D19.DC58" + "(" + _dafny.String(data.Cf105) + ", " + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ")"
		}
	case D19_DC59:
		{
			return "D19.DC59" + "(" + _dafny.String(data.Cf108) + ", " + _dafny.String(data.Cf109) + ")"
		}
	case D19_DC56:
		{
			return "D19.DC56" + "(" + _dafny.String(data.Cf99) + ")"
		}
	case D19_DC60:
		{
			return "D19.DC60" + "(" + _dafny.String(data.Cf110) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC57:
		{
			data2, ok := other.Get_().(D19_DC57)
			return ok && data1.Cf100 == data2.Cf100 && data1.Cf101.Cmp(data2.Cf101) == 0 && _dafny.AreEqual(data1.Cf102, data2.Cf102) && data1.Cf103 == data2.Cf103 && data1.Cf104.Cmp(data2.Cf104) == 0
		}
	case D19_DC58:
		{
			data2, ok := other.Get_().(D19_DC58)
			return ok && data1.Cf105 == data2.Cf105 && data1.Cf106 == data2.Cf106 && data1.Cf107 == data2.Cf107
		}
	case D19_DC59:
		{
			data2, ok := other.Get_().(D19_DC59)
			return ok && data1.Cf108 == data2.Cf108 && data1.Cf109.Cmp(data2.Cf109) == 0
		}
	case D19_DC56:
		{
			data2, ok := other.Get_().(D19_DC56)
			return ok && data1.Cf99.Equals(data2.Cf99)
		}
	case D19_DC60:
		{
			data2, ok := other.Get_().(D19_DC60)
			return ok && data1.Cf110.Equals(data2.Cf110)
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

type D20_DC62 struct {
}

func (D20_DC62) isD20() {}

func (CompanionStruct_D20_) Create_DC62_() D20 {
	return D20{D20_DC62{}}
}

func (_this D20) Is_DC62() bool {
	_, ok := _this.Get_().(D20_DC62)
	return ok
}

type D20_DC63 struct {
	Cf112 _dafny.Int
}

func (D20_DC63) isD20() {}

func (CompanionStruct_D20_) Create_DC63_(Cf112 _dafny.Int) D20 {
	return D20{D20_DC63{Cf112}}
}

func (_this D20) Is_DC63() bool {
	_, ok := _this.Get_().(D20_DC63)
	return ok
}

type D20_DC61 struct {
	Cf111 _dafny.Sequence
}

func (D20_DC61) isD20() {}

func (CompanionStruct_D20_) Create_DC61_(Cf111 _dafny.Sequence) D20 {
	return D20{D20_DC61{Cf111}}
}

func (_this D20) Is_DC61() bool {
	_, ok := _this.Get_().(D20_DC61)
	return ok
}

type D20_DC64 struct {
	Cf113 D20
}

func (D20_DC64) isD20() {}

func (CompanionStruct_D20_) Create_DC64_(Cf113 D20) D20 {
	return D20{D20_DC64{Cf113}}
}

func (_this D20) Is_DC64() bool {
	_, ok := _this.Get_().(D20_DC64)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC62_()
}

func (_this D20) Dtor_cf112() _dafny.Int {
	return _this.Get_().(D20_DC63).Cf112
}

func (_this D20) Dtor_cf111() _dafny.Sequence {
	return _this.Get_().(D20_DC61).Cf111
}

func (_this D20) Dtor_cf113() D20 {
	return _this.Get_().(D20_DC64).Cf113
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC62:
		{
			return "D20.DC62"
		}
	case D20_DC63:
		{
			return "D20.DC63" + "(" + _dafny.String(data.Cf112) + ")"
		}
	case D20_DC61:
		{
			return "D20.DC61" + "(" + _dafny.String(data.Cf111) + ")"
		}
	case D20_DC64:
		{
			return "D20.DC64" + "(" + _dafny.String(data.Cf113) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC62:
		{
			_, ok := other.Get_().(D20_DC62)
			return ok
		}
	case D20_DC63:
		{
			data2, ok := other.Get_().(D20_DC63)
			return ok && data1.Cf112.Cmp(data2.Cf112) == 0
		}
	case D20_DC61:
		{
			data2, ok := other.Get_().(D20_DC61)
			return ok && data1.Cf111.Equals(data2.Cf111)
		}
	case D20_DC64:
		{
			data2, ok := other.Get_().(D20_DC64)
			return ok && data1.Cf113.Equals(data2.Cf113)
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

type D21_DC66 struct {
	Cf115 _dafny.Int
}

func (D21_DC66) isD21() {}

func (CompanionStruct_D21_) Create_DC66_(Cf115 _dafny.Int) D21 {
	return D21{D21_DC66{Cf115}}
}

func (_this D21) Is_DC66() bool {
	_, ok := _this.Get_().(D21_DC66)
	return ok
}

type D21_DC67 struct {
	Cf116 _dafny.Int
}

func (D21_DC67) isD21() {}

func (CompanionStruct_D21_) Create_DC67_(Cf116 _dafny.Int) D21 {
	return D21{D21_DC67{Cf116}}
}

func (_this D21) Is_DC67() bool {
	_, ok := _this.Get_().(D21_DC67)
	return ok
}

type D21_DC65 struct {
	Cf114 _dafny.Set
}

func (D21_DC65) isD21() {}

func (CompanionStruct_D21_) Create_DC65_(Cf114 _dafny.Set) D21 {
	return D21{D21_DC65{Cf114}}
}

func (_this D21) Is_DC65() bool {
	_, ok := _this.Get_().(D21_DC65)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC66_(_dafny.Zero)
}

func (_this D21) Dtor_cf115() _dafny.Int {
	return _this.Get_().(D21_DC66).Cf115
}

func (_this D21) Dtor_cf116() _dafny.Int {
	return _this.Get_().(D21_DC67).Cf116
}

func (_this D21) Dtor_cf114() _dafny.Set {
	return _this.Get_().(D21_DC65).Cf114
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC66:
		{
			return "D21.DC66" + "(" + _dafny.String(data.Cf115) + ")"
		}
	case D21_DC67:
		{
			return "D21.DC67" + "(" + _dafny.String(data.Cf116) + ")"
		}
	case D21_DC65:
		{
			return "D21.DC65" + "(" + _dafny.String(data.Cf114) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC66:
		{
			data2, ok := other.Get_().(D21_DC66)
			return ok && data1.Cf115.Cmp(data2.Cf115) == 0
		}
	case D21_DC67:
		{
			data2, ok := other.Get_().(D21_DC67)
			return ok && data1.Cf116.Cmp(data2.Cf116) == 0
		}
	case D21_DC65:
		{
			data2, ok := other.Get_().(D21_DC65)
			return ok && data1.Cf114.Equals(data2.Cf114)
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

// Definition of trait T0
type T0 interface {
	String() string
	M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool)
	M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool)
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
	M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool)
	M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool)
	F30() _dafny.Sequence
	F30_set_(value _dafny.Sequence)
	F29() bool
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
	F1   _dafny.Int
	F2   _dafny.Int
	F3   _dafny.Sequence
	F7   _dafny.Sequence
	F11  _dafny.Sequence
	F13  bool
	F14  _dafny.Int
	F15  bool
	F16  _dafny.Map
	F17  _dafny.Array
	F18  bool
	F19  _dafny.Int
	F20  _dafny.Sequence
	F21  _dafny.Int
	F28  bool
	_f0  bool
	_f4  bool
	_f5  _dafny.MultiSet
	_f6  _dafny.Int
	_f8  _dafny.Int
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f12 _dafny.Int
	_f22 _dafny.Map
	_f23 _dafny.Int
	_f24 bool
	_f25 _dafny.Sequence
	_f26 _dafny.Int
	_f27 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F2 = _dafny.Zero
	_this.F3 = _dafny.EmptySeq
	_this.F7 = _dafny.EmptySeq
	_this.F11 = _dafny.EmptySeq
	_this.F13 = false
	_this.F14 = _dafny.Zero
	_this.F15 = false
	_this.F16 = _dafny.EmptyMap
	_this.F17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F18 = false
	_this.F19 = _dafny.Zero
	_this.F20 = _dafny.EmptySeq
	_this.F21 = _dafny.Zero
	_this.F28 = false
	_this._f0 = false
	_this._f4 = false
	_this._f5 = _dafny.EmptyMultiSet
	_this._f6 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f22 = _dafny.EmptyMap
	_this._f23 = _dafny.Zero
	_this._f24 = false
	_this._f25 = _dafny.EmptySeq
	_this._f26 = _dafny.Zero
	_this._f27 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Sequence, f4 bool, f5 _dafny.MultiSet, f6 _dafny.Int, f7 _dafny.Sequence, f8 _dafny.Int, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Sequence, f12 _dafny.Int, f13 bool, f14 _dafny.Int, f15 bool, f16 _dafny.Map, f17 _dafny.Array, f18 bool, f19 _dafny.Int, f20 _dafny.Sequence, f21 _dafny.Int, f22 _dafny.Map, f23 _dafny.Int, f24 bool, f25 _dafny.Sequence, f26 _dafny.Int, f27 bool, f28 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this).F14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
		(_this).F17 = f17
		(_this).F18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f27 = f27
		(_this).F28 = f28
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.MultiSet {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
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
func (_this *GlobalState) F22() _dafny.Map {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Int {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() bool {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F25() _dafny.Sequence {
	{
		return _this._f25
	}
}
func (_this *GlobalState) F26() _dafny.Int {
	{
		return _this._f26
	}
}
func (_this *GlobalState) F27() bool {
	{
		return _this._f27
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F37 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F37 = false
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

func (_this *C0) Ctor__(f37 bool) {
	{
		(_this).F37 = f37
	}
}
func (_this *C0) Fm11(p0 bool, globalState *GlobalState) bool {
	{
		return _this.F37
	}
}
func (_this *C0) Fm12(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(-345)).Minus((func() _dafny.Int {
			if _this.F37 {
				return _dafny.IntOfInt64(696)
			}
			return _dafny.IntOfInt64(343)
		})())
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f40 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f40 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f40 bool) {
	{
		(_this)._f40 = f40
	}
}
func (_this *C1) Fm30(p0 D5, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ovq")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-114)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F40(), _dafny.IntOfInt64(-638))).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, !((_this).F40()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('n'), _dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('r'))).Cardinality()), _dafny.IntOfInt64(141)))
	}
}
func (_this *C1) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _373_v0 _dafny.Set
		_ = _373_v0
		_373_v0 = _dafny.SetOf(p0)
		var _374_v1 D6
		_ = _374_v1
		_374_v1 = Companion_D6_.Create_DC21_(_373_v0)
		var _375_v2 _dafny.Sequence
		_ = _375_v2
		_375_v2 = _dafny.SeqOf(_374_v1)
		var _376_v3 D8
		_ = _376_v3
		_376_v3 = Companion_D8_.Create_DC25_(_375_v2)
		var _377_v4 _dafny.Int
		_ = _377_v4
		_377_v4 = _dafny.IntOfInt64(651)
		var _378_v5 _dafny.MultiSet
		_ = _378_v5
		_378_v5 = _dafny.MultiSetOf((_376_v3).Dtor_cf39(), Companion_Default___.Fm31(_377_v4, _dafny.IntOfInt64(241), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(660))).Uint32(), func(coer52 func(_dafny.Int) D6) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}((func(_379_v1 D6) func(_dafny.Int) D6 {
			return func(_380_i0 _dafny.Int) D6 {
				return _379_v1
			}
		})(_374_v1))), _375_v2)
		var _381_v6 _dafny.Map
		_ = _381_v6
		_381_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v4, _dafny.SetOf(p0, (_this).F40()))
		(globalState).F18 = !((_378_v5).IsDisjointFrom(_dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_375_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.IntOfUint32((_375_v2).Cardinality()))).Uint32(), Companion_D6_.Create_DC21_((func() _dafny.Set {
			if (_381_v6).Contains(_377_v4) {
				return (_381_v6).Get(_377_v4).(_dafny.Set)
			}
			return _dafny.SetOf((_this).F40())
		})())))))
		var _382_v7 _dafny.Map
		_ = _382_v7
		_382_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F40(), p0)
		var _383_v8 _dafny.Sequence
		_ = _383_v8
		_383_v8 = _dafny.SeqOf(_377_v4)
		var _384_v9 _dafny.Sequence
		_ = _384_v9
		_384_v9 = _dafny.SeqOf((_382_v7).Merge(_382_v7), Companion_Default___.Fm32(p0, Companion_Default___.Fm33(true, _dafny.IntOfInt64(-164), _377_v4, globalState), _383_v8, globalState), _382_v7)
		var _385_v10 _dafny.Sequence
		_ = _385_v10
		_385_v10 = _dafny.SeqOf(_384_v9, _dafny.Companion_Sequence_.Concatenate(_384_v9, _384_v9))
		_384_v9 = (_385_v10).Select((Companion_Default___.SafeIndex((_377_v4).Plus(_377_v4), _dafny.IntOfUint32((_385_v10).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _386_v11 _dafny.MultiSet
		_ = _386_v11
		_386_v11 = _dafny.MultiSetOf(_dafny.CodePoint('j'))
		var _387_v12 _dafny.Sequence
		_ = _387_v12
		_387_v12 = _dafny.UnicodeSeqOfUtf8Bytes("xo")
		var _388_v13 _dafny.Array
		_ = _388_v13
		var _nwElement0_7 bool = (_386_v11).IsDisjointFrom(_386_v11)
		_ = _nwElement0_7
		var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(17))
		_ = _nw61
		(_nw61).ArraySet1(_nwElement0_7, 0)
		(_nw61).ArraySet1(!((_this).F40()), 1)
		(_nw61).ArraySet1(p0, 2)
		(_nw61).ArraySet1(p0, 3)
		(_nw61).ArraySet1((_this).F40(), 4)
		(_nw61).ArraySet1((func() bool {
			if (_this).F40() {
				return (_this).F40()
			}
			return (_this).F40()
		})(), 5)
		(_nw61).ArraySet1(p0, 6)
		(_nw61).ArraySet1(p0, 7)
		(_nw61).ArraySet1((_this).F40(), 8)
		(_nw61).ArraySet1((_this).F40(), 9)
		(_nw61).ArraySet1(Companion_Default___.Fm33(p0, _dafny.IntOfInt64(364), _dafny.IntOfInt64(-247), globalState), 10)
		(_nw61).ArraySet1(p0, 11)
		(_nw61).ArraySet1((_this).F40(), 12)
		(_nw61).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(286))).Cardinality())).Cmp(_dafny.IntOfUint32((_387_v12).Cardinality())) < 0, 13)
		(_nw61).ArraySet1((_this).F40(), 14)
		(_nw61).ArraySet1(false, 15)
		(_nw61).ArraySet1((p0) || (p0), 16)
		_388_v13 = _nw61
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))
		_ = _index50
		(_388_v13).ArraySet1((_this).F40(), (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))
		_ = _index51
		(_388_v13).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_387_v12, _387_v12), (_index51).Int())
		var _389_v14 _dafny.Array
		_ = _389_v14
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw62
		_389_v14 = _nw62
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_389_v14), 0))
		_ = _index52
		(_389_v14).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-162))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg53 _dafny.Int) interface{} {
				return coer53(arg53)
			}
		}(func(_390_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_389_v14), 0))
		_ = _index53
		(_389_v14).ArraySet1(Companion_Default___.Fm34(_377_v4, _377_v4, p0, p0, globalState), (_index53).Int())
		var _391_v15 _dafny.Sequence
		_ = _391_v15
		_391_v15 = _dafny.SeqOf((_388_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))).Int()).(bool), ((_this).F40()) || (!((_this).F40())), ((_388_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))).Int()).(bool)) && ((_this).F40()), true, (_this).F40())
		var _392_v16 D4
		_ = _392_v16
		_392_v16 = Companion_D4_.Create_DC15_(_377_v4, _377_v4)
		var _393_v17 _dafny.CodePoint
		_ = _393_v17
		_393_v17 = _dafny.CodePoint('i')
		var _394_v18 D2
		_ = _394_v18
		_394_v18 = Companion_D2_.Create_DC8_(_393_v17, (_388_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))).Int()).(bool))
		var _395_v19 _dafny.Sequence
		_ = _395_v19
		_395_v19 = _dafny.SeqOf(Companion_D2_.Create_DC8_(_393_v17, p0), _394_v18, _394_v18)
		var _rhs58 bool = (func() bool {
			if (_382_v7).Contains((_388_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))).Int()).(bool)) {
				return (_382_v7).Get((_388_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_388_v13), 0))).Int()).(bool)).(bool)
			}
			return true
		})()
		_ = _rhs58
		var _rhs59 _dafny.Int = _377_v4
		_ = _rhs59
		var _rhs60 _dafny.Sequence = Companion_Default___.Fm35(_dafny.Companion_Sequence_.Concatenate(_395_v19, _395_v19), _377_v4, globalState)
		_ = _rhs60
		var _rhs61 D4 = _392_v16
		_ = _rhs61
		var _lhs49 *GlobalState = globalState
		_ = _lhs49
		var _lhs50 *GlobalState = globalState
		_ = _lhs50
		_lhs49.F28 = _rhs58
		_lhs50.F19 = _rhs59
		_391_v15 = _rhs60
		_392_v16 = _rhs61
		var _396_v20 *C0
		_ = _396_v20
		var _nw63 *C0 = New_C0_()
		_ = _nw63
		_nw63.Ctor__(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("rrjunsmd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg54 _dafny.Int) interface{} {
				return coer54(arg54)
			}
		}((func(_397_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_398_i2 _dafny.Int) _dafny.CodePoint {
				return _397_v17
			}
		})(_393_v17)))))
		_396_v20 = _nw63
		var _399_v21 _dafny.Sequence
		_ = _399_v21
		_399_v21 = _dafny.SeqOf(_391_v15)
		r0 = _399_v21
		r1 = p0
		return r0, r1
	}
}
func (_this *C1) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _400_v0 _dafny.MultiSet
		_ = _400_v0
		_400_v0 = _dafny.MultiSetOf((_dafny.Zero).Minus(p2))
		var _401_v1 _dafny.Sequence
		_ = _401_v1
		_401_v1 = _dafny.SeqOf(p2, p2, p2)
		var _402_v2 _dafny.Map
		_ = _402_v2
		_402_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_400_v0).Intersection(_dafny.MultiSetOf(p2)), _dafny.Companion_Sequence_.Concatenate(_401_v1, _401_v1))
		var _403_v3 _dafny.Array
		_ = _403_v3
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_6
		var _nw64 _dafny.Array
		_ = _nw64
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw64 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) bool = (func(_404_p0 bool) func(_dafny.Int) bool {
				return func(_405_i0 _dafny.Int) bool {
					return _404_p0
				}
			})(p0)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw64 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw64).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw64).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_403_v3 = _nw64
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
		_ = _index54
		(_403_v3).ArraySet1(!((_this).F40()), (_index54).Int())
		var _406_v4 _dafny.Map
		_ = _406_v4
		_406_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F40(), p0)
		var _407_v5 D9
		_ = _407_v5
		_407_v5 = Companion_D9_.Create_DC27_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}(func(_408_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(102)
		})))
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
		_ = _index55
		var _rhs62 bool = (_this).F40()
		_ = _rhs62
		var _rhs63 bool = p0
		_ = _rhs63
		var _rhs64 _dafny.Map = ((_402_v2).Update(_dafny.MultiSetOf((_406_v4).Cardinality(), p2, p2, p2, _dafny.IntOfInt64(478)), (_407_v5).Dtor_cf40())).Merge(_402_v2)
		_ = _rhs64
		var _rhs65 bool = !(!(p0)) || (p0)
		_ = _rhs65
		var _lhs51 *GlobalState = globalState
		_ = _lhs51
		var _lhs52 *GlobalState = globalState
		_ = _lhs52
		var _lhs53 _dafny.Array = _403_v3
		_ = _lhs53
		var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
		_ = _lhs54
		_lhs51.F18 = _rhs62
		_lhs52.F18 = _rhs63
		_402_v2 = _rhs64
		(_lhs53).ArraySet1(_rhs65, (_lhs54).Int())
		if !((_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool)) || (!((_this).F40())) {
			var _409_v6 D4
			_ = _409_v6
			_409_v6 = Companion_D4_.Create_DC14_((_dafny.Zero).Minus(p2))
			var _pat_let_tv9 = p2
			_ = _pat_let_tv9
			_409_v6 = func(_pat_let4_0 D4) D4 {
				return func(_410_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let5_0 _dafny.Int) D4 {
						return func(_411_dt__update_hcf23_h0 _dafny.Int) D4 {
							return Companion_D4_.Create_DC14_(_411_dt__update_hcf23_h0)
						}(_pat_let5_0)
					}(_pat_let_tv9)
				}(_pat_let4_0)
			}(_409_v6)
			var _412_v7 *C0
			_ = _412_v7
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__((_this).F40())
			_412_v7 = _nw65
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
			_ = _index56
			(_403_v3).ArraySet1(!(_412_v7.F37), (_index56).Int())
			var _413_v8 D0
			_ = _413_v8
			_413_v8 = Companion_D0_.Create_DC1_(_403_v3, p2, p2)
			var _414_v9 _dafny.Array
			_ = _414_v9
			var _nwElement0_8 _dafny.Array = (_413_v8).Dtor_cf1()
			_ = _nwElement0_8
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(9))
			_ = _nw66
			(_nw66).ArraySet1(_nwElement0_8, 0)
			(_nw66).ArraySet1(_403_v3, 1)
			(_nw66).ArraySet1((_413_v8).Dtor_cf1(), 2)
			(_nw66).ArraySet1(_403_v3, 3)
			(_nw66).ArraySet1(_403_v3, 4)
			(_nw66).ArraySet1((func() _dafny.Array {
				if _412_v7.F37 {
					return _403_v3
				}
				return _403_v3
			})(), 5)
			(_nw66).ArraySet1(_403_v3, 6)
			(_nw66).ArraySet1(_403_v3, 7)
			(_nw66).ArraySet1(_403_v3, 8)
			_414_v9 = _nw66
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_414_v9), 0))
			_ = _index57
			(_414_v9).ArraySet1(_403_v3, (_index57).Int())
			var _415_v10 _dafny.Sequence
			_ = _415_v10
			_415_v10 = _dafny.UnicodeSeqOfUtf8Bytes("eakffkqhw")
			var _416_v11 _dafny.Sequence
			_ = _416_v11
			_416_v11 = _dafny.SeqOf(_415_v10, _dafny.UnicodeSeqOfUtf8Bytes("dkbchhlud"), _dafny.Companion_Sequence_.Concatenate(_415_v10, _415_v10))
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_414_v9), 0))
			_ = _index58
			var _rhs66 _dafny.Array = (func() _dafny.Array {
				if (_this).F40() {
					return _403_v3
				}
				return _403_v3
			})()
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _415_v10
			_ = _rhs67
			var _rhs68 _dafny.Int = _dafny.IntOfUint32((_416_v11).Cardinality())
			_ = _rhs68
			var _rhs69 bool = _412_v7.F37
			_ = _rhs69
			var _rhs70 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(101)).Times(p2), Companion_Default___.SafeDivisionInt((_412_v7).Fm12(globalState), p2))
			_ = _rhs70
			var _lhs55 _dafny.Array = _414_v9
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_414_v9), 0))
			_ = _lhs56
			var _lhs57 *GlobalState = globalState
			_ = _lhs57
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			(_lhs55).ArraySet1(_rhs66, (_lhs56).Int())
			_lhs57.F20 = _rhs67
			_lhs58.F19 = _rhs68
			_lhs59.F28 = _rhs69
			_lhs60.F21 = _rhs70
			var _417_v12 _dafny.Map
			_ = _417_v12
			_417_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v1, p0)
			var _418_v14 _dafny.MultiSet
			_ = _418_v14
			_418_v14 = _dafny.MultiSetOf(_401_v1, _dafny.SeqOf(p2))
			var _419_v15 *C0
			_ = _419_v15
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__((_417_v12).Equals(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter27 := _dafny.Iterate((_418_v14).Elements()); ; {
					_compr_25, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _420_v13 _dafny.Sequence
					_420_v13 = interface{}(_compr_25).(_dafny.Sequence)
					if (_418_v14).Contains(_420_v13) {
						_coll25.Add(_420_v13, !(_412_v7.F37))
					}
				}
				return _coll25.ToMap()
			}()))
			_419_v15 = _nw67
		} else {
			var _421_v16 _dafny.Set
			_ = _421_v16
			_421_v16 = _dafny.SetOf(p2)
			var _422_v17 _dafny.CodePoint
			_ = _422_v17
			_422_v17 = _dafny.CodePoint('q')
			var _423_v18 _dafny.Sequence
			_ = _423_v18
			_423_v18 = _dafny.SeqOf(_422_v17, _422_v17, _dafny.CodePoint('j'), _422_v17)
			var _nwElement0_9 _dafny.Int = p2
			_ = _nwElement0_9
			var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(21))
			_ = _nw68
			(_nw68).ArraySet1(_nwElement0_9, 0)
			(_nw68).ArraySet1((Companion_Default___.Fm36(globalState)).Plus(Companion_Default___.Fm36(globalState)), 1)
			(_nw68).ArraySet1((_421_v16).Cardinality(), 2)
			(_nw68).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p2), 3)
			(_nw68).ArraySet1(p2, 4)
			(_nw68).ArraySet1(p2, 5)
			(_nw68).ArraySet1((_dafny.IntOfInt64(136)).Minus(p2), 6)
			(_nw68).ArraySet1(p2, 7)
			(_nw68).ArraySet1(p2, 8)
			(_nw68).ArraySet1(p2, 9)
			(_nw68).ArraySet1(_dafny.IntOfInt64(-137), 10)
			(_nw68).ArraySet1(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(389)), 11)
			(_nw68).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 12)
			(_nw68).ArraySet1(p2, 13)
			(_nw68).ArraySet1(p2, 14)
			(_nw68).ArraySet1(p2, 15)
			(_nw68).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_423_v18, _423_v18)).Cardinality()), 16)
			(_nw68).ArraySet1((_401_v1).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_401_v1).Cardinality()))).Uint32()).(_dafny.Int), 17)
			(_nw68).ArraySet1(p2, 18)
			(_nw68).ArraySet1((func() _dafny.Int {
				if (_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool) {
					return p2
				}
				return _dafny.IntOfInt64(241)
			})(), 19)
			(_nw68).ArraySet1(p2, 20)
			(globalState).F17 = _nw68
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
			_ = _index59
			(_403_v3).ArraySet1((p2).Cmp(p2) > 0, (_index59).Int())
			var _424_v19 _dafny.Set
			_ = _424_v19
			_424_v19 = _dafny.SetOf((_this).F40(), p0, (_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool), (_this).F40())
			var _425_v20 _dafny.Map
			_ = _425_v20
			_425_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
			var _426_v21 _dafny.Array
			_ = _426_v21
			var _nwElement0_10 _dafny.Int = ((_424_v19).Difference(_424_v19)).Cardinality()
			_ = _nwElement0_10
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(9))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_10, 0)
			(_nw69).ArraySet1(Companion_Default___.SafeDivisionInt(p2, (_dafny.Zero).Minus(Companion_Default___.Fm36(globalState))), 1)
			(_nw69).ArraySet1((func() _dafny.Int {
				if (_this).F40() {
					return _dafny.IntOfInt64(619)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_423_v18, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_423_v18).Cardinality()))).Uint32(), _422_v17)).Cardinality())
			})(), 2)
			(_nw69).ArraySet1((Companion_D4_.Create_DC15_(p2, _dafny.IntOfUint32((_423_v18).Cardinality()))).Dtor_cf24(), 3)
			(_nw69).ArraySet1(p2, 4)
			(_nw69).ArraySet1((_dafny.Zero).Minus((_425_v20).Cardinality()), 5)
			(_nw69).ArraySet1(_dafny.IntOfInt64(773), 6)
			(_nw69).ArraySet1(p2, 7)
			(_nw69).ArraySet1(Companion_Default___.Fm36(globalState), 8)
			_426_v21 = _nw69
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
			_ = _index60
			var _rhs71 _dafny.Array = _426_v21
			_ = _rhs71
			var _rhs72 bool = (Companion_D6_.Create_DC23_((_this).F40(), p2)).Dtor_cf36()
			_ = _rhs72
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			var _lhs62 _dafny.Array = _403_v3
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
			_ = _lhs63
			_lhs61.F17 = _rhs71
			(_lhs62).ArraySet1(_rhs72, (_lhs63).Int())
			if (_this).F40() {
				var _427_v22 _dafny.Set
				_ = _427_v22
				_427_v22 = _dafny.SetOf(_422_v17)
				var _428_v23 _dafny.MultiSet
				_ = _428_v23
				_428_v23 = _dafny.MultiSetOf(_427_v22, _dafny.SetOf(_422_v17), _427_v22, _427_v22)
				var _429_v24 _dafny.Sequence
				_ = _429_v24
				_429_v24 = _dafny.SeqOf(_427_v22)
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))
				_ = _index61
				(_403_v3).ArraySet1(((_dafny.MultiSetFromSeq(_429_v24)).Intersection(_428_v23)).IsProperSubsetOf(_428_v23), (_index61).Int())
				var _430_v25 _dafny.Sequence
				_ = _430_v25
				_430_v25 = _dafny.SeqOf(!(_424_v19).Equals(_dafny.SetOf((_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool))), ((_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool)) || (p0))
				(globalState).F13 = (_430_v25).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_430_v25).Cardinality()))).Uint32()).(bool)
				var _431_v26 *C0
				_ = _431_v26
				var _nw70 *C0 = New_C0_()
				_ = _nw70
				_nw70.Ctor__((p2).Cmp(_dafny.IntOfInt64(981)) < 0)
				_431_v26 = _nw70
				var _rhs73 _dafny.Int = _dafny.IntOfInt64(-164)
				_ = _rhs73
				var _rhs74 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _425_v20)).Contains(p2)
				_ = _rhs74
				var _rhs75 _dafny.Int = p2
				_ = _rhs75
				var _lhs64 *GlobalState = globalState
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				var _lhs66 *GlobalState = globalState
				_ = _lhs66
				_lhs64.F19 = _rhs73
				_lhs65.F28 = _rhs74
				_lhs66.F1 = _rhs75
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_426_v21), 0))
				_ = _index62
				(_426_v21).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool) {
						return _430_v25
					}
					return _430_v25
				})()).Cardinality()), (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_426_v21), 0))
				_ = _index63
				var _rhs76 _dafny.Array = _403_v3
				_ = _rhs76
				var _rhs77 _dafny.Sequence = (func() _dafny.Sequence {
					if ((_dafny.Zero).Minus(_dafny.IntOfUint32((_430_v25).Cardinality()))).Cmp(p2) >= 0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg56 _dafny.Int) interface{} {
								return coer56(arg56)
							}
						}((func(_432_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_433_i2 _dafny.Int) _dafny.CodePoint {
								return _432_v17
							}
						})(_422_v17)))
					}
					return (func() _dafny.Sequence {
						if _431_v26.F37 {
							return _dafny.Companion_Sequence_.Update(_423_v18, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_423_v18).Cardinality()))).Uint32(), _422_v17)
						}
						return _423_v18
					})()
				})()
				_ = _rhs77
				var _rhs78 _dafny.Int = _dafny.IntOfInt64(-296)
				_ = _rhs78
				var _rhs79 bool = !(p0)
				_ = _rhs79
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				var _lhs68 _dafny.Array = _426_v21
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_426_v21), 0))
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				_403_v3 = _rhs76
				_lhs67.F20 = _rhs77
				(_lhs68).ArraySet1(_rhs78, (_lhs69).Int())
				_lhs70.F28 = _rhs79
			} else {
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_426_v21), 0))
				_ = _index64
				(_426_v21).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm36(globalState)), (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_426_v21), 0))
				_ = _index65
				(_426_v21).ArraySet1(p2, (_index65).Int())
				(globalState).F28 = (_this).F40()
				var _434_v27 _dafny.Sequence
				_ = _434_v27
				_434_v27 = _dafny.SeqOf((Companion_Default___.Fm36(globalState)).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-877))) >= 0)
				var _435_v28 _dafny.Map
				_ = _435_v28
				_435_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_426_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_426_v21), 0))).Int()).(_dafny.Int), _434_v27)
				_434_v27 = (func() _dafny.Sequence {
					if !(p0) {
						return _434_v27
					}
					return (func() _dafny.Sequence {
						if (_435_v28).Contains(p2) {
							return (_435_v28).Get(p2).(_dafny.Sequence)
						}
						return _434_v27
					})()
				})()
				_425_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_426_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_426_v21), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf(p0)).IsSubsetOf(_dafny.MultiSetOf(p0)))
				var _436_v29 _dafny.Array
				_ = _436_v29
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_7
				var _nw71 _dafny.Array
				_ = _nw71
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw71 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Int = (func(_437_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_438_i3 _dafny.Int) _dafny.Int {
							return (_438_i3).Minus(_437_p2)
						}
					})(p2)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw71 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw71).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw71).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_436_v29 = _nw71
				var _rhs80 _dafny.CodePoint = _422_v17
				_ = _rhs80
				var _rhs81 _dafny.Array = _436_v29
				_ = _rhs81
				var _rhs82 _dafny.Int = (_426_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_426_v21), 0))).Int()).(_dafny.Int)
				_ = _rhs82
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				_422_v17 = _rhs80
				_lhs71.F17 = _rhs81
				_lhs72.F14 = _rhs82
			}
			(globalState).F13 = p0
		}
		var _439_v30 D3
		_ = _439_v30
		_439_v30 = Companion_D3_.Create_DC11_((_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool), !(p0))
		var _440_i4 _dafny.Int
		_ = _440_i4
		_440_i4 = _dafny.Zero
		{
			for (_439_v30).Dtor_cf20() {
				{
					if (_440_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_440_i4 = (_440_i4).Plus(_dafny.One)
					var _441_v31 _dafny.Array
					_ = _441_v31
					var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
					_ = _nw72
					_441_v31 = _nw72
					var _442_v32 _dafny.Map
					_ = _442_v32
					_442_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v31, (_407_v5).Dtor_cf40())
					_442_v32 = (_442_v32).Update(_441_v31, _dafny.SeqOf(p2, p2))
					(globalState).F18 = (_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool)
					(globalState).F14 = (p2).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(706))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg57 _dafny.Int) interface{} {
							return coer57(arg57)
						}
					}(func(_443_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality()))
					var _444_v33 _dafny.Sequence
					_ = _444_v33
					_444_v33 = _dafny.UnicodeSeqOfUtf8Bytes("evxiub")
					var _445_v34 _dafny.MultiSet
					_ = _445_v34
					_445_v34 = _dafny.MultiSetOf(false)
					var _446_v35 _dafny.CodePoint
					_ = _446_v35
					_446_v35 = _dafny.CodePoint('p')
					var _447_v36 D5
					_ = _447_v36
					_447_v36 = Companion_D5_.Create_DC19_(_dafny.IntOfUint32((_444_v33).Cardinality()), (_401_v1).Select((Companion_Default___.SafeIndex((_445_v34).Cardinality(), _dafny.IntOfUint32((_401_v1).Cardinality()))).Uint32()).(_dafny.Int), _446_v35)
					(globalState).F13 = (p2).Cmp((_447_v36).Dtor_cf28()) != 0
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _448_v37 *C0
		_ = _448_v37
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__((_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool))
		_448_v37 = _nw73
		var _449_v38 _dafny.Map
		_ = _449_v38
		_449_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _450_v39 _dafny.Map
		_ = _450_v39
		_450_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v38, (Companion_Default___.Fm37(globalState)).Dtor_cf33())
		var _451_v40 _dafny.Sequence
		_ = _451_v40
		_451_v40 = _dafny.UnicodeSeqOfUtf8Bytes("bobyhh")
		var _452_v41 _dafny.CodePoint
		_ = _452_v41
		_452_v41 = _dafny.CodePoint('t')
		var _rhs83 bool = !(_450_v39).Contains((_449_v38).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm33(p0, p2, _dafny.IntOfInt64(945), globalState), p2)))
		_ = _rhs83
		var _rhs84 bool = _448_v37.F37
		_ = _rhs84
		var _rhs85 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_451_v40, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32(((Companion_Default___.Fm38(globalState)).Dtor_cf40()).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_451_v40).Cardinality()))).Uint32(), _452_v41), _451_v40), _451_v40)
		_ = _rhs85
		var _lhs73 *GlobalState = globalState
		_ = _lhs73
		var _lhs74 *GlobalState = globalState
		_ = _lhs74
		_lhs73.F18 = _rhs83
		_lhs74.F15 = _rhs84
		r1 = _rhs85
		var _hi0 _dafny.Int = p2
		_ = _hi0
		for _453_i6 := p2; _453_i6.Cmp(_hi0) < 0; _453_i6 = _453_i6.Plus(_dafny.One) {
			(globalState).F1 = (p2).Times(p2)
			var _454_v42 _dafny.Array
			_ = _454_v42
			var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw74
			_454_v42 = _nw74
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_454_v42), 0))
			_ = _index66
			(_454_v42).ArraySet1(p2, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_454_v42), 0))
			_ = _index67
			(_454_v42).ArraySet1(Companion_Default___.Fm36(globalState), (_index67).Int())
			var _455_v43 D6
			_ = _455_v43
			_455_v43 = Companion_D6_.Create_DC22_(p2, (_449_v38).Cardinality(), (_403_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_403_v3), 0))).Int()).(bool))
			var _456_v44 _dafny.Map
			_ = _456_v44
			_456_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v3, (_455_v43).Dtor_cf33())
			_449_v38 = (_449_v38).Update((p2).Cmp(((_456_v44).Update(_403_v3, _453_i6)).Cardinality()) < 0, (_454_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_454_v42), 0))).Int()).(_dafny.Int))
			_400_v0 = (_400_v0).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_454_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_454_v42), 0))).Int()).(_dafny.Int), _449_v38)).Cardinality(), Companion_Default___.Abs(_453_i6))
		}
		r0 = _451_v40
		r1 = !(false)
		return r0, r1
	}
}
func (_this *C1) F40() bool {
	{
		return _this._f40
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f30 _dafny.Sequence
	_f29 bool
	_f38 _dafny.Int
	_f39 _dafny.Set
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f30 = _dafny.EmptySeq
	_this._f29 = false
	_this._f38 = _dafny.Zero
	_this._f39 = _dafny.EmptySet
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C2) F30_set_(value _dafny.Sequence) {
	_this._f30 = value
}
func (_this *C2) F29() bool {
	return _this._f29
}
func (_this *C2) Ctor__(f38 _dafny.Int, f39 _dafny.Set, f29 bool, f30 _dafny.Sequence) {
	{
		(_this)._f38 = f38
		(_this)._f39 = f39
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C2) Fm16(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F38()
	}
}
func (_this *C2) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		(globalState).F28 = (_this).F29()
		if !((_this).F29()) || (((_this).F38()).Cmp((_dafny.Zero).Minus((_this).F38())) > 0) {
			(globalState).F15 = (_this).F29()
			(globalState).F2 = ((_this).F38()).Minus((_this).F38())
			var _457_v0 _dafny.MultiSet
			_ = _457_v0
			_457_v0 = _dafny.MultiSetOf(p0)
			var _458_v1 D4
			_ = _458_v1
			_458_v1 = Companion_D4_.Create_DC15_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ewafp")).Cardinality()), _dafny.IntOfInt64(469))
			var _459_v2 _dafny.Array
			_ = _459_v2
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw75
			_459_v2 = _nw75
			var _460_v3 _dafny.MultiSet
			_ = _460_v3
			_460_v3 = _dafny.MultiSetOf(_459_v2, _459_v2)
			var _461_v4 _dafny.Sequence
			_ = _461_v4
			_461_v4 = _dafny.SeqOf(p0)
			var _462_v5 _dafny.Array
			_ = _462_v5
			var _nwElement0_11 _dafny.Int = (_this).F38()
			_ = _nwElement0_11
			var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(27))
			_ = _nw76
			(_nw76).ArraySet1(_nwElement0_11, 0)
			(_nw76).ArraySet1(_dafny.IntOfInt64(838), 1)
			(_nw76).ArraySet1((_this).F38(), 2)
			(_nw76).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F38(), (_this).F38()), 3)
			(_nw76).ArraySet1(_dafny.IntOfInt64(13), 4)
			(_nw76).ArraySet1((_this).Fm16(_457_v0, globalState), 5)
			(_nw76).ArraySet1((_this).F38(), 6)
			(_nw76).ArraySet1((_this).F38(), 7)
			(_nw76).ArraySet1((_this).F38(), 8)
			(_nw76).ArraySet1((_this).F38(), 9)
			(_nw76).ArraySet1((_dafny.Zero).Minus((_this).F38()), 10)
			(_nw76).ArraySet1((_458_v1).Dtor_cf25(), 11)
			(_nw76).ArraySet1((_this).F38(), 12)
			(_nw76).ArraySet1((_this).F38(), 13)
			(_nw76).ArraySet1(Companion_Default___.SafeModuloInt((_this).F38(), (Companion_Default___.Fm17((_this).F38(), globalState)).Cardinality()), 14)
			(_nw76).ArraySet1((_this).F38(), 15)
			(_nw76).ArraySet1((_this).Fm16(_457_v0, globalState), 16)
			(_nw76).ArraySet1(((Companion_Default___.Fm18((_this).F38(), (_460_v3).Cardinality(), (_this).F38(), _dafny.IntOfUint32((_461_v4).Cardinality()), globalState)).Dtor_cf15()).Times((_this).F38()), 17)
			(_nw76).ArraySet1((_this).F38(), 18)
			(_nw76).ArraySet1(_dafny.IntOfUint32((_461_v4).Cardinality()), 19)
			(_nw76).ArraySet1((_this).F38(), 20)
			(_nw76).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F29())).Cardinality()), 21)
			(_nw76).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F38(), (_this).F38()), 22)
			(_nw76).ArraySet1((func() _dafny.Int {
				if p0 {
					return (_this).F38()
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_463_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality())
			})(), 23)
			(_nw76).ArraySet1((_this).F38(), 24)
			(_nw76).ArraySet1((_this).F38(), 25)
			(_nw76).ArraySet1((_this).F38(), 26)
			_462_v5 = _nw76
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_462_v5), 0))
			_ = _index68
			(_462_v5).ArraySet1((_this).F38(), (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_462_v5), 0))
			_ = _index69
			(_462_v5).ArraySet1((_this).F38(), (_index69).Int())
			(globalState).F14 = (_this).F38()
			var _464_v6 _dafny.Sequence
			_ = _464_v6
			_464_v6 = _dafny.SeqOf(_459_v2)
			(globalState).F1 = _dafny.IntOfUint32((_464_v6).Cardinality())
		} else {
			if p0 {
				var _465_v7 *C0
				_ = _465_v7
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__(p0)
				_465_v7 = _nw77
				var _466_v8 _dafny.MultiSet
				_ = _466_v8
				_466_v8 = _dafny.MultiSetOf(_465_v7.F37)
				var _467_v9 _dafny.Array
				_ = _467_v9
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_8
				var _nw78 _dafny.Array
				_ = _nw78
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw78 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = func(_468_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_468_i1, (_this).F38())
					}
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw78 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw78).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw78).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_467_v9 = _nw78
				var _469_v10 _dafny.Map
				_ = _469_v10
				_469_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F30()).Select((Companion_Default___.SafeIndex((_this).Fm16(_466_v8, globalState), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32()).(_dafny.CodePoint), _467_v9)
				var _470_v11 _dafny.CodePoint
				_ = _470_v11
				_470_v11 = _dafny.CodePoint('e')
				(globalState).F17 = (func() _dafny.Array {
					if (_469_v10).Contains(_470_v11) {
						return (_469_v10).Get(_470_v11).(_dafny.Array)
					}
					return _467_v9
				})()
				var _471_v12 _dafny.Array
				_ = _471_v12
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_9
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = (func(_472_p0 bool, _473_v7 *C0) func(_dafny.Int) bool {
						return func(_474_i2 _dafny.Int) bool {
							return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.SetOf(_472_p0, _472_p0)), _dafny.SeqOf(_dafny.SetOf(_473_v7.F37), _dafny.SetOf(true), _dafny.SetOf(_473_v7.F37, _472_p0), _dafny.SetOf(_473_v7.F37), _dafny.SetOf((_this).F29())))
						}
					})(p0, _465_v7)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw79 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw79).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw79).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_471_v12 = _nw79
				_471_v12 = _471_v12
				(globalState).F14 = (_this).F38()
				var _475_v13 _dafny.Map
				_ = _475_v13
				_475_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_470_v11, _470_v11)
				var _476_v14 _dafny.Sequence
				_ = _476_v14
				_476_v14 = _dafny.SeqOf(_475_v13)
				_476_v14 = _476_v14
			} else {
				(globalState).F28 = p0
				(globalState).F28 = ((_this).F38()).Cmp((_this).F38()) > 0
				var _477_v15 _dafny.MultiSet
				_ = _477_v15
				_477_v15 = _dafny.MultiSetOf((_this).F38(), (_this).F38(), (_this).F38(), (_this).F38(), (_this).F38())
				var _478_v16 D3
				_ = _478_v16
				_478_v16 = Companion_D3_.Create_DC10_(_477_v15)
				var _479_v17 _dafny.CodePoint
				_ = _479_v17
				_479_v17 = _dafny.CodePoint('i')
				var _480_v18 *C0
				_ = _480_v18
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__((_this).F29())
				_480_v18 = _nw80
				var _481_v19 _dafny.Map
				_ = _481_v19
				_481_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F30()).Cardinality()), _480_v18)
				var _482_v20 _dafny.Set
				_ = _482_v20
				_482_v20 = _dafny.SetOf((_481_v19).Cardinality(), _dafny.IntOfInt64(-408))
				var _483_v21 _dafny.Sequence
				_ = _483_v21
				_483_v21 = _dafny.SeqOf(p0)
				var _484_v22 _dafny.Sequence
				_ = _484_v22
				_484_v22 = _dafny.SeqOf((_482_v20).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_480_v18.F37)).Cardinality()), _dafny.IntOfInt64(-607), _dafny.IntOfInt64(730), _dafny.IntOfUint32((_483_v21).Cardinality()))
				var _rhs86 bool = (_dafny.MultiSetOf((_this).F38(), (_this).F38(), _dafny.IntOfInt64(283))).IsDisjointFrom((func() _dafny.MultiSet {
					if Companion_Default___.Fm19(_479_v17, globalState) {
						return _dafny.MultiSetFromSeq(_484_v22)
					}
					return _477_v15
				})())
				_ = _rhs86
				var _rhs87 bool = _480_v18.F37
				_ = _rhs87
				var _rhs88 D3 = _478_v16
				_ = _rhs88
				var _rhs89 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(817))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}(func(_485_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				}))
				_ = _rhs89
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				_lhs75.F28 = _rhs86
				_lhs76.F15 = _rhs87
				_478_v16 = _rhs88
				_lhs77.F11 = _rhs89
				var _486_v23 _dafny.MultiSet
				_ = _486_v23
				_486_v23 = _dafny.MultiSetOf((_this).F29(), (_this).F29(), _480_v18.F37, _480_v18.F37, _480_v18.F37)
				var _487_v24 _dafny.Map
				_ = _487_v24
				_487_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _486_v23)
				var _488_v25 _dafny.Map
				_ = _488_v25
				_488_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_this.F30()).Cardinality()))
				var _489_v26 _dafny.Map
				_ = _489_v26
				_489_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _480_v18.F37)
				var _490_v27 _dafny.Array
				_ = _490_v27
				var _nwElement0_12 _dafny.Int = (_dafny.Zero).Minus((_this).F38())
				_ = _nwElement0_12
				var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(24))
				_ = _nw81
				(_nw81).ArraySet1(_nwElement0_12, 0)
				(_nw81).ArraySet1((_this).F38(), 1)
				(_nw81).ArraySet1(_dafny.IntOfInt64(106), 2)
				(_nw81).ArraySet1(_dafny.IntOfInt64(-464), 3)
				(_nw81).ArraySet1(_dafny.IntOfInt64(977), 4)
				(_nw81).ArraySet1((_this).F38(), 5)
				(_nw81).ArraySet1((_dafny.Zero).Minus((_this).F38()), 6)
				(_nw81).ArraySet1((_this).Fm16((func() _dafny.MultiSet {
					if (_487_v24).Contains(_480_v18.F37) {
						return (_487_v24).Get(_480_v18.F37).(_dafny.MultiSet)
					}
					return _486_v23
				})(), globalState), 7)
				(_nw81).ArraySet1((_this).F38(), 8)
				(_nw81).ArraySet1((func() _dafny.Int {
					if !(!((_this).F29())) {
						return (_this).F38()
					}
					return (_dafny.Zero).Minus((_this).F38())
				})(), 9)
				(_nw81).ArraySet1((_this).F38(), 10)
				(_nw81).ArraySet1((_480_v18).Fm12(globalState), 11)
				(_nw81).ArraySet1((_this).F38(), 12)
				(_nw81).ArraySet1((_this).F38(), 13)
				(_nw81).ArraySet1((_488_v25).Cardinality(), 14)
				(_nw81).ArraySet1(_dafny.IntOfInt64(754), 15)
				(_nw81).ArraySet1((_this).F38(), 16)
				(_nw81).ArraySet1((func() _dafny.Int {
					if (_this).F29() {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_489_v26, _480_v18.F37)).Cardinality()
					}
					return _dafny.IntOfInt64(-789)
				})(), 17)
				(_nw81).ArraySet1(_dafny.IntOfInt64(788), 18)
				(_nw81).ArraySet1((_this).F38(), 19)
				(_nw81).ArraySet1(((_482_v20).Cardinality()).Plus((_this).F38()), 20)
				(_nw81).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_484_v22).Cardinality())), 21)
				(_nw81).ArraySet1((_this).Fm16(_486_v23, globalState), 22)
				(_nw81).ArraySet1((_this).F38(), 23)
				_490_v27 = _nw81
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_490_v27), 0))
				_ = _index70
				(_490_v27).ArraySet1((_dafny.Zero).Minus((_this).F38()), (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_490_v27), 0))
				_ = _index71
				(_490_v27).ArraySet1((_this).F38(), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_490_v27), 0))
				_ = _index72
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_490_v27), 0))
				_ = _index73
				var _rhs90 _dafny.Int = _dafny.IntOfUint32((_this.F30()).Cardinality())
				_ = _rhs90
				var _rhs91 _dafny.Int = (_480_v18).Fm12(globalState)
				_ = _rhs91
				var _rhs92 _dafny.Int = (_this).F38()
				_ = _rhs92
				var _rhs93 _dafny.CodePoint = _479_v17
				_ = _rhs93
				var _lhs78 _dafny.Array = _490_v27
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_490_v27), 0))
				_ = _lhs79
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 _dafny.Array = _490_v27
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_490_v27), 0))
				_ = _lhs82
				(_lhs78).ArraySet1(_rhs90, (_lhs79).Int())
				_lhs80.F2 = _rhs91
				(_lhs81).ArraySet1(_rhs92, (_lhs82).Int())
				_479_v17 = _rhs93
				var _491_v28 _dafny.Array
				_ = _491_v28
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw82
				_491_v28 = _nw82
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_491_v28), 0))
				_ = _index74
				(_491_v28).ArraySet1((_this).F29(), (_index74).Int())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_491_v28), 0))
				_ = _index75
				(_491_v28).ArraySet1(_480_v18.F37, (_index75).Int())
			}
			(globalState).F15 = p0
			var _492_v29 _dafny.Map
			_ = _492_v29
			_492_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F29(), p0, (_this).F29(), true)).Cardinality()), !(!(true)))
			(globalState).F21 = ((_492_v29).Merge(_492_v29)).Cardinality()
			if true {
				var _493_v30 _dafny.Sequence
				_ = _493_v30
				_493_v30 = _dafny.SeqOf(p0)
				var _494_v31 _dafny.Sequence
				_ = _494_v31
				_494_v31 = _dafny.SeqOf(_dafny.SeqOf((_this).F29()), _dafny.Companion_Sequence_.Update(_493_v30, (Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_493_v30).Cardinality()))).Uint32(), !(p0)))
				var _495_v32 _dafny.Sequence
				_ = _495_v32
				_495_v32 = _dafny.SeqOf((_494_v31).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality())), _dafny.IntOfUint32((_494_v31).Cardinality()))).Uint32()).(_dafny.Sequence), _493_v30)
				var _496_v33 _dafny.CodePoint
				_ = _496_v33
				_496_v33 = _dafny.CodePoint('t')
				var _497_v34 _dafny.Array
				_ = _497_v34
				var _nwElement0_13 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_494_v31).Cardinality()), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _496_v33)
				_ = _nwElement0_13
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(3))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_13, 0)
				(_nw83).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm20(true, (_this).F29(), _493_v30, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), globalState), (Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((Companion_Default___.Fm20(true, (_this).F29(), _493_v30, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), globalState)).Cardinality()))).Uint32(), _496_v33), 1)
				(_nw83).ArraySet1(_this.F30(), 2)
				_497_v34 = _nw83
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_497_v34), 0))
				_ = _index76
				(_497_v34).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sar"), (_index76).Int())
				var _498_v35 D1
				_ = _498_v35
				_498_v35 = Companion_D1_.Create_DC3_(_493_v30)
				var _pat_let_tv10 = _493_v30
				_ = _pat_let_tv10
				var _499_v36 _dafny.Map
				_ = _499_v36
				_499_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), func(_pat_let6_0 D1) D1 {
					return func(_500_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let7_0 _dafny.Sequence) D1 {
							return func(_501_dt__update_hcf7_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_501_dt__update_hcf7_h0)
							}(_pat_let7_0)
						}(_pat_let_tv10)
					}(_pat_let6_0)
				}(_498_v35))
				var _502_v37 _dafny.Map
				_ = _502_v37
				_502_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F30())
				var _503_v38 _dafny.Map
				_ = _503_v38
				_503_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F38())
				var _504_v39 _dafny.Sequence
				_ = _504_v39
				_504_v39 = _dafny.SeqOf((_this).F38(), _dafny.IntOfInt64(567), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_503_v38).Cardinality(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _496_v33)).Cardinality()))
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_497_v34), 0))
				_ = _index77
				var _rhs94 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_502_v37).Contains((_this).F29()) {
						return (_502_v37).Get((_this).F29()).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _496_v33)
				})(), Companion_Default___.Fm20(!(p0), true, _dafny.SeqOf((_this).F29(), true, !((_493_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.IntOfUint32((_493_v30).Cardinality()))).Uint32()).(bool)), p0), (_dafny.Zero).Minus((_dafny.MultiSetOf((_this).F38(), (_this).F38())).Cardinality()), globalState))
				_ = _rhs94
				var _rhs95 _dafny.Int = _dafny.IntOfUint32((_504_v39).Cardinality())
				_ = _rhs95
				var _rhs96 _dafny.Map = ((Companion_Default___.Fm21(p0, (_this).F29(), p0, globalState)).Dtor_cf26()).Merge((_499_v36).Merge(_499_v36))
				_ = _rhs96
				var _lhs83 _dafny.Array = _497_v34
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_497_v34), 0))
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				(_lhs83).ArraySet1(_rhs94, (_lhs84).Int())
				_lhs85.F19 = _rhs95
				_499_v36 = _rhs96
				(globalState).F18 = false
				var _505_v40 _dafny.Array
				_ = _505_v40
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw84
				_505_v40 = _nw84
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_505_v40), 0))
				_ = _index78
				(_505_v40).ArraySet1((_this).F29(), (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_505_v40), 0))
				_ = _index79
				(_505_v40).ArraySet1(((_this).F38()).Cmp((_this).F38()) < 0, (_index79).Int())
				var _506_v41 _dafny.Array
				_ = _506_v41
				var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw85
				_506_v41 = _nw85
				var _507_v42 _dafny.Sequence
				_ = _507_v42
				_507_v42 = _dafny.SeqOf(_506_v41)
				_506_v41 = (_507_v42).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F38(), (_this).F38()), _dafny.IntOfUint32((_507_v42).Cardinality()))).Uint32()).(_dafny.Array)
				(globalState).F28 = Companion_Default___.Fm19(_496_v33, globalState)
			} else {
				var _508_v43 _dafny.CodePoint
				_ = _508_v43
				_508_v43 = _dafny.CodePoint('s')
				var _509_v44 _dafny.MultiSet
				_ = _509_v44
				_509_v44 = _dafny.MultiSetOf(_508_v43)
				var _510_v45 D2
				_ = _510_v45
				_510_v45 = Companion_D2_.Create_DC6_(_509_v44)
				_510_v45 = _510_v45
				var _511_v46 _dafny.Map
				_ = _511_v46
				_511_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_492_v29).Cardinality(), (_this).F38())
				_511_v46 = (_511_v46).Update(_dafny.IntOfInt64(-643), (_this).F38())
				var _512_v47 _dafny.Array
				_ = _512_v47
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
				_ = _nw86
				_512_v47 = _nw86
				var _513_v48 _dafny.Sequence
				_ = _513_v48
				_513_v48 = _dafny.SeqOf(true)
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_512_v47), 0))
				_ = _index80
				(_512_v47).ArraySet1(_dafny.MultiSetFromSeq(_513_v48), (_index80).Int())
				var _514_v49 _dafny.Sequence
				_ = _514_v49
				_514_v49 = _dafny.SeqOf(_this.F30())
				var _515_v50 _dafny.Map
				_ = _515_v50
				_515_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29())
				var _516_v51 _dafny.MultiSet
				_ = _516_v51
				_516_v51 = _dafny.MultiSetOf((func() bool {
					if (_492_v29).Contains((_this).F38()) {
						return (_492_v29).Get((_this).F38()).(bool)
					}
					return p0
				})(), (_this).F29())
				var _517_v52 _dafny.Map
				_ = _517_v52
				_517_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(301), _dafny.MultiSetOf(p0, (_this).F29()))
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_512_v47), 0))
				_ = _index81
				var _rhs97 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_514_v49).Cardinality()), ((_515_v50).Cardinality()).Minus((_this).F38()))
				_ = _rhs97
				var _rhs98 _dafny.MultiSet = (_516_v51).Difference((func() _dafny.MultiSet {
					if (_517_v52).Contains((_this).F38()) {
						return (_517_v52).Get((_this).F38()).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_513_v48)
				})())
				_ = _rhs98
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _512_v47
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_512_v47), 0))
				_ = _lhs88
				_lhs86.F1 = _rhs97
				(_lhs87).ArraySet1(_rhs98, (_lhs88).Int())
				var _518_v53 _dafny.Sequence
				_ = _518_v53
				_518_v53 = _dafny.SeqOf(_513_v48)
				(globalState).F11 = Companion_Default___.Fm20((_this).F29(), p0, (_518_v53).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_518_v53).Cardinality()))).Uint32()).(_dafny.Sequence), (_this).F38(), globalState)
				var _519_v54 _dafny.Set
				_ = _519_v54
				_519_v54 = _dafny.SetOf(p0)
				var _520_v55 _dafny.Sequence
				_ = _520_v55
				_520_v55 = _dafny.SeqOf((Companion_D6_.Create_DC21_(_519_v54)).Dtor_cf32(), Companion_Default___.Fm22((_this).F38(), !(p0), _508_v43, globalState), _519_v54)
				(globalState).F19 = ((_520_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.IntOfUint32((_520_v55).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
			}
			var _521_v56 _dafny.Map
			_ = _521_v56
			_521_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) == ((_this).F29()), (_this).F38())
			_521_v56 = (_521_v56).Update((_this).F29(), Companion_Default___.SafeDivisionInt((_this).F38(), (_this).F38()))
		}
		var _522_v57 _dafny.Array
		_ = _522_v57
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_10
		var _nw87 _dafny.Array
		_ = _nw87
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw87 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) bool = func(_523_i4 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.Contains(_this.F30(), _dafny.CodePoint('p'))
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw87 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw87).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw87).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_522_v57 = _nw87
		var _524_v58 _dafny.CodePoint
		_ = _524_v58
		_524_v58 = _dafny.CodePoint('c')
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_522_v57), 0))
		_ = _index82
		(_522_v57).ArraySet1(!(Companion_Default___.Fm19(_524_v58, globalState)), (_index82).Int())
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_522_v57), 0))
		_ = _index83
		(_522_v57).ArraySet1(p0, (_index83).Int())
		(globalState).F19 = (_dafny.Zero).Minus((_dafny.IntOfInt64(-489)).Times((_this).F38()))
		var _pat_let_tv11 = p0
		_ = _pat_let_tv11
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _pat_let_tv13 = _522_v57
		_ = _pat_let_tv13
		var _pat_let_tv14 = _522_v57
		_ = _pat_let_tv14
		(globalState).F28 = func(_source8 D2) bool {
			if _source8.Is_DC7() {
				var _525___mcc_h0 bool = _source8.Get_().(D2_DC7).Cf11
				_ = _525___mcc_h0
				var _526_cf11 bool = _525___mcc_h0
				_ = _526_cf11
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F30(), _this.F30())
			} else if _source8.Is_DC8() {
				var _527___mcc_h1 _dafny.CodePoint = _source8.Get_().(D2_DC8).Cf12
				_ = _527___mcc_h1
				var _528___mcc_h2 bool = _source8.Get_().(D2_DC8).Cf13
				_ = _528___mcc_h2
				var _529_cf13 bool = _528___mcc_h2
				_ = _529_cf13
				var _530_cf12 _dafny.CodePoint = _527___mcc_h1
				_ = _530_cf12
				return (_pat_let_tv11) && (_529_cf13)
			} else if _source8.Is_DC9() {
				var _531___mcc_h3 _dafny.Sequence = _source8.Get_().(D2_DC9).Cf14
				_ = _531___mcc_h3
				var _532___mcc_h4 _dafny.Int = _source8.Get_().(D2_DC9).Cf15
				_ = _532___mcc_h4
				var _533___mcc_h5 bool = _source8.Get_().(D2_DC9).Cf16
				_ = _533___mcc_h5
				var _534___mcc_h6 _dafny.Int = _source8.Get_().(D2_DC9).Cf17
				_ = _534___mcc_h6
				var _535_cf17 _dafny.Int = _534___mcc_h6
				_ = _535_cf17
				var _536_cf16 bool = _533___mcc_h5
				_ = _536_cf16
				var _537_cf15 _dafny.Int = _532___mcc_h4
				_ = _537_cf15
				var _538_cf14 _dafny.Sequence = _531___mcc_h3
				_ = _538_cf14
				return (false) == (_pat_let_tv12)
			} else {
				var _539___mcc_h7 _dafny.MultiSet = _source8.Get_().(D2_DC6).Cf10
				_ = _539___mcc_h7
				var _540_cf10 _dafny.MultiSet = _539___mcc_h7
				_ = _540_cf10
				return (_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(bool)
			}
		}(Companion_D2_.Create_DC9_(_this.F30(), (_this).F38(), p0, (_this).F38()))
		var _541_v59 _dafny.Map
		_ = _541_v59
		_541_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p0) || (p0))
		_541_v59 = (_541_v59).Update(!((_this).F29()), p0)
		r0 = Companion_Default___.Fm23(globalState)
		r1 = p0
		return r0, r1
	}
}
func (_this *C2) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _542_v0 _dafny.Array
		_ = _542_v0
		var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
		_ = _nw88
		_542_v0 = _nw88
		for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_542_v0), 0))); ; {
			_guard_loop_2, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			var _543_i0 _dafny.Int
			_543_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_543_i0).Sign() != -1) && ((_543_i0).Cmp(_dafny.ArrayLen((_542_v0), 0)) < 0)) {
				(_542_v0).ArraySet1CodePoint(_dafny.CodePoint('j'), (_543_i0).Int())
			}
		}
		var _544_v1 _dafny.Set
		_ = _544_v1
		_544_v1 = _dafny.SetOf((_this).F29())
		var _545_v3 _dafny.Map
		_ = _545_v3
		_545_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(50))
		(globalState).F19 = (func() _dafny.Int {
			if (_544_v1).IsSubsetOf(_544_v1) {
				return (_this).F38()
			}
			return ((func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(588), _dafny.IntOfInt64(387))); ; {
					_compr_26, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _546_v2 _dafny.Int
					_546_v2 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(588)).Cmp(_546_v2) <= 0) && ((_546_v2).Cmp(_dafny.IntOfInt64(387)) < 0) {
						_coll26.Add(Companion_Default___.SafeModuloInt(_546_v2, p2), _dafny.IntOfUint32((_dafny.SeqOf(p2, (_this).F38(), (_this).F38())).Cardinality()))
					}
				}
				return _coll26.ToMap()
			}()).Cardinality()).Minus((_545_v3).Cardinality())
		})()
		var _547_v4 _dafny.Array
		_ = _547_v4
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_11
		var _nw89 _dafny.Array
		_ = _nw89
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw89 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = func(_548_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_548_i1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality())))
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw89 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw89).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw89).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_547_v4 = _nw89
		var _549_v5 _dafny.MultiSet
		_ = _549_v5
		_549_v5 = _dafny.MultiSetOf(_547_v4)
		var _rhs99 _dafny.Int = _dafny.IntOfInt64(727)
		_ = _rhs99
		var _rhs100 _dafny.MultiSet = _549_v5
		_ = _rhs100
		var _lhs89 *GlobalState = globalState
		_ = _lhs89
		_lhs89.F19 = _rhs99
		_549_v5 = _rhs100
		if Companion_Default___.Fm19(_dafny.CodePoint('k'), globalState) {
			var _550_v6 *C0
			_ = _550_v6
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__((_this).F29())
			_550_v6 = _nw90
			(globalState).F18 = (Companion_Default___.Fm24(globalState)).Dtor_cf8()
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))
			_ = _index84
			(_547_v4).ArraySet1(_dafny.IntOfInt64(-588), (_index84).Int())
			var _551_v7 _dafny.Sequence
			_ = _551_v7
			_551_v7 = _dafny.SeqOf((_this).F29())
			var _552_v8 _dafny.Map
			_ = _552_v8
			_552_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_551_v7, (_this).F38())
			var _553_v9 D4
			_ = _553_v9
			_553_v9 = Companion_D4_.Create_DC13_(_552_v8)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))
			_ = _index85
			var _rhs101 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfInt64(855))
			_ = _rhs101
			var _rhs102 D4 = _553_v9
			_ = _rhs102
			var _rhs103 bool = true
			_ = _rhs103
			var _lhs90 _dafny.Array = _547_v4
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))
			_ = _lhs91
			var _lhs92 *GlobalState = globalState
			_ = _lhs92
			(_lhs90).ArraySet1(_rhs101, (_lhs91).Int())
			_553_v9 = _rhs102
			_lhs92.F18 = _rhs103
			var _554_v10 D2
			_ = _554_v10
			_554_v10 = Companion_D2_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("xe"), (_547_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))).Int()).(_dafny.Int), p0, _dafny.IntOfUint32((_this.F30()).Cardinality()))
			var _555_v11 _dafny.MultiSet
			_ = _555_v11
			_555_v11 = _dafny.MultiSetOf(_dafny.SetOf((_this).F29(), (_this).F29()))
			var _556_v12 _dafny.MultiSet
			_ = _556_v12
			_556_v12 = _dafny.MultiSetOf((_this).F29(), _dafny.Companion_Sequence_.Contains(_this.F30(), _dafny.CodePoint('t')), ((_554_v10).Dtor_cf15()).Cmp((func() _dafny.Int {
				if (_555_v11).Contains(_544_v1) {
					return (_555_v11).Multiplicity(_544_v1)
				}
				return p2
			})()) >= 0)
			_556_v12 = _556_v12
			var _557_v13 _dafny.Array
			_ = _557_v13
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw91
			_557_v13 = _nw91
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_557_v13), 0))
			_ = _index86
			(_557_v13).ArraySet1((_this).F29(), (_index86).Int())
			var _558_v14 _dafny.Array
			_ = _558_v14
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
			_ = _nw92
			_558_v14 = _nw92
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_558_v14), 0))
			_ = _index87
			(_558_v14).ArraySet1((_544_v1).Union(_544_v1), (_index87).Int())
			var _559_v15 _dafny.CodePoint
			_ = _559_v15
			_559_v15 = _dafny.CodePoint('h')
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))
			_ = _index88
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_557_v13), 0))
			_ = _index89
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_558_v14), 0))
			_ = _index90
			var _rhs104 _dafny.Int = _dafny.IntOfInt64(211)
			_ = _rhs104
			var _rhs105 bool = (p2).Cmp((_this).F38()) >= 0
			_ = _rhs105
			var _rhs106 _dafny.Int = (_547_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))).Int()).(_dafny.Int)
			_ = _rhs106
			var _rhs107 bool = Companion_Default___.Fm19(_559_v15, globalState)
			_ = _rhs107
			var _rhs108 _dafny.Set = _544_v1
			_ = _rhs108
			var _lhs93 *GlobalState = globalState
			_ = _lhs93
			var _lhs94 *GlobalState = globalState
			_ = _lhs94
			var _lhs95 _dafny.Array = _547_v4
			_ = _lhs95
			var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_547_v4), 0))
			_ = _lhs96
			var _lhs97 _dafny.Array = _557_v13
			_ = _lhs97
			var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_557_v13), 0))
			_ = _lhs98
			var _lhs99 _dafny.Array = _558_v14
			_ = _lhs99
			var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_558_v14), 0))
			_ = _lhs100
			_lhs93.F14 = _rhs104
			_lhs94.F18 = _rhs105
			(_lhs95).ArraySet1(_rhs106, (_lhs96).Int())
			(_lhs97).ArraySet1(_rhs107, (_lhs98).Int())
			(_lhs99).ArraySet1(_rhs108, (_lhs100).Int())
		} else {
			var _560_v16 _dafny.Sequence
			_ = _560_v16
			_560_v16 = _dafny.SeqOf(_dafny.IntOfUint32((_this.F30()).Cardinality()))
			var _561_v17 _dafny.Set
			_ = _561_v17
			_561_v17 = _dafny.SetOf(_560_v16, (func() _dafny.Sequence {
				if (_this).F29() {
					return _560_v16
				}
				return _dafny.SeqOf(_dafny.IntOfInt64(-539), p2)
			})(), _560_v16, _560_v16)
			_561_v17 = ((_561_v17).Difference(_561_v17)).Intersection(Companion_Default___.Fm25((_this).F29(), p0, _dafny.UnicodeSeqOfUtf8Bytes("lyyloxwhs"), globalState))
			(globalState).F1 = (_this).F38()
			var _562_v18 _dafny.Array
			_ = _562_v18
			var _nwElement0_14 bool = true
			_ = _nwElement0_14
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(17))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_14, 0)
			(_nw93).ArraySet1((_this).F29(), 1)
			(_nw93).ArraySet1(p0, 2)
			(_nw93).ArraySet1((_this).F29(), 3)
			(_nw93).ArraySet1(p0, 4)
			(_nw93).ArraySet1(p0, 5)
			(_nw93).ArraySet1((_this).F29(), 6)
			(_nw93).ArraySet1((_this).F29(), 7)
			(_nw93).ArraySet1((func() bool {
				if (_this).F29() {
					return true
				}
				return (_this).F29()
			})(), 8)
			(_nw93).ArraySet1((p2).Cmp(p2) == 0, 9)
			(_nw93).ArraySet1((_this).F29(), 10)
			(_nw93).ArraySet1(p0, 11)
			(_nw93).ArraySet1((_this).F29(), 12)
			(_nw93).ArraySet1((_this).F29(), 13)
			(_nw93).ArraySet1(p0, 14)
			(_nw93).ArraySet1(p0, 15)
			(_nw93).ArraySet1(p0, 16)
			_562_v18 = _nw93
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_562_v18), 0))
			_ = _index91
			(_562_v18).ArraySet1((_this).F29(), (_index91).Int())
			var _563_v19 _dafny.CodePoint
			_ = _563_v19
			_563_v19 = _dafny.CodePoint('x')
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_562_v18), 0))
			_ = _index92
			var _rhs109 _dafny.Int = (_dafny.Zero).Minus((_560_v16).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_560_v16).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs109
			var _rhs110 _dafny.Int = (_this).F38()
			_ = _rhs110
			var _rhs111 bool = ((func() _dafny.Int {
				if (_this).F29() {
					return p2
				}
				return (_this).F38()
			})()).Cmp(((_dafny.Zero).Minus(p2)).Times((_this).F38())) >= 0
			_ = _rhs111
			var _rhs112 bool = p0
			_ = _rhs112
			var _rhs113 _dafny.CodePoint = _563_v19
			_ = _rhs113
			var _lhs101 *GlobalState = globalState
			_ = _lhs101
			var _lhs102 *GlobalState = globalState
			_ = _lhs102
			var _lhs103 _dafny.Array = _562_v18
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_562_v18), 0))
			_ = _lhs104
			_lhs101.F14 = _rhs109
			_lhs102.F14 = _rhs110
			r1 = _rhs111
			(_lhs103).ArraySet1(_rhs112, (_lhs104).Int())
			_563_v19 = _rhs113
			var _564_v20 D3
			_ = _564_v20
			_564_v20 = Companion_D3_.Create_DC11_((_this).F29(), (_this).F29())
			var _565_v21 _dafny.Map
			_ = _565_v21
			_565_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v20, _563_v19)
			(globalState).F19 = (Companion_Default___.SafeModuloInt(p2, (_dafny.Zero).Minus((_this).F38()))).Plus((_565_v21).Cardinality())
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_562_v18), 0))
			_ = _index93
			(_562_v18).ArraySet1((_this).F29(), (_index93).Int())
		}
		var _566_v22 _dafny.Array
		_ = _566_v22
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_12
		var _nw94 _dafny.Array
		_ = _nw94
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw94 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) bool = (func(_567_p2 _dafny.Int) func(_dafny.Int) bool {
				return func(_568_i3 _dafny.Int) bool {
					return !((_567_p2).Cmp(_567_p2) < 0)
				}
			})(p2)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw94 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw94).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw94).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_566_v22 = _nw94
		for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_566_v22), 0))); ; {
			_guard_loop_3, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _569_i2 _dafny.Int
			_569_i2 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_569_i2).Sign() != -1) && ((_569_i2).Cmp(_dafny.ArrayLen((_566_v22), 0)) < 0)) {
				(_566_v22).ArraySet1((((_dafny.MultiSetOf((Companion_D2_.Create_DC9_(_this.F30(), (_this).F38(), p0, (_this).F38())).Dtor_cf16(), (_this).F29())).Cardinality()).Plus((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(p0, p0))).Cardinality()))).Cmp(_dafny.IntOfInt64(474)) < 0, (_569_i2).Int())
			}
		}
		if p0 {
			(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pwxebygro"), _this.F30())
			var _570_v23 _dafny.Set
			_ = _570_v23
			_570_v23 = _dafny.SetOf((_this).F38())
			var _571_v24 _dafny.MultiSet
			_ = _571_v24
			_571_v24 = _dafny.MultiSetOf((_this).F38())
			var _rhs114 _dafny.Set = _dafny.SetOf((_this).F38())
			_ = _rhs114
			var _rhs115 bool = (_this).F29()
			_ = _rhs115
			var _rhs116 _dafny.Set = (_570_v23).Union((_570_v23).Union(func() _dafny.Set {
				var _coll27 = _dafny.NewBuilder()
				_ = _coll27
				for _iter31 := _dafny.Iterate((_571_v24).Elements()); ; {
					_compr_27, _ok31 := _iter31()
					if !_ok31 {
						break
					}
					var _572_v25 _dafny.Int
					_572_v25 = interface{}(_compr_27).(_dafny.Int)
					if (_571_v24).Contains(_572_v25) {
						_coll27.Add((_572_v25).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tsl")).Cardinality())))
					}
				}
				return _coll27.ToSet()
			}()))
			_ = _rhs116
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			_570_v23 = _rhs114
			_lhs105.F28 = _rhs115
			_570_v23 = _rhs116
			var _573_v26 *C0
			_ = _573_v26
			var _nw95 *C0 = New_C0_()
			_ = _nw95
			_nw95.Ctor__(!((_this).F29()))
			_573_v26 = _nw95
			var _574_v27 *C0
			_ = _574_v27
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(_573_v26.F37)
			_574_v27 = _nw96
			var _575_v28 _dafny.Set
			_ = _575_v28
			var _576_v29 _dafny.Int
			_ = _576_v29
			var _out10 _dafny.Set
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out10, _out11 = (_this).M8(_566_v22, globalState)
			_575_v28 = _out10
			_576_v29 = _out11
		} else {
			(globalState).F19 = p2
			var _577_v30 _dafny.Sequence
			_ = _577_v30
			_577_v30 = _dafny.SeqOf((_this).F29())
			(globalState).F14 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F38(), _dafny.IntOfUint32((_577_v30).Cardinality())))
			(globalState).F13 = (_this).F29()
			var _578_v31 _dafny.Sequence
			_ = _578_v31
			_578_v31 = _dafny.SeqOf(_566_v22, _566_v22, _566_v22)
			var _579_v32 _dafny.Set
			_ = _579_v32
			_579_v32 = _dafny.SetOf(_566_v22, (_578_v31).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_578_v31).Cardinality()))).Uint32()).(_dafny.Array), _566_v22, _566_v22)
			(globalState).F13 = ((_dafny.SetOf(_566_v22)).IsDisjointFrom(_579_v32)) == (p0)
			(globalState).F14 = p2
		}
		r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("trbmjfc"), _this.F30())
		r1 = (_this).F29()
		return r0, r1
	}
}
func (_this *C2) M8(p0 _dafny.Array, globalState *GlobalState) (_dafny.Set, _dafny.Int) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _580_v0 _dafny.Sequence
		_ = _580_v0
		_580_v0 = _dafny.SeqOf((_this).F29())
		var _581_v1 _dafny.Sequence
		_ = _581_v1
		_581_v1 = _dafny.SeqOf(_dafny.IntOfUint32((_580_v0).Cardinality()), (_this).F38())
		var _582_v2 _dafny.Map
		_ = _582_v2
		_582_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), ((_this).Fm16(_dafny.MultiSetOf((_this).F29()), globalState)).Minus((_581_v1).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_581_v1).Cardinality()))).Uint32()).(_dafny.Int)))
		_582_v2 = (_582_v2).Update((_this).F38(), (_this).F38())
		var _583_v3 _dafny.Array
		_ = _583_v3
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_13
		var _nw97 _dafny.Array
		_ = _nw97
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw97 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) bool = func(_584_i1 _dafny.Int) bool {
				return (_this).F29()
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw97 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw97).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw97).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_583_v3 = _nw97
		for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_583_v3), 0))); ; {
			_guard_loop_4, _ok32 := _iter32()
			if !_ok32 {
				break
			}
			var _585_i0 _dafny.Int
			_585_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_585_i0).Sign() != -1) && ((_585_i0).Cmp(_dafny.ArrayLen((_583_v3), 0)) < 0)) {
				(_583_v3).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_586_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), _this.F30()), (_585_i0).Int())
			}
		}
		(globalState).F13 = (_this).F29()
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_583_v3), 0))); ; {
			_guard_loop_5, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _587_i3 _dafny.Int
			_587_i3 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_587_i3).Sign() != -1) && ((_587_i3).Cmp(_dafny.ArrayLen((_583_v3), 0)) < 0)) {
				(_583_v3).ArraySet1(((_this).F38()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fpqtgpu")).Cardinality())) >= 0, (_587_i3).Int())
			}
		}
		var _588_v4 _dafny.MultiSet
		_ = _588_v4
		_588_v4 = _dafny.MultiSetOf(false)
		var _hi1 _dafny.Int = (_this).F38()
		_ = _hi1
		for _589_i4 := Companion_Default___.SafeModuloInt((_this).Fm16(_588_v4, globalState), (_this).F38()); _589_i4.Cmp(_hi1) < 0; _589_i4 = _589_i4.Plus(_dafny.One) {
			var _590_v5 _dafny.Set
			_ = _590_v5
			_590_v5 = _dafny.SetOf(_588_v4, (_588_v4).Update((_this).F29(), Companion_Default___.Abs((_this).F38())))
			(globalState).F7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_581_v1, _581_v1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F38()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_581_v1, _581_v1)).Cardinality()))).Uint32(), (_590_v5).Cardinality()), _581_v1)
			var _591_v6 _dafny.MultiSet
			_ = _591_v6
			_591_v6 = _dafny.MultiSetOf((_this).F38())
			var _592_v7 D3
			_ = _592_v7
			_592_v7 = Companion_D3_.Create_DC10_(_591_v6)
			var _593_v8 D3
			_ = _593_v8
			_593_v8 = Companion_D3_.Create_DC12_(_592_v7)
			var _source9 D3 = _593_v8
			_ = _source9
			if _source9.Is_DC11() {
				var _594___mcc_h0 bool = _source9.Get_().(D3_DC11).Cf19
				_ = _594___mcc_h0
				var _595___mcc_h1 bool = _source9.Get_().(D3_DC11).Cf20
				_ = _595___mcc_h1
				var _596_cf20 bool = _595___mcc_h1
				_ = _596_cf20
				var _597_cf19 bool = _594___mcc_h0
				_ = _597_cf19
				var _598_v9 _dafny.Map
				_ = _598_v9
				_598_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), _dafny.SetOf((_this).F38()))
				var _599_v10 _dafny.Set
				_ = _599_v10
				_599_v10 = _dafny.SetOf(p0)
				_598_v9 = (_598_v9).Update((_dafny.Zero).Minus(_589_i4), _dafny.SetOf(_589_i4, _dafny.IntOfUint32((_this.F30()).Cardinality()), (_this).F38(), (_599_v10).Cardinality()))
				var _600_v11 *C0
				_ = _600_v11
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__((_this).F29())
				_600_v11 = _nw98
				_582_v2 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-507), (_this).F38())).Merge(_582_v2)
				var _601_v12 _dafny.Map
				_ = _601_v12
				_601_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_600_v11.F37, _589_i4)
				var _602_v13 _dafny.Array
				_ = _602_v13
				var _nwElement0_15 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_580_v0, (Companion_Default___.SafeIndex(_589_i4, _dafny.IntOfUint32((_580_v0).Cardinality()))).Uint32(), (_this).F29())).Cardinality())
				_ = _nwElement0_15
				var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(26))
				_ = _nw99
				(_nw99).ArraySet1(_nwElement0_15, 0)
				(_nw99).ArraySet1((_this).F38(), 1)
				(_nw99).ArraySet1((_dafny.MultiSetFromSeq(_580_v0)).Cardinality(), 2)
				(_nw99).ArraySet1((_this).F38(), 3)
				(_nw99).ArraySet1((_600_v11).Fm12(globalState), 4)
				(_nw99).ArraySet1((_this).F38(), 5)
				(_nw99).ArraySet1((_this).F38(), 6)
				(_nw99).ArraySet1((_this).F38(), 7)
				(_nw99).ArraySet1(_dafny.IntOfUint32((_this.F30()).Cardinality()), 8)
				(_nw99).ArraySet1((_this).F38(), 9)
				(_nw99).ArraySet1(_589_i4, 10)
				(_nw99).ArraySet1((_this).F38(), 11)
				(_nw99).ArraySet1(_dafny.IntOfInt64(670), 12)
				(_nw99).ArraySet1((_this).F38(), 13)
				(_nw99).ArraySet1((_dafny.Zero).Minus((_this).F38()), 14)
				(_nw99).ArraySet1((func() _dafny.Int {
					if (_601_v12).Contains(false) {
						return (_601_v12).Get(false).(_dafny.Int)
					}
					return (_this).F38()
				})(), 15)
				(_nw99).ArraySet1((_dafny.Zero).Minus((_this).F38()), 16)
				(_nw99).ArraySet1(_589_i4, 17)
				(_nw99).ArraySet1((_this).Fm16(_588_v4, globalState), 18)
				(_nw99).ArraySet1((_this).F38(), 19)
				(_nw99).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(698))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_603_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality()), 20)
				(_nw99).ArraySet1(_dafny.IntOfUint32((_this.F30()).Cardinality()), 21)
				(_nw99).ArraySet1((_this).F38(), 22)
				(_nw99).ArraySet1(_dafny.IntOfInt64(329), 23)
				(_nw99).ArraySet1((_this).F38(), 24)
				(_nw99).ArraySet1(_589_i4, 25)
				_602_v13 = _nw99
				var _604_v14 _dafny.Map
				_ = _604_v14
				_604_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v13, !(_600_v11.F37))
				(globalState).F15 = (func() bool {
					if (_604_v14).Contains(_602_v13) {
						return (_604_v14).Get(_602_v13).(bool)
					}
					return _596_cf20
				})()
			} else if _source9.Is_DC10() {
				var _605___mcc_h2 _dafny.MultiSet = _source9.Get_().(D3_DC10).Cf18
				_ = _605___mcc_h2
				var _606_cf18 _dafny.MultiSet = _605___mcc_h2
				_ = _606_cf18
				var _607_v15 _dafny.Array
				_ = _607_v15
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw100
				_607_v15 = _nw100
				var _608_v16 _dafny.Map
				_ = _608_v16
				_608_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm16(_588_v4, globalState), _607_v15)
				_608_v16 = (_608_v16).Update(_589_i4, _607_v15)
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_583_v3), 0))
				_ = _index94
				(_583_v3).ArraySet1((_this).F29(), (_index94).Int())
				var _609_v17 _dafny.CodePoint
				_ = _609_v17
				_609_v17 = _dafny.CodePoint('k')
				var _610_v18 _dafny.Map
				_ = _610_v18
				_610_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), (_this).F38())
				var _611_v19 _dafny.Map
				_ = _611_v19
				_611_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _610_v18)
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_583_v3), 0))
				_ = _index95
				(_583_v3).ArraySet1((func() bool {
					if Companion_Default___.Fm19(_609_v17, globalState) {
						return (_this).F29()
					}
					return !((func() _dafny.Map {
						if (_611_v19).Contains((_this).F29()) {
							return (_611_v19).Get((_this).F29()).(_dafny.Map)
						}
						return Companion_Default___.Fm26((_this).F29(), (_this).F38(), (_this).F29(), globalState)
					})()).Equals(_610_v18)
				})(), (_index95).Int())
				(globalState).F21 = _589_i4
				var _612_v20 D3
				_ = _612_v20
				_612_v20 = Companion_D3_.Create_DC10_((_591_v6).Update(_589_i4, Companion_Default___.Abs((_this).F38())))
				var _613_v22 _dafny.Map
				_ = _613_v22
				_613_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_589_i4), (_this).F29())
				var _614_v23 _dafny.Array
				_ = _614_v23
				var _nwElement0_16 _dafny.MultiSet = _dafny.MultiSetFromSeq(_581_v1)
				_ = _nwElement0_16
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(29))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_16, 0)
				(_nw101).ArraySet1((_612_v20).Dtor_cf18(), 1)
				(_nw101).ArraySet1(_606_cf18, 2)
				(_nw101).ArraySet1(_606_cf18, 3)
				(_nw101).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_581_v1, _581_v1)), 4)
				(_nw101).ArraySet1(Companion_Default___.Fm27((_this).F38(), false, globalState), 5)
				(_nw101).ArraySet1(_591_v6, 6)
				(_nw101).ArraySet1(_606_cf18, 7)
				(_nw101).ArraySet1(_dafny.MultiSetFromSeq(_581_v1), 8)
				(_nw101).ArraySet1(_606_cf18, 9)
				(_nw101).ArraySet1(_591_v6, 10)
				(_nw101).ArraySet1(_dafny.MultiSetOf(_589_i4, _589_i4), 11)
				(_nw101).ArraySet1((_606_cf18).Update((_this).F38(), Companion_Default___.Abs(_589_i4)), 12)
				(_nw101).ArraySet1((func() _dafny.MultiSet {
					if (_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool) {
						return _dafny.MultiSetOf((_this).F38(), _589_i4, _589_i4)
					}
					return _606_cf18
				})(), 13)
				(_nw101).ArraySet1((_606_cf18).Update(_589_i4, Companion_Default___.Abs(_589_i4)), 14)
				(_nw101).ArraySet1((_591_v6).Intersection(_591_v6), 15)
				(_nw101).ArraySet1(_606_cf18, 16)
				(_nw101).ArraySet1((_606_cf18).Difference(_591_v6), 17)
				(_nw101).ArraySet1(_606_cf18, 18)
				(_nw101).ArraySet1(Companion_Default___.Fm27((_this).F38(), (_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool), globalState), 19)
				(_nw101).ArraySet1(_dafny.MultiSetOf((func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter34 := _dafny.Iterate((_613_v22).Keys().Elements()); ; {
						_compr_28, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _615_v21 _dafny.Int
						_615_v21 = interface{}(_compr_28).(_dafny.Int)
						if (_613_v22).Contains(_615_v21) {
							_coll28.Add(Companion_Default___.SafeModuloInt(_615_v21, (_this).F38()), (_606_cf18).Cardinality())
						}
					}
					return _coll28.ToMap()
				}()).Cardinality()), 20)
				(_nw101).ArraySet1((_591_v6).Difference((_612_v20).Dtor_cf18()), 21)
				(_nw101).ArraySet1(_606_cf18, 22)
				(_nw101).ArraySet1(_606_cf18, 23)
				(_nw101).ArraySet1(_606_cf18, 24)
				(_nw101).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfInt64(79)), 25)
				(_nw101).ArraySet1((_591_v6).Intersection(_591_v6), 26)
				(_nw101).ArraySet1(_dafny.MultiSetOf(_589_i4), 27)
				(_nw101).ArraySet1(_dafny.MultiSetOf((_this).F38(), (_this).F38()), 28)
				_614_v23 = _nw101
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_614_v23), 0))
				_ = _index96
				(_614_v23).ArraySet1((_dafny.MultiSetOf(_589_i4)).Update(_589_i4, Companion_Default___.Abs((_this).F38())), (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_614_v23), 0))
				_ = _index97
				var _rhs117 _dafny.Int = _589_i4
				_ = _rhs117
				var _rhs118 _dafny.MultiSet = (_591_v6).Union(_606_cf18)
				_ = _rhs118
				var _rhs119 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _609_v17)
				_ = _rhs119
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				var _lhs107 _dafny.Array = _614_v23
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_614_v23), 0))
				_ = _lhs108
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				_lhs106.F21 = _rhs117
				(_lhs107).ArraySet1(_rhs118, (_lhs108).Int())
				_lhs109.F11 = _rhs119
			} else {
				var _616___mcc_h3 D3 = _source9.Get_().(D3_DC12).Cf21
				_ = _616___mcc_h3
				var _617_cf21 D3 = _616___mcc_h3
				_ = _617_cf21
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))
				_ = _index98
				(_583_v3).ArraySet1((_this).F29(), (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))
				_ = _index99
				(_583_v3).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F29(), (_this).F29()), _580_v0), (_588_v4).Contains(!((_this).F29()))), (_index99).Int())
				(globalState).F13 = ((_this).Fm16(_588_v4, globalState)).Cmp(_589_i4) == 0
				(globalState).F28 = (_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool)
				var _618_v24 _dafny.Array
				_ = _618_v24
				var _nw102 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw102
				_618_v24 = _nw102
				var _619_v25 _dafny.Array
				_ = _619_v25
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw103
				_619_v25 = _nw103
				var _620_v26 _dafny.Map
				_ = _620_v26
				_620_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool))
				var _621_v27 _dafny.CodePoint
				_ = _621_v27
				_621_v27 = _dafny.CodePoint('v')
				var _622_v28 _dafny.Map
				_ = _622_v28
				_622_v28 = _620_v26
				var _623_v29 _dafny.Array
				_ = _623_v29
				var _nwElement0_17 _dafny.Map = _620_v26
				_ = _nwElement0_17
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(27))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_17, 0)
				(_nw104).ArraySet1(Companion_Default___.Fm28((_this).F29(), _589_i4, (_this).F38(), (_this).F29(), globalState), 1)
				(_nw104).ArraySet1(_620_v26, 2)
				(_nw104).ArraySet1(_620_v26, 3)
				(_nw104).ArraySet1(_620_v26, 4)
				(_nw104).ArraySet1(_620_v26, 5)
				(_nw104).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F29()), (_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool)), 6)
				(_nw104).ArraySet1(_620_v26, 7)
				(_nw104).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), Companion_Default___.Fm19(_621_v27, globalState)), 8)
				(_nw104).ArraySet1(_620_v26, 9)
				(_nw104).ArraySet1(_620_v26, 10)
				(_nw104).ArraySet1(_620_v26, 11)
				(_nw104).ArraySet1(_620_v26, 12)
				(_nw104).ArraySet1(_620_v26, 13)
				(_nw104).ArraySet1(_620_v26, 14)
				(_nw104).ArraySet1(_620_v26, 15)
				(_nw104).ArraySet1(_620_v26, 16)
				(_nw104).ArraySet1((_622_v28), 17)
				(_nw104).ArraySet1(_620_v26, 18)
				(_nw104).ArraySet1(_620_v26, 19)
				(_nw104).ArraySet1(_620_v26, 20)
				(_nw104).ArraySet1(_620_v26, 21)
				(_nw104).ArraySet1(_620_v26, 22)
				(_nw104).ArraySet1(_620_v26, 23)
				(_nw104).ArraySet1(Companion_Default___.Fm28((_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool), _589_i4, (_this).F38(), (_this).F29(), globalState), 24)
				(_nw104).ArraySet1(Companion_Default___.Fm28((_583_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_583_v3), 0))).Int()).(bool), (_this).F38(), (_this).F38(), (_this).F29(), globalState), 25)
				(_nw104).ArraySet1(_620_v26, 26)
				_623_v29 = _nw104
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_619_v25), 0))
				_ = _index100
				(_619_v25).ArraySet1(_623_v29, (_index100).Int())
				var _624_v30 _dafny.Array
				_ = _624_v30
				var _nw105 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw105
				_624_v30 = _nw105
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_624_v30), 0))
				_ = _index101
				(_624_v30).ArraySet1((_dafny.Zero).Minus(_589_i4), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_619_v25), 0))
				_ = _index102
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_624_v30), 0))
				_ = _index103
				var _rhs120 _dafny.Array = _618_v24
				_ = _rhs120
				var _rhs121 _dafny.Array = _623_v29
				_ = _rhs121
				var _rhs122 _dafny.Int = (_589_i4).Times((_this).F38())
				_ = _rhs122
				var _lhs110 _dafny.Array = _619_v25
				_ = _lhs110
				var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_619_v25), 0))
				_ = _lhs111
				var _lhs112 _dafny.Array = _624_v30
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_624_v30), 0))
				_ = _lhs113
				_618_v24 = _rhs120
				(_lhs110).ArraySet1(_rhs121, (_lhs111).Int())
				(_lhs112).ArraySet1(_rhs122, (_lhs113).Int())
			}
			(globalState).F13 = !(true) || ((_this).F29())
			(globalState).F14 = (_this).F38()
		}
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_583_v3), 0))); ; {
			_guard_loop_6, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _625_i6 _dafny.Int
			_625_i6 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_625_i6).Sign() != -1) && ((_625_i6).Cmp(_dafny.ArrayLen((_583_v3), 0)) < 0)) {
				(_583_v3).ArraySet1(((_this).F38()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())).Cardinality())) <= 0, (_625_i6).Int())
			}
		}
		var _626_v34 _dafny.Map
		_ = _626_v34
		_626_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F29())
		var _627_v35 _dafny.CodePoint
		_ = _627_v35
		_627_v35 = _dafny.CodePoint('q')
		var _628_v36 _dafny.MultiSet
		_ = _628_v36
		_628_v36 = _dafny.MultiSetOf(_627_v35)
		var _629_v37 _dafny.Set
		_ = _629_v37
		_629_v37 = _dafny.SetOf((_626_v34).Cardinality(), (_628_v36).Cardinality(), (_this).F38(), _dafny.IntOfInt64(859))
		var _630_v38 _dafny.Map
		_ = _630_v38
		_630_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(419))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg62 _dafny.Int) interface{} {
				return coer62(arg62)
			}
		}((func(_631_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_632_i7 _dafny.Int) _dafny.CodePoint {
				return _631_v35
			}
		})(_627_v35)))).Cardinality()), _dafny.SetOf((_this).F38()))
		r0 = (func() _dafny.Set {
			var _coll29 = _dafny.NewBuilder()
			_ = _coll29
			for _iter36 := _dafny.Iterate((_581_v1).Elements()); ; {
				_compr_29, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _633_v31 _dafny.Int
				_633_v31 = interface{}(_compr_29).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_581_v1, _633_v31) {
					_coll29.Add((_633_v31).Times((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-479)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (func() _dafny.Set {
						var _coll30 = _dafny.NewBuilder()
						_ = _coll30
						for _iter37 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('q'))).Elements()); ; {
							_compr_30, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _634_v32 _dafny.CodePoint
							_634_v32 = interface{}(_compr_30).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('q')), _634_v32) {
								_coll30.Add(_634_v32)
							}
						}
						return _coll30.ToSet()
					}()).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(317)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Map {
						var _coll31 = _dafny.NewMapBuilder()
						_ = _coll31
						for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-483), _dafny.IntOfInt64(-255))); ; {
							_compr_31, _ok38 := _iter38()
							if !_ok38 {
								break
							}
							var _635_v33 _dafny.Int
							_635_v33 = interface{}(_compr_31).(_dafny.Int)
							if ((_dafny.IntOfInt64(-483)).Cmp(_635_v33) <= 0) && ((_635_v33).Cmp(_dafny.IntOfInt64(-255)) < 0) {
								_coll31.Add(Companion_Default___.SafeModuloInt(_635_v33, (_dafny.SetOf(_dafny.IntOfInt64(-81))).Cardinality()), true)
							}
						}
						return _coll31.ToMap()
					}()).Cardinality()))).Cardinality()))
				}
			}
			return _coll29.ToSet()
		}()).Intersection((_629_v37).Intersection((func() _dafny.Set {
			if (_630_v38).Contains((_this).F38()) {
				return (_630_v38).Get((_this).F38()).(_dafny.Set)
			}
			return _629_v37
		})()))
		var _636_v39 D3
		_ = _636_v39
		_636_v39 = Companion_D3_.Create_DC12_(Companion_Default___.Fm29((_this).F29(), (_this).F38(), (_this).F29(), globalState))
		var _637_v40 D3
		_ = _637_v40
		_637_v40 = Companion_D3_.Create_DC11_((_580_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.IntOfUint32((_580_v0).Cardinality()))).Uint32()).(bool), (_this).F29())
		var _pat_let_tv15 = _637_v40
		_ = _pat_let_tv15
		r1 = func(_source10 D3) _dafny.Int {
			if _source10.Is_DC11() {
				var _638___mcc_h4 bool = _source10.Get_().(D3_DC11).Cf19
				_ = _638___mcc_h4
				var _639___mcc_h5 bool = _source10.Get_().(D3_DC11).Cf20
				_ = _639___mcc_h5
				var _640_cf20 bool = _639___mcc_h5
				_ = _640_cf20
				var _641_cf19 bool = _638___mcc_h4
				_ = _641_cf19
				return _dafny.IntOfInt64(967)
			} else if _source10.Is_DC10() {
				var _642___mcc_h6 _dafny.MultiSet = _source10.Get_().(D3_DC10).Cf18
				_ = _642___mcc_h6
				var _643_cf18 _dafny.MultiSet = _642___mcc_h6
				_ = _643_cf18
				return (_this).F38()
			} else {
				var _644___mcc_h7 D3 = _source10.Get_().(D3_DC12).Cf21
				_ = _644___mcc_h7
				var _645_cf21 D3 = _644___mcc_h7
				_ = _645_cf21
				return (_this).F38()
			}
		}(func(_pat_let8_0 D3) D3 {
			return func(_646_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let9_0 D3) D3 {
					return func(_647_dt__update_hcf21_h0 D3) D3 {
						return Companion_D3_.Create_DC12_(_647_dt__update_hcf21_h0)
					}(_pat_let9_0)
				}(_pat_let_tv15)
			}(_pat_let8_0)
		}(_636_v39))
		return r0, r1
	}
}
func (_this *C2) M9(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var _648_v0 _dafny.MultiSet
		_ = _648_v0
		_648_v0 = _dafny.MultiSetOf(p1, (_dafny.Zero).Minus((_this).F38()), p1)
		var _649_v1 _dafny.Map
		_ = _649_v1
		_649_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F38())
		var _650_v2 _dafny.Sequence
		_ = _650_v2
		_650_v2 = _dafny.SeqOf((_this).F38())
		var _651_v3 _dafny.Map
		_ = _651_v3
		_651_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v1, (_dafny.MultiSetFromSeq(_650_v2)).Update(_dafny.IntOfUint32((_this.F30()).Cardinality()), Companion_Default___.Abs(p1)))
		_648_v0 = (func() _dafny.MultiSet {
			if (_651_v3).Contains(_649_v1) {
				return (_651_v3).Get(_649_v1).(_dafny.MultiSet)
			}
			return (_648_v0).Difference(_648_v0)
		})()
		(globalState).F13 = (_this).F29()
		(globalState).F13 = p2
		(globalState).F28 = !((_this).F29()) || (p2)
		var _652_v4 D0
		_ = _652_v4
		_652_v4 = Companion_D0_.Create_DC0_((func() bool {
			if (_this).F29() {
				return p3
			}
			return !((_this).F29())
		})())
		_652_v4 = (func() D0 {
			if (_this).F29() {
				return _652_v4
			}
			return _652_v4
		})()
		var _653_v5 _dafny.Array
		_ = _653_v5
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_14
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) bool = (func(_654_p2 bool) func(_dafny.Int) bool {
				return func(_655_i0 _dafny.Int) bool {
					return _654_p2
				}
			})(p2)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw106 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw106).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw106).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_653_v5 = _nw106
		var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_653_v5), 0))
		_ = _index104
		(_653_v5).ArraySet1(false, (_index104).Int())
		var _656_v6 _dafny.CodePoint
		_ = _656_v6
		_656_v6 = _dafny.CodePoint('q')
		var _657_v7 _dafny.Set
		_ = _657_v7
		_657_v7 = _dafny.SetOf(_656_v6)
		var _pat_let_tv16 = p2
		_ = _pat_let_tv16
		var _pat_let_tv17 = p3
		_ = _pat_let_tv17
		var _pat_let_tv18 = p3
		_ = _pat_let_tv18
		var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_653_v5), 0))
		_ = _index105
		(_653_v5).ArraySet1(func(_source11 D2) bool {
			if _source11.Is_DC7() {
				var _658___mcc_h0 bool = _source11.Get_().(D2_DC7).Cf11
				_ = _658___mcc_h0
				var _659_cf11 bool = _658___mcc_h0
				_ = _659_cf11
				return (_pat_let_tv16) && (_659_cf11)
			} else if _source11.Is_DC8() {
				var _660___mcc_h1 _dafny.CodePoint = _source11.Get_().(D2_DC8).Cf12
				_ = _660___mcc_h1
				var _661___mcc_h2 bool = _source11.Get_().(D2_DC8).Cf13
				_ = _661___mcc_h2
				var _662_cf13 bool = _661___mcc_h2
				_ = _662_cf13
				var _663_cf12 _dafny.CodePoint = _660___mcc_h1
				_ = _663_cf12
				return _662_cf13
			} else if _source11.Is_DC9() {
				var _664___mcc_h3 _dafny.Sequence = _source11.Get_().(D2_DC9).Cf14
				_ = _664___mcc_h3
				var _665___mcc_h4 _dafny.Int = _source11.Get_().(D2_DC9).Cf15
				_ = _665___mcc_h4
				var _666___mcc_h5 bool = _source11.Get_().(D2_DC9).Cf16
				_ = _666___mcc_h5
				var _667___mcc_h6 _dafny.Int = _source11.Get_().(D2_DC9).Cf17
				_ = _667___mcc_h6
				var _668_cf17 _dafny.Int = _667___mcc_h6
				_ = _668_cf17
				var _669_cf16 bool = _666___mcc_h5
				_ = _669_cf16
				var _670_cf15 _dafny.Int = _665___mcc_h4
				_ = _670_cf15
				var _671_cf14 _dafny.Sequence = _664___mcc_h3
				_ = _671_cf14
				return _pat_let_tv17
			} else {
				var _672___mcc_h7 _dafny.MultiSet = _source11.Get_().(D2_DC6).Cf10
				_ = _672___mcc_h7
				var _673_cf10 _dafny.MultiSet = _672___mcc_h7
				_ = _673_cf10
				return !(_pat_let_tv18)
			}
		}(func(_pat_let10_0 D2) D2 {
			return func(_675_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let11_0 _dafny.Sequence) D2 {
					return func(_676_dt__update_hcf14_h0 _dafny.Sequence) D2 {
						return func(_pat_let12_0 bool) D2 {
							return func(_677_dt__update_hcf16_h0 bool) D2 {
								return func(_pat_let13_0 _dafny.Int) D2 {
									return func(_678_dt__update_hcf15_h0 _dafny.Int) D2 {
										return Companion_D2_.Create_DC9_(_676_dt__update_hcf14_h0, _678_dt__update_hcf15_h0, _677_dt__update_hcf16_h0, (_675_dt__update__tmp_h0).Dtor_cf17())
									}(_pat_let13_0)
								}(_dafny.IntOfInt64(271))
							}(_pat_let12_0)
						}(true)
					}(_pat_let11_0)
				}(_this.F30())
			}(_pat_let10_0)
		}(Companion_D2_.Create_DC9_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg63 _dafny.Int) interface{} {
				return coer63(arg63)
			}
		}(func(_674_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), (_dafny.Zero).Minus((_657_v7).Cardinality()), p3, p1))), (_index105).Int())
		var _679_v8 T0
		_ = _679_v8
		var _nw107 *C1 = New_C1_()
		_ = _nw107
		_nw107.Ctor__(p2)
		_679_v8 = _nw107
		var _680_v9 D10
		_ = _680_v9
		_680_v9 = Companion_D10_.Create_DC30_(_679_v8)
		var _681_v10 _dafny.Set
		_ = _681_v10
		_681_v10 = _dafny.SetOf(_679_v8, (_680_v9).Dtor_cf49(), _679_v8)
		var _682_v11 _dafny.Set
		_ = _682_v11
		_682_v11 = _dafny.SetOf((_681_v10).Intersection(_681_v10))
		r0 = _682_v11
		return r0
	}
}
func (_this *C2) F38() _dafny.Int {
	{
		return _this._f38
	}
}
func (_this *C2) F39() _dafny.Set {
	{
		return _this._f39
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f30 _dafny.Sequence
	_f29 bool
	_f35 bool
	_f36 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f30 = _dafny.EmptySeq
	_this._f29 = false
	_this._f35 = false
	_this._f36 = _dafny.Zero
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

func (_this *C3) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C3) F30_set_(value _dafny.Sequence) {
	_this._f30 = value
}
func (_this *C3) F29() bool {
	return _this._f29
}
func (_this *C3) Ctor__(f35 bool, f36 _dafny.Int, f29 bool, f30 _dafny.Sequence) {
	{
		(_this)._f35 = f35
		(_this)._f36 = f36
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C3) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _683_v0 _dafny.Array
		_ = _683_v0
		var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
		_ = _nw108
		_683_v0 = _nw108
		var _684_v1 _dafny.Array
		_ = _684_v1
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_15
		var _nw109 _dafny.Array
		_ = _nw109
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw109 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = func(_685_i4 _dafny.Int) bool {
				return (_this).F35()
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw109 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw109).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw109).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_684_v1 = _nw109
		var _686_v2 _dafny.Map
		_ = _686_v2
		_686_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F35(), _684_v1)
		var _687_v3 _dafny.CodePoint
		_ = _687_v3
		_687_v3 = _dafny.CodePoint('f')
		var _688_v4 _dafny.Array
		_ = _688_v4
		var _nwElement0_18 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}(func(_689_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))
		_ = _nwElement0_18
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(24))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_18, 0)
		(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg65 _dafny.Int) interface{} {
				return coer65(arg65)
			}
		}(func(_690_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), 1)
		(_nw110).ArraySet1(_this.F30(), 2)
		(_nw110).ArraySet1(_this.F30(), 3)
		(_nw110).ArraySet1(_this.F30(), 4)
		(_nw110).ArraySet1(_this.F30(), 5)
		(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_691_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), 6)
		(_nw110).ArraySet1(_this.F30(), 7)
		(_nw110).ArraySet1(_this.F30(), 8)
		(_nw110).ArraySet1(_this.F30(), 9)
		(_nw110).ArraySet1(_this.F30(), 10)
		(_nw110).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cdiohjl"), 11)
		(_nw110).ArraySet1(_this.F30(), 12)
		(_nw110).ArraySet1(_this.F30(), 13)
		(_nw110).ArraySet1(_this.F30(), 14)
		(_nw110).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg67 _dafny.Int) interface{} {
				return coer67(arg67)
			}
		}(func(_692_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), 15)
		(_nw110).ArraySet1(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_686_v2).Cardinality(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _687_v3), 16)
		(_nw110).ArraySet1(_this.F30(), 17)
		(_nw110).ArraySet1(_this.F30(), 18)
		(_nw110).ArraySet1(_this.F30(), 19)
		(_nw110).ArraySet1(_this.F30(), 20)
		(_nw110).ArraySet1(_this.F30(), 21)
		(_nw110).ArraySet1(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _687_v3), 22)
		(_nw110).ArraySet1(_this.F30(), 23)
		_688_v4 = _nw110
		var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_683_v0), 0))
		_ = _index106
		(_683_v0).ArraySet1(_688_v4, (_index106).Int())
		var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_683_v0), 0))
		_ = _index107
		(_683_v0).ArraySet1(_688_v4, (_index107).Int())
		var _693_v5 _dafny.Sequence
		_ = _693_v5
		_693_v5 = _dafny.SeqOf(p0, p0)
		var _source12 D1 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Concatenate(_693_v5, _dafny.SeqOf(true)))
		_ = _source12
		if _source12.Is_DC4() {
			var _694___mcc_h0 bool = _source12.Get_().(D1_DC4).Cf8
			_ = _694___mcc_h0
			var _695_cf8 bool = _694___mcc_h0
			_ = _695_cf8
			var _696_v6 _dafny.Array
			_ = _696_v6
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_16
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Int = func(_697_i5 _dafny.Int) _dafny.Int {
					return (_697_i5).Minus((_this).F36())
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw111 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw111).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw111).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_696_v6 = _nw111
			(globalState).F17 = _696_v6
			(globalState).F15 = p0
			var _698_v7 _dafny.Map
			_ = _698_v7
			_698_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F36(), _687_v3)
			var _699_v8 D2
			_ = _699_v8
			_699_v8 = Companion_D2_.Create_DC8_((func() _dafny.CodePoint {
				if (_698_v7).Contains((_this).F36()) {
					return (_698_v7).Get((_this).F36()).(_dafny.CodePoint)
				}
				return _687_v3
			})(), (_this).F29())
			var _source13 D2 = _699_v8
			_ = _source13
			if _source13.Is_DC7() {
				var _700___mcc_h3 bool = _source13.Get_().(D2_DC7).Cf11
				_ = _700___mcc_h3
				var _701_cf11 bool = _700___mcc_h3
				_ = _701_cf11
				var _702_v9 _dafny.Sequence
				_ = _702_v9
				_702_v9 = _dafny.SeqOf(_693_v5)
				r0 = _702_v9
				var _703_v10 _dafny.Array
				_ = _703_v10
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(23))
				_ = _nw112
				_703_v10 = _nw112
				var _704_v11 _dafny.Map
				_ = _704_v11
				_704_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uogups"), _this.F30()), _703_v10)
				_704_v11 = (_704_v11).Update(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _687_v3)), _703_v10)
				r1 = _701_cf11
				var _705_v12 _dafny.Set
				_ = _705_v12
				_705_v12 = _dafny.SetOf(_701_cf11, (_this).F35())
				(globalState).F19 = ((_705_v12).Difference(_705_v12)).Cardinality()
			} else if _source13.Is_DC8() {
				var _706___mcc_h4 _dafny.CodePoint = _source13.Get_().(D2_DC8).Cf12
				_ = _706___mcc_h4
				var _707___mcc_h5 bool = _source13.Get_().(D2_DC8).Cf13
				_ = _707___mcc_h5
				var _708_cf13 bool = _707___mcc_h5
				_ = _708_cf13
				var _709_cf12 _dafny.CodePoint = _706___mcc_h4
				_ = _709_cf12
				var _710_v13 D2
				_ = _710_v13
				_710_v13 = Companion_D2_.Create_DC7_((_this).F29())
				(globalState).F15 = (_710_v13).Dtor_cf11()
				var _711_v14 _dafny.Map
				_ = _711_v14
				_711_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(22), _684_v1)
				_711_v14 = (_711_v14).Update(_dafny.IntOfUint32((_693_v5).Cardinality()), (func() _dafny.Array {
					if (_this).F29() {
						return _684_v1
					}
					return _684_v1
				})())
				var _712_v15 D0
				_ = _712_v15
				_712_v15 = Companion_D0_.Create_DC1_(_684_v1, (_this).F36(), (_this).F36())
				var _rhs123 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_this).F35() {
						return (_this).F36()
					}
					return Companion_Default___.Fm9(globalState)
				})(), (_dafny.IntOfUint32((_this.F30()).Cardinality())).Times((_this).F36()))
				_ = _rhs123
				var _rhs124 _dafny.Int = (_this).F36()
				_ = _rhs124
				var _rhs125 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((Companion_Default___.Fm10((_this).F36(), (_this).F36(), globalState)) == ((_693_v5).Select((Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_693_v5).Cardinality()))).Uint32()).(bool))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F36()), _dafny.IntOfUint32((_dafny.SeqOf((Companion_Default___.Fm10((_this).F36(), (_this).F36(), globalState)) == ((_693_v5).Select((Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_693_v5).Cardinality()))).Uint32()).(bool)))).Cardinality()))).Uint32(), (_this).F35())
				_ = _rhs125
				var _rhs126 _dafny.Int = (_712_v15).Dtor_cf2()
				_ = _rhs126
				var _rhs127 D2 = (func() D2 {
					if _695_cf8 {
						return _710_v13
					}
					return _710_v13
				})()
				_ = _rhs127
				var _lhs114 *GlobalState = globalState
				_ = _lhs114
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				_lhs114.F1 = _rhs123
				_lhs115.F14 = _rhs124
				_693_v5 = _rhs125
				_lhs116.F21 = _rhs126
				_710_v13 = _rhs127
				(globalState).F1 = ((_this).F36()).Plus((_this).F36())
			} else if _source13.Is_DC9() {
				var _713___mcc_h6 _dafny.Sequence = _source13.Get_().(D2_DC9).Cf14
				_ = _713___mcc_h6
				var _714___mcc_h7 _dafny.Int = _source13.Get_().(D2_DC9).Cf15
				_ = _714___mcc_h7
				var _715___mcc_h8 bool = _source13.Get_().(D2_DC9).Cf16
				_ = _715___mcc_h8
				var _716___mcc_h9 _dafny.Int = _source13.Get_().(D2_DC9).Cf17
				_ = _716___mcc_h9
				var _717_cf17 _dafny.Int = _716___mcc_h9
				_ = _717_cf17
				var _718_cf16 bool = _715___mcc_h8
				_ = _718_cf16
				var _719_cf15 _dafny.Int = _714___mcc_h7
				_ = _719_cf15
				var _720_cf14 _dafny.Sequence = _713___mcc_h6
				_ = _720_cf14
				_687_v3 = _687_v3
				(globalState).F2 = (_dafny.Zero).Minus((_this).F36())
				var _721_v16 _dafny.Array
				_ = _721_v16
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(11))
				_ = _nw113
				_721_v16 = _nw113
				var _722_v17 D1
				_ = _722_v17
				_722_v17 = Companion_D1_.Create_DC3_(_693_v5)
				var _723_v18 D1
				_ = _723_v18
				_723_v18 = Companion_D1_.Create_DC5_(_722_v17)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_721_v16), 0))
				_ = _index108
				(_721_v16).ArraySet1(_723_v18, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_721_v16), 0))
				_ = _index109
				(_721_v16).ArraySet1(_723_v18, (_index109).Int())
				var _724_v19 _dafny.Map
				_ = _724_v19
				_724_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_696_v6, (_this).F36())
				(globalState).F2 = (_724_v19).Cardinality()
			} else {
				var _725___mcc_h10 _dafny.MultiSet = _source13.Get_().(D2_DC6).Cf10
				_ = _725___mcc_h10
				var _726_cf10 _dafny.MultiSet = _725___mcc_h10
				_ = _726_cf10
				(globalState).F1 = (func() _dafny.Int {
					if (_this).F35() {
						return (_this).F36()
					}
					return ((_this).F36()).Minus((_this).F36())
				})()
				(globalState).F11 = _this.F30()
				var _727_v20 *C0
				_ = _727_v20
				var _nw114 *C0 = New_C0_()
				_ = _nw114
				_nw114.Ctor__(p0)
				_727_v20 = _nw114
				var _728_v21 _dafny.Array
				_ = _728_v21
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_17
				var _nw115 _dafny.Array
				_ = _nw115
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw115 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.CodePoint = (func(_729_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_730_i6 _dafny.Int) _dafny.CodePoint {
							return _729_v3
						}
					})(_687_v3)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw115 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw115).ArraySet1CodePoint(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw115).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_728_v21 = _nw115
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_728_v21), 0))
				_ = _index110
				(_728_v21).ArraySet1CodePoint((_this.F30()).Select((Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_728_v21), 0))
				_ = _index111
				(_728_v21).ArraySet1CodePoint(_687_v3, (_index111).Int())
			}
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_696_v6), 0))
			_ = _index112
			(_696_v6).ArraySet1((_this).F36(), (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_696_v6), 0))
			_ = _index113
			(_696_v6).ArraySet1(_dafny.IntOfUint32((_this.F30()).Cardinality()), (_index113).Int())
		} else if _source12.Is_DC3() {
			var _731___mcc_h1 _dafny.Sequence = _source12.Get_().(D1_DC3).Cf7
			_ = _731___mcc_h1
			var _732_cf7 _dafny.Sequence = _731___mcc_h1
			_ = _732_cf7
			var _733_v22 _dafny.Map
			_ = _733_v22
			_733_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_cf7, (_dafny.Zero).Minus((_this).F36()))
			_733_v22 = ((_733_v22).Merge((Companion_Default___.Fm13((_this).F36(), (func() _dafny.Map {
				var _coll32 = _dafny.NewMapBuilder()
				_ = _coll32
				for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-784), _dafny.IntOfInt64(-112))); ; {
					_compr_32, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _734_v23 _dafny.Int
					_734_v23 = interface{}(_compr_32).(_dafny.Int)
					if ((_dafny.IntOfInt64(-784)).Cmp(_734_v23) <= 0) && ((_734_v23).Cmp(_dafny.IntOfInt64(-112)) < 0) {
						_coll32.Add(Companion_Default___.SafeModuloInt(_734_v23, _dafny.IntOfInt64(471)), (_this).F36())
					}
				}
				return _coll32.ToMap()
			}()).Cardinality(), (_this).F36(), !((_this).F35()), globalState)).Dtor_cf22())).Merge(_733_v22)
			(globalState).F20 = _this.F30()
			(globalState).F19 = (_this).F36()
			var _735_v24 D1
			_ = _735_v24
			_735_v24 = Companion_D1_.Create_DC4_(!(p0) || (p0))
			var _source14 D1 = _735_v24
			_ = _source14
			if _source14.Is_DC4() {
				var _736___mcc_h11 bool = _source14.Get_().(D1_DC4).Cf8
				_ = _736___mcc_h11
				var _737_cf8 bool = _736___mcc_h11
				_ = _737_cf8
				var _738_v25 _dafny.Array
				_ = _738_v25
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_18
				var _nw116 _dafny.Array
				_ = _nw116
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw116 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.CodePoint = (func(_739_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_740_i7 _dafny.Int) _dafny.CodePoint {
							return _739_v3
						}
					})(_687_v3)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw116 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw116).ArraySet1CodePoint(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw116).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_738_v25 = _nw116
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_738_v25), 0))
				_ = _index114
				(_738_v25).ArraySet1CodePoint(_687_v3, (_index114).Int())
				var _741_v26 _dafny.MultiSet
				_ = _741_v26
				_741_v26 = _dafny.MultiSetOf((_this).F36())
				var _742_v27 _dafny.Map
				_ = _742_v27
				_742_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F35()), _741_v26)
				var _743_v28 D3
				_ = _743_v28
				_743_v28 = Companion_D3_.Create_DC10_((func() _dafny.MultiSet {
					if (_742_v27).Contains(_737_cf8) {
						return (_742_v27).Get(_737_cf8).(_dafny.MultiSet)
					}
					return _741_v26
				})())
				var _744_v29 D3
				_ = _744_v29
				_744_v29 = Companion_D3_.Create_DC12_(_743_v28)
				var _745_v30 _dafny.Map
				_ = _745_v30
				_745_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(774)).Cmp((_this).F36()) <= 0, _744_v29)
				var _746_v31 D2
				_ = _746_v31
				_746_v31 = Companion_D2_.Create_DC8_(_687_v3, p0)
				var _747_v32 _dafny.Array
				_ = _747_v32
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw117
				_747_v32 = _nw117
				var _pat_let_tv19 = _743_v28
				_ = _pat_let_tv19
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_738_v25), 0))
				_ = _index115
				var _rhs128 _dafny.CodePoint = (_746_v31).Dtor_cf12()
				_ = _rhs128
				var _rhs129 _dafny.Map = (((_745_v30).Update((_this).F35(), _744_v29)).Merge((_745_v30).Update(false, func(_pat_let14_0 D3) D3 {
					return func(_748_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let15_0 D3) D3 {
							return func(_749_dt__update_hcf21_h0 D3) D3 {
								return Companion_D3_.Create_DC12_(_749_dt__update_hcf21_h0)
							}(_pat_let15_0)
						}(_pat_let_tv19)
					}(_pat_let14_0)
				}(_744_v29)))).Merge(_745_v30)
				_ = _rhs129
				var _rhs130 _dafny.Array = _684_v1
				_ = _rhs130
				var _rhs131 _dafny.Array = _747_v32
				_ = _rhs131
				var _lhs117 _dafny.Array = _738_v25
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_738_v25), 0))
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				(_lhs117).ArraySet1CodePoint(_rhs128, (_lhs118).Int())
				_745_v30 = _rhs129
				_684_v1 = _rhs130
				_lhs119.F17 = _rhs131
				var _750_v33 *C0
				_ = _750_v33
				var _nw118 *C0 = New_C0_()
				_ = _nw118
				_nw118.Ctor__((func() bool {
					if (_this).F29() {
						return true
					}
					return p0
				})())
				_750_v33 = _nw118
				var _751_v34 _dafny.Map
				_ = _751_v34
				_751_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_cf8, _dafny.IntOfUint32((_693_v5).Cardinality()))
				var _752_v36 _dafny.MultiSet
				_ = _752_v36
				_752_v36 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_750_v33.F37, (_this).F36()), _751_v34, Companion_Default___.Fm14((_this).F35(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(900))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_753_i9 _dafny.Int) _dafny.Int {
					return (_this).F36()
				})), _737_cf8, (_dafny.SetOf(_684_v1)).Cardinality(), globalState), _751_v34)
				var _754_v38 _dafny.Set
				_ = _754_v38
				_754_v38 = _dafny.SetOf(p0, _750_v33.F37)
				var _755_v39 _dafny.Set
				_ = _755_v39
				_755_v39 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_754_v38).Cardinality()))
				var _756_v40 _dafny.Map
				_ = _756_v40
				_756_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_751_v34, true)
				var _757_v42 _dafny.Map
				_ = _757_v42
				_757_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_751_v34, (_this).F36())
				var _758_v44 _dafny.Array
				_ = _758_v44
				var _nwElement0_19 _dafny.Set = _dafny.SetOf(_751_v34, _751_v34)
				_ = _nwElement0_19
				var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(9))
				_ = _nw119
				(_nw119).ArraySet1(_nwElement0_19, 0)
				(_nw119).ArraySet1(func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter40 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg69 _dafny.Int) interface{} {
							return coer69(arg69)
						}
					}((func(_759_v34 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_760_i8 _dafny.Int) _dafny.Map {
							return _759_v34
						}
					})(_751_v34)))).Elements()); ; {
						_compr_33, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _761_v35 _dafny.Map
						_761_v35 = interface{}(_compr_33).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg70 _dafny.Int) interface{} {
								return coer70(arg70)
							}
						}((func(_762_v34 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_760_i8 _dafny.Int) _dafny.Map {
								return _762_v34
							}
						})(_751_v34))), _761_v35) {
							_coll33.Add(_761_v35)
						}
					}
					return _coll33.ToSet()
				}(), 1)
				(_nw119).ArraySet1(func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter41 := _dafny.Iterate((_752_v36).Elements()); ; {
						_compr_34, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _763_v37 _dafny.Map
						_763_v37 = interface{}(_compr_34).(_dafny.Map)
						if (_752_v36).Contains(_763_v37) {
							_coll34.Add(_763_v37)
						}
					}
					return _coll34.ToSet()
				}(), 2)
				(_nw119).ArraySet1((_755_v39).Intersection(func() _dafny.Set {
					var _coll35 = _dafny.NewBuilder()
					_ = _coll35
					for _iter42 := _dafny.Iterate((_756_v40).Keys().Elements()); ; {
						_compr_35, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _764_v41 _dafny.Map
						_764_v41 = interface{}(_compr_35).(_dafny.Map)
						if (_756_v40).Contains(_764_v41) {
							_coll35.Add(_764_v41)
						}
					}
					return _coll35.ToSet()
				}()), 3)
				(_nw119).ArraySet1(_755_v39, 4)
				(_nw119).ArraySet1(func() _dafny.Set {
					var _coll36 = _dafny.NewBuilder()
					_ = _coll36
					for _iter43 := _dafny.Iterate(((_757_v42).Update(_751_v34, (_this).F36())).Keys().Elements()); ; {
						_compr_36, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _765_v43 _dafny.Map
						_765_v43 = interface{}(_compr_36).(_dafny.Map)
						if ((_757_v42).Update(_751_v34, (_this).F36())).Contains(_765_v43) {
							_coll36.Add(_765_v43)
						}
					}
					return _coll36.ToSet()
				}(), 5)
				(_nw119).ArraySet1(_755_v39, 6)
				(_nw119).ArraySet1(_dafny.SetOf(_751_v34, _751_v34, (_751_v34).Update(_737_cf8, (_754_v38).Cardinality()), _751_v34, _751_v34), 7)
				(_nw119).ArraySet1((_dafny.SetOf(_751_v34)).Intersection(_755_v39), 8)
				_758_v44 = _nw119
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_758_v44), 0))
				_ = _index116
				(_758_v44).ArraySet1(_755_v39, (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_758_v44), 0))
				_ = _index117
				(_758_v44).ArraySet1(Companion_Default___.Fm15((_this).F36(), globalState), (_index117).Int())
				(globalState).F19 = (_dafny.Zero).Minus((_this).F36())
			} else if _source14.Is_DC3() {
				var _766___mcc_h12 _dafny.Sequence = _source14.Get_().(D1_DC3).Cf7
				_ = _766___mcc_h12
				var _767_cf7 _dafny.Sequence = _766___mcc_h12
				_ = _767_cf7
				(globalState).F1 = _dafny.IntOfInt64(58)
				var _768_v45 _dafny.Array
				_ = _768_v45
				var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw120
				_768_v45 = _nw120
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_768_v45), 0))
				_ = _index118
				(_768_v45).ArraySet1((_this).F36(), (_index118).Int())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_768_v45), 0))
				_ = _index119
				(_768_v45).ArraySet1(_dafny.IntOfInt64(943), (_index119).Int())
				(globalState).F21 = (_768_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_768_v45), 0))).Int()).(_dafny.Int)
				(globalState).F28 = p0
			} else {
				var _769___mcc_h13 D1 = _source14.Get_().(D1_DC5).Cf9
				_ = _769___mcc_h13
				var _770_cf9 D1 = _769___mcc_h13
				_ = _770_cf9
				var _771_v46 _dafny.Set
				_ = _771_v46
				_771_v46 = _dafny.SetOf(p0, (_this).F29())
				var _772_v47 _dafny.Sequence
				_ = _772_v47
				_772_v47 = _dafny.SeqOf(_771_v46)
				(globalState).F3 = _dafny.Companion_Sequence_.Concatenate(_772_v47, _772_v47)
				(globalState).F1 = (Companion_Default___.Fm9(globalState)).Plus((_this).F36())
				var _773_v48 *C0
				_ = _773_v48
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__((func() bool {
					if (_this).F35() {
						return false
					}
					return (_this).F35()
				})())
				_773_v48 = _nw121
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_684_v1), 0))
				_ = _index120
				(_684_v1).ArraySet1((_this).F29(), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_684_v1), 0))
				_ = _index121
				(_684_v1).ArraySet1((_this).F35(), (_index121).Int())
			}
		} else {
			var _774___mcc_h2 D1 = _source12.Get_().(D1_DC5).Cf9
			_ = _774___mcc_h2
			var _775_cf9 D1 = _774___mcc_h2
			_ = _775_cf9
			var _776_v49 _dafny.Map
			_ = _776_v49
			_776_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F35(), (_this).F36())
			var _777_v50 _dafny.Array
			_ = _777_v50
			var _nwElement0_20 _dafny.Int = (_this).F36()
			_ = _nwElement0_20
			var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(17))
			_ = _nw122
			(_nw122).ArraySet1(_nwElement0_20, 0)
			(_nw122).ArraySet1(_dafny.IntOfUint32((_this.F30()).Cardinality()), 1)
			(_nw122).ArraySet1((_this).F36(), 2)
			(_nw122).ArraySet1((_this).F36(), 3)
			(_nw122).ArraySet1((_this).F36(), 4)
			(_nw122).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-598))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}(func(_778_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))).Cardinality()), 5)
			(_nw122).ArraySet1((_this).F36(), 6)
			(_nw122).ArraySet1((_this).F36(), 7)
			(_nw122).ArraySet1((_776_v49).Cardinality(), 8)
			(_nw122).ArraySet1((_this).F36(), 9)
			(_nw122).ArraySet1(_dafny.IntOfUint32((_693_v5).Cardinality()), 10)
			(_nw122).ArraySet1((_this).F36(), 11)
			(_nw122).ArraySet1((_this).F36(), 12)
			(_nw122).ArraySet1((_this).F36(), 13)
			(_nw122).ArraySet1((_this).F36(), 14)
			(_nw122).ArraySet1((_this).F36(), 15)
			(_nw122).ArraySet1((_this).F36(), 16)
			_777_v50 = _nw122
			var _779_v51 _dafny.Sequence
			_ = _779_v51
			_779_v51 = _dafny.SeqOf(_777_v50, _777_v50)
			var _780_v52 _dafny.Set
			_ = _780_v52
			_780_v52 = _dafny.SetOf(_777_v50, (_779_v51).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm9(globalState), _dafny.IntOfUint32((_779_v51).Cardinality()))).Uint32()).(_dafny.Array))
			var _781_v53 T1
			_ = _781_v53
			var _nw123 *C2 = New_C2_()
			_ = _nw123
			_nw123.Ctor__((_this).F36(), _780_v52, p0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gt"), _this.F30()))
			_781_v53 = _nw123
			_781_v53 = _781_v53
			var _782_v54 _dafny.MultiSet
			_ = _782_v54
			_782_v54 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("vccvow"))
			(globalState).F28 = (_782_v54).Contains(_this.F30())
			var _783_v55 _dafny.Array
			_ = _783_v55
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_19
			var _nw124 _dafny.Array
			_ = _nw124
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw124 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.CodePoint = (func(_784_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_785_i11 _dafny.Int) _dafny.CodePoint {
						return _784_v3
					}
				})(_687_v3)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw124 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw124).ArraySet1CodePoint(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw124).ArraySet1CodePoint(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_783_v55 = _nw124
			var _786_v56 _dafny.Map
			_ = _786_v56
			_786_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F35(), _783_v55)
			var _787_v57 _dafny.MultiSet
			_ = _787_v57
			_787_v57 = _dafny.MultiSetOf(_783_v55, _783_v55, _783_v55, (func() _dafny.Array {
				if (_786_v56).Contains(p0) {
					return (_786_v56).Get(p0).(_dafny.Array)
				}
				return _783_v55
			})())
			_787_v57 = _dafny.MultiSetOf(_783_v55)
			(globalState).F21 = Companion_Default___.SafeDivisionInt((_this).F36(), _dafny.IntOfInt64(871))
		}
		var _788_v58 _dafny.Map
		_ = _788_v58
		_788_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F36(), (_this).F36())
		_788_v58 = (_788_v58).Update(((_this).F36()).Minus((_this).F36()), (_this).F36())
		var _789_v59 _dafny.Map
		_ = _789_v59
		_789_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v3, p0)
		var _790_v60 _dafny.Set
		_ = _790_v60
		_790_v60 = _dafny.SetOf(_789_v59, _789_v59)
		var _791_v61 _dafny.Map
		_ = _791_v61
		_791_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F35(), (_790_v60).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v3, false), _789_v59)))
		_791_v61 = (_791_v61).Update((_this).F29(), _dafny.SetOf((_789_v59).Update(_dafny.CodePoint('u'), (_this).F35()), _789_v59))
		var _792_v62 _dafny.MultiSet
		_ = _792_v62
		_792_v62 = _dafny.MultiSetOf((_this).F35())
		var _793_v63 D0
		_ = _793_v63
		_793_v63 = Companion_D0_.Create_DC2_(_792_v62, _dafny.UnicodeSeqOfUtf8Bytes("ngrrpfx"), (_this).F29())
		var _794_v64 _dafny.Array
		_ = _794_v64
		var _nwElement0_21 _dafny.MultiSet = (func() _dafny.MultiSet {
			if false {
				return (_793_v63).Dtor_cf4()
			}
			return _792_v62
		})()
		_ = _nwElement0_21
		var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(16))
		_ = _nw125
		(_nw125).ArraySet1(_nwElement0_21, 0)
		(_nw125).ArraySet1(((_792_v62).Update((_this).F29(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F36())))).Union(Companion_Default___.Fm39(_687_v3, (_this).F36(), globalState)), 1)
		(_nw125).ArraySet1((_792_v62).Difference(_792_v62), 2)
		(_nw125).ArraySet1(_dafny.MultiSetOf(p0), 3)
		(_nw125).ArraySet1((_792_v62).Update(!((_this).F29()), Companion_Default___.Abs((_this).F36())), 4)
		(_nw125).ArraySet1(_dafny.MultiSetOf(p0), 5)
		(_nw125).ArraySet1(_792_v62, 6)
		(_nw125).ArraySet1(_792_v62, 7)
		(_nw125).ArraySet1((_793_v63).Dtor_cf4(), 8)
		(_nw125).ArraySet1(_dafny.MultiSetOf(p0, true, p0, (_this).F29(), p0), 9)
		(_nw125).ArraySet1(_dafny.MultiSetFromSeq(_693_v5), 10)
		(_nw125).ArraySet1(_792_v62, 11)
		(_nw125).ArraySet1((_793_v63).Dtor_cf4(), 12)
		(_nw125).ArraySet1((Companion_Default___.Fm39(_687_v3, (_this).F36(), globalState)).Intersection((_793_v63).Dtor_cf4()), 13)
		(_nw125).ArraySet1(_792_v62, 14)
		(_nw125).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm19(_dafny.CodePoint('k'), globalState)), 15)
		_794_v64 = _nw125
		_794_v64 = _794_v64
		var _795_v65 *C1
		_ = _795_v65
		var _nw126 *C1 = New_C1_()
		_ = _nw126
		_nw126.Ctor__((_this).F35())
		_795_v65 = _nw126
		var _796_v66 _dafny.Sequence
		_ = _796_v66
		_796_v66 = _dafny.SeqOf(_693_v5, _693_v5)
		r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_796_v66, _796_v66), (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_796_v66, _796_v66)).Cardinality()))).Uint32(), _693_v5), _dafny.Companion_Sequence_.Concatenate(_796_v66, _796_v66))
		r1 = false
		return r0, r1
	}
}
func (_this *C3) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _797_v0 _dafny.Sequence
		_ = _797_v0
		_797_v0 = _dafny.SeqOf((_this).F29(), (_this).F29())
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_797_v0, Companion_Default___.Fm35(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(531))).Uint32(), func(coer72 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_798_i0 _dafny.Int) D2 {
			return Companion_D2_.Create_DC8_(_dafny.CodePoint('o'), (_this).F35())
		})), (_dafny.Zero).Minus((_this).F36()), globalState)), (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_797_v0, Companion_Default___.Fm35(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(531))).Uint32(), func(coer73 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
			return func(arg73 _dafny.Int) interface{} {
				return coer73(arg73)
			}
		}(func(_799_i0 _dafny.Int) D2 {
			return Companion_D2_.Create_DC8_(_dafny.CodePoint('o'), (_this).F35())
		})), (_dafny.Zero).Minus((_this).F36()), globalState))).Cardinality()))).Uint32(), p0), (_this).F35()) {
			var _800_v1 T0
			_ = _800_v1
			var _nw127 *C1 = New_C1_()
			_ = _nw127
			_nw127.Ctor__((_this).F29())
			_800_v1 = _nw127
			var _801_v2 D10
			_ = _801_v2
			_801_v2 = Companion_D10_.Create_DC30_(_800_v1)
			(globalState).F21 = _dafny.IntOfUint32((_dafny.SeqOf(_801_v2)).Cardinality())
			var _802_v3 _dafny.Map
			_ = _802_v3
			_802_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _dafny.IntOfUint32((_this.F30()).Cardinality()))
			var _803_v4 _dafny.Sequence
			_ = _803_v4
			_803_v4 = _dafny.SeqOf((_dafny.Zero).Minus(p2), (_this).F36())
			(globalState).F1 = (Companion_Default___.Fm9(globalState)).Minus((func() _dafny.Int {
				if (_802_v3).Contains((_this).F35()) {
					return (_802_v3).Get((_this).F35()).(_dafny.Int)
				}
				return (_803_v4).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_803_v4).Cardinality()))).Uint32()).(_dafny.Int)
			})())
			var _804_v5 _dafny.Sequence
			_ = _804_v5
			var _805_v6 bool
			_ = _805_v6
			var _out12 _dafny.Sequence
			_ = _out12
			var _out13 bool
			_ = _out13
			_out12, _out13 = (_800_v1).M2(p0, globalState)
			_804_v5 = _out12
			_805_v6 = _out13
			var _806_v7 _dafny.Array
			_ = _806_v7
			var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw128
			_806_v7 = _nw128
			var _807_v8 _dafny.Set
			_ = _807_v8
			_807_v8 = _dafny.SetOf(_806_v7)
			var _808_v9 *C2
			_ = _808_v9
			var _nw129 *C2 = New_C2_()
			_ = _nw129
			_nw129.Ctor__(p2, _807_v8, p0, Companion_Default___.Fm20((_this).F29(), !((_this).F35()), _797_v0, p2, globalState))
			_808_v9 = _nw129
			var _809_v10 *C1
			_ = _809_v10
			var _nw130 *C1 = New_C1_()
			_ = _nw130
			_nw130.Ctor__(true)
			_809_v10 = _nw130
		} else {
			var _810_v11 _dafny.Array
			_ = _810_v11
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw131
			_810_v11 = _nw131
			var _811_v12 _dafny.CodePoint
			_ = _811_v12
			_811_v12 = _dafny.CodePoint('n')
			var _812_v13 _dafny.Sequence
			_ = _812_v13
			_812_v13 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _811_v12))
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_810_v11), 0))
			_ = _index122
			(_810_v11).ArraySet1(_dafny.Companion_Sequence_.Update(_812_v13, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_812_v13).Cardinality()))).Uint32(), _this.F30()), (_index122).Int())
			var _813_v14 _dafny.Array
			_ = _813_v14
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw132
			_813_v14 = _nw132
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_813_v14), 0))
			_ = _index123
			(_813_v14).ArraySet1(_dafny.IntOfInt64(62), (_index123).Int())
			var _814_v15 _dafny.MultiSet
			_ = _814_v15
			_814_v15 = _dafny.MultiSetOf(p0)
			var _815_v16 _dafny.Sequence
			_ = _815_v16
			_815_v16 = _dafny.SeqOf(_814_v15)
			var _816_v17 _dafny.Array
			_ = _816_v17
			var _nwElement0_22 bool = (_this).F29()
			_ = _nwElement0_22
			var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(7))
			_ = _nw133
			(_nw133).ArraySet1(_nwElement0_22, 0)
			(_nw133).ArraySet1((_this).F29(), 1)
			(_nw133).ArraySet1(((_815_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-495), _dafny.IntOfUint32((_815_v16).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf((_814_v15).Update(true, Companion_Default___.Abs(p2))), 2)
			(_nw133).ArraySet1((_this).F35(), 3)
			(_nw133).ArraySet1(p0, 4)
			(_nw133).ArraySet1((_this).F35(), 5)
			(_nw133).ArraySet1(p0, 6)
			_816_v17 = _nw133
			var _817_v18 _dafny.Set
			_ = _817_v18
			_817_v18 = _dafny.SetOf(p0, (_this).F29())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_810_v11), 0))
			_ = _index124
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_813_v14), 0))
			_ = _index125
			var _rhs132 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_812_v13, (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_812_v13).Cardinality()))).Uint32(), _this.F30()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}(func(_818_i1 _dafny.Int) _dafny.Sequence {
				return _this.F30()
			})))
			_ = _rhs132
			var _rhs133 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("wjkkiurju"))).Cardinality())).Minus((func() _dafny.Int {
				if (_this).F35() {
					return (_this).F36()
				}
				return _dafny.IntOfInt64(970)
			})())
			_ = _rhs133
			var _rhs134 _dafny.Int = ((func() _dafny.Set {
				if false {
					return _817_v18
				}
				return (func() _dafny.Set {
					if p0 {
						return _817_v18
					}
					return _dafny.SetOf(p0, (_this).F35())
				})()
			})()).Cardinality()
			_ = _rhs134
			var _rhs135 _dafny.Array = _816_v17
			_ = _rhs135
			var _lhs120 _dafny.Array = _810_v11
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_810_v11), 0))
			_ = _lhs121
			var _lhs122 _dafny.Array = _813_v14
			_ = _lhs122
			var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_813_v14), 0))
			_ = _lhs123
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			(_lhs120).ArraySet1(_rhs132, (_lhs121).Int())
			(_lhs122).ArraySet1(_rhs133, (_lhs123).Int())
			_lhs124.F14 = _rhs134
			_816_v17 = _rhs135
			var _819_v19 _dafny.MultiSet
			_ = _819_v19
			_819_v19 = _dafny.MultiSetOf((_813_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_813_v14), 0))).Int()).(_dafny.Int))
			(globalState).F13 = (_819_v19).IsProperSubsetOf(_819_v19)
			(globalState).F15 = ((_this).F36()).Cmp(_dafny.IntOfInt64(-139)) != 0
			var _820_v20 _dafny.Array
			_ = _820_v20
			var _nwElement0_23 _dafny.CodePoint = _811_v12
			_ = _nwElement0_23
			var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(4))
			_ = _nw134
			(_nw134).ArraySet1CodePoint(_nwElement0_23, 0)
			(_nw134).ArraySet1CodePoint(_811_v12, 1)
			(_nw134).ArraySet1CodePoint(_dafny.CodePoint('r'), 2)
			(_nw134).ArraySet1CodePoint(_811_v12, 3)
			_820_v20 = _nw134
			var _821_v21 _dafny.Sequence
			_ = _821_v21
			_821_v21 = _dafny.SeqOf(_820_v20, _820_v20)
			var _822_v22 _dafny.Sequence
			_ = _822_v22
			_822_v22 = _dafny.SeqOf(_820_v20, _820_v20, _820_v20, (_821_v21).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_821_v21).Cardinality()))).Uint32()).(_dafny.Array), _820_v20)
			_820_v20 = (_822_v22).Select((Companion_Default___.SafeIndex((p2).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_797_v0).Cardinality()))), _dafny.IntOfUint32((_822_v22).Cardinality()))).Uint32()).(_dafny.Array)
			var _rhs136 _dafny.Array = _813_v14
			_ = _rhs136
			var _rhs137 bool = Companion_Default___.Fm33(((_this).F35()) && ((_this).F29()), (_813_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_813_v14), 0))).Int()).(_dafny.Int), p2, globalState)
			_ = _rhs137
			var _lhs125 *GlobalState = globalState
			_ = _lhs125
			var _lhs126 *GlobalState = globalState
			_ = _lhs126
			_lhs125.F17 = _rhs136
			_lhs126.F13 = _rhs137
		}
		var _823_v23 _dafny.Map
		_ = _823_v23
		_823_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F36(), (_this).F36())
		var _824_v24 _dafny.CodePoint
		_ = _824_v24
		_824_v24 = _dafny.CodePoint('h')
		var _825_v25 _dafny.Map
		_ = _825_v25
		_825_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v24, (_this).F35())
		var _826_v26 _dafny.Map
		_ = _826_v26
		_826_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F35(), (_825_v25).Cardinality())
		var _827_v27 D4
		_ = _827_v27
		_827_v27 = Companion_D4_.Create_DC15_((func() _dafny.Int {
			if (_823_v23).Contains(_dafny.IntOfInt64(-597)) {
				return (_823_v23).Get(_dafny.IntOfInt64(-597)).(_dafny.Int)
			}
			return (_826_v26).Cardinality()
		})(), (_this).F36())
		_827_v27 = func(_pat_let16_0 D4) D4 {
			return func(_828_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let17_0 _dafny.Int) D4 {
					return func(_829_dt__update_hcf25_h0 _dafny.Int) D4 {
						return Companion_D4_.Create_DC15_((_828_dt__update__tmp_h0).Dtor_cf24(), _829_dt__update_hcf25_h0)
					}(_pat_let17_0)
				}((_this).F36())
			}(_pat_let16_0)
		}(Companion_D4_.Create_DC15_((_this).F36(), _dafny.IntOfUint32((_this.F30()).Cardinality())))
		if (_this).F29() {
			var _830_v28 _dafny.Sequence
			_ = _830_v28
			_830_v28 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), (_this).F36())
			(globalState).F21 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_830_v28, _830_v28), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg75 _dafny.Int) interface{} {
					return coer75(arg75)
				}
			}((func(_831_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_832_i2 _dafny.Int) _dafny.Int {
					return _831_p2
				}
			})(p2))))).Cardinality())
			var _833_v29 _dafny.Array
			_ = _833_v29
			var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw135
			_833_v29 = _nw135
			var _834_v30 _dafny.Map
			_ = _834_v30
			_834_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _833_v29)
			var _835_v31 D10
			_ = _835_v31
			_835_v31 = Companion_D10_.Create_DC33_(p0, (_this).F36())
			_834_v30 = (_834_v30).Update((_835_v31).Dtor_cf55(), _833_v29)
			(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F36()), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _824_v24)).Cardinality()), Companion_Default___.Fm9(globalState)))
			r0 = _this.F30()
			var _836_v32 D6
			_ = _836_v32
			_836_v32 = Companion_D6_.Create_DC22_((_this).F36(), (_this).F36(), (_this).F29())
			var _837_v33 D9
			_ = _837_v33
			_837_v33 = Companion_D9_.Create_DC29_((_this).F36(), (_836_v32).Dtor_cf34(), (_this).F29())
			var _838_v34 D9
			_ = _838_v34
			_838_v34 = Companion_D9_.Create_DC27_(_dafny.SeqOf((_837_v33).Dtor_cf47(), p2, (_this).F36(), (_this).F36()))
			var _839_v35 _dafny.Set
			_ = _839_v35
			_839_v35 = _dafny.SetOf(_838_v34, _838_v34, _838_v34)
			(globalState).F14 = (func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter44 := _dafny.Iterate(((_839_v35).Union(_839_v35)).Elements()); ; {
					_compr_37, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _840_v36 D9
					_840_v36 = interface{}(_compr_37).(D9)
					if ((_839_v35).Union(_839_v35)).Contains(_840_v36) {
						_coll37.Add(_840_v36)
					}
				}
				return _coll37.ToSet()
			}()).Cardinality()
		} else {
			(globalState).F18 = true
			var _841_v37 _dafny.MultiSet
			_ = _841_v37
			_841_v37 = _dafny.MultiSetOf((_this).F35())
			var _842_v38 D0
			_ = _842_v38
			_842_v38 = Companion_D0_.Create_DC2_((_841_v37).Update((_this).F35(), Companion_Default___.Abs(p2)), _this.F30(), (_this).F29())
			var _843_v39 _dafny.Array
			_ = _843_v39
			var _nwElement0_24 bool = p0
			_ = _nwElement0_24
			var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(23))
			_ = _nw136
			(_nw136).ArraySet1(_nwElement0_24, 0)
			(_nw136).ArraySet1((_this).F35(), 1)
			(_nw136).ArraySet1((_this).F35(), 2)
			(_nw136).ArraySet1(p0, 3)
			(_nw136).ArraySet1((_this).F35(), 4)
			(_nw136).ArraySet1((_this).F35(), 5)
			(_nw136).ArraySet1(p0, 6)
			(_nw136).ArraySet1(p0, 7)
			(_nw136).ArraySet1((_this).F29(), 8)
			(_nw136).ArraySet1((_this).F29(), 9)
			(_nw136).ArraySet1((_this).F29(), 10)
			(_nw136).ArraySet1((_this).F35(), 11)
			(_nw136).ArraySet1((_this).F29(), 12)
			(_nw136).ArraySet1((_this).F35(), 13)
			(_nw136).ArraySet1((_this).F29(), 14)
			(_nw136).ArraySet1(p0, 15)
			(_nw136).ArraySet1((_this).F29(), 16)
			(_nw136).ArraySet1(p0, 17)
			(_nw136).ArraySet1(p0, 18)
			(_nw136).ArraySet1((_this).F29(), 19)
			(_nw136).ArraySet1((_this).F29(), 20)
			(_nw136).ArraySet1(p0, 21)
			(_nw136).ArraySet1((_this).F35(), 22)
			_843_v39 = _nw136
			var _844_v40 _dafny.Sequence
			_ = _844_v40
			_844_v40 = _dafny.SeqOf(_843_v39)
			var _845_v41 _dafny.Map
			_ = _845_v41
			_845_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F35())
			var _846_v42 _dafny.Array
			_ = _846_v42
			var _nwElement0_25 D0 = Companion_D0_.Create_DC2_(_841_v37, _this.F30(), (_this).F29())
			_ = _nwElement0_25
			var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(15))
			_ = _nw137
			(_nw137).ArraySet1(_nwElement0_25, 0)
			(_nw137).ArraySet1((func() D0 {
				if (_this).F29() {
					return _842_v38
				}
				return _842_v38
			})(), 1)
			(_nw137).ArraySet1(_842_v38, 2)
			(_nw137).ArraySet1(_842_v38, 3)
			(_nw137).ArraySet1(Companion_Default___.Fm40(p0, globalState), 4)
			(_nw137).ArraySet1(Companion_D0_.Create_DC2_(_dafny.MultiSetFromSeq(_797_v0), (Companion_D2_.Create_DC9_(_this.F30(), _dafny.IntOfUint32((_844_v40).Cardinality()), (_this).F29(), p2)).Dtor_cf14(), (_this).F35()), 5)
			(_nw137).ArraySet1(_842_v38, 6)
			(_nw137).ArraySet1(Companion_D0_.Create_DC2_(Companion_Default___.Fm39(_824_v24, p2, globalState), _this.F30(), (func() bool {
				if (_845_v41).Contains((_this).F29()) {
					return (_845_v41).Get((_this).F29()).(bool)
				}
				return !(p0)
			})()), 7)
			(_nw137).ArraySet1(_842_v38, 8)
			(_nw137).ArraySet1(Companion_D0_.Create_DC2_(_dafny.MultiSetOf((_this).F35()), _this.F30(), p0), 9)
			(_nw137).ArraySet1(Companion_D0_.Create_DC2_(_841_v37, _this.F30(), (_this).F29()), 10)
			(_nw137).ArraySet1(_842_v38, 11)
			(_nw137).ArraySet1(Companion_D0_.Create_DC2_(_841_v37, _this.F30(), true), 12)
			(_nw137).ArraySet1(_842_v38, 13)
			(_nw137).ArraySet1(_842_v38, 14)
			_846_v42 = _nw137
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_846_v42), 0))
			_ = _index126
			(_846_v42).ArraySet1(Companion_D0_.Create_DC2_(_841_v37, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg76 _dafny.Int) interface{} {
					return coer76(arg76)
				}
			}((func(_847_v24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_848_i3 _dafny.Int) _dafny.CodePoint {
					return _847_v24
				}
			})(_824_v24))), p0), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_846_v42), 0))
			_ = _index127
			(_846_v42).ArraySet1(_842_v38, (_index127).Int())
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_843_v39), 0))
			_ = _index128
			(_843_v39).ArraySet1((_this).F29(), (_index128).Int())
			var _849_v43 _dafny.MultiSet
			_ = _849_v43
			_849_v43 = _dafny.MultiSetOf(_843_v39)
			var _850_v44 _dafny.Sequence
			_ = _850_v44
			_850_v44 = _dafny.SeqOf(_849_v43, _849_v43, (_849_v43).Update(_843_v39, Companion_Default___.Abs(Companion_Default___.Fm36(globalState))))
			var _851_v45 _dafny.Sequence
			_ = _851_v45
			_851_v45 = _dafny.SeqOf((_this).F36(), p2)
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_843_v39), 0))
			_ = _index129
			(_843_v39).ArraySet1(((_850_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_851_v45).Cardinality()), _dafny.IntOfUint32((_850_v44).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_dafny.MultiSetOf(_843_v39, _843_v39, _843_v39)), (_index129).Int())
			(globalState).F18 = !((_843_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_843_v39), 0))).Int()).(bool))
			(globalState).F21 = Companion_Default___.Fm9(globalState)
		}
		var _852_v46 _dafny.Set
		_ = _852_v46
		_852_v46 = _dafny.SetOf(_824_v24, _824_v24, _824_v24)
		_852_v46 = _852_v46
		var _853_v47 _dafny.Array
		_ = _853_v47
		var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw138
		_853_v47 = _nw138
		var _854_v48 _dafny.Set
		_ = _854_v48
		_854_v48 = _dafny.SetOf(_853_v47)
		var _855_v49 *C2
		_ = _855_v49
		var _nw139 *C2 = New_C2_()
		_ = _nw139
		_nw139.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("ivbrw"))).Cardinality()), _854_v48, (_this).F35(), _this.F30())
		_855_v49 = _nw139
		var _856_v50 _dafny.Set
		_ = _856_v50
		_856_v50 = _dafny.SetOf((_this).F35())
		var _857_v51 _dafny.Map
		_ = _857_v51
		_857_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F30()).Cardinality()), (_this).F29())
		var _hi2 _dafny.Int = (_857_v51).Cardinality()
		_ = _hi2
		for _858_i4 := ((_856_v50).Difference(_856_v50)).Cardinality(); _858_i4.Cmp(_hi2) < 0; _858_i4 = _858_i4.Plus(_dafny.One) {
			(globalState).F13 = !(p0) || (p0)
			(globalState).F13 = Companion_Default___.Fm19(_dafny.CodePoint('h'), globalState)
			var _859_v52 *C0
			_ = _859_v52
			var _nw140 *C0 = New_C0_()
			_ = _nw140
			_nw140.Ctor__((_this).F35())
			_859_v52 = _nw140
			var _860_v53 _dafny.MultiSet
			_ = _860_v53
			_860_v53 = _dafny.MultiSetOf(_dafny.IntOfInt64(162), (Companion_Default___.Fm39(_824_v24, _dafny.IntOfInt64(380), globalState)).Cardinality())
			var _861_v55 _dafny.Sequence
			_ = _861_v55
			_861_v55 = _dafny.SeqOf((_855_v49).F38())
			var _862_v56 _dafny.Sequence
			_ = _862_v56
			_862_v56 = _dafny.SeqOf(_861_v55, _861_v55)
			var _863_v57 _dafny.Array
			_ = _863_v57
			var _nwElement0_26 bool = (_this).F29()
			_ = _nwElement0_26
			var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(27))
			_ = _nw141
			(_nw141).ArraySet1(_nwElement0_26, 0)
			(_nw141).ArraySet1(p0, 1)
			(_nw141).ArraySet1((_this).F35(), 2)
			(_nw141).ArraySet1(Companion_Default___.Fm33(true, (_this).F36(), (func() _dafny.Int {
				if (_860_v53).Contains(((func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(777), _dafny.IntOfInt64(-289))); ; {
						_compr_38, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _864_v54 _dafny.Int
						_864_v54 = interface{}(_compr_38).(_dafny.Int)
						if ((_dafny.IntOfInt64(777)).Cmp(_864_v54) <= 0) && ((_864_v54).Cmp(_dafny.IntOfInt64(-289)) < 0) {
							_coll38.Add((_864_v54).Times((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality()), (_855_v49).F38())
						}
					}
					return _coll38.ToMap()
				}()).Update((_this).F36(), (_this).F36())).Cardinality()) {
					return (_860_v53).Multiplicity(((func() _dafny.Map {
						var _coll39 = _dafny.NewMapBuilder()
						_ = _coll39
						for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(777), _dafny.IntOfInt64(-289))); ; {
							_compr_39, _ok46 := _iter46()
							if !_ok46 {
								break
							}
							var _865_v54 _dafny.Int
							_865_v54 = interface{}(_compr_39).(_dafny.Int)
							if ((_dafny.IntOfInt64(777)).Cmp(_865_v54) <= 0) && ((_865_v54).Cmp(_dafny.IntOfInt64(-289)) < 0) {
								_coll39.Add((_865_v54).Times((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality()), (_855_v49).F38())
							}
						}
						return _coll39.ToMap()
					}()).Update((_this).F36(), (_this).F36())).Cardinality())
				}
				return (_855_v49).F38()
			})(), globalState), 3)
			(_nw141).ArraySet1(_859_v52.F37, 4)
			(_nw141).ArraySet1((_this).F35(), 5)
			(_nw141).ArraySet1(_859_v52.F37, 6)
			(_nw141).ArraySet1((_this).F29(), 7)
			(_nw141).ArraySet1((_859_v52).Fm11(true, globalState), 8)
			(_nw141).ArraySet1(((Companion_Default___.Fm41(_dafny.IntOfUint32((_862_v56).Cardinality()), (_this).F29(), (_this).F29(), _859_v52.F37, globalState)).Cardinality()).Cmp((_this).F36()) <= 0, 9)
			(_nw141).ArraySet1(p0, 10)
			(_nw141).ArraySet1(true, 11)
			(_nw141).ArraySet1(_dafny.Companion_Sequence_.Equal(_797_v0, _797_v0), 12)
			(_nw141).ArraySet1((_this).F29(), 13)
			(_nw141).ArraySet1(_859_v52.F37, 14)
			(_nw141).ArraySet1(p0, 15)
			(_nw141).ArraySet1((_this).F35(), 16)
			(_nw141).ArraySet1((_this).F29(), 17)
			(_nw141).ArraySet1((_dafny.IntOfInt64(777)).Cmp((func() _dafny.Int {
				if (_823_v23).Contains(_dafny.IntOfInt64(793)) {
					return (_823_v23).Get(_dafny.IntOfInt64(793)).(_dafny.Int)
				}
				return (_this).F36()
			})()) < 0, 18)
			(_nw141).ArraySet1(((_855_v49).F38()).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_859_v52, (_this).F29())).Cardinality()) <= 0, 19)
			(_nw141).ArraySet1(p0, 20)
			(_nw141).ArraySet1((_this).F35(), 21)
			(_nw141).ArraySet1((_797_v0).Select((Companion_Default___.SafeIndex((_855_v49).F38(), _dafny.IntOfUint32((_797_v0).Cardinality()))).Uint32()).(bool), 22)
			(_nw141).ArraySet1((_this).F29(), 23)
			(_nw141).ArraySet1((_this).F29(), 24)
			(_nw141).ArraySet1(!((_this).F35()), 25)
			(_nw141).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_824_v24, Companion_Default___.Fm42(_859_v52.F37, _859_v52.F37, globalState)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(688))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}((func(_866_v24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_867_i5 _dafny.Int) _dafny.CodePoint {
					return _866_v24
				}
			})(_824_v24)))), 26)
			_863_v57 = _nw141
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_863_v57), 0))
			_ = _index130
			(_863_v57).ArraySet1((_this).F35(), (_index130).Int())
			var _868_v58 _dafny.Array
			_ = _868_v58
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(29))
			_ = _nw142
			_868_v58 = _nw142
			var _869_v59 D1
			_ = _869_v59
			_869_v59 = Companion_D1_.Create_DC3_(_797_v0)
			var _870_v60 D1
			_ = _870_v60
			_870_v60 = Companion_D1_.Create_DC3_((_869_v59).Dtor_cf7())
			var _871_v61 D1
			_ = _871_v61
			_871_v61 = Companion_D1_.Create_DC5_(_870_v60)
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_868_v58), 0))
			_ = _index131
			(_868_v58).ArraySet1(_871_v61, (_index131).Int())
			var _872_v62 D10
			_ = _872_v62
			_872_v62 = Companion_D10_.Create_DC33_(false, p2)
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_863_v57), 0))
			_ = _index132
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_868_v58), 0))
			_ = _index133
			var _rhs138 bool = (_859_v52).Fm11((_872_v62).Dtor_cf54(), globalState)
			_ = _rhs138
			var _rhs139 D1 = _871_v61
			_ = _rhs139
			var _lhs127 _dafny.Array = _863_v57
			_ = _lhs127
			var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_863_v57), 0))
			_ = _lhs128
			var _lhs129 _dafny.Array = _868_v58
			_ = _lhs129
			var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_868_v58), 0))
			_ = _lhs130
			(_lhs127).ArraySet1(_rhs138, (_lhs128).Int())
			(_lhs129).ArraySet1(_rhs139, (_lhs130).Int())
		}
		var _873_v63 _dafny.Sequence
		_ = _873_v63
		_873_v63 = _dafny.SeqOf(_this.F30())
		r0 = (_873_v63).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfUint32((_873_v63).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _874_v64 _dafny.MultiSet
		_ = _874_v64
		_874_v64 = _dafny.MultiSetOf((_dafny.Zero).Minus(p2), (_this).F36())
		r1 = (Companion_Default___.SafeDivisionInt((_855_v49).F38(), (_874_v64).Cardinality())).Cmp(p2) > 0
		return r0, r1
	}
}
func (_this *C3) F35() bool {
	{
		return _this._f35
	}
}
func (_this *C3) F36() _dafny.Int {
	{
		return _this._f36
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f33 bool
	_f34 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f33 = false
	_this._f34 = _dafny.Zero
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

func (_this *C4) Ctor__(f33 bool, f34 _dafny.Int) {
	{
		(_this)._f33 = f33
		(_this)._f34 = f34
	}
}
func (_this *C4) Fm5(p0 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F34(), (_dafny.Zero).Minus((_this).F34()))
	}
}
func (_this *C4) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _875_i0 _dafny.Int
		_ = _875_i0
		_875_i0 = _dafny.Zero
		{
			for Companion_Default___.Fm6(((_this).F34()).Times((_this).F34()), globalState) {
				{
					if (_875_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_875_i0 = (_875_i0).Plus(_dafny.One)
					var _876_v0 _dafny.Map
					_ = _876_v0
					_876_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F33())
					_876_v0 = (_876_v0).Update((_this).F34(), (_this).F33())
					if (_this).F33() {
						(globalState).F28 = p0
						var _877_v1 _dafny.Set
						_ = _877_v1
						_877_v1 = _dafny.SetOf(p0)
						(globalState).F2 = (_this).Fm5((_877_v1).Union(_dafny.SetOf(Companion_Default___.Fm6((_this).F34(), globalState))), globalState)
						(globalState).F18 = !((_this).F33())
						var _878_v2 _dafny.Array
						_ = _878_v2
						var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
						_ = _nw143
						_878_v2 = _nw143
						(globalState).F17 = _878_v2
						(globalState).F21 = (((_this).F34()).Plus((_this).F34())).Times((_this).F34())
					} else {
						var _879_v3 _dafny.Sequence
						_ = _879_v3
						_879_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yk")
						(globalState).F7 = _dafny.SeqOf(((_this).F34()).Times(_dafny.IntOfUint32((_879_v3).Cardinality())))
						(globalState).F28 = !(p0) || (p0)
						var _880_v4 _dafny.Array
						_ = _880_v4
						var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
						_ = _nw144
						_880_v4 = _nw144
						var _881_v5 _dafny.Array
						_ = _881_v5
						var _len0_20 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_20
						var _nw145 _dafny.Array
						_ = _nw145
						if _len0_20.Cmp(_dafny.Zero) == 0 {
							_nw145 = _dafny.NewArray(_len0_20)
						} else {
							var _init20 func(_dafny.Int) _dafny.Sequence = (func(_882_p0 bool) func(_dafny.Int) _dafny.Sequence {
								return func(_883_i1 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf((_this).F33(), _882_p0)
								}
							})(p0)
							_ = _init20
							var _element0_20 = _init20(_dafny.Zero)
							_ = _element0_20
							_nw145 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
							(_nw145).ArraySet1(_element0_20, 0)
							var _nativeLen0_20 = (_len0_20).Int()
							_ = _nativeLen0_20
							for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
								(_nw145).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
							}
						}
						_881_v5 = _nw145
						var _884_v6 _dafny.Sequence
						_ = _884_v6
						_884_v6 = _dafny.SeqOf(_881_v5)
						var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_880_v4), 0))
						_ = _index134
						(_880_v4).ArraySet1((_884_v6).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_884_v6).Cardinality()))).Uint32()).(_dafny.Array), (_index134).Int())
						var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_880_v4), 0))
						_ = _index135
						(_880_v4).ArraySet1(_881_v5, (_index135).Int())
						var _885_v7 _dafny.Array
						_ = _885_v7
						var _nw146 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw146
						_885_v7 = _nw146
						_885_v7 = _885_v7
						(globalState).F15 = p0
					}
					var _886_v8 _dafny.Sequence
					_ = _886_v8
					_886_v8 = _dafny.UnicodeSeqOfUtf8Bytes("ue")
					var _887_v9 _dafny.Sequence
					_ = _887_v9
					_887_v9 = _dafny.SeqOf(_886_v8)
					(globalState).F2 = _dafny.IntOfUint32((_887_v9).Cardinality())
					var _888_v10 _dafny.Sequence
					_ = _888_v10
					_888_v10 = _dafny.SeqOf((_this).F33(), false, (_this).F33(), false, p0)
					(globalState).F15 = (func() bool {
						if false {
							return !_dafny.Companion_Sequence_.Contains(_888_v10, (_this).F33())
						}
						return (_this).F33()
					})()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _hi3 _dafny.Int = (_this).F34()
		_ = _hi3
		for _889_i2 := _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fwevwbc")).Cardinality()); _889_i2.Cmp(_hi3) < 0; _889_i2 = _889_i2.Plus(_dafny.One) {
			(globalState).F15 = !(Companion_Default___.Fm6(_889_i2, globalState))
			var _890_v11 _dafny.Map
			_ = _890_v11
			_890_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(748), (_this).F34())
			_890_v11 = (_890_v11).Update(_dafny.IntOfInt64(-437), (_this).F34())
			var _891_v12 _dafny.Array
			_ = _891_v12
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_21
			var _nw147 _dafny.Array
			_ = _nw147
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw147 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) bool = func(_892_i3 _dafny.Int) bool {
					return (_dafny.IntOfInt64(352)).Cmp((_dafny.Zero).Minus((_this).F34())) != 0
				}
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw147 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw147).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw147).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_891_v12 = _nw147
			_891_v12 = _891_v12
			var _893_v13 _dafny.Sequence
			_ = _893_v13
			_893_v13 = _dafny.SeqOf(_891_v12, _891_v12)
			var _894_v14 _dafny.Array
			_ = _894_v14
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw148
			_894_v14 = _nw148
			var _895_v15 _dafny.Array
			_ = _895_v15
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_22
			var _nw149 _dafny.Array
			_ = _nw149
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw149 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = func(_896_i4 _dafny.Int) _dafny.Int {
					return (_896_i4).Minus(_dafny.IntOfInt64(537))
				}
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw149 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw149).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw149).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_895_v15 = _nw149
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_894_v14), 0))
			_ = _index136
			(_894_v14).ArraySet1(_895_v15, (_index136).Int())
			var _897_v16 _dafny.Sequence
			_ = _897_v16
			_897_v16 = _dafny.UnicodeSeqOfUtf8Bytes("hvp")
			var _898_v17 _dafny.MultiSet
			_ = _898_v17
			_898_v17 = _dafny.MultiSetOf(p0)
			var _899_v18 _dafny.Sequence
			_ = _899_v18
			_899_v18 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jkowf"), Companion_Default___.Fm7((_this).F34(), globalState))
			var _900_v19 _dafny.CodePoint
			_ = _900_v19
			_900_v19 = _dafny.CodePoint('j')
			var _901_v20 _dafny.Map
			_ = _901_v20
			_901_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_i2, _900_v19)
			var _pat_let_tv20 = globalState
			_ = _pat_let_tv20
			var _902_v21 D2
			_ = _902_v21
			_902_v21 = Companion_D2_.Create_DC9_(_897_v16, ((_890_v11).Update(_dafny.IntOfUint32((_897_v16).Cardinality()), _889_i2)).Cardinality(), (func(_pat_let18_0 D0) D0 {
				return func(_903_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let19_0 _dafny.Sequence) D0 {
						return func(_904_dt__update_hcf5_h0 _dafny.Sequence) D0 {
							return func(_pat_let20_0 bool) D0 {
								return func(_905_dt__update_hcf6_h0 bool) D0 {
									return Companion_D0_.Create_DC2_((_903_dt__update__tmp_h0).Dtor_cf4(), _904_dt__update_hcf5_h0, _905_dt__update_hcf6_h0)
								}(_pat_let20_0)
							}(true)
						}(_pat_let19_0)
					}(Companion_Default___.Fm7(_889_i2, _pat_let_tv20))
				}(_pat_let18_0)
			}(Companion_D0_.Create_DC2_(_898_v17, (_899_v18).Select((Companion_Default___.SafeIndex(((_901_v20).Update(_889_i2, _900_v19)).Cardinality(), _dafny.IntOfUint32((_899_v18).Cardinality()))).Uint32()).(_dafny.Sequence), p0))).Dtor_cf6(), (_this).F34())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_894_v14), 0))
			_ = _index137
			var _rhs140 _dafny.Sequence = _dafny.SeqOf((func() _dafny.Array {
				if p0 {
					return _891_v12
				}
				return _891_v12
			})(), _891_v12, _891_v12, _891_v12)
			_ = _rhs140
			var _rhs141 _dafny.Array = _895_v15
			_ = _rhs141
			var _rhs142 _dafny.Int = (_902_v21).Dtor_cf17()
			_ = _rhs142
			var _lhs131 _dafny.Array = _894_v14
			_ = _lhs131
			var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_894_v14), 0))
			_ = _lhs132
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			_893_v13 = _rhs140
			(_lhs131).ArraySet1(_rhs141, (_lhs132).Int())
			_lhs133.F2 = _rhs142
		}
		var _906_v22 _dafny.Array
		_ = _906_v22
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_23
		var _nw150 _dafny.Array
		_ = _nw150
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw150 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) _dafny.Int = func(_907_i5 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_907_i5, (_this).F34())
			}
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw150 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw150).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw150).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_906_v22 = _nw150
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))
		_ = _index138
		(_906_v22).ArraySet1((_this).F34(), (_index138).Int())
		var _908_v23 _dafny.Sequence
		_ = _908_v23
		_908_v23 = _dafny.SeqOf(true)
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))
		_ = _index139
		var _rhs143 _dafny.Int = (_this).F34()
		_ = _rhs143
		var _rhs144 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if p0 {
				return _dafny.SeqOf((_this).F33())
			}
			return _dafny.SeqOf((_this).F33())
		})(), _908_v23)
		_ = _rhs144
		var _lhs134 _dafny.Array = _906_v22
		_ = _lhs134
		var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))
		_ = _lhs135
		(_lhs134).ArraySet1(_rhs143, (_lhs135).Int())
		_908_v23 = _rhs144
		var _909_v24 _dafny.MultiSet
		_ = _909_v24
		_909_v24 = _dafny.MultiSetOf((_906_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))).Int()).(_dafny.Int), (_906_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))).Int()).(_dafny.Int), (_this).F34())
		var _910_v25 D3
		_ = _910_v25
		_910_v25 = Companion_D3_.Create_DC10_(_909_v24)
		(globalState).F28 = (_909_v24).Equals((_910_v25).Dtor_cf18())
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))
		_ = _index140
		(_906_v22).ArraySet1((_this).F34(), (_index140).Int())
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_24
		var _nw151 _dafny.Array
		_ = _nw151
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw151 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) _dafny.Int = func(_911_i6 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_911_i6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bmap")).Cardinality()))
			}
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw151 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw151).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw151).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_906_v22 = _nw151
		var _912_v26 _dafny.Sequence
		_ = _912_v26
		_912_v26 = _dafny.UnicodeSeqOfUtf8Bytes("hoalen")
		var _913_v27 _dafny.Sequence
		_ = _913_v27
		_913_v27 = _dafny.SeqOf(_908_v23, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, (_this).F33(), p0, p0, Companion_Default___.Fm6(_dafny.IntOfUint32((_912_v26).Cardinality()), globalState)), _908_v23))
		r0 = _913_v27
		r1 = !((Companion_Default___.SafeDivisionInt((_906_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))).Int()).(_dafny.Int), (_906_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))).Int()).(_dafny.Int))).Cmp((_906_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_906_v22), 0))).Int()).(_dafny.Int)) <= 0)
		return r0, r1
	}
}
func (_this *C4) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _914_v0 _dafny.Array
		_ = _914_v0
		var _nw152 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw152
		_914_v0 = _nw152
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))
		_ = _index141
		(_914_v0).ArraySet1((_this).F33(), (_index141).Int())
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))
		_ = _index142
		(_914_v0).ArraySet1(!((_this).F33()) || ((_this).F33()), (_index142).Int())
		(globalState).F28 = (_this).F33()
		var _915_v1 _dafny.Sequence
		_ = _915_v1
		_915_v1 = _dafny.UnicodeSeqOfUtf8Bytes("fckyvw")
		var _916_v2 D2
		_ = _916_v2
		_916_v2 = Companion_D2_.Create_DC9_(_915_v1, p2, p0, (_this).F34())
		r0 = _dafny.Companion_Sequence_.Update(_915_v1, (Companion_Default___.SafeIndex((_916_v2).Dtor_cf15(), _dafny.IntOfUint32((_915_v1).Cardinality()))).Uint32(), (_915_v1).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_915_v1).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _917_v3 _dafny.Set
		_ = _917_v3
		_917_v3 = _dafny.SetOf(false, (_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool))
		_916_v2 = Companion_Default___.Fm8(!(((_this).F34()).Cmp((_this).Fm5(_917_v3, globalState)) > 0), (_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool), globalState)
		var _918_v4 _dafny.CodePoint
		_ = _918_v4
		_918_v4 = _dafny.CodePoint('u')
		var _919_v5 _dafny.Array
		_ = _919_v5
		var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw153
		_919_v5 = _nw153
		var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_919_v5), 0))
		_ = _index143
		(_919_v5).ArraySet1(p2, (_index143).Int())
		var _920_v6 _dafny.Set
		_ = _920_v6
		_920_v6 = _dafny.SetOf(p2)
		var _921_v7 _dafny.MultiSet
		_ = _921_v7
		_921_v7 = _dafny.MultiSetOf(p2, (_this).F34(), (func() _dafny.Int {
			if p0 {
				return (_this).F34()
			}
			return (_this).F34()
		})(), (_920_v6).Cardinality(), _dafny.IntOfInt64(223))
		var _922_v8 _dafny.Sequence
		_ = _922_v8
		_922_v8 = _dafny.SeqOf((_this).F33(), p0, (_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool))
		var _923_v9 _dafny.Map
		_ = _923_v9
		_923_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _924_v10 _dafny.MultiSet
		_ = _924_v10
		_924_v10 = _dafny.MultiSetOf(_922_v8)
		var _925_v11 _dafny.Sequence
		_ = _925_v11
		_925_v11 = _dafny.SeqOf((_this).F34(), (_924_v10).Cardinality(), p2)
		var _926_v12 _dafny.Map
		_ = _926_v12
		_926_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool), _dafny.SeqOf((_dafny.Zero).Minus(p2)))
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_919_v5), 0))
		_ = _index144
		var _rhs145 _dafny.Int = (func() _dafny.Int {
			if (_921_v7).Contains((_this).Fm5(_dafny.SetOf((_922_v8).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_922_v8).Cardinality()))).Uint32()).(bool)), globalState)) {
				return (_921_v7).Multiplicity((_this).Fm5(_dafny.SetOf((_922_v8).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_922_v8).Cardinality()))).Uint32()).(bool)), globalState))
			}
			return (_923_v9).Cardinality()
		})()
		_ = _rhs145
		var _rhs146 bool = !((_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool))
		_ = _rhs146
		var _rhs147 _dafny.CodePoint = _918_v4
		_ = _rhs147
		var _rhs148 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F34(), (_this).F34(), p2, _dafny.IntOfUint32((_922_v8).Cardinality()), (_this).F34()), _925_v11), (func() _dafny.Sequence {
			if (_926_v12).Contains((_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool)) {
				return (_926_v12).Get((_914_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_914_v0), 0))).Int()).(bool)).(_dafny.Sequence)
			}
			return _925_v11
		})())
		_ = _rhs148
		var _rhs149 _dafny.Int = p2
		_ = _rhs149
		var _lhs136 *GlobalState = globalState
		_ = _lhs136
		var _lhs137 *GlobalState = globalState
		_ = _lhs137
		var _lhs138 *GlobalState = globalState
		_ = _lhs138
		var _lhs139 _dafny.Array = _919_v5
		_ = _lhs139
		var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_919_v5), 0))
		_ = _lhs140
		_lhs136.F21 = _rhs145
		_lhs137.F28 = _rhs146
		_918_v4 = _rhs147
		_lhs138.F7 = _rhs148
		(_lhs139).ArraySet1(_rhs149, (_lhs140).Int())
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((p1), 0))
		_ = _index145
		(p1).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-838))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg78 _dafny.Int) interface{} {
				return coer78(arg78)
			}
		}((func(_927_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_928_i0 _dafny.Int) _dafny.CodePoint {
				return (Companion_D2_.Create_DC8_(_927_v4, (_this).F33())).Dtor_cf12()
			}
		})(_918_v4))), (_index145).Int())
		var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((p1), 0))
		_ = _index146
		(p1).ArraySet1(_915_v1, (_index146).Int())
		r0 = _dafny.Companion_Sequence_.Concatenate(_915_v1, _dafny.UnicodeSeqOfUtf8Bytes("keatndrhj"))
		r1 = Companion_Default___.Fm6(_dafny.IntOfInt64(82), globalState)
		return r0, r1
	}
}
func (_this *C4) M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 D0, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _929_v0 _dafny.Array
		_ = _929_v0
		var _nwElement0_27 bool = p0
		_ = _nwElement0_27
		var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(29))
		_ = _nw154
		(_nw154).ArraySet1(_nwElement0_27, 0)
		(_nw154).ArraySet1(p0, 1)
		(_nw154).ArraySet1((_this).F33(), 2)
		(_nw154).ArraySet1(false, 3)
		(_nw154).ArraySet1(p0, 4)
		(_nw154).ArraySet1(p0, 5)
		(_nw154).ArraySet1(!(p0), 6)
		(_nw154).ArraySet1((_this).F33(), 7)
		(_nw154).ArraySet1(!(true), 8)
		(_nw154).ArraySet1((_this).F33(), 9)
		(_nw154).ArraySet1((_this).F33(), 10)
		(_nw154).ArraySet1(p0, 11)
		(_nw154).ArraySet1((_this).F33(), 12)
		(_nw154).ArraySet1((_this).F33(), 13)
		(_nw154).ArraySet1((_this).F33(), 14)
		(_nw154).ArraySet1(true, 15)
		(_nw154).ArraySet1((_this).F33(), 16)
		(_nw154).ArraySet1(p0, 17)
		(_nw154).ArraySet1((_this).F33(), 18)
		(_nw154).ArraySet1(p0, 19)
		(_nw154).ArraySet1((_this).F33(), 20)
		(_nw154).ArraySet1(p0, 21)
		(_nw154).ArraySet1((_this).F33(), 22)
		(_nw154).ArraySet1(p0, 23)
		(_nw154).ArraySet1(true, 24)
		(_nw154).ArraySet1(true, 25)
		(_nw154).ArraySet1(p0, 26)
		(_nw154).ArraySet1((_this).F33(), 27)
		(_nw154).ArraySet1(p0, 28)
		_929_v0 = _nw154
		var _930_v1 _dafny.Sequence
		_ = _930_v1
		_930_v1 = _dafny.SeqOf(_929_v0)
		var _931_v2 _dafny.MultiSet
		_ = _931_v2
		_931_v2 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_930_v1).Cardinality()), (_this).F34()))
		var _932_v3 _dafny.Set
		_ = _932_v3
		_932_v3 = _dafny.SetOf((_this).F33())
		var _933_v4 _dafny.Sequence
		_ = _933_v4
		_933_v4 = _dafny.SeqOf(p0)
		var _rhs150 _dafny.Int = (_this).Fm5(_932_v3, globalState)
		_ = _rhs150
		var _rhs151 bool = (_933_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.IntOfUint32((_933_v4).Cardinality()))).Uint32()).(bool)
		_ = _rhs151
		var _rhs152 _dafny.Int = _dafny.IntOfUint32((p2).Cardinality())
		_ = _rhs152
		var _rhs153 bool = !((_this).F33())
		_ = _rhs153
		var _rhs154 _dafny.MultiSet = _931_v2
		_ = _rhs154
		var _lhs141 *GlobalState = globalState
		_ = _lhs141
		var _lhs142 *GlobalState = globalState
		_ = _lhs142
		var _lhs143 *GlobalState = globalState
		_ = _lhs143
		_lhs141.F2 = _rhs150
		r0 = _rhs151
		_lhs142.F21 = _rhs152
		_lhs143.F13 = _rhs153
		_931_v2 = _rhs154
		var _934_v5 _dafny.Sequence
		_ = _934_v5
		_934_v5 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_933_v4, (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_933_v4).Cardinality()))).Uint32(), (_this).F33()), _933_v4)
		var _935_v6 _dafny.Sequence
		_ = _935_v6
		_935_v6 = _dafny.SeqOf((_934_v5).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_934_v5).Cardinality()))).Uint32()).(_dafny.Sequence))
		(globalState).F2 = (func() _dafny.Int {
			if Companion_Default___.Fm6((_this).F34(), globalState) {
				return (_dafny.IntOfUint32(((_934_v5).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_934_v5).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Minus((_this).F34())
			}
			return (_this).Fm5(_932_v3, globalState)
		})()
		var _936_i0 _dafny.Int
		_ = _936_i0
		_936_i0 = _dafny.Zero
		{
			for !(p0) {
				{
					if (_936_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_936_i0 = (_936_i0).Plus(_dafny.One)
					(globalState).F28 = ((_this).F34()).Cmp((_this).F34()) <= 0
					var _937_v7 _dafny.Array
					_ = _937_v7
					var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
					_ = _nw155
					_937_v7 = _nw155
					var _938_v8 _dafny.CodePoint
					_ = _938_v8
					_938_v8 = _dafny.CodePoint('k')
					var _939_v9 _dafny.Map
					_ = _939_v9
					_939_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v8, (_933_v4).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_933_v4).Cardinality()))).Uint32()).(bool))
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_937_v7), 0))
					_ = _index147
					(_937_v7).ArraySet1(_939_v9, (_index147).Int())
					var _940_v11 _dafny.Map
					_ = _940_v11
					_940_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (_dafny.Zero).Minus((_this).F34()))
					var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_937_v7), 0))
					_ = _index148
					var _rhs155 _dafny.Map = func() _dafny.Map {
						var _coll40 = _dafny.NewMapBuilder()
						_ = _coll40
						for _iter47 := _dafny.Iterate((_940_v11).Keys().Elements()); ; {
							_compr_40, _ok47 := _iter47()
							if !_ok47 {
								break
							}
							var _941_v10 _dafny.CodePoint
							_941_v10 = interface{}(_compr_40).(_dafny.CodePoint)
							if (_940_v11).Contains(_941_v10) {
								_coll40.Add(_941_v10, false)
							}
						}
						return _coll40.ToMap()
					}()
					_ = _rhs155
					var _rhs156 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F34(), (_this).F34()))
					_ = _rhs156
					var _rhs157 _dafny.Int = (_this).F34()
					_ = _rhs157
					var _lhs144 _dafny.Array = _937_v7
					_ = _lhs144
					var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_937_v7), 0))
					_ = _lhs145
					var _lhs146 *GlobalState = globalState
					_ = _lhs146
					var _lhs147 *GlobalState = globalState
					_ = _lhs147
					(_lhs144).ArraySet1(_rhs155, (_lhs145).Int())
					_lhs146.F2 = _rhs156
					_lhs147.F14 = _rhs157
					(globalState).F2 = ((_this).F34()).Plus((_this).F34())
					var _942_v12 _dafny.Map
					_ = _942_v12
					_942_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F34(), (_this).F34()), p0)
					_942_v12 = (_942_v12).Update((_this).F34(), p0)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (_this).F33()
		var _943_v13 _dafny.Set
		_ = _943_v13
		_943_v13 = _dafny.SetOf((_this).F34(), (_this).F34())
		(globalState).F14 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_this).F34(), (_943_v13).Cardinality()), (_this).F34())
		var _hi4 _dafny.Int = (_this).F34()
		_ = _hi4
		for _944_i1 := (_this).F34(); _944_i1.Cmp(_hi4) < 0; _944_i1 = _944_i1.Plus(_dafny.One) {
			var _945_v14 _dafny.MultiSet
			_ = _945_v14
			_945_v14 = _dafny.MultiSetOf(Companion_Default___.Fm6((_this).F34(), globalState))
			var _946_v15 _dafny.Sequence
			_ = _946_v15
			_946_v15 = _dafny.SeqOf(_944_i1)
			var _947_v16 _dafny.CodePoint
			_ = _947_v16
			_947_v16 = _dafny.CodePoint('p')
			var _948_v17 _dafny.Map
			_ = _948_v17
			_948_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_944_i1, _947_v16)
			var _949_v18 _dafny.Array
			_ = _949_v18
			var _nwElement0_28 _dafny.Int = _944_i1
			_ = _nwElement0_28
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(11))
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_28, 0)
			(_nw156).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F34(), (_this).F34(), _944_i1, _944_i1, (_dafny.Zero).Minus((_dafny.Zero).Minus(_944_i1))), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_945_v14).Contains(true) {
					return (_945_v14).Multiplicity(true)
				}
				return _944_i1
			})(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F34(), (_this).F34(), _944_i1, _944_i1, (_dafny.Zero).Minus((_dafny.Zero).Minus(_944_i1)))).Cardinality()))).Uint32(), _944_i1), _946_v15)).Cardinality()), 1)
			(_nw156).ArraySet1(Companion_Default___.SafeDivisionInt(_944_i1, _944_i1), 2)
			(_nw156).ArraySet1(_944_i1, 3)
			(_nw156).ArraySet1((_this).F34(), 4)
			(_nw156).ArraySet1(_944_i1, 5)
			(_nw156).ArraySet1(_dafny.IntOfInt64(305), 6)
			(_nw156).ArraySet1(_944_i1, 7)
			(_nw156).ArraySet1(_944_i1, 8)
			(_nw156).ArraySet1((_dafny.Zero).Minus((_948_v17).Cardinality()), 9)
			(_nw156).ArraySet1(_944_i1, 10)
			_949_v18 = _nw156
			(globalState).F17 = _949_v18
			(globalState).F15 = !((_this).F33())
			var _950_v19 *C1
			_ = _950_v19
			var _nw157 *C1 = New_C1_()
			_ = _nw157
			_nw157.Ctor__(p0)
			_950_v19 = _nw157
			if (_950_v19).F40() {
				(globalState).F17 = _949_v18
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))
				_ = _index149
				(_929_v0).ArraySet1(!((_this).F33()), (_index149).Int())
				var _951_v20 _dafny.Map
				_ = _951_v20
				_951_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_944_i1, (_this).F34())
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))
				_ = _index150
				var _rhs158 bool = (Companion_Default___.Fm6((_951_v20).Cardinality(), globalState)) == (true)
				_ = _rhs158
				var _rhs159 bool = (_932_v3).IsDisjointFrom((func() _dafny.Set {
					if !((_this).F33()) {
						return _932_v3
					}
					return _932_v3
				})())
				_ = _rhs159
				var _rhs160 bool = ((_this).F34()).Cmp((_this).F34()) != 0
				_ = _rhs160
				var _lhs148 *GlobalState = globalState
				_ = _lhs148
				var _lhs149 _dafny.Array = _929_v0
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))
				_ = _lhs150
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				_lhs148.F13 = _rhs158
				(_lhs149).ArraySet1(_rhs159, (_lhs150).Int())
				_lhs151.F13 = _rhs160
				var _952_v21 _dafny.Array
				_ = _952_v21
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw158
				_952_v21 = _nw158
				var _953_v22 _dafny.Map
				_ = _953_v22
				_953_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(21), Companion_Default___.Fm19(_947_v16, globalState))
				var _954_v23 _dafny.Map
				_ = _954_v23
				_954_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_950_v19).F40()), p0)
				var _955_v24 _dafny.Array
				_ = _955_v24
				var _nwElement0_29 bool = (_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool)
				_ = _nwElement0_29
				var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(28))
				_ = _nw159
				(_nw159).ArraySet1(_nwElement0_29, 0)
				(_nw159).ArraySet1(true, 1)
				(_nw159).ArraySet1((_933_v4).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_933_v4).Cardinality()))).Uint32()).(bool), 2)
				(_nw159).ArraySet1((p3).Dtor_cf0(), 3)
				(_nw159).ArraySet1(true, 4)
				(_nw159).ArraySet1((_950_v19).F40(), 5)
				(_nw159).ArraySet1((_950_v19).F40(), 6)
				(_nw159).ArraySet1(p0, 7)
				(_nw159).ArraySet1((_950_v19).F40(), 8)
				(_nw159).ArraySet1((_this).F33(), 9)
				(_nw159).ArraySet1((func() bool {
					if (_953_v22).Contains((_this).F34()) {
						return (_953_v22).Get((_this).F34()).(bool)
					}
					return (_this).F33()
				})(), 10)
				(_nw159).ArraySet1((_this).F33(), 11)
				(_nw159).ArraySet1((_this).F33(), 12)
				(_nw159).ArraySet1((_950_v19).F40(), 13)
				(_nw159).ArraySet1((func() bool {
					if (_954_v23).Contains((_this).F33()) {
						return (_954_v23).Get((_this).F33()).(bool)
					}
					return (_950_v19).F40()
				})(), 14)
				(_nw159).ArraySet1((_950_v19).F40(), 15)
				(_nw159).ArraySet1((_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool), 16)
				(_nw159).ArraySet1((_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool), 17)
				(_nw159).ArraySet1(p0, 18)
				(_nw159).ArraySet1((_950_v19).F40(), 19)
				(_nw159).ArraySet1(p0, 20)
				(_nw159).ArraySet1((_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool), 21)
				(_nw159).ArraySet1((_this).F33(), 22)
				(_nw159).ArraySet1((_950_v19).F40(), 23)
				(_nw159).ArraySet1((_950_v19).F40(), 24)
				(_nw159).ArraySet1(p0, 25)
				(_nw159).ArraySet1((_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool), 26)
				(_nw159).ArraySet1((_929_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_929_v0), 0))).Int()).(bool), 27)
				_955_v24 = _nw159
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_952_v21), 0))
				_ = _index151
				(_952_v21).ArraySet1(_955_v24, (_index151).Int())
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_952_v21), 0))
				_ = _index152
				(_952_v21).ArraySet1(_955_v24, (_index152).Int())
				_954_v23 = (_954_v23).Update((_dafny.SetOf(_944_i1)).Equals(_943_v13), true)
				var _956_v25 _dafny.Map
				_ = _956_v25
				_956_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_949_v18, _944_i1)
				_956_v25 = (_956_v25).Update(_949_v18, (_this).F34())
			} else {
				_948_v17 = (_948_v17).Update(_dafny.IntOfInt64(-901), _dafny.CodePoint('y'))
				(globalState).F14 = _dafny.IntOfInt64(566)
				(globalState).F28 = (_950_v19).F40()
				(globalState).F2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_933_v4, _933_v4)).Cardinality())
				var _957_v26 _dafny.Array
				_ = _957_v26
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_25
				var _nw160 _dafny.Array
				_ = _nw160
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw160 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) D8 = func(_958_i2 _dafny.Int) D8 {
						return Companion_D8_.Create_DC26_()
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw160 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw160).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw160).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_957_v26 = _nw160
				var _959_v27 D8
				_ = _959_v27
				_959_v27 = Companion_D8_.Create_DC26_()
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_957_v26), 0))
				_ = _index153
				(_957_v26).ArraySet1((func() D8 {
					if (_950_v19).F40() {
						return _959_v27
					}
					return Companion_D8_.Create_DC26_()
				})(), (_index153).Int())
				var _960_v28 D2
				_ = _960_v28
				_960_v28 = Companion_D2_.Create_DC8_(_947_v16, (_this).F33())
				var _961_v29 _dafny.Sequence
				_ = _961_v29
				_961_v29 = _dafny.SeqOf(_960_v28)
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_957_v26), 0))
				_ = _index154
				var _rhs161 bool = (((_dafny.Zero).Minus((_931_v2).Cardinality())).Minus(_dafny.IntOfUint32((Companion_Default___.Fm35(_961_v29, _944_i1, globalState)).Cardinality()))).Cmp(_944_i1) <= 0
				_ = _rhs161
				var _rhs162 D8 = _959_v27
				_ = _rhs162
				var _rhs163 bool = p0
				_ = _rhs163
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				var _lhs153 _dafny.Array = _957_v26
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_957_v26), 0))
				_ = _lhs154
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				_lhs152.F15 = _rhs161
				(_lhs153).ArraySet1(_rhs162, (_lhs154).Int())
				_lhs155.F13 = _rhs163
			}
		}
		r0 = (func() bool {
			if p0 {
				return p0
			}
			return (_this).F33()
		})()
		r1 = (_this).F33()
		return r0, r1
	}
}
func (_this *C4) M7(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _962_v0 _dafny.Array
		_ = _962_v0
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_26
		var _nw161 _dafny.Array
		_ = _nw161
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw161 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) _dafny.Int = func(_963_i1 _dafny.Int) _dafny.Int {
				return (_963_i1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F33())).Cardinality())
			}
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw161 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw161).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw161).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_962_v0 = _nw161
		for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_962_v0), 0))); ; {
			_guard_loop_7, _ok48 := _iter48()
			if !_ok48 {
				break
			}
			var _964_i0 _dafny.Int
			_964_i0 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_964_i0).Sign() != -1) && ((_964_i0).Cmp(_dafny.ArrayLen((_962_v0), 0)) < 0)) {
				(_962_v0).ArraySet1(Companion_Default___.SafeModuloInt(_964_i0, p0), (_964_i0).Int())
			}
		}
		var _965_v1 _dafny.Sequence
		_ = _965_v1
		_965_v1 = _dafny.UnicodeSeqOfUtf8Bytes("aqds")
		var _966_v2 *C3
		_ = _966_v2
		var _nw162 *C3 = New_C3_()
		_ = _nw162
		_nw162.Ctor__((_this).F33(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), p0)).Cardinality(), ((_this).F34()).Cmp((_dafny.Zero).Minus((_this).F34())) != 0, _965_v1)
		_966_v2 = _nw162
		var _967_v3 _dafny.Sequence
		_ = _967_v3
		_967_v3 = _dafny.SeqOf((_this).F33())
		var _968_v4 D1
		_ = _968_v4
		_968_v4 = Companion_D1_.Create_DC3_(_967_v3)
		var _source15 D1 = (func() D1 {
			if (_966_v2).F35() {
				return Companion_D1_.Create_DC3_(_dafny.SeqOf((_this).F33()))
			}
			return _968_v4
		})()
		_ = _source15
		if _source15.Is_DC4() {
			var _969___mcc_h0 bool = _source15.Get_().(D1_DC4).Cf8
			_ = _969___mcc_h0
			var _970_cf8 bool = _969___mcc_h0
			_ = _970_cf8
			var _971_v5 _dafny.Array
			_ = _971_v5
			var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw163
			_971_v5 = _nw163
			var _972_v6 _dafny.CodePoint
			_ = _972_v6
			_972_v6 = _dafny.CodePoint('a')
			var _973_v7 _dafny.Array
			_ = _973_v7
			var _nwElement0_30 bool = !(false)
			_ = _nwElement0_30
			var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(9))
			_ = _nw164
			(_nw164).ArraySet1(_nwElement0_30, 0)
			(_nw164).ArraySet1((_966_v2).F35(), 1)
			(_nw164).ArraySet1((_966_v2).F35(), 2)
			(_nw164).ArraySet1(_970_cf8, 3)
			(_nw164).ArraySet1(false, 4)
			(_nw164).ArraySet1((_966_v2).F35(), 5)
			(_nw164).ArraySet1(Companion_Default___.Fm19(_972_v6, globalState), 6)
			(_nw164).ArraySet1((_966_v2).F35(), 7)
			(_nw164).ArraySet1(_970_cf8, 8)
			_973_v7 = _nw164
			var _974_v8 _dafny.Sequence
			_ = _974_v8
			_974_v8 = _dafny.SeqOf(_973_v7, _973_v7)
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_971_v5), 0))
			_ = _index155
			(_971_v5).ArraySet1((_974_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-173), _dafny.IntOfUint32((_974_v8).Cardinality()))).Uint32()).(_dafny.Array), (_index155).Int())
			var _975_v9 _dafny.Map
			_ = _975_v9
			_975_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-577), false)
			var _976_v10 D4
			_ = _976_v10
			_976_v10 = Companion_D4_.Create_DC14_((_975_v9).Cardinality())
			var _977_v11 _dafny.Array
			_ = _977_v11
			var _nwElement0_31 D4 = func(_pat_let21_0 D4) D4 {
				return func(_978_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let22_0 _dafny.Int) D4 {
						return func(_979_dt__update_hcf23_h0 _dafny.Int) D4 {
							return Companion_D4_.Create_DC14_(_979_dt__update_hcf23_h0)
						}(_pat_let22_0)
					}((_dafny.Zero).Minus((_this).F34()))
				}(_pat_let21_0)
			}(_976_v10)
			_ = _nwElement0_31
			var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
			_ = _nw165
			(_nw165).ArraySet1(_nwElement0_31, 0)
			(_nw165).ArraySet1(Companion_D4_.Create_DC14_((_966_v2).F36()), 1)
			(_nw165).ArraySet1(_976_v10, 2)
			(_nw165).ArraySet1(Companion_D4_.Create_DC14_(_dafny.IntOfUint32((Companion_Default___.Fm20((_this).F33(), (_966_v2).F35(), _967_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(565))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg79 _dafny.Int) interface{} {
					return coer79(arg79)
				}
			}((func(_980_v2 *C3) func(_dafny.Int) _dafny.Int {
				return func(_981_i2 _dafny.Int) _dafny.Int {
					return (_980_v2).F36()
				}
			})(_966_v2)))).Cardinality()), globalState)).Cardinality())), 3)
			(_nw165).ArraySet1(_976_v10, 4)
			(_nw165).ArraySet1(_976_v10, 5)
			(_nw165).ArraySet1(_976_v10, 6)
			(_nw165).ArraySet1(_976_v10, 7)
			_977_v11 = _nw165
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_977_v11), 0))
			_ = _index156
			(_977_v11).ArraySet1(_976_v10, (_index156).Int())
			var _pat_let_tv21 = _966_v2
			_ = _pat_let_tv21
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_971_v5), 0))
			_ = _index157
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_977_v11), 0))
			_ = _index158
			var _rhs164 _dafny.Array = _973_v7
			_ = _rhs164
			var _rhs165 D4 = func(_pat_let23_0 D4) D4 {
				return func(_982_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let24_0 _dafny.Int) D4 {
						return func(_983_dt__update_hcf23_h1 _dafny.Int) D4 {
							return Companion_D4_.Create_DC14_(_983_dt__update_hcf23_h1)
						}(_pat_let24_0)
					}((_pat_let_tv21).F36())
				}(_pat_let23_0)
			}(_976_v10)
			_ = _rhs165
			var _rhs166 bool = _970_cf8
			_ = _rhs166
			var _lhs156 _dafny.Array = _971_v5
			_ = _lhs156
			var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_971_v5), 0))
			_ = _lhs157
			var _lhs158 _dafny.Array = _977_v11
			_ = _lhs158
			var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_977_v11), 0))
			_ = _lhs159
			(_lhs156).ArraySet1(_rhs164, (_lhs157).Int())
			(_lhs158).ArraySet1(_rhs165, (_lhs159).Int())
			r0 = _rhs166
			var _984_v12 _dafny.Map
			_ = _984_v12
			_984_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F33())
			var _985_v13 _dafny.Set
			_ = _985_v13
			_985_v13 = _dafny.SetOf((_984_v12).Cardinality(), (_966_v2).F36())
			var _986_v14 D9
			_ = _986_v14
			_986_v14 = Companion_D9_.Create_DC28_(_967_v3, (_966_v2).F35(), (_966_v2).F36(), (_966_v2).F35(), (_966_v2).F35())
			var _source16 D1 = Companion_Default___.Fm43(_dafny.MultiSetOf((_985_v13).Cardinality()), (_986_v14).Dtor_cf43(), globalState)
			_ = _source16
			if _source16.Is_DC4() {
				var _987___mcc_h3 bool = _source16.Get_().(D1_DC4).Cf8
				_ = _987___mcc_h3
				var _988_cf8 bool = _987___mcc_h3
				_ = _988_cf8
				var _rhs167 _dafny.Int = (_this).F34()
				_ = _rhs167
				var _rhs168 bool = (_967_v3).Select((Companion_Default___.SafeIndex((_966_v2).F36(), _dafny.IntOfUint32((_967_v3).Cardinality()))).Uint32()).(bool)
				_ = _rhs168
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				_lhs160.F19 = _rhs167
				_lhs161.F18 = _rhs168
				var _989_v15 _dafny.Sequence
				_ = _989_v15
				_989_v15 = _dafny.SeqOf(_dafny.IntOfUint32((_965_v1).Cardinality()), (_966_v2).F36())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_962_v0), 0))
				_ = _index159
				(_962_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_989_v15, _989_v15)).Cardinality()), (_index159).Int())
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_962_v0), 0))
				_ = _index160
				(_962_v0).ArraySet1(p0, (_index160).Int())
				(globalState).F11 = _965_v1
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_962_v0), 0))
				_ = _index161
				(_962_v0).ArraySet1(((_966_v2).F36()).Minus(p0), (_index161).Int())
			} else if _source16.Is_DC3() {
				var _990___mcc_h4 _dafny.Sequence = _source16.Get_().(D1_DC3).Cf7
				_ = _990___mcc_h4
				var _991_cf7 _dafny.Sequence = _990___mcc_h4
				_ = _991_cf7
				_985_v13 = _985_v13
				var _992_v17 _dafny.Set
				_ = _992_v17
				_992_v17 = _dafny.SetOf(_970_cf8, _970_cf8, (_this).F33())
				var _993_v18 _dafny.Map
				_ = _993_v18
				_993_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_992_v17, (_966_v2).F35())
				var _994_v19 _dafny.Map
				_ = _994_v19
				_994_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_966_v2).F36(), (func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter49 := _dafny.Iterate((_993_v18).Keys().Elements()); ; {
						_compr_41, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _995_v16 _dafny.Set
						_995_v16 = interface{}(_compr_41).(_dafny.Set)
						if (_993_v18).Contains(_995_v16) {
							_coll41.Add(_995_v16, (_966_v2).F36())
						}
					}
					return _coll41.ToMap()
				}()).Cardinality())
				var _rhs169 bool = (_this).F33()
				_ = _rhs169
				var _rhs170 _dafny.Map = (func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(73), _dafny.IntOfInt64(979))); ; {
						_compr_42, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _996_v20 _dafny.Int
						_996_v20 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(73)).Cmp(_996_v20) <= 0) && ((_996_v20).Cmp(_dafny.IntOfInt64(979)) < 0) {
							_coll42.Add(Companion_Default___.SafeModuloInt(_996_v20, p0), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))
						}
					}
					return _coll42.ToMap()
				}()).Update(_dafny.IntOfInt64(551), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_965_v1, (Companion_Default___.SafeIndex((_966_v2).F36(), _dafny.IntOfUint32((_965_v1).Cardinality()))).Uint32(), _972_v6)).Cardinality()))
				_ = _rhs170
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				_lhs162.F28 = _rhs169
				_994_v19 = _rhs170
				var _997_v21 _dafny.Map
				_ = _997_v21
				_997_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_991_cf7, _991_cf7), p0)
				_997_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_cf7, _dafny.IntOfInt64(-239))
				(globalState).F20 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7((_966_v2).F36(), globalState), _965_v1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(((_966_v2).F36()).Times((_966_v2).F36()))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7((_966_v2).F36(), globalState), _965_v1)).Cardinality()))).Uint32(), _972_v6)
			} else {
				var _998___mcc_h5 D1 = _source16.Get_().(D1_DC5).Cf9
				_ = _998___mcc_h5
				var _999_cf9 D1 = _998___mcc_h5
				_ = _999_cf9
				var _1000_v22 *C1
				_ = _1000_v22
				var _nw166 *C1 = New_C1_()
				_ = _nw166
				_nw166.Ctor__((_966_v2).F35())
				_1000_v22 = _nw166
				var _1001_v23 _dafny.Set
				_ = _1001_v23
				_1001_v23 = _dafny.SetOf(_962_v0)
				var _1002_v24 *C2
				_ = _1002_v24
				var _nw167 *C2 = New_C2_()
				_ = _nw167
				_nw167.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_967_v3, _967_v3)).Cardinality()), _1001_v23, (_966_v2).F35(), _965_v1)
				_1002_v24 = _nw167
				(globalState).F28 = _970_cf8
				var _1003_v25 _dafny.Sequence
				_ = _1003_v25
				_1003_v25 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jj"))
				_1003_v25 = _dafny.Companion_Sequence_.Update(_1003_v25, (Companion_Default___.SafeIndex((_966_v2).F36(), _dafny.IntOfUint32((_1003_v25).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_965_v1, _965_v1))
			}
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_973_v7), 0))
			_ = _index162
			(_973_v7).ArraySet1(false, (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_973_v7), 0))
			_ = _index163
			(_973_v7).ArraySet1((_966_v2).F35(), (_index163).Int())
			(globalState).F14 = p0
		} else if _source15.Is_DC3() {
			var _1004___mcc_h1 _dafny.Sequence = _source15.Get_().(D1_DC3).Cf7
			_ = _1004___mcc_h1
			var _1005_cf7 _dafny.Sequence = _1004___mcc_h1
			_ = _1005_cf7
			(globalState).F28 = (_966_v2).F35()
			r0 = ((_966_v2).F35()) == ((_966_v2).F35())
			var _1006_v26 _dafny.Array
			_ = _1006_v26
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_27
			var _nw168 _dafny.Array
			_ = _nw168
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw168 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = (func(_1007_v2 *C3) func(_dafny.Int) bool {
					return func(_1008_i3 _dafny.Int) bool {
						return (_1007_v2).F35()
					}
				})(_966_v2)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw168 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw168).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw168).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1006_v26 = _nw168
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_1006_v26), 0))
			_ = _index164
			(_1006_v26).ArraySet1((_966_v2).F35(), (_index164).Int())
			var _1009_v27 _dafny.MultiSet
			_ = _1009_v27
			_1009_v27 = _dafny.MultiSetOf(_962_v0, _962_v0)
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_1006_v26), 0))
			_ = _index165
			(_1006_v26).ArraySet1(((_1009_v27).Intersection(_1009_v27)).Equals(_1009_v27), (_index165).Int())
			(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_965_v1, _dafny.Companion_Sequence_.Concatenate(_965_v1, _965_v1))
		} else {
			var _1010___mcc_h2 D1 = _source15.Get_().(D1_DC5).Cf9
			_ = _1010___mcc_h2
			var _1011_cf9 D1 = _1010___mcc_h2
			_ = _1011_cf9
			var _1012_v28 _dafny.Sequence
			_ = _1012_v28
			var _1013_v29 bool
			_ = _1013_v29
			var _out14 _dafny.Sequence
			_ = _out14
			var _out15 bool
			_ = _out15
			_out14, _out15 = (_966_v2).M2(!((_966_v2).F35()), globalState)
			_1012_v28 = _out14
			_1013_v29 = _out15
			var _1014_v30 _dafny.MultiSet
			_ = _1014_v30
			_1014_v30 = _dafny.MultiSetOf(_dafny.IntOfInt64(-282), (_this).F34())
			var _1015_v31 D3
			_ = _1015_v31
			_1015_v31 = Companion_D3_.Create_DC10_(_1014_v30)
			var _1016_v32 _dafny.Sequence
			_ = _1016_v32
			_1016_v32 = _dafny.SeqOf(Companion_D3_.Create_DC10_(_1014_v30), _1015_v31, _1015_v31)
			var _1017_v33 _dafny.Map
			_ = _1017_v33
			_1017_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(26), _1016_v32)
			var _1018_v34 _dafny.CodePoint
			_ = _1018_v34
			_1018_v34 = _dafny.CodePoint('e')
			_1017_v33 = ((func() _dafny.Map {
				if Companion_Default___.Fm19(_1018_v34, globalState) {
					return _1017_v33
				}
				return func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(717), _dafny.IntOfInt64(624))); ; {
						_compr_43, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1019_v35 _dafny.Int
						_1019_v35 = interface{}(_compr_43).(_dafny.Int)
						if ((_dafny.IntOfInt64(717)).Cmp(_1019_v35) <= 0) && ((_1019_v35).Cmp(_dafny.IntOfInt64(624)) < 0) {
							_coll43.Add((_1019_v35).Plus(p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(888))).Uint32(), func(coer80 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
								return func(arg80 _dafny.Int) interface{} {
									return coer80(arg80)
								}
							}((func(_1020_v31 D3) func(_dafny.Int) D3 {
								return func(_1021_i4 _dafny.Int) D3 {
									return _1020_v31
								}
							})(_1015_v31))))
						}
					}
					return _coll43.ToMap()
				}()
			})()).Update((_dafny.Zero).Minus(((_966_v2).F36()).Minus(p0)), _dafny.Companion_Sequence_.Concatenate(_1016_v32, _1016_v32))
			var _1022_v36 T0
			_ = _1022_v36
			var _nw169 *C1 = New_C1_()
			_ = _nw169
			_nw169.Ctor__((_966_v2).F35())
			_1022_v36 = _nw169
			var _1023_v37 _dafny.Array
			_ = _1023_v37
			var _nwElement0_32 T0 = _1022_v36
			_ = _nwElement0_32
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(18))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_32, 0)
			(_nw170).ArraySet1(_1022_v36, 1)
			(_nw170).ArraySet1(_1022_v36, 2)
			(_nw170).ArraySet1(_1022_v36, 3)
			(_nw170).ArraySet1(_1022_v36, 4)
			(_nw170).ArraySet1(_1022_v36, 5)
			(_nw170).ArraySet1(_1022_v36, 6)
			(_nw170).ArraySet1(_1022_v36, 7)
			(_nw170).ArraySet1(_1022_v36, 8)
			(_nw170).ArraySet1(_1022_v36, 9)
			(_nw170).ArraySet1(_1022_v36, 10)
			(_nw170).ArraySet1(_1022_v36, 11)
			(_nw170).ArraySet1(_1022_v36, 12)
			(_nw170).ArraySet1(_1022_v36, 13)
			(_nw170).ArraySet1(_1022_v36, 14)
			(_nw170).ArraySet1((func() T0 {
				if Companion_Default___.Fm19(_dafny.CodePoint('e'), globalState) {
					return _1022_v36
				}
				return _1022_v36
			})(), 15)
			(_nw170).ArraySet1(_1022_v36, 16)
			(_nw170).ArraySet1(_1022_v36, 17)
			_1023_v37 = _nw170
			var _1024_v38 _dafny.Map
			_ = _1024_v38
			_1024_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_966_v2).F35(), _1023_v37)
			_1023_v37 = (func() _dafny.Array {
				if (_1024_v38).Contains(!(true)) {
					return (_1024_v38).Get(!(true)).(_dafny.Array)
				}
				return _1023_v37
			})()
			if !((_966_v2).F35()) || (!(((_966_v2).F36()).Cmp(_dafny.IntOfInt64(872)) < 0)) {
				(globalState).F2 = (_966_v2).F36()
				var _1025_v39 _dafny.Array
				_ = _1025_v39
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw171
				_1025_v39 = _nw171
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_1025_v39), 0))
				_ = _index166
				(_1025_v39).ArraySet1(_1013_v29, (_index166).Int())
				var _1026_v40 *C0
				_ = _1026_v40
				var _nw172 *C0 = New_C0_()
				_ = _nw172
				_nw172.Ctor__((_966_v2).F35())
				_1026_v40 = _nw172
				var _1027_v41 _dafny.Sequence
				_ = _1027_v41
				_1027_v41 = _dafny.SeqOf(_1026_v40, _1026_v40, _1026_v40)
				var _1028_v42 _dafny.MultiSet
				_ = _1028_v42
				_1028_v42 = _dafny.MultiSetOf(_1027_v41, _1027_v41, _dafny.Companion_Sequence_.Concatenate(_1027_v41, _1027_v41))
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_962_v0), 0))
				_ = _index167
				(_962_v0).ArraySet1((_966_v2).F36(), (_index167).Int())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_1025_v39), 0))
				_ = _index168
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_962_v0), 0))
				_ = _index169
				var _rhs171 bool = !(_1026_v40.F37)
				_ = _rhs171
				var _rhs172 _dafny.Array = _1025_v39
				_ = _rhs172
				var _rhs173 bool = (_967_v3).Select((Companion_Default___.SafeIndex((_966_v2).F36(), _dafny.IntOfUint32((_967_v3).Cardinality()))).Uint32()).(bool)
				_ = _rhs173
				var _rhs174 _dafny.MultiSet = _1028_v42
				_ = _rhs174
				var _rhs175 _dafny.Int = (_dafny.Zero).Minus((_966_v2).F36())
				_ = _rhs175
				var _lhs163 _dafny.Array = _1025_v39
				_ = _lhs163
				var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_1025_v39), 0))
				_ = _lhs164
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				var _lhs166 _dafny.Array = _962_v0
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_962_v0), 0))
				_ = _lhs167
				(_lhs163).ArraySet1(_rhs171, (_lhs164).Int())
				_1025_v39 = _rhs172
				_lhs165.F28 = _rhs173
				_1028_v42 = _rhs174
				(_lhs166).ArraySet1(_rhs175, (_lhs167).Int())
				(globalState).F19 = (_966_v2).F36()
				var _1029_v43 _dafny.Set
				_ = _1029_v43
				_1029_v43 = _dafny.SetOf((_966_v2).F35())
				var _1030_v44 _dafny.Sequence
				_ = _1030_v44
				_1030_v44 = _dafny.SeqOf(_dafny.SetOf(_1013_v29))
				var _1031_v45 _dafny.Map
				_ = _1031_v45
				_1031_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_966_v2).F36(), _1029_v43)
				var _1032_v46 _dafny.Array
				_ = _1032_v46
				var _nwElement0_33 _dafny.Set = _1029_v43
				_ = _nwElement0_33
				var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(22))
				_ = _nw173
				(_nw173).ArraySet1(_nwElement0_33, 0)
				(_nw173).ArraySet1(_1029_v43, 1)
				(_nw173).ArraySet1(_1029_v43, 2)
				(_nw173).ArraySet1(_1029_v43, 3)
				(_nw173).ArraySet1(_1029_v43, 4)
				(_nw173).ArraySet1(Companion_Default___.Fm22((_this).F34(), (_this).F33(), _1018_v34, globalState), 5)
				(_nw173).ArraySet1((_1029_v43).Union(_1029_v43), 6)
				(_nw173).ArraySet1((_dafny.SetOf(true)).Union(_1029_v43), 7)
				(_nw173).ArraySet1((_1029_v43).Intersection(_dafny.SetOf(_1026_v40.F37, (_966_v2).F35(), _1026_v40.F37)), 8)
				(_nw173).ArraySet1(_1029_v43, 9)
				(_nw173).ArraySet1(_dafny.SetOf((_1025_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_1025_v39), 0))).Int()).(bool)), 10)
				(_nw173).ArraySet1((_1030_v44).Select((Companion_Default___.SafeIndex((_966_v2).F36(), _dafny.IntOfUint32((_1030_v44).Cardinality()))).Uint32()).(_dafny.Set), 11)
				(_nw173).ArraySet1(_1029_v43, 12)
				(_nw173).ArraySet1(_1029_v43, 13)
				(_nw173).ArraySet1(_dafny.SetOf(_1026_v40.F37), 14)
				(_nw173).ArraySet1(_1029_v43, 15)
				(_nw173).ArraySet1((Companion_Default___.Fm22((_966_v2).F36(), _1013_v29, _1018_v34, globalState)).Difference(_1029_v43), 16)
				(_nw173).ArraySet1(_dafny.SetOf((_966_v2).F35()), 17)
				(_nw173).ArraySet1(_dafny.SetOf(_1026_v40.F37), 18)
				(_nw173).ArraySet1((_1029_v43).Difference((func() _dafny.Set {
					if (_1031_v45).Contains((_966_v2).F36()) {
						return (_1031_v45).Get((_966_v2).F36()).(_dafny.Set)
					}
					return _1029_v43
				})()), 19)
				(_nw173).ArraySet1(_1029_v43, 20)
				(_nw173).ArraySet1((_1029_v43).Difference(_1029_v43), 21)
				_1032_v46 = _nw173
				_1032_v46 = _1032_v46
				(globalState).F18 = (_966_v2).F35()
			} else {
				(globalState).F28 = Companion_Default___.Fm6(_dafny.IntOfInt64(-894), globalState)
				_1018_v34 = _1018_v34
				var _1033_v47 _dafny.Set
				_ = _1033_v47
				_1033_v47 = _dafny.SetOf((_966_v2).F35(), (_966_v2).F35(), !((_966_v2).F35()), (_966_v2).F35(), (_966_v2).F35())
				_1033_v47 = _1033_v47
				var _1034_v48 _dafny.Array
				_ = _1034_v48
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_28
				var _nw174 _dafny.Array
				_ = _nw174
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw174 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Sequence = (func(_1035_v2 *C3, _1036_p0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1037_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1035_v2).F36()), _dafny.SeqOf(_1036_p0, _dafny.IntOfInt64(814), _1036_p0, (_1035_v2).F36(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1035_v2).F36(), (_1035_v2).F36())).Cardinality()))
						}
					})(_966_v2, p0)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw174 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw174).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw174).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1034_v48 = _nw174
				var _1038_v49 _dafny.Sequence
				_ = _1038_v49
				_1038_v49 = _dafny.SeqOf((_this).F34(), (_dafny.Zero).Minus((_966_v2).F36()))
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1034_v48), 0))
				_ = _index170
				(_1034_v48).ArraySet1(_1038_v49, (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1034_v48), 0))
				_ = _index171
				(_1034_v48).ArraySet1(_1038_v49, (_index171).Int())
				(globalState).F21 = (_1038_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-844), _dafny.IntOfUint32((_1038_v49).Cardinality()))).Uint32()).(_dafny.Int)
			}
		}
		(globalState).F21 = (p0).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oxjnf")).Cardinality()), (_this).F34()))
		(globalState).F21 = ((_966_v2).F36()).Times(Companion_Default___.SafeModuloInt((_966_v2).F36(), (_966_v2).F36()))
		var _1039_v50 _dafny.CodePoint
		_ = _1039_v50
		_1039_v50 = _dafny.CodePoint('g')
		var _1040_v51 _dafny.Set
		_ = _1040_v51
		_1040_v51 = _dafny.SetOf(_1039_v50, Companion_Default___.Fm42(true, !((_this).F33()), globalState), (func() _dafny.CodePoint {
			if (_this).F33() {
				return _1039_v50
			}
			return _1039_v50
		})())
		_1040_v51 = _dafny.SetOf(_1039_v50, _1039_v50, Companion_Default___.Fm42((_this).F33(), (_966_v2).F35(), globalState), _1039_v50, _1039_v50)
		r0 = (_966_v2).F35()
		return r0
	}
}
func (_this *C4) F33() bool {
	{
		return _this._f33
	}
}
func (_this *C4) F34() _dafny.Int {
	{
		return _this._f34
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f30 _dafny.Sequence
	_f29 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f30 = _dafny.EmptySeq
	_this._f29 = false
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

func (_this *C5) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C5) F30_set_(value _dafny.Sequence) {
	_this._f30 = value
}
func (_this *C5) F29() bool {
	return _this._f29
}
func (_this *C5) Ctor__(f29 bool, f30 _dafny.Sequence) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C5) Fm3(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		if ((_this).F29()) == ((_this).F29()) {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())).Cardinality())
		} else {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality()))
		}
	}
}
func (_this *C5) Fm4(globalState *GlobalState) _dafny.Int {
	{
		return ((func() _dafny.Int {
			if (_this).F29() {
				return ((Companion_D2_.Create_DC6_(_dafny.MultiSetOf(_dafny.CodePoint('p')))).Dtor_cf10()).Cardinality()
			}
			return _dafny.IntOfInt64(255)
		})()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kxmebkric"), _dafny.UnicodeSeqOfUtf8Bytes("qe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-273))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1041_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), _dafny.UnicodeSeqOfUtf8Bytes("ugd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(515))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1042_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})))).Cardinality()))
	}
}
func (_this *C5) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1043_v0 _dafny.Array
		_ = _1043_v0
		var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw175
		_1043_v0 = _nw175
		var _1044_v1 _dafny.Array
		_ = _1044_v1
		var _nwElement0_34 bool = (_this).F29()
		_ = _nwElement0_34
		var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(13))
		_ = _nw176
		(_nw176).ArraySet1(_nwElement0_34, 0)
		(_nw176).ArraySet1((_this).F29(), 1)
		(_nw176).ArraySet1((_this).F29(), 2)
		(_nw176).ArraySet1(!(false), 3)
		(_nw176).ArraySet1((_this).F29(), 4)
		(_nw176).ArraySet1(p0, 5)
		(_nw176).ArraySet1((_this).F29(), 6)
		(_nw176).ArraySet1(p0, 7)
		(_nw176).ArraySet1(p0, 8)
		(_nw176).ArraySet1((_this).F29(), 9)
		(_nw176).ArraySet1((_this).F29(), 10)
		(_nw176).ArraySet1((_this).F29(), 11)
		(_nw176).ArraySet1((_this).F29(), 12)
		_1044_v1 = _nw176
		var _1045_v2 _dafny.Int
		_ = _1045_v2
		_1045_v2 = _dafny.IntOfInt64(70)
		var _1046_v3 _dafny.Sequence
		_ = _1046_v3
		_1046_v3 = _dafny.SeqOf((_this).F29())
		var _1047_v4 _dafny.Map
		_ = _1047_v4
		_1047_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1043_v0, (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1044_v1), (Companion_Default___.SafeIndex(_1045_v2, _dafny.IntOfUint32((_dafny.SeqOf(_1044_v1)).Cardinality()))).Uint32(), _1044_v1)).Cardinality())).Times((_this).Fm3(_1046_v3, globalState))))
		_1047_v4 = (_1047_v4).Update(_1043_v0, (_this).Fm4(globalState))
		var _1048_v5 _dafny.Sequence
		_ = _1048_v5
		_1048_v5 = _dafny.SeqOf(_1045_v2, _1045_v2)
		var _hi5 _dafny.Int = (_1048_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1048_v5).Cardinality()), _dafny.IntOfUint32((_1048_v5).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi5
		for _1049_i0 := (_1045_v2).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg83 _dafny.Int) interface{} {
				return coer83(arg83)
			}
		}((func(_1053_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1054_i1 _dafny.Int) _dafny.Int {
				return _1053_v2
			}
		})(_1045_v2)))).Cardinality())); _1049_i0.Cmp(_hi5) < 0; _1049_i0 = _1049_i0.Plus(_dafny.One) {
			var _1050_v6 _dafny.MultiSet
			_ = _1050_v6
			_1050_v6 = _dafny.MultiSetOf((_this).F29(), (_this).F29())
			var _rhs176 bool = (_1050_v6).IsSubsetOf(_1050_v6)
			_ = _rhs176
			var _rhs177 bool = (_1049_i0).Cmp(_1045_v2) < 0
			_ = _rhs177
			var _lhs168 *GlobalState = globalState
			_ = _lhs168
			var _lhs169 *GlobalState = globalState
			_ = _lhs169
			_lhs168.F18 = _rhs176
			_lhs169.F18 = _rhs177
			(globalState).F18 = false
			var _1051_v7 _dafny.Array
			_ = _1051_v7
			var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw177
			_1051_v7 = _nw177
			var _1052_v8 _dafny.Map
			_ = _1052_v8
			_1052_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _1051_v7)
			_1052_v8 = (_1052_v8).Update(((_1050_v6).Cardinality()).Cmp(_1045_v2) >= 0, _1051_v7)
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1051_v7), 0))
			_ = _index172
			(_1051_v7).ArraySet1(_dafny.IntOfInt64(967), (_index172).Int())
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1051_v7), 0))
			_ = _index173
			(_1051_v7).ArraySet1((_1045_v2).Times(_1045_v2), (_index173).Int())
		}
		(globalState).F21 = _1045_v2
		var _hi6 _dafny.Int = _1045_v2
		_ = _hi6
		for _1055_i2 := _1045_v2; _1055_i2.Cmp(_hi6) < 0; _1055_i2 = _1055_i2.Plus(_dafny.One) {
			var _1056_v9 _dafny.Array
			_ = _1056_v9
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_29
			var _nw178 _dafny.Array
			_ = _nw178
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw178 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.MultiSet = (func(_1057_v3 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
					return func(_1058_i3 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(Companion_D1_.Create_DC3_(_1057_v3))
					}
				})(_1046_v3)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw178 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw178).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw178).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1056_v9 = _nw178
			_1056_v9 = _1056_v9
			(globalState).F13 = p0
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))
			_ = _index174
			(_1044_v1).ArraySet1((_this).F29(), (_index174).Int())
			var _1059_v10 _dafny.Map
			_ = _1059_v10
			_1059_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_i2, p0)
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))
			_ = _index175
			(_1044_v1).ArraySet1((func() bool {
				if p0 {
					return (func() bool {
						if (_1059_v10).Contains(((func() _dafny.Map {
							var _coll44 = _dafny.NewMapBuilder()
							_ = _coll44
							for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(228), _dafny.IntOfInt64(630))); ; {
								_compr_44, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _1060_v11 _dafny.Int
								_1060_v11 = interface{}(_compr_44).(_dafny.Int)
								if ((_dafny.IntOfInt64(228)).Cmp(_1060_v11) <= 0) && ((_1060_v11).Cmp(_dafny.IntOfInt64(630)) < 0) {
									_coll44.Add((_1060_v11).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spsbhq")).Cardinality())), _1055_i2)
								}
							}
							return _coll44.ToMap()
						}()).Update(_dafny.IntOfInt64(416), _1055_i2)).Cardinality()) {
							return (_1059_v10).Get(((func() _dafny.Map {
								var _coll45 = _dafny.NewMapBuilder()
								_ = _coll45
								for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(228), _dafny.IntOfInt64(630))); ; {
									_compr_45, _ok53 := _iter53()
									if !_ok53 {
										break
									}
									var _1061_v11 _dafny.Int
									_1061_v11 = interface{}(_compr_45).(_dafny.Int)
									if ((_dafny.IntOfInt64(228)).Cmp(_1061_v11) <= 0) && ((_1061_v11).Cmp(_dafny.IntOfInt64(630)) < 0) {
										_coll45.Add((_1061_v11).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spsbhq")).Cardinality())), _1055_i2)
									}
								}
								return _coll45.ToMap()
							}()).Update(_dafny.IntOfInt64(416), _1055_i2)).Cardinality()).(bool)
						}
						return p0
					})()
				}
				return (_1055_i2).Cmp(_1045_v2) == 0
			})(), (_index175).Int())
			if p0 {
				var _1062_v12 _dafny.Array
				_ = _1062_v12
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw179
				_1062_v12 = _nw179
				var _1063_v13 D1
				_ = _1063_v13
				_1063_v13 = Companion_D1_.Create_DC3_(_1046_v3)
				var _1064_v14 D1
				_ = _1064_v14
				_1064_v14 = Companion_D1_.Create_DC5_(_1063_v13)
				var _1065_v15 D1
				_ = _1065_v15
				_1065_v15 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(_1063_v13))
				var _1066_v16 D1
				_ = _1066_v16
				_1066_v16 = Companion_D1_.Create_DC5_(_1063_v13)
				var _1067_v17 _dafny.Map
				_ = _1067_v17
				_1067_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v12, _1066_v16)
				_1067_v17 = ((func() _dafny.Map {
					if (_this).F29() {
						return _1067_v17
					}
					return (_1067_v17).Update(_1062_v12, _1066_v16)
				})()).Update(_1062_v12, _1066_v16)
				var _1068_v18 _dafny.Set
				_ = _1068_v18
				_1068_v18 = _dafny.SetOf(p0, p0, true)
				(globalState).F21 = ((_1055_i2).Plus((_1068_v18).Cardinality())).Plus(_dafny.IntOfInt64(556))
				var _1069_v19 _dafny.Map
				_ = _1069_v19
				_1069_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v12, (_1044_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))).Int()).(bool))
				var _1070_v20 _dafny.Map
				_ = _1070_v20
				_1070_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1069_v19).Cardinality(), _1045_v2)
				_1070_v20 = (_1070_v20).Update(_dafny.IntOfUint32((_this.F30()).Cardinality()), _1045_v2)
				var _1071_v21 _dafny.Map
				_ = _1071_v21
				_1071_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p0)
				_1070_v20 = (_1070_v20).Update((_1045_v2).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1072_v20 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_1073_i4 _dafny.Int) _dafny.Int {
						return (_1072_v20).Cardinality()
					}
				})(_1070_v20))), (Companion_Default___.SafeIndex((_1071_v21).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1074_v20 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_1075_i4 _dafny.Int) _dafny.Int {
						return (_1074_v20).Cardinality()
					}
				})(_1070_v20)))).Cardinality()))).Uint32(), _1055_i2)).Cardinality())), (_1055_i2).Minus(_dafny.IntOfInt64(826)))
				var _1076_v22 _dafny.Array
				_ = _1076_v22
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_30
				var _nw180 _dafny.Array
				_ = _nw180
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw180 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = (func(_1077_p0 bool) func(_dafny.Int) bool {
						return func(_1078_i5 _dafny.Int) bool {
							return _1077_p0
						}
					})(p0)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw180 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw180).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw180).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1076_v22 = _nw180
				var _1079_v23 D0
				_ = _1079_v23
				_1079_v23 = Companion_D0_.Create_DC1_(_1076_v22, (func() _dafny.Int {
					if (_1070_v20).Contains(_1055_i2) {
						return (_1070_v20).Get(_1055_i2).(_dafny.Int)
					}
					return _1055_i2
				})(), (_dafny.Zero).Minus(_1055_i2))
				_1076_v22 = (_1079_v23).Dtor_cf1()
			} else {
				var _1080_v24 _dafny.Array
				_ = _1080_v24
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw181
				_1080_v24 = _nw181
				var _1081_v25 _dafny.MultiSet
				_ = _1081_v25
				_1081_v25 = _dafny.MultiSetOf(p0)
				_1080_v24 = (func() _dafny.Array {
					if (_dafny.MultiSetFromSeq(_1046_v3)).IsDisjointFrom(_1081_v25) {
						return _1080_v24
					}
					return _1080_v24
				})()
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))
				_ = _index176
				(_1044_v1).ArraySet1((_1055_i2).Cmp(_1045_v2) != 0, (_index176).Int())
				var _1082_v26 _dafny.Map
				_ = _1082_v26
				_1082_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1044_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))).Int()).(bool), _1045_v2)
				(globalState).F21 = (_dafny.Zero).Minus((_1045_v2).Minus(Companion_Default___.SafeModuloInt(_1045_v2, (func() _dafny.Int {
					if (_1082_v26).Contains((_this).F29()) {
						return (_1082_v26).Get((_this).F29()).(_dafny.Int)
					}
					return _1045_v2
				})())))
				r1 = (_1044_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1044_v1), 0))).Int()).(bool)
				var _1083_v27 *C0
				_ = _1083_v27
				var _nw182 *C0 = New_C0_()
				_ = _nw182
				_nw182.Ctor__(p0)
				_1083_v27 = _nw182
			}
		}
		var _1084_v28 D9
		_ = _1084_v28
		_1084_v28 = Companion_D9_.Create_DC27_(_1048_v5)
		var _1085_v29 _dafny.Array
		_ = _1085_v29
		var _nwElement0_35 _dafny.Sequence = (_1084_v28).Dtor_cf40()
		_ = _nwElement0_35
		var _nw183 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(6))
		_ = _nw183
		(_nw183).ArraySet1(_nwElement0_35, 0)
		(_nw183).ArraySet1(_1048_v5, 1)
		(_nw183).ArraySet1(_1048_v5, 2)
		(_nw183).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1048_v5, _1048_v5), 3)
		(_nw183).ArraySet1((func() _dafny.Sequence {
			if (_this).F29() {
				return _1048_v5
			}
			return _dafny.SeqOf(_dafny.IntOfInt64(566))
		})(), 4)
		(_nw183).ArraySet1(_dafny.SeqOf(_1045_v2, _1045_v2), 5)
		_1085_v29 = _nw183
		var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1085_v29), 0))
		_ = _index177
		(_1085_v29).ArraySet1(_1048_v5, (_index177).Int())
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1085_v29), 0))
		_ = _index178
		(_1085_v29).ArraySet1((func() _dafny.Sequence {
			if (_this).F29() {
				return (func() _dafny.Sequence {
					if p0 {
						return _1048_v5
					}
					return _1048_v5
				})()
			}
			return _1048_v5
		})(), (_index178).Int())
		var _1086_v30 _dafny.MultiSet
		_ = _1086_v30
		_1086_v30 = _dafny.MultiSetOf((p0) == (true), (_1046_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfUint32((_1046_v3).Cardinality()))).Uint32()).(bool), (_this).F29())
		_1086_v30 = _1086_v30
		var _1087_v31 _dafny.CodePoint
		_ = _1087_v31
		_1087_v31 = _dafny.CodePoint('e')
		var _1088_v32 D2
		_ = _1088_v32
		_1088_v32 = Companion_D2_.Create_DC8_(_1087_v31, (_this).F29())
		var _1089_v33 _dafny.Sequence
		_ = _1089_v33
		_1089_v33 = _dafny.SeqOf(_1088_v32)
		var _1090_v34 _dafny.Sequence
		_ = _1090_v34
		_1090_v34 = _dafny.SeqOf(Companion_Default___.Fm35(_1089_v33, _dafny.IntOfInt64(654), globalState), _dafny.SeqOf((_1046_v3).Select((Companion_Default___.SafeIndex(_1045_v2, _dafny.IntOfUint32((_1046_v3).Cardinality()))).Uint32()).(bool), (_this).F29(), p0), _1046_v3)
		r0 = _dafny.Companion_Sequence_.Concatenate(_1090_v34, _1090_v34)
		r1 = (_this).F29()
		return r0, r1
	}
}
func (_this *C5) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1091_v0 D4
		_ = _1091_v0
		_1091_v0 = Companion_D4_.Create_DC15_(p2, p2)
		var _1092_i0 _dafny.Int
		_ = _1092_i0
		_1092_i0 = _dafny.Zero
		{
			var _pat_let_tv22 = p0
			_ = _pat_let_tv22
			for func(_source17 D4) bool {
				if _source17.Is_DC14() {
					var _1097___mcc_h0 _dafny.Int = _source17.Get_().(D4_DC14).Cf23
					_ = _1097___mcc_h0
					var _1098_cf23 _dafny.Int = _1097___mcc_h0
					_ = _1098_cf23
					return (_this).F29()
				} else if _source17.Is_DC15() {
					var _1099___mcc_h1 _dafny.Int = _source17.Get_().(D4_DC15).Cf24
					_ = _1099___mcc_h1
					var _1100___mcc_h2 _dafny.Int = _source17.Get_().(D4_DC15).Cf25
					_ = _1100___mcc_h2
					var _1101_cf25 _dafny.Int = _1100___mcc_h2
					_ = _1101_cf25
					var _1102_cf24 _dafny.Int = _1099___mcc_h1
					_ = _1102_cf24
					return _pat_let_tv22
				} else {
					var _1103___mcc_h3 _dafny.Map = _source17.Get_().(D4_DC13).Cf22
					_ = _1103___mcc_h3
					var _1104_cf22 _dafny.Map = _1103___mcc_h3
					_ = _1104_cf22
					return (_this).F29()
				}
			}(_1091_v0) {
				{
					if (_1092_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1092_i0 = (_1092_i0).Plus(_dafny.One)
					var _rhs178 bool = (_this).F29()
					_ = _rhs178
					var _rhs179 bool = (_this).F29()
					_ = _rhs179
					var _lhs170 *GlobalState = globalState
					_ = _lhs170
					var _lhs171 *GlobalState = globalState
					_ = _lhs171
					_lhs170.F13 = _rhs178
					_lhs171.F15 = _rhs179
					var _1093_v1 _dafny.Array
					_ = _1093_v1
					var _nw184 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
					_ = _nw184
					_1093_v1 = _nw184
					var _1094_v2 T0
					_ = _1094_v2
					var _nw185 *C4 = New_C4_()
					_ = _nw185
					_nw185.Ctor__(p0, p2)
					_1094_v2 = _nw185
					var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_1093_v1), 0))
					_ = _index179
					(_1093_v1).ArraySet1(_1094_v2, (_index179).Int())
					var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_1093_v1), 0))
					_ = _index180
					var _rhs180 _dafny.Sequence = _this.F30()
					_ = _rhs180
					var _rhs181 T0 = _1094_v2
					_ = _rhs181
					var _rhs182 bool = Companion_Default___.Fm6(p2, globalState)
					_ = _rhs182
					var _rhs183 _dafny.Int = p2
					_ = _rhs183
					var _lhs172 _dafny.Array = _1093_v1
					_ = _lhs172
					var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_1093_v1), 0))
					_ = _lhs173
					var _lhs174 *GlobalState = globalState
					_ = _lhs174
					var _lhs175 *GlobalState = globalState
					_ = _lhs175
					r0 = _rhs180
					(_lhs172).ArraySet1(_rhs181, (_lhs173).Int())
					_lhs174.F18 = _rhs182
					_lhs175.F19 = _rhs183
					(globalState).F2 = p2
					var _1095_v3 D10
					_ = _1095_v3
					_1095_v3 = Companion_D10_.Create_DC33_(p0, p2)
					var _1096_v4 *C3
					_ = _1096_v4
					var _nw186 *C3 = New_C3_()
					_ = _nw186
					_nw186.Ctor__((_1095_v3).Dtor_cf54(), p2, (p2).Cmp(_dafny.IntOfInt64(285)) == 0, _this.F30())
					_1096_v4 = _nw186
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1105_v5 _dafny.Sequence
		_ = _1105_v5
		_1105_v5 = _dafny.SeqOf(p0, (_this).F29())
		var _1106_v6 _dafny.Sequence
		_ = _1106_v6
		_1106_v6 = _dafny.SeqOf(_1105_v5, _1105_v5)
		var _1107_v7 D1
		_ = _1107_v7
		_1107_v7 = Companion_D1_.Create_DC3_((_1106_v6).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1106_v6).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1108_v8 _dafny.Map
		_ = _1108_v8
		_1108_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1107_v7)
		var _1109_v9 D5
		_ = _1109_v9
		_1109_v9 = Companion_D5_.Create_DC16_(_1108_v8)
		var _pat_let_tv23 = _1108_v8
		_ = _pat_let_tv23
		var _source18 D5 = func(_pat_let25_0 D5) D5 {
			return func(_1110_dt__update__tmp_h0 D5) D5 {
				return func(_pat_let26_0 _dafny.Map) D5 {
					return func(_1111_dt__update_hcf26_h0 _dafny.Map) D5 {
						return Companion_D5_.Create_DC16_(_1111_dt__update_hcf26_h0)
					}(_pat_let26_0)
				}(_pat_let_tv23)
			}(_pat_let25_0)
		}(_1109_v9)
		_ = _source18
		if _source18.Is_DC17() {
			var _1112___mcc_h4 _dafny.Set = _source18.Get_().(D5_DC17).Cf27
			_ = _1112___mcc_h4
			var _1113_cf27 _dafny.Set = _1112___mcc_h4
			_ = _1113_cf27
			(globalState).F18 = (_this).F29()
			var _1114_v10 _dafny.CodePoint
			_ = _1114_v10
			_1114_v10 = _dafny.CodePoint('r')
			var _1115_v11 _dafny.MultiSet
			_ = _1115_v11
			_1115_v11 = _dafny.MultiSetOf((_this).F29())
			(globalState).F20 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _1114_v10), (Companion_Default___.SafeIndex((_1115_v11).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _1114_v10)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if (_this).F29() {
					return _1114_v10
				}
				return _1114_v10
			})())
			var _1116_v12 _dafny.Map
			_ = _1116_v12
			_1116_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(Companion_Default___.Fm6(p2, globalState)))
			var _1117_v13 _dafny.Sequence
			_ = _1117_v13
			_1117_v13 = _dafny.SeqOf(p2)
			var _1118_v14 _dafny.Sequence
			_ = _1118_v14
			_1118_v14 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F29()), _dafny.IntOfInt64(342))).Cardinality(), ((func() _dafny.Map {
				if (_this).F29() {
					return _1116_v12
				}
				return Companion_Default___.Fm32((_this).F29(), (_this).F29(), _1117_v13, globalState)
			})()).Cardinality())
			(globalState).F7 = _1117_v13
			(globalState).F1 = Companion_Default___.SafeDivisionInt(p2, (_dafny.Zero).Minus(p2))
		} else if _source18.Is_DC18() {
			(globalState).F18 = (_this).F29()
			var _1119_v15 _dafny.Sequence
			_ = _1119_v15
			_1119_v15 = _dafny.SeqOf(p2)
			var _1120_v16 _dafny.Map
			_ = _1120_v16
			_1120_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1119_v15)
			var _1121_v17 D9
			_ = _1121_v17
			_1121_v17 = Companion_D9_.Create_DC27_((func() _dafny.Sequence {
				if (_1120_v16).Contains(p2) {
					return (_1120_v16).Get(p2).(_dafny.Sequence)
				}
				return Companion_Default___.Fm44((_this).F29(), p2, p0, p0, globalState)
			})())
			_1121_v17 = Companion_D9_.Create_DC27_(_1119_v15)
			var _1122_v18 D0
			_ = _1122_v18
			_1122_v18 = Companion_D0_.Create_DC2_(_dafny.MultiSetOf((_this).F29()), _this.F30(), true)
			var _1123_v19 _dafny.CodePoint
			_ = _1123_v19
			_1123_v19 = _dafny.CodePoint('u')
			var _pat_let_tv24 = _1123_v19
			_ = _pat_let_tv24
			var _pat_let_tv25 = p2
			_ = _pat_let_tv25
			var _pat_let_tv26 = globalState
			_ = _pat_let_tv26
			var _source19 D0 = func(_pat_let27_0 D0) D0 {
				return func(_1124_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let28_0 _dafny.MultiSet) D0 {
						return func(_1125_dt__update_hcf4_h0 _dafny.MultiSet) D0 {
							return Companion_D0_.Create_DC2_(_1125_dt__update_hcf4_h0, (_1124_dt__update__tmp_h1).Dtor_cf5(), (_1124_dt__update__tmp_h1).Dtor_cf6())
						}(_pat_let28_0)
					}(Companion_Default___.Fm39(_pat_let_tv24, _pat_let_tv25, _pat_let_tv26))
				}(_pat_let27_0)
			}(_1122_v18)
			_ = _source19
			if _source19.Is_DC0() {
				var _1126___mcc_h10 bool = _source19.Get_().(D0_DC0).Cf0
				_ = _1126___mcc_h10
				var _1127_cf0 bool = _1126___mcc_h10
				_ = _1127_cf0
				(globalState).F20 = _this.F30()
				var _1128_v20 _dafny.Array
				_ = _1128_v20
				var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw187
				_1128_v20 = _nw187
				var _1129_v21 *C2
				_ = _1129_v21
				var _nw188 *C2 = New_C2_()
				_ = _nw188
				_nw188.Ctor__(p2, (_dafny.SetOf(_1128_v20, _1128_v20)).Intersection(_dafny.SetOf(_1128_v20, _1128_v20, _1128_v20, _1128_v20, _1128_v20)), p0, _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()))
				_1129_v21 = _nw188
				var _1130_v22 _dafny.Array
				_ = _1130_v22
				var _nw189 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw189
				_1130_v22 = _nw189
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1130_v22), 0))
				_ = _index181
				(_1130_v22).ArraySet1(((_this).F29()) || (_1127_cf0), (_index181).Int())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1130_v22), 0))
				_ = _index182
				(_1130_v22).ArraySet1(_1127_cf0, (_index182).Int())
				(globalState).F1 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1129_v21).F38()), p2)).Cardinality()).Times((_1129_v21).F38())
			} else if _source19.Is_DC1() {
				var _1131___mcc_h11 _dafny.Array = _source19.Get_().(D0_DC1).Cf1
				_ = _1131___mcc_h11
				var _1132___mcc_h12 _dafny.Int = _source19.Get_().(D0_DC1).Cf2
				_ = _1132___mcc_h12
				var _1133___mcc_h13 _dafny.Int = _source19.Get_().(D0_DC1).Cf3
				_ = _1133___mcc_h13
				var _1134_cf3 _dafny.Int = _1133___mcc_h13
				_ = _1134_cf3
				var _1135_cf2 _dafny.Int = _1132___mcc_h12
				_ = _1135_cf2
				var _1136_cf1 _dafny.Array = _1131___mcc_h11
				_ = _1136_cf1
				var _1137_v23 _dafny.Array
				_ = _1137_v23
				var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw190
				_1137_v23 = _nw190
				var _1138_v24 _dafny.Array
				_ = _1138_v24
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_31
				var _nw191 _dafny.Array
				_ = _nw191
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw191 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Map = (func(_1139_v15 _dafny.Sequence, _1140_cf2 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1141_i1 _dafny.Int) _dafny.Map {
							return (Companion_D11_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1139_v15, _1140_cf2))).Dtor_cf56()
						}
					})(_1119_v15, _1135_cf2)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw191 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw191).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw191).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1138_v24 = _nw191
				var _1142_v25 _dafny.Map
				_ = _1142_v25
				_1142_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1119_v15, _1134_cf3)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1138_v24), 0))
				_ = _index183
				(_1138_v24).ArraySet1(((_1142_v25).Update(_1119_v15, _1135_cf2)).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1135_cf2), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1135_cf2), _dafny.IntOfUint32((_dafny.SeqOf(_1135_cf2)).Cardinality()))).Uint32(), _1134_cf3), _1135_cf2), (_index183).Int())
				var _pat_let_tv27 = _1135_cf2
				_ = _pat_let_tv27
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1138_v24), 0))
				_ = _index184
				var _rhs184 _dafny.Array = _1137_v23
				_ = _rhs184
				var _rhs185 _dafny.Map = Companion_Default___.Fm45(globalState)
				_ = _rhs185
				var _rhs186 D4 = func(_pat_let29_0 D4) D4 {
					return func(_1143_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let30_0 _dafny.Int) D4 {
							return func(_1144_dt__update_hcf25_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC15_((_1143_dt__update__tmp_h2).Dtor_cf24(), _1144_dt__update_hcf25_h0)
							}(_pat_let30_0)
						}(_pat_let_tv27)
					}(_pat_let29_0)
				}(_1091_v0)
				_ = _rhs186
				var _rhs187 _dafny.Int = _1135_cf2
				_ = _rhs187
				var _rhs188 bool = !(!((p2).Cmp((_dafny.Zero).Minus(Companion_Default___.Fm9(globalState))) == 0))
				_ = _rhs188
				var _lhs176 _dafny.Array = _1138_v24
				_ = _lhs176
				var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1138_v24), 0))
				_ = _lhs177
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				_1137_v23 = _rhs184
				(_lhs176).ArraySet1(_rhs185, (_lhs177).Int())
				_1091_v0 = _rhs186
				_lhs178.F21 = _rhs187
				_lhs179.F15 = _rhs188
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_1136_cf1), 0))
				_ = _index185
				(_1136_cf1).ArraySet1((_this).F29(), (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_1136_cf1), 0))
				_ = _index186
				var _rhs189 _dafny.Int = _1134_cf3
				_ = _rhs189
				var _rhs190 bool = (_this).F29()
				_ = _rhs190
				var _rhs191 bool = p0
				_ = _rhs191
				var _rhs192 bool = Companion_Default___.Fm19(_1123_v19, globalState)
				_ = _rhs192
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 _dafny.Array = _1136_cf1
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_1136_cf1), 0))
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				_lhs180.F19 = _rhs189
				(_lhs181).ArraySet1(_rhs190, (_lhs182).Int())
				_lhs183.F28 = _rhs191
				_lhs184.F15 = _rhs192
				_1135_cf2 = (p2).Minus(_1135_cf2)
				(globalState).F28 = (_1135_cf2).Cmp(p2) <= 0
			} else {
				var _1145___mcc_h14 _dafny.MultiSet = _source19.Get_().(D0_DC2).Cf4
				_ = _1145___mcc_h14
				var _1146___mcc_h15 _dafny.Sequence = _source19.Get_().(D0_DC2).Cf5
				_ = _1146___mcc_h15
				var _1147___mcc_h16 bool = _source19.Get_().(D0_DC2).Cf6
				_ = _1147___mcc_h16
				var _1148_cf6 bool = _1147___mcc_h16
				_ = _1148_cf6
				var _1149_cf5 _dafny.Sequence = _1146___mcc_h15
				_ = _1149_cf5
				var _1150_cf4 _dafny.MultiSet = _1145___mcc_h14
				_ = _1150_cf4
				var _rhs193 _dafny.Int = (_dafny.Zero).Minus(p2)
				_ = _rhs193
				var _rhs194 bool = !(_1148_cf6)
				_ = _rhs194
				var _rhs195 _dafny.MultiSet = _1150_cf4
				_ = _rhs195
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				_lhs185.F19 = _rhs193
				_lhs186.F18 = _rhs194
				_1150_cf4 = _rhs195
				var _1151_v26 *C3
				_ = _1151_v26
				var _nw192 *C3 = New_C3_()
				_ = _nw192
				_nw192.Ctor__(!((_this).F29()), p2, (func() bool {
					if p0 {
						return p0
					}
					return (_this).F29()
				})(), (Companion_D11_.Create_DC35_(_this.F30(), _1148_cf6)).Dtor_cf57())
				_1151_v26 = _nw192
				(globalState).F13 = (_this).F29()
				var _1152_v27 _dafny.Array
				_ = _1152_v27
				var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw193
				_1152_v27 = _nw193
				var _1153_v28 _dafny.Sequence
				_ = _1153_v28
				_1153_v28 = _dafny.SeqOf(_1152_v27, _1152_v27, _1152_v27, _1152_v27, _1152_v27)
				var _1154_v29 _dafny.Array
				_ = _1154_v29
				var _nwElement0_36 _dafny.Array = _1152_v27
				_ = _nwElement0_36
				var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(28))
				_ = _nw194
				(_nw194).ArraySet1(_nwElement0_36, 0)
				(_nw194).ArraySet1(_1152_v27, 1)
				(_nw194).ArraySet1(_1152_v27, 2)
				(_nw194).ArraySet1(_1152_v27, 3)
				(_nw194).ArraySet1(_1152_v27, 4)
				(_nw194).ArraySet1(_1152_v27, 5)
				(_nw194).ArraySet1(_1152_v27, 6)
				(_nw194).ArraySet1(_1152_v27, 7)
				(_nw194).ArraySet1(_1152_v27, 8)
				(_nw194).ArraySet1(_1152_v27, 9)
				(_nw194).ArraySet1(_1152_v27, 10)
				(_nw194).ArraySet1(_1152_v27, 11)
				(_nw194).ArraySet1(_1152_v27, 12)
				(_nw194).ArraySet1(_1152_v27, 13)
				(_nw194).ArraySet1(_1152_v27, 14)
				(_nw194).ArraySet1((func() _dafny.Array {
					if !((_this).F29()) {
						return _1152_v27
					}
					return _1152_v27
				})(), 15)
				(_nw194).ArraySet1(_1152_v27, 16)
				(_nw194).ArraySet1(_1152_v27, 17)
				(_nw194).ArraySet1(_1152_v27, 18)
				(_nw194).ArraySet1(_1152_v27, 19)
				(_nw194).ArraySet1(_1152_v27, 20)
				(_nw194).ArraySet1(_1152_v27, 21)
				(_nw194).ArraySet1(_1152_v27, 22)
				(_nw194).ArraySet1(_1152_v27, 23)
				(_nw194).ArraySet1((_1153_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.IntOfUint32((_1153_v28).Cardinality()))).Uint32()).(_dafny.Array), 24)
				(_nw194).ArraySet1(_1152_v27, 25)
				(_nw194).ArraySet1(_1152_v27, 26)
				(_nw194).ArraySet1(_1152_v27, 27)
				_1154_v29 = _nw194
				var _1155_v30 D12
				_ = _1155_v30
				_1155_v30 = Companion_D12_.Create_DC38_(_1152_v27)
				var _1156_v31 _dafny.Map
				_ = _1156_v31
				_1156_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1151_v26).F35(), (_1155_v30).Dtor_cf64())
				var _nwElement0_37 _dafny.Array = _1152_v27
				_ = _nwElement0_37
				var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(23))
				_ = _nw195
				(_nw195).ArraySet1(_nwElement0_37, 0)
				(_nw195).ArraySet1(_1152_v27, 1)
				(_nw195).ArraySet1(_1152_v27, 2)
				(_nw195).ArraySet1(_1152_v27, 3)
				(_nw195).ArraySet1(_1152_v27, 4)
				(_nw195).ArraySet1(_1152_v27, 5)
				(_nw195).ArraySet1(_1152_v27, 6)
				(_nw195).ArraySet1(_1152_v27, 7)
				(_nw195).ArraySet1(_1152_v27, 8)
				(_nw195).ArraySet1(_1152_v27, 9)
				(_nw195).ArraySet1(_1152_v27, 10)
				(_nw195).ArraySet1((func() _dafny.Array {
					if (_1156_v31).Contains(_1148_cf6) {
						return (_1156_v31).Get(_1148_cf6).(_dafny.Array)
					}
					return _1152_v27
				})(), 11)
				(_nw195).ArraySet1(_1152_v27, 12)
				(_nw195).ArraySet1(_1152_v27, 13)
				(_nw195).ArraySet1((_1153_v28).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1153_v28).Cardinality()))).Uint32()).(_dafny.Array), 14)
				(_nw195).ArraySet1((_1155_v30).Dtor_cf64(), 15)
				(_nw195).ArraySet1(_1152_v27, 16)
				(_nw195).ArraySet1(_1152_v27, 17)
				(_nw195).ArraySet1(_1152_v27, 18)
				(_nw195).ArraySet1(_1152_v27, 19)
				(_nw195).ArraySet1(_1152_v27, 20)
				(_nw195).ArraySet1(_1152_v27, 21)
				(_nw195).ArraySet1(_1152_v27, 22)
				_1154_v29 = _nw195
			}
			(globalState).F1 = p2
		} else if _source18.Is_DC19() {
			var _1157___mcc_h5 _dafny.Int = _source18.Get_().(D5_DC19).Cf28
			_ = _1157___mcc_h5
			var _1158___mcc_h6 _dafny.Int = _source18.Get_().(D5_DC19).Cf29
			_ = _1158___mcc_h6
			var _1159___mcc_h7 _dafny.CodePoint = _source18.Get_().(D5_DC19).Cf30
			_ = _1159___mcc_h7
			var _1160_cf30 _dafny.CodePoint = _1159___mcc_h7
			_ = _1160_cf30
			var _1161_cf29 _dafny.Int = _1158___mcc_h6
			_ = _1161_cf29
			var _1162_cf28 _dafny.Int = _1157___mcc_h5
			_ = _1162_cf28
			(globalState).F15 = (p0) || ((_1161_cf29).Cmp(_dafny.IntOfUint32((_this.F30()).Cardinality())) <= 0)
			var _1163_v32 _dafny.Array
			_ = _1163_v32
			var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
			_ = _nw196
			_1163_v32 = _nw196
			var _1164_v33 D2
			_ = _1164_v33
			_1164_v33 = Companion_D2_.Create_DC7_((_this).F29())
			var _1165_v34 _dafny.Map
			_ = _1165_v34
			_1165_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1164_v33).Dtor_cf11())
			var _1166_v35 _dafny.Map
			_ = _1166_v35
			_1166_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1165_v34)
			var _1167_v36 _dafny.Map
			_ = _1167_v36
			_1167_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1166_v35)
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1163_v32), 0))
			_ = _index187
			(_1163_v32).ArraySet1(((func() _dafny.Map {
				if (_1167_v36).Contains(false) {
					return (_1167_v36).Get(false).(_dafny.Map)
				}
				return Companion_Default___.Fm46(_dafny.CodePoint('k'), _this.F30(), globalState)
			})()).Update(!((_this).F29()), _1165_v34), (_index187).Int())
			var _1168_v37 _dafny.MultiSet
			_ = _1168_v37
			_1168_v37 = _dafny.MultiSetOf((_this).F29(), (_this).F29())
			var _1169_v38 _dafny.Sequence
			_ = _1169_v38
			_1169_v38 = _dafny.SeqOf((_1168_v37).Cardinality())
			var _1170_v39 _dafny.MultiSet
			_ = _1170_v39
			_1170_v39 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1162_cf28, p0), _1165_v34, _1165_v34, _1165_v34, _1165_v34)
			var _1171_v40 _dafny.Set
			_ = _1171_v40
			_1171_v40 = _dafny.SetOf(_dafny.IntOfInt64(197))
			var _1172_v41 _dafny.Array
			_ = _1172_v41
			var _nwElement0_38 _dafny.Int = (_1169_v38).Select((Companion_Default___.SafeIndex(_1162_cf28, _dafny.IntOfUint32((_1169_v38).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _nwElement0_38
			var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(9))
			_ = _nw197
			(_nw197).ArraySet1(_nwElement0_38, 0)
			(_nw197).ArraySet1(_1161_cf29, 1)
			(_nw197).ArraySet1(p2, 2)
			(_nw197).ArraySet1(_1162_cf28, 3)
			(_nw197).ArraySet1((_this).Fm4(globalState), 4)
			(_nw197).ArraySet1((func() _dafny.Int {
				if (_1170_v39).Contains(Companion_Default___.Fm41(_dafny.IntOfInt64(18), p0, p0, !((_this).F29()), globalState)) {
					return (_1170_v39).Multiplicity(Companion_Default___.Fm41(_dafny.IntOfInt64(18), p0, p0, !((_this).F29()), globalState))
				}
				return (_1171_v40).Cardinality()
			})(), 5)
			(_nw197).ArraySet1(_1161_cf29, 6)
			(_nw197).ArraySet1(p2, 7)
			(_nw197).ArraySet1(_1161_cf29, 8)
			_1172_v41 = _nw197
			var _1173_v42 _dafny.Sequence
			_ = _1173_v42
			_1173_v42 = _dafny.SeqOf(_1172_v41, _1172_v41)
			var _1174_v43 _dafny.Map
			_ = _1174_v43
			_1174_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1173_v42, _dafny.IntOfUint32((_this.F30()).Cardinality()))
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1163_v32), 0))
			_ = _index188
			var _rhs196 _dafny.Map = _1166_v35
			_ = _rhs196
			var _rhs197 _dafny.Int = _1161_cf29
			_ = _rhs197
			var _rhs198 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())
			_ = _rhs198
			var _rhs199 _dafny.Int = (func() _dafny.Int {
				if (_1174_v43).Contains(_1173_v42) {
					return (_1174_v43).Get(_1173_v42).(_dafny.Int)
				}
				return Companion_Default___.SafeDivisionInt(p2, _1162_cf28)
			})()
			_ = _rhs199
			var _lhs187 _dafny.Array = _1163_v32
			_ = _lhs187
			var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1163_v32), 0))
			_ = _lhs188
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			var _lhs191 *GlobalState = globalState
			_ = _lhs191
			(_lhs187).ArraySet1(_rhs196, (_lhs188).Int())
			_lhs189.F14 = _rhs197
			_lhs190.F11 = _rhs198
			_lhs191.F21 = _rhs199
			var _1175_v44 *C4
			_ = _1175_v44
			var _nw198 *C4 = New_C4_()
			_ = _nw198
			_nw198.Ctor__(p0, p2)
			_1175_v44 = _nw198
			(globalState).F21 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}((func(_1176_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1177_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_1176_v5, _1176_v5)
				}
			})(_1105_v5)))).Cardinality()))
		} else if _source18.Is_DC16() {
			var _1178___mcc_h8 _dafny.Map = _source18.Get_().(D5_DC16).Cf26
			_ = _1178___mcc_h8
			var _1179_cf26 _dafny.Map = _1178___mcc_h8
			_ = _1179_cf26
			var _1180_v45 _dafny.CodePoint
			_ = _1180_v45
			_1180_v45 = _dafny.CodePoint('c')
			var _1181_v46 _dafny.Array
			_ = _1181_v46
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_32
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Int = (func(_1182_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1183_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1183_i3, _1182_p2)
					}
				})(p2)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw199 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw199).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw199).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1181_v46 = _nw199
			var _1184_v47 _dafny.Map
			_ = _1184_v47
			_1184_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v45, _1181_v46)
			var _1185_v48 _dafny.Array
			_ = _1185_v48
			var _nwElement0_39 _dafny.Map = _1184_v47
			_ = _nwElement0_39
			var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(16))
			_ = _nw200
			(_nw200).ArraySet1(_nwElement0_39, 0)
			(_nw200).ArraySet1(_1184_v47, 1)
			(_nw200).ArraySet1(_1184_v47, 2)
			(_nw200).ArraySet1(_1184_v47, 3)
			(_nw200).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v45, _1181_v46)).Update(Companion_Default___.Fm42((_this).F29(), p0, globalState), _1181_v46), 4)
			(_nw200).ArraySet1(_1184_v47, 5)
			(_nw200).ArraySet1((_1184_v47).Update(_1180_v45, _1181_v46), 6)
			(_nw200).ArraySet1(_1184_v47, 7)
			(_nw200).ArraySet1(_1184_v47, 8)
			(_nw200).ArraySet1((_1184_v47).Update(_1180_v45, _1181_v46), 9)
			(_nw200).ArraySet1(_1184_v47, 10)
			(_nw200).ArraySet1((_1184_v47).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v45, _1181_v46)), 11)
			(_nw200).ArraySet1(_1184_v47, 12)
			(_nw200).ArraySet1(_1184_v47, 13)
			(_nw200).ArraySet1(_1184_v47, 14)
			(_nw200).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v45, _1181_v46), 15)
			_1185_v48 = _nw200
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_1185_v48), 0))
			_ = _index189
			(_1185_v48).ArraySet1(_1184_v47, (_index189).Int())
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_1185_v48), 0))
			_ = _index190
			(_1185_v48).ArraySet1((_1184_v47).Update(_dafny.CodePoint('d'), _1181_v46), (_index190).Int())
			(globalState).F15 = (_this).F29()
			(globalState).F18 = p0
			(globalState).F18 = (_this).F29()
		} else {
			var _1186___mcc_h9 D5 = _source18.Get_().(D5_DC20).Cf31
			_ = _1186___mcc_h9
			var _1187_cf31 D5 = _1186___mcc_h9
			_ = _1187_cf31
			(globalState).F14 = p2
			var _1188_v49 _dafny.Array
			_ = _1188_v49
			var _nw201 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw201
			_1188_v49 = _nw201
			var _1189_v50 _dafny.Sequence
			_ = _1189_v50
			_1189_v50 = _dafny.SeqOf(_dafny.IntOfInt64(186), p2, _dafny.IntOfUint32((_1105_v5).Cardinality()))
			var _1190_v51 D0
			_ = _1190_v51
			_1190_v51 = Companion_D0_.Create_DC1_(_1188_v49, _dafny.IntOfUint32((_1105_v5).Cardinality()), (_1189_v50).Select((Companion_Default___.SafeIndex((_this).Fm4(globalState), _dafny.IntOfUint32((_1189_v50).Cardinality()))).Uint32()).(_dafny.Int))
			var _1191_v52 _dafny.Sequence
			_ = _1191_v52
			_1191_v52 = _dafny.SeqOf(_1188_v49)
			var _1192_v53 _dafny.Array
			_ = _1192_v53
			var _nwElement0_40 _dafny.Array = _1188_v49
			_ = _nwElement0_40
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(29))
			_ = _nw202
			(_nw202).ArraySet1(_nwElement0_40, 0)
			(_nw202).ArraySet1(_1188_v49, 1)
			(_nw202).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1188_v49
				}
				return _1188_v49
			})(), 2)
			(_nw202).ArraySet1(_1188_v49, 3)
			(_nw202).ArraySet1(_1188_v49, 4)
			(_nw202).ArraySet1(_1188_v49, 5)
			(_nw202).ArraySet1(_1188_v49, 6)
			(_nw202).ArraySet1(_1188_v49, 7)
			(_nw202).ArraySet1((_1190_v51).Dtor_cf1(), 8)
			(_nw202).ArraySet1(_1188_v49, 9)
			(_nw202).ArraySet1(_1188_v49, 10)
			(_nw202).ArraySet1((_1191_v52).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1191_v52).Cardinality()))).Uint32()).(_dafny.Array), 11)
			(_nw202).ArraySet1(_1188_v49, 12)
			(_nw202).ArraySet1(_1188_v49, 13)
			(_nw202).ArraySet1(_1188_v49, 14)
			(_nw202).ArraySet1(_1188_v49, 15)
			(_nw202).ArraySet1(_1188_v49, 16)
			(_nw202).ArraySet1((func() _dafny.Array {
				if (_this).F29() {
					return _1188_v49
				}
				return _1188_v49
			})(), 17)
			(_nw202).ArraySet1(_1188_v49, 18)
			(_nw202).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1188_v49
				}
				return _1188_v49
			})(), 19)
			(_nw202).ArraySet1(_1188_v49, 20)
			(_nw202).ArraySet1(_1188_v49, 21)
			(_nw202).ArraySet1(_1188_v49, 22)
			(_nw202).ArraySet1(_1188_v49, 23)
			(_nw202).ArraySet1(_1188_v49, 24)
			(_nw202).ArraySet1(_1188_v49, 25)
			(_nw202).ArraySet1(_1188_v49, 26)
			(_nw202).ArraySet1(_1188_v49, 27)
			(_nw202).ArraySet1(_1188_v49, 28)
			_1192_v53 = _nw202
			_1192_v53 = (func() _dafny.Array {
				if (_this).F29() {
					return _1192_v53
				}
				return _1192_v53
			})()
			var _1193_v54 T0
			_ = _1193_v54
			var _nw203 *C1 = New_C1_()
			_ = _nw203
			_nw203.Ctor__((_this).F29())
			_1193_v54 = _nw203
			var _1194_v55 _dafny.Map
			_ = _1194_v55
			_1194_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1193_v54, (_this).F29())
			var _1195_v56 _dafny.Map
			_ = _1195_v56
			_1195_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((_1194_v55).Cardinality()).Cmp(p2) < 0)
			_1195_v56 = func() _dafny.Map {
				var _coll46 = _dafny.NewMapBuilder()
				_ = _coll46
				for _iter54 := _dafny.Iterate((_1189_v50).Elements()); ; {
					_compr_46, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1196_v57 _dafny.Int
					_1196_v57 = interface{}(_compr_46).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1189_v50, _1196_v57) {
						_coll46.Add((_1196_v57).Minus(p2), p0)
					}
				}
				return _coll46.ToMap()
			}()
			var _1197_v58 _dafny.Sequence
			_ = _1197_v58
			var _1198_v59 bool
			_ = _1198_v59
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 bool
			_ = _out17
			_out16, _out17 = (_1193_v54).M2(p0, globalState)
			_1197_v58 = _out16
			_1198_v59 = _out17
		}
		var _1199_v60 _dafny.Map
		_ = _1199_v60
		_1199_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _dafny.SetOf((_this).F29()))
		if !((p2).Cmp((_1199_v60).Cardinality()) < 0) || (p0) {
			(globalState).F21 = _dafny.IntOfInt64(744)
			var _1200_v61 _dafny.Map
			_ = _1200_v61
			_1200_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			(globalState).F28 = ((_1200_v61).Cardinality()).Cmp(p2) >= 0
			(globalState).F18 = (p2).Cmp((p2).Times(_dafny.IntOfUint32((_this.F30()).Cardinality()))) >= 0
			(globalState).F19 = (func() _dafny.Int {
				if p0 {
					return Companion_Default___.SafeModuloInt(p2, p2)
				}
				return (_dafny.SetOf(p2)).Cardinality()
			})()
			(globalState).F28 = p0
		} else {
			var _1201_v62 T0
			_ = _1201_v62
			var _nw204 *C4 = New_C4_()
			_ = _nw204
			_nw204.Ctor__((_this).F29(), p2)
			_1201_v62 = _nw204
			var _1202_v63 D10
			_ = _1202_v63
			_1202_v63 = Companion_D10_.Create_DC30_(_1201_v62)
			var _source20 D10 = _1202_v63
			_ = _source20
			if _source20.Is_DC31() {
				var _1203___mcc_h17 bool = _source20.Get_().(D10_DC31).Cf50
				_ = _1203___mcc_h17
				var _1204___mcc_h18 bool = _source20.Get_().(D10_DC31).Cf51
				_ = _1204___mcc_h18
				var _1205_cf51 bool = _1204___mcc_h18
				_ = _1205_cf51
				var _1206_cf50 bool = _1203___mcc_h17
				_ = _1206_cf50
				var _1207_v64 D9
				_ = _1207_v64
				_1207_v64 = Companion_D9_.Create_DC29_(p2, p2, _1205_cf51)
				_1207_v64 = _1207_v64
				var _1208_v65 *C3
				_ = _1208_v65
				var _nw205 *C3 = New_C3_()
				_ = _nw205
				_nw205.Ctor__(_1205_cf51, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(476), p2), (_this).F29(), _this.F30())
				_1208_v65 = _nw205
				var _1209_v66 _dafny.Map
				_ = _1209_v66
				_1209_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1206_cf50, (_1208_v65).F35())
				var _rhs200 _dafny.Map = (_1209_v66).Merge((_1209_v66).Merge(_1209_v66))
				_ = _rhs200
				var _rhs201 _dafny.Int = p2
				_ = _rhs201
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				_1209_v66 = _rhs200
				_lhs192.F21 = _rhs201
				var _1210_v67 _dafny.Array
				_ = _1210_v67
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw206
				_1210_v67 = _nw206
				var _1211_v68 _dafny.Map
				_ = _1211_v68
				_1211_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1210_v67, true)
				var _1212_v69 _dafny.CodePoint
				_ = _1212_v69
				_1212_v69 = _dafny.CodePoint('n')
				var _1213_v70 D2
				_ = _1213_v70
				_1213_v70 = Companion_D2_.Create_DC8_(_1212_v69, _1206_cf50)
				var _rhs202 _dafny.Map = _1211_v68
				_ = _rhs202
				var _rhs203 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm35(_dafny.SeqOf(_1213_v70), p2, globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1208_v65).F35(), _1205_cf51), _1105_v5))
				_ = _rhs203
				var _rhs204 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _1212_v69), _this.F30()), _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()))
				_ = _rhs204
				var _rhs205 _dafny.Sequence = _1105_v5
				_ = _rhs205
				var _rhs206 bool = (_1206_cf50) || (p0)
				_ = _rhs206
				var _lhs193 *C5 = _this
				_ = _lhs193
				var _lhs194 *GlobalState = globalState
				_ = _lhs194
				_1211_v68 = _rhs202
				_1105_v5 = _rhs203
				_lhs193.F30_set_(_rhs204)
				_1105_v5 = _rhs205
				_lhs194.F13 = _rhs206
			} else if _source20.Is_DC32() {
				var _1214___mcc_h19 _dafny.CodePoint = _source20.Get_().(D10_DC32).Cf52
				_ = _1214___mcc_h19
				var _1215___mcc_h20 bool = _source20.Get_().(D10_DC32).Cf53
				_ = _1215___mcc_h20
				var _1216_cf53 bool = _1215___mcc_h20
				_ = _1216_cf53
				var _1217_cf52 _dafny.CodePoint = _1214___mcc_h19
				_ = _1217_cf52
				var _1218_v71 *C0
				_ = _1218_v71
				var _nw207 *C0 = New_C0_()
				_ = _nw207
				_nw207.Ctor__((_this).F29())
				_1218_v71 = _nw207
				var _1219_v72 _dafny.Array
				_ = _1219_v72
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw208
				_1219_v72 = _nw208
				_1219_v72 = _1219_v72
				(globalState).F19 = p2
				var _1220_v73 D9
				_ = _1220_v73
				_1220_v73 = Companion_D9_.Create_DC28_(_dafny.SeqOf(p0, !((_this).F29())), _1216_cf53, p2, _1216_cf53, _1218_v71.F37)
				var _1221_v74 _dafny.Map
				_ = _1221_v74
				_1221_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.Fm42(_1218_v71.F37, (_1220_v73).Dtor_cf45(), globalState))
				_1221_v74 = (_1221_v74).Update((p2).Times(p2), _1217_cf52)
			} else if _source20.Is_DC33() {
				var _1222___mcc_h21 bool = _source20.Get_().(D10_DC33).Cf54
				_ = _1222___mcc_h21
				var _1223___mcc_h22 _dafny.Int = _source20.Get_().(D10_DC33).Cf55
				_ = _1223___mcc_h22
				var _1224_cf55 _dafny.Int = _1223___mcc_h22
				_ = _1224_cf55
				var _1225_cf54 bool = _1222___mcc_h21
				_ = _1225_cf54
				var _1226_v75 *C1
				_ = _1226_v75
				var _nw209 *C1 = New_C1_()
				_ = _nw209
				_nw209.Ctor__(p0)
				_1226_v75 = _nw209
				var _1227_v76 *C1
				_ = _1227_v76
				var _nw210 *C1 = New_C1_()
				_ = _nw210
				_nw210.Ctor__((_this).F29())
				_1227_v76 = _nw210
				var _1228_v77 _dafny.Array
				_ = _1228_v77
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw211
				_1228_v77 = _nw211
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1228_v77), 0))
				_ = _index191
				(_1228_v77).ArraySet1(p2, (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1228_v77), 0))
				_ = _index192
				(_1228_v77).ArraySet1(p2, (_index192).Int())
				(globalState).F1 = (_1228_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1228_v77), 0))).Int()).(_dafny.Int)
			} else {
				var _1229___mcc_h23 T0 = _source20.Get_().(D10_DC30).Cf49
				_ = _1229___mcc_h23
				var _1230_cf49 T0 = _1229___mcc_h23
				_ = _1230_cf49
				var _1231_v78 _dafny.MultiSet
				_ = _1231_v78
				_1231_v78 = _dafny.MultiSetOf(p0)
				_1231_v78 = _1231_v78
				var _1232_v79 _dafny.Array
				_ = _1232_v79
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_33
				var _nw212 _dafny.Array
				_ = _nw212
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw212 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1233_p2 _dafny.Int, _1234_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1235_i4 _dafny.Int) _dafny.Int {
							return (_1235_i4).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(_1233_p2), _dafny.MultiSetFromSeq(_dafny.SeqOf(_1233_p2, _1233_p2, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_1234_v5).Cardinality()), _1233_p2, _1233_p2, _dafny.IntOfInt64(962)))).Cardinality())))).Cardinality()))
						}
					})(p2, _1105_v5)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw212 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw212).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw212).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1232_v79 = _nw212
				(globalState).F17 = _1232_v79
				(globalState).F15 = (_this).F29()
				var _1236_v80 _dafny.CodePoint
				_ = _1236_v80
				_1236_v80 = _dafny.CodePoint('c')
				_1236_v80 = _1236_v80
			}
			var _1237_v81 _dafny.Array
			_ = _1237_v81
			var _nw213 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw213
			_1237_v81 = _nw213
			var _1238_v82 _dafny.Array
			_ = _1238_v82
			var _nwElement0_41 _dafny.Array = _1237_v81
			_ = _nwElement0_41
			var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(5))
			_ = _nw214
			(_nw214).ArraySet1(_nwElement0_41, 0)
			(_nw214).ArraySet1(_1237_v81, 1)
			(_nw214).ArraySet1(_1237_v81, 2)
			(_nw214).ArraySet1(_1237_v81, 3)
			(_nw214).ArraySet1(_1237_v81, 4)
			_1238_v82 = _nw214
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1238_v82), 0))
			_ = _index193
			(_1238_v82).ArraySet1(_1237_v81, (_index193).Int())
			var _1239_v83 _dafny.Map
			_ = _1239_v83
			_1239_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v81, Companion_Default___.Fm10(p2, p2, globalState))
			var _1240_v84 _dafny.Sequence
			_ = _1240_v84
			_1240_v84 = _dafny.SeqOf((_1239_v83).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality())))
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1238_v82), 0))
			_ = _index194
			var _rhs207 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1240_v84, _dafny.Companion_Sequence_.Concatenate(_1240_v84, _1240_v84))).Cardinality())
			_ = _rhs207
			var _rhs208 _dafny.Int = p2
			_ = _rhs208
			var _rhs209 _dafny.Array = _1237_v81
			_ = _rhs209
			var _rhs210 bool = (_this).F29()
			_ = _rhs210
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			var _lhs196 *GlobalState = globalState
			_ = _lhs196
			var _lhs197 _dafny.Array = _1238_v82
			_ = _lhs197
			var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1238_v82), 0))
			_ = _lhs198
			var _lhs199 *GlobalState = globalState
			_ = _lhs199
			_lhs195.F19 = _rhs207
			_lhs196.F21 = _rhs208
			(_lhs197).ArraySet1(_rhs209, (_lhs198).Int())
			_lhs199.F15 = _rhs210
			var _1241_v85 _dafny.MultiSet
			_ = _1241_v85
			_1241_v85 = _dafny.MultiSetOf((_this).F29(), p0)
			var _1242_v86 _dafny.Set
			_ = _1242_v86
			_1242_v86 = _dafny.SetOf(p2, p2, _dafny.IntOfUint32((_dafny.SeqOf((_1241_v85).Cardinality())).Cardinality()), p2, p2)
			var _1243_v87 D8
			_ = _1243_v87
			_1243_v87 = Companion_D8_.Create_DC26_()
			var _1244_v88 _dafny.Map
			_ = _1244_v88
			_1244_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v86, _1243_v87)
			_1244_v88 = (_1244_v88).Update(_1242_v86, _1243_v87)
			var _1245_v89 _dafny.Array
			_ = _1245_v89
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(Companion_D10_.Default(), _dafny.IntOfInt64(14))
			_ = _nw215
			_1245_v89 = _nw215
			var _1246_v90 _dafny.Map
			_ = _1246_v90
			_1246_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm34(p2, p2, p0, true, globalState), _1245_v89)
			var _1247_v91 _dafny.Map
			_ = _1247_v91
			_1247_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _rhs211 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Array {
				if (_1246_v90).Contains(_dafny.UnicodeSeqOfUtf8Bytes("strdxawaa")) {
					return (_1246_v90).Get(_dafny.UnicodeSeqOfUtf8Bytes("strdxawaa")).(_dafny.Array)
				}
				return _1245_v89
			})(), _1245_v89, _1245_v89, _1245_v89, _1245_v89)).Cardinality())
			_ = _rhs211
			var _rhs212 bool = (_this).F29()
			_ = _rhs212
			var _rhs213 bool = !(false) || ((p2).Cmp(p2) == 0)
			_ = _rhs213
			var _rhs214 _dafny.Int = (_dafny.Zero).Minus(p2)
			_ = _rhs214
			var _rhs215 _dafny.Int = (p2).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_this.F30()).Cardinality()), p2, (_1247_v91).Cardinality())).Cardinality()))
			_ = _rhs215
			var _lhs200 *GlobalState = globalState
			_ = _lhs200
			var _lhs201 *GlobalState = globalState
			_ = _lhs201
			var _lhs202 *GlobalState = globalState
			_ = _lhs202
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 *GlobalState = globalState
			_ = _lhs204
			_lhs200.F2 = _rhs211
			_lhs201.F28 = _rhs212
			_lhs202.F15 = _rhs213
			_lhs203.F19 = _rhs214
			_lhs204.F2 = _rhs215
			var _1248_v92 D3
			_ = _1248_v92
			_1248_v92 = Companion_D3_.Create_DC11_(p0, p0)
			var _1249_v93 D3
			_ = _1249_v93
			_1249_v93 = Companion_D3_.Create_DC12_(_1248_v92)
			var _1250_v94 _dafny.Set
			_ = _1250_v94
			_1250_v94 = _dafny.SetOf(_1249_v93)
			var _1251_v95 _dafny.Set
			_ = _1251_v95
			_1251_v95 = _dafny.SetOf(_1250_v94, _1250_v94)
			(globalState).F18 = (_1251_v95).IsSubsetOf((_dafny.SetOf(_1250_v94)).Union(_1251_v95))
		}
		var _1252_i5 _dafny.Int
		_ = _1252_i5
		_1252_i5 = _dafny.Zero
		{
			for (_this).F29() {
				{
					if (_1252_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1252_i5 = (_1252_i5).Plus(_dafny.One)
					var _1253_v96 _dafny.Map
					_ = _1253_v96
					_1253_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p0)
					var _1254_v97 D11
					_ = _1254_v97
					_1254_v97 = Companion_D11_.Create_DC37_(p2, p0)
					var _1255_v98 _dafny.Array
					_ = _1255_v98
					var _nwElement0_42 bool = p0
					_ = _nwElement0_42
					var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(14))
					_ = _nw216
					(_nw216).ArraySet1(_nwElement0_42, 0)
					(_nw216).ArraySet1((_this).F29(), 1)
					(_nw216).ArraySet1(!(!(p0)), 2)
					(_nw216).ArraySet1(p0, 3)
					(_nw216).ArraySet1(p0, 4)
					(_nw216).ArraySet1(p0, 5)
					(_nw216).ArraySet1(true, 6)
					(_nw216).ArraySet1(Companion_Default___.Fm33(p0, p2, p2, globalState), 7)
					(_nw216).ArraySet1((_this).F29(), 8)
					(_nw216).ArraySet1((_this).F29(), 9)
					(_nw216).ArraySet1((func() bool {
						if (_1253_v96).Contains(p0) {
							return (_1253_v96).Get(p0).(bool)
						}
						return (_1105_v5).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1105_v5).Cardinality()))).Uint32()).(bool)
					})(), 10)
					(_nw216).ArraySet1((_1254_v97).Dtor_cf63(), 11)
					(_nw216).ArraySet1((_this).F29(), 12)
					(_nw216).ArraySet1(p0, 13)
					_1255_v98 = _nw216
					var _1256_v99 _dafny.Map
					_ = _1256_v99
					_1256_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1255_v98)
					_1256_v99 = (_1256_v99).Update(p2, _1255_v98)
					var _1257_v100 *C4
					_ = _1257_v100
					var _nw217 *C4 = New_C4_()
					_ = _nw217
					_nw217.Ctor__(true, (_dafny.IntOfInt64(-450)).Times(p2))
					_1257_v100 = _nw217
					(globalState).F15 = !((_1257_v100).F33()) || ((_this).F29())
					var _1258_v101 _dafny.CodePoint
					_ = _1258_v101
					_1258_v101 = _dafny.CodePoint('y')
					var _1259_v102 D5
					_ = _1259_v102
					_1259_v102 = Companion_D5_.Create_DC19_(_dafny.IntOfInt64(453), (_dafny.MultiSetFromSeq(_1105_v5)).Cardinality(), _1258_v101)
					_1258_v101 = (_1259_v102).Dtor_cf30()
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1260_v103 D9
		_ = _1260_v103
		_1260_v103 = Companion_D9_.Create_DC29_(p2, p2, (_this).F29())
		(globalState).F21 = (_1260_v103).Dtor_cf46()
		if p0 {
			var _1261_v104 _dafny.MultiSet
			_ = _1261_v104
			_1261_v104 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(p2, p2))
			var _1262_v105 _dafny.Map
			_ = _1262_v105
			_1262_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1261_v104)
			var _1263_v106 _dafny.Map
			_ = _1263_v106
			_1263_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _1264_v107 _dafny.Map
			_ = _1264_v107
			_1264_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1263_v106).Cardinality())
			_1261_v104 = (func() _dafny.MultiSet {
				if (_1262_v105).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1264_v107, !((_this).F29()))).Cardinality()) {
					return (_1262_v105).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1264_v107, !((_this).F29()))).Cardinality()).(_dafny.MultiSet)
				}
				return (_1261_v104).Update(p2, Companion_Default___.Abs(_dafny.IntOfUint32((_this.F30()).Cardinality())))
			})()
			var _1265_v108 _dafny.CodePoint
			_ = _1265_v108
			_1265_v108 = _dafny.CodePoint('e')
			_1265_v108 = _1265_v108
			var _1266_v109 _dafny.Array
			_ = _1266_v109
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw218
			_1266_v109 = _nw218
			var _1267_v110 _dafny.Map
			_ = _1267_v110
			_1267_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F29())
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1266_v109), 0))
			_ = _index195
			(_1266_v109).ArraySet1(_1267_v110, (_index195).Int())
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1266_v109), 0))
			_ = _index196
			(_1266_v109).ArraySet1(_1267_v110, (_index196).Int())
			(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _1265_v108)), _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()))
			var _1268_v111 _dafny.Sequence
			_ = _1268_v111
			_1268_v111 = _dafny.SeqOf(_this.F30())
			var _1269_v112 _dafny.MultiSet
			_ = _1269_v112
			_1269_v112 = _dafny.MultiSetOf(_this.F30(), (_1268_v111).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-182), _dafny.IntOfUint32((_1268_v111).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1270_v113 _dafny.Sequence
			_ = _1270_v113
			_1270_v113 = _dafny.SeqOf(_1269_v112)
			(globalState).F18 = ((_1270_v113).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1270_v113).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf((_1269_v112).Difference(_1269_v112))
		} else {
			var _1271_v114 _dafny.Array
			_ = _1271_v114
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_34
			var _nw219 _dafny.Array
			_ = _nw219
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw219 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Int = (func(_1272_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1273_i6 _dafny.Int) _dafny.Int {
						return (_1273_i6).Times(_1272_p2)
					}
				})(p2)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw219 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw219).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw219).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1271_v114 = _nw219
			var _1274_v115 _dafny.Set
			_ = _1274_v115
			_1274_v115 = _dafny.SetOf(_1271_v114)
			var _1275_v116 *C2
			_ = _1275_v116
			var _nw220 *C2 = New_C2_()
			_ = _nw220
			_nw220.Ctor__(p2, _1274_v115, (_this).F29(), _this.F30())
			_1275_v116 = _nw220
			var _1276_v117 _dafny.Array
			_ = _1276_v117
			var _nw221 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw221
			_1276_v117 = _nw221
			var _1277_v118 _dafny.Sequence
			_ = _1277_v118
			_1277_v118 = _dafny.SeqOf(_1276_v117)
			var _1278_v119 _dafny.Map
			_ = _1278_v119
			_1278_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1275_v116, _1277_v118)
			var _1279_v120 _dafny.Map
			_ = _1279_v120
			_1279_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_1278_v119).Contains(_1275_v116) {
					return (_1278_v119).Get(_1275_v116).(_dafny.Sequence)
				}
				return _1277_v118
			})(), (_1275_v116).F38())
			_1279_v120 = (_1279_v120).Update(_1277_v118, _dafny.IntOfInt64(80))
			var _1280_v121 _dafny.CodePoint
			_ = _1280_v121
			_1280_v121 = _dafny.CodePoint('v')
			var _1281_v122 D1
			_ = _1281_v122
			_1281_v122 = Companion_D1_.Create_DC4_((_this).F29())
			var _1282_v123 _dafny.Map
			_ = _1282_v123
			_1282_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F30()).Cardinality()), Companion_D11_.Create_DC35_(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex((_1275_v116).F38(), _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _1280_v121), (_1281_v122).Dtor_cf8()))
			var _1283_v124 _dafny.Set
			_ = _1283_v124
			_1283_v124 = _dafny.SetOf(_1282_v123)
			var _1284_v125 _dafny.MultiSet
			_ = _1284_v125
			_1284_v125 = _dafny.MultiSetOf(p0)
			var _1285_v126 _dafny.Set
			_ = _1285_v126
			_1285_v126 = _dafny.SetOf((_dafny.MultiSetOf((_this).F29())).Equals(_1284_v125))
			var _1286_v127 D11
			_ = _1286_v127
			_1286_v127 = Companion_D11_.Create_DC37_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg87 _dafny.Int) interface{} {
					return coer87(arg87)
				}
			}((func(_1287_v121 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1288_i7 _dafny.Int) _dafny.CodePoint {
					return _1287_v121
				}
			})(_1280_v121)))).Cardinality()), (_this).F29())
			var _1289_v128 _dafny.Sequence
			_ = _1289_v128
			_1289_v128 = _dafny.SeqOf((_1275_v116).F38(), (_1275_v116).F38())
			var _1290_v129 _dafny.Map
			_ = _1290_v129
			_1290_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v128, _this.F30())
			var _rhs216 bool = true
			_ = _rhs216
			var _rhs217 _dafny.Int = (_1286_v127).Dtor_cf62()
			_ = _rhs217
			var _rhs218 _dafny.Set = Companion_Default___.Fm47(p0, (_this).F29(), globalState)
			_ = _rhs218
			var _rhs219 _dafny.Set = _1285_v126
			_ = _rhs219
			var _rhs220 bool = _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
				if (_this).F29() {
					return _this.F30()
				}
				return (func() _dafny.Sequence {
					if (_1290_v129).Contains(_1289_v128) {
						return (_1290_v129).Get(_1289_v128).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("ysikfqdgi")
				})()
			})(), _1280_v121)
			_ = _rhs220
			var _lhs205 *GlobalState = globalState
			_ = _lhs205
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			var _lhs207 *GlobalState = globalState
			_ = _lhs207
			_lhs205.F18 = _rhs216
			_lhs206.F2 = _rhs217
			_1283_v124 = _rhs218
			_1285_v126 = _rhs219
			_lhs207.F18 = _rhs220
			var _1291_v130 _dafny.Array
			_ = _1291_v130
			var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw222
			_1291_v130 = _nw222
			_1291_v130 = _1291_v130
			var _1292_v131 *C0
			_ = _1292_v131
			var _nw223 *C0 = New_C0_()
			_ = _nw223
			_nw223.Ctor__(p0)
			_1292_v131 = _nw223
			var _1293_v132 _dafny.Sequence
			_ = _1293_v132
			_1293_v132 = _dafny.SeqOf(_1292_v131)
			(globalState).F21 = ((_1275_v116).F38()).Plus((_dafny.IntOfUint32((_1105_v5).Cardinality())).Minus(_dafny.IntOfUint32((_1293_v132).Cardinality())))
			_1285_v126 = _1285_v126
		}
		r0 = _this.F30()
		var _1294_v133 _dafny.Map
		_ = _1294_v133
		_1294_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1105_v5)
		r1 = (func() bool {
			if (_this).F29() {
				return _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
					if (_1294_v133).Contains(false) {
						return (_1294_v133).Get(false).(_dafny.Sequence)
					}
					return _1105_v5
				})(), !((_this).F29()))
			}
			return p0
		})()
		return r0, r1
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f30 _dafny.Sequence
	_f29 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f30 = _dafny.EmptySeq
	_this._f29 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C6{}
var _ T1 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C6) F30_set_(value _dafny.Sequence) {
	_this._f30 = value
}
func (_this *C6) F29() bool {
	return _this._f29
}
func (_this *C6) Ctor__(f29 bool, f30 _dafny.Sequence) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C6) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F29()
	}
}
func (_this *C6) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1295_v0 _dafny.Int
		_ = _1295_v0
		_1295_v0 = _dafny.IntOfInt64(914)
		var _1296_v1 _dafny.MultiSet
		_ = _1296_v1
		_1296_v1 = _dafny.MultiSetOf(_1295_v0, _1295_v0, _1295_v0, _1295_v0)
		var _hi7 _dafny.Int = _1295_v0
		_ = _hi7
		for _1297_i0 := (_1296_v1).Cardinality(); _1297_i0.Cmp(_hi7) < 0; _1297_i0 = _1297_i0.Plus(_dafny.One) {
			(globalState).F19 = (_1297_i0).Plus((_dafny.Zero).Minus(_1297_i0))
			var _1298_v2 _dafny.Array
			_ = _1298_v2
			var _nw224 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw224
			_1298_v2 = _nw224
			var _nw225 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw225
			_1298_v2 = _nw225
			if !(!((_this).F29()) || ((_this).F29())) {
				var _1299_v3 _dafny.Array
				_ = _1299_v3
				var _nw226 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw226
				_1299_v3 = _nw226
				var _1300_v4 D0
				_ = _1300_v4
				_1300_v4 = Companion_D0_.Create_DC1_(_1299_v3, _1295_v0, _1297_i0)
				(globalState).F1 = (_1300_v4).Dtor_cf3()
				var _1301_v5 _dafny.MultiSet
				_ = _1301_v5
				_1301_v5 = _dafny.MultiSetOf(p0)
				var _1302_v6 D0
				_ = _1302_v6
				_1302_v6 = Companion_D0_.Create_DC2_(_1301_v5, _this.F30(), p0)
				var _1303_v7 _dafny.Map
				_ = _1303_v7
				_1303_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1302_v6, _1295_v0)
				_1303_v7 = _1303_v7
				var _1304_v8 *C1
				_ = _1304_v8
				var _nw227 *C1 = New_C1_()
				_ = _nw227
				_nw227.Ctor__(p0)
				_1304_v8 = _nw227
				var _1305_v9 _dafny.Map
				_ = _1305_v9
				_1305_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())).Cardinality()), (_1295_v0).Cmp(_1295_v0) >= 0)
				var _1306_v10 _dafny.Sequence
				_ = _1306_v10
				_1306_v10 = _dafny.SeqOf((_1304_v8).F40())
				_1305_v9 = (_1305_v9).Update((_1297_i0).Times(_dafny.IntOfUint32((_1306_v10).Cardinality())), (p0) == (p0))
				var _rhs221 _dafny.Int = Companion_Default___.Fm36(globalState)
				_ = _rhs221
				var _rhs222 bool = p0
				_ = _rhs222
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				var _lhs209 *GlobalState = globalState
				_ = _lhs209
				_lhs208.F21 = _rhs221
				_lhs209.F13 = _rhs222
			} else {
				var _1307_v11 _dafny.Sequence
				_ = _1307_v11
				_1307_v11 = _dafny.SeqOf(p0)
				(globalState).F21 = (_dafny.IntOfUint32((_1307_v11).Cardinality())).Times(_1295_v0)
				r1 = !(true)
				var _1308_v12 _dafny.Array
				_ = _1308_v12
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw228
				_1308_v12 = _nw228
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1308_v12), 0))
				_ = _index197
				(_1308_v12).ArraySet1(!(p0), (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1308_v12), 0))
				_ = _index198
				(_1308_v12).ArraySet1(!((_this).F29()), (_index198).Int())
				(globalState).F18 = _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(207))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}(func(_1309_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})))
				var _1310_v13 *C5
				_ = _1310_v13
				var _nw229 *C5 = New_C5_()
				_ = _nw229
				_nw229.Ctor__(true, _this.F30())
				_1310_v13 = _nw229
			}
			(globalState).F14 = _1295_v0
		}
		var _1311_v14 D6
		_ = _1311_v14
		_1311_v14 = Companion_D6_.Create_DC23_((_this).F29(), _1295_v0)
		var _1312_v15 _dafny.Set
		_ = _1312_v15
		_1312_v15 = _dafny.SetOf((_this).F29())
		var _1313_v16 D11
		_ = _1313_v16
		_1313_v16 = Companion_D11_.Create_DC36_(_1311_v14, _1312_v15, (_this).F29())
		var _1314_v17 _dafny.Array
		_ = _1314_v17
		var _nwElement0_43 bool = (_1313_v16).Dtor_cf61()
		_ = _nwElement0_43
		var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.One)
		_ = _nw230
		(_nw230).ArraySet1(_nwElement0_43, 0)
		_1314_v17 = _nw230
		var _1315_v18 D0
		_ = _1315_v18
		_1315_v18 = Companion_D0_.Create_DC1_(_1314_v17, _1295_v0, _1295_v0)
		var _1316_v19 _dafny.Array
		_ = _1316_v19
		var _nwElement0_44 _dafny.Array = (func() _dafny.Array {
			if (_this).Fm2(p0, globalState) {
				return _1314_v17
			}
			return _1314_v17
		})()
		_ = _nwElement0_44
		var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(11))
		_ = _nw231
		(_nw231).ArraySet1(_nwElement0_44, 0)
		(_nw231).ArraySet1(_1314_v17, 1)
		(_nw231).ArraySet1(_1314_v17, 2)
		(_nw231).ArraySet1((func() _dafny.Array {
			if (_this).F29() {
				return _1314_v17
			}
			return _1314_v17
		})(), 3)
		(_nw231).ArraySet1(_1314_v17, 4)
		(_nw231).ArraySet1(_1314_v17, 5)
		(_nw231).ArraySet1(_1314_v17, 6)
		(_nw231).ArraySet1(_1314_v17, 7)
		(_nw231).ArraySet1((_1315_v18).Dtor_cf1(), 8)
		(_nw231).ArraySet1(_1314_v17, 9)
		(_nw231).ArraySet1(_1314_v17, 10)
		_1316_v19 = _nw231
		var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1316_v19), 0))
		_ = _index199
		(_1316_v19).ArraySet1(_1314_v17, (_index199).Int())
		var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1316_v19), 0))
		_ = _index200
		(_1316_v19).ArraySet1(_1314_v17, (_index200).Int())
		var _1317_v20 _dafny.Set
		_ = _1317_v20
		_1317_v20 = _dafny.SetOf(_1295_v0, _1295_v0)
		var _1318_v21 _dafny.Sequence
		_ = _1318_v21
		_1318_v21 = _dafny.SeqOf(p0, (_1317_v20).IsSubsetOf(_1317_v20), (_this).F29())
		(globalState).F13 = (_1318_v21).Select((Companion_Default___.SafeIndex(_1295_v0, _dafny.IntOfUint32((_1318_v21).Cardinality()))).Uint32()).(bool)
		for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1314_v17), 0))); ; {
			_guard_loop_8, _ok55 := _iter55()
			if !_ok55 {
				break
			}
			var _1319_i2 _dafny.Int
			_1319_i2 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_1319_i2).Sign() != -1) && ((_1319_i2).Cmp(_dafny.ArrayLen((_1314_v17), 0)) < 0)) {
				(_1314_v17).ArraySet1(p0, (_1319_i2).Int())
			}
		}
		var _1320_v22 _dafny.Map
		_ = _1320_v22
		_1320_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0) || (false), _1295_v0)
		var _1321_v23 _dafny.Sequence
		_ = _1321_v23
		_1321_v23 = _dafny.SeqOf(_1295_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kyov")).Cardinality()))
		_1320_v22 = (_1320_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1321_v23).Cardinality())))
		r1 = (_this).F29()
		var _1322_v24 _dafny.Map
		_ = _1322_v24
		_1322_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(99), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(112))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}((func(_1323_v21 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_1324_i3 _dafny.Int) _dafny.Sequence {
				return _1323_v21
			}
		})(_1318_v21))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F29())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(112))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}((func(_1325_v21 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_1326_i3 _dafny.Int) _dafny.Sequence {
				return _1325_v21
			}
		})(_1318_v21)))).Cardinality()))).Uint32(), _dafny.SeqOf(p0)))
		var _1327_v25 _dafny.Sequence
		_ = _1327_v25
		_1327_v25 = _dafny.SeqOf(Companion_Default___.Fm35(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer91 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
			return func(arg91 _dafny.Int) interface{} {
				return coer91(arg91)
			}
		}(func(_1328_i4 _dafny.Int) D2 {
			return Companion_D2_.Create_DC8_(_dafny.CodePoint('d'), (_this).F29())
		})), _1295_v0, globalState), _1318_v21)
		r0 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1322_v24).Contains(_1295_v0) {
				return (_1322_v24).Get(_1295_v0).(_dafny.Sequence)
			}
			return _1327_v25
		})(), _dafny.SeqOf(_1318_v21, _1318_v21))
		r1 = p0
		return r0, r1
	}
}
func (_this *C6) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1329_v0 _dafny.Map
		_ = _1329_v0
		_1329_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), !((_this).F29()))
		(globalState).F15 = (func() bool {
			if (_1329_v0).Contains((_this).F29()) {
				return (_1329_v0).Get((_this).F29()).(bool)
			}
			return (func() bool {
				if p0 {
					return true
				}
				return (_this).F29()
			})()
		})()
		var _1330_v1 _dafny.MultiSet
		_ = _1330_v1
		_1330_v1 = _dafny.MultiSetOf(_1329_v0, _1329_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29()))
		var _1331_v2 _dafny.Sequence
		_ = _1331_v2
		_1331_v2 = _dafny.SeqOf((_1330_v1).IsProperSubsetOf(_1330_v1), (_this).F29())
		var _1332_v3 D9
		_ = _1332_v3
		_1332_v3 = Companion_D9_.Create_DC28_(_1331_v2, p0, (_dafny.Zero).Minus(p2), !(true), (_this).F29())
		var _1333_v4 _dafny.Sequence
		_ = _1333_v4
		_1333_v4 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1331_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kgmvdm")).Cardinality()), _dafny.IntOfUint32((_1331_v2).Cardinality()))).Uint32(), p0), _dafny.SeqOf(p0))
		_1331_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1331_v2, _dafny.Companion_Sequence_.Update((_1332_v3).Dtor_cf41(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_1332_v3).Dtor_cf41()).Cardinality()))).Uint32(), p0)), (_1333_v4).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1333_v4).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _hi8 _dafny.Int = p2
		_ = _hi8
		for _1334_i0 := (p2).Plus(p2); _1334_i0.Cmp(_hi8) < 0; _1334_i0 = _1334_i0.Plus(_dafny.One) {
			_1331_v2 = _1331_v2
			var _1335_v5 _dafny.Array
			_ = _1335_v5
			var _nw232 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw232
			_1335_v5 = _nw232
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1335_v5), 0))
			_ = _index201
			(_1335_v5).ArraySet1(true, (_index201).Int())
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1335_v5), 0))
			_ = _index202
			(_1335_v5).ArraySet1(!((_this).F29()) || (p0), (_index202).Int())
			var _1336_v6 _dafny.Map
			_ = _1336_v6
			_1336_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), (_this).F29())
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1335_v5), 0))
			_ = _index203
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1335_v5), 0))
			_ = _index204
			var _rhs223 bool = p0
			_ = _rhs223
			var _rhs224 bool = (func() bool {
				if (_dafny.IntOfInt64(40)).Cmp((_1336_v6).Cardinality()) != 0 {
					return p0
				}
				return Companion_Default___.Fm6(p2, globalState)
			})()
			_ = _rhs224
			var _rhs225 bool = p0
			_ = _rhs225
			var _rhs226 _dafny.Int = Companion_Default___.Fm36(globalState)
			_ = _rhs226
			var _rhs227 bool = p0
			_ = _rhs227
			var _lhs210 *GlobalState = globalState
			_ = _lhs210
			var _lhs211 *GlobalState = globalState
			_ = _lhs211
			var _lhs212 _dafny.Array = _1335_v5
			_ = _lhs212
			var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1335_v5), 0))
			_ = _lhs213
			var _lhs214 *GlobalState = globalState
			_ = _lhs214
			var _lhs215 _dafny.Array = _1335_v5
			_ = _lhs215
			var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1335_v5), 0))
			_ = _lhs216
			_lhs210.F15 = _rhs223
			_lhs211.F18 = _rhs224
			(_lhs212).ArraySet1(_rhs225, (_lhs213).Int())
			_lhs214.F2 = _rhs226
			(_lhs215).ArraySet1(_rhs227, (_lhs216).Int())
			var _1337_v7 _dafny.Sequence
			_ = _1337_v7
			_1337_v7 = _dafny.SeqOf(_this.F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}(func(_1338_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})))
			var _1339_v8 _dafny.Sequence
			_ = _1339_v8
			_1339_v8 = _dafny.SeqOf(_1334_i0, (_dafny.Zero).Minus(p2), _1334_i0, p2, _1334_i0)
			var _1340_v9 bool
			_ = _1340_v9
			var _1341_v10 _dafny.Sequence
			_ = _1341_v10
			var _out18 bool
			_ = _out18
			var _out19 _dafny.Sequence
			_ = _out19
			_out18, _out19 = (_this).M5((_1337_v7).Select((Companion_Default___.SafeIndex(_1334_i0, _dafny.IntOfUint32((_1337_v7).Cardinality()))).Uint32()).(_dafny.Sequence), (func() bool {
				if !(p0) {
					return (_1335_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1335_v5), 0))).Int()).(bool)
				}
				return p0
			})(), (_dafny.IntOfUint32((_1339_v8).Cardinality())).Minus(Companion_Default___.Fm36(globalState)), globalState)
			_1340_v9 = _out18
			_1341_v10 = _out19
			var _1342_v11 *C0
			_ = _1342_v11
			var _nw233 *C0 = New_C0_()
			_ = _nw233
			_nw233.Ctor__(false)
			_1342_v11 = _nw233
			var _1343_v12 _dafny.Sequence
			_ = _1343_v12
			_1343_v12 = _dafny.SeqOf(_1342_v11, _1342_v11, _1342_v11)
			var _1344_v13 D9
			_ = _1344_v13
			_1344_v13 = Companion_D9_.Create_DC29_(_dafny.IntOfUint32((_1343_v12).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1334_i0)).Cardinality(), (_this).F29())
			_1329_v0 = (_1329_v0).Update((_1344_v13).Dtor_cf48(), (_1342_v11).Fm11(_1340_v9, globalState))
		}
		if (_this).F29() {
			var _1345_v14 _dafny.Array
			_ = _1345_v14
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw234
			_1345_v14 = _nw234
			var _rhs228 _dafny.Int = (_dafny.Zero).Minus(p2)
			_ = _rhs228
			var _rhs229 _dafny.Array = _1345_v14
			_ = _rhs229
			var _rhs230 _dafny.Array = _1345_v14
			_ = _rhs230
			var _lhs217 *GlobalState = globalState
			_ = _lhs217
			_lhs217.F2 = _rhs228
			_1345_v14 = _rhs229
			_1345_v14 = _rhs230
			var _1346_v15 *C1
			_ = _1346_v15
			var _nw235 *C1 = New_C1_()
			_ = _nw235
			_nw235.Ctor__((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-156))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}(func(_1347_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Cardinality())).Cmp(p2) <= 0)
			_1346_v15 = _nw235
			var _1348_v16 _dafny.MultiSet
			_ = _1348_v16
			_1348_v16 = _dafny.MultiSetOf(p2, p2)
			var _1349_v17 _dafny.Map
			_ = _1349_v17
			_1349_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(495), _1348_v16)
			var _1350_v18 _dafny.Map
			_ = _1350_v18
			_1350_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1351_v19 _dafny.Sequence
			_ = _1351_v19
			_1351_v19 = _dafny.SeqOf(p2, p2, ((func() _dafny.MultiSet {
				if (_1349_v17).Contains((_1350_v18).Cardinality()) {
					return (_1349_v17).Get((_1350_v18).Cardinality()).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(Companion_Default___.Fm36(globalState))
			})()).Cardinality())
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1351_v19, _1351_v19), _1351_v19) {
				var _1352_v20 _dafny.Array
				_ = _1352_v20
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw236
				_1352_v20 = _nw236
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1352_v20), 0))
				_ = _index205
				(_1352_v20).ArraySet1(p2, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1352_v20), 0))
				_ = _index206
				(_1352_v20).ArraySet1(p2, (_index206).Int())
				var _1353_v21 _dafny.Set
				_ = _1353_v21
				_1353_v21 = _dafny.SetOf(_1352_v20, _1352_v20)
				var _1354_v22 *C2
				_ = _1354_v22
				var _nw237 *C2 = New_C2_()
				_ = _nw237
				_nw237.Ctor__((_1352_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1352_v20), 0))).Int()).(_dafny.Int), _1353_v21, (_1346_v15).F40(), _this.F30())
				_1354_v22 = _nw237
				var _1355_v23 _dafny.Map
				_ = _1355_v23
				_1355_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1354_v22, _1329_v0)
				var _1356_v24 _dafny.Map
				_ = _1356_v24
				_1356_v24 = _1329_v0
				var _1357_v25 _dafny.Array
				_ = _1357_v25
				var _nwElement0_45 _dafny.Map = _1329_v0
				_ = _nwElement0_45
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(14))
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_45, 0)
				(_nw238).ArraySet1(_1329_v0, 1)
				(_nw238).ArraySet1(_1329_v0, 2)
				(_nw238).ArraySet1(_1329_v0, 3)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1346_v15).F40(), (_1346_v15).F40()), 4)
				(_nw238).ArraySet1((func() _dafny.Map {
					if (_1355_v23).Contains(_1354_v22) {
						return (_1355_v23).Get(_1354_v22).(_dafny.Map)
					}
					return _1329_v0
				})(), 5)
				(_nw238).ArraySet1((_1329_v0).Merge(_1329_v0), 6)
				(_nw238).ArraySet1(((_1329_v0).Update((_this).F29(), (_this).F29())).Merge(_1329_v0), 7)
				(_nw238).ArraySet1(_1329_v0, 8)
				(_nw238).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1346_v15).F40(), (_this).F29()), 9)
				(_nw238).ArraySet1((_1356_v24), 10)
				(_nw238).ArraySet1(_1329_v0, 11)
				(_nw238).ArraySet1(_1329_v0, 12)
				(_nw238).ArraySet1(_1329_v0, 13)
				_1357_v25 = _nw238
				var _1358_v26 _dafny.CodePoint
				_ = _1358_v26
				_1358_v26 = _dafny.CodePoint('k')
				var _1359_v27 D2
				_ = _1359_v27
				_1359_v27 = Companion_D2_.Create_DC8_(_1358_v26, (_this).F29())
				var _1360_v28 D5
				_ = _1360_v28
				_1360_v28 = Companion_D5_.Create_DC19_((_1352_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1352_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1331_v2).Cardinality()), (_1359_v27).Dtor_cf12())
				var _rhs231 _dafny.Array = _1357_v25
				_ = _rhs231
				var _rhs232 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1351_v19, (Companion_Default___.SafeIndex((_1360_v28).Dtor_cf28(), _dafny.IntOfUint32((_1351_v19).Cardinality()))).Uint32(), Companion_Default___.SafeDivisionInt((_1354_v22).F38(), (_1354_v22).F38()))
				_ = _rhs232
				var _rhs233 _dafny.Int = p2
				_ = _rhs233
				var _rhs234 bool = _dafny.Companion_Sequence_.Contains(_1331_v2, _dafny.Companion_Sequence_.Equal(_this.F30(), _this.F30()))
				_ = _rhs234
				var _rhs235 _dafny.Int = ((_1354_v22).F38()).Times((_1354_v22).F38())
				_ = _rhs235
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				_1357_v25 = _rhs231
				_lhs218.F7 = _rhs232
				_lhs219.F21 = _rhs233
				_lhs220.F13 = _rhs234
				_lhs221.F1 = _rhs235
				var _1361_v29 *C0
				_ = _1361_v29
				var _nw239 *C0 = New_C0_()
				_ = _nw239
				_nw239.Ctor__((_this).F29())
				_1361_v29 = _nw239
				var _1362_v30 _dafny.Map
				_ = _1362_v30
				_1362_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1361_v29.F37, _dafny.IntOfInt64(994))
				(globalState).F21 = Companion_Default___.SafeModuloInt((_1362_v30).Cardinality(), ((_1354_v22).F38()).Minus((_1361_v29).Fm12(globalState)))
				var _1363_v31 _dafny.Set
				_ = _1363_v31
				_1363_v31 = _dafny.SetOf((_dafny.IntOfInt64(-182)).Plus(p2), (p2).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality())), p2, (_1352_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1352_v20), 0))).Int()).(_dafny.Int))
				_1363_v31 = _1363_v31
			} else {
				(globalState).F1 = p2
				var _1364_v32 _dafny.Array
				_ = _1364_v32
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw240
				_1364_v32 = _nw240
				var _1365_v33 _dafny.Array
				_ = _1365_v33
				var _nwElement0_46 _dafny.Array = _1364_v32
				_ = _nwElement0_46
				var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(9))
				_ = _nw241
				(_nw241).ArraySet1(_nwElement0_46, 0)
				(_nw241).ArraySet1(_1364_v32, 1)
				(_nw241).ArraySet1(_1364_v32, 2)
				(_nw241).ArraySet1(_1364_v32, 3)
				(_nw241).ArraySet1(_1364_v32, 4)
				(_nw241).ArraySet1(_1364_v32, 5)
				(_nw241).ArraySet1(_1364_v32, 6)
				(_nw241).ArraySet1(_1364_v32, 7)
				(_nw241).ArraySet1(_1364_v32, 8)
				_1365_v33 = _nw241
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_1365_v33), 0))
				_ = _index207
				(_1365_v33).ArraySet1(_1364_v32, (_index207).Int())
				var _1366_v34 _dafny.Set
				_ = _1366_v34
				_1366_v34 = _dafny.SetOf(p2, p2, p2, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), _dafny.IntOfUint32((_1331_v2).Cardinality()))
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_1365_v33), 0))
				_ = _index208
				var _rhs236 _dafny.Array = _1364_v32
				_ = _rhs236
				var _rhs237 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_this.F30()).Cardinality()), p2)
				_ = _rhs237
				var _rhs238 _dafny.Int = ((_1366_v34).Cardinality()).Times((func() _dafny.Int {
					if !(p0) {
						return p2
					}
					return p2
				})())
				_ = _rhs238
				var _rhs239 bool = (_this).F29()
				_ = _rhs239
				var _lhs222 _dafny.Array = _1365_v33
				_ = _lhs222
				var _lhs223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_1365_v33), 0))
				_ = _lhs223
				var _lhs224 *GlobalState = globalState
				_ = _lhs224
				var _lhs225 *GlobalState = globalState
				_ = _lhs225
				var _lhs226 *GlobalState = globalState
				_ = _lhs226
				(_lhs222).ArraySet1(_rhs236, (_lhs223).Int())
				_lhs224.F21 = _rhs237
				_lhs225.F21 = _rhs238
				_lhs226.F28 = _rhs239
				_1350_v18 = (_1350_v18).Update(_dafny.IntOfUint32((_1331_v2).Cardinality()), (_this).F29())
				(_this).F30_set_(_this.F30())
				var _1367_v35 _dafny.CodePoint
				_ = _1367_v35
				_1367_v35 = _dafny.CodePoint('e')
				var _rhs240 _dafny.Array = _1345_v14
				_ = _rhs240
				var _rhs241 _dafny.CodePoint = _1367_v35
				_ = _rhs241
				var _rhs242 _dafny.Map = _1350_v18
				_ = _rhs242
				_1345_v14 = _rhs240
				_1367_v35 = _rhs241
				_1350_v18 = _rhs242
			}
			var _1368_v36 _dafny.Array
			_ = _1368_v36
			var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw242
			_1368_v36 = _nw242
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))
			_ = _index209
			(_1368_v36).ArraySet1(_1345_v14, (_index209).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))
			_ = _index210
			(_1368_v36).ArraySet1(_1345_v14, (_index210).Int())
			var _1369_v37 D3
			_ = _1369_v37
			_1369_v37 = Companion_D3_.Create_DC10_(_1348_v16)
			if (_1348_v16).IsSubsetOf(((_1369_v37).Dtor_cf18()).Union(_dafny.MultiSetOf(p2))) {
				var _1370_v38 D12
				_ = _1370_v38
				_1370_v38 = Companion_D12_.Create_DC39_((_1329_v0).Cardinality(), (_1346_v15).F40(), p2)
				var _1371_v39 D6
				_ = _1371_v39
				_1371_v39 = Companion_D6_.Create_DC23_(p0, p2)
				var _pat_let_tv28 = _1346_v15
				_ = _pat_let_tv28
				var _rhs243 _dafny.Int = (func(_pat_let31_0 D6) D6 {
					return func(_1372_dt__update__tmp_h0 D6) D6 {
						return func(_pat_let32_0 bool) D6 {
							return func(_1373_dt__update_hcf36_h0 bool) D6 {
								return Companion_D6_.Create_DC23_(_1373_dt__update_hcf36_h0, (_1372_dt__update__tmp_h0).Dtor_cf37())
							}(_pat_let32_0)
						}((_pat_let_tv28).F40())
					}(_pat_let31_0)
				}(_1371_v39)).Dtor_cf37()
				_ = _rhs243
				var _rhs244 bool = !(!(p0))
				_ = _rhs244
				var _rhs245 D12 = _1370_v38
				_ = _rhs245
				var _rhs246 _dafny.Int = p2
				_ = _rhs246
				var _rhs247 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((p2).Times(_dafny.IntOfInt64(-620)), _dafny.IntOfInt64(-108)))
				_ = _rhs247
				var _lhs227 *GlobalState = globalState
				_ = _lhs227
				var _lhs228 *GlobalState = globalState
				_ = _lhs228
				var _lhs229 *GlobalState = globalState
				_ = _lhs229
				_lhs227.F19 = _rhs243
				r1 = _rhs244
				_1370_v38 = _rhs245
				_lhs228.F19 = _rhs246
				_lhs229.F14 = _rhs247
				var _1374_v40 _dafny.Array
				_ = _1374_v40
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_35
				var _nw243 _dafny.Array
				_ = _nw243
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw243 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) D11 = (func(_1375_v15 *C1, _1376_p2 _dafny.Int) func(_dafny.Int) D11 {
						return func(_1377_i3 _dafny.Int) D11 {
							return (func() D11 {
								if (_1375_v15).F40() {
									return Companion_D11_.Create_DC37_(_1376_p2, (_1375_v15).F40())
								}
								return Companion_D11_.Create_DC37_(_dafny.IntOfUint32((_this.F30()).Cardinality()), false)
							})()
						}
					})(_1346_v15, p2)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw243 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw243).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw243).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1374_v40 = _nw243
				var _1378_v41 D11
				_ = _1378_v41
				_1378_v41 = Companion_D11_.Create_DC37_(p2, (_this).F29())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1374_v40), 0))
				_ = _index211
				(_1374_v40).ArraySet1(_1378_v41, (_index211).Int())
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))
				_ = _arr0
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))), 0))
				_ = _index212
				(_arr0).ArraySet1((_1346_v15).F40(), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1374_v40), 0))
				_ = _index213
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))
				_ = _arr1
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))), 0))
				_ = _index214
				var _rhs248 D11 = _1378_v41
				_ = _rhs248
				var _rhs249 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1351_v19, _dafny.SeqOf((_1329_v0).Cardinality())), _1351_v19)
				_ = _rhs249
				var _rhs250 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf((_1346_v15).F40(), p0))
				_ = _rhs250
				var _lhs230 _dafny.Array = _1374_v40
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1374_v40), 0))
				_ = _lhs231
				var _lhs232 _dafny.Array = _dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_dafny.ArrayCastTo((_1368_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1368_v36), 0))).Int()))), 0))
				_ = _lhs233
				(_lhs230).ArraySet1(_rhs248, (_lhs231).Int())
				(_lhs232).ArraySet1(_rhs249, (_lhs233).Int())
				_1331_v2 = _rhs250
				(globalState).F28 = (_1346_v15).F40()
				var _1379_v42 _dafny.CodePoint
				_ = _1379_v42
				_1379_v42 = _dafny.CodePoint('f')
				var _1380_v44 _dafny.Set
				_ = _1380_v44
				_1380_v44 = _dafny.SetOf(p2)
				var _1381_v45 _dafny.Sequence
				_ = _1381_v45
				_1381_v45 = _dafny.SeqOf(func() _dafny.Set {
					var _coll47 = _dafny.NewBuilder()
					_ = _coll47
					for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(644), _dafny.IntOfInt64(-791))); ; {
						_compr_47, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1382_v43 _dafny.Int
						_1382_v43 = interface{}(_compr_47).(_dafny.Int)
						if ((_dafny.IntOfInt64(644)).Cmp(_1382_v43) <= 0) && ((_1382_v43).Cmp(_dafny.IntOfInt64(-791)) < 0) {
							_coll47.Add((_1382_v43).Plus(p2))
						}
					}
					return _coll47.ToSet()
				}(), _1380_v44, _1380_v44, _1380_v44)
				_1379_v42 = Companion_Default___.Fm42((Companion_Default___.Fm39(_1379_v42, _dafny.IntOfInt64(883), globalState)).Equals(Companion_Default___.Fm39(_1379_v42, (_dafny.Zero).Minus(p2), globalState)), _dafny.Companion_Sequence_.IsPrefixOf(_1381_v45, _1381_v45), globalState)
				var _1383_v46 _dafny.Array
				_ = _1383_v46
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_36
				var _nw244 _dafny.Array
				_ = _nw244
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw244 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = func(_1384_i4 _dafny.Int) _dafny.Int {
						return (_1384_i4).Minus(_dafny.IntOfInt64(142))
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw244 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw244).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw244).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1383_v46 = _nw244
				var _1385_v47 D12
				_ = _1385_v47
				_1385_v47 = Companion_D12_.Create_DC38_(_1383_v46)
				(globalState).F17 = (_1385_v47).Dtor_cf64()
			} else {
				(globalState).F20 = _this.F30()
				(globalState).F15 = ((p0) && ((_1346_v15).F40())) || ((_this).F29())
				var _1386_v48 _dafny.Map
				_ = _1386_v48
				_1386_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_dafny.Zero).Minus((_dafny.Zero).Minus(p2))).Cmp(p2) < 0), (p2).Plus(p2))
				var _1387_v49 _dafny.Map
				_ = _1387_v49
				_1387_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				_1386_v48 = (_1386_v48).Update(!(p0), (func() _dafny.Int {
					if (_1387_v49).Contains(Companion_Default___.Fm9(globalState)) {
						return (_1387_v49).Get(Companion_Default___.Fm9(globalState)).(_dafny.Int)
					}
					return p2
				})())
				var _1388_v50 _dafny.Array
				_ = _1388_v50
				var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw245
				_1388_v50 = _nw245
				var _1389_v51 _dafny.Set
				_ = _1389_v51
				_1389_v51 = _dafny.SetOf(_1388_v50, _1388_v50)
				var _1390_v52 *C3
				_ = _1390_v52
				var _nw246 *C3 = New_C3_()
				_ = _nw246
				_nw246.Ctor__(!(_1350_v18).Equals(_1350_v18), (p2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wgqvib")).Cardinality())), (_1389_v51).IsProperSubsetOf(_1389_v51), _this.F30())
				_1390_v52 = _nw246
				_1390_v52 = _1390_v52
				_1388_v50 = _1388_v50
			}
		} else {
			var _1391_v53 _dafny.MultiSet
			_ = _1391_v53
			_1391_v53 = _dafny.MultiSetOf(Companion_D11_.Create_DC37_(p2, p0), Companion_D11_.Create_DC37_(p2, p0))
			var _1392_v54 D11
			_ = _1392_v54
			_1392_v54 = Companion_D11_.Create_DC37_(p2, p0)
			_1391_v53 = (_1391_v53).Difference(_dafny.MultiSetOf(_1392_v54))
			var _1393_v55 *C3
			_ = _1393_v55
			var _nw247 *C3 = New_C3_()
			_ = _nw247
			_nw247.Ctor__((func() bool {
				if p0 {
					return p0
				}
				return !(p0)
			})(), p2, _dafny.Companion_Sequence_.IsPrefixOf(_1331_v2, _1331_v2), _this.F30())
			_1393_v55 = _nw247
			(globalState).F19 = (_1393_v55).F36()
			var _1394_v56 _dafny.Array
			_ = _1394_v56
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_37
			var _nw248 _dafny.Array
			_ = _nw248
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw248 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = (func(_1395_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1396_i5 _dafny.Int) _dafny.Int {
						return (_1396_i5).Plus(_1395_p2)
					}
				})(p2)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw248 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw248).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw248).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1394_v56 = _nw248
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1394_v56), 0))
			_ = _index215
			(_1394_v56).ArraySet1(p2, (_index215).Int())
			var _1397_v57 _dafny.Array
			_ = _1397_v57
			var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw249
			_1397_v57 = _nw249
			var _1398_v58 _dafny.Array
			_ = _1398_v58
			var _nw250 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw250
			_1398_v58 = _nw250
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1397_v57), 0))
			_ = _index216
			(_1397_v57).ArraySet1(_1398_v58, (_index216).Int())
			var _1399_v59 _dafny.Map
			_ = _1399_v59
			_1399_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1393_v55).F36())
			var _1400_v60 D0
			_ = _1400_v60
			_1400_v60 = Companion_D0_.Create_DC1_(_1398_v58, (_1393_v55).F36(), (_1332_v3).Dtor_cf43())
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1394_v56), 0))
			_ = _index217
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1397_v57), 0))
			_ = _index218
			var _rhs251 _dafny.Int = ((p2).Minus((_dafny.Zero).Minus((_1393_v55).F36()))).Minus((func() _dafny.Int {
				if (_1399_v59).Contains((_1393_v55).F36()) {
					return (_1399_v59).Get((_1393_v55).F36()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_1393_v55).F36())).Cardinality()))
			})())
			_ = _rhs251
			var _rhs252 _dafny.Int = (_1393_v55).F36()
			_ = _rhs252
			var _rhs253 _dafny.Array = (_1400_v60).Dtor_cf1()
			_ = _rhs253
			var _lhs234 _dafny.Array = _1394_v56
			_ = _lhs234
			var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1394_v56), 0))
			_ = _lhs235
			var _lhs236 *GlobalState = globalState
			_ = _lhs236
			var _lhs237 _dafny.Array = _1397_v57
			_ = _lhs237
			var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1397_v57), 0))
			_ = _lhs238
			(_lhs234).ArraySet1(_rhs251, (_lhs235).Int())
			_lhs236.F19 = _rhs252
			(_lhs237).ArraySet1(_rhs253, (_lhs238).Int())
			var _1401_v61 *C4
			_ = _1401_v61
			var _nw251 *C4 = New_C4_()
			_ = _nw251
			_nw251.Ctor__(p0, p2)
			_1401_v61 = _nw251
			var _1402_v62 _dafny.Array
			_ = _1402_v62
			var _nwElement0_47 *C4 = _1401_v61
			_ = _nwElement0_47
			var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(20))
			_ = _nw252
			(_nw252).ArraySet1(_nwElement0_47, 0)
			(_nw252).ArraySet1(_1401_v61, 1)
			(_nw252).ArraySet1(_1401_v61, 2)
			(_nw252).ArraySet1(_1401_v61, 3)
			(_nw252).ArraySet1(_1401_v61, 4)
			(_nw252).ArraySet1(_1401_v61, 5)
			(_nw252).ArraySet1(_1401_v61, 6)
			(_nw252).ArraySet1(_1401_v61, 7)
			(_nw252).ArraySet1(_1401_v61, 8)
			(_nw252).ArraySet1(_1401_v61, 9)
			(_nw252).ArraySet1(_1401_v61, 10)
			(_nw252).ArraySet1(_1401_v61, 11)
			(_nw252).ArraySet1(_1401_v61, 12)
			(_nw252).ArraySet1(_1401_v61, 13)
			(_nw252).ArraySet1(_1401_v61, 14)
			(_nw252).ArraySet1(_1401_v61, 15)
			(_nw252).ArraySet1(_1401_v61, 16)
			(_nw252).ArraySet1(_1401_v61, 17)
			(_nw252).ArraySet1((func() *C4 {
				if (_1393_v55).F35() {
					return _1401_v61
				}
				return _1401_v61
			})(), 18)
			(_nw252).ArraySet1(_1401_v61, 19)
			_1402_v62 = _nw252
			var _1403_v63 D13
			_ = _1403_v63
			_1403_v63 = Companion_D13_.Create_DC40_(_1401_v61)
			var _1404_v64 _dafny.Sequence
			_ = _1404_v64
			_1404_v64 = _dafny.SeqOf(_1401_v61)
			var _1405_v65 _dafny.Map
			_ = _1405_v65
			_1405_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_1404_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.IntOfUint32((_1404_v64).Cardinality()))).Uint32()).(*C4))
			var _1406_v66 D3
			_ = _1406_v66
			_1406_v66 = Companion_D3_.Create_DC11_((_1393_v55).F35(), (_1401_v61).F33())
			var _nwElement0_48 *C4 = _1401_v61
			_ = _nwElement0_48
			var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(18))
			_ = _nw253
			(_nw253).ArraySet1(_nwElement0_48, 0)
			(_nw253).ArraySet1((func() *C4 {
				if (_1393_v55).F35() {
					return (_1403_v63).Dtor_cf68()
				}
				return _1401_v61
			})(), 1)
			(_nw253).ArraySet1(_1401_v61, 2)
			(_nw253).ArraySet1(_1401_v61, 3)
			(_nw253).ArraySet1((func() *C4 {
				if (_1405_v65).Contains((_1401_v61).F33()) {
					return (_1405_v65).Get((_1401_v61).F33()).(*C4)
				}
				return _1401_v61
			})(), 4)
			(_nw253).ArraySet1(_1401_v61, 5)
			(_nw253).ArraySet1(_1401_v61, 6)
			(_nw253).ArraySet1(_1401_v61, 7)
			(_nw253).ArraySet1(_1401_v61, 8)
			(_nw253).ArraySet1(_1401_v61, 9)
			(_nw253).ArraySet1((func() *C4 {
				if p0 {
					return _1401_v61
				}
				return _1401_v61
			})(), 10)
			(_nw253).ArraySet1((func() *C4 {
				if (_this).F29() {
					return _1401_v61
				}
				return _1401_v61
			})(), 11)
			(_nw253).ArraySet1(_1401_v61, 12)
			(_nw253).ArraySet1(_1401_v61, 13)
			(_nw253).ArraySet1(_1401_v61, 14)
			(_nw253).ArraySet1(_1401_v61, 15)
			(_nw253).ArraySet1((func() *C4 {
				if (_1406_v66).Dtor_cf20() {
					return _1401_v61
				}
				return _1401_v61
			})(), 16)
			(_nw253).ArraySet1(_1401_v61, 17)
			_1402_v62 = _nw253
		}
		var _1407_v68 _dafny.Array
		_ = _1407_v68
		var _len0_38 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_38
		var _nw254 _dafny.Array
		_ = _nw254
		if _len0_38.Cmp(_dafny.Zero) == 0 {
			_nw254 = _dafny.NewArray(_len0_38)
		} else {
			var _init38 func(_dafny.Int) bool = (func(_1408_p0 bool) func(_dafny.Int) bool {
				return func(_1409_i7 _dafny.Int) bool {
					return (_dafny.MultiSetOf(func() _dafny.Map {
						var _coll48 = _dafny.NewMapBuilder()
						_ = _coll48
						for _iter57 := _dafny.Iterate((_this.F30()).Elements()); ; {
							_compr_48, _ok57 := _iter57()
							if !_ok57 {
								break
							}
							var _1410_v67 _dafny.CodePoint
							_1410_v67 = interface{}(_compr_48).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_this.F30(), _1410_v67) {
								_coll48.Add(_1410_v67, false)
							}
						}
						return _coll48.ToMap()
					}())).IsSubsetOf(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _1408_p0)))
				}
			})(p0)
			_ = _init38
			var _element0_38 = _init38(_dafny.Zero)
			_ = _element0_38
			_nw254 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
			(_nw254).ArraySet1(_element0_38, 0)
			var _nativeLen0_38 = (_len0_38).Int()
			_ = _nativeLen0_38
			for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
				(_nw254).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
			}
		}
		_1407_v68 = _nw254
		for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1407_v68), 0))); ; {
			_guard_loop_9, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _1411_i6 _dafny.Int
			_1411_i6 = interface{}(_guard_loop_9).(_dafny.Int)
			if (true) && (((_1411_i6).Sign() != -1) && ((_1411_i6).Cmp(_dafny.ArrayLen((_1407_v68), 0)) < 0)) {
				(_1407_v68).ArraySet1((Companion_D11_.Create_DC35_(_this.F30(), p0)).Dtor_cf58(), (_1411_i6).Int())
			}
		}
		var _1412_v69 _dafny.Array
		_ = _1412_v69
		var _nwElement0_49 _dafny.Sequence = _this.F30()
		_ = _nwElement0_49
		var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(14))
		_ = _nw255
		(_nw255).ArraySet1(_nwElement0_49, 0)
		(_nw255).ArraySet1(_this.F30(), 1)
		(_nw255).ArraySet1(_this.F30(), 2)
		(_nw255).ArraySet1(_this.F30(), 3)
		(_nw255).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()), 4)
		(_nw255).ArraySet1(_this.F30(), 5)
		(_nw255).ArraySet1(_this.F30(), 6)
		(_nw255).ArraySet1(_this.F30(), 7)
		(_nw255).ArraySet1(_this.F30(), 8)
		(_nw255).ArraySet1(_this.F30(), 9)
		(_nw255).ArraySet1(_this.F30(), 10)
		(_nw255).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("nnvokyy")), 11)
		(_nw255).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uqootstpb"), _this.F30()), 12)
		(_nw255).ArraySet1((func() _dafny.Sequence {
			if (_this).F29() {
				return _this.F30()
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-207))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg94 _dafny.Int) interface{} {
					return coer94(arg94)
				}
			}(func(_1413_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))
		})(), 13)
		_1412_v69 = _nw255
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1412_v69), 0))); ; {
			_guard_loop_10, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1414_i8 _dafny.Int
			_1414_i8 = interface{}(_guard_loop_10).(_dafny.Int)
			if (true) && (((_1414_i8).Sign() != -1) && ((_1414_i8).Cmp(_dafny.ArrayLen((_1412_v69), 0)) < 0)) {
				(_1412_v69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()), (_1414_i8).Int())
			}
		}
		var _1415_v70 _dafny.Sequence
		_ = _1415_v70
		_1415_v70 = _dafny.SeqOf(_dafny.IntOfInt64(542))
		r0 = Companion_Default___.Fm34(p2, Companion_Default___.Fm36(globalState), p0, _dafny.Companion_Sequence_.IsPrefixOf(_1415_v70, _1415_v70), globalState)
		r1 = p0
		return r0, r1
	}
}
func (_this *C6) M5(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		if (_this).F29() {
			var _1416_v0 _dafny.Set
			_ = _1416_v0
			_1416_v0 = _dafny.SetOf((_dafny.SetOf(p1)).Cardinality())
			if ((_1416_v0).Union(_1416_v0)).Equals(_1416_v0) {
				var _1417_v1 _dafny.Array
				_ = _1417_v1
				var _nw256 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(13))
				_ = _nw256
				_1417_v1 = _nw256
				var _1418_v2 _dafny.CodePoint
				_ = _1418_v2
				_1418_v2 = _dafny.CodePoint('q')
				var _1419_v3 _dafny.MultiSet
				_ = _1419_v3
				_1419_v3 = _dafny.MultiSetOf(_dafny.CodePoint('i'), _1418_v2)
				var _1420_v4 D2
				_ = _1420_v4
				_1420_v4 = Companion_D2_.Create_DC6_(_1419_v3)
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1417_v1), 0))
				_ = _index219
				(_1417_v1).ArraySet1(_1420_v4, (_index219).Int())
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1417_v1), 0))
				_ = _index220
				(_1417_v1).ArraySet1(_1420_v4, (_index220).Int())
				var _1421_v5 *C0
				_ = _1421_v5
				var _nw257 *C0 = New_C0_()
				_ = _nw257
				_nw257.Ctor__(p1)
				_1421_v5 = _nw257
				var _1422_v6 _dafny.Sequence
				_ = _1422_v6
				_1422_v6 = _dafny.SeqOf(!(_1421_v5.F37) || ((_this).F29()))
				_1422_v6 = _dafny.Companion_Sequence_.Update(_1422_v6, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1422_v6).Cardinality()))).Uint32(), _1421_v5.F37)
				(globalState).F21 = p2
				_1418_v2 = Companion_Default___.Fm42(!(p1) || (p1), (_this).F29(), globalState)
			} else {
				var _1423_v7 _dafny.Map
				_ = _1423_v7
				_1423_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm48(globalState)).Cardinality(), (_dafny.Zero).Minus(p2))
				var _1424_v8 _dafny.Map
				_ = _1424_v8
				_1424_v8 = _1423_v7
				var _1425_v9 _dafny.Sequence
				_ = _1425_v9
				_1425_v9 = _dafny.SeqOf((_1424_v8), _1423_v7)
				var _1426_v10 _dafny.CodePoint
				_ = _1426_v10
				_1426_v10 = _dafny.CodePoint('k')
				_1423_v7 = ((_1425_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rksaphr"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rksaphr")).Cardinality()))).Uint32(), _1426_v10)).Cardinality()), _dafny.IntOfUint32((_1425_v9).Cardinality()))).Uint32()).(_dafny.Map)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Merge((_1423_v7).Update(p2, p2)))
				(globalState).F13 = (!((_this).F29())) && ((_this).F29())
				(globalState).F2 = (func() _dafny.Int {
					if p1 {
						return _dafny.IntOfInt64(605)
					}
					return p2
				})()
				_1426_v10 = _1426_v10
				(globalState).F14 = p2
			}
			var _1427_v11 _dafny.MultiSet
			_ = _1427_v11
			_1427_v11 = _dafny.MultiSetOf(p2)
			_1427_v11 = (_1427_v11).Intersection(_1427_v11)
			var _1428_v12 _dafny.Set
			_ = _1428_v12
			_1428_v12 = _dafny.SetOf(p1)
			var _1429_v13 _dafny.Map
			_ = _1429_v13
			_1429_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1428_v12).Union(_1428_v12)).Cardinality(), (p2).Cmp(p2) == 0)
			_1429_v13 = _1429_v13
			var _1430_v14 D15
			_ = _1430_v14
			_1430_v14 = Companion_D15_.Create_DC43_(_1429_v13)
			var _rhs254 _dafny.Map = (_1430_v14).Dtor_cf70()
			_ = _rhs254
			var _rhs255 bool = (_this).F29()
			_ = _rhs255
			var _lhs239 *GlobalState = globalState
			_ = _lhs239
			_1429_v13 = _rhs254
			_lhs239.F28 = _rhs255
			var _1431_v15 *C0
			_ = _1431_v15
			var _nw258 *C0 = New_C0_()
			_ = _nw258
			_nw258.Ctor__(true)
			_1431_v15 = _nw258
		} else {
			var _1432_v16 _dafny.CodePoint
			_ = _1432_v16
			_1432_v16 = _dafny.CodePoint('j')
			var _1433_v17 _dafny.Map
			_ = _1433_v17
			_1433_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-76), !((_this).F29()))
			var _1434_v18 D10
			_ = _1434_v18
			_1434_v18 = Companion_D10_.Create_DC32_(_1432_v16, p1)
			var _1435_v19 _dafny.Array
			_ = _1435_v19
			var _nwElement0_50 D10 = Companion_Default___.Fm49(_1432_v16, globalState)
			_ = _nwElement0_50
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(13))
			_ = _nw259
			(_nw259).ArraySet1(_nwElement0_50, 0)
			(_nw259).ArraySet1(Companion_D10_.Create_DC32_(_dafny.CodePoint('j'), (func() bool {
				if (_1433_v17).Contains(p2) {
					return (_1433_v17).Get(p2).(bool)
				}
				return (_this).F29()
			})()), 1)
			(_nw259).ArraySet1(_1434_v18, 2)
			(_nw259).ArraySet1(_1434_v18, 3)
			(_nw259).ArraySet1(_1434_v18, 4)
			(_nw259).ArraySet1(_1434_v18, 5)
			(_nw259).ArraySet1(_1434_v18, 6)
			(_nw259).ArraySet1(_1434_v18, 7)
			(_nw259).ArraySet1(_1434_v18, 8)
			(_nw259).ArraySet1(_1434_v18, 9)
			(_nw259).ArraySet1(Companion_D10_.Create_DC32_(_1432_v16, p1), 10)
			(_nw259).ArraySet1(Companion_D10_.Create_DC32_(_1432_v16, p1), 11)
			(_nw259).ArraySet1(_1434_v18, 12)
			_1435_v19 = _nw259
			var _1436_v20 _dafny.Array
			_ = _1436_v20
			var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw260
			_1436_v20 = _nw260
			var _1437_v21 _dafny.Sequence
			_ = _1437_v21
			_1437_v21 = _dafny.SeqOf(p1, (_this).F29(), p1)
			var _1438_v22 D15
			_ = _1438_v22
			_1438_v22 = Companion_D15_.Create_DC44_((_dafny.Zero).Minus(p2), p1, _1435_v19, _1436_v20, _dafny.IntOfUint32((_1437_v21).Cardinality()))
			_1438_v22 = (func() D15 {
				if false {
					return _1438_v22
				}
				return _1438_v22
			})()
			var _1439_v23 _dafny.MultiSet
			_ = _1439_v23
			_1439_v23 = _dafny.MultiSetOf(p2)
			var _1440_v24 _dafny.MultiSet
			_ = _1440_v24
			_1440_v24 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F29(), p1, (_this).F29()), Companion_Default___.Fm35(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-147))).Uint32(), func(coer95 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
				return func(arg95 _dafny.Int) interface{} {
					return coer95(arg95)
				}
			}(func(_1441_i0 _dafny.Int) D2 {
				return Companion_D2_.Create_DC8_(_dafny.CodePoint('n'), (_this).F29())
			})), (_1439_v23).Cardinality(), globalState)))
			(globalState).F1 = (func() _dafny.Int {
				if (_1440_v24).Contains(_dafny.Companion_Sequence_.Concatenate(_1437_v21, _1437_v21)) {
					return (_1440_v24).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_1437_v21, _1437_v21))
				}
				return p2
			})()
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1436_v20), 0))
			_ = _index221
			(_1436_v20).ArraySet1(Companion_Default___.Fm9(globalState), (_index221).Int())
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1436_v20), 0))
			_ = _index222
			(_1436_v20).ArraySet1((p2).Minus(p2), (_index222).Int())
			var _1442_v25 *C4
			_ = _1442_v25
			var _nw261 *C4 = New_C4_()
			_ = _nw261
			_nw261.Ctor__(p1, p2)
			_1442_v25 = _nw261
			(globalState).F28 = (_this).F29()
		}
		var _1443_v26 _dafny.Array
		_ = _1443_v26
		var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw262
		_1443_v26 = _nw262
		var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
		_ = _index223
		(_1443_v26).ArraySet1(Companion_Default___.SafeModuloInt(p2, (_dafny.Zero).Minus(p2)), (_index223).Int())
		var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
		_ = _index224
		(_1443_v26).ArraySet1(Companion_Default___.Fm36(globalState), (_index224).Int())
		var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
		_ = _index225
		(_1443_v26).ArraySet1(Companion_Default___.SafeModuloInt(p2, (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int)), (_index225).Int())
		var _hi9 _dafny.Int = (p2).Plus(p2)
		_ = _hi9
		for _1444_i1 := (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int); _1444_i1.Cmp(_hi9) < 0; _1444_i1 = _1444_i1.Plus(_dafny.One) {
			var _1445_v27 *C1
			_ = _1445_v27
			var _nw263 *C1 = New_C1_()
			_ = _nw263
			_nw263.Ctor__(!((_this).F29()))
			_1445_v27 = _nw263
			var _1446_v28 _dafny.Sequence
			_ = _1446_v28
			_1446_v28 = _dafny.SeqOf(_this.F30())
			var _1447_v29 _dafny.Sequence
			_ = _1447_v29
			_1447_v29 = _dafny.SeqOf((_1446_v28).Select((Companion_Default___.SafeIndex(_1444_i1, _dafny.IntOfUint32((_1446_v28).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1448_v30 _dafny.MultiSet
			_ = _1448_v30
			_1448_v30 = _dafny.MultiSetOf(p0, _this.F30(), p0, p0)
			var _1449_v31 _dafny.Map
			_ = _1449_v31
			_1449_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
			var _1450_v32 _dafny.CodePoint
			_ = _1450_v32
			_1450_v32 = _dafny.CodePoint('k')
			var _1451_v33 _dafny.MultiSet
			_ = _1451_v33
			_1451_v33 = _dafny.MultiSetOf(_1450_v32)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
			_ = _index226
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
			_ = _index227
			var _rhs256 _dafny.Int = _1444_i1
			_ = _rhs256
			var _rhs257 bool = !(((_dafny.MultiSetFromSeq(_1447_v29)).Update(_this.F30(), Companion_Default___.Abs(_1444_i1))).IsSubsetOf(_1448_v30))
			_ = _rhs257
			var _rhs258 bool = Companion_Default___.Fm33(Companion_Default___.Fm33((_1445_v27).F40(), (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), (_1449_v31).Cardinality(), globalState), ((_1451_v33).Update(_1450_v32, Companion_Default___.Abs(p2))).Cardinality(), _1444_i1, globalState)
			_ = _rhs258
			var _rhs259 _dafny.Int = _1444_i1
			_ = _rhs259
			var _lhs240 _dafny.Array = _1443_v26
			_ = _lhs240
			var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
			_ = _lhs241
			var _lhs242 *GlobalState = globalState
			_ = _lhs242
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			var _lhs244 _dafny.Array = _1443_v26
			_ = _lhs244
			var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))
			_ = _lhs245
			(_lhs240).ArraySet1(_rhs256, (_lhs241).Int())
			_lhs242.F18 = _rhs257
			_lhs243.F28 = _rhs258
			(_lhs244).ArraySet1(_rhs259, (_lhs245).Int())
			var _1452_v34 _dafny.Set
			_ = _1452_v34
			_1452_v34 = _dafny.SetOf((_this).F29(), (_1445_v27).F40(), p1)
			var _1453_v35 D6
			_ = _1453_v35
			_1453_v35 = Companion_D6_.Create_DC21_(_1452_v34)
			var _1454_v36 _dafny.Sequence
			_ = _1454_v36
			_1454_v36 = _dafny.SeqOf(p1)
			_1453_v35 = Companion_Default___.Fm50((_this).F29(), Companion_Default___.SafeDivisionInt(_1444_i1, _1444_i1), _dafny.Companion_Sequence_.Concatenate(_this.F30(), p0), (func() _dafny.Sequence {
				if (_1445_v27).F40() {
					return _1454_v36
				}
				return _dafny.SeqOf((_this).F29())
			})(), globalState)
			var _1455_v37 _dafny.Sequence
			_ = _1455_v37
			_1455_v37 = _dafny.SeqOf(Companion_Default___.Fm51(globalState))
			(globalState).F21 = _dafny.IntOfUint32((_1455_v37).Cardinality())
		}
		var _1456_v38 _dafny.Sequence
		_ = _1456_v38
		_1456_v38 = _dafny.SeqOf(true)
		var _1457_v39 _dafny.Map
		_ = _1457_v39
		_1457_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _dafny.IntOfUint32((_1456_v38).Cardinality()))
		var _hi10 _dafny.Int = p2
		_ = _hi10
		for _1458_i2 := Companion_Default___.SafeModuloInt(p2, (_1457_v39).Cardinality()); _1458_i2.Cmp(_hi10) < 0; _1458_i2 = _1458_i2.Plus(_dafny.One) {
			var _1459_v40 _dafny.MultiSet
			_ = _1459_v40
			_1459_v40 = _dafny.MultiSetOf(p1)
			var _1460_v41 _dafny.CodePoint
			_ = _1460_v41
			_1460_v41 = _dafny.CodePoint('l')
			var _1461_v42 D2
			_ = _1461_v42
			_1461_v42 = Companion_D2_.Create_DC8_(_1460_v41, p1)
			var _1462_v43 _dafny.Sequence
			_ = _1462_v43
			_1462_v43 = _dafny.SeqOf(_1461_v42, _1461_v42)
			_1459_v40 = ((func() _dafny.MultiSet {
				if p1 {
					return _dafny.MultiSetOf(true, (_this).F29())
				}
				return _1459_v40
			})()).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm35(_1462_v43, (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), globalState), _1456_v38)))
			var _1463_v44 D9
			_ = _1463_v44
			_1463_v44 = Companion_D9_.Create_DC29_((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), _1458_i2, false)
			var _1464_v45 _dafny.Map
			_ = _1464_v45
			_1464_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F29())
			(globalState).F28 = (_1464_v45).Contains((_1463_v44).Dtor_cf48())
			var _1465_v46 _dafny.MultiSet
			_ = _1465_v46
			_1465_v46 = _dafny.MultiSetOf(_1458_i2, (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), _1458_i2, (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), p2)
			var _1466_v47 D3
			_ = _1466_v47
			_1466_v47 = Companion_D3_.Create_DC10_(_1465_v46)
			var _1467_v48 _dafny.Map
			_ = _1467_v48
			_1467_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1458_i2, _1466_v47)
			_1467_v48 = (_1467_v48).Update(p2, _1466_v47)
			if ((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int)).Cmp((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int)) < 0 {
				(globalState).F21 = _1458_i2
				(globalState).F28 = (_this).F29()
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_39
				var _nw264 _dafny.Array
				_ = _nw264
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw264 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.Int = (func(_1468_v26 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_1469_i3 _dafny.Int) _dafny.Int {
							return (_1469_i3).Plus((_1468_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1468_v26), 0))).Int()).(_dafny.Int))
						}
					})(_1443_v26)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw264 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw264).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw264).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				(globalState).F17 = _nw264
				_1460_v41 = _1460_v41
				var _rhs260 _dafny.Int = (_1463_v44).Dtor_cf46()
				_ = _rhs260
				var _rhs261 _dafny.CodePoint = _1460_v41
				_ = _rhs261
				var _lhs246 *GlobalState = globalState
				_ = _lhs246
				_lhs246.F19 = _rhs260
				_1460_v41 = _rhs261
			} else {
				(globalState).F15 = !(!(p1) || (false)) || ((_this).F29())
				(globalState).F20 = p0
				(globalState).F14 = p2
				var _1470_v49 *C0
				_ = _1470_v49
				var _nw265 *C0 = New_C0_()
				_ = _nw265
				_nw265.Ctor__(p1)
				_1470_v49 = _nw265
				var _rhs262 *C0 = _1470_v49
				_ = _rhs262
				var _rhs263 bool = (_1457_v39).Contains((p1) && (Companion_Default___.Fm19(_1460_v41, globalState)))
				_ = _rhs263
				var _lhs247 *GlobalState = globalState
				_ = _lhs247
				_1470_v49 = _rhs262
				_lhs247.F28 = _rhs263
				var _1471_v50 _dafny.Array
				_ = _1471_v50
				var _nw266 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw266
				_1471_v50 = _nw266
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1471_v50), 0))
				_ = _index228
				(_1471_v50).ArraySet1((p1) && (_1470_v49.F37), (_index228).Int())
				var _1472_v51 _dafny.Sequence
				_ = _1472_v51
				_1472_v51 = _dafny.SeqOf((_1459_v40).Intersection(_dafny.MultiSetOf(p1, p1, p1)))
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1471_v50), 0))
				_ = _index229
				var _rhs264 bool = _1470_v49.F37
				_ = _rhs264
				var _rhs265 bool = _1470_v49.F37
				_ = _rhs265
				var _rhs266 _dafny.Sequence = _1472_v51
				_ = _rhs266
				var _rhs267 bool = p1
				_ = _rhs267
				var _lhs248 *GlobalState = globalState
				_ = _lhs248
				var _lhs249 _dafny.Array = _1471_v50
				_ = _lhs249
				var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1471_v50), 0))
				_ = _lhs250
				var _lhs251 *GlobalState = globalState
				_ = _lhs251
				_lhs248.F18 = _rhs264
				(_lhs249).ArraySet1(_rhs265, (_lhs250).Int())
				_1472_v51 = _rhs266
				_lhs251.F28 = _rhs267
			}
		}
		var _1473_v52 _dafny.Array
		_ = _1473_v52
		var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw267
		_1473_v52 = _nw267
		var _1474_v53 _dafny.Map
		_ = _1474_v53
		_1474_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1456_v38).Cardinality()), p2)
		var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1473_v52), 0))
		_ = _index230
		(_1473_v52).ArraySet1(_1474_v53, (_index230).Int())
		var _1475_v55 _dafny.CodePoint
		_ = _1475_v55
		_1475_v55 = _dafny.CodePoint('t')
		var _1476_v56 _dafny.Map
		_ = _1476_v56
		_1476_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_1475_v55, globalState), p1)
		var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1473_v52), 0))
		_ = _index231
		(_1473_v52).ArraySet1(func() _dafny.Map {
			var _coll49 = _dafny.NewMapBuilder()
			_ = _coll49
			for _iter60 := _dafny.Iterate((_dafny.MultiSetOf(Companion_Default___.Fm36(globalState), (_1476_v56).Cardinality(), (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-57))).Elements()); ; {
				_compr_49, _ok60 := _iter60()
				if !_ok60 {
					break
				}
				var _1477_v54 _dafny.Int
				_1477_v54 = interface{}(_compr_49).(_dafny.Int)
				if (_dafny.MultiSetOf(Companion_Default___.Fm36(globalState), (_1476_v56).Cardinality(), (_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-57))).Contains(_1477_v54) {
					_coll49.Add(Companion_Default___.SafeDivisionInt(_1477_v54, p2), Companion_Default___.SafeDivisionInt((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(93), p2)).Cardinality()))
				}
			}
			return _coll49.ToMap()
		}(), (_index231).Int())
		r0 = p1
		var _1478_v57 _dafny.MultiSet
		_ = _1478_v57
		_1478_v57 = _dafny.MultiSetOf((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int))
		r1 = Companion_Default___.Fm34((p2).Minus((_1443_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_1443_v26), 0))).Int()).(_dafny.Int)), (_dafny.IntOfInt64(-80)).Plus(p2), (_this).F29(), (_1478_v57).IsSubsetOf(_dafny.MultiSetOf(p2, p2)), globalState)
		return r0, r1
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f30 _dafny.Sequence
	_f29 bool
	_f31 bool
	_f32 bool
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f30 = _dafny.EmptySeq
	_this._f29 = false
	_this._f31 = false
	_this._f32 = false
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

func (_this *C7) F30() _dafny.Sequence {
	return _this._f30
}
func (_this *C7) F30_set_(value _dafny.Sequence) {
	_this._f30 = value
}
func (_this *C7) F29() bool {
	return _this._f29
}
func (_this *C7) Ctor__(f31 bool, f32 bool, f29 bool, f30 _dafny.Sequence) {
	{
		(_this)._f31 = f31
		(_this)._f32 = f32
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C7) M2(p0 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1479_v0 _dafny.Int
		_ = _1479_v0
		_1479_v0 = _dafny.IntOfInt64(601)
		var _1480_v1 _dafny.Array
		_ = _1480_v1
		var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw268
		_1480_v1 = _nw268
		var _1481_v2 _dafny.Map
		_ = _1481_v2
		_1481_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v0, _1480_v1)
		var _1482_v3 _dafny.Array
		_ = _1482_v3
		var _nwElement0_51 _dafny.Int = _1479_v0
		_ = _nwElement0_51
		var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(11))
		_ = _nw269
		(_nw269).ArraySet1(_nwElement0_51, 0)
		(_nw269).ArraySet1(_1479_v0, 1)
		(_nw269).ArraySet1(_1479_v0, 2)
		(_nw269).ArraySet1(_1479_v0, 3)
		(_nw269).ArraySet1(_1479_v0, 4)
		(_nw269).ArraySet1(_1479_v0, 5)
		(_nw269).ArraySet1(_1479_v0, 6)
		(_nw269).ArraySet1((_1481_v2).Cardinality(), 7)
		(_nw269).ArraySet1(_1479_v0, 8)
		(_nw269).ArraySet1(_1479_v0, 9)
		(_nw269).ArraySet1(_1479_v0, 10)
		_1482_v3 = _nw269
		var _1483_v4 _dafny.Set
		_ = _1483_v4
		_1483_v4 = _dafny.SetOf(_1482_v3, _1482_v3, _1480_v1, _1480_v1, _1480_v1)
		var _1484_v5 *C2
		_ = _1484_v5
		var _nw270 *C2 = New_C2_()
		_ = _nw270
		_nw270.Ctor__((_dafny.IntOfInt64(-639)).Plus(_1479_v0), (_1483_v4).Difference(_1483_v4), p0, _dafny.UnicodeSeqOfUtf8Bytes("nicg"))
		_1484_v5 = _nw270
		var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1480_v1), 0))
		_ = _index232
		(_1480_v1).ArraySet1(_1479_v0, (_index232).Int())
		var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1480_v1), 0))
		_ = _index233
		(_1480_v1).ArraySet1(_1479_v0, (_index233).Int())
		var _1485_i0 _dafny.Int
		_ = _1485_i0
		_1485_i0 = _dafny.Zero
		{
			for Companion_Default___.Fm33(false, _dafny.IntOfInt64(689), (_1480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1480_v1), 0))).Int()).(_dafny.Int), globalState) {
				{
					if (_1485_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1485_i0 = (_1485_i0).Plus(_dafny.One)
					var _1486_v6 _dafny.Set
					_ = _1486_v6
					_1486_v6 = _dafny.SetOf(!(p0))
					_1486_v6 = ((func() _dafny.Set {
						if (_this).F29() {
							return _1486_v6
						}
						return _1486_v6
					})()).Difference((_1486_v6).Union(_1486_v6))
					var _1487_v7 _dafny.CodePoint
					_ = _1487_v7
					_1487_v7 = _dafny.CodePoint('u')
					var _1488_v8 _dafny.Map
					_ = _1488_v8
					_1488_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1482_v3, (_this).F31())
					var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1480_v1), 0))
					_ = _index234
					var _rhs268 _dafny.CodePoint = _1487_v7
					_ = _rhs268
					var _rhs269 _dafny.Map = _1488_v8
					_ = _rhs269
					var _rhs270 _dafny.Int = _dafny.IntOfInt64(-103)
					_ = _rhs270
					var _rhs271 _dafny.Int = (_1484_v5).F38()
					_ = _rhs271
					var _rhs272 bool = !(!(Companion_Default___.Fm10((_1484_v5).F38(), (_1484_v5).F38(), globalState)))
					_ = _rhs272
					var _lhs252 _dafny.Array = _1480_v1
					_ = _lhs252
					var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1480_v1), 0))
					_ = _lhs253
					var _lhs254 *GlobalState = globalState
					_ = _lhs254
					_1487_v7 = _rhs268
					_1488_v8 = _rhs269
					(_lhs252).ArraySet1(_rhs270, (_lhs253).Int())
					_lhs254.F21 = _rhs271
					r1 = _rhs272
					var _1489_v9 _dafny.Sequence
					_ = _1489_v9
					_1489_v9 = _dafny.SeqOf(p0)
					var _1490_v10 _dafny.Map
					_ = _1490_v10
					_1490_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())).Cardinality()), _dafny.Companion_Sequence_.Contains(_1489_v9, (_this).F31()))
					var _1491_v11 D12
					_ = _1491_v11
					_1491_v11 = Companion_D12_.Create_DC39_(_dafny.IntOfInt64(460), Companion_Default___.Fm19(_1487_v7, globalState), (_1484_v5).F38())
					_1490_v10 = (_1490_v10).Update((_1491_v11).Dtor_cf65(), (p0) || (p0))
					var _1492_v12 _dafny.Map
					_ = _1492_v12
					_1492_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _1479_v0)
					var _1493_v13 *C4
					_ = _1493_v13
					var _nw271 *C4 = New_C4_()
					_ = _nw271
					_nw271.Ctor__((_this).F32(), (_1484_v5).F38())
					_1493_v13 = _nw271
					var _1494_v14 _dafny.Set
					_ = _1494_v14
					_1494_v14 = _dafny.SetOf(_1493_v13)
					_1492_v12 = (_1492_v12).Update((_this).F29(), ((_1494_v14).Difference(_1494_v14)).Cardinality())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1495_v15 _dafny.Sequence
		_ = _1495_v15
		_1495_v15 = _dafny.SeqOf(_1482_v3, _1480_v1)
		var _1496_v16 _dafny.Sequence
		_ = _1496_v16
		_1496_v16 = _dafny.SeqOf(_1482_v3, (_1495_v15).Select((Companion_Default___.SafeIndex(_1479_v0, _dafny.IntOfUint32((_1495_v15).Cardinality()))).Uint32()).(_dafny.Array), _1482_v3)
		var _1497_v17 _dafny.Array
		_ = _1497_v17
		var _nwElement0_52 _dafny.Sequence = _1496_v16
		_ = _nwElement0_52
		var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(5))
		_ = _nw272
		(_nw272).ArraySet1(_nwElement0_52, 0)
		(_nw272).ArraySet1(_dafny.SeqOf(_1482_v3), 1)
		(_nw272).ArraySet1(_1495_v15, 2)
		(_nw272).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1495_v15, _1495_v15), 3)
		(_nw272).ArraySet1(_1495_v15, 4)
		_1497_v17 = _nw272
		var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1497_v17), 0))
		_ = _index235
		(_1497_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1496_v16, _1495_v15), (_index235).Int())
		var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1497_v17), 0))
		_ = _index236
		(_1497_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1496_v16, _dafny.SeqOf(_1480_v1, _1482_v3)), (_index236).Int())
		var _1498_v18 _dafny.MultiSet
		_ = _1498_v18
		_1498_v18 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1479_v0))
		var _1499_v19 _dafny.Map
		_ = _1499_v19
		_1499_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _1498_v18)
		_1499_v19 = (_1499_v19).Update((_this).F32(), _dafny.MultiSetOf(_1479_v0))
		var _1500_v20 D3
		_ = _1500_v20
		_1500_v20 = Companion_D3_.Create_DC10_((_1498_v18).Union(_1498_v18))
		_1500_v20 = _1500_v20
		var _1501_v21 _dafny.Sequence
		_ = _1501_v21
		_1501_v21 = _dafny.SeqOf(false)
		var _1502_v22 _dafny.Sequence
		_ = _1502_v22
		_1502_v22 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1501_v21, _1501_v21), _dafny.SeqOf((_this).F32()), _1501_v21)
		r0 = _1502_v22
		r1 = p0
		return r0, r1
	}
}
func (_this *C7) M3(p0 bool, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1503_v0 _dafny.Map
		_ = _1503_v0
		_1503_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(258))
		var _1504_v1 _dafny.Map
		_ = _1504_v1
		_1504_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_1503_v0).Cardinality())
		var _1505_v2 _dafny.Sequence
		_ = _1505_v2
		_1505_v2 = _dafny.SeqOf(!((_this).F31()))
		var _1506_v3 _dafny.Sequence
		_ = _1506_v3
		_1506_v3 = _dafny.SeqOf(_1503_v0, _1503_v0, Companion_Default___.Fm14(!(p0), _dafny.SeqOf(_dafny.IntOfUint32((_1505_v2).Cardinality())), (_this).F31(), p2, globalState), _1504_v1, _1503_v0)
		var _1507_v4 _dafny.Array
		_ = _1507_v4
		var _nwElement0_53 bool = !((_1503_v0).Equals((_1506_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1506_v3).Cardinality()))).Uint32()).(_dafny.Map)))
		_ = _nwElement0_53
		var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.One)
		_ = _nw273
		(_nw273).ArraySet1(_nwElement0_53, 0)
		_1507_v4 = _nw273
		var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1507_v4), 0))
		_ = _index237
		(_1507_v4).ArraySet1((_this).F32(), (_index237).Int())
		var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1507_v4), 0))
		_ = _index238
		(_1507_v4).ArraySet1((_this).F31(), (_index238).Int())
		_1507_v4 = _1507_v4
		r1 = (_this).F31()
		var _1508_v5 _dafny.Array
		_ = _1508_v5
		var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw274
		_1508_v5 = _nw274
		var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1508_v5), 0))
		_ = _index239
		(_1508_v5).ArraySet1((_dafny.Zero).Minus(p2), (_index239).Int())
		var _1509_v6 _dafny.MultiSet
		_ = _1509_v6
		_1509_v6 = _dafny.MultiSetOf(p2)
		var _1510_v7 _dafny.MultiSet
		_ = _1510_v7
		_1510_v7 = _dafny.MultiSetOf((_this).F31(), (_this).F31(), (_1505_v2).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm36(globalState), _dafny.IntOfUint32((_1505_v2).Cardinality()))).Uint32()).(bool), true)
		var _1511_v8 _dafny.Map
		_ = _1511_v8
		_1511_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1510_v7).Equals(_1510_v7), !((p2).Cmp(p2) != 0))
		var _1512_v9 D3
		_ = _1512_v9
		_1512_v9 = Companion_D3_.Create_DC10_(_1509_v6)
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1508_v5), 0))
		_ = _index240
		var _rhs273 bool = p0
		_ = _rhs273
		var _rhs274 _dafny.Int = (_1511_v8).Cardinality()
		_ = _rhs274
		var _rhs275 _dafny.MultiSet = (_1512_v9).Dtor_cf18()
		_ = _rhs275
		var _rhs276 _dafny.Int = (p2).Times(p2)
		_ = _rhs276
		var _rhs277 _dafny.Array = _1507_v4
		_ = _rhs277
		var _lhs255 *GlobalState = globalState
		_ = _lhs255
		var _lhs256 _dafny.Array = _1508_v5
		_ = _lhs256
		var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1508_v5), 0))
		_ = _lhs257
		var _lhs258 *GlobalState = globalState
		_ = _lhs258
		_lhs255.F13 = _rhs273
		(_lhs256).ArraySet1(_rhs274, (_lhs257).Int())
		_1509_v6 = _rhs275
		_lhs258.F19 = _rhs276
		_1507_v4 = _rhs277
		var _1513_v10 _dafny.Array
		_ = _1513_v10
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_40
		var _nw275 _dafny.Array
		_ = _nw275
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw275 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.Map = (func(_1514_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
				return func(_1515_i0 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), Companion_D1_.Create_DC3_(_1514_v2))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D1_.Create_DC3_(_1514_v2)))
				}
			})(_1505_v2)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw275 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw275).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw275).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1513_v10 = _nw275
		var _1516_v11 D1
		_ = _1516_v11
		_1516_v11 = Companion_D1_.Create_DC3_(_1505_v2)
		var _1517_v12 _dafny.Map
		_ = _1517_v12
		_1517_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _1516_v11)
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1513_v10), 0))
		_ = _index241
		(_1513_v10).ArraySet1(_1517_v12, (_index241).Int())
		var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1513_v10), 0))
		_ = _index242
		(_1513_v10).ArraySet1(_1517_v12, (_index242).Int())
		var _1518_v13 _dafny.Array
		_ = _1518_v13
		var _nwElement0_54 _dafny.Array = p1
		_ = _nwElement0_54
		var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(5))
		_ = _nw276
		(_nw276).ArraySet1(_nwElement0_54, 0)
		(_nw276).ArraySet1(p1, 1)
		(_nw276).ArraySet1(p1, 2)
		(_nw276).ArraySet1(p1, 3)
		(_nw276).ArraySet1(p1, 4)
		_1518_v13 = _nw276
		var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_1518_v13), 0))
		_ = _index243
		(_1518_v13).ArraySet1(p1, (_index243).Int())
		var _1519_v14 _dafny.Map
		_ = _1519_v14
		_1519_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), _dafny.IntOfInt64(-828))
		var _1520_v16 _dafny.CodePoint
		_ = _1520_v16
		_1520_v16 = _dafny.CodePoint('d')
		var _1521_v17 _dafny.Set
		_ = _1521_v17
		_1521_v17 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(371))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg96 _dafny.Int) interface{} {
				return coer96(arg96)
			}
		}(func(_1522_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(371))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg97 _dafny.Int) interface{} {
				return coer97(arg97)
			}
		}(func(_1523_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()))).Uint32(), _1520_v16), _dafny.UnicodeSeqOfUtf8Bytes("mmynhama"))
		var _1524_v18 D5
		_ = _1524_v18
		_1524_v18 = Companion_D5_.Create_DC17_(_1521_v17)
		var _1525_v19 D5
		_ = _1525_v19
		_1525_v19 = Companion_D5_.Create_DC20_(_1524_v18)
		var _pat_let_tv29 = p2
		_ = _pat_let_tv29
		var _pat_let_tv30 = _1508_v5
		_ = _pat_let_tv30
		var _pat_let_tv31 = _1508_v5
		_ = _pat_let_tv31
		var _pat_let_tv32 = _1508_v5
		_ = _pat_let_tv32
		var _pat_let_tv33 = _1508_v5
		_ = _pat_let_tv33
		var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_1518_v13), 0))
		_ = _index244
		var _rhs278 _dafny.Array = p1
		_ = _rhs278
		var _rhs279 _dafny.Map = func() _dafny.Map {
			var _coll50 = _dafny.NewMapBuilder()
			_ = _coll50
			for _iter61 := _dafny.Iterate((_1521_v17).Elements()); ; {
				_compr_50, _ok61 := _iter61()
				if !_ok61 {
					break
				}
				var _1526_v15 _dafny.Sequence
				_1526_v15 = interface{}(_compr_50).(_dafny.Sequence)
				if (_1521_v17).Contains(_1526_v15) {
					_coll50.Add(_1526_v15, (_1509_v6).Cardinality())
				}
			}
			return _coll50.ToMap()
		}()
		_ = _rhs279
		var _rhs280 _dafny.Int = func(_source21 D5) _dafny.Int {
			if _source21.Is_DC17() {
				var _1527___mcc_h0 _dafny.Set = _source21.Get_().(D5_DC17).Cf27
				_ = _1527___mcc_h0
				var _1528_cf27 _dafny.Set = _1527___mcc_h0
				_ = _1528_cf27
				return _pat_let_tv29
			} else if _source21.Is_DC18() {
				return _dafny.IntOfUint32((_this.F30()).Cardinality())
			} else if _source21.Is_DC19() {
				var _1529___mcc_h1 _dafny.Int = _source21.Get_().(D5_DC19).Cf28
				_ = _1529___mcc_h1
				var _1530___mcc_h2 _dafny.Int = _source21.Get_().(D5_DC19).Cf29
				_ = _1530___mcc_h2
				var _1531___mcc_h3 _dafny.CodePoint = _source21.Get_().(D5_DC19).Cf30
				_ = _1531___mcc_h3
				var _1532_cf30 _dafny.CodePoint = _1531___mcc_h3
				_ = _1532_cf30
				var _1533_cf29 _dafny.Int = _1530___mcc_h2
				_ = _1533_cf29
				var _1534_cf28 _dafny.Int = _1529___mcc_h1
				_ = _1534_cf28
				return _dafny.IntOfInt64(538)
			} else if _source21.Is_DC16() {
				var _1535___mcc_h4 _dafny.Map = _source21.Get_().(D5_DC16).Cf26
				_ = _1535___mcc_h4
				var _1536_cf26 _dafny.Map = _1535___mcc_h4
				_ = _1536_cf26
				return (_pat_let_tv31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_pat_let_tv30), 0))).Int()).(_dafny.Int)
			} else {
				var _1537___mcc_h5 D5 = _source21.Get_().(D5_DC20).Cf31
				_ = _1537___mcc_h5
				var _1538_cf31 D5 = _1537___mcc_h5
				_ = _1538_cf31
				return (_pat_let_tv33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_pat_let_tv32), 0))).Int()).(_dafny.Int)
			}
		}(_1525_v19)
		_ = _rhs280
		var _lhs259 _dafny.Array = _1518_v13
		_ = _lhs259
		var _lhs260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_1518_v13), 0))
		_ = _lhs260
		var _lhs261 *GlobalState = globalState
		_ = _lhs261
		(_lhs259).ArraySet1(_rhs278, (_lhs260).Int())
		_1519_v14 = _rhs279
		_lhs261.F21 = _rhs280
		var _1539_v20 _dafny.Sequence
		_ = _1539_v20
		_1539_v20 = _dafny.SeqOf(_this.F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-759))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg98 _dafny.Int) interface{} {
				return coer98(arg98)
			}
		}((func(_1540_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1541_i2 _dafny.Int) _dafny.CodePoint {
				return _1540_v16
			}
		})(_1520_v16))), (func() _dafny.Sequence {
			if true {
				return _this.F30()
			}
			return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _dafny.CodePoint('i')), (Companion_Default___.SafeIndex((_1504_v1).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F30(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F30()).Cardinality()))).Uint32(), _dafny.CodePoint('i'))).Cardinality()))).Uint32(), _1520_v16)
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("o"), _dafny.UnicodeSeqOfUtf8Bytes("mjfyn")), _this.F30())
		r0 = (_1539_v20).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1539_v20).Cardinality()))).Uint32()).(_dafny.Sequence)
		r1 = ((_dafny.Zero).Minus((_1508_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1508_v5), 0))).Int()).(_dafny.Int))).Cmp((_dafny.Zero).Minus((p2).Times(_dafny.IntOfInt64(89)))) >= 0
		return r0, r1
	}
}
func (_this *C7) M4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int, _dafny.Set) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		var _1542_v0 _dafny.Array
		_ = _1542_v0
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_41
		var _nw277 _dafny.Array
		_ = _nw277
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw277 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) bool = func(_1543_i0 _dafny.Int) bool {
				return (_this).F29()
			}
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw277 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw277).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw277).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		_1542_v0 = _nw277
		var _1544_v1 _dafny.Set
		_ = _1544_v1
		_1544_v1 = _dafny.SetOf(_1542_v0, _1542_v0)
		var _1545_v2 *C4
		_ = _1545_v2
		var _nw278 *C4 = New_C4_()
		_ = _nw278
		_nw278.Ctor__(!((Companion_D16_.Create_DC47_(_1544_v1)).Dtor_cf81()).Contains(_1542_v0), p0)
		_1545_v2 = _nw278
		var _1546_v3 D4
		_ = _1546_v3
		_1546_v3 = Companion_D4_.Create_DC15_(p0, p0)
		var _1547_v4 _dafny.CodePoint
		_ = _1547_v4
		_1547_v4 = _dafny.CodePoint('i')
		var _1548_v5 D10
		_ = _1548_v5
		_1548_v5 = Companion_D10_.Create_DC32_(_1547_v4, (_this).F32())
		var _1549_v6 _dafny.Array
		_ = _1549_v6
		var _nwElement0_55 D10 = _1548_v5
		_ = _nwElement0_55
		var _nw279 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(9))
		_ = _nw279
		(_nw279).ArraySet1(_nwElement0_55, 0)
		(_nw279).ArraySet1(_1548_v5, 1)
		(_nw279).ArraySet1(_1548_v5, 2)
		(_nw279).ArraySet1(Companion_D10_.Create_DC32_(_1547_v4, (_this).F29()), 3)
		(_nw279).ArraySet1(_1548_v5, 4)
		(_nw279).ArraySet1(_1548_v5, 5)
		(_nw279).ArraySet1(_1548_v5, 6)
		(_nw279).ArraySet1(_1548_v5, 7)
		(_nw279).ArraySet1(Companion_D10_.Create_DC32_(_1547_v4, (_this).F32()), 8)
		_1549_v6 = _nw279
		var _1550_v7 _dafny.Array
		_ = _1550_v7
		var _nw280 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw280
		_1550_v7 = _nw280
		var _1551_v8 _dafny.Map
		_ = _1551_v8
		_1551_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F30()).Cardinality()), !((_1545_v2).F33()))
		var _1552_v9 D15
		_ = _1552_v9
		_1552_v9 = Companion_D15_.Create_DC44_((_1546_v3).Dtor_cf25(), (_this).F32(), _1549_v6, _1550_v7, (_1551_v8).Cardinality())
		var _1553_v10 D3
		_ = _1553_v10
		_1553_v10 = Companion_D3_.Create_DC11_((_1545_v2).F33(), (_this).F32())
		var _1554_v11 _dafny.MultiSet
		_ = _1554_v11
		_1554_v11 = _dafny.MultiSetOf((_1545_v2).F34())
		var _1555_v12 D11
		_ = _1555_v12
		_1555_v12 = Companion_D11_.Create_DC37_((_1552_v9).Dtor_cf75(), Companion_Default___.Fm33(!((Companion_D3_.Create_DC11_((_1553_v10).Dtor_cf19(), (_this).F32())).Dtor_cf19()), p0, (_1554_v11).Cardinality(), globalState))
		var _source22 D11 = _1555_v12
		_ = _source22
		if _source22.Is_DC35() {
			var _1556___mcc_h0 _dafny.Sequence = _source22.Get_().(D11_DC35).Cf57
			_ = _1556___mcc_h0
			var _1557___mcc_h1 bool = _source22.Get_().(D11_DC35).Cf58
			_ = _1557___mcc_h1
			var _1558_cf58 bool = _1557___mcc_h1
			_ = _1558_cf58
			var _1559_cf57 _dafny.Sequence = _1556___mcc_h0
			_ = _1559_cf57
			(globalState).F15 = _1558_cf58
			if (p1).Cmp(p1) >= 0 {
				(globalState).F2 = (Companion_Default___.Fm36(globalState)).Minus(Companion_Default___.SafeDivisionInt((_1545_v2).F34(), (_1545_v2).F34()))
				var _1560_v13 _dafny.Set
				_ = _1560_v13
				_1560_v13 = _dafny.SetOf((_1545_v2).F34(), (_1545_v2).F34())
				var _1561_v14 D17
				_ = _1561_v14
				_1561_v14 = Companion_D17_.Create_DC49_(_1560_v13)
				var _1562_v15 _dafny.Sequence
				_ = _1562_v15
				_1562_v15 = _dafny.SeqOf((_1545_v2).F33())
				var _1563_v16 _dafny.Map
				_ = _1563_v16
				_1563_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F34(), (_1545_v2).F34())
				var _1564_v17 _dafny.Map
				_ = _1564_v17
				_1564_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), _1560_v13)
				var _1565_v19 _dafny.Array
				_ = _1565_v19
				var _nwElement0_56 _dafny.Set = _dafny.SetOf((_1545_v2).F34(), (_1545_v2).F34(), p1, p0, (_1545_v2).F34())
				_ = _nwElement0_56
				var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(20))
				_ = _nw281
				(_nw281).ArraySet1(_nwElement0_56, 0)
				(_nw281).ArraySet1((_1560_v13).Intersection(_1560_v13), 1)
				(_nw281).ArraySet1(_1560_v13, 2)
				(_nw281).ArraySet1((_1561_v14).Dtor_cf82(), 3)
				(_nw281).ArraySet1(_1560_v13, 4)
				(_nw281).ArraySet1(_1560_v13, 5)
				(_nw281).ArraySet1((_1560_v13).Intersection(_1560_v13), 6)
				(_nw281).ArraySet1((_1560_v13).Difference(_1560_v13), 7)
				(_nw281).ArraySet1(_1560_v13, 8)
				(_nw281).ArraySet1(_1560_v13, 9)
				(_nw281).ArraySet1(_1560_v13, 10)
				(_nw281).ArraySet1(_1560_v13, 11)
				(_nw281).ArraySet1(_1560_v13, 12)
				(_nw281).ArraySet1(_1560_v13, 13)
				(_nw281).ArraySet1((func() _dafny.Set {
					if (_1545_v2).F33() {
						return Companion_Default___.Fm17((_1545_v2).F34(), globalState)
					}
					return _1560_v13
				})(), 14)
				(_nw281).ArraySet1((_1560_v13).Difference(_1560_v13), 15)
				(_nw281).ArraySet1(_1560_v13, 16)
				(_nw281).ArraySet1((_dafny.SetOf((_1545_v2).F34(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(730))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}((func(_1566_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1567_i1 _dafny.Int) _dafny.CodePoint {
						return _1566_v4
					}
				})(_1547_v4))), (Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(730))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1568_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1569_i1 _dafny.Int) _dafny.CodePoint {
						return _1568_v4
					}
				})(_1547_v4)))).Cardinality()))).Uint32(), _1547_v4)).Cardinality()), _dafny.IntOfUint32((_1562_v15).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), _1563_v16)).Cardinality(), p1)).Difference((func() _dafny.Set {
					if (_1564_v17).Contains(_this.F30()) {
						return (_1564_v17).Get(_this.F30()).(_dafny.Set)
					}
					return _dafny.SetOf(_dafny.IntOfUint32((_this.F30()).Cardinality()))
				})()), 17)
				(_nw281).ArraySet1((func() _dafny.Set {
					var _coll51 = _dafny.NewBuilder()
					_ = _coll51
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(735), _dafny.IntOfInt64(598))); ; {
						_compr_51, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1570_v18 _dafny.Int
						_1570_v18 = interface{}(_compr_51).(_dafny.Int)
						if ((_dafny.IntOfInt64(735)).Cmp(_1570_v18) <= 0) && ((_1570_v18).Cmp(_dafny.IntOfInt64(598)) < 0) {
							_coll51.Add(Companion_Default___.SafeDivisionInt(_1570_v18, p0))
						}
					}
					return _coll51.ToSet()
				}()).Difference(_1560_v13), 18)
				(_nw281).ArraySet1(_1560_v13, 19)
				_1565_v19 = _nw281
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1565_v19), 0))
				_ = _index245
				(_1565_v19).ArraySet1((_1560_v13).Intersection(_1560_v13), (_index245).Int())
				var _1571_v20 _dafny.Map
				_ = _1571_v20
				_1571_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1548_v5).Dtor_cf52(), (_1560_v13).Difference(_1560_v13))
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1565_v19), 0))
				_ = _index246
				(_1565_v19).ArraySet1((func() _dafny.Set {
					if (_1571_v20).Contains(_dafny.CodePoint('u')) {
						return (_1571_v20).Get(_dafny.CodePoint('u')).(_dafny.Set)
					}
					return _1560_v13
				})(), (_index246).Int())
				var _1572_v21 _dafny.Map
				_ = _1572_v21
				_1572_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v0, _1547_v4)
				var _1573_v22 D0
				_ = _1573_v22
				_1573_v22 = Companion_D0_.Create_DC1_(_1542_v0, p0, _dafny.IntOfUint32((_1559_cf57).Cardinality()))
				_1572_v21 = (_1572_v21).Update((_1573_v22).Dtor_cf1(), _1547_v4)
				var _1574_v23 *C5
				_ = _1574_v23
				var _nw282 *C5 = New_C5_()
				_ = _nw282
				_nw282.Ctor__((func() bool {
					if (_this).F32() {
						return (_this).F32()
					}
					return (_this).F31()
				})(), _1559_cf57)
				_1574_v23 = _nw282
				var _1575_v24 _dafny.MultiSet
				_ = _1575_v24
				_1575_v24 = _dafny.MultiSetOf((_1545_v2).F33(), (_1545_v2).F33())
				var _1576_v25 _dafny.Sequence
				_ = _1576_v25
				_1576_v25 = _dafny.SeqOf((_1575_v24).Cardinality())
				var _1577_v26 *C1
				_ = _1577_v26
				var _nw283 *C1 = New_C1_()
				_ = _nw283
				_nw283.Ctor__(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm44((_1545_v2).F33(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1578_v2 *C4) func(_dafny.Int) _dafny.Int {
					return func(_1579_i2 _dafny.Int) _dafny.Int {
						return (_1578_v2).F34()
					}
				})(_1545_v2)))).Cardinality()), (_1545_v2).F33(), (_1545_v2).F33(), globalState), _1576_v25))
				_1577_v26 = _nw283
			} else {
				(globalState).F13 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("dnewr")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wf"), _dafny.UnicodeSeqOfUtf8Bytes("efgrsogq")))
				(globalState).F15 = (_1545_v2).F33()
				var _1580_v27 _dafny.Array
				_ = _1580_v27
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(23))
				_ = _nw284
				_1580_v27 = _nw284
				var _1581_v28 D12
				_ = _1581_v28
				_1581_v28 = Companion_D12_.Create_DC39_(p1, (_1545_v2).F33(), _dafny.IntOfInt64(146))
				var _1582_v29 D2
				_ = _1582_v29
				_1582_v29 = Companion_D2_.Create_DC9_(_this.F30(), p1, _1558_cf58, (_1581_v28).Dtor_cf67())
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1580_v27), 0))
				_ = _index247
				(_1580_v27).ArraySet1(_1582_v29, (_index247).Int())
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1580_v27), 0))
				_ = _index248
				(_1580_v27).ArraySet1(_1582_v29, (_index248).Int())
				var _1583_v30 D17
				_ = _1583_v30
				_1583_v30 = Companion_D17_.Create_DC52_(Companion_Default___.Fm6(_dafny.IntOfUint32((_1559_cf57).Cardinality()), globalState), (_1545_v2).F33(), (_1545_v2).F34(), Companion_Default___.Fm27((_1545_v2).F34(), true, globalState))
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index249
				(_1550_v7).ArraySet1(Companion_Default___.SafeModuloInt((_1583_v30).Dtor_cf90(), (_1583_v30).Dtor_cf90()), (_index249).Int())
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index250
				(_1550_v7).ArraySet1(_dafny.IntOfInt64(-602), (_index250).Int())
				var _1584_v32 _dafny.Map
				_ = _1584_v32
				_1584_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(622), (func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter63 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(190))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg102 _dafny.Int) interface{} {
							return coer102(arg102)
						}
					}((func(_1585_v2 *C4) func(_dafny.Int) _dafny.Int {
						return func(_1586_i3 _dafny.Int) _dafny.Int {
							return (_1585_v2).F34()
						}
					})(_1545_v2)))).Elements()); ; {
						_compr_52, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1587_v31 _dafny.Int
						_1587_v31 = interface{}(_compr_52).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(190))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg103 _dafny.Int) interface{} {
								return coer103(arg103)
							}
						}((func(_1588_v2 *C4) func(_dafny.Int) _dafny.Int {
							return func(_1586_i3 _dafny.Int) _dafny.Int {
								return (_1588_v2).F34()
							}
						})(_1545_v2))), _1587_v31) {
							_coll52.Add(Companion_Default___.SafeModuloInt(_1587_v31, (_1545_v2).F34()), (_1545_v2).F34())
						}
					}
					return _coll52.ToMap()
				}()).Cardinality())
				(globalState).F19 = ((_1584_v32).Cardinality()).Times((_1545_v2).F34())
			}
			var _1589_v33 _dafny.Sequence
			_ = _1589_v33
			_1589_v33 = _dafny.SeqOf(p0)
			if !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(301))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}(func(_1590_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(899)
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-501))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1591_v2 *C4) func(_dafny.Int) _dafny.Int {
				return func(_1592_i5 _dafny.Int) _dafny.Int {
					return (_1591_v2).F34()
				}
			})(_1545_v2)))), _1589_v33)) {
				_1547_v4 = _1547_v4
				(globalState).F13 = !((_1558_cf58) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_1559_cf57, _this.F30())))
				_1554_v11 = (_1554_v11).Intersection(_1554_v11)
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index251
				(_1550_v7).ArraySet1(p1, (_index251).Int())
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index252
				var _rhs281 _dafny.Int = (_1589_v33).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(847)).Minus(_dafny.IntOfUint32((_1589_v33).Cardinality())), _dafny.IntOfUint32((_1589_v33).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs281
				var _rhs282 bool = (func() bool {
					if (_1545_v2).F33() {
						return (_1545_v2).F33()
					}
					return _1558_cf58
				})()
				_ = _rhs282
				var _rhs283 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1545_v2).F33() {
						return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm34(p1, (_1545_v2).F34(), (_this).F31(), !(Companion_Default___.Fm33((_this).F29(), p1, _dafny.IntOfInt64(848), globalState)), globalState), _1559_cf57)
					}
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("c"), _this.F30())
				})()
				_ = _rhs283
				var _rhs284 _dafny.Int = (_1545_v2).F34()
				_ = _rhs284
				var _rhs285 _dafny.Array = _1550_v7
				_ = _rhs285
				var _lhs262 _dafny.Array = _1550_v7
				_ = _lhs262
				var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1550_v7), 0))
				_ = _lhs263
				var _lhs264 *GlobalState = globalState
				_ = _lhs264
				var _lhs265 *C7 = _this
				_ = _lhs265
				var _lhs266 *GlobalState = globalState
				_ = _lhs266
				var _lhs267 *GlobalState = globalState
				_ = _lhs267
				(_lhs262).ArraySet1(_rhs281, (_lhs263).Int())
				_lhs264.F18 = _rhs282
				_lhs265.F30_set_(_rhs283)
				_lhs266.F19 = _rhs284
				_lhs267.F17 = _rhs285
				(globalState).F1 = (_1545_v2).F34()
			} else {
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index253
				(_1542_v0).ArraySet1(false, (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index254
				(_1542_v0).ArraySet1((p1).Cmp((_1545_v2).F34()) == 0, (_index254).Int())
				var _1593_v34 _dafny.Set
				_ = _1593_v34
				_1593_v34 = _dafny.SetOf((_1545_v2).F33(), (_this).F31())
				var _1594_v35 _dafny.Sequence
				_ = _1594_v35
				_1594_v35 = _dafny.SeqOf((_this).F31(), _1558_cf58)
				var _1595_v36 _dafny.Map
				_ = _1595_v36
				_1595_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), Companion_Default___.Fm20((_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool), (_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool), _1594_v35, (_1545_v2).F34(), globalState))
				var _rhs286 bool = (_1554_v11).IsProperSubsetOf(_1554_v11)
				_ = _rhs286
				var _rhs287 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F30(), Companion_Default___.Fm34((_1593_v34).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1595_v36).Contains(false) {
						return (_1595_v36).Get(false).(_dafny.Sequence)
					}
					return _1559_cf57
				})()).Cardinality()), (_this).F31(), (_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool), globalState))
				_ = _rhs287
				var _rhs288 _dafny.Int = (_1589_v33).Select((Companion_Default___.SafeIndex(((_1545_v2).F34()).Plus((_1545_v2).F34()), _dafny.IntOfUint32((_1589_v33).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs288
				var _rhs289 _dafny.Int = p1
				_ = _rhs289
				var _lhs268 *GlobalState = globalState
				_ = _lhs268
				var _lhs269 *GlobalState = globalState
				_ = _lhs269
				var _lhs270 *GlobalState = globalState
				_ = _lhs270
				var _lhs271 *GlobalState = globalState
				_ = _lhs271
				_lhs268.F18 = _rhs286
				_lhs269.F11 = _rhs287
				_lhs270.F1 = _rhs288
				_lhs271.F14 = _rhs289
				var _1596_v37 _dafny.Set
				_ = _1596_v37
				_1596_v37 = _dafny.SetOf(_1550_v7)
				var _1597_v38 _dafny.Sequence
				_ = _1597_v38
				_1597_v38 = _dafny.SeqOf(_1596_v37)
				var _1598_v39 *C2
				_ = _1598_v39
				var _nw285 *C2 = New_C2_()
				_ = _nw285
				_nw285.Ctor__((_1545_v2).F34(), (_1597_v38).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1597_v38).Cardinality()))).Uint32()).(_dafny.Set), (_this).F29(), _this.F30())
				_1598_v39 = _nw285
				var _1599_v40 _dafny.MultiSet
				_ = _1599_v40
				_1599_v40 = _dafny.MultiSetOf(_1550_v7, _1550_v7)
				var _1600_v41 _dafny.Array
				_ = _1600_v41
				var _nw286 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw286
				_1600_v41 = _nw286
				var _1601_v42 _dafny.MultiSet
				_ = _1601_v42
				_1601_v42 = _dafny.MultiSetOf(_1600_v41, _1600_v41)
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index255
				(_1542_v0).ArraySet1(((_1599_v40).Union((_1599_v40).Update(_1550_v7, Companion_Default___.Abs((_1601_v42).Cardinality())))).IsSubsetOf((_dafny.MultiSetOf(_1550_v7)).Union(_1599_v40)), (_index255).Int())
				(globalState).F2 = Companion_Default___.Fm9(globalState)
			}
			var _1602_v43 *C3
			_ = _1602_v43
			var _nw287 *C3 = New_C3_()
			_ = _nw287
			_nw287.Ctor__(true, ((_1545_v2).F34()).Times(p1), (_this).F32(), _1559_cf57)
			_1602_v43 = _nw287
		} else if _source22.Is_DC36() {
			var _1603___mcc_h2 D6 = _source22.Get_().(D11_DC36).Cf59
			_ = _1603___mcc_h2
			var _1604___mcc_h3 _dafny.Set = _source22.Get_().(D11_DC36).Cf60
			_ = _1604___mcc_h3
			var _1605___mcc_h4 bool = _source22.Get_().(D11_DC36).Cf61
			_ = _1605___mcc_h4
			var _1606_cf61 bool = _1605___mcc_h4
			_ = _1606_cf61
			var _1607_cf60 _dafny.Set = _1604___mcc_h3
			_ = _1607_cf60
			var _1608_cf59 D6 = _1603___mcc_h2
			_ = _1608_cf59
			var _1609_v44 _dafny.Sequence
			_ = _1609_v44
			_1609_v44 = _dafny.SeqOf((_1545_v2).F33(), true, (func() bool {
				if (_this).F32() {
					return (_1545_v2).F33()
				}
				return (_1545_v2).F33()
			})(), (_this).F29())
			(globalState).F13 = (_1609_v44).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm9(globalState), _dafny.IntOfUint32((_1609_v44).Cardinality()))).Uint32()).(bool)
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index256
			(_1550_v7).ArraySet1((_1545_v2).F34(), (_index256).Int())
			var _1610_v45 _dafny.Set
			_ = _1610_v45
			_1610_v45 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(250))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}((func(_1611_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1612_i6 _dafny.Int) _dafny.CodePoint {
					return _1611_v4
				}
			})(_1547_v4)))).Cardinality()))
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index257
			(_1550_v7).ArraySet1((p0).Plus(((_1610_v45).Cardinality()).Minus(p1)), (_index257).Int())
			var _1613_v46 D3
			_ = _1613_v46
			_1613_v46 = Companion_D3_.Create_DC10_(_1554_v11)
			var _1614_v47 D3
			_ = _1614_v47
			_1614_v47 = Companion_D3_.Create_DC12_(_1613_v46)
			var _source23 D3 = (func() D3 {
				if (_this).F29() {
					return _1614_v47
				}
				return _1614_v47
			})()
			_ = _source23
			if _source23.Is_DC11() {
				var _1615___mcc_h8 bool = _source23.Get_().(D3_DC11).Cf19
				_ = _1615___mcc_h8
				var _1616___mcc_h9 bool = _source23.Get_().(D3_DC11).Cf20
				_ = _1616___mcc_h9
				var _1617_cf20 bool = _1616___mcc_h9
				_ = _1617_cf20
				var _1618_cf19 bool = _1615___mcc_h8
				_ = _1618_cf19
				var _1619_v48 *C1
				_ = _1619_v48
				var _nw288 *C1 = New_C1_()
				_ = _nw288
				_nw288.Ctor__((_this).F29())
				_1619_v48 = _nw288
				var _1620_v49 _dafny.Sequence
				_ = _1620_v49
				_1620_v49 = _dafny.SeqOf(_1619_v48, _1619_v48)
				var _1621_v50 _dafny.Map
				_ = _1621_v50
				_1621_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1617_cf20, _dafny.MultiSetOf(_1619_v48, _1619_v48, (_1620_v49).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1620_v49).Cardinality()))).Uint32()).(*C1), _1619_v48))
				var _1622_v51 _dafny.Map
				_ = _1622_v51
				_1622_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1621_v50, (func() bool {
					if (_1551_v8).Contains((_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int)) {
						return (_1551_v8).Get((_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int)).(bool)
					}
					return !(false)
				})())
				var _1623_v52 _dafny.MultiSet
				_ = _1623_v52
				_1623_v52 = _dafny.MultiSetOf(_1619_v48)
				var _1624_v53 _dafny.Sequence
				_ = _1624_v53
				_1624_v53 = _dafny.SeqOf(p1, (_dafny.Zero).Minus((_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int)), (_dafny.IntOfUint32((_1609_v44).Cardinality())).Plus((_dafny.SetOf((_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int), (_1545_v2).F34(), p0)).Cardinality()), p1)
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index258
				var _rhs290 bool = (func() bool {
					if (_1622_v51).Contains((_1621_v50).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1606_cf61, _1623_v52))) {
						return (_1622_v51).Get((_1621_v50).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1606_cf61, _1623_v52))).(bool)
					}
					return ((_this).F31()) == (_1618_cf19)
				})()
				_ = _rhs290
				var _rhs291 _dafny.Int = (Companion_Default___.SafeModuloInt((_1545_v2).F34(), (_1545_v2).F34())).Minus((_1545_v2).F34())
				_ = _rhs291
				var _rhs292 _dafny.Sequence = _1624_v53
				_ = _rhs292
				var _lhs272 *GlobalState = globalState
				_ = _lhs272
				var _lhs273 _dafny.Array = _1550_v7
				_ = _lhs273
				var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))
				_ = _lhs274
				var _lhs275 *GlobalState = globalState
				_ = _lhs275
				_lhs272.F28 = _rhs290
				(_lhs273).ArraySet1(_rhs291, (_lhs274).Int())
				_lhs275.F7 = _rhs292
				var _1625_v54 _dafny.Sequence
				_ = _1625_v54
				_1625_v54 = _dafny.SeqOf(_1614_v47, Companion_Default___.Fm29((_1545_v2).F33(), (_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int), _1606_cf61, globalState))
				var _rhs293 bool = !(_1617_cf20) || (_1617_cf20)
				_ = _rhs293
				var _rhs294 bool = (_dafny.IntOfInt64(847)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1625_v54, _1625_v54)).Cardinality())) != 0
				_ = _rhs294
				var _rhs295 _dafny.Set = (_1607_cf60).Union((func() _dafny.Set {
					if _1617_cf20 {
						return _1607_cf60
					}
					return _1607_cf60
				})())
				_ = _rhs295
				var _rhs296 _dafny.Int = (p0).Times((_1545_v2).F34())
				_ = _rhs296
				var _lhs276 *GlobalState = globalState
				_ = _lhs276
				var _lhs277 *GlobalState = globalState
				_ = _lhs277
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				_lhs276.F15 = _rhs293
				_lhs277.F18 = _rhs294
				_1607_cf60 = _rhs295
				_lhs278.F14 = _rhs296
				var _1626_v55 _dafny.Array
				_ = _1626_v55
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_42
				var _nw289 _dafny.Array
				_ = _nw289
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw289 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.CodePoint = (func(_1627_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1628_i7 _dafny.Int) _dafny.CodePoint {
							return _1627_v4
						}
					})(_1547_v4)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw289 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw289).ArraySet1CodePoint(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw289).ArraySet1CodePoint(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1626_v55 = _nw289
				var _rhs297 _dafny.Array = _1626_v55
				_ = _rhs297
				var _rhs298 bool = !(_1606_cf61)
				_ = _rhs298
				var _rhs299 _dafny.Int = ((func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(228), _dafny.IntOfInt64(-272))); ; {
						_compr_53, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _1629_v56 _dafny.Int
						_1629_v56 = interface{}(_compr_53).(_dafny.Int)
						if ((_dafny.IntOfInt64(228)).Cmp(_1629_v56) <= 0) && ((_1629_v56).Cmp(_dafny.IntOfInt64(-272)) < 0) {
							_coll53.Add((_1629_v56).Plus(p1), (_1619_v48).F40())
						}
					}
					return _coll53.ToMap()
				}()).Merge((Companion_Default___.Fm41(p0, (_this).F29(), _1606_cf61, (_1545_v2).F33(), globalState)).Merge(_1551_v8))).Cardinality()
				_ = _rhs299
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				var _lhs280 *GlobalState = globalState
				_ = _lhs280
				r1 = _rhs297
				_lhs279.F13 = _rhs298
				_lhs280.F2 = _rhs299
				(globalState).F15 = _1618_cf19
			} else if _source23.Is_DC10() {
				var _1630___mcc_h10 _dafny.MultiSet = _source23.Get_().(D3_DC10).Cf18
				_ = _1630___mcc_h10
				var _1631_cf18 _dafny.MultiSet = _1630___mcc_h10
				_ = _1631_cf18
				var _1632_v57 _dafny.MultiSet
				_ = _1632_v57
				_1632_v57 = _dafny.MultiSetOf(_1547_v4)
				(globalState).F21 = (((_1632_v57).Intersection(_1632_v57)).Difference(_dafny.MultiSetOf(_1547_v4))).Cardinality()
				var _rhs300 _dafny.Int = ((_dafny.Zero).Minus((_1545_v2).F34())).Times((_1545_v2).F34())
				_ = _rhs300
				var _rhs301 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())
				_ = _rhs301
				var _rhs302 _dafny.Int = (_1545_v2).Fm5(_dafny.SetOf((_this).F31(), _1606_cf61), globalState)
				_ = _rhs302
				var _rhs303 bool = Companion_Default___.Fm33(((_1545_v2).F34()).Cmp((_1545_v2).F34()) <= 0, p1, (_1545_v2).F34(), globalState)
				_ = _rhs303
				var _rhs304 _dafny.Int = _dafny.IntOfInt64(-554)
				_ = _rhs304
				var _lhs281 *GlobalState = globalState
				_ = _lhs281
				var _lhs282 *GlobalState = globalState
				_ = _lhs282
				var _lhs283 *GlobalState = globalState
				_ = _lhs283
				var _lhs284 *GlobalState = globalState
				_ = _lhs284
				var _lhs285 *GlobalState = globalState
				_ = _lhs285
				_lhs281.F2 = _rhs300
				_lhs282.F11 = _rhs301
				_lhs283.F21 = _rhs302
				_lhs284.F28 = _rhs303
				_lhs285.F21 = _rhs304
				var _rhs305 _dafny.Array = _1542_v0
				_ = _rhs305
				var _rhs306 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("uef"))
				_ = _rhs306
				var _lhs286 *GlobalState = globalState
				_ = _lhs286
				_1542_v0 = _rhs305
				_lhs286.F11 = _rhs306
				var _1633_v58 _dafny.Array
				_ = _1633_v58
				var _nw290 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
				_ = _nw290
				_1633_v58 = _nw290
				var _1634_v59 _dafny.Array
				_ = _1634_v59
				var _nw291 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw291
				_1634_v59 = _nw291
				var _1635_v60 _dafny.Set
				_ = _1635_v60
				_1635_v60 = _dafny.SetOf(_1634_v59)
				var _1636_v61 *C2
				_ = _1636_v61
				var _nw292 *C2 = New_C2_()
				_ = _nw292
				_nw292.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1609_v44, (_this).F31())).Cardinality(), _1635_v60, (_this).F31(), _this.F30())
				_1636_v61 = _nw292
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_1633_v58), 0))
				_ = _index259
				(_1633_v58).ArraySet1(_1636_v61, (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_1633_v58), 0))
				_ = _index260
				(_1633_v58).ArraySet1(_1636_v61, (_index260).Int())
			} else {
				var _1637___mcc_h11 D3 = _source23.Get_().(D3_DC12).Cf21
				_ = _1637___mcc_h11
				var _1638_cf21 D3 = _1637___mcc_h11
				_ = _1638_cf21
				var _1639_v62 _dafny.Array
				_ = _1639_v62
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_43
				var _nw293 _dafny.Array
				_ = _nw293
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw293 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Set = (func(_1640_v2 *C4) func(_dafny.Int) _dafny.Set {
						return func(_1641_i8 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_dafny.IntOfInt64(146), (_1640_v2).F34())
						}
					})(_1545_v2)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw293 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw293).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw293).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1639_v62 = _nw293
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1639_v62), 0))
				_ = _index261
				(_1639_v62).ArraySet1(_1610_v45, (_index261).Int())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1639_v62), 0))
				_ = _index262
				(_1639_v62).ArraySet1((func() _dafny.Set {
					if (_1545_v2).F33() {
						return _1610_v45
					}
					return (_dafny.SetOf(_dafny.IntOfUint32((_this.F30()).Cardinality()), (_1545_v2).F34())).Intersection(_1610_v45)
				})(), (_index262).Int())
				var _1642_v63 T0
				_ = _1642_v63
				var _nw294 *C1 = New_C1_()
				_ = _nw294
				_nw294.Ctor__(_1606_cf61)
				_1642_v63 = _nw294
				var _1643_v64 _dafny.Sequence
				_ = _1643_v64
				_1643_v64 = _dafny.SeqOf(_1642_v63)
				var _1644_v65 _dafny.Array
				_ = _1644_v65
				var _nwElement0_57 T0 = _1642_v63
				_ = _nwElement0_57
				var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(11))
				_ = _nw295
				(_nw295).ArraySet1(_nwElement0_57, 0)
				(_nw295).ArraySet1(_1642_v63, 1)
				(_nw295).ArraySet1((_1643_v64).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1643_v64).Cardinality()))).Uint32()).(T0), 2)
				(_nw295).ArraySet1(_1642_v63, 3)
				(_nw295).ArraySet1(_1642_v63, 4)
				(_nw295).ArraySet1(_1642_v63, 5)
				(_nw295).ArraySet1(_1642_v63, 6)
				(_nw295).ArraySet1((func() T0 {
					if (_this).F29() {
						return _1642_v63
					}
					return _1642_v63
				})(), 7)
				(_nw295).ArraySet1(_1642_v63, 8)
				(_nw295).ArraySet1(_1642_v63, 9)
				(_nw295).ArraySet1(_1642_v63, 10)
				_1644_v65 = _nw295
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1644_v65), 0))
				_ = _index263
				(_1644_v65).ArraySet1(_1642_v63, (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index264
				(_1542_v0).ArraySet1(Companion_Default___.Fm19(_1547_v4, globalState), (_index264).Int())
				var _1645_v66 _dafny.Array
				_ = _1645_v66
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_44
				var _nw296 _dafny.Array
				_ = _nw296
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw296 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) _dafny.Sequence = func(_1646_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("agpxc")
					}
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw296 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw296).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw296).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1645_v66 = _nw296
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1645_v66), 0))
				_ = _index265
				(_1645_v66).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg107 _dafny.Int) interface{} {
						return coer107(arg107)
					}
				}((func(_1647_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1648_i10 _dafny.Int) _dafny.CodePoint {
						return _1647_v4
					}
				})(_1547_v4))), (_index265).Int())
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1644_v65), 0))
				_ = _index266
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index267
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1645_v66), 0))
				_ = _index268
				var _rhs307 bool = (_1545_v2).F33()
				_ = _rhs307
				var _rhs308 T0 = _1642_v63
				_ = _rhs308
				var _rhs309 bool = (_1545_v2).F33()
				_ = _rhs309
				var _rhs310 _dafny.Int = _dafny.IntOfInt64(-943)
				_ = _rhs310
				var _rhs311 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("wrfecmi")
				_ = _rhs311
				var _lhs287 *GlobalState = globalState
				_ = _lhs287
				var _lhs288 _dafny.Array = _1644_v65
				_ = _lhs288
				var _lhs289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1644_v65), 0))
				_ = _lhs289
				var _lhs290 _dafny.Array = _1542_v0
				_ = _lhs290
				var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1542_v0), 0))
				_ = _lhs291
				var _lhs292 *GlobalState = globalState
				_ = _lhs292
				var _lhs293 _dafny.Array = _1645_v66
				_ = _lhs293
				var _lhs294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1645_v66), 0))
				_ = _lhs294
				_lhs287.F18 = _rhs307
				(_lhs288).ArraySet1(_rhs308, (_lhs289).Int())
				(_lhs290).ArraySet1(_rhs309, (_lhs291).Int())
				_lhs292.F21 = _rhs310
				(_lhs293).ArraySet1(_rhs311, (_lhs294).Int())
				var _1649_v67 _dafny.Map
				_ = _1649_v67
				_1649_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), (_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool))
				var _1650_v68 _dafny.Sequence
				_ = _1650_v68
				_1650_v68 = _dafny.SeqOf(p0, _dafny.IntOfUint32((Companion_Default___.Fm7((_dafny.Zero).Minus((_1545_v2).F34()), globalState)).Cardinality()), (_1545_v2).F34(), (_1545_v2).F34(), (_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int))
				var _rhs312 _dafny.Sequence = (func() _dafny.Sequence {
					if (_this).F32() {
						return _1650_v68
					}
					return _dafny.SeqOf((_1545_v2).F34())
				})()
				_ = _rhs312
				var _rhs313 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1650_v68, _1650_v68)
				_ = _rhs313
				var _rhs314 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), (_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool))
				_ = _rhs314
				var _lhs295 *GlobalState = globalState
				_ = _lhs295
				var _lhs296 *GlobalState = globalState
				_ = _lhs296
				_lhs295.F7 = _rhs312
				_lhs296.F7 = _rhs313
				_1649_v67 = _rhs314
				var _1651_v69 *C0
				_ = _1651_v69
				var _nw297 *C0 = New_C0_()
				_ = _nw297
				_nw297.Ctor__(((_1639_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1639_v62), 0))).Int()).(_dafny.Set)).IsDisjointFrom((_1639_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1639_v62), 0))).Int()).(_dafny.Set)))
				_1651_v69 = _nw297
				var _nw298 *C0 = New_C0_()
				_ = _nw298
				_nw298.Ctor__((_1554_v11).Equals(_1554_v11))
				_1651_v69 = _nw298
			}
			(globalState).F28 = (_1551_v8).Contains(_dafny.IntOfInt64(-874))
		} else if _source22.Is_DC37() {
			var _1652___mcc_h5 _dafny.Int = _source22.Get_().(D11_DC37).Cf62
			_ = _1652___mcc_h5
			var _1653___mcc_h6 bool = _source22.Get_().(D11_DC37).Cf63
			_ = _1653___mcc_h6
			var _1654_cf63 bool = _1653___mcc_h6
			_ = _1654_cf63
			var _1655_cf62 _dafny.Int = _1652___mcc_h5
			_ = _1655_cf62
			if (func() bool {
				if (_dafny.IntOfInt64(-359)).Cmp(p0) <= 0 {
					return (func() bool {
						if (_1545_v2).F33() {
							return _1654_cf63
						}
						return false
					})()
				}
				return ((_1545_v2).F34()).Cmp(p1) >= 0
			})() {
				var _1656_v70 _dafny.Array
				_ = _1656_v70
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_45
				var _nw299 _dafny.Array
				_ = _nw299
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw299 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.Sequence = (func(_1657_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1658_i11 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(550))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg108 _dafny.Int) interface{} {
									return coer108(arg108)
								}
							}((func(_1659_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1660_i12 _dafny.Int) _dafny.CodePoint {
									return _1659_v4
								}
							})(_1657_v4)))
						}
					})(_1547_v4)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw299 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw299).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw299).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1656_v70 = _nw299
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1656_v70), 0))
				_ = _index269
				(_1656_v70).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-415))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}((func(_1661_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1662_i13 _dafny.Int) _dafny.CodePoint {
						return _1661_v4
					}
				})(_1547_v4))), (_index269).Int())
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1656_v70), 0))
				_ = _index270
				(_1656_v70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}(func(_1663_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))), (_index270).Int())
				var _1664_v71 _dafny.Map
				_ = _1664_v71
				_1664_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), _1547_v4)
				var _1665_v72 _dafny.Sequence
				_ = _1665_v72
				_1665_v72 = _dafny.SeqOf(_1664_v71)
				var _1666_v73 D2
				_ = _1666_v73
				_1666_v73 = Companion_D2_.Create_DC8_(_1547_v4, (_this).F31())
				(globalState).F18 = (func() bool {
					if (_1551_v8).Contains(_1655_cf62) {
						return (_1551_v8).Get(_1655_cf62).(bool)
					}
					return ((((_1665_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1665_v72).Cardinality()))).Uint32()).(_dafny.Map)).Update(_1654_cf63, (_1666_v73).Dtor_cf12())).Cardinality()).Cmp(_dafny.IntOfInt64(-222)) > 0
				})()
				var _rhs315 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F30(), (_1656_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1656_v70), 0))).Int()).(_dafny.Sequence)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(795))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}((func(_1667_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1668_i15 _dafny.Int) _dafny.CodePoint {
						return _1667_v4
					}
				})(_1547_v4))))
				_ = _rhs315
				var _rhs316 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), (_dafny.Zero).Minus((_1545_v2).F34()))).Cardinality(), (_1545_v2).F34())).Cardinality(), (_1545_v2).F34())
				_ = _rhs316
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				_lhs297.F20 = _rhs315
				_lhs298.F1 = _rhs316
				var _1669_v74 _dafny.Sequence
				_ = _1669_v74
				_1669_v74 = _dafny.SeqOf((_1656_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1656_v70), 0))).Int()).(_dafny.Sequence), _this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("pbew"), (_1656_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1656_v70), 0))).Int()).(_dafny.Sequence), _this.F30())
				_1551_v8 = (_1551_v8).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1669_v74, _1669_v74)).Cardinality()), (_this).F32())
				var _1670_v75 _dafny.Sequence
				_ = _1670_v75
				_1670_v75 = _dafny.SeqOf(_1654_cf63, (_this).F31())
				(globalState).F28 = _dafny.Companion_Sequence_.Equal(_1670_v75, _dafny.Companion_Sequence_.Concatenate(_1670_v75, _dafny.SeqOf((_this).F32())))
			} else {
				_1655_cf62 = (func() _dafny.Int {
					if (_1554_v11).Contains(_1655_cf62) {
						return (_1554_v11).Multiplicity(_1655_cf62)
					}
					return p0
				})()
				var _1671_v76 *C5
				_ = _1671_v76
				var _nw300 *C5 = New_C5_()
				_ = _nw300
				_nw300.Ctor__(_1654_cf63, _this.F30())
				_1671_v76 = _nw300
				(globalState).F2 = p1
				r2 = p1
				var _1672_v77 _dafny.Map
				_ = _1672_v77
				_1672_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1654_cf63, _this.F30())
				_1655_cf62 = (((Companion_Default___.Fm52(globalState)).Update((_this).F29(), _this.F30())).Merge(_1672_v77)).Cardinality()
			}
			_1547_v4 = _dafny.CodePoint('c')
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1542_v0), 0))
			_ = _index271
			(_1542_v0).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("ejv")), (_index271).Int())
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1542_v0), 0))
			_ = _index272
			(_1542_v0).ArraySet1(Companion_Default___.Fm19(_1547_v4, globalState), (_index272).Int())
			var _1673_v78 D11
			_ = _1673_v78
			_1673_v78 = Companion_D11_.Create_DC35_(_dafny.UnicodeSeqOfUtf8Bytes("qqnwtssk"), (_1545_v2).F33())
			var _1674_v79 _dafny.MultiSet
			_ = _1674_v79
			_1674_v79 = _dafny.MultiSetOf(!((_1673_v78).Dtor_cf58()))
			var _1675_v80 _dafny.MultiSet
			_ = _1675_v80
			_1675_v80 = _dafny.MultiSetOf(_1547_v4, _1547_v4)
			var _1676_v81 _dafny.Sequence
			_ = _1676_v81
			_1676_v81 = _dafny.SeqOf(_1654_cf63)
			var _1677_v82 _dafny.Sequence
			_ = _1677_v82
			_1677_v82 = _dafny.SeqOf(_1676_v81, _1676_v81, _1676_v81, _1676_v81)
			_1674_v79 = (_dafny.MultiSetOf((_this).F31(), (_1545_v2).F33())).Union(((_1674_v79).Update(_1654_cf63, Companion_Default___.Abs((_1675_v80).Cardinality()))).Union(_dafny.MultiSetFromSeq((_1677_v82).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1677_v82).Cardinality()))).Uint32()).(_dafny.Sequence))))
		} else {
			var _1678___mcc_h7 _dafny.Map = _source22.Get_().(D11_DC34).Cf56
			_ = _1678___mcc_h7
			var _1679_cf56 _dafny.Map = _1678___mcc_h7
			_ = _1679_cf56
			var _1680_v83 _dafny.Sequence
			_ = _1680_v83
			_1680_v83 = _dafny.SeqOf((_1545_v2).F34())
			var _1681_v84 D9
			_ = _1681_v84
			_1681_v84 = Companion_D9_.Create_DC27_(_1680_v83)
			var _1682_v85 D2
			_ = _1682_v85
			_1682_v85 = Companion_D2_.Create_DC9_(_this.F30(), _dafny.IntOfUint32(((_1681_v84).Dtor_cf40()).Cardinality()), (_1545_v2).F33(), (_1545_v2).F34())
			var _source24 D2 = _1682_v85
			_ = _source24
			if _source24.Is_DC7() {
				var _1683___mcc_h12 bool = _source24.Get_().(D2_DC7).Cf11
				_ = _1683___mcc_h12
				var _1684_cf11 bool = _1683___mcc_h12
				_ = _1684_cf11
				var _1685_v86 _dafny.Map
				_ = _1685_v86
				_1685_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_1545_v2).F34(), (_1545_v2).F34()), _this.F30())
				_1685_v86 = (_1685_v86).Update(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(663))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}(func(_1686_i16 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				})))
				(globalState).F13 = _1684_cf11
				(globalState).F1 = (_dafny.Zero).Minus(p1)
				var _1687_v87 _dafny.Sequence
				_ = _1687_v87
				_1687_v87 = _dafny.SeqOf((_this).F32(), (_this).F32(), (p1).Cmp((_1545_v2).F34()) >= 0)
				(globalState).F13 = (_1687_v87).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())).Cardinality()), _dafny.IntOfUint32((_1687_v87).Cardinality()))).Uint32()).(bool)
			} else if _source24.Is_DC8() {
				var _1688___mcc_h13 _dafny.CodePoint = _source24.Get_().(D2_DC8).Cf12
				_ = _1688___mcc_h13
				var _1689___mcc_h14 bool = _source24.Get_().(D2_DC8).Cf13
				_ = _1689___mcc_h14
				var _1690_cf13 bool = _1689___mcc_h14
				_ = _1690_cf13
				var _1691_cf12 _dafny.CodePoint = _1688___mcc_h13
				_ = _1691_cf12
				var _1692_v88 _dafny.MultiSet
				_ = _1692_v88
				_1692_v88 = _dafny.MultiSetOf(_1690_cf13, (_this).F31())
				var _1693_v89 _dafny.Sequence
				_ = _1693_v89
				_1693_v89 = _dafny.SeqOf((_this).F32(), (_1545_v2).F33(), false)
				_1692_v88 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1693_v89, _1693_v89))
				(globalState).F2 = (_1545_v2).F34()
				var _rhs317 bool = (func() bool {
					if _1690_cf13 {
						return (p0).Cmp(p1) >= 0
					}
					return (_this).F32()
				})()
				_ = _rhs317
				var _rhs318 bool = _1690_cf13
				_ = _rhs318
				var _rhs319 bool = (_1545_v2).F33()
				_ = _rhs319
				var _rhs320 _dafny.Int = (_1545_v2).F34()
				_ = _rhs320
				var _lhs299 *GlobalState = globalState
				_ = _lhs299
				var _lhs300 *GlobalState = globalState
				_ = _lhs300
				var _lhs301 *GlobalState = globalState
				_ = _lhs301
				var _lhs302 *GlobalState = globalState
				_ = _lhs302
				_lhs299.F18 = _rhs317
				_lhs300.F18 = _rhs318
				_lhs301.F13 = _rhs319
				_lhs302.F2 = _rhs320
				var _1694_v90 _dafny.Set
				_ = _1694_v90
				_1694_v90 = _dafny.SetOf((_1545_v2).F33())
				var _1695_v91 D6
				_ = _1695_v91
				_1695_v91 = Companion_D6_.Create_DC21_(_1694_v90)
				var _1696_v92 _dafny.Sequence
				_ = _1696_v92
				_1696_v92 = _dafny.SeqOf(_1695_v91)
				var _1697_v93 D8
				_ = _1697_v93
				_1697_v93 = Companion_D8_.Create_DC25_(_1696_v92)
				var _1698_v94 D8
				_ = _1698_v94
				_1698_v94 = Companion_D8_.Create_DC25_((_1697_v93).Dtor_cf39())
				var _1699_v95 _dafny.Array
				_ = _1699_v95
				var _nwElement0_58 D8 = _1697_v93
				_ = _nwElement0_58
				var _nw301 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(20))
				_ = _nw301
				(_nw301).ArraySet1(_nwElement0_58, 0)
				(_nw301).ArraySet1(_1698_v94, 1)
				(_nw301).ArraySet1(_1698_v94, 2)
				(_nw301).ArraySet1(_1697_v93, 3)
				(_nw301).ArraySet1(_1698_v94, 4)
				(_nw301).ArraySet1(_1698_v94, 5)
				(_nw301).ArraySet1(Companion_D8_.Create_DC25_(_1696_v92), 6)
				(_nw301).ArraySet1(_1698_v94, 7)
				(_nw301).ArraySet1(Companion_D8_.Create_DC25_(_dafny.SeqOf(_1695_v91, Companion_D6_.Create_DC21_(_1694_v90))), 8)
				(_nw301).ArraySet1(_1697_v93, 9)
				(_nw301).ArraySet1(_1698_v94, 10)
				(_nw301).ArraySet1(_1697_v93, 11)
				(_nw301).ArraySet1(_1697_v93, 12)
				(_nw301).ArraySet1(_1697_v93, 13)
				(_nw301).ArraySet1(_1697_v93, 14)
				(_nw301).ArraySet1(_1698_v94, 15)
				(_nw301).ArraySet1(_1698_v94, 16)
				(_nw301).ArraySet1(_1698_v94, 17)
				(_nw301).ArraySet1(_1698_v94, 18)
				(_nw301).ArraySet1(_1697_v93, 19)
				_1699_v95 = _nw301
				var _1700_v96 _dafny.Map
				_ = _1700_v96
				_1700_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1699_v95, (_this).F29())
				var _rhs321 _dafny.Int = (_1700_v96).Cardinality()
				_ = _rhs321
				var _rhs322 _dafny.Int = p0
				_ = _rhs322
				var _lhs303 *GlobalState = globalState
				_ = _lhs303
				var _lhs304 *GlobalState = globalState
				_ = _lhs304
				_lhs303.F19 = _rhs321
				_lhs304.F21 = _rhs322
			} else if _source24.Is_DC9() {
				var _1701___mcc_h15 _dafny.Sequence = _source24.Get_().(D2_DC9).Cf14
				_ = _1701___mcc_h15
				var _1702___mcc_h16 _dafny.Int = _source24.Get_().(D2_DC9).Cf15
				_ = _1702___mcc_h16
				var _1703___mcc_h17 bool = _source24.Get_().(D2_DC9).Cf16
				_ = _1703___mcc_h17
				var _1704___mcc_h18 _dafny.Int = _source24.Get_().(D2_DC9).Cf17
				_ = _1704___mcc_h18
				var _1705_cf17 _dafny.Int = _1704___mcc_h18
				_ = _1705_cf17
				var _1706_cf16 bool = _1703___mcc_h17
				_ = _1706_cf16
				var _1707_cf15 _dafny.Int = _1702___mcc_h16
				_ = _1707_cf15
				var _1708_cf14 _dafny.Sequence = _1701___mcc_h15
				_ = _1708_cf14
				var _1709_v97 *C5
				_ = _1709_v97
				var _nw302 *C5 = New_C5_()
				_ = _nw302
				_nw302.Ctor__((_this).F31(), _this.F30())
				_1709_v97 = _nw302
				(globalState).F13 = !(((_1545_v2).F33()) == ((_this).F29()))
				var _rhs323 _dafny.Array = _1550_v7
				_ = _rhs323
				var _rhs324 _dafny.Int = _1705_cf17
				_ = _rhs324
				var _rhs325 bool = (_1545_v2).F33()
				_ = _rhs325
				var _lhs305 *GlobalState = globalState
				_ = _lhs305
				var _lhs306 *GlobalState = globalState
				_ = _lhs306
				_lhs305.F17 = _rhs323
				r2 = _rhs324
				_lhs306.F18 = _rhs325
				var _1710_v98 _dafny.Array
				_ = _1710_v98
				var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
				_ = _nw303
				_1710_v98 = _nw303
				_1710_v98 = _1710_v98
			} else {
				var _1711___mcc_h19 _dafny.MultiSet = _source24.Get_().(D2_DC6).Cf10
				_ = _1711___mcc_h19
				var _1712_cf10 _dafny.MultiSet = _1711___mcc_h19
				_ = _1712_cf10
				var _1713_v99 _dafny.Map
				_ = _1713_v99
				_1713_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1545_v2).F34()).Cmp((_1545_v2).F34()) >= 0, p0)
				var _1714_v100 _dafny.Sequence
				_ = _1714_v100
				_1714_v100 = _dafny.SeqOf(_1542_v0, _1542_v0, _1542_v0)
				var _1715_v101 D11
				_ = _1715_v101
				_1715_v101 = Companion_D11_.Create_DC34_(Companion_Default___.Fm45(globalState))
				var _1716_v102 _dafny.Sequence
				_ = _1716_v102
				_1716_v102 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), p1), _1713_v99, _1713_v99, _1713_v99)
				var _1717_v103 _dafny.Map
				_ = _1717_v103
				_1717_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F34(), p0)
				var _pat_let_tv34 = _1679_cf56
				_ = _pat_let_tv34
				var _rhs326 _dafny.Map = (_1716_v102).Select((Companion_Default___.SafeIndex(((_1717_v103).Cardinality()).Plus((_dafny.Zero).Minus((_1545_v2).F34())), _dafny.IntOfUint32((_1716_v102).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _rhs326
				var _rhs327 _dafny.Sequence = _1714_v100
				_ = _rhs327
				var _rhs328 D11 = func(_pat_let33_0 D11) D11 {
					return func(_1718_dt__update__tmp_h0 D11) D11 {
						return func(_pat_let34_0 _dafny.Map) D11 {
							return func(_1719_dt__update_hcf56_h0 _dafny.Map) D11 {
								return Companion_D11_.Create_DC34_(_1719_dt__update_hcf56_h0)
							}(_pat_let34_0)
						}(_pat_let_tv34)
					}(_pat_let33_0)
				}(_1715_v101)
				_ = _rhs328
				_1713_v99 = _rhs326
				_1714_v100 = _rhs327
				_1715_v101 = _rhs328
				var _1720_v104 *C1
				_ = _1720_v104
				var _nw304 *C1 = New_C1_()
				_ = _nw304
				_nw304.Ctor__((_1545_v2).F33())
				_1720_v104 = _nw304
				(globalState).F2 = (_1545_v2).F34()
				var _1721_v105 D3
				_ = _1721_v105
				_1721_v105 = Companion_D3_.Create_DC11_((_this).F31(), (_this).F31())
				var _1722_v106 D3
				_ = _1722_v106
				_1722_v106 = Companion_D3_.Create_DC12_(_1721_v105)
				var _1723_v107 D3
				_ = _1723_v107
				_1723_v107 = Companion_D3_.Create_DC12_((_1722_v106).Dtor_cf21())
				var _1724_v108 _dafny.Map
				_ = _1724_v108
				_1724_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1720_v104).F40(), _1723_v107)
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index273
				(_1542_v0).ArraySet1(true, (_index273).Int())
				var _1725_v109 _dafny.Array
				_ = _1725_v109
				var _nw305 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
				_ = _nw305
				_1725_v109 = _nw305
				var _1726_v110 _dafny.Sequence
				_ = _1726_v110
				_1726_v110 = _dafny.SeqOf(_dafny.MultiSetOf((_1545_v2).F34()))
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1725_v109), 0))
				_ = _index274
				(_1725_v109).ArraySet1(_1726_v110, (_index274).Int())
				var _1727_v111 _dafny.Set
				_ = _1727_v111
				_1727_v111 = _dafny.SetOf((_1545_v2).F33())
				var _1728_v112 _dafny.Array
				_ = _1728_v112
				var _nwElement0_59 _dafny.Set = _1727_v111
				_ = _nwElement0_59
				var _nw306 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(7))
				_ = _nw306
				(_nw306).ArraySet1(_nwElement0_59, 0)
				(_nw306).ArraySet1((_1727_v111).Difference(_1727_v111), 1)
				(_nw306).ArraySet1((func() _dafny.Set {
					if (_1720_v104).F40() {
						return _1727_v111
					}
					return Companion_Default___.Fm22(Companion_Default___.Fm36(globalState), Companion_Default___.Fm33((_this).F29(), (_1545_v2).F34(), p1, globalState), _1547_v4, globalState)
				})(), 2)
				(_nw306).ArraySet1((_1727_v111).Union(_1727_v111), 3)
				(_nw306).ArraySet1(_1727_v111, 4)
				(_nw306).ArraySet1(_1727_v111, 5)
				(_nw306).ArraySet1(Companion_Default___.Fm22((_1545_v2).F34(), (_this).F31(), _1547_v4, globalState), 6)
				_1728_v112 = _nw306
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_1728_v112), 0))
				_ = _index275
				(_1728_v112).ArraySet1(_1727_v111, (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1542_v0), 0))
				_ = _index276
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1725_v109), 0))
				_ = _index277
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_1728_v112), 0))
				_ = _index278
				var _rhs329 _dafny.Map = Companion_Default___.Fm53(!(Companion_Default___.Fm6((_1545_v2).F34(), globalState)), (_1545_v2).F34(), _this.F30(), p0, globalState)
				_ = _rhs329
				var _rhs330 bool = (_1545_v2).F33()
				_ = _rhs330
				var _rhs331 _dafny.Sequence = _1726_v110
				_ = _rhs331
				var _rhs332 _dafny.Int = _dafny.IntOfInt64(384)
				_ = _rhs332
				var _rhs333 _dafny.Set = Companion_Default___.Fm22((func() _dafny.Int {
					if (_this).F31() {
						return (_1727_v111).Cardinality()
					}
					return (_1545_v2).F34()
				})(), (_1545_v2).F33(), _1547_v4, globalState)
				_ = _rhs333
				var _lhs307 _dafny.Array = _1542_v0
				_ = _lhs307
				var _lhs308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1542_v0), 0))
				_ = _lhs308
				var _lhs309 _dafny.Array = _1725_v109
				_ = _lhs309
				var _lhs310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1725_v109), 0))
				_ = _lhs310
				var _lhs311 *GlobalState = globalState
				_ = _lhs311
				var _lhs312 _dafny.Array = _1728_v112
				_ = _lhs312
				var _lhs313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_1728_v112), 0))
				_ = _lhs313
				_1724_v108 = _rhs329
				(_lhs307).ArraySet1(_rhs330, (_lhs308).Int())
				(_lhs309).ArraySet1(_rhs331, (_lhs310).Int())
				_lhs311.F2 = _rhs332
				(_lhs312).ArraySet1(_rhs333, (_lhs313).Int())
			}
			var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1542_v0), 0))
			_ = _index279
			(_1542_v0).ArraySet1((_1545_v2).F33(), (_index279).Int())
			var _1729_v113 _dafny.Sequence
			_ = _1729_v113
			_1729_v113 = _dafny.SeqOf((_this).F29())
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1542_v0), 0))
			_ = _index280
			(_1542_v0).ArraySet1(((_1545_v2).F33()) && (_dafny.Companion_Sequence_.Contains(_1729_v113, !((_this).F29()))), (_index280).Int())
			(globalState).F15 = (Companion_D17_.Create_DC52_((_1542_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1542_v0), 0))).Int()).(bool), (_1545_v2).F33(), (_1545_v2).F34(), _1554_v11)).Dtor_cf89()
			r2 = (_1545_v2).Fm5(_dafny.SetOf((_this).F29()), globalState)
		}
		var _1730_v114 _dafny.Array
		_ = _1730_v114
		var _nw307 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
		_ = _nw307
		_1730_v114 = _nw307
		var _1731_v115 _dafny.Map
		_ = _1731_v115
		_1731_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _1542_v0)
		var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1730_v114), 0))
		_ = _index281
		(_1730_v114).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _1542_v0)).Merge(_1731_v115), (_index281).Int())
		var _1732_v116 _dafny.Set
		_ = _1732_v116
		_1732_v116 = _dafny.SetOf(_this.F30())
		var _1733_v117 _dafny.Sequence
		_ = _1733_v117
		_1733_v117 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("rjcvfdu"), _dafny.UnicodeSeqOfUtf8Bytes("lnab"))
		var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1730_v114), 0))
		_ = _index282
		var _rhs334 _dafny.Int = (_1545_v2).Fm5(_dafny.SetOf(false), globalState)
		_ = _rhs334
		var _rhs335 bool = (_1545_v2).F33()
		_ = _rhs335
		var _rhs336 _dafny.Map = _1731_v115
		_ = _rhs336
		var _rhs337 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F30(), (_1733_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-811), _dafny.IntOfUint32((_1733_v117).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()))
		_ = _rhs337
		var _rhs338 _dafny.Set = _1732_v116
		_ = _rhs338
		var _lhs314 *GlobalState = globalState
		_ = _lhs314
		var _lhs315 *GlobalState = globalState
		_ = _lhs315
		var _lhs316 _dafny.Array = _1730_v114
		_ = _lhs316
		var _lhs317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1730_v114), 0))
		_ = _lhs317
		var _lhs318 *C7 = _this
		_ = _lhs318
		_lhs314.F19 = _rhs334
		_lhs315.F18 = _rhs335
		(_lhs316).ArraySet1(_rhs336, (_lhs317).Int())
		_lhs318.F30_set_(_rhs337)
		_1732_v116 = _rhs338
		if (_this).F29() {
			var _1734_v118 _dafny.MultiSet
			_ = _1734_v118
			_1734_v118 = _dafny.MultiSetOf(_1542_v0)
			if !((_1734_v118).Union(_1734_v118)).Equals((_1734_v118).Update(_1542_v0, Companion_Default___.Abs((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1547_v4, (_this).F32())).Cardinality()))) {
				(globalState).F20 = _this.F30()
				var _1735_v119 _dafny.Array
				_ = _1735_v119
				var _nw308 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
				_ = _nw308
				_1735_v119 = _nw308
				var _1736_v120 _dafny.Array
				_ = _1736_v120
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_46
				var _nw309 _dafny.Array
				_ = _nw309
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw309 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Sequence = func(_1737_i17 _dafny.Int) _dafny.Sequence {
						return _this.F30()
					}
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw309 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw309).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw309).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1736_v120 = _nw309
				var _1738_v121 _dafny.Sequence
				_ = _1738_v121
				_1738_v121 = _dafny.SeqOf(p1, p0)
				var _1739_v122 _dafny.Map
				_ = _1739_v122
				_1739_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1736_v120, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(360))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_1740_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1741_i18 _dafny.Int) _dafny.CodePoint {
						return _1740_v4
					}
				})(_1547_v4))), _dafny.IntOfUint32((_1738_v121).Cardinality())))
				var _1742_v123 D15
				_ = _1742_v123
				_1742_v123 = Companion_D15_.Create_DC45_((_1545_v2).F34(), p1, (_1545_v2).F33(), (_this).F31(), _1736_v120)
				var _1743_v125 _dafny.Map
				_ = _1743_v125
				_1743_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), _dafny.IntOfUint32((_this.F30()).Cardinality()))
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1735_v119), 0))
				_ = _index283
				(_1735_v119).ArraySet1((func() _dafny.Map {
					if (_1739_v122).Contains((_1742_v123).Dtor_cf80()) {
						return (_1739_v122).Get((_1742_v123).Dtor_cf80()).(_dafny.Map)
					}
					return func() _dafny.Map {
						var _coll54 = _dafny.NewMapBuilder()
						_ = _coll54
						for _iter65 := _dafny.Iterate((_1743_v125).Keys().Elements()); ; {
							_compr_54, _ok65 := _iter65()
							if !_ok65 {
								break
							}
							var _1744_v124 _dafny.Sequence
							_1744_v124 = interface{}(_compr_54).(_dafny.Sequence)
							if (_1743_v125).Contains(_1744_v124) {
								_coll54.Add(_1744_v124, p0)
							}
						}
						return _coll54.ToMap()
					}()
				})(), (_index283).Int())
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1735_v119), 0))
				_ = _index284
				(_1735_v119).ArraySet1(_1743_v125, (_index284).Int())
				(globalState).F28 = (_1545_v2).F33()
				var _1745_v126 _dafny.Map
				_ = _1745_v126
				_1745_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), (_1545_v2).F33())).Cardinality(), _1542_v0)
				var _1746_v127 *C4
				_ = _1746_v127
				var _nw310 *C4 = New_C4_()
				_ = _nw310
				_nw310.Ctor__(Companion_Default___.Fm33((_this).F32(), (_1545_v2).F34(), p0, globalState), ((_1745_v126).Merge(_1745_v126)).Cardinality())
				_1746_v127 = _nw310
				_1542_v0 = _1542_v0
			} else {
				var _1747_v128 _dafny.Map
				_ = _1747_v128
				_1747_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29())
				var _1748_v129 _dafny.Sequence
				_ = _1748_v129
				_1748_v129 = _dafny.SeqOf((_this).F31())
				var _1749_v130 T1
				_ = _1749_v130
				var _nw311 *C5 = New_C5_()
				_ = _nw311
				_nw311.Ctor__(false, Companion_Default___.Fm7(p0, globalState))
				_1749_v130 = _nw311
				var _1750_v131 D17
				_ = _1750_v131
				_1750_v131 = Companion_D17_.Create_DC50_((func() bool {
					if (_1747_v128).Contains((_1748_v129).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1748_v129).Cardinality()))).Uint32()).(bool)) {
						return (_1747_v128).Get((_1748_v129).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1748_v129).Cardinality()))).Uint32()).(bool)).(bool)
					}
					return (_this).F31()
				})(), true, _1749_v130)
				_1750_v131 = Companion_D17_.Create_DC50_((_1545_v2).F33(), (true) && ((_1545_v2).F33()), _1749_v130)
				_1748_v129 = _1748_v129
				(globalState).F13 = true
				var _1751_v133 _dafny.Map
				_ = _1751_v133
				_1751_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F34(), p1)
				var _1752_v134 _dafny.Sequence
				_ = _1752_v134
				_1752_v134 = _dafny.SeqOf((_1545_v2).F34(), (_dafny.Zero).Minus((_1545_v2).F34()), Companion_Default___.Fm36(globalState), (func() _dafny.Int {
					if (_1751_v133).Contains(p0) {
						return (_1751_v133).Get(p0).(_dafny.Int)
					}
					return (_1545_v2).F34()
				})())
				(globalState).F15 = !(_1551_v8).Equals((func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter66 := _dafny.Iterate((_1752_v134).Elements()); ; {
						_compr_55, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _1753_v132 _dafny.Int
						_1753_v132 = interface{}(_compr_55).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1752_v134, _1753_v132) {
							_coll55.Add(Companion_Default___.SafeDivisionInt(_1753_v132, p1), (_this).F29())
						}
					}
					return _coll55.ToMap()
				}()).Merge(_1551_v8))
				var _1754_v135 _dafny.Array
				_ = _1754_v135
				var _nwElement0_60 _dafny.Array = _1550_v7
				_ = _nwElement0_60
				var _nw312 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(27))
				_ = _nw312
				(_nw312).ArraySet1(_nwElement0_60, 0)
				(_nw312).ArraySet1((_1552_v9).Dtor_cf74(), 1)
				(_nw312).ArraySet1(_1550_v7, 2)
				(_nw312).ArraySet1(_1550_v7, 3)
				(_nw312).ArraySet1(_1550_v7, 4)
				(_nw312).ArraySet1((func() _dafny.Array {
					if (_1545_v2).F33() {
						return _1550_v7
					}
					return _1550_v7
				})(), 5)
				(_nw312).ArraySet1(_1550_v7, 6)
				(_nw312).ArraySet1(_1550_v7, 7)
				(_nw312).ArraySet1(_1550_v7, 8)
				(_nw312).ArraySet1(_1550_v7, 9)
				(_nw312).ArraySet1(_1550_v7, 10)
				(_nw312).ArraySet1(_1550_v7, 11)
				(_nw312).ArraySet1(_1550_v7, 12)
				(_nw312).ArraySet1(_1550_v7, 13)
				(_nw312).ArraySet1(_1550_v7, 14)
				(_nw312).ArraySet1(_1550_v7, 15)
				(_nw312).ArraySet1(_1550_v7, 16)
				(_nw312).ArraySet1(_1550_v7, 17)
				(_nw312).ArraySet1(_1550_v7, 18)
				(_nw312).ArraySet1(_1550_v7, 19)
				(_nw312).ArraySet1(_1550_v7, 20)
				(_nw312).ArraySet1(_1550_v7, 21)
				(_nw312).ArraySet1(_1550_v7, 22)
				(_nw312).ArraySet1(_1550_v7, 23)
				(_nw312).ArraySet1(_1550_v7, 24)
				(_nw312).ArraySet1(_1550_v7, 25)
				(_nw312).ArraySet1(_1550_v7, 26)
				_1754_v135 = _nw312
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1754_v135), 0))
				_ = _index285
				(_1754_v135).ArraySet1(_1550_v7, (_index285).Int())
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1754_v135), 0))
				_ = _index286
				(_1754_v135).ArraySet1((func() _dafny.Array {
					if (_1545_v2).F33() {
						return _1550_v7
					}
					return (func() _dafny.Array {
						if (_1545_v2).F33() {
							return _1550_v7
						}
						return _1550_v7
					})()
				})(), (_index286).Int())
			}
			var _1755_v136 _dafny.Map
			_ = _1755_v136
			_1755_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1550_v7, (_1545_v2).F33())
			var _1756_v137 _dafny.Sequence
			_ = _1756_v137
			_1756_v137 = _dafny.SeqOf(_1550_v7)
			var _1757_v138 _dafny.Map
			_ = _1757_v138
			_1757_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_1545_v2).F34())
			_1755_v136 = (_1755_v136).Update((_1756_v137).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1757_v138).Contains((_this).F32()) {
					return (_1757_v138).Get((_this).F32()).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_1756_v137).Cardinality()))).Uint32()).(_dafny.Array), (true) || (!((_this).F32())))
			(globalState).F11 = _this.F30()
			var _1758_v139 _dafny.Set
			_ = _1758_v139
			_1758_v139 = _dafny.SetOf((_this).F29(), (_this).F29(), (_this).F31(), (_this).F31(), (_1545_v2).F33())
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index287
			(_1550_v7).ArraySet1((_1545_v2).Fm5(_1758_v139, globalState), (_index287).Int())
			var _1759_v140 _dafny.Sequence
			_ = _1759_v140
			_1759_v140 = _dafny.SeqOf(Companion_Default___.Fm9(globalState))
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index288
			var _rhs339 bool = (_this).F32()
			_ = _rhs339
			var _rhs340 _dafny.Int = (func() _dafny.Int {
				if (_this).F32() {
					return _dafny.IntOfUint32((_1759_v140).Cardinality())
				}
				return (_1545_v2).F34()
			})()
			_ = _rhs340
			var _lhs319 *GlobalState = globalState
			_ = _lhs319
			var _lhs320 _dafny.Array = _1550_v7
			_ = _lhs320
			var _lhs321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
			_ = _lhs321
			_lhs319.F28 = _rhs339
			(_lhs320).ArraySet1(_rhs340, (_lhs321).Int())
			if (func() bool {
				if (_this).F31() {
					return (_1548_v5).Dtor_cf53()
				}
				return (_this).F29()
			})() {
				_1551_v8 = (_1551_v8).Update(p0, (_1545_v2).F33())
				(globalState).F18 = (_1545_v2).F33()
				var _1760_v141 _dafny.Array
				_ = _1760_v141
				var _nw313 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw313
				_1760_v141 = _nw313
				var _1761_v142 D17
				_ = _1761_v142
				_1761_v142 = Companion_D17_.Create_DC51_(_dafny.IntOfUint32((_this.F30()).Cardinality()), _1760_v141)
				var _1762_v143 _dafny.Array
				_ = _1762_v143
				var _nwElement0_61 _dafny.Array = _1760_v141
				_ = _nwElement0_61
				var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(13))
				_ = _nw314
				(_nw314).ArraySet1(_nwElement0_61, 0)
				(_nw314).ArraySet1((_1761_v142).Dtor_cf87(), 1)
				(_nw314).ArraySet1(_1760_v141, 2)
				(_nw314).ArraySet1(_1760_v141, 3)
				(_nw314).ArraySet1(_1760_v141, 4)
				(_nw314).ArraySet1((_1761_v142).Dtor_cf87(), 5)
				(_nw314).ArraySet1(_1760_v141, 6)
				(_nw314).ArraySet1(_1760_v141, 7)
				(_nw314).ArraySet1(_1760_v141, 8)
				(_nw314).ArraySet1(_1760_v141, 9)
				(_nw314).ArraySet1(_1760_v141, 10)
				(_nw314).ArraySet1(_1760_v141, 11)
				(_nw314).ArraySet1(_1760_v141, 12)
				_1762_v143 = _nw314
				var _1763_v144 _dafny.Sequence
				_ = _1763_v144
				_1763_v144 = _dafny.SeqOf(_1762_v143)
				_1762_v143 = (_1763_v144).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1763_v144).Cardinality()))).Uint32()).(_dafny.Array)
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index289
				var _rhs341 _dafny.Int = (_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int)
				_ = _rhs341
				var _rhs342 bool = !((_1545_v2).F33())
				_ = _rhs342
				var _rhs343 _dafny.Int = Companion_Default___.SafeDivisionInt((_1545_v2).F34(), (p0).Minus((_1545_v2).F34()))
				_ = _rhs343
				var _lhs322 *GlobalState = globalState
				_ = _lhs322
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				var _lhs324 _dafny.Array = _1550_v7
				_ = _lhs324
				var _lhs325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
				_ = _lhs325
				_lhs322.F19 = _rhs341
				_lhs323.F18 = _rhs342
				(_lhs324).ArraySet1(_rhs343, (_lhs325).Int())
				var _1764_v145 D0
				_ = _1764_v145
				_1764_v145 = Companion_D0_.Create_DC1_(_1542_v0, (_1545_v2).F34(), (_1545_v2).F34())
				var _1765_v146 D0
				_ = _1765_v146
				_1765_v146 = Companion_D0_.Create_DC1_((_1764_v145).Dtor_cf1(), p0, _dafny.IntOfUint32((_this.F30()).Cardinality()))
				_1764_v145 = _1765_v146
			} else {
				var _1766_v147 _dafny.Sequence
				_ = _1766_v147
				_1766_v147 = _dafny.SeqOf((_1545_v2).F33(), ((_1545_v2).F33()) == ((_1545_v2).F33()))
				_1766_v147 = _1766_v147
				var _1767_v148 _dafny.Array
				_ = _1767_v148
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_47
				var _nw315 _dafny.Array
				_ = _nw315
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw315 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Sequence = (func(_1768_v140 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1769_i19 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_1768_v140, _1768_v140)
						}
					})(_1759_v140)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw315 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw315).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw315).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1767_v148 = _nw315
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_1767_v148), 0))
				_ = _index290
				(_1767_v148).ArraySet1(_1759_v140, (_index290).Int())
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
				_ = _index291
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_1767_v148), 0))
				_ = _index292
				var _rhs344 bool = (_this).F32()
				_ = _rhs344
				var _rhs345 _dafny.Int = p1
				_ = _rhs345
				var _rhs346 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1545_v2).F33() {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg114 _dafny.Int) interface{} {
								return coer114(arg114)
							}
						}((func(_1770_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1771_i20 _dafny.Int) _dafny.Int {
								return _1770_p1
							}
						})(p1))), _1759_v140)
					}
					return _dafny.Companion_Sequence_.Concatenate(_1759_v140, _1759_v140)
				})()
				_ = _rhs346
				var _rhs347 bool = (_1545_v2).F33()
				_ = _rhs347
				var _lhs326 *GlobalState = globalState
				_ = _lhs326
				var _lhs327 _dafny.Array = _1550_v7
				_ = _lhs327
				var _lhs328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))
				_ = _lhs328
				var _lhs329 _dafny.Array = _1767_v148
				_ = _lhs329
				var _lhs330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_1767_v148), 0))
				_ = _lhs330
				var _lhs331 *GlobalState = globalState
				_ = _lhs331
				_lhs326.F15 = _rhs344
				(_lhs327).ArraySet1(_rhs345, (_lhs328).Int())
				(_lhs329).ArraySet1(_rhs346, (_lhs330).Int())
				_lhs331.F18 = _rhs347
				(globalState).F2 = (_dafny.Zero).Minus((_1550_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_1550_v7), 0))).Int()).(_dafny.Int))
				var _1772_v149 _dafny.Array
				_ = _1772_v149
				var _nw316 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
				_ = _nw316
				_1772_v149 = _nw316
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1772_v149), 0))
				_ = _index293
				(_1772_v149).ArraySet1CodePoint((_1548_v5).Dtor_cf52(), (_index293).Int())
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1772_v149), 0))
				_ = _index294
				(_1772_v149).ArraySet1CodePoint(_1547_v4, (_index294).Int())
				(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30())
			}
		} else {
			_1547_v4 = _1547_v4
			(globalState).F15 = (func() bool {
				if (func() bool {
					if (_1551_v8).Contains((_1545_v2).F34()) {
						return (_1551_v8).Get((_1545_v2).F34()).(bool)
					}
					return (_1545_v2).F33()
				})() {
					return (_this).F32()
				}
				return (_1545_v2).F33()
			})()
			if (_1545_v2).F33() {
				(globalState).F2 = (_1545_v2).F34()
				(globalState).F1 = (p0).Times((_1545_v2).Fm5(Companion_Default___.Fm22(_dafny.IntOfInt64(641), (_this).F31(), _1547_v4, globalState), globalState))
				(globalState).F21 = (_1545_v2).F34()
				(globalState).F21 = ((_dafny.Zero).Minus((_1545_v2).F34())).Times((_1545_v2).F34())
				var _1773_v150 _dafny.Array
				_ = _1773_v150
				var _nw317 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw317
				_1773_v150 = _nw317
				var _1774_v151 _dafny.Sequence
				_ = _1774_v151
				_1774_v151 = _dafny.SeqOf(_dafny.IntOfInt64(970), (_1545_v2).F34())
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1773_v150), 0))
				_ = _index295
				(_1773_v150).ArraySet1(_1774_v151, (_index295).Int())
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1773_v150), 0))
				_ = _index296
				var _rhs348 _dafny.Sequence = _1774_v151
				_ = _rhs348
				var _rhs349 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1774_v151, _dafny.SeqOf(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfInt64(965), (_1774_v151).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1774_v151).Cardinality()))).Uint32()).(_dafny.Int), (_1545_v2).F34(), Companion_Default___.Fm36(globalState))), _1774_v151)
				_ = _rhs349
				var _lhs332 *GlobalState = globalState
				_ = _lhs332
				var _lhs333 _dafny.Array = _1773_v150
				_ = _lhs333
				var _lhs334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1773_v150), 0))
				_ = _lhs334
				_lhs332.F7 = _rhs348
				(_lhs333).ArraySet1(_rhs349, (_lhs334).Int())
			} else {
				var _1775_v152 _dafny.Sequence
				_ = _1775_v152
				_1775_v152 = _dafny.SeqOf((_this).F31())
				r0 = ((_1775_v152).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.IntOfUint32((_1775_v152).Cardinality()))).Uint32()).(bool)) == ((_this).F32())
				var _1776_v153 _dafny.Sequence
				_ = _1776_v153
				_1776_v153 = _dafny.SeqOf(_dafny.IntOfUint32((_this.F30()).Cardinality()))
				var _1777_v154 _dafny.Map
				_ = _1777_v154
				_1777_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1545_v2).F33()) && (Companion_Default___.Fm33((_1545_v2).F33(), _dafny.IntOfUint32((_1776_v153).Cardinality()), (_1545_v2).F34(), globalState)), (_this).F32())
				_1777_v154 = (_1777_v154).Update((_1545_v2).F33(), !(((_1545_v2).F34()).Cmp((_1545_v2).F34()) < 0))
				var _1778_v155 _dafny.Array
				_ = _1778_v155
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_48
				var _nw318 _dafny.Array
				_ = _nw318
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw318 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Map = (func(_1779_v8 _dafny.Map, _1780_v2 *C4) func(_dafny.Int) _dafny.Map {
						return func(_1781_i21 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1779_v8).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(194))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg115 _dafny.Int) interface{} {
									return coer115(arg115)
								}
							}((func(_1782_v2 *C4) func(_dafny.Int) _dafny.Int {
								return func(_1783_i22 _dafny.Int) _dafny.Int {
									return (_1782_v2).F34()
								}
							})(_1780_v2)))).Cardinality()))
						}
					})(_1551_v8, _1545_v2)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw318 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw318).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw318).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1778_v155 = _nw318
				var _1784_v156 _dafny.Map
				_ = _1784_v156
				_1784_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tlk")).Cardinality()), (_1554_v11).Cardinality())
				var _1785_v158 _dafny.Map
				_ = _1785_v158
				_1785_v158 = (_1784_v156).Merge(func() _dafny.Map {
					var _coll56 = _dafny.NewMapBuilder()
					_ = _coll56
					for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(18), _dafny.IntOfInt64(970))); ; {
						_compr_56, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1786_v157 _dafny.Int
						_1786_v157 = interface{}(_compr_56).(_dafny.Int)
						if ((_dafny.IntOfInt64(18)).Cmp(_1786_v157) <= 0) && ((_1786_v157).Cmp(_dafny.IntOfInt64(970)) < 0) {
							_coll56.Add((_1786_v157).Plus(p1), _dafny.IntOfInt64(716))
						}
					}
					return _coll56.ToMap()
				}())
				var _rhs350 _dafny.Array = _1778_v155
				_ = _rhs350
				var _rhs351 _dafny.Sequence = _this.F30()
				_ = _rhs351
				var _rhs352 bool = !((_this).F29())
				_ = _rhs352
				var _rhs353 _dafny.Map = (func() _dafny.Map {
					if (_this).F32() {
						return _1785_v158
					}
					return _1785_v158
				})()
				_ = _rhs353
				var _rhs354 _dafny.Int = (func() _dafny.Int {
					if (_this).F31() {
						return (_1545_v2).F34()
					}
					return (_1545_v2).F34()
				})()
				_ = _rhs354
				var _lhs335 *GlobalState = globalState
				_ = _lhs335
				var _lhs336 *GlobalState = globalState
				_ = _lhs336
				var _lhs337 *GlobalState = globalState
				_ = _lhs337
				_1778_v155 = _rhs350
				_lhs335.F11 = _rhs351
				_lhs336.F13 = _rhs352
				_1785_v158 = _rhs353
				_lhs337.F14 = _rhs354
				(globalState).F15 = (func() bool {
					if (_this).F31() {
						return false
					}
					return Companion_Default___.Fm6((_1545_v2).F34(), globalState)
				})()
				var _1787_v159 _dafny.Array
				_ = _1787_v159
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_49
				var _nw319 _dafny.Array
				_ = _nw319
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw319 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Sequence = func(_1788_i23 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("suupvv")
					}
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw319 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw319).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw319).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1787_v159 = _nw319
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1787_v159), 0))
				_ = _index297
				(_1787_v159).ArraySet1((_1733_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1776_v153).Cardinality()), _dafny.IntOfUint32((_1733_v117).Cardinality()))).Uint32()).(_dafny.Sequence), (_index297).Int())
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1787_v159), 0))
				_ = _index298
				var _rhs355 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("y")
				_ = _rhs355
				var _rhs356 bool = Companion_Default___.Fm6(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer116 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}((func(_1789_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1790_i24 _dafny.Int) _dafny.Int {
						return _1789_p0
					}
				})(p0)))).Cardinality()), globalState)
				_ = _rhs356
				var _rhs357 _dafny.Sequence = _this.F30()
				_ = _rhs357
				var _rhs358 bool = (_1775_v152).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1775_v152).Cardinality()))).Uint32()).(bool)
				_ = _rhs358
				var _lhs338 *GlobalState = globalState
				_ = _lhs338
				var _lhs339 *GlobalState = globalState
				_ = _lhs339
				var _lhs340 _dafny.Array = _1787_v159
				_ = _lhs340
				var _lhs341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1787_v159), 0))
				_ = _lhs341
				var _lhs342 *GlobalState = globalState
				_ = _lhs342
				_lhs338.F20 = _rhs355
				_lhs339.F28 = _rhs356
				(_lhs340).ArraySet1(_rhs357, (_lhs341).Int())
				_lhs342.F15 = _rhs358
			}
			var _1791_v160 _dafny.Sequence
			_ = _1791_v160
			_1791_v160 = _dafny.SeqOf(p1, (_1545_v2).F34())
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index299
			(_1550_v7).ArraySet1((func() _dafny.Int {
				if (_this).F29() {
					return p1
				}
				return (_1791_v160).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1791_v160).Cardinality()))).Uint32()).(_dafny.Int)
			})(), (_index299).Int())
			var _1792_v161 _dafny.Set
			_ = _1792_v161
			_1792_v161 = _dafny.SetOf((_1545_v2).F34(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg117 _dafny.Int) interface{} {
					return coer117(arg117)
				}
			}((func(_1793_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1794_i25 _dafny.Int) _dafny.CodePoint {
					return _1793_v4
				}
			})(_1547_v4)))).Cardinality())), (_1545_v2).F34())
			var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1550_v7), 0))
			_ = _index300
			(_1550_v7).ArraySet1(Companion_Default___.SafeDivisionInt((_1545_v2).F34(), (_1792_v161).Cardinality()), (_index300).Int())
			var _1795_v162 *C4
			_ = _1795_v162
			var _nw320 *C4 = New_C4_()
			_ = _nw320
			_nw320.Ctor__((_this).F29(), (_1545_v2).F34())
			_1795_v162 = _nw320
		}
		var _1796_v163 _dafny.MultiSet
		_ = _1796_v163
		_1796_v163 = _dafny.MultiSetOf((_this).F29(), false, (_this).F32())
		if ((func() _dafny.MultiSet {
			if false {
				return _1796_v163
			}
			return _1796_v163
		})()).IsProperSubsetOf((_1796_v163).Intersection(_1796_v163)) {
			r2 = (_1545_v2).F34()
			var _1797_v164 *C3
			_ = _1797_v164
			var _nw321 *C3 = New_C3_()
			_ = _nw321
			_nw321.Ctor__(!(((_1545_v2).F33()) && ((_this).F32())), p0, false, _this.F30())
			_1797_v164 = _nw321
			var _1798_v165 D2
			_ = _1798_v165
			_1798_v165 = Companion_D2_.Create_DC8_(_1547_v4, (_1545_v2).F33())
			var _1799_v166 _dafny.Array
			_ = _1799_v166
			var _nwElement0_62 _dafny.CodePoint = _1547_v4
			_ = _nwElement0_62
			var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(20))
			_ = _nw322
			(_nw322).ArraySet1CodePoint(_nwElement0_62, 0)
			(_nw322).ArraySet1CodePoint(_1547_v4, 1)
			(_nw322).ArraySet1CodePoint(Companion_Default___.Fm42((_1545_v2).F33(), !((_this).F32()), globalState), 2)
			(_nw322).ArraySet1CodePoint(_1547_v4, 3)
			(_nw322).ArraySet1CodePoint(_dafny.CodePoint('e'), 4)
			(_nw322).ArraySet1CodePoint(_1547_v4, 5)
			(_nw322).ArraySet1CodePoint(Companion_Default___.Fm42((_this).F31(), false, globalState), 6)
			(_nw322).ArraySet1CodePoint(_1547_v4, 7)
			(_nw322).ArraySet1CodePoint(_1547_v4, 8)
			(_nw322).ArraySet1CodePoint(_1547_v4, 9)
			(_nw322).ArraySet1CodePoint(_1547_v4, 10)
			(_nw322).ArraySet1CodePoint(_1547_v4, 11)
			(_nw322).ArraySet1CodePoint(_1547_v4, 12)
			(_nw322).ArraySet1CodePoint(_1547_v4, 13)
			(_nw322).ArraySet1CodePoint(_1547_v4, 14)
			(_nw322).ArraySet1CodePoint((_1798_v165).Dtor_cf12(), 15)
			(_nw322).ArraySet1CodePoint(_1547_v4, 16)
			(_nw322).ArraySet1CodePoint(_dafny.CodePoint('m'), 17)
			(_nw322).ArraySet1CodePoint(_1547_v4, 18)
			(_nw322).ArraySet1CodePoint(Companion_Default___.Fm42((_this).F31(), (_this).F32(), globalState), 19)
			_1799_v166 = _nw322
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1799_v166), 0))
			_ = _index301
			(_1799_v166).ArraySet1CodePoint(_1547_v4, (_index301).Int())
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1799_v166), 0))
			_ = _index302
			(_1799_v166).ArraySet1CodePoint(_1547_v4, (_index302).Int())
			var _1800_v167 _dafny.Sequence
			_ = _1800_v167
			_1800_v167 = _dafny.SeqOf((_this).F29())
			var _1801_v168 _dafny.Map
			_ = _1801_v168
			_1801_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1799_v166).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1799_v166), 0))).Int()), _dafny.IntOfUint32((_1800_v167).Cardinality()))
			var _1802_v169 _dafny.Map
			_ = _1802_v169
			_1802_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_this.F30(), _this.F30()), (func() _dafny.Int {
				if (_1801_v168).Contains(_1547_v4) {
					return (_1801_v168).Get(_1547_v4).(_dafny.Int)
				}
				return p1
			})())
			var _1803_v170 *C6
			_ = _1803_v170
			var _nw323 *C6 = New_C6_()
			_ = _nw323
			_nw323.Ctor__(true, _this.F30())
			_1803_v170 = _nw323
			var _1804_v171 _dafny.Sequence
			_ = _1804_v171
			_1804_v171 = _dafny.SeqOf(_1803_v170)
			var _1805_v172 _dafny.Sequence
			_ = _1805_v172
			_1805_v172 = _dafny.SeqOf((_1797_v164).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1804_v171, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F30()).Cardinality()), _dafny.IntOfUint32((_1804_v171).Cardinality()))).Uint32(), _1803_v170)).Cardinality()))
			_1802_v169 = ((_1802_v169).Merge(_1802_v169)).Merge((_1802_v169).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), _dafny.IntOfUint32((_1805_v172).Cardinality()))))
			var _1806_v173 _dafny.Map
			_ = _1806_v173
			_1806_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Times(p1), (_1545_v2).F34())
			var _pat_let_tv35 = _1547_v4
			_ = _pat_let_tv35
			var _1807_v174 _dafny.Sequence
			_ = _1807_v174
			_1807_v174 = _dafny.SeqOf(Companion_D2_.Create_DC8_((_1799_v166).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1799_v166), 0))).Int()), (_1797_v164).F35()), func(_pat_let35_0 D2) D2 {
				return func(_1808_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let36_0 _dafny.CodePoint) D2 {
						return func(_1809_dt__update_hcf12_h0 _dafny.CodePoint) D2 {
							return Companion_D2_.Create_DC8_(_1809_dt__update_hcf12_h0, (_1808_dt__update__tmp_h1).Dtor_cf13())
						}(_pat_let36_0)
					}(_pat_let_tv35)
				}(_pat_let35_0)
			}(Companion_D2_.Create_DC8_((_1799_v166).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1799_v166), 0))).Int()), true)))
			_1806_v173 = (_1806_v173).Update(Companion_Default___.SafeModuloInt((_1545_v2).F34(), _dafny.IntOfUint32((Companion_Default___.Fm35(_1807_v174, (_1797_v164).F36(), globalState)).Cardinality())), _dafny.IntOfInt64(980))
		} else {
			var _1810_v175 _dafny.Sequence
			_ = _1810_v175
			_1810_v175 = _dafny.SeqOf(p0)
			var _1811_v176 _dafny.Sequence
			_ = _1811_v176
			_1811_v176 = _dafny.SeqOf(_dafny.IntOfInt64(265), p0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1810_v175).Cardinality())))
			(globalState).F15 = !((_this).F29()) || (_dafny.Companion_Sequence_.Contains(_1811_v176, (_1545_v2).F34()))
			var _1812_v177 _dafny.Array
			_ = _1812_v177
			var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw324
			_1812_v177 = _nw324
			var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))
			_ = _index303
			(_1812_v177).ArraySet1(_1550_v7, (_index303).Int())
			var _1813_v178 D10
			_ = _1813_v178
			_1813_v178 = Companion_D10_.Create_DC33_((_this).F31(), (_1545_v2).F34())
			var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))
			_ = _index304
			var _rhs359 _dafny.Array = (func() _dafny.Array {
				if (_1545_v2).F33() {
					return _1550_v7
				}
				return _1550_v7
			})()
			_ = _rhs359
			var _rhs360 _dafny.Sequence = _this.F30()
			_ = _rhs360
			var _rhs361 bool = (_1813_v178).Dtor_cf54()
			_ = _rhs361
			var _lhs343 _dafny.Array = _1812_v177
			_ = _lhs343
			var _lhs344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))
			_ = _lhs344
			var _lhs345 *GlobalState = globalState
			_ = _lhs345
			var _lhs346 *GlobalState = globalState
			_ = _lhs346
			(_lhs343).ArraySet1(_rhs359, (_lhs344).Int())
			_lhs345.F11 = _rhs360
			_lhs346.F15 = _rhs361
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))
			_ = _arr2
			var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))
			_ = _index305
			(_arr2).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_this.F30()).Cardinality())), (_index305).Int())
			var _1814_v180 D11
			_ = _1814_v180
			_1814_v180 = Companion_D11_.Create_DC36_(Companion_D6_.Create_DC23_(true, (_1545_v2).F34()), _dafny.SetOf(false, (_this).F31()), (_1545_v2).F33())
			var _1815_v181 _dafny.Sequence
			_ = _1815_v181
			_1815_v181 = _dafny.SeqOf(_1550_v7)
			var _1816_v182 D6
			_ = _1816_v182
			_1816_v182 = Companion_D6_.Create_DC23_((_1545_v2).F33(), _dafny.IntOfUint32((_1815_v181).Cardinality()))
			var _pat_let_tv36 = _1816_v182
			_ = _pat_let_tv36
			var _1817_v183 _dafny.Map
			_ = _1817_v183
			_1817_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), func(_pat_let37_0 D11) D11 {
				return func(_1818_dt__update__tmp_h2 D11) D11 {
					return func(_pat_let38_0 D6) D11 {
						return func(_1819_dt__update_hcf59_h0 D6) D11 {
							return Companion_D11_.Create_DC36_(_1819_dt__update_hcf59_h0, (_1818_dt__update__tmp_h2).Dtor_cf60(), (_1818_dt__update__tmp_h2).Dtor_cf61())
						}(_pat_let38_0)
					}(_pat_let_tv36)
				}(_pat_let37_0)
			}(_1814_v180))
			var _1820_v184 _dafny.Set
			_ = _1820_v184
			_1820_v184 = _dafny.SetOf(_1817_v183, _1817_v183)
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))
			_ = _arr3
			var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))
			_ = _index306
			var _rhs362 _dafny.Int = (_1545_v2).F34()
			_ = _rhs362
			var _rhs363 bool = (func() _dafny.Set {
				var _coll57 = _dafny.NewBuilder()
				_ = _coll57
				for _iter68 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer118 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_1821_v2 *C4) func(_dafny.Int) _dafny.Map {
					return func(_1822_i26 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), Companion_D11_.Create_DC36_(Companion_D6_.Create_DC23_((_1821_v2).F33(), _dafny.IntOfUint32((_dafny.SeqOf((_1821_v2).F33())).Cardinality())), _dafny.SetOf(false), (_this).F29()))
					}
				})(_1545_v2)))).Elements()); ; {
					_compr_57, _ok68 := _iter68()
					if !_ok68 {
						break
					}
					var _1823_v179 _dafny.Map
					_1823_v179 = interface{}(_compr_57).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer119 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg119 _dafny.Int) interface{} {
							return coer119(arg119)
						}
					}((func(_1824_v2 *C4) func(_dafny.Int) _dafny.Map {
						return func(_1822_i26 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), Companion_D11_.Create_DC36_(Companion_D6_.Create_DC23_((_1824_v2).F33(), _dafny.IntOfUint32((_dafny.SeqOf((_1824_v2).F33())).Cardinality())), _dafny.SetOf(false), (_this).F29()))
						}
					})(_1545_v2))), _1823_v179) {
						_coll57.Add(_1823_v179)
					}
				}
				return _coll57.ToSet()
			}()).Equals(_1820_v184)
			_ = _rhs363
			var _lhs347 _dafny.Array = _dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))
			_ = _lhs347
			var _lhs348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))
			_ = _lhs348
			var _lhs349 *GlobalState = globalState
			_ = _lhs349
			(_lhs347).ArraySet1(_rhs362, (_lhs348).Int())
			_lhs349.F13 = _rhs363
			var _1825_v185 *C5
			_ = _1825_v185
			var _nw325 *C5 = New_C5_()
			_ = _nw325
			_nw325.Ctor__((_1545_v2).F33(), _this.F30())
			_1825_v185 = _nw325
			var _1826_v186 _dafny.Array
			_ = _1826_v186
			var _nw326 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
			_ = _nw326
			_1826_v186 = _nw326
			var _1827_v187 _dafny.Map
			_ = _1827_v187
			_1827_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1796_v163, p0)
			var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_1826_v186), 0))
			_ = _index307
			(_1826_v186).ArraySet1(_1827_v187, (_index307).Int())
			var _1828_v188 _dafny.Array
			_ = _1828_v188
			var _nw327 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
			_ = _nw327
			_1828_v188 = _nw327
			var _1829_v189 _dafny.Map
			_ = _1829_v189
			_1829_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F34(), (_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))).Int()).(_dafny.Int))
			var _1830_v190 _dafny.Set
			_ = _1830_v190
			_1830_v190 = _dafny.SetOf(_1829_v189)
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1828_v188), 0))
			_ = _index308
			(_1828_v188).ArraySet1(_1830_v190, (_index308).Int())
			var _1831_v191 _dafny.Sequence
			_ = _1831_v191
			_1831_v191 = _dafny.SeqOf(_1825_v185, _1825_v185)
			var _arr4 _dafny.Array = _dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))
			_ = _arr4
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))
			_ = _index309
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_1826_v186), 0))
			_ = _index310
			var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1828_v188), 0))
			_ = _index311
			var _rhs364 _dafny.Int = _dafny.IntOfInt64(493)
			_ = _rhs364
			var _rhs365 *C5 = (_1831_v191).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1831_v191).Cardinality()))).Uint32()).(*C5)
			_ = _rhs365
			var _rhs366 _dafny.Map = _1827_v187
			_ = _rhs366
			var _rhs367 _dafny.Set = _1830_v190
			_ = _rhs367
			var _lhs350 _dafny.Array = _dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))
			_ = _lhs350
			var _lhs351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_dafny.ArrayCastTo((_1812_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1812_v177), 0))).Int()))), 0))
			_ = _lhs351
			var _lhs352 _dafny.Array = _1826_v186
			_ = _lhs352
			var _lhs353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_1826_v186), 0))
			_ = _lhs353
			var _lhs354 _dafny.Array = _1828_v188
			_ = _lhs354
			var _lhs355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1828_v188), 0))
			_ = _lhs355
			(_lhs350).ArraySet1(_rhs364, (_lhs351).Int())
			_1825_v185 = _rhs365
			(_lhs352).ArraySet1(_rhs366, (_lhs353).Int())
			(_lhs354).ArraySet1(_rhs367, (_lhs355).Int())
			var _1832_v192 *C1
			_ = _1832_v192
			var _nw328 *C1 = New_C1_()
			_ = _nw328
			_nw328.Ctor__((_this).F31())
			_1832_v192 = _nw328
		}
		var _1833_i27 _dafny.Int
		_ = _1833_i27
		_1833_i27 = _dafny.Zero
		{
			for (_this).F31() {
				{
					if (_1833_i27).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1833_i27 = (_1833_i27).Plus(_dafny.One)
					(globalState).F18 = _dafny.Companion_Sequence_.Contains(_this.F30(), _1547_v4)
					(globalState).F18 = (_this).F31()
					var _1834_v193 _dafny.Sequence
					_ = _1834_v193
					_1834_v193 = _dafny.SeqOf((_1545_v2).F33())
					(globalState).F18 = !(_dafny.Companion_Sequence_.Contains(_1834_v193, (_1545_v2).F33()))
					if (_1545_v2).F33() {
						var _1835_v194 *C3
						_ = _1835_v194
						var _nw329 *C3 = New_C3_()
						_ = _nw329
						_nw329.Ctor__((_this).F31(), (_1545_v2).F34(), (_this).F31(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg120 _dafny.Int) interface{} {
								return coer120(arg120)
							}
						}((func(_1836_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1837_i28 _dafny.Int) _dafny.CodePoint {
								return _1836_v4
							}
						})(_1547_v4))))
						_1835_v194 = _nw329
						_1835_v194 = _1835_v194
						_1542_v0 = _1542_v0
						var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1542_v0), 0))
						_ = _index312
						(_1542_v0).ArraySet1((p1).Cmp((_1544_v1).Cardinality()) <= 0, (_index312).Int())
						var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1542_v0), 0))
						_ = _index313
						(_1542_v0).ArraySet1(((func() _dafny.Int {
							if !((_this).F32()) {
								return (_dafny.Zero).Minus(p1)
							}
							return (_1545_v2).F34()
						})()).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
							if (_1545_v2).F33() {
								return (_1835_v194).F36()
							}
							return _dafny.IntOfUint32((_this.F30()).Cardinality())
						})())) <= 0, (_index313).Int())
						(globalState).F1 = p1
						var _1838_v195 D11
						_ = _1838_v195
						_1838_v195 = Companion_D11_.Create_DC35_(_this.F30(), (_this).F32())
						var _1839_v196 D2
						_ = _1839_v196
						_1839_v196 = Companion_D2_.Create_DC9_((_1838_v195).Dtor_cf57(), (_1835_v194).F36(), (_1545_v2).F33(), ((_1554_v11).Update((_1545_v2).F34(), Companion_Default___.Abs((_1835_v194).F36()))).Cardinality())
						var _1840_v197 _dafny.MultiSet
						_ = _1840_v197
						_1840_v197 = _dafny.MultiSetOf(_1796_v163)
						var _1841_v198 _dafny.Sequence
						_ = _1841_v198
						_1841_v198 = _dafny.SeqOf(p1)
						_1839_v196 = Companion_Default___.Fm18((func() _dafny.Int {
							if (_1796_v163).Contains((_this).F31()) {
								return (_1796_v163).Multiplicity((_this).F31())
							}
							return (_1545_v2).F34()
						})(), (_dafny.Zero).Minus(((_1545_v2).F34()).Plus((_1840_v197).Cardinality())), (_1545_v2).F34(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1841_v198, _1841_v198)).Cardinality()), globalState)
					} else {
						var _1842_v200 _dafny.Set
						_ = _1842_v200
						_1842_v200 = _dafny.SetOf(p1)
						var _1843_v201 _dafny.Map
						_ = _1843_v201
						_1843_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), (func() _dafny.Set {
							var _coll58 = _dafny.NewBuilder()
							_ = _coll58
							for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(555), _dafny.IntOfInt64(896))); ; {
								_compr_58, _ok69 := _iter69()
								if !_ok69 {
									break
								}
								var _1844_v199 _dafny.Int
								_1844_v199 = interface{}(_compr_58).(_dafny.Int)
								if ((_dafny.IntOfInt64(555)).Cmp(_1844_v199) <= 0) && ((_1844_v199).Cmp(_dafny.IntOfInt64(896)) < 0) {
									_coll58.Add(Companion_Default___.SafeDivisionInt(_1844_v199, p0))
								}
							}
							return _coll58.ToSet()
						}()).IsDisjointFrom(_1842_v200))
						var _1845_v202 _dafny.Set
						_ = _1845_v202
						_1845_v202 = _dafny.SetOf((_1545_v2).F33())
						_1843_v201 = (_1843_v201).Update(!(_1845_v202).Contains((_this).F31()), (_this).F31())
						var _1846_v203 _dafny.Map
						_ = _1846_v203
						_1846_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30(), (_1834_v193).Select((Companion_Default___.SafeIndex((_1545_v2).F34(), _dafny.IntOfUint32((_1834_v193).Cardinality()))).Uint32()).(bool))
						var _1847_v204 _dafny.Set
						_ = _1847_v204
						_1847_v204 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1545_v2).F33(), (_1545_v2).F34()))
						var _rhs368 _dafny.Array = (func() _dafny.Array {
							if (_1847_v204).IsProperSubsetOf(_1847_v204) {
								return _1550_v7
							}
							return (func() _dafny.Array {
								if false {
									return _1550_v7
								}
								return _1550_v7
							})()
						})()
						_ = _rhs368
						var _rhs369 _dafny.Map = func() _dafny.Map {
							var _coll59 = _dafny.NewMapBuilder()
							_ = _coll59
							for _iter70 := _dafny.Iterate((_1733_v117).Elements()); ; {
								_compr_59, _ok70 := _iter70()
								if !_ok70 {
									break
								}
								var _1848_v205 _dafny.Sequence
								_1848_v205 = interface{}(_compr_59).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_1733_v117, _1848_v205) {
									_coll59.Add(_1848_v205, !((_1545_v2).F33()))
								}
							}
							return _coll59.ToMap()
						}()
						_ = _rhs369
						var _lhs356 *GlobalState = globalState
						_ = _lhs356
						_lhs356.F17 = _rhs368
						_1846_v203 = _rhs369
						var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1550_v7), 0))
						_ = _index314
						(_1550_v7).ArraySet1((_1545_v2).F34(), (_index314).Int())
						var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1550_v7), 0))
						_ = _index315
						var _rhs370 _dafny.CodePoint = _1547_v4
						_ = _rhs370
						var _rhs371 _dafny.Int = (p1).Plus((_1545_v2).F34())
						_ = _rhs371
						var _rhs372 bool = !_dafny.Companion_Sequence_.Equal(_this.F30(), _dafny.UnicodeSeqOfUtf8Bytes("jtnaivbv"))
						_ = _rhs372
						var _lhs357 _dafny.Array = _1550_v7
						_ = _lhs357
						var _lhs358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1550_v7), 0))
						_ = _lhs358
						var _lhs359 *GlobalState = globalState
						_ = _lhs359
						_1547_v4 = _rhs370
						(_lhs357).ArraySet1(_rhs371, (_lhs358).Int())
						_lhs359.F15 = _rhs372
						(globalState).F21 = _dafny.IntOfUint32((_this.F30()).Cardinality())
						(globalState).F13 = (_this).F31()
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		r0 = !((_this).F32())
		var _1849_v206 _dafny.Array
		_ = _1849_v206
		var _len0_50 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_50
		var _nw330 _dafny.Array
		_ = _nw330
		if _len0_50.Cmp(_dafny.Zero) == 0 {
			_nw330 = _dafny.NewArray(_len0_50)
		} else {
			var _init50 func(_dafny.Int) _dafny.CodePoint = (func(_1850_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1851_i29 _dafny.Int) _dafny.CodePoint {
					return _1850_v4
				}
			})(_1547_v4)
			_ = _init50
			var _element0_50 = _init50(_dafny.Zero)
			_ = _element0_50
			_nw330 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
			(_nw330).ArraySet1CodePoint(_element0_50, 0)
			var _nativeLen0_50 = (_len0_50).Int()
			_ = _nativeLen0_50
			for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
				(_nw330).ArraySet1CodePoint(_init50(_dafny.IntOf(_i0_50)), _i0_50)
			}
		}
		_1849_v206 = _nw330
		r1 = _1849_v206
		r2 = p0
		var _1852_v207 _dafny.Sequence
		_ = _1852_v207
		_1852_v207 = _dafny.SeqOf((_1545_v2).F34())
		var _1853_v208 _dafny.Map
		_ = _1853_v208
		_1853_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1852_v207, _1547_v4)
		var _1854_v209 _dafny.Sequence
		_ = _1854_v209
		_1854_v209 = _dafny.SeqOf((_1853_v208).Cardinality(), (_1554_v11).Cardinality())
		var _1855_v210 _dafny.Map
		_ = _1855_v210
		_1855_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1852_v207, (_1545_v2).F33())
		r3 = func() _dafny.Set {
			var _coll60 = _dafny.NewBuilder()
			_ = _coll60
			for _iter71 := _dafny.Iterate((_1855_v210).Keys().Elements()); ; {
				_compr_60, _ok71 := _iter71()
				if !_ok71 {
					break
				}
				var _1856_v211 _dafny.Sequence
				_1856_v211 = interface{}(_compr_60).(_dafny.Sequence)
				if (_1855_v210).Contains(_1856_v211) {
					_coll60.Add(_1856_v211)
				}
			}
			return _coll60.ToSet()
		}()
		return r0, r1, r2, r3
	}
}
func (_this *C7) F31() bool {
	{
		return _this._f31
	}
}
func (_this *C7) F32() bool {
	{
		return _this._f32
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	dummy byte
}

func New_C8_() *C8 {
	_this := C8{}

	return &_this
}

type CompanionStruct_C8_ struct {
}

var Companion_C8_ = CompanionStruct_C8_{}

func (_this *C8) Equals(other *C8) bool {
	return _this == other
}

func (_this *C8) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C8)
	return ok && _this.Equals(other)
}

func (*C8) String() string {
	return "_module.C8"
}

func Type_C8_() _dafny.TypeDescriptor {
	return type_C8_{}
}

type type_C8_ struct {
}

func (_this type_C8_) Default() interface{} {
	return (*C8)(nil)
}

func (_this type_C8_) String() string {
	return "main.C8"
}
func (_this *C8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) Ctor__() {
	{
	}
}
func (_this *C8) M0(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1857_v0 bool
		_ = _1857_v0
		_1857_v0 = true
		(globalState).F28 = _1857_v0
		var _1858_v1 _dafny.Array
		_ = _1858_v1
		var _len0_51 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_51
		var _nw331 _dafny.Array
		_ = _nw331
		if _len0_51.Cmp(_dafny.Zero) == 0 {
			_nw331 = _dafny.NewArray(_len0_51)
		} else {
			var _init51 func(_dafny.Int) bool = func(_1859_i1 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("r"), _dafny.UnicodeSeqOfUtf8Bytes("g"))
			}
			_ = _init51
			var _element0_51 = _init51(_dafny.Zero)
			_ = _element0_51
			_nw331 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
			(_nw331).ArraySet1(_element0_51, 0)
			var _nativeLen0_51 = (_len0_51).Int()
			_ = _nativeLen0_51
			for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
				(_nw331).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
			}
		}
		_1858_v1 = _nw331
		for _iter72 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1858_v1), 0))); ; {
			_guard_loop_11, _ok72 := _iter72()
			if !_ok72 {
				break
			}
			var _1860_i0 _dafny.Int
			_1860_i0 = interface{}(_guard_loop_11).(_dafny.Int)
			if (true) && (((_1860_i0).Sign() != -1) && ((_1860_i0).Cmp(_dafny.ArrayLen((_1858_v1), 0)) < 0)) {
				(_1858_v1).ArraySet1((((_dafny.SetOf(false, _1857_v0)).Union(_dafny.SetOf(_1857_v0))).Cardinality()).Cmp(p0) >= 0, (_1860_i0).Int())
			}
		}
		var _1861_v2 _dafny.Array
		_ = _1861_v2
		var _nw332 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw332
		_1861_v2 = _nw332
		var _1862_v3 _dafny.MultiSet
		_ = _1862_v3
		_1862_v3 = _dafny.MultiSetOf(_1857_v0)
		var _1863_v4 D0
		_ = _1863_v4
		_1863_v4 = Companion_D0_.Create_DC2_(_1862_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg121 _dafny.Int) interface{} {
				return coer121(arg121)
			}
		}(func(_1864_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), _1857_v0)
		var _1865_v5 _dafny.Sequence
		_ = _1865_v5
		_1865_v5 = _dafny.SeqOf(_1857_v0)
		var _1866_v6 _dafny.Sequence
		_ = _1866_v6
		_1866_v6 = _dafny.UnicodeSeqOfUtf8Bytes("xnoahtmjj")
		var _1867_v7 D0
		_ = _1867_v7
		_1867_v7 = Companion_D0_.Create_DC0_(_1857_v0)
		var _pat_let_tv37 = _1866_v6
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1866_v6
		_ = _pat_let_tv38
		var _pat_let_tv39 = _1866_v6
		_ = _pat_let_tv39
		var _pat_let_tv40 = _1865_v5
		_ = _pat_let_tv40
		var _pat_let_tv41 = p0
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1865_v5
		_ = _pat_let_tv42
		var _pat_let_tv43 = p3
		_ = _pat_let_tv43
		var _pat_let_tv44 = _1857_v0
		_ = _pat_let_tv44
		var _rhs373 bool = !(!(func(_source25 D0) bool {
			if _source25.Is_DC0() {
				var _1868___mcc_h0 bool = _source25.Get_().(D0_DC0).Cf0
				_ = _1868___mcc_h0
				var _1869_cf0 bool = _1868___mcc_h0
				_ = _1869_cf0
				return _1869_cf0
			} else if _source25.Is_DC1() {
				var _1870___mcc_h1 _dafny.Array = _source25.Get_().(D0_DC1).Cf1
				_ = _1870___mcc_h1
				var _1871___mcc_h2 _dafny.Int = _source25.Get_().(D0_DC1).Cf2
				_ = _1871___mcc_h2
				var _1872___mcc_h3 _dafny.Int = _source25.Get_().(D0_DC1).Cf3
				_ = _1872___mcc_h3
				var _1873_cf3 _dafny.Int = _1872___mcc_h3
				_ = _1873_cf3
				var _1874_cf2 _dafny.Int = _1871___mcc_h2
				_ = _1874_cf2
				var _1875_cf1 _dafny.Array = _1870___mcc_h1
				_ = _1875_cf1
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_pat_let_tv37, _pat_let_tv38)
			} else {
				var _1876___mcc_h4 _dafny.MultiSet = _source25.Get_().(D0_DC2).Cf4
				_ = _1876___mcc_h4
				var _1877___mcc_h5 _dafny.Sequence = _source25.Get_().(D0_DC2).Cf5
				_ = _1877___mcc_h5
				var _1878___mcc_h6 bool = _source25.Get_().(D0_DC2).Cf6
				_ = _1878___mcc_h6
				var _1879_cf6 bool = _1878___mcc_h6
				_ = _1879_cf6
				var _1880_cf5 _dafny.Sequence = _1877___mcc_h5
				_ = _1880_cf5
				var _1881_cf4 _dafny.MultiSet = _1876___mcc_h4
				_ = _1881_cf4
				return _1879_cf6
			}
		}(func(_pat_let39_0 D0) D0 {
			return func(_1882_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let40_0 _dafny.Sequence) D0 {
					return func(_1883_dt__update_hcf5_h0 _dafny.Sequence) D0 {
						return func(_pat_let41_0 bool) D0 {
							return func(_1884_dt__update_hcf6_h0 bool) D0 {
								return Companion_D0_.Create_DC2_((_1882_dt__update__tmp_h0).Dtor_cf4(), _1883_dt__update_hcf5_h0, _1884_dt__update_hcf6_h0)
							}(_pat_let41_0)
						}((_pat_let_tv40).Select((Companion_Default___.SafeIndex(_pat_let_tv41, _dafny.IntOfUint32((_pat_let_tv42).Cardinality()))).Uint32()).(bool))
					}(_pat_let40_0)
				}(_pat_let_tv39)
			}(_pat_let39_0)
		}(_1863_v4))))
		_ = _rhs373
		var _rhs374 _dafny.Array = _1861_v2
		_ = _rhs374
		var _rhs375 bool = func(_source26 D0) bool {
			if _source26.Is_DC0() {
				var _1885___mcc_h7 bool = _source26.Get_().(D0_DC0).Cf0
				_ = _1885___mcc_h7
				var _1886_cf0 bool = _1885___mcc_h7
				_ = _1886_cf0
				return _1886_cf0
			} else if _source26.Is_DC1() {
				var _1887___mcc_h8 _dafny.Array = _source26.Get_().(D0_DC1).Cf1
				_ = _1887___mcc_h8
				var _1888___mcc_h9 _dafny.Int = _source26.Get_().(D0_DC1).Cf2
				_ = _1888___mcc_h9
				var _1889___mcc_h10 _dafny.Int = _source26.Get_().(D0_DC1).Cf3
				_ = _1889___mcc_h10
				var _1890_cf3 _dafny.Int = _1889___mcc_h10
				_ = _1890_cf3
				var _1891_cf2 _dafny.Int = _1888___mcc_h9
				_ = _1891_cf2
				var _1892_cf1 _dafny.Array = _1887___mcc_h8
				_ = _1892_cf1
				return (_dafny.IntOfInt64(27)).Cmp(_pat_let_tv43) >= 0
			} else {
				var _1893___mcc_h11 _dafny.MultiSet = _source26.Get_().(D0_DC2).Cf4
				_ = _1893___mcc_h11
				var _1894___mcc_h12 _dafny.Sequence = _source26.Get_().(D0_DC2).Cf5
				_ = _1894___mcc_h12
				var _1895___mcc_h13 bool = _source26.Get_().(D0_DC2).Cf6
				_ = _1895___mcc_h13
				var _1896_cf6 bool = _1895___mcc_h13
				_ = _1896_cf6
				var _1897_cf5 _dafny.Sequence = _1894___mcc_h12
				_ = _1897_cf5
				var _1898_cf4 _dafny.MultiSet = _1893___mcc_h11
				_ = _1898_cf4
				return _pat_let_tv44
			}
		}(_1867_v7)
		_ = _rhs375
		var _rhs376 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1866_v6, _1866_v6)
		_ = _rhs376
		var _rhs377 bool = _1857_v0
		_ = _rhs377
		var _lhs360 *GlobalState = globalState
		_ = _lhs360
		var _lhs361 *GlobalState = globalState
		_ = _lhs361
		var _lhs362 *GlobalState = globalState
		_ = _lhs362
		_1857_v0 = _rhs373
		_1861_v2 = _rhs374
		_lhs360.F13 = _rhs375
		_lhs361.F11 = _rhs376
		_lhs362.F13 = _rhs377
		var _1899_i3 _dafny.Int
		_ = _1899_i3
		_1899_i3 = _dafny.Zero
		{
			for !(_1857_v0) {
				{
					if (_1899_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1899_i3 = (_1899_i3).Plus(_dafny.One)
					var _1900_v8 _dafny.Array
					_ = _1900_v8
					var _nw333 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(20))
					_ = _nw333
					_1900_v8 = _nw333
					var _pat_let_tv45 = p0
					_ = _pat_let_tv45
					var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1900_v8), 0))
					_ = _index316
					(_1900_v8).ArraySet1(func(_pat_let42_0 D0) D0 {
						return func(_1901_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let43_0 _dafny.Int) D0 {
								return func(_1902_dt__update_hcf3_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC1_((_1901_dt__update__tmp_h1).Dtor_cf1(), (_1901_dt__update__tmp_h1).Dtor_cf2(), _1902_dt__update_hcf3_h0)
								}(_pat_let43_0)
							}(_pat_let_tv45)
						}(_pat_let42_0)
					}(Companion_D0_.Create_DC1_(_1858_v1, p3, Companion_Default___.Fm0(p3, _1857_v0, _1857_v0, globalState))), (_index316).Int())
					var _1903_v9 D1
					_ = _1903_v9
					_1903_v9 = Companion_D1_.Create_DC3_(_1865_v5)
					var _1904_v10 D0
					_ = _1904_v10
					_1904_v10 = Companion_D0_.Create_DC1_(_1858_v1, Companion_Default___.Fm0(_dafny.IntOfUint32(((_1903_v9).Dtor_cf7()).Cardinality()), _1857_v0, _1857_v0, globalState), p0)
					var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1900_v8), 0))
					_ = _index317
					(_1900_v8).ArraySet1((func() D0 {
						if false {
							return _1904_v10
						}
						return Companion_D0_.Create_DC1_(_1858_v1, p0, p3)
					})(), (_index317).Int())
					var _1905_v11 _dafny.Sequence
					_ = _1905_v11
					_1905_v11 = _dafny.SeqOf(p3, p3)
					var _1906_v12 _dafny.Map
					_ = _1906_v12
					_1906_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1905_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1905_v11).Cardinality()))).Uint32()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("jr"))
					var _1907_v13 _dafny.Sequence
					_ = _1907_v13
					_1907_v13 = _dafny.SeqOf(p1)
					var _1908_v14 _dafny.Array
					_ = _1908_v14
					var _nwElement0_63 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1907_v13, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.IntOfUint32((_1907_v13).Cardinality()))).Uint32(), p1)
					_ = _nwElement0_63
					var _nw334 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(21))
					_ = _nw334
					(_nw334).ArraySet1(_nwElement0_63, 0)
					(_nw334).ArraySet1((func() _dafny.Sequence {
						if _1857_v0 {
							return _1907_v13
						}
						return _dafny.SeqOf(p1, p1)
					})(), 1)
					(_nw334).ArraySet1(_1907_v13, 2)
					(_nw334).ArraySet1(_1907_v13, 3)
					(_nw334).ArraySet1(_1907_v13, 4)
					(_nw334).ArraySet1(_1907_v13, 5)
					(_nw334).ArraySet1(_dafny.Companion_Sequence_.Update(_1907_v13, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1907_v13).Cardinality()))).Uint32(), p1), 6)
					(_nw334).ArraySet1(_1907_v13, 7)
					(_nw334).ArraySet1(_1907_v13, 8)
					(_nw334).ArraySet1(_dafny.SeqOf(p1, p1, p1), 9)
					(_nw334).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1907_v13, _1907_v13), 10)
					(_nw334).ArraySet1(_1907_v13, 11)
					(_nw334).ArraySet1(_1907_v13, 12)
					(_nw334).ArraySet1(_1907_v13, 13)
					(_nw334).ArraySet1(_1907_v13, 14)
					(_nw334).ArraySet1(_1907_v13, 15)
					(_nw334).ArraySet1(_dafny.SeqOf(p1, p1), 16)
					(_nw334).ArraySet1(_1907_v13, 17)
					(_nw334).ArraySet1((func() _dafny.Sequence {
						if _1857_v0 {
							return _1907_v13
						}
						return _1907_v13
					})(), 18)
					(_nw334).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1907_v13, _dafny.SeqOf(p1)), 19)
					(_nw334).ArraySet1(_1907_v13, 20)
					_1908_v14 = _nw334
					var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1908_v14), 0))
					_ = _index318
					(_1908_v14).ArraySet1(_dafny.Companion_Sequence_.Update(_1907_v13, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1907_v13).Cardinality()))).Uint32(), p1), (_index318).Int())
					var _1909_v15 _dafny.CodePoint
					_ = _1909_v15
					_1909_v15 = _dafny.CodePoint('e')
					var _1910_v16 _dafny.Map
					_ = _1910_v16
					_1910_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(p3))
					var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1908_v14), 0))
					_ = _index319
					var _rhs378 _dafny.Int = (func() _dafny.Int {
						if _1857_v0 {
							return Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p3))
						}
						return p0
					})()
					_ = _rhs378
					var _rhs379 _dafny.Map = _1906_v12
					_ = _rhs379
					var _rhs380 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1907_v13, _1907_v13)
					_ = _rhs380
					var _rhs381 _dafny.Int = p3
					_ = _rhs381
					var _rhs382 bool = (func() bool {
						if Companion_Default___.Fm1(_1909_v15, _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1910_v16).Contains(p3) {
								return (_1910_v16).Get(p3).(_dafny.Sequence)
							}
							return _1905_v11
						})()).Cardinality()), (_dafny.Zero).Minus(p3), _1857_v0, globalState) {
							return !(_1857_v0) || (_1857_v0)
						}
						return (Companion_Default___.Fm0(p0, _1857_v0, _1857_v0, globalState)).Cmp(p3) > 0
					})()
					_ = _rhs382
					var _lhs363 *GlobalState = globalState
					_ = _lhs363
					var _lhs364 _dafny.Array = _1908_v14
					_ = _lhs364
					var _lhs365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1908_v14), 0))
					_ = _lhs365
					var _lhs366 *GlobalState = globalState
					_ = _lhs366
					var _lhs367 *GlobalState = globalState
					_ = _lhs367
					_lhs363.F1 = _rhs378
					_1906_v12 = _rhs379
					(_lhs364).ArraySet1(_rhs380, (_lhs365).Int())
					_lhs366.F21 = _rhs381
					_lhs367.F13 = _rhs382
					var _1911_v17 _dafny.Array
					_ = _1911_v17
					var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
					_ = _nw335
					_1911_v17 = _nw335
					var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1911_v17), 0))
					_ = _index320
					(_1911_v17).ArraySet1(_1866_v6, (_index320).Int())
					var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1911_v17), 0))
					_ = _index321
					(_1911_v17).ArraySet1((func() _dafny.Sequence {
						if (_1906_v12).Contains((Companion_Default___.Fm0((_dafny.Zero).Minus(p3), _1857_v0, _1857_v0, globalState)).Plus(Companion_Default___.Fm0(_dafny.IntOfInt64(74), _1857_v0, _1857_v0, globalState))) {
							return (_1906_v12).Get((Companion_Default___.Fm0((_dafny.Zero).Minus(p3), _1857_v0, _1857_v0, globalState)).Plus(Companion_Default___.Fm0(_dafny.IntOfInt64(74), _1857_v0, _1857_v0, globalState))).(_dafny.Sequence)
						}
						return _1866_v6
					})(), (_index321).Int())
					var _1912_v18 _dafny.Set
					_ = _1912_v18
					_1912_v18 = _dafny.SetOf(_1857_v0)
					var _rhs383 _dafny.Array = p1
					_ = _rhs383
					var _rhs384 _dafny.Sequence = _1905_v11
					_ = _rhs384
					var _rhs385 _dafny.Int = _dafny.IntOfUint32(((_1911_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1911_v17), 0))).Int()).(_dafny.Sequence)).Cardinality())
					_ = _rhs385
					var _rhs386 _dafny.Set = _dafny.SetOf(_dafny.Companion_Sequence_.Equal(_1905_v11, _dafny.SeqOf(_dafny.IntOfInt64(-408))))
					_ = _rhs386
					var _lhs368 *GlobalState = globalState
					_ = _lhs368
					var _lhs369 *GlobalState = globalState
					_ = _lhs369
					var _lhs370 *GlobalState = globalState
					_ = _lhs370
					_lhs368.F17 = _rhs383
					_lhs369.F7 = _rhs384
					_lhs370.F21 = _rhs385
					_1912_v18 = _rhs386
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _source27 D0 = _1863_v4
		_ = _source27
		if _source27.Is_DC0() {
			var _1913___mcc_h14 bool = _source27.Get_().(D0_DC0).Cf0
			_ = _1913___mcc_h14
			var _1914_cf0 bool = _1913___mcc_h14
			_ = _1914_cf0
			var _1915_v19 *C0
			_ = _1915_v19
			var _nw336 *C0 = New_C0_()
			_ = _nw336
			_nw336.Ctor__((_1862_v3).IsDisjointFrom(_dafny.MultiSetFromSeq(_1865_v5)))
			_1915_v19 = _nw336
			var _1916_v20 _dafny.Array
			_ = _1916_v20
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_52
			var _nw337 _dafny.Array
			_ = _nw337
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw337 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) D6 = (func(_1917_v19 *C0, _1918_p3 _dafny.Int) func(_dafny.Int) D6 {
					return func(_1919_i4 _dafny.Int) D6 {
						return Companion_D6_.Create_DC23_(_1917_v19.F37, (_dafny.Zero).Minus(_1918_p3))
					}
				})(_1915_v19, p3)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw337 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw337).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw337).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1916_v20 = _nw337
			_1916_v20 = _1916_v20
			(globalState).F20 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg122 _dafny.Int) interface{} {
					return coer122(arg122)
				}
			}(func(_1920_i5 _dafny.Int) _dafny.CodePoint {
				return (func() _dafny.CodePoint {
					if false {
						return _dafny.CodePoint('s')
					}
					return _dafny.CodePoint('g')
				})()
			}))
			var _1921_v21 _dafny.CodePoint
			_ = _1921_v21
			_1921_v21 = _dafny.CodePoint('r')
			_1921_v21 = _1921_v21
		} else if _source27.Is_DC1() {
			var _1922___mcc_h15 _dafny.Array = _source27.Get_().(D0_DC1).Cf1
			_ = _1922___mcc_h15
			var _1923___mcc_h16 _dafny.Int = _source27.Get_().(D0_DC1).Cf2
			_ = _1923___mcc_h16
			var _1924___mcc_h17 _dafny.Int = _source27.Get_().(D0_DC1).Cf3
			_ = _1924___mcc_h17
			var _1925_cf3 _dafny.Int = _1924___mcc_h17
			_ = _1925_cf3
			var _1926_cf2 _dafny.Int = _1923___mcc_h16
			_ = _1926_cf2
			var _1927_cf1 _dafny.Array = _1922___mcc_h15
			_ = _1927_cf1
			var _1928_v22 D17
			_ = _1928_v22
			_1928_v22 = Companion_D17_.Create_DC51_(p0, p1)
			if (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p3), (_1928_v22).Dtor_cf86())).Cmp(p3) <= 0 {
				var _1929_v23 _dafny.CodePoint
				_ = _1929_v23
				_1929_v23 = _dafny.CodePoint('f')
				_1929_v23 = _1929_v23
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1927_cf1), 0))
				_ = _index322
				(_1927_cf1).ArraySet1(_1857_v0, (_index322).Int())
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1927_cf1), 0))
				_ = _index323
				(_1927_cf1).ArraySet1(_dafny.Companion_Sequence_.Contains(_1866_v6, _1929_v23), (_index323).Int())
				var _1930_v24 *C6
				_ = _1930_v24
				var _nw338 *C6 = New_C6_()
				_ = _nw338
				_nw338.Ctor__((_1927_cf1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1927_cf1), 0))).Int()).(bool), _1866_v6)
				_1930_v24 = _nw338
				var _1931_v25 _dafny.Map
				_ = _1931_v25
				_1931_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1857_v0, _dafny.IntOfInt64(181))
				_1931_v25 = (_1931_v25).Update(_1857_v0, _1925_cf3)
				(globalState).F13 = _1857_v0
			} else {
				(globalState).F13 = _1857_v0
				var _1932_v26 *C3
				_ = _1932_v26
				var _nw339 *C3 = New_C3_()
				_ = _nw339
				_nw339.Ctor__(_1857_v0, _1925_cf3, _dafny.Companion_Sequence_.Equal(_1866_v6, _1866_v6), _1866_v6)
				_1932_v26 = _nw339
				(globalState).F21 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1866_v6, _1866_v6)).Cardinality())
				(globalState).F18 = Companion_Default___.Fm1(Companion_Default___.Fm42(_1857_v0, (_1932_v26).F35(), globalState), (_dafny.Zero).Minus(p3), _dafny.IntOfInt64(753), (_1857_v0) || (_1857_v0), globalState)
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1858_v1), 0))
				_ = _index324
				(_1858_v1).ArraySet1(!(_dafny.MultiSetFromSeq(_1865_v5)).Contains(false), (_index324).Int())
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1858_v1), 0))
				_ = _index325
				(_1858_v1).ArraySet1(_1857_v0, (_index325).Int())
			}
			var _1933_v27 _dafny.Sequence
			_ = _1933_v27
			_1933_v27 = _dafny.SeqOf(p3, p0, _dafny.IntOfInt64(743), p3)
			var _1934_v28 _dafny.Set
			_ = _1934_v28
			_1934_v28 = _dafny.SetOf(_1933_v27)
			var _1935_v30 _dafny.Sequence
			_ = _1935_v30
			_1935_v30 = _dafny.SeqOf(_1933_v27, _1933_v27)
			var _1936_v32 _dafny.CodePoint
			_ = _1936_v32
			_1936_v32 = _dafny.CodePoint('n')
			var _1937_v33 _dafny.Sequence
			_ = _1937_v33
			_1937_v33 = _dafny.SeqOf((_1935_v30).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll61 = _dafny.NewMapBuilder()
				_ = _coll61
				for _iter73 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1867_v7, _1936_v32)).Keys().Elements()); ; {
					_compr_61, _ok73 := _iter73()
					if !_ok73 {
						break
					}
					var _1938_v31 D0
					_1938_v31 = interface{}(_compr_61).(D0)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1867_v7, _1936_v32)).Contains(_1938_v31) {
						_coll61.Add(_1938_v31, true)
					}
				}
				return _coll61.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_1935_v30).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1939_v35 _dafny.Map
			_ = _1939_v35
			_1939_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1857_v0, _1857_v0)
			var _1940_v36 _dafny.Map
			_ = _1940_v36
			_1940_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_1865_v5).Cardinality()), _1926_cf2), ((_1939_v35).Update(!(_1857_v0), _1857_v0)).Cardinality())
			var _1941_v42 _dafny.Array
			_ = _1941_v42
			var _nwElement0_64 _dafny.Set = (_1934_v28).Difference(_1934_v28)
			_ = _nwElement0_64
			var _nw340 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(24))
			_ = _nw340
			(_nw340).ArraySet1(_nwElement0_64, 0)
			(_nw340).ArraySet1(_1934_v28, 1)
			(_nw340).ArraySet1(_1934_v28, 2)
			(_nw340).ArraySet1(_1934_v28, 3)
			(_nw340).ArraySet1((Companion_Default___.Fm25(false, _1857_v0, _1866_v6, globalState)).Union(func() _dafny.Set {
				var _coll62 = _dafny.NewBuilder()
				_ = _coll62
				for _iter74 := _dafny.Iterate((_dafny.MultiSetOf(_1933_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(833))).Uint32(), func(coer123 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_1942_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1943_i6 _dafny.Int) _dafny.Int {
						return _1942_p3
					}
				})(p3))))).Elements()); ; {
					_compr_62, _ok74 := _iter74()
					if !_ok74 {
						break
					}
					var _1944_v29 _dafny.Sequence
					_1944_v29 = interface{}(_compr_62).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_1933_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(833))).Uint32(), func(coer124 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg124 _dafny.Int) interface{} {
							return coer124(arg124)
						}
					}((func(_1945_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1943_i6 _dafny.Int) _dafny.Int {
							return _1945_p3
						}
					})(p3))))).Contains(_1944_v29) {
						_coll62.Add(_1944_v29)
					}
				}
				return _coll62.ToSet()
			}()), 4)
			(_nw340).ArraySet1(Companion_Default___.Fm25(_1857_v0, _1857_v0, _dafny.UnicodeSeqOfUtf8Bytes("t"), globalState), 5)
			(_nw340).ArraySet1(func() _dafny.Set {
				var _coll63 = _dafny.NewBuilder()
				_ = _coll63
				for _iter75 := _dafny.Iterate((_1935_v30).Elements()); ; {
					_compr_63, _ok75 := _iter75()
					if !_ok75 {
						break
					}
					var _1946_v34 _dafny.Sequence
					_1946_v34 = interface{}(_compr_63).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_1935_v30, _1946_v34) {
						_coll63.Add(_1946_v34)
					}
				}
				return _coll63.ToSet()
			}(), 6)
			(_nw340).ArraySet1(_1934_v28, 7)
			(_nw340).ArraySet1(_1934_v28, 8)
			(_nw340).ArraySet1((_dafny.SetOf(_dafny.SeqOf((_1939_v35).Cardinality()), _1933_v27)).Intersection(func() _dafny.Set {
				var _coll64 = _dafny.NewBuilder()
				_ = _coll64
				for _iter76 := _dafny.Iterate((_1940_v36).Keys().Elements()); ; {
					_compr_64, _ok76 := _iter76()
					if !_ok76 {
						break
					}
					var _1947_v37 _dafny.Sequence
					_1947_v37 = interface{}(_compr_64).(_dafny.Sequence)
					if (_1940_v36).Contains(_1947_v37) {
						_coll64.Add(_1947_v37)
					}
				}
				return _coll64.ToSet()
			}()), 9)
			(_nw340).ArraySet1(func() _dafny.Set {
				var _coll65 = _dafny.NewBuilder()
				_ = _coll65
				for _iter77 := _dafny.Iterate((_1940_v36).Keys().Elements()); ; {
					_compr_65, _ok77 := _iter77()
					if !_ok77 {
						break
					}
					var _1948_v38 _dafny.Sequence
					_1948_v38 = interface{}(_compr_65).(_dafny.Sequence)
					if (_1940_v36).Contains(_1948_v38) {
						_coll65.Add(_1948_v38)
					}
				}
				return _coll65.ToSet()
			}(), 10)
			(_nw340).ArraySet1(_1934_v28, 11)
			(_nw340).ArraySet1((_1934_v28).Intersection(_1934_v28), 12)
			(_nw340).ArraySet1((_1934_v28).Difference(_dafny.SetOf(_1933_v27)), 13)
			(_nw340).ArraySet1((func() _dafny.Set {
				var _coll66 = _dafny.NewBuilder()
				_ = _coll66
				for _iter78 := _dafny.Iterate((_1937_v33).Elements()); ; {
					_compr_66, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _1949_v39 _dafny.Sequence
					_1949_v39 = interface{}(_compr_66).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_1937_v33, _1949_v39) {
						_coll66.Add(_1949_v39)
					}
				}
				return _coll66.ToSet()
			}()).Difference(_1934_v28), 14)
			(_nw340).ArraySet1(_1934_v28, 15)
			(_nw340).ArraySet1(_1934_v28, 16)
			(_nw340).ArraySet1((_1934_v28).Intersection(_1934_v28), 17)
			(_nw340).ArraySet1(_1934_v28, 18)
			(_nw340).ArraySet1((_1934_v28).Intersection(_1934_v28), 19)
			(_nw340).ArraySet1(func() _dafny.Set {
				var _coll67 = _dafny.NewBuilder()
				_ = _coll67
				for _iter79 := _dafny.Iterate((func() _dafny.Set {
					var _coll68 = _dafny.NewBuilder()
					_ = _coll68
					for _iter80 := _dafny.Iterate((_1935_v30).Elements()); ; {
						_compr_68, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _1950_v40 _dafny.Sequence
						_1950_v40 = interface{}(_compr_68).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_1935_v30, _1950_v40) {
							_coll68.Add(_1950_v40)
						}
					}
					return _coll68.ToSet()
				}()).Elements()); ; {
					_compr_67, _ok79 := _iter79()
					if !_ok79 {
						break
					}
					var _1951_v41 _dafny.Sequence
					_1951_v41 = interface{}(_compr_67).(_dafny.Sequence)
					if (func() _dafny.Set {
						var _coll69 = _dafny.NewBuilder()
						_ = _coll69
						for _iter81 := _dafny.Iterate((_1935_v30).Elements()); ; {
							_compr_69, _ok81 := _iter81()
							if !_ok81 {
								break
							}
							var _1952_v40 _dafny.Sequence
							_1952_v40 = interface{}(_compr_69).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_1935_v30, _1952_v40) {
								_coll69.Add(_1952_v40)
							}
						}
						return _coll69.ToSet()
					}()).Contains(_1951_v41) {
						_coll67.Add(_1951_v41)
					}
				}
				return _coll67.ToSet()
			}(), 20)
			(_nw340).ArraySet1(_dafny.SetOf(_1933_v27, _1933_v27), 21)
			(_nw340).ArraySet1(_1934_v28, 22)
			(_nw340).ArraySet1(_1934_v28, 23)
			_1941_v42 = _nw340
			var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1941_v42), 0))
			_ = _index326
			(_1941_v42).ArraySet1(_1934_v28, (_index326).Int())
			var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1941_v42), 0))
			_ = _index327
			(_1941_v42).ArraySet1(_1934_v28, (_index327).Int())
			var _1953_v43 D9
			_ = _1953_v43
			_1953_v43 = Companion_D9_.Create_DC28_(_1865_v5, _1857_v0, (_1933_v27).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_1926_cf2)).Cardinality(), _dafny.IntOfUint32((_1933_v27).Cardinality()))).Uint32()).(_dafny.Int), true, _1857_v0)
			(globalState).F19 = (_1953_v43).Dtor_cf43()
			(globalState).F1 = p0
		} else {
			var _1954___mcc_h18 _dafny.MultiSet = _source27.Get_().(D0_DC2).Cf4
			_ = _1954___mcc_h18
			var _1955___mcc_h19 _dafny.Sequence = _source27.Get_().(D0_DC2).Cf5
			_ = _1955___mcc_h19
			var _1956___mcc_h20 bool = _source27.Get_().(D0_DC2).Cf6
			_ = _1956___mcc_h20
			var _1957_cf6 bool = _1956___mcc_h20
			_ = _1957_cf6
			var _1958_cf5 _dafny.Sequence = _1955___mcc_h19
			_ = _1958_cf5
			var _1959_cf4 _dafny.MultiSet = _1954___mcc_h18
			_ = _1959_cf4
			var _1960_v44 _dafny.Map
			_ = _1960_v44
			_1960_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1857_v0, true)).Cardinality(), _1957_cf6)
			var _1961_v45 _dafny.Sequence
			_ = _1961_v45
			_1961_v45 = _dafny.SeqOf(p1)
			var _1962_v46 _dafny.Map
			_ = _1962_v46
			_1962_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1960_v44, (_1961_v45).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1961_v45).Cardinality()))).Uint32()).(_dafny.Array))
			var _1963_v47 _dafny.Map
			_ = _1963_v47
			_1963_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1957_cf6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()))
			var _1964_v48 _dafny.Map
			_ = _1964_v48
			_1964_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1963_v47).Cardinality(), p1)
			var _1965_v49 _dafny.Array
			_ = _1965_v49
			var _nwElement0_65 _dafny.Array = (func() _dafny.Array {
				if _1857_v0 {
					return p1
				}
				return p1
			})()
			_ = _nwElement0_65
			var _nw341 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(19))
			_ = _nw341
			(_nw341).ArraySet1(_nwElement0_65, 0)
			(_nw341).ArraySet1(p1, 1)
			(_nw341).ArraySet1(p1, 2)
			(_nw341).ArraySet1(p1, 3)
			(_nw341).ArraySet1(p1, 4)
			(_nw341).ArraySet1((func() _dafny.Array {
				if (_1962_v46).Contains(_1960_v44) {
					return (_1962_v46).Get(_1960_v44).(_dafny.Array)
				}
				return p1
			})(), 5)
			(_nw341).ArraySet1(p1, 6)
			(_nw341).ArraySet1(p1, 7)
			(_nw341).ArraySet1(p1, 8)
			(_nw341).ArraySet1((func() _dafny.Array {
				if (_1964_v48).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-917))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg125 _dafny.Int) interface{} {
						return coer125(arg125)
					}
				}(func(_1966_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality())) {
					return (_1964_v48).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-917))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg126 _dafny.Int) interface{} {
							return coer126(arg126)
						}
					}(func(_1967_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('g')
					}))).Cardinality())).(_dafny.Array)
				}
				return p1
			})(), 9)
			(_nw341).ArraySet1(p1, 10)
			(_nw341).ArraySet1(p1, 11)
			(_nw341).ArraySet1((func() _dafny.Array {
				if (_1964_v48).Contains(Companion_Default___.Fm0(p3, (_1865_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1960_v44).Cardinality()), _dafny.IntOfUint32((_1865_v5).Cardinality()))).Uint32()).(bool), _1957_cf6, globalState)) {
					return (_1964_v48).Get(Companion_Default___.Fm0(p3, (_1865_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1960_v44).Cardinality()), _dafny.IntOfUint32((_1865_v5).Cardinality()))).Uint32()).(bool), _1957_cf6, globalState)).(_dafny.Array)
				}
				return p1
			})(), 12)
			(_nw341).ArraySet1(p1, 13)
			(_nw341).ArraySet1(p1, 14)
			(_nw341).ArraySet1(p1, 15)
			(_nw341).ArraySet1(p1, 16)
			(_nw341).ArraySet1((func() _dafny.Array {
				if false {
					return p1
				}
				return p1
			})(), 17)
			(_nw341).ArraySet1(p1, 18)
			_1965_v49 = _nw341
			var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_1965_v49), 0))
			_ = _index328
			(_1965_v49).ArraySet1((func() _dafny.Array {
				if _1857_v0 {
					return p1
				}
				return p1
			})(), (_index328).Int())
			var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_1965_v49), 0))
			_ = _index329
			(_1965_v49).ArraySet1(p1, (_index329).Int())
			var _1968_v50 _dafny.CodePoint
			_ = _1968_v50
			_1968_v50 = _dafny.CodePoint('u')
			var _1969_v51 D2
			_ = _1969_v51
			_1969_v51 = Companion_D2_.Create_DC8_(_1968_v50, _1957_cf6)
			var _1970_v52 _dafny.Sequence
			_ = _1970_v52
			_1970_v52 = _dafny.SeqOf(_1969_v51, _1969_v51)
			_1865_v5 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm35(_1970_v52, p0, globalState), _1865_v5)
			var _1971_v53 *C6
			_ = _1971_v53
			var _nw342 *C6 = New_C6_()
			_ = _nw342
			_nw342.Ctor__(_1857_v0, _1866_v6)
			_1971_v53 = _nw342
			(globalState).F11 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1958_cf5, _1958_cf5), _1866_v6)
		}
		if _1857_v0 {
			var _pat_let_tv46 = _1866_v6
			_ = _pat_let_tv46
			var _1972_v54 _dafny.Map
			_ = _1972_v54
			_1972_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, func(_pat_let44_0 D0) D0 {
				return func(_1973_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let45_0 _dafny.Sequence) D0 {
						return func(_1974_dt__update_hcf5_h1 _dafny.Sequence) D0 {
							return Companion_D0_.Create_DC2_((_1973_dt__update__tmp_h2).Dtor_cf4(), _1974_dt__update_hcf5_h1, (_1973_dt__update__tmp_h2).Dtor_cf6())
						}(_pat_let45_0)
					}(_pat_let_tv46)
				}(_pat_let44_0)
			}(_1863_v4))
			_1972_v54 = (_1972_v54).Update(p1, _1863_v4)
			(globalState).F15 = !((!((func() bool {
				if _1857_v0 {
					return Companion_Default___.Fm10(p0, p0, globalState)
				}
				return !(false)
			})())) == (_1857_v0))
			(globalState).F13 = _1857_v0
			(globalState).F1 = ((_dafny.IntOfInt64(483)).Plus(p3)).Times(p0)
			var _1975_v55 _dafny.CodePoint
			_ = _1975_v55
			_1975_v55 = _dafny.CodePoint('j')
			var _1976_v56 D2
			_ = _1976_v56
			_1976_v56 = Companion_D2_.Create_DC8_(_1975_v55, _1857_v0)
			var _1977_v57 _dafny.Sequence
			_ = _1977_v57
			_1977_v57 = _dafny.SeqOf(_1976_v56, Companion_D2_.Create_DC8_(_1975_v55, _1857_v0))
			var _1978_v58 D18
			_ = _1978_v58
			_1978_v58 = Companion_D18_.Create_DC54_(_1977_v57)
			var _1979_v59 D1
			_ = _1979_v59
			_1979_v59 = Companion_D1_.Create_DC3_(Companion_Default___.Fm35((_1978_v58).Dtor_cf93(), p0, globalState))
			var _pat_let_tv47 = _1865_v5
			_ = _pat_let_tv47
			var _source28 D1 = func(_pat_let46_0 D1) D1 {
				return func(_1980_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let47_0 _dafny.Sequence) D1 {
						return func(_1981_dt__update_hcf7_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC3_(_1981_dt__update_hcf7_h0)
						}(_pat_let47_0)
					}(_pat_let_tv47)
				}(_pat_let46_0)
			}(_1979_v59)
			_ = _source28
			if _source28.Is_DC4() {
				var _1982___mcc_h21 bool = _source28.Get_().(D1_DC4).Cf8
				_ = _1982___mcc_h21
				var _1983_cf8 bool = _1982___mcc_h21
				_ = _1983_cf8
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((p1), 0))
				_ = _index330
				(p1).ArraySet1((p3).Plus(p0), (_index330).Int())
				var _1984_v60 _dafny.Map
				_ = _1984_v60
				_1984_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1858_v1, p3)
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((p1), 0))
				_ = _index331
				(p1).ArraySet1((func() _dafny.Int {
					if (_1984_v60).Contains(_1858_v1) {
						return (_1984_v60).Get(_1858_v1).(_dafny.Int)
					}
					return (p3).Minus(p0)
				})(), (_index331).Int())
				var _1985_v61 _dafny.MultiSet
				_ = _1985_v61
				_1985_v61 = _dafny.MultiSetOf(p0, p3)
				var _1986_v62 *C7
				_ = _1986_v62
				var _nw343 *C7 = New_C7_()
				_ = _nw343
				_nw343.Ctor__(_1983_cf8, !(_1983_cf8) || (Companion_Default___.Fm10(_dafny.IntOfUint32((_1866_v6).Cardinality()), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), globalState)), (_1985_v61).IsSubsetOf(_1985_v61), _1866_v6)
				_1986_v62 = _nw343
				var _1987_v63 *C1
				_ = _1987_v63
				var _nw344 *C1 = New_C1_()
				_ = _nw344
				_nw344.Ctor__(true)
				_1987_v63 = _nw344
				var _1988_v64 _dafny.Set
				_ = _1988_v64
				_1988_v64 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(349))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg127 _dafny.Int) interface{} {
						return coer127(arg127)
					}
				}((func(_1989_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1990_i8 _dafny.Int) _dafny.CodePoint {
						return _1989_v55
					}
				})(_1975_v55))))
				var _1991_v65 _dafny.Array
				_ = _1991_v65
				var _len0_53 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_53
				var _nw345 _dafny.Array
				_ = _nw345
				if _len0_53.Cmp(_dafny.Zero) == 0 {
					_nw345 = _dafny.NewArray(_len0_53)
				} else {
					var _init53 func(_dafny.Int) _dafny.Int = (func(_1992_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1993_i9 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1993_i9, _dafny.IntOfUint32((_1992_v6).Cardinality()))
						}
					})(_1866_v6)
					_ = _init53
					var _element0_53 = _init53(_dafny.Zero)
					_ = _element0_53
					_nw345 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
					(_nw345).ArraySet1(_element0_53, 0)
					var _nativeLen0_53 = (_len0_53).Int()
					_ = _nativeLen0_53
					for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
						(_nw345).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
					}
				}
				_1991_v65 = _nw345
				var _1994_v66 _dafny.Map
				_ = _1994_v66
				_1994_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1991_v65, Companion_Default___.Fm6((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), globalState))
				_1983_cf8 = (((_1988_v64).Cardinality()).Minus((_1994_v66).Cardinality())).Cmp(p0) == 0
			} else if _source28.Is_DC3() {
				var _1995___mcc_h22 _dafny.Sequence = _source28.Get_().(D1_DC3).Cf7
				_ = _1995___mcc_h22
				var _1996_cf7 _dafny.Sequence = _1995___mcc_h22
				_ = _1996_cf7
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_1858_v1), 0))
				_ = _index332
				(_1858_v1).ArraySet1(!(_1857_v0) || (_1857_v0), (_index332).Int())
				var _1997_v67 _dafny.MultiSet
				_ = _1997_v67
				_1997_v67 = _dafny.MultiSetOf(_1996_cf7)
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_1858_v1), 0))
				_ = _index333
				(_1858_v1).ArraySet1((_1997_v67).IsProperSubsetOf(_dafny.MultiSetOf(_1865_v5)), (_index333).Int())
				var _1998_v68 _dafny.Map
				_ = _1998_v68
				_1998_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_1975_v55, (_dafny.Zero).Minus(p0), p0, false, globalState), _1979_v59)
				var _1999_v69 D5
				_ = _1999_v69
				_1999_v69 = Companion_D5_.Create_DC20_(Companion_D5_.Create_DC16_(_1998_v68))
				var _2000_v70 _dafny.Map
				_ = _2000_v70
				_2000_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1999_v69)
				_2000_v70 = _2000_v70
				(globalState).F11 = _1866_v6
				var _2001_v71 _dafny.Set
				_ = _2001_v71
				_2001_v71 = _dafny.SetOf(_1866_v6)
				var _2002_v72 D5
				_ = _2002_v72
				_2002_v72 = Companion_D5_.Create_DC17_(_2001_v71)
				_2002_v72 = _2002_v72
			} else {
				var _2003___mcc_h23 D1 = _source28.Get_().(D1_DC5).Cf9
				_ = _2003___mcc_h23
				var _2004_cf9 D1 = _2003___mcc_h23
				_ = _2004_cf9
				(globalState).F28 = !((_1863_v4).Dtor_cf6())
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))
				_ = _index334
				(p1).ArraySet1(p0, (_index334).Int())
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))
				_ = _index335
				(p1).ArraySet1(_dafny.IntOfInt64(360), (_index335).Int())
				var _2005_v73 _dafny.MultiSet
				_ = _2005_v73
				_2005_v73 = _dafny.MultiSetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
				var _2006_v74 _dafny.Array
				_ = _2006_v74
				var _nwElement0_66 _dafny.Int = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
				_ = _nwElement0_66
				var _nw346 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(20))
				_ = _nw346
				(_nw346).ArraySet1(_nwElement0_66, 0)
				(_nw346).ArraySet1(p0, 1)
				(_nw346).ArraySet1(p3, 2)
				(_nw346).ArraySet1(p3, 3)
				(_nw346).ArraySet1((p3).Minus((_2005_v73).Cardinality()), 4)
				(_nw346).ArraySet1(p0, 5)
				(_nw346).ArraySet1(p0, 6)
				(_nw346).ArraySet1(p0, 7)
				(_nw346).ArraySet1(_dafny.IntOfInt64(-609), 8)
				(_nw346).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), 9)
				(_nw346).ArraySet1(Companion_Default___.Fm0(p0, _1857_v0, _1857_v0, globalState), 10)
				(_nw346).ArraySet1((p3).Plus(p3), 11)
				(_nw346).ArraySet1(p3, 12)
				(_nw346).ArraySet1(_dafny.IntOfUint32((_1866_v6).Cardinality()), 13)
				(_nw346).ArraySet1(_dafny.IntOfInt64(-837), 14)
				(_nw346).ArraySet1(p3, 15)
				(_nw346).ArraySet1(p3, 16)
				(_nw346).ArraySet1(Companion_Default___.Fm9(globalState), 17)
				(_nw346).ArraySet1((_2005_v73).Cardinality(), 18)
				(_nw346).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), 19)
				_2006_v74 = _nw346
				var _2007_v75 _dafny.Sequence
				_ = _2007_v75
				_2007_v75 = _dafny.SeqOf(_1866_v6)
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))
				_ = _index336
				var _rhs387 _dafny.CodePoint = _1975_v55
				_ = _rhs387
				var _rhs388 _dafny.Array = _2006_v74
				_ = _rhs388
				var _rhs389 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1866_v6, _1866_v6)).Cardinality()))
				_ = _rhs389
				var _rhs390 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_2007_v75).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.IntOfUint32((_2007_v75).Cardinality()))).Uint32()).(_dafny.Sequence), (func() _dafny.Sequence {
					if _1857_v0 {
						return _1866_v6
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer128 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg128 _dafny.Int) interface{} {
							return coer128(arg128)
						}
					}((func(_2008_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2009_i10 _dafny.Int) _dafny.CodePoint {
							return _2008_v55
						}
					})(_1975_v55)))
				})())
				_ = _rhs390
				var _lhs371 *GlobalState = globalState
				_ = _lhs371
				var _lhs372 _dafny.Array = p1
				_ = _lhs372
				var _lhs373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((p1), 0))
				_ = _lhs373
				var _lhs374 *GlobalState = globalState
				_ = _lhs374
				_1975_v55 = _rhs387
				_lhs371.F17 = _rhs388
				(_lhs372).ArraySet1(_rhs389, (_lhs373).Int())
				_lhs374.F11 = _rhs390
				var _2010_v76 *C7
				_ = _2010_v76
				var _nw347 *C7 = New_C7_()
				_ = _nw347
				_nw347.Ctor__(_1857_v0, _1857_v0, _1857_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1866_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(998))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg129 _dafny.Int) interface{} {
						return coer129(arg129)
					}
				}((func(_2011_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2012_i11 _dafny.Int) _dafny.CodePoint {
						return _2011_v55
					}
				})(_1975_v55)))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1866_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(998))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}((func(_2013_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2014_i11 _dafny.Int) _dafny.CodePoint {
						return _2013_v55
					}
				})(_1975_v55))))).Cardinality()))).Uint32(), _1975_v55))
				_2010_v76 = _nw347
			}
		} else {
			var _2015_v77 _dafny.Array
			_ = _2015_v77
			var _nw348 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
			_ = _nw348
			_2015_v77 = _nw348
			var _2016_v78 _dafny.Map
			_ = _2016_v78
			_2016_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2015_v77, _1857_v0)
			var _2017_v79 _dafny.Sequence
			_ = _2017_v79
			_2017_v79 = _dafny.SeqOf(_2015_v77)
			if !((func() bool {
				if (_2016_v78).Contains((_2017_v79).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2017_v79).Cardinality()))).Uint32()).(_dafny.Array)) {
					return (_2016_v78).Get((_2017_v79).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2017_v79).Cardinality()))).Uint32()).(_dafny.Array)).(bool)
				}
				return _1857_v0
			})()) {
				(globalState).F13 = !(_1857_v0)
				var _2018_v80 _dafny.Sequence
				_ = _2018_v80
				_2018_v80 = _dafny.SeqOf(p0, p0)
				(globalState).F19 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2018_v80, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.MultiSetOf(_1857_v0, _1857_v0)).Cardinality()), _dafny.IntOfUint32((_2018_v80).Cardinality()))).Uint32(), p0)).Cardinality())
				var _2019_v81 _dafny.Map
				_ = _2019_v81
				_2019_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
				_2019_v81 = ((_2019_v81).Merge(_2019_v81)).Merge(_2019_v81)
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_2015_v77), 0))
				_ = _index337
				(_2015_v77).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rpkac"), (_index337).Int())
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_2015_v77), 0))
				_ = _index338
				(_2015_v77).ArraySet1(_1866_v6, (_index338).Int())
				var _2020_v82 _dafny.MultiSet
				_ = _2020_v82
				_2020_v82 = _dafny.MultiSetOf(_2018_v80)
				var _rhs391 _dafny.Int = ((_2020_v82).Update(_2018_v80, Companion_Default___.Abs(p0))).Cardinality()
				_ = _rhs391
				var _rhs392 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1866_v6, _1866_v6), _dafny.Companion_Sequence_.Concatenate(_1866_v6, (_2015_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_2015_v77), 0))).Int()).(_dafny.Sequence)))
				_ = _rhs392
				var _rhs393 bool = !(_1857_v0)
				_ = _rhs393
				var _lhs375 *GlobalState = globalState
				_ = _lhs375
				var _lhs376 *GlobalState = globalState
				_ = _lhs376
				var _lhs377 *GlobalState = globalState
				_ = _lhs377
				_lhs375.F1 = _rhs391
				_lhs376.F15 = _rhs392
				_lhs377.F13 = _rhs393
			} else {
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((p1), 0))
				_ = _index339
				(p1).ArraySet1(p0, (_index339).Int())
				var _2021_v83 _dafny.MultiSet
				_ = _2021_v83
				_2021_v83 = _dafny.MultiSetOf(p3, p0)
				var _2022_v84 _dafny.MultiSet
				_ = _2022_v84
				_2022_v84 = _dafny.MultiSetOf((_2021_v83).Difference(_2021_v83))
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((p1), 0))
				_ = _index340
				(p1).ArraySet1((_2022_v84).Cardinality(), (_index340).Int())
				var _2023_v85 _dafny.Array
				_ = _2023_v85
				var _nw349 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw349
				_2023_v85 = _nw349
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2023_v85), 0))
				_ = _index341
				(_2023_v85).ArraySet1(p2, (_index341).Int())
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2023_v85), 0))
				_ = _index342
				(_2023_v85).ArraySet1(p2, (_index342).Int())
				var _2024_v86 _dafny.Array
				_ = _2024_v86
				var _nw350 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw350
				_2024_v86 = _nw350
				_2024_v86 = _2024_v86
				_1858_v1 = _1858_v1
				(globalState).F13 = !(Companion_Default___.Fm19((_1866_v6).Select((Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1866_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState))
			}
			var _2025_v87 _dafny.CodePoint
			_ = _2025_v87
			_2025_v87 = _dafny.CodePoint('d')
			_1857_v0 = Companion_Default___.Fm19(_2025_v87, globalState)
			(globalState).F13 = _1857_v0
			var _2026_v88 _dafny.Set
			_ = _2026_v88
			_2026_v88 = _dafny.SetOf(p1, p1, p1)
			var _2027_v89 *C2
			_ = _2027_v89
			var _nw351 *C2 = New_C2_()
			_ = _nw351
			_nw351.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1865_v5, _1865_v5)).Cardinality()), _2026_v88, false, _dafny.UnicodeSeqOfUtf8Bytes("wxlbbb"))
			_2027_v89 = _nw351
			var _2028_v90 _dafny.Array
			_ = _2028_v90
			var _nw352 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw352
			_2028_v90 = _nw352
			var _2029_v91 _dafny.Array
			_ = _2029_v91
			var _nw353 _dafny.Array = _dafny.NewArray(_dafny.One)
			_ = _nw353
			_2029_v91 = _nw353
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2028_v90), 0))
			_ = _index343
			(_2028_v90).ArraySet1(_2029_v91, (_index343).Int())
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2028_v90), 0))
			_ = _index344
			(_2028_v90).ArraySet1(_2029_v91, (_index344).Int())
		}
		r0 = !(_1857_v0)
		return r0
	}
}
func (_this *C8) M1(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _2030_v0 _dafny.Int
		_ = _2030_v0
		_2030_v0 = _dafny.IntOfInt64(553)
		var _2031_v1 _dafny.MultiSet
		_ = _2031_v1
		_2031_v1 = _dafny.MultiSetOf(_2030_v0)
		_2031_v1 = _2031_v1
		var _2032_v2 _dafny.Array
		_ = _2032_v2
		var _len0_54 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_54
		var _nw354 _dafny.Array
		_ = _nw354
		if _len0_54.Cmp(_dafny.Zero) == 0 {
			_nw354 = _dafny.NewArray(_len0_54)
		} else {
			var _init54 func(_dafny.Int) bool = func(_2033_i1 _dafny.Int) bool {
				return (Companion_D0_.Create_DC0_(false)).Dtor_cf0()
			}
			_ = _init54
			var _element0_54 = _init54(_dafny.Zero)
			_ = _element0_54
			_nw354 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
			(_nw354).ArraySet1(_element0_54, 0)
			var _nativeLen0_54 = (_len0_54).Int()
			_ = _nativeLen0_54
			for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
				(_nw354).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
			}
		}
		_2032_v2 = _nw354
		for _iter82 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2032_v2), 0))); ; {
			_guard_loop_12, _ok82 := _iter82()
			if !_ok82 {
				break
			}
			var _2034_i0 _dafny.Int
			_2034_i0 = interface{}(_guard_loop_12).(_dafny.Int)
			if (true) && (((_2034_i0).Sign() != -1) && ((_2034_i0).Cmp(_dafny.ArrayLen((_2032_v2), 0)) < 0)) {
				(_2032_v2).ArraySet1(((_dafny.SetOf(_dafny.IntOfInt64(852), _2030_v0)).Intersection(_dafny.SetOf(_2030_v0))).IsDisjointFrom(func() _dafny.Set {
					var _coll70 = _dafny.NewBuilder()
					_ = _coll70
					for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(833), _dafny.IntOfInt64(56))); ; {
						_compr_70, _ok83 := _iter83()
						if !_ok83 {
							break
						}
						var _2035_v3 _dafny.Int
						_2035_v3 = interface{}(_compr_70).(_dafny.Int)
						if ((_dafny.IntOfInt64(833)).Cmp(_2035_v3) <= 0) && ((_2035_v3).Cmp(_dafny.IntOfInt64(56)) < 0) {
							_coll70.Add((_2035_v3).Times((_dafny.MultiSetOf(!(true))).Cardinality()))
						}
					}
					return _coll70.ToSet()
				}()), (_2034_i0).Int())
			}
		}
		var _hi11 _dafny.Int = _dafny.IntOfInt64(518)
		_ = _hi11
		for _2036_i2 := Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(969), _2030_v0); _2036_i2.Cmp(_hi11) < 0; _2036_i2 = _2036_i2.Plus(_dafny.One) {
			(globalState).F1 = _2030_v0
			(globalState).F14 = _2030_v0
			var _2037_v4 _dafny.Sequence
			_ = _2037_v4
			_2037_v4 = _dafny.SeqOf(p0)
			_2037_v4 = _2037_v4
			(globalState).F20 = _dafny.UnicodeSeqOfUtf8Bytes("iepcaxj")
		}
		var _2038_v5 bool
		_ = _2038_v5
		_2038_v5 = false
		var _2039_v6 _dafny.Sequence
		_ = _2039_v6
		_2039_v6 = _dafny.UnicodeSeqOfUtf8Bytes("d")
		var _2040_v7 *C4
		_ = _2040_v7
		var _nw355 *C4 = New_C4_()
		_ = _nw355
		_nw355.Ctor__(!(_2038_v5), _dafny.IntOfUint32((_2039_v6).Cardinality()))
		_2040_v7 = _nw355
		var _2041_v8 _dafny.MultiSet
		_ = _2041_v8
		_2041_v8 = _dafny.MultiSetOf(_2040_v7, _2040_v7)
		_2041_v8 = (_2041_v8).Difference(_dafny.MultiSetOf(_2040_v7))
		if false {
			var _2042_v9 _dafny.Sequence
			_ = _2042_v9
			_2042_v9 = _dafny.SeqOf(_2031_v1)
			var _2043_v10 _dafny.Sequence
			_ = _2043_v10
			_2043_v10 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if true {
					return _dafny.SeqOf(_2031_v1, _2031_v1)
				}
				return _2042_v9
			})(), (Companion_Default___.SafeIndex(_2030_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if true {
					return _dafny.SeqOf(_2031_v1, _2031_v1)
				}
				return _2042_v9
			})()).Cardinality()))).Uint32(), _2031_v1))
			_2042_v9 = (_2043_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_2043_v10).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _2044_v11 _dafny.Map
			_ = _2044_v11
			_2044_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v5, (_dafny.One).Cmp((_2040_v7).F34()) <= 0)
			(globalState).F28 = (func() bool {
				if (_2044_v11).Contains((_2040_v7).F33()) {
					return (_2044_v11).Get((_2040_v7).F33()).(bool)
				}
				return _2038_v5
			})()
			var _2045_v12 _dafny.Array
			_ = _2045_v12
			var _nwElement0_67 _dafny.MultiSet = _2031_v1
			_ = _nwElement0_67
			var _nw356 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.One)
			_ = _nw356
			(_nw356).ArraySet1(_nwElement0_67, 0)
			_2045_v12 = _nw356
			_2045_v12 = _2045_v12
			var _nw357 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw357
			_2032_v2 = _nw357
			if (func() bool {
				if (_2040_v7).F33() {
					return (_2040_v7).F33()
				}
				return (_2040_v7).F33()
			})() {
				r0 = (_dafny.IntOfUint32((_2039_v6).Cardinality())).Plus((_2040_v7).F34())
				var _2046_v13 D2
				_ = _2046_v13
				_2046_v13 = Companion_D2_.Create_DC7_((_2040_v7).F33())
				var _2047_v14 _dafny.Map
				_ = _2047_v14
				_2047_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2046_v13, (_2040_v7).F33())
				_2047_v14 = (_2047_v14).Update(_2046_v13, true)
				var _2048_v15 _dafny.Array
				_ = _2048_v15
				var _nw358 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw358
				_2048_v15 = _nw358
				var _2049_v16 _dafny.CodePoint
				_ = _2049_v16
				_2049_v16 = _dafny.CodePoint('h')
				var _2050_v17 _dafny.Array
				_ = _2050_v17
				var _nwElement0_68 _dafny.Sequence = _dafny.SeqOf(_2049_v16)
				_ = _nwElement0_68
				var _nw359 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(20))
				_ = _nw359
				(_nw359).ArraySet1(_nwElement0_68, 0)
				(_nw359).ArraySet1(_2039_v6, 1)
				(_nw359).ArraySet1(_dafny.SeqOf(_2049_v16), 2)
				(_nw359).ArraySet1(_2039_v6, 3)
				(_nw359).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg131 _dafny.Int) interface{} {
						return coer131(arg131)
					}
				}((func(_2051_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2052_i3 _dafny.Int) _dafny.CodePoint {
						return _2051_v16
					}
				})(_2049_v16))), (Companion_Default___.SafeIndex((_2040_v7).F34(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer132 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg132 _dafny.Int) interface{} {
						return coer132(arg132)
					}
				}((func(_2053_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2054_i3 _dafny.Int) _dafny.CodePoint {
						return _2053_v16
					}
				})(_2049_v16)))).Cardinality()))).Uint32(), _2049_v16), 4)
				(_nw359).ArraySet1(_2039_v6, 5)
				(_nw359).ArraySet1(_2039_v6, 6)
				(_nw359).ArraySet1(_2039_v6, 7)
				(_nw359).ArraySet1(_2039_v6, 8)
				(_nw359).ArraySet1(_2039_v6, 9)
				(_nw359).ArraySet1(_2039_v6, 10)
				(_nw359).ArraySet1(_2039_v6, 11)
				(_nw359).ArraySet1(_2039_v6, 12)
				(_nw359).ArraySet1(_dafny.SeqOf(_dafny.CodePoint('i')), 13)
				(_nw359).ArraySet1(_2039_v6, 14)
				(_nw359).ArraySet1(_2039_v6, 15)
				(_nw359).ArraySet1(_2039_v6, 16)
				(_nw359).ArraySet1(_2039_v6, 17)
				(_nw359).ArraySet1(_2039_v6, 18)
				(_nw359).ArraySet1(_2039_v6, 19)
				_2050_v17 = _nw359
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2048_v15), 0))
				_ = _index345
				(_2048_v15).ArraySet1(_2050_v17, (_index345).Int())
				var _2055_v18 _dafny.Array
				_ = _2055_v18
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_55
				var _nw360 _dafny.Array
				_ = _nw360
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw360 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) _dafny.Int = (func(_2056_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2057_i4 _dafny.Int) _dafny.Int {
							return (_2057_i4).Times(_2056_v0)
						}
					})(_2030_v0)
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw360 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw360).ArraySet1(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw360).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_2055_v18 = _nw360
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_2055_v18), 0))
				_ = _index346
				(_2055_v18).ArraySet1((_dafny.Zero).Minus((p0).Select((Companion_Default___.SafeIndex(_2030_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)), (_index346).Int())
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2048_v15), 0))
				_ = _index347
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_2055_v18), 0))
				_ = _index348
				var _rhs394 _dafny.Array = _2050_v17
				_ = _rhs394
				var _rhs395 bool = (_2030_v0).Cmp(_2030_v0) > 0
				_ = _rhs395
				var _rhs396 _dafny.Int = _dafny.IntOfInt64(-241)
				_ = _rhs396
				var _lhs378 _dafny.Array = _2048_v15
				_ = _lhs378
				var _lhs379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2048_v15), 0))
				_ = _lhs379
				var _lhs380 *GlobalState = globalState
				_ = _lhs380
				var _lhs381 _dafny.Array = _2055_v18
				_ = _lhs381
				var _lhs382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_2055_v18), 0))
				_ = _lhs382
				(_lhs378).ArraySet1(_rhs394, (_lhs379).Int())
				_lhs380.F15 = _rhs395
				(_lhs381).ArraySet1(_rhs396, (_lhs382).Int())
				var _2058_v19 _dafny.Sequence
				_ = _2058_v19
				_2058_v19 = _dafny.SeqOf((_2040_v7).F33())
				(globalState).F18 = (_2058_v19).Select((Companion_Default___.SafeIndex((_2040_v7).F34(), _dafny.IntOfUint32((_2058_v19).Cardinality()))).Uint32()).(bool)
				(globalState).F15 = (_2040_v7).F33()
			} else {
				(globalState).F14 = Companion_Default___.Fm0(_2030_v0, (_2040_v7).F33(), (func() bool {
					if (_2044_v11).Contains((_2040_v7).F33()) {
						return (_2044_v11).Get((_2040_v7).F33()).(bool)
					}
					return (_2040_v7).F33()
				})(), globalState)
				var _2059_v20 _dafny.CodePoint
				_ = _2059_v20
				_2059_v20 = _dafny.CodePoint('i')
				_2059_v20 = _2059_v20
				var _2060_v21 _dafny.Map
				_ = _2060_v21
				_2060_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v5, _dafny.IntOfInt64(-52))
				var _2061_v22 D10
				_ = _2061_v22
				_2061_v22 = Companion_D10_.Create_DC32_(_2059_v20, Companion_Default___.Fm6((func() _dafny.Int {
					if (_2060_v21).Contains(_2038_v5) {
						return (_2060_v21).Get(_2038_v5).(_dafny.Int)
					}
					return (_2040_v7).F34()
				})(), globalState))
				(globalState).F28 = (_2061_v22).Dtor_cf53()
				var _2062_v23 _dafny.Array
				_ = _2062_v23
				var _nw361 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(3))
				_ = _nw361
				_2062_v23 = _nw361
				var _2063_v24 _dafny.Set
				_ = _2063_v24
				_2063_v24 = _dafny.SetOf((_2040_v7).F33())
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_2062_v23), 0))
				_ = _index349
				(_2062_v23).ArraySet1(_2063_v24, (_index349).Int())
				var _2064_v25 _dafny.Sequence
				_ = _2064_v25
				_2064_v25 = _dafny.SeqOf((_2040_v7).F33(), (_2040_v7).F33(), _2038_v5)
				var _2065_v26 _dafny.Map
				_ = _2065_v26
				_2065_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2064_v25).Cardinality()), (_2040_v7).F33())
				var _2066_v29 _dafny.Array
				_ = _2066_v29
				var _nwElement0_69 _dafny.Map = _2065_v26
				_ = _nwElement0_69
				var _nw362 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(21))
				_ = _nw362
				(_nw362).ArraySet1(_nwElement0_69, 0)
				(_nw362).ArraySet1(_2065_v26, 1)
				(_nw362).ArraySet1(_2065_v26, 2)
				(_nw362).ArraySet1(Companion_Default___.Fm41((_dafny.Zero).Minus((_2040_v7).F34()), (_2040_v7).F33(), !((_2040_v7).F33()), _2038_v5, globalState), 3)
				(_nw362).ArraySet1(func() _dafny.Map {
					var _coll71 = _dafny.NewMapBuilder()
					_ = _coll71
					for _iter84 := _dafny.Iterate((p0).Elements()); ; {
						_compr_71, _ok84 := _iter84()
						if !_ok84 {
							break
						}
						var _2067_v27 _dafny.Int
						_2067_v27 = interface{}(_compr_71).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(p0, _2067_v27) {
							_coll71.Add((_2067_v27).Plus((_2040_v7).F34()), _2038_v5)
						}
					}
					return _coll71.ToMap()
				}(), 4)
				(_nw362).ArraySet1(_2065_v26, 5)
				(_nw362).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(354), (_2040_v7).F33())).Update((_2040_v7).F34(), (_2040_v7).F33()), 6)
				(_nw362).ArraySet1(_2065_v26, 7)
				(_nw362).ArraySet1(func() _dafny.Map {
					var _coll72 = _dafny.NewMapBuilder()
					_ = _coll72
					for _iter85 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-623), _dafny.IntOfInt64(476))); ; {
						_compr_72, _ok85 := _iter85()
						if !_ok85 {
							break
						}
						var _2068_v28 _dafny.Int
						_2068_v28 = interface{}(_compr_72).(_dafny.Int)
						if ((_dafny.IntOfInt64(-623)).Cmp(_2068_v28) <= 0) && ((_2068_v28).Cmp(_dafny.IntOfInt64(476)) < 0) {
							_coll72.Add((_2068_v28).Times((_2040_v7).F34()), (_2040_v7).F33())
						}
					}
					return _coll72.ToMap()
				}(), 8)
				(_nw362).ArraySet1(_2065_v26, 9)
				(_nw362).ArraySet1(_2065_v26, 10)
				(_nw362).ArraySet1((_2065_v26).Merge(_2065_v26), 11)
				(_nw362).ArraySet1(_2065_v26, 12)
				(_nw362).ArraySet1((_2065_v26).Merge(_2065_v26), 13)
				(_nw362).ArraySet1(_2065_v26, 14)
				(_nw362).ArraySet1(_2065_v26, 15)
				(_nw362).ArraySet1(_2065_v26, 16)
				(_nw362).ArraySet1(_2065_v26, 17)
				(_nw362).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2040_v7).F34(), (_2040_v7).F33())).Merge(_2065_v26), 18)
				(_nw362).ArraySet1(_2065_v26, 19)
				(_nw362).ArraySet1((_2065_v26).Merge(_2065_v26), 20)
				_2066_v29 = _nw362
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_2066_v29), 0))
				_ = _index350
				(_2066_v29).ArraySet1(_2065_v26, (_index350).Int())
				var _2069_v30 *C3
				_ = _2069_v30
				var _nw363 *C3 = New_C3_()
				_ = _nw363
				_nw363.Ctor__((_2040_v7).F33(), (_2040_v7).F34(), false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-792))).Uint32(), func(coer133 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg133 _dafny.Int) interface{} {
						return coer133(arg133)
					}
				}((func(_2070_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2071_i5 _dafny.Int) _dafny.CodePoint {
						return _2070_v20
					}
				})(_2059_v20))))
				_2069_v30 = _nw363
				var _2072_v31 _dafny.Sequence
				_ = _2072_v31
				_2072_v31 = _dafny.SeqOf(_2069_v30, _2069_v30)
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_2062_v23), 0))
				_ = _index351
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_2066_v29), 0))
				_ = _index352
				var _rhs397 _dafny.Set = _2063_v24
				_ = _rhs397
				var _rhs398 _dafny.MultiSet = (_2031_v1).Update(Companion_Default___.Fm36(globalState), Companion_Default___.Abs((_dafny.Zero).Minus(((_2060_v21).Update(!(false), _dafny.IntOfUint32((_2072_v31).Cardinality()))).Cardinality())))
				_ = _rhs398
				var _rhs399 bool = (_2064_v25).Select((Companion_Default___.SafeIndex(_2030_v0, _dafny.IntOfUint32((_2064_v25).Cardinality()))).Uint32()).(bool)
				_ = _rhs399
				var _rhs400 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, p0)
				_ = _rhs400
				var _rhs401 _dafny.Map = (_2065_v26).Merge(_2065_v26)
				_ = _rhs401
				var _lhs383 _dafny.Array = _2062_v23
				_ = _lhs383
				var _lhs384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_2062_v23), 0))
				_ = _lhs384
				var _lhs385 *GlobalState = globalState
				_ = _lhs385
				var _lhs386 *GlobalState = globalState
				_ = _lhs386
				var _lhs387 _dafny.Array = _2066_v29
				_ = _lhs387
				var _lhs388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_2066_v29), 0))
				_ = _lhs388
				(_lhs383).ArraySet1(_rhs397, (_lhs384).Int())
				_2031_v1 = _rhs398
				_lhs385.F13 = _rhs399
				_lhs386.F7 = _rhs400
				(_lhs387).ArraySet1(_rhs401, (_lhs388).Int())
				var _2073_v32 _dafny.Array
				_ = _2073_v32
				var _nwElement0_70 _dafny.Int = (_2040_v7).F34()
				_ = _nwElement0_70
				var _nw364 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(4))
				_ = _nw364
				(_nw364).ArraySet1(_nwElement0_70, 0)
				(_nw364).ArraySet1(_dafny.IntOfUint32((_2039_v6).Cardinality()), 1)
				(_nw364).ArraySet1((_2040_v7).F34(), 2)
				(_nw364).ArraySet1((_2040_v7).F34(), 3)
				_2073_v32 = _nw364
				var _2074_v33 D12
				_ = _2074_v33
				_2074_v33 = Companion_D12_.Create_DC38_(_2073_v32)
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_2032_v2), 0))
				_ = _index353
				(_2032_v2).ArraySet1((_2064_v25).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_2040_v7).F34()), _dafny.IntOfUint32((_2064_v25).Cardinality()))).Uint32()).(bool), (_index353).Int())
				var _2075_v34 T0
				_ = _2075_v34
				var _nw365 *C4 = New_C4_()
				_ = _nw365
				_nw365.Ctor__((_2040_v7).F33(), _dafny.IntOfUint32((_2039_v6).Cardinality()))
				_2075_v34 = _nw365
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_2032_v2), 0))
				_ = _index354
				var _rhs402 D12 = (func() D12 {
					if (_2040_v7).F33() {
						return _2074_v33
					}
					return _2074_v33
				})()
				_ = _rhs402
				var _rhs403 bool = !(Companion_Default___.Fm1(_dafny.CodePoint('o'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v5, _dafny.IntOfInt64(673))).Cardinality(), (_2040_v7).F34(), !(_2038_v5), globalState))
				_ = _rhs403
				var _rhs404 _dafny.Int = _dafny.IntOfInt64(406)
				_ = _rhs404
				var _rhs405 T0 = _2075_v34
				_ = _rhs405
				var _rhs406 _dafny.Map = _2060_v21
				_ = _rhs406
				var _lhs389 _dafny.Array = _2032_v2
				_ = _lhs389
				var _lhs390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_2032_v2), 0))
				_ = _lhs390
				var _lhs391 *GlobalState = globalState
				_ = _lhs391
				_2074_v33 = _rhs402
				(_lhs389).ArraySet1(_rhs403, (_lhs390).Int())
				_lhs391.F19 = _rhs404
				_2075_v34 = _rhs405
				_2060_v21 = _rhs406
			}
		} else {
			var _2076_v35 *C5
			_ = _2076_v35
			var _nw366 *C5 = New_C5_()
			_ = _nw366
			_nw366.Ctor__((_2040_v7).F33(), _dafny.Companion_Sequence_.Concatenate(_2039_v6, _2039_v6))
			_2076_v35 = _nw366
			var _nw367 *C5 = New_C5_()
			_ = _nw367
			_nw367.Ctor__((_2040_v7).F33(), _2039_v6)
			_2076_v35 = _nw367
			r0 = _dafny.IntOfUint32((_2039_v6).Cardinality())
			var _2077_v36 T0
			_ = _2077_v36
			var _nw368 *C4 = New_C4_()
			_ = _nw368
			_nw368.Ctor__(_2038_v5, _2030_v0)
			_2077_v36 = _nw368
			var _2078_v37 _dafny.Map
			_ = _2078_v37
			_2078_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2040_v7).F34(), _2077_v36)
			var _2079_v38 _dafny.Array
			_ = _2079_v38
			var _nwElement0_71 T0 = _2077_v36
			_ = _nwElement0_71
			var _nw369 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(24))
			_ = _nw369
			(_nw369).ArraySet1(_nwElement0_71, 0)
			(_nw369).ArraySet1(_2077_v36, 1)
			(_nw369).ArraySet1(_2077_v36, 2)
			(_nw369).ArraySet1(_2077_v36, 3)
			(_nw369).ArraySet1(_2077_v36, 4)
			(_nw369).ArraySet1(_2077_v36, 5)
			(_nw369).ArraySet1(_2077_v36, 6)
			(_nw369).ArraySet1(_2077_v36, 7)
			(_nw369).ArraySet1(_2077_v36, 8)
			(_nw369).ArraySet1(_2077_v36, 9)
			(_nw369).ArraySet1(_2077_v36, 10)
			(_nw369).ArraySet1((func() T0 {
				if (_2078_v37).Contains((_2040_v7).F34()) {
					return (_2078_v37).Get((_2040_v7).F34()).(T0)
				}
				return _2077_v36
			})(), 11)
			(_nw369).ArraySet1(_2077_v36, 12)
			(_nw369).ArraySet1(_2077_v36, 13)
			(_nw369).ArraySet1(_2077_v36, 14)
			(_nw369).ArraySet1(_2077_v36, 15)
			(_nw369).ArraySet1(_2077_v36, 16)
			(_nw369).ArraySet1(_2077_v36, 17)
			(_nw369).ArraySet1(_2077_v36, 18)
			(_nw369).ArraySet1(_2077_v36, 19)
			(_nw369).ArraySet1(_2077_v36, 20)
			(_nw369).ArraySet1(_2077_v36, 21)
			(_nw369).ArraySet1(_2077_v36, 22)
			(_nw369).ArraySet1(_2077_v36, 23)
			_2079_v38 = _nw369
			var _2080_v39 _dafny.Map
			_ = _2080_v39
			_2080_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2039_v6, _2039_v6), _2079_v38)
			_2080_v39 = (_2080_v39).Update(_2039_v6, (func() _dafny.Array {
				if _2038_v5 {
					return _2079_v38
				}
				return _2079_v38
			})())
			if (!((func() bool {
				if true {
					return (_2040_v7).F33()
				}
				return (_2040_v7).F33()
			})())) && ((_2040_v7).F33()) {
				(globalState).F14 = Companion_Default___.SafeModuloInt(_2030_v0, (_2040_v7).F34())
				var _2081_v40 _dafny.MultiSet
				_ = _2081_v40
				_2081_v40 = _dafny.MultiSetOf(_2038_v5, (_2040_v7).F33(), (_2040_v7).F33(), (_2040_v7).F33(), (_2040_v7).F33())
				(globalState).F28 = ((_2081_v40).Intersection(_2081_v40)).IsDisjointFrom((_2081_v40).Update((_2040_v7).F33(), Companion_Default___.Abs(_2030_v0)))
				var _2082_v41 _dafny.Set
				_ = _2082_v41
				_2082_v41 = _dafny.SetOf((_2040_v7).F33())
				var _2083_v42 _dafny.Map
				_ = _2083_v42
				_2083_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2082_v41).Cardinality(), _dafny.IntOfUint32((_2039_v6).Cardinality()))
				var _2084_v43 _dafny.Array
				_ = _2084_v43
				var _nwElement0_72 _dafny.Int = ((_2031_v1).Difference(_2031_v1)).Cardinality()
				_ = _nwElement0_72
				var _nw370 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(19))
				_ = _nw370
				(_nw370).ArraySet1(_nwElement0_72, 0)
				(_nw370).ArraySet1((_2040_v7).F34(), 1)
				(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt((_2040_v7).F34(), (_2040_v7).F34()), 2)
				(_nw370).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kljunxavc")).Cardinality()), 3)
				(_nw370).ArraySet1((_2040_v7).F34(), 4)
				(_nw370).ArraySet1((_2040_v7).F34(), 5)
				(_nw370).ArraySet1(_dafny.IntOfInt64(-367), 6)
				(_nw370).ArraySet1(_2030_v0, 7)
				(_nw370).ArraySet1((func() _dafny.Int {
					if (_2083_v42).Contains(_dafny.IntOfInt64(179)) {
						return (_2083_v42).Get(_dafny.IntOfInt64(179)).(_dafny.Int)
					}
					return (_2040_v7).F34()
				})(), 8)
				(_nw370).ArraySet1(_2030_v0, 9)
				(_nw370).ArraySet1((_2040_v7).F34(), 10)
				(_nw370).ArraySet1((_2040_v7).F34(), 11)
				(_nw370).ArraySet1((_2040_v7).F34(), 12)
				(_nw370).ArraySet1(Companion_Default___.Fm36(globalState), 13)
				(_nw370).ArraySet1(_2030_v0, 14)
				(_nw370).ArraySet1((_2040_v7).F34(), 15)
				(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt((_2040_v7).F34(), (_2040_v7).F34()), 16)
				(_nw370).ArraySet1((func() _dafny.Int {
					if (_2040_v7).F33() {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xgc")).Cardinality())
					}
					return (_2040_v7).F34()
				})(), 17)
				(_nw370).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_2040_v7).F34())), 18)
				_2084_v43 = _nw370
				(globalState).F17 = _2084_v43
				(globalState).F14 = Companion_Default___.SafeDivisionInt(_2030_v0, (_2040_v7).F34())
				var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2084_v43), 0))
				_ = _index355
				(_2084_v43).ArraySet1((_dafny.Zero).Minus(_2030_v0), (_index355).Int())
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_2084_v43), 0))
				_ = _index356
				(_2084_v43).ArraySet1(_2030_v0, (_index356).Int())
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2084_v43), 0))
				_ = _index357
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_2084_v43), 0))
				_ = _index358
				var _rhs407 _dafny.Int = (_2040_v7).F34()
				_ = _rhs407
				var _rhs408 _dafny.Int = _dafny.IntOfInt64(309)
				_ = _rhs408
				var _lhs392 _dafny.Array = _2084_v43
				_ = _lhs392
				var _lhs393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2084_v43), 0))
				_ = _lhs393
				var _lhs394 _dafny.Array = _2084_v43
				_ = _lhs394
				var _lhs395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_2084_v43), 0))
				_ = _lhs395
				(_lhs392).ArraySet1(_rhs407, (_lhs393).Int())
				(_lhs394).ArraySet1(_rhs408, (_lhs395).Int())
			} else {
				var _2085_v44 _dafny.Array
				_ = _2085_v44
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_56
				var _nw371 _dafny.Array
				_ = _nw371
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw371 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) _dafny.Sequence = (func(_2086_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2087_i6 _dafny.Int) _dafny.Sequence {
							return _2086_v6
						}
					})(_2039_v6)
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw371 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw371).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw371).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2085_v44 = _nw371
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2085_v44), 0))
				_ = _index359
				(_2085_v44).ArraySet1(_2039_v6, (_index359).Int())
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2085_v44), 0))
				_ = _index360
				(_2085_v44).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2039_v6, _2039_v6), (_index360).Int())
				(globalState).F15 = true
				(globalState).F15 = (_2040_v7).F33()
				(globalState).F20 = _dafny.Companion_Sequence_.Concatenate((_2085_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2085_v44), 0))).Int()).(_dafny.Sequence), (_2085_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2085_v44), 0))).Int()).(_dafny.Sequence))
				(globalState).F1 = (func() _dafny.Map {
					var _coll73 = _dafny.NewMapBuilder()
					_ = _coll73
					for _iter86 := _dafny.Iterate((_2031_v1).Elements()); ; {
						_compr_73, _ok86 := _iter86()
						if !_ok86 {
							break
						}
						var _2088_v45 _dafny.Int
						_2088_v45 = interface{}(_compr_73).(_dafny.Int)
						if (_2031_v1).Contains(_2088_v45) {
							_coll73.Add((_2088_v45).Times((_2040_v7).F34()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2040_v7).F34(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
								var _coll74 = _dafny.NewBuilder()
								_ = _coll74
								for _iter87 := _dafny.Iterate((_2031_v1).Elements()); ; {
									_compr_74, _ok87 := _iter87()
									if !_ok87 {
										break
									}
									var _2089_v46 _dafny.Int
									_2089_v46 = interface{}(_compr_74).(_dafny.Int)
									if (_2031_v1).Contains(_2089_v46) {
										_coll74.Add((_2089_v46).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true)).Cardinality())).Cardinality()))
									}
								}
								return _coll74.ToSet()
							}()).Cardinality())).Cardinality()))
						}
					}
					return _coll73.ToMap()
				}()).Cardinality()
			}
			var _2090_v47 _dafny.MultiSet
			_ = _2090_v47
			_2090_v47 = _dafny.MultiSetOf(false)
			var _2091_v48 _dafny.Map
			_ = _2091_v48
			_2091_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2040_v7).F34(), _2090_v47)
			var _2092_v49 _dafny.Sequence
			_ = _2092_v49
			_2092_v49 = _dafny.SeqOf(_2039_v6)
			var _2093_v50 D0
			_ = _2093_v50
			_2093_v50 = Companion_D0_.Create_DC2_((func() _dafny.MultiSet {
				if (_2091_v48).Contains(_2030_v0) {
					return (_2091_v48).Get(_2030_v0).(_dafny.MultiSet)
				}
				return _2090_v47
			})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("umiruj"), (_2092_v49).Select((Companion_Default___.SafeIndex(_2030_v0, _dafny.IntOfUint32((_2092_v49).Cardinality()))).Uint32()).(_dafny.Sequence)), (_2040_v7).F33())
			_2093_v50 = Companion_D0_.Create_DC2_(_dafny.MultiSetOf((_2040_v7).F33(), !((_2040_v7).F33()), (_2040_v7).F33()), _2039_v6, false)
		}
		var _hi12 _dafny.Int = _2030_v0
		_ = _hi12
		for _2094_i7 := (_2030_v0).Minus((_dafny.Zero).Minus((_2040_v7).F34())); _2094_i7.Cmp(_hi12) < 0; _2094_i7 = _2094_i7.Plus(_dafny.One) {
			var _2095_v51 _dafny.Array
			_ = _2095_v51
			var _nw372 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
			_ = _nw372
			_2095_v51 = _nw372
			var _2096_v52 _dafny.Array
			_ = _2096_v52
			var _len0_57 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_57
			var _nw373 _dafny.Array
			_ = _nw373
			if _len0_57.Cmp(_dafny.Zero) == 0 {
				_nw373 = _dafny.NewArray(_len0_57)
			} else {
				var _init57 func(_dafny.Int) _dafny.Int = (func(_2097_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2098_i8 _dafny.Int) _dafny.Int {
						return (_2098_i8).Plus(_2097_v0)
					}
				})(_2030_v0)
				_ = _init57
				var _element0_57 = _init57(_dafny.Zero)
				_ = _element0_57
				_nw373 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
				(_nw373).ArraySet1(_element0_57, 0)
				var _nativeLen0_57 = (_len0_57).Int()
				_ = _nativeLen0_57
				for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
					(_nw373).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
				}
			}
			_2096_v52 = _nw373
			var _2099_v53 _dafny.Set
			_ = _2099_v53
			_2099_v53 = _dafny.SetOf(_2096_v52, _2096_v52)
			var _2100_v54 _dafny.CodePoint
			_ = _2100_v54
			_2100_v54 = _dafny.CodePoint('r')
			var _2101_v55 *C2
			_ = _2101_v55
			var _nw374 *C2 = New_C2_()
			_ = _nw374
			_nw374.Ctor__((_dafny.MultiSetOf(_2040_v7)).Cardinality(), _2099_v53, _2038_v5, _dafny.Companion_Sequence_.Update(_2039_v6, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2039_v6).Cardinality()), _dafny.IntOfUint32((_2039_v6).Cardinality()))).Uint32(), _2100_v54))
			_2101_v55 = _nw374
			var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_2095_v51), 0))
			_ = _index361
			(_2095_v51).ArraySet1(_2101_v55, (_index361).Int())
			var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_2095_v51), 0))
			_ = _index362
			(_2095_v51).ArraySet1(_2101_v55, (_index362).Int())
			var _2102_v56 _dafny.Map
			_ = _2102_v56
			_2102_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v5, _dafny.IntOfInt64(-916))
			var _2103_v57 T1
			_ = _2103_v57
			var _nw375 *C3 = New_C3_()
			_ = _nw375
			_nw375.Ctor__((_2040_v7).F33(), _dafny.IntOfInt64(941), !(_2102_v56).Equals(_2102_v56), _2039_v6)
			_2103_v57 = _nw375
			_2103_v57 = _2103_v57
			var _2104_v58 _dafny.MultiSet
			_ = _2104_v58
			_2104_v58 = _dafny.MultiSetOf((_2103_v57).F29())
			var _2105_v59 D11
			_ = _2105_v59
			_2105_v59 = Companion_D11_.Create_DC37_((_2104_v58).Cardinality(), _2038_v5)
			(globalState).F2 = (func(_pat_let48_0 D11) D11 {
				return func(_2106_dt__update__tmp_h0 D11) D11 {
					return func(_pat_let49_0 bool) D11 {
						return func(_2107_dt__update_hcf63_h0 bool) D11 {
							return Companion_D11_.Create_DC37_((_2106_dt__update__tmp_h0).Dtor_cf62(), _2107_dt__update_hcf63_h0)
						}(_pat_let49_0)
					}(false)
				}(_pat_let48_0)
			}(_2105_v59)).Dtor_cf62()
			var _2108_v60 T0
			_ = _2108_v60
			var _nw376 *C1 = New_C1_()
			_ = _nw376
			_nw376.Ctor__(false)
			_2108_v60 = _nw376
			var _2109_v61 _dafny.Map
			_ = _2109_v61
			_2109_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2108_v60, ((_2101_v55).F38()).Cmp((_2101_v55).F38()) != 0)
			_2109_v61 = (_2109_v61).Update(_2108_v60, _2038_v5)
		}
		r0 = _2030_v0
		return r0
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
