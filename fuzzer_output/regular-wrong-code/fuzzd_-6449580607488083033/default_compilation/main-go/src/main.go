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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('v'), _dafny.CodePoint('x'), _dafny.CodePoint('l'))).Cardinality(), _dafny.IntOfInt64(216))).Cardinality())).Times(_dafny.IntOfInt64(906)), (func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("jyfwejrpi")
	})(), true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(168))).Cardinality())).Cardinality())), true)
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((Companion_D14_.Create_DC32_(true, (_dafny.SetOf(_dafny.IntOfInt64(-907))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.SeqOf(_dafny.CodePoint('s')), true)).Dtor_cf55(), _dafny.UnicodeSeqOfUtf8Bytes("racyt")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("irsogxd"), _dafny.UnicodeSeqOfUtf8Bytes("maxquypl")))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(932), _dafny.IntOfInt64(220))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wgfhyrgb")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
	})()).Cardinality()).Minus(_dafny.IntOfInt64(417))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 _dafny.MultiSet = (func() _dafny.MultiSet {
		if !(true) {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(-389))
		}
		return _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-550))))
	})()
	_ = _source0
	var _1___mcc_h0 _dafny.MultiSet = _source0
	_ = _1___mcc_h0
	var _2_cf18 _dafny.MultiSet = _1___mcc_h0
	_ = _2_cf18
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_2_cf18).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(67)))
}
func (_static *CompanionStruct_Default___) Fm4(p0 D0, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(879), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), false)).Cardinality()))).Contains((_dafny.SetOf(true)).Cardinality()) {
		return _dafny.CodePoint('o')
	} else {
		return _dafny.CodePoint('a')
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_((func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('u')
		}
		return _dafny.CodePoint('p')
	})())
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(-113))).Union(_dafny.SetOf((_dafny.SetOf(!(true))).Cardinality()))).Intersection((Companion_D6_.Create_DC13_(_dafny.IntOfInt64(7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwvbfw")).Cardinality()), _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-416)), _dafny.IntOfInt64(686)))).Dtor_cf22())
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(!(true), true))).Intersection(_dafny.SetOf(_dafny.SeqOf(true), (_dafny.SeqOf(true))))).Difference((_dafny.SetOf(_dafny.SeqOf(false, false))).Difference(_dafny.SetOf(_dafny.SeqOf(true, false))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality())).Difference((_dafny.MultiSetOf(_dafny.IntOfInt64(-786), _dafny.IntOfInt64(57))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(!(false) || (false), !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(false), _dafny.SetOf(false)), _dafny.SetOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(true, !(true), true, !((_dafny.IntOfInt64(-314)).Cmp((_dafny.SetOf(_dafny.IntOfInt64(653))).Cardinality()) != 0), !(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(597))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), _dafny.IntOfInt64(125))).Cardinality()
	})), (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true)))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(595), _dafny.IntOfInt64(298))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(595)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(298)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_4_v0, _dafny.IntOfInt64(642)))
			}
		}
		return _coll0.ToSet()
	}()).Intersection((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pejhxpg")).Cardinality()), _dafny.IntOfInt64(161))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v1 _dafny.Int
			_5_v1 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pejhxpg")).Cardinality()), _dafny.IntOfInt64(161)), _5_v1) {
				_coll1.Add((_5_v1).Times(_dafny.IntOfInt64(-835)))
			}
		}
		return _coll1.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfInt64(24), _dafny.IntOfInt64(-476), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(262), _dafny.CodePoint('v'))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(true), ((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(true, true, true, !(true), false)).Cardinality()))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), _dafny.IntOfInt64(782), _dafny.IntOfInt64(-196), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ef")).Cardinality())), _dafny.IntOfInt64(464)))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	if !(false) || (true) {
		return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.CodePoint('h')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('u')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('o')))
	} else {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('s'))))).Intersection(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('e'))))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(false))))).Difference(_dafny.MultiSetOf(true, !(false)))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ufhvh"), _dafny.CodePoint('e')))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Map, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Difference(_dafny.SetOf(true, false))).Intersection(_dafny.SetOf(true, true))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.MultiSet {
	return (func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(-988), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(632))).Cardinality())))
		}
		return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true, !(false))).Cardinality())).Cardinality())).Cardinality())
	})()
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('v'))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) D4 {
	var _source1 D0 = Companion_D0_.Create_DC2_()
	_ = _source1
	if _source1.Is_DC1() {
		var _6___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
		_ = _6___mcc_h0
		var _7___mcc_h1 _dafny.Sequence = _source1.Get_().(D0_DC1).Cf2
		_ = _7___mcc_h1
		var _8___mcc_h2 bool = _source1.Get_().(D0_DC1).Cf3
		_ = _8___mcc_h2
		var _9___mcc_h3 _dafny.Int = _source1.Get_().(D0_DC1).Cf4
		_ = _9___mcc_h3
		var _10___mcc_h4 bool = _source1.Get_().(D0_DC1).Cf5
		_ = _10___mcc_h4
		var _11_cf5 bool = _10___mcc_h4
		_ = _11_cf5
		var _12_cf4 _dafny.Int = _9___mcc_h3
		_ = _12_cf4
		var _13_cf3 bool = _8___mcc_h2
		_ = _13_cf3
		var _14_cf2 _dafny.Sequence = _7___mcc_h1
		_ = _14_cf2
		var _15_cf1 _dafny.Int = _6___mcc_h0
		_ = _15_cf1
		return Companion_D4_.Create_DC10_(_dafny.CodePoint('m'), _dafny.CodePoint('a'))
	} else if _source1.Is_DC2() {
		if true {
			return Companion_D4_.Create_DC10_(_dafny.CodePoint('s'), _dafny.CodePoint('d'))
		} else {
			return Companion_D4_.Create_DC10_(_dafny.CodePoint('w'), _dafny.CodePoint('c'))
		}
	} else {
		var _16___mcc_h5 bool = _source1.Get_().(D0_DC0).Cf0
		_ = _16___mcc_h5
		var _17_cf0 bool = _16___mcc_h5
		_ = _17_cf0
		return Companion_D4_.Create_DC10_(_dafny.CodePoint('w'), _dafny.CodePoint('f'))
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 D6, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true), !(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true, false), true)))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) D11 {
	if (!(!(true))) == (false) {
		return Companion_D11_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
	} else {
		return Companion_D11_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.CodePoint('l'))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(687)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(566)))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(false)).Cardinality())).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(459)))).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(138))))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Union(_dafny.MultiSetOf(!(false), true)), (Companion_D4_.Create_DC10_(_dafny.CodePoint('h'), _dafny.CodePoint('v'))).Dtor_cf16())
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(297), Companion_D14_.Create_DC32_(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjwmcxnlo")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('f')), false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC9_(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC32_((_dafny.SetOf(_dafny.CodePoint('n'), _dafny.CodePoint('f'))).IsProperSubsetOf((Companion_D22_.Create_DC50_(_dafny.SetOf(_dafny.CodePoint('w')))).Dtor_cf80()), (_dafny.IntOfInt64(290)).Plus(_dafny.IntOfInt64(-436)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('h')), _dafny.SeqOf(_dafny.CodePoint('r'), _dafny.CodePoint('b'))), true)
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC14_(((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(715), _dafny.IntOfInt64(609))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(330), !(true))).Cardinality())).Cardinality())))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(260), _dafny.IntOfInt64(-500)))).Cardinality(), _dafny.IntOfInt64(586), (_dafny.MultiSetOf(!(false))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(848)))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality())))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(848)))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality())))).Contains(_18_v0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_18_v0, _dafny.IntOfInt64(-263)), true)
			}
		}
		return _coll2.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-464), _dafny.IntOfInt64(557))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-464)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(557)) < 0) {
				_coll3.Add(Companion_Default___.SafeModuloInt(_19_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(false, true)).Cardinality(), (_dafny.SetOf(false, true)).Cardinality())).Cardinality()))), _dafny.IntOfInt64(200))
			}
		}
		return _coll3.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(282))).Uint32(), func(coer2 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_20_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_21_i1 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(328))).Cardinality()
		})))
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(28))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_22_i2 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(true)).Cardinality()
	}))), _dafny.MultiSetOf(_dafny.IntOfInt64(647), _dafny.IntOfInt64(-625)), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), true)).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(645), _dafny.IntOfInt64(732), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_23_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality())), (func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-771))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_24_i4 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
		}))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _25_v0 _dafny.Int
			_25_v0 = interface{}(_compr_4).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-771))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_24_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
			})), _25_v0) {
				_coll4.Add((_25_v0).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_26_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-688))).Cardinality())
				}))).Cardinality()))))
			}
		}
		return _coll4.ToSet()
	}()).Cardinality()), _dafny.MultiSetOf(_dafny.IntOfInt64(687), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(-742))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(586))).Cardinality())))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _27_v0 _dafny.Int
	_ = _27_v0
	_27_v0 = _dafny.IntOfInt64(629)
	var _28_i0 _dafny.Int
	_ = _28_i0
	_28_i0 = _dafny.Zero
	{
		for (_27_v0).Cmp(_27_v0) < 0 {
			{
				if (_28_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_28_i0 = (_28_i0).Plus(_dafny.One)
				if p0 {
					var _29_v1 *C1
					_ = _29_v1
					var _nw0 *C1 = New_C1_()
					_ = _nw0
					_nw0.Ctor__(p0)
					_29_v1 = _nw0
					(globalState).F2 = _27_v0
					_27_v0 = (_dafny.Zero).Minus(_27_v0)
					var _30_v2 D0
					_ = _30_v2
					_30_v2 = Companion_D0_.Create_DC0_((_29_v1).F32())
					var _31_v3 *C2
					_ = _31_v3
					var _nw1 *C2 = New_C2_()
					_ = _nw1
					_nw1.Ctor__(_27_v0, p0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-851)))
					_31_v3 = _nw1
					var _32_v4 _dafny.Map
					_ = _32_v4
					_32_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v2, _dafny.IntOfUint32((_dafny.SeqOf(_31_v3)).Cardinality()))
					var _33_v5 D10
					_ = _33_v5
					_33_v5 = Companion_D10_.Create_DC23_()
					var _34_v6 _dafny.Sequence
					_ = _34_v6
					_34_v6 = _dafny.UnicodeSeqOfUtf8Bytes("ehxqgq")
					var _35_v7 _dafny.MultiSet
					_ = _35_v7
					_35_v7 = _dafny.MultiSetOf((_31_v3).F31(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}(func(_36_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					}))).Cardinality()), _dafny.IntOfUint32((_34_v6).Cardinality()), _dafny.IntOfInt64(411))
					var _37_v8 _dafny.Map
					_ = _37_v8
					_37_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_29_v1).F32()), (_29_v1).F32())
					var _38_v9 _dafny.Sequence
					_ = _38_v9
					_38_v9 = _dafny.SeqOf((_29_v1).F32(), p0, p0)
					var _39_v10 _dafny.Array
					_ = _39_v10
					var _nwElement0_0 bool = Companion_Default___.Fm5((_29_v1).F32(), _32_v4, (_31_v3).F31(), globalState)
					_ = _nwElement0_0
					var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(26))
					_ = _nw2
					(_nw2).ArraySet1(_nwElement0_0, 0)
					(_nw2).ArraySet1(p0, 1)
					(_nw2).ArraySet1((_dafny.IntOfInt64(-184)).Cmp((_31_v3).F31()) >= 0, 2)
					(_nw2).ArraySet1(p0, 3)
					(_nw2).ArraySet1(p0, 4)
					(_nw2).ArraySet1((_29_v1).F32(), 5)
					(_nw2).ArraySet1((_29_v1).F32(), 6)
					(_nw2).ArraySet1(p0, 7)
					(_nw2).ArraySet1(p0, 8)
					(_nw2).ArraySet1(((_dafny.SetOf(_33_v5, _33_v5)).Cardinality()).Cmp(_27_v0) <= 0, 9)
					(_nw2).ArraySet1((_27_v0).Cmp(_dafny.IntOfInt64(362)) < 0, 10)
					(_nw2).ArraySet1((_29_v1).F32(), 11)
					(_nw2).ArraySet1((_27_v0).Cmp(_27_v0) > 0, 12)
					(_nw2).ArraySet1(p0, 13)
					(_nw2).ArraySet1((_29_v1).F32(), 14)
					(_nw2).ArraySet1((_35_v7).Equals(_35_v7), 15)
					(_nw2).ArraySet1((_29_v1).F32(), 16)
					(_nw2).ArraySet1((_dafny.IntOfUint32((_34_v6).Cardinality())).Cmp(_dafny.IntOfInt64(144)) == 0, 17)
					(_nw2).ArraySet1(!(_37_v8).Equals((_37_v8).Update(p0, (_38_v9).Select((Companion_Default___.SafeIndex(_27_v0, _dafny.IntOfUint32((_38_v9).Cardinality()))).Uint32()).(bool))), 18)
					(_nw2).ArraySet1((_dafny.IntOfUint32((_34_v6).Cardinality())).Cmp((_31_v3).F31()) < 0, 19)
					(_nw2).ArraySet1(p0, 20)
					(_nw2).ArraySet1((_29_v1).F32(), 21)
					(_nw2).ArraySet1(p0, 22)
					(_nw2).ArraySet1((_29_v1).F32(), 23)
					(_nw2).ArraySet1(!(p0) || ((_29_v1).F32()), 24)
					(_nw2).ArraySet1(!(((_31_v3).F31()).Cmp((_31_v3).F31()) <= 0), 25)
					_39_v10 = _nw2
					var _rhs0 _dafny.Array = _39_v10
					_ = _rhs0
					var _rhs1 bool = !((_38_v9).Select((Companion_Default___.SafeIndex((_31_v3).F31(), _dafny.IntOfUint32((_38_v9).Cardinality()))).Uint32()).(bool))
					_ = _rhs1
					var _lhs0 *GlobalState = globalState
					_ = _lhs0
					_39_v10 = _rhs0
					_lhs0.F1 = _rhs1
					var _40_v11 _dafny.MultiSet
					_ = _40_v11
					var _41_v12 _dafny.Map
					_ = _41_v12
					var _out0 _dafny.MultiSet
					_ = _out0
					var _out1 _dafny.Map
					_ = _out1
					_out0, _out1 = (_29_v1).M7((_31_v3).Fm7(globalState), _30_v2, _34_v6, globalState)
					_40_v11 = _out0
					_41_v12 = _out1
				} else {
					var _42_v13 *C2
					_ = _42_v13
					var _nw3 *C2 = New_C2_()
					_ = _nw3
					_nw3.Ctor__(_27_v0, !(p0), Companion_Default___.SafeModuloInt(_27_v0, _27_v0))
					_42_v13 = _nw3
					(globalState).F0 = (_dafny.Zero).Minus(_27_v0)
					var _43_v14 _dafny.Sequence
					_ = _43_v14
					_43_v14 = _dafny.UnicodeSeqOfUtf8Bytes("myjeh")
					(globalState).F1 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("iplftk"), _43_v14), _dafny.Companion_Sequence_.Concatenate(_43_v14, _43_v14))
					var _44_v15 _dafny.Map
					_ = _44_v15
					_44_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v0, p0)
					var _45_v16 _dafny.Map
					_ = _45_v16
					_45_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hfaji")).Cardinality()), _27_v0)
					(globalState).F1 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_45_v16), Companion_Default___.Fm43((_44_v15).Cardinality(), p0, (_dafny.MultiSetOf(_27_v0)).Cardinality(), (_dafny.Zero).Minus((_42_v13).F31()), globalState))
					var _46_v17 _dafny.Map
					_ = _46_v17
					_46_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _47_v18 _dafny.Map
					_ = _47_v18
					_47_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v0, _46_v17)
					var _48_v19 _dafny.Sequence
					_ = _48_v19
					_48_v19 = _dafny.SeqOf(_dafny.IntOfInt64(75), _27_v0, (_42_v13).F31())
					var _49_v20 _dafny.Map
					_ = _49_v20
					_49_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						if (_47_v18).Contains(_dafny.IntOfUint32((_48_v19).Cardinality())) {
							return (_47_v18).Get(_dafny.IntOfUint32((_48_v19).Cardinality())).(_dafny.Map)
						}
						return _46_v17
					})(), (_42_v13).F31())
					_49_v20 = (_49_v20).Update(_46_v17, _27_v0)
				}
				var _50_v21 D0
				_ = _50_v21
				_50_v21 = Companion_D0_.Create_DC0_(p0)
				var _pat_let_tv0 = p0
				_ = _pat_let_tv0
				var _51_v22 _dafny.Map
				_ = _51_v22
				_51_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let0_0 D0) D0 {
					return func(_52_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let1_0 bool) D0 {
							return func(_53_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC0_(_53_dt__update_hcf0_h0)
							}(_pat_let1_0)
						}(_pat_let_tv0)
					}(_pat_let0_0)
				}(_50_v21), _27_v0)
				var _54_v23 _dafny.Set
				_ = _54_v23
				_54_v23 = _dafny.SetOf(p0, false)
				var _55_v24 *C2
				_ = _55_v24
				var _nw4 *C2 = New_C2_()
				_ = _nw4
				_nw4.Ctor__((_dafny.Zero).Minus((_27_v0).Minus(_27_v0)), Companion_Default___.Fm5(p0, _51_v22, _27_v0, globalState), (_54_v23).Cardinality())
				_55_v24 = _nw4
				var _56_v25 _dafny.Sequence
				_ = _56_v25
				_56_v25 = _dafny.UnicodeSeqOfUtf8Bytes("ww")
				var _57_v26 _dafny.Sequence
				_ = _57_v26
				_57_v26 = _dafny.SeqOf((_dafny.Zero).Minus(((_55_v24).F31()).Minus(_dafny.IntOfUint32((_56_v25).Cardinality()))), _27_v0, Companion_Default___.SafeModuloInt((_55_v24).F31(), (_55_v24).F31()))
				(globalState).F0 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_57_v26).Cardinality()))
				var _58_v27 _dafny.MultiSet
				_ = _58_v27
				_58_v27 = _dafny.MultiSetOf(p0)
				_58_v27 = _58_v27
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _59_v28 _dafny.MultiSet
	_ = _59_v28
	_59_v28 = _dafny.MultiSetOf(_27_v0, _27_v0, Companion_Default___.Fm2(_27_v0, _27_v0, globalState))
	var _60_v29 _dafny.Sequence
	_ = _60_v29
	_60_v29 = _dafny.UnicodeSeqOfUtf8Bytes("hfda")
	var _61_v30 _dafny.Array
	_ = _61_v30
	var _nwElement0_1 _dafny.Int = _27_v0
	_ = _nwElement0_1
	var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(11))
	_ = _nw5
	(_nw5).ArraySet1(_nwElement0_1, 0)
	(_nw5).ArraySet1(Companion_Default___.SafeDivisionInt(_27_v0, _27_v0), 1)
	(_nw5).ArraySet1(_27_v0, 2)
	(_nw5).ArraySet1(((func() _dafny.Int {
		if (_59_v28).Contains(_27_v0) {
			return (_59_v28).Multiplicity(_27_v0)
		}
		return _27_v0
	})()).Plus(_27_v0), 3)
	(_nw5).ArraySet1(_27_v0, 4)
	(_nw5).ArraySet1(Companion_Default___.SafeModuloInt(_27_v0, _27_v0), 5)
	(_nw5).ArraySet1((_dafny.MultiSetFromSeq(Companion_Default___.Fm44(globalState))).Cardinality(), 6)
	(_nw5).ArraySet1(_27_v0, 7)
	(_nw5).ArraySet1(_27_v0, 8)
	(_nw5).ArraySet1(_dafny.IntOfUint32((_60_v29).Cardinality()), 9)
	(_nw5).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ewsm")).Cardinality()), _27_v0), 10)
	_61_v30 = _nw5
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))
	_ = _index0
	(_61_v30).ArraySet1(_27_v0, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))
	_ = _index1
	(_61_v30).ArraySet1(_27_v0, (_index1).Int())
	(globalState).F0 = _27_v0
	_60_v29 = _dafny.UnicodeSeqOfUtf8Bytes("fba")
	(globalState).F0 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_62_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality())
	if !(!(p0) || (!(p0))) {
		(globalState).F0 = ((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)).Plus(Companion_Default___.SafeDivisionInt(_27_v0, _27_v0))
		var _63_v31 _dafny.Map
		_ = _63_v31
		_63_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_v29, (_dafny.Zero).Minus(_27_v0))
		_63_v31 = (_63_v31).Update(_60_v29, _dafny.IntOfInt64(970))
		(globalState).F2 = (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)
		var _64_v32 _dafny.Map
		_ = _64_v32
		_64_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v0, p0)
		_64_v32 = (_64_v32).Update((_dafny.SetOf(p0, p0)).Cardinality(), p0)
		var _65_v33 D0
		_ = _65_v33
		_65_v33 = Companion_D0_.Create_DC0_(true)
		var _66_v34 D0
		_ = _66_v34
		_66_v34 = Companion_D0_.Create_DC0_(Companion_Default___.Fm5(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v33, _dafny.IntOfUint32((_60_v29).Cardinality())), _27_v0, globalState))
		var _67_v35 _dafny.Map
		_ = _67_v35
		_67_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v33, (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int))
		var _68_v36 *C3
		_ = _68_v36
		var _nw6 *C3 = New_C3_()
		_ = _nw6
		_nw6.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("nj"), !(p0), p0, (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int))
		_68_v36 = _nw6
		var _69_v37 _dafny.Set
		_ = _69_v37
		_69_v37 = _dafny.SetOf(_68_v36)
		var _70_v38 _dafny.Sequence
		_ = _70_v38
		_70_v38 = _dafny.SeqOf(false)
		var _71_v39 _dafny.Array
		_ = _71_v39
		var _nwElement0_2 bool = (_68_v36).F30()
		_ = _nwElement0_2
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(23))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_2, 0)
		(_nw7).ArraySet1((_70_v38).Select((Companion_Default___.SafeIndex((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_70_v38).Cardinality()))).Uint32()).(bool), 1)
		(_nw7).ArraySet1(p0, 2)
		(_nw7).ArraySet1((_68_v36).F30(), 3)
		(_nw7).ArraySet1((_68_v36).F30(), 4)
		(_nw7).ArraySet1((_68_v36).F30(), 5)
		(_nw7).ArraySet1(p0, 6)
		(_nw7).ArraySet1(p0, 7)
		(_nw7).ArraySet1((_68_v36).F30(), 8)
		(_nw7).ArraySet1(p0, 9)
		(_nw7).ArraySet1(false, 10)
		(_nw7).ArraySet1((_68_v36).F30(), 11)
		(_nw7).ArraySet1(p0, 12)
		(_nw7).ArraySet1((_68_v36).F30(), 13)
		(_nw7).ArraySet1(!((_68_v36).F30()), 14)
		(_nw7).ArraySet1(true, 15)
		(_nw7).ArraySet1((_68_v36).F30(), 16)
		(_nw7).ArraySet1(!(p0), 17)
		(_nw7).ArraySet1(p0, 18)
		(_nw7).ArraySet1((_68_v36).F30(), 19)
		(_nw7).ArraySet1(p0, 20)
		(_nw7).ArraySet1(p0, 21)
		(_nw7).ArraySet1(p0, 22)
		_71_v39 = _nw7
		var _72_v40 *C0
		_ = _72_v40
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__(_60_v29, Companion_Default___.Fm5(!(p0), _67_v35, (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int), globalState), (_69_v37).Cardinality(), _71_v39, (_68_v36).F30())
		_72_v40 = _nw8
		var _73_v41 D17
		_ = _73_v41
		_73_v41 = Companion_D17_.Create_DC39_(Companion_D17_.Create_DC38_())
		var _74_v42 *C2
		_ = _74_v42
		var _nw9 *C2 = New_C2_()
		_ = _nw9
		_nw9.Ctor__((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int), false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(p0, _67_v35, _dafny.IntOfUint32((_70_v38).Cardinality()), globalState), (_68_v36).F30())).Cardinality())
		_74_v42 = _nw9
		var _75_v43 D17
		_ = _75_v43
		_75_v43 = Companion_D17_.Create_DC37_(_74_v42)
		var _76_v44 _dafny.Sequence
		_ = _76_v44
		_76_v44 = _dafny.SeqOf((_73_v41).Dtor_cf64(), _75_v43, _75_v43, _75_v43)
		var _77_v45 _dafny.Sequence
		_ = _77_v45
		_77_v45 = _dafny.SeqOf(_72_v40.F35)
		var _pat_let_tv1 = _76_v44
		_ = _pat_let_tv1
		var _pat_let_tv2 = _61_v30
		_ = _pat_let_tv2
		var _pat_let_tv3 = _61_v30
		_ = _pat_let_tv3
		var _pat_let_tv4 = _76_v44
		_ = _pat_let_tv4
		var _rhs2 *C0 = _72_v40
		_ = _rhs2
		var _rhs3 D17 = func(_pat_let2_0 D17) D17 {
			return func(_78_dt__update__tmp_h1 D17) D17 {
				return func(_pat_let3_0 D17) D17 {
					return func(_79_dt__update_hcf64_h0 D17) D17 {
						return Companion_D17_.Create_DC39_(_79_dt__update_hcf64_h0)
					}(_pat_let3_0)
				}((_pat_let_tv1).Select((Companion_Default___.SafeIndex((_pat_let_tv3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_pat_let_tv2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_pat_let_tv4).Cardinality()))).Uint32()).(D17))
			}(_pat_let2_0)
		}(_73_v41)
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_70_v38, _70_v38), _70_v38)
		_ = _rhs4
		var _rhs5 _dafny.Int = (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)
		_ = _rhs5
		var _rhs6 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int), _27_v0))).Minus((_dafny.IntOfUint32((_77_v45).Cardinality())).Times((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)))
		_ = _rhs6
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		_72_v40 = _rhs2
		_73_v41 = _rhs3
		_70_v38 = _rhs4
		_lhs1.F2 = _rhs5
		_lhs2.F0 = _rhs6
	} else {
		var _80_v46 _dafny.CodePoint
		_ = _80_v46
		_80_v46 = _dafny.CodePoint('r')
		_80_v46 = _80_v46
		var _81_v47 _dafny.Array
		_ = _81_v47
		var _nw10 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
		_ = _nw10
		_81_v47 = _nw10
		_81_v47 = _81_v47
		var _82_v48 _dafny.Sequence
		_ = _82_v48
		_82_v48 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_83_v29 _dafny.Sequence, _84_v30 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_85_i3 _dafny.Int) _dafny.CodePoint {
				return (_83_v29).Select((Companion_Default___.SafeIndex((_84_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_84_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_83_v29).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_60_v29, _61_v30))), _60_v29)
		var _86_v49 _dafny.Array
		_ = _86_v49
		var _nw11 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
		_ = _nw11
		_86_v49 = _nw11
		var _87_v50 _dafny.Map
		_ = _87_v50
		_87_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v48, _86_v49)
		var _88_v51 D21
		_ = _88_v51
		_88_v51 = Companion_D21_.Create_DC48_(_86_v49)
		_87_v50 = (_87_v50).Update(_dafny.Companion_Sequence_.Concatenate(_82_v48, _82_v48), (_88_v51).Dtor_cf77())
		var _89_v52 _dafny.Array
		_ = _89_v52
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_0
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_90_p0 bool) func(_dafny.Int) bool {
				return func(_91_i4 _dafny.Int) bool {
					return _90_p0
				}
			})(p0)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw12 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw12).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw12).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_89_v52 = _nw12
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_89_v52), 0))
		_ = _index2
		(_89_v52).ArraySet1((_59_v28).IsSubsetOf(_dafny.MultiSetOf((_dafny.Zero).Minus((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)))), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_89_v52), 0))
		_ = _index3
		(_89_v52).ArraySet1(p0, (_index3).Int())
		var _92_v53 *C1
		_ = _92_v53
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__((_89_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_89_v52), 0))).Int()).(bool))
		_92_v53 = _nw13
		var _93_v54 D10
		_ = _93_v54
		_93_v54 = Companion_D10_.Create_DC24_(_27_v0, _92_v53, _80_v46)
		var _94_v55 D10
		_ = _94_v55
		_94_v55 = Companion_D10_.Create_DC24_((_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int), (_93_v54).Dtor_cf37(), _80_v46)
		var _pat_let_tv5 = _92_v53
		_ = _pat_let_tv5
		var _pat_let_tv6 = _80_v46
		_ = _pat_let_tv6
		_94_v55 = func(_pat_let4_0 D10) D10 {
			return func(_95_dt__update__tmp_h2 D10) D10 {
				return func(_pat_let5_0 *C1) D10 {
					return func(_96_dt__update_hcf37_h0 *C1) D10 {
						return func(_pat_let6_0 _dafny.CodePoint) D10 {
							return func(_97_dt__update_hcf38_h0 _dafny.CodePoint) D10 {
								return Companion_D10_.Create_DC24_((_95_dt__update__tmp_h2).Dtor_cf36(), _96_dt__update_hcf37_h0, _97_dt__update_hcf38_h0)
							}(_pat_let6_0)
						}(_pat_let_tv6)
					}(_pat_let5_0)
				}(_pat_let_tv5)
			}(_pat_let4_0)
		}(_94_v55)
	}
	r0 = (_61_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_61_v30), 0))).Int()).(_dafny.Int)
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _98_v0 _dafny.Array
	_ = _98_v0
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
	_ = _nw14
	_98_v0 = _nw14
	var _99_v1 _dafny.Sequence
	_ = _99_v1
	_99_v1 = _dafny.UnicodeSeqOfUtf8Bytes("oduxsblwp")
	var _100_v2 bool
	_ = _100_v2
	_100_v2 = false
	var _101_v3 _dafny.Int
	_ = _101_v3
	_101_v3 = _dafny.IntOfInt64(189)
	var _102_v4 _dafny.Map
	_ = _102_v4
	_102_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v2, _101_v3)
	var _103_v5 _dafny.Sequence
	_ = _103_v5
	_103_v5 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v2, _101_v3), _102_v4)
	var _104_v6 _dafny.Set
	_ = _104_v6
	_104_v6 = _dafny.SetOf(_101_v3)
	var _105_globalState *GlobalState
	_ = _105_globalState
	var _nw15 *GlobalState = New_GlobalState_()
	_ = _nw15
	_nw15.Ctor__(_dafny.IntOfInt64(281), false, _dafny.IntOfInt64(-82), _98_v0, true, _dafny.IntOfInt64(-974), _dafny.IntOfInt64(777), false, _dafny.Companion_Sequence_.Concatenate(_99_v1, _dafny.UnicodeSeqOfUtf8Bytes("n")), _103_v5, _dafny.IntOfInt64(-345), true, _dafny.UnicodeSeqOfUtf8Bytes("yrxdsa"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v2, _dafny.UnicodeSeqOfUtf8Bytes("jxiyve")), _dafny.IntOfInt64(362), _dafny.IntOfInt64(72), _dafny.IntOfInt64(327), (_104_v6).Union(_104_v6), true, _dafny.IntOfInt64(-736), true, true, false, false, true)
	_105_globalState = _nw15
	_100_v2 = _100_v2
	var _106_v7 _dafny.Array
	_ = _106_v7
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_1
	var _nw16 _dafny.Array
	_ = _nw16
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw16 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = (func(_107_v2 bool, _108_v1 _dafny.Sequence) func(_dafny.Int) bool {
			return func(_109_i0 _dafny.Int) bool {
				return !((_107_v2) && ((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-205), _108_v1, _107_v2, _dafny.IntOfUint32((_108_v1).Cardinality()), _107_v2)).Dtor_cf5()))
			}
		})(_100_v2, _99_v1)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw16 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw16).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw16).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_106_v7 = _nw16
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
	_ = _index4
	(_106_v7).ArraySet1(!(_100_v2) || (_100_v2), (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
	_ = _index5
	(_106_v7).ArraySet1((_101_v3).Cmp(_101_v3) <= 0, (_index5).Int())
	var _110_i1 _dafny.Int
	_ = _110_i1
	_110_i1 = _dafny.Zero
	{
		for !(((_102_v4).Cardinality()).Cmp(_101_v3) <= 0) || (!(false)) {
			{
				if (_110_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_110_i1 = (_110_i1).Plus(_dafny.One)
				(_105_globalState).F0 = _101_v3
				var _111_v8 _dafny.MultiSet
				_ = _111_v8
				_111_v8 = _dafny.MultiSetOf(_100_v2)
				var _112_v9 _dafny.CodePoint
				_ = _112_v9
				_112_v9 = _dafny.CodePoint('w')
				var _113_v10 _dafny.MultiSet
				_ = _113_v10
				_113_v10 = _dafny.MultiSetOf(_112_v9)
				var _114_v11 D0
				_ = _114_v11
				_114_v11 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(315), Companion_Default___.Fm1(_105_globalState), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), Companion_Default___.Fm2((func() _dafny.Int {
					if (_111_v8).Contains(_100_v2) {
						return (_111_v8).Multiplicity(_100_v2)
					}
					return (_113_v10).Cardinality()
				})(), _101_v3, _105_globalState), _100_v2)
				var _115_v12 _dafny.Sequence
				_ = _115_v12
				_115_v12 = _dafny.SeqOf(_101_v3)
				var _116_v13 _dafny.Map
				_ = _116_v13
				_116_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_115_v12).Cardinality()), _101_v3)
				var _117_v14 _dafny.Array
				_ = _117_v14
				var _nwElement0_3 D0 = Companion_Default___.Fm0(_101_v3, _105_globalState)
				_ = _nwElement0_3
				var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(9))
				_ = _nw17
				(_nw17).ArraySet1(_nwElement0_3, 0)
				(_nw17).ArraySet1(_114_v11, 1)
				(_nw17).ArraySet1(_114_v11, 2)
				(_nw17).ArraySet1(_114_v11, 3)
				(_nw17).ArraySet1(Companion_Default___.Fm0(_101_v3, _105_globalState), 4)
				(_nw17).ArraySet1(_114_v11, 5)
				(_nw17).ArraySet1(_114_v11, 6)
				(_nw17).ArraySet1(_114_v11, 7)
				(_nw17).ArraySet1(Companion_D0_.Create_DC1_((func() _dafny.Int {
					if (_116_v13).Contains(_101_v3) {
						return (_116_v13).Get(_101_v3).(_dafny.Int)
					}
					return _101_v3
				})(), _99_v1, false, _101_v3, _100_v2), 8)
				_117_v14 = _nw17
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_117_v14), 0))
				_ = _index6
				(_117_v14).ArraySet1(_114_v11, (_index6).Int())
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_117_v14), 0))
				_ = _index7
				(_117_v14).ArraySet1(_114_v11, (_index7).Int())
				var _118_v15 _dafny.Map
				_ = _118_v15
				_118_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v3, _100_v2)
				(_105_globalState).F0 = (_115_v12).Select((Companion_Default___.SafeIndex((_118_v15).Cardinality(), _dafny.IntOfUint32((_115_v12).Cardinality()))).Uint32()).(_dafny.Int)
				_111_v8 = (_111_v8).Union(_111_v8)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _hi0 _dafny.Int = _101_v3
	_ = _hi0
	for _119_i2 := _101_v3; _119_i2.Cmp(_hi0) < 0; _119_i2 = _119_i2.Plus(_dafny.One) {
		var _120_v16 _dafny.CodePoint
		_ = _120_v16
		_120_v16 = _dafny.CodePoint('o')
		_120_v16 = _120_v16
		var _121_v17 _dafny.Sequence
		_ = _121_v17
		_121_v17 = _dafny.SeqOf(_101_v3)
		var _122_v18 _dafny.MultiSet
		_ = _122_v18
		_122_v18 = _dafny.MultiSetOf((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _100_v2, _100_v2, _100_v2)
		var _123_v19 _dafny.Array
		_ = _123_v19
		var _nwElement0_4 _dafny.Int = _dafny.IntOfInt64(-393)
		_ = _nwElement0_4
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_4, 0)
		(_nw18).ArraySet1((func() _dafny.Int {
			if false {
				return _101_v3
			}
			return _101_v3
		})(), 1)
		(_nw18).ArraySet1(Companion_Default___.SafeDivisionInt(_119_i2, _119_i2), 2)
		(_nw18).ArraySet1(_101_v3, 3)
		(_nw18).ArraySet1(_101_v3, 4)
		(_nw18).ArraySet1(_119_i2, 5)
		(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_121_v17, (Companion_Default___.SafeIndex(_119_i2, _dafny.IntOfUint32((_121_v17).Cardinality()))).Uint32(), _101_v3), _121_v17)).Cardinality()), 6)
		(_nw18).ArraySet1((func() _dafny.Int {
			if (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool) {
				return (_122_v18).Cardinality()
			}
			return _119_i2
		})(), 7)
		(_nw18).ArraySet1((_121_v17).Select((Companion_Default___.SafeIndex(_101_v3, _dafny.IntOfUint32((_121_v17).Cardinality()))).Uint32()).(_dafny.Int), 8)
		_123_v19 = _nw18
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))
		_ = _index8
		(_123_v19).ArraySet1(Companion_Default___.SafeModuloInt(_101_v3, (_122_v18).Cardinality()), (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))
		_ = _index9
		(_123_v19).ArraySet1(Companion_Default___.SafeDivisionInt(_101_v3, _101_v3), (_index9).Int())
		var _124_v20 D0
		_ = _124_v20
		_124_v20 = Companion_D0_.Create_DC0_((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
		var _source2 D0 = _124_v20
		_ = _source2
		if _source2.Is_DC1() {
			var _125___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
			_ = _125___mcc_h0
			var _126___mcc_h1 _dafny.Sequence = _source2.Get_().(D0_DC1).Cf2
			_ = _126___mcc_h1
			var _127___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf3
			_ = _127___mcc_h2
			var _128___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf4
			_ = _128___mcc_h3
			var _129___mcc_h4 bool = _source2.Get_().(D0_DC1).Cf5
			_ = _129___mcc_h4
			var _130_cf5 bool = _129___mcc_h4
			_ = _130_cf5
			var _131_cf4 _dafny.Int = _128___mcc_h3
			_ = _131_cf4
			var _132_cf3 bool = _127___mcc_h2
			_ = _132_cf3
			var _133_cf2 _dafny.Sequence = _126___mcc_h1
			_ = _133_cf2
			var _134_cf1 _dafny.Int = _125___mcc_h0
			_ = _134_cf1
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index10
			(_106_v7).ArraySet1(_130_cf5, (_index10).Int())
			(_105_globalState).F0 = _119_i2
			var _135_v21 _dafny.Sequence
			_ = _135_v21
			_135_v21 = _dafny.SeqOf(_132_cf3, _100_v2)
			(_105_globalState).F1 = (_135_v21).Select((Companion_Default___.SafeIndex(_134_cf1, _dafny.IntOfUint32((_135_v21).Cardinality()))).Uint32()).(bool)
			var _136_v22 D0
			_ = _136_v22
			_136_v22 = Companion_D0_.Create_DC2_()
			var _137_v23 _dafny.Array
			_ = _137_v23
			var _nwElement0_5 D0 = _136_v22
			_ = _nwElement0_5
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(15))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_5, 0)
			(_nw19).ArraySet1(_136_v22, 1)
			(_nw19).ArraySet1(_136_v22, 2)
			(_nw19).ArraySet1(Companion_D0_.Create_DC2_(), 3)
			(_nw19).ArraySet1(_136_v22, 4)
			(_nw19).ArraySet1(_136_v22, 5)
			(_nw19).ArraySet1(_136_v22, 6)
			(_nw19).ArraySet1(_136_v22, 7)
			(_nw19).ArraySet1(_136_v22, 8)
			(_nw19).ArraySet1((func() D0 {
				if _100_v2 {
					return _136_v22
				}
				return _136_v22
			})(), 9)
			(_nw19).ArraySet1(Companion_D0_.Create_DC2_(), 10)
			(_nw19).ArraySet1(Companion_D0_.Create_DC2_(), 11)
			(_nw19).ArraySet1(_136_v22, 12)
			(_nw19).ArraySet1(Companion_D0_.Create_DC2_(), 13)
			(_nw19).ArraySet1(_136_v22, 14)
			_137_v23 = _nw19
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_137_v23), 0))
			_ = _index11
			(_137_v23).ArraySet1(_136_v22, (_index11).Int())
			var _138_v24 _dafny.Map
			_ = _138_v24
			_138_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(_130_cf5), _132_cf3, _100_v2)).Cardinality()), (_119_i2).Times((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int)))
			var _139_v25 _dafny.Array
			_ = _139_v25
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_2
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_140_cf2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_141_i3 _dafny.Int) _dafny.Sequence {
						return _140_cf2
					}
				})(_133_cf2)
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
			_139_v25 = _nw20
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_139_v25), 0))
			_ = _index12
			(_139_v25).ArraySet1((func() _dafny.Sequence {
				if _132_cf3 {
					return _133_cf2
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_142_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_143_i4 _dafny.Int) _dafny.CodePoint {
						return _142_v16
					}
				})(_120_v16)))
			})(), (_index12).Int())
			var _144_v26 _dafny.Array
			_ = _144_v26
			var _nwElement0_6 _dafny.Sequence = _121_v17
			_ = _nwElement0_6
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(15))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_6, 0)
			(_nw21).ArraySet1(_121_v17, 1)
			(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_121_v17, _121_v17), 2)
			(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_121_v17, _121_v17), 3)
			(_nw21).ArraySet1(_121_v17, 4)
			(_nw21).ArraySet1((func() _dafny.Sequence {
				if _100_v2 {
					return _121_v17
				}
				return _dafny.SeqOf((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int), _101_v3, _101_v3, _134_cf1, _dafny.IntOfUint32((_133_cf2).Cardinality()))
			})(), 5)
			(_nw21).ArraySet1(_121_v17, 6)
			(_nw21).ArraySet1(_121_v17, 7)
			(_nw21).ArraySet1(_121_v17, 8)
			(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_121_v17, _121_v17), 9)
			(_nw21).ArraySet1(_121_v17, 10)
			(_nw21).ArraySet1(_dafny.SeqOf(_119_i2, _134_cf1, (_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int), _101_v3, _101_v3), 11)
			(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int), _134_cf1), _121_v17), 12)
			(_nw21).ArraySet1(Companion_Default___.Fm3(_134_cf1, _132_cf3, _101_v3, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _105_globalState), 13)
			(_nw21).ArraySet1(_121_v17, 14)
			_144_v26 = _nw21
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_144_v26), 0))
			_ = _index13
			(_144_v26).ArraySet1(_121_v17, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_137_v23), 0))
			_ = _index14
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_139_v25), 0))
			_ = _index15
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_144_v26), 0))
			_ = _index16
			var _rhs7 D0 = _136_v22
			_ = _rhs7
			var _rhs8 _dafny.Map = _138_v24
			_ = _rhs8
			var _rhs9 _dafny.Sequence = _99_v1
			_ = _rhs9
			var _rhs10 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_121_v17, _dafny.SeqOf(_101_v3, Companion_Default___.Fm2(_119_i2, _134_cf1, _105_globalState))), (Companion_Default___.SafeIndex(Companion_Default___.Fm2(_101_v3, (func() _dafny.Set {
				var _coll5 = _dafny.NewBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(388), _dafny.IntOfInt64(253))); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _145_v27 _dafny.Int
					_145_v27 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(388)).Cmp(_145_v27) <= 0) && ((_145_v27).Cmp(_dafny.IntOfInt64(253)) < 0) {
						_coll5.Add((_145_v27).Plus(_dafny.IntOfUint32((_121_v17).Cardinality())))
					}
				}
				return _coll5.ToSet()
			}()).Cardinality(), _105_globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_121_v17, _dafny.SeqOf(_101_v3, Companion_Default___.Fm2(_119_i2, _134_cf1, _105_globalState)))).Cardinality()))).Uint32(), (_dafny.IntOfUint32((_135_v21).Cardinality())).Times(_131_cf4))
			_ = _rhs10
			var _lhs3 _dafny.Array = _137_v23
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_137_v23), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = _139_v25
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_139_v25), 0))
			_ = _lhs6
			var _lhs7 _dafny.Array = _144_v26
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_144_v26), 0))
			_ = _lhs8
			(_lhs3).ArraySet1(_rhs7, (_lhs4).Int())
			_138_v24 = _rhs8
			(_lhs5).ArraySet1(_rhs9, (_lhs6).Int())
			(_lhs7).ArraySet1(_rhs10, (_lhs8).Int())
		} else if _source2.Is_DC2() {
			var _146_v28 _dafny.Sequence
			_ = _146_v28
			_146_v28 = _dafny.SeqOf(_121_v17, _121_v17)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index17
			var _rhs11 bool = !(_100_v2)
			_ = _rhs11
			var _rhs12 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-947))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_147_v17 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_148_i5 _dafny.Int) _dafny.Sequence {
					return _147_v17
				}
			})(_121_v17)))
			_ = _rhs12
			var _rhs13 _dafny.Int = (_dafny.Zero).Minus(_119_i2)
			_ = _rhs13
			var _lhs9 _dafny.Array = _106_v7
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = _105_globalState
			_ = _lhs11
			(_lhs9).ArraySet1(_rhs11, (_lhs10).Int())
			_146_v28 = _rhs12
			_lhs11.F0 = _rhs13
			var _149_v29 _dafny.MultiSet
			_ = _149_v29
			_149_v29 = _dafny.MultiSetOf(_120_v16, _120_v16, Companion_Default___.Fm4(Companion_D0_.Create_DC1_(_101_v3, _99_v1, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _119_i2, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), !(_100_v2), (_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int), _105_globalState))
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))
			_ = _index19
			var _rhs14 _dafny.MultiSet = _149_v29
			_ = _rhs14
			var _rhs15 bool = ((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int)).Cmp((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int)) <= 0
			_ = _rhs15
			var _rhs16 bool = _100_v2
			_ = _rhs16
			var _rhs17 _dafny.Int = _119_i2
			_ = _rhs17
			var _lhs12 _dafny.Array = _106_v7
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _123_v19
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))
			_ = _lhs15
			_149_v29 = _rhs14
			(_lhs12).ArraySet1(_rhs15, (_lhs13).Int())
			_100_v2 = _rhs16
			(_lhs14).ArraySet1(_rhs17, (_lhs15).Int())
			(_105_globalState).F2 = _119_i2
			_100_v2 = (_101_v3).Cmp(_119_i2) < 0
		} else {
			var _150___mcc_h5 bool = _source2.Get_().(D0_DC0).Cf0
			_ = _150___mcc_h5
			var _151_cf0 bool = _150___mcc_h5
			_ = _151_cf0
			(_105_globalState).F1 = _100_v2
			var _152_v30 _dafny.Int
			_ = _152_v30
			var _out2 _dafny.Int
			_ = _out2
			_out2 = Companion_Default___.M0(!(_151_cf0), _105_globalState)
			_152_v30 = _out2
			var _153_v31 _dafny.Map
			_ = _153_v31
			_153_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v20, _119_i2)
			_151_cf0 = !(Companion_Default___.Fm5(_151_cf0, (_153_v31).Merge(_153_v31), _152_v30, _105_globalState))
			var _154_v32 _dafny.Int
			_ = _154_v32
			var _out3 _dafny.Int
			_ = _out3
			_out3 = Companion_Default___.M0((_100_v2) == (_100_v2), _105_globalState)
			_154_v32 = _out3
		}
		var _155_v33 *C1
		_ = _155_v33
		var _nw22 *C1 = New_C1_()
		_ = _nw22
		_nw22.Ctor__((_dafny.SetOf(false)).Equals((Companion_Default___.Fm39((_123_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_123_v19), 0))).Int()).(_dafny.Int), !((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), _105_globalState)).Dtor_cf15()))
		_155_v33 = _nw22
	}
	var _156_v34 _dafny.Sequence
	_ = _156_v34
	_156_v34 = _dafny.SeqOf(!(_100_v2), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
	var _157_v35 _dafny.Sequence
	_ = _157_v35
	_157_v35 = _dafny.SeqOf(_101_v3)
	var _158_v36 _dafny.CodePoint
	_ = _158_v36
	_158_v36 = _dafny.CodePoint('r')
	var _159_v38 D0
	_ = _159_v38
	_159_v38 = Companion_D0_.Create_DC0_((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
	var _160_v39 _dafny.Sequence
	_ = _160_v39
	_160_v39 = _dafny.SeqOf(Companion_D0_.Create_DC0_((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), Companion_D0_.Create_DC0_(_100_v2), _159_v38)
	var _161_v40 _dafny.Map
	_ = _161_v40
	_161_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v36, func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_160_v39).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _162_v37 D0
			_162_v37 = interface{}(_compr_6).(D0)
			if _dafny.Companion_Sequence_.Contains(_160_v39, _162_v37) {
				_coll6.Add(_162_v37, (_102_v4).Cardinality())
			}
		}
		return _coll6.ToMap()
	}())
	var _163_v41 *C3
	_ = _163_v41
	var _nw23 *C3 = New_C3_()
	_ = _nw23
	_nw23.Ctor__((func() _dafny.Sequence {
		if !((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)) {
			return _99_v1
		}
		return _99_v1
	})(), _100_v2, Companion_Default___.Fm5((_156_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_157_v35).Cardinality()), _dafny.IntOfUint32((_156_v34).Cardinality()))).Uint32()).(bool), (func() _dafny.Map {
		if (_161_v40).Contains(_158_v36) {
			return (_161_v40).Get(_158_v36).(_dafny.Map)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_v38, _dafny.IntOfUint32((_99_v1).Cardinality()))
	})(), _101_v3, _105_globalState), Companion_Default___.SafeDivisionInt(_101_v3, _101_v3))
	_163_v41 = _nw23
	var _164_v42 *C2
	_ = _164_v42
	var _nw24 *C2 = New_C2_()
	_ = _nw24
	_nw24.Ctor__((_dafny.Zero).Minus(_101_v3), !((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), _101_v3)
	_164_v42 = _nw24
	_164_v42 = _164_v42
	var _165_v43 _dafny.Array
	_ = _165_v43
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
	_ = _nw25
	_165_v43 = _nw25
	var _166_v44 _dafny.Array
	_ = _166_v44
	var _nw26 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
	_ = _nw26
	_166_v44 = _nw26
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_165_v43), 0))
	_ = _index20
	(_165_v43).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).F31(), _166_v44), (_index20).Int())
	var _167_v45 _dafny.Map
	_ = _167_v45
	_167_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).F31(), _166_v44)
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_165_v43), 0))
	_ = _index21
	(_165_v43).ArraySet1(_167_v45, (_index21).Int())
	var _hi1 _dafny.Int = (_164_v42).F31()
	_ = _hi1
	for _168_i6 := (_163_v41).Fm6(_100_v2, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _101_v3, _105_globalState); _168_i6.Cmp(_hi1) < 0; _168_i6 = _168_i6.Plus(_dafny.One) {
		if (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool) {
			var _169_v46 _dafny.Map
			_ = _169_v46
			_169_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_i6, _156_v34)
			var _170_v47 _dafny.Sequence
			_ = _170_v47
			_170_v47 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v3, _156_v34), _169_v46, _169_v46)
			_102_v4 = (_102_v4).Update((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), ((_170_v47).Select((Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_170_v47).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
			_158_v36 = _dafny.CodePoint('f')
			var _171_v48 _dafny.Map
			_ = _171_v48
			_171_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v2, _99_v1)
			(_105_globalState).F1 = (func() bool {
				if (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool) {
					return !((_163_v41).F30())
				}
				return ((_dafny.Zero).Minus(_101_v3)).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_171_v48).Contains(_100_v2) {
						return (_171_v48).Get(_100_v2).(_dafny.Sequence)
					}
					return (_163_v41).F29()
				})()).Cardinality())) != 0
			})()
			var _172_v49 _dafny.Map
			_ = _172_v49
			_172_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).Fm11(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v1, _99_v1), (_164_v42).F31(), _dafny.IntOfUint32(((_163_v41).F29()).Cardinality()), _105_globalState), _163_v41)
			_172_v49 = (_172_v49).Update(_168_i6, _163_v41)
			var _173_v50 _dafny.Array
			_ = _173_v50
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw27
			_173_v50 = _nw27
			var _174_v51 _dafny.Array
			_ = _174_v51
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_3
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_175_v42 *C2) func(_dafny.Int) _dafny.Int {
					return func(_176_i7 _dafny.Int) _dafny.Int {
						return (_176_i7).Minus((_175_v42).F31())
					}
				})(_164_v42)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw28 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw28).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw28).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_174_v51 = _nw28
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_173_v50), 0))
			_ = _index22
			(_173_v50).ArraySet1(_174_v51, (_index22).Int())
			var _177_v52 D12
			_ = _177_v52
			_177_v52 = Companion_D12_.Create_DC28_(_174_v51)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_173_v50), 0))
			_ = _index23
			(_173_v50).ArraySet1((_177_v52).Dtor_cf44(), (_index23).Int())
		} else {
			(_105_globalState).F2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(805), _168_i6)
			var _rhs18 _dafny.Sequence = (_163_v41).F29()
			_ = _rhs18
			var _rhs19 _dafny.Int = (_164_v42).Fm6((_163_v41).F30(), (_163_v41).F30(), (_dafny.Zero).Minus((_164_v42).F31()), _105_globalState)
			_ = _rhs19
			var _rhs20 bool = _100_v2
			_ = _rhs20
			var _lhs16 *GlobalState = _105_globalState
			_ = _lhs16
			_99_v1 = _rhs18
			_lhs16.F2 = _rhs19
			_100_v2 = _rhs20
			var _178_v53 _dafny.Set
			_ = _178_v53
			_178_v53 = _dafny.SetOf(_106_v7, _106_v7, _106_v7, _106_v7)
			(_105_globalState).F1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_99_v1, (Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_99_v1).Cardinality()))).Uint32(), _158_v36), _99_v1)).Cardinality())).Cmp(((_164_v42).F31()).Times((_178_v53).Cardinality())) == 0
			_106_v7 = (func() _dafny.Array {
				if (_163_v41).F30() {
					return _106_v7
				}
				return _106_v7
			})()
			var _179_v54 _dafny.Array
			_ = _179_v54
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_4
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) D9 = (func(_180_v41 *C3, _181_v3 _dafny.Int) func(_dafny.Int) D9 {
					return func(_182_i8 _dafny.Int) D9 {
						return Companion_D9_.Create_DC19_(Companion_D6_.Create_DC14_(_dafny.IntOfUint32(((_180_v41).F29()).Cardinality()), _181_v3, _dafny.MultiSetOf((_180_v41).F30())))
					}
				})(_163_v41, _101_v3)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw29 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw29).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw29).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_179_v54 = _nw29
			_179_v54 = _179_v54
		}
		var _183_v55 _dafny.Array
		_ = _183_v55
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
		_ = _nw30
		_183_v55 = _nw30
		var _184_v56 _dafny.Map
		_ = _184_v56
		_184_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)) && ((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), _183_v55)
		_183_v55 = (func() _dafny.Array {
			if (_184_v56).Contains((_163_v41).F30()) {
				return (_184_v56).Get((_163_v41).F30()).(_dafny.Array)
			}
			return _183_v55
		})()
		(_105_globalState).F2 = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_164_v42).F31(), (_164_v42).F31())).Plus(_dafny.IntOfInt64(377)))
		var _185_v57 *C1
		_ = _185_v57
		var _nw31 *C1 = New_C1_()
		_ = _nw31
		_nw31.Ctor__((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
		_185_v57 = _nw31
	}
	if !(!((_163_v41).F30()) || (!(!_dafny.Companion_Sequence_.Equal((_163_v41).F29(), _99_v1)))) {
		if (_163_v41).F30() {
			var _186_v58 *C1
			_ = _186_v58
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__((_163_v41).F30())
			_186_v58 = _nw32
			_186_v58 = _186_v58
			var _187_v59 _dafny.MultiSet
			_ = _187_v59
			_187_v59 = _dafny.MultiSetOf((_163_v41).F30(), (_186_v58).F32(), (_186_v58).F32())
			var _188_v60 _dafny.Map
			_ = _188_v60
			_188_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).F31(), !((_163_v41).F30()))
			var _rhs21 bool = (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)
			_ = _rhs21
			var _rhs22 _dafny.Int = (func() _dafny.Int {
				if false {
					return (_164_v42).F31()
				}
				return (_dafny.IntOfUint32(((_163_v41).F29()).Cardinality())).Plus((_164_v42).F31())
			})()
			_ = _rhs22
			var _rhs23 bool = (_163_v41).F30()
			_ = _rhs23
			var _rhs24 _dafny.Int = (_164_v42).F31()
			_ = _rhs24
			var _rhs25 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_187_v59).Contains((_163_v41).F30()) {
					return (_187_v59).Multiplicity((_163_v41).F30())
				}
				return (_164_v42).F31()
			})(), (_163_v41).F30())).Merge(_188_v60)).Cardinality()
			_ = _rhs25
			var _lhs17 *GlobalState = _105_globalState
			_ = _lhs17
			var _lhs18 *GlobalState = _105_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _105_globalState
			_ = _lhs19
			var _lhs20 *GlobalState = _105_globalState
			_ = _lhs20
			var _lhs21 *GlobalState = _105_globalState
			_ = _lhs21
			_lhs17.F1 = _rhs21
			_lhs18.F0 = _rhs22
			_lhs19.F1 = _rhs23
			_lhs20.F2 = _rhs24
			_lhs21.F0 = _rhs25
			var _189_v61 D17
			_ = _189_v61
			_189_v61 = Companion_D17_.Create_DC37_(_164_v42)
			var _190_v62 _dafny.Sequence
			_ = _190_v62
			_190_v62 = _dafny.SeqOf(_164_v42, _164_v42, (_189_v61).Dtor_cf63(), (func() *C2 {
				if (_186_v58).F32() {
					return _164_v42
				}
				return _164_v42
			})(), _164_v42)
			_190_v62 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_190_v62, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_156_v34).Cardinality()), _dafny.IntOfUint32((_190_v62).Cardinality()))).Uint32(), _164_v42), _dafny.SeqOf(_164_v42, _164_v42, _164_v42, _164_v42, _164_v42))
			(_105_globalState).F2 = (func() _dafny.Int {
				if (_156_v34).Select((Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_156_v34).Cardinality()))).Uint32()).(bool) {
					return (_164_v42).F31()
				}
				return _101_v3
			})()
			var _191_v63 _dafny.Array
			_ = _191_v63
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw33
			_191_v63 = _nw33
			var _192_v64 _dafny.Sequence
			_ = _192_v64
			_192_v64 = _dafny.SeqOf(_191_v63, (func() _dafny.Array {
				if (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool) {
					return _191_v63
				}
				return _191_v63
			})(), _191_v63, _191_v63)
			(_105_globalState).F0 = _dafny.IntOfUint32((_192_v64).Cardinality())
		} else {
			var _193_v65 *C4
			_ = _193_v65
			var _nw34 *C4 = New_C4_()
			_ = _nw34
			_nw34.Ctor__(_101_v3, (_163_v41).F29(), (_163_v41).F30(), (_164_v42).F31())
			_193_v65 = _nw34
			var _194_v66 _dafny.Array
			_ = _194_v66
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw35
			_194_v66 = _nw35
			_194_v66 = _194_v66
			var _195_v67 D16
			_ = _195_v67
			_195_v67 = Companion_D16_.Create_DC36_((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _157_v35, _106_v7)
			var _196_v68 _dafny.Int
			_ = _196_v68
			var _197_v69 bool
			_ = _197_v69
			var _198_v70 _dafny.Map
			_ = _198_v70
			var _199_v71 _dafny.Map
			_ = _199_v71
			var _out4 _dafny.Int
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 _dafny.Map
			_ = _out6
			var _out7 _dafny.Map
			_ = _out7
			_out4, _out5, _out6, _out7 = (_193_v65).M2((_195_v67).Dtor_cf60(), _105_globalState)
			_196_v68 = _out4
			_197_v69 = _out5
			_198_v70 = _out6
			_199_v71 = _out7
			_99_v1 = (_163_v41).F29()
			var _200_v72 _dafny.Array
			_ = _200_v72
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw36
			_200_v72 = _nw36
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_200_v72), 0))
			_ = _index24
			(_200_v72).ArraySet1(_194_v66, (_index24).Int())
			var _201_v73 _dafny.Array
			_ = _201_v73
			var _nwElement0_7 _dafny.Sequence = _157_v35
			_ = _nwElement0_7
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_7, 0)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _157_v35), 1)
			(_nw37).ArraySet1(_157_v35, 2)
			(_nw37).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-357))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_202_v42 *C2) func(_dafny.Int) _dafny.Int {
				return func(_203_i9 _dafny.Int) _dafny.Int {
					return (_202_v42).F31()
				}
			})(_164_v42))), 3)
			(_nw37).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(430))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_204_v65 *C4) func(_dafny.Int) _dafny.Int {
				return func(_205_i10 _dafny.Int) _dafny.Int {
					return (_204_v65).F27()
				}
			})(_193_v65))), 4)
			(_nw37).ArraySet1(_157_v35, 5)
			(_nw37).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_206_v42 *C2) func(_dafny.Int) _dafny.Int {
				return func(_207_i11 _dafny.Int) _dafny.Int {
					return (_206_v42).F31()
				}
			})(_164_v42))), 6)
			_201_v73 = _nw37
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_201_v73), 0))
			_ = _index25
			(_201_v73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _157_v35), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_200_v72), 0))
			_ = _index26
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_201_v73), 0))
			_ = _index27
			var _rhs26 _dafny.Array = _194_v66
			_ = _rhs26
			var _rhs27 _dafny.Int = (_164_v42).F31()
			_ = _rhs27
			var _rhs28 _dafny.Sequence = (Companion_Default___.Fm40(_100_v2, _105_globalState)).Dtor_cf55()
			_ = _rhs28
			var _rhs29 _dafny.Sequence = _dafny.SeqOf(_101_v3, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_163_v41).F30() {
					return Companion_Default___.Fm1(_105_globalState)
				}
				return (_163_v41).F29()
			})()).Cardinality()), (_164_v42).F31())
			_ = _rhs29
			var _lhs22 _dafny.Array = _200_v72
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_200_v72), 0))
			_ = _lhs23
			var _lhs24 *GlobalState = _105_globalState
			_ = _lhs24
			var _lhs25 *C4 = _193_v65
			_ = _lhs25
			var _lhs26 _dafny.Array = _201_v73
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_201_v73), 0))
			_ = _lhs27
			(_lhs22).ArraySet1(_rhs26, (_lhs23).Int())
			_lhs24.F2 = _rhs27
			_lhs25.F28 = _rhs28
			(_lhs26).ArraySet1(_rhs29, (_lhs27).Int())
		}
		(_105_globalState).F1 = (_163_v41).F30()
		_156_v34 = _dafny.Companion_Sequence_.Concatenate(_156_v34, _156_v34)
		var _208_v74 _dafny.Map
		_ = _208_v74
		_208_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _100_v2)
		var _209_v75 *C4
		_ = _209_v75
		var _nw38 *C4 = New_C4_()
		_ = _nw38
		_nw38.Ctor__((_208_v74).Cardinality(), _dafny.Companion_Sequence_.Update(_99_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_163_v41).F29()).Cardinality()), _dafny.IntOfUint32((_99_v1).Cardinality()))).Uint32(), _158_v36), _100_v2, (_164_v42).F31())
		_209_v75 = _nw38
		var _210_v76 _dafny.MultiSet
		_ = _210_v76
		_210_v76 = _dafny.MultiSetOf(_209_v75)
		var _211_v77 _dafny.Sequence
		_ = _211_v77
		_211_v77 = _dafny.SeqOf(_209_v75)
		var _212_v78 _dafny.MultiSet
		_ = _212_v78
		_212_v78 = _dafny.MultiSetFromSeq(_211_v77)
		var _213_v79 _dafny.Set
		_ = _213_v79
		_213_v79 = _dafny.SetOf(_210_v76, (_210_v76).Difference(_210_v76), _210_v76, (_212_v78), _210_v76)
		var _rhs30 _dafny.Int = _101_v3
		_ = _rhs30
		var _rhs31 _dafny.Int = Companion_Default___.Fm2(((_209_v75).F27()).Minus(_101_v3), (_209_v75).Fm6(_100_v2, (_163_v41).F30(), _101_v3, _105_globalState), _105_globalState)
		_ = _rhs31
		var _rhs32 _dafny.Set = _dafny.SetOf(_dafny.MultiSetOf(_209_v75, _209_v75, _209_v75))
		_ = _rhs32
		var _rhs33 _dafny.Int = Companion_Default___.SafeModuloInt((_163_v41).Fm9(_105_globalState), _dafny.IntOfInt64(915))
		_ = _rhs33
		var _lhs28 *GlobalState = _105_globalState
		_ = _lhs28
		var _lhs29 *GlobalState = _105_globalState
		_ = _lhs29
		_101_v3 = _rhs30
		_lhs28.F0 = _rhs31
		_213_v79 = _rhs32
		_lhs29.F0 = _rhs33
		var _214_v80 _dafny.Set
		_ = _214_v80
		_214_v80 = _dafny.SetOf(_156_v34, _156_v34, _156_v34, _156_v34)
		var _215_v81 _dafny.Sequence
		_ = _215_v81
		_215_v81 = _dafny.SeqOf(_156_v34, _dafny.SeqOf((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), _156_v34, _156_v34, _156_v34)
		(_105_globalState).F0 = (_dafny.Zero).Minus(((_214_v80).Union(func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_215_v81).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _216_v82 _dafny.Sequence
				_216_v82 = interface{}(_compr_7).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_215_v81, _216_v82) {
					_coll7.Add(_216_v82)
				}
			}
			return _coll7.ToSet()
		}())).Cardinality())
	} else {
		_99_v1 = _99_v1
		if (func() bool {
			if !((_163_v41).F30()) {
				return _100_v2
			}
			return !(true)
		})() {
			(_105_globalState).F0 = _101_v3
			(_105_globalState).F0 = _dafny.IntOfInt64(338)
			var _217_v83 *C1
			_ = _217_v83
			var _nw39 *C1 = New_C1_()
			_ = _nw39
			_nw39.Ctor__(false)
			_217_v83 = _nw39
			var _218_v84 *C4
			_ = _218_v84
			var _nw40 *C4 = New_C4_()
			_ = _nw40
			_nw40.Ctor__(_dafny.IntOfInt64(903), (_163_v41).F29(), (_217_v83).F32(), (_164_v42).F31())
			_218_v84 = _nw40
			var _219_v85 _dafny.MultiSet
			_ = _219_v85
			_219_v85 = _dafny.MultiSetOf(_218_v84, _218_v84)
			var _220_v86 D19
			_ = _220_v86
			_220_v86 = Companion_D19_.Create_DC41_(_218_v84)
			var _221_v87 _dafny.Sequence
			_ = _221_v87
			_221_v87 = _dafny.SeqOf(_dafny.SetOf((_218_v84).F27()))
			var _222_v88 _dafny.Sequence
			_ = _222_v88
			_222_v88 = _dafny.SeqOf(_157_v35)
			var _223_v89 _dafny.Sequence
			_ = _223_v89
			_223_v89 = _dafny.SeqOf(_106_v7)
			var _224_v90 _dafny.Map
			_ = _224_v90
			_224_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_218_v84).F27())).Cardinality(), _dafny.Companion_Sequence_.Update(_99_v1, (Companion_Default___.SafeIndex((_218_v84).F27(), _dafny.IntOfUint32((_99_v1).Cardinality()))).Uint32(), _158_v36))
			var _225_v91 _dafny.Map
			_ = _225_v91
			_225_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_163_v41).F30()), _106_v7)
			var _226_v92 D20
			_ = _226_v92
			_226_v92 = Companion_D20_.Create_DC44_(_225_v91)
			var _227_v93 _dafny.Map
			_ = _227_v93
			_227_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_226_v92).Dtor_cf73(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_228_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_229_i13 _dafny.Int) _dafny.CodePoint {
					return _228_v36
				}
			})(_158_v36))))
			var _230_v94 _dafny.Set
			_ = _230_v94
			_230_v94 = _dafny.SetOf(!((_164_v42).Fm7(_105_globalState)), (_217_v83).F32())
			var _231_v95 _dafny.Map
			_ = _231_v95
			_231_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v94, false)
			var _232_v96 _dafny.Map
			_ = _232_v96
			_232_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v95, (_164_v42).F31())
			var _233_v97 _dafny.Array
			_ = _233_v97
			var _nwElement0_8 _dafny.Int = (func() _dafny.Int {
				if (_219_v85).Contains((_220_v86).Dtor_cf66()) {
					return (_219_v85).Multiplicity((_220_v86).Dtor_cf66())
				}
				return (_218_v84).F27()
			})()
			_ = _nwElement0_8
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(22))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_8, 0)
			(_nw41).ArraySet1((_dafny.IntOfUint32(((_163_v41).F29()).Cardinality())).Minus((_218_v84).F27()), 1)
			(_nw41).ArraySet1((_164_v42).Fm11(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_234_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_235_i12 _dafny.Int) _dafny.CodePoint {
					return _234_v36
				}
			})(_158_v36))), _99_v1), _101_v3, (_164_v42).F31(), _105_globalState), 2)
			(_nw41).ArraySet1(Companion_Default___.SafeDivisionInt(_101_v3, (_164_v42).F31()), 3)
			(_nw41).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_221_v87).Cardinality()), (_218_v84).F27()), 4)
			(_nw41).ArraySet1((_dafny.Zero).Minus((_218_v84).F27()), 5)
			(_nw41).ArraySet1(_dafny.IntOfUint32((_222_v88).Cardinality()), 6)
			(_nw41).ArraySet1((_164_v42).F31(), 7)
			(_nw41).ArraySet1((_dafny.IntOfInt64(671)).Plus((_218_v84).F27()), 8)
			(_nw41).ArraySet1((_dafny.Zero).Minus((_164_v42).F31()), 9)
			(_nw41).ArraySet1(_101_v3, 10)
			(_nw41).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_223_v89).Select((Companion_Default___.SafeIndex((_224_v90).Cardinality(), _dafny.IntOfUint32((_223_v89).Cardinality()))).Uint32()).(_dafny.Array), (_164_v42).F31())).Cardinality(), 11)
			(_nw41).ArraySet1((_164_v42).F31(), 12)
			(_nw41).ArraySet1(Companion_Default___.SafeModuloInt((_164_v42).F31(), (_164_v42).F31()), 13)
			(_nw41).ArraySet1((_218_v84).F27(), 14)
			(_nw41).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm1(_105_globalState)).Cardinality()), 15)
			(_nw41).ArraySet1((_dafny.Zero).Minus((_164_v42).F31()), 16)
			(_nw41).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_227_v93).Contains(_225_v91) {
					return (_227_v93).Get(_225_v91).(_dafny.Sequence)
				}
				return (_163_v41).F29()
			})()).Cardinality())).Times(_dafny.IntOfInt64(890))), 17)
			(_nw41).ArraySet1(_101_v3, 18)
			(_nw41).ArraySet1(_dafny.IntOfInt64(-639), 19)
			(_nw41).ArraySet1((_164_v42).F31(), 20)
			(_nw41).ArraySet1((func() _dafny.Int {
				if (_232_v96).Contains(_231_v95) {
					return (_232_v96).Get(_231_v95).(_dafny.Int)
				}
				return (_218_v84).F27()
			})(), 21)
			_233_v97 = _nw41
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_233_v97), 0))
			_ = _index28
			(_233_v97).ArraySet1((_164_v42).F31(), (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_233_v97), 0))
			_ = _index29
			(_233_v97).ArraySet1((((_218_v84).F27()).Times((_164_v42).F31())).Times(_dafny.IntOfUint32(((_163_v41).F29()).Cardinality())), (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index30
			(_106_v7).ArraySet1((_164_v42).Fm7(_105_globalState), (_index30).Int())
		} else {
			(_105_globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("j"), (_163_v41).F29())
			_157_v35 = _dafny.SeqOf(_dafny.IntOfUint32(((_163_v41).F29()).Cardinality()))
			var _rhs34 _dafny.Sequence = (_163_v41).F29()
			_ = _rhs34
			var _rhs35 bool = (_163_v41).F30()
			_ = _rhs35
			var _lhs30 *GlobalState = _105_globalState
			_ = _lhs30
			_99_v1 = _rhs34
			_lhs30.F1 = _rhs35
			var _236_v98 _dafny.Set
			_ = _236_v98
			_236_v98 = _dafny.SetOf(_99_v1)
			(_105_globalState).F1 = (_236_v98).IsSubsetOf(_236_v98)
			var _237_v99 _dafny.Map
			_ = _237_v99
			_237_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v41).F30(), _100_v2)
			var _238_v100 _dafny.Set
			_ = _238_v100
			_238_v100 = _dafny.SetOf((_163_v41).F30())
			var _239_v101 *C4
			_ = _239_v101
			var _nw42 *C4 = New_C4_()
			_ = _nw42
			_nw42.Ctor__((_164_v42).F31(), (_163_v41).F29(), (_163_v41).F30(), (_238_v100).Cardinality())
			_239_v101 = _nw42
			var _240_v102 _dafny.Set
			_ = _240_v102
			_240_v102 = _dafny.SetOf(_239_v101, _239_v101)
			var _241_v103 D10
			_ = _241_v103
			_241_v103 = Companion_D10_.Create_DC23_()
			var _242_v104 *C0
			_ = _242_v104
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(_239_v101.F28, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _101_v3, _106_v7, (_163_v41).F30())
			_242_v104 = _nw43
			var _243_v105 _dafny.Map
			_ = _243_v105
			_243_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v104, _157_v35)
			var _244_v107 _dafny.MultiSet
			_ = _244_v107
			_244_v107 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_102_v4).Contains((_242_v104).Fm17(!((_163_v41).F30()), (_163_v41).F30(), _105_globalState)) {
					return (_102_v4).Get((_242_v104).Fm17(!((_163_v41).F30()), (_163_v41).F30(), _105_globalState)).(_dafny.Int)
				}
				return (_164_v42).F31()
			})(), (_164_v42).F31(), (_164_v42).F31())
			var _245_v108 T0
			_ = _245_v108
			var _nw44 *C3 = New_C3_()
			_ = _nw44
			_nw44.Ctor__(_239_v101.F28, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _dafny.IntOfInt64(637))
			_245_v108 = _nw44
			var _246_v109 _dafny.Map
			_ = _246_v109
			_246_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).F31(), _245_v108)
			var _247_v110 _dafny.Array
			_ = _247_v110
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_5
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_248_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_249_i14 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_249_i14, _248_v3)
					}
				})(_101_v3)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw45 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw45).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw45).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_247_v110 = _nw45
			var _250_v111 _dafny.Sequence
			_ = _250_v111
			_250_v111 = _dafny.SeqOf(_247_v110)
			var _251_v112 _dafny.Sequence
			_ = _251_v112
			_251_v112 = _dafny.SeqOf(Companion_Default___.Fm3((_246_v109).Cardinality(), (_163_v41).F30(), _dafny.IntOfUint32((_250_v111).Cardinality()), !((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), _105_globalState))
			var _252_v113 _dafny.Array
			_ = _252_v113
			var _nwElement0_9 _dafny.Sequence = _dafny.SeqOf((_163_v41).Fm6((_163_v41).F30(), (_163_v41).F30(), (_164_v42).F31(), _105_globalState))
			_ = _nwElement0_9
			var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(28))
			_ = _nw46
			(_nw46).ArraySet1(_nwElement0_9, 0)
			(_nw46).ArraySet1(_157_v35, 1)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _dafny.SeqOf(_dafny.IntOfInt64(125), (_237_v99).Cardinality(), (_164_v42).F31(), (_dafny.Zero).Minus((_164_v42).F31()))), 2)
			(_nw46).ArraySet1((func() _dafny.Sequence {
				if (_163_v41).F30() {
					return _dafny.Companion_Sequence_.Update(_157_v35, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_164_v42).F31()), _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32(), (_164_v42).F31())
				}
				return _157_v35
			})(), 3)
			(_nw46).ArraySet1(_157_v35, 4)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _157_v35), 5)
			(_nw46).ArraySet1(_157_v35, 6)
			(_nw46).ArraySet1(_dafny.SeqOf((_164_v42).F31()), 7)
			(_nw46).ArraySet1(_157_v35, 8)
			(_nw46).ArraySet1(_157_v35, 9)
			(_nw46).ArraySet1(_dafny.SeqOf((_dafny.MultiSetOf(_240_v102)).Cardinality()), 10)
			(_nw46).ArraySet1(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_241_v103, Companion_D10_.Create_DC23_()), _106_v7)).Cardinality(), _dafny.IntOfUint32((_157_v35).Cardinality()), (_164_v42).F31(), _101_v3, (_dafny.Zero).Minus((_164_v42).F31())), 11)
			(_nw46).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(811)), 12)
			(_nw46).ArraySet1(_157_v35, 13)
			(_nw46).ArraySet1(_157_v35, 14)
			(_nw46).ArraySet1(_157_v35, 15)
			(_nw46).ArraySet1(_157_v35, 16)
			(_nw46).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(307), (_164_v42).F31()), 17)
			(_nw46).ArraySet1((func() _dafny.Sequence {
				if (_243_v105).Contains(_242_v104) {
					return (_243_v105).Get(_242_v104).(_dafny.Sequence)
				}
				return _157_v35
			})(), 18)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _157_v35), 19)
			(_nw46).ArraySet1(_157_v35, 20)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Update(_157_v35, (Companion_Default___.SafeIndex(_101_v3, _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32(), _101_v3), 21)
			(_nw46).ArraySet1(_157_v35, 22)
			(_nw46).ArraySet1(_157_v35, 23)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _157_v35), 24)
			(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v35, _dafny.Companion_Sequence_.Update(_157_v35, (Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_244_v107).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _253_v106 _dafny.Int
					_253_v106 = interface{}(_compr_8).(_dafny.Int)
					if (_244_v107).Contains(_253_v106) {
						_coll8.Add((_253_v106).Minus((_164_v42).F31()), _dafny.IntOfInt64(166))
					}
				}
				return _coll8.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32(), (_164_v42).F31())), 25)
			(_nw46).ArraySet1(_157_v35, 26)
			(_nw46).ArraySet1((_251_v112).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_245_v108).F26()), _dafny.IntOfUint32((_251_v112).Cardinality()))).Uint32()).(_dafny.Sequence), 27)
			_252_v113 = _nw46
			var _254_v114 _dafny.Sequence
			_ = _254_v114
			_254_v114 = _dafny.SeqOf((func() _dafny.Array {
				if (_163_v41).F30() {
					return _252_v113
				}
				return _252_v113
			})(), _252_v113)
			_252_v113 = (_254_v114).Select((Companion_Default___.SafeIndex((_245_v108).F26(), _dafny.IntOfUint32((_254_v114).Cardinality()))).Uint32()).(_dafny.Array)
		}
		var _255_v115 T1
		_ = _255_v115
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("gkpbjugb"), (_163_v41).F30(), (_164_v42).F31(), _106_v7, ((_163_v41).F30()) || (!((_163_v41).Fm7(_105_globalState))))
		_255_v115 = _nw47
		_255_v115 = _255_v115
		var _256_v116 _dafny.Map
		_ = _256_v116
		_256_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true) || (_100_v2), (_163_v41).F30())
		var _257_v117 _dafny.MultiSet
		_ = _257_v117
		_257_v117 = _dafny.MultiSetOf(_100_v2, (_163_v41).F30(), _100_v2, (_163_v41).F30())
		_256_v116 = (_256_v116).Update((Companion_Default___.Fm27(_101_v3, _105_globalState)).IsProperSubsetOf(_257_v117), _100_v2)
		(_105_globalState).F1 = (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)
	}
	var _258_v118 _dafny.Map
	_ = _258_v118
	_258_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _158_v36)
	if _dafny.Companion_Sequence_.Contains(_156_v34, !(_258_v118).Equals(_258_v118)) {
		var _259_v119 _dafny.Set
		_ = _259_v119
		_259_v119 = _dafny.SetOf((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_98_v0), 0))
		_ = _index31
		(_98_v0).ArraySet1(_259_v119, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_98_v0), 0))
		_ = _index32
		(_98_v0).ArraySet1(_259_v119, (_index32).Int())
		var _260_v120 D1
		_ = _260_v120
		_260_v120 = Companion_D1_.Create_DC4_(_100_v2, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
		var _261_v121 D1
		_ = _261_v121
		_261_v121 = Companion_D1_.Create_DC5_(_260_v120)
		var _262_v122 D0
		_ = _262_v122
		_262_v122 = Companion_D0_.Create_DC1_(_101_v3, (_163_v41).F29(), (_163_v41).F30(), (_164_v42).F31(), (_163_v41).F30())
		var _263_v123 bool
		_ = _263_v123
		var _264_v124 bool
		_ = _264_v124
		var _out8 bool
		_ = _out8
		var _out9 bool
		_ = _out9
		_out8, _out9 = (_164_v42).M5(_100_v2, _261_v121, Companion_Default___.Fm4(_262_v122, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _dafny.IntOfUint32(((_163_v41).F29()).Cardinality()), _105_globalState), _105_globalState)
		_263_v123 = _out8
		_264_v124 = _out9
		_158_v36 = _158_v36
		_101_v3 = (((_dafny.Zero).Minus(_101_v3)).Minus(_dafny.IntOfInt64(554))).Minus((_dafny.IntOfUint32((_157_v35).Cardinality())).Times(_101_v3))
		var _265_v125 _dafny.Array
		_ = _265_v125
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw48
		_265_v125 = _nw48
		var _266_v126 _dafny.Map
		_ = _266_v126
		_266_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(172), _156_v34)
		var _267_v127 *C4
		_ = _267_v127
		var _nw49 *C4 = New_C4_()
		_ = _nw49
		_nw49.Ctor__(_dafny.IntOfInt64(-205), (_163_v41).F29(), (_163_v41).F30(), (_259_v119).Cardinality())
		_267_v127 = _nw49
		var _268_v128 _dafny.Set
		_ = _268_v128
		_268_v128 = _dafny.SetOf(_267_v127, _267_v127)
		var _269_v129 _dafny.Array
		_ = _269_v129
		var _nwElement0_10 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eblm")).Cardinality())
		_ = _nwElement0_10
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(25))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_10, 0)
		(_nw50).ArraySet1((_164_v42).F31(), 1)
		(_nw50).ArraySet1((_164_v42).F31(), 2)
		(_nw50).ArraySet1((_164_v42).F31(), 3)
		(_nw50).ArraySet1(_101_v3, 4)
		(_nw50).ArraySet1(_dafny.IntOfInt64(-139), 5)
		(_nw50).ArraySet1(_dafny.IntOfInt64(420), 6)
		(_nw50).ArraySet1(_101_v3, 7)
		(_nw50).ArraySet1((_164_v42).F31(), 8)
		(_nw50).ArraySet1(_101_v3, 9)
		(_nw50).ArraySet1((_164_v42).F31(), 10)
		(_nw50).ArraySet1((_266_v126).Cardinality(), 11)
		(_nw50).ArraySet1(_dafny.IntOfInt64(952), 12)
		(_nw50).ArraySet1(((_98_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_98_v0), 0))).Int()).(_dafny.Set)).Cardinality(), 13)
		(_nw50).ArraySet1(_101_v3, 14)
		(_nw50).ArraySet1((_164_v42).F31(), 15)
		(_nw50).ArraySet1((_164_v42).F31(), 16)
		(_nw50).ArraySet1((_164_v42).F31(), 17)
		(_nw50).ArraySet1((_164_v42).F31(), 18)
		(_nw50).ArraySet1(_101_v3, 19)
		(_nw50).ArraySet1(_dafny.IntOfInt64(576), 20)
		(_nw50).ArraySet1((_164_v42).F31(), 21)
		(_nw50).ArraySet1((_268_v128).Cardinality(), 22)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_263_v123, _263_v123, false)).Cardinality()), 23)
		(_nw50).ArraySet1((_267_v127).F27(), 24)
		_269_v129 = _nw50
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_265_v125), 0))
		_ = _index33
		(_265_v125).ArraySet1(_269_v129, (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_265_v125), 0))
		_ = _index34
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
		_ = _index35
		var _rhs36 _dafny.Array = _269_v129
		_ = _rhs36
		var _rhs37 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(89), (_dafny.Zero).Minus(((_164_v42).F31()).Times((_164_v42).F31())))
		_ = _rhs37
		var _rhs38 bool = (_267_v127).Fm7(_105_globalState)
		_ = _rhs38
		var _lhs31 _dafny.Array = _265_v125
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_265_v125), 0))
		_ = _lhs32
		var _lhs33 *GlobalState = _105_globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = _106_v7
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
		_ = _lhs35
		(_lhs31).ArraySet1(_rhs36, (_lhs32).Int())
		_lhs33.F0 = _rhs37
		(_lhs34).ArraySet1(_rhs38, (_lhs35).Int())
	} else {
		var _270_v130 _dafny.Sequence
		_ = _270_v130
		_270_v130 = _dafny.SeqOf(_163_v41, _163_v41)
		var _271_v131 bool
		_ = _271_v131
		var _out10 bool
		_ = _out10
		_out10 = (_164_v42).M6(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v42).F31(), (_164_v42).F31()), _dafny.IntOfUint32((_270_v130).Cardinality()), _dafny.IntOfInt64(257), _105_globalState)
		_271_v131 = _out10
		if !((func() bool {
			if (_163_v41).F30() {
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_99_v1, _99_v1)
			}
			return (_163_v41).F30()
		})()) {
			var _272_v132 _dafny.Array
			_ = _272_v132
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw51
			_272_v132 = _nw51
			var _273_v133 _dafny.Array
			_ = _273_v133
			var _len0_6 _dafny.Int = _dafny.One
			_ = _len0_6
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_274_v42 *C2) func(_dafny.Int) _dafny.Int {
					return func(_275_i15 _dafny.Int) _dafny.Int {
						return (_275_i15).Plus((_274_v42).F31())
					}
				})(_164_v42)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw52 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw52).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw52).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_273_v133 = _nw52
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_272_v132), 0))
			_ = _index36
			(_272_v132).ArraySet1(_273_v133, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_272_v132), 0))
			_ = _index37
			(_272_v132).ArraySet1(_273_v133, (_index37).Int())
			var _276_v134 T0
			_ = _276_v134
			var _nw53 *C4 = New_C4_()
			_ = _nw53
			_nw53.Ctor__((_dafny.Zero).Minus((_164_v42).F31()), (_163_v41).F29(), (_163_v41).F30(), (_164_v42).F31())
			_276_v134 = _nw53
			var _277_v135 _dafny.Set
			_ = _277_v135
			_277_v135 = _dafny.SetOf(_276_v134)
			var _278_v136 _dafny.Sequence
			_ = _278_v136
			_278_v136 = _dafny.SeqOf(_277_v135)
			var _279_v137 D20
			_ = _279_v137
			_279_v137 = Companion_D20_.Create_DC46_(true, (_164_v42).F31())
			var _280_v138 bool
			_ = _280_v138
			var _out11 bool
			_ = _out11
			_out11 = (_164_v42).M6(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm35((_dafny.Zero).Minus(_101_v3), (_dafny.MultiSetFromSeq(_278_v136)).Cardinality(), _105_globalState)).Cardinality(), (_279_v137).Dtor_cf76()), _101_v3, (_276_v134).Fm6(_100_v2, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _101_v3, _105_globalState), _105_globalState)
			_280_v138 = _out11
			var _281_v139 D2
			_ = _281_v139
			_281_v139 = Companion_D2_.Create_DC6_(_104_v6)
			var _282_v140 _dafny.Sequence
			_ = _282_v140
			_282_v140 = _dafny.SeqOf(Companion_D2_.Create_DC6_(_104_v6), _281_v139, _281_v139)
			_282_v140 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_282_v140, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer19 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_283_v139 D2) func(_dafny.Int) D2 {
				return func(_284_i16 _dafny.Int) D2 {
					return _283_v139
				}
			})(_281_v139)))), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_163_v41).F30() {
					return _101_v3
				}
				return (_164_v42).F31()
			})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_282_v140, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer20 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_285_v139 D2) func(_dafny.Int) D2 {
				return func(_286_i16 _dafny.Int) D2 {
					return _285_v139
				}
			})(_281_v139))))).Cardinality()))).Uint32(), Companion_D2_.Create_DC6_(_104_v6))
			var _287_v141 _dafny.MultiSet
			_ = _287_v141
			_287_v141 = _dafny.MultiSetOf((_164_v42).F31())
			var _288_v142 bool
			_ = _288_v142
			var _out12 bool
			_ = _out12
			_out12 = (_163_v41).M1(_287_v141, (_163_v41).F30(), _100_v2, _105_globalState)
			_288_v142 = _out12
			var _pat_let_tv7 = _164_v42
			_ = _pat_let_tv7
			var _pat_let_tv8 = _164_v42
			_ = _pat_let_tv8
			_101_v3 = (func(_pat_let7_0 D6) D6 {
				return func(_289_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let8_0 _dafny.Int) D6 {
						return func(_290_dt__update_hcf24_h0 _dafny.Int) D6 {
							return func(_pat_let9_0 _dafny.Int) D6 {
								return func(_291_dt__update_hcf23_h0 _dafny.Int) D6 {
									return Companion_D6_.Create_DC14_(_291_dt__update_hcf23_h0, _290_dt__update_hcf24_h0, (_289_dt__update__tmp_h0).Dtor_cf25())
								}(_pat_let9_0)
							}((_pat_let_tv8).F31())
						}(_pat_let8_0)
					}((_pat_let_tv7).F31())
				}(_pat_let7_0)
			}(Companion_Default___.Fm41(_288_v142, true, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), _105_globalState))).Dtor_cf23()
		} else {
			_104_v6 = (Companion_D2_.Create_DC6_(_104_v6)).Dtor_cf10()
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index38
			(_106_v7).ArraySet1(_100_v2, (_index38).Int())
			var _292_v143 *C1
			_ = _292_v143
			var _nw54 *C1 = New_C1_()
			_ = _nw54
			_nw54.Ctor__((_163_v41).F30())
			_292_v143 = _nw54
			var _293_v144 D10
			_ = _293_v144
			_293_v144 = Companion_D10_.Create_DC24_(_dafny.IntOfInt64(547), _292_v143, _158_v36)
			var _294_v145 _dafny.MultiSet
			_ = _294_v145
			_294_v145 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_163_v41).F29()).Cardinality()))
			var _295_v146 _dafny.MultiSet
			_ = _295_v146
			_295_v146 = _dafny.MultiSetOf(_158_v36)
			var _296_v147 _dafny.Sequence
			_ = _296_v147
			_296_v147 = _dafny.SeqOf(_164_v42)
			var _297_v148 D19
			_ = _297_v148
			_297_v148 = Companion_D19_.Create_DC43_((_296_v147).Select((Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_296_v147).Cardinality()))).Uint32()).(*C2), (_164_v42).F31(), (_292_v143).F32())
			var _298_v149 _dafny.MultiSet
			_ = _298_v149
			_298_v149 = _dafny.MultiSetOf(true, _100_v2)
			var _299_v150 D6
			_ = _299_v150
			_299_v150 = Companion_D6_.Create_DC13_((_164_v42).F31(), (_164_v42).F31(), _dafny.SetOf((_298_v149).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ygqeqe")).Cardinality()), (_164_v42).F31()))
			var _300_v151 _dafny.Array
			_ = _300_v151
			var _nwElement0_11 _dafny.Int = (_163_v41).Fm6((_163_v41).F30(), _100_v2, (_164_v42).F31(), _105_globalState)
			_ = _nwElement0_11
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(10))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_11, 0)
			(_nw55).ArraySet1(_dafny.IntOfInt64(969), 1)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_99_v1).Cardinality()), 2)
			(_nw55).ArraySet1((_164_v42).F31(), 3)
			(_nw55).ArraySet1((_164_v42).F31(), 4)
			(_nw55).ArraySet1((_299_v150).Dtor_cf20(), 5)
			(_nw55).ArraySet1(Companion_Default___.Fm2((_164_v42).F31(), _dafny.IntOfInt64(421), _105_globalState), 6)
			(_nw55).ArraySet1((_164_v42).F31(), 7)
			(_nw55).ArraySet1((_164_v42).F31(), 8)
			(_nw55).ArraySet1((_164_v42).F31(), 9)
			_300_v151 = _nw55
			var _301_v152 _dafny.Set
			_ = _301_v152
			_301_v152 = _dafny.SetOf(_300_v151)
			var _302_v153 T2
			_ = _302_v153
			var _nw56 *C0 = New_C0_()
			_ = _nw56
			_nw56.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("yjpy"), false, _101_v3, _106_v7, (_292_v143).F32())
			_302_v153 = _nw56
			var _303_v154 _dafny.Array
			_ = _303_v154
			var _nwElement0_12 _dafny.Int = (_164_v42).F31()
			_ = _nwElement0_12
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(24))
			_ = _nw57
			(_nw57).ArraySet1(_nwElement0_12, 0)
			(_nw57).ArraySet1((_164_v42).F31(), 1)
			(_nw57).ArraySet1((_164_v42).F31(), 2)
			(_nw57).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(776))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_304_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_305_i17 _dafny.Int) _dafny.CodePoint {
					return _304_v36
				}
			})(_158_v36)))).Cardinality()), 3)
			(_nw57).ArraySet1((_164_v42).F31(), 4)
			(_nw57).ArraySet1((_164_v42).F31(), 5)
			(_nw57).ArraySet1((_293_v144).Dtor_cf36(), 6)
			(_nw57).ArraySet1(_dafny.IntOfInt64(187), 7)
			(_nw57).ArraySet1((_dafny.Zero).Minus((_164_v42).F31()), 8)
			(_nw57).ArraySet1(Companion_Default___.SafeModuloInt((_164_v42).F31(), (_164_v42).F31()), 9)
			(_nw57).ArraySet1((_164_v42).F31(), 10)
			(_nw57).ArraySet1(_101_v3, 11)
			(_nw57).ArraySet1(Companion_Default___.SafeDivisionInt((_164_v42).F31(), (_164_v42).F31()), 12)
			(_nw57).ArraySet1((_164_v42).F31(), 13)
			(_nw57).ArraySet1(_101_v3, 14)
			(_nw57).ArraySet1((_dafny.Zero).Minus((_294_v145).Cardinality()), 15)
			(_nw57).ArraySet1((func() _dafny.Int {
				if false {
					return _101_v3
				}
				return _dafny.IntOfUint32((_156_v34).Cardinality())
			})(), 16)
			(_nw57).ArraySet1((func() _dafny.Int {
				if (_292_v143).F32() {
					return (_295_v146).Cardinality()
				}
				return (_164_v42).F31()
			})(), 17)
			(_nw57).ArraySet1(_dafny.IntOfInt64(979), 18)
			(_nw57).ArraySet1((_164_v42).F31(), 19)
			(_nw57).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_292_v143).F32() {
					return _101_v3
				}
				return (_164_v42).F31()
			})()), 20)
			(_nw57).ArraySet1((_297_v148).Dtor_cf71(), 21)
			(_nw57).ArraySet1(((_301_v152).Union(_301_v152)).Cardinality(), 22)
			(_nw57).ArraySet1((_dafny.SetOf(_302_v153, _302_v153)).Cardinality(), 23)
			_303_v154 = _nw57
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))
			_ = _index39
			(_300_v151).ArraySet1((_164_v42).F31(), (_index39).Int())
			var _306_v155 *C3
			_ = _306_v155
			var _nw58 *C3 = New_C3_()
			_ = _nw58
			_nw58.Ctor__(Companion_Default___.Fm1(_105_globalState), _271_v131, _271_v131, _dafny.IntOfInt64(56))
			_306_v155 = _nw58
			var _307_v156 _dafny.Map
			_ = _307_v156
			_307_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v3, (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))
			_ = _index40
			var _rhs39 _dafny.Int = ((_307_v156).Cardinality()).Minus((_164_v42).F31())
			_ = _rhs39
			var _rhs40 *C3 = _306_v155
			_ = _rhs40
			var _rhs41 bool = (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)
			_ = _rhs41
			var _lhs36 _dafny.Array = _300_v151
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))
			_ = _lhs37
			var _lhs38 *GlobalState = _105_globalState
			_ = _lhs38
			(_lhs36).ArraySet1(_rhs39, (_lhs37).Int())
			_306_v155 = _rhs40
			_lhs38.F1 = _rhs41
			var _308_v157 _dafny.Set
			_ = _308_v157
			_308_v157 = _dafny.SetOf((_292_v143).F32(), (_292_v143).F32())
			var _309_v158 *C2
			_ = _309_v158
			var _nw59 *C2 = New_C2_()
			_ = _nw59
			_nw59.Ctor__((_300_v151).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))).Int()).(_dafny.Int), (_292_v143).Fm14(_308_v157, (_300_v151).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))).Int()).(_dafny.Int), _105_globalState), (_300_v151).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_300_v151), 0))).Int()).(_dafny.Int))
			_309_v158 = _nw59
			_102_v4 = _102_v4
		}
		var _310_v159 _dafny.Array
		_ = _310_v159
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_7
		var _nw60 _dafny.Array
		_ = _nw60
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw60 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = (func(_311_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_312_i18 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-591))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_313_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_314_i19 _dafny.Int) _dafny.CodePoint {
							return _313_v36
						}
					})(_311_v36)))
				}
			})(_158_v36)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw60 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw60).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw60).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_310_v159 = _nw60
		var _315_v160 _dafny.Sequence
		_ = _315_v160
		_315_v160 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-774))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_316_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_317_i20 _dafny.Int) _dafny.CodePoint {
				return _316_v36
			}
		})(_158_v36))))
		var _318_v161 _dafny.MultiSet
		_ = _318_v161
		_318_v161 = _dafny.MultiSetOf((_164_v42).F31())
		var _319_v162 T1
		_ = _319_v162
		var _nw61 *C0 = New_C0_()
		_ = _nw61
		_nw61.Ctor__(_99_v1, _271_v131, (_318_v161).Cardinality(), _106_v7, (_163_v41).F30())
		_319_v162 = _nw61
		var _320_v163 _dafny.Sequence
		_ = _320_v163
		_320_v163 = _dafny.SeqOf(_319_v162)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_310_v159), 0))
		_ = _index41
		(_310_v159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_99_v1, (_315_v160).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_320_v163).Cardinality()), _dafny.IntOfUint32((_315_v160).Cardinality()))).Uint32()).(_dafny.Sequence)), (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_310_v159), 0))
		_ = _index42
		(_310_v159).ArraySet1((_163_v41).F29(), (_index42).Int())
		_101_v3 = (_164_v42).F31()
		var _321_v164 _dafny.Map
		_ = _321_v164
		_321_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vrd"), _100_v2)
		_321_v164 = (_321_v164).Update(_dafny.UnicodeSeqOfUtf8Bytes("tdm"), _271_v131)
	}
	var _322_v165 _dafny.Set
	_ = _322_v165
	_322_v165 = _dafny.SetOf((_163_v41).F30())
	if (_322_v165).IsSubsetOf(_322_v165) {
		var _323_v166 _dafny.Array
		_ = _323_v166
		var _nwElement0_13 _dafny.Int = Companion_Default___.SafeModuloInt((_164_v42).F31(), (_164_v42).F31())
		_ = _nwElement0_13
		var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
		_ = _nw62
		(_nw62).ArraySet1(_nwElement0_13, 0)
		_323_v166 = _nw62
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
		_ = _index43
		(_323_v166).ArraySet1(_dafny.IntOfInt64(-575), (_index43).Int())
		var _324_v167 _dafny.MultiSet
		_ = _324_v167
		_324_v167 = _dafny.MultiSetOf(_101_v3)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
		_ = _index44
		(_323_v166).ArraySet1((_dafny.IntOfInt64(-461)).Minus((_324_v167).Cardinality()), (_index44).Int())
		var _325_v168 _dafny.Sequence
		_ = _325_v168
		_325_v168 = _dafny.SeqOf(_104_v6, _104_v6, _104_v6, _104_v6, _104_v6)
		var _326_v169 _dafny.Map
		_ = _326_v169
		_326_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(80), _100_v2)
		var _327_v170 D12
		_ = _327_v170
		_327_v170 = Companion_D12_.Create_DC29_(_dafny.IntOfInt64(381), _100_v2, _325_v168, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v2, (_326_v169).Cardinality())).Cardinality(), _dafny.IntOfInt64(-170))
		var _source3 D12 = (func() D12 {
			if _100_v2 {
				return _327_v170
			}
			return _327_v170
		})()
		_ = _source3
		if _source3.Is_DC29() {
			var _328___mcc_h6 _dafny.Int = _source3.Get_().(D12_DC29).Cf45
			_ = _328___mcc_h6
			var _329___mcc_h7 bool = _source3.Get_().(D12_DC29).Cf46
			_ = _329___mcc_h7
			var _330___mcc_h8 _dafny.Sequence = _source3.Get_().(D12_DC29).Cf47
			_ = _330___mcc_h8
			var _331___mcc_h9 _dafny.Int = _source3.Get_().(D12_DC29).Cf48
			_ = _331___mcc_h9
			var _332___mcc_h10 _dafny.Int = _source3.Get_().(D12_DC29).Cf49
			_ = _332___mcc_h10
			var _333_cf49 _dafny.Int = _332___mcc_h10
			_ = _333_cf49
			var _334_cf48 _dafny.Int = _331___mcc_h9
			_ = _334_cf48
			var _335_cf47 _dafny.Sequence = _330___mcc_h8
			_ = _335_cf47
			var _336_cf46 bool = _329___mcc_h7
			_ = _336_cf46
			var _337_cf45 _dafny.Int = _328___mcc_h6
			_ = _337_cf45
			var _338_v171 _dafny.Sequence
			_ = _338_v171
			_338_v171 = _dafny.SeqOf(_156_v34, _156_v34, _156_v34)
			_156_v34 = _dafny.Companion_Sequence_.Concatenate(_156_v34, (_338_v171).Select((Companion_Default___.SafeIndex(_337_cf45, _dafny.IntOfUint32((_338_v171).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _339_v172 *C3
			_ = _339_v172
			var _nw63 *C3 = New_C3_()
			_ = _nw63
			_nw63.Ctor__((_163_v41).F29(), _336_cf46, (_163_v41).F30(), _337_cf45)
			_339_v172 = _nw63
			_99_v1 = _dafny.UnicodeSeqOfUtf8Bytes("kh")
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index45
			(_106_v7).ArraySet1((_163_v41).F30(), (_index45).Int())
		} else {
			var _340___mcc_h11 _dafny.Array = _source3.Get_().(D12_DC28).Cf44
			_ = _340___mcc_h11
			var _341_cf44 _dafny.Array = _340___mcc_h11
			_ = _341_cf44
			var _342_v173 _dafny.MultiSet
			_ = _342_v173
			_342_v173 = _dafny.MultiSetOf(!((_163_v41).F30()))
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index46
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index47
			var _rhs42 bool = (_100_v2) && ((_163_v41).F30())
			_ = _rhs42
			var _rhs43 _dafny.Int = Companion_Default___.SafeModuloInt((_164_v42).Fm6(true, true, (_dafny.Zero).Minus((_164_v42).F31()), _105_globalState), (_157_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_99_v1).Cardinality()), _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs43
			var _rhs44 bool = !(_100_v2)
			_ = _rhs44
			var _rhs45 bool = (_342_v173).IsDisjointFrom((_342_v173).Intersection(_342_v173))
			_ = _rhs45
			var _rhs46 _dafny.Array = _166_v44
			_ = _rhs46
			var _lhs39 _dafny.Array = _106_v7
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _lhs40
			var _lhs41 *GlobalState = _105_globalState
			_ = _lhs41
			var _lhs42 _dafny.Array = _106_v7
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _lhs43
			var _lhs44 *GlobalState = _105_globalState
			_ = _lhs44
			(_lhs39).ArraySet1(_rhs42, (_lhs40).Int())
			_lhs41.F0 = _rhs43
			(_lhs42).ArraySet1(_rhs44, (_lhs43).Int())
			_lhs44.F1 = _rhs45
			_166_v44 = _rhs46
			var _343_v174 _dafny.Array
			_ = _343_v174
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw64
			_343_v174 = _nw64
			var _344_v175 _dafny.Map
			_ = _344_v175
			_344_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v7, _343_v174)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _index49
			var _rhs47 _dafny.Int = (_164_v42).F31()
			_ = _rhs47
			var _rhs48 bool = ((_164_v42).F31()).Cmp(_101_v3) <= 0
			_ = _rhs48
			var _rhs49 _dafny.Array = (func() _dafny.Array {
				if (_344_v175).Contains(_106_v7) {
					return (_344_v175).Get(_106_v7).(_dafny.Array)
				}
				return _343_v174
			})()
			_ = _rhs49
			var _rhs50 _dafny.MultiSet = Companion_Default___.Fm21(Companion_Default___.Fm1(_105_globalState), (_164_v42).F31(), _dafny.IntOfInt64(82), _105_globalState)
			_ = _rhs50
			var _rhs51 bool = (_163_v41).F30()
			_ = _rhs51
			var _lhs45 _dafny.Array = _323_v166
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _lhs46
			var _lhs47 _dafny.Array = _106_v7
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
			_ = _lhs48
			(_lhs45).ArraySet1(_rhs47, (_lhs46).Int())
			(_lhs47).ArraySet1(_rhs48, (_lhs48).Int())
			_343_v174 = _rhs49
			_324_v167 = _rhs50
			_100_v2 = _rhs51
			var _345_v176 bool
			_ = _345_v176
			var _out13 bool
			_ = _out13
			_out13 = (_164_v42).M1(_324_v167, ((_342_v173).Cardinality()).Cmp((_164_v42).F31()) >= 0, (_163_v41).F30(), _105_globalState)
			_345_v176 = _out13
			_100_v2 = (_163_v41).F30()
		}
		var _346_v177 *C1
		_ = _346_v177
		var _nw65 *C1 = New_C1_()
		_ = _nw65
		_nw65.Ctor__(false)
		_346_v177 = _nw65
		var _347_v178 _dafny.Map
		_ = _347_v178
		_347_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v177, (func() *C3 {
			if (_346_v177).F32() {
				return _163_v41
			}
			return _163_v41
		})())
		var _348_v179 _dafny.Sequence
		_ = _348_v179
		_348_v179 = _dafny.SeqOf(_347_v178)
		var _349_v180 T2
		_ = _349_v180
		var _nw66 *C0 = New_C0_()
		_ = _nw66
		_nw66.Ctor__((_163_v41).F29(), (_163_v41).F30(), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_163_v41).F29()).Cardinality())), _106_v7, (_163_v41).F30())
		_349_v180 = _nw66
		var _rhs52 _dafny.Map = (_348_v179).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_349_v180, _349_v180, _349_v180, _349_v180)).Cardinality()), _dafny.IntOfUint32((_348_v179).Cardinality()))).Uint32()).(_dafny.Map)
		_ = _rhs52
		var _rhs53 _dafny.Array = _106_v7
		_ = _rhs53
		var _rhs54 _dafny.Int = (_164_v42).Fm6(!(((_164_v42).F31()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jhjne")).Cardinality())) <= 0), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_324_v167).Contains((_164_v42).F31()) {
				return (_324_v167).Multiplicity((_164_v42).F31())
			}
			return (_323_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))).Int()).(_dafny.Int)
		})(), (_323_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))).Int()).(_dafny.Int)), _105_globalState)
		_ = _rhs54
		var _rhs55 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(283), _dafny.IntOfInt64(161), _dafny.IntOfInt64(353))).Cardinality()), (_164_v42).F31())
		_ = _rhs55
		var _lhs49 *GlobalState = _105_globalState
		_ = _lhs49
		var _lhs50 *GlobalState = _105_globalState
		_ = _lhs50
		_347_v178 = _rhs52
		_106_v7 = _rhs53
		_lhs49.F2 = _rhs54
		_lhs50.F0 = _rhs55
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
		_ = _index50
		var _rhs56 bool = !(_dafny.MultiSetOf(_349_v180.F34(), (_163_v41).F30())).Contains(_100_v2)
		_ = _rhs56
		var _rhs57 bool = (_346_v177).F32()
		_ = _rhs57
		var _rhs58 _dafny.Array = _323_v166
		_ = _rhs58
		var _rhs59 _dafny.Int = (_164_v42).F31()
		_ = _rhs59
		var _rhs60 _dafny.Sequence = (_163_v41).F29()
		_ = _rhs60
		var _lhs51 T2 = _349_v180
		_ = _lhs51
		var _lhs52 _dafny.Array = _106_v7
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
		_ = _lhs53
		_lhs51.F34_set_(_rhs56)
		(_lhs52).ArraySet1(_rhs57, (_lhs53).Int())
		_323_v166 = _rhs58
		_101_v3 = _rhs59
		_99_v1 = _rhs60
		var _350_v181 _dafny.MultiSet
		_ = _350_v181
		_350_v181 = _dafny.MultiSetOf((_163_v41).F30(), _100_v2)
		if !(((func() _dafny.Int {
			if (_350_v181).Contains(_100_v2) {
				return (_350_v181).Multiplicity(_100_v2)
			}
			return _101_v3
		})()).Cmp((_164_v42).F31()) == 0) {
			var _351_v182 *C4
			_ = _351_v182
			var _nw67 *C4 = New_C4_()
			_ = _nw67
			_nw67.Ctor__((_164_v42).F31(), _99_v1, (_346_v177).F32(), (_164_v42).F31())
			_351_v182 = _nw67
			var _352_v183 _dafny.Map
			_ = _352_v183
			_352_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_351_v182, (_164_v42).F31())
			var _353_v184 _dafny.Map
			_ = _353_v184
			_353_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v41).F30(), (func() bool {
				if (_326_v169).Contains((_351_v182).F27()) {
					return (_326_v169).Get((_351_v182).F27()).(bool)
				}
				return (_163_v41).F30()
			})())
			var _354_v185 D14
			_ = _354_v185
			_354_v185 = Companion_D14_.Create_DC32_((_163_v41).F30(), _101_v3, _353_v184, (_163_v41).F29(), (_346_v177).F32())
			var _355_v186 *C4
			_ = _355_v186
			var _nw68 *C4 = New_C4_()
			_ = _nw68
			_nw68.Ctor__((func() _dafny.Int {
				if (_352_v183).Contains(_351_v182) {
					return (_352_v183).Get(_351_v182).(_dafny.Int)
				}
				return (_164_v42).F31()
			})(), (_354_v185).Dtor_cf55(), (_346_v177).F32(), (_102_v4).Cardinality())
			_355_v186 = _nw68
			var _356_v187 _dafny.Array
			_ = _356_v187
			var _nwElement0_14 _dafny.Array = _323_v166
			_ = _nwElement0_14
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(27))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_14, 0)
			(_nw69).ArraySet1(_323_v166, 1)
			(_nw69).ArraySet1(_323_v166, 2)
			(_nw69).ArraySet1((func() _dafny.Array {
				if (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool) {
					return _323_v166
				}
				return _323_v166
			})(), 3)
			(_nw69).ArraySet1(_323_v166, 4)
			(_nw69).ArraySet1(_323_v166, 5)
			(_nw69).ArraySet1((func() _dafny.Array {
				if true {
					return _323_v166
				}
				return _323_v166
			})(), 6)
			(_nw69).ArraySet1(_323_v166, 7)
			(_nw69).ArraySet1(_323_v166, 8)
			(_nw69).ArraySet1(_323_v166, 9)
			(_nw69).ArraySet1(_323_v166, 10)
			(_nw69).ArraySet1(_323_v166, 11)
			(_nw69).ArraySet1(_323_v166, 12)
			(_nw69).ArraySet1(_323_v166, 13)
			(_nw69).ArraySet1(_323_v166, 14)
			(_nw69).ArraySet1(_323_v166, 15)
			(_nw69).ArraySet1(_323_v166, 16)
			(_nw69).ArraySet1(_323_v166, 17)
			(_nw69).ArraySet1(_323_v166, 18)
			(_nw69).ArraySet1(_323_v166, 19)
			(_nw69).ArraySet1(_323_v166, 20)
			(_nw69).ArraySet1(_323_v166, 21)
			(_nw69).ArraySet1(_323_v166, 22)
			(_nw69).ArraySet1(_323_v166, 23)
			(_nw69).ArraySet1(_323_v166, 24)
			(_nw69).ArraySet1(_323_v166, 25)
			(_nw69).ArraySet1(_323_v166, 26)
			_356_v187 = _nw69
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _index51
			var _rhs61 bool = (_156_v34).Select((Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_156_v34).Cardinality()))).Uint32()).(bool)
			_ = _rhs61
			var _rhs62 _dafny.Int = _101_v3
			_ = _rhs62
			var _rhs63 _dafny.Array = _356_v187
			_ = _rhs63
			var _rhs64 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_157_v35, (Companion_Default___.SafeIndex((_323_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32(), (_164_v42).F31())).Cardinality())
			_ = _rhs64
			var _lhs54 T2 = _349_v180
			_ = _lhs54
			var _lhs55 _dafny.Array = _323_v166
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _lhs56
			var _lhs57 *GlobalState = _105_globalState
			_ = _lhs57
			_lhs54.F34_set_(_rhs61)
			(_lhs55).ArraySet1(_rhs62, (_lhs56).Int())
			_356_v187 = _rhs63
			_lhs57.F0 = _rhs64
			_102_v4 = (_102_v4).Update((_163_v41).F30(), _dafny.IntOfInt64(-43))
			(_105_globalState).F2 = (_164_v42).F31()
			_104_v6 = _104_v6
		} else {
			var _357_v188 _dafny.Map
			_ = _357_v188
			_357_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v36, _106_v7)
			var _358_v189 _dafny.MultiSet
			_ = _358_v189
			_358_v189 = _dafny.MultiSetOf((_357_v188).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v36, (_349_v180).F33())))
			(_105_globalState).F0 = (func() _dafny.Int {
				if (_358_v189).Contains(_357_v188) {
					return (_358_v189).Multiplicity(_357_v188)
				}
				return (_164_v42).F31()
			})()
			_99_v1 = _dafny.Companion_Sequence_.Concatenate((_163_v41).F29(), _dafny.Companion_Sequence_.Concatenate((_163_v41).F29(), (_163_v41).F29()))
			(_105_globalState).F0 = (_164_v42).F31()
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _index52
			(_323_v166).ArraySet1(Companion_Default___.SafeModuloInt(((_164_v42).F31()).Times((_164_v42).F31()), _101_v3), (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))
			_ = _index53
			(_323_v166).ArraySet1((_323_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_323_v166), 0))).Int()).(_dafny.Int), (_index53).Int())
		}
	} else {
		_157_v35 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3((_164_v42).F31(), (_163_v41).F30(), (_164_v42).F31(), _100_v2, _105_globalState), _157_v35), _157_v35)
		var _359_v190 _dafny.Map
		_ = _359_v190
		_359_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_156_v34).Cardinality()), _dafny.IntOfUint32((_156_v34).Cardinality()))).Uint32()).(bool), true)
		var _360_v191 _dafny.MultiSet
		_ = _360_v191
		_360_v191 = _dafny.MultiSetOf(_359_v190, _359_v190, (_359_v190).Merge(_359_v190), (_359_v190).Update(!((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)), true), (_359_v190).Update((_163_v41).F30(), (_163_v41).F30()))
		_360_v191 = ((_dafny.MultiSetOf(_359_v190)).Difference(_dafny.MultiSetOf(_359_v190, _359_v190, _359_v190))).Difference(_dafny.MultiSetOf(_359_v190))
		var _361_v192 _dafny.Sequence
		_ = _361_v192
		_361_v192 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bqwbfped"), (_163_v41).F29(), _99_v1, _99_v1)
		_361_v192 = _dafny.Companion_Sequence_.Concatenate(_361_v192, _361_v192)
		var _362_v193 _dafny.Map
		_ = _362_v193
		_362_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v35, _106_v7)
		_362_v193 = _362_v193
		var _363_v194 _dafny.Array
		_ = _363_v194
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw70
		_363_v194 = _nw70
		var _364_v195 _dafny.Array
		_ = _364_v195
		var _nwElement0_15 _dafny.Int = (_359_v190).Cardinality()
		_ = _nwElement0_15
		var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(2))
		_ = _nw71
		(_nw71).ArraySet1(_nwElement0_15, 0)
		(_nw71).ArraySet1(_101_v3, 1)
		_364_v195 = _nw71
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_363_v194), 0))
		_ = _index54
		(_363_v194).ArraySet1(_364_v195, (_index54).Int())
		var _365_v196 _dafny.Map
		_ = _365_v196
		_365_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_164_v42).F31()).Plus(_101_v3), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))
		var _366_v197 _dafny.Sequence
		_ = _366_v197
		_366_v197 = _dafny.SeqOf(_156_v34)
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_363_v194), 0))
		_ = _index55
		var _rhs65 _dafny.Array = _364_v195
		_ = _rhs65
		var _rhs66 _dafny.Array = (func() _dafny.Array {
			if (func() bool {
				if (_163_v41).F30() {
					return (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool)
				}
				return (_163_v41).F30()
			})() {
				return _364_v195
			}
			return _364_v195
		})()
		_ = _rhs66
		var _rhs67 _dafny.Map = Companion_Default___.Fm42(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3(_dafny.IntOfUint32((_dafny.SeqOf((_164_v42).F31())).Cardinality()), (_163_v41).F30(), (_164_v42).F31(), (_163_v41).F30(), _105_globalState), _dafny.Companion_Sequence_.Update(_157_v35, (Companion_Default___.SafeIndex((_164_v42).F31(), _dafny.IntOfUint32((_157_v35).Cardinality()))).Uint32(), (_164_v42).F31())), _100_v2, _366_v197, _105_globalState)
		_ = _rhs67
		var _lhs58 _dafny.Array = _363_v194
		_ = _lhs58
		var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_363_v194), 0))
		_ = _lhs59
		(_lhs58).ArraySet1(_rhs65, (_lhs59).Int())
		_364_v195 = _rhs66
		_365_v196 = _rhs67
	}
	var _367_v198 _dafny.Map
	_ = _367_v198
	_367_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v1, (_163_v41).F29())
	var _368_v199 _dafny.MultiSet
	_ = _368_v199
	_368_v199 = _dafny.MultiSetOf((_164_v42).Fm11(_367_v198, _dafny.IntOfUint32(((_163_v41).F29()).Cardinality()), (_164_v42).F31(), _105_globalState), _101_v3, (_164_v42).F31(), (_164_v42).F31())
	_368_v199 = (_368_v199).Union(_368_v199)
	var _369_v200 _dafny.Map
	_ = _369_v200
	_369_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(false), _101_v3)
	var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))
	_ = _index56
	(_106_v7).ArraySet1(!(Companion_Default___.Fm5((_dafny.MultiSetOf((_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool), !((_163_v41).F30()), (_163_v41).F30(), (_106_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_106_v7), 0))).Int()).(bool))).IsDisjointFrom(_dafny.MultiSetOf((_163_v41).F30())), _369_v200, _101_v3, _105_globalState)), (_index56).Int())
	var _370_v201 _dafny.Map
	_ = _370_v201
	_370_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v7, _dafny.Companion_Sequence_.Concatenate(_156_v34, _156_v34))
	_370_v201 = (_370_v201).Update(_106_v7, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_163_v41).F30()), _156_v34))
	var _371_v202 _dafny.Array
	_ = _371_v202
	var _len0_8 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_8
	var _nw72 _dafny.Array
	_ = _nw72
	if _len0_8.Cmp(_dafny.Zero) == 0 {
		_nw72 = _dafny.NewArray(_len0_8)
	} else {
		var _init8 func(_dafny.Int) _dafny.Int = (func(_372_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_373_i21 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_373_i21, _372_v3)
			}
		})(_101_v3)
		_ = _init8
		var _element0_8 = _init8(_dafny.Zero)
		_ = _element0_8
		_nw72 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
		(_nw72).ArraySet1(_element0_8, 0)
		var _nativeLen0_8 = (_len0_8).Int()
		_ = _nativeLen0_8
		for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
			(_nw72).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
		}
	}
	_371_v202 = _nw72
	var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_371_v202), 0))
	_ = _index57
	(_371_v202).ArraySet1(_dafny.IntOfInt64(923), (_index57).Int())
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_371_v202), 0))
	_ = _index58
	var _rhs68 _dafny.Int = _101_v3
	_ = _rhs68
	var _rhs69 _dafny.CodePoint = (_99_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_164_v42).F31(), (_164_v42).F31()), _dafny.IntOfUint32((_99_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
	_ = _rhs69
	var _rhs70 bool = (_dafny.IntOfInt64(224)).Cmp(_101_v3) > 0
	_ = _rhs70
	var _rhs71 _dafny.Sequence = _156_v34
	_ = _rhs71
	var _rhs72 _dafny.Int = (_164_v42).F31()
	_ = _rhs72
	var _lhs60 _dafny.Array = _371_v202
	_ = _lhs60
	var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_371_v202), 0))
	_ = _lhs61
	_101_v3 = _rhs68
	_158_v36 = _rhs69
	_100_v2 = _rhs70
	_156_v34 = _rhs71
	(_lhs60).ArraySet1(_rhs72, (_lhs61).Int())
	var _374_v203 _dafny.Array
	_ = _374_v203
	var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
	_ = _nw73
	_374_v203 = _nw73
	var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_374_v203), 0))
	_ = _index59
	(_374_v203).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_163_v41).F29(), _dafny.UnicodeSeqOfUtf8Bytes("cd")), (_index59).Int())
	var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_374_v203), 0))
	_ = _index60
	(_374_v203).ArraySet1(_99_v1, (_index60).Int())
	_dafny.Print(_99_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_100_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_101_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-43))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_103_v5, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(189)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(189)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v6).Equals(_dafny.SetOf(_dafny.IntOfInt64(189))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_105_globalState).F9(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(189)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(189)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F12().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F13()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("jxiyve"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F17()).Equals(_dafny.SetOf(_dafny.IntOfInt64(189))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v7).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_110_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_156_v34, _dafny.SeqOf(true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_157_v35, _dafny.SeqOf(_dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v36)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v38).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_160_v39, _dafny.SeqOf(Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v40).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(true), _dafny.One).UpdateUnsafe(Companion_D0_.Create_DC0_(false), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v41).F29().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v41).F30())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v42).F31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v43).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v45).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v118).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('r'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v165).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_367_v198).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("oduxsblwp"), _dafny.UnicodeSeqOfUtf8Bytes("oduxsblwp"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_368_v199).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(991), _dafny.IntOfInt64(991), _dafny.IntOfInt64(-189), _dafny.IntOfInt64(-189), _dafny.IntOfInt64(-189), _dafny.IntOfInt64(-189), _dafny.IntOfInt64(-189), _dafny.IntOfInt64(-189))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_369_v200).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(false), _dafny.IntOfInt64(-189))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_370_v201).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v202).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_374_v203).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
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
	Cf1 _dafny.Int
	Cf2 _dafny.Sequence
	Cf3 bool
	Cf4 _dafny.Int
	Cf5 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Sequence, Cf3 bool, Cf4 _dafny.Int, Cf5 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_() D0 {
	return D0{D0_DC2{}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.EmptySeq, false, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC1).Cf5
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + data.Cf2.VerbatimString(true) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D0_DC2:
		{
			_, ok := other.Get_().(D0_DC2)
			return ok
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

type D1_DC4 struct {
	Cf7 bool
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 bool, Cf8 bool) D1 {
	return D1{D1_DC4{Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf6}}
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
	return Companion_D1_.Create_DC4_(false, false)
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf6
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
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6 == data2.Cf6
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
	Cf12 bool
	Cf13 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool, Cf12 bool, Cf13 _dafny.Int) D2 {
	return D2{D2_DC7{Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf10 _dafny.Set
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.Set) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, false, _dafny.Zero)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf10() _dafny.Set {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13.Cmp(data2.Cf13) == 0
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

type D3_DC8 struct {
	Cf14 _dafny.MultiSet
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.MultiSet) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D3) Dtor_cf14() _dafny.MultiSet {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D4_DC10 struct {
	Cf16 _dafny.CodePoint
	Cf17 _dafny.CodePoint
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 _dafny.CodePoint, Cf17 _dafny.CodePoint) D4 {
	return D4{D4_DC10{Cf16, Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf15 _dafny.Set
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 _dafny.Set) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.CodePoint('D'), _dafny.CodePoint('D'))
}

func (_this D4) Dtor_cf16() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf15() _dafny.Set {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17 == data2.Cf17
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D5_DC11 struct {
	Cf18 _dafny.MultiSet
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf18 _dafny.MultiSet) D5 {
	return D5{D5_DC11{Cf18}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D5_DC11).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
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

type D6_DC13 struct {
	Cf20 _dafny.Int
	Cf21 _dafny.Int
	Cf22 _dafny.Set
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf20 _dafny.Int, Cf21 _dafny.Int, Cf22 _dafny.Set) D6 {
	return D6{D6_DC13{Cf20, Cf21, Cf22}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC14 struct {
	Cf23 _dafny.Int
	Cf24 _dafny.Int
	Cf25 _dafny.MultiSet
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf23 _dafny.Int, Cf24 _dafny.Int, Cf25 _dafny.MultiSet) D6 {
	return D6{D6_DC14{Cf23, Cf24, Cf25}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC12 struct {
	Cf19 _dafny.Set
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf19 _dafny.Set) D6 {
	return D6{D6_DC12{Cf19}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC13_(_dafny.Zero, _dafny.Zero, _dafny.EmptySet)
}

func (_this D6) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf20
}

func (_this D6) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf21
}

func (_this D6) Dtor_cf22() _dafny.Set {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) Dtor_cf25() _dafny.MultiSet {
	return _this.Get_().(D6_DC14).Cf25
}

func (_this D6) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D6_DC12).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Equals(data2.Cf22)
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25)
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf26 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf26 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf26}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf26) + ")"
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
	Cf28 D0
	Cf29 bool
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf28 D0, Cf29 bool) D8 {
	return D8{D8_DC17{Cf28, Cf29}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC16 struct {
	Cf27 _dafny.Array
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf27 _dafny.Array) D8 {
	return D8{D8_DC16{Cf27}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC17_(Companion_D0_.Default(), false)
}

func (_this D8) Dtor_cf28() D0 {
	return _this.Get_().(D8_DC17).Cf28
}

func (_this D8) Dtor_cf29() bool {
	return _this.Get_().(D8_DC17).Cf29
}

func (_this D8) Dtor_cf27() _dafny.Array {
	return _this.Get_().(D8_DC16).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf28.Equals(data2.Cf28) && data1.Cf29 == data2.Cf29
		}
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf27 == data2.Cf27
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
	Cf31 D6
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf31 D6) D9 {
	return D9{D9_DC19{Cf31}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC18 struct {
	Cf30 _dafny.Map
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf30 _dafny.Map) D9 {
	return D9{D9_DC18{Cf30}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

type D9_DC20 struct {
	Cf32 D9
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf32 D9) D9 {
	return D9{D9_DC20{Cf32}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC19_(Companion_D6_.Default())
}

func (_this D9) Dtor_cf31() D6 {
	return _this.Get_().(D9_DC19).Cf31
}

func (_this D9) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D9_DC18).Cf30
}

func (_this D9) Dtor_cf32() D9 {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf31) + ")"
		}
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf32) + ")"
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
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf32.Equals(data2.Cf32)
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
	Cf34 _dafny.Array
	Cf35 _dafny.Sequence
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf34 _dafny.Array, Cf35 _dafny.Sequence) D10 {
	return D10{D10_DC22{Cf34, Cf35}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC23 struct {
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_() D10 {
	return D10{D10_DC23{}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

type D10_DC24 struct {
	Cf36 _dafny.Int
	Cf37 *C1
	Cf38 _dafny.CodePoint
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf36 _dafny.Int, Cf37 *C1, Cf38 _dafny.CodePoint) D10 {
	return D10{D10_DC24{Cf36, Cf37, Cf38}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

type D10_DC21 struct {
	Cf33 _dafny.Array
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf33 _dafny.Array) D10 {
	return D10{D10_DC21{Cf33}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq)
}

func (_this D10) Dtor_cf34() _dafny.Array {
	return _this.Get_().(D10_DC22).Cf34
}

func (_this D10) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D10_DC22).Cf35
}

func (_this D10) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D10_DC24).Cf36
}

func (_this D10) Dtor_cf37() *C1 {
	return _this.Get_().(D10_DC24).Cf37
}

func (_this D10) Dtor_cf38() _dafny.CodePoint {
	return _this.Get_().(D10_DC24).Cf38
}

func (_this D10) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D10_DC21).Cf33
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf34) + ", " + data.Cf35.VerbatimString(true) + ")"
		}
	case D10_DC23:
		{
			return "D10.DC23"
		}
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35.Equals(data2.Cf35)
		}
	case D10_DC23:
		{
			_, ok := other.Get_().(D10_DC23)
			return ok
		}
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf33 == data2.Cf33
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

type D11_DC26 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Int
	Cf42 bool
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf40 _dafny.Int, Cf41 _dafny.Int, Cf42 bool) D11 {
	return D11{D11_DC26{Cf40, Cf41, Cf42}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC25 struct {
	Cf39 _dafny.Map
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf39 _dafny.Map) D11 {
	return D11{D11_DC25{Cf39}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC27 struct {
	Cf43 D11
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf43 D11) D11 {
	return D11{D11_DC27{Cf43}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC26_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D11) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D11_DC26).Cf40
}

func (_this D11) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D11_DC26).Cf41
}

func (_this D11) Dtor_cf42() bool {
	return _this.Get_().(D11_DC26).Cf42
}

func (_this D11) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D11_DC25).Cf39
}

func (_this D11) Dtor_cf43() D11 {
	return _this.Get_().(D11_DC27).Cf43
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D12_DC29 struct {
	Cf45 _dafny.Int
	Cf46 bool
	Cf47 _dafny.Sequence
	Cf48 _dafny.Int
	Cf49 _dafny.Int
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf45 _dafny.Int, Cf46 bool, Cf47 _dafny.Sequence, Cf48 _dafny.Int, Cf49 _dafny.Int) D12 {
	return D12{D12_DC29{Cf45, Cf46, Cf47, Cf48, Cf49}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC28 struct {
	Cf44 _dafny.Array
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf44 _dafny.Array) D12 {
	return D12{D12_DC28{Cf44}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC29_(_dafny.Zero, false, _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D12) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf45
}

func (_this D12) Dtor_cf46() bool {
	return _this.Get_().(D12_DC29).Cf46
}

func (_this D12) Dtor_cf47() _dafny.Sequence {
	return _this.Get_().(D12_DC29).Cf47
}

func (_this D12) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf48
}

func (_this D12) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf49
}

func (_this D12) Dtor_cf44() _dafny.Array {
	return _this.Get_().(D12_DC28).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47) && data1.Cf48.Cmp(data2.Cf48) == 0 && data1.Cf49.Cmp(data2.Cf49) == 0
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf44 == data2.Cf44
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

type D13_DC30 struct {
	Cf50 _dafny.MultiSet
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf50 _dafny.MultiSet) D13 {
	return D13{D13_DC30{Cf50}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D13) Dtor_cf50() _dafny.MultiSet {
	return _this.Get_().(D13_DC30).Cf50
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D14_DC32 struct {
	Cf52 bool
	Cf53 _dafny.Int
	Cf54 _dafny.Map
	Cf55 _dafny.Sequence
	Cf56 bool
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_(Cf52 bool, Cf53 _dafny.Int, Cf54 _dafny.Map, Cf55 _dafny.Sequence, Cf56 bool) D14 {
	return D14{D14_DC32{Cf52, Cf53, Cf54, Cf55, Cf56}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

type D14_DC31 struct {
	Cf51 _dafny.Map
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf51 _dafny.Map) D14 {
	return D14{D14_DC31{Cf51}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC32_(false, _dafny.Zero, _dafny.EmptyMap, _dafny.EmptySeq, false)
}

func (_this D14) Dtor_cf52() bool {
	return _this.Get_().(D14_DC32).Cf52
}

func (_this D14) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D14_DC32).Cf53
}

func (_this D14) Dtor_cf54() _dafny.Map {
	return _this.Get_().(D14_DC32).Cf54
}

func (_this D14) Dtor_cf55() _dafny.Sequence {
	return _this.Get_().(D14_DC32).Cf55
}

func (_this D14) Dtor_cf56() bool {
	return _this.Get_().(D14_DC32).Cf56
}

func (_this D14) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D14_DC31).Cf51
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC32:
		{
			return "D14.DC32" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + data.Cf55.VerbatimString(true) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC32:
		{
			data2, ok := other.Get_().(D14_DC32)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54.Equals(data2.Cf54) && data1.Cf55.Equals(data2.Cf55) && data1.Cf56 == data2.Cf56
		}
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D15_DC34 struct {
	Cf58 _dafny.Int
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf58 _dafny.Int) D15 {
	return D15{D15_DC34{Cf58}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

type D15_DC33 struct {
	Cf57 _dafny.Array
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf57 _dafny.Array) D15 {
	return D15{D15_DC33{Cf57}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC34_(_dafny.Zero)
}

func (_this D15) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D15_DC34).Cf58
}

func (_this D15) Dtor_cf57() _dafny.Array {
	return _this.Get_().(D15_DC33).Cf57
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC34:
		{
			return "D15.DC34" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC34:
		{
			data2, ok := other.Get_().(D15_DC34)
			return ok && data1.Cf58.Cmp(data2.Cf58) == 0
		}
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf57 == data2.Cf57
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

type D16_DC36 struct {
	Cf60 bool
	Cf61 _dafny.Sequence
	Cf62 _dafny.Array
}

func (D16_DC36) isD16() {}

func (CompanionStruct_D16_) Create_DC36_(Cf60 bool, Cf61 _dafny.Sequence, Cf62 _dafny.Array) D16 {
	return D16{D16_DC36{Cf60, Cf61, Cf62}}
}

func (_this D16) Is_DC36() bool {
	_, ok := _this.Get_().(D16_DC36)
	return ok
}

type D16_DC35 struct {
	Cf59 _dafny.Map
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf59 _dafny.Map) D16 {
	return D16{D16_DC35{Cf59}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC36_(false, _dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D16) Dtor_cf60() bool {
	return _this.Get_().(D16_DC36).Cf60
}

func (_this D16) Dtor_cf61() _dafny.Sequence {
	return _this.Get_().(D16_DC36).Cf61
}

func (_this D16) Dtor_cf62() _dafny.Array {
	return _this.Get_().(D16_DC36).Cf62
}

func (_this D16) Dtor_cf59() _dafny.Map {
	return _this.Get_().(D16_DC35).Cf59
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC36:
		{
			return "D16.DC36" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC36:
		{
			data2, ok := other.Get_().(D16_DC36)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61.Equals(data2.Cf61) && data1.Cf62 == data2.Cf62
		}
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf59.Equals(data2.Cf59)
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

type D17_DC38 struct {
}

func (D17_DC38) isD17() {}

func (CompanionStruct_D17_) Create_DC38_() D17 {
	return D17{D17_DC38{}}
}

func (_this D17) Is_DC38() bool {
	_, ok := _this.Get_().(D17_DC38)
	return ok
}

type D17_DC37 struct {
	Cf63 *C2
}

func (D17_DC37) isD17() {}

func (CompanionStruct_D17_) Create_DC37_(Cf63 *C2) D17 {
	return D17{D17_DC37{Cf63}}
}

func (_this D17) Is_DC37() bool {
	_, ok := _this.Get_().(D17_DC37)
	return ok
}

type D17_DC39 struct {
	Cf64 D17
}

func (D17_DC39) isD17() {}

func (CompanionStruct_D17_) Create_DC39_(Cf64 D17) D17 {
	return D17{D17_DC39{Cf64}}
}

func (_this D17) Is_DC39() bool {
	_, ok := _this.Get_().(D17_DC39)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC38_()
}

func (_this D17) Dtor_cf63() *C2 {
	return _this.Get_().(D17_DC37).Cf63
}

func (_this D17) Dtor_cf64() D17 {
	return _this.Get_().(D17_DC39).Cf64
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC38:
		{
			return "D17.DC38"
		}
	case D17_DC37:
		{
			return "D17.DC37" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D17_DC39:
		{
			return "D17.DC39" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC38:
		{
			_, ok := other.Get_().(D17_DC38)
			return ok
		}
	case D17_DC37:
		{
			data2, ok := other.Get_().(D17_DC37)
			return ok && data1.Cf63 == data2.Cf63
		}
	case D17_DC39:
		{
			data2, ok := other.Get_().(D17_DC39)
			return ok && data1.Cf64.Equals(data2.Cf64)
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

type D18_DC40 struct {
	Cf65 _dafny.MultiSet
}

func (D18_DC40) isD18() {}

func (CompanionStruct_D18_) Create_DC40_(Cf65 _dafny.MultiSet) D18 {
	return D18{D18_DC40{Cf65}}
}

func (_this D18) Is_DC40() bool {
	_, ok := _this.Get_().(D18_DC40)
	return ok
}

func (CompanionStruct_D18_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D18) Dtor_cf65() _dafny.MultiSet {
	return _this.Get_().(D18_DC40).Cf65
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC40:
		{
			return "D18.DC40" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC40:
		{
			data2, ok := other.Get_().(D18_DC40)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D19_DC42 struct {
	Cf67 _dafny.MultiSet
	Cf68 _dafny.Int
	Cf69 _dafny.Int
}

func (D19_DC42) isD19() {}

func (CompanionStruct_D19_) Create_DC42_(Cf67 _dafny.MultiSet, Cf68 _dafny.Int, Cf69 _dafny.Int) D19 {
	return D19{D19_DC42{Cf67, Cf68, Cf69}}
}

func (_this D19) Is_DC42() bool {
	_, ok := _this.Get_().(D19_DC42)
	return ok
}

type D19_DC43 struct {
	Cf70 *C2
	Cf71 _dafny.Int
	Cf72 bool
}

func (D19_DC43) isD19() {}

func (CompanionStruct_D19_) Create_DC43_(Cf70 *C2, Cf71 _dafny.Int, Cf72 bool) D19 {
	return D19{D19_DC43{Cf70, Cf71, Cf72}}
}

func (_this D19) Is_DC43() bool {
	_, ok := _this.Get_().(D19_DC43)
	return ok
}

type D19_DC41 struct {
	Cf66 *C4
}

func (D19_DC41) isD19() {}

func (CompanionStruct_D19_) Create_DC41_(Cf66 *C4) D19 {
	return D19{D19_DC41{Cf66}}
}

func (_this D19) Is_DC41() bool {
	_, ok := _this.Get_().(D19_DC41)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC42_(_dafny.EmptyMultiSet, _dafny.Zero, _dafny.Zero)
}

func (_this D19) Dtor_cf67() _dafny.MultiSet {
	return _this.Get_().(D19_DC42).Cf67
}

func (_this D19) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D19_DC42).Cf68
}

func (_this D19) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D19_DC42).Cf69
}

func (_this D19) Dtor_cf70() *C2 {
	return _this.Get_().(D19_DC43).Cf70
}

func (_this D19) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D19_DC43).Cf71
}

func (_this D19) Dtor_cf72() bool {
	return _this.Get_().(D19_DC43).Cf72
}

func (_this D19) Dtor_cf66() *C4 {
	return _this.Get_().(D19_DC41).Cf66
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC42:
		{
			return "D19.DC42" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D19_DC43:
		{
			return "D19.DC43" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D19_DC41:
		{
			return "D19.DC41" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC42:
		{
			data2, ok := other.Get_().(D19_DC42)
			return ok && data1.Cf67.Equals(data2.Cf67) && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69.Cmp(data2.Cf69) == 0
		}
	case D19_DC43:
		{
			data2, ok := other.Get_().(D19_DC43)
			return ok && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72 == data2.Cf72
		}
	case D19_DC41:
		{
			data2, ok := other.Get_().(D19_DC41)
			return ok && data1.Cf66 == data2.Cf66
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

type D20_DC45 struct {
	Cf74 _dafny.Int
}

func (D20_DC45) isD20() {}

func (CompanionStruct_D20_) Create_DC45_(Cf74 _dafny.Int) D20 {
	return D20{D20_DC45{Cf74}}
}

func (_this D20) Is_DC45() bool {
	_, ok := _this.Get_().(D20_DC45)
	return ok
}

type D20_DC46 struct {
	Cf75 bool
	Cf76 _dafny.Int
}

func (D20_DC46) isD20() {}

func (CompanionStruct_D20_) Create_DC46_(Cf75 bool, Cf76 _dafny.Int) D20 {
	return D20{D20_DC46{Cf75, Cf76}}
}

func (_this D20) Is_DC46() bool {
	_, ok := _this.Get_().(D20_DC46)
	return ok
}

type D20_DC47 struct {
}

func (D20_DC47) isD20() {}

func (CompanionStruct_D20_) Create_DC47_() D20 {
	return D20{D20_DC47{}}
}

func (_this D20) Is_DC47() bool {
	_, ok := _this.Get_().(D20_DC47)
	return ok
}

type D20_DC44 struct {
	Cf73 _dafny.Map
}

func (D20_DC44) isD20() {}

func (CompanionStruct_D20_) Create_DC44_(Cf73 _dafny.Map) D20 {
	return D20{D20_DC44{Cf73}}
}

func (_this D20) Is_DC44() bool {
	_, ok := _this.Get_().(D20_DC44)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC45_(_dafny.Zero)
}

func (_this D20) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D20_DC45).Cf74
}

func (_this D20) Dtor_cf75() bool {
	return _this.Get_().(D20_DC46).Cf75
}

func (_this D20) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D20_DC46).Cf76
}

func (_this D20) Dtor_cf73() _dafny.Map {
	return _this.Get_().(D20_DC44).Cf73
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC45:
		{
			return "D20.DC45" + "(" + _dafny.String(data.Cf74) + ")"
		}
	case D20_DC46:
		{
			return "D20.DC46" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D20_DC47:
		{
			return "D20.DC47"
		}
	case D20_DC44:
		{
			return "D20.DC44" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC45:
		{
			data2, ok := other.Get_().(D20_DC45)
			return ok && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D20_DC46:
		{
			data2, ok := other.Get_().(D20_DC46)
			return ok && data1.Cf75 == data2.Cf75 && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D20_DC47:
		{
			_, ok := other.Get_().(D20_DC47)
			return ok
		}
	case D20_DC44:
		{
			data2, ok := other.Get_().(D20_DC44)
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D21_DC49 struct {
	Cf78 bool
	Cf79 _dafny.Array
}

func (D21_DC49) isD21() {}

func (CompanionStruct_D21_) Create_DC49_(Cf78 bool, Cf79 _dafny.Array) D21 {
	return D21{D21_DC49{Cf78, Cf79}}
}

func (_this D21) Is_DC49() bool {
	_, ok := _this.Get_().(D21_DC49)
	return ok
}

type D21_DC48 struct {
	Cf77 _dafny.Array
}

func (D21_DC48) isD21() {}

func (CompanionStruct_D21_) Create_DC48_(Cf77 _dafny.Array) D21 {
	return D21{D21_DC48{Cf77}}
}

func (_this D21) Is_DC48() bool {
	_, ok := _this.Get_().(D21_DC48)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC49_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D21) Dtor_cf78() bool {
	return _this.Get_().(D21_DC49).Cf78
}

func (_this D21) Dtor_cf79() _dafny.Array {
	return _this.Get_().(D21_DC49).Cf79
}

func (_this D21) Dtor_cf77() _dafny.Array {
	return _this.Get_().(D21_DC48).Cf77
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC49:
		{
			return "D21.DC49" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D21_DC48:
		{
			return "D21.DC48" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC49:
		{
			data2, ok := other.Get_().(D21_DC49)
			return ok && data1.Cf78 == data2.Cf78 && data1.Cf79 == data2.Cf79
		}
	case D21_DC48:
		{
			data2, ok := other.Get_().(D21_DC48)
			return ok && data1.Cf77 == data2.Cf77
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

type D22_DC51 struct {
	Cf81 _dafny.Int
	Cf82 bool
}

func (D22_DC51) isD22() {}

func (CompanionStruct_D22_) Create_DC51_(Cf81 _dafny.Int, Cf82 bool) D22 {
	return D22{D22_DC51{Cf81, Cf82}}
}

func (_this D22) Is_DC51() bool {
	_, ok := _this.Get_().(D22_DC51)
	return ok
}

type D22_DC50 struct {
	Cf80 _dafny.Set
}

func (D22_DC50) isD22() {}

func (CompanionStruct_D22_) Create_DC50_(Cf80 _dafny.Set) D22 {
	return D22{D22_DC50{Cf80}}
}

func (_this D22) Is_DC50() bool {
	_, ok := _this.Get_().(D22_DC50)
	return ok
}

type D22_DC52 struct {
	Cf83 D22
}

func (D22_DC52) isD22() {}

func (CompanionStruct_D22_) Create_DC52_(Cf83 D22) D22 {
	return D22{D22_DC52{Cf83}}
}

func (_this D22) Is_DC52() bool {
	_, ok := _this.Get_().(D22_DC52)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC51_(_dafny.Zero, false)
}

func (_this D22) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D22_DC51).Cf81
}

func (_this D22) Dtor_cf82() bool {
	return _this.Get_().(D22_DC51).Cf82
}

func (_this D22) Dtor_cf80() _dafny.Set {
	return _this.Get_().(D22_DC50).Cf80
}

func (_this D22) Dtor_cf83() D22 {
	return _this.Get_().(D22_DC52).Cf83
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC51:
		{
			return "D22.DC51" + "(" + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D22_DC50:
		{
			return "D22.DC50" + "(" + _dafny.String(data.Cf80) + ")"
		}
	case D22_DC52:
		{
			return "D22.DC52" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC51:
		{
			data2, ok := other.Get_().(D22_DC51)
			return ok && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82 == data2.Cf82
		}
	case D22_DC50:
		{
			data2, ok := other.Get_().(D22_DC50)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	case D22_DC52:
		{
			data2, ok := other.Get_().(D22_DC52)
			return ok && data1.Cf83.Equals(data2.Cf83)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int
	Fm7(globalState *GlobalState) bool
	M1(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) bool
	F25() bool
	F26() _dafny.Int
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
	Fm16(globalState *GlobalState) _dafny.Int
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
	F34() bool
	F34_set_(value bool)
	Fm17(p0 bool, p1 bool, globalState *GlobalState) bool
	Fm18(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map
	F33() _dafny.Array
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
	F1   bool
	F2   _dafny.Int
	_f3  _dafny.Array
	_f4  bool
	_f5  _dafny.Int
	_f6  _dafny.Int
	_f7  bool
	_f8  _dafny.Sequence
	_f9  _dafny.Sequence
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.Sequence
	_f13 _dafny.Map
	_f14 _dafny.Int
	_f15 _dafny.Int
	_f16 _dafny.Int
	_f17 _dafny.Set
	_f18 bool
	_f19 _dafny.Int
	_f20 bool
	_f21 bool
	_f22 bool
	_f23 bool
	_f24 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F1 = false
	_this.F2 = _dafny.Zero
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.EmptySeq
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.EmptySeq
	_this._f13 = _dafny.EmptyMap
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.EmptySet
	_this._f18 = false
	_this._f19 = _dafny.Zero
	_this._f20 = false
	_this._f21 = false
	_this._f22 = false
	_this._f23 = false
	_this._f24 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Int, f3 _dafny.Array, f4 bool, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 _dafny.Sequence, f9 _dafny.Sequence, f10 _dafny.Int, f11 bool, f12 _dafny.Sequence, f13 _dafny.Map, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Int, f17 _dafny.Set, f18 bool, f19 _dafny.Int, f20 bool, f21 bool, f22 bool, f23 bool, f24 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *GlobalState) F3() _dafny.Array {
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
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
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
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Map {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
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
func (_this *GlobalState) F17() _dafny.Set {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
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
func (_this *GlobalState) F21() bool {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() bool {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() bool {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() bool {
	{
		return _this._f24
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f34 bool
	_f25 bool
	_f26 _dafny.Int
	_f33 _dafny.Array
	F35  _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f34 = false
	_this._f25 = false
	_this._f26 = _dafny.Zero
	_this._f33 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F35 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_, Companion_T2_.TraitID_}
}

var _ T0 = &C0{}
var _ T1 = &C0{}
var _ T2 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F34() bool {
	return _this._f34
}
func (_this *C0) F34_set_(value bool) {
	_this._f34 = value
}
func (_this *C0) F25() bool {
	return _this._f25
}
func (_this *C0) F26() _dafny.Int {
	return _this._f26
}
func (_this *C0) F33() _dafny.Array {
	return _this._f33
}
func (_this *C0) Ctor__(f35 _dafny.Sequence, f25 bool, f26 _dafny.Int, f33 _dafny.Array, f34 bool) {
	{
		(_this).F35 = f35
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f33 = f33
		(_this)._f34 = f34
	}
}
func (_this *C0) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if (_dafny.MultiSetOf(_dafny.SetOf((_this).F25(), !(_this.F34())))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.SetOf((_this).F25(), _this.F34()))) {
			return (_this).F26()
		} else {
			return (_this).F26()
		}
	}
}
func (_this *C0) Fm7(globalState *GlobalState) bool {
	{
		return (_this).F25()
	}
}
func (_this *C0) Fm16(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(162)).Plus((_this).F26())
	}
}
func (_this *C0) Fm17(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return ((_this).F26()).Cmp(((_this).F26()).Minus((_this).F26())) != 0
	}
}
func (_this *C0) Fm18(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(((_dafny.MultiSetOf(_this.F35, _this.F35, _this.F35, _this.F35, _this.F35)).Difference(_dafny.MultiSetOf(_this.F35, _this.F35, _this.F35))).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _375_v0 _dafny.Sequence
				_375_v0 = interface{}(_compr_9).(_dafny.Sequence)
				if ((_dafny.MultiSetOf(_this.F35, _this.F35, _this.F35, _this.F35, _this.F35)).Difference(_dafny.MultiSetOf(_this.F35, _this.F35, _this.F35))).Contains(_375_v0) {
					_coll9.Add(_375_v0, (_this).F25())
				}
			}
			return _coll9.ToMap()
		}()
	}
}
func (_this *C0) Fm19(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.MultiSetOf((func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_this.F35).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _376_v0 _dafny.CodePoint
				_376_v0 = interface{}(_compr_10).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_this.F35, _376_v0) {
					_coll10.Add(_376_v0, _this.F34())
				}
			}
			return _coll10.ToMap()
		}()).Cardinality())).Intersection((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F26()))))).Cardinality()
	}
}
func (_this *C0) Fm20(p0 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	{
		var _source4 D2 = Companion_D2_.Create_DC7_((_this).F25(), true, (_this).F26())
		_ = _source4
		if _source4.Is_DC7() {
			var _377___mcc_h0 bool = _source4.Get_().(D2_DC7).Cf11
			_ = _377___mcc_h0
			var _378___mcc_h1 bool = _source4.Get_().(D2_DC7).Cf12
			_ = _378___mcc_h1
			var _379___mcc_h2 _dafny.Int = _source4.Get_().(D2_DC7).Cf13
			_ = _379___mcc_h2
			var _380_cf13 _dafny.Int = _379___mcc_h2
			_ = _380_cf13
			var _381_cf12 bool = _378___mcc_h1
			_ = _381_cf12
			var _382_cf11 bool = _377___mcc_h0
			_ = _382_cf11
			return _dafny.CodePoint('m')
		} else {
			var _383___mcc_h3 _dafny.Set = _source4.Get_().(D2_DC6).Cf10
			_ = _383___mcc_h3
			var _384_cf10 _dafny.Set = _383___mcc_h3
			_ = _384_cf10
			return _dafny.CodePoint('w')
		}
	}
}
func (_this *C0) M1(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F2 = (_this).F26()
		var _385_v0 _dafny.MultiSet
		_ = _385_v0
		_385_v0 = _dafny.MultiSetOf(_this.F35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(449))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_386_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})))
		(globalState).F2 = (func() _dafny.Int {
			if (_385_v0).Contains(_dafny.Companion_Sequence_.Concatenate(_this.F35, _dafny.UnicodeSeqOfUtf8Bytes("hfhbn"))) {
				return (_385_v0).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_this.F35, _dafny.UnicodeSeqOfUtf8Bytes("hfhbn")))
			}
			return (func() _dafny.Int {
				if p1 {
					return (_this).F26()
				}
				return (p0).Cardinality()
			})()
		})()
		var _387_v1 _dafny.Array
		_ = _387_v1
		var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw74
		_387_v1 = _nw74
		var _388_v2 _dafny.Sequence
		_ = _388_v2
		_388_v2 = _dafny.SeqOf((_this).F33())
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_387_v1), 0))
		_ = _index61
		(_387_v1).ArraySet1(_388_v2, (_index61).Int())
		var _389_v3 _dafny.Array
		_ = _389_v3
		var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw75
		_389_v3 = _nw75
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_389_v3), 0))
		_ = _index62
		(_389_v3).ArraySet1(_this.F35, (_index62).Int())
		var _390_v4 _dafny.Array
		_ = _390_v4
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw76
		_390_v4 = _nw76
		var _391_v5 _dafny.Map
		_ = _391_v5
		_391_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v4, Companion_Default___.SafeModuloInt((_this).F26(), _dafny.IntOfInt64(574)))
		var _392_v6 _dafny.Map
		_ = _392_v6
		_392_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.Companion_Sequence_.Concatenate(_388_v2, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_388_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.IntOfUint32((_388_v2).Cardinality()))).Uint32(), (_this).F33()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F35).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_388_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.IntOfUint32((_388_v2).Cardinality()))).Uint32(), (_this).F33())).Cardinality()))).Uint32(), (_this).F33())))
		var _393_v7 _dafny.Map
		_ = _393_v7
		_393_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
		var _394_v8 _dafny.Map
		_ = _394_v8
		_394_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_393_v7).Cardinality(), p1)
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_387_v1), 0))
		_ = _index63
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_389_v3), 0))
		_ = _index64
		var _rhs73 _dafny.Sequence = (func() _dafny.Sequence {
			if (_392_v6).Contains((func() bool {
				if (_394_v8).Contains((_this).F26()) {
					return (_394_v8).Get((_this).F26()).(bool)
				}
				return _this.F34()
			})()) {
				return (_392_v6).Get((func() bool {
					if (_394_v8).Contains((_this).F26()) {
						return (_394_v8).Get((_this).F26()).(bool)
					}
					return _this.F34()
				})()).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_388_v2, _dafny.SeqOf((_this).F33(), (_this).F33()))
		})()
		_ = _rhs73
		var _rhs74 bool = (_this).F25()
		_ = _rhs74
		var _rhs75 _dafny.Sequence = _this.F35
		_ = _rhs75
		var _rhs76 _dafny.Map = _391_v5
		_ = _rhs76
		var _lhs62 _dafny.Array = _387_v1
		_ = _lhs62
		var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_387_v1), 0))
		_ = _lhs63
		var _lhs64 *GlobalState = globalState
		_ = _lhs64
		var _lhs65 _dafny.Array = _389_v3
		_ = _lhs65
		var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_389_v3), 0))
		_ = _lhs66
		(_lhs62).ArraySet1(_rhs73, (_lhs63).Int())
		_lhs64.F1 = _rhs74
		(_lhs65).ArraySet1(_rhs75, (_lhs66).Int())
		_391_v5 = _rhs76
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))
		_ = _index65
		(_390_v4).ArraySet1(Companion_Default___.SafeModuloInt((_this).F26(), _dafny.IntOfUint32(((_389_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_389_v3), 0))).Int()).(_dafny.Sequence)).Cardinality())), (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))
		_ = _index66
		var _rhs77 bool = p1
		_ = _rhs77
		var _rhs78 _dafny.Int = (_dafny.Zero).Minus((_this).F26())
		_ = _rhs78
		var _rhs79 bool = p2
		_ = _rhs79
		var _lhs67 _dafny.Array = _390_v4
		_ = _lhs67
		var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))
		_ = _lhs68
		r0 = _rhs77
		(_lhs67).ArraySet1(_rhs78, (_lhs68).Int())
		r0 = _rhs79
		(globalState).F2 = Companion_Default___.SafeDivisionInt((_this).F26(), ((_390_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))).Int()).(_dafny.Int)).Times((_390_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))).Int()).(_dafny.Int)))
		var _395_v9 _dafny.Map
		_ = _395_v9
		_395_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34(), (_this).F26())
		(globalState).F2 = (func() _dafny.Int {
			if !(_395_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F34(), _dafny.IntOfUint32(((_389_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_389_v3), 0))).Int()).(_dafny.Sequence)).Cardinality()))) {
				return (_this).F26()
			}
			return (_dafny.IntOfUint32((_dafny.SeqOf(!(_this.F34()))).Cardinality())).Minus((_390_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_390_v4), 0))).Int()).(_dafny.Int))
		})()
		r0 = p2
		return r0
	}
}
func (_this *C0) M9(p0 _dafny.MultiSet, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _396_i0 _dafny.Int
		_ = _396_i0
		_396_i0 = _dafny.Zero
		{
			for (true) && ((_this).F25()) {
				{
					if (_396_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_396_i0 = (_396_i0).Plus(_dafny.One)
					(globalState).F1 = ((_this).F26()).Cmp((_this).F26()) == 0
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))
					_ = _index67
					((_this).F33()).ArraySet1(_this.F34(), (_index67).Int())
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))
					_ = _index68
					var _rhs80 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-188), (_dafny.Zero).Minus((_this).F26()))
					_ = _rhs80
					var _rhs81 bool = _this.F34()
					_ = _rhs81
					var _rhs82 bool = _this.F34()
					_ = _rhs82
					var _rhs83 _dafny.Int = (func() _dafny.Int {
						if ((_this).F26()).Cmp((_this).F26()) <= 0 {
							return (_this).F26()
						}
						return (_this).F26()
					})()
					_ = _rhs83
					var _lhs69 *GlobalState = globalState
					_ = _lhs69
					var _lhs70 _dafny.Array = (_this).F33()
					_ = _lhs70
					var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))
					_ = _lhs71
					var _lhs72 *GlobalState = globalState
					_ = _lhs72
					var _lhs73 *GlobalState = globalState
					_ = _lhs73
					_lhs69.F0 = _rhs80
					(_lhs70).ArraySet1(_rhs81, (_lhs71).Int())
					_lhs72.F1 = _rhs82
					_lhs73.F2 = _rhs83
					var _397_v0 _dafny.Sequence
					_ = _397_v0
					_397_v0 = _dafny.SeqOf((_this).F26())
					(globalState).F2 = Companion_Default___.Fm2((func() _dafny.Int {
						if _this.F34() {
							return (_397_v0).Select((Companion_Default___.SafeIndex((_this).Fm16(globalState), _dafny.IntOfUint32((_397_v0).Cardinality()))).Uint32()).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-33)
					})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_397_v0).Cardinality())), globalState)
					var _398_v1 _dafny.Set
					_ = _398_v1
					_398_v1 = _dafny.SetOf(((_this).F33()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))).Int()).(bool), ((_this).F33()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))).Int()).(bool))
					var _399_v2 _dafny.Map
					_ = _399_v2
					_399_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_398_v1).Cardinality())
					var _400_v3 D2
					_ = _400_v3
					_400_v3 = Companion_D2_.Create_DC7_((_this).F25(), _this.F34(), (func() _dafny.Int {
						if (_399_v2).Contains((_this).F26()) {
							return (_399_v2).Get((_this).F26()).(_dafny.Int)
						}
						return (_397_v0).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_397_v0).Cardinality()))).Uint32()).(_dafny.Int)
					})())
					var _401_v4 _dafny.Array
					_ = _401_v4
					var _nwElement0_16 D2 = _400_v3
					_ = _nwElement0_16
					var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(21))
					_ = _nw77
					(_nw77).ArraySet1(_nwElement0_16, 0)
					(_nw77).ArraySet1(_400_v3, 1)
					(_nw77).ArraySet1(_400_v3, 2)
					(_nw77).ArraySet1(_400_v3, 3)
					(_nw77).ArraySet1(_400_v3, 4)
					(_nw77).ArraySet1(_400_v3, 5)
					(_nw77).ArraySet1(Companion_D2_.Create_DC7_(_this.F34(), ((_this).F33()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))).Int()).(bool), (_this).F26()), 6)
					(_nw77).ArraySet1(_400_v3, 7)
					(_nw77).ArraySet1(Companion_D2_.Create_DC7_((_this).F25(), ((_this).F33()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen(((_this).F33()), 0))).Int()).(bool), _dafny.IntOfInt64(442)), 8)
					(_nw77).ArraySet1(_400_v3, 9)
					(_nw77).ArraySet1(_400_v3, 10)
					(_nw77).ArraySet1(_400_v3, 11)
					(_nw77).ArraySet1(_400_v3, 12)
					(_nw77).ArraySet1(_400_v3, 13)
					(_nw77).ArraySet1((func() D2 {
						if (_this).F25() {
							return _400_v3
						}
						return _400_v3
					})(), 14)
					(_nw77).ArraySet1(_400_v3, 15)
					(_nw77).ArraySet1(_400_v3, 16)
					(_nw77).ArraySet1(_400_v3, 17)
					(_nw77).ArraySet1(_400_v3, 18)
					(_nw77).ArraySet1(_400_v3, 19)
					(_nw77).ArraySet1(_400_v3, 20)
					_401_v4 = _nw77
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_401_v4), 0))
					_ = _index69
					(_401_v4).ArraySet1(_400_v3, (_index69).Int())
					var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_401_v4), 0))
					_ = _index70
					(_401_v4).ArraySet1(_400_v3, (_index70).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _402_v5 _dafny.Map
		_ = _402_v5
		_402_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F26(), (_this).F26()), p0)
		_402_v5 = (_402_v5).Update(((p0).Intersection(p0)).Cardinality(), (p0).Update((_this).F25(), Companion_Default___.Abs((_this).F26())))
		(_this).F34_set_(_this.F34())
		(_this).F34_set_(_this.F34())
		var _403_v6 _dafny.Map
		_ = _403_v6
		_403_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
		var _404_v7 _dafny.Map
		_ = _404_v7
		_404_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _403_v6)
		var _hi2 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (p0).Contains((_this).F25()) {
				return (p0).Multiplicity((_this).F25())
			}
			return (_404_v7).Cardinality()
		})(), (_this).F26())
		_ = _hi2
		for _405_i1 := (_dafny.Zero).Minus(((_this).Fm6((_this).F25(), (_this).F25(), (_this).F26(), globalState)).Plus((_this).F26())); _405_i1.Cmp(_hi2) < 0; _405_i1 = _405_i1.Plus(_dafny.One) {
			(globalState).F2 = _dafny.IntOfInt64(480)
			(globalState).F1 = (_405_i1).Cmp((_this).F26()) <= 0
			(globalState).F1 = (_this).F25()
			(globalState).F1 = true
		}
		(globalState).F0 = (_this).F26()
		var _406_v8 _dafny.CodePoint
		_ = _406_v8
		_406_v8 = _dafny.CodePoint('g')
		r0 = _406_v8
		r1 = (_this).F33()
		return r0, r1
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f32 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f32 = false
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

func (_this *C1) Ctor__(f32 bool) {
	{
		(_this)._f32 = f32
	}
}
func (_this *C1) Fm14(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ehss"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(556))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_407_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))), _dafny.UnicodeSeqOfUtf8Bytes("kbxgttvgb"))
	}
}
func (_this *C1) Fm15(p0 _dafny.MultiSet, p1 _dafny.Int, p2 D1, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(967))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_408_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(644)
		}))).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_409_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality()), _dafny.IntOfInt64(-971), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-428))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_410_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(907))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_411_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), _dafny.IntOfInt64(943))).Cardinality()))
	}
}
func (_this *C1) M7(p0 bool, p1 D0, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.MultiSet, _dafny.Map) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _412_v0 _dafny.Array
		_ = _412_v0
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_9
		var _nw78 _dafny.Array
		_ = _nw78
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw78 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) bool = func(_413_i0 _dafny.Int) bool {
				return (_this).F32()
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw78 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw78).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw78).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_412_v0 = _nw78
		var _414_v1 _dafny.CodePoint
		_ = _414_v1
		_414_v1 = _dafny.CodePoint('o')
		var _415_v2 _dafny.Map
		_ = _415_v2
		_415_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v1, (_this).F32())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
		_ = _index71
		(_412_v0).ArraySet1((func() bool {
			if (_415_v2).Contains(_414_v1) {
				return (_415_v2).Get(_414_v1).(bool)
			}
			return (_this).F32()
		})(), (_index71).Int())
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
		_ = _index72
		(_412_v0).ArraySet1((_this).F32(), (_index72).Int())
		var _416_v3 _dafny.Sequence
		_ = _416_v3
		_416_v3 = _dafny.UnicodeSeqOfUtf8Bytes("qby")
		_416_v3 = p2
		var _417_v4 _dafny.Int
		_ = _417_v4
		_417_v4 = _dafny.IntOfInt64(238)
		var _418_v5 _dafny.Map
		_ = _418_v5
		_418_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _417_v4)
		if Companion_Default___.Fm5(!((_this).F32()), _418_v5, _417_v4, globalState) {
			var _419_v6 _dafny.Map
			_ = _419_v6
			_419_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v4, p0)
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
			_ = _index73
			(_412_v0).ArraySet1(((_419_v6).Update(_417_v4, (_this).F32())).Equals((_419_v6).Merge(_419_v6)), (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
			_ = _index74
			var _rhs84 _dafny.Int = _417_v4
			_ = _rhs84
			var _rhs85 bool = !((Companion_D0_.Create_DC1_(_417_v4, p2, (_this).F32(), (_dafny.Zero).Minus(_417_v4), p0)).Dtor_cf5())
			_ = _rhs85
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			var _lhs75 _dafny.Array = _412_v0
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
			_ = _lhs76
			_lhs74.F2 = _rhs84
			(_lhs75).ArraySet1(_rhs85, (_lhs76).Int())
			(globalState).F1 = (_this).F32()
			var _420_v7 _dafny.Set
			_ = _420_v7
			_420_v7 = _dafny.SetOf(false)
			var _421_v8 D4
			_ = _421_v8
			_421_v8 = Companion_D4_.Create_DC9_(_420_v7)
			var _rhs86 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_416_v3, _416_v3)
			_ = _rhs86
			var _rhs87 _dafny.Int = Companion_Default___.Fm2(_417_v4, (((_421_v8).Dtor_cf15()).Difference(_420_v7)).Cardinality(), globalState)
			_ = _rhs87
			var _rhs88 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F32() {
					return (_dafny.Zero).Minus(_417_v4)
				}
				return ((_420_v7).Difference(_dafny.SetOf((_this).F32()))).Cardinality()
			})())
			_ = _rhs88
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			_416_v3 = _rhs86
			_lhs77.F2 = _rhs87
			_lhs78.F0 = _rhs88
			(globalState).F0 = _417_v4
		} else {
			var _422_v9 _dafny.Map
			_ = _422_v9
			_422_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool), _dafny.IntOfInt64(347))
			var _423_v10 _dafny.Sequence
			_ = _423_v10
			_423_v10 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _417_v4), _422_v9)
			var _424_v11 _dafny.Map
			_ = _424_v11
			_424_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v4, _dafny.MultiSetFromSeq(_423_v10))
			var _425_v12 _dafny.Set
			_ = _425_v12
			_425_v12 = _dafny.SetOf((_this).F32(), true, p0)
			var _426_v13 _dafny.MultiSet
			_ = _426_v13
			_426_v13 = _dafny.MultiSetOf((_422_v9).Update((_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool), _417_v4))
			var _427_v14 _dafny.MultiSet
			_ = _427_v14
			_427_v14 = (func() _dafny.MultiSet {
				if (_424_v11).Contains((_425_v12).Cardinality()) {
					return (_424_v11).Get((_425_v12).Cardinality()).(_dafny.MultiSet)
				}
				return _426_v13
			})()
			r0 = _427_v14
			var _428_v15 D0
			_ = _428_v15
			_428_v15 = Companion_D0_.Create_DC2_()
			var _429_v16 _dafny.Map
			_ = _429_v16
			_429_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_428_v15, (_this).F32())
			if (_429_v16).Contains(_428_v15) {
				var _430_v17 _dafny.Array
				_ = _430_v17
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_10
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = (func(_431_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_432_i1 _dafny.Int) _dafny.Int {
							return (_432_i1).Times(_431_v4)
						}
					})(_417_v4)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw79 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw79).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw79).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_430_v17 = _nw79
				var _rhs89 bool = (_417_v4).Cmp((_417_v4).Plus(_417_v4)) < 0
				_ = _rhs89
				var _rhs90 _dafny.Array = _430_v17
				_ = _rhs90
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_lhs79.F1 = _rhs89
				_430_v17 = _rhs90
				_416_v3 = _dafny.Companion_Sequence_.Concatenate(_416_v3, _416_v3)
				(globalState).F2 = _417_v4
				var _433_v18 D0
				_ = _433_v18
				_433_v18 = Companion_D0_.Create_DC1_(_417_v4, _416_v3, p0, _417_v4, (_this).F32())
				var _434_v19 *C0
				_ = _434_v19
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__((func(_pat_let10_0 D0) D0 {
					return func(_435_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let11_0 bool) D0 {
							return func(_436_dt__update_hcf5_h0 bool) D0 {
								return Companion_D0_.Create_DC1_((_435_dt__update__tmp_h0).Dtor_cf1(), (_435_dt__update__tmp_h0).Dtor_cf2(), (_435_dt__update__tmp_h0).Dtor_cf3(), (_435_dt__update__tmp_h0).Dtor_cf4(), _436_dt__update_hcf5_h0)
							}(_pat_let11_0)
						}((_this).F32())
					}(_pat_let10_0)
				}(_433_v18)).Dtor_cf2(), !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("yggum"), _416_v3), _417_v4, _412_v0, (_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool))
				_434_v19 = _nw80
				(globalState).F1 = p0
			} else {
				(globalState).F0 = _417_v4
				_414_v1 = _414_v1
				var _437_v20 _dafny.Sequence
				_ = _437_v20
				_437_v20 = _dafny.SeqOf((_this).F32())
				var _438_v21 _dafny.Array
				_ = _438_v21
				var _nwElement0_17 _dafny.Int = (_dafny.IntOfInt64(224)).Minus(_417_v4)
				_ = _nwElement0_17
				var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(10))
				_ = _nw81
				(_nw81).ArraySet1(_nwElement0_17, 0)
				(_nw81).ArraySet1(_417_v4, 1)
				(_nw81).ArraySet1(_dafny.IntOfInt64(931), 2)
				(_nw81).ArraySet1(_417_v4, 3)
				(_nw81).ArraySet1(_417_v4, 4)
				(_nw81).ArraySet1((_417_v4).Minus(_dafny.IntOfInt64(-173)), 5)
				(_nw81).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 6)
				(_nw81).ArraySet1(_417_v4, 7)
				(_nw81).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_437_v20).Cardinality()), _417_v4), 8)
				(_nw81).ArraySet1((_dafny.Zero).Minus(_417_v4), 9)
				_438_v21 = _nw81
				var _439_v22 D0
				_ = _439_v22
				_439_v22 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(791), _416_v3, (_this).F32(), _417_v4, true)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_438_v21), 0))
				_ = _index75
				(_438_v21).ArraySet1(Companion_Default___.Fm2(_417_v4, (_439_v22).Dtor_cf1(), globalState), (_index75).Int())
				var _440_v23 _dafny.Map
				_ = _440_v23
				_440_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v4, _417_v4)
				var _441_v24 _dafny.Map
				_ = _441_v24
				_441_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_438_v21, _440_v23)
				var _442_v25 _dafny.Map
				_ = _442_v25
				_442_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _441_v24)
				var _443_v27 _dafny.Sequence
				_ = _443_v27
				_443_v27 = _dafny.SeqOf(_417_v4)
				var _444_v28 _dafny.MultiSet
				_ = _444_v28
				_444_v28 = _dafny.MultiSetOf(_dafny.IntOfUint32((_443_v27).Cardinality()), _417_v4)
				var _445_v29 _dafny.Sequence
				_ = _445_v29
				_445_v29 = _dafny.SeqOf(p1)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_438_v21), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
				_ = _index77
				var _rhs91 _dafny.Int = _417_v4
				_ = _rhs91
				var _rhs92 _dafny.Map = ((func() _dafny.Map {
					if (_442_v25).Contains((_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool)) {
						return (_442_v25).Get((_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool)).(_dafny.Map)
					}
					return _441_v24
				})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_438_v21, (func() _dafny.Map {
					var _coll11 = _dafny.NewMapBuilder()
					_ = _coll11
					for _iter11 := _dafny.Iterate((_444_v28).Elements()); ; {
						_compr_11, _ok11 := _iter11()
						if !_ok11 {
							break
						}
						var _446_v26 _dafny.Int
						_446_v26 = interface{}(_compr_11).(_dafny.Int)
						if (_444_v28).Contains(_446_v26) {
							_coll11.Add((_446_v26).Minus(_417_v4), _417_v4)
						}
					}
					return _coll11.ToMap()
				}()).Update(_dafny.IntOfInt64(617), _417_v4)))
				_ = _rhs92
				var _rhs93 bool = (_dafny.IntOfUint32((_445_v29).Cardinality())).Cmp(_417_v4) > 0
				_ = _rhs93
				var _rhs94 _dafny.Int = (_417_v4).Plus(_417_v4)
				_ = _rhs94
				var _lhs80 _dafny.Array = _438_v21
				_ = _lhs80
				var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_438_v21), 0))
				_ = _lhs81
				var _lhs82 _dafny.Array = _412_v0
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
				_ = _lhs83
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				(_lhs80).ArraySet1(_rhs91, (_lhs81).Int())
				_441_v24 = _rhs92
				(_lhs82).ArraySet1(_rhs93, (_lhs83).Int())
				_lhs84.F2 = _rhs94
				_422_v9 = (_422_v9).Merge(_422_v9)
				_416_v3 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-795))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_447_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_448_i2 _dafny.Int) _dafny.CodePoint {
						return _447_v1
					}
				})(_414_v1)))
			}
			var _449_v30 _dafny.Map
			_ = _449_v30
			_449_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21(_dafny.UnicodeSeqOfUtf8Bytes("adicjv"), _417_v4, _417_v4, globalState), (_this).F32())
			_449_v30 = _449_v30
			_414_v1 = _414_v1
			var _450_v31 _dafny.Map
			_ = _450_v31
			_450_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_417_v4).Plus(_417_v4), p0)
			var _451_v32 _dafny.MultiSet
			_ = _451_v32
			_451_v32 = _dafny.MultiSetOf(_dafny.CodePoint('v'), _414_v1)
			_450_v31 = (_450_v31).Update(_417_v4, (_451_v32).IsSubsetOf(_451_v32))
		}
		var _452_v33 _dafny.Set
		_ = _452_v33
		_452_v33 = _dafny.SetOf((_this).F32())
		var _453_v34 D1
		_ = _453_v34
		_453_v34 = Companion_D1_.Create_DC4_((_this).Fm14(_452_v33, _dafny.IntOfInt64(614), globalState), (_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool))
		var _454_v35 _dafny.MultiSet
		_ = _454_v35
		_454_v35 = _dafny.MultiSetOf(_453_v34)
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))
		_ = _index78
		(_412_v0).ArraySet1((_454_v35).IsDisjointFrom(_454_v35), (_index78).Int())
		var _pat_let_tv9 = _417_v4
		_ = _pat_let_tv9
		var _pat_let_tv10 = _417_v4
		_ = _pat_let_tv10
		var _pat_let_tv11 = _417_v4
		_ = _pat_let_tv11
		var _rhs95 D1 = Companion_Default___.Fm22(((_452_v33).Cardinality()).Plus(_417_v4), _417_v4, _414_v1, globalState)
		_ = _rhs95
		var _rhs96 bool = true
		_ = _rhs96
		var _rhs97 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_416_v3, _416_v3), p2)
		_ = _rhs97
		var _rhs98 _dafny.Int = func(_source5 D1) _dafny.Int {
			if _source5.Is_DC4() {
				var _455___mcc_h0 bool = _source5.Get_().(D1_DC4).Cf7
				_ = _455___mcc_h0
				var _456___mcc_h1 bool = _source5.Get_().(D1_DC4).Cf8
				_ = _456___mcc_h1
				var _457_cf8 bool = _456___mcc_h1
				_ = _457_cf8
				var _458_cf7 bool = _455___mcc_h0
				_ = _458_cf7
				return _dafny.IntOfInt64(-400)
			} else if _source5.Is_DC3() {
				var _459___mcc_h2 _dafny.CodePoint = _source5.Get_().(D1_DC3).Cf6
				_ = _459___mcc_h2
				var _460_cf6 _dafny.CodePoint = _459___mcc_h2
				_ = _460_cf6
				return Companion_Default___.SafeDivisionInt(_pat_let_tv9, _pat_let_tv10)
			} else {
				var _461___mcc_h3 D1 = _source5.Get_().(D1_DC5).Cf9
				_ = _461___mcc_h3
				var _462_cf9 D1 = _461___mcc_h3
				_ = _462_cf9
				return _pat_let_tv11
			}
		}(_453_v34)
		_ = _rhs98
		var _rhs99 _dafny.Int = _dafny.IntOfInt64(230)
		_ = _rhs99
		var _lhs85 *GlobalState = globalState
		_ = _lhs85
		var _lhs86 *GlobalState = globalState
		_ = _lhs86
		var _lhs87 *GlobalState = globalState
		_ = _lhs87
		_453_v34 = _rhs95
		_lhs85.F1 = _rhs96
		_416_v3 = _rhs97
		_lhs86.F0 = _rhs98
		_lhs87.F0 = _rhs99
		var _463_v36 _dafny.Sequence
		_ = _463_v36
		_463_v36 = _dafny.SeqOf(p0, _dafny.Companion_Sequence_.Contains(_416_v3, _414_v1), (_this).F32(), (_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool))
		_463_v36 = Companion_Default___.Fm23(globalState)
		var _464_v37 _dafny.Map
		_ = _464_v37
		_464_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_412_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_412_v0), 0))).Int()).(bool), _417_v4)
		var _465_v38 _dafny.MultiSet
		_ = _465_v38
		_465_v38 = _dafny.MultiSetOf(_464_v37, _464_v37)
		var _466_v39 _dafny.MultiSet
		_ = _466_v39
		_466_v39 = _465_v38
		r0 = _466_v39
		var _467_v40 _dafny.Map
		_ = _467_v40
		_467_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_417_v4).Plus(_417_v4), (_this).Fm14(_452_v33, _dafny.IntOfUint32((_463_v36).Cardinality()), globalState))
		r1 = _467_v40
		return r0, r1
	}
}
func (_this *C1) M8(p0 _dafny.Array, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _468_v0 _dafny.CodePoint
		_ = _468_v0
		_468_v0 = _dafny.CodePoint('o')
		var _469_v1 _dafny.Map
		_ = _469_v1
		_469_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v0, p1)
		(globalState).F1 = ((_this).F32()) == (!((func() bool {
			if (_469_v1).Contains(_468_v0) {
				return (_469_v1).Get(_468_v0).(bool)
			}
			return (_this).F32()
		})()))
		var _hi3 _dafny.Int = p2
		_ = _hi3
		for _470_i0 := Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfgmwfj")).Cardinality()), p2); _470_i0.Cmp(_hi3) < 0; _470_i0 = _470_i0.Plus(_dafny.One) {
			var _471_v2 _dafny.Sequence
			_ = _471_v2
			_471_v2 = _dafny.SeqOf((_this).F32(), (_this).F32(), p1, p1, p1)
			var _472_v3 _dafny.Set
			_ = _472_v3
			_472_v3 = _dafny.SetOf(p2)
			(globalState).F1 = (_471_v2).Select((Companion_Default___.SafeIndex(((_472_v3).Difference(_472_v3)).Cardinality(), _dafny.IntOfUint32((_471_v2).Cardinality()))).Uint32()).(bool)
			(globalState).F1 = (p2).Cmp(_470_i0) != 0
			r1 = (_this).F32()
			(globalState).F1 = ((Companion_Default___.Fm24((_this).F32(), globalState)).IsSubsetOf(_472_v3)) || (p1)
		}
		if p1 {
			if p1 {
				var _473_v4 _dafny.Sequence
				_ = _473_v4
				_473_v4 = _dafny.UnicodeSeqOfUtf8Bytes("qh")
				var _474_v5 _dafny.Map
				_ = _474_v5
				_474_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Minus(p2), _dafny.Companion_Sequence_.Concatenate(_473_v4, _473_v4))
				var _475_v6 _dafny.Set
				_ = _475_v6
				_475_v6 = _dafny.SetOf((_this).F32(), p1)
				var _rhs100 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality())
				_ = _rhs100
				var _rhs101 bool = (p1) == (p1)
				_ = _rhs101
				var _rhs102 _dafny.Int = (func() _dafny.Int {
					if p1 {
						return Companion_Default___.SafeModuloInt(p2, (_dafny.Zero).Minus(p2))
					}
					return p2
				})()
				_ = _rhs102
				var _rhs103 _dafny.Sequence = (func() _dafny.Sequence {
					if (_474_v5).Contains(p2) {
						return (_474_v5).Get(p2).(_dafny.Sequence)
					}
					return _473_v4
				})()
				_ = _rhs103
				var _rhs104 bool = ((_475_v6).Cardinality()).Cmp(_dafny.IntOfInt64(113)) == 0
				_ = _rhs104
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_lhs88.F0 = _rhs100
				_lhs89.F1 = _rhs101
				_lhs90.F2 = _rhs102
				r0 = _rhs103
				_lhs91.F1 = _rhs104
				(globalState).F2 = (p2).Minus(p2)
				var _476_v7 *C0
				_ = _476_v7
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__(_dafny.Companion_Sequence_.Concatenate(_473_v4, _dafny.UnicodeSeqOfUtf8Bytes("tgruw")), (_this).F32(), p2, p0, !_dafny.Companion_Sequence_.Equal(_473_v4, _473_v4))
				_476_v7 = _nw82
				var _477_v8 _dafny.Map
				_ = _477_v8
				_477_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _478_v9 _dafny.MultiSet
				_ = _478_v9
				_478_v9 = _dafny.MultiSetOf((_this).F32(), (func() bool {
					if (_477_v8).Contains(true) {
						return (_477_v8).Get(true).(bool)
					}
					return (_this).F32()
				})(), (_this).F32(), (_this).F32(), !(!(!(p1))))
				var _rhs105 bool = (func() bool {
					if (_478_v9).IsDisjointFrom(_478_v9) {
						return ((_this).F32()) || (p1)
					}
					return true
				})()
				_ = _rhs105
				var _rhs106 *C0 = _476_v7
				_ = _rhs106
				var _rhs107 _dafny.CodePoint = _468_v0
				_ = _rhs107
				var _rhs108 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_479_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_480_i1 _dafny.Int) _dafny.CodePoint {
						return _479_v0
					}
				})(_468_v0))), _dafny.UnicodeSeqOfUtf8Bytes("h"))
				_ = _rhs108
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				_lhs92.F1 = _rhs105
				_476_v7 = _rhs106
				_468_v0 = _rhs107
				_lhs93.F1 = _rhs108
				(globalState).F0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-248))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_481_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_482_i2 _dafny.Int) _dafny.CodePoint {
						return _481_v0
					}
				})(_468_v0))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-248))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_483_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_484_i2 _dafny.Int) _dafny.CodePoint {
						return _483_v0
					}
				})(_468_v0)))).Cardinality()))).Uint32(), (_476_v7.F35).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_476_v7.F35).Cardinality()))).Uint32()).(_dafny.CodePoint)), (Companion_Default___.SafeIndex((p2).Times(p2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-248))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}((func(_485_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_486_i2 _dafny.Int) _dafny.CodePoint {
						return _485_v0
					}
				})(_468_v0))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-248))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_487_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_488_i2 _dafny.Int) _dafny.CodePoint {
						return _487_v0
					}
				})(_468_v0)))).Cardinality()))).Uint32(), (_476_v7.F35).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_476_v7.F35).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()))).Uint32(), (_476_v7.F35).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(p2)), _dafny.IntOfUint32((_476_v7.F35).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality())
				var _489_v10 _dafny.Array
				_ = _489_v10
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw83
				_489_v10 = _nw83
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_489_v10), 0))
				_ = _index79
				(_489_v10).ArraySet1(p2, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_489_v10), 0))
				_ = _index80
				(_489_v10).ArraySet1(p2, (_index80).Int())
			} else {
				var _490_v11 _dafny.Array
				_ = _490_v11
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw84
				_490_v11 = _nw84
				var _491_v12 _dafny.Set
				_ = _491_v12
				_491_v12 = _dafny.SetOf(_490_v11)
				var _492_v13 _dafny.MultiSet
				_ = _492_v13
				_492_v13 = _dafny.MultiSetOf(p2)
				var _493_v14 _dafny.Sequence
				_ = _493_v14
				_493_v14 = _dafny.UnicodeSeqOfUtf8Bytes("lpv")
				var _494_v15 _dafny.Sequence
				_ = _494_v15
				_494_v15 = _dafny.SeqOf(p1)
				var _495_v16 _dafny.Array
				_ = _495_v16
				var _nwElement0_18 _dafny.MultiSet = _492_v13
				_ = _nwElement0_18
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(8))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_18, 0)
				(_nw85).ArraySet1(_dafny.MultiSetOf(p2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_493_v14).Cardinality())), p2), 1)
				(_nw85).ArraySet1((_492_v13).Difference(_492_v13), 2)
				(_nw85).ArraySet1((Companion_Default___.Fm21(_493_v14, _dafny.IntOfUint32((_494_v15).Cardinality()), p2, globalState)).Union(_492_v13), 3)
				(_nw85).ArraySet1(_492_v13, 4)
				(_nw85).ArraySet1((_492_v13).Intersection(_492_v13), 5)
				(_nw85).ArraySet1(_492_v13, 6)
				(_nw85).ArraySet1(_492_v13, 7)
				_495_v16 = _nw85
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_495_v16), 0))
				_ = _index81
				(_495_v16).ArraySet1(_492_v13, (_index81).Int())
				var _496_v17 _dafny.MultiSet
				_ = _496_v17
				_496_v17 = _492_v13
				var _497_v18 _dafny.Set
				_ = _497_v18
				_497_v18 = _dafny.SetOf(_496_v17, _496_v17, _496_v17, _492_v13)
				var _498_v19 D6
				_ = _498_v19
				_498_v19 = Companion_D6_.Create_DC12_(_491_v12)
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_495_v16), 0))
				_ = _index82
				var _rhs109 _dafny.Int = (_dafny.Zero).Minus(p2)
				_ = _rhs109
				var _rhs110 _dafny.Set = ((_498_v19).Dtor_cf19()).Difference(_491_v12)
				_ = _rhs110
				var _rhs111 _dafny.MultiSet = _492_v13
				_ = _rhs111
				var _rhs112 _dafny.Set = ((_497_v18).Union(_497_v18)).Intersection(_497_v18)
				_ = _rhs112
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				var _lhs95 _dafny.Array = _495_v16
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_495_v16), 0))
				_ = _lhs96
				_lhs94.F2 = _rhs109
				_491_v12 = _rhs110
				(_lhs95).ArraySet1(_rhs111, (_lhs96).Int())
				_497_v18 = _rhs112
				var _499_v20 _dafny.Array
				_ = _499_v20
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_11
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Sequence = (func(_500_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_501_i3 _dafny.Int) _dafny.Sequence {
							return _500_v14
						}
					})(_493_v14)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw86 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw86).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw86).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_499_v20 = _nw86
				_499_v20 = _499_v20
				(globalState).F0 = (_dafny.IntOfInt64(553)).Plus(_dafny.IntOfInt64(740))
				(globalState).F0 = p2
				(globalState).F1 = (_this).F32()
			}
			var _502_v21 D4
			_ = _502_v21
			_502_v21 = Companion_D4_.Create_DC10_(_468_v0, _dafny.CodePoint('k'))
			var _503_v22 D1
			_ = _503_v22
			_503_v22 = Companion_D1_.Create_DC3_(_468_v0)
			var _pat_let_tv12 = _503_v22
			_ = _pat_let_tv12
			_502_v21 = func(_pat_let12_0 D4) D4 {
				return func(_504_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let13_0 _dafny.CodePoint) D4 {
						return func(_505_dt__update_hcf16_h0 _dafny.CodePoint) D4 {
							return Companion_D4_.Create_DC10_(_505_dt__update_hcf16_h0, (_504_dt__update__tmp_h1).Dtor_cf17())
						}(_pat_let13_0)
					}((_pat_let_tv12).Dtor_cf6())
				}(_pat_let12_0)
			}(_502_v21)
			var _506_v23 _dafny.Sequence
			_ = _506_v23
			_506_v23 = _dafny.UnicodeSeqOfUtf8Bytes("u")
			var _507_v24 *C0
			_ = _507_v24
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__(_506_v23, ((_this).F32()) && (p1), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("norpc")).Cardinality()), p0, p1)
			_507_v24 = _nw87
			(globalState).F0 = (_dafny.Zero).Minus(p2)
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((p0), 0))
			_ = _index83
			(p0).ArraySet1((p2).Cmp(p2) <= 0, (_index83).Int())
			var _508_v25 _dafny.Array
			_ = _508_v25
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw88
			_508_v25 = _nw88
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((p0), 0))
			_ = _index84
			var _rhs113 bool = (p1) == ((p2).Cmp(_dafny.IntOfUint32((_506_v23).Cardinality())) != 0)
			_ = _rhs113
			var _rhs114 _dafny.Array = _508_v25
			_ = _rhs114
			var _lhs97 _dafny.Array = p0
			_ = _lhs97
			var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((p0), 0))
			_ = _lhs98
			(_lhs97).ArraySet1(_rhs113, (_lhs98).Int())
			_508_v25 = _rhs114
		} else {
			var _509_v26 _dafny.Array
			_ = _509_v26
			var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw89
			_509_v26 = _nw89
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))
			_ = _index85
			(_509_v26).ArraySet1(p2, (_index85).Int())
			var _510_v27 _dafny.Sequence
			_ = _510_v27
			_510_v27 = _dafny.UnicodeSeqOfUtf8Bytes("scemdf")
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))
			_ = _index86
			var _rhs115 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_510_v27, _dafny.Companion_Sequence_.Concatenate(_510_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(195))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_511_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_512_i4 _dafny.Int) _dafny.CodePoint {
					return _511_v0
				}
			})(_468_v0)))))).Cardinality()))
			_ = _rhs115
			var _rhs116 _dafny.Int = p2
			_ = _rhs116
			var _lhs99 _dafny.Array = _509_v26
			_ = _lhs99
			var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))
			_ = _lhs100
			var _lhs101 *GlobalState = globalState
			_ = _lhs101
			(_lhs99).ArraySet1(_rhs115, (_lhs100).Int())
			_lhs101.F0 = _rhs116
			(globalState).F2 = p2
			if (Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(660))).Cmp((_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int)) <= 0 {
				r1 = false
				var _513_v28 _dafny.Array
				_ = _513_v28
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_12
				var _nw90 _dafny.Array
				_ = _nw90
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw90 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.CodePoint = (func(_514_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_515_i5 _dafny.Int) _dafny.CodePoint {
							return _514_v0
						}
					})(_468_v0)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw90 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw90).ArraySet1CodePoint(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw90).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_513_v28 = _nw90
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_513_v28), 0))
				_ = _index87
				(_513_v28).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_513_v28), 0))
				_ = _index88
				(_513_v28).ArraySet1CodePoint(_468_v0, (_index88).Int())
				var _516_v29 _dafny.Set
				_ = _516_v29
				_516_v29 = _dafny.SetOf((_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int))
				(globalState).F1 = (((_516_v29).Union(_dafny.SetOf((_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int)))).Cardinality()).Cmp((_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int)) > 0
				var _517_v30 _dafny.Array
				_ = _517_v30
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_13
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_518_p1 bool, _519_p2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_520_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_518_p1), _dafny.SeqOf((Companion_D2_.Create_DC7_(_518_p1, (_this).F32(), _519_p2)).Dtor_cf11()))
						}
					})(p1, p2)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw91 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw91).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw91).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_517_v30 = _nw91
				var _521_v31 _dafny.Sequence
				_ = _521_v31
				_521_v31 = _dafny.SeqOf(p1)
				var _522_v32 _dafny.Sequence
				_ = _522_v32
				_522_v32 = _521_v31
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_517_v30), 0))
				_ = _index89
				(_517_v30).ArraySet1((_522_v32), (_index89).Int())
				var _523_v33 D6
				_ = _523_v33
				_523_v33 = Companion_D6_.Create_DC12_(_dafny.SetOf(_509_v26))
				var _524_v34 _dafny.Set
				_ = _524_v34
				_524_v34 = _dafny.SetOf(_509_v26)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_517_v30), 0))
				_ = _index90
				(_517_v30).ArraySet1((func() _dafny.Sequence {
					if (_dafny.MultiSetOf(_523_v33)).Equals(_dafny.MultiSetOf(Companion_D6_.Create_DC12_(_524_v34), _523_v33)) {
						return _dafny.Companion_Sequence_.Concatenate(_521_v31, _521_v31)
					}
					return _dafny.Companion_Sequence_.Concatenate(_521_v31, _dafny.SeqOf((_this).F32(), (_this).F32()))
				})(), (_index90).Int())
				(globalState).F1 = (_this).F32()
			} else {
				var _525_v35 _dafny.Set
				_ = _525_v35
				_525_v35 = _dafny.SetOf(p1)
				var _526_v36 *C0
				_ = _526_v36
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__(_510_v27, (_this).Fm14(_525_v35, (_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int), globalState), (_dafny.Zero).Minus(p2), p0, p1)
				_526_v36 = _nw92
				var _527_v37 _dafny.Set
				_ = _527_v37
				_527_v37 = _dafny.SetOf(_526_v36)
				_527_v37 = _527_v37
				(globalState).F2 = (_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int)
				_525_v35 = _525_v35
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))
				_ = _index91
				(_509_v26).ArraySet1((p2).Times((_dafny.Zero).Minus((_509_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_509_v26), 0))).Int()).(_dafny.Int))), (_index91).Int())
				var _528_v38 *C0
				_ = _528_v38
				var _nw93 *C0 = New_C0_()
				_ = _nw93
				_nw93.Ctor__(_526_v36.F35, (_this).F32(), _dafny.IntOfUint32((_510_v27).Cardinality()), p0, (_526_v36).Fm7(globalState))
				_528_v38 = _nw93
			}
			(globalState).F1 = (_this).F32()
			_509_v26 = _509_v26
		}
		var _529_v39 D0
		_ = _529_v39
		_529_v39 = Companion_D0_.Create_DC0_((_this).F32())
		_529_v39 = _529_v39
		var _530_v40 _dafny.Array
		_ = _530_v40
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_14
		var _nw94 _dafny.Array
		_ = _nw94
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw94 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Map = (func(_531_p1 bool) func(_dafny.Int) _dafny.Map {
				return func(_532_i7 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_p1, (_this).F32())
				}
			})(p1)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw94 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw94).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw94).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_530_v40 = _nw94
		var _533_v41 D8
		_ = _533_v41
		_533_v41 = Companion_D8_.Create_DC16_(_530_v40)
		var _534_v42 _dafny.Map
		_ = _534_v42
		_534_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_533_v41).Dtor_cf27(), !(p1))
		var _535_v43 _dafny.Sequence
		_ = _535_v43
		_535_v43 = _dafny.SeqOf(p1, (_this).F32(), !(false))
		var _536_v44 _dafny.Sequence
		_ = _536_v44
		_536_v44 = _dafny.SeqOf(p2)
		_534_v42 = (_534_v42).Update(_530_v40, (_535_v43).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_536_v44).Cardinality()), _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32()).(bool))
		var _537_v45 D1
		_ = _537_v45
		_537_v45 = Companion_D1_.Create_DC3_(_468_v0)
		var _538_v46 D1
		_ = _538_v46
		_538_v46 = Companion_D1_.Create_DC5_(_537_v45)
		var _source6 D1 = _538_v46
		_ = _source6
		if _source6.Is_DC4() {
			var _539___mcc_h0 bool = _source6.Get_().(D1_DC4).Cf7
			_ = _539___mcc_h0
			var _540___mcc_h1 bool = _source6.Get_().(D1_DC4).Cf8
			_ = _540___mcc_h1
			var _541_cf8 bool = _540___mcc_h1
			_ = _541_cf8
			var _542_cf7 bool = _539___mcc_h0
			_ = _542_cf7
			if (_535_v43).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32()).(bool) {
				var _pat_let_tv13 = p1
				_ = _pat_let_tv13
				var _543_v47 _dafny.Sequence
				_ = _543_v47
				_543_v47 = _dafny.SeqOf(_529_v39, func(_pat_let14_0 D0) D0 {
					return func(_544_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let15_0 bool) D0 {
							return func(_545_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC0_(_545_dt__update_hcf0_h0)
							}(_pat_let15_0)
						}(_pat_let_tv13)
					}(_pat_let14_0)
				}(_529_v39), Companion_D0_.Create_DC0_(_542_cf7))
				_543_v47 = _543_v47
				(globalState).F1 = (_this).F32()
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))
				_ = _index92
				(p0).ArraySet1((_this).F32(), (_index92).Int())
				var _546_v48 _dafny.Set
				_ = _546_v48
				_546_v48 = _dafny.SetOf(p2)
				var _547_v49 _dafny.MultiSet
				_ = _547_v49
				_547_v49 = _dafny.MultiSetOf(p1)
				var _548_v50 D6
				_ = _548_v50
				_548_v50 = Companion_D6_.Create_DC14_(_dafny.IntOfInt64(444), (_546_v48).Cardinality(), _547_v49)
				var _549_v51 _dafny.Map
				_ = _549_v51
				_549_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(p2, Companion_Default___.Fm2(p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_548_v50).Dtor_cf24())).Cardinality(), globalState), globalState), p2)
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))
				_ = _index93
				(p0).ArraySet1(!(!(_549_v51).Equals((_549_v51).Merge(_549_v51))), (_index93).Int())
				var _550_v52 _dafny.Array
				_ = _550_v52
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw95
				_550_v52 = _nw95
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))
				_ = _index94
				(_550_v52).ArraySet1(p2, (_index94).Int())
				var _551_v53 _dafny.Set
				_ = _551_v53
				_551_v53 = _dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))).Int()).(bool), _541_cf8, (_this).F32())
				var _552_v54 _dafny.Set
				_ = _552_v54
				_552_v54 = _dafny.SetOf((_this).F32(), (_this).Fm14(_551_v53, p2, globalState))
				var _553_v55 D0
				_ = _553_v55
				_553_v55 = Companion_D0_.Create_DC1_(p2, _dafny.UnicodeSeqOfUtf8Bytes("jcg"), !((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))).Int()).(bool)), p2, (_this).Fm14(_551_v53, p2, globalState))
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))
				_ = _index95
				var _rhs117 _dafny.CodePoint = Companion_Default___.Fm4(_553_v55, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))).Int()).(bool), p2, globalState)
				_ = _rhs117
				var _rhs118 _dafny.Int = p2
				_ = _rhs118
				var _lhs102 _dafny.Array = _550_v52
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))
				_ = _lhs103
				_468_v0 = _rhs117
				(_lhs102).ArraySet1(_rhs118, (_lhs103).Int())
				var _554_v56 _dafny.Map
				_ = _554_v56
				_554_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _dafny.Companion_Sequence_.IsPrefixOf(_536_v44, _536_v44))
				var _555_v57 _dafny.Array
				_ = _555_v57
				var _nwElement0_19 _dafny.Sequence = _535_v43
				_ = _nwElement0_19
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(20))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_19, 0)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_535_v43, _dafny.Companion_Sequence_.Update(_535_v43, (Companion_Default___.SafeIndex((_550_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32(), p1)), 1)
				(_nw96).ArraySet1(_535_v43, 2)
				(_nw96).ArraySet1(_dafny.SeqOf(!((_this).F32())), 3)
				(_nw96).ArraySet1(_dafny.SeqOf(_542_cf7), 4)
				(_nw96).ArraySet1(_535_v43, 5)
				(_nw96).ArraySet1(_535_v43, 6)
				(_nw96).ArraySet1(_535_v43, 7)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F32(), (_535_v43).Select((Companion_Default___.SafeIndex((_550_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32()).(bool)), (Companion_Default___.SafeIndex((_550_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_this).F32(), (_535_v43).Select((Companion_Default___.SafeIndex((_550_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32()).(bool))).Cardinality()))).Uint32(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((p0), 0))).Int()).(bool)), 8)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_535_v43, _dafny.SeqOf(_542_cf7)), 9)
				(_nw96).ArraySet1(_535_v43, 10)
				(_nw96).ArraySet1(_535_v43, 11)
				(_nw96).ArraySet1(_535_v43, 12)
				(_nw96).ArraySet1(_535_v43, 13)
				(_nw96).ArraySet1(_dafny.SeqOf(true), 14)
				(_nw96).ArraySet1(_535_v43, 15)
				(_nw96).ArraySet1(_535_v43, 16)
				(_nw96).ArraySet1(_535_v43, 17)
				(_nw96).ArraySet1(_535_v43, 18)
				(_nw96).ArraySet1(_535_v43, 19)
				_555_v57 = _nw96
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_555_v57), 0))
				_ = _index96
				(_555_v57).ArraySet1(_535_v43, (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))
				_ = _index97
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_555_v57), 0))
				_ = _index98
				var _rhs119 _dafny.Int = p2
				_ = _rhs119
				var _rhs120 bool = _542_cf7
				_ = _rhs120
				var _rhs121 _dafny.Map = _554_v56
				_ = _rhs121
				var _rhs122 _dafny.Sequence = _535_v43
				_ = _rhs122
				var _lhs104 _dafny.Array = _550_v52
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_550_v52), 0))
				_ = _lhs105
				var _lhs106 _dafny.Array = _555_v57
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_555_v57), 0))
				_ = _lhs107
				(_lhs104).ArraySet1(_rhs119, (_lhs105).Int())
				_542_cf7 = _rhs120
				_554_v56 = _rhs121
				(_lhs106).ArraySet1(_rhs122, (_lhs107).Int())
			} else {
				var _556_v58 _dafny.Map
				_ = _556_v58
				_556_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(609))
				(globalState).F0 = Companion_Default___.SafeDivisionInt(p2, (func() _dafny.Int {
					if (_556_v58).Contains(!((_this).F32())) {
						return (_556_v58).Get(!((_this).F32())).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(p2)
				})())
				var _557_v59 _dafny.Array
				_ = _557_v59
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw97
				_557_v59 = _nw97
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))
				_ = _index99
				(_557_v59).ArraySet1(p0, (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))
				_ = _index100
				(_557_v59).ArraySet1(p0, (_index100).Int())
				(globalState).F0 = (p2).Plus(p2)
				(globalState).F0 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-529), (Companion_Default___.Fm2(p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(625), false)).Cardinality(), globalState)).Times(p2))
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))
				_ = _arr0
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))), 0))
				_ = _index101
				(_arr0).ArraySet1(p1, (_index101).Int())
				var _558_v60 _dafny.Sequence
				_ = _558_v60
				_558_v60 = _dafny.UnicodeSeqOfUtf8Bytes("d")
				var _559_v61 T0
				_ = _559_v61
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__(_dafny.Companion_Sequence_.Concatenate(_558_v60, _558_v60), (_this).F32(), Companion_Default___.SafeModuloInt(p2, p2), _dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int())), true)
				_559_v61 = _nw98
				var _560_v62 *C0
				_ = _560_v62
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__(_558_v60, false, p2, p0, (_this).F32())
				_560_v62 = _nw99
				var _561_v63 D6
				_ = _561_v63
				_561_v63 = Companion_D6_.Create_DC14_((_559_v61).F26(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _535_v43)).Cardinality(), _dafny.MultiSetOf(!(_542_cf7)))
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))
				_ = _arr1
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))), 0))
				_ = _index102
				var _rhs123 bool = ((_561_v63).Dtor_cf23()).Cmp(p2) != 0
				_ = _rhs123
				var _rhs124 _dafny.Int = ((_559_v61).F26()).Minus((_559_v61).F26())
				_ = _rhs124
				var _rhs125 T0 = _559_v61
				_ = _rhs125
				var _rhs126 *C0 = _560_v62
				_ = _rhs126
				var _lhs108 _dafny.Array = _dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_dafny.ArrayCastTo((_557_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_557_v59), 0))).Int()))), 0))
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				(_lhs108).ArraySet1(_rhs123, (_lhs109).Int())
				_lhs110.F0 = _rhs124
				_559_v61 = _rhs125
				_560_v62 = _rhs126
			}
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((p0), 0))
			_ = _index103
			(p0).ArraySet1(_541_cf8, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((p0), 0))
			_ = _index104
			(p0).ArraySet1(true, (_index104).Int())
			var _562_v64 _dafny.Array
			_ = _562_v64
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_15
			var _nw100 _dafny.Array
			_ = _nw100
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw100 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) D1 = (func(_563_v46 D1) func(_dafny.Int) D1 {
					return func(_564_i8 _dafny.Int) D1 {
						return _563_v46
					}
				})(_538_v46)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw100 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw100).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw100).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_562_v64 = _nw100
			_562_v64 = _562_v64
			(globalState).F1 = (_this).F32()
		} else if _source6.Is_DC3() {
			var _565___mcc_h2 _dafny.CodePoint = _source6.Get_().(D1_DC3).Cf6
			_ = _565___mcc_h2
			var _566_cf6 _dafny.CodePoint = _565___mcc_h2
			_ = _566_cf6
			var _567_v65 _dafny.MultiSet
			_ = _567_v65
			var _568_v66 _dafny.Map
			_ = _568_v66
			var _out14 _dafny.MultiSet
			_ = _out14
			var _out15 _dafny.Map
			_ = _out15
			_out14, _out15 = (_this).M7(p1, _529_v39, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(666))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}(func(_569_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), globalState)
			_567_v65 = _out14
			_568_v66 = _out15
			var _570_v67 _dafny.Set
			_ = _570_v67
			_570_v67 = _dafny.SetOf(p1)
			var _571_v68 D4
			_ = _571_v68
			_571_v68 = Companion_D4_.Create_DC9_((_570_v67).Union(_570_v67))
			var _source7 D4 = _571_v68
			_ = _source7
			if _source7.Is_DC10() {
				var _572___mcc_h4 _dafny.CodePoint = _source7.Get_().(D4_DC10).Cf16
				_ = _572___mcc_h4
				var _573___mcc_h5 _dafny.CodePoint = _source7.Get_().(D4_DC10).Cf17
				_ = _573___mcc_h5
				var _574_cf17 _dafny.CodePoint = _573___mcc_h5
				_ = _574_cf17
				var _575_cf16 _dafny.CodePoint = _572___mcc_h4
				_ = _575_cf16
				var _576_v69 _dafny.Array
				_ = _576_v69
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
				_ = _nw101
				_576_v69 = _nw101
				var _577_v70 _dafny.Sequence
				_ = _577_v70
				_577_v70 = _dafny.SeqOf(_576_v69, _576_v69, _576_v69)
				var _578_v71 _dafny.Sequence
				_ = _578_v71
				_578_v71 = _dafny.UnicodeSeqOfUtf8Bytes("cc")
				_576_v69 = (_577_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_578_v71, _578_v71)).Cardinality()), _dafny.IntOfUint32((_577_v70).Cardinality()))).Uint32()).(_dafny.Array)
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((p0), 0))
				_ = _index105
				(p0).ArraySet1(p1, (_index105).Int())
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((p0), 0))
				_ = _index106
				(p0).ArraySet1(true, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((p0), 0))
				_ = _index107
				(p0).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_578_v71, _575_cf16)), (_index107).Int())
				var _579_v72 _dafny.Map
				_ = _579_v72
				_579_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ukvt"), p1)
				_579_v72 = _579_v72
			} else {
				var _580___mcc_h6 _dafny.Set = _source7.Get_().(D4_DC9).Cf15
				_ = _580___mcc_h6
				var _581_cf15 _dafny.Set = _580___mcc_h6
				_ = _581_cf15
				var _582_v73 _dafny.Sequence
				_ = _582_v73
				_582_v73 = _dafny.UnicodeSeqOfUtf8Bytes("qten")
				var _583_v74 *C0
				_ = _583_v74
				var _nw102 *C0 = New_C0_()
				_ = _nw102
				_nw102.Ctor__(_582_v73, !((_this).F32()), p2, p0, true)
				_583_v74 = _nw102
				var _584_v75 _dafny.Sequence
				_ = _584_v75
				_584_v75 = _dafny.SeqOf(_583_v74, _583_v74)
				var _585_v76 _dafny.Map
				_ = _585_v76
				_585_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_535_v43).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_535_v43).Cardinality()))).Uint32()).(bool), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_584_v75, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_584_v75).Cardinality()))).Uint32(), _583_v74), _584_v75))
				_585_v76 = (_585_v76).Update(p1, _584_v75)
				var _586_v77 _dafny.Array
				_ = _586_v77
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_16
				var _nw103 _dafny.Array
				_ = _nw103
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw103 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_587_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_588_i10 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_588_i10, _587_p2)
						}
					})(p2)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw103 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw103).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw103).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_586_v77 = _nw103
				var _589_v78 _dafny.Array
				_ = _589_v78
				var _nwElement0_20 _dafny.Array = _586_v77
				_ = _nwElement0_20
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(4))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_20, 0)
				(_nw104).ArraySet1(_586_v77, 1)
				(_nw104).ArraySet1(_586_v77, 2)
				(_nw104).ArraySet1(_586_v77, 3)
				_589_v78 = _nw104
				_589_v78 = _589_v78
				var _590_v79 _dafny.MultiSet
				_ = _590_v79
				_590_v79 = _dafny.MultiSetOf((_this).F32())
				var _591_v80 D6
				_ = _591_v80
				_591_v80 = Companion_D6_.Create_DC14_((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(63))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_592_cf6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_593_i11 _dafny.Int) _dafny.CodePoint {
						return _592_cf6
					}
				})(_566_cf6))))).Cardinality(), p2, _590_v79)
				var _594_v81 _dafny.Set
				_ = _594_v81
				_594_v81 = _dafny.SetOf(p2)
				var _595_v82 D6
				_ = _595_v82
				_595_v82 = Companion_D6_.Create_DC13_((_591_v80).Dtor_cf23(), _dafny.IntOfInt64(-119), _594_v81)
				var _596_v83 _dafny.Map
				_ = _596_v83
				_596_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), p2)
				_595_v82 = Companion_D6_.Create_DC13_((func() _dafny.Int {
					if (_596_v83).Contains(!(!(true))) {
						return (_596_v83).Get(!(!(true))).(_dafny.Int)
					}
					return _dafny.IntOfInt64(682)
				})(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_535_v43).Cardinality()), p2), _594_v81)
				(globalState).F2 = p2
			}
			var _597_v84 _dafny.MultiSet
			_ = _597_v84
			_597_v84 = _dafny.MultiSetOf((_this).F32(), (_this).F32(), !(p1))
			_597_v84 = (_dafny.MultiSetOf(p1)).Difference(_597_v84)
			if (_this).F32() {
				_536_v44 = _dafny.Companion_Sequence_.Concatenate(_536_v44, _536_v44)
				r0 = Companion_Default___.Fm1(globalState)
				var _598_v85 _dafny.Array
				_ = _598_v85
				var _nw105 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(23))
				_ = _nw105
				_598_v85 = _nw105
				var _599_v86 D0
				_ = _599_v86
				_599_v86 = Companion_D0_.Create_DC2_()
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_598_v85), 0))
				_ = _index108
				(_598_v85).ArraySet1(_599_v86, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_598_v85), 0))
				_ = _index109
				(_598_v85).ArraySet1(Companion_D0_.Create_DC2_(), (_index109).Int())
				var _600_v87 _dafny.Map
				_ = _600_v87
				_600_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(p2))
				var _601_v88 _dafny.MultiSet
				_ = _601_v88
				_601_v88 = _dafny.MultiSetOf(_600_v87)
				var _rhs127 bool = p1
				_ = _rhs127
				var _rhs128 _dafny.MultiSet = (_601_v88).Difference(_601_v88)
				_ = _rhs128
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				_lhs111.F1 = _rhs127
				_567_v65 = _rhs128
				var _602_v89 _dafny.Sequence
				_ = _602_v89
				_602_v89 = _dafny.UnicodeSeqOfUtf8Bytes("kbsu")
				var _603_v90 *C0
				_ = _603_v90
				var _nw106 *C0 = New_C0_()
				_ = _nw106
				_nw106.Ctor__(_602_v89, p1, p2, p0, !(p1))
				_603_v90 = _nw106
			} else {
				var _604_v91 _dafny.Sequence
				_ = _604_v91
				_604_v91 = _dafny.UnicodeSeqOfUtf8Bytes("doq")
				var _605_v92 _dafny.Array
				_ = _605_v92
				var _nwElement0_21 _dafny.Sequence = _604_v91
				_ = _nwElement0_21
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(9))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_21, 0)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xqocpbbe"), _604_v91), 1)
				(_nw107).ArraySet1(_604_v91, 2)
				(_nw107).ArraySet1(_604_v91, 3)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_604_v91, _604_v91), 4)
				(_nw107).ArraySet1(_604_v91, 5)
				(_nw107).ArraySet1(_604_v91, 6)
				(_nw107).ArraySet1(_604_v91, 7)
				(_nw107).ArraySet1(_604_v91, 8)
				_605_v92 = _nw107
				_605_v92 = _605_v92
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_605_v92), 0))
				_ = _index110
				(_605_v92).ArraySet1(_604_v91, (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_605_v92), 0))
				_ = _index111
				var _rhs129 bool = ((p2).Plus(_dafny.IntOfInt64(724))).Cmp(p2) != 0
				_ = _rhs129
				var _rhs130 _dafny.Sequence = _604_v91
				_ = _rhs130
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				var _lhs113 _dafny.Array = _605_v92
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_605_v92), 0))
				_ = _lhs114
				_lhs112.F1 = _rhs129
				(_lhs113).ArraySet1(_rhs130, (_lhs114).Int())
				r1 = (((_570_v67).Cardinality()).Cmp(p2) > 0) == (p1)
				var _606_v93 _dafny.Array
				_ = _606_v93
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_17
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = func(_607_i12 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_607_i12, _dafny.IntOfInt64(-966))
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw108 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw108).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw108).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_606_v93 = _nw108
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _index112
				(_606_v93).ArraySet1((_dafny.Zero).Minus(p2), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _index113
				(_606_v93).ArraySet1((p2).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(478))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_608_cf6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_609_i13 _dafny.Int) _dafny.CodePoint {
						return _608_cf6
					}
				})(_566_cf6)))).Cardinality())), (_index113).Int())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _index115
				var _rhs131 _dafny.Int = _dafny.IntOfInt64(-988)
				_ = _rhs131
				var _rhs132 _dafny.Int = p2
				_ = _rhs132
				var _lhs115 _dafny.Array = _606_v93
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _lhs116
				var _lhs117 _dafny.Array = _606_v93
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_606_v93), 0))
				_ = _lhs118
				(_lhs115).ArraySet1(_rhs131, (_lhs116).Int())
				(_lhs117).ArraySet1(_rhs132, (_lhs118).Int())
			}
		} else {
			var _610___mcc_h3 D1 = _source6.Get_().(D1_DC5).Cf9
			_ = _610___mcc_h3
			var _611_cf9 D1 = _610___mcc_h3
			_ = _611_cf9
			var _612_v94 _dafny.Sequence
			_ = _612_v94
			_612_v94 = _dafny.UnicodeSeqOfUtf8Bytes("rmndt")
			var _613_v95 *C0
			_ = _613_v95
			var _nw109 *C0 = New_C0_()
			_ = _nw109
			_nw109.Ctor__(_612_v94, (_this).F32(), p2, p0, p1)
			_613_v95 = _nw109
			var _614_v96 _dafny.Array
			_ = _614_v96
			var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw110
			_614_v96 = _nw110
			var _615_v97 D4
			_ = _615_v97
			_615_v97 = Companion_D4_.Create_DC10_(_dafny.CodePoint('o'), _dafny.CodePoint('p'))
			var _616_v98 _dafny.Array
			_ = _616_v98
			var _nwElement0_22 _dafny.CodePoint = _468_v0
			_ = _nwElement0_22
			var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(26))
			_ = _nw111
			(_nw111).ArraySet1CodePoint(_nwElement0_22, 0)
			(_nw111).ArraySet1CodePoint(_468_v0, 1)
			(_nw111).ArraySet1CodePoint(_dafny.CodePoint('i'), 2)
			(_nw111).ArraySet1CodePoint(_468_v0, 3)
			(_nw111).ArraySet1CodePoint(_468_v0, 4)
			(_nw111).ArraySet1CodePoint(_468_v0, 5)
			(_nw111).ArraySet1CodePoint((_615_v97).Dtor_cf16(), 6)
			(_nw111).ArraySet1CodePoint(_468_v0, 7)
			(_nw111).ArraySet1CodePoint(_468_v0, 8)
			(_nw111).ArraySet1CodePoint(_468_v0, 9)
			(_nw111).ArraySet1CodePoint(_dafny.CodePoint('b'), 10)
			(_nw111).ArraySet1CodePoint(_468_v0, 11)
			(_nw111).ArraySet1CodePoint(_dafny.CodePoint('s'), 12)
			(_nw111).ArraySet1CodePoint(_468_v0, 13)
			(_nw111).ArraySet1CodePoint(_dafny.CodePoint('o'), 14)
			(_nw111).ArraySet1CodePoint(_dafny.CodePoint('x'), 15)
			(_nw111).ArraySet1CodePoint(_468_v0, 16)
			(_nw111).ArraySet1CodePoint((_613_v95.F35).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_613_v95.F35).Cardinality()))).Uint32()).(_dafny.CodePoint), 17)
			(_nw111).ArraySet1CodePoint(_468_v0, 18)
			(_nw111).ArraySet1CodePoint(_468_v0, 19)
			(_nw111).ArraySet1CodePoint(_468_v0, 20)
			(_nw111).ArraySet1CodePoint(_468_v0, 21)
			(_nw111).ArraySet1CodePoint(_468_v0, 22)
			(_nw111).ArraySet1CodePoint(_468_v0, 23)
			(_nw111).ArraySet1CodePoint(_468_v0, 24)
			(_nw111).ArraySet1CodePoint(_468_v0, 25)
			_616_v98 = _nw111
			var _617_v99 _dafny.Map
			_ = _617_v99
			_617_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _616_v98)
			var _618_v100 _dafny.Map
			_ = _618_v100
			_618_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_617_v99, _614_v96)
			_614_v96 = (func() _dafny.Array {
				if (_618_v100).Contains((_617_v99).Merge(_617_v99)) {
					return (_618_v100).Get((_617_v99).Merge(_617_v99)).(_dafny.Array)
				}
				return _614_v96
			})()
			var _619_v101 _dafny.Array
			_ = _619_v101
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw112
			_619_v101 = _nw112
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_619_v101), 0))
			_ = _index116
			(_619_v101).ArraySet1(p2, (_index116).Int())
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((p0), 0))
			_ = _index117
			(p0).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_535_v43, _dafny.SeqOf(p1, p1)), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_619_v101), 0))
			_ = _index118
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((p0), 0))
			_ = _index119
			var _rhs133 _dafny.Int = (p2).Times((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_613_v95.F35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.IntOfUint32((_613_v95.F35).Cardinality()))).Uint32(), _468_v0)).Cardinality())).Times(p2))
			_ = _rhs133
			var _rhs134 bool = false
			_ = _rhs134
			var _lhs119 _dafny.Array = _619_v101
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_619_v101), 0))
			_ = _lhs120
			var _lhs121 _dafny.Array = p0
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((p0), 0))
			_ = _lhs122
			(_lhs119).ArraySet1(_rhs133, (_lhs120).Int())
			(_lhs121).ArraySet1(_rhs134, (_lhs122).Int())
			(globalState).F1 = ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((p0), 0))).Int()).(bool)) == (p1)
		}
		var _620_v102 _dafny.Sequence
		_ = _620_v102
		_620_v102 = _dafny.UnicodeSeqOfUtf8Bytes("km")
		r0 = _620_v102
		r1 = p1
		return r0, r1
	}
}
func (_this *C1) F32() bool {
	{
		return _this._f32
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f25 bool
	_f26 _dafny.Int
	_f31 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f25 = false
	_this._f26 = _dafny.Zero
	_this._f31 = _dafny.Zero
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

func (_this *C2) F25() bool {
	return _this._f25
}
func (_this *C2) F26() _dafny.Int {
	return _this._f26
}
func (_this *C2) Ctor__(f31 _dafny.Int, f25 bool, f26 _dafny.Int) {
	{
		(_this)._f31 = f31
		(_this)._f25 = f25
		(_this)._f26 = f26
	}
}
func (_this *C2) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F31()
	}
}
func (_this *C2) Fm7(globalState *GlobalState) bool {
	{
		return (Companion_D0_.Create_DC0_((_this).F25())).Dtor_cf0()
	}
}
func (_this *C2) Fm11(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(func(_source8 D0) _dafny.Int {
			if _source8.Is_DC1() {
				var _621___mcc_h0 _dafny.Int = _source8.Get_().(D0_DC1).Cf1
				_ = _621___mcc_h0
				var _622___mcc_h1 _dafny.Sequence = _source8.Get_().(D0_DC1).Cf2
				_ = _622___mcc_h1
				var _623___mcc_h2 bool = _source8.Get_().(D0_DC1).Cf3
				_ = _623___mcc_h2
				var _624___mcc_h3 _dafny.Int = _source8.Get_().(D0_DC1).Cf4
				_ = _624___mcc_h3
				var _625___mcc_h4 bool = _source8.Get_().(D0_DC1).Cf5
				_ = _625___mcc_h4
				var _626_cf5 bool = _625___mcc_h4
				_ = _626_cf5
				var _627_cf4 _dafny.Int = _624___mcc_h3
				_ = _627_cf4
				var _628_cf3 bool = _623___mcc_h2
				_ = _628_cf3
				var _629_cf2 _dafny.Sequence = _622___mcc_h1
				_ = _629_cf2
				var _630_cf1 _dafny.Int = _621___mcc_h0
				_ = _630_cf1
				return _627_cf4
			} else if _source8.Is_DC2() {
				return Companion_Default___.SafeModuloInt((_this).F26(), (_this).F31())
			} else {
				var _631___mcc_h5 bool = _source8.Get_().(D0_DC0).Cf0
				_ = _631___mcc_h5
				var _632_cf0 bool = _631___mcc_h5
				_ = _632_cf0
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(523), (_this).F26())
			}
		}(Companion_D0_.Create_DC1_((_this).F31(), _dafny.UnicodeSeqOfUtf8Bytes("lyagi"), (_this).F25(), _dafny.IntOfInt64(-991), (_this).F25())))
	}
}
func (_this *C2) M1(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _hi4 _dafny.Int = (_this).F31()
		_ = _hi4
		for _633_i0 := (_this).F31(); _633_i0.Cmp(_hi4) < 0; _633_i0 = _633_i0.Plus(_dafny.One) {
			var _634_v0 _dafny.Map
			_ = _634_v0
			_634_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			_634_v0 = (_634_v0).Update((_this).F25(), p2)
			var _635_v1 D0
			_ = _635_v1
			_635_v1 = Companion_D0_.Create_DC2_()
			var _636_v2 _dafny.Map
			_ = _636_v2
			_636_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(64), p1)
			var _637_v3 _dafny.Sequence
			_ = _637_v3
			_637_v3 = _dafny.UnicodeSeqOfUtf8Bytes("dxrsfndo")
			var _638_v4 _dafny.Map
			_ = _638_v4
			_638_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(154), _dafny.IntOfUint32((_637_v3).Cardinality()))
			var _639_v5 _dafny.Set
			_ = _639_v5
			_639_v5 = _dafny.SetOf((p0).Cardinality())
			var _640_v6 D2
			_ = _640_v6
			_640_v6 = Companion_D2_.Create_DC6_(_dafny.SetOf((_this).F26()))
			var _rhs135 bool = (_this).F25()
			_ = _rhs135
			var _rhs136 D0 = _635_v1
			_ = _rhs136
			var _rhs137 _dafny.Int = (_this).Fm6(p2, (func() bool {
				if (_636_v2).Contains((func() _dafny.Int {
					if (_638_v4).Contains((_this).F26()) {
						return (_638_v4).Get((_this).F26()).(_dafny.Int)
					}
					return _633_i0
				})()) {
					return (_636_v2).Get((func() _dafny.Int {
						if (_638_v4).Contains((_this).F26()) {
							return (_638_v4).Get((_this).F26()).(_dafny.Int)
						}
						return _633_i0
					})()).(bool)
				}
				return p2
			})(), Companion_Default___.SafeDivisionInt((_this).F26(), _633_i0), globalState)
			_ = _rhs137
			var _rhs138 bool = (_639_v5).Equals(((_640_v6).Dtor_cf10()).Difference(_639_v5))
			_ = _rhs138
			var _lhs123 *GlobalState = globalState
			_ = _lhs123
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			var _lhs125 *GlobalState = globalState
			_ = _lhs125
			_lhs123.F1 = _rhs135
			_635_v1 = _rhs136
			_lhs124.F0 = _rhs137
			_lhs125.F1 = _rhs138
			var _641_v7 D0
			_ = _641_v7
			_641_v7 = Companion_D0_.Create_DC0_(true)
			var _642_v8 _dafny.Map
			_ = _642_v8
			_642_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_641_v7, (_this).F26())
			(globalState).F2 = (func() _dafny.Int {
				if !(!(Companion_Default___.Fm5(p1, _642_v8, (func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-118))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}(func(_643_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _644_v9 _dafny.CodePoint
						_644_v9 = interface{}(_compr_12).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-118))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg41 _dafny.Int) interface{} {
								return coer41(arg41)
							}
						}(func(_643_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						})), _644_v9) {
							_coll12.Add(_644_v9)
						}
					}
					return _coll12.ToSet()
				}()).Cardinality(), globalState))) {
					return _633_i0
				}
				return (_this).F26()
			})()
			var _645_v10 _dafny.Map
			_ = _645_v10
			_645_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _637_v3)
			_645_v10 = (_645_v10).Update(Companion_Default___.Fm5(p1, _642_v8, (_this).F31(), globalState), _637_v3)
		}
		(globalState).F0 = (_this).F31()
		var _646_v11 _dafny.Sequence
		_ = _646_v11
		_646_v11 = _dafny.UnicodeSeqOfUtf8Bytes("mt")
		var _647_v12 D1
		_ = _647_v12
		_647_v12 = Companion_D1_.Create_DC4_(p1, p1)
		var _pat_let_tv14 = _647_v12
		_ = _pat_let_tv14
		var _pat_let_tv15 = p1
		_ = _pat_let_tv15
		var _pat_let_tv16 = p2
		_ = _pat_let_tv16
		var _pat_let_tv17 = p2
		_ = _pat_let_tv17
		var _pat_let_tv18 = p1
		_ = _pat_let_tv18
		var _pat_let_tv19 = p1
		_ = _pat_let_tv19
		var _rhs139 bool = p1
		_ = _rhs139
		var _rhs140 _dafny.Sequence = Companion_Default___.Fm1(globalState)
		_ = _rhs140
		var _rhs141 bool = func(_source9 D1) bool {
			if _source9.Is_DC4() {
				var _648___mcc_h0 bool = _source9.Get_().(D1_DC4).Cf7
				_ = _648___mcc_h0
				var _649___mcc_h1 bool = _source9.Get_().(D1_DC4).Cf8
				_ = _649___mcc_h1
				var _650_cf8 bool = _649___mcc_h1
				_ = _650_cf8
				var _651_cf7 bool = _648___mcc_h0
				_ = _651_cf7
				return (_pat_let_tv14).Dtor_cf8()
			} else if _source9.Is_DC3() {
				var _652___mcc_h2 _dafny.CodePoint = _source9.Get_().(D1_DC3).Cf6
				_ = _652___mcc_h2
				var _653_cf6 _dafny.CodePoint = _652___mcc_h2
				_ = _653_cf6
				return (_this).F25()
			} else {
				var _654___mcc_h3 D1 = _source9.Get_().(D1_DC5).Cf9
				_ = _654___mcc_h3
				var _655_cf9 D1 = _654___mcc_h3
				_ = _655_cf9
				return (_dafny.MultiSetOf((_this).F25())).IsProperSubsetOf(_dafny.MultiSetOf(_pat_let_tv15, _pat_let_tv16, _pat_let_tv17, _pat_let_tv18, _pat_let_tv19))
			}
		}(_647_v12)
		_ = _rhs141
		var _rhs142 bool = !(!(!(p1)))
		_ = _rhs142
		var _rhs143 bool = ((_this).F31()).Cmp((_this).F26()) < 0
		_ = _rhs143
		var _lhs126 *GlobalState = globalState
		_ = _lhs126
		var _lhs127 *GlobalState = globalState
		_ = _lhs127
		r0 = _rhs139
		_646_v11 = _rhs140
		_lhs126.F1 = _rhs141
		_lhs127.F1 = _rhs142
		r0 = _rhs143
		var _hi5 _dafny.Int = ((_this).F31()).Times((_this).F26())
		_ = _hi5
		for _656_i2 := Companion_Default___.SafeModuloInt((_this).F26(), (_this).F31()); _656_i2.Cmp(_hi5) < 0; _656_i2 = _656_i2.Plus(_dafny.One) {
			var _657_v13 _dafny.Map
			_ = _657_v13
			_657_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p2)
			var _658_v14 _dafny.Map
			_ = _658_v14
			_658_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v13, true)
			var _659_v17 D0
			_ = _659_v17
			_659_v17 = Companion_D0_.Create_DC0_(p1)
			var _660_v18 D0
			_ = _660_v18
			_660_v18 = Companion_D0_.Create_DC1_(_656_i2, _646_v11, p1, (_this).F26(), (func() bool {
				if (_657_v13).Contains(_656_i2) {
					return (_657_v13).Get(_656_i2).(bool)
				}
				return Companion_Default___.Fm5((_this).F25(), func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter13 := _dafny.Iterate((func() _dafny.Map {
						var _coll14 = _dafny.NewMapBuilder()
						_ = _coll14
						for _iter14 := _dafny.Iterate((_dafny.SeqOf(_659_v17, _659_v17)).Elements()); ; {
							_compr_14, _ok14 := _iter14()
							if !_ok14 {
								break
							}
							var _661_v16 D0
							_661_v16 = interface{}(_compr_14).(D0)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_659_v17, _659_v17), _661_v16) {
								_coll14.Add(_661_v16, (_this).F25())
							}
						}
						return _coll14.ToMap()
					}()).Keys().Elements()); ; {
						_compr_13, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _662_v15 D0
						_662_v15 = interface{}(_compr_13).(D0)
						if (func() _dafny.Map {
							var _coll15 = _dafny.NewMapBuilder()
							_ = _coll15
							for _iter15 := _dafny.Iterate((_dafny.SeqOf(_659_v17, _659_v17)).Elements()); ; {
								_compr_15, _ok15 := _iter15()
								if !_ok15 {
									break
								}
								var _661_v16 D0
								_661_v16 = interface{}(_compr_15).(D0)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_659_v17, _659_v17), _661_v16) {
									_coll15.Add(_661_v16, (_this).F25())
								}
							}
							return _coll15.ToMap()
						}()).Contains(_662_v15) {
							_coll13.Add(_662_v15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_656_i2, (_this).F31())).Cardinality())
						}
					}
					return _coll13.ToMap()
				}(), _dafny.IntOfUint32((_646_v11).Cardinality()), globalState)
			})())
			var _663_v19 _dafny.Map
			_ = _663_v19
			_663_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_658_v14).Cardinality(), (_dafny.Zero).Minus((_660_v18).Dtor_cf4()))
			var _664_v20 _dafny.Map
			_ = _664_v20
			_664_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_646_v11, _646_v11)
			_663_v19 = (_663_v19).Update((_this).F26(), (_this).Fm11(_664_v20, (_this).F31(), (_this).F31(), globalState))
			(globalState).F0 = ((_dafny.IntOfInt64(281)).Times((_this).F31())).Times((_this).F31())
			var _665_v21 _dafny.Map
			_ = _665_v21
			_665_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((p0).Cardinality()).Cmp(_dafny.IntOfUint32((_646_v11).Cardinality())) != 0)
			_665_v21 = (_665_v21).Update(p2, !(p1))
			_663_v19 = (_663_v19).Update(((_665_v21).Cardinality()).Minus((_this).F31()), (_656_i2).Times((_this).F31()))
		}
		var _666_v22 _dafny.Sequence
		_ = _666_v22
		_666_v22 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F31()), _dafny.IntOfInt64(537))
		var _667_v23 _dafny.Map
		_ = _667_v23
		_667_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_666_v22).Cardinality()))
		var _668_v24 _dafny.Set
		_ = _668_v24
		_668_v24 = _dafny.SetOf(Companion_D2_.Create_DC6_(Companion_Default___.Fm12((_667_v23).Cardinality(), (_this).F31(), globalState)))
		var _669_v25 _dafny.Set
		_ = _669_v25
		_669_v25 = _dafny.SetOf((_this).F31())
		var _670_v26 D2
		_ = _670_v26
		_670_v26 = Companion_D2_.Create_DC6_(_669_v25)
		_668_v24 = (_668_v24).Difference(_dafny.SetOf(_670_v26, Companion_D2_.Create_DC6_(_dafny.SetOf((_this).F26(), _dafny.IntOfInt64(969), (_this).F31(), (_this).F31())), _670_v26))
		var _hi6 _dafny.Int = ((_this).F26()).Plus((_this).F26())
		_ = _hi6
		for _671_i3 := (_this).F26(); _671_i3.Cmp(_hi6) < 0; _671_i3 = _671_i3.Plus(_dafny.One) {
			(globalState).F0 = _dafny.IntOfInt64(-20)
			var _672_v27 _dafny.Array
			_ = _672_v27
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
			_ = _nw113
			_672_v27 = _nw113
			var _673_v28 _dafny.Sequence
			_ = _673_v28
			_673_v28 = _dafny.SeqOf((_this).F25())
			var _674_v29 _dafny.Set
			_ = _674_v29
			_674_v29 = _dafny.SetOf(_673_v28, _673_v28)
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_672_v27), 0))
			_ = _index120
			(_672_v27).ArraySet1(_674_v29, (_index120).Int())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_672_v27), 0))
			_ = _index121
			(_672_v27).ArraySet1(Companion_Default___.Fm13(p2, globalState), (_index121).Int())
			var _675_v30 D0
			_ = _675_v30
			_675_v30 = Companion_D0_.Create_DC2_()
			var _source10 D0 = _675_v30
			_ = _source10
			if _source10.Is_DC1() {
				var _676___mcc_h4 _dafny.Int = _source10.Get_().(D0_DC1).Cf1
				_ = _676___mcc_h4
				var _677___mcc_h5 _dafny.Sequence = _source10.Get_().(D0_DC1).Cf2
				_ = _677___mcc_h5
				var _678___mcc_h6 bool = _source10.Get_().(D0_DC1).Cf3
				_ = _678___mcc_h6
				var _679___mcc_h7 _dafny.Int = _source10.Get_().(D0_DC1).Cf4
				_ = _679___mcc_h7
				var _680___mcc_h8 bool = _source10.Get_().(D0_DC1).Cf5
				_ = _680___mcc_h8
				var _681_cf5 bool = _680___mcc_h8
				_ = _681_cf5
				var _682_cf4 _dafny.Int = _679___mcc_h7
				_ = _682_cf4
				var _683_cf3 bool = _678___mcc_h6
				_ = _683_cf3
				var _684_cf2 _dafny.Sequence = _677___mcc_h5
				_ = _684_cf2
				var _685_cf1 _dafny.Int = _676___mcc_h4
				_ = _685_cf1
				var _686_v31 _dafny.Array
				_ = _686_v31
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
				_ = _nw114
				_686_v31 = _nw114
				var _687_v32 _dafny.MultiSet
				_ = _687_v32
				_687_v32 = _dafny.MultiSetOf(_667_v23)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_686_v31), 0))
				_ = _index122
				(_686_v31).ArraySet1((_dafny.MultiSetOf(_667_v23)).Union(_687_v32), (_index122).Int())
				var _688_v33 _dafny.MultiSet
				_ = _688_v33
				_688_v33 = _687_v32
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_686_v31), 0))
				_ = _index123
				(_686_v31).ArraySet1((_688_v33), (_index123).Int())
				var _689_v34 D2
				_ = _689_v34
				_689_v34 = Companion_D2_.Create_DC7_((_this).F25(), p2, (_this).F31())
				var _690_v35 _dafny.Array
				_ = _690_v35
				var _nwElement0_23 bool = _681_cf5
				_ = _nwElement0_23
				var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(7))
				_ = _nw115
				(_nw115).ArraySet1(_nwElement0_23, 0)
				(_nw115).ArraySet1(_681_cf5, 1)
				(_nw115).ArraySet1(true, 2)
				(_nw115).ArraySet1((_681_cf5) && (p2), 3)
				(_nw115).ArraySet1((_689_v34).Dtor_cf12(), 4)
				(_nw115).ArraySet1(_681_cf5, 5)
				(_nw115).ArraySet1(p2, 6)
				_690_v35 = _nw115
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_690_v35), 0))
				_ = _index124
				(_690_v35).ArraySet1(p1, (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_690_v35), 0))
				_ = _index125
				(_690_v35).ArraySet1(p1, (_index125).Int())
				var _691_v36 _dafny.CodePoint
				_ = _691_v36
				_691_v36 = _dafny.CodePoint('c')
				_691_v36 = _691_v36
				var _692_v37 _dafny.Map
				_ = _692_v37
				_692_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_682_cf4, _685_cf1)
				var _693_v38 _dafny.Map
				_ = _693_v38
				_693_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_692_v37, _690_v35)
				var _rhs144 bool = !(_693_v38).Equals((_693_v38).Merge(_693_v38))
				_ = _rhs144
				var _rhs145 _dafny.Set = _dafny.SetOf(_685_cf1, (_this).F26(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(960), _682_cf4))
				_ = _rhs145
				var _rhs146 _dafny.Int = (func() _dafny.Int {
					if (_690_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_690_v35), 0))).Int()).(bool) {
						return (_this).F26()
					}
					return ((_this).F31()).Minus(_682_cf4)
				})()
				_ = _rhs146
				var _rhs147 _dafny.Int = (_this).F31()
				_ = _rhs147
				var _rhs148 bool = false
				_ = _rhs148
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_lhs128.F1 = _rhs144
				_669_v25 = _rhs145
				_lhs129.F2 = _rhs146
				_lhs130.F0 = _rhs147
				_683_cf3 = _rhs148
			} else if _source10.Is_DC2() {
				var _694_v39 _dafny.Sequence
				_ = _694_v39
				_694_v39 = _dafny.SeqOf(_646_v11)
				var _695_v40 _dafny.Map
				_ = _695_v40
				_695_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F31())
				_666_v22 = Companion_Default___.Fm3(_671_i3, !(((_dafny.Zero).Minus((_this).F31())).Cmp(_dafny.IntOfUint32((_694_v39).Cardinality())) == 0), _dafny.IntOfInt64(766), !(_695_v40).Equals(func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(914), _dafny.IntOfInt64(-407))); ; {
						_compr_16, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _696_v41 _dafny.Int
						_696_v41 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(914)).Cmp(_696_v41) <= 0) && ((_696_v41).Cmp(_dafny.IntOfInt64(-407)) < 0) {
							_coll16.Add((_696_v41).Plus((_this).F31()), (_dafny.Zero).Minus((_this).F31()))
						}
					}
					return _coll16.ToMap()
				}()), globalState)
				(globalState).F1 = (_this).Fm7(globalState)
				var _697_v42 *C1
				_ = _697_v42
				var _nw116 *C1 = New_C1_()
				_ = _nw116
				_nw116.Ctor__(p2)
				_697_v42 = _nw116
				(globalState).F1 = (_673_v28).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_673_v28).Cardinality()))).Uint32()).(bool)
			} else {
				var _698___mcc_h9 bool = _source10.Get_().(D0_DC0).Cf0
				_ = _698___mcc_h9
				var _699_cf0 bool = _698___mcc_h9
				_ = _699_cf0
				var _700_v43 D0
				_ = _700_v43
				_700_v43 = Companion_D0_.Create_DC0_(_699_cf0)
				var _701_v44 _dafny.Array
				_ = _701_v44
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw117
				_701_v44 = _nw117
				var _702_v45 *C0
				_ = _702_v45
				var _nw118 *C0 = New_C0_()
				_ = _nw118
				_nw118.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("me"), (_700_v43).Dtor_cf0(), _671_i3, _701_v44, false)
				_702_v45 = _nw118
				var _703_v46 *C1
				_ = _703_v46
				var _nw119 *C1 = New_C1_()
				_ = _nw119
				_nw119.Ctor__((func() bool {
					if p1 {
						return p1
					}
					return !(p2)
				})())
				_703_v46 = _nw119
				_673_v28 = _673_v28
				var _704_v47 _dafny.Sequence
				_ = _704_v47
				_704_v47 = _673_v28
				var _705_v48 _dafny.Sequence
				_ = _705_v48
				_705_v48 = _dafny.SeqOf(_673_v28, _704_v47, _704_v47, _704_v47, _704_v47)
				var _706_v49 _dafny.Sequence
				_ = _706_v49
				_706_v49 = _dafny.SeqOf(_701_v44)
				(globalState).F1 = _dafny.Companion_Sequence_.Equal(_705_v48, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_704_v47, _704_v47, _704_v47, _704_v47), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_706_v49).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_704_v47, _704_v47, _704_v47, _704_v47)).Cardinality()))).Uint32(), _dafny.SeqOf(p1)), _705_v48))
			}
			var _707_v50 _dafny.MultiSet
			_ = _707_v50
			_707_v50 = _dafny.MultiSetOf(!(p1), true)
			var _708_v51 _dafny.Map
			_ = _708_v51
			_708_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(p2), (_707_v50).Cardinality())
			var _709_v52 _dafny.Sequence
			_ = _709_v52
			_709_v52 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_710_v11 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_711_i4 _dafny.Int) _dafny.CodePoint {
					return (_710_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ferpotds")).Cardinality()), _dafny.IntOfUint32((_710_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_646_v11))))
			if (Companion_Default___.Fm5(p1, _708_v51, (_this).F31(), globalState)) && ((_dafny.IntOfInt64(-875)).Cmp(_dafny.IntOfUint32(((_709_v52).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_709_v52).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) > 0) {
				var _712_v53 _dafny.Map
				_ = _712_v53
				_712_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_671_i3, (_673_v28).Select((Companion_Default___.SafeIndex(_671_i3, _dafny.IntOfUint32((_673_v28).Cardinality()))).Uint32()).(bool))
				var _713_v54 _dafny.CodePoint
				_ = _713_v54
				_713_v54 = _dafny.CodePoint('r')
				var _714_v55 _dafny.Array
				_ = _714_v55
				var _nwElement0_24 _dafny.Int = (_dafny.Zero).Minus((_this).F26())
				_ = _nwElement0_24
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_24, 0)
				(_nw120).ArraySet1((_this).F31(), 1)
				(_nw120).ArraySet1((_this).F26(), 2)
				(_nw120).ArraySet1((((_712_v53).Update((_this).F31(), (_this).F25())).Merge(_712_v53)).Cardinality(), 3)
				(_nw120).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F26()), (_this).F26()), 4)
				(_nw120).ArraySet1((func() _dafny.Int {
					if (p0).Contains(_dafny.IntOfInt64(926)) {
						return (p0).Multiplicity(_dafny.IntOfInt64(926))
					}
					return (_this).F31()
				})(), 5)
				(_nw120).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("b"), (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()))).Uint32(), _713_v54), (Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("b"), (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()))).Uint32(), _713_v54)).Cardinality()))).Uint32(), _dafny.CodePoint('s'))).Cardinality()), 6)
				(_nw120).ArraySet1(_671_i3, 7)
				(_nw120).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F26(), (_this).F31()), 8)
				(_nw120).ArraySet1(_671_i3, 9)
				(_nw120).ArraySet1((_this).F31(), 10)
				(_nw120).ArraySet1(Companion_Default___.Fm2(_671_i3, _dafny.IntOfInt64(-150), globalState), 11)
				_714_v55 = _nw120
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_714_v55), 0))
				_ = _index126
				(_714_v55).ArraySet1((func() _dafny.Int {
					if p1 {
						return (_this).F31()
					}
					return _dafny.IntOfUint32((_dafny.SeqOf((_this).F26(), _671_i3)).Cardinality())
				})(), (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_714_v55), 0))
				_ = _index127
				(_714_v55).ArraySet1(_671_i3, (_index127).Int())
				var _715_v56 _dafny.Array
				_ = _715_v56
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_18
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) bool = (func(_716_p2 bool) func(_dafny.Int) bool {
						return func(_717_i5 _dafny.Int) bool {
							return _716_p2
						}
					})(p2)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw121 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw121).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw121).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_715_v56 = _nw121
				var _718_v57 *C0
				_ = _718_v57
				var _nw122 *C0 = New_C0_()
				_ = _nw122
				_nw122.Ctor__(_646_v11, (_this).F25(), _dafny.IntOfInt64(296), _715_v56, p2)
				_718_v57 = _nw122
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_715_v56), 0))
				_ = _index128
				(_715_v56).ArraySet1(p2, (_index128).Int())
				var _719_v58 D0
				_ = _719_v58
				_719_v58 = Companion_D0_.Create_DC0_(p2)
				var _pat_let_tv20 = p1
				_ = _pat_let_tv20
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_715_v56), 0))
				_ = _index129
				(_715_v56).ArraySet1(((_this).F25()) && (Companion_Default___.Fm5((_this).F25(), (_708_v51).Update(func(_pat_let16_0 D0) D0 {
					return func(_720_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let17_0 bool) D0 {
							return func(_721_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC0_(_721_dt__update_hcf0_h0)
							}(_pat_let17_0)
						}(_pat_let_tv20)
					}(_pat_let16_0)
				}(_719_v58), (_714_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_714_v55), 0))).Int()).(_dafny.Int)), (_this).F26(), globalState)), (_index129).Int())
				var _722_v59 _dafny.Sequence
				_ = _722_v59
				_722_v59 = _dafny.SeqOf(_715_v56)
				var _723_v60 _dafny.Sequence
				_ = _723_v60
				_723_v60 = _dafny.SeqOf((_722_v59).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_722_v59).Cardinality()))).Uint32()).(_dafny.Array))
				_723_v60 = (func() _dafny.Sequence {
					if (p0).IsProperSubsetOf(p0) {
						return _dafny.Companion_Sequence_.Concatenate(_722_v59, _723_v60)
					}
					return _723_v60
				})()
				(globalState).F2 = _671_i3
			} else {
				var _724_v61 _dafny.Array
				_ = _724_v61
				var _nwElement0_25 bool = (_this).F25()
				_ = _nwElement0_25
				var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(5))
				_ = _nw123
				(_nw123).ArraySet1(_nwElement0_25, 0)
				(_nw123).ArraySet1((_this).F25(), 1)
				(_nw123).ArraySet1(p1, 2)
				(_nw123).ArraySet1(!(p1), 3)
				(_nw123).ArraySet1(p1, 4)
				_724_v61 = _nw123
				var _725_v62 T0
				_ = _725_v62
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__(_646_v11, p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kiivsm")).Cardinality()), _724_v61, p2)
				_725_v62 = _nw124
				var _726_v63 _dafny.Sequence
				_ = _726_v63
				_726_v63 = _dafny.SeqOf(_725_v62)
				var _727_v64 _dafny.Map
				_ = _727_v64
				_727_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _726_v63)
				var _728_v65 _dafny.Map
				_ = _728_v65
				_728_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Sequence {
					if (_727_v64).Contains((_725_v62).F25()) {
						return (_727_v64).Get((_725_v62).F25()).(_dafny.Sequence)
					}
					return _726_v63
				})())
				var _729_v66 D9
				_ = _729_v66
				_729_v66 = Companion_D9_.Create_DC18_(_727_v64)
				_728_v65 = (((_729_v66).Dtor_cf30()).Update((_725_v62).F25(), _726_v63)).Merge(_727_v64)
				(globalState).F1 = (func() bool {
					if (_725_v62).F25() {
						return true
					}
					return _dafny.Companion_Sequence_.IsProperPrefixOf(_646_v11, _dafny.UnicodeSeqOfUtf8Bytes("wmubls"))
				})()
				(globalState).F2 = _dafny.IntOfUint32((_673_v28).Cardinality())
				var _730_v67 _dafny.Int
				_ = _730_v67
				var _out16 _dafny.Int
				_ = _out16
				_out16 = Companion_Default___.M0(p1, globalState)
				_730_v67 = _out16
				var _731_v68 _dafny.Map
				_ = _731_v68
				_731_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_646_v11, _730_v67)
				(globalState).F1 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_646_v11, (_this).Fm6((_this).F25(), (_this).F25(), _671_i3, globalState))).Equals(_731_v68)
			}
		}
		r0 = (_this).Fm7(globalState)
		return r0
	}
}
func (_this *C2) M5(p0 bool, p1 D1, p2 _dafny.CodePoint, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _732_v0 _dafny.Set
		_ = _732_v0
		_732_v0 = _dafny.SetOf(p0, (_this).F25(), (_this).F25(), (_this).F25(), !((_this).F25()) || (!(p0)))
		var _733_v1 _dafny.Array
		_ = _733_v1
		var _nw125 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw125
		_733_v1 = _nw125
		var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
		_ = _index130
		(_733_v1).ArraySet1(p0, (_index130).Int())
		var _734_v2 D0
		_ = _734_v2
		_734_v2 = Companion_D0_.Create_DC0_(p0)
		var _735_v3 _dafny.Map
		_ = _735_v3
		_735_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5((_this).F25(), Companion_Default___.Fm25((_this).F26(), (_734_v2).Dtor_cf0(), p0, globalState), _dafny.IntOfInt64(345), globalState), _732_v0)
		var _736_v4 _dafny.Sequence
		_ = _736_v4
		_736_v4 = _dafny.UnicodeSeqOfUtf8Bytes("nf")
		var _737_v5 _dafny.Sequence
		_ = _737_v5
		_737_v5 = _dafny.SeqOf(Companion_Default___.Fm26((_this).F26(), p2, p0, globalState))
		var _738_v6 _dafny.Map
		_ = _738_v6
		_738_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)
		var _739_v7 _dafny.MultiSet
		_ = _739_v7
		_739_v7 = _dafny.MultiSetOf(_738_v6, _738_v6)
		var _740_v8 _dafny.Map
		_ = _740_v8
		_740_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v2, (_this).F31())
		var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
		_ = _index131
		var _rhs149 _dafny.Set = (func() _dafny.Set {
			if (_735_v3).Contains(!(p0) || ((_this).F25())) {
				return (_735_v3).Get(!(p0) || ((_this).F25())).(_dafny.Set)
			}
			return _732_v0
		})()
		_ = _rhs149
		var _rhs150 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).F25() {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_736_v4, _736_v4)).Cardinality())
			}
			return (_this).F31()
		})())
		_ = _rhs150
		var _rhs151 bool = ((_739_v7).Update(_738_v6, Companion_Default___.Abs((_this).F31()))).IsProperSubsetOf((_737_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F31()), _dafny.IntOfUint32((_737_v5).Cardinality()))).Uint32()).(_dafny.MultiSet))
		_ = _rhs151
		var _rhs152 bool = (func() bool {
			if Companion_Default___.Fm5((_this).F25(), _740_v8, (_this).F31(), globalState) {
				return _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_741_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_742_i0 _dafny.Int) _dafny.CodePoint {
						return _741_p2
					}
				})(p2))), p2)
			}
			return (_this).F25()
		})()
		_ = _rhs152
		var _lhs131 *GlobalState = globalState
		_ = _lhs131
		var _lhs132 *GlobalState = globalState
		_ = _lhs132
		var _lhs133 _dafny.Array = _733_v1
		_ = _lhs133
		var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
		_ = _lhs134
		_732_v0 = _rhs149
		_lhs131.F0 = _rhs150
		_lhs132.F1 = _rhs151
		(_lhs133).ArraySet1(_rhs152, (_lhs134).Int())
		var _hi7 _dafny.Int = (_this).F31()
		_ = _hi7
		for _743_i1 := (_this).F31(); _743_i1.Cmp(_hi7) < 0; _743_i1 = _743_i1.Plus(_dafny.One) {
			var _744_v9 _dafny.Array
			_ = _744_v9
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_19
			var _nw126 _dafny.Array
			_ = _nw126
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw126 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Sequence = (func(_745_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_746_i2 _dafny.Int) _dafny.Sequence {
						return _745_v4
					}
				})(_736_v4)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw126 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw126).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw126).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_744_v9 = _nw126
			var _747_v10 _dafny.Sequence
			_ = _747_v10
			_747_v10 = _dafny.SeqOf((_this).F31())
			var _748_v11 _dafny.Map
			_ = _748_v11
			_748_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_v9, _747_v10)
			_748_v11 = (_748_v11).Update(_744_v9, _747_v10)
			var _749_v12 _dafny.Map
			_ = _749_v12
			_749_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_747_v10, p2)
			var _750_v13 *C0
			_ = _750_v13
			var _nw127 *C0 = New_C0_()
			_ = _nw127
			_nw127.Ctor__(_736_v4, !(((_this).F26()).Cmp((_this).F31()) > 0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_736_v4, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.IntOfUint32((_736_v4).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if (_749_v12).Contains(_747_v10) {
					return (_749_v12).Get(_747_v10).(_dafny.CodePoint)
				}
				return p2
			})()), (Companion_Default___.SafeIndex(_743_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_736_v4, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.IntOfUint32((_736_v4).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if (_749_v12).Contains(_747_v10) {
					return (_749_v12).Get(_747_v10).(_dafny.CodePoint)
				}
				return p2
			})())).Cardinality()))).Uint32(), p2)).Cardinality()), _733_v1, !((_this).F25()))
			_750_v13 = _nw127
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
			_ = _index132
			(_733_v1).ArraySet1(!((_this).F25()), (_index132).Int())
			var _751_v14 _dafny.Map
			_ = _751_v14
			_751_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _743_i1)
			var _752_v15 _dafny.MultiSet
			_ = _752_v15
			_752_v15 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), (_this).F31()), (_751_v14).Update((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _dafny.IntOfInt64(-691)))
			var _753_v16 _dafny.MultiSet
			_ = _753_v16
			_753_v16 = _752_v15
			var _source11 _dafny.MultiSet = _dafny.MultiSetOf(_751_v14)
			_ = _source11
			var _754___mcc_h0 _dafny.MultiSet = _source11
			_ = _754___mcc_h0
			var _755_cf14 _dafny.MultiSet = _754___mcc_h0
			_ = _755_cf14
			var _756_v17 _dafny.MultiSet
			_ = _756_v17
			_756_v17 = _dafny.MultiSetOf(p0)
			var _757_v18 *C0
			_ = _757_v18
			var _nw128 *C0 = New_C0_()
			_ = _nw128
			_nw128.Ctor__(_750_v13.F35, (_756_v17).IsDisjointFrom(_756_v17), ((_732_v0).Intersection(_dafny.SetOf(p0))).Cardinality(), _733_v1, (Companion_Default___.Fm27((_this).F26(), globalState)).IsDisjointFrom(_756_v17))
			_757_v18 = _nw128
			var _758_v19 _dafny.Sequence
			_ = _758_v19
			_758_v19 = _dafny.SeqOf(!(p0))
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
			_ = _index133
			var _rhs153 _dafny.Sequence = _758_v19
			_ = _rhs153
			var _rhs154 _dafny.Int = _743_i1
			_ = _rhs154
			var _rhs155 bool = ((_756_v17).Intersection(_756_v17)).IsProperSubsetOf(_756_v17)
			_ = _rhs155
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			var _lhs136 _dafny.Array = _733_v1
			_ = _lhs136
			var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
			_ = _lhs137
			_758_v19 = _rhs153
			_lhs135.F0 = _rhs154
			(_lhs136).ArraySet1(_rhs155, (_lhs137).Int())
			(globalState).F2 = Companion_Default___.SafeDivisionInt(((_this).F26()).Minus((_this).F26()), (_this).F31())
			var _759_v20 _dafny.Array
			_ = _759_v20
			var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw129
			_759_v20 = _nw129
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_759_v20), 0))
			_ = _index134
			(_759_v20).ArraySet1((_this).F26(), (_index134).Int())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_759_v20), 0))
			_ = _index135
			(_759_v20).ArraySet1((_this).F26(), (_index135).Int())
		}
		var _760_v21 _dafny.Array
		_ = _760_v21
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw130
		_760_v21 = _nw130
		var _761_v22 _dafny.MultiSet
		_ = _761_v22
		_761_v22 = _dafny.MultiSetOf((_this).F25(), (_this).F25(), (_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool))
		var _762_v23 _dafny.Sequence
		_ = _762_v23
		_762_v23 = _dafny.SeqOf((_this).F25())
		var _763_v24 _dafny.Set
		_ = _763_v24
		_763_v24 = _dafny.SetOf((_this).F31(), (_this).F31())
		var _764_v25 _dafny.Map
		_ = _764_v25
		_764_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_762_v23)).Cardinality(), _763_v24)
		var _765_v26 D6
		_ = _765_v26
		_765_v26 = Companion_D6_.Create_DC13_((_this).F31(), (_this).F31(), (func() _dafny.Set {
			if (_764_v25).Contains((_this).F26()) {
				return (_764_v25).Get((_this).F26()).(_dafny.Set)
			}
			return _dafny.SetOf((_this).F26())
		})())
		var _766_v27 _dafny.Array
		_ = _766_v27
		var _nw131 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
		_ = _nw131
		_766_v27 = _nw131
		var _767_v28 _dafny.Sequence
		_ = _767_v28
		_767_v28 = _dafny.SeqOf(_766_v27, _766_v27)
		var _768_v29 _dafny.Array
		_ = _768_v29
		var _nwElement0_26 _dafny.Int = (_this).F26()
		_ = _nwElement0_26
		var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(10))
		_ = _nw132
		(_nw132).ArraySet1(_nwElement0_26, 0)
		(_nw132).ArraySet1((func() _dafny.Int {
			if (_761_v22).Contains(p0) {
				return (_761_v22).Multiplicity(p0)
			}
			return (_this).F26()
		})(), 1)
		(_nw132).ArraySet1(_dafny.IntOfUint32((_736_v4).Cardinality()), 2)
		(_nw132).ArraySet1(_dafny.IntOfInt64(880), 3)
		(_nw132).ArraySet1((_this).F31(), 4)
		(_nw132).ArraySet1((_this).F31(), 5)
		(_nw132).ArraySet1(_dafny.IntOfUint32((_736_v4).Cardinality()), 6)
		(_nw132).ArraySet1((_this).F31(), 7)
		(_nw132).ArraySet1((_765_v26).Dtor_cf20(), 8)
		(_nw132).ArraySet1(_dafny.IntOfUint32((_767_v28).Cardinality()), 9)
		_768_v29 = _nw132
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))
		_ = _index136
		(_760_v21).ArraySet1(_768_v29, (_index136).Int())
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))
		_ = _index137
		(_760_v21).ArraySet1(_768_v29, (_index137).Int())
		var _hi8 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F26())).Cardinality()
		_ = _hi8
		for _769_i3 := (_this).F31(); _769_i3.Cmp(_hi8) < 0; _769_i3 = _769_i3.Plus(_dafny.One) {
			(globalState).F0 = _769_i3
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr2
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index138
			(_arr2).ArraySet1(_769_i3, (_index138).Int())
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr3
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index139
			(_arr3).ArraySet1((_this).F26(), (_index139).Int())
			var _770_v30 _dafny.Map
			_ = _770_v30
			_770_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.Fm5((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _740_v8, (_this).F26(), globalState))
			var _771_v31 _dafny.Set
			_ = _771_v31
			_771_v31 = _dafny.SetOf(_770_v30, _770_v30, _770_v30)
			var _772_v32 _dafny.MultiSet
			_ = _772_v32
			_772_v32 = _dafny.MultiSetOf((_dafny.Zero).Minus((_771_v31).Cardinality()), (_this).F26(), _dafny.IntOfInt64(-489))
			var _arr4 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr4
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index140
			var _arr5 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr5
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index141
			var _rhs156 _dafny.Int = (_this).F26()
			_ = _rhs156
			var _rhs157 _dafny.Int = (func() _dafny.Int {
				if (_761_v22).Contains(p0) {
					return (_761_v22).Multiplicity(p0)
				}
				return (_this).F31()
			})()
			_ = _rhs157
			var _rhs158 _dafny.Int = ((_this).F31()).Times(_769_i3)
			_ = _rhs158
			var _rhs159 _dafny.Int = Companion_Default___.SafeModuloInt(_769_i3, (_772_v32).Cardinality())
			_ = _rhs159
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			var _lhs139 *GlobalState = globalState
			_ = _lhs139
			var _lhs140 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _lhs141
			var _lhs142 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _lhs142
			var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _lhs143
			_lhs138.F2 = _rhs156
			_lhs139.F0 = _rhs157
			(_lhs140).ArraySet1(_rhs158, (_lhs141).Int())
			(_lhs142).ArraySet1(_rhs159, (_lhs143).Int())
			var _773_v33 *C1
			_ = _773_v33
			var _nw133 *C1 = New_C1_()
			_ = _nw133
			_nw133.Ctor__((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool))
			_773_v33 = _nw133
			var _774_v35 _dafny.Map
			_ = _774_v35
			_774_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_736_v4, _736_v4)
			var _775_v36 _dafny.Map
			_ = _775_v36
			_775_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm11(_774_v35, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer44 func(_dafny.Int) D9) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_776_i3 _dafny.Int, _777_v22 _dafny.MultiSet) func(_dafny.Int) D9 {
				return func(_778_i4 _dafny.Int) D9 {
					return Companion_D9_.Create_DC19_(Companion_D6_.Create_DC14_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_776_i3, (_this).F31())).Cardinality(), (_this).F31(), _777_v22))
				}
			})(_769_i3, _761_v22)))).Cardinality()), _769_i3, globalState), (_772_v32).Cardinality())
			var _779_v37 bool
			_ = _779_v37
			var _out17 bool
			_ = _out17
			_out17 = (_this).M6((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate(((_772_v32).Update((_this).F26(), Companion_Default___.Abs((_this).F31()))).Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _780_v34 _dafny.Int
					_780_v34 = interface{}(_compr_17).(_dafny.Int)
					if ((_772_v32).Update((_this).F26(), Companion_Default___.Abs((_this).F31()))).Contains(_780_v34) {
						_coll17.Add((_780_v34).Plus((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))).Int()).(_dafny.Int)), _769_i3)
					}
				}
				return _coll17.ToMap()
			}()).Merge(_775_v36), (_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))).Int()).(_dafny.Int), (_this).F31(), globalState)
			_779_v37 = _out17
		}
		(globalState).F0 = (_dafny.Zero).Minus((_this).F26())
		var _781_v39 _dafny.Sequence
		_ = _781_v39
		_781_v39 = _dafny.SeqOf(_dafny.IntOfInt64(801))
		var _782_v40 D0
		_ = _782_v40
		_782_v40 = Companion_D0_.Create_DC1_((_this).F26(), _dafny.UnicodeSeqOfUtf8Bytes("ac"), (_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _dafny.IntOfUint32((_762_v23).Cardinality()), Companion_Default___.Fm5((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _740_v8, (func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_781_v39).Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _783_v38 _dafny.Int
				_783_v38 = interface{}(_compr_18).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_781_v39, _783_v38) {
					_coll18.Add((_783_v38).Minus(_dafny.IntOfInt64(619)), p0)
				}
			}
			return _coll18.ToMap()
		}()).Cardinality(), globalState))
		var _784_v41 _dafny.Map
		_ = _784_v41
		_784_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_dafny.IntOfUint32((_736_v4).Cardinality()), _dafny.IntOfUint32((_762_v23).Cardinality()), globalState), p0)
		var _pat_let_tv21 = _784_v41
		_ = _pat_let_tv21
		var _pat_let_tv22 = p0
		_ = _pat_let_tv22
		if (_dafny.IntOfUint32(((func(_pat_let18_0 D0) D0 {
			return func(_785_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let19_0 _dafny.Int) D0 {
					return func(_786_dt__update_hcf1_h0 _dafny.Int) D0 {
						return func(_pat_let20_0 bool) D0 {
							return func(_787_dt__update_hcf5_h0 bool) D0 {
								return func(_pat_let21_0 _dafny.Int) D0 {
									return func(_788_dt__update_hcf4_h0 _dafny.Int) D0 {
										return Companion_D0_.Create_DC1_(_786_dt__update_hcf1_h0, (_785_dt__update__tmp_h1).Dtor_cf2(), (_785_dt__update__tmp_h1).Dtor_cf3(), _788_dt__update_hcf4_h0, _787_dt__update_hcf5_h0)
									}(_pat_let21_0)
								}((_this).F31())
							}(_pat_let20_0)
						}(_pat_let_tv22)
					}(_pat_let19_0)
				}((_pat_let_tv21).Cardinality())
			}(_pat_let18_0)
		}(_782_v40)).Dtor_cf2()).Cardinality())).Cmp((_this).F26()) >= 0 {
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))
			_ = _index142
			(_733_v1).ArraySet1(((_dafny.Zero).Minus((_this).F31())).Cmp((_this).F31()) != 0, (_index142).Int())
			(globalState).F1 = (_this).F25()
			(globalState).F0 = Companion_Default___.SafeModuloInt((_this).F26(), ((_dafny.Zero).Minus((_this).F26())).Plus((_this).F31()))
			var _arr6 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr6
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index143
			(_arr6).ArraySet1((_this).F31(), (_index143).Int())
			var _arr7 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr7
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index144
			(_arr7).ArraySet1((_this).F26(), (_index144).Int())
			var _789_v42 _dafny.Map
			_ = _789_v42
			_789_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_dafny.MultiSetFromSeq(Companion_Default___.Fm23(globalState))).Cardinality())
			_789_v42 = (_789_v42).Update(p0, (_this).F31())
		} else {
			(globalState).F0 = _dafny.IntOfInt64(432)
			(globalState).F0 = ((_dafny.IntOfUint32((_736_v4).Cardinality())).Times((_this).F26())).Plus(_dafny.IntOfInt64(800))
			var _790_v43 D6
			_ = _790_v43
			_790_v43 = Companion_D6_.Create_DC14_(_dafny.IntOfUint32((_736_v4).Cardinality()), (_this).F26(), _761_v22)
			var _791_v44 D9
			_ = _791_v44
			_791_v44 = Companion_D9_.Create_DC19_(_790_v43)
			var _792_v45 D9
			_ = _792_v45
			_792_v45 = Companion_D9_.Create_DC20_(_791_v44)
			var _793_v46 D9
			_ = _793_v46
			_793_v46 = Companion_D9_.Create_DC20_(_792_v45)
			var _794_v47 _dafny.Map
			_ = _794_v47
			_794_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_768_v29, _793_v46)
			var _795_v48 _dafny.Map
			_ = _795_v48
			_795_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _794_v47)
			_795_v48 = (_795_v48).Update((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), _794_v47)
			var _796_v49 *C1
			_ = _796_v49
			var _nw134 *C1 = New_C1_()
			_ = _nw134
			_nw134.Ctor__((_this).F25())
			_796_v49 = _nw134
			var _arr8 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr8
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index145
			(_arr8).ArraySet1((_761_v22).Cardinality(), (_index145).Int())
			var _arr9 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _arr9
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _index146
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))
			_ = _index147
			var _rhs160 *C1 = _796_v49
			_ = _rhs160
			var _rhs161 bool = !(p0)
			_ = _rhs161
			var _rhs162 _dafny.Int = (_this).F26()
			_ = _rhs162
			var _rhs163 _dafny.Array = _768_v29
			_ = _rhs163
			var _lhs144 *GlobalState = globalState
			_ = _lhs144
			var _lhs145 _dafny.Array = _dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_dafny.ArrayCastTo((_760_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))).Int()))), 0))
			_ = _lhs146
			var _lhs147 _dafny.Array = _760_v21
			_ = _lhs147
			var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_760_v21), 0))
			_ = _lhs148
			_796_v49 = _rhs160
			_lhs144.F1 = _rhs161
			(_lhs145).ArraySet1(_rhs162, (_lhs146).Int())
			(_lhs147).ArraySet1(_rhs163, (_lhs148).Int())
			var _797_v50 _dafny.Map
			_ = _797_v50
			_797_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7(globalState), (_796_v49).F32())
			var _798_v51 _dafny.Map
			_ = _798_v51
			_798_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_781_v39).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_781_v39).Cardinality()))).Uint32()).(_dafny.Int), _797_v50)
			var _799_v52 D9
			_ = _799_v52
			_799_v52 = Companion_D9_.Create_DC19_(_790_v43)
			var _800_v53 _dafny.Map
			_ = _800_v53
			_800_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if !((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool)) {
					return _797_v50
				}
				return (func() _dafny.Map {
					if (_798_v51).Contains((_this).F31()) {
						return (_798_v51).Get((_this).F31()).(_dafny.Map)
					}
					return _797_v50
				})()
			})(), _799_v52)
			_800_v53 = _800_v53
		}
		r0 = !(_dafny.Companion_Sequence_.Equal(_762_v23, _dafny.SeqOf((_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool), (_733_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_733_v1), 0))).Int()).(bool))))
		var _801_v54 _dafny.MultiSet
		_ = _801_v54
		_801_v54 = _dafny.MultiSetOf((_this).Fm6(false, false, (_this).F31(), globalState), (_this).F26(), (_this).F31())
		var _pat_let_tv23 = p0
		_ = _pat_let_tv23
		r1 = func(_source12 _dafny.MultiSet) bool {
			var _802___mcc_h1 _dafny.MultiSet = _source12
			_ = _802___mcc_h1
			var _803_cf18 _dafny.MultiSet = _802___mcc_h1
			_ = _803_cf18
			return _pat_let_tv23
		}(_801_v54)
		return r0, r1
	}
}
func (_this *C2) M6(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _804_v0 _dafny.CodePoint
		_ = _804_v0
		_804_v0 = _dafny.CodePoint('y')
		_804_v0 = _804_v0
		if (func() bool {
			if false {
				return (_this).F25()
			}
			return Companion_Default___.Fm5((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_this).F25()), p2), _dafny.IntOfInt64(726), globalState)
		})() {
			var _805_v1 _dafny.MultiSet
			_ = _805_v1
			_805_v1 = _dafny.MultiSetOf(_dafny.IntOfInt64(-40))
			var _806_v2 _dafny.Sequence
			_ = _806_v2
			_806_v2 = _dafny.SeqOf(p2, p2, p2, (_dafny.Zero).Minus((func() _dafny.Int {
				if (_805_v1).Contains(_dafny.IntOfInt64(7)) {
					return (_805_v1).Multiplicity(_dafny.IntOfInt64(7))
				}
				return p2
			})()))
			var _807_v3 _dafny.Set
			_ = _807_v3
			_807_v3 = _dafny.SetOf((_this).F26(), (_this).F31())
			var _808_v4 D6
			_ = _808_v4
			_808_v4 = Companion_D6_.Create_DC13_((_this).F26(), (_this).F26(), _807_v3)
			var _809_v5 _dafny.Sequence
			_ = _809_v5
			_809_v5 = _dafny.SeqOf(_806_v2, _806_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}(func(_810_i1 _dafny.Int) _dafny.Int {
				return (_this).F26()
			})))
			var _811_v6 _dafny.Sequence
			_ = _811_v6
			_811_v6 = _dafny.SeqOf((_this).F25())
			var _812_v7 _dafny.Array
			_ = _812_v7
			var _nwElement0_27 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_806_v2, (Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_806_v2).Cardinality()))).Uint32(), (_this).F26())
			_ = _nwElement0_27
			var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(26))
			_ = _nw135
			(_nw135).ArraySet1(_nwElement0_27, 0)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), 1)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), 2)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), 3)
			(_nw135).ArraySet1(_806_v2, 4)
			(_nw135).ArraySet1(_806_v2, 5)
			(_nw135).ArraySet1(Companion_Default___.Fm3(p1, false, p2, (_this).F25(), globalState), 6)
			(_nw135).ArraySet1(_806_v2, 7)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), 8)
			(_nw135).ArraySet1(_806_v2, 9)
			(_nw135).ArraySet1(_806_v2, 10)
			(_nw135).ArraySet1(_dafny.SeqOf(p2, p2, p1), 11)
			(_nw135).ArraySet1(_806_v2, 12)
			(_nw135).ArraySet1(_806_v2, 13)
			(_nw135).ArraySet1(_806_v2, 14)
			(_nw135).ArraySet1(_806_v2, 15)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), (Companion_Default___.SafeIndex((_808_v4).Dtor_cf20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2)).Cardinality()))).Uint32(), _dafny.IntOfInt64(383)), 16)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), 17)
			(_nw135).ArraySet1(_806_v2, 18)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(327), (_this).F31(), p2, (_this).F26()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-679))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_813_i0 _dafny.Int) _dafny.Int {
				return (_this).F26()
			}))), 19)
			(_nw135).ArraySet1(_dafny.SeqOf((_this).F26(), p1, p1), 20)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_806_v2, _806_v2)).Cardinality()))).Uint32(), p1), 21)
			(_nw135).ArraySet1((_809_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_811_v6).Cardinality()), _dafny.IntOfUint32((_809_v5).Cardinality()))).Uint32()).(_dafny.Sequence), 22)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_806_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_814_i2 _dafny.Int) _dafny.Int {
				return (_this).F26()
			}))), 23)
			(_nw135).ArraySet1(_806_v2, 24)
			(_nw135).ArraySet1(_806_v2, 25)
			_812_v7 = _nw135
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_812_v7), 0))
			_ = _index148
			(_812_v7).ArraySet1(_806_v2, (_index148).Int())
			var _815_v8 _dafny.MultiSet
			_ = _815_v8
			_815_v8 = _dafny.MultiSetOf(true)
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_812_v7), 0))
			_ = _index149
			(_812_v7).ArraySet1(_dafny.Companion_Sequence_.Update((_809_v5).Select((Companion_Default___.SafeIndex(((_this).F31()).Times((_this).F26()), _dafny.IntOfUint32((_809_v5).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(((_815_v8).Update((_this).F25(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F26())))).Cardinality(), _dafny.IntOfUint32(((_809_v5).Select((Companion_Default___.SafeIndex(((_this).F31()).Times((_this).F26()), _dafny.IntOfUint32((_809_v5).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), (_this).F31()), (_index149).Int())
			var _816_v9 _dafny.Int
			_ = _816_v9
			var _out18 _dafny.Int
			_ = _out18
			_out18 = Companion_Default___.M0(!((_this).F25()), globalState)
			_816_v9 = _out18
			var _817_v10 _dafny.Int
			_ = _817_v10
			var _out19 _dafny.Int
			_ = _out19
			_out19 = Companion_Default___.M0((_this).F25(), globalState)
			_817_v10 = _out19
			if !((_this).F25()) {
				var _818_v11 _dafny.Sequence
				_ = _818_v11
				_818_v11 = _dafny.UnicodeSeqOfUtf8Bytes("sxl")
				var _819_v12 _dafny.Array
				_ = _819_v12
				var _nwElement0_28 bool = (_this).F25()
				_ = _nwElement0_28
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(28))
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_28, 0)
				(_nw136).ArraySet1(!((_this).F25()), 1)
				(_nw136).ArraySet1((_this).F25(), 2)
				(_nw136).ArraySet1(!((_this).F25()), 3)
				(_nw136).ArraySet1((_this).F25(), 4)
				(_nw136).ArraySet1((_this).F25(), 5)
				(_nw136).ArraySet1((_this).F25(), 6)
				(_nw136).ArraySet1((_this).F25(), 7)
				(_nw136).ArraySet1((_this).F25(), 8)
				(_nw136).ArraySet1((_this).F25(), 9)
				(_nw136).ArraySet1((_this).F25(), 10)
				(_nw136).ArraySet1((_this).F25(), 11)
				(_nw136).ArraySet1((_this).F25(), 12)
				(_nw136).ArraySet1((_this).F25(), 13)
				(_nw136).ArraySet1((_this).F25(), 14)
				(_nw136).ArraySet1((_this).F25(), 15)
				(_nw136).ArraySet1((_this).F25(), 16)
				(_nw136).ArraySet1((_this).F25(), 17)
				(_nw136).ArraySet1((_this).F25(), 18)
				(_nw136).ArraySet1((_this).F25(), 19)
				(_nw136).ArraySet1((_this).F25(), 20)
				(_nw136).ArraySet1((_this).F25(), 21)
				(_nw136).ArraySet1(false, 22)
				(_nw136).ArraySet1((_this).F25(), 23)
				(_nw136).ArraySet1((_this).F25(), 24)
				(_nw136).ArraySet1(true, 25)
				(_nw136).ArraySet1((_this).F25(), 26)
				(_nw136).ArraySet1((_this).F25(), 27)
				_819_v12 = _nw136
				var _820_v13 *C0
				_ = _820_v13
				var _nw137 *C0 = New_C0_()
				_ = _nw137
				_nw137.Ctor__(_818_v11, (_this).F25(), (_this).F26(), _819_v12, !((_this).F25()))
				_820_v13 = _nw137
				var _821_v14 *C1
				_ = _821_v14
				var _nw138 *C1 = New_C1_()
				_ = _nw138
				_nw138.Ctor__((func() bool {
					if (_this).F25() {
						return false
					}
					return (_this).F25()
				})())
				_821_v14 = _nw138
				var _822_v15 _dafny.Array
				_ = _822_v15
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw139
				_822_v15 = _nw139
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_822_v15), 0))
				_ = _index150
				(_822_v15).ArraySet1((_this).F31(), (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_822_v15), 0))
				_ = _index151
				(_822_v15).ArraySet1(p1, (_index151).Int())
				var _823_v16 _dafny.Set
				_ = _823_v16
				_823_v16 = _dafny.SetOf(_822_v15, _822_v15)
				var _824_v17 D6
				_ = _824_v17
				_824_v17 = Companion_D6_.Create_DC12_(_823_v16)
				var _825_v18 D6
				_ = _825_v18
				_825_v18 = Companion_D6_.Create_DC12_((_823_v16).Union((_824_v17).Dtor_cf19()))
				_824_v17 = _825_v18
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_819_v12), 0))
				_ = _index152
				(_819_v12).ArraySet1((func() bool {
					if (_this).F25() {
						return (_821_v14).F32()
					}
					return (_this).F25()
				})(), (_index152).Int())
				var _826_v19 _dafny.Set
				_ = _826_v19
				_826_v19 = _dafny.SetOf(_820_v13.F35)
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_822_v15), 0))
				_ = _index153
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_819_v12), 0))
				_ = _index154
				var _rhs164 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F31(), _816_v9), (_this).F26())
				_ = _rhs164
				var _rhs165 bool = (_821_v14).F32()
				_ = _rhs165
				var _rhs166 _dafny.Int = _dafny.IntOfInt64(727)
				_ = _rhs166
				var _rhs167 bool = ((_826_v19).Intersection(_826_v19)).IsProperSubsetOf(_826_v19)
				_ = _rhs167
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				var _lhs150 _dafny.Array = _822_v15
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_822_v15), 0))
				_ = _lhs151
				var _lhs152 _dafny.Array = _819_v12
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_819_v12), 0))
				_ = _lhs153
				_817_v10 = _rhs164
				_lhs149.F1 = _rhs165
				(_lhs150).ArraySet1(_rhs166, (_lhs151).Int())
				(_lhs152).ArraySet1(_rhs167, (_lhs153).Int())
			} else {
				var _827_v20 _dafny.Map
				_ = _827_v20
				_827_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_805_v1).Contains(_816_v9) {
						return (_805_v1).Multiplicity(_816_v9)
					}
					return _dafny.IntOfInt64(-600)
				})(), (func() _dafny.Sequence {
					if (_this).F25() {
						return _811_v6
					}
					return _811_v6
				})())
				_827_v20 = (func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter19 := _dafny.Iterate(((_812_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_812_v7), 0))).Int()).(_dafny.Sequence)).Elements()); ; {
						_compr_19, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _828_v21 _dafny.Int
						_828_v21 = interface{}(_compr_19).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains((_812_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_812_v7), 0))).Int()).(_dafny.Sequence), _828_v21) {
							_coll19.Add((_828_v21).Minus((_dafny.Zero).Minus(p2)), _811_v6)
						}
					}
					return _coll19.ToMap()
				}()).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_812_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_812_v7), 0))).Int()).(_dafny.Sequence)).Cardinality()), p1), _811_v6)
				(globalState).F1 = (_this).F25()
				_812_v7 = _812_v7
				var _829_v22 _dafny.Array
				_ = _829_v22
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_20
				var _nw140 _dafny.Array
				_ = _nw140
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw140 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = func(_830_i3 _dafny.Int) _dafny.Int {
						return (_830_i3).Minus(_dafny.IntOfInt64(948))
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw140 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw140).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw140).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_829_v22 = _nw140
				_829_v22 = _829_v22
				r0 = (p1).Cmp(p1) > 0
			}
			r0 = !((_this).F25())
		} else {
			var _831_v23 _dafny.Array
			_ = _831_v23
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_21
			var _nw141 _dafny.Array
			_ = _nw141
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw141 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_832_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_833_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_833_i4, (_dafny.Zero).Minus(_832_p2))
					}
				})(p2)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw141 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw141).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw141).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_831_v23 = _nw141
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))
			_ = _index155
			(_831_v23).ArraySet1(((_this).F31()).Minus(p2), (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))
			_ = _index156
			(_831_v23).ArraySet1((func() _dafny.Int {
				if (_this).F25() {
					return (_dafny.Zero).Minus(p2)
				}
				return (_this).F26()
			})(), (_index156).Int())
			var _834_v24 _dafny.Sequence
			_ = _834_v24
			_834_v24 = _dafny.UnicodeSeqOfUtf8Bytes("pwperytjt")
			var _835_v25 _dafny.Map
			_ = _835_v25
			_835_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)
			var _source13 D0 = Companion_Default___.Fm28(_834_v24, (_834_v24).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_834_v24).Cardinality()))).Uint32()).(_dafny.CodePoint), Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)).Update((_this).F25(), (func() _dafny.Int {
				if (_835_v25).Contains((_this).F25()) {
					return (_835_v25).Get((_this).F25()).(_dafny.Int)
				}
				return (_this).F31()
			})()))).Cardinality(), Companion_Default___.Fm2(p1, (_dafny.Zero).Minus(((_835_v25).Update((_this).F25(), (_831_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))).Int()).(_dafny.Int))).Cardinality()), globalState)), globalState)
			_ = _source13
			if _source13.Is_DC1() {
				var _836___mcc_h0 _dafny.Int = _source13.Get_().(D0_DC1).Cf1
				_ = _836___mcc_h0
				var _837___mcc_h1 _dafny.Sequence = _source13.Get_().(D0_DC1).Cf2
				_ = _837___mcc_h1
				var _838___mcc_h2 bool = _source13.Get_().(D0_DC1).Cf3
				_ = _838___mcc_h2
				var _839___mcc_h3 _dafny.Int = _source13.Get_().(D0_DC1).Cf4
				_ = _839___mcc_h3
				var _840___mcc_h4 bool = _source13.Get_().(D0_DC1).Cf5
				_ = _840___mcc_h4
				var _841_cf5 bool = _840___mcc_h4
				_ = _841_cf5
				var _842_cf4 _dafny.Int = _839___mcc_h3
				_ = _842_cf4
				var _843_cf3 bool = _838___mcc_h2
				_ = _843_cf3
				var _844_cf2 _dafny.Sequence = _837___mcc_h1
				_ = _844_cf2
				var _845_cf1 _dafny.Int = _836___mcc_h0
				_ = _845_cf1
				var _846_v26 _dafny.Array
				_ = _846_v26
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_22
				var _nw142 _dafny.Array
				_ = _nw142
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw142 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) bool = (func(_847_cf5 bool) func(_dafny.Int) bool {
						return func(_848_i6 _dafny.Int) bool {
							return _847_cf5
						}
					})(_841_cf5)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw142 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw142).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw142).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_846_v26 = _nw142
				var _849_v27 T1
				_ = _849_v27
				var _nw143 *C0 = New_C0_()
				_ = _nw143
				_nw143.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(685))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_850_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_851_i5 _dafny.Int) _dafny.CodePoint {
						return _850_v0
					}
				})(_804_v0))), _843_cf3, p2, _846_v26, (_this).F25())
				_849_v27 = _nw143
				var _852_v28 _dafny.Sequence
				_ = _852_v28
				_852_v28 = _dafny.SeqOf(_849_v27)
				var _853_v29 _dafny.Set
				_ = _853_v29
				_853_v29 = _dafny.SetOf(_843_cf3)
				var _854_v30 _dafny.Sequence
				_ = _854_v30
				_854_v30 = _dafny.SeqOf(_853_v29)
				var _855_v31 _dafny.Map
				_ = _855_v31
				_855_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_852_v28, _dafny.Companion_Sequence_.Concatenate(_854_v30, _dafny.SeqOf(_853_v29)))
				_855_v31 = (_855_v31).Update(_852_v28, _854_v30)
				var _856_v32 D2
				_ = _856_v32
				_856_v32 = Companion_D2_.Create_DC7_(_843_cf3, false, _842_cf4)
				var _857_v33 _dafny.Sequence
				_ = _857_v33
				_857_v33 = _dafny.SeqOf(_841_cf5)
				var _858_v34 _dafny.Map
				_ = _858_v34
				_858_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _843_cf3)
				var _859_v35 _dafny.MultiSet
				_ = _859_v35
				_859_v35 = _dafny.MultiSetOf(_845_cf1, (_this).F31())
				var _860_v36 _dafny.Map
				_ = _860_v36
				_860_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_857_v33).Cardinality()), !((func() bool {
					if (_858_v34).Contains((func() _dafny.Int {
						if (_859_v35).Contains((_859_v35).Cardinality()) {
							return (_859_v35).Multiplicity((_859_v35).Cardinality())
						}
						return (_this).F26()
					})()) {
						return (_858_v34).Get((func() _dafny.Int {
							if (_859_v35).Contains((_859_v35).Cardinality()) {
								return (_859_v35).Multiplicity((_859_v35).Cardinality())
							}
							return (_this).F26()
						})()).(bool)
					}
					return _843_cf3
				})()))
				var _861_v37 _dafny.MultiSet
				_ = _861_v37
				_861_v37 = _dafny.MultiSetOf(_860_v36)
				var _862_v38 _dafny.Map
				_ = _862_v38
				_862_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_856_v32, ((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-923), _841_cf5))).Update(_858_v34, Companion_Default___.Abs(_845_cf1))).Difference(_861_v37))
				var _863_v39 _dafny.Sequence
				_ = _863_v39
				_863_v39 = _dafny.SeqOf(_861_v37)
				_862_v38 = (_862_v38).Update(_856_v32, (_863_v39).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_863_v39).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _864_v40 _dafny.Set
				_ = _864_v40
				_864_v40 = _dafny.SetOf(_804_v0)
				var _865_v41 _dafny.Int
				_ = _865_v41
				var _out20 _dafny.Int
				_ = _out20
				_out20 = Companion_Default___.M0(!(!(_864_v40).Contains(_804_v0)), globalState)
				_865_v41 = _out20
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))
				_ = _index157
				(_831_v23).ArraySet1(((_this).F26()).Times((_dafny.Zero).Minus(_842_cf4)), (_index157).Int())
			} else if _source13.Is_DC2() {
				var _866_v42 _dafny.Map
				_ = _866_v42
				_866_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))).Int()).(_dafny.Int), (p1).Cmp((_831_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))).Int()).(_dafny.Int)) == 0)
				_866_v42 = (_866_v42).Update((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F26(), (_this).F31())), !((_this).F25()))
				var _867_v43 _dafny.Map
				_ = _867_v43
				_867_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _834_v24)
				var _868_v44 _dafny.Array
				_ = _868_v44
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_23
				var _nw144 _dafny.Array
				_ = _nw144
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw144 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) bool = func(_869_i7 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw144 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw144).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw144).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_868_v44 = _nw144
				var _870_v45 _dafny.Map
				_ = _870_v45
				_870_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _868_v44)
				var _871_v46 T0
				_ = _871_v46
				var _nw145 *C0 = New_C0_()
				_ = _nw145
				_nw145.Ctor__(_834_v24, !(false), (_867_v43).Cardinality(), (func() _dafny.Array {
					if (_870_v45).Contains((_this).F25()) {
						return (_870_v45).Get((_this).F25()).(_dafny.Array)
					}
					return _868_v44
				})(), (_this).F25())
				_871_v46 = _nw145
				var _872_v47 _dafny.Sequence
				_ = _872_v47
				_872_v47 = _dafny.SeqOf(_871_v46, _871_v46)
				var _873_v48 _dafny.Map
				_ = _873_v48
				_873_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _872_v47)
				var _874_v49 D9
				_ = _874_v49
				_874_v49 = Companion_D9_.Create_DC18_(_873_v48)
				var _pat_let_tv24 = _873_v48
				_ = _pat_let_tv24
				var _875_v50 _dafny.Array
				_ = _875_v50
				var _nwElement0_29 D9 = _874_v49
				_ = _nwElement0_29
				var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(29))
				_ = _nw146
				(_nw146).ArraySet1(_nwElement0_29, 0)
				(_nw146).ArraySet1(_874_v49, 1)
				(_nw146).ArraySet1(_874_v49, 2)
				(_nw146).ArraySet1(_874_v49, 3)
				(_nw146).ArraySet1(Companion_D9_.Create_DC18_(_873_v48), 4)
				(_nw146).ArraySet1(Companion_D9_.Create_DC18_(_873_v48), 5)
				(_nw146).ArraySet1(func(_pat_let22_0 D9) D9 {
					return func(_876_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let23_0 _dafny.Map) D9 {
							return func(_877_dt__update_hcf30_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC18_(_877_dt__update_hcf30_h0)
							}(_pat_let23_0)
						}(_pat_let_tv24)
					}(_pat_let22_0)
				}(_874_v49), 6)
				(_nw146).ArraySet1(_874_v49, 7)
				(_nw146).ArraySet1(Companion_D9_.Create_DC18_(_873_v48), 8)
				(_nw146).ArraySet1(_874_v49, 9)
				(_nw146).ArraySet1(_874_v49, 10)
				(_nw146).ArraySet1(_874_v49, 11)
				(_nw146).ArraySet1(_874_v49, 12)
				(_nw146).ArraySet1(_874_v49, 13)
				(_nw146).ArraySet1(_874_v49, 14)
				(_nw146).ArraySet1(_874_v49, 15)
				(_nw146).ArraySet1(_874_v49, 16)
				(_nw146).ArraySet1(_874_v49, 17)
				(_nw146).ArraySet1(_874_v49, 18)
				(_nw146).ArraySet1(_874_v49, 19)
				(_nw146).ArraySet1(_874_v49, 20)
				(_nw146).ArraySet1(_874_v49, 21)
				(_nw146).ArraySet1(_874_v49, 22)
				(_nw146).ArraySet1(_874_v49, 23)
				(_nw146).ArraySet1(_874_v49, 24)
				(_nw146).ArraySet1(Companion_D9_.Create_DC18_(_873_v48), 25)
				(_nw146).ArraySet1(Companion_D9_.Create_DC18_(_873_v48), 26)
				(_nw146).ArraySet1(_874_v49, 27)
				(_nw146).ArraySet1(_874_v49, 28)
				_875_v50 = _nw146
				var _878_v51 _dafny.Array
				_ = _878_v51
				var _nwElement0_30 _dafny.Array = _875_v50
				_ = _nwElement0_30
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(25))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_30, 0)
				(_nw147).ArraySet1(_875_v50, 1)
				(_nw147).ArraySet1(_875_v50, 2)
				(_nw147).ArraySet1(_875_v50, 3)
				(_nw147).ArraySet1(_875_v50, 4)
				(_nw147).ArraySet1(_875_v50, 5)
				(_nw147).ArraySet1(_875_v50, 6)
				(_nw147).ArraySet1(_875_v50, 7)
				(_nw147).ArraySet1(_875_v50, 8)
				(_nw147).ArraySet1(_875_v50, 9)
				(_nw147).ArraySet1(_875_v50, 10)
				(_nw147).ArraySet1(_875_v50, 11)
				(_nw147).ArraySet1(_875_v50, 12)
				(_nw147).ArraySet1(_875_v50, 13)
				(_nw147).ArraySet1(_875_v50, 14)
				(_nw147).ArraySet1(_875_v50, 15)
				(_nw147).ArraySet1(_875_v50, 16)
				(_nw147).ArraySet1(_875_v50, 17)
				(_nw147).ArraySet1(_875_v50, 18)
				(_nw147).ArraySet1(_875_v50, 19)
				(_nw147).ArraySet1(_875_v50, 20)
				(_nw147).ArraySet1(_875_v50, 21)
				(_nw147).ArraySet1(_875_v50, 22)
				(_nw147).ArraySet1(_875_v50, 23)
				(_nw147).ArraySet1(_875_v50, 24)
				_878_v51 = _nw147
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_878_v51), 0))
				_ = _index158
				(_878_v51).ArraySet1(_875_v50, (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_878_v51), 0))
				_ = _index159
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(3))
				_ = _nw148
				(_878_v51).ArraySet1(_nw148, (_index159).Int())
				var _879_v52 *C1
				_ = _879_v52
				var _nw149 *C1 = New_C1_()
				_ = _nw149
				_nw149.Ctor__((_this).F25())
				_879_v52 = _nw149
				_879_v52 = _879_v52
				var _880_v53 _dafny.Map
				_ = _880_v53
				_880_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(956), (_834_v24).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_834_v24).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _881_v54 D4
				_ = _881_v54
				_881_v54 = Companion_D4_.Create_DC10_(_804_v0, _804_v0)
				_880_v53 = (_880_v53).Update(p1, (_881_v54).Dtor_cf16())
			} else {
				var _882___mcc_h5 bool = _source13.Get_().(D0_DC0).Cf0
				_ = _882___mcc_h5
				var _883_cf0 bool = _882___mcc_h5
				_ = _883_cf0
				var _884_v55 _dafny.Int
				_ = _884_v55
				var _out21 _dafny.Int
				_ = _out21
				_out21 = Companion_Default___.M0(!(true), globalState)
				_884_v55 = _out21
				(globalState).F0 = Companion_Default___.SafeModuloInt((_this).F26(), (_this).F26())
				var _885_v56 _dafny.Sequence
				_ = _885_v56
				_885_v56 = _dafny.SeqOf(_dafny.SeqOf((_this).F25(), false, (_this).F25()))
				var _886_v57 _dafny.Array
				_ = _886_v57
				var _nwElement0_31 bool = _883_cf0
				_ = _nwElement0_31
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(23))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_31, 0)
				(_nw150).ArraySet1((_this).F25(), 1)
				(_nw150).ArraySet1((_this).F25(), 2)
				(_nw150).ArraySet1((_this).F25(), 3)
				(_nw150).ArraySet1((_this).F25(), 4)
				(_nw150).ArraySet1(!(_883_cf0), 5)
				(_nw150).ArraySet1((_this).F25(), 6)
				(_nw150).ArraySet1(_883_cf0, 7)
				(_nw150).ArraySet1((_this).F25(), 8)
				(_nw150).ArraySet1((p1).Cmp((_831_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))).Int()).(_dafny.Int)) != 0, 9)
				(_nw150).ArraySet1(_883_cf0, 10)
				(_nw150).ArraySet1(_883_cf0, 11)
				(_nw150).ArraySet1((_this).F25(), 12)
				(_nw150).ArraySet1(false, 13)
				(_nw150).ArraySet1(!(_883_cf0), 14)
				(_nw150).ArraySet1(_883_cf0, 15)
				(_nw150).ArraySet1((_this).F25(), 16)
				(_nw150).ArraySet1(_883_cf0, 17)
				(_nw150).ArraySet1(_883_cf0, 18)
				(_nw150).ArraySet1(true, 19)
				(_nw150).ArraySet1((p0).Contains((_dafny.SetOf(_dafny.IntOfUint32((_885_v56).Cardinality()))).Cardinality()), 20)
				(_nw150).ArraySet1(_883_cf0, 21)
				(_nw150).ArraySet1(_883_cf0, 22)
				_886_v57 = _nw150
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_886_v57), 0))
				_ = _index160
				(_886_v57).ArraySet1((_this).F25(), (_index160).Int())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_886_v57), 0))
				_ = _index161
				(_886_v57).ArraySet1(_dafny.Companion_Sequence_.Equal(_834_v24, _834_v24), (_index161).Int())
				var _887_v58 _dafny.Map
				_ = _887_v58
				_887_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F25())
				var _888_v59 *C1
				_ = _888_v59
				var _nw151 *C1 = New_C1_()
				_ = _nw151
				_nw151.Ctor__(_883_cf0)
				_888_v59 = _nw151
				var _889_v60 _dafny.Sequence
				_ = _889_v60
				_889_v60 = _dafny.SeqOf(_888_v59, _888_v59, _888_v59)
				var _890_v61 _dafny.Sequence
				_ = _890_v61
				_890_v61 = _dafny.SeqOf(_889_v60)
				_887_v58 = ((_887_v58).Update((_this).F25(), true)).Update(((_this).F26()).Cmp(_dafny.IntOfUint32((_890_v61).Cardinality())) == 0, !_dafny.Companion_Sequence_.Equal(_834_v24, _834_v24))
			}
			(globalState).F1 = true
			(globalState).F1 = (func() bool {
				if (_this).F25() {
					return (_this).F25()
				}
				return false
			})()
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_831_v23), 0))
			_ = _index162
			(_831_v23).ArraySet1((_this).F26(), (_index162).Int())
		}
		var _891_i8 _dafny.Int
		_ = _891_i8
		_891_i8 = _dafny.Zero
		{
			for !((_this).F25()) {
				{
					if (_891_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_891_i8 = (_891_i8).Plus(_dafny.One)
					var _892_v62 D0
					_ = _892_v62
					_892_v62 = Companion_D0_.Create_DC2_()
					var _893_v63 _dafny.Map
					_ = _893_v63
					_893_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _892_v62)
					var _894_v64 _dafny.Sequence
					_ = _894_v64
					_894_v64 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_893_v63).Cardinality(), (_this).F31()))
					(globalState).F0 = (_894_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.IntOfUint32((_894_v64).Cardinality()))).Uint32()).(_dafny.Int)
					var _895_v65 _dafny.Set
					_ = _895_v65
					_895_v65 = _dafny.SetOf(Companion_Default___.Fm2(_dafny.IntOfInt64(-334), (_this).F31(), globalState), p1)
					if !(!((_895_v65).IsDisjointFrom(_895_v65))) {
						var _896_v66 _dafny.Sequence
						_ = _896_v66
						_896_v66 = _dafny.SeqOf((_this).F25(), (_this).F25(), (_this).F25(), (_this).F25())
						var _897_v67 _dafny.Sequence
						_ = _897_v67
						_897_v67 = _dafny.UnicodeSeqOfUtf8Bytes("ynxpxku")
						var _898_v68 _dafny.Map
						_ = _898_v68
						_898_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _897_v67)
						var _899_v69 _dafny.Map
						_ = _899_v69
						_899_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_898_v68).Contains(_dafny.IntOfInt64(-3)) {
								return (_898_v68).Get(_dafny.IntOfInt64(-3)).(_dafny.Sequence)
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg49 _dafny.Int) interface{} {
									return coer49(arg49)
								}
							}(func(_900_i9 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('n')
							}))
						})()).Cardinality()), (_this).F25())
						var _901_v70 _dafny.Map
						_ = _901_v70
						_901_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)
						var _902_v71 _dafny.Sequence
						_ = _902_v71
						_902_v71 = _dafny.SeqOf(_901_v70, _901_v70)
						var _903_v72 _dafny.Sequence
						_ = _903_v72
						_903_v72 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_896_v66, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_899_v69).Cardinality()), _dafny.IntOfUint32((_896_v66).Cardinality()))).Uint32(), (_this).F25()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_896_v66, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_899_v69).Cardinality()), _dafny.IntOfUint32((_896_v66).Cardinality()))).Uint32(), (_this).F25())).Cardinality()))).Uint32(), (_this).F25()), _dafny.Companion_Sequence_.Update(_896_v66, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_902_v71).Cardinality()), _dafny.IntOfUint32((_896_v66).Cardinality()))).Uint32(), (_this).F25()))
						_903_v72 = _903_v72
						var _904_v73 _dafny.MultiSet
						_ = _904_v73
						_904_v73 = _dafny.MultiSetOf(_804_v0, _804_v0, (func() _dafny.CodePoint {
							if (_this).F25() {
								return _804_v0
							}
							return _804_v0
						})())
						var _905_v74 D0
						_ = _905_v74
						_905_v74 = Companion_D0_.Create_DC1_((_this).F31(), _897_v67, (_this).F25(), (_this).F31(), (_this).F25())
						var _906_v76 _dafny.Set
						_ = _906_v76
						_906_v76 = _dafny.SetOf((_this).F25(), true)
						var _907_v77 _dafny.Array
						_ = _907_v77
						var _nwElement0_32 _dafny.Int = Companion_Default___.SafeDivisionInt(p2, p1)
						_ = _nwElement0_32
						var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
						_ = _nw152
						(_nw152).ArraySet1(_nwElement0_32, 0)
						(_nw152).ArraySet1((func() _dafny.Int {
							if (_this).F25() {
								return (_this).F31()
							}
							return p1
						})(), 1)
						(_nw152).ArraySet1(p1, 2)
						(_nw152).ArraySet1((_this).F26(), 3)
						(_nw152).ArraySet1((_905_v74).Dtor_cf4(), 4)
						(_nw152).ArraySet1(_dafny.IntOfInt64(22), 5)
						(_nw152).ArraySet1(((_this).F26()).Times((_this).F31()), 6)
						(_nw152).ArraySet1(p2, 7)
						(_nw152).ArraySet1((func() _dafny.Map {
							var _coll20 = _dafny.NewMapBuilder()
							_ = _coll20
							for _iter20 := _dafny.Iterate((p0).Keys().Elements()); ; {
								_compr_20, _ok20 := _iter20()
								if !_ok20 {
									break
								}
								var _908_v75 _dafny.Int
								_908_v75 = interface{}(_compr_20).(_dafny.Int)
								if (p0).Contains(_908_v75) {
									_coll20.Add((_908_v75).Times((_this).F26()), !((_this).F25()))
								}
							}
							return _coll20.ToMap()
						}()).Cardinality(), 8)
						(_nw152).ArraySet1((_dafny.IntOfUint32((_894_v64).Cardinality())).Times(p2), 9)
						(_nw152).ArraySet1((_this).F31(), 10)
						(_nw152).ArraySet1(((_906_v76).Cardinality()).Minus((_this).F31()), 11)
						(_nw152).ArraySet1((_this).F31(), 12)
						_907_v77 = _nw152
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_907_v77), 0))
						_ = _index163
						(_907_v77).ArraySet1(p1, (_index163).Int())
						var _909_v78 D2
						_ = _909_v78
						_909_v78 = Companion_D2_.Create_DC7_((_this).F25(), (_this).F25(), p1)
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_907_v77), 0))
						_ = _index164
						var _rhs168 _dafny.MultiSet = _904_v73
						_ = _rhs168
						var _rhs169 _dafny.Int = ((_909_v78).Dtor_cf13()).Minus(p1)
						_ = _rhs169
						var _rhs170 _dafny.Int = (_this).F26()
						_ = _rhs170
						var _rhs171 bool = !(!((func() bool {
							if !((_this).F25()) {
								return (_this).F25()
							}
							return (_this).F25()
						})())) || ((_this).F25())
						_ = _rhs171
						var _rhs172 bool = (_this).Fm7(globalState)
						_ = _rhs172
						var _lhs154 *GlobalState = globalState
						_ = _lhs154
						var _lhs155 _dafny.Array = _907_v77
						_ = _lhs155
						var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_907_v77), 0))
						_ = _lhs156
						var _lhs157 *GlobalState = globalState
						_ = _lhs157
						_904_v73 = _rhs168
						_lhs154.F2 = _rhs169
						(_lhs155).ArraySet1(_rhs170, (_lhs156).Int())
						_lhs157.F1 = _rhs171
						r0 = _rhs172
						var _910_v79 _dafny.Sequence
						_ = _910_v79
						_910_v79 = _dafny.SeqOf(_897_v67, _897_v67)
						_897_v67 = _dafny.Companion_Sequence_.Concatenate(_897_v67, _dafny.Companion_Sequence_.Concatenate(_897_v67, _dafny.Companion_Sequence_.Update((_910_v79).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-414)), _dafny.IntOfUint32((_910_v79).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_897_v67, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_897_v67).Cardinality()), _dafny.IntOfUint32((_897_v67).Cardinality()))).Uint32(), _804_v0)).Cardinality()), _dafny.IntOfUint32(((_910_v79).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-414)), _dafny.IntOfUint32((_910_v79).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _804_v0)))
						var _911_v80 _dafny.MultiSet
						_ = _911_v80
						_911_v80 = _dafny.MultiSetOf((_this).F25())
						var _rhs173 _dafny.MultiSet = ((_911_v80).Update((_this).F25(), Companion_Default___.Abs(p2))).Difference(_911_v80)
						_ = _rhs173
						var _rhs174 bool = !(!((_this).F25()))
						_ = _rhs174
						var _rhs175 _dafny.Set = (_906_v76).Intersection(_906_v76)
						_ = _rhs175
						var _lhs158 *GlobalState = globalState
						_ = _lhs158
						_911_v80 = _rhs173
						_lhs158.F1 = _rhs174
						_906_v76 = _rhs175
						var _912_v81 D6
						_ = _912_v81
						_912_v81 = Companion_D6_.Create_DC13_(_dafny.IntOfInt64(990), p1, _dafny.SetOf((_this).F26(), (_907_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_907_v77), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_897_v67).Cardinality())))
						var _913_v82 _dafny.Map
						_ = _913_v82
						_913_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_804_v0, (_dafny.Zero).Minus((_912_v81).Dtor_cf21()))
						_913_v82 = _913_v82
					} else {
						var _914_v83 _dafny.Sequence
						_ = _914_v83
						_914_v83 = _dafny.UnicodeSeqOfUtf8Bytes("kg")
						var _915_v84 D0
						_ = _915_v84
						_915_v84 = Companion_D0_.Create_DC0_((_this).F25())
						var _916_v85 _dafny.Map
						_ = _916_v85
						_916_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v84, (_dafny.Zero).Minus((_this).F31()))
						var _917_v86 _dafny.Array
						_ = _917_v86
						var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw153
						_917_v86 = _nw153
						var _918_v87 T2
						_ = _918_v87
						var _nw154 *C0 = New_C0_()
						_ = _nw154
						_nw154.Ctor__(_914_v83, Companion_Default___.Fm5(true, _916_v85, _dafny.IntOfInt64(814), globalState), (_this).F26(), _917_v86, false)
						_918_v87 = _nw154
						var _919_v88 _dafny.Sequence
						_ = _919_v88
						_919_v88 = _dafny.SeqOf(false)
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_917_v86), 0))
						_ = _index165
						(_917_v86).ArraySet1((_919_v88).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_919_v88).Cardinality()))).Uint32()).(bool), (_index165).Int())
						var _920_v89 _dafny.Array
						_ = _920_v89
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(10)
						_ = _len0_24
						var _nw155 _dafny.Array
						_ = _nw155
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw155 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) _dafny.Map = (func(_921_v64 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
								return func(_922_i10 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_921_v64, true)
								}
							})(_894_v64)
							_ = _init24
							var _element0_24 = _init24(_dafny.Zero)
							_ = _element0_24
							_nw155 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
							(_nw155).ArraySet1(_element0_24, 0)
							var _nativeLen0_24 = (_len0_24).Int()
							_ = _nativeLen0_24
							for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
								(_nw155).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
							}
						}
						_920_v89 = _nw155
						var _923_v91 _dafny.Sequence
						_ = _923_v91
						_923_v91 = _dafny.SeqOf(_894_v64)
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_920_v89), 0))
						_ = _index166
						(_920_v89).ArraySet1(func() _dafny.Map {
							var _coll21 = _dafny.NewMapBuilder()
							_ = _coll21
							for _iter21 := _dafny.Iterate((_923_v91).Elements()); ; {
								_compr_21, _ok21 := _iter21()
								if !_ok21 {
									break
								}
								var _924_v90 _dafny.Sequence
								_924_v90 = interface{}(_compr_21).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_923_v91, _924_v90) {
									_coll21.Add(_924_v90, (_this).F25())
								}
							}
							return _coll21.ToMap()
						}(), (_index166).Int())
						var _925_v92 _dafny.Set
						_ = _925_v92
						_925_v92 = _dafny.SetOf(_918_v87.F34())
						var _926_v93 _dafny.Map
						_ = _926_v93
						_926_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf(p2, _dafny.IntOfUint32((_914_v83).Cardinality())))
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_917_v86), 0))
						_ = _index167
						var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_920_v89), 0))
						_ = _index168
						var _rhs176 _dafny.Int = (_this).F31()
						_ = _rhs176
						var _rhs177 T2 = _918_v87
						_ = _rhs177
						var _rhs178 bool = (_this).F25()
						_ = _rhs178
						var _rhs179 _dafny.Int = ((func() _dafny.Int {
							if (_this).F25() {
								return p2
							}
							return (Companion_D2_.Create_DC7_(true, _918_v87.F34(), (_925_v92).Cardinality())).Dtor_cf13()
						})()).Times(p2)
						_ = _rhs179
						var _rhs180 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
							if (_926_v93).Contains(_dafny.IntOfInt64(-64)) {
								return (_926_v93).Get(_dafny.IntOfInt64(-64)).(_dafny.Sequence)
							}
							return _894_v64
						})(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_894_v64, (_this).F25()))
						_ = _rhs180
						var _lhs159 *GlobalState = globalState
						_ = _lhs159
						var _lhs160 _dafny.Array = _917_v86
						_ = _lhs160
						var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_917_v86), 0))
						_ = _lhs161
						var _lhs162 *GlobalState = globalState
						_ = _lhs162
						var _lhs163 _dafny.Array = _920_v89
						_ = _lhs163
						var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_920_v89), 0))
						_ = _lhs164
						_lhs159.F0 = _rhs176
						_918_v87 = _rhs177
						(_lhs160).ArraySet1(_rhs178, (_lhs161).Int())
						_lhs162.F2 = _rhs179
						(_lhs163).ArraySet1(_rhs180, (_lhs164).Int())
						var _927_v94 _dafny.Map
						_ = _927_v94
						_927_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v83, _914_v83)
						_927_v94 = _927_v94
						var _928_v95 _dafny.Array
						_ = _928_v95
						var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
						_ = _nw156
						_928_v95 = _nw156
						var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_928_v95), 0))
						_ = _index169
						(_928_v95).ArraySet1(p2, (_index169).Int())
						var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_928_v95), 0))
						_ = _index170
						(_928_v95).ArraySet1((_this).F31(), (_index170).Int())
						var _929_v96 D1
						_ = _929_v96
						_929_v96 = Companion_D1_.Create_DC4_((_917_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_917_v86), 0))).Int()).(bool), true)
						(globalState).F0 = Companion_Default___.Fm2((func() _dafny.Int {
							if (_929_v96).Dtor_cf7() {
								return (_928_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_928_v95), 0))).Int()).(_dafny.Int)
							}
							return p1
						})(), (_928_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_928_v95), 0))).Int()).(_dafny.Int), globalState)
						var _930_v97 _dafny.Set
						_ = _930_v97
						_930_v97 = _dafny.SetOf(_895_v65)
						var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_917_v86), 0))
						_ = _index171
						(_917_v86).ArraySet1((_930_v97).IsDisjointFrom(_930_v97), (_index171).Int())
					}
					var _931_v98 _dafny.Sequence
					_ = _931_v98
					_931_v98 = _dafny.UnicodeSeqOfUtf8Bytes("kbigaspv")
					var _932_v99 _dafny.Set
					_ = _932_v99
					_932_v99 = _dafny.SetOf((_this).F25())
					var _933_v100 _dafny.Map
					_ = _933_v100
					_933_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_931_v98).Cardinality()), _932_v99)
					var _934_v101 _dafny.Sequence
					_ = _934_v101
					_934_v101 = _dafny.SeqOf((_this).F25())
					_933_v100 = (_933_v100).Update((_dafny.MultiSetOf(_931_v98)).Cardinality(), (_dafny.SetOf((_934_v101).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_934_v101).Cardinality()))).Uint32()).(bool), (_this).F25())).Difference(Companion_Default___.Fm29((_this).Fm7(globalState), p0, _804_v0, (_this).F26(), globalState)))
					_932_v99 = _932_v99
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _935_v102 _dafny.Int
		_ = _935_v102
		var _out22 _dafny.Int
		_ = _out22
		_out22 = Companion_Default___.M0((_this).F25(), globalState)
		_935_v102 = _out22
		var _936_v103 _dafny.MultiSet
		_ = _936_v103
		_936_v103 = _dafny.MultiSetOf(Companion_Default___.Fm2((_dafny.MultiSetOf(p2)).Cardinality(), (_this).F31(), globalState))
		var _937_v104 _dafny.Sequence
		_ = _937_v104
		_937_v104 = _dafny.SeqOf((_this).F26())
		var _938_v105 _dafny.MultiSet
		_ = _938_v105
		_938_v105 = _dafny.MultiSetOf((_937_v104).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_937_v104).Cardinality()))).Uint32()).(_dafny.Int))
		var _939_v106 _dafny.Array
		_ = _939_v106
		var _nwElement0_33 _dafny.MultiSet = _936_v103
		_ = _nwElement0_33
		var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(17))
		_ = _nw157
		(_nw157).ArraySet1(_nwElement0_33, 0)
		(_nw157).ArraySet1(_936_v103, 1)
		(_nw157).ArraySet1(_938_v105, 2)
		(_nw157).ArraySet1(_936_v103, 3)
		(_nw157).ArraySet1(_936_v103, 4)
		(_nw157).ArraySet1(_936_v103, 5)
		(_nw157).ArraySet1(Companion_Default___.Fm30(globalState), 6)
		(_nw157).ArraySet1(_936_v103, 7)
		(_nw157).ArraySet1(_936_v103, 8)
		(_nw157).ArraySet1(Companion_Default___.Fm30(globalState), 9)
		(_nw157).ArraySet1(_936_v103, 10)
		(_nw157).ArraySet1(_938_v105, 11)
		(_nw157).ArraySet1(_936_v103, 12)
		(_nw157).ArraySet1((_938_v105).Update(p1, Companion_Default___.Abs((_this).F26())), 13)
		(_nw157).ArraySet1(_936_v103, 14)
		(_nw157).ArraySet1(_936_v103, 15)
		(_nw157).ArraySet1(_936_v103, 16)
		_939_v106 = _nw157
		var _940_v107 _dafny.Sequence
		_ = _940_v107
		_940_v107 = _dafny.UnicodeSeqOfUtf8Bytes("nnxuic")
		var _941_v108 D0
		_ = _941_v108
		_941_v108 = Companion_D0_.Create_DC1_(_935_v102, _940_v107, true, p1, (_this).F25())
		var _942_v109 _dafny.MultiSet
		_ = _942_v109
		_942_v109 = _dafny.MultiSetOf(true, (_this).F25(), false)
		var _rhs181 _dafny.Array = _939_v106
		_ = _rhs181
		var _rhs182 _dafny.Int = (_941_v108).Dtor_cf4()
		_ = _rhs182
		var _rhs183 bool = (_942_v109).IsSubsetOf(_942_v109)
		_ = _rhs183
		var _rhs184 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_937_v104, _dafny.Companion_Sequence_.Concatenate(_937_v104, _937_v104))
		_ = _rhs184
		var _lhs165 *GlobalState = globalState
		_ = _lhs165
		var _lhs166 *GlobalState = globalState
		_ = _lhs166
		_939_v106 = _rhs181
		_lhs165.F2 = _rhs182
		_lhs166.F1 = _rhs183
		_937_v104 = _rhs184
		var _943_v110 _dafny.Array
		_ = _943_v110
		var _nw158 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw158
		_943_v110 = _nw158
		var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_943_v110), 0))
		_ = _index172
		(_943_v110).ArraySet1((_this).F25(), (_index172).Int())
		var _944_v111 D0
		_ = _944_v111
		_944_v111 = Companion_D0_.Create_DC0_((_this).F25())
		var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_943_v110), 0))
		_ = _index173
		(_943_v110).ArraySet1(func(_source14 D0) bool {
			if _source14.Is_DC1() {
				var _945___mcc_h6 _dafny.Int = _source14.Get_().(D0_DC1).Cf1
				_ = _945___mcc_h6
				var _946___mcc_h7 _dafny.Sequence = _source14.Get_().(D0_DC1).Cf2
				_ = _946___mcc_h7
				var _947___mcc_h8 bool = _source14.Get_().(D0_DC1).Cf3
				_ = _947___mcc_h8
				var _948___mcc_h9 _dafny.Int = _source14.Get_().(D0_DC1).Cf4
				_ = _948___mcc_h9
				var _949___mcc_h10 bool = _source14.Get_().(D0_DC1).Cf5
				_ = _949___mcc_h10
				var _950_cf5 bool = _949___mcc_h10
				_ = _950_cf5
				var _951_cf4 _dafny.Int = _948___mcc_h9
				_ = _951_cf4
				var _952_cf3 bool = _947___mcc_h8
				_ = _952_cf3
				var _953_cf2 _dafny.Sequence = _946___mcc_h7
				_ = _953_cf2
				var _954_cf1 _dafny.Int = _945___mcc_h6
				_ = _954_cf1
				return (_this).F25()
			} else if _source14.Is_DC2() {
				return (_this).F25()
			} else {
				var _955___mcc_h11 bool = _source14.Get_().(D0_DC0).Cf0
				_ = _955___mcc_h11
				var _956_cf0 bool = _955___mcc_h11
				_ = _956_cf0
				return _956_cf0
			}
		}(_944_v111), (_index173).Int())
		r0 = (_943_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_943_v110), 0))).Int()).(bool)
		return r0
	}
}
func (_this *C2) F31() _dafny.Int {
	{
		return _this._f31
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f25 bool
	_f26 _dafny.Int
	_f29 _dafny.Sequence
	_f30 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f25 = false
	_this._f26 = _dafny.Zero
	_this._f29 = _dafny.EmptySeq
	_this._f30 = false
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

func (_this *C3) F25() bool {
	return _this._f25
}
func (_this *C3) F26() _dafny.Int {
	return _this._f26
}
func (_this *C3) Ctor__(f29 _dafny.Sequence, f30 bool, f25 bool, f26 _dafny.Int) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f25 = f25
		(_this)._f26 = f26
	}
}
func (_this *C3) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F26()).Times(((_this).F26()).Minus((_this).F26()))
	}
}
func (_this *C3) Fm7(globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(926)).Cmp((_this).F26()) != 0
	}
}
func (_this *C3) Fm9(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_this).F26())
	}
}
func (_this *C3) M1(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _957_v0 _dafny.CodePoint
		_ = _957_v0
		_957_v0 = _dafny.CodePoint('u')
		_957_v0 = (Companion_Default___.Fm10((_this).F25(), p2, globalState)).Dtor_cf6()
		var _958_v1 _dafny.Map
		_ = _958_v1
		_958_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), p1)
		if (func() bool {
			if (func() bool {
				if (func() bool {
					if (_958_v1).Contains(true) {
						return (_958_v1).Get(true).(bool)
					}
					return (_this).F25()
				})() {
					return true
				}
				return (_this).F30()
			})() {
				return (_this).F25()
			}
			return p2
		})() {
			var _959_v2 _dafny.Map
			_ = _959_v2
			_959_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), !((_this).F30()) || ((_this).F30()))
			_959_v2 = (_959_v2).Update((_this).F26(), (_this).F25())
			var _960_v3 _dafny.Array
			_ = _960_v3
			var _nw159 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw159
			_960_v3 = _nw159
			var _961_v4 _dafny.Sequence
			_ = _961_v4
			_961_v4 = _dafny.SeqOf(_960_v3)
			var _962_v5 _dafny.Array
			_ = _962_v5
			var _nwElement0_34 _dafny.Array = _960_v3
			_ = _nwElement0_34
			var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(26))
			_ = _nw160
			(_nw160).ArraySet1(_nwElement0_34, 0)
			(_nw160).ArraySet1(_960_v3, 1)
			(_nw160).ArraySet1(_960_v3, 2)
			(_nw160).ArraySet1(_960_v3, 3)
			(_nw160).ArraySet1(_960_v3, 4)
			(_nw160).ArraySet1(_960_v3, 5)
			(_nw160).ArraySet1(_960_v3, 6)
			(_nw160).ArraySet1(_960_v3, 7)
			(_nw160).ArraySet1(_960_v3, 8)
			(_nw160).ArraySet1(_960_v3, 9)
			(_nw160).ArraySet1(_960_v3, 10)
			(_nw160).ArraySet1(_960_v3, 11)
			(_nw160).ArraySet1(_960_v3, 12)
			(_nw160).ArraySet1(_960_v3, 13)
			(_nw160).ArraySet1(_960_v3, 14)
			(_nw160).ArraySet1(_960_v3, 15)
			(_nw160).ArraySet1(_960_v3, 16)
			(_nw160).ArraySet1(_960_v3, 17)
			(_nw160).ArraySet1(_960_v3, 18)
			(_nw160).ArraySet1(_960_v3, 19)
			(_nw160).ArraySet1(_960_v3, 20)
			(_nw160).ArraySet1(_960_v3, 21)
			(_nw160).ArraySet1(_960_v3, 22)
			(_nw160).ArraySet1((func() _dafny.Array {
				if (_this).F30() {
					return _960_v3
				}
				return _960_v3
			})(), 23)
			(_nw160).ArraySet1((_961_v4).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_961_v4).Cardinality()))).Uint32()).(_dafny.Array), 24)
			(_nw160).ArraySet1((_961_v4).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_961_v4).Cardinality()))).Uint32()).(_dafny.Array), 25)
			_962_v5 = _nw160
			var _963_v6 _dafny.Set
			_ = _963_v6
			_963_v6 = _dafny.SetOf((_this).F30(), (_this).F25())
			var _964_v7 D1
			_ = _964_v7
			_964_v7 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_957_v0))
			var _965_v8 D1
			_ = _965_v8
			_965_v8 = Companion_D1_.Create_DC5_(_964_v7)
			var _pat_let_tv25 = _964_v7
			_ = _pat_let_tv25
			var _966_v9 _dafny.MultiSet
			_ = _966_v9
			_966_v9 = _dafny.MultiSetOf(Companion_D1_.Create_DC5_(_964_v7), func(_pat_let24_0 D1) D1 {
				return func(_967_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let25_0 D1) D1 {
						return func(_968_dt__update_hcf9_h0 D1) D1 {
							return Companion_D1_.Create_DC5_(_968_dt__update_hcf9_h0)
						}(_pat_let25_0)
					}(_pat_let_tv25)
				}(_pat_let24_0)
			}(_965_v8), _965_v8)
			var _rhs185 bool = !((_963_v6).IsDisjointFrom(_963_v6)) || (p2)
			_ = _rhs185
			var _rhs186 _dafny.Int = (_this).F26()
			_ = _rhs186
			var _rhs187 _dafny.Array = _962_v5
			_ = _rhs187
			var _rhs188 bool = (func() bool {
				if (_959_v2).Contains((func() _dafny.Int {
					if (_966_v9).Contains(_965_v8) {
						return (_966_v9).Multiplicity(_965_v8)
					}
					return (_this).F26()
				})()) {
					return (_959_v2).Get((func() _dafny.Int {
						if (_966_v9).Contains(_965_v8) {
							return (_966_v9).Multiplicity(_965_v8)
						}
						return (_this).F26()
					})()).(bool)
				}
				return ((_this).F25()) || (p1)
			})()
			_ = _rhs188
			var _lhs167 *GlobalState = globalState
			_ = _lhs167
			var _lhs168 *GlobalState = globalState
			_ = _lhs168
			r0 = _rhs185
			_lhs167.F2 = _rhs186
			_962_v5 = _rhs187
			_lhs168.F1 = _rhs188
			var _969_v10 _dafny.Array
			_ = _969_v10
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_25
			var _nw161 _dafny.Array
			_ = _nw161
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw161 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Sequence = func(_970_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(true, true)
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw161 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw161).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw161).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_969_v10 = _nw161
			var _971_v11 _dafny.Sequence
			_ = _971_v11
			_971_v11 = _dafny.SeqOf(p1)
			var _972_v12 D0
			_ = _972_v12
			_972_v12 = Companion_D0_.Create_DC0_((_this).F30())
			var _973_v13 _dafny.Map
			_ = _973_v13
			_973_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_972_v12, (_this).F26())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_969_v10), 0))
			_ = _index174
			(_969_v10).ArraySet1(_dafny.Companion_Sequence_.Update(_971_v11, (Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_971_v11).Cardinality()))).Uint32(), Companion_Default___.Fm5((_this).F30(), _973_v13, (_this).F26(), globalState)), (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_969_v10), 0))
			_ = _index175
			(_969_v10).ArraySet1(_dafny.SeqOf(p1), (_index175).Int())
			var _974_v14 D0
			_ = _974_v14
			_974_v14 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(689))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_975_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_976_i1 _dafny.Int) _dafny.CodePoint {
					return _975_v0
				}
			})(_957_v0)))).Cardinality()), (_this).F29(), (_this).F30(), (_this).F26(), p2)
			_957_v0 = Companion_Default___.Fm4(_974_v14, (_this).F25(), (_dafny.Zero).Minus((_this).F26()), globalState)
			(globalState).F0 = (_this).F26()
		} else {
			var _977_v15 _dafny.Sequence
			_ = _977_v15
			_977_v15 = _dafny.SeqOf(p2)
			var _978_v16 T0
			_ = _978_v16
			var _nw162 *C2 = New_C2_()
			_ = _nw162
			_nw162.Ctor__((_dafny.Zero).Minus((_this).F26()), !(!((_this).F25())), (func() _dafny.Int {
				if p2 {
					return (_this).F26()
				}
				return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_977_v15).Cardinality()), (p0).Cardinality())).Cardinality())
			})())
			_978_v16 = _nw162
			_978_v16 = (func() T0 {
				if p1 {
					return _978_v16
				}
				return _978_v16
			})()
			_977_v15 = _977_v15
			(globalState).F1 = (Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_978_v16).F26()), (_this).F29(), true, (_978_v16).F26(), !((func() bool {
				if (_958_v1).Contains(true) {
					return (_958_v1).Get(true).(bool)
				}
				return p1
			})()))).Dtor_cf5()
			var _rhs189 bool = (_this).F30()
			_ = _rhs189
			var _rhs190 _dafny.Int = (_dafny.Zero).Minus((_this).F26())
			_ = _rhs190
			var _lhs169 *GlobalState = globalState
			_ = _lhs169
			var _lhs170 *GlobalState = globalState
			_ = _lhs170
			_lhs169.F1 = _rhs189
			_lhs170.F2 = _rhs190
			var _979_v17 _dafny.Set
			_ = _979_v17
			_979_v17 = _dafny.SetOf(Companion_Default___.Fm2((_this).F26(), (_this).F26(), globalState), (_this).F26())
			var _980_v18 D6
			_ = _980_v18
			_980_v18 = Companion_D6_.Create_DC13_((_dafny.Zero).Minus((_this).F26()), (_978_v16).F26(), _979_v17)
			var _981_v19 _dafny.Map
			_ = _981_v19
			_981_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_980_v18).Dtor_cf21(), (_this).F26())
			_981_v19 = _981_v19
		}
		var _982_i2 _dafny.Int
		_ = _982_i2
		_982_i2 = _dafny.Zero
		{
			for !(((_this).F26()).Cmp(((_this).F26()).Minus((_this).F26())) == 0) {
				{
					if (_982_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_982_i2 = (_982_i2).Plus(_dafny.One)
					var _983_v20 _dafny.Array
					_ = _983_v20
					var _nw163 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
					_ = _nw163
					_983_v20 = _nw163
					var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))
					_ = _index176
					(_983_v20).ArraySet1((_this).F25(), (_index176).Int())
					var _984_v21 _dafny.Sequence
					_ = _984_v21
					_984_v21 = _dafny.SeqOf(false)
					var _985_v22 _dafny.Sequence
					_ = _985_v22
					_985_v22 = _dafny.SeqOf(_984_v21, _984_v21)
					var _986_v23 _dafny.MultiSet
					_ = _986_v23
					_986_v23 = _dafny.MultiSetOf((_985_v22).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_985_v22).Cardinality()))).Uint32()).(_dafny.Sequence), _984_v21, _984_v21)
					var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))
					_ = _index177
					(_983_v20).ArraySet1(!(!(_986_v23).Contains(_984_v21)), (_index177).Int())
					if !(p1) {
						var _987_v24 _dafny.Map
						_ = _987_v24
						_987_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), ((_this).F26()).Plus((_this).F26()))
						var _988_v25 _dafny.Array
						_ = _988_v25
						var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
						_ = _nw164
						_988_v25 = _nw164
						var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))
						_ = _index178
						(_988_v25).ArraySet1(Companion_Default___.Fm3(_dafny.IntOfInt64(973), (_983_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))).Int()).(bool), (_this).F26(), (_this).F30(), globalState), (_index178).Int())
						var _989_v26 _dafny.Sequence
						_ = _989_v26
						_989_v26 = _dafny.SeqOf((_this).F26(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()))
						var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))
						_ = _index179
						var _rhs191 _dafny.Map = _987_v24
						_ = _rhs191
						var _rhs192 _dafny.Sequence = _989_v26
						_ = _rhs192
						var _lhs171 _dafny.Array = _988_v25
						_ = _lhs171
						var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))
						_ = _lhs172
						_987_v24 = _rhs191
						(_lhs171).ArraySet1(_rhs192, (_lhs172).Int())
						var _990_v27 _dafny.Array
						_ = _990_v27
						var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
						_ = _nw165
						_990_v27 = _nw165
						var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_990_v27), 0))
						_ = _index180
						(_990_v27).ArraySet1CodePoint(_957_v0, (_index180).Int())
						var _991_v28 _dafny.Set
						_ = _991_v28
						_991_v28 = _dafny.SetOf(((_this).F26()).Cmp((_this).F26()) >= 0)
						var _pat_let_tv26 = _957_v0
						_ = _pat_let_tv26
						var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_990_v27), 0))
						_ = _index181
						var _rhs193 _dafny.Array = _983_v20
						_ = _rhs193
						var _rhs194 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_984_v21, _984_v21)
						_ = _rhs194
						var _rhs195 _dafny.CodePoint = (func(_pat_let26_0 D4) D4 {
							return func(_992_dt__update__tmp_h1 D4) D4 {
								return func(_pat_let27_0 _dafny.CodePoint) D4 {
									return func(_993_dt__update_hcf17_h0 _dafny.CodePoint) D4 {
										return Companion_D4_.Create_DC10_((_992_dt__update__tmp_h1).Dtor_cf16(), _993_dt__update_hcf17_h0)
									}(_pat_let27_0)
								}(_pat_let_tv26)
							}(_pat_let26_0)
						}(Companion_D4_.Create_DC10_(_957_v0, _957_v0))).Dtor_cf17()
						_ = _rhs195
						var _rhs196 _dafny.Int = (_this).F26()
						_ = _rhs196
						var _rhs197 _dafny.Set = _991_v28
						_ = _rhs197
						var _lhs173 _dafny.Array = _990_v27
						_ = _lhs173
						var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_990_v27), 0))
						_ = _lhs174
						var _lhs175 *GlobalState = globalState
						_ = _lhs175
						_983_v20 = _rhs193
						_984_v21 = _rhs194
						(_lhs173).ArraySet1CodePoint(_rhs195, (_lhs174).Int())
						_lhs175.F0 = _rhs196
						_991_v28 = _rhs197
						var _994_v29 _dafny.Map
						_ = _994_v29
						_994_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_983_v20, (_983_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))).Int()).(bool))
						var _995_v30 _dafny.Sequence
						_ = _995_v30
						_995_v30 = _dafny.SeqOf(_994_v29, _994_v29, _994_v29, _994_v29, _994_v29)
						var _996_v31 *C2
						_ = _996_v31
						var _nw166 *C2 = New_C2_()
						_ = _nw166
						_nw166.Ctor__((_this).F26(), !((_this).F25()) || (false), (((_995_v30).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_995_v30).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Times(_dafny.IntOfUint32((_984_v21).Cardinality())))
						_996_v31 = _nw166
						_957_v0 = (_990_v27).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_990_v27), 0))).Int())
						var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))
						_ = _index182
						(_988_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
							if p1 {
								return _989_v26
							}
							return (_988_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))).Int()).(_dafny.Sequence)
						})(), (_988_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_988_v25), 0))).Int()).(_dafny.Sequence)), (_index182).Int())
					} else {
						var _997_v32 D0
						_ = _997_v32
						_997_v32 = Companion_D0_.Create_DC0_(p2)
						var _998_v33 _dafny.Map
						_ = _998_v33
						_998_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_997_v32, _dafny.IntOfInt64(649))
						var _999_v34 _dafny.Map
						_ = _999_v34
						_999_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(326), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F29()).Cardinality())))
						var _1000_v35 _dafny.MultiSet
						_ = _1000_v35
						_1000_v35 = _dafny.MultiSetOf((_983_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))).Int()).(bool), Companion_Default___.Fm5((_this).F25(), _998_v33, (func() _dafny.Int {
							if (_999_v34).Contains((_this).F26()) {
								return (_999_v34).Get((_this).F26()).(_dafny.Int)
							}
							return (_this).F26()
						})(), globalState))
						(globalState).F1 = !((_1000_v35).Contains(true))
						var _1001_v36 _dafny.Array
						_ = _1001_v36
						var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
						_ = _nw167
						_1001_v36 = _nw167
						var _1002_v37 _dafny.Sequence
						_ = _1002_v37
						_1002_v37 = _dafny.SeqOf(_1001_v36)
						_1001_v36 = (func() _dafny.Array {
							if Companion_Default___.Fm5(p1, _998_v33, (_this).F26(), globalState) {
								return _1001_v36
							}
							return (_1002_v37).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1002_v37).Cardinality()))).Uint32()).(_dafny.Array)
						})()
						var _1003_v38 _dafny.Array
						_ = _1003_v38
						var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
						_ = _nw168
						_1003_v38 = _nw168
						var _1004_v39 _dafny.Map
						_ = _1004_v39
						_1004_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _1000_v35)
						var _1005_v40 D6
						_ = _1005_v40
						_1005_v40 = Companion_D6_.Create_DC14_((_this).F26(), (_this).F26(), (func() _dafny.MultiSet {
							if (_1004_v39).Contains(_dafny.IntOfInt64(796)) {
								return (_1004_v39).Get(_dafny.IntOfInt64(796)).(_dafny.MultiSet)
							}
							return (_1000_v35).Update((_this).F25(), Companion_Default___.Abs((_this).F26()))
						})())
						var _1006_v41 D9
						_ = _1006_v41
						_1006_v41 = Companion_D9_.Create_DC19_(_1005_v40)
						var _1007_v42 D9
						_ = _1007_v42
						_1007_v42 = Companion_D9_.Create_DC20_(_1006_v41)
						var _1008_v43 D0
						_ = _1008_v43
						_1008_v43 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-977), (_this).F29(), p1, (_this).F26(), (_984_v21).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_984_v21).Cardinality()))).Uint32()).(bool))
						var _pat_let_tv27 = p1
						_ = _pat_let_tv27
						var _pat_let_tv28 = p1
						_ = _pat_let_tv28
						var _1009_v44 _dafny.Map
						_ = _1009_v44
						_1009_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v42, _dafny.SetOf(_957_v0, Companion_Default___.Fm4(func(_pat_let28_0 D0) D0 {
							return func(_1010_dt__update__tmp_h2 D0) D0 {
								return func(_pat_let29_0 bool) D0 {
									return func(_1011_dt__update_hcf3_h0 bool) D0 {
										return func(_pat_let30_0 _dafny.Int) D0 {
											return func(_1012_dt__update_hcf4_h0 _dafny.Int) D0 {
												return func(_pat_let31_0 bool) D0 {
													return func(_1013_dt__update_hcf5_h0 bool) D0 {
														return Companion_D0_.Create_DC1_((_1010_dt__update__tmp_h2).Dtor_cf1(), (_1010_dt__update__tmp_h2).Dtor_cf2(), _1011_dt__update_hcf3_h0, _1012_dt__update_hcf4_h0, _1013_dt__update_hcf5_h0)
													}(_pat_let31_0)
												}(_pat_let_tv28)
											}(_pat_let30_0)
										}((_this).F26())
									}(_pat_let29_0)
								}(_pat_let_tv27)
							}(_pat_let28_0)
						}(_1008_v43), (_983_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))).Int()).(bool), (_this).F26(), globalState), _957_v0, _957_v0))
						var _1014_v45 _dafny.Set
						_ = _1014_v45
						_1014_v45 = _dafny.SetOf(_957_v0, _957_v0)
						var _pat_let_tv29 = _1007_v42
						_ = _pat_let_tv29
						var _pat_let_tv30 = _1007_v42
						_ = _pat_let_tv30
						var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1003_v38), 0))
						_ = _index183
						(_1003_v38).ArraySet1((func() _dafny.Set {
							if (_1009_v44).Contains(func(_pat_let32_0 D9) D9 {
								return func(_1015_dt__update__tmp_h3 D9) D9 {
									return func(_pat_let33_0 D9) D9 {
										return func(_1016_dt__update_hcf32_h0 D9) D9 {
											return Companion_D9_.Create_DC20_(_1016_dt__update_hcf32_h0)
										}(_pat_let33_0)
									}((_pat_let_tv29).Dtor_cf32())
								}(_pat_let32_0)
							}(_1007_v42)) {
								return (_1009_v44).Get(func(_pat_let34_0 D9) D9 {
									return func(_1017_dt__update__tmp_h4 D9) D9 {
										return func(_pat_let35_0 D9) D9 {
											return func(_1018_dt__update_hcf32_h1 D9) D9 {
												return Companion_D9_.Create_DC20_(_1018_dt__update_hcf32_h1)
											}(_pat_let35_0)
										}((_pat_let_tv30).Dtor_cf32())
									}(_pat_let34_0)
								}(_1007_v42)).(_dafny.Set)
							}
							return _1014_v45
						})(), (_index183).Int())
						var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1003_v38), 0))
						_ = _index184
						(_1003_v38).ArraySet1(Companion_Default___.Fm31(globalState), (_index184).Int())
						var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1001_v36), 0))
						_ = _index185
						(_1001_v36).ArraySet1((_this).F26(), (_index185).Int())
						var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1001_v36), 0))
						_ = _index186
						(_1001_v36).ArraySet1(_dafny.IntOfInt64(625), (_index186).Int())
						var _1019_v46 _dafny.Sequence
						_ = _1019_v46
						_1019_v46 = _dafny.SeqOf((_this).F26())
						var _1020_v47 _dafny.Set
						_ = _1020_v47
						_1020_v47 = _dafny.SetOf((_this).Fm7(globalState), p1)
						var _1021_v48 _dafny.Array
						_ = _1021_v48
						var _nwElement0_35 _dafny.Sequence = _1019_v46
						_ = _nwElement0_35
						var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(21))
						_ = _nw169
						(_nw169).ArraySet1(_nwElement0_35, 0)
						(_nw169).ArraySet1(Companion_Default___.Fm3((_1001_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1001_v36), 0))).Int()).(_dafny.Int), false, (_this).F26(), true, globalState), 1)
						(_nw169).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1019_v46, _1019_v46), 2)
						(_nw169).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(352)), 3)
						(_nw169).ArraySet1(_1019_v46, 4)
						(_nw169).ArraySet1(_1019_v46, 5)
						(_nw169).ArraySet1(_1019_v46, 6)
						(_nw169).ArraySet1(_1019_v46, 7)
						(_nw169).ArraySet1(_1019_v46, 8)
						(_nw169).ArraySet1(Companion_Default___.Fm3((_this).F26(), p2, (_1020_v47).Cardinality(), false, globalState), 9)
						(_nw169).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26(), (_1001_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1001_v36), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(784)), _1019_v46), 10)
						(_nw169).ArraySet1(_1019_v46, 11)
						(_nw169).ArraySet1(_1019_v46, 12)
						(_nw169).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26(), (_this).F26()), _1019_v46), 13)
						(_nw169).ArraySet1(_1019_v46, 14)
						(_nw169).ArraySet1(_dafny.Companion_Sequence_.Update(_1019_v46, (Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1019_v46).Cardinality()))).Uint32(), (_this).F26()), 15)
						(_nw169).ArraySet1(_1019_v46, 16)
						(_nw169).ArraySet1(_1019_v46, 17)
						(_nw169).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus((_this).F26())), 18)
						(_nw169).ArraySet1(_1019_v46, 19)
						(_nw169).ArraySet1(_dafny.SeqOf((_1001_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1001_v36), 0))).Int()).(_dafny.Int)), 20)
						_1021_v48 = _nw169
						var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1021_v48), 0))
						_ = _index187
						(_1021_v48).ArraySet1(_1019_v46, (_index187).Int())
						var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1021_v48), 0))
						_ = _index188
						(_1021_v48).ArraySet1(_1019_v46, (_index188).Int())
					}
					var _1022_v49 *C2
					_ = _1022_v49
					var _nw170 *C2 = New_C2_()
					_ = _nw170
					_nw170.Ctor__(_dafny.IntOfInt64(144), (p2) == ((_this).F25()), ((_this).F26()).Times((_this).F26()))
					_1022_v49 = _nw170
					var _1023_v50 _dafny.Map
					_ = _1023_v50
					_1023_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_983_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_983_v20), 0))).Int()).(bool), (_this).F26())
					_1023_v50 = (_1023_v50).Update(p1, (_dafny.Zero).Minus((_1022_v49).F31()))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		(globalState).F0 = (_this).F26()
		var _1024_i3 _dafny.Int
		_ = _1024_i3
		_1024_i3 = _dafny.Zero
		{
			for (_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt((_this).F26(), Companion_Default___.Fm2((_this).F26(), (_this).F26(), globalState))) >= 0 {
				{
					if (_1024_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1024_i3 = (_1024_i3).Plus(_dafny.One)
					var _1025_v51 _dafny.Int
					_ = _1025_v51
					var _out23 _dafny.Int
					_ = _out23
					_out23 = Companion_Default___.M0((_this).F25(), globalState)
					_1025_v51 = _out23
					(globalState).F0 = _dafny.IntOfInt64(-983)
					_958_v1 = (_958_v1).Update((_this).F25(), p2)
					_957_v0 = _957_v0
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _hi9 _dafny.Int = (_this).F26()
		_ = _hi9
		for _1026_i4 := ((_this).F26()).Times((_dafny.Zero).Minus((_this).F26())); _1026_i4.Cmp(_hi9) < 0; _1026_i4 = _1026_i4.Plus(_dafny.One) {
			var _1027_v52 _dafny.Array
			_ = _1027_v52
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_26
			var _nw171 _dafny.Array
			_ = _nw171
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw171 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = (func(_1028_p1 bool) func(_dafny.Int) bool {
					return func(_1029_i5 _dafny.Int) bool {
						return _1028_p1
					}
				})(p1)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw171 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw171).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw171).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1027_v52 = _nw171
			var _1030_v53 _dafny.Sequence
			_ = _1030_v53
			_1030_v53 = _dafny.SeqOf(_1027_v52)
			var _1031_v54 _dafny.Map
			_ = _1031_v54
			_1031_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1027_v52)
			_1030_v53 = _dafny.SeqOf((func() _dafny.Array {
				if (_1031_v54).Contains(!(!(p2))) {
					return (_1031_v54).Get(!(!(p2))).(_dafny.Array)
				}
				return _1027_v52
			})(), _1027_v52, _1027_v52, _1027_v52)
			(globalState).F0 = (_1026_i4).Plus((_this).F26())
			var _1032_v55 _dafny.Array
			_ = _1032_v55
			var _nw172 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(17))
			_ = _nw172
			_1032_v55 = _nw172
			var _1033_v56 _dafny.Map
			_ = _1033_v56
			_1033_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F26())
			var _1034_v57 _dafny.MultiSet
			_ = _1034_v57
			_1034_v57 = _dafny.MultiSetOf(_1033_v56, _1033_v56)
			var _1035_v58 _dafny.MultiSet
			_ = _1035_v58
			_1035_v58 = _1034_v57
			var _rhs198 _dafny.Array = _1032_v55
			_ = _rhs198
			var _rhs199 _dafny.MultiSet = _1035_v58
			_ = _rhs199
			var _rhs200 _dafny.Int = (_this).F26()
			_ = _rhs200
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			_1032_v55 = _rhs198
			_1035_v58 = _rhs199
			_lhs176.F2 = _rhs200
			var _1036_v59 _dafny.Array
			_ = _1036_v59
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_27
			var _nw173 _dafny.Array
			_ = _nw173
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw173 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Sequence = func(_1037_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F26())
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw173 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw173).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw173).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1036_v59 = _nw173
			var _1038_v60 _dafny.Sequence
			_ = _1038_v60
			_1038_v60 = _dafny.SeqOf(_1026_i4)
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_1036_v59), 0))
			_ = _index189
			(_1036_v59).ArraySet1(_1038_v60, (_index189).Int())
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_1036_v59), 0))
			_ = _index190
			(_1036_v59).ArraySet1(_1038_v60, (_index190).Int())
		}
		var _1039_v61 D1
		_ = _1039_v61
		_1039_v61 = Companion_D1_.Create_DC3_(_957_v0)
		var _1040_v62 _dafny.Sequence
		_ = _1040_v62
		_1040_v62 = _dafny.SeqOf(_1039_v61)
		var _1041_v63 _dafny.Map
		_ = _1041_v63
		_1041_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _1040_v62)
		r0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1041_v63).Contains((_this).F26()) {
				return (_1041_v63).Get((_this).F26()).(_dafny.Sequence)
			}
			return _1040_v62
		})(), _1040_v62), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(304))).Uint32(), func(coer51 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}((func(_1042_v61 D1) func(_dafny.Int) D1 {
			return func(_1043_i7 _dafny.Int) D1 {
				return _1042_v61
			}
		})(_1039_v61))))
		return r0
	}
}
func (_this *C3) M4(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _source15 D4 = Companion_Default___.Fm32(globalState)
		_ = _source15
		if _source15.Is_DC10() {
			var _1044___mcc_h0 _dafny.CodePoint = _source15.Get_().(D4_DC10).Cf16
			_ = _1044___mcc_h0
			var _1045___mcc_h1 _dafny.CodePoint = _source15.Get_().(D4_DC10).Cf17
			_ = _1045___mcc_h1
			var _1046_cf17 _dafny.CodePoint = _1045___mcc_h1
			_ = _1046_cf17
			var _1047_cf16 _dafny.CodePoint = _1044___mcc_h0
			_ = _1047_cf16
			var _1048_v0 _dafny.MultiSet
			_ = _1048_v0
			_1048_v0 = _dafny.MultiSetOf((_this).F26(), (_this).F26())
			var _1049_v1 _dafny.MultiSet
			_ = _1049_v1
			_1049_v1 = _1048_v0
			var _1050_v2 _dafny.Map
			_ = _1050_v2
			_1050_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _1051_v3 _dafny.Map
			_ = _1051_v3
			_1051_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), !(p1))
			var _1052_v4 _dafny.Array
			_ = _1052_v4
			var _nwElement0_36 bool = true
			_ = _nwElement0_36
			var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(21))
			_ = _nw174
			(_nw174).ArraySet1(_nwElement0_36, 0)
			(_nw174).ArraySet1((func() bool {
				if (_1050_v2).Contains(p1) {
					return (_1050_v2).Get(p1).(bool)
				}
				return false
			})(), 1)
			(_nw174).ArraySet1((_this).F25(), 2)
			(_nw174).ArraySet1(p1, 3)
			(_nw174).ArraySet1(p2, 4)
			(_nw174).ArraySet1(p1, 5)
			(_nw174).ArraySet1((_this).F25(), 6)
			(_nw174).ArraySet1(p1, 7)
			(_nw174).ArraySet1((func() bool {
				if (_1051_v3).Contains(_dafny.IntOfInt64(-745)) {
					return (_1051_v3).Get(_dafny.IntOfInt64(-745)).(bool)
				}
				return p2
			})(), 8)
			(_nw174).ArraySet1(p2, 9)
			(_nw174).ArraySet1(!(p1), 10)
			(_nw174).ArraySet1(p2, 11)
			(_nw174).ArraySet1((_this).F25(), 12)
			(_nw174).ArraySet1(true, 13)
			(_nw174).ArraySet1((_this).F25(), 14)
			(_nw174).ArraySet1(p1, 15)
			(_nw174).ArraySet1((_this).F30(), 16)
			(_nw174).ArraySet1(p2, 17)
			(_nw174).ArraySet1(p2, 18)
			(_nw174).ArraySet1(p2, 19)
			(_nw174).ArraySet1(p1, 20)
			_1052_v4 = _nw174
			var _1053_v5 T0
			_ = _1053_v5
			var _nw175 *C0 = New_C0_()
			_ = _nw175
			_nw175.Ctor__((_this).F29(), true, p0, _1052_v4, (_this).F30())
			_1053_v5 = _nw175
			var _1054_v6 _dafny.Set
			_ = _1054_v6
			_1054_v6 = _dafny.SetOf((_this).F25())
			var _1055_v7 _dafny.Map
			_ = _1055_v7
			_1055_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1053_v5, (_1054_v6).Cardinality())
			var _1056_v8 _dafny.Sequence
			_ = _1056_v8
			_1056_v8 = _dafny.SeqOf((_this).F26(), (_this).F26(), (func() _dafny.Int {
				if (_1055_v7).Contains(_1053_v5) {
					return (_1055_v7).Get(_1053_v5).(_dafny.Int)
				}
				return (_this).F26()
			})())
			var _1057_v9 _dafny.MultiSet
			_ = _1057_v9
			_1057_v9 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_1056_v8))
			var _1058_v10 _dafny.Array
			_ = _1058_v10
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_28
			var _nw176 _dafny.Array
			_ = _nw176
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw176 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Map = (func(_1059_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1060_i0 _dafny.Int) _dafny.Map {
						return _1059_v2
					}
				})(_1050_v2)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw176 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw176).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw176).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1058_v10 = _nw176
			var _1061_v11 _dafny.Map
			_ = _1061_v11
			_1061_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1053_v5).F26(), Companion_D8_.Create_DC16_(_1058_v10))
			var _1062_v12 _dafny.MultiSet
			_ = _1062_v12
			_1062_v12 = _dafny.MultiSetOf(_1061_v11, _1061_v11)
			var _1063_v13 _dafny.MultiSet
			_ = _1063_v13
			_1063_v13 = _dafny.MultiSetOf(Companion_Default___.Fm1(globalState))
			var _1064_v14 _dafny.Array
			_ = _1064_v14
			var _nwElement0_37 _dafny.Int = (_this).F26()
			_ = _nwElement0_37
			var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(21))
			_ = _nw177
			(_nw177).ArraySet1(_nwElement0_37, 0)
			(_nw177).ArraySet1((_1057_v9).Cardinality(), 1)
			(_nw177).ArraySet1((_1053_v5).F26(), 2)
			(_nw177).ArraySet1((_this).F26(), 3)
			(_nw177).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1062_v12).Contains(_1061_v11) {
					return (_1062_v12).Multiplicity(_1061_v11)
				}
				return (_1053_v5).F26()
			})())).Plus((_this).F26()), 4)
			(_nw177).ArraySet1((func() _dafny.Int {
				if (_1063_v13).Contains(_dafny.UnicodeSeqOfUtf8Bytes("gesulblvl")) {
					return (_1063_v13).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("gesulblvl"))
				}
				return (_dafny.Zero).Minus((_1053_v5).F26())
			})(), 5)
			(_nw177).ArraySet1((_this).F26(), 6)
			(_nw177).ArraySet1((func() _dafny.Int {
				if (_this).F30() {
					return (_1051_v3).Cardinality()
				}
				return p0
			})(), 7)
			(_nw177).ArraySet1((_1053_v5).F26(), 8)
			(_nw177).ArraySet1(p0, 9)
			(_nw177).ArraySet1((_1056_v8).Select((Companion_Default___.SafeIndex((_1053_v5).F26(), _dafny.IntOfUint32((_1056_v8).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw177).ArraySet1((_1053_v5).F26(), 11)
			(_nw177).ArraySet1(_dafny.IntOfInt64(-274), 12)
			(_nw177).ArraySet1(_dafny.IntOfUint32((_1056_v8).Cardinality()), 13)
			(_nw177).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F29(), (_this).F29())).Cardinality()), 14)
			(_nw177).ArraySet1((_this).F26(), 15)
			(_nw177).ArraySet1((_dafny.IntOfInt64(312)).Minus((_1053_v5).F26()), 16)
			(_nw177).ArraySet1(p0, 17)
			(_nw177).ArraySet1((_this).Fm6((_this).F25(), (_this).F25(), _dafny.IntOfInt64(132), globalState), 18)
			(_nw177).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p1)).Cardinality()).Minus((_1053_v5).F26()), 19)
			(_nw177).ArraySet1((_this).F26(), 20)
			_1064_v14 = _nw177
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index191
			(_1064_v14).ArraySet1(p0, (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index192
			(_1064_v14).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F29(), _dafny.UnicodeSeqOfUtf8Bytes("ix"))).Cardinality()), (_index192).Int())
			var _1065_v15 _dafny.Set
			_ = _1065_v15
			_1065_v15 = _dafny.SetOf((func() _dafny.Int {
				if p2 {
					return (_1064_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))).Int()).(_dafny.Int)
				}
				return (_1053_v5).F26()
			})())
			var _1066_v16 _dafny.Map
			_ = _1066_v16
			_1066_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1053_v5).F26(), Companion_Default___.Fm1(globalState))
			var _1067_v17 _dafny.Sequence
			_ = _1067_v17
			_1067_v17 = _dafny.SeqOf(p1, p2, p1, (((_1066_v16).Update((_1064_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))).Int()).(_dafny.Int), (_this).F29())).Cardinality()).Cmp((_1053_v5).F26()) < 0)
			var _1068_v18 _dafny.Array
			_ = _1068_v18
			var _nw178 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(9))
			_ = _nw178
			_1068_v18 = _nw178
			var _1069_v19 D6
			_ = _1069_v19
			_1069_v19 = Companion_D6_.Create_DC13_(Companion_Default___.Fm2(Companion_Default___.Fm2((_this).F26(), _dafny.IntOfInt64(102), globalState), (_1064_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))).Int()).(_dafny.Int), globalState), (_1053_v5).F26(), _1065_v15)
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1068_v18), 0))
			_ = _index193
			(_1068_v18).ArraySet1(_1069_v19, (_index193).Int())
			var _pat_let_tv31 = _1064_v14
			_ = _pat_let_tv31
			var _pat_let_tv32 = _1064_v14
			_ = _pat_let_tv32
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1068_v18), 0))
			_ = _index194
			var _rhs201 _dafny.Set = _1065_v15
			_ = _rhs201
			var _rhs202 _dafny.Map = _1050_v2
			_ = _rhs202
			var _rhs203 _dafny.Sequence = _dafny.SeqOf((_1053_v5).F25())
			_ = _rhs203
			var _rhs204 D6 = func(_pat_let36_0 D6) D6 {
				return func(_1070_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let37_0 _dafny.Int) D6 {
						return func(_1071_dt__update_hcf21_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC13_((_1070_dt__update__tmp_h1).Dtor_cf20(), _1071_dt__update_hcf21_h0, (_1070_dt__update__tmp_h1).Dtor_cf22())
						}(_pat_let37_0)
					}((_pat_let_tv32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_pat_let_tv31), 0))).Int()).(_dafny.Int))
				}(_pat_let36_0)
			}(_1069_v19)
			_ = _rhs204
			var _rhs205 _dafny.CodePoint = _1047_cf16
			_ = _rhs205
			var _lhs177 _dafny.Array = _1068_v18
			_ = _lhs177
			var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_1068_v18), 0))
			_ = _lhs178
			_1065_v15 = _rhs201
			_1050_v2 = _rhs202
			_1067_v17 = _rhs203
			(_lhs177).ArraySet1(_rhs204, (_lhs178).Int())
			_1047_cf16 = _rhs205
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index195
			(_1064_v14).ArraySet1((_1064_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1064_v14), 0))).Int()).(_dafny.Int), (_index195).Int())
			(globalState).F1 = p2
		} else {
			var _1072___mcc_h2 _dafny.Set = _source15.Get_().(D4_DC9).Cf15
			_ = _1072___mcc_h2
			var _1073_cf15 _dafny.Set = _1072___mcc_h2
			_ = _1073_cf15
			var _1074_v20 _dafny.Map
			_ = _1074_v20
			_1074_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), p2)
			r1 = (func() _dafny.Int {
				if !((func() bool {
					if (_1074_v20).Contains((_this).F26()) {
						return (_1074_v20).Get((_this).F26()).(bool)
					}
					return (_this).F30()
				})()) {
					return (_this).F26()
				}
				return (_this).F26()
			})()
			var _1075_v21 _dafny.Array
			_ = _1075_v21
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_29
			var _nw179 _dafny.Array
			_ = _nw179
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw179 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) bool = (func(_1076_p1 bool) func(_dafny.Int) bool {
					return func(_1077_i1 _dafny.Int) bool {
						return _1076_p1
					}
				})(p1)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw179 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw179).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw179).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1075_v21 = _nw179
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1075_v21), 0))
			_ = _index196
			(_1075_v21).ArraySet1(_dafny.Companion_Sequence_.Equal((_this).F29(), (_this).F29()), (_index196).Int())
			var _1078_v22 _dafny.Map
			_ = _1078_v22
			_1078_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1074_v20).Cardinality(), p2), p1)
			var _1079_v23 _dafny.CodePoint
			_ = _1079_v23
			_1079_v23 = _dafny.CodePoint('d')
			var _1080_v24 _dafny.Map
			_ = _1080_v24
			_1080_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1079_v23, (_this).F26())
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1075_v21), 0))
			_ = _index197
			(_1075_v21).ArraySet1(!(_dafny.MultiSetOf((_1080_v24).Cardinality())).Contains(((_1078_v22).Cardinality()).Plus(p0)), (_index197).Int())
			var _1081_v25 *C1
			_ = _1081_v25
			var _nw180 *C1 = New_C1_()
			_ = _nw180
			_nw180.Ctor__(p1)
			_1081_v25 = _nw180
			var _1082_v26 *C2
			_ = _1082_v26
			var _nw181 *C2 = New_C2_()
			_ = _nw181
			_nw181.Ctor__(p0, (_1075_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1075_v21), 0))).Int()).(bool), _dafny.IntOfInt64(-12))
			_1082_v26 = _nw181
		}
		var _hi10 _dafny.Int = p0
		_ = _hi10
		for _1083_i2 := (func() _dafny.Int {
			if p2 {
				return _dafny.IntOfInt64(606)
			}
			return (_this).F26()
		})(); _1083_i2.Cmp(_hi10) < 0; _1083_i2 = _1083_i2.Plus(_dafny.One) {
			var _1084_v27 _dafny.Array
			_ = _1084_v27
			var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw182
			_1084_v27 = _nw182
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))
			_ = _index198
			(_1084_v27).ArraySet1(_1083_i2, (_index198).Int())
			var _1085_v28 _dafny.MultiSet
			_ = _1085_v28
			_1085_v28 = _dafny.MultiSetOf(_dafny.IntOfUint32((Companion_Default___.Fm3(_1083_i2, p2, _dafny.IntOfInt64(-651), p1, globalState)).Cardinality()))
			var _1086_v29 _dafny.MultiSet
			_ = _1086_v29
			_1086_v29 = _dafny.MultiSetOf(_1085_v28)
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))
			_ = _index199
			(_1084_v27).ArraySet1((func() _dafny.Int {
				if (_1086_v29).Contains(Companion_Default___.Fm21((_this).F29(), (_this).F26(), p0, globalState)) {
					return (_1086_v29).Multiplicity(Companion_Default___.Fm21((_this).F29(), (_this).F26(), p0, globalState))
				}
				return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(833))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1087_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1088_i3 _dafny.Int) _dafny.Int {
						return _1087_i2
					}
				})(_1083_i2)))).Cardinality()), _dafny.IntOfUint32(((_this).F29()).Cardinality()))
			})(), (_index199).Int())
			var _1089_v30 _dafny.MultiSet
			_ = _1089_v30
			_1089_v30 = _dafny.MultiSetOf((_this).F30())
			var _1090_v31 *C2
			_ = _1090_v31
			var _nw183 *C2 = New_C2_()
			_ = _nw183
			_nw183.Ctor__((_1089_v30).Cardinality(), false, (_dafny.Zero).Minus((_this).F26()))
			_1090_v31 = _nw183
			var _1091_v32 _dafny.MultiSet
			_ = _1091_v32
			_1091_v32 = _dafny.MultiSetOf(_1090_v31)
			if (_dafny.MultiSetOf(_1090_v31, _1090_v31)).IsProperSubsetOf(_1091_v32) {
				r1 = p0
				var _1092_v33 _dafny.Array
				_ = _1092_v33
				var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw184
				_1092_v33 = _nw184
				var _1093_v34 _dafny.Sequence
				_ = _1093_v34
				_1093_v34 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1084_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))).Int()).(_dafny.Int)))
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1092_v33), 0))
				_ = _index200
				(_1092_v33).ArraySet1(_1093_v34, (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1092_v33), 0))
				_ = _index201
				(_1092_v33).ArraySet1(_1093_v34, (_index201).Int())
				var _1094_v35 D0
				_ = _1094_v35
				_1094_v35 = Companion_D0_.Create_DC0_((_this).F30())
				(globalState).F1 = Companion_Default___.Fm5((_this).F30(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1094_v35, (_this).F26()), (_this).F26(), globalState)
				var _1095_v36 _dafny.Set
				_ = _1095_v36
				_1095_v36 = _dafny.SetOf((_1084_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))).Int()).(_dafny.Int))
				var _1096_v37 _dafny.Sequence
				_ = _1096_v37
				_1096_v37 = _dafny.SeqOf((_1090_v31).F31())
				r0 = ((_1095_v36).Cardinality()).Minus(_dafny.IntOfUint32((_1096_v37).Cardinality()))
				var _1097_v38 D0
				_ = _1097_v38
				_1097_v38 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_1083_i2), (_this).F29(), p1, _dafny.IntOfInt64(145), p1)
				(globalState).F0 = (_1097_v38).Dtor_cf1()
			} else {
				var _1098_v39 D1
				_ = _1098_v39
				_1098_v39 = Companion_D1_.Create_DC4_((_this).F30(), (_this).F25())
				var _1099_v40 _dafny.Sequence
				_ = _1099_v40
				_1099_v40 = _dafny.SeqOf(p1, (_1098_v39).Dtor_cf7(), (_this).F30())
				(globalState).F2 = (_dafny.IntOfUint32((_1099_v40).Cardinality())).Times((func() _dafny.Int {
					if p2 {
						return (_1084_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))).Int()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(720)
				})())
				var _1100_v41 _dafny.Sequence
				_ = _1100_v41
				_1100_v41 = _dafny.UnicodeSeqOfUtf8Bytes("pq")
				var _1101_v42 _dafny.MultiSet
				_ = _1101_v42
				_1101_v42 = _dafny.MultiSetOf(_1100_v41, (_this).F29(), (_this).F29())
				var _1102_v43 _dafny.Set
				_ = _1102_v43
				_1102_v43 = _dafny.SetOf((_1084_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))).Int()).(_dafny.Int))
				var _1103_v44 _dafny.Array
				_ = _1103_v44
				var _nwElement0_38 bool = ((_1090_v31).F31()).Cmp(_dafny.IntOfInt64(-210)) <= 0
				_ = _nwElement0_38
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(23))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_38, 0)
				(_nw185).ArraySet1((_1101_v42).IsProperSubsetOf(_dafny.MultiSetOf((_this).F29(), _1100_v41, _1100_v41)), 1)
				(_nw185).ArraySet1(p2, 2)
				(_nw185).ArraySet1((_this).F30(), 3)
				(_nw185).ArraySet1(!((p1) && ((_this).F25())), 4)
				(_nw185).ArraySet1((func() bool {
					if true {
						return (_this).F25()
					}
					return p2
				})(), 5)
				(_nw185).ArraySet1((_this).F30(), 6)
				(_nw185).ArraySet1((_1085_v28).IsSubsetOf(_1085_v28), 7)
				(_nw185).ArraySet1((_this).F30(), 8)
				(_nw185).ArraySet1((_this).F25(), 9)
				(_nw185).ArraySet1(!((_this).F30()) || ((_this).F25()), 10)
				(_nw185).ArraySet1(p2, 11)
				(_nw185).ArraySet1((_this).F25(), 12)
				(_nw185).ArraySet1(p1, 13)
				(_nw185).ArraySet1((_this).F25(), 14)
				(_nw185).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("nwtpj"), _1100_v41), 15)
				(_nw185).ArraySet1(!((_this).F30()), 16)
				(_nw185).ArraySet1((_this).F25(), 17)
				(_nw185).ArraySet1(p1, 18)
				(_nw185).ArraySet1((_dafny.SetOf((_1090_v31).F31())).IsProperSubsetOf(_1102_v43), 19)
				(_nw185).ArraySet1((_this).F25(), 20)
				(_nw185).ArraySet1(!((_this).F30()), 21)
				(_nw185).ArraySet1(p1, 22)
				_1103_v44 = _nw185
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1103_v44), 0))
				_ = _index202
				(_1103_v44).ArraySet1(((_1089_v30).Cardinality()).Cmp(_1083_i2) < 0, (_index202).Int())
				var _1104_v45 _dafny.Array
				_ = _1104_v45
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_30
				var _nw186 _dafny.Array
				_ = _nw186
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw186 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.MultiSet = (func(_1105_v28 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_1106_i4 _dafny.Int) _dafny.MultiSet {
							return _1105_v28
						}
					})(_1085_v28)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw186 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw186).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw186).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1104_v45 = _nw186
				var _1107_v46 D10
				_ = _1107_v46
				_1107_v46 = Companion_D10_.Create_DC21_(_1104_v45)
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1103_v44), 0))
				_ = _index203
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))
				_ = _index204
				var _rhs206 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F29(), (_this).F29())
				_ = _rhs206
				var _rhs207 bool = true
				_ = _rhs207
				var _rhs208 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-169), ((_dafny.Zero).Minus((_1084_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))).Int()).(_dafny.Int))).Times(_1083_i2))
				_ = _rhs208
				var _rhs209 _dafny.Array = (_1107_v46).Dtor_cf33()
				_ = _rhs209
				var _rhs210 bool = (_this).F25()
				_ = _rhs210
				var _lhs179 _dafny.Array = _1103_v44
				_ = _lhs179
				var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1103_v44), 0))
				_ = _lhs180
				var _lhs181 _dafny.Array = _1084_v27
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1084_v27), 0))
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				_1100_v41 = _rhs206
				(_lhs179).ArraySet1(_rhs207, (_lhs180).Int())
				(_lhs181).ArraySet1(_rhs208, (_lhs182).Int())
				_1104_v45 = _rhs209
				_lhs183.F1 = _rhs210
				var _1108_v47 _dafny.Sequence
				_ = _1108_v47
				_1108_v47 = _dafny.SeqOf(_1084_v27)
				_1108_v47 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1108_v47, (Companion_Default___.SafeIndex((_1090_v31).F31(), _dafny.IntOfUint32((_1108_v47).Cardinality()))).Uint32(), _1084_v27), _dafny.Companion_Sequence_.Concatenate(_1108_v47, _1108_v47))
				(globalState).F1 = !((_this).F25())
				var _1109_v48 D0
				_ = _1109_v48
				_1109_v48 = Companion_D0_.Create_DC0_((_1103_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1103_v44), 0))).Int()).(bool))
				(globalState).F1 = (_1109_v48).Dtor_cf0()
			}
			var _1110_v49 _dafny.Array
			_ = _1110_v49
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_31
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) bool = func(_1111_i5 _dafny.Int) bool {
					return (_this).F25()
				}
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw187 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw187).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw187).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1110_v49 = _nw187
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1110_v49), 0))
			_ = _index205
			(_1110_v49).ArraySet1(p2, (_index205).Int())
			var _1112_v50 _dafny.Array
			_ = _1112_v50
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_32
			var _nw188 _dafny.Array
			_ = _nw188
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw188 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1113_p2 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1114_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1113_p2, (_this).F30()), _dafny.CodePoint('p'))).Cardinality(), (_this).F26())
					}
				})(p2)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw188 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw188).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw188).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1112_v50 = _nw188
			var _1115_v51 _dafny.Sequence
			_ = _1115_v51
			_1115_v51 = _dafny.SeqOf(p1)
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1110_v49), 0))
			_ = _index206
			var _rhs211 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1115_v51, _dafny.SeqOf((_this).F30(), (_this).F30()))
			_ = _rhs211
			var _rhs212 bool = (_1090_v31).Fm7(globalState)
			_ = _rhs212
			var _rhs213 _dafny.Array = _1112_v50
			_ = _rhs213
			var _lhs184 _dafny.Array = _1110_v49
			_ = _lhs184
			var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_1110_v49), 0))
			_ = _lhs185
			r2 = _rhs211
			(_lhs184).ArraySet1(_rhs212, (_lhs185).Int())
			_1112_v50 = _rhs213
		}
		if (_this).F25() {
			var _1116_v52 _dafny.Map
			_ = _1116_v52
			_1116_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm7(globalState))
			var _1117_v53 _dafny.MultiSet
			_ = _1117_v53
			_1117_v53 = _dafny.MultiSetOf((_this).F26())
			var _1118_v54 _dafny.Sequence
			_ = _1118_v54
			_1118_v54 = _dafny.SeqOf((func() bool {
				if (_1116_v52).Contains((_1117_v53).Cardinality()) {
					return (_1116_v52).Get((_1117_v53).Cardinality()).(bool)
				}
				return (_this).F30()
			})())
			var _1119_v55 _dafny.Map
			_ = _1119_v55
			_1119_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v54, (_this).F30())
			var _1120_v56 _dafny.Map
			_ = _1120_v56
			_1120_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1119_v55)
			var _1121_v57 _dafny.MultiSet
			_ = _1121_v57
			_1121_v57 = _dafny.MultiSetOf(true, (_this).F25(), p1)
			var _1122_v58 D6
			_ = _1122_v58
			_1122_v58 = Companion_D6_.Create_DC14_((_this).F26(), _dafny.IntOfInt64(-684), _1121_v57)
			r0 = ((_1120_v56).Merge((Companion_Default___.Fm33(_1122_v58, p0, globalState)).Merge(_1120_v56))).Cardinality()
			(globalState).F1 = ((_this).F26()).Cmp((_this).F26()) < 0
			var _1123_v59 _dafny.Array
			_ = _1123_v59
			var _len0_33 _dafny.Int = _dafny.One
			_ = _len0_33
			var _nw189 _dafny.Array
			_ = _nw189
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw189 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Int = (func(_1124_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1125_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1125_i7, _1124_p0)
					}
				})(p0)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw189 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw189).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw189).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1123_v59 = _nw189
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1123_v59), 0))
			_ = _index207
			(_1123_v59).ArraySet1((_this).F26(), (_index207).Int())
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1123_v59), 0))
			_ = _index208
			(_1123_v59).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(221), p0), (_this).F26()), (_index208).Int())
			r0 = (_this).F26()
			var _1126_v60 _dafny.Array
			_ = _1126_v60
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_34
			var _nw190 _dafny.Array
			_ = _nw190
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw190 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Map = (func(_1127_p2 bool) func(_dafny.Int) _dafny.Map {
					return func(_1128_i8 _dafny.Int) _dafny.Map {
						return (Companion_D11_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1127_p2, _1127_p2))).Dtor_cf39()
					}
				})(p2)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw190 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw190).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw190).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1126_v60 = _nw190
			var _1129_v61 D8
			_ = _1129_v61
			_1129_v61 = Companion_D8_.Create_DC16_(_1126_v60)
			var _1130_v62 _dafny.Map
			_ = _1130_v62
			_1130_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v54, _1129_v61)
			var _1131_v63 *C2
			_ = _1131_v63
			var _nw191 *C2 = New_C2_()
			_ = _nw191
			_nw191.Ctor__((_1130_v62).Cardinality(), p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1118_v54).Cardinality())))
			_1131_v63 = _nw191
		} else {
			var _1132_v64 _dafny.Map
			_ = _1132_v64
			_1132_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfInt64(545))
			_1132_v64 = _1132_v64
			var _1133_v65 _dafny.Array
			_ = _1133_v65
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_35
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = func(_1134_i9 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1134_i9, _dafny.IntOfInt64(843))
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw192 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw192).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw192).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1133_v65 = _nw192
			var _1135_v66 _dafny.Array
			_ = _1135_v66
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_36
			var _nw193 _dafny.Array
			_ = _nw193
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw193 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Sequence = func(_1136_i10 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F29(), (_this).F29(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(128))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}(func(_1137_i11 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					})))
				}
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw193 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw193).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw193).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1135_v66 = _nw193
			var _1138_v67 _dafny.Sequence
			_ = _1138_v67
			_1138_v67 = _dafny.SeqOf((_this).F29())
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1135_v66), 0))
			_ = _index209
			(_1135_v66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1138_v67, _1138_v67), (_index209).Int())
			var _1139_v68 _dafny.CodePoint
			_ = _1139_v68
			_1139_v68 = _dafny.CodePoint('c')
			var _1140_v69 _dafny.MultiSet
			_ = _1140_v69
			_1140_v69 = _dafny.MultiSetOf(_dafny.IntOfInt64(397), (_this).F26(), (_this).F26())
			var _1141_v70 _dafny.Map
			_ = _1141_v70
			_1141_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm28((_this).F29(), _1139_v68, (_this).F26(), globalState), _dafny.IntOfUint32(((_this).F29()).Cardinality()))
			var _1142_v73 D0
			_ = _1142_v73
			_1142_v73 = Companion_D0_.Create_DC0_((_this).F30())
			var _1143_v74 _dafny.Sequence
			_ = _1143_v74
			_1143_v74 = _dafny.SeqOf(_1142_v73, _1142_v73)
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1135_v66), 0))
			_ = _index210
			var _rhs214 _dafny.Array = _1133_v65
			_ = _rhs214
			var _rhs215 _dafny.Sequence = (func() _dafny.Sequence {
				if (_this).Fm7(globalState) {
					return _1138_v67
				}
				return _1138_v67
			})()
			_ = _rhs215
			var _rhs216 bool = (_this).F25()
			_ = _rhs216
			var _rhs217 bool = Companion_Default___.Fm5((_1140_v69).IsDisjointFrom(Companion_Default___.Fm21(_dafny.UnicodeSeqOfUtf8Bytes("ncdfeu"), (_this).F26(), (_this).F26(), globalState)), (func() _dafny.Map {
				if false {
					return _1141_v70
				}
				return func() _dafny.Map {
					var _coll22 = _dafny.NewMapBuilder()
					_ = _coll22
					for _iter22 := _dafny.Iterate((func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter23 := _dafny.Iterate((_1143_v74).Elements()); ; {
							_compr_23, _ok23 := _iter23()
							if !_ok23 {
								break
							}
							var _1144_v72 D0
							_1144_v72 = interface{}(_compr_23).(D0)
							if _dafny.Companion_Sequence_.Contains(_1143_v74, _1144_v72) {
								_coll23.Add(_1144_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F30())).Cardinality())
							}
						}
						return _coll23.ToMap()
					}()).Keys().Elements()); ; {
						_compr_22, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _1145_v71 D0
						_1145_v71 = interface{}(_compr_22).(D0)
						if (func() _dafny.Map {
							var _coll24 = _dafny.NewMapBuilder()
							_ = _coll24
							for _iter24 := _dafny.Iterate((_1143_v74).Elements()); ; {
								_compr_24, _ok24 := _iter24()
								if !_ok24 {
									break
								}
								var _1144_v72 D0
								_1144_v72 = interface{}(_compr_24).(D0)
								if _dafny.Companion_Sequence_.Contains(_1143_v74, _1144_v72) {
									_coll24.Add(_1144_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F30())).Cardinality())
								}
							}
							return _coll24.ToMap()
						}()).Contains(_1145_v71) {
							_coll22.Add(_1145_v71, _dafny.IntOfUint32(((_this).F29()).Cardinality()))
						}
					}
					return _coll22.ToMap()
				}()
			})(), ((_this).F26()).Times((func() _dafny.Int {
				if (_1132_v64).Contains((_this).F26()) {
					return (_1132_v64).Get((_this).F26()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_this).F26())
			})()), globalState)
			_ = _rhs217
			var _rhs218 _dafny.CodePoint = _1139_v68
			_ = _rhs218
			var _lhs186 _dafny.Array = _1135_v66
			_ = _lhs186
			var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1135_v66), 0))
			_ = _lhs187
			_1133_v65 = _rhs214
			(_lhs186).ArraySet1(_rhs215, (_lhs187).Int())
			r2 = _rhs216
			r2 = _rhs217
			_1139_v68 = _rhs218
			var _1146_v75 _dafny.Array
			_ = _1146_v75
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
			_ = _nw194
			_1146_v75 = _nw194
			_1146_v75 = _1146_v75
			var _1147_v76 _dafny.Array
			_ = _1147_v76
			var _nw195 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw195
			_1147_v76 = _nw195
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1147_v76), 0))
			_ = _index211
			(_1147_v76).ArraySet1((_this).F25(), (_index211).Int())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1147_v76), 0))
			_ = _index212
			(_1147_v76).ArraySet1((Companion_Default___.Fm2(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1138_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.IntOfUint32((_1138_v67).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))), globalState)).Cmp((_this).F26()) >= 0, (_index212).Int())
			var _1148_v77 _dafny.Set
			_ = _1148_v77
			_1148_v77 = _dafny.SetOf((_dafny.SetOf(false, (_this).F25(), p1)).Cardinality(), p0)
			var _1149_v78 _dafny.Sequence
			_ = _1149_v78
			_1149_v78 = _dafny.SeqOf((_dafny.SetOf((_this).F26())).IsDisjointFrom(_1148_v77), true)
			var _1150_v79 *C0
			_ = _1150_v79
			var _nw196 *C0 = New_C0_()
			_ = _nw196
			_nw196.Ctor__((_this).F29(), (_1147_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1147_v76), 0))).Int()).(bool), (_this).F26(), _1147_v76, false)
			_1150_v79 = _nw196
			var _1151_v80 _dafny.Map
			_ = _1151_v80
			_1151_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v79, _1149_v78)
			_1149_v78 = (func() _dafny.Sequence {
				if (_1151_v80).Contains(_1150_v79) {
					return (_1151_v80).Get(_1150_v79).(_dafny.Sequence)
				}
				return Companion_Default___.Fm23(globalState)
			})()
		}
		var _1152_v81 D11
		_ = _1152_v81
		_1152_v81 = Companion_D11_.Create_DC26_((_this).F26(), (_this).F26(), (_this).F25())
		var _1153_v82 _dafny.Map
		_ = _1153_v82
		_1153_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1152_v81).Dtor_cf41(), (_this).Fm6(!(p2), (_1152_v81).Dtor_cf42(), (_dafny.Zero).Minus(p0), globalState))
		var _1154_v83 _dafny.Sequence
		_ = _1154_v83
		_1154_v83 = _dafny.SeqOf(p0, (_this).F26())
		_1153_v82 = (_1153_v82).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-976))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg54 _dafny.Int) interface{} {
				return coer54(arg54)
			}
		}(func(_1155_i12 _dafny.Int) _dafny.Int {
			return (_this).F26()
		})), _1154_v83)).Cardinality()), p0)
		var _1156_v84 _dafny.Array
		_ = _1156_v84
		var _len0_37 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_37
		var _nw197 _dafny.Array
		_ = _nw197
		if _len0_37.Cmp(_dafny.Zero) == 0 {
			_nw197 = _dafny.NewArray(_len0_37)
		} else {
			var _init37 func(_dafny.Int) bool = func(_1157_i13 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(638))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}(func(_1158_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), (_this).F29())
			}
			_ = _init37
			var _element0_37 = _init37(_dafny.Zero)
			_ = _element0_37
			_nw197 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
			(_nw197).ArraySet1(_element0_37, 0)
			var _nativeLen0_37 = (_len0_37).Int()
			_ = _nativeLen0_37
			for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
				(_nw197).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
			}
		}
		_1156_v84 = _nw197
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1156_v84), 0))
		_ = _index213
		(_1156_v84).ArraySet1(p1, (_index213).Int())
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1156_v84), 0))
		_ = _index214
		(_1156_v84).ArraySet1(!((_this).F30()), (_index214).Int())
		var _hi11 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F26(), (_dafny.Zero).Minus((_this).F26()))
		_ = _hi11
		for _1159_i15 := p0; _1159_i15.Cmp(_hi11) < 0; _1159_i15 = _1159_i15.Plus(_dafny.One) {
			var _1160_v85 _dafny.Array
			_ = _1160_v85
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_38
			var _nw198 _dafny.Array
			_ = _nw198
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw198 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Set = (func(_1161_v84 _dafny.Array) func(_dafny.Int) _dafny.Set {
					return func(_1162_i16 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(false)).Difference(_dafny.SetOf((_1161_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1161_v84), 0))).Int()).(bool)))
					}
				})(_1156_v84)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw198 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw198).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw198).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1160_v85 = _nw198
			var _1163_v86 _dafny.Map
			_ = _1163_v86
			_1163_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1154_v83).Cardinality()), (_this).F25())
			var _1164_v87 _dafny.Set
			_ = _1164_v87
			_1164_v87 = _dafny.SetOf((func() bool {
				if (_1163_v86).Contains(_1159_i15) {
					return (_1163_v86).Get(_1159_i15).(bool)
				}
				return (_this).F30()
			})())
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1160_v85), 0))
			_ = _index215
			(_1160_v85).ArraySet1((_1164_v87).Union(_dafny.SetOf((_this).F30())), (_index215).Int())
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1160_v85), 0))
			_ = _index216
			(_1160_v85).ArraySet1((_1164_v87).Intersection(_1164_v87), (_index216).Int())
			var _1165_v88 _dafny.Map
			_ = _1165_v88
			_1165_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1156_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1156_v84), 0))).Int()).(bool))
			r2 = (func() bool {
				if (_1165_v88).Contains((_this).F25()) {
					return (_1165_v88).Get((_this).F25()).(bool)
				}
				return (_1156_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1156_v84), 0))).Int()).(bool)
			})()
			var _1166_v89 _dafny.Sequence
			_ = _1166_v89
			_1166_v89 = _dafny.SeqOf((_1156_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_1156_v84), 0))).Int()).(bool))
			var _1167_v90 _dafny.Set
			_ = _1167_v90
			_1167_v90 = _dafny.SetOf(p0, _1159_i15, (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1166_v89, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1166_v89).Cardinality()))).Uint32(), true))).Cardinality())
			var _1168_v91 *C2
			_ = _1168_v91
			var _nw199 *C2 = New_C2_()
			_ = _nw199
			_nw199.Ctor__(_dafny.IntOfInt64(-927), (_this).F25(), (_1167_v90).Cardinality())
			_1168_v91 = _nw199
			var _1169_v92 _dafny.Sequence
			_ = _1169_v92
			_1169_v92 = _dafny.UnicodeSeqOfUtf8Bytes("wxlotbcg")
			var _rhs219 _dafny.Int = (_this).F26()
			_ = _rhs219
			var _rhs220 bool = p1
			_ = _rhs220
			var _rhs221 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("c")
			_ = _rhs221
			var _lhs188 *GlobalState = globalState
			_ = _lhs188
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			_lhs188.F0 = _rhs219
			_lhs189.F1 = _rhs220
			_1169_v92 = _rhs221
		}
		var _1170_v93 *C2
		_ = _1170_v93
		var _nw200 *C2 = New_C2_()
		_ = _nw200
		_nw200.Ctor__((_this).F26(), true, p0)
		_1170_v93 = _nw200
		var _1171_v94 _dafny.Map
		_ = _1171_v94
		_1171_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1170_v93)
		r0 = ((_1171_v94).Merge((_1171_v94).Merge(_1171_v94))).Cardinality()
		r1 = (_this).F26()
		r2 = (_this).F30()
		return r0, r1, r2
	}
}
func (_this *C3) F29() _dafny.Sequence {
	{
		return _this._f29
	}
}
func (_this *C3) F30() bool {
	{
		return _this._f30
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f25 bool
	_f26 _dafny.Int
	F28  _dafny.Sequence
	_f27 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f25 = false
	_this._f26 = _dafny.Zero
	_this.F28 = _dafny.EmptySeq
	_this._f27 = _dafny.Zero
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

func (_this *C4) F25() bool {
	return _this._f25
}
func (_this *C4) F26() _dafny.Int {
	return _this._f26
}
func (_this *C4) Ctor__(f27 _dafny.Int, f28 _dafny.Sequence, f25 bool, f26 _dafny.Int) {
	{
		(_this)._f27 = f27
		(_this).F28 = f28
		(_this)._f25 = f25
		(_this)._f26 = f26
	}
}
func (_this *C4) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F27()
	}
}
func (_this *C4) Fm7(globalState *GlobalState) bool {
	{
		return !(!(!((Companion_D0_.Create_DC0_(true)).Dtor_cf0())))
	}
}
func (_this *C4) Fm8(p0 bool, globalState *GlobalState) D0 {
	{
		var _source16 D0 = Companion_D0_.Create_DC2_()
		_ = _source16
		if _source16.Is_DC1() {
			var _1172___mcc_h0 _dafny.Int = _source16.Get_().(D0_DC1).Cf1
			_ = _1172___mcc_h0
			var _1173___mcc_h1 _dafny.Sequence = _source16.Get_().(D0_DC1).Cf2
			_ = _1173___mcc_h1
			var _1174___mcc_h2 bool = _source16.Get_().(D0_DC1).Cf3
			_ = _1174___mcc_h2
			var _1175___mcc_h3 _dafny.Int = _source16.Get_().(D0_DC1).Cf4
			_ = _1175___mcc_h3
			var _1176___mcc_h4 bool = _source16.Get_().(D0_DC1).Cf5
			_ = _1176___mcc_h4
			var _1177_cf5 bool = _1176___mcc_h4
			_ = _1177_cf5
			var _1178_cf4 _dafny.Int = _1175___mcc_h3
			_ = _1178_cf4
			var _1179_cf3 bool = _1174___mcc_h2
			_ = _1179_cf3
			var _1180_cf2 _dafny.Sequence = _1173___mcc_h1
			_ = _1180_cf2
			var _1181_cf1 _dafny.Int = _1172___mcc_h0
			_ = _1181_cf1
			return Companion_D0_.Create_DC2_()
		} else if _source16.Is_DC2() {
			return Companion_D0_.Create_DC2_()
		} else {
			var _1182___mcc_h5 bool = _source16.Get_().(D0_DC0).Cf0
			_ = _1182___mcc_h5
			var _1183_cf0 bool = _1182___mcc_h5
			_ = _1183_cf0
			return Companion_D0_.Create_DC2_()
		}
	}
}
func (_this *C4) M1(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F1 = (_this).F25()
		var _1184_v0 _dafny.Sequence
		_ = _1184_v0
		_1184_v0 = _dafny.SeqOf((_this).F25(), (_this).F25())
		if (_1184_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1184_v0).Cardinality()))).Uint32()).(bool) {
			var _1185_v1 _dafny.Set
			_ = _1185_v1
			_1185_v1 = _dafny.SetOf(_this.F28)
			var _1186_v2 D0
			_ = _1186_v2
			_1186_v2 = Companion_D0_.Create_DC0_(p2)
			var _1187_v3 _dafny.Map
			_ = _1187_v3
			_1187_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v2, (_this).F26())
			var _1188_v4 _dafny.Sequence
			_ = _1188_v4
			_1188_v4 = _dafny.SeqOf((_this).F26(), (_this).F27())
			var _1189_v5 _dafny.Array
			_ = _1189_v5
			var _nwElement0_39 bool = (true) || (p2)
			_ = _nwElement0_39
			var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(25))
			_ = _nw201
			(_nw201).ArraySet1(_nwElement0_39, 0)
			(_nw201).ArraySet1(p2, 1)
			(_nw201).ArraySet1((_this).F25(), 2)
			(_nw201).ArraySet1(p2, 3)
			(_nw201).ArraySet1((_this).F25(), 4)
			(_nw201).ArraySet1((_this).F25(), 5)
			(_nw201).ArraySet1(!(p1), 6)
			(_nw201).ArraySet1(!(p1), 7)
			(_nw201).ArraySet1(true, 8)
			(_nw201).ArraySet1(true, 9)
			(_nw201).ArraySet1(p1, 10)
			(_nw201).ArraySet1((_1185_v1).IsSubsetOf(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wrebsokg"), _this.F28)), 11)
			(_nw201).ArraySet1((_1184_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1184_v0).Cardinality()))).Uint32()).(bool), 12)
			(_nw201).ArraySet1(Companion_Default___.Fm5((_1186_v2).Dtor_cf0(), (_1187_v3).Update(Companion_D0_.Create_DC0_((_this).F25()), (_1188_v4).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1188_v4).Cardinality()))).Uint32()).(_dafny.Int)), (_this).F26(), globalState), 13)
			(_nw201).ArraySet1((_this).F25(), 14)
			(_nw201).ArraySet1((_this).F25(), 15)
			(_nw201).ArraySet1(false, 16)
			(_nw201).ArraySet1(!(p2), 17)
			(_nw201).ArraySet1((_this).F25(), 18)
			(_nw201).ArraySet1(p1, 19)
			(_nw201).ArraySet1(!(p2) || ((_this).F25()), 20)
			(_nw201).ArraySet1((_this).F25(), 21)
			(_nw201).ArraySet1((_this).F25(), 22)
			(_nw201).ArraySet1(p2, 23)
			(_nw201).ArraySet1(p1, 24)
			_1189_v5 = _nw201
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))
			_ = _index217
			(_1189_v5).ArraySet1((_dafny.IntOfUint32((_this.F28).Cardinality())).Cmp((_this).F27()) <= 0, (_index217).Int())
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))
			_ = _index218
			(_1189_v5).ArraySet1(p2, (_index218).Int())
			if (_this).F25() {
				var _1190_v6 _dafny.Array
				_ = _1190_v6
				var _nwElement0_40 _dafny.Sequence = _this.F28
				_ = _nwElement0_40
				var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.One)
				_ = _nw202
				(_nw202).ArraySet1(_nwElement0_40, 0)
				_1190_v6 = _nw202
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_1190_v6), 0))
				_ = _index219
				(_1190_v6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28), (_index219).Int())
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_1190_v6), 0))
				_ = _index220
				(_1190_v6).ArraySet1(Companion_Default___.Fm1(globalState), (_index220).Int())
				var _1191_v7 _dafny.Map
				_ = _1191_v7
				_1191_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool), (_this).F26())
				var _1192_v8 D0
				_ = _1192_v8
				_1192_v8 = Companion_D0_.Create_DC2_()
				var _1193_v9 _dafny.MultiSet
				_ = _1193_v9
				_1193_v9 = _dafny.MultiSetOf(_1192_v8, _1192_v8)
				var _1194_v10 _dafny.Sequence
				_ = _1194_v10
				_1194_v10 = _dafny.SeqOf(_1192_v8)
				var _1195_v11 D0
				_ = _1195_v11
				_1195_v11 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(535), (_1190_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_1190_v6), 0))).Int()).(_dafny.Sequence), p2, (_this).F26(), (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool))
				var _rhs222 bool = (_dafny.MultiSetOf(_1192_v8, _1192_v8, _1192_v8, Companion_D0_.Create_DC2_())).IsSubsetOf((_1193_v9).Intersection(_dafny.MultiSetFromSeq(_1194_v10)))
				_ = _rhs222
				var _rhs223 bool = (_this).F25()
				_ = _rhs223
				var _rhs224 bool = (_dafny.SetOf((_this).F26(), (_this).F27(), (_this).F27(), (_dafny.SetOf(false)).Cardinality())).Contains(Companion_Default___.SafeModuloInt((_this).F26(), _dafny.IntOfInt64(50)))
				_ = _rhs224
				var _rhs225 _dafny.Int = (_1195_v11).Dtor_cf4()
				_ = _rhs225
				var _rhs226 _dafny.Map = (_1191_v7).Merge(_1191_v7)
				_ = _rhs226
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				var _lhs193 *GlobalState = globalState
				_ = _lhs193
				_lhs190.F1 = _rhs222
				_lhs191.F1 = _rhs223
				_lhs192.F1 = _rhs224
				_lhs193.F2 = _rhs225
				_1191_v7 = _rhs226
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))
				_ = _index221
				(_1189_v5).ArraySet1((_this).F25(), (_index221).Int())
				var _1196_v12 _dafny.CodePoint
				_ = _1196_v12
				_1196_v12 = _dafny.CodePoint('b')
				_1196_v12 = _1196_v12
				var _1197_v13 _dafny.Array
				_ = _1197_v13
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
				_ = _nw203
				_1197_v13 = _nw203
				var _1198_v14 _dafny.Array
				_ = _1198_v14
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
				_ = _nw204
				_1198_v14 = _nw204
				var _1199_v15 _dafny.Set
				_ = _1199_v15
				_1199_v15 = _dafny.SetOf(_dafny.IntOfUint32((_1188_v4).Cardinality()), (_this).F27(), (_this).F26(), (_this).F27())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1197_v13), 0))
				_ = _index222
				(_1197_v13).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(696), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1198_v14)).Cardinality())).Intersection(_1199_v15), (_index222).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1197_v13), 0))
				_ = _index223
				var _rhs227 _dafny.Int = (_dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality())).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(261), (_this).F26()))
				_ = _rhs227
				var _rhs228 _dafny.Set = _1199_v15
				_ = _rhs228
				var _rhs229 _dafny.CodePoint = _1196_v12
				_ = _rhs229
				var _lhs194 *GlobalState = globalState
				_ = _lhs194
				var _lhs195 _dafny.Array = _1197_v13
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1197_v13), 0))
				_ = _lhs196
				_lhs194.F0 = _rhs227
				(_lhs195).ArraySet1(_rhs228, (_lhs196).Int())
				_1196_v12 = _rhs229
			} else {
				(globalState).F1 = (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool)
				var _1200_v16 _dafny.Set
				_ = _1200_v16
				_1200_v16 = _dafny.SetOf((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool), (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool))
				var _1201_v17 _dafny.CodePoint
				_ = _1201_v17
				_1201_v17 = _dafny.CodePoint('i')
				var _1202_v18 _dafny.Map
				_ = _1202_v18
				_1202_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1200_v16, _1201_v17)
				var _1203_v19 _dafny.Map
				_ = _1203_v19
				_1203_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1202_v18).Update(_1200_v16, _1201_v17), (_dafny.Zero).Minus(((_this).F26()).Minus((_this).F27())))
				_1203_v19 = (_1203_v19).Update((_1202_v18).Merge(_1202_v18), (_this).F26())
				_1186_v2 = _1186_v2
				var _1204_v20 T0
				_ = _1204_v20
				var _nw205 *C3 = New_C3_()
				_ = _nw205
				_nw205.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-475))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_1205_i0 _dafny.Int) _dafny.CodePoint {
					return (_this.F28).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32()).(_dafny.CodePoint)
				})), _this.F28), p1, ((_this).F26()).Cmp((_dafny.Zero).Minus((_this).F27())) != 0, (_this).F27())
				_1204_v20 = _nw205
				var _1206_v21 _dafny.Array
				_ = _1206_v21
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw206
				_1206_v21 = _nw206
				var _1207_v22 _dafny.Map
				_ = _1207_v22
				_1207_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1206_v21, (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool))
				var _1208_v23 _dafny.Map
				_ = _1208_v23
				_1208_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_dafny.Zero).Minus((_1207_v22).Cardinality())).Minus((_1204_v20).F26()))
				var _1209_v24 D0
				_ = _1209_v24
				_1209_v24 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_this.F28).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("cmdoif"), p2, (_this).F26(), (_this).F25())
				var _rhs230 bool = (_this).F25()
				_ = _rhs230
				var _rhs231 T0 = _1204_v20
				_ = _rhs231
				var _rhs232 bool = !((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool)) || ((_1204_v20).Fm7(globalState))
				_ = _rhs232
				var _rhs233 _dafny.Map = (func() _dafny.Map {
					if (_1204_v20).F25() {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F26())
					}
					return _1208_v23
				})()
				_ = _rhs233
				var _rhs234 D0 = _1209_v24
				_ = _rhs234
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				_lhs197.F1 = _rhs230
				_1204_v20 = _rhs231
				_lhs198.F1 = _rhs232
				_1208_v23 = _rhs233
				_1209_v24 = _rhs234
				var _1210_v25 _dafny.Array
				_ = _1210_v25
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
				_ = _nw207
				_1210_v25 = _nw207
				var _1211_v26 _dafny.Sequence
				_ = _1211_v26
				_1211_v26 = _dafny.SeqOf(_1210_v25, _1210_v25)
				(globalState).F1 = ((_1204_v20).F26()).Cmp(_dafny.IntOfUint32((_1211_v26).Cardinality())) <= 0
			}
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))
			_ = _index224
			(_1189_v5).ArraySet1((_1184_v0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1184_v0).Cardinality()))).Uint32()).(bool), (_index224).Int())
			var _1212_v27 _dafny.Array
			_ = _1212_v27
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_39
			var _nw208 _dafny.Array
			_ = _nw208
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw208 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Sequence = func(_1213_i1 _dafny.Int) _dafny.Sequence {
					return _this.F28
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw208 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw208).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw208).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1212_v27 = _nw208
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))
			_ = _index225
			(_1212_v27).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28), (_index225).Int())
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))
			_ = _index226
			(_1212_v27).ArraySet1(_this.F28, (_index226).Int())
			var _1214_v28 _dafny.Set
			_ = _1214_v28
			_1214_v28 = _dafny.SetOf(_1184_v0)
			if (_1214_v28).IsProperSubsetOf(Companion_Default___.Fm13((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool), globalState)) {
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))
				_ = _index227
				(_1189_v5).ArraySet1(p2, (_index227).Int())
				var _1215_v29 _dafny.MultiSet
				_ = _1215_v29
				_1215_v29 = _dafny.MultiSetOf(_1189_v5)
				var _1216_v30 _dafny.Map
				_ = _1216_v30
				_1216_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool))
				_1215_v29 = ((_1215_v29).Union((_1215_v29).Update(_1189_v5, Companion_Default___.Abs((_1216_v30).Cardinality())))).Difference((_1215_v29).Union(_1215_v29))
				var _1217_v31 _dafny.Int
				_ = _1217_v31
				var _1218_v32 bool
				_ = _1218_v32
				var _1219_v33 _dafny.Map
				_ = _1219_v33
				var _1220_v34 _dafny.Map
				_ = _1220_v34
				var _out24 _dafny.Int
				_ = _out24
				var _out25 bool
				_ = _out25
				var _out26 _dafny.Map
				_ = _out26
				var _out27 _dafny.Map
				_ = _out27
				_out24, _out25, _out26, _out27 = (_this).M2(p1, globalState)
				_1217_v31 = _out24
				_1218_v32 = _out25
				_1219_v33 = _out26
				_1220_v34 = _out27
				var _1221_v35 _dafny.CodePoint
				_ = _1221_v35
				_1221_v35 = _dafny.CodePoint('b')
				_1221_v35 = _1221_v35
				var _1222_v36 _dafny.Sequence
				_ = _1222_v36
				_1222_v36 = _dafny.SeqOf((func() _dafny.Array {
					if p2 {
						return _1189_v5
					}
					return _1189_v5
				})(), _1189_v5, _1189_v5)
				var _rhs235 bool = (func() bool {
					if p2 {
						return _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F28, (_1212_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))).Int()).(_dafny.Sequence))
					}
					return (_this).F25()
				})()
				_ = _rhs235
				var _rhs236 _dafny.Sequence = _dafny.SeqOf(_1189_v5, _1189_v5, _1189_v5, _1189_v5, _1189_v5)
				_ = _rhs236
				_1218_v32 = _rhs235
				_1222_v36 = _rhs236
			} else {
				(_this).F28 = _this.F28
				var _1223_v37 _dafny.MultiSet
				_ = _1223_v37
				_1223_v37 = _dafny.MultiSetOf((_this).F25(), p1)
				(globalState).F1 = ((_1223_v37).Difference(_dafny.MultiSetOf(p1, p1))).IsProperSubsetOf((_dafny.MultiSetOf((_this).F25(), (_this).F25(), (_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool), p1, p1)).Union(_1223_v37))
				var _1224_v38 _dafny.Map
				_ = _1224_v38
				_1224_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1212_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))).Int()).(_dafny.Sequence))
				var _1225_v39 _dafny.Sequence
				_ = _1225_v39
				_1225_v39 = _dafny.SeqOf((_1212_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm1(globalState), _this.F28, _dafny.UnicodeSeqOfUtf8Bytes("wddepf"), (func() _dafny.Sequence {
					if (_1224_v38).Contains((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool)) {
						return (_1224_v38).Get((_1189_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1189_v5), 0))).Int()).(bool)).(_dafny.Sequence)
					}
					return (_1212_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1212_v27), 0))).Int()).(_dafny.Sequence)
				})())
				var _1226_v40 _dafny.Set
				_ = _1226_v40
				_1226_v40 = _dafny.SetOf(Companion_Default___.Fm5((_this).F25(), _1187_v3, _dafny.IntOfUint32((_1184_v0).Cardinality()), globalState))
				var _1227_v41 *C3
				_ = _1227_v41
				var _nw209 *C3 = New_C3_()
				_ = _nw209
				_nw209.Ctor__((_1225_v39).Select((Companion_Default___.SafeIndex((_1226_v40).Cardinality(), _dafny.IntOfUint32((_1225_v39).Cardinality()))).Uint32()).(_dafny.Sequence), !(_1226_v40).Equals(_1226_v40), p1, (_dafny.Zero).Minus((_this).F26()))
				_1227_v41 = _nw209
				(globalState).F1 = !(p2) || (!(true) || (p1))
				var _1228_v42 _dafny.Map
				_ = _1228_v42
				_1228_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1227_v41).F30(), true)
				var _1229_v43 D11
				_ = _1229_v43
				_1229_v43 = Companion_D11_.Create_DC25_(_1228_v42)
				var _1230_v44 _dafny.CodePoint
				_ = _1230_v44
				_1230_v44 = _dafny.CodePoint('j')
				_1229_v43 = Companion_Default___.Fm34(_dafny.Companion_Sequence_.Update((_1227_v41).F29(), (Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32(((_1227_v41).F29()).Cardinality()))).Uint32(), _1230_v44), (_this).F27(), globalState)
			}
		} else {
			var _1231_v45 _dafny.Sequence
			_ = _1231_v45
			_1231_v45 = _dafny.SeqOf((_this).F26(), (_this).F27(), (_this).F27(), _dafny.IntOfInt64(844))
			var _1232_v46 _dafny.Set
			_ = _1232_v46
			_1232_v46 = _dafny.SetOf(p2)
			var _1233_v47 _dafny.Map
			_ = _1233_v47
			_1233_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _1232_v46)
			(globalState).F1 = (_1233_v47).Contains((func() _dafny.Int {
				if (_this).F25() {
					return (_1231_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.IntOfUint32((_1231_v45).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return (_this).F27()
			})())
			r0 = p1
			var _1234_v48 _dafny.Map
			_ = _1234_v48
			_1234_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_dafny.IntOfInt64(649)).Cmp((_this).F26()) > 0)
			var _1235_v49 _dafny.MultiSet
			_ = _1235_v49
			_1235_v49 = _dafny.MultiSetOf(p0)
			var _1236_v50 D0
			_ = _1236_v50
			_1236_v50 = Companion_D0_.Create_DC1_((_1235_v49).Cardinality(), _this.F28, p1, (_this).F26(), p2)
			_1234_v48 = (_1234_v48).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-915), (_this).F27()), (_1236_v50).Dtor_cf3())
			var _1237_v51 _dafny.CodePoint
			_ = _1237_v51
			_1237_v51 = _dafny.CodePoint('s')
			var _1238_v52 _dafny.Map
			_ = _1238_v52
			_1238_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v51, p2)
			var _1239_v54 _dafny.Set
			_ = _1239_v54
			_1239_v54 = _dafny.SetOf(_1237_v51)
			_1238_v52 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), p2)).Merge(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((_1239_v54).Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _1240_v53 _dafny.CodePoint
					_1240_v53 = interface{}(_compr_25).(_dafny.CodePoint)
					if (_1239_v54).Contains(_1240_v53) {
						_coll25.Add(_1240_v53, true)
					}
				}
				return _coll25.ToMap()
			}())
			(_this).F28 = Companion_Default___.Fm1(globalState)
		}
		var _1241_v55 _dafny.Sequence
		_ = _1241_v55
		_1241_v55 = _dafny.SeqOf((_this).F26())
		if (Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27())).Cmp((_1241_v55).Select((Companion_Default___.SafeIndex((_1241_v55).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1241_v55).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1241_v55).Cardinality()))).Uint32()).(_dafny.Int)) != 0 {
			var _1242_v56 D1
			_ = _1242_v56
			_1242_v56 = Companion_D1_.Create_DC4_(false, p1)
			var _source17 D1 = _1242_v56
			_ = _source17
			if _source17.Is_DC4() {
				var _1243___mcc_h0 bool = _source17.Get_().(D1_DC4).Cf7
				_ = _1243___mcc_h0
				var _1244___mcc_h1 bool = _source17.Get_().(D1_DC4).Cf8
				_ = _1244___mcc_h1
				var _1245_cf8 bool = _1244___mcc_h1
				_ = _1245_cf8
				var _1246_cf7 bool = _1243___mcc_h0
				_ = _1246_cf7
				var _1247_v57 D6
				_ = _1247_v57
				_1247_v57 = Companion_D6_.Create_DC14_(_dafny.IntOfInt64(-775), (_this).F26(), _dafny.MultiSetOf(false, !(_1246_cf7)))
				var _1248_v58 _dafny.Set
				_ = _1248_v58
				_1248_v58 = _dafny.SetOf((_1247_v57).Dtor_cf23())
				var _1249_v59 _dafny.Map
				_ = _1249_v59
				_1249_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F26())
				var _rhs237 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28), _this.F28)
				_ = _rhs237
				var _rhs238 bool = ((_1248_v58).Cardinality()).Cmp((_this).F26()) > 0
				_ = _rhs238
				var _rhs239 _dafny.Set = ((func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter26 := _dafny.Iterate((_1249_v59).Keys().Elements()); ; {
						_compr_26, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _1250_v60 _dafny.Int
						_1250_v60 = interface{}(_compr_26).(_dafny.Int)
						if (_1249_v59).Contains(_1250_v60) {
							_coll26.Add((_1250_v60).Minus(_dafny.IntOfInt64(746)))
						}
					}
					return _coll26.ToSet()
				}()).Difference(_1248_v58)).Intersection(_dafny.SetOf((_this).F27()))
				_ = _rhs239
				var _lhs199 *C4 = _this
				_ = _lhs199
				_lhs199.F28 = _rhs237
				_1245_cf8 = _rhs238
				_1248_v58 = _rhs239
				_1184_v0 = _1184_v0
				var _1251_v61 D2
				_ = _1251_v61
				_1251_v61 = Companion_D2_.Create_DC7_(_1246_cf7, _1245_cf8, (_this).F27())
				var _1252_v62 _dafny.Array
				_ = _1252_v62
				var _nwElement0_41 _dafny.Int = _dafny.IntOfUint32((_this.F28).Cardinality())
				_ = _nwElement0_41
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(20))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_41, 0)
				(_nw210).ArraySet1((_this).F26(), 1)
				(_nw210).ArraySet1(((_1241_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-480), _dafny.IntOfUint32((_1241_v55).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_this).F27()), 2)
				(_nw210).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(901), (p0).Cardinality()), 3)
				(_nw210).ArraySet1((_this).F26(), 4)
				(_nw210).ArraySet1(Companion_Default___.Fm2((_this).F27(), (p0).Cardinality(), globalState), 5)
				(_nw210).ArraySet1(((_dafny.Zero).Minus((_1248_v58).Cardinality())).Minus((_this).F26()), 6)
				(_nw210).ArraySet1((_this).F26(), 7)
				(_nw210).ArraySet1(Companion_Default___.Fm2((_this).F26(), Companion_Default___.Fm2((_this).F27(), _dafny.IntOfUint32((_1241_v55).Cardinality()), globalState), globalState), 8)
				(_nw210).ArraySet1(_dafny.IntOfUint32((_1184_v0).Cardinality()), 9)
				(_nw210).ArraySet1(((_this).F27()).Minus((_1251_v61).Dtor_cf13()), 10)
				(_nw210).ArraySet1((_this).F26(), 11)
				(_nw210).ArraySet1((_this).F26(), 12)
				(_nw210).ArraySet1((_1247_v57).Dtor_cf24(), 13)
				(_nw210).ArraySet1((_this).F27(), 14)
				(_nw210).ArraySet1(((_this).F26()).Times((_this).F26()), 15)
				(_nw210).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F27(), (_this).F27())), 16)
				(_nw210).ArraySet1((_this).F27(), 17)
				(_nw210).ArraySet1((_this).F27(), 18)
				(_nw210).ArraySet1((_this).F27(), 19)
				_1252_v62 = _nw210
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw211
				_1252_v62 = _nw211
				var _1253_v63 _dafny.Map
				_ = _1253_v63
				_1253_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v56, (_this).F27())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1252_v62), 0))
				_ = _index228
				(_1252_v62).ArraySet1((_dafny.Zero).Minus((_this).F26()), (_index228).Int())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1252_v62), 0))
				_ = _index229
				var _rhs240 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(p1, true), (_dafny.IntOfInt64(8)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-569))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}(func(_1254_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				}))).Cardinality())))
				_ = _rhs240
				var _rhs241 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F27()), (_this).F26())
				_ = _rhs241
				var _lhs200 _dafny.Array = _1252_v62
				_ = _lhs200
				var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1252_v62), 0))
				_ = _lhs201
				_1253_v63 = _rhs240
				(_lhs200).ArraySet1(_rhs241, (_lhs201).Int())
			} else if _source17.Is_DC3() {
				var _1255___mcc_h2 _dafny.CodePoint = _source17.Get_().(D1_DC3).Cf6
				_ = _1255___mcc_h2
				var _1256_cf6 _dafny.CodePoint = _1255___mcc_h2
				_ = _1256_cf6
				var _1257_v64 D2
				_ = _1257_v64
				_1257_v64 = Companion_D2_.Create_DC6_(_dafny.SetOf((_this).F26()))
				_1257_v64 = _1257_v64
				var _1258_v65 _dafny.Map
				_ = _1258_v65
				_1258_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())
				var _1259_v66 _dafny.Map
				_ = _1259_v66
				_1259_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F27())
				var _1260_v67 D0
				_ = _1260_v67
				_1260_v67 = Companion_D0_.Create_DC1_((_this).F27(), _dafny.UnicodeSeqOfUtf8Bytes("ssmnmk"), false, (_this).F27(), (_this).F25())
				var _1261_v68 _dafny.Array
				_ = _1261_v68
				var _nwElement0_42 _dafny.Int = (func() _dafny.Int {
					if (_1258_v65).Contains((_this).F27()) {
						return (_1258_v65).Get((_this).F27()).(_dafny.Int)
					}
					return (_this).F27()
				})()
				_ = _nwElement0_42
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(27))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_42, 0)
				(_nw212).ArraySet1(Companion_Default___.SafeModuloInt(((_1259_v66).Update(p1, (_this).F26())).Cardinality(), (_this).F27()), 1)
				(_nw212).ArraySet1((_1260_v67).Dtor_cf4(), 2)
				(_nw212).ArraySet1(((_this).F26()).Times((_this).F26()), 3)
				(_nw212).ArraySet1((func() _dafny.Int {
					if (_1259_v66).Contains(!(p2)) {
						return (_1259_v66).Get(!(p2)).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-159)
				})(), 4)
				(_nw212).ArraySet1((_dafny.MultiSetFromSeq(_this.F28)).Cardinality(), 5)
				(_nw212).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()), 6)
				(_nw212).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27()), 7)
				(_nw212).ArraySet1((_this).F27(), 8)
				(_nw212).ArraySet1((_this).F26(), 9)
				(_nw212).ArraySet1((_this).F26(), 10)
				(_nw212).ArraySet1(_dafny.IntOfInt64(619), 11)
				(_nw212).ArraySet1(Companion_Default___.Fm2((_this).F26(), _dafny.IntOfInt64(731), globalState), 12)
				(_nw212).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-263)), 13)
				(_nw212).ArraySet1(((_this).F26()).Times((_this).F26()), 14)
				(_nw212).ArraySet1((_this).F26(), 15)
				(_nw212).ArraySet1((_this).F27(), 16)
				(_nw212).ArraySet1(Companion_Default___.SafeModuloInt((_this).F27(), (_this).F26()), 17)
				(_nw212).ArraySet1((_this).F26(), 18)
				(_nw212).ArraySet1((_this).F27(), 19)
				(_nw212).ArraySet1(((_this).F26()).Times((_this).F26()), 20)
				(_nw212).ArraySet1((_this).F26(), 21)
				(_nw212).ArraySet1((func() _dafny.Int {
					if (p0).Contains((_this).F26()) {
						return (p0).Multiplicity((_this).F26())
					}
					return (_this).F26()
				})(), 22)
				(_nw212).ArraySet1((_this).F26(), 23)
				(_nw212).ArraySet1((_this).F27(), 24)
				(_nw212).ArraySet1((_this).F26(), 25)
				(_nw212).ArraySet1(_dafny.IntOfUint32((_this.F28).Cardinality()), 26)
				_1261_v68 = _nw212
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw213
				_1261_v68 = _nw213
				(globalState).F0 = (_this).F26()
				(globalState).F1 = (func() bool {
					if !(p2) || (p2) {
						return p1
					}
					return !((_this).F25()) || (p1)
				})()
			} else {
				var _1262___mcc_h3 D1 = _source17.Get_().(D1_DC5).Cf9
				_ = _1262___mcc_h3
				var _1263_cf9 D1 = _1262___mcc_h3
				_ = _1263_cf9
				var _1264_v69 _dafny.Array
				_ = _1264_v69
				var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw214
				_1264_v69 = _nw214
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_40
				var _nw215 _dafny.Array
				_ = _nw215
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw215 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Int = func(_1265_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1265_i3, (_this).F26())
					}
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw215 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw215).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw215).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1264_v69 = _nw215
				var _1266_v70 D0
				_ = _1266_v70
				_1266_v70 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_this.F28).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("dl"), true, (_this).F26(), !(false))
				var _1267_v71 _dafny.CodePoint
				_ = _1267_v71
				_1267_v71 = _dafny.CodePoint('q')
				(_this).F28 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_1266_v70).Dtor_cf2(), (func() _dafny.Sequence {
					if (_this).F25() {
						return _this.F28
					}
					return _this.F28
				})()), (Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Cardinality()).Plus((_this).F26()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1266_v70).Dtor_cf2(), (func() _dafny.Sequence {
					if (_this).F25() {
						return _this.F28
					}
					return _this.F28
				})())).Cardinality()))).Uint32(), _1267_v71)
				_1242_v56 = Companion_D1_.Create_DC4_(((_this).F26()).Cmp((_this).F26()) > 0, (_this).F25())
				var _1268_v72 _dafny.Array
				_ = _1268_v72
				var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
				_ = _nw216
				_1268_v72 = _nw216
				var _1269_v73 _dafny.Array
				_ = _1269_v73
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_41
				var _nw217 _dafny.Array
				_ = _nw217
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw217 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) _dafny.Map = (func(_1270_p1 bool) func(_dafny.Int) _dafny.Map {
						return func(_1271_i4 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F25()), _1270_p1)
						}
					})(p1)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw217 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw217).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw217).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1269_v73 = _nw217
				var _1272_v74 D8
				_ = _1272_v74
				_1272_v74 = Companion_D8_.Create_DC16_(_1269_v73)
				var _1273_v75 _dafny.Map
				_ = _1273_v75
				_1273_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
				var _1274_v76 D12
				_ = _1274_v76
				_1274_v76 = Companion_D12_.Create_DC28_(_1264_v69)
				var _rhs242 _dafny.Int = _dafny.IntOfInt64(576)
				_ = _rhs242
				var _rhs243 _dafny.Array = (func() _dafny.Array {
					if !(Companion_Default___.Fm5((_this).F25(), Companion_Default___.Fm25((_this).F26(), (_this).F25(), (_this).F25(), globalState), (_1273_v75).Cardinality(), globalState)) {
						return (_1274_v76).Dtor_cf44()
					}
					return _1264_v69
				})()
				_ = _rhs243
				var _rhs244 _dafny.Array = _1268_v72
				_ = _rhs244
				var _rhs245 D8 = _1272_v74
				_ = _rhs245
				var _rhs246 bool = (_this).F25()
				_ = _rhs246
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				var _lhs203 *GlobalState = globalState
				_ = _lhs203
				_lhs202.F2 = _rhs242
				_1264_v69 = _rhs243
				_1268_v72 = _rhs244
				_1272_v74 = _rhs245
				_lhs203.F1 = _rhs246
			}
			var _1275_v77 _dafny.Int
			_ = _1275_v77
			var _out28 _dafny.Int
			_ = _out28
			_out28 = Companion_Default___.M0((func() bool {
				if (_this).F25() {
					return true
				}
				return (_this).F25()
			})(), globalState)
			_1275_v77 = _out28
			var _1276_v78 _dafny.Array
			_ = _1276_v78
			var _nwElement0_43 bool = (_this).F25()
			_ = _nwElement0_43
			var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(19))
			_ = _nw218
			(_nw218).ArraySet1(_nwElement0_43, 0)
			(_nw218).ArraySet1((_this).F25(), 1)
			(_nw218).ArraySet1(p1, 2)
			(_nw218).ArraySet1(p1, 3)
			(_nw218).ArraySet1(p2, 4)
			(_nw218).ArraySet1(!((_this).F25()), 5)
			(_nw218).ArraySet1((_this).Fm7(globalState), 6)
			(_nw218).ArraySet1((_this).F25(), 7)
			(_nw218).ArraySet1(p2, 8)
			(_nw218).ArraySet1((_this).F25(), 9)
			(_nw218).ArraySet1(p2, 10)
			(_nw218).ArraySet1(p2, 11)
			(_nw218).ArraySet1((_this).F25(), 12)
			(_nw218).ArraySet1(p2, 13)
			(_nw218).ArraySet1(p1, 14)
			(_nw218).ArraySet1(p2, 15)
			(_nw218).ArraySet1(false, 16)
			(_nw218).ArraySet1((_this).F25(), 17)
			(_nw218).ArraySet1(p2, 18)
			_1276_v78 = _nw218
			var _1277_v79 _dafny.CodePoint
			_ = _1277_v79
			_1277_v79 = _dafny.CodePoint('p')
			var _1278_v80 _dafny.Map
			_ = _1278_v80
			_1278_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1276_v78, _1277_v79)
			var _1279_v81 _dafny.Set
			_ = _1279_v81
			_1279_v81 = _dafny.SetOf(_1275_v77, _1275_v77, (_this).F27(), (_1278_v80).Cardinality(), _dafny.IntOfUint32((_this.F28).Cardinality()))
			var _1280_v82 _dafny.MultiSet
			_ = _1280_v82
			_1280_v82 = _dafny.MultiSetOf(_1279_v81)
			_1280_v82 = (_1280_v82).Difference((func() _dafny.MultiSet {
				if (_this).F25() {
					return _1280_v82
				}
				return _1280_v82
			})())
			var _1281_v83 _dafny.Array
			_ = _1281_v83
			var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw219
			_1281_v83 = _nw219
			_1281_v83 = _1281_v83
			var _1282_v84 *C1
			_ = _1282_v84
			var _nw220 *C1 = New_C1_()
			_ = _nw220
			_nw220.Ctor__((_this).F25())
			_1282_v84 = _nw220
			var _1283_v85 _dafny.Map
			_ = _1283_v85
			_1283_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() *C1 {
				if false {
					return _1282_v84
				}
				return _1282_v84
			})())
			_1283_v85 = (_1283_v85).Update(!(p2), _1282_v84)
		} else {
			if (_this).F25() {
				var _1284_v86 _dafny.CodePoint
				_ = _1284_v86
				_1284_v86 = _dafny.CodePoint('x')
				var _1285_v87 D6
				_ = _1285_v87
				_1285_v87 = Companion_D6_.Create_DC14_((_this).F26(), (_this).F26(), _dafny.MultiSetOf(p2, (_this).F25()))
				var _1286_v88 D9
				_ = _1286_v88
				_1286_v88 = Companion_D9_.Create_DC19_(_1285_v87)
				var _1287_v89 D9
				_ = _1287_v89
				_1287_v89 = Companion_D9_.Create_DC20_((func() D9 {
					if p1 {
						return _1286_v88
					}
					return _1286_v88
				})())
				var _1288_v90 _dafny.Set
				_ = _1288_v90
				_1288_v90 = _dafny.SetOf(!((_1184_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(760))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_1289_i5 _dafny.Int) _dafny.Sequence {
					return _this.F28
				}))).Cardinality()), _dafny.IntOfUint32((_1184_v0).Cardinality()))).Uint32()).(bool)))
				var _pat_let_tv33 = _1286_v88
				_ = _pat_let_tv33
				var _rhs247 _dafny.CodePoint = _1284_v86
				_ = _rhs247
				var _rhs248 bool = !(p2) || (!((_1288_v90).Equals(_1288_v90)))
				_ = _rhs248
				var _rhs249 D9 = func(_pat_let38_0 D9) D9 {
					return func(_1290_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let39_0 D9) D9 {
							return func(_1291_dt__update_hcf32_h0 D9) D9 {
								return Companion_D9_.Create_DC20_(_1291_dt__update_hcf32_h0)
							}(_pat_let39_0)
						}(Companion_D9_.Create_DC20_(_pat_let_tv33))
					}(_pat_let38_0)
				}(_1287_v89)
				_ = _rhs249
				var _rhs250 bool = (_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F25() {
						return _this.F28
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg59 _dafny.Int) interface{} {
							return coer59(arg59)
						}
					}((func(_1292_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1293_i6 _dafny.Int) _dafny.CodePoint {
							return _1292_v86
						}
					})(_1284_v86)))
				})()).Cardinality())).Cmp((_dafny.IntOfUint32((_1241_v55).Cardinality())).Times((_this).F26())) < 0
				_ = _rhs250
				var _lhs204 *GlobalState = globalState
				_ = _lhs204
				_1284_v86 = _rhs247
				r0 = _rhs248
				_1287_v89 = _rhs249
				_lhs204.F1 = _rhs250
				var _1294_v91 *C3
				_ = _1294_v91
				var _nw221 *C3 = New_C3_()
				_ = _nw221
				_nw221.Ctor__(_this.F28, p2, p2, (_this).F26())
				_1294_v91 = _nw221
				var _1295_v92 _dafny.Sequence
				_ = _1295_v92
				_1295_v92 = _dafny.SeqOf(Companion_Default___.Fm1(globalState))
				_1295_v92 = _dafny.SeqOf(_this.F28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(150))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_1296_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1297_i7 _dafny.Int) _dafny.CodePoint {
						return _1296_v86
					}
				})(_1284_v86))), _this.F28, _this.F28)
				var _1298_v93 _dafny.Array
				_ = _1298_v93
				var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(16))
				_ = _nw222
				_1298_v93 = _nw222
				var _1299_v94 D10
				_ = _1299_v94
				_1299_v94 = Companion_D10_.Create_DC22_(_1298_v93, _this.F28)
				var _1300_v95 _dafny.Map
				_ = _1300_v95
				_1300_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28, (_1299_v94).Dtor_cf35())
				(globalState).F1 = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
					if (_1300_v95).Contains(_this.F28) {
						return (_1300_v95).Get(_this.F28).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("rj")
				})(), _this.F28)
				var _1301_v96 _dafny.Array
				_ = _1301_v96
				var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw223
				_1301_v96 = _nw223
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1301_v96), 0))
				_ = _index230
				(_1301_v96).ArraySet1((_this).F27(), (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1301_v96), 0))
				_ = _index231
				(_1301_v96).ArraySet1((_this).F27(), (_index231).Int())
			} else {
				(globalState).F1 = p2
				(globalState).F1 = p1
				var _1302_v97 _dafny.Map
				_ = _1302_v97
				_1302_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), false)
				var _1303_v98 *C1
				_ = _1303_v98
				var _nw224 *C1 = New_C1_()
				_ = _nw224
				_nw224.Ctor__(p2)
				_1303_v98 = _nw224
				var _1304_v99 _dafny.CodePoint
				_ = _1304_v99
				_1304_v99 = _dafny.CodePoint('i')
				var _1305_v100 D10
				_ = _1305_v100
				_1305_v100 = Companion_D10_.Create_DC24_((_this).F27(), _1303_v98, _1304_v99)
				var _1306_v101 _dafny.Map
				_ = _1306_v101
				_1306_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1305_v100, _dafny.MultiSetOf((_this).F27()))
				_1302_v97 = (_1302_v97).Update(!(false), ((_dafny.MultiSetOf(p2)).Cardinality()).Cmp(((func() _dafny.MultiSet {
					if (_1306_v101).Contains(_1305_v100) {
						return (_1306_v101).Get(_1305_v100).(_dafny.MultiSet)
					}
					return p0
				})()).Cardinality()) == 0)
				var _1307_v102 *C1
				_ = _1307_v102
				var _nw225 *C1 = New_C1_()
				_ = _nw225
				_nw225.Ctor__(p1)
				_1307_v102 = _nw225
				(globalState).F1 = (_this).F25()
			}
			var _1308_v103 _dafny.Array
			_ = _1308_v103
			var _nw226 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw226
			_1308_v103 = _nw226
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))
			_ = _index232
			(_1308_v103).ArraySet1((_this).F25(), (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))
			_ = _index233
			(_1308_v103).ArraySet1(true, (_index233).Int())
			var _1309_v104 _dafny.Array
			_ = _1309_v104
			var _nwElement0_44 _dafny.Array = _1308_v103
			_ = _nwElement0_44
			var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(15))
			_ = _nw227
			(_nw227).ArraySet1(_nwElement0_44, 0)
			(_nw227).ArraySet1(_1308_v103, 1)
			(_nw227).ArraySet1(_1308_v103, 2)
			(_nw227).ArraySet1(_1308_v103, 3)
			(_nw227).ArraySet1(_1308_v103, 4)
			(_nw227).ArraySet1(_1308_v103, 5)
			(_nw227).ArraySet1(_1308_v103, 6)
			(_nw227).ArraySet1(_1308_v103, 7)
			(_nw227).ArraySet1(_1308_v103, 8)
			(_nw227).ArraySet1(_1308_v103, 9)
			(_nw227).ArraySet1(_1308_v103, 10)
			(_nw227).ArraySet1(_1308_v103, 11)
			(_nw227).ArraySet1((func() _dafny.Array {
				if p2 {
					return _1308_v103
				}
				return _1308_v103
			})(), 12)
			(_nw227).ArraySet1(_1308_v103, 13)
			(_nw227).ArraySet1(_1308_v103, 14)
			_1309_v104 = _nw227
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1309_v104), 0))
			_ = _index234
			(_1309_v104).ArraySet1(_1308_v103, (_index234).Int())
			var _1310_v105 _dafny.Sequence
			_ = _1310_v105
			_1310_v105 = _dafny.SeqOf(_1308_v103, _1308_v103, _1308_v103, _1308_v103)
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1309_v104), 0))
			_ = _index235
			(_1309_v104).ArraySet1((_1310_v105).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.IntOfUint32((_1310_v105).Cardinality()))).Uint32()).(_dafny.Array), (_index235).Int())
			var _1311_v106 *C1
			_ = _1311_v106
			var _nw228 *C1 = New_C1_()
			_ = _nw228
			_nw228.Ctor__(((_1308_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))).Int()).(bool)) && (p2))
			_1311_v106 = _nw228
			if (Companion_Default___.SafeModuloInt((_this).F27(), (_this).F26())).Cmp((_this).F27()) >= 0 {
				var _1312_v107 _dafny.Map
				_ = _1312_v107
				_1312_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_1311_v106).Fm14(_dafny.SetOf(p2), (_this).F26(), globalState))
				var _1313_v108 _dafny.Set
				_ = _1313_v108
				_1313_v108 = _dafny.SetOf((_1312_v107).Cardinality(), (_this).F26())
				_1313_v108 = (_1313_v108).Difference(func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-465), _dafny.IntOfInt64(-266))); ; {
						_compr_27, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _1314_v109 _dafny.Int
						_1314_v109 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(-465)).Cmp(_1314_v109) <= 0) && ((_1314_v109).Cmp(_dafny.IntOfInt64(-266)) < 0) {
							_coll27.Add((_1314_v109).Times((_this).F27()))
						}
					}
					return _coll27.ToSet()
				}())
				var _1315_v110 _dafny.CodePoint
				_ = _1315_v110
				_1315_v110 = _dafny.CodePoint('m')
				var _1316_v111 _dafny.MultiSet
				_ = _1316_v111
				_1316_v111 = _dafny.MultiSetOf(_1315_v110)
				var _1317_v112 _dafny.Map
				_ = _1317_v112
				_1317_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1311_v106).F32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qtwjddafo")).Cardinality()))
				var _1318_v113 _dafny.Sequence
				_ = _1318_v113
				_1318_v113 = _dafny.SeqOf(_1316_v111)
				var _1319_v114 _dafny.Array
				_ = _1319_v114
				var _nwElement0_45 _dafny.MultiSet = _1316_v111
				_ = _nwElement0_45
				var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(22))
				_ = _nw229
				(_nw229).ArraySet1(_nwElement0_45, 0)
				(_nw229).ArraySet1(_1316_v111, 1)
				(_nw229).ArraySet1(_1316_v111, 2)
				(_nw229).ArraySet1((_1316_v111).Intersection(_1316_v111), 3)
				(_nw229).ArraySet1(_1316_v111, 4)
				(_nw229).ArraySet1(_1316_v111, 5)
				(_nw229).ArraySet1((func() _dafny.MultiSet {
					if p1 {
						return _1316_v111
					}
					return Companion_Default___.Fm35((_this).F27(), (_1317_v112).Cardinality(), globalState)
				})(), 6)
				(_nw229).ArraySet1(_1316_v111, 7)
				(_nw229).ArraySet1((_1316_v111).Difference(_dafny.MultiSetOf(_1315_v110)), 8)
				(_nw229).ArraySet1((_1316_v111).Difference(_dafny.MultiSetOf(_1315_v110)), 9)
				(_nw229).ArraySet1(_1316_v111, 10)
				(_nw229).ArraySet1(_dafny.MultiSetOf(_1315_v110), 11)
				(_nw229).ArraySet1((func() _dafny.MultiSet {
					if (_1308_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))).Int()).(bool) {
						return _1316_v111
					}
					return _dafny.MultiSetOf(_1315_v110, _1315_v110, _1315_v110)
				})(), 12)
				(_nw229).ArraySet1(_1316_v111, 13)
				(_nw229).ArraySet1((_1316_v111).Update(_1315_v110, Companion_Default___.Abs((_dafny.Zero).Minus((_this).F26()))), 14)
				(_nw229).ArraySet1(_1316_v111, 15)
				(_nw229).ArraySet1(_dafny.MultiSetOf(_1315_v110, (_this.F28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1184_v0).Cardinality()), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32()).(_dafny.CodePoint)), 16)
				(_nw229).ArraySet1((_1318_v113).Select((Companion_Default___.SafeIndex(((_dafny.MultiSetOf((_1311_v106).F32(), true)).Update((_this).F25(), Companion_Default___.Abs((_this).F27()))).Cardinality(), _dafny.IntOfUint32((_1318_v113).Cardinality()))).Uint32()).(_dafny.MultiSet), 17)
				(_nw229).ArraySet1(_1316_v111, 18)
				(_nw229).ArraySet1(_1316_v111, 19)
				(_nw229).ArraySet1((_1316_v111).Difference(_dafny.MultiSetOf(_1315_v110, _1315_v110)), 20)
				(_nw229).ArraySet1(_1316_v111, 21)
				_1319_v114 = _nw229
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1319_v114), 0))
				_ = _index236
				(_1319_v114).ArraySet1(_1316_v111, (_index236).Int())
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1319_v114), 0))
				_ = _index237
				(_1319_v114).ArraySet1(_1316_v111, (_index237).Int())
				(globalState).F0 = (_this).F27()
				var _1320_v115 _dafny.Map
				_ = _1320_v115
				_1320_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1309_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1309_v104), 0))).Int())), (_this).F26())
				var _1321_v116 _dafny.Int
				_ = _1321_v116
				var _1322_v117 bool
				_ = _1322_v117
				var _1323_v118 _dafny.Map
				_ = _1323_v118
				var _1324_v119 _dafny.Map
				_ = _1324_v119
				var _out29 _dafny.Int
				_ = _out29
				var _out30 bool
				_ = _out30
				var _out31 _dafny.Map
				_ = _out31
				var _out32 _dafny.Map
				_ = _out32
				_out29, _out30, _out31, _out32 = (_this).M2(((_this).F26()).Cmp((_1320_v115).Cardinality()) == 0, globalState)
				_1321_v116 = _out29
				_1322_v117 = _out30
				_1323_v118 = _out31
				_1324_v119 = _out32
				r0 = p1
			} else {
				var _1325_v120 _dafny.Set
				_ = _1325_v120
				_1325_v120 = _dafny.SetOf((_this).F27())
				var _1326_v121 _dafny.Array
				_ = _1326_v121
				var _nwElement0_46 _dafny.MultiSet = Companion_Default___.Fm21(_this.F28, (_this).F27(), (_1325_v120).Cardinality(), globalState)
				_ = _nwElement0_46
				var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(10))
				_ = _nw230
				(_nw230).ArraySet1(_nwElement0_46, 0)
				(_nw230).ArraySet1(p0, 1)
				(_nw230).ArraySet1((p0).Update((_this).F26(), Companion_Default___.Abs(_dafny.IntOfUint32((_this.F28).Cardinality()))), 2)
				(_nw230).ArraySet1(p0, 3)
				(_nw230).ArraySet1(_dafny.MultiSetOf((_this).F26()), 4)
				(_nw230).ArraySet1(p0, 5)
				(_nw230).ArraySet1(_dafny.MultiSetFromSeq(_1241_v55), 6)
				(_nw230).ArraySet1(p0, 7)
				(_nw230).ArraySet1(p0, 8)
				(_nw230).ArraySet1(p0, 9)
				_1326_v121 = _nw230
				var _1327_v122 D10
				_ = _1327_v122
				_1327_v122 = Companion_D10_.Create_DC21_(_1326_v121)
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))
				_ = _index238
				var _rhs251 D10 = Companion_D10_.Create_DC21_(_1326_v121)
				_ = _rhs251
				var _rhs252 bool = (_this).F25()
				_ = _rhs252
				var _lhs205 _dafny.Array = _1308_v103
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))
				_ = _lhs206
				_1327_v122 = _rhs251
				(_lhs205).ArraySet1(_rhs252, (_lhs206).Int())
				var _1328_v123 _dafny.Set
				_ = _1328_v123
				_1328_v123 = _dafny.SetOf((_this).F25())
				var _1329_v124 _dafny.Sequence
				_ = _1329_v124
				_1329_v124 = _dafny.SeqOf(_1328_v123)
				var _1330_v125 _dafny.MultiSet
				_ = _1330_v125
				_1330_v125 = _dafny.MultiSetOf((_1311_v106).F32(), (_this).F25(), (_1311_v106).F32(), (_1311_v106).Fm14(_1328_v123, (_this).F27(), globalState), (_1308_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))).Int()).(bool))
				var _rhs253 _dafny.Sequence = _1329_v124
				_ = _rhs253
				var _rhs254 bool = p1
				_ = _rhs254
				var _rhs255 bool = (func() bool {
					if (_1330_v125).Equals(_dafny.MultiSetOf((_1311_v106).F32(), !(p1))) {
						return false
					}
					return p1
				})()
				_ = _rhs255
				var _rhs256 _dafny.Int = (_dafny.Zero).Minus((_this).F26())
				_ = _rhs256
				var _lhs207 *GlobalState = globalState
				_ = _lhs207
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				var _lhs209 *GlobalState = globalState
				_ = _lhs209
				_1329_v124 = _rhs253
				_lhs207.F1 = _rhs254
				_lhs208.F1 = _rhs255
				_lhs209.F0 = _rhs256
				var _1331_v126 _dafny.Array
				_ = _1331_v126
				var _nw231 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw231
				_1331_v126 = _nw231
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1331_v126), 0))
				_ = _index239
				(_1331_v126).ArraySet1(_1184_v0, (_index239).Int())
				var _1332_v127 _dafny.Sequence
				_ = _1332_v127
				_1332_v127 = _1184_v0
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1331_v126), 0))
				_ = _index240
				var _rhs257 _dafny.Sequence = _1332_v127
				_ = _rhs257
				var _rhs258 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28)).Cardinality())
				_ = _rhs258
				var _rhs259 _dafny.MultiSet = (func() _dafny.MultiSet {
					if p2 {
						return (_dafny.MultiSetOf((_1311_v106).F32(), !((_1311_v106).F32()))).Union(Companion_Default___.Fm27(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1309_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1309_v104), 0))).Int())), (_1308_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1308_v103), 0))).Int()).(bool))).Update(_dafny.ArrayCastTo((_1309_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1309_v104), 0))).Int())), true)).Cardinality(), globalState))
					}
					return _dafny.MultiSetFromSeq(_1184_v0)
				})()
				_ = _rhs259
				var _lhs210 _dafny.Array = _1331_v126
				_ = _lhs210
				var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1331_v126), 0))
				_ = _lhs211
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				(_lhs210).ArraySet1(_rhs257, (_lhs211).Int())
				_lhs212.F0 = _rhs258
				_1330_v125 = _rhs259
				(globalState).F2 = (_this).F27()
				var _1333_v129 _dafny.Sequence
				_ = _1333_v129
				_1333_v129 = _dafny.SeqOf((_1325_v120).Union(_1325_v120), (_1325_v120).Difference(func() _dafny.Set {
					var _coll28 = _dafny.NewBuilder()
					_ = _coll28
					for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(141), _dafny.IntOfInt64(841))); ; {
						_compr_28, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _1334_v128 _dafny.Int
						_1334_v128 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(141)).Cmp(_1334_v128) <= 0) && ((_1334_v128).Cmp(_dafny.IntOfInt64(841)) < 0) {
							_coll28.Add(Companion_Default___.SafeDivisionInt(_1334_v128, (_this).F27()))
						}
					}
					return _coll28.ToSet()
				}()), _1325_v120)
				var _rhs260 _dafny.Int = ((_1333_v129).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1333_v129).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
				_ = _rhs260
				var _rhs261 bool = (_1311_v106).F32()
				_ = _rhs261
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				_lhs213.F0 = _rhs260
				_lhs214.F1 = _rhs261
			}
		}
		var _1335_v130 _dafny.Array
		_ = _1335_v130
		var _nw232 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
		_ = _nw232
		_1335_v130 = _nw232
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1335_v130), 0))
		_ = _index241
		(_1335_v130).ArraySet1(p2, (_index241).Int())
		var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1335_v130), 0))
		_ = _index242
		(_1335_v130).ArraySet1(p1, (_index242).Int())
		var _1336_v131 D1
		_ = _1336_v131
		_1336_v131 = Companion_D1_.Create_DC4_(p2, (_1335_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1335_v130), 0))).Int()).(bool))
		var _1337_v132 D1
		_ = _1337_v132
		_1337_v132 = Companion_D1_.Create_DC5_(_1336_v131)
		(globalState).F0 = func(_source18 D1) _dafny.Int {
			if _source18.Is_DC4() {
				var _1338___mcc_h4 bool = _source18.Get_().(D1_DC4).Cf7
				_ = _1338___mcc_h4
				var _1339___mcc_h5 bool = _source18.Get_().(D1_DC4).Cf8
				_ = _1339___mcc_h5
				var _1340_cf8 bool = _1339___mcc_h5
				_ = _1340_cf8
				var _1341_cf7 bool = _1338___mcc_h4
				_ = _1341_cf7
				return ((_dafny.SetOf(_dafny.IntOfUint32((_this.F28).Cardinality()), (_this).F27())).Intersection(_dafny.SetOf((_this).F26()))).Cardinality()
			} else if _source18.Is_DC3() {
				var _1342___mcc_h6 _dafny.CodePoint = _source18.Get_().(D1_DC3).Cf6
				_ = _1342___mcc_h6
				var _1343_cf6 _dafny.CodePoint = _1342___mcc_h6
				_ = _1343_cf6
				return (_this).F27()
			} else {
				var _1344___mcc_h7 D1 = _source18.Get_().(D1_DC5).Cf9
				_ = _1344___mcc_h7
				var _1345_cf9 D1 = _1344___mcc_h7
				_ = _1345_cf9
				return (_this).F27()
			}
		}(_1337_v132)
		var _source19 _dafny.Sequence = _1184_v0
		_ = _source19
		var _1346___mcc_h8 _dafny.Sequence = _source19
		_ = _1346___mcc_h8
		var _1347_cf26 _dafny.Sequence = _1346___mcc_h8
		_ = _1347_cf26
		(globalState).F0 = (_this).F27()
		(globalState).F1 = ((_this).F26()).Cmp((_this).F26()) >= 0
		var _1348_v133 _dafny.CodePoint
		_ = _1348_v133
		_1348_v133 = _dafny.CodePoint('q')
		_1348_v133 = (func() _dafny.CodePoint {
			if p2 {
				return _1348_v133
			}
			return _1348_v133
		})()
		(globalState).F0 = (_this).F27()
		r0 = (func() bool {
			if (_1335_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1335_v130), 0))).Int()).(bool) {
				return p2
			}
			return (_1335_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1335_v130), 0))).Int()).(bool)
		})()
		return r0
	}
}
func (_this *C4) M2(p0 bool, globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, _dafny.Map) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Map = _dafny.EmptyMap
		_ = r3
		var _1349_v0 _dafny.Sequence
		_ = _1349_v0
		_1349_v0 = _dafny.SeqOf((_this).F25(), (_this).F25(), p0, p0, !((_this).F25()))
		var _1350_v1 _dafny.Sequence
		_ = _1350_v1
		_1350_v1 = _dafny.SeqOf(_1349_v0)
		var _hi12 _dafny.Int = ((_this).F27()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1350_v1).Cardinality())))
		_ = _hi12
		for _1351_i0 := (_this).F27(); _1351_i0.Cmp(_hi12) < 0; _1351_i0 = _1351_i0.Plus(_dafny.One) {
			var _1352_v2 _dafny.CodePoint
			_ = _1352_v2
			_1352_v2 = _dafny.CodePoint('s')
			var _1353_v3 _dafny.Map
			_ = _1353_v3
			_1353_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1352_v2, (_this).F25())
			var _1354_v4 _dafny.Array
			_ = _1354_v4
			var _nwElement0_47 bool = (_this).F25()
			_ = _nwElement0_47
			var _nw233 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(16))
			_ = _nw233
			(_nw233).ArraySet1(_nwElement0_47, 0)
			(_nw233).ArraySet1(p0, 1)
			(_nw233).ArraySet1(false, 2)
			(_nw233).ArraySet1(!(p0), 3)
			(_nw233).ArraySet1(p0, 4)
			(_nw233).ArraySet1((_this).F25(), 5)
			(_nw233).ArraySet1(p0, 6)
			(_nw233).ArraySet1(p0, 7)
			(_nw233).ArraySet1((_this).F25(), 8)
			(_nw233).ArraySet1(p0, 9)
			(_nw233).ArraySet1((_this).F25(), 10)
			(_nw233).ArraySet1((_this).F25(), 11)
			(_nw233).ArraySet1((func() bool {
				if (_1353_v3).Contains(_1352_v2) {
					return (_1353_v3).Get(_1352_v2).(bool)
				}
				return false
			})(), 12)
			(_nw233).ArraySet1(p0, 13)
			(_nw233).ArraySet1((_this).F25(), 14)
			(_nw233).ArraySet1((_this).F25(), 15)
			_1354_v4 = _nw233
			var _1355_v5 _dafny.Map
			_ = _1355_v5
			_1355_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28, _1354_v4)
			var _1356_v6 _dafny.Map
			_ = _1356_v6
			_1356_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_1355_v5).Contains(_this.F28) {
					return (_1355_v5).Get(_this.F28).(_dafny.Array)
				}
				return _1354_v4
			})(), p0)
			_1356_v6 = (_1356_v6).Update(_1354_v4, true)
			var _rhs262 bool = (_this).F25()
			_ = _rhs262
			var _rhs263 bool = true
			_ = _rhs263
			var _lhs215 *GlobalState = globalState
			_ = _lhs215
			var _lhs216 *GlobalState = globalState
			_ = _lhs216
			_lhs215.F1 = _rhs262
			_lhs216.F1 = _rhs263
			var _1357_v7 _dafny.Array
			_ = _1357_v7
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw234
			_1357_v7 = _nw234
			var _1358_v8 _dafny.Array
			_ = _1358_v8
			var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw235
			_1358_v8 = _nw235
			var _1359_v9 _dafny.Sequence
			_ = _1359_v9
			_1359_v9 = _dafny.SeqOf(_1358_v8)
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_1357_v7), 0))
			_ = _index243
			(_1357_v7).ArraySet1(_1359_v9, (_index243).Int())
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_1354_v4), 0))
			_ = _index244
			(_1354_v4).ArraySet1((_this).F25(), (_index244).Int())
			var _1360_v10 D0
			_ = _1360_v10
			_1360_v10 = Companion_D0_.Create_DC1_(_1351_i0, _this.F28, (_this).F25(), (_this).F26(), (_this).F25())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_1357_v7), 0))
			_ = _index245
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_1354_v4), 0))
			_ = _index246
			var _rhs264 bool = true
			_ = _rhs264
			var _rhs265 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1359_v9, _dafny.SeqOf(_1358_v8, _1358_v8, _1358_v8))
			_ = _rhs265
			var _rhs266 bool = p0
			_ = _rhs266
			var _rhs267 _dafny.CodePoint = Companion_Default___.Fm4(_1360_v10, (_this).F25(), Companion_Default___.SafeDivisionInt((_this).F26(), (_this).F27()), globalState)
			_ = _rhs267
			var _rhs268 bool = (_1349_v0).Select((Companion_Default___.SafeIndex(((_this).F27()).Minus((_this).F27()), _dafny.IntOfUint32((_1349_v0).Cardinality()))).Uint32()).(bool)
			_ = _rhs268
			var _lhs217 _dafny.Array = _1357_v7
			_ = _lhs217
			var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_1357_v7), 0))
			_ = _lhs218
			var _lhs219 _dafny.Array = _1354_v4
			_ = _lhs219
			var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_1354_v4), 0))
			_ = _lhs220
			var _lhs221 *GlobalState = globalState
			_ = _lhs221
			r1 = _rhs264
			(_lhs217).ArraySet1(_rhs265, (_lhs218).Int())
			(_lhs219).ArraySet1(_rhs266, (_lhs220).Int())
			_1352_v2 = _rhs267
			_lhs221.F1 = _rhs268
			_1354_v4 = _1354_v4
		}
		var _1361_v11 _dafny.CodePoint
		_ = _1361_v11
		_1361_v11 = _dafny.CodePoint('t')
		var _1362_v12 _dafny.Map
		_ = _1362_v12
		_1362_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1361_v11, p0)
		var _1363_v13 _dafny.Map
		_ = _1363_v13
		_1363_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1362_v12, (_this).F26())
		_1363_v13 = _1363_v13
		var _1364_v14 D1
		_ = _1364_v14
		_1364_v14 = Companion_D1_.Create_DC4_((_1349_v0).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2((_this).F26(), (_this).F26(), globalState), _dafny.IntOfUint32((_1349_v0).Cardinality()))).Uint32()).(bool), p0)
		var _1365_v15 _dafny.Sequence
		_ = _1365_v15
		_1365_v15 = _dafny.SeqOf(_1364_v14)
		var _1366_v16 D0
		_ = _1366_v16
		_1366_v16 = Companion_D0_.Create_DC2_()
		var _1367_v17 D8
		_ = _1367_v17
		_1367_v17 = Companion_D8_.Create_DC17_(_1366_v16, p0)
		var _pat_let_tv34 = p0
		_ = _pat_let_tv34
		_1365_v15 = _dafny.SeqOf((func() D1 {
			if (_1367_v17).Dtor_cf29() {
				return _1364_v14
			}
			return func(_pat_let40_0 D1) D1 {
				return func(_1368_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let41_0 bool) D1 {
						return func(_1369_dt__update_hcf7_h0 bool) D1 {
							return Companion_D1_.Create_DC4_(_1369_dt__update_hcf7_h0, (_1368_dt__update__tmp_h0).Dtor_cf8())
						}(_pat_let41_0)
					}(_pat_let_tv34)
				}(_pat_let40_0)
			}(_1364_v14)
		})())
		var _1370_v18 _dafny.MultiSet
		_ = _1370_v18
		_1370_v18 = _dafny.MultiSetOf((_this).F26(), _dafny.IntOfUint32((_this.F28).Cardinality()))
		var _1371_v19 _dafny.Set
		_ = _1371_v19
		_1371_v19 = _dafny.SetOf(p0)
		var _1372_v20 _dafny.Map
		_ = _1372_v20
		_1372_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1370_v18, (_1371_v19).IsDisjointFrom(_1371_v19))
		_1372_v20 = func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter29 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F27())), _1370_v18)).Elements()); ; {
				_compr_29, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _1373_v21 _dafny.MultiSet
				_1373_v21 = interface{}(_compr_29).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F27())), _1370_v18), _1373_v21) {
					_coll29.Add(_1373_v21, p0)
				}
			}
			return _coll29.ToMap()
		}()
		r0 = (_this).F26()
		r1 = p0
		r0 = (_dafny.Zero).Minus(Companion_Default___.Fm2((_this).F27(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if p0 {
				return _this.F28
			}
			return _this.F28
		})()).Cardinality()), globalState))
		r1 = (_this).F25()
		var _1374_v22 _dafny.MultiSet
		_ = _1374_v22
		_1374_v22 = _dafny.MultiSetOf((_this).F25(), (_this).F25(), (_this).F25())
		var _1375_v23 _dafny.Map
		_ = _1375_v23
		_1375_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v22, (_this).F27())
		r2 = (_1375_v23).Merge((func() _dafny.Map {
			if (_this).F25() {
				return _1375_v23
			}
			return _1375_v23
		})())
		var _1376_v24 _dafny.Map
		_ = _1376_v24
		_1376_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).Plus((_this).F26()), (func() _dafny.Sequence {
			if true {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_1377_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))
			}
			return _this.F28
		})())
		r3 = _1376_v24
		return r0, r1, r2, r3
	}
}
func (_this *C4) M3(globalState *GlobalState) (_dafny.Array, bool, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		if (_this).F25() {
			r2 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(398))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_1378_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality())
			var _1379_v0 _dafny.Array
			_ = _1379_v0
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_42
			var _nw236 _dafny.Array
			_ = _nw236
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw236 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Int = func(_1380_i1 _dafny.Int) _dafny.Int {
					return (_1380_i1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-692))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg63 _dafny.Int) interface{} {
							return coer63(arg63)
						}
					}(func(_1381_i2 _dafny.Int) _dafny.Int {
						return (_this).F27()
					}))), _dafny.MultiSetOf(_dafny.IntOfInt64(659)), _dafny.MultiSetOf(_dafny.IntOfInt64(-667), (_this).F27()), _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F27())))).Cardinality()))
				}
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw236 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw236).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw236).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1379_v0 = _nw236
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
			_ = _index247
			(_1379_v0).ArraySet1((_dafny.Zero).Minus((_this).F27()), (_index247).Int())
			var _1382_v1 D12
			_ = _1382_v1
			_1382_v1 = Companion_D12_.Create_DC28_(_1379_v0)
			var _1383_v2 _dafny.Array
			_ = _1383_v2
			var _nw237 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw237
			_1383_v2 = _nw237
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))
			_ = _index248
			(_1383_v2).ArraySet1(true, (_index248).Int())
			var _1384_v3 _dafny.Set
			_ = _1384_v3
			_1384_v3 = _dafny.SetOf((_this).F25(), (_this).F25(), (_this).F25())
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
			_ = _index249
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))
			_ = _index250
			var _rhs269 _dafny.Int = (func() _dafny.Int {
				if (_this).F25() {
					return (_this).F26()
				}
				return _dafny.IntOfInt64(229)
			})()
			_ = _rhs269
			var _rhs270 _dafny.Int = (_this).F27()
			_ = _rhs270
			var _rhs271 _dafny.Sequence = _this.F28
			_ = _rhs271
			var _rhs272 D12 = _1382_v1
			_ = _rhs272
			var _rhs273 bool = (_1384_v3).IsSubsetOf(_dafny.SetOf(!((_this).F25())))
			_ = _rhs273
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			var _lhs223 _dafny.Array = _1379_v0
			_ = _lhs223
			var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
			_ = _lhs224
			var _lhs225 *C4 = _this
			_ = _lhs225
			var _lhs226 _dafny.Array = _1383_v2
			_ = _lhs226
			var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))
			_ = _lhs227
			_lhs222.F0 = _rhs269
			(_lhs223).ArraySet1(_rhs270, (_lhs224).Int())
			_lhs225.F28 = _rhs271
			_1382_v1 = _rhs272
			(_lhs226).ArraySet1(_rhs273, (_lhs227).Int())
			(_this).F28 = _this.F28
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
			_ = _index251
			var _rhs274 _dafny.Int = (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs274
			var _rhs275 _dafny.Int = (_this).Fm6((_this).F25(), (_this).F25(), (_this).F26(), globalState)
			_ = _rhs275
			var _lhs228 _dafny.Array = _1379_v0
			_ = _lhs228
			var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
			_ = _lhs229
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			(_lhs228).ArraySet1(_rhs274, (_lhs229).Int())
			_lhs230.F2 = _rhs275
			if ((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)).Cmp((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)) <= 0 {
				var _1385_v4 _dafny.Map
				_ = _1385_v4
				_1385_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool))
				var _1386_v5 *C1
				_ = _1386_v5
				var _nw238 *C1 = New_C1_()
				_ = _nw238
				_nw238.Ctor__((_this).F25())
				_1386_v5 = _nw238
				var _1387_v6 _dafny.MultiSet
				_ = _1387_v6
				_1387_v6 = _dafny.MultiSetOf(_1386_v5)
				var _1388_v8 D0
				_ = _1388_v8
				_1388_v8 = Companion_D0_.Create_DC0_((_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool))
				var _1389_v9 _dafny.Set
				_ = _1389_v9
				_1389_v9 = _dafny.SetOf(_1388_v8)
				var _1390_v10 _dafny.Map
				_ = _1390_v10
				_1390_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (func() bool {
					if (_1385_v4).Contains((func() _dafny.Int {
						if (_1387_v6).Contains(_1386_v5) {
							return (_1387_v6).Multiplicity(_1386_v5)
						}
						return (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)
					})()) {
						return (_1385_v4).Get((func() _dafny.Int {
							if (_1387_v6).Contains(_1386_v5) {
								return (_1387_v6).Multiplicity(_1386_v5)
							}
							return (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)
						})()).(bool)
					}
					return Companion_Default___.Fm5((Companion_D1_.Create_DC4_((_this).F25(), (_this).F25())).Dtor_cf8(), func() _dafny.Map {
						var _coll30 = _dafny.NewMapBuilder()
						_ = _coll30
						for _iter30 := _dafny.Iterate((_1389_v9).Elements()); ; {
							_compr_30, _ok30 := _iter30()
							if !_ok30 {
								break
							}
							var _1391_v7 D0
							_1391_v7 = interface{}(_compr_30).(D0)
							if (_1389_v9).Contains(_1391_v7) {
								_coll30.Add(_1391_v7, (_this).F26())
							}
						}
						return _coll30.ToMap()
					}(), (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int), globalState)
				})())
				var _1392_v11 _dafny.Sequence
				_ = _1392_v11
				_1392_v11 = _dafny.SeqOf((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int), (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int))
				var _1393_v12 _dafny.Sequence
				_ = _1393_v12
				_1393_v12 = _dafny.SeqOf((_1386_v5).F32())
				_1390_v10 = (_1390_v10).Update(((_1390_v10).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1392_v11, (Companion_Default___.SafeIndex((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1392_v11).Cardinality()))).Uint32(), (_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int))).Cardinality())), (func() bool {
					if (_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool) {
						return (_1393_v12).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1393_v12).Cardinality()))).Uint32()).(bool)
					}
					return (_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool)
				})())
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))
				_ = _index252
				(_1383_v2).ArraySet1(((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)).Cmp((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int)) != 0, (_index252).Int())
				var _1394_v13 _dafny.Map
				_ = _1394_v13
				_1394_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1379_v0, (_this).F25())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
				_ = _index253
				(_1379_v0).ArraySet1((_1394_v13).Cardinality(), (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))
				_ = _index254
				(_1383_v2).ArraySet1(!((_this).F25()), (_index254).Int())
				_1383_v2 = _1383_v2
			} else {
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))
				_ = _index255
				(_1379_v0).ArraySet1((_dafny.Zero).Minus((_this).F26()), (_index255).Int())
				var _1395_v14 _dafny.Map
				_ = _1395_v14
				_1395_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F27())
				var _1396_v15 _dafny.Map
				_ = _1396_v15
				_1396_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1395_v14).Contains((_this).F25()) {
						return (_1395_v14).Get((_this).F25()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_this.F28).Cardinality())
				})(), _this.F28)
				_1396_v15 = (_1396_v15).Update((_1379_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1379_v0), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28))
				var _1397_v16 _dafny.CodePoint
				_ = _1397_v16
				_1397_v16 = _dafny.CodePoint('u')
				var _1398_v17 _dafny.Map
				_ = _1398_v17
				_1398_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (func() _dafny.CodePoint {
					if (_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool) {
						return (_this.F28).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _1397_v16
				})())
				var _1399_v18 _dafny.Sequence
				_ = _1399_v18
				_1399_v18 = _dafny.SeqOf(false)
				_1398_v17 = (_1398_v17).Update((_1383_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1383_v2), 0))).Int()).(bool), (func() _dafny.CodePoint {
					if (_1399_v18).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1399_v18).Cardinality()))).Uint32()).(bool) {
						return _1397_v16
					}
					return _1397_v16
				})())
				var _1400_v19 *C1
				_ = _1400_v19
				var _nw239 *C1 = New_C1_()
				_ = _nw239
				_nw239.Ctor__(false)
				_1400_v19 = _nw239
				var _1401_v20 _dafny.Sequence
				_ = _1401_v20
				_1401_v20 = _dafny.SeqOf((_this).F26(), (_this).F26())
				var _1402_v21 D11
				_ = _1402_v21
				_1402_v21 = Companion_D11_.Create_DC26_((_this).F26(), _dafny.IntOfInt64(797), (_this).F25())
				_1401_v20 = Companion_Default___.Fm3((_1402_v21).Dtor_cf41(), !(true), (_this).F26(), (_this).F25(), globalState)
			}
		} else {
			var _1403_v22 D0
			_ = _1403_v22
			_1403_v22 = Companion_D0_.Create_DC0_((_this).F25())
			var _1404_v23 _dafny.Map
			_ = _1404_v23
			_1404_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1403_v22, (_this).F26()), (_this).F27(), globalState), (_this).F25())
			_1404_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), !((_this).F25()))
			(globalState).F0 = (_this).F27()
			var _1405_v24 _dafny.CodePoint
			_ = _1405_v24
			_1405_v24 = _dafny.CodePoint('l')
			var _1406_v25 _dafny.Map
			_ = _1406_v25
			_1406_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v24, (_this).F27())
			r1 = ((_this).F26()).Cmp(((_this).F27()).Minus((func() _dafny.Int {
				if (_1406_v25).Contains(_dafny.CodePoint('f')) {
					return (_1406_v25).Get(_dafny.CodePoint('f')).(_dafny.Int)
				}
				return (_this).F27()
			})())) >= 0
			var _1407_v26 _dafny.Array
			_ = _1407_v26
			var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw240
			_1407_v26 = _nw240
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_1407_v26), 0))
			_ = _index256
			(_1407_v26).ArraySet1((_this).F26(), (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_1407_v26), 0))
			_ = _index257
			(_1407_v26).ArraySet1((_this).F26(), (_index257).Int())
			if (_this).F25() {
				var _rhs276 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F28, _dafny.UnicodeSeqOfUtf8Bytes("y"))
				_ = _rhs276
				var _rhs277 _dafny.Int = (_this).F26()
				_ = _rhs277
				var _lhs231 *C4 = _this
				_ = _lhs231
				var _lhs232 *GlobalState = globalState
				_ = _lhs232
				_lhs231.F28 = _rhs276
				_lhs232.F0 = _rhs277
				var _1408_v27 *C2
				_ = _1408_v27
				var _nw241 *C2 = New_C2_()
				_ = _nw241
				_nw241.Ctor__((_this).F27(), false, (_this).F26())
				_1408_v27 = _nw241
				var _1409_v28 *C1
				_ = _1409_v28
				var _nw242 *C1 = New_C1_()
				_ = _nw242
				_nw242.Ctor__(!(false))
				_1409_v28 = _nw242
				var _1410_v29 _dafny.Array
				_ = _1410_v29
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_43
				var _nw243 _dafny.Array
				_ = _nw243
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw243 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Sequence = func(_1411_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28)
					}
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw243 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw243).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw243).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1410_v29 = _nw243
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1410_v29), 0))
				_ = _index258
				(_1410_v29).ArraySet1(_this.F28, (_index258).Int())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1410_v29), 0))
				_ = _index259
				(_1410_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rgnv"), _dafny.UnicodeSeqOfUtf8Bytes("cqotdmdy")), _dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28)), (_index259).Int())
				var _1412_v30 _dafny.Sequence
				_ = _1412_v30
				_1412_v30 = _dafny.SeqOf((_this).F25())
				var _1413_v31 *C3
				_ = _1413_v31
				var _nw244 *C3 = New_C3_()
				_ = _nw244
				_nw244.Ctor__(_this.F28, (_this).F25(), (_this).F25(), _dafny.IntOfUint32((_1412_v30).Cardinality()))
				_1413_v31 = _nw244
				var _1414_v32 _dafny.Map
				_ = _1414_v32
				_1414_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1413_v31, true)
				var _1415_v33 _dafny.Sequence
				_ = _1415_v33
				_1415_v33 = _dafny.SeqOf(_1413_v31)
				var _1416_v34 _dafny.MultiSet
				_ = _1416_v34
				_1416_v34 = _dafny.MultiSetOf((_this).F25(), true, false)
				var _1417_v35 _dafny.Map
				_ = _1417_v35
				_1417_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1410_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1410_v29), 0))).Int()).(_dafny.Sequence)).Cardinality()), _1416_v34)
				_1414_v32 = (_1414_v32).Update((_1415_v33).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((func() _dafny.MultiSet {
					if (_1417_v35).Contains((_this).F26()) {
						return (_1417_v35).Get((_this).F26()).(_dafny.MultiSet)
					}
					return _1416_v34
				})()).Cardinality()), _dafny.IntOfUint32((_1415_v33).Cardinality()))).Uint32()).(*C3), (_1409_v28).F32())
			} else {
				var _1418_v36 D1
				_ = _1418_v36
				_1418_v36 = Companion_D1_.Create_DC4_((_this).F25(), (_this).F25())
				(globalState).F1 = (_1418_v36).Dtor_cf7()
				(globalState).F1 = (_this).F25()
				(globalState).F2 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_this.F28).Cardinality()), _dafny.IntOfInt64(-450))
				var _1419_v37 _dafny.Sequence
				_ = _1419_v37
				_1419_v37 = _dafny.SeqOf((_this).F26())
				_1419_v37 = _1419_v37
				var _1420_v38 _dafny.Set
				_ = _1420_v38
				_1420_v38 = _dafny.SetOf((_this).F25())
				(globalState).F1 = ((_dafny.SetOf((_this).F25())).Difference(_1420_v38)).IsSubsetOf(_1420_v38)
			}
		}
		var _1421_v39 _dafny.CodePoint
		_ = _1421_v39
		_1421_v39 = _dafny.CodePoint('o')
		var _1422_v40 D1
		_ = _1422_v40
		_1422_v40 = Companion_D1_.Create_DC3_(_1421_v39)
		if func(_source20 D1) bool {
			if _source20.Is_DC4() {
				var _1423___mcc_h0 bool = _source20.Get_().(D1_DC4).Cf7
				_ = _1423___mcc_h0
				var _1424___mcc_h1 bool = _source20.Get_().(D1_DC4).Cf8
				_ = _1424___mcc_h1
				var _1425_cf8 bool = _1424___mcc_h1
				_ = _1425_cf8
				var _1426_cf7 bool = _1423___mcc_h0
				_ = _1426_cf7
				return (func() _dafny.Set {
					var _coll31 = _dafny.NewBuilder()
					_ = _coll31
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-349), _dafny.IntOfInt64(15))); ; {
						_compr_31, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _1427_v41 _dafny.Int
						_1427_v41 = interface{}(_compr_31).(_dafny.Int)
						if ((_dafny.IntOfInt64(-349)).Cmp(_1427_v41) <= 0) && ((_1427_v41).Cmp(_dafny.IntOfInt64(15)) < 0) {
							_coll31.Add((_1427_v41).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F27(), (_this).F27())).Cardinality())))
						}
					}
					return _coll31.ToSet()
				}()).IsProperSubsetOf(_dafny.SetOf((_this).F27()))
			} else if _source20.Is_DC3() {
				var _1428___mcc_h2 _dafny.CodePoint = _source20.Get_().(D1_DC3).Cf6
				_ = _1428___mcc_h2
				var _1429_cf6 _dafny.CodePoint = _1428___mcc_h2
				_ = _1429_cf6
				return (_this).F25()
			} else {
				var _1430___mcc_h3 D1 = _source20.Get_().(D1_DC5).Cf9
				_ = _1430___mcc_h3
				var _1431_cf9 D1 = _1430___mcc_h3
				_ = _1431_cf9
				return (_this).F25()
			}
		}(_1422_v40) {
			_1421_v39 = _1421_v39
			(globalState).F0 = (_this).F26()
			var _1432_v42 _dafny.Array
			_ = _1432_v42
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_44
			var _nw245 _dafny.Array
			_ = _nw245
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw245 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) _dafny.Map = func(_1433_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
				}
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw245 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw245).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw245).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_1432_v42 = _nw245
			var _1434_v43 D8
			_ = _1434_v43
			_1434_v43 = Companion_D8_.Create_DC16_(_1432_v42)
			var _1435_v44 _dafny.Array
			_ = _1435_v44
			var _nwElement0_48 D8 = Companion_D8_.Create_DC16_(_1432_v42)
			_ = _nwElement0_48
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(7))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_48, 0)
			(_nw246).ArraySet1(_1434_v43, 1)
			(_nw246).ArraySet1(_1434_v43, 2)
			(_nw246).ArraySet1(_1434_v43, 3)
			(_nw246).ArraySet1(_1434_v43, 4)
			(_nw246).ArraySet1(_1434_v43, 5)
			(_nw246).ArraySet1(_1434_v43, 6)
			_1435_v44 = _nw246
			var _1436_v45 _dafny.Map
			_ = _1436_v45
			_1436_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _1435_v44)
			var _1437_v46 _dafny.Sequence
			_ = _1437_v46
			_1437_v46 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _1435_v44))
			_1436_v45 = (_1437_v46).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1437_v46).Cardinality()))).Uint32()).(_dafny.Map)
			var _1438_v47 _dafny.Sequence
			_ = _1438_v47
			_1438_v47 = _dafny.SeqOf((_this).F27(), (_this).F27())
			(globalState).F2 = (_dafny.IntOfInt64(703)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1438_v47, (_this).F25())).Cardinality())
			var _1439_v48 *C3
			_ = _1439_v48
			var _nw247 *C3 = New_C3_()
			_ = _nw247
			_nw247.Ctor__(_this.F28, (_this).F25(), (_this).F25(), _dafny.IntOfInt64(-752))
			_1439_v48 = _nw247
		} else {
			var _1440_v49 _dafny.Set
			_ = _1440_v49
			_1440_v49 = _dafny.SetOf((_this).F25())
			if ((_1440_v49).Union(_1440_v49)).IsDisjointFrom(_dafny.SetOf((_this).F25(), (_this).F25())) {
				var _1441_v50 _dafny.Array
				_ = _1441_v50
				var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw248
				_1441_v50 = _nw248
				_1441_v50 = _1441_v50
				var _1442_v51 _dafny.Map
				_ = _1442_v51
				_1442_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
				_1442_v51 = (_1442_v51).Update((_this).F25(), !((_this).F25()))
				var _1443_v52 _dafny.Array
				_ = _1443_v52
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_45
				var _nw249 _dafny.Array
				_ = _nw249
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw249 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) bool = func(_1444_i5 _dafny.Int) bool {
						return false
					}
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw249 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw249).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw249).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1443_v52 = _nw249
				var _1445_v53 *C0
				_ = _1445_v53
				var _nw250 *C0 = New_C0_()
				_ = _nw250
				_nw250.Ctor__(_this.F28, (_this).F25(), (_this).F26(), _1443_v52, (_this).F25())
				_1445_v53 = _nw250
				var _1446_v54 *C3
				_ = _1446_v54
				var _nw251 *C3 = New_C3_()
				_ = _nw251
				_nw251.Ctor__(_this.F28, (_this).F25(), (_this).F25(), (_this).F26())
				_1446_v54 = _nw251
				var _1447_v55 _dafny.Array
				_ = _1447_v55
				var _nwElement0_49 *C3 = _1446_v54
				_ = _nwElement0_49
				var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(12))
				_ = _nw252
				(_nw252).ArraySet1(_nwElement0_49, 0)
				(_nw252).ArraySet1(_1446_v54, 1)
				(_nw252).ArraySet1(_1446_v54, 2)
				(_nw252).ArraySet1(_1446_v54, 3)
				(_nw252).ArraySet1(_1446_v54, 4)
				(_nw252).ArraySet1(_1446_v54, 5)
				(_nw252).ArraySet1(_1446_v54, 6)
				(_nw252).ArraySet1(_1446_v54, 7)
				(_nw252).ArraySet1(_1446_v54, 8)
				(_nw252).ArraySet1((func() *C3 {
					if false {
						return _1446_v54
					}
					return _1446_v54
				})(), 9)
				(_nw252).ArraySet1(_1446_v54, 10)
				(_nw252).ArraySet1(_1446_v54, 11)
				_1447_v55 = _nw252
				var _1448_v56 _dafny.Sequence
				_ = _1448_v56
				_1448_v56 = _dafny.SeqOf(_1443_v52)
				var _1449_v57 D0
				_ = _1449_v57
				_1449_v57 = Companion_D0_.Create_DC0_((_1446_v54).F30())
				var _1450_v58 _dafny.Map
				_ = _1450_v58
				_1450_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1449_v57, (_dafny.Zero).Minus((_this).F26()))
				var _1451_v59 _dafny.MultiSet
				_ = _1451_v59
				_1451_v59 = _dafny.MultiSetOf((_this).F25())
				var _1452_v60 _dafny.Map
				_ = _1452_v60
				_1452_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1451_v59).Cardinality(), (_this).F26())
				var _rhs278 _dafny.Int = (_this).F26()
				_ = _rhs278
				var _rhs279 _dafny.Array = (func() _dafny.Array {
					if _dafny.Companion_Sequence_.Contains(_1448_v56, _1443_v52) {
						return _1447_v55
					}
					return _1447_v55
				})()
				_ = _rhs279
				var _rhs280 bool = Companion_Default___.Fm5((_1446_v54).F30(), _1450_v58, (_this).F27(), globalState)
				_ = _rhs280
				var _rhs281 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if !(_1452_v60).Contains((_1446_v54).Fm9(globalState)) {
						return (func() _dafny.Int {
							if (_1446_v54).F30() {
								return (_this).F26()
							}
							return (_this).F26()
						})()
					}
					return Companion_Default___.SafeModuloInt((_this).F27(), (_this).F26())
				})())
				_ = _rhs281
				var _lhs233 *GlobalState = globalState
				_ = _lhs233
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				_lhs233.F0 = _rhs278
				_1447_v55 = _rhs279
				_lhs234.F1 = _rhs280
				_lhs235.F2 = _rhs281
				var _1453_v61 _dafny.Array
				_ = _1453_v61
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_46
				var _nw253 _dafny.Array
				_ = _nw253
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw253 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Sequence = (func(_1454_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1455_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg64 _dafny.Int) interface{} {
									return coer64(arg64)
								}
							}((func(_1456_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1457_i7 _dafny.Int) _dafny.CodePoint {
									return _1456_v39
								}
							})(_1454_v39)))
						}
					})(_1421_v39)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw253 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw253).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw253).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1453_v61 = _nw253
				_1453_v61 = _1453_v61
			} else {
				var _1458_v62 _dafny.MultiSet
				_ = _1458_v62
				_1458_v62 = _dafny.MultiSetOf((_this).F27())
				var _1459_v63 *C2
				_ = _1459_v63
				var _nw254 *C2 = New_C2_()
				_ = _nw254
				_nw254.Ctor__(((func() _dafny.Int {
					if (_1458_v62).Contains(_dafny.IntOfInt64(824)) {
						return (_1458_v62).Multiplicity(_dafny.IntOfInt64(824))
					}
					return (_this).F26()
				})()).Plus((_this).F27()), (_this).F25(), (_this).F27())
				_1459_v63 = _nw254
				_1459_v63 = _1459_v63
				(globalState).F0 = (Companion_Default___.SafeDivisionInt((_1459_v63).F31(), (_this).F27())).Times((_1459_v63).F31())
				var _1460_v64 _dafny.Array
				_ = _1460_v64
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw255
				_1460_v64 = _nw255
				var _1461_v65 _dafny.Array
				_ = _1461_v65
				var _nwElement0_50 _dafny.Int = (_this).F27()
				_ = _nwElement0_50
				var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(9))
				_ = _nw256
				(_nw256).ArraySet1(_nwElement0_50, 0)
				(_nw256).ArraySet1((_this).F26(), 1)
				(_nw256).ArraySet1((_1459_v63).F31(), 2)
				(_nw256).ArraySet1((_1459_v63).F31(), 3)
				(_nw256).ArraySet1((_1459_v63).F31(), 4)
				(_nw256).ArraySet1((_this).F27(), 5)
				(_nw256).ArraySet1((_this).F27(), 6)
				(_nw256).ArraySet1((_1459_v63).F31(), 7)
				(_nw256).ArraySet1(_dafny.IntOfInt64(514), 8)
				_1461_v65 = _nw256
				var _1462_v66 _dafny.Map
				_ = _1462_v66
				_1462_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1461_v65, (_this).F25())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1460_v64), 0))
				_ = _index260
				(_1460_v64).ArraySet1(_1462_v66, (_index260).Int())
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1460_v64), 0))
				_ = _index261
				(_1460_v64).ArraySet1(_1462_v66, (_index261).Int())
				(globalState).F0 = (func() _dafny.Int {
					if (_1458_v62).Contains(_dafny.IntOfInt64(271)) {
						return (_1458_v62).Multiplicity(_dafny.IntOfInt64(271))
					}
					return (_this).F27()
				})()
				var _1463_v67 _dafny.MultiSet
				_ = _1463_v67
				_1463_v67 = _dafny.MultiSetOf(_this.F28, _this.F28)
				_1463_v67 = (_1463_v67).Difference((_1463_v67).Intersection(_1463_v67))
			}
			var _1464_v68 _dafny.Map
			_ = _1464_v68
			_1464_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F25())
			_1464_v68 = (_1464_v68).Update(_dafny.IntOfUint32((_this.F28).Cardinality()), !((_this).F25()))
			var _1465_v69 _dafny.Sequence
			_ = _1465_v69
			_1465_v69 = _dafny.SeqOf(_this.F28)
			var _1466_v70 *C3
			_ = _1466_v70
			var _nw257 *C3 = New_C3_()
			_ = _nw257
			_nw257.Ctor__(_dafny.Companion_Sequence_.Concatenate(_this.F28, (_1465_v69).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1465_v69).Cardinality()))).Uint32()).(_dafny.Sequence)), (_this).F25(), (_this).F25(), (_this).F27())
			_1466_v70 = _nw257
			var _1467_v71 _dafny.Array
			_ = _1467_v71
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw258
			_1467_v71 = _nw258
			var _1468_v72 _dafny.Sequence
			_ = _1468_v72
			_1468_v72 = _dafny.SeqOf((_1466_v70).F30())
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1467_v71), 0))
			_ = _index262
			(_1467_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1468_v72, _1468_v72), (_index262).Int())
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1467_v71), 0))
			_ = _index263
			(_1467_v71).ArraySet1(_1468_v72, (_index263).Int())
			var _1469_v73 _dafny.MultiSet
			_ = _1469_v73
			_1469_v73 = _dafny.MultiSetOf(_dafny.IntOfUint32((_this.F28).Cardinality()), (_this).F26(), (_this).Fm6((_this).F25(), (_1466_v70).F30(), (_dafny.Zero).Minus((_this).F27()), globalState))
			(globalState).F0 = (func() _dafny.Int {
				if (_1469_v73).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1440_v49, (_this).Fm7(globalState))).Cardinality()) {
					return (_1469_v73).Multiplicity((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1440_v49, (_this).Fm7(globalState))).Cardinality())
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F28, _this.F28)).Cardinality())
			})()
		}
		var _hi13 _dafny.Int = _dafny.IntOfInt64(801)
		_ = _hi13
		for _1470_i8 := ((_this).F26()).Plus((_this).F27()); _1470_i8.Cmp(_hi13) < 0; _1470_i8 = _1470_i8.Plus(_dafny.One) {
			var _1471_v74 _dafny.Array
			_ = _1471_v74
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_47
			var _nw259 _dafny.Array
			_ = _nw259
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw259 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) bool = func(_1472_i9 _dafny.Int) bool {
					return (_this).F25()
				}
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw259 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw259).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw259).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1471_v74 = _nw259
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))
			_ = _index264
			(_1471_v74).ArraySet1((_this).F25(), (_index264).Int())
			var _1473_v75 _dafny.Array
			_ = _1473_v75
			var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(21))
			_ = _nw260
			_1473_v75 = _nw260
			var _1474_v76 _dafny.MultiSet
			_ = _1474_v76
			_1474_v76 = _dafny.MultiSetOf(_1471_v74, _1471_v74, _1471_v74)
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1473_v75), 0))
			_ = _index265
			(_1473_v75).ArraySet1(_1474_v76, (_index265).Int())
			var _1475_v77 _dafny.Array
			_ = _1475_v77
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw261
			_1475_v77 = _nw261
			var _1476_v78 _dafny.Array
			_ = _1476_v78
			var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
			_ = _nw262
			_1476_v78 = _nw262
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1475_v77), 0))
			_ = _index266
			(_1475_v77).ArraySet1(_1476_v78, (_index266).Int())
			var _1477_v79 _dafny.Sequence
			_ = _1477_v79
			_1477_v79 = _dafny.SeqOf((_this).F25(), (func() bool {
				if (_this).F25() {
					return !((_this).F25())
				}
				return (_this).F25()
			})(), ((_this).F25()) || ((_this).F25()))
			var _1478_v80 _dafny.Map
			_ = _1478_v80
			_1478_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1476_v78)
			var _1479_v81 _dafny.Set
			_ = _1479_v81
			_1479_v81 = _dafny.SetOf((_this).F25())
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))
			_ = _index267
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1473_v75), 0))
			_ = _index268
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1475_v77), 0))
			_ = _index269
			var _rhs282 bool = !((_this).F25())
			_ = _rhs282
			var _rhs283 bool = (_1477_v79).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1477_v79).Cardinality()))).Uint32()).(bool)
			_ = _rhs283
			var _rhs284 _dafny.MultiSet = _1474_v76
			_ = _rhs284
			var _rhs285 _dafny.Array = (func() _dafny.Array {
				if (_1478_v80).Contains((Companion_Default___.Fm36((_this).F27(), _dafny.IntOfUint32((_this.F28).Cardinality()), _this.F28, (_this).F27(), globalState)).Cardinality()) {
					return (_1478_v80).Get((Companion_Default___.Fm36((_this).F27(), _dafny.IntOfUint32((_this.F28).Cardinality()), _this.F28, (_this).F27(), globalState)).Cardinality()).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if (_this).F25() {
						return _1476_v78
					}
					return _1476_v78
				})()
			})()
			_ = _rhs285
			var _rhs286 _dafny.Int = (_1479_v81).Cardinality()
			_ = _rhs286
			var _lhs236 _dafny.Array = _1471_v74
			_ = _lhs236
			var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))
			_ = _lhs237
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			var _lhs239 _dafny.Array = _1473_v75
			_ = _lhs239
			var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1473_v75), 0))
			_ = _lhs240
			var _lhs241 _dafny.Array = _1475_v77
			_ = _lhs241
			var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1475_v77), 0))
			_ = _lhs242
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			(_lhs236).ArraySet1(_rhs282, (_lhs237).Int())
			_lhs238.F1 = _rhs283
			(_lhs239).ArraySet1(_rhs284, (_lhs240).Int())
			(_lhs241).ArraySet1(_rhs285, (_lhs242).Int())
			_lhs243.F0 = _rhs286
			(globalState).F2 = Companion_Default___.Fm2((_dafny.Zero).Minus(Companion_Default___.Fm2((_this).F26(), (_this).F27(), globalState)), (_this).F27(), globalState)
			var _1480_v82 _dafny.Array
			_ = _1480_v82
			var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
			_ = _nw263
			_1480_v82 = _nw263
			var _1481_v83 D0
			_ = _1481_v83
			_1481_v83 = Companion_D0_.Create_DC0_((_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool))
			var _1482_v84 _dafny.Map
			_ = _1482_v84
			_1482_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), (_this).F25())
			var _1483_v85 _dafny.Map
			_ = _1483_v85
			_1483_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1481_v83, ((_1482_v84).Update(false, (_this).F25())).Cardinality())
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1480_v82), 0))
			_ = _index270
			(_1480_v82).ArraySet1((_1479_v81).Intersection(_dafny.SetOf((_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), (_this).F25(), (_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), (_this).F25(), Companion_Default___.Fm5((_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), _1483_v85, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-270))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}(func(_1484_i10 _dafny.Int) _dafny.Int {
				return (_this).F27()
			}))).Cardinality()), globalState))), (_index270).Int())
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1480_v82), 0))
			_ = _index271
			(_1480_v82).ArraySet1(_1479_v81, (_index271).Int())
			var _1485_v86 _dafny.Map
			_ = _1485_v86
			_1485_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(605), (_this).F27())
			var _1486_v87 _dafny.MultiSet
			_ = _1486_v87
			_1486_v87 = _dafny.MultiSetOf((_this).F26())
			var _1487_v88 _dafny.Map
			_ = _1487_v88
			_1487_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1480_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1480_v82), 0))).Int()).(_dafny.Set)).Cardinality(), _1486_v87)
			var _1488_v89 *C2
			_ = _1488_v89
			var _nw264 *C2 = New_C2_()
			_ = _nw264
			_nw264.Ctor__((_this).F27(), (_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), (_dafny.Zero).Minus((_1487_v88).Cardinality()))
			_1488_v89 = _nw264
			var _1489_v90 _dafny.Set
			_ = _1489_v90
			_1489_v90 = _dafny.SetOf(_1488_v89)
			var _1490_v91 _dafny.Set
			_ = _1490_v91
			_1490_v91 = _dafny.SetOf((_this).F26(), _1470_i8, (_1489_v90).Cardinality())
			var _1491_v92 _dafny.Sequence
			_ = _1491_v92
			_1491_v92 = _dafny.SeqOf((_1490_v91).Cardinality(), _1470_i8, (_this).F26(), (_1488_v89).F31(), (_1488_v89).F31())
			(globalState).F2 = Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3((_this).F27(), (_1471_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1471_v74), 0))).Int()).(bool), (_1485_v86).Cardinality(), !(!((_this).F25())), globalState), _1491_v92)).Cardinality()), (_this).F27(), globalState)
		}
		if !(!((_this).F25()) || (!((_this).F25()))) || ((func() bool {
			if (_this).F25() {
				return false
			}
			return (_this).F25()
		})()) {
			(globalState).F1 = (_this).F25()
			var _1492_v93 _dafny.Sequence
			_ = _1492_v93
			_1492_v93 = _dafny.SeqOf(false)
			var _1493_v94 _dafny.Map
			_ = _1493_v94
			_1493_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1492_v93).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1492_v93).Cardinality()))).Uint32()).(bool))
			_1493_v94 = (_1493_v94).Update((_this).F25(), (_this).F25())
			var _1494_v95 *C2
			_ = _1494_v95
			var _nw265 *C2 = New_C2_()
			_ = _nw265
			_nw265.Ctor__((_this).F26(), (_this).F25(), (_this).F26())
			_1494_v95 = _nw265
			var _1495_v96 _dafny.Sequence
			_ = _1495_v96
			_1495_v96 = _dafny.SeqOf(_1494_v95, _1494_v95)
			var _1496_v97 _dafny.Map
			_ = _1496_v97
			_1496_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F25()), (_1495_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-67), _dafny.IntOfUint32((_1495_v96).Cardinality()))).Uint32()).(*C2))
			var _1497_v98 _dafny.Sequence
			_ = _1497_v98
			_1497_v98 = _dafny.SeqOf((func() *C2 {
				if (_1496_v97).Contains((_this).F25()) {
					return (_1496_v97).Get((_this).F25()).(*C2)
				}
				return _1494_v95
			})())
			var _1498_v99 _dafny.MultiSet
			_ = _1498_v99
			_1498_v99 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1497_v98, _1495_v96))
			var _1499_v100 _dafny.MultiSet
			_ = _1499_v100
			_1499_v100 = (_1498_v99).Update(_1497_v98, Companion_Default___.Abs((_this).F26()))
			var _1500_v101 _dafny.Sequence
			_ = _1500_v101
			_1500_v101 = _dafny.SeqOf(_1497_v98)
			_1498_v99 = ((_1498_v99).Union(_dafny.MultiSetOf(_1495_v96))).Union((_dafny.MultiSetFromSeq(_1500_v101)))
			var _1501_v102 _dafny.Map
			_ = _1501_v102
			_1501_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1494_v95).F31(), (_this).F25())
			_1501_v102 = (_1501_v102).Update((_1494_v95).F31(), true)
			(globalState).F1 = (_this).F25()
		} else {
			var _1502_v103 _dafny.MultiSet
			_ = _1502_v103
			_1502_v103 = _dafny.MultiSetOf((_this).F26())
			var _1503_v104 _dafny.Map
			_ = _1503_v104
			_1503_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(541), (_1502_v103).Cardinality())
			var _1504_v105 _dafny.Map
			_ = _1504_v105
			_1504_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1503_v104, (_this).F25())
			var _1505_v106 _dafny.Sequence
			_ = _1505_v106
			_1505_v106 = _dafny.SeqOf((_this).F25(), (_this).F25(), (func() bool {
				if (_1504_v105).Contains(_1503_v104) {
					return (_1504_v105).Get(_1503_v104).(bool)
				}
				return !((_this).F25())
			})(), (_this).F25())
			if _dafny.Companion_Sequence_.IsPrefixOf(_1505_v106, _1505_v106) {
				var _1506_v107 _dafny.Array
				_ = _1506_v107
				var _nw266 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw266
				_1506_v107 = _nw266
				var _1507_v108 D12
				_ = _1507_v108
				_1507_v108 = Companion_D12_.Create_DC28_(_1506_v107)
				var _1508_v109 _dafny.Array
				_ = _1508_v109
				var _nwElement0_51 _dafny.Array = _1506_v107
				_ = _nwElement0_51
				var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(27))
				_ = _nw267
				(_nw267).ArraySet1(_nwElement0_51, 0)
				(_nw267).ArraySet1(_1506_v107, 1)
				(_nw267).ArraySet1(_1506_v107, 2)
				(_nw267).ArraySet1(_1506_v107, 3)
				(_nw267).ArraySet1(_1506_v107, 4)
				(_nw267).ArraySet1(_1506_v107, 5)
				(_nw267).ArraySet1(_1506_v107, 6)
				(_nw267).ArraySet1(_1506_v107, 7)
				(_nw267).ArraySet1(_1506_v107, 8)
				(_nw267).ArraySet1((func() _dafny.Array {
					if false {
						return _1506_v107
					}
					return _1506_v107
				})(), 9)
				(_nw267).ArraySet1(_1506_v107, 10)
				(_nw267).ArraySet1(_1506_v107, 11)
				(_nw267).ArraySet1(_1506_v107, 12)
				(_nw267).ArraySet1((_1507_v108).Dtor_cf44(), 13)
				(_nw267).ArraySet1(_1506_v107, 14)
				(_nw267).ArraySet1(_1506_v107, 15)
				(_nw267).ArraySet1((func() _dafny.Array {
					if (_this).F25() {
						return _1506_v107
					}
					return _1506_v107
				})(), 16)
				(_nw267).ArraySet1(_1506_v107, 17)
				(_nw267).ArraySet1(_1506_v107, 18)
				(_nw267).ArraySet1((func() _dafny.Array {
					if (_this).F25() {
						return _1506_v107
					}
					return _1506_v107
				})(), 19)
				(_nw267).ArraySet1(_1506_v107, 20)
				(_nw267).ArraySet1(_1506_v107, 21)
				(_nw267).ArraySet1(_1506_v107, 22)
				(_nw267).ArraySet1(_1506_v107, 23)
				(_nw267).ArraySet1(_1506_v107, 24)
				(_nw267).ArraySet1(_1506_v107, 25)
				(_nw267).ArraySet1(_1506_v107, 26)
				_1508_v109 = _nw267
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1508_v109), 0))
				_ = _index272
				(_1508_v109).ArraySet1(_1506_v107, (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1508_v109), 0))
				_ = _index273
				(_1508_v109).ArraySet1(_1506_v107, (_index273).Int())
				var _1509_v110 _dafny.Array
				_ = _1509_v110
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_48
				var _nw268 _dafny.Array
				_ = _nw268
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw268 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) bool = func(_1510_i11 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw268 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw268).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw268).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1509_v110 = _nw268
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1509_v110), 0))
				_ = _index274
				(_1509_v110).ArraySet1(!((_this).F25()) || (true), (_index274).Int())
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1509_v110), 0))
				_ = _index275
				(_1509_v110).ArraySet1((_this).F25(), (_index275).Int())
				(globalState).F0 = Companion_Default___.SafeModuloInt((_this).F26(), (_this).F27())
				var _1511_v111 *C3
				_ = _1511_v111
				var _nw269 *C3 = New_C3_()
				_ = _nw269
				_nw269.Ctor__(_this.F28, (_this).F25(), false, (_this).F27())
				_1511_v111 = _nw269
				var _1512_v112 _dafny.Map
				_ = _1512_v112
				_1512_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1502_v103).Contains((_this).F27()) {
						return (_1502_v103).Multiplicity((_this).F27())
					}
					return _dafny.IntOfInt64(205)
				})(), _1511_v111)
				_1503_v104 = (_1503_v104).Update((_dafny.IntOfInt64(-314)).Minus((_1512_v112).Cardinality()), (_this).F27())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1509_v110), 0))
				_ = _index276
				(_1509_v110).ArraySet1((_1502_v103).Equals(_1502_v103), (_index276).Int())
			} else {
				var _1513_v113 D1
				_ = _1513_v113
				_1513_v113 = Companion_D1_.Create_DC4_((_this).F25(), (_this).F25())
				var _1514_v114 D1
				_ = _1514_v114
				_1514_v114 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(_1513_v113))
				var _1515_v115 _dafny.Sequence
				_ = _1515_v115
				_1515_v115 = _dafny.SeqOf(_1514_v114, _1514_v114)
				var _1516_v116 _dafny.Set
				_ = _1516_v116
				_1516_v116 = _dafny.SetOf(_1515_v115)
				var _1517_v117 _dafny.Sequence
				_ = _1517_v117
				_1517_v117 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(839))).Uint32(), func(coer66 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_1518_v113 D1) func(_dafny.Int) D1 {
					return func(_1519_i12 _dafny.Int) D1 {
						return Companion_D1_.Create_DC5_(_1518_v113)
					}
				})(_1513_v113))))
				var _1520_v119 _dafny.Map
				_ = _1520_v119
				_1520_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_1516_v116).Difference(func() _dafny.Set {
					var _coll32 = _dafny.NewBuilder()
					_ = _coll32
					for _iter32 := _dafny.Iterate((_1517_v117).Elements()); ; {
						_compr_32, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _1521_v118 _dafny.Sequence
						_1521_v118 = interface{}(_compr_32).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_1517_v117, _1521_v118) {
							_coll32.Add(_1521_v118)
						}
					}
					return _coll32.ToSet()
				}()))
				_1520_v119 = (_1520_v119).Update((_this).F25(), _1516_v116)
				var _1522_v120 _dafny.Map
				_ = _1522_v120
				_1522_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F27())
				var _1523_v121 _dafny.Array
				_ = _1523_v121
				var _nw270 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw270
				_1523_v121 = _nw270
				r0 = (func() _dafny.Array {
					if !(_1522_v120).Equals(_1522_v120) {
						return _1523_v121
					}
					return _1523_v121
				})()
				r1 = _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm1(globalState), _1421_v39)
				(globalState).F0 = ((func() _dafny.Int {
					if (_1503_v104).Contains((_this).F26()) {
						return (_1503_v104).Get((_this).F26()).(_dafny.Int)
					}
					return (_this).F26()
				})()).Minus((_this).F27())
				r1 = (_this).F25()
			}
			var _1524_v122 _dafny.Map
			_ = _1524_v122
			_1524_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F26()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(295))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}(func(_1525_i13 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())).Cardinality()
			}))).Cardinality()), (_this).F25()))
			r2 = Companion_Default___.Fm2((_1524_v122).Cardinality(), (_this).F26(), globalState)
			var _1526_v123 _dafny.Int
			_ = _1526_v123
			var _out33 _dafny.Int
			_ = _out33
			_out33 = Companion_Default___.M0((_1505_v106).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1505_v106).Cardinality()))).Uint32()).(bool), globalState)
			_1526_v123 = _out33
			r2 = (_this).F27()
			(globalState).F1 = (_this).F25()
		}
		(globalState).F2 = (_dafny.Zero).Minus(((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(782))).Cardinality())).Plus((_this).F26())).Plus((_dafny.Zero).Minus((_this).F27())))
		var _1527_v124 D0
		_ = _1527_v124
		_1527_v124 = Companion_D0_.Create_DC0_((_this).F25())
		var _1528_v125 _dafny.Map
		_ = _1528_v125
		_1528_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm28(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}((func(_1529_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1530_i14 _dafny.Int) _dafny.CodePoint {
				return _1529_v39
			}
		})(_1421_v39))), _1421_v39, _dafny.IntOfInt64(-187), globalState), (_this).F26())
		var _1531_v126 _dafny.MultiSet
		_ = _1531_v126
		_1531_v126 = _dafny.MultiSetOf(true, (_this).F25())
		if Companion_Default___.Fm5((_1527_v124).Dtor_cf0(), _1528_v125, (func() _dafny.Int {
			if (_1531_v126).Contains((_this).F25()) {
				return (_1531_v126).Multiplicity((_this).F25())
			}
			return (_this).F27()
		})(), globalState) {
			var _1532_v127 _dafny.Map
			_ = _1532_v127
			_1532_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F27())
			var _1533_v128 _dafny.Map
			_ = _1533_v128
			_1533_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality())).Cmp((_1532_v127).Cardinality()) != 0)
			var _1534_v129 _dafny.MultiSet
			_ = _1534_v129
			_1534_v129 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umwt")).Cardinality()))
			var _1535_v130 D11
			_ = _1535_v130
			_1535_v130 = Companion_D11_.Create_DC26_((_dafny.Zero).Minus((_this).F26()), (_this).F27(), (_this).F25())
			var _1536_v131 _dafny.Set
			_ = _1536_v131
			_1536_v131 = _dafny.SetOf((_this).F27(), _dafny.IntOfInt64(649))
			var _1537_v132 _dafny.MultiSet
			_ = _1537_v132
			_1537_v132 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_1534_v129).Contains((_1535_v130).Dtor_cf41()) {
					return (_1534_v129).Multiplicity((_1535_v130).Dtor_cf41())
				}
				return Companion_Default___.Fm2((_this).F27(), (_1536_v131).Cardinality(), globalState)
			})(), (_this).F26(), (_this).F27())
			_1533_v128 = (_1533_v128).Update((_1537_v132).Cardinality(), (_this).F25())
			var _1538_v133 _dafny.Map
			_ = _1538_v133
			_1538_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v126, _1421_v39)
			var _1539_v134 D14
			_ = _1539_v134
			_1539_v134 = Companion_D14_.Create_DC31_((_1538_v133).Update(_dafny.MultiSetOf((_this).F25(), (_this).F25()), _1421_v39))
			var _1540_v135 _dafny.Array
			_ = _1540_v135
			var _nwElement0_52 _dafny.Map = (_1538_v133).Merge(_1538_v133)
			_ = _nwElement0_52
			var _nw271 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(22))
			_ = _nw271
			(_nw271).ArraySet1(_nwElement0_52, 0)
			(_nw271).ArraySet1(_1538_v133, 1)
			(_nw271).ArraySet1(_1538_v133, 2)
			(_nw271).ArraySet1(((_1539_v134).Dtor_cf51()).Update(_1531_v126, _1421_v39), 3)
			(_nw271).ArraySet1(_1538_v133, 4)
			(_nw271).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm27((_this).F26(), globalState), _dafny.CodePoint('n')), 5)
			(_nw271).ArraySet1(_1538_v133, 6)
			(_nw271).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v126, _1421_v39), 7)
			(_nw271).ArraySet1(_1538_v133, 8)
			(_nw271).ArraySet1(_1538_v133, 9)
			(_nw271).ArraySet1(_1538_v133, 10)
			(_nw271).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v126, _1421_v39), 11)
			(_nw271).ArraySet1(_1538_v133, 12)
			(_nw271).ArraySet1(_1538_v133, 13)
			(_nw271).ArraySet1(_1538_v133, 14)
			(_nw271).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v126, _1421_v39), 15)
			(_nw271).ArraySet1((_1538_v133).Merge(_1538_v133), 16)
			(_nw271).ArraySet1(_1538_v133, 17)
			(_nw271).ArraySet1((_1538_v133).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v126, _1421_v39)), 18)
			(_nw271).ArraySet1(Companion_Default___.Fm37((_this).F25(), globalState), 19)
			(_nw271).ArraySet1(_1538_v133, 20)
			(_nw271).ArraySet1(_1538_v133, 21)
			_1540_v135 = _nw271
			var _1541_v136 _dafny.Sequence
			_ = _1541_v136
			_1541_v136 = _dafny.SeqOf(_1536_v131, _dafny.SetOf((_this).F27()))
			var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1540_v135), 0))
			_ = _index277
			(_1540_v135).ArraySet1(Companion_Default___.Fm37((Companion_D12_.Create_DC29_((_this).F27(), (_this).F25(), _1541_v136, (_dafny.Zero).Minus((_this).F26()), (_1532_v127).Cardinality())).Dtor_cf46(), globalState), (_index277).Int())
			var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1540_v135), 0))
			_ = _index278
			(_1540_v135).ArraySet1(_1538_v133, (_index278).Int())
			var _source21 D11 = Companion_Default___.Fm34(_this.F28, _dafny.IntOfInt64(277), globalState)
			_ = _source21
			if _source21.Is_DC26() {
				var _1542___mcc_h4 _dafny.Int = _source21.Get_().(D11_DC26).Cf40
				_ = _1542___mcc_h4
				var _1543___mcc_h5 _dafny.Int = _source21.Get_().(D11_DC26).Cf41
				_ = _1543___mcc_h5
				var _1544___mcc_h6 bool = _source21.Get_().(D11_DC26).Cf42
				_ = _1544___mcc_h6
				var _1545_cf42 bool = _1544___mcc_h6
				_ = _1545_cf42
				var _1546_cf41 _dafny.Int = _1543___mcc_h5
				_ = _1546_cf41
				var _1547_cf40 _dafny.Int = _1542___mcc_h4
				_ = _1547_cf40
				var _1548_v137 _dafny.Sequence
				_ = _1548_v137
				_1548_v137 = _dafny.SeqOf(_1547_cf40, (_this).F27(), _1546_cf41, (_this).F26())
				_1548_v137 = _1548_v137
				var _1549_v138 _dafny.Array
				_ = _1549_v138
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_49
				var _nw272 _dafny.Array
				_ = _nw272
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw272 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) bool = func(_1550_i15 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw272 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw272).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw272).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1549_v138 = _nw272
				var _1551_v139 *C0
				_ = _1551_v139
				var _nw273 *C0 = New_C0_()
				_ = _nw273
				_nw273.Ctor__(_this.F28, _1545_cf42, (_this).F27(), _1549_v138, _1545_cf42)
				_1551_v139 = _nw273
				var _1552_v140 _dafny.Array
				_ = _1552_v140
				var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
				_ = _nw274
				_1552_v140 = _nw274
				_1552_v140 = _1552_v140
				var _1553_v141 _dafny.Map
				_ = _1553_v141
				_1553_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1545_cf42)
				var _1554_v142 D14
				_ = _1554_v142
				_1554_v142 = Companion_D14_.Create_DC32_((_this).F25(), (_this).F26(), _1553_v141, _dafny.SeqOf(_1421_v39), (_this).F25())
				var _1555_v143 *C1
				_ = _1555_v143
				var _nw275 *C1 = New_C1_()
				_ = _nw275
				_nw275.Ctor__((_this).F25())
				_1555_v143 = _nw275
				var _1556_v144 D10
				_ = _1556_v144
				_1556_v144 = Companion_D10_.Create_DC24_((_1554_v142).Dtor_cf53(), _1555_v143, _1421_v39)
				(globalState).F2 = (_1556_v144).Dtor_cf36()
			} else if _source21.Is_DC25() {
				var _1557___mcc_h7 _dafny.Map = _source21.Get_().(D11_DC25).Cf39
				_ = _1557___mcc_h7
				var _1558_cf39 _dafny.Map = _1557___mcc_h7
				_ = _1558_cf39
				var _1559_v145 _dafny.Array
				_ = _1559_v145
				var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw276
				_1559_v145 = _nw276
				var _rhs287 _dafny.Int = (_this).F27()
				_ = _rhs287
				var _rhs288 _dafny.Array = _1559_v145
				_ = _rhs288
				var _rhs289 bool = (_this).F25()
				_ = _rhs289
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				var _lhs245 *GlobalState = globalState
				_ = _lhs245
				_lhs244.F2 = _rhs287
				_1559_v145 = _rhs288
				_lhs245.F1 = _rhs289
				var _1560_v146 _dafny.Array
				_ = _1560_v146
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_50
				var _nw277 _dafny.Array
				_ = _nw277
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw277 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) bool = func(_1561_i16 _dafny.Int) bool {
						return !((_this).F25())
					}
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw277 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw277).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw277).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1560_v146 = _nw277
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1560_v146), 0))
				_ = _index279
				(_1560_v146).ArraySet1(!((_this).F25()), (_index279).Int())
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1560_v146), 0))
				_ = _index280
				(_1560_v146).ArraySet1((_this).F25(), (_index280).Int())
				var _1562_v147 _dafny.Array
				_ = _1562_v147
				var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw278
				_1562_v147 = _nw278
				var _1563_v148 D15
				_ = _1563_v148
				_1563_v148 = Companion_D15_.Create_DC33_(_1562_v147)
				_1562_v147 = ((func() D15 {
					if (_1560_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1560_v146), 0))).Int()).(bool) {
						return _1563_v148
					}
					return Companion_D15_.Create_DC33_(_1562_v147)
				})()).Dtor_cf57()
				(globalState).F2 = (_this).F27()
			} else {
				var _1564___mcc_h8 D11 = _source21.Get_().(D11_DC27).Cf43
				_ = _1564___mcc_h8
				var _1565_cf43 D11 = _1564___mcc_h8
				_ = _1565_cf43
				(globalState).F1 = (_this).F25()
				var _1566_v149 _dafny.Array
				_ = _1566_v149
				var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw279
				_1566_v149 = _nw279
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1566_v149), 0))
				_ = _index281
				(_1566_v149).ArraySet1(Companion_Default___.SafeModuloInt((_1531_v126).Cardinality(), Companion_Default___.Fm2((_this).F27(), (_this).F27(), globalState)), (_index281).Int())
				var _1567_v150 _dafny.Map
				_ = _1567_v150
				_1567_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}((func(_1568_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1569_i17 _dafny.Int) _dafny.CodePoint {
						return _1568_v39
					}
				})(_1421_v39)))).Cardinality()), (_this).F27())
				var _1570_v151 _dafny.Map
				_ = _1570_v151
				_1570_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1567_v150, (_1534_v129).Cardinality())
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1566_v149), 0))
				_ = _index282
				(_1566_v149).ArraySet1(Companion_Default___.SafeDivisionInt(((_1570_v151).Merge(_1570_v151)).Cardinality(), Companion_Default___.Fm2((_this).F27(), (_this).F26(), globalState)), (_index282).Int())
				r2 = ((_this).F27()).Times((_this).F27())
				(globalState).F2 = Companion_Default___.SafeModuloInt((_this).F27(), (_this).F26())
			}
			var _1571_v153 _dafny.Map
			_ = _1571_v153
			_1571_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), func() _dafny.Set {
				var _coll33 = _dafny.NewBuilder()
				_ = _coll33
				for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(881), _dafny.IntOfInt64(80))); ; {
					_compr_33, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _1572_v152 _dafny.Int
					_1572_v152 = interface{}(_compr_33).(_dafny.Int)
					if ((_dafny.IntOfInt64(881)).Cmp(_1572_v152) <= 0) && ((_1572_v152).Cmp(_dafny.IntOfInt64(80)) < 0) {
						_coll33.Add(Companion_Default___.SafeModuloInt(_1572_v152, _dafny.IntOfInt64(328)))
					}
				}
				return _coll33.ToSet()
			}())
			var _1573_v154 D2
			_ = _1573_v154
			_1573_v154 = Companion_D2_.Create_DC6_((func() _dafny.Set {
				if (_1571_v153).Contains((_this).F27()) {
					return (_1571_v153).Get((_this).F27()).(_dafny.Set)
				}
				return _1536_v131
			})())
			var _1574_v155 _dafny.Map
			_ = _1574_v155
			_1574_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1536_v131).Cardinality(), (_this).F26())
			var _1575_v156 _dafny.Map
			_ = _1575_v156
			_1575_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
			var _1576_v157 D14
			_ = _1576_v157
			_1576_v157 = Companion_D14_.Create_DC32_((_this).F25(), (_this).F26(), _1575_v156, _this.F28, false)
			var _pat_let_tv35 = _1575_v156
			_ = _pat_let_tv35
			var _rhs290 bool = (_this).F25()
			_ = _rhs290
			var _rhs291 D2 = func(_pat_let42_0 D2) D2 {
				return func(_1577_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let43_0 _dafny.Set) D2 {
						return func(_1578_dt__update_hcf10_h0 _dafny.Set) D2 {
							return Companion_D2_.Create_DC6_(_1578_dt__update_hcf10_h0)
						}(_pat_let43_0)
					}((_dafny.SetOf((_this).F26(), (_this).F27(), _dafny.IntOfInt64(700))).Union(_dafny.SetOf((_this).F27())))
				}(_pat_let42_0)
			}(Companion_Default___.Fm38((func() bool {
				if (_1533_v128).Contains((_this).F27()) {
					return (_1533_v128).Get((_this).F27()).(bool)
				}
				return (_this).F25()
			})(), (_this).F26(), (_dafny.Zero).Minus((_1574_v155).Cardinality()), Companion_Default___.Fm5((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_this).F25()), (_this).F27()), (_this).Fm6((_this).F25(), !((_this).F25()), _dafny.IntOfInt64(220), globalState), globalState), globalState))
			_ = _rhs291
			var _rhs292 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1532_v127).Contains(true) {
					return (_1532_v127).Get(true).(_dafny.Int)
				}
				return (func(_pat_let44_0 D14) D14 {
					return func(_1579_dt__update__tmp_h2 D14) D14 {
						return func(_pat_let45_0 _dafny.Map) D14 {
							return func(_1580_dt__update_hcf54_h0 _dafny.Map) D14 {
								return Companion_D14_.Create_DC32_((_1579_dt__update__tmp_h2).Dtor_cf52(), (_1579_dt__update__tmp_h2).Dtor_cf53(), _1580_dt__update_hcf54_h0, (_1579_dt__update__tmp_h2).Dtor_cf55(), (_1579_dt__update__tmp_h2).Dtor_cf56())
							}(_pat_let45_0)
						}(_pat_let_tv35)
					}(_pat_let44_0)
				}(_1576_v157)).Dtor_cf53()
			})(), (_this).F26())
			_ = _rhs292
			var _rhs293 _dafny.Int = (_this).F27()
			_ = _rhs293
			var _lhs246 *GlobalState = globalState
			_ = _lhs246
			var _lhs247 *GlobalState = globalState
			_ = _lhs247
			var _lhs248 *GlobalState = globalState
			_ = _lhs248
			_lhs246.F1 = _rhs290
			_1573_v154 = _rhs291
			_lhs247.F2 = _rhs292
			_lhs248.F2 = _rhs293
			var _1581_v158 _dafny.Array
			_ = _1581_v158
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_51
			var _nw280 _dafny.Array
			_ = _nw280
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw280 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) bool = func(_1582_i18 _dafny.Int) bool {
					return true
				}
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw280 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw280).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw280).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_1581_v158 = _nw280
			var _1583_v159 *C0
			_ = _1583_v159
			var _nw281 *C0 = New_C0_()
			_ = _nw281
			_nw281.Ctor__(_this.F28, ((_this).F25()) && ((_this).F25()), (_this).F27(), _1581_v158, (_this).F25())
			_1583_v159 = _nw281
		} else {
			var _1584_v160 _dafny.Map
			_ = _1584_v160
			_1584_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()))
			var _1585_v161 _dafny.MultiSet
			_ = _1585_v161
			_1585_v161 = _dafny.MultiSetOf(_this.F28)
			var _1586_v162 _dafny.Array
			_ = _1586_v162
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_52
			var _nw282 _dafny.Array
			_ = _nw282
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw282 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) bool = func(_1587_i19 _dafny.Int) bool {
					return (_this).F25()
				}
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw282 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw282).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw282).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1586_v162 = _nw282
			var _1588_v163 *C0
			_ = _1588_v163
			var _nw283 *C0 = New_C0_()
			_ = _nw283
			_nw283.Ctor__(_this.F28, (_this).F25(), (_this).F27(), _1586_v162, (_this).F25())
			_1588_v163 = _nw283
			var _1589_v164 _dafny.Map
			_ = _1589_v164
			_1589_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v163, (_this).F25())
			var _1590_v165 D0
			_ = _1590_v165
			_1590_v165 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(39), _1588_v163.F35, false, (_dafny.MultiSetOf((_this).F27(), _dafny.IntOfInt64(434), (_this).F26())).Cardinality(), (_this).F25())
			var _1591_v166 *C3
			_ = _1591_v166
			var _nw284 *C3 = New_C3_()
			_ = _nw284
			_nw284.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}((func(_1592_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1593_i20 _dafny.Int) _dafny.CodePoint {
					return _1592_v39
				}
			})(_1421_v39))), (_this).F25(), (_this).F25(), (_this).F26())
			_1591_v166 = _nw284
			var _1594_v169 D16
			_ = _1594_v169
			_1594_v169 = Companion_D16_.Create_DC35_(_1584_v160)
			var _1595_v172 _dafny.MultiSet
			_ = _1595_v172
			_1595_v172 = _dafny.MultiSetOf((_this).F26())
			var _1596_v174 _dafny.Sequence
			_ = _1596_v174
			_1596_v174 = _dafny.SeqOf(_1584_v160, _1584_v160)
			var _1597_v175 _dafny.Map
			_ = _1597_v175
			_1597_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1421_v39)
			var _1598_v176 _dafny.Map
			_ = _1598_v176
			_1598_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v175).Cardinality(), _this.F28)
			var _1599_v177 _dafny.Array
			_ = _1599_v177
			var _nwElement0_53 _dafny.Map = _1584_v160
			_ = _nwElement0_53
			var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(29))
			_ = _nw285
			(_nw285).ArraySet1(_nwElement0_53, 0)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (func() _dafny.Int {
				if (_1585_v161).Contains(_dafny.Companion_Sequence_.Update(_this.F28, (Companion_Default___.SafeIndex((_1589_v164).Cardinality(), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32(), Companion_Default___.Fm4(_1590_v165, (_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1591_v166, (_1591_v166).F30())).Cardinality(), globalState))) {
					return (_1585_v161).Multiplicity(_dafny.Companion_Sequence_.Update(_this.F28, (Companion_Default___.SafeIndex((_1589_v164).Cardinality(), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32(), Companion_Default___.Fm4(_1590_v165, (_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1591_v166, (_1591_v166).F30())).Cardinality(), globalState)))
				}
				return (_this).F26()
			})()), 1)
			(_nw285).ArraySet1((func() _dafny.Map {
				var _coll34 = _dafny.NewMapBuilder()
				_ = _coll34
				for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(83), _dafny.IntOfInt64(340))); ; {
					_compr_34, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _1600_v167 _dafny.Int
					_1600_v167 = interface{}(_compr_34).(_dafny.Int)
					if ((_dafny.IntOfInt64(83)).Cmp(_1600_v167) <= 0) && ((_1600_v167).Cmp(_dafny.IntOfInt64(340)) < 0) {
						_coll34.Add((_1600_v167).Minus((_this).F27()), (_this).F27())
					}
				}
				return _coll34.ToMap()
			}()).Merge(_1584_v160), 2)
			(_nw285).ArraySet1((_1584_v160).Merge(_1584_v160), 3)
			(_nw285).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())).Update((_this).F27(), (_this).F26()), 4)
			(_nw285).ArraySet1(_1584_v160, 5)
			(_nw285).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F26()), (_this).Fm6((_1591_v166).F30(), !((_1591_v166).F30()), (_this).F26(), globalState))).Merge(_1584_v160), 6)
			(_nw285).ArraySet1(_1584_v160, 7)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27()), 8)
			(_nw285).ArraySet1(_1584_v160, 9)
			(_nw285).ArraySet1((_1584_v160).Merge(_1584_v160), 10)
			(_nw285).ArraySet1(func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(491), _dafny.IntOfInt64(-882))); ; {
					_compr_35, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _1601_v168 _dafny.Int
					_1601_v168 = interface{}(_compr_35).(_dafny.Int)
					if ((_dafny.IntOfInt64(491)).Cmp(_1601_v168) <= 0) && ((_1601_v168).Cmp(_dafny.IntOfInt64(-882)) < 0) {
						_coll35.Add((_1601_v168).Plus((_dafny.MultiSetFromSeq((_1591_v166).F29())).Cardinality()), _dafny.IntOfInt64(343))
					}
				}
				return _coll35.ToMap()
			}(), 11)
			(_nw285).ArraySet1(_1584_v160, 12)
			(_nw285).ArraySet1((_1594_v169).Dtor_cf59(), 13)
			(_nw285).ArraySet1((func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(580), _dafny.IntOfInt64(30))); ; {
					_compr_36, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _1602_v170 _dafny.Int
					_1602_v170 = interface{}(_compr_36).(_dafny.Int)
					if ((_dafny.IntOfInt64(580)).Cmp(_1602_v170) <= 0) && ((_1602_v170).Cmp(_dafny.IntOfInt64(30)) < 0) {
						_coll36.Add((_1602_v170).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_1591_v166).F30())).Cardinality()), (_this).F26())
					}
				}
				return _coll36.ToMap()
			}()).Merge(_1584_v160), 14)
			(_nw285).ArraySet1((func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter37 := _dafny.Iterate((_1595_v172).Elements()); ; {
					_compr_37, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1603_v171 _dafny.Int
					_1603_v171 = interface{}(_compr_37).(_dafny.Int)
					if (_1595_v172).Contains(_1603_v171) {
						_coll37.Add(Companion_Default___.SafeModuloInt(_1603_v171, (_this).F27()), (_this).F26())
					}
				}
				return _coll37.ToMap()
			}()).Merge(_1584_v160), 15)
			(_nw285).ArraySet1(func() _dafny.Map {
				var _coll38 = _dafny.NewMapBuilder()
				_ = _coll38
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(77), _dafny.IntOfInt64(-552))); ; {
					_compr_38, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _1604_v173 _dafny.Int
					_1604_v173 = interface{}(_compr_38).(_dafny.Int)
					if ((_dafny.IntOfInt64(77)).Cmp(_1604_v173) <= 0) && ((_1604_v173).Cmp(_dafny.IntOfInt64(-552)) < 0) {
						_coll38.Add((_1604_v173).Plus(_dafny.IntOfInt64(-918)), (_this).F26())
					}
				}
				return _coll38.ToMap()
			}(), 16)
			(_nw285).ArraySet1((_1596_v174).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1596_v174).Cardinality()))).Uint32()).(_dafny.Map), 17)
			(_nw285).ArraySet1((func() _dafny.Map {
				if (_this).F25() {
					return _1584_v160
				}
				return _1584_v160
			})(), 18)
			(_nw285).ArraySet1(_1584_v160, 19)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F27()), 20)
			(_nw285).ArraySet1(_1584_v160, 21)
			(_nw285).ArraySet1(_1584_v160, 22)
			(_nw285).ArraySet1((_1584_v160).Merge(_1584_v160), 23)
			(_nw285).ArraySet1((_1584_v160).Merge(_1584_v160), 24)
			(_nw285).ArraySet1(_1584_v160, 25)
			(_nw285).ArraySet1((_1584_v160).Update(_dafny.IntOfInt64(159), (_1598_v176).Cardinality()), 26)
			(_nw285).ArraySet1((_1584_v160).Merge(_1584_v160), 27)
			(_nw285).ArraySet1(_1584_v160, 28)
			_1599_v177 = _nw285
			var _nw286 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw286
			_1599_v177 = _nw286
			var _1605_v178 _dafny.Array
			_ = _1605_v178
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_53
			var _nw287 _dafny.Array
			_ = _nw287
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw287 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Int = func(_1606_i21 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1606_i21, (_this).F26())
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw287 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw287).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw287).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_1605_v178 = _nw287
			_1605_v178 = (func() _dafny.Array {
				if (_1591_v166).F30() {
					return _1605_v178
				}
				return _1605_v178
			})()
			r1 = (_1591_v166).F30()
			_1421_v39 = _1421_v39
			var _1607_v179 _dafny.Array
			_ = _1607_v179
			var _nw288 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(8))
			_ = _nw288
			_1607_v179 = _nw288
			var _1608_v180 _dafny.Set
			_ = _1608_v180
			_1608_v180 = _dafny.SetOf(_1607_v179)
			(globalState).F1 = (_1608_v180).IsProperSubsetOf(_dafny.SetOf(_1607_v179, _1607_v179))
		}
		var _1609_v181 _dafny.Array
		_ = _1609_v181
		var _nw289 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw289
		_1609_v181 = _nw289
		r0 = _1609_v181
		var _1610_v182 _dafny.Sequence
		_ = _1610_v182
		_1610_v182 = _dafny.SeqOf((_this).F25())
		var _1611_v183 _dafny.MultiSet
		_ = _1611_v183
		_1611_v183 = _dafny.MultiSetOf((_this).F27())
		var _1612_v184 _dafny.MultiSet
		_ = _1612_v184
		_1612_v184 = _dafny.MultiSetOf(_1610_v182, _dafny.SeqOf((_this).F25()))
		r1 = !((_1612_v184).Contains(_dafny.Companion_Sequence_.Update(_1610_v182, (Companion_Default___.SafeIndex(Companion_Default___.Fm2((_this).F27(), (func() _dafny.Int {
			if (_1611_v183).Contains(_dafny.IntOfUint32((_1610_v182).Cardinality())) {
				return (_1611_v183).Multiplicity(_dafny.IntOfUint32((_1610_v182).Cardinality()))
			}
			return (_this).F26()
		})(), globalState), _dafny.IntOfUint32((_1610_v182).Cardinality()))).Uint32(), (_this).F25())))
		r2 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_this).F25() {
				return (_this).F26()
			}
			return (_this).F26()
		})(), (_this).F27())
		return r0, r1, r2
	}
}
func (_this *C4) F27() _dafny.Int {
	{
		return _this._f27
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
