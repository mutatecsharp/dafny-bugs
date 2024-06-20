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
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	var _source0 _dafny.Sequence = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-702))))
	_ = _source0
	var _0___mcc_h0 _dafny.Sequence = _source0
	_ = _0___mcc_h0
	var _1_cf0 _dafny.Sequence = _0___mcc_h0
	_ = _1_cf0
	return true
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.MultiSet, globalState *GlobalState) bool {
	if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("urqk"), _dafny.IntOfInt64(865))).Cardinality()) >= 0 {
		return true
	} else {
		return true
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	if !(!((false) == (true))) {
		return _dafny.SetOf(true)
	} else {
		return (_dafny.SetOf(false, false, false, !(true))).Intersection(_dafny.SetOf(true, (Companion_D16_.Create_DC45_(false)).Dtor_cf61(), false))
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)), _dafny.SeqOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(592), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(227))).Cardinality(), true))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(128)
	})), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(411), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jsxnajks")).Cardinality()))).Cardinality())).Cardinality())), _dafny.IntOfInt64(437)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	if (_dafny.MultiSetOf(_dafny.CodePoint('x'))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.CodePoint('f'), _dafny.CodePoint('r'), _dafny.CodePoint('n'), (Companion_D13_.Create_DC35_(_dafny.CodePoint('c'))).Dtor_cf45(), _dafny.CodePoint('w'))) {
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(405), _dafny.IntOfInt64(994))).Intersection(_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality()))
	} else {
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qo")).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xrq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))), _dafny.UnicodeSeqOfUtf8Bytes("jhpmme"))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Union(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D11_.Create_DC30_(false, _dafny.MultiSetFromSeq(_dafny.SeqOf(true, !(false), true, true)))).Dtor_cf40()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC16_(_dafny.SetOf(false)), Companion_D7_.Create_DC16_(_dafny.SetOf(true, false)), Companion_D7_.Create_DC16_(_dafny.SetOf(!(false), true, !(true), false))), _dafny.SeqOf(Companion_D7_.Create_DC16_(_dafny.SetOf(false))))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(true)))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source1 bool = !(!(true))
	_ = _source1
	var _5___mcc_h0 bool = _source1
	_ = _5___mcc_h0
	var _6_cf1 bool = _5___mcc_h0
	_ = _6_cf1
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qaete")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(687))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!(false), false, false)).Difference(_dafny.MultiSetOf(false, !(!(false)), !(false)))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 _dafny.Set, p3 bool, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC11_(_dafny.IntOfInt64(-410), !((_dafny.SetOf(false)).IsProperSubsetOf(_dafny.SetOf(true, false))), (!(false)) == (!(false)))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(528), _dafny.IntOfInt64(-787))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(528)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(-787)) < 0) {
				_coll0.Add((_7_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(298)))).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.IntOfInt64(949)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, globalState *GlobalState) _dafny.Sequence {
	if false {
		return _dafny.UnicodeSeqOfUtf8Bytes("fix")
	} else {
		return (Companion_D6_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("w"))).Dtor_cf21()
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(634))).Merge((func() _dafny.Map {
		if (Companion_D4_.Create_DC10_(_dafny.MultiSetOf(false), true)).Dtor_cf16() {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(945))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-925))
	})())
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(277))).Cardinality())), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, true, !(true), false, false), false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-614), (_dafny.Zero).Minus((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(6), _dafny.IntOfInt64(99))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _8_v0 _dafny.Int
			_8_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(6)).Cmp(_8_v0) <= 0) && ((_8_v0).Cmp(_dafny.IntOfInt64(99)) < 0) {
				_coll1.Add((_8_v0).Plus((_dafny.MultiSetOf(false)).Cardinality()), true)
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tnsrwmjof")).Cardinality()))).Cardinality()))), true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ilip")).Cardinality()), true)).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(46), _dafny.IntOfInt64(717))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _9_v1 _dafny.Int
			_9_v1 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(46)).Cmp(_9_v1) <= 0) && ((_9_v1).Cmp(_dafny.IntOfInt64(717)) < 0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_9_v1, _dafny.IntOfInt64(132)), (false))
			}
		}
		return _coll2.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Sequence {
	var _source2 D12 = Companion_D12_.Create_DC33_()
	_ = _source2
	if _source2.Is_DC33() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('d'))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(349))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_10_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('v'))
		})))
	} else if _source2.Is_DC32() {
		var _11___mcc_h0 _dafny.Sequence = _source2.Get_().(D12_DC32).Cf43
		_ = _11___mcc_h0
		var _12_cf43 _dafny.Sequence = _11___mcc_h0
		_ = _12_cf43
		return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('d')))
	} else {
		var _13___mcc_h1 D12 = _source2.Get_().(D12_DC34).Cf44
		_ = _13___mcc_h1
		var _14_cf44 D12 = _13___mcc_h1
		_ = _14_cf44
		return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('g'), _dafny.CodePoint('y'), _dafny.CodePoint('d'), _dafny.CodePoint('u'), _dafny.CodePoint('j')), _dafny.MultiSetOf(_dafny.CodePoint('t'), _dafny.CodePoint('u')), _dafny.MultiSetOf(_dafny.CodePoint('i'), _dafny.CodePoint('b')))
	}
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('n'))), _dafny.MultiSetOf(_dafny.CodePoint('g'), _dafny.CodePoint('w')), _dafny.MultiSetOf(_dafny.CodePoint('q')))).Difference(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.CodePoint('f'), _dafny.CodePoint('s'), _dafny.CodePoint('l'))))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.CodePoint('b'))).Difference(_dafny.MultiSetOf(_dafny.CodePoint('k')))).Difference(_dafny.MultiSetOf(_dafny.CodePoint('n')))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qe"), _dafny.UnicodeSeqOfUtf8Bytes("bmhgftehd"), _dafny.UnicodeSeqOfUtf8Bytes("yatjnok"), (Companion_D6_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("ebufvm"))).Dtor_cf21(), _dafny.UnicodeSeqOfUtf8Bytes("ytlonbky")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("lw"))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("oojfql")))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(760))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_15_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(!(false))
	})), !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Contains(_dafny.CodePoint('p')))
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(36))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_16_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(945), true)
	}))).Cardinality())).Cmp(_dafny.IntOfInt64(703)) >= 0, (_dafny.SetOf(!(false))).Difference((Companion_D7_.Create_DC16_(_dafny.SetOf(true))).Dtor_cf24()))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC20_()
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(730), _dafny.IntOfInt64(70), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thphighp")).Cardinality())))).Cardinality(), _dafny.IntOfInt64(-40), _dafny.IntOfInt64(563))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(210), _dafny.IntOfInt64(532), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_17_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality())
	}))).Cardinality()), _dafny.IntOfInt64(561)))).Union(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(163), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(290), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(320))).Cardinality())).Cardinality())).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-678), _dafny.IntOfInt64(-417))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-678)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(-417)) < 0) {
				_coll3.Add((_18_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(!(true))).Cardinality()), !((Companion_D7_.Create_DC17_(true)).Dtor_cf25()))).Cardinality()), false)
			}
		}
		return _coll3.ToMap()
	}(), _dafny.IntOfInt64(235))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(998), _dafny.IntOfInt64(447))).Cardinality())).Cardinality(), !(true)), _dafny.IntOfInt64(951))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()), false), _dafny.IntOfInt64(709))))
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC25_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(457))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(619))))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(true), _dafny.SeqOf((Companion_D16_.Create_DC45_(!(true))).Dtor_cf61()))).Intersection(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(true))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _19_v0 _dafny.Sequence
			_19_v0 = interface{}(_compr_4).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(true)), _19_v0) {
				_coll4.Add(_19_v0)
			}
		}
		return _coll4.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 bool, globalState *GlobalState) D3 {
	if !(((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bvsgx")).Cardinality())), false)).Cardinality())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tanewc")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(410), _dafny.IntOfInt64(632))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(410)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(632)) < 0) {
				_coll5.Add((_20_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('f'))).Cardinality())), _dafny.IntOfInt64(78))
			}
		}
		return _coll5.ToMap()
	}())).Cardinality())))).Cardinality()) >= 0) {
		return Companion_D3_.Create_DC6_((_dafny.Zero).Minus((func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(452))).Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _21_v2 _dafny.Int
					_21_v2 = interface{}(_compr_7).(_dafny.Int)
					if (_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(452))).Contains(_21_v2) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_21_v2, _dafny.IntOfInt64(-222)))
					}
				}
				return _coll7.ToSet()
			}()).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _22_v1 _dafny.Int
				_22_v1 = interface{}(_compr_6).(_dafny.Int)
				if (func() _dafny.Set {
					var _coll8 = _dafny.NewBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(452))).Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _23_v2 _dafny.Int
						_23_v2 = interface{}(_compr_8).(_dafny.Int)
						if (_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(452))).Contains(_23_v2) {
							_coll8.Add(Companion_Default___.SafeModuloInt(_23_v2, _dafny.IntOfInt64(-222)))
						}
					}
					return _coll8.ToSet()
				}()).Contains(_22_v1) {
					_coll6.Add(Companion_Default___.SafeModuloInt(_22_v1, (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Cardinality()), true)
				}
			}
			return _coll6.ToMap()
		}()).Cardinality()))
	} else {
		return Companion_D3_.Create_DC6_((func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(162))).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _24_v3 _dafny.Int
				_24_v3 = interface{}(_compr_9).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(162))).Contains(_24_v3) {
					_coll9.Add((_24_v3).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-187))))
				}
			}
			return _coll9.ToSet()
		}()).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _25_v0 _dafny.Sequence
	_ = _25_v0
	_25_v0 = _dafny.UnicodeSeqOfUtf8Bytes("mrqsc")
	var _26_v1 _dafny.Int
	_ = _26_v1
	_26_v1 = _dafny.IntOfInt64(-60)
	var _27_v2 _dafny.Array
	_ = _27_v2
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
	_ = _nw0
	_27_v2 = _nw0
	var _28_v3 _dafny.Map
	_ = _28_v3
	_28_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v1, _27_v2)
	var _29_v4 _dafny.Set
	_ = _29_v4
	_29_v4 = _dafny.SetOf(_26_v1)
	var _30_v5 _dafny.Map
	_ = _30_v5
	_30_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v1, _29_v4)
	var _31_v6 _dafny.Map
	_ = _31_v6
	_31_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((func() _dafny.Set {
		if (_30_v5).Contains(_26_v1) {
			return (_30_v5).Get(_26_v1).(_dafny.Set)
		}
		return _29_v4
	})()).Cardinality())
	var _32_v7 bool
	_ = _32_v7
	_32_v7 = true
	var _33_v8 _dafny.Map
	_ = _33_v8
	_33_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v1, _26_v1)
	var _34_v9 _dafny.Set
	_ = _34_v9
	_34_v9 = _dafny.SetOf(_32_v7)
	var _35_v10 _dafny.Map
	_ = _35_v10
	_35_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_34_v9).Cardinality(), _32_v7)
	var _36_v12 _dafny.Map
	_ = _36_v12
	_36_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_32_v7), _33_v8)
	var _37_v13 _dafny.Map
	_ = _37_v13
	_37_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_26_v1), _33_v8)
	var _38_v16 _dafny.Sequence
	_ = _38_v16
	_38_v16 = _dafny.SeqOf(_26_v1, _26_v1)
	var _39_v17 _dafny.Map
	_ = _39_v17
	_39_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_25_v0).Cardinality()), _38_v16)
	var _40_v18 _dafny.Array
	_ = _40_v18
	var _nwElement0_0 _dafny.Map = (_33_v8).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_32_v7, _32_v7, true), (Companion_Default___.SafeIndex(_26_v1, _dafny.IntOfUint32((_dafny.SeqOf(_32_v7, _32_v7, true)).Cardinality()))).Uint32(), _32_v7)).Cardinality()), (_35_v10).Cardinality())
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(19))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1(_33_v8, 1)
	(_nw1).ArraySet1(_33_v8, 2)
	(_nw1).ArraySet1(_33_v8, 3)
	(_nw1).ArraySet1(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(419), _dafny.IntOfInt64(570))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _41_v11 _dafny.Int
			_41_v11 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(419)).Cmp(_41_v11) <= 0) && ((_41_v11).Cmp(_dafny.IntOfInt64(570)) < 0) {
				_coll10.Add(Companion_Default___.SafeDivisionInt(_41_v11, _26_v1), _dafny.IntOfInt64(898))
			}
		}
		return _coll10.ToMap()
	}(), 4)
	(_nw1).ArraySet1(_33_v8, 5)
	(_nw1).ArraySet1(_33_v8, 6)
	(_nw1).ArraySet1((func() _dafny.Map {
		if (_36_v12).Contains(_32_v7) {
			return (_36_v12).Get(_32_v7).(_dafny.Map)
		}
		return _33_v8
	})(), 7)
	(_nw1).ArraySet1(_33_v8, 8)
	(_nw1).ArraySet1(_33_v8, 9)
	(_nw1).ArraySet1(_33_v8, 10)
	(_nw1).ArraySet1(_33_v8, 11)
	(_nw1).ArraySet1((func() _dafny.Map {
		if (_37_v13).Contains(_26_v1) {
			return (_37_v13).Get(_26_v1).(_dafny.Map)
		}
		return _33_v8
	})(), 12)
	(_nw1).ArraySet1(_33_v8, 13)
	(_nw1).ArraySet1(_33_v8, 14)
	(_nw1).ArraySet1((func() _dafny.Map {
		if (_37_v13).Contains(_26_v1) {
			return (_37_v13).Get(_26_v1).(_dafny.Map)
		}
		return _33_v8
	})(), 15)
	(_nw1).ArraySet1(_33_v8, 16)
	(_nw1).ArraySet1(func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_29_v4).Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _42_v14 _dafny.Int
			_42_v14 = interface{}(_compr_11).(_dafny.Int)
			if (_29_v4).Contains(_42_v14) {
				_coll11.Add(Companion_Default___.SafeDivisionInt(_42_v14, _26_v1), _26_v1)
			}
		}
		return _coll11.ToMap()
	}(), 17)
	(_nw1).ArraySet1(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_39_v17).Keys().Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _43_v15 _dafny.Int
			_43_v15 = interface{}(_compr_12).(_dafny.Int)
			if (_39_v17).Contains(_43_v15) {
				_coll12.Add((_43_v15).Minus(_26_v1), _26_v1)
			}
		}
		return _coll12.ToMap()
	}(), 18)
	_40_v18 = _nw1
	var _44_v19 _dafny.Sequence
	_ = _44_v19
	_44_v19 = _dafny.SeqOf(_25_v0)
	var _45_v20 _dafny.Map
	_ = _45_v20
	_45_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v2, _dafny.IntOfUint32((_44_v19).Cardinality()))
	var _46_v21 _dafny.CodePoint
	_ = _46_v21
	_46_v21 = _dafny.CodePoint('m')
	var _47_globalState *GlobalState
	_ = _47_globalState
	var _nw2 *GlobalState = New_GlobalState_()
	_ = _nw2
	_nw2.Ctor__(_dafny.IntOfInt64(687), _dafny.Companion_Sequence_.Concatenate(_25_v0, _25_v0), true, false, false, _dafny.IntOfInt64(351), (func() _dafny.Array {
		if (_28_v3).Contains((func() _dafny.Int {
			if (_31_v6).Contains(_32_v7) {
				return (_31_v6).Get(_32_v7).(_dafny.Int)
			}
			return _26_v1
		})()) {
			return (_28_v3).Get((func() _dafny.Int {
				if (_31_v6).Contains(_32_v7) {
					return (_31_v6).Get(_32_v7).(_dafny.Int)
				}
				return _26_v1
			})()).(_dafny.Array)
		}
		return _27_v2
	})(), false, _40_v18, _dafny.IntOfInt64(825), false, _dafny.IntOfInt64(-252), false, _dafny.IntOfInt64(318), true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(774))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), (_31_v6).Merge(_31_v6), _dafny.IntOfInt64(124), false, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_25_v0, (Companion_Default___.SafeIndex((_45_v20).Cardinality(), _dafny.IntOfUint32((_25_v0).Cardinality()))).Uint32(), _46_v21), _25_v0))
	_47_globalState = _nw2
	var _49_v22 _dafny.Array
	_ = _49_v22
	var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
	_ = _nw3
	_49_v22 = _nw3
	var _50_v23 *C1
	_ = _50_v23
	var _nw4 *C1 = New_C1_()
	_ = _nw4
	_nw4.Ctor__(_26_v1, _49_v22, _dafny.IntOfInt64(542), _32_v7)
	_50_v23 = _nw4
	_34_v9 = _34_v9
	var _51_v24 _dafny.Map
	_ = _51_v24
	_51_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v7, _34_v9)
	_34_v9 = (((func() _dafny.Set {
		if (_51_v24).Contains(_32_v7) {
			return (_51_v24).Get(_32_v7).(_dafny.Set)
		}
		return _34_v9
	})()).Intersection(_34_v9)).Union((_34_v9).Difference(_34_v9))
	var _52_v25 _dafny.Sequence
	_ = _52_v25
	var _53_v26 bool
	_ = _53_v26
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 bool
	_ = _out1
	_out0, _out1 = (_50_v23).M0(Companion_Default___.SafeModuloInt(_26_v1, (_dafny.Zero).Minus(_26_v1)), _32_v7, _47_globalState)
	_52_v25 = _out0
	_53_v26 = _out1
	(_47_globalState).F3 = _32_v7
	var _54_v27 _dafny.Array
	_ = _54_v27
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_0
	var _nw5 _dafny.Array
	_ = _nw5
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw5 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = (func(_55_v10 _dafny.Map, _56_v23 *C1) func(_dafny.Int) _dafny.Map {
			return func(_57_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v10, _56_v23.F27)
			}
		})(_35_v10, _50_v23)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw5 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw5).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw5).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_54_v27 = _nw5
	var _58_v28 _dafny.Map
	_ = _58_v28
	_58_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v23.F27, _53_v26))
	var _59_v29 _dafny.Map
	_ = _59_v29
	_59_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		if (_58_v28).Contains(_32_v7) {
			return (_58_v28).Get(_32_v7).(_dafny.Map)
		}
		return _35_v10
	})(), _50_v23.F27)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_54_v27), 0))
	_ = _index0
	(_54_v27).ArraySet1((_59_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v10, _50_v23.F27)), (_index0).Int())
	var _60_v30 *C5
	_ = _60_v30
	var _nw6 *C5 = New_C5_()
	_ = _nw6
	_nw6.Ctor__(_53_v26, _26_v1, _53_v26)
	_60_v30 = _nw6
	var _61_v31 _dafny.Map
	_ = _61_v31
	_61_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v26, _60_v30)
	var _62_v34 _dafny.Map
	_ = _62_v34
	_62_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v26, _60_v30.F23)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_54_v27), 0))
	_ = _index1
	(_54_v27).ArraySet1((func() _dafny.Map {
		if !(_61_v31).Contains(_32_v7) {
			return (func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate(((func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate(((_59_v29).Update(_35_v10, (_62_v34).Cardinality())).Keys().Elements()); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _63_v33 _dafny.Map
						_63_v33 = interface{}(_compr_14).(_dafny.Map)
						if ((_59_v29).Update(_35_v10, (_62_v34).Cardinality())).Contains(_63_v33) {
							_coll14.Add(_63_v33, false)
						}
					}
					return _coll14.ToMap()
				}()).Update(_35_v10, _60_v30.F23)).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _64_v32 _dafny.Map
					_64_v32 = interface{}(_compr_13).(_dafny.Map)
					if ((func() _dafny.Map {
						var _coll15 = _dafny.NewMapBuilder()
						_ = _coll15
						for _iter15 := _dafny.Iterate(((_59_v29).Update(_35_v10, (_62_v34).Cardinality())).Keys().Elements()); ; {
							_compr_15, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _63_v33 _dafny.Map
							_63_v33 = interface{}(_compr_15).(_dafny.Map)
							if ((_59_v29).Update(_35_v10, (_62_v34).Cardinality())).Contains(_63_v33) {
								_coll15.Add(_63_v33, false)
							}
						}
						return _coll15.ToMap()
					}()).Update(_35_v10, _60_v30.F23)).Contains(_64_v32) {
						_coll13.Add(_64_v32, _50_v23.F27)
					}
				}
				return _coll13.ToMap()
			}()).Merge(_59_v29)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v10, _50_v23.F27)
	})(), (_index1).Int())
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_46_v21)).Cardinality())
	_ = _hi0
	for _65_i2 := (_dafny.IntOfInt64(676)).Minus(_50_v23.F27); _65_i2.Cmp(_hi0) < 0; _65_i2 = _65_i2.Plus(_dafny.One) {
		var _66_v35 _dafny.Array
		_ = _66_v35
		_66_v35 = _49_v22
		var _67_v36 _dafny.Map
		_ = _67_v36
		_67_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
			if _60_v30.F23 {
				return Companion_Default___.Fm17(_53_v26, _47_globalState)
			}
			return _46_v21
		})(), _66_v35)
		var _68_v37 _dafny.Array
		_ = _68_v37
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_1
		var _nw7 _dafny.Array
		_ = _nw7
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw7 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.MultiSet = (func(_69_v23 *C1) func(_dafny.Int) _dafny.MultiSet {
				return func(_70_i3 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_69_v23.F27, _dafny.IntOfInt64(195))
				}
			})(_50_v23)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw7 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw7).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw7).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_68_v37 = _nw7
		var _71_v38 _dafny.MultiSet
		_ = _71_v38
		_71_v38 = _dafny.MultiSetOf(_26_v1)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_68_v37), 0))
		_ = _index2
		(_68_v37).ArraySet1((_71_v38).Difference((_71_v38).Update(_50_v23.F27, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality())))), (_index2).Int())
		var _72_v39 _dafny.Array
		_ = _72_v39
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
		_ = _nw8
		_72_v39 = _nw8
		var _73_v40 D6
		_ = _73_v40
		_73_v40 = Companion_D6_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("ip"))
		var _74_v41 _dafny.Sequence
		_ = _74_v41
		_74_v41 = _dafny.SeqOf(_62_v34)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_68_v37), 0))
		_ = _index3
		var _rhs0 bool = !((_50_v23).Fm0((_33_v8).Merge(_33_v8), (_73_v40).Dtor_cf21(), _dafny.Companion_Sequence_.Concatenate(_74_v41, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_74_v41, (Companion_Default___.SafeIndex(_50_v23.F27, _dafny.IntOfUint32((_74_v41).Cardinality()))).Uint32(), _62_v34), (Companion_Default___.SafeIndex(_50_v23.F27, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_74_v41, (Companion_Default___.SafeIndex(_50_v23.F27, _dafny.IntOfUint32((_74_v41).Cardinality()))).Uint32(), _62_v34)).Cardinality()))).Uint32(), _62_v34)), _47_globalState))
		_ = _rhs0
		var _rhs1 _dafny.Map = _67_v36
		_ = _rhs1
		var _rhs2 _dafny.MultiSet = ((_71_v38).Update(_dafny.IntOfInt64(14), Companion_Default___.Abs(_50_v23.F27))).Union((_71_v38).Union(_71_v38))
		_ = _rhs2
		var _rhs3 bool = _53_v26
		_ = _rhs3
		var _rhs4 _dafny.Array = _72_v39
		_ = _rhs4
		var _lhs0 *GlobalState = _47_globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _68_v37
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_68_v37), 0))
		_ = _lhs2
		var _lhs3 *GlobalState = _47_globalState
		_ = _lhs3
		_lhs0.F14 = _rhs0
		_67_v36 = _rhs1
		(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
		_lhs3.F14 = _rhs3
		_72_v39 = _rhs4
		(_50_v23).F27 = _65_i2
		_49_v22 = _49_v22
		(_47_globalState).F3 = (_50_v23.F27).Cmp(_dafny.IntOfInt64(214)) != 0
	}
	(_47_globalState).F17 = _50_v23.F27
	var _75_v42 D13
	_ = _75_v42
	_75_v42 = Companion_D13_.Create_DC36_(_50_v23.F27, (_dafny.SetOf(_50_v23.F27, _50_v23.F27)).IsProperSubsetOf(_29_v4), _32_v7)
	var _source3 D13 = _75_v42
	_ = _source3
	if _source3.Is_DC36() {
		var _76___mcc_h0 _dafny.Int = _source3.Get_().(D13_DC36).Cf46
		_ = _76___mcc_h0
		var _77___mcc_h1 bool = _source3.Get_().(D13_DC36).Cf47
		_ = _77___mcc_h1
		var _78___mcc_h2 bool = _source3.Get_().(D13_DC36).Cf48
		_ = _78___mcc_h2
		var _79_cf48 bool = _78___mcc_h2
		_ = _79_cf48
		var _80_cf47 bool = _77___mcc_h1
		_ = _80_cf47
		var _81_cf46 _dafny.Int = _76___mcc_h0
		_ = _81_cf46
		(_47_globalState).F4 = _60_v30.F23
		var _82_v43 bool
		_ = _82_v43
		_82_v43 = _32_v7
		var _83_v44 *C2
		_ = _83_v44
		var _nw9 *C2 = New_C2_()
		_ = _nw9
		_nw9.Ctor__(Companion_Default___.SafeDivisionInt(_50_v23.F27, _50_v23.F27), (_82_v43))
		_83_v44 = _nw9
		var _84_v45 *C4
		_ = _84_v45
		var _nw10 *C4 = New_C4_()
		_ = _nw10
		_nw10.Ctor__(!(_79_cf48), _50_v23.F27, _32_v7)
		_84_v45 = _nw10
		var _85_v46 _dafny.MultiSet
		_ = _85_v46
		_85_v46 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_25_v0).Cardinality()), _26_v1)).Cardinality(), _26_v1, _81_cf46)
		var _86_v47 _dafny.Map
		_ = _86_v47
		_86_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v45, _85_v46)
		var _87_v48 _dafny.Sequence
		_ = _87_v48
		_87_v48 = _dafny.SeqOf(_53_v26)
		var _88_v49 _dafny.Map
		_ = _88_v49
		_88_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.MultiSet {
			if (_86_v47).Contains(_84_v45) {
				return (_86_v47).Get(_84_v45).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_26_v1, (_dafny.Zero).Minus(_26_v1), _50_v23.F27, _26_v1, _dafny.IntOfUint32((_87_v48).Cardinality()))
		})()).IsDisjointFrom(_85_v46), _25_v0)
		_81_cf46 = (_88_v49).Cardinality()
		(_60_v30).F23 = (func() bool {
			if (_35_v10).Contains(_50_v23.F27) {
				return (_35_v10).Get(_50_v23.F27).(bool)
			}
			return (_60_v30.F23) == (_32_v7)
		})()
	} else if _source3.Is_DC37() {
		var _89___mcc_h3 bool = _source3.Get_().(D13_DC37).Cf49
		_ = _89___mcc_h3
		var _90___mcc_h4 _dafny.Int = _source3.Get_().(D13_DC37).Cf50
		_ = _90___mcc_h4
		var _91_cf50 _dafny.Int = _90___mcc_h4
		_ = _91_cf50
		var _92_cf49 bool = _89___mcc_h3
		_ = _92_cf49
		(_47_globalState).F18 = _92_cf49
		var _93_v50 _dafny.Map
		_ = _93_v50
		_93_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_29_v4).Intersection(_29_v4), _dafny.SetOf(!(_60_v30.F23)))
		_91_cf50 = (_93_v50).Cardinality()
		_46_v21 = _dafny.CodePoint('t')
		var _94_v51 _dafny.Sequence
		_ = _94_v51
		_94_v51 = _dafny.SeqOf(_32_v7)
		if (_dafny.IntOfUint32((_94_v51).Cardinality())).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_26_v1, _91_cf50))) <= 0 {
			var _95_v52 _dafny.Sequence
			_ = _95_v52
			_95_v52 = _dafny.SeqOf(_27_v2, _27_v2)
			_95_v52 = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_27_v2, _27_v2, _27_v2), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_26_v1, _26_v1), _dafny.IntOfUint32((_dafny.SeqOf(_27_v2, _27_v2, _27_v2)).Cardinality()))).Uint32(), _27_v2)
			_33_v8 = (_33_v8).Update((func() _dafny.Int {
				if _60_v30.F23 {
					return _50_v23.F27
				}
				return (_dafny.Zero).Minus((_33_v8).Cardinality())
			})(), _50_v23.F27)
			var _96_v53 T0
			_ = _96_v53
			var _nw11 *C5 = New_C5_()
			_ = _nw11
			_nw11.Ctor__(_60_v30.F23, _50_v23.F27, _32_v7)
			_96_v53 = _nw11
			_96_v53 = _96_v53
			_46_v21 = _46_v21
			var _97_v54 _dafny.Map
			_ = _97_v54
			_97_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v21, _29_v4)
			_97_v54 = _97_v54
		} else {
			(_47_globalState).F5 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('d'), _46_v21), _25_v0)).Cardinality())
			_31_v6 = (_31_v6).Update(_60_v30.F23, _91_cf50)
			var _98_v55 _dafny.Array
			_ = _98_v55
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw12
			_98_v55 = _nw12
			var _99_v56 D18
			_ = _99_v56
			_99_v56 = Companion_D18_.Create_DC50_(_98_v55)
			var _100_v57 _dafny.Array
			_ = _100_v57
			var _nwElement0_1 _dafny.Array = _98_v55
			_ = _nwElement0_1
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(17))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_1, 0)
			(_nw13).ArraySet1(_98_v55, 1)
			(_nw13).ArraySet1(_98_v55, 2)
			(_nw13).ArraySet1(_98_v55, 3)
			(_nw13).ArraySet1(_98_v55, 4)
			(_nw13).ArraySet1((_99_v56).Dtor_cf69(), 5)
			(_nw13).ArraySet1(_98_v55, 6)
			(_nw13).ArraySet1(_98_v55, 7)
			(_nw13).ArraySet1(_98_v55, 8)
			(_nw13).ArraySet1(_98_v55, 9)
			(_nw13).ArraySet1(_98_v55, 10)
			(_nw13).ArraySet1(_98_v55, 11)
			(_nw13).ArraySet1(_98_v55, 12)
			(_nw13).ArraySet1(_98_v55, 13)
			(_nw13).ArraySet1((func() _dafny.Array {
				if _53_v26 {
					return (Companion_D18_.Create_DC50_(_98_v55)).Dtor_cf69()
				}
				return _98_v55
			})(), 14)
			(_nw13).ArraySet1(_98_v55, 15)
			(_nw13).ArraySet1(_98_v55, 16)
			_100_v57 = _nw13
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_100_v57), 0))
			_ = _index4
			(_100_v57).ArraySet1(_98_v55, (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_100_v57), 0))
			_ = _index5
			(_100_v57).ArraySet1(_98_v55, (_index5).Int())
			var _101_v58 _dafny.Sequence
			_ = _101_v58
			_101_v58 = _dafny.SeqOf(_34_v9, _34_v9, _34_v9)
			var _102_v59 D3
			_ = _102_v59
			_102_v59 = Companion_D3_.Create_DC7_(_101_v58, _60_v30.F23)
			var _103_v60 *C4
			_ = _103_v60
			var _nw14 *C4 = New_C4_()
			_ = _nw14
			_nw14.Ctor__(true, _50_v23.F27, _32_v7)
			_103_v60 = _nw14
			var _104_v61 _dafny.Map
			_ = _104_v61
			_104_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v60, _50_v23.F27)
			var _105_v62 D19
			_ = _105_v62
			_105_v62 = Companion_D19_.Create_DC54_(_104_v61)
			var _106_v63 *C3
			_ = _106_v63
			var _nw15 *C3 = New_C3_()
			_ = _nw15
			_nw15.Ctor__(_50_v23.F27, (_102_v59).Dtor_cf12(), (_dafny.Zero).Minus(_50_v23.F27), ((_105_v62).Dtor_cf72()).Contains(_103_v60))
			_106_v63 = _nw15
			(_60_v30).F23 = ((_34_v9).Union(_34_v9)).IsDisjointFrom((_34_v9).Union(_dafny.SetOf((_103_v60).F24())))
		}
	} else {
		var _107___mcc_h5 _dafny.CodePoint = _source3.Get_().(D13_DC35).Cf45
		_ = _107___mcc_h5
		var _108_cf45 _dafny.CodePoint = _107___mcc_h5
		_ = _108_cf45
		var _109_v64 *C2
		_ = _109_v64
		var _nw16 *C2 = New_C2_()
		_ = _nw16
		_nw16.Ctor__(_50_v23.F27, _32_v7)
		_109_v64 = _nw16
		_109_v64 = _109_v64
		var _110_v65 _dafny.Array
		_ = _110_v65
		_110_v65 = _49_v22
		var _111_v66 _dafny.Map
		_ = _111_v66
		_111_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_v30.F23, (_50_v23).F28())
		var _112_v67 _dafny.Array
		_ = _112_v67
		var _nwElement0_2 _dafny.Array = _49_v22
		_ = _nwElement0_2
		var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(16))
		_ = _nw17
		(_nw17).ArraySet1(_nwElement0_2, 0)
		(_nw17).ArraySet1(_110_v65, 1)
		(_nw17).ArraySet1(_110_v65, 2)
		(_nw17).ArraySet1(_110_v65, 3)
		(_nw17).ArraySet1(_110_v65, 4)
		(_nw17).ArraySet1((func() _dafny.Array {
			if false {
				return (func() _dafny.Array {
					if (_111_v66).Contains(_60_v30.F23) {
						return (_111_v66).Get(_60_v30.F23).(_dafny.Array)
					}
					return (_50_v23).F28()
				})()
			}
			return (_50_v23).F28()
		})(), 5)
		(_nw17).ArraySet1(_110_v65, 6)
		(_nw17).ArraySet1(_110_v65, 7)
		(_nw17).ArraySet1((_50_v23).F28(), 8)
		(_nw17).ArraySet1(_110_v65, 9)
		(_nw17).ArraySet1(_110_v65, 10)
		(_nw17).ArraySet1((_50_v23).F28(), 11)
		(_nw17).ArraySet1(_110_v65, 12)
		(_nw17).ArraySet1((_50_v23).F28(), 13)
		(_nw17).ArraySet1(_110_v65, 14)
		(_nw17).ArraySet1(_110_v65, 15)
		_112_v67 = _nw17
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_112_v67), 0))
		_ = _index6
		(_112_v67).ArraySet1((_50_v23).F28(), (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_112_v67), 0))
		_ = _index7
		var _rhs5 _dafny.Array = _110_v65
		_ = _rhs5
		var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_38_v16, _38_v16)
		_ = _rhs6
		var _rhs7 _dafny.Int = _50_v23.F27
		_ = _rhs7
		var _lhs4 _dafny.Array = _112_v67
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_112_v67), 0))
		_ = _lhs5
		var _lhs6 *GlobalState = _47_globalState
		_ = _lhs6
		(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
		_38_v16 = _rhs6
		_lhs6.F17 = _rhs7
		(_47_globalState).F4 = ((_62_v34).Cardinality()).Cmp(_26_v1) < 0
		var _113_v68 *C5
		_ = _113_v68
		var _nw18 *C5 = New_C5_()
		_ = _nw18
		_nw18.Ctor__(!(_60_v30.F23), _dafny.IntOfInt64(651), _32_v7)
		_113_v68 = _nw18
	}
	(_47_globalState).F17 = _26_v1
	(_47_globalState).F5 = _26_v1
	var _114_v69 D17
	_ = _114_v69
	_114_v69 = Companion_D17_.Create_DC47_(Companion_Default___.SafeModuloInt((Companion_Default___.Fm41((_dafny.Zero).Minus(_26_v1), (func() _dafny.Int {
		if (_31_v6).Contains(_32_v7) {
			return (_31_v6).Get(_32_v7).(_dafny.Int)
		}
		return (_50_v23).Fm1(_26_v1, _26_v1, _60_v30.F23, _47_globalState)
	})(), _47_globalState)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_25_v0).Cardinality()))), _60_v30.F23, (_26_v1).Plus((_33_v8).Cardinality()))
	_114_v69 = _114_v69
	var _115_v70 *C0
	_ = _115_v70
	var _nw19 *C0 = New_C0_()
	_ = _nw19
	_nw19.Ctor__()
	_115_v70 = _nw19
	var _116_v71 _dafny.Array
	_ = _116_v71
	var _nwElement0_3 *C0 = _115_v70
	_ = _nwElement0_3
	var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(12))
	_ = _nw20
	(_nw20).ArraySet1(_nwElement0_3, 0)
	(_nw20).ArraySet1(_115_v70, 1)
	(_nw20).ArraySet1(_115_v70, 2)
	(_nw20).ArraySet1(_115_v70, 3)
	(_nw20).ArraySet1(_115_v70, 4)
	(_nw20).ArraySet1(_115_v70, 5)
	(_nw20).ArraySet1(_115_v70, 6)
	(_nw20).ArraySet1((func() *C0 {
		if _60_v30.F23 {
			return _115_v70
		}
		return _115_v70
	})(), 7)
	(_nw20).ArraySet1(_115_v70, 8)
	(_nw20).ArraySet1(_115_v70, 9)
	(_nw20).ArraySet1(_115_v70, 10)
	(_nw20).ArraySet1(_115_v70, 11)
	_116_v71 = _nw20
	var _117_v72 _dafny.Sequence
	_ = _117_v72
	_117_v72 = _dafny.SeqOf(_115_v70)
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_116_v71), 0))
	_ = _index8
	(_116_v71).ArraySet1((_117_v72).Select((Companion_Default___.SafeIndex(_50_v23.F27, _dafny.IntOfUint32((_117_v72).Cardinality()))).Uint32()).(*C0), (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_116_v71), 0))
	_ = _index9
	(_116_v71).ArraySet1(_115_v70, (_index9).Int())
	var _118_v73 _dafny.Array
	_ = _118_v73
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_2
	var _nw21 _dafny.Array
	_ = _nw21
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw21 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Set = (func(_119_v9 _dafny.Set) func(_dafny.Int) _dafny.Set {
			return func(_120_i5 _dafny.Int) _dafny.Set {
				return _119_v9
			}
		})(_34_v9)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw21 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw21).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw21).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_118_v73 = _nw21
	for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_118_v73), 0))); ; {
		_guard_loop_0, _ok16 := _iter16()
		if !_ok16 {
			break
		}
		var _121_i4 _dafny.Int
		_121_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_121_i4).Sign() != -1) && ((_121_i4).Cmp(_dafny.ArrayLen((_118_v73), 0)) < 0)) {
			(_118_v73).ArraySet1((Companion_D7_.Create_DC16_(_34_v9)).Dtor_cf24(), (_121_i4).Int())
		}
	}
	_46_v21 = Companion_Default___.Fm17(_53_v26, _47_globalState)
	var _122_v74 D12
	_ = _122_v74
	_122_v74 = Companion_D12_.Create_DC33_()
	var _123_i6 _dafny.Int
	_ = _123_i6
	_123_i6 = _dafny.Zero
	{
		var _pat_let_tv0 = _50_v23
		_ = _pat_let_tv0
		var _pat_let_tv1 = _50_v23
		_ = _pat_let_tv1
		var _pat_let_tv2 = _60_v30
		_ = _pat_let_tv2
		var _pat_let_tv3 = _32_v7
		_ = _pat_let_tv3
		for func(_source4 D12) bool {
			if _source4.Is_DC33() {
				return (_pat_let_tv0.F27).Cmp(_pat_let_tv1.F27) != 0
			} else if _source4.Is_DC32() {
				var _133___mcc_h6 _dafny.Sequence = _source4.Get_().(D12_DC32).Cf43
				_ = _133___mcc_h6
				var _134_cf43 _dafny.Sequence = _133___mcc_h6
				_ = _134_cf43
				return _pat_let_tv2.F23
			} else {
				var _135___mcc_h7 D12 = _source4.Get_().(D12_DC34).Cf44
				_ = _135___mcc_h7
				var _136_cf44 D12 = _135___mcc_h7
				_ = _136_cf44
				return _pat_let_tv3
			}
		}(_122_v74) {
			{
				if (_123_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_123_i6 = (_123_i6).Plus(_dafny.One)
				var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_27_v2), 0))
				_ = _index10
				(_27_v2).ArraySet1(_60_v30.F23, (_index10).Int())
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_27_v2), 0))
				_ = _index11
				(_27_v2).ArraySet1(_32_v7, (_index11).Int())
				var _124_v75 _dafny.Array
				_ = _124_v75
				var _nw22 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
				_ = _nw22
				_124_v75 = _nw22
				var _125_v76 _dafny.Sequence
				_ = _125_v76
				_125_v76 = _dafny.SeqOf(_53_v26)
				var _126_v77 *C6
				_ = _126_v77
				var _nw23 *C6 = New_C6_()
				_ = _nw23
				_nw23.Ctor__(_25_v0, (_31_v6).Cardinality(), (_125_v76).Select((Companion_Default___.SafeIndex(_26_v1, _dafny.IntOfUint32((_125_v76).Cardinality()))).Uint32()).(bool))
				_126_v77 = _nw23
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_124_v75), 0))
				_ = _index12
				(_124_v75).ArraySet1(_126_v77, (_index12).Int())
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_27_v2), 0))
				_ = _index13
				(_27_v2).ArraySet1(_53_v26, (_index13).Int())
				var _127_v78 _dafny.MultiSet
				_ = _127_v78
				_127_v78 = _dafny.MultiSetOf((_126_v77).F22())
				var _128_v79 D16
				_ = _128_v79
				_128_v79 = Companion_D16_.Create_DC43_(_26_v1, _127_v78)
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_27_v2), 0))
				_ = _index14
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_27_v2), 0))
				_ = _index15
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_124_v75), 0))
				_ = _index16
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_27_v2), 0))
				_ = _index17
				var _rhs8 bool = _60_v30.F23
				_ = _rhs8
				var _rhs9 bool = _53_v26
				_ = _rhs9
				var _rhs10 _dafny.Int = (func() _dafny.Int {
					if _32_v7 {
						return ((_60_v30).Fm1((_50_v23).Fm1((_128_v79).Dtor_cf55(), _26_v1, _53_v26, _47_globalState), _dafny.IntOfUint32((Companion_Default___.Fm21(_47_globalState)).Cardinality()), _60_v30.F23, _47_globalState)).Minus(_26_v1)
					}
					return (_dafny.IntOfUint32(((_126_v77).F22()).Cardinality())).Times(_50_v23.F27)
				})()
				_ = _rhs10
				var _rhs11 *C6 = (func() *C6 {
					if _53_v26 {
						return _126_v77
					}
					return _126_v77
				})()
				_ = _rhs11
				var _rhs12 bool = false
				_ = _rhs12
				var _lhs7 _dafny.Array = _27_v2
				_ = _lhs7
				var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_27_v2), 0))
				_ = _lhs8
				var _lhs9 _dafny.Array = _27_v2
				_ = _lhs9
				var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_27_v2), 0))
				_ = _lhs10
				var _lhs11 *GlobalState = _47_globalState
				_ = _lhs11
				var _lhs12 _dafny.Array = _124_v75
				_ = _lhs12
				var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_124_v75), 0))
				_ = _lhs13
				var _lhs14 _dafny.Array = _27_v2
				_ = _lhs14
				var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_27_v2), 0))
				_ = _lhs15
				(_lhs7).ArraySet1(_rhs8, (_lhs8).Int())
				(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
				_lhs11.F5 = _rhs10
				(_lhs12).ArraySet1(_rhs11, (_lhs13).Int())
				(_lhs14).ArraySet1(_rhs12, (_lhs15).Int())
				(_47_globalState).F5 = (_50_v23.F27).Plus(_26_v1)
				var _129_v80 D13
				_ = _129_v80
				_129_v80 = Companion_D13_.Create_DC35_(_46_v21)
				_129_v80 = (func() D13 {
					if (Companion_D17_.Create_DC48_(_60_v30.F23, (_52_v25).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm42(_53_v26, (_27_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_27_v2), 0))).Int()).(bool), _47_globalState)).Dtor_cf10(), _dafny.IntOfUint32((_52_v25).Cardinality()))).Uint32()).(_dafny.Array))).Dtor_cf66() {
						return _129_v80
					}
					return Companion_D13_.Create_DC35_(_46_v21)
				})()
				var _130_v81 _dafny.Array
				_ = _130_v81
				var _len0_3 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_3
				var _nw24 _dafny.Array
				_ = _nw24
				if _len0_3.Cmp(_dafny.Zero) == 0 {
					_nw24 = _dafny.NewArray(_len0_3)
				} else {
					var _init3 func(_dafny.Int) bool = (func(_131_v2 _dafny.Array) func(_dafny.Int) bool {
						return func(_132_i7 _dafny.Int) bool {
							return (_131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_131_v2), 0))).Int()).(bool)
						}
					})(_27_v2)
					_ = _init3
					var _element0_3 = _init3(_dafny.Zero)
					_ = _element0_3
					_nw24 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
					(_nw24).ArraySet1(_element0_3, 0)
					var _nativeLen0_3 = (_len0_3).Int()
					_ = _nativeLen0_3
					for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
						(_nw24).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
					}
				}
				_130_v81 = _nw24
				_28_v3 = (_28_v3).Update(_50_v23.F27, _130_v81)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_dafny.Print(_25_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_26_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_v3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_29_v4).Equals(_dafny.SetOf(_dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_30_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.SetOf(_dafny.IntOfInt64(-60)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_31_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_32_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_33_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_34_v9).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_35_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_36_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_37_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(60), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_38_v16, _dafny.SeqOf(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_39_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(5), _dafny.SeqOf(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60)).UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-6), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-7), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-8), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-9), _dafny.IntOfInt64(898))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_40_v18).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_44_v19, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("mrqsc"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_45_v20).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_46_v21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60)).UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-6), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-7), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-8), _dafny.IntOfInt64(898)).UpdateUnsafe(_dafny.IntOfInt64(-9), _dafny.IntOfInt64(898))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState.F8).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F15.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_47_globalState).F16()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_47_globalState).F19().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_50_v23.F27)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v24).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_52_v25).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_53_v26)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_54_v27).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_58_v28).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v29).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-60), true), _dafny.IntOfInt64(-60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_v30.F23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_61_v31).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_62_v34).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v42).Dtor_cf46())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v42).Dtor_cf47())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v42).Dtor_cf48())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v69).Dtor_cf63())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v69).Dtor_cf64())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v69).Dtor_cf65())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_117_v72).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_v73).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_123_i6)
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

type D1_DC1 struct {
	Cf1 bool
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 bool) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() bool {
	return false
}

func (_this D1) Dtor_cf1() bool {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
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

type D2_DC3 struct {
	Cf3 _dafny.Int
	Cf4 bool
	Cf5 bool
	Cf6 _dafny.Int
	Cf7 bool
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf3 _dafny.Int, Cf4 bool, Cf5 bool, Cf6 _dafny.Int, Cf7 bool) D2 {
	return D2{D2_DC3{Cf3, Cf4, Cf5, Cf6, Cf7}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC2 struct {
	Cf2 _dafny.MultiSet
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 _dafny.MultiSet) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

type D2_DC4 struct {
	Cf8 D2
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf8 D2) D2 {
	return D2{D2_DC4{Cf8}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC3_(_dafny.Zero, false, false, _dafny.Zero, false)
}

func (_this D2) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) Dtor_cf4() bool {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC3).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC3).Cf7
}

func (_this D2) Dtor_cf2() _dafny.MultiSet {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) Dtor_cf8() D2 {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7
		}
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf10 _dafny.Int
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf10 _dafny.Int) D3 {
	return D3{D3_DC6{Cf10}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC7 struct {
	Cf11 _dafny.Sequence
	Cf12 bool
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Sequence, Cf12 bool) D3 {
	return D3{D3_DC7{Cf11, Cf12}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC5 struct {
	Cf9 _dafny.Sequence
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf9 _dafny.Sequence) D3 {
	return D3{D3_DC5{Cf9}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

type D3_DC8 struct {
	Cf13 D3
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf13 D3) D3 {
	return D3{D3_DC8{Cf13}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.Zero)
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf10
}

func (_this D3) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC7).Cf12
}

func (_this D3) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D3_DC5).Cf9
}

func (_this D3) Dtor_cf13() D3 {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11) && data1.Cf12 == data2.Cf12
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf15 _dafny.MultiSet
	Cf16 bool
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf15 _dafny.MultiSet, Cf16 bool) D4 {
	return D4{D4_DC10{Cf15, Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf17 _dafny.Int
	Cf18 bool
	Cf19 bool
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.Int, Cf18 bool, Cf19 bool) D4 {
	return D4{D4_DC11{Cf17, Cf18, Cf19}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC9 struct {
	Cf14 _dafny.Array
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf14 _dafny.Array) D4 {
	return D4{D4_DC9{Cf14}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.EmptyMultiSet, false)
}

func (_this D4) Dtor_cf15() _dafny.MultiSet {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf19() bool {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15) && data1.Cf16 == data2.Cf16
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf14 == data2.Cf14
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

type D5_DC12 struct {
	Cf20 _dafny.Array
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf20 _dafny.Array) D5 {
	return D5{D5_DC12{Cf20}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D5_DC12).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf20 == data2.Cf20
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

type D6_DC14 struct {
	Cf22 _dafny.Int
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf22 _dafny.Int) D6 {
	return D6{D6_DC14{Cf22}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC15 struct {
	Cf23 _dafny.Array
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.Array) D6 {
	return D6{D6_DC15{Cf23}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC13 struct {
	Cf21 _dafny.Sequence
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf21 _dafny.Sequence) D6 {
	return D6{D6_DC13{Cf21}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.Zero)
}

func (_this D6) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Array {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D6_DC13).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + data.Cf21.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
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

type D7_DC17 struct {
	Cf25 bool
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf25 bool) D7 {
	return D7{D7_DC17{Cf25}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC18 struct {
	Cf26 _dafny.Sequence
	Cf27 _dafny.Int
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf26 _dafny.Sequence, Cf27 _dafny.Int) D7 {
	return D7{D7_DC18{Cf26, Cf27}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC16 struct {
	Cf24 _dafny.Set
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf24 _dafny.Set) D7 {
	return D7{D7_DC16{Cf24}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC17_(false)
}

func (_this D7) Dtor_cf25() bool {
	return _this.Get_().(D7_DC17).Cf25
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC18).Cf26
}

func (_this D7) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D7_DC18).Cf27
}

func (_this D7) Dtor_cf24() _dafny.Set {
	return _this.Get_().(D7_DC16).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf25 == data2.Cf25
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26) && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_() D8 {
	return D8{D8_DC20{}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC21 struct {
	Cf29 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf29 _dafny.Int) D8 {
	return D8{D8_DC21{Cf29}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC19 struct {
	Cf28 _dafny.Map
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf28 _dafny.Map) D8 {
	return D8{D8_DC19{Cf28}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_()
}

func (_this D8) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf29
}

func (_this D8) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D8_DC19).Cf28
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf28) + ")"
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
			_, ok := other.Get_().(D8_DC20)
			return ok
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
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

type D9_DC23 struct {
	Cf31 _dafny.Int
	Cf32 bool
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf31 _dafny.Int, Cf32 bool) D9 {
	return D9{D9_DC23{Cf31, Cf32}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC22 struct {
	Cf30 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf30 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf30}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC24 struct {
	Cf33 D9
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf33 D9) D9 {
	return D9{D9_DC24{Cf33}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(_dafny.Zero, false)
}

func (_this D9) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf31
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC23).Cf32
}

func (_this D9) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf30
}

func (_this D9) Dtor_cf33() D9 {
	return _this.Get_().(D9_DC24).Cf33
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

type D10_DC26 struct {
	Cf35 _dafny.Int
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf35 _dafny.Int) D10 {
	return D10{D10_DC26{Cf35}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC27 struct {
	Cf36 _dafny.Int
	Cf37 _dafny.Sequence
	Cf38 _dafny.CodePoint
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf36 _dafny.Int, Cf37 _dafny.Sequence, Cf38 _dafny.CodePoint) D10 {
	return D10{D10_DC27{Cf36, Cf37, Cf38}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC25 struct {
	Cf34 _dafny.MultiSet
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf34 _dafny.MultiSet) D10 {
	return D10{D10_DC25{Cf34}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_(_dafny.Zero)
}

func (_this D10) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D10_DC26).Cf35
}

func (_this D10) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf36
}

func (_this D10) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D10_DC27).Cf37
}

func (_this D10) Dtor_cf38() _dafny.CodePoint {
	return _this.Get_().(D10_DC27).Cf38
}

func (_this D10) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D10_DC25).Cf34
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Equals(data2.Cf37) && data1.Cf38 == data2.Cf38
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
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

type D11_DC29 struct {
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_() D11 {
	return D11{D11_DC29{}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC30 struct {
	Cf40 bool
	Cf41 _dafny.MultiSet
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf40 bool, Cf41 _dafny.MultiSet) D11 {
	return D11{D11_DC30{Cf40, Cf41}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC28 struct {
	Cf39 _dafny.Array
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf39 _dafny.Array) D11 {
	return D11{D11_DC28{Cf39}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC31 struct {
	Cf42 D11
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf42 D11) D11 {
	return D11{D11_DC31{Cf42}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC29_()
}

func (_this D11) Dtor_cf40() bool {
	return _this.Get_().(D11_DC30).Cf40
}

func (_this D11) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D11_DC30).Cf41
}

func (_this D11) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D11_DC28).Cf39
}

func (_this D11) Dtor_cf42() D11 {
	return _this.Get_().(D11_DC31).Cf42
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC29:
		{
			return "D11.DC29"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC29:
		{
			_, ok := other.Get_().(D11_DC29)
			return ok
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41.Equals(data2.Cf41)
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf39 == data2.Cf39
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

type D12_DC33 struct {
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_() D12 {
	return D12{D12_DC33{}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf43 _dafny.Sequence
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf43 _dafny.Sequence) D12 {
	return D12{D12_DC32{Cf43}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC34 struct {
	Cf44 D12
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf44 D12) D12 {
	return D12{D12_DC34{Cf44}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_()
}

func (_this D12) Dtor_cf43() _dafny.Sequence {
	return _this.Get_().(D12_DC32).Cf43
}

func (_this D12) Dtor_cf44() D12 {
	return _this.Get_().(D12_DC34).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC33:
		{
			_, ok := other.Get_().(D12_DC33)
			return ok
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D13_DC36 struct {
	Cf46 _dafny.Int
	Cf47 bool
	Cf48 bool
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf46 _dafny.Int, Cf47 bool, Cf48 bool) D13 {
	return D13{D13_DC36{Cf46, Cf47, Cf48}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC37 struct {
	Cf49 bool
	Cf50 _dafny.Int
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf49 bool, Cf50 _dafny.Int) D13 {
	return D13{D13_DC37{Cf49, Cf50}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

type D13_DC35 struct {
	Cf45 _dafny.CodePoint
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf45 _dafny.CodePoint) D13 {
	return D13{D13_DC35{Cf45}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC36_(_dafny.Zero, false, false)
}

func (_this D13) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D13_DC36).Cf46
}

func (_this D13) Dtor_cf47() bool {
	return _this.Get_().(D13_DC36).Cf47
}

func (_this D13) Dtor_cf48() bool {
	return _this.Get_().(D13_DC36).Cf48
}

func (_this D13) Dtor_cf49() bool {
	return _this.Get_().(D13_DC37).Cf49
}

func (_this D13) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D13_DC37).Cf50
}

func (_this D13) Dtor_cf45() _dafny.CodePoint {
	return _this.Get_().(D13_DC35).Cf45
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48
		}
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf45 == data2.Cf45
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

type D14_DC39 struct {
}

func (D14_DC39) isD14() {}

func (CompanionStruct_D14_) Create_DC39_() D14 {
	return D14{D14_DC39{}}
}

func (_this D14) Is_DC39() bool {
	_, ok := _this.Get_().(D14_DC39)
	return ok
}

type D14_DC38 struct {
	Cf51 _dafny.Array
}

func (D14_DC38) isD14() {}

func (CompanionStruct_D14_) Create_DC38_(Cf51 _dafny.Array) D14 {
	return D14{D14_DC38{Cf51}}
}

func (_this D14) Is_DC38() bool {
	_, ok := _this.Get_().(D14_DC38)
	return ok
}

type D14_DC40 struct {
	Cf52 D14
}

func (D14_DC40) isD14() {}

func (CompanionStruct_D14_) Create_DC40_(Cf52 D14) D14 {
	return D14{D14_DC40{Cf52}}
}

func (_this D14) Is_DC40() bool {
	_, ok := _this.Get_().(D14_DC40)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC39_()
}

func (_this D14) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D14_DC38).Cf51
}

func (_this D14) Dtor_cf52() D14 {
	return _this.Get_().(D14_DC40).Cf52
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC39:
		{
			return "D14.DC39"
		}
	case D14_DC38:
		{
			return "D14.DC38" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D14_DC40:
		{
			return "D14.DC40" + "(" + _dafny.String(data.Cf52) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC39:
		{
			_, ok := other.Get_().(D14_DC39)
			return ok
		}
	case D14_DC38:
		{
			data2, ok := other.Get_().(D14_DC38)
			return ok && data1.Cf51 == data2.Cf51
		}
	case D14_DC40:
		{
			data2, ok := other.Get_().(D14_DC40)
			return ok && data1.Cf52.Equals(data2.Cf52)
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

type D15_DC41 struct {
	Cf53 *C1
}

func (D15_DC41) isD15() {}

func (CompanionStruct_D15_) Create_DC41_(Cf53 *C1) D15 {
	return D15{D15_DC41{Cf53}}
}

func (_this D15) Is_DC41() bool {
	_, ok := _this.Get_().(D15_DC41)
	return ok
}

func (CompanionStruct_D15_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D15) Dtor_cf53() *C1 {
	return _this.Get_().(D15_DC41).Cf53
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC41:
		{
			return "D15.DC41" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC41:
		{
			data2, ok := other.Get_().(D15_DC41)
			return ok && data1.Cf53 == data2.Cf53
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

type D16_DC43 struct {
	Cf55 _dafny.Int
	Cf56 _dafny.MultiSet
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf55 _dafny.Int, Cf56 _dafny.MultiSet) D16 {
	return D16{D16_DC43{Cf55, Cf56}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

type D16_DC44 struct {
	Cf57 _dafny.Int
	Cf58 bool
	Cf59 _dafny.Int
	Cf60 _dafny.Array
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf57 _dafny.Int, Cf58 bool, Cf59 _dafny.Int, Cf60 _dafny.Array) D16 {
	return D16{D16_DC44{Cf57, Cf58, Cf59, Cf60}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

type D16_DC45 struct {
	Cf61 bool
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf61 bool) D16 {
	return D16{D16_DC45{Cf61}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC42 struct {
	Cf54 _dafny.Map
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf54 _dafny.Map) D16 {
	return D16{D16_DC42{Cf54}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC43_(_dafny.Zero, _dafny.EmptyMultiSet)
}

func (_this D16) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D16_DC43).Cf55
}

func (_this D16) Dtor_cf56() _dafny.MultiSet {
	return _this.Get_().(D16_DC43).Cf56
}

func (_this D16) Dtor_cf57() _dafny.Int {
	return _this.Get_().(D16_DC44).Cf57
}

func (_this D16) Dtor_cf58() bool {
	return _this.Get_().(D16_DC44).Cf58
}

func (_this D16) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D16_DC44).Cf59
}

func (_this D16) Dtor_cf60() _dafny.Array {
	return _this.Get_().(D16_DC44).Cf60
}

func (_this D16) Dtor_cf61() bool {
	return _this.Get_().(D16_DC45).Cf61
}

func (_this D16) Dtor_cf54() _dafny.Map {
	return _this.Get_().(D16_DC42).Cf54
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf61) + ")"
		}
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf55.Cmp(data2.Cf55) == 0 && data1.Cf56.Equals(data2.Cf56)
		}
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf57.Cmp(data2.Cf57) == 0 && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60 == data2.Cf60
		}
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf61 == data2.Cf61
		}
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D17_DC47 struct {
	Cf63 _dafny.Int
	Cf64 bool
	Cf65 _dafny.Int
}

func (D17_DC47) isD17() {}

func (CompanionStruct_D17_) Create_DC47_(Cf63 _dafny.Int, Cf64 bool, Cf65 _dafny.Int) D17 {
	return D17{D17_DC47{Cf63, Cf64, Cf65}}
}

func (_this D17) Is_DC47() bool {
	_, ok := _this.Get_().(D17_DC47)
	return ok
}

type D17_DC48 struct {
	Cf66 bool
	Cf67 _dafny.Array
}

func (D17_DC48) isD17() {}

func (CompanionStruct_D17_) Create_DC48_(Cf66 bool, Cf67 _dafny.Array) D17 {
	return D17{D17_DC48{Cf66, Cf67}}
}

func (_this D17) Is_DC48() bool {
	_, ok := _this.Get_().(D17_DC48)
	return ok
}

type D17_DC49 struct {
	Cf68 bool
}

func (D17_DC49) isD17() {}

func (CompanionStruct_D17_) Create_DC49_(Cf68 bool) D17 {
	return D17{D17_DC49{Cf68}}
}

func (_this D17) Is_DC49() bool {
	_, ok := _this.Get_().(D17_DC49)
	return ok
}

type D17_DC46 struct {
	Cf62 _dafny.Sequence
}

func (D17_DC46) isD17() {}

func (CompanionStruct_D17_) Create_DC46_(Cf62 _dafny.Sequence) D17 {
	return D17{D17_DC46{Cf62}}
}

func (_this D17) Is_DC46() bool {
	_, ok := _this.Get_().(D17_DC46)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC47_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D17) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf63
}

func (_this D17) Dtor_cf64() bool {
	return _this.Get_().(D17_DC47).Cf64
}

func (_this D17) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf65
}

func (_this D17) Dtor_cf66() bool {
	return _this.Get_().(D17_DC48).Cf66
}

func (_this D17) Dtor_cf67() _dafny.Array {
	return _this.Get_().(D17_DC48).Cf67
}

func (_this D17) Dtor_cf68() bool {
	return _this.Get_().(D17_DC49).Cf68
}

func (_this D17) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D17_DC46).Cf62
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC47:
		{
			return "D17.DC47" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D17_DC48:
		{
			return "D17.DC48" + "(" + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D17_DC49:
		{
			return "D17.DC49" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D17_DC46:
		{
			return "D17.DC46" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC47:
		{
			data2, ok := other.Get_().(D17_DC47)
			return ok && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64 == data2.Cf64 && data1.Cf65.Cmp(data2.Cf65) == 0
		}
	case D17_DC48:
		{
			data2, ok := other.Get_().(D17_DC48)
			return ok && data1.Cf66 == data2.Cf66 && data1.Cf67 == data2.Cf67
		}
	case D17_DC49:
		{
			data2, ok := other.Get_().(D17_DC49)
			return ok && data1.Cf68 == data2.Cf68
		}
	case D17_DC46:
		{
			data2, ok := other.Get_().(D17_DC46)
			return ok && data1.Cf62.Equals(data2.Cf62)
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

type D18_DC51 struct {
}

func (D18_DC51) isD18() {}

func (CompanionStruct_D18_) Create_DC51_() D18 {
	return D18{D18_DC51{}}
}

func (_this D18) Is_DC51() bool {
	_, ok := _this.Get_().(D18_DC51)
	return ok
}

type D18_DC52 struct {
	Cf70 D10
}

func (D18_DC52) isD18() {}

func (CompanionStruct_D18_) Create_DC52_(Cf70 D10) D18 {
	return D18{D18_DC52{Cf70}}
}

func (_this D18) Is_DC52() bool {
	_, ok := _this.Get_().(D18_DC52)
	return ok
}

type D18_DC50 struct {
	Cf69 _dafny.Array
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf69 _dafny.Array) D18 {
	return D18{D18_DC50{Cf69}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

type D18_DC53 struct {
	Cf71 D18
}

func (D18_DC53) isD18() {}

func (CompanionStruct_D18_) Create_DC53_(Cf71 D18) D18 {
	return D18{D18_DC53{Cf71}}
}

func (_this D18) Is_DC53() bool {
	_, ok := _this.Get_().(D18_DC53)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC51_()
}

func (_this D18) Dtor_cf70() D10 {
	return _this.Get_().(D18_DC52).Cf70
}

func (_this D18) Dtor_cf69() _dafny.Array {
	return _this.Get_().(D18_DC50).Cf69
}

func (_this D18) Dtor_cf71() D18 {
	return _this.Get_().(D18_DC53).Cf71
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC51:
		{
			return "D18.DC51"
		}
	case D18_DC52:
		{
			return "D18.DC52" + "(" + _dafny.String(data.Cf70) + ")"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf69) + ")"
		}
	case D18_DC53:
		{
			return "D18.DC53" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC51:
		{
			_, ok := other.Get_().(D18_DC51)
			return ok
		}
	case D18_DC52:
		{
			data2, ok := other.Get_().(D18_DC52)
			return ok && data1.Cf70.Equals(data2.Cf70)
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf69 == data2.Cf69
		}
	case D18_DC53:
		{
			data2, ok := other.Get_().(D18_DC53)
			return ok && data1.Cf71.Equals(data2.Cf71)
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

type D19_DC55 struct {
	Cf73 _dafny.Int
	Cf74 _dafny.Int
}

func (D19_DC55) isD19() {}

func (CompanionStruct_D19_) Create_DC55_(Cf73 _dafny.Int, Cf74 _dafny.Int) D19 {
	return D19{D19_DC55{Cf73, Cf74}}
}

func (_this D19) Is_DC55() bool {
	_, ok := _this.Get_().(D19_DC55)
	return ok
}

type D19_DC54 struct {
	Cf72 _dafny.Map
}

func (D19_DC54) isD19() {}

func (CompanionStruct_D19_) Create_DC54_(Cf72 _dafny.Map) D19 {
	return D19{D19_DC54{Cf72}}
}

func (_this D19) Is_DC54() bool {
	_, ok := _this.Get_().(D19_DC54)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC55_(_dafny.Zero, _dafny.Zero)
}

func (_this D19) Dtor_cf73() _dafny.Int {
	return _this.Get_().(D19_DC55).Cf73
}

func (_this D19) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D19_DC55).Cf74
}

func (_this D19) Dtor_cf72() _dafny.Map {
	return _this.Get_().(D19_DC54).Cf72
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC55:
		{
			return "D19.DC55" + "(" + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D19_DC54:
		{
			return "D19.DC54" + "(" + _dafny.String(data.Cf72) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC55:
		{
			data2, ok := other.Get_().(D19_DC55)
			return ok && data1.Cf73.Cmp(data2.Cf73) == 0 && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D19_DC54:
		{
			data2, ok := other.Get_().(D19_DC54)
			return ok && data1.Cf72.Equals(data2.Cf72)
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

type D20_DC56 struct {
	Cf75 _dafny.Map
}

func (D20_DC56) isD20() {}

func (CompanionStruct_D20_) Create_DC56_(Cf75 _dafny.Map) D20 {
	return D20{D20_DC56{Cf75}}
}

func (_this D20) Is_DC56() bool {
	_, ok := _this.Get_().(D20_DC56)
	return ok
}

func (CompanionStruct_D20_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D20) Dtor_cf75() _dafny.Map {
	return _this.Get_().(D20_DC56).Cf75
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC56:
		{
			return "D20.DC56" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC56:
		{
			data2, ok := other.Get_().(D20_DC56)
			return ok && data1.Cf75.Equals(data2.Cf75)
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

type D21_DC57 struct {
	Cf76 _dafny.Map
}

func (D21_DC57) isD21() {}

func (CompanionStruct_D21_) Create_DC57_(Cf76 _dafny.Map) D21 {
	return D21{D21_DC57{Cf76}}
}

func (_this D21) Is_DC57() bool {
	_, ok := _this.Get_().(D21_DC57)
	return ok
}

func (CompanionStruct_D21_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D21) Dtor_cf76() _dafny.Map {
	return _this.Get_().(D21_DC57).Cf76
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC57:
		{
			return "D21.DC57" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC57:
		{
			data2, ok := other.Get_().(D21_DC57)
			return ok && data1.Cf76.Equals(data2.Cf76)
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
	Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool
	Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool)
	F20() _dafny.Int
	F21() bool
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

// Definition of class GlobalState
type GlobalState struct {
	F3   bool
	F4   bool
	F5   _dafny.Int
	F8   _dafny.Array
	F14  bool
	F15  _dafny.Sequence
	F17  _dafny.Int
	F18  bool
	_f0  _dafny.Int
	_f1  _dafny.Sequence
	_f2  bool
	_f6  _dafny.Array
	_f7  bool
	_f9  _dafny.Int
	_f10 bool
	_f11 _dafny.Int
	_f12 bool
	_f13 _dafny.Int
	_f16 _dafny.Map
	_f19 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = false
	_this.F4 = false
	_this.F5 = _dafny.Zero
	_this.F8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F14 = false
	_this.F15 = _dafny.EmptySeq
	_this.F17 = _dafny.Zero
	_this.F18 = false
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f2 = false
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f7 = false
	_this._f9 = _dafny.Zero
	_this._f10 = false
	_this._f11 = _dafny.Zero
	_this._f12 = false
	_this._f13 = _dafny.Zero
	_this._f16 = _dafny.EmptyMap
	_this._f19 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 bool, f3 bool, f4 bool, f5 _dafny.Int, f6 _dafny.Array, f7 bool, f8 _dafny.Array, f9 _dafny.Int, f10 bool, f11 _dafny.Int, f12 bool, f13 _dafny.Int, f14 bool, f15 _dafny.Sequence, f16 _dafny.Map, f17 _dafny.Int, f18 bool, f19 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this).F18 = f18
		(_this)._f19 = f19
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
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
func (_this *GlobalState) F6() _dafny.Array {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
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
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F16() _dafny.Map {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F19() _dafny.Sequence {
	{
		return _this._f19
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
func (_this *C0) Fm3(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	{
		return !(false)
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(!(true))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f20 _dafny.Int
	_f21 bool
	F27  _dafny.Int
	_f28 _dafny.Array
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
	_this.F27 = _dafny.Zero
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C1) F20() _dafny.Int {
	return _this._f20
}
func (_this *C1) F21() bool {
	return _this._f21
}
func (_this *C1) Ctor__(f27 _dafny.Int, f28 _dafny.Array, f20 _dafny.Int, f21 bool) {
	{
		(_this).F27 = f27
		(_this)._f28 = f28
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C1) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F21()
	}
}
func (_this *C1) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F27
	}
}
func (_this *C1) Fm16(p0 _dafny.MultiSet, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Set {
			if (_this).F21() {
				return _dafny.SetOf((_this).F21(), (_this).F21(), (_this).F21(), (_this).F21(), (_this).F21())
			}
			return _dafny.SetOf((_this).F21(), (_this).F21(), (_this).F21())
		})()).IsProperSubsetOf((_dafny.SetOf((_this).F21())).Difference(_dafny.SetOf((_this).F21())))
	}
}
func (_this *C1) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _137_v0 _dafny.MultiSet
		_ = _137_v0
		_137_v0 = _dafny.MultiSetOf(_this.F27)
		_137_v0 = _dafny.MultiSetOf(_this.F27)
		var _138_i0 _dafny.Int
		_ = _138_i0
		_138_i0 = _dafny.Zero
		{
			for !(p1) || ((_this.F27).Cmp((_this).F20()) <= 0) {
				{
					if (_138_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_138_i0 = (_138_i0).Plus(_dafny.One)
					(globalState).F18 = p1
					var _139_v1 _dafny.Map
					_ = _139_v1
					_139_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())
					var _140_v2 _dafny.Map
					_ = _140_v2
					_140_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), !(p1))
					var _141_v3 _dafny.MultiSet
					_ = _141_v3
					_141_v3 = _dafny.MultiSetOf((func() bool {
						if (_140_v2).Contains(p1) {
							return (_140_v2).Get(p1).(bool)
						}
						return (_this).F21()
					})())
					_139_v1 = (_139_v1).Update(((_141_v3).Cardinality()).Times((_this).F20()), (_this).F28())
					(globalState).F5 = (_this.F27).Plus(p0)
					(globalState).F5 = (func() _dafny.Int {
						if (_this).F21() {
							return (_this).F20()
						}
						return p0
					})()
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _142_v4 _dafny.MultiSet
		_ = _142_v4
		_142_v4 = _dafny.MultiSetOf((_this).F21())
		var _143_v5 _dafny.Sequence
		_ = _143_v5
		_143_v5 = _dafny.SeqOf(p0, (_137_v0).Cardinality())
		var _144_v6 _dafny.Map
		_ = _144_v6
		_144_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _143_v5)
		if (_this).Fm16(_142_v4, (func() _dafny.Sequence {
			if (_144_v6).Contains(_dafny.IntOfInt64(824)) {
				return (_144_v6).Get(_dafny.IntOfInt64(824)).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-961))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_145_i1 _dafny.Int) _dafny.Int {
				return _this.F27
			}))
		})(), globalState) {
			var _146_v7 D4
			_ = _146_v7
			_146_v7 = Companion_D4_.Create_DC11_((_this).F20(), (_this).F21(), (_this).F21())
			var _source5 D4 = _146_v7
			_ = _source5
			if _source5.Is_DC10() {
				var _147___mcc_h0 _dafny.MultiSet = _source5.Get_().(D4_DC10).Cf15
				_ = _147___mcc_h0
				var _148___mcc_h1 bool = _source5.Get_().(D4_DC10).Cf16
				_ = _148___mcc_h1
				var _149_cf16 bool = _148___mcc_h1
				_ = _149_cf16
				var _150_cf15 _dafny.MultiSet = _147___mcc_h0
				_ = _150_cf15
				var _151_v8 _dafny.Array
				_ = _151_v8
				var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw25
				_151_v8 = _nw25
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_151_v8), 0))
				_ = _index18
				(_151_v8).ArraySet1((_this).F28(), (_index18).Int())
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_151_v8), 0))
				_ = _index19
				(_151_v8).ArraySet1((_this).F28(), (_index19).Int())
				var _152_v9 *C0
				_ = _152_v9
				var _nw26 *C0 = New_C0_()
				_ = _nw26
				_nw26.Ctor__()
				_152_v9 = _nw26
				(globalState).F3 = p1
				(globalState).F14 = ((_this).F20()).Cmp(_this.F27) <= 0
			} else if _source5.Is_DC11() {
				var _153___mcc_h2 _dafny.Int = _source5.Get_().(D4_DC11).Cf17
				_ = _153___mcc_h2
				var _154___mcc_h3 bool = _source5.Get_().(D4_DC11).Cf18
				_ = _154___mcc_h3
				var _155___mcc_h4 bool = _source5.Get_().(D4_DC11).Cf19
				_ = _155___mcc_h4
				var _156_cf19 bool = _155___mcc_h4
				_ = _156_cf19
				var _157_cf18 bool = _154___mcc_h3
				_ = _157_cf18
				var _158_cf17 _dafny.Int = _153___mcc_h2
				_ = _158_cf17
				(globalState).F14 = p1
				var _159_v10 *C0
				_ = _159_v10
				var _nw27 *C0 = New_C0_()
				_ = _nw27
				_nw27.Ctor__()
				_159_v10 = _nw27
				(globalState).F18 = !(_156_cf19)
				(globalState).F18 = (_this).F21()
			} else {
				var _160___mcc_h5 _dafny.Array = _source5.Get_().(D4_DC9).Cf14
				_ = _160___mcc_h5
				var _161_cf14 _dafny.Array = _160___mcc_h5
				_ = _161_cf14
				var _162_v11 _dafny.Map
				_ = _162_v11
				_162_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, (_this).F28())
				var _163_v12 _dafny.Sequence
				_ = _163_v12
				_163_v12 = _dafny.SeqOf((func() _dafny.Array {
					if (_162_v11).Contains(_dafny.IntOfInt64(35)) {
						return (_162_v11).Get(_dafny.IntOfInt64(35)).(_dafny.Array)
					}
					return (_this).F28()
				})())
				_163_v12 = _163_v12
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_161_cf14), 0))
				_ = _index20
				(_161_cf14).ArraySet1(p1, (_index20).Int())
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_161_cf14), 0))
				_ = _index21
				(_161_cf14).ArraySet1(!((_this).F21()) || (p1), (_index21).Int())
				var _164_v13 _dafny.Array
				_ = _164_v13
				var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
				_ = _nw28
				_164_v13 = _nw28
				var _165_v14 _dafny.Sequence
				_ = _165_v14
				_165_v14 = _dafny.UnicodeSeqOfUtf8Bytes("mgaj")
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_164_v13), 0))
				_ = _index22
				(_164_v13).ArraySet1(_165_v14, (_index22).Int())
				var _166_v15 _dafny.CodePoint
				_ = _166_v15
				_166_v15 = _dafny.CodePoint('k')
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_164_v13), 0))
				_ = _index23
				(_164_v13).ArraySet1((func() _dafny.Sequence {
					if true {
						return _165_v14
					}
					return _dafny.Companion_Sequence_.Update(_165_v14, (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_165_v14).Cardinality()))).Uint32(), _166_v15)
				})(), (_index23).Int())
				var _167_v16 D6
				_ = _167_v16
				_167_v16 = Companion_D6_.Create_DC13_(_165_v14)
				_165_v14 = (_167_v16).Dtor_cf21()
			}
			var _168_v17 _dafny.Array
			_ = _168_v17
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw29
			_168_v17 = _nw29
			_168_v17 = _168_v17
			(globalState).F3 = p1
			(globalState).F4 = !(false) || (p1)
			var _169_v18 _dafny.Map
			_ = _169_v18
			_169_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (p0).Times(_this.F27))
			var _170_v19 _dafny.Sequence
			_ = _170_v19
			_170_v19 = _dafny.UnicodeSeqOfUtf8Bytes("e")
			_169_v18 = (_169_v18).Update((_this).F28(), (Companion_D2_.Create_DC3_((func() _dafny.Int {
				if (_137_v0).Contains(p0) {
					return (_137_v0).Multiplicity(p0)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_170_v19).Cardinality()))
			})(), (_this).F21(), !((_this).F21()), (_dafny.Zero).Minus((_this).F20()), p1)).Dtor_cf6())
		} else {
			(globalState).F18 = (_this).F21()
			var _171_v20 _dafny.Array
			_ = _171_v20
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw30
			_171_v20 = _nw30
			var _172_v21 _dafny.Sequence
			_ = _172_v21
			_172_v21 = _dafny.SeqOf(_171_v20, _171_v20)
			var _173_v22 _dafny.Array
			_ = _173_v22
			var _nwElement0_4 _dafny.Array = (func() _dafny.Array {
				if true {
					return _171_v20
				}
				return _171_v20
			})()
			_ = _nwElement0_4
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(7))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_4, 0)
			(_nw31).ArraySet1(_171_v20, 1)
			(_nw31).ArraySet1(_171_v20, 2)
			(_nw31).ArraySet1(_171_v20, 3)
			(_nw31).ArraySet1(_171_v20, 4)
			(_nw31).ArraySet1(_171_v20, 5)
			(_nw31).ArraySet1((_172_v21).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_172_v21).Cardinality()))).Uint32()).(_dafny.Array), 6)
			_173_v22 = _nw31
			var _174_v23 _dafny.Sequence
			_ = _174_v23
			_174_v23 = _dafny.SeqOf(_173_v22)
			_173_v22 = (_174_v23).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_174_v23).Cardinality()))).Uint32()).(_dafny.Array)
			(_this).F27 = _this.F27
			(globalState).F17 = p0
			var _175_v24 _dafny.CodePoint
			_ = _175_v24
			_175_v24 = _dafny.CodePoint('j')
			var _176_v25 _dafny.MultiSet
			_ = _176_v25
			_176_v25 = _dafny.MultiSetOf(_175_v24, _dafny.CodePoint('l'), _175_v24, _175_v24)
			_176_v25 = _dafny.MultiSetOf(_175_v24, _175_v24, Companion_Default___.Fm17(!((_this).F21()), globalState), _175_v24, _175_v24)
		}
		var _177_v26 _dafny.Sequence
		_ = _177_v26
		_177_v26 = _dafny.UnicodeSeqOfUtf8Bytes("ogosrhy")
		(globalState).F15 = _dafny.Companion_Sequence_.Concatenate(_177_v26, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(531))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_178_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _177_v26))
		var _179_v27 _dafny.Array
		_ = _179_v27
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw32
		_179_v27 = _nw32
		var _180_v28 _dafny.Set
		_ = _180_v28
		_180_v28 = _dafny.SetOf((_this).F21(), (_this).F21())
		var _181_v29 _dafny.Sequence
		_ = _181_v29
		_181_v29 = _dafny.SeqOf(_180_v28, _180_v28)
		var _182_v30 _dafny.Map
		_ = _182_v30
		_182_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F27)
		var _183_v31 _dafny.CodePoint
		_ = _183_v31
		_183_v31 = _dafny.CodePoint('m')
		var _184_v32 _dafny.Array
		_ = _184_v32
		var _nwElement0_5 _dafny.Set = _180_v28
		_ = _nwElement0_5
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_5, 0)
		(_nw33).ArraySet1(_180_v28, 1)
		(_nw33).ArraySet1(_180_v28, 2)
		(_nw33).ArraySet1((_181_v29).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_181_v29).Cardinality()))).Uint32()).(_dafny.Set), 3)
		(_nw33).ArraySet1(_180_v28, 4)
		(_nw33).ArraySet1(_dafny.SetOf(!(p1), !((_this).F21())), 5)
		(_nw33).ArraySet1(_180_v28, 6)
		(_nw33).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()), (func() _dafny.Int {
			if (_182_v30).Contains(p1) {
				return (_182_v30).Get(p1).(_dafny.Int)
			}
			return _dafny.IntOfInt64(221)
		})(), _183_v31, globalState), 7)
		(_nw33).ArraySet1(_180_v28, 8)
		(_nw33).ArraySet1(_180_v28, 9)
		(_nw33).ArraySet1(_180_v28, 10)
		(_nw33).ArraySet1((Companion_D7_.Create_DC16_(_180_v28)).Dtor_cf24(), 11)
		(_nw33).ArraySet1(_180_v28, 12)
		(_nw33).ArraySet1(_dafny.SetOf(p1, (_this).F21()), 13)
		(_nw33).ArraySet1(_180_v28, 14)
		(_nw33).ArraySet1(_180_v28, 15)
		(_nw33).ArraySet1(_180_v28, 16)
		(_nw33).ArraySet1(_180_v28, 17)
		(_nw33).ArraySet1(_180_v28, 18)
		(_nw33).ArraySet1(_180_v28, 19)
		(_nw33).ArraySet1(_dafny.SetOf(true, p1), 20)
		(_nw33).ArraySet1(_180_v28, 21)
		(_nw33).ArraySet1(_180_v28, 22)
		(_nw33).ArraySet1(_180_v28, 23)
		(_nw33).ArraySet1(Companion_Default___.Fm18(p0, p0, _183_v31, globalState), 24)
		_184_v32 = _nw33
		var _185_v33 _dafny.Map
		_ = _185_v33
		_185_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v27, _184_v32)
		_185_v33 = (_185_v33).Update(_179_v27, _184_v32)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_179_v27), 0))
		_ = _index24
		(_179_v27).ArraySet1(p1, (_index24).Int())
		var _186_v34 _dafny.Sequence
		_ = _186_v34
		_186_v34 = _dafny.SeqOf(p1)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_179_v27), 0))
		_ = _index25
		(_179_v27).ArraySet1(((_186_v34).Select((Companion_Default___.SafeIndex((_182_v30).Cardinality(), _dafny.IntOfUint32((_186_v34).Cardinality()))).Uint32()).(bool)) && ((_this).F21()), (_index25).Int())
		var _187_v35 _dafny.Sequence
		_ = _187_v35
		_187_v35 = _dafny.SeqOf((_this).F28(), (_this).F28())
		var _188_v36 _dafny.Sequence
		_ = _188_v36
		_188_v36 = _dafny.SeqOf(_187_v35)
		r0 = _dafny.Companion_Sequence_.Concatenate((_188_v36).Select((Companion_Default___.SafeIndex((_143_v5).Select((Companion_Default___.SafeIndex((_180_v28).Cardinality(), _dafny.IntOfUint32((_143_v5).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_188_v36).Cardinality()))).Uint32()).(_dafny.Sequence), _187_v35)
		r1 = p1
		return r0, r1
	}
}
func (_this *C1) F28() _dafny.Array {
	{
		return _this._f28
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f20 _dafny.Int
	_f21 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
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

func (_this *C2) F20() _dafny.Int {
	return _this._f20
}
func (_this *C2) F21() bool {
	return _this._f21
}
func (_this *C2) Ctor__(f20 _dafny.Int, f21 bool) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C2) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_this).F20(), (_this).F20()), (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf((_this).F21(), (_this).F21(), true, (_this).F21(), (_this).F21())).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(189))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_189_i0 _dafny.Int) _dafny.Int {
			return (_this).F20()
		}))).Cardinality())))))
	}
}
func (_this *C2) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F20()
	}
}
func (_this *C2) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _190_v0 _dafny.Sequence
		_ = _190_v0
		_190_v0 = _dafny.SeqOf(true)
		var _191_v1 _dafny.Int
		_ = _191_v1
		var _192_v2 bool
		_ = _192_v2
		var _193_v3 _dafny.Int
		_ = _193_v3
		var _194_v4 _dafny.Int
		_ = _194_v4
		var _out2 _dafny.Int
		_ = _out2
		var _out3 bool
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		var _out5 _dafny.Int
		_ = _out5
		_out2, _out3, _out4, _out5 = (_this).M3(_dafny.Companion_Sequence_.Equal(_190_v0, _190_v0), globalState)
		_191_v1 = _out2
		_192_v2 = _out3
		_193_v3 = _out4
		_194_v4 = _out5
		var _195_v5 _dafny.Set
		_ = _195_v5
		_195_v5 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _191_v1)).Cardinality(), (_this).F20())
		var _196_v6 _dafny.Map
		_ = _196_v6
		_196_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_195_v5).Cardinality())
		_196_v6 = (_196_v6).Update(p1, (_this).F20())
		if true {
			var _197_v7 _dafny.Sequence
			_ = _197_v7
			_197_v7 = _dafny.SeqOf(p0, (_dafny.Zero).Minus((_this).F20()))
			_197_v7 = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(418))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_198_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_199_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F20()), _198_p0)
				}
			})(p0))), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(418))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_200_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_201_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F20()), _200_p0)
				}
			})(p0)))).Cardinality()))).Uint32(), p0)
			var _202_v8 _dafny.CodePoint
			_ = _202_v8
			_202_v8 = _dafny.CodePoint('v')
			_202_v8 = _202_v8
			var _203_v9 _dafny.MultiSet
			_ = _203_v9
			_203_v9 = _dafny.MultiSetOf((_this).F21())
			_193_v3 = (_this).Fm1(_191_v1, (_this).Fm1(_194_v4, _194_v4, _192_v2, globalState), (_203_v9).IsSubsetOf(_203_v9), globalState)
			var _204_v10 *C0
			_ = _204_v10
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__()
			_204_v10 = _nw34
			var _205_v11 *C0
			_ = _205_v11
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__()
			_205_v11 = _nw35
		} else {
			var _206_v12 _dafny.Array
			_ = _206_v12
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw36
			_206_v12 = _nw36
			var _207_v13 _dafny.Sequence
			_ = _207_v13
			_207_v13 = _dafny.SeqOf(_206_v12)
			var _208_v14 D3
			_ = _208_v14
			_208_v14 = Companion_D3_.Create_DC5_(_207_v13)
			var _pat_let_tv4 = _207_v13
			_ = _pat_let_tv4
			var _pat_let_tv5 = _207_v13
			_ = _pat_let_tv5
			var _source6 D3 = func(_pat_let0_0 D3) D3 {
				return func(_211_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let3_0 _dafny.Sequence) D3 {
						return func(_212_dt__update_hcf9_h1 _dafny.Sequence) D3 {
							return Companion_D3_.Create_DC5_(_212_dt__update_hcf9_h1)
						}(_pat_let3_0)
					}(_pat_let_tv5)
				}(_pat_let0_0)
			}((func() D3 {
				if (_this).F21() {
					return _208_v14
				}
				return func(_pat_let1_0 D3) D3 {
					return func(_209_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let2_0 _dafny.Sequence) D3 {
							return func(_210_dt__update_hcf9_h0 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC5_(_210_dt__update_hcf9_h0)
							}(_pat_let2_0)
						}(_pat_let_tv4)
					}(_pat_let1_0)
				}(_208_v14)
			})())
			_ = _source6
			if _source6.Is_DC6() {
				var _213___mcc_h0 _dafny.Int = _source6.Get_().(D3_DC6).Cf10
				_ = _213___mcc_h0
				var _214_cf10 _dafny.Int = _213___mcc_h0
				_ = _214_cf10
				var _215_v15 _dafny.Array
				_ = _215_v15
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_4
				var _nw37 _dafny.Array
				_ = _nw37
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw37 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = (func(_216_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_217_i1 _dafny.Int) _dafny.Int {
							return (_217_i1).Minus(_216_v3)
						}
					})(_193_v3)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw37 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw37).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw37).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_215_v15 = _nw37
				var _218_v16 T0
				_ = _218_v16
				var _nw38 *C1 = New_C1_()
				_ = _nw38
				_nw38.Ctor__((_dafny.SetOf(_192_v2)).Cardinality(), _215_v15, _194_v4, true)
				_218_v16 = _nw38
				var _219_v17 _dafny.Map
				_ = _219_v17
				_219_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _218_v16)
				var _220_v18 _dafny.Set
				_ = _220_v18
				_220_v18 = _dafny.SetOf(_192_v2, p1, _192_v2, _192_v2)
				var _221_v19 _dafny.Map
				_ = _221_v19
				_221_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F20()).Plus((_219_v17).Cardinality()), _220_v18)
				var _222_v20 _dafny.CodePoint
				_ = _222_v20
				_222_v20 = _dafny.CodePoint('f')
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_206_v12), 0))
				_ = _index26
				(_206_v12).ArraySet1((_220_v18).IsDisjointFrom(Companion_Default___.Fm18(_214_cf10, _194_v4, _222_v20, globalState)), (_index26).Int())
				var _223_v21 _dafny.Map
				_ = _223_v21
				_223_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
				var _224_v22 _dafny.Sequence
				_ = _224_v22
				_224_v22 = _dafny.UnicodeSeqOfUtf8Bytes("nlhaskp")
				var _225_v23 _dafny.Map
				_ = _225_v23
				_225_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_190_v0).Select((Companion_Default___.SafeIndex(_193_v3, _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32()).(bool))
				var _226_v24 _dafny.Sequence
				_ = _226_v24
				_226_v24 = _dafny.SeqOf(Companion_Default___.Fm19(true, globalState), _225_v23)
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_206_v12), 0))
				_ = _index27
				var _rhs13 _dafny.Map = _221_v19
				_ = _rhs13
				var _rhs14 _dafny.Int = (_this).F20()
				_ = _rhs14
				var _rhs15 bool = (_218_v16).Fm0(_223_v21, _224_v22, _226_v24, globalState)
				_ = _rhs15
				var _rhs16 bool = !(p1)
				_ = _rhs16
				var _lhs16 *GlobalState = globalState
				_ = _lhs16
				var _lhs17 _dafny.Array = _206_v12
				_ = _lhs17
				var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_206_v12), 0))
				_ = _lhs18
				_221_v19 = _rhs13
				_214_cf10 = _rhs14
				_lhs16.F18 = _rhs15
				(_lhs17).ArraySet1(_rhs16, (_lhs18).Int())
				_190_v0 = _dafny.Companion_Sequence_.Concatenate(_190_v0, _190_v0)
				var _227_v25 _dafny.Array
				_ = _227_v25
				_227_v25 = _215_v15
				var _228_v26 _dafny.Sequence
				_ = _228_v26
				_228_v26 = _dafny.SeqOf(_215_v15)
				var _229_v27 _dafny.Map
				_ = _229_v27
				_229_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v25, (_228_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.IntOfUint32((_228_v26).Cardinality()))).Uint32()).(_dafny.Array))
				var _230_v28 _dafny.Set
				_ = _230_v28
				_230_v28 = _dafny.SetOf((func() _dafny.Array {
					if (_229_v27).Contains(_227_v25) {
						return (_229_v27).Get(_227_v25).(_dafny.Array)
					}
					return _215_v15
				})(), _215_v15, _215_v15)
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_206_v12), 0))
				_ = _index28
				(_206_v12).ArraySet1(!(!(!((_230_v28).IsSubsetOf(_230_v28)))), (_index28).Int())
				(globalState).F3 = (_this).F21()
			} else if _source6.Is_DC7() {
				var _231___mcc_h1 _dafny.Sequence = _source6.Get_().(D3_DC7).Cf11
				_ = _231___mcc_h1
				var _232___mcc_h2 bool = _source6.Get_().(D3_DC7).Cf12
				_ = _232___mcc_h2
				var _233_cf12 bool = _232___mcc_h2
				_ = _233_cf12
				var _234_cf11 _dafny.Sequence = _231___mcc_h1
				_ = _234_cf11
				_191_v1 = (_this).F20()
				var _235_v29 _dafny.Array
				_ = _235_v29
				var _nw39 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
				_ = _nw39
				_235_v29 = _nw39
				var _236_v30 _dafny.Map
				_ = _236_v30
				_236_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v2, _235_v29)
				var _237_v31 _dafny.Sequence
				_ = _237_v31
				_237_v31 = _dafny.SeqOf(_236_v30, _236_v30, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _235_v29))
				_192_v2 = (((_237_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_237_v31).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Cmp(p0) == 0
				var _238_v32 _dafny.Sequence
				_ = _238_v32
				_238_v32 = _dafny.SeqOf(_194_v4)
				_193_v3 = ((_193_v3).Plus((_238_v32).Select((Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_238_v32).Cardinality()))).Uint32()).(_dafny.Int))).Minus(p0)
				var _239_v33 _dafny.Map
				_ = _239_v33
				_239_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v2, p1)
				var _240_v34 _dafny.Sequence
				_ = _240_v34
				_240_v34 = _dafny.UnicodeSeqOfUtf8Bytes("uj")
				var _241_v35 _dafny.Map
				_ = _241_v35
				_241_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(518), _191_v1)
				var _242_v36 bool
				_ = _242_v36
				_242_v36 = p1
				var _243_v37 _dafny.Map
				_ = _243_v37
				_243_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v36, false)
				var _244_v38 _dafny.Map
				_ = _244_v38
				_244_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_206_v12, false)
				var _245_v39 _dafny.Array
				_ = _245_v39
				var _nwElement0_6 _dafny.Int = (_239_v33).Cardinality()
				_ = _nwElement0_6
				var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(28))
				_ = _nw40
				(_nw40).ArraySet1(_nwElement0_6, 0)
				(_nw40).ArraySet1(_191_v1, 1)
				(_nw40).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v34, (_241_v35).Cardinality())).Cardinality(), 2)
				(_nw40).ArraySet1(_193_v3, 3)
				(_nw40).ArraySet1(_194_v4, 4)
				(_nw40).ArraySet1((_this).Fm1(_191_v1, (_241_v35).Cardinality(), _233_cf12, globalState), 5)
				(_nw40).ArraySet1(_dafny.IntOfInt64(388), 6)
				(_nw40).ArraySet1(_191_v1, 7)
				(_nw40).ArraySet1(_193_v3, 8)
				(_nw40).ArraySet1(_dafny.IntOfInt64(135), 9)
				(_nw40).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(p0)).Cardinality(), _194_v4), 10)
				(_nw40).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_240_v34).Cardinality()))).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-625))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}(func(_246_i2 _dafny.Int) _dafny.Int {
					return (_this).F20()
				}))).Cardinality())), 11)
				(_nw40).ArraySet1((func() _dafny.Int {
					if _233_cf12 {
						return (_this).F20()
					}
					return (_243_v37).Cardinality()
				})(), 12)
				(_nw40).ArraySet1((_244_v38).Cardinality(), 13)
				(_nw40).ArraySet1((_this).F20(), 14)
				(_nw40).ArraySet1(_194_v4, 15)
				(_nw40).ArraySet1((_this).F20(), 16)
				(_nw40).ArraySet1(_191_v1, 17)
				(_nw40).ArraySet1(_dafny.IntOfInt64(730), 18)
				(_nw40).ArraySet1(_191_v1, 19)
				(_nw40).ArraySet1(p0, 20)
				(_nw40).ArraySet1((_this).F20(), 21)
				(_nw40).ArraySet1((_this).F20(), 22)
				(_nw40).ArraySet1(_193_v3, 23)
				(_nw40).ArraySet1(_193_v3, 24)
				(_nw40).ArraySet1(p0, 25)
				(_nw40).ArraySet1(_191_v1, 26)
				(_nw40).ArraySet1(_dafny.IntOfInt64(525), 27)
				_245_v39 = _nw40
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_245_v39), 0))
				_ = _index29
				(_245_v39).ArraySet1(_193_v3, (_index29).Int())
				var _247_v40 _dafny.Set
				_ = _247_v40
				_247_v40 = _dafny.SetOf(p1)
				var _248_v41 D7
				_ = _248_v41
				_248_v41 = Companion_D7_.Create_DC16_(_247_v40)
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_245_v39), 0))
				_ = _index30
				(_245_v39).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_248_v41, _248_v41), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(374))).Uint32(), func(coer14 func(_dafny.Int) D7) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}((func(_249_v41 D7) func(_dafny.Int) D7 {
					return func(_250_i3 _dafny.Int) D7 {
						return _249_v41
					}
				})(_248_v41))), Companion_Default___.Fm20((_this).F20(), _193_v3, _193_v3, globalState)))).Cardinality()), (_index30).Int())
			} else if _source6.Is_DC5() {
				var _251___mcc_h3 _dafny.Sequence = _source6.Get_().(D3_DC5).Cf9
				_ = _251___mcc_h3
				var _252_cf9 _dafny.Sequence = _251___mcc_h3
				_ = _252_cf9
				var _253_v42 _dafny.Map
				_ = _253_v42
				_253_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v4, true)
				var _254_v43 _dafny.Map
				_ = _254_v43
				_254_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC16_(_dafny.SetOf((func() bool {
					if (_253_v42).Contains(_dafny.IntOfInt64(636)) {
						return (_253_v42).Get(_dafny.IntOfInt64(636)).(bool)
					}
					return (_this).F21()
				})(), (_this).F21())), p1)
				var _255_v45 _dafny.Sequence
				_ = _255_v45
				_255_v45 = _dafny.UnicodeSeqOfUtf8Bytes("vnfdr")
				var _256_v46 _dafny.Map
				_ = _256_v46
				_256_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p1)
				var _257_v47 _dafny.Sequence
				_ = _257_v47
				_257_v47 = _dafny.SeqOf(_256_v46)
				var _258_v48 _dafny.Map
				_ = _258_v48
				_258_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_192_v2), (_this).Fm0(func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(926), _dafny.IntOfInt64(940))); ; {
						_compr_16, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _259_v44 _dafny.Int
						_259_v44 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(926)).Cmp(_259_v44) <= 0) && ((_259_v44).Cmp(_dafny.IntOfInt64(940)) < 0) {
							_coll16.Add((_259_v44).Plus(_dafny.IntOfUint32((_190_v0).Cardinality())), (_dafny.SetOf(true, p1, _192_v2)).Cardinality())
						}
					}
					return _coll16.ToMap()
				}(), _255_v45, _257_v47, globalState))
				var _260_v49 _dafny.Set
				_ = _260_v49
				_260_v49 = _dafny.SetOf((func() bool {
					if (_256_v46).Contains(_192_v2) {
						return (_256_v46).Get(_192_v2).(bool)
					}
					return true
				})())
				var _261_v50 D7
				_ = _261_v50
				_261_v50 = Companion_D7_.Create_DC16_(_260_v49)
				(globalState).F3 = (func() bool {
					if (_254_v43).Contains(_261_v50) {
						return (_254_v43).Get(_261_v50).(bool)
					}
					return !(p1)
				})()
				var _262_v51 _dafny.Map
				_ = _262_v51
				_262_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_255_v45).Cardinality()), (_this).F20())
				var _263_v52 _dafny.Map
				_ = _263_v52
				_263_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_193_v3), _257_v47)
				var _264_v53 _dafny.MultiSet
				_ = _264_v53
				_264_v53 = _dafny.MultiSetOf((_this).Fm0(_262_v51, _255_v45, (func() _dafny.Sequence {
					if (_263_v52).Contains(p0) {
						return (_263_v52).Get(p0).(_dafny.Sequence)
					}
					return _257_v47
				})(), globalState), _192_v2)
				var _265_v54 _dafny.MultiSet
				_ = _265_v54
				_265_v54 = _dafny.MultiSetOf(((_this).F21()) == ((_this).F21()), (_264_v53).IsProperSubsetOf(_dafny.MultiSetOf(!(!(p1)), (_this).F21())), p1)
				var _266_v55 _dafny.Sequence
				_ = _266_v55
				_266_v55 = _dafny.SeqOf(_265_v54)
				_264_v53 = ((_266_v55).Select((Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_266_v55).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference((_dafny.MultiSetOf(!(p1), true)).Difference(_264_v53))
				var _267_v56 _dafny.Sequence
				_ = _267_v56
				_267_v56 = _dafny.SeqOf(p0, _193_v3)
				var _268_v57 D7
				_ = _268_v57
				_268_v57 = Companion_D7_.Create_DC18_(_267_v56, p0)
				var _269_v58 _dafny.Array
				_ = _269_v58
				var _nwElement0_7 _dafny.Int = _191_v1
				_ = _nwElement0_7
				var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(9))
				_ = _nw41
				(_nw41).ArraySet1(_nwElement0_7, 0)
				(_nw41).ArraySet1(_dafny.IntOfUint32((_255_v45).Cardinality()), 1)
				(_nw41).ArraySet1(_191_v1, 2)
				(_nw41).ArraySet1(_191_v1, 3)
				(_nw41).ArraySet1((_this).F20(), 4)
				(_nw41).ArraySet1(_194_v4, 5)
				(_nw41).ArraySet1(p0, 6)
				(_nw41).ArraySet1((_268_v57).Dtor_cf27(), 7)
				(_nw41).ArraySet1((_this).F20(), 8)
				_269_v58 = _nw41
				var _270_v59 _dafny.Array
				_ = _270_v59
				_270_v59 = _269_v58
				var _271_v60 _dafny.MultiSet
				_ = _271_v60
				_271_v60 = _dafny.MultiSetOf(_270_v59, _270_v59)
				var _272_v61 _dafny.Map
				_ = _272_v61
				_272_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v4, _271_v60)
				var _273_v62 D6
				_ = _273_v62
				_273_v62 = Companion_D6_.Create_DC14_(_dafny.IntOfInt64(764))
				var _274_v63 _dafny.Sequence
				_ = _274_v63
				_274_v63 = _dafny.SeqOf(_269_v58)
				var _275_v64 _dafny.Map
				_ = _275_v64
				_275_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), (_this).F21())
				var _276_v65 D8
				_ = _276_v65
				_276_v65 = Companion_D8_.Create_DC19_(_262_v51)
				var _277_v66 _dafny.Array
				_ = _277_v66
				var _nwElement0_8 _dafny.MultiSet = _271_v60
				_ = _nwElement0_8
				var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(26))
				_ = _nw42
				(_nw42).ArraySet1(_nwElement0_8, 0)
				(_nw42).ArraySet1(_271_v60, 1)
				(_nw42).ArraySet1((_271_v60).Difference(_271_v60), 2)
				(_nw42).ArraySet1(_271_v60, 3)
				(_nw42).ArraySet1(((func() _dafny.MultiSet {
					if (_272_v61).Contains((_273_v62).Dtor_cf22()) {
						return (_272_v61).Get((_273_v62).Dtor_cf22()).(_dafny.MultiSet)
					}
					return (func() _dafny.MultiSet {
						if (_272_v61).Contains(_191_v1) {
							return (_272_v61).Get(_191_v1).(_dafny.MultiSet)
						}
						return _271_v60
					})()
				})()).Union(_271_v60), 4)
				(_nw42).ArraySet1((_271_v60).Intersection(_dafny.MultiSetOf(_270_v59)), 5)
				(_nw42).ArraySet1(_271_v60, 6)
				(_nw42).ArraySet1(_271_v60, 7)
				(_nw42).ArraySet1((_271_v60).Update(_270_v59, Companion_Default___.Abs(_193_v3)), 8)
				(_nw42).ArraySet1(_271_v60, 9)
				(_nw42).ArraySet1(_271_v60, 10)
				(_nw42).ArraySet1((_271_v60).Union(_271_v60), 11)
				(_nw42).ArraySet1(_271_v60, 12)
				(_nw42).ArraySet1(_dafny.MultiSetFromSeq(_274_v63), 13)
				(_nw42).ArraySet1(((_271_v60).Update(_270_v59, Companion_Default___.Abs((_275_v64).Cardinality()))).Intersection(_271_v60), 14)
				(_nw42).ArraySet1(_271_v60, 15)
				(_nw42).ArraySet1((_271_v60).Update(_270_v59, Companion_Default___.Abs(_193_v3)), 16)
				(_nw42).ArraySet1((_dafny.MultiSetOf(_269_v58, _269_v58, _270_v59)).Update(_270_v59, Companion_Default___.Abs((func() _dafny.Int {
					if (_262_v51).Contains(((_276_v65).Dtor_cf28()).Cardinality()) {
						return (_262_v51).Get(((_276_v65).Dtor_cf28()).Cardinality()).(_dafny.Int)
					}
					return _194_v4
				})())), 17)
				(_nw42).ArraySet1(_271_v60, 18)
				(_nw42).ArraySet1(_271_v60, 19)
				(_nw42).ArraySet1(_271_v60, 20)
				(_nw42).ArraySet1(_271_v60, 21)
				(_nw42).ArraySet1((_271_v60).Intersection(_271_v60), 22)
				(_nw42).ArraySet1(_271_v60, 23)
				(_nw42).ArraySet1(_dafny.MultiSetOf(_270_v59), 24)
				(_nw42).ArraySet1(_271_v60, 25)
				_277_v66 = _nw42
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_277_v66), 0))
				_ = _index31
				(_277_v66).ArraySet1(_271_v60, (_index31).Int())
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_277_v66), 0))
				_ = _index32
				var _rhs17 bool = !((func() bool {
					if p1 {
						return (_this).F21()
					}
					return p1
				})())
				_ = _rhs17
				var _rhs18 _dafny.Int = (_dafny.Zero).Minus((_this).F20())
				_ = _rhs18
				var _rhs19 _dafny.MultiSet = _271_v60
				_ = _rhs19
				var _lhs19 *GlobalState = globalState
				_ = _lhs19
				var _lhs20 *GlobalState = globalState
				_ = _lhs20
				var _lhs21 _dafny.Array = _277_v66
				_ = _lhs21
				var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_277_v66), 0))
				_ = _lhs22
				_lhs19.F4 = _rhs17
				_lhs20.F5 = _rhs18
				(_lhs21).ArraySet1(_rhs19, (_lhs22).Int())
				var _278_v67 _dafny.Sequence
				_ = _278_v67
				_278_v67 = _dafny.SeqOf(_190_v0)
				var _279_v68 _dafny.CodePoint
				_ = _279_v68
				_279_v68 = _dafny.CodePoint('s')
				var _280_v69 _dafny.Set
				_ = _280_v69
				_280_v69 = _dafny.SetOf(_279_v68)
				var _281_v70 D9
				_ = _281_v70
				_281_v70 = Companion_D9_.Create_DC22_(_190_v0)
				var _282_v71 D9
				_ = _282_v71
				_282_v71 = Companion_D9_.Create_DC23_(_193_v3, _192_v2)
				var _283_v72 _dafny.Sequence
				_ = _283_v72
				_283_v72 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_284_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_285_i4 _dafny.Int) _dafny.CodePoint {
						return _284_v68
					}
				})(_279_v68)))).Cardinality())))
				var _286_v73 _dafny.Map
				_ = _286_v73
				_286_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v72, p1)
				var _287_v74 _dafny.Array
				_ = _287_v74
				var _nwElement0_9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F21(), _192_v2), _190_v0)
				_ = _nwElement0_9
				var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
				_ = _nw43
				(_nw43).ArraySet1(_nwElement0_9, 0)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm21(globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-21), _dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality()))).Uint32(), (_this).F21()), 1)
				(_nw43).ArraySet1((func() _dafny.Sequence {
					if p1 {
						return _190_v0
					}
					return _dafny.SeqOf(p1, p1)
				})(), 2)
				(_nw43).ArraySet1((_278_v67).Select((Companion_Default___.SafeIndex((_280_v69).Cardinality(), _dafny.IntOfUint32((_278_v67).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Update(_190_v0, (Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32(), p1), 4)
				(_nw43).ArraySet1(_190_v0, 5)
				(_nw43).ArraySet1(Companion_Default___.Fm21(globalState), 6)
				(_nw43).ArraySet1(_dafny.SeqOf(p1), 7)
				(_nw43).ArraySet1(_190_v0, 8)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Update(_190_v0, (Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32(), _192_v2), 9)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_190_v0, _190_v0), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_190_v0, _190_v0)).Cardinality()))).Uint32(), true), (Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_190_v0, _190_v0), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_190_v0, _190_v0)).Cardinality()))).Uint32(), true)).Cardinality()))).Uint32(), _192_v2), 10)
				(_nw43).ArraySet1(_190_v0, 11)
				(_nw43).ArraySet1(_190_v0, 12)
				(_nw43).ArraySet1(_190_v0, 13)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, _192_v2), (_281_v70).Dtor_cf30()), (Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, _192_v2), (_281_v70).Dtor_cf30())).Cardinality()))).Uint32(), (_this).F21()), 14)
				(_nw43).ArraySet1(Companion_Default___.Fm21(globalState), 15)
				(_nw43).ArraySet1((func() _dafny.Sequence {
					if (_282_v71).Dtor_cf32() {
						return _190_v0
					}
					return _190_v0
				})(), 16)
				(_nw43).ArraySet1(_dafny.SeqOf((_this).F21()), 17)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_190_v0, _dafny.SeqOf(p1, (_this).Fm0(Companion_Default___.Fm22(_193_v3, globalState), _255_v45, _dafny.SeqOf(_258_v48), globalState))), 18)
				(_nw43).ArraySet1(Companion_Default___.Fm21(globalState), 19)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_190_v0, _dafny.SeqOf((_this).F21(), (func() bool {
					if (_286_v73).Contains(_283_v72) {
						return (_286_v73).Get(_283_v72).(bool)
					}
					return (_this).F21()
				})(), (_190_v0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F20()), _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32()).(bool), _192_v2)), 20)
				(_nw43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_190_v0, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F21(), p1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_194_v4), _dafny.IntOfUint32((_dafny.SeqOf((_this).F21(), p1)).Cardinality()))).Uint32(), (_this).F21())), 21)
				_287_v74 = _nw43
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_287_v74), 0))
				_ = _index33
				(_287_v74).ArraySet1(_190_v0, (_index33).Int())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_287_v74), 0))
				_ = _index34
				(_287_v74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_192_v2), _190_v0), (_index34).Int())
			} else {
				var _288___mcc_h4 D3 = _source6.Get_().(D3_DC8).Cf13
				_ = _288___mcc_h4
				var _289_cf13 D3 = _288___mcc_h4
				_ = _289_cf13
				var _290_v75 _dafny.Sequence
				_ = _290_v75
				_290_v75 = _dafny.SeqOf((_this).F20())
				var _291_v76 D7
				_ = _291_v76
				_291_v76 = Companion_D7_.Create_DC18_(_290_v75, (_this).F20())
				var _292_v77 _dafny.Map
				_ = _292_v77
				_292_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v2, p1)
				var _293_v78 _dafny.Sequence
				_ = _293_v78
				_293_v78 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v2, p1)).Update(_192_v2, p1), _292_v77, _292_v77)
				var _294_v79 _dafny.Map
				_ = _294_v79
				_294_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_v1, (_this).F21())
				var _rhs20 _dafny.Int = (_194_v4).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_291_v76).Dtor_cf26(), (Companion_Default___.SafeIndex(((_293_v78).Select((Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_293_v78).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32(((_291_v76).Dtor_cf26()).Cardinality()))).Uint32(), _193_v3), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_291_v76).Dtor_cf26(), (Companion_Default___.SafeIndex(((_293_v78).Select((Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_293_v78).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32(((_291_v76).Dtor_cf26()).Cardinality()))).Uint32(), _193_v3)).Cardinality()))).Uint32(), (_294_v79).Cardinality())).Cardinality()))
				_ = _rhs20
				var _rhs21 _dafny.Int = p0
				_ = _rhs21
				var _lhs23 *GlobalState = globalState
				_ = _lhs23
				_lhs23.F17 = _rhs20
				_194_v4 = _rhs21
				var _295_v80 _dafny.Sequence
				_ = _295_v80
				_295_v80 = _dafny.UnicodeSeqOfUtf8Bytes("lmfshk")
				var _296_v81 _dafny.Map
				_ = _296_v81
				_296_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_295_v80, _dafny.IntOfInt64(170))
				_296_v81 = (_296_v81).Update(_dafny.Companion_Sequence_.Concatenate(_295_v80, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_297_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))), _191_v1)
				_194_v4 = (_this).Fm1(_194_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-454))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_298_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(840)
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(85))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_299_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_300_i7 _dafny.Int) _dafny.Int {
						return _299_p0
					}
				})(p0))))).Cardinality()), _192_v2, globalState)
				var _301_v82 _dafny.CodePoint
				_ = _301_v82
				_301_v82 = _dafny.CodePoint('m')
				_301_v82 = _301_v82
			}
			var _302_v83 _dafny.Sequence
			_ = _302_v83
			_302_v83 = _dafny.UnicodeSeqOfUtf8Bytes("xoj")
			(globalState).F15 = _dafny.Companion_Sequence_.Concatenate(_302_v83, _dafny.UnicodeSeqOfUtf8Bytes("hsyle"))
			var _303_v84 _dafny.Sequence
			_ = _303_v84
			_303_v84 = _dafny.SeqOf(_191_v1)
			var _304_v85 _dafny.MultiSet
			_ = _304_v85
			_304_v85 = _dafny.MultiSetOf(_192_v2)
			var _305_v86 _dafny.MultiSet
			_ = _305_v86
			_305_v86 = _dafny.MultiSetOf((Companion_Default___.Fm23(globalState)).Update((_this).F21(), Companion_Default___.Abs(_194_v4)), _304_v85, _304_v85)
			var _306_v87 _dafny.Array
			_ = _306_v87
			var _nwElement0_10 _dafny.Int = (_this).F20()
			_ = _nwElement0_10
			var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(23))
			_ = _nw44
			(_nw44).ArraySet1(_nwElement0_10, 0)
			(_nw44).ArraySet1((_this).F20(), 1)
			(_nw44).ArraySet1((func() _dafny.Int {
				if (_this).F21() {
					return (_this).F20()
				}
				return (_dafny.Zero).Minus(p0)
			})(), 2)
			(_nw44).ArraySet1(_191_v1, 3)
			(_nw44).ArraySet1((_this).F20(), 4)
			(_nw44).ArraySet1(_191_v1, 5)
			(_nw44).ArraySet1(_194_v4, 6)
			(_nw44).ArraySet1((_195_v5).Cardinality(), 7)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt(_194_v4, _191_v1), 8)
			(_nw44).ArraySet1(p0, 9)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt(_193_v3, _194_v4), 10)
			(_nw44).ArraySet1(_194_v4, 11)
			(_nw44).ArraySet1(_193_v3, 12)
			(_nw44).ArraySet1(_194_v4, 13)
			(_nw44).ArraySet1((_dafny.Zero).Minus(_191_v1), 14)
			(_nw44).ArraySet1((_dafny.Zero).Minus(_194_v4), 15)
			(_nw44).ArraySet1((_303_v84).Select((Companion_Default___.SafeIndex((_dafny.SetOf(p1)).Cardinality(), _dafny.IntOfUint32((_303_v84).Cardinality()))).Uint32()).(_dafny.Int), 16)
			(_nw44).ArraySet1(p0, 17)
			(_nw44).ArraySet1((_this).F20(), 18)
			(_nw44).ArraySet1(Companion_Default___.SafeModuloInt(_194_v4, (_this).F20()), 19)
			(_nw44).ArraySet1(((_this).F20()).Times(_194_v4), 20)
			(_nw44).ArraySet1(p0, 21)
			(_nw44).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_305_v86).Cardinality())), 22)
			_306_v87 = _nw44
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_306_v87), 0))
			_ = _index35
			(_306_v87).ArraySet1(p0, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_306_v87), 0))
			_ = _index36
			(_306_v87).ArraySet1(((p0).Times(_dafny.IntOfUint32((_302_v83).Cardinality()))).Minus(((_this).F20()).Plus((_this).Fm1(_194_v4, _193_v3, (_this).F21(), globalState))), (_index36).Int())
			(globalState).F5 = _191_v1
			var _307_v89 _dafny.MultiSet
			_ = _307_v89
			_307_v89 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F20()))
			var _308_v90 D8
			_ = _308_v90
			_308_v90 = Companion_D8_.Create_DC19_(func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter18 := _dafny.Iterate((_307_v89).Elements()); ; {
					_compr_17, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _309_v88 _dafny.Int
					_309_v88 = interface{}(_compr_17).(_dafny.Int)
					if (_307_v89).Contains(_309_v88) {
						_coll17.Add((_309_v88).Times(_dafny.IntOfUint32((_302_v83).Cardinality())), _193_v3)
					}
				}
				return _coll17.ToMap()
			}())
			var _310_v91 _dafny.Map
			_ = _310_v91
			_310_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_307_v89).Cardinality(), _dafny.IntOfInt64(409))
			_308_v90 = Companion_D8_.Create_DC19_(_310_v91)
		}
		var _311_v92 D4
		_ = _311_v92
		_311_v92 = Companion_D4_.Create_DC11_(_191_v1, (_190_v0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32()).(bool), p1)
		var _312_v93 _dafny.Set
		_ = _312_v93
		_312_v93 = _dafny.SetOf(p1)
		var _pat_let_tv6 = _192_v2
		_ = _pat_let_tv6
		var _pat_let_tv7 = _193_v3
		_ = _pat_let_tv7
		var _313_v94 _dafny.Array
		_ = _313_v94
		var _nwElement0_11 D4 = _311_v92
		_ = _nwElement0_11
		var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(28))
		_ = _nw45
		(_nw45).ArraySet1(_nwElement0_11, 0)
		(_nw45).ArraySet1(_311_v92, 1)
		(_nw45).ArraySet1(Companion_D4_.Create_DC11_((_this).F20(), p1, _192_v2), 2)
		(_nw45).ArraySet1(_311_v92, 3)
		(_nw45).ArraySet1(_311_v92, 4)
		(_nw45).ArraySet1(_311_v92, 5)
		(_nw45).ArraySet1(Companion_Default___.Fm24(_191_v1, _192_v2, _312_v93, p1, globalState), 6)
		(_nw45).ArraySet1(_311_v92, 7)
		(_nw45).ArraySet1(_311_v92, 8)
		(_nw45).ArraySet1(_311_v92, 9)
		(_nw45).ArraySet1(_311_v92, 10)
		(_nw45).ArraySet1(_311_v92, 11)
		(_nw45).ArraySet1(_311_v92, 12)
		(_nw45).ArraySet1(_311_v92, 13)
		(_nw45).ArraySet1(Companion_D4_.Create_DC11_(_191_v1, (_this).F21(), _192_v2), 14)
		(_nw45).ArraySet1(_311_v92, 15)
		(_nw45).ArraySet1(_311_v92, 16)
		(_nw45).ArraySet1(Companion_Default___.Fm24(_dafny.IntOfInt64(127), p1, _312_v93, _192_v2, globalState), 17)
		(_nw45).ArraySet1(_311_v92, 18)
		(_nw45).ArraySet1(_311_v92, 19)
		(_nw45).ArraySet1(_311_v92, 20)
		(_nw45).ArraySet1(func(_pat_let4_0 D4) D4 {
			return func(_314_dt__update__tmp_h2 D4) D4 {
				return func(_pat_let5_0 bool) D4 {
					return func(_315_dt__update_hcf19_h0 bool) D4 {
						return Companion_D4_.Create_DC11_((_314_dt__update__tmp_h2).Dtor_cf17(), (_314_dt__update__tmp_h2).Dtor_cf18(), _315_dt__update_hcf19_h0)
					}(_pat_let5_0)
				}((_this).F21())
			}(_pat_let4_0)
		}(_311_v92), 21)
		(_nw45).ArraySet1(func(_pat_let6_0 D4) D4 {
			return func(_316_dt__update__tmp_h3 D4) D4 {
				return func(_pat_let7_0 bool) D4 {
					return func(_317_dt__update_hcf18_h0 bool) D4 {
						return Companion_D4_.Create_DC11_((_316_dt__update__tmp_h3).Dtor_cf17(), _317_dt__update_hcf18_h0, (_316_dt__update__tmp_h3).Dtor_cf19())
					}(_pat_let7_0)
				}(_pat_let_tv6)
			}(_pat_let6_0)
		}(_311_v92), 22)
		(_nw45).ArraySet1(_311_v92, 23)
		(_nw45).ArraySet1(_311_v92, 24)
		(_nw45).ArraySet1(_311_v92, 25)
		(_nw45).ArraySet1(func(_pat_let8_0 D4) D4 {
			return func(_318_dt__update__tmp_h4 D4) D4 {
				return func(_pat_let9_0 bool) D4 {
					return func(_319_dt__update_hcf18_h1 bool) D4 {
						return func(_pat_let10_0 _dafny.Int) D4 {
							return func(_320_dt__update_hcf17_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC11_(_320_dt__update_hcf17_h0, _319_dt__update_hcf18_h1, (_318_dt__update__tmp_h4).Dtor_cf19())
							}(_pat_let10_0)
						}(_pat_let_tv7)
					}(_pat_let9_0)
				}((_this).F21())
			}(_pat_let8_0)
		}(_311_v92), 26)
		(_nw45).ArraySet1(_311_v92, 27)
		_313_v94 = _nw45
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_313_v94), 0))
		_ = _index37
		(_313_v94).ArraySet1(Companion_D4_.Create_DC11_((_this).F20(), !((_this).F21()), (_this).F21()), (_index37).Int())
		var _pat_let_tv8 = _191_v1
		_ = _pat_let_tv8
		var _pat_let_tv9 = _191_v1
		_ = _pat_let_tv9
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_313_v94), 0))
		_ = _index38
		(_313_v94).ArraySet1(func(_pat_let11_0 D4) D4 {
			return func(_321_dt__update__tmp_h5 D4) D4 {
				return func(_pat_let12_0 bool) D4 {
					return func(_322_dt__update_hcf19_h1 bool) D4 {
						return Companion_D4_.Create_DC11_((_321_dt__update__tmp_h5).Dtor_cf17(), (_321_dt__update__tmp_h5).Dtor_cf18(), _322_dt__update_hcf19_h1)
					}(_pat_let12_0)
				}((_pat_let_tv8).Cmp(_pat_let_tv9) < 0)
			}(_pat_let11_0)
		}(_311_v92), (_index38).Int())
		var _323_v95 _dafny.CodePoint
		_ = _323_v95
		_323_v95 = _dafny.CodePoint('w')
		var _324_v96 _dafny.Map
		_ = _324_v96
		_324_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_323_v95, p1)
		if (_192_v2) && (!((func() bool {
			if (_324_v96).Contains(Companion_Default___.Fm17(_192_v2, globalState)) {
				return (_324_v96).Get(Companion_Default___.Fm17(_192_v2, globalState)).(bool)
			}
			return (_190_v0).Select((Companion_Default___.SafeIndex(_191_v1, _dafny.IntOfUint32((_190_v0).Cardinality()))).Uint32()).(bool)
		})()) || (p1)) {
			var _325_v97 _dafny.Sequence
			_ = _325_v97
			_325_v97 = _dafny.SeqOf(_193_v3, _191_v1)
			var _326_v98 _dafny.Map
			_ = _326_v98
			_326_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_v1, _dafny.IntOfUint32((_325_v97).Cardinality()))
			_326_v98 = (_326_v98).Merge(Companion_Default___.Fm22(_dafny.IntOfInt64(398), globalState))
			var _327_v99 _dafny.Sequence
			_ = _327_v99
			_327_v99 = _dafny.UnicodeSeqOfUtf8Bytes("wqqgtwun")
			var _328_v100 _dafny.Map
			_ = _328_v100
			_328_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v99, _323_v95)
			var _329_v101 _dafny.Sequence
			_ = _329_v101
			_329_v101 = _dafny.SeqOf(_325_v97)
			_325_v97 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_328_v100).Cardinality()), (_329_v101).Select((Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_329_v101).Cardinality()))).Uint32()).(_dafny.Sequence)), _325_v97)
			var _330_v102 _dafny.Array
			_ = _330_v102
			var _nwElement0_12 bool = (_this).F21()
			_ = _nwElement0_12
			var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
			_ = _nw46
			(_nw46).ArraySet1(_nwElement0_12, 0)
			(_nw46).ArraySet1(false, 1)
			(_nw46).ArraySet1(false, 2)
			(_nw46).ArraySet1(p1, 3)
			_330_v102 = _nw46
			var _331_v103 _dafny.Sequence
			_ = _331_v103
			_331_v103 = _dafny.SeqOf(_330_v102, _330_v102, _330_v102, _330_v102, _330_v102)
			var _332_v104 D3
			_ = _332_v104
			_332_v104 = Companion_D3_.Create_DC5_(_331_v103)
			_332_v104 = _332_v104
			var _333_v105 _dafny.Sequence
			_ = _333_v105
			_333_v105 = _dafny.SeqOf((_dafny.SetOf((_196_v6).Cardinality(), (_this).F20())).Difference(Companion_Default___.Fm25(globalState)), _195_v5, _195_v5)
			_333_v105 = _dafny.SeqOf(_195_v5, _195_v5, _195_v5)
			_192_v2 = p1
		} else {
			var _334_v106 _dafny.Sequence
			_ = _334_v106
			_334_v106 = _dafny.SeqOf(_194_v4)
			(globalState).F5 = ((_334_v106).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_334_v106).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_191_v1)
			var _335_v107 *C0
			_ = _335_v107
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__()
			_335_v107 = _nw47
			var _336_v108 _dafny.Sequence
			_ = _336_v108
			_336_v108 = _dafny.UnicodeSeqOfUtf8Bytes("ssmu")
			var _337_v109 _dafny.MultiSet
			_ = _337_v109
			_337_v109 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_336_v108, (Companion_Default___.SafeIndex(_194_v4, _dafny.IntOfUint32((_336_v108).Cardinality()))).Uint32(), Companion_Default___.Fm17(p1, globalState))).Cardinality()), Companion_Default___.SafeDivisionInt(_194_v4, p0), (_dafny.Zero).Minus((_this).F20()), (_dafny.IntOfUint32((_336_v108).Cardinality())).Minus(_194_v4), (_this).Fm1(_191_v1, _193_v3, _192_v2, globalState))
			_337_v109 = (_337_v109).Union(_337_v109)
			var _338_v110 *C0
			_ = _338_v110
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__()
			_338_v110 = _nw48
			var _339_v111 D8
			_ = _339_v111
			_339_v111 = Companion_D8_.Create_DC20_()
			_339_v111 = Companion_D8_.Create_DC20_()
		}
		var _340_v112 _dafny.Map
		_ = _340_v112
		_340_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _194_v4)
		var _341_v113 _dafny.Sequence
		_ = _341_v113
		_341_v113 = _dafny.UnicodeSeqOfUtf8Bytes("ac")
		_340_v112 = (_340_v112).Update(((_this).F20()).Plus(_dafny.IntOfUint32((_341_v113).Cardinality())), p0)
		var _342_v114 _dafny.Array
		_ = _342_v114
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw49
		_342_v114 = _nw49
		var _343_v115 _dafny.Sequence
		_ = _343_v115
		_343_v115 = _dafny.SeqOf(_342_v114, _342_v114, _342_v114)
		r0 = (func() _dafny.Sequence {
			if _192_v2 {
				return _343_v115
			}
			return _343_v115
		})()
		r1 = p1
		return r0, r1
	}
}
func (_this *C2) M3(p0 bool, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _hi1 _dafny.Int = (_this).F20()
		_ = _hi1
		for _344_i0 := _dafny.IntOfUint32((Companion_Default___.Fm26((_this).F21(), globalState)).Cardinality()); _344_i0.Cmp(_hi1) < 0; _344_i0 = _344_i0.Plus(_dafny.One) {
			(globalState).F3 = false
			var _345_v0 _dafny.Sequence
			_ = _345_v0
			_345_v0 = _dafny.UnicodeSeqOfUtf8Bytes("nfbohx")
			var _346_v1 _dafny.CodePoint
			_ = _346_v1
			_346_v1 = _dafny.CodePoint('r')
			(globalState).F15 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_345_v0, (Companion_Default___.SafeIndex(_344_i0, _dafny.IntOfUint32((_345_v0).Cardinality()))).Uint32(), _346_v1), _345_v0)
			(globalState).F15 = _345_v0
			(globalState).F3 = (_this).F21()
		}
		var _347_v2 _dafny.Sequence
		_ = _347_v2
		_347_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fahojwr")
		(globalState).F5 = (_this).Fm1((_this).F20(), _dafny.IntOfUint32((_347_v2).Cardinality()), (_this).F21(), globalState)
		var _348_v3 _dafny.Array
		_ = _348_v3
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_5
		var _nw50 _dafny.Array
		_ = _nw50
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw50 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_349_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_349_i1, (_this).F20())
			}
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
		_348_v3 = _nw50
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
		_ = _index39
		(_348_v3).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F20(), (_this).F20())).Cardinality()), (_index39).Int())
		var _350_v4 _dafny.CodePoint
		_ = _350_v4
		_350_v4 = _dafny.CodePoint('g')
		var _351_v5 _dafny.Sequence
		_ = _351_v5
		_351_v5 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_347_v2, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_347_v2).Cardinality()))).Uint32(), _350_v4)).Cardinality()))
		var _352_v6 _dafny.Map
		_ = _352_v6
		_352_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_351_v5).Cardinality()))
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
		_ = _index40
		(_348_v3).ArraySet1(((_this).F20()).Minus((func() _dafny.Int {
			if (_352_v6).Contains((_this).F21()) {
				return (_352_v6).Get((_this).F21()).(_dafny.Int)
			}
			return (_this).F20()
		})()), (_index40).Int())
		var _353_v7 _dafny.Map
		_ = _353_v7
		_353_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wpkgyo")).Cardinality()))
		var _354_v8 _dafny.MultiSet
		_ = _354_v8
		_354_v8 = _dafny.MultiSetOf(_dafny.CodePoint('a'))
		var _355_v9 D7
		_ = _355_v9
		_355_v9 = Companion_D7_.Create_DC18_(_351_v5, (_354_v8).Cardinality())
		var _hi2 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_353_v7).Contains(_348_v3) {
				return (_353_v7).Get(_348_v3).(_dafny.Int)
			}
			return (_355_v9).Dtor_cf27()
		})(), (_this).F20())
		_ = _hi2
		for _356_i2 := (_dafny.IntOfUint32((_347_v2).Cardinality())).Plus((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)); _356_i2.Cmp(_hi2) < 0; _356_i2 = _356_i2.Plus(_dafny.One) {
			var _357_v10 D3
			_ = _357_v10
			_357_v10 = Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ksptncyh")).Cardinality()))
			var _358_v11 _dafny.Map
			_ = _358_v11
			_358_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_356_i2, (_dafny.Zero).Minus(_356_i2))
			var _359_v12 _dafny.Map
			_ = _359_v12
			_359_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), !((_this).F21()))
			var _360_v13 _dafny.Sequence
			_ = _360_v13
			_360_v13 = _dafny.SeqOf(_359_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false), _359_v12, _359_v12)
			var _361_v14 _dafny.Sequence
			_ = _361_v14
			_361_v14 = _dafny.SeqOf(_359_v12, (_360_v13).Select((Companion_Default___.SafeIndex((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_360_v13).Cardinality()))).Uint32()).(_dafny.Map), _359_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21()))
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _index41
			var _rhs22 _dafny.CodePoint = _350_v4
			_ = _rhs22
			var _rhs23 _dafny.Int = (_357_v10).Dtor_cf10()
			_ = _rhs23
			var _rhs24 bool = (_this).F21()
			_ = _rhs24
			var _rhs25 bool = !((_this).F21()) || (!((_this).Fm0(_358_v11, _347_v2, _361_v14, globalState)))
			_ = _rhs25
			var _lhs24 _dafny.Array = _348_v3
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _lhs25
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			_350_v4 = _rhs22
			(_lhs24).ArraySet1(_rhs23, (_lhs25).Int())
			_lhs26.F14 = _rhs24
			_lhs27.F18 = _rhs25
			var _362_v15 _dafny.Sequence
			_ = _362_v15
			_362_v15 = _dafny.SeqOf((_this).Fm0(_358_v11, _347_v2, _361_v14, globalState), p0, (_this).Fm0(_358_v11, _dafny.UnicodeSeqOfUtf8Bytes("mgw"), _360_v13, globalState))
			var _363_v16 D4
			_ = _363_v16
			_363_v16 = Companion_D4_.Create_DC11_((_this).F20(), (_this).F21(), !((_this).F21()))
			if (_362_v15).Select((Companion_Default___.SafeIndex((_363_v16).Dtor_cf17(), _dafny.IntOfUint32((_362_v15).Cardinality()))).Uint32()).(bool) {
				var _364_v17 _dafny.Array
				_ = _364_v17
				var _nw51 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw51
				_364_v17 = _nw51
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_364_v17), 0))
				_ = _index42
				(_364_v17).ArraySet1((_this).Fm0(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-127), _dafny.IntOfInt64(39))); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _365_v18 _dafny.Int
						_365_v18 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(-127)).Cmp(_365_v18) <= 0) && ((_365_v18).Cmp(_dafny.IntOfInt64(39)) < 0) {
							_coll18.Add((_365_v18).Minus((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)), _356_i2)
						}
					}
					return _coll18.ToMap()
				}(), _347_v2, _361_v14, globalState), (_index42).Int())
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_364_v17), 0))
				_ = _index43
				(_364_v17).ArraySet1((_this).F21(), (_index43).Int())
				var _366_v19 D4
				_ = _366_v19
				_366_v19 = Companion_D4_.Create_DC9_(_364_v17)
				var _367_v20 _dafny.Map
				_ = _367_v20
				_367_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_364_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_364_v17), 0))).Int()).(bool), (_366_v19).Dtor_cf14())
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
				_ = _index44
				(_348_v3).ArraySet1((_this).Fm1(_356_i2, (_367_v20).Cardinality(), p0, globalState), (_index44).Int())
				(globalState).F17 = (_dafny.IntOfUint32((_347_v2).Cardinality())).Plus((_this).F20())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_364_v17), 0))
				_ = _index45
				var _rhs26 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(297), _356_i2), _356_i2)
				_ = _rhs26
				var _rhs27 _dafny.Int = (_dafny.Zero).Minus((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
				_ = _rhs27
				var _lhs28 _dafny.Array = _364_v17
				_ = _lhs28
				var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_364_v17), 0))
				_ = _lhs29
				var _lhs30 *GlobalState = globalState
				_ = _lhs30
				(_lhs28).ArraySet1(_rhs26, (_lhs29).Int())
				_lhs30.F5 = _rhs27
				var _368_v21 _dafny.Map
				_ = _368_v21
				_368_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mqpn"), ((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)).Times(_356_i2))
				var _369_v22 _dafny.Map
				_ = _369_v22
				_369_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_347_v2).Cardinality()), (_this).F21())
				var _370_v23 D9
				_ = _370_v23
				_370_v23 = Companion_D9_.Create_DC23_(_356_i2, p0)
				_368_v21 = (_368_v21).Update(Companion_Default___.Fm26(!((func() bool {
					if (_369_v22).Contains((_370_v23).Dtor_cf31()) {
						return (_369_v22).Get((_370_v23).Dtor_cf31()).(bool)
					}
					return true
				})()), globalState), (_356_i2).Minus(_356_i2))
			} else {
				_351_v5 = _351_v5
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
				_ = _index46
				(_348_v3).ArraySet1((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_index46).Int())
				var _371_v24 _dafny.Array
				_ = _371_v24
				var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
				_ = _nw52
				_371_v24 = _nw52
				var _rhs28 _dafny.Array = _371_v24
				_ = _rhs28
				var _rhs29 _dafny.Int = Companion_Default___.SafeDivisionInt((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)).Plus(_356_i2)))
				_ = _rhs29
				var _rhs30 _dafny.Int = (_dafny.Zero).Minus((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
				_ = _rhs30
				var _rhs31 bool = (((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(193))).Cmp((_dafny.Zero).Minus((_this).F20())) == 0
				_ = _rhs31
				var _lhs31 *GlobalState = globalState
				_ = _lhs31
				var _lhs32 *GlobalState = globalState
				_ = _lhs32
				_371_v24 = _rhs28
				r3 = _rhs29
				_lhs31.F17 = _rhs30
				_lhs32.F4 = _rhs31
				var _372_v25 _dafny.Map
				_ = _372_v25
				_372_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _362_v15)
				(globalState).F5 = ((_dafny.Zero).Minus((_this).F20())).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _362_v15)).Merge(_372_v25)).Cardinality())
				var _373_v26 _dafny.Array
				_ = _373_v26
				var _nwElement0_13 bool = (_this).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F20()), (_this).Fm1((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), p0, globalState)), _dafny.UnicodeSeqOfUtf8Bytes("vbcdjta"), _360_v13, globalState)
				_ = _nwElement0_13
				var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
				_ = _nw53
				(_nw53).ArraySet1(_nwElement0_13, 0)
				_373_v26 = _nw53
				var _374_v27 D8
				_ = _374_v27
				_374_v27 = Companion_D8_.Create_DC20_()
				var _rhs32 bool = (_this).F21()
				_ = _rhs32
				var _rhs33 _dafny.Array = _373_v26
				_ = _rhs33
				var _rhs34 D8 = Companion_D8_.Create_DC20_()
				_ = _rhs34
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				_lhs33.F3 = _rhs32
				_373_v26 = _rhs33
				_374_v27 = _rhs34
			}
			var _rhs35 bool = (_this).F21()
			_ = _rhs35
			var _rhs36 _dafny.Int = _dafny.IntOfInt64(948)
			_ = _rhs36
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			_lhs34.F4 = _rhs35
			_lhs35.F5 = _rhs36
			var _375_v28 _dafny.MultiSet
			_ = _375_v28
			_375_v28 = _dafny.MultiSetOf(!(p0))
			var _376_v29 *C1
			_ = _376_v29
			var _nw54 *C1 = New_C1_()
			_ = _nw54
			_nw54.Ctor__((_this).F20(), _348_v3, Companion_Default___.SafeModuloInt((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_this).F20()), (Companion_Default___.Fm23(globalState)).IsSubsetOf(_375_v28))
			_376_v29 = _nw54
		}
		if (_this).F21() {
			r0 = (_this).F20()
			var _377_v30 _dafny.Map
			_ = _377_v30
			_377_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_350_v4, (_this).F21())
			var _378_v31 _dafny.Map
			_ = _378_v31
			_378_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v30, (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
			_378_v31 = (_378_v31).Update(_377_v30, (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
			var _379_v32 _dafny.Map
			_ = _379_v32
			_379_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
			var _380_v33 _dafny.Map
			_ = _380_v33
			_380_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
			var _381_v34 _dafny.Set
			_ = _381_v34
			_381_v34 = _dafny.SetOf(p0, p0, (_this).F21(), (_this).F21())
			r1 = (_381_v34).IsProperSubsetOf((_dafny.SetOf((_this).Fm0(_379_v32, _347_v2, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)), (Companion_Default___.SafeIndex((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))).Cardinality()))).Uint32(), _380_v33), globalState), p0, p0)).Union(_dafny.SetOf(p0)))
			var _382_v35 _dafny.Array
			_ = _382_v35
			var _nwElement0_14 _dafny.CodePoint = _350_v4
			_ = _nwElement0_14
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(10))
			_ = _nw55
			(_nw55).ArraySet1CodePoint(_nwElement0_14, 0)
			(_nw55).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _350_v4
				}
				return _dafny.CodePoint('t')
			})(), 1)
			(_nw55).ArraySet1CodePoint(_350_v4, 2)
			(_nw55).ArraySet1CodePoint(Companion_Default___.Fm17(p0, globalState), 3)
			(_nw55).ArraySet1CodePoint(_350_v4, 4)
			(_nw55).ArraySet1CodePoint(_350_v4, 5)
			(_nw55).ArraySet1CodePoint(_350_v4, 6)
			(_nw55).ArraySet1CodePoint(_350_v4, 7)
			(_nw55).ArraySet1CodePoint(_350_v4, 8)
			(_nw55).ArraySet1CodePoint(_dafny.CodePoint('k'), 9)
			_382_v35 = _nw55
			var _383_v36 _dafny.MultiSet
			_ = _383_v36
			_383_v36 = _dafny.MultiSetOf(true, false, (_this).F21())
			var _384_v37 D3
			_ = _384_v37
			_384_v37 = Companion_D3_.Create_DC6_((_383_v36).Cardinality())
			var _385_v38 _dafny.MultiSet
			_ = _385_v38
			_385_v38 = _dafny.MultiSetOf((_dafny.SetOf((_this).F21(), (_this).F21(), (_this).F21())).Difference(_381_v34))
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _index47
			var _rhs37 _dafny.Array = _382_v35
			_ = _rhs37
			var _rhs38 D3 = _384_v37
			_ = _rhs38
			var _rhs39 _dafny.Int = (_this).F20()
			_ = _rhs39
			var _rhs40 _dafny.Int = (func() _dafny.Int {
				if ((_this).F20()).Cmp((_this).F20()) >= 0 {
					return _dafny.IntOfUint32((_347_v2).Cardinality())
				}
				return (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)
			})()
			_ = _rhs40
			var _rhs41 _dafny.MultiSet = _385_v38
			_ = _rhs41
			var _lhs36 _dafny.Array = _348_v3
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _lhs37
			_382_v35 = _rhs37
			_384_v37 = _rhs38
			r0 = _rhs39
			(_lhs36).ArraySet1(_rhs40, (_lhs37).Int())
			_385_v38 = _rhs41
			var _386_v39 _dafny.Sequence
			_ = _386_v39
			_386_v39 = _dafny.SeqOf(_380_v33)
			(globalState).F3 = (_this).Fm0(_379_v32, Companion_Default___.Fm26(p0, globalState), _dafny.Companion_Sequence_.Concatenate(_386_v39, _386_v39), globalState)
		} else {
			(globalState).F14 = (_this).F21()
			(globalState).F5 = (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)
			(globalState).F17 = (_this).F20()
			var _387_v40 _dafny.Map
			_ = _387_v40
			_387_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_347_v2, (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _index48
			(_348_v3).ArraySet1(Companion_Default___.SafeModuloInt((_387_v40).Cardinality(), (_this).F20()), (_index48).Int())
			var _388_v41 bool
			_ = _388_v41
			_388_v41 = p0
			_388_v41 = _388_v41
		}
		var _389_v42 _dafny.Map
		_ = _389_v42
		_389_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_this).F20())
		var _390_v43 _dafny.Map
		_ = _390_v43
		_390_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _391_v44 _dafny.Sequence
		_ = _391_v44
		_391_v44 = _dafny.SeqOf(_390_v43)
		var _392_v45 _dafny.Map
		_ = _392_v45
		_392_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).Fm0(_389_v42, _347_v2, _391_v44, globalState))
		var _393_v46 _dafny.Map
		_ = _393_v46
		_393_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _390_v43)
		if !(((_390_v43).Merge(_392_v45)).Equals(((func() _dafny.Map {
			if (_393_v46).Contains(_dafny.IntOfInt64(240)) {
				return (_393_v46).Get(_dafny.IntOfInt64(240)).(_dafny.Map)
			}
			return Companion_Default___.Fm19((_this).F21(), globalState)
		})()).Merge(_392_v45))) {
			var _394_v47 _dafny.Set
			_ = _394_v47
			_394_v47 = _dafny.SetOf(_350_v4, _350_v4)
			var _395_v48 _dafny.Sequence
			_ = _395_v48
			_395_v48 = _dafny.SeqOf(_351_v5, _351_v5)
			_392_v45 = (_392_v45).Update(((_394_v47).Cardinality()).Cmp(_dafny.IntOfUint32((_395_v48).Cardinality())) < 0, true)
			var _396_v49 _dafny.Array
			_ = _396_v49
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_6
			var _nw56 _dafny.Array
			_ = _nw56
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw56 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = func(_397_i3 _dafny.Int) bool {
					return (_this).F21()
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw56 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw56).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw56).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_396_v49 = _nw56
			var _398_v50 D4
			_ = _398_v50
			_398_v50 = Companion_D4_.Create_DC9_(_396_v49)
			var _pat_let_tv10 = _396_v49
			_ = _pat_let_tv10
			var _399_v51 _dafny.MultiSet
			_ = _399_v51
			_399_v51 = _dafny.MultiSetOf(func(_pat_let13_0 D4) D4 {
				return func(_400_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let14_0 _dafny.Array) D4 {
						return func(_401_dt__update_hcf14_h0 _dafny.Array) D4 {
							return Companion_D4_.Create_DC9_(_401_dt__update_hcf14_h0)
						}(_pat_let14_0)
					}(_pat_let_tv10)
				}(_pat_let13_0)
			}(_398_v50), _398_v50, Companion_D4_.Create_DC9_(_396_v49), _398_v50)
			var _402_v52 _dafny.MultiSet
			_ = _402_v52
			_402_v52 = _dafny.MultiSetOf((_this).F20(), (_dafny.Zero).Minus((_this).F20()), (func() _dafny.Int {
				if (_399_v51).Contains(_398_v50) {
					return (_399_v51).Multiplicity(_398_v50)
				}
				return (_dafny.Zero).Minus((_this).F20())
			})(), _dafny.IntOfInt64(969))
			var _403_v53 _dafny.Map
			_ = _403_v53
			_403_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_402_v52, _355_v9)
			var _pat_let_tv11 = _351_v5
			_ = _pat_let_tv11
			_403_v53 = (_403_v53).Update(_dafny.MultiSetOf((_this).F20(), (_this).F20(), _dafny.IntOfInt64(-398), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)), func(_pat_let15_0 D7) D7 {
				return func(_404_dt__update__tmp_h1 D7) D7 {
					return func(_pat_let16_0 _dafny.Sequence) D7 {
						return func(_405_dt__update_hcf26_h0 _dafny.Sequence) D7 {
							return Companion_D7_.Create_DC18_(_405_dt__update_hcf26_h0, (_404_dt__update__tmp_h1).Dtor_cf27())
						}(_pat_let16_0)
					}(_pat_let_tv11)
				}(_pat_let15_0)
			}(Companion_D7_.Create_DC18_(_351_v5, (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))))
			var _406_v54 _dafny.MultiSet
			_ = _406_v54
			_406_v54 = _dafny.MultiSetOf((_this).F21())
			var _407_v55 _dafny.Map
			_ = _407_v55
			_407_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_406_v54).Cardinality(), _348_v3)
			var _408_v56 _dafny.Set
			_ = _408_v56
			_408_v56 = _dafny.SetOf((_this).F21())
			var _409_v57 *C1
			_ = _409_v57
			var _nw57 *C1 = New_C1_()
			_ = _nw57
			_nw57.Ctor__((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (func() _dafny.Array {
				if (_407_v55).Contains((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)) {
					return (_407_v55).Get((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)).(_dafny.Array)
				}
				return _348_v3
			})(), Companion_Default___.SafeDivisionInt((_this).F20(), (_dafny.Zero).Minus((_408_v56).Cardinality())), !((_this).F21()))
			_409_v57 = _nw57
			var _rhs42 _dafny.Int = _dafny.IntOfInt64(632)
			_ = _rhs42
			var _rhs43 *C1 = _409_v57
			_ = _rhs43
			r3 = _rhs42
			_409_v57 = _rhs43
			var _410_v58 _dafny.Map
			_ = _410_v58
			_410_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_396_v49, (_dafny.Zero).Minus((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)))
			_410_v58 = (_410_v58).Update(_396_v49, (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))
			_ = _index49
			(_348_v3).ArraySet1((_dafny.Zero).Minus(_409_v57.F27), (_index49).Int())
		} else {
			var _411_v59 _dafny.Array
			_ = _411_v59
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_7
			var _nw58 _dafny.Array
			_ = _nw58
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw58 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_412_p0 bool) func(_dafny.Int) bool {
					return func(_413_i4 _dafny.Int) bool {
						return _412_p0
					}
				})(p0)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw58 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw58).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw58).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_411_v59 = _nw58
			var _414_v60 D4
			_ = _414_v60
			_414_v60 = Companion_D4_.Create_DC9_(_411_v59)
			var _415_v61 _dafny.Sequence
			_ = _415_v61
			_415_v61 = _dafny.SeqOf(!(((_this).F20()).Cmp((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)) != 0), (_this).F21())
			var _rhs44 D4 = _414_v60
			_ = _rhs44
			var _rhs45 _dafny.Sequence = _415_v61
			_ = _rhs45
			var _rhs46 bool = (_this).F21()
			_ = _rhs46
			var _rhs47 _dafny.Int = (_this).F20()
			_ = _rhs47
			var _lhs38 *GlobalState = globalState
			_ = _lhs38
			_414_v60 = _rhs44
			_415_v61 = _rhs45
			r1 = _rhs46
			_lhs38.F17 = _rhs47
			var _416_v62 _dafny.Map
			_ = _416_v62
			_416_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), p0)
			_389_v42 = (_389_v42).Update(((_this).F20()).Times((_416_v62).Cardinality()), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int))
			var _417_v63 _dafny.Set
			_ = _417_v63
			_417_v63 = _dafny.SetOf((_this).F21(), (_this).F21())
			var _418_v64 _dafny.MultiSet
			_ = _418_v64
			_418_v64 = _dafny.MultiSetOf(_417_v63)
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_411_v59), 0))
			_ = _index50
			(_411_v59).ArraySet1(((_this).F20()).Cmp((_418_v64).Cardinality()) >= 0, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_411_v59), 0))
			_ = _index51
			(_411_v59).ArraySet1(((_this).F21()) && (p0), (_index51).Int())
			(globalState).F4 = (_411_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_411_v59), 0))).Int()).(bool)
			var _419_v65 _dafny.Sequence
			_ = _419_v65
			_419_v65 = _dafny.SeqOf(_dafny.SeqOf((_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int), (_this).F20(), (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)), _351_v5)
			var _420_v66 _dafny.Sequence
			_ = _420_v66
			_420_v66 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_351_v5, (_419_v65).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_419_v65).Cardinality()))).Uint32()).(_dafny.Sequence)), _351_v5)
			_419_v65 = _420_v66
		}
		r0 = (_348_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_348_v3), 0))).Int()).(_dafny.Int)
		r1 = (_this).Fm0(_389_v42, _347_v2, _391_v44, globalState)
		r2 = _dafny.IntOfInt64(-558)
		var _421_v67 _dafny.Map
		_ = _421_v67
		_421_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_351_v5).Cardinality()), p0)
		var _422_v68 _dafny.Map
		_ = _422_v68
		_422_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_421_v67, p0)
		r3 = Companion_Default___.SafeModuloInt((_422_v68).Cardinality(), (_this).F20())
		return r0, r1, r2, r3
	}
}
func (_this *C2) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 D3, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _hi3 _dafny.Int = (_this).F20()
		_ = _hi3
		for _423_i0 := (_dafny.Zero).Minus(p0); _423_i0.Cmp(_hi3) < 0; _423_i0 = _423_i0.Plus(_dafny.One) {
			var _424_v0 _dafny.Array
			_ = _424_v0
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw59
			_424_v0 = _nw59
			var _425_v1 _dafny.Array
			_ = _425_v1
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw60
			_425_v1 = _nw60
			var _426_v2 D3
			_ = _426_v2
			_426_v2 = Companion_D3_.Create_DC6_(_423_i0)
			var _427_v3 _dafny.Sequence
			_ = _427_v3
			_427_v3 = _dafny.UnicodeSeqOfUtf8Bytes("qvejr")
			var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_427_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-478), _dafny.IntOfUint32((_427_v3).Cardinality()))).Uint32(), _dafny.CodePoint('r')), _427_v3)
			_ = _rhs48
			var _rhs49 _dafny.Array = _424_v0
			_ = _rhs49
			var _rhs50 _dafny.Array = _425_v1
			_ = _rhs50
			var _rhs51 D3 = Companion_D3_.Create_DC6_(_423_i0)
			_ = _rhs51
			var _rhs52 _dafny.Int = p0
			_ = _rhs52
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			var _lhs40 *GlobalState = globalState
			_ = _lhs40
			_lhs39.F15 = _rhs48
			_424_v0 = _rhs49
			_425_v1 = _rhs50
			_426_v2 = _rhs51
			_lhs40.F5 = _rhs52
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))
			_ = _index52
			(_424_v0).ArraySet1(_425_v1, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))
			_ = _index53
			(_424_v0).ArraySet1(_425_v1, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_425_v1), 0))
			_ = _index54
			(_425_v1).ArraySet1(p1, (_index54).Int())
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))
			_ = _arr0
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))), 0))
			_ = _index55
			(_arr0).ArraySet1((_this).F21(), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_425_v1), 0))
			_ = _index56
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))
			_ = _arr1
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))), 0))
			_ = _index57
			var _rhs53 bool = (_this).F21()
			_ = _rhs53
			var _rhs54 bool = p1
			_ = _rhs54
			var _lhs41 _dafny.Array = _425_v1
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_425_v1), 0))
			_ = _lhs42
			var _lhs43 _dafny.Array = _dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_dafny.ArrayCastTo((_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_424_v0), 0))).Int()))), 0))
			_ = _lhs44
			(_lhs41).ArraySet1(_rhs53, (_lhs42).Int())
			(_lhs43).ArraySet1(_rhs54, (_lhs44).Int())
			var _428_v4 _dafny.Set
			_ = _428_v4
			_428_v4 = _dafny.SetOf((_425_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_425_v1), 0))).Int()).(bool))
			var _429_v5 _dafny.Array
			_ = _429_v5
			var _nwElement0_15 _dafny.Int = _423_i0
			_ = _nwElement0_15
			var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(6))
			_ = _nw61
			(_nw61).ArraySet1(_nwElement0_15, 0)
			(_nw61).ArraySet1((_this).F20(), 1)
			(_nw61).ArraySet1(_dafny.IntOfInt64(-412), 2)
			(_nw61).ArraySet1((_dafny.Zero).Minus(_423_i0), 3)
			(_nw61).ArraySet1(_423_i0, 4)
			(_nw61).ArraySet1((_428_v4).Cardinality(), 5)
			_429_v5 = _nw61
			var _430_v6 *C1
			_ = _430_v6
			var _nw62 *C1 = New_C1_()
			_ = _nw62
			_nw62.Ctor__(p0, _429_v5, _423_i0, p1)
			_430_v6 = _nw62
		}
		var _431_i1 _dafny.Int
		_ = _431_i1
		_431_i1 = _dafny.Zero
		{
			for (_this).F21() {
				{
					if (_431_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_431_i1 = (_431_i1).Plus(_dafny.One)
					var _432_v7 _dafny.Sequence
					_ = _432_v7
					_432_v7 = _dafny.SeqOf(p0)
					var _433_v8 _dafny.Sequence
					_ = _433_v8
					_433_v8 = _dafny.UnicodeSeqOfUtf8Bytes("ogxg")
					var _434_v9 _dafny.Sequence
					_ = _434_v9
					_434_v9 = _dafny.SeqOf((_this).F21(), p1)
					var _435_v10 _dafny.Array
					_ = _435_v10
					var _nwElement0_16 _dafny.Int = (_432_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_432_v7).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _nwElement0_16
					var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(24))
					_ = _nw63
					(_nw63).ArraySet1(_nwElement0_16, 0)
					(_nw63).ArraySet1(_dafny.IntOfUint32((_433_v8).Cardinality()), 1)
					(_nw63).ArraySet1((_this).F20(), 2)
					(_nw63).ArraySet1((_this).F20(), 3)
					(_nw63).ArraySet1((_432_v7).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_432_v7).Cardinality()))).Uint32()).(_dafny.Int), 4)
					(_nw63).ArraySet1((_this).F20(), 5)
					(_nw63).ArraySet1(p0, 6)
					(_nw63).ArraySet1(_dafny.IntOfUint32((_433_v8).Cardinality()), 7)
					(_nw63).ArraySet1(_dafny.IntOfInt64(551), 8)
					(_nw63).ArraySet1(p0, 9)
					(_nw63).ArraySet1((_this).Fm1(_dafny.IntOfInt64(-469), (_this).F20(), (_this).F21(), globalState), 10)
					(_nw63).ArraySet1(_dafny.IntOfInt64(158), 11)
					(_nw63).ArraySet1(p0, 12)
					(_nw63).ArraySet1(p0, 13)
					(_nw63).ArraySet1(p0, 14)
					(_nw63).ArraySet1((_this).F20(), 15)
					(_nw63).ArraySet1(_dafny.IntOfUint32((_433_v8).Cardinality()), 16)
					(_nw63).ArraySet1(_dafny.IntOfUint32((_434_v9).Cardinality()), 17)
					(_nw63).ArraySet1(p0, 18)
					(_nw63).ArraySet1((_dafny.SetOf(true, p1, (_this).F21(), (_this).F21())).Cardinality(), 19)
					(_nw63).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_432_v7, p1)).Cardinality(), 20)
					(_nw63).ArraySet1((_dafny.Zero).Minus((_this).Fm1(_dafny.IntOfUint32((_433_v8).Cardinality()), (_this).F20(), p1, globalState)), 21)
					(_nw63).ArraySet1(p0, 22)
					(_nw63).ArraySet1(p0, 23)
					_435_v10 = _nw63
					var _436_v11 *C1
					_ = _436_v11
					var _nw64 *C1 = New_C1_()
					_ = _nw64
					_nw64.Ctor__(_dafny.IntOfInt64(-261), _435_v10, _dafny.IntOfInt64(556), (_this).F21())
					_436_v11 = _nw64
					(globalState).F14 = !(!(p1))
					if (_434_v9).Select((Companion_Default___.SafeIndex(_436_v11.F27, _dafny.IntOfUint32((_434_v9).Cardinality()))).Uint32()).(bool) {
						r1 = false
						var _437_v12 _dafny.CodePoint
						_ = _437_v12
						_437_v12 = _dafny.CodePoint('w')
						_437_v12 = _dafny.CodePoint('i')
						(globalState).F14 = (p0).Cmp(p0) == 0
						var _438_v13 _dafny.Array
						_ = _438_v13
						var _len0_8 _dafny.Int = _dafny.One
						_ = _len0_8
						var _nw65 _dafny.Array
						_ = _nw65
						if _len0_8.Cmp(_dafny.Zero) == 0 {
							_nw65 = _dafny.NewArray(_len0_8)
						} else {
							var _init8 func(_dafny.Int) _dafny.Sequence = (func(_439_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_440_i2 _dafny.Int) _dafny.Sequence {
									return _439_v8
								}
							})(_433_v8)
							_ = _init8
							var _element0_8 = _init8(_dafny.Zero)
							_ = _element0_8
							_nw65 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
							(_nw65).ArraySet1(_element0_8, 0)
							var _nativeLen0_8 = (_len0_8).Int()
							_ = _nativeLen0_8
							for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
								(_nw65).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
							}
						}
						_438_v13 = _nw65
						var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_438_v13), 0))
						_ = _index58
						(_438_v13).ArraySet1(_433_v8, (_index58).Int())
						var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_438_v13), 0))
						_ = _index59
						(_438_v13).ArraySet1(_433_v8, (_index59).Int())
						(globalState).F17 = (_436_v11.F27).Plus(_436_v11.F27)
					} else {
						var _441_v14 _dafny.Int
						_ = _441_v14
						var _442_v15 bool
						_ = _442_v15
						var _443_v16 _dafny.Int
						_ = _443_v16
						var _444_v17 _dafny.Int
						_ = _444_v17
						var _out6 _dafny.Int
						_ = _out6
						var _out7 bool
						_ = _out7
						var _out8 _dafny.Int
						_ = _out8
						var _out9 _dafny.Int
						_ = _out9
						_out6, _out7, _out8, _out9 = (_this).M3((func() bool {
							if p1 {
								return (_this).F21()
							}
							return (_this).F21()
						})(), globalState)
						_441_v14 = _out6
						_442_v15 = _out7
						_443_v16 = _out8
						_444_v17 = _out9
						var _445_v18 _dafny.Set
						_ = _445_v18
						_445_v18 = _dafny.SetOf(_433_v8, _dafny.Companion_Sequence_.Concatenate(_433_v8, _433_v8), _433_v8, _433_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg19 _dafny.Int) interface{} {
								return coer19(arg19)
							}
						}(func(_446_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						})))
						_445_v18 = _445_v18
						(globalState).F5 = _443_v16
						var _447_v19 _dafny.MultiSet
						_ = _447_v19
						_447_v19 = _dafny.MultiSetOf(p1, (_this).F21(), _442_v15)
						var _448_v20 _dafny.Sequence
						_ = _448_v20
						_448_v20 = _dafny.SeqOf(_436_v11.F27, _443_v16)
						var _449_v21 *C1
						_ = _449_v21
						var _nw66 *C1 = New_C1_()
						_ = _nw66
						_nw66.Ctor__(_436_v11.F27, (_436_v11).F28(), (_436_v11).Fm1(_dafny.IntOfUint32((_434_v9).Cardinality()), _436_v11.F27, false, globalState), (_436_v11).Fm16(_447_v19, _448_v20, globalState))
						_449_v21 = _nw66
						var _450_v22 D7
						_ = _450_v22
						_450_v22 = Companion_D7_.Create_DC17_(!(p1))
						(globalState).F4 = (_450_v22).Dtor_cf25()
					}
					if !((_436_v11.F27).Cmp(p0) != 0) || (true) {
						var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))
						_ = _index60
						((_436_v11).F28()).ArraySet1(_436_v11.F27, (_index60).Int())
						var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))
						_ = _index61
						((_436_v11).F28()).ArraySet1(_436_v11.F27, (_index61).Int())
						var _451_v23 _dafny.Set
						_ = _451_v23
						_451_v23 = _dafny.SetOf(p0, p0, _436_v11.F27, _436_v11.F27, p0)
						var _452_v24 _dafny.Int
						_ = _452_v24
						var _453_v25 bool
						_ = _453_v25
						var _454_v26 _dafny.Int
						_ = _454_v26
						var _455_v27 _dafny.Int
						_ = _455_v27
						var _out10 _dafny.Int
						_ = _out10
						var _out11 bool
						_ = _out11
						var _out12 _dafny.Int
						_ = _out12
						var _out13 _dafny.Int
						_ = _out13
						_out10, _out11, _out12, _out13 = (_this).M3(((_451_v23).Cardinality()).Cmp(_436_v11.F27) < 0, globalState)
						_452_v24 = _out10
						_453_v25 = _out11
						_454_v26 = _out12
						_455_v27 = _out13
						var _456_v28 _dafny.Array
						_ = _456_v28
						var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
						_ = _nw67
						_456_v28 = _nw67
						var _457_v29 _dafny.MultiSet
						_ = _457_v29
						_457_v29 = _dafny.MultiSetOf(false)
						var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_456_v28), 0))
						_ = _index62
						(_456_v28).ArraySet1((_457_v29).IsSubsetOf(_dafny.MultiSetOf((_this).F21(), _453_v25)), (_index62).Int())
						var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_456_v28), 0))
						_ = _index63
						(_456_v28).ArraySet1((((_436_v11).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))).Int()).(_dafny.Int)).Cmp(_452_v24) < 0, (_index63).Int())
						var _458_v30 *C1
						_ = _458_v30
						var _nw68 *C1 = New_C1_()
						_ = _nw68
						_nw68.Ctor__(_436_v11.F27, (_436_v11).F28(), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(((_436_v11).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))).Int()).(_dafny.Int)), _436_v11.F27), (((_436_v11).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(992)) == 0)
						_458_v30 = _nw68
						var _459_v31 _dafny.MultiSet
						_ = _459_v31
						_459_v31 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_436_v11).Fm1(_454_v26, (_451_v23).Cardinality(), true, globalState), (_dafny.Zero).Minus(_dafny.IntOfUint32((_433_v8).Cardinality()))), _dafny.IntOfUint32((_433_v8).Cardinality()))
						var _460_v32 D10
						_ = _460_v32
						_460_v32 = Companion_D10_.Create_DC25_(_459_v31)
						var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))
						_ = _index64
						var _rhs55 bool = (_this).F21()
						_ = _rhs55
						var _rhs56 _dafny.MultiSet = ((_460_v32).Dtor_cf34()).Intersection(_459_v31)
						_ = _rhs56
						var _rhs57 _dafny.Int = (_dafny.Zero).Minus(_452_v24)
						_ = _rhs57
						var _rhs58 bool = p1
						_ = _rhs58
						var _lhs45 *GlobalState = globalState
						_ = _lhs45
						var _lhs46 _dafny.Array = (_436_v11).F28()
						_ = _lhs46
						var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen(((_436_v11).F28()), 0))
						_ = _lhs47
						var _lhs48 *GlobalState = globalState
						_ = _lhs48
						_lhs45.F18 = _rhs55
						_459_v31 = _rhs56
						(_lhs46).ArraySet1(_rhs57, (_lhs47).Int())
						_lhs48.F18 = _rhs58
					} else {
						(_436_v11).F27 = _dafny.IntOfUint32((_434_v9).Cardinality())
						var _461_v33 bool
						_ = _461_v33
						_461_v33 = !(_dafny.Companion_Sequence_.IsPrefixOf(_433_v8, _433_v8))
						_461_v33 = _461_v33
						var _462_v34 _dafny.Map
						_ = _462_v34
						_462_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(150), p0)
						var _463_v35 _dafny.Map
						_ = _463_v35
						_463_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F21())
						var _464_v36 _dafny.Sequence
						_ = _464_v36
						_464_v36 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21()), _463_v35, _463_v35)
						(globalState).F3 = (func() bool {
							if (_this).F21() {
								return !((_this).Fm0(_462_v34, _433_v8, _464_v36, globalState))
							}
							return (_this).F21()
						})()
						var _465_v37 _dafny.Map
						_ = _465_v37
						_465_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _436_v11.F27)
						_465_v37 = ((Companion_Default___.Fm27((_this).F20(), globalState)).Merge(_465_v37)).Merge(p2)
						var _466_v38 *C1
						_ = _466_v38
						var _nw69 *C1 = New_C1_()
						_ = _nw69
						_nw69.Ctor__((_436_v11.F27).Times(_dafny.IntOfInt64(687)), (_436_v11).F28(), p0, (_this).F21())
						_466_v38 = _nw69
					}
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		r1 = (func() bool {
			if !(!((_this).F21())) {
				return (_this).F21()
			}
			return p1
		})()
		var _467_v39 _dafny.Map
		_ = _467_v39
		_467_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F21())
		var _468_v40 _dafny.CodePoint
		_ = _468_v40
		_468_v40 = _dafny.CodePoint('k')
		var _469_v41 _dafny.Sequence
		_ = _469_v41
		_469_v41 = _dafny.UnicodeSeqOfUtf8Bytes("xuposfy")
		var _470_v42 _dafny.MultiSet
		_ = _470_v42
		_470_v42 = _dafny.MultiSetOf(p0)
		var _471_v43 _dafny.Sequence
		_ = _471_v43
		_471_v43 = _dafny.SeqOf(p0)
		var _472_v44 _dafny.Set
		_ = _472_v44
		_472_v44 = _dafny.SetOf(_dafny.IntOfInt64(862), p0)
		var _473_v45 _dafny.Map
		_ = _473_v45
		_473_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v40, (_472_v44).Cardinality())
		var _474_v46 _dafny.Array
		_ = _474_v46
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_9
		var _nw70 _dafny.Array
		_ = _nw70
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw70 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Int = func(_475_i4 _dafny.Int) _dafny.Int {
				return (_475_i4).Plus((_this).F20())
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw70 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw70).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw70).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_474_v46 = _nw70
		var _476_v47 *C1
		_ = _476_v47
		var _nw71 *C1 = New_C1_()
		_ = _nw71
		_nw71.Ctor__((_473_v45).Cardinality(), _474_v46, p0, p1)
		_476_v47 = _nw71
		var _477_v48 _dafny.Set
		_ = _477_v48
		_477_v48 = _dafny.SetOf(_476_v47, _476_v47)
		var _478_v49 _dafny.Map
		_ = _478_v49
		_478_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F21())
		var _479_v50 _dafny.Array
		_ = _479_v50
		var _nwElement0_17 _dafny.Int = _dafny.IntOfUint32((_469_v41).Cardinality())
		_ = _nwElement0_17
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(28))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_17, 0)
		(_nw72).ArraySet1(p0, 1)
		(_nw72).ArraySet1((_470_v42).Cardinality(), 2)
		(_nw72).ArraySet1((_this).F20(), 3)
		(_nw72).ArraySet1((_471_v43).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_471_v43).Cardinality()))).Uint32()).(_dafny.Int), 4)
		(_nw72).ArraySet1(p0, 5)
		(_nw72).ArraySet1((_this).F20(), 6)
		(_nw72).ArraySet1(p0, 7)
		(_nw72).ArraySet1((_this).F20(), 8)
		(_nw72).ArraySet1((_467_v39).Cardinality(), 9)
		(_nw72).ArraySet1(p0, 10)
		(_nw72).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_477_v48).Cardinality())).Cardinality()), 11)
		(_nw72).ArraySet1(p0, 12)
		(_nw72).ArraySet1(_476_v47.F27, 13)
		(_nw72).ArraySet1((_478_v49).Cardinality(), 14)
		(_nw72).ArraySet1(_476_v47.F27, 15)
		(_nw72).ArraySet1(_476_v47.F27, 16)
		(_nw72).ArraySet1(_dafny.IntOfUint32((_469_v41).Cardinality()), 17)
		(_nw72).ArraySet1((_this).F20(), 18)
		(_nw72).ArraySet1(_476_v47.F27, 19)
		(_nw72).ArraySet1((_this).F20(), 20)
		(_nw72).ArraySet1(_476_v47.F27, 21)
		(_nw72).ArraySet1(_476_v47.F27, 22)
		(_nw72).ArraySet1((_this).F20(), 23)
		(_nw72).ArraySet1(p0, 24)
		(_nw72).ArraySet1((_this).F20(), 25)
		(_nw72).ArraySet1((_473_v45).Cardinality(), 26)
		(_nw72).ArraySet1(_dafny.IntOfUint32((_471_v43).Cardinality()), 27)
		_479_v50 = _nw72
		var _480_v51 *C1
		_ = _480_v51
		var _nw73 *C1 = New_C1_()
		_ = _nw73
		_nw73.Ctor__((Companion_D10_.Create_DC26_(_dafny.IntOfInt64(545))).Dtor_cf35(), _474_v46, p0, p1)
		_480_v51 = _nw73
		var _481_v52 _dafny.Map
		_ = _481_v52
		_481_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v40, _476_v47)
		var _482_v53 _dafny.Sequence
		_ = _482_v53
		_482_v53 = _dafny.SeqOf(_481_v52, _481_v52, _481_v52)
		var _483_v54 D10
		_ = _483_v54
		_483_v54 = Companion_D10_.Create_DC27_((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_467_v39).Cardinality()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf((_467_v39).Cardinality())).Cardinality()))).Uint32(), p0))).Cardinality(), _482_v53, Companion_Default___.Fm17((_this).F21(), globalState))
		var _484_v55 _dafny.Sequence
		_ = _484_v55
		_484_v55 = _471_v43
		var _pat_let_tv12 = _480_v51
		_ = _pat_let_tv12
		var _pat_let_tv13 = p1
		_ = _pat_let_tv13
		var _pat_let_tv14 = _484_v55
		_ = _pat_let_tv14
		var _pat_let_tv15 = globalState
		_ = _pat_let_tv15
		var _pat_let_tv16 = globalState
		_ = _pat_let_tv16
		var _source7 D10 = func(_pat_let17_0 D10) D10 {
			return func(_485_dt__update__tmp_h0 D10) D10 {
				return func(_pat_let18_0 _dafny.Int) D10 {
					return func(_486_dt__update_hcf36_h0 _dafny.Int) D10 {
						return func(_pat_let19_0 _dafny.CodePoint) D10 {
							return func(_487_dt__update_hcf38_h0 _dafny.CodePoint) D10 {
								return Companion_D10_.Create_DC27_(_486_dt__update_hcf36_h0, (_485_dt__update__tmp_h0).Dtor_cf37(), _487_dt__update_hcf38_h0)
							}(_pat_let19_0)
						}(Companion_Default___.Fm17((_pat_let_tv12).Fm16(_dafny.MultiSetOf(true, !(_pat_let_tv13)), _pat_let_tv14, _pat_let_tv15), _pat_let_tv16))
					}(_pat_let18_0)
				}((_this).F20())
			}(_pat_let17_0)
		}(_483_v54)
		_ = _source7
		if _source7.Is_DC26() {
			var _488___mcc_h0 _dafny.Int = _source7.Get_().(D10_DC26).Cf35
			_ = _488___mcc_h0
			var _489_cf35 _dafny.Int = _488___mcc_h0
			_ = _489_cf35
			(globalState).F18 = (_dafny.IntOfInt64(350)).Cmp(p0) <= 0
			var _490_v56 _dafny.Array
			_ = _490_v56
			var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw74
			_490_v56 = _nw74
			var _491_v57 _dafny.Map
			_ = _491_v57
			_491_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_476_v47.F27), _476_v47.F27)
			var _492_v58 _dafny.Sequence
			_ = _492_v58
			_492_v58 = _dafny.SeqOf(_478_v49)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))
			_ = _index65
			(_490_v56).ArraySet1((_this).Fm0(_491_v57, _469_v41, _dafny.Companion_Sequence_.Update(_492_v58, (Companion_Default___.SafeIndex(_476_v47.F27, _dafny.IntOfUint32((_492_v58).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p1)), globalState), (_index65).Int())
			var _493_v59 _dafny.Set
			_ = _493_v59
			_493_v59 = _dafny.SetOf((_this).F21(), p1)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))
			_ = _index66
			((_480_v51).F28()).ArraySet1((Companion_Default___.Fm24(_476_v47.F27, !(p1), _493_v59, p1, globalState)).Dtor_cf17(), (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_490_v56), 0))
			_ = _index67
			(_490_v56).ArraySet1(p1, (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))
			_ = _index68
			(_474_v46).ArraySet1(_480_v51.F27, (_index68).Int())
			var _494_v60 _dafny.MultiSet
			_ = _494_v60
			_494_v60 = _dafny.MultiSetOf(_490_v56, _490_v56)
			var _495_v61 _dafny.Map
			_ = _495_v61
			_495_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_472_v44, _480_v51)
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))
			_ = _index69
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))
			_ = _index70
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_490_v56), 0))
			_ = _index71
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))
			_ = _index72
			var _rhs59 bool = (_this).F21()
			_ = _rhs59
			var _rhs60 _dafny.Int = _489_cf35
			_ = _rhs60
			var _rhs61 bool = !(_495_v61).Equals(_495_v61)
			_ = _rhs61
			var _rhs62 _dafny.Int = _480_v51.F27
			_ = _rhs62
			var _rhs63 _dafny.MultiSet = ((_494_v60).Union(_494_v60)).Intersection(_494_v60)
			_ = _rhs63
			var _lhs49 _dafny.Array = _490_v56
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = (_480_v51).F28()
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))
			_ = _lhs52
			var _lhs53 _dafny.Array = _490_v56
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_490_v56), 0))
			_ = _lhs54
			var _lhs55 _dafny.Array = _474_v46
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))
			_ = _lhs56
			(_lhs49).ArraySet1(_rhs59, (_lhs50).Int())
			(_lhs51).ArraySet1(_rhs60, (_lhs52).Int())
			(_lhs53).ArraySet1(_rhs61, (_lhs54).Int())
			(_lhs55).ArraySet1(_rhs62, (_lhs56).Int())
			_494_v60 = _rhs63
			var _496_v62 _dafny.Array
			_ = _496_v62
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
			_ = _nw75
			_496_v62 = _nw75
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))
			_ = _index73
			(_496_v62).ArraySet1(_469_v41, (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))
			_ = _index74
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))
			_ = _index75
			var _rhs64 _dafny.Sequence = _469_v41
			_ = _rhs64
			var _rhs65 _dafny.Int = Companion_Default___.SafeModuloInt(_476_v47.F27, _476_v47.F27)
			_ = _rhs65
			var _rhs66 bool = !((_this).F21())
			_ = _rhs66
			var _lhs57 _dafny.Array = _496_v62
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))
			_ = _lhs58
			var _lhs59 _dafny.Array = (_480_v51).F28()
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))
			_ = _lhs60
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			(_lhs57).ArraySet1(_rhs64, (_lhs58).Int())
			(_lhs59).ArraySet1(_rhs65, (_lhs60).Int())
			_lhs61.F3 = _rhs66
			if (func() bool {
				if (_478_v49).Contains((_490_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))).Int()).(bool)) {
					return (_478_v49).Get((_490_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))).Int()).(bool)).(bool)
				}
				return p1
			})() {
				var _497_v64 _dafny.Array
				_ = _497_v64
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_10
				var _nw76 _dafny.Array
				_ = _nw76
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw76 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Sequence = (func(_498_v44 _dafny.Set, _499_p2 _dafny.Map, _500_v57 _dafny.Map) func(_dafny.Int) _dafny.Sequence {
						return func(_501_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_498_v44, func() _dafny.Set {
								var _coll19 = _dafny.NewBuilder()
								_ = _coll19
								for _iter20 := _dafny.Iterate((_500_v57).Keys().Elements()); ; {
									_compr_19, _ok20 := _iter20()
									if !_ok20 {
										break
									}
									var _502_v63 _dafny.Int
									_502_v63 = interface{}(_compr_19).(_dafny.Int)
									if (_500_v57).Contains(_502_v63) {
										_coll19.Add((_502_v63).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("afyfapfi")).Cardinality())))
									}
								}
								return _coll19.ToSet()
							}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ml")).Cardinality()), (func() _dafny.Int {
								if (_499_p2).Contains((_this).F21()) {
									return (_499_p2).Get((_this).F21()).(_dafny.Int)
								}
								return (_this).F20()
							})()))
						}
					})(_472_v44, p2, _491_v57)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw76 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw76).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw76).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_497_v64 = _nw76
				var _503_v65 _dafny.MultiSet
				_ = _503_v65
				_503_v65 = _dafny.MultiSetOf(p1, true)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_497_v64), 0))
				_ = _index76
				(_497_v64).ArraySet1(_dafny.SeqOf(_dafny.SetOf((func() _dafny.Int {
					if (_503_v65).Contains((_this).F21()) {
						return (_503_v65).Multiplicity((_this).F21())
					}
					return _dafny.IntOfInt64(684)
				})(), (_480_v51).Fm1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-418))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_504_cf35 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_505_i6 _dafny.Int) _dafny.Int {
						return _504_cf35
					}
				})(_489_cf35)))).Cardinality()), _dafny.IntOfInt64(-357), (_490_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))).Int()).(bool), globalState)), _472_v44), (_index76).Int())
				var _506_v66 _dafny.Sequence
				_ = _506_v66
				_506_v66 = _dafny.SeqOf(_dafny.SetOf((p2).Cardinality(), (_dafny.Zero).Minus(_489_cf35), ((_480_v51).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_480_v51).F28()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_507_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_508_i7 _dafny.Int) _dafny.CodePoint {
						return _507_v40
					}
				})(_468_v40)))).Cardinality())))
				var _509_v67 _dafny.Sequence
				_ = _509_v67
				_509_v67 = _dafny.SeqOf(_506_v66, _506_v66, _506_v66, _dafny.SeqOf((_506_v66).Select((Companion_Default___.SafeIndex(_480_v51.F27, _dafny.IntOfUint32((_506_v66).Cardinality()))).Uint32()).(_dafny.Set)), _506_v66)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_497_v64), 0))
				_ = _index77
				(_497_v64).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_509_v67).Select((Companion_Default___.SafeIndex(_480_v51.F27, _dafny.IntOfUint32((_509_v67).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_476_v47.F27, _dafny.IntOfUint32(((_509_v67).Select((Companion_Default___.SafeIndex(_480_v51.F27, _dafny.IntOfUint32((_509_v67).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.SetOf(_dafny.IntOfUint32(((_496_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))).Int()).(_dafny.Sequence)).Cardinality()))), _506_v66), (_index77).Int())
				var _510_v68 _dafny.Map
				_ = _510_v68
				_510_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_490_v56, (_489_cf35).Plus((_472_v44).Cardinality()))
				_510_v68 = (_510_v68).Update(_490_v56, _476_v47.F27)
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))
				_ = _index78
				(_496_v62).ArraySet1(_469_v41, (_index78).Int())
				var _511_v69 _dafny.MultiSet
				_ = _511_v69
				_511_v69 = _dafny.MultiSetOf(Companion_Default___.Fm17(p1, globalState))
				(globalState).F5 = (_476_v47.F27).Minus(((_467_v39).Cardinality()).Times((func() _dafny.Int {
					if (_511_v69).Contains(_468_v40) {
						return (_511_v69).Multiplicity(_468_v40)
					}
					return (_this).F20()
				})()))
				var _512_v70 _dafny.Sequence
				_ = _512_v70
				_512_v70 = _dafny.SeqOf(p1, !((_480_v51.F27).Cmp(_480_v51.F27) > 0), (_503_v65).IsSubsetOf(_503_v65))
				(globalState).F18 = (_512_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-639), _dafny.IntOfUint32((_512_v70).Cardinality()))).Uint32()).(bool)
			} else {
				(globalState).F5 = (_471_v43).Select((Companion_Default___.SafeIndex(((_474_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))).Int()).(_dafny.Int)).Minus(_489_cf35), _dafny.IntOfUint32((_471_v43).Cardinality()))).Uint32()).(_dafny.Int)
				var _513_v71 _dafny.Sequence
				_ = _513_v71
				_513_v71 = _dafny.SeqOf((_476_v47).F28())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))
				_ = _index79
				(_490_v56).ArraySet1((_dafny.IntOfUint32((_513_v71).Cardinality())).Cmp(_480_v51.F27) == 0, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))
				_ = _index80
				(_474_v46).ArraySet1((_480_v51).Fm1(_480_v51.F27, _489_cf35, !(p1), globalState), (_index80).Int())
				var _514_v72 D9
				_ = _514_v72
				_514_v72 = Companion_D9_.Create_DC23_((_474_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_474_v46), 0))).Int()).(_dafny.Int), p1)
				var _515_v73 D9
				_ = _515_v73
				_515_v73 = Companion_D9_.Create_DC24_(_514_v72)
				var _516_v74 D9
				_ = _516_v74
				_516_v74 = Companion_D9_.Create_DC24_(_515_v73)
				_516_v74 = _516_v74
				var _517_v75 *C1
				_ = _517_v75
				var _nw77 *C1 = New_C1_()
				_ = _nw77
				_nw77.Ctor__((_dafny.Zero).Minus(_476_v47.F27), (_480_v51).F28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_496_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_496_v62), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("ien"))).Cardinality()), (_490_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_490_v56), 0))).Int()).(bool))
				_517_v75 = _nw77
			}
		} else if _source7.Is_DC27() {
			var _518___mcc_h1 _dafny.Int = _source7.Get_().(D10_DC27).Cf36
			_ = _518___mcc_h1
			var _519___mcc_h2 _dafny.Sequence = _source7.Get_().(D10_DC27).Cf37
			_ = _519___mcc_h2
			var _520___mcc_h3 _dafny.CodePoint = _source7.Get_().(D10_DC27).Cf38
			_ = _520___mcc_h3
			var _521_cf38 _dafny.CodePoint = _520___mcc_h3
			_ = _521_cf38
			var _522_cf37 _dafny.Sequence = _519___mcc_h2
			_ = _522_cf37
			var _523_cf36 _dafny.Int = _518___mcc_h1
			_ = _523_cf36
			(globalState).F18 = p1
			var _524_v76 _dafny.Array
			_ = _524_v76
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw78
			_524_v76 = _nw78
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v76), 0))
			_ = _index81
			(_524_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_469_v41, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_469_v41).Cardinality()), _dafny.IntOfUint32((_469_v41).Cardinality()))).Uint32(), _dafny.CodePoint('a')), _469_v41), (_index81).Int())
			var _525_v77 D6
			_ = _525_v77
			_525_v77 = Companion_D6_.Create_DC13_(_469_v41)
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v76), 0))
			_ = _index82
			(_524_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ndkkvypt"), _dafny.UnicodeSeqOfUtf8Bytes("dat")), (_525_v77).Dtor_cf21()), (_index82).Int())
			var _526_v78 _dafny.Sequence
			_ = _526_v78
			_526_v78 = _dafny.SeqOf((_this).F21(), (_this).F21(), (_this).F21())
			var _527_v79 _dafny.MultiSet
			_ = _527_v79
			_527_v79 = _dafny.MultiSetOf(!((_526_v78).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_471_v43).Cardinality()), _dafny.IntOfUint32((_526_v78).Cardinality()))).Uint32()).(bool)), (_this).F21())
			var _528_v80 _dafny.Sequence
			_ = _528_v80
			_528_v80 = _dafny.SeqOf((_480_v51).Fm16(_527_v79, _dafny.SeqOf(_523_cf36, _480_v51.F27), globalState), p1)
			var _529_v81 _dafny.Sequence
			_ = _529_v81
			_529_v81 = _dafny.SeqOf(_478_v49, (_478_v49).Update((_this).F21(), p1), (func() _dafny.Map {
				if !((_528_v80).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.IntOfUint32((_528_v80).Cardinality()))).Uint32()).(bool)) {
					return _478_v49
				}
				return _478_v49
			})())
			var _530_v82 D7
			_ = _530_v82
			_530_v82 = Companion_D7_.Create_DC18_(_471_v43, _476_v47.F27)
			var _531_v83 _dafny.Map
			_ = _531_v83
			_531_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_v82, (true) && (p1))
			var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21()), _478_v49), _529_v81), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_476_v47.F27), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21()), _478_v49), _529_v81)).Cardinality()))).Uint32(), _478_v49)
			_ = _rhs67
			var _rhs68 bool = (func() bool {
				if (_531_v83).Contains(_530_v82) {
					return (_531_v83).Get(_530_v82).(bool)
				}
				return p1
			})()
			_ = _rhs68
			var _rhs69 _dafny.Int = _dafny.IntOfUint32((_471_v43).Cardinality())
			_ = _rhs69
			var _lhs62 *GlobalState = globalState
			_ = _lhs62
			var _lhs63 *C1 = _476_v47
			_ = _lhs63
			_529_v81 = _rhs67
			_lhs62.F3 = _rhs68
			_lhs63.F27 = _rhs69
			var _532_v84 _dafny.Array
			_ = _532_v84
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_11
			var _nw79 _dafny.Array
			_ = _nw79
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw79 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = func(_533_i8 _dafny.Int) bool {
					return (_this).F21()
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw79 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw79).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw79).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_532_v84 = _nw79
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_532_v84), 0))
			_ = _index83
			(_532_v84).ArraySet1((_this).F21(), (_index83).Int())
			var _534_v85 _dafny.Map
			_ = _534_v85
			_534_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _476_v47.F27)
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_532_v84), 0))
			_ = _index84
			(_532_v84).ArraySet1((_480_v51).Fm0(_534_v85, _469_v41, _529_v81, globalState), (_index84).Int())
		} else {
			var _535___mcc_h4 _dafny.MultiSet = _source7.Get_().(D10_DC25).Cf34
			_ = _535___mcc_h4
			var _536_cf34 _dafny.MultiSet = _535___mcc_h4
			_ = _536_cf34
			var _537_v86 _dafny.Array
			_ = _537_v86
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_12
			var _nw80 _dafny.Array
			_ = _nw80
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw80 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = func(_538_i9 _dafny.Int) bool {
					return (_this).F21()
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw80 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw80).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw80).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_537_v86 = _nw80
			var _539_v87 _dafny.MultiSet
			_ = _539_v87
			_539_v87 = _dafny.MultiSetOf(_537_v86, _537_v86)
			var _source8 D2 = Companion_D2_.Create_DC2_((_dafny.MultiSetOf(_537_v86)).Difference(_539_v87))
			_ = _source8
			if _source8.Is_DC3() {
				var _540___mcc_h5 _dafny.Int = _source8.Get_().(D2_DC3).Cf3
				_ = _540___mcc_h5
				var _541___mcc_h6 bool = _source8.Get_().(D2_DC3).Cf4
				_ = _541___mcc_h6
				var _542___mcc_h7 bool = _source8.Get_().(D2_DC3).Cf5
				_ = _542___mcc_h7
				var _543___mcc_h8 _dafny.Int = _source8.Get_().(D2_DC3).Cf6
				_ = _543___mcc_h8
				var _544___mcc_h9 bool = _source8.Get_().(D2_DC3).Cf7
				_ = _544___mcc_h9
				var _545_cf7 bool = _544___mcc_h9
				_ = _545_cf7
				var _546_cf6 _dafny.Int = _543___mcc_h8
				_ = _546_cf6
				var _547_cf5 bool = _542___mcc_h7
				_ = _547_cf5
				var _548_cf4 bool = _541___mcc_h6
				_ = _548_cf4
				var _549_cf3 _dafny.Int = _540___mcc_h5
				_ = _549_cf3
				(globalState).F4 = _548_cf4
				(_476_v47).F27 = _480_v51.F27
				_546_cf6 = (_476_v47.F27).Plus(Companion_Default___.SafeDivisionInt(_480_v51.F27, _476_v47.F27))
				var _550_v88 *C0
				_ = _550_v88
				var _nw81 *C0 = New_C0_()
				_ = _nw81
				_nw81.Ctor__()
				_550_v88 = _nw81
			} else if _source8.Is_DC2() {
				var _551___mcc_h10 _dafny.MultiSet = _source8.Get_().(D2_DC2).Cf2
				_ = _551___mcc_h10
				var _552_cf2 _dafny.MultiSet = _551___mcc_h10
				_ = _552_cf2
				var _553_v89 _dafny.Sequence
				_ = _553_v89
				_553_v89 = _dafny.SeqOf(p1)
				var _554_v90 _dafny.Sequence
				_ = _554_v90
				_554_v90 = _dafny.SeqOf(_478_v49)
				var _555_v91 _dafny.MultiSet
				_ = _555_v91
				_555_v91 = _dafny.MultiSetOf((_this).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_480_v51).Fm1((_dafny.MultiSetOf(_476_v47.F27)).Cardinality(), (_this).Fm1(_dafny.IntOfUint32((_553_v89).Cardinality()), _476_v47.F27, (_this).F21(), globalState), p1, globalState), _dafny.IntOfUint32((_553_v89).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("folerbj"), _554_v90, globalState), (_this).F21())
				var _rhs70 _dafny.Int = ((_this).F20()).Times((_dafny.Zero).Minus((func() _dafny.Int {
					if (_555_v91).Contains(p1) {
						return (_555_v91).Multiplicity(p1)
					}
					return _480_v51.F27
				})()))
				_ = _rhs70
				var _rhs71 bool = p1
				_ = _rhs71
				var _lhs64 *C1 = _476_v47
				_ = _lhs64
				_lhs64.F27 = _rhs70
				r1 = _rhs71
				_478_v49 = (_478_v49).Update((func() bool {
					if p1 {
						return p1
					}
					return false
				})(), (_this).F21())
				(globalState).F18 = !(!(p1))
				var _556_v92 *C1
				_ = _556_v92
				var _nw82 *C1 = New_C1_()
				_ = _nw82
				_nw82.Ctor__((_this).F20(), (_476_v47).F28(), Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()), p1)
				_556_v92 = _nw82
			} else {
				var _557___mcc_h11 D2 = _source8.Get_().(D2_DC4).Cf8
				_ = _557___mcc_h11
				var _558_cf8 D2 = _557___mcc_h11
				_ = _558_cf8
				var _559_v93 *C1
				_ = _559_v93
				var _nw83 *C1 = New_C1_()
				_ = _nw83
				_nw83.Ctor__(_480_v51.F27, _474_v46, (_480_v51.F27).Times((_this).F20()), (_this).F21())
				_559_v93 = _nw83
				var _560_v94 _dafny.Array
				_ = _560_v94
				var _nwElement0_18 _dafny.Array = (_476_v47).F28()
				_ = _nwElement0_18
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(28))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_18, 0)
				(_nw84).ArraySet1((_559_v93).F28(), 1)
				(_nw84).ArraySet1(_479_v50, 2)
				(_nw84).ArraySet1((_480_v51).F28(), 3)
				(_nw84).ArraySet1((_480_v51).F28(), 4)
				(_nw84).ArraySet1(_479_v50, 5)
				(_nw84).ArraySet1(_479_v50, 6)
				(_nw84).ArraySet1((_559_v93).F28(), 7)
				(_nw84).ArraySet1((_559_v93).F28(), 8)
				(_nw84).ArraySet1((_476_v47).F28(), 9)
				(_nw84).ArraySet1((_476_v47).F28(), 10)
				(_nw84).ArraySet1((_559_v93).F28(), 11)
				(_nw84).ArraySet1(_474_v46, 12)
				(_nw84).ArraySet1((_480_v51).F28(), 13)
				(_nw84).ArraySet1(_479_v50, 14)
				(_nw84).ArraySet1((_476_v47).F28(), 15)
				(_nw84).ArraySet1((_476_v47).F28(), 16)
				(_nw84).ArraySet1((_559_v93).F28(), 17)
				(_nw84).ArraySet1(_479_v50, 18)
				(_nw84).ArraySet1((_559_v93).F28(), 19)
				(_nw84).ArraySet1(_479_v50, 20)
				(_nw84).ArraySet1((_480_v51).F28(), 21)
				(_nw84).ArraySet1(_479_v50, 22)
				(_nw84).ArraySet1(_474_v46, 23)
				(_nw84).ArraySet1((_476_v47).F28(), 24)
				(_nw84).ArraySet1((_476_v47).F28(), 25)
				(_nw84).ArraySet1(_474_v46, 26)
				(_nw84).ArraySet1((_480_v51).F28(), 27)
				_560_v94 = _nw84
				var _561_v95 _dafny.Sequence
				_ = _561_v95
				_561_v95 = _dafny.SeqOf(_560_v94)
				var _562_v96 _dafny.Array
				_ = _562_v96
				var _nwElement0_19 _dafny.Array = (_561_v95).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_536_cf34).Contains(_476_v47.F27) {
						return (_536_cf34).Multiplicity(_476_v47.F27)
					}
					return _dafny.IntOfInt64(592)
				})(), _dafny.IntOfUint32((_561_v95).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _nwElement0_19
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(29))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_19, 0)
				(_nw85).ArraySet1(_560_v94, 1)
				(_nw85).ArraySet1(_560_v94, 2)
				(_nw85).ArraySet1(_560_v94, 3)
				(_nw85).ArraySet1(_560_v94, 4)
				(_nw85).ArraySet1(_560_v94, 5)
				(_nw85).ArraySet1(_560_v94, 6)
				(_nw85).ArraySet1(_560_v94, 7)
				(_nw85).ArraySet1(_560_v94, 8)
				(_nw85).ArraySet1(_560_v94, 9)
				(_nw85).ArraySet1(_560_v94, 10)
				(_nw85).ArraySet1((func() _dafny.Array {
					if p1 {
						return _560_v94
					}
					return _560_v94
				})(), 11)
				(_nw85).ArraySet1(_560_v94, 12)
				(_nw85).ArraySet1(_560_v94, 13)
				(_nw85).ArraySet1(_560_v94, 14)
				(_nw85).ArraySet1(_560_v94, 15)
				(_nw85).ArraySet1((_561_v95).Select((Companion_Default___.SafeIndex(_476_v47.F27, _dafny.IntOfUint32((_561_v95).Cardinality()))).Uint32()).(_dafny.Array), 16)
				(_nw85).ArraySet1(_560_v94, 17)
				(_nw85).ArraySet1(_560_v94, 18)
				(_nw85).ArraySet1(_560_v94, 19)
				(_nw85).ArraySet1(_560_v94, 20)
				(_nw85).ArraySet1(_560_v94, 21)
				(_nw85).ArraySet1(_560_v94, 22)
				(_nw85).ArraySet1(_560_v94, 23)
				(_nw85).ArraySet1(_560_v94, 24)
				(_nw85).ArraySet1(_560_v94, 25)
				(_nw85).ArraySet1(_560_v94, 26)
				(_nw85).ArraySet1(_560_v94, 27)
				(_nw85).ArraySet1(_560_v94, 28)
				_562_v96 = _nw85
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_562_v96), 0))
				_ = _index85
				(_562_v96).ArraySet1(_560_v94, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_562_v96), 0))
				_ = _index86
				(_562_v96).ArraySet1(_560_v94, (_index86).Int())
				var _563_v97 D7
				_ = _563_v97
				_563_v97 = Companion_D7_.Create_DC17_((p1))
				_563_v97 = _563_v97
				var _564_v98 _dafny.Array
				_ = _564_v98
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_13
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_565_v41 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_566_i10 _dafny.Int) _dafny.Sequence {
							return _565_v41
						}
					})(_469_v41)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw86 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw86).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw86).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_564_v98 = _nw86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_564_v98), 0))
				_ = _index87
				(_564_v98).ArraySet1(_469_v41, (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_564_v98), 0))
				_ = _index88
				(_564_v98).ArraySet1(_469_v41, (_index88).Int())
			}
			(globalState).F5 = (_this).F20()
			if !((_this).F21()) || ((_this).F21()) {
				var _567_v99 _dafny.Array
				_ = _567_v99
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw87
				_567_v99 = _nw87
				var _568_v100 _dafny.Array
				_ = _568_v100
				var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw88
				_568_v100 = _nw88
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_567_v99), 0))
				_ = _index89
				(_567_v99).ArraySet1(_568_v100, (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_567_v99), 0))
				_ = _index90
				(_567_v99).ArraySet1(_568_v100, (_index90).Int())
				var _569_v101 D8
				_ = _569_v101
				_569_v101 = Companion_D8_.Create_DC20_()
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_568_v100), 0))
				_ = _index91
				(_568_v100).ArraySet1((_480_v51).F28(), (_index91).Int())
				var _570_v102 _dafny.Sequence
				_ = _570_v102
				_570_v102 = _dafny.SeqOf((_480_v51).Fm16(_dafny.MultiSetOf(p1), _484_v55, globalState), !(p1))
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_568_v100), 0))
				_ = _index92
				var _rhs72 D8 = _569_v101
				_ = _rhs72
				var _rhs73 bool = p1
				_ = _rhs73
				var _rhs74 _dafny.Array = (_476_v47).F28()
				_ = _rhs74
				var _rhs75 bool = (_this).F21()
				_ = _rhs75
				var _rhs76 _dafny.Sequence = _570_v102
				_ = _rhs76
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				var _lhs66 _dafny.Array = _568_v100
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_568_v100), 0))
				_ = _lhs67
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				_569_v101 = _rhs72
				_lhs65.F18 = _rhs73
				(_lhs66).ArraySet1(_rhs74, (_lhs67).Int())
				_lhs68.F3 = _rhs75
				_570_v102 = _rhs76
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_474_v46), 0))
				_ = _index93
				(_474_v46).ArraySet1((_480_v51.F27).Plus(_476_v47.F27), (_index93).Int())
				var _571_v103 D7
				_ = _571_v103
				_571_v103 = Companion_D7_.Create_DC17_(!((_this).F21()))
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_474_v46), 0))
				_ = _index94
				(_474_v46).ArraySet1(Companion_Default___.SafeDivisionInt(_476_v47.F27, (_480_v51).Fm1((_this).F20(), _476_v47.F27, !((_571_v103).Dtor_cf25()), globalState)), (_index94).Int())
				(globalState).F5 = _480_v51.F27
				var _pat_let_tv17 = _482_v53
				_ = _pat_let_tv17
				var _572_v104 _dafny.Map
				_ = _572_v104
				_572_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let20_0 D10) D10 {
					return func(_573_dt__update__tmp_h1 D10) D10 {
						return func(_pat_let21_0 _dafny.Sequence) D10 {
							return func(_574_dt__update_hcf37_h0 _dafny.Sequence) D10 {
								return Companion_D10_.Create_DC27_((_573_dt__update__tmp_h1).Dtor_cf36(), _574_dt__update_hcf37_h0, (_573_dt__update__tmp_h1).Dtor_cf38())
							}(_pat_let21_0)
						}(_pat_let_tv17)
					}(_pat_let20_0)
				}(_483_v54), !(!((_this).F21())))
				(_480_v51).F27 = (func() _dafny.Int {
					if (_this).F21() {
						return _476_v47.F27
					}
					return (_572_v104).Cardinality()
				})()
			} else {
				(globalState).F4 = (_this).F21()
				var _575_v105 _dafny.Array
				_ = _575_v105
				var _nwElement0_20 _dafny.Sequence = Companion_Default___.Fm28(_480_v51.F27, globalState)
				_ = _nwElement0_20
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(10))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_20, 0)
				(_nw89).ArraySet1(_471_v43, 1)
				(_nw89).ArraySet1(_471_v43, 2)
				(_nw89).ArraySet1(_471_v43, 3)
				(_nw89).ArraySet1(_471_v43, 4)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Update(_471_v43, (Companion_Default___.SafeIndex(_480_v51.F27, _dafny.IntOfUint32((_471_v43).Cardinality()))).Uint32(), _476_v47.F27), 5)
				(_nw89).ArraySet1(_471_v43, 6)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_471_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_576_v51 *C1) func(_dafny.Int) _dafny.Int {
					return func(_577_i11 _dafny.Int) _dafny.Int {
						return _576_v51.F27
					}
				})(_480_v51)))), 7)
				(_nw89).ArraySet1(_471_v43, 8)
				(_nw89).ArraySet1(_dafny.SeqOf(p0, (_this).F20(), _476_v47.F27), 9)
				_575_v105 = _nw89
				var _578_v106 _dafny.Map
				_ = _578_v106
				_578_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_480_v51.F27, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}((func(_579_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_580_i12 _dafny.Int) _dafny.CodePoint {
						return _579_v40
					}
				})(_468_v40)))).Cardinality()))
				var _581_v107 _dafny.Sequence
				_ = _581_v107
				_581_v107 = _dafny.SeqOf(_478_v49, _478_v49)
				var _582_v108 _dafny.Map
				_ = _582_v108
				_582_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_476_v47).Fm0(_578_v106, _469_v41, _581_v107, globalState)), _dafny.CodePoint('q'))
				var _rhs77 _dafny.Int = (_480_v51).Fm1(_480_v51.F27, _480_v51.F27, p1, globalState)
				_ = _rhs77
				var _rhs78 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), _480_v51.F27)
				_ = _rhs78
				var _rhs79 bool = (func() bool {
					if (_this).F21() {
						return p1
					}
					return ((_this).F20()).Cmp(_476_v47.F27) >= 0
				})()
				_ = _rhs79
				var _rhs80 _dafny.Array = (func() _dafny.Array {
					if p1 {
						return _575_v105
					}
					return _575_v105
				})()
				_ = _rhs80
				var _rhs81 _dafny.Map = _582_v108
				_ = _rhs81
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 *C1 = _480_v51
				_ = _lhs70
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				_lhs69.F17 = _rhs77
				_lhs70.F27 = _rhs78
				_lhs71.F18 = _rhs79
				_575_v105 = _rhs80
				_582_v108 = _rhs81
				(globalState).F18 = p1
				var _583_v109 _dafny.Map
				_ = _583_v109
				_583_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _478_v49)
				(globalState).F5 = (((_583_v109).Cardinality()).Minus(_dafny.IntOfUint32((Companion_Default___.Fm26(p1, globalState)).Cardinality()))).Minus(_476_v47.F27)
				var _584_v110 _dafny.Array
				_ = _584_v110
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_14
				var _nw90 _dafny.Array
				_ = _nw90
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw90 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Map = (func(_585_p2 _dafny.Map, _586_p1 bool, _587_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
						return func(_588_i13 _dafny.Int) _dafny.Map {
							return (_585_p2).Update(_586_p1, _dafny.IntOfUint32((_dafny.SeqOf(_587_v40, _587_v40)).Cardinality()))
						}
					})(p2, p1, _468_v40)
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw90 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw90).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw90).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_584_v110 = _nw90
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_15
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Map = (func(_589_p2 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_590_i14 _dafny.Int) _dafny.Map {
							return (_589_p2).Merge(_589_p2)
						}
					})(p2)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw91 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw91).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw91).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_584_v110 = _nw91
			}
			(_480_v51).F27 = (_472_v44).Cardinality()
		}
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_474_v46), 0))
		_ = _index95
		(_474_v46).ArraySet1(p0, (_index95).Int())
		var _591_v111 D6
		_ = _591_v111
		_591_v111 = Companion_D6_.Create_DC13_(_469_v41)
		var _592_v112 _dafny.Array
		_ = _592_v112
		var _nwElement0_21 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_469_v41, (Companion_Default___.SafeIndex(_480_v51.F27, _dafny.IntOfUint32((_469_v41).Cardinality()))).Uint32(), _468_v40)
		_ = _nwElement0_21
		var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(22))
		_ = _nw92
		(_nw92).ArraySet1(_nwElement0_21, 0)
		(_nw92).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("miga"), 1)
		(_nw92).ArraySet1(_469_v41, 2)
		(_nw92).ArraySet1(_469_v41, 3)
		(_nw92).ArraySet1(_469_v41, 4)
		(_nw92).ArraySet1(_469_v41, 5)
		(_nw92).ArraySet1(_469_v41, 6)
		(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_469_v41, _469_v41), 7)
		(_nw92).ArraySet1(_469_v41, 8)
		(_nw92).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_593_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_594_i15 _dafny.Int) _dafny.CodePoint {
				return _593_v40
			}
		})(_468_v40))), 9)
		(_nw92).ArraySet1(_469_v41, 10)
		(_nw92).ArraySet1(_469_v41, 11)
		(_nw92).ArraySet1(Companion_Default___.Fm26(true, globalState), 12)
		(_nw92).ArraySet1(_469_v41, 13)
		(_nw92).ArraySet1(_469_v41, 14)
		(_nw92).ArraySet1((_591_v111).Dtor_cf21(), 15)
		(_nw92).ArraySet1(_469_v41, 16)
		(_nw92).ArraySet1(_469_v41, 17)
		(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_469_v41, _469_v41), 18)
		(_nw92).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_595_i16 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), 19)
		(_nw92).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qifnsfgu"), 20)
		(_nw92).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(922))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_596_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_597_i17 _dafny.Int) _dafny.CodePoint {
				return _596_v40
			}
		})(_468_v40))), 21)
		_592_v112 = _nw92
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_592_v112), 0))
		_ = _index96
		(_592_v112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bwsydd"), _469_v41), (_index96).Int())
		var _598_v113 _dafny.Array
		_ = _598_v113
		var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
		_ = _nw93
		_598_v113 = _nw93
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_598_v113), 0))
		_ = _index97
		(_598_v113).ArraySet1((p0).Cmp(_480_v51.F27) == 0, (_index97).Int())
		var _599_v114 *C0
		_ = _599_v114
		var _nw94 *C0 = New_C0_()
		_ = _nw94
		_nw94.Ctor__()
		_599_v114 = _nw94
		var _600_v115 _dafny.MultiSet
		_ = _600_v115
		_600_v115 = _dafny.MultiSetOf(_599_v114)
		var _601_v116 _dafny.Sequence
		_ = _601_v116
		_601_v116 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_600_v115).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _480_v51.F27))
		var _602_v117 _dafny.Sequence
		_ = _602_v117
		_602_v117 = _dafny.SeqOf(_469_v41)
		var _603_v118 D11
		_ = _603_v118
		_603_v118 = Companion_D11_.Create_DC28_(_592_v112)
		var _604_v119 _dafny.Map
		_ = _604_v119
		_604_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(144), _dafny.IntOfUint32((_469_v41).Cardinality()))
		var _605_v120 _dafny.Sequence
		_ = _605_v120
		_605_v120 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_476_v47).Fm16(_dafny.MultiSetOf((_this).F21()), _484_v55, globalState), p1), _478_v49, _478_v49)
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_474_v46), 0))
		_ = _index98
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_592_v112), 0))
		_ = _index99
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_598_v113), 0))
		_ = _index100
		var _rhs82 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_601_v116, _dafny.Companion_Sequence_.Concatenate(_601_v116, _601_v116))).Cardinality())
		_ = _rhs82
		var _rhs83 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_602_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_469_v41).Cardinality()), _dafny.IntOfUint32((_602_v117).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_469_v41, _469_v41))
		_ = _rhs83
		var _rhs84 _dafny.Array = (_603_v118).Dtor_cf39()
		_ = _rhs84
		var _rhs85 bool = (_this).Fm0(((_604_v119).Update((_this).F20(), _476_v47.F27)).Merge(_604_v119), _469_v41, _dafny.Companion_Sequence_.Update(_605_v120, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.IntOfUint32((_605_v120).Cardinality()))).Uint32(), _478_v49), globalState)
		_ = _rhs85
		var _lhs72 _dafny.Array = _474_v46
		_ = _lhs72
		var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_474_v46), 0))
		_ = _lhs73
		var _lhs74 _dafny.Array = _592_v112
		_ = _lhs74
		var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_592_v112), 0))
		_ = _lhs75
		var _lhs76 _dafny.Array = _598_v113
		_ = _lhs76
		var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_598_v113), 0))
		_ = _lhs77
		(_lhs72).ArraySet1(_rhs82, (_lhs73).Int())
		(_lhs74).ArraySet1(_rhs83, (_lhs75).Int())
		_592_v112 = _rhs84
		(_lhs76).ArraySet1(_rhs85, (_lhs77).Int())
		var _606_v122 _dafny.Sequence
		_ = _606_v122
		_606_v122 = _dafny.SeqOf(!((_598_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_598_v113), 0))).Int()).(bool)), p1)
		var _607_v123 _dafny.Map
		_ = _607_v123
		_607_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_606_v122).Select((Companion_Default___.SafeIndex((_this).Fm1(_476_v47.F27, (_dafny.MultiSetFromSeq(_602_v117)).Cardinality(), p1, globalState), _dafny.IntOfUint32((_606_v122).Cardinality()))).Uint32()).(bool), _604_v119)
		var _608_v124 _dafny.Array
		_ = _608_v124
		var _nwElement0_22 _dafny.Map = _604_v119
		_ = _nwElement0_22
		var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(11))
		_ = _nw95
		(_nw95).ArraySet1(_nwElement0_22, 0)
		(_nw95).ArraySet1(_604_v119, 1)
		(_nw95).ArraySet1(_604_v119, 2)
		(_nw95).ArraySet1((func() _dafny.Map {
			if p1 {
				return _604_v119
			}
			return _604_v119
		})(), 3)
		(_nw95).ArraySet1((_604_v119).Update(_dafny.IntOfInt64(-384), _476_v47.F27), 4)
		(_nw95).ArraySet1(_604_v119, 5)
		(_nw95).ArraySet1((_604_v119).Merge(_604_v119), 6)
		(_nw95).ArraySet1(_604_v119, 7)
		(_nw95).ArraySet1((func() _dafny.Map {
			if false {
				return _604_v119
			}
			return func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-508), _dafny.IntOfInt64(-87))); ; {
					_compr_20, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _609_v121 _dafny.Int
					_609_v121 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(-508)).Cmp(_609_v121) <= 0) && ((_609_v121).Cmp(_dafny.IntOfInt64(-87)) < 0) {
						_coll20.Add(Companion_Default___.SafeModuloInt(_609_v121, _dafny.IntOfInt64(392)), _480_v51.F27)
					}
				}
				return _coll20.ToMap()
			}()
		})(), 8)
		(_nw95).ArraySet1((_604_v119).Merge((func() _dafny.Map {
			if (_607_v123).Contains(false) {
				return (_607_v123).Get(false).(_dafny.Map)
			}
			return _604_v119
		})()), 9)
		(_nw95).ArraySet1(_604_v119, 10)
		_608_v124 = _nw95
		(globalState).F8 = _608_v124
		r0 = ((_472_v44).IsProperSubsetOf(func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(72), _dafny.IntOfInt64(341))); ; {
				_compr_21, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _610_v125 _dafny.Int
				_610_v125 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(72)).Cmp(_610_v125) <= 0) && ((_610_v125).Cmp(_dafny.IntOfInt64(341)) < 0) {
					_coll21.Add((_610_v125).Times((_this).F20()))
				}
			}
			return _coll21.ToSet()
		}())) && ((_470_v42).Equals(_470_v42))
		r1 = (_472_v44).IsProperSubsetOf((_472_v44).Intersection(func() _dafny.Set {
			var _coll22 = _dafny.NewBuilder()
			_ = _coll22
			for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-389), _dafny.IntOfInt64(-64))); ; {
				_compr_22, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _611_v126 _dafny.Int
				_611_v126 = interface{}(_compr_22).(_dafny.Int)
				if ((_dafny.IntOfInt64(-389)).Cmp(_611_v126) <= 0) && ((_611_v126).Cmp(_dafny.IntOfInt64(-64)) < 0) {
					_coll22.Add((_611_v126).Minus(_476_v47.F27))
				}
			}
			return _coll22.ToSet()
		}()))
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f20 _dafny.Int
	_f21 bool
	_f25 _dafny.Int
	_f26 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
	_this._f25 = _dafny.Zero
	_this._f26 = false
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

func (_this *C3) F20() _dafny.Int {
	return _this._f20
}
func (_this *C3) F21() bool {
	return _this._f21
}
func (_this *C3) Ctor__(f25 _dafny.Int, f26 bool, f20 _dafny.Int, f21 bool) {
	{
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C3) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F26()
	}
}
func (_this *C3) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F25()
	}
}
func (_this *C3) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _612_v0 D3
		_ = _612_v0
		_612_v0 = Companion_D3_.Create_DC6_((_this).F25())
		var _pat_let_tv18 = p1
		_ = _pat_let_tv18
		var _pat_let_tv19 = p1
		_ = _pat_let_tv19
		var _pat_let_tv20 = p1
		_ = _pat_let_tv20
		var _pat_let_tv21 = p1
		_ = _pat_let_tv21
		var _source9 D4 = func(_source10 D3) D4 {
			if _source10.Is_DC6() {
				var _613___mcc_h6 _dafny.Int = _source10.Get_().(D3_DC6).Cf10
				_ = _613___mcc_h6
				var _614_cf10 _dafny.Int = _613___mcc_h6
				_ = _614_cf10
				return Companion_D4_.Create_DC10_(_dafny.MultiSetOf((false)), (_this).F21())
			} else if _source10.Is_DC7() {
				var _615___mcc_h7 _dafny.Sequence = _source10.Get_().(D3_DC7).Cf11
				_ = _615___mcc_h7
				var _616___mcc_h8 bool = _source10.Get_().(D3_DC7).Cf12
				_ = _616___mcc_h8
				var _617_cf12 bool = _616___mcc_h8
				_ = _617_cf12
				var _618_cf11 _dafny.Sequence = _615___mcc_h7
				_ = _618_cf11
				return Companion_D4_.Create_DC10_(_dafny.MultiSetOf(_617_cf12, (_this).F26(), _pat_let_tv18), (_this).F21())
			} else if _source10.Is_DC5() {
				var _619___mcc_h9 _dafny.Sequence = _source10.Get_().(D3_DC5).Cf9
				_ = _619___mcc_h9
				var _620_cf9 _dafny.Sequence = _619___mcc_h9
				_ = _620_cf9
				return Companion_D4_.Create_DC10_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, (_this).F26(), _pat_let_tv19, false)), _pat_let_tv20)
			} else {
				var _621___mcc_h10 D3 = _source10.Get_().(D3_DC8).Cf13
				_ = _621___mcc_h10
				var _622_cf13 D3 = _621___mcc_h10
				_ = _622_cf13
				return Companion_D4_.Create_DC10_(_dafny.MultiSetOf(!(_pat_let_tv21)), (_this).F26())
			}
		}(_612_v0)
		_ = _source9
		if _source9.Is_DC10() {
			var _623___mcc_h0 _dafny.MultiSet = _source9.Get_().(D4_DC10).Cf15
			_ = _623___mcc_h0
			var _624___mcc_h1 bool = _source9.Get_().(D4_DC10).Cf16
			_ = _624___mcc_h1
			var _625_cf16 bool = _624___mcc_h1
			_ = _625_cf16
			var _626_cf15 _dafny.MultiSet = _623___mcc_h0
			_ = _626_cf15
			var _627_v1 _dafny.Sequence
			_ = _627_v1
			_627_v1 = _dafny.UnicodeSeqOfUtf8Bytes("dknejhjps")
			var _628_v2 _dafny.Array
			_ = _628_v2
			var _nwElement0_23 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F25(), (_this).F20())
			_ = _nwElement0_23
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(16))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_23, 0)
			(_nw96).ArraySet1((_this).F25(), 1)
			(_nw96).ArraySet1((_this).Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nbras")).Cardinality()), (_this).F20(), (_this).F26(), globalState), 2)
			(_nw96).ArraySet1(p0, 3)
			(_nw96).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F25(), (_this).F20()), 4)
			(_nw96).ArraySet1(p0, 5)
			(_nw96).ArraySet1(_dafny.IntOfInt64(205), 6)
			(_nw96).ArraySet1((func() _dafny.Int {
				if (_this).F21() {
					return (_this).F20()
				}
				return (_this).F25()
			})(), 7)
			(_nw96).ArraySet1(((_this).F20()).Minus(_dafny.IntOfUint32((_627_v1).Cardinality())), 8)
			(_nw96).ArraySet1((_this).F25(), 9)
			(_nw96).ArraySet1((_this).F25(), 10)
			(_nw96).ArraySet1(Companion_Default___.SafeModuloInt((_this).F20(), p0), 11)
			(_nw96).ArraySet1((_this).Fm1(p0, _dafny.IntOfInt64(978), (_this).F21(), globalState), 12)
			(_nw96).ArraySet1((_this).F25(), 13)
			(_nw96).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20())), 14)
			(_nw96).ArraySet1((_this).F20(), 15)
			_628_v2 = _nw96
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))
			_ = _index101
			(_628_v2).ArraySet1((_this).Fm1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(759))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_629_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()), (_this).F20(), p1, globalState), (_index101).Int())
			var _630_v3 _dafny.Array
			_ = _630_v3
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_16
			var _nw97 _dafny.Array
			_ = _nw97
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw97 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) D4 = (func(_631_p0 _dafny.Int) func(_dafny.Int) D4 {
					return func(_632_i1 _dafny.Int) D4 {
						return Companion_D4_.Create_DC11_(_631_p0, (_this).F21(), false)
					}
				})(p0)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw97 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw97).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw97).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_630_v3 = _nw97
			var _633_v4 _dafny.Array
			_ = _633_v4
			_633_v4 = _628_v2
			var _634_v6 _dafny.Map
			_ = _634_v6
			_634_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F21()), (_this).F21())
			var _635_v7 _dafny.Sequence
			_ = _635_v7
			_635_v7 = _dafny.SeqOf(_634_v6)
			var _636_v8 _dafny.Sequence
			_ = _636_v8
			_636_v8 = _dafny.SeqOf((_this).F20(), p0)
			var _637_v9 _dafny.Array
			_ = _637_v9
			var _nwElement0_24 _dafny.Array = (_633_v4)
			_ = _nwElement0_24
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(25))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_24, 0)
			(_nw98).ArraySet1(_628_v2, 1)
			(_nw98).ArraySet1(_628_v2, 2)
			(_nw98).ArraySet1((func() _dafny.Array {
				if (_this).F21() {
					return _628_v2
				}
				return _628_v2
			})(), 3)
			(_nw98).ArraySet1(_628_v2, 4)
			(_nw98).ArraySet1(_628_v2, 5)
			(_nw98).ArraySet1(_628_v2, 6)
			(_nw98).ArraySet1((func() _dafny.Array {
				if (_this).Fm0(func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(82), _dafny.IntOfInt64(-571))); ; {
						_compr_23, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _638_v5 _dafny.Int
						_638_v5 = interface{}(_compr_23).(_dafny.Int)
						if ((_dafny.IntOfInt64(82)).Cmp(_638_v5) <= 0) && ((_638_v5).Cmp(_dafny.IntOfInt64(-571)) < 0) {
							_coll23.Add(Companion_Default___.SafeModuloInt(_638_v5, _dafny.IntOfInt64(939)), (_this).F20())
						}
					}
					return _coll23.ToMap()
				}(), _627_v1, _dafny.Companion_Sequence_.Update(_635_v7, (Companion_Default___.SafeIndex((_636_v8).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_636_v8).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_635_v7).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _625_cf16)), globalState) {
					return _628_v2
				}
				return _628_v2
			})(), 7)
			(_nw98).ArraySet1(_628_v2, 8)
			(_nw98).ArraySet1(_628_v2, 9)
			(_nw98).ArraySet1(_628_v2, 10)
			(_nw98).ArraySet1(_628_v2, 11)
			(_nw98).ArraySet1(_628_v2, 12)
			(_nw98).ArraySet1(_628_v2, 13)
			(_nw98).ArraySet1(_628_v2, 14)
			(_nw98).ArraySet1(_628_v2, 15)
			(_nw98).ArraySet1(_628_v2, 16)
			(_nw98).ArraySet1(_628_v2, 17)
			(_nw98).ArraySet1(_628_v2, 18)
			(_nw98).ArraySet1(_628_v2, 19)
			(_nw98).ArraySet1(_628_v2, 20)
			(_nw98).ArraySet1(_628_v2, 21)
			(_nw98).ArraySet1(_628_v2, 22)
			(_nw98).ArraySet1(_628_v2, 23)
			(_nw98).ArraySet1(_628_v2, 24)
			_637_v9 = _nw98
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_637_v9), 0))
			_ = _index102
			(_637_v9).ArraySet1(_628_v2, (_index102).Int())
			var _639_v10 _dafny.Array
			_ = _639_v10
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_17
			var _nw99 _dafny.Array
			_ = _nw99
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw99 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Map = (func(_640_v6 _dafny.Map, _641_p1 bool) func(_dafny.Int) _dafny.Map {
					return func(_642_i2 _dafny.Int) _dafny.Map {
						return (_640_v6).Update((_this).F21(), _641_p1)
					}
				})(_634_v6, p1)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw99 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw99).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw99).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_639_v10 = _nw99
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_639_v10), 0))
			_ = _index103
			(_639_v10).ArraySet1(_634_v6, (_index103).Int())
			var _643_v11 _dafny.CodePoint
			_ = _643_v11
			_643_v11 = _dafny.CodePoint('b')
			var _644_v12 _dafny.MultiSet
			_ = _644_v12
			_644_v12 = _dafny.MultiSetOf(_643_v11)
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))
			_ = _index104
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_637_v9), 0))
			_ = _index105
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_639_v10), 0))
			_ = _index106
			var _rhs86 _dafny.Int = ((_dafny.Zero).Minus((p0).Minus((_644_v12).Cardinality()))).Minus(_dafny.IntOfInt64(-180))
			_ = _rhs86
			var _rhs87 _dafny.Array = _630_v3
			_ = _rhs87
			var _rhs88 _dafny.Array = _628_v2
			_ = _rhs88
			var _rhs89 bool = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_645_i3 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
			}))).Cardinality()))).Cmp((_this).F20()) != 0
			_ = _rhs89
			var _rhs90 _dafny.Map = (func() _dafny.Map {
				if (_this).F21() {
					return (_634_v6).Update((_this).F26(), (_this).F21())
				}
				return Companion_Default___.Fm15((_this).F20(), (_this).F25(), (_this).F25(), _dafny.IntOfInt64(-399), globalState)
			})()
			_ = _rhs90
			var _lhs78 _dafny.Array = _628_v2
			_ = _lhs78
			var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))
			_ = _lhs79
			var _lhs80 _dafny.Array = _637_v9
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_637_v9), 0))
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			var _lhs83 _dafny.Array = _639_v10
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_639_v10), 0))
			_ = _lhs84
			(_lhs78).ArraySet1(_rhs86, (_lhs79).Int())
			_630_v3 = _rhs87
			(_lhs80).ArraySet1(_rhs88, (_lhs81).Int())
			_lhs82.F3 = _rhs89
			(_lhs83).ArraySet1(_rhs90, (_lhs84).Int())
			if ((_this).F20()).Cmp((_this).Fm1(_dafny.IntOfUint32((_636_v8).Cardinality()), (_628_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))).Int()).(_dafny.Int), false, globalState)) == 0 {
				(globalState).F15 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(348))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_646_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_647_i4 _dafny.Int) _dafny.CodePoint {
						return _646_v11
					}
				})(_643_v11)))
				var _648_v13 _dafny.MultiSet
				_ = _648_v13
				_648_v13 = _dafny.MultiSetOf((_dafny.Zero).Minus(((_this).F20()).Plus((_this).F25())))
				_648_v13 = _648_v13
				var _649_v14 _dafny.Array
				_ = _649_v14
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw100
				_649_v14 = _nw100
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_649_v14), 0))
				_ = _index107
				(_649_v14).ArraySet1((_this).F21(), (_index107).Int())
				var _650_v15 D4
				_ = _650_v15
				_650_v15 = Companion_D4_.Create_DC10_(_626_cf15, p1)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_649_v14), 0))
				_ = _index108
				(_649_v14).ArraySet1((_650_v15).Dtor_cf16(), (_index108).Int())
				var _651_v16 _dafny.Map
				_ = _651_v16
				_651_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_628_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))).Int()).(_dafny.Int), (_this).F25())
				var _652_v17 _dafny.Map
				_ = _652_v17
				_652_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_625_cf16, _dafny.IntOfUint32((_627_v1).Cardinality()))
				var _653_v18 _dafny.Map
				_ = _653_v18
				_653_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_651_v16, (_652_v17).Cardinality())
				(globalState).F17 = (func() _dafny.Int {
					if (_653_v18).Contains(_651_v16) {
						return (_653_v18).Get(_651_v16).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-292)
				})()
				var _654_v19 *C0
				_ = _654_v19
				var _nw101 *C0 = New_C0_()
				_ = _nw101
				_nw101.Ctor__()
				_654_v19 = _nw101
			} else {
				var _655_v20 _dafny.Set
				_ = _655_v20
				_655_v20 = _dafny.SetOf(_626_cf15)
				var _656_v21 _dafny.Sequence
				_ = _656_v21
				_656_v21 = _dafny.SeqOf(_dafny.SetOf(_626_cf15, _626_cf15, _626_cf15), _dafny.SetOf(_626_cf15, _dafny.MultiSetOf(false, p1), _626_cf15))
				_655_v20 = (_656_v21).Select((Companion_Default___.SafeIndex(((_this).F20()).Plus((_this).F25()), _dafny.IntOfUint32((_656_v21).Cardinality()))).Uint32()).(_dafny.Set)
				var _657_v22 _dafny.Array
				_ = _657_v22
				var _nwElement0_25 _dafny.Array = _633_v4
				_ = _nwElement0_25
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(6))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_25, 0)
				(_nw102).ArraySet1(_633_v4, 1)
				(_nw102).ArraySet1(_633_v4, 2)
				(_nw102).ArraySet1(_633_v4, 3)
				(_nw102).ArraySet1(_633_v4, 4)
				(_nw102).ArraySet1(_633_v4, 5)
				_657_v22 = _nw102
				var _658_v23 _dafny.Map
				_ = _658_v23
				_658_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v22, p1)
				_658_v23 = (_658_v23).Update(_657_v22, (_this).F26())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))
				_ = _index109
				(_628_v2).ArraySet1(p0, (_index109).Int())
				var _659_v24 *C0
				_ = _659_v24
				var _nw103 *C0 = New_C0_()
				_ = _nw103
				_nw103.Ctor__()
				_659_v24 = _nw103
				var _660_v25 _dafny.Map
				_ = _660_v25
				_660_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_659_v24, (_dafny.Zero).Minus(((_this).F25()).Times((_628_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))).Int()).(_dafny.Int))))
				var _661_v26 _dafny.Sequence
				_ = _661_v26
				_661_v26 = _dafny.SeqOf(_659_v24)
				_660_v25 = (_660_v25).Update((_661_v26).Select((Companion_Default___.SafeIndex((_628_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_661_v26).Cardinality()))).Uint32()).(*C0), (_this).F25())
				var _662_v27 _dafny.Sequence
				_ = _662_v27
				_662_v27 = _dafny.SeqOf((_633_v4), _dafny.ArrayCastTo((_637_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_637_v9), 0))).Int())), _dafny.ArrayCastTo((_637_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_637_v9), 0))).Int())))
				var _663_v28 T0
				_ = _663_v28
				var _nw104 *C1 = New_C1_()
				_ = _nw104
				_nw104.Ctor__(((_636_v8).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_636_v8).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p0), (_662_v27).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm29((_this).F25(), (_this).F26(), globalState)).Cardinality(), _dafny.IntOfUint32((_662_v27).Cardinality()))).Uint32()).(_dafny.Array), _dafny.IntOfUint32((_627_v1).Cardinality()), (_this).F21())
				_663_v28 = _nw104
				var _664_v29 _dafny.Map
				_ = _664_v29
				_664_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_639_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_639_v10), 0))).Int()).(_dafny.Map))).Cardinality()), (_663_v28).F21())
				var _665_v30 D9
				_ = _665_v30
				_665_v30 = Companion_D9_.Create_DC23_((_628_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_628_v2), 0))).Int()).(_dafny.Int), (_this).F26())
				var _666_v31 _dafny.Array
				_ = _666_v31
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_18
				var _nw105 _dafny.Array
				_ = _nw105
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw105 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) D6 = (func(_667_v1 _dafny.Sequence) func(_dafny.Int) D6 {
						return func(_668_i5 _dafny.Int) D6 {
							return Companion_D6_.Create_DC13_(_667_v1)
						}
					})(_627_v1)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw105 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw105).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw105).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_666_v31 = _nw105
				var _669_v32 _dafny.Map
				_ = _669_v32
				_669_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _666_v31)
				var _670_v33 _dafny.Array
				_ = _670_v33
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw106
				_670_v33 = _nw106
				var _671_v34 _dafny.Sequence
				_ = _671_v34
				_671_v34 = _dafny.SeqOf(_670_v33, _670_v33)
				var _672_v35 D3
				_ = _672_v35
				_672_v35 = Companion_D3_.Create_DC5_(_671_v34)
				var _673_v36 _dafny.MultiSet
				_ = _673_v36
				_673_v36 = _dafny.MultiSetOf(_672_v35)
				var _674_v37 _dafny.Map
				_ = _674_v37
				_674_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (func() _dafny.Int {
					if (_673_v36).Contains(_672_v35) {
						return (_673_v36).Multiplicity(_672_v35)
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}((func(_675_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_676_i6 _dafny.Int) _dafny.CodePoint {
							return _675_v11
						}
					})(_643_v11)))).Cardinality())
				})())
				var _677_v38 _dafny.Sequence
				_ = _677_v38
				_677_v38 = _dafny.SeqOf(_674_v37)
				var _678_v39 *C1
				_ = _678_v39
				var _nw107 *C1 = New_C1_()
				_ = _nw107
				_nw107.Ctor__(p0, _628_v2, (_663_v28).F20(), _625_cf16)
				_678_v39 = _nw107
				var _679_v40 _dafny.Sequence
				_ = _679_v40
				_679_v40 = _dafny.SeqOf(_678_v39)
				var _680_v41 _dafny.Array
				_ = _680_v41
				var _nwElement0_26 bool = p1
				_ = _nwElement0_26
				var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(19))
				_ = _nw108
				(_nw108).ArraySet1(_nwElement0_26, 0)
				(_nw108).ArraySet1((_663_v28).F21(), 1)
				(_nw108).ArraySet1((func() bool {
					if (_664_v29).Contains((_663_v28).F20()) {
						return (_664_v29).Get((_663_v28).F20()).(bool)
					}
					return true
				})(), 2)
				(_nw108).ArraySet1((_665_v30).Dtor_cf32(), 3)
				(_nw108).ArraySet1((_this).F21(), 4)
				(_nw108).ArraySet1((_this).F26(), 5)
				(_nw108).ArraySet1((_663_v28).F21(), 6)
				(_nw108).ArraySet1(_625_cf16, 7)
				(_nw108).ArraySet1(false, 8)
				(_nw108).ArraySet1(((_663_v28).F21()) == ((_this).F21()), 9)
				(_nw108).ArraySet1((func() bool {
					if _625_cf16 {
						return (_this).F21()
					}
					return _625_cf16
				})(), 10)
				(_nw108).ArraySet1(!((func() bool {
					if _625_cf16 {
						return (_this).F21()
					}
					return _625_cf16
				})()), 11)
				(_nw108).ArraySet1((_659_v24).Fm4((_669_v32).Cardinality(), _677_v38, (_this).F20(), globalState), 12)
				(_nw108).ArraySet1((_this).F21(), 13)
				(_nw108).ArraySet1(_dafny.Companion_Sequence_.Contains(_679_v40, _678_v39), 14)
				(_nw108).ArraySet1((_678_v39).Fm0(_674_v37, _dafny.UnicodeSeqOfUtf8Bytes("safjb"), _635_v7, globalState), 15)
				(_nw108).ArraySet1(p1, 16)
				(_nw108).ArraySet1((_this).F21(), 17)
				(_nw108).ArraySet1((_this).F26(), 18)
				_680_v41 = _nw108
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_670_v33), 0))
				_ = _index110
				(_670_v33).ArraySet1(p1, (_index110).Int())
				var _681_v42 _dafny.MultiSet
				_ = _681_v42
				_681_v42 = _dafny.MultiSetOf(_dafny.IntOfInt64(286), (_this).F25())
				var _682_v44 _dafny.Set
				_ = _682_v44
				_682_v44 = _dafny.SetOf((_this).F26(), (_this).F21(), (_this).F26())
				var _683_v45 _dafny.Sequence
				_ = _683_v45
				_683_v45 = _dafny.SeqOf((_682_v44).Cardinality())
				var _684_v46 _dafny.Map
				_ = _684_v46
				_684_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _dafny.IntOfInt64(-665))
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_670_v33), 0))
				_ = _index111
				var _rhs91 _dafny.Int = (func() _dafny.Int {
					if (_663_v28).F21() {
						return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-584), (func() _dafny.Set {
							var _coll24 = _dafny.NewBuilder()
							_ = _coll24
							for _iter25 := _dafny.Iterate((_681_v42).Elements()); ; {
								_compr_24, _ok25 := _iter25()
								if !_ok25 {
									break
								}
								var _685_v43 _dafny.Int
								_685_v43 = interface{}(_compr_24).(_dafny.Int)
								if (_681_v42).Contains(_685_v43) {
									_coll24.Add(Companion_Default___.SafeDivisionInt(_685_v43, _dafny.IntOfInt64(145)))
								}
							}
							return _coll24.ToSet()
						}()).Cardinality())
					}
					return (_this).F25()
				})()
				_ = _rhs91
				var _rhs92 T0 = _663_v28
				_ = _rhs92
				var _rhs93 bool = !((_663_v28).F21())
				_ = _rhs93
				var _rhs94 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_686_i7 _dafny.Int) _dafny.Int {
					return (_this).F20()
				}))
				_ = _rhs94
				var _rhs95 bool = !(!((_678_v39).Fm16(_626_cf15, _683_v45, globalState)) || (p1)) || (!(_684_v46).Contains((_this).F26()))
				_ = _rhs95
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _670_v33
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_670_v33), 0))
				_ = _lhs88
				_lhs85.F17 = _rhs91
				_663_v28 = _rhs92
				_lhs86.F3 = _rhs93
				_636_v8 = _rhs94
				(_lhs87).ArraySet1(_rhs95, (_lhs88).Int())
			}
			var _687_v47 _dafny.MultiSet
			_ = _687_v47
			_687_v47 = _dafny.MultiSetOf(_dafny.IntOfInt64(512))
			var _688_v48 _dafny.Map
			_ = _688_v48
			_688_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v47, _dafny.UnicodeSeqOfUtf8Bytes("auwfsukpw"))
			_688_v48 = (_688_v48).Update(_687_v47, _dafny.Companion_Sequence_.Concatenate(_627_v1, _627_v1))
			var _689_v49 _dafny.Sequence
			_ = _689_v49
			_689_v49 = _dafny.SeqOf(_687_v47)
			_689_v49 = _689_v49
		} else if _source9.Is_DC11() {
			var _690___mcc_h2 _dafny.Int = _source9.Get_().(D4_DC11).Cf17
			_ = _690___mcc_h2
			var _691___mcc_h3 bool = _source9.Get_().(D4_DC11).Cf18
			_ = _691___mcc_h3
			var _692___mcc_h4 bool = _source9.Get_().(D4_DC11).Cf19
			_ = _692___mcc_h4
			var _693_cf19 bool = _692___mcc_h4
			_ = _693_cf19
			var _694_cf18 bool = _691___mcc_h3
			_ = _694_cf18
			var _695_cf17 _dafny.Int = _690___mcc_h2
			_ = _695_cf17
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(372))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_696_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(332))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_697_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))) {
				var _698_v50 _dafny.Sequence
				_ = _698_v50
				_698_v50 = _dafny.UnicodeSeqOfUtf8Bytes("iagyloo")
				var _699_v51 _dafny.Array
				_ = _699_v51
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_19
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Int = func(_700_i10 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_700_i10, (_dafny.Zero).Minus((_this).F25()))
					}
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw109 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw109).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw109).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_699_v51 = _nw109
				var _701_v52 *C1
				_ = _701_v52
				var _nw110 *C1 = New_C1_()
				_ = _nw110
				_nw110.Ctor__((_dafny.IntOfUint32((_698_v50).Cardinality())).Plus((_this).F20()), _699_v51, p0, _694_cf18)
				_701_v52 = _nw110
				var _702_v53 *C0
				_ = _702_v53
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__()
				_702_v53 = _nw111
				var _703_v54 _dafny.Array
				_ = _703_v54
				var _nwElement0_27 bool = _694_cf18
				_ = _nwElement0_27
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_27, 0)
				(_nw112).ArraySet1((_this).F21(), 1)
				(_nw112).ArraySet1((_this).F21(), 2)
				(_nw112).ArraySet1((_this).F26(), 3)
				(_nw112).ArraySet1(_693_cf19, 4)
				(_nw112).ArraySet1(false, 5)
				(_nw112).ArraySet1((_this).F21(), 6)
				(_nw112).ArraySet1((_this).F26(), 7)
				(_nw112).ArraySet1(_693_cf19, 8)
				(_nw112).ArraySet1((_this).F26(), 9)
				(_nw112).ArraySet1(_694_cf18, 10)
				(_nw112).ArraySet1(!(_694_cf18), 11)
				(_nw112).ArraySet1(_694_cf18, 12)
				(_nw112).ArraySet1(_693_cf19, 13)
				_703_v54 = _nw112
				var _704_v55 _dafny.Set
				_ = _704_v55
				_704_v55 = _dafny.SetOf(_703_v54, _703_v54, _703_v54)
				r1 = (_704_v55).Contains(_703_v54)
				var _705_v56 _dafny.CodePoint
				_ = _705_v56
				_705_v56 = _dafny.CodePoint('j')
				var _706_v57 _dafny.MultiSet
				_ = _706_v57
				_706_v57 = _dafny.MultiSetOf(_705_v56, _705_v56)
				var _707_v58 _dafny.MultiSet
				_ = _707_v58
				_707_v58 = _dafny.MultiSetOf(_706_v57)
				var _708_v59 _dafny.Sequence
				_ = _708_v59
				_708_v59 = _dafny.SeqOf(_707_v58, _dafny.MultiSetOf(_706_v57, _dafny.MultiSetOf(_705_v56)))
				var _709_v60 _dafny.Sequence
				_ = _709_v60
				_709_v60 = _dafny.SeqOf(_707_v58, _707_v58, _dafny.MultiSetOf(_706_v57), (_708_v59).Select((Companion_Default___.SafeIndex(_695_cf17, _dafny.IntOfUint32((_708_v59).Cardinality()))).Uint32()).(_dafny.MultiSet), Companion_Default___.Fm31(globalState))
				_693_cf19 = !((_dafny.MultiSetFromSeq(Companion_Default___.Fm30(globalState))).Equals((_708_v59).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_708_v59).Cardinality()))).Uint32()).(_dafny.MultiSet)))
				(globalState).F4 = (_this).F21()
			} else {
				var _710_v61 _dafny.MultiSet
				_ = _710_v61
				_710_v61 = _dafny.MultiSetOf(p0, _695_cf17)
				var _711_v62 _dafny.Map
				_ = _711_v62
				_711_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).Fm1(p0, (_710_v61).Cardinality(), _693_cf19, globalState))
				var _712_v63 _dafny.Map
				_ = _712_v63
				_712_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if !(true) {
						return (_this).F26()
					}
					return (_this).F21()
				})(), (_this).Fm1((func() _dafny.Int {
					if (_711_v62).Contains((_this).F26()) {
						return (_711_v62).Get((_this).F26()).(_dafny.Int)
					}
					return p0
				})(), p0, (_this).F21(), globalState))
				_712_v63 = (_712_v63).Update((_this).F26(), _dafny.IntOfInt64(-402))
				var _713_v64 _dafny.Sequence
				_ = _713_v64
				_713_v64 = _dafny.UnicodeSeqOfUtf8Bytes("qvfhbcfbh")
				var _714_v65 _dafny.Sequence
				_ = _714_v65
				_714_v65 = _dafny.SeqOf(Companion_Default___.Fm26(_694_cf18, globalState), _713_v64)
				_714_v65 = _714_v65
				var _715_v66 _dafny.Map
				_ = _715_v66
				_715_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
				var _716_v67 _dafny.Sequence
				_ = _716_v67
				_716_v67 = _dafny.SeqOf((func() bool {
					if (_715_v66).Contains(_694_cf18) {
						return (_715_v66).Get(_694_cf18).(bool)
					}
					return _693_cf19
				})())
				var _717_v68 _dafny.MultiSet
				_ = _717_v68
				_717_v68 = _dafny.MultiSetOf(_716_v67, _716_v67)
				var _718_v69 _dafny.Sequence
				_ = _718_v69
				_718_v69 = _dafny.SeqOf(_dafny.IntOfInt64(594), (_717_v68).Cardinality(), (_this).Fm1(_695_cf17, (_715_v66).Cardinality(), _694_cf18, globalState))
				var _719_v70 _dafny.Set
				_ = _719_v70
				_719_v70 = _dafny.SetOf((_dafny.Zero).Minus(_695_cf17))
				var _720_v71 _dafny.Sequence
				_ = _720_v71
				_720_v71 = _dafny.SeqOf(_719_v70, _719_v70)
				var _721_v72 _dafny.Sequence
				_ = _721_v72
				_721_v72 = _dafny.SeqOf(_695_cf17, (_this).F25(), (Companion_D7_.Create_DC18_(_718_v69, _695_cf17)).Dtor_cf27(), (_this).F25(), (Companion_D8_.Create_DC21_(((_720_v71).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_720_v71).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())).Dtor_cf29())
				var _722_v73 _dafny.Array
				_ = _722_v73
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw113
				_722_v73 = _nw113
				var _723_v74 _dafny.Sequence
				_ = _723_v74
				_723_v74 = _dafny.SeqOf(_722_v73)
				var _rhs96 bool = _694_cf18
				_ = _rhs96
				var _rhs97 _dafny.Int = (func() _dafny.Int {
					if !(true) {
						return (_dafny.Zero).Minus(_695_cf17)
					}
					return (_dafny.Zero).Minus(p0)
				})()
				_ = _rhs97
				var _rhs98 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_721_v72, _721_v72)
				_ = _rhs98
				var _rhs99 _dafny.Sequence = _723_v74
				_ = _rhs99
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_lhs89.F18 = _rhs96
				_lhs90.F5 = _rhs97
				_lhs91.F14 = _rhs98
				r0 = _rhs99
				var _724_v75 _dafny.Map
				_ = _724_v75
				_724_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_718_v69).Cardinality()))
				var _725_v76 _dafny.Map
				_ = _725_v76
				_725_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_724_v75, (_this).F20())
				_725_v76 = (_725_v76).Update(_724_v75, p0)
				var _726_v77 _dafny.CodePoint
				_ = _726_v77
				_726_v77 = _dafny.CodePoint('v')
				var _727_v78 *C1
				_ = _727_v78
				var _nw114 *C1 = New_C1_()
				_ = _nw114
				_nw114.Ctor__(_695_cf17, _722_v73, (_this).F25(), _694_cf18)
				_727_v78 = _nw114
				var _728_v79 _dafny.Map
				_ = _728_v79
				_728_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v77, _727_v78)
				var _729_v80 _dafny.Sequence
				_ = _729_v80
				_729_v80 = _dafny.SeqOf(_728_v79, _728_v79)
				var _730_v81 D10
				_ = _730_v81
				_730_v81 = Companion_D10_.Create_DC27_(_695_cf17, _729_v80, _726_v77)
				var _731_v82 _dafny.MultiSet
				_ = _731_v82
				_731_v82 = _dafny.MultiSetOf((_730_v81).Dtor_cf38())
				var _732_v83 _dafny.Sequence
				_ = _732_v83
				_732_v83 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_713_v64), _731_v82)
				var _733_v84 _dafny.Array
				_ = _733_v84
				var _nwElement0_28 _dafny.MultiSet = (_731_v82).Union(Companion_Default___.Fm32(globalState))
				_ = _nwElement0_28
				var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(17))
				_ = _nw115
				(_nw115).ArraySet1(_nwElement0_28, 0)
				(_nw115).ArraySet1(_731_v82, 1)
				(_nw115).ArraySet1(_731_v82, 2)
				(_nw115).ArraySet1((_732_v83).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_732_v83).Cardinality()))).Uint32()).(_dafny.MultiSet), 3)
				(_nw115).ArraySet1(_731_v82, 4)
				(_nw115).ArraySet1(_731_v82, 5)
				(_nw115).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('r'), _726_v77)), 6)
				(_nw115).ArraySet1((func() _dafny.MultiSet {
					if _693_cf19 {
						return _dafny.MultiSetOf(_726_v77, _dafny.CodePoint('v'))
					}
					return _731_v82
				})(), 7)
				(_nw115).ArraySet1((_731_v82).Union((_732_v83).Select((Companion_Default___.SafeIndex(_727_v78.F27, _dafny.IntOfUint32((_732_v83).Cardinality()))).Uint32()).(_dafny.MultiSet)), 8)
				(_nw115).ArraySet1(_dafny.MultiSetOf(_726_v77, _726_v77, _726_v77, _726_v77, _726_v77), 9)
				(_nw115).ArraySet1(_731_v82, 10)
				(_nw115).ArraySet1(_731_v82, 11)
				(_nw115).ArraySet1(_dafny.MultiSetOf(_726_v77, _726_v77), 12)
				(_nw115).ArraySet1(_dafny.MultiSetOf(_726_v77), 13)
				(_nw115).ArraySet1(_731_v82, 14)
				(_nw115).ArraySet1((_731_v82).Difference(_731_v82), 15)
				(_nw115).ArraySet1(_731_v82, 16)
				_733_v84 = _nw115
				_733_v84 = _733_v84
			}
			(globalState).F17 = (_this).F25()
			var _734_v85 bool
			_ = _734_v85
			_734_v85 = _693_cf19
			if _734_v85 {
				var _735_v86 _dafny.Array
				_ = _735_v86
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw116
				_735_v86 = _nw116
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_735_v86), 0))
				_ = _index112
				(_735_v86).ArraySet1((_this).F25(), (_index112).Int())
				var _736_v87 D7
				_ = _736_v87
				_736_v87 = Companion_D7_.Create_DC17_(_694_cf18)
				var _737_v88 _dafny.Set
				_ = _737_v88
				_737_v88 = _dafny.SetOf(false)
				var _738_v89 _dafny.Sequence
				_ = _738_v89
				_738_v89 = _dafny.SeqOf(_737_v88, _737_v88)
				var _739_v90 _dafny.Map
				_ = _739_v90
				_739_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_736_v87, _738_v89)
				var _740_v91 D3
				_ = _740_v91
				_740_v91 = Companion_D3_.Create_DC7_((func() _dafny.Sequence {
					if (_739_v90).Contains(_736_v87) {
						return (_739_v90).Get(_736_v87).(_dafny.Sequence)
					}
					return _738_v89
				})(), _693_cf19)
				var _741_v92 _dafny.MultiSet
				_ = _741_v92
				_741_v92 = _dafny.MultiSetOf((_this).F20(), (_this).Fm1(_dafny.IntOfInt64(-252), (_dafny.Zero).Minus(_695_cf17), (_740_v91).Dtor_cf12(), globalState), (_this).F25())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_735_v86), 0))
				_ = _index113
				(_735_v86).ArraySet1((func() _dafny.Int {
					if (_741_v92).Contains((_this).F20()) {
						return (_741_v92).Multiplicity((_this).F20())
					}
					return (_this).F25()
				})(), (_index113).Int())
				var _742_v93 _dafny.Array
				_ = _742_v93
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_20
				var _nw117 _dafny.Array
				_ = _nw117
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw117 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = (func(_743_p1 bool) func(_dafny.Int) bool {
						return func(_744_i11 _dafny.Int) bool {
							return _743_p1
						}
					})(p1)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw117 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw117).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw117).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_742_v93 = _nw117
				var _nwElement0_29 bool = p1
				_ = _nwElement0_29
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(9))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_29, 0)
				(_nw118).ArraySet1(!(_694_cf18) || (p1), 1)
				(_nw118).ArraySet1(true, 2)
				(_nw118).ArraySet1(p1, 3)
				(_nw118).ArraySet1((_this).F21(), 4)
				(_nw118).ArraySet1(_694_cf18, 5)
				(_nw118).ArraySet1(p1, 6)
				(_nw118).ArraySet1(!((!(p1)) && (_694_cf18)), 7)
				(_nw118).ArraySet1(true, 8)
				_742_v93 = _nw118
				var _745_v94 _dafny.Array
				_ = _745_v94
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw119
				_745_v94 = _nw119
				var _746_v95 _dafny.Sequence
				_ = _746_v95
				_746_v95 = _dafny.SeqOf(_735_v86, _735_v86)
				var _747_v96 _dafny.Sequence
				_ = _747_v96
				_747_v96 = _dafny.UnicodeSeqOfUtf8Bytes("ysgax")
				var _rhs100 _dafny.Array = _745_v94
				_ = _rhs100
				var _rhs101 _dafny.Int = _dafny.IntOfInt64(785)
				_ = _rhs101
				var _rhs102 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_746_v95).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_746_v95).Cardinality()))).Uint32()).(_dafny.Array), _735_v86)).Cardinality())
				_ = _rhs102
				var _rhs103 bool = _693_cf19
				_ = _rhs103
				var _rhs104 _dafny.Int = Companion_Default___.SafeModuloInt(p0, (p0).Minus(_dafny.IntOfUint32((_747_v96).Cardinality())))
				_ = _rhs104
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				_745_v94 = _rhs100
				_lhs92.F17 = _rhs101
				_695_cf17 = _rhs102
				_lhs93.F14 = _rhs103
				_lhs94.F5 = _rhs104
				var _748_v97 _dafny.Sequence
				_ = _748_v97
				_748_v97 = _dafny.SeqOf((_this).F21(), _694_cf18)
				var _749_v98 _dafny.Map
				_ = _749_v98
				_749_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _748_v97)
				(globalState).F18 = ((_749_v98).Cardinality()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_747_v96, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-800), _dafny.IntOfUint32((_747_v96).Cardinality()))).Uint32(), _dafny.CodePoint('f'))).Cardinality()))) == 0
				(globalState).F4 = (_694_cf18) || (false)
			} else {
				var _750_v99 _dafny.Array
				_ = _750_v99
				var _nw120 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(25))
				_ = _nw120
				_750_v99 = _nw120
				var _751_v100 _dafny.Sequence
				_ = _751_v100
				_751_v100 = _dafny.SeqOf(_dafny.MultiSetOf((_this).F25(), (_this).F20(), _dafny.IntOfInt64(81), p0))
				var _752_v101 _dafny.Sequence
				_ = _752_v101
				_752_v101 = _dafny.SeqOf((_this).F26())
				var _rhs105 _dafny.Array = (func() _dafny.Array {
					if p1 {
						return _750_v99
					}
					return _750_v99
				})()
				_ = _rhs105
				var _rhs106 _dafny.Sequence = _751_v100
				_ = _rhs106
				var _rhs107 bool = (_752_v101).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvtiyhvfl")).Cardinality()), _dafny.IntOfUint32((_752_v101).Cardinality()))).Uint32()).(bool)
				_ = _rhs107
				var _rhs108 _dafny.Int = (_this).F20()
				_ = _rhs108
				var _rhs109 _dafny.Int = _695_cf17
				_ = _rhs109
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				_750_v99 = _rhs105
				_751_v100 = _rhs106
				_lhs95.F14 = _rhs107
				_695_cf17 = _rhs108
				_695_cf17 = _rhs109
				var _753_v102 _dafny.Array
				_ = _753_v102
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_21
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) bool = func(_754_i12 _dafny.Int) bool {
						return (_dafny.SetOf((_this).F21())).IsProperSubsetOf(_dafny.SetOf((_this).F26()))
					}
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw121 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw121).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw121).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_753_v102 = _nw121
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_753_v102), 0))
				_ = _index114
				(_753_v102).ArraySet1((_694_cf18) == (_693_cf19), (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_753_v102), 0))
				_ = _index115
				(_753_v102).ArraySet1((_this).F26(), (_index115).Int())
				var _755_v103 _dafny.Sequence
				_ = _755_v103
				_755_v103 = _dafny.UnicodeSeqOfUtf8Bytes("etpbgjoh")
				var _756_v104 D6
				_ = _756_v104
				_756_v104 = Companion_D6_.Create_DC13_(_755_v103)
				(globalState).F15 = (_756_v104).Dtor_cf21()
				var _757_v105 _dafny.Map
				_ = _757_v105
				_757_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _755_v103)
				var _758_v106 _dafny.Sequence
				_ = _758_v106
				_758_v106 = _dafny.SeqOf(_755_v103, _755_v103, _755_v103, _755_v103)
				_757_v105 = (_757_v105).Update(p0, (_758_v106).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_758_v106).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _759_v107 D9
				_ = _759_v107
				_759_v107 = Companion_D9_.Create_DC23_(_dafny.IntOfInt64(75), p1)
				var _760_v108 D9
				_ = _760_v108
				_760_v108 = Companion_D9_.Create_DC24_(Companion_D9_.Create_DC24_(_759_v107))
				var _761_v109 _dafny.Array
				_ = _761_v109
				var _nwElement0_30 D9 = Companion_D9_.Create_DC24_(_759_v107)
				_ = _nwElement0_30
				var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(11))
				_ = _nw122
				(_nw122).ArraySet1(_nwElement0_30, 0)
				(_nw122).ArraySet1(_760_v108, 1)
				(_nw122).ArraySet1(_760_v108, 2)
				(_nw122).ArraySet1(_760_v108, 3)
				(_nw122).ArraySet1(_760_v108, 4)
				(_nw122).ArraySet1(_760_v108, 5)
				(_nw122).ArraySet1(Companion_D9_.Create_DC24_(Companion_D9_.Create_DC23_(_695_cf17, _694_cf18)), 6)
				(_nw122).ArraySet1(Companion_D9_.Create_DC24_(Companion_D9_.Create_DC24_(Companion_D9_.Create_DC23_(_dafny.IntOfInt64(722), _694_cf18))), 7)
				(_nw122).ArraySet1(_760_v108, 8)
				(_nw122).ArraySet1(_760_v108, 9)
				(_nw122).ArraySet1(_760_v108, 10)
				_761_v109 = _nw122
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_761_v109), 0))
				_ = _index116
				(_761_v109).ArraySet1(_760_v108, (_index116).Int())
				var _762_v110 _dafny.Map
				_ = _762_v110
				_762_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_753_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_753_v102), 0))).Int()).(bool))
				var _763_v111 _dafny.Sequence
				_ = _763_v111
				_763_v111 = _dafny.SeqOf(_762_v110)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_761_v109), 0))
				_ = _index117
				(_761_v109).ArraySet1((func() D9 {
					if (_this).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmql")).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("rbvsweaf"), _763_v111, globalState) {
						return (func() D9 {
							if !(p1) {
								return _760_v108
							}
							return _760_v108
						})()
					}
					return _760_v108
				})(), (_index117).Int())
			}
			var _764_v112 _dafny.Array
			_ = _764_v112
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw123
			_764_v112 = _nw123
			var _765_v113 _dafny.Sequence
			_ = _765_v113
			_765_v113 = _dafny.SeqOf(p0, (_this).F20())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_764_v112), 0))
			_ = _index118
			(_764_v112).ArraySet1(_765_v113, (_index118).Int())
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_764_v112), 0))
			_ = _index119
			(_764_v112).ArraySet1(_765_v113, (_index119).Int())
		} else {
			var _766___mcc_h5 _dafny.Array = _source9.Get_().(D4_DC9).Cf14
			_ = _766___mcc_h5
			var _767_cf14 _dafny.Array = _766___mcc_h5
			_ = _767_cf14
			var _768_v114 _dafny.Sequence
			_ = _768_v114
			_768_v114 = _dafny.SeqOf((_this).F25(), p0)
			var _769_v115 _dafny.Set
			_ = _769_v115
			_769_v115 = _dafny.SetOf((func() _dafny.Sequence {
				if p1 {
					return _768_v114
				}
				return _768_v114
			})())
			_769_v115 = _769_v115
			var _770_v116 *C2
			_ = _770_v116
			var _nw124 *C2 = New_C2_()
			_ = _nw124
			_nw124.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngcig")).Cardinality()), (_this).F26())
			_770_v116 = _nw124
			var _771_v117 _dafny.CodePoint
			_ = _771_v117
			_771_v117 = _dafny.CodePoint('a')
			_771_v117 = _771_v117
			var _772_v118 _dafny.Set
			_ = _772_v118
			_772_v118 = _dafny.SetOf(p1)
			var _773_v119 _dafny.Sequence
			_ = _773_v119
			_773_v119 = _dafny.SeqOf(_dafny.SetOf(p1, (_this).F26()))
			var _774_v120 D3
			_ = _774_v120
			_774_v120 = Companion_D3_.Create_DC7_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_772_v118, _772_v118), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.SeqOf(_772_v118, _772_v118)).Cardinality()))).Uint32(), _772_v118), _773_v119), p1)
			var _source11 D3 = _774_v120
			_ = _source11
			if _source11.Is_DC6() {
				var _775___mcc_h11 _dafny.Int = _source11.Get_().(D3_DC6).Cf10
				_ = _775___mcc_h11
				var _776_cf10 _dafny.Int = _775___mcc_h11
				_ = _776_cf10
				var _777_v121 _dafny.Sequence
				_ = _777_v121
				_777_v121 = _dafny.SeqOf(p1)
				var _778_v122 _dafny.Array
				_ = _778_v122
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_22
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = func(_779_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("bty")
					}
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw125 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw125).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw125).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_778_v122 = _nw125
				var _780_v123 _dafny.Map
				_ = _780_v123
				_780_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p0)
				var _781_v124 _dafny.Map
				_ = _781_v124
				_781_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), true)
				var _782_v125 _dafny.Sequence
				_ = _782_v125
				_782_v125 = _dafny.SeqOf(_781_v124)
				var _783_v126 _dafny.Sequence
				_ = _783_v126
				_783_v126 = _dafny.SeqOf(_782_v125)
				var _rhs110 _dafny.Int = (_dafny.IntOfInt64(332)).Minus((func() _dafny.Int {
					if (_this).Fm0(_780_v123, _dafny.UnicodeSeqOfUtf8Bytes("ivwxw"), (_783_v126).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_783_v126).Cardinality()))).Uint32()).(_dafny.Sequence), globalState) {
						return p0
					}
					return _dafny.IntOfInt64(-46)
				})())
				_ = _rhs110
				var _rhs111 _dafny.Sequence = _777_v121
				_ = _rhs111
				var _rhs112 _dafny.Array = _778_v122
				_ = _rhs112
				var _rhs113 bool = (_this).F26()
				_ = _rhs113
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				_lhs96.F17 = _rhs110
				_777_v121 = _rhs111
				_778_v122 = _rhs112
				_lhs97.F4 = _rhs113
				var _784_v128 _dafny.Map
				_ = _784_v128
				_784_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _785_v129 _dafny.Map
				_ = _785_v129
				_785_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_768_v114).Cardinality()), (func() bool {
					if (_784_v128).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F20())).Cardinality()) {
						return (_784_v128).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F20())).Cardinality()).(bool)
					}
					return p1
				})())
				var _786_v131 _dafny.Sequence
				_ = _786_v131
				_786_v131 = _dafny.UnicodeSeqOfUtf8Bytes("kdacsl")
				var _787_v132 _dafny.Map
				_ = _787_v132
				_787_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
					var _coll25 = _dafny.NewBuilder()
					_ = _coll25
					for _iter26 := _dafny.Iterate((_785_v129).Keys().Elements()); ; {
						_compr_25, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _788_v130 _dafny.Int
						_788_v130 = interface{}(_compr_25).(_dafny.Int)
						if (_785_v129).Contains(_788_v130) {
							_coll25.Add(Companion_Default___.SafeModuloInt(_788_v130, _dafny.IntOfInt64(486)))
						}
					}
					return _coll25.ToSet()
				}(), _786_v131)
				var _789_v133 _dafny.Array
				_ = _789_v133
				var _nwElement0_31 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F20(), p0)
				_ = _nwElement0_31
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(14))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_31, 0)
				(_nw126).ArraySet1(((_781_v124).Merge(_781_v124)).Cardinality(), 1)
				(_nw126).ArraySet1((_772_v118).Cardinality(), 2)
				(_nw126).ArraySet1((_this).F20(), 3)
				(_nw126).ArraySet1((_this).F25(), 4)
				(_nw126).ArraySet1(_776_cf10, 5)
				(_nw126).ArraySet1((_this).F25(), 6)
				(_nw126).ArraySet1(p0, 7)
				(_nw126).ArraySet1((_780_v123).Cardinality(), 8)
				(_nw126).ArraySet1((_this).F20(), 9)
				(_nw126).ArraySet1(((func() _dafny.Map {
					var _coll26 = _dafny.NewMapBuilder()
					_ = _coll26
					for _iter27 := _dafny.Iterate((_787_v132).Keys().Elements()); ; {
						_compr_26, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _790_v127 _dafny.Set
						_790_v127 = interface{}(_compr_26).(_dafny.Set)
						if (_787_v132).Contains(_790_v127) {
							_coll26.Add(_790_v127, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()))
						}
					}
					return _coll26.ToMap()
				}()).Cardinality()).Plus((_this).F25()), 10)
				(_nw126).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pttigqy")).Cardinality()), 11)
				(_nw126).ArraySet1(p0, 12)
				(_nw126).ArraySet1(p0, 13)
				_789_v133 = _nw126
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))
				_ = _index120
				(_789_v133).ArraySet1(_776_cf10, (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))
				_ = _index121
				var _rhs114 bool = !((_772_v118).IsDisjointFrom(_772_v118))
				_ = _rhs114
				var _rhs115 _dafny.Int = (_770_v116).Fm1(_776_cf10, (_this).F20(), (_this).F21(), globalState)
				_ = _rhs115
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 _dafny.Array = _789_v133
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))
				_ = _lhs100
				_lhs98.F18 = _rhs114
				(_lhs99).ArraySet1(_rhs115, (_lhs100).Int())
				var _rhs116 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F21())).Update(p1, !((_this).F26()))
				_ = _rhs116
				var _rhs117 _dafny.Int = (_dafny.Zero).Minus((_789_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))).Int()).(_dafny.Int))
				_ = _rhs117
				var _rhs118 _dafny.Sequence = _786_v131
				_ = _rhs118
				var _rhs119 bool = p1
				_ = _rhs119
				var _rhs120 bool = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Equals(Companion_Default___.Fm15((_789_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))).Int()).(_dafny.Int), (_789_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(953), _dafny.IntOfUint32((_786_v131).Cardinality()), globalState))
				_ = _rhs120
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				_781_v124 = _rhs116
				_776_cf10 = _rhs117
				_786_v131 = _rhs118
				_lhs101.F3 = _rhs119
				_lhs102.F3 = _rhs120
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_789_v133), 0))
				_ = _index122
				(_789_v133).ArraySet1((_dafny.Zero).Minus(p0), (_index122).Int())
			} else if _source11.Is_DC7() {
				var _791___mcc_h12 _dafny.Sequence = _source11.Get_().(D3_DC7).Cf11
				_ = _791___mcc_h12
				var _792___mcc_h13 bool = _source11.Get_().(D3_DC7).Cf12
				_ = _792___mcc_h13
				var _793_cf12 bool = _792___mcc_h13
				_ = _793_cf12
				var _794_cf11 _dafny.Sequence = _791___mcc_h12
				_ = _794_cf11
				var _795_v134 _dafny.Array
				_ = _795_v134
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw127
				_795_v134 = _nw127
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))
				_ = _index123
				(_795_v134).ArraySet1((_this).F20(), (_index123).Int())
				var _796_v135 _dafny.Array
				_ = _796_v135
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_23
				var _nw128 _dafny.Array
				_ = _nw128
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw128 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) D2 = (func(_797_p0 _dafny.Int, _798_cf12 bool) func(_dafny.Int) D2 {
						return func(_799_i14 _dafny.Int) D2 {
							return Companion_D2_.Create_DC3_(_797_p0, (_this).F26(), _798_cf12, _dafny.IntOfInt64(593), (_this).F26())
						}
					})(p0, _793_cf12)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw128 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw128).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw128).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_796_v135 = _nw128
				var _800_v136 D2
				_ = _800_v136
				_800_v136 = Companion_D2_.Create_DC3_((_this).F25(), p1, false, (_this).F25(), (_this).F21())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_796_v135), 0))
				_ = _index124
				(_796_v135).ArraySet1(_800_v136, (_index124).Int())
				var _801_v137 _dafny.Map
				_ = _801_v137
				_801_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _770_v116)
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))
				_ = _index125
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_796_v135), 0))
				_ = _index126
				var _rhs121 _dafny.Int = (_801_v137).Cardinality()
				_ = _rhs121
				var _rhs122 D2 = _800_v136
				_ = _rhs122
				var _rhs123 bool = p1
				_ = _rhs123
				var _rhs124 _dafny.Int = (_768_v114).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.IntOfUint32((_768_v114).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs124
				var _lhs103 _dafny.Array = _795_v134
				_ = _lhs103
				var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))
				_ = _lhs104
				var _lhs105 _dafny.Array = _796_v135
				_ = _lhs105
				var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_796_v135), 0))
				_ = _lhs106
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				(_lhs103).ArraySet1(_rhs121, (_lhs104).Int())
				(_lhs105).ArraySet1(_rhs122, (_lhs106).Int())
				_lhs107.F3 = _rhs123
				_lhs108.F17 = _rhs124
				var _802_v138 _dafny.Sequence
				_ = _802_v138
				_802_v138 = _dafny.UnicodeSeqOfUtf8Bytes("qgfualrx")
				(globalState).F5 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_802_v138, _802_v138)).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if _793_cf12 {
						return _dafny.SeqOf(_dafny.IntOfInt64(756))
					}
					return _768_v114
				})(), (Companion_Default___.SafeIndex((_768_v114).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_768_v114).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _793_cf12 {
						return _dafny.SeqOf(_dafny.IntOfInt64(756))
					}
					return _768_v114
				})()).Cardinality()))).Uint32(), (_795_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))).Int()).(_dafny.Int))).Cardinality()))
				var _803_v139 _dafny.Map
				_ = _803_v139
				_803_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F26())
				var _804_v140 _dafny.Map
				_ = _804_v140
				_804_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v138, _803_v139)
				var _805_v141 _dafny.Sequence
				_ = _805_v141
				_805_v141 = _dafny.SeqOf((_this).F26())
				var _806_v142 _dafny.Map
				_ = _806_v142
				_806_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_804_v140, _dafny.Companion_Sequence_.IsPrefixOf(_805_v141, _805_v141))
				_806_v142 = (_806_v142).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v138, _803_v139), (_this).F26())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_767_cf14), 0))
				_ = _index127
				(_767_cf14).ArraySet1(_793_cf12, (_index127).Int())
				var _807_v143 _dafny.Set
				_ = _807_v143
				_807_v143 = _dafny.SetOf(_771_v117, _771_v117)
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_767_cf14), 0))
				_ = _index128
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))
				_ = _index129
				var _rhs125 bool = _793_cf12
				_ = _rhs125
				var _rhs126 bool = ((_795_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))).Int()).(_dafny.Int)).Cmp((_795_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))).Int()).(_dafny.Int)) <= 0
				_ = _rhs126
				var _rhs127 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(((_795_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_802_v138, _802_v138), (Companion_Default___.SafeIndex((_807_v143).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_802_v138, _802_v138)).Cardinality()))).Uint32(), _771_v117)).Cardinality()))))
				_ = _rhs127
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 _dafny.Array = _767_cf14
				_ = _lhs110
				var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_767_cf14), 0))
				_ = _lhs111
				var _lhs112 _dafny.Array = _795_v134
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_795_v134), 0))
				_ = _lhs113
				_lhs109.F14 = _rhs125
				(_lhs110).ArraySet1(_rhs126, (_lhs111).Int())
				(_lhs112).ArraySet1(_rhs127, (_lhs113).Int())
			} else if _source11.Is_DC5() {
				var _808___mcc_h14 _dafny.Sequence = _source11.Get_().(D3_DC5).Cf9
				_ = _808___mcc_h14
				var _809_cf9 _dafny.Sequence = _808___mcc_h14
				_ = _809_cf9
				var _810_v144 _dafny.Map
				_ = _810_v144
				_810_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _768_v114)
				_810_v144 = (_810_v144).Update(!(p1), _dafny.SeqOf(p0, p0, (_this).F25()))
				(globalState).F5 = (_dafny.Zero).Minus((_this).F25())
				var _811_v145 _dafny.Sequence
				_ = _811_v145
				_811_v145 = _dafny.UnicodeSeqOfUtf8Bytes("pd")
				var _812_v146 _dafny.Array
				_ = _812_v146
				var _nwElement0_32 _dafny.Int = (_this).F20()
				_ = _nwElement0_32
				var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(2))
				_ = _nw129
				(_nw129).ArraySet1(_nwElement0_32, 0)
				(_nw129).ArraySet1((_this).F20(), 1)
				_812_v146 = _nw129
				var _813_v147 _dafny.Map
				_ = _813_v147
				_813_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_811_v145, _812_v146)
				var _814_v148 *C1
				_ = _814_v148
				var _nw130 *C1 = New_C1_()
				_ = _nw130
				_nw130.Ctor__(p0, (func() _dafny.Array {
					if (_813_v147).Contains(_811_v145) {
						return (_813_v147).Get(_811_v145).(_dafny.Array)
					}
					return _812_v146
				})(), Companion_Default___.SafeDivisionInt((_this).F25(), _dafny.IntOfInt64(656)), p1)
				_814_v148 = _nw130
				var _815_v149 *C1
				_ = _815_v149
				var _nw131 *C1 = New_C1_()
				_ = _nw131
				_nw131.Ctor__(((_dafny.Zero).Minus((_this).F25())).Times(p0), (_814_v148).F28(), ((_this).F25()).Plus(_dafny.IntOfInt64(316)), (_this).F26())
				_815_v149 = _nw131
			} else {
				var _816___mcc_h15 D3 = _source11.Get_().(D3_DC8).Cf13
				_ = _816___mcc_h15
				var _817_cf13 D3 = _816___mcc_h15
				_ = _817_cf13
				var _rhs128 _dafny.Int = p0
				_ = _rhs128
				var _rhs129 bool = true
				_ = _rhs129
				var _rhs130 _dafny.Int = (_770_v116).Fm1(p0, (_this).F25(), (_this).F26(), globalState)
				_ = _rhs130
				var _lhs114 *GlobalState = globalState
				_ = _lhs114
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				_lhs114.F17 = _rhs128
				_lhs115.F18 = _rhs129
				_lhs116.F17 = _rhs130
				var _818_v151 _dafny.Array
				_ = _818_v151
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_24
				var _nw132 _dafny.Array
				_ = _nw132
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw132 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Set = (func(_819_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_820_i15 _dafny.Int) _dafny.Set {
							return func() _dafny.Set {
								var _coll27 = _dafny.NewBuilder()
								_ = _coll27
								for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(439), _dafny.IntOfInt64(582))); ; {
									_compr_27, _ok28 := _iter28()
									if !_ok28 {
										break
									}
									var _821_v150 _dafny.Int
									_821_v150 = interface{}(_compr_27).(_dafny.Int)
									if ((_dafny.IntOfInt64(439)).Cmp(_821_v150) <= 0) && ((_821_v150).Cmp(_dafny.IntOfInt64(582)) < 0) {
										_coll27.Add((_821_v150).Plus(_819_p0))
									}
								}
								return _coll27.ToSet()
							}()
						}
					})(p0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw132 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw132).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw132).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_818_v151 = _nw132
				var _822_v152 D8
				_ = _822_v152
				_822_v152 = Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20()))
				var _823_v153 _dafny.Sequence
				_ = _823_v153
				_823_v153 = _dafny.UnicodeSeqOfUtf8Bytes("fvwgksww")
				var _824_v154 _dafny.Map
				_ = _824_v154
				_824_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), true)
				var _825_v155 _dafny.Map
				_ = _825_v155
				_825_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Update((_770_v116).Fm0((_822_v152).Dtor_cf28(), _dafny.Companion_Sequence_.Update(_823_v153, (Companion_Default___.SafeIndex((_770_v116).Fm1((_this).F20(), p0, p1, globalState), _dafny.IntOfUint32((_823_v153).Cardinality()))).Uint32(), _771_v117), _dafny.SeqOf(_824_v154, _824_v154, _824_v154), globalState), (_this).F21()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_768_v114).Cardinality()), (_this).F26()))
				var _rhs131 _dafny.Array = _818_v151
				_ = _rhs131
				var _rhs132 _dafny.Map = (_825_v155).Merge(_825_v155)
				_ = _rhs132
				var _rhs133 bool = (_this).F21()
				_ = _rhs133
				var _lhs117 *GlobalState = globalState
				_ = _lhs117
				_818_v151 = _rhs131
				_825_v155 = _rhs132
				_lhs117.F3 = _rhs133
				var _826_v156 _dafny.Map
				_ = _826_v156
				_826_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_767_cf14, (_dafny.Zero).Minus((_this).F25()))
				var _827_v157 _dafny.Sequence
				_ = _827_v157
				_827_v157 = _dafny.SeqOf(_826_v156, _826_v156, (_826_v156).Merge(_826_v156), (_826_v156).Merge(_826_v156))
				_826_v156 = (_827_v157).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-318), _dafny.IntOfUint32((_827_v157).Cardinality()))).Uint32()).(_dafny.Map)
				var _828_v158 D4
				_ = _828_v158
				_828_v158 = Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqebepvf")).Cardinality()), p1, p1)
				var _829_v159 _dafny.Map
				_ = _829_v159
				_829_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v118, p0)
				var _830_v160 _dafny.Array
				_ = _830_v160
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw133
				_830_v160 = _nw133
				var _831_v161 _dafny.MultiSet
				_ = _831_v161
				_831_v161 = _dafny.MultiSetOf(_830_v160)
				var _832_v162 _dafny.Array
				_ = _832_v162
				var _nwElement0_33 _dafny.Int = (_828_v158).Dtor_cf17()
				_ = _nwElement0_33
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(12))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_33, 0)
				(_nw134).ArraySet1((_this).Fm1((_this).F20(), (_this).F25(), p1, globalState), 1)
				(_nw134).ArraySet1(p0, 2)
				(_nw134).ArraySet1(((func() _dafny.Int {
					if (_829_v159).Contains(_772_v118) {
						return (_829_v159).Get(_772_v118).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_831_v161).Contains(_830_v160) {
							return (_831_v161).Multiplicity(_830_v160)
						}
						return (_this).F25()
					})()
				})()).Times((_this).F25()), 3)
				(_nw134).ArraySet1(_dafny.IntOfInt64(489), 4)
				(_nw134).ArraySet1(p0, 5)
				(_nw134).ArraySet1(p0, 6)
				(_nw134).ArraySet1(Companion_Default___.SafeModuloInt((_this).F20(), (_824_v154).Cardinality()), 7)
				(_nw134).ArraySet1(_dafny.IntOfInt64(-12), 8)
				(_nw134).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F20(), (_this).F25())).Cardinality()), 9)
				(_nw134).ArraySet1(((_this).F25()).Plus((_dafny.Zero).Minus(p0)), 10)
				(_nw134).ArraySet1(p0, 11)
				_832_v162 = _nw134
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_832_v162), 0))
				_ = _index130
				(_832_v162).ArraySet1(p0, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_832_v162), 0))
				_ = _index131
				(_832_v162).ArraySet1((_this).F20(), (_index131).Int())
			}
		}
		var _833_v163 _dafny.Sequence
		_ = _833_v163
		_833_v163 = _dafny.UnicodeSeqOfUtf8Bytes("uftkj")
		var _834_v164 _dafny.Sequence
		_ = _834_v164
		_834_v164 = _dafny.SeqOf(_833_v163, (func() _dafny.Sequence {
			if (_this).F21() {
				return _833_v163
			}
			return _833_v163
		})())
		var _835_v165 _dafny.MultiSet
		_ = _835_v165
		_835_v165 = _dafny.MultiSetOf((_this).F25())
		var _rhs134 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_835_v165).Cardinality(), (_this).F25()), (_this).F20())
		_ = _rhs134
		var _rhs135 _dafny.Int = (_dafny.Zero).Minus((((_this).F20()).Times((_this).F25())).Times((_this).F20()))
		_ = _rhs135
		var _rhs136 _dafny.Sequence = _834_v164
		_ = _rhs136
		var _rhs137 bool = (_this).F26()
		_ = _rhs137
		var _rhs138 _dafny.Int = p0
		_ = _rhs138
		var _lhs118 *GlobalState = globalState
		_ = _lhs118
		var _lhs119 *GlobalState = globalState
		_ = _lhs119
		var _lhs120 *GlobalState = globalState
		_ = _lhs120
		var _lhs121 *GlobalState = globalState
		_ = _lhs121
		_lhs118.F5 = _rhs134
		_lhs119.F17 = _rhs135
		_834_v164 = _rhs136
		_lhs120.F14 = _rhs137
		_lhs121.F5 = _rhs138
		var _836_v166 _dafny.CodePoint
		_ = _836_v166
		_836_v166 = _dafny.CodePoint('w')
		var _837_v167 _dafny.MultiSet
		_ = _837_v167
		_837_v167 = _dafny.MultiSetOf(p1)
		var _838_v168 _dafny.Sequence
		_ = _838_v168
		_838_v168 = _dafny.SeqOf((_this).F21(), (_this).F21())
		var _839_v169 _dafny.Set
		_ = _839_v169
		_839_v169 = _dafny.SetOf((_this).F26())
		var _840_v170 _dafny.Sequence
		_ = _840_v170
		_840_v170 = _dafny.SeqOf((_this).F20(), (_this).F20())
		var _841_v171 D6
		_ = _841_v171
		_841_v171 = Companion_D6_.Create_DC13_(_833_v163)
		var _842_v172 _dafny.Map
		_ = _842_v172
		_842_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-993), _dafny.IntOfUint32(((_841_v171).Dtor_cf21()).Cardinality()))
		var _843_v174 D8
		_ = _843_v174
		_843_v174 = Companion_D8_.Create_DC19_(func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(377), _dafny.IntOfInt64(479))); ; {
				_compr_28, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _844_v173 _dafny.Int
				_844_v173 = interface{}(_compr_28).(_dafny.Int)
				if ((_dafny.IntOfInt64(377)).Cmp(_844_v173) <= 0) && ((_844_v173).Cmp(_dafny.IntOfInt64(479)) < 0) {
					_coll28.Add(Companion_Default___.SafeDivisionInt(_844_v173, (_this).F20()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F20())).Cardinality())
				}
			}
			return _coll28.ToMap()
		}())
		var _845_v175 _dafny.Array
		_ = _845_v175
		var _nwElement0_34 _dafny.Int = (_this).F25()
		_ = _nwElement0_34
		var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(27))
		_ = _nw135
		(_nw135).ArraySet1(_nwElement0_34, 0)
		(_nw135).ArraySet1((_this).F25(), 1)
		(_nw135).ArraySet1((_this).F25(), 2)
		(_nw135).ArraySet1((_this).F20(), 3)
		(_nw135).ArraySet1((func() _dafny.Int {
			if (_837_v167).Contains((_this).F21()) {
				return (_837_v167).Multiplicity((_this).F21())
			}
			return p0
		})(), 4)
		(_nw135).ArraySet1(p0, 5)
		(_nw135).ArraySet1((_this).F20(), 6)
		(_nw135).ArraySet1(_dafny.IntOfInt64(489), 7)
		(_nw135).ArraySet1((Companion_Default___.Fm18((_this).F20(), (_this).F20(), _836_v166, globalState)).Cardinality(), 8)
		(_nw135).ArraySet1(_dafny.IntOfInt64(248), 9)
		(_nw135).ArraySet1((_this).F25(), 10)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_838_v168).Cardinality()), 11)
		(_nw135).ArraySet1((_839_v169).Cardinality(), 12)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("esaxrsb")).Cardinality()), 13)
		(_nw135).ArraySet1(_dafny.IntOfInt64(268), 14)
		(_nw135).ArraySet1(p0, 15)
		(_nw135).ArraySet1((_this).F25(), 16)
		(_nw135).ArraySet1((_this).F20(), 17)
		(_nw135).ArraySet1((_this).F25(), 18)
		(_nw135).ArraySet1(p0, 19)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_840_v170).Cardinality()), 20)
		(_nw135).ArraySet1((_this).F20(), 21)
		(_nw135).ArraySet1(_dafny.IntOfInt64(48), 22)
		(_nw135).ArraySet1((_this).F25(), 23)
		(_nw135).ArraySet1((_842_v172).Cardinality(), 24)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_838_v168).Cardinality()), 25)
		(_nw135).ArraySet1(((_843_v174).Dtor_cf28()).Cardinality(), 26)
		_845_v175 = _nw135
		var _846_v176 *C1
		_ = _846_v176
		var _nw136 *C1 = New_C1_()
		_ = _nw136
		_nw136.Ctor__((_this).F20(), _845_v175, (_this).F20(), (_this).F21())
		_846_v176 = _nw136
		var _847_v177 _dafny.Map
		_ = _847_v177
		_847_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_v166, _846_v176)
		var _848_v178 _dafny.Sequence
		_ = _848_v178
		_848_v178 = _dafny.SeqOf(_847_v177)
		var _849_v179 D10
		_ = _849_v179
		_849_v179 = Companion_D10_.Create_DC27_(_dafny.IntOfUint32((_833_v163).Cardinality()), _848_v178, _dafny.CodePoint('p'))
		var _850_v180 _dafny.Array
		_ = _850_v180
		var _nwElement0_35 _dafny.CodePoint = _836_v166
		_ = _nwElement0_35
		var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(16))
		_ = _nw137
		(_nw137).ArraySet1CodePoint(_nwElement0_35, 0)
		(_nw137).ArraySet1CodePoint(_836_v166, 1)
		(_nw137).ArraySet1CodePoint(_836_v166, 2)
		(_nw137).ArraySet1CodePoint(_836_v166, 3)
		(_nw137).ArraySet1CodePoint(_836_v166, 4)
		(_nw137).ArraySet1CodePoint(_836_v166, 5)
		(_nw137).ArraySet1CodePoint(_836_v166, 6)
		(_nw137).ArraySet1CodePoint((_849_v179).Dtor_cf38(), 7)
		(_nw137).ArraySet1CodePoint(_836_v166, 8)
		(_nw137).ArraySet1CodePoint(_836_v166, 9)
		(_nw137).ArraySet1CodePoint(_836_v166, 10)
		(_nw137).ArraySet1CodePoint(_836_v166, 11)
		(_nw137).ArraySet1CodePoint(_836_v166, 12)
		(_nw137).ArraySet1CodePoint(_836_v166, 13)
		(_nw137).ArraySet1CodePoint(_836_v166, 14)
		(_nw137).ArraySet1CodePoint(_dafny.CodePoint('e'), 15)
		_850_v180 = _nw137
		for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_850_v180), 0))); ; {
			_guard_loop_1, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _851_i16 _dafny.Int
			_851_i16 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_851_i16).Sign() != -1) && ((_851_i16).Cmp(_dafny.ArrayLen((_850_v180), 0)) < 0)) {
				(_850_v180).ArraySet1CodePoint((_833_v163).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v170, _846_v176.F27)).Cardinality(), _dafny.IntOfUint32((_833_v163).Cardinality()))).Uint32()).(_dafny.CodePoint), (_851_i16).Int())
			}
		}
		_834_v164 = Companion_Default___.Fm33(_846_v176.F27, (_this).F21(), _dafny.Companion_Sequence_.Concatenate(_833_v163, _833_v163), globalState)
		var _852_i17 _dafny.Int
		_ = _852_i17
		_852_i17 = _dafny.Zero
		{
			for (_this).F26() {
				{
					if (_852_i17).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_852_i17 = (_852_i17).Plus(_dafny.One)
					var _853_v181 _dafny.Map
					_ = _853_v181
					_853_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
					var _854_v182 _dafny.Map
					_ = _854_v182
					_854_v182 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F26()) || ((func() bool {
						if (_853_v181).Contains(false) {
							return (_853_v181).Get(false).(bool)
						}
						return (_this).F21()
					})()), (_dafny.IntOfInt64(290)).Minus((_this).F25()))
					(_846_v176).F27 = (func() _dafny.Int {
						if (_854_v182).Contains((_this).F26()) {
							return (_854_v182).Get((_this).F26()).(_dafny.Int)
						}
						return (func() _dafny.Int {
							if (_842_v172).Contains((_this).F20()) {
								return (_842_v172).Get((_this).F20()).(_dafny.Int)
							}
							return (_this).F20()
						})()
					})()
					(globalState).F5 = Companion_Default___.SafeModuloInt(_846_v176.F27, _846_v176.F27)
					var _855_v183 _dafny.Map
					_ = _855_v183
					_855_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_v165, p1)
					(globalState).F5 = Companion_Default___.SafeModuloInt((_846_v176).Fm1(_846_v176.F27, p0, (func() bool {
						if (_855_v183).Contains(_835_v165) {
							return (_855_v183).Get(_835_v165).(bool)
						}
						return (_this).F26()
					})(), globalState), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _dafny.IntOfInt64(793))).Merge(_854_v182)).Cardinality())
					var _856_v184 _dafny.Map
					_ = _856_v184
					_856_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F26())
					_856_v184 = (_856_v184).Update(p0, (func() bool {
						if (_856_v184).Contains(_846_v176.F27) {
							return (_856_v184).Get(_846_v176.F27).(bool)
						}
						return (_this).F21()
					})())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		(globalState).F14 = false
		var _857_v185 _dafny.Sequence
		_ = _857_v185
		_857_v185 = _dafny.SeqOf((_846_v176).F28())
		var _858_v186 _dafny.Sequence
		_ = _858_v186
		_858_v186 = _dafny.SeqOf(_857_v185, _857_v185)
		r0 = (_858_v186).Select((Companion_Default___.SafeIndex(((_this).F20()).Minus(_dafny.IntOfInt64(463)), _dafny.IntOfUint32((_858_v186).Cardinality()))).Uint32()).(_dafny.Sequence)
		r1 = true
		return r0, r1
	}
}
func (_this *C3) M2(p0 _dafny.Array, p1 _dafny.Array, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _859_v0 _dafny.Set
		_ = _859_v0
		_859_v0 = _dafny.SetOf((_this).F26(), (_this).F26())
		var _860_v1 _dafny.Map
		_ = _860_v1
		_860_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_859_v0).Cardinality(), _dafny.IntOfInt64(856))
		var _861_v2 _dafny.Map
		_ = _861_v2
		_861_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F20())), true)
		var _862_v3 _dafny.Sequence
		_ = _862_v3
		_862_v3 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F20()), (_861_v2).Cardinality(), (_this).F20(), (_this).F20())
		_860_v1 = (_860_v1).Update(_dafny.IntOfUint32((_862_v3).Cardinality()), (_dafny.Zero).Minus((_this).F25()))
		var _863_v4 _dafny.Sequence
		_ = _863_v4
		_863_v4 = _dafny.SeqOf((_this).F21(), (_this).F26())
		if (_863_v4).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_863_v4).Cardinality()))).Uint32()).(bool) {
			var _864_v5 _dafny.Array
			_ = _864_v5
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_25
			var _nw138 _dafny.Array
			_ = _nw138
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw138 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Sequence = func(_865_i0 _dafny.Int) _dafny.Sequence {
					return (func() _dafny.Sequence {
						if (_this).F21() {
							return _dafny.UnicodeSeqOfUtf8Bytes("rjsrayy")
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("asut")
					})()
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw138 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw138).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw138).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_864_v5 = _nw138
			var _866_v6 _dafny.Sequence
			_ = _866_v6
			_866_v6 = _dafny.UnicodeSeqOfUtf8Bytes("qxyux")
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_864_v5), 0))
			_ = _index132
			(_864_v5).ArraySet1(_866_v6, (_index132).Int())
			var _867_v7 _dafny.CodePoint
			_ = _867_v7
			_867_v7 = _dafny.CodePoint('l')
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_864_v5), 0))
			_ = _index133
			var _rhs139 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_868_i1 _dafny.Int) _dafny.Int {
				return (_this).F25()
			})), _dafny.SeqOf((_this).F25()))).Cardinality())).Times(_dafny.IntOfInt64(821)))
			_ = _rhs139
			var _rhs140 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_866_v6, _866_v6), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_866_v6, _866_v6)).Cardinality()))).Uint32(), _867_v7), _866_v6)
			_ = _rhs140
			var _rhs141 _dafny.Int = (_this).F25()
			_ = _rhs141
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			var _lhs123 _dafny.Array = _864_v5
			_ = _lhs123
			var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_864_v5), 0))
			_ = _lhs124
			var _lhs125 *GlobalState = globalState
			_ = _lhs125
			_lhs122.F17 = _rhs139
			(_lhs123).ArraySet1(_rhs140, (_lhs124).Int())
			_lhs125.F17 = _rhs141
			(globalState).F18 = (_this).F26()
			var _869_v8 D2
			_ = _869_v8
			_869_v8 = Companion_D2_.Create_DC3_((_this).F20(), (_this).F26(), !(!((_this).F21())), (_this).F20(), (_this).F26())
			var _870_v9 _dafny.Sequence
			_ = _870_v9
			_870_v9 = _dafny.SeqOf(_859_v0)
			var _871_v10 D3
			_ = _871_v10
			_871_v10 = Companion_D3_.Create_DC7_(_870_v9, (_this).F26())
			var _pat_let_tv22 = _870_v9
			_ = _pat_let_tv22
			var _source12 D3 = func(_pat_let22_0 D3) D3 {
				return func(_872_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let23_0 _dafny.Sequence) D3 {
						return func(_873_dt__update_hcf11_h0 _dafny.Sequence) D3 {
							return Companion_D3_.Create_DC7_(_873_dt__update_hcf11_h0, (_872_dt__update__tmp_h0).Dtor_cf12())
						}(_pat_let23_0)
					}(_pat_let_tv22)
				}(_pat_let22_0)
			}((func() D3 {
				if (_869_v8).Dtor_cf7() {
					return _871_v10
				}
				return Companion_Default___.Fm34(_dafny.CodePoint('v'), (_this).F21(), (_this).F20(), (_this).F25(), globalState)
			})())
			_ = _source12
			if _source12.Is_DC6() {
				var _874___mcc_h0 _dafny.Int = _source12.Get_().(D3_DC6).Cf10
				_ = _874___mcc_h0
				var _875_cf10 _dafny.Int = _874___mcc_h0
				_ = _875_cf10
				_867_v7 = _867_v7
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((p1), 0))
				_ = _index134
				(p1).ArraySet1((func() _dafny.Int {
					if (_860_v1).Contains(_875_cf10) {
						return (_860_v1).Get(_875_cf10).(_dafny.Int)
					}
					return _875_cf10
				})(), (_index134).Int())
				var _876_v11 _dafny.Set
				_ = _876_v11
				_876_v11 = _dafny.SetOf((_this).F25())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((p1), 0))
				_ = _index135
				(p1).ArraySet1((_dafny.Zero).Minus(((Companion_Default___.Fm25(globalState)).Intersection(_876_v11)).Cardinality()), (_index135).Int())
				var _877_v12 _dafny.Array
				_ = _877_v12
				var _nwElement0_36 bool = (_this).F21()
				_ = _nwElement0_36
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(4))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_36, 0)
				(_nw139).ArraySet1((_this).F26(), 1)
				(_nw139).ArraySet1((_this).F21(), 2)
				(_nw139).ArraySet1(true, 3)
				_877_v12 = _nw139
				var _878_v13 _dafny.Sequence
				_ = _878_v13
				_878_v13 = _dafny.SeqOf(_877_v12)
				var _879_v14 _dafny.MultiSet
				_ = _879_v14
				_879_v14 = _dafny.MultiSetOf((_878_v13).Select((Companion_Default___.SafeIndex((Companion_D2_.Create_DC3_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), !((_this).F21()), false, (_this).F20(), (_this).F21())).Dtor_cf3(), _dafny.IntOfUint32((_878_v13).Cardinality()))).Uint32()).(_dafny.Array), _877_v12)
				(globalState).F4 = (_879_v14).IsSubsetOf(_879_v14)
				(globalState).F5 = (_this).F20()
			} else if _source12.Is_DC7() {
				var _880___mcc_h1 _dafny.Sequence = _source12.Get_().(D3_DC7).Cf11
				_ = _880___mcc_h1
				var _881___mcc_h2 bool = _source12.Get_().(D3_DC7).Cf12
				_ = _881___mcc_h2
				var _882_cf12 bool = _881___mcc_h2
				_ = _882_cf12
				var _883_cf11 _dafny.Sequence = _880___mcc_h1
				_ = _883_cf11
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((p1), 0))
				_ = _index136
				(p1).ArraySet1((_this).F20(), (_index136).Int())
				var _884_v15 *C0
				_ = _884_v15
				var _nw140 *C0 = New_C0_()
				_ = _nw140
				_nw140.Ctor__()
				_884_v15 = _nw140
				var _885_v16 _dafny.Set
				_ = _885_v16
				_885_v16 = _dafny.SetOf((_this).F25())
				var _886_v17 D4
				_ = _886_v17
				_886_v17 = Companion_D4_.Create_DC11_((_this).F25(), _882_cf12, (_this).F26())
				var _887_v18 _dafny.Map
				_ = _887_v18
				_887_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _886_v17)
				var _888_v19 _dafny.Map
				_ = _888_v19
				_888_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _887_v18)
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((p1), 0))
				_ = _index137
				var _rhs142 bool = !((_884_v15).Fm3((_this).F20(), (_this).F21(), globalState))
				_ = _rhs142
				var _rhs143 _dafny.Int = ((_888_v19).Merge(_888_v19)).Cardinality()
				_ = _rhs143
				var _rhs144 *C0 = _884_v15
				_ = _rhs144
				var _rhs145 _dafny.Set = (_885_v16).Union((_dafny.SetOf((_this).F25())).Union(_885_v16))
				_ = _rhs145
				var _rhs146 bool = _882_cf12
				_ = _rhs146
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				var _lhs127 _dafny.Array = p1
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((p1), 0))
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				_lhs126.F18 = _rhs142
				(_lhs127).ArraySet1(_rhs143, (_lhs128).Int())
				_884_v15 = _rhs144
				_885_v16 = _rhs145
				_lhs129.F3 = _rhs146
				(globalState).F17 = (_this).F20()
				(globalState).F17 = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ysfj")).Cardinality())).Plus((_this).F25())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((p0), 0))
				_ = _index138
				(p0).ArraySet1(_860_v1, (_index138).Int())
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((p0), 0))
				_ = _index139
				(p0).ArraySet1(_860_v1, (_index139).Int())
			} else if _source12.Is_DC5() {
				var _889___mcc_h3 _dafny.Sequence = _source12.Get_().(D3_DC5).Cf9
				_ = _889___mcc_h3
				var _890_cf9 _dafny.Sequence = _889___mcc_h3
				_ = _890_cf9
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((p1), 0))
				_ = _index140
				(p1).ArraySet1((_this).F25(), (_index140).Int())
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((p1), 0))
				_ = _index141
				(p1).ArraySet1(_dafny.IntOfUint32((_862_v3).Cardinality()), (_index141).Int())
				(globalState).F17 = _dafny.IntOfInt64(-305)
				var _891_v20 *C0
				_ = _891_v20
				var _nw141 *C0 = New_C0_()
				_ = _nw141
				_nw141.Ctor__()
				_891_v20 = _nw141
				var _892_v21 _dafny.Map
				_ = _892_v21
				_892_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _891_v20)
				var _893_v22 _dafny.Set
				_ = _893_v22
				_893_v22 = _dafny.SetOf(_891_v20, (func() *C0 {
					if (_892_v21).Contains((_this).F25()) {
						return (_892_v21).Get((_this).F25()).(*C0)
					}
					return _891_v20
				})(), _891_v20)
				var _rhs147 _dafny.Set = _893_v22
				_ = _rhs147
				var _rhs148 bool = (_this).F21()
				_ = _rhs148
				var _rhs149 _dafny.Int = (func() _dafny.Int {
					if (_860_v1).Contains((_this).F20()) {
						return (_860_v1).Get((_this).F20()).(_dafny.Int)
					}
					return (_this).F25()
				})()
				_ = _rhs149
				var _rhs150 _dafny.Int = (_this).F25()
				_ = _rhs150
				var _rhs151 _dafny.Int = (func() _dafny.Int {
					if (_this).F21() {
						return (_862_v3).Select((Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32()).(_dafny.Int)
					}
					return ((_this).F25()).Minus((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
				})()
				_ = _rhs151
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				_893_v22 = _rhs147
				_lhs130.F3 = _rhs148
				_lhs131.F5 = _rhs149
				_lhs132.F17 = _rhs150
				_lhs133.F17 = _rhs151
				(globalState).F4 = (_this).F26()
			} else {
				var _894___mcc_h4 D3 = _source12.Get_().(D3_DC8).Cf13
				_ = _894___mcc_h4
				var _895_cf13 D3 = _894___mcc_h4
				_ = _895_cf13
				(globalState).F17 = (_dafny.Zero).Minus((_860_v1).Cardinality())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_864_v5), 0))
				_ = _index142
				(_864_v5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm17((_this).F26(), globalState)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_867_v7, _dafny.CodePoint('u')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-627))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_896_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_897_i2 _dafny.Int) _dafny.CodePoint {
						return _896_v7
					}
				})(_867_v7))))), (_index142).Int())
				var _898_v23 _dafny.Array
				_ = _898_v23
				var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw142
				_898_v23 = _nw142
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_26
				var _nw143 _dafny.Array
				_ = _nw143
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw143 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = func(_899_i3 _dafny.Int) bool {
						return ((_this).F20()).Cmp((_this).F25()) > 0
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw143 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw143).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw143).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_898_v23 = _nw143
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p1), 0))
				_ = _index143
				(p1).ArraySet1((_this).F20(), (_index143).Int())
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p1), 0))
				_ = _index144
				(p1).ArraySet1((_this).Fm1((_dafny.Zero).Minus((_this).Fm1((_this).F20(), (_this).F25(), !((_this).F26()), globalState)), _dafny.IntOfInt64(900), false, globalState), (_index144).Int())
			}
			(globalState).F17 = (_this).F25()
			_866_v6 = (_864_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_864_v5), 0))).Int()).(_dafny.Sequence)
		} else {
			var _900_v24 _dafny.Array
			_ = _900_v24
			var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw144
			_900_v24 = _nw144
			var _901_v25 _dafny.Array
			_ = _901_v25
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw145
			_901_v25 = _nw145
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_901_v25), 0))
			_ = _index145
			(_901_v25).ArraySet1((_this).F26(), (_index145).Int())
			var _902_v26 _dafny.MultiSet
			_ = _902_v26
			_902_v26 = _dafny.MultiSetOf((_this).F21())
			var _903_v27 D11
			_ = _903_v27
			_903_v27 = Companion_D11_.Create_DC31_(Companion_D11_.Create_DC30_((_this).F26(), (_902_v26).Update(true, Companion_Default___.Abs((_this).F20()))))
			var _904_v28 _dafny.Map
			_ = _904_v28
			_904_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v27, _dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality()))
			var _905_v29 _dafny.Sequence
			_ = _905_v29
			_905_v29 = _dafny.SeqOf((_904_v28).Merge(_904_v28), _904_v28)
			var _906_v30 _dafny.Sequence
			_ = _906_v30
			_906_v30 = _dafny.UnicodeSeqOfUtf8Bytes("uielcygj")
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_901_v25), 0))
			_ = _index146
			var _rhs152 _dafny.Array = _900_v24
			_ = _rhs152
			var _rhs153 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ie"), _dafny.Companion_Sequence_.Concatenate(_906_v30, _906_v30))
			_ = _rhs153
			var _rhs154 bool = false
			_ = _rhs154
			var _rhs155 _dafny.Sequence = _905_v29
			_ = _rhs155
			var _rhs156 _dafny.Sequence = _906_v30
			_ = _rhs156
			var _lhs134 *GlobalState = globalState
			_ = _lhs134
			var _lhs135 _dafny.Array = _901_v25
			_ = _lhs135
			var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_901_v25), 0))
			_ = _lhs136
			var _lhs137 *GlobalState = globalState
			_ = _lhs137
			_900_v24 = _rhs152
			_lhs134.F15 = _rhs153
			(_lhs135).ArraySet1(_rhs154, (_lhs136).Int())
			_905_v29 = _rhs155
			_lhs137.F15 = _rhs156
			(globalState).F4 = (_901_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_901_v25), 0))).Int()).(bool)
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((p1), 0))
			_ = _index147
			(p1).ArraySet1((_this).F25(), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((p1), 0))
			_ = _index148
			(p1).ArraySet1((_this).F25(), (_index148).Int())
			var _907_v31 _dafny.Array
			_ = _907_v31
			var _nw146 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw146
			_907_v31 = _nw146
			var _908_v32 T0
			_ = _908_v32
			var _nw147 *C2 = New_C2_()
			_ = _nw147
			_nw147.Ctor__(_dafny.IntOfUint32((_906_v30).Cardinality()), (_901_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_901_v25), 0))).Int()).(bool))
			_908_v32 = _nw147
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_907_v31), 0))
			_ = _index149
			(_907_v31).ArraySet1(_908_v32, (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_907_v31), 0))
			_ = _index150
			(_907_v31).ArraySet1(_908_v32, (_index150).Int())
			var _909_v33 _dafny.Set
			_ = _909_v33
			_909_v33 = _dafny.SetOf(_860_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_908_v32).F20(), (_this).F20()), _860_v1, _860_v1)
			var _910_v34 *C1
			_ = _910_v34
			var _nw148 *C1 = New_C1_()
			_ = _nw148
			_nw148.Ctor__((_this).F25(), _900_v24, (_909_v33).Cardinality(), (_this).F26())
			_910_v34 = _nw148
		}
		var _911_v35 _dafny.CodePoint
		_ = _911_v35
		_911_v35 = _dafny.CodePoint('b')
		var _912_v36 _dafny.Map
		_ = _912_v36
		_912_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((!(!((_this).F26()))) && ((_this).F21()), _911_v35)
		_912_v36 = (_912_v36).Update((_this).F21(), _911_v35)
		var _913_v37 _dafny.Sequence
		_ = _913_v37
		_913_v37 = _dafny.SeqOf(p1, p1, (func() _dafny.Array {
			if (_this).F26() {
				return p1
			}
			return p1
		})(), p1, p1)
		_913_v37 = _913_v37
		var _914_v38 _dafny.MultiSet
		_ = _914_v38
		_914_v38 = _dafny.MultiSetOf((_this).F21(), (_this).F26())
		var _source13 D11 = Companion_D11_.Create_DC30_((_this).F21(), _914_v38)
		_ = _source13
		if _source13.Is_DC29() {
			var _915_v39 _dafny.Sequence
			_ = _915_v39
			_915_v39 = _dafny.UnicodeSeqOfUtf8Bytes("yakhfedov")
			var _916_v40 _dafny.Map
			_ = _916_v40
			_916_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), false)
			var _917_v41 _dafny.Sequence
			_ = _917_v41
			_917_v41 = _dafny.SeqOf(_916_v40, _916_v40, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21()))
			var _rhs157 bool = (_this).F21()
			_ = _rhs157
			var _rhs158 bool = (_this).Fm0(_860_v1, _915_v39, _dafny.Companion_Sequence_.Concatenate(_917_v41, _917_v41), globalState)
			_ = _rhs158
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			r0 = _rhs157
			_lhs138.F4 = _rhs158
			(globalState).F17 = (_this).F25()
			var _918_v42 _dafny.Array
			_ = _918_v42
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
			_ = _nw149
			_918_v42 = _nw149
			var _919_v43 _dafny.Array
			_ = _919_v43
			var _nwElement0_37 _dafny.Array = _918_v42
			_ = _nwElement0_37
			var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
			_ = _nw150
			(_nw150).ArraySet1(_nwElement0_37, 0)
			(_nw150).ArraySet1(_918_v42, 1)
			_919_v43 = _nw150
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_919_v43), 0))
			_ = _index151
			(_919_v43).ArraySet1(_918_v42, (_index151).Int())
			var _920_v44 _dafny.Map
			_ = _920_v44
			_920_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfInt64(495))
			var _921_v45 _dafny.Map
			_ = _921_v45
			_921_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_920_v44, _918_v42)
			var _922_v46 _dafny.Sequence
			_ = _922_v46
			_922_v46 = _dafny.SeqOf(_918_v42, (func() _dafny.Array {
				if (_921_v45).Contains(_920_v44) {
					return (_921_v45).Get(_920_v44).(_dafny.Array)
				}
				return _918_v42
			})())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_919_v43), 0))
			_ = _index152
			(_919_v43).ArraySet1((_922_v46).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_915_v39).Cardinality()), (_this).F20()), _dafny.IntOfUint32((_922_v46).Cardinality()))).Uint32()).(_dafny.Array), (_index152).Int())
			var _923_v47 _dafny.MultiSet
			_ = _923_v47
			_923_v47 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_920_v44).Contains((_this).F26()) {
					return (_920_v44).Get((_this).F26()).(_dafny.Int)
				}
				return (_this).F25()
			})())
			if !(_923_v47).Contains((_this).F20()) {
				var _924_v48 _dafny.MultiSet
				_ = _924_v48
				_924_v48 = _dafny.MultiSetOf(_911_v35)
				var _925_v49 D7
				_ = _925_v49
				_925_v49 = Companion_D7_.Create_DC18_(_dafny.SeqOf(_dafny.IntOfUint32((_915_v39).Cardinality())), (_this).F20())
				var _926_v50 *C1
				_ = _926_v50
				var _nw151 *C1 = New_C1_()
				_ = _nw151
				_nw151.Ctor__((_924_v48).Cardinality(), p1, _dafny.IntOfUint32(((_925_v49).Dtor_cf26()).Cardinality()), !_dafny.Companion_Sequence_.Equal(_915_v39, _915_v39))
				_926_v50 = _nw151
				_911_v35 = _911_v35
				var _927_v51 _dafny.Set
				_ = _927_v51
				_927_v51 = _dafny.SetOf((_this).F20(), _dafny.IntOfInt64(438), (_this).F25(), _dafny.IntOfInt64(-926), (_this).F20())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((p1), 0))
				_ = _index153
				(p1).ArraySet1((_this).F20(), (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((p1), 0))
				_ = _index154
				var _rhs159 _dafny.Int = (Companion_Default___.Fm35((_this).F26(), (_this).F21(), (_this).Fm1((_dafny.Zero).Minus(_926_v50.F27), _dafny.IntOfInt64(981), (_this).F26(), globalState), globalState)).Cardinality()
				_ = _rhs159
				var _rhs160 bool = (((_this).F26()) == ((_863_v4).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_863_v4).Cardinality()))).Uint32()).(bool))) || ((_this).F21())
				_ = _rhs160
				var _rhs161 _dafny.Set = (_927_v51).Union(Companion_Default___.Fm25(globalState))
				_ = _rhs161
				var _rhs162 _dafny.Int = _926_v50.F27
				_ = _rhs162
				var _rhs163 bool = ((_this).F25()).Cmp(((_926_v50).Fm1(_926_v50.F27, (_this).F25(), (_this).F21(), globalState)).Plus((_this).F25())) > 0
				_ = _rhs163
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				var _lhs140 _dafny.Array = p1
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((p1), 0))
				_ = _lhs141
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				_lhs139.F5 = _rhs159
				r0 = _rhs160
				_927_v51 = _rhs161
				(_lhs140).ArraySet1(_rhs162, (_lhs141).Int())
				_lhs142.F3 = _rhs163
				r0 = ((_this).F21()) && (!_dafny.Companion_Sequence_.Equal(_915_v39, _915_v39))
				var _928_v52 *C0
				_ = _928_v52
				var _nw152 *C0 = New_C0_()
				_ = _nw152
				_nw152.Ctor__()
				_928_v52 = _nw152
			} else {
				var _929_v53 _dafny.Array
				_ = _929_v53
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw153
				_929_v53 = _nw153
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_929_v53), 0))
				_ = _index155
				(_929_v53).ArraySet1((_this).F26(), (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_929_v53), 0))
				_ = _index156
				(_929_v53).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_915_v39, _915_v39), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_915_v39, _915_v39)).Cardinality()))).Uint32(), _911_v35), _dafny.Companion_Sequence_.Update(_915_v39, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_915_v39).Cardinality()))).Uint32(), _911_v35)), (_index156).Int())
				(globalState).F17 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(-536)).Plus((_this).F20()), (_this).F25())
				var _930_v54 _dafny.Sequence
				_ = _930_v54
				_930_v54 = _dafny.SeqOf(_929_v53)
				_930_v54 = _930_v54
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p1), 0))
				_ = _index157
				(p1).ArraySet1(((_this).F20()).Minus((_this).F20()), (_index157).Int())
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p1), 0))
				_ = _index158
				var _rhs164 bool = (_this).F26()
				_ = _rhs164
				var _rhs165 bool = ((_this).F20()).Cmp((_this).F20()) == 0
				_ = _rhs165
				var _rhs166 _dafny.Set = _859_v0
				_ = _rhs166
				var _rhs167 _dafny.Int = (func() _dafny.Int {
					if (_860_v1).Contains(((_862_v3).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32()).(_dafny.Int)).Times((_this).F20())) {
						return (_860_v1).Get(((_862_v3).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32()).(_dafny.Int)).Times((_this).F20())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(961)
				})()
				_ = _rhs167
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 _dafny.Array = p1
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p1), 0))
				_ = _lhs146
				_lhs143.F3 = _rhs164
				_lhs144.F4 = _rhs165
				_859_v0 = _rhs166
				(_lhs145).ArraySet1(_rhs167, (_lhs146).Int())
				_862_v3 = (func() _dafny.Sequence {
					if (_929_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_929_v53), 0))).Int()).(bool) {
						return _862_v3
					}
					return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_862_v3, (Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32(), (_this).F25()), _862_v3)
				})()
			}
		} else if _source13.Is_DC30() {
			var _931___mcc_h5 bool = _source13.Get_().(D11_DC30).Cf40
			_ = _931___mcc_h5
			var _932___mcc_h6 _dafny.MultiSet = _source13.Get_().(D11_DC30).Cf41
			_ = _932___mcc_h6
			var _933_cf41 _dafny.MultiSet = _932___mcc_h6
			_ = _933_cf41
			var _934_cf40 bool = _931___mcc_h5
			_ = _934_cf40
			var _935_v55 _dafny.Array
			_ = _935_v55
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_27
			var _nw154 _dafny.Array
			_ = _nw154
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw154 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = func(_936_i4 _dafny.Int) bool {
					return (_this).F26()
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw154 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw154).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw154).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_935_v55 = _nw154
			var _937_v56 _dafny.Map
			_ = _937_v56
			_937_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _935_v55)
			var _938_v57 D8
			_ = _938_v57
			_938_v57 = Companion_D8_.Create_DC21_(((_937_v56).Merge(_937_v56)).Cardinality())
			_938_v57 = _938_v57
			var _939_v58 _dafny.MultiSet
			_ = _939_v58
			_939_v58 = _dafny.MultiSetOf((_this).F25())
			_860_v1 = (_860_v1).Update((_939_v58).Cardinality(), (_this).F25())
			if (_this).F21() {
				var _940_v59 _dafny.Map
				_ = _940_v59
				_940_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_911_v35, (_this).F25())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))
				_ = _index159
				(p1).ArraySet1(((_this).F20()).Plus((func() _dafny.Int {
					if (_940_v59).Contains(_911_v35) {
						return (_940_v59).Get(_911_v35).(_dafny.Int)
					}
					return (_this).F20()
				})()), (_index159).Int())
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))
				_ = _index160
				(p1).ArraySet1(((_this).F20()).Times(((_this).F25()).Times((_this).F20())), (_index160).Int())
				var _941_v60 _dafny.Set
				_ = _941_v60
				_941_v60 = _dafny.SetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_this).F20(), (_this).F25(), (_this).F25(), (_dafny.Zero).Minus((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)))
				var _942_v61 _dafny.Sequence
				_ = _942_v61
				_942_v61 = _dafny.UnicodeSeqOfUtf8Bytes("hwv")
				_941_v60 = _dafny.SetOf(((func() _dafny.Int {
					if (_939_v58).Contains(_dafny.IntOfUint32((_942_v61).Cardinality())) {
						return (_939_v58).Multiplicity(_dafny.IntOfUint32((_942_v61).Cardinality()))
					}
					return (_860_v1).Cardinality()
				})()).Minus((_this).F20()), _dafny.IntOfInt64(221), (_this).F20(), _dafny.IntOfInt64(123), (_this).F20())
				var _943_v62 D8
				_ = _943_v62
				_943_v62 = Companion_D8_.Create_DC20_()
				_943_v62 = Companion_Default___.Fm36(globalState)
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))
				_ = _index161
				(p1).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))).Cardinality())), Companion_Default___.SafeModuloInt((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_863_v4).Cardinality()))), (_index161).Int())
				var _944_v63 *C2
				_ = _944_v63
				var _nw155 *C2 = New_C2_()
				_ = _nw155
				_nw155.Ctor__(((_this).F25()).Times((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)), (func() bool {
					if (_861_v2).Contains((_this).Fm1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_this).F20(), true, globalState)) {
						return (_861_v2).Get((_this).Fm1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_this).F20(), true, globalState)).(bool)
					}
					return _934_cf40
				})())
				_944_v63 = _nw155
			} else {
				(globalState).F17 = ((_this).F25()).Minus(Companion_Default___.SafeModuloInt((_this).F25(), (_this).F20()))
				var _945_v64 _dafny.Set
				_ = _945_v64
				_945_v64 = _dafny.SetOf((_this).F20(), (_this).F20(), (_this).F25(), Companion_Default___.SafeModuloInt((_this).F20(), _dafny.IntOfInt64(752)), (_dafny.IntOfInt64(129)).Plus((_this).F25()))
				var _946_v65 _dafny.Array
				_ = _946_v65
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
				_ = _nw156
				_946_v65 = _nw156
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_946_v65), 0))
				_ = _index162
				(_946_v65).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F25())), (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_946_v65), 0))
				_ = _index163
				var _rhs168 _dafny.Set = func() _dafny.Set {
					var _coll29 = _dafny.NewBuilder()
					_ = _coll29
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(632), _dafny.IntOfInt64(-38))); ; {
						_compr_29, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _947_v66 _dafny.Int
						_947_v66 = interface{}(_compr_29).(_dafny.Int)
						if ((_dafny.IntOfInt64(632)).Cmp(_947_v66) <= 0) && ((_947_v66).Cmp(_dafny.IntOfInt64(-38)) < 0) {
							_coll29.Add(Companion_Default___.SafeDivisionInt(_947_v66, _dafny.IntOfUint32((_863_v4).Cardinality())))
						}
					}
					return _coll29.ToSet()
				}()
				_ = _rhs168
				var _rhs169 _dafny.MultiSet = _939_v58
				_ = _rhs169
				var _rhs170 _dafny.Int = _dafny.IntOfInt64(-424)
				_ = _rhs170
				var _lhs147 _dafny.Array = _946_v65
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_946_v65), 0))
				_ = _lhs148
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				_945_v64 = _rhs168
				(_lhs147).ArraySet1(_rhs169, (_lhs148).Int())
				_lhs149.F17 = _rhs170
				var _948_v67 _dafny.Sequence
				_ = _948_v67
				_948_v67 = _dafny.SeqOf(_dafny.SeqOf((_this).F20(), _dafny.IntOfInt64(860)))
				_948_v67 = _948_v67
				var _949_v68 _dafny.Map
				_ = _949_v68
				_949_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_863_v4, (_this).F20())
				_949_v68 = (_949_v68).Update(_dafny.SeqOf((func() bool {
					if (_861_v2).Contains(_dafny.IntOfUint32((_862_v3).Cardinality())) {
						return (_861_v2).Get(_dafny.IntOfUint32((_862_v3).Cardinality())).(bool)
					}
					return (_this).F21()
				})(), (_this).F26()), _dafny.IntOfUint32((_dafny.SeqOf(_934_cf40, (_this).F26(), (_this).F26(), false)).Cardinality()))
				var _950_v69 _dafny.MultiSet
				_ = _950_v69
				_950_v69 = _dafny.MultiSetOf(_863_v4, _dafny.Companion_Sequence_.Concatenate(_863_v4, _863_v4))
				var _951_v70 _dafny.Sequence
				_ = _951_v70
				_951_v70 = _dafny.SeqOf(_863_v4, _863_v4)
				var _952_v71 _dafny.Sequence
				_ = _952_v71
				_952_v71 = _dafny.SeqOf(_863_v4, _863_v4, (_951_v70).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_951_v70).Cardinality()))).Uint32()).(_dafny.Sequence), _863_v4)
				var _953_v72 D12
				_ = _953_v72
				_953_v72 = Companion_D12_.Create_DC32_(_952_v71)
				_950_v69 = _dafny.MultiSetFromSeq((_953_v72).Dtor_cf43())
			}
			_861_v2 = (_861_v2).Update((_this).F20(), !(((_this).F25()).Cmp((_this).F25()) != 0))
		} else if _source13.Is_DC28() {
			var _954___mcc_h7 _dafny.Array = _source13.Get_().(D11_DC28).Cf39
			_ = _954___mcc_h7
			var _955_cf39 _dafny.Array = _954___mcc_h7
			_ = _955_cf39
			var _956_v73 _dafny.Map
			_ = _956_v73
			_956_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F25())
			var _957_v74 *C1
			_ = _957_v74
			var _nw157 *C1 = New_C1_()
			_ = _nw157
			_nw157.Ctor__(Companion_Default___.SafeModuloInt((_this).F25(), (_this).F20()), p1, (_dafny.Zero).Minus((func() _dafny.Int {
				if (_956_v73).Contains((_this).F26()) {
					return (_956_v73).Get((_this).F26()).(_dafny.Int)
				}
				return (_this).F20()
			})()), ((_this).F26()) || ((_this).F26()))
			_957_v74 = _nw157
			_957_v74 = _957_v74
			var _958_v75 _dafny.Array
			_ = _958_v75
			_958_v75 = (_957_v74).F28()
			var _959_v76 _dafny.Map
			_ = _959_v76
			_959_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_863_v4).Cardinality()), _958_v75)
			var _960_v77 _dafny.Sequence
			_ = _960_v77
			_960_v77 = _dafny.SeqOf(_959_v76)
			(_957_v74).F27 = (Companion_Default___.SafeModuloInt(_957_v74.F27, (_dafny.Zero).Minus(((_960_v77).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_960_v77).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()))).Plus(Companion_Default___.SafeDivisionInt(_957_v74.F27, _dafny.IntOfInt64(-289)))
			var _961_v78 _dafny.Array
			_ = _961_v78
			var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(16))
			_ = _nw158
			_961_v78 = _nw158
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_961_v78), 0))
			_ = _index164
			(_961_v78).ArraySet1(_914_v38, (_index164).Int())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_961_v78), 0))
			_ = _index165
			(_961_v78).ArraySet1(_dafny.MultiSetOf(!((_this).F21())), (_index165).Int())
		} else {
			var _962___mcc_h8 D11 = _source13.Get_().(D11_DC31).Cf42
			_ = _962___mcc_h8
			var _963_cf42 D11 = _962___mcc_h8
			_ = _963_cf42
			var _964_v79 _dafny.Array
			_ = _964_v79
			var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw159
			_964_v79 = _nw159
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_964_v79), 0))
			_ = _index166
			(_964_v79).ArraySet1(_863_v4, (_index166).Int())
			var _965_v80 _dafny.Set
			_ = _965_v80
			_965_v80 = _dafny.SetOf((_this).F25())
			var _966_v81 _dafny.Array
			_ = _966_v81
			var _nwElement0_38 _dafny.CodePoint = _911_v35
			_ = _nwElement0_38
			var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(12))
			_ = _nw160
			(_nw160).ArraySet1CodePoint(_nwElement0_38, 0)
			(_nw160).ArraySet1CodePoint(_911_v35, 1)
			(_nw160).ArraySet1CodePoint(Companion_Default___.Fm17((_this).F26(), globalState), 2)
			(_nw160).ArraySet1CodePoint(_911_v35, 3)
			(_nw160).ArraySet1CodePoint(_911_v35, 4)
			(_nw160).ArraySet1CodePoint(_911_v35, 5)
			(_nw160).ArraySet1CodePoint(_911_v35, 6)
			(_nw160).ArraySet1CodePoint(_911_v35, 7)
			(_nw160).ArraySet1CodePoint(_911_v35, 8)
			(_nw160).ArraySet1CodePoint(_911_v35, 9)
			(_nw160).ArraySet1CodePoint(_911_v35, 10)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('f'), 11)
			_966_v81 = _nw160
			var _967_v82 _dafny.Map
			_ = _967_v82
			_967_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), (_this).F25())
			var _968_v83 _dafny.Map
			_ = _968_v83
			_968_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_966_v81, _967_v82)
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_964_v79), 0))
			_ = _index167
			var _rhs171 _dafny.Sequence = _dafny.SeqOf((_this).F26(), (_this).F26(), false, (_965_v80).IsProperSubsetOf(_965_v80), (_this).F26())
			_ = _rhs171
			var _rhs172 bool = ((_862_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_862_v3).Cardinality()), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(((_this).F20()).Times((_this).F20())) > 0
			_ = _rhs172
			var _rhs173 _dafny.Int = (func() _dafny.Int {
				if (_this).F21() {
					return (_968_v83).Cardinality()
				}
				return _dafny.IntOfInt64(153)
			})()
			_ = _rhs173
			var _rhs174 _dafny.Int = (_this).F25()
			_ = _rhs174
			var _rhs175 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F25(), (_dafny.Zero).Minus((_this).F25()))
			_ = _rhs175
			var _lhs150 _dafny.Array = _964_v79
			_ = _lhs150
			var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_964_v79), 0))
			_ = _lhs151
			var _lhs152 *GlobalState = globalState
			_ = _lhs152
			var _lhs153 *GlobalState = globalState
			_ = _lhs153
			var _lhs154 *GlobalState = globalState
			_ = _lhs154
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			(_lhs150).ArraySet1(_rhs171, (_lhs151).Int())
			_lhs152.F14 = _rhs172
			_lhs153.F17 = _rhs173
			_lhs154.F5 = _rhs174
			_lhs155.F17 = _rhs175
			var _969_v85 _dafny.Sequence
			_ = _969_v85
			_969_v85 = _dafny.UnicodeSeqOfUtf8Bytes("dr")
			if !(!((_this).Fm0((func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(587), _dafny.IntOfInt64(-160))); ; {
					_compr_30, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _970_v84 _dafny.Int
					_970_v84 = interface{}(_compr_30).(_dafny.Int)
					if ((_dafny.IntOfInt64(587)).Cmp(_970_v84) <= 0) && ((_970_v84).Cmp(_dafny.IntOfInt64(-160)) < 0) {
						_coll30.Add((_970_v84).Times((_this).F20()), (_this).F25())
					}
				}
				return _coll30.ToMap()
			}()).Merge(_860_v1), _969_v85, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-466))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_971_i5 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
			})), globalState))) {
				(globalState).F17 = _dafny.IntOfInt64(88)
				var _972_v86 D6
				_ = _972_v86
				_972_v86 = Companion_D6_.Create_DC14_((_this).Fm1((_this).F20(), (_this).F20(), (_this).F21(), globalState))
				var _973_v87 D2
				_ = _973_v87
				_973_v87 = Companion_D2_.Create_DC3_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_969_v85, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nw")).Cardinality()), _dafny.IntOfUint32((_969_v85).Cardinality()))).Uint32(), _911_v35)).Cardinality()), (_this).F26(), (_this).F26(), (_this).F20(), (_this).F21())
				var _974_v88 _dafny.Map
				_ = _974_v88
				_974_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F21())
				var _975_v89 _dafny.Sequence
				_ = _975_v89
				_975_v89 = _dafny.SeqOf(_974_v88)
				var _976_v90 _dafny.Map
				_ = _976_v90
				_976_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_863_v4).Cardinality()), (_this).F25()), _dafny.UnicodeSeqOfUtf8Bytes("rwox"), _975_v89, globalState), _dafny.IntOfInt64(798))
				var _977_v91 _dafny.Array
				_ = _977_v91
				var _nwElement0_39 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_863_v4).Cardinality()))
				_ = _nwElement0_39
				var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(23))
				_ = _nw161
				(_nw161).ArraySet1(_nwElement0_39, 0)
				(_nw161).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F25()), _dafny.IntOfInt64(-804)), 1)
				(_nw161).ArraySet1((_972_v86).Dtor_cf22(), 2)
				(_nw161).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()), 3)
				(_nw161).ArraySet1((_this).F25(), 4)
				(_nw161).ArraySet1((_this).F20(), 5)
				(_nw161).ArraySet1((_this).F25(), 6)
				(_nw161).ArraySet1((func() _dafny.Int {
					if (_this).F21() {
						return (_this).F25()
					}
					return (_this).F20()
				})(), 7)
				(_nw161).ArraySet1((_this).Fm1(_dafny.IntOfUint32((_969_v85).Cardinality()), (_this).F25(), (_this).F26(), globalState), 8)
				(_nw161).ArraySet1((_this).F20(), 9)
				(_nw161).ArraySet1((_this).F25(), 10)
				(_nw161).ArraySet1(Companion_Default___.SafeModuloInt((_this).F25(), (_this).F20()), 11)
				(_nw161).ArraySet1((_this).F25(), 12)
				(_nw161).ArraySet1((_973_v87).Dtor_cf6(), 13)
				(_nw161).ArraySet1((_this).F20(), 14)
				(_nw161).ArraySet1((_this).F25(), 15)
				(_nw161).ArraySet1((func() _dafny.Int {
					if (_976_v90).Contains((_this).F26()) {
						return (_976_v90).Get((_this).F26()).(_dafny.Int)
					}
					return (_this).F20()
				})(), 16)
				(_nw161).ArraySet1((_this).F20(), 17)
				(_nw161).ArraySet1((_974_v88).Cardinality(), 18)
				(_nw161).ArraySet1((_dafny.Zero).Minus((_this).F25()), 19)
				(_nw161).ArraySet1((_this).F25(), 20)
				(_nw161).ArraySet1((func() _dafny.Int {
					if (_this).F26() {
						return (_859_v0).Cardinality()
					}
					return (_this).F20()
				})(), 21)
				(_nw161).ArraySet1((_this).F25(), 22)
				_977_v91 = _nw161
				var _rhs176 bool = !((_this).F21())
				_ = _rhs176
				var _rhs177 _dafny.Array = p1
				_ = _rhs177
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				_lhs156.F18 = _rhs176
				_977_v91 = _rhs177
				(globalState).F3 = ((_914_v38).Union(Companion_Default___.Fm23(globalState))).IsDisjointFrom((_914_v38).Intersection(_dafny.MultiSetFromSeq(Companion_Default___.Fm21(globalState))))
				var _nwElement0_40 _dafny.Int = (_this).F20()
				_ = _nwElement0_40
				var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(2))
				_ = _nw162
				(_nw162).ArraySet1(_nwElement0_40, 0)
				(_nw162).ArraySet1((_this).F20(), 1)
				_977_v91 = _nw162
				var _978_v92 *C0
				_ = _978_v92
				var _nw163 *C0 = New_C0_()
				_ = _nw163
				_nw163.Ctor__()
				_978_v92 = _nw163
			} else {
				(globalState).F5 = (_this).F25()
				_969_v85 = _dafny.UnicodeSeqOfUtf8Bytes("tbh")
				var _979_v93 _dafny.Array
				_ = _979_v93
				_979_v93 = p1
				_979_v93 = _979_v93
				(globalState).F5 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_862_v3).Cardinality()))
				var _980_v94 T0
				_ = _980_v94
				var _nw164 *C2 = New_C2_()
				_ = _nw164
				_nw164.Ctor__(_dafny.IntOfInt64(340), (_this).F21())
				_980_v94 = _nw164
				var _981_v95 _dafny.Map
				_ = _981_v95
				_981_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T0 {
					if (_this).F26() {
						return _980_v94
					}
					return _980_v94
				})(), (_dafny.Zero).Minus(((_this).Fm1((func() _dafny.Int {
					if (_860_v1).Contains((_this).F20()) {
						return (_860_v1).Get((_this).F20()).(_dafny.Int)
					}
					return (_this).F20()
				})(), (_this).F25(), true, globalState)).Times((_980_v94).F20())))
				_981_v95 = (_981_v95).Update((func() T0 {
					if (_this).F21() {
						return _980_v94
					}
					return _980_v94
				})(), (_980_v94).F20())
			}
			var _982_v96 D8
			_ = _982_v96
			_982_v96 = Companion_D8_.Create_DC20_()
			_982_v96 = _982_v96
			var _983_v97 _dafny.MultiSet
			_ = _983_v97
			_983_v97 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F25())))
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((p1), 0))
			_ = _index168
			(p1).ArraySet1(((func() _dafny.MultiSet {
				if (_this).F21() {
					return (_983_v97).Update((_this).F20(), Companion_Default___.Abs((_861_v2).Cardinality()))
				}
				return _983_v97
			})()).Cardinality(), (_index168).Int())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((p1), 0))
			_ = _index169
			(p1).ArraySet1((_dafny.IntOfInt64(-521)).Times((_this).F25()), (_index169).Int())
		}
		var _984_v98 _dafny.Array
		_ = _984_v98
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_28
		var _nw165 _dafny.Array
		_ = _nw165
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw165 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) D8 = func(_985_i6 _dafny.Int) D8 {
				return Companion_D8_.Create_DC20_()
			}
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw165 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw165).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw165).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_984_v98 = _nw165
		var _986_v99 _dafny.Array
		_ = _986_v99
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_29
		var _nw166 _dafny.Array
		_ = _nw166
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw166 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) _dafny.CodePoint = (func(_987_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_988_i7 _dafny.Int) _dafny.CodePoint {
					return (Companion_D13_.Create_DC35_(_987_v35)).Dtor_cf45()
				}
			})(_911_v35)
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw166 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw166).ArraySet1CodePoint(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw166).ArraySet1CodePoint(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_986_v99 = _nw166
		var _989_v100 _dafny.Sequence
		_ = _989_v100
		_989_v100 = _dafny.SeqOf(_986_v99)
		var _990_v101 _dafny.Map
		_ = _990_v101
		_990_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F20())
		var _991_v102 _dafny.Map
		_ = _991_v102
		_991_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _862_v3)
		var _992_v103 _dafny.Sequence
		_ = _992_v103
		_992_v103 = _dafny.UnicodeSeqOfUtf8Bytes("cntlqc")
		var _rhs178 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_989_v100, _dafny.SeqOf(_986_v99, _986_v99, _986_v99, _986_v99, _986_v99)), _989_v100)
		_ = _rhs178
		var _rhs179 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_990_v101).Contains(p1) {
				return (_990_v101).Get(p1).(_dafny.Int)
			}
			return _dafny.IntOfInt64(862)
		})(), (_this).F20())
		_ = _rhs179
		var _rhs180 _dafny.Array = _984_v98
		_ = _rhs180
		var _rhs181 _dafny.Int = ((Companion_D8_.Create_DC21_((_this).F25())).Dtor_cf29()).Plus(Companion_Default___.SafeModuloInt((_991_v102).Cardinality(), (_dafny.Zero).Minus((_this).F20())))
		_ = _rhs181
		var _rhs182 bool = !((_this).F26()) || ((_dafny.IntOfUint32((_992_v103).Cardinality())).Cmp((_this).F20()) == 0)
		_ = _rhs182
		var _lhs157 *GlobalState = globalState
		_ = _lhs157
		var _lhs158 *GlobalState = globalState
		_ = _lhs158
		var _lhs159 *GlobalState = globalState
		_ = _lhs159
		_lhs157.F14 = _rhs178
		_lhs158.F5 = _rhs179
		_984_v98 = _rhs180
		_lhs159.F5 = _rhs181
		r0 = _rhs182
		var _993_v104 D8
		_ = _993_v104
		_993_v104 = Companion_D8_.Create_DC19_(_860_v1)
		var _994_v105 _dafny.Map
		_ = _994_v105
		_994_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F21())
		var _995_v106 _dafny.Sequence
		_ = _995_v106
		_995_v106 = _dafny.SeqOf(_994_v105)
		r0 = (_this).Fm0((_993_v104).Dtor_cf28(), _992_v103, _995_v106, globalState)
		return r0
	}
}
func (_this *C3) F25() _dafny.Int {
	{
		return _this._f25
	}
}
func (_this *C3) F26() bool {
	{
		return _this._f26
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f20 _dafny.Int
	_f21 bool
	_f24 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
	_this._f24 = false
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

func (_this *C4) F20() _dafny.Int {
	return _this._f20
}
func (_this *C4) F21() bool {
	return _this._f21
}
func (_this *C4) Ctor__(f24 bool, f20 _dafny.Int, f21 bool) {
	{
		(_this)._f24 = f24
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C4) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C4) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.SetOf((_this).F24())).Cardinality()).Times((_this).F20())).Times((Companion_D4_.Create_DC11_((_this).F20(), (_this).F21(), (_this).F24())).Dtor_cf17())
	}
}
func (_this *C4) Fm13(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F21())).Cardinality()).Plus((_this).F20())).Cmp((_dafny.Zero).Minus((_this).F20())) == 0
	}
}
func (_this *C4) Fm14(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F20(), (_this).F20()))
	}
}
func (_this *C4) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _996_v0 _dafny.Array
		_ = _996_v0
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
		_ = _nw167
		_996_v0 = _nw167
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_996_v0), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _997_i0 _dafny.Int
			_997_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_997_i0).Sign() != -1) && ((_997_i0).Cmp(_dafny.ArrayLen((_996_v0), 0)) < 0)) {
				(_996_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F21(), (_this).F21()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24(), p1), _dafny.SeqOf((_this).F24()))), (_997_i0).Int())
			}
		}
		var _998_v1 D4
		_ = _998_v1
		_998_v1 = Companion_D4_.Create_DC10_(_dafny.MultiSetOf((_this).F21()), (_this).F24())
		(globalState).F5 = _dafny.IntOfUint32((func(_source14 D4) _dafny.Sequence {
			if _source14.Is_DC10() {
				var _999___mcc_h0 _dafny.MultiSet = _source14.Get_().(D4_DC10).Cf15
				_ = _999___mcc_h0
				var _1000___mcc_h1 bool = _source14.Get_().(D4_DC10).Cf16
				_ = _1000___mcc_h1
				var _1001_cf16 bool = _1000___mcc_h1
				_ = _1001_cf16
				var _1002_cf15 _dafny.MultiSet = _999___mcc_h0
				_ = _1002_cf15
				return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("lqfaswbj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-778))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}(func(_1003_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_1004_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				})))
			} else if _source14.Is_DC11() {
				var _1005___mcc_h2 _dafny.Int = _source14.Get_().(D4_DC11).Cf17
				_ = _1005___mcc_h2
				var _1006___mcc_h3 bool = _source14.Get_().(D4_DC11).Cf18
				_ = _1006___mcc_h3
				var _1007___mcc_h4 bool = _source14.Get_().(D4_DC11).Cf19
				_ = _1007___mcc_h4
				var _1008_cf19 bool = _1007___mcc_h4
				_ = _1008_cf19
				var _1009_cf18 bool = _1006___mcc_h3
				_ = _1009_cf18
				var _1010_cf17 _dafny.Int = _1005___mcc_h2
				_ = _1010_cf17
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}(func(_1011_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("tlweiv")
				}))
			} else {
				var _1012___mcc_h5 _dafny.Array = _source14.Get_().(D4_DC9).Cf14
				_ = _1012___mcc_h5
				var _1013_cf14 _dafny.Array = _1012___mcc_h5
				_ = _1013_cf14
				return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cnq"))
			}
		}(_998_v1)).Cardinality())
		var _1014_v2 D10
		_ = _1014_v2
		_1014_v2 = Companion_D10_.Create_DC26_(p0)
		var _1015_v3 _dafny.Set
		_ = _1015_v3
		_1015_v3 = _dafny.SetOf(_1014_v2)
		var _1016_v4 _dafny.Sequence
		_ = _1016_v4
		_1016_v4 = _dafny.SeqOf(!(false))
		var _1017_v5 _dafny.Array
		_ = _1017_v5
		var _nwElement0_41 _dafny.Int = (_this).F20()
		_ = _nwElement0_41
		var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(9))
		_ = _nw168
		(_nw168).ArraySet1(_nwElement0_41, 0)
		(_nw168).ArraySet1(_dafny.IntOfUint32((_1016_v4).Cardinality()), 1)
		(_nw168).ArraySet1((_this).F20(), 2)
		(_nw168).ArraySet1((_this).F20(), 3)
		(_nw168).ArraySet1((_this).F20(), 4)
		(_nw168).ArraySet1(p0, 5)
		(_nw168).ArraySet1((_this).F20(), 6)
		(_nw168).ArraySet1(p0, 7)
		(_nw168).ArraySet1((_dafny.SetOf(p0, (_this).F20())).Cardinality(), 8)
		_1017_v5 = _nw168
		var _1018_v6 T0
		_ = _1018_v6
		var _nw169 *C1 = New_C1_()
		_ = _nw169
		_nw169.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(749), (_1015_v3).Cardinality()), _1017_v5, p0, (func() bool {
			if p1 {
				return (_this).F24()
			}
			return (_this).F24()
		})())
		_1018_v6 = _nw169
		_1018_v6 = _1018_v6
		if false {
			var _1019_v7 _dafny.Sequence
			_ = _1019_v7
			_1019_v7 = _dafny.UnicodeSeqOfUtf8Bytes("qxdrfss")
			var _1020_v8 D6
			_ = _1020_v8
			_1020_v8 = Companion_D6_.Create_DC13_(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm26((_this).F21(), globalState), _1019_v7))
			var _source15 D6 = _1020_v8
			_ = _source15
			if _source15.Is_DC14() {
				var _1021___mcc_h6 _dafny.Int = _source15.Get_().(D6_DC14).Cf22
				_ = _1021___mcc_h6
				var _1022_cf22 _dafny.Int = _1021___mcc_h6
				_ = _1022_cf22
				var _1023_v9 *C1
				_ = _1023_v9
				var _nw170 *C1 = New_C1_()
				_ = _nw170
				_nw170.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F20()), (_1018_v6).F20())), _1017_v5, _dafny.IntOfUint32((_1019_v7).Cardinality()), (_1018_v6).F21())
				_1023_v9 = _nw170
				(globalState).F5 = (_dafny.Zero).Minus((_1018_v6).F20())
				var _1024_v10 _dafny.Array
				_ = _1024_v10
				var _1025_v11 _dafny.Int
				_ = _1025_v11
				var _out14 _dafny.Array
				_ = _out14
				var _out15 _dafny.Int
				_ = _out15
				_out14, _out15 = (_this).M1(_1022_cf22, (_1018_v6).F21(), _1018_v6, globalState)
				_1024_v10 = _out14
				_1025_v11 = _out15
				(_1023_v9).F27 = (_1018_v6).F20()
			} else if _source15.Is_DC15() {
				var _1026___mcc_h7 _dafny.Array = _source15.Get_().(D6_DC15).Cf23
				_ = _1026___mcc_h7
				var _1027_cf23 _dafny.Array = _1026___mcc_h7
				_ = _1027_cf23
				var _1028_v12 _dafny.MultiSet
				_ = _1028_v12
				_1028_v12 = _dafny.MultiSetOf((_this).F20())
				var _1029_v13 _dafny.Sequence
				_ = _1029_v13
				_1029_v13 = _dafny.SeqOf((_1018_v6).F20(), _dafny.IntOfUint32((_1019_v7).Cardinality()), _dafny.IntOfUint32((_1019_v7).Cardinality()), p0, (_1018_v6).F20())
				var _1030_v14 _dafny.MultiSet
				_ = _1030_v14
				_1030_v14 = _dafny.MultiSetOf(p1)
				var _1031_v15 _dafny.Sequence
				_ = _1031_v15
				_1031_v15 = _dafny.SeqOf((_1018_v6).F20(), (_this).F20(), (_this).Fm14((_1018_v6).F21(), (_1018_v6).F20(), (func() _dafny.Int {
					if (_1028_v12).Contains((_1029_v13).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_1028_v12).Contains((_1030_v14).Cardinality()) {
							return (_1028_v12).Multiplicity((_1030_v14).Cardinality())
						}
						return (_1018_v6).F20()
					})(), _dafny.IntOfUint32((_1029_v13).Cardinality()))).Uint32()).(_dafny.Int)) {
						return (_1028_v12).Multiplicity((_1029_v13).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_1028_v12).Contains((_1030_v14).Cardinality()) {
								return (_1028_v12).Multiplicity((_1030_v14).Cardinality())
							}
							return (_1018_v6).F20()
						})(), _dafny.IntOfUint32((_1029_v13).Cardinality()))).Uint32()).(_dafny.Int))
					}
					return (_this).F20()
				})(), p1, globalState))
				_1031_v15 = _1031_v15
				var _1032_v16 *C0
				_ = _1032_v16
				var _nw171 *C0 = New_C0_()
				_ = _nw171
				_nw171.Ctor__()
				_1032_v16 = _nw171
				(globalState).F3 = !((_this).F21()) || (((_1018_v6).F21()) && ((Companion_D13_.Create_DC36_(p0, !(true), (_this).F24())).Dtor_cf48()))
				var _1033_v17 *C2
				_ = _1033_v17
				var _nw172 *C2 = New_C2_()
				_ = _nw172
				_nw172.Ctor__(Companion_Default___.SafeModuloInt((_1018_v6).F20(), (_1018_v6).F20()), (p0).Cmp(_dafny.IntOfUint32((_1019_v7).Cardinality())) != 0)
				_1033_v17 = _nw172
			} else {
				var _1034___mcc_h8 _dafny.Sequence = _source15.Get_().(D6_DC13).Cf21
				_ = _1034___mcc_h8
				var _1035_cf21 _dafny.Sequence = _1034___mcc_h8
				_ = _1035_cf21
				var _1036_v18 _dafny.Array
				_ = _1036_v18
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_30
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = func(_1037_i4 _dafny.Int) bool {
						return (_this).F21()
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw173 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw173).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw173).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1036_v18 = _nw173
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1036_v18), 0))
				_ = _index170
				(_1036_v18).ArraySet1(false, (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1036_v18), 0))
				_ = _index171
				var _rhs183 bool = _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}(func(_1038_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})), _1035_cf21)
				_ = _rhs183
				var _rhs184 bool = !((_this).F24())
				_ = _rhs184
				var _lhs160 _dafny.Array = _1036_v18
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1036_v18), 0))
				_ = _lhs161
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				(_lhs160).ArraySet1(_rhs183, (_lhs161).Int())
				_lhs162.F14 = _rhs184
				_1017_v5 = _1017_v5
				var _1039_v19 _dafny.Map
				_ = _1039_v19
				_1039_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1018_v6).F21()) == ((_this).F24()), p0)
				_1039_v19 = (_1039_v19).Update((_1018_v6).F21(), p0)
				var _1040_v20 _dafny.Sequence
				_ = _1040_v20
				_1040_v20 = _dafny.SeqOf(p0, _dafny.IntOfInt64(478), _dafny.IntOfInt64(-570))
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1036_v18), 0))
				_ = _index172
				var _rhs185 bool = (_1018_v6).F21()
				_ = _rhs185
				var _rhs186 _dafny.Int = _dafny.IntOfUint32((_1040_v20).Cardinality())
				_ = _rhs186
				var _rhs187 _dafny.Int = (p0).Plus(_dafny.IntOfUint32((_1040_v20).Cardinality()))
				_ = _rhs187
				var _rhs188 bool = (_this).F21()
				_ = _rhs188
				var _rhs189 bool = (_1018_v6).F21()
				_ = _rhs189
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				var _lhs166 _dafny.Array = _1036_v18
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1036_v18), 0))
				_ = _lhs167
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				_lhs163.F18 = _rhs185
				_lhs164.F5 = _rhs186
				_lhs165.F5 = _rhs187
				(_lhs166).ArraySet1(_rhs188, (_lhs167).Int())
				_lhs168.F14 = _rhs189
			}
			var _1041_v21 _dafny.Array
			_ = _1041_v21
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_31
			var _nw174 _dafny.Array
			_ = _nw174
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw174 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1042_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1043_i6 _dafny.Int) _dafny.Sequence {
						return _1042_v7
					}
				})(_1019_v7)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw174 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw174).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw174).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1041_v21 = _nw174
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_1041_v21), 0))
			_ = _index173
			(_1041_v21).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}(func(_1044_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), (_index173).Int())
			var _1045_v22 _dafny.Sequence
			_ = _1045_v22
			_1045_v22 = _dafny.SeqOf(p0, p0, (_1018_v6).F20())
			var _1046_v23 _dafny.Sequence
			_ = _1046_v23
			_1046_v23 = _dafny.SeqOf(_1045_v22, _1045_v22, _1045_v22)
			var _1047_v24 D9
			_ = _1047_v24
			_1047_v24 = Companion_D9_.Create_DC22_(_1016_v4)
			var _1048_v25 D9
			_ = _1048_v25
			_1048_v25 = Companion_D9_.Create_DC24_(_1047_v24)
			var _1049_v26 D9
			_ = _1049_v26
			_1049_v26 = Companion_D9_.Create_DC24_(_1048_v25)
			var _1050_v27 _dafny.CodePoint
			_ = _1050_v27
			_1050_v27 = _dafny.CodePoint('j')
			var _1051_v28 _dafny.Map
			_ = _1051_v28
			_1051_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F20(), (_1018_v6).F20())
			var _1052_v29 _dafny.Map
			_ = _1052_v29
			_1052_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21())
			var _1053_v30 _dafny.Sequence
			_ = _1053_v30
			_1053_v30 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1018_v6).F21()), _1052_v29)
			var _1054_v31 *C3
			_ = _1054_v31
			var _nw175 *C3 = New_C3_()
			_ = _nw175
			_nw175.Ctor__(_dafny.IntOfInt64(382), (_this).Fm0(_1051_v28, _1019_v7, _1053_v30, globalState), (_1018_v6).F20(), (_this).Fm13((_1018_v6).F20(), globalState))
			_1054_v31 = _nw175
			var _1055_v32 _dafny.Map
			_ = _1055_v32
			_1055_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v31, _1046_v23)
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_1041_v21), 0))
			_ = _index174
			var _rhs190 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1019_v7, (Companion_Default___.SafeIndex((_1018_v6).F20(), _dafny.IntOfUint32((_1019_v7).Cardinality()))).Uint32(), _1050_v27)
			_ = _rhs190
			var _rhs191 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(657), p0)
			_ = _rhs191
			var _rhs192 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1046_v23, (func() _dafny.Sequence {
				if (_1055_v32).Contains(_1054_v31) {
					return (_1055_v32).Get(_1054_v31).(_dafny.Sequence)
				}
				return _1046_v23
			})()), _1046_v23)
			_ = _rhs192
			var _rhs193 D9 = _1049_v26
			_ = _rhs193
			var _rhs194 _dafny.Array = _1041_v21
			_ = _rhs194
			var _lhs169 _dafny.Array = _1041_v21
			_ = _lhs169
			var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_1041_v21), 0))
			_ = _lhs170
			var _lhs171 *GlobalState = globalState
			_ = _lhs171
			(_lhs169).ArraySet1(_rhs190, (_lhs170).Int())
			_lhs171.F5 = _rhs191
			_1046_v23 = _rhs192
			_1049_v26 = _rhs193
			_1041_v21 = _rhs194
			var _1056_v33 _dafny.Array
			_ = _1056_v33
			var _nw176 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw176
			_1056_v33 = _nw176
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1056_v33), 0))
			_ = _index175
			(_1056_v33).ArraySet1((func() _dafny.Set {
				var _coll31 = _dafny.NewBuilder()
				_ = _coll31
				for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(841), _dafny.IntOfInt64(77))); ; {
					_compr_31, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _1057_v34 _dafny.Int
					_1057_v34 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(841)).Cmp(_1057_v34) <= 0) && ((_1057_v34).Cmp(_dafny.IntOfInt64(77)) < 0) {
						_coll31.Add((_1057_v34).Times((_1018_v6).F20()))
					}
				}
				return _coll31.ToSet()
			}()).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(987))), (_index175).Int())
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1056_v33), 0))
			_ = _index176
			(_1056_v33).ArraySet1((_dafny.IntOfUint32((_1019_v7).Cardinality())).Cmp(_dafny.IntOfUint32((_1019_v7).Cardinality())) > 0, (_index176).Int())
			var _1058_v35 _dafny.Map
			_ = _1058_v35
			_1058_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm34(_1050_v27, (_this).F21(), (_1018_v6).F20(), p0, globalState)).Dtor_cf12(), _dafny.IntOfInt64(-844))
			var _1059_v36 _dafny.MultiSet
			_ = _1059_v36
			_1059_v36 = _dafny.MultiSetOf(p0)
			var _1060_v37 _dafny.Sequence
			_ = _1060_v37
			_1060_v37 = _dafny.SeqOf(_1058_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1056_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1056_v33), 0))).Int()).(bool), (_1059_v36).Cardinality())).Merge(_1058_v35), (func() _dafny.Map {
				if (_this).F21() {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1056_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1056_v33), 0))).Int()).(bool), p0)
				}
				return _1058_v35
			})(), (_1058_v35).Merge(Companion_Default___.Fm27((_1018_v6).F20(), globalState)), _1058_v35)
			_1060_v37 = _1060_v37
			var _1061_v38 *C1
			_ = _1061_v38
			var _nw177 *C1 = New_C1_()
			_ = _nw177
			_nw177.Ctor__(p0, _1017_v5, _dafny.IntOfInt64(342), (_1054_v31).F26())
			_1061_v38 = _nw177
			var _1062_v39 _dafny.Map
			_ = _1062_v39
			_1062_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F20(), _1061_v38)
			var _1063_v40 *C1
			_ = _1063_v40
			var _nw178 *C1 = New_C1_()
			_ = _nw178
			_nw178.Ctor__((_dafny.Zero).Minus(p0), _1017_v5, (_1062_v39).Cardinality(), (_1061_v38.F27).Cmp((_1018_v6).F20()) != 0)
			_1063_v40 = _nw178
		} else {
			(globalState).F18 = (_1018_v6).F21()
			var _rhs195 _dafny.Int = p0
			_ = _rhs195
			var _rhs196 _dafny.Int = (_1018_v6).F20()
			_ = _rhs196
			var _rhs197 bool = (p1) && (((_1018_v6).F20()).Cmp(p0) <= 0)
			_ = _rhs197
			var _lhs172 *GlobalState = globalState
			_ = _lhs172
			var _lhs173 *GlobalState = globalState
			_ = _lhs173
			var _lhs174 *GlobalState = globalState
			_ = _lhs174
			_lhs172.F17 = _rhs195
			_lhs173.F17 = _rhs196
			_lhs174.F14 = _rhs197
			var _1064_v41 _dafny.Map
			_ = _1064_v41
			_1064_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_1018_v6).F21()))
			var _1065_v42 _dafny.Map
			_ = _1065_v42
			_1065_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if (_1064_v41).Contains((_this).F20()) {
					return (_1064_v41).Get((_this).F20()).(bool)
				}
				return (_this).F21()
			})()), _dafny.IntOfInt64(-741))
			_1065_v42 = (_1065_v42).Update((func() bool {
				if (_1018_v6).F21() {
					return p1
				}
				return (_1018_v6).F21()
			})(), (_this).F20())
			_1017_v5 = _1017_v5
			(globalState).F17 = Companion_Default___.SafeModuloInt((_this).F20(), (_1018_v6).F20())
		}
		var _1066_v43 _dafny.Set
		_ = _1066_v43
		_1066_v43 = _dafny.SetOf((_this).F21(), p1)
		var _1067_v44 _dafny.Set
		_ = _1067_v44
		_1067_v44 = _dafny.SetOf((_1066_v43).Cardinality())
		var _1068_v45 _dafny.MultiSet
		_ = _1068_v45
		_1068_v45 = _dafny.MultiSetOf((_1067_v44).Cardinality())
		if (_1068_v45).IsProperSubsetOf((_1068_v45).Difference(_1068_v45)) {
			var _1069_v46 D13
			_ = _1069_v46
			_1069_v46 = Companion_D13_.Create_DC36_((_this).F20(), !(!(((_1018_v6).F21()) && ((_1018_v6).F21()))), (_1016_v4).Select((Companion_Default___.SafeIndex((_1018_v6).F20(), _dafny.IntOfUint32((_1016_v4).Cardinality()))).Uint32()).(bool))
			var _1070_v47 _dafny.Map
			_ = _1070_v47
			_1070_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
			var _1071_v48 _dafny.Sequence
			_ = _1071_v48
			_1071_v48 = _dafny.UnicodeSeqOfUtf8Bytes("hadgdwudf")
			var _rhs198 _dafny.Int = (Companion_Default___.SafeModuloInt((_1018_v6).F20(), (_this).F20())).Plus((func() _dafny.Int {
				if (_1070_v47).Contains(p0) {
					return (_1070_v47).Get(p0).(_dafny.Int)
				}
				return (_this).F20()
			})())
			_ = _rhs198
			var _rhs199 _dafny.Int = _dafny.IntOfUint32((_1071_v48).Cardinality())
			_ = _rhs199
			var _rhs200 D13 = _1069_v46
			_ = _rhs200
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			_lhs175.F5 = _rhs198
			_lhs176.F5 = _rhs199
			_1069_v46 = _rhs200
			var _1072_v49 _dafny.Map
			_ = _1072_v49
			_1072_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1071_v48).Cardinality()), _1018_v6)
			if ((_1018_v6).F20()).Cmp((_1072_v49).Cardinality()) > 0 {
				_1017_v5 = _1017_v5
				(globalState).F5 = (_1018_v6).F20()
				var _1073_v50 _dafny.MultiSet
				_ = _1073_v50
				_1073_v50 = _dafny.MultiSetOf(p1)
				var _1074_v51 _dafny.Sequence
				_ = _1074_v51
				_1074_v51 = _dafny.SeqOf((func() _dafny.Int {
					if (_1073_v50).Contains((_this).F21()) {
						return (_1073_v50).Multiplicity((_this).F21())
					}
					return (_1018_v6).Fm1((_1018_v6).F20(), (_dafny.SetOf((_this).F20(), p0, _dafny.IntOfInt64(809))).Cardinality(), (_1018_v6).F21(), globalState)
				})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}(func(_1075_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))).Cardinality()))
				var _1076_v52 _dafny.CodePoint
				_ = _1076_v52
				_1076_v52 = _dafny.CodePoint('v')
				(globalState).F18 = _dafny.Companion_Sequence_.IsPrefixOf(_1071_v48, (func() _dafny.Sequence {
					if (_1018_v6).F21() {
						return _1071_v48
					}
					return _dafny.Companion_Sequence_.Update(_1071_v48, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1074_v51).Cardinality()), _dafny.IntOfUint32((_1071_v48).Cardinality()))).Uint32(), _1076_v52)
				})())
				var _1077_v53 _dafny.Array
				_ = _1077_v53
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw179
				_1077_v53 = _nw179
				var _1078_v56 _dafny.Map
				_ = _1078_v56
				_1078_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm37(_dafny.IntOfUint32((_1071_v48).Cardinality()), globalState), (func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter35 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg43 _dafny.Int) interface{} {
							return coer43(arg43)
						}
					}((func(_1079_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1080_i9 _dafny.Int) _dafny.CodePoint {
							return _1079_v52
						}
					})(_1076_v52)))).Elements()); ; {
						_compr_32, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _1081_v55 _dafny.CodePoint
						_1081_v55 = interface{}(_compr_32).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_1082_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1080_i9 _dafny.Int) _dafny.CodePoint {
								return _1082_v52
							}
						})(_1076_v52))), _1081_v55) {
							_coll32.Add(_1081_v55, (_1018_v6).F20())
						}
					}
					return _coll32.ToMap()
				}()).Cardinality())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1077_v53), 0))
				_ = _index177
				(_1077_v53).ArraySet1((func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter36 := _dafny.Iterate((_1078_v56).Keys().Elements()); ; {
						_compr_33, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1083_v54 _dafny.MultiSet
						_1083_v54 = interface{}(_compr_33).(_dafny.MultiSet)
						if (_1078_v56).Contains(_1083_v54) {
							_coll33.Add(_1083_v54, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg45 _dafny.Int) interface{} {
									return coer45(arg45)
								}
							}((func(_1084_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1085_i10 _dafny.Int) _dafny.CodePoint {
									return _1084_v52
								}
							})(_1076_v52)))).Cardinality()))
						}
					}
					return _coll33.ToMap()
				}()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F20()), (_this).F20())), (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1077_v53), 0))
				_ = _index178
				(_1077_v53).ArraySet1((_1018_v6).F21(), (_index178).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1077_v53), 0))
				_ = _index179
				(_1077_v53).ArraySet1(!((_1077_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1077_v53), 0))).Int()).(bool)) || ((p0).Cmp(_dafny.IntOfInt64(208)) == 0), (_index179).Int())
			} else {
				var _1086_v57 *C3
				_ = _1086_v57
				var _nw180 *C3 = New_C3_()
				_ = _nw180
				_nw180.Ctor__(((_1018_v6).F20()).Minus((_1018_v6).F20()), (p0).Cmp((_1018_v6).F20()) >= 0, _dafny.IntOfInt64(582), ((_1018_v6).F21()) == ((_1018_v6).F21()))
				_1086_v57 = _nw180
				var _1087_v58 _dafny.Sequence
				_ = _1087_v58
				_1087_v58 = _dafny.SeqOf(_dafny.IntOfInt64(251))
				var _1088_v59 _dafny.MultiSet
				_ = _1088_v59
				_1088_v59 = _dafny.MultiSetOf(_1067_v44)
				var _1089_v60 _dafny.Map
				_ = _1089_v60
				_1089_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_1086_v57).F25())
				var _1090_v61 _dafny.CodePoint
				_ = _1090_v61
				_1090_v61 = _dafny.CodePoint('m')
				var _1091_v62 _dafny.Map
				_ = _1091_v62
				_1091_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1089_v60, _1090_v61)
				var _1092_v63 _dafny.Array
				_ = _1092_v63
				var _nwElement0_42 _dafny.Sequence = _1087_v58
				_ = _nwElement0_42
				var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(11))
				_ = _nw181
				(_nw181).ArraySet1(_nwElement0_42, 0)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1087_v58, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1088_v59).Contains(_1067_v44) {
						return (_1088_v59).Multiplicity(_1067_v44)
					}
					return _dafny.IntOfUint32((_1071_v48).Cardinality())
				})(), _dafny.IntOfUint32((_1087_v58).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1071_v48).Cardinality())), (Companion_Default___.SafeIndex((_1091_v62).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1087_v58, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1088_v59).Contains(_1067_v44) {
						return (_1088_v59).Multiplicity(_1067_v44)
					}
					return _dafny.IntOfUint32((_1071_v48).Cardinality())
				})(), _dafny.IntOfUint32((_1087_v58).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1071_v48).Cardinality()))).Cardinality()))).Uint32(), (_1018_v6).F20()), 1)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(73))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_1093_i11 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gxcytgdfb")).Cardinality())
				})), _1087_v58), 2)
				(_nw181).ArraySet1(_dafny.SeqOf((_1018_v6).F20()), 3)
				(_nw181).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(636))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_1094_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1095_i12 _dafny.Int) _dafny.Int {
						return _1094_p0
					}
				})(p0))), 4)
				(_nw181).ArraySet1(Companion_Default___.Fm28((_1018_v6).F20(), globalState), 5)
				(_nw181).ArraySet1(_1087_v58, 6)
				(_nw181).ArraySet1(_1087_v58, 7)
				(_nw181).ArraySet1(_1087_v58, 8)
				(_nw181).ArraySet1(_1087_v58, 9)
				(_nw181).ArraySet1(_1087_v58, 10)
				_1092_v63 = _nw181
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1092_v63), 0))
				_ = _index180
				(_1092_v63).ArraySet1(_1087_v58, (_index180).Int())
				var _1096_v64 _dafny.MultiSet
				_ = _1096_v64
				_1096_v64 = _dafny.MultiSetOf(true, (_this).F24())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1092_v63), 0))
				_ = _index181
				var _rhs201 _dafny.Int = (_dafny.Zero).Minus((((_1096_v64).Cardinality()).Plus(_dafny.IntOfInt64(908))).Plus((_1018_v6).F20()))
				_ = _rhs201
				var _rhs202 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1087_v58, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1087_v58).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1016_v4).Cardinality())), Companion_Default___.Fm28((_1018_v6).F20(), globalState))
				_ = _rhs202
				var _lhs177 *GlobalState = globalState
				_ = _lhs177
				var _lhs178 _dafny.Array = _1092_v63
				_ = _lhs178
				var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1092_v63), 0))
				_ = _lhs179
				_lhs177.F17 = _rhs201
				(_lhs178).ArraySet1(_rhs202, (_lhs179).Int())
				(globalState).F18 = !((_1018_v6).F21()) || (!(!((_1018_v6).F21())))
				_1017_v5 = _1017_v5
				var _1097_v65 _dafny.Array
				_ = _1097_v65
				var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw182
				_1097_v65 = _nw182
				var _1098_v66 _dafny.Map
				_ = _1098_v66
				_1098_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1097_v65, (_1018_v6).F21())
				var _1099_v67 _dafny.Map
				_ = _1099_v67
				_1099_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F20(), p1), (_1086_v57).F25())
				var _1100_v68 _dafny.Map
				_ = _1100_v68
				_1100_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F20(), (_this).F21())
				(globalState).F17 = (((func() _dafny.Map {
					if (_this).F21() {
						return Companion_Default___.Fm38((_dafny.Zero).Minus((_1086_v57).F25()), (func() bool {
							if (_1098_v66).Contains(_1097_v65) {
								return (_1098_v66).Get(_1097_v65).(bool)
							}
							return p1
						})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qrgiiat")).Cardinality()), (_1018_v6).F21(), globalState)
					}
					return _1099_v67
				})()).Update(_1100_v68, (p0).Minus((_1086_v57).F25()))).Cardinality()
			}
			(globalState).F5 = (_1018_v6).F20()
			var _1101_v69 _dafny.Sequence
			_ = _1101_v69
			_1101_v69 = _dafny.SeqOf(_1071_v48)
			var _1102_v70 _dafny.Array
			_ = _1102_v70
			var _nwElement0_43 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-767))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_1103_i13 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))
			_ = _nwElement0_43
			var _nw183 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(13))
			_ = _nw183
			(_nw183).ArraySet1(_nwElement0_43, 0)
			(_nw183).ArraySet1(_1071_v48, 1)
			(_nw183).ArraySet1(Companion_Default___.Fm26((_1018_v6).F21(), globalState), 2)
			(_nw183).ArraySet1(_1071_v48, 3)
			(_nw183).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm26((_this).F24(), globalState), _1071_v48), 4)
			(_nw183).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1071_v48, _dafny.UnicodeSeqOfUtf8Bytes("gntqqvui")), 5)
			(_nw183).ArraySet1(_1071_v48, 6)
			(_nw183).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1101_v69).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1101_v69).Cardinality()))).Uint32()).(_dafny.Sequence), _1071_v48), 7)
			(_nw183).ArraySet1(_1071_v48, 8)
			(_nw183).ArraySet1(_1071_v48, 9)
			(_nw183).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pxfbmf"), 10)
			(_nw183).ArraySet1(_1071_v48, 11)
			(_nw183).ArraySet1(_1071_v48, 12)
			_1102_v70 = _nw183
			_1102_v70 = _1102_v70
			if !(true) {
				(globalState).F4 = (_1018_v6).F21()
				(globalState).F15 = _dafny.UnicodeSeqOfUtf8Bytes("yhpdkduv")
				var _1104_v71 _dafny.CodePoint
				_ = _1104_v71
				_1104_v71 = _dafny.CodePoint('r')
				var _1105_v72 _dafny.Array
				_ = _1105_v72
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_32
				var _nw184 _dafny.Array
				_ = _nw184
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw184 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = func(_1106_i14 _dafny.Int) bool {
						return (_this).F21()
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw184 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw184).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw184).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1105_v72 = _nw184
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1105_v72), 0))
				_ = _index182
				(_1105_v72).ArraySet1((_1018_v6).F21(), (_index182).Int())
				var _1107_v73 _dafny.Map
				_ = _1107_v73
				_1107_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_this).F21())
				var _1108_v74 _dafny.Map
				_ = _1108_v74
				_1108_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), ((_1107_v73).Update((_this).F24(), (_1018_v6).F21())).Merge(_1107_v73))
				var _1109_v75 _dafny.Array
				_ = _1109_v75
				var _nwElement0_44 _dafny.Array = _1017_v5
				_ = _nwElement0_44
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(23))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_44, 0)
				(_nw185).ArraySet1(_1017_v5, 1)
				(_nw185).ArraySet1(_1017_v5, 2)
				(_nw185).ArraySet1(_1017_v5, 3)
				(_nw185).ArraySet1(_1017_v5, 4)
				(_nw185).ArraySet1(_1017_v5, 5)
				(_nw185).ArraySet1(_1017_v5, 6)
				(_nw185).ArraySet1(_1017_v5, 7)
				(_nw185).ArraySet1(_1017_v5, 8)
				(_nw185).ArraySet1(_1017_v5, 9)
				(_nw185).ArraySet1(_1017_v5, 10)
				(_nw185).ArraySet1(_1017_v5, 11)
				(_nw185).ArraySet1(_1017_v5, 12)
				(_nw185).ArraySet1(_1017_v5, 13)
				(_nw185).ArraySet1(_1017_v5, 14)
				(_nw185).ArraySet1(_1017_v5, 15)
				(_nw185).ArraySet1(_1017_v5, 16)
				(_nw185).ArraySet1(_1017_v5, 17)
				(_nw185).ArraySet1(_1017_v5, 18)
				(_nw185).ArraySet1(_1017_v5, 19)
				(_nw185).ArraySet1((func() _dafny.Array {
					if true {
						return _1017_v5
					}
					return _1017_v5
				})(), 20)
				(_nw185).ArraySet1((func() _dafny.Array {
					if (_this).F24() {
						return _1017_v5
					}
					return _1017_v5
				})(), 21)
				(_nw185).ArraySet1((func() _dafny.Array {
					if (_1018_v6).F21() {
						return _1017_v5
					}
					return _1017_v5
				})(), 22)
				_1109_v75 = _nw185
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))
				_ = _index183
				(_1109_v75).ArraySet1(_1017_v5, (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1105_v72), 0))
				_ = _index184
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))
				_ = _index185
				var _rhs203 _dafny.CodePoint = _1104_v71
				_ = _rhs203
				var _rhs204 bool = (_1018_v6).F21()
				_ = _rhs204
				var _rhs205 bool = (_1018_v6).F21()
				_ = _rhs205
				var _rhs206 _dafny.Map = _1108_v74
				_ = _rhs206
				var _rhs207 _dafny.Array = _1017_v5
				_ = _rhs207
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 _dafny.Array = _1105_v72
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1105_v72), 0))
				_ = _lhs182
				var _lhs183 _dafny.Array = _1109_v75
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))
				_ = _lhs184
				_1104_v71 = _rhs203
				_lhs180.F14 = _rhs204
				(_lhs181).ArraySet1(_rhs205, (_lhs182).Int())
				_1108_v74 = _rhs206
				(_lhs183).ArraySet1(_rhs207, (_lhs184).Int())
				var _1110_v76 _dafny.MultiSet
				_ = _1110_v76
				_1110_v76 = _dafny.MultiSetOf(p1)
				var _1111_v77 _dafny.Map
				_ = _1111_v77
				_1111_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_1018_v6).F20())
				var _1112_v78 _dafny.Sequence
				_ = _1112_v78
				_1112_v78 = _dafny.SeqOf((_this).F20(), (_1018_v6).F20(), (_this).F20(), (_1111_v77).Cardinality())
				var _rhs208 bool = (((_1110_v76).Update((_1018_v6).F21(), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Int {
					if (_1068_v45).Contains((_1018_v6).F20()) {
						return (_1068_v45).Multiplicity((_1018_v6).F20())
					}
					return (_this).F20()
				})(), (_this).F20(), _dafny.IntOfUint32((_1112_v78).Cardinality()), (_1018_v6).F20())).Cardinality())))).Union(_dafny.MultiSetFromSeq(_1016_v4))).IsProperSubsetOf(_1110_v76)
				_ = _rhs208
				var _rhs209 _dafny.Int = Companion_Default___.SafeDivisionInt((_1018_v6).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1101_v69).Select((Companion_Default___.SafeIndex((_1018_v6).F20(), _dafny.IntOfUint32((_1101_v69).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("jwncqvclf"))).Cardinality()))
				_ = _rhs209
				var _rhs210 _dafny.Int = (_1018_v6).F20()
				_ = _rhs210
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				_lhs185.F4 = _rhs208
				_lhs186.F5 = _rhs209
				_lhs187.F17 = _rhs210
				var _1113_v79 _dafny.Sequence
				_ = _1113_v79
				_1113_v79 = _dafny.SeqOf(_1068_v45, (_1068_v45).Union(_1068_v45))
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _arr2
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _index186
				(_arr2).ArraySet1(Companion_Default___.SafeModuloInt((_this).F20(), (_1018_v6).F20()), (_index186).Int())
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1102_v70), 0))
				_ = _index187
				(_1102_v70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("u"), (_index187).Int())
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _arr3
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _index188
				(_arr3).ArraySet1(((_dafny.Zero).Minus((_1018_v6).F20())).Times(p0), (_index188).Int())
				var _arr4 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _arr4
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _index189
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1102_v70), 0))
				_ = _index190
				var _arr5 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _arr5
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _index191
				var _rhs211 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(185))).Uint32(), func(coer49 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_1114_v45 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1115_i15 _dafny.Int) _dafny.MultiSet {
						return _1114_v45
					}
				})(_1068_v45)))
				_ = _rhs211
				var _rhs212 _dafny.Int = (_dafny.Zero).Minus((_this).Fm14((_1110_v76).IsSubsetOf(_dafny.MultiSetOf(!((_1018_v6).F21()), (_1018_v6).F21(), !((_1018_v6).F21()), (_1018_v6).F21())), (_1018_v6).F20(), (_1018_v6).F20(), (_1018_v6).F21(), globalState))
				_ = _rhs212
				var _rhs213 _dafny.Int = (_1018_v6).F20()
				_ = _rhs213
				var _rhs214 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1071_v48, _1071_v48)
				_ = _rhs214
				var _rhs215 _dafny.Int = ((func() _dafny.Int {
					if (_1110_v76).Contains((_1018_v6).F21()) {
						return (_1110_v76).Multiplicity((_1018_v6).F21())
					}
					return (_dafny.Zero).Minus((_1018_v6).F20())
				})()).Plus(Companion_Default___.SafeDivisionInt((_1018_v6).F20(), (_1018_v6).F20()))
				_ = _rhs215
				var _lhs188 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _lhs188
				var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				var _lhs191 _dafny.Array = _1102_v70
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1102_v70), 0))
				_ = _lhs192
				var _lhs193 _dafny.Array = _dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_dafny.ArrayCastTo((_1109_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1109_v75), 0))).Int()))), 0))
				_ = _lhs194
				_1113_v79 = _rhs211
				(_lhs188).ArraySet1(_rhs212, (_lhs189).Int())
				_lhs190.F5 = _rhs213
				(_lhs191).ArraySet1(_rhs214, (_lhs192).Int())
				(_lhs193).ArraySet1(_rhs215, (_lhs194).Int())
			} else {
				(globalState).F17 = (_1018_v6).Fm1((_1018_v6).F20(), p0, (_this).F21(), globalState)
				var _1116_v80 _dafny.Sequence
				_ = _1116_v80
				_1116_v80 = _dafny.SeqOf((_dafny.Zero).Minus(p0), (_1018_v6).F20(), (_this).F20(), (_1018_v6).F20())
				_1116_v80 = _dafny.Companion_Sequence_.Concatenate(_1116_v80, _1116_v80)
				var _1117_v81 *C2
				_ = _1117_v81
				var _nw186 *C2 = New_C2_()
				_ = _nw186
				_nw186.Ctor__(Companion_Default___.SafeDivisionInt((_1018_v6).F20(), (_1018_v6).F20()), (_this).Fm13((_1018_v6).F20(), globalState))
				_1117_v81 = _nw186
				var _1118_v82 _dafny.Array
				_ = _1118_v82
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_33
				var _nw187 _dafny.Array
				_ = _nw187
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw187 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1119_v6 T0) func(_dafny.Int) bool {
						return func(_1120_i16 _dafny.Int) bool {
							return !((_1119_v6).F21())
						}
					})(_1018_v6)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw187 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw187).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw187).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1118_v82 = _nw187
				var _1121_v83 _dafny.MultiSet
				_ = _1121_v83
				_1121_v83 = _dafny.MultiSetOf(_1118_v82, _1118_v82)
				_1121_v83 = _dafny.MultiSetOf(_1118_v82)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1118_v82), 0))
				_ = _index192
				(_1118_v82).ArraySet1((_this).F24(), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1118_v82), 0))
				_ = _index193
				(_1118_v82).ArraySet1((_this).F24(), (_index193).Int())
				var _1122_v84 _dafny.MultiSet
				_ = _1122_v84
				_1122_v84 = _dafny.MultiSetOf((_1071_v48).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1071_v48).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _1123_v85 bool
				_ = _1123_v85
				_1123_v85 = (_1018_v6).F21()
				var _1124_v86 _dafny.Sequence
				_ = _1124_v86
				_1124_v86 = _1116_v80
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1118_v82), 0))
				_ = _index194
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1118_v82), 0))
				_ = _index195
				var _rhs216 bool = !(((_1122_v84).Cardinality()).Cmp((_1018_v6).F20()) <= 0)
				_ = _rhs216
				var _rhs217 bool = (func() bool {
					if (_1018_v6).F21() {
						return (_this).F21()
					}
					return (_1123_v85)
				})()
				_ = _rhs217
				var _rhs218 bool = (_1018_v6).F21()
				_ = _rhs218
				var _rhs219 bool = (_1018_v6).F21()
				_ = _rhs219
				var _rhs220 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1116_v80, (func() _dafny.Sequence {
					if (_this).F24() {
						return _1116_v80
					}
					return (_1124_v86)
				})())).Cardinality()))
				_ = _rhs220
				var _lhs195 _dafny.Array = _1118_v82
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1118_v82), 0))
				_ = _lhs196
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 _dafny.Array = _1118_v82
				_ = _lhs198
				var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1118_v82), 0))
				_ = _lhs199
				var _lhs200 *GlobalState = globalState
				_ = _lhs200
				r1 = _rhs216
				(_lhs195).ArraySet1(_rhs217, (_lhs196).Int())
				_lhs197.F4 = _rhs218
				(_lhs198).ArraySet1(_rhs219, (_lhs199).Int())
				_lhs200.F17 = _rhs220
			}
		} else {
			(globalState).F17 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_1125_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})), _dafny.UnicodeSeqOfUtf8Bytes("h"))).Cardinality()), _dafny.IntOfInt64(409))
			var _1126_v88 _dafny.Map
			_ = _1126_v88
			_1126_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_this).F21())
			var _1127_v89 _dafny.Map
			_ = _1127_v89
			_1127_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1126_v88, (_1018_v6).F21())
			var _1128_v90 _dafny.Map
			_ = _1128_v90
			_1128_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1066_v43).Cardinality(), (_1018_v6).F21())
			var _1129_v91 _dafny.Map
			_ = _1129_v91
			_1129_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1128_v90)
			var _1130_v92 _dafny.Map
			_ = _1130_v92
			_1130_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_1129_v91).Contains(p0) {
					return (_1129_v91).Get(p0).(_dafny.Map)
				}
				return _1128_v90
			})(), (_this).F24())
			(globalState).F14 = (_1130_v92).Contains(func() _dafny.Map {
				var _coll34 = _dafny.NewMapBuilder()
				_ = _coll34
				for _iter37 := _dafny.Iterate((_dafny.MultiSetOf((_1127_v89).Cardinality())).Elements()); ; {
					_compr_34, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1131_v87 _dafny.Int
					_1131_v87 = interface{}(_compr_34).(_dafny.Int)
					if (_dafny.MultiSetOf((_1127_v89).Cardinality())).Contains(_1131_v87) {
						_coll34.Add((_1131_v87).Times((_1018_v6).F20()), (_998_v1).Dtor_cf16())
					}
				}
				return _coll34.ToMap()
			}())
			if p1 {
				var _1132_v93 _dafny.Array
				_ = _1132_v93
				var _nwElement0_45 _dafny.MultiSet = _1068_v45
				_ = _nwElement0_45
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(5))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_45, 0)
				(_nw188).ArraySet1(_1068_v45, 1)
				(_nw188).ArraySet1(_1068_v45, 2)
				(_nw188).ArraySet1((_1068_v45).Difference(_1068_v45), 3)
				(_nw188).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfInt64(993)), 4)
				_1132_v93 = _nw188
				var _1133_v94 _dafny.Sequence
				_ = _1133_v94
				_1133_v94 = _dafny.SeqOf((_1018_v6).F20())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1132_v93), 0))
				_ = _index196
				(_1132_v93).ArraySet1((_dafny.MultiSetFromSeq(_1133_v94)).Intersection(_1068_v45), (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1132_v93), 0))
				_ = _index197
				(_1132_v93).ArraySet1(_1068_v45, (_index197).Int())
				r1 = ((_1066_v43).Difference(_1066_v43)).Equals((_1066_v43).Difference(_1066_v43))
				var _1134_v95 _dafny.CodePoint
				_ = _1134_v95
				_1134_v95 = _dafny.CodePoint('u')
				var _1135_v96 *C1
				_ = _1135_v96
				var _nw189 *C1 = New_C1_()
				_ = _nw189
				_nw189.Ctor__((_dafny.IntOfInt64(955)).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_1136_i18 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				})), (Companion_Default___.SafeIndex(((_1132_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1132_v93), 0))).Int()).(_dafny.MultiSet)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}(func(_1137_i18 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality()))).Uint32(), _1134_v95)).Cardinality()))), (func() _dafny.Array {
					if (_this).F21() {
						return _1017_v5
					}
					return _1017_v5
				})(), (_1018_v6).F20(), (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(420))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1138_v95 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1139_i19 _dafny.Int) _dafny.CodePoint {
						return _1138_v95
					}
				})(_1134_v95))), (Companion_Default___.SafeIndex((_1018_v6).F20(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(420))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1140_v95 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1141_i19 _dafny.Int) _dafny.CodePoint {
						return _1140_v95
					}
				})(_1134_v95)))).Cardinality()))).Uint32(), _1134_v95)).Cardinality())).Cmp((_1018_v6).F20()) < 0)
				_1135_v96 = _nw189
				var _1142_v97 _dafny.Sequence
				_ = _1142_v97
				_1142_v97 = _dafny.UnicodeSeqOfUtf8Bytes("d")
				var _1143_v98 _dafny.Map
				_ = _1143_v98
				_1143_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(453), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1142_v97).Cardinality())))
				(globalState).F18 = (_1018_v6).Fm0(_1143_v98, _1142_v97, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1018_v6).F21())), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1133_v94).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1018_v6).F21()))).Cardinality()))).Uint32(), _1126_v88), globalState)
				(globalState).F18 = (_1135_v96.F27).Cmp((_1018_v6).F20()) > 0
			} else {
				var _1144_v99 *C1
				_ = _1144_v99
				var _nw190 *C1 = New_C1_()
				_ = _nw190
				_nw190.Ctor__(p0, _1017_v5, p0, !(!(!(true) || ((_this).F21()))))
				_1144_v99 = _nw190
				var _1145_v100 _dafny.Array
				_ = _1145_v100
				var _nw191 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw191
				_1145_v100 = _nw191
				_1145_v100 = _1145_v100
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1145_v100), 0))
				_ = _index198
				(_1145_v100).ArraySet1((_this).F24(), (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1145_v100), 0))
				_ = _index199
				(_1145_v100).ArraySet1(p1, (_index199).Int())
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1017_v5), 0))
				_ = _index200
				(_1017_v5).ArraySet1(_dafny.IntOfInt64(753), (_index200).Int())
				var _1146_v101 _dafny.Map
				_ = _1146_v101
				_1146_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1016_v4)
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1017_v5), 0))
				_ = _index201
				(_1017_v5).ArraySet1(((_1146_v101).Merge(_1146_v101)).Cardinality(), (_index201).Int())
				_1067_v44 = (_1067_v44).Union((_1067_v44).Difference(_1067_v44))
			}
			(globalState).F5 = Companion_Default___.SafeDivisionInt((_1018_v6).F20(), p0)
			(globalState).F14 = (_this).F21()
		}
		var _1147_v102 _dafny.MultiSet
		_ = _1147_v102
		_1147_v102 = _dafny.MultiSetOf(p1)
		if (!((_dafny.MultiSetOf((_1018_v6).F21())).IsDisjointFrom(_1147_v102))) && ((_this).F24()) {
			var _1148_v103 _dafny.Sequence
			_ = _1148_v103
			_1148_v103 = _dafny.SeqOf(p0, p0, (_this).F20())
			var _1149_v105 _dafny.Map
			_ = _1149_v105
			_1149_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F20())
			var _1150_v106 _dafny.Map
			_ = _1150_v106
			_1150_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _dafny.UnicodeSeqOfUtf8Bytes("keo"))
			var _1151_v107 _dafny.Map
			_ = _1151_v107
			_1151_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
			var _1152_v108 _dafny.Sequence
			_ = _1152_v108
			_1152_v108 = _dafny.UnicodeSeqOfUtf8Bytes("rupn")
			var _1153_v109 _dafny.Map
			_ = _1153_v109
			_1153_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_1018_v6).F21())
			var _rhs221 bool = (_this).Fm0((func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(305))); ; {
					_compr_35, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _1154_v104 _dafny.Int
					_1154_v104 = interface{}(_compr_35).(_dafny.Int)
					if ((_dafny.IntOfInt64(297)).Cmp(_1154_v104) <= 0) && ((_1154_v104).Cmp(_dafny.IntOfInt64(305)) < 0) {
						_coll35.Add((_1154_v104).Times((_1018_v6).F20()), (_this).F20())
					}
				}
				return _coll35.ToMap()
			}()).Merge(_1149_v105), (func() _dafny.Sequence {
				if (_1150_v106).Contains((func() bool {
					if (_1151_v107).Contains((_1018_v6).F20()) {
						return (_1151_v107).Get((_1018_v6).F20()).(bool)
					}
					return p1
				})()) {
					return (_1150_v106).Get((func() bool {
						if (_1151_v107).Contains((_1018_v6).F20()) {
							return (_1151_v107).Get((_1018_v6).F20()).(bool)
						}
						return p1
					})()).(_dafny.Sequence)
				}
				return _1152_v108
			})(), _dafny.SeqOf(_1153_v109), globalState)
			_ = _rhs221
			var _rhs222 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1148_v103, _1148_v103)
			_ = _rhs222
			var _lhs201 *GlobalState = globalState
			_ = _lhs201
			_lhs201.F4 = _rhs221
			_1148_v103 = _rhs222
			(globalState).F5 = Companion_Default___.SafeModuloInt(p0, (_1018_v6).F20())
			(globalState).F5 = (_dafny.Zero).Minus((_1018_v6).F20())
			if false {
				var _1155_v110 _dafny.Map
				_ = _1155_v110
				_1155_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), Companion_Default___.Fm17((_this).F21(), globalState))
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1017_v5), 0))
				_ = _index202
				(_1017_v5).ArraySet1(((_1155_v110).Merge(_1155_v110)).Cardinality(), (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1017_v5), 0))
				_ = _index203
				var _rhs223 _dafny.MultiSet = _1147_v102
				_ = _rhs223
				var _rhs224 _dafny.Int = (_dafny.Zero).Minus((_this).F20())
				_ = _rhs224
				var _rhs225 _dafny.Int = (_1147_v102).Cardinality()
				_ = _rhs225
				var _rhs226 _dafny.Int = Companion_Default___.SafeModuloInt(((_1018_v6).F20()).Times((_this).F20()), (_this).F20())
				_ = _rhs226
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				var _lhs203 _dafny.Array = _1017_v5
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1017_v5), 0))
				_ = _lhs204
				var _lhs205 *GlobalState = globalState
				_ = _lhs205
				_1147_v102 = _rhs223
				_lhs202.F17 = _rhs224
				(_lhs203).ArraySet1(_rhs225, (_lhs204).Int())
				_lhs205.F5 = _rhs226
				var _1156_v111 D10
				_ = _1156_v111
				_1156_v111 = Companion_D10_.Create_DC25_(_1068_v45)
				var _1157_v112 _dafny.Sequence
				_ = _1157_v112
				_1157_v112 = _dafny.SeqOf((_1156_v111).Dtor_cf34())
				var _1158_v113 _dafny.Array
				_ = _1158_v113
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw192
				_1158_v113 = _nw192
				var _1159_v114 _dafny.Sequence
				_ = _1159_v114
				_1159_v114 = _dafny.SeqOf(_1158_v113)
				var _1160_v115 _dafny.Map
				_ = _1160_v115
				_1160_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(776))).Uint32(), func(coer55 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1161_v45 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1162_i20 _dafny.Int) _dafny.MultiSet {
						return _1161_v45
					}
				})(_1068_v45))), _1157_v112), _dafny.Companion_Sequence_.Concatenate(_1159_v114, _1159_v114))
				var _1163_v116 D3
				_ = _1163_v116
				_1163_v116 = Companion_D3_.Create_DC5_(_1159_v114)
				_1160_v115 = (_1160_v115).Update(_dafny.SeqOf(_1068_v45), (_1163_v116).Dtor_cf9())
				var _1164_v117 _dafny.Sequence
				_ = _1164_v117
				_1164_v117 = _dafny.SeqOf(_1016_v4, _dafny.Companion_Sequence_.Update(_1016_v4, (Companion_Default___.SafeIndex((_1153_v109).Cardinality(), _dafny.IntOfUint32((_1016_v4).Cardinality()))).Uint32(), (_1018_v6).F21()), _1016_v4, _1016_v4, _1016_v4)
				var _1165_v118 _dafny.Map
				_ = _1165_v118
				_1165_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1164_v117, _dafny.UnicodeSeqOfUtf8Bytes("cbrjrlv"))
				var _1166_v119 _dafny.CodePoint
				_ = _1166_v119
				_1166_v119 = _dafny.CodePoint('t')
				var _rhs227 bool = (_this).F21()
				_ = _rhs227
				var _rhs228 bool = (_1018_v6).F21()
				_ = _rhs228
				var _rhs229 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1152_v108, (func() _dafny.Sequence {
					if (_1165_v118).Contains(_1164_v117) {
						return (_1165_v118).Get(_1164_v117).(_dafny.Sequence)
					}
					return _1152_v108
				})()), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1152_v108, (func() _dafny.Sequence {
					if (_1165_v118).Contains(_1164_v117) {
						return (_1165_v118).Get(_1164_v117).(_dafny.Sequence)
					}
					return _1152_v108
				})())).Cardinality()))).Uint32(), _1166_v119)
				_ = _rhs229
				var _lhs206 *GlobalState = globalState
				_ = _lhs206
				var _lhs207 *GlobalState = globalState
				_ = _lhs207
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				_lhs206.F14 = _rhs227
				_lhs207.F14 = _rhs228
				_lhs208.F15 = _rhs229
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1017_v5), 0))
				_ = _index204
				(_1017_v5).ArraySet1((_this).Fm14(true, (_1018_v6).F20(), (_1018_v6).F20(), !_dafny.Companion_Sequence_.Equal(_1152_v108, Companion_Default___.Fm26((_this).F21(), globalState)), globalState), (_index204).Int())
				_1153_v109 = (_1153_v109).Update(true, (func() bool {
					if (_1151_v107).Contains((_1153_v109).Cardinality()) {
						return (_1151_v107).Get((_1153_v109).Cardinality()).(bool)
					}
					return (_1018_v6).F21()
				})())
			} else {
				var _1167_v120 _dafny.Map
				_ = _1167_v120
				_1167_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(93))
				(globalState).F5 = (_dafny.Zero).Minus(((func() _dafny.Int {
					if (_1167_v120).Contains((_1018_v6).F21()) {
						return (_1167_v120).Get((_1018_v6).F21()).(_dafny.Int)
					}
					return (_1018_v6).F20()
				})()).Minus((_1018_v6).F20()))
				var _1168_v121 _dafny.Array
				_ = _1168_v121
				var _nw193 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw193
				_1168_v121 = _nw193
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1168_v121), 0))
				_ = _index205
				(_1168_v121).ArraySet1((_this).F21(), (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1168_v121), 0))
				_ = _index206
				(_1168_v121).ArraySet1((_1016_v4).Select((Companion_Default___.SafeIndex((_1018_v6).F20(), _dafny.IntOfUint32((_1016_v4).Cardinality()))).Uint32()).(bool), (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1168_v121), 0))
				_ = _index207
				(_1168_v121).ArraySet1((_this).Fm13(p0, globalState), (_index207).Int())
				(globalState).F5 = (((_1149_v105).Merge(_1149_v105)).Cardinality()).Plus((_1018_v6).F20())
				(globalState).F17 = (_this).F20()
			}
			var _1169_v122 _dafny.Array
			_ = _1169_v122
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw194
			_1169_v122 = _nw194
			_1169_v122 = _1169_v122
		} else {
			_1017_v5 = (func() _dafny.Array {
				if (_this).F21() {
					return _1017_v5
				}
				return _1017_v5
			})()
			var _1170_v123 _dafny.Sequence
			_ = _1170_v123
			_1170_v123 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			var _1171_v124 _dafny.Map
			_ = _1171_v124
			_1171_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1170_v123).Cardinality()), _dafny.IntOfUint32((_1016_v4).Cardinality()))
			var _1172_v125 _dafny.Map
			_ = _1172_v125
			_1172_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F20(), (func() _dafny.Int {
				if (_1171_v124).Contains((_this).F20()) {
					return (_1171_v124).Get((_this).F20()).(_dafny.Int)
				}
				return (_this).F20()
			})())
			(globalState).F4 = (_1018_v6).Fm0(_1172_v125, _1170_v123, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-561))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_1173_v6 T0) func(_dafny.Int) _dafny.Map {
				return func(_1174_i21 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1173_v6).F21(), (_1173_v6).F21())
				}
			})(_1018_v6))), globalState)
			var _1175_v126 _dafny.Array
			_ = _1175_v126
			var _1176_v127 _dafny.Int
			_ = _1176_v127
			var _out16 _dafny.Array
			_ = _out16
			var _out17 _dafny.Int
			_ = _out17
			_out16, _out17 = (_this).M1(_dafny.IntOfUint32((_1170_v123).Cardinality()), (_this).F24(), _1018_v6, globalState)
			_1175_v126 = _out16
			_1176_v127 = _out17
			var _1177_v128 _dafny.Map
			_ = _1177_v128
			_1177_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F21())
			var _1178_v129 _dafny.Sequence
			_ = _1178_v129
			_1178_v129 = _dafny.SeqOf((_1018_v6).F20())
			var _1179_v130 _dafny.CodePoint
			_ = _1179_v130
			_1179_v130 = _dafny.CodePoint('n')
			var _1180_v131 _dafny.Map
			_ = _1180_v131
			_1180_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_1018_v6).F21())
			var _1181_v132 _dafny.Sequence
			_ = _1181_v132
			_1181_v132 = _dafny.SeqOf(_1180_v131)
			var _1182_v133 _dafny.Map
			_ = _1182_v133
			_1182_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1018_v6).Fm0(_1171_v124, _1170_v123, _1181_v132, globalState))
			_1177_v128 = (_1177_v128).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1178_v129, _dafny.Companion_Sequence_.Update(_1178_v129, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1179_v130)).Cardinality()), _dafny.IntOfUint32((_1178_v129).Cardinality()))).Uint32(), (_1068_v45).Cardinality()))).Cardinality()), (func() bool {
				if (_1182_v133).Contains((_1018_v6).F21()) {
					return (_1182_v133).Get((_1018_v6).F21()).(bool)
				}
				return (_this).F21()
			})())
			var _1183_v134 _dafny.Sequence
			_ = _1183_v134
			_1183_v134 = _dafny.SeqOf(_1175_v126, _1175_v126, _1175_v126, _1017_v5, _1175_v126)
			r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1183_v134, _1183_v134), (func() _dafny.Sequence {
				if (_this).F21() {
					return _dafny.SeqOf(_1175_v126, _1175_v126)
				}
				return _1183_v134
			})())
		}
		var _1184_v135 _dafny.Sequence
		_ = _1184_v135
		_1184_v135 = _dafny.SeqOf(_1017_v5)
		r0 = _dafny.Companion_Sequence_.Concatenate(_1184_v135, _1184_v135)
		var _1185_v138 _dafny.Map
		_ = _1185_v138
		_1185_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F21())
		var _1186_v139 _dafny.Map
		_ = _1186_v139
		_1186_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _1187_v140 _dafny.Sequence
		_ = _1187_v140
		_1187_v140 = _dafny.SeqOf(_1185_v138, Companion_Default___.Fm15(_dafny.IntOfInt64(522), (_1018_v6).F20(), (func() _dafny.Int {
			if (_1186_v139).Contains((_this).F24()) {
				return (_1186_v139).Get((_this).F24()).(_dafny.Int)
			}
			return p0
		})(), (_1018_v6).F20(), globalState), _1185_v138, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v6).F21(), (_this).F21()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()))
		r1 = (_this).Fm0(func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(702), _dafny.IntOfInt64(756))); ; {
				_compr_36, _ok39 := _iter39()
				if !_ok39 {
					break
				}
				var _1188_v136 _dafny.Int
				_1188_v136 = interface{}(_compr_36).(_dafny.Int)
				if ((_dafny.IntOfInt64(702)).Cmp(_1188_v136) <= 0) && ((_1188_v136).Cmp(_dafny.IntOfInt64(756)) < 0) {
					_coll36.Add((_1188_v136).Times((func() _dafny.Map {
						var _coll37 = _dafny.NewMapBuilder()
						_ = _coll37
						for _iter40 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), p1)).Cardinality()))).Elements()); ; {
							_compr_37, _ok40 := _iter40()
							if !_ok40 {
								break
							}
							var _1189_v137 _dafny.Int
							_1189_v137 = interface{}(_compr_37).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), p1)).Cardinality())), _1189_v137) {
								_coll37.Add((_1189_v137).Minus((_1018_v6).F20()), (_1067_v44).Cardinality())
							}
						}
						return _coll37.ToMap()
					}()).Cardinality()), (_dafny.Zero).Minus((_1018_v6).F20()))
				}
			}
			return _coll36.ToMap()
		}(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_1190_i22 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		})), _1187_v140, globalState)
		return r0, r1
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 bool, p2 T0, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		(globalState).F17 = p0
		if (_this).F24() {
			var _1191_v0 _dafny.Map
			_ = _1191_v0
			_1191_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (p2).F20())
			var _1192_v1 _dafny.Map
			_ = _1192_v1
			_1192_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F20(), _dafny.UnicodeSeqOfUtf8Bytes("tyiwlx"))
			var _1193_v2 _dafny.Map
			_ = _1193_v2
			_1193_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21())
			var _1194_v3 _dafny.Sequence
			_ = _1194_v3
			_1194_v3 = _dafny.SeqOf(_1193_v2, _1193_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (Companion_D7_.Create_DC17_((_this).F24())).Dtor_cf25()))
			var _1195_v4 *C3
			_ = _1195_v4
			var _nw195 *C3 = New_C3_()
			_ = _nw195
			_nw195.Ctor__((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(_1191_v0, (func() _dafny.Sequence {
				if (_1192_v1).Contains((p2).F20()) {
					return (_1192_v1).Get((p2).F20()).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-688))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_1196_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))
			})(), _1194_v3, globalState), true)).Cardinality()), (_this).F21(), p0, (_this).F21())
			_1195_v4 = _nw195
			var _nw196 *C3 = New_C3_()
			_ = _nw196
			_nw196.Ctor__(_dafny.IntOfInt64(365), (_this).F24(), _dafny.IntOfInt64(599), (p2).F21())
			_1195_v4 = _nw196
			if !((_1195_v4).F26()) {
				var _1197_v5 _dafny.Array
				_ = _1197_v5
				var _nw197 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw197
				_1197_v5 = _nw197
				var _1198_v6 _dafny.Sequence
				_ = _1198_v6
				_1198_v6 = _dafny.SeqOf(_1197_v5)
				var _1199_v7 D3
				_ = _1199_v7
				_1199_v7 = Companion_D3_.Create_DC5_(_1198_v6)
				var _pat_let_tv23 = _1198_v6
				_ = _pat_let_tv23
				var _1200_v8 _dafny.Array
				_ = _1200_v8
				var _nwElement0_46 D3 = _1199_v7
				_ = _nwElement0_46
				var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(15))
				_ = _nw198
				(_nw198).ArraySet1(_nwElement0_46, 0)
				(_nw198).ArraySet1(_1199_v7, 1)
				(_nw198).ArraySet1(_1199_v7, 2)
				(_nw198).ArraySet1(_1199_v7, 3)
				(_nw198).ArraySet1(_1199_v7, 4)
				(_nw198).ArraySet1(_1199_v7, 5)
				(_nw198).ArraySet1(_1199_v7, 6)
				(_nw198).ArraySet1(_1199_v7, 7)
				(_nw198).ArraySet1(func(_pat_let24_0 D3) D3 {
					return func(_1201_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let25_0 _dafny.Sequence) D3 {
							return func(_1202_dt__update_hcf9_h0 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC5_(_1202_dt__update_hcf9_h0)
							}(_pat_let25_0)
						}(_pat_let_tv23)
					}(_pat_let24_0)
				}(_1199_v7), 8)
				(_nw198).ArraySet1(_1199_v7, 9)
				(_nw198).ArraySet1(Companion_D3_.Create_DC5_(_1198_v6), 10)
				(_nw198).ArraySet1(Companion_D3_.Create_DC5_(_1198_v6), 11)
				(_nw198).ArraySet1(_1199_v7, 12)
				(_nw198).ArraySet1(_1199_v7, 13)
				(_nw198).ArraySet1(_1199_v7, 14)
				_1200_v8 = _nw198
				var _1203_v9 _dafny.Array
				_ = _1203_v9
				var _nwElement0_47 _dafny.Array = _1200_v8
				_ = _nwElement0_47
				var _nw199 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(23))
				_ = _nw199
				(_nw199).ArraySet1(_nwElement0_47, 0)
				(_nw199).ArraySet1(_1200_v8, 1)
				(_nw199).ArraySet1(_1200_v8, 2)
				(_nw199).ArraySet1(_1200_v8, 3)
				(_nw199).ArraySet1(_1200_v8, 4)
				(_nw199).ArraySet1(_1200_v8, 5)
				(_nw199).ArraySet1((Companion_D14_.Create_DC38_(_1200_v8)).Dtor_cf51(), 6)
				(_nw199).ArraySet1(_1200_v8, 7)
				(_nw199).ArraySet1(_1200_v8, 8)
				(_nw199).ArraySet1(_1200_v8, 9)
				(_nw199).ArraySet1(_1200_v8, 10)
				(_nw199).ArraySet1(_1200_v8, 11)
				(_nw199).ArraySet1(_1200_v8, 12)
				(_nw199).ArraySet1(_1200_v8, 13)
				(_nw199).ArraySet1(_1200_v8, 14)
				(_nw199).ArraySet1(_1200_v8, 15)
				(_nw199).ArraySet1(_1200_v8, 16)
				(_nw199).ArraySet1(_1200_v8, 17)
				(_nw199).ArraySet1(_1200_v8, 18)
				(_nw199).ArraySet1(_1200_v8, 19)
				(_nw199).ArraySet1(_1200_v8, 20)
				(_nw199).ArraySet1(_1200_v8, 21)
				(_nw199).ArraySet1(_1200_v8, 22)
				_1203_v9 = _nw199
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1203_v9), 0))
				_ = _index208
				(_1203_v9).ArraySet1(_1200_v8, (_index208).Int())
				var _1204_v10 _dafny.Sequence
				_ = _1204_v10
				_1204_v10 = _dafny.UnicodeSeqOfUtf8Bytes("dpcrmxv")
				var _1205_v11 _dafny.Set
				_ = _1205_v11
				_1205_v11 = _dafny.SetOf((_this).F24(), !(!((_1195_v4).F26())))
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1203_v9), 0))
				_ = _index209
				var _rhs230 _dafny.Int = Companion_Default___.SafeModuloInt(((_1195_v4).F25()).Times((_1195_v4).F25()), (_1195_v4).Fm1(_dafny.IntOfUint32((_1204_v10).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), p1)).Cardinality(), false, globalState))
				_ = _rhs230
				var _rhs231 bool = ((_this).F20()).Cmp(((_1205_v11).Cardinality()).Times((p2).F20())) <= 0
				_ = _rhs231
				var _rhs232 _dafny.Int = ((_1193_v2).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1204_v10, _1204_v10)).Cardinality()))
				_ = _rhs232
				var _rhs233 _dafny.Array = _1200_v8
				_ = _rhs233
				var _lhs209 *GlobalState = globalState
				_ = _lhs209
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				var _lhs211 *GlobalState = globalState
				_ = _lhs211
				var _lhs212 _dafny.Array = _1203_v9
				_ = _lhs212
				var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1203_v9), 0))
				_ = _lhs213
				_lhs209.F17 = _rhs230
				_lhs210.F4 = _rhs231
				_lhs211.F5 = _rhs232
				(_lhs212).ArraySet1(_rhs233, (_lhs213).Int())
				var _1206_v12 _dafny.Sequence
				_ = _1206_v12
				_1206_v12 = _dafny.SeqOf((p2).F21(), (p2).F21())
				var _1207_v13 _dafny.Sequence
				_ = _1207_v13
				_1207_v13 = _dafny.SeqOf((_1206_v12).Select((Companion_Default___.SafeIndex((_1195_v4).F25(), _dafny.IntOfUint32((_1206_v12).Cardinality()))).Uint32()).(bool))
				var _1208_v14 _dafny.Sequence
				_ = _1208_v14
				_1208_v14 = _dafny.SeqOf(_1206_v12, _dafny.SeqOf((p2).F21()))
				var _1209_v15 *C2
				_ = _1209_v15
				var _nw200 *C2 = New_C2_()
				_ = _nw200
				_nw200.Ctor__(_dafny.IntOfUint32((_1208_v14).Cardinality()), !((_this).Fm13((p2).F20(), globalState)))
				_1209_v15 = _nw200
				var _1210_v18 _dafny.Map
				_ = _1210_v18
				_1210_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm0(_1191_v0, _1204_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1211_v3 _dafny.Sequence, _1212_v4 *C3) func(_dafny.Int) _dafny.Map {
					return func(_1213_i1 _dafny.Int) _dafny.Map {
						return (_1211_v3).Select((Companion_Default___.SafeIndex((_1212_v4).F25(), _dafny.IntOfUint32((_1211_v3).Cardinality()))).Uint32()).(_dafny.Map)
					}
				})(_1194_v3, _1195_v4))), globalState))
				var _1214_v19 _dafny.Set
				_ = _1214_v19
				_1214_v19 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F20(), (_this).F24()), _1210_v18, _1210_v18)
				var _1215_v20 _dafny.CodePoint
				_ = _1215_v20
				_1215_v20 = _dafny.CodePoint('w')
				var _1216_v21 _dafny.Map
				_ = _1216_v21
				_1216_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter41 := _dafny.Iterate((func() _dafny.Map {
						var _coll39 = _dafny.NewMapBuilder()
						_ = _coll39
						for _iter42 := _dafny.Iterate((_1214_v19).Elements()); ; {
							_compr_39, _ok42 := _iter42()
							if !_ok42 {
								break
							}
							var _1217_v17 _dafny.Map
							_1217_v17 = interface{}(_compr_39).(_dafny.Map)
							if (_1214_v19).Contains(_1217_v17) {
								_coll39.Add(_1217_v17, (_this).F24())
							}
						}
						return _coll39.ToMap()
					}()).Keys().Elements()); ; {
						_compr_38, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _1218_v16 _dafny.Map
						_1218_v16 = interface{}(_compr_38).(_dafny.Map)
						if (func() _dafny.Map {
							var _coll40 = _dafny.NewMapBuilder()
							_ = _coll40
							for _iter43 := _dafny.Iterate((_1214_v19).Elements()); ; {
								_compr_40, _ok43 := _iter43()
								if !_ok43 {
									break
								}
								var _1217_v17 _dafny.Map
								_1217_v17 = interface{}(_compr_40).(_dafny.Map)
								if (_1214_v19).Contains(_1217_v17) {
									_coll40.Add(_1217_v17, (_this).F24())
								}
							}
							return _coll40.ToMap()
						}()).Contains(_1218_v16) {
							_coll38.Add(_1218_v16, p1)
						}
					}
					return _coll38.ToMap()
				}(), _1215_v20)
				var _1219_v22 bool
				_ = _1219_v22
				_1219_v22 = (_this).F24()
				var _1220_v23 _dafny.Map
				_ = _1220_v23
				_1220_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1210_v18, _1219_v22)
				var _1221_v24 _dafny.Map
				_ = _1221_v24
				_1221_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1215_v20)
				_1216_v21 = (_1216_v21).Update(_1220_v23, (func() _dafny.CodePoint {
					if (_1221_v24).Contains((_1195_v4).F26()) {
						return (_1221_v24).Get((_1195_v4).F26()).(_dafny.CodePoint)
					}
					return _1215_v20
				})())
				var _1222_v26 _dafny.Sequence
				_ = _1222_v26
				_1222_v26 = _dafny.SeqOf((_1195_v4).F25())
				var _pat_let_tv24 = _1207_v13
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1207_v13
				_ = _pat_let_tv25
				var _1223_v27 D8
				_ = _1223_v27
				_1223_v27 = Companion_D8_.Create_DC19_(func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter44 := _dafny.Iterate((func(_pat_let26_0 _dafny.Sequence) _dafny.Sequence {
						return func(_1224_dt__update__tmp_h1 _dafny.Sequence) _dafny.Sequence {
							return func(_pat_let27_0 _dafny.Sequence) _dafny.Sequence {
								return func(_1225_dt__update_hcf0_h0 _dafny.Sequence) _dafny.Sequence {
									return _1225_dt__update_hcf0_h0
								}(_pat_let27_0)
							}(_dafny.SeqOf((_dafny.MultiSetFromSeq(_pat_let_tv24)).Cardinality()))
						}(_pat_let26_0)
					}(_1222_v26)).Elements()); ; {
						_compr_41, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _1226_v25 _dafny.Int
						_1226_v25 = interface{}(_compr_41).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains((func(_pat_let28_0 _dafny.Sequence) _dafny.Sequence {
							return func(_1227_dt__update__tmp_h1 _dafny.Sequence) _dafny.Sequence {
								return func(_pat_let29_0 _dafny.Sequence) _dafny.Sequence {
									return func(_1228_dt__update_hcf0_h0 _dafny.Sequence) _dafny.Sequence {
										return _1228_dt__update_hcf0_h0
									}(_pat_let29_0)
								}(_dafny.SeqOf((_dafny.MultiSetFromSeq(_pat_let_tv25)).Cardinality()))
							}(_pat_let28_0)
						}(_1222_v26)), _1226_v25) {
							_coll41.Add((_1226_v25).Times(p0), (p2).F20())
						}
					}
					return _coll41.ToMap()
				}())
				var _1229_v28 _dafny.Sequence
				_ = _1229_v28
				_1229_v28 = _dafny.SeqOf((_1223_v27).Dtor_cf28(), (_1191_v0).Update((_1195_v4).F25(), p0))
				(globalState).F18 = (Companion_Default___.SafeModuloInt((p2).F20(), ((_1229_v28).Select((Companion_Default___.SafeIndex((p2).F20(), _dafny.IntOfUint32((_1229_v28).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())).Cmp((p2).F20()) != 0
				var _1230_v29 D13
				_ = _1230_v29
				_1230_v29 = Companion_D13_.Create_DC37_(p1, (_this).F20())
				var _1231_v30 D4
				_ = _1231_v30
				_1231_v30 = Companion_D4_.Create_DC11_((p2).F20(), (_1230_v29).Dtor_cf49(), (_1195_v4).F26())
				var _1232_v31 _dafny.Map
				_ = _1232_v31
				_1232_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), _1231_v30)
				(globalState).F4 = (_1206_v12).Select((Companion_Default___.SafeIndex(((_1232_v31).Update((_1195_v4).F26(), _1231_v30)).Cardinality(), _dafny.IntOfUint32((_1206_v12).Cardinality()))).Uint32()).(bool)
			} else {
				(globalState).F14 = (p2).F21()
				var _1233_v32 _dafny.Array
				_ = _1233_v32
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_34
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Int = (func(_1234_v4 *C3) func(_dafny.Int) _dafny.Int {
						return func(_1235_i2 _dafny.Int) _dafny.Int {
							return (_1235_i2).Times((_1234_v4).F25())
						}
					})(_1195_v4)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw201 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw201).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw201).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1233_v32 = _nw201
				var _1236_v33 _dafny.Set
				_ = _1236_v33
				_1236_v33 = _dafny.SetOf(_1233_v32)
				var _1237_v34 *C1
				_ = _1237_v34
				var _nw202 *C1 = New_C1_()
				_ = _nw202
				_nw202.Ctor__(_dafny.IntOfInt64(675), _1233_v32, (p2).F20(), (_1195_v4).F26())
				_1237_v34 = _nw202
				var _1238_v35 _dafny.Array
				_ = _1238_v35
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw203
				_1238_v35 = _nw203
				var _1239_v36 _dafny.Sequence
				_ = _1239_v36
				_1239_v36 = _dafny.SeqOf(_1238_v35, _1238_v35)
				var _1240_v37 D3
				_ = _1240_v37
				_1240_v37 = Companion_D3_.Create_DC5_(_1239_v36)
				var _1241_v38 D3
				_ = _1241_v38
				_1241_v38 = Companion_D3_.Create_DC8_(_1240_v37)
				var _1242_v39 _dafny.Map
				_ = _1242_v39
				_1242_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v34, _1241_v38)
				var _1243_v40 _dafny.Set
				_ = _1243_v40
				_1243_v40 = _dafny.SetOf(_1237_v34.F27, _dafny.IntOfInt64(-407), _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
				var _1244_v41 _dafny.Sequence
				_ = _1244_v41
				_1244_v41 = _dafny.SeqOf((_this).F24())
				var _1245_v42 _dafny.Map
				_ = _1245_v42
				_1245_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v39, (_1243_v40).Equals(_dafny.SetOf((_this).F20(), _dafny.IntOfUint32((_1244_v41).Cardinality()), (_1237_v34).Fm1((p2).F20(), _dafny.IntOfInt64(-317), p1, globalState))))
				var _1246_v43 _dafny.Sequence
				_ = _1246_v43
				_1246_v43 = _dafny.UnicodeSeqOfUtf8Bytes("klgw")
				var _rhs234 _dafny.Set = _1236_v33
				_ = _rhs234
				var _rhs235 _dafny.Map = _1245_v42
				_ = _rhs235
				var _rhs236 bool = (_1237_v34).Fm0(_1191_v0, _1246_v43, _1194_v3, globalState)
				_ = _rhs236
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				_1236_v33 = _rhs234
				_1245_v42 = _rhs235
				_lhs214.F4 = _rhs236
				var _1247_v44 _dafny.Map
				_ = _1247_v44
				_1247_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F20(), _1191_v0)
				(globalState).F4 = (p2).Fm0((func() _dafny.Map {
					if (_1247_v44).Contains((p2).F20()) {
						return (_1247_v44).Get((p2).F20()).(_dafny.Map)
					}
					return _1191_v0
				})(), _dafny.UnicodeSeqOfUtf8Bytes("vtqfdj"), _1194_v3, globalState)
				var _1248_v45 _dafny.MultiSet
				_ = _1248_v45
				_1248_v45 = _dafny.MultiSetOf(!((_this).F24()))
				_1248_v45 = _dafny.MultiSetFromSeq(_1244_v41)
				(globalState).F4 = (func() bool {
					if (_1195_v4).F26() {
						return (Companion_Default___.Fm37((p2).F20(), globalState)).IsProperSubsetOf(_dafny.MultiSetOf((_this).F20()))
					}
					return (p2).F21()
				})()
			}
			var _1249_v46 _dafny.Sequence
			_ = _1249_v46
			_1249_v46 = _dafny.UnicodeSeqOfUtf8Bytes("tyhopkq")
			var _1250_v47 D8
			_ = _1250_v47
			_1250_v47 = Companion_D8_.Create_DC19_(_1191_v0)
			var _1251_v48 _dafny.Map
			_ = _1251_v48
			_1251_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1249_v46, true)
			var _1252_v49 _dafny.Set
			_ = _1252_v49
			_1252_v49 = _dafny.SetOf(!((p2).F21()))
			(globalState).F4 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lwcov"), _1249_v46)).Cardinality())).Cmp((((_1250_v47).Dtor_cf28()).Update((_1251_v48).Cardinality(), (_1252_v49).Cardinality())).Cardinality()) > 0
			(globalState).F15 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_this).Fm13((_dafny.Zero).Minus((p2).F20()), globalState) {
					return _1249_v46
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_1253_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))
			})(), _1249_v46)
			var _1254_v50 _dafny.Array
			_ = _1254_v50
			var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw204
			_1254_v50 = _nw204
			var _1255_v51 T0
			_ = _1255_v51
			var _nw205 *C1 = New_C1_()
			_ = _nw205
			_nw205.Ctor__((_this).F20(), (func() _dafny.Array {
				if (_this).F24() {
					return _1254_v50
				}
				return _1254_v50
			})(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ndkax")).Cardinality())).Times(_dafny.IntOfInt64(447)), (p2).F21())
			_1255_v51 = _nw205
			_1255_v51 = _1255_v51
		} else {
			(globalState).F17 = (_dafny.Zero).Minus((p2).F20())
			var _1256_v52 _dafny.Sequence
			_ = _1256_v52
			_1256_v52 = _dafny.UnicodeSeqOfUtf8Bytes("ucx")
			var _1257_v53 _dafny.Sequence
			_ = _1257_v53
			_1257_v53 = _dafny.SeqOf((_this).F20())
			var _1258_v54 _dafny.Map
			_ = _1258_v54
			_1258_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1257_v53).Cardinality()))
			var _1259_v55 _dafny.Map
			_ = _1259_v55
			_1259_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p1)
			var _1260_v56 _dafny.Sequence
			_ = _1260_v56
			_1260_v56 = _dafny.SeqOf(_1259_v55, _1259_v55, _1259_v55, _1259_v55)
			var _1261_v57 _dafny.Sequence
			_ = _1261_v57
			_1261_v57 = _dafny.SeqOf((_this).F24(), !((_this).Fm0(_1258_v54, _1256_v52, _1260_v56, globalState)))
			var _1262_v58 _dafny.Array
			_ = _1262_v58
			var _nwElement0_48 _dafny.Int = _dafny.IntOfInt64(-892)
			_ = _nwElement0_48
			var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(8))
			_ = _nw206
			(_nw206).ArraySet1(_nwElement0_48, 0)
			(_nw206).ArraySet1((_dafny.MultiSetOf(p1, (_this).F21(), (_this).F24(), !(p1), (p2).F21())).Cardinality(), 1)
			(_nw206).ArraySet1(_dafny.IntOfUint32((_1256_v52).Cardinality()), 2)
			(_nw206).ArraySet1((p2).F20(), 3)
			(_nw206).ArraySet1(_dafny.IntOfInt64(120), 4)
			(_nw206).ArraySet1(_dafny.IntOfUint32((_1261_v57).Cardinality()), 5)
			(_nw206).ArraySet1(p0, 6)
			(_nw206).ArraySet1((_this).F20(), 7)
			_1262_v58 = _nw206
			var _1263_v59 *C1
			_ = _1263_v59
			var _nw207 *C1 = New_C1_()
			_ = _nw207
			_nw207.Ctor__((p2).F20(), _1262_v58, (_dafny.Zero).Minus((p2).F20()), (p2).F21())
			_1263_v59 = _nw207
			_1263_v59 = _1263_v59
			var _1264_v60 _dafny.Sequence
			_ = _1264_v60
			_1264_v60 = _dafny.SeqOf(_1261_v57)
			_1261_v57 = _dafny.Companion_Sequence_.Concatenate((_1264_v60).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1264_v60).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf((p2).F21(), (p2).F21()))
			var _1265_v61 _dafny.Array
			_ = _1265_v61
			var _nw208 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
			_ = _nw208
			_1265_v61 = _nw208
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1265_v61), 0))
			_ = _index210
			(_1265_v61).ArraySet1(_1263_v59, (_index210).Int())
			var _1266_v62 *C1
			_ = _1266_v62
			_1266_v62 = _1263_v59
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1265_v61), 0))
			_ = _index211
			(_1265_v61).ArraySet1((_1266_v62), (_index211).Int())
			(globalState).F18 = (p2).F21()
		}
		var _1267_v63 _dafny.Map
		_ = _1267_v63
		_1267_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
		var _1268_i4 _dafny.Int
		_ = _1268_i4
		_1268_i4 = _dafny.Zero
		{
			for ((func() bool {
				if (_1267_v63).Contains(_dafny.IntOfInt64(-116)) {
					return (_1267_v63).Get(_dafny.IntOfInt64(-116)).(bool)
				}
				return (_this).F24()
			})()) && ((p2).F21()) {
				{
					if (_1268_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1268_i4 = (_1268_i4).Plus(_dafny.One)
					var _1269_v64 _dafny.Sequence
					_ = _1269_v64
					_1269_v64 = _dafny.UnicodeSeqOfUtf8Bytes("uqwrp")
					var _1270_v65 _dafny.Sequence
					_ = _1270_v65
					_1270_v65 = _dafny.SeqOf((p2).F20())
					var _1271_v66 _dafny.CodePoint
					_ = _1271_v66
					_1271_v66 = _dafny.CodePoint('x')
					(globalState).F3 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1269_v64, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1270_v65).Cardinality()), _dafny.IntOfUint32((_1269_v64).Cardinality()))).Uint32(), _1271_v66), _1269_v64), _dafny.Companion_Sequence_.Concatenate(_1269_v64, _dafny.Companion_Sequence_.Update(_1269_v64, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1271_v66, _1271_v66, _1271_v66, _1271_v66, _dafny.CodePoint('f'))).Cardinality()), _dafny.IntOfUint32((_1269_v64).Cardinality()))).Uint32(), _1271_v66)))
					var _1272_v67 _dafny.Array
					_ = _1272_v67
					var _len0_35 _dafny.Int = _dafny.IntOfInt64(17)
					_ = _len0_35
					var _nw209 _dafny.Array
					_ = _nw209
					if _len0_35.Cmp(_dafny.Zero) == 0 {
						_nw209 = _dafny.NewArray(_len0_35)
					} else {
						var _init35 func(_dafny.Int) _dafny.Int = func(_1273_i5 _dafny.Int) _dafny.Int {
							return (_1273_i5).Minus((_this).F20())
						}
						_ = _init35
						var _element0_35 = _init35(_dafny.Zero)
						_ = _element0_35
						_nw209 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
						(_nw209).ArraySet1(_element0_35, 0)
						var _nativeLen0_35 = (_len0_35).Int()
						_ = _nativeLen0_35
						for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
							(_nw209).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
						}
					}
					_1272_v67 = _nw209
					var _1274_v68 *C1
					_ = _1274_v68
					var _nw210 *C1 = New_C1_()
					_ = _nw210
					_nw210.Ctor__((_this).F20(), _1272_v67, _dafny.IntOfInt64(284), !((_this).F21()))
					_1274_v68 = _nw210
					if (p2).F21() {
						(globalState).F17 = (p2).F20()
						var _1275_v70 _dafny.Map
						_ = _1275_v70
						_1275_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), (p2).F21())
						var _1276_v71 D2
						_ = _1276_v71
						_1276_v71 = Companion_D2_.Create_DC3_((p2).F20(), !((_this).F24()), p1, (_this).F20(), (_1274_v68).Fm0(func() _dafny.Map {
							var _coll42 = _dafny.NewMapBuilder()
							_ = _coll42
							for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(260), _dafny.IntOfInt64(408))); ; {
								_compr_42, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _1277_v69 _dafny.Int
								_1277_v69 = interface{}(_compr_42).(_dafny.Int)
								if ((_dafny.IntOfInt64(260)).Cmp(_1277_v69) <= 0) && ((_1277_v69).Cmp(_dafny.IntOfInt64(408)) < 0) {
									_coll42.Add(Companion_Default___.SafeModuloInt(_1277_v69, _dafny.IntOfUint32((_1269_v64).Cardinality())), _1274_v68.F27)
								}
							}
							return _coll42.ToMap()
						}(), _1269_v64, _dafny.SeqOf(_1275_v70, _1275_v70), globalState))
						var _1278_v72 _dafny.Sequence
						_ = _1278_v72
						_1278_v72 = _dafny.SeqOf(!((_1276_v71).Dtor_cf7()))
						var _1279_v73 *C0
						_ = _1279_v73
						var _nw211 *C0 = New_C0_()
						_ = _nw211
						_nw211.Ctor__()
						_1279_v73 = _nw211
						var _1280_v74 _dafny.Map
						_ = _1280_v74
						_1280_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1279_v73, p2)
						var _1281_v75 _dafny.Set
						_ = _1281_v75
						_1281_v75 = _dafny.SetOf((func() T0 {
							if (_1280_v74).Contains(_1279_v73) {
								return (_1280_v74).Get(_1279_v73).(T0)
							}
							return p2
						})())
						var _1282_v76 _dafny.Array
						_ = _1282_v76
						var _nwElement0_49 bool = (p2).F21()
						_ = _nwElement0_49
						var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(25))
						_ = _nw212
						(_nw212).ArraySet1(_nwElement0_49, 0)
						(_nw212).ArraySet1((p2).F21(), 1)
						(_nw212).ArraySet1(p1, 2)
						(_nw212).ArraySet1(p1, 3)
						(_nw212).ArraySet1((p2).F21(), 4)
						(_nw212).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1269_v64, Companion_Default___.Fm26((p2).F21(), globalState)), 5)
						(_nw212).ArraySet1(!(!(!((p2).F21()))), 6)
						(_nw212).ArraySet1((_1267_v63).Equals(_1267_v63), 7)
						(_nw212).ArraySet1((p2).F21(), 8)
						(_nw212).ArraySet1((_this).F24(), 9)
						(_nw212).ArraySet1((_this).F21(), 10)
						(_nw212).ArraySet1(p1, 11)
						(_nw212).ArraySet1((p0).Cmp((p2).F20()) >= 0, 12)
						(_nw212).ArraySet1((p2).F21(), 13)
						(_nw212).ArraySet1((p1) || ((_this).F24()), 14)
						(_nw212).ArraySet1((_1278_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.IntOfUint32((_1278_v72).Cardinality()))).Uint32()).(bool), 15)
						(_nw212).ArraySet1((p2).F21(), 16)
						(_nw212).ArraySet1(!((_1281_v75).IsProperSubsetOf(_1281_v75)), 17)
						(_nw212).ArraySet1((_this).F24(), 18)
						(_nw212).ArraySet1((p2).F21(), 19)
						(_nw212).ArraySet1((_1274_v68.F27).Cmp(_dafny.IntOfInt64(-465)) <= 0, 20)
						(_nw212).ArraySet1(!((p2).F21()), 21)
						(_nw212).ArraySet1((func() bool {
							if (p2).F21() {
								return !((p2).F21())
							}
							return (_this).F24()
						})(), 22)
						(_nw212).ArraySet1(true, 23)
						(_nw212).ArraySet1(((_this).F21()) == ((p2).F21()), 24)
						_1282_v76 = _nw212
						var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_1282_v76), 0))
						_ = _index212
						(_1282_v76).ArraySet1((p2).F21(), (_index212).Int())
						var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_1282_v76), 0))
						_ = _index213
						(_1282_v76).ArraySet1(p1, (_index213).Int())
						var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen(((_1274_v68).F28()), 0))
						_ = _index214
						((_1274_v68).F28()).ArraySet1(p0, (_index214).Int())
						var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen(((_1274_v68).F28()), 0))
						_ = _index215
						var _rhs237 _dafny.Int = (_dafny.Zero).Minus(_1274_v68.F27)
						_ = _rhs237
						var _rhs238 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("qfdxlwohg")
						_ = _rhs238
						var _lhs215 _dafny.Array = (_1274_v68).F28()
						_ = _lhs215
						var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen(((_1274_v68).F28()), 0))
						_ = _lhs216
						(_lhs215).ArraySet1(_rhs237, (_lhs216).Int())
						_1269_v64 = _rhs238
						var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_1282_v76), 0))
						_ = _index216
						(_1282_v76).ArraySet1((p2).F21(), (_index216).Int())
						(globalState).F5 = ((p2).F20()).Plus(_dafny.IntOfInt64(86))
					} else {
						var _1283_v77 _dafny.Array
						_ = _1283_v77
						var _nw213 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw213
						_1283_v77 = _nw213
						var _1284_v78 _dafny.Map
						_ = _1284_v78
						_1284_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), _1283_v77)
						var _1285_v79 _dafny.Map
						_ = _1285_v79
						_1285_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1269_v64).Cardinality()), (_1284_v78).Cardinality())
						var _1286_v80 _dafny.MultiSet
						_ = _1286_v80
						_1286_v80 = _dafny.MultiSetOf((_this).F24(), (_this).F24())
						var _1287_v81 _dafny.Sequence
						_ = _1287_v81
						_1287_v81 = _1270_v65
						var _1288_v82 _dafny.Map
						_ = _1288_v82
						_1288_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), (_1274_v68).Fm16(_1286_v80, _1287_v81, globalState))
						var _1289_v83 _dafny.Sequence
						_ = _1289_v83
						_1289_v83 = _dafny.SeqOf(_1288_v82, _1288_v82, _1288_v82, _1288_v82)
						var _1290_v84 _dafny.Map
						_ = _1290_v84
						_1290_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), _1269_v64)
						(globalState).F15 = (func() _dafny.Sequence {
							if !((_1274_v68).Fm0(_1285_v79, _dafny.Companion_Sequence_.Update(_1269_v64, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1269_v64).Cardinality()))).Uint32(), _1271_v66), _1289_v83, globalState)) || ((p2).F21()) {
								return _dafny.UnicodeSeqOfUtf8Bytes("tyibl")
							}
							return (func() _dafny.Sequence {
								if (_1290_v84).Contains((p2).F21()) {
									return (_1290_v84).Get((p2).F21()).(_dafny.Sequence)
								}
								return _1269_v64
							})()
						})()
						var _1291_v85 *C0
						_ = _1291_v85
						var _nw214 *C0 = New_C0_()
						_ = _nw214
						_nw214.Ctor__()
						_1291_v85 = _nw214
						(globalState).F18 = _dafny.Companion_Sequence_.Equal(_1270_v65, _dafny.Companion_Sequence_.Concatenate(_1270_v65, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg61 _dafny.Int) interface{} {
								return coer61(arg61)
							}
						}(func(_1292_i6 _dafny.Int) _dafny.Int {
							return (_this).F20()
						}))))
						var _1293_v86 _dafny.MultiSet
						_ = _1293_v86
						_1293_v86 = _dafny.MultiSetOf((p2).Fm1((p2).F20(), _dafny.IntOfUint32((Companion_Default___.Fm28((p2).F20(), globalState)).Cardinality()), (Companion_D9_.Create_DC23_(p0, (p2).F21())).Dtor_cf32(), globalState), (_dafny.Zero).Minus((p2).Fm1((_1267_v63).Cardinality(), (_dafny.Zero).Minus((p2).F20()), p1, globalState)), _1274_v68.F27, _dafny.IntOfInt64(855))
						var _1294_v87 D10
						_ = _1294_v87
						_1294_v87 = Companion_D10_.Create_DC25_(_1293_v86)
						var _1295_v88 _dafny.Array
						_ = _1295_v88
						var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
						_ = _nw215
						_1295_v88 = _nw215
						var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1295_v88), 0))
						_ = _index217
						(_1295_v88).ArraySet1((func() _dafny.Array {
							if (_this).F21() {
								return _1283_v77
							}
							return _1283_v77
						})(), (_index217).Int())
						var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_1283_v77), 0))
						_ = _index218
						(_1283_v77).ArraySet1((_this).F21(), (_index218).Int())
						var _1296_v89 _dafny.Sequence
						_ = _1296_v89
						_1296_v89 = _dafny.SeqOf(_1285_v79)
						var _1297_v90 _dafny.Map
						_ = _1297_v90
						_1297_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1274_v68).F28(), false)
						var _1298_v91 _dafny.Sequence
						_ = _1298_v91
						_1298_v91 = _dafny.SeqOf(_1283_v77)
						var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1295_v88), 0))
						_ = _index219
						var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_1283_v77), 0))
						_ = _index220
						var _rhs239 bool = (p0).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()))) == 0
						_ = _rhs239
						var _rhs240 D10 = Companion_Default___.Fm39(globalState)
						_ = _rhs240
						var _rhs241 bool = !((((_1291_v85).Fm4((_1286_v80).Cardinality(), _1296_v89, (_dafny.Zero).Minus((_1297_v90).Cardinality()), globalState)) && ((p2).F21())) && ((_this).F24()))
						_ = _rhs241
						var _rhs242 _dafny.Array = (_1298_v91).Select((Companion_Default___.SafeIndex((p2).F20(), _dafny.IntOfUint32((_1298_v91).Cardinality()))).Uint32()).(_dafny.Array)
						_ = _rhs242
						var _rhs243 bool = (p2).F21()
						_ = _rhs243
						var _lhs217 *GlobalState = globalState
						_ = _lhs217
						var _lhs218 *GlobalState = globalState
						_ = _lhs218
						var _lhs219 _dafny.Array = _1295_v88
						_ = _lhs219
						var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1295_v88), 0))
						_ = _lhs220
						var _lhs221 _dafny.Array = _1283_v77
						_ = _lhs221
						var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_1283_v77), 0))
						_ = _lhs222
						_lhs217.F18 = _rhs239
						_1294_v87 = _rhs240
						_lhs218.F3 = _rhs241
						(_lhs219).ArraySet1(_rhs242, (_lhs220).Int())
						(_lhs221).ArraySet1(_rhs243, (_lhs222).Int())
						(globalState).F14 = (p2).F21()
					}
					var _1299_v92 _dafny.MultiSet
					_ = _1299_v92
					_1299_v92 = _dafny.MultiSetOf((_this).F21())
					var _1300_v93 D9
					_ = _1300_v93
					_1300_v93 = Companion_D9_.Create_DC23_(_dafny.IntOfInt64(455), p1)
					var _1301_v94 *C0
					_ = _1301_v94
					var _nw216 *C0 = New_C0_()
					_ = _nw216
					_nw216.Ctor__()
					_1301_v94 = _nw216
					var _1302_v95 _dafny.MultiSet
					_ = _1302_v95
					_1302_v95 = _dafny.MultiSetOf(_1301_v94)
					var _1303_v96 _dafny.Map
					_ = _1303_v96
					_1303_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).F21(), (p2).F21())
					var _1304_v97 _dafny.MultiSet
					_ = _1304_v97
					_1304_v97 = _dafny.MultiSetOf((_dafny.Zero).Minus((p2).F20()), (p2).F20())
					var _1305_v98 _dafny.Array
					_ = _1305_v98
					var _nwElement0_50 bool = false
					_ = _nwElement0_50
					var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(17))
					_ = _nw217
					(_nw217).ArraySet1(_nwElement0_50, 0)
					(_nw217).ArraySet1(p1, 1)
					(_nw217).ArraySet1((p2).F21(), 2)
					(_nw217).ArraySet1((_1299_v92).IsSubsetOf(_1299_v92), 3)
					(_nw217).ArraySet1(!((_1300_v93).Dtor_cf32()), 4)
					(_nw217).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1269_v64, _1269_v64), 5)
					(_nw217).ArraySet1(p1, 6)
					(_nw217).ArraySet1(p1, 7)
					(_nw217).ArraySet1((_1302_v95).Contains(_1301_v94), 8)
					(_nw217).ArraySet1((p2).F21(), 9)
					(_nw217).ArraySet1((_1274_v68).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1274_v68.F27, (p2).F20()), _1269_v64, _dafny.SeqOf(_1303_v96, _1303_v96), globalState), 10)
					(_nw217).ArraySet1((p2).F21(), 11)
					(_nw217).ArraySet1((_this).F21(), 12)
					(_nw217).ArraySet1((_this).F24(), 13)
					(_nw217).ArraySet1(((p2).F20()).Cmp((_1304_v97).Cardinality()) <= 0, 14)
					(_nw217).ArraySet1((_this).F21(), 15)
					(_nw217).ArraySet1(p1, 16)
					_1305_v98 = _nw217
					_1305_v98 = _1305_v98
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		_1267_v63 = (_1267_v63).Update(_dafny.IntOfInt64(605), (p2).F21())
		(globalState).F17 = (p2).F20()
		var _1306_v99 _dafny.Sequence
		_ = _1306_v99
		_1306_v99 = _dafny.SeqOf(false, (_this).F24(), (_this).F24())
		var _1307_v100 _dafny.Set
		_ = _1307_v100
		_1307_v100 = _dafny.SetOf((_this).F20(), (_this).F20(), _dafny.IntOfUint32((_1306_v99).Cardinality()), p0)
		var _1308_v101 *C3
		_ = _1308_v101
		var _nw218 *C3 = New_C3_()
		_ = _nw218
		_nw218.Ctor__(Companion_Default___.SafeDivisionInt(p0, (_this).F20()), false, (_this).F20(), (_1307_v100).IsDisjointFrom(_1307_v100))
		_1308_v101 = _nw218
		var _1309_v102 _dafny.MultiSet
		_ = _1309_v102
		_1309_v102 = _dafny.MultiSetOf(p1)
		var _1310_v104 _dafny.Sequence
		_ = _1310_v104
		_1310_v104 = _dafny.SeqOf(func() _dafny.Map {
			var _coll43 = _dafny.NewMapBuilder()
			_ = _coll43
			for _iter46 := _dafny.Iterate((_1267_v63).Keys().Elements()); ; {
				_compr_43, _ok46 := _iter46()
				if !_ok46 {
					break
				}
				var _1311_v103 _dafny.Int
				_1311_v103 = interface{}(_compr_43).(_dafny.Int)
				if (_1267_v63).Contains(_1311_v103) {
					_coll43.Add(Companion_Default___.SafeModuloInt(_1311_v103, (_1308_v101).F25()), (_1308_v101).F26())
				}
			}
			return _coll43.ToMap()
		}())
		var _1312_v105 _dafny.Sequence
		_ = _1312_v105
		_1312_v105 = _dafny.UnicodeSeqOfUtf8Bytes("rukj")
		var _1313_v106 _dafny.Sequence
		_ = _1313_v106
		_1313_v106 = _dafny.SeqOf(_dafny.IntOfUint32((_1312_v105).Cardinality()))
		var _1314_v107 _dafny.Array
		_ = _1314_v107
		var _nwElement0_51 _dafny.Int = (p2).F20()
		_ = _nwElement0_51
		var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(23))
		_ = _nw219
		(_nw219).ArraySet1(_nwElement0_51, 0)
		(_nw219).ArraySet1((_1308_v101).F25(), 1)
		(_nw219).ArraySet1(((_1308_v101).F25()).Minus(p0), 2)
		(_nw219).ArraySet1((_1309_v102).Cardinality(), 3)
		(_nw219).ArraySet1(_dafny.IntOfUint32((_1310_v104).Cardinality()), 4)
		(_nw219).ArraySet1((_dafny.IntOfInt64(325)).Plus((_1308_v101).F25()), 5)
		(_nw219).ArraySet1(Companion_Default___.SafeModuloInt((_1308_v101).F25(), (_this).Fm1((p2).F20(), (_1308_v101).F25(), (_1308_v101).F26(), globalState)), 6)
		(_nw219).ArraySet1((_this).Fm1(p0, (_1308_v101).F25(), (_this).F24(), globalState), 7)
		(_nw219).ArraySet1(_dafny.IntOfUint32((_1312_v105).Cardinality()), 8)
		(_nw219).ArraySet1((_1308_v101).F25(), 9)
		(_nw219).ArraySet1(p0, 10)
		(_nw219).ArraySet1((p2).Fm1(_dafny.IntOfUint32((_1306_v99).Cardinality()), (_this).F20(), (p2).F21(), globalState), 11)
		(_nw219).ArraySet1((_1308_v101).F25(), 12)
		(_nw219).ArraySet1(p0, 13)
		(_nw219).ArraySet1((_1308_v101).F25(), 14)
		(_nw219).ArraySet1((p2).Fm1(p0, p0, (_this).F24(), globalState), 15)
		(_nw219).ArraySet1((_1308_v101).F25(), 16)
		(_nw219).ArraySet1(_dafny.IntOfUint32((_1313_v106).Cardinality()), 17)
		(_nw219).ArraySet1((_this).F20(), 18)
		(_nw219).ArraySet1((_dafny.IntOfInt64(-384)).Times((_this).F20()), 19)
		(_nw219).ArraySet1((_1308_v101).F25(), 20)
		(_nw219).ArraySet1((p2).F20(), 21)
		(_nw219).ArraySet1((_dafny.Zero).Minus((_this).F20()), 22)
		_1314_v107 = _nw219
		r0 = _1314_v107
		r1 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(500), p0), (_dafny.Zero).Minus((_this).F20()))
		return r0, r1
	}
}
func (_this *C4) F24() bool {
	{
		return _this._f24
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f20 _dafny.Int
	_f21 bool
	F23  bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
	_this.F23 = false
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

func (_this *C5) F20() _dafny.Int {
	return _this._f20
}
func (_this *C5) F21() bool {
	return _this._f21
}
func (_this *C5) Ctor__(f23 bool, f20 _dafny.Int, f21 bool) {
	{
		(_this).F23 = f23
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C5) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((_this).F21())
	}
}
func (_this *C5) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.SetOf((_this).F20(), (_this).F20())).Cardinality()).Plus(((func() _dafny.Map {
			var _coll44 = _dafny.NewMapBuilder()
			_ = _coll44
			for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-451), _dafny.IntOfInt64(-31))); ; {
				_compr_44, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _1315_v0 _dafny.Int
				_1315_v0 = interface{}(_compr_44).(_dafny.Int)
				if ((_dafny.IntOfInt64(-451)).Cmp(_1315_v0) <= 0) && ((_1315_v0).Cmp(_dafny.IntOfInt64(-31)) < 0) {
					_coll44.Add((_1315_v0).Plus((_this).F20()), _this.F23)
				}
			}
			return _coll44.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F21()))).Cardinality())
	}
}
func (_this *C5) Fm5(p0 bool, globalState *GlobalState) bool {
	{
		return _this.F23
	}
}
func (_this *C5) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1316_v0 _dafny.Sequence
		_ = _1316_v0
		_1316_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gjgjhukoy")
		var _1317_v1 _dafny.Map
		_ = _1317_v1
		_1317_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1316_v0).Cardinality()), (_this).F21())
		var _1318_v2 _dafny.MultiSet
		_ = _1318_v2
		_1318_v2 = _dafny.MultiSetOf(_1317_v1)
		var _1319_v3 _dafny.Sequence
		_ = _1319_v3
		_1319_v3 = _dafny.SeqOf(Companion_Default___.Fm6(false, _1318_v2, globalState))
		var _1320_v4 _dafny.Map
		_ = _1320_v4
		_1320_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7((_this).F20(), p0, globalState), p1)
		var _1321_v5 _dafny.Map
		_ = _1321_v5
		_1321_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1319_v3, (p0).Plus((_1320_v4).Cardinality()))
		var _1322_v6 _dafny.Set
		_ = _1322_v6
		_1322_v6 = _dafny.SetOf((_this).Fm5(!(p1), globalState), (_this).F21())
		var _1323_v7 _dafny.Sequence
		_ = _1323_v7
		_1323_v7 = _dafny.SeqOf((_this).F20(), (_dafny.Zero).Minus((_1322_v6).Cardinality()))
		_1321_v5 = (_1321_v5).Update((func() _dafny.Sequence {
			if (_this).F21() {
				return _1319_v3
			}
			return Companion_Default___.Fm8((_this).F20(), (_this).F21(), _1323_v7, _this.F23, globalState)
		})(), _dafny.IntOfInt64(820))
		var _1324_v9 _dafny.Sequence
		_ = _1324_v9
		_1324_v9 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p1))
		var _1325_v10 _dafny.CodePoint
		_ = _1325_v10
		_1325_v10 = _dafny.CodePoint('p')
		var _1326_v11 _dafny.Map
		_ = _1326_v11
		_1326_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfInt64(121))
		var _1327_v12 _dafny.Sequence
		_ = _1327_v12
		_1327_v12 = _dafny.SeqOf((_this).Fm0(_1326_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(384))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg62 _dafny.Int) interface{} {
				return coer62(arg62)
			}
		}((func(_1328_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1329_i0 _dafny.Int) _dafny.CodePoint {
				return _1328_v10
			}
		})(_1325_v10))), _1324_v9, globalState), _this.F23, !(_this.F23))
		var _1330_v13 _dafny.Map
		_ = _1330_v13
		_1330_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, p1)
		var _1331_v14 _dafny.Array
		_ = _1331_v14
		var _nwElement0_52 bool = !((_this).F21())
		_ = _nwElement0_52
		var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(22))
		_ = _nw220
		(_nw220).ArraySet1(_nwElement0_52, 0)
		(_nw220).ArraySet1(p1, 1)
		(_nw220).ArraySet1(p1, 2)
		(_nw220).ArraySet1(p1, 3)
		(_nw220).ArraySet1(!((p1) == (_this.F23)), 4)
		(_nw220).ArraySet1(true, 5)
		(_nw220).ArraySet1(((_this).F20()).Cmp((_this).F20()) <= 0, 6)
		(_nw220).ArraySet1(p1, 7)
		(_nw220).ArraySet1((_this).F21(), 8)
		(_nw220).ArraySet1(false, 9)
		(_nw220).ArraySet1(p1, 10)
		(_nw220).ArraySet1((_this).Fm0(func() _dafny.Map {
			var _coll45 = _dafny.NewMapBuilder()
			_ = _coll45
			for _iter48 := _dafny.Iterate((_1323_v7).Elements()); ; {
				_compr_45, _ok48 := _iter48()
				if !_ok48 {
					break
				}
				var _1332_v8 _dafny.Int
				_1332_v8 = interface{}(_compr_45).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_1323_v7, _1332_v8) {
					_coll45.Add(Companion_Default___.SafeDivisionInt(_1332_v8, p0), (_this).F20())
				}
			}
			return _coll45.ToMap()
		}(), _dafny.UnicodeSeqOfUtf8Bytes("i"), _1324_v9, globalState), 11)
		(_nw220).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1316_v0, _1325_v10), 12)
		(_nw220).ArraySet1(((_this).F20()).Cmp((_this).F20()) >= 0, 13)
		(_nw220).ArraySet1(_this.F23, 14)
		(_nw220).ArraySet1(_this.F23, 15)
		(_nw220).ArraySet1((_1327_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eghu")).Cardinality()), _dafny.IntOfUint32((_1327_v12).Cardinality()))).Uint32()).(bool), 16)
		(_nw220).ArraySet1(_this.F23, 17)
		(_nw220).ArraySet1(((_this).F20()).Cmp((_1330_v13).Cardinality()) == 0, 18)
		(_nw220).ArraySet1((_this).F21(), 19)
		(_nw220).ArraySet1(!((_this).F21()) || (_this.F23), 20)
		(_nw220).ArraySet1((p0).Cmp((_this).F20()) >= 0, 21)
		_1331_v14 = _nw220
		var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
		_ = _index221
		(_1331_v14).ArraySet1(p1, (_index221).Int())
		var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
		_ = _index222
		(_1331_v14).ArraySet1(!(_this.F23), (_index222).Int())
		var _1333_i1 _dafny.Int
		_ = _1333_i1
		_1333_i1 = _dafny.Zero
		{
			for (_dafny.IntOfUint32((_1316_v0).Cardinality())).Cmp((_this).F20()) <= 0 {
				{
					if (_1333_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1333_i1 = (_1333_i1).Plus(_dafny.One)
					var _1334_v15 _dafny.Array
					_ = _1334_v15
					var _len0_36 _dafny.Int = _dafny.IntOfInt64(15)
					_ = _len0_36
					var _nw221 _dafny.Array
					_ = _nw221
					if _len0_36.Cmp(_dafny.Zero) == 0 {
						_nw221 = _dafny.NewArray(_len0_36)
					} else {
						var _init36 func(_dafny.Int) _dafny.Int = (func(_1335_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1336_i2 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_1336_i2, _1335_p0)
							}
						})(p0)
						_ = _init36
						var _element0_36 = _init36(_dafny.Zero)
						_ = _element0_36
						_nw221 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
						(_nw221).ArraySet1(_element0_36, 0)
						var _nativeLen0_36 = (_len0_36).Int()
						_ = _nativeLen0_36
						for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
							(_nw221).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
						}
					}
					_1334_v15 = _nw221
					var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))
					_ = _index223
					(_1334_v15).ArraySet1(p0, (_index223).Int())
					var _1337_v16 _dafny.MultiSet
					_ = _1337_v16
					_1337_v16 = _dafny.MultiSetOf((_this).F20(), (_this).F20())
					var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))
					_ = _index224
					(_1334_v15).ArraySet1((p0).Minus((_1337_v16).Cardinality()), (_index224).Int())
					(globalState).F18 = _this.F23
					var _1338_v17 _dafny.Map
					_ = _1338_v17
					_1338_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1316_v0, (_1334_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))).Int()).(_dafny.Int))
					var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))
					_ = _index225
					(_1334_v15).ArraySet1((((_1338_v17).Merge(_1338_v17)).Merge(_1338_v17)).Cardinality(), (_index225).Int())
					if ((_this).F20()).Cmp((_this).F20()) > 0 {
						var _1339_v18 _dafny.Array
						_ = _1339_v18
						var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
						_ = _nw222
						_1339_v18 = _nw222
						var _1340_v19 _dafny.Map
						_ = _1340_v19
						_1340_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1339_v18, !((_this).F21()) || (_this.F23))
						var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))
						_ = _index226
						var _rhs244 _dafny.Array = _1331_v14
						_ = _rhs244
						var _rhs245 _dafny.Int = _dafny.IntOfInt64(273)
						_ = _rhs245
						var _rhs246 bool = (_this).Fm0(_1326_v11, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(916))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg63 _dafny.Int) interface{} {
								return coer63(arg63)
							}
						}((func(_1341_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1342_i3 _dafny.Int) _dafny.CodePoint {
								return _1341_v10
							}
						})(_1325_v10))), _1316_v0), _1324_v9, globalState)
						_ = _rhs246
						var _rhs247 bool = (func() bool {
							if (_1340_v19).Contains(_1339_v18) {
								return (_1340_v19).Get(_1339_v18).(bool)
							}
							return (_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool)
						})()
						_ = _rhs247
						var _lhs223 _dafny.Array = _1334_v15
						_ = _lhs223
						var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))
						_ = _lhs224
						var _lhs225 *GlobalState = globalState
						_ = _lhs225
						var _lhs226 *GlobalState = globalState
						_ = _lhs226
						_1331_v14 = _rhs244
						(_lhs223).ArraySet1(_rhs245, (_lhs224).Int())
						_lhs225.F18 = _rhs246
						_lhs226.F3 = _rhs247
						var _1343_v20 _dafny.Array
						_ = _1343_v20
						var _len0_37 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_37
						var _nw223 _dafny.Array
						_ = _nw223
						if _len0_37.Cmp(_dafny.Zero) == 0 {
							_nw223 = _dafny.NewArray(_len0_37)
						} else {
							var _init37 func(_dafny.Int) bool = func(_1344_i4 _dafny.Int) bool {
								return (_this).F21()
							}
							_ = _init37
							var _element0_37 = _init37(_dafny.Zero)
							_ = _element0_37
							_nw223 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
							(_nw223).ArraySet1(_element0_37, 0)
							var _nativeLen0_37 = (_len0_37).Int()
							_ = _nativeLen0_37
							for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
								(_nw223).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
							}
						}
						_1343_v20 = _nw223
						var _1345_v21 _dafny.Sequence
						_ = _1345_v21
						_1345_v21 = _dafny.SeqOf(_1343_v20)
						_1345_v21 = _dafny.SeqOf(_1343_v20, _1343_v20)
						(globalState).F18 = (p0).Cmp(_dafny.IntOfUint32((_1327_v12).Cardinality())) > 0
						var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
						_ = _index227
						(_1331_v14).ArraySet1((_this).F21(), (_index227).Int())
						var _1346_v22 _dafny.Sequence
						_ = _1346_v22
						_1346_v22 = _dafny.SeqOf(_1331_v14, _1331_v14)
						var _1347_v23 _dafny.Map
						_ = _1347_v23
						_1347_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1346_v22)
						_1347_v23 = (_1347_v23).Update((_this).F20(), _1346_v22)
					} else {
						(globalState).F4 = (_this).F21()
						var _1348_v24 _dafny.MultiSet
						_ = _1348_v24
						_1348_v24 = _dafny.MultiSetOf(_1331_v14, _1331_v14)
						var _1349_v25 _dafny.Sequence
						_ = _1349_v25
						_1349_v25 = _dafny.SeqOf(_1348_v24)
						var _1350_v26 _dafny.Sequence
						_ = _1350_v26
						_1350_v26 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1323_v7, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1323_v7).Cardinality()))).Uint32(), (_this).F20()))
						var _1351_v27 _dafny.Sequence
						_ = _1351_v27
						_1351_v27 = (_1350_v26).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1350_v26).Cardinality()))).Uint32()).(_dafny.Sequence)
						var _1352_v28 _dafny.Map
						_ = _1352_v28
						_1352_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1334_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))).Int()).(_dafny.Int), _1348_v24)
						var _1353_v29 D2
						_ = _1353_v29
						_1353_v29 = Companion_D2_.Create_DC2_(_1348_v24)
						var _1354_v30 _dafny.Map
						_ = _1354_v30
						_1354_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F20())
						var _1355_v31 _dafny.Array
						_ = _1355_v31
						var _nwElement0_53 _dafny.MultiSet = _1348_v24
						_ = _nwElement0_53
						var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(21))
						_ = _nw224
						(_nw224).ArraySet1(_nwElement0_53, 0)
						(_nw224).ArraySet1(_1348_v24, 1)
						(_nw224).ArraySet1((_1349_v25).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm9(_dafny.IntOfUint32((_1316_v0).Cardinality()), _1351_v27, (_this).F20(), globalState)).Cardinality(), _dafny.IntOfUint32((_1349_v25).Cardinality()))).Uint32()).(_dafny.MultiSet), 2)
						(_nw224).ArraySet1(_1348_v24, 3)
						(_nw224).ArraySet1((_1349_v25).Select((Companion_Default___.SafeIndex((_1334_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1334_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1349_v25).Cardinality()))).Uint32()).(_dafny.MultiSet), 4)
						(_nw224).ArraySet1(_1348_v24, 5)
						(_nw224).ArraySet1((func() _dafny.MultiSet {
							if (_1352_v28).Contains((_this).F20()) {
								return (_1352_v28).Get((_this).F20()).(_dafny.MultiSet)
							}
							return _1348_v24
						})(), 6)
						(_nw224).ArraySet1(_1348_v24, 7)
						(_nw224).ArraySet1(_1348_v24, 8)
						(_nw224).ArraySet1(_1348_v24, 9)
						(_nw224).ArraySet1((_1348_v24).Update(_1331_v14, Companion_Default___.Abs(p0)), 10)
						(_nw224).ArraySet1((_1353_v29).Dtor_cf2(), 11)
						(_nw224).ArraySet1((_dafny.MultiSetOf(_1331_v14)).Difference(_1348_v24), 12)
						(_nw224).ArraySet1(_1348_v24, 13)
						(_nw224).ArraySet1(_1348_v24, 14)
						(_nw224).ArraySet1(_1348_v24, 15)
						(_nw224).ArraySet1((_1348_v24).Update(_1331_v14, Companion_Default___.Abs((_this).Fm1((_this).F20(), (_1354_v30).Cardinality(), true, globalState))), 16)
						(_nw224).ArraySet1((_1348_v24).Intersection(_dafny.MultiSetOf(_1331_v14, _1331_v14)), 17)
						(_nw224).ArraySet1(_1348_v24, 18)
						(_nw224).ArraySet1(_1348_v24, 19)
						(_nw224).ArraySet1(_1348_v24, 20)
						_1355_v31 = _nw224
						var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1355_v31), 0))
						_ = _index228
						(_1355_v31).ArraySet1(_1348_v24, (_index228).Int())
						var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1355_v31), 0))
						_ = _index229
						(_1355_v31).ArraySet1((_1349_v25).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1349_v25).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index229).Int())
						(globalState).F3 = !((_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool))
						(globalState).F18 = !((_this).F21())
						var _1356_v32 *C0
						_ = _1356_v32
						var _nw225 *C0 = New_C0_()
						_ = _nw225
						_nw225.Ctor__()
						_1356_v32 = _nw225
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if _this.F23 {
			var _1357_v33 _dafny.Array
			_ = _1357_v33
			var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw226
			_1357_v33 = _nw226
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _index230
			(_1357_v33).ArraySet1((func() _dafny.Int {
				if (_this).F21() {
					return (_this).F20()
				}
				return (_this).F20()
			})(), (_index230).Int())
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _index231
			(_1357_v33).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()), (_index231).Int())
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _index232
			(_1357_v33).ArraySet1(((_1326_v11).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfInt64(204)))).Cardinality(), (_index232).Int())
			_1325_v10 = _1325_v10
			var _1358_v34 _dafny.Sequence
			_ = _1358_v34
			_1358_v34 = _dafny.SeqOf(_1331_v14, _1331_v14)
			var _1359_v35 D3
			_ = _1359_v35
			_1359_v35 = Companion_D3_.Create_DC5_(_1358_v34)
			var _1360_v36 _dafny.MultiSet
			_ = _1360_v36
			_1360_v36 = _dafny.MultiSetOf(_dafny.CodePoint('q'), _1325_v10, _1325_v10)
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _index233
			var _rhs248 _dafny.Int = (_this).F20()
			_ = _rhs248
			var _rhs249 _dafny.Sequence = ((func() D3 {
				if true {
					return _1359_v35
				}
				return _1359_v35
			})()).Dtor_cf9()
			_ = _rhs249
			var _rhs250 bool = (_dafny.MultiSetOf(_1325_v10)).IsSubsetOf(_1360_v36)
			_ = _rhs250
			var _lhs227 _dafny.Array = _1357_v33
			_ = _lhs227
			var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _lhs228
			(_lhs227).ArraySet1(_rhs248, (_lhs228).Int())
			_1358_v34 = _rhs249
			r1 = _rhs250
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
			_ = _index234
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _index235
			var _rhs251 bool = !((_dafny.IntOfInt64(714)).Cmp((_this).F20()) > 0)
			_ = _rhs251
			var _rhs252 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}(func(_1361_i5 _dafny.Int) _dafny.Int {
				return (_this).F20()
			}))
			_ = _rhs252
			var _rhs253 _dafny.Int = Companion_Default___.SafeModuloInt((p0).Times(_dafny.IntOfUint32((_1327_v12).Cardinality())), p0)
			_ = _rhs253
			var _lhs229 _dafny.Array = _1331_v14
			_ = _lhs229
			var _lhs230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
			_ = _lhs230
			var _lhs231 _dafny.Array = _1357_v33
			_ = _lhs231
			var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1357_v33), 0))
			_ = _lhs232
			(_lhs229).ArraySet1(_rhs251, (_lhs230).Int())
			_1323_v7 = _rhs252
			(_lhs231).ArraySet1(_rhs253, (_lhs232).Int())
		} else {
			var _1362_v37 _dafny.MultiSet
			_ = _1362_v37
			_1362_v37 = _dafny.MultiSetOf(_1331_v14)
			var _1363_v38 D2
			_ = _1363_v38
			_1363_v38 = Companion_D2_.Create_DC2_((_1362_v37).Intersection((_dafny.MultiSetOf(_1331_v14, _1331_v14, _1331_v14, _1331_v14, _1331_v14)).Update(_1331_v14, Companion_Default___.Abs(_dafny.IntOfInt64(438)))))
			_1363_v38 = _1363_v38
			var _1364_v39 _dafny.Sequence
			_ = _1364_v39
			_1364_v39 = _1323_v7
			var _1365_v40 _dafny.Set
			_ = _1365_v40
			_1365_v40 = _dafny.SetOf(_1364_v39)
			var _1366_v41 _dafny.MultiSet
			_ = _1366_v41
			_1366_v41 = _dafny.MultiSetOf(Companion_Default___.Fm10(true, (_this).F21(), _dafny.IntOfUint32((_1316_v0).Cardinality()), globalState))
			(globalState).F5 = ((_1365_v40).Intersection(func() _dafny.Set {
				var _coll46 = _dafny.NewBuilder()
				_ = _coll46
				for _iter49 := _dafny.Iterate((_1366_v41).Elements()); ; {
					_compr_46, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _1367_v42 _dafny.Sequence
					_1367_v42 = interface{}(_compr_46).(_dafny.Sequence)
					if (_1366_v41).Contains(_1367_v42) {
						_coll46.Add(_1367_v42)
					}
				}
				return _coll46.ToSet()
			}())).Cardinality()
			var _1368_v43 *C0
			_ = _1368_v43
			var _nw227 *C0 = New_C0_()
			_ = _nw227
			_nw227.Ctor__()
			_1368_v43 = _nw227
			var _1369_v44 _dafny.Array
			_ = _1369_v44
			var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw228
			_1369_v44 = _nw228
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1369_v44), 0))
			_ = _index236
			(_1369_v44).ArraySet1(p0, (_index236).Int())
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1369_v44), 0))
			_ = _index237
			var _rhs254 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs254
			var _rhs255 _dafny.Int = (_this).F20()
			_ = _rhs255
			var _rhs256 _dafny.Int = ((_this).F20()).Plus(Companion_Default___.SafeModuloInt((_this).F20(), _dafny.IntOfUint32((_1316_v0).Cardinality())))
			_ = _rhs256
			var _lhs233 _dafny.Array = _1369_v44
			_ = _lhs233
			var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1369_v44), 0))
			_ = _lhs234
			var _lhs235 *GlobalState = globalState
			_ = _lhs235
			var _lhs236 *GlobalState = globalState
			_ = _lhs236
			(_lhs233).ArraySet1(_rhs254, (_lhs234).Int())
			_lhs235.F17 = _rhs255
			_lhs236.F17 = _rhs256
			(globalState).F18 = !(false) || (_this.F23)
		}
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1331_v14), 0))); ; {
			_guard_loop_3, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _1370_i6 _dafny.Int
			_1370_i6 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1370_i6).Sign() != -1) && ((_1370_i6).Cmp(_dafny.ArrayLen((_1331_v14), 0)) < 0)) {
				(_1331_v14).ArraySet1(!(p1), (_1370_i6).Int())
			}
		}
		var _1371_v45 _dafny.MultiSet
		_ = _1371_v45
		_1371_v45 = _dafny.MultiSetOf(_this.F23, _this.F23)
		var _1372_v46 D2
		_ = _1372_v46
		_1372_v46 = Companion_D2_.Create_DC3_(_dafny.IntOfUint32((_1327_v12).Cardinality()), !(_1371_v45).Contains(_this.F23), (_this).F21(), p0, ((_this).Fm5((_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool), globalState)) || ((_this).F21()))
		var _source16 D2 = _1372_v46
		_ = _source16
		if _source16.Is_DC3() {
			var _1373___mcc_h0 _dafny.Int = _source16.Get_().(D2_DC3).Cf3
			_ = _1373___mcc_h0
			var _1374___mcc_h1 bool = _source16.Get_().(D2_DC3).Cf4
			_ = _1374___mcc_h1
			var _1375___mcc_h2 bool = _source16.Get_().(D2_DC3).Cf5
			_ = _1375___mcc_h2
			var _1376___mcc_h3 _dafny.Int = _source16.Get_().(D2_DC3).Cf6
			_ = _1376___mcc_h3
			var _1377___mcc_h4 bool = _source16.Get_().(D2_DC3).Cf7
			_ = _1377___mcc_h4
			var _1378_cf7 bool = _1377___mcc_h4
			_ = _1378_cf7
			var _1379_cf6 _dafny.Int = _1376___mcc_h3
			_ = _1379_cf6
			var _1380_cf5 bool = _1375___mcc_h2
			_ = _1380_cf5
			var _1381_cf4 bool = _1374___mcc_h1
			_ = _1381_cf4
			var _1382_cf3 _dafny.Int = _1373___mcc_h0
			_ = _1382_cf3
			var _1383_v47 _dafny.Array
			_ = _1383_v47
			var _nw229 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(6))
			_ = _nw229
			_1383_v47 = _nw229
			var _1384_v48 _dafny.MultiSet
			_ = _1384_v48
			_1384_v48 = _dafny.MultiSetOf(_1331_v14, (Companion_D4_.Create_DC9_(_1331_v14)).Dtor_cf14(), _1331_v14)
			var _1385_v49 D2
			_ = _1385_v49
			_1385_v49 = Companion_D2_.Create_DC2_(_1384_v48)
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1383_v47), 0))
			_ = _index238
			(_1383_v47).ArraySet1(_1385_v49, (_index238).Int())
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1383_v47), 0))
			_ = _index239
			(_1383_v47).ArraySet1(_1385_v49, (_index239).Int())
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))
			_ = _index240
			(_1331_v14).ArraySet1((_this).F21(), (_index240).Int())
			(globalState).F17 = p0
			var _1386_v50 _dafny.Sequence
			_ = _1386_v50
			_1386_v50 = _dafny.SeqOf(_1316_v0)
			var _1387_v51 _dafny.Sequence
			_ = _1387_v51
			_1387_v51 = _dafny.SeqOf(_1316_v0, _1316_v0, _dafny.Companion_Sequence_.Concatenate(_1316_v0, _1316_v0), _1316_v0, (_1386_v50).Select((Companion_Default___.SafeIndex(_1379_cf6, _dafny.IntOfUint32((_1386_v50).Cardinality()))).Uint32()).(_dafny.Sequence))
			_1386_v50 = _1387_v51
		} else if _source16.Is_DC2() {
			var _1388___mcc_h5 _dafny.MultiSet = _source16.Get_().(D2_DC2).Cf2
			_ = _1388___mcc_h5
			var _1389_cf2 _dafny.MultiSet = _1388___mcc_h5
			_ = _1389_cf2
			(globalState).F5 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(696))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_1390_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1391_i7 _dafny.Int) _dafny.CodePoint {
					return _1390_v10
				}
			})(_1325_v10)))).Cardinality())
			(_this).F23 = (_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool)
			(globalState).F17 = (_this).F20()
			var _1392_v52 _dafny.Array
			_ = _1392_v52
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_38
			var _nw230 _dafny.Array
			_ = _nw230
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw230 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Sequence = (func(_1393_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1394_i8 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_1393_v0, _1393_v0)
					}
				})(_1316_v0)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw230 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw230).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw230).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1392_v52 = _nw230
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1392_v52), 0))
			_ = _index241
			(_1392_v52).ArraySet1(_1316_v0, (_index241).Int())
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1392_v52), 0))
			_ = _index242
			(_1392_v52).ArraySet1(_1316_v0, (_index242).Int())
		} else {
			var _1395___mcc_h6 D2 = _source16.Get_().(D2_DC4).Cf8
			_ = _1395___mcc_h6
			var _1396_cf8 D2 = _1395___mcc_h6
			_ = _1396_cf8
			var _1397_v53 *C0
			_ = _1397_v53
			var _nw231 *C0 = New_C0_()
			_ = _nw231
			_nw231.Ctor__()
			_1397_v53 = _nw231
			var _1398_v54 _dafny.Array
			_ = _1398_v54
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_39
			var _nw232 _dafny.Array
			_ = _nw232
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw232 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Sequence = (func(_1399_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_1400_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-378))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg66 _dafny.Int) interface{} {
								return coer66(arg66)
							}
						}((func(_1401_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1402_i10 _dafny.Int) _dafny.CodePoint {
								return _1401_v10
							}
						})(_1399_v10)))
					}
				})(_1325_v10)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw232 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw232).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw232).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1398_v54 = _nw232
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1398_v54), 0))
			_ = _index243
			(_1398_v54).ArraySet1(_1316_v0, (_index243).Int())
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1398_v54), 0))
			_ = _index244
			(_1398_v54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("iby"), _1316_v0), (_index244).Int())
			if (p0).Cmp((_this).F20()) > 0 {
				(globalState).F17 = (_this).Fm1((_this).F20(), (_this).F20(), (func() bool {
					if p1 {
						return p1
					}
					return (_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool)
				})(), globalState)
				var _1403_v55 _dafny.Map
				_ = _1403_v55
				_1403_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool), _1398_v54)
				var _1404_v56 _dafny.MultiSet
				_ = _1404_v56
				_1404_v56 = _dafny.MultiSetOf((_this).F20(), p0, (_this).F20())
				_1398_v54 = (func() _dafny.Array {
					if (_1403_v55).Contains((_1404_v56).IsProperSubsetOf(_1404_v56)) {
						return (_1403_v55).Get((_1404_v56).IsProperSubsetOf(_1404_v56)).(_dafny.Array)
					}
					return _1398_v54
				})()
				var _1405_v57 _dafny.Map
				_ = _1405_v57
				_1405_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1372_v46).Dtor_cf6(), _1404_v56)
				_1405_v57 = (_1405_v57).Update(_dafny.IntOfInt64(797), Companion_Default___.Fm11(_dafny.IntOfUint32((_1323_v7).Cardinality()), (_this).F21(), globalState))
				(_this).F23 = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm12(p0, _1325_v10, (_this).F21(), (func() _dafny.Int {
					if (_1326_v11).Contains((_this).F20()) {
						return (_1326_v11).Get((_this).F20()).(_dafny.Int)
					}
					return (_this).F20()
				})(), globalState), _dafny.Companion_Sequence_.Concatenate((_1398_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1398_v54), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("avxghk")))
				_1323_v7 = (func() _dafny.Sequence {
					if p1 {
						return _1323_v7
					}
					return _1323_v7
				})()
			} else {
				var _1406_v58 *C0
				_ = _1406_v58
				var _nw233 *C0 = New_C0_()
				_ = _nw233
				_nw233.Ctor__()
				_1406_v58 = _nw233
				var _1407_v59 _dafny.Map
				_ = _1407_v59
				_1407_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _1317_v1)
				var _1408_v60 _dafny.Sequence
				_ = _1408_v60
				_1408_v60 = _dafny.SeqOf(_1326_v11)
				var _1409_v61 _dafny.Map
				_ = _1409_v61
				_1409_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC3_(p0, _this.F23, (_this).F21(), p0, true), (_1397_v53).Fm4(p0, _1408_v60, (_this).F20(), globalState))
				var _1410_v62 _dafny.Sequence
				_ = _1410_v62
				_1410_v62 = _dafny.SeqOf(_1331_v14)
				var _1411_v63 D3
				_ = _1411_v63
				_1411_v63 = Companion_D3_.Create_DC6_((_this).F20())
				var _1412_v64 _dafny.Set
				_ = _1412_v64
				_1412_v64 = _dafny.SetOf(p0, (_this).F20(), _dafny.IntOfUint32(((_1398_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1398_v54), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				var _1413_v65 D4
				_ = _1413_v65
				_1413_v65 = Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_1327_v12).Cardinality()), (_this).F21(), p1)
				var _1414_v66 _dafny.Array
				_ = _1414_v66
				var _nwElement0_54 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _nwElement0_54
				var _nw234 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(25))
				_ = _nw234
				(_nw234).ArraySet1(_nwElement0_54, 0)
				(_nw234).ArraySet1((_1407_v59).Cardinality(), 1)
				(_nw234).ArraySet1(p0, 2)
				(_nw234).ArraySet1(p0, 3)
				(_nw234).ArraySet1(_dafny.IntOfUint32((_1327_v12).Cardinality()), 4)
				(_nw234).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v46, (_this).F21())).Merge(_1409_v61)).Cardinality(), 5)
				(_nw234).ArraySet1(_dafny.IntOfUint32((_1410_v62).Cardinality()), 6)
				(_nw234).ArraySet1(p0, 7)
				(_nw234).ArraySet1((p0).Times(_dafny.IntOfInt64(418)), 8)
				(_nw234).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_this).F20()), 9)
				(_nw234).ArraySet1(((_this).F20()).Minus((_1411_v63).Dtor_cf10()), 10)
				(_nw234).ArraySet1((_1412_v64).Cardinality(), 11)
				(_nw234).ArraySet1((_dafny.Zero).Minus(p0), 12)
				(_nw234).ArraySet1((_dafny.Zero).Minus(((_this).F20()).Times(p0)), 13)
				(_nw234).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm12((_this).F20(), _dafny.CodePoint('n'), (_1397_v53).Fm3(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(431))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1415_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1416_i11 _dafny.Int) _dafny.CodePoint {
						return _1415_v10
					}
				})(_1325_v10)))).Cardinality()), false, globalState), (_this).Fm1(_dafny.IntOfInt64(-961), _dafny.IntOfInt64(442), (_this).F21(), globalState), globalState)).Cardinality()), 14)
				(_nw234).ArraySet1(_dafny.IntOfInt64(390), 15)
				(_nw234).ArraySet1((_this).F20(), 16)
				(_nw234).ArraySet1(_dafny.IntOfUint32((_1323_v7).Cardinality()), 17)
				(_nw234).ArraySet1((_this).F20(), 18)
				(_nw234).ArraySet1((p0).Plus((_1413_v65).Dtor_cf17()), 19)
				(_nw234).ArraySet1(p0, 20)
				(_nw234).ArraySet1((_this).F20(), 21)
				(_nw234).ArraySet1((_this).F20(), 22)
				(_nw234).ArraySet1(p0, 23)
				(_nw234).ArraySet1(p0, 24)
				_1414_v66 = _nw234
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_1414_v66), 0))
				_ = _index245
				(_1414_v66).ArraySet1((_dafny.Zero).Minus(p0), (_index245).Int())
				var _1417_v67 _dafny.MultiSet
				_ = _1417_v67
				_1417_v67 = _dafny.MultiSetOf((_1330_v13).Cardinality())
				var _1418_v68 _dafny.Sequence
				_ = _1418_v68
				_1418_v68 = _dafny.SeqOf(_dafny.MultiSetOf(p0), _1417_v67, _1417_v67)
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_1414_v66), 0))
				_ = _index246
				(_1414_v66).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(498))).Uint32(), func(coer68 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_1419_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_1420_i12 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(_1419_p0)
					}
				})(p0))), _1418_v68)).Cardinality())), (_index246).Int())
				var _1421_v69 *C0
				_ = _1421_v69
				var _nw235 *C0 = New_C0_()
				_ = _nw235
				_nw235.Ctor__()
				_1421_v69 = _nw235
				var _1422_v70 *C0
				_ = _1422_v70
				var _nw236 *C0 = New_C0_()
				_ = _nw236
				_nw236.Ctor__()
				_1422_v70 = _nw236
				_1414_v66 = _1414_v66
			}
			(globalState).F5 = (func() _dafny.Map {
				var _coll47 = _dafny.NewMapBuilder()
				_ = _coll47
				for _iter51 := _dafny.Iterate((_dafny.SeqOf(_1323_v7)).Elements()); ; {
					_compr_47, _ok51 := _iter51()
					if !_ok51 {
						break
					}
					var _1423_v71 _dafny.Sequence
					_1423_v71 = interface{}(_compr_47).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1323_v7), _1423_v71) {
						_coll47.Add(_1423_v71, p1)
					}
				}
				return _coll47.ToMap()
			}()).Cardinality()
		}
		var _1424_v72 _dafny.Map
		_ = _1424_v72
		_1424_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1331_v14, (_this).F21())
		var _1425_v73 *C0
		_ = _1425_v73
		var _nw237 *C0 = New_C0_()
		_ = _nw237
		_nw237.Ctor__()
		_1425_v73 = _nw237
		var _1426_v74 _dafny.Set
		_ = _1426_v74
		_1426_v74 = _dafny.SetOf(_1425_v73)
		var _1427_v75 T0
		_ = _1427_v75
		var _nw238 *C4 = New_C4_()
		_ = _nw238
		_nw238.Ctor__(p1, _dafny.IntOfInt64(471), true)
		_1427_v75 = _nw238
		var _1428_v76 _dafny.Map
		_ = _1428_v76
		_1428_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _1427_v75)
		var _1429_v77 _dafny.MultiSet
		_ = _1429_v77
		_1429_v77 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0), (_this).F20())
		var _1430_v78 _dafny.Map
		_ = _1430_v78
		_1430_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _1429_v77)
		var _1431_v79 _dafny.Map
		_ = _1431_v79
		_1431_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1427_v75, p0)
		var _1432_v80 _dafny.Map
		_ = _1432_v80
		_1432_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _1371_v45)
		var _1433_v81 _dafny.Array
		_ = _1433_v81
		var _nwElement0_55 _dafny.Int = (_this).F20()
		_ = _nwElement0_55
		var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(16))
		_ = _nw239
		(_nw239).ArraySet1(_nwElement0_55, 0)
		(_nw239).ArraySet1(p0, 1)
		(_nw239).ArraySet1((_1426_v74).Cardinality(), 2)
		(_nw239).ArraySet1(_dafny.IntOfInt64(815), 3)
		(_nw239).ArraySet1(p0, 4)
		(_nw239).ArraySet1((_1428_v76).Cardinality(), 5)
		(_nw239).ArraySet1((_this).F20(), 6)
		(_nw239).ArraySet1((_1427_v75).F20(), 7)
		(_nw239).ArraySet1((_this).F20(), 8)
		(_nw239).ArraySet1((_1427_v75).F20(), 9)
		(_nw239).ArraySet1((_this).F20(), 10)
		(_nw239).ArraySet1(p0, 11)
		(_nw239).ArraySet1(((func() _dafny.MultiSet {
			if (_1430_v78).Contains((_1431_v79).Cardinality()) {
				return (_1430_v78).Get((_1431_v79).Cardinality()).(_dafny.MultiSet)
			}
			return _1429_v77
		})()).Cardinality(), 12)
		(_nw239).ArraySet1((_dafny.Zero).Minus((_1371_v45).Cardinality()), 13)
		(_nw239).ArraySet1((_1326_v11).Cardinality(), 14)
		(_nw239).ArraySet1((_1432_v80).Cardinality(), 15)
		_1433_v81 = _nw239
		var _1434_v82 _dafny.Sequence
		_ = _1434_v82
		_1434_v82 = _dafny.SeqOf(_1433_v81, _1433_v81)
		var _1435_v83 _dafny.Map
		_ = _1435_v83
		_1435_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1424_v72).Cardinality(), _dafny.Companion_Sequence_.Update(_1434_v82, (Companion_Default___.SafeIndex((_1323_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.IntOfUint32((_1323_v7).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1434_v82).Cardinality()))).Uint32(), _1433_v81))
		r0 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1435_v83).Contains(p0) {
				return (_1435_v83).Get(p0).(_dafny.Sequence)
			}
			return _1434_v82
		})(), _1434_v82)
		r1 = (_1331_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1331_v14), 0))).Int()).(bool)
		return r0, r1
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f20 _dafny.Int
	_f21 bool
	_f22 _dafny.Sequence
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f20 = _dafny.Zero
	_this._f21 = false
	_this._f22 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F20() _dafny.Int {
	return _this._f20
}
func (_this *C6) F21() bool {
	return _this._f21
}
func (_this *C6) Ctor__(f22 _dafny.Sequence, f20 _dafny.Int, f21 bool) {
	{
		(_this)._f22 = f22
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C6) Fm0(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F21()
	}
}
func (_this *C6) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_this).F20())).Minus((_this).F20())
	}
}
func (_this *C6) M0(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1436_i0 _dafny.Int
		_ = _1436_i0
		_1436_i0 = _dafny.Zero
		{
			for (_this).F21() {
				{
					if (_1436_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1436_i0 = (_1436_i0).Plus(_dafny.One)
					var _1437_v0 _dafny.Sequence
					_ = _1437_v0
					_1437_v0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg69 _dafny.Int) interface{} {
							return coer69(arg69)
						}
					}((func(_1438_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1439_i1 _dafny.Int) _dafny.Int {
							return _1438_p0
						}
					})(p0)))
					var _1440_v1 _dafny.MultiSet
					_ = _1440_v1
					_1440_v1 = _dafny.MultiSetOf(p0)
					var _1441_v2 _dafny.Sequence
					_ = _1441_v2
					_1441_v2 = _dafny.SeqOf(_1440_v1, _1440_v1, _1440_v1)
					var _1442_v3 _dafny.MultiSet
					_ = _1442_v3
					_1442_v3 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_1441_v2).Cardinality()))
					var _1443_v4 _dafny.Sequence
					_ = _1443_v4
					_1443_v4 = _dafny.SeqOf((func() _dafny.Int {
						if (_1442_v3).Contains(p0) {
							return (_1442_v3).Multiplicity(p0)
						}
						return (_this).F20()
					})())
					var _1444_v5 _dafny.Map
					_ = _1444_v5
					_1444_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((true) == ((_this).F21()), (_1443_v4))
					var _1445_v6 _dafny.Map
					_ = _1445_v6
					_1445_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfInt64(-218))
					var _1446_v7 _dafny.Map
					_ = _1446_v7
					_1446_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
					var _1447_v8 bool
					_ = _1447_v8
					_1447_v8 = (_this).F21()
					var _1448_v9 _dafny.Sequence
					_ = _1448_v9
					_1448_v9 = _dafny.SeqOf(_1446_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_1447_v8)))
					_1444_v5 = (_1444_v5).Update((_this).Fm0(_1445_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(787))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg70 _dafny.Int) interface{} {
							return coer70(arg70)
						}
					}(func(_1449_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					})), _1448_v9, globalState), _1443_v4)
					var _source17 _dafny.Sequence = _1437_v0
					_ = _source17
					var _1450___mcc_h0 _dafny.Sequence = _source17
					_ = _1450___mcc_h0
					var _1451_cf0 _dafny.Sequence = _1450___mcc_h0
					_ = _1451_cf0
					var _1452_v10 _dafny.Set
					_ = _1452_v10
					_1452_v10 = _dafny.SetOf(p1, (_this).F21())
					(globalState).F3 = !(_dafny.SetOf((_this).F21())).Equals(_1452_v10)
					var _1453_v11 _dafny.MultiSet
					_ = _1453_v11
					_1453_v11 = _dafny.MultiSetOf((_this).F21())
					var _1454_v12 _dafny.Array
					_ = _1454_v12
					var _nwElement0_56 _dafny.Int = p0
					_ = _nwElement0_56
					var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(8))
					_ = _nw240
					(_nw240).ArraySet1(_nwElement0_56, 0)
					(_nw240).ArraySet1((_1453_v11).Cardinality(), 1)
					(_nw240).ArraySet1((_this).Fm1(p0, _dafny.IntOfInt64(-743), p1, globalState), 2)
					(_nw240).ArraySet1(p0, 3)
					(_nw240).ArraySet1((p0).Times(p0), 4)
					(_nw240).ArraySet1(_dafny.IntOfInt64(296), 5)
					(_nw240).ArraySet1((_1451_cf0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F20()), _dafny.IntOfUint32((_1451_cf0).Cardinality()))).Uint32()).(_dafny.Int), 6)
					(_nw240).ArraySet1((p0).Times(p0), 7)
					_1454_v12 = _nw240
					var _1455_v13 _dafny.Map
					_ = _1455_v13
					_1455_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v4, _1454_v12)
					_1454_v12 = (func() _dafny.Array {
						if (_1455_v13).Contains((func() _dafny.Sequence {
							if (_this).F21() {
								return _1437_v0
							}
							return _1437_v0
						})()) {
							return (_1455_v13).Get((func() _dafny.Sequence {
								if (_this).F21() {
									return _1437_v0
								}
								return _1437_v0
							})()).(_dafny.Array)
						}
						return _1454_v12
					})()
					var _1456_v14 _dafny.Array
					_ = _1456_v14
					var _len0_40 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_40
					var _nw241 _dafny.Array
					_ = _nw241
					if _len0_40.Cmp(_dafny.Zero) == 0 {
						_nw241 = _dafny.NewArray(_len0_40)
					} else {
						var _init40 func(_dafny.Int) bool = (func(_1457_v8 bool) func(_dafny.Int) bool {
							return func(_1458_i3 _dafny.Int) bool {
								return _1457_v8
							}
						})(_1447_v8)
						_ = _init40
						var _element0_40 = _init40(_dafny.Zero)
						_ = _element0_40
						_nw241 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
						(_nw241).ArraySet1(_element0_40, 0)
						var _nativeLen0_40 = (_len0_40).Int()
						_ = _nativeLen0_40
						for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
							(_nw241).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
						}
					}
					_1456_v14 = _nw241
					var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1456_v14), 0))
					_ = _index247
					(_1456_v14).ArraySet1(_1447_v8, (_index247).Int())
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1456_v14), 0))
					_ = _index248
					var _rhs257 bool = _1447_v8
					_ = _rhs257
					var _rhs258 _dafny.Int = (_1452_v10).Cardinality()
					_ = _rhs258
					var _rhs259 bool = ((func() bool {
						if (_this).F21() {
							return p1
						}
						return _1447_v8
					})())
					_ = _rhs259
					var _rhs260 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32(((_this).F22()).Cardinality())), _1443_v4), _1451_cf0)
					_ = _rhs260
					var _rhs261 bool = p1
					_ = _rhs261
					var _lhs237 _dafny.Array = _1456_v14
					_ = _lhs237
					var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1456_v14), 0))
					_ = _lhs238
					var _lhs239 *GlobalState = globalState
					_ = _lhs239
					var _lhs240 *GlobalState = globalState
					_ = _lhs240
					var _lhs241 *GlobalState = globalState
					_ = _lhs241
					(_lhs237).ArraySet1(_rhs257, (_lhs238).Int())
					_lhs239.F5 = _rhs258
					_lhs240.F14 = _rhs259
					_1451_cf0 = _rhs260
					_lhs241.F14 = _rhs261
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1454_v12), 0))
					_ = _index249
					(_1454_v12).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_this).Fm1(p0, (_this).F20(), p1, globalState)), (_index249).Int())
					var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1454_v12), 0))
					_ = _index250
					(_1454_v12).ArraySet1(_dafny.IntOfInt64(157), (_index250).Int())
					_1446_v7 = (_1446_v7).Update(!((_this).F21()), p1)
					var _1459_v15 _dafny.Array
					_ = _1459_v15
					var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
					_ = _nw242
					_1459_v15 = _nw242
					var _1460_v16 _dafny.Map
					_ = _1460_v16
					_1460_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1459_v15)
					_1460_v16 = (_1460_v16).Update((_this).F21(), _1459_v15)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _hi4 _dafny.Int = p0
		_ = _hi4
		for _1461_i4 := (_dafny.Zero).Minus(_dafny.IntOfInt64(-898)); _1461_i4.Cmp(_hi4) < 0; _1461_i4 = _1461_i4.Plus(_dafny.One) {
			var _1462_v17 _dafny.CodePoint
			_ = _1462_v17
			_1462_v17 = _dafny.CodePoint('x')
			var _1463_v18 _dafny.Map
			_ = _1463_v18
			_1463_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1462_v17, Companion_Default___.Fm2(!((_this).F21()), (_this).F20(), p1, globalState))
			var _1464_v19 bool
			_ = _1464_v19
			_1464_v19 = (_this).F21()
			_1463_v18 = (_1463_v18).Update(_1462_v17, _1464_v19)
			(globalState).F5 = (_this).F20()
			var _1465_v20 _dafny.Array
			_ = _1465_v20
			var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw243
			_1465_v20 = _nw243
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1465_v20), 0))
			_ = _index251
			(_1465_v20).ArraySet1((p0).Times(_dafny.IntOfUint32(((_this).F22()).Cardinality())), (_index251).Int())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1465_v20), 0))
			_ = _index252
			var _rhs262 _dafny.Int = (_dafny.Zero).Minus((_this).F20())
			_ = _rhs262
			var _rhs263 _dafny.Sequence = (_this).F22()
			_ = _rhs263
			var _lhs242 _dafny.Array = _1465_v20
			_ = _lhs242
			var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1465_v20), 0))
			_ = _lhs243
			var _lhs244 *GlobalState = globalState
			_ = _lhs244
			(_lhs242).ArraySet1(_rhs262, (_lhs243).Int())
			_lhs244.F15 = _rhs263
			var _1466_v21 _dafny.Sequence
			_ = _1466_v21
			_1466_v21 = _dafny.SeqOf(p0)
			var _source18 _dafny.Sequence = _1466_v21
			_ = _source18
			var _1467___mcc_h1 _dafny.Sequence = _source18
			_ = _1467___mcc_h1
			var _1468_cf0 _dafny.Sequence = _1467___mcc_h1
			_ = _1468_cf0
			(globalState).F18 = (_this).F21()
			var _1469_v22 *C0
			_ = _1469_v22
			var _nw244 *C0 = New_C0_()
			_ = _nw244
			_nw244.Ctor__()
			_1469_v22 = _nw244
			var _1470_v23 _dafny.Sequence
			_ = _1470_v23
			_1470_v23 = _1468_cf0
			_1470_v23 = _1470_v23
			var _1471_v24 _dafny.Map
			_ = _1471_v24
			_1471_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1464_v19, p1)
			var _1472_v25 _dafny.Map
			_ = _1472_v25
			_1472_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
			var _1473_v26 _dafny.Map
			_ = _1473_v26
			_1473_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1465_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1465_v20), 0))).Int()).(_dafny.Int))
			var _1474_v27 _dafny.Sequence
			_ = _1474_v27
			_1474_v27 = _dafny.SeqOf(_1472_v25)
			_1471_v24 = (_1471_v24).Update(_1464_v19, (func() bool {
				if (func() bool {
					if (_1472_v25).Contains((_this).Fm0(_1473_v26, (_this).F22(), _1474_v27, globalState)) {
						return (_1472_v25).Get((_this).Fm0(_1473_v26, (_this).F22(), _1474_v27, globalState)).(bool)
					}
					return p1
				})() {
					return !((_this).F21())
				}
				return (_this).F21()
			})())
		}
		var _1475_v28 _dafny.Array
		_ = _1475_v28
		var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw245
		_1475_v28 = _nw245
		var _1476_v29 _dafny.Array
		_ = _1476_v29
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_41
		var _nw246 _dafny.Array
		_ = _nw246
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw246 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) _dafny.Int = (func(_1477_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1478_i5 _dafny.Int) _dafny.Int {
					return (_1478_i5).Plus(_1477_p0)
				}
			})(p0)
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw246 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw246).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw246).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		_1476_v29 = _nw246
		var _1479_v30 T0
		_ = _1479_v30
		var _nw247 *C1 = New_C1_()
		_ = _nw247
		_nw247.Ctor__((_this).F20(), _1476_v29, (_this).Fm1((_this).F20(), p0, p1, globalState), false)
		_1479_v30 = _nw247
		var _1480_v31 _dafny.Map
		_ = _1480_v31
		_1480_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v30, (_this).F21())
		var _1481_v32 _dafny.Map
		_ = _1481_v32
		_1481_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1480_v31, _1475_v28)
		var _1482_v33 _dafny.Map
		_ = _1482_v33
		_1482_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1479_v30).F20())
		var _1483_v34 _dafny.Map
		_ = _1483_v34
		_1483_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)
		var _1484_v35 _dafny.Sequence
		_ = _1484_v35
		_1484_v35 = _dafny.SeqOf(_1483_v34)
		var _1485_v36 _dafny.Array
		_ = _1485_v36
		var _nwElement0_57 _dafny.Array = _1475_v28
		_ = _nwElement0_57
		var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(13))
		_ = _nw248
		(_nw248).ArraySet1(_nwElement0_57, 0)
		(_nw248).ArraySet1(_1475_v28, 1)
		(_nw248).ArraySet1(_1475_v28, 2)
		(_nw248).ArraySet1(_1475_v28, 3)
		(_nw248).ArraySet1(_1475_v28, 4)
		(_nw248).ArraySet1(_1475_v28, 5)
		(_nw248).ArraySet1((func() _dafny.Array {
			if p1 {
				return _1475_v28
			}
			return _1475_v28
		})(), 6)
		(_nw248).ArraySet1(_1475_v28, 7)
		(_nw248).ArraySet1((func() _dafny.Array {
			if (_1481_v32).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v30, (_this).Fm0(_1482_v33, (_this).F22(), _1484_v35, globalState))) {
				return (_1481_v32).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v30, (_this).Fm0(_1482_v33, (_this).F22(), _1484_v35, globalState))).(_dafny.Array)
			}
			return _1475_v28
		})(), 8)
		(_nw248).ArraySet1(_1475_v28, 9)
		(_nw248).ArraySet1(_1475_v28, 10)
		(_nw248).ArraySet1(_1475_v28, 11)
		(_nw248).ArraySet1(_1475_v28, 12)
		_1485_v36 = _nw248
		var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1485_v36), 0))
		_ = _index253
		(_1485_v36).ArraySet1(_1475_v28, (_index253).Int())
		var _1486_v37 _dafny.CodePoint
		_ = _1486_v37
		_1486_v37 = _dafny.CodePoint('i')
		var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1485_v36), 0))
		_ = _index254
		var _rhs264 _dafny.Array = _1475_v28
		_ = _rhs264
		var _rhs265 _dafny.Int = (func() _dafny.Int {
			if false {
				return (_dafny.IntOfUint32((Companion_Default___.Fm12((_1479_v30).F20(), _1486_v37, p1, (_this).F20(), globalState)).Cardinality())).Plus((_1479_v30).F20())
			}
			return Companion_Default___.SafeModuloInt((_1479_v30).F20(), p0)
		})()
		_ = _rhs265
		var _rhs266 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(469))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg71 _dafny.Int) interface{} {
				return coer71(arg71)
			}
		}((func(_1487_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1488_i6 _dafny.Int) _dafny.CodePoint {
				return _1487_v37
			}
		})(_1486_v37))), (func() _dafny.CodePoint {
			if p1 {
				return _1486_v37
			}
			return _1486_v37
		})())
		_ = _rhs266
		var _lhs245 _dafny.Array = _1485_v36
		_ = _lhs245
		var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1485_v36), 0))
		_ = _lhs246
		var _lhs247 *GlobalState = globalState
		_ = _lhs247
		var _lhs248 *GlobalState = globalState
		_ = _lhs248
		(_lhs245).ArraySet1(_rhs264, (_lhs246).Int())
		_lhs247.F17 = _rhs265
		_lhs248.F4 = _rhs266
		_1479_v30 = _1479_v30
		var _hi5 _dafny.Int = _dafny.IntOfInt64(585)
		_ = _hi5
		for _1489_i7 := (func() _dafny.Int {
			if (_1479_v30).F21() {
				return (_this).F20()
			}
			return (func() _dafny.Int {
				if (_1482_v33).Contains((_1479_v30).F20()) {
					return (_1482_v33).Get((_1479_v30).F20()).(_dafny.Int)
				}
				return (_this).F20()
			})()
		})(); _1489_i7.Cmp(_hi5) < 0; _1489_i7 = _1489_i7.Plus(_dafny.One) {
			var _1490_v38 bool
			_ = _1490_v38
			_1490_v38 = p1
			_1490_v38 = _1490_v38
			var _1491_v39 *C1
			_ = _1491_v39
			var _nw249 *C1 = New_C1_()
			_ = _nw249
			_nw249.Ctor__((_1479_v30).F20(), _1476_v29, p0, (_1479_v30).F21())
			_1491_v39 = _nw249
			_1491_v39 = _1491_v39
			var _1492_v40 _dafny.Set
			_ = _1492_v40
			_1492_v40 = _dafny.SetOf((_1483_v34).Cardinality(), (_this).F20())
			var _1493_v41 _dafny.Sequence
			_ = _1493_v41
			_1493_v41 = _dafny.SeqOf((_1492_v40).Difference(_1492_v40))
			_1493_v41 = _1493_v41
			if (_1479_v30).F21() {
				(_1491_v39).F27 = p0
				var _1494_v42 D8
				_ = _1494_v42
				_1494_v42 = Companion_D8_.Create_DC21_(_1491_v39.F27)
				var _1495_v43 _dafny.MultiSet
				_ = _1495_v43
				_1495_v43 = _dafny.MultiSetOf(((_1494_v42).Dtor_cf29()).Cmp((_this).F20()) > 0)
				_1495_v43 = (_1495_v43).Union(_1495_v43)
				var _1496_v44 _dafny.Sequence
				_ = _1496_v44
				_1496_v44 = _dafny.SeqOf(_dafny.IntOfUint32(((_this).F22()).Cardinality()))
				_1496_v44 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}(func(_1497_i8 _dafny.Int) _dafny.Int {
					return (_this).F20()
				})), _dafny.Companion_Sequence_.Concatenate(_1496_v44, _1496_v44))
				(globalState).F15 = (_this).F22()
				_1476_v29 = (_1491_v39).F28()
			} else {
				var _1498_v45 _dafny.Sequence
				_ = _1498_v45
				_1498_v45 = _dafny.SeqOf(_1491_v39.F27, (_1479_v30).F20())
				var _1499_v47 _dafny.Set
				_ = _1499_v47
				_1499_v47 = _dafny.SetOf(_1492_v40, _1492_v40, _1492_v40, func() _dafny.Set {
					var _coll48 = _dafny.NewBuilder()
					_ = _coll48
					for _iter52 := _dafny.Iterate((_dafny.MultiSetFromSeq(_1498_v45)).Elements()); ; {
						_compr_48, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1500_v46 _dafny.Int
						_1500_v46 = interface{}(_compr_48).(_dafny.Int)
						if (_dafny.MultiSetFromSeq(_1498_v45)).Contains(_1500_v46) {
							_coll48.Add((_1500_v46).Times((_dafny.MultiSetOf(!(false))).Cardinality()))
						}
					}
					return _coll48.ToSet()
				}(), _1492_v40)
				var _rhs267 _dafny.Int = (_1479_v30).F20()
				_ = _rhs267
				var _rhs268 bool = (_this).F21()
				_ = _rhs268
				var _rhs269 _dafny.Set = (_1499_v47).Difference(_1499_v47)
				_ = _rhs269
				var _lhs249 *GlobalState = globalState
				_ = _lhs249
				var _lhs250 *GlobalState = globalState
				_ = _lhs250
				_lhs249.F17 = _rhs267
				_lhs250.F4 = _rhs268
				_1499_v47 = _rhs269
				(globalState).F5 = p0
				var _1501_v48 D13
				_ = _1501_v48
				_1501_v48 = Companion_D13_.Create_DC37_((_1479_v30).F21(), (_this).F20())
				(globalState).F17 = ((_1479_v30).F20()).Plus((_1501_v48).Dtor_cf50())
				var _1502_v49 _dafny.Array
				_ = _1502_v49
				var _nw250 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw250
				_1502_v49 = _nw250
				var _1503_v50 _dafny.Set
				_ = _1503_v50
				_1503_v50 = _dafny.SetOf((_this).F21())
				var _1504_v51 _dafny.Sequence
				_ = _1504_v51
				_1504_v51 = _dafny.SeqOf(_1503_v50)
				var _1505_v52 D3
				_ = _1505_v52
				_1505_v52 = Companion_D3_.Create_DC7_(_1504_v51, p1)
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1502_v49), 0))
				_ = _index255
				(_1502_v49).ArraySet1((_1505_v52).Dtor_cf12(), (_index255).Int())
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1502_v49), 0))
				_ = _index256
				(_1502_v49).ArraySet1(p1, (_index256).Int())
				var _1506_v53 _dafny.Array
				_ = _1506_v53
				_1506_v53 = _1476_v29
				var _1507_v54 _dafny.Set
				_ = _1507_v54
				_1507_v54 = _dafny.SetOf(_1506_v53)
				var _1508_v55 _dafny.Sequence
				_ = _1508_v55
				_1508_v55 = _dafny.SeqOf((_1479_v30).F21())
				var _1509_v56 _dafny.Map
				_ = _1509_v56
				_1509_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _1508_v55)
				var _1510_v57 _dafny.Map
				_ = _1510_v57
				_1510_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v54, (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1509_v56).Contains((_1479_v30).F21()) {
						return (_1509_v56).Get((_1479_v30).F21()).(_dafny.Sequence)
					}
					return _1508_v55
				})()).Cardinality())))
				_1510_v57 = _1510_v57
			}
		}
		var _1511_v58 _dafny.Set
		_ = _1511_v58
		_1511_v58 = _dafny.SetOf(_1476_v29)
		var _hi6 _dafny.Int = (_1511_v58).Cardinality()
		_ = _hi6
		for _1512_i9 := Companion_Default___.SafeModuloInt((_this).F20(), p0); _1512_i9.Cmp(_hi6) < 0; _1512_i9 = _1512_i9.Plus(_dafny.One) {
			var _1513_v59 _dafny.Map
			_ = _1513_v59
			_1513_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F21())
			var _1514_v60 _dafny.Map
			_ = _1514_v60
			_1514_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _1513_v59)
			var _1515_v61 D16
			_ = _1515_v61
			_1515_v61 = Companion_D16_.Create_DC42_(Companion_Default___.Fm40((_this).F21(), (_1479_v30).F21(), (_1479_v30).F20(), (_1514_v60).Cardinality(), globalState))
			(globalState).F17 = ((_1515_v61).Dtor_cf54()).Cardinality()
			(globalState).F15 = _dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22())
			_1483_v34 = _1483_v34
			if !((_dafny.Companion_Sequence_.Contains((_this).F22(), _1486_v37)) == (!((_1479_v30).F21()))) {
				(globalState).F17 = (_this).F20()
				(globalState).F3 = true
				(globalState).F4 = (_this).Fm0(_1482_v33, _dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()), _1484_v35, globalState)
				var _1516_v62 _dafny.Array
				_ = _1516_v62
				var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw251
				_1516_v62 = _nw251
				var _1517_v63 _dafny.Sequence
				_ = _1517_v63
				_1517_v63 = _dafny.SeqOf((_1479_v30).F21())
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_1516_v62), 0))
				_ = _index257
				(_1516_v62).ArraySet1(_1517_v63, (_index257).Int())
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_1516_v62), 0))
				_ = _index258
				var _rhs270 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1517_v63, _dafny.SeqOf(p1, (_this).F21(), (_1479_v30).F21(), (_1479_v30).F21()))
				_ = _rhs270
				var _rhs271 _dafny.Int = (func() _dafny.Int {
					if (_1482_v33).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eopxabkhd")).Cardinality())) {
						return (_1482_v33).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eopxabkhd")).Cardinality())).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if p1 {
							return _dafny.IntOfInt64(255)
						}
						return (_1479_v30).F20()
					})()
				})()
				_ = _rhs271
				var _lhs251 _dafny.Array = _1516_v62
				_ = _lhs251
				var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_1516_v62), 0))
				_ = _lhs252
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				(_lhs251).ArraySet1(_rhs270, (_lhs252).Int())
				_lhs253.F5 = _rhs271
				var _1518_v64 _dafny.Array
				_ = _1518_v64
				var _nwElement0_58 _dafny.Sequence = (_this).F22()
				_ = _nwElement0_58
				var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(19))
				_ = _nw252
				(_nw252).ArraySet1(_nwElement0_58, 0)
				(_nw252).ArraySet1((_this).F22(), 1)
				(_nw252).ArraySet1((_this).F22(), 2)
				(_nw252).ArraySet1((_this).F22(), 3)
				(_nw252).ArraySet1((_this).F22(), 4)
				(_nw252).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("diwso"), 5)
				(_nw252).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xkkwmspk"), 6)
				(_nw252).ArraySet1((_this).F22(), 7)
				(_nw252).ArraySet1((_this).F22(), 8)
				(_nw252).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1519_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1520_i10 _dafny.Int) _dafny.CodePoint {
						return _1519_v37
					}
				})(_1486_v37))), 9)
				(_nw252).ArraySet1((_this).F22(), 10)
				(_nw252).ArraySet1((_this).F22(), 11)
				(_nw252).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-284))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1521_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1522_i11 _dafny.Int) _dafny.CodePoint {
						return _1521_v37
					}
				})(_1486_v37))), 12)
				(_nw252).ArraySet1((_this).F22(), 13)
				(_nw252).ArraySet1((_this).F22(), 14)
				(_nw252).ArraySet1((_this).F22(), 15)
				(_nw252).ArraySet1((_this).F22(), 16)
				(_nw252).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ikkiep"), 17)
				(_nw252).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yvy"), 18)
				_1518_v64 = _nw252
				var _1523_v65 _dafny.Map
				_ = _1523_v65
				_1523_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D11_.Create_DC28_(_1518_v64), (_1479_v30).F21())
				var _1524_v66 _dafny.Sequence
				_ = _1524_v66
				_1524_v66 = _dafny.SeqOf(_1523_v65)
				var _1525_v67 D2
				_ = _1525_v67
				_1525_v67 = Companion_D2_.Create_DC3_((_1479_v30).F20(), (_this).F21(), (_1479_v30).F21(), (_1479_v30).F20(), (_1479_v30).F21())
				var _1526_v68 _dafny.Sequence
				_ = _1526_v68
				_1526_v68 = _dafny.SeqOf(_1517_v63)
				var _1527_v69 _dafny.Array
				_ = _1527_v69
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_42
				var _nw253 _dafny.Array
				_ = _nw253
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw253 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = (func(_1528_p1 bool) func(_dafny.Int) bool {
						return func(_1529_i12 _dafny.Int) bool {
							return _1528_p1
						}
					})(p1)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw253 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw253).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw253).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1527_v69 = _nw253
				var _1530_v70 _dafny.Set
				_ = _1530_v70
				_1530_v70 = _dafny.SetOf((_1479_v30).F20(), (_1479_v30).F20())
				var _1531_v71 _dafny.Map
				_ = _1531_v71
				_1531_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1527_v69, _1530_v70)
				var _1532_v72 _dafny.Set
				_ = _1532_v72
				_1532_v72 = _dafny.SetOf(_dafny.SeqOf((func() bool {
					if (_1483_v34).Contains((_1479_v30).F21()) {
						return (_1483_v34).Get((_1479_v30).F21()).(bool)
					}
					return !((_1525_v67).Dtor_cf4())
				})()), (_1526_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1517_v63, (Companion_Default___.SafeIndex((_1531_v71).Cardinality(), _dafny.IntOfUint32((_1517_v63).Cardinality()))).Uint32(), (_1479_v30).F21())).Cardinality()), _dafny.IntOfUint32((_1526_v68).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _1533_v73 _dafny.MultiSet
				_ = _1533_v73
				_1533_v73 = _dafny.MultiSetOf(!((_1479_v30).F21()))
				var _1534_v74 D4
				_ = _1534_v74
				_1534_v74 = Companion_D4_.Create_DC10_(_1533_v73, (_this).F21())
				var _1535_v75 _dafny.Sequence
				_ = _1535_v75
				_1535_v75 = _dafny.SeqOf((_1479_v30).F20())
				var _rhs272 _dafny.Int = _dafny.IntOfInt64(-116)
				_ = _rhs272
				var _rhs273 _dafny.Int = (_1479_v30).Fm1((_dafny.Zero).Minus(p0), ((_1479_v30).F20()).Times((_1479_v30).F20()), (_1534_v74).Dtor_cf16(), globalState)
				_ = _rhs273
				var _rhs274 _dafny.Sequence = _1524_v66
				_ = _rhs274
				var _rhs275 _dafny.Set = _1532_v72
				_ = _rhs275
				var _rhs276 _dafny.Int = Companion_Default___.SafeDivisionInt((_1479_v30).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1535_v75, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1535_v75).Cardinality()))).Uint32(), _1512_i9), _dafny.SeqOf((_this).F20()))).Cardinality()))
				_ = _rhs276
				var _lhs254 *GlobalState = globalState
				_ = _lhs254
				var _lhs255 *GlobalState = globalState
				_ = _lhs255
				var _lhs256 *GlobalState = globalState
				_ = _lhs256
				_lhs254.F17 = _rhs272
				_lhs255.F17 = _rhs273
				_1524_v66 = _rhs274
				_1532_v72 = _rhs275
				_lhs256.F17 = _rhs276
			} else {
				_1513_v59 = (_1513_v59).Update(p0, true)
				(globalState).F18 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()), _1486_v37)
				var _1536_v76 _dafny.Array
				_ = _1536_v76
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw254
				_1536_v76 = _nw254
				var _1537_v77 D11
				_ = _1537_v77
				_1537_v77 = Companion_D11_.Create_DC28_(_1536_v76)
				_1537_v77 = Companion_D11_.Create_DC28_(_1536_v76)
				var _1538_v78 _dafny.Array
				_ = _1538_v78
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw255
				_1538_v78 = _nw255
				var _1539_v79 _dafny.MultiSet
				_ = _1539_v79
				_1539_v79 = _dafny.MultiSetOf((_1479_v30).F21(), (_1479_v30).F21())
				var _1540_v81 D11
				_ = _1540_v81
				_1540_v81 = Companion_D11_.Create_DC30_(false, _1539_v79)
				var _1541_v82 _dafny.Map
				_ = _1541_v82
				_1541_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1479_v30).F20(), _1540_v81)
				var _1542_v83 _dafny.Array
				_ = _1542_v83
				var _nwElement0_59 bool = (_this).F21()
				_ = _nwElement0_59
				var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(13))
				_ = _nw256
				(_nw256).ArraySet1(_nwElement0_59, 0)
				(_nw256).ArraySet1((_1479_v30).F21(), 1)
				(_nw256).ArraySet1(p1, 2)
				(_nw256).ArraySet1((_this).Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1539_v79).Cardinality(), (_1479_v30).F20()), (_this).F22(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(452))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1543_v34 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1544_i13 _dafny.Int) _dafny.Map {
						return _1543_v34
					}
				})(_1483_v34))), globalState), 3)
				(_nw256).ArraySet1((_1479_v30).F21(), 4)
				(_nw256).ArraySet1((_this).Fm0(_1482_v33, (_this).F22(), _1484_v35, globalState), 5)
				(_nw256).ArraySet1((_this).F21(), 6)
				(_nw256).ArraySet1((func() bool {
					if (_1513_v59).Contains(_dafny.IntOfUint32(((_this).F22()).Cardinality())) {
						return (_1513_v59).Get(_dafny.IntOfUint32(((_this).F22()).Cardinality())).(bool)
					}
					return (_1479_v30).F21()
				})(), 7)
				(_nw256).ArraySet1((_1479_v30).F21(), 8)
				(_nw256).ArraySet1((_this).F21(), 9)
				(_nw256).ArraySet1((_1479_v30).F21(), 10)
				(_nw256).ArraySet1(true, 11)
				(_nw256).ArraySet1((_this).Fm0(func() _dafny.Map {
					var _coll49 = _dafny.NewMapBuilder()
					_ = _coll49
					for _iter53 := _dafny.Iterate((_1541_v82).Keys().Elements()); ; {
						_compr_49, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1545_v80 _dafny.Int
						_1545_v80 = interface{}(_compr_49).(_dafny.Int)
						if (_1541_v82).Contains(_1545_v80) {
							_coll49.Add(Companion_Default___.SafeModuloInt(_1545_v80, (_1479_v30).F20()), p0)
						}
					}
					return _coll49.ToMap()
				}(), (_this).F22(), _1484_v35, globalState), 12)
				_1542_v83 = _nw256
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1538_v78), 0))
				_ = _index259
				(_1538_v78).ArraySet1(_1542_v83, (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1538_v78), 0))
				_ = _index260
				(_1538_v78).ArraySet1(_1542_v83, (_index260).Int())
				(globalState).F5 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(787), (_this).F20())
			}
		}
		var _1546_v84 _dafny.Sequence
		_ = _1546_v84
		_1546_v84 = _dafny.SeqOf(_1476_v29)
		var _1547_v85 D17
		_ = _1547_v85
		_1547_v85 = Companion_D17_.Create_DC46_(_1546_v84)
		r0 = _dafny.Companion_Sequence_.Concatenate((_1547_v85).Dtor_cf62(), (_1547_v85).Dtor_cf62())
		var _1548_v86 D6
		_ = _1548_v86
		_1548_v86 = Companion_D6_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-363))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}((func(_1549_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1550_i14 _dafny.Int) _dafny.CodePoint {
				return _1549_v37
			}
		})(_1486_v37))))
		r1 = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jhqis"), _dafny.Companion_Sequence_.Concatenate((_1548_v86).Dtor_cf21(), (_this).F22()))
		return r0, r1
	}
}
func (_this *C6) F22() _dafny.Sequence {
	{
		return _this._f22
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
