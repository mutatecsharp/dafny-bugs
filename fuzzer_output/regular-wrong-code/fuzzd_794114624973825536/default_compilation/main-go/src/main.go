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
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool {
	return ((_dafny.Zero).Minus((func() _dafny.Int {
		if true {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.CodePoint('n'))).Cardinality()
		}
		return _dafny.IntOfInt64(98)
	})())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bdayefsjg"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})))).Cardinality()))) <= 0
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(988))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))
	}))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(61))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
	})))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('o')
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false, !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))), !(false))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf((_dafny.SetOf(true)).Cardinality())).Cardinality(), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(468), _dafny.IntOfInt64(-404))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(468)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(-404)) < 0) {
				_coll0.Add((_4_v0).Times(_dafny.IntOfInt64(203)), _dafny.IntOfInt64(415))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()))).Plus(_dafny.IntOfInt64(665))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false, true)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, p2 D2, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kjyj")).Cardinality())), _dafny.IntOfInt64(-123))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pmifknqvs")).Cardinality())))).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(781)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-166))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(406)
	})))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, p2 _dafny.Set, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fxkcy")).Cardinality()))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_7_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-110)
	}))), _dafny.SeqOf(_dafny.IntOfInt64(194)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(57), _dafny.IntOfInt64(825)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(897)
	})), _dafny.SeqOf(_dafny.IntOfInt64(-853), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("inrv")).Cardinality())))).Cardinality()), _dafny.IntOfInt64(-354), _dafny.IntOfInt64(-817))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(4))))).Union((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(5))).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvdonglnj")).Cardinality()), false)).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_9_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(527)
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, true, false)).Difference(_dafny.SetOf(!(true)))).Intersection((func() _dafny.Set {
		if true {
			return _dafny.SetOf(false, true, (Companion_D4_.Create_DC9_(_dafny.IntOfInt64(293), _dafny.CodePoint('i'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_10_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(767)
			})), _dafny.IntOfInt64(37), false)).Dtor_cf20(), true)
		}
		return _dafny.SetOf(true, true)
	})())
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.Zero).Minus((_dafny.IntOfInt64(190)).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-831))))).Times((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.CodePoint('r'))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(503), false, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality())), true)).Cardinality()).Plus(_dafny.IntOfInt64(829)))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(true)), false, false), _dafny.SeqOf(true))).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-948))) < 0
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC3_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(884))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_11_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(925)
	}))), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_12_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqwbp")).Cardinality())
	}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(769))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_13_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(837)
	})))).Cardinality(), _dafny.IntOfInt64(471))), _dafny.IntOfInt64(819))))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(394)), _dafny.SeqOf(_dafny.IntOfInt64(904)))).Cardinality())).Cmp((_dafny.IntOfInt64(239)).Minus(_dafny.IntOfInt64(156))) >= 0
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Set, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(((_dafny.IntOfInt64(938)).Times(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Times((_dafny.IntOfInt64(634)).Times(_dafny.IntOfInt64(619))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D13 = Companion_D13_.Create_DC25_((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()))), _dafny.IntOfInt64(787), false)
	_ = _source0
	if _source0.Is_DC25() {
		var _14___mcc_h0 _dafny.Int = _source0.Get_().(D13_DC25).Cf42
		_ = _14___mcc_h0
		var _15___mcc_h1 _dafny.Int = _source0.Get_().(D13_DC25).Cf43
		_ = _15___mcc_h1
		var _16___mcc_h2 bool = _source0.Get_().(D13_DC25).Cf44
		_ = _16___mcc_h2
		var _17_cf44 bool = _16___mcc_h2
		_ = _17_cf44
		var _18_cf43 _dafny.Int = _15___mcc_h1
		_ = _18_cf43
		var _19_cf42 _dafny.Int = _14___mcc_h0
		_ = _19_cf42
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("li"), _dafny.UnicodeSeqOfUtf8Bytes("ggkk"))
	} else {
		var _20___mcc_h3 *C2 = _source0.Get_().(D13_DC24).Cf41
		_ = _20___mcc_h3
		var _21_cf41 *C2 = _20___mcc_h3
		_ = _21_cf41
		return _dafny.UnicodeSeqOfUtf8Bytes("vnno")
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('v')
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("hls"), _dafny.UnicodeSeqOfUtf8Bytes("nhuwclrv")) {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_22_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_dafny.IntOfInt64(-387))
		}))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(((Companion_D14_.Create_DC26_(_dafny.MultiSetOf(false))).Dtor_cf45()).Cardinality()), (Companion_D4_.Create_DC9_(_dafny.IntOfInt64(317), _dafny.CodePoint('l'), _dafny.SeqOf(_dafny.IntOfInt64(111), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('v'))).Cardinality()), _dafny.IntOfInt64(447), !(true))).Dtor_cf18())
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(729))).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(24), _dafny.IntOfInt64(239))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(24)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(239)) < 0) {
				_coll1.Add((_23_v0).Times(_dafny.IntOfInt64(758)), _dafny.IntOfInt64(627))
			}
		}
		return _coll1.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(167), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) D2 {
	if !(!(true)) {
		return Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.IntOfInt64(-659))).Cardinality()), !(true), true, true, !(false))
	} else {
		return Companion_D2_.Create_DC5_(_dafny.IntOfInt64(63), true, !(false), false, false)
	}
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false, false)).Union((_dafny.SetOf(false, true)).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, !(!(!(true)))), _dafny.SeqOf(true, true)), _dafny.SeqOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Sequence, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf(true))).Intersection(_dafny.MultiSetOf(_dafny.SeqOf(false, false)))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC7_(_dafny.IntOfInt64(397), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Dtor_cf14(), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(637))).Cardinality(), _dafny.IntOfInt64(462))).Cardinality()), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(449), !(true))))
}
func (_static *CompanionStruct_Default___) Fm35(globalState *GlobalState) _dafny.Set {
	if !(false) || (true) {
		return _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_24_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})))
	} else {
		return _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-615))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_25_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), _dafny.UnicodeSeqOfUtf8Bytes("b"))
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC14_((_dafny.IntOfInt64(814)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sqkmgu")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm37(p0 D1, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.Zero).Minus(_dafny.IntOfInt64(-689)))).Cardinality(), _dafny.IntOfInt64(-952), _dafny.IntOfInt64(228))).Cardinality())).Union((_dafny.SetOf(_dafny.IntOfInt64(621))).Difference(_dafny.SetOf(_dafny.IntOfInt64(-276))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(!(true)), false)).Intersection(_dafny.MultiSetOf(true))).Intersection((_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D8 {
	var _source1 D1 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
	_ = _source1
	if _source1.Is_DC2() {
		var _26___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC2).Cf2
		_ = _26___mcc_h0
		var _27___mcc_h1 bool = _source1.Get_().(D1_DC2).Cf3
		_ = _27___mcc_h1
		var _28___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC2).Cf4
		_ = _28___mcc_h2
		var _29_cf4 _dafny.Int = _28___mcc_h2
		_ = _29_cf4
		var _30_cf3 bool = _27___mcc_h1
		_ = _30_cf3
		var _31_cf2 _dafny.Int = _26___mcc_h0
		_ = _31_cf2
		return Companion_D8_.Create_DC16_(_29_cf4, _30_cf3, _30_cf3, _31_cf2)
	} else {
		var _32___mcc_h3 _dafny.Sequence = _source1.Get_().(D1_DC1).Cf1
		_ = _32___mcc_h3
		var _33_cf1 _dafny.Sequence = _32___mcc_h3
		_ = _33_cf1
		return Companion_D8_.Create_DC16_((_dafny.SetOf(false)).Cardinality(), false, true, (_dafny.SetOf(!(false))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(((func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _34_v1 _dafny.Sequence
					_34_v1 = interface{}(_compr_4).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Contains(_34_v1) {
						_coll4.Add(_34_v1, false)
					}
				}
				return _coll4.ToMap()
			}()).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _35_v2 _dafny.Sequence
				_35_v2 = interface{}(_compr_3).(_dafny.Sequence)
				if (func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _34_v1 _dafny.Sequence
						_34_v1 = interface{}(_compr_5).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Contains(_34_v1) {
							_coll5.Add(_34_v1, false)
						}
					}
					return _coll5.ToMap()
				}()).Contains(_35_v2) {
					_coll3.Add(_35_v2)
				}
			}
			return _coll3.ToSet()
		}()).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_36_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})), _dafny.UnicodeSeqOfUtf8Bytes("kbeim")))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _37_v0 _dafny.Sequence
			_37_v0 = interface{}(_compr_2).(_dafny.Sequence)
			if ((func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((func() _dafny.Map {
					var _coll7 = _dafny.NewMapBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _34_v1 _dafny.Sequence
						_34_v1 = interface{}(_compr_7).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Contains(_34_v1) {
							_coll7.Add(_34_v1, false)
						}
					}
					return _coll7.ToMap()
				}()).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _38_v2 _dafny.Sequence
					_38_v2 = interface{}(_compr_6).(_dafny.Sequence)
					if (func() _dafny.Map {
						var _coll8 = _dafny.NewMapBuilder()
						_ = _coll8
						for _iter8 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Elements()); ; {
							_compr_8, _ok8 := _iter8()
							if !_ok8 {
								break
							}
							var _34_v1 _dafny.Sequence
							_34_v1 = interface{}(_compr_8).(_dafny.Sequence)
							if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.UnicodeSeqOfUtf8Bytes("qwc"), _dafny.UnicodeSeqOfUtf8Bytes("prgm"))).Contains(_34_v1) {
								_coll8.Add(_34_v1, false)
							}
						}
						return _coll8.ToMap()
					}()).Contains(_38_v2) {
						_coll6.Add(_38_v2)
					}
				}
				return _coll6.ToSet()
			}()).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_36_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), _dafny.UnicodeSeqOfUtf8Bytes("kbeim")))).Contains(_37_v0) {
				_coll2.Add(_37_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false)))
			}
		}
		return _coll2.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("snxgcwim")).Cardinality()), _dafny.IntOfInt64(430)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false)), _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ejalkcx")).Cardinality()))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('r'), _dafny.CodePoint('y'), _dafny.CodePoint('p'), _dafny.CodePoint('r'))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC20_(!(!(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_39_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.UnicodeSeqOfUtf8Bytes("fc")))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var _40_v0 bool
	_ = _40_v0
	_40_v0 = true
	var _41_v1 _dafny.CodePoint
	_ = _41_v1
	_41_v1 = _dafny.CodePoint('r')
	var _42_v2 _dafny.Sequence
	_ = _42_v2
	_42_v2 = _dafny.SeqOf(_41_v1, _41_v1)
	var _43_v3 _dafny.Sequence
	_ = _43_v3
	_43_v3 = _dafny.SeqOf(_42_v2)
	var _44_v4 _dafny.Array
	_ = _44_v4
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
	_ = _nw0
	_44_v4 = _nw0
	var _45_v5 _dafny.Map
	_ = _45_v5
	_45_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, _44_v4)
	var _46_v6 _dafny.Array
	_ = _46_v6
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
	_ = _nw1
	_46_v6 = _nw1
	var _47_v7 *C4
	_ = _47_v7
	var _nw2 *C4 = New_C4_()
	_ = _nw2
	_nw2.Ctor__((_40_v0) == (_40_v0), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.Fm17(_43_v3, (_45_v5).Cardinality(), p0, _dafny.IntOfUint32((_42_v2).Cardinality()), globalState)))), _46_v6)
	_47_v7 = _nw2
	_40_v0 = (_47_v7).F7()
	var _48_v8 _dafny.Sequence
	_ = _48_v8
	_48_v8 = _dafny.SeqOf((_47_v7).F7(), (_47_v7).F7(), _40_v0)
	var _49_v9 _dafny.Map
	_ = _49_v9
	_49_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_47_v7).F7(), (_dafny.MultiSetFromSeq(_48_v8)).Cardinality())
	var _50_v10 *C6
	_ = _50_v10
	var _nw3 *C6 = New_C6_()
	_ = _nw3
	_nw3.Ctor__((_49_v9).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, p0)).Update((_47_v7).F7(), p0)), p0)
	_50_v10 = _nw3
	var _51_v11 *C5
	_ = _51_v11
	var _nw4 *C5 = New_C5_()
	_ = _nw4
	_nw4.Ctor__(_dafny.CodePoint('m'))
	_51_v11 = _nw4
	var _52_v12 _dafny.Map
	_ = _52_v12
	_52_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v11, _50_v10.F5)
	_52_v12 = (_52_v12).Update(_51_v11, _dafny.IntOfInt64(821))
	(_50_v10).F5 = Companion_Default___.SafeDivisionInt(p0, _50_v10.F5)
	var _53_v13 *C2
	_ = _53_v13
	var _nw5 *C2 = New_C2_()
	_ = _nw5
	_nw5.Ctor__(_40_v0, (_dafny.Zero).Minus(p0), _46_v6)
	_53_v13 = _nw5
	var _54_v14 _dafny.Array
	_ = _54_v14
	var _nwElement0_0 *C2 = _53_v13
	_ = _nwElement0_0
	var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
	_ = _nw6
	(_nw6).ArraySet1(_nwElement0_0, 0)
	(_nw6).ArraySet1(_53_v13, 1)
	(_nw6).ArraySet1(_53_v13, 2)
	(_nw6).ArraySet1(_53_v13, 3)
	(_nw6).ArraySet1(_53_v13, 4)
	(_nw6).ArraySet1(_53_v13, 5)
	(_nw6).ArraySet1(_53_v13, 6)
	(_nw6).ArraySet1(_53_v13, 7)
	(_nw6).ArraySet1(_53_v13, 8)
	(_nw6).ArraySet1(_53_v13, 9)
	(_nw6).ArraySet1(_53_v13, 10)
	(_nw6).ArraySet1(_53_v13, 11)
	(_nw6).ArraySet1(_53_v13, 12)
	(_nw6).ArraySet1(_53_v13, 13)
	(_nw6).ArraySet1(_53_v13, 14)
	(_nw6).ArraySet1(_53_v13, 15)
	(_nw6).ArraySet1(_53_v13, 16)
	(_nw6).ArraySet1(_53_v13, 17)
	(_nw6).ArraySet1(_53_v13, 18)
	(_nw6).ArraySet1(_53_v13, 19)
	_54_v14 = _nw6
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_54_v14), 0))
	_ = _index0
	(_54_v14).ArraySet1(_53_v13, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_54_v14), 0))
	_ = _index1
	(_54_v14).ArraySet1(_53_v13, (_index1).Int())
	r0 = false
	r1 = ((_53_v13).F12()).Cmp(p0) < 0
	var _55_v15 D6
	_ = _55_v15
	_55_v15 = Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("tw"))
	var _pat_let_tv0 = _42_v2
	_ = _pat_let_tv0
	var _pat_let_tv1 = _42_v2
	_ = _pat_let_tv1
	r2 = _dafny.IntOfUint32((func(_source2 D6) _dafny.Sequence {
		if _source2.Is_DC12() {
			var _56___mcc_h0 bool = _source2.Get_().(D6_DC12).Cf23
			_ = _56___mcc_h0
			var _57___mcc_h1 _dafny.CodePoint = _source2.Get_().(D6_DC12).Cf24
			_ = _57___mcc_h1
			var _58_cf24 _dafny.CodePoint = _57___mcc_h1
			_ = _58_cf24
			var _59_cf23 bool = _56___mcc_h0
			_ = _59_cf23
			return _pat_let_tv0
		} else {
			var _60___mcc_h2 _dafny.Sequence = _source2.Get_().(D6_DC11).Cf22
			_ = _60___mcc_h2
			var _61_cf22 _dafny.Sequence = _60___mcc_h2
			_ = _61_cf22
			return _pat_let_tv1
		}
	}(_55_v15)).Cardinality())
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _62_v0 bool
	_ = _62_v0
	_62_v0 = false
	var _63_v1 _dafny.Sequence
	_ = _63_v1
	_63_v1 = _dafny.SeqOf(_62_v0)
	var _64_v2 _dafny.Sequence
	_ = _64_v2
	_64_v2 = _63_v1
	var _65_v3 _dafny.Int
	_ = _65_v3
	_65_v3 = _dafny.IntOfInt64(124)
	var _66_v4 _dafny.Array
	_ = _66_v4
	var _nwElement0_1 _dafny.Sequence = _63_v1
	_ = _nwElement0_1
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(11))
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_1, 0)
	(_nw7).ArraySet1(_63_v1, 1)
	(_nw7).ArraySet1((_64_v2), 2)
	(_nw7).ArraySet1(_63_v1, 3)
	(_nw7).ArraySet1(_dafny.Companion_Sequence_.Update(_63_v1, (Companion_Default___.SafeIndex(_65_v3, _dafny.IntOfUint32((_63_v1).Cardinality()))).Uint32(), _62_v0), 4)
	(_nw7).ArraySet1(_63_v1, 5)
	(_nw7).ArraySet1(_63_v1, 6)
	(_nw7).ArraySet1(_63_v1, 7)
	(_nw7).ArraySet1(_63_v1, 8)
	(_nw7).ArraySet1(_dafny.SeqOf(_62_v0), 9)
	(_nw7).ArraySet1(_63_v1, 10)
	_66_v4 = _nw7
	var _67_globalState *GlobalState
	_ = _67_globalState
	var _nw8 *GlobalState = New_GlobalState_()
	_ = _nw8
	_nw8.Ctor__(_66_v4, false, _dafny.IntOfInt64(47))
	_67_globalState = _nw8
	var _68_v5 _dafny.Set
	_ = _68_v5
	_68_v5 = _dafny.SetOf(_65_v3)
	if (_68_v5).IsProperSubsetOf(_68_v5) {
		var _69_v6 bool
		_ = _69_v6
		var _70_v7 bool
		_ = _70_v7
		var _71_v8 _dafny.Int
		_ = _71_v8
		var _out0 bool
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.Int
		_ = _out2
		_out0, _out1, _out2 = Companion_Default___.M0(_65_v3, _67_globalState)
		_69_v6 = _out0
		_70_v7 = _out1
		_71_v8 = _out2
		_62_v0 = _62_v0
		if !(_70_v7) {
			var _72_v9 _dafny.Map
			_ = _72_v9
			_72_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-674), false)
			var _73_v11 _dafny.Sequence
			_ = _73_v11
			_73_v11 = _dafny.SeqOf(_65_v3)
			var _74_v12 _dafny.Map
			_ = _74_v12
			_74_v12 = _72_v9
			var _75_v13 _dafny.CodePoint
			_ = _75_v13
			_75_v13 = _dafny.CodePoint('v')
			var _76_v14 _dafny.Sequence
			_ = _76_v14
			_76_v14 = _dafny.SeqOf(_75_v13)
			var _77_v17 _dafny.Array
			_ = _77_v17
			var _nwElement0_2 _dafny.Map = _72_v9
			_ = _nwElement0_2
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_2, 0)
			(_nw9).ArraySet1(_72_v9, 1)
			(_nw9).ArraySet1(_72_v9, 2)
			(_nw9).ArraySet1(_72_v9, 3)
			(_nw9).ArraySet1(_72_v9, 4)
			(_nw9).ArraySet1(func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_73_v11).Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _78_v10 _dafny.Int
					_78_v10 = interface{}(_compr_9).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_73_v11, _78_v10) {
						_coll9.Add((_78_v10).Times(_dafny.IntOfInt64(789)), _70_v7)
					}
				}
				return _coll9.ToMap()
			}(), 5)
			(_nw9).ArraySet1(_72_v9, 6)
			(_nw9).ArraySet1((_74_v12), 7)
			(_nw9).ArraySet1(Companion_Default___.Fm34(_62_v0, _75_v13, _73_v11, _67_globalState), 8)
			(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm17(_dafny.SeqOf(_76_v14, _76_v14), _71_v8, _71_v8, _71_v8, _67_globalState), _69_v6), 9)
			(_nw9).ArraySet1(_72_v9, 10)
			(_nw9).ArraySet1(_72_v9, 11)
			(_nw9).ArraySet1(func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v8, _71_v8)).Keys().Elements()); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _79_v15 _dafny.Int
					_79_v15 = interface{}(_compr_10).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v8, _71_v8)).Contains(_79_v15) {
						_coll10.Add((_79_v15).Plus(_65_v3), _62_v0)
					}
				}
				return _coll10.ToMap()
			}(), 12)
			(_nw9).ArraySet1(func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_73_v11).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _80_v16 _dafny.Int
					_80_v16 = interface{}(_compr_11).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_73_v11, _80_v16) {
						_coll11.Add((_80_v16).Minus(_71_v8), _70_v7)
					}
				}
				return _coll11.ToMap()
			}(), 13)
			(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _70_v7), 14)
			(_nw9).ArraySet1(_72_v9, 15)
			(_nw9).ArraySet1(_72_v9, 16)
			_77_v17 = _nw9
			var _81_v18 *C2
			_ = _81_v18
			var _nw10 *C2 = New_C2_()
			_ = _nw10
			_nw10.Ctor__(_69_v6, (_dafny.Zero).Minus(_65_v3), _77_v17)
			_81_v18 = _nw10
			var _82_v19 D6
			_ = _82_v19
			_82_v19 = Companion_D6_.Create_DC11_(_76_v14)
			var _83_v20 _dafny.Array
			_ = _83_v20
			var _nwElement0_3 D6 = _82_v19
			_ = _nwElement0_3
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(7))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_3, 0)
			(_nw11).ArraySet1(_82_v19, 1)
			(_nw11).ArraySet1(_82_v19, 2)
			(_nw11).ArraySet1(_82_v19, 3)
			(_nw11).ArraySet1(_82_v19, 4)
			(_nw11).ArraySet1(_82_v19, 5)
			(_nw11).ArraySet1(_82_v19, 6)
			_83_v20 = _nw11
			var _rhs0 _dafny.Int = (_81_v18).F12()
			_ = _rhs0
			var _rhs1 _dafny.Int = _65_v3
			_ = _rhs1
			var _rhs2 _dafny.Array = (func() _dafny.Array {
				if _62_v0 {
					return _83_v20
				}
				return _83_v20
			})()
			_ = _rhs2
			var _rhs3 _dafny.Int = (_65_v3).Plus(_65_v3)
			_ = _rhs3
			var _lhs0 *GlobalState = _67_globalState
			_ = _lhs0
			var _lhs1 *GlobalState = _67_globalState
			_ = _lhs1
			_lhs0.F2 = _rhs0
			_lhs1.F2 = _rhs1
			_83_v20 = _rhs2
			_65_v3 = _rhs3
			_69_v6 = true
			var _84_v21 _dafny.Map
			_ = _84_v21
			_84_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_v18).F11(), (_81_v18).F12())
			var _85_v22 _dafny.Map
			_ = _85_v22
			_85_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_69_v6), (func() _dafny.Int {
				if (_84_v21).Contains((_81_v18).F11()) {
					return (_84_v21).Get((_81_v18).F11()).(_dafny.Int)
				}
				return _71_v8
			})())
			_71_v8 = (func() _dafny.Int {
				if (_85_v22).Contains(true) {
					return (_85_v22).Get(true).(_dafny.Int)
				}
				return _dafny.IntOfInt64(325)
			})()
			_85_v22 = (_85_v22).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_86_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_87_i0 _dafny.Int) _dafny.CodePoint {
					return _86_v13
				}
			})(_75_v13))), _76_v14), _71_v8)
		} else {
			var _88_v23 _dafny.Array
			_ = _88_v23
			var _nw12 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
			_ = _nw12
			_88_v23 = _nw12
			_88_v23 = _88_v23
			var _89_v24 _dafny.Sequence
			_ = _89_v24
			_89_v24 = _dafny.SeqOf(_65_v3, _71_v8, _71_v8)
			var _90_v25 _dafny.Sequence
			_ = _90_v25
			_90_v25 = _dafny.UnicodeSeqOfUtf8Bytes("ro")
			var _91_v26 _dafny.CodePoint
			_ = _91_v26
			_91_v26 = _dafny.CodePoint('j')
			var _92_v27 _dafny.Array
			_ = _92_v27
			var _nwElement0_4 bool = (func() bool {
				if _62_v0 {
					return _70_v7
				}
				return _62_v0
			})()
			_ = _nwElement0_4
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_4, 0)
			(_nw13).ArraySet1(_69_v6, 1)
			(_nw13).ArraySet1(Companion_Default___.Fm21(_dafny.IntOfUint32((_90_v25).Cardinality()), _91_v26, false, _70_v7, _67_globalState), 2)
			(_nw13).ArraySet1(_70_v7, 3)
			(_nw13).ArraySet1(false, 4)
			(_nw13).ArraySet1((func() bool {
				if _69_v6 {
					return Companion_Default___.Fm21(_71_v8, _91_v26, _70_v7, !(_70_v7), _67_globalState)
				}
				return _69_v6
			})(), 5)
			(_nw13).ArraySet1(_62_v0, 6)
			(_nw13).ArraySet1(!((func() bool {
				if _70_v7 {
					return !(Companion_Default___.Fm2(_dafny.CodePoint('f'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg20 _dafny.Int) interface{} {
							return coer20(arg20)
						}
					}((func(_93_v7 bool) func(_dafny.Int) _dafny.Map {
						return func(_94_i1 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v7, _93_v7)
						}
					})(_70_v7))), _70_v7, _67_globalState))
				}
				return _70_v7
			})()), 7)
			(_nw13).ArraySet1(!(true) || (_70_v7), 8)
			_92_v27 = _nw13
			var _rhs4 _dafny.Sequence = _89_v24
			_ = _rhs4
			var _rhs5 bool = (func() bool {
				if _69_v6 {
					return (_69_v6) == (_70_v7)
				}
				return _62_v0
			})()
			_ = _rhs5
			var _rhs6 _dafny.Array = _92_v27
			_ = _rhs6
			var _rhs7 bool = !_dafny.Companion_Sequence_.Equal(_89_v24, _89_v24)
			_ = _rhs7
			var _rhs8 _dafny.Int = _65_v3
			_ = _rhs8
			_89_v24 = _rhs4
			_69_v6 = _rhs5
			_92_v27 = _rhs6
			_69_v6 = _rhs7
			_71_v8 = _rhs8
			var _95_v28 _dafny.Array
			_ = _95_v28
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
			_ = _nw14
			_95_v28 = _nw14
			var _96_v29 *C1
			_ = _96_v29
			var _nw15 *C1 = New_C1_()
			_ = _nw15
			_nw15.Ctor__(_71_v8, _95_v28)
			_96_v29 = _nw15
			var _97_v30 _dafny.Array
			_ = _97_v30
			var _nwElement0_5 _dafny.Sequence = _64_v2
			_ = _nwElement0_5
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(14))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_5, 0)
			(_nw16).ArraySet1(_64_v2, 1)
			(_nw16).ArraySet1(_64_v2, 2)
			(_nw16).ArraySet1(_64_v2, 3)
			(_nw16).ArraySet1(Companion_Default___.Fm44(_67_globalState), 4)
			(_nw16).ArraySet1(Companion_Default___.Fm44(_67_globalState), 5)
			(_nw16).ArraySet1(_64_v2, 6)
			(_nw16).ArraySet1(_64_v2, 7)
			(_nw16).ArraySet1(_64_v2, 8)
			(_nw16).ArraySet1(Companion_Default___.Fm44(_67_globalState), 9)
			(_nw16).ArraySet1(_64_v2, 10)
			(_nw16).ArraySet1(_64_v2, 11)
			(_nw16).ArraySet1(_64_v2, 12)
			(_nw16).ArraySet1(_64_v2, 13)
			_97_v30 = _nw16
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_97_v30), 0))
			_ = _index2
			(_97_v30).ArraySet1(_64_v2, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_97_v30), 0))
			_ = _index3
			var _rhs9 _dafny.Sequence = _64_v2
			_ = _rhs9
			var _rhs10 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_90_v25, _90_v25)).Cardinality())
			_ = _rhs10
			var _rhs11 _dafny.Int = (_96_v29).F14()
			_ = _rhs11
			var _rhs12 bool = (_63_v1).Select((Companion_Default___.SafeIndex((_96_v29).F14(), _dafny.IntOfUint32((_63_v1).Cardinality()))).Uint32()).(bool)
			_ = _rhs12
			var _rhs13 _dafny.CodePoint = _91_v26
			_ = _rhs13
			var _lhs2 _dafny.Array = _97_v30
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_97_v30), 0))
			_ = _lhs3
			var _lhs4 *GlobalState = _67_globalState
			_ = _lhs4
			(_lhs2).ArraySet1(_rhs9, (_lhs3).Int())
			_65_v3 = _rhs10
			_lhs4.F2 = _rhs11
			_70_v7 = _rhs12
			_91_v26 = _rhs13
			var _98_v31 _dafny.Sequence
			_ = _98_v31
			_98_v31 = _dafny.SeqOf(_90_v25)
			var _99_v32 _dafny.Set
			_ = _99_v32
			_99_v32 = _dafny.SetOf(_92_v27, _92_v27)
			var _100_v33 _dafny.Set
			_ = _100_v33
			_100_v33 = _dafny.SetOf(_96_v29)
			_69_v6 = ((_dafny.IntOfUint32((_90_v25).Cardinality())).Minus(Companion_Default___.Fm17(_98_v31, (_89_v24).Select((Companion_Default___.SafeIndex((_96_v29).F14(), _dafny.IntOfUint32((_89_v24).Cardinality()))).Uint32()).(_dafny.Int), (_99_v32).Cardinality(), (_96_v29).F14(), _67_globalState))).Cmp(Companion_Default___.Fm17(_dafny.SeqOf(_90_v25, _90_v25, _dafny.Companion_Sequence_.Update(_90_v25, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.IntOfUint32((_90_v25).Cardinality()))).Uint32(), _dafny.CodePoint('b'))), (_96_v29).F14(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, (_100_v33).Cardinality())).Cardinality(), _65_v3, _67_globalState)) <= 0
		}
		_62_v0 = _70_v7
		var _101_v34 bool
		_ = _101_v34
		var _102_v35 bool
		_ = _102_v35
		var _103_v36 _dafny.Int
		_ = _103_v36
		var _out3 bool
		_ = _out3
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Int
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.M0((_dafny.Zero).Minus(_65_v3), _67_globalState)
		_101_v34 = _out3
		_102_v35 = _out4
		_103_v36 = _out5
	} else {
		_62_v0 = _62_v0
		var _104_v37 _dafny.Map
		_ = _104_v37
		_104_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, !(_62_v0) || (_62_v0))
		var _105_v38 _dafny.CodePoint
		_ = _105_v38
		_105_v38 = _dafny.CodePoint('b')
		_104_v37 = (_104_v37).Update(Companion_Default___.Fm21(_65_v3, _105_v38, _62_v0, false, _67_globalState), _62_v0)
		var _106_v39 _dafny.Sequence
		_ = _106_v39
		_106_v39 = _dafny.UnicodeSeqOfUtf8Bytes("fgobb")
		_106_v39 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_107_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_108_i2 _dafny.Int) _dafny.CodePoint {
				return _107_v38
			}
		})(_105_v38))), _106_v39)
		_62_v0 = (func() bool {
			if _62_v0 {
				return (_68_v5).IsDisjointFrom(_68_v5)
			}
			return _62_v0
		})()
		if _62_v0 {
			_62_v0 = (false) && (_62_v0)
			var _109_v40 _dafny.Map
			_ = _109_v40
			_109_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _65_v3)
			var _110_v41 bool
			_ = _110_v41
			var _111_v42 bool
			_ = _111_v42
			var _112_v43 _dafny.Int
			_ = _112_v43
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			var _out8 _dafny.Int
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M0((func() _dafny.Int {
				if (_109_v40).Contains(_65_v3) {
					return (_109_v40).Get(_65_v3).(_dafny.Int)
				}
				return _65_v3
			})(), _67_globalState)
			_110_v41 = _out6
			_111_v42 = _out7
			_112_v43 = _out8
			_110_v41 = _62_v0
			var _113_v44 _dafny.MultiSet
			_ = _113_v44
			_113_v44 = _dafny.MultiSetOf(_110_v41)
			var _114_v45 _dafny.Map
			_ = _114_v45
			_114_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_106_v39).Cardinality()), (_113_v44).Update(_110_v41, Companion_Default___.Abs(_65_v3)))
			_114_v45 = (_114_v45).Update(_65_v3, _113_v44)
			var _115_v46 _dafny.MultiSet
			_ = _115_v46
			_115_v46 = _dafny.MultiSetOf(_112_v43, _65_v3, _112_v43, _65_v3)
			var _116_v47 *C3
			_ = _116_v47
			var _nw17 *C3 = New_C3_()
			_ = _nw17
			_nw17.Ctor__(_115_v46, _110_v41)
			_116_v47 = _nw17
			var _117_v48 _dafny.Sequence
			_ = _117_v48
			_117_v48 = _dafny.SeqOf(_116_v47)
			var _118_v49 _dafny.Set
			_ = _118_v49
			_118_v49 = _dafny.SetOf(_109_v40)
			var _119_v50 _dafny.Array
			_ = _119_v50
			var _nwElement0_6 *C3 = _116_v47
			_ = _nwElement0_6
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(6))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_6, 0)
			(_nw18).ArraySet1(_116_v47, 1)
			(_nw18).ArraySet1((_117_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_118_v49).Cardinality(), _112_v43)).Cardinality()), _dafny.IntOfUint32((_117_v48).Cardinality()))).Uint32()).(*C3), 2)
			(_nw18).ArraySet1(_116_v47, 3)
			(_nw18).ArraySet1(_116_v47, 4)
			(_nw18).ArraySet1(_116_v47, 5)
			_119_v50 = _nw18
			_119_v50 = _119_v50
		} else {
			var _120_v51 _dafny.Array
			_ = _120_v51
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw19
			_120_v51 = _nw19
			var _121_v52 _dafny.Array
			_ = _121_v52
			var _nwElement0_7 bool = _62_v0
			_ = _nwElement0_7
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(10))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_7, 0)
			(_nw20).ArraySet1(!(_62_v0), 1)
			(_nw20).ArraySet1(_62_v0, 2)
			(_nw20).ArraySet1(Companion_Default___.Fm21(_65_v3, _105_v38, _62_v0, _62_v0, _67_globalState), 3)
			(_nw20).ArraySet1(_62_v0, 4)
			(_nw20).ArraySet1(_62_v0, 5)
			(_nw20).ArraySet1(Companion_Default___.Fm23(_106_v39, _65_v3, _65_v3, _62_v0, _67_globalState), 6)
			(_nw20).ArraySet1(false, 7)
			(_nw20).ArraySet1(_62_v0, 8)
			(_nw20).ArraySet1(_62_v0, 9)
			_121_v52 = _nw20
			var _122_v53 _dafny.Set
			_ = _122_v53
			_122_v53 = _dafny.SetOf(_121_v52, _121_v52, _121_v52, _121_v52, _121_v52)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_120_v51), 0))
			_ = _index4
			(_120_v51).ArraySet1((_122_v53).Cardinality(), (_index4).Int())
			var _123_v54 _dafny.Sequence
			_ = _123_v54
			_123_v54 = _dafny.SeqOf(Companion_Default___.Fm25(_62_v0, _65_v3, _67_globalState))
			var _124_v55 _dafny.Map
			_ = _124_v55
			_124_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _65_v3)
			var _125_v56 _dafny.Sequence
			_ = _125_v56
			_125_v56 = _dafny.SeqOf(_124_v55)
			var _126_v57 D4
			_ = _126_v57
			_126_v57 = Companion_D4_.Create_DC9_(_dafny.IntOfUint32((_125_v56).Cardinality()), _105_v38, _dafny.SeqOf(_65_v3), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("muudfjey")).Cardinality()), Companion_Default___.Fm23(_106_v39, _65_v3, (_dafny.MultiSetOf(Companion_Default___.Fm17(_123_v54, _65_v3, _dafny.IntOfInt64(-758), _65_v3, _67_globalState))).Cardinality(), false, _67_globalState))
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_120_v51), 0))
			_ = _index5
			(_120_v51).ArraySet1(Companion_Default___.Fm17(_123_v54, _dafny.IntOfUint32((_106_v39).Cardinality()), (_126_v57).Dtor_cf19(), _65_v3, _67_globalState), (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_121_v52), 0))
			_ = _index6
			(_121_v52).ArraySet1(_62_v0, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_121_v52), 0))
			_ = _index7
			(_121_v52).ArraySet1(_62_v0, (_index7).Int())
			_124_v55 = (_124_v55).Update(_62_v0, (_120_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_120_v51), 0))).Int()).(_dafny.Int))
			var _127_v58 _dafny.Map
			_ = _127_v58
			_127_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v38, (_dafny.IntOfUint32((_106_v39).Cardinality())).Plus(_65_v3))
			_127_v58 = (_127_v58).Update(_105_v38, (_120_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_120_v51), 0))).Int()).(_dafny.Int))
			var _128_v59 _dafny.Array
			_ = _128_v59
			var _len0_0 _dafny.Int = _dafny.One
			_ = _len0_0
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Map = (func(_129_v3 _dafny.Int, _130_v0 bool) func(_dafny.Int) _dafny.Map {
					return func(_131_i3 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v3, _130_v0)
					}
				})(_65_v3, _62_v0)
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
			_128_v59 = _nw21
			var _132_v60 *C7
			_ = _132_v60
			var _nw22 *C7 = New_C7_()
			_ = _nw22
			_nw22.Ctor__(_128_v59)
			_132_v60 = _nw22
			_132_v60 = _132_v60
		}
	}
	var _hi0 _dafny.Int = (_65_v3).Plus(_65_v3)
	_ = _hi0
	for _133_i4 := _65_v3; _133_i4.Cmp(_hi0) < 0; _133_i4 = _133_i4.Plus(_dafny.One) {
		var _134_v61 _dafny.CodePoint
		_ = _134_v61
		_134_v61 = _dafny.CodePoint('k')
		var _135_v62 _dafny.MultiSet
		_ = _135_v62
		_135_v62 = _dafny.MultiSetOf(_134_v61, Companion_Default___.Fm5(_65_v3, _67_globalState))
		var _136_v63 D1
		_ = _136_v63
		_136_v63 = Companion_D1_.Create_DC2_((_135_v62).Cardinality(), _62_v0, _dafny.IntOfInt64(502))
		var _137_v64 _dafny.Set
		_ = _137_v64
		_137_v64 = _dafny.SetOf(false, false)
		var _138_v65 _dafny.Sequence
		_ = _138_v65
		_138_v65 = _dafny.SeqOf(_137_v64, _137_v64)
		var _139_v66 _dafny.Array
		_ = _139_v66
		var _nwElement0_8 _dafny.Int = ((_138_v65).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_65_v3)), _dafny.IntOfUint32((_138_v65).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
		_ = _nwElement0_8
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.One)
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_8, 0)
		_139_v66 = _nw23
		var _140_v67 _dafny.Sequence
		_ = _140_v67
		_140_v67 = _dafny.UnicodeSeqOfUtf8Bytes("qm")
		var _141_v68 _dafny.Map
		_ = _141_v68
		_141_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v66, _dafny.IntOfUint32((_140_v67).Cardinality()))
		var _142_v69 _dafny.Array
		_ = _142_v69
		var _nwElement0_9 bool = _62_v0
		_ = _nwElement0_9
		var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
		_ = _nw24
		(_nw24).ArraySet1(_nwElement0_9, 0)
		(_nw24).ArraySet1((func() bool {
			if _62_v0 {
				return _62_v0
			}
			return _62_v0
		})(), 1)
		(_nw24).ArraySet1(_62_v0, 2)
		(_nw24).ArraySet1((_136_v63).Dtor_cf3(), 3)
		(_nw24).ArraySet1(_62_v0, 4)
		(_nw24).ArraySet1(_62_v0, 5)
		(_nw24).ArraySet1(_62_v0, 6)
		(_nw24).ArraySet1(_62_v0, 7)
		(_nw24).ArraySet1(_62_v0, 8)
		(_nw24).ArraySet1(_62_v0, 9)
		(_nw24).ArraySet1(false, 10)
		(_nw24).ArraySet1(!(_141_v68).Contains(_139_v66), 11)
		(_nw24).ArraySet1(true, 12)
		(_nw24).ArraySet1(_62_v0, 13)
		_142_v69 = _nw24
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_142_v69), 0))
		_ = _index8
		(_142_v69).ArraySet1(_62_v0, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_142_v69), 0))
		_ = _index9
		(_142_v69).ArraySet1(_62_v0, (_index9).Int())
		_62_v0 = !(_62_v0)
		var _143_v70 _dafny.MultiSet
		_ = _143_v70
		_143_v70 = _dafny.MultiSetOf(_dafny.IntOfUint32((_140_v67).Cardinality()))
		var _144_v71 *C3
		_ = _144_v71
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__(_143_v70, true)
		_144_v71 = _nw25
		var _145_v72 _dafny.Map
		_ = _145_v72
		_145_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, (_144_v71).F10())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_142_v69), 0))
		_ = _index10
		(_142_v69).ArraySet1((func() bool {
			if (_145_v72).Contains(_62_v0) {
				return (_145_v72).Get(_62_v0).(bool)
			}
			return (_142_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_142_v69), 0))).Int()).(bool)
		})(), (_index10).Int())
	}
	_62_v0 = _62_v0
	var _146_v73 _dafny.Array
	_ = _146_v73
	var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
	_ = _nw26
	_146_v73 = _nw26
	_146_v73 = _146_v73
	var _147_v74 _dafny.MultiSet
	_ = _147_v74
	_147_v74 = _dafny.MultiSetOf(_62_v0)
	var _148_i5 _dafny.Int
	_ = _148_i5
	_148_i5 = _dafny.Zero
	{
		for !((func() _dafny.MultiSet {
			if _62_v0 {
				return _147_v74
			}
			return _147_v74
		})()).Equals((_147_v74).Union(_147_v74)) {
			{
				if (_148_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_148_i5 = (_148_i5).Plus(_dafny.One)
				var _149_v75 _dafny.Sequence
				_ = _149_v75
				_149_v75 = _dafny.SeqOf(_68_v5)
				var _150_v76 _dafny.Map
				_ = _150_v76
				_150_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_149_v75).Cardinality()), _dafny.IntOfUint32((_63_v1).Cardinality())), _dafny.SeqOf(_65_v3, _65_v3))
				_150_v76 = _150_v76
				var _151_v77 _dafny.Map
				_ = _151_v77
				_151_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _65_v3)
				var _152_v78 _dafny.Sequence
				_ = _152_v78
				_152_v78 = _dafny.SeqOf(_dafny.IntOfInt64(742))
				var _153_v79 _dafny.Map
				_ = _153_v79
				_153_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_152_v78).Cardinality()), _62_v0)
				var _154_v80 _dafny.Map
				_ = _154_v80
				_154_v80 = _153_v79
				var _155_v81 _dafny.MultiSet
				_ = _155_v81
				_155_v81 = _dafny.MultiSetOf(_154_v80, _154_v80)
				var _156_v82 _dafny.Map
				_ = _156_v82
				_156_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v78, _155_v81)
				var _157_v83 *C6
				_ = _157_v83
				var _nw27 *C6 = New_C6_()
				_ = _nw27
				_nw27.Ctor__(_151_v77, ((_156_v82).Merge(_156_v82)).Cardinality())
				_157_v83 = _nw27
				var _158_v84 _dafny.CodePoint
				_ = _158_v84
				_158_v84 = _dafny.CodePoint('m')
				_158_v84 = _158_v84
				_63_v1 = _63_v1
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _159_v85 _dafny.CodePoint
	_ = _159_v85
	_159_v85 = _dafny.CodePoint('y')
	var _160_v86 *C5
	_ = _160_v86
	var _nw28 *C5 = New_C5_()
	_ = _nw28
	_nw28.Ctor__(_159_v85)
	_160_v86 = _nw28
	var _161_i6 _dafny.Int
	_ = _161_i6
	_161_i6 = _dafny.Zero
	{
		for !(_62_v0) {
			{
				if (_161_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_161_i6 = (_161_i6).Plus(_dafny.One)
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_146_v73), 0))
				_ = _index11
				(_146_v73).ArraySet1(_62_v0, (_index11).Int())
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_146_v73), 0))
				_ = _index12
				(_146_v73).ArraySet1(_62_v0, (_index12).Int())
				(_67_globalState).F2 = _65_v3
				var _162_v87 _dafny.Array
				_ = _162_v87
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_1
				var _nw29 _dafny.Array
				_ = _nw29
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw29 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.Map = (func(_163_v3 _dafny.Int, _164_v0 bool) func(_dafny.Int) _dafny.Map {
						return func(_165_i7 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v3, _164_v0)
						}
					})(_65_v3, _62_v0)
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
				_162_v87 = _nw29
				var _166_v88 T0
				_ = _166_v88
				var _nw30 *C2 = New_C2_()
				_ = _nw30
				_nw30.Ctor__(_62_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_65_v3)), _162_v87)
				_166_v88 = _nw30
				var _167_v89 _dafny.Array
				_ = _167_v89
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_2
				var _nw31 _dafny.Array
				_ = _nw31
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw31 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) _dafny.Int = (func(_168_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_169_i8 _dafny.Int) _dafny.Int {
							return (_169_i8).Times(_168_v3)
						}
					})(_65_v3)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw31 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw31).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw31).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_167_v89 = _nw31
				var _170_v90 _dafny.Map
				_ = _170_v90
				_170_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _167_v89)
				var _171_v91 D2
				_ = _171_v91
				_171_v91 = Companion_D2_.Create_DC5_((_170_v90).Cardinality(), (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool), _62_v0, !(_62_v0), (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool))
				var _172_v92 _dafny.Sequence
				_ = _172_v92
				_172_v92 = _dafny.SeqOf(_171_v91)
				var _173_v93 D9
				_ = _173_v93
				_173_v93 = Companion_D9_.Create_DC19_(_166_v88, (_dafny.MultiSetOf(_172_v92)).Cardinality(), _65_v3)
				var _174_v94 _dafny.Array
				_ = _174_v94
				var _nwElement0_10 T0 = _166_v88
				_ = _nwElement0_10
				var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(5))
				_ = _nw32
				(_nw32).ArraySet1(_nwElement0_10, 0)
				(_nw32).ArraySet1(_166_v88, 1)
				(_nw32).ArraySet1((_173_v93).Dtor_cf34(), 2)
				(_nw32).ArraySet1(_166_v88, 3)
				(_nw32).ArraySet1(_166_v88, 4)
				_174_v94 = _nw32
				var _175_v95 _dafny.Sequence
				_ = _175_v95
				_175_v95 = _dafny.SeqOf(_166_v88, _166_v88)
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_174_v94), 0))
				_ = _index13
				(_174_v94).ArraySet1((_175_v95).Select((Companion_Default___.SafeIndex(_65_v3, _dafny.IntOfUint32((_175_v95).Cardinality()))).Uint32()).(T0), (_index13).Int())
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_174_v94), 0))
				_ = _index14
				(_174_v94).ArraySet1((_173_v93).Dtor_cf34(), (_index14).Int())
				var _176_v96 _dafny.Sequence
				_ = _176_v96
				_176_v96 = _dafny.UnicodeSeqOfUtf8Bytes("jfj")
				_176_v96 = _dafny.Companion_Sequence_.Concatenate(_176_v96, _dafny.Companion_Sequence_.Concatenate(_176_v96, _dafny.UnicodeSeqOfUtf8Bytes("wye")))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _177_v97 _dafny.Set
	_ = _177_v97
	_177_v97 = _dafny.SetOf(_62_v0, _62_v0, _62_v0, false, _62_v0)
	var _178_v98 _dafny.Set
	_ = _178_v98
	_178_v98 = _177_v97
	var _179_v99 _dafny.Map
	_ = _179_v99
	_179_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _178_v98)
	var _rhs14 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _178_v98)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _178_v98))
	_ = _rhs14
	var _rhs15 _dafny.Int = _65_v3
	_ = _rhs15
	var _lhs5 *GlobalState = _67_globalState
	_ = _lhs5
	_179_v99 = _rhs14
	_lhs5.F2 = _rhs15
	var _180_v100 _dafny.Map
	_ = _180_v100
	_180_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _62_v0)
	var _181_v101 _dafny.Sequence
	_ = _181_v101
	_181_v101 = _dafny.SeqOf(_159_v85)
	var _182_v102 _dafny.MultiSet
	_ = _182_v102
	_182_v102 = _dafny.MultiSetOf(_65_v3, _dafny.IntOfUint32((_181_v101).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_183_i9 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()), _65_v3)
	var _184_v103 _dafny.Sequence
	_ = _184_v103
	_184_v103 = _dafny.SeqOf((_182_v102).Cardinality(), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_v85, _182_v102)).Cardinality())).Cardinality())
	var _185_v106 _dafny.Map
	_ = _185_v106
	_185_v106 = _180_v100
	var _186_v107 _dafny.Map
	_ = _186_v107
	_186_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _65_v3)
	var _187_v108 *C6
	_ = _187_v108
	var _nw33 *C6 = New_C6_()
	_ = _nw33
	_nw33.Ctor__(_186_v107, _65_v3)
	_187_v108 = _nw33
	var _188_v109 *C0
	_ = _188_v109
	var _nw34 *C0 = New_C0_()
	_ = _nw34
	_nw34.Ctor__(Companion_Default___.Fm9(_65_v3, _67_globalState))
	_188_v109 = _nw34
	var _189_v110 _dafny.Sequence
	_ = _189_v110
	_189_v110 = _dafny.SeqOf(_188_v109, _188_v109)
	var _190_v111 _dafny.Map
	_ = _190_v111
	_190_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v108, (_dafny.Zero).Minus(_dafny.IntOfUint32((_189_v110).Cardinality())))
	var _191_v112 _dafny.Array
	_ = _191_v112
	var _nwElement0_11 _dafny.Map = _180_v100
	_ = _nwElement0_11
	var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(29))
	_ = _nw35
	(_nw35).ArraySet1(_nwElement0_11, 0)
	(_nw35).ArraySet1(_180_v100, 1)
	(_nw35).ArraySet1(_180_v100, 2)
	(_nw35).ArraySet1(_180_v100, 3)
	(_nw35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _62_v0)).Update(_dafny.IntOfInt64(147), _62_v0), 4)
	(_nw35).ArraySet1(_180_v100, 5)
	(_nw35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _62_v0)).Update(_65_v3, _62_v0), 6)
	(_nw35).ArraySet1(Companion_Default___.Fm34(_62_v0, (_160_v86).F6(), _184_v103, _67_globalState), 7)
	(_nw35).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_181_v101).Cardinality()), _62_v0), 8)
	(_nw35).ArraySet1(_180_v100, 9)
	(_nw35).ArraySet1(_180_v100, 10)
	(_nw35).ArraySet1(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_182_v102).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _192_v104 _dafny.Int
			_192_v104 = interface{}(_compr_12).(_dafny.Int)
			if (_182_v102).Contains(_192_v104) {
				_coll12.Add((_192_v104).Plus(_65_v3), _62_v0)
			}
		}
		return _coll12.ToMap()
	}(), 11)
	(_nw35).ArraySet1(_180_v100, 12)
	(_nw35).ArraySet1(_180_v100, 13)
	(_nw35).ArraySet1((_180_v100).Update(_65_v3, _62_v0), 14)
	(_nw35).ArraySet1(_180_v100, 15)
	(_nw35).ArraySet1(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(80), _dafny.IntOfInt64(355))); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _193_v105 _dafny.Int
			_193_v105 = interface{}(_compr_13).(_dafny.Int)
			if ((_dafny.IntOfInt64(80)).Cmp(_193_v105) <= 0) && ((_193_v105).Cmp(_dafny.IntOfInt64(355)) < 0) {
				_coll13.Add((_193_v105).Plus(_65_v3), _62_v0)
			}
		}
		return _coll13.ToMap()
	}(), 16)
	(_nw35).ArraySet1(Companion_Default___.Fm34(_62_v0, _159_v85, _184_v103, _67_globalState), 17)
	(_nw35).ArraySet1((_185_v106), 18)
	(_nw35).ArraySet1(_180_v100, 19)
	(_nw35).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v3, _62_v0), 20)
	(_nw35).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_190_v111).Contains(_187_v108) {
			return (_190_v111).Get(_187_v108).(_dafny.Int)
		}
		return _dafny.IntOfInt64(684)
	})(), _62_v0), 21)
	(_nw35).ArraySet1(_180_v100, 22)
	(_nw35).ArraySet1(_180_v100, 23)
	(_nw35).ArraySet1(_180_v100, 24)
	(_nw35).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_65_v3)).Cardinality(), _62_v0), 25)
	(_nw35).ArraySet1(_180_v100, 26)
	(_nw35).ArraySet1(_180_v100, 27)
	(_nw35).ArraySet1(_180_v100, 28)
	_191_v112 = _nw35
	var _194_v113 *C7
	_ = _194_v113
	var _nw36 *C7 = New_C7_()
	_ = _nw36
	_nw36.Ctor__(_191_v112)
	_194_v113 = _nw36
	var _195_v114 T1
	_ = _195_v114
	var _nw37 *C7 = New_C7_()
	_ = _nw37
	_nw37.Ctor__(_191_v112)
	_195_v114 = _nw37
	var _196_v115 _dafny.Sequence
	_ = _196_v115
	_196_v115 = _dafny.SeqOf(_195_v114, _195_v114)
	var _197_v116 *C2
	_ = _197_v116
	var _nw38 *C2 = New_C2_()
	_ = _nw38
	_nw38.Ctor__(_62_v0, (_184_v103).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_196_v115).Cardinality()), _dafny.IntOfUint32((_184_v103).Cardinality()))).Uint32()).(_dafny.Int), _191_v112)
	_197_v116 = _nw38
	var _198_v117 _dafny.Int
	_ = _198_v117
	var _199_v118 _dafny.CodePoint
	_ = _199_v118
	var _out9 _dafny.Int
	_ = _out9
	var _out10 _dafny.CodePoint
	_ = _out10
	_out9, _out10 = (_194_v113).M1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(119))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}((func(_200_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_201_i10 _dafny.Int) _dafny.Int {
			return _200_v3
		}
	})(_65_v3))), ((_188_v109).F13()).Cmp((_197_v116).F12()) == 0, _181_v101, _67_globalState)
	_198_v117 = _out9
	_199_v118 = _out10
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
	_ = _index15
	(_146_v73).ArraySet1((_dafny.SetOf((_188_v109).F13())).IsSubsetOf(_68_v5), (_index15).Int())
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
	_ = _index16
	(_146_v73).ArraySet1(Companion_Default___.Fm21(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(71), _198_v117), _199_v118, !((_197_v116).F11()) || (_62_v0), _62_v0, _67_globalState), (_index16).Int())
	_62_v0 = !(((_197_v116).F11()) == ((_147_v74).IsSubsetOf(_147_v74)))
	var _202_v119 D8
	_ = _202_v119
	_202_v119 = Companion_D8_.Create_DC15_(_146_v73)
	var _203_v120 _dafny.MultiSet
	_ = _203_v120
	_203_v120 = _dafny.MultiSetOf(_202_v119)
	var _204_v121 _dafny.Sequence
	_ = _204_v121
	_204_v121 = _dafny.SeqOf(_182_v102, _182_v102)
	var _205_v122 _dafny.Map
	_ = _205_v122
	_205_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool), _62_v0)
	var _206_v123 _dafny.MultiSet
	_ = _206_v123
	_206_v123 = _dafny.MultiSetOf(_184_v103)
	var _207_v124 _dafny.Map
	_ = _207_v124
	_207_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v101, (_dafny.Zero).Minus((_147_v74).Cardinality()))
	var _208_v125 _dafny.Array
	_ = _208_v125
	var _nwElement0_12 _dafny.Int = ((func() _dafny.MultiSet {
		if (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool) {
			return _203_v120
		}
		return _203_v120
	})()).Cardinality()
	_ = _nwElement0_12
	var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(9))
	_ = _nw39
	(_nw39).ArraySet1(_nwElement0_12, 0)
	(_nw39).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_204_v121, (Companion_Default___.SafeIndex((_197_v116).F12(), _dafny.IntOfUint32((_204_v121).Cardinality()))).Uint32(), _dafny.MultiSetOf(_198_v117, (_197_v116).F12())), _204_v121)).Cardinality()), 1)
	(_nw39).ArraySet1((func() _dafny.Int {
		if _62_v0 {
			return (_197_v116).F12()
		}
		return ((_182_v102).Update((_205_v122).Cardinality(), Companion_Default___.Abs(_198_v117))).Cardinality()
	})(), 2)
	(_nw39).ArraySet1((func() _dafny.Int {
		if (_206_v123).Contains(_dafny.SeqOf(_65_v3)) {
			return (_206_v123).Multiplicity(_dafny.SeqOf(_65_v3))
		}
		return _65_v3
	})(), 3)
	(_nw39).ArraySet1((func() _dafny.Int {
		if (_207_v124).Contains(_181_v101) {
			return (_207_v124).Get(_181_v101).(_dafny.Int)
		}
		return _187_v108.F5
	})(), 4)
	(_nw39).ArraySet1((_65_v3).Times((_184_v103).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_181_v101).Cardinality()), _dafny.IntOfUint32((_184_v103).Cardinality()))).Uint32()).(_dafny.Int)), 5)
	(_nw39).ArraySet1((_188_v109).F13(), 6)
	(_nw39).ArraySet1(_dafny.IntOfInt64(689), 7)
	(_nw39).ArraySet1((_188_v109).F13(), 8)
	_208_v125 = _nw39
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))
	_ = _index17
	(_208_v125).ArraySet1((_188_v109).F13(), (_index17).Int())
	var _209_v126 _dafny.Map
	_ = _209_v126
	_209_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v125, (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool))
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))
	_ = _index18
	(_208_v125).ArraySet1((_209_v126).Cardinality(), (_index18).Int())
	var _source3 D3 = Companion_D3_.Create_DC7_(_dafny.IntOfUint32((_63_v1).Cardinality()), Companion_Default___.SafeDivisionInt(_187_v108.F5, (_208_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))).Int()).(_dafny.Int)))
	_ = _source3
	if _source3.Is_DC7() {
		var _210___mcc_h0 _dafny.Int = _source3.Get_().(D3_DC7).Cf13
		_ = _210___mcc_h0
		var _211___mcc_h1 _dafny.Int = _source3.Get_().(D3_DC7).Cf14
		_ = _211___mcc_h1
		var _212_cf14 _dafny.Int = _211___mcc_h1
		_ = _212_cf14
		var _213_cf13 _dafny.Int = _210___mcc_h0
		_ = _213_cf13
		var _214_v127 _dafny.Int
		_ = _214_v127
		var _215_v128 _dafny.CodePoint
		_ = _215_v128
		var _out11 _dafny.Int
		_ = _out11
		var _out12 _dafny.CodePoint
		_ = _out12
		_out11, _out12 = (_194_v113).M1(Companion_Default___.Fm12(!((_197_v116).F11()), _187_v108.F5, _177_v97, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_197_v116).F12(), (_68_v5).Cardinality()), _67_globalState), (_197_v116).F11(), _181_v101, _67_globalState)
		_214_v127 = _out11
		_215_v128 = _out12
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
		_ = _index19
		(_146_v73).ArraySet1((_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool), (_index19).Int())
		var _216_v130 _dafny.Map
		_ = _216_v130
		_216_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_cf14, _213_cf13)
		var _217_v131 _dafny.Map
		_ = _217_v131
		_217_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm24(func() _dafny.Set {
			var _coll14 = _dafny.NewBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(676), _dafny.IntOfInt64(163))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _218_v129 _dafny.Int
				_218_v129 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(676)).Cmp(_218_v129) <= 0) && ((_218_v129).Cmp(_dafny.IntOfInt64(163)) < 0) {
					_coll14.Add(Companion_Default___.SafeModuloInt(_218_v129, (_177_v97).Cardinality()))
				}
			}
			return _coll14.ToSet()
		}(), _67_globalState)).Times((_216_v130).Cardinality()), (_188_v109).F13())
		_216_v130 = _217_v131
		var _219_v132 _dafny.Array
		_ = _219_v132
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw40
		_219_v132 = _nw40
		var _220_v133 D9
		_ = _220_v133
		_220_v133 = Companion_D9_.Create_DC18_(_219_v132)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))
		_ = _index20
		var _rhs16 D9 = _220_v133
		_ = _rhs16
		var _rhs17 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_181_v101).Cardinality()), ((_197_v116).F12()).Plus(_212_cf14))
		_ = _rhs17
		var _rhs18 _dafny.Array = _146_v73
		_ = _rhs18
		var _rhs19 _dafny.Int = ((_dafny.IntOfInt64(332)).Plus(_dafny.IntOfUint32((_181_v101).Cardinality()))).Plus(_187_v108.F5)
		_ = _rhs19
		var _lhs6 _dafny.Array = _208_v125
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))
		_ = _lhs7
		_220_v133 = _rhs16
		_213_cf13 = _rhs17
		_146_v73 = _rhs18
		(_lhs6).ArraySet1(_rhs19, (_lhs7).Int())
	} else {
		var _221___mcc_h2 _dafny.CodePoint = _source3.Get_().(D3_DC6).Cf12
		_ = _221___mcc_h2
		var _222_cf12 _dafny.CodePoint = _221___mcc_h2
		_ = _222_cf12
		_182_v102 = (_182_v102).Intersection((_182_v102).Union(_182_v102))
		_68_v5 = (func() _dafny.Set {
			if (_197_v116).F11() {
				return _68_v5
			}
			return _68_v5
		})()
		_65_v3 = (func() _dafny.Int {
			if _62_v0 {
				return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_197_v116)).Cardinality()), _187_v108.F5)
			}
			return (_197_v116).F12()
		})()
		_62_v0 = true
	}
	var _223_v134 D4
	_ = _223_v134
	_223_v134 = Companion_D4_.Create_DC8_(_188_v109)
	var _source4 D4 = _223_v134
	_ = _source4
	if _source4.Is_DC9() {
		var _224___mcc_h3 _dafny.Int = _source4.Get_().(D4_DC9).Cf16
		_ = _224___mcc_h3
		var _225___mcc_h4 _dafny.CodePoint = _source4.Get_().(D4_DC9).Cf17
		_ = _225___mcc_h4
		var _226___mcc_h5 _dafny.Sequence = _source4.Get_().(D4_DC9).Cf18
		_ = _226___mcc_h5
		var _227___mcc_h6 _dafny.Int = _source4.Get_().(D4_DC9).Cf19
		_ = _227___mcc_h6
		var _228___mcc_h7 bool = _source4.Get_().(D4_DC9).Cf20
		_ = _228___mcc_h7
		var _229_cf20 bool = _228___mcc_h7
		_ = _229_cf20
		var _230_cf19 _dafny.Int = _227___mcc_h6
		_ = _230_cf19
		var _231_cf18 _dafny.Sequence = _226___mcc_h5
		_ = _231_cf18
		var _232_cf17 _dafny.CodePoint = _225___mcc_h4
		_ = _232_cf17
		var _233_cf16 _dafny.Int = _224___mcc_h3
		_ = _233_cf16
		_62_v0 = _229_cf20
		_198_v117 = (_194_v113).Fm0((_188_v109).F13(), _67_globalState)
		var _234_v135 _dafny.Int
		_ = _234_v135
		var _235_v136 _dafny.CodePoint
		_ = _235_v136
		var _out13 _dafny.Int
		_ = _out13
		var _out14 _dafny.CodePoint
		_ = _out14
		_out13, _out14 = (_197_v116).M1(_dafny.SeqOf(_dafny.IntOfUint32((_63_v1).Cardinality())), Companion_Default___.Fm23(_181_v101, (_182_v102).Cardinality(), (_197_v116).F12(), _229_cf20, _67_globalState), Companion_Default___.Fm25(true, _187_v108.F5, _67_globalState), _67_globalState)
		_234_v135 = _out13
		_235_v136 = _out14
		if _62_v0 {
			_186_v107 = (_186_v107).Update((_197_v116).F11(), _65_v3)
			var _236_v137 _dafny.Array
			_ = _236_v137
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_3
			var _nw41 _dafny.Array
			_ = _nw41
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw41 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_237_cf18 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_238_i11 _dafny.Int) _dafny.Sequence {
						return _237_cf18
					}
				})(_231_cf18)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw41 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw41).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw41).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_236_v137 = _nw41
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))
			_ = _index21
			(_236_v137).ArraySet1(_184_v103, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))
			_ = _index22
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
			_ = _index23
			var _rhs20 _dafny.Map = (((_187_v108).F4()).Merge(_186_v107)).Merge((_187_v108).F4())
			_ = _rhs20
			var _rhs21 _dafny.Int = _198_v117
			_ = _rhs21
			var _rhs22 _dafny.Int = (_188_v109).F13()
			_ = _rhs22
			var _rhs23 _dafny.Sequence = _231_cf18
			_ = _rhs23
			var _rhs24 bool = !(!((_dafny.MultiSetFromSeq(_63_v1)).Update((_197_v116).F11(), Companion_Default___.Abs((((_205_v122).Update(false, (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool))).Update(_62_v0, (_197_v116).F11())).Cardinality()))).Equals(_dafny.MultiSetFromSeq(_63_v1)))
			_ = _rhs24
			var _lhs8 *GlobalState = _67_globalState
			_ = _lhs8
			var _lhs9 _dafny.Array = _236_v137
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))
			_ = _lhs10
			var _lhs11 _dafny.Array = _146_v73
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
			_ = _lhs12
			_186_v107 = _rhs20
			_lhs8.F2 = _rhs21
			_65_v3 = _rhs22
			(_lhs9).ArraySet1(_rhs23, (_lhs10).Int())
			(_lhs11).ArraySet1(_rhs24, (_lhs12).Int())
			var _239_v138 _dafny.Map
			_ = _239_v138
			_239_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v113, (_236_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))).Int()).(_dafny.Sequence))
			var _240_v139 _dafny.Sequence
			_ = _240_v139
			_240_v139 = _dafny.SeqOf((_236_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))).Int()).(_dafny.Sequence), (_236_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_236_v137), 0))).Int()).(_dafny.Sequence))
			_239_v138 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v113, _184_v103)).Update(_194_v113, (_240_v139).Select((Companion_Default___.SafeIndex((_188_v109).F13(), _dafny.IntOfUint32((_240_v139).Cardinality()))).Uint32()).(_dafny.Sequence))).Merge(_239_v138)
			var _241_v140 D9
			_ = _241_v140
			_241_v140 = Companion_D9_.Create_DC20_((_197_v116).F11())
			_241_v140 = Companion_Default___.Fm45((_147_v74).IsProperSubsetOf(_dafny.MultiSetOf(Companion_Default___.Fm21(_234_v135, _dafny.CodePoint('f'), (_197_v116).F11(), !((_197_v116).F11()), _67_globalState))), (_197_v116).F12(), _67_globalState)
			var _242_v141 _dafny.Map
			_ = _242_v141
			_242_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v117, _181_v101)
			_208_v125 = (func() _dafny.Array {
				if (_160_v86).Fm7(_dafny.IntOfInt64(380), (_197_v116).F11(), _242_v141, _67_globalState) {
					return _208_v125
				}
				return _208_v125
			})()
		} else {
			var _243_v142 _dafny.Sequence
			_ = _243_v142
			_243_v142 = _dafny.SeqOf(_dafny.SeqOf(false, (func() bool {
				if (_180_v100).Contains((_197_v116).F12()) {
					return (_180_v100).Get((_197_v116).F12()).(bool)
				}
				return _229_cf20
			})()), (func() _dafny.Sequence {
				if (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool) {
					return _63_v1
				}
				return _63_v1
			})(), _63_v1)
			var _244_v143 T0
			_ = _244_v143
			var _nw42 *C2 = New_C2_()
			_ = _nw42
			_nw42.Ctor__((_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool), _187_v108.F5, _191_v112)
			_244_v143 = _nw42
			var _245_v144 _dafny.Map
			_ = _245_v144
			_245_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_244_v143, _243_v142)
			_243_v142 = (func() _dafny.Sequence {
				if (_245_v144).Contains(_244_v143) {
					return (_245_v144).Get(_244_v143).(_dafny.Sequence)
				}
				return _243_v142
			})()
			_235_v136 = (_160_v86).F6()
			_189_v110 = _189_v110
			_194_v113 = (func() *C7 {
				if (_182_v102).IsProperSubsetOf(_182_v102) {
					return _194_v113
				}
				return _194_v113
			})()
			var _246_v145 _dafny.Array
			_ = _246_v145
			var _nwElement0_13 _dafny.CodePoint = (_181_v101).Select((Companion_Default___.SafeIndex((_188_v109).F13(), _dafny.IntOfUint32((_181_v101).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _nwElement0_13
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(12))
			_ = _nw43
			(_nw43).ArraySet1CodePoint(_nwElement0_13, 0)
			(_nw43).ArraySet1CodePoint(_159_v85, 1)
			(_nw43).ArraySet1CodePoint((_181_v101).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_181_v101).Cardinality()), _dafny.IntOfUint32((_181_v101).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
			(_nw43).ArraySet1CodePoint(_159_v85, 3)
			(_nw43).ArraySet1CodePoint((_160_v86).F6(), 4)
			(_nw43).ArraySet1CodePoint(_199_v118, 5)
			(_nw43).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(true) {
					return (Companion_D3_.Create_DC6_(_199_v118)).Dtor_cf12()
				}
				return (_160_v86).F6()
			})(), 6)
			(_nw43).ArraySet1CodePoint((_160_v86).F6(), 7)
			(_nw43).ArraySet1CodePoint((_160_v86).F6(), 8)
			(_nw43).ArraySet1CodePoint(_232_cf17, 9)
			(_nw43).ArraySet1CodePoint((_160_v86).F6(), 10)
			(_nw43).ArraySet1CodePoint((_160_v86).F6(), 11)
			_246_v145 = _nw43
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_246_v145), 0))
			_ = _index24
			(_246_v145).ArraySet1CodePoint((_160_v86).F6(), (_index24).Int())
			var _247_v146 _dafny.Map
			_ = _247_v146
			_247_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v0, _177_v97)
			var _248_v147 _dafny.Map
			_ = _248_v147
			_248_v147 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_181_v101).Cardinality()), ((func() _dafny.Set {
				if (_247_v146).Contains(_229_cf20) {
					return (_247_v146).Get(_229_cf20).(_dafny.Set)
				}
				return _177_v97
			})()).Cardinality())).Update(_187_v108.F5, _65_v3)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_246_v145), 0))
			_ = _index25
			(_246_v145).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _229_cf20 {
					return _235_v136
				}
				return (_160_v86).F6()
			})(), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_246_v145), 0))
			_ = _index26
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_246_v145), 0))
			_ = _index27
			var _rhs25 bool = (_197_v116).F11()
			_ = _rhs25
			var _rhs26 _dafny.CodePoint = _199_v118
			_ = _rhs26
			var _rhs27 _dafny.Map = _248_v147
			_ = _rhs27
			var _rhs28 _dafny.CodePoint = (_181_v101).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-851), (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("lbovk"), _dafny.UnicodeSeqOfUtf8Bytes("kydcrfrb"))).Cardinality()), _dafny.IntOfUint32((_181_v101).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs28
			var _lhs13 _dafny.Array = _246_v145
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_246_v145), 0))
			_ = _lhs14
			var _lhs15 _dafny.Array = _246_v145
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_246_v145), 0))
			_ = _lhs16
			_62_v0 = _rhs25
			(_lhs13).ArraySet1CodePoint(_rhs26, (_lhs14).Int())
			_248_v147 = _rhs27
			(_lhs15).ArraySet1CodePoint(_rhs28, (_lhs16).Int())
		}
	} else {
		var _249___mcc_h8 *C0 = _source4.Get_().(D4_DC8).Cf15
		_ = _249___mcc_h8
		var _250_cf15 *C0 = _249___mcc_h8
		_ = _250_cf15
		if ((_250_cf15).F13()).Cmp((_dafny.IntOfInt64(-485)).Minus((_188_v109).F13())) == 0 {
			_204_v121 = _204_v121
			(_67_globalState).F2 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm9(_dafny.IntOfInt64(836), _67_globalState), (_208_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))).Int()).(_dafny.Int)))
			(_187_v108).F5 = (func() _dafny.Int {
				if (_186_v107).Contains((func() bool {
					if true {
						return (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool)
					}
					return Companion_Default___.Fm2(_dafny.CodePoint('r'), _dafny.SeqOf(_205_v122), false, _67_globalState)
				})()) {
					return (_186_v107).Get((func() bool {
						if true {
							return (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool)
						}
						return Companion_Default___.Fm2(_dafny.CodePoint('r'), _dafny.SeqOf(_205_v122), false, _67_globalState)
					})()).(_dafny.Int)
				}
				return (_198_v117).Plus((_250_cf15).F13())
			})()
			(_67_globalState).F2 = _198_v117
			_65_v3 = Companion_Default___.Fm9(((_250_cf15).F13()).Times((_197_v116).F12()), _67_globalState)
		} else {
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))
			_ = _index28
			(_146_v73).ArraySet1(true, (_index28).Int())
			_68_v5 = (_68_v5).Difference(_68_v5)
			var _251_v148 _dafny.Array
			_ = _251_v148
			var _nw44 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
			_ = _nw44
			_251_v148 = _nw44
			_251_v148 = _251_v148
			var _252_v149 _dafny.Sequence
			_ = _252_v149
			_252_v149 = _dafny.SeqOf(_197_v116, _197_v116)
			var _253_v150 D13
			_ = _253_v150
			_253_v150 = Companion_D13_.Create_DC24_(_197_v116)
			var _254_v151 _dafny.Array
			_ = _254_v151
			var _nwElement0_14 *C2 = _197_v116
			_ = _nwElement0_14
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(23))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_14, 0)
			(_nw45).ArraySet1(_197_v116, 1)
			(_nw45).ArraySet1(_197_v116, 2)
			(_nw45).ArraySet1((_252_v149).Select((Companion_Default___.SafeIndex((_197_v116).F12(), _dafny.IntOfUint32((_252_v149).Cardinality()))).Uint32()).(*C2), 3)
			(_nw45).ArraySet1(_197_v116, 4)
			(_nw45).ArraySet1(_197_v116, 5)
			(_nw45).ArraySet1(_197_v116, 6)
			(_nw45).ArraySet1(_197_v116, 7)
			(_nw45).ArraySet1((func() *C2 {
				if (_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool) {
					return _197_v116
				}
				return _197_v116
			})(), 8)
			(_nw45).ArraySet1(_197_v116, 9)
			(_nw45).ArraySet1(_197_v116, 10)
			(_nw45).ArraySet1(_197_v116, 11)
			(_nw45).ArraySet1(_197_v116, 12)
			(_nw45).ArraySet1(_197_v116, 13)
			(_nw45).ArraySet1(_197_v116, 14)
			(_nw45).ArraySet1((func() *C2 {
				if (_197_v116).F11() {
					return _197_v116
				}
				return (_253_v150).Dtor_cf41()
			})(), 15)
			(_nw45).ArraySet1(_197_v116, 16)
			(_nw45).ArraySet1((_252_v149).Select((Companion_Default___.SafeIndex((_208_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_252_v149).Cardinality()))).Uint32()).(*C2), 17)
			(_nw45).ArraySet1(_197_v116, 18)
			(_nw45).ArraySet1(_197_v116, 19)
			(_nw45).ArraySet1(_197_v116, 20)
			(_nw45).ArraySet1(_197_v116, 21)
			(_nw45).ArraySet1(_197_v116, 22)
			_254_v151 = _nw45
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_254_v151), 0))
			_ = _index29
			(_254_v151).ArraySet1(_197_v116, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_254_v151), 0))
			_ = _index30
			(_254_v151).ArraySet1(_197_v116, (_index30).Int())
			var _255_v152 *C7
			_ = _255_v152
			var _nw46 *C7 = New_C7_()
			_ = _nw46
			_nw46.Ctor__(_191_v112)
			_255_v152 = _nw46
		}
		var _256_v153 *C4
		_ = _256_v153
		var _nw47 *C4 = New_C4_()
		_ = _nw47
		_nw47.Ctor__((_146_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_146_v73), 0))).Int()).(bool), _182_v102, _191_v112)
		_256_v153 = _nw47
		var _257_v154 _dafny.Map
		_ = _257_v154
		_257_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_v153, _62_v0)
		_181_v101 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_181_v101, _181_v101), (Companion_Default___.SafeIndex((_257_v154).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_181_v101, _181_v101)).Cardinality()))).Uint32(), (_160_v86).F6()), _181_v101)
		var _258_v155 _dafny.Sequence
		_ = _258_v155
		_258_v155 = _dafny.SeqOf(_181_v101)
		var _259_v156 _dafny.Map
		_ = _259_v156
		_259_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v117, Companion_Default___.Fm17(_258_v155, Companion_Default___.Fm9((_188_v109).F13(), _67_globalState), (_197_v116).F12(), (_197_v116).F12(), _67_globalState))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_208_v125), 0))
		_ = _index31
		(_208_v125).ArraySet1((func() _dafny.Int {
			if (_259_v156).Contains((_188_v109).F13()) {
				return (_259_v156).Get((_188_v109).F13()).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-30)
		})(), (_index31).Int())
		var _260_v157 _dafny.Int
		_ = _260_v157
		var _261_v158 bool
		_ = _261_v158
		var _out15 _dafny.Int
		_ = _out15
		var _out16 bool
		_ = _out16
		_out15, _out16 = (_187_v108).M2((_197_v116).F11(), _67_globalState)
		_260_v157 = _out15
		_261_v158 = _out16
	}
	_dafny.Print(_62_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_63_v1, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_64_v2), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_65_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_66_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_67_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_67_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_67_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_68_v5).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v73).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v73).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v74).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_148_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_v85)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v86).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_i6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v97).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v98).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v99).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v100).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_181_v101, _dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('y'), _dafny.CodePoint('y'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v102).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(124), _dafny.IntOfInt64(124), _dafny.One, _dafny.IntOfInt64(652))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_184_v103, _dafny.SeqOf(_dafny.IntOfInt64(4), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v106).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v107).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(124))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_187_v108).F4()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(124))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_187_v108.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v109).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_189_v110).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v111).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false).UpdateUnsafe(_dafny.IntOfInt64(147), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true).UpdateUnsafe(_dafny.IntOfInt64(2), false).UpdateUnsafe(_dafny.IntOfInt64(449), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(248), false).UpdateUnsafe(_dafny.IntOfInt64(125), false).UpdateUnsafe(_dafny.IntOfInt64(776), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(204), false).UpdateUnsafe(_dafny.IntOfInt64(205), false).UpdateUnsafe(_dafny.IntOfInt64(206), false).UpdateUnsafe(_dafny.IntOfInt64(207), false).UpdateUnsafe(_dafny.IntOfInt64(208), false).UpdateUnsafe(_dafny.IntOfInt64(209), false).UpdateUnsafe(_dafny.IntOfInt64(210), false).UpdateUnsafe(_dafny.IntOfInt64(211), false).UpdateUnsafe(_dafny.IntOfInt64(212), false).UpdateUnsafe(_dafny.IntOfInt64(213), false).UpdateUnsafe(_dafny.IntOfInt64(214), false).UpdateUnsafe(_dafny.IntOfInt64(215), false).UpdateUnsafe(_dafny.IntOfInt64(216), false).UpdateUnsafe(_dafny.IntOfInt64(217), false).UpdateUnsafe(_dafny.IntOfInt64(218), false).UpdateUnsafe(_dafny.IntOfInt64(219), false).UpdateUnsafe(_dafny.IntOfInt64(220), false).UpdateUnsafe(_dafny.IntOfInt64(221), false).UpdateUnsafe(_dafny.IntOfInt64(222), false).UpdateUnsafe(_dafny.IntOfInt64(223), false).UpdateUnsafe(_dafny.IntOfInt64(224), false).UpdateUnsafe(_dafny.IntOfInt64(225), false).UpdateUnsafe(_dafny.IntOfInt64(226), false).UpdateUnsafe(_dafny.IntOfInt64(227), false).UpdateUnsafe(_dafny.IntOfInt64(228), false).UpdateUnsafe(_dafny.IntOfInt64(229), false).UpdateUnsafe(_dafny.IntOfInt64(230), false).UpdateUnsafe(_dafny.IntOfInt64(231), false).UpdateUnsafe(_dafny.IntOfInt64(232), false).UpdateUnsafe(_dafny.IntOfInt64(233), false).UpdateUnsafe(_dafny.IntOfInt64(234), false).UpdateUnsafe(_dafny.IntOfInt64(235), false).UpdateUnsafe(_dafny.IntOfInt64(236), false).UpdateUnsafe(_dafny.IntOfInt64(237), false).UpdateUnsafe(_dafny.IntOfInt64(238), false).UpdateUnsafe(_dafny.IntOfInt64(239), false).UpdateUnsafe(_dafny.IntOfInt64(240), false).UpdateUnsafe(_dafny.IntOfInt64(241), false).UpdateUnsafe(_dafny.IntOfInt64(242), false).UpdateUnsafe(_dafny.IntOfInt64(243), false).UpdateUnsafe(_dafny.IntOfInt64(244), false).UpdateUnsafe(_dafny.IntOfInt64(245), false).UpdateUnsafe(_dafny.IntOfInt64(246), false).UpdateUnsafe(_dafny.IntOfInt64(247), false).UpdateUnsafe(_dafny.IntOfInt64(248), false).UpdateUnsafe(_dafny.IntOfInt64(249), false).UpdateUnsafe(_dafny.IntOfInt64(250), false).UpdateUnsafe(_dafny.IntOfInt64(251), false).UpdateUnsafe(_dafny.IntOfInt64(252), false).UpdateUnsafe(_dafny.IntOfInt64(253), false).UpdateUnsafe(_dafny.IntOfInt64(254), false).UpdateUnsafe(_dafny.IntOfInt64(255), false).UpdateUnsafe(_dafny.IntOfInt64(256), false).UpdateUnsafe(_dafny.IntOfInt64(257), false).UpdateUnsafe(_dafny.IntOfInt64(258), false).UpdateUnsafe(_dafny.IntOfInt64(259), false).UpdateUnsafe(_dafny.IntOfInt64(260), false).UpdateUnsafe(_dafny.IntOfInt64(261), false).UpdateUnsafe(_dafny.IntOfInt64(262), false).UpdateUnsafe(_dafny.IntOfInt64(263), false).UpdateUnsafe(_dafny.IntOfInt64(264), false).UpdateUnsafe(_dafny.IntOfInt64(265), false).UpdateUnsafe(_dafny.IntOfInt64(266), false).UpdateUnsafe(_dafny.IntOfInt64(267), false).UpdateUnsafe(_dafny.IntOfInt64(268), false).UpdateUnsafe(_dafny.IntOfInt64(269), false).UpdateUnsafe(_dafny.IntOfInt64(270), false).UpdateUnsafe(_dafny.IntOfInt64(271), false).UpdateUnsafe(_dafny.IntOfInt64(272), false).UpdateUnsafe(_dafny.IntOfInt64(273), false).UpdateUnsafe(_dafny.IntOfInt64(274), false).UpdateUnsafe(_dafny.IntOfInt64(275), false).UpdateUnsafe(_dafny.IntOfInt64(276), false).UpdateUnsafe(_dafny.IntOfInt64(277), false).UpdateUnsafe(_dafny.IntOfInt64(278), false).UpdateUnsafe(_dafny.IntOfInt64(279), false).UpdateUnsafe(_dafny.IntOfInt64(280), false).UpdateUnsafe(_dafny.IntOfInt64(281), false).UpdateUnsafe(_dafny.IntOfInt64(282), false).UpdateUnsafe(_dafny.IntOfInt64(283), false).UpdateUnsafe(_dafny.IntOfInt64(284), false).UpdateUnsafe(_dafny.IntOfInt64(285), false).UpdateUnsafe(_dafny.IntOfInt64(286), false).UpdateUnsafe(_dafny.IntOfInt64(287), false).UpdateUnsafe(_dafny.IntOfInt64(288), false).UpdateUnsafe(_dafny.IntOfInt64(289), false).UpdateUnsafe(_dafny.IntOfInt64(290), false).UpdateUnsafe(_dafny.IntOfInt64(291), false).UpdateUnsafe(_dafny.IntOfInt64(292), false).UpdateUnsafe(_dafny.IntOfInt64(293), false).UpdateUnsafe(_dafny.IntOfInt64(294), false).UpdateUnsafe(_dafny.IntOfInt64(295), false).UpdateUnsafe(_dafny.IntOfInt64(296), false).UpdateUnsafe(_dafny.IntOfInt64(297), false).UpdateUnsafe(_dafny.IntOfInt64(298), false).UpdateUnsafe(_dafny.IntOfInt64(299), false).UpdateUnsafe(_dafny.IntOfInt64(300), false).UpdateUnsafe(_dafny.IntOfInt64(301), false).UpdateUnsafe(_dafny.IntOfInt64(302), false).UpdateUnsafe(_dafny.IntOfInt64(303), false).UpdateUnsafe(_dafny.IntOfInt64(304), false).UpdateUnsafe(_dafny.IntOfInt64(305), false).UpdateUnsafe(_dafny.IntOfInt64(306), false).UpdateUnsafe(_dafny.IntOfInt64(307), false).UpdateUnsafe(_dafny.IntOfInt64(308), false).UpdateUnsafe(_dafny.IntOfInt64(309), false).UpdateUnsafe(_dafny.IntOfInt64(310), false).UpdateUnsafe(_dafny.IntOfInt64(311), false).UpdateUnsafe(_dafny.IntOfInt64(312), false).UpdateUnsafe(_dafny.IntOfInt64(313), false).UpdateUnsafe(_dafny.IntOfInt64(314), false).UpdateUnsafe(_dafny.IntOfInt64(315), false).UpdateUnsafe(_dafny.IntOfInt64(316), false).UpdateUnsafe(_dafny.IntOfInt64(317), false).UpdateUnsafe(_dafny.IntOfInt64(318), false).UpdateUnsafe(_dafny.IntOfInt64(319), false).UpdateUnsafe(_dafny.IntOfInt64(320), false).UpdateUnsafe(_dafny.IntOfInt64(321), false).UpdateUnsafe(_dafny.IntOfInt64(322), false).UpdateUnsafe(_dafny.IntOfInt64(323), false).UpdateUnsafe(_dafny.IntOfInt64(324), false).UpdateUnsafe(_dafny.IntOfInt64(325), false).UpdateUnsafe(_dafny.IntOfInt64(326), false).UpdateUnsafe(_dafny.IntOfInt64(327), false).UpdateUnsafe(_dafny.IntOfInt64(328), false).UpdateUnsafe(_dafny.IntOfInt64(329), false).UpdateUnsafe(_dafny.IntOfInt64(330), false).UpdateUnsafe(_dafny.IntOfInt64(331), false).UpdateUnsafe(_dafny.IntOfInt64(332), false).UpdateUnsafe(_dafny.IntOfInt64(333), false).UpdateUnsafe(_dafny.IntOfInt64(334), false).UpdateUnsafe(_dafny.IntOfInt64(335), false).UpdateUnsafe(_dafny.IntOfInt64(336), false).UpdateUnsafe(_dafny.IntOfInt64(337), false).UpdateUnsafe(_dafny.IntOfInt64(338), false).UpdateUnsafe(_dafny.IntOfInt64(339), false).UpdateUnsafe(_dafny.IntOfInt64(340), false).UpdateUnsafe(_dafny.IntOfInt64(341), false).UpdateUnsafe(_dafny.IntOfInt64(342), false).UpdateUnsafe(_dafny.IntOfInt64(343), false).UpdateUnsafe(_dafny.IntOfInt64(344), false).UpdateUnsafe(_dafny.IntOfInt64(345), false).UpdateUnsafe(_dafny.IntOfInt64(346), false).UpdateUnsafe(_dafny.IntOfInt64(347), false).UpdateUnsafe(_dafny.IntOfInt64(348), false).UpdateUnsafe(_dafny.IntOfInt64(349), false).UpdateUnsafe(_dafny.IntOfInt64(350), false).UpdateUnsafe(_dafny.IntOfInt64(351), false).UpdateUnsafe(_dafny.IntOfInt64(352), false).UpdateUnsafe(_dafny.IntOfInt64(353), false).UpdateUnsafe(_dafny.IntOfInt64(354), false).UpdateUnsafe(_dafny.IntOfInt64(355), false).UpdateUnsafe(_dafny.IntOfInt64(356), false).UpdateUnsafe(_dafny.IntOfInt64(357), false).UpdateUnsafe(_dafny.IntOfInt64(358), false).UpdateUnsafe(_dafny.IntOfInt64(359), false).UpdateUnsafe(_dafny.IntOfInt64(360), false).UpdateUnsafe(_dafny.IntOfInt64(361), false).UpdateUnsafe(_dafny.IntOfInt64(362), false).UpdateUnsafe(_dafny.IntOfInt64(363), false).UpdateUnsafe(_dafny.IntOfInt64(364), false).UpdateUnsafe(_dafny.IntOfInt64(365), false).UpdateUnsafe(_dafny.IntOfInt64(366), false).UpdateUnsafe(_dafny.IntOfInt64(367), false).UpdateUnsafe(_dafny.IntOfInt64(368), false).UpdateUnsafe(_dafny.IntOfInt64(369), false).UpdateUnsafe(_dafny.IntOfInt64(370), false).UpdateUnsafe(_dafny.IntOfInt64(371), false).UpdateUnsafe(_dafny.IntOfInt64(372), false).UpdateUnsafe(_dafny.IntOfInt64(373), false).UpdateUnsafe(_dafny.IntOfInt64(374), false).UpdateUnsafe(_dafny.IntOfInt64(375), false).UpdateUnsafe(_dafny.IntOfInt64(376), false).UpdateUnsafe(_dafny.IntOfInt64(377), false).UpdateUnsafe(_dafny.IntOfInt64(378), false).UpdateUnsafe(_dafny.IntOfInt64(379), false).UpdateUnsafe(_dafny.IntOfInt64(380), false).UpdateUnsafe(_dafny.IntOfInt64(381), false).UpdateUnsafe(_dafny.IntOfInt64(382), false).UpdateUnsafe(_dafny.IntOfInt64(383), false).UpdateUnsafe(_dafny.IntOfInt64(384), false).UpdateUnsafe(_dafny.IntOfInt64(385), false).UpdateUnsafe(_dafny.IntOfInt64(386), false).UpdateUnsafe(_dafny.IntOfInt64(387), false).UpdateUnsafe(_dafny.IntOfInt64(388), false).UpdateUnsafe(_dafny.IntOfInt64(389), false).UpdateUnsafe(_dafny.IntOfInt64(390), false).UpdateUnsafe(_dafny.IntOfInt64(391), false).UpdateUnsafe(_dafny.IntOfInt64(392), false).UpdateUnsafe(_dafny.IntOfInt64(393), false).UpdateUnsafe(_dafny.IntOfInt64(394), false).UpdateUnsafe(_dafny.IntOfInt64(395), false).UpdateUnsafe(_dafny.IntOfInt64(396), false).UpdateUnsafe(_dafny.IntOfInt64(397), false).UpdateUnsafe(_dafny.IntOfInt64(398), false).UpdateUnsafe(_dafny.IntOfInt64(399), false).UpdateUnsafe(_dafny.IntOfInt64(400), false).UpdateUnsafe(_dafny.IntOfInt64(401), false).UpdateUnsafe(_dafny.IntOfInt64(402), false).UpdateUnsafe(_dafny.IntOfInt64(403), false).UpdateUnsafe(_dafny.IntOfInt64(404), false).UpdateUnsafe(_dafny.IntOfInt64(405), false).UpdateUnsafe(_dafny.IntOfInt64(406), false).UpdateUnsafe(_dafny.IntOfInt64(407), false).UpdateUnsafe(_dafny.IntOfInt64(408), false).UpdateUnsafe(_dafny.IntOfInt64(409), false).UpdateUnsafe(_dafny.IntOfInt64(410), false).UpdateUnsafe(_dafny.IntOfInt64(411), false).UpdateUnsafe(_dafny.IntOfInt64(412), false).UpdateUnsafe(_dafny.IntOfInt64(413), false).UpdateUnsafe(_dafny.IntOfInt64(414), false).UpdateUnsafe(_dafny.IntOfInt64(415), false).UpdateUnsafe(_dafny.IntOfInt64(416), false).UpdateUnsafe(_dafny.IntOfInt64(417), false).UpdateUnsafe(_dafny.IntOfInt64(418), false).UpdateUnsafe(_dafny.IntOfInt64(419), false).UpdateUnsafe(_dafny.IntOfInt64(420), false).UpdateUnsafe(_dafny.IntOfInt64(421), false).UpdateUnsafe(_dafny.IntOfInt64(422), false).UpdateUnsafe(_dafny.IntOfInt64(423), false).UpdateUnsafe(_dafny.IntOfInt64(424), false).UpdateUnsafe(_dafny.IntOfInt64(425), false).UpdateUnsafe(_dafny.IntOfInt64(426), false).UpdateUnsafe(_dafny.IntOfInt64(427), false).UpdateUnsafe(_dafny.IntOfInt64(428), false).UpdateUnsafe(_dafny.IntOfInt64(429), false).UpdateUnsafe(_dafny.IntOfInt64(430), false).UpdateUnsafe(_dafny.IntOfInt64(431), false).UpdateUnsafe(_dafny.IntOfInt64(432), false).UpdateUnsafe(_dafny.IntOfInt64(433), false).UpdateUnsafe(_dafny.IntOfInt64(434), false).UpdateUnsafe(_dafny.IntOfInt64(435), false).UpdateUnsafe(_dafny.IntOfInt64(436), false).UpdateUnsafe(_dafny.IntOfInt64(437), false).UpdateUnsafe(_dafny.IntOfInt64(438), false).UpdateUnsafe(_dafny.IntOfInt64(439), false).UpdateUnsafe(_dafny.IntOfInt64(440), false).UpdateUnsafe(_dafny.IntOfInt64(441), false).UpdateUnsafe(_dafny.IntOfInt64(442), false).UpdateUnsafe(_dafny.IntOfInt64(443), false).UpdateUnsafe(_dafny.IntOfInt64(444), false).UpdateUnsafe(_dafny.IntOfInt64(445), false).UpdateUnsafe(_dafny.IntOfInt64(446), false).UpdateUnsafe(_dafny.IntOfInt64(447), false).UpdateUnsafe(_dafny.IntOfInt64(448), false).UpdateUnsafe(_dafny.IntOfInt64(449), false).UpdateUnsafe(_dafny.IntOfInt64(450), false).UpdateUnsafe(_dafny.IntOfInt64(451), false).UpdateUnsafe(_dafny.IntOfInt64(452), false).UpdateUnsafe(_dafny.IntOfInt64(453), false).UpdateUnsafe(_dafny.IntOfInt64(454), false).UpdateUnsafe(_dafny.IntOfInt64(455), false).UpdateUnsafe(_dafny.IntOfInt64(456), false).UpdateUnsafe(_dafny.IntOfInt64(457), false).UpdateUnsafe(_dafny.IntOfInt64(458), false).UpdateUnsafe(_dafny.IntOfInt64(459), false).UpdateUnsafe(_dafny.IntOfInt64(460), false).UpdateUnsafe(_dafny.IntOfInt64(461), false).UpdateUnsafe(_dafny.IntOfInt64(462), false).UpdateUnsafe(_dafny.IntOfInt64(463), false).UpdateUnsafe(_dafny.IntOfInt64(464), false).UpdateUnsafe(_dafny.IntOfInt64(465), false).UpdateUnsafe(_dafny.IntOfInt64(466), false).UpdateUnsafe(_dafny.IntOfInt64(467), false).UpdateUnsafe(_dafny.IntOfInt64(468), false).UpdateUnsafe(_dafny.IntOfInt64(469), false).UpdateUnsafe(_dafny.IntOfInt64(470), false).UpdateUnsafe(_dafny.IntOfInt64(471), false).UpdateUnsafe(_dafny.IntOfInt64(472), false).UpdateUnsafe(_dafny.IntOfInt64(473), false).UpdateUnsafe(_dafny.IntOfInt64(474), false).UpdateUnsafe(_dafny.IntOfInt64(475), false).UpdateUnsafe(_dafny.IntOfInt64(476), false).UpdateUnsafe(_dafny.IntOfInt64(477), false).UpdateUnsafe(_dafny.IntOfInt64(478), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true).UpdateUnsafe(_dafny.IntOfInt64(2), false).UpdateUnsafe(_dafny.IntOfInt64(449), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-2), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_191_v112).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(124), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_196_v115).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v116).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v116).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_198_v117)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_199_v118)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_202_v119).Dtor_cf27()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_202_v119).Dtor_cf27()).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v120).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_204_v121, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(124), _dafny.IntOfInt64(124), _dafny.One, _dafny.IntOfInt64(652)), _dafny.MultiSetOf(_dafny.IntOfInt64(124), _dafny.IntOfInt64(124), _dafny.One, _dafny.IntOfInt64(652)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v122).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v123).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(4), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_207_v124).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('y')), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v125).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_209_v126).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_223_v134).Dtor_cf15()).F13())
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
	Cf2 _dafny.Int
	Cf3 bool
	Cf4 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 bool, Cf4 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Sequence
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Sequence) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf6 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf6 _dafny.Int) D2 {
	return D2{D2_DC4{Cf6}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
	Cf7  _dafny.Int
	Cf8  bool
	Cf9  bool
	Cf10 bool
	Cf11 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf7 _dafny.Int, Cf8 bool, Cf9 bool, Cf10 bool, Cf11 bool) D2 {
	return D2{D2_DC5{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf5 _dafny.Map
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf5 _dafny.Map) D2 {
	return D2{D2_DC3{Cf5}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero)
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC5).Cf11
}

func (_this D2) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D3_DC7 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.Int
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf13 _dafny.Int, Cf14 _dafny.Int) D3 {
	return D3{D3_DC7{Cf13, Cf14}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC6 struct {
	Cf12 _dafny.CodePoint
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf12 _dafny.CodePoint) D3 {
	return D3{D3_DC6{Cf12}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(_dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf14
}

func (_this D3) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D3_DC6).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf12 == data2.Cf12
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

type D4_DC9 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.CodePoint
	Cf18 _dafny.Sequence
	Cf19 _dafny.Int
	Cf20 bool
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf16 _dafny.Int, Cf17 _dafny.CodePoint, Cf18 _dafny.Sequence, Cf19 _dafny.Int, Cf20 bool) D4 {
	return D4{D4_DC9{Cf16, Cf17, Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC8 struct {
	Cf15 *C0
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf15 *C0) D4 {
	return D4{D4_DC8{Cf15}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf16
}

func (_this D4) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D4_DC9).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D4_DC9).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC9).Cf20
}

func (_this D4) Dtor_cf15() *C0 {
	return _this.Get_().(D4_DC8).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
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

type D5_DC10 struct {
	Cf21 _dafny.Set
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf21 _dafny.Set) D5 {
	return D5{D5_DC10{Cf21}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D5) Dtor_cf21() _dafny.Set {
	return _this.Get_().(D5_DC10).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf21) + ")"
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

type D6_DC12 struct {
	Cf23 bool
	Cf24 _dafny.CodePoint
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf23 bool, Cf24 _dafny.CodePoint) D6 {
	return D6{D6_DC12{Cf23, Cf24}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC11 struct {
	Cf22 _dafny.Sequence
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC11{Cf22}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC12_(false, _dafny.CodePoint('D'))
}

func (_this D6) Dtor_cf23() bool {
	return _this.Get_().(D6_DC12).Cf23
}

func (_this D6) Dtor_cf24() _dafny.CodePoint {
	return _this.Get_().(D6_DC12).Cf24
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC11).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + data.Cf22.VerbatimString(true) + ")"
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
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
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

type D7_DC14 struct {
	Cf26 _dafny.Int
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf26 _dafny.Int) D7 {
	return D7{D7_DC14{Cf26}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC13 struct {
	Cf25 _dafny.Set
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf25 _dafny.Set) D7 {
	return D7{D7_DC13{Cf25}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC14_(_dafny.Zero)
}

func (_this D7) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf26
}

func (_this D7) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D7_DC13).Cf25
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf28 _dafny.Int
	Cf29 bool
	Cf30 bool
	Cf31 _dafny.Int
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf28 _dafny.Int, Cf29 bool, Cf30 bool, Cf31 _dafny.Int) D8 {
	return D8{D8_DC16{Cf28, Cf29, Cf30, Cf31}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC15 struct {
	Cf27 _dafny.Array
}

func (D8_DC15) isD8() {}

func (CompanionStruct_D8_) Create_DC15_(Cf27 _dafny.Array) D8 {
	return D8{D8_DC15{Cf27}}
}

func (_this D8) Is_DC15() bool {
	_, ok := _this.Get_().(D8_DC15)
	return ok
}

type D8_DC17 struct {
	Cf32 D8
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf32 D8) D8 {
	return D8{D8_DC17{Cf32}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC16_(_dafny.Zero, false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D8_DC16).Cf28
}

func (_this D8) Dtor_cf29() bool {
	return _this.Get_().(D8_DC16).Cf29
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC16).Cf30
}

func (_this D8) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D8_DC16).Cf31
}

func (_this D8) Dtor_cf27() _dafny.Array {
	return _this.Get_().(D8_DC15).Cf27
}

func (_this D8) Dtor_cf32() D8 {
	return _this.Get_().(D8_DC17).Cf32
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC15:
		{
			return "D8.DC15" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf32) + ")"
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
			return ok && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D8_DC15:
		{
			data2, ok := other.Get_().(D8_DC15)
			return ok && data1.Cf27 == data2.Cf27
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D9_DC19 struct {
	Cf34 T0
	Cf35 _dafny.Int
	Cf36 _dafny.Int
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf34 T0, Cf35 _dafny.Int, Cf36 _dafny.Int) D9 {
	return D9{D9_DC19{Cf34, Cf35, Cf36}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC20 struct {
	Cf37 bool
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf37 bool) D9 {
	return D9{D9_DC20{Cf37}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC18 struct {
	Cf33 _dafny.Array
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf33 _dafny.Array) D9 {
	return D9{D9_DC18{Cf33}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC19_((T0)(nil), _dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf34() T0 {
	return _this.Get_().(D9_DC19).Cf34
}

func (_this D9) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf36
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC20).Cf37
}

func (_this D9) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D9_DC18).Cf33
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && _dafny.AreEqual(data1.Cf34, data2.Cf34) && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf37 == data2.Cf37
		}
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf33 == data2.Cf33
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

type D10_DC21 struct {
	Cf38 _dafny.Array
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf38 _dafny.Array) D10 {
	return D10{D10_DC21{Cf38}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf38() _dafny.Array {
	return _this.Get_().(D10_DC21).Cf38
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf38 == data2.Cf38
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

type D11_DC22 struct {
	Cf39 _dafny.Map
}

func (D11_DC22) isD11() {}

func (CompanionStruct_D11_) Create_DC22_(Cf39 _dafny.Map) D11 {
	return D11{D11_DC22{Cf39}}
}

func (_this D11) Is_DC22() bool {
	_, ok := _this.Get_().(D11_DC22)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D11) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D11_DC22).Cf39
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC22:
		{
			return "D11.DC22" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC22:
		{
			data2, ok := other.Get_().(D11_DC22)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D12_DC23 struct {
	Cf40 _dafny.Map
}

func (D12_DC23) isD12() {}

func (CompanionStruct_D12_) Create_DC23_(Cf40 _dafny.Map) D12 {
	return D12{D12_DC23{Cf40}}
}

func (_this D12) Is_DC23() bool {
	_, ok := _this.Get_().(D12_DC23)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D12) Dtor_cf40() _dafny.Map {
	return _this.Get_().(D12_DC23).Cf40
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC23:
		{
			return "D12.DC23" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC23:
		{
			data2, ok := other.Get_().(D12_DC23)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D13_DC25 struct {
	Cf42 _dafny.Int
	Cf43 _dafny.Int
	Cf44 bool
}

func (D13_DC25) isD13() {}

func (CompanionStruct_D13_) Create_DC25_(Cf42 _dafny.Int, Cf43 _dafny.Int, Cf44 bool) D13 {
	return D13{D13_DC25{Cf42, Cf43, Cf44}}
}

func (_this D13) Is_DC25() bool {
	_, ok := _this.Get_().(D13_DC25)
	return ok
}

type D13_DC24 struct {
	Cf41 *C2
}

func (D13_DC24) isD13() {}

func (CompanionStruct_D13_) Create_DC24_(Cf41 *C2) D13 {
	return D13{D13_DC24{Cf41}}
}

func (_this D13) Is_DC24() bool {
	_, ok := _this.Get_().(D13_DC24)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC25_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D13) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D13_DC25).Cf42
}

func (_this D13) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D13_DC25).Cf43
}

func (_this D13) Dtor_cf44() bool {
	return _this.Get_().(D13_DC25).Cf44
}

func (_this D13) Dtor_cf41() *C2 {
	return _this.Get_().(D13_DC24).Cf41
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC25:
		{
			return "D13.DC25" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D13_DC24:
		{
			return "D13.DC24" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC25:
		{
			data2, ok := other.Get_().(D13_DC25)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44 == data2.Cf44
		}
	case D13_DC24:
		{
			data2, ok := other.Get_().(D13_DC24)
			return ok && data1.Cf41 == data2.Cf41
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

type D14_DC27 struct {
	Cf46 _dafny.Int
	Cf47 bool
	Cf48 bool
	Cf49 bool
	Cf50 _dafny.Int
}

func (D14_DC27) isD14() {}

func (CompanionStruct_D14_) Create_DC27_(Cf46 _dafny.Int, Cf47 bool, Cf48 bool, Cf49 bool, Cf50 _dafny.Int) D14 {
	return D14{D14_DC27{Cf46, Cf47, Cf48, Cf49, Cf50}}
}

func (_this D14) Is_DC27() bool {
	_, ok := _this.Get_().(D14_DC27)
	return ok
}

type D14_DC28 struct {
	Cf51 bool
	Cf52 bool
	Cf53 bool
	Cf54 _dafny.Int
}

func (D14_DC28) isD14() {}

func (CompanionStruct_D14_) Create_DC28_(Cf51 bool, Cf52 bool, Cf53 bool, Cf54 _dafny.Int) D14 {
	return D14{D14_DC28{Cf51, Cf52, Cf53, Cf54}}
}

func (_this D14) Is_DC28() bool {
	_, ok := _this.Get_().(D14_DC28)
	return ok
}

type D14_DC26 struct {
	Cf45 _dafny.MultiSet
}

func (D14_DC26) isD14() {}

func (CompanionStruct_D14_) Create_DC26_(Cf45 _dafny.MultiSet) D14 {
	return D14{D14_DC26{Cf45}}
}

func (_this D14) Is_DC26() bool {
	_, ok := _this.Get_().(D14_DC26)
	return ok
}

type D14_DC29 struct {
	Cf55 D14
}

func (D14_DC29) isD14() {}

func (CompanionStruct_D14_) Create_DC29_(Cf55 D14) D14 {
	return D14{D14_DC29{Cf55}}
}

func (_this D14) Is_DC29() bool {
	_, ok := _this.Get_().(D14_DC29)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC27_(_dafny.Zero, false, false, false, _dafny.Zero)
}

func (_this D14) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D14_DC27).Cf46
}

func (_this D14) Dtor_cf47() bool {
	return _this.Get_().(D14_DC27).Cf47
}

func (_this D14) Dtor_cf48() bool {
	return _this.Get_().(D14_DC27).Cf48
}

func (_this D14) Dtor_cf49() bool {
	return _this.Get_().(D14_DC27).Cf49
}

func (_this D14) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D14_DC27).Cf50
}

func (_this D14) Dtor_cf51() bool {
	return _this.Get_().(D14_DC28).Cf51
}

func (_this D14) Dtor_cf52() bool {
	return _this.Get_().(D14_DC28).Cf52
}

func (_this D14) Dtor_cf53() bool {
	return _this.Get_().(D14_DC28).Cf53
}

func (_this D14) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D14_DC28).Cf54
}

func (_this D14) Dtor_cf45() _dafny.MultiSet {
	return _this.Get_().(D14_DC26).Cf45
}

func (_this D14) Dtor_cf55() D14 {
	return _this.Get_().(D14_DC29).Cf55
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC27:
		{
			return "D14.DC27" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D14_DC28:
		{
			return "D14.DC28" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D14_DC26:
		{
			return "D14.DC26" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D14_DC29:
		{
			return "D14.DC29" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC27:
		{
			data2, ok := other.Get_().(D14_DC27)
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49 && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D14_DC28:
		{
			data2, ok := other.Get_().(D14_DC28)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D14_DC26:
		{
			data2, ok := other.Get_().(D14_DC26)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D14_DC29:
		{
			data2, ok := other.Get_().(D14_DC29)
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

// Definition of trait T0
type T0 interface {
	String() string
	F3() _dafny.Array
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
	M1(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint)
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
	F2  _dafny.Int
	_f0 _dafny.Array
	_f1 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 bool, f2 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
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
	_f13 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f13 = _dafny.Zero
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

func (_this *C0) Ctor__(f13 _dafny.Int) {
	{
		(_this)._f13 = f13
	}
}
func (_this *C0) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f3  _dafny.Array
	_f14 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = _dafny.Zero
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

func (_this *C1) F3() _dafny.Array {
	return _this._f3
}
func (_this *C1) Ctor__(f14 _dafny.Int, f3 _dafny.Array) {
	{
		(_this)._f14 = f14
		(_this)._f3 = f3
	}
}
func (_this *C1) Fm29(p0 bool, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('y')
	}
}
func (_this *C1) F14() _dafny.Int {
	{
		return _this._f14
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f3  _dafny.Array
	_f11 bool
	_f12 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = false
	_this._f12 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C2{}
var _ T1 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F3() _dafny.Array {
	return _this._f3
}
func (_this *C2) Ctor__(f11 bool, f12 _dafny.Int, f3 _dafny.Array) {
	{
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f3 = f3
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var _262_v0 bool
		_ = _262_v0
		_262_v0 = true
		_262_v0 = (_this).F11()
		var _263_v1 _dafny.Map
		_ = _263_v1
		_263_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), Companion_Default___.Fm23(p2, (_this).F12(), (_this).F12(), true, globalState))
		var _264_v2 _dafny.Sequence
		_ = _264_v2
		_264_v2 = _dafny.SeqOf(_263_v1, _263_v1)
		var _265_v3 _dafny.Map
		_ = _265_v3
		_265_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm24(_dafny.SetOf((_this).F12(), (_this).F12(), _dafny.IntOfUint32((p2).Cardinality())), globalState), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int), _264_v2)
		var _266_v4 D1
		_ = _266_v4
		_266_v4 = Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Concatenate(_264_v2, (func() _dafny.Sequence {
			if (_265_v3).Contains((_this).F12()) {
				return (_265_v3).Get((_this).F12()).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_263_v1)
		})()))
		var _source5 D1 = _266_v4
		_ = _source5
		if _source5.Is_DC2() {
			var _267___mcc_h0 _dafny.Int = _source5.Get_().(D1_DC2).Cf2
			_ = _267___mcc_h0
			var _268___mcc_h1 bool = _source5.Get_().(D1_DC2).Cf3
			_ = _268___mcc_h1
			var _269___mcc_h2 _dafny.Int = _source5.Get_().(D1_DC2).Cf4
			_ = _269___mcc_h2
			var _270_cf4 _dafny.Int = _269___mcc_h2
			_ = _270_cf4
			var _271_cf3 bool = _268___mcc_h1
			_ = _271_cf3
			var _272_cf2 _dafny.Int = _267___mcc_h0
			_ = _272_cf2
			var _273_v5 _dafny.Array
			_ = _273_v5
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw48
			_273_v5 = _nw48
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_273_v5), 0))
			_ = _index32
			(_273_v5).ArraySet1(_262_v0, (_index32).Int())
			var _274_v6 _dafny.Set
			_ = _274_v6
			_274_v6 = _dafny.SetOf(_272_cf2, _dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfInt64(-72))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_273_v5), 0))
			_ = _index33
			(_273_v5).ArraySet1((_274_v6).IsSubsetOf((_274_v6).Difference(_274_v6)), (_index33).Int())
			_274_v6 = _274_v6
			if _271_cf3 {
				var _275_v7 _dafny.Sequence
				_ = _275_v7
				_275_v7 = _dafny.SeqOf(true, _262_v0)
				var _276_v8 *C0
				_ = _276_v8
				var _nw49 *C0 = New_C0_()
				_ = _nw49
				_nw49.Ctor__((_270_cf4).Minus(_dafny.IntOfUint32((_275_v7).Cardinality())))
				_276_v8 = _nw49
				var _277_v9 _dafny.Sequence
				_ = _277_v9
				_277_v9 = _dafny.UnicodeSeqOfUtf8Bytes("hdgtom")
				_277_v9 = (func() _dafny.Sequence {
					if !(false) {
						return _277_v9
					}
					return _277_v9
				})()
				var _278_v10 _dafny.Set
				_ = _278_v10
				_278_v10 = _dafny.SetOf((_this).F11(), p1)
				var _279_v11 _dafny.Array
				_ = _279_v11
				var _nwElement0_15 _dafny.Set = _dafny.SetOf(_262_v0)
				_ = _nwElement0_15
				var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(14))
				_ = _nw50
				(_nw50).ArraySet1(_nwElement0_15, 0)
				(_nw50).ArraySet1((_278_v10).Union(_278_v10), 1)
				(_nw50).ArraySet1(_278_v10, 2)
				(_nw50).ArraySet1(_dafny.SetOf((_273_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_273_v5), 0))).Int()).(bool), true), 3)
				(_nw50).ArraySet1(_278_v10, 4)
				(_nw50).ArraySet1(_278_v10, 5)
				(_nw50).ArraySet1(_278_v10, 6)
				(_nw50).ArraySet1(_278_v10, 7)
				(_nw50).ArraySet1((func() _dafny.Set {
					if true {
						return _278_v10
					}
					return _278_v10
				})(), 8)
				(_nw50).ArraySet1((_dafny.SetOf(_262_v0)).Difference(_278_v10), 9)
				(_nw50).ArraySet1(_278_v10, 10)
				(_nw50).ArraySet1(_278_v10, 11)
				(_nw50).ArraySet1(_278_v10, 12)
				(_nw50).ArraySet1(_278_v10, 13)
				_279_v11 = _nw50
				_279_v11 = _279_v11
				var _280_v12 _dafny.MultiSet
				_ = _280_v12
				_280_v12 = _dafny.MultiSetOf((_this).F12(), _dafny.IntOfInt64(634), (_276_v8).F13())
				var _281_v13 _dafny.Map
				_ = _281_v13
				_281_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v12, _270_cf4)
				var _282_v14 D2
				_ = _282_v14
				_282_v14 = Companion_D2_.Create_DC3_(_281_v13)
				var _283_v15 _dafny.Map
				_ = _283_v15
				_283_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_v14, p1)
				var _284_v16 _dafny.Map
				_ = _284_v16
				_284_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_cf2, ((_274_v6).Cardinality()).Cmp((_283_v15).Cardinality()) >= 0)
				_284_v16 = (_284_v16).Update((_276_v8).F13(), (_273_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_273_v5), 0))).Int()).(bool))
				var _285_v17 _dafny.Array
				_ = _285_v17
				var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw51
				_285_v17 = _nw51
				var _286_v18 _dafny.Sequence
				_ = _286_v18
				_286_v18 = _dafny.SeqOf(_274_v6)
				var _287_v19 _dafny.Map
				_ = _287_v19
				_287_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_285_v17, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_286_v18, _dafny.SeqOf(_274_v6, _274_v6, _274_v6))).Cardinality()))
				_287_v19 = (_287_v19).Update(_285_v17, (_this).F12())
			} else {
				_262_v0 = !(_271_cf3)
				var _288_v20 bool
				_ = _288_v20
				var _289_v21 bool
				_ = _289_v21
				var _290_v22 _dafny.Int
				_ = _290_v22
				var _out17 bool
				_ = _out17
				var _out18 bool
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				_out17, _out18, _out19 = Companion_Default___.M0(_270_cf4, globalState)
				_288_v20 = _out17
				_289_v21 = _out18
				_290_v22 = _out19
				var _291_v23 _dafny.Sequence
				_ = _291_v23
				_291_v23 = _dafny.SeqOf(_271_cf3, p1)
				var _292_v25 _dafny.Sequence
				_ = _292_v25
				_292_v25 = _dafny.SeqOf(func() _dafny.Set {
					var _coll15 = _dafny.NewBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate((p0).Elements()); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _293_v24 _dafny.Int
						_293_v24 = interface{}(_compr_15).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(p0, _293_v24) {
							_coll15.Add((_293_v24).Times(_dafny.IntOfInt64(495)))
						}
					}
					return _coll15.ToSet()
				}())
				var _294_v26 _dafny.Map
				_ = _294_v26
				_294_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.SeqOf((_291_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.IntOfUint32((_291_v23).Cardinality()))).Uint32()).(bool), _271_cf3)).Cardinality())).Times(Companion_Default___.Fm24((_292_v25).Select((Companion_Default___.SafeIndex((p0).Select((Companion_Default___.SafeIndex(_290_v22, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_292_v25).Cardinality()))).Uint32()).(_dafny.Set), globalState)), _262_v0)
				_294_v26 = (_294_v26).Update((_this).F12(), !(!(_288_v20)) || (_262_v0))
				var _295_v27 _dafny.Array
				_ = _295_v27
				var _nwElement0_16 _dafny.Int = (_this).F12()
				_ = _nwElement0_16
				var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.One)
				_ = _nw52
				(_nw52).ArraySet1(_nwElement0_16, 0)
				_295_v27 = _nw52
				var _296_v28 _dafny.Map
				_ = _296_v28
				_296_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v0, _274_v6)
				var _297_v29 _dafny.MultiSet
				_ = _297_v29
				_297_v29 = _dafny.MultiSetOf(!(false), _271_cf3)
				var _298_v30 _dafny.Map
				_ = _298_v30
				_298_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (func() _dafny.Int {
					if (_297_v29).Contains(_288_v20) {
						return (_297_v29).Multiplicity(_288_v20)
					}
					return _dafny.IntOfInt64(293)
				})())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_295_v27), 0))
				_ = _index34
				(_295_v27).ArraySet1(Companion_Default___.SafeModuloInt(((func() _dafny.Set {
					if (_296_v28).Contains((func() bool {
						if (_263_v1).Contains(p1) {
							return (_263_v1).Get(p1).(bool)
						}
						return false
					})()) {
						return (_296_v28).Get((func() bool {
							if (_263_v1).Contains(p1) {
								return (_263_v1).Get(p1).(bool)
							}
							return false
						})()).(_dafny.Set)
					}
					return (_292_v25).Select((Companion_Default___.SafeIndex((_298_v30).Cardinality(), _dafny.IntOfUint32((_292_v25).Cardinality()))).Uint32()).(_dafny.Set)
				})()).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.Fm24(_274_v6, globalState))), (_index34).Int())
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_295_v27), 0))
				_ = _index35
				(_295_v27).ArraySet1((_290_v22).Minus(_272_cf2), (_index35).Int())
				var _299_v31 *C0
				_ = _299_v31
				var _nw53 *C0 = New_C0_()
				_ = _nw53
				_nw53.Ctor__((_dafny.Zero).Minus(_dafny.IntOfInt64(-817)))
				_299_v31 = _nw53
			}
			(globalState).F2 = (_this).F12()
		} else {
			var _300___mcc_h3 _dafny.Sequence = _source5.Get_().(D1_DC1).Cf1
			_ = _300___mcc_h3
			var _301_cf1 _dafny.Sequence = _300___mcc_h3
			_ = _301_cf1
			var _302_v32 _dafny.Array
			_ = _302_v32
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_4
			var _nw54 _dafny.Array
			_ = _nw54
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw54 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = func(_303_i0 _dafny.Int) bool {
					return (_this).F11()
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw54 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw54).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw54).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_302_v32 = _nw54
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))
			_ = _index36
			(_302_v32).ArraySet1(_262_v0, (_index36).Int())
			var _304_v33 D2
			_ = _304_v33
			_304_v33 = Companion_D2_.Create_DC5_((_this).F12(), (_this).F11(), (_this).F11(), _262_v0, _262_v0)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))
			_ = _index37
			(_302_v32).ArraySet1((_304_v33).Dtor_cf11(), (_index37).Int())
			var _305_v34 _dafny.MultiSet
			_ = _305_v34
			_305_v34 = _dafny.MultiSetOf(_302_v32)
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))
			_ = _index38
			(_302_v32).ArraySet1(((_305_v34).Difference(_305_v34)).IsSubsetOf(_305_v34), (_index38).Int())
			var _306_v35 _dafny.Set
			_ = _306_v35
			_306_v35 = _dafny.SetOf((_this).F11(), _262_v0, (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool))
			var _307_v36 _dafny.Map
			_ = _307_v36
			_307_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v35, _306_v35)
			if (_306_v35).IsDisjointFrom((func() _dafny.Set {
				if (_307_v36).Contains(_dafny.SetOf(_262_v0, _262_v0)) {
					return (_307_v36).Get(_dafny.SetOf(_262_v0, _262_v0)).(_dafny.Set)
				}
				return _dafny.SetOf(_262_v0)
			})()) {
				var _308_v37 _dafny.Map
				_ = _308_v37
				_308_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), p0)
				var _309_v38 _dafny.Sequence
				_ = _309_v38
				_309_v38 = _dafny.SeqOf((_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), (_this).F11(), (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), false)
				var _310_v39 _dafny.Map
				_ = _310_v39
				_310_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), (_this).F12())
				var _311_v40 _dafny.Array
				_ = _311_v40
				var _nwElement0_17 _dafny.Sequence = p0
				_ = _nwElement0_17
				var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(6))
				_ = _nw55
				(_nw55).ArraySet1(_nwElement0_17, 0)
				(_nw55).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_308_v37).Contains(!(Companion_Default___.Fm23(p2, (_this).F12(), _dafny.IntOfInt64(784), (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), globalState))) {
						return (_308_v37).Get(!(Companion_Default___.Fm23(p2, (_this).F12(), _dafny.IntOfInt64(784), (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), globalState))).(_dafny.Sequence)
					}
					return p0
				})(), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_309_v38).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_310_v39).Cardinality())), 1)
				(_nw55).ArraySet1(_dafny.SeqOf((_this).F12(), (_this).F12()), 2)
				(_nw55).ArraySet1(p0, 3)
				(_nw55).ArraySet1(p0, 4)
				(_nw55).ArraySet1(p0, 5)
				_311_v40 = _nw55
				var _312_v41 _dafny.Map
				_ = _312_v41
				_312_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _311_v40)
				var _313_v42 _dafny.Sequence
				_ = _313_v42
				_313_v42 = _dafny.SeqOf((func() _dafny.Array {
					if (_312_v41).Contains((_this).F12()) {
						return (_312_v41).Get((_this).F12()).(_dafny.Array)
					}
					return _311_v40
				})())
				_311_v40 = (_313_v42).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_313_v42).Cardinality()))).Uint32()).(_dafny.Array)
				var _314_v43 _dafny.CodePoint
				_ = _314_v43
				_314_v43 = _dafny.CodePoint('l')
				r1 = _314_v43
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))
				_ = _index39
				(_302_v32).ArraySet1((!(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2), (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))).Uint32(), p2), p2))) && (!(Companion_Default___.Fm23(p2, (_this).F12(), (_this).F12(), (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), globalState))), (_index39).Int())
				var _315_v44 _dafny.Array
				_ = _315_v44
				var _nw56 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
				_ = _nw56
				_315_v44 = _nw56
				var _316_v45 *C0
				_ = _316_v45
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__((_this).F12())
				_316_v45 = _nw57
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_315_v44), 0))
				_ = _index40
				(_315_v44).ArraySet1(_316_v45, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_315_v44), 0))
				_ = _index41
				(_315_v44).ArraySet1(_316_v45, (_index41).Int())
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_311_v40), 0))
				_ = _index42
				(_311_v40).ArraySet1(p0, (_index42).Int())
				var _317_v46 _dafny.Sequence
				_ = _317_v46
				_317_v46 = _dafny.SeqOf(_309_v38)
				var _318_v47 _dafny.MultiSet
				_ = _318_v47
				_318_v47 = _dafny.MultiSetOf(_309_v38, (_317_v46).Select((Companion_Default___.SafeIndex((_316_v45).F13(), _dafny.IntOfUint32((_317_v46).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _319_v48 _dafny.MultiSet
				_ = _319_v48
				_319_v48 = _dafny.MultiSetOf(_262_v0)
				var _320_v49 _dafny.Map
				_ = _320_v49
				_320_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v47, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((p0).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_319_v48).Contains((_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool)) {
						return (_319_v48).Multiplicity((_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool))
					}
					return _dafny.IntOfUint32((p2).Cardinality())
				})(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_this).F12()), p0))
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_311_v40), 0))
				_ = _index43
				(_311_v40).ArraySet1((func() _dafny.Sequence {
					if (_320_v49).Contains(_318_v47) {
						return (_320_v49).Get(_318_v47).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg24 _dafny.Int) interface{} {
							return coer24(arg24)
						}
					}(func(_321_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-644)
					}))
				})(), (_index43).Int())
			} else {
				var _322_v50 _dafny.Set
				_ = _322_v50
				_322_v50 = _dafny.SetOf((_this).F12())
				var _323_v51 _dafny.Sequence
				_ = _323_v51
				_323_v51 = _dafny.SeqOf(true, p1)
				var _324_v52 _dafny.MultiSet
				_ = _324_v52
				_324_v52 = _dafny.MultiSetOf(p1, p1)
				var _325_v53 _dafny.Map
				_ = _325_v53
				_325_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F12())
				var _326_v54 _dafny.Array
				_ = _326_v54
				var _nwElement0_18 _dafny.Int = _dafny.IntOfInt64(413)
				_ = _nwElement0_18
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(24))
				_ = _nw58
				(_nw58).ArraySet1(_nwElement0_18, 0)
				(_nw58).ArraySet1((_this).F12(), 1)
				(_nw58).ArraySet1((_this).F12(), 2)
				(_nw58).ArraySet1((_this).F12(), 3)
				(_nw58).ArraySet1((_this).F12(), 4)
				(_nw58).ArraySet1((_this).F12(), 5)
				(_nw58).ArraySet1((_dafny.Zero).Minus((_this).F12()), 6)
				(_nw58).ArraySet1(Companion_Default___.Fm24(_322_v50, globalState), 7)
				(_nw58).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_323_v51).Cardinality())), 8)
				(_nw58).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 9)
				(_nw58).ArraySet1((_dafny.Zero).Minus((_this).F12()), 10)
				(_nw58).ArraySet1((_dafny.Zero).Minus((_this).F12()), 11)
				(_nw58).ArraySet1((_this).F12(), 12)
				(_nw58).ArraySet1((_this).F12(), 13)
				(_nw58).ArraySet1(_dafny.IntOfInt64(825), 14)
				(_nw58).ArraySet1((_this).F12(), 15)
				(_nw58).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p2).Cardinality())), 16)
				(_nw58).ArraySet1((_this).F12(), 17)
				(_nw58).ArraySet1((_this).F12(), 18)
				(_nw58).ArraySet1((func() _dafny.Int {
					if (_324_v52).Contains((_this).F11()) {
						return (_324_v52).Multiplicity((_this).F11())
					}
					return (_325_v53).Cardinality()
				})(), 19)
				(_nw58).ArraySet1((_this).F12(), 20)
				(_nw58).ArraySet1((_324_v52).Cardinality(), 21)
				(_nw58).ArraySet1((_this).F12(), 22)
				(_nw58).ArraySet1(_dafny.IntOfInt64(241), 23)
				_326_v54 = _nw58
				var _327_v55 _dafny.Sequence
				_ = _327_v55
				_327_v55 = _dafny.SeqOf(_326_v54, _326_v54)
				var _328_v56 *C0
				_ = _328_v56
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__((_this).F12())
				_328_v56 = _nw59
				var _rhs29 _dafny.Sequence = _dafny.SeqOf(_326_v54, _326_v54, _326_v54)
				_ = _rhs29
				var _rhs30 *C0 = _328_v56
				_ = _rhs30
				_327_v55 = _rhs29
				_328_v56 = _rhs30
				var _329_v57 _dafny.Sequence
				_ = _329_v57
				_329_v57 = _dafny.SeqOf(_322_v50, _dafny.SetOf((_this).F12()), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}((func(_330_p2 _dafny.Sequence, _331_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_332_i2 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_p2, _331_v0)).Cardinality()
					}
				})(p2, _262_v0)))).Cardinality())))
				var _333_v58 _dafny.MultiSet
				_ = _333_v58
				_333_v58 = _dafny.MultiSetOf(_dafny.IntOfInt64(848), Companion_Default___.Fm24((_329_v57).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_329_v57).Cardinality()))).Uint32()).(_dafny.Set), globalState), (_dafny.SetOf(_328_v56, _328_v56, _328_v56)).Cardinality())
				var _334_v59 _dafny.Map
				_ = _334_v59
				_334_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v1).Cardinality(), (_this).F12())
				(globalState).F2 = (func() _dafny.Int {
					if (_333_v58).Contains(_dafny.IntOfInt64(725)) {
						return (_333_v58).Multiplicity(_dafny.IntOfInt64(725))
					}
					return _dafny.IntOfUint32(((func() _dafny.Sequence {
						if Companion_Default___.Fm23(p2, (_328_v56).F13(), (_334_v59).Cardinality(), (_302_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_302_v32), 0))).Int()).(bool), globalState) {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg26 _dafny.Int) interface{} {
									return coer26(arg26)
								}
							}(func(_335_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('y')
							}))
						}
						return p2
					})()).Cardinality())
				})()
				var _336_v60 _dafny.Sequence
				_ = _336_v60
				_336_v60 = _dafny.SeqOf(_323_v51)
				_262_v0 = _dafny.Companion_Sequence_.Equal(_336_v60, _336_v60)
				(globalState).F2 = ((func() D2 {
					if !(_262_v0) {
						return _304_v33
					}
					return _304_v33
				})()).Dtor_cf7()
				r0 = ((_this).F12()).Minus((_328_v56).F13())
			}
			var _337_v61 *C0
			_ = _337_v61
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__((_dafny.IntOfInt64(733)).Times((_this).F12()))
			_337_v61 = _nw60
		}
		var _338_v62 _dafny.Map
		_ = _338_v62
		_338_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		_338_v62 = (_338_v62).Update(p0, p2)
		var _339_v63 _dafny.Array
		_ = _339_v63
		var _len0_5 _dafny.Int = _dafny.One
		_ = _len0_5
		var _nw61 _dafny.Array
		_ = _nw61
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw61 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_340_i5 _dafny.Int) _dafny.Int {
				return (_340_i5).Minus(_dafny.IntOfInt64(-892))
			}
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw61 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw61).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw61).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_339_v63 = _nw61
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_339_v63), 0))); ; {
			_guard_loop_0, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _341_i4 _dafny.Int
			_341_i4 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_341_i4).Sign() != -1) && ((_341_i4).Cmp(_dafny.ArrayLen((_339_v63), 0)) < 0)) {
				(_339_v63).ArraySet1(Companion_Default___.SafeModuloInt(_341_i4, (_this).F12()), (_341_i4).Int())
			}
		}
		r0 = (_this).F12()
		var _342_v64 _dafny.Set
		_ = _342_v64
		_342_v64 = _dafny.SetOf(_dafny.CodePoint('y'))
		var _hi1 _dafny.Int = (_342_v64).Cardinality()
		_ = _hi1
		for _343_i6 := (_dafny.Zero).Minus((_this).F12()); _343_i6.Cmp(_hi1) < 0; _343_i6 = _343_i6.Plus(_dafny.One) {
			var _344_v65 D3
			_ = _344_v65
			_344_v65 = Companion_D3_.Create_DC7_((_343_i6).Plus(_343_i6), (_this).F12())
			var _source6 D3 = _344_v65
			_ = _source6
			if _source6.Is_DC7() {
				var _345___mcc_h4 _dafny.Int = _source6.Get_().(D3_DC7).Cf13
				_ = _345___mcc_h4
				var _346___mcc_h5 _dafny.Int = _source6.Get_().(D3_DC7).Cf14
				_ = _346___mcc_h5
				var _347_cf14 _dafny.Int = _346___mcc_h5
				_ = _347_cf14
				var _348_cf13 _dafny.Int = _345___mcc_h4
				_ = _348_cf13
				var _349_v66 _dafny.Map
				_ = _349_v66
				_349_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_cf13, !((_this).F11()))
				(globalState).F2 = (_348_cf13).Plus((_349_v66).Cardinality())
				_262_v0 = !(_dafny.Companion_Sequence_.IsPrefixOf(p2, p2))
				var _350_v67 *C0
				_ = _350_v67
				var _nw62 *C0 = New_C0_()
				_ = _nw62
				_nw62.Ctor__(_347_cf14)
				_350_v67 = _nw62
				_262_v0 = true
			} else {
				var _351___mcc_h6 _dafny.CodePoint = _source6.Get_().(D3_DC6).Cf12
				_ = _351___mcc_h6
				var _352_cf12 _dafny.CodePoint = _351___mcc_h6
				_ = _352_cf12
				_262_v0 = ((_343_i6).Minus(_343_i6)).Cmp(_dafny.IntOfUint32((p2).Cardinality())) > 0
				var _353_v68 _dafny.Sequence
				_ = _353_v68
				_353_v68 = _dafny.UnicodeSeqOfUtf8Bytes("chue")
				var _354_v69 _dafny.Array
				_ = _354_v69
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw63
				_354_v69 = _nw63
				var _355_v70 _dafny.Set
				_ = _355_v70
				_355_v70 = _dafny.SetOf(_354_v69, _354_v69)
				var _356_v71 _dafny.MultiSet
				_ = _356_v71
				_356_v71 = _dafny.MultiSetOf((func() _dafny.Int {
					if true {
						return _343_i6
					}
					return (_355_v70).Cardinality()
				})())
				var _rhs31 _dafny.Sequence = p2
				_ = _rhs31
				var _rhs32 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (true) || ((_this).F11()) {
						return _343_i6
					}
					return (_this).F12()
				})())
				_ = _rhs32
				var _rhs33 _dafny.MultiSet = (_356_v71).Difference((_dafny.MultiSetFromSeq(p0)).Intersection(_356_v71))
				_ = _rhs33
				_353_v68 = _rhs31
				r0 = _rhs32
				_356_v71 = _rhs33
				_354_v69 = _354_v69
				_266_v4 = _266_v4
			}
			var _357_v72 *C0
			_ = _357_v72
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__((_this).F12())
			_357_v72 = _nw64
			(globalState).F2 = _dafny.IntOfInt64(99)
			var _358_v73 _dafny.MultiSet
			_ = _358_v73
			_358_v73 = _dafny.MultiSetOf(p1, (_this).F11())
			var _359_v74 _dafny.Map
			_ = _359_v74
			_359_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v73, !(p1))
			var _360_v75 _dafny.Map
			_ = _360_v75
			_360_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v73, p1)).Merge(_359_v74), Companion_Default___.Fm25(p1, (_357_v72).F13(), globalState))
			_360_v75 = (_360_v75).Update((_359_v74).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v73, p1)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(147))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_361_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})))
		}
		r0 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F12(), (Companion_D2_.Create_DC4_((_this).F12())).Dtor_cf6()))
		var _362_v76 _dafny.CodePoint
		_ = _362_v76
		_362_v76 = _dafny.CodePoint('q')
		r1 = _362_v76
		return r0, r1
	}
}
func (_this *C2) M6(globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		r2 = ((_this).F12()).Cmp(((_this).F12()).Times((_this).F12())) != 0
		var _363_v0 _dafny.Map
		_ = _363_v0
		_363_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), true)
		var _364_v1 _dafny.Map
		_ = _364_v1
		_364_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _363_v0)
		var _365_v2 _dafny.Set
		_ = _365_v2
		_365_v2 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(760))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_366_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality()), (_this).F12(), (_364_v1).Cardinality())
		var _hi2 _dafny.Int = Companion_Default___.Fm24(_365_v2, globalState)
		_ = _hi2
		for _367_i0 := (_this).F12(); _367_i0.Cmp(_hi2) < 0; _367_i0 = _367_i0.Plus(_dafny.One) {
			r2 = (func() bool {
				if (_this).F11() {
					return (_this).F11()
				}
				return (_this).F11()
			})()
			(globalState).F2 = _367_i0
			var _368_v3 _dafny.Sequence
			_ = _368_v3
			_368_v3 = _dafny.UnicodeSeqOfUtf8Bytes("y")
			if !(Companion_Default___.Fm23(_368_v3, Companion_Default___.SafeDivisionInt(_367_i0, _dafny.IntOfInt64(-836)), _367_i0, (_this).F11(), globalState)) {
				_368_v3 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_368_v3, _368_v3), _368_v3), (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_368_v3, _368_v3), _368_v3)).Cardinality()))).Uint32(), _dafny.CodePoint('f'))
				var _369_v4 _dafny.Array
				_ = _369_v4
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw65
				_369_v4 = _nw65
				var _370_v5 D2
				_ = _370_v5
				_370_v5 = Companion_D2_.Create_DC5_(_367_i0, (_this).F11(), (_this).F11(), (_this).F11(), (_this).F11())
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))
				_ = _index44
				(_369_v4).ArraySet1(((_370_v5).Dtor_cf9()) || ((_this).F11()), (_index44).Int())
				var _371_v6 _dafny.Array
				_ = _371_v6
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_6
				var _nw66 _dafny.Array
				_ = _nw66
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw66 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) D2 = (func(_372_v5 D2) func(_dafny.Int) D2 {
						return func(_373_i2 _dafny.Int) D2 {
							return _372_v5
						}
					})(_370_v5)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw66 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw66).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw66).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_371_v6 = _nw66
				var _374_v7 _dafny.Sequence
				_ = _374_v7
				_374_v7 = _dafny.SeqOf(!((_this).F11()), (_this).F11(), true, (_this).F11(), !((_this).F11()))
				var _375_v8 _dafny.MultiSet
				_ = _375_v8
				_375_v8 = _dafny.MultiSetOf((_this).F11())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))
				_ = _index45
				var _rhs34 bool = ((_this).F12()).Cmp(_dafny.IntOfUint32((_374_v7).Cardinality())) != 0
				_ = _rhs34
				var _rhs35 bool = !((_375_v8).Update(false, Companion_Default___.Abs(_367_i0))).Equals((_375_v8).Difference(_375_v8))
				_ = _rhs35
				var _rhs36 bool = (_this).F11()
				_ = _rhs36
				var _rhs37 _dafny.Array = _371_v6
				_ = _rhs37
				var _lhs17 _dafny.Array = _369_v4
				_ = _lhs17
				var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))
				_ = _lhs18
				r2 = _rhs34
				(_lhs17).ArraySet1(_rhs35, (_lhs18).Int())
				r2 = _rhs36
				_371_v6 = _rhs37
				var _376_v9 _dafny.CodePoint
				_ = _376_v9
				_376_v9 = _dafny.CodePoint('j')
				var _377_v10 D1
				_ = _377_v10
				_377_v10 = Companion_D1_.Create_DC2_((_this).F12(), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool), _dafny.IntOfInt64(-552))
				var _378_v11 _dafny.Array
				_ = _378_v11
				var _nwElement0_19 _dafny.Sequence = _374_v7
				_ = _nwElement0_19
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(4))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_19, 0)
				(_nw67).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_374_v7, _374_v7), 1)
				(_nw67).ArraySet1(_374_v7, 2)
				(_nw67).ArraySet1((func() _dafny.Sequence {
					if (_377_v10).Dtor_cf3() {
						return _374_v7
					}
					return _dafny.SeqOf((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool))
				})(), 3)
				_378_v11 = _nw67
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_378_v11), 0))
				_ = _index46
				(_378_v11).ArraySet1(_374_v7, (_index46).Int())
				var _379_v12 _dafny.Map
				_ = _379_v12
				_379_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _376_v9)
				var _380_v13 _dafny.Set
				_ = _380_v13
				_380_v13 = _dafny.SetOf((func() _dafny.CodePoint {
					if (_379_v12).Contains((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool)) {
						return (_379_v12).Get((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool)).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('w')
				})(), _376_v9)
				var _381_v14 _dafny.Array
				_ = _381_v14
				var _nwElement0_20 _dafny.Int = _dafny.IntOfUint32((_368_v3).Cardinality())
				_ = _nwElement0_20
				var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(7))
				_ = _nw68
				(_nw68).ArraySet1(_nwElement0_20, 0)
				(_nw68).ArraySet1((_380_v13).Cardinality(), 1)
				(_nw68).ArraySet1((_this).F12(), 2)
				(_nw68).ArraySet1(_367_i0, 3)
				(_nw68).ArraySet1((_this).F12(), 4)
				(_nw68).ArraySet1(_367_i0, 5)
				(_nw68).ArraySet1((_this).F12(), 6)
				_381_v14 = _nw68
				var _382_v15 _dafny.Sequence
				_ = _382_v15
				_382_v15 = _dafny.SeqOf(_381_v14, _381_v14, _381_v14)
				var _383_v16 _dafny.Array
				_ = _383_v16
				var _nwElement0_21 _dafny.Array = (_382_v15).Select((Companion_Default___.SafeIndex(_367_i0, _dafny.IntOfUint32((_382_v15).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _nwElement0_21
				var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(3))
				_ = _nw69
				(_nw69).ArraySet1(_nwElement0_21, 0)
				(_nw69).ArraySet1(_381_v14, 1)
				(_nw69).ArraySet1(_381_v14, 2)
				_383_v16 = _nw69
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_383_v16), 0))
				_ = _index47
				(_383_v16).ArraySet1(_381_v14, (_index47).Int())
				var _384_v17 _dafny.Set
				_ = _384_v17
				_384_v17 = _dafny.SetOf((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool))
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_378_v11), 0))
				_ = _index48
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_383_v16), 0))
				_ = _index49
				var _rhs38 _dafny.CodePoint = Companion_Default___.Fm26(Companion_Default___.SafeDivisionInt(_367_i0, _367_i0), ((_dafny.SetOf((_this).F11())).Union(_384_v17)).Cardinality(), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool), globalState)
				_ = _rhs38
				var _rhs39 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_374_v7, _374_v7)
				_ = _rhs39
				var _rhs40 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_dafny.IntOfUint32((_368_v3).Cardinality())).Cmp(_367_i0) <= 0 {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_374_v7, _374_v7)).Cardinality())
					}
					return (_this).F12()
				})())
				_ = _rhs40
				var _rhs41 _dafny.Array = _381_v14
				_ = _rhs41
				var _rhs42 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_374_v7, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_374_v7).Cardinality()))).Uint32(), (func() bool {
					if true {
						return (_this).F11()
					}
					return (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))).Int()).(bool)
				})())
				_ = _rhs42
				var _lhs19 _dafny.Array = _378_v11
				_ = _lhs19
				var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_378_v11), 0))
				_ = _lhs20
				var _lhs21 *GlobalState = globalState
				_ = _lhs21
				var _lhs22 _dafny.Array = _383_v16
				_ = _lhs22
				var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_383_v16), 0))
				_ = _lhs23
				_376_v9 = _rhs38
				(_lhs19).ArraySet1(_rhs39, (_lhs20).Int())
				_lhs21.F2 = _rhs40
				(_lhs22).ArraySet1(_rhs41, (_lhs23).Int())
				_374_v7 = _rhs42
				(globalState).F2 = _367_i0
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_369_v4), 0))
				_ = _index50
				(_369_v4).ArraySet1((_this).F11(), (_index50).Int())
			} else {
				var _385_v18 _dafny.Array
				_ = _385_v18
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_7
				var _nw70 _dafny.Array
				_ = _nw70
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw70 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Int = func(_386_i3 _dafny.Int) _dafny.Int {
						return (_386_i3).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kisi")).Cardinality()))
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw70 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw70).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw70).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_385_v18 = _nw70
				var _387_v19 _dafny.MultiSet
				_ = _387_v19
				_387_v19 = _dafny.MultiSetOf(_385_v18)
				(globalState).F2 = (func() _dafny.Int {
					if (_387_v19).Contains(_385_v18) {
						return (_387_v19).Multiplicity(_385_v18)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("styupjoxw")).Cardinality())
				})()
				var _388_v20 _dafny.Sequence
				_ = _388_v20
				_388_v20 = _dafny.SeqOf((_this).F11())
				var _389_v21 *C0
				_ = _389_v21
				var _nw71 *C0 = New_C0_()
				_ = _nw71
				_nw71.Ctor__(_367_i0)
				_389_v21 = _nw71
				var _390_v22 D4
				_ = _390_v22
				_390_v22 = Companion_D4_.Create_DC8_(_389_v21)
				var _391_v23 _dafny.Map
				_ = _391_v23
				_391_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_389_v21).F13(), _389_v21)
				var _392_v24 _dafny.Array
				_ = _392_v24
				var _nwElement0_22 *C0 = _389_v21
				_ = _nwElement0_22
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(27))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_22, 0)
				(_nw72).ArraySet1(_389_v21, 1)
				(_nw72).ArraySet1(_389_v21, 2)
				(_nw72).ArraySet1(_389_v21, 3)
				(_nw72).ArraySet1(_389_v21, 4)
				(_nw72).ArraySet1(_389_v21, 5)
				(_nw72).ArraySet1(_389_v21, 6)
				(_nw72).ArraySet1(_389_v21, 7)
				(_nw72).ArraySet1(_389_v21, 8)
				(_nw72).ArraySet1(_389_v21, 9)
				(_nw72).ArraySet1((_390_v22).Dtor_cf15(), 10)
				(_nw72).ArraySet1(_389_v21, 11)
				(_nw72).ArraySet1((func() *C0 {
					if (_391_v23).Contains((_389_v21).F13()) {
						return (_391_v23).Get((_389_v21).F13()).(*C0)
					}
					return _389_v21
				})(), 12)
				(_nw72).ArraySet1(_389_v21, 13)
				(_nw72).ArraySet1(_389_v21, 14)
				(_nw72).ArraySet1(_389_v21, 15)
				(_nw72).ArraySet1(_389_v21, 16)
				(_nw72).ArraySet1(_389_v21, 17)
				(_nw72).ArraySet1(_389_v21, 18)
				(_nw72).ArraySet1(_389_v21, 19)
				(_nw72).ArraySet1(_389_v21, 20)
				(_nw72).ArraySet1(_389_v21, 21)
				(_nw72).ArraySet1(_389_v21, 22)
				(_nw72).ArraySet1((func() *C0 {
					if (_391_v23).Contains((_389_v21).F13()) {
						return (_391_v23).Get((_389_v21).F13()).(*C0)
					}
					return _389_v21
				})(), 23)
				(_nw72).ArraySet1(_389_v21, 24)
				(_nw72).ArraySet1(_389_v21, 25)
				(_nw72).ArraySet1(_389_v21, 26)
				_392_v24 = _nw72
				var _393_v25 _dafny.Map
				_ = _393_v25
				_393_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v24, (_389_v21).F13())
				var _394_v26 _dafny.Set
				_ = _394_v26
				_394_v26 = _dafny.SetOf((_this).F11(), (_dafny.IntOfUint32((_388_v20).Cardinality())).Cmp((_393_v25).Cardinality()) == 0, (_this).F11())
				var _395_v27 _dafny.Set
				_ = _395_v27
				_395_v27 = _394_v26
				_394_v26 = (_394_v26).Intersection((_395_v27))
				var _396_v28 _dafny.Sequence
				_ = _396_v28
				_396_v28 = _dafny.SeqOf(_367_i0)
				(globalState).F2 = (func() _dafny.Int {
					if !(_dafny.Companion_Sequence_.IsProperPrefixOf(_396_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(297))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg29 _dafny.Int) interface{} {
							return coer29(arg29)
						}
					}(func(_397_i4 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(885)
					})))) {
						return (_this).F12()
					}
					return (_this).F12()
				})()
				var _398_v29 _dafny.MultiSet
				_ = _398_v29
				_398_v29 = _dafny.MultiSetOf((_this).F11())
				(globalState).F2 = (func() _dafny.Int {
					if (_398_v29).Contains(true) {
						return (_398_v29).Multiplicity(true)
					}
					return _367_i0
				})()
				r2 = (_this).F11()
			}
			var _399_v30 _dafny.Array
			_ = _399_v30
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_8
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = func(_400_i5 _dafny.Int) _dafny.Int {
					return (_400_i5).Times((_this).F12())
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw73 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw73).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw73).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_399_v30 = _nw73
			_399_v30 = _399_v30
		}
		var _401_v31 _dafny.Array
		_ = _401_v31
		var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw74
		_401_v31 = _nw74
		var _402_v32 _dafny.Array
		_ = _402_v32
		var _nwElement0_23 _dafny.Array = _401_v31
		_ = _nwElement0_23
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(22))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_23, 0)
		(_nw75).ArraySet1(_401_v31, 1)
		(_nw75).ArraySet1(_401_v31, 2)
		(_nw75).ArraySet1(_401_v31, 3)
		(_nw75).ArraySet1(_401_v31, 4)
		(_nw75).ArraySet1(_401_v31, 5)
		(_nw75).ArraySet1(_401_v31, 6)
		(_nw75).ArraySet1(_401_v31, 7)
		(_nw75).ArraySet1(_401_v31, 8)
		(_nw75).ArraySet1(_401_v31, 9)
		(_nw75).ArraySet1(_401_v31, 10)
		(_nw75).ArraySet1(_401_v31, 11)
		(_nw75).ArraySet1(_401_v31, 12)
		(_nw75).ArraySet1(_401_v31, 13)
		(_nw75).ArraySet1(_401_v31, 14)
		(_nw75).ArraySet1(_401_v31, 15)
		(_nw75).ArraySet1(_401_v31, 16)
		(_nw75).ArraySet1(_401_v31, 17)
		(_nw75).ArraySet1(_401_v31, 18)
		(_nw75).ArraySet1(_401_v31, 19)
		(_nw75).ArraySet1(_401_v31, 20)
		(_nw75).ArraySet1(_401_v31, 21)
		_402_v32 = _nw75
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))
		_ = _index51
		(_402_v32).ArraySet1(_401_v31, (_index51).Int())
		var _403_v33 _dafny.Array
		_ = _403_v33
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_9
		var _nw76 _dafny.Array
		_ = _nw76
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw76 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = func(_404_i6 _dafny.Int) _dafny.Sequence {
				return (Companion_D6_.Create_DC11_((Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("fwqyvo"))).Dtor_cf22())).Dtor_cf22()
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw76 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw76).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw76).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_403_v33 = _nw76
		var _405_v34 _dafny.Sequence
		_ = _405_v34
		_405_v34 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		var _406_v35 _dafny.CodePoint
		_ = _406_v35
		_406_v35 = _dafny.CodePoint('x')
		var _407_v36 D6
		_ = _407_v36
		_407_v36 = Companion_D6_.Create_DC11_(_dafny.Companion_Sequence_.Update(_405_v34, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.IntOfUint32((_405_v34).Cardinality()))).Uint32(), _406_v35))
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))
		_ = _index52
		(_403_v33).ArraySet1((_407_v36).Dtor_cf22(), (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))
		_ = _index53
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))
		_ = _index54
		var _rhs43 _dafny.Array = (func() _dafny.Array {
			if (_this).F11() {
				return _401_v31
			}
			return _401_v31
		})()
		_ = _rhs43
		var _rhs44 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(48))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_408_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_409_i7 _dafny.Int) _dafny.CodePoint {
				return _408_v35
			}
		})(_406_v35))), _405_v34), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_410_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_411_i8 _dafny.Int) _dafny.CodePoint {
				return _410_v35
			}
		})(_406_v35))), _405_v34))
		_ = _rhs44
		var _rhs45 _dafny.Int = (_this).F12()
		_ = _rhs45
		var _rhs46 bool = (_this).F11()
		_ = _rhs46
		var _lhs24 _dafny.Array = _402_v32
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))
		_ = _lhs25
		var _lhs26 _dafny.Array = _403_v33
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))
		_ = _lhs27
		(_lhs24).ArraySet1(_rhs43, (_lhs25).Int())
		(_lhs26).ArraySet1(_rhs44, (_lhs27).Int())
		r1 = _rhs45
		r2 = _rhs46
		var _412_v37 bool
		_ = _412_v37
		var _413_v38 bool
		_ = _413_v38
		var _414_v39 _dafny.Int
		_ = _414_v39
		var _out20 bool
		_ = _out20
		var _out21 bool
		_ = _out21
		var _out22 _dafny.Int
		_ = _out22
		_out20, _out21, _out22 = Companion_Default___.M0((_this).F12(), globalState)
		_412_v37 = _out20
		_413_v38 = _out21
		_414_v39 = _out22
		var _415_v40 _dafny.Sequence
		_ = _415_v40
		_415_v40 = _dafny.SeqOf((_this).F12())
		var _hi3 _dafny.Int = _dafny.IntOfUint32((_415_v40).Cardinality())
		_ = _hi3
		for _416_i9 := _414_v39; _416_i9.Cmp(_hi3) < 0; _416_i9 = _416_i9.Plus(_dafny.One) {
			var _417_v41 _dafny.Sequence
			_ = _417_v41
			_417_v41 = _dafny.SeqOf(Companion_Default___.Fm27(_415_v40, !(_412_v37), _412_v37, globalState))
			r0 = _dafny.IntOfUint32((_417_v41).Cardinality())
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))
			_ = _arr0
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))
			_ = _index55
			(_arr0).ArraySet1(_dafny.IntOfInt64(483), (_index55).Int())
			var _418_v42 _dafny.Sequence
			_ = _418_v42
			_418_v42 = _dafny.SeqOf((_363_v0).Update(_413_v38, false))
			var _419_v43 D1
			_ = _419_v43
			_419_v43 = Companion_D1_.Create_DC1_(_418_v42)
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))
			_ = _arr1
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))
			_ = _index56
			var _rhs47 _dafny.CodePoint = _dafny.CodePoint('r')
			_ = _rhs47
			var _rhs48 _dafny.Int = _dafny.IntOfInt64(707)
			_ = _rhs48
			var _rhs49 D1 = _419_v43
			_ = _rhs49
			var _lhs28 _dafny.Array = _dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))
			_ = _lhs29
			_406_v35 = _rhs47
			(_lhs28).ArraySet1(_rhs48, (_lhs29).Int())
			_419_v43 = _rhs49
			_414_v39 = Companion_Default___.SafeModuloInt((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))).Int()).(_dafny.Int), ((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_415_v40).Cardinality())))
			var _420_v44 _dafny.Set
			_ = _420_v44
			_420_v44 = _dafny.SetOf(_401_v31)
			var _421_v45 D1
			_ = _421_v45
			_421_v45 = Companion_D1_.Create_DC2_((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))).Int()).(_dafny.Int), _412_v37, (_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))).Int()).(_dafny.Int))
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))
			_ = _arr2
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_dafny.ArrayCastTo((_402_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_402_v32), 0))).Int()))), 0))
			_ = _index57
			(_arr2).ArraySet1(((_420_v44).Cardinality()).Plus((_421_v45).Dtor_cf4()), (_index57).Int())
		}
		var _pat_let_tv2 = _414_v39
		_ = _pat_let_tv2
		var _pat_let_tv3 = _414_v39
		_ = _pat_let_tv3
		var _pat_let_tv4 = _414_v39
		_ = _pat_let_tv4
		var _pat_let_tv5 = _414_v39
		_ = _pat_let_tv5
		var _pat_let_tv6 = _414_v39
		_ = _pat_let_tv6
		var _pat_let_tv7 = _414_v39
		_ = _pat_let_tv7
		var _pat_let_tv8 = _414_v39
		_ = _pat_let_tv8
		var _pat_let_tv9 = _414_v39
		_ = _pat_let_tv9
		var _pat_let_tv10 = _414_v39
		_ = _pat_let_tv10
		var _source7 D2 = func(_source8 D1) D2 {
			if _source8.Is_DC2() {
				var _422___mcc_h7 _dafny.Int = _source8.Get_().(D1_DC2).Cf2
				_ = _422___mcc_h7
				var _423___mcc_h8 bool = _source8.Get_().(D1_DC2).Cf3
				_ = _423___mcc_h8
				var _424___mcc_h9 _dafny.Int = _source8.Get_().(D1_DC2).Cf4
				_ = _424___mcc_h9
				var _425_cf4 _dafny.Int = _424___mcc_h9
				_ = _425_cf4
				var _426_cf3 bool = _423___mcc_h8
				_ = _426_cf3
				var _427_cf2 _dafny.Int = _422___mcc_h7
				_ = _427_cf2
				return Companion_D2_.Create_DC3_(func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter17 := _dafny.Iterate((_dafny.SetOf(_dafny.MultiSetOf(_pat_let_tv2), _dafny.MultiSetOf(_pat_let_tv3, _pat_let_tv4, (_this).F12(), _pat_let_tv5))).Elements()); ; {
						_compr_16, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _428_v46 _dafny.MultiSet
						_428_v46 = interface{}(_compr_16).(_dafny.MultiSet)
						if (_dafny.SetOf(_dafny.MultiSetOf(_pat_let_tv6), _dafny.MultiSetOf(_pat_let_tv7, _pat_let_tv8, (_this).F12(), _pat_let_tv9))).Contains(_428_v46) {
							_coll16.Add(_428_v46, _427_cf2)
						}
					}
					return _coll16.ToMap()
				}())
			} else {
				var _429___mcc_h10 _dafny.Sequence = _source8.Get_().(D1_DC1).Cf1
				_ = _429___mcc_h10
				var _430_cf1 _dafny.Sequence = _429___mcc_h10
				_ = _430_cf1
				return Companion_D2_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-146)), _pat_let_tv10))
			}
		}(Companion_D1_.Create_DC2_(_414_v39, (_this).F11(), (_this).F12()))
		_ = _source7
		if _source7.Is_DC4() {
			var _431___mcc_h0 _dafny.Int = _source7.Get_().(D2_DC4).Cf6
			_ = _431___mcc_h0
			var _432_cf6 _dafny.Int = _431___mcc_h0
			_ = _432_cf6
			_412_v37 = (_412_v37) == (false)
			var _433_v47 *C0
			_ = _433_v47
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(_432_cf6)
			_433_v47 = _nw77
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_401_v31), 0))
			_ = _index58
			(_401_v31).ArraySet1(_dafny.IntOfInt64(809), (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_401_v31), 0))
			_ = _index59
			(_401_v31).ArraySet1((_414_v39).Times((_432_cf6).Times(_432_cf6)), (_index59).Int())
			_432_cf6 = (_dafny.Zero).Minus((_433_v47).F13())
		} else if _source7.Is_DC5() {
			var _434___mcc_h1 _dafny.Int = _source7.Get_().(D2_DC5).Cf7
			_ = _434___mcc_h1
			var _435___mcc_h2 bool = _source7.Get_().(D2_DC5).Cf8
			_ = _435___mcc_h2
			var _436___mcc_h3 bool = _source7.Get_().(D2_DC5).Cf9
			_ = _436___mcc_h3
			var _437___mcc_h4 bool = _source7.Get_().(D2_DC5).Cf10
			_ = _437___mcc_h4
			var _438___mcc_h5 bool = _source7.Get_().(D2_DC5).Cf11
			_ = _438___mcc_h5
			var _439_cf11 bool = _438___mcc_h5
			_ = _439_cf11
			var _440_cf10 bool = _437___mcc_h4
			_ = _440_cf10
			var _441_cf9 bool = _436___mcc_h3
			_ = _441_cf9
			var _442_cf8 bool = _435___mcc_h2
			_ = _442_cf8
			var _443_cf7 _dafny.Int = _434___mcc_h1
			_ = _443_cf7
			var _444_v48 _dafny.Sequence
			_ = _444_v48
			_444_v48 = _dafny.SeqOf(_405_v34)
			var _445_v49 _dafny.Map
			_ = _445_v49
			_445_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _406_v35)
			var _446_v50 _dafny.Array
			_ = _446_v50
			var _nwElement0_24 bool = _441_cf9
			_ = _nwElement0_24
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_24, 0)
			(_nw78).ArraySet1(!(true), 1)
			(_nw78).ArraySet1(_412_v37, 2)
			(_nw78).ArraySet1(_440_cf10, 3)
			(_nw78).ArraySet1(_413_v38, 4)
			(_nw78).ArraySet1(_412_v37, 5)
			(_nw78).ArraySet1(!(_439_cf11), 6)
			(_nw78).ArraySet1(_412_v37, 7)
			(_nw78).ArraySet1(_442_cf8, 8)
			(_nw78).ArraySet1(_440_cf10, 9)
			(_nw78).ArraySet1(!_dafny.Companion_Sequence_.Equal(_444_v48, _444_v48), 10)
			(_nw78).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("odq"), (func() _dafny.CodePoint {
				if (_445_v49).Contains((func() bool {
					if (_363_v0).Contains(_439_cf11) {
						return (_363_v0).Get(_439_cf11).(bool)
					}
					return _440_cf10
				})()) {
					return (_445_v49).Get((func() bool {
						if (_363_v0).Contains(_439_cf11) {
							return (_363_v0).Get(_439_cf11).(bool)
						}
						return _440_cf10
					})()).(_dafny.CodePoint)
				}
				return _406_v35
			})()), 11)
			_446_v50 = _nw78
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_446_v50), 0))
			_ = _index60
			(_446_v50).ArraySet1(_439_cf11, (_index60).Int())
			var _447_v52 _dafny.Map
			_ = _447_v52
			_447_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_403_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))).Int()).(_dafny.Sequence)).Cardinality()), _413_v38)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_446_v50), 0))
			_ = _index61
			var _rhs50 _dafny.Int = _443_cf7
			_ = _rhs50
			var _rhs51 _dafny.Int = (_dafny.Zero).Minus((((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(865), _dafny.IntOfInt64(185))); ; {
					_compr_17, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _448_v51 _dafny.Int
					_448_v51 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(865)).Cmp(_448_v51) <= 0) && ((_448_v51).Cmp(_dafny.IntOfInt64(185)) < 0) {
						_coll17.Add(Companion_Default___.SafeDivisionInt(_448_v51, _dafny.IntOfInt64(235)), true)
					}
				}
				return _coll17.ToMap()
			}()).Merge(((_447_v52).Update((_this).F12(), _439_cf11)).Update(_dafny.IntOfUint32((_415_v40).Cardinality()), Companion_Default___.Fm23((_403_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))).Int()).(_dafny.Sequence), (_this).F12(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_449_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_450_i10 _dafny.Int) _dafny.CodePoint {
					return _449_v35
				}
			})(_406_v35)))).Cardinality()), _439_cf11, globalState)))).Merge((_447_v52).Update((_this).F12(), _439_cf11))).Cardinality())
			_ = _rhs51
			var _rhs52 bool = _442_cf8
			_ = _rhs52
			var _rhs53 bool = true
			_ = _rhs53
			var _lhs30 _dafny.Array = _446_v50
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_446_v50), 0))
			_ = _lhs31
			r1 = _rhs50
			r1 = _rhs51
			(_lhs30).ArraySet1(_rhs52, (_lhs31).Int())
			_441_cf9 = _rhs53
			var _451_v53 D3
			_ = _451_v53
			_451_v53 = Companion_D3_.Create_DC6_(_406_v35)
			_406_v35 = (_451_v53).Dtor_cf12()
			r2 = (_446_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_446_v50), 0))).Int()).(bool)
			var _452_v54 D2
			_ = _452_v54
			_452_v54 = Companion_D2_.Create_DC5_(_443_cf7, _439_cf11, (_this).F11(), _439_cf11, _442_cf8)
			var _pat_let_tv11 = _439_cf11
			_ = _pat_let_tv11
			var _pat_let_tv12 = _446_v50
			_ = _pat_let_tv12
			var _pat_let_tv13 = _446_v50
			_ = _pat_let_tv13
			_442_cf8 = (func(_pat_let0_0 D2) D2 {
				return func(_453_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 bool) D2 {
						return func(_454_dt__update_hcf9_h0 bool) D2 {
							return func(_pat_let2_0 bool) D2 {
								return func(_455_dt__update_hcf8_h0 bool) D2 {
									return func(_pat_let3_0 bool) D2 {
										return func(_456_dt__update_hcf11_h0 bool) D2 {
											return Companion_D2_.Create_DC5_((_453_dt__update__tmp_h0).Dtor_cf7(), _455_dt__update_hcf8_h0, _454_dt__update_hcf9_h0, (_453_dt__update__tmp_h0).Dtor_cf10(), _456_dt__update_hcf11_h0)
										}(_pat_let3_0)
									}((_this).F11())
								}(_pat_let2_0)
							}((_pat_let_tv13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_pat_let_tv12), 0))).Int()).(bool))
						}(_pat_let1_0)
					}(!(!(_pat_let_tv11)))
				}(_pat_let0_0)
			}(_452_v54)).Dtor_cf11()
		} else {
			var _457___mcc_h6 _dafny.Map = _source7.Get_().(D2_DC3).Cf5
			_ = _457___mcc_h6
			var _458_cf5 _dafny.Map = _457___mcc_h6
			_ = _458_cf5
			r1 = _414_v39
			var _459_v55 *C0
			_ = _459_v55
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__((_this).F12())
			_459_v55 = _nw79
			var _460_v56 _dafny.Map
			_ = _460_v56
			_460_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_459_v55).F13())
			_413_v38 = ((_460_v56).Contains(_414_v39)) && ((_this).F11())
			r1 = (_this).F12()
		}
		r0 = _414_v39
		var _461_v57 _dafny.MultiSet
		_ = _461_v57
		_461_v57 = _dafny.MultiSetOf(true)
		var _462_v58 _dafny.Sequence
		_ = _462_v58
		_462_v58 = _dafny.SeqOf((_this).F11(), _412_v37)
		r1 = ((_414_v39).Minus((_461_v57).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_403_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_462_v58).Cardinality()), _dafny.IntOfUint32(((_403_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_403_v33), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _406_v35)).Cardinality()))
		r2 = _412_v37
		return r0, r1, r2
	}
}
func (_this *C2) M7(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _463_v0 _dafny.Array
		_ = _463_v0
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw80
		_463_v0 = _nw80
		var _464_v1 _dafny.Sequence
		_ = _464_v1
		_464_v1 = _dafny.SeqOf((_this).F12())
		var _465_v2 _dafny.Map
		_ = _465_v2
		_465_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_463_v0, _464_v1)
		_465_v2 = ((_465_v2).Merge(_465_v2)).Merge((_465_v2).Merge(_465_v2))
		if p1 {
			var _466_v3 _dafny.Sequence
			_ = _466_v3
			_466_v3 = _dafny.UnicodeSeqOfUtf8Bytes("jmjjfn")
			r1 = _dafny.Companion_Sequence_.Concatenate(_466_v3, _466_v3)
			var _467_v4 _dafny.Array
			_ = _467_v4
			var _nwElement0_25 _dafny.Array = _463_v0
			_ = _nwElement0_25
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(6))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_25, 0)
			(_nw81).ArraySet1(_463_v0, 1)
			(_nw81).ArraySet1(_463_v0, 2)
			(_nw81).ArraySet1(_463_v0, 3)
			(_nw81).ArraySet1(_463_v0, 4)
			(_nw81).ArraySet1(_463_v0, 5)
			_467_v4 = _nw81
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_467_v4), 0))
			_ = _index62
			(_467_v4).ArraySet1(_463_v0, (_index62).Int())
			var _468_v5 _dafny.Map
			_ = _468_v5
			_468_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F12()).Plus((_this).F12()), _463_v0)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_467_v4), 0))
			_ = _index63
			(_467_v4).ArraySet1((func() _dafny.Array {
				if (_468_v5).Contains((_this).F12()) {
					return (_468_v5).Get((_this).F12()).(_dafny.Array)
				}
				return _463_v0
			})(), (_index63).Int())
			var _469_v6 _dafny.Map
			_ = _469_v6
			_469_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p1)
			var _470_v7 _dafny.Set
			_ = _470_v7
			_470_v7 = _dafny.SetOf(_469_v6, _469_v6, _469_v6)
			_470_v7 = _470_v7
			var _471_v8 bool
			_ = _471_v8
			_471_v8 = true
			var _472_v9 _dafny.Sequence
			_ = _472_v9
			_472_v9 = _dafny.SeqOf(_469_v6)
			_471_v8 = Companion_Default___.Fm23(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_473_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), _dafny.IntOfUint32((_472_v9).Cardinality()), (_this).F12(), p0, globalState)
			_471_v8 = ((_this).F12()).Cmp((func() _dafny.Int {
				if p0 {
					return (_this).F12()
				}
				return _dafny.IntOfInt64(941)
			})()) < 0
		} else {
			var _474_v10 _dafny.Set
			_ = _474_v10
			_474_v10 = _dafny.SetOf((_this).F12())
			(globalState).F2 = Companion_Default___.Fm24((_474_v10).Difference(_474_v10), globalState)
			var _475_v11 *C0
			_ = _475_v11
			var _nw82 *C0 = New_C0_()
			_ = _nw82
			_nw82.Ctor__((_this).F12())
			_475_v11 = _nw82
			var _476_v12 _dafny.Sequence
			_ = _476_v12
			_476_v12 = _dafny.SeqOf(_475_v11)
			var _477_v13 _dafny.Array
			_ = _477_v13
			var _nwElement0_26 _dafny.Int = _dafny.IntOfUint32((_476_v12).Cardinality())
			_ = _nwElement0_26
			var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(6))
			_ = _nw83
			(_nw83).ArraySet1(_nwElement0_26, 0)
			(_nw83).ArraySet1(Companion_Default___.SafeModuloInt((_475_v11).F13(), (_475_v11).F13()), 1)
			(_nw83).ArraySet1(((_this).F12()).Minus((_dafny.Zero).Minus((_475_v11).F13())), 2)
			(_nw83).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_464_v1, _dafny.SeqOf((_475_v11).F13(), _dafny.IntOfUint32((_464_v1).Cardinality()), (_this).F12())), (Companion_Default___.SafeIndex((_475_v11).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_464_v1, _dafny.SeqOf((_475_v11).F13(), _dafny.IntOfUint32((_464_v1).Cardinality()), (_this).F12()))).Cardinality()))).Uint32(), (_475_v11).F13())).Cardinality()), 3)
			(_nw83).ArraySet1((_this).F12(), 4)
			(_nw83).ArraySet1((_this).F12(), 5)
			_477_v13 = _nw83
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_477_v13), 0))
			_ = _index64
			(_477_v13).ArraySet1((_this).F12(), (_index64).Int())
			var _478_v14 D3
			_ = _478_v14
			_478_v14 = Companion_D3_.Create_DC7_((_this).F12(), (_dafny.IntOfInt64(316)).Times((_this).F12()))
			var _479_v15 _dafny.MultiSet
			_ = _479_v15
			_479_v15 = _dafny.MultiSetOf((_475_v11).F13(), (_this).F12())
			var _480_v16 _dafny.Map
			_ = _480_v16
			_480_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_479_v15, _dafny.IntOfInt64(-719))
			var _481_v17 D2
			_ = _481_v17
			_481_v17 = Companion_D2_.Create_DC3_(_480_v16)
			var _482_v18 D2
			_ = _482_v18
			_482_v18 = Companion_D2_.Create_DC3_((_481_v17).Dtor_cf5())
			var _483_v19 D4
			_ = _483_v19
			_483_v19 = Companion_D4_.Create_DC9_((_this).F12(), _dafny.CodePoint('e'), _464_v1, (_475_v11).F13(), (_this).F11())
			var _484_v20 _dafny.Sequence
			_ = _484_v20
			_484_v20 = _dafny.SeqOf(_483_v19)
			var _485_v21 _dafny.Sequence
			_ = _485_v21
			_485_v21 = _dafny.SeqOf(_484_v20)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_477_v13), 0))
			_ = _index65
			var _rhs54 _dafny.Int = Companion_Default___.SafeModuloInt((_475_v11).F13(), (_475_v11).F13())
			_ = _rhs54
			var _rhs55 _dafny.Int = _dafny.IntOfUint32((_485_v21).Cardinality())
			_ = _rhs55
			var _rhs56 D3 = _478_v14
			_ = _rhs56
			var _rhs57 D2 = _481_v17
			_ = _rhs57
			var _lhs32 _dafny.Array = _477_v13
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_477_v13), 0))
			_ = _lhs33
			r0 = _rhs54
			(_lhs32).ArraySet1(_rhs55, (_lhs33).Int())
			_478_v14 = _rhs56
			_481_v17 = _rhs57
			var _486_v22 _dafny.Int
			_ = _486_v22
			var _487_v23 _dafny.Int
			_ = _487_v23
			var _488_v24 bool
			_ = _488_v24
			var _out23 _dafny.Int
			_ = _out23
			var _out24 _dafny.Int
			_ = _out24
			var _out25 bool
			_ = _out25
			_out23, _out24, _out25 = (_this).M6(globalState)
			_486_v22 = _out23
			_487_v23 = _out24
			_488_v24 = _out25
			_488_v24 = !((func() bool {
				if p1 {
					return (_this).F11()
				}
				return true
			})())
			var _489_v25 _dafny.Array
			_ = _489_v25
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw84
			_489_v25 = _nw84
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_489_v25), 0))
			_ = _index66
			(_489_v25).ArraySet1(_477_v13, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_463_v0), 0))
			_ = _index67
			(_463_v0).ArraySet1(!(!(false)), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_477_v13), 0))
			_ = _index68
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_489_v25), 0))
			_ = _index69
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_463_v0), 0))
			_ = _index70
			var _rhs58 _dafny.Int = _dafny.IntOfInt64(365)
			_ = _rhs58
			var _rhs59 _dafny.Array = _477_v13
			_ = _rhs59
			var _rhs60 bool = p1
			_ = _rhs60
			var _lhs34 _dafny.Array = _477_v13
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_477_v13), 0))
			_ = _lhs35
			var _lhs36 _dafny.Array = _489_v25
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_489_v25), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _463_v0
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_463_v0), 0))
			_ = _lhs39
			(_lhs34).ArraySet1(_rhs58, (_lhs35).Int())
			(_lhs36).ArraySet1(_rhs59, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs60, (_lhs39).Int())
		}
		var _490_v26 _dafny.Sequence
		_ = _490_v26
		_490_v26 = _dafny.UnicodeSeqOfUtf8Bytes("vtttnk")
		var _491_v28 _dafny.Array
		_ = _491_v28
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_10
		var _nw85 _dafny.Array
		_ = _nw85
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw85 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) D2 = func(_492_i1 _dafny.Int) D2 {
				return Companion_D2_.Create_DC3_(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(91)), _dafny.MultiSetOf((_this).F12()))).Elements()); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _493_v27 _dafny.MultiSet
						_493_v27 = interface{}(_compr_18).(_dafny.MultiSet)
						if (_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(91)), _dafny.MultiSetOf((_this).F12()))).Contains(_493_v27) {
							_coll18.Add(_493_v27, (_this).F12())
						}
					}
					return _coll18.ToMap()
				}())
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw85 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw85).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw85).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_491_v28 = _nw85
		var _494_v29 _dafny.Map
		_ = _494_v29
		_494_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_491_v28, _490_v26)
		var _495_v30 _dafny.CodePoint
		_ = _495_v30
		_495_v30 = _dafny.CodePoint('u')
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25(!((_this).F11()), (_this).F12(), globalState), _490_v26), (func() _dafny.Sequence {
			if (_494_v29).Contains(_491_v28) {
				return (_494_v29).Get(_491_v28).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("i")
		})()), (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25(!((_this).F11()), (_this).F12(), globalState), _490_v26), (func() _dafny.Sequence {
			if (_494_v29).Contains(_491_v28) {
				return (_494_v29).Get(_491_v28).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("i")
		})())).Cardinality()))).Uint32(), _495_v30)
		(globalState).F2 = _dafny.IntOfInt64(238)
		var _496_v31 _dafny.Map
		_ = _496_v31
		_496_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _463_v0)
		var _497_v32 _dafny.Map
		_ = _497_v32
		_497_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (func() _dafny.Array {
			if (_496_v31).Contains((_this).F11()) {
				return (_496_v31).Get((_this).F11()).(_dafny.Array)
			}
			return _463_v0
		})())
		_496_v31 = _497_v32
		var _498_v33 _dafny.Set
		_ = _498_v33
		_498_v33 = _dafny.SetOf(_dafny.IntOfInt64(616), (_this).F12())
		var _499_v35 D1
		_ = _499_v35
		_499_v35 = Companion_D1_.Create_DC2_(((_this).F12()).Plus(_dafny.IntOfInt64(611)), !((_498_v33).Equals(func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(835), _dafny.IntOfInt64(567))); ; {
				_compr_19, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _500_v34 _dafny.Int
				_500_v34 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(835)).Cmp(_500_v34) <= 0) && ((_500_v34).Cmp(_dafny.IntOfInt64(567)) < 0) {
					_coll19.Add((_500_v34).Minus(_dafny.IntOfInt64(97)))
				}
			}
			return _coll19.ToSet()
		}())), (_this).F12())
		var _source9 D1 = _499_v35
		_ = _source9
		if _source9.Is_DC2() {
			var _501___mcc_h0 _dafny.Int = _source9.Get_().(D1_DC2).Cf2
			_ = _501___mcc_h0
			var _502___mcc_h1 bool = _source9.Get_().(D1_DC2).Cf3
			_ = _502___mcc_h1
			var _503___mcc_h2 _dafny.Int = _source9.Get_().(D1_DC2).Cf4
			_ = _503___mcc_h2
			var _504_cf4 _dafny.Int = _503___mcc_h2
			_ = _504_cf4
			var _505_cf3 bool = _502___mcc_h1
			_ = _505_cf3
			var _506_cf2 _dafny.Int = _501___mcc_h0
			_ = _506_cf2
			_505_cf3 = (_this).F11()
			var _507_v36 *C0
			_ = _507_v36
			var _nw86 *C0 = New_C0_()
			_ = _nw86
			_nw86.Ctor__(_dafny.IntOfInt64(-580))
			_507_v36 = _nw86
			var _508_v37 _dafny.Set
			_ = _508_v37
			_508_v37 = _dafny.SetOf(_495_v30, _495_v30)
			var _509_v38 D7
			_ = _509_v38
			_509_v38 = Companion_D7_.Create_DC13_(_508_v37)
			_508_v37 = (_508_v37).Union((_509_v38).Dtor_cf25())
			var _510_v39 _dafny.Map
			_ = _510_v39
			_510_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _dafny.IntOfUint32((_490_v26).Cardinality()))
			var _511_v40 _dafny.MultiSet
			_ = _511_v40
			_511_v40 = _dafny.MultiSetOf(_510_v39)
			var _512_v41 D2
			_ = _512_v41
			_512_v41 = Companion_D2_.Create_DC4_((func() _dafny.Int {
				if (_511_v40).Contains(Companion_Default___.Fm28(_506_cf2, globalState)) {
					return (_511_v40).Multiplicity(Companion_Default___.Fm28(_506_cf2, globalState))
				}
				return _506_cf2
			})())
			var _513_v42 _dafny.Sequence
			_ = _513_v42
			_513_v42 = _dafny.SeqOf((_this).F11(), (_this).F11())
			var _514_v43 _dafny.Map
			_ = _514_v43
			_514_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC13_(_508_v37), _dafny.IntOfInt64(-23))
			_512_v41 = (func() D2 {
				if (_513_v42).Select((Companion_Default___.SafeIndex((_514_v43).Cardinality(), _dafny.IntOfUint32((_513_v42).Cardinality()))).Uint32()).(bool) {
					return (func() D2 {
						if p1 {
							return _512_v41
						}
						return Companion_D2_.Create_DC4_((_this).F12())
					})()
				}
				return _512_v41
			})()
		} else {
			var _515___mcc_h3 _dafny.Sequence = _source9.Get_().(D1_DC1).Cf1
			_ = _515___mcc_h3
			var _516_cf1 _dafny.Sequence = _515___mcc_h3
			_ = _516_cf1
			if p0 {
				var _517_v44 _dafny.Set
				_ = _517_v44
				_517_v44 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("c"), _490_v26)
				var _518_v45 D6
				_ = _518_v45
				_518_v45 = Companion_D6_.Create_DC12_((_this).F11(), Companion_Default___.Fm26((_this).F12(), (_this).F12(), (_this).F11(), globalState))
				_490_v26 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("xramcle"), (Companion_Default___.SafeIndex((_517_v44).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xramcle")).Cardinality()))).Uint32(), (_518_v45).Dtor_cf24()), _490_v26)
				var _519_v46 *C0
				_ = _519_v46
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__((_this).F12())
				_519_v46 = _nw87
				var _520_v47 bool
				_ = _520_v47
				_520_v47 = false
				var _521_v48 _dafny.Sequence
				_ = _521_v48
				_521_v48 = _dafny.SeqOf(_519_v46, _519_v46, _519_v46)
				var _522_v49 _dafny.Map
				_ = _522_v49
				_522_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_519_v46).F13(), _490_v26)
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_463_v0), 0))
				_ = _index71
				(_463_v0).ArraySet1(p1, (_index71).Int())
				var _523_v50 D8
				_ = _523_v50
				_523_v50 = Companion_D8_.Create_DC15_(_463_v0)
				var _524_v51 _dafny.Map
				_ = _524_v51
				_524_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_523_v50).Dtor_cf27(), p1)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_463_v0), 0))
				_ = _index72
				var _rhs61 bool = _520_v47
				_ = _rhs61
				var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_521_v48, _dafny.Companion_Sequence_.Update(_521_v48, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_521_v48).Cardinality()))).Uint32(), (_521_v48).Select((Companion_Default___.SafeIndex((_519_v46).F13(), _dafny.IntOfUint32((_521_v48).Cardinality()))).Uint32()).(*C0)))
				_ = _rhs62
				var _rhs63 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), Companion_Default___.Fm25((func() bool {
					if (_524_v51).Contains(_463_v0) {
						return (_524_v51).Get(_463_v0).(bool)
					}
					return p0
				})(), _dafny.IntOfInt64(185), globalState))
				_ = _rhs63
				var _rhs64 bool = true
				_ = _rhs64
				var _lhs40 _dafny.Array = _463_v0
				_ = _lhs40
				var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_463_v0), 0))
				_ = _lhs41
				_520_v47 = _rhs61
				_521_v48 = _rhs62
				_522_v49 = _rhs63
				(_lhs40).ArraySet1(_rhs64, (_lhs41).Int())
				r0 = ((_this).F12()).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_490_v26, _490_v26)).Cardinality()))
				var _525_v52 _dafny.MultiSet
				_ = _525_v52
				_525_v52 = _dafny.MultiSetOf((_519_v46).F13())
				var _526_v53 _dafny.Sequence
				_ = _526_v53
				_526_v53 = _dafny.SeqOf((_525_v52).IsDisjointFrom(_dafny.MultiSetOf((_519_v46).F13())), p1, p0, (_this).F11())
				_526_v53 = (func() _dafny.Sequence {
					if false {
						return _526_v53
					}
					return _dafny.SeqOf(p1, _520_v47, p1)
				})()
			} else {
				var _527_v54 _dafny.Map
				_ = _527_v54
				_527_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _dafny.Companion_Sequence_.IsProperPrefixOf(_490_v26, _490_v26))
				var _528_v55 _dafny.MultiSet
				_ = _528_v55
				_528_v55 = _dafny.MultiSetOf(_464_v1, _464_v1, _dafny.SeqOf(_dafny.IntOfInt64(-84)))
				_527_v54 = (_527_v54).Update((func() _dafny.Int {
					if (_528_v55).Contains(_464_v1) {
						return (_528_v55).Multiplicity(_464_v1)
					}
					return (_this).F12()
				})(), (_this).F11())
				var _529_v56 bool
				_ = _529_v56
				var _530_v57 bool
				_ = _530_v57
				var _531_v58 _dafny.Int
				_ = _531_v58
				var _out26 bool
				_ = _out26
				var _out27 bool
				_ = _out27
				var _out28 _dafny.Int
				_ = _out28
				_out26, _out27, _out28 = Companion_Default___.M0((_this).F12(), globalState)
				_529_v56 = _out26
				_530_v57 = _out27
				_531_v58 = _out28
				_529_v56 = ((func() _dafny.Set {
					if _530_v57 {
						return _498_v33
					}
					return _498_v33
				})()).IsProperSubsetOf(_498_v33)
				_463_v0 = _463_v0
				var _532_v59 bool
				_ = _532_v59
				var _533_v60 bool
				_ = _533_v60
				var _534_v61 _dafny.Int
				_ = _534_v61
				var _out29 bool
				_ = _out29
				var _out30 bool
				_ = _out30
				var _out31 _dafny.Int
				_ = _out31
				_out29, _out30, _out31 = Companion_Default___.M0((_this).F12(), globalState)
				_532_v59 = _out29
				_533_v60 = _out30
				_534_v61 = _out31
			}
			var _535_v62 D8
			_ = _535_v62
			_535_v62 = Companion_D8_.Create_DC15_(_463_v0)
			var _pat_let_tv14 = _463_v0
			_ = _pat_let_tv14
			var _source10 D8 = func(_pat_let4_0 D8) D8 {
				return func(_536_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let5_0 _dafny.Array) D8 {
						return func(_537_dt__update_hcf27_h0 _dafny.Array) D8 {
							return Companion_D8_.Create_DC15_(_537_dt__update_hcf27_h0)
						}(_pat_let5_0)
					}(_pat_let_tv14)
				}(_pat_let4_0)
			}(_535_v62)
			_ = _source10
			if _source10.Is_DC16() {
				var _538___mcc_h4 _dafny.Int = _source10.Get_().(D8_DC16).Cf28
				_ = _538___mcc_h4
				var _539___mcc_h5 bool = _source10.Get_().(D8_DC16).Cf29
				_ = _539___mcc_h5
				var _540___mcc_h6 bool = _source10.Get_().(D8_DC16).Cf30
				_ = _540___mcc_h6
				var _541___mcc_h7 _dafny.Int = _source10.Get_().(D8_DC16).Cf31
				_ = _541___mcc_h7
				var _542_cf31 _dafny.Int = _541___mcc_h7
				_ = _542_cf31
				var _543_cf30 bool = _540___mcc_h6
				_ = _543_cf30
				var _544_cf29 bool = _539___mcc_h5
				_ = _544_cf29
				var _545_cf28 _dafny.Int = _538___mcc_h4
				_ = _545_cf28
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_463_v0), 0))
				_ = _index73
				(_463_v0).ArraySet1((_498_v33).IsSubsetOf(_498_v33), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_463_v0), 0))
				_ = _index74
				(_463_v0).ArraySet1(false, (_index74).Int())
				var _546_v63 _dafny.MultiSet
				_ = _546_v63
				_546_v63 = _dafny.MultiSetOf(_542_cf31, _dafny.IntOfUint32((_490_v26).Cardinality()))
				(globalState).F2 = ((_545_cf28).Plus((_546_v63).Cardinality())).Times(_542_cf31)
				var _547_v64 _dafny.Set
				_ = _547_v64
				_547_v64 = _dafny.SetOf(p0)
				var _548_v65 _dafny.Array
				_ = _548_v65
				var _nwElement0_27 _dafny.Int = _542_cf31
				_ = _nwElement0_27
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(2))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_27, 0)
				(_nw88).ArraySet1(Companion_Default___.SafeModuloInt(_545_cf28, (_547_v64).Cardinality()), 1)
				_548_v65 = _nw88
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_548_v65), 0))
				_ = _index75
				(_548_v65).ArraySet1(_545_cf28, (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_548_v65), 0))
				_ = _index76
				(_548_v65).ArraySet1(_dafny.IntOfInt64(-366), (_index76).Int())
				var _549_v66 D8
				_ = _549_v66
				_549_v66 = Companion_D8_.Create_DC16_(_542_cf31, p0, (_this).F11(), _545_cf28)
				_544_cf29 = (_549_v66).Dtor_cf29()
			} else if _source10.Is_DC15() {
				var _550___mcc_h8 _dafny.Array = _source10.Get_().(D8_DC15).Cf27
				_ = _550___mcc_h8
				var _551_cf27 _dafny.Array = _550___mcc_h8
				_ = _551_cf27
				var _552_v67 T0
				_ = _552_v67
				var _nw89 *C1 = New_C1_()
				_ = _nw89
				_nw89.Ctor__((_464_v1).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_464_v1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F3())
				_552_v67 = _nw89
				_552_v67 = _552_v67
				var _553_v68 _dafny.Int
				_ = _553_v68
				var _554_v69 _dafny.Int
				_ = _554_v69
				var _555_v70 bool
				_ = _555_v70
				var _out32 _dafny.Int
				_ = _out32
				var _out33 _dafny.Int
				_ = _out33
				var _out34 bool
				_ = _out34
				_out32, _out33, _out34 = (_this).M6(globalState)
				_553_v68 = _out32
				_554_v69 = _out33
				_555_v70 = _out34
				r1 = _dafny.Companion_Sequence_.Concatenate(_490_v26, _dafny.Companion_Sequence_.Concatenate(_490_v26, _490_v26))
				r1 = _490_v26
			} else {
				var _556___mcc_h9 D8 = _source10.Get_().(D8_DC17).Cf32
				_ = _556___mcc_h9
				var _557_cf32 D8 = _556___mcc_h9
				_ = _557_cf32
				var _558_v71 bool
				_ = _558_v71
				_558_v71 = false
				var _559_v72 *C0
				_ = _559_v72
				var _nw90 *C0 = New_C0_()
				_ = _nw90
				_nw90.Ctor__((_this).F12())
				_559_v72 = _nw90
				var _560_v73 D4
				_ = _560_v73
				_560_v73 = Companion_D4_.Create_DC8_(_559_v72)
				var _rhs65 bool = p0
				_ = _rhs65
				var _rhs66 bool = p0
				_ = _rhs66
				var _rhs67 D4 = _560_v73
				_ = _rhs67
				_558_v71 = _rhs65
				_558_v71 = _rhs66
				_560_v73 = _rhs67
				var _561_v74 _dafny.MultiSet
				_ = _561_v74
				_561_v74 = _dafny.MultiSetOf(_558_v71)
				var _562_v75 _dafny.Array
				_ = _562_v75
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw91
				_562_v75 = _nw91
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_562_v75), 0))
				_ = _index77
				(_562_v75).ArraySet1((_this).F12(), (_index77).Int())
				var _563_v76 _dafny.Map
				_ = _563_v76
				_563_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_558_v71, _561_v74)
				var _564_v77 _dafny.Map
				_ = _564_v77
				_564_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("eeosvtnfb"), (func() _dafny.MultiSet {
					if (_563_v76).Contains((_this).F11()) {
						return (_563_v76).Get((_this).F11()).(_dafny.MultiSet)
					}
					return _561_v74
				})())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_562_v75), 0))
				_ = _index78
				var _rhs68 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_490_v26, _490_v26), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_490_v26, _490_v26), (Companion_Default___.SafeIndex((_559_v72).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_490_v26, _490_v26)).Cardinality()))).Uint32(), _495_v30))).Cardinality())
				_ = _rhs68
				var _rhs69 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_559_v72).F13()), (_559_v72).F13())
				_ = _rhs69
				var _rhs70 _dafny.MultiSet = (func() _dafny.MultiSet {
					if (_564_v77).Contains((func() _dafny.Sequence {
						if Companion_Default___.Fm23(_490_v26, (_559_v72).F13(), (_559_v72).F13(), _558_v71, globalState) {
							return _490_v26
						}
						return _dafny.Companion_Sequence_.Update(_490_v26, (Companion_Default___.SafeIndex((Companion_Default___.Fm30((_559_v72).F13(), (_this).F12(), _dafny.UnicodeSeqOfUtf8Bytes("hkitllpy"), globalState)).Dtor_cf7(), _dafny.IntOfUint32((_490_v26).Cardinality()))).Uint32(), _dafny.CodePoint('f'))
					})()) {
						return (_564_v77).Get((func() _dafny.Sequence {
							if Companion_Default___.Fm23(_490_v26, (_559_v72).F13(), (_559_v72).F13(), _558_v71, globalState) {
								return _490_v26
							}
							return _dafny.Companion_Sequence_.Update(_490_v26, (Companion_Default___.SafeIndex((Companion_Default___.Fm30((_559_v72).F13(), (_this).F12(), _dafny.UnicodeSeqOfUtf8Bytes("hkitllpy"), globalState)).Dtor_cf7(), _dafny.IntOfUint32((_490_v26).Cardinality()))).Uint32(), _dafny.CodePoint('f'))
						})()).(_dafny.MultiSet)
					}
					return _561_v74
				})()
				_ = _rhs70
				var _rhs71 _dafny.Int = (_this).F12()
				_ = _rhs71
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				var _lhs43 *GlobalState = globalState
				_ = _lhs43
				var _lhs44 _dafny.Array = _562_v75
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_562_v75), 0))
				_ = _lhs45
				_lhs42.F2 = _rhs68
				_lhs43.F2 = _rhs69
				_561_v74 = _rhs70
				(_lhs44).ArraySet1(_rhs71, (_lhs45).Int())
				var _565_v78 *C0
				_ = _565_v78
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__((_464_v1).Select((Companion_Default___.SafeIndex((_559_v72).F13(), _dafny.IntOfUint32((_464_v1).Cardinality()))).Uint32()).(_dafny.Int))
				_565_v78 = _nw92
				var _566_v79 _dafny.Sequence
				_ = _566_v79
				_566_v79 = _dafny.SeqOf(!(_558_v71))
				var _567_v80 _dafny.Sequence
				_ = _567_v80
				_567_v80 = _566_v79
				var _568_v81 _dafny.Map
				_ = _568_v81
				_568_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_567_v80, p1)
				_568_v81 = (_568_v81).Update(_567_v80, (_this).F11())
			}
			var _569_v82 _dafny.Set
			_ = _569_v82
			_569_v82 = _dafny.SetOf(false)
			_569_v82 = (Companion_Default___.Fm31(p1, (_this).F12(), (_this).F11(), globalState)).Difference(_569_v82)
			_463_v0 = _463_v0
		}
		r0 = (_this).F12()
		var _570_v83 _dafny.Map
		_ = _570_v83
		_570_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), p0)
		r1 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25((func() bool {
			if (_570_v83).Contains((_this).F12()) {
				return (_570_v83).Get((_this).F12()).(bool)
			}
			return p0
		})(), _dafny.IntOfUint32((Companion_Default___.Fm32(true, p1, (_this).F12(), (_this).F11(), globalState)).Cardinality()), globalState), _490_v26)
		r2 = _dafny.IntOfUint32((_490_v26).Cardinality())
		return r0, r1, r2
	}
}
func (_this *C2) F11() bool {
	{
		return _this._f11
	}
}
func (_this *C2) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f9  _dafny.MultiSet
	_f10 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f9 = _dafny.EmptyMultiSet
	_this._f10 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__(f9 _dafny.MultiSet, f10 bool) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C3) Fm19(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if !(false) {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(550), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-925), _dafny.IntOfInt64(484), (_dafny.SetOf(true)).Cardinality())).Cardinality(), _dafny.IntOfInt64(-771))).Cardinality()))))).Cardinality())
		} else {
			return (_dafny.Zero).Minus(_dafny.IntOfInt64(-312))
		}
	}
}
func (_this *C3) Fm20(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(483)
	}
}
func (_this *C3) M5(p0 _dafny.Int, p1 bool, globalState *GlobalState) (bool, bool, _dafny.CodePoint, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r2
		var r3 bool = false
		_ = r3
		var _571_v0 _dafny.Array
		_ = _571_v0
		var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
		_ = _nw93
		_571_v0 = _nw93
		var _572_v1 _dafny.Map
		_ = _572_v1
		_572_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _573_v2 _dafny.Array
		_ = _573_v2
		var _nwElement0_28 _dafny.Int = p0
		_ = _nwElement0_28
		var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(3))
		_ = _nw94
		(_nw94).ArraySet1(_nwElement0_28, 0)
		(_nw94).ArraySet1(p0, 1)
		(_nw94).ArraySet1((_572_v1).Cardinality(), 2)
		_573_v2 = _nw94
		var _574_v3 _dafny.Map
		_ = _574_v3
		_574_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_573_v2, _573_v2)
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_571_v0), 0))
		_ = _index79
		(_571_v0).ArraySet1((_574_v3).Merge(_574_v3), (_index79).Int())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_571_v0), 0))
		_ = _index80
		(_571_v0).ArraySet1(_574_v3, (_index80).Int())
		var _575_v5 _dafny.Sequence
		_ = _575_v5
		_575_v5 = _dafny.SeqOf((_this).F9())
		var _576_v6 D2
		_ = _576_v6
		_576_v6 = Companion_D2_.Create_DC3_(func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter21 := _dafny.Iterate((_575_v5).Elements()); ; {
				_compr_20, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _577_v4 _dafny.MultiSet
				_577_v4 = interface{}(_compr_20).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_575_v5, _577_v4) {
					_coll20.Add(_577_v4, (_dafny.MultiSetFromSeq(_dafny.SeqOf(p0, p0))).Cardinality())
				}
			}
			return _coll20.ToMap()
		}())
		var _source11 D2 = _576_v6
		_ = _source11
		if _source11.Is_DC4() {
			var _578___mcc_h0 _dafny.Int = _source11.Get_().(D2_DC4).Cf6
			_ = _578___mcc_h0
			var _579_cf6 _dafny.Int = _578___mcc_h0
			_ = _579_cf6
			var _580_v7 _dafny.Map
			_ = _580_v7
			_580_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1) && (p1), (p0).Minus(p0))
			_580_v7 = (_580_v7).Update(p1, _579_cf6)
			var _581_v8 _dafny.Sequence
			_ = _581_v8
			_581_v8 = _dafny.UnicodeSeqOfUtf8Bytes("lfvoqiom")
			var _582_v9 _dafny.Set
			_ = _582_v9
			_582_v9 = _dafny.SetOf(p0, p0, p0)
			var _583_v10 _dafny.Sequence
			_ = _583_v10
			_583_v10 = _dafny.SeqOf(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_581_v8).Cardinality())), p0, _dafny.IntOfInt64(30)), _582_v9)
			var _584_v11 _dafny.CodePoint
			_ = _584_v11
			_584_v11 = _dafny.CodePoint('k')
			var _585_v12 _dafny.Map
			_ = _585_v12
			_585_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _573_v2)
			var _586_v13 _dafny.Sequence
			_ = _586_v13
			_586_v13 = _dafny.SeqOf(_579_cf6)
			var _587_v14 _dafny.Sequence
			_ = _587_v14
			_587_v14 = _dafny.SeqOf((_this).F10(), p1)
			var _588_v15 D2
			_ = _588_v15
			_588_v15 = Companion_D2_.Create_DC5_(p0, (_this).F10(), p1, (_this).F10(), (_this).F10())
			var _589_v17 _dafny.Array
			_ = _589_v17
			var _nwElement0_29 bool = Companion_Default___.Fm21(_dafny.IntOfUint32((_583_v10).Cardinality()), _584_v11, Companion_Default___.Fm21((_dafny.Zero).Minus(((_585_v12).Update((_this).Fm19(_584_v11, (_586_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_587_v14).Cardinality()), _dafny.IntOfUint32((_586_v13).Cardinality()))).Uint32()).(_dafny.Int), globalState), _573_v2)).Cardinality()), _584_v11, p1, (_this).F10(), globalState), p1, globalState)
			_ = _nwElement0_29
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(27))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_29, 0)
			(_nw95).ArraySet1((_582_v9).IsSubsetOf(_582_v9), 1)
			(_nw95).ArraySet1(((_this).F9()).IsProperSubsetOf(_dafny.MultiSetFromSeq(_586_v13)), 2)
			(_nw95).ArraySet1((_this).F10(), 3)
			(_nw95).ArraySet1(p1, 4)
			(_nw95).ArraySet1(Companion_Default___.Fm21(_dafny.IntOfUint32((_581_v8).Cardinality()), _584_v11, (_this).F10(), false, globalState), 5)
			(_nw95).ArraySet1(p1, 6)
			(_nw95).ArraySet1((_this).F10(), 7)
			(_nw95).ArraySet1((func() bool {
				if false {
					return p1
				}
				return (_this).F10()
			})(), 8)
			(_nw95).ArraySet1(p1, 9)
			(_nw95).ArraySet1(p1, 10)
			(_nw95).ArraySet1((_588_v15).Dtor_cf11(), 11)
			(_nw95).ArraySet1((func() bool {
				if (_this).F10() {
					return true
				}
				return (_this).F10()
			})(), 12)
			(_nw95).ArraySet1(!((_this).F10()), 13)
			(_nw95).ArraySet1((_this).F10(), 14)
			(_nw95).ArraySet1(p1, 15)
			(_nw95).ArraySet1(false, 16)
			(_nw95).ArraySet1((_this).F10(), 17)
			(_nw95).ArraySet1(((_this).F10()) == (!((_this).F10())), 18)
			(_nw95).ArraySet1(p1, 19)
			(_nw95).ArraySet1((p1) && ((_this).F10()), 20)
			(_nw95).ArraySet1(p1, 21)
			(_nw95).ArraySet1(false, 22)
			(_nw95).ArraySet1((_dafny.SetOf(p0, _579_cf6)).IsProperSubsetOf(func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter22 := _dafny.Iterate((_586_v13).Elements()); ; {
					_compr_21, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _590_v16 _dafny.Int
					_590_v16 = interface{}(_compr_21).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_586_v13, _590_v16) {
						_coll21.Add(Companion_Default___.SafeModuloInt(_590_v16, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(994))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}(func(_591_i0 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(121)
						}))).Cardinality())))
					}
				}
				return _coll21.ToSet()
			}()), 23)
			(_nw95).ArraySet1(p1, 24)
			(_nw95).ArraySet1(_dafny.Companion_Sequence_.Equal(_586_v13, _586_v13), 25)
			(_nw95).ArraySet1(Companion_Default___.Fm21(_579_cf6, _584_v11, (func() bool {
				if (_572_v1).Contains(true) {
					return (_572_v1).Get(true).(bool)
				}
				return (_this).F10()
			})(), p1, globalState), 26)
			_589_v17 = _nw95
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_589_v17), 0))
			_ = _index81
			(_589_v17).ArraySet1(p1, (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_589_v17), 0))
			_ = _index82
			(_589_v17).ArraySet1((_this).F10(), (_index82).Int())
			_573_v2 = _573_v2
			_589_v17 = _589_v17
		} else if _source11.Is_DC5() {
			var _592___mcc_h1 _dafny.Int = _source11.Get_().(D2_DC5).Cf7
			_ = _592___mcc_h1
			var _593___mcc_h2 bool = _source11.Get_().(D2_DC5).Cf8
			_ = _593___mcc_h2
			var _594___mcc_h3 bool = _source11.Get_().(D2_DC5).Cf9
			_ = _594___mcc_h3
			var _595___mcc_h4 bool = _source11.Get_().(D2_DC5).Cf10
			_ = _595___mcc_h4
			var _596___mcc_h5 bool = _source11.Get_().(D2_DC5).Cf11
			_ = _596___mcc_h5
			var _597_cf11 bool = _596___mcc_h5
			_ = _597_cf11
			var _598_cf10 bool = _595___mcc_h4
			_ = _598_cf10
			var _599_cf9 bool = _594___mcc_h3
			_ = _599_cf9
			var _600_cf8 bool = _593___mcc_h2
			_ = _600_cf8
			var _601_cf7 _dafny.Int = _592___mcc_h1
			_ = _601_cf7
			_597_cf11 = (p0).Cmp((_601_cf7).Minus(_601_cf7)) != 0
			var _602_v18 _dafny.Sequence
			_ = _602_v18
			_602_v18 = _dafny.UnicodeSeqOfUtf8Bytes("vsovhibx")
			var _603_v19 _dafny.Sequence
			_ = _603_v19
			_603_v19 = _dafny.SeqOf(_dafny.IntOfInt64(270), _dafny.IntOfUint32((_602_v18).Cardinality()))
			var _604_v20 _dafny.CodePoint
			_ = _604_v20
			_604_v20 = _dafny.CodePoint('q')
			var _605_v21 _dafny.Sequence
			_ = _605_v21
			_605_v21 = _dafny.SeqOf(_600_cf8)
			r1 = (func() bool {
				if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_576_v6, p1)).Contains(Companion_Default___.Fm22(_603_v19, _601_cf7, _604_v20, globalState)) {
					return Companion_Default___.Fm21(p0, _604_v20, (_605_v21).Select((Companion_Default___.SafeIndex(_601_cf7, _dafny.IntOfUint32((_605_v21).Cardinality()))).Uint32()).(bool), _598_cf10, globalState)
				}
				return Companion_Default___.Fm21(_601_cf7, _604_v20, (_this).F10(), _598_cf10, globalState)
			})()
			var _606_v22 _dafny.Map
			_ = _606_v22
			_606_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_601_cf7, _573_v2)
			_606_v22 = ((_606_v22).Merge((_606_v22).Update(_dafny.IntOfInt64(963), _573_v2))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _573_v2))
			var _607_v23 _dafny.Map
			_ = _607_v23
			_607_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_598_cf10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_597_cf11, _dafny.IntOfInt64(-126)))
			var _608_v24 _dafny.Array
			_ = _608_v24
			var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw96
			_608_v24 = _nw96
			var _609_v25 *C1
			_ = _609_v25
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__((_dafny.Zero).Minus((_607_v23).Cardinality()), _608_v24)
			_609_v25 = _nw97
		} else {
			var _610___mcc_h6 _dafny.Map = _source11.Get_().(D2_DC3).Cf5
			_ = _610___mcc_h6
			var _611_cf5 _dafny.Map = _610___mcc_h6
			_ = _611_cf5
			var _612_v26 _dafny.Sequence
			_ = _612_v26
			_612_v26 = _dafny.UnicodeSeqOfUtf8Bytes("lnqpg")
			var _613_v27 _dafny.CodePoint
			_ = _613_v27
			_613_v27 = _dafny.CodePoint('b')
			var _614_v28 D4
			_ = _614_v28
			_614_v28 = Companion_D4_.Create_DC9_(_dafny.IntOfInt64(991), _613_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(119))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_615_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(141)
			})), p0, (_this).F10())
			var _616_v29 _dafny.Sequence
			_ = _616_v29
			_616_v29 = _dafny.SeqOf(p0, _dafny.IntOfInt64(-848))
			var _617_v30 _dafny.Sequence
			_ = _617_v30
			_617_v30 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_612_v26, Companion_Default___.Fm25((_614_v28).Dtor_cf20(), _dafny.IntOfUint32((_616_v29).Cardinality()), globalState))).Cardinality()))
			(globalState).F2 = (_616_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_616_v29).Cardinality()))).Uint32()).(_dafny.Int)
			(globalState).F2 = p0
			var _618_v31 _dafny.Array
			_ = _618_v31
			var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw98
			_618_v31 = _nw98
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_618_v31), 0))
			_ = _index83
			(_618_v31).ArraySet1(Companion_Default___.Fm25((_this).F10(), _dafny.IntOfInt64(444), globalState), (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_618_v31), 0))
			_ = _index84
			(_618_v31).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_619_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_620_i2 _dafny.Int) _dafny.CodePoint {
					return _619_v27
				}
			})(_613_v27))), (_index84).Int())
			r3 = p1
		}
		var _621_v32 _dafny.Array
		_ = _621_v32
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
		_ = _nw99
		_621_v32 = _nw99
		var _622_v33 _dafny.Map
		_ = _622_v33
		_622_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_621_v32, p1)
		r0 = (func() bool {
			if (_622_v33).Contains(_621_v32) {
				return (_622_v33).Get(_621_v32).(bool)
			}
			return (_this).F10()
		})()
		var _623_v34 _dafny.Array
		_ = _623_v34
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_11
		var _nw100 _dafny.Array
		_ = _nw100
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw100 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Map = func(_624_i3 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _dafny.UnicodeSeqOfUtf8Bytes("rvkrqon"))
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw100 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw100).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw100).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_623_v34 = _nw100
		var _625_v35 _dafny.CodePoint
		_ = _625_v35
		_625_v35 = _dafny.CodePoint('t')
		var _626_v36 _dafny.Sequence
		_ = _626_v36
		_626_v36 = _dafny.UnicodeSeqOfUtf8Bytes("aewdkgg")
		var _627_v37 _dafny.Map
		_ = _627_v37
		_627_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_625_v35, Companion_Default___.Fm25(false, _dafny.IntOfUint32((_626_v36).Cardinality()), globalState))
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_623_v34), 0))
		_ = _index85
		(_623_v34).ArraySet1(_627_v37, (_index85).Int())
		var _628_v38 _dafny.Map
		_ = _628_v38
		_628_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), p1, !((_this).F10()), (_this).F10(), p1)).Cardinality())))
		var _pat_let_tv15 = _627_v37
		_ = _pat_let_tv15
		var _pat_let_tv16 = _625_v35
		_ = _pat_let_tv16
		var _pat_let_tv17 = _626_v36
		_ = _pat_let_tv17
		var _pat_let_tv18 = _627_v37
		_ = _pat_let_tv18
		var _pat_let_tv19 = _627_v37
		_ = _pat_let_tv19
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_623_v34), 0))
		_ = _index86
		(_623_v34).ArraySet1(func(_source12 D2) _dafny.Map {
			if _source12.Is_DC4() {
				var _629___mcc_h7 _dafny.Int = _source12.Get_().(D2_DC4).Cf6
				_ = _629___mcc_h7
				var _630_cf6 _dafny.Int = _629___mcc_h7
				_ = _630_cf6
				return _pat_let_tv15
			} else if _source12.Is_DC5() {
				var _631___mcc_h8 _dafny.Int = _source12.Get_().(D2_DC5).Cf7
				_ = _631___mcc_h8
				var _632___mcc_h9 bool = _source12.Get_().(D2_DC5).Cf8
				_ = _632___mcc_h9
				var _633___mcc_h10 bool = _source12.Get_().(D2_DC5).Cf9
				_ = _633___mcc_h10
				var _634___mcc_h11 bool = _source12.Get_().(D2_DC5).Cf10
				_ = _634___mcc_h11
				var _635___mcc_h12 bool = _source12.Get_().(D2_DC5).Cf11
				_ = _635___mcc_h12
				var _636_cf11 bool = _635___mcc_h12
				_ = _636_cf11
				var _637_cf10 bool = _634___mcc_h11
				_ = _637_cf10
				var _638_cf9 bool = _633___mcc_h10
				_ = _638_cf9
				var _639_cf8 bool = _632___mcc_h9
				_ = _639_cf8
				var _640_cf7 _dafny.Int = _631___mcc_h8
				_ = _640_cf7
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv16, _pat_let_tv17)).Merge(_pat_let_tv18)
			} else {
				var _641___mcc_h13 _dafny.Map = _source12.Get_().(D2_DC3).Cf5
				_ = _641___mcc_h13
				var _642_cf5 _dafny.Map = _641___mcc_h13
				_ = _642_cf5
				return _pat_let_tv19
			}
		}(Companion_D2_.Create_DC3_(_628_v38)), (_index86).Int())
		var _643_v40 _dafny.Array
		_ = _643_v40
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_12
		var _nw101 _dafny.Array
		_ = _nw101
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw101 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_644_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_645_i4 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll22 = _dafny.NewMapBuilder()
						_ = _coll22
						for _iter23 := _dafny.Iterate((_dafny.SetOf(_644_p0)).Elements()); ; {
							_compr_22, _ok23 := _iter23()
							if !_ok23 {
								break
							}
							var _646_v39 _dafny.Int
							_646_v39 = interface{}(_compr_22).(_dafny.Int)
							if (_dafny.SetOf(_644_p0)).Contains(_646_v39) {
								_coll22.Add((_646_v39).Minus(_644_p0), (_this).F10())
							}
						}
						return _coll22.ToMap()
					}()
				}
			})(p0)
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
		_643_v40 = _nw101
		var _647_v41 *C2
		_ = _647_v41
		var _nw102 *C2 = New_C2_()
		_ = _nw102
		_nw102.Ctor__((_this).F10(), p0, _643_v40)
		_647_v41 = _nw102
		var _648_v42 _dafny.Array
		_ = _648_v42
		var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw103
		_648_v42 = _nw103
		_648_v42 = _648_v42
		r0 = p1
		r1 = (_647_v41).F11()
		var _649_v43 _dafny.Array
		_ = _649_v43
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_13
		var _nw104 _dafny.Array
		_ = _nw104
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw104 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Sequence = (func(_650_p0 _dafny.Int, _651_v41 *C2) func(_dafny.Int) _dafny.Sequence {
				return func(_652_i5 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_650_p0, _dafny.IntOfUint32((_dafny.SeqOf((_651_v41).F11(), (_this).F10())).Cardinality()), (_651_v41).F12())
				}
			})(p0, _647_v41)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw104 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw104).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw104).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_649_v43 = _nw104
		var _653_v44 _dafny.Map
		_ = _653_v44
		_653_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v43, (_647_v41).F12())
		var _654_v45 _dafny.Map
		_ = _654_v45
		_654_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_647_v41).F12())
		var _655_v46 _dafny.Sequence
		_ = _655_v46
		_655_v46 = _dafny.SeqOf(p1, (_647_v41).F11(), (_647_v41).F11(), (_this).F10(), (_647_v41).F11())
		var _656_v47 _dafny.Sequence
		_ = _656_v47
		_656_v47 = _dafny.SeqOf(_dafny.IntOfUint32((_655_v46).Cardinality()), p0)
		var _657_v48 D4
		_ = _657_v48
		_657_v48 = Companion_D4_.Create_DC9_((func() _dafny.Int {
			if (_653_v44).Contains(_649_v43) {
				return (_653_v44).Get(_649_v43).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_654_v45).Contains((_647_v41).F11()) {
					return (_654_v45).Get((_647_v41).F11()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(273)
			})()
		})(), _625_v35, _656_v47, p0, (_647_v41).F11())
		r2 = (_657_v48).Dtor_cf17()
		r3 = ((_this).Fm19(_625_v35, (_647_v41).F12(), globalState)).Cmp(_dafny.IntOfInt64(636)) <= 0
		return r0, r1, r2, r3
	}
}
func (_this *C3) F9() _dafny.MultiSet {
	{
		return _this._f9
	}
}
func (_this *C3) F10() bool {
	{
		return _this._f10
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f3 _dafny.Array
	_f7 bool
	_f8 _dafny.MultiSet
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f7 = false
	_this._f8 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F3() _dafny.Array {
	return _this._f3
}
func (_this *C4) Ctor__(f7 bool, f8 _dafny.MultiSet, f3 _dafny.Array) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f3 = f3
	}
}
func (_this *C4) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	{
		return ((func() _dafny.Set {
			if true {
				return _dafny.SetOf(_dafny.IntOfInt64(196), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yapel")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ruu")).Cardinality())))).Cardinality()))
			}
			return _dafny.SetOf(_dafny.IntOfInt64(532), _dafny.IntOfInt64(424))
		})()).Intersection(_dafny.SetOf(_dafny.IntOfInt64(566)))
	}
}
func (_this *C4) Fm14(p0 _dafny.Map, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("aljqjwu")
	}
}
func (_this *C4) M1(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var _658_v0 _dafny.Int
		_ = _658_v0
		_658_v0 = _dafny.IntOfInt64(671)
		var _659_v1 D1
		_ = _659_v1
		_659_v1 = Companion_D1_.Create_DC2_(_658_v0, p1, _658_v0)
		var _660_v2 _dafny.Array
		_ = _660_v2
		var _nwElement0_30 bool = (_this).F7()
		_ = _nwElement0_30
		var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(10))
		_ = _nw105
		(_nw105).ArraySet1(_nwElement0_30, 0)
		(_nw105).ArraySet1((_658_v0).Cmp(_658_v0) >= 0, 1)
		(_nw105).ArraySet1((_659_v1).Dtor_cf3(), 2)
		(_nw105).ArraySet1(!(p1), 3)
		(_nw105).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Cmp(_658_v0) > 0, 4)
		(_nw105).ArraySet1(p1, 5)
		(_nw105).ArraySet1((_659_v1).Dtor_cf3(), 6)
		(_nw105).ArraySet1((_658_v0).Cmp(_658_v0) <= 0, 7)
		(_nw105).ArraySet1((_this).F7(), 8)
		(_nw105).ArraySet1((_this).F7(), 9)
		_660_v2 = _nw105
		var _661_v3 _dafny.Array
		_ = _661_v3
		var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
		_ = _nw106
		_661_v3 = _nw106
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))
		_ = _index87
		(_661_v3).ArraySet1(_658_v0, (_index87).Int())
		var _662_v6 _dafny.Set
		_ = _662_v6
		_662_v6 = _dafny.SetOf(_658_v0, _dafny.IntOfUint32((p2).Cardinality()))
		var _663_v7 _dafny.Map
		_ = _663_v7
		_663_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll23 = _dafny.NewMapBuilder()
			_ = _coll23
			for _iter24 := _dafny.Iterate((_662_v6).Elements()); ; {
				_compr_23, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _664_v5 _dafny.Int
				_664_v5 = interface{}(_compr_23).(_dafny.Int)
				if (_662_v6).Contains(_664_v5) {
					_coll23.Add((_664_v5).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), (_this).F7())).Cardinality()), true)
				}
			}
			return _coll23.ToMap()
		}(), (_this).F7())
		var _665_v8 _dafny.Map
		_ = _665_v8
		_665_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, p1)
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))
		_ = _index88
		var _rhs72 _dafny.Array = _660_v2
		_ = _rhs72
		var _rhs73 _dafny.Int = _658_v0
		_ = _rhs73
		var _rhs74 _dafny.Int = ((func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter25 := _dafny.Iterate((_663_v7).Keys().Elements()); ; {
				_compr_24, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _666_v4 _dafny.Map
				_666_v4 = interface{}(_compr_24).(_dafny.Map)
				if (_663_v7).Contains(_666_v4) {
					_coll24.Add(_666_v4, _658_v0)
				}
			}
			return _coll24.ToMap()
		}()).Update(_665_v8, (_658_v0).Times(_658_v0))).Cardinality()
		_ = _rhs74
		var _lhs46 _dafny.Array = _661_v3
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))
		_ = _lhs47
		var _lhs48 *GlobalState = globalState
		_ = _lhs48
		_660_v2 = _rhs72
		(_lhs46).ArraySet1(_rhs73, (_lhs47).Int())
		_lhs48.F2 = _rhs74
		var _667_i0 _dafny.Int
		_ = _667_i0
		_667_i0 = _dafny.Zero
		{
			for p1 {
				{
					if (_667_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_667_i0 = (_667_i0).Plus(_dafny.One)
					var _668_v9 _dafny.Map
					_ = _668_v9
					_668_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm15(!(p1), true, globalState), _658_v0)
					var _669_v10 D2
					_ = _669_v10
					_669_v10 = Companion_D2_.Create_DC3_(_668_v9)
					var _670_v11 _dafny.Map
					_ = _670_v11
					_670_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_658_v0), _dafny.IntOfInt64(226))
					var _671_v12 _dafny.Sequence
					_ = _671_v12
					_671_v12 = _dafny.SeqOf(_660_v2)
					var _672_v13 _dafny.Array
					_ = _672_v13
					var _nwElement0_31 D2 = Companion_D2_.Create_DC3_(_668_v9)
					_ = _nwElement0_31
					var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
					_ = _nw107
					(_nw107).ArraySet1(_nwElement0_31, 0)
					(_nw107).ArraySet1(_669_v10, 1)
					(_nw107).ArraySet1(Companion_D2_.Create_DC3_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}((func(_673_v3 _dafny.Array, _674_p2 _dafny.Sequence, _675_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_676_i1 _dafny.Int) _dafny.Int {
							return ((_dafny.MultiSetOf((_673_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_673_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-947))).Update(_dafny.IntOfUint32((_674_p2).Cardinality()), Companion_Default___.Abs(_675_v0))).Cardinality()
						}
					})(_661_v3, p2, _658_v0)))).Cardinality()))).Update((_this).F8(), (func() _dafny.Int {
						if (_670_v11).Contains(_dafny.IntOfUint32((_671_v12).Cardinality())) {
							return (_670_v11).Get(_dafny.IntOfUint32((_671_v12).Cardinality())).(_dafny.Int)
						}
						return _658_v0
					})())), 2)
					(_nw107).ArraySet1(_669_v10, 3)
					(_nw107).ArraySet1(_669_v10, 4)
					(_nw107).ArraySet1(_669_v10, 5)
					(_nw107).ArraySet1(_669_v10, 6)
					(_nw107).ArraySet1(_669_v10, 7)
					_672_v13 = _nw107
					var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_672_v13), 0))
					_ = _index89
					(_672_v13).ArraySet1(_669_v10, (_index89).Int())
					var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_672_v13), 0))
					_ = _index90
					(_672_v13).ArraySet1(_669_v10, (_index90).Int())
					var _677_v14 _dafny.Sequence
					_ = _677_v14
					_677_v14 = _dafny.UnicodeSeqOfUtf8Bytes("c")
					_677_v14 = _dafny.Companion_Sequence_.Concatenate(_677_v14, _dafny.Companion_Sequence_.Concatenate(p2, p2))
					if (_this).F7() {
						var _678_v15 bool
						_ = _678_v15
						_678_v15 = false
						_678_v15 = ((_662_v6).Intersection(func() _dafny.Set {
							var _coll25 = _dafny.NewBuilder()
							_ = _coll25
							for _iter26 := _dafny.Iterate((_662_v6).Elements()); ; {
								_compr_25, _ok26 := _iter26()
								if !_ok26 {
									break
								}
								var _679_v16 _dafny.Int
								_679_v16 = interface{}(_compr_25).(_dafny.Int)
								if (_662_v6).Contains(_679_v16) {
									_coll25.Add((_679_v16).Minus(_dafny.IntOfInt64(713)))
								}
							}
							return _coll25.ToSet()
						}())).IsDisjointFrom(_662_v6)
						_678_v15 = p1
						var _680_v17 _dafny.Sequence
						_ = _680_v17
						_680_v17 = _dafny.SeqOf(p1, p1)
						var _681_v18 _dafny.Map
						_ = _681_v18
						_681_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_658_v0), _680_v17)
						_681_v18 = (_681_v18).Update((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F7(), _678_v15), _680_v17))
						_678_v15 = _678_v15
						var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))
						_ = _index91
						(_661_v3).ArraySet1(_658_v0, (_index91).Int())
					} else {
						var _682_v19 D3
						_ = _682_v19
						_682_v19 = Companion_D3_.Create_DC6_(_dafny.CodePoint('l'))
						_658_v0 = (p0).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm16((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), (_682_v19).Dtor_cf12(), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), globalState)).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
						var _683_v20 _dafny.Sequence
						_ = _683_v20
						_683_v20 = _dafny.SeqOf(p2, _dafny.SeqOf(_dafny.CodePoint('r')))
						(globalState).F2 = (Companion_Default___.Fm17(_dafny.SeqOf(_677_v14), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), globalState)).Plus(Companion_Default___.Fm17(_683_v20, _658_v0, (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), _658_v0, globalState))
						var _684_v21 _dafny.Array
						_ = _684_v21
						var _nw108 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(22))
						_ = _nw108
						_684_v21 = _nw108
						_684_v21 = _684_v21
						var _685_v22 _dafny.Map
						_ = _685_v22
						_685_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
						var _686_v23 _dafny.Sequence
						_ = _686_v23
						_686_v23 = _dafny.SeqOf((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), ((_685_v22).Cardinality()).Times((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus(_dafny.IntOfInt64(-212)), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(571), _658_v0))
						var _687_v24 _dafny.Set
						_ = _687_v24
						_687_v24 = _dafny.SetOf(_660_v2)
						var _rhs75 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_677_v14).Cardinality())), _658_v0))
						_ = _rhs75
						var _rhs76 _dafny.Int = _658_v0
						_ = _rhs76
						var _rhs77 D1 = Companion_Default___.Fm18((_687_v24).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()), globalState)
						_ = _rhs77
						var _rhs78 _dafny.Int = (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)
						_ = _rhs78
						var _rhs79 _dafny.Sequence = p0
						_ = _rhs79
						var _lhs49 *GlobalState = globalState
						_ = _lhs49
						var _lhs50 *GlobalState = globalState
						_ = _lhs50
						_658_v0 = _rhs75
						_lhs49.F2 = _rhs76
						_659_v1 = _rhs77
						_lhs50.F2 = _rhs78
						_686_v23 = _rhs79
						var _688_v25 _dafny.Array
						_ = _688_v25
						var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
						_ = _nw109
						_688_v25 = _nw109
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_688_v25), 0))
						_ = _index92
						(_688_v25).ArraySet1(_661_v3, (_index92).Int())
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_688_v25), 0))
						_ = _index93
						(_688_v25).ArraySet1(_661_v3, (_index93).Int())
					}
					var _689_v26 _dafny.Set
					_ = _689_v26
					_689_v26 = _dafny.SetOf(true)
					var _690_v27 _dafny.Map
					_ = _690_v27
					_690_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_689_v26, _dafny.IntOfUint32((_677_v14).Cardinality()))
					_690_v27 = (_690_v27).Update(_689_v26, _dafny.IntOfInt64(680))
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _691_i2 _dafny.Int
		_ = _691_i2
		_691_i2 = _dafny.Zero
		{
			for p1 {
				{
					if (_691_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_691_i2 = (_691_i2).Plus(_dafny.One)
					(globalState).F2 = (_dafny.Zero).Minus((((func() _dafny.Map {
						if p1 {
							return _665_v8
						}
						return (_665_v8).Update((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), (_this).F7())
					})()).Merge(_665_v8)).Cardinality())
					var _692_v28 *C0
					_ = _692_v28
					var _nw110 *C0 = New_C0_()
					_ = _nw110
					_nw110.Ctor__((func() _dafny.Int {
						if p1 {
							return Companion_Default___.Fm17(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-339))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg38 _dafny.Int) interface{} {
									return coer38(arg38)
								}
							}((func(_693_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_694_i3 _dafny.Int) _dafny.Sequence {
									return _693_p2
								}
							})(p2))), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), _658_v0, globalState)
						}
						return (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)
					})())
					_692_v28 = _nw110
					var _695_v29 *C3
					_ = _695_v29
					var _nw111 *C3 = New_C3_()
					_ = _nw111
					_nw111.Ctor__((_this).F8(), (_this).F7())
					_695_v29 = _nw111
					(globalState).F2 = _658_v0
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _696_v30 bool
		_ = _696_v30
		_696_v30 = false
		_696_v30 = p1
		var _697_i4 _dafny.Int
		_ = _697_i4
		_697_i4 = _dafny.Zero
		{
			for !_dafny.Companion_Sequence_.Equal(p0, p0) {
				{
					if (_697_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_697_i4 = (_697_i4).Plus(_dafny.One)
					var _698_v31 _dafny.CodePoint
					_ = _698_v31
					_698_v31 = _dafny.CodePoint('x')
					var _rhs80 _dafny.CodePoint = (func() _dafny.CodePoint {
						if _696_v30 {
							return _698_v31
						}
						return (p2).Select((Companion_Default___.SafeIndex(_658_v0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.CodePoint)
					})()
					_ = _rhs80
					var _rhs81 bool = (_658_v0).Cmp(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_658_v0), _658_v0)) <= 0
					_ = _rhs81
					var _rhs82 bool = (((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)).Times(_658_v0)).Cmp((func() _dafny.Int {
						if p1 {
							return _658_v0
						}
						return _658_v0
					})()) == 0
					_ = _rhs82
					r1 = _rhs80
					_696_v30 = _rhs81
					_696_v30 = _rhs82
					var _699_v32 _dafny.Array
					_ = _699_v32
					var _len0_14 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_14
					var _nw112 _dafny.Array
					_ = _nw112
					if _len0_14.Cmp(_dafny.Zero) == 0 {
						_nw112 = _dafny.NewArray(_len0_14)
					} else {
						var _init14 func(_dafny.Int) _dafny.Set = (func(_700_v30 bool, _701_p1 bool) func(_dafny.Int) _dafny.Set {
							return func(_702_i5 _dafny.Int) _dafny.Set {
								return (_dafny.SetOf(_700_v30, !((_this).F7()), _701_p1)).Difference(_dafny.SetOf(_701_p1))
							}
						})(_696_v30, p1)
						_ = _init14
						var _element0_14 = _init14(_dafny.Zero)
						_ = _element0_14
						_nw112 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
						(_nw112).ArraySet1(_element0_14, 0)
						var _nativeLen0_14 = (_len0_14).Int()
						_ = _nativeLen0_14
						for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
							(_nw112).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
						}
					}
					_699_v32 = _nw112
					var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_699_v32), 0))
					_ = _index94
					(_699_v32).ArraySet1(_dafny.SetOf(true, p1, (_this).F7(), !(_696_v30)), (_index94).Int())
					var _703_v33 _dafny.Set
					_ = _703_v33
					_703_v33 = _dafny.SetOf((_this).F7())
					var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_699_v32), 0))
					_ = _index95
					(_699_v32).ArraySet1(_703_v33, (_index95).Int())
					_696_v30 = ((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)).Cmp(_658_v0) > 0
					var _704_v34 D2
					_ = _704_v34
					_704_v34 = Companion_D2_.Create_DC4_(Companion_Default___.SafeModuloInt(_658_v0, _658_v0))
					var _source13 D2 = _704_v34
					_ = _source13
					if _source13.Is_DC4() {
						var _705___mcc_h0 _dafny.Int = _source13.Get_().(D2_DC4).Cf6
						_ = _705___mcc_h0
						var _706_cf6 _dafny.Int = _705___mcc_h0
						_ = _706_cf6
						_698_v31 = _698_v31
						_661_v3 = _661_v3
						var _707_v35 _dafny.Map
						_ = _707_v35
						_707_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(169))
						var _708_v36 _dafny.Map
						_ = _708_v36
						_708_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
							if (_707_v35).Contains(_658_v0) {
								return (_707_v35).Get(_658_v0).(_dafny.Int)
							}
							return Companion_Default___.Fm24(_662_v6, globalState)
						})()).Plus((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)), (_706_cf6).Times(_706_cf6))
						var _709_v37 _dafny.MultiSet
						_ = _709_v37
						_709_v37 = _dafny.MultiSetOf((_this).Fm13(_658_v0, _dafny.IntOfInt64(495), p1, (_dafny.Zero).Minus(_706_cf6), globalState))
						_707_v35 = (_707_v35).Update(_706_cf6, (func() _dafny.Int {
							if (_709_v37).Contains(_dafny.SetOf(_706_cf6)) {
								return (_709_v37).Multiplicity(_dafny.SetOf(_706_cf6))
							}
							return _dafny.IntOfUint32((p2).Cardinality())
						})())
						_696_v30 = Companion_Default___.Fm23((func() _dafny.Sequence {
							if !(p1) {
								return p2
							}
							return p2
						})(), _658_v0, _706_cf6, _696_v30, globalState)
					} else if _source13.Is_DC5() {
						var _710___mcc_h1 _dafny.Int = _source13.Get_().(D2_DC5).Cf7
						_ = _710___mcc_h1
						var _711___mcc_h2 bool = _source13.Get_().(D2_DC5).Cf8
						_ = _711___mcc_h2
						var _712___mcc_h3 bool = _source13.Get_().(D2_DC5).Cf9
						_ = _712___mcc_h3
						var _713___mcc_h4 bool = _source13.Get_().(D2_DC5).Cf10
						_ = _713___mcc_h4
						var _714___mcc_h5 bool = _source13.Get_().(D2_DC5).Cf11
						_ = _714___mcc_h5
						var _715_cf11 bool = _714___mcc_h5
						_ = _715_cf11
						var _716_cf10 bool = _713___mcc_h4
						_ = _716_cf10
						var _717_cf9 bool = _712___mcc_h3
						_ = _717_cf9
						var _718_cf8 bool = _711___mcc_h2
						_ = _718_cf8
						var _719_cf7 _dafny.Int = _710___mcc_h1
						_ = _719_cf7
						var _720_v38 _dafny.Array
						_ = _720_v38
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_15
						var _nw113 _dafny.Array
						_ = _nw113
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw113 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) _dafny.CodePoint = (func(_721_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_722_i6 _dafny.Int) _dafny.CodePoint {
									return _721_v31
								}
							})(_698_v31)
							_ = _init15
							var _element0_15 = _init15(_dafny.Zero)
							_ = _element0_15
							_nw113 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
							(_nw113).ArraySet1CodePoint(_element0_15, 0)
							var _nativeLen0_15 = (_len0_15).Int()
							_ = _nativeLen0_15
							for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
								(_nw113).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
							}
						}
						_720_v38 = _nw113
						var _723_v39 _dafny.Sequence
						_ = _723_v39
						_723_v39 = _dafny.SeqOf(_720_v38, _720_v38)
						_658_v0 = ((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfUint32((_723_v39).Cardinality()))
						var _724_v40 *C0
						_ = _724_v40
						var _nw114 *C0 = New_C0_()
						_ = _nw114
						_nw114.Ctor__(_719_cf7)
						_724_v40 = _nw114
						var _725_v41 _dafny.Sequence
						_ = _725_v41
						_725_v41 = _dafny.SeqOf(_718_cf8)
						var _726_v42 _dafny.MultiSet
						_ = _726_v42
						_726_v42 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_725_v41, _725_v41), _dafny.Companion_Sequence_.Update(_725_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.IntOfUint32((_725_v41).Cardinality()))).Uint32(), !(true)), _dafny.SeqOf(_718_cf8), _725_v41)
						var _727_v43 _dafny.Map
						_ = _727_v43
						_727_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_716_cf10, Companion_Default___.Fm33(_716_cf10, _dafny.SeqOf((_dafny.Zero).Minus((_724_v40).F13())), _717_cf9, _698_v31, globalState))
						_726_v42 = (func() _dafny.MultiSet {
							if (_727_v43).Contains(!(_696_v30)) {
								return (_727_v43).Get(!(_696_v30)).(_dafny.MultiSet)
							}
							return _726_v42
						})()
						var _728_v44 _dafny.Array
						_ = _728_v44
						var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_16
						var _nw115 _dafny.Array
						_ = _nw115
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw115 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) _dafny.Sequence = (func(_729_cf11 bool) func(_dafny.Int) _dafny.Sequence {
								return func(_730_i7 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_729_cf11, _729_cf11)
								}
							})(_715_cf11)
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw115 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw115).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw115).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_728_v44 = _nw115
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_728_v44), 0))
						_ = _index96
						(_728_v44).ArraySet1(_725_v41, (_index96).Int())
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_728_v44), 0))
						_ = _index97
						(_728_v44).ArraySet1(_725_v41, (_index97).Int())
					} else {
						var _731___mcc_h6 _dafny.Map = _source13.Get_().(D2_DC3).Cf5
						_ = _731___mcc_h6
						var _732_cf5 _dafny.Map = _731___mcc_h6
						_ = _732_cf5
						var _733_v45 *C2
						_ = _733_v45
						var _nw116 *C2 = New_C2_()
						_ = _nw116
						_nw116.Ctor__((_this).F7(), _dafny.IntOfInt64(296), (func() _dafny.Array {
							if p1 {
								return (_this).F3()
							}
							return (_this).F3()
						})())
						_733_v45 = _nw116
						var _734_v46 _dafny.MultiSet
						_ = _734_v46
						_734_v46 = _dafny.MultiSetOf(p1, !((_733_v45).F11()), true)
						var _735_v47 _dafny.Map
						_ = _735_v47
						_735_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _661_v3)
						var _736_v48 _dafny.Map
						_ = _736_v48
						_736_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), _661_v3)
						var _737_v49 _dafny.Array
						_ = _737_v49
						var _nwElement0_32 _dafny.Array = _661_v3
						_ = _nwElement0_32
						var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(18))
						_ = _nw117
						(_nw117).ArraySet1(_nwElement0_32, 0)
						(_nw117).ArraySet1(_661_v3, 1)
						(_nw117).ArraySet1((func() _dafny.Array {
							if (_735_v47).Contains((_733_v45).F11()) {
								return (_735_v47).Get((_733_v45).F11()).(_dafny.Array)
							}
							return _661_v3
						})(), 2)
						(_nw117).ArraySet1(_661_v3, 3)
						(_nw117).ArraySet1(_661_v3, 4)
						(_nw117).ArraySet1(_661_v3, 5)
						(_nw117).ArraySet1(_661_v3, 6)
						(_nw117).ArraySet1(_661_v3, 7)
						(_nw117).ArraySet1(_661_v3, 8)
						(_nw117).ArraySet1(_661_v3, 9)
						(_nw117).ArraySet1(_661_v3, 10)
						(_nw117).ArraySet1(_661_v3, 11)
						(_nw117).ArraySet1(_661_v3, 12)
						(_nw117).ArraySet1(_661_v3, 13)
						(_nw117).ArraySet1(_661_v3, 14)
						(_nw117).ArraySet1(_661_v3, 15)
						(_nw117).ArraySet1(_661_v3, 16)
						(_nw117).ArraySet1((func() _dafny.Array {
							if (_736_v48).Contains((_dafny.Zero).Minus((_733_v45).F12())) {
								return (_736_v48).Get((_dafny.Zero).Minus((_733_v45).F12())).(_dafny.Array)
							}
							return _661_v3
						})(), 17)
						_737_v49 = _nw117
						var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_737_v49), 0))
						_ = _index98
						(_737_v49).ArraySet1(_661_v3, (_index98).Int())
						var _738_v50 _dafny.Map
						_ = _738_v50
						_738_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_dafny.MultiSetOf(p1)).Intersection(_734_v46))
						var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_737_v49), 0))
						_ = _index99
						var _rhs83 _dafny.MultiSet = (func() _dafny.MultiSet {
							if (_738_v50).Contains(p1) {
								return (_738_v50).Get(p1).(_dafny.MultiSet)
							}
							return _734_v46
						})()
						_ = _rhs83
						var _rhs84 bool = p1
						_ = _rhs84
						var _rhs85 _dafny.Array = (func() _dafny.Array {
							if (_736_v48).Contains(((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)).Times((_733_v45).F12())) {
								return (_736_v48).Get(((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int)).Times((_733_v45).F12())).(_dafny.Array)
							}
							return _661_v3
						})()
						_ = _rhs85
						var _lhs51 _dafny.Array = _737_v49
						_ = _lhs51
						var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_737_v49), 0))
						_ = _lhs52
						_734_v46 = _rhs83
						_696_v30 = _rhs84
						(_lhs51).ArraySet1(_rhs85, (_lhs52).Int())
						var _739_v51 _dafny.Map
						_ = _739_v51
						_739_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v0, (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int))
						var _740_v52 _dafny.Map
						_ = _740_v52
						_740_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v51, ((_this).F8()).Intersection(_dafny.MultiSetOf((_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int))))
						_740_v52 = (_740_v52).Update(((_739_v51).Update((_733_v45).F12(), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int))).Update(((_this).F8()).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf((_this).F7())).Cardinality())), (_this).F8())
						var _rhs86 bool = true
						_ = _rhs86
						var _rhs87 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_733_v45).F12(), (_661_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_661_v3), 0))).Int()).(_dafny.Int))).Times(_658_v0))
						_ = _rhs87
						var _rhs88 _dafny.Int = (_733_v45).F12()
						_ = _rhs88
						var _rhs89 bool = (_this).F7()
						_ = _rhs89
						var _rhs90 bool = (Companion_Default___.SafeModuloInt(_658_v0, _658_v0)).Cmp((_733_v45).F12()) > 0
						_ = _rhs90
						var _lhs53 *GlobalState = globalState
						_ = _lhs53
						_696_v30 = _rhs86
						r0 = _rhs87
						_lhs53.F2 = _rhs88
						_696_v30 = _rhs89
						_696_v30 = _rhs90
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		_696_v30 = ((Companion_D2_.Create_DC5_(_658_v0, _696_v30, (_this).F7(), !(_696_v30), _696_v30)).Dtor_cf7()).Cmp(_dafny.IntOfUint32((p2).Cardinality())) == 0
		r0 = _658_v0
		var _741_v53 _dafny.CodePoint
		_ = _741_v53
		_741_v53 = _dafny.CodePoint('u')
		r1 = _741_v53
		return r0, r1
	}
}
func (_this *C4) F7() bool {
	{
		return _this._f7
	}
}
func (_this *C4) F8() _dafny.MultiSet {
	{
		return _this._f8
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f6 _dafny.CodePoint
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f6 = _dafny.CodePoint('D')
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f6 _dafny.CodePoint) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, p1 bool, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Int {
			if false {
				return _dafny.IntOfInt64(976)
			}
			return _dafny.IntOfInt64(852)
		})()).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.MultiSetOf((Companion_D1_.Create_DC2_(_dafny.IntOfInt64(840), true, _dafny.IntOfInt64(-480))).Dtor_cf3(), true, false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, false))).Cardinality()))).Cardinality()) <= 0
	}
}
func (_this *C5) M4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _742_v0 bool
		_ = _742_v0
		_742_v0 = false
		if !(_742_v0) {
			var _743_v1 _dafny.Sequence
			_ = _743_v1
			_743_v1 = _dafny.UnicodeSeqOfUtf8Bytes("otnenv")
			_742_v0 = !_dafny.Companion_Sequence_.Equal(_743_v1, _dafny.UnicodeSeqOfUtf8Bytes("sypxlsr"))
			var _744_v2 _dafny.Array
			_ = _744_v2
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
			_ = _nw118
			_744_v2 = _nw118
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_744_v2), 0))
			_ = _index100
			(_744_v2).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), (_index100).Int())
			var _745_v3 _dafny.Map
			_ = _745_v3
			_745_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _742_v0)
			var _746_v4 _dafny.MultiSet
			_ = _746_v4
			_746_v4 = _dafny.MultiSetOf(p0, p0)
			var _747_v5 _dafny.Map
			_ = _747_v5
			_747_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_746_v4, p1)
			var _748_v6 D2
			_ = _748_v6
			_748_v6 = Companion_D2_.Create_DC3_(_747_v5)
			var _749_v7 _dafny.Sequence
			_ = _749_v7
			_749_v7 = _dafny.SeqOf(((_748_v6).Dtor_cf5()).Cardinality(), p0)
			var _750_v8 _dafny.Map
			_ = _750_v8
			_750_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((p0).Times((_745_v3).Cardinality())), _dafny.IntOfUint32((_749_v7).Cardinality()))
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_744_v2), 0))
			_ = _index101
			(_744_v2).ArraySet1(_750_v8, (_index101).Int())
			var _751_v9 _dafny.Map
			_ = _751_v9
			_751_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_742_v0), _dafny.IntOfInt64(-76))
			var _752_v10 _dafny.MultiSet
			_ = _752_v10
			_752_v10 = _dafny.MultiSetOf(!(_742_v0))
			_751_v9 = (_751_v9).Update(!(_742_v0), (_752_v10).Cardinality())
			var _753_v11 _dafny.Array
			_ = _753_v11
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_17
			var _nw119 _dafny.Array
			_ = _nw119
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw119 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Set = (func(_754_v3 _dafny.Map, _755_p1 _dafny.Int, _756_v0 bool) func(_dafny.Int) _dafny.Set {
					return func(_757_i0 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(((_754_v3).Update(_755_p1, _756_v0)).Cardinality())
					}
				})(_745_v3, p1, _742_v0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw119 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw119).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw119).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_753_v11 = _nw119
			var _758_v12 _dafny.Array
			_ = _758_v12
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_18
			var _nw120 _dafny.Array
			_ = _nw120
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw120 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = (func(_759_v0 bool) func(_dafny.Int) bool {
					return func(_760_i1 _dafny.Int) bool {
						return _759_v0
					}
				})(_742_v0)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw120 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw120).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw120).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_758_v12 = _nw120
			var _rhs91 bool = !((func() bool {
				if !(_742_v0) {
					return _742_v0
				}
				return _742_v0
			})())
			_ = _rhs91
			var _rhs92 _dafny.Array = _753_v11
			_ = _rhs92
			var _rhs93 bool = !(_742_v0) || (_742_v0)
			_ = _rhs93
			var _rhs94 bool = _742_v0
			_ = _rhs94
			var _rhs95 _dafny.Array = _758_v12
			_ = _rhs95
			_742_v0 = _rhs91
			_753_v11 = _rhs92
			_742_v0 = _rhs93
			_742_v0 = _rhs94
			_758_v12 = _rhs95
			_743_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_743_v1, Companion_Default___.Fm8(p1, globalState)), _743_v1)
		} else {
			_742_v0 = _742_v0
			var _761_v13 _dafny.Map
			_ = _761_v13
			_761_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _dafny.UnicodeSeqOfUtf8Bytes("uqv"))
			var _762_v14 _dafny.Sequence
			_ = _762_v14
			_762_v14 = _dafny.UnicodeSeqOfUtf8Bytes("keousjicc")
			var _763_v15 _dafny.Sequence
			_ = _763_v15
			_763_v15 = _dafny.SeqOf((_dafny.Zero).Minus((_761_v13).Cardinality()), _dafny.IntOfUint32((_762_v14).Cardinality()), p0, p0, (_dafny.Zero).Minus(p1))
			var _764_v16 _dafny.Map
			_ = _764_v16
			_764_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if false {
					return p0
				}
				return p0
			})(), _763_v15)
			var _765_v17 _dafny.Array
			_ = _765_v17
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_19
			var _nw121 _dafny.Array
			_ = _nw121
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw121 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) D1 = (func(_766_p0 _dafny.Int, _767_v0 bool, _768_p1 _dafny.Int) func(_dafny.Int) D1 {
					return func(_769_i2 _dafny.Int) D1 {
						return Companion_D1_.Create_DC2_(_766_p0, _767_v0, _768_p1)
					}
				})(p0, _742_v0, p1)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw121 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw121).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw121).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_765_v17 = _nw121
			var _770_v18 D1
			_ = _770_v18
			_770_v18 = Companion_D1_.Create_DC2_(p0, true, p1)
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_765_v17), 0))
			_ = _index102
			(_765_v17).ArraySet1(_770_v18, (_index102).Int())
			var _771_v19 _dafny.Map
			_ = _771_v19
			_771_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _762_v14)
			var _772_v20 _dafny.Set
			_ = _772_v20
			_772_v20 = _dafny.SetOf((_this).Fm7(p0, _742_v0, _771_v19, globalState), false)
			var _773_v21 _dafny.Map
			_ = _773_v21
			_773_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, p1)
			var _774_v22 _dafny.Map
			_ = _774_v22
			_774_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v21, _dafny.UnicodeSeqOfUtf8Bytes("grdlrs"))
			var _775_v23 _dafny.Array
			_ = _775_v23
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw122
			_775_v23 = _nw122
			var _776_v24 _dafny.Map
			_ = _776_v24
			_776_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_775_v23, _742_v0)
			var _777_v25 _dafny.Array
			_ = _777_v25
			var _nwElement0_33 _dafny.Int = (_dafny.Zero).Minus(p1)
			_ = _nwElement0_33
			var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(25))
			_ = _nw123
			(_nw123).ArraySet1(_nwElement0_33, 0)
			(_nw123).ArraySet1(p1, 1)
			(_nw123).ArraySet1(p0, 2)
			(_nw123).ArraySet1(p0, 3)
			(_nw123).ArraySet1(p1, 4)
			(_nw123).ArraySet1(p1, 5)
			(_nw123).ArraySet1((_dafny.Zero).Minus((_772_v20).Cardinality()), 6)
			(_nw123).ArraySet1(_dafny.IntOfInt64(519), 7)
			(_nw123).ArraySet1(_dafny.IntOfInt64(118), 8)
			(_nw123).ArraySet1(p0, 9)
			(_nw123).ArraySet1((_763_v15).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_763_v15).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw123).ArraySet1(p1, 11)
			(_nw123).ArraySet1(p0, 12)
			(_nw123).ArraySet1(Companion_Default___.Fm9((_dafny.Zero).Minus(p1), globalState), 13)
			(_nw123).ArraySet1(p0, 14)
			(_nw123).ArraySet1(p0, 15)
			(_nw123).ArraySet1((p1).Times((_dafny.Zero).Minus(p1)), 16)
			(_nw123).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_774_v22).Contains(_773_v21) {
					return (_774_v22).Get(_773_v21).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}(func(_778_i3 _dafny.Int) _dafny.CodePoint {
					return (_this).F6()
				}))
			})()).Cardinality()), 17)
			(_nw123).ArraySet1((func() _dafny.Int {
				if (_773_v21).Contains(_742_v0) {
					return (_773_v21).Get(_742_v0).(_dafny.Int)
				}
				return p1
			})(), 18)
			(_nw123).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm9(p0, globalState), _dafny.IntOfUint32((_762_v14).Cardinality())), 19)
			(_nw123).ArraySet1(p1, 20)
			(_nw123).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0), 21)
			(_nw123).ArraySet1(((_776_v24).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_775_v23, _742_v0))).Cardinality(), 22)
			(_nw123).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngqqsqvhs")).Cardinality()), 23)
			(_nw123).ArraySet1(p1, 24)
			_777_v25 = _nw123
			var _779_v26 _dafny.Sequence
			_ = _779_v26
			_779_v26 = _dafny.SeqOf(_742_v0, _742_v0, _742_v0)
			var _780_v27 _dafny.Sequence
			_ = _780_v27
			_780_v27 = _dafny.SeqOf(Companion_Default___.Fm10(_742_v0, globalState), _779_v26)
			var _781_v28 _dafny.Sequence
			_ = _781_v28
			_781_v28 = (_780_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.IntOfUint32((_780_v27).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
			_ = _index103
			(_777_v25).ArraySet1(_dafny.IntOfUint32((_781_v28).Cardinality()), (_index103).Int())
			var _782_v29 _dafny.MultiSet
			_ = _782_v29
			_782_v29 = _dafny.MultiSetOf(p0, (_dafny.IntOfUint32((_763_v15).Cardinality())).Times(p0))
			var _783_v30 _dafny.Map
			_ = _783_v30
			_783_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_773_v21).Cardinality(), p1)
			var _784_v31 _dafny.MultiSet
			_ = _784_v31
			_784_v31 = _dafny.MultiSetOf(Companion_Default___.Fm12(_742_v0, p0, _772_v20, _783_v30, globalState))
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_765_v17), 0))
			_ = _index104
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
			_ = _index105
			var _rhs96 _dafny.Map = (_764_v16).Merge(_764_v16)
			_ = _rhs96
			var _rhs97 D1 = _770_v18
			_ = _rhs97
			var _rhs98 _dafny.Int = p1
			_ = _rhs98
			var _rhs99 _dafny.MultiSet = Companion_Default___.Fm11(_742_v0, _742_v0, Companion_D2_.Create_DC5_(p0, _742_v0, _742_v0, _742_v0, _742_v0), globalState)
			_ = _rhs99
			var _rhs100 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Int {
				if (_784_v31).Contains(_763_v15) {
					return (_784_v31).Multiplicity(_763_v15)
				}
				return _dafny.IntOfUint32((_779_v26).Cardinality())
			})()).Plus(p1))
			_ = _rhs100
			var _lhs54 _dafny.Array = _765_v17
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_765_v17), 0))
			_ = _lhs55
			var _lhs56 _dafny.Array = _777_v25
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
			_ = _lhs57
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			_764_v16 = _rhs96
			(_lhs54).ArraySet1(_rhs97, (_lhs55).Int())
			(_lhs56).ArraySet1(_rhs98, (_lhs57).Int())
			_782_v29 = _rhs99
			_lhs58.F2 = _rhs100
			if _742_v0 {
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
				_ = _index106
				(_777_v25).ArraySet1((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))
				_ = _index107
				(_775_v23).ArraySet1(_742_v0, (_index107).Int())
				var _785_v32 _dafny.Map
				_ = _785_v32
				_785_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _742_v0)
				var _786_v33 _dafny.Set
				_ = _786_v33
				_786_v33 = _dafny.SetOf((_dafny.Zero).Minus((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int)), (_783_v30).Cardinality())
				var _787_v34 _dafny.Sequence
				_ = _787_v34
				_787_v34 = _dafny.SeqOf(_773_v21)
				var _788_v35 D2
				_ = _788_v35
				_788_v35 = Companion_D2_.Create_DC5_((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _742_v0, _742_v0, (func() bool {
					if (_785_v32).Contains(_742_v0) {
						return (_785_v32).Get(_742_v0).(bool)
					}
					return _742_v0
				})(), (_this).Fm7(p0, _742_v0, _771_v19, globalState))
				var _789_v36 _dafny.Map
				_ = _789_v36
				_789_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_775_v23, _788_v35)
				var _790_v37 _dafny.Map
				_ = _790_v37
				_790_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v36, _786_v33)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))
				_ = _index108
				var _rhs101 bool = !(true)
				_ = _rhs101
				var _rhs102 _dafny.Int = (((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int)).Plus(p0)).Plus((p1).Plus(p0))
				_ = _rhs102
				var _rhs103 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}(func(_791_i4 _dafny.Int) _dafny.CodePoint {
					return (_this).F6()
				}))
				_ = _rhs103
				var _rhs104 _dafny.Map = (_785_v32).Update((_773_v21).Equals((_787_v34).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_787_v34).Cardinality()))).Uint32()).(_dafny.Map)), _dafny.Companion_Sequence_.Contains(_762_v14, (_this).F6()))
				_ = _rhs104
				var _rhs105 _dafny.Set = (func() _dafny.Set {
					if (_790_v37).Contains(_789_v36) {
						return (_790_v37).Get(_789_v36).(_dafny.Set)
					}
					return _786_v33
				})()
				_ = _rhs105
				var _lhs59 _dafny.Array = _775_v23
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))
				_ = _lhs60
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				(_lhs59).ArraySet1(_rhs101, (_lhs60).Int())
				_lhs61.F2 = _rhs102
				_762_v14 = _rhs103
				_785_v32 = _rhs104
				_786_v33 = _rhs105
				var _792_v38 _dafny.Map
				_ = _792_v38
				_792_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), (_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool))
				var _793_v41 _dafny.Map
				_ = _793_v41
				_793_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool))
				var _794_v42 _dafny.Array
				_ = _794_v42
				var _nwElement0_34 _dafny.Map = _792_v38
				_ = _nwElement0_34
				var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(14))
				_ = _nw124
				(_nw124).ArraySet1(_nwElement0_34, 0)
				(_nw124).ArraySet1(_792_v38, 1)
				(_nw124).ArraySet1(func() _dafny.Map {
					var _coll26 = _dafny.NewMapBuilder()
					_ = _coll26
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(305), _dafny.IntOfInt64(57))); ; {
						_compr_26, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _795_v39 _dafny.Int
						_795_v39 = interface{}(_compr_26).(_dafny.Int)
						if ((_dafny.IntOfInt64(305)).Cmp(_795_v39) <= 0) && ((_795_v39).Cmp(_dafny.IntOfInt64(57)) < 0) {
							_coll26.Add((_795_v39).Minus(p1), _742_v0)
						}
					}
					return _coll26.ToMap()
				}(), 2)
				(_nw124).ArraySet1(_792_v38, 3)
				(_nw124).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), (_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool)), 4)
				(_nw124).ArraySet1(_792_v38, 5)
				(_nw124).ArraySet1(_792_v38, 6)
				(_nw124).ArraySet1(func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter28 := _dafny.Iterate((_792_v38).Keys().Elements()); ; {
						_compr_27, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _796_v40 _dafny.Int
						_796_v40 = interface{}(_compr_27).(_dafny.Int)
						if (_792_v38).Contains(_796_v40) {
							_coll27.Add(Companion_Default___.SafeModuloInt(_796_v40, p0), (_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool))
						}
					}
					return _coll27.ToMap()
				}(), 7)
				(_nw124).ArraySet1(Companion_Default___.Fm34((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool), _dafny.CodePoint('k'), _763_v15, globalState), 8)
				(_nw124).ArraySet1(_792_v38, 9)
				(_nw124).ArraySet1(_792_v38, 10)
				(_nw124).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_783_v30).Cardinality(), (func() bool {
					if (_793_v41).Contains((_this).F6()) {
						return (_793_v41).Get((_this).F6()).(bool)
					}
					return (func() bool {
						if (_785_v32).Contains((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool)) {
							return (_785_v32).Get((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool)).(bool)
						}
						return _742_v0
					})()
				})()), 11)
				(_nw124).ArraySet1(_792_v38, 12)
				(_nw124).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _742_v0), 13)
				_794_v42 = _nw124
				var _797_v43 T0
				_ = _797_v43
				var _nw125 *C2 = New_C2_()
				_ = _nw125
				_nw125.Ctor__(true, (_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _794_v42)
				_797_v43 = _nw125
				var _798_v44 _dafny.Map
				_ = _798_v44
				_798_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _742_v0, _771_v19, globalState), _797_v43)
				var _799_v45 _dafny.Sequence
				_ = _799_v45
				_799_v45 = _dafny.SeqOf(_797_v43)
				_798_v44 = (_798_v44).Update(true, (_799_v45).Select((Companion_Default___.SafeIndex((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_799_v45).Cardinality()))).Uint32()).(T0))
				var _800_v46 D3
				_ = _800_v46
				_800_v46 = Companion_D3_.Create_DC7_(_dafny.IntOfUint32((_762_v14).Cardinality()), p0)
				var _801_v47 _dafny.Map
				_ = _801_v47
				_801_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_800_v46, _777_v25)
				var _802_v48 _dafny.Map
				_ = _802_v48
				_802_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _777_v25)
				var _803_v49 _dafny.Array
				_ = _803_v49
				var _nwElement0_35 _dafny.Array = (func() _dafny.Array {
					if (_801_v47).Contains(_800_v46) {
						return (_801_v47).Get(_800_v46).(_dafny.Array)
					}
					return _777_v25
				})()
				_ = _nwElement0_35
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(17))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_35, 0)
				(_nw126).ArraySet1(_777_v25, 1)
				(_nw126).ArraySet1(_777_v25, 2)
				(_nw126).ArraySet1(_777_v25, 3)
				(_nw126).ArraySet1(_777_v25, 4)
				(_nw126).ArraySet1(_777_v25, 5)
				(_nw126).ArraySet1(_777_v25, 6)
				(_nw126).ArraySet1((func() _dafny.Array {
					if (_802_v48).Contains((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool)) {
						return (_802_v48).Get((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool)).(_dafny.Array)
					}
					return _777_v25
				})(), 7)
				(_nw126).ArraySet1(_777_v25, 8)
				(_nw126).ArraySet1(_777_v25, 9)
				(_nw126).ArraySet1(_777_v25, 10)
				(_nw126).ArraySet1(_777_v25, 11)
				(_nw126).ArraySet1(_777_v25, 12)
				(_nw126).ArraySet1(_777_v25, 13)
				(_nw126).ArraySet1(_777_v25, 14)
				(_nw126).ArraySet1(_777_v25, 15)
				(_nw126).ArraySet1(_777_v25, 16)
				_803_v49 = _nw126
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw127
				_803_v49 = _nw127
				_785_v32 = (_785_v32).Update((_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool), (_775_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_775_v23), 0))).Int()).(bool))
			} else {
				var _804_v50 D3
				_ = _804_v50
				_804_v50 = Companion_D3_.Create_DC7_((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), p0)
				var _805_v51 _dafny.Sequence
				_ = _805_v51
				_805_v51 = _dafny.SeqOf(_804_v50, _804_v50, _804_v50)
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_775_v23), 0))
				_ = _index109
				(_775_v23).ArraySet1(_742_v0, (_index109).Int())
				var _806_v52 D4
				_ = _806_v52
				_806_v52 = Companion_D4_.Create_DC9_(p0, (_this).F6(), _763_v15, _dafny.IntOfUint32((_762_v14).Cardinality()), true)
				var _807_v53 _dafny.Map
				_ = _807_v53
				_807_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_806_v52, _775_v23)
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_775_v23), 0))
				_ = _index110
				var _rhs106 _dafny.Sequence = _805_v51
				_ = _rhs106
				var _rhs107 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_762_v14, (Companion_Default___.SafeIndex((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_762_v14).Cardinality()))).Uint32(), (_this).F6())
				_ = _rhs107
				var _rhs108 _dafny.Int = Companion_Default___.Fm17(_dafny.SeqOf(_762_v14), p1, Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(-351)), (_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), globalState)
				_ = _rhs108
				var _rhs109 _dafny.Array = (func() _dafny.Array {
					if (_807_v53).Contains(_806_v52) {
						return (_807_v53).Get(_806_v52).(_dafny.Array)
					}
					return _775_v23
				})()
				_ = _rhs109
				var _rhs110 bool = (func() bool {
					if (_779_v26).Select((Companion_Default___.SafeIndex((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_779_v26).Cardinality()))).Uint32()).(bool) {
						return _742_v0
					}
					return !((_dafny.IntOfUint32((_762_v14).Cardinality())).Cmp((_777_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))).Int()).(_dafny.Int)) > 0)
				})()
				_ = _rhs110
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 _dafny.Array = _775_v23
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_775_v23), 0))
				_ = _lhs64
				_805_v51 = _rhs106
				_762_v14 = _rhs107
				_lhs62.F2 = _rhs108
				_775_v23 = _rhs109
				(_lhs63).ArraySet1(_rhs110, (_lhs64).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_775_v23), 0))
				_ = _index111
				(_775_v23).ArraySet1(Companion_Default___.Fm23(_dafny.Companion_Sequence_.Concatenate(_762_v14, _dafny.Companion_Sequence_.Update(_762_v14, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_763_v15).Cardinality()), _dafny.IntOfUint32((_762_v14).Cardinality()))).Uint32(), (_this).F6())), p1, p1, !(false), globalState), (_index111).Int())
				var _808_v54 _dafny.Sequence
				_ = _808_v54
				_808_v54 = _dafny.SeqOf(_781_v28)
				_742_v0 = _dafny.Companion_Sequence_.IsPrefixOf(_808_v54, _808_v54)
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
				_ = _index112
				(_777_v25).ArraySet1(p0, (_index112).Int())
				_773_v21 = _773_v21
			}
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
			_ = _index113
			var _rhs111 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(356))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}(func(_809_i5 _dafny.Int) _dafny.CodePoint {
				return (_this).F6()
			}))).Cardinality())
			_ = _rhs111
			var _rhs112 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.Companion_Sequence_.Concatenate(_762_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(121))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}(func(_810_i6 _dafny.Int) _dafny.CodePoint {
				return (_this).F6()
			}))))
			_ = _rhs112
			var _lhs65 _dafny.Array = _777_v25
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_777_v25), 0))
			_ = _lhs66
			(_lhs65).ArraySet1(_rhs111, (_lhs66).Int())
			_761_v13 = _rhs112
			var _811_v55 _dafny.Array
			_ = _811_v55
			var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw128
			_811_v55 = _nw128
			var _812_v56 *C2
			_ = _812_v56
			var _nw129 *C2 = New_C2_()
			_ = _nw129
			_nw129.Ctor__((_779_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_762_v14).Cardinality()), _dafny.IntOfUint32((_779_v26).Cardinality()))).Uint32()).(bool), _dafny.IntOfUint32((_779_v26).Cardinality()), _811_v55)
			_812_v56 = _nw129
		}
		var _813_v57 D6
		_ = _813_v57
		_813_v57 = Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("x"))
		var _source14 D6 = _813_v57
		_ = _source14
		if _source14.Is_DC12() {
			var _814___mcc_h0 bool = _source14.Get_().(D6_DC12).Cf23
			_ = _814___mcc_h0
			var _815___mcc_h1 _dafny.CodePoint = _source14.Get_().(D6_DC12).Cf24
			_ = _815___mcc_h1
			var _816_cf24 _dafny.CodePoint = _815___mcc_h1
			_ = _816_cf24
			var _817_cf23 bool = _814___mcc_h0
			_ = _817_cf23
			if _817_cf23 {
				var _818_v58 _dafny.Sequence
				_ = _818_v58
				_818_v58 = _dafny.UnicodeSeqOfUtf8Bytes("wgyx")
				var _819_v59 _dafny.Array
				_ = _819_v59
				var _nwElement0_36 _dafny.Sequence = _818_v58
				_ = _nwElement0_36
				var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(6))
				_ = _nw130
				(_nw130).ArraySet1(_nwElement0_36, 0)
				(_nw130).ArraySet1(_818_v58, 1)
				(_nw130).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(536))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_820_cf24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_821_i7 _dafny.Int) _dafny.CodePoint {
						return _820_cf24
					}
				})(_816_cf24))), 2)
				(_nw130).ArraySet1(_818_v58, 3)
				(_nw130).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ef"), 4)
				(_nw130).ArraySet1(_818_v58, 5)
				_819_v59 = _nw130
				var _822_v60 D9
				_ = _822_v60
				_822_v60 = Companion_D9_.Create_DC18_(_819_v59)
				var _823_v61 _dafny.Array
				_ = _823_v61
				var _nwElement0_37 _dafny.Array = _819_v59
				_ = _nwElement0_37
				var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(18))
				_ = _nw131
				(_nw131).ArraySet1(_nwElement0_37, 0)
				(_nw131).ArraySet1(_819_v59, 1)
				(_nw131).ArraySet1(_819_v59, 2)
				(_nw131).ArraySet1(_819_v59, 3)
				(_nw131).ArraySet1(_819_v59, 4)
				(_nw131).ArraySet1(_819_v59, 5)
				(_nw131).ArraySet1(_819_v59, 6)
				(_nw131).ArraySet1(_819_v59, 7)
				(_nw131).ArraySet1(_819_v59, 8)
				(_nw131).ArraySet1(_819_v59, 9)
				(_nw131).ArraySet1(_819_v59, 10)
				(_nw131).ArraySet1((_822_v60).Dtor_cf33(), 11)
				(_nw131).ArraySet1(_819_v59, 12)
				(_nw131).ArraySet1(_819_v59, 13)
				(_nw131).ArraySet1(_819_v59, 14)
				(_nw131).ArraySet1((func() _dafny.Array {
					if _817_cf23 {
						return _819_v59
					}
					return _819_v59
				})(), 15)
				(_nw131).ArraySet1(_819_v59, 16)
				(_nw131).ArraySet1(_819_v59, 17)
				_823_v61 = _nw131
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_823_v61), 0))
				_ = _index114
				(_823_v61).ArraySet1(_819_v59, (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_823_v61), 0))
				_ = _index115
				(_823_v61).ArraySet1(_819_v59, (_index115).Int())
				var _824_v62 _dafny.Array
				_ = _824_v62
				var _nw132 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw132
				_824_v62 = _nw132
				_824_v62 = _824_v62
				_817_cf23 = !(Companion_Default___.Fm21((p1).Times(_dafny.IntOfInt64(73)), _dafny.CodePoint('i'), _dafny.Companion_Sequence_.IsProperPrefixOf(_818_v58, _818_v58), _817_cf23, globalState))
				var _825_v63 _dafny.Set
				_ = _825_v63
				_825_v63 = _dafny.SetOf(!(_742_v0))
				var _826_v64 _dafny.Map
				_ = _826_v64
				_826_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_v63, _818_v58)
				_826_v64 = (_826_v64).Update(_825_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(981))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}(func(_827_i8 _dafny.Int) _dafny.CodePoint {
					return (_this).F6()
				})))
				var _828_v66 D2
				_ = _828_v66
				_828_v66 = Companion_D2_.Create_DC5_((func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(345), _dafny.IntOfInt64(-615))); ; {
						_compr_28, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _829_v65 _dafny.Int
						_829_v65 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(345)).Cmp(_829_v65) <= 0) && ((_829_v65).Cmp(_dafny.IntOfInt64(-615)) < 0) {
							_coll28.Add((_829_v65).Times(p0), p1)
						}
					}
					return _coll28.ToMap()
				}()).Cardinality(), _742_v0, _817_cf23, _817_cf23, _817_cf23)
				var _pat_let_tv20 = _817_cf23
				_ = _pat_let_tv20
				var _pat_let_tv21 = _817_cf23
				_ = _pat_let_tv21
				_828_v66 = func(_pat_let6_0 D2) D2 {
					return func(_830_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let7_0 bool) D2 {
							return func(_831_dt__update_hcf8_h0 bool) D2 {
								return func(_pat_let8_0 bool) D2 {
									return func(_832_dt__update_hcf10_h0 bool) D2 {
										return Companion_D2_.Create_DC5_((_830_dt__update__tmp_h0).Dtor_cf7(), _831_dt__update_hcf8_h0, (_830_dt__update__tmp_h0).Dtor_cf9(), _832_dt__update_hcf10_h0, (_830_dt__update__tmp_h0).Dtor_cf11())
									}(_pat_let8_0)
								}(_pat_let_tv21)
							}(_pat_let7_0)
						}(_pat_let_tv20)
					}(_pat_let6_0)
				}(_828_v66)
			} else {
				var _833_v67 _dafny.Map
				_ = _833_v67
				_833_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_742_v0) && (_817_cf23), p1)
				_833_v67 = (_833_v67).Update(_817_cf23, p0)
				var _834_v68 _dafny.Sequence
				_ = _834_v68
				_834_v68 = _dafny.SeqOf(_dafny.IntOfInt64(291), p0)
				var _835_v69 _dafny.Sequence
				_ = _835_v69
				_835_v69 = _dafny.UnicodeSeqOfUtf8Bytes("ulgwrp")
				var _836_v70 _dafny.Array
				_ = _836_v70
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
				_ = _nw133
				_836_v70 = _nw133
				var _837_v71 T0
				_ = _837_v71
				var _nw134 *C2 = New_C2_()
				_ = _nw134
				_nw134.Ctor__(true, p1, _836_v70)
				_837_v71 = _nw134
				var _838_v72 _dafny.Set
				_ = _838_v72
				_838_v72 = _dafny.SetOf(p0)
				var _839_v73 *C0
				_ = _839_v73
				var _nw135 *C0 = New_C0_()
				_ = _nw135
				_nw135.Ctor__(p0)
				_839_v73 = _nw135
				var _840_v74 _dafny.MultiSet
				_ = _840_v74
				_840_v74 = _dafny.MultiSetOf(_839_v73, _839_v73)
				var _841_v75 _dafny.Array
				_ = _841_v75
				var _nwElement0_38 _dafny.Int = p0
				_ = _nwElement0_38
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(15))
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_38, 0)
				(_nw136).ArraySet1(_dafny.IntOfUint32((_835_v69).Cardinality()), 1)
				(_nw136).ArraySet1(p0, 2)
				(_nw136).ArraySet1(p1, 3)
				(_nw136).ArraySet1(p1, 4)
				(_nw136).ArraySet1(p1, 5)
				(_nw136).ArraySet1(_dafny.IntOfUint32((_834_v68).Cardinality()), 6)
				(_nw136).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v71, _817_cf23)).Cardinality(), 7)
				(_nw136).ArraySet1((_838_v72).Cardinality(), 8)
				(_nw136).ArraySet1(p1, 9)
				(_nw136).ArraySet1(p1, 10)
				(_nw136).ArraySet1(p1, 11)
				(_nw136).ArraySet1(_dafny.IntOfUint32((_834_v68).Cardinality()), 12)
				(_nw136).ArraySet1((_840_v74).Cardinality(), 13)
				(_nw136).ArraySet1(_dafny.IntOfInt64(-423), 14)
				_841_v75 = _nw136
				var _842_v76 _dafny.Map
				_ = _842_v76
				_842_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_841_v75, _742_v0)
				var _843_v77 D8
				_ = _843_v77
				_843_v77 = Companion_D8_.Create_DC16_((_834_v68).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_834_v68).Cardinality()))).Uint32()).(_dafny.Int), (func() bool {
					if (_842_v76).Contains(_841_v75) {
						return (_842_v76).Get(_841_v75).(bool)
					}
					return _817_cf23
				})(), _817_cf23, p1)
				_742_v0 = (_843_v77).Dtor_cf29()
				var _844_v78 _dafny.Sequence
				_ = _844_v78
				_844_v78 = _dafny.SeqOf(_835_v69)
				_835_v69 = _dafny.Companion_Sequence_.Concatenate(_835_v69, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_835_v69, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}(func(_845_i9 _dafny.Int) _dafny.CodePoint {
					return (_this).F6()
				}))), (Companion_Default___.SafeIndex((_834_v68).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm17(_844_v78, p0, _dafny.IntOfUint32((_835_v69).Cardinality()), (_838_v72).Cardinality(), globalState), _dafny.IntOfUint32((_834_v68).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_835_v69, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_846_i9 _dafny.Int) _dafny.CodePoint {
					return (_this).F6()
				})))).Cardinality()))).Uint32(), _816_cf24))
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_841_v75), 0))
				_ = _index116
				(_841_v75).ArraySet1(p0, (_index116).Int())
				var _847_v79 _dafny.Map
				_ = _847_v79
				_847_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm35(globalState), _817_cf23)
				var _848_v80 _dafny.Set
				_ = _848_v80
				_848_v80 = _dafny.SetOf(_835_v69)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_841_v75), 0))
				_ = _index117
				var _rhs113 _dafny.Int = p1
				_ = _rhs113
				var _rhs114 D6 = _813_v57
				_ = _rhs114
				var _rhs115 _dafny.Int = p1
				_ = _rhs115
				var _rhs116 bool = (func() bool {
					if (_847_v79).Contains(_848_v80) {
						return (_847_v79).Get(_848_v80).(bool)
					}
					return !(_742_v0)
				})()
				_ = _rhs116
				var _rhs117 _dafny.Set = _838_v72
				_ = _rhs117
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				var _lhs68 _dafny.Array = _841_v75
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_841_v75), 0))
				_ = _lhs69
				_lhs67.F2 = _rhs113
				_813_v57 = _rhs114
				(_lhs68).ArraySet1(_rhs115, (_lhs69).Int())
				_742_v0 = _rhs116
				_838_v72 = _rhs117
				var _849_v81 _dafny.Map
				_ = _849_v81
				_849_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_841_v75, p0)
				_849_v81 = (_849_v81).Update(_841_v75, _dafny.IntOfInt64(657))
			}
			(globalState).F2 = p1
			var _850_v82 _dafny.Sequence
			_ = _850_v82
			_850_v82 = _dafny.UnicodeSeqOfUtf8Bytes("a")
			_850_v82 = _dafny.Companion_Sequence_.Concatenate(_850_v82, (func() _dafny.Sequence {
				if _742_v0 {
					return _850_v82
				}
				return _850_v82
			})())
			if Companion_Default___.Fm23(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_851_cf24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_852_i10 _dafny.Int) _dafny.CodePoint {
					return _851_cf24
				}
			})(_816_cf24))), _dafny.IntOfInt64(296), Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_850_v82).Cardinality())), _817_cf23, globalState) {
				var _853_v83 _dafny.Array
				_ = _853_v83
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_20
				var _nw137 _dafny.Array
				_ = _nw137
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw137 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = (func(_854_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_855_i11 _dafny.Int) _dafny.Int {
							return (_855_i11).Minus(_854_p0)
						}
					})(p0)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw137 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw137).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw137).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_853_v83 = _nw137
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_853_v83), 0))
				_ = _index118
				(_853_v83).ArraySet1(p0, (_index118).Int())
				var _856_v84 _dafny.Sequence
				_ = _856_v84
				_856_v84 = _dafny.SeqOf(_850_v82, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_857_cf24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_858_i12 _dafny.Int) _dafny.CodePoint {
						return _857_cf24
					}
				})(_816_cf24))))
				var _859_v85 _dafny.Sequence
				_ = _859_v85
				_859_v85 = _dafny.SeqOf((_856_v84).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_856_v84).Cardinality()))).Uint32()).(_dafny.Sequence), _850_v82, Companion_Default___.Fm25(_742_v0, p0, globalState))
				var _860_v86 _dafny.Map
				_ = _860_v86
				_860_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _850_v82)
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_853_v83), 0))
				_ = _index119
				var _rhs118 _dafny.Int = Companion_Default___.Fm17((func() _dafny.Sequence {
					if _742_v0 {
						return _859_v85
					}
					return _859_v85
				})(), p1, _dafny.IntOfInt64(82), p0, globalState)
				_ = _rhs118
				var _rhs119 bool = (_this).Fm7(p0, _817_cf23, (_860_v86).Update(p0, _850_v82), globalState)
				_ = _rhs119
				var _lhs70 _dafny.Array = _853_v83
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_853_v83), 0))
				_ = _lhs71
				(_lhs70).ArraySet1(_rhs118, (_lhs71).Int())
				_817_cf23 = _rhs119
				var _861_v87 _dafny.Array
				_ = _861_v87
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_21
				var _nw138 _dafny.Array
				_ = _nw138
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw138 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Sequence = (func(_862_v0 bool, _863_cf23 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_864_i13 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_862_v0, _863_cf23), _dafny.SeqOf(!(true)))
						}
					})(_742_v0, _817_cf23)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw138 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw138).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw138).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_861_v87 = _nw138
				var _865_v88 _dafny.Sequence
				_ = _865_v88
				_865_v88 = _dafny.SeqOf(_817_cf23)
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_861_v87), 0))
				_ = _index120
				(_861_v87).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_865_v88, _865_v88), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_861_v87), 0))
				_ = _index121
				(_861_v87).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_817_cf23), _865_v88), (_index121).Int())
				var _866_v89 _dafny.Sequence
				_ = _866_v89
				_866_v89 = _dafny.SeqOf(((_853_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_853_v83), 0))).Int()).(_dafny.Int)).Plus(p0))
				var _867_v90 D4
				_ = _867_v90
				_867_v90 = Companion_D4_.Create_DC9_((_dafny.Zero).Minus((_853_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_853_v83), 0))).Int()).(_dafny.Int)), (_this).F6(), _866_v89, p0, _817_cf23)
				var _868_v91 _dafny.Set
				_ = _868_v91
				_868_v91 = _dafny.SetOf(_742_v0, _817_cf23)
				var _869_v92 _dafny.Map
				_ = _869_v92
				_869_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_866_v89 = (func() _dafny.Sequence {
					if _817_cf23 {
						return _866_v89
					}
					return _dafny.Companion_Sequence_.Concatenate((_867_v90).Dtor_cf18(), Companion_Default___.Fm12(true, p0, _868_v91, _869_v92, globalState))
				})()
				var _870_v93 _dafny.Map
				_ = _870_v93
				_870_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
				_870_v93 = (_870_v93).Update(_817_cf23, (_dafny.Zero).Minus(p1))
				_742_v0 = (func() bool {
					if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(900))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_871_v83 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_872_i14 _dafny.Int) _dafny.Int {
							return (_871_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_871_v83), 0))).Int()).(_dafny.Int)
						}
					})(_853_v83))), _dafny.SeqOf(p0, p0)) {
						return (_dafny.IntOfInt64(834)).Cmp(p1) == 0
					}
					return _742_v0
				})()
			} else {
				var _873_v94 _dafny.Array
				_ = _873_v94
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_22
				var _nw139 _dafny.Array
				_ = _nw139
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw139 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Int = (func(_874_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_875_i15 _dafny.Int) _dafny.Int {
							return (_875_i15).Plus(_874_p1)
						}
					})(p1)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw139 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw139).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw139).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_873_v94 = _nw139
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_873_v94), 0))
				_ = _index122
				(_873_v94).ArraySet1(p1, (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_873_v94), 0))
				_ = _index123
				(_873_v94).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_850_v82, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_850_v82).Cardinality()))).Uint32(), _816_cf24), _dafny.UnicodeSeqOfUtf8Bytes("q")), _dafny.UnicodeSeqOfUtf8Bytes("lojxgqr"))).Cardinality()), (_index123).Int())
				var _876_v95 *C0
				_ = _876_v95
				var _nw140 *C0 = New_C0_()
				_ = _nw140
				_nw140.Ctor__(p0)
				_876_v95 = _nw140
				var _877_v96 D4
				_ = _877_v96
				_877_v96 = Companion_D4_.Create_DC8_(_876_v95)
				var _878_v97 _dafny.MultiSet
				_ = _878_v97
				_878_v97 = _dafny.MultiSetOf(_877_v96, _877_v96)
				var _879_v98 D7
				_ = _879_v98
				_879_v98 = Companion_D7_.Create_DC14_((func() _dafny.Int {
					if (_878_v97).Contains(_877_v96) {
						return (_878_v97).Multiplicity(_877_v96)
					}
					return p0
				})())
				var _pat_let_tv22 = _873_v94
				_ = _pat_let_tv22
				var _pat_let_tv23 = _873_v94
				_ = _pat_let_tv23
				_879_v98 = func(_pat_let9_0 D7) D7 {
					return func(_880_dt__update__tmp_h1 D7) D7 {
						return func(_pat_let10_0 _dafny.Int) D7 {
							return func(_881_dt__update_hcf26_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC14_(_881_dt__update_hcf26_h0)
							}(_pat_let10_0)
						}((_pat_let_tv23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_pat_let_tv22), 0))).Int()).(_dafny.Int))
					}(_pat_let9_0)
				}(Companion_Default___.Fm36(_742_v0, globalState))
				var _882_v99 _dafny.Array
				_ = _882_v99
				_882_v99 = _873_v94
				var _883_v100 _dafny.MultiSet
				_ = _883_v100
				_883_v100 = _dafny.MultiSetOf((_882_v99), _873_v94, _873_v94)
				_883_v100 = _dafny.MultiSetOf(_873_v94, (func() _dafny.Array {
					if _742_v0 {
						return _873_v94
					}
					return _873_v94
				})(), _873_v94, _873_v94, _873_v94)
				_883_v100 = _883_v100
				_817_cf23 = _817_cf23
			}
		} else {
			var _884___mcc_h2 _dafny.Sequence = _source14.Get_().(D6_DC11).Cf22
			_ = _884___mcc_h2
			var _885_cf22 _dafny.Sequence = _884___mcc_h2
			_ = _885_cf22
			var _886_v101 _dafny.Sequence
			_ = _886_v101
			_886_v101 = _dafny.SeqOf(_885_cf22, _885_cf22)
			var _887_v102 _dafny.Map
			_ = _887_v102
			_887_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains((_813_v57).Dtor_cf22(), (_this).F6()), Companion_Default___.Fm17(_886_v101, p1, _dafny.IntOfInt64(803), p0, globalState))
			_887_v102 = (_887_v102).Update(_742_v0, (func() _dafny.Int {
				if _742_v0 {
					return p1
				}
				return _dafny.IntOfUint32((_885_cf22).Cardinality())
			})())
			var _888_v103 _dafny.Sequence
			_ = _888_v103
			_888_v103 = _dafny.SeqOf(p0, _dafny.IntOfInt64(192))
			var _889_v104 D2
			_ = _889_v104
			_889_v104 = Companion_D2_.Create_DC4_((_888_v103).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_888_v103).Cardinality()))).Uint32()).(_dafny.Int))
			var _890_v105 _dafny.Sequence
			_ = _890_v105
			_890_v105 = _dafny.SeqOf(_889_v104)
			_890_v105 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_889_v104), _dafny.Companion_Sequence_.Concatenate(_890_v105, _890_v105))
			var _891_v106 _dafny.Array
			_ = _891_v106
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw141
			_891_v106 = _nw141
			var _892_v107 _dafny.Sequence
			_ = _892_v107
			_892_v107 = _dafny.SeqOf(_891_v106)
			var _893_v108 _dafny.Sequence
			_ = _893_v108
			_893_v108 = _dafny.SeqOf((_892_v107).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_892_v107).Cardinality()))).Uint32()).(_dafny.Array), _891_v106)
			var _894_v109 D8
			_ = _894_v109
			_894_v109 = Companion_D8_.Create_DC15_((_893_v108).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_893_v108).Cardinality()))).Uint32()).(_dafny.Array))
			var _source15 D8 = _894_v109
			_ = _source15
			if _source15.Is_DC16() {
				var _895___mcc_h3 _dafny.Int = _source15.Get_().(D8_DC16).Cf28
				_ = _895___mcc_h3
				var _896___mcc_h4 bool = _source15.Get_().(D8_DC16).Cf29
				_ = _896___mcc_h4
				var _897___mcc_h5 bool = _source15.Get_().(D8_DC16).Cf30
				_ = _897___mcc_h5
				var _898___mcc_h6 _dafny.Int = _source15.Get_().(D8_DC16).Cf31
				_ = _898___mcc_h6
				var _899_cf31 _dafny.Int = _898___mcc_h6
				_ = _899_cf31
				var _900_cf30 bool = _897___mcc_h5
				_ = _900_cf30
				var _901_cf29 bool = _896___mcc_h4
				_ = _901_cf29
				var _902_cf28 _dafny.Int = _895___mcc_h3
				_ = _902_cf28
				_742_v0 = _742_v0
				_900_cf30 = _900_cf30
				var _903_v110 _dafny.Array
				_ = _903_v110
				var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw142
				_903_v110 = _nw142
				var _904_v111 _dafny.MultiSet
				_ = _904_v111
				_904_v111 = _dafny.MultiSetOf(p1)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_903_v110), 0))
				_ = _index124
				(_903_v110).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_904_v111).Contains(_899_cf31) {
						return (_904_v111).Multiplicity(_899_cf31)
					}
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_cf29, p0)).Cardinality()
				})(), (_889_v104).Dtor_cf6()), (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_903_v110), 0))
				_ = _index125
				(_903_v110).ArraySet1(_902_cf28, (_index125).Int())
				var _905_v112 _dafny.CodePoint
				_ = _905_v112
				_905_v112 = _dafny.CodePoint('m')
				_905_v112 = _905_v112
			} else if _source15.Is_DC15() {
				var _906___mcc_h7 _dafny.Array = _source15.Get_().(D8_DC15).Cf27
				_ = _906___mcc_h7
				var _907_cf27 _dafny.Array = _906___mcc_h7
				_ = _907_cf27
				var _908_v113 _dafny.Sequence
				_ = _908_v113
				_908_v113 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _742_v0))
				var _909_v114 _dafny.Map
				_ = _909_v114
				_909_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_908_v113).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_908_v113).Cardinality()))).Uint32()).(_dafny.Map))
				var _910_v115 _dafny.Set
				_ = _910_v115
				_910_v115 = _dafny.SetOf(p1)
				var _911_v116 _dafny.Map
				_ = _911_v116
				_911_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _742_v0)
				var _912_v117 _dafny.MultiSet
				_ = _912_v117
				_912_v117 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _742_v0), (func() _dafny.Map {
					if (_909_v114).Contains(Companion_Default___.Fm24(_910_v115, globalState)) {
						return (_909_v114).Get(Companion_Default___.Fm24(_910_v115, globalState)).(_dafny.Map)
					}
					return _911_v116
				})())
				var _913_v118 _dafny.Sequence
				_ = _913_v118
				_913_v118 = _dafny.SeqOf(_912_v117)
				var _914_v119 _dafny.Sequence
				_ = _914_v119
				_914_v119 = _dafny.SeqOf(_742_v0, _742_v0, _742_v0, _742_v0)
				_742_v0 = (((_913_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_914_v119).Cardinality()), _dafny.IntOfUint32((_913_v118).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_912_v117)).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_908_v113, Companion_Default___.Fm37(Companion_D1_.Create_DC2_(p1, false, p0), _742_v0, globalState))))
				_742_v0 = _742_v0
				(globalState).F2 = _dafny.IntOfInt64(14)
				var _915_v121 _dafny.Map
				_ = _915_v121
				_915_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_885_cf22, _dafny.IntOfInt64(746))
				var _916_v122 _dafny.MultiSet
				_ = _916_v122
				_916_v122 = _dafny.MultiSetOf(p1, (func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate((_915_v121).Keys().Elements()); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _917_v120 _dafny.Sequence
						_917_v120 = interface{}(_compr_29).(_dafny.Sequence)
						if (_915_v121).Contains(_917_v120) {
							_coll29.Add(_917_v120, _742_v0)
						}
					}
					return _coll29.ToMap()
				}()).Cardinality(), p1)
				var _918_v123 *C3
				_ = _918_v123
				var _nw143 *C3 = New_C3_()
				_ = _nw143
				_nw143.Ctor__(_916_v122, _742_v0)
				_918_v123 = _nw143
				var _nw144 *C3 = New_C3_()
				_ = _nw144
				_nw144.Ctor__((_918_v123).F9(), true)
				_918_v123 = _nw144
			} else {
				var _919___mcc_h8 D8 = _source15.Get_().(D8_DC17).Cf32
				_ = _919___mcc_h8
				var _920_cf32 D8 = _919___mcc_h8
				_ = _920_cf32
				(globalState).F2 = p0
				var _921_v124 _dafny.MultiSet
				_ = _921_v124
				_921_v124 = _dafny.MultiSetOf(p1, p1)
				var _922_v125 _dafny.Array
				_ = _922_v125
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw145
				_922_v125 = _nw145
				var _923_v126 *C4
				_ = _923_v126
				var _nw146 *C4 = New_C4_()
				_ = _nw146
				_nw146.Ctor__(_742_v0, _921_v124, _922_v125)
				_923_v126 = _nw146
				var _924_v127 _dafny.Array
				_ = _924_v127
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_23
				var _nw147 _dafny.Array
				_ = _nw147
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw147 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = func(_925_i16 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_925_i16, _dafny.IntOfInt64(778))
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw147 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw147).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw147).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_924_v127 = _nw147
				var _926_v128 _dafny.Map
				_ = _926_v128
				_926_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _924_v127)
				var _927_v129 _dafny.Array
				_ = _927_v129
				_927_v129 = _924_v127
				var _928_v130 _dafny.Map
				_ = _928_v130
				_928_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v0, _924_v127)
				var _929_v131 _dafny.Array
				_ = _929_v131
				var _nwElement0_39 _dafny.Array = (func() _dafny.Array {
					if _742_v0 {
						return _924_v127
					}
					return _924_v127
				})()
				_ = _nwElement0_39
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(17))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_39, 0)
				(_nw148).ArraySet1(_924_v127, 1)
				(_nw148).ArraySet1(_924_v127, 2)
				(_nw148).ArraySet1(_924_v127, 3)
				(_nw148).ArraySet1(_924_v127, 4)
				(_nw148).ArraySet1((func() _dafny.Array {
					if (_926_v128).Contains((_this).F6()) {
						return (_926_v128).Get((_this).F6()).(_dafny.Array)
					}
					return _924_v127
				})(), 5)
				(_nw148).ArraySet1(_924_v127, 6)
				(_nw148).ArraySet1(_924_v127, 7)
				(_nw148).ArraySet1((_927_v129), 8)
				(_nw148).ArraySet1(_924_v127, 9)
				(_nw148).ArraySet1(_924_v127, 10)
				(_nw148).ArraySet1(_924_v127, 11)
				(_nw148).ArraySet1(_924_v127, 12)
				(_nw148).ArraySet1(_924_v127, 13)
				(_nw148).ArraySet1((func() _dafny.Array {
					if (_928_v130).Contains((_923_v126).F7()) {
						return (_928_v130).Get((_923_v126).F7()).(_dafny.Array)
					}
					return _924_v127
				})(), 14)
				(_nw148).ArraySet1(_924_v127, 15)
				(_nw148).ArraySet1(_924_v127, 16)
				_929_v131 = _nw148
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_929_v131), 0))
				_ = _index126
				(_929_v131).ArraySet1(_924_v127, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_929_v131), 0))
				_ = _index127
				(_929_v131).ArraySet1(_924_v127, (_index127).Int())
				var _rhs120 bool = _742_v0
				_ = _rhs120
				var _rhs121 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (p0).Cmp(_dafny.IntOfInt64(439)) == 0 {
						return _dafny.Companion_Sequence_.Concatenate(_885_cf22, _885_cf22)
					}
					return _885_cf22
				})()).Cardinality())
				_ = _rhs121
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				_742_v0 = _rhs120
				_lhs72.F2 = _rhs121
			}
			var _930_v132 _dafny.Array
			_ = _930_v132
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw149
			_930_v132 = _nw149
			var _931_v133 _dafny.MultiSet
			_ = _931_v133
			_931_v133 = _dafny.MultiSetOf(p1, p0, p0, p1)
			var _932_v134 _dafny.Map
			_ = _932_v134
			_932_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_742_v0), _742_v0)
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_930_v132), 0))
			_ = _index128
			(_930_v132).ArraySet1((func() _dafny.Int {
				if (_931_v133).Contains(Companion_Default___.Fm9(_dafny.IntOfUint32((_885_cf22).Cardinality()), globalState)) {
					return (_931_v133).Multiplicity(Companion_Default___.Fm9(_dafny.IntOfUint32((_885_cf22).Cardinality()), globalState))
				}
				return (_dafny.Zero).Minus((_932_v134).Cardinality())
			})(), (_index128).Int())
			var _933_v135 _dafny.Map
			_ = _933_v135
			_933_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.Companion_Sequence_.Concatenate(_885_cf22, _885_cf22))
			var _934_v136 _dafny.Map
			_ = _934_v136
			_934_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _930_v132)
			var _935_v138 _dafny.Set
			_ = _935_v138
			_935_v138 = _dafny.SetOf((_this).F6())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_930_v132), 0))
			_ = _index129
			var _rhs122 _dafny.Int = (func() _dafny.Int {
				if !(true) {
					return p0
				}
				return p0
			})()
			_ = _rhs122
			var _rhs123 _dafny.Map = (func() _dafny.Map {
				if _742_v0 {
					return _933_v135
				}
				return (func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter31 := _dafny.Iterate((_935_v138).Elements()); ; {
						_compr_30, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _936_v137 _dafny.CodePoint
						_936_v137 = interface{}(_compr_30).(_dafny.CodePoint)
						if (_935_v138).Contains(_936_v137) {
							_coll30.Add(_936_v137, _885_cf22)
						}
					}
					return _coll30.ToMap()
				}()).Update(_dafny.CodePoint('f'), _885_cf22)
			})()
			_ = _rhs123
			var _rhs124 _dafny.Int = Companion_Default___.SafeModuloInt((p1).Plus(p1), p1)
			_ = _rhs124
			var _rhs125 _dafny.Int = p1
			_ = _rhs125
			var _rhs126 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _930_v132)
			_ = _rhs126
			var _lhs73 _dafny.Array = _930_v132
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_930_v132), 0))
			_ = _lhs74
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			(_lhs73).ArraySet1(_rhs122, (_lhs74).Int())
			_933_v135 = _rhs123
			_lhs75.F2 = _rhs124
			_lhs76.F2 = _rhs125
			_934_v136 = _rhs126
		}
		var _937_v139 _dafny.Sequence
		_ = _937_v139
		_937_v139 = _dafny.UnicodeSeqOfUtf8Bytes("f")
		var _938_v140 _dafny.Sequence
		_ = _938_v140
		_938_v140 = _dafny.SeqOf(p0, p0)
		var _939_v141 _dafny.Map
		_ = _939_v141
		_939_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_938_v140).Cardinality()), !(_742_v0))
		var _940_v142 _dafny.Array
		_ = _940_v142
		var _nwElement0_40 _dafny.Map = _939_v141
		_ = _nwElement0_40
		var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(11))
		_ = _nw150
		(_nw150).ArraySet1(_nwElement0_40, 0)
		(_nw150).ArraySet1(_939_v141, 1)
		(_nw150).ArraySet1(_939_v141, 2)
		(_nw150).ArraySet1(_939_v141, 3)
		(_nw150).ArraySet1(_939_v141, 4)
		(_nw150).ArraySet1(_939_v141, 5)
		(_nw150).ArraySet1(_939_v141, 6)
		(_nw150).ArraySet1(_939_v141, 7)
		(_nw150).ArraySet1(_939_v141, 8)
		(_nw150).ArraySet1(_939_v141, 9)
		(_nw150).ArraySet1(_939_v141, 10)
		_940_v142 = _nw150
		var _941_v143 T0
		_ = _941_v143
		var _nw151 *C2 = New_C2_()
		_ = _nw151
		_nw151.Ctor__(_742_v0, _dafny.IntOfUint32((_937_v139).Cardinality()), _940_v142)
		_941_v143 = _nw151
		var _942_v144 _dafny.Map
		_ = _942_v144
		_942_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _941_v143)
		var _943_v145 D2
		_ = _943_v145
		_943_v145 = Companion_D2_.Create_DC5_(p1, false, _742_v0, _742_v0, _742_v0)
		var _source16 D2 = Companion_D2_.Create_DC5_((_942_v144).Cardinality(), (_943_v145).Dtor_cf8(), _742_v0, _742_v0, !_dafny.Companion_Sequence_.Contains(_937_v139, (_this).F6()))
		_ = _source16
		if _source16.Is_DC4() {
			var _944___mcc_h9 _dafny.Int = _source16.Get_().(D2_DC4).Cf6
			_ = _944___mcc_h9
			var _945_cf6 _dafny.Int = _944___mcc_h9
			_ = _945_cf6
			var _946_v146 *C2
			_ = _946_v146
			var _nw152 *C2 = New_C2_()
			_ = _nw152
			_nw152.Ctor__(false, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if !(_742_v0) {
					return _938_v140
				}
				return _938_v140
			})()).Cardinality()), (_941_v143).F3())
			_946_v146 = _nw152
			var _947_v147 _dafny.Array
			_ = _947_v147
			var _nwElement0_41 bool = (_946_v146).F11()
			_ = _nwElement0_41
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(26))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_41, 0)
			(_nw153).ArraySet1((_946_v146).F11(), 1)
			(_nw153).ArraySet1(_742_v0, 2)
			(_nw153).ArraySet1(_742_v0, 3)
			(_nw153).ArraySet1((_946_v146).F11(), 4)
			(_nw153).ArraySet1((_946_v146).F11(), 5)
			(_nw153).ArraySet1(Companion_Default___.Fm23(Companion_Default___.Fm25(Companion_Default___.Fm23(_937_v139, _dafny.IntOfInt64(924), (_946_v146).F12(), _742_v0, globalState), p1, globalState), _dafny.IntOfInt64(379), _945_cf6, (_946_v146).F11(), globalState), 6)
			(_nw153).ArraySet1(!(_742_v0), 7)
			(_nw153).ArraySet1(_742_v0, 8)
			(_nw153).ArraySet1(_742_v0, 9)
			(_nw153).ArraySet1(_742_v0, 10)
			(_nw153).ArraySet1(_742_v0, 11)
			(_nw153).ArraySet1(_742_v0, 12)
			(_nw153).ArraySet1(false, 13)
			(_nw153).ArraySet1(_742_v0, 14)
			(_nw153).ArraySet1((_946_v146).F11(), 15)
			(_nw153).ArraySet1(false, 16)
			(_nw153).ArraySet1(true, 17)
			(_nw153).ArraySet1((_946_v146).F11(), 18)
			(_nw153).ArraySet1(!((_946_v146).F11()), 19)
			(_nw153).ArraySet1(_742_v0, 20)
			(_nw153).ArraySet1((_946_v146).F11(), 21)
			(_nw153).ArraySet1((_946_v146).F11(), 22)
			(_nw153).ArraySet1((_946_v146).F11(), 23)
			(_nw153).ArraySet1((_946_v146).F11(), 24)
			(_nw153).ArraySet1(_742_v0, 25)
			_947_v147 = _nw153
			var _948_v148 _dafny.MultiSet
			_ = _948_v148
			_948_v148 = _dafny.MultiSetOf((_939_v141).Cardinality())
			var _949_v149 *C3
			_ = _949_v149
			var _nw154 *C3 = New_C3_()
			_ = _nw154
			_nw154.Ctor__(_948_v148, _742_v0)
			_949_v149 = _nw154
			var _950_v150 _dafny.Map
			_ = _950_v150
			_950_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v147, _949_v149)
			_742_v0 = (func() bool {
				if _742_v0 {
					return !(_950_v150).Contains(_947_v147)
				}
				return (_949_v149).F10()
			})()
			var _951_v151 _dafny.MultiSet
			_ = _951_v151
			_951_v151 = _dafny.MultiSetOf((_949_v149).F10(), (_949_v149).F10(), (_946_v146).F11())
			var _952_v152 _dafny.Sequence
			_ = _952_v152
			_952_v152 = _dafny.SeqOf((_951_v151).IsDisjointFrom(_951_v151), (_949_v149).F10(), true, (_946_v146).F11(), (_949_v149).F10())
			if (_952_v152).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_952_v152).Cardinality()))).Uint32()).(bool) {
				var _953_v154 _dafny.Map
				_ = _953_v154
				_953_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_946_v146).F12(), (_dafny.Zero).Minus((_946_v146).F12()))
				var _954_v155 _dafny.Set
				_ = _954_v155
				_954_v155 = _dafny.SetOf((func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter32 := _dafny.Iterate((_953_v154).Keys().Elements()); ; {
						_compr_31, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _955_v153 _dafny.Int
						_955_v153 = interface{}(_compr_31).(_dafny.Int)
						if (_953_v154).Contains(_955_v153) {
							_coll31.Add(Companion_Default___.SafeModuloInt(_955_v153, p1), (_946_v146).F11())
						}
					}
					return _coll31.ToMap()
				}()).Cardinality())
				(globalState).F2 = (_dafny.Zero).Minus(Companion_Default___.Fm24((func() _dafny.Set {
					if true {
						return _954_v155
					}
					return _954_v155
				})(), globalState))
				var _956_v156 _dafny.MultiSet
				_ = _956_v156
				_956_v156 = _dafny.MultiSetOf(_947_v147, _947_v147)
				var _957_v157 _dafny.Sequence
				_ = _957_v157
				_957_v157 = _dafny.SeqOf(_956_v156)
				var _958_v158 _dafny.Sequence
				_ = _958_v158
				_958_v158 = _dafny.SeqOf(_947_v147, _947_v147, _947_v147)
				_956_v156 = ((_957_v157).Select((Companion_Default___.SafeIndex(_945_cf6, _dafny.IntOfUint32((_957_v157).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference((_dafny.MultiSetOf(_947_v147, (_958_v158).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-888), _dafny.IntOfUint32((_958_v158).Cardinality()))).Uint32()).(_dafny.Array), _947_v147)).Union(_956_v156))
				var _959_v159 _dafny.Sequence
				_ = _959_v159
				_959_v159 = _dafny.SeqOf((_949_v149).F10())
				(globalState).F2 = _dafny.IntOfUint32((_959_v159).Cardinality())
				_945_cf6 = (_946_v146).F12()
				_937_v139 = Companion_Default___.Fm8(_dafny.IntOfInt64(826), globalState)
			} else {
				(globalState).F2 = (_dafny.Zero).Minus((_946_v146).F12())
				var _960_v160 _dafny.Map
				_ = _960_v160
				_960_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_949_v149).F10(), _dafny.Companion_Sequence_.Equal(_937_v139, (_813_v57).Dtor_cf22()))
				_742_v0 = (func() bool {
					if (_960_v160).Contains((_946_v146).F11()) {
						return (_960_v160).Get((_946_v146).F11()).(bool)
					}
					return !(!(((func() bool {
						if (_939_v141).Contains((_dafny.Zero).Minus((_946_v146).F12())) {
							return (_939_v141).Get((_dafny.Zero).Minus((_946_v146).F12())).(bool)
						}
						return (_949_v149).F10()
					})()) && ((_949_v149).F10())))
				})()
				var _961_v161 _dafny.Array
				_ = _961_v161
				var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw155
				_961_v161 = _nw155
				_961_v161 = _961_v161
				_960_v160 = (_960_v160).Update(_dafny.Companion_Sequence_.Contains(_937_v139, (_this).F6()), (_946_v146).F11())
				var _962_v162 _dafny.Map
				_ = _962_v162
				_962_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_946_v146).F12(), p0)
				var _963_v164 _dafny.Map
				_ = _963_v164
				_963_v164 = _962_v162
				var _964_v165 _dafny.Array
				_ = _964_v165
				var _nwElement0_42 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_946_v146).F12(), _dafny.IntOfUint32((_937_v139).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(57), p1))
				_ = _nwElement0_42
				var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(5))
				_ = _nw156
				(_nw156).ArraySet1(_nwElement0_42, 0)
				(_nw156).ArraySet1(_962_v162, 1)
				(_nw156).ArraySet1(func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(572), _dafny.IntOfInt64(990))); ; {
						_compr_32, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _965_v163 _dafny.Int
						_965_v163 = interface{}(_compr_32).(_dafny.Int)
						if ((_dafny.IntOfInt64(572)).Cmp(_965_v163) <= 0) && ((_965_v163).Cmp(_dafny.IntOfInt64(990)) < 0) {
							_coll32.Add((_965_v163).Times(p1), p0)
						}
					}
					return _coll32.ToMap()
				}(), 2)
				(_nw156).ArraySet1((_962_v162).Merge((_963_v164)), 3)
				(_nw156).ArraySet1(_962_v162, 4)
				_964_v165 = _nw156
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_964_v165), 0))
				_ = _index130
				(_964_v165).ArraySet1(_962_v162, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_964_v165), 0))
				_ = _index131
				(_964_v165).ArraySet1((_962_v162).Update(p1, (func() _dafny.Int {
					if (_949_v149).F10() {
						return _dafny.IntOfUint32((_952_v152).Cardinality())
					}
					return _dafny.IntOfInt64(-935)
				})()), (_index131).Int())
			}
			if (_949_v149).F10() {
				var _966_v166 _dafny.Int
				_ = _966_v166
				var _967_v167 _dafny.CodePoint
				_ = _967_v167
				var _out35 _dafny.Int
				_ = _out35
				var _out36 _dafny.CodePoint
				_ = _out36
				_out35, _out36 = (_946_v146).M1(_938_v140, ((_946_v146).F11()) || ((_946_v146).F11()), _937_v139, globalState)
				_966_v166 = _out35
				_967_v167 = _out36
				(globalState).F2 = (p0).Minus((_946_v146).F12())
				var _968_v168 bool
				_ = _968_v168
				var _969_v169 bool
				_ = _969_v169
				var _970_v170 _dafny.CodePoint
				_ = _970_v170
				var _971_v171 bool
				_ = _971_v171
				var _out37 bool
				_ = _out37
				var _out38 bool
				_ = _out38
				var _out39 _dafny.CodePoint
				_ = _out39
				var _out40 bool
				_ = _out40
				_out37, _out38, _out39, _out40 = (_949_v149).M5(p1, (_949_v149).F10(), globalState)
				_968_v168 = _out37
				_969_v169 = _out38
				_970_v170 = _out39
				_971_v171 = _out40
				var _972_v172 D9
				_ = _972_v172
				_972_v172 = Companion_D9_.Create_DC19_(_941_v143, (_dafny.Zero).Minus((func() _dafny.Int {
					if (_948_v148).Contains((_946_v146).F12()) {
						return (_948_v148).Multiplicity((_946_v146).F12())
					}
					return (_951_v151).Cardinality()
				})()), Companion_Default___.Fm9(_dafny.IntOfUint32((_938_v140).Cardinality()), globalState))
				_972_v172 = _972_v172
				var _973_v173 _dafny.Array
				_ = _973_v173
				var _nwElement0_43 _dafny.MultiSet = _951_v151
				_ = _nwElement0_43
				var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(2))
				_ = _nw157
				(_nw157).ArraySet1(_nwElement0_43, 0)
				(_nw157).ArraySet1(_951_v151, 1)
				_973_v173 = _nw157
				var _974_v174 *C0
				_ = _974_v174
				var _nw158 *C0 = New_C0_()
				_ = _nw158
				_nw158.Ctor__(p1)
				_974_v174 = _nw158
				var _975_v175 _dafny.Set
				_ = _975_v175
				_975_v175 = _dafny.SetOf((_946_v146).F11())
				var _rhs127 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _971_v171 {
						return _967_v167
					}
					return _970_v170
				})()
				_ = _rhs127
				var _rhs128 bool = ((_975_v175).Union(_975_v175)).Equals(_975_v175)
				_ = _rhs128
				var _rhs129 _dafny.Array = _973_v173
				_ = _rhs129
				var _rhs130 *C0 = _974_v174
				_ = _rhs130
				_967_v167 = _rhs127
				_968_v168 = _rhs128
				_973_v173 = _rhs129
				_974_v174 = _rhs130
			} else {
				var _976_v176 _dafny.Set
				_ = _976_v176
				_976_v176 = _dafny.SetOf((_946_v146).F12(), (_938_v140).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), _dafny.IntOfUint32((_938_v140).Cardinality()))).Uint32()).(_dafny.Int))
				var _977_v177 _dafny.Map
				_ = _977_v177
				_977_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_946_v146).F11(), _976_v176)
				var _978_v178 _dafny.Map
				_ = _978_v178
				_978_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_938_v140, _938_v140), (Companion_Default___.Fm38((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_949_v149).F10(), (_952_v152).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.IntOfUint32((_952_v152).Cardinality()))).Uint32()).(bool))).Cardinality(), (_952_v152).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_952_v152).Cardinality()))).Uint32()).(bool), (_949_v149).F10(), (_949_v149).F10(), globalState)).Union((func() _dafny.Set {
					if (_977_v177).Contains((_949_v149).F10()) {
						return (_977_v177).Get((_949_v149).F10()).(_dafny.Set)
					}
					return _dafny.SetOf((_946_v146).F12(), p1)
				})()))
				_977_v177 = (_977_v177).Update(!(!(_976_v176).Contains(_945_cf6)), (func() _dafny.Set {
					if false {
						return func() _dafny.Set {
							var _coll33 = _dafny.NewBuilder()
							_ = _coll33
							for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(121), _dafny.IntOfInt64(74))); ; {
								_compr_33, _ok34 := _iter34()
								if !_ok34 {
									break
								}
								var _979_v179 _dafny.Int
								_979_v179 = interface{}(_compr_33).(_dafny.Int)
								if ((_dafny.IntOfInt64(121)).Cmp(_979_v179) <= 0) && ((_979_v179).Cmp(_dafny.IntOfInt64(74)) < 0) {
									_coll33.Add(Companion_Default___.SafeModuloInt(_979_v179, (func() _dafny.Map {
										var _coll34 = _dafny.NewMapBuilder()
										_ = _coll34
										for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(124), _dafny.IntOfInt64(935))); ; {
											_compr_34, _ok35 := _iter35()
											if !_ok35 {
												break
											}
											var _980_v180 _dafny.Int
											_980_v180 = interface{}(_compr_34).(_dafny.Int)
											if ((_dafny.IntOfInt64(124)).Cmp(_980_v180) <= 0) && ((_980_v180).Cmp(_dafny.IntOfInt64(935)) < 0) {
												_coll34.Add((_980_v180).Minus(_945_cf6), p1)
											}
										}
										return _coll34.ToMap()
									}()).Cardinality()))
								}
							}
							return _coll33.ToSet()
						}()
					}
					return _976_v176
				})())
				var _981_v181 _dafny.Set
				_ = _981_v181
				_981_v181 = _dafny.SetOf((_this).F6())
				var _982_v182 _dafny.Set
				_ = _982_v182
				_982_v182 = _dafny.SetOf((_949_v149).F10(), (_981_v181).IsDisjointFrom(_981_v181), false)
				_982_v182 = _982_v182
				(globalState).F2 = _945_cf6
				_976_v176 = _976_v176
				var _983_v183 _dafny.Set
				_ = _983_v183
				_983_v183 = _982_v182
				_983_v183 = (_dafny.SetOf((_949_v149).F10())).Difference(_982_v182)
			}
		} else if _source16.Is_DC5() {
			var _984___mcc_h10 _dafny.Int = _source16.Get_().(D2_DC5).Cf7
			_ = _984___mcc_h10
			var _985___mcc_h11 bool = _source16.Get_().(D2_DC5).Cf8
			_ = _985___mcc_h11
			var _986___mcc_h12 bool = _source16.Get_().(D2_DC5).Cf9
			_ = _986___mcc_h12
			var _987___mcc_h13 bool = _source16.Get_().(D2_DC5).Cf10
			_ = _987___mcc_h13
			var _988___mcc_h14 bool = _source16.Get_().(D2_DC5).Cf11
			_ = _988___mcc_h14
			var _989_cf11 bool = _988___mcc_h14
			_ = _989_cf11
			var _990_cf10 bool = _987___mcc_h13
			_ = _990_cf10
			var _991_cf9 bool = _986___mcc_h12
			_ = _991_cf9
			var _992_cf8 bool = _985___mcc_h11
			_ = _992_cf8
			var _993_cf7 _dafny.Int = _984___mcc_h10
			_ = _993_cf7
			var _994_v184 _dafny.Array
			_ = _994_v184
			var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(16))
			_ = _nw159
			_994_v184 = _nw159
			var _995_v185 _dafny.Array
			_ = _995_v185
			var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw160
			_995_v185 = _nw160
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_994_v184), 0))
			_ = _index132
			(_994_v184).ArraySet1((_dafny.MultiSetOf(_995_v185, _995_v185, _995_v185)).Update(_995_v185, Companion_Default___.Abs(p1)), (_index132).Int())
			var _996_v186 _dafny.MultiSet
			_ = _996_v186
			_996_v186 = _dafny.MultiSetOf(_995_v185, _995_v185, _995_v185, _995_v185)
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_994_v184), 0))
			_ = _index133
			(_994_v184).ArraySet1((func() _dafny.MultiSet {
				if _992_cf8 {
					return (_996_v186).Difference(_996_v186)
				}
				return (_996_v186).Union(_996_v186)
			})(), (_index133).Int())
			var _997_v187 _dafny.MultiSet
			_ = _997_v187
			_997_v187 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("polxqwadi"))
			_997_v187 = ((_997_v187).Union(_997_v187)).Update(_dafny.UnicodeSeqOfUtf8Bytes("uiiokhs"), Companion_Default___.Abs(_dafny.IntOfInt64(-980)))
			_937_v139 = (_813_v57).Dtor_cf22()
			var _998_v188 _dafny.Array
			_ = _998_v188
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_24
			var _nw161 _dafny.Array
			_ = _nw161
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw161 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Set = (func(_999_cf8 bool) func(_dafny.Int) _dafny.Set {
					return func(_1000_i17 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(_999_cf8)
					}
				})(_992_cf8)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw161 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw161).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw161).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_998_v188 = _nw161
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_998_v188), 0))
			_ = _index134
			(_998_v188).ArraySet1(Companion_Default___.Fm16(p0, _dafny.CodePoint('l'), _993_cf7, globalState), (_index134).Int())
			var _1001_v189 _dafny.Set
			_ = _1001_v189
			_1001_v189 = _dafny.SetOf(p1)
			var _1002_v190 _dafny.Sequence
			_ = _1002_v190
			_1002_v190 = _dafny.SeqOf(!(false))
			var _1003_v191 _dafny.Set
			_ = _1003_v191
			_1003_v191 = _dafny.SetOf((_1001_v189).IsSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(544))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_1004_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1005_i18 _dafny.Int) _dafny.Int {
					return _1004_p1
				}
			})(p1)))).Cardinality()))), _990_cf10, (_1002_v190).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1002_v190).Cardinality()))).Uint32()).(bool))
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_998_v188), 0))
			_ = _index135
			(_998_v188).ArraySet1(_1003_v191, (_index135).Int())
		} else {
			var _1006___mcc_h15 _dafny.Map = _source16.Get_().(D2_DC3).Cf5
			_ = _1006___mcc_h15
			var _1007_cf5 _dafny.Map = _1006___mcc_h15
			_ = _1007_cf5
			var _1008_v192 *C2
			_ = _1008_v192
			var _nw162 *C2 = New_C2_()
			_ = _nw162
			_nw162.Ctor__(_742_v0, Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(130)), _940_v142)
			_1008_v192 = _nw162
			if (_1008_v192).F11() {
				_742_v0 = _742_v0
				var _1009_v193 _dafny.Array
				_ = _1009_v193
				var _nwElement0_44 bool = false
				_ = _nwElement0_44
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(20))
				_ = _nw163
				(_nw163).ArraySet1(_nwElement0_44, 0)
				(_nw163).ArraySet1((_1008_v192).F11(), 1)
				(_nw163).ArraySet1(_742_v0, 2)
				(_nw163).ArraySet1(_742_v0, 3)
				(_nw163).ArraySet1((_1008_v192).F11(), 4)
				(_nw163).ArraySet1((_1008_v192).F11(), 5)
				(_nw163).ArraySet1((_1008_v192).F11(), 6)
				(_nw163).ArraySet1((_1008_v192).F11(), 7)
				(_nw163).ArraySet1((_1008_v192).F11(), 8)
				(_nw163).ArraySet1((_1008_v192).F11(), 9)
				(_nw163).ArraySet1(false, 10)
				(_nw163).ArraySet1((_1008_v192).F11(), 11)
				(_nw163).ArraySet1((_1008_v192).F11(), 12)
				(_nw163).ArraySet1((_1008_v192).F11(), 13)
				(_nw163).ArraySet1(true, 14)
				(_nw163).ArraySet1(_742_v0, 15)
				(_nw163).ArraySet1(_742_v0, 16)
				(_nw163).ArraySet1(_742_v0, 17)
				(_nw163).ArraySet1((_1008_v192).F11(), 18)
				(_nw163).ArraySet1((_1008_v192).F11(), 19)
				_1009_v193 = _nw163
				var _1010_v194 _dafny.Map
				_ = _1010_v194
				_1010_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC15_(_1009_v193), Companion_D9_.Create_DC19_(_941_v143, p0, (_1008_v192).F12()))
				_939_v141 = (_939_v141).Update((_dafny.Zero).Minus(((_1010_v194).Cardinality()).Plus((_1008_v192).F12())), (_1008_v192).F11())
				var _1011_v195 D8
				_ = _1011_v195
				_1011_v195 = Companion_D8_.Create_DC15_(_1009_v193)
				var _pat_let_tv24 = _1009_v193
				_ = _pat_let_tv24
				_1011_v195 = func(_pat_let11_0 D8) D8 {
					return func(_1012_dt__update__tmp_h3 D8) D8 {
						return func(_pat_let12_0 _dafny.Array) D8 {
							return func(_1013_dt__update_hcf27_h0 _dafny.Array) D8 {
								return Companion_D8_.Create_DC15_(_1013_dt__update_hcf27_h0)
							}(_pat_let12_0)
						}(_pat_let_tv24)
					}(_pat_let11_0)
				}(_1011_v195)
				_742_v0 = (p0).Cmp((_1008_v192).F12()) <= 0
				_742_v0 = (func() bool {
					if ((_dafny.Zero).Minus((_938_v140).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_938_v140).Cardinality()))).Uint32()).(_dafny.Int))).Cmp(_dafny.IntOfUint32((_937_v139).Cardinality())) >= 0 {
						return (_1008_v192).F11()
					}
					return ((_1008_v192).F11()) == (Companion_Default___.Fm23(_937_v139, (_dafny.MultiSetOf((_1008_v192).F12())).Cardinality(), p1, (_1008_v192).F11(), globalState))
				})()
			} else {
				var _rhs131 bool = !(!(!((_1008_v192).F11())))
				_ = _rhs131
				var _rhs132 _dafny.Sequence = _937_v139
				_ = _rhs132
				_742_v0 = _rhs131
				_937_v139 = _rhs132
				(globalState).F2 = (((_dafny.SetOf(_dafny.IntOfInt64(36), _dafny.IntOfUint32((_937_v139).Cardinality()))).Intersection(func() _dafny.Set {
					var _coll35 = _dafny.NewBuilder()
					_ = _coll35
					for _iter36 := _dafny.Iterate((_dafny.SeqOf(p1)).Elements()); ; {
						_compr_35, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1014_v196 _dafny.Int
						_1014_v196 = interface{}(_compr_35).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p1), _1014_v196) {
							_coll35.Add((_1014_v196).Times(_dafny.IntOfInt64(489)))
						}
					}
					return _coll35.ToSet()
				}())).Cardinality()).Minus(p1)
				_742_v0 = _742_v0
				var _1015_v197 _dafny.Map
				_ = _1015_v197
				_1015_v197 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1008_v192).F12(), (_939_v141).Cardinality())
				(globalState).F2 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_1015_v197).Contains((_1008_v192).F12()) {
						return (_1015_v197).Get((_1008_v192).F12()).(_dafny.Int)
					}
					return (_1008_v192).F12()
				})(), p1)
				var _1016_v198 _dafny.CodePoint
				_ = _1016_v198
				_1016_v198 = _dafny.CodePoint('e')
				_1016_v198 = (_this).F6()
			}
			(globalState).F2 = (_dafny.Zero).Minus(((_1008_v192).F12()).Plus((_dafny.SetOf(false)).Cardinality()))
			(globalState).F2 = ((p1).Plus(p1)).Times(((_dafny.Zero).Minus(_dafny.IntOfInt64(-916))).Minus(p0))
		}
		var _1017_v199 _dafny.Array
		_ = _1017_v199
		var _nw164 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw164
		_1017_v199 = _nw164
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1017_v199), 0))
		_ = _index136
		(_1017_v199).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-512))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_1018_i19 _dafny.Int) _dafny.CodePoint {
			return (_this).F6()
		})), _937_v139), (_index136).Int())
		var _1019_v200 _dafny.Map
		_ = _1019_v200
		_1019_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _1020_v201 D3
		_ = _1020_v201
		_1020_v201 = Companion_D3_.Create_DC6_((_this).F6())
		var _1021_v202 _dafny.Array
		_ = _1021_v202
		var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw165
		_1021_v202 = _nw165
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1021_v202), 0))
		_ = _index137
		(_1021_v202).ArraySet1(p1, (_index137).Int())
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1017_v199), 0))
		_ = _index138
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1021_v202), 0))
		_ = _index139
		var _rhs133 bool = _742_v0
		_ = _rhs133
		var _rhs134 _dafny.Map = _1019_v200
		_ = _rhs134
		var _rhs135 _dafny.Int = ((p0).Minus(p1)).Minus(((_dafny.Zero).Minus(p1)).Minus(_dafny.IntOfInt64(468)))
		_ = _rhs135
		var _rhs136 D3 = _1020_v201
		_ = _rhs136
		var _rhs137 _dafny.Int = p1
		_ = _rhs137
		var _lhs77 _dafny.Array = _1017_v199
		_ = _lhs77
		var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1017_v199), 0))
		_ = _lhs78
		var _lhs79 *GlobalState = globalState
		_ = _lhs79
		var _lhs80 _dafny.Array = _1021_v202
		_ = _lhs80
		var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1021_v202), 0))
		_ = _lhs81
		(_lhs77).ArraySet1(_rhs133, (_lhs78).Int())
		_1019_v200 = _rhs134
		_lhs79.F2 = _rhs135
		_1020_v201 = _rhs136
		(_lhs80).ArraySet1(_rhs137, (_lhs81).Int())
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1017_v199), 0))
		_ = _index140
		(_1017_v199).ArraySet1(!(true), (_index140).Int())
		var _1022_v203 _dafny.MultiSet
		_ = _1022_v203
		_1022_v203 = _dafny.MultiSetOf(p0)
		var _1023_v204 *C3
		_ = _1023_v204
		var _nw166 *C3 = New_C3_()
		_ = _nw166
		_nw166.Ctor__(_1022_v203, _742_v0)
		_1023_v204 = _nw166
		var _1024_v205 _dafny.Map
		_ = _1024_v205
		_1024_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1023_v204, (_this).F6())
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1017_v199), 0))
		_ = _index141
		(_1017_v199).ArraySet1(!(_1024_v205).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1023_v204, (Companion_D3_.Create_DC6_((_this).F6())).Dtor_cf12())), (_index141).Int())
	}
}
func (_this *C5) F6() _dafny.CodePoint {
	{
		return _this._f6
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	F5  _dafny.Int
	_f4 _dafny.Map
}

func New_C6_() *C6 {
	_this := C6{}

	_this.F5 = _dafny.Zero
	_this._f4 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__(f4 _dafny.Map, f5 _dafny.Int) {
	{
		(_this)._f4 = f4
		(_this).F5 = f5
	}
}
func (_this *C6) M2(p0 bool, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _1025_i0 _dafny.Int
		_ = _1025_i0
		_1025_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_1025_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1025_i0 = (_1025_i0).Plus(_dafny.One)
					var _1026_v0 _dafny.Array
					_ = _1026_v0
					var _len0_25 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_25
					var _nw167 _dafny.Array
					_ = _nw167
					if _len0_25.Cmp(_dafny.Zero) == 0 {
						_nw167 = _dafny.NewArray(_len0_25)
					} else {
						var _init25 func(_dafny.Int) D1 = (func(_1027_p0 bool) func(_dafny.Int) D1 {
							return func(_1028_i1 _dafny.Int) D1 {
								return Companion_D1_.Create_DC2_(_this.F5, _1027_p0, _this.F5)
							}
						})(p0)
						_ = _init25
						var _element0_25 = _init25(_dafny.Zero)
						_ = _element0_25
						_nw167 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
						(_nw167).ArraySet1(_element0_25, 0)
						var _nativeLen0_25 = (_len0_25).Int()
						_ = _nativeLen0_25
						for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
							(_nw167).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
						}
					}
					_1026_v0 = _nw167
					var _1029_v1 D1
					_ = _1029_v1
					_1029_v1 = Companion_D1_.Create_DC2_(_this.F5, p0, _this.F5)
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1026_v0), 0))
					_ = _index142
					(_1026_v0).ArraySet1(_1029_v1, (_index142).Int())
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1026_v0), 0))
					_ = _index143
					(_1026_v0).ArraySet1(_1029_v1, (_index143).Int())
					var _1030_v2 _dafny.CodePoint
					_ = _1030_v2
					_1030_v2 = _dafny.CodePoint('x')
					var _1031_v3 *C5
					_ = _1031_v3
					var _nw168 *C5 = New_C5_()
					_ = _nw168
					_nw168.Ctor__(_1030_v2)
					_1031_v3 = _nw168
					var _1032_v4 _dafny.Map
					_ = _1032_v4
					_1032_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
					var _1033_v5 _dafny.Sequence
					_ = _1033_v5
					_1033_v5 = _dafny.SeqOf((_dafny.Zero).Minus(_this.F5), _this.F5)
					var _1034_v6 _dafny.Map
					_ = _1034_v6
					_1034_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1032_v4, _1033_v5)
					_1034_v6 = (_1034_v6).Update(Companion_Default___.Fm39(false, globalState), _1033_v5)
					var _1035_v7 _dafny.Array
					_ = _1035_v7
					var _nw169 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
					_ = _nw169
					_1035_v7 = _nw169
					var _1036_v8 _dafny.Set
					_ = _1036_v8
					_1036_v8 = _dafny.SetOf(_1035_v7, _1035_v7, _1035_v7, _1035_v7, _1035_v7)
					var _1037_v9 _dafny.Sequence
					_ = _1037_v9
					_1037_v9 = _dafny.UnicodeSeqOfUtf8Bytes("eueyat")
					var _1038_v10 D6
					_ = _1038_v10
					_1038_v10 = Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("kuxdxnku"))
					var _1039_v11 _dafny.Map
					_ = _1039_v11
					_1039_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1036_v8, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1037_v9, (_1038_v10).Dtor_cf22())).Cardinality()))
					var _1040_v12 _dafny.Sequence
					_ = _1040_v12
					_1040_v12 = _dafny.SeqOf(_1031_v3, _1031_v3)
					var _1041_v13 _dafny.MultiSet
					_ = _1041_v13
					_1041_v13 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1040_v12).Cardinality()), _this.F5)
					var _1042_v14 _dafny.Sequence
					_ = _1042_v14
					_1042_v14 = _dafny.SeqOf(p0)
					_1039_v11 = (_1039_v11).Update((_1036_v8).Difference(_1036_v8), (func() _dafny.Int {
						if p0 {
							return (_1041_v13).Cardinality()
						}
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1033_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1042_v14).Cardinality()), _dafny.IntOfUint32((_1033_v5).Cardinality()))).Uint32()).(_dafny.Int), _1032_v4)).Cardinality()
					})())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1043_v15 _dafny.MultiSet
		_ = _1043_v15
		_1043_v15 = _dafny.MultiSetOf(true, !(p0), p0, p0)
		var _1044_v16 _dafny.CodePoint
		_ = _1044_v16
		_1044_v16 = _dafny.CodePoint('x')
		var _1045_v17 _dafny.Map
		_ = _1045_v17
		_1045_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(884), Companion_Default___.Fm21(_this.F5, _1044_v16, p0, false, globalState))
		var _1046_v18 _dafny.Sequence
		_ = _1046_v18
		_1046_v18 = _dafny.UnicodeSeqOfUtf8Bytes("tbwc")
		var _1047_v19 _dafny.MultiSet
		_ = _1047_v19
		_1047_v19 = _dafny.MultiSetOf(((_1043_v15).Cardinality()).Times(_this.F5), _this.F5, Companion_Default___.SafeModuloInt((_1045_v17).Cardinality(), _dafny.IntOfUint32((_1046_v18).Cardinality())))
		var _1048_v20 _dafny.Array
		_ = _1048_v20
		var _nw170 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(19))
		_ = _nw170
		_1048_v20 = _nw170
		var _1049_v21 _dafny.Sequence
		_ = _1049_v21
		_1049_v21 = _dafny.SeqOf(_1048_v20)
		var _rhs138 _dafny.MultiSet = (_1047_v19).Intersection(_1047_v19)
		_ = _rhs138
		var _rhs139 _dafny.Int = _this.F5
		_ = _rhs139
		var _rhs140 _dafny.Sequence = _dafny.SeqOf(_1048_v20)
		_ = _rhs140
		var _lhs82 *C6 = _this
		_ = _lhs82
		_1047_v19 = _rhs138
		_lhs82.F5 = _rhs139
		_1049_v21 = _rhs140
		var _1050_v22 _dafny.Array
		_ = _1050_v22
		var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw171
		_1050_v22 = _nw171
		var _1051_v23 _dafny.Array
		_ = _1051_v23
		var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw172
		_1051_v23 = _nw172
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1050_v22), 0))
		_ = _index144
		(_1050_v22).ArraySet1(_1051_v23, (_index144).Int())
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1050_v22), 0))
		_ = _index145
		(_1050_v22).ArraySet1(_1051_v23, (_index145).Int())
		var _hi4 _dafny.Int = _this.F5
		_ = _hi4
		for _1052_i2 := ((_dafny.Zero).Minus(_this.F5)).Minus(_this.F5); _1052_i2.Cmp(_hi4) < 0; _1052_i2 = _1052_i2.Plus(_dafny.One) {
			r1 = p0
			_1043_v15 = _1043_v15
			var _1053_v24 _dafny.Array
			_ = _1053_v24
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_26
			var _nw173 _dafny.Array
			_ = _nw173
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw173 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.MultiSet = (func(_1054_v19 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1055_i3 _dafny.Int) _dafny.MultiSet {
						return (_1054_v19).Difference(_1054_v19)
					}
				})(_1047_v19)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw173 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw173).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw173).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1053_v24 = _nw173
			_1053_v24 = _1053_v24
			var _1056_v25 _dafny.Sequence
			_ = _1056_v25
			_1056_v25 = _dafny.SeqOf(((_this).F4()).Cardinality(), _1052_i2, _1052_i2, _this.F5, _dafny.IntOfUint32((_1046_v18).Cardinality()))
			var _1057_v26 _dafny.Set
			_ = _1057_v26
			_1057_v26 = _dafny.SetOf(false, p0)
			var _1058_v27 _dafny.Map
			_ = _1058_v27
			_1058_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F5, _this.F5)
			var _1059_v28 _dafny.Array
			_ = _1059_v28
			var _nwElement0_45 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1056_v25, Companion_Default___.Fm12(p0, _dafny.IntOfUint32((_1056_v25).Cardinality()), _1057_v26, _1058_v27, globalState))
			_ = _nwElement0_45
			var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(21))
			_ = _nw174
			(_nw174).ArraySet1(_nwElement0_45, 0)
			(_nw174).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_1060_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1061_i4 _dafny.Int) _dafny.Int {
					return _1060_i2
				}
			})(_1052_i2))), 1)
			(_nw174).ArraySet1(Companion_Default___.Fm12(p0, _dafny.IntOfUint32((_1046_v18).Cardinality()), _1057_v26, _1058_v27, globalState), 2)
			(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1056_v25, _1056_v25), 3)
			(_nw174).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(240))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_1062_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1063_i5 _dafny.Int) _dafny.Int {
					return _1062_i2
				}
			})(_1052_i2))), 4)
			(_nw174).ArraySet1(_dafny.SeqOf(_this.F5), 5)
			(_nw174).ArraySet1(_1056_v25, 6)
			(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F5, _1052_i2), (Companion_Default___.SafeIndex(_1052_i2, _dafny.IntOfUint32((_dafny.SeqOf(_this.F5, _1052_i2)).Cardinality()))).Uint32(), _this.F5), _1056_v25), 7)
			(_nw174).ArraySet1(_1056_v25, 8)
			(_nw174).ArraySet1(_1056_v25, 9)
			(_nw174).ArraySet1(_1056_v25, 10)
			(_nw174).ArraySet1(_1056_v25, 11)
			(_nw174).ArraySet1(_1056_v25, 12)
			(_nw174).ArraySet1(_1056_v25, 13)
			(_nw174).ArraySet1(_1056_v25, 14)
			(_nw174).ArraySet1(_dafny.SeqOf(_1052_i2), 15)
			(_nw174).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_1056_v25).Cardinality()), _1052_i2, _1052_i2, _1052_i2), 16)
			(_nw174).ArraySet1(_1056_v25, 17)
			(_nw174).ArraySet1(_1056_v25, 18)
			(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1056_v25, _dafny.SeqOf(_dafny.IntOfInt64(729))), 19)
			(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1056_v25, _1056_v25), 20)
			_1059_v28 = _nw174
			var _1064_v29 _dafny.Map
			_ = _1064_v29
			_1064_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1056_v25).Select((Companion_Default___.SafeIndex(_1052_i2, _dafny.IntOfUint32((_1056_v25).Cardinality()))).Uint32()).(_dafny.Int), _1059_v28)
			_1059_v28 = (func() _dafny.Array {
				if (_1064_v29).Contains(_this.F5) {
					return (_1064_v29).Get(_this.F5).(_dafny.Array)
				}
				return _1059_v28
			})()
		}
		var _1065_i6 _dafny.Int
		_ = _1065_i6
		_1065_i6 = _dafny.Zero
		{
			for !(p0) {
				{
					if (_1065_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1065_i6 = (_1065_i6).Plus(_dafny.One)
					var _1066_v30 _dafny.Map
					_ = _1066_v30
					_1066_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1051_v23)
					var _1067_v31 _dafny.Array
					_ = _1067_v31
					var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
					_ = _nw175
					_1067_v31 = _nw175
					var _1068_v32 *C2
					_ = _1068_v32
					var _nw176 *C2 = New_C2_()
					_ = _nw176
					_nw176.Ctor__(Companion_Default___.Fm21(_this.F5, _1044_v16, p0, p0, globalState), _this.F5, _1067_v31)
					_1068_v32 = _nw176
					var _1069_v33 _dafny.Sequence
					_ = _1069_v33
					_1069_v33 = _dafny.SeqOf(((_1066_v30).Update(p0, _dafny.ArrayCastTo((_1050_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_1050_v22), 0))).Int())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_1068_v32)).Cardinality()), _this.F5, (_1068_v32).F12())
					var _1070_v34 _dafny.Sequence
					_ = _1070_v34
					_1070_v34 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_1069_v33))
					var _1071_v35 *C3
					_ = _1071_v35
					var _nw177 *C3 = New_C3_()
					_ = _nw177
					_nw177.Ctor__((_dafny.MultiSetOf((_1047_v19).Cardinality(), _this.F5)).Union((_1070_v34).Select((Companion_Default___.SafeIndex((_1068_v32).F12(), _dafny.IntOfUint32((_1070_v34).Cardinality()))).Uint32()).(_dafny.MultiSet)), p0)
					_1071_v35 = _nw177
					var _1072_v36 _dafny.Sequence
					_ = _1072_v36
					_1072_v36 = _dafny.SeqOf((_1068_v32).F11())
					var _1073_v37 D4
					_ = _1073_v37
					_1073_v37 = Companion_D4_.Create_DC9_((_dafny.MultiSetFromSeq(_1072_v36)).Cardinality(), _1044_v16, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_1068_v32).F12(), _this.F5)).Cardinality())), _this.F5, p0)
					_1044_v16 = (_1073_v37).Dtor_cf17()
					var _1074_v38 _dafny.Sequence
					_ = _1074_v38
					_1074_v38 = _dafny.SeqOf(Companion_Default___.Fm40((_1071_v35).F10(), (_1068_v32).F12(), _1044_v16, (_1068_v32).F11(), globalState), (_dafny.MultiSetOf((_1068_v32).F11())).Update((_1068_v32).F11(), Companion_Default___.Abs(_this.F5)), _1043_v15)
					var _1075_v39 _dafny.Set
					_ = _1075_v39
					_1075_v39 = _dafny.SetOf((_1068_v32).F12())
					var _1076_v40 _dafny.Sequence
					_ = _1076_v40
					_1076_v40 = _dafny.SeqOf(_1075_v39, _1075_v39, _1075_v39, _1075_v39)
					var _1077_v41 _dafny.Array
					_ = _1077_v41
					var _nwElement0_46 _dafny.Int = (_dafny.Zero).Minus((_1068_v32).F12())
					_ = _nwElement0_46
					var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(5))
					_ = _nw178
					(_nw178).ArraySet1(_nwElement0_46, 0)
					(_nw178).ArraySet1(((_1043_v15).Intersection((_1074_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.IntOfUint32((_1074_v38).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality(), 1)
					(_nw178).ArraySet1(_dafny.IntOfUint32((_1076_v40).Cardinality()), 2)
					(_nw178).ArraySet1(_this.F5, 3)
					(_nw178).ArraySet1((_1068_v32).F12(), 4)
					_1077_v41 = _nw178
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1077_v41), 0))
					_ = _index146
					(_1077_v41).ArraySet1(_this.F5, (_index146).Int())
					var _1078_v42 _dafny.Set
					_ = _1078_v42
					_1078_v42 = _dafny.SetOf(p0)
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1077_v41), 0))
					_ = _index147
					var _rhs141 _dafny.Int = (_1068_v32).F12()
					_ = _rhs141
					var _rhs142 _dafny.Int = (((_1078_v42).Intersection(_1078_v42)).Difference(_dafny.SetOf(true, false))).Cardinality()
					_ = _rhs142
					var _rhs143 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}((func(_1079_v32 *C2, _1080_v17 _dafny.Map) func(_dafny.Int) _dafny.Int {
						return func(_1081_i7 _dafny.Int) _dafny.Int {
							return ((_1079_v32).F12()).Times((_1080_v17).Cardinality())
						}
					})(_1068_v32, _1045_v17))), (Companion_Default___.SafeIndex(_this.F5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}((func(_1082_v32 *C2, _1083_v17 _dafny.Map) func(_dafny.Int) _dafny.Int {
						return func(_1084_i7 _dafny.Int) _dafny.Int {
							return ((_1082_v32).F12()).Times((_1083_v17).Cardinality())
						}
					})(_1068_v32, _1045_v17)))).Cardinality()))).Uint32(), (_dafny.Zero).Minus(((_dafny.Zero).Minus(_this.F5)).Times((_1068_v32).F12())))).Cardinality())
					_ = _rhs143
					var _rhs144 bool = (_1071_v35).F10()
					_ = _rhs144
					var _lhs83 *GlobalState = globalState
					_ = _lhs83
					var _lhs84 *C6 = _this
					_ = _lhs84
					var _lhs85 _dafny.Array = _1077_v41
					_ = _lhs85
					var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1077_v41), 0))
					_ = _lhs86
					_lhs83.F2 = _rhs141
					_lhs84.F5 = _rhs142
					(_lhs85).ArraySet1(_rhs143, (_lhs86).Int())
					r1 = _rhs144
					var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_1051_v23), 0))
					_ = _index148
					(_1051_v23).ArraySet1((_this.F5).Cmp((_1077_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1077_v41), 0))).Int()).(_dafny.Int)) != 0, (_index148).Int())
					var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_1051_v23), 0))
					_ = _index149
					(_1051_v23).ArraySet1((func() bool {
						if ((_1068_v32).F12()).Cmp(_this.F5) > 0 {
							return (_1068_v32).F11()
						}
						return _dafny.Companion_Sequence_.Contains(_1046_v18, _1044_v16)
					})(), (_index149).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1085_v43 _dafny.Sequence
		_ = _1085_v43
		_1085_v43 = _dafny.SeqOf((_this.F5).Plus(_this.F5))
		var _1086_v44 _dafny.Map
		_ = _1086_v44
		_1086_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F5, _1085_v43)
		_1085_v43 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1085_v43, (func() _dafny.Sequence {
			if (_1086_v44).Contains(_this.F5) {
				return (_1086_v44).Get(_this.F5).(_dafny.Sequence)
			}
			return _1085_v43
		})()), _1085_v43)
		var _1087_v45 _dafny.Map
		_ = _1087_v45
		_1087_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1046_v18).Cardinality()), _1044_v16)
		var _1088_v46 _dafny.MultiSet
		_ = _1088_v46
		_1088_v46 = _dafny.MultiSetOf(_1087_v45)
		r0 = Companion_Default___.Fm9(((_1088_v46).Update(_1087_v45, Companion_Default___.Abs(_this.F5))).Cardinality(), globalState)
		r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1046_v18, _1046_v18)).Cardinality())).Cmp(_this.F5) <= 0
		return r0, r1
	}
}
func (_this *C6) M3(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var _1089_v0 _dafny.Array
		_ = _1089_v0
		var _nw179 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw179
		_1089_v0 = _nw179
		var _1090_v1 D8
		_ = _1090_v1
		_1090_v1 = Companion_D8_.Create_DC15_(_1089_v0)
		var _source17 D8 = _1090_v1
		_ = _source17
		if _source17.Is_DC16() {
			var _1091___mcc_h0 _dafny.Int = _source17.Get_().(D8_DC16).Cf28
			_ = _1091___mcc_h0
			var _1092___mcc_h1 bool = _source17.Get_().(D8_DC16).Cf29
			_ = _1092___mcc_h1
			var _1093___mcc_h2 bool = _source17.Get_().(D8_DC16).Cf30
			_ = _1093___mcc_h2
			var _1094___mcc_h3 _dafny.Int = _source17.Get_().(D8_DC16).Cf31
			_ = _1094___mcc_h3
			var _1095_cf31 _dafny.Int = _1094___mcc_h3
			_ = _1095_cf31
			var _1096_cf30 bool = _1093___mcc_h2
			_ = _1096_cf30
			var _1097_cf29 bool = _1092___mcc_h1
			_ = _1097_cf29
			var _1098_cf28 _dafny.Int = _1091___mcc_h0
			_ = _1098_cf28
			var _1099_v2 _dafny.Array
			_ = _1099_v2
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw180
			_1099_v2 = _nw180
			var _1100_v3 _dafny.Set
			_ = _1100_v3
			_1100_v3 = _dafny.SetOf(_1097_cf29)
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1099_v2), 0))
			_ = _index150
			(_1099_v2).ArraySet1(((func() _dafny.Set {
				if _1096_cf30 {
					return _1100_v3
				}
				return _1100_v3
			})()).Cardinality(), (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1099_v2), 0))
			_ = _index151
			(_1099_v2).ArraySet1(_1095_cf31, (_index151).Int())
			_1098_cf28 = _1098_cf28
			var _1101_v4 _dafny.CodePoint
			_ = _1101_v4
			_1101_v4 = _dafny.CodePoint('x')
			var _1102_v5 D3
			_ = _1102_v5
			_1102_v5 = Companion_D3_.Create_DC6_(_1101_v4)
			var _1103_v6 *C5
			_ = _1103_v6
			var _nw181 *C5 = New_C5_()
			_ = _nw181
			_nw181.Ctor__((_1102_v5).Dtor_cf12())
			_1103_v6 = _nw181
			var _1104_v7 _dafny.MultiSet
			_ = _1104_v7
			_1104_v7 = _dafny.MultiSetOf((_1099_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1099_v2), 0))).Int()).(_dafny.Int), p1)
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1099_v2), 0))
			_ = _index152
			(_1099_v2).ArraySet1(Companion_Default___.SafeModuloInt(p1, (_this.F5).Times((func() _dafny.Int {
				if (_1104_v7).Contains(_this.F5) {
					return (_1104_v7).Multiplicity(_this.F5)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality())
			})())), (_index152).Int())
		} else if _source17.Is_DC15() {
			var _1105___mcc_h4 _dafny.Array = _source17.Get_().(D8_DC15).Cf27
			_ = _1105___mcc_h4
			var _1106_cf27 _dafny.Array = _1105___mcc_h4
			_ = _1106_cf27
			var _1107_v8 bool
			_ = _1107_v8
			_1107_v8 = true
			_1107_v8 = (Companion_Default___.Fm41(_this.F5, _this.F5, globalState)).Dtor_cf30()
			var _1108_v9 _dafny.Map
			_ = _1108_v9
			_1108_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			var _1109_v10 _dafny.Map
			_ = _1109_v10
			_1109_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
			var _1110_v11 _dafny.Array
			_ = _1110_v11
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_27
			var _nw182 _dafny.Array
			_ = _nw182
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw182 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Map = (func(_1111_v9 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1112_i0 _dafny.Int) _dafny.Map {
						return _1111_v9
					}
				})(_1108_v9)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw182 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw182).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw182).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1110_v11 = _nw182
			var _1113_v12 *C2
			_ = _1113_v12
			var _nw183 *C2 = New_C2_()
			_ = _nw183
			_nw183.Ctor__((func() bool {
				if (_1108_v9).Contains((func() _dafny.Int {
					if ((_this).F4()).Contains(!(_1107_v8)) {
						return ((_this).F4()).Get(!(_1107_v8)).(_dafny.Int)
					}
					return p1
				})()) {
					return (_1108_v9).Get((func() _dafny.Int {
						if ((_this).F4()).Contains(!(_1107_v8)) {
							return ((_this).F4()).Get(!(_1107_v8)).(_dafny.Int)
						}
						return p1
					})()).(bool)
				}
				return p2
			})(), (_1109_v10).Cardinality(), _1110_v11)
			_1113_v12 = _nw183
			var _1114_v13 _dafny.CodePoint
			_ = _1114_v13
			_1114_v13 = _dafny.CodePoint('l')
			var _1115_v14 _dafny.Sequence
			_ = _1115_v14
			_1115_v14 = _dafny.SeqOf(_1114_v13)
			var _1116_v15 _dafny.Sequence
			_ = _1116_v15
			_1116_v15 = _dafny.SeqOf(_dafny.SeqOf(_1114_v13), _1115_v14)
			(globalState).F2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.Fm17(_1116_v15, (_dafny.Zero).Minus(p1), (_1108_v9).Cardinality(), p1, globalState)), p1))
			var _1117_v16 _dafny.Array
			_ = _1117_v16
			var _nwElement0_47 _dafny.Int = (_1113_v12).F12()
			_ = _nwElement0_47
			var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(5))
			_ = _nw184
			(_nw184).ArraySet1(_nwElement0_47, 0)
			(_nw184).ArraySet1(p3, 1)
			(_nw184).ArraySet1((func() _dafny.Int {
				if _1107_v8 {
					return p3
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_1115_v14).Cardinality()))
			})(), 2)
			(_nw184).ArraySet1((p3).Plus(p1), 3)
			(_nw184).ArraySet1(p3, 4)
			_1117_v16 = _nw184
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_28
			var _nw185 _dafny.Array
			_ = _nw185
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw185 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = func(_1118_i1 _dafny.Int) _dafny.Int {
					return (_1118_i1).Plus(_this.F5)
				}
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw185 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw185).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw185).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1117_v16 = _nw185
		} else {
			var _1119___mcc_h5 D8 = _source17.Get_().(D8_DC17).Cf32
			_ = _1119___mcc_h5
			var _1120_cf32 D8 = _1119___mcc_h5
			_ = _1120_cf32
			var _1121_v17 *C5
			_ = _1121_v17
			var _nw186 *C5 = New_C5_()
			_ = _nw186
			_nw186.Ctor__(Companion_Default___.Fm26(_this.F5, (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if ((_this).F4()).Contains(p2) {
					return ((_this).F4()).Get(p2).(_dafny.Int)
				}
				return p1
			})())), p2, globalState))
			_1121_v17 = _nw186
			var _1122_v18 _dafny.Map
			_ = _1122_v18
			_1122_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1121_v17).F6())
			var _1123_v19 _dafny.Sequence
			_ = _1123_v19
			_1123_v19 = _dafny.SeqOf(p1, _dafny.IntOfInt64(-934), (_1122_v18).Cardinality())
			var _1124_v20 _dafny.Set
			_ = _1124_v20
			_1124_v20 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F5, p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}(func(_1125_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(228)
			}))).Cardinality()), p1), _1123_v19))
			_1124_v20 = _1124_v20
			var _1126_v21 _dafny.Array
			_ = _1126_v21
			var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw187
			_1126_v21 = _nw187
			var _1127_v22 _dafny.Sequence
			_ = _1127_v22
			_1127_v22 = _dafny.SeqOf(_1126_v21)
			var _1128_v23 _dafny.Array
			_ = _1128_v23
			var _nwElement0_48 _dafny.Array = _1126_v21
			_ = _nwElement0_48
			var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(18))
			_ = _nw188
			(_nw188).ArraySet1(_nwElement0_48, 0)
			(_nw188).ArraySet1(_1126_v21, 1)
			(_nw188).ArraySet1(_1126_v21, 2)
			(_nw188).ArraySet1(_1126_v21, 3)
			(_nw188).ArraySet1(_1126_v21, 4)
			(_nw188).ArraySet1(_1126_v21, 5)
			(_nw188).ArraySet1(_1126_v21, 6)
			(_nw188).ArraySet1(_1126_v21, 7)
			(_nw188).ArraySet1((func() _dafny.Array {
				if p2 {
					return _1126_v21
				}
				return _1126_v21
			})(), 8)
			(_nw188).ArraySet1((_1127_v22).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1127_v22).Cardinality()))).Uint32()).(_dafny.Array), 9)
			(_nw188).ArraySet1(_1126_v21, 10)
			(_nw188).ArraySet1(_1126_v21, 11)
			(_nw188).ArraySet1(_1126_v21, 12)
			(_nw188).ArraySet1((_1127_v22).Select((Companion_Default___.SafeIndex(_this.F5, _dafny.IntOfUint32((_1127_v22).Cardinality()))).Uint32()).(_dafny.Array), 13)
			(_nw188).ArraySet1(_1126_v21, 14)
			(_nw188).ArraySet1((func() _dafny.Array {
				if true {
					return _1126_v21
				}
				return _1126_v21
			})(), 15)
			(_nw188).ArraySet1(_1126_v21, 16)
			(_nw188).ArraySet1(_1126_v21, 17)
			_1128_v23 = _nw188
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_1128_v23), 0))
			_ = _index153
			(_1128_v23).ArraySet1(_1126_v21, (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_1128_v23), 0))
			_ = _index154
			(_1128_v23).ArraySet1(_1126_v21, (_index154).Int())
			var _1129_v24 _dafny.Map
			_ = _1129_v24
			_1129_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			if ((p2) || (p2)) == ((func() bool {
				if (_1129_v24).Contains(_this.F5) {
					return (_1129_v24).Get(_this.F5).(bool)
				}
				return p2
			})()) {
				var _1130_v25 bool
				_ = _1130_v25
				_1130_v25 = false
				_1130_v25 = p2
				var _1131_v26 *C0
				_ = _1131_v26
				var _nw189 *C0 = New_C0_()
				_ = _nw189
				_nw189.Ctor__((_this.F5).Minus(_this.F5))
				_1131_v26 = _nw189
				var _1132_v27 _dafny.Sequence
				_ = _1132_v27
				_1132_v27 = _dafny.SeqOf(Companion_Default___.Fm25(false, _dafny.IntOfUint32((_1123_v19).Cardinality()), globalState))
				var _1133_v28 _dafny.Set
				_ = _1133_v28
				_1133_v28 = _dafny.SetOf(p1)
				var _rhs145 _dafny.Int = Companion_Default___.Fm17(_1132_v27, (_1131_v26).F13(), Companion_Default___.Fm9((_1133_v28).Cardinality(), globalState), (_dafny.Zero).Minus(p1), globalState)
				_ = _rhs145
				var _rhs146 *C0 = _1131_v26
				_ = _rhs146
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				_lhs87.F2 = _rhs145
				_1131_v26 = _rhs146
				_1130_v25 = _1130_v25
				var _1134_v29 _dafny.Map
				_ = _1134_v29
				_1134_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1130_v25, _1130_v25)
				var _1135_v30 _dafny.Set
				_ = _1135_v30
				_1135_v30 = _dafny.SetOf((func() bool {
					if (_1134_v29).Contains(_1130_v25) {
						return (_1134_v29).Get(_1130_v25).(bool)
					}
					return _1130_v25
				})())
				_1130_v25 = !(_1135_v30).Contains(p2)
				var _1136_v31 _dafny.Sequence
				_ = _1136_v31
				_1136_v31 = _dafny.UnicodeSeqOfUtf8Bytes("myiqun")
				var _1137_v32 D6
				_ = _1137_v32
				_1137_v32 = Companion_D6_.Create_DC11_(_1136_v31)
				var _1138_v33 _dafny.Sequence
				_ = _1138_v33
				_1138_v33 = _dafny.SeqOf(_1137_v32)
				var _1139_v34 _dafny.Array
				_ = _1139_v34
				var _nw190 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
				_ = _nw190
				_1139_v34 = _nw190
				var _1140_v35 _dafny.Array
				_ = _1140_v35
				var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw191
				_1140_v35 = _nw191
				var _1141_v36 _dafny.MultiSet
				_ = _1141_v36
				_1141_v36 = _dafny.MultiSetOf(_1140_v35, _1140_v35)
				var _1142_v37 _dafny.Map
				_ = _1142_v37
				_1142_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1138_v33, _1139_v34), _1141_v36)
				var _1143_v38 _dafny.Map
				_ = _1143_v38
				_1143_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1137_v32), _1139_v34)
				_1142_v37 = (_1142_v37).Update(_1143_v38, _1141_v36)
			} else {
				var _1144_v39 _dafny.Array
				_ = _1144_v39
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_29
				var _nw192 _dafny.Array
				_ = _nw192
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw192 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = (func(_1145_p1 _dafny.Int, _1146_p2 bool) func(_dafny.Int) _dafny.Int {
						return func(_1147_i3 _dafny.Int) _dafny.Int {
							return (_1147_i3).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1145_p1, _dafny.SeqOf(_1146_p2, _1146_p2))).Cardinality())
						}
					})(p1, p2)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw192 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw192).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw192).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1144_v39 = _nw192
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))
				_ = _index155
				(_1144_v39).ArraySet1(_this.F5, (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))
				_ = _index156
				(_1144_v39).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2, p2), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality()))).Uint32(), p2)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1089_v0)).Cardinality())), (_index156).Int())
				var _1148_v40 _dafny.Int
				_ = _1148_v40
				var _1149_v41 bool
				_ = _1149_v41
				var _out41 _dafny.Int
				_ = _out41
				var _out42 bool
				_ = _out42
				_out41, _out42 = (_this).M2(p2, globalState)
				_1148_v40 = _out41
				_1149_v41 = _out42
				var _1150_v42 _dafny.Sequence
				_ = _1150_v42
				_1150_v42 = _dafny.UnicodeSeqOfUtf8Bytes("yxohwdduo")
				var _1151_v43 _dafny.Map
				_ = _1151_v43
				_1151_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1121_v17).F6(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1150_v42).Cardinality())))
				_1151_v43 = (_1151_v43).Update((_1121_v17).F6(), _dafny.IntOfInt64(435))
				_1149_v41 = Companion_Default___.Fm21((_1144_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))).Int()).(_dafny.Int), (_1121_v17).F6(), _1149_v41, (func() bool {
					if _1149_v41 {
						return false
					}
					return _1149_v41
				})(), globalState)
				var _1152_v44 _dafny.Map
				_ = _1152_v44
				_1152_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
				_1152_v44 = (_1152_v44).Update(((_1144_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))).Int()).(_dafny.Int)).Cmp((_1144_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))).Int()).(_dafny.Int)) != 0, (_1148_v40).Plus((_1144_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1144_v39), 0))).Int()).(_dafny.Int)))
			}
		}
		if p2 {
			var _1153_v45 _dafny.Map
			_ = _1153_v45
			_1153_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3).Times(p3), _1089_v0)
			var _1154_v47 _dafny.MultiSet
			_ = _1154_v47
			_1154_v47 = _dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.Fm24(func() _dafny.Set {
				var _coll36 = _dafny.NewBuilder()
				_ = _coll36
				for _iter37 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-165), _dafny.IntOfInt64(396))).Elements()); ; {
					_compr_36, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1155_v46 _dafny.Int
					_1155_v46 = interface{}(_compr_36).(_dafny.Int)
					if (_dafny.MultiSetOf(_dafny.IntOfInt64(-165), _dafny.IntOfInt64(396))).Contains(_1155_v46) {
						_coll36.Add((_1155_v46).Times(_dafny.IntOfInt64(-395)))
					}
				}
				return _coll36.ToSet()
			}(), globalState)), p1)
			var _1156_v48 _dafny.CodePoint
			_ = _1156_v48
			_1156_v48 = _dafny.CodePoint('d')
			var _1157_v49 _dafny.Sequence
			_ = _1157_v49
			_1157_v49 = _dafny.SeqOf(p3)
			var _1158_v50 _dafny.Sequence
			_ = _1158_v50
			_1158_v50 = _dafny.SeqOf(_1157_v49)
			var _1159_v51 _dafny.Sequence
			_ = _1159_v51
			_1159_v51 = _dafny.SeqOf(_dafny.IntOfUint32((_1158_v50).Cardinality()), p3)
			var _1160_v52 D4
			_ = _1160_v52
			_1160_v52 = Companion_D4_.Create_DC9_((_1154_v47).Cardinality(), _1156_v48, _1157_v49, _this.F5, p2)
			_1153_v45 = (_1153_v45).Update((_1160_v52).Dtor_cf19(), _1089_v0)
			var _1161_v53 bool
			_ = _1161_v53
			_1161_v53 = false
			var _1162_v54 D8
			_ = _1162_v54
			_1162_v54 = Companion_D8_.Create_DC16_(_dafny.IntOfUint32((_1159_v51).Cardinality()), _1161_v53, false, _this.F5)
			var _rhs147 bool = (_1162_v54).Dtor_cf30()
			_ = _rhs147
			var _rhs148 bool = p2
			_ = _rhs148
			var _rhs149 _dafny.Int = p3
			_ = _rhs149
			var _rhs150 _dafny.Int = p1
			_ = _rhs150
			var _lhs88 *C6 = _this
			_ = _lhs88
			_1161_v53 = _rhs147
			_1161_v53 = _rhs148
			r1 = _rhs149
			_lhs88.F5 = _rhs150
			var _1163_v55 _dafny.Sequence
			_ = _1163_v55
			_1163_v55 = _dafny.SeqOf(_1161_v53)
			var _1164_v56 _dafny.Map
			_ = _1164_v56
			_1164_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1163_v55)
			var _1165_v57 _dafny.Sequence
			_ = _1165_v57
			_1165_v57 = _dafny.UnicodeSeqOfUtf8Bytes("isekkdev")
			var _1166_v58 _dafny.MultiSet
			_ = _1166_v58
			_1166_v58 = _dafny.MultiSetOf(!(_1161_v53))
			var _1167_v59 _dafny.Sequence
			_ = _1167_v59
			_1167_v59 = _dafny.SeqOf(_1165_v57, _1165_v57, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1165_v57, (Companion_Default___.SafeIndex(_this.F5, _dafny.IntOfUint32((_1165_v57).Cardinality()))).Uint32(), _1156_v48), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1166_v58).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1165_v57, (Companion_Default___.SafeIndex(_this.F5, _dafny.IntOfUint32((_1165_v57).Cardinality()))).Uint32(), _1156_v48)).Cardinality()))).Uint32(), _1156_v48))
			var _1168_v60 _dafny.Map
			_ = _1168_v60
			_1168_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1154_v47).Cardinality(), _1161_v53)
			var _1169_v61 _dafny.Sequence
			_ = _1169_v61
			_1169_v61 = _dafny.SeqOf(_1168_v60, _1168_v60)
			var _1170_v63 D6
			_ = _1170_v63
			_1170_v63 = Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("wwxjc"))
			var _1171_v65 _dafny.Array
			_ = _1171_v65
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
			_ = _nw193
			_1171_v65 = _nw193
			var _1172_v66 T0
			_ = _1172_v66
			var _nw194 *C4 = New_C4_()
			_ = _nw194
			_nw194.Ctor__(_1161_v53, _1154_v47, _1171_v65)
			_1172_v66 = _nw194
			var _1173_v67 D9
			_ = _1173_v67
			_1173_v67 = Companion_D9_.Create_DC19_(_1172_v66, _dafny.IntOfInt64(80), (_dafny.Zero).Minus(_this.F5))
			var _1174_v68 _dafny.Map
			_ = _1174_v68
			_1174_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1157_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.IntOfUint32((_1157_v49).Cardinality()))).Uint32()).(_dafny.Int), _1156_v48)
			var _1175_v69 _dafny.Array
			_ = _1175_v69
			var _nwElement0_49 _dafny.Int = p1
			_ = _nwElement0_49
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(27))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_49, 0)
			(_nw195).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1163_v55)).Merge(_1164_v56)).Cardinality(), 1)
			(_nw195).ArraySet1((p3).Plus(_this.F5), 2)
			(_nw195).ArraySet1(p3, 3)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_1167_v59).Cardinality()), 4)
			(_nw195).ArraySet1(p1, 5)
			(_nw195).ArraySet1(Companion_Default___.Fm24(func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter38 := _dafny.Iterate(((_1169_v61).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1169_v61).Cardinality()))).Uint32()).(_dafny.Map)).Keys().Elements()); ; {
					_compr_37, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _1176_v62 _dafny.Int
					_1176_v62 = interface{}(_compr_37).(_dafny.Int)
					if ((_1169_v61).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1169_v61).Cardinality()))).Uint32()).(_dafny.Map)).Contains(_1176_v62) {
						_coll37.Add((_1176_v62).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(559), _dafny.IntOfInt64(505))).Cardinality()))))
					}
				}
				return _coll37.ToSet()
			}(), globalState), 6)
			(_nw195).ArraySet1((func() _dafny.Int {
				if (_1154_v47).Contains(p1) {
					return (_1154_v47).Multiplicity(p1)
				}
				return (Companion_D2_.Create_DC4_(p3)).Dtor_cf6()
			})(), 7)
			(_nw195).ArraySet1((_dafny.IntOfUint32((_1165_v57).Cardinality())).Minus((Companion_D7_.Create_DC14_(p1)).Dtor_cf26()), 8)
			(_nw195).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_1170_v63).Dtor_cf22()).Cardinality()), (_1159_v51).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1159_v51).Cardinality()))).Uint32()).(_dafny.Int)), 9)
			(_nw195).ArraySet1(Companion_Default___.Fm24(func() _dafny.Set {
				var _coll38 = _dafny.NewBuilder()
				_ = _coll38
				for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(190), _dafny.IntOfInt64(666))); ; {
					_compr_38, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _1177_v64 _dafny.Int
					_1177_v64 = interface{}(_compr_38).(_dafny.Int)
					if ((_dafny.IntOfInt64(190)).Cmp(_1177_v64) <= 0) && ((_1177_v64).Cmp(_dafny.IntOfInt64(666)) < 0) {
						_coll38.Add((_1177_v64).Minus((_1166_v58).Cardinality()))
					}
				}
				return _coll38.ToSet()
			}(), globalState), 10)
			(_nw195).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1159_v51).Cardinality()), p1), 11)
			(_nw195).ArraySet1(p1, 12)
			(_nw195).ArraySet1(p3, 13)
			(_nw195).ArraySet1(_this.F5, 14)
			(_nw195).ArraySet1(_this.F5, 15)
			(_nw195).ArraySet1((p1).Minus(_this.F5), 16)
			(_nw195).ArraySet1(p3, 17)
			(_nw195).ArraySet1(_this.F5, 18)
			(_nw195).ArraySet1((_1173_v67).Dtor_cf35(), 19)
			(_nw195).ArraySet1((p1).Minus(_this.F5), 20)
			(_nw195).ArraySet1(_this.F5, 21)
			(_nw195).ArraySet1(_this.F5, 22)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(225), _dafny.CodePoint('l')), _1174_v68, _1174_v68, _1174_v68)).Cardinality()), 23)
			(_nw195).ArraySet1(p1, 24)
			(_nw195).ArraySet1(p3, 25)
			(_nw195).ArraySet1(p3, 26)
			_1175_v69 = _nw195
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_1175_v69), 0))
			_ = _index157
			(_1175_v69).ArraySet1(p3, (_index157).Int())
			var _1178_v70 _dafny.Array
			_ = _1178_v70
			var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
			_ = _nw196
			_1178_v70 = _nw196
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1089_v0), 0))
			_ = _index158
			(_1089_v0).ArraySet1((func() bool {
				if p2 {
					return p2
				}
				return _1161_v53
			})(), (_index158).Int())
			var _1179_v72 _dafny.Map
			_ = _1179_v72
			_1179_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1165_v57, _1165_v57)
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_1175_v69), 0))
			_ = _index159
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1089_v0), 0))
			_ = _index160
			var _rhs151 _dafny.Int = Companion_Default___.Fm17(_1167_v59, (func() _dafny.Map {
				var _coll39 = _dafny.NewMapBuilder()
				_ = _coll39
				for _iter40 := _dafny.Iterate((_1179_v72).Keys().Elements()); ; {
					_compr_39, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _1180_v71 _dafny.Sequence
					_1180_v71 = interface{}(_compr_39).(_dafny.Sequence)
					if (_1179_v72).Contains(_1180_v71) {
						_coll39.Add(_1180_v71, _1165_v57)
					}
				}
				return _coll39.ToMap()
			}()).Cardinality(), p1, _this.F5, globalState)
			_ = _rhs151
			var _rhs152 _dafny.Array = _1178_v70
			_ = _rhs152
			var _rhs153 bool = Companion_Default___.Fm21((func() _dafny.Int {
				if ((_this).F4()).Contains((_1162_v54).Dtor_cf30()) {
					return ((_this).F4()).Get((_1162_v54).Dtor_cf30()).(_dafny.Int)
				}
				return _this.F5
			})(), _1156_v48, p2, p2, globalState)
			_ = _rhs153
			var _lhs89 _dafny.Array = _1175_v69
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_1175_v69), 0))
			_ = _lhs90
			var _lhs91 _dafny.Array = _1089_v0
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1089_v0), 0))
			_ = _lhs92
			(_lhs89).ArraySet1(_rhs151, (_lhs90).Int())
			_1178_v70 = _rhs152
			(_lhs91).ArraySet1(_rhs153, (_lhs92).Int())
			var _1181_v73 _dafny.MultiSet
			_ = _1181_v73
			_1181_v73 = _dafny.MultiSetOf(_1163_v55)
			var _1182_v74 _dafny.Map
			_ = _1182_v74
			_1182_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1161_v53, p2)
			var _1183_v75 _dafny.Array
			_ = _1183_v75
			_1183_v75 = _1175_v69
			var _1184_v76 _dafny.Map
			_ = _1184_v76
			_1184_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm8((func() _dafny.Int {
				if (_1181_v73).Contains(_1163_v55) {
					return (_1181_v73).Multiplicity(_1163_v55)
				}
				return (_1182_v74).Cardinality()
			})(), globalState), _1183_v75)
			_1184_v76 = (_1184_v76).Update(_1165_v57, _1183_v75)
			var _1185_v77 _dafny.Map
			_ = _1185_v77
			_1185_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _this.F5)
			(globalState).F2 = Companion_Default___.SafeDivisionInt(p3, (func() _dafny.Int {
				if p2 {
					return (_1185_v77).Cardinality()
				}
				return _dafny.IntOfUint32((_1165_v57).Cardinality())
			})())
		} else {
			var _1186_v78 bool
			_ = _1186_v78
			_1186_v78 = true
			_1186_v78 = (true) || (p2)
			var _1187_v79 _dafny.Set
			_ = _1187_v79
			_1187_v79 = _dafny.SetOf(_this.F5)
			if !((Companion_Default___.Fm38(_dafny.IntOfInt64(773), true, !(p2), _1186_v78, globalState)).Union(_1187_v79)).Contains(p3) {
				var _1188_v80 _dafny.CodePoint
				_ = _1188_v80
				_1188_v80 = _dafny.CodePoint('v')
				_1188_v80 = Companion_Default___.Fm26((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v78, _1186_v78)).Cardinality(), (_1187_v79).Cardinality(), p2, globalState)
				var _1189_v81 *C5
				_ = _1189_v81
				var _nw197 *C5 = New_C5_()
				_ = _nw197
				_nw197.Ctor__(_1188_v80)
				_1189_v81 = _nw197
				_1188_v80 = (_1189_v81).F6()
				r0 = p1
				var _1190_v82 _dafny.Sequence
				_ = _1190_v82
				_1190_v82 = _dafny.UnicodeSeqOfUtf8Bytes("cvmfvuww")
				var _1191_v83 _dafny.Sequence
				_ = _1191_v83
				_1191_v83 = _dafny.SeqOf(p2, _1186_v78, _1186_v78)
				var _1192_v84 _dafny.Map
				_ = _1192_v84
				_1192_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1190_v82, _1191_v83)
				var _1193_v85 _dafny.Map
				_ = _1193_v85
				_1193_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.UnicodeSeqOfUtf8Bytes("oaxbduxb"))
				_1192_v84 = (_1192_v84).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_1193_v85).Contains(_dafny.IntOfInt64(229)) {
						return (_1193_v85).Get(_dafny.IntOfInt64(229)).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-136))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg57 _dafny.Int) interface{} {
							return coer57(arg57)
						}
					}((func(_1194_v81 *C5) func(_dafny.Int) _dafny.CodePoint {
						return func(_1195_i4 _dafny.Int) _dafny.CodePoint {
							return (_1194_v81).F6()
						}
					})(_1189_v81)))
				})(), (Companion_Default___.SafeIndex(_this.F5, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1193_v85).Contains(_dafny.IntOfInt64(229)) {
						return (_1193_v85).Get(_dafny.IntOfInt64(229)).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-136))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg58 _dafny.Int) interface{} {
							return coer58(arg58)
						}
					}((func(_1196_v81 *C5) func(_dafny.Int) _dafny.CodePoint {
						return func(_1197_i4 _dafny.Int) _dafny.CodePoint {
							return (_1196_v81).F6()
						}
					})(_1189_v81)))
				})()).Cardinality()))).Uint32(), _1188_v80), _1190_v82), _1191_v83)
			} else {
				var _1198_v86 _dafny.MultiSet
				_ = _1198_v86
				_1198_v86 = _dafny.MultiSetOf(p3)
				var _1199_v87 *C3
				_ = _1199_v87
				var _nw198 *C3 = New_C3_()
				_ = _nw198
				_nw198.Ctor__((_1198_v86).Difference(_1198_v86), !(false))
				_1199_v87 = _nw198
				var _1200_v88 _dafny.Sequence
				_ = _1200_v88
				_1200_v88 = _dafny.UnicodeSeqOfUtf8Bytes("tplrn")
				_1186_v78 = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("jjpsq"), _dafny.Companion_Sequence_.Concatenate(_1200_v88, _dafny.UnicodeSeqOfUtf8Bytes("diviu"))))
				var _1201_v89 _dafny.Map
				_ = _1201_v89
				_1201_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v78, _1186_v78)
				_1201_v89 = (_1201_v89).Update(true, p2)
				var _1202_v90 _dafny.Array
				_ = _1202_v90
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_30
				var _nw199 _dafny.Array
				_ = _nw199
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw199 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Int = (func(_1203_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1204_i5 _dafny.Int) _dafny.Int {
							return (_1204_i5).Times(_1203_p3)
						}
					})(p3)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw199 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw199).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw199).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1202_v90 = _nw199
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1202_v90), 0))
				_ = _index161
				(_1202_v90).ArraySet1(p1, (_index161).Int())
				var _1205_v91 _dafny.CodePoint
				_ = _1205_v91
				_1205_v91 = _dafny.CodePoint('h')
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1202_v90), 0))
				_ = _index162
				(_1202_v90).ArraySet1((_1199_v87).Fm19(_1205_v91, _dafny.IntOfInt64(272), globalState), (_index162).Int())
				(globalState).F2 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(74), (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1198_v86).Contains(_this.F5) {
						return (_1198_v86).Multiplicity(_this.F5)
					}
					return _this.F5
				})())), p1)
			}
			(_this).F5 = Companion_Default___.Fm24(_1187_v79, globalState)
			var _1206_v92 _dafny.Int
			_ = _1206_v92
			var _1207_v93 bool
			_ = _1207_v93
			var _out43 _dafny.Int
			_ = _out43
			var _out44 bool
			_ = _out44
			_out43, _out44 = (_this).M2(p2, globalState)
			_1206_v92 = _out43
			_1207_v93 = _out44
			var _1208_v94 _dafny.Map
			_ = _1208_v94
			_1208_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-249), _1207_v93)
			_1208_v94 = (_1208_v94).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vpbxb")).Cardinality()), p2)
		}
		(_this).F5 = (p1).Times(p3)
		var _1209_i6 _dafny.Int
		_ = _1209_i6
		_1209_i6 = _dafny.Zero
		{
			for p2 {
				{
					if (_1209_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1209_i6 = (_1209_i6).Plus(_dafny.One)
					var _1210_v95 _dafny.MultiSet
					_ = _1210_v95
					_1210_v95 = _dafny.MultiSetOf(p3)
					var _1211_v96 _dafny.Sequence
					_ = _1211_v96
					_1211_v96 = _dafny.SeqOf(_1210_v95, _1210_v95, _1210_v95)
					(globalState).F2 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((_1211_v96).Cardinality()))), _this.F5)
					var _1212_v97 _dafny.Array
					_ = _1212_v97
					var _nw200 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
					_ = _nw200
					_1212_v97 = _nw200
					_1212_v97 = _1212_v97
					var _1213_v98 bool
					_ = _1213_v98
					_1213_v98 = false
					var _1214_v99 _dafny.Sequence
					_ = _1214_v99
					_1214_v99 = _dafny.UnicodeSeqOfUtf8Bytes("vm")
					var _1215_v100 _dafny.Map
					_ = _1215_v100
					_1215_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F5, _dafny.Companion_Sequence_.IsProperPrefixOf(_1214_v99, _1214_v99))
					var _1216_v101 _dafny.Set
					_ = _1216_v101
					_1216_v101 = _dafny.SetOf(_1089_v0)
					var _rhs154 bool = !((func() bool {
						if (_1215_v100).Contains(((_1216_v101).Cardinality()).Minus(_this.F5)) {
							return (_1215_v100).Get(((_1216_v101).Cardinality()).Minus(_this.F5)).(bool)
						}
						return (func() bool {
							if p2 {
								return p2
							}
							return false
						})()
					})())
					_ = _rhs154
					var _rhs155 bool = !(_1213_v98)
					_ = _rhs155
					_1213_v98 = _rhs154
					_1213_v98 = _rhs155
					(globalState).F2 = p3
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1217_v102 _dafny.Sequence
		_ = _1217_v102
		_1217_v102 = _dafny.UnicodeSeqOfUtf8Bytes("tp")
		var _1218_v103 _dafny.Map
		_ = _1218_v103
		_1218_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.Companion_Sequence_.Concatenate(_1217_v102, _1217_v102))
		var _1219_v104 _dafny.CodePoint
		_ = _1219_v104
		_1219_v104 = _dafny.CodePoint('b')
		_1218_v103 = (_1218_v103).Update(_1219_v104, _dafny.UnicodeSeqOfUtf8Bytes("eltjfofh"))
		var _1220_v105 _dafny.Set
		_ = _1220_v105
		_1220_v105 = _dafny.SetOf(false)
		var _1221_v106 _dafny.Map
		_ = _1221_v106
		_1221_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
		var _1222_v107 _dafny.Map
		_ = _1222_v107
		_1222_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1221_v106).Cardinality())
		var _1223_v108 _dafny.Map
		_ = _1223_v108
		_1223_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p2, p2, p2)).IsProperSubsetOf(_1220_v105), ((_1222_v107).Update(p1, _this.F5)).Cardinality())
		_1223_v108 = (_1223_v108).Update(p2, _this.F5)
		r0 = _dafny.IntOfInt64(883)
		r1 = (_dafny.Zero).Minus(p1)
		var _1224_v109 D2
		_ = _1224_v109
		_1224_v109 = Companion_D2_.Create_DC5_(p3, p2, false, p2, p2)
		r2 = (func() _dafny.Set {
			if !((_1224_v109).Dtor_cf10()) {
				return func() _dafny.Set {
					var _coll40 = _dafny.NewBuilder()
					_ = _coll40
					for _iter41 := _dafny.Iterate((Companion_Default___.Fm42(false, p3, p2, globalState)).Keys().Elements()); ; {
						_compr_40, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _1225_v110 _dafny.Sequence
						_1225_v110 = interface{}(_compr_40).(_dafny.Sequence)
						if (Companion_Default___.Fm42(false, p3, p2, globalState)).Contains(_1225_v110) {
							_coll40.Add(_1225_v110)
						}
					}
					return _coll40.ToSet()
				}()
			}
			return Companion_Default___.Fm35(globalState)
		})()
		return r0, r1, r2
	}
}
func (_this *C6) F4() _dafny.Map {
	{
		return _this._f4
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f3 _dafny.Array
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C7{}
var _ T1 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F3() _dafny.Array {
	return _this._f3
}
func (_this *C7) Ctor__(f3 _dafny.Array) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C7) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(297))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fuwqkmw")).Cardinality()))).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}(func(_1226_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()))
	}
}
func (_this *C7) Fm1(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(133)
	}
}
func (_this *C7) M1(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var _1227_v0 _dafny.Int
		_ = _1227_v0
		_1227_v0 = _dafny.IntOfInt64(-681)
		r0 = (func() _dafny.Int {
			if p1 {
				return _1227_v0
			}
			return _1227_v0
		})()
		if p1 {
			var _1228_v1 _dafny.Sequence
			_ = _1228_v1
			_1228_v1 = _dafny.SeqOf(p1)
			var _1229_v2 _dafny.Map
			_ = _1229_v2
			_1229_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1228_v1, (Companion_Default___.SafeIndex(_1227_v0, _dafny.IntOfUint32((_1228_v1).Cardinality()))).Uint32(), p1), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_1227_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _dafny.IntOfUint32((p2).Cardinality())))
			var _1230_v3 _dafny.CodePoint
			_ = _1230_v3
			_1230_v3 = _dafny.CodePoint('t')
			var _1231_v4 _dafny.Map
			_ = _1231_v4
			_1231_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1228_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.IntOfUint32((_1228_v1).Cardinality()))).Uint32()).(bool))
			var _1232_v5 _dafny.Sequence
			_ = _1232_v5
			_1232_v5 = _dafny.SeqOf(_1231_v4, _1231_v4)
			var _1233_v6 D1
			_ = _1233_v6
			_1233_v6 = Companion_D1_.Create_DC1_(_1232_v5)
			var _1234_v7 _dafny.MultiSet
			_ = _1234_v7
			_1234_v7 = _dafny.MultiSetOf(p1)
			var _1235_v8 _dafny.Array
			_ = _1235_v8
			var _nwElement0_50 bool = (_1229_v2).Contains(_1228_v1)
			_ = _nwElement0_50
			var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(29))
			_ = _nw201
			(_nw201).ArraySet1(_nwElement0_50, 0)
			(_nw201).ArraySet1(p1, 1)
			(_nw201).ArraySet1(p1, 2)
			(_nw201).ArraySet1((p1) || (false), 3)
			(_nw201).ArraySet1(!(p1) || (p1), 4)
			(_nw201).ArraySet1(p1, 5)
			(_nw201).ArraySet1(p1, 6)
			(_nw201).ArraySet1((_1227_v0).Cmp(_1227_v0) < 0, 7)
			(_nw201).ArraySet1(p1, 8)
			(_nw201).ArraySet1(p1, 9)
			(_nw201).ArraySet1(!(p1) || (p1), 10)
			(_nw201).ArraySet1(p1, 11)
			(_nw201).ArraySet1((_1227_v0).Cmp(_1227_v0) > 0, 12)
			(_nw201).ArraySet1(!(Companion_Default___.Fm2(_1230_v3, (_1233_v6).Dtor_cf1(), p1, globalState)), 13)
			(_nw201).ArraySet1(p1, 14)
			(_nw201).ArraySet1((_dafny.IntOfInt64(-190)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v0, true)).Cardinality()) <= 0, 15)
			(_nw201).ArraySet1(p1, 16)
			(_nw201).ArraySet1(Companion_Default___.Fm2(_1230_v3, Companion_Default___.Fm3(globalState), false, globalState), 17)
			(_nw201).ArraySet1((_1227_v0).Cmp(_dafny.IntOfInt64(-359)) <= 0, 18)
			(_nw201).ArraySet1(p1, 19)
			(_nw201).ArraySet1(p1, 20)
			(_nw201).ArraySet1(p1, 21)
			(_nw201).ArraySet1(p1, 22)
			(_nw201).ArraySet1(p1, 23)
			(_nw201).ArraySet1(p1, 24)
			(_nw201).ArraySet1(!((_1234_v7).Contains(true)), 25)
			(_nw201).ArraySet1(p1, 26)
			(_nw201).ArraySet1(p1, 27)
			(_nw201).ArraySet1(p1, 28)
			_1235_v8 = _nw201
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
			_ = _index163
			(_1235_v8).ArraySet1(p1, (_index163).Int())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
			_ = _index164
			(_1235_v8).ArraySet1(p1, (_index164).Int())
			var _1236_v9 _dafny.Array
			_ = _1236_v9
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_31
			var _nw202 _dafny.Array
			_ = _nw202
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw202 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Map = (func(_1237_v0 _dafny.Int, _1238_v1 _dafny.Sequence, _1239_v8 _dafny.Array, _1240_p1 bool) func(_dafny.Int) _dafny.Map {
					return func(_1241_i0 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-156), false, _1237_v0), _1238_v1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1239_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1239_v8), 0))).Int()).(bool), _1237_v0)).Cardinality(), _1240_p1, _1237_v0), _1238_v1))
					}
				})(_1227_v0, _1228_v1, _1235_v8, p1)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw202 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw202).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw202).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1236_v9 = _nw202
			var _1242_v10 _dafny.MultiSet
			_ = _1242_v10
			_1242_v10 = _dafny.MultiSetOf(_1227_v0)
			var _1243_v11 D1
			_ = _1243_v11
			_1243_v11 = Companion_D1_.Create_DC2_((_1242_v10).Cardinality(), (_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), _1227_v0)
			var _1244_v12 _dafny.Map
			_ = _1244_v12
			_1244_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v11, _1228_v1)
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1236_v9), 0))
			_ = _index165
			(_1236_v9).ArraySet1((_1244_v12).Update(_1243_v11, _1228_v1), (_index165).Int())
			var _1245_v13 _dafny.Set
			_ = _1245_v13
			_1245_v13 = _dafny.SetOf(p1)
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1236_v9), 0))
			_ = _index166
			(_1236_v9).ArraySet1(((_1244_v12).Update(Companion_D1_.Create_DC2_(_1227_v0, (_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), (_1245_v13).Cardinality()), _1228_v1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v11, _1228_v1)), (_index166).Int())
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
			_ = _index167
			(_1235_v8).ArraySet1((_1243_v11).Dtor_cf3(), (_index167).Int())
			_1242_v10 = _1242_v10
			if Companion_Default___.Fm2(_1230_v3, _1232_v5, p1, globalState) {
				_1233_v6 = Companion_Default___.Fm4(p1, _1230_v3, _1227_v0, Companion_Default___.Fm5(_1227_v0, globalState), globalState)
				var _1246_v14 _dafny.Array
				_ = _1246_v14
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_32
				var _nw203 _dafny.Array
				_ = _nw203
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw203 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1247_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1248_i1 _dafny.Int) _dafny.Sequence {
							return _1247_p2
						}
					})(p2)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw203 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw203).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw203).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1246_v14 = _nw203
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1246_v14), 0))
				_ = _index168
				(_1246_v14).ArraySet1(p2, (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1246_v14), 0))
				_ = _index169
				(_1246_v14).ArraySet1(p2, (_index169).Int())
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
				_ = _index170
				(_1235_v8).ArraySet1((_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), (_index170).Int())
				var _1249_v15 _dafny.Map
				_ = _1249_v15
				_1249_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.CodePoint('n'))
				_1249_v15 = (_1249_v15).Update(p2, Companion_Default___.Fm5(_dafny.IntOfUint32((_1228_v1).Cardinality()), globalState))
				r0 = _1227_v0
			} else {
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
				_ = _index171
				(_1235_v8).ArraySet1((_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), (_index171).Int())
				var _1250_v16 _dafny.Array
				_ = _1250_v16
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_33
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) D1 = (func(_1251_v11 D1) func(_dafny.Int) D1 {
						return func(_1252_i2 _dafny.Int) D1 {
							return _1251_v11
						}
					})(_1243_v11)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw204 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw204).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw204).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1250_v16 = _nw204
				var _1253_v17 _dafny.Sequence
				_ = _1253_v17
				_1253_v17 = _dafny.SeqOf(_1235_v8, _1235_v8, _1235_v8, _1235_v8)
				var _1254_v18 _dafny.Map
				_ = _1254_v18
				_1254_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v16, _dafny.Companion_Sequence_.Concatenate(_1253_v17, _1253_v17))
				_1254_v18 = (_1254_v18).Update(_1250_v16, _dafny.SeqOf(_1235_v8, _1235_v8))
				r0 = _dafny.IntOfInt64(597)
				var _1255_v19 _dafny.Sequence
				_ = _1255_v19
				_1255_v19 = _dafny.SeqOf(_1228_v1)
				var _1256_v20 _dafny.Sequence
				_ = _1256_v20
				_1256_v20 = _1228_v1
				var _1257_v21 _dafny.Map
				_ = _1257_v21
				_1257_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1243_v11).Dtor_cf3(), _dafny.Companion_Sequence_.Concatenate((_1255_v19).Select((Companion_Default___.SafeIndex(_1227_v0, _dafny.IntOfUint32((_1255_v19).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm6(_dafny.IntOfUint32((p2).Cardinality()), (_this).Fm1((_1228_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-124), _dafny.IntOfUint32((_1228_v1).Cardinality()))).Uint32()).(bool), _1230_v3, _1256_v20, (_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), globalState), _1230_v3, globalState)))
				_1257_v21 = (_1257_v21).Update((_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool), _1228_v1)
				var _1258_v22 _dafny.Array
				_ = _1258_v22
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_34
				var _nw205 _dafny.Array
				_ = _nw205
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw205 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Int = (func(_1259_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1260_i3 _dafny.Int) _dafny.Int {
							return (_1260_i3).Plus(_dafny.IntOfUint32((_1259_p2).Cardinality()))
						}
					})(p2)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw205 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw205).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw205).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1258_v22 = _nw205
				var _1261_v23 _dafny.Sequence
				_ = _1261_v23
				_1261_v23 = _dafny.UnicodeSeqOfUtf8Bytes("bgdm")
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
				_ = _index172
				var _rhs156 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("khi"), _1261_v23)).Cardinality())
				_ = _rhs156
				var _rhs157 _dafny.Array = _1258_v22
				_ = _rhs157
				var _rhs158 bool = !((Companion_Default___.Fm2(_1230_v3, _1232_v5, !(!(true)), globalState)) && ((_1235_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))).Int()).(bool))) || ((_1227_v0).Cmp(_1227_v0) == 0)
				_ = _rhs158
				var _rhs159 _dafny.Sequence = p2
				_ = _rhs159
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 _dafny.Array = _1235_v8
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1235_v8), 0))
				_ = _lhs95
				_lhs93.F2 = _rhs156
				_1258_v22 = _rhs157
				(_lhs94).ArraySet1(_rhs158, (_lhs95).Int())
				_1261_v23 = _rhs159
			}
		} else {
			(globalState).F2 = _1227_v0
			var _1262_v24 bool
			_ = _1262_v24
			_1262_v24 = false
			_1262_v24 = _1262_v24
			var _1263_v25 _dafny.Array
			_ = _1263_v25
			var _nw206 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw206
			_1263_v25 = _nw206
			var _1264_v26 _dafny.CodePoint
			_ = _1264_v26
			_1264_v26 = _dafny.CodePoint('i')
			var _1265_v27 _dafny.Map
			_ = _1265_v27
			_1265_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1262_v24, p1)
			var _1266_v28 _dafny.Sequence
			_ = _1266_v28
			_1266_v28 = _dafny.SeqOf(_1265_v27)
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1263_v25), 0))
			_ = _index173
			(_1263_v25).ArraySet1(Companion_Default___.Fm2(_1264_v26, _1266_v28, _1262_v24, globalState), (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1263_v25), 0))
			_ = _index174
			(_1263_v25).ArraySet1(_1262_v24, (_index174).Int())
			_1265_v27 = (_1265_v27).Merge(_1265_v27)
			var _1267_v29 _dafny.MultiSet
			_ = _1267_v29
			_1267_v29 = _dafny.MultiSetOf(_1227_v0, _1227_v0, _1227_v0)
			var _1268_v30 _dafny.Map
			_ = _1268_v30
			_1268_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v0, _1227_v0)
			var _1269_v31 _dafny.MultiSet
			_ = _1269_v31
			_1269_v31 = _dafny.MultiSetOf(p1)
			var _1270_v32 _dafny.Array
			_ = _1270_v32
			var _nwElement0_51 _dafny.Int = (_1267_v29).Cardinality()
			_ = _nwElement0_51
			var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(18))
			_ = _nw207
			(_nw207).ArraySet1(_nwElement0_51, 0)
			(_nw207).ArraySet1(_1227_v0, 1)
			(_nw207).ArraySet1(_1227_v0, 2)
			(_nw207).ArraySet1(_1227_v0, 3)
			(_nw207).ArraySet1(_1227_v0, 4)
			(_nw207).ArraySet1(_1227_v0, 5)
			(_nw207).ArraySet1(_1227_v0, 6)
			(_nw207).ArraySet1(_dafny.IntOfInt64(297), 7)
			(_nw207).ArraySet1(_1227_v0, 8)
			(_nw207).ArraySet1((Companion_D1_.Create_DC2_(_dafny.IntOfInt64(591), (_1263_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1263_v25), 0))).Int()).(bool), _1227_v0)).Dtor_cf4(), 9)
			(_nw207).ArraySet1(_1227_v0, 10)
			(_nw207).ArraySet1((_1268_v30).Cardinality(), 11)
			(_nw207).ArraySet1(_1227_v0, 12)
			(_nw207).ArraySet1(_1227_v0, 13)
			(_nw207).ArraySet1(_1227_v0, 14)
			(_nw207).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 15)
			(_nw207).ArraySet1(_1227_v0, 16)
			(_nw207).ArraySet1((func() _dafny.Int {
				if (_1269_v31).Contains(p1) {
					return (_1269_v31).Multiplicity(p1)
				}
				return _1227_v0
			})(), 17)
			_1270_v32 = _nw207
			var _1271_v33 _dafny.Map
			_ = _1271_v33
			_1271_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.SetOf(p1, (_1263_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1263_v25), 0))).Int()).(bool))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-718))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_1272_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))).Cardinality())) >= 0, _1270_v32)
			_1271_v33 = (_1271_v33).Merge(_1271_v33)
		}
		var _1273_v34 _dafny.Array
		_ = _1273_v34
		var _nwElement0_52 bool = p1
		_ = _nwElement0_52
		var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(2))
		_ = _nw208
		(_nw208).ArraySet1(_nwElement0_52, 0)
		(_nw208).ArraySet1(p1, 1)
		_1273_v34 = _nw208
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
		_ = _index175
		(_1273_v34).ArraySet1(!(false) || (p1), (_index175).Int())
		var _1274_v35 _dafny.Map
		_ = _1274_v35
		_1274_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
		var _1275_v36 _dafny.Sequence
		_ = _1275_v36
		_1275_v36 = _dafny.SeqOf(_1274_v35)
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
		_ = _index176
		(_1273_v34).ArraySet1(!(!(p1)) || (Companion_Default___.Fm2(Companion_Default___.Fm5(_1227_v0, globalState), _1275_v36, p1, globalState)), (_index176).Int())
		var _1276_v37 *C0
		_ = _1276_v37
		var _nw209 *C0 = New_C0_()
		_ = _nw209
		_nw209.Ctor__(_1227_v0)
		_1276_v37 = _nw209
		var _1277_v38 _dafny.MultiSet
		_ = _1277_v38
		_1277_v38 = _dafny.MultiSetOf(_1227_v0, _dafny.IntOfUint32((p2).Cardinality()), (_1276_v37).F13())
		var _hi5 _dafny.Int = _1227_v0
		_ = _hi5
		for _1278_i5 := Companion_Default___.SafeDivisionInt((_1277_v38).Cardinality(), _1227_v0); _1278_i5.Cmp(_hi5) < 0; _1278_i5 = _1278_i5.Plus(_dafny.One) {
			_1274_v35 = (_1274_v35).Update(p1, true)
			if p1 {
				var _1279_v39 _dafny.Sequence
				_ = _1279_v39
				_1279_v39 = _dafny.SeqOf(_dafny.CodePoint('l'))
				_1279_v39 = p2
				var _1280_v40 _dafny.Sequence
				_ = _1280_v40
				_1280_v40 = _dafny.SeqOf(_1277_v38, _1277_v38, _1277_v38)
				var _1281_v41 _dafny.Set
				_ = _1281_v41
				_1281_v41 = _dafny.SetOf((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
				var _1282_v42 _dafny.Sequence
				_ = _1282_v42
				_1282_v42 = _dafny.SeqOf((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
				var _1283_v43 _dafny.Array
				_ = _1283_v43
				var _nwElement0_53 _dafny.MultiSet = _dafny.MultiSetOf((_1276_v37).F13())
				_ = _nwElement0_53
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(16))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_53, 0)
				(_nw210).ArraySet1(_1277_v38, 1)
				(_nw210).ArraySet1(_1277_v38, 2)
				(_nw210).ArraySet1((_1280_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.IntOfUint32((_1280_v40).Cardinality()))).Uint32()).(_dafny.MultiSet), 3)
				(_nw210).ArraySet1(_1277_v38, 4)
				(_nw210).ArraySet1(_1277_v38, 5)
				(_nw210).ArraySet1(_dafny.MultiSetOf((_1276_v37).F13()), 6)
				(_nw210).ArraySet1(_dafny.MultiSetOf((_1276_v37).F13(), _dafny.IntOfInt64(907)), 7)
				(_nw210).ArraySet1(_1277_v38, 8)
				(_nw210).ArraySet1(_dafny.MultiSetOf(_1278_i5, (_1276_v37).F13()), 9)
				(_nw210).ArraySet1(_1277_v38, 10)
				(_nw210).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ffaybi")).Cardinality()), (_1281_v41).Cardinality(), _1278_i5, (_dafny.Zero).Minus((_1276_v37).F13())), 11)
				(_nw210).ArraySet1(_1277_v38, 12)
				(_nw210).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_1282_v42).Cardinality())), 13)
				(_nw210).ArraySet1(_1277_v38, 14)
				(_nw210).ArraySet1(_1277_v38, 15)
				_1283_v43 = _nw210
				var _1284_v44 _dafny.Map
				_ = _1284_v44
				_1284_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1283_v43, _1227_v0)
				_1284_v44 = (_1284_v44).Update(_1283_v43, (_1276_v37).F13())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _index177
				(_1273_v34).ArraySet1(!(p1), (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _index178
				(_1273_v34).ArraySet1(false, (_index178).Int())
				(globalState).F2 = _dafny.IntOfInt64(804)
			} else {
				var _1285_v45 _dafny.MultiSet
				_ = _1285_v45
				_1285_v45 = _dafny.MultiSetOf(p1, (_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
				var _1286_v46 _dafny.Sequence
				_ = _1286_v46
				_1286_v46 = _dafny.SeqOf((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool), p1, p1)
				var _1287_v47 _dafny.MultiSet
				_ = _1287_v47
				_1287_v47 = _dafny.MultiSetOf((_1285_v45).Difference(_dafny.MultiSetFromSeq(_1286_v46)))
				_1287_v47 = ((func() _dafny.MultiSet {
					if p1 {
						return _1287_v47
					}
					return _1287_v47
				})()).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.MultiSetOf(p1), _dafny.MultiSetOf(p1, Companion_Default___.Fm23(p2, (_1276_v37).F13(), (_1276_v37).F13(), false, globalState)), _1285_v45, _1285_v45)))
				var _1288_v48 D1
				_ = _1288_v48
				_1288_v48 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((p2).Cardinality()), p1, (_1276_v37).F13())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _index179
				(_1273_v34).ArraySet1(((_1288_v48).Dtor_cf4()).Cmp((_1276_v37).F13()) >= 0, (_index179).Int())
				var _1289_v49 _dafny.Array
				_ = _1289_v49
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw211
				_1289_v49 = _nw211
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1289_v49), 0))
				_ = _index180
				(_1289_v49).ArraySet1(Companion_Default___.Fm43(_dafny.MultiSetFromSeq(p0), _1278_i5, _1278_i5, globalState), (_index180).Int())
				var _1290_v50 _dafny.Map
				_ = _1290_v50
				_1290_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus((_1276_v37).F13()))
				var _1291_v51 _dafny.Map
				_ = _1291_v51
				_1291_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1285_v45, _dafny.MultiSetOf((_1276_v37).F13()))
				var _1292_v52 _dafny.Sequence
				_ = _1292_v52
				_1292_v52 = _dafny.SeqOf(_1290_v50)
				var _1293_v53 _dafny.Set
				_ = _1293_v53
				_1293_v53 = _dafny.SetOf((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool), p1)
				var _1294_v54 _dafny.Sequence
				_ = _1294_v54
				_1294_v54 = _dafny.SeqOf(_1290_v50, (_1292_v52).Select((Companion_Default___.SafeIndex((_1293_v53).Cardinality(), _dafny.IntOfUint32((_1292_v52).Cardinality()))).Uint32()).(_dafny.Map))
				var _1295_v56 _dafny.Map
				_ = _1295_v56
				_1295_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1289_v49), 0))
				_ = _index181
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _index182
				var _rhs160 _dafny.Map = (_1291_v51).Merge(_1291_v51)
				_ = _rhs160
				var _rhs161 _dafny.Map = ((_1290_v50).Merge((_1294_v54).Select((Companion_Default___.SafeIndex(_1278_i5, _dafny.IntOfUint32((_1294_v54).Cardinality()))).Uint32()).(_dafny.Map))).Merge(func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter42 := _dafny.Iterate((_1295_v56).Keys().Elements()); ; {
						_compr_41, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _1296_v55 _dafny.Sequence
						_1296_v55 = interface{}(_compr_41).(_dafny.Sequence)
						if (_1295_v56).Contains(_1296_v55) {
							_coll41.Add(_1296_v55, (_1276_v37).F13())
						}
					}
					return _coll41.ToMap()
				}())
				_ = _rhs161
				var _rhs162 bool = ((_1276_v37).F13()).Cmp(_1278_i5) > 0
				_ = _rhs162
				var _lhs96 _dafny.Array = _1289_v49
				_ = _lhs96
				var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1289_v49), 0))
				_ = _lhs97
				var _lhs98 _dafny.Array = _1273_v34
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _lhs99
				(_lhs96).ArraySet1(_rhs160, (_lhs97).Int())
				_1290_v50 = _rhs161
				(_lhs98).ArraySet1(_rhs162, (_lhs99).Int())
				_1227_v0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1276_v37).F13(), Companion_Default___.SafeDivisionInt(_1278_i5, _dafny.IntOfInt64(803)))))
				var _1297_v57 _dafny.Map
				_ = _1297_v57
				_1297_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1278_i5, p1)
				var _1298_v58 _dafny.Map
				_ = _1298_v58
				_1298_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1297_v57, _dafny.IntOfInt64(-541))
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
				_ = _index183
				(_1273_v34).ArraySet1(!(_1298_v58).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1278_i5), p1)), (_index183).Int())
			}
			var _1299_v59 D1
			_ = _1299_v59
			_1299_v59 = Companion_D1_.Create_DC2_((_1276_v37).F13(), (_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool), _1278_i5)
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))
			_ = _index184
			(_1273_v34).ArraySet1((_1299_v59).Dtor_cf3(), (_index184).Int())
			var _1300_v60 D2
			_ = _1300_v60
			_1300_v60 = Companion_D2_.Create_DC4_((_1277_v38).Cardinality())
			_1300_v60 = _1300_v60
		}
		if p1 {
			var _1301_v61 _dafny.Sequence
			_ = _1301_v61
			_1301_v61 = _dafny.UnicodeSeqOfUtf8Bytes("hyoeph")
			_1301_v61 = _1301_v61
			var _1302_v62 _dafny.CodePoint
			_ = _1302_v62
			_1302_v62 = _dafny.CodePoint('c')
			r1 = _1302_v62
			var _1303_v63 _dafny.Array
			_ = _1303_v63
			var _nwElement0_54 _dafny.Int = (_1276_v37).F13()
			_ = _nwElement0_54
			var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(8))
			_ = _nw212
			(_nw212).ArraySet1(_nwElement0_54, 0)
			(_nw212).ArraySet1((_1276_v37).F13(), 1)
			(_nw212).ArraySet1(_1227_v0, 2)
			(_nw212).ArraySet1((_1276_v37).F13(), 3)
			(_nw212).ArraySet1(_dafny.IntOfInt64(159), 4)
			(_nw212).ArraySet1((p0).Select((Companion_Default___.SafeIndex(_1227_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int), 5)
			(_nw212).ArraySet1((_1276_v37).F13(), 6)
			(_nw212).ArraySet1((_dafny.MultiSetFromSeq(p2)).Cardinality(), 7)
			_1303_v63 = _nw212
			var _1304_v64 _dafny.Array
			_ = _1304_v64
			_1304_v64 = _1303_v63
			var _1305_v65 _dafny.Map
			_ = _1305_v65
			_1305_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1276_v37).F13(), _1304_v64)
			var _1306_v66 _dafny.Set
			_ = _1306_v66
			_1306_v66 = _dafny.SetOf(p1)
			var _1307_v67 _dafny.Map
			_ = _1307_v67
			_1307_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v0, (_1276_v37).F13())
			_1305_v65 = (_1305_v65).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12(!(p1), _1227_v0, _1306_v66, _1307_v67, globalState), p0)).Cardinality()), _1304_v64)
			var _1308_v68 *C2
			_ = _1308_v68
			var _nw213 *C2 = New_C2_()
			_ = _nw213
			_nw213.Ctor__((false) && ((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool)), Companion_Default___.SafeModuloInt(_1227_v0, (_1276_v37).F13()), (_this).F3())
			_1308_v68 = _nw213
			var _1309_v69 _dafny.Map
			_ = _1309_v69
			_1309_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1308_v68)
			_1308_v68 = (func() *C2 {
				if (_1309_v69).Contains((_1308_v68).F11()) {
					return (_1309_v69).Get((_1308_v68).F11()).(*C2)
				}
				return _1308_v68
			})()
			var _1310_v70 D4
			_ = _1310_v70
			_1310_v70 = Companion_D4_.Create_DC8_(_1276_v37)
			var _pat_let_tv25 = _1276_v37
			_ = _pat_let_tv25
			_1310_v70 = func(_pat_let13_0 D4) D4 {
				return func(_1311_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let14_0 *C0) D4 {
						return func(_1312_dt__update_hcf15_h0 *C0) D4 {
							return Companion_D4_.Create_DC8_(_1312_dt__update_hcf15_h0)
						}(_pat_let14_0)
					}(_pat_let_tv25)
				}(_pat_let13_0)
			}(_1310_v70)
		} else {
			var _1313_v71 T1
			_ = _1313_v71
			var _nw214 *C2 = New_C2_()
			_ = _nw214
			_nw214.Ctor__(p1, (_1276_v37).F13(), (_this).F3())
			_1313_v71 = _nw214
			var _1314_v72 _dafny.Array
			_ = _1314_v72
			var _nwElement0_55 T1 = _1313_v71
			_ = _nwElement0_55
			var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(2))
			_ = _nw215
			(_nw215).ArraySet1(_nwElement0_55, 0)
			(_nw215).ArraySet1(_1313_v71, 1)
			_1314_v72 = _nw215
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1314_v72), 0))
			_ = _index185
			(_1314_v72).ArraySet1(_1313_v71, (_index185).Int())
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1314_v72), 0))
			_ = _index186
			(_1314_v72).ArraySet1(_1313_v71, (_index186).Int())
			var _1315_v73 _dafny.MultiSet
			_ = _1315_v73
			_1315_v73 = _dafny.MultiSetOf((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool), (_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
			var _1316_v74 _dafny.Sequence
			_ = _1316_v74
			_1316_v74 = _dafny.SeqOf((func() _dafny.Int {
				if (_1315_v73).Contains((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool)) {
					return (_1315_v73).Multiplicity((_1273_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1273_v34), 0))).Int()).(bool))
				}
				return (_1276_v37).F13()
			})(), (_dafny.Zero).Minus((_1276_v37).F13()), _1227_v0, (_1276_v37).F13(), (_1276_v37).F13())
			_1316_v74 = _1316_v74
			var _1317_v75 _dafny.Array
			_ = _1317_v75
			var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
			_ = _nw216
			_1317_v75 = _nw216
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw217
			_1317_v75 = _nw217
			_1315_v73 = _1315_v73
			var _1318_v76 _dafny.Map
			_ = _1318_v76
			_1318_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1276_v37).F13(), (_1276_v37).F13())
			var _1319_v77 D1
			_ = _1319_v77
			_1319_v77 = Companion_D1_.Create_DC2_((func() _dafny.Int {
				if (_1318_v76).Contains((_1276_v37).F13()) {
					return (_1318_v76).Get((_1276_v37).F13()).(_dafny.Int)
				}
				return _1227_v0
			})(), _dafny.Companion_Sequence_.IsPrefixOf(p2, p2), _1227_v0)
			_1319_v77 = _1319_v77
		}
		r0 = _1227_v0
		var _1320_v78 _dafny.CodePoint
		_ = _1320_v78
		_1320_v78 = _dafny.CodePoint('n')
		r1 = _1320_v78
		return r0, r1
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
