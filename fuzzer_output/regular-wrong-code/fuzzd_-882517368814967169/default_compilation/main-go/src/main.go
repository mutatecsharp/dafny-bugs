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
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.UnicodeSeqOfUtf8Bytes("hwfd")
	} else {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(844))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})), _dafny.CodePoint('o'))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(true, !(false) || (true), _dafny.CodePoint('t'), ((_dafny.SetOf(true, false)).Cardinality()).Minus(_dafny.IntOfInt64(31)))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) bool {
	return (_dafny.SetOf(false, true, true, false)).IsProperSubsetOf(_dafny.SetOf(false, false, false))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.IntOfInt64(687)).Times((_dafny.SetOf(_dafny.CodePoint('k'))).Cardinality())), true)
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bkklhmuw"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))), (Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("rw"))).Dtor_cf21())
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(834), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-435))).Cardinality())).Cardinality())).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('j'))).Cardinality())) <= 0)
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.CodePoint('s'), _dafny.CodePoint('u'), _dafny.CodePoint('u'))
		}
		return _dafny.MultiSetOf(_dafny.CodePoint('o'))
	})()).Union((_dafny.MultiSetOf(_dafny.CodePoint('k'))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('g')))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("imlqi")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality(), _dafny.IntOfInt64(881), (_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-937))).Cardinality()))).Intersection(func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(104), _dafny.IntOfInt64(923))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(104)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(923)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_4_v0, _dafny.IntOfInt64(703)))
			}
		}
		return _coll0.ToSet()
	}())).Difference(_dafny.SetOf((_dafny.MultiSetOf(!(true))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.MultiSetOf(true, !(true), false, true)).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(500)))).Union(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_5_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-557)
			}))
		}
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(996), true)).Cardinality())
	})()))
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), true))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ay"), _dafny.UnicodeSeqOfUtf8Bytes("poxqgk")), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("xqxl"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-618))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})), _dafny.SeqOf(_dafny.CodePoint('r')), _dafny.SeqOf(_dafny.CodePoint('d')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_8_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})))).Union(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('w')), _dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('q'))))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('j')
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Sequence, globalState *GlobalState) D0 {
	if true {
		return Companion_D0_.Create_DC0_(true)
	} else {
		return Companion_D0_.Create_DC0_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source0 D7 = Companion_D7_.Create_DC14_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_9_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lvtm")).Cardinality())
	})))
	_ = _source0
	if _source0.Is_DC15() {
		var _10___mcc_h0 _dafny.Int = _source0.Get_().(D7_DC15).Cf27
		_ = _10___mcc_h0
		var _11___mcc_h1 _dafny.CodePoint = _source0.Get_().(D7_DC15).Cf28
		_ = _11___mcc_h1
		var _12_cf28 _dafny.CodePoint = _11___mcc_h1
		_ = _12_cf28
		var _13_cf27 _dafny.Int = _10___mcc_h0
		_ = _13_cf27
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
	} else {
		var _14___mcc_h2 _dafny.Sequence = _source0.Get_().(D7_DC14).Cf26
		_ = _14___mcc_h2
		var _15_cf26 _dafny.Sequence = _14___mcc_h2
		_ = _15_cf26
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), !(false))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khmjbaxct")).Cardinality()), (_dafny.IntOfInt64(696)).Times((_dafny.MultiSetOf(_dafny.CodePoint('b'))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 bool, p2 bool, p3 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("vxkrw")
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC10_(!((!(false)) && (true)), _dafny.IntOfInt64(117), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rbplxx"), func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(801), _dafny.IntOfInt64(233))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _16_v0 _dafny.Int
			_16_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(801)).Cmp(_16_v0) <= 0) && ((_16_v0).Cmp(_dafny.IntOfInt64(233)) < 0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_16_v0, (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality()), true)
			}
		}
		return _coll1.ToMap()
	}()), !(false))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("acqcse"))).Difference(_dafny.MultiSetOf((Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-477))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})))).Dtor_cf21(), _dafny.UnicodeSeqOfUtf8Bytes("oc"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_18_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("teyeal"), _dafny.UnicodeSeqOfUtf8Bytes("jcehpgj")))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(382))).Union((_dafny.SetOf(_dafny.IntOfInt64(962), (_dafny.Zero).Minus(_dafny.IntOfInt64(-593)))).Intersection(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(883))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_2).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfInt64(883))).Contains(_19_v0) {
				_coll2.Add((_19_v0).Times((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality()))
			}
		}
		return _coll2.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-334))).Cardinality()), _dafny.IntOfInt64(564))).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-49))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_20_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(913)
	})))), _dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-550))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ulsmcuwk")).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(150))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_21_i1 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(false, true)
	}))).Cardinality()), _dafny.IntOfInt64(80))).Cardinality(), _dafny.IntOfInt64(561)))).Cardinality(), _dafny.IntOfInt64(374), _dafny.IntOfInt64(-739), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(25)))), _dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-788))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(621), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(909))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_23_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(782), (_dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(750))).Cardinality())).Cardinality())))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(134))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(300))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(-30), _dafny.IntOfInt64(30), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fpgtqpqa")).Cardinality()), _dafny.IntOfInt64(159)))), _dafny.SeqOf(Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ppysr")).Cardinality()), (_dafny.SetOf(false, true)).Cardinality(), _dafny.IntOfInt64(636), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gvfexkl")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_24_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qmdcartnm")).Cardinality()))), Companion_D7_.Create_DC14_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-803))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_25_i3 _dafny.Int) _dafny.Int {
		return (func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(577), _dafny.IntOfInt64(-231))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _26_v0 _dafny.Int
				_26_v0 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(577)).Cmp(_26_v0) <= 0) && ((_26_v0).Cmp(_dafny.IntOfInt64(-231)) < 0) {
					_coll3.Add((_26_v0).Plus(_dafny.IntOfInt64(8)))
				}
			}
			return _coll3.ToSet()
		}()).Cardinality()
	}))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(-929))), Companion_D7_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(612)))))))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if !(true) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ckmfb"), false)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pfwhyaky"), false)
	})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("hvc"), !(true))))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.CodePoint, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Difference((_dafny.MultiSetOf(_dafny.IntOfInt64(257), (_dafny.SetOf(false)).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(482))))).Union((_dafny.MultiSetOf((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-548), _dafny.IntOfInt64(152))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _28_v0 _dafny.Int
			_28_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-548)).Cmp(_28_v0) <= 0) && ((_28_v0).Cmp(_dafny.IntOfInt64(152)) < 0) {
				_coll4.Add((_28_v0).Minus(_dafny.IntOfInt64(896)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxd")).Cardinality()))).Cardinality())
			}
		}
		return _coll4.ToMap()
	}()).Cardinality())).Intersection(_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(395)))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('k'), _dafny.CodePoint('j'), _dafny.CodePoint('y'))).Cardinality()), _dafny.IntOfInt64(861)))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('x')
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true)).Intersection(_dafny.MultiSetOf(false, false, !(false)))
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false, false)).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.CodePoint('w'), _dafny.CodePoint('j'))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(!(true), true), _dafny.MultiSetOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), true)), _dafny.MultiSetOf(false, true, true, true)), _dafny.SeqOf(_dafny.MultiSetOf(false))))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Sequence, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return (Companion_D14_.Create_DC26_(Companion_D0_.Create_DC0_(!(false)), _dafny.SetOf(false, true, false, false, false))).Dtor_cf44()
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC11_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cxwyhsabm"), _dafny.UnicodeSeqOfUtf8Bytes("gojgfnrgd")))
}
func (_static *CompanionStruct_Default___) Fm43(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.MultiSetOf(_dafny.IntOfInt64(212)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.MultiSetOf(_dafny.IntOfInt64(383), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lu")).Cardinality()))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(348))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_29_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()), _dafny.IntOfInt64(-327))))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(-445), _dafny.IntOfInt64(-810), _dafny.IntOfInt64(358)), true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(777)), !(true))).Merge(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(381)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_30_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		})))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _31_v0 _dafny.Sequence
			_31_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(381)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_30_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
			})))).Contains(_31_v0) {
				_coll5.Add(_31_v0, false)
			}
		}
		return _coll5.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC22_(!(!(false) || (true)))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Set, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC21_((func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('k')
		}
		return _dafny.CodePoint('e')
	})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pldd")).Cardinality()), _dafny.IntOfInt64(-615))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-131))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.IntOfInt64(391))))
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(788))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (Companion_D5_.Create_DC10_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('b')))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("afr"), func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(249), _dafny.IntOfInt64(-663))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _32_v0 _dafny.Int
			_32_v0 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(249)).Cmp(_32_v0) <= 0) && ((_32_v0).Cmp(_dafny.IntOfInt64(-663)) < 0) {
				_coll6.Add((_32_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahorcb")).Cardinality())), false)
			}
		}
		return _coll6.ToMap()
	}()), false)).Dtor_cf17()))
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D7 {
	var _source1 D14 = Companion_D14_.Create_DC27_(!(true))
	_ = _source1
	if _source1.Is_DC26() {
		var _33___mcc_h0 D0 = _source1.Get_().(D14_DC26).Cf43
		_ = _33___mcc_h0
		var _34___mcc_h1 _dafny.Set = _source1.Get_().(D14_DC26).Cf44
		_ = _34___mcc_h1
		var _35_cf44 _dafny.Set = _34___mcc_h1
		_ = _35_cf44
		var _36_cf43 D0 = _33___mcc_h0
		_ = _36_cf43
		return Companion_D7_.Create_DC15_(_dafny.IntOfInt64(85), _dafny.CodePoint('d'))
	} else if _source1.Is_DC27() {
		var _37___mcc_h2 bool = _source1.Get_().(D14_DC27).Cf45
		_ = _37___mcc_h2
		var _38_cf45 bool = _37___mcc_h2
		_ = _38_cf45
		return Companion_D7_.Create_DC15_((_dafny.SetOf(_38_cf45, _38_cf45, _38_cf45, _38_cf45, _38_cf45)).Cardinality(), _dafny.CodePoint('b'))
	} else if _source1.Is_DC28() {
		return Companion_D7_.Create_DC15_(_dafny.IntOfInt64(404), _dafny.CodePoint('n'))
	} else {
		var _39___mcc_h3 _dafny.Array = _source1.Get_().(D14_DC25).Cf42
		_ = _39___mcc_h3
		var _40_cf42 _dafny.Array = _39___mcc_h3
		_ = _40_cf42
		return Companion_D7_.Create_DC15_(_dafny.IntOfInt64(717), _dafny.CodePoint('m'))
	}
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC28_()
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(Companion_D6_.Create_DC11_((Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_41_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	})))).Dtor_cf21()), Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("lqncemshr")))).Intersection(_dafny.SetOf(Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(718))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_42_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))), Companion_D6_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("ghyhpuqt"))))
}
func (_static *CompanionStruct_Default___) Fm53(p0 bool, p1 bool, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC26_(Companion_D0_.Create_DC0_(true), _dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-926))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_43_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-743)
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(845), _dafny.IntOfInt64(377))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _44_v0 _dafny.Int
			_44_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(845)).Cmp(_44_v0) <= 0) && ((_44_v0).Cmp(_dafny.IntOfInt64(377)) < 0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_44_v0, _dafny.IntOfInt64(674)), false)
			}
		}
		return _coll7.ToMap()
	}()).Cardinality())).Cardinality(), ((Companion_D24_.Create_DC48_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.IntOfInt64(-110))).Cardinality()))).Dtor_cf67()).Cardinality(), (_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(893))).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality())).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_45_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(426), !(!(false)))).Cardinality()
	}))))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 _dafny.Sequence, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false)), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true))
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(747)), _dafny.MultiSetOf(_dafny.IntOfInt64(-652)))).Union((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(745), _dafny.IntOfInt64(760)))).Union(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-439)))))
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, globalState *GlobalState) _dafny.Array {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var _46_v0 _dafny.Int
	_ = _46_v0
	_46_v0 = _dafny.IntOfInt64(-220)
	var _47_v1 *C2
	_ = _47_v1
	var _nw0 *C2 = New_C2_()
	_ = _nw0
	_nw0.Ctor__(_46_v0)
	_47_v1 = _nw0
	var _48_v2 bool
	_ = _48_v2
	_48_v2 = true
	_48_v2 = Companion_Default___.Fm5(globalState)
	var _49_v3 T0
	_ = _49_v3
	var _nw1 *C2 = New_C2_()
	_ = _nw1
	_nw1.Ctor__(_46_v0)
	_49_v3 = _nw1
	var _50_v4 T0
	_ = _50_v4
	_50_v4 = _49_v3
	var _source2 T0 = _50_v4
	_ = _source2
	var _51___mcc_h0 T0 = _source2
	_ = _51___mcc_h0
	var _52_cf53 T0 = _51___mcc_h0
	_ = _52_cf53
	var _53_v5 _dafny.Sequence
	_ = _53_v5
	_53_v5 = _dafny.UnicodeSeqOfUtf8Bytes("dxgtgy")
	var _54_v6 _dafny.Array
	_ = _54_v6
	var _nwElement0_0 T0 = (_50_v4)
	_ = _nwElement0_0
	var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(25))
	_ = _nw2
	(_nw2).ArraySet1(_nwElement0_0, 0)
	(_nw2).ArraySet1(_52_cf53, 1)
	(_nw2).ArraySet1(_49_v3, 2)
	(_nw2).ArraySet1(_52_cf53, 3)
	(_nw2).ArraySet1(_49_v3, 4)
	(_nw2).ArraySet1(_52_cf53, 5)
	(_nw2).ArraySet1(_49_v3, 6)
	(_nw2).ArraySet1(_52_cf53, 7)
	(_nw2).ArraySet1(_52_cf53, 8)
	(_nw2).ArraySet1(_49_v3, 9)
	(_nw2).ArraySet1(_52_cf53, 10)
	(_nw2).ArraySet1(_52_cf53, 11)
	(_nw2).ArraySet1(_49_v3, 12)
	(_nw2).ArraySet1(_49_v3, 13)
	(_nw2).ArraySet1(_52_cf53, 14)
	(_nw2).ArraySet1(_49_v3, 15)
	(_nw2).ArraySet1(_49_v3, 16)
	(_nw2).ArraySet1(_52_cf53, 17)
	(_nw2).ArraySet1(_52_cf53, 18)
	(_nw2).ArraySet1(_52_cf53, 19)
	(_nw2).ArraySet1(_49_v3, 20)
	(_nw2).ArraySet1(_49_v3, 21)
	(_nw2).ArraySet1(_52_cf53, 22)
	(_nw2).ArraySet1(_49_v3, 23)
	(_nw2).ArraySet1(_49_v3, 24)
	_54_v6 = _nw2
	var _55_v7 _dafny.Array
	_ = _55_v7
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_0
	var _nw3 _dafny.Array
	_ = _nw3
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw3 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_56_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_57_i0 _dafny.Int) _dafny.Int {
				return (_57_i0).Minus(_56_v0)
			}
		})(_46_v0)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw3 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw3).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw3).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_55_v7 = _nw3
	var _58_v8 _dafny.Map
	_ = _58_v8
	_58_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v7, _54_v6)
	var _59_v9 _dafny.Array
	_ = _59_v9
	var _nwElement0_1 _dafny.Array = _54_v6
	_ = _nwElement0_1
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(24))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_1, 0)
	(_nw4).ArraySet1(_54_v6, 1)
	(_nw4).ArraySet1(_54_v6, 2)
	(_nw4).ArraySet1(_54_v6, 3)
	(_nw4).ArraySet1(_54_v6, 4)
	(_nw4).ArraySet1(_54_v6, 5)
	(_nw4).ArraySet1(_54_v6, 6)
	(_nw4).ArraySet1(_54_v6, 7)
	(_nw4).ArraySet1(_54_v6, 8)
	(_nw4).ArraySet1(_54_v6, 9)
	(_nw4).ArraySet1(_54_v6, 10)
	(_nw4).ArraySet1(_54_v6, 11)
	(_nw4).ArraySet1(_54_v6, 12)
	(_nw4).ArraySet1(_54_v6, 13)
	(_nw4).ArraySet1(_54_v6, 14)
	(_nw4).ArraySet1(_54_v6, 15)
	(_nw4).ArraySet1(_54_v6, 16)
	(_nw4).ArraySet1(_54_v6, 17)
	(_nw4).ArraySet1((func() _dafny.Array {
		if (_58_v8).Contains(_55_v7) {
			return (_58_v8).Get(_55_v7).(_dafny.Array)
		}
		return _54_v6
	})(), 18)
	(_nw4).ArraySet1(_54_v6, 19)
	(_nw4).ArraySet1(_54_v6, 20)
	(_nw4).ArraySet1(_54_v6, 21)
	(_nw4).ArraySet1(_54_v6, 22)
	(_nw4).ArraySet1(_54_v6, 23)
	_59_v9 = _nw4
	var _60_v10 _dafny.Map
	_ = _60_v10
	_60_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v5, _59_v9)
	_60_v10 = (_60_v10).Update(_53_v5, _59_v9)
	if !(_48_v2) {
		var _61_v11 _dafny.Array
		_ = _61_v11
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw5
		_61_v11 = _nw5
		var _62_v12 _dafny.Sequence
		_ = _62_v12
		_62_v12 = _dafny.SeqOf(_61_v11, _61_v11)
		var _63_v13 _dafny.MultiSet
		_ = _63_v13
		_63_v13 = _dafny.MultiSetOf(_dafny.IntOfInt64(-971), _46_v0)
		var _64_v14 _dafny.Set
		_ = _64_v14
		_64_v14 = _dafny.SetOf(!_dafny.Companion_Sequence_.Equal(_62_v12, _62_v12), !((_63_v13).IsDisjointFrom(_63_v13)), (_47_v1).Fm8(_53_v5, p0, globalState), p1)
		var _65_v15 _dafny.Array
		_ = _65_v15
		var _nwElement0_2 T0 = _50_v4
		_ = _nwElement0_2
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(13))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_2, 0)
		(_nw6).ArraySet1(_50_v4, 1)
		(_nw6).ArraySet1(_50_v4, 2)
		(_nw6).ArraySet1(_50_v4, 3)
		(_nw6).ArraySet1(_50_v4, 4)
		(_nw6).ArraySet1(_50_v4, 5)
		(_nw6).ArraySet1(_50_v4, 6)
		(_nw6).ArraySet1(_50_v4, 7)
		(_nw6).ArraySet1(_50_v4, 8)
		(_nw6).ArraySet1(_52_cf53, 9)
		(_nw6).ArraySet1(_50_v4, 10)
		(_nw6).ArraySet1(_50_v4, 11)
		(_nw6).ArraySet1(_50_v4, 12)
		_65_v15 = _nw6
		var _rhs0 _dafny.Int = _46_v0
		_ = _rhs0
		var _rhs1 _dafny.Set = _64_v14
		_ = _rhs1
		var _rhs2 _dafny.Array = _65_v15
		_ = _rhs2
		_46_v0 = _rhs0
		_64_v14 = _rhs1
		_65_v15 = _rhs2
		var _66_v16 _dafny.Map
		_ = _66_v16
		_66_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_53_v5).Cardinality()), !(_48_v2))
		var _67_v17 _dafny.Map
		_ = _67_v17
		_67_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v5, _66_v16)
		var _68_v18 D5
		_ = _68_v18
		_68_v18 = Companion_D5_.Create_DC10_(p1, _46_v0, _67_v17, p0)
		var _69_v19 _dafny.Map
		_ = _69_v19
		_69_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _68_v18)
		_69_v19 = (func() _dafny.Map {
			if p0 {
				return _69_v19
			}
			return _69_v19
		})()
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_61_v11), 0))
		_ = _index0
		(_61_v11).ArraySet1(p1, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_61_v11), 0))
		_ = _index1
		(_61_v11).ArraySet1(p0, (_index1).Int())
		var _70_v20 *C7
		_ = _70_v20
		var _nw7 *C7 = New_C7_()
		_ = _nw7
		_nw7.Ctor__()
		_70_v20 = _nw7
		var _71_v21 D20
		_ = _71_v21
		_71_v21 = Companion_D20_.Create_DC38_(_70_v20)
		_70_v20 = (_71_v21).Dtor_cf57()
		var _72_v22 D11
		_ = _72_v22
		_72_v22 = Companion_D11_.Create_DC22_(p1)
		var _73_v23 _dafny.Map
		_ = _73_v23
		_73_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v22, _46_v0)
		var _74_v24 _dafny.Set
		_ = _74_v24
		_74_v24 = _dafny.SetOf(_73_v23, _73_v23, _73_v23)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_55_v7), 0))
		_ = _index2
		(_55_v7).ArraySet1((_74_v24).Cardinality(), (_index2).Int())
		var _75_v25 _dafny.Map
		_ = _75_v25
		_75_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-954))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_76_v11 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_77_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), (_76_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_76_v11), 0))).Int()).(bool))
			}
		})(_61_v11))), (_49_v3).Fm1((_61_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_61_v11), 0))).Int()).(bool), globalState))
		var _78_v26 _dafny.CodePoint
		_ = _78_v26
		_78_v26 = _dafny.CodePoint('l')
		var _79_v27 _dafny.Map
		_ = _79_v27
		_79_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v26, p1)
		var _80_v28 _dafny.Sequence
		_ = _80_v28
		_80_v28 = _dafny.SeqOf(p0)
		var _81_v29 _dafny.Sequence
		_ = _81_v29
		_81_v29 = _dafny.SeqOf(_79_v27, Companion_Default___.Fm49(_80_v28, !(p0), _46_v0, globalState))
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_55_v7), 0))
		_ = _index3
		(_55_v7).ArraySet1((func() _dafny.Int {
			if (_75_v25).Contains(_81_v29) {
				return (_75_v25).Get(_81_v29).(_dafny.Int)
			}
			return _46_v0
		})(), (_index3).Int())
	} else {
		_46_v0 = (_dafny.Zero).Minus(_46_v0)
		var _82_v30 *C3
		_ = _82_v30
		var _nw8 *C3 = New_C3_()
		_ = _nw8
		_nw8.Ctor__(((_dafny.Zero).Minus(_46_v0)).Plus(_46_v0))
		_82_v30 = _nw8
		var _83_v31 *C4
		_ = _83_v31
		var _nw9 *C4 = New_C4_()
		_ = _nw9
		_nw9.Ctor__(_46_v0)
		_83_v31 = _nw9
		var _84_v32 _dafny.CodePoint
		_ = _84_v32
		_84_v32 = _dafny.CodePoint('u')
		var _85_v33 _dafny.Array
		_ = _85_v33
		var _nwElement0_3 _dafny.CodePoint = _84_v32
		_ = _nwElement0_3
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.One)
		_ = _nw10
		(_nw10).ArraySet1CodePoint(_nwElement0_3, 0)
		_85_v33 = _nw10
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_85_v33), 0))
		_ = _index4
		(_85_v33).ArraySet1CodePoint(_84_v32, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_85_v33), 0))
		_ = _index5
		(_85_v33).ArraySet1CodePoint(_84_v32, (_index5).Int())
		var _86_v34 _dafny.Array
		_ = _86_v34
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw11
		_86_v34 = _nw11
		var _87_v35 _dafny.Map
		_ = _87_v35
		_87_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _46_v0)
		var _88_v36 _dafny.MultiSet
		_ = _88_v36
		_88_v36 = _dafny.MultiSetOf(_48_v2)
		var _89_v37 _dafny.Map
		_ = _89_v37
		_89_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _47_v1)
		var _90_v38 _dafny.Map
		_ = _90_v38
		_90_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v0, _46_v0)
		var _91_v39 _dafny.MultiSet
		_ = _91_v39
		_91_v39 = _dafny.MultiSetOf(_90_v38)
		var _nwElement0_4 bool = p0
		_ = _nwElement0_4
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(18))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_4, 0)
		(_nw12).ArraySet1((_87_v35).Equals(_87_v35), 1)
		(_nw12).ArraySet1(p0, 2)
		(_nw12).ArraySet1(!((_88_v36).IsDisjointFrom(_88_v36)), 3)
		(_nw12).ArraySet1(p0, 4)
		(_nw12).ArraySet1((_89_v37).Contains(p1), 5)
		(_nw12).ArraySet1((_48_v2) == (_48_v2), 6)
		(_nw12).ArraySet1(p0, 7)
		(_nw12).ArraySet1(_48_v2, 8)
		(_nw12).ArraySet1((func() bool {
			if p1 {
				return p0
			}
			return _48_v2
		})(), 9)
		(_nw12).ArraySet1((func() bool {
			if _48_v2 {
				return _48_v2
			}
			return p1
		})(), 10)
		(_nw12).ArraySet1(_48_v2, 11)
		(_nw12).ArraySet1(p1, 12)
		(_nw12).ArraySet1(Companion_Default___.Fm3(_48_v2, _dafny.IntOfUint32((_dafny.SeqOf(_82_v30)).Cardinality()), _46_v0, globalState), 13)
		(_nw12).ArraySet1(p0, 14)
		(_nw12).ArraySet1(p1, 15)
		(_nw12).ArraySet1((_91_v39).IsDisjointFrom(_91_v39), 16)
		(_nw12).ArraySet1((_88_v36).IsSubsetOf(_88_v36), 17)
		_86_v34 = _nw12
	}
	var _92_v40 _dafny.Map
	_ = _92_v40
	_92_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
	var _93_v41 _dafny.Map
	_ = _93_v41
	_93_v41 = _92_v40
	var _94_v42 _dafny.Sequence
	_ = _94_v42
	_94_v42 = _dafny.SeqOf(_46_v0)
	var _95_v43 _dafny.MultiSet
	_ = _95_v43
	_95_v43 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_94_v42, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_53_v5).Cardinality()), _dafny.IntOfUint32((_94_v42).Cardinality()))).Uint32(), _46_v0)).Cardinality()))
	var _96_v44 _dafny.Sequence
	_ = _96_v44
	_96_v44 = _dafny.SeqOf((_92_v40).Update(!(true), p1), _92_v40)
	var _97_v45 _dafny.Map
	_ = _97_v45
	_97_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p1), p0)
	var _98_v46 _dafny.Array
	_ = _98_v46
	var _nwElement0_5 _dafny.Map = _93_v41
	_ = _nwElement0_5
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
	_ = _nw13
	(_nw13).ArraySet1(_nwElement0_5, 0)
	(_nw13).ArraySet1(_93_v41, 1)
	(_nw13).ArraySet1(_93_v41, 2)
	(_nw13).ArraySet1(_93_v41, 3)
	(_nw13).ArraySet1(_92_v40, 4)
	(_nw13).ArraySet1(_93_v41, 5)
	(_nw13).ArraySet1(_93_v41, 6)
	(_nw13).ArraySet1(_92_v40, 7)
	(_nw13).ArraySet1(_93_v41, 8)
	(_nw13).ArraySet1(_93_v41, 9)
	(_nw13).ArraySet1(_93_v41, 10)
	(_nw13).ArraySet1(_93_v41, 11)
	(_nw13).ArraySet1(_93_v41, 12)
	(_nw13).ArraySet1(_93_v41, 13)
	(_nw13).ArraySet1(_93_v41, 14)
	(_nw13).ArraySet1(Companion_Default___.Fm57(_95_v43, globalState), 15)
	(_nw13).ArraySet1(Companion_Default___.Fm57(_95_v43, globalState), 16)
	(_nw13).ArraySet1(_93_v41, 17)
	(_nw13).ArraySet1((_96_v44).Select((Companion_Default___.SafeIndex(_46_v0, _dafny.IntOfUint32((_96_v44).Cardinality()))).Uint32()).(_dafny.Map), 18)
	(_nw13).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if (_97_v45).Contains(_dafny.SeqOf(_48_v2, p0)) {
			return (_97_v45).Get(_dafny.SeqOf(_48_v2, p0)).(bool)
		}
		return !(!(_48_v2))
	})(), p1), 19)
	(_nw13).ArraySet1(_93_v41, 20)
	(_nw13).ArraySet1(_92_v40, 21)
	(_nw13).ArraySet1(_93_v41, 22)
	(_nw13).ArraySet1(_93_v41, 23)
	(_nw13).ArraySet1(_93_v41, 24)
	_98_v46 = _nw13
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_98_v46), 0))
	_ = _index6
	(_98_v46).ArraySet1(_93_v41, (_index6).Int())
	var _99_v47 *C7
	_ = _99_v47
	var _nw14 *C7 = New_C7_()
	_ = _nw14
	_nw14.Ctor__()
	_99_v47 = _nw14
	var _100_v48 *C1
	_ = _100_v48
	var _nw15 *C1 = New_C1_()
	_ = _nw15
	_nw15.Ctor__(_53_v5)
	_100_v48 = _nw15
	var _101_v49 _dafny.Sequence
	_ = _101_v49
	_101_v49 = _dafny.SeqOf(_100_v48)
	var _102_v50 _dafny.Set
	_ = _102_v50
	_102_v50 = _dafny.SetOf(p0, p1, !(_dafny.Companion_Sequence_.Contains(_101_v49, (_101_v49).Select((Companion_Default___.SafeIndex((_47_v1).Fm1(!(p0), globalState), _dafny.IntOfUint32((_101_v49).Cardinality()))).Uint32()).(*C1))), p1)
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_98_v46), 0))
	_ = _index7
	var _rhs3 _dafny.Map = Companion_Default___.Fm57(_95_v43, globalState)
	_ = _rhs3
	var _rhs4 _dafny.Int = _46_v0
	_ = _rhs4
	var _rhs5 *C7 = _99_v47
	_ = _rhs5
	var _rhs6 _dafny.Set = (_102_v50).Difference((Companion_Default___.Fm41(_94_v42, p0, p1, _48_v2, globalState)).Union(_dafny.SetOf(!(_48_v2))))
	_ = _rhs6
	var _lhs0 _dafny.Array = _98_v46
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_98_v46), 0))
	_ = _lhs1
	(_lhs0).ArraySet1(_rhs3, (_lhs1).Int())
	_46_v0 = _rhs4
	_99_v47 = _rhs5
	_102_v50 = _rhs6
	_46_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_46_v0, _dafny.IntOfInt64(-418)), _46_v0))
	var _103_v51 T2
	_ = _103_v51
	var _nw16 *C3 = New_C3_()
	_ = _nw16
	_nw16.Ctor__(_dafny.IntOfInt64(735))
	_103_v51 = _nw16
	var _104_v52 T2
	_ = _104_v52
	_104_v52 = _103_v51
	var _source3 T2 = _104_v52
	_ = _source3
	var _105___mcc_h1 T2 = _source3
	_ = _105___mcc_h1
	var _106_cf33 T2 = _105___mcc_h1
	_ = _106_cf33
	var _107_v53 *C6
	_ = _107_v53
	var _nw17 *C6 = New_C6_()
	_ = _nw17
	_nw17.Ctor__()
	_107_v53 = _nw17
	var _108_v54 _dafny.Array
	_ = _108_v54
	var _nwElement0_6 _dafny.Int = (_103_v51.F2()).Plus(_103_v51.F2())
	_ = _nwElement0_6
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(6))
	_ = _nw18
	(_nw18).ArraySet1(_nwElement0_6, 0)
	(_nw18).ArraySet1(_103_v51.F2(), 1)
	(_nw18).ArraySet1((_46_v0).Times(_106_cf33.F2()), 2)
	(_nw18).ArraySet1((_dafny.IntOfInt64(997)).Minus(_dafny.IntOfInt64(-856)), 3)
	(_nw18).ArraySet1(_dafny.IntOfInt64(721), 4)
	(_nw18).ArraySet1(_46_v0, 5)
	_108_v54 = _nw18
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_108_v54), 0))
	_ = _index8
	(_108_v54).ArraySet1((_103_v51.F2()).Minus(_103_v51.F2()), (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_108_v54), 0))
	_ = _index9
	(_108_v54).ArraySet1(_106_cf33.F2(), (_index9).Int())
	var _109_v55 _dafny.Array
	_ = _109_v55
	var _len0_1 _dafny.Int = _dafny.One
	_ = _len0_1
	var _nw19 _dafny.Array
	_ = _nw19
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw19 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = func(_110_i2 _dafny.Int) bool {
			return true
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw19 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw19).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw19).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_109_v55 = _nw19
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_109_v55), 0))
	_ = _index10
	(_109_v55).ArraySet1(p0, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_109_v55), 0))
	_ = _index11
	(_109_v55).ArraySet1(p0, (_index11).Int())
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_109_v55), 0))
	_ = _index12
	(_109_v55).ArraySet1(p0, (_index12).Int())
	var _111_v56 _dafny.Array
	_ = _111_v56
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_2
	var _nw20 _dafny.Array
	_ = _nw20
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw20 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = (func(_112_p0 bool) func(_dafny.Int) bool {
			return func(_113_i3 _dafny.Int) bool {
				return _112_p0
			}
		})(p0)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw20 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw20).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw20).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_111_v56 = _nw20
	var _pat_let_tv0 = _111_v56
	_ = _pat_let_tv0
	var _pat_let_tv1 = _111_v56
	_ = _pat_let_tv1
	var _source4 D11 = func(_pat_let0_0 D11) D11 {
		return func(_116_dt__update__tmp_h3 D11) D11 {
			return func(_pat_let3_0 _dafny.Array) D11 {
				return func(_117_dt__update_hcf35_h1 _dafny.Array) D11 {
					return Companion_D11_.Create_DC20_(_117_dt__update_hcf35_h1)
				}(_pat_let3_0)
			}(_pat_let_tv1)
		}(_pat_let0_0)
	}(func(_pat_let1_0 D11) D11 {
		return func(_114_dt__update__tmp_h2 D11) D11 {
			return func(_pat_let2_0 _dafny.Array) D11 {
				return func(_115_dt__update_hcf35_h0 _dafny.Array) D11 {
					return Companion_D11_.Create_DC20_(_115_dt__update_hcf35_h0)
				}(_pat_let2_0)
			}(_pat_let_tv0)
		}(_pat_let1_0)
	}(Companion_D11_.Create_DC20_(_111_v56)))
	_ = _source4
	if _source4.Is_DC21() {
		var _118___mcc_h2 _dafny.CodePoint = _source4.Get_().(D11_DC21).Cf36
		_ = _118___mcc_h2
		var _119___mcc_h3 _dafny.Int = _source4.Get_().(D11_DC21).Cf37
		_ = _119___mcc_h3
		var _120___mcc_h4 _dafny.Int = _source4.Get_().(D11_DC21).Cf38
		_ = _120___mcc_h4
		var _121_cf38 _dafny.Int = _120___mcc_h4
		_ = _121_cf38
		var _122_cf37 _dafny.Int = _119___mcc_h3
		_ = _122_cf37
		var _123_cf36 _dafny.CodePoint = _118___mcc_h2
		_ = _123_cf36
		_48_v2 = p0
		var _nw21 *C3 = New_C3_()
		_ = _nw21
		_nw21.Ctor__(_122_cf37)
		_103_v51 = _nw21
		var _124_v57 _dafny.Sequence
		_ = _124_v57
		_124_v57 = _dafny.SeqOf(_103_v51.F2())
		var _125_v58 _dafny.Sequence
		_ = _125_v58
		_125_v58 = _dafny.SeqOf(_48_v2, _48_v2)
		var _126_v59 _dafny.Map
		_ = _126_v59
		_126_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_124_v57).Cardinality()), _125_v58)
		var _127_v60 _dafny.Map
		_ = _127_v60
		_127_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v51.F2(), _122_cf37)
		var _128_v61 _dafny.MultiSet
		_ = _128_v61
		_128_v61 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_122_cf37)), _103_v51.F2(), _46_v0, (_126_v59).Cardinality(), (_dafny.Zero).Minus((func() _dafny.Int {
			if (_127_v60).Contains(_103_v51.F2()) {
				return (_127_v60).Get(_103_v51.F2()).(_dafny.Int)
			}
			return _103_v51.F2()
		})()))
		var _129_v62 _dafny.Sequence
		_ = _129_v62
		_129_v62 = _dafny.SeqOf(_128_v61, (_dafny.MultiSetOf((_dafny.Zero).Minus(_122_cf37))).Intersection(_128_v61), _128_v61)
		_129_v62 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer26 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_130_cf37 _dafny.Int, _131_cf38 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
			return func(_132_i4 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_130_cf37, _131_cf38)
			}
		})(_122_cf37, _121_cf38)))
		(_103_v51).F2_set_((((_47_v1).Fm1(_48_v2, globalState)).Minus(_46_v0)).Plus(Companion_Default___.SafeModuloInt(_121_cf38, (func() _dafny.Int {
			if (_128_v61).Contains(_103_v51.F2()) {
				return (_128_v61).Multiplicity(_103_v51.F2())
			}
			return _121_cf38
		})())))
	} else if _source4.Is_DC22() {
		var _133___mcc_h5 bool = _source4.Get_().(D11_DC22).Cf39
		_ = _133___mcc_h5
		var _134_cf39 bool = _133___mcc_h5
		_ = _134_cf39
		if p0 {
			_48_v2 = !(p0)
			_134_cf39 = !(Companion_Default___.Fm5(globalState))
			_48_v2 = p1
			var _135_v63 _dafny.Array
			_ = _135_v63
			var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw22
			_135_v63 = _nw22
			var _136_v64 _dafny.Array
			_ = _136_v64
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw23
			_136_v64 = _nw23
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_135_v63), 0))
			_ = _index13
			(_135_v63).ArraySet1(_136_v64, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_135_v63), 0))
			_ = _index14
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw24
			(_135_v63).ArraySet1(_nw24, (_index14).Int())
			(_103_v51).F2_set_((_dafny.Zero).Minus(_103_v51.F2()))
		} else {
			(_103_v51).F2_set_(((_103_v51.F2()).Times(_103_v51.F2())).Plus(_103_v51.F2()))
			_48_v2 = p0
			(_103_v51).F2_set_(_46_v0)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_111_v56), 0))
			_ = _index15
			(_111_v56).ArraySet1(_48_v2, (_index15).Int())
			var _137_v65 _dafny.Sequence
			_ = _137_v65
			_137_v65 = _dafny.SeqOf(_134_cf39, (func() bool {
				if p0 {
					return _48_v2
				}
				return _48_v2
			})(), !(true) || (p1))
			var _138_v66 _dafny.Sequence
			_ = _138_v66
			_138_v66 = _dafny.UnicodeSeqOfUtf8Bytes("pekaqfyha")
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_111_v56), 0))
			_ = _index16
			(_111_v56).ArraySet1((_137_v65).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm54(_134_cf39, _138_v66, globalState)).Plus(_46_v0), _dafny.IntOfUint32((_137_v65).Cardinality()))).Uint32()).(bool), (_index16).Int())
			var _139_v67 *C6
			_ = _139_v67
			var _nw25 *C6 = New_C6_()
			_ = _nw25
			_nw25.Ctor__()
			_139_v67 = _nw25
			_139_v67 = _139_v67
		}
		var _140_v68 _dafny.Sequence
		_ = _140_v68
		_140_v68 = _dafny.UnicodeSeqOfUtf8Bytes("bkantehx")
		var _141_v69 *C7
		_ = _141_v69
		var _nw26 *C7 = New_C7_()
		_ = _nw26
		_nw26.Ctor__()
		_141_v69 = _nw26
		var _142_v70 D20
		_ = _142_v70
		_142_v70 = Companion_D20_.Create_DC38_(_141_v69)
		var _rhs7 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("bjadvjecw")
		_ = _rhs7
		var _rhs8 D20 = _142_v70
		_ = _rhs8
		var _rhs9 _dafny.Int = _46_v0
		_ = _rhs9
		var _lhs2 T2 = _103_v51
		_ = _lhs2
		_140_v68 = _rhs7
		_142_v70 = _rhs8
		_lhs2.F2_set_(_rhs9)
		var _143_v71 _dafny.CodePoint
		_ = _143_v71
		_143_v71 = _dafny.CodePoint('b')
		var _144_v72 _dafny.Sequence
		_ = _144_v72
		_144_v72 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_140_v68, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("s"), (Companion_Default___.SafeIndex(_103_v51.F2(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))).Uint32(), _143_v71)))
		_144_v72 = _144_v72
		var _145_v73 _dafny.Array
		_ = _145_v73
		var _nw27 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
		_ = _nw27
		_145_v73 = _nw27
		var _146_v74 *C3
		_ = _146_v74
		var _nw28 *C3 = New_C3_()
		_ = _nw28
		_nw28.Ctor__(_103_v51.F2())
		_146_v74 = _nw28
		var _147_v75 _dafny.Sequence
		_ = _147_v75
		_147_v75 = _dafny.SeqOf(_146_v74, _146_v74)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_145_v73), 0))
		_ = _index17
		(_145_v73).ArraySet1((_147_v75).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_140_v68).Cardinality()), _dafny.IntOfUint32((_147_v75).Cardinality()))).Uint32()).(*C3), (_index17).Int())
		var _148_v76 _dafny.Set
		_ = _148_v76
		_148_v76 = _dafny.SetOf(_141_v69)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_145_v73), 0))
		_ = _index18
		var _nw29 *C3 = New_C3_()
		_ = _nw29
		_nw29.Ctor__((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_148_v76).Cardinality(), _46_v0)).Update(_46_v0, _103_v51.F2())).Cardinality()).Times(_103_v51.F2()))
		(_145_v73).ArraySet1(_nw29, (_index18).Int())
	} else {
		var _149___mcc_h6 _dafny.Array = _source4.Get_().(D11_DC20).Cf35
		_ = _149___mcc_h6
		var _150_cf35 _dafny.Array = _149___mcc_h6
		_ = _150_cf35
		var _151_v77 _dafny.CodePoint
		_ = _151_v77
		_151_v77 = _dafny.CodePoint('d')
		var _152_v78 _dafny.Sequence
		_ = _152_v78
		_152_v78 = _dafny.SeqOf(_151_v77)
		_151_v77 = (_152_v78).Select((Companion_Default___.SafeIndex(_103_v51.F2(), _dafny.IntOfUint32((_152_v78).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _153_v79 _dafny.Array
		_ = _153_v79
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw30
		_153_v79 = _nw30
		var _154_v80 _dafny.Sequence
		_ = _154_v80
		_154_v80 = _dafny.SeqOf(_46_v0)
		var _155_v81 _dafny.Map
		_ = _155_v81
		_155_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_154_v80)).Cardinality(), _48_v2)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_153_v79), 0))
		_ = _index19
		(_153_v79).ArraySet1(_155_v81, (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_153_v79), 0))
		_ = _index20
		(_153_v79).ArraySet1(_155_v81, (_index20).Int())
		var _156_v82 _dafny.MultiSet
		_ = _156_v82
		_156_v82 = _dafny.MultiSetOf(_48_v2)
		var _157_v83 _dafny.Map
		_ = _157_v83
		_157_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v0, (_156_v82).Cardinality())
		var _158_v84 _dafny.Map
		_ = _158_v84
		_158_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_103_v51.F2()), _46_v0)).Merge(_157_v83), _150_cf35)
		_158_v84 = _158_v84
		var _159_v85 _dafny.Map
		_ = _159_v85
		_159_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _dafny.IntOfInt64(622))
		var _160_v86 _dafny.Map
		_ = _160_v86
		_160_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D14_.Create_DC27_(true), (_dafny.Zero).Minus((_159_v85).Cardinality()))
		var _161_v87 *C2
		_ = _161_v87
		var _nw31 *C2 = New_C2_()
		_ = _nw31
		_nw31.Ctor__(((_160_v86).Merge(_160_v86)).Cardinality())
		_161_v87 = _nw31
	}
	var _162_v89 _dafny.Set
	_ = _162_v89
	_162_v89 = func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-646), _dafny.IntOfInt64(248))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _163_v88 _dafny.Int
			_163_v88 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(-646)).Cmp(_163_v88) <= 0) && ((_163_v88).Cmp(_dafny.IntOfInt64(248)) < 0) {
				_coll8.Add((_163_v88).Times(_dafny.IntOfUint32((_dafny.SeqOf(p1, p0, p1, p0)).Cardinality())))
			}
		}
		return _coll8.ToSet()
	}()
	var _pat_let_tv2 = p0
	_ = _pat_let_tv2
	_48_v2 = func(_source5 _dafny.Set) bool {
		var _164___mcc_h7 _dafny.Set = _source5
		_ = _164___mcc_h7
		var _165_cf48 _dafny.Set = _164___mcc_h7
		_ = _165_cf48
		return _pat_let_tv2
	}(_162_v89)
	var _166_v90 _dafny.Array
	_ = _166_v90
	var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
	_ = _nw32
	_166_v90 = _nw32
	r0 = _166_v90
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _167_globalState *GlobalState
	_ = _167_globalState
	var _nw33 *GlobalState = New_GlobalState_()
	_ = _nw33
	_nw33.Ctor__()
	_167_globalState = _nw33
	var _168_v0 bool
	_ = _168_v0
	_168_v0 = true
	var _169_v1 _dafny.Array
	_ = _169_v1
	var _out0 _dafny.Array
	_ = _out0
	_out0 = Companion_Default___.M0(_168_v0, (_168_v0) || (true), _167_globalState)
	_169_v1 = _out0
	_168_v0 = !(_168_v0)
	if _168_v0 {
		var _170_v2 _dafny.CodePoint
		_ = _170_v2
		_170_v2 = _dafny.CodePoint('y')
		_170_v2 = _170_v2
		var _171_v3 _dafny.Sequence
		_ = _171_v3
		_171_v3 = _dafny.SeqOf(_170_v2)
		var _172_v4 _dafny.Array
		_ = _172_v4
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_3
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_173_i0 _dafny.Int) bool {
				return !(true)
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw34 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw34).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw34).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_172_v4 = _nw34
		var _174_v5 *C5
		_ = _174_v5
		var _nw35 *C5 = New_C5_()
		_ = _nw35
		_nw35.Ctor__(_dafny.IntOfUint32((_171_v3).Cardinality()), _172_v4)
		_174_v5 = _nw35
		var _175_v6 _dafny.Sequence
		_ = _175_v6
		_175_v6 = _dafny.SeqOf((_174_v5).F0(), _dafny.IntOfUint32((_171_v3).Cardinality()))
		var _176_v7 _dafny.MultiSet
		_ = _176_v7
		_176_v7 = _dafny.MultiSetOf((_174_v5).F0())
		var _rhs10 bool = ((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_175_v6, (Companion_Default___.SafeIndex((_174_v5).F0(), _dafny.IntOfUint32((_175_v6).Cardinality()))).Uint32(), (_174_v5).F0()))).Difference(_176_v7)).IsProperSubsetOf(_176_v7)
		_ = _rhs10
		var _rhs11 bool = _168_v0
		_ = _rhs11
		var _rhs12 bool = _168_v0
		_ = _rhs12
		var _rhs13 bool = !_dafny.Companion_Sequence_.Contains(_171_v3, _170_v2)
		_ = _rhs13
		_168_v0 = _rhs10
		_168_v0 = _rhs11
		_168_v0 = _rhs12
		_168_v0 = _rhs13
		var _177_v8 D4
		_ = _177_v8
		_177_v8 = Companion_D4_.Create_DC5_(_dafny.SetOf(_168_v0))
		var _178_v9 _dafny.Int
		_ = _178_v9
		_178_v9 = _dafny.IntOfInt64(-538)
		var _179_v10 _dafny.Sequence
		_ = _179_v10
		_179_v10 = _dafny.SeqOf(_168_v0, _168_v0, _168_v0)
		var _rhs14 bool = (func() bool {
			if _168_v0 {
				return _168_v0
			}
			return (_179_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_179_v10).Cardinality()), _dafny.IntOfUint32((_179_v10).Cardinality()))).Uint32()).(bool)
		})()
		_ = _rhs14
		var _rhs15 D4 = _177_v8
		_ = _rhs15
		var _rhs16 _dafny.Int = (_dafny.Zero).Minus((_178_v9).Minus(_178_v9))
		_ = _rhs16
		_168_v0 = _rhs14
		_177_v8 = _rhs15
		_178_v9 = _rhs16
		var _180_v11 _dafny.Array
		_ = _180_v11
		var _out1 _dafny.Array
		_ = _out1
		_out1 = Companion_Default___.M0(_168_v0, Companion_Default___.Fm3((_174_v5).Fm6(_168_v0, _167_globalState), _178_v9, (_174_v5).F0(), _167_globalState), _167_globalState)
		_180_v11 = _out1
	} else {
		var _181_v12 _dafny.Int
		_ = _181_v12
		_181_v12 = _dafny.IntOfInt64(229)
		var _182_v13 _dafny.Map
		_ = _182_v13
		_182_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v12, _dafny.IntOfInt64(-100))
		var _183_v14 _dafny.MultiSet
		_ = _183_v14
		_183_v14 = _dafny.MultiSetOf(_182_v13, _182_v13, _182_v13)
		var _184_v15 _dafny.Sequence
		_ = _184_v15
		_184_v15 = _dafny.SeqOf((_183_v14).Cardinality())
		_184_v15 = _184_v15
		var _185_v16 _dafny.Sequence
		_ = _185_v16
		_185_v16 = _dafny.UnicodeSeqOfUtf8Bytes("t")
		var _186_v17 _dafny.CodePoint
		_ = _186_v17
		_186_v17 = _dafny.CodePoint('a')
		var _187_v18 *C1
		_ = _187_v18
		var _nw36 *C1 = New_C1_()
		_ = _nw36
		_nw36.Ctor__(_dafny.Companion_Sequence_.Update(_185_v16, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_184_v15).Cardinality()), _dafny.IntOfUint32((_185_v16).Cardinality()))).Uint32(), _186_v17))
		_187_v18 = _nw36
		_168_v0 = (Companion_D8_.Create_DC17_(_168_v0, _187_v18, _181_v12)).Dtor_cf30()
		var _188_v19 _dafny.Sequence
		_ = _188_v19
		_188_v19 = _dafny.SeqOf(_168_v0, _168_v0, _168_v0)
		var _189_v20 _dafny.Array
		_ = _189_v20
		var _nwElement0_7 bool = _168_v0
		_ = _nwElement0_7
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(3))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_7, 0)
		(_nw37).ArraySet1((_188_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_187_v18).F3()).Cardinality()), _dafny.IntOfUint32((_188_v19).Cardinality()))).Uint32()).(bool), 1)
		(_nw37).ArraySet1(_168_v0, 2)
		_189_v20 = _nw37
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_189_v20), 0))
		_ = _index21
		(_189_v20).ArraySet1(!(_168_v0), (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_189_v20), 0))
		_ = _index22
		(_189_v20).ArraySet1(_168_v0, (_index22).Int())
		var _190_v21 D6
		_ = _190_v21
		_190_v21 = Companion_D6_.Create_DC11_((_187_v18).F3())
		var _191_v22 _dafny.Map
		_ = _191_v22
		_191_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_190_v21, _181_v12)
		_191_v22 = (_191_v22).Update(_190_v21, Companion_Default___.SafeModuloInt(_181_v12, (_dafny.Zero).Minus(_181_v12)))
		var _192_v23 *C3
		_ = _192_v23
		var _nw38 *C3 = New_C3_()
		_ = _nw38
		_nw38.Ctor__(_181_v12)
		_192_v23 = _nw38
		var _193_v24 _dafny.Sequence
		_ = _193_v24
		_193_v24 = _dafny.SeqOf(_192_v23, _192_v23)
		var _194_v25 _dafny.Array
		_ = _194_v25
		var _nwElement0_8 *C3 = _192_v23
		_ = _nwElement0_8
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_8, 0)
		(_nw39).ArraySet1(_192_v23, 1)
		(_nw39).ArraySet1(_192_v23, 2)
		(_nw39).ArraySet1(_192_v23, 3)
		(_nw39).ArraySet1(_192_v23, 4)
		(_nw39).ArraySet1(_192_v23, 5)
		(_nw39).ArraySet1((_193_v24).Select((Companion_Default___.SafeIndex(_181_v12, _dafny.IntOfUint32((_193_v24).Cardinality()))).Uint32()).(*C3), 6)
		(_nw39).ArraySet1((func() *C3 {
			if _168_v0 {
				return _192_v23
			}
			return _192_v23
		})(), 7)
		(_nw39).ArraySet1((func() *C3 {
			if true {
				return _192_v23
			}
			return _192_v23
		})(), 8)
		(_nw39).ArraySet1(_192_v23, 9)
		(_nw39).ArraySet1(_192_v23, 10)
		(_nw39).ArraySet1(_192_v23, 11)
		(_nw39).ArraySet1(_192_v23, 12)
		(_nw39).ArraySet1(_192_v23, 13)
		(_nw39).ArraySet1(_192_v23, 14)
		(_nw39).ArraySet1(_192_v23, 15)
		(_nw39).ArraySet1(_192_v23, 16)
		(_nw39).ArraySet1(_192_v23, 17)
		(_nw39).ArraySet1(_192_v23, 18)
		(_nw39).ArraySet1(_192_v23, 19)
		(_nw39).ArraySet1(_192_v23, 20)
		(_nw39).ArraySet1((func() *C3 {
			if (_192_v23).Fm8(_185_v16, (_189_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_189_v20), 0))).Int()).(bool), _167_globalState) {
				return _192_v23
			}
			return _192_v23
		})(), 21)
		(_nw39).ArraySet1(_192_v23, 22)
		(_nw39).ArraySet1(_192_v23, 23)
		(_nw39).ArraySet1(_192_v23, 24)
		_194_v25 = _nw39
		_194_v25 = _194_v25
	}
	var _195_v27 _dafny.Set
	_ = _195_v27
	_195_v27 = func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(691), _dafny.IntOfInt64(586))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _196_v26 _dafny.Int
			_196_v26 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(691)).Cmp(_196_v26) <= 0) && ((_196_v26).Cmp(_dafny.IntOfInt64(586)) < 0) {
				_coll9.Add((_196_v26).Plus(_dafny.IntOfInt64(462)))
			}
		}
		return _coll9.ToSet()
	}()
	var _pat_let_tv3 = _168_v0
	_ = _pat_let_tv3
	if func(_source6 _dafny.Set) bool {
		var _197___mcc_h0 _dafny.Set = _source6
		_ = _197___mcc_h0
		var _198_cf48 _dafny.Set = _197___mcc_h0
		_ = _198_cf48
		return _pat_let_tv3
	}(_195_v27) {
		var _199_v28 _dafny.Int
		_ = _199_v28
		_199_v28 = _dafny.IntOfInt64(269)
		var _200_v29 _dafny.Sequence
		_ = _200_v29
		_200_v29 = _dafny.SeqOf((_199_v28).Times(_dafny.IntOfInt64(-921)))
		_200_v29 = _200_v29
		var _201_v30 _dafny.Map
		_ = _201_v30
		_201_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, !(false))
		var _202_v31 _dafny.MultiSet
		_ = _202_v31
		_202_v31 = _dafny.MultiSetOf((func() bool {
			if (_201_v30).Contains(_168_v0) {
				return (_201_v30).Get(_168_v0).(bool)
			}
			return _168_v0
		})())
		if (_202_v31).IsProperSubsetOf((_202_v31).Difference(_202_v31)) {
			var _203_v32 _dafny.Array
			_ = _203_v32
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw40
			_203_v32 = _nw40
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))
			_ = _index23
			(_203_v32).ArraySet1(_169_v1, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))
			_ = _index24
			(_203_v32).ArraySet1(_169_v1, (_index24).Int())
			_168_v0 = _168_v0
			var _204_v33 _dafny.Array
			_ = _204_v33
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw41
			_204_v33 = _nw41
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_204_v33), 0))
			_ = _index25
			(_204_v33).ArraySet1(Companion_Default___.Fm5(_167_globalState), (_index25).Int())
			var _205_v34 _dafny.Sequence
			_ = _205_v34
			_205_v34 = _dafny.SeqOf(Companion_Default___.Fm5(_167_globalState), _168_v0, _168_v0)
			var _206_v35 _dafny.Sequence
			_ = _206_v35
			_206_v35 = _dafny.UnicodeSeqOfUtf8Bytes("qgfsf")
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_204_v33), 0))
			_ = _index26
			var _rhs17 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_205_v34).Cardinality()), _dafny.IntOfUint32((_206_v35).Cardinality()))).Cmp((_dafny.Zero).Minus(((_202_v31).Intersection(_202_v31)).Cardinality())) == 0
			_ = _rhs17
			var _rhs18 bool = _168_v0
			_ = _rhs18
			var _lhs3 _dafny.Array = _204_v33
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_204_v33), 0))
			_ = _lhs4
			_168_v0 = _rhs17
			(_lhs3).ArraySet1(_rhs18, (_lhs4).Int())
			var _207_v36 _dafny.Map
			_ = _207_v36
			_207_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v28, _199_v28)
			var _208_v37 _dafny.Sequence
			_ = _208_v37
			_208_v37 = _dafny.SeqOf(_207_v36)
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_203_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))).Int()))
			_ = _arr0
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_dafny.ArrayCastTo((_203_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))).Int()))), 0))
			_ = _index27
			(_arr0).ArraySet1((_dafny.IntOfUint32((_200_v29).Cardinality())).Times(_dafny.IntOfUint32((_208_v37).Cardinality())), (_index27).Int())
			var _209_v38 *C6
			_ = _209_v38
			var _nw42 *C6 = New_C6_()
			_ = _nw42
			_nw42.Ctor__()
			_209_v38 = _nw42
			var _210_v39 _dafny.Sequence
			_ = _210_v39
			_210_v39 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_209_v38, _199_v28))
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_203_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))).Int()))
			_ = _arr1
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_dafny.ArrayCastTo((_203_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_203_v32), 0))).Int()))), 0))
			_ = _index28
			(_arr1).ArraySet1(_dafny.IntOfUint32((_210_v39).Cardinality()), (_index28).Int())
			_168_v0 = (_204_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_204_v33), 0))).Int()).(bool)
		} else {
			var _211_v40 _dafny.Sequence
			_ = _211_v40
			_211_v40 = _dafny.UnicodeSeqOfUtf8Bytes("cpimjma")
			_211_v40 = _211_v40
			var _212_v41 *C1
			_ = _212_v41
			var _nw43 *C1 = New_C1_()
			_ = _nw43
			_nw43.Ctor__(_211_v40)
			_212_v41 = _nw43
			_211_v40 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_213_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))
			var _214_v42 _dafny.Array
			_ = _214_v42
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_4
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = func(_215_i2 _dafny.Int) bool {
					return !(false)
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw44 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw44).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw44).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_214_v42 = _nw44
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_214_v42), 0))
			_ = _index29
			(_214_v42).ArraySet1(_168_v0, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_214_v42), 0))
			_ = _index30
			(_214_v42).ArraySet1((_199_v28).Cmp(_199_v28) < 0, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_169_v1), 0))
			_ = _index31
			(_169_v1).ArraySet1(_199_v28, (_index31).Int())
			var _216_v43 _dafny.Array
			_ = _216_v43
			var _nw45 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
			_ = _nw45
			_216_v43 = _nw45
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_169_v1), 0))
			_ = _index32
			var _rhs19 _dafny.Int = _199_v28
			_ = _rhs19
			var _rhs20 _dafny.Int = (_199_v28).Minus(Companion_Default___.SafeDivisionInt(_199_v28, (_212_v41).Fm1(_168_v0, _167_globalState)))
			_ = _rhs20
			var _rhs21 _dafny.Array = _216_v43
			_ = _rhs21
			var _lhs5 _dafny.Array = _169_v1
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_169_v1), 0))
			_ = _lhs6
			_199_v28 = _rhs19
			(_lhs5).ArraySet1(_rhs20, (_lhs6).Int())
			_216_v43 = _rhs21
		}
		var _217_v44 _dafny.Array
		_ = _217_v44
		var _out2 _dafny.Array
		_ = _out2
		_out2 = Companion_Default___.M0((_202_v31).IsDisjointFrom(_202_v31), _168_v0, _167_globalState)
		_217_v44 = _out2
		var _218_v45 _dafny.Array
		_ = _218_v45
		var _out3 _dafny.Array
		_ = _out3
		_out3 = Companion_Default___.M0(_168_v0, _168_v0, _167_globalState)
		_218_v45 = _out3
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_218_v45), 0))
		_ = _index33
		(_218_v45).ArraySet1(_dafny.IntOfInt64(-497), (_index33).Int())
		var _219_v46 *C6
		_ = _219_v46
		var _nw46 *C6 = New_C6_()
		_ = _nw46
		_nw46.Ctor__()
		_219_v46 = _nw46
		var _220_v47 _dafny.Map
		_ = _220_v47
		_220_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v46, _199_v28)
		var _221_v48 _dafny.Map
		_ = _221_v48
		_221_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v28, _169_v1)
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_218_v45), 0))
		_ = _index34
		(_218_v45).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_220_v47).Contains(_219_v46) {
				return (_220_v47).Get(_219_v46).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-686)
		})(), (_221_v48).Cardinality()), (_index34).Int())
	} else {
		var _222_v49 _dafny.Sequence
		_ = _222_v49
		_222_v49 = _dafny.UnicodeSeqOfUtf8Bytes("htengavdr")
		var _223_v50 _dafny.Int
		_ = _223_v50
		_223_v50 = _dafny.IntOfInt64(-705)
		var _224_v51 _dafny.CodePoint
		_ = _224_v51
		_224_v51 = _dafny.CodePoint('q')
		_222_v49 = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wafife"), (Companion_Default___.SafeIndex(_223_v50, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wafife")).Cardinality()))).Uint32(), _224_v51)
		var _225_v52 _dafny.Sequence
		_ = _225_v52
		_225_v52 = _dafny.SeqOf(_223_v50)
		var _226_v53 _dafny.Set
		_ = _226_v53
		_226_v53 = _dafny.SetOf(_168_v0)
		var _227_v54 _dafny.Map
		_ = _227_v54
		_227_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_225_v52).Select((Companion_Default___.SafeIndex((_226_v53).Cardinality(), _dafny.IntOfUint32((_225_v52).Cardinality()))).Uint32()).(_dafny.Int), _169_v1)
		var _228_v55 _dafny.Array
		_ = _228_v55
		var _nwElement0_9 _dafny.Array = (func() _dafny.Array {
			if (_227_v54).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_223_v50, _223_v50)).Cardinality())) {
				return (_227_v54).Get(_dafny.IntOfUint32((_dafny.SeqOf(_223_v50, _223_v50)).Cardinality())).(_dafny.Array)
			}
			return _169_v1
		})()
		_ = _nwElement0_9
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_9, 0)
		(_nw47).ArraySet1(_169_v1, 1)
		(_nw47).ArraySet1(_169_v1, 2)
		(_nw47).ArraySet1(_169_v1, 3)
		(_nw47).ArraySet1(_169_v1, 4)
		(_nw47).ArraySet1((func() _dafny.Array {
			if _168_v0 {
				return _169_v1
			}
			return _169_v1
		})(), 5)
		(_nw47).ArraySet1(_169_v1, 6)
		(_nw47).ArraySet1(_169_v1, 7)
		(_nw47).ArraySet1(_169_v1, 8)
		(_nw47).ArraySet1(_169_v1, 9)
		_228_v55 = _nw47
		var _229_v56 _dafny.Sequence
		_ = _229_v56
		_229_v56 = _dafny.SeqOf(_228_v55, (func() _dafny.Array {
			if _168_v0 {
				return _228_v55
			}
			return _228_v55
		})())
		_228_v55 = (_229_v56).Select((Companion_Default___.SafeIndex(_223_v50, _dafny.IntOfUint32((_229_v56).Cardinality()))).Uint32()).(_dafny.Array)
		var _230_v57 _dafny.MultiSet
		_ = _230_v57
		_230_v57 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(464))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_231_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_232_i3 _dafny.Int) _dafny.CodePoint {
				return _231_v51
			}
		})(_224_v51))), _222_v49)
		var _233_v58 *C0
		_ = _233_v58
		var _nw48 *C0 = New_C0_()
		_ = _nw48
		_nw48.Ctor__(!(_dafny.MultiSetOf(_222_v49, _222_v49)).Equals(_230_v57), _168_v0)
		_233_v58 = _nw48
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_169_v1), 0))
		_ = _index35
		(_169_v1).ArraySet1(_223_v50, (_index35).Int())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_169_v1), 0))
		_ = _index36
		(_169_v1).ArraySet1(_223_v50, (_index36).Int())
		var _234_v59 *C7
		_ = _234_v59
		var _nw49 *C7 = New_C7_()
		_ = _nw49
		_nw49.Ctor__()
		_234_v59 = _nw49
	}
	if _168_v0 {
		var _235_v60 _dafny.Int
		_ = _235_v60
		_235_v60 = _dafny.IntOfInt64(823)
		var _236_v61 _dafny.Sequence
		_ = _236_v61
		_236_v61 = _dafny.SeqOf(_235_v60)
		var _237_v62 _dafny.Sequence
		_ = _237_v62
		_237_v62 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_236_v61, (Companion_Default___.SafeIndex(_235_v60, _dafny.IntOfUint32((_236_v61).Cardinality()))).Uint32(), _235_v60))
		var _238_v63 _dafny.Map
		_ = _238_v63
		_238_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _dafny.IntOfUint32((_236_v61).Cardinality()))
		var _239_v64 _dafny.Sequence
		_ = _239_v64
		_239_v64 = _dafny.UnicodeSeqOfUtf8Bytes("tjwbxhkeo")
		var _240_v65 _dafny.Set
		_ = _240_v65
		_240_v65 = _dafny.SetOf(_235_v60, _dafny.IntOfUint32((_237_v62).Cardinality()), (func() _dafny.Int {
			if (_238_v63).Contains(_168_v0) {
				return (_238_v63).Get(_168_v0).(_dafny.Int)
			}
			return _235_v60
		})(), _235_v60, (func() _dafny.Int {
			if _168_v0 {
				return _235_v60
			}
			return _dafny.IntOfUint32((_239_v64).Cardinality())
		})())
		_240_v65 = _240_v65
		var _241_v66 _dafny.Array
		_ = _241_v66
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_5
		var _nw50 _dafny.Array
		_ = _nw50
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw50 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_242_v0 bool) func(_dafny.Int) bool {
				return func(_243_i4 _dafny.Int) bool {
					return _242_v0
				}
			})(_168_v0)
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
		_241_v66 = _nw50
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))
		_ = _index37
		(_241_v66).ArraySet1(_168_v0, (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_241_v66), 0))
		_ = _index38
		(_241_v66).ArraySet1(_168_v0, (_index38).Int())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))
		_ = _index39
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_241_v66), 0))
		_ = _index40
		var _rhs22 bool = (_168_v0) && (_168_v0)
		_ = _rhs22
		var _rhs23 bool = _168_v0
		_ = _rhs23
		var _lhs7 _dafny.Array = _241_v66
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))
		_ = _lhs8
		var _lhs9 _dafny.Array = _241_v66
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_241_v66), 0))
		_ = _lhs10
		(_lhs7).ArraySet1(_rhs22, (_lhs8).Int())
		(_lhs9).ArraySet1(_rhs23, (_lhs10).Int())
		var _244_v67 _dafny.Array
		_ = _244_v67
		var _out4 _dafny.Array
		_ = _out4
		_out4 = Companion_Default___.M0(_168_v0, (_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool), _167_globalState)
		_244_v67 = _out4
		var _245_v68 *C4
		_ = _245_v68
		var _nw51 *C4 = New_C4_()
		_ = _nw51
		_nw51.Ctor__(_235_v60)
		_245_v68 = _nw51
		var _246_v69 _dafny.Map
		_ = _246_v69
		_246_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, (_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool))
		var _247_v70 _dafny.Sequence
		_ = _247_v70
		_247_v70 = _dafny.SeqOf(_246_v69)
		var _248_v71 _dafny.Map
		_ = _248_v71
		_248_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool), _246_v69)
		var _249_v72 _dafny.Array
		_ = _249_v72
		var _nwElement0_10 _dafny.Map = (_246_v69).Merge(_246_v69)
		_ = _nwElement0_10
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(26))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_10, 0)
		(_nw52).ArraySet1((_246_v69).Merge(_246_v69), 1)
		(_nw52).ArraySet1(_246_v69, 2)
		(_nw52).ArraySet1(_246_v69, 3)
		(_nw52).ArraySet1(_246_v69, 4)
		(_nw52).ArraySet1((_246_v69).Merge(_246_v69), 5)
		(_nw52).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, (_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool))).Merge(_246_v69), 6)
		(_nw52).ArraySet1(_246_v69, 7)
		(_nw52).ArraySet1((_246_v69).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, (_245_v68).Fm8(_239_v64, (_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool), _167_globalState))), 8)
		(_nw52).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, _168_v0), 9)
		(_nw52).ArraySet1((_247_v70).Select((Companion_Default___.SafeIndex(_235_v60, _dafny.IntOfUint32((_247_v70).Cardinality()))).Uint32()).(_dafny.Map), 10)
		(_nw52).ArraySet1(_246_v69, 11)
		(_nw52).ArraySet1(_246_v69, 12)
		(_nw52).ArraySet1(_246_v69, 13)
		(_nw52).ArraySet1((func() _dafny.Map {
			if (_248_v71).Contains(_168_v0) {
				return (_248_v71).Get(_168_v0).(_dafny.Map)
			}
			return _246_v69
		})(), 14)
		(_nw52).ArraySet1((_246_v69).Merge(_246_v69), 15)
		(_nw52).ArraySet1(_246_v69, 16)
		(_nw52).ArraySet1((Companion_D18_.Create_DC33_(_246_v69)).Dtor_cf50(), 17)
		(_nw52).ArraySet1((_246_v69).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, _168_v0)).Update(_245_v68, _168_v0)), 18)
		(_nw52).ArraySet1(_246_v69, 19)
		(_nw52).ArraySet1(_246_v69, 20)
		(_nw52).ArraySet1((_246_v69).Merge(_246_v69), 21)
		(_nw52).ArraySet1(_246_v69, 22)
		(_nw52).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v68, _168_v0), 23)
		(_nw52).ArraySet1(_246_v69, 24)
		(_nw52).ArraySet1(((_246_v69).Update(_245_v68, false)).Merge(_246_v69), 25)
		_249_v72 = _nw52
		_249_v72 = _249_v72
		if (_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool) {
			var _250_v73 _dafny.Array
			_ = _250_v73
			var _nw53 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
			_ = _nw53
			_250_v73 = _nw53
			var _251_v74 _dafny.MultiSet
			_ = _251_v74
			_251_v74 = _dafny.MultiSetOf(_235_v60, _235_v60)
			var _252_v75 _dafny.MultiSet
			_ = _252_v75
			_252_v75 = _dafny.MultiSetOf(_251_v74)
			var _253_v76 _dafny.MultiSet
			_ = _253_v76
			_253_v76 = _dafny.MultiSetOf((_252_v75).Cardinality(), _dafny.IntOfUint32((_239_v64).Cardinality()))
			var _254_v77 T1
			_ = _254_v77
			var _nw54 *C4 = New_C4_()
			_ = _nw54
			_nw54.Ctor__(_235_v60)
			_254_v77 = _nw54
			var _255_v78 _dafny.MultiSet
			_ = _255_v78
			_255_v78 = _dafny.MultiSetOf(_254_v77)
			var _256_v79 T0
			_ = _256_v79
			var _nw55 *C5 = New_C5_()
			_ = _nw55
			_nw55.Ctor__((func() _dafny.Int {
				if (_251_v74).Contains((_245_v68).Fm1((_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool), _167_globalState)) {
					return (_251_v74).Multiplicity((_245_v68).Fm1((_241_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))).Int()).(bool), _167_globalState))
				}
				return (_255_v78).Cardinality()
			})(), _241_v66)
			_256_v79 = _nw55
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_250_v73), 0))
			_ = _index41
			(_250_v73).ArraySet1((_256_v79), (_index41).Int())
			var _257_v80 _dafny.Map
			_ = _257_v80
			_257_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v60, _254_v77)
			var _258_v81 _dafny.Map
			_ = _258_v81
			_258_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v60, _257_v80)
			var _259_v82 _dafny.Set
			_ = _259_v82
			_259_v82 = _dafny.SetOf((_257_v80).Merge((func() _dafny.Map {
				if (_258_v81).Contains(_235_v60) {
					return (_258_v81).Get(_235_v60).(_dafny.Map)
				}
				return _257_v80
			})()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v60, _254_v77), (_257_v80).Update(_235_v60, _254_v77), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_235_v60), _254_v77), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v60, _254_v77))
			var _260_v83 _dafny.Sequence
			_ = _260_v83
			_260_v83 = _dafny.SeqOf(_256_v79)
			var _261_v84 _dafny.Map
			_ = _261_v84
			_261_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v61, (_260_v83).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-19)), _dafny.IntOfUint32((_260_v83).Cardinality()))).Uint32()).(T0))
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_250_v73), 0))
			_ = _index42
			var _rhs24 T0 = (func() T0 {
				if (_261_v84).Contains(_236_v61) {
					return (_261_v84).Get(_236_v61).(T0)
				}
				return _256_v79
			})()
			_ = _rhs24
			var _rhs25 _dafny.Set = _259_v82
			_ = _rhs25
			var _lhs11 _dafny.Array = _250_v73
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_250_v73), 0))
			_ = _lhs12
			(_lhs11).ArraySet1(_rhs24, (_lhs12).Int())
			_259_v82 = _rhs25
			var _262_v85 *C1
			_ = _262_v85
			var _nw56 *C1 = New_C1_()
			_ = _nw56
			_nw56.Ctor__(_239_v64)
			_262_v85 = _nw56
			var _263_v86 _dafny.Set
			_ = _263_v86
			_263_v86 = _dafny.SetOf(_262_v85, _262_v85)
			var _264_v87 _dafny.Set
			_ = _264_v87
			_264_v87 = _dafny.SetOf((_263_v86).Intersection(_263_v86))
			_264_v87 = (_264_v87).Difference(_264_v87)
			var _265_v88 _dafny.Array
			_ = _265_v88
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw57
			_265_v88 = _nw57
			var _266_v89 _dafny.Array
			_ = _266_v89
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw58
			_266_v89 = _nw58
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_265_v88), 0))
			_ = _index43
			(_265_v88).ArraySet1(_266_v89, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_265_v88), 0))
			_ = _index44
			(_265_v88).ArraySet1(_266_v89, (_index44).Int())
			_168_v0 = _168_v0
			var _267_v90 _dafny.Array
			_ = _267_v90
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw59
			_267_v90 = _nw59
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_267_v90), 0))
			_ = _index45
			(_267_v90).ArraySet1(_241_v66, (_index45).Int())
			var _268_v91 _dafny.MultiSet
			_ = _268_v91
			_268_v91 = _dafny.MultiSetOf(_169_v1)
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))
			_ = _index46
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_267_v90), 0))
			_ = _index47
			var _rhs26 bool = ((_268_v91).Union(_268_v91)).IsSubsetOf(_268_v91)
			_ = _rhs26
			var _rhs27 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uookla"), (_262_v85).F3())
			_ = _rhs27
			var _rhs28 _dafny.Int = _235_v60
			_ = _rhs28
			var _rhs29 _dafny.Array = _241_v66
			_ = _rhs29
			var _lhs13 _dafny.Array = _241_v66
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_241_v66), 0))
			_ = _lhs14
			var _lhs15 _dafny.Array = _267_v90
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_267_v90), 0))
			_ = _lhs16
			(_lhs13).ArraySet1(_rhs26, (_lhs14).Int())
			_239_v64 = _rhs27
			_235_v60 = _rhs28
			(_lhs15).ArraySet1(_rhs29, (_lhs16).Int())
		} else {
			var _269_v92 _dafny.MultiSet
			_ = _269_v92
			_269_v92 = _dafny.MultiSetOf(_235_v60, _235_v60)
			_235_v60 = ((_235_v60).Minus((_269_v92).Cardinality())).Times(_235_v60)
			var _270_v93 _dafny.CodePoint
			_ = _270_v93
			_270_v93 = _dafny.CodePoint('y')
			_270_v93 = _270_v93
			_270_v93 = _270_v93
			_235_v60 = _235_v60
			_235_v60 = _235_v60
		}
	} else {
		var _271_v94 _dafny.Sequence
		_ = _271_v94
		_271_v94 = _dafny.UnicodeSeqOfUtf8Bytes("dlkwbx")
		_271_v94 = _dafny.Companion_Sequence_.Concatenate(_271_v94, _dafny.Companion_Sequence_.Concatenate(_271_v94, _271_v94))
		var _272_v95 _dafny.Int
		_ = _272_v95
		_272_v95 = _dafny.IntOfInt64(785)
		var _273_v96 _dafny.CodePoint
		_ = _273_v96
		_273_v96 = _dafny.CodePoint('y')
		var _274_v97 _dafny.Map
		_ = _274_v97
		_274_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_271_v94, (Companion_Default___.SafeIndex(_272_v95, _dafny.IntOfUint32((_271_v94).Cardinality()))).Uint32(), _273_v96), _271_v94))
		var _275_v98 _dafny.MultiSet
		_ = _275_v98
		_275_v98 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("rllwbgim"))
		_274_v97 = (_274_v97).Update(true, (_275_v98).Intersection(_275_v98))
		var _276_v99 _dafny.Map
		_ = _276_v99
		_276_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _168_v0)
		var _277_v100 _dafny.Map
		_ = _277_v100
		_277_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _dafny.SeqOf((func() bool {
			if (_276_v99).Contains(false) {
				return (_276_v99).Get(false).(bool)
			}
			return !(false)
		})()))
		var _278_v101 _dafny.Sequence
		_ = _278_v101
		_278_v101 = _dafny.SeqOf(_168_v0, _168_v0)
		_277_v100 = (_277_v100).Update(_168_v0, _dafny.Companion_Sequence_.Concatenate(_278_v101, Companion_Default___.Fm16(_167_globalState)))
		var _279_v102 D6
		_ = _279_v102
		_279_v102 = Companion_D6_.Create_DC11_(_271_v94)
		var _280_v103 _dafny.Map
		_ = _280_v103
		_280_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v102, _272_v95)
		var _281_v104 _dafny.Set
		_ = _281_v104
		_281_v104 = _dafny.SetOf(_280_v103)
		var _282_v105 _dafny.Set
		_ = _282_v105
		_282_v105 = _281_v104
		var _283_v106 _dafny.Array
		_ = _283_v106
		var _nw60 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
		_ = _nw60
		_283_v106 = _nw60
		var _284_v107 _dafny.Map
		_ = _284_v107
		_284_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_v105, _283_v106)
		_284_v107 = (_284_v107).Update(_282_v105, _283_v106)
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_169_v1), 0))
		_ = _index48
		(_169_v1).ArraySet1(_272_v95, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_169_v1), 0))
		_ = _index49
		(_169_v1).ArraySet1(Companion_Default___.SafeModuloInt(_272_v95, (_dafny.Zero).Minus(_272_v95)), (_index49).Int())
	}
	_169_v1 = _169_v1
	if _168_v0 {
		var _285_v108 _dafny.Sequence
		_ = _285_v108
		_285_v108 = _dafny.UnicodeSeqOfUtf8Bytes("roxp")
		var _286_v109 D6
		_ = _286_v109
		_286_v109 = Companion_D6_.Create_DC11_(_285_v108)
		var _pat_let_tv4 = _285_v108
		_ = _pat_let_tv4
		var _287_v110 _dafny.Set
		_ = _287_v110
		_287_v110 = _dafny.SetOf(func(_pat_let4_0 D6) D6 {
			return func(_288_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let5_0 _dafny.Sequence) D6 {
					return func(_289_dt__update_hcf21_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC11_(_289_dt__update_hcf21_h0)
					}(_pat_let5_0)
				}(_pat_let_tv4)
			}(_pat_let4_0)
		}(_286_v109), _286_v109)
		var _290_v113 _dafny.Int
		_ = _290_v113
		_290_v113 = _dafny.IntOfInt64(-20)
		var _291_v114 _dafny.Sequence
		_ = _291_v114
		_291_v114 = _dafny.SeqOf(Companion_Default___.Fm5(_167_globalState))
		var _292_v115 _dafny.MultiSet
		_ = _292_v115
		_292_v115 = _dafny.MultiSetOf(_286_v109, Companion_D6_.Create_DC11_(_285_v108))
		var _293_v116 D20
		_ = _293_v116
		_293_v116 = Companion_D20_.Create_DC36_(_292_v115)
		var _294_v119 _dafny.Map
		_ = _294_v119
		_294_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v109, _168_v0)
		var _295_v121 _dafny.Array
		_ = _295_v121
		var _nwElement0_11 _dafny.Set = _287_v110
		_ = _nwElement0_11
		var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(14))
		_ = _nw61
		(_nw61).ArraySet1(_nwElement0_11, 0)
		(_nw61).ArraySet1(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_287_v110).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _296_v111 D6
					_296_v111 = interface{}(_compr_11).(D6)
					if (_287_v110).Contains(_296_v111) {
						_coll11.Add(_296_v111)
					}
				}
				return _coll11.ToSet()
			}()).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _297_v112 D6
				_297_v112 = interface{}(_compr_10).(D6)
				if (func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_287_v110).Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _298_v111 D6
						_298_v111 = interface{}(_compr_12).(D6)
						if (_287_v110).Contains(_298_v111) {
							_coll12.Add(_298_v111)
						}
					}
					return _coll12.ToSet()
				}()).Contains(_297_v112) {
					_coll10.Add(_297_v112)
				}
			}
			return _coll10.ToSet()
		}(), 1)
		(_nw61).ArraySet1(_dafny.SetOf(_286_v109, _286_v109, _286_v109, _286_v109, _286_v109), 2)
		(_nw61).ArraySet1(_dafny.SetOf(_286_v109, _286_v109), 3)
		(_nw61).ArraySet1(_287_v110, 4)
		(_nw61).ArraySet1(_287_v110, 5)
		(_nw61).ArraySet1(_287_v110, 6)
		(_nw61).ArraySet1(Companion_Default___.Fm52(_168_v0, _290_v113, _dafny.IntOfUint32((Companion_Default___.Fm16(_167_globalState)).Cardinality()), (_291_v114).Select((Companion_Default___.SafeIndex(_290_v113, _dafny.IntOfUint32((_291_v114).Cardinality()))).Uint32()).(bool), _167_globalState), 7)
		(_nw61).ArraySet1(_287_v110, 8)
		(_nw61).ArraySet1(_287_v110, 9)
		(_nw61).ArraySet1(_287_v110, 10)
		(_nw61).ArraySet1((func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate(((_293_v116).Dtor_cf54()).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _299_v117 D6
				_299_v117 = interface{}(_compr_13).(D6)
				if ((_293_v116).Dtor_cf54()).Contains(_299_v117) {
					_coll13.Add(_299_v117)
				}
			}
			return _coll13.ToSet()
		}()).Union(func() _dafny.Set {
			var _coll14 = _dafny.NewBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v109, _290_v113)).Keys().Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _300_v118 D6
				_300_v118 = interface{}(_compr_14).(D6)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v109, _290_v113)).Contains(_300_v118) {
					_coll14.Add(_300_v118)
				}
			}
			return _coll14.ToSet()
		}()), 11)
		(_nw61).ArraySet1(_287_v110, 12)
		(_nw61).ArraySet1(func() _dafny.Set {
			var _coll15 = _dafny.NewBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate((_294_v119).Keys().Elements()); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _301_v120 D6
				_301_v120 = interface{}(_compr_15).(D6)
				if (_294_v119).Contains(_301_v120) {
					_coll15.Add(_301_v120)
				}
			}
			return _coll15.ToSet()
		}(), 13)
		_295_v121 = _nw61
		var _302_v122 _dafny.Set
		_ = _302_v122
		_302_v122 = _dafny.SetOf(!(_168_v0), _168_v0)
		_295_v121 = (func() _dafny.Array {
			if (_302_v122).IsDisjointFrom(_302_v122) {
				return _295_v121
			}
			return _295_v121
		})()
		var _rhs30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_285_v108, _285_v108)
		_ = _rhs30
		var _rhs31 bool = !_dafny.Companion_Sequence_.Equal(_285_v108, _285_v108)
		_ = _rhs31
		var _rhs32 bool = !(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(492), _dafny.IntOfInt64(-52))); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _303_v123 _dafny.Int
				_303_v123 = interface{}(_compr_16).(_dafny.Int)
				if ((_dafny.IntOfInt64(492)).Cmp(_303_v123) <= 0) && ((_303_v123).Cmp(_dafny.IntOfInt64(-52)) < 0) {
					_coll16.Add((_303_v123).Plus(_290_v113), _290_v113)
				}
			}
			return _coll16.ToMap()
		}()).Contains(_290_v113)
		_ = _rhs32
		var _rhs33 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qgybp"), _dafny.UnicodeSeqOfUtf8Bytes("ubiaagdx")), _dafny.Companion_Sequence_.Concatenate(_285_v108, _285_v108))
		_ = _rhs33
		_285_v108 = _rhs30
		_168_v0 = _rhs31
		_168_v0 = _rhs32
		_285_v108 = _rhs33
		_168_v0 = _168_v0
		var _304_v124 _dafny.Array
		_ = _304_v124
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw62
		_304_v124 = _nw62
		var _305_v125 _dafny.Map
		_ = _305_v125
		_305_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v113, _304_v124)
		var _306_v126 _dafny.Array
		_ = _306_v126
		var _nwElement0_12 _dafny.Array = _304_v124
		_ = _nwElement0_12
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(8))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_12, 0)
		(_nw63).ArraySet1(_304_v124, 1)
		(_nw63).ArraySet1(_304_v124, 2)
		(_nw63).ArraySet1(_304_v124, 3)
		(_nw63).ArraySet1(_304_v124, 4)
		(_nw63).ArraySet1((func() _dafny.Array {
			if _168_v0 {
				return _304_v124
			}
			return _304_v124
		})(), 5)
		(_nw63).ArraySet1(_304_v124, 6)
		(_nw63).ArraySet1((func() _dafny.Array {
			if (_305_v125).Contains(_290_v113) {
				return (_305_v125).Get(_290_v113).(_dafny.Array)
			}
			return _304_v124
		})(), 7)
		_306_v126 = _nw63
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_306_v126), 0))
		_ = _index50
		(_306_v126).ArraySet1(_304_v124, (_index50).Int())
		var _307_v127 D0
		_ = _307_v127
		_307_v127 = Companion_D0_.Create_DC0_(_168_v0)
		var _308_v128 D14
		_ = _308_v128
		_308_v128 = Companion_D14_.Create_DC26_(_307_v127, _302_v122)
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_306_v126), 0))
		_ = _index51
		var _rhs34 _dafny.Sequence = _285_v108
		_ = _rhs34
		var _rhs35 _dafny.Array = _304_v124
		_ = _rhs35
		var _rhs36 D14 = Companion_Default___.Fm53(_168_v0, _168_v0, _167_globalState)
		_ = _rhs36
		var _lhs17 _dafny.Array = _306_v126
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_306_v126), 0))
		_ = _lhs18
		_285_v108 = _rhs34
		(_lhs17).ArraySet1(_rhs35, (_lhs18).Int())
		_308_v128 = _rhs36
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_169_v1), 0))
		_ = _index52
		(_169_v1).ArraySet1(_290_v113, (_index52).Int())
		var _309_v129 _dafny.MultiSet
		_ = _309_v129
		_309_v129 = _dafny.MultiSetOf(_169_v1, _169_v1)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_169_v1), 0))
		_ = _index53
		var _rhs37 _dafny.Int = (_309_v129).Cardinality()
		_ = _rhs37
		var _rhs38 _dafny.Int = Companion_Default___.Fm54(_168_v0, _dafny.Companion_Sequence_.Concatenate(_285_v108, _dafny.UnicodeSeqOfUtf8Bytes("ef")), _167_globalState)
		_ = _rhs38
		var _rhs39 _dafny.Sequence = _285_v108
		_ = _rhs39
		var _rhs40 _dafny.Array = _169_v1
		_ = _rhs40
		var _lhs19 _dafny.Array = _169_v1
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_169_v1), 0))
		_ = _lhs20
		(_lhs19).ArraySet1(_rhs37, (_lhs20).Int())
		_290_v113 = _rhs38
		_285_v108 = _rhs39
		_169_v1 = _rhs40
	} else {
		_168_v0 = false
		var _310_v130 _dafny.Int
		_ = _310_v130
		_310_v130 = _dafny.IntOfInt64(627)
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_169_v1), 0))
		_ = _index54
		(_169_v1).ArraySet1(_310_v130, (_index54).Int())
		var _311_v131 _dafny.Sequence
		_ = _311_v131
		_311_v131 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()))
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_169_v1), 0))
		_ = _index55
		(_169_v1).ArraySet1(((_310_v130).Minus(_dafny.IntOfUint32((_311_v131).Cardinality()))).Plus(_dafny.IntOfInt64(-306)), (_index55).Int())
		_168_v0 = _168_v0
		_168_v0 = _168_v0
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_169_v1), 0))
		_ = _index56
		(_169_v1).ArraySet1(Companion_Default___.SafeDivisionInt(_310_v130, _310_v130), (_index56).Int())
	}
	var _312_v132 _dafny.Int
	_ = _312_v132
	_312_v132 = _dafny.IntOfInt64(786)
	var _313_v133 _dafny.Map
	_ = _313_v133
	_313_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, (_dafny.Zero).Minus(_312_v132))
	_313_v133 = (_313_v133).Update(_168_v0, _312_v132)
	var _314_v134 _dafny.Sequence
	_ = _314_v134
	_314_v134 = _dafny.SeqOf(_312_v132, _dafny.IntOfInt64(390))
	var _315_v135 _dafny.Map
	_ = _315_v135
	_315_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_312_v132, _168_v0)
	var _316_v136 _dafny.Sequence
	_ = _316_v136
	_316_v136 = _dafny.UnicodeSeqOfUtf8Bytes("bmigsl")
	var _317_v137 *C2
	_ = _317_v137
	var _nw64 *C2 = New_C2_()
	_ = _nw64
	_nw64.Ctor__(Companion_Default___.Fm54(_168_v0, _316_v136, _167_globalState))
	_317_v137 = _nw64
	var _318_v138 _dafny.Sequence
	_ = _318_v138
	_318_v138 = _dafny.SeqOf(_317_v137, _317_v137, _317_v137, _317_v137, _317_v137)
	var _hi0 _dafny.Int = Companion_Default___.SafeDivisionInt((_315_v135).Cardinality(), _dafny.IntOfUint32((_318_v138).Cardinality()))
	_ = _hi0
	for _319_i5 := (_314_v134).Select((Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32()).(_dafny.Int); _319_i5.Cmp(_hi0) < 0; _319_i5 = _319_i5.Plus(_dafny.One) {
		var _320_v139 *C6
		_ = _320_v139
		var _nw65 *C6 = New_C6_()
		_ = _nw65
		_nw65.Ctor__()
		_320_v139 = _nw65
		var _321_v140 _dafny.Set
		_ = _321_v140
		_321_v140 = _dafny.SetOf(_312_v132)
		var _322_v141 _dafny.CodePoint
		_ = _322_v141
		_322_v141 = _dafny.CodePoint('s')
		var _323_v142 D0
		_ = _323_v142
		_323_v142 = Companion_D0_.Create_DC0_(_168_v0)
		var _324_v143 _dafny.Set
		_ = _324_v143
		_324_v143 = _dafny.SetOf(_168_v0, _168_v0)
		var _325_v144 D14
		_ = _325_v144
		_325_v144 = Companion_D14_.Create_DC26_(_323_v142, _324_v143)
		var _326_v145 _dafny.Map
		_ = _326_v145
		_326_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _168_v0)
		var _327_v146 _dafny.Array
		_ = _327_v146
		var _nwElement0_13 _dafny.Sequence = _dafny.SeqOf((_dafny.Zero).Minus(_319_i5))
		_ = _nwElement0_13
		var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(17))
		_ = _nw66
		(_nw66).ArraySet1(_nwElement0_13, 0)
		(_nw66).ArraySet1(_314_v134, 1)
		(_nw66).ArraySet1(_314_v134, 2)
		(_nw66).ArraySet1(_314_v134, 3)
		(_nw66).ArraySet1(_314_v134, 4)
		(_nw66).ArraySet1(Companion_Default___.Fm36(_168_v0, _322_v141, _312_v132, _167_globalState), 5)
		(_nw66).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_312_v132, ((_325_v144).Dtor_cf44()).Cardinality()), (Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_dafny.SeqOf(_312_v132, ((_325_v144).Dtor_cf44()).Cardinality())).Cardinality()))).Uint32(), _319_i5), 6)
		(_nw66).ArraySet1(_314_v134, 7)
		(_nw66).ArraySet1(_314_v134, 8)
		(_nw66).ArraySet1(_314_v134, 9)
		(_nw66).ArraySet1(_314_v134, 10)
		(_nw66).ArraySet1(_314_v134, 11)
		(_nw66).ArraySet1(_314_v134, 12)
		(_nw66).ArraySet1(_314_v134, 13)
		(_nw66).ArraySet1(_dafny.Companion_Sequence_.Update(_314_v134, (Companion_Default___.SafeIndex((_326_v145).Cardinality(), _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32(), _319_i5), 14)
		(_nw66).ArraySet1(_dafny.Companion_Sequence_.Update(_314_v134, (Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32(), _319_i5), 15)
		(_nw66).ArraySet1(_314_v134, 16)
		_327_v146 = _nw66
		var _328_v147 bool
		_ = _328_v147
		var _329_v148 _dafny.Array
		_ = _329_v148
		var _330_v149 _dafny.Int
		_ = _330_v149
		var _out5 bool
		_ = _out5
		var _out6 _dafny.Array
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out5, _out6, _out7 = (_317_v137).M1(((_321_v140).Union(_321_v140)).Cardinality(), _327_v146, _168_v0, _dafny.Companion_Sequence_.IsProperPrefixOf(_314_v134, _314_v134), _167_globalState)
		_328_v147 = _out5
		_329_v148 = _out6
		_330_v149 = _out7
		_328_v147 = Companion_Default___.Fm5(_167_globalState)
		var _331_v150 *C4
		_ = _331_v150
		var _nw67 *C4 = New_C4_()
		_ = _nw67
		_nw67.Ctor__(_319_i5)
		_331_v150 = _nw67
		var _332_v151 D21
		_ = _332_v151
		_332_v151 = Companion_D21_.Create_DC39_(_331_v150)
		var _333_v152 _dafny.Sequence
		_ = _333_v152
		_333_v152 = _dafny.SeqOf((_332_v151).Dtor_cf58(), _331_v150, _331_v150, _331_v150)
		_331_v150 = (_333_v152).Select((Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_333_v152).Cardinality()))).Uint32()).(*C4)
	}
	if (func() bool {
		if _168_v0 {
			return _168_v0
		}
		return _168_v0
	})() {
		var _334_v153 _dafny.Int
		_ = _334_v153
		var _335_v154 _dafny.Int
		_ = _335_v154
		var _out8 _dafny.Int
		_ = _out8
		var _out9 _dafny.Int
		_ = _out9
		_out8, _out9 = (_317_v137).M3(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_314_v134, _dafny.SeqOf(_dafny.IntOfInt64(553))), (Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_314_v134, _dafny.SeqOf(_dafny.IntOfInt64(553)))).Cardinality()))).Uint32(), _312_v132)), _169_v1, _168_v0, !(_168_v0), _167_globalState)
		_334_v153 = _out8
		_335_v154 = _out9
		var _336_v155 T2
		_ = _336_v155
		var _nw68 *C4 = New_C4_()
		_ = _nw68
		_nw68.Ctor__(_334_v153)
		_336_v155 = _nw68
		var _337_v156 _dafny.Array
		_ = _337_v156
		var _nwElement0_14 T2 = _336_v155
		_ = _nwElement0_14
		var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(2))
		_ = _nw69
		(_nw69).ArraySet1(_nwElement0_14, 0)
		(_nw69).ArraySet1(_336_v155, 1)
		_337_v156 = _nw69
		_337_v156 = _337_v156
		var _338_v157 *C7
		_ = _338_v157
		var _nw70 *C7 = New_C7_()
		_ = _nw70
		_nw70.Ctor__()
		_338_v157 = _nw70
		var _339_v158 D20
		_ = _339_v158
		_339_v158 = Companion_D20_.Create_DC38_(_338_v157)
		var _340_v159 _dafny.Set
		_ = _340_v159
		_340_v159 = _dafny.SetOf((func() D20 {
			if _168_v0 {
				return _339_v158
			}
			return Companion_D20_.Create_DC38_(_338_v157)
		})())
		var _341_v160 _dafny.MultiSet
		_ = _341_v160
		_341_v160 = _dafny.MultiSetOf(_168_v0, _168_v0)
		var _pat_let_tv5 = _340_v159
		_ = _pat_let_tv5
		var _rhs41 bool = !(_168_v0) || (!((_341_v160).IsSubsetOf(_341_v160)))
		_ = _rhs41
		var _rhs42 _dafny.Set = (func(_pat_let6_0 D22) D22 {
			return func(_342_dt__update__tmp_h1 D22) D22 {
				return func(_pat_let7_0 _dafny.Set) D22 {
					return func(_343_dt__update_hcf62_h0 _dafny.Set) D22 {
						return Companion_D22_.Create_DC43_(_343_dt__update_hcf62_h0)
					}(_pat_let7_0)
				}(_pat_let_tv5)
			}(_pat_let6_0)
		}(Companion_D22_.Create_DC43_(_340_v159))).Dtor_cf62()
		_ = _rhs42
		var _rhs43 bool = _168_v0
		_ = _rhs43
		_168_v0 = _rhs41
		_340_v159 = _rhs42
		_168_v0 = _rhs43
		_334_v153 = Companion_Default___.SafeDivisionInt(_335_v154, _312_v132)
		var _344_v161 _dafny.Sequence
		_ = _344_v161
		_344_v161 = _dafny.SeqOf(_168_v0, !(_168_v0))
		var _345_v162 D22
		_ = _345_v162
		_345_v162 = Companion_D22_.Create_DC44_(Companion_Default___.Fm41(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(930))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}((func(_346_v160 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
			return func(_347_i6 _dafny.Int) _dafny.Int {
				return (_346_v160).Cardinality()
			}
		})(_341_v160))), _168_v0, (_344_v161).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.IntOfUint32((_344_v161).Cardinality()))).Uint32()).(bool), _168_v0, _167_globalState))
		_345_v162 = _345_v162
	} else {
		var _348_v163 T2
		_ = _348_v163
		var _nw71 *C2 = New_C2_()
		_ = _nw71
		_nw71.Ctor__(_312_v132)
		_348_v163 = _nw71
		_348_v163 = _348_v163
		var _349_v164 _dafny.Sequence
		_ = _349_v164
		_349_v164 = _dafny.SeqOf(_168_v0)
		_313_v133 = (_313_v133).Update((func() bool {
			if _168_v0 {
				return _168_v0
			}
			return true
		})(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_349_v164).Cardinality()), _312_v132))
		var _350_v165 _dafny.MultiSet
		_ = _350_v165
		_350_v165 = _dafny.MultiSetOf(_168_v0, _168_v0)
		var _351_v166 _dafny.Array
		_ = _351_v166
		var _nwElement0_15 bool = _168_v0
		_ = _nwElement0_15
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(23))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_15, 0)
		(_nw72).ArraySet1(_168_v0, 1)
		(_nw72).ArraySet1((Companion_D5_.Create_DC9_(_168_v0, _312_v132, _350_v165)).Dtor_cf14(), 2)
		(_nw72).ArraySet1(_168_v0, 3)
		(_nw72).ArraySet1(_168_v0, 4)
		(_nw72).ArraySet1(_168_v0, 5)
		(_nw72).ArraySet1(!(_168_v0), 6)
		(_nw72).ArraySet1(_168_v0, 7)
		(_nw72).ArraySet1(_168_v0, 8)
		(_nw72).ArraySet1(_168_v0, 9)
		(_nw72).ArraySet1(_168_v0, 10)
		(_nw72).ArraySet1((_349_v164).Select((Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_349_v164).Cardinality()))).Uint32()).(bool), 11)
		(_nw72).ArraySet1(_168_v0, 12)
		(_nw72).ArraySet1(_168_v0, 13)
		(_nw72).ArraySet1(_168_v0, 14)
		(_nw72).ArraySet1(_168_v0, 15)
		(_nw72).ArraySet1(_168_v0, 16)
		(_nw72).ArraySet1(true, 17)
		(_nw72).ArraySet1(_168_v0, 18)
		(_nw72).ArraySet1(true, 19)
		(_nw72).ArraySet1(!(_168_v0), 20)
		(_nw72).ArraySet1(_168_v0, 21)
		(_nw72).ArraySet1(_168_v0, 22)
		_351_v166 = _nw72
		var _352_v167 _dafny.Map
		_ = _352_v167
		_352_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v0, _168_v0)
		var _353_v168 _dafny.Map
		_ = _353_v168
		_353_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_351_v166, !(((Companion_Default___.Fm26(_312_v132, _168_v0, _168_v0, Companion_Default___.Fm5(_167_globalState), _167_globalState)).Update(_168_v0, _168_v0)).Equals(_352_v167)))
		_353_v168 = (_353_v168).Update(_351_v166, _168_v0)
		if _168_v0 {
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_351_v166), 0))
			_ = _index57
			(_351_v166).ArraySet1(_168_v0, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_351_v166), 0))
			_ = _index58
			(_351_v166).ArraySet1(_168_v0, (_index58).Int())
			_168_v0 = false
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_351_v166), 0))
			_ = _index59
			var _rhs44 bool = ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bn")).Cardinality())).Times(_312_v132)).Cmp(_348_v163.F2()) != 0
			_ = _rhs44
			var _rhs45 _dafny.Int = _348_v163.F2()
			_ = _rhs45
			var _rhs46 bool = false
			_ = _rhs46
			var _lhs21 _dafny.Array = _351_v166
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_351_v166), 0))
			_ = _lhs22
			var _lhs23 T2 = _348_v163
			_ = _lhs23
			(_lhs21).ArraySet1(_rhs44, (_lhs22).Int())
			_lhs23.F2_set_(_rhs45)
			_168_v0 = _rhs46
			var _354_v169 D6
			_ = _354_v169
			_354_v169 = Companion_D6_.Create_DC12_(_168_v0, (_351_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_351_v166), 0))).Int()).(bool))
			_354_v169 = _354_v169
			var _355_v170 bool
			_ = _355_v170
			var _356_v171 _dafny.Int
			_ = _356_v171
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out10, _out11 = (_317_v137).M6(_348_v163.F2(), _167_globalState)
			_355_v170 = _out10
			_356_v171 = _out11
		} else {
			var _357_v172 D6
			_ = _357_v172
			_357_v172 = Companion_D6_.Create_DC13_(true, _312_v132)
			_168_v0 = (_357_v172).Dtor_cf24()
			var _358_v173 _dafny.Map
			_ = _358_v173
			_358_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v134, (_dafny.Zero).Minus(_348_v163.F2()))
			var _359_v174 _dafny.Map
			_ = _359_v174
			_359_v174 = _358_v173
			_359_v174 = _358_v173
			var _360_v175 _dafny.Map
			_ = _360_v175
			_360_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_348_v163.F2()).Times(_348_v163.F2()), _314_v134)
			_314_v134 = (func() _dafny.Sequence {
				if (_360_v175).Contains(_312_v132) {
					return (_360_v175).Get(_312_v132).(_dafny.Sequence)
				}
				return _314_v134
			})()
			var _361_v176 _dafny.Array
			_ = _361_v176
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_6
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Sequence = (func(_362_v134 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_363_i7 _dafny.Int) _dafny.Sequence {
						return _362_v134
					}
				})(_314_v134)
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
			_361_v176 = _nw73
			var _364_v177 bool
			_ = _364_v177
			var _365_v178 _dafny.Array
			_ = _365_v178
			var _366_v179 _dafny.Int
			_ = _366_v179
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Array
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			_out12, _out13, _out14 = (_348_v163).M1(_348_v163.F2(), _361_v176, _168_v0, _168_v0, _167_globalState)
			_364_v177 = _out12
			_365_v178 = _out13
			_366_v179 = _out14
			var _367_v180 *C1
			_ = _367_v180
			var _nw74 *C1 = New_C1_()
			_ = _nw74
			_nw74.Ctor__(_dafny.Companion_Sequence_.Concatenate(_316_v136, _dafny.UnicodeSeqOfUtf8Bytes("fifwva")))
			_367_v180 = _nw74
			_367_v180 = _367_v180
		}
		var _368_v181 *C0
		_ = _368_v181
		var _nw75 *C0 = New_C0_()
		_ = _nw75
		_nw75.Ctor__((_348_v163.F2()).Cmp((_dafny.Zero).Minus(_312_v132)) < 0, (_349_v164).Select((Companion_Default___.SafeIndex(_348_v163.F2(), _dafny.IntOfUint32((_349_v164).Cardinality()))).Uint32()).(bool))
		_368_v181 = _nw75
	}
	var _369_v182 bool
	_ = _369_v182
	var _370_v183 _dafny.Int
	_ = _370_v183
	var _out15 bool
	_ = _out15
	var _out16 _dafny.Int
	_ = _out16
	_out15, _out16 = (_317_v137).M6(_312_v132, _167_globalState)
	_369_v182 = _out15
	_370_v183 = _out16
	var _hi1 _dafny.Int = _370_v183
	_ = _hi1
	for _371_i8 := _370_v183; _371_i8.Cmp(_hi1) < 0; _371_i8 = _371_i8.Plus(_dafny.One) {
		var _372_v184 _dafny.CodePoint
		_ = _372_v184
		_372_v184 = _dafny.CodePoint('j')
		var _373_v185 D7
		_ = _373_v185
		_373_v185 = Companion_D7_.Create_DC15_(_312_v132, _372_v184)
		var _374_v186 T1
		_ = _374_v186
		var _nw76 *C6 = New_C6_()
		_ = _nw76
		_nw76.Ctor__()
		_374_v186 = _nw76
		var _pat_let_tv6 = _372_v184
		_ = _pat_let_tv6
		var _375_v187 _dafny.Map
		_ = _375_v187
		_375_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let8_0 D7) D7 {
			return func(_376_dt__update__tmp_h2 D7) D7 {
				return func(_pat_let9_0 _dafny.CodePoint) D7 {
					return func(_377_dt__update_hcf28_h0 _dafny.CodePoint) D7 {
						return Companion_D7_.Create_DC15_((_376_dt__update__tmp_h2).Dtor_cf27(), _377_dt__update_hcf28_h0)
					}(_pat_let9_0)
				}(_pat_let_tv6)
			}(_pat_let8_0)
		}(_373_v185)).Dtor_cf27(), _374_v186)
		_375_v187 = (_375_v187).Update((_dafny.Zero).Minus((func() _dafny.Int {
			if true {
				return _370_v183
			}
			return _371_i8
		})()), _374_v186)
		var _378_v188 _dafny.Map
		_ = _378_v188
		_378_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_316_v136, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_i8, _369_v182))
		var _379_v189 _dafny.Sequence
		_ = _379_v189
		_379_v189 = _dafny.SeqOf(_369_v182, false, _168_v0, false)
		var _380_v190 D5
		_ = _380_v190
		_380_v190 = Companion_D5_.Create_DC10_(_168_v0, _370_v183, (_378_v188).Merge(_378_v188), !((_379_v189).Select((Companion_Default___.SafeIndex((_314_v134).Select((Companion_Default___.SafeIndex(_370_v183, _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_379_v189).Cardinality()))).Uint32()).(bool)) || (_168_v0))
		var _source7 D5 = _380_v190
		_ = _source7
		if _source7.Is_DC9() {
			var _381___mcc_h1 bool = _source7.Get_().(D5_DC9).Cf14
			_ = _381___mcc_h1
			var _382___mcc_h2 _dafny.Int = _source7.Get_().(D5_DC9).Cf15
			_ = _382___mcc_h2
			var _383___mcc_h3 _dafny.MultiSet = _source7.Get_().(D5_DC9).Cf16
			_ = _383___mcc_h3
			var _384_cf16 _dafny.MultiSet = _383___mcc_h3
			_ = _384_cf16
			var _385_cf15 _dafny.Int = _382___mcc_h2
			_ = _385_cf15
			var _386_cf14 bool = _381___mcc_h1
			_ = _386_cf14
			_168_v0 = false
			var _387_v191 _dafny.Int
			_ = _387_v191
			var _out17 _dafny.Int
			_ = _out17
			_out17 = (_317_v137).M2(_312_v132, _dafny.IntOfInt64(597), !(_168_v0) || (!(_369_v182)), _167_globalState)
			_387_v191 = _out17
			_370_v183 = ((func() _dafny.Int {
				if _369_v182 {
					return _312_v132
				}
				return _370_v183
			})()).Minus((_374_v186).Fm1(true, _167_globalState))
			_168_v0 = (_379_v189).Select((Companion_Default___.SafeIndex(_385_cf15, _dafny.IntOfUint32((_379_v189).Cardinality()))).Uint32()).(bool)
		} else if _source7.Is_DC10() {
			var _388___mcc_h4 bool = _source7.Get_().(D5_DC10).Cf17
			_ = _388___mcc_h4
			var _389___mcc_h5 _dafny.Int = _source7.Get_().(D5_DC10).Cf18
			_ = _389___mcc_h5
			var _390___mcc_h6 _dafny.Map = _source7.Get_().(D5_DC10).Cf19
			_ = _390___mcc_h6
			var _391___mcc_h7 bool = _source7.Get_().(D5_DC10).Cf20
			_ = _391___mcc_h7
			var _392_cf20 bool = _391___mcc_h7
			_ = _392_cf20
			var _393_cf19 _dafny.Map = _390___mcc_h6
			_ = _393_cf19
			var _394_cf18 _dafny.Int = _389___mcc_h5
			_ = _394_cf18
			var _395_cf17 bool = _388___mcc_h4
			_ = _395_cf17
			var _396_v192 _dafny.Int
			_ = _396_v192
			var _out18 _dafny.Int
			_ = _out18
			_out18 = (_317_v137).M2(_312_v132, _dafny.IntOfInt64(223), _369_v182, _167_globalState)
			_396_v192 = _out18
			_315_v135 = (_315_v135).Update((_371_i8).Minus(_394_cf18), (func() bool {
				if _395_cf17 {
					return _168_v0
				}
				return _395_cf17
			})())
			_317_v137 = _317_v137
			_317_v137 = _317_v137
		} else {
			var _397___mcc_h8 _dafny.MultiSet = _source7.Get_().(D5_DC8).Cf13
			_ = _397___mcc_h8
			var _398_cf13 _dafny.MultiSet = _397___mcc_h8
			_ = _398_cf13
			var _399_v193 *C0
			_ = _399_v193
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(true, _168_v0)
			_399_v193 = _nw77
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_169_v1), 0))
			_ = _index60
			(_169_v1).ArraySet1(_371_i8, (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_169_v1), 0))
			_ = _index61
			(_169_v1).ArraySet1(Companion_Default___.SafeModuloInt(_370_v183, _dafny.IntOfInt64(-579)), (_index61).Int())
			_369_v182 = false
			var _400_v194 _dafny.Set
			_ = _400_v194
			_400_v194 = _dafny.SetOf(_399_v193.F4, _399_v193.F4, _399_v193.F4, _399_v193.F4)
			_400_v194 = (_400_v194).Intersection(_400_v194)
		}
		var _401_v195 *C4
		_ = _401_v195
		var _nw78 *C4 = New_C4_()
		_ = _nw78
		_nw78.Ctor__((func() _dafny.Int {
			if (_313_v133).Contains(false) {
				return (_313_v133).Get(false).(_dafny.Int)
			}
			return _dafny.IntOfInt64(988)
		})())
		_401_v195 = _nw78
		var _nw79 *C4 = New_C4_()
		_ = _nw79
		_nw79.Ctor__(_370_v183)
		_401_v195 = _nw79
		var _402_v196 *C7
		_ = _402_v196
		var _nw80 *C7 = New_C7_()
		_ = _nw80
		_nw80.Ctor__()
		_402_v196 = _nw80
		_402_v196 = _402_v196
	}
	for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_169_v1), 0))); ; {
		_guard_loop_0, _ok17 := _iter17()
		if !_ok17 {
			break
		}
		var _403_i9 _dafny.Int
		_403_i9 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_403_i9).Sign() != -1) && ((_403_i9).Cmp(_dafny.ArrayLen((_169_v1), 0)) < 0)) {
			(_169_v1).ArraySet1((_403_i9).Plus(_312_v132), (_403_i9).Int())
		}
	}
	if (Companion_Default___.SafeModuloInt(_370_v183, _370_v183)).Cmp(_370_v183) <= 0 {
		if _168_v0 {
			var _404_v197 T0
			_ = _404_v197
			var _nw81 *C6 = New_C6_()
			_ = _nw81
			_nw81.Ctor__()
			_404_v197 = _nw81
			var _405_v198 T0
			_ = _405_v198
			_405_v198 = _404_v197
			var _406_v199 _dafny.Array
			_ = _406_v199
			var _nwElement0_16 T0 = _405_v198
			_ = _nwElement0_16
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(15))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_16, 0)
			(_nw82).ArraySet1(_405_v198, 1)
			(_nw82).ArraySet1(_404_v197, 2)
			(_nw82).ArraySet1(_405_v198, 3)
			(_nw82).ArraySet1(_405_v198, 4)
			(_nw82).ArraySet1(_405_v198, 5)
			(_nw82).ArraySet1(_405_v198, 6)
			(_nw82).ArraySet1(_405_v198, 7)
			(_nw82).ArraySet1(_405_v198, 8)
			(_nw82).ArraySet1(_404_v197, 9)
			(_nw82).ArraySet1(_405_v198, 10)
			(_nw82).ArraySet1(_405_v198, 11)
			(_nw82).ArraySet1(_405_v198, 12)
			(_nw82).ArraySet1(_405_v198, 13)
			(_nw82).ArraySet1(_405_v198, 14)
			_406_v199 = _nw82
			_406_v199 = _406_v199
			var _407_v200 _dafny.Set
			_ = _407_v200
			_407_v200 = _dafny.SetOf(_369_v182)
			_168_v0 = (_407_v200).IsDisjointFrom((_407_v200).Difference(_407_v200))
			var _408_v201 D6
			_ = _408_v201
			_408_v201 = Companion_D6_.Create_DC11_(_316_v136)
			_316_v136 = _dafny.Companion_Sequence_.Concatenate(_316_v136, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bhskj"), (_408_v201).Dtor_cf21()))
			var _409_v202 _dafny.Array
			_ = _409_v202
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
			_ = _nw83
			_409_v202 = _nw83
			_409_v202 = _409_v202
			var _410_v203 _dafny.Array
			_ = _410_v203
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw84
			_410_v203 = _nw84
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_410_v203), 0))
			_ = _index62
			(_410_v203).ArraySet1(_169_v1, (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_410_v203), 0))
			_ = _index63
			(_410_v203).ArraySet1((func() _dafny.Array {
				if true {
					return _169_v1
				}
				return _169_v1
			})(), (_index63).Int())
		} else {
			_369_v182 = !(((_317_v137).Fm1(false, _167_globalState)).Cmp(Companion_Default___.SafeDivisionInt(_312_v132, _dafny.IntOfInt64(785))) <= 0)
			var _411_v204 _dafny.Array
			_ = _411_v204
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw85
			_411_v204 = _nw85
			var _412_v205 _dafny.Array
			_ = _412_v205
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw86
			_412_v205 = _nw86
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_411_v204), 0))
			_ = _index64
			(_411_v204).ArraySet1(_412_v205, (_index64).Int())
			var _413_v206 *C7
			_ = _413_v206
			var _nw87 *C7 = New_C7_()
			_ = _nw87
			_nw87.Ctor__()
			_413_v206 = _nw87
			var _414_v207 D20
			_ = _414_v207
			_414_v207 = Companion_D20_.Create_DC38_(_413_v206)
			var _415_v208 _dafny.Set
			_ = _415_v208
			_415_v208 = _dafny.SetOf(_414_v207, _414_v207, _414_v207)
			var _416_v209 D22
			_ = _416_v209
			_416_v209 = Companion_D22_.Create_DC43_(_415_v208)
			var _417_v210 _dafny.Sequence
			_ = _417_v210
			_417_v210 = _dafny.SeqOf(_416_v209)
			var _418_v211 *C4
			_ = _418_v211
			var _nw88 *C4 = New_C4_()
			_ = _nw88
			_nw88.Ctor__(_312_v132)
			_418_v211 = _nw88
			var _419_v212 _dafny.Map
			_ = _419_v212
			_419_v212 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_314_v134).Select((Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32()).(_dafny.Int), _418_v211)
			var _420_v213 _dafny.Sequence
			_ = _420_v213
			_420_v213 = _dafny.SeqOf(_418_v211, _418_v211, _418_v211)
			var _421_v214 _dafny.Array
			_ = _421_v214
			var _nwElement0_17 *C4 = _418_v211
			_ = _nwElement0_17
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_17, 0)
			(_nw89).ArraySet1(_418_v211, 1)
			(_nw89).ArraySet1(_418_v211, 2)
			(_nw89).ArraySet1(_418_v211, 3)
			(_nw89).ArraySet1((func() *C4 {
				if _168_v0 {
					return _418_v211
				}
				return _418_v211
			})(), 4)
			(_nw89).ArraySet1(_418_v211, 5)
			(_nw89).ArraySet1(_418_v211, 6)
			(_nw89).ArraySet1(_418_v211, 7)
			(_nw89).ArraySet1((func() *C4 {
				if (_419_v212).Contains(_312_v132) {
					return (_419_v212).Get(_312_v132).(*C4)
				}
				return _418_v211
			})(), 8)
			(_nw89).ArraySet1(_418_v211, 9)
			(_nw89).ArraySet1((_420_v213).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("orqkbxy")).Cardinality()), _dafny.IntOfUint32((_420_v213).Cardinality()))).Uint32()).(*C4), 10)
			(_nw89).ArraySet1(_418_v211, 11)
			(_nw89).ArraySet1(_418_v211, 12)
			(_nw89).ArraySet1(_418_v211, 13)
			(_nw89).ArraySet1(_418_v211, 14)
			(_nw89).ArraySet1(_418_v211, 15)
			(_nw89).ArraySet1(_418_v211, 16)
			(_nw89).ArraySet1(_418_v211, 17)
			(_nw89).ArraySet1(_418_v211, 18)
			(_nw89).ArraySet1(_418_v211, 19)
			(_nw89).ArraySet1(_418_v211, 20)
			(_nw89).ArraySet1(_418_v211, 21)
			(_nw89).ArraySet1(_418_v211, 22)
			(_nw89).ArraySet1(_418_v211, 23)
			(_nw89).ArraySet1(_418_v211, 24)
			(_nw89).ArraySet1(_418_v211, 25)
			(_nw89).ArraySet1(_418_v211, 26)
			(_nw89).ArraySet1(_418_v211, 27)
			(_nw89).ArraySet1(_418_v211, 28)
			_421_v214 = _nw89
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_421_v214), 0))
			_ = _index65
			(_421_v214).ArraySet1(_418_v211, (_index65).Int())
			var _422_v215 _dafny.CodePoint
			_ = _422_v215
			_422_v215 = _dafny.CodePoint('r')
			var _423_v216 _dafny.MultiSet
			_ = _423_v216
			_423_v216 = _dafny.MultiSetOf(_369_v182)
			var _424_v217 _dafny.Map
			_ = _424_v217
			_424_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v216, _369_v182)
			var _425_v218 D21
			_ = _425_v218
			_425_v218 = Companion_D21_.Create_DC39_(_418_v211)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_411_v204), 0))
			_ = _index66
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_421_v214), 0))
			_ = _index67
			var _rhs47 _dafny.Array = _412_v205
			_ = _rhs47
			var _rhs48 _dafny.Int = ((Companion_Default___.Fm55(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_426_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), _422_v215, _167_globalState)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v216, true)).Merge(_424_v217))).Cardinality()
			_ = _rhs48
			var _rhs49 _dafny.Int = Companion_Default___.Fm54(_168_v0, _dafny.UnicodeSeqOfUtf8Bytes("kmnormpf"), _167_globalState)
			_ = _rhs49
			var _rhs50 _dafny.Sequence = _417_v210
			_ = _rhs50
			var _rhs51 *C4 = (func() *C4 {
				if _369_v182 {
					return _418_v211
				}
				return (_425_v218).Dtor_cf58()
			})()
			_ = _rhs51
			var _lhs24 _dafny.Array = _411_v204
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_411_v204), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _421_v214
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_421_v214), 0))
			_ = _lhs27
			(_lhs24).ArraySet1(_rhs47, (_lhs25).Int())
			_370_v183 = _rhs48
			_312_v132 = _rhs49
			_417_v210 = _rhs50
			(_lhs26).ArraySet1(_rhs51, (_lhs27).Int())
			_370_v183 = (_418_v211).Fm1(_168_v0, _167_globalState)
			var _427_v219 _dafny.Map
			_ = _427_v219
			_427_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v1, _369_v182)
			var _rhs52 bool = _369_v182
			_ = _rhs52
			var _rhs53 bool = _369_v182
			_ = _rhs53
			var _rhs54 _dafny.Map = ((_427_v219).Merge(_427_v219)).Merge(_427_v219)
			_ = _rhs54
			_168_v0 = _rhs52
			_168_v0 = _rhs53
			_427_v219 = _rhs54
			var _428_v220 _dafny.Array
			_ = _428_v220
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_7
			var _nw90 _dafny.Array
			_ = _nw90
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw90 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_429_v183 _dafny.Int, _430_v135 _dafny.Map, _431_v0 bool) func(_dafny.Int) bool {
					return func(_432_i11 _dafny.Int) bool {
						return (func() bool {
							if (_430_v135).Contains(_429_v183) {
								return (_430_v135).Get(_429_v183).(bool)
							}
							return _431_v0
						})()
					}
				})(_370_v183, _315_v135, _168_v0)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw90 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw90).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw90).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_428_v220 = _nw90
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_428_v220), 0))
			_ = _index68
			(_428_v220).ArraySet1(_168_v0, (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_428_v220), 0))
			_ = _index69
			(_428_v220).ArraySet1(_369_v182, (_index69).Int())
		}
		var _433_v221 _dafny.Array
		_ = _433_v221
		var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(19))
		_ = _nw91
		_433_v221 = _nw91
		var _434_v222 T2
		_ = _434_v222
		var _nw92 *C4 = New_C4_()
		_ = _nw92
		_nw92.Ctor__(_370_v183)
		_434_v222 = _nw92
		var _435_v223 T2
		_ = _435_v223
		_435_v223 = _434_v222
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_433_v221), 0))
		_ = _index70
		(_433_v221).ArraySet1(_dafny.MultiSetOf(_435_v223, _435_v223), (_index70).Int())
		var _436_v224 _dafny.Set
		_ = _436_v224
		_436_v224 = _dafny.SetOf(_434_v222.F2(), _434_v222.F2(), _dafny.IntOfInt64(658), _434_v222.F2(), (_370_v183).Times((_dafny.Zero).Minus((Companion_Default___.Fm38(_168_v0, _167_globalState)).Cardinality())))
		var _437_v225 _dafny.MultiSet
		_ = _437_v225
		_437_v225 = _dafny.MultiSetOf(_435_v223)
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_433_v221), 0))
		_ = _index71
		var _rhs55 _dafny.Int = (_436_v224).Cardinality()
		_ = _rhs55
		var _rhs56 _dafny.MultiSet = _437_v225
		_ = _rhs56
		var _lhs28 _dafny.Array = _433_v221
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_433_v221), 0))
		_ = _lhs29
		_370_v183 = _rhs55
		(_lhs28).ArraySet1(_rhs56, (_lhs29).Int())
		var _438_v226 D21
		_ = _438_v226
		_438_v226 = Companion_D21_.Create_DC41_(_369_v182)
		var _439_v227 *C3
		_ = _439_v227
		var _nw93 *C3 = New_C3_()
		_ = _nw93
		_nw93.Ctor__((_314_v134).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm36(_369_v182, _dafny.CodePoint('i'), _312_v132, _167_globalState)).Cardinality()), _dafny.IntOfUint32((_314_v134).Cardinality()))).Uint32()).(_dafny.Int))
		_439_v227 = _nw93
		var _440_v228 _dafny.Map
		_ = _440_v228
		_440_v228 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_438_v226).Dtor_cf60(), _439_v227)
		var _441_v229 _dafny.Sequence
		_ = _441_v229
		_441_v229 = _dafny.SeqOf(_439_v227)
		_440_v228 = (_440_v228).Update((func() bool {
			if _168_v0 {
				return _168_v0
			}
			return _168_v0
		})(), (func() *C3 {
			if (_440_v228).Contains(false) {
				return (_440_v228).Get(false).(*C3)
			}
			return (_441_v229).Select((Companion_Default___.SafeIndex(_312_v132, _dafny.IntOfUint32((_441_v229).Cardinality()))).Uint32()).(*C3)
		})())
		var _442_v230 _dafny.CodePoint
		_ = _442_v230
		_442_v230 = _dafny.CodePoint('w')
		var _443_v231 D20
		_ = _443_v231
		_443_v231 = Companion_D20_.Create_DC37_(_442_v230, !(_168_v0))
		var _444_v232 _dafny.Map
		_ = _444_v232
		_444_v232 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_312_v132)), _443_v231)
		var _445_v233 _dafny.MultiSet
		_ = _445_v233
		_445_v233 = _dafny.MultiSetOf(_316_v136)
		var _pat_let_tv7 = _369_v182
		_ = _pat_let_tv7
		_444_v232 = (_444_v232).Update((func() _dafny.Int {
			if (_445_v233).Contains(_316_v136) {
				return (_445_v233).Multiplicity(_316_v136)
			}
			return _434_v222.F2()
		})(), func(_pat_let10_0 D20) D20 {
			return func(_446_dt__update__tmp_h4 D20) D20 {
				return func(_pat_let11_0 bool) D20 {
					return func(_447_dt__update_hcf56_h0 bool) D20 {
						return Companion_D20_.Create_DC37_((_446_dt__update__tmp_h4).Dtor_cf55(), _447_dt__update_hcf56_h0)
					}(_pat_let11_0)
				}(_pat_let_tv7)
			}(_pat_let10_0)
		}(_443_v231))
		var _448_v234 _dafny.Array
		_ = _448_v234
		var _nwElement0_18 bool = _369_v182
		_ = _nwElement0_18
		var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(9))
		_ = _nw94
		(_nw94).ArraySet1(_nwElement0_18, 0)
		(_nw94).ArraySet1(_168_v0, 1)
		(_nw94).ArraySet1(_168_v0, 2)
		(_nw94).ArraySet1(_168_v0, 3)
		(_nw94).ArraySet1(_168_v0, 4)
		(_nw94).ArraySet1(_369_v182, 5)
		(_nw94).ArraySet1(!(_369_v182), 6)
		(_nw94).ArraySet1(_369_v182, 7)
		(_nw94).ArraySet1(_369_v182, 8)
		_448_v234 = _nw94
		var _449_v235 _dafny.Map
		_ = _449_v235
		_449_v235 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v182, _448_v234)
		var _450_v236 _dafny.Array
		_ = _450_v236
		var _nwElement0_19 _dafny.Array = _448_v234
		_ = _nwElement0_19
		var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(16))
		_ = _nw95
		(_nw95).ArraySet1(_nwElement0_19, 0)
		(_nw95).ArraySet1(_448_v234, 1)
		(_nw95).ArraySet1(_448_v234, 2)
		(_nw95).ArraySet1(_448_v234, 3)
		(_nw95).ArraySet1(_448_v234, 4)
		(_nw95).ArraySet1(_448_v234, 5)
		(_nw95).ArraySet1(_448_v234, 6)
		(_nw95).ArraySet1(_448_v234, 7)
		(_nw95).ArraySet1(_448_v234, 8)
		(_nw95).ArraySet1(_448_v234, 9)
		(_nw95).ArraySet1(_448_v234, 10)
		(_nw95).ArraySet1((func() _dafny.Array {
			if (_449_v235).Contains(true) {
				return (_449_v235).Get(true).(_dafny.Array)
			}
			return _448_v234
		})(), 11)
		(_nw95).ArraySet1(_448_v234, 12)
		(_nw95).ArraySet1(_448_v234, 13)
		(_nw95).ArraySet1(_448_v234, 14)
		(_nw95).ArraySet1((func() _dafny.Array {
			if !(_168_v0) {
				return _448_v234
			}
			return _448_v234
		})(), 15)
		_450_v236 = _nw95
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_450_v236), 0))
		_ = _index72
		(_450_v236).ArraySet1(_448_v234, (_index72).Int())
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_450_v236), 0))
		_ = _index73
		(_450_v236).ArraySet1(_448_v234, (_index73).Int())
	} else {
		var _451_v237 _dafny.Array
		_ = _451_v237
		var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
		_ = _nw96
		_451_v237 = _nw96
		var _452_v238 _dafny.Array
		_ = _452_v238
		var _nwElement0_20 _dafny.Array = _451_v237
		_ = _nwElement0_20
		var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(9))
		_ = _nw97
		(_nw97).ArraySet1(_nwElement0_20, 0)
		(_nw97).ArraySet1(_451_v237, 1)
		(_nw97).ArraySet1(_451_v237, 2)
		(_nw97).ArraySet1(_451_v237, 3)
		(_nw97).ArraySet1(_451_v237, 4)
		(_nw97).ArraySet1(_451_v237, 5)
		(_nw97).ArraySet1(_451_v237, 6)
		(_nw97).ArraySet1(_451_v237, 7)
		(_nw97).ArraySet1(_451_v237, 8)
		_452_v238 = _nw97
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_452_v238), 0))
		_ = _index74
		(_452_v238).ArraySet1(_451_v237, (_index74).Int())
		var _453_v239 _dafny.Sequence
		_ = _453_v239
		_453_v239 = _dafny.SeqOf(_451_v237)
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_452_v238), 0))
		_ = _index75
		(_452_v238).ArraySet1((_453_v239).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_316_v136).Cardinality()), _dafny.IntOfUint32((_453_v239).Cardinality()))).Uint32()).(_dafny.Array), (_index75).Int())
		_370_v183 = (_dafny.Zero).Minus(((_dafny.IntOfUint32((_314_v134).Cardinality())).Times(_370_v183)).Times((_dafny.MultiSetOf((Companion_Default___.Fm56(_168_v0, _167_globalState)).Cardinality(), _370_v183, _dafny.IntOfUint32((_316_v136).Cardinality()), _312_v132)).Cardinality()))
		var _454_v240 _dafny.CodePoint
		_ = _454_v240
		_454_v240 = _dafny.CodePoint('a')
		_454_v240 = _454_v240
		_454_v240 = _454_v240
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_169_v1), 0))
		_ = _index76
		(_169_v1).ArraySet1(_dafny.IntOfInt64(390), (_index76).Int())
		var _455_v241 _dafny.MultiSet
		_ = _455_v241
		_455_v241 = _dafny.MultiSetOf(_454_v240, _454_v240)
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_169_v1), 0))
		_ = _index77
		var _rhs57 _dafny.Int = (_dafny.Zero).Minus(_370_v183)
		_ = _rhs57
		var _rhs58 bool = !(_455_v241).Contains(_454_v240)
		_ = _rhs58
		var _rhs59 _dafny.Int = (_dafny.Zero).Minus(_370_v183)
		_ = _rhs59
		var _lhs30 _dafny.Array = _169_v1
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_169_v1), 0))
		_ = _lhs31
		_312_v132 = _rhs57
		_168_v0 = _rhs58
		(_lhs30).ArraySet1(_rhs59, (_lhs31).Int())
	}
	_312_v132 = _370_v183
	var _456_v242 _dafny.Array
	_ = _456_v242
	var _nw98 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw98
	_456_v242 = _nw98
	var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_456_v242), 0))
	_ = _index78
	(_456_v242).ArraySet1(Companion_Default___.Fm3(_168_v0, _312_v132, _312_v132, _167_globalState), (_index78).Int())
	var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_456_v242), 0))
	_ = _index79
	(_456_v242).ArraySet1(!_dafny.Companion_Sequence_.Equal(_316_v136, _316_v136), (_index79).Int())
	_dafny.Print(_168_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v27).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_312_v132)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v133).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(786)).UpdateUnsafe(true, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_314_v134, _dafny.SeqOf(_dafny.IntOfInt64(786), _dafny.IntOfInt64(390))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v135).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(786), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_316_v136.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_318_v138).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_369_v182)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_370_v183)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_456_v242).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
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
	Cf2 bool
	Cf3 _dafny.CodePoint
	Cf4 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 _dafny.CodePoint, Cf4 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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

type D1_DC2 struct {
	Cf5 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf5 _dafny.Map) D1 {
	return D1{D1_DC2{Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D2_DC3 struct {
	Cf6 _dafny.Map
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf6 _dafny.Map) D2 {
	return D2{D2_DC3{Cf6}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D2_DC3).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf6) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
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

type D3_DC4 struct {
	Cf7 _dafny.Array
}

func (D3_DC4) isD3() {}

func (CompanionStruct_D3_) Create_DC4_(Cf7 _dafny.Array) D3 {
	return D3{D3_DC4{Cf7}}
}

func (_this D3) Is_DC4() bool {
	_, ok := _this.Get_().(D3_DC4)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D3) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D3_DC4).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC4:
		{
			return "D3.DC4" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC4:
		{
			data2, ok := other.Get_().(D3_DC4)
			return ok && data1.Cf7 == data2.Cf7
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

type D4_DC6 struct {
	Cf9  _dafny.Int
	Cf10 _dafny.Array
	Cf11 _dafny.Int
}

func (D4_DC6) isD4() {}

func (CompanionStruct_D4_) Create_DC6_(Cf9 _dafny.Int, Cf10 _dafny.Array, Cf11 _dafny.Int) D4 {
	return D4{D4_DC6{Cf9, Cf10, Cf11}}
}

func (_this D4) Is_DC6() bool {
	_, ok := _this.Get_().(D4_DC6)
	return ok
}

type D4_DC5 struct {
	Cf8 _dafny.Set
}

func (D4_DC5) isD4() {}

func (CompanionStruct_D4_) Create_DC5_(Cf8 _dafny.Set) D4 {
	return D4{D4_DC5{Cf8}}
}

func (_this D4) Is_DC5() bool {
	_, ok := _this.Get_().(D4_DC5)
	return ok
}

type D4_DC7 struct {
	Cf12 D4
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf12 D4) D4 {
	return D4{D4_DC7{Cf12}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC6_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D4) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D4_DC6).Cf9
}

func (_this D4) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D4_DC6).Cf10
}

func (_this D4) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D4_DC6).Cf11
}

func (_this D4) Dtor_cf8() _dafny.Set {
	return _this.Get_().(D4_DC5).Cf8
}

func (_this D4) Dtor_cf12() D4 {
	return _this.Get_().(D4_DC7).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC6:
		{
			return "D4.DC6" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D4_DC5:
		{
			return "D4.DC5" + "(" + _dafny.String(data.Cf8) + ")"
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
	case D4_DC6:
		{
			data2, ok := other.Get_().(D4_DC6)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D4_DC5:
		{
			data2, ok := other.Get_().(D4_DC5)
			return ok && data1.Cf8.Equals(data2.Cf8)
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D5_DC9 struct {
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 _dafny.MultiSet
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf14 bool, Cf15 _dafny.Int, Cf16 _dafny.MultiSet) D5 {
	return D5{D5_DC9{Cf14, Cf15, Cf16}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

type D5_DC10 struct {
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 _dafny.Map
	Cf20 bool
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf17 bool, Cf18 _dafny.Int, Cf19 _dafny.Map, Cf20 bool) D5 {
	return D5{D5_DC10{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

type D5_DC8 struct {
	Cf13 _dafny.MultiSet
}

func (D5_DC8) isD5() {}

func (CompanionStruct_D5_) Create_DC8_(Cf13 _dafny.MultiSet) D5 {
	return D5{D5_DC8{Cf13}}
}

func (_this D5) Is_DC8() bool {
	_, ok := _this.Get_().(D5_DC8)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC9_(false, _dafny.Zero, _dafny.EmptyMultiSet)
}

func (_this D5) Dtor_cf14() bool {
	return _this.Get_().(D5_DC9).Cf14
}

func (_this D5) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D5_DC9).Cf15
}

func (_this D5) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D5_DC9).Cf16
}

func (_this D5) Dtor_cf17() bool {
	return _this.Get_().(D5_DC10).Cf17
}

func (_this D5) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D5_DC10).Cf18
}

func (_this D5) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D5_DC10).Cf19
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC10).Cf20
}

func (_this D5) Dtor_cf13() _dafny.MultiSet {
	return _this.Get_().(D5_DC8).Cf13
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC9:
		{
			return "D5.DC9" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC8:
		{
			return "D5.DC8" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC9:
		{
			data2, ok := other.Get_().(D5_DC9)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Equals(data2.Cf16)
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Equals(data2.Cf19) && data1.Cf20 == data2.Cf20
		}
	case D5_DC8:
		{
			data2, ok := other.Get_().(D5_DC8)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf22 bool
	Cf23 bool
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf22 bool, Cf23 bool) D6 {
	return D6{D6_DC12{Cf22, Cf23}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC13 struct {
	Cf24 bool
	Cf25 _dafny.Int
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf24 bool, Cf25 _dafny.Int) D6 {
	return D6{D6_DC13{Cf24, Cf25}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC11 struct {
	Cf21 _dafny.Sequence
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf21 _dafny.Sequence) D6 {
	return D6{D6_DC11{Cf21}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC12_(false, false)
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC12).Cf22
}

func (_this D6) Dtor_cf23() bool {
	return _this.Get_().(D6_DC12).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC13).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf25
}

func (_this D6) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D6_DC11).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + data.Cf21.VerbatimString(true) + ")"
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
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D7_DC15 struct {
	Cf27 _dafny.Int
	Cf28 _dafny.CodePoint
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf27 _dafny.Int, Cf28 _dafny.CodePoint) D7 {
	return D7{D7_DC15{Cf27, Cf28}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC14 struct {
	Cf26 _dafny.Sequence
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf26 _dafny.Sequence) D7 {
	return D7{D7_DC14{Cf26}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC15_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D7) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf27
}

func (_this D7) Dtor_cf28() _dafny.CodePoint {
	return _this.Get_().(D7_DC15).Cf28
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC14).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D8_DC17 struct {
	Cf30 bool
	Cf31 *C1
	Cf32 _dafny.Int
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf30 bool, Cf31 *C1, Cf32 _dafny.Int) D8 {
	return D8{D8_DC17{Cf30, Cf31, Cf32}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC16 struct {
	Cf29 *C1
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf29 *C1) D8 {
	return D8{D8_DC16{Cf29}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC17_(false, (*C1)(nil), _dafny.Zero)
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC17).Cf30
}

func (_this D8) Dtor_cf31() *C1 {
	return _this.Get_().(D8_DC17).Cf31
}

func (_this D8) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf32
}

func (_this D8) Dtor_cf29() *C1 {
	return _this.Get_().(D8_DC16).Cf29
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0
		}
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf29 == data2.Cf29
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
	Cf33 T2
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf33 T2) D9 {
	return D9{D9_DC18{Cf33}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

func (CompanionStruct_D9_) Default() T2 {
	return (T2)(nil)
}

func (_this D9) Dtor_cf33() T2 {
	return _this.Get_().(D9_DC18).Cf33
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && _dafny.AreEqual(data1.Cf33, data2.Cf33)
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

type D10_DC19 struct {
	Cf34 _dafny.MultiSet
}

func (D10_DC19) isD10() {}

func (CompanionStruct_D10_) Create_DC19_(Cf34 _dafny.MultiSet) D10 {
	return D10{D10_DC19{Cf34}}
}

func (_this D10) Is_DC19() bool {
	_, ok := _this.Get_().(D10_DC19)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D10) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D10_DC19).Cf34
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC19:
		{
			return "D10.DC19" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC19:
		{
			data2, ok := other.Get_().(D10_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D11_DC21 struct {
	Cf36 _dafny.CodePoint
	Cf37 _dafny.Int
	Cf38 _dafny.Int
}

func (D11_DC21) isD11() {}

func (CompanionStruct_D11_) Create_DC21_(Cf36 _dafny.CodePoint, Cf37 _dafny.Int, Cf38 _dafny.Int) D11 {
	return D11{D11_DC21{Cf36, Cf37, Cf38}}
}

func (_this D11) Is_DC21() bool {
	_, ok := _this.Get_().(D11_DC21)
	return ok
}

type D11_DC22 struct {
	Cf39 bool
}

func (D11_DC22) isD11() {}

func (CompanionStruct_D11_) Create_DC22_(Cf39 bool) D11 {
	return D11{D11_DC22{Cf39}}
}

func (_this D11) Is_DC22() bool {
	_, ok := _this.Get_().(D11_DC22)
	return ok
}

type D11_DC20 struct {
	Cf35 _dafny.Array
}

func (D11_DC20) isD11() {}

func (CompanionStruct_D11_) Create_DC20_(Cf35 _dafny.Array) D11 {
	return D11{D11_DC20{Cf35}}
}

func (_this D11) Is_DC20() bool {
	_, ok := _this.Get_().(D11_DC20)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC21_(_dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D11) Dtor_cf36() _dafny.CodePoint {
	return _this.Get_().(D11_DC21).Cf36
}

func (_this D11) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D11_DC21).Cf37
}

func (_this D11) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D11_DC21).Cf38
}

func (_this D11) Dtor_cf39() bool {
	return _this.Get_().(D11_DC22).Cf39
}

func (_this D11) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D11_DC20).Cf35
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC21:
		{
			return "D11.DC21" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D11_DC22:
		{
			return "D11.DC22" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D11_DC20:
		{
			return "D11.DC20" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC21:
		{
			data2, ok := other.Get_().(D11_DC21)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D11_DC22:
		{
			data2, ok := other.Get_().(D11_DC22)
			return ok && data1.Cf39 == data2.Cf39
		}
	case D11_DC20:
		{
			data2, ok := other.Get_().(D11_DC20)
			return ok && data1.Cf35 == data2.Cf35
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
	Cf40 _dafny.Array
}

func (D12_DC23) isD12() {}

func (CompanionStruct_D12_) Create_DC23_(Cf40 _dafny.Array) D12 {
	return D12{D12_DC23{Cf40}}
}

func (_this D12) Is_DC23() bool {
	_, ok := _this.Get_().(D12_DC23)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D12) Dtor_cf40() _dafny.Array {
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
			return ok && data1.Cf40 == data2.Cf40
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

type D13_DC24 struct {
	Cf41 _dafny.Set
}

func (D13_DC24) isD13() {}

func (CompanionStruct_D13_) Create_DC24_(Cf41 _dafny.Set) D13 {
	return D13{D13_DC24{Cf41}}
}

func (_this D13) Is_DC24() bool {
	_, ok := _this.Get_().(D13_DC24)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D13) Dtor_cf41() _dafny.Set {
	return _this.Get_().(D13_DC24).Cf41
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D13_DC24:
		{
			data2, ok := other.Get_().(D13_DC24)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D14_DC26 struct {
	Cf43 D0
	Cf44 _dafny.Set
}

func (D14_DC26) isD14() {}

func (CompanionStruct_D14_) Create_DC26_(Cf43 D0, Cf44 _dafny.Set) D14 {
	return D14{D14_DC26{Cf43, Cf44}}
}

func (_this D14) Is_DC26() bool {
	_, ok := _this.Get_().(D14_DC26)
	return ok
}

type D14_DC27 struct {
	Cf45 bool
}

func (D14_DC27) isD14() {}

func (CompanionStruct_D14_) Create_DC27_(Cf45 bool) D14 {
	return D14{D14_DC27{Cf45}}
}

func (_this D14) Is_DC27() bool {
	_, ok := _this.Get_().(D14_DC27)
	return ok
}

type D14_DC28 struct {
}

func (D14_DC28) isD14() {}

func (CompanionStruct_D14_) Create_DC28_() D14 {
	return D14{D14_DC28{}}
}

func (_this D14) Is_DC28() bool {
	_, ok := _this.Get_().(D14_DC28)
	return ok
}

type D14_DC25 struct {
	Cf42 _dafny.Array
}

func (D14_DC25) isD14() {}

func (CompanionStruct_D14_) Create_DC25_(Cf42 _dafny.Array) D14 {
	return D14{D14_DC25{Cf42}}
}

func (_this D14) Is_DC25() bool {
	_, ok := _this.Get_().(D14_DC25)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC26_(Companion_D0_.Default(), _dafny.EmptySet)
}

func (_this D14) Dtor_cf43() D0 {
	return _this.Get_().(D14_DC26).Cf43
}

func (_this D14) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D14_DC26).Cf44
}

func (_this D14) Dtor_cf45() bool {
	return _this.Get_().(D14_DC27).Cf45
}

func (_this D14) Dtor_cf42() _dafny.Array {
	return _this.Get_().(D14_DC25).Cf42
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC26:
		{
			return "D14.DC26" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D14_DC27:
		{
			return "D14.DC27" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D14_DC28:
		{
			return "D14.DC28"
		}
	case D14_DC25:
		{
			return "D14.DC25" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC26:
		{
			data2, ok := other.Get_().(D14_DC26)
			return ok && data1.Cf43.Equals(data2.Cf43) && data1.Cf44.Equals(data2.Cf44)
		}
	case D14_DC27:
		{
			data2, ok := other.Get_().(D14_DC27)
			return ok && data1.Cf45 == data2.Cf45
		}
	case D14_DC28:
		{
			_, ok := other.Get_().(D14_DC28)
			return ok
		}
	case D14_DC25:
		{
			data2, ok := other.Get_().(D14_DC25)
			return ok && data1.Cf42 == data2.Cf42
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

type D15_DC30 struct {
	Cf47 bool
}

func (D15_DC30) isD15() {}

func (CompanionStruct_D15_) Create_DC30_(Cf47 bool) D15 {
	return D15{D15_DC30{Cf47}}
}

func (_this D15) Is_DC30() bool {
	_, ok := _this.Get_().(D15_DC30)
	return ok
}

type D15_DC29 struct {
	Cf46 _dafny.Map
}

func (D15_DC29) isD15() {}

func (CompanionStruct_D15_) Create_DC29_(Cf46 _dafny.Map) D15 {
	return D15{D15_DC29{Cf46}}
}

func (_this D15) Is_DC29() bool {
	_, ok := _this.Get_().(D15_DC29)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC30_(false)
}

func (_this D15) Dtor_cf47() bool {
	return _this.Get_().(D15_DC30).Cf47
}

func (_this D15) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D15_DC29).Cf46
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC30:
		{
			return "D15.DC30" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D15_DC29:
		{
			return "D15.DC29" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC30:
		{
			data2, ok := other.Get_().(D15_DC30)
			return ok && data1.Cf47 == data2.Cf47
		}
	case D15_DC29:
		{
			data2, ok := other.Get_().(D15_DC29)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D16_DC31 struct {
	Cf48 _dafny.Set
}

func (D16_DC31) isD16() {}

func (CompanionStruct_D16_) Create_DC31_(Cf48 _dafny.Set) D16 {
	return D16{D16_DC31{Cf48}}
}

func (_this D16) Is_DC31() bool {
	_, ok := _this.Get_().(D16_DC31)
	return ok
}

func (CompanionStruct_D16_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D16) Dtor_cf48() _dafny.Set {
	return _this.Get_().(D16_DC31).Cf48
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC31:
		{
			return "D16.DC31" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC31:
		{
			data2, ok := other.Get_().(D16_DC31)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

type D17_DC32 struct {
	Cf49 _dafny.Set
}

func (D17_DC32) isD17() {}

func (CompanionStruct_D17_) Create_DC32_(Cf49 _dafny.Set) D17 {
	return D17{D17_DC32{Cf49}}
}

func (_this D17) Is_DC32() bool {
	_, ok := _this.Get_().(D17_DC32)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D17) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D17_DC32).Cf49
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC32:
		{
			return "D17.DC32" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC32:
		{
			data2, ok := other.Get_().(D17_DC32)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D18_DC34 struct {
	Cf51 bool
	Cf52 bool
}

func (D18_DC34) isD18() {}

func (CompanionStruct_D18_) Create_DC34_(Cf51 bool, Cf52 bool) D18 {
	return D18{D18_DC34{Cf51, Cf52}}
}

func (_this D18) Is_DC34() bool {
	_, ok := _this.Get_().(D18_DC34)
	return ok
}

type D18_DC33 struct {
	Cf50 _dafny.Map
}

func (D18_DC33) isD18() {}

func (CompanionStruct_D18_) Create_DC33_(Cf50 _dafny.Map) D18 {
	return D18{D18_DC33{Cf50}}
}

func (_this D18) Is_DC33() bool {
	_, ok := _this.Get_().(D18_DC33)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC34_(false, false)
}

func (_this D18) Dtor_cf51() bool {
	return _this.Get_().(D18_DC34).Cf51
}

func (_this D18) Dtor_cf52() bool {
	return _this.Get_().(D18_DC34).Cf52
}

func (_this D18) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D18_DC33).Cf50
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC34:
		{
			return "D18.DC34" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D18_DC33:
		{
			return "D18.DC33" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC34:
		{
			data2, ok := other.Get_().(D18_DC34)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52 == data2.Cf52
		}
	case D18_DC33:
		{
			data2, ok := other.Get_().(D18_DC33)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D19_DC35 struct {
	Cf53 T0
}

func (D19_DC35) isD19() {}

func (CompanionStruct_D19_) Create_DC35_(Cf53 T0) D19 {
	return D19{D19_DC35{Cf53}}
}

func (_this D19) Is_DC35() bool {
	_, ok := _this.Get_().(D19_DC35)
	return ok
}

func (CompanionStruct_D19_) Default() T0 {
	return (T0)(nil)
}

func (_this D19) Dtor_cf53() T0 {
	return _this.Get_().(D19_DC35).Cf53
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC35:
		{
			return "D19.DC35" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC35:
		{
			data2, ok := other.Get_().(D19_DC35)
			return ok && _dafny.AreEqual(data1.Cf53, data2.Cf53)
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

type D20_DC37 struct {
	Cf55 _dafny.CodePoint
	Cf56 bool
}

func (D20_DC37) isD20() {}

func (CompanionStruct_D20_) Create_DC37_(Cf55 _dafny.CodePoint, Cf56 bool) D20 {
	return D20{D20_DC37{Cf55, Cf56}}
}

func (_this D20) Is_DC37() bool {
	_, ok := _this.Get_().(D20_DC37)
	return ok
}

type D20_DC38 struct {
	Cf57 *C7
}

func (D20_DC38) isD20() {}

func (CompanionStruct_D20_) Create_DC38_(Cf57 *C7) D20 {
	return D20{D20_DC38{Cf57}}
}

func (_this D20) Is_DC38() bool {
	_, ok := _this.Get_().(D20_DC38)
	return ok
}

type D20_DC36 struct {
	Cf54 _dafny.MultiSet
}

func (D20_DC36) isD20() {}

func (CompanionStruct_D20_) Create_DC36_(Cf54 _dafny.MultiSet) D20 {
	return D20{D20_DC36{Cf54}}
}

func (_this D20) Is_DC36() bool {
	_, ok := _this.Get_().(D20_DC36)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC37_(_dafny.CodePoint('D'), false)
}

func (_this D20) Dtor_cf55() _dafny.CodePoint {
	return _this.Get_().(D20_DC37).Cf55
}

func (_this D20) Dtor_cf56() bool {
	return _this.Get_().(D20_DC37).Cf56
}

func (_this D20) Dtor_cf57() *C7 {
	return _this.Get_().(D20_DC38).Cf57
}

func (_this D20) Dtor_cf54() _dafny.MultiSet {
	return _this.Get_().(D20_DC36).Cf54
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC37:
		{
			return "D20.DC37" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D20_DC38:
		{
			return "D20.DC38" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D20_DC36:
		{
			return "D20.DC36" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC37:
		{
			data2, ok := other.Get_().(D20_DC37)
			return ok && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56
		}
	case D20_DC38:
		{
			data2, ok := other.Get_().(D20_DC38)
			return ok && data1.Cf57 == data2.Cf57
		}
	case D20_DC36:
		{
			data2, ok := other.Get_().(D20_DC36)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D21_DC40 struct {
	Cf59 *C3
}

func (D21_DC40) isD21() {}

func (CompanionStruct_D21_) Create_DC40_(Cf59 *C3) D21 {
	return D21{D21_DC40{Cf59}}
}

func (_this D21) Is_DC40() bool {
	_, ok := _this.Get_().(D21_DC40)
	return ok
}

type D21_DC41 struct {
	Cf60 bool
}

func (D21_DC41) isD21() {}

func (CompanionStruct_D21_) Create_DC41_(Cf60 bool) D21 {
	return D21{D21_DC41{Cf60}}
}

func (_this D21) Is_DC41() bool {
	_, ok := _this.Get_().(D21_DC41)
	return ok
}

type D21_DC39 struct {
	Cf58 *C4
}

func (D21_DC39) isD21() {}

func (CompanionStruct_D21_) Create_DC39_(Cf58 *C4) D21 {
	return D21{D21_DC39{Cf58}}
}

func (_this D21) Is_DC39() bool {
	_, ok := _this.Get_().(D21_DC39)
	return ok
}

type D21_DC42 struct {
	Cf61 D21
}

func (D21_DC42) isD21() {}

func (CompanionStruct_D21_) Create_DC42_(Cf61 D21) D21 {
	return D21{D21_DC42{Cf61}}
}

func (_this D21) Is_DC42() bool {
	_, ok := _this.Get_().(D21_DC42)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC40_((*C3)(nil))
}

func (_this D21) Dtor_cf59() *C3 {
	return _this.Get_().(D21_DC40).Cf59
}

func (_this D21) Dtor_cf60() bool {
	return _this.Get_().(D21_DC41).Cf60
}

func (_this D21) Dtor_cf58() *C4 {
	return _this.Get_().(D21_DC39).Cf58
}

func (_this D21) Dtor_cf61() D21 {
	return _this.Get_().(D21_DC42).Cf61
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC40:
		{
			return "D21.DC40" + "(" + _dafny.String(data.Cf59) + ")"
		}
	case D21_DC41:
		{
			return "D21.DC41" + "(" + _dafny.String(data.Cf60) + ")"
		}
	case D21_DC39:
		{
			return "D21.DC39" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D21_DC42:
		{
			return "D21.DC42" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC40:
		{
			data2, ok := other.Get_().(D21_DC40)
			return ok && data1.Cf59 == data2.Cf59
		}
	case D21_DC41:
		{
			data2, ok := other.Get_().(D21_DC41)
			return ok && data1.Cf60 == data2.Cf60
		}
	case D21_DC39:
		{
			data2, ok := other.Get_().(D21_DC39)
			return ok && data1.Cf58 == data2.Cf58
		}
	case D21_DC42:
		{
			data2, ok := other.Get_().(D21_DC42)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D22_DC44 struct {
	Cf63 _dafny.Set
}

func (D22_DC44) isD22() {}

func (CompanionStruct_D22_) Create_DC44_(Cf63 _dafny.Set) D22 {
	return D22{D22_DC44{Cf63}}
}

func (_this D22) Is_DC44() bool {
	_, ok := _this.Get_().(D22_DC44)
	return ok
}

type D22_DC45 struct {
	Cf64 _dafny.Map
}

func (D22_DC45) isD22() {}

func (CompanionStruct_D22_) Create_DC45_(Cf64 _dafny.Map) D22 {
	return D22{D22_DC45{Cf64}}
}

func (_this D22) Is_DC45() bool {
	_, ok := _this.Get_().(D22_DC45)
	return ok
}

type D22_DC43 struct {
	Cf62 _dafny.Set
}

func (D22_DC43) isD22() {}

func (CompanionStruct_D22_) Create_DC43_(Cf62 _dafny.Set) D22 {
	return D22{D22_DC43{Cf62}}
}

func (_this D22) Is_DC43() bool {
	_, ok := _this.Get_().(D22_DC43)
	return ok
}

type D22_DC46 struct {
	Cf65 D22
}

func (D22_DC46) isD22() {}

func (CompanionStruct_D22_) Create_DC46_(Cf65 D22) D22 {
	return D22{D22_DC46{Cf65}}
}

func (_this D22) Is_DC46() bool {
	_, ok := _this.Get_().(D22_DC46)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC44_(_dafny.EmptySet)
}

func (_this D22) Dtor_cf63() _dafny.Set {
	return _this.Get_().(D22_DC44).Cf63
}

func (_this D22) Dtor_cf64() _dafny.Map {
	return _this.Get_().(D22_DC45).Cf64
}

func (_this D22) Dtor_cf62() _dafny.Set {
	return _this.Get_().(D22_DC43).Cf62
}

func (_this D22) Dtor_cf65() D22 {
	return _this.Get_().(D22_DC46).Cf65
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC44:
		{
			return "D22.DC44" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D22_DC45:
		{
			return "D22.DC45" + "(" + _dafny.String(data.Cf64) + ")"
		}
	case D22_DC43:
		{
			return "D22.DC43" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D22_DC46:
		{
			return "D22.DC46" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC44:
		{
			data2, ok := other.Get_().(D22_DC44)
			return ok && data1.Cf63.Equals(data2.Cf63)
		}
	case D22_DC45:
		{
			data2, ok := other.Get_().(D22_DC45)
			return ok && data1.Cf64.Equals(data2.Cf64)
		}
	case D22_DC43:
		{
			data2, ok := other.Get_().(D22_DC43)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	case D22_DC46:
		{
			data2, ok := other.Get_().(D22_DC46)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D23_DC47 struct {
	Cf66 _dafny.Map
}

func (D23_DC47) isD23() {}

func (CompanionStruct_D23_) Create_DC47_(Cf66 _dafny.Map) D23 {
	return D23{D23_DC47{Cf66}}
}

func (_this D23) Is_DC47() bool {
	_, ok := _this.Get_().(D23_DC47)
	return ok
}

func (CompanionStruct_D23_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D23) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D23_DC47).Cf66
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC47:
		{
			return "D23.DC47" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC47:
		{
			data2, ok := other.Get_().(D23_DC47)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D24_DC49 struct {
	Cf68 _dafny.Int
	Cf69 _dafny.Int
	Cf70 _dafny.Int
}

func (D24_DC49) isD24() {}

func (CompanionStruct_D24_) Create_DC49_(Cf68 _dafny.Int, Cf69 _dafny.Int, Cf70 _dafny.Int) D24 {
	return D24{D24_DC49{Cf68, Cf69, Cf70}}
}

func (_this D24) Is_DC49() bool {
	_, ok := _this.Get_().(D24_DC49)
	return ok
}

type D24_DC48 struct {
	Cf67 _dafny.Map
}

func (D24_DC48) isD24() {}

func (CompanionStruct_D24_) Create_DC48_(Cf67 _dafny.Map) D24 {
	return D24{D24_DC48{Cf67}}
}

func (_this D24) Is_DC48() bool {
	_, ok := _this.Get_().(D24_DC48)
	return ok
}

type D24_DC50 struct {
	Cf71 D24
}

func (D24_DC50) isD24() {}

func (CompanionStruct_D24_) Create_DC50_(Cf71 D24) D24 {
	return D24{D24_DC50{Cf71}}
}

func (_this D24) Is_DC50() bool {
	_, ok := _this.Get_().(D24_DC50)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC49_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D24) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D24_DC49).Cf68
}

func (_this D24) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D24_DC49).Cf69
}

func (_this D24) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D24_DC49).Cf70
}

func (_this D24) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D24_DC48).Cf67
}

func (_this D24) Dtor_cf71() D24 {
	return _this.Get_().(D24_DC50).Cf71
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC49:
		{
			return "D24.DC49" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D24_DC48:
		{
			return "D24.DC48" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D24_DC50:
		{
			return "D24.DC50" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC49:
		{
			data2, ok := other.Get_().(D24_DC49)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69.Cmp(data2.Cf69) == 0 && data1.Cf70.Cmp(data2.Cf70) == 0
		}
	case D24_DC48:
		{
			data2, ok := other.Get_().(D24_DC48)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	case D24_DC50:
		{
			data2, ok := other.Get_().(D24_DC50)
			return ok && data1.Cf71.Equals(data2.Cf71)
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
	Fm0(globalState *GlobalState) _dafny.Map
	Fm1(p0 bool, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int)
	M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
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
	Fm0(globalState *GlobalState) _dafny.Map
	Fm1(p0 bool, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int)
	M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
	M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int)
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
	Fm0(globalState *GlobalState) _dafny.Map
	Fm1(p0 bool, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int)
	M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
	M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int)
	F2() _dafny.Int
	F2_set_(value _dafny.Int)
	Fm8(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool
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
	dummy byte
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

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

func (_this *GlobalState) Ctor__() {
	{
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F4 bool
	F5 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F4 = false
	_this.F5 = false
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

func (_this *C0) Ctor__(f4 bool, f5 bool) {
	{
		(_this).F4 = f4
		(_this).F5 = f5
	}
}
func (_this *C0) Fm14(p0 bool, p1 bool, p2 _dafny.Set, globalState *GlobalState) bool {
	{
		return _this.F4
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f3 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f3 = _dafny.EmptySeq
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

func (_this *C1) Ctor__(f3 _dafny.Sequence) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C1) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_457_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32(((_this).F3()).Cardinality())
		})), _dafny.IntOfUint32(((_this).F3()).Cardinality()))).Merge((func() _dafny.Map {
			if false {
				return func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter18 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}(func(_458_i1 _dafny.Int) _dafny.Int {
						return (_dafny.MultiSetOf(false)).Cardinality()
					})), _dafny.IntOfInt64(352))).Keys().Elements()); ; {
						_compr_17, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _459_v0 _dafny.Sequence
						_459_v0 = interface{}(_compr_17).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}(func(_458_i1 _dafny.Int) _dafny.Int {
							return (_dafny.MultiSetOf(false)).Cardinality()
						})), _dafny.IntOfInt64(352))).Contains(_459_v0) {
							_coll17.Add(_459_v0, _dafny.IntOfInt64(894))
						}
					}
					return _coll17.ToMap()
				}()
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(266)), _dafny.IntOfInt64(-986))
		})())
	}
}
func (_this *C1) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true))).Cardinality())).Plus(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if !(false) {
				return _dafny.SeqOf(false, false)
			}
			return _dafny.SeqOf(!(!(false)), false)
		})()).Cardinality()))
	}
}
func (_this *C1) Fm11(globalState *GlobalState) bool {
	{
		return (_dafny.SetOf(true)).IsDisjointFrom((_dafny.SetOf(true)).Intersection(_dafny.SetOf(false)))
	}
}
func (_this *C1) Fm12(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-906)), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-299)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ob")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(730))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_460_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality())))
	}
}
func (_this *C1) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _461_v0 _dafny.Array
		_ = _461_v0
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw99
		_461_v0 = _nw99
		var _462_v1 _dafny.Array
		_ = _462_v1
		var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw100
		_462_v1 = _nw100
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_461_v0), 0))
		_ = _index80
		(_461_v0).ArraySet1(_462_v1, (_index80).Int())
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_461_v0), 0))
		_ = _index81
		(_461_v0).ArraySet1(_462_v1, (_index81).Int())
		var _463_v2 _dafny.CodePoint
		_ = _463_v2
		_463_v2 = _dafny.CodePoint('t')
		var _464_v3 _dafny.Int
		_ = _464_v3
		_464_v3 = _dafny.IntOfInt64(469)
		var _465_v4 D0
		_ = _465_v4
		_465_v4 = Companion_D0_.Create_DC1_(p3, p3, _463_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _464_v3)).Cardinality())
		var _pat_let_tv8 = _464_v3
		_ = _pat_let_tv8
		var _pat_let_tv9 = _464_v3
		_ = _pat_let_tv9
		r0 = func(_source8 D0) _dafny.Int {
			if _source8.Is_DC1() {
				var _466___mcc_h0 bool = _source8.Get_().(D0_DC1).Cf1
				_ = _466___mcc_h0
				var _467___mcc_h1 bool = _source8.Get_().(D0_DC1).Cf2
				_ = _467___mcc_h1
				var _468___mcc_h2 _dafny.CodePoint = _source8.Get_().(D0_DC1).Cf3
				_ = _468___mcc_h2
				var _469___mcc_h3 _dafny.Int = _source8.Get_().(D0_DC1).Cf4
				_ = _469___mcc_h3
				var _470_cf4 _dafny.Int = _469___mcc_h3
				_ = _470_cf4
				var _471_cf3 _dafny.CodePoint = _468___mcc_h2
				_ = _471_cf3
				var _472_cf2 bool = _467___mcc_h1
				_ = _472_cf2
				var _473_cf1 bool = _466___mcc_h0
				_ = _473_cf1
				return _pat_let_tv8
			} else {
				var _474___mcc_h4 bool = _source8.Get_().(D0_DC0).Cf0
				_ = _474___mcc_h4
				var _475_cf0 bool = _474___mcc_h4
				_ = _475_cf0
				return _pat_let_tv9
			}
		}(_465_v4)
		var _476_v5 bool
		_ = _476_v5
		_476_v5 = false
		var _477_v6 _dafny.Array
		_ = _477_v6
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_8
		var _nw101 _dafny.Array
		_ = _nw101
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw101 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.MultiSet = (func(_478_p3 bool, _479_v5 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_480_i0 _dafny.Int) _dafny.MultiSet {
					return (_dafny.MultiSetOf(_478_p3, _479_v5)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_478_p3)))
				}
			})(p3, _476_v5)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw101 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw101).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw101).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_477_v6 = _nw101
		var _481_v7 _dafny.MultiSet
		_ = _481_v7
		_481_v7 = _dafny.MultiSetOf(p2, _476_v5)
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_477_v6), 0))
		_ = _index82
		(_477_v6).ArraySet1(_481_v7, (_index82).Int())
		var _482_v8 _dafny.Map
		_ = _482_v8
		_482_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm13(globalState))
		var _483_v9 _dafny.Map
		_ = _483_v9
		_483_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v3, p3)
		var _pat_let_tv10 = p3
		_ = _pat_let_tv10
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_477_v6), 0))
		_ = _index83
		var _rhs60 _dafny.CodePoint = _463_v2
		_ = _rhs60
		var _rhs61 bool = func(_source9 D0) bool {
			if _source9.Is_DC1() {
				var _484___mcc_h5 bool = _source9.Get_().(D0_DC1).Cf1
				_ = _484___mcc_h5
				var _485___mcc_h6 bool = _source9.Get_().(D0_DC1).Cf2
				_ = _485___mcc_h6
				var _486___mcc_h7 _dafny.CodePoint = _source9.Get_().(D0_DC1).Cf3
				_ = _486___mcc_h7
				var _487___mcc_h8 _dafny.Int = _source9.Get_().(D0_DC1).Cf4
				_ = _487___mcc_h8
				var _488_cf4 _dafny.Int = _487___mcc_h8
				_ = _488_cf4
				var _489_cf3 _dafny.CodePoint = _486___mcc_h7
				_ = _489_cf3
				var _490_cf2 bool = _485___mcc_h6
				_ = _490_cf2
				var _491_cf1 bool = _484___mcc_h5
				_ = _491_cf1
				return _pat_let_tv10
			} else {
				var _492___mcc_h9 bool = _source9.Get_().(D0_DC0).Cf0
				_ = _492___mcc_h9
				var _493_cf0 bool = _492___mcc_h9
				_ = _493_cf0
				return _493_cf0
			}
		}(Companion_D0_.Create_DC1_(p3, _476_v5, _463_v2, _464_v3))
		_ = _rhs61
		var _rhs62 _dafny.MultiSet = (_481_v7).Union(_481_v7)
		_ = _rhs62
		var _rhs63 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
			if (_482_v8).Contains(_476_v5) {
				return (_482_v8).Get(_476_v5).(_dafny.Map)
			}
			return _483_v9
		})())).Cardinality())).Times((_dafny.Zero).Minus(_464_v3))
		_ = _rhs63
		var _lhs32 _dafny.Array = _477_v6
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_477_v6), 0))
		_ = _lhs33
		_463_v2 = _rhs60
		_476_v5 = _rhs61
		(_lhs32).ArraySet1(_rhs62, (_lhs33).Int())
		r1 = _rhs63
		var _494_v10 _dafny.Sequence
		_ = _494_v10
		_494_v10 = _dafny.SeqOf(_dafny.IntOfInt64(18), _464_v3, Companion_Default___.SafeDivisionInt(_464_v3, _464_v3))
		var _495_v11 _dafny.Sequence
		_ = _495_v11
		_495_v11 = _dafny.SeqOf(_476_v5, true)
		_494_v10 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).Fm12(p3, _476_v5, (_this).Fm1((_495_v11).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _464_v3)).Cardinality(), _dafny.IntOfUint32((_495_v11).Cardinality()))).Uint32()).(bool), globalState), _476_v5, globalState), (Companion_Default___.SafeIndex(_464_v3, _dafny.IntOfUint32(((_this).Fm12(p3, _476_v5, (_this).Fm1((_495_v11).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _464_v3)).Cardinality(), _dafny.IntOfUint32((_495_v11).Cardinality()))).Uint32()).(bool), globalState), _476_v5, globalState)).Cardinality()))).Uint32(), _464_v3), (_this).Fm12((_this).Fm11(globalState), p3, _464_v3, _476_v5, globalState))
		var _496_v12 *C0
		_ = _496_v12
		var _nw102 *C0 = New_C0_()
		_ = _nw102
		_nw102.Ctor__((_464_v3).Cmp(_464_v3) >= 0, p2)
		_496_v12 = _nw102
		if p2 {
			var _497_v13 _dafny.Map
			_ = _497_v13
			_497_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_476_v5, true)
			var _source10 _dafny.Map = (_497_v13).Merge(_497_v13)
			_ = _source10
			var _498___mcc_h10 _dafny.Map = _source10
			_ = _498___mcc_h10
			var _499_cf5 _dafny.Map = _498___mcc_h10
			_ = _499_cf5
			var _500_v14 _dafny.Map
			_ = _500_v14
			_500_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v3, _464_v3)
			_464_v3 = Companion_Default___.SafeModuloInt(_464_v3, (func() _dafny.Int {
				if (_500_v14).Contains(_464_v3) {
					return (_500_v14).Get(_464_v3).(_dafny.Int)
				}
				return _dafny.IntOfInt64(583)
			})())
			var _501_v15 _dafny.Map
			_ = _501_v15
			_501_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_463_v2, _dafny.UnicodeSeqOfUtf8Bytes("rmd"))
			_501_v15 = (_501_v15).Update(_463_v2, _dafny.Companion_Sequence_.Concatenate((_this).F3(), (_this).F3()))
			var _502_v16 _dafny.Set
			_ = _502_v16
			_502_v16 = _dafny.SetOf(_dafny.ArrayCastTo((_461_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_461_v0), 0))).Int())))
			var _503_v17 _dafny.Map
			_ = _503_v17
			_503_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v12.F5, (_502_v16).Cardinality())
			var _504_v18 _dafny.Map
			_ = _504_v18
			_504_v18 = (_497_v13).Update(_496_v12.F5, p2)
			var _505_v19 _dafny.Sequence
			_ = _505_v19
			_505_v19 = _dafny.SeqOf(_499_cf5, _504_v18)
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((p1), 0))
			_ = _index84
			(p1).ArraySet1(((_503_v17).Cardinality()).Plus(_dafny.IntOfUint32((_505_v19).Cardinality())), (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((p1), 0))
			_ = _index85
			(p1).ArraySet1((_464_v3).Minus((_494_v10).Select((Companion_Default___.SafeIndex(_464_v3, _dafny.IntOfUint32((_494_v10).Cardinality()))).Uint32()).(_dafny.Int)), (_index85).Int())
			r0 = _464_v3
			var _506_v20 _dafny.Array
			_ = _506_v20
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw103
			_506_v20 = _nw103
			var _507_v21 _dafny.Sequence
			_ = _507_v21
			_507_v21 = _dafny.SeqOf(p1, p1, p1)
			var _508_v22 _dafny.Array
			_ = _508_v22
			var _nwElement0_21 _dafny.Array = _506_v20
			_ = _nwElement0_21
			var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(14))
			_ = _nw104
			(_nw104).ArraySet1(_nwElement0_21, 0)
			(_nw104).ArraySet1(_506_v20, 1)
			(_nw104).ArraySet1(p1, 2)
			(_nw104).ArraySet1((_507_v21).Select((Companion_Default___.SafeIndex(_464_v3, _dafny.IntOfUint32((_507_v21).Cardinality()))).Uint32()).(_dafny.Array), 3)
			(_nw104).ArraySet1(p1, 4)
			(_nw104).ArraySet1(p1, 5)
			(_nw104).ArraySet1(_506_v20, 6)
			(_nw104).ArraySet1(p1, 7)
			(_nw104).ArraySet1(p1, 8)
			(_nw104).ArraySet1(_506_v20, 9)
			(_nw104).ArraySet1(_506_v20, 10)
			(_nw104).ArraySet1(_506_v20, 11)
			(_nw104).ArraySet1(_506_v20, 12)
			(_nw104).ArraySet1(_506_v20, 13)
			_508_v22 = _nw104
			var _509_v23 _dafny.Array
			_ = _509_v23
			_509_v23 = _508_v22
			var _rhs64 _dafny.Array = _461_v0
			_ = _rhs64
			var _rhs65 _dafny.Array = _506_v20
			_ = _rhs65
			var _rhs66 _dafny.Array = (_509_v23)
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_494_v10, _494_v10)
			_ = _rhs67
			_461_v0 = _rhs64
			_506_v20 = _rhs65
			_508_v22 = _rhs66
			_494_v10 = _rhs67
			var _510_v24 _dafny.Map
			_ = _510_v24
			_510_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v3, _dafny.ArrayCastTo((_461_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_461_v0), 0))).Int())))
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_461_v0), 0))
			_ = _index86
			(_461_v0).ArraySet1((func() _dafny.Array {
				if (_510_v24).Contains(_464_v3) {
					return (_510_v24).Get(_464_v3).(_dafny.Array)
				}
				return _462_v1
			})(), (_index86).Int())
			_476_v5 = !(((_464_v3).Times(_dafny.IntOfUint32(((_this).F3()).Cardinality()))).Cmp(Companion_Default___.SafeDivisionInt(_464_v3, _464_v3)) < 0)
			_476_v5 = !(!((_465_v4).Dtor_cf2()))
		} else {
			(_496_v12).F4 = !(_476_v5)
			(_496_v12).F5 = !(_496_v12.F5)
			r0 = _464_v3
			(_496_v12).F4 = !(_496_v12.F4)
			(_496_v12).F4 = p3
		}
		r0 = Companion_Default___.SafeModuloInt(_464_v3, (_464_v3).Plus(_464_v3))
		r1 = (_464_v3).Times(_dafny.IntOfUint32(((_this).F3()).Cardinality()))
		return r0, r1
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _511_v0 _dafny.MultiSet
		_ = _511_v0
		_511_v0 = _dafny.MultiSetOf(p2)
		if !((_511_v0).IsProperSubsetOf(((_511_v0).Update(p3, Companion_Default___.Abs(p0))).Difference(_511_v0))) {
			var _512_v1 _dafny.Sequence
			_ = _512_v1
			_512_v1 = _dafny.SeqOf(p0)
			var _513_v2 _dafny.CodePoint
			_ = _513_v2
			_513_v2 = _dafny.CodePoint('a')
			var _514_v3 _dafny.Map
			_ = _514_v3
			_514_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_v2, _dafny.CodePoint('q'))
			var _515_v4 _dafny.MultiSet
			_ = _515_v4
			_515_v4 = _dafny.MultiSetOf((_514_v3).Cardinality())
			r2 = (_512_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((func() _dafny.Set {
				var _coll18 = _dafny.NewBuilder()
				_ = _coll18
				for _iter19 := _dafny.Iterate((_515_v4).Elements()); ; {
					_compr_18, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _516_v5 _dafny.Int
					_516_v5 = interface{}(_compr_18).(_dafny.Int)
					if (_515_v4).Contains(_516_v5) {
						_coll18.Add((_516_v5).Plus(_dafny.IntOfInt64(315)))
					}
				}
				return _coll18.ToSet()
			}()).Cardinality(), p0), _dafny.IntOfUint32((_512_v1).Cardinality()))).Uint32()).(_dafny.Int)
			var _517_v6 _dafny.Set
			_ = _517_v6
			_517_v6 = _dafny.SetOf(p2)
			var _518_v7 D4
			_ = _518_v7
			_518_v7 = Companion_D4_.Create_DC5_(_517_v6)
			var _519_v8 *C0
			_ = _519_v8
			var _nw105 *C0 = New_C0_()
			_ = _nw105
			_nw105.Ctor__(p3, ((_518_v7).Dtor_cf8()).Equals(_517_v6))
			_519_v8 = _nw105
			r2 = (func() _dafny.Int {
				if _519_v8.F4 {
					return p0
				}
				return (p0).Times(p0)
			})()
			var _520_v9 *C0
			_ = _520_v9
			var _nw106 *C0 = New_C0_()
			_ = _nw106
			_nw106.Ctor__(_519_v8.F5, _519_v8.F4)
			_520_v9 = _nw106
			var _521_v10 _dafny.Sequence
			_ = _521_v10
			_521_v10 = _dafny.UnicodeSeqOfUtf8Bytes("gragg")
			_521_v10 = Companion_Default___.Fm15(globalState)
		} else {
			var _522_v11 _dafny.Array
			_ = _522_v11
			var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw107
			_522_v11 = _nw107
			var _523_v12 _dafny.Sequence
			_ = _523_v12
			_523_v12 = _dafny.SeqOf(_522_v11, _522_v11)
			var _524_v13 D4
			_ = _524_v13
			_524_v13 = Companion_D4_.Create_DC6_(Companion_Default___.SafeDivisionInt(p0, p0), (_523_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F3()).Cardinality()), _dafny.IntOfUint32((_523_v12).Cardinality()))).Uint32()).(_dafny.Array), (p0).Minus(p0))
			var _source11 D4 = _524_v13
			_ = _source11
			if _source11.Is_DC6() {
				var _525___mcc_h0 _dafny.Int = _source11.Get_().(D4_DC6).Cf9
				_ = _525___mcc_h0
				var _526___mcc_h1 _dafny.Array = _source11.Get_().(D4_DC6).Cf10
				_ = _526___mcc_h1
				var _527___mcc_h2 _dafny.Int = _source11.Get_().(D4_DC6).Cf11
				_ = _527___mcc_h2
				var _528_cf11 _dafny.Int = _527___mcc_h2
				_ = _528_cf11
				var _529_cf10 _dafny.Array = _526___mcc_h1
				_ = _529_cf10
				var _530_cf9 _dafny.Int = _525___mcc_h0
				_ = _530_cf9
				var _531_v14 _dafny.CodePoint
				_ = _531_v14
				_531_v14 = _dafny.CodePoint('q')
				var _532_v15 D0
				_ = _532_v15
				_532_v15 = Companion_D0_.Create_DC1_(p3, p2, _531_v14, _dafny.IntOfInt64(971))
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_522_v11), 0))
				_ = _index87
				(_522_v11).ArraySet1((func() _dafny.Int {
					if p2 {
						return (_532_v15).Dtor_cf4()
					}
					return _530_cf9
				})(), (_index87).Int())
				var _533_v16 _dafny.Map
				_ = _533_v16
				_533_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_528_cf11, _dafny.IntOfUint32(((_this).F3()).Cardinality()))
				var _534_v17 _dafny.Map
				_ = _534_v17
				_534_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _535_v18 _dafny.Sequence
				_ = _535_v18
				_535_v18 = _dafny.SeqOf(p2, p2, (_this).Fm11(globalState), p3)
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_522_v11), 0))
				_ = _index88
				(_522_v11).ArraySet1((func() _dafny.Int {
					if !(p3) || (!(false)) {
						return (func() _dafny.Int {
							if (_533_v16).Contains((_dafny.Zero).Minus((_534_v17).Cardinality())) {
								return (_533_v16).Get((_dafny.Zero).Minus((_534_v17).Cardinality())).(_dafny.Int)
							}
							return _530_cf9
						})()
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_535_v18, _535_v18)).Cardinality())
				})(), (_index88).Int())
				var _536_v19 _dafny.Array
				_ = _536_v19
				var _nwElement0_22 bool = p2
				_ = _nwElement0_22
				var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(13))
				_ = _nw108
				(_nw108).ArraySet1(_nwElement0_22, 0)
				(_nw108).ArraySet1(p2, 1)
				(_nw108).ArraySet1(p3, 2)
				(_nw108).ArraySet1(p3, 3)
				(_nw108).ArraySet1(p2, 4)
				(_nw108).ArraySet1(p3, 5)
				(_nw108).ArraySet1(p2, 6)
				(_nw108).ArraySet1(p2, 7)
				(_nw108).ArraySet1((_535_v18).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_dafny.IntOfInt64(254), _530_cf9, _dafny.IntOfUint32((_dafny.SeqOf((_511_v0).Cardinality())).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_535_v18).Cardinality()))).Uint32()).(bool), 8)
				(_nw108).ArraySet1(p2, 9)
				(_nw108).ArraySet1(p2, 10)
				(_nw108).ArraySet1(p3, 11)
				(_nw108).ArraySet1(p3, 12)
				_536_v19 = _nw108
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_536_v19), 0))
				_ = _index89
				(_536_v19).ArraySet1((_this).Fm11(globalState), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_536_v19), 0))
				_ = _index90
				(_536_v19).ArraySet1(false, (_index90).Int())
				var _537_v20 _dafny.Sequence
				_ = _537_v20
				_537_v20 = _dafny.SeqOf((func() _dafny.Int {
					if (_533_v16).Contains(p0) {
						return (_533_v16).Get(p0).(_dafny.Int)
					}
					return (_522_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_522_v11), 0))).Int()).(_dafny.Int)
				})())
				_537_v20 = _537_v20
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_536_v19), 0))
				_ = _index91
				(_536_v19).ArraySet1((_535_v18).Select((Companion_Default___.SafeIndex((_522_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_522_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_535_v18).Cardinality()))).Uint32()).(bool), (_index91).Int())
			} else if _source11.Is_DC5() {
				var _538___mcc_h3 _dafny.Set = _source11.Get_().(D4_DC5).Cf8
				_ = _538___mcc_h3
				var _539_cf8 _dafny.Set = _538___mcc_h3
				_ = _539_cf8
				var _540_v21 *C0
				_ = _540_v21
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__(p2, p2)
				_540_v21 = _nw109
				var _541_v22 _dafny.Sequence
				_ = _541_v22
				_541_v22 = _dafny.UnicodeSeqOfUtf8Bytes("n")
				var _542_v23 _dafny.Array
				_ = _542_v23
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_9
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = func(_543_i0 _dafny.Int) bool {
						return true
					}
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw110 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw110).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw110).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_542_v23 = _nw110
				var _544_v24 _dafny.Array
				_ = _544_v24
				var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(29))
				_ = _nw111
				_544_v24 = _nw111
				var _545_v25 _dafny.CodePoint
				_ = _545_v25
				_545_v25 = _dafny.CodePoint('n')
				var _546_v26 D0
				_ = _546_v26
				_546_v26 = Companion_D0_.Create_DC1_(_540_v21.F5, p3, _545_v25, p0)
				var _547_v27 _dafny.Set
				_ = _547_v27
				_547_v27 = _dafny.SetOf(_546_v26, Companion_D0_.Create_DC1_(_540_v21.F5, p3, _545_v25, p0), _546_v26, _546_v26)
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_544_v24), 0))
				_ = _index92
				(_544_v24).ArraySet1(_547_v27, (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_544_v24), 0))
				_ = _index93
				var _rhs68 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_548_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_549_i1 _dafny.Int) _dafny.CodePoint {
						return _548_v25
					}
				})(_545_v25))), (_this).F3())
				_ = _rhs68
				var _rhs69 _dafny.Array = _542_v23
				_ = _rhs69
				var _rhs70 bool = !(((_dafny.IntOfUint32((_541_v22).Cardinality())).Times(p0)).Cmp(p0) < 0)
				_ = _rhs70
				var _rhs71 _dafny.Set = _547_v27
				_ = _rhs71
				var _rhs72 bool = _540_v21.F5
				_ = _rhs72
				var _lhs34 *C0 = _540_v21
				_ = _lhs34
				var _lhs35 _dafny.Array = _544_v24
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_544_v24), 0))
				_ = _lhs36
				_541_v22 = _rhs68
				_542_v23 = _rhs69
				_lhs34.F4 = _rhs70
				(_lhs35).ArraySet1(_rhs71, (_lhs36).Int())
				r0 = _rhs72
				var _550_v28 _dafny.Array
				_ = _550_v28
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
				_ = _nw112
				_550_v28 = _nw112
				_550_v28 = _550_v28
				r0 = false
			} else {
				var _551___mcc_h4 D4 = _source11.Get_().(D4_DC7).Cf12
				_ = _551___mcc_h4
				var _552_cf12 D4 = _551___mcc_h4
				_ = _552_cf12
				var _553_v29 _dafny.Sequence
				_ = _553_v29
				_553_v29 = _dafny.SeqOf(p0, Companion_Default___.SafeModuloInt(p0, (_this).Fm1(p2, globalState)), Companion_Default___.SafeModuloInt(p0, p0), p0, p0)
				_553_v29 = _553_v29
				r0 = !(p2)
				var _554_v30 *C0
				_ = _554_v30
				var _nw113 *C0 = New_C0_()
				_ = _nw113
				_nw113.Ctor__(p2, !(p2))
				_554_v30 = _nw113
				var _555_v31 _dafny.Array
				_ = _555_v31
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw114
				_555_v31 = _nw114
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_555_v31), 0))
				_ = _index94
				(_555_v31).ArraySet1(_554_v30.F5, (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_555_v31), 0))
				_ = _index95
				(_555_v31).ArraySet1(p3, (_index95).Int())
			}
			r0 = !(p3)
			var _556_v32 _dafny.Sequence
			_ = _556_v32
			_556_v32 = _dafny.UnicodeSeqOfUtf8Bytes("nutlbe")
			_556_v32 = _dafny.Companion_Sequence_.Concatenate((_this).F3(), _dafny.UnicodeSeqOfUtf8Bytes("xdn"))
			r2 = p0
			var _557_v33 _dafny.Array
			_ = _557_v33
			var _nwElement0_23 bool = p2
			_ = _nwElement0_23
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(8))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_23, 0)
			(_nw115).ArraySet1(p2, 1)
			(_nw115).ArraySet1(false, 2)
			(_nw115).ArraySet1((_this).Fm11(globalState), 3)
			(_nw115).ArraySet1(p2, 4)
			(_nw115).ArraySet1((func() bool {
				if p2 {
					return false
				}
				return p2
			})(), 5)
			(_nw115).ArraySet1(p3, 6)
			(_nw115).ArraySet1(p3, 7)
			_557_v33 = _nw115
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_557_v33), 0))
			_ = _index96
			(_557_v33).ArraySet1(p3, (_index96).Int())
			var _558_v34 _dafny.Sequence
			_ = _558_v34
			_558_v34 = _dafny.SeqOf(p2, p2)
			var _559_v35 *C0
			_ = _559_v35
			var _nw116 *C0 = New_C0_()
			_ = _nw116
			_nw116.Ctor__(false, p2)
			_559_v35 = _nw116
			var _560_v36 _dafny.Sequence
			_ = _560_v36
			_560_v36 = _dafny.SeqOf(_559_v35)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_557_v33), 0))
			_ = _index97
			(_557_v33).ArraySet1(!((func() bool {
				if p2 {
					return (_558_v34).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_558_v34).Cardinality()))).Uint32()).(bool)
				}
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_560_v36, _dafny.SeqOf(_559_v35, _559_v35))
			})()), (_index97).Int())
		}
		var _561_v37 _dafny.MultiSet
		_ = _561_v37
		_561_v37 = _dafny.MultiSetOf(p0)
		var _hi2 _dafny.Int = (_561_v37).Cardinality()
		_ = _hi2
		for _562_i2 := p0; _562_i2.Cmp(_hi2) < 0; _562_i2 = _562_i2.Plus(_dafny.One) {
			var _563_v38 _dafny.Map
			_ = _563_v38
			_563_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _562_i2)
			var _564_v39 _dafny.Array
			_ = _564_v39
			var _nwElement0_24 _dafny.Int = _dafny.IntOfInt64(-17)
			_ = _nwElement0_24
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_24, 0)
			(_nw117).ArraySet1(_562_i2, 1)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()), 2)
			(_nw117).ArraySet1(_562_i2, 3)
			(_nw117).ArraySet1(p0, 4)
			(_nw117).ArraySet1(_562_i2, 5)
			(_nw117).ArraySet1((_563_v38).Cardinality(), 6)
			(_nw117).ArraySet1(p0, 7)
			(_nw117).ArraySet1(_562_i2, 8)
			(_nw117).ArraySet1(_562_i2, 9)
			(_nw117).ArraySet1(_562_i2, 10)
			(_nw117).ArraySet1(_562_i2, 11)
			_564_v39 = _nw117
			var _565_v40 _dafny.MultiSet
			_ = _565_v40
			_565_v40 = _dafny.MultiSetOf(_564_v39, _564_v39, _564_v39)
			var _566_v41 D5
			_ = _566_v41
			_566_v41 = Companion_D5_.Create_DC8_(_565_v40)
			var _567_v42 _dafny.Array
			_ = _567_v42
			var _nwElement0_25 _dafny.MultiSet = _dafny.MultiSetOf(_564_v39, _564_v39)
			_ = _nwElement0_25
			var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(21))
			_ = _nw118
			(_nw118).ArraySet1(_nwElement0_25, 0)
			(_nw118).ArraySet1(_565_v40, 1)
			(_nw118).ArraySet1(_dafny.MultiSetOf(_564_v39, _564_v39, _564_v39), 2)
			(_nw118).ArraySet1((func() _dafny.MultiSet {
				if p3 {
					return _565_v40
				}
				return _565_v40
			})(), 3)
			(_nw118).ArraySet1(_dafny.MultiSetOf(_564_v39, _564_v39, _564_v39, _564_v39), 4)
			(_nw118).ArraySet1(_565_v40, 5)
			(_nw118).ArraySet1((_dafny.MultiSetOf(_564_v39)).Intersection(_565_v40), 6)
			(_nw118).ArraySet1((_565_v40).Difference(_565_v40), 7)
			(_nw118).ArraySet1(_dafny.MultiSetOf(_564_v39, _564_v39, _564_v39), 8)
			(_nw118).ArraySet1(_565_v40, 9)
			(_nw118).ArraySet1(_565_v40, 10)
			(_nw118).ArraySet1(_dafny.MultiSetOf(_564_v39), 11)
			(_nw118).ArraySet1(_565_v40, 12)
			(_nw118).ArraySet1(_565_v40, 13)
			(_nw118).ArraySet1(_565_v40, 14)
			(_nw118).ArraySet1((_566_v41).Dtor_cf13(), 15)
			(_nw118).ArraySet1(_565_v40, 16)
			(_nw118).ArraySet1(_565_v40, 17)
			(_nw118).ArraySet1(_565_v40, 18)
			(_nw118).ArraySet1(((_565_v40).Update(_564_v39, Companion_Default___.Abs(p0))).Difference(_565_v40), 19)
			(_nw118).ArraySet1(_565_v40, 20)
			_567_v42 = _nw118
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_567_v42), 0))
			_ = _index98
			(_567_v42).ArraySet1(_565_v40, (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_567_v42), 0))
			_ = _index99
			(_567_v42).ArraySet1(_565_v40, (_index99).Int())
			var _568_v43 _dafny.Map
			_ = _568_v43
			_568_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_511_v0).Intersection(_511_v0))
			_568_v43 = (_568_v43).Merge((_568_v43).Update(p2, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm16(globalState)).Cardinality()))).Uint32(), p3))))
			var _569_v44 _dafny.CodePoint
			_ = _569_v44
			_569_v44 = _dafny.CodePoint('h')
			var _570_v45 D0
			_ = _570_v45
			_570_v45 = Companion_D0_.Create_DC1_(p2, p2, _569_v44, p0)
			var _571_v46 _dafny.MultiSet
			_ = _571_v46
			_571_v46 = _dafny.MultiSetOf(_570_v45)
			var _572_v47 _dafny.MultiSet
			_ = _572_v47
			_572_v47 = _dafny.MultiSetOf(_569_v44, _569_v44)
			var _573_v48 _dafny.Set
			_ = _573_v48
			_573_v48 = _dafny.SetOf((func() bool {
				if p3 {
					return !(p2)
				}
				return p3
			})(), (_572_v47).IsProperSubsetOf(Companion_Default___.Fm17((_this).Fm11(globalState), globalState)), true, p2)
			var _574_v49 D5
			_ = _574_v49
			_574_v49 = Companion_D5_.Create_DC9_(p2, _dafny.IntOfInt64(-599), _dafny.MultiSetOf(p2, p3))
			var _pat_let_tv11 = p2
			_ = _pat_let_tv11
			var _rhs73 _dafny.MultiSet = _dafny.MultiSetOf(func(_pat_let12_0 D0) D0 {
				return func(_575_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let13_0 bool) D0 {
						return func(_576_dt__update_hcf2_h0 bool) D0 {
							return Companion_D0_.Create_DC1_((_575_dt__update__tmp_h0).Dtor_cf1(), _576_dt__update_hcf2_h0, (_575_dt__update__tmp_h0).Dtor_cf3(), (_575_dt__update__tmp_h0).Dtor_cf4())
						}(_pat_let13_0)
					}(_pat_let_tv11)
				}(_pat_let12_0)
			}(_570_v45))
			_ = _rhs73
			var _rhs74 _dafny.Set = _573_v48
			_ = _rhs74
			var _rhs75 _dafny.Int = _562_i2
			_ = _rhs75
			var _rhs76 _dafny.CodePoint = ((_this).F3()).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if p3 {
					return (_574_v49).Dtor_cf15()
				}
				return _562_i2
			})(), _dafny.IntOfUint32(((_this).F3()).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs76
			var _rhs77 bool = (_dafny.MultiSetFromSeq(_dafny.SeqOf(_562_i2))).IsDisjointFrom(_561_v37)
			_ = _rhs77
			_571_v46 = _rhs73
			_573_v48 = _rhs74
			r2 = _rhs75
			_569_v44 = _rhs76
			r0 = _rhs77
			var _577_v50 _dafny.Map
			_ = _577_v50
			_577_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
			r2 = (func() _dafny.Int {
				if (_this).Fm11(globalState) {
					return (_577_v50).Cardinality()
				}
				return _562_i2
			})()
		}
		var _578_v51 _dafny.Array
		_ = _578_v51
		var _nwElement0_26 _dafny.Int = p0
		_ = _nwElement0_26
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(13))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_26, 0)
		(_nw119).ArraySet1(_dafny.IntOfInt64(962), 1)
		(_nw119).ArraySet1(_dafny.IntOfUint32(((_this).F3()).Cardinality()), 2)
		(_nw119).ArraySet1(_dafny.IntOfUint32(((_this).F3()).Cardinality()), 3)
		(_nw119).ArraySet1(p0, 4)
		(_nw119).ArraySet1(p0, 5)
		(_nw119).ArraySet1(p0, 6)
		(_nw119).ArraySet1(p0, 7)
		(_nw119).ArraySet1(p0, 8)
		(_nw119).ArraySet1(p0, 9)
		(_nw119).ArraySet1(p0, 10)
		(_nw119).ArraySet1(p0, 11)
		(_nw119).ArraySet1(p0, 12)
		_578_v51 = _nw119
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _579_v52 _dafny.Sequence
		_ = _579_v52
		_579_v52 = _dafny.SeqOf((func(_pat_let14_0 D4) D4 {
			return func(_580_dt__update__tmp_h1 D4) D4 {
				return func(_pat_let15_0 _dafny.Int) D4 {
					return func(_581_dt__update_hcf9_h0 _dafny.Int) D4 {
						return Companion_D4_.Create_DC6_(_581_dt__update_hcf9_h0, (_580_dt__update__tmp_h1).Dtor_cf10(), (_580_dt__update__tmp_h1).Dtor_cf11())
					}(_pat_let15_0)
				}(_pat_let_tv12)
			}(_pat_let14_0)
		}(Companion_D4_.Create_DC6_((_511_v0).Cardinality(), _578_v51, p0))).Dtor_cf9(), (_dafny.Zero).Minus(p0), p0)
		var _582_v53 *C0
		_ = _582_v53
		var _nw120 *C0 = New_C0_()
		_ = _nw120
		_nw120.Ctor__(p2, p2)
		_582_v53 = _nw120
		var _583_v54 _dafny.Set
		_ = _583_v54
		_583_v54 = _dafny.SetOf(_582_v53, _582_v53)
		var _584_v55 _dafny.Map
		_ = _584_v55
		_584_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(_582_v53))
		var _585_v56 _dafny.Set
		_ = _585_v56
		_585_v56 = _dafny.SetOf(p3)
		var _rhs78 bool = (((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))).Minus(p0)).Cmp(_dafny.IntOfUint32((_579_v52).Cardinality())) <= 0
		_ = _rhs78
		var _rhs79 _dafny.Int = p0
		_ = _rhs79
		var _rhs80 bool = ((func() _dafny.Set {
			if (_584_v55).Contains(p0) {
				return (_584_v55).Get(p0).(_dafny.Set)
			}
			return _583_v54
		})()).IsProperSubsetOf(_583_v54)
		_ = _rhs80
		var _rhs81 _dafny.Int = (((_585_v56).Union(_585_v56)).Union(_585_v56)).Cardinality()
		_ = _rhs81
		r0 = _rhs78
		r2 = _rhs79
		r0 = _rhs80
		r2 = _rhs81
		r0 = _582_v53.F5
		var _rhs82 _dafny.Int = p0
		_ = _rhs82
		var _rhs83 bool = _582_v53.F4
		_ = _rhs83
		var _rhs84 bool = ((p0).Times(p0)).Cmp((_this).Fm1(p3, globalState)) != 0
		_ = _rhs84
		r2 = _rhs82
		r0 = _rhs83
		r0 = _rhs84
		var _586_v57 _dafny.Map
		_ = _586_v57
		_586_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
		var _587_v58 _dafny.Map
		_ = _587_v58
		_587_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _586_v57)
		var _source12 D5 = Companion_D5_.Create_DC10_(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tt")).Cardinality()), _587_v58, true)
		_ = _source12
		if _source12.Is_DC9() {
			var _588___mcc_h5 bool = _source12.Get_().(D5_DC9).Cf14
			_ = _588___mcc_h5
			var _589___mcc_h6 _dafny.Int = _source12.Get_().(D5_DC9).Cf15
			_ = _589___mcc_h6
			var _590___mcc_h7 _dafny.MultiSet = _source12.Get_().(D5_DC9).Cf16
			_ = _590___mcc_h7
			var _591_cf16 _dafny.MultiSet = _590___mcc_h7
			_ = _591_cf16
			var _592_cf15 _dafny.Int = _589___mcc_h6
			_ = _592_cf15
			var _593_cf14 bool = _588___mcc_h5
			_ = _593_cf14
			var _594_v59 D4
			_ = _594_v59
			_594_v59 = Companion_D4_.Create_DC6_((_dafny.SetOf(p2)).Cardinality(), _578_v51, _dafny.IntOfUint32((_dafny.SeqOf((_579_v52).Select((Companion_Default___.SafeIndex(_592_cf15, _dafny.IntOfUint32((_579_v52).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality()))
			var _595_v60 _dafny.Map
			_ = _595_v60
			_595_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _594_v59)
			var _596_v61 _dafny.Map
			_ = _596_v61
			_596_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v60, p0)
			_596_v61 = (_596_v61).Merge(_596_v61)
			var _597_v62 _dafny.Sequence
			_ = _597_v62
			_597_v62 = _dafny.SeqOf(!(p3), _582_v53.F4, (func() bool {
				if p3 {
					return p2
				}
				return _582_v53.F5
			})())
			var _rhs85 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_597_v62, _597_v62)
			_ = _rhs85
			var _rhs86 _dafny.Int = _dafny.IntOfUint32(((_this).F3()).Cardinality())
			_ = _rhs86
			_597_v62 = _rhs85
			_592_cf15 = _rhs86
			var _598_v63 _dafny.Sequence
			_ = _598_v63
			_598_v63 = _dafny.SeqOf(_511_v0)
			_511_v0 = (_511_v0).Union((_dafny.MultiSetOf(_582_v53.F4, false)).Difference((_598_v63).Select((Companion_Default___.SafeIndex(_592_cf15, _dafny.IntOfUint32((_598_v63).Cardinality()))).Uint32()).(_dafny.MultiSet)))
			var _599_v64 _dafny.Array
			_ = _599_v64
			var _nw121 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw121
			_599_v64 = _nw121
			var _600_v65 _dafny.MultiSet
			_ = _600_v65
			_600_v65 = _dafny.MultiSetOf(_582_v53, _582_v53)
			var _601_v66 _dafny.Map
			_ = _601_v66
			_601_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_599_v64, _600_v65)
			(_582_v53).F4 = (_597_v62).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _593_cf14 {
					return p0
				}
				return (_601_v66).Cardinality()
			})(), _dafny.IntOfUint32((_597_v62).Cardinality()))).Uint32()).(bool)
		} else if _source12.Is_DC10() {
			var _602___mcc_h8 bool = _source12.Get_().(D5_DC10).Cf17
			_ = _602___mcc_h8
			var _603___mcc_h9 _dafny.Int = _source12.Get_().(D5_DC10).Cf18
			_ = _603___mcc_h9
			var _604___mcc_h10 _dafny.Map = _source12.Get_().(D5_DC10).Cf19
			_ = _604___mcc_h10
			var _605___mcc_h11 bool = _source12.Get_().(D5_DC10).Cf20
			_ = _605___mcc_h11
			var _606_cf20 bool = _605___mcc_h11
			_ = _606_cf20
			var _607_cf19 _dafny.Map = _604___mcc_h10
			_ = _607_cf19
			var _608_cf18 _dafny.Int = _603___mcc_h9
			_ = _608_cf18
			var _609_cf17 bool = _602___mcc_h8
			_ = _609_cf17
			var _610_v67 _dafny.Sequence
			_ = _610_v67
			_610_v67 = _dafny.UnicodeSeqOfUtf8Bytes("wwqob")
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_578_v51), 0))
			_ = _index100
			(_578_v51).ArraySet1((_579_v52).Select((Companion_Default___.SafeIndex((_585_v56).Cardinality(), _dafny.IntOfUint32((_579_v52).Cardinality()))).Uint32()).(_dafny.Int), (_index100).Int())
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_578_v51), 0))
			_ = _index101
			var _rhs87 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_610_v67, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_611_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))
			_ = _rhs87
			var _rhs88 _dafny.Int = (_dafny.IntOfInt64(827)).Plus(_608_cf18)
			_ = _rhs88
			var _rhs89 _dafny.Int = _608_cf18
			_ = _rhs89
			var _lhs37 _dafny.Array = _578_v51
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_578_v51), 0))
			_ = _lhs38
			_610_v67 = _rhs87
			(_lhs37).ArraySet1(_rhs88, (_lhs38).Int())
			_608_cf18 = _rhs89
			_606_cf20 = _609_cf17
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_578_v51), 0))
			_ = _index102
			(_578_v51).ArraySet1((_578_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_578_v51), 0))).Int()).(_dafny.Int), (_index102).Int())
			(_582_v53).F5 = _582_v53.F5
		} else {
			var _612___mcc_h12 _dafny.MultiSet = _source12.Get_().(D5_DC8).Cf13
			_ = _612___mcc_h12
			var _613_cf13 _dafny.MultiSet = _612___mcc_h12
			_ = _613_cf13
			var _614_v68 _dafny.Set
			_ = _614_v68
			_614_v68 = _dafny.SetOf(p0, _dafny.IntOfInt64(-983))
			_511_v0 = (_511_v0).Update((_614_v68).IsProperSubsetOf(Companion_Default___.Fm18(_582_v53.F4, p3, _511_v0, globalState)), Companion_Default___.Abs(p0))
			var _615_v69 _dafny.Map
			_ = _615_v69
			_615_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _582_v53.F5)
			var _616_v70 _dafny.MultiSet
			_ = _616_v70
			_616_v70 = _dafny.MultiSetOf(_615_v69, _615_v69)
			r2 = (((_dafny.MultiSetOf((_dafny.Zero).Minus((_616_v70).Cardinality()))).Cardinality()).Minus((_579_v52).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_579_v52).Cardinality()))).Uint32()).(_dafny.Int))).Plus(p0)
			var _617_v71 *C0
			_ = _617_v71
			var _nw122 *C0 = New_C0_()
			_ = _nw122
			_nw122.Ctor__(_582_v53.F5, !(_582_v53.F5))
			_617_v71 = _nw122
			_511_v0 = _dafny.MultiSetOf(!(false) || (_617_v71.F4))
		}
		r0 = p3
		var _618_v72 _dafny.Set
		_ = _618_v72
		_618_v72 = _dafny.SetOf(_561_v37, _561_v37)
		var _619_v73 _dafny.Map
		_ = _619_v73
		_619_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
		var _620_v74 _dafny.Array
		_ = _620_v74
		var _nwElement0_27 _dafny.Map = _619_v73
		_ = _nwElement0_27
		var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(5))
		_ = _nw123
		(_nw123).ArraySet1(_nwElement0_27, 0)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_582_v53.F5, _582_v53.F4), 1)
		(_nw123).ArraySet1(_619_v73, 2)
		(_nw123).ArraySet1(_619_v73, 3)
		(_nw123).ArraySet1(_619_v73, 4)
		_620_v74 = _nw123
		var _621_v75 _dafny.Map
		_ = _621_v75
		_621_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_618_v72, _620_v74)
		r1 = (func() _dafny.Array {
			if (_621_v75).Contains(_dafny.SetOf(_561_v37, _dafny.MultiSetOf(p0), Companion_Default___.Fm19(p0, p0, _582_v53.F4, p0, globalState))) {
				return (_621_v75).Get(_dafny.SetOf(_561_v37, _dafny.MultiSetOf(p0), Companion_Default___.Fm19(p0, p0, _582_v53.F4, p0, globalState))).(_dafny.Array)
			}
			return _620_v74
		})()
		r2 = (_dafny.Zero).Minus(((_this).Fm1(p2, globalState)).Minus((_579_v52).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_579_v52).Cardinality()))).Uint32()).(_dafny.Int)))
		return r0, r1, r2
	}
}
func (_this *C1) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _622_v0 _dafny.Array
		_ = _622_v0
		var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw124
		_622_v0 = _nw124
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_622_v0), 0))); ; {
			_guard_loop_1, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _623_i0 _dafny.Int
			_623_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_623_i0).Sign() != -1) && ((_623_i0).Cmp(_dafny.ArrayLen((_622_v0), 0)) < 0)) {
				(_622_v0).ArraySet1((_this).F3(), (_623_i0).Int())
			}
		}
		var _624_i1 _dafny.Int
		_ = _624_i1
		_624_i1 = _dafny.Zero
		{
			for p2 {
				{
					if (_624_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_624_i1 = (_624_i1).Plus(_dafny.One)
					var _625_v1 bool
					_ = _625_v1
					_625_v1 = true
					_625_v1 = (_this).Fm11(globalState)
					var _626_v2 _dafny.Array
					_ = _626_v2
					var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
					_ = _nw125
					_626_v2 = _nw125
					var _627_v3 _dafny.Array
					_ = _627_v3
					var _nw126 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(5))
					_ = _nw126
					_627_v3 = _nw126
					var _628_v4 _dafny.Map
					_ = _628_v4
					_628_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_627_v3, _625_v1)
					var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_626_v2), 0))
					_ = _index103
					(_626_v2).ArraySet1((_628_v4).Merge(_628_v4), (_index103).Int())
					var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_626_v2), 0))
					_ = _index104
					(_626_v2).ArraySet1(_628_v4, (_index104).Int())
					var _629_v5 _dafny.CodePoint
					_ = _629_v5
					_629_v5 = _dafny.CodePoint('h')
					_629_v5 = ((_this).F3()).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_this).F3()).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _630_v6 _dafny.Set
					_ = _630_v6
					_630_v6 = _dafny.SetOf(p2, _625_v1)
					var _631_v7 _dafny.Map
					_ = _631_v7
					_631_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_630_v6).Union(_630_v6), p1)
					_631_v7 = (_631_v7).Merge((_631_v7).Merge(_631_v7))
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _632_v8 _dafny.Sequence
		_ = _632_v8
		_632_v8 = _dafny.SeqOf(true, p2)
		var _633_v9 _dafny.Map
		_ = _633_v9
		_633_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (_632_v8).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_632_v8).Cardinality()))).Uint32()).(bool))
		var _634_v10 _dafny.Sequence
		_ = _634_v10
		_634_v10 = _dafny.SeqOf(_633_v9, _633_v9, _633_v9, _633_v9, _633_v9)
		_634_v10 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2))
		var _635_v11 bool
		_ = _635_v11
		_635_v11 = false
		_635_v11 = _635_v11
		_635_v11 = _635_v11
		var _636_v12 _dafny.Set
		_ = _636_v12
		_636_v12 = _dafny.SetOf(p0)
		_636_v12 = _636_v12
		r0 = (_this).Fm1(p2, globalState)
		return r0
	}
}
func (_this *C1) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f2 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f2 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C2{}
var _ T2 = &C2{}
var _ T1 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F2() _dafny.Int {
	return _this._f2
}
func (_this *C2) F2_set_(value _dafny.Int) {
	_this._f2 = value
}
func (_this *C2) Ctor__(f2 _dafny.Int) {
	{
		(_this)._f2 = f2
	}
}
func (_this *C2) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F2()), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_637_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(_this.F2(), _this.F2())).Cardinality())
		})))).Cardinality())).Merge(func() _dafny.Map {
			var _coll19 = _dafny.NewMapBuilder()
			_ = _coll19
			for _iter21 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F2()), true)).Keys().Elements()); ; {
				_compr_19, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _638_v0 _dafny.Sequence
				_638_v0 = interface{}(_compr_19).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F2()), true)).Contains(_638_v0) {
					_coll19.Add(_638_v0, _this.F2())
				}
			}
			return _coll19.ToMap()
		}())).Merge((func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter22 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()), (_dafny.SetOf(true, true, true, !(false), false)).Cardinality()))).Elements()); ; {
				_compr_20, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _639_v1 _dafny.Sequence
				_639_v1 = interface{}(_compr_20).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()), (_dafny.SetOf(true, true, true, !(false), false)).Cardinality())), _639_v1) {
					_coll20.Add(_639_v1, (func() _dafny.Map {
						var _coll21 = _dafny.NewMapBuilder()
						_ = _coll21
						for _iter23 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('m'))).Elements()); ; {
							_compr_21, _ok23 := _iter23()
							if !_ok23 {
								break
							}
							var _640_v2 _dafny.CodePoint
							_640_v2 = interface{}(_compr_21).(_dafny.CodePoint)
							if (_dafny.MultiSetOf(_dafny.CodePoint('m'))).Contains(_640_v2) {
								_coll21.Add(_640_v2, _this.F2())
							}
						}
						return _coll21.ToMap()
					}()).Cardinality())
				}
			}
			return _coll20.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F2())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _this.F2())).Cardinality())))
	}
}
func (_this *C2) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F2()
	}
}
func (_this *C2) Fm8(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf((_dafny.SetOf(true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D5_.Create_DC10_(true, _this.F2(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(436))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_641_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(934), _dafny.IntOfInt64(570))); ; {
				_compr_22, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _642_v0 _dafny.Int
				_642_v0 = interface{}(_compr_22).(_dafny.Int)
				if ((_dafny.IntOfInt64(934)).Cmp(_642_v0) <= 0) && ((_642_v0).Cmp(_dafny.IntOfInt64(570)) < 0) {
					_coll22.Add(Companion_Default___.SafeModuloInt(_642_v0, _this.F2()), false)
				}
			}
			return _coll22.ToMap()
		}()), true)).Dtor_cf18(), !(true))).Cardinality(), _this.F2())).Union(_dafny.SetOf(_this.F2(), _this.F2(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cawq")).Cardinality()), _this.F2(), _this.F2()))).Equals((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ydjxbae")).Cardinality()))).Difference(_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _this.F2())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(911))).Cardinality())))
	}
}
func (_this *C2) Fm24(p0 _dafny.Sequence, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC14_(_dafny.SeqOf(_this.F2(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-59))).Cardinality())))).Dtor_cf26(), _this.F2())
	}
}
func (_this *C2) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _643_v0 _dafny.Set
		_ = _643_v0
		_643_v0 = _dafny.SetOf(_this.F2())
		var _644_v1 _dafny.Map
		_ = _644_v1
		_644_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _643_v0)
		var _645_v2 _dafny.Map
		_ = _645_v2
		_645_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
		var _646_v3 _dafny.Sequence
		_ = _646_v3
		_646_v3 = _dafny.UnicodeSeqOfUtf8Bytes("rftn")
		var _647_v4 _dafny.Array
		_ = _647_v4
		var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw127
		_647_v4 = _nw127
		var _648_v5 _dafny.Array
		_ = _648_v5
		var _nwElement0_28 bool = p3
		_ = _nwElement0_28
		var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(18))
		_ = _nw128
		(_nw128).ArraySet1(_nwElement0_28, 0)
		(_nw128).ArraySet1(p3, 1)
		(_nw128).ArraySet1(p3, 2)
		(_nw128).ArraySet1(p2, 3)
		(_nw128).ArraySet1(p3, 4)
		(_nw128).ArraySet1(false, 5)
		(_nw128).ArraySet1(p3, 6)
		(_nw128).ArraySet1(p2, 7)
		(_nw128).ArraySet1(p3, 8)
		(_nw128).ArraySet1(p2, 9)
		(_nw128).ArraySet1(p3, 10)
		(_nw128).ArraySet1(p3, 11)
		(_nw128).ArraySet1(true, 12)
		(_nw128).ArraySet1(p3, 13)
		(_nw128).ArraySet1(p2, 14)
		(_nw128).ArraySet1(p3, 15)
		(_nw128).ArraySet1(p2, 16)
		(_nw128).ArraySet1(p2, 17)
		_648_v5 = _nw128
		var _649_v6 _dafny.Map
		_ = _649_v6
		_649_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_648_v5, p0)
		var _650_v7 D4
		_ = _650_v7
		_650_v7 = Companion_D4_.Create_DC6_((_this).Fm1(p3, globalState), _647_v4, (_649_v6).Cardinality())
		var _651_v8 D4
		_ = _651_v8
		_651_v8 = Companion_D4_.Create_DC7_(_650_v7)
		var _652_v9 _dafny.Sequence
		_ = _652_v9
		_652_v9 = _dafny.SeqOf(_651_v8)
		var _653_v10 _dafny.Sequence
		_ = _653_v10
		_653_v10 = _dafny.SeqOf(p3, p2, p3, p3, p2)
		var _654_v11 _dafny.CodePoint
		_ = _654_v11
		_654_v11 = _dafny.CodePoint('e')
		var _655_v12 _dafny.Array
		_ = _655_v12
		var _nwElement0_29 _dafny.Int = Companion_Default___.SafeDivisionInt(((func() _dafny.Set {
			if (_644_v1).Contains(p3) {
				return (_644_v1).Get(p3).(_dafny.Set)
			}
			return _643_v0
		})()).Cardinality(), (_645_v2).Cardinality())
		_ = _nwElement0_29
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(16))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_29, 0)
		(_nw129).ArraySet1(_dafny.IntOfUint32((_646_v3).Cardinality()), 1)
		(_nw129).ArraySet1(p0, 2)
		(_nw129).ArraySet1(_dafny.IntOfUint32((_652_v9).Cardinality()), 3)
		(_nw129).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F2(), _this.F2()), 4)
		(_nw129).ArraySet1(_dafny.IntOfUint32((_653_v10).Cardinality()), 5)
		(_nw129).ArraySet1(p0, 6)
		(_nw129).ArraySet1((_dafny.MultiSetOf(_654_v11)).Cardinality(), 7)
		(_nw129).ArraySet1((_this).Fm1(false, globalState), 8)
		(_nw129).ArraySet1(_this.F2(), 9)
		(_nw129).ArraySet1(p0, 10)
		(_nw129).ArraySet1(_this.F2(), 11)
		(_nw129).ArraySet1((_dafny.Zero).Minus(_this.F2()), 12)
		(_nw129).ArraySet1(p0, 13)
		(_nw129).ArraySet1((_dafny.IntOfInt64(378)).Plus(_dafny.IntOfInt64(236)), 14)
		(_nw129).ArraySet1((_dafny.MultiSetFromSeq(_653_v10)).Cardinality(), 15)
		_655_v12 = _nw129
		var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
		_ = _index105
		(_647_v4).ArraySet1(_dafny.IntOfUint32((_646_v3).Cardinality()), (_index105).Int())
		var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
		_ = _index106
		var _rhs90 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_646_v3).Cardinality()), (_dafny.Zero).Minus((_this).Fm1(p2, globalState))))).Minus(_this.F2())
		_ = _rhs90
		var _rhs91 _dafny.Int = (_dafny.Zero).Minus(_this.F2())
		_ = _rhs91
		var _lhs39 _dafny.Array = _647_v4
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
		_ = _lhs40
		r2 = _rhs90
		(_lhs39).ArraySet1(_rhs91, (_lhs40).Int())
		if (_this.F2()).Cmp((_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int)) <= 0 {
			var _656_v13 _dafny.Map
			_ = _656_v13
			_656_v13 = (_this).Fm0(globalState)
			var _source13 _dafny.Map = _656_v13
			_ = _source13
			var _657___mcc_h0 _dafny.Map = _source13
			_ = _657___mcc_h0
			var _658_cf6 _dafny.Map = _657___mcc_h0
			_ = _658_cf6
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index107
			(_647_v4).ArraySet1((_dafny.Zero).Minus((_this).Fm1(p3, globalState)), (_index107).Int())
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index108
			(_647_v4).ArraySet1(p0, (_index108).Int())
			r0 = p2
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index109
			(_647_v4).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int)), (_index109).Int())
			var _659_v14 _dafny.Set
			_ = _659_v14
			_659_v14 = _dafny.SetOf(p2)
			_659_v14 = (func() _dafny.Set {
				if (p0).Cmp(p0) == 0 {
					return (_659_v14).Intersection(_659_v14)
				}
				return _659_v14
			})()
			var _660_v15 _dafny.Sequence
			_ = _660_v15
			_660_v15 = _dafny.SeqOf(_dafny.IntOfInt64(293), (_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int))
			var _661_v16 _dafny.Map
			_ = _661_v16
			_661_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v15, _dafny.IntOfUint32((_646_v3).Cardinality()))
			var _662_v17 _dafny.Sequence
			_ = _662_v17
			_662_v17 = _dafny.SeqOf(_656_v13, _661_v16)
			r0 = (p3) && (!_dafny.Companion_Sequence_.Equal(_662_v17, _662_v17))
			var _663_v18 D4
			_ = _663_v18
			_663_v18 = Companion_D4_.Create_DC6_(p0, _655_v12, _this.F2())
			var _664_v19 _dafny.Map
			_ = _664_v19
			_664_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_663_v18).Dtor_cf10(), p3)
			var _665_v20 *C0
			_ = _665_v20
			var _nw130 *C0 = New_C0_()
			_ = _nw130
			_nw130.Ctor__(p3, (func() bool {
				if (_664_v19).Contains(_655_v12) {
					return (_664_v19).Get(_655_v12).(bool)
				}
				return p3
			})())
			_665_v20 = _nw130
			(_this).F2_set_((Companion_Default___.SafeModuloInt(p0, p0)).Plus((_this.F2()).Times(_this.F2())))
		} else {
			var _666_v21 D7
			_ = _666_v21
			_666_v21 = Companion_D7_.Create_DC15_((_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int), _654_v11)
			var _667_v22 _dafny.Map
			_ = _667_v22
			_667_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_666_v21, p2)
			var _668_v23 _dafny.MultiSet
			_ = _668_v23
			_668_v23 = _dafny.MultiSetOf(_654_v11)
			var _669_v24 _dafny.MultiSet
			_ = _669_v24
			_669_v24 = _dafny.MultiSetOf((_645_v2).Cardinality())
			var _670_v25 _dafny.Sequence
			_ = _670_v25
			_670_v25 = _dafny.SeqOf(_this.F2(), p0)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index110
			var _rhs92 bool = ((_667_v22).Cardinality()).Cmp(Companion_Default___.SafeModuloInt(p0, _this.F2())) < 0
			_ = _rhs92
			var _rhs93 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).Fm1(p3, globalState), p0)
			_ = _rhs93
			var _rhs94 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_653_v10).Cardinality())))).Times((_668_v23).Cardinality())
			_ = _rhs94
			var _rhs95 _dafny.Int = Companion_Default___.SafeModuloInt((_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int), _this.F2())
			_ = _rhs95
			var _rhs96 _dafny.Int = ((_669_v24).Union((_dafny.MultiSetFromSeq(_670_v25)).Union(_669_v24))).Cardinality()
			_ = _rhs96
			var _lhs41 *C2 = _this
			_ = _lhs41
			var _lhs42 *C2 = _this
			_ = _lhs42
			var _lhs43 _dafny.Array = _647_v4
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _lhs44
			r0 = _rhs92
			_lhs41.F2_set_(_rhs93)
			r2 = _rhs94
			_lhs42.F2_set_(_rhs95)
			(_lhs43).ArraySet1(_rhs96, (_lhs44).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index111
			(_647_v4).ArraySet1(_dafny.IntOfInt64(814), (_index111).Int())
			_653_v10 = _653_v10
			var _671_v26 _dafny.MultiSet
			_ = _671_v26
			_671_v26 = _dafny.MultiSetOf(p2)
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))
			_ = _index112
			(_647_v4).ArraySet1(Companion_Default___.SafeDivisionInt(p0, ((func() _dafny.MultiSet {
				if p3 {
					return _671_v26
				}
				return _dafny.MultiSetOf(p3)
			})()).Cardinality()), (_index112).Int())
			var _672_v27 _dafny.Map
			_ = _672_v27
			_672_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_646_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), false))
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_648_v5), 0))
			_ = _index113
			(_648_v5).ArraySet1((Companion_D5_.Create_DC10_(p2, (_643_v0).Cardinality(), _672_v27, p2)).Dtor_cf20(), (_index113).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_648_v5), 0))
			_ = _index114
			(_648_v5).ArraySet1(p2, (_index114).Int())
		}
		_648_v5 = _648_v5
		var _673_v28 _dafny.Sequence
		_ = _673_v28
		_673_v28 = _dafny.SeqOf(_dafny.IntOfInt64(99))
		var _pat_let_tv13 = _654_v11
		_ = _pat_let_tv13
		var _pat_let_tv14 = p3
		_ = _pat_let_tv14
		var _pat_let_tv15 = p0
		_ = _pat_let_tv15
		var _pat_let_tv16 = _647_v4
		_ = _pat_let_tv16
		var _pat_let_tv17 = _647_v4
		_ = _pat_let_tv17
		r2 = func(_source14 D0) _dafny.Int {
			if _source14.Is_DC1() {
				var _674___mcc_h1 bool = _source14.Get_().(D0_DC1).Cf1
				_ = _674___mcc_h1
				var _675___mcc_h2 bool = _source14.Get_().(D0_DC1).Cf2
				_ = _675___mcc_h2
				var _676___mcc_h3 _dafny.CodePoint = _source14.Get_().(D0_DC1).Cf3
				_ = _676___mcc_h3
				var _677___mcc_h4 _dafny.Int = _source14.Get_().(D0_DC1).Cf4
				_ = _677___mcc_h4
				var _678_cf4 _dafny.Int = _677___mcc_h4
				_ = _678_cf4
				var _679_cf3 _dafny.CodePoint = _676___mcc_h3
				_ = _679_cf3
				var _680_cf2 bool = _675___mcc_h2
				_ = _680_cf2
				var _681_cf1 bool = _674___mcc_h1
				_ = _681_cf1
				return (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F2(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_678_cf4, _pat_let_tv13)).Cardinality()))
			} else {
				var _682___mcc_h5 bool = _source14.Get_().(D0_DC0).Cf0
				_ = _682___mcc_h5
				var _683_cf0 bool = _682___mcc_h5
				_ = _683_cf0
				return Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv14, _pat_let_tv15)).Cardinality(), (_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(_dafny.Int))
			}
		}(Companion_Default___.Fm25(_dafny.IntOfUint32((_673_v28).Cardinality()), (_this).Fm8(_646_v3, p2, globalState), _dafny.MultiSetOf(_dafny.IntOfInt64(245), _this.F2()), _dafny.SeqOf(p2), globalState))
		var _684_v29 D7
		_ = _684_v29
		_684_v29 = Companion_D7_.Create_DC14_(_673_v28)
		var _685_v30 _dafny.Map
		_ = _685_v30
		_685_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), p0)
		var _686_v31 _dafny.Array
		_ = _686_v31
		var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_673_v28, _673_v28)
		_ = _nwElement0_30
		var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(27))
		_ = _nw131
		(_nw131).ArraySet1(_nwElement0_30, 0)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F2()), _673_v28), 1)
		(_nw131).ArraySet1(_673_v28, 2)
		(_nw131).ArraySet1(_673_v28, 3)
		(_nw131).ArraySet1(_673_v28, 4)
		(_nw131).ArraySet1(_dafny.SeqOf(_this.F2()), 5)
		(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_687_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_688_i1 _dafny.Int) _dafny.Int {
				return _687_p0
			}
		})(p0))), 6)
		(_nw131).ArraySet1(_dafny.SeqOf(_this.F2(), (_this).Fm1(p2, globalState)), 7)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_673_v28, _dafny.SeqOf(p0)), 8)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_673_v28, _673_v28), 9)
		(_nw131).ArraySet1(_673_v28, 10)
		(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(443))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_689_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_690_i2 _dafny.Int) _dafny.Int {
				return _689_p0
			}
		})(p0))), 11)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_673_v28, _673_v28), 12)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F2(), (_643_v0).Cardinality(), (_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int)), _673_v28), 13)
		(_nw131).ArraySet1(_673_v28, 14)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F2()), _dafny.SeqOf(_this.F2(), (_673_v28).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_673_v28).Cardinality()))).Uint32()).(_dafny.Int))), 15)
		(_nw131).ArraySet1(_673_v28, 16)
		(_nw131).ArraySet1(_673_v28, 17)
		(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_691_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_692_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_691_v3).Cardinality())
			}
		})(_646_v3))), 18)
		(_nw131).ArraySet1(_673_v28, 19)
		(_nw131).ArraySet1((_684_v29).Dtor_cf26(), 20)
		(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_647_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_647_v4), 0))).Int()).(_dafny.Int), (_685_v30).Cardinality()), _dafny.Companion_Sequence_.Update(_673_v28, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_673_v28).Cardinality()))).Uint32(), _this.F2())), 21)
		(_nw131).ArraySet1(_673_v28, 22)
		(_nw131).ArraySet1(_673_v28, 23)
		(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(700))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg42 _dafny.Int) interface{} {
				return coer42(arg42)
			}
		}(func(_693_i4 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(334)
		})), 24)
		(_nw131).ArraySet1(_dafny.SeqOf(_this.F2()), 25)
		(_nw131).ArraySet1(_673_v28, 26)
		_686_v31 = _nw131
		for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_686_v31), 0))); ; {
			_guard_loop_2, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _694_i0 _dafny.Int
			_694_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_694_i0).Sign() != -1) && ((_694_i0).Cmp(_dafny.ArrayLen((_686_v31), 0)) < 0)) {
				(_686_v31).ArraySet1(_673_v28, (_694_i0).Int())
			}
		}
		for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_647_v4), 0))); ; {
			_guard_loop_3, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _695_i5 _dafny.Int
			_695_i5 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_695_i5).Sign() != -1) && ((_695_i5).Cmp(_dafny.ArrayLen((_647_v4), 0)) < 0)) {
				(_647_v4).ArraySet1((_695_i5).Plus(p0), (_695_i5).Int())
			}
		}
		r0 = (!(p2) || (false)) || (p3)
		var _696_v32 _dafny.Array
		_ = _696_v32
		var _len0_10 _dafny.Int = _dafny.One
		_ = _len0_10
		var _nw132 _dafny.Array
		_ = _nw132
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw132 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Map = (func(_697_p3 bool) func(_dafny.Int) _dafny.Map {
				return func(_698_i6 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_p3, _697_p3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _697_p3))
				}
			})(p3)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw132 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw132).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw132).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_696_v32 = _nw132
		r1 = _696_v32
		r2 = Companion_Default___.SafeDivisionInt(p0, (_this.F2()).Minus(p0))
		return r0, r1, r2
	}
}
func (_this *C2) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		if p2 {
			var _699_v0 _dafny.Array
			_ = _699_v0
			var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw133
			_699_v0 = _nw133
			var _700_v1 _dafny.Array
			_ = _700_v1
			_700_v1 = _699_v0
			_700_v1 = _699_v0
			var _701_v2 _dafny.Sequence
			_ = _701_v2
			_701_v2 = _dafny.SeqOf(p2)
			var _702_v3 _dafny.Sequence
			_ = _702_v3
			_702_v3 = _dafny.SeqOf(p2, (_701_v2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_701_v2).Cardinality()))).Uint32()).(bool))
			var _703_v4 _dafny.MultiSet
			_ = _703_v4
			_703_v4 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kvtp")).Cardinality()), p1, (_this.F2()).Plus(_dafny.IntOfInt64(514)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2), _702_v3)).Cardinality()))
			var _704_v5 _dafny.Map
			_ = _704_v5
			_704_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			(_this).F2_set_((func() _dafny.Int {
				if (_703_v4).Contains(_this.F2()) {
					return (_703_v4).Multiplicity(_this.F2())
				}
				return ((_704_v5).Merge(_704_v5)).Cardinality()
			})())
			var _705_v6 _dafny.Sequence
			_ = _705_v6
			_705_v6 = _dafny.UnicodeSeqOfUtf8Bytes("h")
			var _706_v7 _dafny.Sequence
			_ = _706_v7
			_706_v7 = _dafny.SeqOf(_705_v6, _705_v6)
			_706_v7 = _dafny.Companion_Sequence_.Concatenate(_706_v7, _dafny.SeqOf(_705_v6, _705_v6))
			var _707_v8 *C1
			_ = _707_v8
			var _nw134 *C1 = New_C1_()
			_ = _nw134
			_nw134.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("bawn"))
			_707_v8 = _nw134
			if ((func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-299), _dafny.IntOfInt64(388))); ; {
					_compr_23, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _708_v9 _dafny.Int
					_708_v9 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(-299)).Cmp(_708_v9) <= 0) && ((_708_v9).Cmp(_dafny.IntOfInt64(388)) < 0) {
						_coll23.Add((_708_v9).Plus(_this.F2()))
					}
				}
				return _coll23.ToSet()
			}()).Cardinality()).Cmp(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_702_v3).Cardinality())), _this.F2())) <= 0 {
				var _709_v10 *C1
				_ = _709_v10
				var _nw135 *C1 = New_C1_()
				_ = _nw135
				_nw135.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("qlxtt"))
				_709_v10 = _nw135
				var _710_v11 _dafny.Map
				_ = _710_v11
				_710_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_705_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2))
				var _711_v12 _dafny.Array
				_ = _711_v12
				var _nwElement0_31 _dafny.Int = _this.F2()
				_ = _nwElement0_31
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.One)
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_31, 0)
				_711_v12 = _nw136
				var _712_v13 _dafny.Set
				_ = _712_v13
				_712_v13 = _dafny.SetOf(_711_v12)
				var _713_v14 *C0
				_ = _713_v14
				var _nw137 *C0 = New_C0_()
				_ = _nw137
				_nw137.Ctor__(!(p2) || ((Companion_D5_.Create_DC10_(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(48))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_714_v6 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_715_i0 _dafny.Int) _dafny.CodePoint {
						return (_714_v6).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_714_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_705_v6)))).Cardinality()), _710_v11, p2)).Dtor_cf17()), (_712_v13).Equals(_712_v13))
				_713_v14 = _nw137
				var _716_v15 D0
				_ = _716_v15
				_716_v15 = Companion_D0_.Create_DC0_(!(!(p2)))
				var _717_v16 _dafny.CodePoint
				_ = _717_v16
				_717_v16 = _dafny.CodePoint('c')
				var _718_v17 _dafny.Array
				_ = _718_v17
				var _nwElement0_32 D0 = _716_v15
				_ = _nwElement0_32
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(5))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_32, 0)
				(_nw138).ArraySet1(Companion_Default___.Fm25((Companion_Default___.Fm26((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v14.F5, p2))).Cardinality(), _713_v14.F4, _713_v14.F5, _713_v14.F5, globalState)).Cardinality(), _713_v14.F4, _703_v4, _702_v3, globalState), 1)
				(_nw138).ArraySet1(Companion_D0_.Create_DC0_(_713_v14.F5), 2)
				(_nw138).ArraySet1(_716_v15, 3)
				(_nw138).ArraySet1(_716_v15, 4)
				_718_v17 = _nw138
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_718_v17), 0))
				_ = _index115
				(_718_v17).ArraySet1(_716_v15, (_index115).Int())
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_718_v17), 0))
				_ = _index116
				(_718_v17).ArraySet1(Companion_Default___.Fm25(p0, !(_713_v14.F4), _703_v4, _dafny.Companion_Sequence_.Update(_702_v3, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_702_v3).Cardinality()))).Uint32(), !(true)), globalState), (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))
				_ = _index117
				(_711_v12).ArraySet1((_this.F2()).Plus(p0), (_index117).Int())
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))
				_ = _index118
				(_711_v12).ArraySet1((_dafny.IntOfInt64(-544)).Plus(p0), (_index118).Int())
				var _719_v18 _dafny.Set
				_ = _719_v18
				_719_v18 = _dafny.SetOf(_dafny.IntOfUint32((_705_v6).Cardinality()), _dafny.IntOfUint32((_702_v3).Cardinality()))
				var _720_v19 _dafny.Map
				_ = _720_v19
				_720_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_dafny.Zero).Minus((_719_v18).Cardinality())).Minus((_711_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))).Int()).(_dafny.Int)))
				var _721_v20 _dafny.Sequence
				_ = _721_v20
				_721_v20 = _dafny.SeqOf((Companion_Default___.Fm27((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), !(false))).Cardinality(), p2, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_720_v19).Cardinality(), p1)))
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))
				_ = _index119
				var _rhs97 _dafny.Map = (_721_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()), _dafny.IntOfUint32((_721_v20).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _rhs97
				var _rhs98 bool = !(_713_v14.F5)
				_ = _rhs98
				var _rhs99 _dafny.Int = (_711_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))).Int()).(_dafny.Int)
				_ = _rhs99
				var _rhs100 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F2(), (_dafny.IntOfUint32(((_709_v10).F3()).Cardinality())).Times(_this.F2()))
				_ = _rhs100
				var _rhs101 _dafny.Map = _720_v19
				_ = _rhs101
				var _lhs45 *C0 = _713_v14
				_ = _lhs45
				var _lhs46 _dafny.Array = _711_v12
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_711_v12), 0))
				_ = _lhs47
				_720_v19 = _rhs97
				_lhs45.F4 = _rhs98
				r0 = _rhs99
				(_lhs46).ArraySet1(_rhs100, (_lhs47).Int())
				_720_v19 = _rhs101
			} else {
				var _722_v21 *C0
				_ = _722_v21
				var _nw139 *C0 = New_C0_()
				_ = _nw139
				_nw139.Ctor__(!((p1).Cmp(p0) >= 0), !(p2))
				_722_v21 = _nw139
				var _723_v22 _dafny.CodePoint
				_ = _723_v22
				_723_v22 = _dafny.CodePoint('i')
				var _724_v23 _dafny.Map
				_ = _724_v23
				_724_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_704_v5).Cardinality(), _723_v22)
				(_722_v21).F4 = !(!_dafny.Companion_Sequence_.Contains((_707_v8).F3(), (func() _dafny.CodePoint {
					if (_724_v23).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32(((_707_v8).F3()).Cardinality()))) {
						return (_724_v23).Get((_dafny.Zero).Minus(_dafny.IntOfUint32(((_707_v8).F3()).Cardinality()))).(_dafny.CodePoint)
					}
					return _723_v22
				})()))
				var _725_v24 _dafny.Array
				_ = _725_v24
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(28))
				_ = _nw140
				_725_v24 = _nw140
				var _726_v25 _dafny.Map
				_ = _726_v25
				_726_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v22, _725_v24)
				var _727_v26 _dafny.MultiSet
				_ = _727_v26
				_727_v26 = _dafny.MultiSetOf(_725_v24, (func() _dafny.Array {
					if (_726_v25).Contains(_dafny.CodePoint('s')) {
						return (_726_v25).Get(_dafny.CodePoint('s')).(_dafny.Array)
					}
					return _725_v24
				})())
				var _728_v27 _dafny.Map
				_ = _728_v27
				_728_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_727_v26, _this.F2())
				var _729_v28 _dafny.Map
				_ = _729_v28
				_729_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _722_v21.F5)
				_728_v27 = (_728_v27).Update((_727_v26).Update(_725_v24, Companion_Default___.Abs(p1)), (_729_v28).Cardinality())
				var _730_v29 D0
				_ = _730_v29
				_730_v29 = Companion_D0_.Create_DC0_(p2)
				var _731_v30 _dafny.Array
				_ = _731_v30
				var _nwElement0_33 bool = _722_v21.F4
				_ = _nwElement0_33
				var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(19))
				_ = _nw141
				(_nw141).ArraySet1(_nwElement0_33, 0)
				(_nw141).ArraySet1((_this).Fm8((_707_v8).F3(), false, globalState), 1)
				(_nw141).ArraySet1(p2, 2)
				(_nw141).ArraySet1(_722_v21.F4, 3)
				(_nw141).ArraySet1(_722_v21.F5, 4)
				(_nw141).ArraySet1(_722_v21.F4, 5)
				(_nw141).ArraySet1(false, 6)
				(_nw141).ArraySet1(_722_v21.F5, 7)
				(_nw141).ArraySet1(!(!(false)), 8)
				(_nw141).ArraySet1((_722_v21.F5) || (p2), 9)
				(_nw141).ArraySet1((p2) && (_722_v21.F4), 10)
				(_nw141).ArraySet1((_730_v29).Dtor_cf0(), 11)
				(_nw141).ArraySet1(_722_v21.F4, 12)
				(_nw141).ArraySet1(true, 13)
				(_nw141).ArraySet1(p2, 14)
				(_nw141).ArraySet1(true, 15)
				(_nw141).ArraySet1((_703_v4).Contains((func() _dafny.Int {
					if (_703_v4).Contains(p0) {
						return (_703_v4).Multiplicity(p0)
					}
					return p1
				})()), 16)
				(_nw141).ArraySet1(true, 17)
				(_nw141).ArraySet1(_722_v21.F4, 18)
				_731_v30 = _nw141
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _index120
				(_731_v30).ArraySet1((_dafny.IntOfUint32(((_707_v8).F3()).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(208))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_732_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_733_i1 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_732_p0)).Cardinality(), _this.F2())).Cardinality()
					}
				})(p0)))).Cardinality())) > 0, (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _index121
				(_731_v30).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf((_707_v8).F3(), (_707_v8).F3()), (_index121).Int())
				var _734_v31 _dafny.Map
				_ = _734_v31
				_734_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_731_v30, _this.F2())
				var _735_v32 _dafny.Array
				_ = _735_v32
				var _nwElement0_34 _dafny.Int = p0
				_ = _nwElement0_34
				var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(12))
				_ = _nw142
				(_nw142).ArraySet1(_nwElement0_34, 0)
				(_nw142).ArraySet1(p0, 1)
				(_nw142).ArraySet1((_729_v28).Cardinality(), 2)
				(_nw142).ArraySet1(_this.F2(), 3)
				(_nw142).ArraySet1(_dafny.IntOfInt64(864), 4)
				(_nw142).ArraySet1((_dafny.Zero).Minus(_this.F2()), 5)
				(_nw142).ArraySet1((_707_v8).Fm1(p2, globalState), 6)
				(_nw142).ArraySet1(p1, 7)
				(_nw142).ArraySet1((_734_v31).Cardinality(), 8)
				(_nw142).ArraySet1(p0, 9)
				(_nw142).ArraySet1(_dafny.IntOfInt64(321), 10)
				(_nw142).ArraySet1(p0, 11)
				_735_v32 = _nw142
				var _736_v33 _dafny.MultiSet
				_ = _736_v33
				_736_v33 = _dafny.MultiSetOf(_735_v32, _735_v32)
				var _737_v34 D5
				_ = _737_v34
				_737_v34 = Companion_D5_.Create_DC10_(_722_v21.F4, _this.F2(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_707_v8).F3(), _729_v28), _722_v21.F5)
				var _738_v35 _dafny.Map
				_ = _738_v35
				_738_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Times(_dafny.IntOfInt64(87)), _dafny.SetOf(_737_v34))
				var _739_v36 _dafny.Set
				_ = _739_v36
				_739_v36 = _dafny.SetOf(_722_v21.F4)
				var _740_v37 _dafny.MultiSet
				_ = _740_v37
				_740_v37 = _dafny.MultiSetOf(!(_722_v21.F4))
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _index122
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _index123
				var _rhs102 bool = _722_v21.F4
				_ = _rhs102
				var _rhs103 _dafny.MultiSet = (_736_v33).Intersection(_736_v33)
				_ = _rhs103
				var _rhs104 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_705_v6, Companion_Default___.Fm28(p2, _722_v21.F4, _722_v21.F4, _739_v36, globalState)), (_707_v8).F3())
				_ = _rhs104
				var _rhs105 _dafny.Map = (_738_v35).Merge(_738_v35)
				_ = _rhs105
				var _rhs106 _dafny.Sequence = _dafny.SeqOf((_740_v37).IsProperSubsetOf(_740_v37))
				_ = _rhs106
				var _lhs48 _dafny.Array = _731_v30
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _lhs49
				var _lhs50 _dafny.Array = _731_v30
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_731_v30), 0))
				_ = _lhs51
				(_lhs48).ArraySet1(_rhs102, (_lhs49).Int())
				_736_v33 = _rhs103
				(_lhs50).ArraySet1(_rhs104, (_lhs51).Int())
				_738_v35 = _rhs105
				_702_v3 = _rhs106
			}
		} else {
			var _741_v38 bool
			_ = _741_v38
			_741_v38 = true
			var _742_v39 _dafny.Sequence
			_ = _742_v39
			_742_v39 = _dafny.SeqOf(p2, p2, !(p2))
			var _743_v41 _dafny.Sequence
			_ = _743_v41
			_743_v41 = _dafny.SeqOf(p1, p0)
			_741_v38 = (_this).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(202))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}(func(_744_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), (_742_v39).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter28 := _dafny.Iterate((_743_v41).Elements()); ; {
					_compr_24, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _745_v40 _dafny.Int
					_745_v40 = interface{}(_compr_24).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_743_v41, _745_v40) {
						_coll24.Add(Companion_Default___.SafeDivisionInt(_745_v40, (_dafny.Zero).Minus(_this.F2())), _741_v38)
					}
				}
				return _coll24.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_742_v39).Cardinality()))).Uint32()).(bool), globalState)
			(_this).F2_set_(Companion_Default___.SafeDivisionInt(p1, p0))
			var _746_v42 _dafny.Sequence
			_ = _746_v42
			_746_v42 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_747_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})))
			var _748_v43 *C1
			_ = _748_v43
			var _nw143 *C1 = New_C1_()
			_ = _nw143
			_nw143.Ctor__((_746_v42).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_746_v42).Cardinality()))).Uint32()).(_dafny.Sequence))
			_748_v43 = _nw143
			var _source15 D5 = Companion_Default___.Fm29(globalState)
			_ = _source15
			if _source15.Is_DC9() {
				var _749___mcc_h0 bool = _source15.Get_().(D5_DC9).Cf14
				_ = _749___mcc_h0
				var _750___mcc_h1 _dafny.Int = _source15.Get_().(D5_DC9).Cf15
				_ = _750___mcc_h1
				var _751___mcc_h2 _dafny.MultiSet = _source15.Get_().(D5_DC9).Cf16
				_ = _751___mcc_h2
				var _752_cf16 _dafny.MultiSet = _751___mcc_h2
				_ = _752_cf16
				var _753_cf15 _dafny.Int = _750___mcc_h1
				_ = _753_cf15
				var _754_cf14 bool = _749___mcc_h0
				_ = _754_cf14
				var _755_v44 _dafny.Array
				_ = _755_v44
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_11
				var _nw144 _dafny.Array
				_ = _nw144
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw144 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) D7 = func(_756_i4 _dafny.Int) D7 {
						return Companion_D7_.Create_DC15_(_dafny.IntOfInt64(-344), _dafny.CodePoint('a'))
					}
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw144 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw144).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw144).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_755_v44 = _nw144
				var _757_v45 _dafny.Set
				_ = _757_v45
				_757_v45 = _dafny.SetOf(_755_v44, _755_v44)
				_757_v45 = _757_v45
				(_this).F2_set_(p0)
				_753_cf15 = _dafny.IntOfUint32(((_748_v43).F3()).Cardinality())
				_754_cf14 = _754_cf14
			} else if _source15.Is_DC10() {
				var _758___mcc_h3 bool = _source15.Get_().(D5_DC10).Cf17
				_ = _758___mcc_h3
				var _759___mcc_h4 _dafny.Int = _source15.Get_().(D5_DC10).Cf18
				_ = _759___mcc_h4
				var _760___mcc_h5 _dafny.Map = _source15.Get_().(D5_DC10).Cf19
				_ = _760___mcc_h5
				var _761___mcc_h6 bool = _source15.Get_().(D5_DC10).Cf20
				_ = _761___mcc_h6
				var _762_cf20 bool = _761___mcc_h6
				_ = _762_cf20
				var _763_cf19 _dafny.Map = _760___mcc_h5
				_ = _763_cf19
				var _764_cf18 _dafny.Int = _759___mcc_h4
				_ = _764_cf18
				var _765_cf17 bool = _758___mcc_h3
				_ = _765_cf17
				_765_cf17 = p2
				var _766_v46 _dafny.Map
				_ = _766_v46
				_766_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _767_v47 D5
				_ = _767_v47
				_767_v47 = Companion_D5_.Create_DC10_(_765_cf17, _dafny.IntOfInt64(637), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_748_v43).F3(), _766_v46), _765_cf17)
				var _768_v48 _dafny.Set
				_ = _768_v48
				_768_v48 = _dafny.SetOf((_767_v47).Dtor_cf18())
				var _769_v49 _dafny.Array
				_ = _769_v49
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw145
				_769_v49 = _nw145
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_769_v49), 0))
				_ = _index124
				(_769_v49).ArraySet1(_this.F2(), (_index124).Int())
				var _770_v50 _dafny.Map
				_ = _770_v50
				_770_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_769_v49, p2)
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_769_v49), 0))
				_ = _index125
				var _rhs107 _dafny.Set = (_dafny.SetOf(p1)).Intersection(_dafny.SetOf(p0))
				_ = _rhs107
				var _rhs108 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_748_v43).F3(), (_748_v43).F3())).Cardinality())
				_ = _rhs108
				var _rhs109 _dafny.Int = _764_cf18
				_ = _rhs109
				var _rhs110 _dafny.Array = _769_v49
				_ = _rhs110
				var _rhs111 bool = (func() bool {
					if (_770_v50).Contains(_769_v49) {
						return (_770_v50).Get(_769_v49).(bool)
					}
					return (_742_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.IntOfUint32((_742_v39).Cardinality()))).Uint32()).(bool)
				})()
				_ = _rhs111
				var _lhs52 *C2 = _this
				_ = _lhs52
				var _lhs53 _dafny.Array = _769_v49
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_769_v49), 0))
				_ = _lhs54
				_768_v48 = _rhs107
				_lhs52.F2_set_(_rhs108)
				(_lhs53).ArraySet1(_rhs109, (_lhs54).Int())
				_769_v49 = _rhs110
				_741_v38 = _rhs111
				r0 = _this.F2()
				var _771_v51 _dafny.MultiSet
				_ = _771_v51
				_771_v51 = _dafny.MultiSetOf((_748_v43).F3(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(48))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}(func(_772_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})))
				var _773_v52 _dafny.Map
				_ = _773_v52
				_773_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_771_v51, (p0).Cmp(_dafny.IntOfInt64(48)) >= 0)
				_773_v52 = (_773_v52).Update(Companion_Default___.Fm30(_765_cf17, globalState), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okwhp")).Cardinality())).Cmp(_this.F2()) > 0)
			} else {
				var _774___mcc_h7 _dafny.MultiSet = _source15.Get_().(D5_DC8).Cf13
				_ = _774___mcc_h7
				var _775_cf13 _dafny.MultiSet = _774___mcc_h7
				_ = _775_cf13
				_741_v38 = _741_v38
				var _776_v53 _dafny.Array
				_ = _776_v53
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw146
				_776_v53 = _nw146
				var _777_v54 bool
				_ = _777_v54
				var _778_v55 _dafny.Array
				_ = _778_v55
				var _779_v56 _dafny.Int
				_ = _779_v56
				var _out19 bool
				_ = _out19
				var _out20 _dafny.Array
				_ = _out20
				var _out21 _dafny.Int
				_ = _out21
				_out19, _out20, _out21 = (_748_v43).M1((_748_v43).Fm1(false, globalState), _776_v53, (_742_v39).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_742_v39).Cardinality()))).Uint32()).(bool), p2, globalState)
				_777_v54 = _out19
				_778_v55 = _out20
				_779_v56 = _out21
				var _780_v57 D0
				_ = _780_v57
				_780_v57 = Companion_D0_.Create_DC0_((_this.F2()).Cmp(p0) > 0)
				var _rhs112 _dafny.Int = _dafny.IntOfInt64(-629)
				_ = _rhs112
				var _rhs113 _dafny.Int = (_748_v43).Fm1(_777_v54, globalState)
				_ = _rhs113
				var _rhs114 D0 = _780_v57
				_ = _rhs114
				var _rhs115 bool = (_748_v43).Fm11(globalState)
				_ = _rhs115
				var _lhs55 *C2 = _this
				_ = _lhs55
				r0 = _rhs112
				_lhs55.F2_set_(_rhs113)
				_780_v57 = _rhs114
				_741_v38 = _rhs115
				var _781_v58 _dafny.Set
				_ = _781_v58
				_781_v58 = _dafny.SetOf(_this.F2(), _779_v56)
				var _782_v59 _dafny.Map
				_ = _782_v59
				_782_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_741_v38, _779_v56)
				_741_v38 = (_dafny.SetOf(_dafny.IntOfUint32(((_748_v43).F3()).Cardinality()))).IsSubsetOf((_781_v58).Intersection(_dafny.SetOf((_dafny.Zero).Minus((_782_v59).Cardinality()), _this.F2(), (_dafny.Zero).Minus(_779_v56), _779_v56, _dafny.IntOfInt64(-225))))
			}
			var _783_v60 _dafny.Map
			_ = _783_v60
			_783_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(531), _741_v38)
			var _784_v61 _dafny.Map
			_ = _784_v61
			_784_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_748_v43).F3(), _783_v60)
			var _785_v62 D5
			_ = _785_v62
			_785_v62 = Companion_D5_.Create_DC10_(_741_v38, p1, _784_v61, p2)
			var _786_v63 _dafny.MultiSet
			_ = _786_v63
			_786_v63 = _dafny.MultiSetOf(p2, p2, _741_v38, _741_v38)
			var _787_v64 _dafny.Set
			_ = _787_v64
			_787_v64 = _dafny.SetOf(_this.F2())
			var _788_v65 _dafny.Map
			_ = _788_v65
			_788_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_741_v38, p2)
			var _789_v66 _dafny.Array
			_ = _789_v66
			var _nwElement0_35 _dafny.Int = p1
			_ = _nwElement0_35
			var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(25))
			_ = _nw147
			(_nw147).ArraySet1(_nwElement0_35, 0)
			(_nw147).ArraySet1(_this.F2(), 1)
			(_nw147).ArraySet1((func() _dafny.Int {
				if p2 {
					return (_dafny.SetOf(_dafny.IntOfInt64(-236))).Cardinality()
				}
				return p1
			})(), 2)
			(_nw147).ArraySet1(_this.F2(), 3)
			(_nw147).ArraySet1(p1, 4)
			(_nw147).ArraySet1(Companion_Default___.SafeDivisionInt((_785_v62).Dtor_cf18(), _this.F2()), 5)
			(_nw147).ArraySet1((_dafny.Zero).Minus(_this.F2()), 6)
			(_nw147).ArraySet1(p1, 7)
			(_nw147).ArraySet1((func() _dafny.Int {
				if p2 {
					return p1
				}
				return (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_786_v63).Contains(_741_v38) {
						return (_786_v63).Multiplicity(_741_v38)
					}
					return _this.F2()
				})()))
			})(), 8)
			(_nw147).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_this.F2(), _dafny.IntOfInt64(619))).Cardinality())).Minus(p1), 9)
			(_nw147).ArraySet1((_dafny.IntOfInt64(241)).Plus(_dafny.IntOfInt64(-534)), 10)
			(_nw147).ArraySet1(_this.F2(), 11)
			(_nw147).ArraySet1((func() _dafny.Int {
				if _741_v38 {
					return (_dafny.Zero).Minus(p0)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uyp")).Cardinality())
			})(), 12)
			(_nw147).ArraySet1(_dafny.IntOfInt64(276), 13)
			(_nw147).ArraySet1(p0, 14)
			(_nw147).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_743_v41).Cardinality()), _dafny.IntOfUint32((_742_v39).Cardinality())), 15)
			(_nw147).ArraySet1((func() _dafny.Int {
				if _741_v38 {
					return _dafny.IntOfInt64(308)
				}
				return p1
			})(), 16)
			(_nw147).ArraySet1(_this.F2(), 17)
			(_nw147).ArraySet1((_787_v64).Cardinality(), 18)
			(_nw147).ArraySet1(((_788_v65).Cardinality()).Minus(p0), 19)
			(_nw147).ArraySet1((_786_v63).Cardinality(), 20)
			(_nw147).ArraySet1(_this.F2(), 21)
			(_nw147).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), p0), 22)
			(_nw147).ArraySet1(_dafny.IntOfUint32(((_748_v43).F3()).Cardinality()), 23)
			(_nw147).ArraySet1(p0, 24)
			_789_v66 = _nw147
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_789_v66), 0))
			_ = _index126
			(_789_v66).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_743_v41).Cardinality())), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_789_v66), 0))
			_ = _index127
			(_789_v66).ArraySet1(p0, (_index127).Int())
		}
		var _790_v67 _dafny.Sequence
		_ = _790_v67
		_790_v67 = _dafny.UnicodeSeqOfUtf8Bytes("hhsiw")
		_790_v67 = _790_v67
		var _791_v68 _dafny.CodePoint
		_ = _791_v68
		_791_v68 = _dafny.CodePoint('d')
		var _792_v69 _dafny.Map
		_ = _792_v69
		_792_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_791_v68, _dafny.Companion_Sequence_.Concatenate(_790_v67, _790_v67))
		_792_v69 = (_792_v69).Update(_791_v68, _790_v67)
		var _793_v70 _dafny.Sequence
		_ = _793_v70
		_793_v70 = _dafny.SeqOf(p2, p2, p2)
		var _794_v71 _dafny.Sequence
		_ = _794_v71
		_794_v71 = _dafny.SeqOf(_790_v67, _790_v67)
		var _795_v73 _dafny.Map
		_ = _795_v73
		_795_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _796_v74 _dafny.Sequence
		_ = _796_v74
		_796_v74 = _dafny.SeqOf((_795_v73).Cardinality())
		var _797_v75 _dafny.Map
		_ = _797_v75
		_797_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v74, true)
		var _798_v76 _dafny.Array
		_ = _798_v76
		var _nwElement0_36 bool = (_793_v70).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_793_v70).Cardinality()))).Uint32()).(bool)
		_ = _nwElement0_36
		var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(19))
		_ = _nw148
		(_nw148).ArraySet1(_nwElement0_36, 0)
		(_nw148).ArraySet1(p2, 1)
		(_nw148).ArraySet1(p2, 2)
		(_nw148).ArraySet1(p2, 3)
		(_nw148).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((_794_v71).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_794_v71).Cardinality()))).Uint32()).(_dafny.Sequence), _790_v67), 4)
		(_nw148).ArraySet1((_this).Fm8(_790_v67, p2, globalState), 5)
		(_nw148).ArraySet1((_this).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-707))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}((func(_799_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_800_i6 _dafny.Int) _dafny.CodePoint {
				return _799_v68
			}
		})(_791_v68))), p2, globalState), 6)
		(_nw148).ArraySet1(p2, 7)
		(_nw148).ArraySet1(p2, 8)
		(_nw148).ArraySet1((_this.F2()).Cmp(_this.F2()) != 0, 9)
		(_nw148).ArraySet1(p2, 10)
		(_nw148).ArraySet1((p1).Cmp((func() _dafny.Map {
			var _coll25 = _dafny.NewMapBuilder()
			_ = _coll25
			for _iter29 := _dafny.Iterate((_797_v75).Keys().Elements()); ; {
				_compr_25, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _801_v72 _dafny.Sequence
				_801_v72 = interface{}(_compr_25).(_dafny.Sequence)
				if (_797_v75).Contains(_801_v72) {
					_coll25.Add(_801_v72, p2)
				}
			}
			return _coll25.ToMap()
		}()).Cardinality()) <= 0, 11)
		(_nw148).ArraySet1(p2, 12)
		(_nw148).ArraySet1((_793_v70).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_793_v70).Cardinality()))).Uint32()).(bool), 13)
		(_nw148).ArraySet1((_dafny.IntOfInt64(684)).Cmp(_this.F2()) == 0, 14)
		(_nw148).ArraySet1((p2) || (p2), 15)
		(_nw148).ArraySet1(p2, 16)
		(_nw148).ArraySet1(p2, 17)
		(_nw148).ArraySet1((p2) == (p2), 18)
		_798_v76 = _nw148
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
		_ = _index128
		(_798_v76).ArraySet1(p2, (_index128).Int())
		var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
		_ = _index129
		(_798_v76).ArraySet1(p2, (_index129).Int())
		var _802_i7 _dafny.Int
		_ = _802_i7
		_802_i7 = _dafny.Zero
		{
			for !(true) || (false) {
				{
					if (_802_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_802_i7 = (_802_i7).Plus(_dafny.One)
					var _803_v77 _dafny.Map
					_ = _803_v77
					_803_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
						if (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool) {
							return _791_v68
						}
						return _791_v68
					})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), p0))
					var _804_v78 _dafny.MultiSet
					_ = _804_v78
					_804_v78 = _dafny.MultiSetOf(p2)
					var _805_v79 _dafny.Map
					_ = _805_v79
					_805_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_804_v78).Cardinality())
					_803_v77 = (_803_v77).Update(_791_v68, _805_v79)
					if ((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)) && ((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)) {
						r0 = Companion_Default___.SafeModuloInt(_this.F2(), p1)
						r0 = (_dafny.Zero).Minus(p0)
						var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index130
						(_798_v76).ArraySet1(p2, (_index130).Int())
						var _806_v81 _dafny.Set
						_ = _806_v81
						_806_v81 = _dafny.SetOf(p1, _this.F2())
						var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index131
						(_798_v76).ArraySet1(((func() _dafny.Set {
							var _coll26 = _dafny.NewBuilder()
							_ = _coll26
							for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfInt64(-727))); ; {
								_compr_26, _ok30 := _iter30()
								if !_ok30 {
									break
								}
								var _807_v80 _dafny.Int
								_807_v80 = interface{}(_compr_26).(_dafny.Int)
								if ((_807_v80).Sign() != -1) && ((_807_v80).Cmp(_dafny.IntOfInt64(-727)) < 0) {
									_coll26.Add((_807_v80).Plus(p0))
								}
							}
							return _coll26.ToSet()
						}()).Union(_806_v81)).IsSubsetOf(Companion_Default___.Fm31(!((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)), p0, p2, p0, globalState)), (_index131).Int())
						r0 = _dafny.IntOfInt64(-962)
					} else {
						_804_v78 = _804_v78
						var _808_v82 _dafny.MultiSet
						_ = _808_v82
						_808_v82 = _dafny.MultiSetOf(_790_v67)
						var _809_v83 D0
						_ = _809_v83
						_809_v83 = Companion_D0_.Create_DC1_((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), (_808_v82).IsDisjointFrom(_808_v82), _791_v68, _this.F2())
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index132
						var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index133
						var _rhs116 D0 = _809_v83
						_ = _rhs116
						var _rhs117 bool = !(!((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)))
						_ = _rhs117
						var _rhs118 _dafny.CodePoint = _dafny.CodePoint('a')
						_ = _rhs118
						var _rhs119 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_790_v67, _790_v67), _790_v67)
						_ = _rhs119
						var _lhs56 _dafny.Array = _798_v76
						_ = _lhs56
						var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _lhs57
						var _lhs58 _dafny.Array = _798_v76
						_ = _lhs58
						var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _lhs59
						_809_v83 = _rhs116
						(_lhs56).ArraySet1(_rhs117, (_lhs57).Int())
						_791_v68 = _rhs118
						(_lhs58).ArraySet1(_rhs119, (_lhs59).Int())
						var _810_v84 D7
						_ = _810_v84
						_810_v84 = Companion_D7_.Create_DC14_(_796_v74)
						var _811_v85 _dafny.Map
						_ = _811_v85
						_811_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(Companion_Default___.Fm32(globalState))).Intersection(_dafny.MultiSetOf(_810_v84)), _dafny.IntOfUint32((_790_v67).Cardinality()))
						var _812_v86 _dafny.Map
						_ = _812_v86
						_812_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), p1)
						var _813_v87 _dafny.Set
						_ = _813_v87
						_813_v87 = _dafny.SetOf((_dafny.SetOf(p0, _dafny.IntOfInt64(264), _dafny.IntOfUint32((_790_v67).Cardinality()), (func() _dafny.Int {
							if (_805_v79).Contains((_this).Fm1(p2, globalState)) {
								return (_805_v79).Get((_this).Fm1(p2, globalState)).(_dafny.Int)
							}
							return (_812_v86).Cardinality()
						})(), _dafny.IntOfInt64(462))).Cardinality())
						var _814_v88 *C1
						_ = _814_v88
						var _nw149 *C1 = New_C1_()
						_ = _nw149
						_nw149.Ctor__(_790_v67)
						_814_v88 = _nw149
						var _815_v89 _dafny.Map
						_ = _815_v89
						_815_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _814_v88)
						var _816_v90 D8
						_ = _816_v90
						_816_v90 = Companion_D8_.Create_DC16_(_814_v88)
						_811_v85 = (_811_v85).Update(Companion_Default___.Fm33(_791_v68, true, (_813_v87).Cardinality(), p1, globalState), _dafny.IntOfUint32((_dafny.SeqOf((_815_v89).Update(p2, (_816_v90).Dtor_cf29()), _815_v89, _815_v89, (_815_v89).Update((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), _814_v88), _815_v89)).Cardinality()))
						r0 = _this.F2()
						var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index134
						(_798_v76).ArraySet1((_809_v83).Dtor_cf1(), (_index134).Int())
					}
					if p2 {
						var _817_v91 _dafny.Map
						_ = _817_v91
						_817_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool))
						var _818_v92 _dafny.Map
						_ = _818_v92
						_818_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _791_v68)
						var _819_v93 _dafny.MultiSet
						_ = _819_v93
						_819_v93 = _dafny.MultiSetOf(p0, ((_818_v92).Update(_this.F2(), _791_v68)).Cardinality(), _this.F2())
						_817_v91 = (_817_v91).Update((func() _dafny.Int {
							if (_819_v93).Contains(_this.F2()) {
								return (_819_v93).Multiplicity(_this.F2())
							}
							return p1
						})(), (_819_v93).IsProperSubsetOf(_819_v93))
						var _820_v94 _dafny.Map
						_ = _820_v94
						_820_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), _798_v76)
						var _821_v95 _dafny.Set
						_ = _821_v95
						_821_v95 = _dafny.SetOf((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ssamx")).Cardinality())).Plus((_820_v94).Cardinality()))
						_821_v95 = ((func() _dafny.Set {
							if p2 {
								return _821_v95
							}
							return func() _dafny.Set {
								var _coll27 = _dafny.NewBuilder()
								_ = _coll27
								for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-95), _dafny.IntOfInt64(219))); ; {
									_compr_27, _ok31 := _iter31()
									if !_ok31 {
										break
									}
									var _822_v96 _dafny.Int
									_822_v96 = interface{}(_compr_27).(_dafny.Int)
									if ((_dafny.IntOfInt64(-95)).Cmp(_822_v96) <= 0) && ((_822_v96).Cmp(_dafny.IntOfInt64(219)) < 0) {
										_coll27.Add((_822_v96).Times(_dafny.IntOfInt64(685)))
									}
								}
								return _coll27.ToSet()
							}()
						})()).Union((_dafny.SetOf(p0, p0)).Union(_821_v95))
						(_this).F2_set_(_this.F2())
						(_this).F2_set_(_this.F2())
						var _rhs120 _dafny.Sequence = _790_v67
						_ = _rhs120
						var _rhs121 _dafny.Int = p1
						_ = _rhs121
						var _rhs122 _dafny.Int = (_this).Fm1((false) && ((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)), globalState)
						_ = _rhs122
						_790_v67 = _rhs120
						r0 = _rhs121
						r0 = _rhs122
					} else {
						var _823_v97 _dafny.Array
						_ = _823_v97
						var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
						_ = _nw150
						_823_v97 = _nw150
						var _824_v98 _dafny.Array
						_ = _824_v98
						var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
						_ = _nw151
						_824_v98 = _nw151
						var _825_v99 _dafny.Sequence
						_ = _825_v99
						_825_v99 = _dafny.SeqOf(_824_v98, _824_v98)
						var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_823_v97), 0))
						_ = _index135
						(_823_v97).ArraySet1(_825_v99, (_index135).Int())
						var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_823_v97), 0))
						_ = _index136
						(_823_v97).ArraySet1(_825_v99, (_index136).Int())
						var _826_v100 D6
						_ = _826_v100
						_826_v100 = Companion_D6_.Create_DC13_((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), p1)
						var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
						_ = _index137
						(_798_v76).ArraySet1(((func() D6 {
							if false {
								return _826_v100
							}
							return _826_v100
						})()).Dtor_cf24(), (_index137).Int())
						var _827_v101 _dafny.Map
						_ = _827_v101
						_827_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p0, p1), _824_v98)
						var _828_v102 _dafny.Map
						_ = _828_v102
						_828_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_793_v70).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_793_v70).Cardinality()))).Uint32()).(bool), _827_v101)
						_827_v101 = (func() _dafny.Map {
							if (_828_v102).Contains((_this).Fm8(_790_v67, p2, globalState)) {
								return (_828_v102).Get((_this).Fm8(_790_v67, p2, globalState)).(_dafny.Map)
							}
							return _827_v101
						})()
						_790_v67 = _dafny.UnicodeSeqOfUtf8Bytes("weab")
						var _829_v103 _dafny.Array
						_ = _829_v103
						var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
						_ = _nw152
						_829_v103 = _nw152
						var _830_v104 _dafny.Map
						_ = _830_v104
						_830_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), _824_v98)
						var _nwElement0_37 _dafny.Array = _824_v98
						_ = _nwElement0_37
						var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(23))
						_ = _nw153
						(_nw153).ArraySet1(_nwElement0_37, 0)
						(_nw153).ArraySet1(_824_v98, 1)
						(_nw153).ArraySet1(_824_v98, 2)
						(_nw153).ArraySet1(_824_v98, 3)
						(_nw153).ArraySet1(_824_v98, 4)
						(_nw153).ArraySet1((func() _dafny.Array {
							if (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool) {
								return _824_v98
							}
							return _824_v98
						})(), 5)
						(_nw153).ArraySet1(_824_v98, 6)
						(_nw153).ArraySet1(_824_v98, 7)
						(_nw153).ArraySet1(_824_v98, 8)
						(_nw153).ArraySet1(_824_v98, 9)
						(_nw153).ArraySet1(_824_v98, 10)
						(_nw153).ArraySet1(_824_v98, 11)
						(_nw153).ArraySet1(_824_v98, 12)
						(_nw153).ArraySet1(_824_v98, 13)
						(_nw153).ArraySet1(_824_v98, 14)
						(_nw153).ArraySet1(_824_v98, 15)
						(_nw153).ArraySet1(_824_v98, 16)
						(_nw153).ArraySet1((func() _dafny.Array {
							if (_830_v104).Contains((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)) {
								return (_830_v104).Get((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)).(_dafny.Array)
							}
							return _824_v98
						})(), 17)
						(_nw153).ArraySet1(_824_v98, 18)
						(_nw153).ArraySet1(_824_v98, 19)
						(_nw153).ArraySet1((func() _dafny.Array {
							if (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool) {
								return _824_v98
							}
							return _824_v98
						})(), 20)
						(_nw153).ArraySet1(_824_v98, 21)
						(_nw153).ArraySet1(_824_v98, 22)
						_829_v103 = _nw153
					}
					r0 = p1
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _831_v105 *C1
		_ = _831_v105
		var _nw154 *C1 = New_C1_()
		_ = _nw154
		_nw154.Ctor__(_790_v67)
		_831_v105 = _nw154
		var _832_v106 D8
		_ = _832_v106
		_832_v106 = Companion_D8_.Create_DC16_(_831_v105)
		var _source16 D8 = _832_v106
		_ = _source16
		if _source16.Is_DC17() {
			var _833___mcc_h8 bool = _source16.Get_().(D8_DC17).Cf30
			_ = _833___mcc_h8
			var _834___mcc_h9 *C1 = _source16.Get_().(D8_DC17).Cf31
			_ = _834___mcc_h9
			var _835___mcc_h10 _dafny.Int = _source16.Get_().(D8_DC17).Cf32
			_ = _835___mcc_h10
			var _836_cf32 _dafny.Int = _835___mcc_h10
			_ = _836_cf32
			var _837_cf31 *C1 = _834___mcc_h9
			_ = _837_cf31
			var _838_cf30 bool = _833___mcc_h8
			_ = _838_cf30
			var _839_v107 _dafny.Map
			_ = _839_v107
			_839_v107 = _795_v73
			var _840_v108 _dafny.Map
			_ = _840_v108
			_840_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v107, p2)
			var _841_v109 _dafny.Set
			_ = _841_v109
			_841_v109 = _dafny.SetOf((_840_v108).Merge((_840_v108).Update(_839_v107, !(p2))), (_840_v108).Update(_839_v107, _838_cf30))
			var _842_v110 _dafny.Map
			_ = _842_v110
			_842_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v108, p1)
			_841_v109 = (func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter32 := _dafny.Iterate((_842_v110).Keys().Elements()); ; {
					_compr_28, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _843_v111 _dafny.Map
					_843_v111 = interface{}(_compr_28).(_dafny.Map)
					if (_842_v110).Contains(_843_v111) {
						_coll28.Add(_843_v111)
					}
				}
				return _coll28.ToSet()
			}()).Difference(_dafny.SetOf(_840_v108, _840_v108))
			var _844_v112 _dafny.Array
			_ = _844_v112
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_12
			var _nw155 _dafny.Array
			_ = _nw155
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw155 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) D7 = (func(_845_v68 _dafny.CodePoint) func(_dafny.Int) D7 {
					return func(_846_i8 _dafny.Int) D7 {
						return Companion_D7_.Create_DC15_(_this.F2(), _845_v68)
					}
				})(_791_v68)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw155 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw155).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw155).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_844_v112 = _nw155
			_844_v112 = (func() _dafny.Array {
				if p2 {
					return _844_v112
				}
				return _844_v112
			})()
			_838_cf30 = _838_cf30
			var _847_v113 *C1
			_ = _847_v113
			var _nw156 *C1 = New_C1_()
			_ = _nw156
			_nw156.Ctor__(_790_v67)
			_847_v113 = _nw156
		} else {
			var _848___mcc_h11 *C1 = _source16.Get_().(D8_DC16).Cf29
			_ = _848___mcc_h11
			var _849_cf29 *C1 = _848___mcc_h11
			_ = _849_cf29
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _index138
			(_798_v76).ArraySet1(((p1).Cmp(p1) > 0) && (!((_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool))), (_index138).Int())
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _index139
			(_798_v76).ArraySet1(p2, (_index139).Int())
			var _850_v114 _dafny.Set
			_ = _850_v114
			_850_v114 = _dafny.SetOf(_this.F2())
			var _851_v115 _dafny.Array
			_ = _851_v115
			var _nwElement0_38 _dafny.Int = _dafny.IntOfInt64(600)
			_ = _nwElement0_38
			var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(2))
			_ = _nw157
			(_nw157).ArraySet1(_nwElement0_38, 0)
			(_nw157).ArraySet1((_850_v114).Cardinality(), 1)
			_851_v115 = _nw157
			var _852_v116 D4
			_ = _852_v116
			_852_v116 = Companion_D4_.Create_DC6_(p1, _851_v115, p1)
			var _source17 D4 = _852_v116
			_ = _source17
			if _source17.Is_DC6() {
				var _853___mcc_h12 _dafny.Int = _source17.Get_().(D4_DC6).Cf9
				_ = _853___mcc_h12
				var _854___mcc_h13 _dafny.Array = _source17.Get_().(D4_DC6).Cf10
				_ = _854___mcc_h13
				var _855___mcc_h14 _dafny.Int = _source17.Get_().(D4_DC6).Cf11
				_ = _855___mcc_h14
				var _856_cf11 _dafny.Int = _855___mcc_h14
				_ = _856_cf11
				var _857_cf10 _dafny.Array = _854___mcc_h13
				_ = _857_cf10
				var _858_cf9 _dafny.Int = _853___mcc_h12
				_ = _858_cf9
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
				_ = _index140
				(_798_v76).ArraySet1((_this).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("yjjaoweep"), (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool), globalState), (_index140).Int())
				var _859_v117 _dafny.Map
				_ = _859_v117
				_859_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool))
				var _860_v118 *C0
				_ = _860_v118
				var _nw158 *C0 = New_C0_()
				_ = _nw158
				_nw158.Ctor__(p2, (func() bool {
					if (_859_v117).Contains(_this.F2()) {
						return (_859_v117).Get(_this.F2()).(bool)
					}
					return (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)
				})())
				_860_v118 = _nw158
				var _861_v119 _dafny.Map
				_ = _861_v119
				_861_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_860_v118, _795_v73)
				var _862_v120 _dafny.Map
				_ = _862_v120
				_862_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_795_v73, _796_v74)
				_858_cf9 = (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					if (_861_v119).Contains(_860_v118) {
						return (_861_v119).Get(_860_v118).(_dafny.Map)
					}
					return _795_v73
				})(), _796_v74)).Merge(_862_v120)).Cardinality()).Minus((_dafny.Zero).Minus(_this.F2())))
				_859_v117 = (_859_v117).Update(Companion_Default___.SafeDivisionInt(p0, p0), _860_v118.F5)
				var _863_v121 _dafny.Set
				_ = _863_v121
				_863_v121 = _dafny.SetOf(_860_v118.F5, (_831_v105).Fm11(globalState))
				var _864_v122 _dafny.Map
				_ = _864_v122
				_864_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_860_v118).Fm14((_831_v105).Fm11(globalState), p2, _850_v114, globalState), Companion_D6_.Create_DC11_(Companion_Default___.Fm28(true, p2, _860_v118.F4, _863_v121, globalState)))
				var _865_v124 D6
				_ = _865_v124
				_865_v124 = Companion_D6_.Create_DC11_((_831_v105).F3())
				_864_v122 = (_864_v122).Update((func() _dafny.Set {
					var _coll29 = _dafny.NewBuilder()
					_ = _coll29
					for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(461), _dafny.IntOfInt64(893))); ; {
						_compr_29, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _866_v123 _dafny.Int
						_866_v123 = interface{}(_compr_29).(_dafny.Int)
						if ((_dafny.IntOfInt64(461)).Cmp(_866_v123) <= 0) && ((_866_v123).Cmp(_dafny.IntOfInt64(893)) < 0) {
							_coll29.Add((_866_v123).Minus(p0))
						}
					}
					return _coll29.ToSet()
				}()).IsProperSubsetOf(_850_v114), _865_v124)
			} else if _source17.Is_DC5() {
				var _867___mcc_h15 _dafny.Set = _source17.Get_().(D4_DC5).Cf8
				_ = _867___mcc_h15
				var _868_cf8 _dafny.Set = _867___mcc_h15
				_ = _868_cf8
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
				_ = _index141
				(_798_v76).ArraySet1(p2, (_index141).Int())
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw159
				_851_v115 = _nw159
				(_this).F2_set_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qnoyfxqrs")).Cardinality()))
				var _869_v125 _dafny.MultiSet
				_ = _869_v125
				_869_v125 = _dafny.MultiSetOf(_851_v115, _851_v115)
				var _870_v126 D5
				_ = _870_v126
				_870_v126 = Companion_D5_.Create_DC8_(_869_v125)
				var _871_v127 _dafny.Map
				_ = _871_v127
				_871_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _870_v126)
				var _872_v128 _dafny.Sequence
				_ = _872_v128
				_872_v128 = _dafny.SeqOf(_871_v127)
				r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_871_v127), _dafny.Companion_Sequence_.Update(_872_v128, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_872_v128).Cardinality()))).Uint32(), _871_v127))).Cardinality())
			} else {
				var _873___mcc_h16 D4 = _source17.Get_().(D4_DC7).Cf12
				_ = _873___mcc_h16
				var _874_cf12 D4 = _873___mcc_h16
				_ = _874_cf12
				var _875_v129 _dafny.Map
				_ = _875_v129
				_875_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_791_v68, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), p0)).Cardinality()).Plus(_this.F2()))
				_875_v129 = (_875_v129).Update(_791_v68, _this.F2())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
				_ = _index142
				(_798_v76).ArraySet1((_831_v105).Fm11(globalState), (_index142).Int())
				(_this).F2_set_(p0)
				_849_cf29 = _849_cf29
			}
			var _876_v130 _dafny.Map
			_ = _876_v130
			_876_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v105).F3(), p0)
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _index143
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _index144
			var _rhs123 D8 = _832_v106
			_ = _rhs123
			var _rhs124 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yac"), _this.F2())
			_ = _rhs124
			var _rhs125 bool = (p0).Cmp(p1) == 0
			_ = _rhs125
			var _rhs126 _dafny.Array = _851_v115
			_ = _rhs126
			var _rhs127 bool = (_798_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))).Int()).(bool)
			_ = _rhs127
			var _lhs60 _dafny.Array = _798_v76
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _lhs61
			var _lhs62 _dafny.Array = _798_v76
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_798_v76), 0))
			_ = _lhs63
			_832_v106 = _rhs123
			_876_v130 = _rhs124
			(_lhs60).ArraySet1(_rhs125, (_lhs61).Int())
			_851_v115 = _rhs126
			(_lhs62).ArraySet1(_rhs127, (_lhs63).Int())
		}
		var _877_v131 _dafny.Map
		_ = _877_v131
		_877_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}((func(_878_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_879_i9 _dafny.Int) _dafny.CodePoint {
				return _878_v68
			}
		})(_791_v68)))).Cardinality()), _this.F2()))
		var _880_v132 _dafny.Array
		_ = _880_v132
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_13
		var _nw160 _dafny.Array
		_ = _nw160
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw160 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Int = (func(_881_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_882_i11 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_882_i11, _881_p0)
				}
			})(p0)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw160 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw160).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw160).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_880_v132 = _nw160
		r0 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_877_v131).Contains(p2) {
				return (_877_v131).Get(p2).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(482))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_883_i10 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(73)
			}))
		})()).Cardinality()), _880_v132)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F2())).Cardinality(), (_831_v105).Fm11(globalState))).Cardinality())).Cardinality()))
		return r0
	}
}
func (_this *C2) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		if p3 {
			var _884_v0 bool
			_ = _884_v0
			_884_v0 = true
			_884_v0 = _884_v0
			var _885_v1 _dafny.Map
			_ = _885_v1
			_885_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F2())
			var _886_v2 _dafny.Sequence
			_ = _886_v2
			_886_v2 = _dafny.SeqOf(p3)
			_885_v1 = (_885_v1).Update(p1, (_this.F2()).Minus((_dafny.MultiSetFromSeq(_886_v2)).Cardinality()))
			var _887_v3 _dafny.CodePoint
			_ = _887_v3
			_887_v3 = _dafny.CodePoint('m')
			var _888_v4 _dafny.Sequence
			_ = _888_v4
			_888_v4 = _dafny.SeqOf(_dafny.CodePoint('g'), _887_v3)
			var _889_v5 _dafny.Map
			_ = _889_v5
			_889_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_888_v4).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_888_v4).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _890_v6 _dafny.Set
			_ = _890_v6
			_890_v6 = _dafny.SetOf(_889_v5, _889_v5)
			var _891_v7 _dafny.Set
			_ = _891_v7
			_891_v7 = _dafny.SetOf(_884_v0, p3)
			var _892_v8 _dafny.Map
			_ = _892_v8
			_892_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_v5, _891_v7)
			var _893_v9 _dafny.Map
			_ = _893_v9
			_893_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_890_v6, _892_v8)
			_893_v9 = (_893_v9).Update((_890_v6).Difference(_dafny.SetOf(_889_v5, _889_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _887_v3), _889_v5, _889_v5)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_v5, _891_v7))
			var _894_v10 _dafny.Set
			_ = _894_v10
			_894_v10 = _dafny.SetOf(_this.F2(), _this.F2())
			if (_894_v10).IsSubsetOf(_894_v10) {
				_886_v2 = _dafny.SeqOf(p2, _884_v0, p3, p3, !(p3))
				var _895_v11 _dafny.Array
				_ = _895_v11
				var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw161
				_895_v11 = _nw161
				var _896_v12 _dafny.Sequence
				_ = _896_v12
				_896_v12 = _dafny.SeqOf(_888_v4, _888_v4, _888_v4)
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_895_v11), 0))
				_ = _index145
				(_895_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_896_v12).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_896_v12).Cardinality()))).Uint32()).(_dafny.Sequence), _888_v4), (_index145).Int())
				var _897_v13 _dafny.Array
				_ = _897_v13
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw162
				_897_v13 = _nw162
				var _898_v14 _dafny.MultiSet
				_ = _898_v14
				_898_v14 = _dafny.MultiSetOf(p2)
				var _899_v15 _dafny.Map
				_ = _899_v15
				_899_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1(_884_v0, globalState), (_898_v14).Cardinality())
				var _900_v16 _dafny.MultiSet
				_ = _900_v16
				_900_v16 = _dafny.MultiSetOf(_899_v15, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p2, _884_v0)).Cardinality()), _this.F2()))
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))
				_ = _index146
				(_897_v13).ArraySet1((_900_v16).IsProperSubsetOf(_900_v16), (_index146).Int())
				var _901_v17 _dafny.Sequence
				_ = _901_v17
				_901_v17 = _dafny.SeqOf((func() _dafny.Int {
					if (_898_v14).Contains(p2) {
						return (_898_v14).Multiplicity(p2)
					}
					return _this.F2()
				})(), _this.F2(), _this.F2(), _dafny.IntOfInt64(256), _this.F2())
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_895_v11), 0))
				_ = _index147
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))
				_ = _index148
				var _rhs128 bool = _dafny.Companion_Sequence_.IsPrefixOf(_901_v17, _901_v17)
				_ = _rhs128
				var _rhs129 _dafny.Sequence = _888_v4
				_ = _rhs129
				var _rhs130 _dafny.Int = (_this.F2()).Minus((_dafny.Zero).Minus(_this.F2()))
				_ = _rhs130
				var _rhs131 _dafny.Int = (_this).Fm1((func() bool {
					if true {
						return false
					}
					return p3
				})(), globalState)
				_ = _rhs131
				var _rhs132 bool = _884_v0
				_ = _rhs132
				var _lhs64 _dafny.Array = _895_v11
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_895_v11), 0))
				_ = _lhs65
				var _lhs66 *C2 = _this
				_ = _lhs66
				var _lhs67 _dafny.Array = _897_v13
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))
				_ = _lhs68
				_884_v0 = _rhs128
				(_lhs64).ArraySet1(_rhs129, (_lhs65).Int())
				_lhs66.F2_set_(_rhs130)
				r0 = _rhs131
				(_lhs67).ArraySet1(_rhs132, (_lhs68).Int())
				var _902_v18 *C0
				_ = _902_v18
				var _nw163 *C0 = New_C0_()
				_ = _nw163
				_nw163.Ctor__((_this).Fm8(Companion_Default___.Fm28((_897_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))).Int()).(bool), p3, p2, _891_v7, globalState), p2, globalState), (_886_v2).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_886_v2).Cardinality()))).Uint32()).(bool))
				_902_v18 = _nw163
				_897_v13 = _897_v13
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))
				_ = _index149
				(_897_v13).ArraySet1((func() bool {
					if (_897_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))).Int()).(bool) {
						return (_897_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_897_v13), 0))).Int()).(bool)
					}
					return (_this.F2()).Cmp(_this.F2()) < 0
				})(), (_index149).Int())
			} else {
				var _903_v19 _dafny.Array
				_ = _903_v19
				var _nw164 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
				_ = _nw164
				_903_v19 = _nw164
				var _904_v20 *C0
				_ = _904_v20
				var _nw165 *C0 = New_C0_()
				_ = _nw165
				_nw165.Ctor__(p2, (_886_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_888_v4).Cardinality()), _dafny.IntOfUint32((_886_v2).Cardinality()))).Uint32()).(bool))
				_904_v20 = _nw165
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_903_v19), 0))
				_ = _index150
				(_903_v19).ArraySet1(_904_v20, (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_903_v19), 0))
				_ = _index151
				(_903_v19).ArraySet1((func() *C0 {
					if _904_v20.F4 {
						return _904_v20
					}
					return _904_v20
				})(), (_index151).Int())
				r1 = (_dafny.Zero).Minus(_this.F2())
				var _905_v22 D5
				_ = _905_v22
				_905_v22 = Companion_D5_.Create_DC10_((_904_v20).Fm14(_884_v0, _904_v20.F4, _894_v10, globalState), _this.F2(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_888_v4, func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(338), _dafny.IntOfInt64(202))); ; {
						_compr_30, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _906_v21 _dafny.Int
						_906_v21 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(338)).Cmp(_906_v21) <= 0) && ((_906_v21).Cmp(_dafny.IntOfInt64(202)) < 0) {
							_coll30.Add(Companion_Default___.SafeModuloInt(_906_v21, (func() _dafny.Int {
								if (p0).Contains(_this.F2()) {
									return (p0).Multiplicity(_this.F2())
								}
								return _dafny.IntOfInt64(198)
							})()), _904_v20.F5)
						}
					}
					return _coll30.ToMap()
				}()), p2)
				_905_v22 = _905_v22
				var _907_v23 *C0
				_ = _907_v23
				var _nw166 *C0 = New_C0_()
				_ = _nw166
				_nw166.Ctor__(_884_v0, _884_v0)
				_907_v23 = _nw166
				(_907_v23).F4 = (_907_v23).Fm14(_907_v23.F5, (_dafny.MultiSetOf(p2, _884_v0)).IsSubsetOf(_dafny.MultiSetFromSeq(_886_v2)), (_894_v10).Difference(_894_v10), globalState)
			}
			var _908_v24 _dafny.Map
			_ = _908_v24
			_908_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F2()), _884_v0)
			var _909_v25 _dafny.Map
			_ = _909_v25
			_909_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("hrjp"), _908_v24)
			var _910_v26 D5
			_ = _910_v26
			_910_v26 = Companion_D5_.Create_DC10_(!(p2), _this.F2(), _909_v25, p2)
			_884_v0 = (func() bool {
				if p2 {
					return ((_910_v26).Dtor_cf18()).Cmp(_dafny.IntOfInt64(876)) <= 0
				}
				return p2
			})()
		} else {
			var _911_v27 bool
			_ = _911_v27
			_911_v27 = true
			_911_v27 = !(p3) || (_911_v27)
			var _912_v28 _dafny.CodePoint
			_ = _912_v28
			_912_v28 = _dafny.CodePoint('v')
			var _913_v29 _dafny.MultiSet
			_ = _913_v29
			_913_v29 = _dafny.MultiSetOf(_912_v28, _912_v28)
			var _914_v30 _dafny.Sequence
			_ = _914_v30
			_914_v30 = _dafny.UnicodeSeqOfUtf8Bytes("h")
			var _915_v31 _dafny.Array
			_ = _915_v31
			var _nwElement0_39 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("gmqami")
			_ = _nwElement0_39
			var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(7))
			_ = _nw167
			(_nw167).ArraySet1(_nwElement0_39, 0)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_914_v30, _914_v30), 1)
			(_nw167).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tuqtafex"), 2)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("g"), _914_v30), 3)
			(_nw167).ArraySet1(_914_v30, 4)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_914_v30, _914_v30), 5)
			(_nw167).ArraySet1(_914_v30, 6)
			_915_v31 = _nw167
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_915_v31), 0))
			_ = _index152
			(_915_v31).ArraySet1(_914_v30, (_index152).Int())
			var _916_v32 _dafny.Array
			_ = _916_v32
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_14
			var _nw168 _dafny.Array
			_ = _nw168
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw168 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_917_p3 bool) func(_dafny.Int) bool {
					return func(_918_i0 _dafny.Int) bool {
						return _917_p3
					}
				})(p3)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw168 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw168).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw168).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_916_v32 = _nw168
			var _919_v33 _dafny.Map
			_ = _919_v33
			_919_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_916_v32, _dafny.MultiSetOf(_this.F2(), _this.F2()))
			var _920_v34 _dafny.Map
			_ = _920_v34
			_920_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_913_v29).Intersection(_913_v29))
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_915_v31), 0))
			_ = _index153
			var _rhs133 _dafny.MultiSet = (func() _dafny.MultiSet {
				if (_920_v34).Contains(_this.F2()) {
					return (_920_v34).Get(_this.F2()).(_dafny.MultiSet)
				}
				return (_dafny.MultiSetOf(_dafny.CodePoint('u'), _912_v28)).Update(_dafny.CodePoint('d'), Companion_Default___.Abs(_this.F2()))
			})()
			_ = _rhs133
			var _rhs134 bool = p2
			_ = _rhs134
			var _rhs135 _dafny.Sequence = _914_v30
			_ = _rhs135
			var _rhs136 _dafny.Map = _919_v33
			_ = _rhs136
			var _rhs137 _dafny.Int = (_this).Fm1(false, globalState)
			_ = _rhs137
			var _lhs69 _dafny.Array = _915_v31
			_ = _lhs69
			var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_915_v31), 0))
			_ = _lhs70
			var _lhs71 *C2 = _this
			_ = _lhs71
			_913_v29 = _rhs133
			_911_v27 = _rhs134
			(_lhs69).ArraySet1(_rhs135, (_lhs70).Int())
			_919_v33 = _rhs136
			_lhs71.F2_set_(_rhs137)
			var _921_v35 _dafny.Set
			_ = _921_v35
			_921_v35 = _dafny.SetOf(_this.F2())
			var _922_v36 _dafny.MultiSet
			_ = _922_v36
			_922_v36 = _dafny.MultiSetOf(!(_911_v27))
			var _923_v37 _dafny.Sequence
			_ = _923_v37
			_923_v37 = _dafny.SeqOf(_922_v36, _922_v36, _922_v36, _922_v36)
			var _924_v38 _dafny.Map
			_ = _924_v38
			_924_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_921_v35, (_923_v37).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_923_v37).Cardinality()))).Uint32()).(_dafny.MultiSet))
			_924_v38 = (_924_v38).Update(_921_v35, _dafny.MultiSetOf((_this).Fm8((_915_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_915_v31), 0))).Int()).(_dafny.Sequence), p3, globalState)))
			var _925_v39 _dafny.Array
			_ = _925_v39
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw169
			_925_v39 = _nw169
			var _926_v40 _dafny.Array
			_ = _926_v40
			var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw170
			_926_v40 = _nw170
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_925_v39), 0))
			_ = _index154
			(_925_v39).ArraySet1(_926_v40, (_index154).Int())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_916_v32), 0))
			_ = _index155
			(_916_v32).ArraySet1(p3, (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((p1), 0))
			_ = _index156
			(p1).ArraySet1(_this.F2(), (_index156).Int())
			var _927_v41 _dafny.Map
			_ = _927_v41
			_927_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8((_915_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_915_v31), 0))).Int()).(_dafny.Sequence), _911_v27, globalState), _926_v40)
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_925_v39), 0))
			_ = _index157
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_916_v32), 0))
			_ = _index158
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((p1), 0))
			_ = _index159
			var _rhs138 _dafny.Int = (_this.F2()).Minus(_this.F2())
			_ = _rhs138
			var _rhs139 _dafny.Array = (func() _dafny.Array {
				if (_927_v41).Contains(_911_v27) {
					return (_927_v41).Get(_911_v27).(_dafny.Array)
				}
				return _926_v40
			})()
			_ = _rhs139
			var _rhs140 bool = _911_v27
			_ = _rhs140
			var _rhs141 _dafny.Int = _this.F2()
			_ = _rhs141
			var _lhs72 *C2 = _this
			_ = _lhs72
			var _lhs73 _dafny.Array = _925_v39
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_925_v39), 0))
			_ = _lhs74
			var _lhs75 _dafny.Array = _916_v32
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_916_v32), 0))
			_ = _lhs76
			var _lhs77 _dafny.Array = p1
			_ = _lhs77
			var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((p1), 0))
			_ = _lhs78
			_lhs72.F2_set_(_rhs138)
			(_lhs73).ArraySet1(_rhs139, (_lhs74).Int())
			(_lhs75).ArraySet1(_rhs140, (_lhs76).Int())
			(_lhs77).ArraySet1(_rhs141, (_lhs78).Int())
			var _928_v42 *C1
			_ = _928_v42
			var _nw171 *C1 = New_C1_()
			_ = _nw171
			_nw171.Ctor__(_914_v30)
			_928_v42 = _nw171
			var _929_v43 _dafny.Map
			_ = _929_v43
			_929_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_912_v28, _928_v42)
			_929_v43 = (_929_v43).Update(_912_v28, _928_v42)
		}
		var _930_v44 _dafny.Array
		_ = _930_v44
		var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw172
		_930_v44 = _nw172
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _index160
		(_930_v44).ArraySet1(!(p3), (_index160).Int())
		var _931_v45 D6
		_ = _931_v45
		_931_v45 = Companion_D6_.Create_DC13_(p2, _this.F2())
		var _932_v46 _dafny.Sequence
		_ = _932_v46
		_932_v46 = _dafny.UnicodeSeqOfUtf8Bytes("hgoffmr")
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _index161
		var _rhs142 bool = (func() bool {
			if !((p3) || ((_931_v45).Dtor_cf24())) {
				return _dafny.Companion_Sequence_.IsPrefixOf(_932_v46, _dafny.UnicodeSeqOfUtf8Bytes("rabeglte"))
			}
			return p2
		})()
		_ = _rhs142
		var _rhs143 _dafny.Int = _this.F2()
		_ = _rhs143
		var _lhs79 _dafny.Array = _930_v44
		_ = _lhs79
		var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _lhs80
		(_lhs79).ArraySet1(_rhs142, (_lhs80).Int())
		r1 = _rhs143
		var _933_v47 _dafny.Sequence
		_ = _933_v47
		_933_v47 = _dafny.SeqOf(p3)
		var _934_v48 _dafny.Map
		_ = _934_v48
		_934_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _933_v47)
		var _935_i1 _dafny.Int
		_ = _935_i1
		_935_i1 = _dafny.Zero
		{
			for !_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
				if (_934_v48).Contains(_this.F2()) {
					return (_934_v48).Get(_this.F2()).(_dafny.Sequence)
				}
				return _933_v47
			})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)), _933_v47)) {
				{
					if (_935_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_935_i1 = (_935_i1).Plus(_dafny.One)
					var _936_v49 _dafny.Map
					_ = _936_v49
					_936_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _this.F2())
					var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
					_ = _index162
					(p1).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_936_v49, (_dafny.Zero).Minus(_this.F2()))).Cardinality(), (_index162).Int())
					var _937_v50 _dafny.CodePoint
					_ = _937_v50
					_937_v50 = _dafny.CodePoint('v')
					var _938_v51 _dafny.Map
					_ = _938_v51
					_938_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_937_v50, _this.F2())
					var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
					_ = _index163
					(p1).ArraySet1((_938_v51).Cardinality(), (_index163).Int())
					if (func() bool {
						if !(p3) {
							return (Companion_Default___.Fm34(p3, globalState)).Contains(_dafny.UnicodeSeqOfUtf8Bytes("tfbw"))
						}
						return p2
					})() {
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
						_ = _index164
						(_930_v44).ArraySet1((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), (_index164).Int())
						r0 = ((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus((_dafny.Zero).Minus((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))))
						var _939_v52 _dafny.Array
						_ = _939_v52
						var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
						_ = _nw173
						_939_v52 = _nw173
						var _940_v53 _dafny.MultiSet
						_ = _940_v53
						_940_v53 = _dafny.MultiSetOf(_937_v50, _937_v50, _937_v50, _937_v50, _937_v50)
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
						_ = _index165
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
						_ = _index166
						var _rhs144 _dafny.Array = (func() _dafny.Array {
							if !(p2) {
								return _939_v52
							}
							return (func() _dafny.Array {
								if p3 {
									return _939_v52
								}
								return _939_v52
							})()
						})()
						_ = _rhs144
						var _rhs145 _dafny.MultiSet = _dafny.MultiSetFromSeq(_932_v46)
						_ = _rhs145
						var _rhs146 _dafny.Int = (_this).Fm1(p3, globalState)
						_ = _rhs146
						var _rhs147 _dafny.Int = _this.F2()
						_ = _rhs147
						var _lhs81 _dafny.Array = p1
						_ = _lhs81
						var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
						_ = _lhs82
						var _lhs83 _dafny.Array = p1
						_ = _lhs83
						var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))
						_ = _lhs84
						_939_v52 = _rhs144
						_940_v53 = _rhs145
						(_lhs81).ArraySet1(_rhs146, (_lhs82).Int())
						(_lhs83).ArraySet1(_rhs147, (_lhs84).Int())
						r0 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
						var _941_v54 _dafny.Map
						_ = _941_v54
						_941_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_936_v49).Merge(_936_v49), _933_v47)
						_941_v54 = (_941_v54).Update(_936_v49, _dafny.Companion_Sequence_.Concatenate(_933_v47, _dafny.SeqOf(p3)))
					} else {
						_930_v44 = _930_v44
						var _942_v55 _dafny.Array
						_ = _942_v55
						var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
						_ = _nw174
						_942_v55 = _nw174
						var _943_v56 _dafny.Map
						_ = _943_v56
						_943_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_this.F2(), (p0).Cardinality()), (func() _dafny.Array {
							if p2 {
								return _942_v55
							}
							return _942_v55
						})())
						var _944_v57 _dafny.Map
						_ = _944_v57
						_944_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _932_v46)
						var _945_v58 _dafny.Map
						_ = _945_v58
						_945_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(170))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg51 _dafny.Int) interface{} {
								return coer51(arg51)
							}
						}((func(_946_v50 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_947_i2 _dafny.Int) _dafny.CodePoint {
								return _946_v50
							}
						})(_937_v50))), p2, globalState), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_944_v57).Contains((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)) {
								return (_944_v57).Get((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)).(_dafny.Sequence)
							}
							return _932_v46
						})()).Cardinality()))
						_942_v55 = (func() _dafny.Array {
							if (_943_v56).Contains(((_945_v58).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(Companion_Default___.Fm28(p2, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), p3, _dafny.SetOf(p2, true), globalState), (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), globalState), _dafny.IntOfInt64(766)))).Cardinality()) {
								return (_943_v56).Get(((_945_v58).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(Companion_Default___.Fm28(p2, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), p3, _dafny.SetOf(p2, true), globalState), (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), globalState), _dafny.IntOfInt64(766)))).Cardinality()).(_dafny.Array)
							}
							return (func() _dafny.Array {
								if p3 {
									return _942_v55
								}
								return _942_v55
							})()
						})()
						var _948_v59 *C0
						_ = _948_v59
						var _nw175 *C0 = New_C0_()
						_ = _nw175
						_nw175.Ctor__(((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg52 _dafny.Int) interface{} {
								return coer52(arg52)
							}
						}((func(_949_v50 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_950_i3 _dafny.Int) _dafny.CodePoint {
								return _949_v50
							}
						})(_937_v50)))).Cardinality())) <= 0, (_this).Fm8(_dafny.Companion_Sequence_.Update(_932_v46, (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_932_v46).Cardinality()))).Uint32(), _937_v50), false, globalState))
						_948_v59 = _nw175
						(_948_v59).F4 = (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)
						var _951_v60 _dafny.Map
						_ = _951_v60
						_951_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _948_v59.F5)
						(_948_v59).F5 = (func() bool {
							if (_951_v60).Contains(!((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool))) {
								return (_951_v60).Get(!((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool))).(bool)
							}
							return p3
						})()
					}
					var _952_v61 *C1
					_ = _952_v61
					var _nw176 *C1 = New_C1_()
					_ = _nw176
					_nw176.Ctor__(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool) {
							return _932_v46
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("dvbp")
					})(), (Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool) {
							return _932_v46
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("dvbp")
					})()).Cardinality()))).Uint32(), _937_v50))
					_952_v61 = _nw176
					var _953_v62 *C0
					_ = _953_v62
					var _nw177 *C0 = New_C0_()
					_ = _nw177
					_nw177.Ctor__((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), (_dafny.IntOfInt64(-802)).Cmp(_dafny.IntOfInt64(940)) >= 0)
					_953_v62 = _nw177
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _954_v63 _dafny.Map
		_ = _954_v63
		_954_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).Plus(_this.F2()), _this.F2())
		var _955_v64 _dafny.MultiSet
		_ = _955_v64
		_955_v64 = _dafny.MultiSetOf((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), p2)
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))
		_ = _index167
		(p1).ArraySet1((_dafny.IntOfInt64(125)).Plus((_955_v64).Cardinality()), (_index167).Int())
		var _956_v65 D7
		_ = _956_v65
		_956_v65 = Companion_D7_.Create_DC15_((_dafny.SetOf((_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), true)).Cardinality(), _dafny.CodePoint('n'))
		var _957_v66 _dafny.CodePoint
		_ = _957_v66
		_957_v66 = _dafny.CodePoint('n')
		var _958_v67 D0
		_ = _958_v67
		_958_v67 = Companion_D0_.Create_DC1_(!(p2), (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), _957_v66, _this.F2())
		var _pat_let_tv18 = _930_v44
		_ = _pat_let_tv18
		var _pat_let_tv19 = _930_v44
		_ = _pat_let_tv19
		var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))
		_ = _index168
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _index169
		var _rhs148 _dafny.Sequence = Companion_Default___.Fm28((_933_v47).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_933_v47).Cardinality()))).Uint32()).(bool), true, p3, _dafny.SetOf(false, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)), globalState)
		_ = _rhs148
		var _rhs149 _dafny.Map = _954_v63
		_ = _rhs149
		var _rhs150 _dafny.Int = (_this.F2()).Plus(_this.F2())
		_ = _rhs150
		var _rhs151 _dafny.Int = (_956_v65).Dtor_cf27()
		_ = _rhs151
		var _rhs152 bool = func(_source18 D0) bool {
			if _source18.Is_DC1() {
				var _959___mcc_h0 bool = _source18.Get_().(D0_DC1).Cf1
				_ = _959___mcc_h0
				var _960___mcc_h1 bool = _source18.Get_().(D0_DC1).Cf2
				_ = _960___mcc_h1
				var _961___mcc_h2 _dafny.CodePoint = _source18.Get_().(D0_DC1).Cf3
				_ = _961___mcc_h2
				var _962___mcc_h3 _dafny.Int = _source18.Get_().(D0_DC1).Cf4
				_ = _962___mcc_h3
				var _963_cf4 _dafny.Int = _962___mcc_h3
				_ = _963_cf4
				var _964_cf3 _dafny.CodePoint = _961___mcc_h2
				_ = _964_cf3
				var _965_cf2 bool = _960___mcc_h1
				_ = _965_cf2
				var _966_cf1 bool = _959___mcc_h0
				_ = _966_cf1
				return (_pat_let_tv19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_pat_let_tv18), 0))).Int()).(bool)
			} else {
				var _967___mcc_h4 bool = _source18.Get_().(D0_DC0).Cf0
				_ = _967___mcc_h4
				var _968_cf0 bool = _967___mcc_h4
				_ = _968_cf0
				return (_dafny.IntOfInt64(172)).Cmp(_this.F2()) <= 0
			}
		}(_958_v67)
		_ = _rhs152
		var _lhs85 *C2 = _this
		_ = _lhs85
		var _lhs86 _dafny.Array = p1
		_ = _lhs86
		var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))
		_ = _lhs87
		var _lhs88 _dafny.Array = _930_v44
		_ = _lhs88
		var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _lhs89
		_932_v46 = _rhs148
		_954_v63 = _rhs149
		_lhs85.F2_set_(_rhs150)
		(_lhs86).ArraySet1(_rhs151, (_lhs87).Int())
		(_lhs88).ArraySet1(_rhs152, (_lhs89).Int())
		var _969_v68 _dafny.Set
		_ = _969_v68
		_969_v68 = _dafny.SetOf(p2, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), true)
		var _hi3 _dafny.Int = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
		_ = _hi3
		for _970_i4 := (func() _dafny.Int {
			if p2 {
				return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_969_v68).Cardinality(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))).Cardinality()
		})(); _970_i4.Cmp(_hi3) < 0; _970_i4 = _970_i4.Plus(_dafny.One) {
			var _pat_let_tv20 = _932_v46
			_ = _pat_let_tv20
			var _pat_let_tv21 = p2
			_ = _pat_let_tv21
			var _971_v69 _dafny.Set
			_ = _971_v69
			_971_v69 = _dafny.SetOf(_957_v66, (func(_pat_let16_0 D0) D0 {
				return func(_972_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let17_0 _dafny.Int) D0 {
						return func(_973_dt__update_hcf4_h0 _dafny.Int) D0 {
							return func(_pat_let18_0 bool) D0 {
								return func(_974_dt__update_hcf2_h0 bool) D0 {
									return Companion_D0_.Create_DC1_((_972_dt__update__tmp_h0).Dtor_cf1(), _974_dt__update_hcf2_h0, (_972_dt__update__tmp_h0).Dtor_cf3(), _973_dt__update_hcf4_h0)
								}(_pat_let18_0)
							}(_pat_let_tv21)
						}(_pat_let17_0)
					}(_dafny.IntOfUint32((_pat_let_tv20).Cardinality()))
				}(_pat_let16_0)
			}(_958_v67)).Dtor_cf3())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
			_ = _index170
			(_930_v44).ArraySet1(!((func() bool {
				if (p0).IsProperSubsetOf(Companion_Default___.Fm35(_957_v66, (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool), _dafny.MultiSetOf(p3), globalState)) {
					return (_971_v69).IsDisjointFrom(_dafny.SetOf(_dafny.CodePoint('g')))
				}
				return p3
			})()), (_index170).Int())
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))
			_ = _index171
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
			_ = _index172
			var _rhs153 _dafny.Int = (_this).Fm1(p3, globalState)
			_ = _rhs153
			var _rhs154 bool = (_930_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))).Int()).(bool)
			_ = _rhs154
			var _lhs90 _dafny.Array = p1
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))
			_ = _lhs91
			var _lhs92 _dafny.Array = _930_v44
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
			_ = _lhs93
			(_lhs90).ArraySet1(_rhs153, (_lhs91).Int())
			(_lhs92).ArraySet1(_rhs154, (_lhs93).Int())
			_955_v64 = _955_v64
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
			_ = _index173
			(_930_v44).ArraySet1(((p0).Cardinality()).Cmp(_this.F2()) != 0, (_index173).Int())
		}
		var _975_v70 _dafny.MultiSet
		_ = _975_v70
		_975_v70 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_954_v63).Contains(_dafny.IntOfInt64(215)) {
				return (_954_v63).Get(_dafny.IntOfInt64(215)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(405)
		})(), (_dafny.Zero).Minus((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)), ((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus(_this.F2())))
		var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _index174
		var _rhs155 bool = p2
		_ = _rhs155
		var _rhs156 _dafny.MultiSet = (p0).Intersection(_975_v70)
		_ = _rhs156
		var _lhs94 _dafny.Array = _930_v44
		_ = _lhs94
		var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_930_v44), 0))
		_ = _lhs95
		(_lhs94).ArraySet1(_rhs155, (_lhs95).Int())
		_975_v70 = _rhs156
		var _976_v71 _dafny.Set
		_ = _976_v71
		_976_v71 = _dafny.SetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
		var _977_v72 _dafny.Map
		_ = _977_v72
		_977_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), p2)
		var _978_v73 _dafny.Map
		_ = _978_v73
		_978_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_932_v46, _977_v72)
		var _979_v74 D5
		_ = _979_v74
		_979_v74 = Companion_D5_.Create_DC10_(p3, (_dafny.Zero).Minus((_dafny.Zero).Minus((_976_v71).Cardinality())), _978_v73, true)
		r0 = (_979_v74).Dtor_cf18()
		var _980_v75 _dafny.Sequence
		_ = _980_v75
		_980_v75 = _dafny.SeqOf(_976_v71)
		r1 = ((_980_v75).Select((Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_980_v75).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
		return r0, r1
	}
}
func (_this *C2) M6(p0 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _981_v0 _dafny.Array
		_ = _981_v0
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_15
		var _nw178 _dafny.Array
		_ = _nw178
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw178 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.Int = func(_982_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_982_i0, _this.F2())
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw178 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw178).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw178).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_981_v0 = _nw178
		var _983_v1 _dafny.Set
		_ = _983_v1
		_983_v1 = _dafny.SetOf(p0)
		var _984_v2 _dafny.Sequence
		_ = _984_v2
		_984_v2 = _dafny.SeqOf(_dafny.IntOfInt64(603))
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))
		_ = _index175
		(_981_v0).ArraySet1(((_983_v1).Union(func() _dafny.Set {
			var _coll31 = _dafny.NewBuilder()
			_ = _coll31
			for _iter35 := _dafny.Iterate((_984_v2).Elements()); ; {
				_compr_31, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _985_v3 _dafny.Int
				_985_v3 = interface{}(_compr_31).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_984_v2, _985_v3) {
					_coll31.Add(Companion_Default___.SafeDivisionInt(_985_v3, _dafny.IntOfInt64(916)))
				}
			}
			return _coll31.ToSet()
		}())).Cardinality(), (_index175).Int())
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))
		_ = _index176
		(_981_v0).ArraySet1((_dafny.IntOfUint32(((Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg53 _dafny.Int) interface{} {
				return coer53(arg53)
			}
		}(func(_986_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})))).Dtor_cf21()).Cardinality())).Minus(p0), (_index176).Int())
		var _987_v4 _dafny.Array
		_ = _987_v4
		var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
		_ = _nw179
		_987_v4 = _nw179
		var _988_v5 _dafny.Map
		_ = _988_v5
		_988_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_987_v4, _981_v0)
		_988_v5 = (_988_v5).Update(_987_v4, _981_v0)
		var _989_v6 bool
		_ = _989_v6
		_989_v6 = true
		var _rhs157 bool = _989_v6
		_ = _rhs157
		var _rhs158 bool = !(_989_v6) || (_989_v6)
		_ = _rhs158
		var _rhs159 bool = _989_v6
		_ = _rhs159
		r0 = _rhs157
		r0 = _rhs158
		r0 = _rhs159
		var _990_v7 _dafny.Map
		_ = _990_v7
		_990_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v6, _989_v6)
		var _991_v8 _dafny.Sequence
		_ = _991_v8
		_991_v8 = _dafny.SeqOf(_989_v6)
		var _992_v9 _dafny.Array
		_ = _992_v9
		var _nwElement0_40 bool = _989_v6
		_ = _nwElement0_40
		var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(5))
		_ = _nw180
		(_nw180).ArraySet1(_nwElement0_40, 0)
		(_nw180).ArraySet1(_989_v6, 1)
		(_nw180).ArraySet1((_991_v8).Select((Companion_Default___.SafeIndex((_981_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_991_v8).Cardinality()))).Uint32()).(bool), 2)
		(_nw180).ArraySet1(_989_v6, 3)
		(_nw180).ArraySet1(_989_v6, 4)
		_992_v9 = _nw180
		var _993_v10 _dafny.Map
		_ = _993_v10
		_993_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_990_v7).Cardinality(), (_this).Fm1(_989_v6, globalState)), _992_v9)
		var _994_v11 _dafny.Map
		_ = _994_v11
		_994_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v6, p0)
		var _995_v12 _dafny.Sequence
		_ = _995_v12
		_995_v12 = _dafny.SeqOf(_994_v11)
		_993_v10 = (_993_v10).Update(_dafny.IntOfUint32((_995_v12).Cardinality()), _992_v9)
		r0 = _989_v6
		var _996_i2 _dafny.Int
		_ = _996_i2
		_996_i2 = _dafny.Zero
		{
			for _989_v6 {
				{
					if (_996_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_996_i2 = (_996_i2).Plus(_dafny.One)
					var _997_v13 _dafny.CodePoint
					_ = _997_v13
					_997_v13 = _dafny.CodePoint('y')
					var _998_v14 _dafny.Sequence
					_ = _998_v14
					_998_v14 = _dafny.UnicodeSeqOfUtf8Bytes("mxs")
					var _rhs160 bool = (func() bool {
						if (_990_v7).Contains(_989_v6) {
							return (_990_v7).Get(_989_v6).(bool)
						}
						return _989_v6
					})()
					_ = _rhs160
					var _rhs161 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_984_v2, _984_v2), _dafny.Companion_Sequence_.Concatenate(_984_v2, Companion_Default___.Fm36(_989_v6, _997_v13, (_dafny.SetOf(_989_v6, (_this).Fm8(_998_v14, _989_v6, globalState))).Cardinality(), globalState))))
					_ = _rhs161
					_989_v6 = _rhs160
					_989_v6 = _rhs161
					if !(_989_v6) {
						_997_v13 = _997_v13
						_989_v6 = _989_v6
						_989_v6 = _989_v6
						var _999_v15 _dafny.Set
						_ = _999_v15
						_999_v15 = _dafny.SetOf(!(_989_v6), _989_v6, (p0).Cmp(_this.F2()) == 0)
						_999_v15 = (_999_v15).Union(_999_v15)
						_989_v6 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(675)), _dafny.Companion_Sequence_.Update(_984_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_984_v2).Cardinality()))).Uint32(), _dafny.IntOfUint32((_998_v14).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_981_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))).Int()).(_dafny.Int)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-907))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}((func(_1000_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
							return func(_1001_i3 _dafny.Int) _dafny.Int {
								return (_dafny.MultiSetFromSeq(_1000_v8)).Cardinality()
							}
						})(_991_v8)))))
					} else {
						var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_992_v9), 0))
						_ = _index177
						(_992_v9).ArraySet1(_989_v6, (_index177).Int())
						var _1002_v16 _dafny.Set
						_ = _1002_v16
						_1002_v16 = _dafny.SetOf(_997_v13, Companion_Default___.Fm37((_this).Fm1(_989_v6, globalState), globalState))
						var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_992_v9), 0))
						_ = _index178
						(_992_v9).ArraySet1(!((_1002_v16).Equals(_1002_v16)) || ((_dafny.MultiSetFromSeq(_991_v8)).IsDisjointFrom(_dafny.MultiSetOf(true))), (_index178).Int())
						(_this).F2_set_((_981_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))).Int()).(_dafny.Int))
						var _1003_v17 D0
						_ = _1003_v17
						_1003_v17 = Companion_D0_.Create_DC0_(!((_992_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_992_v9), 0))).Int()).(bool)))
						_989_v6 = !_dafny.Companion_Sequence_.Contains(_984_v2, (_this).Fm1((_1003_v17).Dtor_cf0(), globalState))
						var _1004_v18 _dafny.Array
						_ = _1004_v18
						var _out22 _dafny.Array
						_ = _out22
						_out22 = Companion_Default___.M0(_989_v6, (_992_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_992_v9), 0))).Int()).(bool), globalState)
						_1004_v18 = _out22
						var _1005_v19 *C1
						_ = _1005_v19
						var _nw181 *C1 = New_C1_()
						_ = _nw181
						_nw181.Ctor__(_998_v14)
						_1005_v19 = _nw181
						var _1006_v20 D8
						_ = _1006_v20
						_1006_v20 = Companion_D8_.Create_DC17_(_989_v6, _1005_v19, p0)
						var _1007_v21 _dafny.Array
						_ = _1007_v21
						var _nwElement0_41 D8 = _1006_v20
						_ = _nwElement0_41
						var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.One)
						_ = _nw182
						(_nw182).ArraySet1(_nwElement0_41, 0)
						_1007_v21 = _nw182
						var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_1007_v21), 0))
						_ = _index179
						(_1007_v21).ArraySet1(_1006_v20, (_index179).Int())
						var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_1007_v21), 0))
						_ = _index180
						(_1007_v21).ArraySet1(_1006_v20, (_index180).Int())
					}
					var _1008_v22 *C0
					_ = _1008_v22
					var _nw183 *C0 = New_C0_()
					_ = _nw183
					_nw183.Ctor__(_989_v6, _989_v6)
					_1008_v22 = _nw183
					var _1009_v23 *C1
					_ = _1009_v23
					var _nw184 *C1 = New_C1_()
					_ = _nw184
					_nw184.Ctor__(_998_v14)
					_1009_v23 = _nw184
					var _1010_v24 D8
					_ = _1010_v24
					_1010_v24 = Companion_D8_.Create_DC16_(_1009_v23)
					var _source19 D8 = (func() D8 {
						if _989_v6 {
							return _1010_v24
						}
						return _1010_v24
					})()
					_ = _source19
					if _source19.Is_DC17() {
						var _1011___mcc_h0 bool = _source19.Get_().(D8_DC17).Cf30
						_ = _1011___mcc_h0
						var _1012___mcc_h1 *C1 = _source19.Get_().(D8_DC17).Cf31
						_ = _1012___mcc_h1
						var _1013___mcc_h2 _dafny.Int = _source19.Get_().(D8_DC17).Cf32
						_ = _1013___mcc_h2
						var _1014_cf32 _dafny.Int = _1013___mcc_h2
						_ = _1014_cf32
						var _1015_cf31 *C1 = _1012___mcc_h1
						_ = _1015_cf31
						var _1016_cf30 bool = _1011___mcc_h0
						_ = _1016_cf30
						var _1017_v25 _dafny.Array
						_ = _1017_v25
						var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
						_ = _nw185
						_1017_v25 = _nw185
						_1017_v25 = _1017_v25
						var _1018_v26 _dafny.Set
						_ = _1018_v26
						_1018_v26 = _dafny.SetOf(_1016_cf30, !(_1008_v22.F5), _1008_v22.F5)
						var _1019_v27 D4
						_ = _1019_v27
						_1019_v27 = Companion_D4_.Create_DC5_(_1018_v26)
						var _1020_v28 D4
						_ = _1020_v28
						_1020_v28 = Companion_D4_.Create_DC7_(_1019_v27)
						_1020_v28 = Companion_D4_.Create_DC7_(Companion_D4_.Create_DC5_(_dafny.SetOf(true, false, _1016_cf30, _1008_v22.F4, _1008_v22.F4)))
						_998_v14 = _dafny.UnicodeSeqOfUtf8Bytes("cfo")
						var _1021_v29 _dafny.MultiSet
						_ = _1021_v29
						_1021_v29 = _dafny.MultiSetOf(_1008_v22.F5)
						_1021_v29 = ((_1021_v29).Difference(_1021_v29)).Intersection(_1021_v29)
					} else {
						var _1022___mcc_h3 *C1 = _source19.Get_().(D8_DC16).Cf29
						_ = _1022___mcc_h3
						var _1023_cf29 *C1 = _1022___mcc_h3
						_ = _1023_cf29
						var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))
						_ = _index181
						(_981_v0).ArraySet1(((_981_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((_this.F2()).Times((_this).Fm1(_1008_v22.F5, globalState)))), (_index181).Int())
						var _1024_v31 _dafny.Map
						_ = _1024_v31
						_1024_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1009_v23).F3(), (_981_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_981_v0), 0))).Int()).(_dafny.Int))
						var _1025_v32 _dafny.Map
						_ = _1025_v32
						_1025_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1008_v22.F4)
						var _1026_v33 _dafny.Map
						_ = _1026_v33
						_1026_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1009_v23).F3(), _1025_v32)
						var _1027_v34 _dafny.Sequence
						_ = _1027_v34
						_1027_v34 = _dafny.SeqOf(func() _dafny.Map {
							var _coll32 = _dafny.NewMapBuilder()
							_ = _coll32
							for _iter36 := _dafny.Iterate((_1024_v31).Keys().Elements()); ; {
								_compr_32, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _1028_v30 _dafny.Sequence
								_1028_v30 = interface{}(_compr_32).(_dafny.Sequence)
								if (_1024_v31).Contains(_1028_v30) {
									_coll32.Add(_1028_v30, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(_1008_v22.F5)))
								}
							}
							return _coll32.ToMap()
						}(), _1026_v33)
						var _1029_v35 D5
						_ = _1029_v35
						_1029_v35 = Companion_D5_.Create_DC10_((_989_v6) && (false), (_dafny.Zero).Minus(_this.F2()), (_1027_v34).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1027_v34).Cardinality()))).Uint32()).(_dafny.Map), _1008_v22.F5)
						var _rhs162 _dafny.Int = _this.F2()
						_ = _rhs162
						var _rhs163 D5 = _1029_v35
						_ = _rhs163
						r1 = _rhs162
						_1029_v35 = _rhs163
						var _1030_v36 *C1
						_ = _1030_v36
						var _nw186 *C1 = New_C1_()
						_ = _nw186
						_nw186.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("r"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_991_v8).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()))).Uint32(), _997_v13), _dafny.UnicodeSeqOfUtf8Bytes("log")))
						_1030_v36 = _nw186
						var _1031_v37 _dafny.Map
						_ = _1031_v37
						_1031_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1008_v22, _1008_v22.F4)
						_1031_v37 = (_1031_v37).Update(_1008_v22, _1008_v22.F5)
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		r0 = !(_989_v6)
		var _1032_v38 _dafny.Sequence
		_ = _1032_v38
		_1032_v38 = _dafny.UnicodeSeqOfUtf8Bytes("nbprgidjf")
		r1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _989_v6 {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}(func(_1033_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), _1032_v38)
			}
			return _1032_v38
		})()).Cardinality())
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f2 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f2 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C3{}
var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F2() _dafny.Int {
	return _this._f2
}
func (_this *C3) F2_set_(value _dafny.Int) {
	_this._f2 = value
}
func (_this *C3) Ctor__(f2 _dafny.Int) {
	{
		(_this)._f2 = f2
	}
}
func (_this *C3) Fm8(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	{
		return !((_dafny.SetOf(!(false), false)).Union(_dafny.SetOf(!(true)))).Contains(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("haphhjf"), _dafny.UnicodeSeqOfUtf8Bytes("jktmaskl")))
	}
}
func (_this *C3) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
			if !(false) {
				return _dafny.SeqOf(_this.F2())
			}
			return _dafny.SeqOf(_this.F2())
		})(), _this.F2())
	}
}
func (_this *C3) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if !(false) {
				return _this.F2()
			}
			return _this.F2()
		})()))).Plus(Companion_Default___.SafeModuloInt(_this.F2(), _this.F2()))
	}
}
func (_this *C3) Fm10(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if false {
				return _dafny.UnicodeSeqOfUtf8Bytes("tues")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("kchnvrhfl")
		})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}(func(_1034_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})))
	}
}
func (_this *C3) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1035_v0 _dafny.Map
		_ = _1035_v0
		_1035_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), p3)
		_1035_v0 = _1035_v0
		var _1036_v1 _dafny.Sequence
		_ = _1036_v1
		_1036_v1 = _dafny.SeqOf(_this.F2())
		(_this).F2_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1036_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(826))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_1037_i0 _dafny.Int) _dafny.Int {
			return _this.F2()
		})))).Cardinality()))
		var _1038_v2 _dafny.CodePoint
		_ = _1038_v2
		_1038_v2 = _dafny.CodePoint('a')
		var _1039_v3 D0
		_ = _1039_v3
		_1039_v3 = Companion_D0_.Create_DC1_(true, !(p3), _1038_v2, _dafny.IntOfInt64(87))
		var _1040_i1 _dafny.Int
		_ = _1040_i1
		_1040_i1 = _dafny.Zero
		{
			for (((_this).Fm1(p2, globalState)).Cmp(_this.F2()) == 0) || ((func() bool {
				if p2 {
					return p2
				}
				return (_1039_v3).Dtor_cf1()
			})()) {
				{
					if (_1040_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1040_i1 = (_1040_i1).Plus(_dafny.One)
					var _1041_v4 _dafny.Array
					_ = _1041_v4
					var _len0_16 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_16
					var _nw187 _dafny.Array
					_ = _nw187
					if _len0_16.Cmp(_dafny.Zero) == 0 {
						_nw187 = _dafny.NewArray(_len0_16)
					} else {
						var _init16 func(_dafny.Int) _dafny.Sequence = (func(_1042_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
							return func(_1043_i2 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg58 _dafny.Int) interface{} {
										return coer58(arg58)
									}
								}((func(_1044_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_1045_i3 _dafny.Int) _dafny.CodePoint {
										return _1044_v2
									}
								})(_1042_v2))), _dafny.UnicodeSeqOfUtf8Bytes("lap"))
							}
						})(_1038_v2)
						_ = _init16
						var _element0_16 = _init16(_dafny.Zero)
						_ = _element0_16
						_nw187 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
						(_nw187).ArraySet1(_element0_16, 0)
						var _nativeLen0_16 = (_len0_16).Int()
						_ = _nativeLen0_16
						for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
							(_nw187).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
						}
					}
					_1041_v4 = _nw187
					var _1046_v5 _dafny.Sequence
					_ = _1046_v5
					_1046_v5 = _dafny.UnicodeSeqOfUtf8Bytes("amusabog")
					var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1041_v4), 0))
					_ = _index182
					(_1041_v4).ArraySet1(_1046_v5, (_index182).Int())
					var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1041_v4), 0))
					_ = _index183
					(_1041_v4).ArraySet1(_1046_v5, (_index183).Int())
					var _1047_v6 _dafny.Map
					_ = _1047_v6
					_1047_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus((_this).Fm1(p2, globalState)))
					var _1048_v7 _dafny.Sequence
					_ = _1048_v7
					_1048_v7 = _dafny.SeqOf(_1047_v6)
					var _1049_v8 _dafny.Set
					_ = _1049_v8
					_1049_v8 = _dafny.SetOf(p2, p3)
					var _1050_v9 _dafny.Map
					_ = _1050_v9
					_1050_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1049_v8).Cardinality(), _1047_v6)
					var _1051_v10 _dafny.Array
					_ = _1051_v10
					var _nwElement0_42 _dafny.Map = (_1048_v7).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1048_v7).Cardinality()))).Uint32()).(_dafny.Map)
					_ = _nwElement0_42
					var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(19))
					_ = _nw188
					(_nw188).ArraySet1(_nwElement0_42, 0)
					(_nw188).ArraySet1(_1047_v6, 1)
					(_nw188).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F2()), 2)
					(_nw188).ArraySet1(_1047_v6, 3)
					(_nw188).ArraySet1(_1047_v6, 4)
					(_nw188).ArraySet1(_1047_v6, 5)
					(_nw188).ArraySet1((_1047_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F2())), 6)
					(_nw188).ArraySet1((_1047_v6).Update(p1, _this.F2()), 7)
					(_nw188).ArraySet1((_1047_v6).Merge(_1047_v6), 8)
					(_nw188).ArraySet1((_1047_v6).Merge(_1047_v6), 9)
					(_nw188).ArraySet1((_1047_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F2())), 10)
					(_nw188).ArraySet1(_1047_v6, 11)
					(_nw188).ArraySet1(((_1047_v6).Update(p1, _dafny.IntOfUint32(((_1041_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1041_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_1046_v5).Cardinality()))), 12)
					(_nw188).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F2()), 13)
					(_nw188).ArraySet1((_1047_v6).Merge(_1047_v6), 14)
					(_nw188).ArraySet1(_1047_v6, 15)
					(_nw188).ArraySet1(_1047_v6, 16)
					(_nw188).ArraySet1(_1047_v6, 17)
					(_nw188).ArraySet1((func() _dafny.Map {
						if (_1050_v9).Contains(_this.F2()) {
							return (_1050_v9).Get(_this.F2()).(_dafny.Map)
						}
						return _1047_v6
					})(), 18)
					_1051_v10 = _nw188
					var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1051_v10), 0))
					_ = _index184
					(_1051_v10).ArraySet1(_1047_v6, (_index184).Int())
					var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1051_v10), 0))
					_ = _index185
					(_1051_v10).ArraySet1(_1047_v6, (_index185).Int())
					r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1046_v5, _dafny.UnicodeSeqOfUtf8Bytes("l"))).Cardinality())
					var _1052_v11 _dafny.Map
					_ = _1052_v11
					_1052_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _this.F2())
					var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((p1), 0))
					_ = _index186
					(p1).ArraySet1((_1036_v1).Select((Companion_Default___.SafeIndex((_1052_v11).Cardinality(), _dafny.IntOfUint32((_1036_v1).Cardinality()))).Uint32()).(_dafny.Int), (_index186).Int())
					var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((p1), 0))
					_ = _index187
					(p1).ArraySet1(_this.F2(), (_index187).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1053_v12 _dafny.Set
		_ = _1053_v12
		_1053_v12 = _dafny.SetOf(_this.F2())
		(_this).F2_set_((_1053_v12).Cardinality())
		var _1054_v13 bool
		_ = _1054_v13
		_1054_v13 = false
		var _1055_v14 _dafny.Sequence
		_ = _1055_v14
		_1055_v14 = _dafny.UnicodeSeqOfUtf8Bytes("unuhbnudi")
		var _1056_v15 _dafny.Array
		_ = _1056_v15
		var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw189
		_1056_v15 = _nw189
		var _1057_v16 _dafny.Set
		_ = _1057_v16
		_1057_v16 = _dafny.SetOf(_dafny.SetOf(p3))
		var _rhs164 bool = _1054_v13
		_ = _rhs164
		var _rhs165 bool = (_1057_v16).IsSubsetOf(_1057_v16)
		_ = _rhs165
		var _rhs166 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("mlb")
		_ = _rhs166
		var _rhs167 _dafny.Array = _1056_v15
		_ = _rhs167
		_1054_v13 = _rhs164
		_1054_v13 = _rhs165
		_1055_v14 = _rhs166
		_1056_v15 = _rhs167
		var _1058_v17 _dafny.Array
		_ = _1058_v17
		var _nw190 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw190
		_1058_v17 = _nw190
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1058_v17), 0))
		_ = _index188
		(_1058_v17).ArraySet1(_1054_v13, (_index188).Int())
		var _1059_v18 D0
		_ = _1059_v18
		_1059_v18 = Companion_D0_.Create_DC0_(p2)
		var _pat_let_tv22 = p2
		_ = _pat_let_tv22
		var _pat_let_tv23 = _1055_v14
		_ = _pat_let_tv23
		var _pat_let_tv24 = _1059_v18
		_ = _pat_let_tv24
		var _pat_let_tv25 = globalState
		_ = _pat_let_tv25
		var _1060_v19 _dafny.MultiSet
		_ = _1060_v19
		_1060_v19 = _dafny.MultiSetOf((func(_pat_let19_0 D0) D0 {
			return func(_1063_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let22_0 bool) D0 {
					return func(_1064_dt__update_hcf0_h1 bool) D0 {
						return Companion_D0_.Create_DC0_(_1064_dt__update_hcf0_h1)
					}(_pat_let22_0)
				}((_this).Fm8(_pat_let_tv23, (_pat_let_tv24).Dtor_cf0(), _pat_let_tv25))
			}(_pat_let19_0)
		}(func(_pat_let20_0 D0) D0 {
			return func(_1061_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let21_0 bool) D0 {
					return func(_1062_dt__update_hcf0_h0 bool) D0 {
						return Companion_D0_.Create_DC0_(_1062_dt__update_hcf0_h0)
					}(_pat_let21_0)
				}(_pat_let_tv22)
			}(_pat_let20_0)
		}(_1059_v18))).Dtor_cf0())
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1058_v17), 0))
		_ = _index189
		(_1058_v17).ArraySet1(!(_1060_v19).Contains(false), (_index189).Int())
		r0 = (_dafny.Zero).Minus((_this.F2()).Times(_this.F2()))
		var _1065_v20 _dafny.Sequence
		_ = _1065_v20
		_1065_v20 = _dafny.SeqOf(true)
		r1 = (_this).Fm1((_1065_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1055_v14, (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1055_v14).Cardinality()))).Uint32(), _1038_v2)).Cardinality()), _dafny.IntOfUint32((_1065_v20).Cardinality()))).Uint32()).(bool), globalState)
		return r0, r1
	}
}
func (_this *C3) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		if p2 {
			if p2 {
				var _1066_v0 _dafny.MultiSet
				_ = _1066_v0
				_1066_v0 = _dafny.MultiSetOf(_dafny.IntOfInt64(-215), (_this.F2()).Times(_dafny.IntOfInt64(446)), (p0).Minus(p0), p0)
				var _1067_v1 _dafny.Sequence
				_ = _1067_v1
				_1067_v1 = _dafny.SeqOf(_1066_v0, _dafny.MultiSetOf(p0, p0))
				_1066_v0 = (_1066_v0).Intersection((_1067_v1).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1067_v1).Cardinality()))).Uint32()).(_dafny.MultiSet))
				(_this).F2_set_(_this.F2())
				var _1068_v2 _dafny.Set
				_ = _1068_v2
				_1068_v2 = _dafny.SetOf(p0)
				var _1069_v3 _dafny.Map
				_ = _1069_v3
				_1069_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1068_v2, p0)
				_1069_v3 = (_1069_v3).Update(_1068_v2, _this.F2())
				var _1070_v4 _dafny.Map
				_ = _1070_v4
				_1070_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p2)
				var _1071_v5 _dafny.Array
				_ = _1071_v5
				var _nwElement0_43 bool = ((func() bool {
					if (_1070_v4).Contains(p3) {
						return (_1070_v4).Get(p3).(bool)
					}
					return p3
				})()) || (p2)
				_ = _nwElement0_43
				var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(6))
				_ = _nw191
				(_nw191).ArraySet1(_nwElement0_43, 0)
				(_nw191).ArraySet1((p0).Cmp(p0) == 0, 1)
				(_nw191).ArraySet1(true, 2)
				(_nw191).ArraySet1((p0).Cmp(p0) == 0, 3)
				(_nw191).ArraySet1(p3, 4)
				(_nw191).ArraySet1(!(p2), 5)
				_1071_v5 = _nw191
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1071_v5), 0))
				_ = _index190
				(_1071_v5).ArraySet1(p2, (_index190).Int())
				var _1072_v6 _dafny.Sequence
				_ = _1072_v6
				_1072_v6 = _dafny.UnicodeSeqOfUtf8Bytes("ecarnc")
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1071_v5), 0))
				_ = _index191
				(_1071_v5).ArraySet1((_this).Fm8(_1072_v6, p3, globalState), (_index191).Int())
				var _1073_v7 _dafny.Map
				_ = _1073_v7
				_1073_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1072_v6, _this.F2())
				_1073_v7 = (_1073_v7).Update(_1072_v6, Companion_Default___.SafeDivisionInt(p0, p0))
			} else {
				r0 = p3
				var _1074_v8 _dafny.Array
				_ = _1074_v8
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw192
				_1074_v8 = _nw192
				var _1075_v9 _dafny.Array
				_ = _1075_v9
				var _nwElement0_44 _dafny.Array = _1074_v8
				_ = _nwElement0_44
				var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(5))
				_ = _nw193
				(_nw193).ArraySet1(_nwElement0_44, 0)
				(_nw193).ArraySet1(_1074_v8, 1)
				(_nw193).ArraySet1(_1074_v8, 2)
				(_nw193).ArraySet1(_1074_v8, 3)
				(_nw193).ArraySet1(_1074_v8, 4)
				_1075_v9 = _nw193
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_1075_v9), 0))
				_ = _index192
				(_1075_v9).ArraySet1(_1074_v8, (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_1075_v9), 0))
				_ = _index193
				(_1075_v9).ArraySet1(_1074_v8, (_index193).Int())
				r2 = ((func() _dafny.Int {
					if p3 {
						return _this.F2()
					}
					return p0
				})()).Times(Companion_Default___.SafeModuloInt(p0, _this.F2()))
				var _1076_v10 _dafny.Sequence
				_ = _1076_v10
				_1076_v10 = _dafny.UnicodeSeqOfUtf8Bytes("k")
				r0 = (_this).Fm8(_1076_v10, false, globalState)
				var _1077_v11 _dafny.CodePoint
				_ = _1077_v11
				_1077_v11 = _dafny.CodePoint('n')
				_1077_v11 = _1077_v11
			}
			var _1078_v12 _dafny.Array
			_ = _1078_v12
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_17
			var _nw194 _dafny.Array
			_ = _nw194
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw194 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Int = func(_1079_i0 _dafny.Int) _dafny.Int {
					return (_1079_i0).Times(_this.F2())
				}
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw194 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw194).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw194).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_1078_v12 = _nw194
			_1078_v12 = _1078_v12
			if p2 {
				var _1080_v13 _dafny.Array
				_ = _1080_v13
				var _nw195 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw195
				_1080_v13 = _nw195
				var _1081_v14 _dafny.MultiSet
				_ = _1081_v14
				_1081_v14 = _dafny.MultiSetOf(p0)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))
				_ = _index194
				(_1080_v13).ArraySet1((_1081_v14).IsDisjointFrom(_1081_v14), (_index194).Int())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))
				_ = _index195
				(_1080_v13).ArraySet1(p3, (_index195).Int())
				var _1082_v15 _dafny.MultiSet
				_ = _1082_v15
				_1082_v15 = _dafny.MultiSetOf(p3)
				var _1083_v16 _dafny.Array
				_ = _1083_v16
				var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
				_ = _nw196
				_1083_v16 = _nw196
				var _1084_v17 _dafny.Map
				_ = _1084_v17
				_1084_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1082_v15).Union(_1082_v15), _1083_v16)
				_1084_v17 = _1084_v17
				var _1085_v18 _dafny.Sequence
				_ = _1085_v18
				_1085_v18 = _dafny.UnicodeSeqOfUtf8Bytes("tgfihkea")
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))
				_ = _index196
				(_1080_v13).ArraySet1((_this).Fm8(_dafny.Companion_Sequence_.Concatenate(_1085_v18, _1085_v18), (_1080_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))).Int()).(bool), globalState), (_index196).Int())
				var _1086_v19 *C0
				_ = _1086_v19
				var _nw197 *C0 = New_C0_()
				_ = _nw197
				_nw197.Ctor__(p2, !((_1080_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))).Int()).(bool)))
				_1086_v19 = _nw197
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1080_v13), 0))
				_ = _index197
				(_1080_v13).ArraySet1(!(!(!(p3) || (p2))), (_index197).Int())
			} else {
				var _1087_v20 _dafny.Sequence
				_ = _1087_v20
				_1087_v20 = _dafny.SeqOf(p3)
				var _1088_v21 D0
				_ = _1088_v21
				_1088_v21 = Companion_D0_.Create_DC1_(p3, p2, _dafny.CodePoint('u'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tplkvlen")).Cardinality()))
				var _1089_v22 _dafny.CodePoint
				_ = _1089_v22
				_1089_v22 = _dafny.CodePoint('h')
				var _pat_let_tv26 = _1089_v22
				_ = _pat_let_tv26
				var _pat_let_tv27 = p2
				_ = _pat_let_tv27
				var _1090_v23 _dafny.Set
				_ = _1090_v23
				_1090_v23 = _dafny.SetOf(func(_pat_let23_0 D0) D0 {
					return func(_1091_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let24_0 _dafny.CodePoint) D0 {
							return func(_1092_dt__update_hcf3_h0 _dafny.CodePoint) D0 {
								return func(_pat_let25_0 bool) D0 {
									return func(_1093_dt__update_hcf2_h0 bool) D0 {
										return Companion_D0_.Create_DC1_((_1091_dt__update__tmp_h0).Dtor_cf1(), _1093_dt__update_hcf2_h0, _1092_dt__update_hcf3_h0, (_1091_dt__update__tmp_h0).Dtor_cf4())
									}(_pat_let25_0)
								}(_pat_let_tv27)
							}(_pat_let24_0)
						}(_pat_let_tv26)
					}(_pat_let23_0)
				}(_1088_v21), _1088_v21, _1088_v21)
				var _1094_v24 _dafny.Array
				_ = _1094_v24
				var _nwElement0_45 bool = !(p2)
				_ = _nwElement0_45
				var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(14))
				_ = _nw198
				(_nw198).ArraySet1(_nwElement0_45, 0)
				(_nw198).ArraySet1((_dafny.IntOfInt64(175)).Cmp(p0) > 0, 1)
				(_nw198).ArraySet1((_this.F2()).Cmp(_this.F2()) == 0, 2)
				(_nw198).ArraySet1(p2, 3)
				(_nw198).ArraySet1(p3, 4)
				(_nw198).ArraySet1(p3, 5)
				(_nw198).ArraySet1(p3, 6)
				(_nw198).ArraySet1((_1087_v20).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1087_v20).Cardinality()))).Uint32()).(bool), 7)
				(_nw198).ArraySet1(p2, 8)
				(_nw198).ArraySet1(p3, 9)
				(_nw198).ArraySet1(p3, 10)
				(_nw198).ArraySet1((_dafny.SetOf(_1088_v21, _1088_v21)).IsSubsetOf(_1090_v23), 11)
				(_nw198).ArraySet1(p2, 12)
				(_nw198).ArraySet1(p2, 13)
				_1094_v24 = _nw198
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1094_v24), 0))
				_ = _index198
				(_1094_v24).ArraySet1(!(!((_this).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("ihkkeqm"), p2, globalState))), (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1094_v24), 0))
				_ = _index199
				(_1094_v24).ArraySet1(p3, (_index199).Int())
				var _1095_v25 _dafny.Sequence
				_ = _1095_v25
				_1095_v25 = _dafny.UnicodeSeqOfUtf8Bytes("cwutipmey")
				_1087_v20 = _dafny.SeqOf((_dafny.IntOfUint32((_1095_v25).Cardinality())).Cmp(p0) >= 0, p3, !(p3))
				var _1096_v26 *C0
				_ = _1096_v26
				var _nw199 *C0 = New_C0_()
				_ = _nw199
				_nw199.Ctor__(p3, p2)
				_1096_v26 = _nw199
				var _1097_v27 _dafny.Map
				_ = _1097_v27
				_1097_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _this.F2())
				var _1098_v28 _dafny.Set
				_ = _1098_v28
				_1098_v28 = _dafny.SetOf(p0)
				var _1099_v29 _dafny.Map
				_ = _1099_v29
				_1099_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1098_v28)
				var _1100_v30 _dafny.Map
				_ = _1100_v30
				_1100_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.Update(_1095_v25, (Companion_Default___.SafeIndex((_1099_v29).Cardinality(), _dafny.IntOfUint32((_1095_v25).Cardinality()))).Uint32(), _1089_v22))
				_1097_v27 = (_1097_v27).Update((p0).Times(_this.F2()), (_1100_v30).Cardinality())
				r2 = (func() _dafny.Int {
					if _1096_v26.F4 {
						return (p0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}((func(_1101_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1102_i1 _dafny.Int) _dafny.Int {
								return _1101_p0
							}
						})(p0)))).Cardinality()))
					}
					return _this.F2()
				})()
			}
			var _1103_v31 _dafny.Map
			_ = _1103_v31
			_1103_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
			r0 = (func() bool {
				if (_1103_v31).Contains(p3) {
					return (_1103_v31).Get(p3).(bool)
				}
				return p2
			})()
			var _1104_v32 _dafny.Array
			_ = _1104_v32
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_18
			var _nw200 _dafny.Array
			_ = _nw200
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw200 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.Map = (func(_1105_v31 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1106_i2 _dafny.Int) _dafny.Map {
						return _1105_v31
					}
				})(_1103_v31)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw200 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw200).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw200).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_1104_v32 = _nw200
			var _1107_v33 _dafny.Map
			_ = _1107_v33
			_1107_v33 = _1103_v31
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1104_v32), 0))
			_ = _index200
			(_1104_v32).ArraySet1(_1107_v33, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1104_v32), 0))
			_ = _index201
			(_1104_v32).ArraySet1(_1107_v33, (_index201).Int())
		} else {
			r2 = _this.F2()
			var _1108_v34 _dafny.Sequence
			_ = _1108_v34
			_1108_v34 = _dafny.SeqOf(p3, p3, p2, p2)
			var _1109_v35 _dafny.Sequence
			_ = _1109_v35
			_1109_v35 = _dafny.UnicodeSeqOfUtf8Bytes("pdbarimii")
			var _1110_v36 D6
			_ = _1110_v36
			_1110_v36 = Companion_D6_.Create_DC11_(_1109_v35)
			r2 = (func() _dafny.Int {
				if !((_1108_v34).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1108_v34).Cardinality()))).Uint32()).(bool)) {
					return (_dafny.IntOfInt64(121)).Minus(_dafny.IntOfUint32(((_1110_v36).Dtor_cf21()).Cardinality()))
				}
				return _this.F2()
			})()
			var _1111_v37 _dafny.Set
			_ = _1111_v37
			_1111_v37 = _dafny.SetOf(_this.F2())
			var _1112_v38 _dafny.Map
			_ = _1112_v38
			_1112_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1111_v37).Difference(_dafny.SetOf(_dafny.IntOfInt64(511), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), true)).Cardinality(), (_dafny.MultiSetOf(p0)).Cardinality())), p2)
			_1112_v38 = _1112_v38
			var _1113_v39 *C0
			_ = _1113_v39
			var _nw201 *C0 = New_C0_()
			_ = _nw201
			_nw201.Ctor__(false, p2)
			_1113_v39 = _nw201
			r0 = p3
		}
		var _1114_v40 _dafny.Sequence
		_ = _1114_v40
		_1114_v40 = _dafny.SeqOf(p3)
		if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1114_v40, _1114_v40)).Cardinality())).Cmp(p0) == 0 {
			var _1115_v42 _dafny.Map
			_ = _1115_v42
			_1115_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _1116_v43 _dafny.Sequence
			_ = _1116_v43
			_1116_v43 = _dafny.SeqOf(_1115_v42, _1115_v42)
			var _1117_v44 _dafny.Sequence
			_ = _1117_v44
			_1117_v44 = _dafny.SeqOf(_1116_v43, _1116_v43, _1116_v43)
			var _1118_v45 _dafny.Map
			_ = _1118_v45
			_1118_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2), p2)
			var _1119_v46 _dafny.Set
			_ = _1119_v46
			_1119_v46 = _dafny.SetOf(_this.F2(), _this.F2())
			var _1120_v47 _dafny.Map
			_ = _1120_v47
			_1120_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1119_v46, _1118_v45)
			var _1121_v48 _dafny.MultiSet
			_ = _1121_v48
			_1121_v48 = _dafny.MultiSetOf(false, p2)
			var _1122_v51 _dafny.Sequence
			_ = _1122_v51
			_1122_v51 = _dafny.UnicodeSeqOfUtf8Bytes("gpieicl")
			var _1123_v52 _dafny.Map
			_ = _1123_v52
			_1123_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, _1122_v51)
			var _1124_v54 _dafny.Array
			_ = _1124_v54
			var _nwElement0_46 _dafny.Map = func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter37 := _dafny.Iterate(((_1117_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1117_v44).Cardinality()))).Uint32()).(_dafny.Sequence)).Elements()); ; {
					_compr_33, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1125_v41 _dafny.Map
					_1125_v41 = interface{}(_compr_33).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains((_1117_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1117_v44).Cardinality()))).Uint32()).(_dafny.Sequence), _1125_v41) {
						_coll33.Add(_1125_v41, p2)
					}
				}
				return _coll33.ToMap()
			}()
			_ = _nwElement0_46
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(24))
			_ = _nw202
			(_nw202).ArraySet1(_nwElement0_46, 0)
			(_nw202).ArraySet1((_1118_v45).Update(_1115_v42, p2), 1)
			(_nw202).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, (_1114_v40).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1114_v40).Cardinality()))).Uint32()).(bool)), 2)
			(_nw202).ArraySet1((_1118_v45).Update(_1115_v42, p3), 3)
			(_nw202).ArraySet1(_1118_v45, 4)
			(_nw202).ArraySet1((func() _dafny.Map {
				if (_1120_v47).Contains(Companion_Default___.Fm18(p2, p2, _1121_v48, globalState)) {
					return (_1120_v47).Get(Companion_Default___.Fm18(p2, p2, _1121_v48, globalState)).(_dafny.Map)
				}
				return func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter38 := _dafny.Iterate((_1116_v43).Elements()); ; {
						_compr_34, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _1126_v49 _dafny.Map
						_1126_v49 = interface{}(_compr_34).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_1116_v43, _1126_v49) {
							_coll34.Add(_1126_v49, p3)
						}
					}
					return _coll34.ToMap()
				}()
			})(), 5)
			(_nw202).ArraySet1(_1118_v45, 6)
			(_nw202).ArraySet1(_1118_v45, 7)
			(_nw202).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, p2), 8)
			(_nw202).ArraySet1(_1118_v45, 9)
			(_nw202).ArraySet1((_1118_v45).Update(_1115_v42, p3), 10)
			(_nw202).ArraySet1(_1118_v45, 11)
			(_nw202).ArraySet1((_1118_v45).Merge(Companion_Default___.Fm20(globalState)), 12)
			(_nw202).ArraySet1(func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter39 := _dafny.Iterate((_1123_v52).Keys().Elements()); ; {
					_compr_35, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _1127_v50 _dafny.Map
					_1127_v50 = interface{}(_compr_35).(_dafny.Map)
					if (_1123_v52).Contains(_1127_v50) {
						_coll35.Add(_1127_v50, p3)
					}
				}
				return _coll35.ToMap()
			}(), 13)
			(_nw202).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, p3), 14)
			(_nw202).ArraySet1((Companion_Default___.Fm20(globalState)).Merge(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter40 := _dafny.Iterate((_1118_v45).Keys().Elements()); ; {
					_compr_36, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _1128_v53 _dafny.Map
					_1128_v53 = interface{}(_compr_36).(_dafny.Map)
					if (_1118_v45).Contains(_1128_v53) {
						_coll36.Add(_1128_v53, p3)
					}
				}
				return _coll36.ToMap()
			}()), 15)
			(_nw202).ArraySet1(_1118_v45, 16)
			(_nw202).ArraySet1(_1118_v45, 17)
			(_nw202).ArraySet1((_1118_v45).Merge((_1118_v45).Update(_1115_v42, p3)), 18)
			(_nw202).ArraySet1(_1118_v45, 19)
			(_nw202).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, true), 20)
			(_nw202).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, (func() bool {
				if (_1115_v42).Contains(p3) {
					return (_1115_v42).Get(p3).(bool)
				}
				return p3
			})()), 21)
			(_nw202).ArraySet1((_1118_v45).Update(_1115_v42, p2), 22)
			(_nw202).ArraySet1(_1118_v45, 23)
			_1124_v54 = _nw202
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1124_v54), 0))
			_ = _index202
			(_1124_v54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v42, (_1114_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1122_v51).Cardinality()), _dafny.IntOfUint32((_1114_v40).Cardinality()))).Uint32()).(bool)), (_index202).Int())
			var _1129_v56 _dafny.Set
			_ = _1129_v56
			_1129_v56 = _dafny.SetOf(_1115_v42, Companion_Default___.Fm21(globalState))
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1124_v54), 0))
			_ = _index203
			(_1124_v54).ArraySet1(func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter41 := _dafny.Iterate(((_1129_v56).Intersection(_1129_v56)).Elements()); ; {
					_compr_37, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _1130_v55 _dafny.Map
					_1130_v55 = interface{}(_compr_37).(_dafny.Map)
					if ((_1129_v56).Intersection(_1129_v56)).Contains(_1130_v55) {
						_coll37.Add(_1130_v55, (func() bool {
							if (_1115_v42).Contains(p3) {
								return (_1115_v42).Get(p3).(bool)
							}
							return p3
						})())
					}
				}
				return _coll37.ToMap()
			}(), (_index203).Int())
			if p3 {
				(_this).F2_set_(p0)
				var _1131_v57 _dafny.Array
				_ = _1131_v57
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw203
				_1131_v57 = _nw203
				var _1132_v58 _dafny.Sequence
				_ = _1132_v58
				_1132_v58 = _dafny.SeqOf(_1131_v57, _1131_v57)
				r0 = _dafny.Companion_Sequence_.IsPrefixOf(_1132_v58, _1132_v58)
				r0 = (p0).Cmp((p0).Plus(p0)) != 0
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1131_v57), 0))
				_ = _index204
				(_1131_v57).ArraySet1((!(p3)) == (p3), (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1131_v57), 0))
				_ = _index205
				(_1131_v57).ArraySet1(p2, (_index205).Int())
				var _1133_v59 _dafny.Map
				_ = _1133_v59
				_1133_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
				var _1134_v60 _dafny.Map
				_ = _1134_v60
				_1134_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.Fm15(globalState))
				var _1135_v61 _dafny.CodePoint
				_ = _1135_v61
				_1135_v61 = _dafny.CodePoint('t')
				_1133_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).Cmp((_1134_v60).Cardinality()) < 0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1122_v51, _dafny.Companion_Sequence_.Update(_1122_v51, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F2()), _dafny.IntOfUint32((_1122_v51).Cardinality()))).Uint32(), _1135_v61))).Cardinality()))
			} else {
				var _1136_v62 _dafny.Array
				_ = _1136_v62
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_19
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.MultiSet = (func(_1137_v51 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
						return func(_1138_i3 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_1137_v51)).Intersection(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg60 _dafny.Int) interface{} {
									return coer60(arg60)
								}
							}(func(_1139_i4 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('s')
							}))))
						}
					})(_1122_v51)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw204 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw204).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw204).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_1136_v62 = _nw204
				var _1140_v63 _dafny.MultiSet
				_ = _1140_v63
				_1140_v63 = _dafny.MultiSetOf(_1122_v51, _1122_v51, _1122_v51)
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1136_v62), 0))
				_ = _index206
				(_1136_v62).ArraySet1((Companion_Default___.Fm22(p2, _this.F2(), p2, p0, globalState)).Difference(_1140_v63), (_index206).Int())
				var _1141_v64 _dafny.Array
				_ = _1141_v64
				var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw205
				_1141_v64 = _nw205
				var _1142_v65 _dafny.Map
				_ = _1142_v65
				_1142_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1141_v64, (func() _dafny.MultiSet {
					if p2 {
						return _1140_v63
					}
					return _1140_v63
				})())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1136_v62), 0))
				_ = _index207
				(_1136_v62).ArraySet1((func() _dafny.MultiSet {
					if (_1142_v65).Contains(_1141_v64) {
						return (_1142_v65).Get(_1141_v64).(_dafny.MultiSet)
					}
					return _1140_v63
				})(), (_index207).Int())
				var _1143_v66 *C1
				_ = _1143_v66
				var _nw206 *C1 = New_C1_()
				_ = _nw206
				_nw206.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1122_v51, _1122_v51))
				_1143_v66 = _nw206
				var _1144_v67 _dafny.CodePoint
				_ = _1144_v67
				_1144_v67 = _dafny.CodePoint('q')
				var _1145_v68 _dafny.Map
				_ = _1145_v68
				_1145_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _1146_v69 D5
				_ = _1146_v69
				_1146_v69 = Companion_D5_.Create_DC9_(false, p0, _1121_v48)
				var _1147_v70 D6
				_ = _1147_v70
				_1147_v70 = Companion_D6_.Create_DC13_(p3, p0)
				var _1148_v71 _dafny.Sequence
				_ = _1148_v71
				_1148_v71 = _dafny.SeqOf(_1147_v70, _1147_v70)
				var _1149_v72 _dafny.Sequence
				_ = _1149_v72
				_1149_v72 = _dafny.SeqOf(_this.F2(), _dafny.IntOfUint32((_1148_v71).Cardinality()))
				var _1150_v73 D4
				_ = _1150_v73
				_1150_v73 = Companion_D4_.Create_DC6_((_1143_v66).Fm1(p2, globalState), _1141_v64, _dafny.IntOfUint32((_1149_v72).Cardinality()))
				var _1151_v74 _dafny.Array
				_ = _1151_v74
				var _nwElement0_47 bool = p3
				_ = _nwElement0_47
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(24))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_47, 0)
				(_nw207).ArraySet1(p2, 1)
				(_nw207).ArraySet1(p2, 2)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}((func(_1152_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1153_i5 _dafny.Int) _dafny.CodePoint {
						return _1152_v67
					}
				})(_1144_v67))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1154_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1155_i5 _dafny.Int) _dafny.CodePoint {
						return _1154_v67
					}
				})(_1144_v67)))).Cardinality()))).Uint32(), _1144_v67), _1144_v67), 3)
				(_nw207).ArraySet1(p3, 4)
				(_nw207).ArraySet1(p2, 5)
				(_nw207).ArraySet1((func() bool {
					if (_1145_v68).Contains(_this.F2()) {
						return (_1145_v68).Get(_this.F2()).(bool)
					}
					return !(true)
				})(), 6)
				(_nw207).ArraySet1((_1143_v66).Fm11(globalState), 7)
				(_nw207).ArraySet1(p2, 8)
				(_nw207).ArraySet1((_this.F2()).Cmp(_this.F2()) <= 0, 9)
				(_nw207).ArraySet1(p2, 10)
				(_nw207).ArraySet1((_1146_v69).Dtor_cf14(), 11)
				(_nw207).ArraySet1(p3, 12)
				(_nw207).ArraySet1((p0).Cmp(p0) >= 0, 13)
				(_nw207).ArraySet1(p2, 14)
				(_nw207).ArraySet1((func() bool {
					if p3 {
						return !(p3)
					}
					return p2
				})(), 15)
				(_nw207).ArraySet1(p2, 16)
				(_nw207).ArraySet1((_1114_v40).Select((Companion_Default___.SafeIndex((_1150_v73).Dtor_cf11(), _dafny.IntOfUint32((_1114_v40).Cardinality()))).Uint32()).(bool), 17)
				(_nw207).ArraySet1(p3, 18)
				(_nw207).ArraySet1(p3, 19)
				(_nw207).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(p0, _this.F2()), _1149_v72)), 20)
				(_nw207).ArraySet1(p2, 21)
				(_nw207).ArraySet1(true, 22)
				(_nw207).ArraySet1(p3, 23)
				_1151_v74 = _nw207
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1151_v74), 0))
				_ = _index208
				(_1151_v74).ArraySet1(p3, (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1151_v74), 0))
				_ = _index209
				(_1151_v74).ArraySet1(p3, (_index209).Int())
				_1114_v40 = _1114_v40
				_1141_v64 = _1141_v64
			}
			var _1156_v75 _dafny.CodePoint
			_ = _1156_v75
			_1156_v75 = _dafny.CodePoint('m')
			_1156_v75 = _1156_v75
			r0 = p2
			var _1157_v76 _dafny.Map
			_ = _1157_v76
			_1157_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(879), Companion_Default___.Fm19(_this.F2(), p0, false, _this.F2(), globalState))
			var _1158_v77 *C0
			_ = _1158_v77
			var _nw208 *C0 = New_C0_()
			_ = _nw208
			_nw208.Ctor__((_dafny.IntOfInt64(-623)).Cmp((_1157_v76).Cardinality()) != 0, p2)
			_1158_v77 = _nw208
		} else {
			r0 = p2
			var _1159_v78 _dafny.CodePoint
			_ = _1159_v78
			_1159_v78 = _dafny.CodePoint('r')
			var _1160_v79 _dafny.Sequence
			_ = _1160_v79
			_1160_v79 = _dafny.SeqOf(_1159_v78)
			var _1161_v80 _dafny.Array
			_ = _1161_v80
			var _nwElement0_48 _dafny.CodePoint = _1159_v78
			_ = _nwElement0_48
			var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(10))
			_ = _nw209
			(_nw209).ArraySet1CodePoint(_nwElement0_48, 0)
			(_nw209).ArraySet1CodePoint(Companion_Default___.Fm23(globalState), 1)
			(_nw209).ArraySet1CodePoint((_1160_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wa")).Cardinality()), _dafny.IntOfUint32((_1160_v79).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
			(_nw209).ArraySet1CodePoint(_dafny.CodePoint('h'), 3)
			(_nw209).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p2 {
					return _1159_v78
				}
				return _1159_v78
			})(), 4)
			(_nw209).ArraySet1CodePoint(_1159_v78, 5)
			(_nw209).ArraySet1CodePoint(_1159_v78, 6)
			(_nw209).ArraySet1CodePoint(_1159_v78, 7)
			(_nw209).ArraySet1CodePoint(_1159_v78, 8)
			(_nw209).ArraySet1CodePoint(_1159_v78, 9)
			_1161_v80 = _nw209
			var _1162_v81 _dafny.Set
			_ = _1162_v81
			_1162_v81 = _dafny.SetOf(p0, p0, _dafny.IntOfUint32((_1160_v79).Cardinality()), p0)
			var _rhs168 _dafny.CodePoint = _1159_v78
			_ = _rhs168
			var _rhs169 _dafny.Array = _1161_v80
			_ = _rhs169
			var _rhs170 _dafny.Set = _1162_v81
			_ = _rhs170
			_1159_v78 = _rhs168
			_1161_v80 = _rhs169
			_1162_v81 = _rhs170
			var _1163_v82 D0
			_ = _1163_v82
			_1163_v82 = Companion_D0_.Create_DC0_(p2)
			r0 = (_1163_v82).Dtor_cf0()
			if (p0).Cmp((func() _dafny.Int {
				if p3 {
					return _this.F2()
				}
				return _this.F2()
			})()) < 0 {
				var _1164_v83 *C1
				_ = _1164_v83
				var _nw210 *C1 = New_C1_()
				_ = _nw210
				_nw210.Ctor__(_1160_v79)
				_1164_v83 = _nw210
				var _1165_v84 _dafny.Array
				_ = _1165_v84
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw211
				_1165_v84 = _nw211
				var _1166_v85 _dafny.Array
				_ = _1166_v85
				_1166_v85 = _1165_v84
				_1166_v85 = _1166_v85
				var _1167_v86 D6
				_ = _1167_v86
				_1167_v86 = Companion_D6_.Create_DC11_(_dafny.Companion_Sequence_.Update(_1160_v79, (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1160_v79).Cardinality()))).Uint32(), _1159_v78))
				var _1168_v87 D6
				_ = _1168_v87
				_1168_v87 = Companion_D6_.Create_DC11_((_1167_v86).Dtor_cf21())
				var _1169_v88 _dafny.Sequence
				_ = _1169_v88
				_1169_v88 = _dafny.SeqOf(p0)
				var _rhs171 bool = (_this).Fm8(_1160_v79, p2, globalState)
				_ = _rhs171
				var _rhs172 D6 = _1168_v87
				_ = _rhs172
				var _rhs173 _dafny.Int = Companion_Default___.SafeDivisionInt((_1169_v88).Select((Companion_Default___.SafeIndex((_1164_v83).Fm1(p2, globalState), _dafny.IntOfUint32((_1169_v88).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.SafeDivisionInt(p0, _this.F2()))
				_ = _rhs173
				var _lhs96 *C3 = _this
				_ = _lhs96
				r0 = _rhs171
				_1168_v87 = _rhs172
				_lhs96.F2_set_(_rhs173)
				(_this).F2_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).Fm10(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1164_v83).F3(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1164_v83).F3()).Cardinality()))).Uint32(), _dafny.CodePoint('r'))).Cardinality()), _dafny.IntOfInt64(-521)), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1164_v83).F3()).Cardinality()), _dafny.IntOfUint32(((_this).Fm10(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1164_v83).F3(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1164_v83).F3()).Cardinality()))).Uint32(), _dafny.CodePoint('r'))).Cardinality()), _dafny.IntOfInt64(-521)), globalState)).Cardinality()))).Uint32(), _1159_v78)).Cardinality()))
				var _1170_v89 _dafny.Map
				_ = _1170_v89
				_1170_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1164_v83).Fm1(true, globalState), _dafny.IntOfInt64(953))
				var _1171_v90 _dafny.Map
				_ = _1171_v90
				_1171_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p2))
				var _1172_v91 _dafny.Array
				_ = _1172_v91
				var _nwElement0_49 _dafny.Int = p0
				_ = _nwElement0_49
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(24))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_49, 0)
				(_nw212).ArraySet1((p0).Minus(p0), 1)
				(_nw212).ArraySet1((func() _dafny.Int {
					if (_1170_v89).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))) {
						return (_1170_v89).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).(_dafny.Int)
					}
					return _dafny.IntOfInt64(593)
				})(), 2)
				(_nw212).ArraySet1(_this.F2(), 3)
				(_nw212).ArraySet1(p0, 4)
				(_nw212).ArraySet1(p0, 5)
				(_nw212).ArraySet1(p0, 6)
				(_nw212).ArraySet1((_dafny.Zero).Minus(p0), 7)
				(_nw212).ArraySet1((_1164_v83).Fm1(p3, globalState), 8)
				(_nw212).ArraySet1(p0, 9)
				(_nw212).ArraySet1((p0).Plus(_dafny.IntOfInt64(144)), 10)
				(_nw212).ArraySet1((_this.F2()).Minus(p0), 11)
				(_nw212).ArraySet1((_this.F2()).Times(_this.F2()), 12)
				(_nw212).ArraySet1(_this.F2(), 13)
				(_nw212).ArraySet1(p0, 14)
				(_nw212).ArraySet1(p0, 15)
				(_nw212).ArraySet1(p0, 16)
				(_nw212).ArraySet1(_dafny.IntOfUint32((_1160_v79).Cardinality()), 17)
				(_nw212).ArraySet1((_1171_v90).Cardinality(), 18)
				(_nw212).ArraySet1(p0, 19)
				(_nw212).ArraySet1((_this).Fm1(p3, globalState), 20)
				(_nw212).ArraySet1(_this.F2(), 21)
				(_nw212).ArraySet1(_this.F2(), 22)
				(_nw212).ArraySet1(p0, 23)
				_1172_v91 = _nw212
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1172_v91), 0))
				_ = _index210
				(_1172_v91).ArraySet1(_this.F2(), (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1172_v91), 0))
				_ = _index211
				var _rhs174 _dafny.Int = ((_this).Fm1(p2, globalState)).Times(Companion_Default___.SafeModuloInt(_this.F2(), _dafny.IntOfUint32((_dafny.SeqOf(_1159_v78, _1159_v78)).Cardinality())))
				_ = _rhs174
				var _rhs175 _dafny.Int = (_1169_v88).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1169_v88).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs175
				var _rhs176 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1160_v79, _dafny.UnicodeSeqOfUtf8Bytes("cpr")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(451))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1173_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1174_i6 _dafny.Int) _dafny.CodePoint {
						return _1173_v78
					}
				})(_1159_v78))))
				_ = _rhs176
				var _rhs177 _dafny.Int = ((_dafny.IntOfInt64(529)).Plus(p0)).Times(_dafny.IntOfInt64(458))
				_ = _rhs177
				var _lhs97 _dafny.Array = _1172_v91
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1172_v91), 0))
				_ = _lhs98
				r2 = _rhs174
				r2 = _rhs175
				_1160_v79 = _rhs176
				(_lhs97).ArraySet1(_rhs177, (_lhs98).Int())
			} else {
				var _1175_v92 _dafny.MultiSet
				_ = _1175_v92
				_1175_v92 = _dafny.MultiSetOf((p3) && (p3))
				(_this).F2_set_((_1175_v92).Cardinality())
				_1160_v79 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1176_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1177_i7 _dafny.Int) _dafny.CodePoint {
						return _1176_v78
					}
				})(_1159_v78)))
				_1175_v92 = (_1175_v92).Difference(_1175_v92)
				var _1178_v93 _dafny.Array
				_ = _1178_v93
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw213
				_1178_v93 = _nw213
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_1178_v93), 0))
				_ = _index212
				(_1178_v93).ArraySet1(_dafny.IntOfUint32((_1160_v79).Cardinality()), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_1178_v93), 0))
				_ = _index213
				(_1178_v93).ArraySet1(p0, (_index213).Int())
				var _1179_v94 *C1
				_ = _1179_v94
				var _nw214 *C1 = New_C1_()
				_ = _nw214
				_nw214.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_1180_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1181_i8 _dafny.Int) _dafny.CodePoint {
						return _1180_v78
					}
				})(_1159_v78))))
				_1179_v94 = _nw214
			}
			var _1182_v95 _dafny.Map
			_ = _1182_v95
			_1182_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}((func(_1183_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1184_i9 _dafny.Int) _dafny.Int {
					return _1183_p0
				}
			})(p0))), p0)
			var _1185_v96 _dafny.Map
			_ = _1185_v96
			_1185_v96 = _1182_v95
			var _1186_v97 _dafny.Map
			_ = _1186_v97
			_1186_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1185_v96, _dafny.IntOfInt64(30))
			r0 = !(((p0).Plus(p0)).Cmp((func() _dafny.Int {
				if (_1186_v97).Contains(_1185_v96) {
					return (_1186_v97).Get(_1185_v96).(_dafny.Int)
				}
				return p0
			})()) < 0)
		}
		r2 = (_dafny.Zero).Minus(_this.F2())
		var _1187_v98 _dafny.Sequence
		_ = _1187_v98
		_1187_v98 = _dafny.UnicodeSeqOfUtf8Bytes("brhukyi")
		var _hi4 _dafny.Int = (_this.F2()).Plus(_this.F2())
		_ = _hi4
		for _1188_i10 := Companion_Default___.SafeModuloInt(_this.F2(), _dafny.IntOfUint32((_1187_v98).Cardinality())); _1188_i10.Cmp(_hi4) < 0; _1188_i10 = _1188_i10.Plus(_dafny.One) {
			if !(p3) {
				var _1189_v99 T2
				_ = _1189_v99
				var _nw215 *C2 = New_C2_()
				_ = _nw215
				_nw215.Ctor__((_this).Fm1(p2, globalState))
				_1189_v99 = _nw215
				var _1190_v100 T2
				_ = _1190_v100
				_1190_v100 = _1189_v99
				var _rhs178 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, _this.F2())
				_ = _rhs178
				var _rhs179 T2 = (_1190_v100)
				_ = _rhs179
				var _rhs180 bool = (_this).Fm8(_1187_v98, p3, globalState)
				_ = _rhs180
				r2 = _rhs178
				_1189_v99 = _rhs179
				r0 = _rhs180
				var _1191_v101 _dafny.CodePoint
				_ = _1191_v101
				_1191_v101 = _dafny.CodePoint('d')
				var _1192_v102 _dafny.MultiSet
				_ = _1192_v102
				_1192_v102 = _dafny.MultiSetOf(p2, p2, !(p2), true, p3)
				var _1193_v103 _dafny.Map
				_ = _1193_v103
				_1193_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1194_v104 _dafny.Map
				_ = _1194_v104
				_1194_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1193_v103, _1192_v102)
				var _1195_v105 D5
				_ = _1195_v105
				_1195_v105 = Companion_D5_.Create_DC9_(true, _dafny.IntOfInt64(571), _1192_v102)
				var _1196_v106 _dafny.Array
				_ = _1196_v106
				var _nwElement0_50 _dafny.MultiSet = _dafny.MultiSetOf(p2, p3)
				_ = _nwElement0_50
				var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(28))
				_ = _nw216
				(_nw216).ArraySet1(_nwElement0_50, 0)
				(_nw216).ArraySet1(_1192_v102, 1)
				(_nw216).ArraySet1(Companion_Default___.Fm38(p3, globalState), 2)
				(_nw216).ArraySet1((func() _dafny.MultiSet {
					if p2 {
						return Companion_Default___.Fm38(p2, globalState)
					}
					return _dafny.MultiSetOf(p2, p3)
				})(), 3)
				(_nw216).ArraySet1(_1192_v102, 4)
				(_nw216).ArraySet1(_1192_v102, 5)
				(_nw216).ArraySet1(_1192_v102, 6)
				(_nw216).ArraySet1(_1192_v102, 7)
				(_nw216).ArraySet1(_1192_v102, 8)
				(_nw216).ArraySet1(_1192_v102, 9)
				(_nw216).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1114_v40, _1114_v40)), 10)
				(_nw216).ArraySet1(_1192_v102, 11)
				(_nw216).ArraySet1((func() _dafny.MultiSet {
					if (_1189_v99).Fm8(_1187_v98, p2, globalState) {
						return _1192_v102
					}
					return _dafny.MultiSetFromSeq(_1114_v40)
				})(), 12)
				(_nw216).ArraySet1(_1192_v102, 13)
				(_nw216).ArraySet1((_1192_v102).Update(p3, Companion_Default___.Abs(p0)), 14)
				(_nw216).ArraySet1(_1192_v102, 15)
				(_nw216).ArraySet1(_1192_v102, 16)
				(_nw216).ArraySet1(_1192_v102, 17)
				(_nw216).ArraySet1(_1192_v102, 18)
				(_nw216).ArraySet1(_1192_v102, 19)
				(_nw216).ArraySet1(((func() _dafny.MultiSet {
					if (_1194_v104).Contains(_1193_v103) {
						return (_1194_v104).Get(_1193_v103).(_dafny.MultiSet)
					}
					return _1192_v102
				})()).Difference(_1192_v102), 20)
				(_nw216).ArraySet1((_1195_v105).Dtor_cf16(), 21)
				(_nw216).ArraySet1(_1192_v102, 22)
				(_nw216).ArraySet1((_1192_v102).Union(_1192_v102), 23)
				(_nw216).ArraySet1(_1192_v102, 24)
				(_nw216).ArraySet1(_1192_v102, 25)
				(_nw216).ArraySet1((_dafny.MultiSetOf((_this).Fm8(_1187_v98, p2, globalState))).Intersection(_1192_v102), 26)
				(_nw216).ArraySet1(_dafny.MultiSetOf(p2, p3), 27)
				_1196_v106 = _nw216
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1196_v106), 0))
				_ = _index214
				(_1196_v106).ArraySet1(_dafny.MultiSetOf(p3, p3), (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1196_v106), 0))
				_ = _index215
				var _rhs181 _dafny.CodePoint = (func() _dafny.CodePoint {
					if p2 {
						return _1191_v101
					}
					return _1191_v101
				})()
				_ = _rhs181
				var _rhs182 bool = true
				_ = _rhs182
				var _rhs183 _dafny.MultiSet = _1192_v102
				_ = _rhs183
				var _lhs99 _dafny.Array = _1196_v106
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1196_v106), 0))
				_ = _lhs100
				_1191_v101 = _rhs181
				r0 = _rhs182
				(_lhs99).ArraySet1(_rhs183, (_lhs100).Int())
				var _1197_v107 *C1
				_ = _1197_v107
				var _nw217 *C1 = New_C1_()
				_ = _nw217
				_nw217.Ctor__((Companion_D6_.Create_DC11_(_1187_v98)).Dtor_cf21())
				_1197_v107 = _nw217
				var _1198_v108 D8
				_ = _1198_v108
				_1198_v108 = Companion_D8_.Create_DC17_(p2, _1197_v107, _this.F2())
				var _1199_v109 _dafny.Array
				_ = _1199_v109
				var _nwElement0_51 bool = p2
				_ = _nwElement0_51
				var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(20))
				_ = _nw218
				(_nw218).ArraySet1(_nwElement0_51, 0)
				(_nw218).ArraySet1(p3, 1)
				(_nw218).ArraySet1(p3, 2)
				(_nw218).ArraySet1(p2, 3)
				(_nw218).ArraySet1(true, 4)
				(_nw218).ArraySet1(true, 5)
				(_nw218).ArraySet1(!(false), 6)
				(_nw218).ArraySet1(p2, 7)
				(_nw218).ArraySet1(!((_1198_v108).Dtor_cf30()), 8)
				(_nw218).ArraySet1(p2, 9)
				(_nw218).ArraySet1(p2, 10)
				(_nw218).ArraySet1(p2, 11)
				(_nw218).ArraySet1(p3, 12)
				(_nw218).ArraySet1(p3, 13)
				(_nw218).ArraySet1(p3, 14)
				(_nw218).ArraySet1(false, 15)
				(_nw218).ArraySet1(p3, 16)
				(_nw218).ArraySet1(p2, 17)
				(_nw218).ArraySet1((_1197_v107).Fm11(globalState), 18)
				(_nw218).ArraySet1(false, 19)
				_1199_v109 = _nw218
				var _1200_v110 _dafny.Set
				_ = _1200_v110
				_1200_v110 = _dafny.SetOf(_1199_v109)
				_1200_v110 = (_1200_v110).Intersection(_1200_v110)
				var _1201_v111 *C1
				_ = _1201_v111
				var _nw219 *C1 = New_C1_()
				_ = _nw219
				_nw219.Ctor__(_dafny.Companion_Sequence_.Concatenate((_1197_v107).F3(), _1187_v98))
				_1201_v111 = _nw219
				_1191_v101 = _1191_v101
			} else {
				_1114_v40 = _1114_v40
				var _1202_v112 _dafny.MultiSet
				_ = _1202_v112
				_1202_v112 = _dafny.MultiSetOf(!(p2), p2)
				_1202_v112 = _1202_v112
				var _1203_v113 _dafny.Map
				_ = _1203_v113
				_1203_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.SafeDivisionInt(_this.F2(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1187_v98).Cardinality()))))
				var _1204_v114 _dafny.Set
				_ = _1204_v114
				_1204_v114 = _dafny.SetOf(p3, p3, true)
				_1203_v113 = (_1203_v113).Update((_dafny.Zero).Minus((_this).Fm1(p3, globalState)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1204_v114).Cardinality(), p3)).Cardinality())
				var _1205_v115 _dafny.Array
				_ = _1205_v115
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_20
				var _nw220 _dafny.Array
				_ = _nw220
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw220 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Sequence = (func(_1206_v98 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1207_i11 _dafny.Int) _dafny.Sequence {
							return _1206_v98
						}
					})(_1187_v98)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw220 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw220).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw220).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_1205_v115 = _nw220
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1205_v115), 0))
				_ = _index216
				(_1205_v115).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tybfst"), (_index216).Int())
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1205_v115), 0))
				_ = _index217
				(_1205_v115).ArraySet1(_1187_v98, (_index217).Int())
				var _1208_v116 D0
				_ = _1208_v116
				_1208_v116 = Companion_D0_.Create_DC0_(p3)
				var _1209_v117 _dafny.Sequence
				_ = _1209_v117
				_1209_v117 = _dafny.SeqOf(_1208_v116)
				_1209_v117 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer67 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1210_v116 D0) func(_dafny.Int) D0 {
					return func(_1211_i12 _dafny.Int) D0 {
						return _1210_v116
					}
				})(_1208_v116)))
			}
			var _1212_v118 _dafny.CodePoint
			_ = _1212_v118
			_1212_v118 = _dafny.CodePoint('c')
			var _1213_v119 _dafny.MultiSet
			_ = _1213_v119
			_1213_v119 = _dafny.MultiSetOf(_1212_v118)
			_1213_v119 = _1213_v119
			var _1214_v120 _dafny.Array
			_ = _1214_v120
			var _nwElement0_52 bool = (_this).Fm8(_1187_v98, p3, globalState)
			_ = _nwElement0_52
			var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(16))
			_ = _nw221
			(_nw221).ArraySet1(_nwElement0_52, 0)
			(_nw221).ArraySet1(p3, 1)
			(_nw221).ArraySet1(p2, 2)
			(_nw221).ArraySet1(p2, 3)
			(_nw221).ArraySet1(p2, 4)
			(_nw221).ArraySet1(p3, 5)
			(_nw221).ArraySet1(p2, 6)
			(_nw221).ArraySet1(p3, 7)
			(_nw221).ArraySet1(p3, 8)
			(_nw221).ArraySet1(p2, 9)
			(_nw221).ArraySet1(p2, 10)
			(_nw221).ArraySet1(p2, 11)
			(_nw221).ArraySet1(p2, 12)
			(_nw221).ArraySet1(false, 13)
			(_nw221).ArraySet1(p2, 14)
			(_nw221).ArraySet1(p2, 15)
			_1214_v120 = _nw221
			var _1215_v121 _dafny.Set
			_ = _1215_v121
			_1215_v121 = _dafny.SetOf(_1214_v120)
			var _1216_v122 _dafny.Map
			_ = _1216_v122
			_1216_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1188_i10)
			var _1217_v123 _dafny.Set
			_ = _1217_v123
			_1217_v123 = _dafny.SetOf(p2)
			var _1218_v124 _dafny.Sequence
			_ = _1218_v124
			_1218_v124 = _dafny.SeqOf(_1217_v123)
			var _1219_v125 _dafny.Array
			_ = _1219_v125
			var _nwElement0_53 _dafny.Map = _1216_v122
			_ = _nwElement0_53
			var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(12))
			_ = _nw222
			(_nw222).ArraySet1(_nwElement0_53, 0)
			(_nw222).ArraySet1(_1216_v122, 1)
			(_nw222).ArraySet1(_1216_v122, 2)
			(_nw222).ArraySet1(_1216_v122, 3)
			(_nw222).ArraySet1((_1216_v122).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1188_i10)), 4)
			(_nw222).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(_1188_i10)), 5)
			(_nw222).ArraySet1(_1216_v122, 6)
			(_nw222).ArraySet1(_1216_v122, 7)
			(_nw222).ArraySet1(Companion_Default___.Fm27(_1188_i10, p2, globalState), 8)
			(_nw222).ArraySet1((_1216_v122).Update((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1218_v124).Cardinality()))), p0), 9)
			(_nw222).ArraySet1(_1216_v122, 10)
			(_nw222).ArraySet1(_1216_v122, 11)
			_1219_v125 = _nw222
			var _rhs184 _dafny.Set = (_1215_v121).Difference(_1215_v121)
			_ = _rhs184
			var _rhs185 bool = !(!(!((p2) && (false))))
			_ = _rhs185
			var _rhs186 _dafny.Array = _1219_v125
			_ = _rhs186
			var _rhs187 bool = p3
			_ = _rhs187
			var _rhs188 _dafny.Int = (_dafny.Zero).Minus((_this).Fm1(p2, globalState))
			_ = _rhs188
			var _lhs101 *C3 = _this
			_ = _lhs101
			_1215_v121 = _rhs184
			r0 = _rhs185
			_1219_v125 = _rhs186
			r0 = _rhs187
			_lhs101.F2_set_(_rhs188)
			r0 = p2
		}
		var _1220_v126 _dafny.Array
		_ = _1220_v126
		var _nw223 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw223
		_1220_v126 = _nw223
		_1220_v126 = _1220_v126
		var _1221_v127 _dafny.Array
		_ = _1221_v127
		var _nwElement0_54 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uokfbik")).Cardinality())).Minus(p0)
		_ = _nwElement0_54
		var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(2))
		_ = _nw224
		(_nw224).ArraySet1(_nwElement0_54, 0)
		(_nw224).ArraySet1(p0, 1)
		_1221_v127 = _nw224
		_1221_v127 = _1221_v127
		r0 = (func() bool {
			if p3 {
				return (p2) || (p2)
			}
			return p3
		})()
		var _1222_v128 _dafny.Map
		_ = _1222_v128
		_1222_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
		var _1223_v129 _dafny.Sequence
		_ = _1223_v129
		_1223_v129 = _dafny.SeqOf(_1222_v128)
		var _1224_v130 _dafny.Map
		_ = _1224_v130
		_1224_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p3)
		var _1225_v131 _dafny.Sequence
		_ = _1225_v131
		_1225_v131 = _dafny.SeqOf(_this.F2(), _this.F2(), p0, _dafny.IntOfInt64(-887))
		var _1226_v132 _dafny.MultiSet
		_ = _1226_v132
		_1226_v132 = _dafny.MultiSetOf(_1225_v131, _dafny.SeqOf(p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(916))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}(func(_1227_i13 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-392)
		})))
		var _1228_v133 _dafny.Map
		_ = _1228_v133
		_1228_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1226_v132, _1222_v128)
		var _1229_v134 _dafny.Array
		_ = _1229_v134
		var _nwElement0_55 _dafny.Map = _1222_v128
		_ = _nwElement0_55
		var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(18))
		_ = _nw225
		(_nw225).ArraySet1(_nwElement0_55, 0)
		(_nw225).ArraySet1((_1222_v128).Merge((_1223_v129).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1223_v129).Cardinality()))).Uint32()).(_dafny.Map)), 1)
		(_nw225).ArraySet1(_1222_v128, 2)
		(_nw225).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)).Update(p3, p3), 3)
		(_nw225).ArraySet1(_1222_v128, 4)
		(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2), 5)
		(_nw225).ArraySet1(_1222_v128, 6)
		(_nw225).ArraySet1(_1222_v128, 7)
		(_nw225).ArraySet1(_1222_v128, 8)
		(_nw225).ArraySet1(_1222_v128, 9)
		(_nw225).ArraySet1((_1222_v128).Merge(_1222_v128), 10)
		(_nw225).ArraySet1((_1222_v128).Merge(_1222_v128), 11)
		(_nw225).ArraySet1(_1222_v128, 12)
		(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3), 13)
		(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(p3)), 14)
		(_nw225).ArraySet1(_1222_v128, 15)
		(_nw225).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Merge((_1224_v130)), 16)
		(_nw225).ArraySet1((_1222_v128).Merge((func() _dafny.Map {
			if (_1228_v133).Contains(_1226_v132) {
				return (_1228_v133).Get(_1226_v132).(_dafny.Map)
			}
			return _1222_v128
		})()), 17)
		_1229_v134 = _nw225
		r1 = _1229_v134
		var _1230_v135 _dafny.Map
		_ = _1230_v135
		_1230_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1220_v126, _1187_v98)
		var _1231_v136 _dafny.CodePoint
		_ = _1231_v136
		_1231_v136 = _dafny.CodePoint('p')
		r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1230_v135).Contains(_1220_v126) {
				return (_1230_v135).Get(_1220_v126).(_dafny.Sequence)
			}
			return _1187_v98
		})(), _dafny.Companion_Sequence_.Concatenate(_1187_v98, _1187_v98)), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1230_v135).Contains(_1220_v126) {
				return (_1230_v135).Get(_1220_v126).(_dafny.Sequence)
			}
			return _1187_v98
		})(), _dafny.Companion_Sequence_.Concatenate(_1187_v98, _1187_v98))).Cardinality()))).Uint32(), _1231_v136)).Cardinality())
		return r0, r1, r2
	}
}
func (_this *C3) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1232_v0 _dafny.Array
		_ = _1232_v0
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_21
		var _nw226 _dafny.Array
		_ = _nw226
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw226 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) bool = (func(_1233_p2 bool) func(_dafny.Int) bool {
				return func(_1234_i0 _dafny.Int) bool {
					return _1233_p2
				}
			})(p2)
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw226 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw226).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw226).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_1232_v0 = _nw226
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
		_ = _index218
		(_1232_v0).ArraySet1(p2, (_index218).Int())
		var _1235_v1 bool
		_ = _1235_v1
		_1235_v1 = true
		var _1236_v2 _dafny.Map
		_ = _1236_v2
		_1236_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1235_v1)
		var _1237_v4 _dafny.Sequence
		_ = _1237_v4
		_1237_v4 = _dafny.UnicodeSeqOfUtf8Bytes("geclc")
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
		_ = _index219
		var _rhs189 _dafny.Int = _this.F2()
		_ = _rhs189
		var _rhs190 bool = (((_1236_v2).Merge(((func() _dafny.Map {
			var _coll38 = _dafny.NewMapBuilder()
			_ = _coll38
			for _iter42 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1238_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-31))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}(func(_1239_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality())
			}))).Elements()); ; {
				_compr_38, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _1240_v3 _dafny.Int
				_1240_v3 = interface{}(_compr_38).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}(func(_1238_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-31))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg72 _dafny.Int) interface{} {
							return coer72(arg72)
						}
					}(func(_1239_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality())
				})), _1240_v3) {
					_coll38.Add(Companion_Default___.SafeDivisionInt(_1240_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oner")).Cardinality())), p2)
				}
			}
			return _coll38.ToMap()
		}()).Update(p0, p2)).Update(p0, p2))).Cardinality()).Cmp(_this.F2()) < 0
		_ = _rhs190
		var _rhs191 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1237_v4, _1237_v4), _1237_v4)
		_ = _rhs191
		var _rhs192 _dafny.Int = (_dafny.Zero).Minus(p1)
		_ = _rhs192
		var _lhs102 _dafny.Array = _1232_v0
		_ = _lhs102
		var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
		_ = _lhs103
		var _lhs104 *C3 = _this
		_ = _lhs104
		r0 = _rhs189
		(_lhs102).ArraySet1(_rhs190, (_lhs103).Int())
		_1235_v1 = _rhs191
		_lhs104.F2_set_(_rhs192)
		var _1241_v5 _dafny.Set
		_ = _1241_v5
		_1241_v5 = _dafny.SetOf(_this.F2(), _this.F2())
		var _1242_v6 _dafny.MultiSet
		_ = _1242_v6
		_1242_v6 = _dafny.MultiSetOf(_1241_v5)
		var _1243_i3 _dafny.Int
		_ = _1243_i3
		_1243_i3 = _dafny.Zero
		{
			for ((func() _dafny.Int {
				if (_1242_v6).Contains(_1241_v5) {
					return (_1242_v6).Multiplicity(_1241_v5)
				}
				return p0
			})()).Cmp((_this).Fm1((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), globalState)) <= 0 {
				{
					if (_1243_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1243_i3 = (_1243_i3).Plus(_dafny.One)
					_1237_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1237_v4, _1237_v4), _1237_v4)
					if p2 {
						var _1244_v7 _dafny.Set
						_ = _1244_v7
						_1244_v7 = _dafny.SetOf((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool))
						_1244_v7 = _1244_v7
						var _1245_v8 *C1
						_ = _1245_v8
						var _nw227 *C1 = New_C1_()
						_ = _nw227
						_nw227.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("oxtqqto"))
						_1245_v8 = _nw227
						var _1246_v9 _dafny.Map
						_ = _1246_v9
						_1246_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p2, _1235_v1)).Cardinality(), _this.F2())
						var _1247_v11 _dafny.Sequence
						_ = _1247_v11
						_1247_v11 = _dafny.SeqOf(_1246_v9)
						var _1248_v12 _dafny.Array
						_ = _1248_v12
						var _nwElement0_56 _dafny.Map = _1246_v9
						_ = _nwElement0_56
						var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(9))
						_ = _nw228
						(_nw228).ArraySet1(_nwElement0_56, 0)
						(_nw228).ArraySet1(_1246_v9, 1)
						(_nw228).ArraySet1(_1246_v9, 2)
						(_nw228).ArraySet1((_1246_v9).Merge(_1246_v9), 3)
						(_nw228).ArraySet1(_1246_v9, 4)
						(_nw228).ArraySet1(_1246_v9, 5)
						(_nw228).ArraySet1(func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(216), _dafny.IntOfInt64(603))); ; {
								_compr_39, _ok43 := _iter43()
								if !_ok43 {
									break
								}
								var _1249_v10 _dafny.Int
								_1249_v10 = interface{}(_compr_39).(_dafny.Int)
								if ((_dafny.IntOfInt64(216)).Cmp(_1249_v10) <= 0) && ((_1249_v10).Cmp(_dafny.IntOfInt64(603)) < 0) {
									_coll39.Add((_1249_v10).Plus(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality())), p1)
								}
							}
							return _coll39.ToMap()
						}(), 6)
						(_nw228).ArraySet1((_1247_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1247_v11).Cardinality()))).Uint32()).(_dafny.Map), 7)
						(_nw228).ArraySet1(_1246_v9, 8)
						_1248_v12 = _nw228
						var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1248_v12), 0))
						_ = _index220
						(_1248_v12).ArraySet1(_1246_v9, (_index220).Int())
						var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1248_v12), 0))
						_ = _index221
						var _rhs193 _dafny.Map = _1246_v9
						_ = _rhs193
						var _rhs194 _dafny.Int = _this.F2()
						_ = _rhs194
						var _lhs105 _dafny.Array = _1248_v12
						_ = _lhs105
						var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1248_v12), 0))
						_ = _lhs106
						var _lhs107 *C3 = _this
						_ = _lhs107
						(_lhs105).ArraySet1(_rhs193, (_lhs106).Int())
						_lhs107.F2_set_(_rhs194)
						var _1250_v13 _dafny.Sequence
						_ = _1250_v13
						_1250_v13 = _dafny.SeqOf(_dafny.IntOfInt64(690), p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1245_v8).F3(), _1237_v4)).Cardinality()))
						_1250_v13 = _dafny.Companion_Sequence_.Concatenate(_1250_v13, _1250_v13)
						var _1251_v14 _dafny.Array
						_ = _1251_v14
						var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
						_ = _nw229
						_1251_v14 = _nw229
						_1251_v14 = _1251_v14
					} else {
						var _1252_v15 _dafny.Array
						_ = _1252_v15
						var _len0_22 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_22
						var _nw230 _dafny.Array
						_ = _nw230
						if _len0_22.Cmp(_dafny.Zero) == 0 {
							_nw230 = _dafny.NewArray(_len0_22)
						} else {
							var _init22 func(_dafny.Int) _dafny.Map = (func(_1253_v0 _dafny.Array) func(_dafny.Int) _dafny.Map {
								return func(_1254_i4 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1253_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1253_v0), 0))).Int()).(bool), true)
								}
							})(_1232_v0)
							_ = _init22
							var _element0_22 = _init22(_dafny.Zero)
							_ = _element0_22
							_nw230 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
							(_nw230).ArraySet1(_element0_22, 0)
							var _nativeLen0_22 = (_len0_22).Int()
							_ = _nativeLen0_22
							for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
								(_nw230).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
							}
						}
						_1252_v15 = _nw230
						var _1255_v16 _dafny.Sequence
						_ = _1255_v16
						_1255_v16 = _dafny.SeqOf(_1252_v15)
						_1252_v15 = (_1255_v16).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1255_v16).Cardinality()))).Uint32()).(_dafny.Array)
						var _1256_v17 _dafny.Array
						_ = _1256_v17
						var _len0_23 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_23
						var _nw231 _dafny.Array
						_ = _nw231
						if _len0_23.Cmp(_dafny.Zero) == 0 {
							_nw231 = _dafny.NewArray(_len0_23)
						} else {
							var _init23 func(_dafny.Int) _dafny.Int = func(_1257_i5 _dafny.Int) _dafny.Int {
								return (_1257_i5).Minus(_dafny.IntOfInt64(771))
							}
							_ = _init23
							var _element0_23 = _init23(_dafny.Zero)
							_ = _element0_23
							_nw231 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
							(_nw231).ArraySet1(_element0_23, 0)
							var _nativeLen0_23 = (_len0_23).Int()
							_ = _nativeLen0_23
							for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
								(_nw231).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
							}
						}
						_1256_v17 = _nw231
						var _1258_v18 _dafny.Sequence
						_ = _1258_v18
						_1258_v18 = _dafny.SeqOf(_1256_v17)
						var _1259_v19 _dafny.CodePoint
						_ = _1259_v19
						_1259_v19 = _dafny.CodePoint('i')
						var _1260_v20 D4
						_ = _1260_v20
						_1260_v20 = Companion_D4_.Create_DC6_(p1, _1256_v17, p0)
						var _1261_v21 _dafny.Array
						_ = _1261_v21
						var _nwElement0_57 _dafny.Array = _1256_v17
						_ = _nwElement0_57
						var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(20))
						_ = _nw232
						(_nw232).ArraySet1(_nwElement0_57, 0)
						(_nw232).ArraySet1((func() _dafny.Array {
							if p2 {
								return _1256_v17
							}
							return _1256_v17
						})(), 1)
						(_nw232).ArraySet1((func() _dafny.Array {
							if (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool) {
								return _1256_v17
							}
							return (_1258_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("udwhy")).Cardinality()), _dafny.IntOfUint32((_1258_v18).Cardinality()))).Uint32()).(_dafny.Array)
						})(), 2)
						(_nw232).ArraySet1((_1258_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1237_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1237_v4).Cardinality()))).Uint32(), _1259_v19), (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1237_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1237_v4).Cardinality()))).Uint32(), _1259_v19)).Cardinality()))).Uint32(), _1259_v19)).Cardinality()), _dafny.IntOfUint32((_1258_v18).Cardinality()))).Uint32()).(_dafny.Array), 3)
						(_nw232).ArraySet1(_1256_v17, 4)
						(_nw232).ArraySet1(_1256_v17, 5)
						(_nw232).ArraySet1(_1256_v17, 6)
						(_nw232).ArraySet1((func() _dafny.Array {
							if !(true) {
								return _1256_v17
							}
							return _1256_v17
						})(), 7)
						(_nw232).ArraySet1((_1260_v20).Dtor_cf10(), 8)
						(_nw232).ArraySet1(_1256_v17, 9)
						(_nw232).ArraySet1(_1256_v17, 10)
						(_nw232).ArraySet1(_1256_v17, 11)
						(_nw232).ArraySet1(_1256_v17, 12)
						(_nw232).ArraySet1(_1256_v17, 13)
						(_nw232).ArraySet1(_1256_v17, 14)
						(_nw232).ArraySet1(_1256_v17, 15)
						(_nw232).ArraySet1(_1256_v17, 16)
						(_nw232).ArraySet1(_1256_v17, 17)
						(_nw232).ArraySet1(_1256_v17, 18)
						(_nw232).ArraySet1(_1256_v17, 19)
						_1261_v21 = _nw232
						var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1261_v21), 0))
						_ = _index222
						(_1261_v21).ArraySet1((func() _dafny.Array {
							if _1235_v1 {
								return _1256_v17
							}
							return _1256_v17
						})(), (_index222).Int())
						var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1261_v21), 0))
						_ = _index223
						(_1261_v21).ArraySet1((_1260_v20).Dtor_cf10(), (_index223).Int())
						var _1262_v22 _dafny.Map
						_ = _1262_v22
						_1262_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1237_v4)
						var _1263_v23 _dafny.Sequence
						_ = _1263_v23
						_1263_v23 = _dafny.SeqOf(_this.F2(), _this.F2())
						var _1264_v24 _dafny.Sequence
						_ = _1264_v24
						_1264_v24 = _dafny.SeqOf(_1235_v1, _1235_v1)
						var _1265_v25 _dafny.Sequence
						_ = _1265_v25
						_1265_v25 = _dafny.SeqOf((_1263_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1264_v24).Cardinality()), _dafny.IntOfUint32((_1263_v23).Cardinality()))).Uint32()).(_dafny.Int))
						var _1266_v26 _dafny.Map
						_ = _1266_v26
						_1266_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
						_1237_v4 = (func() _dafny.Sequence {
							if (_1262_v22).Contains(p1) {
								return (_1262_v22).Get(p1).(_dafny.Sequence)
							}
							return _dafny.Companion_Sequence_.Update(_1237_v4, (Companion_Default___.SafeIndex((_dafny.SetOf(_1265_v25, _1265_v25, _dafny.SeqOf((_1266_v26).Cardinality()), _1265_v25)).Cardinality(), _dafny.IntOfUint32((_1237_v4).Cardinality()))).Uint32(), _1259_v19)
						})()
						var _1267_v27 _dafny.Array
						_ = _1267_v27
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_24
						var _nw233 _dafny.Array
						_ = _nw233
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw233 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) _dafny.Sequence = (func(_1268_v23 _dafny.Sequence, _1269_p1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
								return func(_1270_i6 _dafny.Int) _dafny.Sequence {
									return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1268_v23, _dafny.Companion_Sequence_.Update(_1268_v23, (Companion_Default___.SafeIndex(_1269_p1, _dafny.IntOfUint32((_1268_v23).Cardinality()))).Uint32(), _this.F2())), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1268_v23).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1268_v23, _dafny.Companion_Sequence_.Update(_1268_v23, (Companion_Default___.SafeIndex(_1269_p1, _dafny.IntOfUint32((_1268_v23).Cardinality()))).Uint32(), _this.F2()))).Cardinality()))).Uint32(), _this.F2())
								}
							})(_1263_v23, p1)
							_ = _init24
							var _element0_24 = _init24(_dafny.Zero)
							_ = _element0_24
							_nw233 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
							(_nw233).ArraySet1(_element0_24, 0)
							var _nativeLen0_24 = (_len0_24).Int()
							_ = _nativeLen0_24
							for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
								(_nw233).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
							}
						}
						_1267_v27 = _nw233
						var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1267_v27), 0))
						_ = _index224
						(_1267_v27).ArraySet1(_1263_v23, (_index224).Int())
						var _1271_v28 _dafny.Set
						_ = _1271_v28
						_1271_v28 = _dafny.SetOf(false)
						var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1267_v27), 0))
						_ = _index225
						(_1267_v27).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_1271_v28, _dafny.SetOf(p2))).Cardinality())), (_index225).Int())
						var _1272_v29 *C2
						_ = _1272_v29
						var _nw234 *C2 = New_C2_()
						_ = _nw234
						_nw234.Ctor__(_this.F2())
						_1272_v29 = _nw234
					}
					var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
					_ = _index226
					(_1232_v0).ArraySet1(p2, (_index226).Int())
					var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
					_ = _index227
					(_1232_v0).ArraySet1(_1235_v1, (_index227).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if _1235_v1 {
			_1235_v1 = _1235_v1
			r0 = (_dafny.Zero).Minus(_this.F2())
			r0 = p0
			(_this).F2_set_(_dafny.IntOfUint32((_1237_v4).Cardinality()))
			var _1273_v30 _dafny.Map
			_ = _1273_v30
			_1273_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _dafny.IntOfInt64(-750))
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index228
			(_1232_v0).ArraySet1((_dafny.IntOfInt64(-831)).Cmp((_1273_v30).Cardinality()) > 0, (_index228).Int())
		} else {
			var _1274_v31 _dafny.Map
			_ = _1274_v31
			_1274_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v1, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool))
			var _1275_v32 _dafny.Sequence
			_ = _1275_v32
			_1275_v32 = _dafny.SeqOf(_dafny.IntOfInt64(-366), _this.F2(), Companion_Default___.SafeModuloInt((_1274_v31).Cardinality(), (_1236_v2).Cardinality()), p0)
			_1275_v32 = _1275_v32
			r0 = (_this).Fm1(p2, globalState)
			var _1276_v33 _dafny.Set
			_ = _1276_v33
			_1276_v33 = _dafny.SetOf(p2, _1235_v1)
			var _1277_v34 _dafny.Sequence
			_ = _1277_v34
			_1277_v34 = _dafny.SeqOf(_dafny.SetOf(true, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)), _1276_v33, _1276_v33, _1276_v33)
			var _1278_v35 _dafny.Map
			_ = _1278_v35
			_1278_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Int {
				if (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool) {
					return p0
				}
				return ((_1277_v34).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1277_v34).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
			})())
			_1278_v35 = (_1278_v35).Update(p0, p0)
			var _1279_v36 _dafny.MultiSet
			_ = _1279_v36
			_1279_v36 = _dafny.MultiSetOf((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), (func() bool {
				if true {
					return p2
				}
				return _1235_v1
			})(), _1235_v1)
			var _rhs195 _dafny.Int = (func() _dafny.Int {
				if (_1279_v36).Contains((_1241_v5).IsSubsetOf(_1241_v5)) {
					return (_1279_v36).Multiplicity((_1241_v5).IsSubsetOf(_1241_v5))
				}
				return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_1237_v4).Cardinality()))).Times(p1)
			})()
			_ = _rhs195
			var _rhs196 _dafny.Int = Companion_Default___.SafeModuloInt((_1276_v33).Cardinality(), (_this.F2()).Times(p0))
			_ = _rhs196
			var _lhs108 *C3 = _this
			_ = _lhs108
			var _lhs109 *C3 = _this
			_ = _lhs109
			_lhs108.F2_set_(_rhs195)
			_lhs109.F2_set_(_rhs196)
			var _1280_v37 _dafny.Array
			_ = _1280_v37
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_25
			var _nw235 _dafny.Array
			_ = _nw235
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw235 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = (func(_1281_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1282_i7 _dafny.Int) _dafny.Int {
						return (_1282_i7).Plus(_1281_p0)
					}
				})(p0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw235 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw235).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw235).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1280_v37 = _nw235
			var _1283_v38 _dafny.MultiSet
			_ = _1283_v38
			_1283_v38 = _dafny.MultiSetOf(_dafny.IntOfInt64(993), _this.F2())
			var _1284_v39 _dafny.CodePoint
			_ = _1284_v39
			_1284_v39 = _dafny.CodePoint('j')
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index229
			var _rhs197 bool = _1235_v1
			_ = _rhs197
			var _rhs198 bool = (_this).Fm8(_dafny.Companion_Sequence_.Update(_1237_v4, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1237_v4).Cardinality()))).Uint32(), _1284_v39), (_1279_v36).IsProperSubsetOf(_1279_v36), globalState)
			_ = _rhs198
			var _rhs199 _dafny.Array = _1280_v37
			_ = _rhs199
			var _rhs200 _dafny.MultiSet = ((_1283_v38).Union(_1283_v38)).Union((_1283_v38).Union(_1283_v38))
			_ = _rhs200
			var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1275_v32, _1275_v32)
			_ = _rhs201
			var _lhs110 _dafny.Array = _1232_v0
			_ = _lhs110
			var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _lhs111
			(_lhs110).ArraySet1(_rhs197, (_lhs111).Int())
			_1235_v1 = _rhs198
			_1280_v37 = _rhs199
			_1283_v38 = _rhs200
			_1275_v32 = _rhs201
		}
		var _1285_v40 _dafny.Map
		_ = _1285_v40
		_1285_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(486))
		var _1286_v41 _dafny.Set
		_ = _1286_v41
		_1286_v41 = _dafny.SetOf(_1285_v40)
		var _hi5 _dafny.Int = (p1).Minus(p1)
		_ = _hi5
		for _1287_i8 := ((_dafny.SetOf((_1285_v40).Update(p2, p1), _1285_v40, _1285_v40)).Difference(_1286_v41)).Cardinality(); _1287_i8.Cmp(_hi5) < 0; _1287_i8 = _1287_i8.Plus(_dafny.One) {
			var _1288_v42 _dafny.Set
			_ = _1288_v42
			_1288_v42 = _dafny.SetOf((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool))
			if (_1288_v42).IsSubsetOf(_1288_v42) {
				var _1289_v43 _dafny.Array
				_ = _1289_v43
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw236
				_1289_v43 = _nw236
				var _1290_v44 _dafny.Set
				_ = _1290_v44
				_1290_v44 = _dafny.SetOf(_1289_v43, _1289_v43, _1289_v43)
				_1290_v44 = ((func() _dafny.Set {
					if false {
						return _dafny.SetOf(_1289_v43)
					}
					return _1290_v44
				})()).Intersection((_1290_v44).Union(_1290_v44))
				var _1291_v45 _dafny.Array
				_ = _1291_v45
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_26
				var _nw237 _dafny.Array
				_ = _nw237
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw237 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Map = (func(_1292_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1293_i9 _dafny.Int) _dafny.Map {
							return _1292_v2
						}
					})(_1236_v2)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw237 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw237).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw237).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1291_v45 = _nw237
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1291_v45), 0))
				_ = _index230
				(_1291_v45).ArraySet1(_1236_v2, (_index230).Int())
				var _1294_v46 _dafny.Map
				_ = _1294_v46
				_1294_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1241_v5).Cardinality(), (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)))
				var _1295_v47 D5
				_ = _1295_v47
				_1295_v47 = Companion_D5_.Create_DC10_(_1235_v1, p0, _1294_v46, _1235_v1)
				var _1296_v48 _dafny.Sequence
				_ = _1296_v48
				_1296_v48 = _dafny.SeqOf(_1235_v1)
				var _1297_v49 _dafny.MultiSet
				_ = _1297_v49
				_1297_v49 = _dafny.MultiSetOf(_1235_v1, (_1296_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1296_v48).Cardinality()), _dafny.IntOfUint32((_1296_v48).Cardinality()))).Uint32()).(bool))
				var _1298_v50 _dafny.CodePoint
				_ = _1298_v50
				_1298_v50 = _dafny.CodePoint('e')
				var _1299_v51 D0
				_ = _1299_v51
				_1299_v51 = Companion_D0_.Create_DC1_(p2, p2, _1298_v50, p1)
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1291_v45), 0))
				_ = _index231
				var _rhs202 _dafny.Int = _dafny.IntOfInt64(-912)
				_ = _rhs202
				var _rhs203 _dafny.Map = (func() _dafny.Map {
					if (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool) {
						return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v1, (_1297_v49).Cardinality())).Update((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), (_1241_v5).Cardinality())).Merge(_1285_v40)
					}
					return _1285_v40
				})()
				_ = _rhs203
				var _rhs204 _dafny.Map = ((_1236_v2).Merge(_1236_v2)).Update((_this).Fm1((_this).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("o"), false, globalState), globalState), (_1299_v51).Dtor_cf1())
				_ = _rhs204
				var _rhs205 D5 = _1295_v47
				_ = _rhs205
				var _lhs112 *C3 = _this
				_ = _lhs112
				var _lhs113 _dafny.Array = _1291_v45
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1291_v45), 0))
				_ = _lhs114
				_lhs112.F2_set_(_rhs202)
				_1285_v40 = _rhs203
				(_lhs113).ArraySet1(_rhs204, (_lhs114).Int())
				_1295_v47 = _rhs205
				var _1300_v52 _dafny.Array
				_ = _1300_v52
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw238
				_1300_v52 = _nw238
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_1300_v52), 0))
				_ = _index232
				(_1300_v52).ArraySet1(_1237_v4, (_index232).Int())
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_1300_v52), 0))
				_ = _index233
				(_1300_v52).ArraySet1(_1237_v4, (_index233).Int())
				var _1301_v53 *C1
				_ = _1301_v53
				var _nw239 *C1 = New_C1_()
				_ = _nw239
				_nw239.Ctor__(_1237_v4)
				_1301_v53 = _nw239
				_1301_v53 = _1301_v53
				var _1302_v54 _dafny.Map
				_ = _1302_v54
				_1302_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1296_v48)
				var _1303_v55 _dafny.Sequence
				_ = _1303_v55
				_1303_v55 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1296_v48, (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1296_v48).Cardinality()))).Uint32(), false), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1296_v48, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1296_v48).Cardinality()))).Uint32(), _1235_v1), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1296_v48, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1296_v48).Cardinality()))).Uint32(), _1235_v1)).Cardinality()))).Uint32(), !(true)), _1296_v48, _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1302_v54).Contains(_dafny.IntOfInt64(-989)) {
						return (_1302_v54).Get(_dafny.IntOfInt64(-989)).(_dafny.Sequence)
					}
					return _1296_v48
				})(), _dafny.Companion_Sequence_.Update(_1296_v48, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1296_v48).Cardinality()))).Uint32(), false)), _1296_v48)
				var _1304_v56 _dafny.MultiSet
				_ = _1304_v56
				_1304_v56 = _dafny.MultiSetOf((_1288_v42).Cardinality(), _this.F2())
				var _1305_v57 _dafny.Map
				_ = _1305_v57
				_1305_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v1, _1235_v1)
				var _1306_v58 D8
				_ = _1306_v58
				_1306_v58 = Companion_D8_.Create_DC17_((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _1301_v53, (_1305_v57).Cardinality())
				_1296_v48 = (_1303_v55).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1304_v56).Contains((_1306_v58).Dtor_cf32()) {
						return (_1304_v56).Multiplicity((_1306_v58).Dtor_cf32())
					}
					return p0
				})(), _dafny.IntOfUint32((_1303_v55).Cardinality()))).Uint32()).(_dafny.Sequence)
			} else {
				var _1307_v59 _dafny.MultiSet
				_ = _1307_v59
				_1307_v59 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qsrqrrepm"))
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
				_ = _index234
				(_1232_v0).ArraySet1((Companion_Default___.Fm30(_1235_v1, globalState)).IsSubsetOf((_1307_v59).Union(_1307_v59)), (_index234).Int())
				var _1308_v60 _dafny.Array
				_ = _1308_v60
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_27
				var _nw240 _dafny.Array
				_ = _nw240
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw240 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Int = (func(_1309_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1310_i10 _dafny.Int) _dafny.Int {
							return (_1310_i10).Minus(_dafny.IntOfUint32((_1309_v4).Cardinality()))
						}
					})(_1237_v4)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw240 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw240).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw240).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1308_v60 = _nw240
				_1308_v60 = _1308_v60
				var _1311_v61 _dafny.MultiSet
				_ = _1311_v61
				_1311_v61 = _dafny.MultiSetOf(_this.F2())
				var _1312_v62 _dafny.CodePoint
				_ = _1312_v62
				_1312_v62 = _dafny.CodePoint('p')
				var _1313_v63 _dafny.MultiSet
				_ = _1313_v63
				_1313_v63 = _dafny.MultiSetOf(_1235_v1)
				var _1314_v64 D0
				_ = _1314_v64
				_1314_v64 = Companion_D0_.Create_DC1_((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _1235_v1, _1312_v62, p0)
				var _1315_v65 _dafny.Map
				_ = _1315_v65
				_1315_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.MultiSetOf((_1314_v64).Dtor_cf4()))
				var _1316_v66 _dafny.MultiSet
				_ = _1316_v66
				_1316_v66 = _1311_v61
				var _1317_v67 _dafny.Sequence
				_ = _1317_v67
				_1317_v67 = _dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_this.F2())))
				var _1318_v68 _dafny.Array
				_ = _1318_v68
				var _nwElement0_58 _dafny.MultiSet = (_1311_v61).Intersection(_1311_v61)
				_ = _nwElement0_58
				var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(25))
				_ = _nw241
				(_nw241).ArraySet1(_nwElement0_58, 0)
				(_nw241).ArraySet1(_1311_v61, 1)
				(_nw241).ArraySet1((_dafny.MultiSetOf(_this.F2(), p1)).Difference(_1311_v61), 2)
				(_nw241).ArraySet1(_1311_v61, 3)
				(_nw241).ArraySet1((Companion_Default___.Fm35(_1312_v62, _1235_v1, _1313_v63, globalState)).Intersection(_1311_v61), 4)
				(_nw241).ArraySet1(_1311_v61, 5)
				(_nw241).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1319_i11 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(924)
				})))).Union(_1311_v61), 6)
				(_nw241).ArraySet1(_dafny.MultiSetOf(p0, p0, p0, p0, _this.F2()), 7)
				(_nw241).ArraySet1(Companion_Default___.Fm19(_dafny.IntOfInt64(712), p0, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _this.F2(), globalState), 8)
				(_nw241).ArraySet1((_1311_v61).Difference((func() _dafny.MultiSet {
					if (_1315_v65).Contains(_1312_v62) {
						return (_1315_v65).Get(_1312_v62).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf((_1236_v2).Cardinality(), p1)
				})()), 9)
				(_nw241).ArraySet1(_1311_v61, 10)
				(_nw241).ArraySet1((_1311_v61).Intersection(_1311_v61), 11)
				(_nw241).ArraySet1(((_1311_v61).Update(p1, Companion_Default___.Abs(p1))).Intersection((_1316_v66)), 12)
				(_nw241).ArraySet1(_1311_v61, 13)
				(_nw241).ArraySet1((_1317_v67).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1317_v67).Cardinality()))).Uint32()).(_dafny.MultiSet), 14)
				(_nw241).ArraySet1(_1311_v61, 15)
				(_nw241).ArraySet1(_1311_v61, 16)
				(_nw241).ArraySet1(_1311_v61, 17)
				(_nw241).ArraySet1((func() _dafny.MultiSet {
					if true {
						return _dafny.MultiSetOf((_this).Fm1(p2, globalState))
					}
					return _1311_v61
				})(), 18)
				(_nw241).ArraySet1(Companion_Default___.Fm35(_1312_v62, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _1313_v63, globalState), 19)
				(_nw241).ArraySet1(_1311_v61, 20)
				(_nw241).ArraySet1((_1311_v61).Update(_1287_i8, Companion_Default___.Abs(_dafny.IntOfInt64(55))), 21)
				(_nw241).ArraySet1(_dafny.MultiSetOf(_1287_i8), 22)
				(_nw241).ArraySet1(_1311_v61, 23)
				(_nw241).ArraySet1(_1311_v61, 24)
				_1318_v68 = _nw241
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1318_v68), 0))
				_ = _index235
				(_1318_v68).ArraySet1((_1311_v61).Difference(_1311_v61), (_index235).Int())
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1318_v68), 0))
				_ = _index236
				(_1318_v68).ArraySet1(_1311_v61, (_index236).Int())
				_1312_v62 = _1312_v62
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
				_ = _index237
				(_1232_v0).ArraySet1((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), (_index237).Int())
			}
			var _1320_v69 _dafny.CodePoint
			_ = _1320_v69
			_1320_v69 = _dafny.CodePoint('s')
			var _1321_v71 _dafny.Array
			_ = _1321_v71
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_28
			var _nw242 _dafny.Array
			_ = _nw242
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw242 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Map = (func(_1322_p0 _dafny.Int, _1323_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1324_i12 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
							var _coll40 = _dafny.NewMapBuilder()
							_ = _coll40
							for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(895), _dafny.IntOfInt64(618))); ; {
								_compr_40, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _1325_v70 _dafny.Int
								_1325_v70 = interface{}(_compr_40).(_dafny.Int)
								if ((_dafny.IntOfInt64(895)).Cmp(_1325_v70) <= 0) && ((_1325_v70).Cmp(_dafny.IntOfInt64(618)) < 0) {
									_coll40.Add((_1325_v70).Times(_dafny.IntOfInt64(190)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_p1, _this.F2()))
								}
							}
							return _coll40.ToMap()
						}()).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1323_p1), _1322_p0))).Cardinality(), _1322_p0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _this.F2()))
					}
				})(p0, p1)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw242 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw242).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw242).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1321_v71 = _nw242
			var _1326_v73 _dafny.MultiSet
			_ = _1326_v73
			_1326_v73 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1237_v4).Cardinality()))
			var _1327_v74 _dafny.Map
			_ = _1327_v74
			_1327_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1326_v73).Contains(_1287_i8) {
					return (_1326_v73).Multiplicity(_1287_i8)
				}
				return _dafny.IntOfInt64(47)
			})(), _1287_i8)
			var _1328_v75 _dafny.Sequence
			_ = _1328_v75
			_1328_v75 = _dafny.SeqOf(_1327_v74)
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1321_v71), 0))
			_ = _index238
			(_1321_v71).ArraySet1(func() _dafny.Map {
				var _coll41 = _dafny.NewMapBuilder()
				_ = _coll41
				for _iter45 := _dafny.Iterate(((_1328_v75).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1328_v75).Cardinality()))).Uint32()).(_dafny.Map)).Keys().Elements()); ; {
					_compr_41, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _1329_v72 _dafny.Int
					_1329_v72 = interface{}(_compr_41).(_dafny.Int)
					if ((_1328_v75).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1328_v75).Cardinality()))).Uint32()).(_dafny.Map)).Contains(_1329_v72) {
						_coll41.Add((_1329_v72).Plus((_dafny.Zero).Minus(_this.F2())), p1)
					}
				}
				return _coll41.ToMap()
			}(), (_index238).Int())
			var _1330_v76 _dafny.Map
			_ = _1330_v76
			_1330_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _1235_v1)
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1321_v71), 0))
			_ = _index239
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index240
			var _rhs206 _dafny.Int = Companion_Default___.SafeDivisionInt((_1330_v76).Cardinality(), p1)
			_ = _rhs206
			var _rhs207 _dafny.CodePoint = _dafny.CodePoint('u')
			_ = _rhs207
			var _rhs208 _dafny.Map = func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(105), _dafny.IntOfInt64(303))); ; {
					_compr_42, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _1331_v77 _dafny.Int
					_1331_v77 = interface{}(_compr_42).(_dafny.Int)
					if ((_dafny.IntOfInt64(105)).Cmp(_1331_v77) <= 0) && ((_1331_v77).Cmp(_dafny.IntOfInt64(303)) < 0) {
						_coll42.Add((_1331_v77).Times(((_1330_v76).Update(p2, _1235_v1)).Cardinality()), _1287_i8)
					}
				}
				return _coll42.ToMap()
			}()
			_ = _rhs208
			var _rhs209 bool = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}(func(_1332_i13 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(169))).Cardinality()))
			}))).Cardinality())).Cmp(_this.F2()) == 0
			_ = _rhs209
			var _rhs210 bool = false
			_ = _rhs210
			var _lhs115 _dafny.Array = _1321_v71
			_ = _lhs115
			var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1321_v71), 0))
			_ = _lhs116
			var _lhs117 _dafny.Array = _1232_v0
			_ = _lhs117
			var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _lhs118
			r0 = _rhs206
			_1320_v69 = _rhs207
			(_lhs115).ArraySet1(_rhs208, (_lhs116).Int())
			(_lhs117).ArraySet1(_rhs209, (_lhs118).Int())
			_1235_v1 = _rhs210
			var _1333_v78 _dafny.Sequence
			_ = _1333_v78
			_1333_v78 = _dafny.SeqOf(_1235_v1)
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index241
			(_1232_v0).ArraySet1((func() bool {
				if p2 {
					return (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)
				}
				return (_1333_v78).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1285_v40).Contains((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)) {
						return (_1285_v40).Get((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)).(_dafny.Int)
					}
					return _1287_i8
				})(), _dafny.IntOfUint32((_1333_v78).Cardinality()))).Uint32()).(bool)
			})(), (_index241).Int())
			var _1334_v79 _dafny.Map
			_ = _1334_v79
			_1334_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v4, ((_1288_v42).Intersection(_1288_v42)).Cardinality())
			_1334_v79 = (_1334_v79).Update(_1237_v4, _dafny.IntOfInt64(255))
		}
		(_this).F2_set_(p1)
		var _1335_v80 _dafny.Sequence
		_ = _1335_v80
		_1335_v80 = _dafny.SeqOf(p0, _this.F2(), p1)
		var _1336_v81 _dafny.Sequence
		_ = _1336_v81
		_1336_v81 = _dafny.SeqOf(_1335_v80)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_1336_v81, _1336_v81) {
			_1235_v1 = p2
			var _1337_v82 _dafny.CodePoint
			_ = _1337_v82
			_1337_v82 = _dafny.CodePoint('k')
			r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1237_v4, (Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1237_v4).Cardinality()))).Uint32(), _1337_v82)).Cardinality())
			var _1338_v83 _dafny.Array
			_ = _1338_v83
			var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw243
			_1338_v83 = _nw243
			var _1339_v84 _dafny.Map
			_ = _1339_v84
			_1339_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(_1237_v4, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), globalState), _1235_v1)
			var _1340_v85 _dafny.Map
			_ = _1340_v85
			_1340_v85 = _1339_v84
			var _1341_v86 _dafny.Map
			_ = _1341_v86
			_1341_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1340_v85)
			var _1342_v87 _dafny.Sequence
			_ = _1342_v87
			_1342_v87 = _dafny.SeqOf(_1341_v86)
			var _1343_v88 _dafny.Array
			_ = _1343_v88
			var _nwElement0_59 _dafny.Map = (_1342_v87).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1342_v87).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _nwElement0_59
			var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(10))
			_ = _nw244
			(_nw244).ArraySet1(_nwElement0_59, 0)
			(_nw244).ArraySet1(_1341_v86, 1)
			(_nw244).ArraySet1(_1341_v86, 2)
			(_nw244).ArraySet1(_1341_v86, 3)
			(_nw244).ArraySet1(_1341_v86, 4)
			(_nw244).ArraySet1(_1341_v86, 5)
			(_nw244).ArraySet1(_1341_v86, 6)
			(_nw244).ArraySet1(Companion_Default___.Fm39(globalState), 7)
			(_nw244).ArraySet1(_1341_v86, 8)
			(_nw244).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(449), _1339_v84), 9)
			_1343_v88 = _nw244
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1338_v83), 0))
			_ = _index242
			(_1338_v83).ArraySet1(_1343_v88, (_index242).Int())
			var _1344_v89 _dafny.Map
			_ = _1344_v89
			_1344_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1241_v5)
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1338_v83), 0))
			_ = _index243
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index244
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index245
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index246
			var _rhs211 _dafny.Array = _1343_v88
			_ = _rhs211
			var _rhs212 bool = (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)
			_ = _rhs212
			var _rhs213 bool = ((_1241_v5).Intersection(_1241_v5)).IsProperSubsetOf(((func() _dafny.Set {
				if (_1344_v89).Contains(_1235_v1) {
					return (_1344_v89).Get(_1235_v1).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll43 = _dafny.NewBuilder()
					_ = _coll43
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(56), _dafny.IntOfInt64(123))); ; {
						_compr_43, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1345_v90 _dafny.Int
						_1345_v90 = interface{}(_compr_43).(_dafny.Int)
						if ((_dafny.IntOfInt64(56)).Cmp(_1345_v90) <= 0) && ((_1345_v90).Cmp(_dafny.IntOfInt64(123)) < 0) {
							_coll43.Add(Companion_Default___.SafeDivisionInt(_1345_v90, _dafny.IntOfUint32((_1237_v4).Cardinality())))
						}
					}
					return _coll43.ToSet()
				}()
			})()).Difference(_1241_v5))
			_ = _rhs213
			var _rhs214 bool = !(!(_1236_v2).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F2()))))
			_ = _rhs214
			var _lhs119 _dafny.Array = _1338_v83
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1338_v83), 0))
			_ = _lhs120
			var _lhs121 _dafny.Array = _1232_v0
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _lhs122
			var _lhs123 _dafny.Array = _1232_v0
			_ = _lhs123
			var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _lhs124
			var _lhs125 _dafny.Array = _1232_v0
			_ = _lhs125
			var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _lhs126
			(_lhs119).ArraySet1(_rhs211, (_lhs120).Int())
			(_lhs121).ArraySet1(_rhs212, (_lhs122).Int())
			(_lhs123).ArraySet1(_rhs213, (_lhs124).Int())
			(_lhs125).ArraySet1(_rhs214, (_lhs126).Int())
			_1235_v1 = (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)
			var _1346_v91 _dafny.Array
			_ = _1346_v91
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw245
			_1346_v91 = _nw245
			var _1347_v92 _dafny.MultiSet
			_ = _1347_v92
			_1347_v92 = _dafny.MultiSetOf(_this.F2())
			var _1348_v93 _dafny.MultiSet
			_ = _1348_v93
			_1348_v93 = _1347_v92
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_1346_v91), 0))
			_ = _index247
			(_1346_v91).ArraySet1(_1348_v93, (_index247).Int())
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_1346_v91), 0))
			_ = _index248
			(_1346_v91).ArraySet1(_1347_v92, (_index248).Int())
		} else {
			if ((p0).Times(p0)).Cmp((func() _dafny.Int {
				if p2 {
					return p0
				}
				return p0
			})()) >= 0 {
				r0 = p1
				var _1349_v94 *C0
				_ = _1349_v94
				var _nw246 *C0 = New_C0_()
				_ = _nw246
				_nw246.Ctor__(p2, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool))
				_1349_v94 = _nw246
				var _1350_v95 _dafny.Sequence
				_ = _1350_v95
				_1350_v95 = _dafny.SeqOf(_1237_v4)
				var _1351_v96 _dafny.Set
				_ = _1351_v96
				_1351_v96 = _dafny.SetOf(_1237_v4, (_1350_v95).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1350_v95).Cardinality()))).Uint32()).(_dafny.Sequence))
				_1351_v96 = _1351_v96
				var _1352_v97 _dafny.MultiSet
				_ = _1352_v97
				_1352_v97 = _dafny.MultiSetOf(_1349_v94.F4)
				var _1353_v98 _dafny.Sequence
				_ = _1353_v98
				_1353_v98 = _dafny.SeqOf(_1352_v97)
				var _1354_v99 _dafny.Array
				_ = _1354_v99
				var _nwElement0_60 _dafny.Int = _dafny.IntOfInt64(746)
				_ = _nwElement0_60
				var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(15))
				_ = _nw247
				(_nw247).ArraySet1(_nwElement0_60, 0)
				(_nw247).ArraySet1(p0, 1)
				(_nw247).ArraySet1(p0, 2)
				(_nw247).ArraySet1(_dafny.IntOfInt64(-554), 3)
				(_nw247).ArraySet1((_dafny.Zero).Minus(p1), 4)
				(_nw247).ArraySet1(p1, 5)
				(_nw247).ArraySet1(_this.F2(), 6)
				(_nw247).ArraySet1(_dafny.IntOfInt64(203), 7)
				(_nw247).ArraySet1(p0, 8)
				(_nw247).ArraySet1(_this.F2(), 9)
				(_nw247).ArraySet1(p0, 10)
				(_nw247).ArraySet1(_this.F2(), 11)
				(_nw247).ArraySet1(_this.F2(), 12)
				(_nw247).ArraySet1(p0, 13)
				(_nw247).ArraySet1(_dafny.IntOfInt64(-538), 14)
				_1354_v99 = _nw247
				var _1355_v100 _dafny.Map
				_ = _1355_v100
				_1355_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pcgsm"), _1354_v99)
				var _1356_v101 _dafny.CodePoint
				_ = _1356_v101
				_1356_v101 = _dafny.CodePoint('q')
				r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1353_v98, _1353_v98), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm40((_1355_v100).Cardinality(), _1356_v101, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), p1, globalState), _1353_v98))).Cardinality())
				(_1349_v94).F4 = (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool)
			} else {
				_1235_v1 = !((_dafny.IntOfInt64(861)).Cmp((_this).Fm1((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), globalState)) > 0)
				_1241_v5 = _1241_v5
				r0 = (Companion_Default___.SafeModuloInt(p1, (_dafny.Zero).Minus(p0))).Plus((_dafny.Zero).Minus(_this.F2()))
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
				_ = _index249
				(_1232_v0).ArraySet1(p2, (_index249).Int())
				_1235_v1 = p2
			}
			var _1357_v102 _dafny.Sequence
			_ = _1357_v102
			_1357_v102 = _dafny.SeqOf(_1237_v4)
			var _1358_v104 _dafny.Set
			_ = _1358_v104
			_1358_v104 = _dafny.SetOf(_1237_v4)
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))
			_ = _index250
			(_1232_v0).ArraySet1((func() _dafny.Set {
				var _coll44 = _dafny.NewBuilder()
				_ = _coll44
				for _iter48 := _dafny.Iterate((_1357_v102).Elements()); ; {
					_compr_44, _ok48 := _iter48()
					if !_ok48 {
						break
					}
					var _1359_v103 _dafny.Sequence
					_1359_v103 = interface{}(_compr_44).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_1357_v102, _1359_v103) {
						_coll44.Add(_1359_v103)
					}
				}
				return _coll44.ToSet()
			}()).Equals(_1358_v104), (_index250).Int())
			var _1360_v105 _dafny.Map
			_ = _1360_v105
			_1360_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0, _dafny.IntOfInt64(167)), _dafny.IntOfUint32((_1335_v80).Cardinality()))
			var _1361_v106 _dafny.Map
			_ = _1361_v106
			_1361_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1360_v105)
			var _1362_v107 _dafny.MultiSet
			_ = _1362_v107
			_1362_v107 = _dafny.MultiSetOf((_1361_v106).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(685), _1360_v105)), _1361_v106)
			var _pat_let_tv28 = _1335_v80
			_ = _pat_let_tv28
			_1362_v107 = _dafny.MultiSetOf(_1361_v106, _1361_v106, func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(517), _dafny.IntOfInt64(472))); ; {
					_compr_45, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _1363_v108 _dafny.Int
					_1363_v108 = interface{}(_compr_45).(_dafny.Int)
					if ((_dafny.IntOfInt64(517)).Cmp(_1363_v108) <= 0) && ((_1363_v108).Cmp(_dafny.IntOfInt64(472)) < 0) {
						_coll45.Add((_1363_v108).Times((_1335_v80).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_dafny.SeqOf(p2))).Cardinality(), _dafny.IntOfUint32((_1335_v80).Cardinality()))).Uint32()).(_dafny.Int)), func(_pat_let26_0 _dafny.Map) _dafny.Map {
							return func(_1364_dt__update__tmp_h1 _dafny.Map) _dafny.Map {
								return func(_pat_let27_0 _dafny.Map) _dafny.Map {
									return func(_1365_dt__update_hcf6_h0 _dafny.Map) _dafny.Map {
										return _1365_dt__update_hcf6_h0
									}(_pat_let27_0)
								}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv28, _this.F2()))
							}(_pat_let26_0)
						}(_1360_v105))
					}
				}
				return _coll45.ToMap()
			}(), _1361_v106, _1361_v106)
			var _1366_v109 _dafny.Array
			_ = _1366_v109
			var _nwElement0_61 _dafny.Int = p0
			_ = _nwElement0_61
			var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(6))
			_ = _nw248
			(_nw248).ArraySet1(_nwElement0_61, 0)
			(_nw248).ArraySet1(_this.F2(), 1)
			(_nw248).ArraySet1((_this).Fm1((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), globalState), 2)
			(_nw248).ArraySet1(((Companion_D6_.Create_DC13_(p2, _this.F2())).Dtor_cf25()).Plus(_dafny.IntOfInt64(825)), 3)
			(_nw248).ArraySet1(p1, 4)
			(_nw248).ArraySet1((func() _dafny.Int {
				if (_1285_v40).Contains(p2) {
					return (_1285_v40).Get(p2).(_dafny.Int)
				}
				return p1
			})(), 5)
			_1366_v109 = _nw248
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_1366_v109), 0))
			_ = _index251
			(_1366_v109).ArraySet1(p1, (_index251).Int())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_1366_v109), 0))
			_ = _index252
			(_1366_v109).ArraySet1((func() _dafny.Int {
				if !(p2) {
					return p0
				}
				return p0
			})(), (_index252).Int())
			var _1367_v110 _dafny.CodePoint
			_ = _1367_v110
			_1367_v110 = _dafny.CodePoint('a')
			var _1368_v111 _dafny.Array
			_ = _1368_v111
			var _nwElement0_62 _dafny.CodePoint = _1367_v110
			_ = _nwElement0_62
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(9))
			_ = _nw249
			(_nw249).ArraySet1CodePoint(_nwElement0_62, 0)
			(_nw249).ArraySet1CodePoint(_dafny.CodePoint('k'), 1)
			(_nw249).ArraySet1CodePoint(Companion_Default___.Fm37(p1, globalState), 2)
			(_nw249).ArraySet1CodePoint(_1367_v110, 3)
			(_nw249).ArraySet1CodePoint(_1367_v110, 4)
			(_nw249).ArraySet1CodePoint(_1367_v110, 5)
			(_nw249).ArraySet1CodePoint((Companion_D0_.Create_DC1_(p2, (_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), _1367_v110, _this.F2())).Dtor_cf3(), 6)
			(_nw249).ArraySet1CodePoint(_1367_v110, 7)
			(_nw249).ArraySet1CodePoint(_1367_v110, 8)
			_1368_v111 = _nw249
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_1368_v111), 0))
			_ = _index253
			(_1368_v111).ArraySet1CodePoint(_dafny.CodePoint('b'), (_index253).Int())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_1368_v111), 0))
			_ = _index254
			(_1368_v111).ArraySet1CodePoint(_1367_v110, (_index254).Int())
		}
		r0 = (p0).Times((_1335_v80).Select((Companion_Default___.SafeIndex((_this).Fm1((_1232_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1232_v0), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_1335_v80).Cardinality()))).Uint32()).(_dafny.Int))
		return r0
	}
}
func (_this *C3) M5(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _1369_v0 _dafny.Array
		_ = _1369_v0
		var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
		_ = _nw250
		_1369_v0 = _nw250
		_1369_v0 = _1369_v0
		var _1370_v1 _dafny.Sequence
		_ = _1370_v1
		_1370_v1 = _dafny.SeqOf(_this.F2())
		var _hi6 _dafny.Int = (_1370_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1370_v1).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi6
		for _1371_i0 := p0; _1371_i0.Cmp(_hi6) < 0; _1371_i0 = _1371_i0.Plus(_dafny.One) {
			var _1372_v2 bool
			_ = _1372_v2
			_1372_v2 = true
			if (func() bool {
				if _1372_v2 {
					return _1372_v2
				}
				return (_1371_i0).Cmp(_this.F2()) <= 0
			})() {
				var _1373_v3 _dafny.Map
				_ = _1373_v3
				_1373_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v2, _1372_v2)
				var _1374_v4 _dafny.Sequence
				_ = _1374_v4
				_1374_v4 = _dafny.SeqOf(_1372_v2, !((func() bool {
					if (_1373_v3).Contains(_1372_v2) {
						return (_1373_v3).Get(_1372_v2).(bool)
					}
					return _1372_v2
				})()), _1372_v2)
				_1373_v3 = (_1373_v3).Update(_1372_v2, (_1374_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F2()), _dafny.IntOfUint32((_1374_v4).Cardinality()))).Uint32()).(bool))
				(_this).F2_set_(p0)
				var _1375_v5 *C0
				_ = _1375_v5
				var _nw251 *C0 = New_C0_()
				_ = _nw251
				_nw251.Ctor__(_1372_v2, _1372_v2)
				_1375_v5 = _nw251
				var _1376_v6 _dafny.Map
				_ = _1376_v6
				_1376_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1375_v5.F5, (_this).Fm1(!(_1375_v5.F4), globalState))
				(_this).F2_set_((_dafny.Zero).Minus((p0).Plus(((_1376_v6).Merge(_1376_v6)).Cardinality())))
				var _1377_v7 _dafny.Array
				_ = _1377_v7
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_29
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = (func(_1378_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1379_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1379_i1, _dafny.IntOfUint32((_1378_v4).Cardinality()))
						}
					})(_1374_v4)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw252 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw252).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw252).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1377_v7 = _nw252
				var _1380_v8 D4
				_ = _1380_v8
				_1380_v8 = Companion_D4_.Create_DC6_((_dafny.Zero).Minus((_this).Fm1(_1375_v5.F5, globalState)), _1377_v7, p0)
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1377_v7), 0))
				_ = _index255
				(_1377_v7).ArraySet1((_1380_v8).Dtor_cf9(), (_index255).Int())
				var _1381_v9 _dafny.Array
				_ = _1381_v9
				var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw253
				_1381_v9 = _nw253
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1381_v9), 0))
				_ = _index256
				(_1381_v9).ArraySet1(_1376_v6, (_index256).Int())
				var _1382_v10 _dafny.Map
				_ = _1382_v10
				_1382_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1371_i0, _1375_v5.F5)
				var _1383_v11 _dafny.Map
				_ = _1383_v11
				_1383_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _1371_i0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(505))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}(func(_1384_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})))
				var _1385_v12 _dafny.CodePoint
				_ = _1385_v12
				_1385_v12 = _dafny.CodePoint('g')
				var _1386_v13 D7
				_ = _1386_v13
				_1386_v13 = Companion_D7_.Create_DC15_(_this.F2(), _1385_v12)
				var _1387_v14 _dafny.Map
				_ = _1387_v14
				_1387_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1383_v11).Cardinality(), _1386_v13)
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1377_v7), 0))
				_ = _index257
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1381_v9), 0))
				_ = _index258
				var _rhs215 _dafny.Int = _this.F2()
				_ = _rhs215
				var _rhs216 _dafny.Int = (func() _dafny.Int {
					if _1375_v5.F5 {
						return ((_1382_v10).Cardinality()).Minus(_dafny.IntOfInt64(-727))
					}
					return ((_1387_v14).Merge(func() _dafny.Map {
						var _coll46 = _dafny.NewMapBuilder()
						_ = _coll46
						for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(183), _dafny.IntOfInt64(417))); ; {
							_compr_46, _ok50 := _iter50()
							if !_ok50 {
								break
							}
							var _1388_v15 _dafny.Int
							_1388_v15 = interface{}(_compr_46).(_dafny.Int)
							if ((_dafny.IntOfInt64(183)).Cmp(_1388_v15) <= 0) && ((_1388_v15).Cmp(_dafny.IntOfInt64(417)) < 0) {
								_coll46.Add((_1388_v15).Plus(p0), _1386_v13)
							}
						}
						return _coll46.ToMap()
					}())).Cardinality()
				})()
				_ = _rhs216
				var _rhs217 _dafny.Int = Companion_Default___.SafeModuloInt(p0, _this.F2())
				_ = _rhs217
				var _rhs218 _dafny.Map = ((_1376_v6).Merge(_1376_v6)).Merge(_1376_v6)
				_ = _rhs218
				var _rhs219 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1370_v1, _dafny.Companion_Sequence_.Concatenate(_1370_v1, _1370_v1))
				_ = _rhs219
				var _lhs127 *C3 = _this
				_ = _lhs127
				var _lhs128 _dafny.Array = _1377_v7
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1377_v7), 0))
				_ = _lhs129
				var _lhs130 _dafny.Array = _1381_v9
				_ = _lhs130
				var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1381_v9), 0))
				_ = _lhs131
				r0 = _rhs215
				_lhs127.F2_set_(_rhs216)
				(_lhs128).ArraySet1(_rhs217, (_lhs129).Int())
				(_lhs130).ArraySet1(_rhs218, (_lhs131).Int())
				_1370_v1 = _rhs219
			} else {
				var _1389_v16 _dafny.Array
				_ = _1389_v16
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw254
				_1389_v16 = _nw254
				var _1390_v17 _dafny.Array
				_ = _1390_v17
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw255
				_1390_v17 = _nw255
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_1389_v16), 0))
				_ = _index259
				(_1389_v16).ArraySet1(_1390_v17, (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_1389_v16), 0))
				_ = _index260
				(_1389_v16).ArraySet1(_1390_v17, (_index260).Int())
				var _1391_v18 *C1
				_ = _1391_v18
				var _nw256 *C1 = New_C1_()
				_ = _nw256
				_nw256.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-92))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}(func(_1392_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				})))
				_1391_v18 = _nw256
				_1391_v18 = _1391_v18
				var _1393_v19 _dafny.Set
				_ = _1393_v19
				_1393_v19 = _dafny.SetOf(_1371_i0)
				var _1394_v20 _dafny.Sequence
				_ = _1394_v20
				_1394_v20 = _dafny.SeqOf(_1393_v19)
				var _1395_v21 _dafny.Map
				_ = _1395_v21
				_1395_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_1394_v20, _1394_v20), _1372_v2)
				_1395_v21 = (_1395_v21).Update(!(_1372_v2), _1372_v2)
				var _1396_v22 _dafny.Sequence
				_ = _1396_v22
				_1396_v22 = _dafny.UnicodeSeqOfUtf8Bytes("hwqgaeod")
				_1396_v22 = _1396_v22
				var _1397_v23 _dafny.Sequence
				_ = _1397_v23
				_1397_v23 = _dafny.SeqOf(_1372_v2)
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1390_v17), 0))
				_ = _index261
				(_1390_v17).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(_1372_v2), _1372_v2), _1397_v23)).Cardinality()), (_index261).Int())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1390_v17), 0))
				_ = _index262
				(_1390_v17).ArraySet1((func() _dafny.Int {
					if _1372_v2 {
						return (_1371_i0).Times(_dafny.IntOfInt64(867))
					}
					return _this.F2()
				})(), (_index262).Int())
			}
			var _1398_v24 _dafny.CodePoint
			_ = _1398_v24
			_1398_v24 = _dafny.CodePoint('s')
			_1398_v24 = _1398_v24
			var _1399_v25 _dafny.Sequence
			_ = _1399_v25
			_1399_v25 = _dafny.UnicodeSeqOfUtf8Bytes("ednpivdn")
			var _1400_v26 _dafny.Sequence
			_ = _1400_v26
			_1400_v26 = _dafny.SeqOf(_1399_v25, _1399_v25)
			_1400_v26 = _1400_v26
			(_this).F2_set_(_1371_i0)
		}
		var _1401_v27 _dafny.Sequence
		_ = _1401_v27
		_1401_v27 = _dafny.UnicodeSeqOfUtf8Bytes("wosyavxm")
		var _1402_v28 D6
		_ = _1402_v28
		_1402_v28 = Companion_D6_.Create_DC11_(_1401_v27)
		r0 = (((_1370_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1370_v1).Cardinality()))).Uint32()).(_dafny.Int)).Minus(p0)).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1402_v28).Dtor_cf21()).Cardinality())))
		var _1403_v29 bool
		_ = _1403_v29
		_1403_v29 = true
		r0 = (func() _dafny.Int {
			if _1403_v29 {
				return (_this.F2()).Plus(_this.F2())
			}
			return _this.F2()
		})()
		var _1404_v30 _dafny.Set
		_ = _1404_v30
		_1404_v30 = _dafny.SetOf((_1370_v1).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1370_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(-226), p0)
		var _1405_v31 _dafny.Sequence
		_ = _1405_v31
		_1405_v31 = _dafny.SeqOf(_1403_v29)
		var _1406_v32 _dafny.MultiSet
		_ = _1406_v32
		_1406_v32 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1405_v31).Cardinality()), _this.F2())
		var _1407_v33 _dafny.Set
		_ = _1407_v33
		_1407_v33 = _dafny.SetOf(false, true)
		var _1408_v34 _dafny.Map
		_ = _1408_v34
		_1408_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1406_v32).Cardinality(), _1407_v33)
		_1404_v30 = ((_1404_v30).Difference(func() _dafny.Set {
			var _coll47 = _dafny.NewBuilder()
			_ = _coll47
			for _iter51 := _dafny.Iterate((Companion_Default___.Fm31(_1403_v29, (_1408_v34).Cardinality(), _1403_v29, (func() _dafny.Map {
				var _coll48 = _dafny.NewMapBuilder()
				_ = _coll48
				for _iter52 := _dafny.Iterate((func() _dafny.Map {
					var _coll49 = _dafny.NewMapBuilder()
					_ = _coll49
					for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(349), _dafny.IntOfInt64(-74))); ; {
						_compr_49, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1409_v36 _dafny.Int
						_1409_v36 = interface{}(_compr_49).(_dafny.Int)
						if ((_dafny.IntOfInt64(349)).Cmp(_1409_v36) <= 0) && ((_1409_v36).Cmp(_dafny.IntOfInt64(-74)) < 0) {
							_coll49.Add(Companion_Default___.SafeDivisionInt(_1409_v36, _this.F2()), (_1406_v32).Cardinality())
						}
					}
					return _coll49.ToMap()
				}()).Keys().Elements()); ; {
					_compr_48, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _1410_v35 _dafny.Int
					_1410_v35 = interface{}(_compr_48).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll50 = _dafny.NewMapBuilder()
						_ = _coll50
						for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(349), _dafny.IntOfInt64(-74))); ; {
							_compr_50, _ok54 := _iter54()
							if !_ok54 {
								break
							}
							var _1409_v36 _dafny.Int
							_1409_v36 = interface{}(_compr_50).(_dafny.Int)
							if ((_dafny.IntOfInt64(349)).Cmp(_1409_v36) <= 0) && ((_1409_v36).Cmp(_dafny.IntOfInt64(-74)) < 0) {
								_coll50.Add(Companion_Default___.SafeDivisionInt(_1409_v36, _this.F2()), (_1406_v32).Cardinality())
							}
						}
						return _coll50.ToMap()
					}()).Contains(_1410_v35) {
						_coll48.Add(Companion_Default___.SafeDivisionInt(_1410_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), (func() _dafny.Map {
							var _coll51 = _dafny.NewMapBuilder()
							_ = _coll51
							for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(443), _dafny.IntOfInt64(563))); ; {
								_compr_51, _ok55 := _iter55()
								if !_ok55 {
									break
								}
								var _1411_v37 _dafny.Int
								_1411_v37 = interface{}(_compr_51).(_dafny.Int)
								if ((_dafny.IntOfInt64(443)).Cmp(_1411_v37) <= 0) && ((_1411_v37).Cmp(_dafny.IntOfInt64(563)) < 0) {
									_coll51.Add((_1411_v37).Times(_this.F2()), p0)
								}
							}
							return _coll51.ToMap()
						}()).Cardinality())).Cardinality()), _1405_v31)
					}
				}
				return _coll48.ToMap()
			}()).Cardinality(), globalState)).Elements()); ; {
				_compr_47, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _1412_v38 _dafny.Int
				_1412_v38 = interface{}(_compr_47).(_dafny.Int)
				if (Companion_Default___.Fm31(_1403_v29, (_1408_v34).Cardinality(), _1403_v29, (func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter56 := _dafny.Iterate((func() _dafny.Map {
						var _coll53 = _dafny.NewMapBuilder()
						_ = _coll53
						for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(349), _dafny.IntOfInt64(-74))); ; {
							_compr_53, _ok57 := _iter57()
							if !_ok57 {
								break
							}
							var _1409_v36 _dafny.Int
							_1409_v36 = interface{}(_compr_53).(_dafny.Int)
							if ((_dafny.IntOfInt64(349)).Cmp(_1409_v36) <= 0) && ((_1409_v36).Cmp(_dafny.IntOfInt64(-74)) < 0) {
								_coll53.Add(Companion_Default___.SafeDivisionInt(_1409_v36, _this.F2()), (_1406_v32).Cardinality())
							}
						}
						return _coll53.ToMap()
					}()).Keys().Elements()); ; {
						_compr_52, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1410_v35 _dafny.Int
						_1410_v35 = interface{}(_compr_52).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll54 = _dafny.NewMapBuilder()
							_ = _coll54
							for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(349), _dafny.IntOfInt64(-74))); ; {
								_compr_54, _ok58 := _iter58()
								if !_ok58 {
									break
								}
								var _1409_v36 _dafny.Int
								_1409_v36 = interface{}(_compr_54).(_dafny.Int)
								if ((_dafny.IntOfInt64(349)).Cmp(_1409_v36) <= 0) && ((_1409_v36).Cmp(_dafny.IntOfInt64(-74)) < 0) {
									_coll54.Add(Companion_Default___.SafeDivisionInt(_1409_v36, _this.F2()), (_1406_v32).Cardinality())
								}
							}
							return _coll54.ToMap()
						}()).Contains(_1410_v35) {
							_coll52.Add(Companion_Default___.SafeDivisionInt(_1410_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), (func() _dafny.Map {
								var _coll55 = _dafny.NewMapBuilder()
								_ = _coll55
								for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(443), _dafny.IntOfInt64(563))); ; {
									_compr_55, _ok59 := _iter59()
									if !_ok59 {
										break
									}
									var _1411_v37 _dafny.Int
									_1411_v37 = interface{}(_compr_55).(_dafny.Int)
									if ((_dafny.IntOfInt64(443)).Cmp(_1411_v37) <= 0) && ((_1411_v37).Cmp(_dafny.IntOfInt64(563)) < 0) {
										_coll55.Add((_1411_v37).Times(_this.F2()), p0)
									}
								}
								return _coll55.ToMap()
							}()).Cardinality())).Cardinality()), _1405_v31)
						}
					}
					return _coll52.ToMap()
				}()).Cardinality(), globalState)).Contains(_1412_v38) {
					_coll47.Add((_1412_v38).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(656), (_dafny.SetOf(true, true)).Cardinality())).Cardinality())).Cardinality()))
				}
			}
			return _coll47.ToSet()
		}())).Difference((func() _dafny.Set {
			if _1403_v29 {
				return _dafny.SetOf(_dafny.IntOfUint32((_1401_v27).Cardinality()))
			}
			return _1404_v30
		})())
		_1408_v34 = (_1408_v34).Update(_this.F2(), _1407_v33)
		var _1413_v39 *C1
		_ = _1413_v39
		var _nw257 *C1 = New_C1_()
		_ = _nw257
		_nw257.Ctor__(_1401_v27)
		_1413_v39 = _nw257
		var _1414_v40 D8
		_ = _1414_v40
		_1414_v40 = Companion_D8_.Create_DC16_(_1413_v39)
		var _1415_v41 _dafny.Map
		_ = _1415_v41
		_1415_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1414_v40, p0)
		r0 = (func() _dafny.Int {
			if (_1415_v41).Contains(_1414_v40) {
				return (_1415_v41).Get(_1414_v40).(_dafny.Int)
			}
			return p0
		})()
		r1 = (((_dafny.MultiSetOf(_1403_v29, true)).Cardinality()).Times((_dafny.Zero).Minus(p0))).Cmp(_this.F2()) <= 0
		r2 = _1403_v29
		return r0, r1, r2
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f2 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f2 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_, Companion_T2_.TraitID_}
}

var _ T0 = &C4{}
var _ T1 = &C4{}
var _ T2 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F2() _dafny.Int {
	return _this._f2
}
func (_this *C4) F2_set_(value _dafny.Int) {
	_this._f2 = value
}
func (_this *C4) Ctor__(f2 _dafny.Int) {
	{
		(_this)._f2 = f2
	}
}
func (_this *C4) Fm0(globalState *GlobalState) _dafny.Map {
	{
		var _source20 D0 = Companion_D0_.Create_DC1_(false, true, _dafny.CodePoint('l'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mephw")).Cardinality()))
		_ = _source20
		if _source20.Is_DC1() {
			var _1416___mcc_h0 bool = _source20.Get_().(D0_DC1).Cf1
			_ = _1416___mcc_h0
			var _1417___mcc_h1 bool = _source20.Get_().(D0_DC1).Cf2
			_ = _1417___mcc_h1
			var _1418___mcc_h2 _dafny.CodePoint = _source20.Get_().(D0_DC1).Cf3
			_ = _1418___mcc_h2
			var _1419___mcc_h3 _dafny.Int = _source20.Get_().(D0_DC1).Cf4
			_ = _1419___mcc_h3
			var _1420_cf4 _dafny.Int = _1419___mcc_h3
			_ = _1420_cf4
			var _1421_cf3 _dafny.CodePoint = _1418___mcc_h2
			_ = _1421_cf3
			var _1422_cf2 bool = _1417___mcc_h1
			_ = _1422_cf2
			var _1423_cf1 bool = _1416___mcc_h0
			_ = _1423_cf1
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1420_cf4, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}((func(_1424_cf3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1425_i0 _dafny.Int) _dafny.CodePoint {
					return _1424_cf3
				}
			})(_1421_cf3)))).Cardinality()), _this.F2(), _1420_cf4, _1420_cf4), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(837))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}(func(_1426_i1 _dafny.Int) _dafny.Int {
				return _this.F2()
			}))).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1420_cf4, _1420_cf4), _this.F2())))
		} else {
			var _1427___mcc_h4 bool = _source20.Get_().(D0_DC0).Cf0
			_ = _1427___mcc_h4
			var _1428_cf0 bool = _1427___mcc_h4
			_ = _1428_cf0
			return func() _dafny.Map {
				var _coll56 = _dafny.NewMapBuilder()
				_ = _coll56
				for _iter60 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(530))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}(func(_1429_i2 _dafny.Int) _dafny.Int {
					return _this.F2()
				})), _dafny.SeqOf(_this.F2(), _this.F2(), (_dafny.MultiSetOf(_this.F2(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_dafny.Zero).Minus(_this.F2()))).Cardinality())).Cardinality()))).Elements()); ; {
					_compr_56, _ok60 := _iter60()
					if !_ok60 {
						break
					}
					var _1430_v0 _dafny.Sequence
					_1430_v0 = interface{}(_compr_56).(_dafny.Sequence)
					if (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(530))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg80 _dafny.Int) interface{} {
							return coer80(arg80)
						}
					}(func(_1429_i2 _dafny.Int) _dafny.Int {
						return _this.F2()
					})), _dafny.SeqOf(_this.F2(), _this.F2(), (_dafny.MultiSetOf(_this.F2(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_dafny.Zero).Minus(_this.F2()))).Cardinality())).Cardinality()))).Contains(_1430_v0) {
						_coll56.Add(_1430_v0, _dafny.IntOfInt64(802))
					}
				}
				return _coll56.ToMap()
			}()
		}
	}
}
func (_this *C4) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.SetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(929), _this.F2()))).Cardinality()
	}
}
func (_this *C4) Fm8(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality()))).Union(_dafny.SetOf(_this.F2(), _this.F2()))).IsDisjointFrom(_dafny.SetOf(_this.F2(), _this.F2()))
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1431_v0 _dafny.Map
		_ = _1431_v0
		_1431_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(p2))
		var _1432_v1 _dafny.Map
		_ = _1432_v1
		_1432_v1 = _1431_v0
		var _source21 _dafny.Map = _1432_v1
		_ = _source21
		var _1433___mcc_h0 _dafny.Map = _source21
		_ = _1433___mcc_h0
		var _1434_cf5 _dafny.Map = _1433___mcc_h0
		_ = _1434_cf5
		var _1435_v2 _dafny.CodePoint
		_ = _1435_v2
		_1435_v2 = _dafny.CodePoint('s')
		var _1436_v3 _dafny.Map
		_ = _1436_v3
		_1436_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1435_v2, (_this.F2()).Plus(_this.F2()))
		_1436_v3 = (_1436_v3).Update(_1435_v2, p0)
		var _1437_v4 _dafny.Set
		_ = _1437_v4
		_1437_v4 = _dafny.SetOf(p0, _dafny.IntOfUint32((Companion_Default___.Fm9(p2, p2, p2, _this.F2(), globalState)).Cardinality()))
		_1437_v4 = _1437_v4
		(_this).F2_set_(p0)
		var _1438_v5 *C2
		_ = _1438_v5
		var _nw258 *C2 = New_C2_()
		_ = _nw258
		_nw258.Ctor__(_this.F2())
		_1438_v5 = _nw258
		var _1439_v6 _dafny.Sequence
		_ = _1439_v6
		_1439_v6 = _dafny.UnicodeSeqOfUtf8Bytes("nuymyuim")
		var _1440_v7 _dafny.Map
		_ = _1440_v7
		_1440_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _1441_v8 D5
		_ = _1441_v8
		_1441_v8 = Companion_D5_.Create_DC10_(p3, _dafny.IntOfUint32((_1439_v6).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1439_v6, _1440_v7), p3)
		var _1442_v9 _dafny.Map
		_ = _1442_v9
		_1442_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1441_v8, _1441_v8), p3)
		var _1443_v10 T0
		_ = _1443_v10
		var _nw259 *C2 = New_C2_()
		_ = _nw259
		_nw259.Ctor__((_this).Fm1(p3, globalState))
		_1443_v10 = _nw259
		var _1444_v11 _dafny.MultiSet
		_ = _1444_v11
		_1444_v11 = _dafny.MultiSetOf(_1443_v10)
		var _1445_v12 _dafny.Array
		_ = _1445_v12
		var _nwElement0_63 bool = p3
		_ = _nwElement0_63
		var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(4))
		_ = _nw260
		(_nw260).ArraySet1(_nwElement0_63, 0)
		(_nw260).ArraySet1(((_1442_v9).Cardinality()).Cmp((_1444_v11).Cardinality()) != 0, 1)
		(_nw260).ArraySet1(p2, 2)
		(_nw260).ArraySet1(false, 3)
		_1445_v12 = _nw260
		_1445_v12 = _1445_v12
		var _1446_v13 *C2
		_ = _1446_v13
		var _nw261 *C2 = New_C2_()
		_ = _nw261
		_nw261.Ctor__(p0)
		_1446_v13 = _nw261
		var _1447_v14 _dafny.Array
		_ = _1447_v14
		var _nwElement0_64 _dafny.Int = p0
		_ = _nwElement0_64
		var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(6))
		_ = _nw262
		(_nw262).ArraySet1(_nwElement0_64, 0)
		(_nw262).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1439_v6, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1439_v6).Cardinality()))).Uint32(), (_1439_v6).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1439_v6).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()), 1)
		(_nw262).ArraySet1(_this.F2(), 2)
		(_nw262).ArraySet1((_dafny.Zero).Minus(_this.F2()), 3)
		(_nw262).ArraySet1(_this.F2(), 4)
		(_nw262).ArraySet1(_this.F2(), 5)
		_1447_v14 = _nw262
		var _1448_v15 _dafny.Sequence
		_ = _1448_v15
		_1448_v15 = _dafny.SeqOf(_this.F2(), p0)
		var _rhs220 _dafny.Array = _1447_v14
		_ = _rhs220
		var _rhs221 bool = p3
		_ = _rhs221
		var _rhs222 _dafny.Int = (p0).Plus((_1448_v15).Select((Companion_Default___.SafeIndex(_this.F2(), _dafny.IntOfUint32((_1448_v15).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _rhs222
		_1447_v14 = _rhs220
		r0 = _rhs221
		r2 = _rhs222
		var _1449_v16 _dafny.CodePoint
		_ = _1449_v16
		_1449_v16 = _dafny.CodePoint('r')
		if !(_dafny.Companion_Sequence_.Contains(_1439_v6, _1449_v16)) {
			var _1450_v17 *C1
			_ = _1450_v17
			var _nw263 *C1 = New_C1_()
			_ = _nw263
			_nw263.Ctor__(_1439_v6)
			_1450_v17 = _nw263
			var _pat_let_tv29 = _1450_v17
			_ = _pat_let_tv29
			var _source22 D8 = func(_pat_let28_0 D8) D8 {
				return func(_1451_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let29_0 *C1) D8 {
						return func(_1452_dt__update_hcf29_h0 *C1) D8 {
							return Companion_D8_.Create_DC16_(_1452_dt__update_hcf29_h0)
						}(_pat_let29_0)
					}(_pat_let_tv29)
				}(_pat_let28_0)
			}(Companion_D8_.Create_DC16_(_1450_v17))
			_ = _source22
			if _source22.Is_DC17() {
				var _1453___mcc_h1 bool = _source22.Get_().(D8_DC17).Cf30
				_ = _1453___mcc_h1
				var _1454___mcc_h2 *C1 = _source22.Get_().(D8_DC17).Cf31
				_ = _1454___mcc_h2
				var _1455___mcc_h3 _dafny.Int = _source22.Get_().(D8_DC17).Cf32
				_ = _1455___mcc_h3
				var _1456_cf32 _dafny.Int = _1455___mcc_h3
				_ = _1456_cf32
				var _1457_cf31 *C1 = _1454___mcc_h2
				_ = _1457_cf31
				var _1458_cf30 bool = _1453___mcc_h1
				_ = _1458_cf30
				_1456_cf32 = (p0).Times((_this.F2()).Times(_1456_cf32))
				var _1459_v18 *C3
				_ = _1459_v18
				var _nw264 *C3 = New_C3_()
				_ = _nw264
				_nw264.Ctor__(p0)
				_1459_v18 = _nw264
				(_this).F2_set_(Companion_Default___.SafeDivisionInt(p0, _this.F2()))
				var _1460_v19 *C3
				_ = _1460_v19
				var _nw265 *C3 = New_C3_()
				_ = _nw265
				_nw265.Ctor__(_1456_cf32)
				_1460_v19 = _nw265
			} else {
				var _1461___mcc_h4 *C1 = _source22.Get_().(D8_DC16).Cf29
				_ = _1461___mcc_h4
				var _1462_cf29 *C1 = _1461___mcc_h4
				_ = _1462_cf29
				var _1463_v20 _dafny.Map
				_ = _1463_v20
				_1463_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_1450_v17).Fm1(p2, globalState))
				var _rhs223 _dafny.Map = (_1463_v20).Merge((_1463_v20).Merge(_1463_v20))
				_ = _rhs223
				var _rhs224 _dafny.Array = _1445_v12
				_ = _rhs224
				var _rhs225 _dafny.Map = (_1440_v7).Merge(_1440_v7)
				_ = _rhs225
				var _rhs226 _dafny.Int = (_this.F2()).Minus(Companion_Default___.SafeDivisionInt(_this.F2(), _this.F2()))
				_ = _rhs226
				_1463_v20 = _rhs223
				_1445_v12 = _rhs224
				_1440_v7 = _rhs225
				r2 = _rhs226
				r2 = _this.F2()
				(_this).F2_set_(_this.F2())
				r2 = (p0).Plus(_this.F2())
			}
			(_this).F2_set_((_dafny.Zero).Minus(_this.F2()))
			var _1464_v21 D6
			_ = _1464_v21
			_1464_v21 = Companion_D6_.Create_DC12_(p2, p3)
			var _1465_v22 _dafny.Map
			_ = _1465_v22
			_1465_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1464_v21).Dtor_cf23(), _this.F2())
			var _rhs227 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1439_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}((func(_1466_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1467_i0 _dafny.Int) _dafny.CodePoint {
					return _1466_v16
				}
			})(_1449_v16))))
			_ = _rhs227
			var _rhs228 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F2(), (func() _dafny.Int {
				if (_1465_v22).Contains(p3) {
					return (_1465_v22).Get(p3).(_dafny.Int)
				}
				return _this.F2()
			})())
			_ = _rhs228
			r0 = _rhs227
			r2 = _rhs228
			var _1468_v23 _dafny.Array
			_ = _1468_v23
			var _nwElement0_65 _dafny.Array = _1447_v14
			_ = _nwElement0_65
			var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(10))
			_ = _nw266
			(_nw266).ArraySet1(_nwElement0_65, 0)
			(_nw266).ArraySet1((func() _dafny.Array {
				if p3 {
					return _1447_v14
				}
				return _1447_v14
			})(), 1)
			(_nw266).ArraySet1(_1447_v14, 2)
			(_nw266).ArraySet1(_1447_v14, 3)
			(_nw266).ArraySet1(_1447_v14, 4)
			(_nw266).ArraySet1((func() _dafny.Array {
				if p2 {
					return _1447_v14
				}
				return _1447_v14
			})(), 5)
			(_nw266).ArraySet1(_1447_v14, 6)
			(_nw266).ArraySet1(_1447_v14, 7)
			(_nw266).ArraySet1(_1447_v14, 8)
			(_nw266).ArraySet1((func() _dafny.Array {
				if p2 {
					return _1447_v14
				}
				return _1447_v14
			})(), 9)
			_1468_v23 = _nw266
			_1468_v23 = _1468_v23
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_1447_v14), 0))
			_ = _index263
			(_1447_v14).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_1448_v15).Cardinality())), (_index263).Int())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_1447_v14), 0))
			_ = _index264
			(_1447_v14).ArraySet1(_this.F2(), (_index264).Int())
		} else {
			var _1469_v24 _dafny.Map
			_ = _1469_v24
			_1469_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F2())), p0)
			_1469_v24 = (_1469_v24).Update(p0, (p0).Times(_dafny.IntOfInt64(-273)))
			(_this).F2_set_((_1446_v13).Fm1(p2, globalState))
			var _1470_v25 _dafny.Sequence
			_ = _1470_v25
			_1470_v25 = _dafny.SeqOf(p2)
			var _1471_v26 _dafny.Sequence
			_ = _1471_v26
			_1471_v26 = _dafny.SeqOf(_dafny.SeqOf(p3))
			var _1472_v27 _dafny.Sequence
			_ = _1472_v27
			_1472_v27 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1470_v25, (_1471_v26).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1471_v26).Cardinality()))).Uint32()).(_dafny.Sequence)), _1470_v25, _dafny.Companion_Sequence_.Concatenate(_1470_v25, _1470_v25), _1470_v25, _dafny.SeqOf(p3, p2, p3))
			_1471_v26 = _dafny.Companion_Sequence_.Concatenate(_1472_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(384))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg82 _dafny.Int) interface{} {
					return coer82(arg82)
				}
			}((func(_1473_v25 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1474_i1 _dafny.Int) _dafny.Sequence {
					return _1473_v25
				}
			})(_1470_v25))))
			var _1475_v28 _dafny.Set
			_ = _1475_v28
			_1475_v28 = _dafny.SetOf(p2)
			var _rhs229 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1448_v15, _dafny.SeqOf(_this.F2(), _this.F2()))
			_ = _rhs229
			var _rhs230 _dafny.Int = (func() _dafny.Int {
				if p2 {
					return ((_1475_v28).Intersection(_1475_v28)).Cardinality()
				}
				return _dafny.IntOfInt64(967)
			})()
			_ = _rhs230
			var _rhs231 bool = p2
			_ = _rhs231
			var _rhs232 _dafny.Int = _dafny.IntOfInt64(-227)
			_ = _rhs232
			_1448_v15 = _rhs229
			r2 = _rhs230
			r0 = _rhs231
			r2 = _rhs232
			var _1476_v29 _dafny.Map
			_ = _1476_v29
			_1476_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v14, p0)
			_1476_v29 = (_1476_v29).Update(_1447_v14, _this.F2())
		}
		var _1477_v30 _dafny.Sequence
		_ = _1477_v30
		_1477_v30 = _dafny.SeqOf(p2)
		var _1478_v31 _dafny.Set
		_ = _1478_v31
		_1478_v31 = _dafny.SetOf(p3, p2, p2, !(p2), p2)
		_1439_v6 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm28((_1477_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.IntOfUint32((_1477_v30).Cardinality()))).Uint32()).(bool), p3, p2, _1478_v31, globalState), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm28(true, !(p3), p2, Companion_Default___.Fm41(_1448_v15, p3, true, p2, globalState), globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm28(true, !(p3), p2, Companion_Default___.Fm41(_1448_v15, p3, true, p2, globalState), globalState)).Cardinality()))).Uint32(), _1449_v16))
		r0 = p3
		var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
		_ = _nw267
		r1 = _nw267
		r2 = _this.F2()
		return r0, r1, r2
	}
}
func (_this *C4) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1479_v0 _dafny.Sequence
		_ = _1479_v0
		_1479_v0 = _dafny.UnicodeSeqOfUtf8Bytes("uxrn")
		var _1480_v1 _dafny.Sequence
		_ = _1480_v1
		_1480_v1 = _dafny.SeqOf(p2)
		var _1481_v2 _dafny.MultiSet
		_ = _1481_v2
		_1481_v2 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-166)))
		var _1482_v3 _dafny.Sequence
		_ = _1482_v3
		_1482_v3 = _dafny.SeqOf(_1481_v2)
		var _1483_v4 _dafny.CodePoint
		_ = _1483_v4
		_1483_v4 = _dafny.CodePoint('j')
		var _1484_v5 _dafny.Array
		_ = _1484_v5
		var _nwElement0_66 bool = (p2) == (p2)
		_ = _nwElement0_66
		var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(25))
		_ = _nw268
		(_nw268).ArraySet1(_nwElement0_66, 0)
		(_nw268).ArraySet1(p2, 1)
		(_nw268).ArraySet1((p0).Cmp(p1) < 0, 2)
		(_nw268).ArraySet1(p2, 3)
		(_nw268).ArraySet1(!(p2), 4)
		(_nw268).ArraySet1(!(p2), 5)
		(_nw268).ArraySet1(p2, 6)
		(_nw268).ArraySet1((func() bool {
			if p2 {
				return p2
			}
			return p2
		})(), 7)
		(_nw268).ArraySet1(p2, 8)
		(_nw268).ArraySet1(p2, 9)
		(_nw268).ArraySet1(p2, 10)
		(_nw268).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1479_v0, _1479_v0), 11)
		(_nw268).ArraySet1(p2, 12)
		(_nw268).ArraySet1(p2, 13)
		(_nw268).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1480_v1, _1480_v1), 14)
		(_nw268).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_1481_v2), _1482_v3), 15)
		(_nw268).ArraySet1(p2, 16)
		(_nw268).ArraySet1(p2, 17)
		(_nw268).ArraySet1(p2, 18)
		(_nw268).ArraySet1(p2, 19)
		(_nw268).ArraySet1(p2, 20)
		(_nw268).ArraySet1(p2, 21)
		(_nw268).ArraySet1(_dafny.Companion_Sequence_.Contains(_1479_v0, _1483_v4), 22)
		(_nw268).ArraySet1(p2, 23)
		(_nw268).ArraySet1(p2, 24)
		_1484_v5 = _nw268
		var _1485_v6 _dafny.Set
		_ = _1485_v6
		_1485_v6 = _dafny.SetOf(p1, p1, p1)
		var _1486_v7 _dafny.Map
		_ = _1486_v7
		_1486_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1483_v4, _1485_v6)
		var _1487_v8 _dafny.Sequence
		_ = _1487_v8
		_1487_v8 = _dafny.SeqOf(_1485_v6)
		var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
		_ = _index265
		(_1484_v5).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((func() _dafny.Set {
			if (_1486_v7).Contains(_1483_v4) {
				return (_1486_v7).Get(_1483_v4).(_dafny.Set)
			}
			return _dafny.SetOf(_dafny.IntOfUint32((_1479_v0).Cardinality()), _dafny.IntOfUint32((_1480_v1).Cardinality()))
		})(), _dafny.SetOf(_this.F2())), _1487_v8), (_index265).Int())
		var _1488_v9 *C2
		_ = _1488_v9
		var _nw269 *C2 = New_C2_()
		_ = _nw269
		_nw269.Ctor__(p1)
		_1488_v9 = _nw269
		var _1489_v10 _dafny.Map
		_ = _1489_v10
		_1489_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1488_v9)
		var _1490_v11 _dafny.Sequence
		_ = _1490_v11
		_1490_v11 = _dafny.SeqOf(p0)
		var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
		_ = _index266
		(_1484_v5).ArraySet1((_this).Fm8(_1479_v0, !(!(((_1489_v10).Cardinality()).Cmp(_dafny.IntOfUint32((_1490_v11).Cardinality())) <= 0)), globalState), (_index266).Int())
		var _1491_v12 _dafny.Array
		_ = _1491_v12
		var _nw270 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw270
		_1491_v12 = _nw270
		for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1491_v12), 0))); ; {
			_guard_loop_4, _ok61 := _iter61()
			if !_ok61 {
				break
			}
			var _1492_i0 _dafny.Int
			_1492_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1492_i0).Sign() != -1) && ((_1492_i0).Cmp(_dafny.ArrayLen((_1491_v12), 0)) < 0)) {
				(_1491_v12).ArraySet1(Companion_Default___.SafeModuloInt(_1492_i0, p1), (_1492_i0).Int())
			}
		}
		r0 = p1
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qsihpmbty")).Cardinality())
		_ = _hi7
		for _1493_i1 := p0; _1493_i1.Cmp(_hi7) < 0; _1493_i1 = _1493_i1.Plus(_dafny.One) {
			var _1494_v13 _dafny.Map
			_ = _1494_v13
			_1494_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1488_v9, _1493_i1)
			_1494_v13 = (_1494_v13).Merge((_1494_v13).Merge(_1494_v13))
			var _1495_v14 *C1
			_ = _1495_v14
			var _nw271 *C1 = New_C1_()
			_ = _nw271
			_nw271.Ctor__(_1479_v0)
			_1495_v14 = _nw271
			var _1496_v15 D8
			_ = _1496_v15
			_1496_v15 = Companion_D8_.Create_DC17_(false, _1495_v14, _1493_i1)
			var _1497_v16 _dafny.Map
			_ = _1497_v16
			_1497_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), _1483_v4)
			var _pat_let_tv30 = _1495_v14
			_ = _pat_let_tv30
			var _pat_let_tv31 = _1497_v16
			_ = _pat_let_tv31
			var _pat_let_tv32 = p1
			_ = _pat_let_tv32
			var _source23 D8 = func(_pat_let30_0 D8) D8 {
				return func(_1498_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let31_0 *C1) D8 {
						return func(_1499_dt__update_hcf31_h0 *C1) D8 {
							return func(_pat_let32_0 bool) D8 {
								return func(_1500_dt__update_hcf30_h0 bool) D8 {
									return Companion_D8_.Create_DC17_(_1500_dt__update_hcf30_h0, _1499_dt__update_hcf31_h0, (_1498_dt__update__tmp_h0).Dtor_cf32())
								}(_pat_let32_0)
							}(((_pat_let_tv31).Cardinality()).Cmp(_pat_let_tv32) != 0)
						}(_pat_let31_0)
					}(_pat_let_tv30)
				}(_pat_let30_0)
			}(_1496_v15)
			_ = _source23
			if _source23.Is_DC17() {
				var _1501___mcc_h0 bool = _source23.Get_().(D8_DC17).Cf30
				_ = _1501___mcc_h0
				var _1502___mcc_h1 *C1 = _source23.Get_().(D8_DC17).Cf31
				_ = _1502___mcc_h1
				var _1503___mcc_h2 _dafny.Int = _source23.Get_().(D8_DC17).Cf32
				_ = _1503___mcc_h2
				var _1504_cf32 _dafny.Int = _1503___mcc_h2
				_ = _1504_cf32
				var _1505_cf31 *C1 = _1502___mcc_h1
				_ = _1505_cf31
				var _1506_cf30 bool = _1501___mcc_h0
				_ = _1506_cf30
				var _rhs233 _dafny.Int = _1504_cf32
				_ = _rhs233
				var _rhs234 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F2(), _dafny.IntOfInt64(401))
				_ = _rhs234
				r0 = _rhs233
				_1504_cf32 = _rhs234
				_1504_cf32 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mqghm")).Cardinality()), _dafny.IntOfInt64(104)), Companion_Default___.SafeModuloInt(_1504_cf32, _1493_i1)))
				var _1507_v17 D11
				_ = _1507_v17
				_1507_v17 = Companion_D11_.Create_DC20_(_1484_v5)
				var _1508_v18 _dafny.Sequence
				_ = _1508_v18
				_1508_v18 = _dafny.SeqOf(_1484_v5)
				var _1509_v19 _dafny.Array
				_ = _1509_v19
				var _nwElement0_67 _dafny.Array = _1484_v5
				_ = _nwElement0_67
				var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(22))
				_ = _nw272
				(_nw272).ArraySet1(_nwElement0_67, 0)
				(_nw272).ArraySet1((_1507_v17).Dtor_cf35(), 1)
				(_nw272).ArraySet1(_1484_v5, 2)
				(_nw272).ArraySet1((func() _dafny.Array {
					if _1506_cf30 {
						return _1484_v5
					}
					return _1484_v5
				})(), 3)
				(_nw272).ArraySet1(_1484_v5, 4)
				(_nw272).ArraySet1(_1484_v5, 5)
				(_nw272).ArraySet1(_1484_v5, 6)
				(_nw272).ArraySet1(_1484_v5, 7)
				(_nw272).ArraySet1(_1484_v5, 8)
				(_nw272).ArraySet1(_1484_v5, 9)
				(_nw272).ArraySet1(_1484_v5, 10)
				(_nw272).ArraySet1(_1484_v5, 11)
				(_nw272).ArraySet1(_1484_v5, 12)
				(_nw272).ArraySet1(_1484_v5, 13)
				(_nw272).ArraySet1(_1484_v5, 14)
				(_nw272).ArraySet1(_1484_v5, 15)
				(_nw272).ArraySet1(_1484_v5, 16)
				(_nw272).ArraySet1(_1484_v5, 17)
				(_nw272).ArraySet1(_1484_v5, 18)
				(_nw272).ArraySet1(_1484_v5, 19)
				(_nw272).ArraySet1((_1508_v18).Select((Companion_Default___.SafeIndex((_1481_v2).Cardinality(), _dafny.IntOfUint32((_1508_v18).Cardinality()))).Uint32()).(_dafny.Array), 20)
				(_nw272).ArraySet1(_1484_v5, 21)
				_1509_v19 = _nw272
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1509_v19), 0))
				_ = _index267
				(_1509_v19).ArraySet1(_1484_v5, (_index267).Int())
				var _1510_v20 _dafny.Set
				_ = _1510_v20
				_1510_v20 = _dafny.SetOf(_1506_cf30, (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool))
				var _1511_v21 _dafny.MultiSet
				_ = _1511_v21
				_1511_v21 = _dafny.MultiSetOf(_1510_v20)
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1509_v19), 0))
				_ = _index268
				var _nwElement0_68 bool = (func() bool {
					if p2 {
						return _1506_cf30
					}
					return p2
				})()
				_ = _nwElement0_68
				var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(11))
				_ = _nw273
				(_nw273).ArraySet1(_nwElement0_68, 0)
				(_nw273).ArraySet1(!((_1511_v21).IsProperSubsetOf(_1511_v21)), 1)
				(_nw273).ArraySet1((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), 2)
				(_nw273).ArraySet1(_1506_cf30, 3)
				(_nw273).ArraySet1((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), 4)
				(_nw273).ArraySet1(!((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool)), 5)
				(_nw273).ArraySet1(true, 6)
				(_nw273).ArraySet1(_1506_cf30, 7)
				(_nw273).ArraySet1((_1481_v2).IsSubsetOf(_1481_v2), 8)
				(_nw273).ArraySet1(false, 9)
				(_nw273).ArraySet1(p2, 10)
				(_1509_v19).ArraySet1(_nw273, (_index268).Int())
				_1506_cf30 = ((func() _dafny.Int {
					if (_1481_v2).Contains(p1) {
						return (_1481_v2).Multiplicity(p1)
					}
					return _1493_i1
				})()).Cmp(p1) != 0
			} else {
				var _1512___mcc_h3 *C1 = _source23.Get_().(D8_DC16).Cf29
				_ = _1512___mcc_h3
				var _1513_cf29 *C1 = _1512___mcc_h3
				_ = _1513_cf29
				var _1514_v22 D6
				_ = _1514_v22
				_1514_v22 = Companion_D6_.Create_DC13_((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), _dafny.IntOfInt64(763))
				var _1515_v23 _dafny.Map
				_ = _1515_v23
				_1515_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1514_v22, _1484_v5)
				_1484_v5 = (func() _dafny.Array {
					if (_1515_v23).Contains(_1514_v22) {
						return (_1515_v23).Get(_1514_v22).(_dafny.Array)
					}
					return _1484_v5
				})()
				r0 = p0
				_1485_v6 = func() _dafny.Set {
					var _coll57 = _dafny.NewBuilder()
					_ = _coll57
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(96), _dafny.IntOfInt64(-363))); ; {
						_compr_57, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1516_v24 _dafny.Int
						_1516_v24 = interface{}(_compr_57).(_dafny.Int)
						if ((_dafny.IntOfInt64(96)).Cmp(_1516_v24) <= 0) && ((_1516_v24).Cmp(_dafny.IntOfInt64(-363)) < 0) {
							_coll57.Add((_1516_v24).Plus(_1493_i1))
						}
					}
					return _coll57.ToSet()
				}()
				(_this).F2_set_(p1)
			}
			_1479_v0 = _1479_v0
			if (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool) {
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
				_ = _index269
				(_1484_v5).ArraySet1(p2, (_index269).Int())
				var _1517_v25 *C1
				_ = _1517_v25
				var _nw274 *C1 = New_C1_()
				_ = _nw274
				_nw274.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(925))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1518_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1519_i2 _dafny.Int) _dafny.CodePoint {
						return _1518_v4
					}
				})(_1483_v4))))
				_1517_v25 = _nw274
				var _1520_v26 _dafny.Sequence
				_ = _1520_v26
				_1520_v26 = _dafny.SeqOf(_1484_v5)
				_1520_v26 = _1520_v26
				var _1521_v27 _dafny.Map
				_ = _1521_v27
				_1521_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(542))
				var _1522_v28 _dafny.Map
				_ = _1522_v28
				_1522_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1521_v27, (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool))
				var _1523_v29 _dafny.Map
				_ = _1523_v29
				_1523_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1491_v12)
				var _1524_v30 _dafny.Array
				_ = _1524_v30
				var _nwElement0_69 _dafny.Sequence = _1490_v11
				_ = _nwElement0_69
				var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(24))
				_ = _nw275
				(_nw275).ArraySet1(_nwElement0_69, 0)
				(_nw275).ArraySet1(_1490_v11, 1)
				(_nw275).ArraySet1(_1490_v11, 2)
				(_nw275).ArraySet1(_1490_v11, 3)
				(_nw275).ArraySet1(_1490_v11, 4)
				(_nw275).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, _dafny.IntOfUint32((_1480_v1).Cardinality()), p0), (Companion_Default___.SafeIndex((_1522_v28).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(p0, _dafny.IntOfUint32((_1480_v1).Cardinality()), p0)).Cardinality()))).Uint32(), _dafny.IntOfInt64(-996)), 5)
				(_nw275).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(715)), 6)
				(_nw275).ArraySet1((_1495_v14).Fm12((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), (_1523_v29).Cardinality(), (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), globalState), 7)
				(_nw275).ArraySet1(_1490_v11, 8)
				(_nw275).ArraySet1(_1490_v11, 9)
				(_nw275).ArraySet1(_1490_v11, 10)
				(_nw275).ArraySet1(_dafny.SeqOf(_1493_i1), 11)
				(_nw275).ArraySet1(_1490_v11, 12)
				(_nw275).ArraySet1(_dafny.SeqOf(p1, _dafny.IntOfUint32(((_1495_v14).F3()).Cardinality())), 13)
				(_nw275).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}(func(_1525_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(451)
				})), 14)
				(_nw275).ArraySet1(_1490_v11, 15)
				(_nw275).ArraySet1(_1490_v11, 16)
				(_nw275).ArraySet1(_1490_v11, 17)
				(_nw275).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1526_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1527_i4 _dafny.Int) _dafny.Int {
						return _1526_p0
					}
				})(p0))), 18)
				(_nw275).ArraySet1(_1490_v11, 19)
				(_nw275).ArraySet1(_1490_v11, 20)
				(_nw275).ArraySet1(_1490_v11, 21)
				(_nw275).ArraySet1(_1490_v11, 22)
				(_nw275).ArraySet1(_1490_v11, 23)
				_1524_v30 = _nw275
				var _1528_v31 bool
				_ = _1528_v31
				var _1529_v32 _dafny.Array
				_ = _1529_v32
				var _1530_v33 _dafny.Int
				_ = _1530_v33
				var _out23 bool
				_ = _out23
				var _out24 _dafny.Array
				_ = _out24
				var _out25 _dafny.Int
				_ = _out25
				_out23, _out24, _out25 = (_1517_v25).M1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ck")).Cardinality()), _1524_v30, !(!((_1493_i1).Cmp(_this.F2()) > 0)), false, globalState)
				_1528_v31 = _out23
				_1529_v32 = _out24
				_1530_v33 = _out25
				r0 = _1493_i1
			} else {
				var _1531_v34 _dafny.Set
				_ = _1531_v34
				_1531_v34 = _dafny.SetOf((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), p2)
				var _1532_v35 D4
				_ = _1532_v35
				_1532_v35 = Companion_D4_.Create_DC5_((func() _dafny.Set {
					if true {
						return _1531_v34
					}
					return _1531_v34
				})())
				var _1533_v38 _dafny.Map
				_ = _1533_v38
				_1533_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll58 = _dafny.NewMapBuilder()
					_ = _coll58
					for _iter63 := _dafny.Iterate((_1485_v6).Elements()); ; {
						_compr_58, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1534_v37 _dafny.Int
						_1534_v37 = interface{}(_compr_58).(_dafny.Int)
						if (_1485_v6).Contains(_1534_v37) {
							_coll58.Add((_1534_v37).Plus(_dafny.IntOfInt64(445)), (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool))
						}
					}
					return _coll58.ToMap()
				}(), _this.F2())
				var _1535_v39 _dafny.Map
				_ = _1535_v39
				_1535_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool))
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
				_ = _index270
				var _rhs235 bool = (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool)
				_ = _rhs235
				var _rhs236 _dafny.Int = (func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter64 := _dafny.Iterate((((_1533_v38).Update(_1535_v39, p1)).Merge(_1533_v38)).Keys().Elements()); ; {
						_compr_59, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _1536_v36 _dafny.Map
						_1536_v36 = interface{}(_compr_59).(_dafny.Map)
						if (((_1533_v38).Update(_1535_v39, p1)).Merge(_1533_v38)).Contains(_1536_v36) {
							_coll59.Add(_1536_v36, Companion_Default___.SafeDivisionInt(_this.F2(), p0))
						}
					}
					return _coll59.ToMap()
				}()).Cardinality()
				_ = _rhs236
				var _rhs237 D4 = _1532_v35
				_ = _rhs237
				var _rhs238 _dafny.CodePoint = _1483_v4
				_ = _rhs238
				var _lhs132 _dafny.Array = _1484_v5
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
				_ = _lhs133
				(_lhs132).ArraySet1(_rhs235, (_lhs133).Int())
				r0 = _rhs236
				_1532_v35 = _rhs237
				_1483_v4 = _rhs238
				_1490_v11 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1490_v11, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1490_v11).Cardinality()))).Uint32(), p1), _1490_v11)
				(_this).F2_set_(((_1488_v9).Fm1((_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), globalState)).Times(p1))
				r0 = _this.F2()
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
				_ = _index271
				(_1484_v5).ArraySet1(p2, (_index271).Int())
			}
		}
		var _1537_v40 _dafny.MultiSet
		_ = _1537_v40
		_1537_v40 = _dafny.MultiSetOf(p2, p2)
		_1537_v40 = (_1537_v40).Union(_1537_v40)
		var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))
		_ = _index272
		(_1484_v5).ArraySet1((_this).Fm8(_dafny.Companion_Sequence_.Concatenate(_1479_v0, _1479_v0), (_1484_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1484_v5), 0))).Int()).(bool), globalState), (_index272).Int())
		var _1538_v41 _dafny.Set
		_ = _1538_v41
		_1538_v41 = _dafny.SetOf(_1484_v5, _1484_v5)
		r0 = Companion_Default___.SafeDivisionInt(_this.F2(), Companion_Default___.SafeModuloInt((_1538_v41).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F2())).Cardinality())))
		return r0
	}
}
func (_this *C4) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1539_v0 _dafny.Map
		_ = _1539_v0
		_1539_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(p3))
		_1539_v0 = (_1539_v0).Update(p3, p3)
		var _1540_i0 _dafny.Int
		_ = _1540_i0
		_1540_i0 = _dafny.Zero
		{
			for p3 {
				{
					if (_1540_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1540_i0 = (_1540_i0).Plus(_dafny.One)
					var _1541_v1 _dafny.CodePoint
					_ = _1541_v1
					_1541_v1 = _dafny.CodePoint('q')
					var _1542_v2 _dafny.Sequence
					_ = _1542_v2
					_1542_v2 = _dafny.SeqOf((_dafny.IntOfInt64(391)).Minus(_dafny.IntOfInt64(438)), (_this.F2()).Times(_this.F2()), _dafny.IntOfInt64(-19), (_this.F2()).Plus(_this.F2()))
					var _1543_v3 _dafny.Map
					_ = _1543_v3
					_1543_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1541_v1)
					var _1544_v4 _dafny.Sequence
					_ = _1544_v4
					_1544_v4 = _dafny.UnicodeSeqOfUtf8Bytes("kwftu")
					var _rhs239 _dafny.CodePoint = (func() _dafny.CodePoint {
						if (_1543_v3).Contains((_this).Fm8(_1544_v4, true, globalState)) {
							return (_1543_v3).Get((_this).Fm8(_1544_v4, true, globalState)).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('k')
					})()
					_ = _rhs239
					var _rhs240 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1542_v2, _1542_v2)
					_ = _rhs240
					var _rhs241 _dafny.Int = (_this).Fm1(_dafny.Companion_Sequence_.IsPrefixOf(_1544_v4, _1544_v4), globalState)
					_ = _rhs241
					_1541_v1 = _rhs239
					_1542_v2 = _rhs240
					r0 = _rhs241
					var _1545_v5 bool
					_ = _1545_v5
					_1545_v5 = true
					_1545_v5 = p3
					var _1546_v6 _dafny.Array
					_ = _1546_v6
					var _out26 _dafny.Array
					_ = _out26
					_out26 = Companion_Default___.M0(p2, _1545_v5, globalState)
					_1546_v6 = _out26
					_1542_v2 = _1542_v2
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if p3 {
			var _1547_v7 bool
			_ = _1547_v7
			_1547_v7 = true
			var _1548_v8 _dafny.Sequence
			_ = _1548_v8
			_1548_v8 = _dafny.UnicodeSeqOfUtf8Bytes("yauevfig")
			_1547_v7 = (_this).Fm8(_1548_v8, (func() bool {
				if _1547_v7 {
					return p3
				}
				return p2
			})(), globalState)
			r1 = Companion_Default___.SafeModuloInt(_this.F2(), _dafny.IntOfUint32((_1548_v8).Cardinality()))
			var _1549_v9 _dafny.CodePoint
			_ = _1549_v9
			_1549_v9 = _dafny.CodePoint('a')
			var _1550_v10 _dafny.Array
			_ = _1550_v10
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_30
			var _nw276 _dafny.Array
			_ = _nw276
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw276 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) D7 = func(_1551_i1 _dafny.Int) D7 {
					return Companion_D7_.Create_DC14_(_dafny.SeqOf(_this.F2(), _this.F2(), (_dafny.SetOf(_this.F2(), _this.F2())).Cardinality(), (_dafny.Zero).Minus(_this.F2()), _this.F2()))
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw276 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw276).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw276).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1550_v10 = _nw276
			var _1552_v11 _dafny.Map
			_ = _1552_v11
			_1552_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1549_v9, _1550_v10)
			var _rhs242 _dafny.Map = (_1552_v11).Merge(_1552_v11)
			_ = _rhs242
			var _rhs243 _dafny.Int = _this.F2()
			_ = _rhs243
			var _rhs244 bool = p2
			_ = _rhs244
			var _rhs245 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F2()))
			_ = _rhs245
			_1552_v11 = _rhs242
			r0 = _rhs243
			_1547_v7 = _rhs244
			r1 = _rhs245
			var _1553_v12 _dafny.Array
			_ = _1553_v12
			var _nw277 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw277
			_1553_v12 = _nw277
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1553_v12), 0))
			_ = _index273
			(_1553_v12).ArraySet1(p2, (_index273).Int())
			var _1554_v13 T0
			_ = _1554_v13
			var _nw278 *C2 = New_C2_()
			_ = _nw278
			_nw278.Ctor__(_dafny.IntOfInt64(693))
			_1554_v13 = _nw278
			var _1555_v14 _dafny.MultiSet
			_ = _1555_v14
			_1555_v14 = _dafny.MultiSetOf(_1554_v13)
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1553_v12), 0))
			_ = _index274
			(_1553_v12).ArraySet1((_1555_v14).Equals(_1555_v14), (_index274).Int())
			var _1556_v15 *C1
			_ = _1556_v15
			var _nw279 *C1 = New_C1_()
			_ = _nw279
			_nw279.Ctor__((func() _dafny.Sequence {
				if p2 {
					return _1548_v8
				}
				return _1548_v8
			})())
			_1556_v15 = _nw279
		} else {
			var _1557_v16 _dafny.Array
			_ = _1557_v16
			var _nw280 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw280
			_1557_v16 = _nw280
			_1557_v16 = (func() _dafny.Array {
				if !(true) || (p2) {
					return _1557_v16
				}
				return _1557_v16
			})()
			var _1558_v17 _dafny.MultiSet
			_ = _1558_v17
			_1558_v17 = _dafny.MultiSetOf(p1)
			var _1559_v18 D5
			_ = _1559_v18
			_1559_v18 = Companion_D5_.Create_DC8_(_1558_v17)
			var _1560_v19 bool
			_ = _1560_v19
			_1560_v19 = true
			var _1561_v20 _dafny.Sequence
			_ = _1561_v20
			_1561_v20 = _dafny.SeqOf(_1560_v19)
			var _1562_v21 *C3
			_ = _1562_v21
			var _nw281 *C3 = New_C3_()
			_ = _nw281
			_nw281.Ctor__(_this.F2())
			_1562_v21 = _nw281
			var _1563_v22 _dafny.Map
			_ = _1563_v22
			_1563_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1561_v20, _1562_v21)
			var _1564_v23 _dafny.Map
			_ = _1564_v23
			_1564_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _dafny.IntOfInt64(479))
			var _1565_v24 _dafny.MultiSet
			_ = _1565_v24
			_1565_v24 = _dafny.MultiSetOf(_this.F2(), Companion_Default___.SafeModuloInt(_this.F2(), (func() _dafny.Int {
				if (_1564_v23).Contains(_dafny.IntOfInt64(61)) {
					return (_1564_v23).Get(_dafny.IntOfInt64(61)).(_dafny.Int)
				}
				return _this.F2()
			})()))
			var _1566_v25 D0
			_ = _1566_v25
			_1566_v25 = Companion_D0_.Create_DC0_(p2)
			var _1567_v26 _dafny.Sequence
			_ = _1567_v26
			_1567_v26 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("hweyh"))
			var _1568_v27 _dafny.Sequence
			_ = _1568_v27
			_1568_v27 = _dafny.UnicodeSeqOfUtf8Bytes("ljkq")
			var _1569_v28 _dafny.Sequence
			_ = _1569_v28
			_1569_v28 = _dafny.SeqOf(_this.F2())
			var _rhs246 D5 = _1559_v18
			_ = _rhs246
			var _rhs247 bool = (_1566_v25).Dtor_cf0()
			_ = _rhs247
			var _rhs248 _dafny.Map = _1563_v22
			_ = _rhs248
			var _rhs249 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1567_v26, _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-522))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}(func(_1570_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), _1568_v27))).Cardinality())
			_ = _rhs249
			var _rhs250 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_1569_v28)).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg87 _dafny.Int) interface{} {
					return coer87(arg87)
				}
			}(func(_1571_i3 _dafny.Int) _dafny.Int {
				return _this.F2()
			}))))
			_ = _rhs250
			_1559_v18 = _rhs246
			_1560_v19 = _rhs247
			_1563_v22 = _rhs248
			r1 = _rhs249
			_1565_v24 = _rhs250
			var _1572_v29 *C3
			_ = _1572_v29
			var _nw282 *C3 = New_C3_()
			_ = _nw282
			_nw282.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-329), (_dafny.Zero).Minus(_this.F2())))
			_1572_v29 = _nw282
			var _1573_v30 _dafny.Array
			_ = _1573_v30
			var _nwElement0_70 _dafny.Sequence = _1568_v27
			_ = _nwElement0_70
			var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(11))
			_ = _nw283
			(_nw283).ArraySet1(_nwElement0_70, 0)
			(_nw283).ArraySet1(_1568_v27, 1)
			(_nw283).ArraySet1(_1568_v27, 2)
			(_nw283).ArraySet1(_1568_v27, 3)
			(_nw283).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-663))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}(func(_1574_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _1568_v27), 4)
			(_nw283).ArraySet1(_1568_v27, 5)
			(_nw283).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gjfrs"), _1568_v27), 6)
			(_nw283).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1568_v27, _1568_v27), 7)
			(_nw283).ArraySet1(_1568_v27, 8)
			(_nw283).ArraySet1(_1568_v27, 9)
			(_nw283).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("boyqhdgay"), 10)
			_1573_v30 = _nw283
			var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1573_v30), 0))
			_ = _index275
			(_1573_v30).ArraySet1(_1568_v27, (_index275).Int())
			var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1573_v30), 0))
			_ = _index276
			(_1573_v30).ArraySet1(_1568_v27, (_index276).Int())
			_1561_v20 = _dafny.Companion_Sequence_.Concatenate(_1561_v20, _dafny.Companion_Sequence_.Concatenate(_1561_v20, _1561_v20))
		}
		var _hi8 _dafny.Int = _this.F2()
		_ = _hi8
		for _1575_i5 := (func() _dafny.Int {
			if p3 {
				return _this.F2()
			}
			return _this.F2()
		})(); _1575_i5.Cmp(_hi8) < 0; _1575_i5 = _1575_i5.Plus(_dafny.One) {
			var _1576_v31 _dafny.Sequence
			_ = _1576_v31
			_1576_v31 = _dafny.UnicodeSeqOfUtf8Bytes("trytchnh")
			var _1577_v32 *C1
			_ = _1577_v32
			var _nw284 *C1 = New_C1_()
			_ = _nw284
			_nw284.Ctor__(_1576_v31)
			_1577_v32 = _nw284
			var _1578_v33 D8
			_ = _1578_v33
			_1578_v33 = Companion_D8_.Create_DC16_(_1577_v32)
			_1578_v33 = _1578_v33
			r1 = _1575_i5
			var _1579_v34 _dafny.Array
			_ = _1579_v34
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_31
			var _nw285 _dafny.Array
			_ = _nw285
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw285 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) bool = (func(_1580_p2 bool) func(_dafny.Int) bool {
					return func(_1581_i6 _dafny.Int) bool {
						return _1580_p2
					}
				})(p2)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw285 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw285).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw285).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1579_v34 = _nw285
			var _1582_v35 _dafny.Map
			_ = _1582_v35
			_1582_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1579_v34)
			_1582_v35 = _1582_v35
			var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1579_v34), 0))
			_ = _index277
			(_1579_v34).ArraySet1((_this).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(507))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg89 _dafny.Int) interface{} {
					return coer89(arg89)
				}
			}(func(_1583_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})), !(p2), globalState), (_index277).Int())
			var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1579_v34), 0))
			_ = _index278
			(_1579_v34).ArraySet1(p3, (_index278).Int())
		}
		var _1584_v36 _dafny.Array
		_ = _1584_v36
		var _nw286 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw286
		_1584_v36 = _nw286
		_1584_v36 = _1584_v36
		var _1585_v37 *C0
		_ = _1585_v37
		var _nw287 *C0 = New_C0_()
		_ = _nw287
		_nw287.Ctor__(p3, !(p2))
		_1585_v37 = _nw287
		r0 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(837), _this.F2())
		r1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F2(), (_dafny.Zero).Minus(_this.F2())))
		return r0, r1
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	F1  _dafny.Array
	_f0 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this.F1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f0 _dafny.Int, f1 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
	}
}
func (_this *C5) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}(func(_1586_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality())
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg91 _dafny.Int) interface{} {
				return coer91(arg91)
			}
		}(func(_1587_i1 _dafny.Int) _dafny.Int {
			return (_this).F0()
		}))), (_dafny.Zero).Minus((_this).F0()))
	}
}
func (_this *C5) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(426)
	}
}
func (_this *C5) Fm6(p0 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('m')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), _dafny.CodePoint('m'))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('b')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('e')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('c')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('l')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('j')))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.CodePoint('i'))))
	}
}
func (_this *C5) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		r0 = p2
		var _1588_v0 _dafny.Sequence
		_ = _1588_v0
		_1588_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gtifm")
		var _hi9 _dafny.Int = (_this).F0()
		_ = _hi9
		for _1589_i0 := _dafny.IntOfUint32((_1588_v0).Cardinality()); _1589_i0.Cmp(_hi9) < 0; _1589_i0 = _1589_i0.Plus(_dafny.One) {
			if p2 {
				var _1590_v1 _dafny.Set
				_ = _1590_v1
				_1590_v1 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1589_i0, p3)).Cardinality())
				var _1591_v2 _dafny.Map
				_ = _1591_v2
				_1591_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
				var _1592_v3 _dafny.Map
				_ = _1592_v3
				_1592_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_1591_v2).Cardinality())).IsSubsetOf(_1590_v1), (_this).F0())
				var _rhs251 _dafny.Int = _dafny.IntOfInt64(-735)
				_ = _rhs251
				var _rhs252 _dafny.Int = p0
				_ = _rhs252
				var _rhs253 _dafny.Int = (_dafny.Zero).Minus((_this).F0())
				_ = _rhs253
				var _rhs254 _dafny.Map = _1592_v3
				_ = _rhs254
				var _rhs255 bool = ((_this).F0()).Cmp(p0) != 0
				_ = _rhs255
				r2 = _rhs251
				r2 = _rhs252
				r2 = _rhs253
				_1592_v3 = _rhs254
				r0 = _rhs255
				var _arr2 _dafny.Array = _this.F1
				_ = _arr2
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_this.F1), 0))
				_ = _index279
				(_arr2).ArraySet1(p2, (_index279).Int())
				var _arr3 _dafny.Array = _this.F1
				_ = _arr3
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_this.F1), 0))
				_ = _index280
				(_arr3).ArraySet1(((_1591_v2).Cardinality()).Cmp(_1589_i0) > 0, (_index280).Int())
				var _1593_v4 _dafny.Map
				_ = _1593_v4
				_1593_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1591_v2)
				var _1594_v5 _dafny.Map
				_ = _1594_v5
				_1594_v5 = _1591_v2
				_1593_v4 = (_1593_v4).Update(p2, (_1594_v5))
				var _arr4 _dafny.Array = _this.F1
				_ = _arr4
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_this.F1), 0))
				_ = _index281
				(_arr4).ArraySet1(!(true), (_index281).Int())
				var _1595_v6 *C2
				_ = _1595_v6
				var _nw288 *C2 = New_C2_()
				_ = _nw288
				_nw288.Ctor__(_1589_i0)
				_1595_v6 = _nw288
			} else {
				var _1596_v7 _dafny.Map
				_ = _1596_v7
				_1596_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F1)
				_1596_v7 = (_1596_v7).Update(p2, _this.F1)
				var _arr5 _dafny.Array = _this.F1
				_ = _arr5
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))
				_ = _index282
				(_arr5).ArraySet1(!(p3) || (p2), (_index282).Int())
				var _1597_v8 _dafny.MultiSet
				_ = _1597_v8
				_1597_v8 = _dafny.MultiSetOf(_this.F1, _this.F1)
				var _arr6 _dafny.Array = _this.F1
				_ = _arr6
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))
				_ = _index283
				var _rhs256 bool = ((_1597_v8).Union(_1597_v8)).IsDisjointFrom((_1597_v8).Intersection(_1597_v8))
				_ = _rhs256
				var _rhs257 _dafny.Sequence = _1588_v0
				_ = _rhs257
				var _lhs134 _dafny.Array = _this.F1
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))
				_ = _lhs135
				(_lhs134).ArraySet1(_rhs256, (_lhs135).Int())
				_1588_v0 = _rhs257
				r0 = (_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool)
				var _1598_v9 *C0
				_ = _1598_v9
				var _nw289 *C0 = New_C0_()
				_ = _nw289
				_nw289.Ctor__((_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool), p3)
				_1598_v9 = _nw289
				var _1599_v10 _dafny.Array
				_ = _1599_v10
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_32
				var _nw290 _dafny.Array
				_ = _nw290
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw290 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Int = (func(_1600_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1601_i1 _dafny.Int) _dafny.Int {
							return (_1601_i1).Times(_1600_i0)
						}
					})(_1589_i0)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw290 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw290).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw290).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1599_v10 = _nw290
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_1599_v10), 0))
				_ = _index284
				(_1599_v10).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-535), _1589_i0), (_index284).Int())
				var _1602_v11 _dafny.Set
				_ = _1602_v11
				_1602_v11 = _dafny.SetOf(_1598_v9.F5)
				var _1603_v12 _dafny.MultiSet
				_ = _1603_v12
				_1603_v12 = _dafny.MultiSetOf(_1602_v11)
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_1599_v10), 0))
				_ = _index285
				(_1599_v10).ArraySet1(((_dafny.MultiSetOf(_dafny.SetOf(_1598_v9.F5, (_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool)))).Difference(_1603_v12)).Cardinality(), (_index285).Int())
			}
			var _1604_v13 _dafny.CodePoint
			_ = _1604_v13
			_1604_v13 = _dafny.CodePoint('j')
			var _1605_v14 _dafny.Array
			_ = _1605_v14
			var _nw291 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw291
			_1605_v14 = _nw291
			var _1606_v15 _dafny.Map
			_ = _1606_v15
			_1606_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
			var _1607_v16 D6
			_ = _1607_v16
			_1607_v16 = Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-155))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}((func(_1608_v0 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_1609_i2 _dafny.Int) _dafny.CodePoint {
					return (_1608_v0).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1608_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1588_v0))))
			var _pat_let_tv33 = _1588_v0
			_ = _pat_let_tv33
			var _rhs258 _dafny.CodePoint = _dafny.CodePoint('u')
			_ = _rhs258
			var _rhs259 _dafny.Array = _1605_v14
			_ = _rhs259
			var _rhs260 bool = !(p3) || (!(p2))
			_ = _rhs260
			var _rhs261 _dafny.Map = _1606_v15
			_ = _rhs261
			var _rhs262 D6 = (func() D6 {
				if p2 {
					return Companion_D6_.Create_DC11_(_1588_v0)
				}
				return func(_pat_let33_0 D6) D6 {
					return func(_1610_dt__update__tmp_h0 D6) D6 {
						return func(_pat_let34_0 _dafny.Sequence) D6 {
							return func(_1611_dt__update_hcf21_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC11_(_1611_dt__update_hcf21_h0)
							}(_pat_let34_0)
						}(_pat_let_tv33)
					}(_pat_let33_0)
				}(Companion_Default___.Fm42(_1589_i0, p3, p2, globalState))
			})()
			_ = _rhs262
			_1604_v13 = _rhs258
			_1605_v14 = _rhs259
			r0 = _rhs260
			_1606_v15 = _rhs261
			_1607_v16 = _rhs262
			var _1612_v17 *C1
			_ = _1612_v17
			var _nw292 *C1 = New_C1_()
			_ = _nw292
			_nw292.Ctor__(_1588_v0)
			_1612_v17 = _nw292
			var _1613_v18 D8
			_ = _1613_v18
			_1613_v18 = Companion_D8_.Create_DC16_(_1612_v17)
			var _rhs263 D8 = _1613_v18
			_ = _rhs263
			var _rhs264 _dafny.Int = (p0).Plus((_dafny.Zero).Minus(_1589_i0))
			_ = _rhs264
			var _rhs265 bool = (func() bool {
				if p2 {
					return false
				}
				return p2
			})()
			_ = _rhs265
			_1613_v18 = _rhs263
			r2 = _rhs264
			r0 = _rhs265
			var _1614_v19 *C3
			_ = _1614_v19
			var _nw293 *C3 = New_C3_()
			_ = _nw293
			_nw293.Ctor__(_1589_i0)
			_1614_v19 = _nw293
		}
		var _1615_v20 _dafny.Sequence
		_ = _1615_v20
		_1615_v20 = _dafny.SeqOf(p3)
		r2 = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1615_v20, _1615_v20)).Cardinality()))
		var _1616_i3 _dafny.Int
		_ = _1616_i3
		_1616_i3 = _dafny.Zero
		{
			for p3 {
				{
					if (_1616_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1616_i3 = (_1616_i3).Plus(_dafny.One)
					_1588_v0 = _1588_v0
					r0 = !(!(p2) || (p3))
					var _1617_v21 _dafny.MultiSet
					_ = _1617_v21
					_1617_v21 = _dafny.MultiSetOf(_dafny.IntOfInt64(660), p0)
					if (Companion_Default___.SafeDivisionInt(p0, (_this).F0())).Cmp((_1617_v21).Cardinality()) <= 0 {
						var _1618_v22 _dafny.Map
						_ = _1618_v22
						_1618_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1588_v0)
						var _1619_v23 _dafny.Map
						_ = _1619_v23
						_1619_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if (_1618_v22).Contains(false) {
								return (_1618_v22).Get(false).(_dafny.Sequence)
							}
							return _1588_v0
						})(), (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1618_v22).Contains(false) {
								return (_1618_v22).Get(false).(_dafny.Sequence)
							}
							return _1588_v0
						})()).Cardinality()))).Uint32(), _dafny.CodePoint('p')), _1588_v0)
						var _1620_v24 _dafny.Map
						_ = _1620_v24
						_1620_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("darjgmy")).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-130))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg93 _dafny.Int) interface{} {
								return coer93(arg93)
							}
						}(func(_1621_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('s')
						})))
						_1619_v23 = (_1619_v23).Update((func() _dafny.Sequence {
							if (_1620_v24).Contains(p0) {
								return (_1620_v24).Get(p0).(_dafny.Sequence)
							}
							return _1588_v0
						})(), _dafny.UnicodeSeqOfUtf8Bytes("min"))
						_1588_v0 = _1588_v0
						r2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(205), (_this).F0())
						var _1622_v25 _dafny.Map
						_ = _1622_v25
						_1622_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F0())).Update((_this).F0(), (_this).F0()), _dafny.UnicodeSeqOfUtf8Bytes("ja"))
						var _1623_v26 _dafny.Map
						_ = _1623_v26
						_1623_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F0())
						_1622_v25 = (_1622_v25).Update(_1623_v26, _dafny.UnicodeSeqOfUtf8Bytes("sxeuxp"))
						var _rhs266 _dafny.Int = Companion_Default___.SafeModuloInt(p0, p0)
						_ = _rhs266
						var _rhs267 bool = p3
						_ = _rhs267
						var _rhs268 bool = p3
						_ = _rhs268
						var _rhs269 bool = (_this).Fm6(true, globalState)
						_ = _rhs269
						r2 = _rhs266
						r0 = _rhs267
						r0 = _rhs268
						r0 = _rhs269
					} else {
						r2 = (_this).F0()
						var _1624_v27 *C3
						_ = _1624_v27
						var _nw294 *C3 = New_C3_()
						_ = _nw294
						_nw294.Ctor__((_this).F0())
						_1624_v27 = _nw294
						var _1625_v28 _dafny.Array
						_ = _1625_v28
						var _nwElement0_71 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F0(), p0)
						_ = _nwElement0_71
						var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(8))
						_ = _nw295
						(_nw295).ArraySet1(_nwElement0_71, 0)
						(_nw295).ArraySet1(((_this).F0()).Minus((_this).F0()), 1)
						(_nw295).ArraySet1(p0, 2)
						(_nw295).ArraySet1((_dafny.Zero).Minus((_this).F0()), 3)
						(_nw295).ArraySet1((_this).F0(), 4)
						(_nw295).ArraySet1((_dafny.Zero).Minus(p0), 5)
						(_nw295).ArraySet1((_this).F0(), 6)
						(_nw295).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F0(), (func() _dafny.Int {
							if (_1617_v21).Contains(p0) {
								return (_1617_v21).Multiplicity(p0)
							}
							return p0
						})()), 7)
						_1625_v28 = _nw295
						var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1625_v28), 0))
						_ = _index286
						(_1625_v28).ArraySet1(_dafny.IntOfInt64(842), (_index286).Int())
						var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1625_v28), 0))
						_ = _index287
						(_1625_v28).ArraySet1(p0, (_index287).Int())
						var _1626_v29 _dafny.Array
						_ = _1626_v29
						var _len0_33 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_33
						var _nw296 _dafny.Array
						_ = _nw296
						if _len0_33.Cmp(_dafny.Zero) == 0 {
							_nw296 = _dafny.NewArray(_len0_33)
						} else {
							var _init33 func(_dafny.Int) _dafny.Map = func(_1627_i5 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(33), _dafny.IntOfInt64(753))
							}
							_ = _init33
							var _element0_33 = _init33(_dafny.Zero)
							_ = _element0_33
							_nw296 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
							(_nw296).ArraySet1(_element0_33, 0)
							var _nativeLen0_33 = (_len0_33).Int()
							_ = _nativeLen0_33
							for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
								(_nw296).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
							}
						}
						_1626_v29 = _nw296
						var _1628_v30 _dafny.Map
						_ = _1628_v30
						_1628_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1615_v20).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1615_v20).Cardinality()))).Uint32()).(bool), p3)
						var _1629_v31 _dafny.Map
						_ = _1629_v31
						_1629_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1626_v29), _1628_v30)
						_1629_v31 = (_1629_v31).Update(_1626_v29, _1628_v30)
						var _1630_v32 _dafny.Map
						_ = _1630_v32
						_1630_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _1631_v33 _dafny.Sequence
						_ = _1631_v33
						_1631_v33 = _dafny.SeqOf((_this).F0())
						r0 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
							if (_1617_v21).Contains((_1630_v32).Cardinality()) {
								return (_1617_v21).Multiplicity((_1630_v32).Cardinality())
							}
							return (_dafny.Zero).Minus(p0)
						})(), !(p3))).Cardinality()).Cmp(_dafny.IntOfUint32((_1631_v33).Cardinality())) <= 0
					}
					var _arr7 _dafny.Array = _this.F1
					_ = _arr7
					var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_this.F1), 0))
					_ = _index288
					(_arr7).ArraySet1((p0).Cmp(p0) == 0, (_index288).Int())
					var _arr8 _dafny.Array = _this.F1
					_ = _arr8
					var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_this.F1), 0))
					_ = _index289
					(_arr8).ArraySet1(p2, (_index289).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1632_v34 *C2
		_ = _1632_v34
		var _nw297 *C2 = New_C2_()
		_ = _nw297
		_nw297.Ctor__(Companion_Default___.SafeModuloInt((_this).F0(), _dafny.IntOfUint32((_1588_v0).Cardinality())))
		_1632_v34 = _nw297
		var _1633_v35 _dafny.Map
		_ = _1633_v35
		_1633_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1588_v0).Cardinality()))
		var _1634_v36 _dafny.Map
		_ = _1634_v36
		_1634_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F0()), p2)
		var _1635_v37 _dafny.Sequence
		_ = _1635_v37
		_1635_v37 = _dafny.SeqOf((_1632_v34).Fm1(true, globalState), (_1634_v36).Cardinality(), p0)
		var _1636_v38 _dafny.MultiSet
		_ = _1636_v38
		_1636_v38 = _dafny.MultiSetOf(p2)
		var _1637_v39 _dafny.Set
		_ = _1637_v39
		_1637_v39 = _dafny.SetOf(p3, p2)
		var _1638_v40 _dafny.Array
		_ = _1638_v40
		var _nwElement0_72 _dafny.Int = (_this).Fm1(p3, globalState)
		_ = _nwElement0_72
		var _nw298 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(22))
		_ = _nw298
		(_nw298).ArraySet1(_nwElement0_72, 0)
		(_nw298).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-749))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}(func(_1639_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), _dafny.UnicodeSeqOfUtf8Bytes("dhwnmp"))).Cardinality()), 1)
		(_nw298).ArraySet1(p0, 2)
		(_nw298).ArraySet1((_this).F0(), 3)
		(_nw298).ArraySet1((_this).F0(), 4)
		(_nw298).ArraySet1((_this).F0(), 5)
		(_nw298).ArraySet1((_this).F0(), 6)
		(_nw298).ArraySet1(((_1633_v35).Cardinality()).Plus(_dafny.IntOfUint32((_1635_v37).Cardinality())), 7)
		(_nw298).ArraySet1((func() _dafny.Int {
			if (_1636_v38).Contains(p3) {
				return (_1636_v38).Multiplicity(p3)
			}
			return (_1637_v39).Cardinality()
		})(), 8)
		(_nw298).ArraySet1((func() _dafny.Int {
			if p2 {
				return (_1632_v34).Fm1(p3, globalState)
			}
			return p0
		})(), 9)
		(_nw298).ArraySet1((p0).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-8))), 10)
		(_nw298).ArraySet1(p0, 11)
		(_nw298).ArraySet1(((_dafny.Zero).Minus((_this).F0())).Times((_dafny.Zero).Minus(p0)), 12)
		(_nw298).ArraySet1(p0, 13)
		(_nw298).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1588_v0).Cardinality()), (_this).F0()), 14)
		(_nw298).ArraySet1(_dafny.IntOfUint32((_1588_v0).Cardinality()), 15)
		(_nw298).ArraySet1(_dafny.IntOfUint32((_1588_v0).Cardinality()), 16)
		(_nw298).ArraySet1((_this).F0(), 17)
		(_nw298).ArraySet1((p0).Minus((_this).F0()), 18)
		(_nw298).ArraySet1((_this).F0(), 19)
		(_nw298).ArraySet1(_dafny.IntOfUint32((_1588_v0).Cardinality()), 20)
		(_nw298).ArraySet1(_dafny.IntOfUint32((_1635_v37).Cardinality()), 21)
		_1638_v40 = _nw298
		_1638_v40 = _1638_v40
		var _1640_v41 _dafny.Set
		_ = _1640_v41
		_1640_v41 = _dafny.SetOf(p0)
		r0 = !(_1640_v41).Equals((_1640_v41).Intersection(_1640_v41))
		var _1641_v42 _dafny.Map
		_ = _1641_v42
		_1641_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
		var _1642_v43 _dafny.Map
		_ = _1642_v43
		_1642_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1588_v0).Cardinality()), _1641_v42)
		var _nwElement0_73 _dafny.Map = _1641_v42
		_ = _nwElement0_73
		var _nw299 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(24))
		_ = _nw299
		(_nw299).ArraySet1(_nwElement0_73, 0)
		(_nw299).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3), 1)
		(_nw299).ArraySet1(_1641_v42, 2)
		(_nw299).ArraySet1(_1641_v42, 3)
		(_nw299).ArraySet1(_1641_v42, 4)
		(_nw299).ArraySet1((func() _dafny.Map {
			if p2 {
				return _1641_v42
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, true)
		})(), 5)
		(_nw299).ArraySet1(_1641_v42, 6)
		(_nw299).ArraySet1((func() _dafny.Map {
			if p2 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			}
			return (_1641_v42).Update(p2, (_this).Fm6(false, globalState))
		})(), 7)
		(_nw299).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p2), 8)
		(_nw299).ArraySet1((_1641_v42).Merge(_1641_v42), 9)
		(_nw299).ArraySet1((_1641_v42).Update(p3, !(p2)), 10)
		(_nw299).ArraySet1(_1641_v42, 11)
		(_nw299).ArraySet1(Companion_Default___.Fm21(globalState), 12)
		(_nw299).ArraySet1(_1641_v42, 13)
		(_nw299).ArraySet1(((func() _dafny.Map {
			if (_1642_v43).Contains(p0) {
				return (_1642_v43).Get(p0).(_dafny.Map)
			}
			return _1641_v42
		})()).Merge(Companion_Default___.Fm21(globalState)), 14)
		(_nw299).ArraySet1((_1641_v42).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)), 15)
		(_nw299).ArraySet1(_1641_v42, 16)
		(_nw299).ArraySet1(_1641_v42, 17)
		(_nw299).ArraySet1(_1641_v42, 18)
		(_nw299).ArraySet1(_1641_v42, 19)
		(_nw299).ArraySet1((_1641_v42).Merge(_1641_v42), 20)
		(_nw299).ArraySet1(_1641_v42, 21)
		(_nw299).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2), 22)
		(_nw299).ArraySet1(_1641_v42, 23)
		r1 = _nw299
		var _1643_v44 _dafny.Map
		_ = _1643_v44
		_1643_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D6_.Create_DC12_(p2, p3)).Dtor_cf22(), (_dafny.Zero).Minus((_this).F0()))
		r2 = (func() _dafny.Int {
			if (_1643_v44).Contains(!(p3)) {
				return (_1643_v44).Get(!(p3)).(_dafny.Int)
			}
			return ((_this).F0()).Plus((_this).F0())
		})()
		return r0, r1, r2
	}
}
func (_this *C5) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1644_v0 _dafny.CodePoint
		_ = _1644_v0
		_1644_v0 = _dafny.CodePoint('f')
		var _1645_v1 _dafny.Array
		_ = _1645_v1
		var _nwElement0_74 _dafny.CodePoint = _1644_v0
		_ = _nwElement0_74
		var _nw300 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(2))
		_ = _nw300
		(_nw300).ArraySet1CodePoint(_nwElement0_74, 0)
		(_nw300).ArraySet1CodePoint((func() _dafny.CodePoint {
			if true {
				return _1644_v0
			}
			return _1644_v0
		})(), 1)
		_1645_v1 = _nw300
		var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
		_ = _index290
		(_1645_v1).ArraySet1CodePoint(_1644_v0, (_index290).Int())
		var _1646_v2 _dafny.Array
		_ = _1646_v2
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_34
		var _nw301 _dafny.Array
		_ = _nw301
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw301 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Sequence = func(_1647_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("tfbmfbdi")
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw301 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw301).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw301).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1646_v2 = _nw301
		var _1648_v3 _dafny.Sequence
		_ = _1648_v3
		_1648_v3 = _dafny.UnicodeSeqOfUtf8Bytes("bu")
		var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
		_ = _index291
		(_1646_v2).ArraySet1(_1648_v3, (_index291).Int())
		var _1649_v4 _dafny.Map
		_ = _1649_v4
		_1649_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F0())
		var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
		_ = _index292
		var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
		_ = _index293
		var _rhs270 _dafny.CodePoint = _dafny.CodePoint('a')
		_ = _rhs270
		var _rhs271 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1648_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1648_v3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1648_v3).Cardinality()), _dafny.IntOfUint32((_1648_v3).Cardinality()))).Uint32(), _1644_v0), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg95 _dafny.Int) interface{} {
				return coer95(arg95)
			}
		}((func(_1650_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1651_i1 _dafny.Int) _dafny.CodePoint {
				return _1650_v0
			}
		})(_1644_v0))), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_1649_v4).Contains(p2) {
				return (_1649_v4).Get(p2).(_dafny.Int)
			}
			return p1
		})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg96 _dafny.Int) interface{} {
				return coer96(arg96)
			}
		}((func(_1652_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1653_i1 _dafny.Int) _dafny.CodePoint {
				return _1652_v0
			}
		})(_1644_v0)))).Cardinality()))).Uint32(), _1644_v0)))
		_ = _rhs271
		var _lhs136 _dafny.Array = _1645_v1
		_ = _lhs136
		var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
		_ = _lhs137
		var _lhs138 _dafny.Array = _1646_v2
		_ = _lhs138
		var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
		_ = _lhs139
		(_lhs136).ArraySet1CodePoint(_rhs270, (_lhs137).Int())
		(_lhs138).ArraySet1(_rhs271, (_lhs139).Int())
		var _1654_v5 D6
		_ = _1654_v5
		_1654_v5 = Companion_D6_.Create_DC11_((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence))
		var _1655_v6 _dafny.Sequence
		_ = _1655_v6
		_1655_v6 = _dafny.SeqOf((_1654_v5).Dtor_cf21(), (_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence))
		var _rhs272 _dafny.Sequence = _1648_v3
		_ = _rhs272
		var _rhs273 _dafny.Int = (_dafny.SetOf((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update(_1648_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1648_v3).Cardinality()))).Uint32(), (_1645_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))).Int())), (_1655_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1655_v6).Cardinality()))).Uint32()).(_dafny.Sequence), (_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence))).Cardinality()
		_ = _rhs273
		_1648_v3 = _rhs272
		r0 = _rhs273
		var _1656_v7 _dafny.Sequence
		_ = _1656_v7
		_1656_v7 = _dafny.SeqOf(_this.F1)
		var _1657_v8 D11
		_ = _1657_v8
		_1657_v8 = Companion_D11_.Create_DC20_((_1656_v7).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1656_v7).Cardinality()))).Uint32()).(_dafny.Array))
		var _1658_v9 _dafny.Sequence
		_ = _1658_v9
		_1658_v9 = _dafny.SeqOf(_1657_v8, _1657_v8)
		var _1659_v10 _dafny.Sequence
		_ = _1659_v10
		_1659_v10 = _dafny.SeqOf(_1658_v9)
		if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_1659_v10).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1659_v10).Cardinality()))).Uint32()).(_dafny.Sequence), _1658_v9), (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1659_v10).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1659_v10).Cardinality()))).Uint32()).(_dafny.Sequence), _1658_v9)).Cardinality()))).Uint32(), _1657_v8), _1657_v8) {
			r0 = (p0).Times(_dafny.IntOfInt64(-22))
			var _1660_v11 _dafny.Array
			_ = _1660_v11
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_35
			var _nw302 _dafny.Array
			_ = _nw302
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw302 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = (func(_1661_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1662_i2 _dafny.Int) _dafny.Int {
						return (_1662_i2).Times(_1661_p0)
					}
				})(p0)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw302 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw302).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw302).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1660_v11 = _nw302
			var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
			_ = _index294
			(_1660_v11).ArraySet1((_dafny.IntOfInt64(859)).Plus(p0), (_index294).Int())
			var _1663_v12 _dafny.Set
			_ = _1663_v12
			_1663_v12 = _dafny.SetOf((_dafny.Zero).Minus(p1), (_this).F0())
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
			_ = _index295
			(_1660_v11).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F0(), ((func() _dafny.Set {
				if p2 {
					return _1663_v12
				}
				return _1663_v12
			})()).Cardinality()), (_index295).Int())
			if p2 {
				var _1664_v13 _dafny.Sequence
				_ = _1664_v13
				_1664_v13 = _dafny.SeqOf(p0)
				var _1665_v14 _dafny.Map
				_ = _1665_v14
				_1665_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1644_v0, _dafny.MultiSetOf((_1660_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))).Int()).(_dafny.Int), (_1664_v13).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1649_v4).Contains(p2) {
						return (_1649_v4).Get(p2).(_dafny.Int)
					}
					return (_this).F0()
				})(), _dafny.IntOfUint32((_1664_v13).Cardinality()))).Uint32()).(_dafny.Int)))
				var _1666_v15 _dafny.Map
				_ = _1666_v15
				_1666_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1644_v0, p0)
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
				_ = _index296
				(_1660_v11).ArraySet1((((_1665_v14).Merge(Companion_Default___.Fm43(globalState))).Cardinality()).Minus((func() _dafny.Int {
					if (_1666_v15).Contains(_1644_v0) {
						return (_1666_v15).Get(_1644_v0).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1664_v13).Cardinality())
				})()), (_index296).Int())
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
				_ = _index297
				(_1645_v1).ArraySet1CodePoint(Companion_Default___.Fm37(p1, globalState), (_index297).Int())
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
				_ = _index298
				(_1660_v11).ArraySet1((_this).F0(), (_index298).Int())
				var _1667_v16 _dafny.Array
				_ = _1667_v16
				var _out27 _dafny.Array
				_ = _out27
				_out27 = Companion_Default___.M0(p2, p2, globalState)
				_1667_v16 = _out27
				r0 = (_1660_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))).Int()).(_dafny.Int)
			} else {
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
				_ = _index299
				(_1660_v11).ArraySet1((_this).F0(), (_index299).Int())
				var _1668_v17 bool
				_ = _1668_v17
				_1668_v17 = false
				_1668_v17 = _1668_v17
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1660_v11), 0))
				_ = _index300
				(_1660_v11).ArraySet1((_this).F0(), (_index300).Int())
				var _1669_v18 *C3
				_ = _1669_v18
				var _nw303 *C3 = New_C3_()
				_ = _nw303
				_nw303.Ctor__(p0)
				_1669_v18 = _nw303
				var _1670_v19 _dafny.Map
				_ = _1670_v19
				_1670_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1671_v20 _dafny.Map
				_ = _1671_v20
				_1671_v20 = _1670_v19
				var _1672_v21 _dafny.Map
				_ = _1672_v21
				_1672_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v20, _this.F1)
				_1672_v21 = (_1672_v21).Update(_1671_v20, _this.F1)
			}
			var _1673_v22 _dafny.Set
			_ = _1673_v22
			_1673_v22 = _dafny.SetOf(!(p2), p2, p2)
			var _1674_v23 _dafny.Map
			_ = _1674_v23
			_1674_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1673_v22, _dafny.UnicodeSeqOfUtf8Bytes("qgpvdxtv"))
			_1674_v23 = (_1674_v23).Update(_1673_v22, _1648_v3)
			var _1675_v24 _dafny.Array
			_ = _1675_v24
			var _nw304 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw304
			_1675_v24 = _nw304
			var _pat_let_tv34 = _1648_v3
			_ = _pat_let_tv34
			var _1676_v25 _dafny.Map
			_ = _1676_v25
			_1676_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let35_0 D6) D6 {
				return func(_1677_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let36_0 _dafny.Sequence) D6 {
						return func(_1678_dt__update_hcf21_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC11_(_1678_dt__update_hcf21_h0)
						}(_pat_let36_0)
					}(_pat_let_tv34)
				}(_pat_let35_0)
			}(_1654_v5)).Dtor_cf21(), _1675_v24)
			_1676_v25 = (_1676_v25).Update(_dafny.Companion_Sequence_.Update(_1648_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-940), _dafny.IntOfUint32((_1648_v3).Cardinality()))).Uint32(), _1644_v0), _1675_v24)
		} else {
			var _1679_v26 *C0
			_ = _1679_v26
			var _nw305 *C0 = New_C0_()
			_ = _nw305
			_nw305.Ctor__((_this).Fm6(p2, globalState), p2)
			_1679_v26 = _nw305
			r0 = _dafny.IntOfInt64(183)
			var _1680_v27 _dafny.Sequence
			_ = _1680_v27
			_1680_v27 = _dafny.SeqOf(_dafny.IntOfUint32((_1648_v3).Cardinality()), (_this).F0(), (_this).F0())
			var _1681_v28 *C3
			_ = _1681_v28
			var _nw306 *C3 = New_C3_()
			_ = _nw306
			_nw306.Ctor__((_1680_v27).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1680_v27).Cardinality()))).Uint32()).(_dafny.Int))
			_1681_v28 = _nw306
			var _arr9 _dafny.Array = _this.F1
			_ = _arr9
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_this.F1), 0))
			_ = _index301
			(_arr9).ArraySet1(_1679_v26.F5, (_index301).Int())
			var _arr10 _dafny.Array = _this.F1
			_ = _arr10
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_this.F1), 0))
			_ = _index302
			(_arr10).ArraySet1(p2, (_index302).Int())
			if _1679_v26.F4 {
				var _1682_v29 *C3
				_ = _1682_v29
				var _nw307 *C3 = New_C3_()
				_ = _nw307
				_nw307.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1648_v3, _1648_v3)).Cardinality()))
				_1682_v29 = _nw307
				_1682_v29 = (func() *C3 {
					if _1679_v26.F5 {
						return _1682_v29
					}
					return _1682_v29
				})()
				r0 = (_this).F0()
				var _1683_v30 _dafny.MultiSet
				_ = _1683_v30
				_1683_v30 = _dafny.MultiSetOf(false)
				r0 = ((p1).Times((_1683_v30).Cardinality())).Minus((_this).F0())
				var _1684_v31 _dafny.Array
				_ = _1684_v31
				var _nw308 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
				_ = _nw308
				_1684_v31 = _nw308
				var _1685_v32 T2
				_ = _1685_v32
				var _nw309 *C4 = New_C4_()
				_ = _nw309
				_nw309.Ctor__(_dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				_1685_v32 = _nw309
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1684_v31), 0))
				_ = _index303
				(_1684_v31).ArraySet1(_1685_v32, (_index303).Int())
				var _1686_v33 _dafny.Array
				_ = _1686_v33
				var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw310
				_1686_v33 = _nw310
				var _1687_v34 _dafny.Sequence
				_ = _1687_v34
				_1687_v34 = _dafny.SeqOf(_1686_v33, _1686_v33)
				var _1688_v35 _dafny.Array
				_ = _1688_v35
				var _nwElement0_75 _dafny.Array = _1686_v33
				_ = _nwElement0_75
				var _nw311 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(16))
				_ = _nw311
				(_nw311).ArraySet1(_nwElement0_75, 0)
				(_nw311).ArraySet1(_1686_v33, 1)
				(_nw311).ArraySet1(_1686_v33, 2)
				(_nw311).ArraySet1(_1686_v33, 3)
				(_nw311).ArraySet1(_1686_v33, 4)
				(_nw311).ArraySet1(_1686_v33, 5)
				(_nw311).ArraySet1(_1686_v33, 6)
				(_nw311).ArraySet1(_1686_v33, 7)
				(_nw311).ArraySet1(_1686_v33, 8)
				(_nw311).ArraySet1(_1686_v33, 9)
				(_nw311).ArraySet1(_1686_v33, 10)
				(_nw311).ArraySet1(_1686_v33, 11)
				(_nw311).ArraySet1(_1686_v33, 12)
				(_nw311).ArraySet1((_1687_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.IntOfUint32((_1687_v34).Cardinality()))).Uint32()).(_dafny.Array), 13)
				(_nw311).ArraySet1(_1686_v33, 14)
				(_nw311).ArraySet1(_1686_v33, 15)
				_1688_v35 = _nw311
				var _1689_v36 _dafny.Map
				_ = _1689_v36
				_1689_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v35, _1680_v27)
				var _1690_v37 _dafny.Map
				_ = _1690_v37
				_1690_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F0())
				var _1691_v38 _dafny.Set
				_ = _1691_v38
				_1691_v38 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1654_v5, (_1690_v37).Cardinality()))
				var _1692_v39 _dafny.Map
				_ = _1692_v39
				_1692_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC11_(_1648_v3), p1)
				var _1693_v40 D11
				_ = _1693_v40
				_1693_v40 = Companion_D11_.Create_DC21_(_1644_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}(func(_1694_i3 _dafny.Int) _dafny.Int {
					return (_this).F0()
				}))).Cardinality()), p0)
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1684_v31), 0))
				_ = _index304
				var _rhs274 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm36(_1679_v26.F4, _dafny.CodePoint('i'), (_this).Fm1(_1679_v26.F5, globalState), globalState), (func() _dafny.Sequence {
					if (_1689_v36).Contains(_1688_v35) {
						return (_1689_v36).Get(_1688_v35).(_dafny.Sequence)
					}
					return _1680_v27
				})())
				_ = _rhs274
				var _rhs275 _dafny.Int = Companion_Default___.SafeModuloInt(p1, (_this).F0())
				_ = _rhs275
				var _rhs276 bool = ((func() _dafny.Set {
					if _1679_v26.F5 {
						return _1691_v38
					}
					return (_dafny.SetOf(_1692_v39, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1654_v5, _1685_v32.F2())))
				})()).IsSubsetOf(_1691_v38)
				_ = _rhs276
				var _rhs277 T2 = _1685_v32
				_ = _rhs277
				var _rhs278 _dafny.Int = (p1).Minus((_1693_v40).Dtor_cf38())
				_ = _rhs278
				var _lhs140 *C0 = _1679_v26
				_ = _lhs140
				var _lhs141 _dafny.Array = _1684_v31
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1684_v31), 0))
				_ = _lhs142
				var _lhs143 T2 = _1685_v32
				_ = _lhs143
				_1680_v27 = _rhs274
				r0 = _rhs275
				_lhs140.F5 = _rhs276
				(_lhs141).ArraySet1(_rhs277, (_lhs142).Int())
				_lhs143.F2_set_(_rhs278)
				var _1695_v41 _dafny.Set
				_ = _1695_v41
				_1695_v41 = _dafny.SetOf(_1679_v26.F5)
				(_1685_v32).F2_set_((_1695_v41).Cardinality())
			} else {
				var _1696_v42 _dafny.Array
				_ = _1696_v42
				var _nw312 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw312
				_1696_v42 = _nw312
				var _1697_v43 _dafny.Sequence
				_ = _1697_v43
				_1697_v43 = _dafny.SeqOf(_1696_v42, _1696_v42)
				r0 = Companion_Default___.SafeDivisionInt(p1, (func() _dafny.Int {
					if _1679_v26.F5 {
						return p0
					}
					return _dafny.IntOfUint32((_1697_v43).Cardinality())
				})())
				var _1698_v44 _dafny.Set
				_ = _1698_v44
				_1698_v44 = _dafny.SetOf(_1679_v26.F4)
				var _1699_v45 D4
				_ = _1699_v45
				_1699_v45 = Companion_D4_.Create_DC5_(_1698_v44)
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
				_ = _index305
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
				_ = _index306
				var _rhs279 _dafny.Int = (_this).Fm1(_1679_v26.F5, globalState)
				_ = _rhs279
				var _rhs280 _dafny.CodePoint = (_1648_v3).Select((Companion_Default___.SafeIndex((p0).Times(_dafny.IntOfInt64(144)), _dafny.IntOfUint32((_1648_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs280
				var _rhs281 _dafny.Int = ((_dafny.SetOf(_1679_v26.F4)).Union((_1699_v45).Dtor_cf8())).Cardinality()
				_ = _rhs281
				var _rhs282 _dafny.Sequence = (_1655_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality()), _dafny.IntOfUint32((_1655_v6).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs282
				var _lhs144 _dafny.Array = _1645_v1
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
				_ = _lhs145
				var _lhs146 _dafny.Array = _1646_v2
				_ = _lhs146
				var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
				_ = _lhs147
				r0 = _rhs279
				(_lhs144).ArraySet1CodePoint(_rhs280, (_lhs145).Int())
				r0 = _rhs281
				(_lhs146).ArraySet1(_rhs282, (_lhs147).Int())
				r0 = (p0).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1680_v27, _1679_v26.F5)).Merge(func() _dafny.Map {
					var _coll60 = _dafny.NewMapBuilder()
					_ = _coll60
					for _iter65 := _dafny.Iterate((Companion_Default___.Fm44(p0, (_this).F0(), (_this).F0(), _1679_v26.F5, globalState)).Keys().Elements()); ; {
						_compr_60, _ok65 := _iter65()
						if !_ok65 {
							break
						}
						var _1700_v46 _dafny.Sequence
						_1700_v46 = interface{}(_compr_60).(_dafny.Sequence)
						if (Companion_Default___.Fm44(p0, (_this).F0(), (_this).F0(), _1679_v26.F5, globalState)).Contains(_1700_v46) {
							_coll60.Add(_1700_v46, false)
						}
					}
					return _coll60.ToMap()
				}())).Cardinality())
				var _1701_v47 _dafny.Set
				_ = _1701_v47
				_1701_v47 = _dafny.SetOf(p1)
				var _1702_v48 _dafny.Map
				_ = _1702_v48
				_1702_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if _1679_v26.F4 {
						return _1679_v26.F4
					}
					return _1679_v26.F5
				})(), (_dafny.SetOf(p0, _dafny.IntOfInt64(418))).Union(_1701_v47))
				_1702_v48 = (_1702_v48).Update(_1679_v26.F4, _1701_v47)
				var _1703_v49 *C0
				_ = _1703_v49
				var _nw313 *C0 = New_C0_()
				_ = _nw313
				_nw313.Ctor__(!((_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool)), _1679_v26.F5)
				_1703_v49 = _nw313
			}
		}
		var _hi10 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ltbe")).Cardinality())
		_ = _hi10
		for _1704_i4 := (_dafny.Zero).Minus((p1).Minus(_dafny.IntOfUint32((_1648_v3).Cardinality()))); _1704_i4.Cmp(_hi10) < 0; _1704_i4 = _1704_i4.Plus(_dafny.One) {
			var _1705_v50 *C4
			_ = _1705_v50
			var _nw314 *C4 = New_C4_()
			_ = _nw314
			_nw314.Ctor__(_dafny.IntOfInt64(-578))
			_1705_v50 = _nw314
			var _1706_v51 _dafny.Set
			_ = _1706_v51
			_1706_v51 = _dafny.SetOf(p0, p0)
			var _1707_v53 _dafny.Map
			_ = _1707_v53
			_1707_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(p2, globalState), (_1706_v51).IsDisjointFrom(func() _dafny.Set {
				var _coll61 = _dafny.NewBuilder()
				_ = _coll61
				for _iter66 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(775), _dafny.IntOfInt64(402))); ; {
					_compr_61, _ok66 := _iter66()
					if !_ok66 {
						break
					}
					var _1708_v52 _dafny.Int
					_1708_v52 = interface{}(_compr_61).(_dafny.Int)
					if ((_dafny.IntOfInt64(775)).Cmp(_1708_v52) <= 0) && ((_1708_v52).Cmp(_dafny.IntOfInt64(402)) < 0) {
						_coll61.Add(Companion_Default___.SafeDivisionInt(_1708_v52, p0))
					}
				}
				return _coll61.ToSet()
			}()))
			_1707_v53 = (_1707_v53).Update((p2) == (p2), p2)
			var _1709_v54 _dafny.Array
			_ = _1709_v54
			var _nw315 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw315
			_1709_v54 = _nw315
			_1709_v54 = _1709_v54
			if p2 {
				var _1710_v55 _dafny.Map
				_ = _1710_v55
				_1710_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1644_v0, _1654_v5)
				var _1711_v56 _dafny.Map
				_ = _1711_v56
				_1711_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1710_v55, _1645_v1)
				r0 = (_1711_v56).Cardinality()
				var _1712_v57 _dafny.Map
				_ = _1712_v57
				_1712_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1707_v53).Cardinality(), ((_dafny.Zero).Minus(p1)).Cmp(_1704_i4) > 0)
				_1712_v57 = ((_1712_v57).Merge(_1712_v57)).Merge(_1712_v57)
				var _1713_v58 _dafny.Set
				_ = _1713_v58
				_1713_v58 = _dafny.SetOf(p2, p2, !(p2))
				var _arr11 _dafny.Array = _this.F1
				_ = _arr11
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))
				_ = _index307
				(_arr11).ArraySet1((_1704_i4).Cmp((_1713_v58).Cardinality()) >= 0, (_index307).Int())
				var _arr12 _dafny.Array = _this.F1
				_ = _arr12
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))
				_ = _index308
				(_arr12).ArraySet1(p2, (_index308).Int())
				var _arr13 _dafny.Array = _this.F1
				_ = _arr13
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))
				_ = _index309
				(_arr13).ArraySet1(!(p2), (_index309).Int())
				var _arr14 _dafny.Array = _this.F1
				_ = _arr14
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))
				_ = _index310
				var _rhs283 bool = (_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool)
				_ = _rhs283
				var _rhs284 _dafny.Int = (func() _dafny.Int {
					if (_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool) {
						return (_1705_v50).Fm1((_this.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))).Int()).(bool), globalState)
					}
					return Companion_Default___.SafeDivisionInt(p0, _1704_i4)
				})()
				_ = _rhs284
				var _lhs148 _dafny.Array = _this.F1
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_this.F1), 0))
				_ = _lhs149
				(_lhs148).ArraySet1(_rhs283, (_lhs149).Int())
				r0 = _rhs284
			} else {
				var _1714_v59 bool
				_ = _1714_v59
				_1714_v59 = false
				var _1715_v60 _dafny.Map
				_ = _1715_v60
				_1715_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _1716_v61 _dafny.Map
				_ = _1716_v61
				_1716_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1717_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1718_i5 _dafny.Int) _dafny.CodePoint {
						return _1717_v0
					}
				})(_1644_v0))), _1715_v60)
				var _1719_v62 D5
				_ = _1719_v62
				_1719_v62 = Companion_D5_.Create_DC10_(false, p1, _1716_v61, p2)
				_1714_v59 = (_1719_v62).Dtor_cf17()
				var _1720_v63 *C3
				_ = _1720_v63
				var _nw316 *C3 = New_C3_()
				_ = _nw316
				_nw316.Ctor__(p1)
				_1720_v63 = _nw316
				var _1721_v64 _dafny.Sequence
				_ = _1721_v64
				_1721_v64 = _dafny.SeqOf(_1704_i4)
				r0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1720_v63).Fm10(_1704_i4, globalState), _1721_v64)).Cardinality()
				r0 = _dafny.IntOfInt64(-570)
				var _1722_v65 _dafny.Array
				_ = _1722_v65
				var _nw317 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw317
				_1722_v65 = _nw317
				var _1723_v66 _dafny.Array
				_ = _1723_v66
				var _nwElement0_76 D5 = Companion_D5_.Create_DC10_(p2, p1, _1716_v61, _1714_v59)
				_ = _nwElement0_76
				var _nw318 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(8))
				_ = _nw318
				(_nw318).ArraySet1(_nwElement0_76, 0)
				(_nw318).ArraySet1(_1719_v62, 1)
				(_nw318).ArraySet1(Companion_D5_.Create_DC10_(p2, _1704_i4, _1716_v61, _1714_v59), 2)
				(_nw318).ArraySet1(_1719_v62, 3)
				(_nw318).ArraySet1(Companion_D5_.Create_DC10_(p2, _dafny.IntOfUint32((_dafny.SeqOf(p1, _1704_i4)).Cardinality()), _1716_v61, p2), 4)
				(_nw318).ArraySet1(_1719_v62, 5)
				(_nw318).ArraySet1(_1719_v62, 6)
				(_nw318).ArraySet1(_1719_v62, 7)
				_1723_v66 = _nw318
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1722_v65), 0))
				_ = _index311
				(_1722_v65).ArraySet1(_1723_v66, (_index311).Int())
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1722_v65), 0))
				_ = _index312
				(_1722_v65).ArraySet1(_1723_v66, (_index312).Int())
			}
		}
		var _hi11 _dafny.Int = p1
		_ = _hi11
		for _1724_i6 := (_dafny.IntOfInt64(58)).Plus((_this).F0()); _1724_i6.Cmp(_hi11) < 0; _1724_i6 = _1724_i6.Plus(_dafny.One) {
			var _pat_let_tv35 = _1644_v0
			_ = _pat_let_tv35
			var _1725_v67 D0
			_ = _1725_v67
			_1725_v67 = Companion_D0_.Create_DC0_((func(_pat_let37_0 D0) D0 {
				return func(_1726_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let38_0 _dafny.CodePoint) D0 {
						return func(_1727_dt__update_hcf3_h0 _dafny.CodePoint) D0 {
							return Companion_D0_.Create_DC1_((_1726_dt__update__tmp_h2).Dtor_cf1(), (_1726_dt__update__tmp_h2).Dtor_cf2(), _1727_dt__update_hcf3_h0, (_1726_dt__update__tmp_h2).Dtor_cf4())
						}(_pat_let38_0)
					}(_pat_let_tv35)
				}(_pat_let37_0)
			}(Companion_D0_.Create_DC1_(!(p2), p2, (_1645_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))).Int()), p0))).Dtor_cf2())
			_1725_v67 = _1725_v67
			var _1728_v68 bool
			_ = _1728_v68
			_1728_v68 = false
			_1728_v68 = p2
			var _1729_v69 _dafny.Array
			_ = _1729_v69
			var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw319
			_1729_v69 = _nw319
			var _1730_v70 _dafny.Sequence
			_ = _1730_v70
			_1730_v70 = _dafny.SeqOf(_1728_v68, _1728_v68)
			var _1731_v71 _dafny.Sequence
			_ = _1731_v71
			_1731_v71 = _dafny.SeqOf(_1730_v70)
			var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_1729_v69), 0))
			_ = _index313
			(_1729_v69).ArraySet1(_1731_v71, (_index313).Int())
			var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_1729_v69), 0))
			_ = _index314
			(_1729_v69).ArraySet1(_1731_v71, (_index314).Int())
			r0 = p0
		}
		var _1732_i7 _dafny.Int
		_ = _1732_i7
		_1732_i7 = _dafny.Zero
		{
			for !(p2) {
				{
					if (_1732_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1732_i7 = (_1732_i7).Plus(_dafny.One)
					r0 = p1
					if ((_this).F0()).Cmp(_dafny.IntOfUint32((_1648_v3).Cardinality())) == 0 {
						var _1733_v72 _dafny.Array
						_ = _1733_v72
						var _len0_36 _dafny.Int = _dafny.IntOfInt64(16)
						_ = _len0_36
						var _nw320 _dafny.Array
						_ = _nw320
						if _len0_36.Cmp(_dafny.Zero) == 0 {
							_nw320 = _dafny.NewArray(_len0_36)
						} else {
							var _init36 func(_dafny.Int) _dafny.Int = (func(_1734_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1735_i8 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_1735_i8, (_dafny.Zero).Minus(_1734_p0))
								}
							})(p0)
							_ = _init36
							var _element0_36 = _init36(_dafny.Zero)
							_ = _element0_36
							_nw320 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
							(_nw320).ArraySet1(_element0_36, 0)
							var _nativeLen0_36 = (_len0_36).Int()
							_ = _nativeLen0_36
							for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
								(_nw320).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
							}
						}
						_1733_v72 = _nw320
						var _1736_v73 D4
						_ = _1736_v73
						_1736_v73 = Companion_D4_.Create_DC6_((_this).F0(), _1733_v72, p0)
						_1733_v72 = (_1736_v73).Dtor_cf10()
						var _1737_v74 _dafny.Map
						_ = _1737_v74
						_1737_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (p0).Cmp(_dafny.IntOfInt64(331)) != 0)
						_1737_v74 = (_1737_v74).Update(p2, !(_1737_v74).Equals(_1737_v74))
						var _1738_v75 _dafny.MultiSet
						_ = _1738_v75
						_1738_v75 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))
						var _1739_v76 _dafny.MultiSet
						_ = _1739_v76
						_1739_v76 = (_1738_v75).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality())))
						_1739_v76 = _1739_v76
						var _1740_v77 *C2
						_ = _1740_v77
						var _nw321 *C2 = New_C2_()
						_ = _nw321
						_nw321.Ctor__(p0)
						_1740_v77 = _nw321
						var _1741_v78 _dafny.Sequence
						_ = _1741_v78
						_1741_v78 = _dafny.SeqOf(_dafny.IntOfInt64(205), p1)
						_1741_v78 = _1741_v78
					} else {
						var _1742_v79 bool
						_ = _1742_v79
						_1742_v79 = false
						var _1743_v80 _dafny.Map
						_ = _1743_v80
						_1743_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _1744_v81 _dafny.Map
						_ = _1744_v81
						_1744_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1743_v80, p2)
						var _1745_v82 _dafny.Sequence
						_ = _1745_v82
						_1745_v82 = _dafny.SeqOf(p2)
						var _1746_v83 _dafny.Map
						_ = _1746_v83
						_1746_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), p0)
						var _rhs285 bool = p2
						_ = _rhs285
						var _rhs286 _dafny.Map = (_1649_v4).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1745_v82).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1745_v82).Cardinality()))).Uint32()).(bool)), (_1746_v83).Cardinality()))
						_ = _rhs286
						var _rhs287 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1743_v80, _1742_v79)).Merge(_1744_v81)
						_ = _rhs287
						var _rhs288 bool = _1742_v79
						_ = _rhs288
						var _rhs289 _dafny.Int = ((_dafny.IntOfUint32((_1745_v82).Cardinality())).Plus(p0)).Minus(p0)
						_ = _rhs289
						_1742_v79 = _rhs285
						_1649_v4 = _rhs286
						_1744_v81 = _rhs287
						_1742_v79 = _rhs288
						r0 = _rhs289
						var _1747_v84 _dafny.MultiSet
						_ = _1747_v84
						_1747_v84 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1648_v3, _dafny.Companion_Sequence_.Update((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_1743_v80).Cardinality(), _dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1644_v0))).Cardinality()))
						_1747_v84 = (_1747_v84).Intersection(_1747_v84)
						_1742_v79 = ((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfInt64(-473)) < 0
						_1742_v79 = (_dafny.IntOfInt64(532)).Cmp(_dafny.IntOfInt64(343)) < 0
						var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
						_ = _index315
						(_1646_v2).ArraySet1(_1648_v3, (_index315).Int())
					}
					var _source24 D11 = _1657_v8
					_ = _source24
					if _source24.Is_DC21() {
						var _1748___mcc_h0 _dafny.CodePoint = _source24.Get_().(D11_DC21).Cf36
						_ = _1748___mcc_h0
						var _1749___mcc_h1 _dafny.Int = _source24.Get_().(D11_DC21).Cf37
						_ = _1749___mcc_h1
						var _1750___mcc_h2 _dafny.Int = _source24.Get_().(D11_DC21).Cf38
						_ = _1750___mcc_h2
						var _1751_cf38 _dafny.Int = _1750___mcc_h2
						_ = _1751_cf38
						var _1752_cf37 _dafny.Int = _1749___mcc_h1
						_ = _1752_cf37
						var _1753_cf36 _dafny.CodePoint = _1748___mcc_h0
						_ = _1753_cf36
						var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))
						_ = _index316
						(_1646_v2).ArraySet1(_1648_v3, (_index316).Int())
						var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
						_ = _index317
						(_1645_v1).ArraySet1CodePoint(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_1752_cf37, _dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index317).Int())
						r0 = (_dafny.Zero).Minus(_1751_cf38)
						var _1754_v85 bool
						_ = _1754_v85
						_1754_v85 = true
						var _1755_v86 _dafny.MultiSet
						_ = _1755_v86
						_1755_v86 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfInt64(717), p1)
						var _1756_v87 _dafny.MultiSet
						_ = _1756_v87
						_1756_v87 = _dafny.MultiSetOf(_1755_v86)
						_1754_v85 = (_1756_v87).Equals(_1756_v87)
					} else if _source24.Is_DC22() {
						var _1757___mcc_h3 bool = _source24.Get_().(D11_DC22).Cf39
						_ = _1757___mcc_h3
						var _1758_cf39 bool = _1757___mcc_h3
						_ = _1758_cf39
						var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1645_v1), 0))
						_ = _index318
						(_1645_v1).ArraySet1CodePoint((_1648_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1648_v3).Cardinality()), _dafny.IntOfUint32((_1648_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index318).Int())
						r0 = p0
						var _1759_v88 _dafny.Map
						_ = _1759_v88
						_1759_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), p2)
						_1758_cf39 = ((_1759_v88).Merge(_1759_v88)).Contains((func() _dafny.Int {
							if !(p2) {
								return (_this).Fm1(!(p2), globalState)
							}
							return (_this).F0()
						})())
						_1759_v88 = _1759_v88
					} else {
						var _1760___mcc_h4 _dafny.Array = _source24.Get_().(D11_DC20).Cf35
						_ = _1760___mcc_h4
						var _1761_cf35 _dafny.Array = _1760___mcc_h4
						_ = _1761_cf35
						(_this).F1 = _1761_cf35
						r0 = (_this).F0()
						var _1762_v89 _dafny.Sequence
						_ = _1762_v89
						_1762_v89 = _dafny.SeqOf(_dafny.IntOfUint32(((_1646_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1646_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))
						var _1763_v90 _dafny.Map
						_ = _1763_v90
						_1763_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1762_v89).Cardinality()), (_this).Fm1((_this).Fm6(p2, globalState), globalState))
						var _1764_v91 _dafny.Sequence
						_ = _1764_v91
						_1764_v91 = _dafny.SeqOf(p1, (_1763_v90).Cardinality(), p1, (_this).F0())
						r0 = Companion_Default___.SafeDivisionInt(p1, (_1764_v91).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1764_v91).Cardinality()))).Uint32()).(_dafny.Int))
						var _1765_v92 _dafny.Set
						_ = _1765_v92
						_1765_v92 = _dafny.SetOf(p0)
						var _rhs290 _dafny.Int = ((_this).F0()).Plus((_1765_v92).Cardinality())
						_ = _rhs290
						var _rhs291 _dafny.Int = p1
						_ = _rhs291
						r0 = _rhs290
						r0 = _rhs291
					}
					var _1766_v93 bool
					_ = _1766_v93
					_1766_v93 = false
					_1766_v93 = !(!(true))
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1767_v94 _dafny.MultiSet
		_ = _1767_v94
		_1767_v94 = _dafny.MultiSetOf((_this).F0(), p1, p0)
		var _1768_v95 _dafny.Set
		_ = _1768_v95
		_1768_v95 = _dafny.SetOf(_1644_v0)
		r0 = (func() _dafny.Int {
			if (_1767_v94).Contains((p1).Times((_1768_v95).Cardinality())) {
				return (_1767_v94).Multiplicity((p1).Times((_1768_v95).Cardinality()))
			}
			return _dafny.IntOfUint32((_1648_v3).Cardinality())
		})()
		return r0
	}
}
func (_this *C5) F0() _dafny.Int {
	{
		return _this._f0
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C6{}
var _ T1 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__() {
	{
	}
}
func (_this *C6) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(320))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg99 _dafny.Int) interface{} {
				return coer99(arg99)
			}
		}(func(_1769_i0 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(469), true)).Cardinality()
		})), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-916), false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(818))).Cardinality())).Cardinality()), _dafny.IntOfInt64(937))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxumvgapi")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(77))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg100 _dafny.Int) interface{} {
				return coer100(arg100)
			}
		}(func(_1770_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(502)
		})), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kqtm")).Cardinality())))
	}
}
func (_this *C6) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("up"), _dafny.UnicodeSeqOfUtf8Bytes("fgl")), _dafny.UnicodeSeqOfUtf8Bytes("lmdmkvq"))).Cardinality())
	}
}
func (_this *C6) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1771_v0 _dafny.Map
		_ = _1771_v0
		_1771_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(689)).Cmp(p0) != 0, p0)
		var _1772_v1 _dafny.Sequence
		_ = _1772_v1
		_1772_v1 = _dafny.SeqOf(p3)
		var _1773_v2 _dafny.Map
		_ = _1773_v2
		_1773_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _1774_v3 _dafny.Map
		_ = _1774_v3
		_1774_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1772_v1).Cardinality()), (_1773_v2).Cardinality())
		_1771_v0 = (_1771_v0).Update(!(_1774_v3).Contains(p0), p0)
		var _1775_v4 _dafny.Array
		_ = _1775_v4
		var _nw322 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw322
		_1775_v4 = _nw322
		var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw323
		_1775_v4 = _nw323
		if p3 {
			r0 = (p0).Cmp(p0) >= 0
			var _1776_v5 _dafny.Array
			_ = _1776_v5
			var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw324
			_1776_v5 = _nw324
			var _1777_v6 _dafny.Array
			_ = _1777_v6
			var _nw325 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw325
			_1777_v6 = _nw325
			var _1778_v7 _dafny.MultiSet
			_ = _1778_v7
			_1778_v7 = _dafny.MultiSetOf(p2)
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
			_ = _index319
			(_1777_v6).ArraySet1((_1778_v7).IsSubsetOf(_1778_v7), (_index319).Int())
			var _1779_v8 _dafny.Map
			_ = _1779_v8
			_1779_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1776_v5)
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
			_ = _index320
			var _rhs292 _dafny.Array = (func() _dafny.Array {
				if (_1779_v8).Contains(p2) {
					return (_1779_v8).Get(p2).(_dafny.Array)
				}
				return _1776_v5
			})()
			_ = _rhs292
			var _rhs293 bool = Companion_Default___.Fm5(globalState)
			_ = _rhs293
			var _lhs150 _dafny.Array = _1777_v6
			_ = _lhs150
			var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
			_ = _lhs151
			_1776_v5 = _rhs292
			(_lhs150).ArraySet1(_rhs293, (_lhs151).Int())
			var _1780_v9 _dafny.Sequence
			_ = _1780_v9
			_1780_v9 = _dafny.SeqOf(p0, p0, p0, p0)
			var _1781_v10 _dafny.MultiSet
			_ = _1781_v10
			_1781_v10 = _dafny.MultiSetOf(_1780_v9)
			var _1782_v11 _dafny.Map
			_ = _1782_v11
			_1782_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1781_v10, false)
			_1782_v11 = (_1782_v11).Update(_1781_v10, Companion_Default___.Fm5(globalState))
			if p3 {
				r0 = !(false)
				var _1783_v12 _dafny.Array
				_ = _1783_v12
				var _nw326 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw326
				_1783_v12 = _nw326
				var _1784_v13 _dafny.Map
				_ = _1784_v13
				_1784_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("vxtfidplh"))
				var _1785_v14 _dafny.Sequence
				_ = _1785_v14
				_1785_v14 = _dafny.UnicodeSeqOfUtf8Bytes("to")
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1783_v12), 0))
				_ = _index321
				(_1783_v12).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1784_v13).Contains(p0) {
						return (_1784_v13).Get(p0).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("o")
				})(), _1785_v14), (_index321).Int())
				var _1786_v15 _dafny.Array
				_ = _1786_v15
				var _nwElement0_77 _dafny.Array = _1775_v4
				_ = _nwElement0_77
				var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(7))
				_ = _nw327
				(_nw327).ArraySet1(_nwElement0_77, 0)
				(_nw327).ArraySet1(_1775_v4, 1)
				(_nw327).ArraySet1(_1775_v4, 2)
				(_nw327).ArraySet1(_1775_v4, 3)
				(_nw327).ArraySet1(_1775_v4, 4)
				(_nw327).ArraySet1((func() _dafny.Array {
					if p2 {
						return _1775_v4
					}
					return _1775_v4
				})(), 5)
				(_nw327).ArraySet1(_1775_v4, 6)
				_1786_v15 = _nw327
				var _1787_v16 _dafny.Sequence
				_ = _1787_v16
				_1787_v16 = _dafny.SeqOf(_1786_v15, _1786_v15)
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index322
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1783_v12), 0))
				_ = _index323
				var _rhs294 bool = p2
				_ = _rhs294
				var _rhs295 _dafny.Sequence = _1785_v14
				_ = _rhs295
				var _rhs296 _dafny.Array = (_1787_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1787_v16).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs296
				var _lhs152 _dafny.Array = _1777_v6
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _lhs153
				var _lhs154 _dafny.Array = _1783_v12
				_ = _lhs154
				var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1783_v12), 0))
				_ = _lhs155
				(_lhs152).ArraySet1(_rhs294, (_lhs153).Int())
				(_lhs154).ArraySet1(_rhs295, (_lhs155).Int())
				_1786_v15 = _rhs296
				r2 = (_dafny.Zero).Minus((_1780_v9).Select((Companion_Default___.SafeIndex((p0).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1783_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_1783_v12), 0))).Int()).(_dafny.Sequence)).Cardinality()))), _dafny.IntOfUint32((_1780_v9).Cardinality()))).Uint32()).(_dafny.Int))
				var _1788_v17 _dafny.CodePoint
				_ = _1788_v17
				_1788_v17 = _dafny.CodePoint('r')
				var _1789_v18 _dafny.Map
				_ = _1789_v18
				_1789_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1778_v7, !((_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool)))
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index324
				var _rhs297 bool = false
				_ = _rhs297
				var _rhs298 bool = (_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool)
				_ = _rhs298
				var _rhs299 _dafny.CodePoint = _1788_v17
				_ = _rhs299
				var _rhs300 _dafny.Map = _1789_v18
				_ = _rhs300
				var _lhs156 _dafny.Array = _1777_v6
				_ = _lhs156
				var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _lhs157
				r0 = _rhs297
				(_lhs156).ArraySet1(_rhs298, (_lhs157).Int())
				_1788_v17 = _rhs299
				_1789_v18 = _rhs300
				var _1790_v19 *C3
				_ = _1790_v19
				var _nw328 *C3 = New_C3_()
				_ = _nw328
				_nw328.Ctor__((p0).Times(p0))
				_1790_v19 = _nw328
			} else {
				var _1791_v20 _dafny.Map
				_ = _1791_v20
				_1791_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1774_v3, _1775_v4)
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index325
				(_1777_v6).ArraySet1(((_1791_v20).Cardinality()).Cmp(p0) >= 0, (_index325).Int())
				var _rhs301 _dafny.Sequence = _1772_v1
				_ = _rhs301
				var _rhs302 _dafny.Int = (p0).Plus((_dafny.Zero).Minus(p0))
				_ = _rhs302
				_1772_v1 = _rhs301
				r2 = _rhs302
				var _1792_v21 _dafny.Array
				_ = _1792_v21
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_37
				var _nw329 _dafny.Array
				_ = _nw329
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw329 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.CodePoint = (func(_1793_p2 bool) func(_dafny.Int) _dafny.CodePoint {
						return func(_1794_i0 _dafny.Int) _dafny.CodePoint {
							return (func() _dafny.CodePoint {
								if _1793_p2 {
									return _dafny.CodePoint('c')
								}
								return _dafny.CodePoint('u')
							})()
						}
					})(p2)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw329 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw329).ArraySet1CodePoint(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw329).ArraySet1CodePoint(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1792_v21 = _nw329
				var _1795_v22 _dafny.CodePoint
				_ = _1795_v22
				_1795_v22 = _dafny.CodePoint('l')
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_1792_v21), 0))
				_ = _index326
				(_1792_v21).ArraySet1CodePoint(_1795_v22, (_index326).Int())
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_1792_v21), 0))
				_ = _index327
				(_1792_v21).ArraySet1CodePoint(_1795_v22, (_index327).Int())
				var _1796_v23 _dafny.MultiSet
				_ = _1796_v23
				_1796_v23 = _dafny.MultiSetOf(p0, p0, p0, p0)
				var _1797_v24 _dafny.MultiSet
				_ = _1797_v24
				_1797_v24 = _1796_v23
				var _1798_v25 _dafny.Array
				_ = _1798_v25
				var _nwElement0_78 _dafny.MultiSet = Companion_Default___.Fm19(p0, p0, (_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool), p0, globalState)
				_ = _nwElement0_78
				var _nw330 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(13))
				_ = _nw330
				(_nw330).ArraySet1(_nwElement0_78, 0)
				(_nw330).ArraySet1(_1796_v23, 1)
				(_nw330).ArraySet1(_1796_v23, 2)
				(_nw330).ArraySet1(_1797_v24, 3)
				(_nw330).ArraySet1(_1797_v24, 4)
				(_nw330).ArraySet1(_1797_v24, 5)
				(_nw330).ArraySet1(_1797_v24, 6)
				(_nw330).ArraySet1(_1797_v24, 7)
				(_nw330).ArraySet1(_1796_v23, 8)
				(_nw330).ArraySet1(_1796_v23, 9)
				(_nw330).ArraySet1(_1797_v24, 10)
				(_nw330).ArraySet1(_1796_v23, 11)
				(_nw330).ArraySet1((func() _dafny.MultiSet {
					if p3 {
						return _1797_v24
					}
					return _1797_v24
				})(), 12)
				_1798_v25 = _nw330
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_1798_v25), 0))
				_ = _index328
				(_1798_v25).ArraySet1(_1797_v24, (_index328).Int())
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_1798_v25), 0))
				_ = _index329
				(_1798_v25).ArraySet1(_1796_v23, (_index329).Int())
				r0 = p2
			}
			if (_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool) {
				var _1799_v26 *C4
				_ = _1799_v26
				var _nw331 *C4 = New_C4_()
				_ = _nw331
				_nw331.Ctor__((_dafny.Zero).Minus(p0))
				_1799_v26 = _nw331
				var _1800_v27 *C5
				_ = _1800_v27
				var _nw332 *C5 = New_C5_()
				_ = _nw332
				_nw332.Ctor__(p0, _1777_v6)
				_1800_v27 = _nw332
				var _1801_v28 D14
				_ = _1801_v28
				_1801_v28 = Companion_D14_.Create_DC25_(_1776_v5)
				_1776_v5 = (_1801_v28).Dtor_cf42()
				var _1802_v29 *C4
				_ = _1802_v29
				var _nw333 *C4 = New_C4_()
				_ = _nw333
				_nw333.Ctor__((_1800_v27).F0())
				_1802_v29 = _nw333
				var _1803_v30 _dafny.Map
				_ = _1803_v30
				_1803_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1773_v2).Merge(_1773_v2), p2)
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index330
				var _rhs303 _dafny.Array = _1775_v4
				_ = _rhs303
				var _rhs304 bool = (_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool)
				_ = _rhs304
				var _rhs305 _dafny.Int = (_dafny.Zero).Minus((_1799_v26).Fm1(p3, globalState))
				_ = _rhs305
				var _rhs306 _dafny.Map = _1803_v30
				_ = _rhs306
				var _rhs307 bool = (_1777_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))).Int()).(bool)
				_ = _rhs307
				var _lhs158 _dafny.Array = _1777_v6
				_ = _lhs158
				var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _lhs159
				_1775_v4 = _rhs303
				r0 = _rhs304
				r2 = _rhs305
				_1803_v30 = _rhs306
				(_lhs158).ArraySet1(_rhs307, (_lhs159).Int())
			} else {
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index331
				(_1777_v6).ArraySet1(!(!(!(p2))), (_index331).Int())
				r2 = p0
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1777_v6), 0))
				_ = _index332
				(_1777_v6).ArraySet1(p2, (_index332).Int())
				var _1804_v31 _dafny.Sequence
				_ = _1804_v31
				var _1805_v32 _dafny.Int
				_ = _1805_v32
				var _1806_v33 _dafny.Int
				_ = _1806_v33
				var _1807_v34 bool
				_ = _1807_v34
				var _out28 _dafny.Sequence
				_ = _out28
				var _out29 _dafny.Int
				_ = _out29
				var _out30 _dafny.Int
				_ = _out30
				var _out31 bool
				_ = _out31
				_out28, _out29, _out30, _out31 = (_this).M4(true, (Companion_Default___.Fm42(p0, p2, p3, globalState)).Dtor_cf21(), p0, p3, globalState)
				_1804_v31 = _out28
				_1805_v32 = _out29
				_1806_v33 = _out30
				_1807_v34 = _out31
				var _1808_v35 _dafny.Sequence
				_ = _1808_v35
				_1808_v35 = _dafny.UnicodeSeqOfUtf8Bytes("bevnw")
				_1808_v35 = _1808_v35
			}
		} else {
			r0 = (p0).Cmp(p0) <= 0
			var _1809_v36 _dafny.Array
			_ = _1809_v36
			var _nwElement0_79 bool = true
			_ = _nwElement0_79
			var _nw334 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(2))
			_ = _nw334
			(_nw334).ArraySet1(_nwElement0_79, 0)
			(_nw334).ArraySet1(p2, 1)
			_1809_v36 = _nw334
			var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))
			_ = _index333
			(_1809_v36).ArraySet1(p3, (_index333).Int())
			var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))
			_ = _index334
			(_1809_v36).ArraySet1(p3, (_index334).Int())
			r2 = p0
			var _1810_v37 _dafny.CodePoint
			_ = _1810_v37
			_1810_v37 = _dafny.CodePoint('a')
			var _1811_v38 D0
			_ = _1811_v38
			_1811_v38 = Companion_D0_.Create_DC1_(!((_1809_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))).Int()).(bool)), (_1809_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))).Int()).(bool), _1810_v37, (p0).Plus(p0))
			_1811_v38 = _1811_v38
			var _1812_v39 _dafny.Map
			_ = _1812_v39
			_1812_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1774_v3)
			var _1813_v40 _dafny.Map
			_ = _1813_v40
			_1813_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1809_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))).Int()).(bool), (_1774_v3).Merge((func() _dafny.Map {
				if (_1812_v39).Contains((_1809_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))).Int()).(bool)) {
					return (_1812_v39).Get((_1809_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1809_v36), 0))).Int()).(bool)).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm38(p3, globalState)).Cardinality(), p0)
			})()))
			_1813_v40 = (_1813_v40).Update(Companion_Default___.Fm5(globalState), _1774_v3)
		}
		var _hi12 _dafny.Int = _dafny.IntOfInt64(575)
		_ = _hi12
		for _1814_i1 := (_dafny.Zero).Minus(p0); _1814_i1.Cmp(_hi12) < 0; _1814_i1 = _1814_i1.Plus(_dafny.One) {
			var _1815_v41 _dafny.CodePoint
			_ = _1815_v41
			_1815_v41 = _dafny.CodePoint('r')
			var _rhs308 bool = !(!((Companion_Default___.Fm45(p0, true, p2, _1814_i1, globalState)).Dtor_cf39()))
			_ = _rhs308
			var _rhs309 bool = !(!((true) == (p3)))
			_ = _rhs309
			var _rhs310 _dafny.CodePoint = _1815_v41
			_ = _rhs310
			var _rhs311 _dafny.Int = _1814_i1
			_ = _rhs311
			r0 = _rhs308
			r0 = _rhs309
			_1815_v41 = _rhs310
			r2 = _rhs311
			r0 = !(p2)
			r2 = p0
			var _1816_v42 _dafny.Set
			_ = _1816_v42
			_1816_v42 = _dafny.SetOf(p2)
			var _1817_v43 D0
			_ = _1817_v43
			_1817_v43 = Companion_D0_.Create_DC0_(p3)
			var _1818_v44 D14
			_ = _1818_v44
			_1818_v44 = Companion_D14_.Create_DC26_(_1817_v43, _1816_v42)
			var _1819_v45 _dafny.Set
			_ = _1819_v45
			_1819_v45 = _dafny.SetOf((_1816_v42).Intersection(_1816_v42), (_1818_v44).Dtor_cf44(), _1816_v42, _dafny.SetOf(false))
			_1819_v45 = _1819_v45
		}
		var _1820_i2 _dafny.Int
		_ = _1820_i2
		_1820_i2 = _dafny.Zero
		{
			for Companion_Default___.Fm5(globalState) {
				{
					if (_1820_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1820_i2 = (_1820_i2).Plus(_dafny.One)
					r2 = p0
					var _1821_v46 _dafny.Sequence
					_ = _1821_v46
					_1821_v46 = _dafny.SeqOf(_dafny.IntOfInt64(-589))
					r0 = ((_1821_v46).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1821_v46).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(p0) == 0
					r2 = Companion_Default___.SafeDivisionInt(((_this).Fm1(p2, globalState)).Times(p0), p0)
					var _1822_v47 _dafny.MultiSet
					_ = _1822_v47
					_1822_v47 = _dafny.MultiSetOf(_1775_v4, _1775_v4)
					var _1823_v48 *C0
					_ = _1823_v48
					var _nw335 *C0 = New_C0_()
					_ = _nw335
					_nw335.Ctor__((_dafny.MultiSetOf(_1775_v4, _1775_v4)).IsProperSubsetOf(_1822_v47), p2)
					_1823_v48 = _nw335
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		if (p0).Cmp(p0) >= 0 {
			var _1824_v49 _dafny.Map
			_ = _1824_v49
			_1824_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm5(globalState)), !(p3))
			var _1825_v50 _dafny.Sequence
			_ = _1825_v50
			_1825_v50 = _dafny.UnicodeSeqOfUtf8Bytes("y")
			r0 = (func() bool {
				if p2 {
					return (func() bool {
						if (_1824_v49).Contains(!(!(false))) {
							return (_1824_v49).Get(!(!(false))).(bool)
						}
						return p3
					})()
				}
				return _dafny.Companion_Sequence_.IsPrefixOf(_1825_v50, _1825_v50)
			})()
			var _1826_v51 _dafny.Set
			_ = _1826_v51
			_1826_v51 = _dafny.SetOf(false, p3)
			r0 = ((_1826_v51).Union(_1826_v51)).IsSubsetOf((_dafny.SetOf(p3, p2, p2)).Difference(_dafny.SetOf(true, p2, p2)))
			if !(_dafny.Companion_Sequence_.Equal(_1772_v1, _1772_v1)) || (!(p3)) {
				var _1827_v52 T1
				_ = _1827_v52
				var _nw336 *C1 = New_C1_()
				_ = _nw336
				_nw336.Ctor__(_1825_v50)
				_1827_v52 = _nw336
				_1827_v52 = _1827_v52
				r2 = p0
				var _1828_v53 _dafny.Sequence
				_ = _1828_v53
				var _1829_v54 _dafny.Int
				_ = _1829_v54
				var _1830_v55 _dafny.Int
				_ = _1830_v55
				var _1831_v56 bool
				_ = _1831_v56
				var _out32 _dafny.Sequence
				_ = _out32
				var _out33 _dafny.Int
				_ = _out33
				var _out34 _dafny.Int
				_ = _out34
				var _out35 bool
				_ = _out35
				_out32, _out33, _out34, _out35 = (_this).M4(p2, _1825_v50, (_this).Fm1(p3, globalState), !((func() bool {
					if p3 {
						return false
					}
					return p3
				})()), globalState)
				_1828_v53 = _out32
				_1829_v54 = _out33
				_1830_v55 = _out34
				_1831_v56 = _out35
				_1831_v56 = _1831_v56
				var _1832_v57 _dafny.Map
				_ = _1832_v57
				_1832_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1772_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1831_v56, _1775_v4)).Cardinality())
				var _1833_v58 _dafny.CodePoint
				_ = _1833_v58
				_1833_v58 = _dafny.CodePoint('h')
				var _1834_v59 _dafny.Set
				_ = _1834_v59
				_1834_v59 = _dafny.SetOf(p0, _dafny.IntOfUint32((_1772_v1).Cardinality()), _1830_v55)
				var _1835_v60 _dafny.Array
				_ = _1835_v60
				var _nwElement0_80 bool = (_1832_v57).Equals(_1832_v57)
				_ = _nwElement0_80
				var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(25))
				_ = _nw337
				(_nw337).ArraySet1(_nwElement0_80, 0)
				(_nw337).ArraySet1(!(p2) || (p3), 1)
				(_nw337).ArraySet1(_1831_v56, 2)
				(_nw337).ArraySet1((func() bool {
					if p3 {
						return p2
					}
					return _1831_v56
				})(), 3)
				(_nw337).ArraySet1((!(p2)) == (false), 4)
				(_nw337).ArraySet1((func() bool {
					if !(p2) {
						return p2
					}
					return Companion_Default___.Fm5(globalState)
				})(), 5)
				(_nw337).ArraySet1(_1831_v56, 6)
				(_nw337).ArraySet1(!(_1831_v56), 7)
				(_nw337).ArraySet1(p2, 8)
				(_nw337).ArraySet1((func() bool {
					if (_1824_v49).Contains(p2) {
						return (_1824_v49).Get(p2).(bool)
					}
					return p3
				})(), 9)
				(_nw337).ArraySet1(p2, 10)
				(_nw337).ArraySet1(_1831_v56, 11)
				(_nw337).ArraySet1(_1831_v56, 12)
				(_nw337).ArraySet1(p3, 13)
				(_nw337).ArraySet1(p3, 14)
				(_nw337).ArraySet1(_1831_v56, 15)
				(_nw337).ArraySet1(p3, 16)
				(_nw337).ArraySet1(p3, 17)
				(_nw337).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1825_v50, _1833_v58), 18)
				(_nw337).ArraySet1((_dafny.SetOf(_1830_v55)).IsDisjointFrom(_1834_v59), 19)
				(_nw337).ArraySet1(p2, 20)
				(_nw337).ArraySet1(!(p2) || (p2), 21)
				(_nw337).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_1772_v1, p3)), 22)
				(_nw337).ArraySet1(_1831_v56, 23)
				(_nw337).ArraySet1(p3, 24)
				_1835_v60 = _nw337
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1835_v60), 0))
				_ = _index335
				(_1835_v60).ArraySet1(true, (_index335).Int())
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1835_v60), 0))
				_ = _index336
				(_1835_v60).ArraySet1((_1826_v51).IsSubsetOf(_1826_v51), (_index336).Int())
			} else {
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index337
				(_1775_v4).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(212), p0), (_index337).Int())
				var _1836_v61 _dafny.Sequence
				_ = _1836_v61
				_1836_v61 = _dafny.SeqOf(_1826_v51)
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index338
				(_1775_v4).ArraySet1((func() _dafny.Int {
					if p2 {
						return (func() _dafny.Int {
							if p2 {
								return p0
							}
							return p0
						})()
					}
					return (func() _dafny.Int {
						if p3 {
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(404))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg101 _dafny.Int) interface{} {
									return coer101(arg101)
								}
							}((func(_1837_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1838_i3 _dafny.Int) _dafny.Int {
									return _1837_p0
								}
							})(p0)))).Cardinality()))
						}
						return ((_1836_v61).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1836_v61).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
					})()
				})(), (_index338).Int())
				var _1839_v62 _dafny.Array
				_ = _1839_v62
				var _out36 _dafny.Array
				_ = _out36
				_out36 = Companion_Default___.M0(p2, p2, globalState)
				_1839_v62 = _out36
				_1774_v3 = (_1774_v3).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-585), (_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)), p0)
				var _1840_v63 _dafny.Array
				_ = _1840_v63
				var _nw338 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw338
				_1840_v63 = _nw338
				var _1841_v64 _dafny.Map
				_ = _1841_v64
				_1841_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1840_v63, !(p3))
				_1841_v64 = (_1841_v64).Update(_1840_v63, false)
				var _1842_v65 _dafny.CodePoint
				_ = _1842_v65
				_1842_v65 = _dafny.CodePoint('e')
				_1842_v65 = _1842_v65
			}
			r2 = (func() _dafny.Int {
				if p3 {
					return (_dafny.Zero).Minus(p0)
				}
				return p0
			})()
			var _1843_v66 _dafny.MultiSet
			_ = _1843_v66
			_1843_v66 = _dafny.MultiSetOf(Companion_Default___.Fm5(globalState))
			if (_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(767))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}(func(_1844_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), _dafny.CodePoint('q'))) || (((func() _dafny.Int {
				if (_1843_v66).Contains(p2) {
					return (_1843_v66).Multiplicity(p2)
				}
				return p0
			})()).Cmp(p0) < 0) {
				var _1845_v67 D11
				_ = _1845_v67
				_1845_v67 = Companion_D11_.Create_DC21_(_dafny.CodePoint('s'), _dafny.IntOfInt64(65), p0)
				var _1846_v68 _dafny.CodePoint
				_ = _1846_v68
				_1846_v68 = _dafny.CodePoint('i')
				var _1847_v70 _dafny.Set
				_ = _1847_v70
				_1847_v70 = _dafny.SetOf((_this).Fm1(p3, globalState), p0, p0, p0)
				var _pat_let_tv36 = _1846_v68
				_ = _pat_let_tv36
				var _1848_v71 _dafny.Array
				_ = _1848_v71
				var _nwElement0_81 D11 = _1845_v67
				_ = _nwElement0_81
				var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(21))
				_ = _nw339
				(_nw339).ArraySet1(_nwElement0_81, 0)
				(_nw339).ArraySet1(_1845_v67, 1)
				(_nw339).ArraySet1((func() D11 {
					if p2 {
						return _1845_v67
					}
					return _1845_v67
				})(), 2)
				(_nw339).ArraySet1(_1845_v67, 3)
				(_nw339).ArraySet1(_1845_v67, 4)
				(_nw339).ArraySet1((func() D11 {
					if p2 {
						return func(_pat_let39_0 D11) D11 {
							return func(_1849_dt__update__tmp_h2 D11) D11 {
								return func(_pat_let40_0 _dafny.CodePoint) D11 {
									return func(_1850_dt__update_hcf36_h0 _dafny.CodePoint) D11 {
										return Companion_D11_.Create_DC21_(_1850_dt__update_hcf36_h0, (_1849_dt__update__tmp_h2).Dtor_cf37(), (_1849_dt__update__tmp_h2).Dtor_cf38())
									}(_pat_let40_0)
								}(_pat_let_tv36)
							}(_pat_let39_0)
						}(_1845_v67)
					}
					return Companion_D11_.Create_DC21_(_1846_v68, p0, p0)
				})(), 5)
				(_nw339).ArraySet1(_1845_v67, 6)
				(_nw339).ArraySet1(_1845_v67, 7)
				(_nw339).ArraySet1(_1845_v67, 8)
				(_nw339).ArraySet1(Companion_D11_.Create_DC21_(_1846_v68, p0, p0), 9)
				(_nw339).ArraySet1(_1845_v67, 10)
				(_nw339).ArraySet1(_1845_v67, 11)
				(_nw339).ArraySet1(_1845_v67, 12)
				(_nw339).ArraySet1(Companion_Default___.Fm46(func() _dafny.Set {
					var _coll62 = _dafny.NewBuilder()
					_ = _coll62
					for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(424), _dafny.IntOfInt64(-760))); ; {
						_compr_62, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1851_v69 _dafny.Int
						_1851_v69 = interface{}(_compr_62).(_dafny.Int)
						if ((_dafny.IntOfInt64(424)).Cmp(_1851_v69) <= 0) && ((_1851_v69).Cmp(_dafny.IntOfInt64(-760)) < 0) {
							_coll62.Add(Companion_Default___.SafeModuloInt(_1851_v69, (_dafny.Zero).Minus(p0)))
						}
					}
					return _coll62.ToSet()
				}(), globalState), 13)
				(_nw339).ArraySet1(Companion_D11_.Create_DC21_(_1846_v68, p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1772_v1).Cardinality()))).Cardinality(), _dafny.IntOfInt64(836))).Cardinality()), 14)
				(_nw339).ArraySet1(Companion_D11_.Create_DC21_(_dafny.CodePoint('o'), p0, p0), 15)
				(_nw339).ArraySet1((func() D11 {
					if p3 {
						return _1845_v67
					}
					return _1845_v67
				})(), 16)
				(_nw339).ArraySet1(_1845_v67, 17)
				(_nw339).ArraySet1(_1845_v67, 18)
				(_nw339).ArraySet1(Companion_D11_.Create_DC21_(_1846_v68, p0, p0), 19)
				(_nw339).ArraySet1(Companion_Default___.Fm46(_1847_v70, globalState), 20)
				_1848_v71 = _nw339
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1848_v71), 0))
				_ = _index339
				(_1848_v71).ArraySet1(_1845_v67, (_index339).Int())
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index340
				(_1775_v4).ArraySet1(p0, (_index340).Int())
				var _1852_v72 D4
				_ = _1852_v72
				_1852_v72 = Companion_D4_.Create_DC6_(p0, _1775_v4, p0)
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1848_v71), 0))
				_ = _index341
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index342
				var _rhs312 D11 = _1845_v67
				_ = _rhs312
				var _rhs313 _dafny.Int = Companion_Default___.SafeModuloInt((_1852_v72).Dtor_cf9(), (_1774_v3).Cardinality())
				_ = _rhs313
				var _lhs160 _dafny.Array = _1848_v71
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1848_v71), 0))
				_ = _lhs161
				var _lhs162 _dafny.Array = _1775_v4
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1775_v4), 0))
				_ = _lhs163
				(_lhs160).ArraySet1(_rhs312, (_lhs161).Int())
				(_lhs162).ArraySet1(_rhs313, (_lhs163).Int())
				var _1853_v73 _dafny.Map
				_ = _1853_v73
				_1853_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1773_v2).Merge(_1773_v2), _dafny.IntOfUint32((_1825_v50).Cardinality()))
				_1853_v73 = (_1853_v73).Update(_1773_v2, _dafny.IntOfInt64(161))
				var _1854_v74 _dafny.Array
				_ = _1854_v74
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_38
				var _nw340 _dafny.Array
				_ = _nw340
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw340 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.CodePoint = (func(_1855_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1856_i5 _dafny.Int) _dafny.CodePoint {
							return _1855_v68
						}
					})(_1846_v68)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw340 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw340).ArraySet1CodePoint(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw340).ArraySet1CodePoint(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1854_v74 = _nw340
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_1854_v74), 0))
				_ = _index343
				(_1854_v74).ArraySet1CodePoint(Companion_Default___.Fm23(globalState), (_index343).Int())
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_1854_v74), 0))
				_ = _index344
				(_1854_v74).ArraySet1CodePoint(Companion_Default___.Fm37((_dafny.IntOfUint32((_1825_v50).Cardinality())).Plus((_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)), globalState), (_index344).Int())
				r0 = (p0).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_1825_v50).Cardinality()))) <= 0
				var _1857_v75 _dafny.Array
				_ = _1857_v75
				var _nw341 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
				_ = _nw341
				_1857_v75 = _nw341
				_1857_v75 = (func() _dafny.Array {
					if !(p2) {
						return _1857_v75
					}
					return _1857_v75
				})()
			} else {
				var _1858_v76 _dafny.Sequence
				_ = _1858_v76
				_1858_v76 = _dafny.SeqOf(p0)
				r0 = (Companion_Default___.Fm41(_1858_v76, p2, p2, p2, globalState)).IsSubsetOf(_1826_v51)
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((p1), 0))
				_ = _index345
				(p1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1858_v76, _dafny.SeqOf(p0)), (_index345).Int())
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((p1), 0))
				_ = _index346
				(p1).ArraySet1(_1858_v76, (_index346).Int())
				r2 = (_dafny.IntOfInt64(131)).Minus(p0)
				_1771_v0 = (_1771_v0).Update((func() bool {
					if (_1773_v2).Contains(p0) {
						return (_1773_v2).Get(p0).(bool)
					}
					return p2
				})(), p0)
				r0 = !(p2)
			}
		} else {
			r2 = (_dafny.Zero).Minus(p0)
			r2 = (_this).Fm1(p2, globalState)
			var _1859_v77 D0
			_ = _1859_v77
			_1859_v77 = Companion_D0_.Create_DC0_(p2)
			var _1860_v78 _dafny.Set
			_ = _1860_v78
			_1860_v78 = _dafny.SetOf(p2)
			var _1861_v79 D14
			_ = _1861_v79
			_1861_v79 = Companion_D14_.Create_DC26_(_1859_v77, _1860_v78)
			var _1862_v80 _dafny.Set
			_ = _1862_v80
			_1862_v80 = _dafny.SetOf(_1861_v79, _1861_v79)
			var _1863_v81 _dafny.Map
			_ = _1863_v81
			_1863_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1862_v80)
			if (_1862_v80).IsSubsetOf((func() _dafny.Set {
				if (_1863_v81).Contains(!(false)) {
					return (_1863_v81).Get(!(false)).(_dafny.Set)
				}
				return _1862_v80
			})()) {
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index347
				(_1775_v4).ArraySet1((p0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mig")).Cardinality())), (_index347).Int())
				var _1864_v82 _dafny.Map
				_ = _1864_v82
				_1864_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1775_v4, p0)
				var _1865_v83 _dafny.Set
				_ = _1865_v83
				_1865_v83 = _dafny.SetOf(p0)
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index348
				var _rhs314 _dafny.Int = ((_this).Fm1(p2, globalState)).Minus(p0)
				_ = _rhs314
				var _rhs315 _dafny.Int = ((_1864_v82).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1775_v4, p0)).Update(_1775_v4, (_1865_v83).Cardinality()))).Cardinality()
				_ = _rhs315
				var _rhs316 bool = p3
				_ = _rhs316
				var _rhs317 _dafny.Int = p0
				_ = _rhs317
				var _lhs164 _dafny.Array = _1775_v4
				_ = _lhs164
				var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))
				_ = _lhs165
				r2 = _rhs314
				(_lhs164).ArraySet1(_rhs315, (_lhs165).Int())
				r0 = _rhs316
				r2 = _rhs317
				var _1866_v84 _dafny.Sequence
				_ = _1866_v84
				_1866_v84 = _dafny.UnicodeSeqOfUtf8Bytes("kaqatlgk")
				var _1867_v85 _dafny.Sequence
				_ = _1867_v85
				_1867_v85 = _dafny.SeqOf((_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(292))).Cardinality()), ((_dafny.Zero).Minus(p0)).Times(p0), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1866_v84).Cardinality()), _dafny.IntOfInt64(785)))
				_1867_v85 = _dafny.SeqOf(((_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)).Minus((_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)), p0, p0)
				var _1868_v86 *C0
				_ = _1868_v86
				var _nw342 *C0 = New_C0_()
				_ = _nw342
				_nw342.Ctor__(p3, p3)
				_1868_v86 = _nw342
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index349
				(_1775_v4).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if _1868_v86.F5 {
						return (_this).Fm1(_1868_v86.F5, globalState)
					}
					return (_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)
				})(), (_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int)), (_index349).Int())
				var _1869_v87 _dafny.Array
				_ = _1869_v87
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
				_ = _nw343
				_1869_v87 = _nw343
				_1869_v87 = _1869_v87
			} else {
				var _1870_v88 _dafny.Array
				_ = _1870_v88
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_39
				var _nw344 _dafny.Array
				_ = _nw344
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw344 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) bool = (func(_1871_p3 bool) func(_dafny.Int) bool {
						return func(_1872_i6 _dafny.Int) bool {
							return _1871_p3
						}
					})(p3)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw344 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw344).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw344).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1870_v88 = _nw344
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _index350
				(_1870_v88).ArraySet1(p3, (_index350).Int())
				var _1873_v89 _dafny.Map
				_ = _1873_v89
				_1873_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1860_v78, p2)
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _index351
				(_1870_v88).ArraySet1(!((func() bool {
					if (_1873_v89).Contains((_dafny.SetOf(p2)).Union(_1860_v78)) {
						return (_1873_v89).Get((_dafny.SetOf(p2)).Union(_1860_v78)).(bool)
					}
					return !(Companion_Default___.Fm5(globalState))
				})()), (_index351).Int())
				var _1874_v90 _dafny.Array
				_ = _1874_v90
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_40
				var _nw345 _dafny.Array
				_ = _nw345
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw345 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.CodePoint = func(_1875_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('i')
					}
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw345 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw345).ArraySet1CodePoint(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw345).ArraySet1CodePoint(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1874_v90 = _nw345
				var _1876_v91 _dafny.CodePoint
				_ = _1876_v91
				_1876_v91 = _dafny.CodePoint('w')
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_1874_v90), 0))
				_ = _index352
				(_1874_v90).ArraySet1CodePoint(_1876_v91, (_index352).Int())
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_1874_v90), 0))
				_ = _index353
				(_1874_v90).ArraySet1CodePoint(_1876_v91, (_index353).Int())
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _index354
				(_1870_v88).ArraySet1(p3, (_index354).Int())
				var _1877_v92 _dafny.MultiSet
				_ = _1877_v92
				_1877_v92 = _dafny.MultiSetOf(!(p3))
				var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _index355
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _index356
				var _rhs318 bool = Companion_Default___.Fm5(globalState)
				_ = _rhs318
				var _rhs319 bool = (p0).Cmp((p0).Plus(_dafny.IntOfInt64(144))) <= 0
				_ = _rhs319
				var _rhs320 bool = (_1877_v92).Equals(_1877_v92)
				_ = _rhs320
				var _lhs166 _dafny.Array = _1870_v88
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _lhs167
				var _lhs168 _dafny.Array = _1870_v88
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))
				_ = _lhs169
				(_lhs166).ArraySet1(_rhs318, (_lhs167).Int())
				r0 = _rhs319
				(_lhs168).ArraySet1(_rhs320, (_lhs169).Int())
				var _1878_v93 D6
				_ = _1878_v93
				_1878_v93 = Companion_D6_.Create_DC13_(p2, p0)
				r0 = ((_1878_v93).Dtor_cf24()) && ((_1870_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1870_v88), 0))).Int()).(bool))
			}
			var _1879_v94 _dafny.Array
			_ = _1879_v94
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_41
			var _nw346 _dafny.Array
			_ = _nw346
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw346 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) _dafny.Set = (func(_1880_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_1881_i8 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_p0, _dafny.CodePoint('o'))).Cardinality(), _1880_p0, _1880_p0, _1880_p0, _1880_p0)).Intersection(_dafny.SetOf(_1880_p0))
					}
				})(p0)
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw346 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw346).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw346).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1879_v94 = _nw346
			_1879_v94 = _1879_v94
			if p2 {
				r0 = Companion_Default___.Fm5(globalState)
				r2 = p0
				var _1882_v97 D14
				_ = _1882_v97
				_1882_v97 = Companion_D14_.Create_DC27_(false)
				var _1883_v98 _dafny.Map
				_ = _1883_v98
				_1883_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1771_v0).Contains(p3) {
						return (_1771_v0).Get(p3).(_dafny.Int)
					}
					return p0
				})(), _1882_v97)
				var _1884_v100 _dafny.Array
				_ = _1884_v100
				var _nwElement0_82 _dafny.Map = _1774_v3
				_ = _nwElement0_82
				var _nw347 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(22))
				_ = _nw347
				(_nw347).ArraySet1(_nwElement0_82, 0)
				(_nw347).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}((func(_1885_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1886_i9 _dafny.Int) _dafny.Int {
						return _1885_p0
					}
				})(p0)))).Cardinality()), p0), 1)
				(_nw347).ArraySet1(_1774_v3, 2)
				(_nw347).ArraySet1(func() _dafny.Map {
					var _coll63 = _dafny.NewMapBuilder()
					_ = _coll63
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(696), _dafny.IntOfInt64(369))); ; {
						_compr_63, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1887_v95 _dafny.Int
						_1887_v95 = interface{}(_compr_63).(_dafny.Int)
						if ((_dafny.IntOfInt64(696)).Cmp(_1887_v95) <= 0) && ((_1887_v95).Cmp(_dafny.IntOfInt64(369)) < 0) {
							_coll63.Add(Companion_Default___.SafeDivisionInt(_1887_v95, (_1860_v78).Cardinality()), p0)
						}
					}
					return _coll63.ToMap()
				}(), 3)
				(_nw347).ArraySet1((_1774_v3).Update(_dafny.IntOfInt64(-183), p0), 4)
				(_nw347).ArraySet1(_1774_v3, 5)
				(_nw347).ArraySet1(_1774_v3, 6)
				(_nw347).ArraySet1(_1774_v3, 7)
				(_nw347).ArraySet1(_1774_v3, 8)
				(_nw347).ArraySet1(_1774_v3, 9)
				(_nw347).ArraySet1((_1774_v3).Update(p0, p0), 10)
				(_nw347).ArraySet1(func() _dafny.Map {
					var _coll64 = _dafny.NewMapBuilder()
					_ = _coll64
					for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(427), _dafny.IntOfInt64(-816))); ; {
						_compr_64, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _1888_v96 _dafny.Int
						_1888_v96 = interface{}(_compr_64).(_dafny.Int)
						if ((_dafny.IntOfInt64(427)).Cmp(_1888_v96) <= 0) && ((_1888_v96).Cmp(_dafny.IntOfInt64(-816)) < 0) {
							_coll64.Add((_1888_v96).Times((_dafny.Zero).Minus(p0)), p0)
						}
					}
					return _coll64.ToMap()
				}(), 11)
				(_nw347).ArraySet1(_1774_v3, 12)
				(_nw347).ArraySet1(_1774_v3, 13)
				(_nw347).ArraySet1(_1774_v3, 14)
				(_nw347).ArraySet1(Companion_Default___.Fm27(p0, p3, globalState), 15)
				(_nw347).ArraySet1(_1774_v3, 16)
				(_nw347).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1883_v98).Cardinality(), (_1860_v78).Cardinality()), 17)
				(_nw347).ArraySet1(func() _dafny.Map {
					var _coll65 = _dafny.NewMapBuilder()
					_ = _coll65
					for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(876), _dafny.IntOfInt64(548))); ; {
						_compr_65, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _1889_v99 _dafny.Int
						_1889_v99 = interface{}(_compr_65).(_dafny.Int)
						if ((_dafny.IntOfInt64(876)).Cmp(_1889_v99) <= 0) && ((_1889_v99).Cmp(_dafny.IntOfInt64(548)) < 0) {
							_coll65.Add((_1889_v99).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('f'))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gg")).Cardinality()))
						}
					}
					return _coll65.ToMap()
				}(), 18)
				(_nw347).ArraySet1(_1774_v3, 19)
				(_nw347).ArraySet1(_1774_v3, 20)
				(_nw347).ArraySet1(_1774_v3, 21)
				_1884_v100 = _nw347
				var _1890_v101 _dafny.Array
				_ = _1890_v101
				_1890_v101 = _1884_v100
				_1890_v101 = _1890_v101
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index357
				(_1775_v4).ArraySet1(p0, (_index357).Int())
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index358
				(_1775_v4).ArraySet1(p0, (_index358).Int())
				var _1891_v102 _dafny.Set
				_ = _1891_v102
				_1891_v102 = _dafny.SetOf((_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int))
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index359
				(_1775_v4).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_dafny.MultiSetOf(_1891_v102, _1891_v102)).Cardinality()), (_index359).Int())
			} else {
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index360
				(_1775_v4).ArraySet1((func() _dafny.Int {
					if p2 {
						return p0
					}
					return _dafny.IntOfInt64(835)
				})(), (_index360).Int())
				var _1892_v103 _dafny.Sequence
				_ = _1892_v103
				_1892_v103 = _dafny.SeqOf(p0)
				var _1893_v104 _dafny.Map
				_ = _1893_v104
				_1893_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1892_v103, p3)
				var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index361
				(_1775_v4).ArraySet1((_1893_v104).Cardinality(), (_index361).Int())
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1775_v4), 0))
				_ = _index362
				(_1775_v4).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, (_1775_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1775_v4), 0))).Int()).(_dafny.Int))), (_index362).Int())
				var _1894_v105 _dafny.Array
				_ = _1894_v105
				var _nw348 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw348
				_1894_v105 = _nw348
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_42
				var _nw349 _dafny.Array
				_ = _nw349
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw349 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Int = (func(_1895_v4 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_1896_i10 _dafny.Int) _dafny.Int {
							return (_1896_i10).Minus((_1895_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1895_v4), 0))).Int()).(_dafny.Int))
						}
					})(_1775_v4)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw349 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw349).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw349).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1894_v105 = _nw349
				var _1897_v106 _dafny.MultiSet
				_ = _1897_v106
				_1897_v106 = _dafny.MultiSetOf(p3, !(p2))
				_1897_v106 = Companion_Default___.Fm38((_dafny.SetOf((_1897_v106).Cardinality())).IsSubsetOf(_dafny.SetOf(p0)), globalState)
				r0 = p3
			}
		}
		var _1898_v107 D6
		_ = _1898_v107
		_1898_v107 = Companion_D6_.Create_DC13_(false, p0)
		var _pat_let_tv37 = p2
		_ = _pat_let_tv37
		var _pat_let_tv38 = p2
		_ = _pat_let_tv38
		r0 = !(func(_source25 D6) bool {
			if _source25.Is_DC12() {
				var _1899___mcc_h0 bool = _source25.Get_().(D6_DC12).Cf22
				_ = _1899___mcc_h0
				var _1900___mcc_h1 bool = _source25.Get_().(D6_DC12).Cf23
				_ = _1900___mcc_h1
				var _1901_cf23 bool = _1900___mcc_h1
				_ = _1901_cf23
				var _1902_cf22 bool = _1899___mcc_h0
				_ = _1902_cf22
				return _pat_let_tv37
			} else if _source25.Is_DC13() {
				var _1903___mcc_h2 bool = _source25.Get_().(D6_DC13).Cf24
				_ = _1903___mcc_h2
				var _1904___mcc_h3 _dafny.Int = _source25.Get_().(D6_DC13).Cf25
				_ = _1904___mcc_h3
				var _1905_cf25 _dafny.Int = _1904___mcc_h3
				_ = _1905_cf25
				var _1906_cf24 bool = _1903___mcc_h2
				_ = _1906_cf24
				return !(_1906_cf24)
			} else {
				var _1907___mcc_h4 _dafny.Sequence = _source25.Get_().(D6_DC11).Cf21
				_ = _1907___mcc_h4
				var _1908_cf21 _dafny.Sequence = _1907___mcc_h4
				_ = _1908_cf21
				return _pat_let_tv38
			}
		}(_1898_v107))
		var _nw350 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw350
		r1 = _nw350
		r2 = Companion_Default___.SafeModuloInt(p0, (_this).Fm1(p2, globalState))
		return r0, r1, r2
	}
}
func (_this *C6) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1909_v0 bool
		_ = _1909_v0
		_1909_v0 = false
		var _1910_v1 _dafny.Array
		_ = _1910_v1
		var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw351
		_1910_v1 = _nw351
		var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
		_ = _index363
		(_1910_v1).ArraySet1(_dafny.IntOfInt64(-432), (_index363).Int())
		var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
		_ = _index364
		var _rhs321 bool = (p1).Cmp((p0).Minus(_dafny.IntOfInt64(189))) == 0
		_ = _rhs321
		var _rhs322 _dafny.Int = p0
		_ = _rhs322
		var _lhs170 _dafny.Array = _1910_v1
		_ = _lhs170
		var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
		_ = _lhs171
		_1909_v0 = _rhs321
		(_lhs170).ArraySet1(_rhs322, (_lhs171).Int())
		var _1911_i0 _dafny.Int
		_ = _1911_i0
		_1911_i0 = _dafny.Zero
		{
			for _1909_v0 {
				{
					if (_1911_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1911_i0 = (_1911_i0).Plus(_dafny.One)
					_1910_v1 = _1910_v1
					var _1912_v2 _dafny.MultiSet
					_ = _1912_v2
					_1912_v2 = _dafny.MultiSetOf(_1909_v0)
					if ((_dafny.MultiSetOf(p2)).Intersection(_1912_v2)).IsSubsetOf((_1912_v2).Union(_1912_v2)) {
						var _1913_v3 _dafny.CodePoint
						_ = _1913_v3
						_1913_v3 = _dafny.CodePoint('h')
						var _1914_v4 _dafny.MultiSet
						_ = _1914_v4
						_1914_v4 = _dafny.MultiSetOf(_1913_v3)
						var _1915_v5 _dafny.Sequence
						_ = _1915_v5
						_1915_v5 = _dafny.SeqOf((_1914_v4).Cardinality(), (_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int))
						var _1916_v6 _dafny.Sequence
						_ = _1916_v6
						_1916_v6 = _dafny.UnicodeSeqOfUtf8Bytes("dsqqc")
						var _1917_v7 _dafny.Sequence
						_ = _1917_v7
						_1917_v7 = _dafny.SeqOf(_1916_v6)
						var _1918_v8 _dafny.Sequence
						_ = _1918_v8
						_1918_v8 = _dafny.SeqOf(_1916_v6, (_1917_v7).Select((Companion_Default___.SafeIndex((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1917_v7).Cardinality()))).Uint32()).(_dafny.Sequence))
						var _1919_v9 _dafny.Map
						_ = _1919_v9
						_1919_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1917_v7).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1917_v7).Cardinality()))).Uint32()).(_dafny.Sequence))
						var _1920_v10 _dafny.Map
						_ = _1920_v10
						_1920_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1915_v5).Select((Companion_Default___.SafeIndex((_1919_v9).Cardinality(), _dafny.IntOfUint32((_1915_v5).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1916_v6).Cardinality()))
						var _1921_v11 _dafny.Set
						_ = _1921_v11
						_1921_v11 = _dafny.SetOf(p2, _1909_v0, _1909_v0)
						var _1922_v12 _dafny.Sequence
						_ = _1922_v12
						_1922_v12 = _dafny.SeqOf(_1921_v11, (_1921_v11).Difference(_1921_v11))
						var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
						_ = _index365
						var _rhs323 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), p0)
						_ = _rhs323
						var _rhs324 _dafny.Sequence = _1922_v12
						_ = _rhs324
						var _rhs325 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm47((Companion_D11_.Create_DC21_(_1913_v3, p1, p0)).Dtor_cf36(), p1, Companion_Default___.SafeDivisionInt((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), p0), globalState)).Cardinality())
						_ = _rhs325
						var _rhs326 bool = _1909_v0
						_ = _rhs326
						var _rhs327 bool = p2
						_ = _rhs327
						var _lhs172 _dafny.Array = _1910_v1
						_ = _lhs172
						var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
						_ = _lhs173
						_1920_v10 = _rhs323
						_1922_v12 = _rhs324
						(_lhs172).ArraySet1(_rhs325, (_lhs173).Int())
						_1909_v0 = _rhs326
						_1909_v0 = _rhs327
						var _1923_v13 _dafny.MultiSet
						_ = _1923_v13
						_1923_v13 = _dafny.MultiSetOf(_1910_v1)
						var _1924_v14 D5
						_ = _1924_v14
						_1924_v14 = Companion_D5_.Create_DC8_(_1923_v13)
						var _1925_v15 _dafny.Set
						_ = _1925_v15
						_1925_v15 = _dafny.SetOf(_1924_v14, _1924_v14)
						_1925_v15 = _1925_v15
						var _1926_v16 _dafny.Map
						_ = _1926_v16
						_1926_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
						_1926_v16 = (_1926_v16).Update(_1909_v0, !(_1920_v10).Contains(p0))
						r0 = (_dafny.Zero).Minus((func() _dafny.Int {
							if (_1912_v2).Contains((p2) && (_1909_v0)) {
								return (_1912_v2).Multiplicity((p2) && (_1909_v0))
							}
							return (_dafny.Zero).Minus(p1)
						})())
						var _1927_v17 _dafny.Array
						_ = _1927_v17
						var _nw352 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
						_ = _nw352
						_1927_v17 = _nw352
						var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_1927_v17), 0))
						_ = _index366
						(_1927_v17).ArraySet1(_1909_v0, (_index366).Int())
						var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_1927_v17), 0))
						_ = _index367
						(_1927_v17).ArraySet1(p2, (_index367).Int())
					} else {
						var _1928_v18 T2
						_ = _1928_v18
						var _nw353 *C4 = New_C4_()
						_ = _nw353
						_nw353.Ctor__(p1)
						_1928_v18 = _nw353
						var _1929_v19 _dafny.Map
						_ = _1929_v19
						_1929_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1928_v18, ((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus(_1928_v18.F2())))
						var _1930_v20 _dafny.Sequence
						_ = _1930_v20
						_1930_v20 = _dafny.UnicodeSeqOfUtf8Bytes("pftld")
						var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
						_ = _index368
						(_1910_v1).ArraySet1((func() _dafny.Int {
							if (_1929_v19).Contains(_1928_v18) {
								return (_1929_v19).Get(_1928_v18).(_dafny.Int)
							}
							return Companion_Default___.SafeModuloInt(_1928_v18.F2(), _dafny.IntOfUint32((_1930_v20).Cardinality()))
						})(), (_index368).Int())
						var _1931_v21 _dafny.Sequence
						_ = _1931_v21
						_1931_v21 = _dafny.SeqOf(_1912_v2, _1912_v2)
						var _1932_v22 _dafny.Map
						_ = _1932_v22
						_1932_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1909_v0, _dafny.IntOfInt64(76))
						var _1933_v23 D5
						_ = _1933_v23
						_1933_v23 = Companion_D5_.Create_DC9_(!(_1909_v0), _dafny.IntOfUint32((_1930_v20).Cardinality()), _1912_v2)
						var _1934_v24 _dafny.Set
						_ = _1934_v24
						_1934_v24 = _dafny.SetOf(_1909_v0)
						var _1935_v25 D4
						_ = _1935_v25
						_1935_v25 = Companion_D4_.Create_DC5_(_1934_v24)
						var _1936_v26 _dafny.Set
						_ = _1936_v26
						_1936_v26 = _dafny.SetOf((_dafny.Zero).Minus(((_1931_v21).Select((Companion_Default___.SafeIndex(_1928_v18.F2(), _dafny.IntOfUint32((_1931_v21).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()), p0, (_1932_v22).Cardinality(), (_1933_v23).Dtor_cf15(), ((_1935_v25).Dtor_cf8()).Cardinality())
						_1909_v0 = (_1936_v26).IsProperSubsetOf(_1936_v26)
						var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
						_ = _index369
						(_1910_v1).ArraySet1((func() _dafny.Int {
							if p2 {
								return p0
							}
							return _dafny.IntOfInt64(339)
						})(), (_index369).Int())
						var _1937_v27 _dafny.Map
						_ = _1937_v27
						_1937_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
						var _1938_v28 _dafny.Sequence
						_ = _1938_v28
						_1938_v28 = _dafny.SeqOf(!(p2), _1909_v0, _1909_v0)
						var _1939_v29 *C0
						_ = _1939_v29
						var _nw354 *C0 = New_C0_()
						_ = _nw354
						_nw354.Ctor__((func() bool {
							if p2 {
								return (func() bool {
									if (_1937_v27).Contains(_1909_v0) {
										return (_1937_v27).Get(_1909_v0).(bool)
									}
									return _1909_v0
								})()
							}
							return !(_1909_v0)
						})(), (_1938_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1938_v28).Cardinality()))).Uint32()).(bool))
						_1939_v29 = _nw354
						(_1939_v29).F4 = (Companion_Default___.SafeModuloInt((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), p0)).Cmp(p1) != 0
					}
					var _1940_v30 _dafny.Sequence
					_ = _1940_v30
					_1940_v30 = _dafny.UnicodeSeqOfUtf8Bytes("nfqqmub")
					var _1941_v31 _dafny.MultiSet
					_ = _1941_v31
					_1941_v31 = _dafny.MultiSetOf((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), p1)
					var _1942_v32 _dafny.Sequence
					_ = _1942_v32
					_1942_v32 = _dafny.SeqOf((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), p1, (_1941_v31).Cardinality())
					var _1943_v33 _dafny.Map
					_ = _1943_v33
					_1943_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1909_v0, _1942_v32)
					var _1944_v34 _dafny.Map
					_ = _1944_v34
					_1944_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1912_v2).Cardinality(), p2)
					var _1945_v35 _dafny.Map
					_ = _1945_v35
					_1945_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1940_v30, _1944_v34)
					var _1946_v36 D5
					_ = _1946_v36
					_1946_v36 = Companion_D5_.Create_DC10_(_1909_v0, p1, _1945_v35, _1909_v0)
					var _1947_v37 *C1
					_ = _1947_v37
					var _nw355 *C1 = New_C1_()
					_ = _nw355
					_nw355.Ctor__(_dafny.Companion_Sequence_.Update(_1940_v30, (Companion_Default___.SafeIndex((_1943_v33).Cardinality(), _dafny.IntOfUint32((_1940_v30).Cardinality()))).Uint32(), Companion_Default___.Fm37((_1946_v36).Dtor_cf18(), globalState)))
					_1947_v37 = _nw355
					var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
					_ = _index370
					(_1910_v1).ArraySet1((func() _dafny.Int {
						if _1909_v0 {
							return (_dafny.Zero).Minus(p1)
						}
						return p1
					})(), (_index370).Int())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1948_v38 _dafny.Array
		_ = _1948_v38
		var _len0_43 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_43
		var _nw356 _dafny.Array
		_ = _nw356
		if _len0_43.Cmp(_dafny.Zero) == 0 {
			_nw356 = _dafny.NewArray(_len0_43)
		} else {
			var _init43 func(_dafny.Int) _dafny.Sequence = func(_1949_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jcqb"), _dafny.UnicodeSeqOfUtf8Bytes("cxvrnpqcq"))
			}
			_ = _init43
			var _element0_43 = _init43(_dafny.Zero)
			_ = _element0_43
			_nw356 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
			(_nw356).ArraySet1(_element0_43, 0)
			var _nativeLen0_43 = (_len0_43).Int()
			_ = _nativeLen0_43
			for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
				(_nw356).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
			}
		}
		_1948_v38 = _nw356
		var _1950_v39 _dafny.Sequence
		_ = _1950_v39
		_1950_v39 = _dafny.UnicodeSeqOfUtf8Bytes("paxuh")
		var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1948_v38), 0))
		_ = _index371
		(_1948_v38).ArraySet1(_1950_v39, (_index371).Int())
		var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1948_v38), 0))
		_ = _index372
		(_1948_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1950_v39, _1950_v39), (_index372).Int())
		var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))
		_ = _index373
		(_1910_v1).ArraySet1(p0, (_index373).Int())
		var _1951_v40 _dafny.Map
		_ = _1951_v40
		_1951_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1950_v39).Cardinality()), (_dafny.MultiSetOf(p2, p2)).Cardinality())
		var _1952_v41 T0
		_ = _1952_v41
		var _nw357 *C4 = New_C4_()
		_ = _nw357
		_nw357.Ctor__((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int))
		_1952_v41 = _nw357
		var _1953_v42 _dafny.Sequence
		_ = _1953_v42
		_1953_v42 = _dafny.SeqOf((_1910_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1910_v1), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
			if (_1951_v40).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1952_v41, _1952_v41)).Cardinality())) {
				return (_1951_v40).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1952_v41, _1952_v41)).Cardinality())).(_dafny.Int)
			}
			return p1
		})())
		var _1954_v43 *C2
		_ = _1954_v43
		var _nw358 *C2 = New_C2_()
		_ = _nw358
		_nw358.Ctor__(p1)
		_1954_v43 = _nw358
		var _1955_v44 _dafny.Map
		_ = _1955_v44
		_1955_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1909_v0, _1954_v43)
		var _1956_v45 _dafny.Map
		_ = _1956_v45
		_1956_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1909_v0, (_1953_v42).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1955_v44).Cardinality()), _dafny.IntOfUint32((_1953_v42).Cardinality()))).Uint32()).(_dafny.Int))
		var _1957_v46 _dafny.Sequence
		_ = _1957_v46
		_1957_v46 = _dafny.SeqOf(true)
		_1956_v45 = (_1956_v45).Update((_1957_v46).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1957_v46).Cardinality()))).Uint32()).(bool), p1)
		var _1958_v47 bool
		_ = _1958_v47
		var _1959_v48 _dafny.Int
		_ = _1959_v48
		var _out37 bool
		_ = _out37
		var _out38 _dafny.Int
		_ = _out38
		_out37, _out38 = (_1954_v43).M6(_dafny.IntOfUint32(((_1948_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1948_v38), 0))).Int()).(_dafny.Sequence)).Cardinality()), globalState)
		_1958_v47 = _out37
		_1959_v48 = _out38
		r0 = p1
		return r0
	}
}
func (_this *C6) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1960_v0 *C1
		_ = _1960_v0
		var _nw359 *C1 = New_C1_()
		_ = _nw359
		_nw359.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("ui"))
		_1960_v0 = _nw359
		var _1961_v1 D8
		_ = _1961_v1
		_1961_v1 = Companion_D8_.Create_DC17_(p2, _1960_v0, (_dafny.Zero).Minus((p0).Cardinality()))
		var _source26 D8 = _1961_v1
		_ = _source26
		if _source26.Is_DC17() {
			var _1962___mcc_h0 bool = _source26.Get_().(D8_DC17).Cf30
			_ = _1962___mcc_h0
			var _1963___mcc_h1 *C1 = _source26.Get_().(D8_DC17).Cf31
			_ = _1963___mcc_h1
			var _1964___mcc_h2 _dafny.Int = _source26.Get_().(D8_DC17).Cf32
			_ = _1964___mcc_h2
			var _1965_cf32 _dafny.Int = _1964___mcc_h2
			_ = _1965_cf32
			var _1966_cf31 *C1 = _1963___mcc_h1
			_ = _1966_cf31
			var _1967_cf30 bool = _1962___mcc_h0
			_ = _1967_cf30
			var _1968_v2 D4
			_ = _1968_v2
			_1968_v2 = Companion_D4_.Create_DC6_(_dafny.IntOfInt64(862), p1, _1965_cf32)
			var _1969_v3 D4
			_ = _1969_v3
			_1969_v3 = Companion_D4_.Create_DC7_(_1968_v2)
			var _1970_v4 _dafny.Sequence
			_ = _1970_v4
			_1970_v4 = _dafny.SeqOf(_1968_v2)
			var _1971_v5 _dafny.Map
			_ = _1971_v5
			_1971_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), p3)
			var _1972_v6 D4
			_ = _1972_v6
			_1972_v6 = Companion_D4_.Create_DC7_((_1970_v4).Select((Companion_Default___.SafeIndex((_1971_v5).Cardinality(), _dafny.IntOfUint32((_1970_v4).Cardinality()))).Uint32()).(D4))
			var _1973_v7 D4
			_ = _1973_v7
			_1973_v7 = Companion_D4_.Create_DC7_((_1972_v6).Dtor_cf12())
			var _source27 D4 = _1973_v7
			_ = _source27
			if _source27.Is_DC6() {
				var _1974___mcc_h4 _dafny.Int = _source27.Get_().(D4_DC6).Cf9
				_ = _1974___mcc_h4
				var _1975___mcc_h5 _dafny.Array = _source27.Get_().(D4_DC6).Cf10
				_ = _1975___mcc_h5
				var _1976___mcc_h6 _dafny.Int = _source27.Get_().(D4_DC6).Cf11
				_ = _1976___mcc_h6
				var _1977_cf11 _dafny.Int = _1976___mcc_h6
				_ = _1977_cf11
				var _1978_cf10 _dafny.Array = _1975___mcc_h5
				_ = _1978_cf10
				var _1979_cf9 _dafny.Int = _1974___mcc_h4
				_ = _1979_cf9
				var _1980_v8 _dafny.Map
				_ = _1980_v8
				_1980_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1979_cf9).Cmp(_1977_cf11) <= 0, _dafny.CodePoint('u'))
				var _rhs328 bool = p3
				_ = _rhs328
				var _rhs329 _dafny.Map = _1980_v8
				_ = _rhs329
				_1967_cf30 = _rhs328
				_1980_v8 = _rhs329
				var _1981_v9 _dafny.CodePoint
				_ = _1981_v9
				_1981_v9 = _dafny.CodePoint('d')
				var _1982_v10 _dafny.Array
				_ = _1982_v10
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_44
				var _nw360 _dafny.Array
				_ = _nw360
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw360 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) bool = (func(_1983_p3 bool) func(_dafny.Int) bool {
						return func(_1984_i0 _dafny.Int) bool {
							return !(_1983_p3)
						}
					})(p3)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw360 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw360).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw360).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1982_v10 = _nw360
				var _1985_v11 *C5
				_ = _1985_v11
				var _nw361 *C5 = New_C5_()
				_ = _nw361
				_nw361.Ctor__((func() _dafny.Int {
					if _1967_cf30 {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wywgutya"), (Companion_Default___.SafeIndex(_1979_cf9, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wywgutya")).Cardinality()))).Uint32(), _1981_v9)).Cardinality())
					}
					return (_dafny.Zero).Minus(_1979_cf9)
				})(), _1982_v10)
				_1985_v11 = _nw361
				var _1986_v12 _dafny.Array
				_ = _1986_v12
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_45
				var _nw362 _dafny.Array
				_ = _nw362
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw362 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.Map = (func(_1987_v0 *C1) func(_dafny.Int) _dafny.Map {
						return func(_1988_i1 _dafny.Int) _dafny.Map {
							return ((Companion_D15_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1987_v0).F3(), _dafny.UnicodeSeqOfUtf8Bytes("luwy")))).Dtor_cf46()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1987_v0).F3(), _dafny.UnicodeSeqOfUtf8Bytes("avaemdf")))
						}
					})(_1960_v0)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw362 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw362).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw362).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1986_v12 = _nw362
				var _1989_v13 _dafny.Map
				_ = _1989_v13
				_1989_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1966_cf31).F3(), _dafny.UnicodeSeqOfUtf8Bytes("qgu"))
				var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_1986_v12), 0))
				_ = _index374
				(_1986_v12).ArraySet1(_1989_v13, (_index374).Int())
				var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_1986_v12), 0))
				_ = _index375
				(_1986_v12).ArraySet1(_1989_v13, (_index375).Int())
				_1967_cf30 = p3
			} else if _source27.Is_DC5() {
				var _1990___mcc_h7 _dafny.Set = _source27.Get_().(D4_DC5).Cf8
				_ = _1990___mcc_h7
				var _1991_cf8 _dafny.Set = _1990___mcc_h7
				_ = _1991_cf8
				r0 = _1965_cf32
				var _1992_v14 T2
				_ = _1992_v14
				var _nw363 *C4 = New_C4_()
				_ = _nw363
				_nw363.Ctor__(_1965_cf32)
				_1992_v14 = _nw363
				var _1993_v15 _dafny.Map
				_ = _1993_v15
				_1993_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1992_v14, _1992_v14.F2())
				var _1994_v16 _dafny.Sequence
				_ = _1994_v16
				_1994_v16 = _dafny.SeqOf((_dafny.Zero).Minus(_1965_cf32), _1965_cf32, ((_1993_v15).Update(_1992_v14, _1992_v14.F2())).Cardinality(), _dafny.IntOfInt64(232), _1965_cf32)
				_1965_cf32 = (func() _dafny.Int {
					if (_dafny.IntOfUint32((_1994_v16).Cardinality())).Cmp(_1992_v14.F2()) == 0 {
						return _1992_v14.F2()
					}
					return _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality())
				})()
				var _1995_v17 _dafny.Set
				_ = _1995_v17
				_1995_v17 = _dafny.SetOf((_1960_v0).Fm1(_1967_cf30, globalState))
				var _rhs330 bool = (_1965_cf32).Cmp(_1965_cf32) == 0
				_ = _rhs330
				var _rhs331 _dafny.Set = (_1995_v17).Intersection(_1995_v17)
				_ = _rhs331
				_1967_cf30 = _rhs330
				_1995_v17 = _rhs331
				var _1996_v18 _dafny.Sequence
				_ = _1996_v18
				_1996_v18 = _dafny.SeqOf(p3, _1967_cf30, !(p3), _1967_cf30)
				var _1997_v19 _dafny.Sequence
				_ = _1997_v19
				_1997_v19 = _dafny.SeqOf(_1996_v18, Companion_Default___.Fm16(globalState))
				var _rhs332 _dafny.Int = _1992_v14.F2()
				_ = _rhs332
				var _rhs333 _dafny.Int = (_dafny.Zero).Minus(_1965_cf32)
				_ = _rhs333
				var _rhs334 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1997_v19, _1997_v19)
				_ = _rhs334
				var _rhs335 bool = p2
				_ = _rhs335
				var _rhs336 bool = false
				_ = _rhs336
				var _lhs174 T2 = _1992_v14
				_ = _lhs174
				r0 = _rhs332
				_lhs174.F2_set_(_rhs333)
				_1997_v19 = _rhs334
				_1967_cf30 = _rhs335
				_1967_cf30 = _rhs336
			} else {
				var _1998___mcc_h8 D4 = _source27.Get_().(D4_DC7).Cf12
				_ = _1998___mcc_h8
				var _1999_cf12 D4 = _1998___mcc_h8
				_ = _1999_cf12
				var _2000_v20 _dafny.Map
				_ = _2000_v20
				_2000_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1965_cf32, (func() _dafny.Int {
					if p3 {
						return _1965_cf32
					}
					return _1965_cf32
				})())
				_2000_v20 = (_2000_v20).Update((Companion_Default___.Fm38(p3, globalState)).Cardinality(), _1965_cf32)
				var _2001_v21 _dafny.Array
				_ = _2001_v21
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_46
				var _nw364 _dafny.Array
				_ = _nw364
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw364 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) bool = (func(_2002_p2 bool) func(_dafny.Int) bool {
						return func(_2003_i2 _dafny.Int) bool {
							return !(_2002_p2)
						}
					})(p2)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw364 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw364).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw364).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_2001_v21 = _nw364
				_2001_v21 = _2001_v21
				var _2004_v22 _dafny.Sequence
				_ = _2004_v22
				_2004_v22 = _dafny.SeqOf(p3, p2)
				var _2005_v23 _dafny.Array
				_ = _2005_v23
				var _nw365 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw365
				_2005_v23 = _nw365
				var _2006_v24 bool
				_ = _2006_v24
				var _2007_v25 _dafny.Array
				_ = _2007_v25
				var _2008_v26 _dafny.Int
				_ = _2008_v26
				var _out39 bool
				_ = _out39
				var _out40 _dafny.Array
				_ = _out40
				var _out41 _dafny.Int
				_ = _out41
				_out39, _out40, _out41 = (_1960_v0).M1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2004_v22, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.IntOfUint32((_2004_v22).Cardinality()))).Uint32(), p2)).Cardinality())).Minus(_dafny.IntOfInt64(-947)), _2005_v23, p3, p3, globalState)
				_2006_v24 = _out39
				_2007_v25 = _out40
				_2008_v26 = _out41
				var _2009_v27 _dafny.Set
				_ = _2009_v27
				_2009_v27 = _dafny.SetOf(_1967_cf30)
				var _2010_v28 _dafny.Sequence
				_ = _2010_v28
				_2010_v28 = _dafny.SeqOf(_dafny.SetOf(_2006_v24), _2009_v27, _2009_v27, _2009_v27, _2009_v27)
				var _2011_v29 _dafny.Sequence
				_ = _2011_v29
				_2011_v29 = _dafny.UnicodeSeqOfUtf8Bytes("cbfxcksb")
				var _rhs337 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(p2), _2009_v27), _2010_v28), (Companion_Default___.SafeIndex(_1965_cf32, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(p2), _2009_v27), _2010_v28)).Cardinality()))).Uint32(), _2009_v27)
				_ = _rhs337
				var _rhs338 _dafny.Sequence = Companion_Default___.Fm15(globalState)
				_ = _rhs338
				var _rhs339 _dafny.Int = (_dafny.IntOfUint32(((_1960_v0).F3()).Cardinality())).Times(_1965_cf32)
				_ = _rhs339
				_2010_v28 = _rhs337
				_2011_v29 = _rhs338
				r0 = _rhs339
			}
			r0 = _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(_1967_cf30)).Cardinality()))).Cardinality())
			var _2012_v30 _dafny.Array
			_ = _2012_v30
			var _nw366 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw366
			_2012_v30 = _nw366
			var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2012_v30), 0))
			_ = _index376
			(_2012_v30).ArraySet1(p3, (_index376).Int())
			var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2012_v30), 0))
			_ = _index377
			(_2012_v30).ArraySet1(p2, (_index377).Int())
			r1 = _dafny.IntOfInt64(-385)
		} else {
			var _2013___mcc_h3 *C1 = _source26.Get_().(D8_DC16).Cf29
			_ = _2013___mcc_h3
			var _2014_cf29 *C1 = _2013___mcc_h3
			_ = _2014_cf29
			var _2015_v31 _dafny.Array
			_ = _2015_v31
			var _nw367 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw367
			_2015_v31 = _nw367
			var _2016_v32 _dafny.Int
			_ = _2016_v32
			_2016_v32 = _dafny.IntOfInt64(246)
			var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_2015_v31), 0))
			_ = _index378
			(_2015_v31).ArraySet1((_2016_v32).Cmp(_dafny.IntOfUint32(((_1960_v0).F3()).Cardinality())) > 0, (_index378).Int())
			var _2017_v33 bool
			_ = _2017_v33
			_2017_v33 = true
			var _2018_v34 _dafny.Map
			_ = _2018_v34
			_2018_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2016_v32).Cmp(_2016_v32) >= 0, !(p3))
			var _2019_v35 _dafny.CodePoint
			_ = _2019_v35
			_2019_v35 = _dafny.CodePoint('k')
			var _2020_v36 _dafny.Map
			_ = _2020_v36
			_2020_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2019_v35, _dafny.IntOfInt64(-702))
			var _2021_v37 _dafny.Map
			_ = _2021_v37
			_2021_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2020_v36).Cardinality(), p3)
			var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_2015_v31), 0))
			_ = _index379
			var _rhs340 bool = (func() bool {
				if (_2018_v34).Contains(p3) {
					return (_2018_v34).Get(p3).(bool)
				}
				return (func() bool {
					if (_2021_v37).Contains(_2016_v32) {
						return (_2021_v37).Get(_2016_v32).(bool)
					}
					return p2
				})()
			})()
			_ = _rhs340
			var _rhs341 bool = _2017_v33
			_ = _rhs341
			var _lhs175 _dafny.Array = _2015_v31
			_ = _lhs175
			var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_2015_v31), 0))
			_ = _lhs176
			(_lhs175).ArraySet1(_rhs340, (_lhs176).Int())
			_2017_v33 = _rhs341
			var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_2015_v31), 0))
			_ = _index380
			(_2015_v31).ArraySet1(!((p0).IsDisjointFrom(p0)), (_index380).Int())
			var _2022_v38 _dafny.Sequence
			_ = _2022_v38
			_2022_v38 = _dafny.SeqOf(p3, p2)
			var _2023_v39 _dafny.Array
			_ = _2023_v39
			var _out42 _dafny.Array
			_ = _out42
			_out42 = Companion_Default___.M0((_dafny.IntOfUint32((_2022_v38).Cardinality())).Cmp((_this).Fm1(Companion_Default___.Fm5(globalState), globalState)) <= 0, _dafny.Companion_Sequence_.IsPrefixOf((_2014_cf29).F3(), _dafny.UnicodeSeqOfUtf8Bytes("prkosnxd")), globalState)
			_2023_v39 = _out42
			_2017_v33 = _2017_v33
		}
		var _2024_v40 bool
		_ = _2024_v40
		_2024_v40 = true
		var _2025_v41 _dafny.Array
		_ = _2025_v41
		var _len0_47 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_47
		var _nw368 _dafny.Array
		_ = _nw368
		if _len0_47.Cmp(_dafny.Zero) == 0 {
			_nw368 = _dafny.NewArray(_len0_47)
		} else {
			var _init47 func(_dafny.Int) bool = func(_2026_i3 _dafny.Int) bool {
				return !(true)
			}
			_ = _init47
			var _element0_47 = _init47(_dafny.Zero)
			_ = _element0_47
			_nw368 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
			(_nw368).ArraySet1(_element0_47, 0)
			var _nativeLen0_47 = (_len0_47).Int()
			_ = _nativeLen0_47
			for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
				(_nw368).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
			}
		}
		_2025_v41 = _nw368
		var _2027_v42 _dafny.Sequence
		_ = _2027_v42
		_2027_v42 = _dafny.SeqOf(_2025_v41)
		var _2028_v43 D14
		_ = _2028_v43
		_2028_v43 = Companion_D14_.Create_DC27_(p3)
		var _pat_let_tv39 = _2024_v40
		_ = _pat_let_tv39
		var _pat_let_tv40 = p3
		_ = _pat_let_tv40
		var _pat_let_tv41 = _1960_v0
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1960_v0
		_ = _pat_let_tv42
		var _pat_let_tv43 = p2
		_ = _pat_let_tv43
		var _pat_let_tv44 = p3
		_ = _pat_let_tv44
		var _pat_let_tv45 = p2
		_ = _pat_let_tv45
		var _rhs342 bool = func(_source28 D14) bool {
			if _source28.Is_DC26() {
				var _2029___mcc_h9 D0 = _source28.Get_().(D14_DC26).Cf43
				_ = _2029___mcc_h9
				var _2030___mcc_h10 _dafny.Set = _source28.Get_().(D14_DC26).Cf44
				_ = _2030___mcc_h10
				var _2031_cf44 _dafny.Set = _2030___mcc_h10
				_ = _2031_cf44
				var _2032_cf43 D0 = _2029___mcc_h9
				_ = _2032_cf43
				return ((Companion_D0_.Create_DC1_(_pat_let_tv39, _pat_let_tv40, _dafny.CodePoint('s'), _dafny.IntOfUint32(((_pat_let_tv41).F3()).Cardinality()))).Dtor_cf4()).Cmp((_dafny.SetOf(_dafny.IntOfUint32(((_pat_let_tv42).F3()).Cardinality()))).Cardinality()) >= 0
			} else if _source28.Is_DC27() {
				var _2033___mcc_h11 bool = _source28.Get_().(D14_DC27).Cf45
				_ = _2033___mcc_h11
				var _2034_cf45 bool = _2033___mcc_h11
				_ = _2034_cf45
				return ((_dafny.SetOf(_2034_cf45)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_2034_cf45, _pat_let_tv43)).Cardinality())) >= 0
			} else if _source28.Is_DC28() {
				return _pat_let_tv44
			} else {
				var _2035___mcc_h12 _dafny.Array = _source28.Get_().(D14_DC25).Cf42
				_ = _2035___mcc_h12
				var _2036_cf42 _dafny.Array = _2035___mcc_h12
				_ = _2036_cf42
				return _pat_let_tv45
			}
		}(_2028_v43)
		_ = _rhs342
		var _rhs343 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2025_v41, _2025_v41, _2025_v41, _2025_v41), _2027_v42)
		_ = _rhs343
		_2024_v40 = _rhs342
		_2027_v42 = _rhs343
		var _2037_v44 _dafny.Set
		_ = _2037_v44
		_2037_v44 = _dafny.SetOf(p1, p1, p1)
		var _2038_v45 _dafny.Sequence
		_ = _2038_v45
		_2038_v45 = _dafny.SeqOf(p3)
		var _2039_v46 _dafny.Map
		_ = _2039_v46
		_2039_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v45, _2024_v40)
		var _2040_v47 _dafny.Int
		_ = _2040_v47
		_2040_v47 = _dafny.IntOfInt64(420)
		var _2041_v48 _dafny.Set
		_ = _2041_v48
		_2041_v48 = _dafny.SetOf(((_2039_v46).Merge(_2039_v46)).Cardinality(), _2040_v47)
		var _2042_v49 _dafny.Map
		_ = _2042_v49
		_2042_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p2, !(true)), _2037_v44)
		var _2043_v50 _dafny.Map
		_ = _2043_v50
		_2043_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()), _2040_v47)
		var _2044_v51 _dafny.Map
		_ = _2044_v51
		_2044_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2025_v41, (func() _dafny.Int {
			if (_2043_v50).Contains(_2040_v47) {
				return (_2043_v50).Get(_2040_v47).(_dafny.Int)
			}
			return _2040_v47
		})())
		var _rhs344 _dafny.Int = Companion_Default___.SafeModuloInt(_2040_v47, _2040_v47)
		_ = _rhs344
		var _rhs345 _dafny.Set = ((func() _dafny.Set {
			if (_2042_v49).Contains((_dafny.MultiSetFromSeq(_2038_v45)).Update(p3, Companion_Default___.Abs(_2040_v47))) {
				return (_2042_v49).Get((_dafny.MultiSetFromSeq(_2038_v45)).Update(p3, Companion_Default___.Abs(_2040_v47))).(_dafny.Set)
			}
			return _2037_v44
		})()).Union(_2037_v44)
		_ = _rhs345
		var _rhs346 _dafny.Set = (Companion_Default___.Fm48(_2024_v40, globalState))
		_ = _rhs346
		var _rhs347 _dafny.Int = (_2044_v51).Cardinality()
		_ = _rhs347
		r0 = _rhs344
		_2037_v44 = _rhs345
		_2041_v48 = _rhs346
		r0 = _rhs347
		var _2045_v52 _dafny.Sequence
		_ = _2045_v52
		_2045_v52 = _dafny.SeqOf(_2040_v47)
		var _2046_v53 D7
		_ = _2046_v53
		_2046_v53 = Companion_D7_.Create_DC14_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2040_v47), _2045_v52))
		var _source29 D7 = _2046_v53
		_ = _source29
		if _source29.Is_DC15() {
			var _2047___mcc_h13 _dafny.Int = _source29.Get_().(D7_DC15).Cf27
			_ = _2047___mcc_h13
			var _2048___mcc_h14 _dafny.CodePoint = _source29.Get_().(D7_DC15).Cf28
			_ = _2048___mcc_h14
			var _2049_cf28 _dafny.CodePoint = _2048___mcc_h14
			_ = _2049_cf28
			var _2050_cf27 _dafny.Int = _2047___mcc_h13
			_ = _2050_cf27
			var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))
			_ = _index381
			(_2025_v41).ArraySet1(_2024_v40, (_index381).Int())
			var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))
			_ = _index382
			(_2025_v41).ArraySet1(_2024_v40, (_index382).Int())
			var _2051_v54 *C4
			_ = _2051_v54
			var _nw369 *C4 = New_C4_()
			_ = _nw369
			_nw369.Ctor__(_2050_cf27)
			_2051_v54 = _nw369
			var _2052_v55 _dafny.Array
			_ = _2052_v55
			var _len0_48 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_48
			var _nw370 _dafny.Array
			_ = _nw370
			if _len0_48.Cmp(_dafny.Zero) == 0 {
				_nw370 = _dafny.NewArray(_len0_48)
			} else {
				var _init48 func(_dafny.Int) bool = (func(_2053_p2 bool) func(_dafny.Int) bool {
					return func(_2054_i4 _dafny.Int) bool {
						return _2053_p2
					}
				})(p2)
				_ = _init48
				var _element0_48 = _init48(_dafny.Zero)
				_ = _element0_48
				_nw370 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
				(_nw370).ArraySet1(_element0_48, 0)
				var _nativeLen0_48 = (_len0_48).Int()
				_ = _nativeLen0_48
				for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
					(_nw370).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
				}
			}
			_2052_v55 = _nw370
			var _2055_v56 _dafny.Array
			_ = _2055_v56
			var _nwElement0_83 _dafny.Array = _2052_v55
			_ = _nwElement0_83
			var _nw371 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(4))
			_ = _nw371
			(_nw371).ArraySet1(_nwElement0_83, 0)
			(_nw371).ArraySet1(_2052_v55, 1)
			(_nw371).ArraySet1(_2052_v55, 2)
			(_nw371).ArraySet1(_2052_v55, 3)
			_2055_v56 = _nw371
			var _2056_v57 D14
			_ = _2056_v57
			_2056_v57 = Companion_D14_.Create_DC25_(_2055_v56)
			var _pat_let_tv46 = _2055_v56
			_ = _pat_let_tv46
			var _source30 D14 = func(_pat_let41_0 D14) D14 {
				return func(_2057_dt__update__tmp_h0 D14) D14 {
					return func(_pat_let42_0 _dafny.Array) D14 {
						return func(_2058_dt__update_hcf42_h0 _dafny.Array) D14 {
							return Companion_D14_.Create_DC25_(_2058_dt__update_hcf42_h0)
						}(_pat_let42_0)
					}(_pat_let_tv46)
				}(_pat_let41_0)
			}(_2056_v57)
			_ = _source30
			if _source30.Is_DC26() {
				var _2059___mcc_h16 D0 = _source30.Get_().(D14_DC26).Cf43
				_ = _2059___mcc_h16
				var _2060___mcc_h17 _dafny.Set = _source30.Get_().(D14_DC26).Cf44
				_ = _2060___mcc_h17
				var _2061_cf44 _dafny.Set = _2060___mcc_h17
				_ = _2061_cf44
				var _2062_cf43 D0 = _2059___mcc_h16
				_ = _2062_cf43
				_2040_v47 = (Companion_Default___.SafeModuloInt(_2050_cf27, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2040_v47, p3)).Cardinality())).Minus(_2040_v47)
				r1 = (_2040_v47).Minus((_this).Fm1((_2025_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))).Int()).(bool), globalState))
				var _2063_v58 _dafny.Map
				_ = _2063_v58
				_2063_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				_2063_v58 = (_2063_v58).Update(p1, (_2025_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))).Int()).(bool))
				var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))
				_ = _index383
				(_2025_v41).ArraySet1(p3, (_index383).Int())
			} else if _source30.Is_DC27() {
				var _2064___mcc_h18 bool = _source30.Get_().(D14_DC27).Cf45
				_ = _2064___mcc_h18
				var _2065_cf45 bool = _2064___mcc_h18
				_ = _2065_cf45
				_2040_v47 = Companion_Default___.SafeDivisionInt(_2050_cf27, _2050_cf27)
				var _2066_v59 *C1
				_ = _2066_v59
				var _nw372 *C1 = New_C1_()
				_ = _nw372
				_nw372.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-916))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg104 _dafny.Int) interface{} {
						return coer104(arg104)
					}
				}((func(_2067_cf28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2068_i5 _dafny.Int) _dafny.CodePoint {
						return _2067_cf28
					}
				})(_2049_cf28))), _dafny.UnicodeSeqOfUtf8Bytes("du")))
				_2066_v59 = _nw372
				r1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_2040_v47).Minus(_2050_cf27)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_2027_v42).Cardinality())))
				var _2069_v60 _dafny.Array
				_ = _2069_v60
				var _nwElement0_84 _dafny.Int = _2040_v47
				_ = _nwElement0_84
				var _nw373 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(6))
				_ = _nw373
				(_nw373).ArraySet1(_nwElement0_84, 0)
				(_nw373).ArraySet1(_2040_v47, 1)
				(_nw373).ArraySet1(_dafny.IntOfInt64(-630), 2)
				(_nw373).ArraySet1(_2050_cf27, 3)
				(_nw373).ArraySet1((_dafny.SetOf(_2049_cf28, _2049_cf28, _2049_cf28)).Cardinality(), 4)
				(_nw373).ArraySet1(_dafny.IntOfInt64(358), 5)
				_2069_v60 = _nw373
				_2069_v60 = p1
			} else if _source30.Is_DC28() {
				_2024_v40 = (_2038_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()), _dafny.IntOfUint32((_2038_v45).Cardinality()))).Uint32()).(bool)
				var _2070_v61 _dafny.MultiSet
				_ = _2070_v61
				_2070_v61 = p0
				var _2071_v62 _dafny.Map
				_ = _2071_v62
				_2071_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2070_v61, _2049_cf28)
				var _2072_v63 _dafny.Set
				_ = _2072_v63
				_2072_v63 = _dafny.SetOf(p3, true, Companion_Default___.Fm5(globalState), p3)
				var _2073_v64 D4
				_ = _2073_v64
				_2073_v64 = Companion_D4_.Create_DC5_(_2072_v63)
				var _2074_v65 _dafny.Map
				_ = _2074_v65
				_2074_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2071_v62, _2073_v64)
				_2074_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2071_v62, _2073_v64)
				var _2075_v66 *C1
				_ = _2075_v66
				var _nw374 *C1 = New_C1_()
				_ = _nw374
				_nw374.Ctor__((_1960_v0).F3())
				_2075_v66 = _nw374
				var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))
				_ = _index384
				(_2025_v41).ArraySet1((_2025_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))).Int()).(bool), (_index384).Int())
			} else {
				var _2076___mcc_h19 _dafny.Array = _source30.Get_().(D14_DC25).Cf42
				_ = _2076___mcc_h19
				var _2077_cf42 _dafny.Array = _2076___mcc_h19
				_ = _2077_cf42
				var _2078_v67 _dafny.Array
				_ = _2078_v67
				var _nw375 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw375
				_2078_v67 = _nw375
				var _2079_v68 bool
				_ = _2079_v68
				var _2080_v69 _dafny.Array
				_ = _2080_v69
				var _2081_v70 _dafny.Int
				_ = _2081_v70
				var _out43 bool
				_ = _out43
				var _out44 _dafny.Array
				_ = _out44
				var _out45 _dafny.Int
				_ = _out45
				_out43, _out44, _out45 = (_2051_v54).M1(_2050_cf27, _2078_v67, !(_dafny.Companion_Sequence_.IsPrefixOf(_2038_v45, _2038_v45)), p2, globalState)
				_2079_v68 = _out43
				_2080_v69 = _out44
				_2081_v70 = _out45
				_2038_v45 = _2038_v45
				var _2082_v71 _dafny.Sequence
				_ = _2082_v71
				_2082_v71 = _dafny.UnicodeSeqOfUtf8Bytes("vdfxyi")
				_2082_v71 = _dafny.Companion_Sequence_.Concatenate((_1960_v0).F3(), _dafny.Companion_Sequence_.Update((_1960_v0).F3(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-474), _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32(), _2049_cf28))
				var _2083_v72 *C3
				_ = _2083_v72
				var _nw376 *C3 = New_C3_()
				_ = _nw376
				_nw376.Ctor__(((_this).Fm1(p2, globalState)).Plus(_2050_cf27))
				_2083_v72 = _nw376
			}
			_2024_v40 = (func() bool {
				if (_2025_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_2025_v41), 0))).Int()).(bool) {
					return p2
				}
				return p3
			})()
		} else {
			var _2084___mcc_h15 _dafny.Sequence = _source29.Get_().(D7_DC14).Cf26
			_ = _2084___mcc_h15
			var _2085_cf26 _dafny.Sequence = _2084___mcc_h15
			_ = _2085_cf26
			var _2086_v73 _dafny.Map
			_ = _2086_v73
			_2086_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1960_v0).F3(), p2)
			var _2087_v74 _dafny.Map
			_ = _2087_v74
			_2087_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2040_v47, _2024_v40)
			_2086_v73 = (_2086_v73).Update((_1960_v0).F3(), (_2024_v40) && ((func() bool {
				if (_2087_v74).Contains(_2040_v47) {
					return (_2087_v74).Get(_2040_v47).(bool)
				}
				return p3
			})()))
			var _2088_v75 *C2
			_ = _2088_v75
			var _nw377 *C2 = New_C2_()
			_ = _nw377
			_nw377.Ctor__((_dafny.Zero).Minus(_2040_v47))
			_2088_v75 = _nw377
			var _2089_v76 _dafny.Sequence
			_ = _2089_v76
			_2089_v76 = _dafny.UnicodeSeqOfUtf8Bytes("fuumxjor")
			var _2090_v77 _dafny.CodePoint
			_ = _2090_v77
			_2090_v77 = _dafny.CodePoint('x')
			_2089_v76 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}(func(_2091_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2089_v76, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2085_cf26).Cardinality()), _dafny.IntOfUint32((_2089_v76).Cardinality()))).Uint32(), ((_1960_v0).F3()).Select((Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32()).(_dafny.CodePoint)), (Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2089_v76, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2085_cf26).Cardinality()), _dafny.IntOfUint32((_2089_v76).Cardinality()))).Uint32(), ((_1960_v0).F3()).Select((Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()))).Uint32(), _2090_v77))
			_2040_v47 = _2040_v47
		}
		var _2092_v78 _dafny.CodePoint
		_ = _2092_v78
		_2092_v78 = _dafny.CodePoint('n')
		var _2093_v79 _dafny.Array
		_ = _2093_v79
		var _nwElement0_85 _dafny.Sequence = (_1960_v0).F3()
		_ = _nwElement0_85
		var _nw378 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(22))
		_ = _nw378
		(_nw378).ArraySet1(_nwElement0_85, 0)
		(_nw378).ArraySet1(_dafny.Companion_Sequence_.Update((_1960_v0).F3(), (Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32(), _2092_v78), 1)
		(_nw378).ArraySet1((_1960_v0).F3(), 2)
		(_nw378).ArraySet1((_1960_v0).F3(), 3)
		(_nw378).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bxjtpgk"), 4)
		(_nw378).ArraySet1((_1960_v0).F3(), 5)
		(_nw378).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bcgsqra"), 6)
		(_nw378).ArraySet1((_1960_v0).F3(), 7)
		(_nw378).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-609))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}(func(_2094_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), 8)
		(_nw378).ArraySet1((_1960_v0).F3(), 9)
		(_nw378).ArraySet1((_1960_v0).F3(), 10)
		(_nw378).ArraySet1((_1960_v0).F3(), 11)
		(_nw378).ArraySet1((_1960_v0).F3(), 12)
		(_nw378).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(365))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg107 _dafny.Int) interface{} {
				return coer107(arg107)
			}
		}((func(_2095_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2096_i8 _dafny.Int) _dafny.CodePoint {
				return _2095_v78
			}
		})(_2092_v78))), 13)
		(_nw378).ArraySet1((_1960_v0).F3(), 14)
		(_nw378).ArraySet1((_1960_v0).F3(), 15)
		(_nw378).ArraySet1((_1960_v0).F3(), 16)
		(_nw378).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("eoy"), (Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eoy")).Cardinality()))).Uint32(), _dafny.CodePoint('l')), 17)
		(_nw378).ArraySet1((_1960_v0).F3(), 18)
		(_nw378).ArraySet1((_1960_v0).F3(), 19)
		(_nw378).ArraySet1((_1960_v0).F3(), 20)
		(_nw378).ArraySet1((_1960_v0).F3(), 21)
		_2093_v79 = _nw378
		var _2097_v80 _dafny.Map
		_ = _2097_v80
		_2097_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2040_v47, _2045_v52)
		var _2098_v81 _dafny.Array
		_ = _2098_v81
		var _nwElement0_86 _dafny.Sequence = _2045_v52
		_ = _nwElement0_86
		var _nw379 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(29))
		_ = _nw379
		(_nw379).ArraySet1(_nwElement0_86, 0)
		(_nw379).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-225))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg108 _dafny.Int) interface{} {
				return coer108(arg108)
			}
		}(func(_2099_i9 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(61)
		})), 1)
		(_nw379).ArraySet1(_2045_v52, 2)
		(_nw379).ArraySet1(_2045_v52, 3)
		(_nw379).ArraySet1(_2045_v52, 4)
		(_nw379).ArraySet1(Companion_Default___.Fm36(p2, _2092_v78, (_2041_v48).Cardinality(), globalState), 5)
		(_nw379).ArraySet1(_2045_v52, 6)
		(_nw379).ArraySet1(_2045_v52, 7)
		(_nw379).ArraySet1(Companion_Default___.Fm36(p3, _dafny.CodePoint('b'), _2040_v47, globalState), 8)
		(_nw379).ArraySet1(_2045_v52, 9)
		(_nw379).ArraySet1(_2045_v52, 10)
		(_nw379).ArraySet1(_2045_v52, 11)
		(_nw379).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32(((_1960_v0).F3()).Cardinality())), 12)
		(_nw379).ArraySet1(_2045_v52, 13)
		(_nw379).ArraySet1(_dafny.SeqOf(_2040_v47), 14)
		(_nw379).ArraySet1(_2045_v52, 15)
		(_nw379).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-172))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg109 _dafny.Int) interface{} {
				return coer109(arg109)
			}
		}((func(_2100_v47 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_2101_i10 _dafny.Int) _dafny.Int {
				return _2100_v47
			}
		})(_2040_v47))), 16)
		(_nw379).ArraySet1(_2045_v52, 17)
		(_nw379).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_2045_v52).Cardinality())), 18)
		(_nw379).ArraySet1(_dafny.Companion_Sequence_.Update(_2045_v52, (Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32((_2045_v52).Cardinality()))).Uint32(), _dafny.IntOfInt64(5)), 19)
		(_nw379).ArraySet1(_2045_v52, 20)
		(_nw379).ArraySet1(_2045_v52, 21)
		(_nw379).ArraySet1((func() _dafny.Sequence {
			if (_2097_v80).Contains(_2040_v47) {
				return (_2097_v80).Get(_2040_v47).(_dafny.Sequence)
			}
			return _2045_v52
		})(), 22)
		(_nw379).ArraySet1(_dafny.SeqOf(_2040_v47, _2040_v47), 23)
		(_nw379).ArraySet1(_2045_v52, 24)
		(_nw379).ArraySet1(_2045_v52, 25)
		(_nw379).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg110 _dafny.Int) interface{} {
				return coer110(arg110)
			}
		}((func(_2102_v50 _dafny.Map) func(_dafny.Int) _dafny.Int {
			return func(_2103_i11 _dafny.Int) _dafny.Int {
				return (_2102_v50).Cardinality()
			}
		})(_2043_v50))), 26)
		(_nw379).ArraySet1(_2045_v52, 27)
		(_nw379).ArraySet1(_dafny.SeqOf((_1960_v0).Fm1(true, globalState)), 28)
		_2098_v81 = _nw379
		var _2104_v82 _dafny.Map
		_ = _2104_v82
		_2104_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2093_v79, _2098_v81)
		var _2105_v83 bool
		_ = _2105_v83
		var _2106_v84 _dafny.Array
		_ = _2106_v84
		var _2107_v85 _dafny.Int
		_ = _2107_v85
		var _out46 bool
		_ = _out46
		var _out47 _dafny.Array
		_ = _out47
		var _out48 _dafny.Int
		_ = _out48
		_out46, _out47, _out48 = (_1960_v0).M1(_2040_v47, (func() _dafny.Array {
			if (_2104_v82).Contains(_2093_v79) {
				return (_2104_v82).Get(_2093_v79).(_dafny.Array)
			}
			return _2098_v81
		})(), p2, p2, globalState)
		_2105_v83 = _out46
		_2106_v84 = _out47
		_2107_v85 = _out48
		var _2108_v86 _dafny.Map
		_ = _2108_v86
		_2108_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2025_v41, ((_1960_v0).F3()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-578), _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _2109_v87 _dafny.Map
		_ = _2109_v87
		_2109_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2108_v86, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v45, _2040_v47))
		var _2110_v88 _dafny.Map
		_ = _2110_v88
		_2110_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2038_v45, _2040_v47)
		_2109_v87 = (_2109_v87).Update((_2108_v86).Merge(_2108_v86), _2110_v88)
		var _2111_v89 _dafny.Set
		_ = _2111_v89
		_2111_v89 = _dafny.SetOf(_2046_v53)
		r0 = (((_2111_v89).Union(_dafny.SetOf(_2046_v53))).Union(_2111_v89)).Cardinality()
		var _2112_v90 _dafny.Sequence
		_ = _2112_v90
		_2112_v90 = _dafny.SeqOf(_2041_v48, _2041_v48)
		r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1960_v0).F3(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2112_v90).Cardinality()), _dafny.IntOfUint32(((_1960_v0).F3()).Cardinality()))).Uint32(), _2092_v78)).Cardinality())).Times((func() _dafny.Int {
			if (_2038_v45).Select((Companion_Default___.SafeIndex(_2040_v47, _dafny.IntOfUint32((_2038_v45).Cardinality()))).Uint32()).(bool) {
				return (_dafny.Zero).Minus(_2107_v85)
			}
			return (_this).Fm1(p2, globalState)
		})())
		return r0, r1
	}
}
func (_this *C6) M4(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		if p3 {
			var _2113_v0 _dafny.Array
			_ = _2113_v0
			var _nw380 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw380
			_2113_v0 = _nw380
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_2113_v0), 0))
			_ = _index385
			(_2113_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()), (_index385).Int())
			var _2114_v1 _dafny.Sequence
			_ = _2114_v1
			_2114_v1 = _dafny.SeqOf(p3, !(p3), !(p3), p0)
			var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_2113_v0), 0))
			_ = _index386
			(_2113_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2114_v1, _dafny.SeqOf(p3))).Cardinality()), (_index386).Int())
			r3 = p3
			var _2115_v2 _dafny.Sequence
			_ = _2115_v2
			_2115_v2 = _dafny.UnicodeSeqOfUtf8Bytes("iawbfjbim")
			var _2116_v3 _dafny.CodePoint
			_ = _2116_v3
			_2116_v3 = _dafny.CodePoint('h')
			var _2117_v4 _dafny.Map
			_ = _2117_v4
			_2117_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2116_v3, p0)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg111 _dafny.Int) interface{} {
					return coer111(arg111)
				}
			}((func(_2118_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2119_i0 _dafny.Int) _dafny.CodePoint {
					return _2118_v3
				}
			})(_2116_v3))))
			_2115_v2 = _dafny.Companion_Sequence_.Concatenate(_2115_v2, (func() _dafny.Sequence {
				if (_2117_v4).Contains((_2113_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_2113_v0), 0))).Int()).(_dafny.Int)) {
					return (_2117_v4).Get((_2113_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_2113_v0), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _2115_v2
			})())
			var _2120_v5 _dafny.Array
			_ = _2120_v5
			var _nw381 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw381
			_2120_v5 = _nw381
			var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_2120_v5), 0))
			_ = _index387
			(_2120_v5).ArraySet1(!(!(true)), (_index387).Int())
			var _2121_v6 _dafny.MultiSet
			_ = _2121_v6
			_2121_v6 = _dafny.MultiSetOf(_dafny.IntOfInt64(8))
			var _2122_v7 _dafny.MultiSet
			_ = _2122_v7
			_2122_v7 = _dafny.MultiSetOf(_2121_v6)
			var _2123_v8 _dafny.Sequence
			_ = _2123_v8
			_2123_v8 = _dafny.SeqOf(p2)
			var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_2120_v5), 0))
			_ = _index388
			(_2120_v5).ArraySet1((func() bool {
				if p0 {
					return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ffeyfgi"), p1)
				}
				return (_2122_v7).IsSubsetOf(_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_2123_v8)))
			})(), (_index388).Int())
			var _2124_v9 _dafny.Array
			_ = _2124_v9
			var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw382
			_2124_v9 = _nw382
			var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_2124_v9), 0))
			_ = _index389
			(_2124_v9).ArraySet1(_2113_v0, (_index389).Int())
			var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_2124_v9), 0))
			_ = _index390
			(_2124_v9).ArraySet1(_2113_v0, (_index390).Int())
		} else {
			var _2125_v10 _dafny.Set
			_ = _2125_v10
			_2125_v10 = _dafny.SetOf(p2)
			var _2126_v11 _dafny.MultiSet
			_ = _2126_v11
			_2126_v11 = _dafny.MultiSetOf(p0)
			var _2127_v12 D5
			_ = _2127_v12
			_2127_v12 = Companion_D5_.Create_DC9_(p0, p2, _2126_v11)
			var _2128_v13 _dafny.Array
			_ = _2128_v13
			var _nwElement0_87 bool = p0
			_ = _nwElement0_87
			var _nw383 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(22))
			_ = _nw383
			(_nw383).ArraySet1(_nwElement0_87, 0)
			(_nw383).ArraySet1(p3, 1)
			(_nw383).ArraySet1(p3, 2)
			(_nw383).ArraySet1(p3, 3)
			(_nw383).ArraySet1(p0, 4)
			(_nw383).ArraySet1(p3, 5)
			(_nw383).ArraySet1(p3, 6)
			(_nw383).ArraySet1(Companion_Default___.Fm5(globalState), 7)
			(_nw383).ArraySet1(p3, 8)
			(_nw383).ArraySet1(p0, 9)
			(_nw383).ArraySet1(p3, 10)
			(_nw383).ArraySet1(p0, 11)
			(_nw383).ArraySet1(p3, 12)
			(_nw383).ArraySet1(p3, 13)
			(_nw383).ArraySet1(p0, 14)
			(_nw383).ArraySet1(p3, 15)
			(_nw383).ArraySet1(p0, 16)
			(_nw383).ArraySet1(p0, 17)
			(_nw383).ArraySet1(!((_2127_v12).Dtor_cf14()), 18)
			(_nw383).ArraySet1(p3, 19)
			(_nw383).ArraySet1(p3, 20)
			(_nw383).ArraySet1(Companion_Default___.Fm5(globalState), 21)
			_2128_v13 = _nw383
			var _2129_v14 *C5
			_ = _2129_v14
			var _nw384 *C5 = New_C5_()
			_ = _nw384
			_nw384.Ctor__((_2125_v10).Cardinality(), _2128_v13)
			_2129_v14 = _nw384
			var _2130_v15 _dafny.Array
			_ = _2130_v15
			var _nw385 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw385
			_2130_v15 = _nw385
			var _2131_v16 _dafny.Sequence
			_ = _2131_v16
			_2131_v16 = _dafny.UnicodeSeqOfUtf8Bytes("yghybw")
			var _rhs348 _dafny.Array = (func() _dafny.Array {
				if p0 {
					return _2130_v15
				}
				return _2130_v15
			})()
			_ = _rhs348
			var _rhs349 _dafny.Int = p2
			_ = _rhs349
			var _rhs350 _dafny.Sequence = _2131_v16
			_ = _rhs350
			_2130_v15 = _rhs348
			r2 = _rhs349
			_2131_v16 = _rhs350
			var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))
			_ = _index391
			(_2130_v15).ArraySet1(p2, (_index391).Int())
			var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))
			_ = _index392
			var _rhs351 _dafny.Array = _2129_v14.F1
			_ = _rhs351
			var _rhs352 bool = p3
			_ = _rhs352
			var _rhs353 _dafny.Int = p2
			_ = _rhs353
			var _lhs177 *C5 = _2129_v14
			_ = _lhs177
			var _lhs178 _dafny.Array = _2130_v15
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))
			_ = _lhs179
			_lhs177.F1 = _rhs351
			r3 = _rhs352
			(_lhs178).ArraySet1(_rhs353, (_lhs179).Int())
			if !(true) || (!(p0) || (p3)) {
				var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))
				_ = _index393
				(_2130_v15).ArraySet1((_2129_v14).F0(), (_index393).Int())
				var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))
				_ = _index394
				(_2130_v15).ArraySet1((_2129_v14).F0(), (_index394).Int())
				var _2132_v17 *C2
				_ = _2132_v17
				var _nw386 *C2 = New_C2_()
				_ = _nw386
				_nw386.Ctor__(p2)
				_2132_v17 = _nw386
				var _2133_v18 _dafny.Map
				_ = _2133_v18
				_2133_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2131_v16)
				r2 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm15(globalState), (func() _dafny.Sequence {
					if (_2133_v18).Contains(!(!(p0))) {
						return (_2133_v18).Get(!(!(p0))).(_dafny.Sequence)
					}
					return p1
				})())).Cardinality())).Times(p2)
				r3 = p0
			} else {
				r3 = p3
				var _2134_v19 *C3
				_ = _2134_v19
				var _nw387 *C3 = New_C3_()
				_ = _nw387
				_nw387.Ctor__((_2130_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_2130_v15), 0))).Int()).(_dafny.Int))
				_2134_v19 = _nw387
				var _2135_v20 _dafny.CodePoint
				_ = _2135_v20
				_2135_v20 = _dafny.CodePoint('s')
				_2135_v20 = _2135_v20
				var _2136_v21 _dafny.Sequence
				_ = _2136_v21
				_2136_v21 = _dafny.SeqOf(p0)
				var _2137_v22 _dafny.Map
				_ = _2137_v22
				_2137_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
				r3 = (_2136_v21).Select((Companion_Default___.SafeIndex(((func() _dafny.Map {
					if p0 {
						return _2137_v22
					}
					return _2137_v22
				})()).Cardinality(), _dafny.IntOfUint32((_2136_v21).Cardinality()))).Uint32()).(bool)
				_2130_v15 = _2130_v15
			}
			r3 = (func() bool {
				if p3 {
					return p0
				}
				return p3
			})()
		}
		if p0 {
			var _2138_v23 _dafny.Array
			_ = _2138_v23
			var _nw388 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw388
			_2138_v23 = _nw388
			var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
			_ = _index395
			(_2138_v23).ArraySet1((_dafny.IntOfInt64(890)).Plus(p2), (_index395).Int())
			var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
			_ = _index396
			(_2138_v23).ArraySet1(p2, (_index396).Int())
			var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
			_ = _index397
			(_2138_v23).ArraySet1(p2, (_index397).Int())
			var _2139_v24 _dafny.Map
			_ = _2139_v24
			_2139_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
			var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
			_ = _index398
			(_2138_v23).ArraySet1(((_2139_v24).Merge(_2139_v24)).Cardinality(), (_index398).Int())
			var _2140_v25 *C2
			_ = _2140_v25
			var _nw389 *C2 = New_C2_()
			_ = _nw389
			_nw389.Ctor__((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int))
			_2140_v25 = _nw389
			var _2141_v26 D6
			_ = _2141_v26
			_2141_v26 = Companion_D6_.Create_DC11_(p1)
			var _source31 D6 = _2141_v26
			_ = _source31
			if _source31.Is_DC12() {
				var _2142___mcc_h0 bool = _source31.Get_().(D6_DC12).Cf22
				_ = _2142___mcc_h0
				var _2143___mcc_h1 bool = _source31.Get_().(D6_DC12).Cf23
				_ = _2143___mcc_h1
				var _2144_cf23 bool = _2143___mcc_h1
				_ = _2144_cf23
				var _2145_cf22 bool = _2142___mcc_h0
				_ = _2145_cf22
				var _2146_v27 _dafny.Sequence
				_ = _2146_v27
				_2146_v27 = _dafny.SeqOf(_2145_cf22, _2145_cf22)
				var _2147_v28 _dafny.Array
				_ = _2147_v28
				var _nwElement0_88 bool = Companion_Default___.Fm5(globalState)
				_ = _nwElement0_88
				var _nw390 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(6))
				_ = _nw390
				(_nw390).ArraySet1(_nwElement0_88, 0)
				(_nw390).ArraySet1(_2145_cf22, 1)
				(_nw390).ArraySet1(_2144_cf23, 2)
				(_nw390).ArraySet1(p0, 3)
				(_nw390).ArraySet1((_2146_v27).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2144_cf23)).Cardinality(), _dafny.IntOfUint32((_2146_v27).Cardinality()))).Uint32()).(bool), 4)
				(_nw390).ArraySet1(_2144_cf23, 5)
				_2147_v28 = _nw390
				var _2148_v29 *C5
				_ = _2148_v29
				var _nw391 *C5 = New_C5_()
				_ = _nw391
				_nw391.Ctor__(p2, (func() _dafny.Array {
					if p3 {
						return _2147_v28
					}
					return _2147_v28
				})())
				_2148_v29 = _nw391
				var _arr15 _dafny.Array = _2148_v29.F1
				_ = _arr15
				var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_2148_v29.F1), 0))
				_ = _index399
				(_arr15).ArraySet1(_2144_cf23, (_index399).Int())
				var _2149_v30 _dafny.Map
				_ = _2149_v30
				_2149_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(151)).Minus(_dafny.IntOfInt64(32)), p0)
				var _arr16 _dafny.Array = _2148_v29.F1
				_ = _arr16
				var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_2148_v29.F1), 0))
				_ = _index400
				(_arr16).ArraySet1((func() bool {
					if (_2149_v30).Contains((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)) {
						return (_2149_v30).Get((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)).(bool)
					}
					return _2144_cf23
				})(), (_index400).Int())
				_2144_cf23 = (_2148_v29.F1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_2148_v29.F1), 0))).Int()).(bool)
				var _2150_v31 _dafny.CodePoint
				_ = _2150_v31
				_2150_v31 = _dafny.CodePoint('l')
				_2150_v31 = _2150_v31
			} else if _source31.Is_DC13() {
				var _2151___mcc_h2 bool = _source31.Get_().(D6_DC13).Cf24
				_ = _2151___mcc_h2
				var _2152___mcc_h3 _dafny.Int = _source31.Get_().(D6_DC13).Cf25
				_ = _2152___mcc_h3
				var _2153_cf25 _dafny.Int = _2152___mcc_h3
				_ = _2153_cf25
				var _2154_cf24 bool = _2151___mcc_h2
				_ = _2154_cf24
				var _2155_v32 _dafny.Map
				_ = _2155_v32
				_2155_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm16(globalState), p3)
				r1 = Companion_Default___.SafeDivisionInt((_2155_v32).Cardinality(), Companion_Default___.SafeModuloInt((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(908)))
				var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
				_ = _index401
				(_2138_v23).ArraySet1((_2140_v25).Fm1(p0, globalState), (_index401).Int())
				var _2156_v33 *C1
				_ = _2156_v33
				var _nw392 *C1 = New_C1_()
				_ = _nw392
				_nw392.Ctor__(p1)
				_2156_v33 = _nw392
				_2156_v33 = _2156_v33
				r3 = true
			} else {
				var _2157___mcc_h4 _dafny.Sequence = _source31.Get_().(D6_DC11).Cf21
				_ = _2157___mcc_h4
				var _2158_cf21 _dafny.Sequence = _2157___mcc_h4
				_ = _2158_cf21
				var _rhs354 _dafny.Int = _dafny.IntOfInt64(853)
				_ = _rhs354
				var _rhs355 bool = ((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)).Cmp(((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)).Times((_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int))) < 0
				_ = _rhs355
				var _rhs356 _dafny.Array = _2138_v23
				_ = _rhs356
				var _rhs357 bool = ((func() _dafny.Int {
					if p0 {
						return (_this).Fm1(p0, globalState)
					}
					return (_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)
				})()).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2158_cf21).Cardinality()), p2)) <= 0
				_ = _rhs357
				var _rhs358 bool = p0
				_ = _rhs358
				r2 = _rhs354
				r3 = _rhs355
				_2138_v23 = _rhs356
				r3 = _rhs357
				r3 = _rhs358
				var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))
				_ = _index402
				(_2138_v23).ArraySet1((func() _dafny.Int {
					if p0 {
						return (_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)
					}
					return (_2138_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2138_v23), 0))).Int()).(_dafny.Int)
				})(), (_index402).Int())
				var _2159_v34 _dafny.MultiSet
				_ = _2159_v34
				_2159_v34 = _dafny.MultiSetOf(p2, p2)
				var _2160_v35 _dafny.MultiSet
				_ = _2160_v35
				_2160_v35 = _2159_v34
				r3 = (_2159_v34).IsProperSubsetOf((_2160_v35).Difference(_2159_v34))
				var _2161_v36 _dafny.Array
				_ = _2161_v36
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_49
				var _nw393 _dafny.Array
				_ = _nw393
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw393 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) bool = (func(_2162_p0 bool) func(_dafny.Int) bool {
						return func(_2163_i1 _dafny.Int) bool {
							return _2162_p0
						}
					})(p0)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw393 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw393).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw393).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_2161_v36 = _nw393
				var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2161_v36), 0))
				_ = _index403
				(_2161_v36).ArraySet1(p0, (_index403).Int())
				var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2161_v36), 0))
				_ = _index404
				(_2161_v36).ArraySet1(p3, (_index404).Int())
			}
		} else {
			if p3 {
				r3 = !(p3)
				var _2164_v37 _dafny.Sequence
				_ = _2164_v37
				_2164_v37 = _dafny.SeqOf(p2)
				var _2165_v38 _dafny.Array
				_ = _2165_v38
				var _nwElement0_89 _dafny.Int = p2
				_ = _nwElement0_89
				var _nw394 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(7))
				_ = _nw394
				(_nw394).ArraySet1(_nwElement0_89, 0)
				(_nw394).ArraySet1(p2, 1)
				(_nw394).ArraySet1((_dafny.Zero).Minus(p2), 2)
				(_nw394).ArraySet1(Companion_Default___.SafeDivisionInt((_this).Fm1(p3, globalState), p2), 3)
				(_nw394).ArraySet1(p2, 4)
				(_nw394).ArraySet1(p2, 5)
				(_nw394).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2164_v37, _2164_v37)).Cardinality()), 6)
				_2165_v38 = _nw394
				var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_2165_v38), 0))
				_ = _index405
				(_2165_v38).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2164_v37, _2164_v37)).Cardinality()), (_index405).Int())
				var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_2165_v38), 0))
				_ = _index406
				(_2165_v38).ArraySet1(Companion_Default___.SafeModuloInt(p2, (Companion_Default___.Fm41(_2164_v37, p0, !(Companion_Default___.Fm5(globalState)), p0, globalState)).Cardinality()), (_index406).Int())
				var _2166_v39 _dafny.Map
				_ = _2166_v39
				_2166_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2165_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_2165_v38), 0))).Int()).(_dafny.Int), _dafny.SeqOf(p2))
				_2166_v39 = (_2166_v39).Update((_dafny.Zero).Minus((_2165_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_2165_v38), 0))).Int()).(_dafny.Int)), _2164_v37)
				var _2167_v40 *C1
				_ = _2167_v40
				var _nw395 *C1 = New_C1_()
				_ = _nw395
				_nw395.Ctor__(p1)
				_2167_v40 = _nw395
				var _2168_v41 T1
				_ = _2168_v41
				var _nw396 *C4 = New_C4_()
				_ = _nw396
				_nw396.Ctor__((func() _dafny.Int {
					if Companion_Default___.Fm5(globalState) {
						return _dafny.IntOfInt64(795)
					}
					return (Companion_D8_.Create_DC17_(p3, _2167_v40, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ycbfarff"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ycbfarff")).Cardinality()))).Uint32(), _dafny.CodePoint('a'))).Cardinality()))).Dtor_cf32()
				})())
				_2168_v41 = _nw396
				_2168_v41 = _2168_v41
				var _2169_v42 _dafny.Array
				_ = _2169_v42
				var _nwElement0_90 bool = true
				_ = _nwElement0_90
				var _nw397 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.One)
				_ = _nw397
				(_nw397).ArraySet1(_nwElement0_90, 0)
				_2169_v42 = _nw397
				var _2170_v43 _dafny.Set
				_ = _2170_v43
				_2170_v43 = _dafny.SetOf(_2169_v42)
				var _2171_v44 *C0
				_ = _2171_v44
				var _nw398 *C0 = New_C0_()
				_ = _nw398
				_nw398.Ctor__((_2170_v43).IsProperSubsetOf(_2170_v43), p0)
				_2171_v44 = _nw398
			} else {
				var _2172_v45 _dafny.Map
				_ = _2172_v45
				_2172_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(331), (_dafny.Zero).Minus(p2))).Cardinality()).Minus(p2))
				var _2173_v46 _dafny.Sequence
				_ = _2173_v46
				_2173_v46 = _dafny.SeqOf(p0)
				var _2174_v47 T2
				_ = _2174_v47
				var _nw399 *C3 = New_C3_()
				_ = _nw399
				_nw399.Ctor__(p2)
				_2174_v47 = _nw399
				var _2175_v48 _dafny.MultiSet
				_ = _2175_v48
				_2175_v48 = _dafny.MultiSetOf(p0)
				var _2176_v49 D5
				_ = _2176_v49
				_2176_v49 = Companion_D5_.Create_DC9_(p0, (_dafny.SetOf(_2174_v47, _2174_v47)).Cardinality(), _2175_v48)
				var _2177_v50 _dafny.Sequence
				_ = _2177_v50
				_2177_v50 = _dafny.SeqOf((_2173_v46).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2173_v46).Cardinality()))).Uint32()).(bool), p0, (_2176_v49).Dtor_cf14(), p3)
				r1 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_2172_v45).Contains((p2).Times(p2)) {
						return (_2172_v45).Get((p2).Times(p2)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_2173_v46).Cardinality())
				})())
				var _2178_v51 _dafny.Set
				_ = _2178_v51
				_2178_v51 = _dafny.SetOf(p2)
				var _2179_v52 _dafny.MultiSet
				_ = _2179_v52
				_2179_v52 = _dafny.MultiSetOf(p2, (_dafny.Zero).Minus(p2))
				var _2180_v53 _dafny.Array
				_ = _2180_v53
				var _nwElement0_91 _dafny.Int = p2
				_ = _nwElement0_91
				var _nw400 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(25))
				_ = _nw400
				(_nw400).ArraySet1(_nwElement0_91, 0)
				(_nw400).ArraySet1((func() _dafny.Int {
					if (_2172_v45).Contains(p2) {
						return (_2172_v45).Get(p2).(_dafny.Int)
					}
					return (_2178_v51).Cardinality()
				})(), 1)
				(_nw400).ArraySet1(p2, 2)
				(_nw400).ArraySet1(p2, 3)
				(_nw400).ArraySet1(_2174_v47.F2(), 4)
				(_nw400).ArraySet1((func() _dafny.Int {
					if (_2179_v52).Contains(_dafny.IntOfUint32((Companion_Default___.Fm15(globalState)).Cardinality())) {
						return (_2179_v52).Multiplicity(_dafny.IntOfUint32((Companion_Default___.Fm15(globalState)).Cardinality()))
					}
					return p2
				})(), 5)
				(_nw400).ArraySet1((_2175_v48).Cardinality(), 6)
				(_nw400).ArraySet1(p2, 7)
				(_nw400).ArraySet1(p2, 8)
				(_nw400).ArraySet1(_2174_v47.F2(), 9)
				(_nw400).ArraySet1((_dafny.Zero).Minus(p2), 10)
				(_nw400).ArraySet1(_2174_v47.F2(), 11)
				(_nw400).ArraySet1((_this).Fm1(p0, globalState), 12)
				(_nw400).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 13)
				(_nw400).ArraySet1(_2174_v47.F2(), 14)
				(_nw400).ArraySet1(p2, 15)
				(_nw400).ArraySet1(_2174_v47.F2(), 16)
				(_nw400).ArraySet1(_dafny.IntOfInt64(267), 17)
				(_nw400).ArraySet1(_2174_v47.F2(), 18)
				(_nw400).ArraySet1(_2174_v47.F2(), 19)
				(_nw400).ArraySet1(p2, 20)
				(_nw400).ArraySet1((_dafny.Zero).Minus(p2), 21)
				(_nw400).ArraySet1(_2174_v47.F2(), 22)
				(_nw400).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 23)
				(_nw400).ArraySet1((_2174_v47).Fm1(p0, globalState), 24)
				_2180_v53 = _nw400
				r1 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _2180_v53)).Cardinality()
				var _2181_v54 _dafny.Sequence
				_ = _2181_v54
				_2181_v54 = _dafny.SeqOf(Companion_Default___.Fm16(globalState))
				var _2182_v55 _dafny.Array
				_ = _2182_v55
				var _nwElement0_92 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _dafny.Companion_Sequence_.Update(_2173_v46, (Companion_Default___.SafeIndex(_2174_v47.F2(), _dafny.IntOfUint32((_2173_v46).Cardinality()))).Uint32(), p0))
				_ = _nwElement0_92
				var _nw401 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(8))
				_ = _nw401
				(_nw401).ArraySet1(_nwElement0_92, 0)
				(_nw401).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2173_v46, _2177_v50), 1)
				(_nw401).ArraySet1((_2181_v54).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_2181_v54).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
				(_nw401).ArraySet1(_2177_v50, 3)
				(_nw401).ArraySet1(_2173_v46, 4)
				(_nw401).ArraySet1(_2177_v50, 5)
				(_nw401).ArraySet1(_2173_v46, 6)
				(_nw401).ArraySet1(_2177_v50, 7)
				_2182_v55 = _nw401
				var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2182_v55), 0))
				_ = _index407
				(_2182_v55).ArraySet1(Companion_Default___.Fm16(globalState), (_index407).Int())
				var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2182_v55), 0))
				_ = _index408
				(_2182_v55).ArraySet1(_2173_v46, (_index408).Int())
				r3 = p0
				var _2183_v56 _dafny.Sequence
				_ = _2183_v56
				_2183_v56 = _dafny.UnicodeSeqOfUtf8Bytes("ndgceyxxq")
				var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2180_v53), 0))
				_ = _index409
				(_2180_v53).ArraySet1((_this).Fm1(p0, globalState), (_index409).Int())
				var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2180_v53), 0))
				_ = _index410
				var _rhs359 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("sxrkeyvv"))
				_ = _rhs359
				var _rhs360 bool = p3
				_ = _rhs360
				var _rhs361 _dafny.Int = _2174_v47.F2()
				_ = _rhs361
				var _rhs362 _dafny.Array = _2180_v53
				_ = _rhs362
				var _rhs363 _dafny.Int = (_this).Fm1((_2174_v47).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("wtikft"), p3, globalState), globalState)
				_ = _rhs363
				var _lhs180 _dafny.Array = _2180_v53
				_ = _lhs180
				var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2180_v53), 0))
				_ = _lhs181
				var _lhs182 T2 = _2174_v47
				_ = _lhs182
				_2183_v56 = _rhs359
				r3 = _rhs360
				(_lhs180).ArraySet1(_rhs361, (_lhs181).Int())
				_2180_v53 = _rhs362
				_lhs182.F2_set_(_rhs363)
			}
			var _2184_v57 _dafny.Array
			_ = _2184_v57
			var _nw402 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(9))
			_ = _nw402
			_2184_v57 = _nw402
			var _2185_v58 _dafny.Map
			_ = _2185_v58
			_2185_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
			var _2186_v59 _dafny.Set
			_ = _2186_v59
			_2186_v59 = _dafny.SetOf((_2185_v58).Cardinality(), p2)
			var _2187_v60 _dafny.Sequence
			_ = _2187_v60
			_2187_v60 = _dafny.SeqOf(_dafny.IntOfInt64(410))
			var _2188_v61 _dafny.Array
			_ = _2188_v61
			var _nwElement0_93 _dafny.Int = p2
			_ = _nwElement0_93
			var _nw403 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_93, nil, _dafny.IntOfInt64(20))
			_ = _nw403
			(_nw403).ArraySet1(_nwElement0_93, 0)
			(_nw403).ArraySet1(p2, 1)
			(_nw403).ArraySet1(p2, 2)
			(_nw403).ArraySet1((_2186_v59).Cardinality(), 3)
			(_nw403).ArraySet1((_dafny.Zero).Minus(p2), 4)
			(_nw403).ArraySet1(p2, 5)
			(_nw403).ArraySet1(p2, 6)
			(_nw403).ArraySet1(p2, 7)
			(_nw403).ArraySet1(p2, 8)
			(_nw403).ArraySet1(_dafny.IntOfUint32((_2187_v60).Cardinality()), 9)
			(_nw403).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())), 10)
			(_nw403).ArraySet1(p2, 11)
			(_nw403).ArraySet1(_dafny.IntOfInt64(934), 12)
			(_nw403).ArraySet1(p2, 13)
			(_nw403).ArraySet1(_dafny.IntOfUint32((_2187_v60).Cardinality()), 14)
			(_nw403).ArraySet1(p2, 15)
			(_nw403).ArraySet1(p2, 16)
			(_nw403).ArraySet1(_dafny.IntOfInt64(-472), 17)
			(_nw403).ArraySet1(p2, 18)
			(_nw403).ArraySet1(_dafny.IntOfInt64(-928), 19)
			_2188_v61 = _nw403
			var _2189_v62 D4
			_ = _2189_v62
			_2189_v62 = Companion_D4_.Create_DC6_(p2, _2188_v61, _dafny.IntOfInt64(994))
			var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_2184_v57), 0))
			_ = _index411
			(_2184_v57).ArraySet1(_2189_v62, (_index411).Int())
			var _2190_v63 _dafny.Array
			_ = _2190_v63
			var _nw404 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw404
			_2190_v63 = _nw404
			var _2191_v64 _dafny.Array
			_ = _2191_v64
			var _nw405 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw405
			_2191_v64 = _nw405
			var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_2190_v63), 0))
			_ = _index412
			(_2190_v63).ArraySet1(_2191_v64, (_index412).Int())
			var _2192_v65 _dafny.Map
			_ = _2192_v65
			_2192_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.CodePoint('n'))
			var _2193_v66 _dafny.Map
			_ = _2193_v66
			_2193_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2192_v65, (func() bool {
				if true {
					return p0
				}
				return p3
			})())
			var _2194_v68 _dafny.Map
			_ = _2194_v68
			_2194_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("fao"))
			var _2195_v69 _dafny.Map
			_ = _2195_v69
			_2195_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Sequence {
				if (_2194_v68).Contains(p0) {
					return (_2194_v68).Get(p0).(_dafny.Sequence)
				}
				return p1
			})())
			var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_2184_v57), 0))
			_ = _index413
			var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_2190_v63), 0))
			_ = _index414
			var _rhs364 D4 = Companion_D4_.Create_DC6_(_dafny.IntOfUint32((_2187_v60).Cardinality()), _2188_v61, p2)
			_ = _rhs364
			var _rhs365 _dafny.Array = _2191_v64
			_ = _rhs365
			var _rhs366 bool = (func() bool {
				if (_2193_v66).Contains((_2192_v65).Merge(func() _dafny.Map {
					var _coll66 = _dafny.NewMapBuilder()
					_ = _coll66
					for _iter71 := _dafny.Iterate((_2195_v69).Keys().Elements()); ; {
						_compr_66, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _2196_v67 _dafny.Sequence
						_2196_v67 = interface{}(_compr_66).(_dafny.Sequence)
						if (_2195_v69).Contains(_2196_v67) {
							_coll66.Add(_2196_v67, _dafny.CodePoint('o'))
						}
					}
					return _coll66.ToMap()
				}())) {
					return (_2193_v66).Get((_2192_v65).Merge(func() _dafny.Map {
						var _coll67 = _dafny.NewMapBuilder()
						_ = _coll67
						for _iter72 := _dafny.Iterate((_2195_v69).Keys().Elements()); ; {
							_compr_67, _ok72 := _iter72()
							if !_ok72 {
								break
							}
							var _2197_v67 _dafny.Sequence
							_2197_v67 = interface{}(_compr_67).(_dafny.Sequence)
							if (_2195_v69).Contains(_2197_v67) {
								_coll67.Add(_2197_v67, _dafny.CodePoint('o'))
							}
						}
						return _coll67.ToMap()
					}())).(bool)
				}
				return p3
			})()
			_ = _rhs366
			var _rhs367 _dafny.Int = _dafny.IntOfInt64(206)
			_ = _rhs367
			var _lhs183 _dafny.Array = _2184_v57
			_ = _lhs183
			var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_2184_v57), 0))
			_ = _lhs184
			var _lhs185 _dafny.Array = _2190_v63
			_ = _lhs185
			var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_2190_v63), 0))
			_ = _lhs186
			(_lhs183).ArraySet1(_rhs364, (_lhs184).Int())
			(_lhs185).ArraySet1(_rhs365, (_lhs186).Int())
			r3 = _rhs366
			r2 = _rhs367
			var _2198_v70 _dafny.CodePoint
			_ = _2198_v70
			_2198_v70 = _dafny.CodePoint('l')
			_2198_v70 = _2198_v70
			var _2199_v71 _dafny.MultiSet
			_ = _2199_v71
			_2199_v71 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()))
			var _2200_v72 _dafny.Sequence
			_ = _2200_v72
			_2200_v72 = _dafny.SeqOf(_2199_v71)
			var _2201_v73 _dafny.Map
			_ = _2201_v73
			_2201_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_2190_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_2190_v63), 0))).Int())), !((_2199_v71).IsSubsetOf((_2200_v72).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2200_v72).Cardinality()))).Uint32()).(_dafny.MultiSet))))
			_2201_v73 = _2201_v73
			r2 = (p2).Times(_dafny.IntOfInt64(119))
		}
		var _2202_v74 _dafny.Map
		_ = _2202_v74
		_2202_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(66), p3)
		var _2203_v75 _dafny.Map
		_ = _2203_v75
		_2203_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2202_v74)
		var _2204_v76 D5
		_ = _2204_v76
		_2204_v76 = Companion_D5_.Create_DC10_(false, p2, (_2203_v75).Update(p1, _2202_v74), p0)
		var _hi13 _dafny.Int = (_2204_v76).Dtor_cf18()
		_ = _hi13
		for _2205_i2 := p2; _2205_i2.Cmp(_hi13) < 0; _2205_i2 = _2205_i2.Plus(_dafny.One) {
			r3 = p0
			var _2206_v77 _dafny.Array
			_ = _2206_v77
			var _nw406 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw406
			_2206_v77 = _nw406
			_2206_v77 = _2206_v77
			if !(p3) {
				var _2207_v78 _dafny.Set
				_ = _2207_v78
				_2207_v78 = _dafny.SetOf(_2205_i2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}(func(_2208_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))).Cardinality()), _2205_i2)
				var _rhs368 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), (_2207_v78).Cardinality())
				_ = _rhs368
				var _rhs369 _dafny.Int = ((func() _dafny.Int {
					if p0 {
						return _2205_i2
					}
					return p2
				})()).Plus(_dafny.IntOfInt64(-664))
				_ = _rhs369
				r1 = _rhs368
				r2 = _rhs369
				r2 = _2205_i2
				r3 = false
				r2 = p2
				var _2209_v79 _dafny.MultiSet
				_ = _2209_v79
				_2209_v79 = _dafny.MultiSetOf(p0, p3)
				var _2210_v80 _dafny.Map
				_ = _2210_v80
				_2210_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2205_i2).Plus((func() _dafny.Int {
					if (_2209_v79).Contains(p0) {
						return (_2209_v79).Multiplicity(p0)
					}
					return _2205_i2
				})()), p2)
				var _2211_v81 _dafny.Map
				_ = _2211_v81
				_2211_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2209_v79)
				_2210_v80 = (_2210_v80).Update((func() _dafny.Int {
					if p3 {
						return p2
					}
					return _2205_i2
				})(), ((func() _dafny.MultiSet {
					if Companion_Default___.Fm5(globalState) {
						return _2209_v79
					}
					return (func() _dafny.MultiSet {
						if (_2211_v81).Contains(p3) {
							return (_2211_v81).Get(p3).(_dafny.MultiSet)
						}
						return _2209_v79
					})()
				})()).Cardinality())
			} else {
				var _2212_v82 _dafny.Sequence
				_ = _2212_v82
				_2212_v82 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), _2205_i2)
				var _2213_v83 _dafny.Map
				_ = _2213_v83
				_2213_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2212_v82, p0)
				var _2214_v84 _dafny.Map
				_ = _2214_v84
				_2214_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("c"), !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-663)), (_2213_v83).Cardinality(), (_this).Fm1(true, globalState)), _2212_v82))
				_2214_v84 = (_2214_v84).Update(p1, p3)
				var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))
				_ = _index415
				(_2206_v77).ArraySet1(!(p0), (_index415).Int())
				var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))
				_ = _index416
				(_2206_v77).ArraySet1(!(p3) || (p3), (_index416).Int())
				var _2215_v85 *C1
				_ = _2215_v85
				var _nw407 *C1 = New_C1_()
				_ = _nw407
				_nw407.Ctor__(_dafny.Companion_Sequence_.Concatenate(p1, p1))
				_2215_v85 = _nw407
				var _2216_v86 _dafny.Array
				_ = _2216_v86
				var _nwElement0_94 bool = false
				_ = _nwElement0_94
				var _nw408 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_94, nil, _dafny.IntOfInt64(26))
				_ = _nw408
				(_nw408).ArraySet1(_nwElement0_94, 0)
				(_nw408).ArraySet1(p3, 1)
				(_nw408).ArraySet1(p0, 2)
				(_nw408).ArraySet1(p0, 3)
				(_nw408).ArraySet1(p3, 4)
				(_nw408).ArraySet1(!(Companion_Default___.Fm5(globalState)), 5)
				(_nw408).ArraySet1(p3, 6)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 7)
				(_nw408).ArraySet1(p3, 8)
				(_nw408).ArraySet1(p0, 9)
				(_nw408).ArraySet1(p3, 10)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 11)
				(_nw408).ArraySet1(p0, 12)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 13)
				(_nw408).ArraySet1(true, 14)
				(_nw408).ArraySet1(p3, 15)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 16)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 17)
				(_nw408).ArraySet1(p0, 18)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 19)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 20)
				(_nw408).ArraySet1(p3, 21)
				(_nw408).ArraySet1(p3, 22)
				(_nw408).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), 23)
				(_nw408).ArraySet1(p0, 24)
				(_nw408).ArraySet1(p3, 25)
				_2216_v86 = _nw408
				var _2217_v87 _dafny.Map
				_ = _2217_v87
				_2217_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2216_v86, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pgdsyug"), (_2215_v85).F3()))
				var _2218_v88 _dafny.CodePoint
				_ = _2218_v88
				_2218_v88 = _dafny.CodePoint('f')
				var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))
				_ = _index417
				var _rhs370 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.CodePoint {
					if false {
						return _2218_v88
					}
					return _2218_v88
				})())).Cardinality())
				_ = _rhs370
				var _rhs371 _dafny.Int = _2205_i2
				_ = _rhs371
				var _rhs372 _dafny.Int = p2
				_ = _rhs372
				var _rhs373 _dafny.Map = (func() _dafny.Map {
					if !((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool)) {
						return (_2217_v87).Merge(_2217_v87)
					}
					return _2217_v87
				})()
				_ = _rhs373
				var _rhs374 bool = p3
				_ = _rhs374
				var _lhs187 _dafny.Array = _2206_v77
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))
				_ = _lhs188
				r2 = _rhs370
				r2 = _rhs371
				r2 = _rhs372
				_2217_v87 = _rhs373
				(_lhs187).ArraySet1(_rhs374, (_lhs188).Int())
				var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))
				_ = _index418
				(_2206_v77).ArraySet1((_2206_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_2206_v77), 0))).Int()).(bool), (_index418).Int())
			}
			r1 = (_this).Fm1(p3, globalState)
		}
		var _2219_v89 D15
		_ = _2219_v89
		_2219_v89 = Companion_D15_.Create_DC30_(true)
		r1 = (_this).Fm1((_2219_v89).Dtor_cf47(), globalState)
		var _2220_v90 _dafny.Array
		_ = _2220_v90
		var _nw409 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw409
		_2220_v90 = _nw409
		var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))
		_ = _index419
		(_2220_v90).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(138), _dafny.IntOfInt64(795)), (_index419).Int())
		var _2221_v91 _dafny.Sequence
		_ = _2221_v91
		_2221_v91 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()))
		var _2222_v92 _dafny.Map
		_ = _2222_v92
		_2222_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2221_v91, p3)
		var _index420 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))
		_ = _index420
		(_2220_v90).ArraySet1(((_2222_v92).Cardinality()).Minus(p2), (_index420).Int())
		var _2223_v93 _dafny.Array
		_ = _2223_v93
		var _nw410 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
		_ = _nw410
		_2223_v93 = _nw410
		var _2224_v95 _dafny.Sequence
		_ = _2224_v95
		_2224_v95 = _dafny.SeqOf(p1, _dafny.UnicodeSeqOfUtf8Bytes("sdflwo"))
		var _2225_v96 D15
		_ = _2225_v96
		_2225_v96 = Companion_D15_.Create_DC29_(func() _dafny.Map {
			var _coll68 = _dafny.NewMapBuilder()
			_ = _coll68
			for _iter73 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_2224_v95, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2224_v95).Cardinality()))).Uint32(), (_2224_v95).Select((Companion_Default___.SafeIndex((_2220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2224_v95).Cardinality()))).Uint32()).(_dafny.Sequence))).Elements()); ; {
				_compr_68, _ok73 := _iter73()
				if !_ok73 {
					break
				}
				var _2226_v94 _dafny.Sequence
				_2226_v94 = interface{}(_compr_68).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_2224_v95, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2224_v95).Cardinality()))).Uint32(), (_2224_v95).Select((Companion_Default___.SafeIndex((_2220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2224_v95).Cardinality()))).Uint32()).(_dafny.Sequence)), _2226_v94) {
					_coll68.Add(_2226_v94, p1)
				}
			}
			return _coll68.ToMap()
		}())
		var _pat_let_tv47 = _2220_v90
		_ = _pat_let_tv47
		var _pat_let_tv48 = _2220_v90
		_ = _pat_let_tv48
		var _pat_let_tv49 = _2220_v90
		_ = _pat_let_tv49
		var _pat_let_tv50 = _2220_v90
		_ = _pat_let_tv50
		var _pat_let_tv51 = _2220_v90
		_ = _pat_let_tv51
		var _pat_let_tv52 = _2220_v90
		_ = _pat_let_tv52
		var _index421 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))
		_ = _index421
		var _rhs375 _dafny.Array = _2223_v93
		_ = _rhs375
		var _rhs376 _dafny.Int = Companion_Default___.SafeDivisionInt((_2220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))).Int()).(_dafny.Int), p2)
		_ = _rhs376
		var _rhs377 _dafny.Int = func(_source32 D15) _dafny.Int {
			if _source32.Is_DC30() {
				var _2227___mcc_h5 bool = _source32.Get_().(D15_DC30).Cf47
				_ = _2227___mcc_h5
				var _2228_cf47 bool = _2227___mcc_h5
				_ = _2228_cf47
				return Companion_Default___.SafeDivisionInt((_pat_let_tv48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_pat_let_tv47), 0))).Int()).(_dafny.Int), (_pat_let_tv50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_pat_let_tv49), 0))).Int()).(_dafny.Int))
			} else {
				var _2229___mcc_h6 _dafny.Map = _source32.Get_().(D15_DC29).Cf46
				_ = _2229___mcc_h6
				var _2230_cf46 _dafny.Map = _2229___mcc_h6
				_ = _2230_cf46
				return (_pat_let_tv52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_pat_let_tv51), 0))).Int()).(_dafny.Int)
			}
		}(_2225_v96)
		_ = _rhs377
		var _rhs378 _dafny.Int = _dafny.IntOfUint32((p1).Cardinality())
		_ = _rhs378
		var _lhs189 _dafny.Array = _2220_v90
		_ = _lhs189
		var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))
		_ = _lhs190
		_2223_v93 = _rhs375
		r1 = _rhs376
		(_lhs189).ArraySet1(_rhs377, (_lhs190).Int())
		r1 = _rhs378
		r0 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if !(p3) {
				return _2224_v95
			}
			return _dafny.SeqOf(p1, p1, p1, p1, p1)
		})(), _dafny.Companion_Sequence_.Concatenate(_2224_v95, _2224_v95))
		var _2231_v97 _dafny.MultiSet
		_ = _2231_v97
		_2231_v97 = _dafny.MultiSetOf(p0)
		r1 = (func() _dafny.Int {
			if (_2231_v97).IsProperSubsetOf(_2231_v97) {
				return _dafny.IntOfUint32((p1).Cardinality())
			}
			return p2
		})()
		r2 = (_2221_v91).Select((Companion_Default___.SafeIndex((_2220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2221_v91).Cardinality()))).Uint32()).(_dafny.Int)
		var _2232_v98 _dafny.CodePoint
		_ = _2232_v98
		_2232_v98 = _dafny.CodePoint('b')
		var _2233_v99 D7
		_ = _2233_v99
		_2233_v99 = Companion_D7_.Create_DC15_((_2220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_2220_v90), 0))).Int()).(_dafny.Int), _2232_v98)
		var _pat_let_tv53 = p3
		_ = _pat_let_tv53
		var _pat_let_tv54 = _2220_v90
		_ = _pat_let_tv54
		var _pat_let_tv55 = _2220_v90
		_ = _pat_let_tv55
		r3 = func(_source33 D7) bool {
			if _source33.Is_DC15() {
				var _2234___mcc_h7 _dafny.Int = _source33.Get_().(D7_DC15).Cf27
				_ = _2234___mcc_h7
				var _2235___mcc_h8 _dafny.CodePoint = _source33.Get_().(D7_DC15).Cf28
				_ = _2235___mcc_h8
				var _2236_cf28 _dafny.CodePoint = _2235___mcc_h8
				_ = _2236_cf28
				var _2237_cf27 _dafny.Int = _2234___mcc_h7
				_ = _2237_cf27
				return _pat_let_tv53
			} else {
				var _2238___mcc_h9 _dafny.Sequence = _source33.Get_().(D7_DC14).Cf26
				_ = _2238___mcc_h9
				var _2239_cf26 _dafny.Sequence = _2238___mcc_h9
				_ = _2239_cf26
				return false
			}
		}(func(_pat_let43_0 D7) D7 {
			return func(_2240_dt__update__tmp_h0 D7) D7 {
				return func(_pat_let44_0 _dafny.Int) D7 {
					return func(_2241_dt__update_hcf27_h0 _dafny.Int) D7 {
						return Companion_D7_.Create_DC15_(_2241_dt__update_hcf27_h0, (_2240_dt__update__tmp_h0).Dtor_cf28())
					}(_pat_let44_0)
				}((_pat_let_tv55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_pat_let_tv54), 0))).Int()).(_dafny.Int))
			}(_pat_let43_0)
		}(_2233_v99))
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
func (_this *C7) Fm0(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(462))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg113 _dafny.Int) interface{} {
				return coer113(arg113)
			}
		}(func(_2242_i0 _dafny.Int) _dafny.Int {
			return (_dafny.MultiSetOf(true)).Cardinality()
		})), _dafny.IntOfInt64(711))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(638))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg114 _dafny.Int) interface{} {
				return coer114(arg114)
			}
		}(func(_2243_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpdthucj")).Cardinality())
		})), _dafny.IntOfInt64(618))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(280), (func() _dafny.Map {
			var _coll69 = _dafny.NewMapBuilder()
			_ = _coll69
			for _iter74 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-18), _dafny.IntOfInt64(784))); ; {
				_compr_69, _ok74 := _iter74()
				if !_ok74 {
					break
				}
				var _2244_v0 _dafny.Int
				_2244_v0 = interface{}(_compr_69).(_dafny.Int)
				if ((_dafny.IntOfInt64(-18)).Cmp(_2244_v0) <= 0) && ((_2244_v0).Cmp(_dafny.IntOfInt64(784)) < 0) {
					_coll69.Add(Companion_Default___.SafeDivisionInt(_2244_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), false)
				}
			}
			return _coll69.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(438), _dafny.IntOfInt64(-11), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg115 _dafny.Int) interface{} {
				return coer115(arg115)
			}
		}(func(_2245_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-503))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg116 _dafny.Int) interface{} {
				return coer116(arg116)
			}
		}(func(_2246_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()))))
	}
}
func (_this *C7) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(((_dafny.SetOf(!(!(false)), false, false)).Intersection(_dafny.SetOf(true, (Companion_D0_.Create_DC0_(true)).Dtor_cf0()))).Cardinality(), _dafny.IntOfInt64(312))
	}
}
func (_this *C7) M3(p0 _dafny.MultiSet, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _2247_v0 _dafny.Sequence
		_ = _2247_v0
		_2247_v0 = _dafny.UnicodeSeqOfUtf8Bytes("hsn")
		_2247_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2247_v0, Companion_Default___.Fm2(globalState)), _2247_v0)
		var _2248_v1 _dafny.Int
		_ = _2248_v1
		_2248_v1 = _dafny.IntOfInt64(583)
		r1 = (_dafny.Zero).Minus(_2248_v1)
		var _2249_v2 bool
		_ = _2249_v2
		_2249_v2 = true
		_2249_v2 = _2249_v2
		_2249_v2 = _2249_v2
		r0 = _2248_v1
		var _2250_v3 _dafny.Map
		_ = _2250_v3
		_2250_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v2, p3)
		var _2251_v4 _dafny.Sequence
		_ = _2251_v4
		_2251_v4 = _dafny.SeqOf(_2249_v2)
		var _2252_v5 _dafny.Sequence
		_ = _2252_v5
		_2252_v5 = _dafny.SeqOf(p2, (func() bool {
			if (_2250_v3).Contains(true) {
				return (_2250_v3).Get(true).(bool)
			}
			return (_2251_v4).Select((Companion_Default___.SafeIndex((_this).Fm1(p3, globalState), _dafny.IntOfUint32((_2251_v4).Cardinality()))).Uint32()).(bool)
		})(), p2)
		var _2253_v6 _dafny.Array
		_ = _2253_v6
		var _nwElement0_95 bool = (p2) == (_2249_v2)
		_ = _nwElement0_95
		var _nw411 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_95, nil, _dafny.IntOfInt64(22))
		_ = _nw411
		(_nw411).ArraySet1(_nwElement0_95, 0)
		(_nw411).ArraySet1(p3, 1)
		(_nw411).ArraySet1(_2249_v2, 2)
		(_nw411).ArraySet1(true, 3)
		(_nw411).ArraySet1(_2249_v2, 4)
		(_nw411).ArraySet1((_2251_v4).Select((Companion_Default___.SafeIndex(_2248_v1, _dafny.IntOfUint32((_2251_v4).Cardinality()))).Uint32()).(bool), 5)
		(_nw411).ArraySet1(_2249_v2, 6)
		(_nw411).ArraySet1(!(_2249_v2), 7)
		(_nw411).ArraySet1(p2, 8)
		(_nw411).ArraySet1(!(_2249_v2), 9)
		(_nw411).ArraySet1(true, 10)
		(_nw411).ArraySet1(!(p3), 11)
		(_nw411).ArraySet1((_2249_v2) && (p2), 12)
		(_nw411).ArraySet1(_2249_v2, 13)
		(_nw411).ArraySet1(p2, 14)
		(_nw411).ArraySet1(Companion_Default___.Fm3(true, _2248_v1, _2248_v1, globalState), 15)
		(_nw411).ArraySet1(!_dafny.Companion_Sequence_.Equal(_2251_v4, _2251_v4), 16)
		(_nw411).ArraySet1((_2248_v1).Cmp(_2248_v1) < 0, 17)
		(_nw411).ArraySet1(!((func() bool {
			if (func() bool {
				if (_2250_v3).Contains(p3) {
					return (_2250_v3).Get(p3).(bool)
				}
				return (func() bool {
					if (_2250_v3).Contains(p2) {
						return (_2250_v3).Get(p2).(bool)
					}
					return p2
				})()
			})() {
				return p2
			}
			return false
		})()), 18)
		(_nw411).ArraySet1(_2249_v2, 19)
		(_nw411).ArraySet1(p2, 20)
		(_nw411).ArraySet1(_2249_v2, 21)
		_2253_v6 = _nw411
		var _2254_v7 _dafny.Sequence
		_ = _2254_v7
		_2254_v7 = _dafny.SeqOf(_2248_v1, (p0).Cardinality(), _2248_v1)
		var _2255_v8 _dafny.Set
		_ = _2255_v8
		_2255_v8 = _dafny.SetOf(_2248_v1, _dafny.IntOfInt64(799))
		var _2256_v9 _dafny.Set
		_ = _2256_v9
		_2256_v9 = _dafny.SetOf(_2255_v8)
		var _2257_v10 _dafny.Sequence
		_ = _2257_v10
		_2257_v10 = _dafny.SeqOf(_2255_v8, _2255_v8)
		var _2258_v12 _dafny.Sequence
		_ = _2258_v12
		_2258_v12 = _dafny.SeqOf(_2256_v9, _dafny.SetOf(_2255_v8), func() _dafny.Set {
			var _coll70 = _dafny.NewBuilder()
			_ = _coll70
			for _iter75 := _dafny.Iterate((_2257_v10).Elements()); ; {
				_compr_70, _ok75 := _iter75()
				if !_ok75 {
					break
				}
				var _2259_v11 _dafny.Set
				_2259_v11 = interface{}(_compr_70).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_2257_v10, _2259_v11) {
					_coll70.Add(_2259_v11)
				}
			}
			return _coll70.ToSet()
		}(), _2256_v9, _2256_v9)
		var _2260_v13 _dafny.CodePoint
		_ = _2260_v13
		_2260_v13 = _dafny.CodePoint('t')
		var _2261_v14 _dafny.Map
		_ = _2261_v14
		_2261_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2260_v13, p3)
		var _2262_v15 D0
		_ = _2262_v15
		_2262_v15 = Companion_D0_.Create_DC1_(_2249_v2, p3, _2260_v13, (_dafny.Zero).Minus(_2248_v1))
		var _2263_v16 _dafny.Map
		_ = _2263_v16
		_2263_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2262_v15, _2249_v2)
		var _nwElement0_96 bool = (p2) || (p2)
		_ = _nwElement0_96
		var _nw412 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_96, nil, _dafny.IntOfInt64(24))
		_ = _nw412
		(_nw412).ArraySet1(_nwElement0_96, 0)
		(_nw412).ArraySet1(!(_2249_v2) || (_2249_v2), 1)
		(_nw412).ArraySet1(p2, 2)
		(_nw412).ArraySet1(_2249_v2, 3)
		(_nw412).ArraySet1(_2249_v2, 4)
		(_nw412).ArraySet1(p3, 5)
		(_nw412).ArraySet1((func() bool {
			if (_2250_v3).Contains(true) {
				return (_2250_v3).Get(true).(bool)
			}
			return p2
		})(), 6)
		(_nw412).ArraySet1(_2249_v2, 7)
		(_nw412).ArraySet1(true, 8)
		(_nw412).ArraySet1(p3, 9)
		(_nw412).ArraySet1(false, 10)
		(_nw412).ArraySet1(p2, 11)
		(_nw412).ArraySet1(true, 12)
		(_nw412).ArraySet1(_2249_v2, 13)
		(_nw412).ArraySet1(false, 14)
		(_nw412).ArraySet1(!_dafny.Companion_Sequence_.Equal(_2254_v7, _2254_v7), 15)
		(_nw412).ArraySet1(_2249_v2, 16)
		(_nw412).ArraySet1(_2249_v2, 17)
		(_nw412).ArraySet1(((_2258_v12).Select((Companion_Default___.SafeIndex((_2261_v14).Cardinality(), _dafny.IntOfUint32((_2258_v12).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf(_dafny.SetOf(_2255_v8)), 18)
		(_nw412).ArraySet1(!(_2263_v16).Contains(Companion_D0_.Create_DC1_(_2249_v2, _2249_v2, _2260_v13, _2248_v1)), 19)
		(_nw412).ArraySet1((_2255_v8).IsSubsetOf(_2255_v8), 20)
		(_nw412).ArraySet1(p2, 21)
		(_nw412).ArraySet1(p3, 22)
		(_nw412).ArraySet1(p3, 23)
		_2253_v6 = _nw412
		r0 = (_2248_v1).Times(_2248_v1)
		var _2264_v17 _dafny.Map
		_ = _2264_v17
		_2264_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2253_v6, _dafny.CodePoint('r'))
		var _2265_v18 _dafny.MultiSet
		_ = _2265_v18
		_2265_v18 = _dafny.MultiSetOf(_2260_v13, _2260_v13, (func() _dafny.CodePoint {
			if (_2264_v17).Contains(_2253_v6) {
				return (_2264_v17).Get(_2253_v6).(_dafny.CodePoint)
			}
			return _2260_v13
		})())
		r1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_2265_v18).Cardinality()), Companion_Default___.SafeModuloInt((_this).Fm1(_2249_v2, globalState), _2248_v1))
		return r0, r1
	}
}
func (_this *C7) M1(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _2266_v1 _dafny.MultiSet
		_ = _2266_v1
		_2266_v1 = _dafny.MultiSetOf(p0)
		if (func() _dafny.Map {
			var _coll71 = _dafny.NewMapBuilder()
			_ = _coll71
			for _iter76 := _dafny.Iterate((_2266_v1).Elements()); ; {
				_compr_71, _ok76 := _iter76()
				if !_ok76 {
					break
				}
				var _2267_v0 _dafny.Int
				_2267_v0 = interface{}(_compr_71).(_dafny.Int)
				if (_2266_v1).Contains(_2267_v0) {
					_coll71.Add((_2267_v0).Times(((_2266_v1).Update(p0, Companion_Default___.Abs(p0))).Cardinality()), p0)
				}
			}
			return _coll71.ToMap()
		}()).Contains(p0) {
			var _2268_v2 _dafny.CodePoint
			_ = _2268_v2
			_2268_v2 = _dafny.CodePoint('x')
			var _2269_v3 _dafny.Set
			_ = _2269_v3
			_2269_v3 = _dafny.SetOf(p3, p2, p2)
			var _2270_v4 D0
			_ = _2270_v4
			_2270_v4 = Companion_D0_.Create_DC1_(p2, p3, _2268_v2, (_2269_v3).Cardinality())
			var _2271_v5 _dafny.MultiSet
			_ = _2271_v5
			_2271_v5 = _dafny.MultiSetOf(_2270_v4, _2270_v4)
			var _2272_v6 _dafny.Map
			_ = _2272_v6
			_2272_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_2271_v5).Update(_2270_v4, Companion_Default___.Abs(p0)))
			var _2273_v7 _dafny.Sequence
			_ = _2273_v7
			_2273_v7 = _dafny.SeqOf(p0, p0)
			var _2274_v8 _dafny.Sequence
			_ = _2274_v8
			_2274_v8 = _dafny.SeqOf(Companion_Default___.Fm4(p0, _dafny.IntOfUint32((_2273_v7).Cardinality()), p2, p2, globalState))
			_2272_v6 = (_2272_v6).Update(p2, _dafny.MultiSetFromSeq(_2274_v8))
			var _2275_v9 _dafny.Array
			_ = _2275_v9
			var _out49 _dafny.Array
			_ = _out49
			_out49 = Companion_Default___.M0((!(p3)) == (p2), p3, globalState)
			_2275_v9 = _out49
			r0 = p3
			var _2276_v10 *C0
			_ = _2276_v10
			var _nw413 *C0 = New_C0_()
			_ = _nw413
			_nw413.Ctor__(p3, p3)
			_2276_v10 = _nw413
			var _2277_v11 _dafny.Array
			_ = _2277_v11
			var _nw414 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw414
			_2277_v11 = _nw414
			var _index422 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_2277_v11), 0))
			_ = _index422
			(_2277_v11).ArraySet1(p2, (_index422).Int())
			var _index423 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_2277_v11), 0))
			_ = _index423
			(_2277_v11).ArraySet1(p2, (_index423).Int())
		} else {
			var _2278_v12 _dafny.Sequence
			_ = _2278_v12
			_2278_v12 = _dafny.UnicodeSeqOfUtf8Bytes("acq")
			var _2279_v13 _dafny.Set
			_ = _2279_v13
			_2279_v13 = _dafny.SetOf(p3)
			var _2280_v14 _dafny.Map
			_ = _2280_v14
			_2280_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if p3 {
					return _2278_v12
				}
				return _2278_v12
			})(), (_dafny.SetOf(p2)).Intersection(_2279_v13))
			var _2281_v15 _dafny.CodePoint
			_ = _2281_v15
			_2281_v15 = _dafny.CodePoint('c')
			_2280_v14 = (_2280_v14).Update(_dafny.Companion_Sequence_.Update(_2278_v12, (Companion_Default___.SafeIndex((_this).Fm1(p2, globalState), _dafny.IntOfUint32((_2278_v12).Cardinality()))).Uint32(), _2281_v15), _2279_v13)
			var _2282_v16 _dafny.Sequence
			_ = _2282_v16
			_2282_v16 = _dafny.SeqOf(p0)
			var _2283_v17 _dafny.Map
			_ = _2283_v17
			_2283_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2266_v1, p3)
			var _2284_v18 _dafny.Set
			_ = _2284_v18
			_2284_v18 = _dafny.SetOf((_this).Fm1(p3, globalState), p0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_2282_v16).Cardinality())), (_2283_v17).Cardinality())
			r2 = Companion_Default___.SafeModuloInt(p0, ((_2284_v18).Difference(_2284_v18)).Cardinality())
			var _2285_v19 _dafny.Map
			_ = _2285_v19
			_2285_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _2286_v20 D6
			_ = _2286_v20
			_2286_v20 = Companion_D6_.Create_DC13_(p2, p0)
			var _2287_v21 _dafny.Array
			_ = _2287_v21
			var _nwElement0_97 bool = p3
			_ = _nwElement0_97
			var _nw415 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_97, nil, _dafny.IntOfInt64(20))
			_ = _nw415
			(_nw415).ArraySet1(_nwElement0_97, 0)
			(_nw415).ArraySet1(p3, 1)
			(_nw415).ArraySet1((func() bool {
				if (_2285_v19).Contains(!(p2)) {
					return (_2285_v19).Get(!(p2)).(bool)
				}
				return p3
			})(), 2)
			(_nw415).ArraySet1(p2, 3)
			(_nw415).ArraySet1(Companion_Default___.Fm5(globalState), 4)
			(_nw415).ArraySet1((func() bool {
				if (_2285_v19).Contains(true) {
					return (_2285_v19).Get(true).(bool)
				}
				return p3
			})(), 5)
			(_nw415).ArraySet1(p2, 6)
			(_nw415).ArraySet1(p2, 7)
			(_nw415).ArraySet1((p0).Cmp(p0) != 0, 8)
			(_nw415).ArraySet1(p2, 9)
			(_nw415).ArraySet1(!(p2), 10)
			(_nw415).ArraySet1(p3, 11)
			(_nw415).ArraySet1(false, 12)
			(_nw415).ArraySet1(p3, 13)
			(_nw415).ArraySet1((func() bool {
				if p2 {
					return p3
				}
				return !(p2)
			})(), 14)
			(_nw415).ArraySet1(p3, 15)
			(_nw415).ArraySet1(!(p2), 16)
			(_nw415).ArraySet1(Companion_Default___.Fm3(p3, p0, p0, globalState), 17)
			(_nw415).ArraySet1(p3, 18)
			(_nw415).ArraySet1((_2286_v20).Dtor_cf24(), 19)
			_2287_v21 = _nw415
			var _2288_v22 _dafny.Map
			_ = _2288_v22
			_2288_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.Companion_Sequence_.Concatenate(_2278_v12, _dafny.SeqOf(_dafny.CodePoint('k'), _2281_v15, _2281_v15)))
			var _rhs379 _dafny.Array = _2287_v21
			_ = _rhs379
			var _rhs380 _dafny.Int = (_dafny.Zero).Minus((_2288_v22).Cardinality())
			_ = _rhs380
			_2287_v21 = _rhs379
			r2 = _rhs380
			var _2289_v23 *C1
			_ = _2289_v23
			var _nw416 *C1 = New_C1_()
			_ = _nw416
			_nw416.Ctor__(_2278_v12)
			_2289_v23 = _nw416
			var _2290_v24 _dafny.MultiSet
			_ = _2290_v24
			_2290_v24 = _dafny.MultiSetOf(_2289_v23, _2289_v23)
			var _2291_v25 _dafny.Map
			_ = _2291_v25
			_2291_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_2290_v24, _2290_v24, _2290_v24), _dafny.Companion_Sequence_.Concatenate(_2282_v16, _2282_v16))
			_2291_v25 = (_2291_v25).Update(_dafny.SetOf((_2290_v24).Update(_2289_v23, Companion_Default___.Abs(p0))), _2282_v16)
			if p3 {
				r2 = p0
				var _2292_v26 T1
				_ = _2292_v26
				var _nw417 *C4 = New_C4_()
				_ = _nw417
				_nw417.Ctor__(p0)
				_2292_v26 = _nw417
				var _2293_v27 _dafny.Set
				_ = _2293_v27
				_2293_v27 = _dafny.SetOf(_2292_v26)
				var _2294_v28 _dafny.Set
				_ = _2294_v28
				_2294_v28 = _2293_v27
				var _2295_v29 _dafny.Sequence
				_ = _2295_v29
				_2295_v29 = _dafny.SeqOf(_2293_v27, (_dafny.SetOf(_2292_v26)).Union(_2293_v27), (_2294_v28))
				_2295_v29 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(_2292_v26, _2292_v26, _2292_v26, _2292_v26), _2293_v27), _2295_v29)
				r0 = (func() bool {
					if (_2285_v19).Contains(Companion_Default___.Fm5(globalState)) {
						return (_2285_v19).Get(Companion_Default___.Fm5(globalState)).(bool)
					}
					return true
				})()
				var _2296_v30 _dafny.Map
				_ = _2296_v30
				_2296_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v21, _2278_v12)
				_2296_v30 = _2296_v30
				var _2297_v31 D14
				_ = _2297_v31
				_2297_v31 = Companion_D14_.Create_DC28_()
				var _2298_v32 _dafny.Sequence
				_ = _2298_v32
				_2298_v32 = _dafny.SeqOf(_2297_v31, _2297_v31)
				var _index424 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index424
				(_2287_v21).ArraySet1(p3, (_index424).Int())
				var _2299_v33 _dafny.Array
				_ = _2299_v33
				var _nw418 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw418
				_2299_v33 = _nw418
				var _index425 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_2299_v33), 0))
				_ = _index425
				(_2299_v33).ArraySet1(_2287_v21, (_index425).Int())
				var _2300_v34 _dafny.Sequence
				_ = _2300_v34
				_2300_v34 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(321))).Uint32(), func(coer117 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}((func(_2301_v31 D14) func(_dafny.Int) D14 {
					return func(_2302_i0 _dafny.Int) D14 {
						return _2301_v31
					}
				})(_2297_v31))), _2298_v32))
				var _2303_v35 _dafny.Map
				_ = _2303_v35
				_2303_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
				var _2304_v36 _dafny.Map
				_ = _2304_v36
				_2304_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2303_v35).Contains(p3), _2287_v21)
				var _index426 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index426
				var _index427 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_2299_v33), 0))
				_ = _index427
				var _rhs381 _dafny.Sequence = (_2300_v34).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2300_v34).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs381
				var _rhs382 bool = p3
				_ = _rhs382
				var _rhs383 _dafny.Array = (func() _dafny.Array {
					if (_2304_v36).Contains(p2) {
						return (_2304_v36).Get(p2).(_dafny.Array)
					}
					return _2287_v21
				})()
				_ = _rhs383
				var _rhs384 bool = !(p3)
				_ = _rhs384
				var _lhs191 _dafny.Array = _2287_v21
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2287_v21), 0))
				_ = _lhs192
				var _lhs193 _dafny.Array = _2299_v33
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_2299_v33), 0))
				_ = _lhs194
				_2298_v32 = _rhs381
				(_lhs191).ArraySet1(_rhs382, (_lhs192).Int())
				(_lhs193).ArraySet1(_rhs383, (_lhs194).Int())
				r0 = _rhs384
			} else {
				var _2305_v37 _dafny.Map
				_ = _2305_v37
				_2305_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.SafeModuloInt(p0, p0))
				_2305_v37 = (_2305_v37).Update(_dafny.IntOfInt64(241), p0)
				var _2306_v38 _dafny.MultiSet
				_ = _2306_v38
				_2306_v38 = _dafny.MultiSetOf(p3)
				var _index428 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index428
				(_2287_v21).ArraySet1((Companion_Default___.Fm38(p2, globalState)).IsSubsetOf(_2306_v38), (_index428).Int())
				var _index429 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index429
				(_2287_v21).ArraySet1(p2, (_index429).Int())
				var _2307_v39 *C4
				_ = _2307_v39
				var _nw419 *C4 = New_C4_()
				_ = _nw419
				_nw419.Ctor__(p0)
				_2307_v39 = _nw419
				var _2308_v40 _dafny.Sequence
				_ = _2308_v40
				_2308_v40 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_2309_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2310_i1 _dafny.Int) _dafny.CodePoint {
						return _2309_v15
					}
				})(_2281_v15))), (_2289_v23).F3(), _2278_v12, (_2289_v23).F3(), (_2289_v23).F3())
				var _index430 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index430
				var _index431 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _index431
				var _rhs385 _dafny.CodePoint = _2281_v15
				_ = _rhs385
				var _rhs386 bool = (_2307_v39).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("pinmtrsal"), p3, globalState)
				_ = _rhs386
				var _rhs387 bool = !((_2287_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))).Int()).(bool))
				_ = _rhs387
				var _rhs388 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer119 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}((func(_2311_v23 *C1, _2312_p0 _dafny.Int, _2313_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_2314_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update((Companion_D6_.Create_DC11_((_2311_v23).F3())).Dtor_cf21(), (Companion_Default___.SafeIndex(_2312_p0, _dafny.IntOfUint32(((Companion_D6_.Create_DC11_((_2311_v23).F3())).Dtor_cf21()).Cardinality()))).Uint32(), _2313_v15)
					}
				})(_2289_v23, p0, _2281_v15)))
				_ = _rhs388
				var _lhs195 _dafny.Array = _2287_v21
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _lhs196
				var _lhs197 _dafny.Array = _2287_v21
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_2287_v21), 0))
				_ = _lhs198
				_2281_v15 = _rhs385
				(_lhs195).ArraySet1(_rhs386, (_lhs196).Int())
				(_lhs197).ArraySet1(_rhs387, (_lhs198).Int())
				_2308_v40 = _rhs388
				var _2315_v41 _dafny.Sequence
				_ = _2315_v41
				_2315_v41 = _dafny.SeqOf(((_2306_v38).Update((func() bool {
					if (_2285_v19).Contains(p2) {
						return (_2285_v19).Get(p2).(bool)
					}
					return p2
				})(), Companion_Default___.Abs(p0))).Union((_2306_v38).Update(p3, Companion_Default___.Abs(_dafny.IntOfInt64(669)))))
				_2315_v41 = _2315_v41
			}
		}
		var _2316_v42 _dafny.CodePoint
		_ = _2316_v42
		_2316_v42 = _dafny.CodePoint('u')
		var _2317_v43 _dafny.Map
		_ = _2317_v43
		_2317_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm9(p3, !(p2), true, (_this).Fm1((func() bool {
			if (_2317_v43).Contains(p3) {
				return (_2317_v43).Get(p3).(bool)
			}
			return false
		})(), globalState), globalState), (Companion_Default___.SafeIndex((_this).Fm1(p3, globalState), _dafny.IntOfUint32((Companion_Default___.Fm9(p3, !(p2), true, (_this).Fm1((func() bool {
			if (_2317_v43).Contains(p3) {
				return (_2317_v43).Get(p3).(bool)
			}
			return false
		})(), globalState), globalState)).Cardinality()))).Uint32(), _2316_v42), _2316_v42) {
			var _2318_v44 _dafny.Sequence
			_ = _2318_v44
			_2318_v44 = _dafny.UnicodeSeqOfUtf8Bytes("tlysn")
			var _2319_v45 D6
			_ = _2319_v45
			_2319_v45 = Companion_D6_.Create_DC11_(_2318_v44)
			var _2320_v46 D5
			_ = _2320_v46
			_2320_v46 = Companion_D5_.Create_DC9_(p2, p0, Companion_Default___.Fm38(p3, globalState))
			var _2321_v47 _dafny.Sequence
			_ = _2321_v47
			_2321_v47 = _dafny.SeqOf(p3, !_dafny.Companion_Sequence_.Equal(_2318_v44, (_2319_v45).Dtor_cf21()), (_2320_v46).Dtor_cf14())
			if (_2321_v47).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2321_v47).Cardinality()))).Uint32()).(bool) {
				r2 = ((_dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality())).Minus(p0)).Plus(Companion_Default___.SafeModuloInt(p0, p0))
				var _2322_v48 _dafny.Array
				_ = _2322_v48
				var _out50 _dafny.Array
				_ = _out50
				_out50 = Companion_Default___.M0(p2, p3, globalState)
				_2322_v48 = _out50
				var _2323_v49 _dafny.Map
				_ = _2323_v49
				_2323_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
				var _2324_v50 _dafny.Sequence
				_ = _2324_v50
				_2324_v50 = _dafny.SeqOf(p0)
				var _2325_v51 _dafny.Map
				_ = _2325_v51
				_2325_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2323_v49).Contains(p2) {
						return (_2323_v49).Get(p2).(_dafny.Int)
					}
					return p0
				})(), (func() _dafny.Sequence {
					if p3 {
						return _dafny.Companion_Sequence_.Update(_2324_v50, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2324_v50).Cardinality()))).Uint32(), (_this).Fm1(p2, globalState))
					}
					return _2324_v50
				})())
				_2325_v51 = (_2325_v51).Update(p0, _dafny.Companion_Sequence_.Concatenate(_2324_v50, _2324_v50))
				r0 = p2
				r0 = !(p3) || (p2)
			} else {
				r0 = (p3) == (p3)
				var _2326_v52 *C3
				_ = _2326_v52
				var _nw420 *C3 = New_C3_()
				_ = _nw420
				_nw420.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qiesjmhyn"), _2318_v44)).Cardinality()))
				_2326_v52 = _nw420
				r0 = !(p3)
				var _2327_v53 _dafny.Map
				_ = _2327_v53
				_2327_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp(p0) >= 0, p0)
				_2327_v53 = (_2327_v53).Update(!(p2), (func() _dafny.Int {
					if true {
						return p0
					}
					return p0
				})())
				r2 = p0
			}
			var _2328_v54 _dafny.Array
			_ = _2328_v54
			var _nwElement0_98 bool = false
			_ = _nwElement0_98
			var _nw421 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_98, nil, _dafny.IntOfInt64(14))
			_ = _nw421
			(_nw421).ArraySet1(_nwElement0_98, 0)
			(_nw421).ArraySet1(p2, 1)
			(_nw421).ArraySet1(p3, 2)
			(_nw421).ArraySet1(p2, 3)
			(_nw421).ArraySet1(p2, 4)
			(_nw421).ArraySet1(Companion_Default___.Fm3(p2, p0, p0, globalState), 5)
			(_nw421).ArraySet1(false, 6)
			(_nw421).ArraySet1(Companion_Default___.Fm3(p3, p0, _dafny.IntOfInt64(967), globalState), 7)
			(_nw421).ArraySet1(Companion_Default___.Fm3(false, p0, p0, globalState), 8)
			(_nw421).ArraySet1(p2, 9)
			(_nw421).ArraySet1(p3, 10)
			(_nw421).ArraySet1(p3, 11)
			(_nw421).ArraySet1(p3, 12)
			(_nw421).ArraySet1(p2, 13)
			_2328_v54 = _nw421
			var _2329_v55 _dafny.Set
			_ = _2329_v55
			_2329_v55 = _dafny.SetOf(_2328_v54, _2328_v54, _2328_v54, _2328_v54, _2328_v54)
			var _rhs389 bool = !(!(p2))
			_ = _rhs389
			var _rhs390 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).Fm1(p2, globalState), p0)
			_ = _rhs390
			var _rhs391 bool = (_dafny.SetOf(_2328_v54)).IsSubsetOf(_2329_v55)
			_ = _rhs391
			r0 = _rhs389
			r2 = _rhs390
			r0 = _rhs391
			var _index432 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))
			_ = _index432
			(_2328_v54).ArraySet1(p3, (_index432).Int())
			var _2330_v56 _dafny.Map
			_ = _2330_v56
			_2330_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2328_v54, p3)
			var _index433 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))
			_ = _index433
			var _rhs392 bool = p3
			_ = _rhs392
			var _rhs393 _dafny.Map = _2330_v56
			_ = _rhs393
			var _rhs394 _dafny.Int = p0
			_ = _rhs394
			var _lhs199 _dafny.Array = _2328_v54
			_ = _lhs199
			var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))
			_ = _lhs200
			(_lhs199).ArraySet1(_rhs392, (_lhs200).Int())
			_2330_v56 = _rhs393
			r2 = _rhs394
			var _2331_v57 _dafny.MultiSet
			_ = _2331_v57
			_2331_v57 = _dafny.MultiSetOf((_2328_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))).Int()).(bool), !(true), !((func() bool {
				if (_2317_v43).Contains((_2328_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))).Int()).(bool)) {
					return (_2317_v43).Get((_2328_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_2328_v54), 0))).Int()).(bool)).(bool)
				}
				return p2
			})()))
			_2331_v57 = (_dafny.MultiSetOf(p3)).Intersection(_2331_v57)
			var _2332_v58 *C1
			_ = _2332_v58
			var _nw422 *C1 = New_C1_()
			_ = _nw422
			_nw422.Ctor__(_dafny.Companion_Sequence_.Update(_2318_v44, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2318_v44).Cardinality()))).Uint32(), _2316_v42))
			_2332_v58 = _nw422
			_2332_v58 = _2332_v58
		} else {
			r0 = !(p2) || (p2)
			var _2333_v59 _dafny.Array
			_ = _2333_v59
			var _nw423 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
			_ = _nw423
			_2333_v59 = _nw423
			var _2334_v60 _dafny.Sequence
			_ = _2334_v60
			_2334_v60 = _dafny.SeqOf(true)
			var _index434 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2333_v59), 0))
			_ = _index434
			(_2333_v59).ArraySet1CodePoint(Companion_Default___.Fm37((Companion_Default___.Fm49(_2334_v60, p2, p0, globalState)).Cardinality(), globalState), (_index434).Int())
			var _2335_v61 _dafny.Map
			_ = _2335_v61
			_2335_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2316_v42, _dafny.UnicodeSeqOfUtf8Bytes("qnctyt"))
			var _2336_v62 _dafny.Map
			_ = _2336_v62
			_2336_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg120 _dafny.Int) interface{} {
					return coer120(arg120)
				}
			}((func(_2337_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2338_i3 _dafny.Int) _dafny.CodePoint {
					return _2337_v42
				}
			})(_2316_v42)))).Cardinality()), p0)
			var _2339_v63 T2
			_ = _2339_v63
			var _nw424 *C2 = New_C2_()
			_ = _nw424
			_nw424.Ctor__((_dafny.Zero).Minus(p0))
			_2339_v63 = _nw424
			var _2340_v64 _dafny.Set
			_ = _2340_v64
			_2340_v64 = _dafny.SetOf(_2339_v63, _2339_v63, _2339_v63)
			var _2341_v65 _dafny.Map
			_ = _2341_v65
			_2341_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2340_v64)
			var _2342_v66 _dafny.Sequence
			_ = _2342_v66
			_2342_v66 = _dafny.SeqOf(p0, _dafny.IntOfInt64(436), (_this).Fm1(true, globalState), (_2341_v65).Cardinality())
			var _2343_v67 _dafny.Sequence
			_ = _2343_v67
			_2343_v67 = _dafny.SeqOf((_2336_v62).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm15(globalState)).Cardinality()), (_2342_v66).Select((Companion_Default___.SafeIndex(_2339_v63.F2(), _dafny.IntOfUint32((_2342_v66).Cardinality()))).Uint32()).(_dafny.Int))
			var _2344_v68 D11
			_ = _2344_v68
			_2344_v68 = Companion_D11_.Create_DC21_(_2316_v42, _2339_v63.F2(), _2339_v63.F2())
			var _2345_v69 D11
			_ = _2345_v69
			_2345_v69 = Companion_D11_.Create_DC22_(p2)
			var _2346_v70 _dafny.MultiSet
			_ = _2346_v70
			_2346_v70 = _dafny.MultiSetOf(_2345_v69, Companion_D11_.Create_DC22_(p3), _2345_v69)
			var _index435 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2333_v59), 0))
			_ = _index435
			var _rhs395 _dafny.CodePoint = (_2344_v68).Dtor_cf36()
			_ = _rhs395
			var _rhs396 _dafny.Map = _2335_v61
			_ = _rhs396
			var _rhs397 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2342_v66, _2343_v67)
			_ = _rhs397
			var _rhs398 bool = !((_2346_v70).IsSubsetOf(_2346_v70))
			_ = _rhs398
			var _lhs201 _dafny.Array = _2333_v59
			_ = _lhs201
			var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2333_v59), 0))
			_ = _lhs202
			(_lhs201).ArraySet1CodePoint(_rhs395, (_lhs202).Int())
			_2335_v61 = _rhs396
			_2342_v66 = _rhs397
			r0 = _rhs398
			var _2347_v71 _dafny.Sequence
			_ = _2347_v71
			_2347_v71 = _dafny.SeqOf(_2345_v69)
			_2347_v71 = _2347_v71
			r2 = p0
			var _source34 D11 = _2345_v69
			_ = _source34
			if _source34.Is_DC21() {
				var _2348___mcc_h0 _dafny.CodePoint = _source34.Get_().(D11_DC21).Cf36
				_ = _2348___mcc_h0
				var _2349___mcc_h1 _dafny.Int = _source34.Get_().(D11_DC21).Cf37
				_ = _2349___mcc_h1
				var _2350___mcc_h2 _dafny.Int = _source34.Get_().(D11_DC21).Cf38
				_ = _2350___mcc_h2
				var _2351_cf38 _dafny.Int = _2350___mcc_h2
				_ = _2351_cf38
				var _2352_cf37 _dafny.Int = _2349___mcc_h1
				_ = _2352_cf37
				var _2353_cf36 _dafny.CodePoint = _2348___mcc_h0
				_ = _2353_cf36
				var _2354_v72 _dafny.Array
				_ = _2354_v72
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_50
				var _nw425 _dafny.Array
				_ = _nw425
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw425 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) _dafny.Set = (func(_2355_cf37 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_2356_i4 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_dafny.IntOfInt64(-495), _2355_cf37, _2355_cf37)
						}
					})(_2352_cf37)
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw425 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw425).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw425).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_2354_v72 = _nw425
				var _2357_v74 _dafny.Set
				_ = _2357_v74
				_2357_v74 = _dafny.SetOf((_2317_v43).Cardinality())
				var _index436 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_2354_v72), 0))
				_ = _index436
				(_2354_v72).ArraySet1((func() _dafny.Set {
					var _coll72 = _dafny.NewBuilder()
					_ = _coll72
					for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-245), _dafny.IntOfInt64(-526))); ; {
						_compr_72, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _2358_v73 _dafny.Int
						_2358_v73 = interface{}(_compr_72).(_dafny.Int)
						if ((_dafny.IntOfInt64(-245)).Cmp(_2358_v73) <= 0) && ((_2358_v73).Cmp(_dafny.IntOfInt64(-526)) < 0) {
							_coll72.Add((_2358_v73).Minus(_2339_v63.F2()))
						}
					}
					return _coll72.ToSet()
				}()).Union(_2357_v74), (_index436).Int())
				var _index437 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_2354_v72), 0))
				_ = _index437
				(_2354_v72).ArraySet1(_2357_v74, (_index437).Int())
				r0 = p2
				_2352_cf37 = _2351_cf38
				var _2359_v75 _dafny.Sequence
				_ = _2359_v75
				_2359_v75 = _dafny.UnicodeSeqOfUtf8Bytes("ogsgwnlyp")
				var _2360_v76 _dafny.Set
				_ = _2360_v76
				_2360_v76 = _dafny.SetOf(Companion_Default___.Fm5(globalState))
				var _2361_v77 _dafny.Array
				_ = _2361_v77
				var _nwElement0_99 _dafny.Sequence = _2359_v75
				_ = _nwElement0_99
				var _nw426 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_99, nil, _dafny.IntOfInt64(29))
				_ = _nw426
				(_nw426).ArraySet1(_nwElement0_99, 0)
				(_nw426).ArraySet1(_2359_v75, 1)
				(_nw426).ArraySet1(Companion_Default___.Fm28(p2, p2, p2, _2360_v76, globalState), 2)
				(_nw426).ArraySet1(_2359_v75, 3)
				(_nw426).ArraySet1(_2359_v75, 4)
				(_nw426).ArraySet1(_2359_v75, 5)
				(_nw426).ArraySet1(_2359_v75, 6)
				(_nw426).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(258))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg121 _dafny.Int) interface{} {
						return coer121(arg121)
					}
				}((func(_2362_cf36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2363_i5 _dafny.Int) _dafny.CodePoint {
						return _2362_cf36
					}
				})(_2353_cf36))), 7)
				(_nw426).ArraySet1(_2359_v75, 8)
				(_nw426).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("myypb"), 9)
				(_nw426).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _2359_v75
					}
					return _2359_v75
				})(), 10)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2359_v75, _2359_v75), 11)
				(_nw426).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _2359_v75
					}
					return _2359_v75
				})(), 12)
				(_nw426).ArraySet1(_2359_v75, 13)
				(_nw426).ArraySet1(_2359_v75, 14)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2359_v75, _2359_v75), 15)
				(_nw426).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}((func(_2364_cf36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2365_i6 _dafny.Int) _dafny.CodePoint {
						return _2364_cf36
					}
				})(_2353_cf36))), 16)
				(_nw426).ArraySet1(_2359_v75, 17)
				(_nw426).ArraySet1(_2359_v75, 18)
				(_nw426).ArraySet1(_2359_v75, 19)
				(_nw426).ArraySet1(_2359_v75, 20)
				(_nw426).ArraySet1(_2359_v75, 21)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2359_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_2366_cf36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2367_i7 _dafny.Int) _dafny.CodePoint {
						return _2366_cf36
					}
				})(_2353_cf36)))), 22)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2359_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}((func(_2368_v75 _dafny.Sequence, _2369_v43 _dafny.Map) func(_dafny.Int) _dafny.CodePoint {
					return func(_2370_i8 _dafny.Int) _dafny.CodePoint {
						return (_2368_v75).Select((Companion_Default___.SafeIndex((_2369_v43).Cardinality(), _dafny.IntOfUint32((_2368_v75).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_2359_v75, _2317_v43)))), 23)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kcnc"), _dafny.UnicodeSeqOfUtf8Bytes("s")), 24)
				(_nw426).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2359_v75, _2359_v75), 25)
				(_nw426).ArraySet1(_2359_v75, 26)
				(_nw426).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _2359_v75
					}
					return _2359_v75
				})(), 27)
				(_nw426).ArraySet1(_2359_v75, 28)
				_2361_v77 = _nw426
				var _index438 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_2361_v77), 0))
				_ = _index438
				(_2361_v77).ArraySet1(_2359_v75, (_index438).Int())
				var _index439 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_2361_v77), 0))
				_ = _index439
				(_2361_v77).ArraySet1(_2359_v75, (_index439).Int())
			} else if _source34.Is_DC22() {
				var _2371___mcc_h3 bool = _source34.Get_().(D11_DC22).Cf39
				_ = _2371___mcc_h3
				var _2372_cf39 bool = _2371___mcc_h3
				_ = _2372_cf39
				var _2373_v78 _dafny.Map
				_ = _2373_v78
				_2373_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2339_v63.F2(), true)
				var _2374_v79 _dafny.Map
				_ = _2374_v79
				_2374_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("cktdo"), _2373_v78)
				var _2375_v80 D5
				_ = _2375_v80
				_2375_v80 = Companion_D5_.Create_DC10_(p2, p0, _2374_v79, !(p3))
				var _rhs399 D5 = Companion_Default___.Fm29(globalState)
				_ = _rhs399
				var _rhs400 bool = !((true) == (p3))
				_ = _rhs400
				_2375_v80 = _rhs399
				r0 = _rhs400
				var _2376_v81 _dafny.Array
				_ = _2376_v81
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_51
				var _nw427 _dafny.Array
				_ = _nw427
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw427 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Sequence = (func(_2377_v66 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2378_i9 _dafny.Int) _dafny.Sequence {
							return _2377_v66
						}
					})(_2342_v66)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw427 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw427).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw427).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_2376_v81 = _nw427
				_2376_v81 = _2376_v81
				var _2379_v82 _dafny.Sequence
				_ = _2379_v82
				_2379_v82 = _dafny.UnicodeSeqOfUtf8Bytes("odck")
				var _2380_v83 *C1
				_ = _2380_v83
				var _nw428 *C1 = New_C1_()
				_ = _nw428
				_nw428.Ctor__(_2379_v82)
				_2380_v83 = _nw428
				var _rhs401 _dafny.Int = Companion_Default___.SafeModuloInt(p0, _2339_v63.F2())
				_ = _rhs401
				var _rhs402 *C1 = _2380_v83
				_ = _rhs402
				var _rhs403 _dafny.CodePoint = _dafny.CodePoint('j')
				_ = _rhs403
				var _lhs203 T2 = _2339_v63
				_ = _lhs203
				_lhs203.F2_set_(_rhs401)
				_2380_v83 = _rhs402
				_2316_v42 = _rhs403
				var _2381_v84 D6
				_ = _2381_v84
				_2381_v84 = Companion_D6_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-691))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg125 _dafny.Int) interface{} {
						return coer125(arg125)
					}
				}((func(_2382_v59 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_2383_i10 _dafny.Int) _dafny.CodePoint {
						return (_2382_v59).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2382_v59), 0))).Int())
					}
				})(_2333_v59))))
				_2379_v82 = (_2381_v84).Dtor_cf21()
			} else {
				var _2384___mcc_h4 _dafny.Array = _source34.Get_().(D11_DC20).Cf35
				_ = _2384___mcc_h4
				var _2385_cf35 _dafny.Array = _2384___mcc_h4
				_ = _2385_cf35
				r0 = p3
				var _2386_v85 _dafny.Array
				_ = _2386_v85
				var _nw429 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
				_ = _nw429
				_2386_v85 = _nw429
				var _2387_v86 _dafny.Sequence
				_ = _2387_v86
				_2387_v86 = _dafny.UnicodeSeqOfUtf8Bytes("gcqafja")
				var _index440 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2386_v85), 0))
				_ = _index440
				(_2386_v85).ArraySet1(_2387_v86, (_index440).Int())
				var _index441 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2386_v85), 0))
				_ = _index441
				(_2386_v85).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg126 _dafny.Int) interface{} {
								return coer126(arg126)
							}
						}((func(_2388_v59 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
							return func(_2389_i11 _dafny.Int) _dafny.CodePoint {
								return (_2388_v59).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2388_v59), 0))).Int())
							}
						})(_2333_v59)))
					}
					return _2387_v86
				})(), (_index441).Int())
				var _2390_v87 _dafny.Array
				_ = _2390_v87
				var _nwElement0_100 _dafny.Int = _2339_v63.F2()
				_ = _nwElement0_100
				var _nw430 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_100, nil, _dafny.IntOfInt64(14))
				_ = _nw430
				(_nw430).ArraySet1(_nwElement0_100, 0)
				(_nw430).ArraySet1(_2339_v63.F2(), 1)
				(_nw430).ArraySet1(_2339_v63.F2(), 2)
				(_nw430).ArraySet1(_2339_v63.F2(), 3)
				(_nw430).ArraySet1(_2339_v63.F2(), 4)
				(_nw430).ArraySet1(_2339_v63.F2(), 5)
				(_nw430).ArraySet1(p0, 6)
				(_nw430).ArraySet1(_2339_v63.F2(), 7)
				(_nw430).ArraySet1(_dafny.IntOfInt64(-400), 8)
				(_nw430).ArraySet1(_2339_v63.F2(), 9)
				(_nw430).ArraySet1(_2339_v63.F2(), 10)
				(_nw430).ArraySet1(_2339_v63.F2(), 11)
				(_nw430).ArraySet1(_2339_v63.F2(), 12)
				(_nw430).ArraySet1(_dafny.IntOfUint32(((_2386_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2386_v85), 0))).Int()).(_dafny.Sequence)).Cardinality()), 13)
				_2390_v87 = _nw430
				var _2391_v88 _dafny.Map
				_ = _2391_v88
				_2391_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2390_v87)
				var _2392_v89 *C1
				_ = _2392_v89
				var _nw431 *C1 = New_C1_()
				_ = _nw431
				_nw431.Ctor__((_2386_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2386_v85), 0))).Int()).(_dafny.Sequence))
				_2392_v89 = _nw431
				var _2393_v90 _dafny.MultiSet
				_ = _2393_v90
				_2393_v90 = _dafny.MultiSetOf(_2392_v89, _2392_v89)
				var _2394_v91 _dafny.Set
				_ = _2394_v91
				_2394_v91 = _dafny.SetOf((_2391_v88).Cardinality(), p0, (_2393_v90).Cardinality(), (_dafny.Zero).Minus(p0), _dafny.IntOfInt64(-947))
				r0 = !((_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg127 _dafny.Int) interface{} {
						return coer127(arg127)
					}
				}((func(_2395_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2396_i12 _dafny.Int) _dafny.CodePoint {
						return _2395_v42
					}
				})(_2316_v42))), _2387_v86)) == ((_2339_v63.F2()).Cmp((_2394_v91).Cardinality()) == 0))
				(_2339_v63).F2_set_(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(_2339_v63.F2())))
			}
		}
		var _2397_v92 _dafny.Array
		_ = _2397_v92
		var _nw432 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
		_ = _nw432
		_2397_v92 = _nw432
		var _rhs404 _dafny.Int = (func() _dafny.Set {
			var _coll73 = _dafny.NewBuilder()
			_ = _coll73
			for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(182), _dafny.IntOfInt64(266))); ; {
				_compr_73, _ok78 := _iter78()
				if !_ok78 {
					break
				}
				var _2398_v93 _dafny.Int
				_2398_v93 = interface{}(_compr_73).(_dafny.Int)
				if ((_dafny.IntOfInt64(182)).Cmp(_2398_v93) <= 0) && ((_2398_v93).Cmp(_dafny.IntOfInt64(266)) < 0) {
					_coll73.Add((_2398_v93).Minus(p0))
				}
			}
			return _coll73.ToSet()
		}()).Cardinality()
		_ = _rhs404
		var _rhs405 _dafny.Array = _2397_v92
		_ = _rhs405
		r2 = _rhs404
		_2397_v92 = _rhs405
		var _2399_v94 T1
		_ = _2399_v94
		var _nw433 *C6 = New_C6_()
		_ = _nw433
		_nw433.Ctor__()
		_2399_v94 = _nw433
		var _2400_v95 _dafny.Sequence
		_ = _2400_v95
		_2400_v95 = _dafny.SeqOf(_2316_v42)
		var _2401_v96 _dafny.Map
		_ = _2401_v96
		_2401_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2400_v95, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2400_v95).Cardinality()))).Uint32(), _2316_v42)).Cardinality()))
		var _2402_v97 _dafny.Map
		_ = _2402_v97
		_2402_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2399_v94, _2401_v96)
		var _2403_v98 _dafny.Set
		_ = _2403_v98
		_2403_v98 = _dafny.SetOf(p0, p0)
		var _2404_v99 _dafny.Sequence
		_ = _2404_v99
		_2404_v99 = _dafny.SeqOf(true)
		var _2405_v100 _dafny.MultiSet
		_ = _2405_v100
		_2405_v100 = _dafny.MultiSetOf(p2)
		var _2406_v101 _dafny.Array
		_ = _2406_v101
		var _nwElement0_101 _dafny.Int = (Companion_D6_.Create_DC13_(p3, p0)).Dtor_cf25()
		_ = _nwElement0_101
		var _nw434 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_101, nil, _dafny.IntOfInt64(14))
		_ = _nw434
		(_nw434).ArraySet1(_nwElement0_101, 0)
		(_nw434).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if p3 {
				return (_dafny.Zero).Minus(p0)
			}
			return p0
		})())), 1)
		(_nw434).ArraySet1(p0, 2)
		(_nw434).ArraySet1(((func() _dafny.Map {
			if (_2402_v97).Contains(_2399_v94) {
				return (_2402_v97).Get(_2399_v94).(_dafny.Map)
			}
			return _2401_v96
		})()).Cardinality(), 3)
		(_nw434).ArraySet1(((_2403_v98).Intersection(_2403_v98)).Cardinality(), 4)
		(_nw434).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2404_v99).Cardinality()), p0), 5)
		(_nw434).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p0), 6)
		(_nw434).ArraySet1(p0, 7)
		(_nw434).ArraySet1((_2266_v1).Cardinality(), 8)
		(_nw434).ArraySet1((_2405_v100).Cardinality(), 9)
		(_nw434).ArraySet1(p0, 10)
		(_nw434).ArraySet1(p0, 11)
		(_nw434).ArraySet1(p0, 12)
		(_nw434).ArraySet1(p0, 13)
		_2406_v101 = _nw434
		var _2407_v102 _dafny.Sequence
		_ = _2407_v102
		_2407_v102 = _dafny.SeqOf((_this).Fm1(p2, globalState), p0, p0, (_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2404_v99).Cardinality()))
		var _index442 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_2406_v101), 0))
		_ = _index442
		(_2406_v101).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm50(!(p2), !(!(p3)), p2, (_2407_v102).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_2317_v43).Cardinality()), _dafny.IntOfUint32((_2407_v102).Cardinality()))).Uint32()).(_dafny.Int), globalState)).Dtor_cf27()), (_index442).Int())
		var _2408_v103 _dafny.Map
		_ = _2408_v103
		_2408_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _index443 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_2406_v101), 0))
		_ = _index443
		(_2406_v101).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_2408_v103).Contains(p3) {
				return (_2408_v103).Get(p3).(_dafny.Int)
			}
			return p0
		})(), _dafny.IntOfUint32((_2407_v102).Cardinality())))), (_index443).Int())
		var _2409_v104 _dafny.Array
		_ = _2409_v104
		var _nw435 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw435
		_2409_v104 = _nw435
		var _index444 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_2409_v104), 0))
		_ = _index444
		(_2409_v104).ArraySet1(p3, (_index444).Int())
		var _2410_v105 _dafny.Map
		_ = _2410_v105
		_2410_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2406_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_2406_v101), 0))).Int()).(_dafny.Int), _2405_v100)
		var _index445 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_2409_v104), 0))
		_ = _index445
		(_2409_v104).ArraySet1((_2405_v100).IsProperSubsetOf((func() _dafny.MultiSet {
			if (_2410_v105).Contains(p0) {
				return (_2410_v105).Get(p0).(_dafny.MultiSet)
			}
			return _2405_v100
		})()), (_index445).Int())
		var _index446 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_2406_v101), 0))
		_ = _index446
		(_2406_v101).ArraySet1(_dafny.IntOfUint32((_2404_v99).Cardinality()), (_index446).Int())
		r0 = p3
		var _2411_v106 _dafny.Array
		_ = _2411_v106
		var _nw436 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
		_ = _nw436
		_2411_v106 = _nw436
		r1 = _2411_v106
		r2 = (p0).Times(p0)
		return r0, r1, r2
	}
}
func (_this *C7) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _2412_v0 bool
		_ = _2412_v0
		_2412_v0 = true
		_2412_v0 = p2
		if false {
			var _2413_v1 _dafny.Map
			_ = _2413_v1
			_2413_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			_2413_v1 = (_2413_v1).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-409), p0), _2412_v0)
			var _2414_v2 T2
			_ = _2414_v2
			var _nw437 *C3 = New_C3_()
			_ = _nw437
			_nw437.Ctor__(((_dafny.Zero).Minus(p1)).Plus(p0))
			_2414_v2 = _nw437
			_2414_v2 = _2414_v2
			var _2415_v3 _dafny.MultiSet
			_ = _2415_v3
			_2415_v3 = _dafny.MultiSetOf(!(true), p2, _2412_v0, !(!(_2412_v0)))
			var _2416_v4 _dafny.Sequence
			_ = _2416_v4
			_2416_v4 = _dafny.SeqOf((_2415_v3).Update(true, Companion_Default___.Abs(_dafny.IntOfInt64(525))))
			_2416_v4 = _dafny.SeqOf((_2415_v3).Union((_2416_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2416_v4).Cardinality()))).Uint32()).(_dafny.MultiSet)))
			var _2417_v5 _dafny.CodePoint
			_ = _2417_v5
			_2417_v5 = _dafny.CodePoint('r')
			var _2418_v6 _dafny.Sequence
			_ = _2418_v6
			_2418_v6 = _dafny.SeqOf(_dafny.CodePoint('c'), _2417_v5)
			var _2419_v7 _dafny.Map
			_ = _2419_v7
			_2419_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2418_v6, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_2417_v5, _dafny.CodePoint('k')), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(_2417_v5, _dafny.CodePoint('k'))).Cardinality()))).Uint32(), _2417_v5)), _2412_v0)
			if (func() bool {
				if (_2419_v7).Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer128 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg128 _dafny.Int) interface{} {
						return coer128(arg128)
					}
				}((func(_2420_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2421_i0 _dafny.Int) _dafny.CodePoint {
						return _2420_v5
					}
				})(_2417_v5))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2418_v6).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg129 _dafny.Int) interface{} {
						return coer129(arg129)
					}
				}((func(_2422_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2423_i0 _dafny.Int) _dafny.CodePoint {
						return _2422_v5
					}
				})(_2417_v5)))).Cardinality()))).Uint32(), _2417_v5)) {
					return (_2419_v7).Get(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg130 _dafny.Int) interface{} {
							return coer130(arg130)
						}
					}((func(_2424_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2425_i0 _dafny.Int) _dafny.CodePoint {
							return _2424_v5
						}
					})(_2417_v5))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2418_v6).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg131 _dafny.Int) interface{} {
							return coer131(arg131)
						}
					}((func(_2426_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2427_i0 _dafny.Int) _dafny.CodePoint {
							return _2426_v5
						}
					})(_2417_v5)))).Cardinality()))).Uint32(), _2417_v5)).(bool)
				}
				return !(p2)
			})() {
				var _2428_v8 _dafny.Set
				_ = _2428_v8
				_2428_v8 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(227))).Uint32(), func(coer132 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg132 _dafny.Int) interface{} {
						return coer132(arg132)
					}
				}((func(_2429_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2430_i1 _dafny.Int) _dafny.CodePoint {
						return _2429_v5
					}
				})(_2417_v5))), _2418_v6, _2418_v6, _2418_v6)
				var _2431_v9 _dafny.Array
				_ = _2431_v9
				var _nwElement0_102 _dafny.Int = _2414_v2.F2()
				_ = _nwElement0_102
				var _nw438 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_102, nil, _dafny.IntOfInt64(5))
				_ = _nw438
				(_nw438).ArraySet1(_nwElement0_102, 0)
				(_nw438).ArraySet1(_dafny.IntOfInt64(-808), 1)
				(_nw438).ArraySet1(_2414_v2.F2(), 2)
				(_nw438).ArraySet1((_2428_v8).Cardinality(), 3)
				(_nw438).ArraySet1(p0, 4)
				_2431_v9 = _nw438
				var _index447 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_2431_v9), 0))
				_ = _index447
				(_2431_v9).ArraySet1(p0, (_index447).Int())
				var _index448 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_2431_v9), 0))
				_ = _index448
				(_2431_v9).ArraySet1((_2414_v2.F2()).Plus(p1), (_index448).Int())
				var _2432_v10 _dafny.Sequence
				_ = _2432_v10
				_2432_v10 = _dafny.SeqOf(p1)
				var _2433_v11 _dafny.Map
				_ = _2433_v11
				_2433_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2432_v10, _2414_v2.F2())
				_2433_v11 = _2433_v11
				var _2434_v12 _dafny.Set
				_ = _2434_v12
				_2434_v12 = _dafny.SetOf(_2412_v0, p2)
				var _rhs406 bool = !(Companion_Default___.Fm3(p2, (func() _dafny.Int {
					if p2 {
						return _dafny.IntOfUint32((_dafny.SeqOf(false, p2)).Cardinality())
					}
					return (_2431_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_2431_v9), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((Companion_Default___.Fm15(globalState)).Cardinality()), globalState))
				_ = _rhs406
				var _rhs407 _dafny.Int = ((_2434_v12).Cardinality()).Plus((_dafny.Zero).Minus(_2414_v2.F2()))
				_ = _rhs407
				var _rhs408 bool = _2412_v0
				_ = _rhs408
				var _rhs409 _dafny.Sequence = _2418_v6
				_ = _rhs409
				var _lhs204 T2 = _2414_v2
				_ = _lhs204
				_2412_v0 = _rhs406
				_lhs204.F2_set_(_rhs407)
				_2412_v0 = _rhs408
				_2418_v6 = _rhs409
				_2417_v5 = _2417_v5
				_2412_v0 = _2412_v0
			} else {
				_2412_v0 = (_2414_v2.F2()).Cmp((_2413_v1).Cardinality()) <= 0
				var _2435_v13 *C4
				_ = _2435_v13
				var _nw439 *C4 = New_C4_()
				_ = _nw439
				_nw439.Ctor__(Companion_Default___.SafeModuloInt(p1, (_2415_v3).Cardinality()))
				_2435_v13 = _nw439
				var _2436_v14 _dafny.Sequence
				_ = _2436_v14
				_2436_v14 = _dafny.SeqOf(p2, p2)
				var _2437_v15 _dafny.Map
				_ = _2437_v15
				_2437_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2436_v14, p1)
				_2437_v15 = (_2437_v15).Update(_2436_v14, (_dafny.Zero).Minus((_2435_v13).Fm1(_2412_v0, globalState)))
				var _2438_v16 _dafny.Array
				_ = _2438_v16
				var _nw440 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw440
				_2438_v16 = _nw440
				var _len0_52 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_52
				var _nw441 _dafny.Array
				_ = _nw441
				if _len0_52.Cmp(_dafny.Zero) == 0 {
					_nw441 = _dafny.NewArray(_len0_52)
				} else {
					var _init52 func(_dafny.Int) _dafny.Sequence = (func(_2439_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2440_i2 _dafny.Int) _dafny.Sequence {
							return _2439_v6
						}
					})(_2418_v6)
					_ = _init52
					var _element0_52 = _init52(_dafny.Zero)
					_ = _element0_52
					_nw441 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
					(_nw441).ArraySet1(_element0_52, 0)
					var _nativeLen0_52 = (_len0_52).Int()
					_ = _nativeLen0_52
					for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
						(_nw441).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
					}
				}
				_2438_v16 = _nw441
				var _2441_v17 _dafny.Sequence
				_ = _2441_v17
				_2441_v17 = _dafny.SeqOf(_2414_v2.F2(), _2414_v2.F2())
				var _2442_v18 _dafny.MultiSet
				_ = _2442_v18
				_2442_v18 = _dafny.MultiSetOf(_2441_v17)
				var _2443_v19 _dafny.Map
				_ = _2443_v19
				_2443_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_2412_v0)).IsProperSubsetOf(_2415_v3), _2442_v18)
				_2443_v19 = (_2443_v19).Update((_dafny.IntOfInt64(-163)).Cmp(_2414_v2.F2()) > 0, _dafny.MultiSetOf(_2441_v17, _2441_v17, _2441_v17))
			}
			if !((p2) == ((_2414_v2).Fm8(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("fvqbna"), (Companion_Default___.SafeIndex(_2414_v2.F2(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvqbna")).Cardinality()))).Uint32(), _dafny.CodePoint('k')), _2412_v0, globalState))) {
				var _2444_v20 _dafny.MultiSet
				_ = _2444_v20
				_2444_v20 = _dafny.MultiSetOf(_2414_v2.F2())
				var _2445_v21 _dafny.Sequence
				_ = _2445_v21
				_2445_v21 = _dafny.SeqOf(_2418_v6)
				var _2446_v22 _dafny.Sequence
				_ = _2446_v22
				_2446_v22 = _dafny.SeqOf(_2412_v0, p2)
				var _2447_v23 D6
				_ = _2447_v23
				_2447_v23 = Companion_D6_.Create_DC12_((_2446_v22).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2446_v22).Cardinality()))).Uint32()).(bool), _2412_v0)
				var _2448_v24 _dafny.Set
				_ = _2448_v24
				_2448_v24 = _dafny.SetOf(_2414_v2.F2(), _2414_v2.F2())
				var _2449_v25 D11
				_ = _2449_v25
				_2449_v25 = Companion_D11_.Create_DC21_(_2417_v5, _2414_v2.F2(), p0)
				var _2450_v26 _dafny.Array
				_ = _2450_v26
				var _nwElement0_103 bool = _2412_v0
				_ = _nwElement0_103
				var _nw442 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_103, nil, _dafny.IntOfInt64(10))
				_ = _nw442
				(_nw442).ArraySet1(_nwElement0_103, 0)
				(_nw442).ArraySet1(p2, 1)
				(_nw442).ArraySet1(_2412_v0, 2)
				(_nw442).ArraySet1(!(p2), 3)
				(_nw442).ArraySet1(_2412_v0, 4)
				(_nw442).ArraySet1(_2412_v0, 5)
				(_nw442).ArraySet1(_2412_v0, 6)
				(_nw442).ArraySet1(true, 7)
				(_nw442).ArraySet1(p2, 8)
				(_nw442).ArraySet1(p2, 9)
				_2450_v26 = _nw442
				var _2451_v27 _dafny.Map
				_ = _2451_v27
				_2451_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2450_v26, _2414_v2.F2())
				var _2452_v28 *C1
				_ = _2452_v28
				var _nw443 *C1 = New_C1_()
				_ = _nw443
				_nw443.Ctor__(_2418_v6)
				_2452_v28 = _nw443
				var _2453_v29 D8
				_ = _2453_v29
				_2453_v29 = Companion_D8_.Create_DC17_(_2412_v0, _2452_v28, p1)
				var _2454_v30 _dafny.Sequence
				_ = _2454_v30
				_2454_v30 = _dafny.SeqOf(p0)
				var _2455_v31 _dafny.Array
				_ = _2455_v31
				var _nwElement0_104 _dafny.Int = (_2414_v2.F2()).Plus(p1)
				_ = _nwElement0_104
				var _nw444 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_104, nil, _dafny.IntOfInt64(26))
				_ = _nw444
				(_nw444).ArraySet1(_nwElement0_104, 0)
				(_nw444).ArraySet1((_2444_v20).Cardinality(), 1)
				(_nw444).ArraySet1((_2414_v2.F2()).Plus(_2414_v2.F2()), 2)
				(_nw444).ArraySet1(_dafny.IntOfInt64(-705), 3)
				(_nw444).ArraySet1(_2414_v2.F2(), 4)
				(_nw444).ArraySet1(_2414_v2.F2(), 5)
				(_nw444).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_2445_v21).Select((Companion_Default___.SafeIndex((_this).Fm1(p2, globalState), _dafny.IntOfUint32((_2445_v21).Cardinality()))).Uint32()).(_dafny.Sequence), _2418_v6)).Cardinality()), 6)
				(_nw444).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if p2 {
						return _2418_v6
					}
					return _2418_v6
				})()).Cardinality()), 7)
				(_nw444).ArraySet1(p1, 8)
				(_nw444).ArraySet1((func() _dafny.Int {
					if p2 {
						return _2414_v2.F2()
					}
					return p0
				})(), 9)
				(_nw444).ArraySet1(p0, 10)
				(_nw444).ArraySet1(_2414_v2.F2(), 11)
				(_nw444).ArraySet1(_2414_v2.F2(), 12)
				(_nw444).ArraySet1((func() _dafny.Int {
					if (_2447_v23).Dtor_cf23() {
						return _2414_v2.F2()
					}
					return _2414_v2.F2()
				})(), 13)
				(_nw444).ArraySet1((_dafny.IntOfUint32((_2446_v22).Cardinality())).Minus(_2414_v2.F2()), 14)
				(_nw444).ArraySet1(_2414_v2.F2(), 15)
				(_nw444).ArraySet1((_2414_v2.F2()).Times(_2414_v2.F2()), 16)
				(_nw444).ArraySet1(_2414_v2.F2(), 17)
				(_nw444).ArraySet1((p0).Times((_2448_v24).Cardinality()), 18)
				(_nw444).ArraySet1((_2449_v25).Dtor_cf37(), 19)
				(_nw444).ArraySet1(Companion_Default___.SafeDivisionInt(_2414_v2.F2(), p0), 20)
				(_nw444).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_2451_v27).Cardinality()), 21)
				(_nw444).ArraySet1((_2453_v29).Dtor_cf32(), 22)
				(_nw444).ArraySet1(Companion_Default___.SafeModuloInt(_2414_v2.F2(), p0), 23)
				(_nw444).ArraySet1(p0, 24)
				(_nw444).ArraySet1(_dafny.IntOfUint32((_2454_v30).Cardinality()), 25)
				_2455_v31 = _nw444
				var _index449 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_2455_v31), 0))
				_ = _index449
				(_2455_v31).ArraySet1(p0, (_index449).Int())
				var _2456_v32 _dafny.Set
				_ = _2456_v32
				_2456_v32 = _dafny.SetOf(_2450_v26)
				var _index450 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_2455_v31), 0))
				_ = _index450
				(_2455_v31).ArraySet1(((_2456_v32).Cardinality()).Times((_dafny.Zero).Minus(_2414_v2.F2())), (_index450).Int())
				var _2457_v33 *C2
				_ = _2457_v33
				var _nw445 *C2 = New_C2_()
				_ = _nw445
				_nw445.Ctor__(_2414_v2.F2())
				_2457_v33 = _nw445
				(_2414_v2).F2_set_(p1)
				var _2458_v34 _dafny.Array
				_ = _2458_v34
				var _nwElement0_105 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer133 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg133 _dafny.Int) interface{} {
						return coer133(arg133)
					}
				}((func(_2459_v2 T2) func(_dafny.Int) _dafny.Int {
					return func(_2460_i3 _dafny.Int) _dafny.Int {
						return _2459_v2.F2()
					}
				})(_2414_v2)))
				_ = _nwElement0_105
				var _nw446 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_105, nil, _dafny.IntOfInt64(25))
				_ = _nw446
				(_nw446).ArraySet1(_nwElement0_105, 0)
				(_nw446).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer134 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg134 _dafny.Int) interface{} {
						return coer134(arg134)
					}
				}((func(_2461_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2462_i4 _dafny.Int) _dafny.Int {
						return _2461_p1
					}
				})(p1))), 1)
				(_nw446).ArraySet1(_dafny.SeqOf(_2414_v2.F2()), 2)
				(_nw446).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-633))).Uint32(), func(coer135 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg135 _dafny.Int) interface{} {
						return coer135(arg135)
					}
				}((func(_2463_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2464_i5 _dafny.Int) _dafny.Int {
						return _2463_p0
					}
				})(p0))), (Companion_Default___.SafeIndex(_2414_v2.F2(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-633))).Uint32(), func(coer136 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg136 _dafny.Int) interface{} {
						return coer136(arg136)
					}
				}((func(_2465_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2466_i5 _dafny.Int) _dafny.Int {
						return _2465_p0
					}
				})(p0)))).Cardinality()))).Uint32(), (_2413_v1).Cardinality()), 3)
				(_nw446).ArraySet1(_2454_v30, 4)
				(_nw446).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32(((_2452_v28).F3()).Cardinality()), p0), 5)
				(_nw446).ArraySet1(_2454_v30, 6)
				(_nw446).ArraySet1(_2454_v30, 7)
				(_nw446).ArraySet1(_2454_v30, 8)
				(_nw446).ArraySet1(_2454_v30, 9)
				(_nw446).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_2454_v30).Cardinality()), _dafny.IntOfUint32((_2454_v30).Cardinality())), 10)
				(_nw446).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer137 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg137 _dafny.Int) interface{} {
						return coer137(arg137)
					}
				}((func(_2467_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2468_i6 _dafny.Int) _dafny.Int {
						return _2467_p1
					}
				})(p1))), 11)
				(_nw446).ArraySet1(_dafny.SeqOf((_2415_v3).Cardinality()), 12)
				(_nw446).ArraySet1(_2454_v30, 13)
				(_nw446).ArraySet1(_2454_v30, 14)
				(_nw446).ArraySet1(_2454_v30, 15)
				(_nw446).ArraySet1(_2454_v30, 16)
				(_nw446).ArraySet1(_2454_v30, 17)
				(_nw446).ArraySet1(_dafny.SeqOf(_2414_v2.F2(), _2414_v2.F2()), 18)
				(_nw446).ArraySet1(_dafny.SeqOf(_2414_v2.F2()), 19)
				(_nw446).ArraySet1(_dafny.SeqOf(_2414_v2.F2()), 20)
				(_nw446).ArraySet1(_dafny.SeqOf((_2455_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_2455_v31), 0))).Int()).(_dafny.Int), (_2457_v33).Fm1(p2, globalState)), 21)
				(_nw446).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(54))).Uint32(), func(coer138 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg138 _dafny.Int) interface{} {
						return coer138(arg138)
					}
				}((func(_2469_v2 T2) func(_dafny.Int) _dafny.Int {
					return func(_2470_i7 _dafny.Int) _dafny.Int {
						return _2469_v2.F2()
					}
				})(_2414_v2))), 22)
				(_nw446).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(383))).Uint32(), func(coer139 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg139 _dafny.Int) interface{} {
						return coer139(arg139)
					}
				}((func(_2471_v2 T2) func(_dafny.Int) _dafny.Int {
					return func(_2472_i8 _dafny.Int) _dafny.Int {
						return _2471_v2.F2()
					}
				})(_2414_v2))), 23)
				(_nw446).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(842))).Uint32(), func(coer140 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg140 _dafny.Int) interface{} {
						return coer140(arg140)
					}
				}((func(_2473_v2 T2) func(_dafny.Int) _dafny.Int {
					return func(_2474_i9 _dafny.Int) _dafny.Int {
						return _2473_v2.F2()
					}
				})(_2414_v2))), 24)
				_2458_v34 = _nw446
				var _2475_v35 bool
				_ = _2475_v35
				var _2476_v36 _dafny.Array
				_ = _2476_v36
				var _2477_v37 _dafny.Int
				_ = _2477_v37
				var _out51 bool
				_ = _out51
				var _out52 _dafny.Array
				_ = _out52
				var _out53 _dafny.Int
				_ = _out53
				_out51, _out52, _out53 = (_2452_v28).M1((_dafny.IntOfUint32((_2454_v30).Cardinality())).Plus(p1), _2458_v34, p2, p2, globalState)
				_2475_v35 = _out51
				_2476_v36 = _out52
				_2477_v37 = _out53
				var _2478_v38 _dafny.Map
				_ = _2478_v38
				_2478_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.SafeDivisionInt((_2414_v2).Fm1((_2452_v28).Fm11(globalState), globalState), _2414_v2.F2()))
				_2478_v38 = (_2478_v38).Update(_2412_v0, _2477_v37)
			} else {
				_2412_v0 = !(p2)
				_2412_v0 = _2412_v0
				r0 = _2414_v2.F2()
				var _2479_v39 _dafny.Sequence
				_ = _2479_v39
				_2479_v39 = _dafny.SeqOf(p0, _2414_v2.F2(), _2414_v2.F2(), p0)
				var _2480_v40 _dafny.Map
				_ = _2480_v40
				_2480_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2417_v5, _dafny.IntOfUint32((_2479_v39).Cardinality()))
				_2480_v40 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2417_v5, p1)).Merge(_2480_v40)).Update(_2417_v5, p0)
				var _2481_v41 _dafny.Array
				_ = _2481_v41
				var _nw447 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw447
				_2481_v41 = _nw447
				var _2482_v43 _dafny.Set
				_ = _2482_v43
				_2482_v43 = _dafny.SetOf(_2414_v2.F2(), (_2414_v2).Fm1(true, globalState), p1)
				var _rhs410 _dafny.Int = ((func() _dafny.Set {
					var _coll74 = _dafny.NewBuilder()
					_ = _coll74
					for _iter79 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(329), _dafny.IntOfInt64(686))); ; {
						_compr_74, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _2483_v42 _dafny.Int
						_2483_v42 = interface{}(_compr_74).(_dafny.Int)
						if ((_dafny.IntOfInt64(329)).Cmp(_2483_v42) <= 0) && ((_2483_v42).Cmp(_dafny.IntOfInt64(686)) < 0) {
							_coll74.Add((_2483_v42).Minus(_dafny.IntOfInt64(270)))
						}
					}
					return _coll74.ToSet()
				}()).Union(_2482_v43)).Cardinality()
				_ = _rhs410
				var _rhs411 _dafny.Array = _2481_v41
				_ = _rhs411
				var _rhs412 _dafny.Int = p0
				_ = _rhs412
				r0 = _rhs410
				_2481_v41 = _rhs411
				r0 = _rhs412
			}
		} else {
			var _2484_v44 _dafny.Sequence
			_ = _2484_v44
			_2484_v44 = _dafny.UnicodeSeqOfUtf8Bytes("jy")
			var _2485_v45 _dafny.Map
			_ = _2485_v45
			_2485_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2412_v0, p0)
			var _2486_v47 _dafny.Sequence
			_ = _2486_v47
			_2486_v47 = _dafny.SeqOf(_dafny.IntOfUint32((_2484_v44).Cardinality()))
			var _2487_v48 _dafny.Sequence
			_ = _2487_v48
			_2487_v48 = _dafny.SeqOf(_2412_v0, true)
			var _2488_v49 _dafny.Map
			_ = _2488_v49
			_2488_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_2486_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2487_v48).Cardinality()), _dafny.IntOfUint32((_2486_v47).Cardinality()))).Uint32()).(_dafny.Int))
			var _2489_v50 _dafny.Set
			_ = _2489_v50
			_2489_v50 = _dafny.SetOf(p0, (_2488_v49).Cardinality())
			var _2490_v51 _dafny.Sequence
			_ = _2490_v51
			_2490_v51 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_2484_v44).Cardinality()), p0, (func() _dafny.Int {
				if (_2485_v45).Contains(_2412_v0) {
					return (_2485_v45).Get(_2412_v0).(_dafny.Int)
				}
				return p1
			})()), func() _dafny.Set {
				var _coll75 = _dafny.NewBuilder()
				_ = _coll75
				for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(345), _dafny.IntOfInt64(363))); ; {
					_compr_75, _ok80 := _iter80()
					if !_ok80 {
						break
					}
					var _2491_v46 _dafny.Int
					_2491_v46 = interface{}(_compr_75).(_dafny.Int)
					if ((_dafny.IntOfInt64(345)).Cmp(_2491_v46) <= 0) && ((_2491_v46).Cmp(_dafny.IntOfInt64(363)) < 0) {
						_coll75.Add((_2491_v46).Plus(p1))
					}
				}
				return _coll75.ToSet()
			}(), _2489_v50)
			var _2492_v52 D7
			_ = _2492_v52
			_2492_v52 = Companion_D7_.Create_DC14_(_2486_v47)
			var _2493_v53 _dafny.Sequence
			_ = _2493_v53
			_2493_v53 = _dafny.SeqOf(_2492_v52, _2492_v52)
			var _2494_v54 _dafny.Array
			_ = _2494_v54
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_53
			var _nw448 _dafny.Array
			_ = _nw448
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw448 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Int = (func(_2495_p2 bool, _2496_p1 _dafny.Int, _2497_v44 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_2498_i10 _dafny.Int) _dafny.Int {
						return (_2498_i10).Plus((Companion_D5_.Create_DC10_(_2495_p2, _2496_p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2497_v44, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2496_p1, _2495_p2)), _2495_p2)).Dtor_cf18())
					}
				})(p2, p1, _2484_v44)
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw448 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw448).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw448).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_2494_v54 = _nw448
			var _2499_v55 _dafny.Sequence
			_ = _2499_v55
			_2499_v55 = _dafny.SeqOf(_2494_v54)
			var _2500_v56 _dafny.MultiSet
			_ = _2500_v56
			_2500_v56 = _dafny.MultiSetOf(_2412_v0, false, p2, _2412_v0)
			var _2501_v57 _dafny.Array
			_ = _2501_v57
			var _nwElement0_106 bool = (p0).Cmp(_dafny.IntOfUint32((_2490_v51).Cardinality())) > 0
			_ = _nwElement0_106
			var _nw449 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_106, nil, _dafny.IntOfInt64(27))
			_ = _nw449
			(_nw449).ArraySet1(_nwElement0_106, 0)
			(_nw449).ArraySet1((_dafny.IntOfUint32((_2493_v53).Cardinality())).Cmp((_this).Fm1(_2412_v0, globalState)) > 0, 1)
			(_nw449).ArraySet1(_2412_v0, 2)
			(_nw449).ArraySet1((_2412_v0) == (p2), 3)
			(_nw449).ArraySet1((true) && (!(false)), 4)
			(_nw449).ArraySet1(_2412_v0, 5)
			(_nw449).ArraySet1(p2, 6)
			(_nw449).ArraySet1((_dafny.IntOfInt64(-180)).Cmp(_dafny.IntOfUint32((_2499_v55).Cardinality())) == 0, 7)
			(_nw449).ArraySet1(_2412_v0, 8)
			(_nw449).ArraySet1(!(p2), 9)
			(_nw449).ArraySet1(_2412_v0, 10)
			(_nw449).ArraySet1(_2412_v0, 11)
			(_nw449).ArraySet1((_2412_v0) && (false), 12)
			(_nw449).ArraySet1(p2, 13)
			(_nw449).ArraySet1((_2500_v56).IsSubsetOf(_2500_v56), 14)
			(_nw449).ArraySet1(p2, 15)
			(_nw449).ArraySet1(_2412_v0, 16)
			(_nw449).ArraySet1(Companion_Default___.Fm3(_2412_v0, p1, p1, globalState), 17)
			(_nw449).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_this).Fm1(_2412_v0, globalState), (_2486_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2484_v44).Cardinality()), _dafny.IntOfUint32((_2486_v47).Cardinality()))).Uint32()).(_dafny.Int)), p0), 18)
			(_nw449).ArraySet1((p0).Cmp(p0) >= 0, 19)
			(_nw449).ArraySet1(_2412_v0, 20)
			(_nw449).ArraySet1(false, 21)
			(_nw449).ArraySet1(!(p2) || (_2412_v0), 22)
			(_nw449).ArraySet1(_2412_v0, 23)
			(_nw449).ArraySet1(p2, 24)
			(_nw449).ArraySet1(_2412_v0, 25)
			(_nw449).ArraySet1(_dafny.Companion_Sequence_.Equal(_2484_v44, _dafny.UnicodeSeqOfUtf8Bytes("krmp")), 26)
			_2501_v57 = _nw449
			var _index451 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
			_ = _index451
			(_2501_v57).ArraySet1(Companion_Default___.Fm5(globalState), (_index451).Int())
			var _index452 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
			_ = _index452
			var _rhs413 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2484_v44, _2484_v44), _2484_v44)
			_ = _rhs413
			var _rhs414 _dafny.Int = p0
			_ = _rhs414
			var _rhs415 bool = _2412_v0
			_ = _rhs415
			var _lhs205 _dafny.Array = _2501_v57
			_ = _lhs205
			var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
			_ = _lhs206
			_2484_v44 = _rhs413
			r0 = _rhs414
			(_lhs205).ArraySet1(_rhs415, (_lhs206).Int())
			if !((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool)) {
				var _2502_v58 _dafny.Array
				_ = _2502_v58
				var _nw450 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw450
				_2502_v58 = _nw450
				var _index453 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2502_v58), 0))
				_ = _index453
				(_2502_v58).ArraySet1(_2494_v54, (_index453).Int())
				var _index454 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2502_v58), 0))
				_ = _index454
				(_2502_v58).ArraySet1(_2494_v54, (_index454).Int())
				var _2503_v59 _dafny.Map
				_ = _2503_v59
				_2503_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), _dafny.CodePoint('r'))
				var _2504_v60 _dafny.Set
				_ = _2504_v60
				_2504_v60 = _dafny.SetOf(_2503_v59)
				var _index455 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index455
				var _rhs416 _dafny.Int = p0
				_ = _rhs416
				var _rhs417 _dafny.Set = (_2504_v60).Union((_dafny.SetOf(_2503_v59)).Union(_2504_v60))
				_ = _rhs417
				var _rhs418 bool = _2412_v0
				_ = _rhs418
				var _rhs419 bool = (p1).Cmp((_dafny.SetOf((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool))).Cardinality()) == 0
				_ = _rhs419
				var _lhs207 _dafny.Array = _2501_v57
				_ = _lhs207
				var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _lhs208
				r0 = _rhs416
				_2504_v60 = _rhs417
				_2412_v0 = _rhs418
				(_lhs207).ArraySet1(_rhs419, (_lhs208).Int())
				var _2505_v61 _dafny.CodePoint
				_ = _2505_v61
				_2505_v61 = _dafny.CodePoint('b')
				var _2506_v62 D5
				_ = _2506_v62
				_2506_v62 = Companion_D5_.Create_DC9_((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), _dafny.IntOfInt64(69), _2500_v56)
				var _index456 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index456
				var _rhs420 bool = (p2) && ((_2487_v48).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2487_v48).Cardinality()))).Uint32()).(bool))
				_ = _rhs420
				var _rhs421 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool) {
						return (_2506_v62).Dtor_cf15()
					}
					return p0
				})())
				_ = _rhs421
				var _rhs422 _dafny.CodePoint = _2505_v61
				_ = _rhs422
				var _lhs209 _dafny.Array = _2501_v57
				_ = _lhs209
				var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _lhs210
				(_lhs209).ArraySet1(_rhs420, (_lhs210).Int())
				r0 = _rhs421
				_2505_v61 = _rhs422
				var _2507_v63 T0
				_ = _2507_v63
				var _nw451 *C6 = New_C6_()
				_ = _nw451
				_nw451.Ctor__()
				_2507_v63 = _nw451
				var _index457 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index457
				var _rhs423 T0 = _2507_v63
				_ = _rhs423
				var _rhs424 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ongtjwil"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(959))).Uint32(), func(coer141 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg141 _dafny.Int) interface{} {
						return coer141(arg141)
					}
				}((func(_2508_v61 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2509_i11 _dafny.Int) _dafny.CodePoint {
						return _2508_v61
					}
				})(_2505_v61))))
				_ = _rhs424
				var _rhs425 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2487_v48, _2487_v48), _dafny.SeqOf(_2412_v0, false))
				_ = _rhs425
				var _rhs426 bool = (_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool)
				_ = _rhs426
				var _lhs211 _dafny.Array = _2501_v57
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _lhs212
				_2507_v63 = _rhs423
				_2484_v44 = _rhs424
				_2487_v48 = _rhs425
				(_lhs211).ArraySet1(_rhs426, (_lhs212).Int())
				var _2510_v64 _dafny.Array
				_ = _2510_v64
				var _nw452 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
				_ = _nw452
				_2510_v64 = _nw452
				var _index458 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_2510_v64), 0))
				_ = _index458
				(_2510_v64).ArraySet1(_2489_v50, (_index458).Int())
				var _index459 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index459
				var _index460 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_2510_v64), 0))
				_ = _index460
				var _rhs427 bool = !((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool))
				_ = _rhs427
				var _rhs428 bool = (p1).Cmp((func() _dafny.Int {
					if _2412_v0 {
						return p1
					}
					return (_2503_v59).Cardinality()
				})()) <= 0
				_ = _rhs428
				var _rhs429 _dafny.Array = _2501_v57
				_ = _rhs429
				var _rhs430 _dafny.Set = _dafny.SetOf(_dafny.IntOfInt64(353), p1, (p0).Plus(_dafny.IntOfInt64(573)))
				_ = _rhs430
				var _rhs431 bool = _2412_v0
				_ = _rhs431
				var _lhs213 _dafny.Array = _2501_v57
				_ = _lhs213
				var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _lhs214
				var _lhs215 _dafny.Array = _2510_v64
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_2510_v64), 0))
				_ = _lhs216
				_2412_v0 = _rhs427
				(_lhs213).ArraySet1(_rhs428, (_lhs214).Int())
				_2501_v57 = _rhs429
				(_lhs215).ArraySet1(_rhs430, (_lhs216).Int())
				_2412_v0 = _rhs431
			} else {
				var _2511_v65 D6
				_ = _2511_v65
				_2511_v65 = Companion_D6_.Create_DC13_((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), (_this).Fm1(_2412_v0, globalState))
				_2412_v0 = (_2511_v65).Dtor_cf24()
				var _2512_v66 _dafny.CodePoint
				_ = _2512_v66
				_2512_v66 = _dafny.CodePoint('w')
				_2485_v45 = Companion_Default___.Fm47(_2512_v66, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2487_v48, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2487_v48).Cardinality()))).Uint32(), !((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool)))).Cardinality()), p1, globalState)
				var _index461 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index461
				(_2501_v57).ArraySet1(((_2500_v56).Difference(_2500_v56)).IsProperSubsetOf(_2500_v56), (_index461).Int())
				var _2513_v67 _dafny.Set
				_ = _2513_v67
				_2513_v67 = _dafny.SetOf(_2412_v0)
				var _index462 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _index462
				var _rhs432 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2484_v44, _dafny.Companion_Sequence_.Concatenate(_2484_v44, _2484_v44))
				_ = _rhs432
				var _rhs433 bool = ((_dafny.SetOf(true, _2412_v0, (_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), _2412_v0, p2)).Intersection(_2513_v67)).IsDisjointFrom(_2513_v67)
				_ = _rhs433
				var _rhs434 bool = !(((p0).Times((Companion_Default___.Fm49(_2487_v48, (_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), p0, globalState)).Cardinality())).Cmp((func() _dafny.Set {
					var _coll76 = _dafny.NewBuilder()
					_ = _coll76
					for _iter81 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(879), _dafny.IntOfInt64(273))); ; {
						_compr_76, _ok81 := _iter81()
						if !_ok81 {
							break
						}
						var _2514_v68 _dafny.Int
						_2514_v68 = interface{}(_compr_76).(_dafny.Int)
						if ((_dafny.IntOfInt64(879)).Cmp(_2514_v68) <= 0) && ((_2514_v68).Cmp(_dafny.IntOfInt64(273)) < 0) {
							_coll76.Add(Companion_Default___.SafeDivisionInt(_2514_v68, p1))
						}
					}
					return _coll76.ToSet()
				}()).Cardinality()) < 0)
				_ = _rhs434
				var _lhs217 _dafny.Array = _2501_v57
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
				_ = _lhs218
				_2484_v44 = _rhs432
				_2412_v0 = _rhs433
				(_lhs217).ArraySet1(_rhs434, (_lhs218).Int())
				var _2515_v69 D14
				_ = _2515_v69
				_2515_v69 = Companion_D14_.Create_DC28_()
				var _2516_v70 _dafny.Sequence
				_ = _2516_v70
				_2516_v70 = _dafny.SeqOf(_2515_v69, Companion_Default___.Fm51(_2484_v44, false, _2412_v0, globalState))
				var _rhs435 bool = (_2412_v0) || (Companion_Default___.Fm5(globalState))
				_ = _rhs435
				var _rhs436 _dafny.Sequence = _2516_v70
				_ = _rhs436
				_2412_v0 = _rhs435
				_2516_v70 = _rhs436
			}
			var _2517_v71 _dafny.CodePoint
			_ = _2517_v71
			_2517_v71 = _dafny.CodePoint('k')
			var _2518_v72 _dafny.Map
			_ = _2518_v72
			_2518_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm36((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), _2517_v71, p0, globalState)).Cardinality()), p2)
			var _2519_v73 _dafny.Array
			_ = _2519_v73
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_54
			var _nw453 _dafny.Array
			_ = _nw453
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw453 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Map = (func(_2520_v49 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_2521_i12 _dafny.Int) _dafny.Map {
						return _2520_v49
					}
				})(_2488_v49)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw453 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw453).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw453).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2519_v73 = _nw453
			var _2522_v74 _dafny.Array
			_ = _2522_v74
			_2522_v74 = _2519_v73
			var _2523_v75 _dafny.MultiSet
			_ = _2523_v75
			_2523_v75 = _dafny.MultiSetOf(_2522_v74, _2522_v74, _2522_v74, _2522_v74)
			var _2524_v76 _dafny.Map
			_ = _2524_v76
			_2524_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2523_v75)
			var _2525_v77 _dafny.Map
			_ = _2525_v77
			_2525_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2518_v72).Cardinality(), _dafny.MultiSetOf(_2522_v74))).Merge(_2524_v76), _2412_v0)
			_2525_v77 = (_2525_v77).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2523_v75), !(p2) || (_2412_v0))
			var _2526_v78 *C0
			_ = _2526_v78
			var _nw454 *C0 = New_C0_()
			_ = _nw454
			_nw454.Ctor__((_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool), (_2501_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))).Int()).(bool))
			_2526_v78 = _nw454
			var _2527_v79 _dafny.MultiSet
			_ = _2527_v79
			_2527_v79 = _dafny.MultiSetOf(p0)
			var _index463 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
			_ = _index463
			var _rhs437 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2484_v44, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer142 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg142 _dafny.Int) interface{} {
					return coer142(arg142)
				}
			}((func(_2528_v71 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2529_i13 _dafny.Int) _dafny.CodePoint {
					return _2528_v71
				}
			})(_2517_v71))), _dafny.UnicodeSeqOfUtf8Bytes("fgyncltxh")))
			_ = _rhs437
			var _rhs438 bool = (p1).Cmp((_2527_v79).Cardinality()) <= 0
			_ = _rhs438
			var _lhs219 _dafny.Array = _2501_v57
			_ = _lhs219
			var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2501_v57), 0))
			_ = _lhs220
			_2484_v44 = _rhs437
			(_lhs219).ArraySet1(_rhs438, (_lhs220).Int())
		}
		var _2530_v80 _dafny.Sequence
		_ = _2530_v80
		_2530_v80 = _dafny.SeqOf(true, Companion_Default___.Fm3(false, p0, p1, globalState), _2412_v0, p2)
		var _2531_v81 _dafny.MultiSet
		_ = _2531_v81
		_2531_v81 = _dafny.MultiSetOf(!(_2412_v0))
		var _2532_v82 _dafny.Array
		_ = _2532_v82
		var _nwElement0_107 bool = _2412_v0
		_ = _nwElement0_107
		var _nw455 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_107, nil, _dafny.IntOfInt64(10))
		_ = _nw455
		(_nw455).ArraySet1(_nwElement0_107, 0)
		(_nw455).ArraySet1(true, 1)
		(_nw455).ArraySet1(!(!(_2412_v0)), 2)
		(_nw455).ArraySet1(_2412_v0, 3)
		(_nw455).ArraySet1((_2530_v80).Select((Companion_Default___.SafeIndex((_2531_v81).Cardinality(), _dafny.IntOfUint32((_2530_v80).Cardinality()))).Uint32()).(bool), 4)
		(_nw455).ArraySet1(p2, 5)
		(_nw455).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_2530_v80, _2530_v80), 6)
		(_nw455).ArraySet1(_2412_v0, 7)
		(_nw455).ArraySet1(p2, 8)
		(_nw455).ArraySet1(!((p2) && (p2)), 9)
		_2532_v82 = _nw455
		for _iter82 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2532_v82), 0))); ; {
			_guard_loop_5, _ok82 := _iter82()
			if !_ok82 {
				break
			}
			var _2533_i14 _dafny.Int
			_2533_i14 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_2533_i14).Sign() != -1) && ((_2533_i14).Cmp(_dafny.ArrayLen((_2532_v82), 0)) < 0)) {
				(_2532_v82).ArraySet1((Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).Cmp(_dafny.IntOfInt64(900)) > 0, (_2533_i14).Int())
			}
		}
		var _2534_v83 D6
		_ = _2534_v83
		_2534_v83 = Companion_D6_.Create_DC13_(!(_2412_v0), p1)
		var _pat_let_tv56 = p0
		_ = _pat_let_tv56
		var _pat_let_tv57 = p1
		_ = _pat_let_tv57
		var _pat_let_tv58 = p1
		_ = _pat_let_tv58
		r0 = func(_source35 D6) _dafny.Int {
			if _source35.Is_DC12() {
				var _2535___mcc_h0 bool = _source35.Get_().(D6_DC12).Cf22
				_ = _2535___mcc_h0
				var _2536___mcc_h1 bool = _source35.Get_().(D6_DC12).Cf23
				_ = _2536___mcc_h1
				var _2537_cf23 bool = _2536___mcc_h1
				_ = _2537_cf23
				var _2538_cf22 bool = _2535___mcc_h0
				_ = _2538_cf22
				return _pat_let_tv56
			} else if _source35.Is_DC13() {
				var _2539___mcc_h2 bool = _source35.Get_().(D6_DC13).Cf24
				_ = _2539___mcc_h2
				var _2540___mcc_h3 _dafny.Int = _source35.Get_().(D6_DC13).Cf25
				_ = _2540___mcc_h3
				var _2541_cf25 _dafny.Int = _2540___mcc_h3
				_ = _2541_cf25
				var _2542_cf24 bool = _2539___mcc_h2
				_ = _2542_cf24
				return _pat_let_tv57
			} else {
				var _2543___mcc_h4 _dafny.Sequence = _source35.Get_().(D6_DC11).Cf21
				_ = _2543___mcc_h4
				var _2544_cf21 _dafny.Sequence = _2543___mcc_h4
				_ = _2544_cf21
				return _pat_let_tv58
			}
		}(_2534_v83)
		var _2545_v84 _dafny.Array
		_ = _2545_v84
		var _nw456 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw456
		_2545_v84 = _nw456
		var _index464 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2545_v84), 0))
		_ = _index464
		(_2545_v84).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(308), p1), (_index464).Int())
		var _index465 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_2545_v84), 0))
		_ = _index465
		(_2545_v84).ArraySet1(p1, (_index465).Int())
		var _2546_v85 D4
		_ = _2546_v85
		_2546_v85 = Companion_D4_.Create_DC6_(p0, _2545_v84, p0)
		_2546_v85 = _2546_v85
		r0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_2531_v81).Cardinality()), (p0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())))
		return r0
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
