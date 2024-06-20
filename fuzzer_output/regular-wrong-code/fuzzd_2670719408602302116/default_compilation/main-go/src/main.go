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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(!(!(!(!(true))))) || ((Companion_D2_.Create_DC6_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(797))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()))).Cardinality()), true, !(true))).Dtor_cf8()), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("fqkb"), _dafny.UnicodeSeqOfUtf8Bytes("ocpjiygo")), (_dafny.IntOfInt64(-380)).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('n'))).Cardinality()))) != 0)
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-386))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), _dafny.UnicodeSeqOfUtf8Bytes("xk"))).Cardinality())
	}))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC3_((func() D0 {
		if true {
			return Companion_D0_.Create_DC1_((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(189))
		}
		return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())
	})())
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(84)
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.CodePoint('y')), _dafny.CodePoint('l'))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(false), _dafny.SeqOf(!(false), true, true, true)), _dafny.CodePoint('k'))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("g")
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("elk"), _dafny.UnicodeSeqOfUtf8Bytes("f")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("f"), _dafny.UnicodeSeqOfUtf8Bytes("jflao")))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('e')
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jos"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _dafny.UnicodeSeqOfUtf8Bytes("ilgan")))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(376)).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("akm"), !(true))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	var _source0 _dafny.Sequence = _dafny.SeqOf(((Companion_D5_.Create_DC13_((Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Dtor_cf19())).Dtor_cf19()).Cardinality(), _dafny.IntOfInt64(382))
	_ = _source0
	var _5___mcc_h0 _dafny.Sequence = _source0
	_ = _5___mcc_h0
	var _6_cf24 _dafny.Sequence = _5___mcc_h0
	_ = _6_cf24
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (Companion_D9_.Create_DC21_(false, _dafny.SeqOf(true), true)).Dtor_cf27())
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(216), _dafny.IntOfInt64(372)))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-715))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("unujxwuqx"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(37))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_8_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	}))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 D3, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	return (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-978), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))).Cardinality())))).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()), _dafny.IntOfInt64(312)))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source1 D19 = Companion_D19_.Create_DC39_(false, _dafny.IntOfInt64(106))
	_ = _source1
	if _source1.Is_DC39() {
		var _10___mcc_h0 bool = _source1.Get_().(D19_DC39).Cf58
		_ = _10___mcc_h0
		var _11___mcc_h1 _dafny.Int = _source1.Get_().(D19_DC39).Cf59
		_ = _11___mcc_h1
		var _12_cf59 _dafny.Int = _11___mcc_h1
		_ = _12_cf59
		var _13_cf58 bool = _10___mcc_h0
		_ = _13_cf58
		return _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(506), (_dafny.MultiSetOf(_12_cf59)).Cardinality(), _12_cf59, (_dafny.MultiSetOf(_12_cf59, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_cf58, _12_cf59)).Cardinality())).Cardinality(), _12_cf59)).Cardinality(), _12_cf59, (_dafny.Zero).Minus(_12_cf59)))
	} else if _source1.Is_DC38() {
		var _14___mcc_h2 _dafny.Map = _source1.Get_().(D19_DC38).Cf57
		_ = _14___mcc_h2
		var _15_cf57 _dafny.Map = _14___mcc_h2
		_ = _15_cf57
		return (_dafny.MultiSetOf((_dafny.MultiSetOf(true, true, true, !(false), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kh")).Cardinality()), _dafny.IntOfInt64(-878), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-273))).Cardinality())).Difference(_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(939))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-120))).Cardinality())))
	} else {
		var _16___mcc_h3 D19 = _source1.Get_().(D19_DC40).Cf60
		_ = _16___mcc_h3
		var _17_cf60 D19 = _16___mcc_h3
		_ = _17_cf60
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(687)))).Difference(_dafny.MultiSetOf((Companion_D0_.Create_DC0_((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-72)))).Cardinality())).Dtor_cf0(), _dafny.IntOfInt64(919)))
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Union(_dafny.SetOf(_dafny.IntOfInt64(564), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(2), _dafny.IntOfInt64(844))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(2)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(844)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_18_v0, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(277))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_19_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qthfycdk")).Cardinality())
	})), !(!(false)))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("heoi")).Cardinality()))).Cardinality()), _dafny.IntOfInt64(387), (_dafny.Zero).Minus((_dafny.MultiSetOf(!(false))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(false, false, false)
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 D9, p2 bool, p3 _dafny.Int, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC11_(!((_dafny.MultiSetOf(_dafny.IntOfInt64(-407))).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(708), _dafny.IntOfInt64(789), _dafny.IntOfInt64(-937)))), (_dafny.IntOfInt64(210)).Cmp(_dafny.IntOfInt64(-759)) != 0, _dafny.IntOfInt64(842), _dafny.Companion_Sequence_.Concatenate((Companion_D10_.Create_DC24_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-675))).Cardinality(), true, _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('y')))).Dtor_cf33(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source2 D10 = Companion_D10_.Create_DC23_(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sisutg")))
	_ = _source2
	if _source2.Is_DC24() {
		var _21___mcc_h0 _dafny.Int = _source2.Get_().(D10_DC24).Cf31
		_ = _21___mcc_h0
		var _22___mcc_h1 bool = _source2.Get_().(D10_DC24).Cf32
		_ = _22___mcc_h1
		var _23___mcc_h2 _dafny.Sequence = _source2.Get_().(D10_DC24).Cf33
		_ = _23___mcc_h2
		var _24_cf33 _dafny.Sequence = _23___mcc_h2
		_ = _24_cf33
		var _25_cf32 bool = _22___mcc_h1
		_ = _25_cf32
		var _26_cf31 _dafny.Int = _21___mcc_h0
		_ = _26_cf31
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_cf32, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_cf32, _26_cf31))
	} else {
		var _27___mcc_h3 _dafny.MultiSet = _source2.Get_().(D10_DC23).Cf30
		_ = _27___mcc_h3
		var _28_cf30 _dafny.MultiSet = _27___mcc_h3
		_ = _28_cf30
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC7_(_dafny.CodePoint('n'), true, _dafny.IntOfInt64(-255))).Dtor_cf10(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(275)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-616))))
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true), !(false)), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true)), false))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(997))
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, true)).Difference(_dafny.SetOf(!(true), !(true), true, false, (Companion_D19_.Create_DC39_(false, _dafny.IntOfInt64(831))).Dtor_cf58()))).Union((_dafny.SetOf(false)).Intersection(_dafny.SetOf(!(true), !(false))))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 _dafny.Set, p2 bool, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC23_(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hx"), _dafny.UnicodeSeqOfUtf8Bytes("kueav")))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetOf(false, false), _dafny.MultiSetOf(false))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _29_v0 _dafny.MultiSet
			_29_v0 = interface{}(_compr_1).(_dafny.MultiSet)
			if (_dafny.MultiSetOf(_dafny.MultiSetOf(false, false), _dafny.MultiSetOf(false))).Contains(_29_v0) {
				_coll1.Add(_29_v0, _dafny.IntOfInt64(737))
			}
		}
		return _coll1.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, true), _dafny.IntOfInt64(790)))).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _30_v1 _dafny.MultiSet
			_30_v1 = interface{}(_compr_2).(_dafny.MultiSet)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Contains(_30_v1) {
				_coll2.Add(_30_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
			}
		}
		return _coll2.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-725), _dafny.IntOfInt64(483)))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D12 {
	var _source3 _dafny.Sequence = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(-707))).Cardinality())
	_ = _source3
	var _31___mcc_h0 _dafny.Sequence = _source3
	_ = _31___mcc_h0
	var _32_cf24 _dafny.Sequence = _31___mcc_h0
	_ = _32_cf24
	return Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(48)))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, globalState *GlobalState) D2 {
	var _source4 D7 = Companion_D7_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gndcte")).Cardinality()), !(true)))
	_ = _source4
	if _source4.Is_DC18() {
		if true {
			return Companion_D2_.Create_DC7_(_dafny.CodePoint('j'), !(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mhiixh")).Cardinality()))
		} else {
			return Companion_D2_.Create_DC7_(_dafny.CodePoint('p'), true, _dafny.IntOfInt64(-733))
		}
	} else {
		var _33___mcc_h0 _dafny.Map = _source4.Get_().(D7_DC17).Cf23
		_ = _33___mcc_h0
		var _34_cf23 _dafny.Map = _33___mcc_h0
		_ = _34_cf23
		return Companion_D2_.Create_DC7_(_dafny.CodePoint('v'), false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(true, true)).Cardinality()))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("fvh"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sfclqlf"), _dafny.UnicodeSeqOfUtf8Bytes("vxt")))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source5 D12 = (func() D12 {
		if true {
			return Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())))
		}
		return Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-332)))
	})()
	_ = _source5
	if _source5.Is_DC27() {
		var _35___mcc_h0 _dafny.Int = _source5.Get_().(D12_DC27).Cf36
		_ = _35___mcc_h0
		var _36___mcc_h1 _dafny.Int = _source5.Get_().(D12_DC27).Cf37
		_ = _36___mcc_h1
		var _37___mcc_h2 _dafny.Int = _source5.Get_().(D12_DC27).Cf38
		_ = _37___mcc_h2
		var _38___mcc_h3 _dafny.Int = _source5.Get_().(D12_DC27).Cf39
		_ = _38___mcc_h3
		var _39_cf39 _dafny.Int = _38___mcc_h3
		_ = _39_cf39
		var _40_cf38 _dafny.Int = _37___mcc_h2
		_ = _40_cf38
		var _41_cf37 _dafny.Int = _36___mcc_h1
		_ = _41_cf37
		var _42_cf36 _dafny.Int = _35___mcc_h0
		_ = _42_cf36
		return (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fup"), _dafny.UnicodeSeqOfUtf8Bytes("xmp"))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _43_v1 _dafny.Sequence
					_43_v1 = interface{}(_compr_4).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fup"), _dafny.UnicodeSeqOfUtf8Bytes("xmp"))).Contains(_43_v1) {
						_coll4.Add(_43_v1, true)
					}
				}
				return _coll4.ToMap()
			}()).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _44_v0 _dafny.Sequence
				_44_v0 = interface{}(_compr_3).(_dafny.Sequence)
				if (func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fup"), _dafny.UnicodeSeqOfUtf8Bytes("xmp"))).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _43_v1 _dafny.Sequence
						_43_v1 = interface{}(_compr_5).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fup"), _dafny.UnicodeSeqOfUtf8Bytes("xmp"))).Contains(_43_v1) {
							_coll5.Add(_43_v1, true)
						}
					}
					return _coll5.ToMap()
				}()).Contains(_44_v0) {
					_coll3.Add(_44_v0, true)
				}
			}
			return _coll3.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mbdnvowfp"), false))
	} else if _source5.Is_DC26() {
		var _45___mcc_h4 _dafny.Map = _source5.Get_().(D12_DC26).Cf35
		_ = _45___mcc_h4
		var _46_cf35 _dafny.Map = _45___mcc_h4
		_ = _46_cf35
		return func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("c"), _dafny.UnicodeSeqOfUtf8Bytes("xegsurtcy"))).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _47_v2 _dafny.Sequence
				_47_v2 = interface{}(_compr_6).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("c"), _dafny.UnicodeSeqOfUtf8Bytes("xegsurtcy")), _47_v2) {
					_coll6.Add(_47_v2, true)
				}
			}
			return _coll6.ToMap()
		}()
	} else {
		var _48___mcc_h5 D12 = _source5.Get_().(D12_DC28).Cf40
		_ = _48___mcc_h5
		var _49_cf40 D12 = _48___mcc_h5
		_ = _49_cf40
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("s"), false)).Merge(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rasmegb"), true)).Keys().Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _50_v3 _dafny.Sequence
				_50_v3 = interface{}(_compr_7).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rasmegb"), true)).Contains(_50_v3) {
					_coll7.Add(_50_v3, true)
				}
			}
			return _coll7.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()))).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(315), _dafny.CodePoint('o'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(6))).Cardinality()), _dafny.CodePoint('s')))
	} else {
		return (func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-535), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}(func(_51_i0 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
				}))).Cardinality()))).Keys().Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _52_v1 _dafny.Int
					_52_v1 = interface{}(_compr_9).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-535), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}(func(_51_i0 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
					}))).Cardinality()))).Contains(_52_v1) {
						_coll9.Add((_52_v1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkdcjuouw")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfInt64(-371)))
					}
				}
				return _coll9.ToMap()
			}()).Keys().Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _53_v0 _dafny.Int
				_53_v0 = interface{}(_compr_8).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll10 = _dafny.NewMapBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-535), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}(func(_51_i0 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
					}))).Cardinality()))).Keys().Elements()); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _52_v1 _dafny.Int
						_52_v1 = interface{}(_compr_10).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-535), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg13 _dafny.Int) interface{} {
								return coer13(arg13)
							}
						}(func(_51_i0 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
						}))).Cardinality()))).Contains(_52_v1) {
							_coll10.Add((_52_v1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkdcjuouw")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfInt64(-371)))
						}
					}
					return _coll10.ToMap()
				}()).Contains(_53_v0) {
					_coll8.Add((_53_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())), _dafny.CodePoint('e'))
				}
			}
			return _coll8.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(827), _dafny.CodePoint('b')))
	}
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(false, true), _dafny.SeqOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if !(!(true)) {
			return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(754))), _dafny.IntOfInt64(665)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(202))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.IntOfInt64(198))).Cardinality()))
		}
		return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-239))), _dafny.IntOfInt64(616)), func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_dafny.SetOf(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())))).Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _54_v0 D12
				_54_v0 = interface{}(_compr_11).(D12)
				if (_dafny.SetOf(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())))).Contains(_54_v0) {
					_coll11.Add(_54_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-170))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg14 _dafny.Int) interface{} {
							return coer14(arg14)
						}
					}(func(_55_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))).Cardinality()))
				}
			}
			return _coll11.ToMap()
		}())
	})()).Difference((_dafny.SetOf(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(333))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(Companion_D4_.Create_DC11_(true, false, _dafny.IntOfInt64(251), _dafny.UnicodeSeqOfUtf8Bytes("uid")))).Cardinality())).Cardinality()))).Cardinality()))), func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))).Cardinality(), _dafny.CodePoint('t'))).Keys().Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _56_v2 _dafny.Int
				_56_v2 = interface{}(_compr_13).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))).Cardinality(), _dafny.CodePoint('t'))).Contains(_56_v2) {
					_coll13.Add((_56_v2).Plus(_dafny.IntOfInt64(842)), _dafny.CodePoint('s'))
				}
			}
			return _coll13.ToMap()
		}())).Keys().Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _57_v1 D12
			_57_v1 = interface{}(_compr_12).(D12)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(333))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(Companion_D4_.Create_DC11_(true, false, _dafny.IntOfInt64(251), _dafny.UnicodeSeqOfUtf8Bytes("uid")))).Cardinality())).Cardinality()))).Cardinality()))), func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))).Cardinality(), _dafny.CodePoint('t'))).Keys().Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _56_v2 _dafny.Int
					_56_v2 = interface{}(_compr_14).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))).Cardinality(), _dafny.CodePoint('t'))).Contains(_56_v2) {
						_coll14.Add((_56_v2).Plus(_dafny.IntOfInt64(842)), _dafny.CodePoint('s'))
					}
				}
				return _coll14.ToMap()
			}())).Contains(_57_v1) {
				_coll12.Add(_57_v1, _dafny.IntOfInt64(230))
			}
		}
		return _coll12.ToMap()
	}())).Intersection(func() _dafny.Set {
		var _coll15 = _dafny.NewBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate((func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter18 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
						var _coll19 = _dafny.NewMapBuilder()
						_ = _coll19
						for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
							_compr_19, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _58_v6 _dafny.Int
							_58_v6 = interface{}(_compr_19).(_dafny.Int)
							if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
								_coll19.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
							}
						}
						return _coll19.ToMap()
					}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
						_compr_18, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _59_v5 D12
						_59_v5 = interface{}(_compr_18).(D12)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
							var _coll20 = _dafny.NewMapBuilder()
							_ = _coll20
							for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
								_compr_20, _ok20 := _iter20()
								if !_ok20 {
									break
								}
								var _58_v6 _dafny.Int
								_58_v6 = interface{}(_compr_20).(_dafny.Int)
								if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
									_coll20.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
								}
							}
							return _coll20.ToMap()
						}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
							_coll18.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
						}
					}
					return _coll18.ToMap()
				}()).Keys().Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _60_v4 D12
					_60_v4 = interface{}(_compr_17).(D12)
					if (func() _dafny.Map {
						var _coll21 = _dafny.NewMapBuilder()
						_ = _coll21
						for _iter21 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
							var _coll22 = _dafny.NewMapBuilder()
							_ = _coll22
							for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
								_compr_22, _ok22 := _iter22()
								if !_ok22 {
									break
								}
								var _58_v6 _dafny.Int
								_58_v6 = interface{}(_compr_22).(_dafny.Int)
								if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
									_coll22.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
								}
							}
							return _coll22.ToMap()
						}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
							_compr_21, _ok21 := _iter21()
							if !_ok21 {
								break
							}
							var _59_v5 D12
							_59_v5 = interface{}(_compr_21).(D12)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll23 = _dafny.NewMapBuilder()
								_ = _coll23
								for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_23, _ok23 := _iter23()
									if !_ok23 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_23).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll23.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll23.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
								_coll21.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
							}
						}
						return _coll21.ToMap()
					}()).Contains(_60_v4) {
						_coll17.Add(_60_v4, (func() _dafny.Map {
							var _coll24 = _dafny.NewMapBuilder()
							_ = _coll24
							for _iter24 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Elements()); ; {
								_compr_24, _ok24 := _iter24()
								if !_ok24 {
									break
								}
								var _61_v7 _dafny.CodePoint
								_61_v7 = interface{}(_compr_24).(_dafny.CodePoint)
								if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Contains(_61_v7) {
									_coll24.Add(_61_v7, true)
								}
							}
							return _coll24.ToMap()
						}()).Cardinality())
					}
				}
				return _coll17.ToMap()
			}()).Keys().Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _62_v3 D12
				_62_v3 = interface{}(_compr_16).(D12)
				if (func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter25 := _dafny.Iterate((func() _dafny.Map {
						var _coll26 = _dafny.NewMapBuilder()
						_ = _coll26
						for _iter26 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
							var _coll27 = _dafny.NewMapBuilder()
							_ = _coll27
							for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
								_compr_27, _ok27 := _iter27()
								if !_ok27 {
									break
								}
								var _58_v6 _dafny.Int
								_58_v6 = interface{}(_compr_27).(_dafny.Int)
								if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
									_coll27.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
								}
							}
							return _coll27.ToMap()
						}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
							_compr_26, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _59_v5 D12
							_59_v5 = interface{}(_compr_26).(D12)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll28 = _dafny.NewMapBuilder()
								_ = _coll28
								for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_28, _ok28 := _iter28()
									if !_ok28 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_28).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll28.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll28.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
								_coll26.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
							}
						}
						return _coll26.ToMap()
					}()).Keys().Elements()); ; {
						_compr_25, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _60_v4 D12
						_60_v4 = interface{}(_compr_25).(D12)
						if (func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter29 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll30 = _dafny.NewMapBuilder()
								_ = _coll30
								for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_30, _ok30 := _iter30()
									if !_ok30 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_30).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll30.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll30.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
								_compr_29, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _59_v5 D12
								_59_v5 = interface{}(_compr_29).(D12)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
									var _coll31 = _dafny.NewMapBuilder()
									_ = _coll31
									for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
										_compr_31, _ok31 := _iter31()
										if !_ok31 {
											break
										}
										var _58_v6 _dafny.Int
										_58_v6 = interface{}(_compr_31).(_dafny.Int)
										if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
											_coll31.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
										}
									}
									return _coll31.ToMap()
								}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
									_coll29.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
								}
							}
							return _coll29.ToMap()
						}()).Contains(_60_v4) {
							_coll25.Add(_60_v4, (func() _dafny.Map {
								var _coll32 = _dafny.NewMapBuilder()
								_ = _coll32
								for _iter32 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Elements()); ; {
									_compr_32, _ok32 := _iter32()
									if !_ok32 {
										break
									}
									var _61_v7 _dafny.CodePoint
									_61_v7 = interface{}(_compr_32).(_dafny.CodePoint)
									if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Contains(_61_v7) {
										_coll32.Add(_61_v7, true)
									}
								}
								return _coll32.ToMap()
							}()).Cardinality())
						}
					}
					return _coll25.ToMap()
				}()).Contains(_62_v3) {
					_coll16.Add(_62_v3, _dafny.IntOfInt64(754))
				}
			}
			return _coll16.ToMap()
		}(), _dafny.CodePoint('m'))).Keys().Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _63_v8 _dafny.Map
			_63_v8 = interface{}(_compr_15).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter33 := _dafny.Iterate((func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter34 := _dafny.Iterate((func() _dafny.Map {
						var _coll35 = _dafny.NewMapBuilder()
						_ = _coll35
						for _iter35 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
								_compr_36, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _58_v6 _dafny.Int
								_58_v6 = interface{}(_compr_36).(_dafny.Int)
								if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
									_coll36.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
								}
							}
							return _coll36.ToMap()
						}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
							_compr_35, _ok35 := _iter35()
							if !_ok35 {
								break
							}
							var _59_v5 D12
							_59_v5 = interface{}(_compr_35).(D12)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll37 = _dafny.NewMapBuilder()
								_ = _coll37
								for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_37, _ok37 := _iter37()
									if !_ok37 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_37).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll37.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll37.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
								_coll35.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
							}
						}
						return _coll35.ToMap()
					}()).Keys().Elements()); ; {
						_compr_34, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _60_v4 D12
						_60_v4 = interface{}(_compr_34).(D12)
						if (func() _dafny.Map {
							var _coll38 = _dafny.NewMapBuilder()
							_ = _coll38
							for _iter38 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll39 = _dafny.NewMapBuilder()
								_ = _coll39
								for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_39, _ok39 := _iter39()
									if !_ok39 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_39).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll39.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll39.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
								_compr_38, _ok38 := _iter38()
								if !_ok38 {
									break
								}
								var _59_v5 D12
								_59_v5 = interface{}(_compr_38).(D12)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
									var _coll40 = _dafny.NewMapBuilder()
									_ = _coll40
									for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
										_compr_40, _ok40 := _iter40()
										if !_ok40 {
											break
										}
										var _58_v6 _dafny.Int
										_58_v6 = interface{}(_compr_40).(_dafny.Int)
										if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
											_coll40.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
										}
									}
									return _coll40.ToMap()
								}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
									_coll38.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
								}
							}
							return _coll38.ToMap()
						}()).Contains(_60_v4) {
							_coll34.Add(_60_v4, (func() _dafny.Map {
								var _coll41 = _dafny.NewMapBuilder()
								_ = _coll41
								for _iter41 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Elements()); ; {
									_compr_41, _ok41 := _iter41()
									if !_ok41 {
										break
									}
									var _61_v7 _dafny.CodePoint
									_61_v7 = interface{}(_compr_41).(_dafny.CodePoint)
									if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Contains(_61_v7) {
										_coll41.Add(_61_v7, true)
									}
								}
								return _coll41.ToMap()
							}()).Cardinality())
						}
					}
					return _coll34.ToMap()
				}()).Keys().Elements()); ; {
					_compr_33, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _62_v3 D12
					_62_v3 = interface{}(_compr_33).(D12)
					if (func() _dafny.Map {
						var _coll42 = _dafny.NewMapBuilder()
						_ = _coll42
						for _iter42 := _dafny.Iterate((func() _dafny.Map {
							var _coll43 = _dafny.NewMapBuilder()
							_ = _coll43
							for _iter43 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
								var _coll44 = _dafny.NewMapBuilder()
								_ = _coll44
								for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
									_compr_44, _ok44 := _iter44()
									if !_ok44 {
										break
									}
									var _58_v6 _dafny.Int
									_58_v6 = interface{}(_compr_44).(_dafny.Int)
									if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
										_coll44.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
									}
								}
								return _coll44.ToMap()
							}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
								_compr_43, _ok43 := _iter43()
								if !_ok43 {
									break
								}
								var _59_v5 D12
								_59_v5 = interface{}(_compr_43).(D12)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
									var _coll45 = _dafny.NewMapBuilder()
									_ = _coll45
									for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
										_compr_45, _ok45 := _iter45()
										if !_ok45 {
											break
										}
										var _58_v6 _dafny.Int
										_58_v6 = interface{}(_compr_45).(_dafny.Int)
										if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
											_coll45.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
										}
									}
									return _coll45.ToMap()
								}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
									_coll43.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
								}
							}
							return _coll43.ToMap()
						}()).Keys().Elements()); ; {
							_compr_42, _ok42 := _iter42()
							if !_ok42 {
								break
							}
							var _60_v4 D12
							_60_v4 = interface{}(_compr_42).(D12)
							if (func() _dafny.Map {
								var _coll46 = _dafny.NewMapBuilder()
								_ = _coll46
								for _iter46 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
									var _coll47 = _dafny.NewMapBuilder()
									_ = _coll47
									for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
										_compr_47, _ok47 := _iter47()
										if !_ok47 {
											break
										}
										var _58_v6 _dafny.Int
										_58_v6 = interface{}(_compr_47).(_dafny.Int)
										if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
											_coll47.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
										}
									}
									return _coll47.ToMap()
								}()).Cardinality())), _dafny.CodePoint('d'))).Keys().Elements()); ; {
									_compr_46, _ok46 := _iter46()
									if !_ok46 {
										break
									}
									var _59_v5 D12
									_59_v5 = interface{}(_compr_46).(D12)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
										var _coll48 = _dafny.NewMapBuilder()
										_ = _coll48
										for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-428), _dafny.IntOfInt64(530))); ; {
											_compr_48, _ok48 := _iter48()
											if !_ok48 {
												break
											}
											var _58_v6 _dafny.Int
											_58_v6 = interface{}(_compr_48).(_dafny.Int)
											if ((_dafny.IntOfInt64(-428)).Cmp(_58_v6) <= 0) && ((_58_v6).Cmp(_dafny.IntOfInt64(530)) < 0) {
												_coll48.Add(Companion_Default___.SafeDivisionInt(_58_v6, _dafny.IntOfInt64(-206)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daydimu")).Cardinality()))
											}
										}
										return _coll48.ToMap()
									}()).Cardinality())), _dafny.CodePoint('d'))).Contains(_59_v5) {
										_coll46.Add(_59_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
									}
								}
								return _coll46.ToMap()
							}()).Contains(_60_v4) {
								_coll42.Add(_60_v4, (func() _dafny.Map {
									var _coll49 = _dafny.NewMapBuilder()
									_ = _coll49
									for _iter49 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Elements()); ; {
										_compr_49, _ok49 := _iter49()
										if !_ok49 {
											break
										}
										var _61_v7 _dafny.CodePoint
										_61_v7 = interface{}(_compr_49).(_dafny.CodePoint)
										if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j')))).Contains(_61_v7) {
											_coll49.Add(_61_v7, true)
										}
									}
									return _coll49.ToMap()
								}()).Cardinality())
							}
						}
						return _coll42.ToMap()
					}()).Contains(_62_v3) {
						_coll33.Add(_62_v3, _dafny.IntOfInt64(754))
					}
				}
				return _coll33.ToMap()
			}(), _dafny.CodePoint('m'))).Contains(_63_v8) {
				_coll15.Add(_63_v8)
			}
		}
		return _coll15.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm46(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(!(true))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(424)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.SetOf(!(!(!(true))))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(470))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('x'))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_64_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})), _dafny.UnicodeSeqOfUtf8Bytes("yjv")))
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if false {
			return true
		}
		return !(false)
	})(), Companion_D7_.Create_DC18_())
}
func (_static *CompanionStruct_Default___) Fm49(p0 bool, globalState *GlobalState) bool {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(-701))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(959))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_65_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(926)
	}))).Cardinality())) > 0
}
func (_static *CompanionStruct_Default___) Fm50(globalState *GlobalState) D3 {
	if true {
		return Companion_D3_.Create_DC9_()
	} else {
		return Companion_D3_.Create_DC9_()
	}
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm52(globalState *GlobalState) _dafny.Map {
	var _source6 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.SeqOf(true))
	_ = _source6
	var _66___mcc_h0 _dafny.MultiSet = _source6
	_ = _66___mcc_h0
	var _67_cf34 _dafny.MultiSet = _66___mcc_h0
	_ = _67_cf34
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(369), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
		var _coll50 = _dafny.NewBuilder()
		_ = _coll50
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(944), _dafny.IntOfInt64(690))); ; {
			_compr_50, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _68_v0 _dafny.Int
			_68_v0 = interface{}(_compr_50).(_dafny.Int)
			if ((_dafny.IntOfInt64(944)).Cmp(_68_v0) <= 0) && ((_68_v0).Cmp(_dafny.IntOfInt64(690)) < 0) {
				_coll50.Add(Companion_Default___.SafeDivisionInt(_68_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kgpvcbka")).Cardinality()))).Cardinality()))
			}
		}
		return _coll50.ToSet()
	}()).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kc")).Cardinality()), false), func() _dafny.Map {
		var _coll51 = _dafny.NewMapBuilder()
		_ = _coll51
		for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-717), _dafny.IntOfInt64(130))); ; {
			_compr_51, _ok51 := _iter51()
			if !_ok51 {
				break
			}
			var _69_v0 _dafny.Int
			_69_v0 = interface{}(_compr_51).(_dafny.Int)
			if ((_dafny.IntOfInt64(-717)).Cmp(_69_v0) <= 0) && ((_69_v0).Cmp(_dafny.IntOfInt64(130)) < 0) {
				_coll51.Add((_69_v0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_70_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality())), true)
			}
		}
		return _coll51.ToMap()
	}())).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll52 = _dafny.NewBuilder()
		_ = _coll52
		for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(634), _dafny.IntOfInt64(678))); ; {
			_compr_52, _ok52 := _iter52()
			if !_ok52 {
				break
			}
			var _71_v1 _dafny.Int
			_71_v1 = interface{}(_compr_52).(_dafny.Int)
			if ((_dafny.IntOfInt64(634)).Cmp(_71_v1) <= 0) && ((_71_v1).Cmp(_dafny.IntOfInt64(678)) < 0) {
				_coll52.Add(Companion_Default___.SafeModuloInt(_71_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oghwvum")).Cardinality())))).Cardinality())))).Cardinality()))
			}
		}
		return _coll52.ToSet()
	}()).Cardinality(), true)))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('d'))
		}
		return _dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('i'))
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-406))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_72_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))))
}
func (_static *CompanionStruct_Default___) M15(globalState *GlobalState) (_dafny.Int, _dafny.Map, bool, *C0) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 *C0 = (*C0)(nil)
	_ = r3
	var _73_v0 _dafny.Sequence
	_ = _73_v0
	_73_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dpud")
	var _74_v1 _dafny.Int
	_ = _74_v1
	_74_v1 = _dafny.IntOfInt64(36)
	var _hi0 _dafny.Int = _74_v1
	_ = _hi0
	for _75_i0 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bheo"), _73_v0)).Cardinality()); _75_i0.Cmp(_hi0) < 0; _75_i0 = _75_i0.Plus(_dafny.One) {
		var _76_v2 bool
		_ = _76_v2
		_76_v2 = true
		var _77_v3 _dafny.Array
		_ = _77_v3
		var _nwElement0_0 bool = _76_v2
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		_77_v3 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_77_v3), 0))
		_ = _index0
		(_77_v3).ArraySet1(_76_v2, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_77_v3), 0))
		_ = _index1
		(_77_v3).ArraySet1(_76_v2, (_index1).Int())
		(globalState).F8 = _75_i0
		var _78_v4 _dafny.Sequence
		_ = _78_v4
		_78_v4 = _dafny.SeqOf(_76_v2, _76_v2)
		var _79_v5 _dafny.MultiSet
		_ = _79_v5
		_79_v5 = _dafny.MultiSetOf(true, _76_v2)
		var _80_v6 _dafny.Sequence
		_ = _80_v6
		_80_v6 = _dafny.SeqOf(_74_v1, _dafny.IntOfUint32((_dafny.SeqOf(_74_v1)).Cardinality()))
		(globalState).F3 = ((_dafny.IntOfUint32((_78_v4).Cardinality())).Times((_dafny.MultiSetOf(_74_v1, _74_v1, (_79_v5).Cardinality(), (_79_v5).Cardinality())).Cardinality())).Minus((_80_v6).Select((Companion_Default___.SafeIndex(_74_v1, _dafny.IntOfUint32((_80_v6).Cardinality()))).Uint32()).(_dafny.Int))
		_76_v2 = !(_76_v2)
	}
	(globalState).F3 = _74_v1
	(globalState).F3 = _dafny.IntOfUint32((_73_v0).Cardinality())
	var _81_v7 bool
	_ = _81_v7
	_81_v7 = false
	var _82_v9 _dafny.Map
	_ = _82_v9
	_82_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v7, func() _dafny.Set {
		var _coll53 = _dafny.NewBuilder()
		_ = _coll53
		for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-904), _dafny.IntOfInt64(973))); ; {
			_compr_53, _ok53 := _iter53()
			if !_ok53 {
				break
			}
			var _83_v8 _dafny.Int
			_83_v8 = interface{}(_compr_53).(_dafny.Int)
			if ((_dafny.IntOfInt64(-904)).Cmp(_83_v8) <= 0) && ((_83_v8).Cmp(_dafny.IntOfInt64(973)) < 0) {
				_coll53.Add((_83_v8).Times(_74_v1))
			}
		}
		return _coll53.ToSet()
	}())
	var _84_v10 _dafny.Sequence
	_ = _84_v10
	_84_v10 = _dafny.SeqOf(_74_v1, _74_v1, _74_v1, _dafny.IntOfInt64(657))
	(globalState).F3 = (_dafny.Zero).Minus((func() _dafny.Int {
		if true {
			return ((_82_v9).Cardinality()).Times((_84_v10).Select((Companion_Default___.SafeIndex(_74_v1, _dafny.IntOfUint32((_84_v10).Cardinality()))).Uint32()).(_dafny.Int))
		}
		return _74_v1
	})())
	var _85_v11 _dafny.Array
	_ = _85_v11
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
	_ = _nw1
	_85_v11 = _nw1
	var _86_v12 _dafny.Sequence
	_ = _86_v12
	_86_v12 = _dafny.SeqOf(_85_v11)
	_85_v11 = (_86_v12).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_74_v1), _dafny.IntOfUint32((_86_v12).Cardinality()))).Uint32()).(_dafny.Array)
	var _87_v13 _dafny.Array
	_ = _87_v13
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
	_ = _nw2
	_87_v13 = _nw2
	for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_87_v13), 0))); ; {
		_guard_loop_0, _ok54 := _iter54()
		if !_ok54 {
			break
		}
		var _88_i1 _dafny.Int
		_88_i1 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_88_i1).Sign() != -1) && ((_88_i1).Cmp(_dafny.ArrayLen((_87_v13), 0)) < 0)) {
			(_87_v13).ArraySet1(Companion_Default___.SafeModuloInt(_88_i1, Companion_Default___.SafeModuloInt(_74_v1, _74_v1)), (_88_i1).Int())
		}
	}
	r0 = _74_v1
	var _89_v14 _dafny.Map
	_ = _89_v14
	_89_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v13, Companion_Default___.Fm54(_74_v1, globalState))
	var _90_v15 _dafny.Map
	_ = _90_v15
	_90_v15 = _89_v14
	r1 = (_90_v15)
	r2 = _81_v7
	var _91_v16 *C0
	_ = _91_v16
	var _nw3 *C0 = New_C0_()
	_ = _nw3
	_nw3.Ctor__()
	_91_v16 = _nw3
	r3 = (func() *C0 {
		if _81_v7 {
			return _91_v16
		}
		return _91_v16
	})()
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _92_v0 _dafny.Int
	_ = _92_v0
	_92_v0 = _dafny.IntOfInt64(545)
	var _93_v1 _dafny.Sequence
	_ = _93_v1
	_93_v1 = _dafny.SeqOf(_92_v0)
	var _94_v2 bool
	_ = _94_v2
	_94_v2 = false
	var _95_v3 _dafny.Map
	_ = _95_v3
	_95_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _92_v0)
	var _96_v4 D0
	_ = _96_v4
	_96_v4 = Companion_D0_.Create_DC0_((_95_v3).Cardinality())
	var _97_v5 _dafny.Sequence
	_ = _97_v5
	_97_v5 = _dafny.UnicodeSeqOfUtf8Bytes("cj")
	var _98_v6 _dafny.Set
	_ = _98_v6
	_98_v6 = _dafny.SetOf(_97_v5)
	var _99_v7 _dafny.MultiSet
	_ = _99_v7
	_99_v7 = _dafny.MultiSetOf(_92_v0)
	var _100_globalState *GlobalState
	_ = _100_globalState
	var _nw4 *GlobalState = New_GlobalState_()
	_ = _nw4
	_nw4.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_93_v1, (Companion_Default___.SafeIndex((_96_v4).Dtor_cf0(), _dafny.IntOfUint32((_93_v1).Cardinality()))).Uint32(), _92_v0), (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_93_v1, (Companion_Default___.SafeIndex((_96_v4).Dtor_cf0(), _dafny.IntOfUint32((_93_v1).Cardinality()))).Uint32(), _92_v0)).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, (_98_v6).Cardinality())).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("w"), _97_v5), _99_v7, _dafny.IntOfInt64(742), false, false, true, _dafny.IntOfInt64(339), _dafny.IntOfInt64(344))
	_100_globalState = _nw4
	var _101_v8 _dafny.Array
	_ = _101_v8
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw5 _dafny.Array
	_ = _nw5
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw5 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_102_v2 bool) func(_dafny.Int) _dafny.Int {
			return func(_103_i0 _dafny.Int) _dafny.Int {
				return (_103_i0).Times((_dafny.SetOf(_102_v2, _102_v2)).Cardinality())
			}
		})(_94_v2)
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
	_101_v8 = _nw5
	var _104_v9 D0
	_ = _104_v9
	_104_v9 = Companion_D0_.Create_DC1_(_92_v0, _92_v0)
	var _rhs0 _dafny.Array = _101_v8
	_ = _rhs0
	var _rhs1 _dafny.Int = (_104_v9).Dtor_cf1()
	_ = _rhs1
	var _lhs0 *GlobalState = _100_globalState
	_ = _lhs0
	_101_v8 = _rhs0
	_lhs0.F3 = _rhs1
	var _105_i1 _dafny.Int
	_ = _105_i1
	_105_i1 = _dafny.Zero
	{
		for _94_v2 {
			{
				if (_105_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_105_i1 = (_105_i1).Plus(_dafny.One)
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))
				_ = _index2
				(_101_v8).ArraySet1(_92_v0, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))
				_ = _index3
				(_101_v8).ArraySet1(_92_v0, (_index3).Int())
				var _source7 D0 = _96_v4
				_ = _source7
				if _source7.Is_DC1() {
					var _106___mcc_h0 _dafny.Int = _source7.Get_().(D0_DC1).Cf1
					_ = _106___mcc_h0
					var _107___mcc_h1 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
					_ = _107___mcc_h1
					var _108_cf2 _dafny.Int = _107___mcc_h1
					_ = _108_cf2
					var _109_cf1 _dafny.Int = _106___mcc_h0
					_ = _109_cf1
					var _110_v10 _dafny.Array
					_ = _110_v10
					var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
					_ = _nw6
					_110_v10 = _nw6
					var _len0_1 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_1
					var _nw7 _dafny.Array
					_ = _nw7
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw7 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) _dafny.Int = (func(_111_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_112_i2 _dafny.Int) _dafny.Int {
								return (_112_i2).Minus(_111_cf2)
							}
						})(_108_cf2)
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
					_110_v10 = _nw7
					(_100_globalState).F3 = _dafny.IntOfInt64(716)
					(_100_globalState).F2 = _99_v7
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))
					_ = _index4
					var _rhs2 bool = _94_v2
					_ = _rhs2
					var _rhs3 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_108_cf2, _92_v0))
					_ = _rhs3
					var _rhs4 bool = (_92_v0).Cmp((_109_cf1).Times(_109_cf1)) != 0
					_ = _rhs4
					var _lhs1 *GlobalState = _100_globalState
					_ = _lhs1
					var _lhs2 _dafny.Array = _101_v8
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))
					_ = _lhs3
					var _lhs4 *GlobalState = _100_globalState
					_ = _lhs4
					_lhs1.F4 = _rhs2
					(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
					_lhs4.F4 = _rhs4
				} else if _source7.Is_DC2() {
					var _113_v11 _dafny.Array
					_ = _113_v11
					var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
					_ = _nw8
					_113_v11 = _nw8
					_113_v11 = _113_v11
					_95_v3 = (_95_v3).Update(!(!(_94_v2) || (_94_v2)), (_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int))
					var _114_v12 _dafny.Map
					_ = _114_v12
					_114_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int), _92_v0)
					(_100_globalState).F0 = (func() _dafny.Sequence {
						if _94_v2 {
							return _93_v1
						}
						return _dafny.SeqOf(_92_v0, (_114_v12).Cardinality())
					})()
					_95_v3 = _95_v3
				} else if _source7.Is_DC0() {
					var _115___mcc_h2 _dafny.Int = _source7.Get_().(D0_DC0).Cf0
					_ = _115___mcc_h2
					var _116_cf0 _dafny.Int = _115___mcc_h2
					_ = _116_cf0
					(_100_globalState).F8 = (Companion_Default___.SafeDivisionInt(_116_cf0, (_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int))).Minus((Companion_D0_.Create_DC0_((Companion_Default___.Fm0(_116_cf0, _100_globalState)).Cardinality())).Dtor_cf0())
					(_100_globalState).F4 = _94_v2
					var _117_v13 _dafny.Array
					_ = _117_v13
					var _nwElement0_1 _dafny.Sequence = _93_v1
					_ = _nwElement0_1
					var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
					_ = _nw9
					(_nw9).ArraySet1(_nwElement0_1, 0)
					_117_v13 = _nw9
					var _118_v14 _dafny.Sequence
					_ = _118_v14
					_118_v14 = _dafny.SeqOf(_94_v2, _94_v2, _94_v2, false, false)
					var _119_v15 _dafny.Set
					_ = _119_v15
					_119_v15 = _dafny.SetOf(_94_v2)
					var _120_v16 _dafny.Map
					_ = _120_v16
					_120_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_119_v15).Cardinality(), _116_cf0)
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_117_v13), 0))
					_ = _index5
					(_117_v13).ArraySet1(Companion_Default___.Fm1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_116_cf0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_118_v14).Cardinality())), _dafny.IntOfUint32((_97_v5).Cardinality()), _dafny.IntOfInt64(454), _dafny.IntOfUint32((_97_v5).Cardinality()))).Cardinality())), _dafny.CodePoint('f'), !(_94_v2), _120_v16, _100_globalState), (_index5).Int())
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_117_v13), 0))
					_ = _index6
					(_117_v13).ArraySet1(_93_v1, (_index6).Int())
					_92_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(681), (_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int)))
				} else {
					var _121___mcc_h3 D0 = _source7.Get_().(D0_DC3).Cf3
					_ = _121___mcc_h3
					var _122_cf3 D0 = _121___mcc_h3
					_ = _122_cf3
					(_100_globalState).F8 = (_dafny.Zero).Minus((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int))
					var _123_v17 *C5
					_ = _123_v17
					var _nw10 *C5 = New_C5_()
					_ = _nw10
					_nw10.Ctor__()
					_123_v17 = _nw10
					var _124_v18 _dafny.Set
					_ = _124_v18
					_124_v18 = _dafny.SetOf((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int), _92_v0)
					var _rhs5 bool = _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ethjt"), (_97_v5).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_97_v5).Cardinality()))).Uint32()).(_dafny.CodePoint))
					_ = _rhs5
					var _rhs6 _dafny.Sequence = _97_v5
					_ = _rhs6
					var _rhs7 bool = ((_124_v18).Cardinality()).Cmp(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int)), (_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int))) >= 0
					_ = _rhs7
					var _rhs8 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_97_v5, _97_v5)
					_ = _rhs8
					var _lhs5 *GlobalState = _100_globalState
					_ = _lhs5
					_94_v2 = _rhs5
					_97_v5 = _rhs6
					_lhs5.F4 = _rhs7
					_97_v5 = _rhs8
					(_100_globalState).F4 = _94_v2
				}
				var _125_v19 *C8
				_ = _125_v19
				var _nw11 *C8 = New_C8_()
				_ = _nw11
				_nw11.Ctor__()
				_125_v19 = _nw11
				var _126_v20 _dafny.Array
				_ = _126_v20
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_2
				var _nw12 _dafny.Array
				_ = _nw12
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw12 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) bool = (func(_127_v2 bool) func(_dafny.Int) bool {
						return func(_128_i3 _dafny.Int) bool {
							return _127_v2
						}
					})(_94_v2)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw12 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw12).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw12).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_126_v20 = _nw12
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_126_v20), 0))
				_ = _index7
				(_126_v20).ArraySet1(!(_94_v2), (_index7).Int())
				var _129_v21 _dafny.Sequence
				_ = _129_v21
				_129_v21 = _dafny.SeqOf(_94_v2, _94_v2, false, _94_v2)
				var _130_v22 _dafny.MultiSet
				_ = _130_v22
				_130_v22 = _dafny.MultiSetOf(_94_v2)
				var _131_v23 _dafny.MultiSet
				_ = _131_v23
				_131_v23 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_129_v21)).Update(_94_v2, Companion_Default___.Abs(_92_v0)), _130_v22, _130_v22)
				var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_126_v20), 0))
				_ = _index8
				(_126_v20).ArraySet1(!(Companion_Default___.Fm49(Companion_Default___.Fm49(true, _100_globalState), _100_globalState)) || ((_131_v23).IsProperSubsetOf(_131_v23)), (_index8).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _132_v24 _dafny.Set
	_ = _132_v24
	_132_v24 = _dafny.SetOf(_dafny.IntOfUint32((_97_v5).Cardinality()), _92_v0)
	if !((_132_v24).IsSubsetOf(_132_v24)) {
		if _94_v2 {
			var _133_v25 _dafny.Array
			_ = _133_v25
			var _nwElement0_2 bool = _94_v2
			_ = _nwElement0_2
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_2, 0)
			(_nw13).ArraySet1(!(_94_v2), 1)
			_133_v25 = _nw13
			var _134_v26 _dafny.Map
			_ = _134_v26
			_134_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, _133_v25)
			_134_v26 = _134_v26
			var _135_v27 *C5
			_ = _135_v27
			var _nw14 *C5 = New_C5_()
			_ = _nw14
			_nw14.Ctor__()
			_135_v27 = _nw14
			var _136_v28 _dafny.Array
			_ = _136_v28
			var _nwElement0_3 *C5 = _135_v27
			_ = _nwElement0_3
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_3, 0)
			(_nw15).ArraySet1(_135_v27, 1)
			(_nw15).ArraySet1(_135_v27, 2)
			(_nw15).ArraySet1(_135_v27, 3)
			(_nw15).ArraySet1(_135_v27, 4)
			_136_v28 = _nw15
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_136_v28), 0))
			_ = _index9
			(_136_v28).ArraySet1(_135_v27, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_136_v28), 0))
			_ = _index10
			(_136_v28).ArraySet1(_135_v27, (_index10).Int())
			var _137_v29 _dafny.Array
			_ = _137_v29
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_3
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = func(_138_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw16 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw16).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw16).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_137_v29 = _nw16
			var _139_v30 _dafny.CodePoint
			_ = _139_v30
			_139_v30 = _dafny.CodePoint('d')
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_137_v29), 0))
			_ = _index11
			(_137_v29).ArraySet1CodePoint(_139_v30, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_137_v29), 0))
			_ = _index12
			(_137_v29).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index12).Int())
			var _140_v31 *C5
			_ = _140_v31
			var _nw17 *C5 = New_C5_()
			_ = _nw17
			_nw17.Ctor__()
			_140_v31 = _nw17
			var _141_v32 *C9
			_ = _141_v32
			var _nw18 *C9 = New_C9_()
			_ = _nw18
			_nw18.Ctor__(_94_v2)
			_141_v32 = _nw18
		} else {
			var _142_v33 D23
			_ = _142_v33
			_142_v33 = Companion_D23_.Create_DC46_()
			var _143_v34 _dafny.MultiSet
			_ = _143_v34
			_143_v34 = _dafny.MultiSetOf(_142_v33)
			var _144_v35 _dafny.Map
			_ = _144_v35
			_144_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _143_v34)
			var _145_v36 _dafny.Map
			_ = _145_v36
			_145_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_144_v35).Cardinality(), (_dafny.Zero).Minus(_92_v0))
			_145_v36 = Companion_Default___.Fm52(_100_globalState)
			(_100_globalState).F3 = _92_v0
			var _146_v37 _dafny.Array
			_ = _146_v37
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_4
			var _nw19 _dafny.Array
			_ = _nw19
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw19 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_147_v2 bool) func(_dafny.Int) bool {
					return func(_148_i5 _dafny.Int) bool {
						return _147_v2
					}
				})(_94_v2)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw19 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw19).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw19).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_146_v37 = _nw19
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_146_v37), 0))
			_ = _index13
			(_146_v37).ArraySet1(!(_94_v2), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_146_v37), 0))
			_ = _index14
			(_146_v37).ArraySet1((_92_v0).Cmp(_92_v0) >= 0, (_index14).Int())
			var _149_v38 _dafny.Map
			_ = _149_v38
			_149_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v8, _146_v37)
			_149_v38 = _149_v38
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_146_v37), 0))
			_ = _index15
			(_146_v37).ArraySet1((_dafny.IntOfUint32((_97_v5).Cardinality())).Cmp(_92_v0) < 0, (_index15).Int())
		}
		var _150_v39 _dafny.Int
		_ = _150_v39
		var _151_v40 _dafny.Map
		_ = _151_v40
		var _152_v41 bool
		_ = _152_v41
		var _153_v42 *C0
		_ = _153_v42
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Map
		_ = _out1
		var _out2 bool
		_ = _out2
		var _out3 *C0
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M15(_100_globalState)
		_150_v39 = _out0
		_151_v40 = _out1
		_152_v41 = _out2
		_153_v42 = _out3
		(_100_globalState).F4 = _152_v41
		var _154_v43 D10
		_ = _154_v43
		_154_v43 = Companion_D10_.Create_DC24_(_dafny.IntOfUint32((Companion_Default___.Fm23(_97_v5, _100_globalState)).Cardinality()), _152_v41, _97_v5)
		var _155_v44 _dafny.MultiSet
		_ = _155_v44
		_155_v44 = _dafny.MultiSetOf(_152_v41, (_154_v43).Dtor_cf32(), _dafny.Companion_Sequence_.IsProperPrefixOf(_97_v5, _97_v5))
		_155_v44 = ((_155_v44).Update(!(_94_v2), Companion_Default___.Abs((Companion_Default___.Fm35(_152_v41, _92_v0, _94_v2, _152_v41, _100_globalState)).Cardinality()))).Difference(_155_v44)
		var _156_v45 _dafny.Int
		_ = _156_v45
		var _157_v46 _dafny.Map
		_ = _157_v46
		var _158_v47 bool
		_ = _158_v47
		var _159_v48 *C0
		_ = _159_v48
		var _out4 _dafny.Int
		_ = _out4
		var _out5 _dafny.Map
		_ = _out5
		var _out6 bool
		_ = _out6
		var _out7 *C0
		_ = _out7
		_out4, _out5, _out6, _out7 = Companion_Default___.M15(_100_globalState)
		_156_v45 = _out4
		_157_v46 = _out5
		_158_v47 = _out6
		_159_v48 = _out7
	} else {
		if _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("q"), Companion_Default___.Fm18(_94_v2, _94_v2, _100_globalState)) {
			(_100_globalState).F8 = (_93_v1).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_93_v1).Cardinality()))).Uint32()).(_dafny.Int)
			(_100_globalState).F8 = (_92_v0).Plus(_92_v0)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_101_v8), 0))
			_ = _index16
			(_101_v8).ArraySet1(_92_v0, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_101_v8), 0))
			_ = _index17
			(_101_v8).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_92_v0), _dafny.IntOfUint32((_dafny.SeqOf((_95_v3).Cardinality())).Cardinality())), (_index17).Int())
			(_100_globalState).F0 = _93_v1
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_101_v8), 0))
			_ = _index18
			(_101_v8).ArraySet1(_92_v0, (_index18).Int())
		} else {
			var _160_v49 *C2
			_ = _160_v49
			var _nw20 *C2 = New_C2_()
			_ = _nw20
			_nw20.Ctor__()
			_160_v49 = _nw20
			var _161_v50 *C7
			_ = _161_v50
			var _nw21 *C7 = New_C7_()
			_ = _nw21
			_nw21.Ctor__()
			_161_v50 = _nw21
			var _162_v51 _dafny.Map
			_ = _162_v51
			_162_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v50, (_161_v50).Fm3(_93_v1, _99_v7, _94_v2, _100_globalState))
			var _163_v52 _dafny.Map
			_ = _163_v52
			_163_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _162_v51)
			(_100_globalState).F4 = (!(_162_v51).Equals((func() _dafny.Map {
				if (_163_v52).Contains(false) {
					return (_163_v52).Get(false).(_dafny.Map)
				}
				return _162_v51
			})())) == (_94_v2)
			var _164_v53 *C9
			_ = _164_v53
			var _nw22 *C9 = New_C9_()
			_ = _nw22
			_nw22.Ctor__(_94_v2)
			_164_v53 = _nw22
			var _165_v54 _dafny.Map
			_ = _165_v54
			_165_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, _164_v53)
			var _166_v55 T0
			_ = _166_v55
			var _nw23 *C5 = New_C5_()
			_ = _nw23
			_nw23.Ctor__()
			_166_v55 = _nw23
			var _167_v56 _dafny.Map
			_ = _167_v56
			_167_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v55, _92_v0)
			var _168_v57 D26
			_ = _168_v57
			_168_v57 = Companion_D26_.Create_DC55_(_166_v55)
			var _169_v58 _dafny.Sequence
			_ = _169_v58
			_169_v58 = _dafny.SeqOf(_94_v2, _94_v2)
			var _170_v59 _dafny.Array
			_ = _170_v59
			var _nwElement0_4 _dafny.Map = _165_v54
			_ = _nwElement0_4
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(5))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_4, 0)
			(_nw24).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_167_v56).Contains((_168_v57).Dtor_cf77()) {
					return (_167_v56).Get((_168_v57).Dtor_cf77()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_169_v58).Cardinality())
			})(), _164_v53), 1)
			(_nw24).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, _164_v53), 2)
			(_nw24).ArraySet1(_165_v54, 3)
			(_nw24).ArraySet1(_165_v54, 4)
			_170_v59 = _nw24
			_170_v59 = _170_v59
			_97_v5 = _97_v5
			(_100_globalState).F8 = _92_v0
		}
		var _171_v60 *C8
		_ = _171_v60
		var _nw25 *C8 = New_C8_()
		_ = _nw25
		_nw25.Ctor__()
		_171_v60 = _nw25
		var _172_v61 _dafny.Map
		_ = _172_v61
		_172_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _97_v5)
		(_171_v60).M4(((_dafny.Zero).Minus(_92_v0)).Cmp(_92_v0) == 0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _97_v5)).Merge(_172_v61), _dafny.IntOfInt64(709), _100_globalState)
		if _94_v2 {
			_101_v8 = (func() _dafny.Array {
				if !(_94_v2) {
					return _101_v8
				}
				return _101_v8
			})()
			(_100_globalState).F4 = !(_94_v2)
			(_171_v60).M4(_94_v2, _172_v61, _92_v0, _100_globalState)
			var _173_v62 _dafny.Sequence
			_ = _173_v62
			_173_v62 = _dafny.SeqOf(_94_v2)
			_173_v62 = _dafny.Companion_Sequence_.Concatenate(_173_v62, _dafny.SeqOf(!(_94_v2), false, _94_v2))
			_101_v8 = _101_v8
		} else {
			var _174_v63 _dafny.Map
			_ = _174_v63
			_174_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, false)
			var _175_v64 _dafny.Array
			_ = _175_v64
			var _nwElement0_5 _dafny.Map = _95_v3
			_ = _nwElement0_5
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(12))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_5, 0)
			(_nw26).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _dafny.IntOfInt64(-686))).Update(_94_v2, _92_v0), 1)
			(_nw26).ArraySet1((Companion_Default___.Fm46(_100_globalState)).Update((func() bool {
				if (_174_v63).Contains(!(_94_v2)) {
					return (_174_v63).Get(!(_94_v2)).(bool)
				}
				return _94_v2
			})(), _dafny.IntOfInt64(379)), 2)
			(_nw26).ArraySet1(_95_v3, 3)
			(_nw26).ArraySet1(_95_v3, 4)
			(_nw26).ArraySet1(_95_v3, 5)
			(_nw26).ArraySet1(_95_v3, 6)
			(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _92_v0), 7)
			(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _92_v0), 8)
			(_nw26).ArraySet1(_95_v3, 9)
			(_nw26).ArraySet1(_95_v3, 10)
			(_nw26).ArraySet1(_95_v3, 11)
			_175_v64 = _nw26
			var _176_v65 T0
			_ = _176_v65
			var _nw27 *C4 = New_C4_()
			_ = _nw27
			_nw27.Ctor__(_175_v64)
			_176_v65 = _nw27
			_176_v65 = _176_v65
			var _177_v66 _dafny.Array
			_ = _177_v66
			var _nwElement0_6 _dafny.Sequence = (func() _dafny.Sequence {
				if true {
					return _93_v1
				}
				return _93_v1
			})()
			_ = _nwElement0_6
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(6))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_6, 0)
			(_nw28).ArraySet1(_93_v1, 1)
			(_nw28).ArraySet1(_93_v1, 2)
			(_nw28).ArraySet1(_dafny.SeqOf(_92_v0, _92_v0, _92_v0), 3)
			(_nw28).ArraySet1(_93_v1, 4)
			(_nw28).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_178_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_179_i6 _dafny.Int) _dafny.Int {
					return _178_v0
				}
			})(_92_v0))), 5)
			_177_v66 = _nw28
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))
			_ = _index19
			(_177_v66).ArraySet1(_93_v1, (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))
			_ = _index20
			(_177_v66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_93_v1, _93_v1), (_index20).Int())
			var _180_v67 _dafny.Array
			_ = _180_v67
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_5
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_181_v3 _dafny.Map) func(_dafny.Int) bool {
					return func(_182_i7 _dafny.Int) bool {
						return !(_181_v3).Equals(_181_v3)
					}
				})(_95_v3)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw29 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw29).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw29).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_180_v67 = _nw29
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_180_v67), 0))
			_ = _index21
			(_180_v67).ArraySet1((_dafny.IntOfUint32((_97_v5).Cardinality())).Cmp(_92_v0) <= 0, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_180_v67), 0))
			_ = _index22
			(_180_v67).ArraySet1(!(_94_v2), (_index22).Int())
			(_100_globalState).F4 = !(!((_92_v0).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((func() bool {
				if (_174_v63).Contains(_94_v2) {
					return (_174_v63).Get(_94_v2).(bool)
				}
				return (_180_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_180_v67), 0))).Int()).(bool)
			})())).Cardinality())) <= 0) || ((_180_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_180_v67), 0))).Int()).(bool)))
			var _183_v68 _dafny.MultiSet
			_ = _183_v68
			_183_v68 = _dafny.MultiSetOf(true)
			var _184_v69 _dafny.Sequence
			_ = _184_v69
			_184_v69 = _dafny.SeqOf(_183_v68)
			var _185_v70 _dafny.Sequence
			_ = _185_v70
			_185_v70 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32(((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _92_v0))
			var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_93_v1, _dafny.SeqOf(((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32(((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int))), (_185_v70).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_185_v70).Cardinality()))).Uint32()).(_dafny.Sequence)), (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_93_v1, _dafny.SeqOf(((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32(((_177_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_177_v66), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int))), (_185_v70).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_185_v70).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()))).Uint32(), _92_v0)
			_ = _rhs9
			var _rhs10 _dafny.Sequence = _184_v69
			_ = _rhs10
			var _rhs11 _dafny.Int = (_92_v0).Plus(_92_v0)
			_ = _rhs11
			var _lhs6 *GlobalState = _100_globalState
			_ = _lhs6
			_93_v1 = _rhs9
			_184_v69 = _rhs10
			_lhs6.F3 = _rhs11
		}
		if (func() bool {
			if _94_v2 {
				return _94_v2
			}
			return _94_v2
		})() {
			var _186_v71 _dafny.CodePoint
			_ = _186_v71
			_186_v71 = _dafny.CodePoint('r')
			var _187_v72 D10
			_ = _187_v72
			_187_v72 = Companion_D10_.Create_DC24_((func() _dafny.Int {
				if (_99_v7).Contains((_dafny.Zero).Minus(_92_v0)) {
					return (_99_v7).Multiplicity((_dafny.Zero).Minus(_92_v0))
				}
				return Companion_Default___.Fm6(_dafny.IntOfInt64(68), _100_globalState)
			})(), _94_v2, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_97_v5, _dafny.SeqOf(_186_v71)), (Companion_Default___.SafeIndex((Companion_Default___.Fm7(_186_v71, _100_globalState)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_97_v5, _dafny.SeqOf(_186_v71))).Cardinality()))).Uint32(), _186_v71))
			var _pat_let_tv0 = _94_v2
			_ = _pat_let_tv0
			var _pat_let_tv1 = _92_v0
			_ = _pat_let_tv1
			var _pat_let_tv2 = _92_v0
			_ = _pat_let_tv2
			var _pat_let_tv3 = _94_v2
			_ = _pat_let_tv3
			var _pat_let_tv4 = _92_v0
			_ = _pat_let_tv4
			_187_v72 = func(_pat_let0_0 D10) D10 {
				return func(_193_dt__update__tmp_h2 D10) D10 {
					return func(_pat_let6_0 _dafny.Int) D10 {
						return func(_194_dt__update_hcf31_h1 _dafny.Int) D10 {
							return Companion_D10_.Create_DC24_(_194_dt__update_hcf31_h1, (_193_dt__update__tmp_h2).Dtor_cf32(), (_193_dt__update__tmp_h2).Dtor_cf33())
						}(_pat_let6_0)
					}(_pat_let_tv4)
				}(_pat_let0_0)
			}(func(_pat_let1_0 D10) D10 {
				return func(_191_dt__update__tmp_h1 D10) D10 {
					return func(_pat_let5_0 bool) D10 {
						return func(_192_dt__update_hcf32_h1 bool) D10 {
							return Companion_D10_.Create_DC24_((_191_dt__update__tmp_h1).Dtor_cf31(), _192_dt__update_hcf32_h1, (_191_dt__update__tmp_h1).Dtor_cf33())
						}(_pat_let5_0)
					}(_pat_let_tv3)
				}(_pat_let1_0)
			}(func(_pat_let2_0 D10) D10 {
				return func(_188_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let3_0 bool) D10 {
						return func(_189_dt__update_hcf32_h0 bool) D10 {
							return func(_pat_let4_0 _dafny.Int) D10 {
								return func(_190_dt__update_hcf31_h0 _dafny.Int) D10 {
									return Companion_D10_.Create_DC24_(_190_dt__update_hcf31_h0, _189_dt__update_hcf32_h0, (_188_dt__update__tmp_h0).Dtor_cf33())
								}(_pat_let4_0)
							}(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv1, _pat_let_tv2)).Cardinality()))
						}(_pat_let3_0)
					}(_pat_let_tv0)
				}(_pat_let2_0)
			}(_187_v72)))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))
			_ = _index23
			(_101_v8).ArraySet1(_92_v0, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))
			_ = _index24
			(_101_v8).ArraySet1(Companion_Default___.SafeDivisionInt(_92_v0, _92_v0), (_index24).Int())
			var _195_v73 _dafny.Array
			_ = _195_v73
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(5))
			_ = _nw30
			_195_v73 = _nw30
			var _196_v74 _dafny.Array
			_ = _196_v74
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw31
			_196_v74 = _nw31
			var _197_v75 _dafny.Sequence
			_ = _197_v75
			_197_v75 = _dafny.SeqOf(_196_v74)
			var _198_v76 D2
			_ = _198_v76
			_198_v76 = Companion_D2_.Create_DC5_((_197_v75).Select((Companion_Default___.SafeIndex((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_197_v75).Cardinality()))).Uint32()).(_dafny.Array))
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_195_v73), 0))
			_ = _index25
			(_195_v73).ArraySet1(_198_v76, (_index25).Int())
			var _199_v77 *C1
			_ = _199_v77
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__()
			_199_v77 = _nw32
			var _200_v78 _dafny.Array
			_ = _200_v78
			var _nwElement0_7 *C1 = (func() *C1 {
				if _94_v2 {
					return _199_v77
				}
				return _199_v77
			})()
			_ = _nwElement0_7
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(4))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_7, 0)
			(_nw33).ArraySet1(_199_v77, 1)
			(_nw33).ArraySet1(_199_v77, 2)
			(_nw33).ArraySet1(_199_v77, 3)
			_200_v78 = _nw33
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_200_v78), 0))
			_ = _index26
			(_200_v78).ArraySet1(_199_v77, (_index26).Int())
			var _201_v79 _dafny.MultiSet
			_ = _201_v79
			_201_v79 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(665))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_202_v71 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_203_i8 _dafny.Int) _dafny.CodePoint {
					return _202_v71
				}
			})(_186_v71))), _97_v5)
			var _204_v80 *C7
			_ = _204_v80
			var _nw34 *C7 = New_C7_()
			_ = _nw34
			_nw34.Ctor__()
			_204_v80 = _nw34
			var _205_v81 _dafny.Map
			_ = _205_v81
			_205_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, _204_v80)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_195_v73), 0))
			_ = _index27
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_200_v78), 0))
			_ = _index28
			var _rhs12 D2 = _198_v76
			_ = _rhs12
			var _rhs13 *C1 = _199_v77
			_ = _rhs13
			var _rhs14 bool = _94_v2
			_ = _rhs14
			var _rhs15 bool = _94_v2
			_ = _rhs15
			var _rhs16 _dafny.Int = ((func() _dafny.Int {
				if (_201_v79).Contains(_97_v5) {
					return (_201_v79).Multiplicity(_97_v5)
				}
				return _92_v0
			})()).Minus(Companion_Default___.SafeModuloInt((_205_v81).Cardinality(), _92_v0))
			_ = _rhs16
			var _lhs7 _dafny.Array = _195_v73
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_195_v73), 0))
			_ = _lhs8
			var _lhs9 _dafny.Array = _200_v78
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_200_v78), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = _100_globalState
			_ = _lhs11
			var _lhs12 *GlobalState = _100_globalState
			_ = _lhs12
			(_lhs7).ArraySet1(_rhs12, (_lhs8).Int())
			(_lhs9).ArraySet1(_rhs13, (_lhs10).Int())
			_94_v2 = _rhs14
			_lhs11.F4 = _rhs15
			_lhs12.F8 = _rhs16
			var _206_v82 _dafny.Map
			_ = _206_v82
			_206_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _171_v60)
			var _207_v83 D0
			_ = _207_v83
			_207_v83 = Companion_D0_.Create_DC2_()
			var _208_v84 _dafny.Map
			_ = _208_v84
			_208_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_206_v82).Merge(_206_v82), _207_v83)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))
			_ = _index29
			var _rhs17 bool = _94_v2
			_ = _rhs17
			var _rhs18 bool = Companion_Default___.Fm49(_94_v2, _100_globalState)
			_ = _rhs18
			var _rhs19 _dafny.Int = Companion_Default___.SafeDivisionInt((_101_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))).Int()).(_dafny.Int), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(104), _92_v0))
			_ = _rhs19
			var _rhs20 bool = !(_94_v2)
			_ = _rhs20
			var _rhs21 _dafny.Int = (_208_v84).Cardinality()
			_ = _rhs21
			var _lhs13 *GlobalState = _100_globalState
			_ = _lhs13
			var _lhs14 _dafny.Array = _101_v8
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_101_v8), 0))
			_ = _lhs15
			var _lhs16 *GlobalState = _100_globalState
			_ = _lhs16
			var _lhs17 *GlobalState = _100_globalState
			_ = _lhs17
			_lhs13.F4 = _rhs17
			_94_v2 = _rhs18
			(_lhs14).ArraySet1(_rhs19, (_lhs15).Int())
			_lhs16.F4 = _rhs20
			_lhs17.F3 = _rhs21
			var _209_v85 _dafny.MultiSet
			_ = _209_v85
			_209_v85 = _dafny.MultiSetOf(Companion_Default___.Fm49(_94_v2, _100_globalState))
			var _210_v86 _dafny.Sequence
			_ = _210_v86
			_210_v86 = _dafny.SeqOf(_97_v5)
			var _211_v87 _dafny.Map
			_ = _211_v87
			_211_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_209_v85).Contains(_94_v2) {
					return (_209_v85).Multiplicity(_94_v2)
				}
				return _92_v0
			})(), (_dafny.IntOfInt64(-212)).Times(_dafny.IntOfUint32(((_210_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jjqbikwfu")).Cardinality()), _dafny.IntOfUint32((_210_v86).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
			var _rhs22 bool = _94_v2
			_ = _rhs22
			var _rhs23 _dafny.Map = ((func() _dafny.Map {
				if _94_v2 {
					return _211_v87
				}
				return _211_v87
			})()).Update(_92_v0, _92_v0)
			_ = _rhs23
			var _lhs18 *GlobalState = _100_globalState
			_ = _lhs18
			_lhs18.F4 = _rhs22
			_211_v87 = _rhs23
		} else {
			var _212_v88 _dafny.Array
			_ = _212_v88
			_212_v88 = _101_v8
			var _213_v89 _dafny.Sequence
			_ = _213_v89
			_213_v89 = _dafny.SeqOf(_101_v8)
			var _214_v90 _dafny.MultiSet
			_ = _214_v90
			_214_v90 = _dafny.MultiSetOf(_94_v2)
			var _215_v91 _dafny.Array
			_ = _215_v91
			var _nwElement0_8 _dafny.Array = (func() _dafny.Array {
				if _94_v2 {
					return _101_v8
				}
				return _101_v8
			})()
			_ = _nwElement0_8
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_8, 0)
			(_nw35).ArraySet1((_212_v88), 1)
			(_nw35).ArraySet1(_101_v8, 2)
			(_nw35).ArraySet1(_101_v8, 3)
			(_nw35).ArraySet1(_101_v8, 4)
			(_nw35).ArraySet1(_101_v8, 5)
			(_nw35).ArraySet1(_101_v8, 6)
			(_nw35).ArraySet1((_213_v89).Select((Companion_Default___.SafeIndex((_214_v90).Cardinality(), _dafny.IntOfUint32((_213_v89).Cardinality()))).Uint32()).(_dafny.Array), 7)
			_215_v91 = _nw35
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_215_v91), 0))
			_ = _index30
			(_215_v91).ArraySet1(_101_v8, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_215_v91), 0))
			_ = _index31
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw36
			(_215_v91).ArraySet1(_nw36, (_index31).Int())
			var _216_v93 _dafny.Map
			_ = _216_v93
			_216_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_92_v0, _92_v0, (_dafny.Zero).Minus((func() _dafny.Set {
				var _coll54 = _dafny.NewBuilder()
				_ = _coll54
				for _iter55 := _dafny.Iterate((_132_v24).Elements()); ; {
					_compr_54, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _217_v92 _dafny.Int
					_217_v92 = interface{}(_compr_54).(_dafny.Int)
					if (_132_v24).Contains(_217_v92) {
						_coll54.Add((_217_v92).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(-834)))).Cardinality()))
					}
				}
				return _coll54.ToSet()
			}()).Cardinality()), _dafny.IntOfInt64(605))).Union(_132_v24), _94_v2)
			_216_v93 = (_216_v93).Merge(_216_v93)
			var _218_v94 _dafny.Sequence
			_ = _218_v94
			_218_v94 = _dafny.SeqOf(_94_v2, _94_v2)
			(_100_globalState).F8 = (_92_v0).Minus((func() _dafny.Int {
				if (_214_v90).Contains((_218_v94).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_218_v94).Cardinality()))).Uint32()).(bool)) {
					return (_214_v90).Multiplicity((_218_v94).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_218_v94).Cardinality()))).Uint32()).(bool))
				}
				return _92_v0
			})())
			var _219_v95 *C9
			_ = _219_v95
			var _nw37 *C9 = New_C9_()
			_ = _nw37
			_nw37.Ctor__((false) && (_94_v2))
			_219_v95 = _nw37
			var _220_v96 *C8
			_ = _220_v96
			var _nw38 *C8 = New_C8_()
			_ = _nw38
			_nw38.Ctor__()
			_220_v96 = _nw38
			_220_v96 = _220_v96
		}
	}
	var _221_v97 *C0
	_ = _221_v97
	var _nw39 *C0 = New_C0_()
	_ = _nw39
	_nw39.Ctor__()
	_221_v97 = _nw39
	var _222_v98 _dafny.Map
	_ = _222_v98
	_222_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, _92_v0)
	var _223_v99 D3
	_ = _223_v99
	_223_v99 = Companion_D3_.Create_DC8_(_222_v98)
	var _224_v100 _dafny.Sequence
	_ = _224_v100
	_224_v100 = _dafny.SeqOf(_94_v2)
	var _225_v101 _dafny.Set
	_ = _225_v101
	_225_v101 = _dafny.SetOf(_94_v2)
	var _226_i9 _dafny.Int
	_ = _226_i9
	_226_i9 = _dafny.Zero
	{
		for Companion_Default___.Fm49((_225_v101).IsSubsetOf(Companion_Default___.Fm35(_94_v2, Companion_Default___.Fm26(_223_v99, _92_v0, _dafny.MultiSetOf(_224_v100, _224_v100, _224_v100), _100_globalState), _94_v2, true, _100_globalState)), _100_globalState) {
			{
				if (_226_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_226_i9 = (_226_i9).Plus(_dafny.One)
				var _227_v102 _dafny.Array
				_ = _227_v102
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
				_ = _nw40
				_227_v102 = _nw40
				var _228_v103 *C4
				_ = _228_v103
				var _nw41 *C4 = New_C4_()
				_ = _nw41
				_nw41.Ctor__(_227_v102)
				_228_v103 = _nw41
				var _229_v104 _dafny.Sequence
				_ = _229_v104
				_229_v104 = _dafny.SeqOf(_228_v103)
				_228_v103 = (_229_v104).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.IntOfUint32((_229_v104).Cardinality()))).Uint32()).(*C4)
				(_100_globalState).F8 = (func() _dafny.Int {
					if !(_94_v2) {
						return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-265), _92_v0)
					}
					return (_92_v0).Plus(_92_v0)
				})()
				var _230_v105 T0
				_ = _230_v105
				var _nw42 *C7 = New_C7_()
				_ = _nw42
				_nw42.Ctor__()
				_230_v105 = _nw42
				var _231_v106 _dafny.CodePoint
				_ = _231_v106
				_231_v106 = _dafny.CodePoint('w')
				var _232_v107 _dafny.Set
				_ = _232_v107
				_232_v107 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v105, _231_v106))
				var _233_v108 _dafny.Map
				_ = _233_v108
				_233_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v105, _dafny.CodePoint('q'))
				_232_v107 = _dafny.SetOf(_233_v108, _233_v108, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v105, _dafny.CodePoint('y')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v105, _231_v106))
				(_228_v103).M9(Companion_Default___.Fm6(_92_v0, _100_globalState), _92_v0, _100_globalState)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _234_v109 D3
	_ = _234_v109
	_234_v109 = Companion_D3_.Create_DC9_()
	_234_v109 = _234_v109
	var _235_v110 D21
	_ = _235_v110
	_235_v110 = Companion_D21_.Create_DC43_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(678))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}((func(_236_v2 bool) func(_dafny.Int) _dafny.Int {
		return func(_237_i11 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _236_v2)).Cardinality()
		}
	})(_94_v2))), false)
	var _238_i10 _dafny.Int
	_ = _238_i10
	_238_i10 = _dafny.Zero
	{
		var _pat_let_tv5 = _92_v0
		_ = _pat_let_tv5
		var _pat_let_tv6 = _92_v0
		_ = _pat_let_tv6
		var _pat_let_tv7 = _94_v2
		_ = _pat_let_tv7
		var _pat_let_tv8 = _94_v2
		_ = _pat_let_tv8
		var _pat_let_tv9 = _100_globalState
		_ = _pat_let_tv9
		for func(_source9 D21) bool {
			if _source9.Is_DC43() {
				var _288___mcc_h11 _dafny.Sequence = _source9.Get_().(D21_DC43).Cf63
				_ = _288___mcc_h11
				var _289___mcc_h12 bool = _source9.Get_().(D21_DC43).Cf64
				_ = _289___mcc_h12
				var _290_cf64 bool = _289___mcc_h12
				_ = _290_cf64
				var _291_cf63 _dafny.Sequence = _288___mcc_h11
				_ = _291_cf63
				return (_pat_let_tv5).Cmp(_pat_let_tv6) >= 0
			} else {
				var _292___mcc_h13 _dafny.Map = _source9.Get_().(D21_DC42).Cf62
				_ = _292___mcc_h13
				var _293_cf62 _dafny.Map = _292___mcc_h13
				_ = _293_cf62
				return _pat_let_tv7
			}
		}(func(_pat_let7_0 D21) D21 {
			return func(_294_dt__update__tmp_h3 D21) D21 {
				return func(_pat_let8_0 bool) D21 {
					return func(_295_dt__update_hcf64_h0 bool) D21 {
						return Companion_D21_.Create_DC43_((_294_dt__update__tmp_h3).Dtor_cf63(), _295_dt__update_hcf64_h0)
					}(_pat_let8_0)
				}(Companion_Default___.Fm49(_pat_let_tv8, _pat_let_tv9))
			}(_pat_let7_0)
		}(_235_v110)) {
			{
				if (_238_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_238_i10 = (_238_i10).Plus(_dafny.One)
				if _94_v2 {
					var _239_v111 _dafny.Array
					_ = _239_v111
					var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
					_ = _nw43
					_239_v111 = _nw43
					var _240_v112 *C4
					_ = _240_v112
					var _nw44 *C4 = New_C4_()
					_ = _nw44
					_nw44.Ctor__(_239_v111)
					_240_v112 = _nw44
					(_100_globalState).F8 = (func() _dafny.Int {
						if _94_v2 {
							return _92_v0
						}
						return (_dafny.Zero).Minus(_92_v0)
					})()
					var _241_v113 _dafny.Map
					_ = _241_v113
					_241_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v0, (_240_v112).Fm3(_93_v1, _dafny.MultiSetOf(_92_v0), !(_94_v2), _100_globalState))
					var _242_v114 _dafny.Set
					_ = _242_v114
					_242_v114 = _dafny.SetOf(_241_v113, _241_v113, _241_v113, (_241_v113).Merge(_241_v113), _241_v113)
					var _243_v115 T0
					_ = _243_v115
					var _nw45 *C2 = New_C2_()
					_ = _nw45
					_nw45.Ctor__()
					_243_v115 = _nw45
					var _244_v116 _dafny.Sequence
					_ = _244_v116
					_244_v116 = _dafny.SeqOf(_243_v115, _243_v115)
					var _245_v117 _dafny.CodePoint
					_ = _245_v117
					_245_v117 = _dafny.CodePoint('y')
					var _246_v118 *C1
					_ = _246_v118
					var _nw46 *C1 = New_C1_()
					_ = _nw46
					_nw46.Ctor__()
					_246_v118 = _nw46
					var _247_v119 _dafny.Map
					_ = _247_v119
					_247_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_246_v118, _dafny.IntOfInt64(110))
					var _rhs24 bool = _dafny.Companion_Sequence_.IsPrefixOf(_97_v5, Companion_Default___.Fm25(_100_globalState))
					_ = _rhs24
					var _rhs25 _dafny.Set = (_dafny.SetOf(_241_v113)).Intersection(Companion_Default___.Fm53(_92_v0, _92_v0, _100_globalState))
					_ = _rhs25
					var _rhs26 T0 = (_244_v116).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_244_v116).Cardinality()))).Uint32()).(T0)
					_ = _rhs26
					var _rhs27 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ukahennov"), (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ukahennov")).Cardinality()))).Uint32(), _245_v117), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uvcx"), _97_v5)), (Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_247_v119).Contains(_246_v118) {
							return (_247_v119).Get(_246_v118).(_dafny.Int)
						}
						return (_dafny.Zero).Minus(_92_v0)
					})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ukahennov"), (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ukahennov")).Cardinality()))).Uint32(), _245_v117), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uvcx"), _97_v5))).Cardinality()))).Uint32(), _245_v117)
					_ = _rhs27
					_94_v2 = _rhs24
					_242_v114 = _rhs25
					_243_v115 = _rhs26
					_97_v5 = _rhs27
					var _248_v120 _dafny.Map
					_ = _248_v120
					_248_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gofjuqlqk"), _94_v2)
					(_100_globalState).F3 = (_248_v120).Cardinality()
					_92_v0 = _dafny.IntOfInt64(-256)
				} else {
					var _249_v121 _dafny.Array
					_ = _249_v121
					var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
					_ = _nw47
					_249_v121 = _nw47
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_249_v121), 0))
					_ = _index32
					(_249_v121).ArraySet1(_94_v2, (_index32).Int())
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_249_v121), 0))
					_ = _index33
					(_249_v121).ArraySet1(_94_v2, (_index33).Int())
					var _250_v122 _dafny.Int
					_ = _250_v122
					var _251_v123 _dafny.Map
					_ = _251_v123
					var _252_v124 bool
					_ = _252_v124
					var _253_v125 *C0
					_ = _253_v125
					var _out8 _dafny.Int
					_ = _out8
					var _out9 _dafny.Map
					_ = _out9
					var _out10 bool
					_ = _out10
					var _out11 *C0
					_ = _out11
					_out8, _out9, _out10, _out11 = Companion_Default___.M15(_100_globalState)
					_250_v122 = _out8
					_251_v123 = _out9
					_252_v124 = _out10
					_253_v125 = _out11
					var _254_v126 _dafny.Map
					_ = _254_v126
					_254_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_252_v124, (_249_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_249_v121), 0))).Int()).(bool))
					(_100_globalState).F4 = !(_254_v126).Equals(_254_v126)
					var _255_v127 _dafny.Int
					_ = _255_v127
					var _256_v128 _dafny.Map
					_ = _256_v128
					var _257_v129 bool
					_ = _257_v129
					var _258_v130 *C0
					_ = _258_v130
					var _out12 _dafny.Int
					_ = _out12
					var _out13 _dafny.Map
					_ = _out13
					var _out14 bool
					_ = _out14
					var _out15 *C0
					_ = _out15
					_out12, _out13, _out14, _out15 = Companion_Default___.M15(_100_globalState)
					_255_v127 = _out12
					_256_v128 = _out13
					_257_v129 = _out14
					_258_v130 = _out15
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_101_v8), 0))
					_ = _index34
					(_101_v8).ArraySet1(_250_v122, (_index34).Int())
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_101_v8), 0))
					_ = _index35
					(_101_v8).ArraySet1(_255_v127, (_index35).Int())
				}
				var _rhs28 _dafny.Int = _dafny.IntOfInt64(-265)
				_ = _rhs28
				var _rhs29 _dafny.Int = (_dafny.Zero).Minus(_92_v0)
				_ = _rhs29
				var _lhs19 *GlobalState = _100_globalState
				_ = _lhs19
				var _lhs20 *GlobalState = _100_globalState
				_ = _lhs20
				_lhs19.F8 = _rhs28
				_lhs20.F8 = _rhs29
				var _259_v131 T0
				_ = _259_v131
				var _nw48 *C7 = New_C7_()
				_ = _nw48
				_nw48.Ctor__()
				_259_v131 = _nw48
				var _260_v132 _dafny.Sequence
				_ = _260_v132
				_260_v132 = _dafny.SeqOf(_259_v131)
				var _261_v133 _dafny.Map
				_ = _261_v133
				_261_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _94_v2)
				var _262_v134 _dafny.MultiSet
				_ = _262_v134
				_262_v134 = _dafny.MultiSetOf(_94_v2)
				var _263_v135 _dafny.Array
				_ = _263_v135
				var _nwElement0_9 bool = (_132_v24).IsDisjointFrom(_132_v24)
				_ = _nwElement0_9
				var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(29))
				_ = _nw49
				(_nw49).ArraySet1(_nwElement0_9, 0)
				(_nw49).ArraySet1(_dafny.Companion_Sequence_.Equal(_224_v100, _dafny.SeqOf(_94_v2, _94_v2, _94_v2, _94_v2, _94_v2)), 1)
				(_nw49).ArraySet1(_94_v2, 2)
				(_nw49).ArraySet1((func() bool {
					if (_261_v133).Contains(_94_v2) {
						return (_261_v133).Get(_94_v2).(bool)
					}
					return _94_v2
				})(), 3)
				(_nw49).ArraySet1(_94_v2, 4)
				(_nw49).ArraySet1(_94_v2, 5)
				(_nw49).ArraySet1(_dafny.Companion_Sequence_.Equal(_224_v100, _224_v100), 6)
				(_nw49).ArraySet1(false, 7)
				(_nw49).ArraySet1(_94_v2, 8)
				(_nw49).ArraySet1(true, 9)
				(_nw49).ArraySet1((func() bool {
					if _94_v2 {
						return _94_v2
					}
					return _94_v2
				})(), 10)
				(_nw49).ArraySet1(_94_v2, 11)
				(_nw49).ArraySet1(_94_v2, 12)
				(_nw49).ArraySet1(_94_v2, 13)
				(_nw49).ArraySet1((_259_v131).Fm3(_dafny.SeqOf(_dafny.IntOfUint32((_97_v5).Cardinality())), _99_v7, _94_v2, _100_globalState), 14)
				(_nw49).ArraySet1(_94_v2, 15)
				(_nw49).ArraySet1(_94_v2, 16)
				(_nw49).ArraySet1(_94_v2, 17)
				(_nw49).ArraySet1(_94_v2, 18)
				(_nw49).ArraySet1(_94_v2, 19)
				(_nw49).ArraySet1(_94_v2, 20)
				(_nw49).ArraySet1(_94_v2, 21)
				(_nw49).ArraySet1(!(_94_v2) || (_94_v2), 22)
				(_nw49).ArraySet1(_94_v2, 23)
				(_nw49).ArraySet1(_94_v2, 24)
				(_nw49).ArraySet1(false, 25)
				(_nw49).ArraySet1((_dafny.MultiSetOf(_94_v2)).IsDisjointFrom(_262_v134), 26)
				(_nw49).ArraySet1(_94_v2, 27)
				(_nw49).ArraySet1(_94_v2, 28)
				_263_v135 = _nw49
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))
				_ = _index36
				(_263_v135).ArraySet1(!(!(_94_v2) || (_94_v2)), (_index36).Int())
				var _264_v137 _dafny.Array
				_ = _264_v137
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_6
				var _nw50 _dafny.Array
				_ = _nw50
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw50 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Map = (func(_265_v2 bool) func(_dafny.Int) _dafny.Map {
						return func(_266_i12 _dafny.Int) _dafny.Map {
							return func() _dafny.Map {
								var _coll55 = _dafny.NewMapBuilder()
								_ = _coll55
								for _iter56 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('w'))).Elements()); ; {
									_compr_55, _ok56 := _iter56()
									if !_ok56 {
										break
									}
									var _267_v136 _dafny.CodePoint
									_267_v136 = interface{}(_compr_55).(_dafny.CodePoint)
									if (_dafny.MultiSetOf(_dafny.CodePoint('w'))).Contains(_267_v136) {
										_coll55.Add(_267_v136, _265_v2)
									}
								}
								return _coll55.ToMap()
							}()
						}
					})(_94_v2)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw50 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw50).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw50).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_264_v137 = _nw50
				var _268_v138 _dafny.CodePoint
				_ = _268_v138
				_268_v138 = _dafny.CodePoint('q')
				var _269_v139 _dafny.Map
				_ = _269_v139
				_269_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v138, _94_v2)
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_264_v137), 0))
				_ = _index37
				(_264_v137).ArraySet1((_269_v139).Update(_268_v138, _94_v2), (_index37).Int())
				var _270_v140 D27
				_ = _270_v140
				_270_v140 = Companion_D27_.Create_DC58_(_269_v139)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))
				_ = _index38
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_264_v137), 0))
				_ = _index39
				var _rhs30 _dafny.Int = _92_v0
				_ = _rhs30
				var _rhs31 _dafny.Sequence = _dafny.SeqOf((func() T0 {
					if _94_v2 {
						return _259_v131
					}
					return _259_v131
				})(), _259_v131, _259_v131, _259_v131, _259_v131)
				_ = _rhs31
				var _rhs32 bool = false
				_ = _rhs32
				var _rhs33 bool = (_259_v131).Fm3(_93_v1, _99_v7, _94_v2, _100_globalState)
				_ = _rhs33
				var _rhs34 _dafny.Map = (_270_v140).Dtor_cf82()
				_ = _rhs34
				var _lhs21 *GlobalState = _100_globalState
				_ = _lhs21
				var _lhs22 _dafny.Array = _263_v135
				_ = _lhs22
				var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))
				_ = _lhs23
				var _lhs24 _dafny.Array = _264_v137
				_ = _lhs24
				var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_264_v137), 0))
				_ = _lhs25
				_lhs21.F8 = _rhs30
				_260_v132 = _rhs31
				_94_v2 = _rhs32
				(_lhs22).ArraySet1(_rhs33, (_lhs23).Int())
				(_lhs24).ArraySet1(_rhs34, (_lhs25).Int())
				var _271_v141 D2
				_ = _271_v141
				_271_v141 = Companion_D2_.Create_DC5_(_263_v135)
				var _source8 D2 = _271_v141
				_ = _source8
				if _source8.Is_DC6() {
					var _272___mcc_h4 _dafny.Int = _source8.Get_().(D2_DC6).Cf6
					_ = _272___mcc_h4
					var _273___mcc_h5 bool = _source8.Get_().(D2_DC6).Cf7
					_ = _273___mcc_h5
					var _274___mcc_h6 bool = _source8.Get_().(D2_DC6).Cf8
					_ = _274___mcc_h6
					var _275_cf8 bool = _274___mcc_h6
					_ = _275_cf8
					var _276_cf7 bool = _273___mcc_h5
					_ = _276_cf7
					var _277_cf6 _dafny.Int = _272___mcc_h4
					_ = _277_cf6
					_95_v3 = (_95_v3).Update(_276_cf7, _277_cf6)
					(_100_globalState).F8 = Companion_Default___.SafeDivisionInt(_92_v0, _92_v0)
					var _278_v142 _dafny.Array
					_ = _278_v142
					var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
					_ = _nw51
					_278_v142 = _nw51
					_278_v142 = _278_v142
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_101_v8), 0))
					_ = _index40
					(_101_v8).ArraySet1(_92_v0, (_index40).Int())
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_101_v8), 0))
					_ = _index41
					(_101_v8).ArraySet1(_277_cf6, (_index41).Int())
				} else if _source8.Is_DC7() {
					var _279___mcc_h7 _dafny.CodePoint = _source8.Get_().(D2_DC7).Cf9
					_ = _279___mcc_h7
					var _280___mcc_h8 bool = _source8.Get_().(D2_DC7).Cf10
					_ = _280___mcc_h8
					var _281___mcc_h9 _dafny.Int = _source8.Get_().(D2_DC7).Cf11
					_ = _281___mcc_h9
					var _282_cf11 _dafny.Int = _281___mcc_h9
					_ = _282_cf11
					var _283_cf10 bool = _280___mcc_h8
					_ = _283_cf10
					var _284_cf9 _dafny.CodePoint = _279___mcc_h7
					_ = _284_cf9
					_283_cf10 = (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool)
					(_100_globalState).F4 = (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool)
					_261_v133 = (_261_v133).Update(!((_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool)) || (_283_cf10), (_92_v0).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dmy")).Cardinality())) >= 0)
					(_100_globalState).F8 = (_dafny.Zero).Minus((func() _dafny.Int {
						if (_95_v3).Contains(_283_cf10) {
							return (_95_v3).Get(_283_cf10).(_dafny.Int)
						}
						return _92_v0
					})())
				} else {
					var _285___mcc_h10 _dafny.Array = _source8.Get_().(D2_DC5).Cf5
					_ = _285___mcc_h10
					var _286_cf5 _dafny.Array = _285___mcc_h10
					_ = _286_cf5
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))
					_ = _index42
					(_263_v135).ArraySet1((func() bool {
						if (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool) {
							return (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool)
						}
						return (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool)
					})(), (_index42).Int())
					_101_v8 = _101_v8
					var _287_v143 D27
					_ = _287_v143
					_287_v143 = Companion_D27_.Create_DC59_(_94_v2, _263_v135, (_263_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_263_v135), 0))).Int()).(bool))
					(_100_globalState).F4 = (_287_v143).Dtor_cf83()
					(_100_globalState).F3 = Companion_Default___.SafeDivisionInt(_92_v0, _92_v0)
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _296_v144 _dafny.Array
	_ = _296_v144
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_7
	var _nw52 _dafny.Array
	_ = _nw52
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw52 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) bool = (func(_297_v2 bool) func(_dafny.Int) bool {
			return func(_298_i13 _dafny.Int) bool {
				return _297_v2
			}
		})(_94_v2)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw52 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw52).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw52).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_296_v144 = _nw52
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
	_ = _index43
	(_296_v144).ArraySet1(_94_v2, (_index43).Int())
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
	_ = _index44
	(_296_v144).ArraySet1(_94_v2, (_index44).Int())
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if _94_v2 {
			return _92_v0
		}
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(250))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_299_i15 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))).Cardinality())
	})()
	_ = _hi1
	for _300_i14 := _92_v0; _300_i14.Cmp(_hi1) < 0; _300_i14 = _300_i14.Plus(_dafny.One) {
		var _301_v145 _dafny.Sequence
		_ = _301_v145
		_301_v145 = _dafny.SeqOf(_296_v144)
		(_100_globalState).F3 = _dafny.IntOfUint32((_dafny.SeqOf((_301_v145).Select((Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_301_v145).Cardinality()))).Uint32()).(_dafny.Array), _296_v144)).Cardinality())
		var _302_v146 _dafny.CodePoint
		_ = _302_v146
		_302_v146 = _dafny.CodePoint('y')
		_302_v146 = _dafny.CodePoint('l')
		_99_v7 = _99_v7
		if !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm23(_dafny.Companion_Sequence_.Update(_97_v5, (Companion_Default___.SafeIndex(_92_v0, _dafny.IntOfUint32((_97_v5).Cardinality()))).Uint32(), _302_v146), _100_globalState), !(false)) {
			var _303_v147 _dafny.Array
			_ = _303_v147
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw53
			_303_v147 = _nw53
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_303_v147), 0))
			_ = _index45
			(_303_v147).ArraySet1(_101_v8, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_303_v147), 0))
			_ = _index46
			(_303_v147).ArraySet1(_101_v8, (_index46).Int())
			var _304_v148 *C5
			_ = _304_v148
			var _nw54 *C5 = New_C5_()
			_ = _nw54
			_nw54.Ctor__()
			_304_v148 = _nw54
			var _305_v149 _dafny.Array
			_ = _305_v149
			var _nwElement0_10 *C5 = _304_v148
			_ = _nwElement0_10
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.One)
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_10, 0)
			_305_v149 = _nw55
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_305_v149), 0))
			_ = _index47
			(_305_v149).ArraySet1((func() *C5 {
				if (_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool) {
					return _304_v148
				}
				return _304_v148
			})(), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_305_v149), 0))
			_ = _index48
			(_305_v149).ArraySet1(_304_v148, (_index48).Int())
			(_100_globalState).F4 = (_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool)
			var _306_v150 _dafny.MultiSet
			_ = _306_v150
			_306_v150 = _dafny.MultiSetOf(_94_v2)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
			_ = _index49
			var _rhs35 _dafny.Int = (func() _dafny.Int {
				if !((_300_i14).Cmp(_300_i14) < 0) {
					return _92_v0
				}
				return _300_i14
			})()
			_ = _rhs35
			var _rhs36 bool = ((_dafny.MultiSetFromSeq(_224_v100)).Difference(_306_v150)).IsProperSubsetOf((_306_v150).Update((_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), Companion_Default___.Abs(_300_i14)))
			_ = _rhs36
			var _rhs37 _dafny.Int = (_92_v0).Minus(_92_v0)
			_ = _rhs37
			var _rhs38 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_224_v100, _224_v100), _224_v100)
			_ = _rhs38
			var _rhs39 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_224_v100, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_92_v0), _dafny.IntOfUint32((_224_v100).Cardinality()))).Uint32(), _94_v2)).Cardinality())
			_ = _rhs39
			var _lhs26 *GlobalState = _100_globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _296_v144
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
			_ = _lhs28
			var _lhs29 *GlobalState = _100_globalState
			_ = _lhs29
			var _lhs30 *GlobalState = _100_globalState
			_ = _lhs30
			_lhs26.F3 = _rhs35
			(_lhs27).ArraySet1(_rhs36, (_lhs28).Int())
			_lhs29.F3 = _rhs37
			_224_v100 = _rhs38
			_lhs30.F3 = _rhs39
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
			_ = _index50
			(_296_v144).ArraySet1((_306_v150).IsDisjointFrom(_306_v150), (_index50).Int())
		} else {
			var _307_v151 *C8
			_ = _307_v151
			var _nw56 *C8 = New_C8_()
			_ = _nw56
			_nw56.Ctor__()
			_307_v151 = _nw56
			var _308_v152 _dafny.Sequence
			_ = _308_v152
			_308_v152 = _dafny.SeqOf(_307_v151, _307_v151)
			_307_v151 = (_308_v152).Select((Companion_Default___.SafeIndex(_300_i14, _dafny.IntOfUint32((_308_v152).Cardinality()))).Uint32()).(*C8)
			(_100_globalState).F3 = _92_v0
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
			_ = _index51
			(_296_v144).ArraySet1(Companion_Default___.Fm49(_94_v2, _100_globalState), (_index51).Int())
			var _309_v153 *C3
			_ = _309_v153
			var _nw57 *C3 = New_C3_()
			_ = _nw57
			_nw57.Ctor__()
			_309_v153 = _nw57
			_309_v153 = _309_v153
			(_100_globalState).F8 = _92_v0
		}
	}
	var _310_v154 *C8
	_ = _310_v154
	var _nw58 *C8 = New_C8_()
	_ = _nw58
	_nw58.Ctor__()
	_310_v154 = _nw58
	var _311_v155 _dafny.Array
	_ = _311_v155
	var _nwElement0_11 *C8 = _310_v154
	_ = _nwElement0_11
	var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(3))
	_ = _nw59
	(_nw59).ArraySet1(_nwElement0_11, 0)
	(_nw59).ArraySet1(_310_v154, 1)
	(_nw59).ArraySet1(_310_v154, 2)
	_311_v155 = _nw59
	var _312_v156 _dafny.Array
	_ = _312_v156
	var _nwElement0_12 _dafny.Array = _311_v155
	_ = _nwElement0_12
	var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
	_ = _nw60
	(_nw60).ArraySet1(_nwElement0_12, 0)
	(_nw60).ArraySet1(_311_v155, 1)
	(_nw60).ArraySet1(_311_v155, 2)
	(_nw60).ArraySet1(_311_v155, 3)
	(_nw60).ArraySet1(_311_v155, 4)
	(_nw60).ArraySet1(_311_v155, 5)
	(_nw60).ArraySet1(_311_v155, 6)
	(_nw60).ArraySet1(_311_v155, 7)
	(_nw60).ArraySet1(_311_v155, 8)
	(_nw60).ArraySet1(_311_v155, 9)
	(_nw60).ArraySet1(_311_v155, 10)
	(_nw60).ArraySet1(_311_v155, 11)
	(_nw60).ArraySet1(_311_v155, 12)
	(_nw60).ArraySet1(_311_v155, 13)
	(_nw60).ArraySet1(_311_v155, 14)
	(_nw60).ArraySet1(_311_v155, 15)
	(_nw60).ArraySet1(_311_v155, 16)
	_312_v156 = _nw60
	var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_312_v156), 0))
	_ = _index52
	(_312_v156).ArraySet1(_311_v155, (_index52).Int())
	var _313_v157 _dafny.Map
	_ = _313_v157
	_313_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), _224_v100)
	var _314_v158 _dafny.MultiSet
	_ = _314_v158
	_314_v158 = _dafny.MultiSetOf((func() _dafny.Sequence {
		if (_313_v157).Contains(_94_v2) {
			return (_313_v157).Get(_94_v2).(_dafny.Sequence)
		}
		return _224_v100
	})())
	var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_312_v156), 0))
	_ = _index53
	var _rhs40 bool = false
	_ = _rhs40
	var _rhs41 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("kmhjp")
	_ = _rhs41
	var _rhs42 _dafny.Int = Companion_Default___.Fm26(_223_v99, _dafny.IntOfUint32((_224_v100).Cardinality()), _314_v158, _100_globalState)
	_ = _rhs42
	var _rhs43 _dafny.Array = _311_v155
	_ = _rhs43
	var _lhs31 *GlobalState = _100_globalState
	_ = _lhs31
	var _lhs32 *GlobalState = _100_globalState
	_ = _lhs32
	var _lhs33 _dafny.Array = _312_v156
	_ = _lhs33
	var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_312_v156), 0))
	_ = _lhs34
	_lhs31.F4 = _rhs40
	_97_v5 = _rhs41
	_lhs32.F3 = _rhs42
	(_lhs33).ArraySet1(_rhs43, (_lhs34).Int())
	(_100_globalState).F4 = !(_94_v2)
	(_100_globalState).F3 = (_92_v0).Times(_92_v0)
	(_100_globalState).F4 = true
	var _315_v159 D2
	_ = _315_v159
	_315_v159 = Companion_D2_.Create_DC6_(_92_v0, (_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), (_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool))
	var _hi2 _dafny.Int = (_92_v0).Minus((_315_v159).Dtor_cf6())
	_ = _hi2
	for _316_i16 := _dafny.IntOfUint32((_224_v100).Cardinality()); _316_i16.Cmp(_hi2) < 0; _316_i16 = _316_i16.Plus(_dafny.One) {
		if _94_v2 {
			var _317_v160 _dafny.CodePoint
			_ = _317_v160
			_317_v160 = _dafny.CodePoint('x')
			var _318_v161 _dafny.Map
			_ = _318_v161
			_318_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_317_v160, _97_v5)
			var _319_v162 _dafny.Map
			_ = _319_v162
			_319_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_318_v161).Contains(_317_v160) {
					return (_318_v161).Get(_317_v160).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-780))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}((func(_320_v0 _dafny.Int, _321_v5 _dafny.Sequence, _322_v2 bool, _323_i16 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_324_i17 _dafny.Int) _dafny.CodePoint {
						return (Companion_D16_.Create_DC35_(_320_v0, (_321_v5).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_321_v5).Cardinality()))).Uint32()).(_dafny.CodePoint), (_dafny.SetOf(_322_v2, _322_v2)).Cardinality(), _323_i16, _320_v0)).Dtor_cf51()
					}
				})(_92_v0, _97_v5, _94_v2, _316_i16)))
			})(), _94_v2)
			var _325_v163 _dafny.Sequence
			_ = _325_v163
			_325_v163 = _dafny.SeqOf(_dafny.SetOf(_92_v0))
			_319_v162 = Companion_Default___.Fm42(_316_i16, _dafny.IntOfInt64(198), ((_325_v163).Select((Companion_Default___.SafeIndex(_316_i16, _dafny.IntOfUint32((_325_v163).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), true, _100_globalState)
			var _326_v164 _dafny.Set
			_ = _326_v164
			_326_v164 = _dafny.SetOf(_222_v98, (_222_v98).Merge(_222_v98))
			var _327_v165 _dafny.Map
			_ = _327_v165
			_327_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), _dafny.SetOf(_222_v98))
			_326_v164 = (func() _dafny.Set {
				if (_327_v165).Contains(!(_94_v2) || (true)) {
					return (_327_v165).Get(!(_94_v2) || (true)).(_dafny.Set)
				}
				return _326_v164
			})()
			(_100_globalState).F3 = _dafny.IntOfInt64(359)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))
			_ = _index54
			(_296_v144).ArraySet1((_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), (_index54).Int())
			var _328_v166 _dafny.MultiSet
			_ = _328_v166
			_328_v166 = _dafny.MultiSetOf(_97_v5, _97_v5)
			(_100_globalState).F3 = Companion_Default___.SafeModuloInt((_328_v166).Cardinality(), _316_i16)
		} else {
			(_100_globalState).F4 = (_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool)
			_95_v3 = (_95_v3).Update(_94_v2, (_98_v6).Cardinality())
			(_100_globalState).F4 = _94_v2
			var _329_v167 D13
			_ = _329_v167
			_329_v167 = Companion_D13_.Create_DC30_((_296_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_296_v144), 0))).Int()).(bool), _94_v2, _132_v24, _221_v97)
			(_100_globalState).F4 = !(!(_94_v2)) || ((_329_v167).Dtor_cf43())
			var _330_v168 _dafny.Map
			_ = _330_v168
			_330_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v2, _97_v5)
			(_310_v154).M4(_94_v2, _330_v168, _dafny.IntOfInt64(239), _100_globalState)
		}
		(_100_globalState).F4 = true
		(_100_globalState).F4 = true
		_101_v8 = _101_v8
	}
	var _331_v169 *C0
	_ = _331_v169
	var _nw61 *C0 = New_C0_()
	_ = _nw61
	_nw61.Ctor__()
	_331_v169 = _nw61
	var _332_v170 _dafny.MultiSet
	_ = _332_v170
	_332_v170 = _dafny.MultiSetOf(_97_v5)
	(_100_globalState).F4 = (_332_v170).IsProperSubsetOf((_332_v170).Union(_332_v170))
	_dafny.Print(_92_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_93_v1, _dafny.SeqOf(_dafny.IntOfInt64(545), _dafny.IntOfInt64(545), _dafny.IntOfInt64(545), _dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_94_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_95_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v4).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_v5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v6).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("cj"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v7).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_100_globalState.F0, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_globalState.F2).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_100_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_100_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_100_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v8).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v9).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v9).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v24).Equals(_dafny.SetOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v98).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(545), _dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_223_v99).Dtor_cf12()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(545), _dafny.IntOfInt64(545))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_224_v100, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v101).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_226_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_235_v110).Dtor_cf63(), _dafny.SeqOf(_dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v110).Dtor_cf64())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_238_i10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v144).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v157).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_314_v158).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v159).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v159).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v159).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_332_v170).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("kmhjp"))))
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
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
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
	Cf3 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf3 D0) D0 {
	return D0{D0_DC3{Cf3}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC3).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC2:
		{
			_, ok := other.Get_().(D0_DC2)
			return ok
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf4 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 bool) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() bool {
	return false
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf4 == data2.Cf4
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

type D2_DC6 struct {
	Cf6 _dafny.Int
	Cf7 bool
	Cf8 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 _dafny.Int, Cf7 bool, Cf8 bool) D2 {
	return D2{D2_DC6{Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf9  _dafny.CodePoint
	Cf10 bool
	Cf11 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf9 _dafny.CodePoint, Cf10 bool, Cf11 _dafny.Int) D2 {
	return D2{D2_DC7{Cf9, Cf10, Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf5 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 _dafny.Array) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, false, false)
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC6).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC6).Cf8
}

func (_this D2) Dtor_cf9() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf5 == data2.Cf5
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
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_() D3 {
	return D3{D3_DC9{}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf12 _dafny.Map
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 _dafny.Map) D3 {
	return D3{D3_DC8{Cf12}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_()
}

func (_this D3) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ")"
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
			_, ok := other.Get_().(D3_DC9)
			return ok
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf14 bool
	Cf15 bool
	Cf16 _dafny.Int
	Cf17 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf14 bool, Cf15 bool, Cf16 _dafny.Int, Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf14, Cf15, Cf16, Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf13 _dafny.Sequence
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf13 _dafny.Sequence) D4 {
	return D4{D4_DC10{Cf13}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC12 struct {
	Cf18 D4
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf18 D4) D4 {
	return D4{D4_DC12{Cf18}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(false, false, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC11).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC11).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf13
}

func (_this D4) Dtor_cf18() D4 {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + data.Cf17.VerbatimString(true) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Equals(data2.Cf17)
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 _dafny.Int
	Cf21 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 _dafny.Int, Cf21 bool) D5 {
	return D5{D5_DC14{Cf20, Cf21}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC15 struct {
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_() D5 {
	return D5{D5_DC15{}}
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

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.Zero, false)
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC14).Cf21
}

func (_this D5) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15"
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
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
	case D5_DC15:
		{
			_, ok := other.Get_().(D5_DC15)
			return ok
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf22 _dafny.MultiSet
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.MultiSet) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D6) Dtor_cf22() _dafny.MultiSet {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
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

type D7_DC18 struct {
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_() D7 {
	return D7{D7_DC18{}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC17 struct {
	Cf23 _dafny.Map
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf23 _dafny.Map) D7 {
	return D7{D7_DC17{Cf23}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_()
}

func (_this D7) Dtor_cf23() _dafny.Map {
	return _this.Get_().(D7_DC17).Cf23
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			_, ok := other.Get_().(D7_DC18)
			return ok
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
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

type D8_DC19 struct {
	Cf24 _dafny.Sequence
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf24 _dafny.Sequence) D8 {
	return D8{D8_DC19{Cf24}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D8_DC19).Cf24
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D9_DC21 struct {
	Cf26 bool
	Cf27 _dafny.Sequence
	Cf28 bool
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf26 bool, Cf27 _dafny.Sequence, Cf28 bool) D9 {
	return D9{D9_DC21{Cf26, Cf27, Cf28}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

type D9_DC22 struct {
	Cf29 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf29 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf29}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC20 struct {
	Cf25 _dafny.Set
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf25 _dafny.Set) D9 {
	return D9{D9_DC20{Cf25}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC21_(false, _dafny.EmptySeq, false)
}

func (_this D9) Dtor_cf26() bool {
	return _this.Get_().(D9_DC21).Cf26
}

func (_this D9) Dtor_cf27() _dafny.Sequence {
	return _this.Get_().(D9_DC21).Cf27
}

func (_this D9) Dtor_cf28() bool {
	return _this.Get_().(D9_DC21).Cf28
}

func (_this D9) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf29
}

func (_this D9) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D9_DC20).Cf25
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Equals(data2.Cf27) && data1.Cf28 == data2.Cf28
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D10_DC24 struct {
	Cf31 _dafny.Int
	Cf32 bool
	Cf33 _dafny.Sequence
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf31 _dafny.Int, Cf32 bool, Cf33 _dafny.Sequence) D10 {
	return D10{D10_DC24{Cf31, Cf32, Cf33}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

type D10_DC23 struct {
	Cf30 _dafny.MultiSet
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf30 _dafny.MultiSet) D10 {
	return D10{D10_DC23{Cf30}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC24_(_dafny.Zero, false, _dafny.EmptySeq)
}

func (_this D10) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D10_DC24).Cf31
}

func (_this D10) Dtor_cf32() bool {
	return _this.Get_().(D10_DC24).Cf32
}

func (_this D10) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D10_DC24).Cf33
}

func (_this D10) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D10_DC23).Cf30
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + data.Cf33.VerbatimString(true) + ")"
		}
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32 && data1.Cf33.Equals(data2.Cf33)
		}
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D11_DC25 struct {
	Cf34 _dafny.MultiSet
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf34 _dafny.MultiSet) D11 {
	return D11{D11_DC25{Cf34}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D11) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D11_DC25).Cf34
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf34.Equals(data2.Cf34)
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
	Cf36 _dafny.Int
	Cf37 _dafny.Int
	Cf38 _dafny.Int
	Cf39 _dafny.Int
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf36 _dafny.Int, Cf37 _dafny.Int, Cf38 _dafny.Int, Cf39 _dafny.Int) D12 {
	return D12{D12_DC27{Cf36, Cf37, Cf38, Cf39}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

type D12_DC26 struct {
	Cf35 _dafny.Map
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf35 _dafny.Map) D12 {
	return D12{D12_DC26{Cf35}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

type D12_DC28 struct {
	Cf40 D12
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf40 D12) D12 {
	return D12{D12_DC28{Cf40}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC27_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D12) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf36
}

func (_this D12) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf37
}

func (_this D12) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf38
}

func (_this D12) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf39
}

func (_this D12) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D12_DC26).Cf35
}

func (_this D12) Dtor_cf40() D12 {
	return _this.Get_().(D12_DC28).Cf40
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf40) + ")"
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
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
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

type D13_DC30 struct {
	Cf42 bool
	Cf43 bool
	Cf44 _dafny.Set
	Cf45 *C0
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf42 bool, Cf43 bool, Cf44 _dafny.Set, Cf45 *C0) D13 {
	return D13{D13_DC30{Cf42, Cf43, Cf44, Cf45}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

type D13_DC29 struct {
	Cf41 _dafny.Set
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf41 _dafny.Set) D13 {
	return D13{D13_DC29{Cf41}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

type D13_DC31 struct {
	Cf46 D13
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf46 D13) D13 {
	return D13{D13_DC31{Cf46}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC30_(false, false, _dafny.EmptySet, (*C0)(nil))
}

func (_this D13) Dtor_cf42() bool {
	return _this.Get_().(D13_DC30).Cf42
}

func (_this D13) Dtor_cf43() bool {
	return _this.Get_().(D13_DC30).Cf43
}

func (_this D13) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D13_DC30).Cf44
}

func (_this D13) Dtor_cf45() *C0 {
	return _this.Get_().(D13_DC30).Cf45
}

func (_this D13) Dtor_cf41() _dafny.Set {
	return _this.Get_().(D13_DC29).Cf41
}

func (_this D13) Dtor_cf46() D13 {
	return _this.Get_().(D13_DC31).Cf46
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf46) + ")"
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
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43 == data2.Cf43 && data1.Cf44.Equals(data2.Cf44) && data1.Cf45 == data2.Cf45
		}
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf41.Equals(data2.Cf41)
		}
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf46.Equals(data2.Cf46)
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
	Cf47 _dafny.Sequence
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_(Cf47 _dafny.Sequence) D14 {
	return D14{D14_DC32{Cf47}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D14) Dtor_cf47() _dafny.Sequence {
	return _this.Get_().(D14_DC32).Cf47
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC32:
		{
			return "D14.DC32" + "(" + _dafny.String(data.Cf47) + ")"
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
			return ok && data1.Cf47.Equals(data2.Cf47)
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

type D15_DC33 struct {
	Cf48 _dafny.MultiSet
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf48 _dafny.MultiSet) D15 {
	return D15{D15_DC33{Cf48}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D15) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D15_DC33).Cf48
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

type D16_DC35 struct {
	Cf50 _dafny.Int
	Cf51 _dafny.CodePoint
	Cf52 _dafny.Int
	Cf53 _dafny.Int
	Cf54 _dafny.Int
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf50 _dafny.Int, Cf51 _dafny.CodePoint, Cf52 _dafny.Int, Cf53 _dafny.Int, Cf54 _dafny.Int) D16 {
	return D16{D16_DC35{Cf50, Cf51, Cf52, Cf53, Cf54}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

type D16_DC34 struct {
	Cf49 *C2
}

func (D16_DC34) isD16() {}

func (CompanionStruct_D16_) Create_DC34_(Cf49 *C2) D16 {
	return D16{D16_DC34{Cf49}}
}

func (_this D16) Is_DC34() bool {
	_, ok := _this.Get_().(D16_DC34)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC35_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf50
}

func (_this D16) Dtor_cf51() _dafny.CodePoint {
	return _this.Get_().(D16_DC35).Cf51
}

func (_this D16) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf52
}

func (_this D16) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf53
}

func (_this D16) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf54
}

func (_this D16) Dtor_cf49() *C2 {
	return _this.Get_().(D16_DC34).Cf49
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D16_DC34:
		{
			return "D16.DC34" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf50.Cmp(data2.Cf50) == 0 && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D16_DC34:
		{
			data2, ok := other.Get_().(D16_DC34)
			return ok && data1.Cf49 == data2.Cf49
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

type D17_DC36 struct {
	Cf55 _dafny.Array
}

func (D17_DC36) isD17() {}

func (CompanionStruct_D17_) Create_DC36_(Cf55 _dafny.Array) D17 {
	return D17{D17_DC36{Cf55}}
}

func (_this D17) Is_DC36() bool {
	_, ok := _this.Get_().(D17_DC36)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D17) Dtor_cf55() _dafny.Array {
	return _this.Get_().(D17_DC36).Cf55
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC36:
		{
			return "D17.DC36" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC36:
		{
			data2, ok := other.Get_().(D17_DC36)
			return ok && data1.Cf55 == data2.Cf55
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

type D18_DC37 struct {
	Cf56 *C3
}

func (D18_DC37) isD18() {}

func (CompanionStruct_D18_) Create_DC37_(Cf56 *C3) D18 {
	return D18{D18_DC37{Cf56}}
}

func (_this D18) Is_DC37() bool {
	_, ok := _this.Get_().(D18_DC37)
	return ok
}

func (CompanionStruct_D18_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D18) Dtor_cf56() *C3 {
	return _this.Get_().(D18_DC37).Cf56
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC37:
		{
			return "D18.DC37" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC37:
		{
			data2, ok := other.Get_().(D18_DC37)
			return ok && data1.Cf56 == data2.Cf56
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

type D19_DC39 struct {
	Cf58 bool
	Cf59 _dafny.Int
}

func (D19_DC39) isD19() {}

func (CompanionStruct_D19_) Create_DC39_(Cf58 bool, Cf59 _dafny.Int) D19 {
	return D19{D19_DC39{Cf58, Cf59}}
}

func (_this D19) Is_DC39() bool {
	_, ok := _this.Get_().(D19_DC39)
	return ok
}

type D19_DC38 struct {
	Cf57 _dafny.Map
}

func (D19_DC38) isD19() {}

func (CompanionStruct_D19_) Create_DC38_(Cf57 _dafny.Map) D19 {
	return D19{D19_DC38{Cf57}}
}

func (_this D19) Is_DC38() bool {
	_, ok := _this.Get_().(D19_DC38)
	return ok
}

type D19_DC40 struct {
	Cf60 D19
}

func (D19_DC40) isD19() {}

func (CompanionStruct_D19_) Create_DC40_(Cf60 D19) D19 {
	return D19{D19_DC40{Cf60}}
}

func (_this D19) Is_DC40() bool {
	_, ok := _this.Get_().(D19_DC40)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC39_(false, _dafny.Zero)
}

func (_this D19) Dtor_cf58() bool {
	return _this.Get_().(D19_DC39).Cf58
}

func (_this D19) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D19_DC39).Cf59
}

func (_this D19) Dtor_cf57() _dafny.Map {
	return _this.Get_().(D19_DC38).Cf57
}

func (_this D19) Dtor_cf60() D19 {
	return _this.Get_().(D19_DC40).Cf60
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC39:
		{
			return "D19.DC39" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D19_DC38:
		{
			return "D19.DC38" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D19_DC40:
		{
			return "D19.DC40" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC39:
		{
			data2, ok := other.Get_().(D19_DC39)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D19_DC38:
		{
			data2, ok := other.Get_().(D19_DC38)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	case D19_DC40:
		{
			data2, ok := other.Get_().(D19_DC40)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

type D20_DC41 struct {
	Cf61 _dafny.Sequence
}

func (D20_DC41) isD20() {}

func (CompanionStruct_D20_) Create_DC41_(Cf61 _dafny.Sequence) D20 {
	return D20{D20_DC41{Cf61}}
}

func (_this D20) Is_DC41() bool {
	_, ok := _this.Get_().(D20_DC41)
	return ok
}

func (CompanionStruct_D20_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D20) Dtor_cf61() _dafny.Sequence {
	return _this.Get_().(D20_DC41).Cf61
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC41:
		{
			return "D20.DC41" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC41:
		{
			data2, ok := other.Get_().(D20_DC41)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D21_DC43 struct {
	Cf63 _dafny.Sequence
	Cf64 bool
}

func (D21_DC43) isD21() {}

func (CompanionStruct_D21_) Create_DC43_(Cf63 _dafny.Sequence, Cf64 bool) D21 {
	return D21{D21_DC43{Cf63, Cf64}}
}

func (_this D21) Is_DC43() bool {
	_, ok := _this.Get_().(D21_DC43)
	return ok
}

type D21_DC42 struct {
	Cf62 _dafny.Map
}

func (D21_DC42) isD21() {}

func (CompanionStruct_D21_) Create_DC42_(Cf62 _dafny.Map) D21 {
	return D21{D21_DC42{Cf62}}
}

func (_this D21) Is_DC42() bool {
	_, ok := _this.Get_().(D21_DC42)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC43_(_dafny.EmptySeq, false)
}

func (_this D21) Dtor_cf63() _dafny.Sequence {
	return _this.Get_().(D21_DC43).Cf63
}

func (_this D21) Dtor_cf64() bool {
	return _this.Get_().(D21_DC43).Cf64
}

func (_this D21) Dtor_cf62() _dafny.Map {
	return _this.Get_().(D21_DC42).Cf62
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC43:
		{
			return "D21.DC43" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D21_DC42:
		{
			return "D21.DC42" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC43:
		{
			data2, ok := other.Get_().(D21_DC43)
			return ok && data1.Cf63.Equals(data2.Cf63) && data1.Cf64 == data2.Cf64
		}
	case D21_DC42:
		{
			data2, ok := other.Get_().(D21_DC42)
			return ok && data1.Cf62.Equals(data2.Cf62)
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
	Cf65 _dafny.MultiSet
}

func (D22_DC44) isD22() {}

func (CompanionStruct_D22_) Create_DC44_(Cf65 _dafny.MultiSet) D22 {
	return D22{D22_DC44{Cf65}}
}

func (_this D22) Is_DC44() bool {
	_, ok := _this.Get_().(D22_DC44)
	return ok
}

func (CompanionStruct_D22_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D22) Dtor_cf65() _dafny.MultiSet {
	return _this.Get_().(D22_DC44).Cf65
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC44:
		{
			return "D22.DC44" + "(" + _dafny.String(data.Cf65) + ")"
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

type D23_DC46 struct {
}

func (D23_DC46) isD23() {}

func (CompanionStruct_D23_) Create_DC46_() D23 {
	return D23{D23_DC46{}}
}

func (_this D23) Is_DC46() bool {
	_, ok := _this.Get_().(D23_DC46)
	return ok
}

type D23_DC47 struct {
	Cf67 _dafny.Int
	Cf68 bool
	Cf69 bool
	Cf70 _dafny.CodePoint
}

func (D23_DC47) isD23() {}

func (CompanionStruct_D23_) Create_DC47_(Cf67 _dafny.Int, Cf68 bool, Cf69 bool, Cf70 _dafny.CodePoint) D23 {
	return D23{D23_DC47{Cf67, Cf68, Cf69, Cf70}}
}

func (_this D23) Is_DC47() bool {
	_, ok := _this.Get_().(D23_DC47)
	return ok
}

type D23_DC45 struct {
	Cf66 _dafny.Array
}

func (D23_DC45) isD23() {}

func (CompanionStruct_D23_) Create_DC45_(Cf66 _dafny.Array) D23 {
	return D23{D23_DC45{Cf66}}
}

func (_this D23) Is_DC45() bool {
	_, ok := _this.Get_().(D23_DC45)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC46_()
}

func (_this D23) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D23_DC47).Cf67
}

func (_this D23) Dtor_cf68() bool {
	return _this.Get_().(D23_DC47).Cf68
}

func (_this D23) Dtor_cf69() bool {
	return _this.Get_().(D23_DC47).Cf69
}

func (_this D23) Dtor_cf70() _dafny.CodePoint {
	return _this.Get_().(D23_DC47).Cf70
}

func (_this D23) Dtor_cf66() _dafny.Array {
	return _this.Get_().(D23_DC45).Cf66
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC46:
		{
			return "D23.DC46"
		}
	case D23_DC47:
		{
			return "D23.DC47" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D23_DC45:
		{
			return "D23.DC45" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC46:
		{
			_, ok := other.Get_().(D23_DC46)
			return ok
		}
	case D23_DC47:
		{
			data2, ok := other.Get_().(D23_DC47)
			return ok && data1.Cf67.Cmp(data2.Cf67) == 0 && data1.Cf68 == data2.Cf68 && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70
		}
	case D23_DC45:
		{
			data2, ok := other.Get_().(D23_DC45)
			return ok && data1.Cf66 == data2.Cf66
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
}

func (D24_DC49) isD24() {}

func (CompanionStruct_D24_) Create_DC49_() D24 {
	return D24{D24_DC49{}}
}

func (_this D24) Is_DC49() bool {
	_, ok := _this.Get_().(D24_DC49)
	return ok
}

type D24_DC50 struct {
}

func (D24_DC50) isD24() {}

func (CompanionStruct_D24_) Create_DC50_() D24 {
	return D24{D24_DC50{}}
}

func (_this D24) Is_DC50() bool {
	_, ok := _this.Get_().(D24_DC50)
	return ok
}

type D24_DC48 struct {
	Cf71 _dafny.Set
}

func (D24_DC48) isD24() {}

func (CompanionStruct_D24_) Create_DC48_(Cf71 _dafny.Set) D24 {
	return D24{D24_DC48{Cf71}}
}

func (_this D24) Is_DC48() bool {
	_, ok := _this.Get_().(D24_DC48)
	return ok
}

type D24_DC51 struct {
	Cf72 D24
}

func (D24_DC51) isD24() {}

func (CompanionStruct_D24_) Create_DC51_(Cf72 D24) D24 {
	return D24{D24_DC51{Cf72}}
}

func (_this D24) Is_DC51() bool {
	_, ok := _this.Get_().(D24_DC51)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC49_()
}

func (_this D24) Dtor_cf71() _dafny.Set {
	return _this.Get_().(D24_DC48).Cf71
}

func (_this D24) Dtor_cf72() D24 {
	return _this.Get_().(D24_DC51).Cf72
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC49:
		{
			return "D24.DC49"
		}
	case D24_DC50:
		{
			return "D24.DC50"
		}
	case D24_DC48:
		{
			return "D24.DC48" + "(" + _dafny.String(data.Cf71) + ")"
		}
	case D24_DC51:
		{
			return "D24.DC51" + "(" + _dafny.String(data.Cf72) + ")"
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
			_, ok := other.Get_().(D24_DC49)
			return ok
		}
	case D24_DC50:
		{
			_, ok := other.Get_().(D24_DC50)
			return ok
		}
	case D24_DC48:
		{
			data2, ok := other.Get_().(D24_DC48)
			return ok && data1.Cf71.Equals(data2.Cf71)
		}
	case D24_DC51:
		{
			data2, ok := other.Get_().(D24_DC51)
			return ok && data1.Cf72.Equals(data2.Cf72)
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

// Definition of datatype D25
type D25 struct {
	Data_D25_
}

func (_this D25) Get_() Data_D25_ {
	return _this.Data_D25_
}

type Data_D25_ interface {
	isD25()
}

type CompanionStruct_D25_ struct {
}

var Companion_D25_ = CompanionStruct_D25_{}

type D25_DC53 struct {
	Cf74 _dafny.Map
	Cf75 _dafny.Int
}

func (D25_DC53) isD25() {}

func (CompanionStruct_D25_) Create_DC53_(Cf74 _dafny.Map, Cf75 _dafny.Int) D25 {
	return D25{D25_DC53{Cf74, Cf75}}
}

func (_this D25) Is_DC53() bool {
	_, ok := _this.Get_().(D25_DC53)
	return ok
}

type D25_DC52 struct {
	Cf73 _dafny.Array
}

func (D25_DC52) isD25() {}

func (CompanionStruct_D25_) Create_DC52_(Cf73 _dafny.Array) D25 {
	return D25{D25_DC52{Cf73}}
}

func (_this D25) Is_DC52() bool {
	_, ok := _this.Get_().(D25_DC52)
	return ok
}

type D25_DC54 struct {
	Cf76 D25
}

func (D25_DC54) isD25() {}

func (CompanionStruct_D25_) Create_DC54_(Cf76 D25) D25 {
	return D25{D25_DC54{Cf76}}
}

func (_this D25) Is_DC54() bool {
	_, ok := _this.Get_().(D25_DC54)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC53_(_dafny.EmptyMap, _dafny.Zero)
}

func (_this D25) Dtor_cf74() _dafny.Map {
	return _this.Get_().(D25_DC53).Cf74
}

func (_this D25) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D25_DC53).Cf75
}

func (_this D25) Dtor_cf73() _dafny.Array {
	return _this.Get_().(D25_DC52).Cf73
}

func (_this D25) Dtor_cf76() D25 {
	return _this.Get_().(D25_DC54).Cf76
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC53:
		{
			return "D25.DC53" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ")"
		}
	case D25_DC52:
		{
			return "D25.DC52" + "(" + _dafny.String(data.Cf73) + ")"
		}
	case D25_DC54:
		{
			return "D25.DC54" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC53:
		{
			data2, ok := other.Get_().(D25_DC53)
			return ok && data1.Cf74.Equals(data2.Cf74) && data1.Cf75.Cmp(data2.Cf75) == 0
		}
	case D25_DC52:
		{
			data2, ok := other.Get_().(D25_DC52)
			return ok && data1.Cf73 == data2.Cf73
		}
	case D25_DC54:
		{
			data2, ok := other.Get_().(D25_DC54)
			return ok && data1.Cf76.Equals(data2.Cf76)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D25) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D25)
	return ok && _this.Equals(typed)
}

func Type_D25_() _dafny.TypeDescriptor {
	return type_D25_{}
}

type type_D25_ struct {
}

func (_this type_D25_) Default() interface{} {
	return Companion_D25_.Default()
}

func (_this type_D25_) String() string {
	return "main.D25"
}
func (_this D25) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D25{}

// End of datatype D25

// Definition of datatype D26
type D26 struct {
	Data_D26_
}

func (_this D26) Get_() Data_D26_ {
	return _this.Data_D26_
}

type Data_D26_ interface {
	isD26()
}

type CompanionStruct_D26_ struct {
}

var Companion_D26_ = CompanionStruct_D26_{}

type D26_DC56 struct {
	Cf78 _dafny.Int
	Cf79 _dafny.Int
}

func (D26_DC56) isD26() {}

func (CompanionStruct_D26_) Create_DC56_(Cf78 _dafny.Int, Cf79 _dafny.Int) D26 {
	return D26{D26_DC56{Cf78, Cf79}}
}

func (_this D26) Is_DC56() bool {
	_, ok := _this.Get_().(D26_DC56)
	return ok
}

type D26_DC57 struct {
	Cf80 _dafny.Int
	Cf81 *C9
}

func (D26_DC57) isD26() {}

func (CompanionStruct_D26_) Create_DC57_(Cf80 _dafny.Int, Cf81 *C9) D26 {
	return D26{D26_DC57{Cf80, Cf81}}
}

func (_this D26) Is_DC57() bool {
	_, ok := _this.Get_().(D26_DC57)
	return ok
}

type D26_DC55 struct {
	Cf77 T0
}

func (D26_DC55) isD26() {}

func (CompanionStruct_D26_) Create_DC55_(Cf77 T0) D26 {
	return D26{D26_DC55{Cf77}}
}

func (_this D26) Is_DC55() bool {
	_, ok := _this.Get_().(D26_DC55)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC56_(_dafny.Zero, _dafny.Zero)
}

func (_this D26) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D26_DC56).Cf78
}

func (_this D26) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D26_DC56).Cf79
}

func (_this D26) Dtor_cf80() _dafny.Int {
	return _this.Get_().(D26_DC57).Cf80
}

func (_this D26) Dtor_cf81() *C9 {
	return _this.Get_().(D26_DC57).Cf81
}

func (_this D26) Dtor_cf77() T0 {
	return _this.Get_().(D26_DC55).Cf77
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC56:
		{
			return "D26.DC56" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D26_DC57:
		{
			return "D26.DC57" + "(" + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D26_DC55:
		{
			return "D26.DC55" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC56:
		{
			data2, ok := other.Get_().(D26_DC56)
			return ok && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D26_DC57:
		{
			data2, ok := other.Get_().(D26_DC57)
			return ok && data1.Cf80.Cmp(data2.Cf80) == 0 && data1.Cf81 == data2.Cf81
		}
	case D26_DC55:
		{
			data2, ok := other.Get_().(D26_DC55)
			return ok && _dafny.AreEqual(data1.Cf77, data2.Cf77)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D26) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D26)
	return ok && _this.Equals(typed)
}

func Type_D26_() _dafny.TypeDescriptor {
	return type_D26_{}
}

type type_D26_ struct {
}

func (_this type_D26_) Default() interface{} {
	return Companion_D26_.Default()
}

func (_this type_D26_) String() string {
	return "main.D26"
}
func (_this D26) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D26{}

// End of datatype D26

// Definition of datatype D27
type D27 struct {
	Data_D27_
}

func (_this D27) Get_() Data_D27_ {
	return _this.Data_D27_
}

type Data_D27_ interface {
	isD27()
}

type CompanionStruct_D27_ struct {
}

var Companion_D27_ = CompanionStruct_D27_{}

type D27_DC59 struct {
	Cf83 bool
	Cf84 _dafny.Array
	Cf85 bool
}

func (D27_DC59) isD27() {}

func (CompanionStruct_D27_) Create_DC59_(Cf83 bool, Cf84 _dafny.Array, Cf85 bool) D27 {
	return D27{D27_DC59{Cf83, Cf84, Cf85}}
}

func (_this D27) Is_DC59() bool {
	_, ok := _this.Get_().(D27_DC59)
	return ok
}

type D27_DC58 struct {
	Cf82 _dafny.Map
}

func (D27_DC58) isD27() {}

func (CompanionStruct_D27_) Create_DC58_(Cf82 _dafny.Map) D27 {
	return D27{D27_DC58{Cf82}}
}

func (_this D27) Is_DC58() bool {
	_, ok := _this.Get_().(D27_DC58)
	return ok
}

func (CompanionStruct_D27_) Default() D27 {
	return Companion_D27_.Create_DC59_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D27) Dtor_cf83() bool {
	return _this.Get_().(D27_DC59).Cf83
}

func (_this D27) Dtor_cf84() _dafny.Array {
	return _this.Get_().(D27_DC59).Cf84
}

func (_this D27) Dtor_cf85() bool {
	return _this.Get_().(D27_DC59).Cf85
}

func (_this D27) Dtor_cf82() _dafny.Map {
	return _this.Get_().(D27_DC58).Cf82
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC59:
		{
			return "D27.DC59" + "(" + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D27_DC58:
		{
			return "D27.DC58" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC59:
		{
			data2, ok := other.Get_().(D27_DC59)
			return ok && data1.Cf83 == data2.Cf83 && data1.Cf84 == data2.Cf84 && data1.Cf85 == data2.Cf85
		}
	case D27_DC58:
		{
			data2, ok := other.Get_().(D27_DC58)
			return ok && data1.Cf82.Equals(data2.Cf82)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D27) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D27)
	return ok && _this.Equals(typed)
}

func Type_D27_() _dafny.TypeDescriptor {
	return type_D27_{}
}

type type_D27_ struct {
}

func (_this type_D27_) Default() interface{} {
	return Companion_D27_.Default()
}

func (_this type_D27_) String() string {
	return "main.D27"
}
func (_this D27) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D27{}

// End of datatype D27

// Definition of datatype D28
type D28 struct {
	Data_D28_
}

func (_this D28) Get_() Data_D28_ {
	return _this.Data_D28_
}

type Data_D28_ interface {
	isD28()
}

type CompanionStruct_D28_ struct {
}

var Companion_D28_ = CompanionStruct_D28_{}

type D28_DC60 struct {
	Cf86 _dafny.Map
}

func (D28_DC60) isD28() {}

func (CompanionStruct_D28_) Create_DC60_(Cf86 _dafny.Map) D28 {
	return D28{D28_DC60{Cf86}}
}

func (_this D28) Is_DC60() bool {
	_, ok := _this.Get_().(D28_DC60)
	return ok
}

func (CompanionStruct_D28_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D28) Dtor_cf86() _dafny.Map {
	return _this.Get_().(D28_DC60).Cf86
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC60:
		{
			return "D28.DC60" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC60:
		{
			data2, ok := other.Get_().(D28_DC60)
			return ok && data1.Cf86.Equals(data2.Cf86)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D28) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D28)
	return ok && _this.Equals(typed)
}

func Type_D28_() _dafny.TypeDescriptor {
	return type_D28_{}
}

type type_D28_ struct {
}

func (_this type_D28_) Default() interface{} {
	return Companion_D28_.Default()
}

func (_this type_D28_) String() string {
	return "main.D28"
}
func (_this D28) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D28{}

// End of datatype D28

// Definition of trait T0
type T0 interface {
	String() string
	Fm2(globalState *GlobalState) _dafny.Map
	Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool
	M0(globalState *GlobalState) (bool, bool)
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
	F0  _dafny.Sequence
	F2  _dafny.MultiSet
	F3  _dafny.Int
	F4  bool
	F8  _dafny.Int
	_f1 _dafny.Sequence
	_f5 bool
	_f6 bool
	_f7 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptySeq
	_this.F2 = _dafny.EmptyMultiSet
	_this.F3 = _dafny.Zero
	_this.F4 = false
	_this.F8 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f5 = false
	_this._f6 = false
	_this._f7 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 _dafny.Sequence, f2 _dafny.MultiSet, f3 _dafny.Int, f4 bool, f5 bool, f6 bool, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
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
func (_this *C0) Fm19(globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(-686)).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(568))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_333_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()))).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(529), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality())
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return func() _dafny.Map {
			var _coll56 = _dafny.NewMapBuilder()
			_ = _coll56
			for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-491), _dafny.IntOfInt64(-220))); ; {
				_compr_56, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _334_v0 _dafny.Int
				_334_v0 = interface{}(_compr_56).(_dafny.Int)
				if ((_dafny.IntOfInt64(-491)).Cmp(_334_v0) <= 0) && ((_334_v0).Cmp(_dafny.IntOfInt64(-220)) < 0) {
					_coll56.Add((_334_v0).Times(_dafny.IntOfInt64(697)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(!(true), true))).Cardinality()))
				}
			}
			return _coll56.ToMap()
		}()
	}
}
func (_this *C1) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return (!(true) || (!(true))) && (_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true, !(false)), _dafny.SeqOf(true)))
	}
}
func (_this *C1) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		if false {
			var _335_v0 _dafny.Int
			_ = _335_v0
			_335_v0 = _dafny.IntOfInt64(847)
			var _336_v1 _dafny.Map
			_ = _336_v1
			_336_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v0, _335_v0)
			var _337_v2 D3
			_ = _337_v2
			_337_v2 = Companion_D3_.Create_DC8_(_336_v1)
			var _338_v3 bool
			_ = _338_v3
			_338_v3 = true
			var _pat_let_tv10 = _336_v1
			_ = _pat_let_tv10
			var _339_v4 _dafny.Array
			_ = _339_v4
			var _nwElement0_13 D3 = _337_v2
			_ = _nwElement0_13
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(13))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_13, 0)
			(_nw62).ArraySet1((func() D3 {
				if !(_338_v3) {
					return _337_v2
				}
				return _337_v2
			})(), 1)
			(_nw62).ArraySet1(_337_v2, 2)
			(_nw62).ArraySet1(Companion_D3_.Create_DC8_(_336_v1), 3)
			(_nw62).ArraySet1(_337_v2, 4)
			(_nw62).ArraySet1(_337_v2, 5)
			(_nw62).ArraySet1(_337_v2, 6)
			(_nw62).ArraySet1(_337_v2, 7)
			(_nw62).ArraySet1(Companion_Default___.Fm24(globalState), 8)
			(_nw62).ArraySet1(_337_v2, 9)
			(_nw62).ArraySet1((func() D3 {
				if false {
					return func(_pat_let9_0 D3) D3 {
						return func(_340_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let10_0 _dafny.Map) D3 {
								return func(_341_dt__update_hcf12_h0 _dafny.Map) D3 {
									return Companion_D3_.Create_DC8_(_341_dt__update_hcf12_h0)
								}(_pat_let10_0)
							}(_pat_let_tv10)
						}(_pat_let9_0)
					}(_337_v2)
				}
				return _337_v2
			})(), 10)
			(_nw62).ArraySet1(_337_v2, 11)
			(_nw62).ArraySet1(_337_v2, 12)
			_339_v4 = _nw62
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_339_v4), 0))
			_ = _index55
			(_339_v4).ArraySet1(_337_v2, (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_339_v4), 0))
			_ = _index56
			(_339_v4).ArraySet1(_337_v2, (_index56).Int())
			var _342_v5 *C0
			_ = _342_v5
			var _nw63 *C0 = New_C0_()
			_ = _nw63
			_nw63.Ctor__()
			_342_v5 = _nw63
			var _343_v6 _dafny.Set
			_ = _343_v6
			_343_v6 = _dafny.SetOf(_335_v0, _335_v0, Companion_Default___.SafeDivisionInt(_335_v0, _335_v0))
			var _344_v7 _dafny.Map
			_ = _344_v7
			_344_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_338_v3, _343_v6)
			_343_v6 = (func() _dafny.Set {
				if _338_v3 {
					return (func() _dafny.Set {
						if (_344_v7).Contains(_338_v3) {
							return (_344_v7).Get(_338_v3).(_dafny.Set)
						}
						return _343_v6
					})()
				}
				return (_343_v6).Difference(_343_v6)
			})()
			var _345_v8 D3
			_ = _345_v8
			_345_v8 = Companion_D3_.Create_DC9_()
			var _source10 D3 = _345_v8
			_ = _source10
			if _source10.Is_DC9() {
				var _346_v9 _dafny.Set
				_ = _346_v9
				_346_v9 = _dafny.SetOf(_338_v3)
				var _347_v10 _dafny.Array
				_ = _347_v10
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_8
				var _nw64 _dafny.Array
				_ = _nw64
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw64 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.MultiSet = (func(_348_v0 _dafny.Int, _349_v3 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_350_i0 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_348_v0, _348_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_349_v3)).Cardinality())), _348_v0))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_348_v0)))
						}
					})(_335_v0, _338_v3)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw64 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw64).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw64).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_347_v10 = _nw64
				var _351_v11 _dafny.Sequence
				_ = _351_v11
				_351_v11 = _dafny.SeqOf(_335_v0, _335_v0, _335_v0, _335_v0, _335_v0)
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_347_v10), 0))
				_ = _index57
				(_347_v10).ArraySet1(_dafny.MultiSetFromSeq(_351_v11), (_index57).Int())
				var _352_v12 _dafny.Sequence
				_ = _352_v12
				_352_v12 = _dafny.SeqOf(_338_v3)
				var _353_v13 _dafny.MultiSet
				_ = _353_v13
				_353_v13 = _dafny.MultiSetOf(_335_v0, _335_v0, (_346_v9).Cardinality(), _335_v0)
				var _354_v14 D4
				_ = _354_v14
				_354_v14 = Companion_D4_.Create_DC10_(_352_v12)
				var _355_v15 _dafny.MultiSet
				_ = _355_v15
				_355_v15 = _dafny.MultiSetOf(_354_v14)
				var _356_v16 _dafny.Map
				_ = _356_v16
				_356_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v15, _335_v0)
				var _357_v17 _dafny.Map
				_ = _357_v17
				_357_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v0, _356_v16)
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_347_v10), 0))
				_ = _index58
				var _rhs44 _dafny.Int = _dafny.IntOfUint32((_352_v12).Cardinality())
				_ = _rhs44
				var _rhs45 _dafny.Int = _335_v0
				_ = _rhs45
				var _rhs46 _dafny.Set = _346_v9
				_ = _rhs46
				var _rhs47 _dafny.MultiSet = ((_353_v13).Difference(_353_v13)).Intersection((_353_v13).Intersection(_353_v13))
				_ = _rhs47
				var _rhs48 _dafny.Int = Companion_Default___.SafeModuloInt(_335_v0, ((func() _dafny.Map {
					if (_357_v17).Contains((_dafny.Zero).Minus(_335_v0)) {
						return (_357_v17).Get((_dafny.Zero).Minus(_335_v0)).(_dafny.Map)
					}
					return _356_v16
				})()).Cardinality())
				_ = _rhs48
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 _dafny.Array = _347_v10
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_347_v10), 0))
				_ = _lhs38
				var _lhs39 *GlobalState = globalState
				_ = _lhs39
				_lhs35.F8 = _rhs44
				_lhs36.F8 = _rhs45
				_346_v9 = _rhs46
				(_lhs37).ArraySet1(_rhs47, (_lhs38).Int())
				_lhs39.F8 = _rhs48
				var _358_v18 _dafny.Sequence
				_ = _358_v18
				_358_v18 = _dafny.UnicodeSeqOfUtf8Bytes("uou")
				var _359_v19 _dafny.CodePoint
				_ = _359_v19
				_359_v19 = _dafny.CodePoint('r')
				var _rhs49 _dafny.Sequence = _358_v18
				_ = _rhs49
				var _rhs50 bool = !_dafny.Companion_Sequence_.Contains(_358_v18, (func() _dafny.CodePoint {
					if _338_v3 {
						return _dafny.CodePoint('j')
					}
					return _359_v19
				})())
				_ = _rhs50
				var _lhs40 *GlobalState = globalState
				_ = _lhs40
				_358_v18 = _rhs49
				_lhs40.F4 = _rhs50
				(globalState).F3 = (_342_v5).Fm19(globalState)
				var _360_v20 D2
				_ = _360_v20
				_360_v20 = Companion_D2_.Create_DC7_(_359_v19, _338_v3, _335_v0)
				var _361_v21 D4
				_ = _361_v21
				_361_v21 = Companion_D4_.Create_DC11_(_338_v3, _338_v3, (_351_v11).Select((Companion_Default___.SafeIndex(_335_v0, _dafny.IntOfUint32((_351_v11).Cardinality()))).Uint32()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("e"), (Companion_Default___.SafeIndex(_335_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()))).Uint32(), _359_v19))
				var _362_v22 _dafny.Sequence
				_ = _362_v22
				_362_v22 = _351_v11
				var _363_v23 _dafny.Map
				_ = _363_v23
				_363_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v0, _343_v6)
				var _364_v24 _dafny.Map
				_ = _364_v24
				_364_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_335_v0), _353_v13)
				var _365_v25 _dafny.Sequence
				_ = _365_v25
				_365_v25 = _dafny.SeqOf(_364_v24)
				var _366_v26 _dafny.Array
				_ = _366_v26
				var _nwElement0_14 _dafny.Int = Companion_Default___.SafeDivisionInt(_335_v0, _335_v0)
				_ = _nwElement0_14
				var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(28))
				_ = _nw65
				(_nw65).ArraySet1(_nwElement0_14, 0)
				(_nw65).ArraySet1(_335_v0, 1)
				(_nw65).ArraySet1(_335_v0, 2)
				(_nw65).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v0, _335_v0)).Cardinality(), 3)
				(_nw65).ArraySet1((_335_v0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}((func(_367_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_368_i1 _dafny.Int) _dafny.CodePoint {
						return _367_v19
					}
				})(_359_v19)))).Cardinality())), 4)
				(_nw65).ArraySet1(_335_v0, 5)
				(_nw65).ArraySet1((func() _dafny.Int {
					if _338_v3 {
						return (_360_v20).Dtor_cf11()
					}
					return (_361_v21).Dtor_cf16()
				})(), 6)
				(_nw65).ArraySet1(_335_v0, 7)
				(_nw65).ArraySet1(_335_v0, 8)
				(_nw65).ArraySet1(_335_v0, 9)
				(_nw65).ArraySet1(_335_v0, 10)
				(_nw65).ArraySet1(_335_v0, 11)
				(_nw65).ArraySet1(_dafny.IntOfInt64(-245), 12)
				(_nw65).ArraySet1((func() _dafny.Int {
					if (_336_v1).Contains((_dafny.Zero).Minus((func() _dafny.Int {
						if (_336_v1).Contains(_335_v0) {
							return (_336_v1).Get(_335_v0).(_dafny.Int)
						}
						return _335_v0
					})())) {
						return (_336_v1).Get((_dafny.Zero).Minus((func() _dafny.Int {
							if (_336_v1).Contains(_335_v0) {
								return (_336_v1).Get(_335_v0).(_dafny.Int)
							}
							return _335_v0
						})())).(_dafny.Int)
					}
					return _335_v0
				})(), 13)
				(_nw65).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
					if ((_347_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_347_v10), 0))).Int()).(_dafny.MultiSet)).Contains(_335_v0) {
						return ((_347_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_347_v10), 0))).Int()).(_dafny.MultiSet)).Multiplicity(_335_v0)
					}
					return _335_v0
				})())).Minus(_335_v0), 14)
				(_nw65).ArraySet1((_335_v0).Plus(_335_v0), 15)
				(_nw65).ArraySet1(_335_v0, 16)
				(_nw65).ArraySet1(_dafny.IntOfUint32((_362_v22).Cardinality()), 17)
				(_nw65).ArraySet1(_335_v0, 18)
				(_nw65).ArraySet1(_335_v0, 19)
				(_nw65).ArraySet1(_335_v0, 20)
				(_nw65).ArraySet1(Companion_Default___.SafeDivisionInt((_363_v23).Cardinality(), _335_v0), 21)
				(_nw65).ArraySet1((_335_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nalyv")).Cardinality())), 22)
				(_nw65).ArraySet1((func() _dafny.Int {
					if true {
						return (_342_v5).Fm19(globalState)
					}
					return _335_v0
				})(), 23)
				(_nw65).ArraySet1(_335_v0, 24)
				(_nw65).ArraySet1(_dafny.IntOfInt64(715), 25)
				(_nw65).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-239))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_369_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_370_i2 _dafny.Int) _dafny.CodePoint {
						return _369_v19
					}
				})(_359_v19)))).Cardinality()), ((_365_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.IntOfUint32((_365_v25).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), 26)
				(_nw65).ArraySet1((_dafny.Zero).Minus(_335_v0), 27)
				_366_v26 = _nw65
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_366_v26), 0))
				_ = _index59
				(_366_v26).ArraySet1(Companion_Default___.SafeDivisionInt(_335_v0, _335_v0), (_index59).Int())
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_366_v26), 0))
				_ = _index60
				(_366_v26).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(865)).Minus(_dafny.IntOfUint32((_358_v18).Cardinality()))), (_index60).Int())
			} else {
				var _371___mcc_h0 _dafny.Map = _source10.Get_().(D3_DC8).Cf12
				_ = _371___mcc_h0
				var _372_cf12 _dafny.Map = _371___mcc_h0
				_ = _372_cf12
				var _373_v27 _dafny.Array
				_ = _373_v27
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw66
				_373_v27 = _nw66
				var _374_v28 _dafny.Sequence
				_ = _374_v28
				_374_v28 = _dafny.UnicodeSeqOfUtf8Bytes("pqitorcr")
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_373_v27), 0))
				_ = _index61
				(_373_v27).ArraySet1((_dafny.IntOfUint32((_374_v28).Cardinality())).Plus(_335_v0), (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_373_v27), 0))
				_ = _index62
				(_373_v27).ArraySet1(_335_v0, (_index62).Int())
				var _375_v29 _dafny.Set
				_ = _375_v29
				_375_v29 = _dafny.SetOf(_338_v3, _338_v3, _338_v3)
				_338_v3 = (_375_v29).IsSubsetOf(_dafny.SetOf(!(_338_v3)))
				(globalState).F8 = (_373_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_373_v27), 0))).Int()).(_dafny.Int)
				var _376_v30 _dafny.Array
				_ = _376_v30
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_9
				var _nw67 _dafny.Array
				_ = _nw67
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw67 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Map = (func(_377_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_378_i3 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _377_v0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfInt64(-308)))
						}
					})(_335_v0)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw67 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw67).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw67).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_376_v30 = _nw67
				_376_v30 = _376_v30
			}
			var _379_v31 _dafny.Sequence
			_ = _379_v31
			_379_v31 = _dafny.SeqOf(false, _338_v3)
			if (_379_v31).Select((Companion_Default___.SafeIndex(_335_v0, _dafny.IntOfUint32((_379_v31).Cardinality()))).Uint32()).(bool) {
				var _380_v32 _dafny.Sequence
				_ = _380_v32
				_380_v32 = _dafny.UnicodeSeqOfUtf8Bytes("prrqlk")
				var _381_v33 _dafny.Array
				_ = _381_v33
				var _nw68 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw68
				_381_v33 = _nw68
				var _382_v34 _dafny.Map
				_ = _382_v34
				_382_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_338_v3, _381_v33)
				var _383_v35 _dafny.CodePoint
				_ = _383_v35
				_383_v35 = _dafny.CodePoint('v')
				(globalState).F8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(_335_v0, _dafny.CodePoint('e'), _338_v3, _336_v1, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-7), _dafny.IntOfUint32((Companion_Default___.Fm1(_335_v0, _dafny.CodePoint('e'), _338_v3, _336_v1, globalState)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_380_v32).Cardinality())), Companion_Default___.Fm1((_382_v34).Cardinality(), _383_v35, false, _336_v1, globalState))).Cardinality())
				var _384_v36 _dafny.Map
				_ = _384_v36
				_384_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_338_v3, _338_v3)
				var _385_v37 _dafny.Map
				_ = _385_v37
				_385_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v36, _338_v3)
				var _386_v38 _dafny.Array
				_ = _386_v38
				var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
				_ = _nw69
				_386_v38 = _nw69
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_386_v38), 0))
				_ = _index63
				(_386_v38).ArraySet1CodePoint(_383_v35, (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_386_v38), 0))
				_ = _index64
				var _rhs51 _dafny.Sequence = Companion_Default___.Fm25(globalState)
				_ = _rhs51
				var _rhs52 _dafny.Int = (func() _dafny.Int {
					if _338_v3 {
						return (_335_v0).Times((func() _dafny.Int {
							if (_336_v1).Contains(_335_v0) {
								return (_336_v1).Get(_335_v0).(_dafny.Int)
							}
							return _335_v0
						})())
					}
					return _335_v0
				})()
				_ = _rhs52
				var _rhs53 _dafny.Map = (_385_v37).Merge(_385_v37)
				_ = _rhs53
				var _rhs54 _dafny.CodePoint = _383_v35
				_ = _rhs54
				var _rhs55 _dafny.Int = Companion_Default___.SafeDivisionInt(_335_v0, (_335_v0).Plus((_dafny.Zero).Minus(_335_v0)))
				_ = _rhs55
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 _dafny.Array = _386_v38
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_386_v38), 0))
				_ = _lhs43
				var _lhs44 *GlobalState = globalState
				_ = _lhs44
				_380_v32 = _rhs51
				_lhs41.F3 = _rhs52
				_385_v37 = _rhs53
				(_lhs42).ArraySet1CodePoint(_rhs54, (_lhs43).Int())
				_lhs44.F8 = _rhs55
				var _387_v39 _dafny.MultiSet
				_ = _387_v39
				_387_v39 = _dafny.MultiSetOf(_335_v0)
				var _388_v40 _dafny.MultiSet
				_ = _388_v40
				_388_v40 = _dafny.MultiSetOf(!(_338_v3))
				var _389_v41 _dafny.Sequence
				_ = _389_v41
				_389_v41 = _dafny.SeqOf((_dafny.Zero).Minus(_335_v0), (_387_v39).Cardinality(), _335_v0, _335_v0, (func() _dafny.Int {
					if (_388_v40).Contains(_338_v3) {
						return (_388_v40).Multiplicity(_338_v3)
					}
					return (_dafny.Zero).Minus(_335_v0)
				})())
				(globalState).F8 = (_389_v41).Select((Companion_Default___.SafeIndex((_335_v0).Times(_335_v0), _dafny.IntOfUint32((_389_v41).Cardinality()))).Uint32()).(_dafny.Int)
				_338_v3 = _338_v3
				(globalState).F8 = (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v35, _dafny.Companion_Sequence_.Equal(_380_v32, _380_v32))).Cardinality())
			} else {
				var _390_v42 _dafny.Array
				_ = _390_v42
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw70
				_390_v42 = _nw70
				var _391_v43 _dafny.Array
				_ = _391_v43
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw71
				_391_v43 = _nw71
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_390_v42), 0))
				_ = _index65
				(_390_v42).ArraySet1(_391_v43, (_index65).Int())
				var _392_v44 D5
				_ = _392_v44
				_392_v44 = Companion_D5_.Create_DC14_(_335_v0, !(_338_v3))
				var _393_v45 _dafny.Sequence
				_ = _393_v45
				_393_v45 = _dafny.SeqOf(_392_v44)
				var _394_v46 _dafny.Set
				_ = _394_v46
				_394_v46 = _dafny.SetOf(_338_v3, _338_v3, _338_v3, _338_v3)
				var _395_v47 D9
				_ = _395_v47
				_395_v47 = Companion_D9_.Create_DC20_(_394_v46)
				var _396_v48 _dafny.Sequence
				_ = _396_v48
				_396_v48 = _dafny.UnicodeSeqOfUtf8Bytes("fw")
				var _397_v49 _dafny.Sequence
				_ = _397_v49
				_397_v49 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_396_v48).Cardinality())))
				var _pat_let_tv11 = _394_v46
				_ = _pat_let_tv11
				var _pat_let_tv12 = _397_v49
				_ = _pat_let_tv12
				var _398_v50 _dafny.Array
				_ = _398_v50
				var _nwElement0_15 bool = _338_v3
				_ = _nwElement0_15
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(19))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_15, 0)
				(_nw72).ArraySet1(_338_v3, 1)
				(_nw72).ArraySet1(_338_v3, 2)
				(_nw72).ArraySet1(_338_v3, 3)
				(_nw72).ArraySet1((_338_v3) || (_338_v3), 4)
				(_nw72).ArraySet1(_338_v3, 5)
				(_nw72).ArraySet1(_338_v3, 6)
				(_nw72).ArraySet1(_338_v3, 7)
				(_nw72).ArraySet1(!(_338_v3), 8)
				(_nw72).ArraySet1(!(_338_v3), 9)
				(_nw72).ArraySet1(_338_v3, 10)
				(_nw72).ArraySet1(_338_v3, 11)
				(_nw72).ArraySet1(_338_v3, 12)
				(_nw72).ArraySet1(_338_v3, 13)
				(_nw72).ArraySet1(_338_v3, 14)
				(_nw72).ArraySet1(!(((func(_pat_let11_0 D9) D9 {
					return func(_399_dt__update__tmp_h1 D9) D9 {
						return func(_pat_let12_0 _dafny.Set) D9 {
							return func(_400_dt__update_hcf25_h0 _dafny.Set) D9 {
								return Companion_D9_.Create_DC20_(_400_dt__update_hcf25_h0)
							}(_pat_let12_0)
						}(_pat_let_tv11)
					}(_pat_let11_0)
				}(_395_v47)).Dtor_cf25()).IsProperSubsetOf(_394_v46)), 15)
				(_nw72).ArraySet1(_338_v3, 16)
				(_nw72).ArraySet1((func(_pat_let13_0 D5) D5 {
					return func(_401_dt__update__tmp_h2 D5) D5 {
						return func(_pat_let14_0 _dafny.Int) D5 {
							return func(_402_dt__update_hcf20_h0 _dafny.Int) D5 {
								return Companion_D5_.Create_DC14_(_402_dt__update_hcf20_h0, (_401_dt__update__tmp_h2).Dtor_cf21())
							}(_pat_let14_0)
						}(_dafny.IntOfUint32((_pat_let_tv12).Cardinality()))
					}(_pat_let13_0)
				}(_392_v44)).Dtor_cf21(), 17)
				(_nw72).ArraySet1(_338_v3, 18)
				_398_v50 = _nw72
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))
				_ = _index66
				(_398_v50).ArraySet1(!(_338_v3), (_index66).Int())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_390_v42), 0))
				_ = _index67
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))
				_ = _index68
				var _rhs56 bool = _338_v3
				_ = _rhs56
				var _rhs57 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("hqv"), _396_v48)
				_ = _rhs57
				var _rhs58 _dafny.Array = _391_v43
				_ = _rhs58
				var _rhs59 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_393_v45, _393_v45)
				_ = _rhs59
				var _rhs60 bool = false
				_ = _rhs60
				var _lhs45 _dafny.Array = _390_v42
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_390_v42), 0))
				_ = _lhs46
				var _lhs47 _dafny.Array = _398_v50
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))
				_ = _lhs48
				_338_v3 = _rhs56
				r0 = _rhs57
				(_lhs45).ArraySet1(_rhs58, (_lhs46).Int())
				_393_v45 = _rhs59
				(_lhs47).ArraySet1(_rhs60, (_lhs48).Int())
				var _403_v51 *C0
				_ = _403_v51
				var _nw73 *C0 = New_C0_()
				_ = _nw73
				_nw73.Ctor__()
				_403_v51 = _nw73
				var _404_v52 _dafny.Array
				_ = _404_v52
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw74
				_404_v52 = _nw74
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_404_v52), 0))
				_ = _index69
				(_404_v52).ArraySet1(_335_v0, (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_404_v52), 0))
				_ = _index70
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))
				_ = _index71
				var _rhs61 _dafny.Int = (_dafny.IntOfInt64(901)).Times((_335_v0).Plus(_335_v0))
				_ = _rhs61
				var _rhs62 _dafny.Int = _335_v0
				_ = _rhs62
				var _rhs63 _dafny.Int = (_342_v5).Fm19(globalState)
				_ = _rhs63
				var _rhs64 bool = (_398_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))).Int()).(bool)
				_ = _rhs64
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 _dafny.Array = _404_v52
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_404_v52), 0))
				_ = _lhs52
				var _lhs53 _dafny.Array = _398_v50
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))
				_ = _lhs54
				_lhs49.F8 = _rhs61
				_lhs50.F3 = _rhs62
				(_lhs51).ArraySet1(_rhs63, (_lhs52).Int())
				(_lhs53).ArraySet1(_rhs64, (_lhs54).Int())
				var _405_v53 _dafny.MultiSet
				_ = _405_v53
				_405_v53 = _dafny.MultiSetOf(_338_v3)
				_405_v53 = _dafny.MultiSetOf(_338_v3, (_398_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))).Int()).(bool), !(_338_v3), (func() bool {
					if (_379_v31).Select((Companion_Default___.SafeIndex((_404_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_404_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_379_v31).Cardinality()))).Uint32()).(bool) {
						return (_398_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_398_v50), 0))).Int()).(bool)
					}
					return _338_v3
				})())
				var _rhs65 bool = _338_v3
				_ = _rhs65
				var _rhs66 bool = false
				_ = _rhs66
				var _rhs67 _dafny.Array = _404_v52
				_ = _rhs67
				var _lhs55 *GlobalState = globalState
				_ = _lhs55
				_lhs55.F4 = _rhs65
				r0 = _rhs66
				_404_v52 = _rhs67
			}
		} else {
			var _406_v54 bool
			_ = _406_v54
			_406_v54 = false
			if _406_v54 {
				var _407_v55 _dafny.Int
				_ = _407_v55
				_407_v55 = _dafny.IntOfInt64(-909)
				var _408_v56 _dafny.Map
				_ = _408_v56
				_408_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v55, _dafny.SeqOf(!(true)))
				var _409_v57 D4
				_ = _409_v57
				_409_v57 = Companion_D4_.Create_DC10_(_dafny.SeqOf(false, _406_v54))
				var _410_v58 _dafny.Sequence
				_ = _410_v58
				_410_v58 = _dafny.SeqOf(Companion_D4_.Create_DC10_((func() _dafny.Sequence {
					if (_408_v56).Contains(_407_v55) {
						return (_408_v56).Get(_407_v55).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_406_v54)
				})()), _409_v57, _409_v57, _409_v57, _409_v57)
				_410_v58 = _410_v58
				var _411_v59 _dafny.Sequence
				_ = _411_v59
				_411_v59 = _dafny.SeqOf(_407_v55)
				var _412_v60 _dafny.Map
				_ = _412_v60
				_412_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_406_v54, _406_v54)
				var _413_v61 _dafny.MultiSet
				_ = _413_v61
				_413_v61 = _dafny.MultiSetOf(_dafny.IntOfInt64(-813), _407_v55, (_412_v60).Cardinality())
				var _414_v62 _dafny.Map
				_ = _414_v62
				_414_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(_411_v59, _413_v61, _406_v54, globalState), _407_v55)
				var _415_v63 _dafny.CodePoint
				_ = _415_v63
				_415_v63 = _dafny.CodePoint('y')
				var _416_v64 _dafny.Map
				_ = _416_v64
				_416_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_407_v55).Minus((_414_v62).Cardinality()), (func() _dafny.CodePoint {
					if _406_v54 {
						return _415_v63
					}
					return _415_v63
				})())
				_416_v64 = (_416_v64).Update((_407_v55).Times(_407_v55), _415_v63)
				var _417_v65 _dafny.Array
				_ = _417_v65
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_10
				var _nw75 _dafny.Array
				_ = _nw75
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw75 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.CodePoint = (func(_418_v63 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_419_i4 _dafny.Int) _dafny.CodePoint {
							return _418_v63
						}
					})(_415_v63)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw75 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw75).ArraySet1CodePoint(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw75).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_417_v65 = _nw75
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_417_v65), 0))
				_ = _index72
				(_417_v65).ArraySet1CodePoint(_415_v63, (_index72).Int())
				var _420_v66 _dafny.Array
				_ = _420_v66
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw76
				_420_v66 = _nw76
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_417_v65), 0))
				_ = _index73
				var _rhs68 _dafny.Int = (_407_v55).Plus(_407_v55)
				_ = _rhs68
				var _rhs69 _dafny.CodePoint = _415_v63
				_ = _rhs69
				var _rhs70 _dafny.Array = _420_v66
				_ = _rhs70
				var _lhs56 *GlobalState = globalState
				_ = _lhs56
				var _lhs57 _dafny.Array = _417_v65
				_ = _lhs57
				var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_417_v65), 0))
				_ = _lhs58
				_lhs56.F3 = _rhs68
				(_lhs57).ArraySet1CodePoint(_rhs69, (_lhs58).Int())
				_420_v66 = _rhs70
				var _421_v67 *C0
				_ = _421_v67
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__()
				_421_v67 = _nw77
				(globalState).F8 = (_411_v59).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_407_v55), _dafny.IntOfUint32((_411_v59).Cardinality()))).Uint32()).(_dafny.Int)
			} else {
				var _422_v68 _dafny.Set
				_ = _422_v68
				_422_v68 = _dafny.SetOf(true)
				var _423_v69 _dafny.Sequence
				_ = _423_v69
				_423_v69 = _dafny.SeqOf(_406_v54, (_dafny.SetOf(true, _406_v54)).IsDisjointFrom(_422_v68))
				var _424_v70 _dafny.Int
				_ = _424_v70
				_424_v70 = _dafny.IntOfInt64(70)
				(globalState).F4 = (_423_v69).Select((Companion_Default___.SafeIndex(_424_v70, _dafny.IntOfUint32((_423_v69).Cardinality()))).Uint32()).(bool)
				r0 = _406_v54
				var _425_v71 _dafny.Array
				_ = _425_v71
				var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
				_ = _nw78
				_425_v71 = _nw78
				var _426_v72 _dafny.CodePoint
				_ = _426_v72
				_426_v72 = _dafny.CodePoint('m')
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_425_v71), 0))
				_ = _index74
				(_425_v71).ArraySet1CodePoint(_426_v72, (_index74).Int())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_425_v71), 0))
				_ = _index75
				(_425_v71).ArraySet1CodePoint(_426_v72, (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_425_v71), 0))
				_ = _index76
				(_425_v71).ArraySet1CodePoint((_425_v71).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_425_v71), 0))).Int()), (_index76).Int())
				var _427_v73 _dafny.Sequence
				_ = _427_v73
				_427_v73 = _dafny.UnicodeSeqOfUtf8Bytes("jcrxtqw")
				_427_v73 = _427_v73
			}
			var _428_v74 _dafny.Sequence
			_ = _428_v74
			_428_v74 = _dafny.UnicodeSeqOfUtf8Bytes("ilxgf")
			var _429_v75 _dafny.Int
			_ = _429_v75
			_429_v75 = _dafny.IntOfInt64(857)
			var _430_v76 _dafny.Map
			_ = _430_v76
			_430_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_428_v74).Cardinality()), _429_v75)
			var _431_v78 _dafny.CodePoint
			_ = _431_v78
			_431_v78 = _dafny.CodePoint('u')
			_430_v76 = (_430_v76).Update(Companion_Default___.SafeDivisionInt((func() _dafny.Map {
				var _coll57 = _dafny.NewMapBuilder()
				_ = _coll57
				for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(513), _dafny.IntOfInt64(938))); ; {
					_compr_57, _ok58 := _iter58()
					if !_ok58 {
						break
					}
					var _432_v77 _dafny.Int
					_432_v77 = interface{}(_compr_57).(_dafny.Int)
					if ((_dafny.IntOfInt64(513)).Cmp(_432_v77) <= 0) && ((_432_v77).Cmp(_dafny.IntOfInt64(938)) < 0) {
						_coll57.Add(Companion_Default___.SafeModuloInt(_432_v77, ((Companion_D10_.Create_DC23_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_428_v74)))).Dtor_cf30()).Cardinality()), _406_v54)
					}
				}
				return _coll57.ToMap()
			}()).Cardinality(), _429_v75), ((func() _dafny.Map {
				if _406_v54 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v78, _431_v78)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v78, _431_v78)
			})()).Cardinality())
			var _433_v79 _dafny.Sequence
			_ = _433_v79
			_433_v79 = _dafny.SeqOf(_406_v54, _406_v54)
			var _434_v80 _dafny.Set
			_ = _434_v80
			_434_v80 = _dafny.SetOf(_406_v54)
			var _435_v81 _dafny.Sequence
			_ = _435_v81
			_435_v81 = _dafny.SeqOf((func() _dafny.Set {
				if _406_v54 {
					return _dafny.SetOf(_406_v54)
				}
				return _434_v80
			})())
			var _rhs71 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_429_v75, _dafny.IntOfInt64(669)), (_dafny.IntOfUint32((_433_v79).Cardinality())).Times(_429_v75)))
			_ = _rhs71
			var _rhs72 _dafny.Int = _dafny.IntOfUint32((_435_v81).Cardinality())
			_ = _rhs72
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			_lhs59.F8 = _rhs71
			_lhs60.F3 = _rhs72
			var _436_v82 D3
			_ = _436_v82
			_436_v82 = Companion_D3_.Create_DC8_(_430_v76)
			var _437_v83 _dafny.MultiSet
			_ = _437_v83
			_437_v83 = _dafny.MultiSetOf(_433_v79)
			var _rhs73 _dafny.Int = _429_v75
			_ = _rhs73
			var _rhs74 _dafny.Int = Companion_Default___.Fm26(_436_v82, _429_v75, _437_v83, globalState)
			_ = _rhs74
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			_lhs61.F3 = _rhs73
			_429_v75 = _rhs74
			var _438_v84 _dafny.Sequence
			_ = _438_v84
			_438_v84 = _dafny.SeqOf(_429_v75)
			var _439_v85 D5
			_ = _439_v85
			_439_v85 = Companion_D5_.Create_DC15_()
			var _source11 D5 = (func() D5 {
				if !_dafny.Companion_Sequence_.Contains(_438_v84, _dafny.IntOfInt64(733)) {
					return _439_v85
				}
				return _439_v85
			})()
			_ = _source11
			if _source11.Is_DC14() {
				var _440___mcc_h1 _dafny.Int = _source11.Get_().(D5_DC14).Cf20
				_ = _440___mcc_h1
				var _441___mcc_h2 bool = _source11.Get_().(D5_DC14).Cf21
				_ = _441___mcc_h2
				var _442_cf21 bool = _441___mcc_h2
				_ = _442_cf21
				var _443_cf20 _dafny.Int = _440___mcc_h1
				_ = _443_cf20
				(globalState).F4 = _442_cf21
				var _444_v86 _dafny.Array
				_ = _444_v86
				var _nwElement0_16 _dafny.Int = _443_cf20
				_ = _nwElement0_16
				var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(23))
				_ = _nw79
				(_nw79).ArraySet1(_nwElement0_16, 0)
				(_nw79).ArraySet1(_dafny.IntOfInt64(340), 1)
				(_nw79).ArraySet1(_443_cf20, 2)
				(_nw79).ArraySet1(_443_cf20, 3)
				(_nw79).ArraySet1(_429_v75, 4)
				(_nw79).ArraySet1(_429_v75, 5)
				(_nw79).ArraySet1(_443_cf20, 6)
				(_nw79).ArraySet1(_429_v75, 7)
				(_nw79).ArraySet1(Companion_Default___.Fm26(_436_v82, _443_cf20, _437_v83, globalState), 8)
				(_nw79).ArraySet1((_dafny.Zero).Minus(_443_cf20), 9)
				(_nw79).ArraySet1(_443_cf20, 10)
				(_nw79).ArraySet1(_dafny.IntOfInt64(78), 11)
				(_nw79).ArraySet1(_429_v75, 12)
				(_nw79).ArraySet1(_429_v75, 13)
				(_nw79).ArraySet1(_443_cf20, 14)
				(_nw79).ArraySet1(_429_v75, 15)
				(_nw79).ArraySet1(_dafny.IntOfUint32((_428_v74).Cardinality()), 16)
				(_nw79).ArraySet1(_443_cf20, 17)
				(_nw79).ArraySet1(_429_v75, 18)
				(_nw79).ArraySet1(_443_cf20, 19)
				(_nw79).ArraySet1(_429_v75, 20)
				(_nw79).ArraySet1(_443_cf20, 21)
				(_nw79).ArraySet1(_dafny.IntOfInt64(-573), 22)
				_444_v86 = _nw79
				var _445_v87 _dafny.Map
				_ = _445_v87
				_445_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_v75, _444_v86)
				var _446_v88 _dafny.Map
				_ = _446_v88
				_446_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((func() _dafny.Array {
					if (_445_v87).Contains(_443_cf20) {
						return (_445_v87).Get(_443_cf20).(_dafny.Array)
					}
					return _444_v86
				})(), _444_v86), !(_406_v54))
				var _447_v89 _dafny.MultiSet
				_ = _447_v89
				_447_v89 = _dafny.MultiSetOf(_444_v86)
				_442_cf21 = (func() bool {
					if (_446_v88).Contains((_447_v89).Union(_447_v89)) {
						return (_446_v88).Get((_447_v89).Union(_447_v89)).(bool)
					}
					return _406_v54
				})()
				r0 = _406_v54
				(globalState).F8 = _443_cf20
			} else if _source11.Is_DC15() {
				(globalState).F3 = _429_v75
				(globalState).F3 = _429_v75
				var _448_v90 D2
				_ = _448_v90
				_448_v90 = Companion_D2_.Create_DC7_(_431_v78, _406_v54, _429_v75)
				_448_v90 = Companion_D2_.Create_DC7_(_431_v78, _406_v54, _429_v75)
				(globalState).F3 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfUint32((_428_v74).Cardinality())).Plus(_429_v75)))
			} else {
				var _449___mcc_h3 _dafny.Map = _source11.Get_().(D5_DC13).Cf19
				_ = _449___mcc_h3
				var _450_cf19 _dafny.Map = _449___mcc_h3
				_ = _450_cf19
				var _451_v91 _dafny.Map
				_ = _451_v91
				_451_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v80, Companion_Default___.Fm26(_436_v82, _dafny.IntOfUint32((_428_v74).Cardinality()), _437_v83, globalState))
				r0 = (_451_v91).Equals(_451_v91)
				var _452_v92 _dafny.MultiSet
				_ = _452_v92
				_452_v92 = _dafny.MultiSetOf(_406_v54)
				r0 = ((_429_v75).Plus(_dafny.IntOfUint32((_428_v74).Cardinality()))).Cmp((_452_v92).Cardinality()) > 0
				r0 = _406_v54
				var _453_v93 _dafny.MultiSet
				_ = _453_v93
				_453_v93 = _dafny.MultiSetOf(_dafny.IntOfInt64(627))
				(_this).M13(_429_v75, (_453_v93).Intersection(_453_v93), !((_dafny.SetOf(false)).IsDisjointFrom(_434_v80)), globalState)
			}
		}
		var _454_v94 bool
		_ = _454_v94
		_454_v94 = true
		(globalState).F4 = _454_v94
		var _455_v95 _dafny.Array
		_ = _455_v95
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_11
		var _nw80 _dafny.Array
		_ = _nw80
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw80 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = (func(_456_v94 bool) func(_dafny.Int) bool {
				return func(_457_i5 _dafny.Int) bool {
					return _456_v94
				}
			})(_454_v94)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw80 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw80).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw80).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_455_v95 = _nw80
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_455_v95), 0))
		_ = _index77
		(_455_v95).ArraySet1((func() bool {
			if _454_v94 {
				return _454_v94
			}
			return _454_v94
		})(), (_index77).Int())
		var _458_v96 _dafny.Array
		_ = _458_v96
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
		_ = _nw81
		_458_v96 = _nw81
		var _459_v97 _dafny.Array
		_ = _459_v97
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
		_ = _nw82
		_459_v97 = _nw82
		var _460_v98 _dafny.Map
		_ = _460_v98
		_460_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_458_v96, _459_v97)
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_455_v95), 0))
		_ = _index78
		(_455_v95).ArraySet1((_460_v98).Equals(_460_v98), (_index78).Int())
		var _461_i6 _dafny.Int
		_ = _461_i6
		_461_i6 = _dafny.Zero
		{
			for (_455_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_455_v95), 0))).Int()).(bool) {
				{
					if (_461_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_461_i6 = (_461_i6).Plus(_dafny.One)
					var _462_v99 _dafny.Int
					_ = _462_v99
					_462_v99 = _dafny.IntOfInt64(-172)
					var _463_v100 _dafny.Sequence
					_ = _463_v100
					_463_v100 = _dafny.SeqOf(_462_v99, _462_v99, _462_v99)
					var _464_v101 _dafny.MultiSet
					_ = _464_v101
					_464_v101 = _dafny.MultiSetOf(_dafny.SeqOf(_462_v99), _463_v100)
					var _465_v102 _dafny.Map
					_ = _465_v102
					_465_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_462_v99), _464_v101)
					var _466_v103 D3
					_ = _466_v103
					_466_v103 = Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v99, _462_v99))
					var _467_v104 _dafny.MultiSet
					_ = _467_v104
					_467_v104 = _dafny.MultiSetOf(_dafny.SeqOf(true))
					_465_v102 = (_465_v102).Update(Companion_Default___.Fm26(_466_v103, _462_v99, _467_v104, globalState), _464_v101)
					_462_v99 = _462_v99
					var _468_v105 *C0
					_ = _468_v105
					var _nw83 *C0 = New_C0_()
					_ = _nw83
					_nw83.Ctor__()
					_468_v105 = _nw83
					(globalState).F8 = ((_462_v99).Plus(_462_v99)).Plus(_462_v99)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _469_v106 _dafny.Int
		_ = _469_v106
		_469_v106 = _dafny.IntOfInt64(160)
		var _470_v107 _dafny.Map
		_ = _470_v107
		_470_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_v106, _469_v106)
		(globalState).F8 = (Companion_Default___.SafeModuloInt(_469_v106, _469_v106)).Times((_dafny.IntOfInt64(602)).Times((_470_v107).Cardinality()))
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_459_v97), 0))
		_ = _index79
		(_459_v97).ArraySet1(_469_v106, (_index79).Int())
		var _471_v108 _dafny.Sequence
		_ = _471_v108
		_471_v108 = _dafny.UnicodeSeqOfUtf8Bytes("p")
		var _472_v109 D2
		_ = _472_v109
		_472_v109 = Companion_D2_.Create_DC6_(_469_v106, true, (_455_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_455_v95), 0))).Int()).(bool))
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_459_v97), 0))
		_ = _index80
		var _rhs75 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_469_v106, (_469_v106).Times((_dafny.Zero).Minus(_469_v106))))
		_ = _rhs75
		var _rhs76 bool = (_455_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_455_v95), 0))).Int()).(bool)
		_ = _rhs76
		var _rhs77 bool = (_472_v109).Dtor_cf8()
		_ = _rhs77
		var _rhs78 _dafny.Sequence = _471_v108
		_ = _rhs78
		var _rhs79 bool = false
		_ = _rhs79
		var _lhs62 _dafny.Array = _459_v97
		_ = _lhs62
		var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_459_v97), 0))
		_ = _lhs63
		var _lhs64 *GlobalState = globalState
		_ = _lhs64
		(_lhs62).ArraySet1(_rhs75, (_lhs63).Int())
		r1 = _rhs76
		_lhs64.F4 = _rhs77
		_471_v108 = _rhs78
		r1 = _rhs79
		r0 = _454_v94
		var _pat_let_tv13 = _459_v97
		_ = _pat_let_tv13
		var _pat_let_tv14 = _459_v97
		_ = _pat_let_tv14
		var _pat_let_tv15 = _454_v94
		_ = _pat_let_tv15
		var _pat_let_tv16 = _455_v95
		_ = _pat_let_tv16
		var _pat_let_tv17 = _455_v95
		_ = _pat_let_tv17
		var _pat_let_tv18 = _454_v94
		_ = _pat_let_tv18
		r1 = func(_source12 D5) bool {
			if _source12.Is_DC14() {
				var _473___mcc_h4 _dafny.Int = _source12.Get_().(D5_DC14).Cf20
				_ = _473___mcc_h4
				var _474___mcc_h5 bool = _source12.Get_().(D5_DC14).Cf21
				_ = _474___mcc_h5
				var _475_cf21 bool = _474___mcc_h5
				_ = _475_cf21
				var _476_cf20 _dafny.Int = _473___mcc_h4
				_ = _476_cf20
				return (Companion_D2_.Create_DC6_((_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(_dafny.Int), true, _475_cf21)).Dtor_cf8()
			} else if _source12.Is_DC15() {
				return (_dafny.SetOf(_pat_let_tv15)).IsSubsetOf(_dafny.SetOf((_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(bool)))
			} else {
				var _477___mcc_h6 _dafny.Map = _source12.Get_().(D5_DC13).Cf19
				_ = _477___mcc_h6
				var _478_cf19 _dafny.Map = _477___mcc_h6
				_ = _478_cf19
				return _pat_let_tv18
			}
		}(Companion_D5_.Create_DC15_())
		return r0, r1
	}
}
func (_this *C1) M13(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _479_v0 _dafny.Sequence
		_ = _479_v0
		_479_v0 = _dafny.UnicodeSeqOfUtf8Bytes("eahalakud")
		if (p1).Contains((_dafny.IntOfUint32((_479_v0).Cardinality())).Minus(_dafny.IntOfInt64(150))) {
			var _480_v1 _dafny.Array
			_ = _480_v1
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw84
			_480_v1 = _nw84
			var _481_v2 _dafny.Sequence
			_ = _481_v2
			_481_v2 = _dafny.SeqOf(_480_v1)
			_481_v2 = _481_v2
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))
			_ = _index81
			(_480_v1).ArraySet1(p0, (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))
			_ = _index82
			(_480_v1).ArraySet1(p0, (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))
			_ = _index83
			(_480_v1).ArraySet1((_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int), (_index83).Int())
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_12
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Int = (func(_482_v1 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_483_i0 _dafny.Int) _dafny.Int {
						return (_483_i0).Plus((_482_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_482_v1), 0))).Int()).(_dafny.Int))
					}
				})(_480_v1)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw85 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw85).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw85).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_480_v1 = _nw85
			var _484_v3 _dafny.Sequence
			_ = _484_v3
			_484_v3 = _dafny.SeqOf(p0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-256)))
			var _485_v4 D9
			_ = _485_v4
			_485_v4 = Companion_D9_.Create_DC21_(false, _dafny.SeqOf(p2, p2, (_this).Fm3(_484_v3, p1, p2, globalState), p2), p2)
			var _486_v5 _dafny.Sequence
			_ = _486_v5
			_486_v5 = _dafny.SeqOf(!(p2), !((_this).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(586))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_487_v1 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_488_i1 _dafny.Int) _dafny.Int {
					return (_487_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_487_v1), 0))).Int()).(_dafny.Int)
				}
			})(_480_v1))), p1, !(false), globalState)), p2)
			var _pat_let_tv19 = p2
			_ = _pat_let_tv19
			if _dafny.Companion_Sequence_.Equal((func(_pat_let15_0 D9) D9 {
				return func(_489_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let16_0 bool) D9 {
						return func(_490_dt__update_hcf26_h0 bool) D9 {
							return Companion_D9_.Create_DC21_(_490_dt__update_hcf26_h0, (_489_dt__update__tmp_h0).Dtor_cf27(), (_489_dt__update__tmp_h0).Dtor_cf28())
						}(_pat_let16_0)
					}(_pat_let_tv19)
				}(_pat_let15_0)
			}(_485_v4)).Dtor_cf27(), _dafny.Companion_Sequence_.Concatenate(_486_v5, _486_v5)) {
				(globalState).F3 = (func() _dafny.Int {
					if p2 {
						return p0
					}
					return _dafny.IntOfInt64(655)
				})()
				_480_v1 = _480_v1
				(globalState).F4 = p2
				var _491_v6 D2
				_ = _491_v6
				_491_v6 = Companion_D2_.Create_DC6_(p0, (p0).Cmp(p0) == 0, p2)
				_491_v6 = _491_v6
				(globalState).F8 = (_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int)
			} else {
				var _492_v7 *C0
				_ = _492_v7
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__()
				_492_v7 = _nw86
				var _493_v9 _dafny.Array
				_ = _493_v9
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_13
				var _nw87 _dafny.Array
				_ = _nw87
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw87 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Map = (func(_494_p2 bool) func(_dafny.Int) _dafny.Map {
						return func(_495_i2 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_p2, _494_p2)).Merge(func() _dafny.Map {
								var _coll58 = _dafny.NewMapBuilder()
								_ = _coll58
								for _iter59 := _dafny.Iterate((_dafny.SetOf(_494_p2, false)).Elements()); ; {
									_compr_58, _ok59 := _iter59()
									if !_ok59 {
										break
									}
									var _496_v8 bool
									_496_v8 = interface{}(_compr_58).(bool)
									if (_dafny.SetOf(_494_p2, false)).Contains(_496_v8) {
										_coll58.Add(_496_v8, _494_p2)
									}
								}
								return _coll58.ToMap()
							}())
						}
					})(p2)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw87 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw87).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw87).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_493_v9 = _nw87
				var _497_v10 bool
				_ = _497_v10
				_497_v10 = p2
				var _498_v11 D2
				_ = _498_v11
				_498_v11 = Companion_D2_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p2))).Cardinality(), p2, p2)
				var _499_v12 _dafny.Map
				_ = _499_v12
				_499_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_497_v10, !((_498_v11).Dtor_cf8()))
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_493_v9), 0))
				_ = _index84
				(_493_v9).ArraySet1(_499_v12, (_index84).Int())
				var _500_v14 _dafny.Set
				_ = _500_v14
				_500_v14 = _dafny.SetOf(_497_v10, _497_v10)
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_493_v9), 0))
				_ = _index85
				(_493_v9).ArraySet1(func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter60 := _dafny.Iterate((func() _dafny.Set {
						var _coll60 = _dafny.NewBuilder()
						_ = _coll60
						for _iter61 := _dafny.Iterate((_500_v14).Elements()); ; {
							_compr_60, _ok61 := _iter61()
							if !_ok61 {
								break
							}
							var _501_v15 bool
							_501_v15 = interface{}(_compr_60).(bool)
							if (_500_v14).Contains(_501_v15) {
								_coll60.Add(_501_v15)
							}
						}
						return _coll60.ToSet()
					}()).Elements()); ; {
						_compr_59, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _502_v13 bool
						_502_v13 = interface{}(_compr_59).(bool)
						if (func() _dafny.Set {
							var _coll61 = _dafny.NewBuilder()
							_ = _coll61
							for _iter62 := _dafny.Iterate((_500_v14).Elements()); ; {
								_compr_61, _ok62 := _iter62()
								if !_ok62 {
									break
								}
								var _503_v15 bool
								_503_v15 = interface{}(_compr_61).(bool)
								if (_500_v14).Contains(_503_v15) {
									_coll61.Add(_503_v15)
								}
							}
							return _coll61.ToSet()
						}()).Contains(_502_v13) {
							_coll59.Add(_502_v13, p2)
						}
					}
					return _coll59.ToMap()
				}(), (_index85).Int())
				(globalState).F8 = (_492_v7).Fm19(globalState)
				var _504_v16 _dafny.CodePoint
				_ = _504_v16
				_504_v16 = _dafny.CodePoint('l')
				var _505_v17 _dafny.Map
				_ = _505_v17
				_505_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v7, _504_v16)
				var _506_v18 _dafny.Map
				_ = _506_v18
				_506_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-817))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}((func(_507_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_508_i5 _dafny.Int) _dafny.CodePoint {
						return _507_v16
					}
				})(_504_v16)))).Cardinality()))
				var _509_v19 _dafny.Map
				_ = _509_v19
				_509_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _510_v20 _dafny.Map
				_ = _510_v20
				_510_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v19, _504_v16)
				var _511_v21 _dafny.Array
				_ = _511_v21
				var _nwElement0_17 _dafny.Sequence = _484_v3
				_ = _nwElement0_17
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(15))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_17, 0)
				(_nw88).ArraySet1(_484_v3, 1)
				(_nw88).ArraySet1(_484_v3, 2)
				(_nw88).ArraySet1(_484_v3, 3)
				(_nw88).ArraySet1((func() _dafny.Sequence {
					if !((_this).Fm3(_484_v3, p1, p2, globalState)) {
						return _484_v3
					}
					return _484_v3
				})(), 4)
				(_nw88).ArraySet1(_484_v3, 5)
				(_nw88).ArraySet1(_dafny.SeqOf((_505_v17).Cardinality(), (_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int)), 6)
				(_nw88).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_512_v1 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_513_i3 _dafny.Int) _dafny.Int {
						return (_512_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_512_v1), 0))).Int()).(_dafny.Int)
					}
				})(_480_v1))), 7)
				(_nw88).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(817))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}(func(_514_i4 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(782)
				})), _484_v3), 8)
				(_nw88).ArraySet1(_484_v3, 9)
				(_nw88).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int), _dafny.CodePoint('j'), p2, _506_v18, globalState), _dafny.SeqOf((_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int), (_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int))), 10)
				(_nw88).ArraySet1(_484_v3, 11)
				(_nw88).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus((_480_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_480_v1), 0))).Int()).(_dafny.Int)), (_510_v20).Cardinality()), 12)
				(_nw88).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_484_v3, _484_v3), 13)
				(_nw88).ArraySet1(_484_v3, 14)
				_511_v21 = _nw88
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_511_v21), 0))
				_ = _index86
				(_511_v21).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_515_v1 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_516_i6 _dafny.Int) _dafny.Int {
						return (_515_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_515_v1), 0))).Int()).(_dafny.Int)
					}
				})(_480_v1))), (_index86).Int())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_511_v21), 0))
				_ = _index87
				(_511_v21).ArraySet1(_484_v3, (_index87).Int())
				var _517_v22 _dafny.Array
				_ = _517_v22
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw89
				_517_v22 = _nw89
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_517_v22), 0))
				_ = _index88
				(_517_v22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fusvcouy"), (_index88).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_517_v22), 0))
				_ = _index89
				(_517_v22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("k"), (_index89).Int())
			}
		} else {
			var _518_v23 _dafny.Array
			_ = _518_v23
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_14
			var _nw90 _dafny.Array
			_ = _nw90
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw90 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_519_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_520_i7 _dafny.Int) _dafny.Int {
						return (_520_i7).Minus(_519_p0)
					}
				})(p0)
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
			_518_v23 = _nw90
			var _521_v24 D0
			_ = _521_v24
			_521_v24 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_((_dafny.Zero).Minus(p0), p0))
			var _522_v25 _dafny.Sequence
			_ = _522_v25
			_522_v25 = _dafny.SeqOf(_521_v24)
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
			_ = _index90
			(_518_v23).ArraySet1(_dafny.IntOfUint32((_522_v25).Cardinality()), (_index90).Int())
			var _523_v26 _dafny.Sequence
			_ = _523_v26
			_523_v26 = _dafny.SeqOf(p2, (p2) || (true), true)
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
			_ = _index91
			var _rhs80 _dafny.Int = ((func() _dafny.Int {
				if p2 {
					return p0
				}
				return p0
			})()).Plus(p0)
			_ = _rhs80
			var _rhs81 _dafny.Int = _dafny.IntOfUint32((_523_v26).Cardinality())
			_ = _rhs81
			var _lhs65 *GlobalState = globalState
			_ = _lhs65
			var _lhs66 _dafny.Array = _518_v23
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
			_ = _lhs67
			_lhs65.F3 = _rhs80
			(_lhs66).ArraySet1(_rhs81, (_lhs67).Int())
			(globalState).F4 = p2
			var _524_v27 _dafny.Set
			_ = _524_v27
			_524_v27 = _dafny.SetOf(p2)
			var _525_v28 D9
			_ = _525_v28
			_525_v28 = Companion_D9_.Create_DC20_(_524_v27)
			var _526_v29 _dafny.Sequence
			_ = _526_v29
			_526_v29 = _dafny.SeqOf(_525_v28)
			var _rhs82 _dafny.Int = (_dafny.Zero).Minus((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
			_ = _rhs82
			var _rhs83 _dafny.Sequence = _526_v29
			_ = _rhs83
			var _lhs68 *GlobalState = globalState
			_ = _lhs68
			_lhs68.F8 = _rhs82
			_526_v29 = _rhs83
			var _527_v30 _dafny.Sequence
			_ = _527_v30
			_527_v30 = _dafny.SeqOf((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), p0, (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), p0, p0)
			var _528_v31 _dafny.Map
			_ = _528_v31
			_528_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p0)).Cardinality(), p0)
			var _529_v32 D3
			_ = _529_v32
			_529_v32 = Companion_D3_.Create_DC8_(_528_v31)
			var _530_v33 _dafny.MultiSet
			_ = _530_v33
			_530_v33 = _dafny.MultiSetOf(_523_v26)
			var _531_v34 _dafny.Array
			_ = _531_v34
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_15
			var _nw91 _dafny.Array
			_ = _nw91
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw91 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.CodePoint = func(_532_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw91 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw91).ArraySet1CodePoint(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw91).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_531_v34 = _nw91
			var _533_v35 _dafny.Map
			_ = _533_v35
			_533_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v34, _dafny.IntOfInt64(-607))
			if (_this).Fm3(_527_v30, (p1).Union(_dafny.MultiSetOf(p0, Companion_Default___.Fm26(_529_v32, Companion_Default___.Fm26(_529_v32, (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), _dafny.MultiSetFromSeq(_dafny.SeqOf(_523_v26)), globalState), _530_v33, globalState), (_533_v35).Cardinality(), p0, p0)), p2, globalState) {
				(globalState).F4 = p2
				var _534_v36 _dafny.Map
				_ = _534_v36
				_534_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
				(globalState).F3 = (func() _dafny.Int {
					if (_534_v36).Contains(p2) {
						return (_534_v36).Get(p2).(_dafny.Int)
					}
					return p0
				})()
				_479_v0 = _479_v0
				_479_v0 = _479_v0
				var _535_v37 _dafny.Set
				_ = _535_v37
				_535_v37 = _dafny.SetOf((p1).Cardinality())
				_535_v37 = (_dafny.SetOf((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Set {
					var _coll62 = _dafny.NewBuilder()
					_ = _coll62
					for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(331), _dafny.IntOfInt64(658))); ; {
						_compr_62, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _536_v38 _dafny.Int
						_536_v38 = interface{}(_compr_62).(_dafny.Int)
						if ((_dafny.IntOfInt64(331)).Cmp(_536_v38) <= 0) && ((_536_v38).Cmp(_dafny.IntOfInt64(658)) < 0) {
							_coll62.Add((_536_v38).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_479_v0).Cardinality()), p0)).Cardinality()))
						}
					}
					return _coll62.ToSet()
				}()).Cardinality())).Cardinality(), (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))).Union(_535_v37)
			} else {
				_521_v24 = _521_v24
				var _537_v39 _dafny.Map
				_ = _537_v39
				_537_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(469), _dafny.SeqOf(p0))
				var _538_v40 _dafny.Array
				_ = _538_v40
				var _nwElement0_18 _dafny.Sequence = _dafny.SeqOf((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
				_ = _nwElement0_18
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(9))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_18, 0)
				(_nw92).ArraySet1(_527_v30, 1)
				(_nw92).ArraySet1(_527_v30, 2)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Update(_527_v30, (Companion_Default___.SafeIndex((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_527_v30).Cardinality()))).Uint32(), _dafny.IntOfInt64(645)), 3)
				(_nw92).ArraySet1(_527_v30, 4)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_537_v39).Contains(_dafny.IntOfInt64(182)) {
						return (_537_v39).Get(_dafny.IntOfInt64(182)).(_dafny.Sequence)
					}
					return _527_v30
				})(), _dafny.Companion_Sequence_.Update(_527_v30, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_527_v30).Cardinality()))).Uint32(), _dafny.IntOfUint32((_479_v0).Cardinality()))), 5)
				(_nw92).ArraySet1(_527_v30, 6)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_527_v30, _527_v30), 7)
				(_nw92).ArraySet1(_527_v30, 8)
				_538_v40 = _nw92
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_538_v40), 0))
				_ = _index92
				(_538_v40).ArraySet1(_527_v30, (_index92).Int())
				var _539_v41 _dafny.CodePoint
				_ = _539_v41
				_539_v41 = _dafny.CodePoint('n')
				var _540_v42 _dafny.Array
				_ = _540_v42
				var _nwElement0_19 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("eypc")
				_ = _nwElement0_19
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(4))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_19, 0)
				(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_479_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_479_v0).Cardinality()))).Uint32(), _539_v41), (Companion_Default___.SafeIndex((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_479_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_479_v0).Cardinality()))).Uint32(), _539_v41)).Cardinality()))).Uint32(), _539_v41), _479_v0), 1)
				(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_479_v0, _479_v0), 2)
				(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_479_v0, _479_v0), 3)
				_540_v42 = _nw93
				var _541_v43 D2
				_ = _541_v43
				_541_v43 = Companion_D2_.Create_DC7_(_539_v41, p2, (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_538_v40), 0))
				_ = _index93
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
				_ = _index94
				var _rhs84 _dafny.Sequence = _527_v30
				_ = _rhs84
				var _rhs85 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), ((_dafny.Zero).Minus((_dafny.Zero).Minus((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int)))).Plus(p0))
				_ = _rhs85
				var _rhs86 _dafny.Array = _540_v42
				_ = _rhs86
				var _rhs87 bool = !(_528_v31).Contains((_541_v43).Dtor_cf11())
				_ = _rhs87
				var _lhs69 _dafny.Array = _538_v40
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_538_v40), 0))
				_ = _lhs70
				var _lhs71 _dafny.Array = _518_v23
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
				_ = _lhs72
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				(_lhs69).ArraySet1(_rhs84, (_lhs70).Int())
				(_lhs71).ArraySet1(_rhs85, (_lhs72).Int())
				_540_v42 = _rhs86
				_lhs73.F4 = _rhs87
				var _542_v44 bool
				_ = _542_v44
				var _out16 bool
				_ = _out16
				_out16 = (_this).M14(p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)), p2, _dafny.IntOfUint32((_527_v30).Cardinality()), globalState)
				_542_v44 = _out16
				var _543_v45 _dafny.Sequence
				_ = _543_v45
				_543_v45 = _dafny.SeqOf(_523_v26)
				var _544_v46 D5
				_ = _544_v46
				_544_v46 = Companion_D5_.Create_DC14_(Companion_Default___.Fm26(_529_v32, (_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), _dafny.MultiSetFromSeq(_dafny.SeqOf(_523_v26)), globalState), _542_v44)
				var _545_v47 _dafny.Map
				_ = _545_v47
				_545_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v46, p0)
				var _546_v48 _dafny.Map
				_ = _546_v48
				_546_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v45, _545_v47)
				_546_v48 = (_546_v48).Update(_543_v45, _545_v47)
				(globalState).F4 = (_dafny.IntOfInt64(628)).Cmp(Companion_Default___.Fm26(_529_v32, ((_538_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_538_v40), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_538_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_538_v40), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), _530_v33, globalState)) != 0
			}
			var _547_v49 D0
			_ = _547_v49
			_547_v49 = Companion_D0_.Create_DC0_((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
			_ = _index95
			var _rhs88 _dafny.Int = p0
			_ = _rhs88
			var _rhs89 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_479_v0, _479_v0)).Cardinality())
			_ = _rhs89
			var _rhs90 bool = p2
			_ = _rhs90
			var _rhs91 D0 = (func() D0 {
				if p2 {
					return Companion_D0_.Create_DC0_((_518_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))).Int()).(_dafny.Int))
				}
				return _547_v49
			})()
			_ = _rhs91
			var _rhs92 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_527_v30, _dafny.Companion_Sequence_.Concatenate(_527_v30, _527_v30))
			_ = _rhs92
			var _lhs74 _dafny.Array = _518_v23
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_518_v23), 0))
			_ = _lhs75
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			(_lhs74).ArraySet1(_rhs88, (_lhs75).Int())
			_lhs76.F8 = _rhs89
			_lhs77.F4 = _rhs90
			_547_v49 = _rhs91
			_lhs78.F0 = _rhs92
		}
		var _548_v50 _dafny.Map
		_ = _548_v50
		_548_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		_548_v50 = (_548_v50).Update(p0, (func() bool {
			if p2 {
				return p2
			}
			return p2
		})())
		(globalState).F4 = p2
		var _549_v51 _dafny.Sequence
		_ = _549_v51
		_549_v51 = _dafny.SeqOf(_dafny.IntOfInt64(72), p0, p0, p0)
		var _550_v52 _dafny.Map
		_ = _550_v52
		_550_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_549_v51, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_549_v51).Cardinality()))).Uint32(), p0), p2)
		var _551_v53 _dafny.Map
		_ = _551_v53
		_551_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v52, p2)
		if ((func() bool {
			if (_551_v53).Contains(_550_v52) {
				return (_551_v53).Get(_550_v52).(bool)
			}
			return p2
		})()) && (p2) {
			var _552_v54 *C0
			_ = _552_v54
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__()
			_552_v54 = _nw94
			(globalState).F8 = p0
			if (_this).Fm3(_549_v51, p1, p2, globalState) {
				var _553_v55 *C0
				_ = _553_v55
				var _nw95 *C0 = New_C0_()
				_ = _nw95
				_nw95.Ctor__()
				_553_v55 = _nw95
				var _rhs93 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("meyetabe")
				_ = _rhs93
				var _rhs94 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _rhs94
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_479_v0 = _rhs93
				_lhs79.F3 = _rhs94
				var _554_v56 *C0
				_ = _554_v56
				var _nw96 *C0 = New_C0_()
				_ = _nw96
				_nw96.Ctor__()
				_554_v56 = _nw96
				(globalState).F4 = (_this).Fm3(_549_v51, p1, (func() bool {
					if p2 {
						return p2
					}
					return p2
				})(), globalState)
				var _555_v57 _dafny.Array
				_ = _555_v57
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_16
				var _nw97 _dafny.Array
				_ = _nw97
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw97 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) bool = (func(_556_p2 bool) func(_dafny.Int) bool {
						return func(_557_i9 _dafny.Int) bool {
							return _556_p2
						}
					})(p2)
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
				_555_v57 = _nw97
				var _558_v58 D5
				_ = _558_v58
				_558_v58 = Companion_D5_.Create_DC14_(p0, p2)
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_555_v57), 0))
				_ = _index96
				(_555_v57).ArraySet1(((_558_v58).Dtor_cf21()) == (p2), (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_555_v57), 0))
				_ = _index97
				(_555_v57).ArraySet1(p2, (_index97).Int())
				var _559_v59 _dafny.Map
				_ = _559_v59
				_559_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24(globalState), _dafny.MultiSetOf(p2, p2, p2, p2))
				var _560_v60 _dafny.MultiSet
				_ = _560_v60
				_560_v60 = _dafny.MultiSetOf(p2)
				var _561_v61 _dafny.Sequence
				_ = _561_v61
				_561_v61 = _dafny.SeqOf(_dafny.MultiSetOf(p2, p2), _560_v60)
				var _562_v62 _dafny.MultiSet
				_ = _562_v62
				_562_v62 = _560_v60
				var _563_v63 _dafny.Map
				_ = _563_v63
				_563_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _560_v60)
				var _564_v64 _dafny.Array
				_ = _564_v64
				var _nwElement0_20 _dafny.MultiSet = ((_560_v60).Update(p2, Companion_Default___.Abs(p0))).Intersection(Companion_Default___.Fm0(p0, globalState))
				_ = _nwElement0_20
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(27))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_20, 0)
				(_nw98).ArraySet1((_561_v61).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_561_v61).Cardinality()))).Uint32()).(_dafny.MultiSet), 1)
				(_nw98).ArraySet1((_560_v60).Difference(_dafny.MultiSetOf(true, p2, p2, p2)), 2)
				(_nw98).ArraySet1((func() _dafny.MultiSet {
					if p2 {
						return _560_v60
					}
					return _dafny.MultiSetFromSeq(_dafny.SeqOf(p2, !(p2)))
				})(), 3)
				(_nw98).ArraySet1(_560_v60, 4)
				(_nw98).ArraySet1((_562_v62), 5)
				(_nw98).ArraySet1(_dafny.MultiSetOf(p2), 6)
				(_nw98).ArraySet1((_dafny.MultiSetOf(p2, p2)).Difference(_560_v60), 7)
				(_nw98).ArraySet1((_560_v60).Intersection(_560_v60), 8)
				(_nw98).ArraySet1(_560_v60, 9)
				(_nw98).ArraySet1(_560_v60, 10)
				(_nw98).ArraySet1(_560_v60, 11)
				(_nw98).ArraySet1(_560_v60, 12)
				(_nw98).ArraySet1((func() _dafny.MultiSet {
					if p2 {
						return (_561_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_549_v51).Cardinality()), _dafny.IntOfUint32((_561_v61).Cardinality()))).Uint32()).(_dafny.MultiSet)
					}
					return _560_v60
				})(), 13)
				(_nw98).ArraySet1(_560_v60, 14)
				(_nw98).ArraySet1(_560_v60, 15)
				(_nw98).ArraySet1(_560_v60, 16)
				(_nw98).ArraySet1((_560_v60).Intersection(_560_v60), 17)
				(_nw98).ArraySet1(_560_v60, 18)
				(_nw98).ArraySet1(_560_v60, 19)
				(_nw98).ArraySet1((_560_v60).Update(p2, Companion_Default___.Abs(p0)), 20)
				(_nw98).ArraySet1(_560_v60, 21)
				(_nw98).ArraySet1((func() _dafny.MultiSet {
					if (_563_v63).Contains(p0) {
						return (_563_v63).Get(p0).(_dafny.MultiSet)
					}
					return _560_v60
				})(), 22)
				(_nw98).ArraySet1(_560_v60, 23)
				(_nw98).ArraySet1(_560_v60, 24)
				(_nw98).ArraySet1((_560_v60).Intersection(_560_v60), 25)
				(_nw98).ArraySet1((_562_v62), 26)
				_564_v64 = _nw98
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_564_v64), 0))
				_ = _index98
				(_564_v64).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_479_v0).Cardinality()), globalState), (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_555_v57), 0))
				_ = _index99
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_555_v57), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_564_v64), 0))
				_ = _index101
				var _rhs95 _dafny.MultiSet = Companion_Default___.Fm27(true, globalState)
				_ = _rhs95
				var _rhs96 bool = p2
				_ = _rhs96
				var _rhs97 bool = !(p2) || (p2)
				_ = _rhs97
				var _rhs98 _dafny.Map = _559_v59
				_ = _rhs98
				var _rhs99 _dafny.MultiSet = _560_v60
				_ = _rhs99
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 _dafny.Array = _555_v57
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_555_v57), 0))
				_ = _lhs82
				var _lhs83 _dafny.Array = _555_v57
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_555_v57), 0))
				_ = _lhs84
				var _lhs85 _dafny.Array = _564_v64
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_564_v64), 0))
				_ = _lhs86
				_lhs80.F2 = _rhs95
				(_lhs81).ArraySet1(_rhs96, (_lhs82).Int())
				(_lhs83).ArraySet1(_rhs97, (_lhs84).Int())
				_559_v59 = _rhs98
				(_lhs85).ArraySet1(_rhs99, (_lhs86).Int())
			} else {
				var _565_v65 _dafny.Array
				_ = _565_v65
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_17
				var _nw99 _dafny.Array
				_ = _nw99
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw99 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) bool = (func(_566_p0 _dafny.Int) func(_dafny.Int) bool {
						return func(_567_i10 _dafny.Int) bool {
							return (_566_p0).Cmp(_566_p0) != 0
						}
					})(p0)
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
				_565_v65 = _nw99
				var _568_v66 _dafny.CodePoint
				_ = _568_v66
				_568_v66 = _dafny.CodePoint('h')
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))
				_ = _index102
				(_565_v65).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ijcbji"), _568_v66), (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))
				_ = _index103
				(_565_v65).ArraySet1(p2, (_index103).Int())
				var _569_v67 _dafny.MultiSet
				_ = _569_v67
				_569_v67 = _dafny.MultiSetOf((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool), (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool), p2, (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))
				var _570_v68 D2
				_ = _570_v68
				_570_v68 = Companion_D2_.Create_DC6_(p0, (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool), p2)
				var _571_v69 _dafny.Sequence
				_ = _571_v69
				_571_v69 = _dafny.SeqOf(p2)
				var _572_v70 _dafny.Array
				_ = _572_v70
				var _nwElement0_21 _dafny.MultiSet = _569_v67
				_ = _nwElement0_21
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(20))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_21, 0)
				(_nw100).ArraySet1(_569_v67, 1)
				(_nw100).ArraySet1((_dafny.MultiSetOf((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))).Difference(_569_v67), 2)
				(_nw100).ArraySet1(_569_v67, 3)
				(_nw100).ArraySet1((_dafny.MultiSetOf((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))).Union(_dafny.MultiSetOf((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))), 4)
				(_nw100).ArraySet1(_569_v67, 5)
				(_nw100).ArraySet1(_dafny.MultiSetOf(p2, (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool), false, !(p2)), 6)
				(_nw100).ArraySet1(_569_v67, 7)
				(_nw100).ArraySet1(_569_v67, 8)
				(_nw100).ArraySet1((func() _dafny.MultiSet {
					if (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool) {
						return _569_v67
					}
					return _569_v67
				})(), 9)
				(_nw100).ArraySet1(_dafny.MultiSetOf((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool)), 10)
				(_nw100).ArraySet1((_569_v67).Difference(_dafny.MultiSetOf((_570_v68).Dtor_cf7(), p2, (_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))), 11)
				(_nw100).ArraySet1(_dafny.MultiSetOf(p2), 12)
				(_nw100).ArraySet1(_569_v67, 13)
				(_nw100).ArraySet1(_dafny.MultiSetOf((Companion_D9_.Create_DC21_(!(!((_565_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_565_v65), 0))).Int()).(bool))), _571_v69, !(p2))).Dtor_cf28()), 14)
				(_nw100).ArraySet1(_569_v67, 15)
				(_nw100).ArraySet1(_569_v67, 16)
				(_nw100).ArraySet1(_dafny.MultiSetFromSeq(_571_v69), 17)
				(_nw100).ArraySet1(_569_v67, 18)
				(_nw100).ArraySet1(_569_v67, 19)
				_572_v70 = _nw100
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_572_v70), 0))
				_ = _index104
				(_572_v70).ArraySet1(_569_v67, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_572_v70), 0))
				_ = _index105
				(_572_v70).ArraySet1((_569_v67).Union(_dafny.MultiSetOf(p2)), (_index105).Int())
				var _573_v71 _dafny.Array
				_ = _573_v71
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw101
				_573_v71 = _nw101
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_573_v71), 0))
				_ = _index106
				(_573_v71).ArraySet1(p0, (_index106).Int())
				var _574_v72 _dafny.Map
				_ = _574_v72
				_574_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-798), p0)
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_573_v71), 0))
				_ = _index107
				(_573_v71).ArraySet1((((_574_v72).Merge(_574_v72)).Cardinality()).Plus(p0), (_index107).Int())
				(globalState).F8 = (p0).Times(_dafny.IntOfInt64(975))
				_479_v0 = _479_v0
			}
			_552_v54 = _552_v54
			var _575_v73 _dafny.Map
			_ = _575_v73
			_575_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_479_v0).Cardinality()))
			(globalState).F3 = (((Companion_D12_.Create_DC26_(_575_v73)).Dtor_cf35()).Cardinality()).Plus(((_548_v50).Merge(func() _dafny.Map {
				var _coll63 = _dafny.NewMapBuilder()
				_ = _coll63
				for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(958), _dafny.IntOfInt64(-919))); ; {
					_compr_63, _ok64 := _iter64()
					if !_ok64 {
						break
					}
					var _576_v74 _dafny.Int
					_576_v74 = interface{}(_compr_63).(_dafny.Int)
					if ((_dafny.IntOfInt64(958)).Cmp(_576_v74) <= 0) && ((_576_v74).Cmp(_dafny.IntOfInt64(-919)) < 0) {
						_coll63.Add((_576_v74).Plus(p0), p2)
					}
				}
				return _coll63.ToMap()
			}())).Cardinality())
		} else {
			var _577_v75 _dafny.Set
			_ = _577_v75
			_577_v75 = _dafny.SetOf(p0, (_dafny.Zero).Minus(p0))
			if (_577_v75).IsProperSubsetOf((_577_v75).Intersection(Companion_Default___.Fm28(p0, globalState))) {
				(globalState).F3 = (_dafny.Zero).Minus(p0)
				var _578_v76 _dafny.Set
				_ = _578_v76
				_578_v76 = _dafny.SetOf(p2, p2, p2)
				var _579_v77 _dafny.MultiSet
				_ = _579_v77
				_579_v77 = _dafny.MultiSetOf(p2, true, p2, p2)
				var _580_v78 _dafny.Sequence
				_ = _580_v78
				_580_v78 = _dafny.SeqOf(_549_v51)
				var _581_v79 _dafny.Map
				_ = _581_v79
				_581_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(p2))
				var _582_v80 _dafny.Array
				_ = _582_v80
				var _nwElement0_22 bool = (p0).Cmp((_dafny.Zero).Minus(p0)) > 0
				_ = _nwElement0_22
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_22, 0)
				(_nw102).ArraySet1(p2, 1)
				(_nw102).ArraySet1((_578_v76).IsSubsetOf(_578_v76), 2)
				(_nw102).ArraySet1(p2, 3)
				(_nw102).ArraySet1(p2, 4)
				(_nw102).ArraySet1((_579_v77).IsDisjointFrom(_579_v77), 5)
				(_nw102).ArraySet1(p2, 6)
				(_nw102).ArraySet1(true, 7)
				(_nw102).ArraySet1((_this).Fm3(_549_v51, _dafny.MultiSetFromSeq((_580_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_580_v78).Cardinality()))).Uint32()).(_dafny.Sequence)), p2, globalState), 8)
				(_nw102).ArraySet1(p2, 9)
				(_nw102).ArraySet1(p2, 10)
				(_nw102).ArraySet1((p1).IsDisjointFrom(p1), 11)
				(_nw102).ArraySet1((func() bool {
					if (_581_v79).Contains(p2) {
						return (_581_v79).Get(p2).(bool)
					}
					return p2
				})(), 12)
				(_nw102).ArraySet1(p2, 13)
				(_nw102).ArraySet1(false, 14)
				(_nw102).ArraySet1(!(!(p2) || (p2)), 15)
				_582_v80 = _nw102
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_582_v80), 0))
				_ = _index108
				(_582_v80).ArraySet1(p2, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_582_v80), 0))
				_ = _index109
				(_582_v80).ArraySet1(!(false), (_index109).Int())
				var _583_v81 _dafny.Array
				_ = _583_v81
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
				_ = _nw103
				_583_v81 = _nw103
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_583_v81), 0))
				_ = _index110
				(_583_v81).ArraySet1(_581_v79, (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_583_v81), 0))
				_ = _index111
				(_583_v81).ArraySet1(_581_v79, (_index111).Int())
				var _584_v82 _dafny.Array
				_ = _584_v82
				var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
				_ = _nw104
				_584_v82 = _nw104
				var _585_v83 _dafny.CodePoint
				_ = _585_v83
				_585_v83 = _dafny.CodePoint('x')
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_584_v82), 0))
				_ = _index112
				(_584_v82).ArraySet1CodePoint(_585_v83, (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_584_v82), 0))
				_ = _index113
				(_584_v82).ArraySet1CodePoint(_585_v83, (_index113).Int())
				(globalState).F4 = ((_579_v77).Union((_dafny.MultiSetOf(p2)))).Equals(_579_v77)
			} else {
				(globalState).F4 = !(_548_v50).Contains((_577_v75).Cardinality())
				var _586_v84 _dafny.MultiSet
				_ = _586_v84
				_586_v84 = _dafny.MultiSetOf(!(p2) || (p2), p2)
				_586_v84 = _586_v84
				var _587_v85 _dafny.CodePoint
				_ = _587_v85
				_587_v85 = _dafny.CodePoint('v')
				_587_v85 = _587_v85
				(globalState).F3 = p0
				var _588_v86 _dafny.Array
				_ = _588_v86
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_18
				var _nw105 _dafny.Array
				_ = _nw105
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw105 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Int = (func(_589_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_590_i11 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_590_i11, _589_p0)
						}
					})(p0)
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
				_588_v86 = _nw105
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_588_v86), 0))
				_ = _index114
				(_588_v86).ArraySet1((_dafny.IntOfInt64(-359)).Plus(p0), (_index114).Int())
				var _591_v87 _dafny.Sequence
				_ = _591_v87
				_591_v87 = _dafny.SeqOf(!(!(true)), p2, p2, p2, p2)
				var _592_v88 _dafny.MultiSet
				_ = _592_v88
				_592_v88 = _dafny.MultiSetOf(_591_v87, _591_v87)
				var _593_v89 _dafny.Map
				_ = _593_v89
				_593_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _594_v90 _dafny.Map
				_ = _594_v90
				_594_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_593_v89).Cardinality(), _591_v87)
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_588_v86), 0))
				_ = _index115
				(_588_v86).ArraySet1((func() _dafny.Int {
					if (_592_v88).Contains((func() _dafny.Sequence {
						if (_594_v90).Contains(p0) {
							return (_594_v90).Get(p0).(_dafny.Sequence)
						}
						return _591_v87
					})()) {
						return (_592_v88).Multiplicity((func() _dafny.Sequence {
							if (_594_v90).Contains(p0) {
								return (_594_v90).Get(p0).(_dafny.Sequence)
							}
							return _591_v87
						})())
					}
					return p0
				})(), (_index115).Int())
			}
			(globalState).F4 = p2
			(globalState).F0 = _dafny.Companion_Sequence_.Concatenate(_549_v51, _549_v51)
			(globalState).F4 = ((func() bool {
				if p2 {
					return p2
				}
				return p2
			})()) || (p2)
			var _595_v91 _dafny.Array
			_ = _595_v91
			var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw106
			_595_v91 = _nw106
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_595_v91), 0))
			_ = _index116
			(_595_v91).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_596_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), (_index116).Int())
			var _597_v92 _dafny.Array
			_ = _597_v92
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_19
			var _nw107 _dafny.Array
			_ = _nw107
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw107 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) bool = (func(_598_p0 _dafny.Int) func(_dafny.Int) bool {
					return func(_599_i13 _dafny.Int) bool {
						return (_598_p0).Cmp(_598_p0) == 0
					}
				})(p0)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw107 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw107).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw107).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_597_v92 = _nw107
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_597_v92), 0))
			_ = _index117
			(_597_v92).ArraySet1(true, (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_595_v91), 0))
			_ = _index118
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_597_v92), 0))
			_ = _index119
			var _rhs100 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("miiuxq")
			_ = _rhs100
			var _rhs101 _dafny.Array = _597_v92
			_ = _rhs101
			var _rhs102 bool = false
			_ = _rhs102
			var _lhs87 _dafny.Array = _595_v91
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_595_v91), 0))
			_ = _lhs88
			var _lhs89 _dafny.Array = _597_v92
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_597_v92), 0))
			_ = _lhs90
			(_lhs87).ArraySet1(_rhs100, (_lhs88).Int())
			_597_v92 = _rhs101
			(_lhs89).ArraySet1(_rhs102, (_lhs90).Int())
		}
		var _600_i14 _dafny.Int
		_ = _600_i14
		_600_i14 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_479_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_611_i15 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))) {
				{
					if (_600_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_600_i14 = (_600_i14).Plus(_dafny.One)
					var _601_v93 _dafny.Sequence
					_ = _601_v93
					_601_v93 = _dafny.SeqOf(!(true))
					var _602_v94 _dafny.MultiSet
					_ = _602_v94
					_602_v94 = _dafny.MultiSetOf((_601_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.IntOfUint32((_601_v93).Cardinality()))).Uint32()).(bool), false)
					var _603_v95 _dafny.MultiSet
					_ = _603_v95
					_603_v95 = _602_v94
					var _604_v96 _dafny.Array
					_ = _604_v96
					var _nwElement0_23 _dafny.MultiSet = _603_v95
					_ = _nwElement0_23
					var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
					_ = _nw108
					(_nw108).ArraySet1(_nwElement0_23, 0)
					(_nw108).ArraySet1(_602_v94, 1)
					(_nw108).ArraySet1(_603_v95, 2)
					(_nw108).ArraySet1(Companion_Default___.Fm29(p0, p0, p2, p0, globalState), 3)
					(_nw108).ArraySet1((func() _dafny.MultiSet {
						if !(p2) {
							return _603_v95
						}
						return _603_v95
					})(), 4)
					(_nw108).ArraySet1(_603_v95, 5)
					(_nw108).ArraySet1(_603_v95, 6)
					(_nw108).ArraySet1(_603_v95, 7)
					(_nw108).ArraySet1(_603_v95, 8)
					(_nw108).ArraySet1(_602_v94, 9)
					(_nw108).ArraySet1(_603_v95, 10)
					(_nw108).ArraySet1(_603_v95, 11)
					(_nw108).ArraySet1(_dafny.MultiSetOf(p2), 12)
					_604_v96 = _nw108
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_604_v96), 0))
					_ = _index120
					(_604_v96).ArraySet1(_603_v95, (_index120).Int())
					var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_604_v96), 0))
					_ = _index121
					(_604_v96).ArraySet1(_603_v95, (_index121).Int())
					var _605_v97 _dafny.Sequence
					_ = _605_v97
					_605_v97 = _dafny.SeqOf(p1)
					var _606_v98 _dafny.Array
					_ = _606_v98
					var _nwElement0_24 bool = p2
					_ = _nwElement0_24
					var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(13))
					_ = _nw109
					(_nw109).ArraySet1(_nwElement0_24, 0)
					(_nw109).ArraySet1(p2, 1)
					(_nw109).ArraySet1(p2, 2)
					(_nw109).ArraySet1(p2, 3)
					(_nw109).ArraySet1((_601_v93).Select((Companion_Default___.SafeIndex((_548_v50).Cardinality(), _dafny.IntOfUint32((_601_v93).Cardinality()))).Uint32()).(bool), 4)
					(_nw109).ArraySet1(p2, 5)
					(_nw109).ArraySet1((p1).IsSubsetOf((_605_v97).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_605_v97).Cardinality()))).Uint32()).(_dafny.MultiSet)), 6)
					(_nw109).ArraySet1(p2, 7)
					(_nw109).ArraySet1(p2, 8)
					(_nw109).ArraySet1(p2, 9)
					(_nw109).ArraySet1(p2, 10)
					(_nw109).ArraySet1(p2, 11)
					(_nw109).ArraySet1(true, 12)
					_606_v98 = _nw109
					var _rhs103 bool = p2
					_ = _rhs103
					var _rhs104 _dafny.Sequence = _549_v51
					_ = _rhs104
					var _rhs105 _dafny.Array = _606_v98
					_ = _rhs105
					var _lhs91 *GlobalState = globalState
					_ = _lhs91
					var _lhs92 *GlobalState = globalState
					_ = _lhs92
					_lhs91.F4 = _rhs103
					_lhs92.F0 = _rhs104
					_606_v98 = _rhs105
					var _607_v99 _dafny.Map
					_ = _607_v99
					_607_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _549_v51)
					var _608_v100 _dafny.Map
					_ = _608_v100
					_608_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (_607_v99).Contains(p2) {
							return (_607_v99).Get(p2).(_dafny.Sequence)
						}
						return _549_v51
					})(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_607_v99).Contains(p2) {
							return (_607_v99).Get(p2).(_dafny.Sequence)
						}
						return _549_v51
					})()).Cardinality()))).Uint32(), _dafny.IntOfInt64(288)))
					(globalState).F0 = (func() _dafny.Sequence {
						if (_608_v100).Contains(p0) {
							return (_608_v100).Get(p0).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(622))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}((func(_609_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_610_i16 _dafny.Int) _dafny.Int {
								return _609_p0
							}
						})(p0)))
					})()
					(globalState).F4 = (func() bool {
						if p2 {
							return p2
						}
						return p2
					})()
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _612_v101 _dafny.Array
		_ = _612_v101
		var _nwElement0_25 _dafny.Int = (func() _dafny.Int {
			if p2 {
				return (_dafny.Zero).Minus(p0)
			}
			return p0
		})()
		_ = _nwElement0_25
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(5))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_25, 0)
		(_nw110).ArraySet1(p0, 1)
		(_nw110).ArraySet1((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Update(p2, p2)).Cardinality()).Plus(_dafny.IntOfInt64(720)), 2)
		(_nw110).ArraySet1(p0, 3)
		(_nw110).ArraySet1(p0, 4)
		_612_v101 = _nw110
		_612_v101 = _612_v101
	}
}
func (_this *C1) M14(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F8 = (p3).Minus(p3)
		var _hi3 _dafny.Int = p3
		_ = _hi3
		for _613_i0 := p3; _613_i0.Cmp(_hi3) < 0; _613_i0 = _613_i0.Plus(_dafny.One) {
			var _614_v0 *C0
			_ = _614_v0
			var _nw111 *C0 = New_C0_()
			_ = _nw111
			_nw111.Ctor__()
			_614_v0 = _nw111
			(globalState).F4 = p0
			(globalState).F4 = p2
			(globalState).F3 = (_dafny.Zero).Minus(_613_i0)
		}
		if p1 {
			var _615_v1 _dafny.Array
			_ = _615_v1
			var _len0_20 _dafny.Int = _dafny.One
			_ = _len0_20
			var _nw112 _dafny.Array
			_ = _nw112
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw112 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Int = (func(_616_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_617_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_617_i1, _616_p3)
					}
				})(p3)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw112 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw112).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw112).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_615_v1 = _nw112
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
			_ = _index122
			(_615_v1).ArraySet1(_dafny.IntOfInt64(199), (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_615_v1), 0))
			_ = _index123
			(_615_v1).ArraySet1(p3, (_index123).Int())
			var _618_v2 _dafny.Set
			_ = _618_v2
			_618_v2 = _dafny.SetOf(p3)
			var _619_v3 _dafny.Sequence
			_ = _619_v3
			_619_v3 = _dafny.SeqOf(_618_v2, _618_v2)
			var _620_v4 _dafny.Sequence
			_ = _620_v4
			_620_v4 = _dafny.SeqOf(p0, !(p1))
			var _621_v5 _dafny.Sequence
			_ = _621_v5
			_621_v5 = _dafny.UnicodeSeqOfUtf8Bytes("auprui")
			var _622_v6 _dafny.Map
			_ = _622_v6
			_622_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_621_v5).Cardinality()), p3)
			var _623_v7 D3
			_ = _623_v7
			_623_v7 = Companion_D3_.Create_DC8_(_622_v6)
			var _624_v8 _dafny.Map
			_ = _624_v8
			_624_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_620_v4, _623_v7)
			var _625_v9 D4
			_ = _625_v9
			_625_v9 = Companion_D4_.Create_DC10_(_620_v4)
			var _626_v10 D4
			_ = _626_v10
			_626_v10 = Companion_D4_.Create_DC12_(_625_v9)
			var _627_v11 _dafny.Map
			_ = _627_v11
			_627_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_626_v10, _618_v2)
			var _628_v12 D9
			_ = _628_v12
			_628_v12 = Companion_D9_.Create_DC21_(p1, _620_v4, !(p0))
			var _629_v13 _dafny.Sequence
			_ = _629_v13
			_629_v13 = _dafny.SeqOf(p3, p3)
			var _630_v14 _dafny.MultiSet
			_ = _630_v14
			_630_v14 = _dafny.MultiSetOf(p3, p3, p3, p3, p3)
			var _631_v15 _dafny.Sequence
			_ = _631_v15
			_631_v15 = _dafny.SeqOf(Companion_Default___.Fm31(p1, p1, p1, globalState))
			var _pat_let_tv20 = _621_v5
			_ = _pat_let_tv20
			var _pat_let_tv21 = _628_v12
			_ = _pat_let_tv21
			var _pat_let_tv22 = p1
			_ = _pat_let_tv22
			var _pat_let_tv23 = _621_v5
			_ = _pat_let_tv23
			var _pat_let_tv24 = globalState
			_ = _pat_let_tv24
			var _pat_let_tv25 = _621_v5
			_ = _pat_let_tv25
			var _pat_let_tv26 = _628_v12
			_ = _pat_let_tv26
			var _pat_let_tv27 = p1
			_ = _pat_let_tv27
			var _pat_let_tv28 = _621_v5
			_ = _pat_let_tv28
			var _pat_let_tv29 = globalState
			_ = _pat_let_tv29
			var _632_v16 _dafny.Array
			_ = _632_v16
			var _nwElement0_26 bool = ((_619_v3).Select((Companion_Default___.SafeIndex((_624_v8).Cardinality(), _dafny.IntOfUint32((_619_v3).Cardinality()))).Uint32()).(_dafny.Set)).IsSubsetOf((func() _dafny.Set {
				if (_627_v11).Contains(func(_pat_let17_0 D4) D4 {
					return func(_633_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let18_0 D4) D4 {
							return func(_634_dt__update_hcf18_h0 D4) D4 {
								return Companion_D4_.Create_DC12_(_634_dt__update_hcf18_h0)
							}(_pat_let18_0)
						}(Companion_Default___.Fm30(_dafny.IntOfUint32((_pat_let_tv20).Cardinality()), _pat_let_tv21, _pat_let_tv22, _dafny.IntOfUint32((_pat_let_tv23).Cardinality()), _pat_let_tv24))
					}(_pat_let17_0)
				}(Companion_D4_.Create_DC12_(_625_v9))) {
					return (_627_v11).Get(func(_pat_let19_0 D4) D4 {
						return func(_635_dt__update__tmp_h1 D4) D4 {
							return func(_pat_let20_0 D4) D4 {
								return func(_636_dt__update_hcf18_h1 D4) D4 {
									return Companion_D4_.Create_DC12_(_636_dt__update_hcf18_h1)
								}(_pat_let20_0)
							}(Companion_Default___.Fm30(_dafny.IntOfUint32((_pat_let_tv25).Cardinality()), _pat_let_tv26, _pat_let_tv27, _dafny.IntOfUint32((_pat_let_tv28).Cardinality()), _pat_let_tv29))
						}(_pat_let19_0)
					}(Companion_D4_.Create_DC12_(_625_v9))).(_dafny.Set)
				}
				return _618_v2
			})())
			_ = _nwElement0_26
			var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(29))
			_ = _nw113
			(_nw113).ArraySet1(_nwElement0_26, 0)
			(_nw113).ArraySet1(p0, 1)
			(_nw113).ArraySet1(p0, 2)
			(_nw113).ArraySet1(p0, 3)
			(_nw113).ArraySet1(p0, 4)
			(_nw113).ArraySet1(p2, 5)
			(_nw113).ArraySet1(!(p1), 6)
			(_nw113).ArraySet1(p0, 7)
			(_nw113).ArraySet1(p1, 8)
			(_nw113).ArraySet1(p2, 9)
			(_nw113).ArraySet1(!(p1), 10)
			(_nw113).ArraySet1(!(false), 11)
			(_nw113).ArraySet1((_this).Fm3(_629_v13, _630_v14, (_this).Fm3(_629_v13, _630_v14, p2, globalState), globalState), 12)
			(_nw113).ArraySet1(p1, 13)
			(_nw113).ArraySet1(p1, 14)
			(_nw113).ArraySet1(p2, 15)
			(_nw113).ArraySet1((p1) == (p1), 16)
			(_nw113).ArraySet1(!(p1), 17)
			(_nw113).ArraySet1(p1, 18)
			(_nw113).ArraySet1((p3).Cmp(_dafny.IntOfUint32((_631_v15).Cardinality())) == 0, 19)
			(_nw113).ArraySet1(p1, 20)
			(_nw113).ArraySet1(p2, 21)
			(_nw113).ArraySet1(p1, 22)
			(_nw113).ArraySet1((p1) && (p1), 23)
			(_nw113).ArraySet1(p0, 24)
			(_nw113).ArraySet1(!(p2), 25)
			(_nw113).ArraySet1(p0, 26)
			(_nw113).ArraySet1(_dafny.Companion_Sequence_.Contains(_629_v13, p3), 27)
			(_nw113).ArraySet1(p2, 28)
			_632_v16 = _nw113
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))
			_ = _index124
			(_632_v16).ArraySet1((_620_v4).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_620_v4).Cardinality()))).Uint32()).(bool), (_index124).Int())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
			_ = _index125
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_615_v1), 0))
			_ = _index126
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))
			_ = _index127
			var _rhs106 _dafny.Int = p3
			_ = _rhs106
			var _rhs107 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.SetOf(p3, p3)).Cardinality(), p3)
			_ = _rhs107
			var _rhs108 _dafny.Int = p3
			_ = _rhs108
			var _rhs109 bool = !(p2)
			_ = _rhs109
			var _rhs110 bool = p1
			_ = _rhs110
			var _lhs93 *GlobalState = globalState
			_ = _lhs93
			var _lhs94 _dafny.Array = _615_v1
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
			_ = _lhs95
			var _lhs96 _dafny.Array = _615_v1
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_615_v1), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _632_v16
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))
			_ = _lhs99
			var _lhs100 *GlobalState = globalState
			_ = _lhs100
			_lhs93.F8 = _rhs106
			(_lhs94).ArraySet1(_rhs107, (_lhs95).Int())
			(_lhs96).ArraySet1(_rhs108, (_lhs97).Int())
			(_lhs98).ArraySet1(_rhs109, (_lhs99).Int())
			_lhs100.F4 = _rhs110
			var _637_v17 _dafny.Sequence
			_ = _637_v17
			_637_v17 = _dafny.SeqOf(_621_v5, _621_v5, _621_v5, _621_v5, _621_v5)
			var _638_v18 _dafny.CodePoint
			_ = _638_v18
			_638_v18 = _dafny.CodePoint('y')
			var _639_v19 _dafny.Map
			_ = _639_v19
			_639_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int)))
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))
			_ = _index128
			(_632_v16).ArraySet1(!(Companion_Default___.Fm32(p3, _637_v17, _638_v18, p0, globalState)).Equals(_639_v19), (_index128).Int())
			var _640_v20 _dafny.Set
			_ = _640_v20
			_640_v20 = _dafny.SetOf(p2)
			if ((_640_v20).IsProperSubsetOf(_640_v20)) || ((_640_v20).Contains(p2)) {
				_621_v5 = _dafny.UnicodeSeqOfUtf8Bytes("mosxmc")
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
				_ = _index129
				(_615_v1).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p3, p3)), _dafny.IntOfInt64(837))), (_index129).Int())
				var _641_v21 _dafny.Map
				_ = _641_v21
				_641_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_632_v16, _638_v18)
				(globalState).F8 = (_641_v21).Cardinality()
				var _642_v22 _dafny.Map
				_ = _642_v22
				_642_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_637_v17).Cardinality()), p1)
				var _643_v25 _dafny.Set
				_ = _643_v25
				_643_v25 = _dafny.SetOf(_618_v2, func() _dafny.Set {
					var _coll64 = _dafny.NewBuilder()
					_ = _coll64
					for _iter65 := _dafny.Iterate((_642_v22).Keys().Elements()); ; {
						_compr_64, _ok65 := _iter65()
						if !_ok65 {
							break
						}
						var _644_v23 _dafny.Int
						_644_v23 = interface{}(_compr_64).(_dafny.Int)
						if (_642_v22).Contains(_644_v23) {
							_coll64.Add(Companion_Default___.SafeModuloInt(_644_v23, (_dafny.SetOf(false, !(false), true, false, !(true))).Cardinality()))
						}
					}
					return _coll64.ToSet()
				}(), func() _dafny.Set {
					var _coll65 = _dafny.NewBuilder()
					_ = _coll65
					for _iter66 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(451))); ; {
						_compr_65, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _645_v24 _dafny.Int
						_645_v24 = interface{}(_compr_65).(_dafny.Int)
						if ((_dafny.IntOfInt64(280)).Cmp(_645_v24) <= 0) && ((_645_v24).Cmp(_dafny.IntOfInt64(451)) < 0) {
							_coll65.Add((_645_v24).Times(p3))
						}
					}
					return _coll65.ToSet()
				}())
				_643_v25 = _643_v25
				var _646_v26 _dafny.Array
				_ = _646_v26
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
				_ = _nw114
				_646_v26 = _nw114
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_646_v26), 0))
				_ = _index130
				(_646_v26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(_629_v13, _630_v14, !((_632_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))).Int()).(bool)), globalState), false), (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_646_v26), 0))
				_ = _index131
				(_646_v26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_622_v6).Contains(_dafny.IntOfInt64(340)), p2), (_index131).Int())
			} else {
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
				_ = _index132
				(_615_v1).ArraySet1(p3, (_index132).Int())
				var _647_v27 _dafny.MultiSet
				_ = _647_v27
				_647_v27 = _dafny.MultiSetOf(_620_v4)
				var _648_v28 _dafny.Sequence
				_ = _648_v28
				_648_v28 = _dafny.SeqOf(_647_v27, _647_v27, _647_v27)
				var _649_v29 _dafny.Map
				_ = _649_v29
				_649_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm26(_623_v7, _dafny.IntOfInt64(390), (_648_v28).Select((Companion_Default___.SafeIndex((_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_648_v28).Cardinality()))).Uint32()).(_dafny.MultiSet), globalState), p1)
				(globalState).F3 = ((_dafny.Zero).Minus((_649_v29).Cardinality())).Minus(_dafny.IntOfInt64(126))
				var _650_v30 _dafny.Array
				_ = _650_v30
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw115
				_650_v30 = _nw115
				var _651_v31 _dafny.Map
				_ = _651_v31
				_651_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, (_632_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_632_v16), 0))).Int()).(bool)), p1)
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_650_v30), 0))
				_ = _index133
				(_650_v30).ArraySet1(_651_v31, (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_650_v30), 0))
				_ = _index134
				var _rhs111 _dafny.Int = _dafny.IntOfInt64(393)
				_ = _rhs111
				var _rhs112 _dafny.Set = (_618_v2).Difference(_dafny.SetOf(_dafny.IntOfUint32((_621_v5).Cardinality())))
				_ = _rhs112
				var _rhs113 _dafny.Map = (Companion_Default___.Fm33(p1, p0, globalState)).Merge((func() _dafny.Map {
					var _coll66 = _dafny.NewMapBuilder()
					_ = _coll66
					for _iter67 := _dafny.Iterate((_631_v15).Elements()); ; {
						_compr_66, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _652_v32 _dafny.Sequence
						_652_v32 = interface{}(_compr_66).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_631_v15, _652_v32) {
							_coll66.Add(_652_v32, p1)
						}
					}
					return _coll66.ToMap()
				}()).Merge(_651_v31))
				_ = _rhs113
				var _rhs114 _dafny.CodePoint = _638_v18
				_ = _rhs114
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				var _lhs102 _dafny.Array = _650_v30
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_650_v30), 0))
				_ = _lhs103
				_lhs101.F3 = _rhs111
				_618_v2 = _rhs112
				(_lhs102).ArraySet1(_rhs113, (_lhs103).Int())
				_638_v18 = _rhs114
				var _653_v33 _dafny.Map
				_ = _653_v33
				_653_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_632_v16, (_dafny.Zero).Minus((_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int)))
				_653_v33 = (_653_v33).Update(_632_v16, Companion_Default___.SafeModuloInt((_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int), p3))
				var _654_v34 *C0
				_ = _654_v34
				var _nw116 *C0 = New_C0_()
				_ = _nw116
				_nw116.Ctor__()
				_654_v34 = _nw116
			}
			var _655_v35 _dafny.Map
			_ = _655_v35
			_655_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_640_v20).Cardinality(), _638_v18)
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))
			_ = _index135
			(_615_v1).ArraySet1(Companion_Default___.SafeDivisionInt((_655_v35).Cardinality(), (_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int)), (_index135).Int())
			var _656_v37 _dafny.MultiSet
			_ = _656_v37
			_656_v37 = _dafny.MultiSetOf(p3, _dafny.IntOfUint32((_621_v5).Cardinality()))
			var _657_v38 _dafny.MultiSet
			_ = _657_v38
			_657_v38 = _dafny.MultiSetOf(_620_v4)
			var _pat_let_tv30 = p2
			_ = _pat_let_tv30
			var _pat_let_tv31 = globalState
			_ = _pat_let_tv31
			var _pat_let_tv32 = p2
			_ = _pat_let_tv32
			var _pat_let_tv33 = globalState
			_ = _pat_let_tv33
			var _658_v39 _dafny.Map
			_ = _658_v39
			_658_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm26(_623_v7, (func() _dafny.Map {
				var _coll67 = _dafny.NewMapBuilder()
				_ = _coll67
				for _iter68 := _dafny.Iterate((_dafny.SeqOf(_656_v37, _630_v14, func(_pat_let21_0 _dafny.MultiSet) _dafny.MultiSet {
					return func(_659_dt__update__tmp_h2 _dafny.MultiSet) _dafny.MultiSet {
						return func(_pat_let22_0 _dafny.MultiSet) _dafny.MultiSet {
							return func(_660_dt__update_hcf22_h0 _dafny.MultiSet) _dafny.MultiSet {
								return _660_dt__update_hcf22_h0
							}(_pat_let22_0)
						}(Companion_Default___.Fm27(_pat_let_tv30, _pat_let_tv31))
					}(_pat_let21_0)
				}(_656_v37))).Elements()); ; {
					_compr_67, _ok68 := _iter68()
					if !_ok68 {
						break
					}
					var _661_v36 _dafny.MultiSet
					_661_v36 = interface{}(_compr_67).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_656_v37, _630_v14, func(_pat_let23_0 _dafny.MultiSet) _dafny.MultiSet {
						return func(_662_dt__update__tmp_h2 _dafny.MultiSet) _dafny.MultiSet {
							return func(_pat_let24_0 _dafny.MultiSet) _dafny.MultiSet {
								return func(_663_dt__update_hcf22_h0 _dafny.MultiSet) _dafny.MultiSet {
									return _663_dt__update_hcf22_h0
								}(_pat_let24_0)
							}(Companion_Default___.Fm27(_pat_let_tv32, _pat_let_tv33))
						}(_pat_let23_0)
					}(_656_v37)), _661_v36) {
						_coll67.Add(_661_v36, true)
					}
				}
				return _coll67.ToMap()
			}()).Cardinality(), _657_v38, globalState))
			_658_v39 = (_658_v39).Update(p2, Companion_Default___.SafeModuloInt(p3, (_615_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_615_v1), 0))).Int()).(_dafny.Int)))
		} else {
			var _664_v40 _dafny.Sequence
			_ = _664_v40
			_664_v40 = _dafny.SeqOf(p3)
			var _665_v41 _dafny.MultiSet
			_ = _665_v41
			_665_v41 = _dafny.MultiSetOf(p3, _dafny.IntOfInt64(492), p3, p3, p3)
			r0 = (_this).Fm3(_664_v40, _665_v41, p1, globalState)
			var _666_v42 _dafny.Set
			_ = _666_v42
			_666_v42 = _dafny.SetOf(p0, p2, p1, (p1) && (p1))
			var _667_v43 _dafny.Array
			_ = _667_v43
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
			_ = _nw117
			_667_v43 = _nw117
			var _668_v44 _dafny.Map
			_ = _668_v44
			_668_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _rhs115 _dafny.Set = (_666_v42).Intersection(_666_v42)
			_ = _rhs115
			var _rhs116 _dafny.Array = _667_v43
			_ = _rhs116
			var _rhs117 _dafny.Int = (func() _dafny.Int {
				if (_665_v41).Contains(Companion_Default___.SafeModuloInt((_668_v44).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(p3)))) {
					return (_665_v41).Multiplicity(Companion_Default___.SafeModuloInt((_668_v44).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(p3))))
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)).Cardinality()
			})()
			_ = _rhs117
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			_666_v42 = _rhs115
			_667_v43 = _rhs116
			_lhs104.F3 = _rhs117
			(globalState).F8 = p3
			var _669_v45 _dafny.Array
			_ = _669_v45
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
			_ = _nw118
			_669_v45 = _nw118
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_669_v45), 0))
			_ = _index136
			(_669_v45).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index136).Int())
			var _670_v46 _dafny.CodePoint
			_ = _670_v46
			_670_v46 = _dafny.CodePoint('s')
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_669_v45), 0))
			_ = _index137
			(_669_v45).ArraySet1CodePoint(_670_v46, (_index137).Int())
			(globalState).F8 = p3
		}
		r0 = p0
		var _671_v47 _dafny.Sequence
		_ = _671_v47
		_671_v47 = _dafny.UnicodeSeqOfUtf8Bytes("mjj")
		var _672_v48 _dafny.Array
		_ = _672_v48
		var _nwElement0_27 bool = false
		_ = _nwElement0_27
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(13))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_27, 0)
		(_nw119).ArraySet1(p1, 1)
		(_nw119).ArraySet1(!(p0), 2)
		(_nw119).ArraySet1(p0, 3)
		(_nw119).ArraySet1(!_dafny.Companion_Sequence_.Contains(_671_v47, _dafny.CodePoint('x')), 4)
		(_nw119).ArraySet1(p1, 5)
		(_nw119).ArraySet1(p2, 6)
		(_nw119).ArraySet1((p3).Cmp((Companion_Default___.Fm34(p0, p1, globalState)).Dtor_cf0()) < 0, 7)
		(_nw119).ArraySet1(p2, 8)
		(_nw119).ArraySet1(p1, 9)
		(_nw119).ArraySet1(p0, 10)
		(_nw119).ArraySet1(p0, 11)
		(_nw119).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vvo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_673_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))), 12)
		_672_v48 = _nw119
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_672_v48), 0))
		_ = _index138
		(_672_v48).ArraySet1(p1, (_index138).Int())
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_672_v48), 0))
		_ = _index139
		(_672_v48).ArraySet1(p2, (_index139).Int())
		var _674_v49 _dafny.Sequence
		_ = _674_v49
		_674_v49 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_671_v47, _671_v47), (_672_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_672_v48), 0))).Int()).(bool), (_672_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_672_v48), 0))).Int()).(bool), p2)
		var _675_v50 _dafny.CodePoint
		_ = _675_v50
		_675_v50 = _dafny.CodePoint('f')
		var _676_v51 _dafny.MultiSet
		_ = _676_v51
		_676_v51 = _dafny.MultiSetOf((_dafny.Zero).Minus(p3))
		var _677_v52 _dafny.Sequence
		_ = _677_v52
		_677_v52 = _dafny.SeqOf((_676_v51).Cardinality())
		var _678_v53 D2
		_ = _678_v53
		_678_v53 = Companion_D2_.Create_DC7_(_675_v50, p1, _dafny.IntOfUint32((_677_v52).Cardinality()))
		var _679_v54 D12
		_ = _679_v54
		_679_v54 = Companion_D12_.Create_DC27_(p3, p3, p3, p3)
		var _pat_let_tv34 = _679_v54
		_ = _pat_let_tv34
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_672_v48), 0))
		_ = _index140
		(_672_v48).ArraySet1((_674_v49).Select((Companion_Default___.SafeIndex((func(_pat_let25_0 D2) D2 {
			return func(_682_dt__update__tmp_h4 D2) D2 {
				return func(_pat_let28_0 bool) D2 {
					return func(_683_dt__update_hcf10_h0 bool) D2 {
						return Companion_D2_.Create_DC7_((_682_dt__update__tmp_h4).Dtor_cf9(), _683_dt__update_hcf10_h0, (_682_dt__update__tmp_h4).Dtor_cf11())
					}(_pat_let28_0)
				}(true)
			}(_pat_let25_0)
		}(func(_pat_let26_0 D2) D2 {
			return func(_680_dt__update__tmp_h3 D2) D2 {
				return func(_pat_let27_0 _dafny.Int) D2 {
					return func(_681_dt__update_hcf11_h0 _dafny.Int) D2 {
						return Companion_D2_.Create_DC7_((_680_dt__update__tmp_h3).Dtor_cf9(), (_680_dt__update__tmp_h3).Dtor_cf10(), _681_dt__update_hcf11_h0)
					}(_pat_let27_0)
				}((_pat_let_tv34).Dtor_cf37())
			}(_pat_let26_0)
		}(_678_v53))).Dtor_cf11(), _dafny.IntOfUint32((_674_v49).Cardinality()))).Uint32()).(bool), (_index140).Int())
		r0 = p1
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
func (_this *C2) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			var _coll68 = _dafny.NewMapBuilder()
			_ = _coll68
			for _iter69 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(920), false)).Keys().Elements()); ; {
				_compr_68, _ok69 := _iter69()
				if !_ok69 {
					break
				}
				var _684_v0 _dafny.Int
				_684_v0 = interface{}(_compr_68).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(920), false)).Contains(_684_v0) {
					_coll68.Add((_684_v0).Minus(_dafny.IntOfInt64(851)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jf")).Cardinality()))
				}
			}
			return _coll68.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(441), _dafny.IntOfInt64(596)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-204), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ee")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(986)))
	}
}
func (_this *C2) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		var _source13 D4 = Companion_D4_.Create_DC11_(true, true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqonprou")).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("svtwbn"))
		_ = _source13
		if _source13.Is_DC11() {
			var _685___mcc_h0 bool = _source13.Get_().(D4_DC11).Cf14
			_ = _685___mcc_h0
			var _686___mcc_h1 bool = _source13.Get_().(D4_DC11).Cf15
			_ = _686___mcc_h1
			var _687___mcc_h2 _dafny.Int = _source13.Get_().(D4_DC11).Cf16
			_ = _687___mcc_h2
			var _688___mcc_h3 _dafny.Sequence = _source13.Get_().(D4_DC11).Cf17
			_ = _688___mcc_h3
			var _689_cf17 _dafny.Sequence = _688___mcc_h3
			_ = _689_cf17
			var _690_cf16 _dafny.Int = _687___mcc_h2
			_ = _690_cf16
			var _691_cf15 bool = _686___mcc_h1
			_ = _691_cf15
			var _692_cf14 bool = _685___mcc_h0
			_ = _692_cf14
			return false
		} else if _source13.Is_DC10() {
			var _693___mcc_h4 _dafny.Sequence = _source13.Get_().(D4_DC10).Cf13
			_ = _693___mcc_h4
			var _694_cf13 _dafny.Sequence = _693___mcc_h4
			_ = _694_cf13
			return (_dafny.IntOfInt64(378)).Cmp(_dafny.IntOfInt64(708)) != 0
		} else {
			var _695___mcc_h5 D4 = _source13.Get_().(D4_DC12).Cf18
			_ = _695___mcc_h5
			var _696_cf18 D4 = _695___mcc_h5
			_ = _696_cf18
			return false
		}
	}
}
func (_this *C2) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _697_v0 _dafny.Int
		_ = _697_v0
		_697_v0 = _dafny.IntOfInt64(558)
		(globalState).F3 = _697_v0
		var _698_v2 _dafny.MultiSet
		_ = _698_v2
		_698_v2 = _dafny.MultiSetOf(_697_v0)
		if (_dafny.MultiSetOf(_697_v0, _697_v0, _697_v0)).IsProperSubsetOf((_dafny.MultiSetOf((func() _dafny.Map {
			var _coll69 = _dafny.NewMapBuilder()
			_ = _coll69
			for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(764), _dafny.IntOfInt64(518))); ; {
				_compr_69, _ok70 := _iter70()
				if !_ok70 {
					break
				}
				var _699_v1 _dafny.Int
				_699_v1 = interface{}(_compr_69).(_dafny.Int)
				if ((_dafny.IntOfInt64(764)).Cmp(_699_v1) <= 0) && ((_699_v1).Cmp(_dafny.IntOfInt64(518)) < 0) {
					_coll69.Add((_699_v1).Minus(_697_v0), _dafny.CodePoint('u'))
				}
			}
			return _coll69.ToMap()
		}()).Cardinality())).Union(_698_v2)) {
			var _700_v3 _dafny.Sequence
			_ = _700_v3
			_700_v3 = _dafny.UnicodeSeqOfUtf8Bytes("xxhq")
			var _701_v4 _dafny.Sequence
			_ = _701_v4
			_701_v4 = _dafny.SeqOf(_700_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-395))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_702_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			})))
			var _703_v5 _dafny.Map
			_ = _703_v5
			_703_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ilkioe"), (_701_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.IntOfUint32((_701_v4).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _704_v6 bool
			_ = _704_v6
			_704_v6 = true
			_700_v3 = (func() _dafny.Sequence {
				if (_703_v5).Contains(_700_v3) {
					return (_703_v5).Get(_700_v3).(_dafny.Sequence)
				}
				return Companion_Default___.Fm17(Companion_Default___.Fm18(_704_v6, _704_v6, globalState), globalState)
			})()
			var _705_v7 _dafny.Sequence
			_ = _705_v7
			_705_v7 = _dafny.SeqOf(_697_v0, _697_v0)
			var _706_v8 _dafny.Sequence
			_ = _706_v8
			_706_v8 = _dafny.SeqOf((_this).Fm3(_705_v7, _698_v2, _704_v6, globalState))
			var _707_v9 D2
			_ = _707_v9
			_707_v9 = Companion_D2_.Create_DC6_((func() _dafny.Int {
				if (_706_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_700_v3).Cardinality()), _dafny.IntOfUint32((_706_v8).Cardinality()))).Uint32()).(bool) {
					return _697_v0
				}
				return _dafny.IntOfInt64(446)
			})(), _704_v6, _704_v6)
			_707_v9 = _707_v9
			var _708_v10 _dafny.Set
			_ = _708_v10
			_708_v10 = _dafny.SetOf(_704_v6)
			(globalState).F8 = Companion_Default___.SafeDivisionInt((_697_v0).Plus(_697_v0), (_708_v10).Cardinality())
			if _704_v6 {
				var _709_v11 *C0
				_ = _709_v11
				var _nw120 *C0 = New_C0_()
				_ = _nw120
				_nw120.Ctor__()
				_709_v11 = _nw120
				var _710_v12 *C0
				_ = _710_v12
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__()
				_710_v12 = _nw121
				var _711_v13 _dafny.Array
				_ = _711_v13
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw122
				_711_v13 = _nw122
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_711_v13), 0))
				_ = _index141
				(_711_v13).ArraySet1(_704_v6, (_index141).Int())
				var _712_v14 _dafny.CodePoint
				_ = _712_v14
				_712_v14 = _dafny.CodePoint('j')
				var _713_v15 _dafny.Array
				_ = _713_v15
				var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw123
				_713_v15 = _nw123
				var _714_v16 _dafny.Sequence
				_ = _714_v16
				_714_v16 = _dafny.SeqOf(_711_v13, _711_v13)
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_713_v15), 0))
				_ = _index142
				(_713_v15).ArraySet1(_714_v16, (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_711_v13), 0))
				_ = _index143
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_713_v15), 0))
				_ = _index144
				var _rhs118 _dafny.Int = _697_v0
				_ = _rhs118
				var _rhs119 _dafny.Int = _dafny.IntOfInt64(635)
				_ = _rhs119
				var _rhs120 bool = _704_v6
				_ = _rhs120
				var _rhs121 _dafny.CodePoint = _712_v14
				_ = _rhs121
				var _rhs122 _dafny.Sequence = _714_v16
				_ = _rhs122
				var _lhs105 *GlobalState = globalState
				_ = _lhs105
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				var _lhs107 _dafny.Array = _711_v13
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_711_v13), 0))
				_ = _lhs108
				var _lhs109 _dafny.Array = _713_v15
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_713_v15), 0))
				_ = _lhs110
				_lhs105.F8 = _rhs118
				_lhs106.F3 = _rhs119
				(_lhs107).ArraySet1(_rhs120, (_lhs108).Int())
				_712_v14 = _rhs121
				(_lhs109).ArraySet1(_rhs122, (_lhs110).Int())
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_711_v13), 0))
				_ = _index145
				(_711_v13).ArraySet1((_711_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_711_v13), 0))).Int()).(bool), (_index145).Int())
				var _715_v17 *C0
				_ = _715_v17
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__()
				_715_v17 = _nw124
			} else {
				var _716_v18 _dafny.CodePoint
				_ = _716_v18
				_716_v18 = _dafny.CodePoint('a')
				_700_v3 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_700_v3, (Companion_Default___.SafeIndex(_697_v0, _dafny.IntOfUint32((_700_v3).Cardinality()))).Uint32(), _716_v18), (Companion_Default___.SafeIndex(Companion_Default___.Fm20(!(false), globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_700_v3, (Companion_Default___.SafeIndex(_697_v0, _dafny.IntOfUint32((_700_v3).Cardinality()))).Uint32(), _716_v18)).Cardinality()))).Uint32(), _716_v18)
				var _717_v19 _dafny.Map
				_ = _717_v19
				_717_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _704_v6)
				var _718_v20 _dafny.Map
				_ = _718_v20
				_718_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v6, _704_v6))
				var _719_v21 D5
				_ = _719_v21
				_719_v21 = Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v6, true))
				var _720_v22 _dafny.Sequence
				_ = _720_v22
				_720_v22 = _dafny.SeqOf(_717_v19)
				var _721_v23 _dafny.Array
				_ = _721_v23
				var _nwElement0_28 _dafny.Map = _717_v19
				_ = _nwElement0_28
				var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(14))
				_ = _nw125
				(_nw125).ArraySet1(_nwElement0_28, 0)
				(_nw125).ArraySet1(_717_v19, 1)
				(_nw125).ArraySet1((func() _dafny.Map {
					if (_718_v20).Contains(_704_v6) {
						return (_718_v20).Get(_704_v6).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v6, _704_v6)
				})(), 2)
				(_nw125).ArraySet1(_717_v19, 3)
				(_nw125).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v6, _704_v6), 4)
				(_nw125).ArraySet1(_717_v19, 5)
				(_nw125).ArraySet1(_717_v19, 6)
				(_nw125).ArraySet1((_719_v21).Dtor_cf19(), 7)
				(_nw125).ArraySet1(_717_v19, 8)
				(_nw125).ArraySet1(Companion_Default___.Fm21(_704_v6, globalState), 9)
				(_nw125).ArraySet1(_717_v19, 10)
				(_nw125).ArraySet1((_720_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.IntOfUint32((_720_v22).Cardinality()))).Uint32()).(_dafny.Map), 11)
				(_nw125).ArraySet1(_717_v19, 12)
				(_nw125).ArraySet1(_717_v19, 13)
				_721_v23 = _nw125
				var _722_v24 _dafny.Map
				_ = _722_v24
				_722_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_721_v23, _704_v6)
				_722_v24 = (_722_v24).Update(_721_v23, (_697_v0).Cmp(_dafny.IntOfUint32((_700_v3).Cardinality())) == 0)
				_697_v0 = _697_v0
				(globalState).F3 = _697_v0
				_708_v10 = _708_v10
			}
			r0 = !(!((_dafny.IntOfUint32((_700_v3).Cardinality())).Cmp(_697_v0) <= 0))
		} else {
			var _723_v25 bool
			_ = _723_v25
			_723_v25 = false
			(globalState).F3 = (func() _dafny.Int {
				if _723_v25 {
					return _697_v0
				}
				return _697_v0
			})()
			var _724_v26 _dafny.Array
			_ = _724_v26
			var _nw126 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw126
			_724_v26 = _nw126
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))
			_ = _index146
			(_724_v26).ArraySet1(_723_v25, (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))
			_ = _index147
			(_724_v26).ArraySet1(_723_v25, (_index147).Int())
			var _725_v27 _dafny.Sequence
			_ = _725_v27
			_725_v27 = _dafny.UnicodeSeqOfUtf8Bytes("ijeb")
			var _726_v28 _dafny.Sequence
			_ = _726_v28
			_726_v28 = _dafny.SeqOf(_dafny.IntOfInt64(-273))
			var _727_v29 _dafny.Sequence
			_ = _727_v29
			_727_v29 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_725_v27).Cardinality()), _697_v0), (func() _dafny.Sequence {
				if (_724_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))).Int()).(bool) {
					return _726_v28
				}
				return _726_v28
			})())
			(globalState).F0 = (_727_v29).Select((Companion_Default___.SafeIndex(_697_v0, _dafny.IntOfUint32((_727_v29).Cardinality()))).Uint32()).(_dafny.Sequence)
			_697_v0 = _697_v0
			var _728_v30 _dafny.Array
			_ = _728_v30
			var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw127
			_728_v30 = _nw127
			var _729_v31 _dafny.Map
			_ = _729_v31
			_729_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_v0, _697_v0)
			var _730_v32 _dafny.Array
			_ = _730_v32
			var _nwElement0_29 _dafny.Int = _697_v0
			_ = _nwElement0_29
			var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(11))
			_ = _nw128
			(_nw128).ArraySet1(_nwElement0_29, 0)
			(_nw128).ArraySet1(_697_v0, 1)
			(_nw128).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_724_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))).Int()).(bool), _723_v25), _dafny.SeqOf((_724_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))).Int()).(bool)))).Cardinality()), 2)
			(_nw128).ArraySet1(Companion_Default___.SafeDivisionInt(_697_v0, Companion_Default___.Fm20((_724_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))).Int()).(bool), globalState)), 3)
			(_nw128).ArraySet1(Companion_Default___.SafeModuloInt(_697_v0, _697_v0), 4)
			(_nw128).ArraySet1(_697_v0, 5)
			(_nw128).ArraySet1(((_729_v31).Cardinality()).Plus(_dafny.IntOfInt64(523)), 6)
			(_nw128).ArraySet1(_697_v0, 7)
			(_nw128).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_697_v0, _697_v0)), 8)
			(_nw128).ArraySet1(Companion_Default___.SafeModuloInt(_697_v0, _697_v0), 9)
			(_nw128).ArraySet1(_697_v0, 10)
			_730_v32 = _nw128
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_730_v32), 0))
			_ = _index148
			(_730_v32).ArraySet1(Companion_Default___.Fm20((_724_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_724_v26), 0))).Int()).(bool), globalState), (_index148).Int())
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_730_v32), 0))
			_ = _index149
			var _rhs123 _dafny.Array = _728_v30
			_ = _rhs123
			var _rhs124 _dafny.Int = Companion_Default___.SafeDivisionInt(_697_v0, _697_v0)
			_ = _rhs124
			var _rhs125 _dafny.Int = _697_v0
			_ = _rhs125
			var _rhs126 bool = _723_v25
			_ = _rhs126
			var _rhs127 _dafny.Array = (func() _dafny.Array {
				if (_697_v0).Cmp(_697_v0) < 0 {
					return _724_v26
				}
				return _724_v26
			})()
			_ = _rhs127
			var _lhs111 *GlobalState = globalState
			_ = _lhs111
			var _lhs112 _dafny.Array = _730_v32
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_730_v32), 0))
			_ = _lhs113
			_728_v30 = _rhs123
			_lhs111.F3 = _rhs124
			(_lhs112).ArraySet1(_rhs125, (_lhs113).Int())
			_723_v25 = _rhs126
			_724_v26 = _rhs127
		}
		var _731_v33 bool
		_ = _731_v33
		_731_v33 = false
		var _732_v34 _dafny.MultiSet
		_ = _732_v34
		_732_v34 = _698_v2
		var _733_v35 _dafny.Map
		_ = _733_v35
		_733_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_v0, _dafny.MultiSetOf(_697_v0, _697_v0))
		var _734_v36 _dafny.Map
		_ = _734_v36
		_734_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _731_v33)
		var _735_v37 _dafny.Sequence
		_ = _735_v37
		_735_v37 = _dafny.UnicodeSeqOfUtf8Bytes("xebnstm")
		var _736_v38 D4
		_ = _736_v38
		_736_v38 = Companion_D4_.Create_DC11_(_731_v33, !(true), _697_v0, _735_v37)
		var _737_v39 _dafny.Sequence
		_ = _737_v39
		_737_v39 = _dafny.SeqOf(!(_731_v33))
		var _738_v40 _dafny.Array
		_ = _738_v40
		var _nwElement0_30 bool = !(_731_v33)
		_ = _nwElement0_30
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(25))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_30, 0)
		(_nw129).ArraySet1(_731_v33, 1)
		(_nw129).ArraySet1(_731_v33, 2)
		(_nw129).ArraySet1(_731_v33, 3)
		(_nw129).ArraySet1(_731_v33, 4)
		(_nw129).ArraySet1(!(true), 5)
		(_nw129).ArraySet1((_732_v34).IsDisjointFrom((func() _dafny.MultiSet {
			if (_733_v35).Contains(_dafny.IntOfInt64(679)) {
				return (_733_v35).Get(_dafny.IntOfInt64(679)).(_dafny.MultiSet)
			}
			return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-318))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}(func(_739_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_740_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))).Cardinality())
			})))).Update((_734_v36).Cardinality(), Companion_Default___.Abs(_697_v0))
		})()), 6)
		(_nw129).ArraySet1(false, 7)
		(_nw129).ArraySet1(_731_v33, 8)
		(_nw129).ArraySet1((_736_v38).Dtor_cf14(), 9)
		(_nw129).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_731_v33, _731_v33), _737_v39), 10)
		(_nw129).ArraySet1(_731_v33, 11)
		(_nw129).ArraySet1(_731_v33, 12)
		(_nw129).ArraySet1(_731_v33, 13)
		(_nw129).ArraySet1(_731_v33, 14)
		(_nw129).ArraySet1(_731_v33, 15)
		(_nw129).ArraySet1(_731_v33, 16)
		(_nw129).ArraySet1(!(_731_v33), 17)
		(_nw129).ArraySet1(_731_v33, 18)
		(_nw129).ArraySet1(_731_v33, 19)
		(_nw129).ArraySet1(_731_v33, 20)
		(_nw129).ArraySet1((_697_v0).Cmp(_697_v0) == 0, 21)
		(_nw129).ArraySet1(true, 22)
		(_nw129).ArraySet1(_731_v33, 23)
		(_nw129).ArraySet1(_731_v33, 24)
		_738_v40 = _nw129
		var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))
		_ = _index150
		(_738_v40).ArraySet1(false, (_index150).Int())
		var _741_v41 _dafny.Sequence
		_ = _741_v41
		_741_v41 = _dafny.SeqOf(_697_v0)
		var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))
		_ = _index151
		(_738_v40).ArraySet1(((_dafny.IntOfUint32((_741_v41).Cardinality())).Minus(_697_v0)).Cmp(_697_v0) >= 0, (_index151).Int())
		var _742_v42 _dafny.Array
		_ = _742_v42
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_21
		var _nw130 _dafny.Array
		_ = _nw130
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw130 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) _dafny.Int = (func(_743_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_744_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_744_i3, (_dafny.Zero).Minus(_743_v0))
				}
			})(_697_v0)
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw130 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw130).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw130).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_742_v42 = _nw130
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))
		_ = _index152
		(_742_v42).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), (_index152).Int())
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))
		_ = _index153
		(_742_v42).ArraySet1(_697_v0, (_index153).Int())
		var _745_i4 _dafny.Int
		_ = _745_i4
		_745_i4 = _dafny.Zero
		{
			for (_738_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))).Int()).(bool) {
				{
					if (_745_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_745_i4 = (_745_i4).Plus(_dafny.One)
					r0 = _731_v33
					(globalState).F8 = (_697_v0).Minus((_742_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))).Int()).(_dafny.Int))
					var _746_v43 _dafny.Map
					_ = _746_v43
					_746_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v42, _741_v41)
					var _747_v44 _dafny.MultiSet
					_ = _747_v44
					_747_v44 = _dafny.MultiSetOf((_738_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))).Int()).(bool), _731_v33, _731_v33)
					_741_v41 = (func() _dafny.Sequence {
						if (_746_v43).Contains(_742_v42) {
							return (_746_v43).Get(_742_v42).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_741_v41, _741_v41), (Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_747_v44).Contains((_738_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))).Int()).(bool)) {
								return (_747_v44).Multiplicity((_738_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_738_v40), 0))).Int()).(bool))
							}
							return _dafny.IntOfInt64(690)
						})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_741_v41, _741_v41)).Cardinality()))).Uint32(), (_742_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))).Int()).(_dafny.Int))
					})()
					var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))
					_ = _index154
					(_742_v42).ArraySet1(((_742_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))).Int()).(_dafny.Int)).Times((_742_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))).Int()).(_dafny.Int)), (_index154).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _748_v45 _dafny.CodePoint
		_ = _748_v45
		_748_v45 = _dafny.CodePoint('h')
		var _749_v46 _dafny.Map
		_ = _749_v46
		_749_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_v0, _742_v42)
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))
		_ = _index155
		var _rhs128 _dafny.Array = (func() _dafny.Array {
			if (_749_v46).Contains(_697_v0) {
				return (_749_v46).Get(_697_v0).(_dafny.Array)
			}
			return _742_v42
		})()
		_ = _rhs128
		var _rhs129 _dafny.CodePoint = _dafny.CodePoint('b')
		_ = _rhs129
		var _rhs130 _dafny.CodePoint = _748_v45
		_ = _rhs130
		var _rhs131 _dafny.Int = ((_742_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))).Int()).(_dafny.Int)).Times(_697_v0)
		_ = _rhs131
		var _lhs114 _dafny.Array = _742_v42
		_ = _lhs114
		var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_742_v42), 0))
		_ = _lhs115
		_742_v42 = _rhs128
		_748_v45 = _rhs129
		_748_v45 = _rhs130
		(_lhs114).ArraySet1(_rhs131, (_lhs115).Int())
		r0 = _731_v33
		var _pat_let_tv35 = _748_v45
		_ = _pat_let_tv35
		var _pat_let_tv36 = _731_v33
		_ = _pat_let_tv36
		var _pat_let_tv37 = _738_v40
		_ = _pat_let_tv37
		var _pat_let_tv38 = _738_v40
		_ = _pat_let_tv38
		r1 = func(_source14 D4) bool {
			if _source14.Is_DC11() {
				var _750___mcc_h0 bool = _source14.Get_().(D4_DC11).Cf14
				_ = _750___mcc_h0
				var _751___mcc_h1 bool = _source14.Get_().(D4_DC11).Cf15
				_ = _751___mcc_h1
				var _752___mcc_h2 _dafny.Int = _source14.Get_().(D4_DC11).Cf16
				_ = _752___mcc_h2
				var _753___mcc_h3 _dafny.Sequence = _source14.Get_().(D4_DC11).Cf17
				_ = _753___mcc_h3
				var _754_cf17 _dafny.Sequence = _753___mcc_h3
				_ = _754_cf17
				var _755_cf16 _dafny.Int = _752___mcc_h2
				_ = _755_cf16
				var _756_cf15 bool = _751___mcc_h1
				_ = _756_cf15
				var _757_cf14 bool = _750___mcc_h0
				_ = _757_cf14
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_754_cf17, _dafny.Companion_Sequence_.Update(_754_cf17, (Companion_Default___.SafeIndex(_755_cf16, _dafny.IntOfUint32((_754_cf17).Cardinality()))).Uint32(), _pat_let_tv35))
			} else if _source14.Is_DC10() {
				var _758___mcc_h4 _dafny.Sequence = _source14.Get_().(D4_DC10).Cf13
				_ = _758___mcc_h4
				var _759_cf13 _dafny.Sequence = _758___mcc_h4
				_ = _759_cf13
				return _pat_let_tv36
			} else {
				var _760___mcc_h5 D4 = _source14.Get_().(D4_DC12).Cf18
				_ = _760___mcc_h5
				var _761_cf18 D4 = _760___mcc_h5
				_ = _761_cf18
				return (_pat_let_tv38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_pat_let_tv37), 0))).Int()).(bool)
			}
		}(_736_v38)
		return r0, r1
	}
}
func (_this *C2) M11(p0 _dafny.Map, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _762_v0 bool
		_ = _762_v0
		_762_v0 = true
		(globalState).F4 = !(_762_v0)
		var _763_v1 _dafny.Int
		_ = _763_v1
		_763_v1 = _dafny.IntOfInt64(159)
		var _rhs132 _dafny.Int = _763_v1
		_ = _rhs132
		var _rhs133 _dafny.Int = _763_v1
		_ = _rhs133
		var _lhs116 *GlobalState = globalState
		_ = _lhs116
		_lhs116.F8 = _rhs132
		r2 = _rhs133
		if _762_v0 {
			var _764_v2 _dafny.Set
			_ = _764_v2
			_764_v2 = _dafny.SetOf(_762_v0)
			_764_v2 = (_764_v2).Difference((func() _dafny.Set {
				if !(_762_v0) {
					return _dafny.SetOf(_762_v0)
				}
				return _764_v2
			})())
			var _765_v3 _dafny.Sequence
			_ = _765_v3
			_765_v3 = _dafny.UnicodeSeqOfUtf8Bytes("cooouputi")
			var _766_v4 _dafny.MultiSet
			_ = _766_v4
			_766_v4 = _dafny.MultiSetOf(_763_v1)
			var _767_v5 _dafny.Map
			_ = _767_v5
			_767_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_762_v0, (_766_v4).Cardinality())
			var _768_v6 _dafny.MultiSet
			_ = _768_v6
			_768_v6 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ligquqw"), _765_v3)).Cardinality()), (_763_v1).Times((func() _dafny.Int {
				if (_766_v4).Contains(_763_v1) {
					return (_766_v4).Multiplicity(_763_v1)
				}
				return (_767_v5).Cardinality()
			})()), Companion_Default___.SafeDivisionInt(_763_v1, _763_v1))
			(globalState).F2 = _766_v4
			var _769_v7 _dafny.CodePoint
			_ = _769_v7
			_769_v7 = _dafny.CodePoint('n')
			var _770_v8 _dafny.Array
			_ = _770_v8
			var _nwElement0_31 _dafny.CodePoint = _769_v7
			_ = _nwElement0_31
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(9))
			_ = _nw131
			(_nw131).ArraySet1CodePoint(_nwElement0_31, 0)
			(_nw131).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _762_v0 {
					return _769_v7
				}
				return _769_v7
			})(), 1)
			(_nw131).ArraySet1CodePoint(_769_v7, 2)
			(_nw131).ArraySet1CodePoint(_769_v7, 3)
			(_nw131).ArraySet1CodePoint(_769_v7, 4)
			(_nw131).ArraySet1CodePoint(_769_v7, 5)
			(_nw131).ArraySet1CodePoint(Companion_Default___.Fm18(_762_v0, _762_v0, globalState), 6)
			(_nw131).ArraySet1CodePoint(_769_v7, 7)
			(_nw131).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _762_v0 {
					return _dafny.CodePoint('t')
				}
				return _769_v7
			})(), 8)
			_770_v8 = _nw131
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_770_v8), 0))
			_ = _index156
			(_770_v8).ArraySet1CodePoint(_dafny.CodePoint('o'), (_index156).Int())
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_770_v8), 0))
			_ = _index157
			(_770_v8).ArraySet1CodePoint(_769_v7, (_index157).Int())
			(globalState).F4 = (_763_v1).Cmp(_763_v1) == 0
			var _771_v9 _dafny.Array
			_ = _771_v9
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw132
			_771_v9 = _nw132
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_771_v9), 0))
			_ = _index158
			(_771_v9).ArraySet1(_763_v1, (_index158).Int())
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_771_v9), 0))
			_ = _index159
			(_771_v9).ArraySet1(Companion_Default___.SafeModuloInt(_763_v1, _dafny.IntOfInt64(217)), (_index159).Int())
		} else {
			var _772_v10 _dafny.Map
			_ = _772_v10
			_772_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(346), _763_v1)
			var _773_v11 D3
			_ = _773_v11
			_773_v11 = Companion_D3_.Create_DC8_(_772_v10)
			var _774_v12 _dafny.Sequence
			_ = _774_v12
			_774_v12 = _dafny.SeqOf((func() D3 {
				if _762_v0 {
					return _773_v11
				}
				return _773_v11
			})(), _773_v11)
			_774_v12 = _dafny.Companion_Sequence_.Concatenate(_774_v12, _774_v12)
			var _775_v13 _dafny.Array
			_ = _775_v13
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_22
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Sequence = (func(_776_v1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_777_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_776_v1, _776_v1)
					}
				})(_763_v1)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw133 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw133).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw133).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_775_v13 = _nw133
			var _778_v14 _dafny.Sequence
			_ = _778_v14
			_778_v14 = _dafny.SeqOf(_763_v1)
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))
			_ = _index160
			(_775_v13).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_778_v14, _778_v14), (_index160).Int())
			var _779_v15 _dafny.MultiSet
			_ = _779_v15
			_779_v15 = _dafny.MultiSetOf(_762_v0)
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))
			_ = _index161
			var _rhs134 _dafny.Sequence = _dafny.SeqOf(_763_v1, _763_v1, _763_v1, (func() _dafny.Int {
				if (_779_v15).Contains(_762_v0) {
					return (_779_v15).Multiplicity(_762_v0)
				}
				return _763_v1
			})())
			_ = _rhs134
			var _rhs135 bool = (_762_v0) && (_762_v0)
			_ = _rhs135
			var _lhs117 _dafny.Array = _775_v13
			_ = _lhs117
			var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))
			_ = _lhs118
			(_lhs117).ArraySet1(_rhs134, (_lhs118).Int())
			_762_v0 = _rhs135
			var _780_v16 _dafny.Map
			_ = _780_v16
			_780_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_762_v0, _763_v1)
			_780_v16 = (_780_v16).Update(_762_v0, ((p0).Merge(p0)).Cardinality())
			var _781_v17 _dafny.Map
			_ = _781_v17
			_781_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, (_this).Fm3((_775_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))).Int()).(_dafny.Sequence), (_dafny.MultiSetOf(_763_v1)).Update(_763_v1, Companion_Default___.Abs(_dafny.IntOfUint32(((_775_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))).Int()).(_dafny.Sequence)).Cardinality()))), _762_v0, globalState))
			var _782_v18 D7
			_ = _782_v18
			_782_v18 = Companion_D7_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, _762_v0))
			_781_v17 = (_782_v18).Dtor_cf23()
			var _783_v19 _dafny.Sequence
			_ = _783_v19
			_783_v19 = _dafny.UnicodeSeqOfUtf8Bytes("rd")
			var _784_v20 D2
			_ = _784_v20
			_784_v20 = Companion_D2_.Create_DC6_((_763_v1).Plus(_dafny.IntOfUint32((_783_v19).Cardinality())), !(_762_v0), !((_763_v1).Cmp(_763_v1) < 0))
			var _source15 D2 = _784_v20
			_ = _source15
			if _source15.Is_DC6() {
				var _785___mcc_h0 _dafny.Int = _source15.Get_().(D2_DC6).Cf6
				_ = _785___mcc_h0
				var _786___mcc_h1 bool = _source15.Get_().(D2_DC6).Cf7
				_ = _786___mcc_h1
				var _787___mcc_h2 bool = _source15.Get_().(D2_DC6).Cf8
				_ = _787___mcc_h2
				var _788_cf8 bool = _787___mcc_h2
				_ = _788_cf8
				var _789_cf7 bool = _786___mcc_h1
				_ = _789_cf7
				var _790_cf6 _dafny.Int = _785___mcc_h0
				_ = _790_cf6
				var _791_v21 _dafny.Sequence
				_ = _791_v21
				_791_v21 = _dafny.SeqOf(_762_v0, _762_v0, true, _762_v0, _788_cf8)
				var _792_v22 _dafny.MultiSet
				_ = _792_v22
				_792_v22 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_791_v21, _791_v21))
				_792_v22 = (_792_v22).Update(_791_v21, Companion_Default___.Abs(_763_v1))
				var _793_v23 _dafny.Array
				_ = _793_v23
				var _nw134 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(12))
				_ = _nw134
				_793_v23 = _nw134
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_793_v23), 0))
				_ = _index162
				(_793_v23).ArraySet1(Companion_D5_.Create_DC15_(), (_index162).Int())
				var _794_v24 D5
				_ = _794_v24
				_794_v24 = Companion_D5_.Create_DC15_()
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_793_v23), 0))
				_ = _index163
				(_793_v23).ArraySet1(_794_v24, (_index163).Int())
				var _795_v25 _dafny.Int
				_ = _795_v25
				var _796_v26 _dafny.Int
				_ = _796_v26
				var _797_v27 _dafny.Int
				_ = _797_v27
				var _798_v28 bool
				_ = _798_v28
				var _out17 _dafny.Int
				_ = _out17
				var _out18 _dafny.Int
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				var _out20 bool
				_ = _out20
				_out17, _out18, _out19, _out20 = (_this).M12(globalState)
				_795_v25 = _out17
				_796_v26 = _out18
				_797_v27 = _out19
				_798_v28 = _out20
				(globalState).F8 = _796_v26
			} else if _source15.Is_DC7() {
				var _799___mcc_h3 _dafny.CodePoint = _source15.Get_().(D2_DC7).Cf9
				_ = _799___mcc_h3
				var _800___mcc_h4 bool = _source15.Get_().(D2_DC7).Cf10
				_ = _800___mcc_h4
				var _801___mcc_h5 _dafny.Int = _source15.Get_().(D2_DC7).Cf11
				_ = _801___mcc_h5
				var _802_cf11 _dafny.Int = _801___mcc_h5
				_ = _802_cf11
				var _803_cf10 bool = _800___mcc_h4
				_ = _803_cf10
				var _804_cf9 _dafny.CodePoint = _799___mcc_h3
				_ = _804_cf9
				var _805_v29 _dafny.MultiSet
				_ = _805_v29
				_805_v29 = _dafny.MultiSetOf(_dafny.IntOfInt64(730))
				var _806_v30 _dafny.Map
				_ = _806_v30
				_806_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_805_v29).Cardinality(), _804_cf9)
				var _807_v31 _dafny.MultiSet
				_ = _807_v31
				_807_v31 = _dafny.MultiSetOf(_dafny.IntOfInt64(770), (_806_v30).Cardinality())
				var _808_v32 _dafny.Map
				_ = _808_v32
				_808_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_762_v0), _762_v0)
				var _809_v33 _dafny.Sequence
				_ = _809_v33
				_809_v33 = _dafny.SeqOf(_807_v31, Companion_Default___.Fm22(_763_v1, globalState))
				var _810_v34 _dafny.Sequence
				_ = _810_v34
				_810_v34 = _dafny.SeqOf(((_805_v29).Update(_dafny.IntOfInt64(510), Companion_Default___.Abs((_808_v32).Cardinality()))).IsProperSubsetOf((_809_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_783_v19).Cardinality()), _dafny.IntOfUint32((_809_v33).Cardinality()))).Uint32()).(_dafny.MultiSet)))
				var _811_v36 _dafny.Map
				_ = _811_v36
				_811_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, (func() _dafny.Map {
					var _coll70 = _dafny.NewMapBuilder()
					_ = _coll70
					for _iter71 := _dafny.Iterate((_772_v10).Keys().Elements()); ; {
						_compr_70, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _812_v35 _dafny.Int
						_812_v35 = interface{}(_compr_70).(_dafny.Int)
						if (_772_v10).Contains(_812_v35) {
							_coll70.Add((_812_v35).Minus(_763_v1), false)
						}
					}
					return _coll70.ToMap()
				}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_cf11, _803_cf10)).Update(_763_v1, false)))
				var _rhs136 _dafny.Sequence = Companion_Default___.Fm23(_783_v19, globalState)
				_ = _rhs136
				var _rhs137 bool = _762_v0
				_ = _rhs137
				var _rhs138 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_cf11, _781_v17)).Merge(_811_v36)
				_ = _rhs138
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				_810_v34 = _rhs136
				_lhs119.F4 = _rhs137
				_811_v36 = _rhs138
				_804_cf9 = _dafny.CodePoint('a')
				_783_v19 = Companion_Default___.Fm17(_804_cf9, globalState)
				var _pat_let_tv39 = _762_v0
				_ = _pat_let_tv39
				var _813_v37 _dafny.Map
				_ = _813_v37
				_813_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_cf11, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-851))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_814_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_815_i1 _dafny.Int) _dafny.Int {
						return _814_v1
					}
				})(_763_v1))), (Companion_Default___.SafeIndex((func(_pat_let29_0 D2) D2 {
					return func(_816_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let30_0 bool) D2 {
							return func(_817_dt__update_hcf7_h0 bool) D2 {
								return func(_pat_let31_0 _dafny.Int) D2 {
									return func(_818_dt__update_hcf6_h0 _dafny.Int) D2 {
										return Companion_D2_.Create_DC6_(_818_dt__update_hcf6_h0, _817_dt__update_hcf7_h0, (_816_dt__update__tmp_h0).Dtor_cf8())
									}(_pat_let31_0)
								}(_dafny.IntOfInt64(455))
							}(_pat_let30_0)
						}(_pat_let_tv39)
					}(_pat_let29_0)
				}(_784_v20)).Dtor_cf6(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-851))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_819_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_820_i1 _dafny.Int) _dafny.Int {
						return _819_v1
					}
				})(_763_v1)))).Cardinality()))).Uint32(), _763_v1), _778_v14))
				_813_v37 = (_813_v37).Update((_dafny.Zero).Minus((_763_v1).Plus(_763_v1)), (_775_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))).Int()).(_dafny.Sequence))
			} else {
				var _821___mcc_h6 _dafny.Array = _source15.Get_().(D2_DC5).Cf5
				_ = _821___mcc_h6
				var _822_cf5 _dafny.Array = _821___mcc_h6
				_ = _822_cf5
				var _823_v38 _dafny.MultiSet
				_ = _823_v38
				_823_v38 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-644))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_824_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality()))).Cardinality()))
				var _825_v39 _dafny.CodePoint
				_ = _825_v39
				_825_v39 = _dafny.CodePoint('h')
				var _826_v40 _dafny.Set
				_ = _826_v40
				_826_v40 = _dafny.SetOf(_dafny.CodePoint('y'), _825_v39)
				var _827_v41 _dafny.Map
				_ = _827_v41
				_827_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v38, ((_826_v40).Cardinality()).Plus(_763_v1))
				_827_v41 = (_827_v41).Update((Companion_Default___.Fm22(_763_v1, globalState)).Update(_763_v1, Companion_Default___.Abs(_763_v1)), _dafny.IntOfInt64(850))
				(globalState).F8 = (_763_v1).Plus(_763_v1)
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_822_cf5), 0))
				_ = _index164
				(_822_cf5).ArraySet1(_762_v0, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_822_cf5), 0))
				_ = _index165
				(_822_cf5).ArraySet1(true, (_index165).Int())
				var _rhs139 bool = !((_822_cf5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_822_cf5), 0))).Int()).(bool))
				_ = _rhs139
				var _rhs140 bool = _762_v0
				_ = _rhs140
				var _rhs141 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_775_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_775_v13), 0))).Int()).(_dafny.Sequence), _778_v14)
				_ = _rhs141
				var _rhs142 bool = _762_v0
				_ = _rhs142
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				r1 = _rhs139
				_lhs120.F4 = _rhs140
				_lhs121.F0 = _rhs141
				_762_v0 = _rhs142
			}
		}
		(globalState).F4 = true
		var _828_v42 _dafny.Array
		_ = _828_v42
		var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw135
		_828_v42 = _nw135
		var _829_v43 _dafny.Set
		_ = _829_v43
		_829_v43 = _dafny.SetOf(_828_v42, _828_v42, _828_v42, _828_v42)
		var _830_i3 _dafny.Int
		_ = _830_i3
		_830_i3 = _dafny.Zero
		{
			for !((_829_v43).Difference(_829_v43)).Contains(_828_v42) {
				{
					if (_830_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_830_i3 = (_830_i3).Plus(_dafny.One)
					(globalState).F4 = _762_v0
					var _831_v44 _dafny.Sequence
					_ = _831_v44
					_831_v44 = _dafny.SeqOf(_762_v0)
					var _832_v45 _dafny.Sequence
					_ = _832_v45
					_832_v45 = _dafny.SeqOf(_762_v0, _dafny.Companion_Sequence_.IsPrefixOf(_831_v44, _831_v44), _762_v0)
					_832_v45 = _dafny.Companion_Sequence_.Concatenate(_831_v44, _dafny.Companion_Sequence_.Update(_831_v44, (Companion_Default___.SafeIndex(_763_v1, _dafny.IntOfUint32((_831_v44).Cardinality()))).Uint32(), _762_v0))
					var _833_v46 _dafny.Array
					_ = _833_v46
					var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
					_ = _nw136
					_833_v46 = _nw136
					var _834_v47 _dafny.Sequence
					_ = _834_v47
					_834_v47 = _dafny.SeqOf(_763_v1, _dafny.IntOfUint32((_832_v45).Cardinality()))
					var _835_v48 _dafny.MultiSet
					_ = _835_v48
					_835_v48 = _dafny.MultiSetOf((_dafny.Zero).Minus(_763_v1))
					var _836_v49 _dafny.Array
					_ = _836_v49
					var _nwElement0_32 bool = !(_762_v0)
					_ = _nwElement0_32
					var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
					_ = _nw137
					(_nw137).ArraySet1(_nwElement0_32, 0)
					(_nw137).ArraySet1(_762_v0, 1)
					(_nw137).ArraySet1(_762_v0, 2)
					(_nw137).ArraySet1((_762_v0), 3)
					(_nw137).ArraySet1(false, 4)
					(_nw137).ArraySet1(_762_v0, 5)
					(_nw137).ArraySet1(!(_762_v0), 6)
					(_nw137).ArraySet1(_762_v0, 7)
					(_nw137).ArraySet1(!(_762_v0), 8)
					(_nw137).ArraySet1(_762_v0, 9)
					(_nw137).ArraySet1(true, 10)
					(_nw137).ArraySet1(_762_v0, 11)
					(_nw137).ArraySet1(_762_v0, 12)
					(_nw137).ArraySet1(false, 13)
					(_nw137).ArraySet1(!(_762_v0), 14)
					(_nw137).ArraySet1(_762_v0, 15)
					(_nw137).ArraySet1(_762_v0, 16)
					(_nw137).ArraySet1(!(true), 17)
					(_nw137).ArraySet1(_762_v0, 18)
					(_nw137).ArraySet1(_762_v0, 19)
					(_nw137).ArraySet1(_762_v0, 20)
					(_nw137).ArraySet1(_762_v0, 21)
					(_nw137).ArraySet1(_762_v0, 22)
					(_nw137).ArraySet1(_762_v0, 23)
					(_nw137).ArraySet1((_this).Fm3(_834_v47, _835_v48, _762_v0, globalState), 24)
					_836_v49 = _nw137
					var _837_v50 D2
					_ = _837_v50
					_837_v50 = Companion_D2_.Create_DC5_(_836_v49)
					var _838_v51 _dafny.MultiSet
					_ = _838_v51
					_838_v51 = _dafny.MultiSetOf(_762_v0, _762_v0)
					var _839_v52 _dafny.Map
					_ = _839_v52
					_839_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, _838_v51)
					var _840_v53 _dafny.Set
					_ = _840_v53
					_840_v53 = _dafny.SetOf(_763_v1)
					var _rhs143 _dafny.Array = _833_v46
					_ = _rhs143
					var _rhs144 bool = (_this).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_841_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_842_i4 _dafny.Int) _dafny.Int {
							return _841_v1
						}
					})(_763_v1))), _835_v48, (_this).Fm3(_dafny.Companion_Sequence_.Update(_834_v47, (Companion_Default___.SafeIndex(_763_v1, _dafny.IntOfUint32((_834_v47).Cardinality()))).Uint32(), _763_v1), _835_v48, false, globalState), globalState)
					_ = _rhs144
					var _rhs145 D2 = _837_v50
					_ = _rhs145
					var _rhs146 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(101)).Minus(_dafny.IntOfInt64(782)), _763_v1)
					_ = _rhs146
					var _rhs147 bool = (_this).Fm3(_dafny.Companion_Sequence_.Update(_834_v47, (Companion_Default___.SafeIndex((_839_v52).Cardinality(), _dafny.IntOfUint32((_834_v47).Cardinality()))).Uint32(), _763_v1), _dafny.MultiSetOf((_840_v53).Cardinality(), _763_v1), false, globalState)
					_ = _rhs147
					var _lhs122 *GlobalState = globalState
					_ = _lhs122
					_833_v46 = _rhs143
					_762_v0 = _rhs144
					_837_v50 = _rhs145
					_lhs122.F8 = _rhs146
					r1 = _rhs147
					if _762_v0 {
						(globalState).F8 = Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_762_v0), _762_v0)).Cardinality(), _dafny.IntOfInt64(945))
						var _843_v54 T0
						_ = _843_v54
						var _nw138 *C1 = New_C1_()
						_ = _nw138
						_nw138.Ctor__()
						_843_v54 = _nw138
						var _844_v55 _dafny.Sequence
						_ = _844_v55
						_844_v55 = _dafny.SeqOf(_843_v54)
						var _845_v56 _dafny.Sequence
						_ = _845_v56
						_845_v56 = _dafny.SeqOf(_844_v55)
						var _846_v57 _dafny.Sequence
						_ = _846_v57
						_846_v57 = _dafny.SeqOf((_845_v56).Select((Companion_Default___.SafeIndex(_763_v1, _dafny.IntOfUint32((_845_v56).Cardinality()))).Uint32()).(_dafny.Sequence))
						(globalState).F8 = (_763_v1).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_845_v56, _846_v57)).Cardinality()))))
						var _847_v58 _dafny.CodePoint
						_ = _847_v58
						_847_v58 = _dafny.CodePoint('t')
						var _848_v59 _dafny.MultiSet
						_ = _848_v59
						_848_v59 = _dafny.MultiSetOf(_847_v58)
						_848_v59 = ((_848_v59).Intersection(_848_v59)).Union(_848_v59)
						var _849_v60 *C1
						_ = _849_v60
						var _nw139 *C1 = New_C1_()
						_ = _nw139
						_nw139.Ctor__()
						_849_v60 = _nw139
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_828_v42), 0))
						_ = _index166
						(_828_v42).ArraySet1(_dafny.IntOfUint32((_832_v45).Cardinality()), (_index166).Int())
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_828_v42), 0))
						_ = _index167
						var _rhs148 _dafny.Int = _763_v1
						_ = _rhs148
						var _rhs149 _dafny.Array = _836_v49
						_ = _rhs149
						var _rhs150 bool = !(_762_v0)
						_ = _rhs150
						var _lhs123 _dafny.Array = _828_v42
						_ = _lhs123
						var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_828_v42), 0))
						_ = _lhs124
						(_lhs123).ArraySet1(_rhs148, (_lhs124).Int())
						_836_v49 = _rhs149
						_762_v0 = _rhs150
					} else {
						(globalState).F4 = (_dafny.SetOf(!(!((_this).Fm3(_dafny.SeqOf(_763_v1), _dafny.MultiSetFromSeq(_834_v47), _762_v0, globalState))), _762_v0)).IsProperSubsetOf(_dafny.SetOf(_762_v0))
						var _850_v61 *C1
						_ = _850_v61
						var _nw140 *C1 = New_C1_()
						_ = _nw140
						_nw140.Ctor__()
						_850_v61 = _nw140
						var _851_v62 _dafny.Map
						_ = _851_v62
						_851_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, _763_v1)
						_851_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v1, Companion_Default___.SafeModuloInt((_840_v53).Cardinality(), _763_v1))
						(globalState).F8 = (func() _dafny.Int {
							if (_835_v48).Contains(_763_v1) {
								return (_835_v48).Multiplicity(_763_v1)
							}
							return (func() _dafny.Int {
								if _762_v0 {
									return _763_v1
								}
								return (_851_v62).Cardinality()
							})()
						})()
						_851_v62 = (_851_v62).Merge(_851_v62)
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _852_v63 _dafny.Array
		_ = _852_v63
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_23
		var _nw141 _dafny.Array
		_ = _nw141
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw141 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) bool = (func(_853_v0 bool) func(_dafny.Int) bool {
				return func(_854_i5 _dafny.Int) bool {
					return _853_v0
				}
			})(_762_v0)
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw141 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw141).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw141).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_852_v63 = _nw141
		var _855_v64 _dafny.Sequence
		_ = _855_v64
		_855_v64 = _dafny.UnicodeSeqOfUtf8Bytes("wdjl")
		var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_852_v63), 0))
		_ = _index168
		(_852_v63).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_855_v64, _855_v64), (_index168).Int())
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_852_v63), 0))
		_ = _index169
		var _rhs151 bool = _762_v0
		_ = _rhs151
		var _rhs152 bool = _762_v0
		_ = _rhs152
		var _rhs153 bool = _762_v0
		_ = _rhs153
		var _rhs154 _dafny.Int = Companion_Default___.SafeModuloInt((_763_v1).Plus(_763_v1), _763_v1)
		_ = _rhs154
		var _rhs155 bool = _762_v0
		_ = _rhs155
		var _lhs125 _dafny.Array = _852_v63
		_ = _lhs125
		var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_852_v63), 0))
		_ = _lhs126
		_762_v0 = _rhs151
		_762_v0 = _rhs152
		r1 = _rhs153
		_763_v1 = _rhs154
		(_lhs125).ArraySet1(_rhs155, (_lhs126).Int())
		var _856_v65 _dafny.Map
		_ = _856_v65
		_856_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_762_v0, (_852_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_852_v63), 0))).Int()).(bool))
		r0 = (func() _dafny.Int {
			if (func() bool {
				if (_856_v65).Contains(_762_v0) {
					return (_856_v65).Get(_762_v0).(bool)
				}
				return _762_v0
			})() {
				return _dafny.IntOfInt64(704)
			}
			return _763_v1
		})()
		r1 = _762_v0
		r2 = _763_v1
		return r0, r1, r2
	}
}
func (_this *C2) M12(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _857_v0 _dafny.Int
		_ = _857_v0
		_857_v0 = _dafny.IntOfInt64(912)
		var _858_v1 bool
		_ = _858_v1
		_858_v1 = true
		(globalState).F4 = (_857_v0).Cmp(Companion_Default___.Fm20(_858_v1, globalState)) != 0
		var _859_v2 _dafny.Array
		_ = _859_v2
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_24
		var _nw142 _dafny.Array
		_ = _nw142
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw142 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) bool = (func(_860_v1 bool) func(_dafny.Int) bool {
				return func(_861_i0 _dafny.Int) bool {
					return _860_v1
				}
			})(_858_v1)
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw142 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw142).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw142).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_859_v2 = _nw142
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))
		_ = _index170
		(_859_v2).ArraySet1(_858_v1, (_index170).Int())
		var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))
		_ = _index171
		(_859_v2).ArraySet1(_858_v1, (_index171).Int())
		var _862_v3 _dafny.Sequence
		_ = _862_v3
		_862_v3 = _dafny.UnicodeSeqOfUtf8Bytes("pncslplhn")
		var _863_v4 _dafny.Map
		_ = _863_v4
		_863_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_862_v3, _857_v0)
		_863_v4 = (_863_v4).Update(_dafny.UnicodeSeqOfUtf8Bytes("momthrphd"), _dafny.IntOfInt64(92))
		var _864_v5 _dafny.Sequence
		_ = _864_v5
		_864_v5 = _dafny.SeqOf(_dafny.IntOfInt64(-139))
		(globalState).F0 = _dafny.Companion_Sequence_.Concatenate(_864_v5, _864_v5)
		var _865_v6 _dafny.Set
		_ = _865_v6
		_865_v6 = _dafny.SetOf(_862_v3)
		var _866_v7 D13
		_ = _866_v7
		_866_v7 = Companion_D13_.Create_DC29_(_865_v6)
		if (_865_v6).Equals(((_866_v7).Dtor_cf41()).Difference(_865_v6)) {
			(globalState).F4 = (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)
			var _nw143 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw143
			_859_v2 = _nw143
			var _867_v8 _dafny.Map
			_ = _867_v8
			_867_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_857_v0, _857_v0)
			var _868_v9 D3
			_ = _868_v9
			_868_v9 = Companion_D3_.Create_DC8_(_867_v8)
			var _869_v10 _dafny.Sequence
			_ = _869_v10
			_869_v10 = _dafny.SeqOf(false)
			var _870_v11 _dafny.MultiSet
			_ = _870_v11
			_870_v11 = _dafny.MultiSetOf(_869_v10, _869_v10, _869_v10, _869_v10)
			var _871_v12 _dafny.MultiSet
			_ = _871_v12
			_871_v12 = _dafny.MultiSetOf(Companion_Default___.Fm26(_868_v9, _857_v0, _870_v11, globalState))
			var _872_v13 _dafny.MultiSet
			_ = _872_v13
			_872_v13 = _871_v12
			(globalState).F4 = !((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)) || ((_872_v13).IsProperSubsetOf(_871_v12))
			var _873_v14 _dafny.MultiSet
			_ = _873_v14
			_873_v14 = _dafny.MultiSetOf(_862_v3)
			var _874_v15 _dafny.MultiSet
			_ = _874_v15
			_874_v15 = _dafny.MultiSetOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), _858_v1, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), _858_v1)
			var _875_v16 _dafny.Sequence
			_ = _875_v16
			_875_v16 = _dafny.SeqOf(_871_v12, _dafny.MultiSetOf((_873_v14).Cardinality(), (func() _dafny.Int {
				if (_874_v15).Contains(_858_v1) {
					return (_874_v15).Multiplicity(_858_v1)
				}
				return _857_v0
			})()))
			(globalState).F8 = _dafny.IntOfUint32((_875_v16).Cardinality())
			var _876_v17 _dafny.Array
			_ = _876_v17
			var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw144
			_876_v17 = _nw144
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_876_v17), 0))
			_ = _index172
			(_876_v17).ArraySet1((func() _dafny.Int {
				if _858_v1 {
					return _857_v0
				}
				return _857_v0
			})(), (_index172).Int())
			var _877_v18 _dafny.Set
			_ = _877_v18
			_877_v18 = _dafny.SetOf(_858_v1)
			var _878_v19 D9
			_ = _878_v19
			_878_v19 = Companion_D9_.Create_DC20_(_877_v18)
			var _879_v20 _dafny.Sequence
			_ = _879_v20
			_879_v20 = _dafny.SeqOf(_877_v18, _877_v18, _877_v18)
			var _pat_let_tv40 = _877_v18
			_ = _pat_let_tv40
			var _880_v21 _dafny.Array
			_ = _880_v21
			var _nwElement0_33 D9 = _878_v19
			_ = _nwElement0_33
			var _nw145 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(26))
			_ = _nw145
			(_nw145).ArraySet1(_nwElement0_33, 0)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_(_877_v18), 1)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_(_dafny.SetOf(_858_v1, _858_v1, _858_v1, _858_v1, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))), 2)
			(_nw145).ArraySet1(_878_v19, 3)
			(_nw145).ArraySet1(_878_v19, 4)
			(_nw145).ArraySet1(_878_v19, 5)
			(_nw145).ArraySet1(_878_v19, 6)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_(_877_v18), 7)
			(_nw145).ArraySet1(_878_v19, 8)
			(_nw145).ArraySet1(_878_v19, 9)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_(_877_v18), 10)
			(_nw145).ArraySet1(_878_v19, 11)
			(_nw145).ArraySet1(_878_v19, 12)
			(_nw145).ArraySet1(_878_v19, 13)
			(_nw145).ArraySet1(_878_v19, 14)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_(Companion_Default___.Fm35(_858_v1, _857_v0, _858_v1, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), globalState)), 15)
			(_nw145).ArraySet1(_878_v19, 16)
			(_nw145).ArraySet1(func(_pat_let32_0 D9) D9 {
				return func(_881_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let33_0 _dafny.Set) D9 {
						return func(_882_dt__update_hcf25_h0 _dafny.Set) D9 {
							return Companion_D9_.Create_DC20_(_882_dt__update_hcf25_h0)
						}(_pat_let33_0)
					}(_pat_let_tv40)
				}(_pat_let32_0)
			}(_878_v19), 17)
			(_nw145).ArraySet1(_878_v19, 18)
			(_nw145).ArraySet1(Companion_D9_.Create_DC20_((_879_v20).Select((Companion_Default___.SafeIndex(_857_v0, _dafny.IntOfUint32((_879_v20).Cardinality()))).Uint32()).(_dafny.Set)), 19)
			(_nw145).ArraySet1(_878_v19, 20)
			(_nw145).ArraySet1(_878_v19, 21)
			(_nw145).ArraySet1(_878_v19, 22)
			(_nw145).ArraySet1(_878_v19, 23)
			(_nw145).ArraySet1(_878_v19, 24)
			(_nw145).ArraySet1(_878_v19, 25)
			_880_v21 = _nw145
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_880_v21), 0))
			_ = _index173
			(_880_v21).ArraySet1(_878_v19, (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))
			_ = _index174
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_876_v17), 0))
			_ = _index175
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_880_v21), 0))
			_ = _index176
			var _rhs156 _dafny.Int = (Companion_Default___.Fm26(Companion_D3_.Create_DC8_(_867_v8), _857_v0, _dafny.MultiSetOf(_869_v10, _dafny.SeqOf(_858_v1), _869_v10), globalState)).Times((_857_v0).Minus(_dafny.IntOfInt64(859)))
			_ = _rhs156
			var _rhs157 bool = _858_v1
			_ = _rhs157
			var _rhs158 _dafny.Int = _857_v0
			_ = _rhs158
			var _rhs159 D9 = Companion_D9_.Create_DC20_(_877_v18)
			_ = _rhs159
			var _lhs127 *GlobalState = globalState
			_ = _lhs127
			var _lhs128 _dafny.Array = _859_v2
			_ = _lhs128
			var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))
			_ = _lhs129
			var _lhs130 _dafny.Array = _876_v17
			_ = _lhs130
			var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_876_v17), 0))
			_ = _lhs131
			var _lhs132 _dafny.Array = _880_v21
			_ = _lhs132
			var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_880_v21), 0))
			_ = _lhs133
			_lhs127.F8 = _rhs156
			(_lhs128).ArraySet1(_rhs157, (_lhs129).Int())
			(_lhs130).ArraySet1(_rhs158, (_lhs131).Int())
			(_lhs132).ArraySet1(_rhs159, (_lhs133).Int())
		} else {
			var _883_v22 _dafny.MultiSet
			_ = _883_v22
			_883_v22 = _dafny.MultiSetOf(_858_v1)
			var _884_v23 *C1
			_ = _884_v23
			var _nw146 *C1 = New_C1_()
			_ = _nw146
			_nw146.Ctor__()
			_884_v23 = _nw146
			var _885_v24 _dafny.Map
			_ = _885_v24
			_885_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_884_v23, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))
			var _886_v25 _dafny.Map
			_ = _886_v25
			_886_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_885_v24).Update(_884_v23, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)), _dafny.MultiSetOf(!((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)))
			var _rhs160 _dafny.Int = _dafny.IntOfUint32((_862_v3).Cardinality())
			_ = _rhs160
			var _rhs161 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _rhs161
			var _rhs162 bool = ((_883_v22).Difference((func() _dafny.MultiSet {
				if (_886_v25).Contains(_885_v24) {
					return (_886_v25).Get(_885_v24).(_dafny.MultiSet)
				}
				return _883_v22
			})())).IsSubsetOf(_883_v22)
			_ = _rhs162
			var _rhs163 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_862_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_887_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})))
			_ = _rhs163
			r0 = _rhs160
			r1 = _rhs161
			r3 = _rhs162
			_862_v3 = _rhs163
			(globalState).F8 = _857_v0
			var _888_v26 _dafny.CodePoint
			_ = _888_v26
			_888_v26 = _dafny.CodePoint('n')
			var _889_v27 _dafny.Map
			_ = _889_v27
			_889_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm17(_888_v26, globalState), _dafny.UnicodeSeqOfUtf8Bytes("isrckjn")), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))
			_889_v27 = (_889_v27).Update(_858_v1, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))
			var _890_v28 _dafny.Array
			_ = _890_v28
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
			_ = _nw147
			_890_v28 = _nw147
			var _891_v29 _dafny.Map
			_ = _891_v29
			_891_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), _857_v0)
			var _892_v30 _dafny.Set
			_ = _892_v30
			_892_v30 = _dafny.SetOf(_891_v29)
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_890_v28), 0))
			_ = _index177
			(_890_v28).ArraySet1(_892_v30, (_index177).Int())
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_890_v28), 0))
			_ = _index178
			(_890_v28).ArraySet1(((_892_v30).Union(_892_v30)).Difference(_892_v30), (_index178).Int())
			var _893_v31 _dafny.Map
			_ = _893_v31
			_893_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_857_v0, _858_v1)
			var _894_v32 D7
			_ = _894_v32
			_894_v32 = Companion_D7_.Create_DC17_((_893_v31).Update(_dafny.IntOfInt64(-103), _858_v1))
			var _source16 D7 = _894_v32
			_ = _source16
			if _source16.Is_DC18() {
				(globalState).F8 = Companion_Default___.SafeDivisionInt(_857_v0, _857_v0)
				var _895_v33 *C1
				_ = _895_v33
				var _nw148 *C1 = New_C1_()
				_ = _nw148
				_nw148.Ctor__()
				_895_v33 = _nw148
				var _896_v35 _dafny.Sequence
				_ = _896_v35
				_896_v35 = _dafny.SeqOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))
				var _897_v36 _dafny.MultiSet
				_ = _897_v36
				_897_v36 = _dafny.MultiSetOf(_896_v35)
				var _898_v37 _dafny.Set
				_ = _898_v37
				_898_v37 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}(func(_899_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()))
				var _900_v38 _dafny.Set
				_ = _900_v38
				_900_v38 = _dafny.SetOf((_898_v37).Cardinality(), _857_v0, _857_v0, (_865_v6).Cardinality(), _857_v0)
				var _901_v39 _dafny.MultiSet
				_ = _901_v39
				_901_v39 = _dafny.MultiSetOf(_857_v0)
				var _902_v40 _dafny.Array
				_ = _902_v40
				var _nwElement0_34 _dafny.Int = _857_v0
				_ = _nwElement0_34
				var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(23))
				_ = _nw149
				(_nw149).ArraySet1(_nwElement0_34, 0)
				(_nw149).ArraySet1((_dafny.Zero).Minus((_857_v0).Plus(_857_v0)), 1)
				(_nw149).ArraySet1(_857_v0, 2)
				(_nw149).ArraySet1(_857_v0, 3)
				(_nw149).ArraySet1(_857_v0, 4)
				(_nw149).ArraySet1(Companion_Default___.SafeModuloInt(_857_v0, _857_v0), 5)
				(_nw149).ArraySet1((func() _dafny.Map {
					var _coll71 = _dafny.NewMapBuilder()
					_ = _coll71
					for _iter72 := _dafny.Iterate((_897_v36).Elements()); ; {
						_compr_71, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _903_v34 _dafny.Sequence
						_903_v34 = interface{}(_compr_71).(_dafny.Sequence)
						if (_897_v36).Contains(_903_v34) {
							_coll71.Add(_903_v34, _857_v0)
						}
					}
					return _coll71.ToMap()
				}()).Cardinality(), 6)
				(_nw149).ArraySet1((_898_v37).Cardinality(), 7)
				(_nw149).ArraySet1(Companion_Default___.SafeDivisionInt(_857_v0, _dafny.IntOfInt64(371)), 8)
				(_nw149).ArraySet1(_857_v0, 9)
				(_nw149).ArraySet1(_857_v0, 10)
				(_nw149).ArraySet1(_857_v0, 11)
				(_nw149).ArraySet1((_891_v29).Cardinality(), 12)
				(_nw149).ArraySet1(_857_v0, 13)
				(_nw149).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_857_v0, _858_v1)).Cardinality(), 14)
				(_nw149).ArraySet1(_dafny.IntOfInt64(-418), 15)
				(_nw149).ArraySet1((_857_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), true, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), _858_v1, _858_v1)).Cardinality())), 16)
				(_nw149).ArraySet1(_857_v0, 17)
				(_nw149).ArraySet1(_857_v0, 18)
				(_nw149).ArraySet1((_864_v5).Select((Companion_Default___.SafeIndex(_857_v0, _dafny.IntOfUint32((_864_v5).Cardinality()))).Uint32()).(_dafny.Int), 19)
				(_nw149).ArraySet1((func() _dafny.Int {
					if (_901_v39).Contains(_857_v0) {
						return (_901_v39).Multiplicity(_857_v0)
					}
					return _857_v0
				})(), 20)
				(_nw149).ArraySet1(_857_v0, 21)
				(_nw149).ArraySet1((func() _dafny.Int {
					if (_891_v29).Contains((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)) {
						return (_891_v29).Get((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_896_v35).Cardinality())
				})(), 22)
				_902_v40 = _nw149
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_902_v40), 0))
				_ = _index179
				(_902_v40).ArraySet1(_857_v0, (_index179).Int())
				var _904_v41 D5
				_ = _904_v41
				_904_v41 = Companion_D5_.Create_DC14_(_dafny.IntOfUint32((Companion_Default___.Fm17(_888_v26, globalState)).Cardinality()), _858_v1)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_902_v40), 0))
				_ = _index180
				(_902_v40).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_857_v0).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_862_v3, (Companion_Default___.SafeIndex((_904_v41).Dtor_cf20(), _dafny.IntOfUint32((_862_v3).Cardinality()))).Uint32(), _888_v26)).Cardinality()))), _857_v0)), (_index180).Int())
				var _905_v42 *C1
				_ = _905_v42
				var _nw150 *C1 = New_C1_()
				_ = _nw150
				_nw150.Ctor__()
				_905_v42 = _nw150
			} else {
				var _906___mcc_h0 _dafny.Map = _source16.Get_().(D7_DC17).Cf23
				_ = _906___mcc_h0
				var _907_cf23 _dafny.Map = _906___mcc_h0
				_ = _907_cf23
				_888_v26 = _888_v26
				r1 = _dafny.IntOfInt64(-346)
				_893_v31 = _893_v31
				var _908_v43 _dafny.Array
				_ = _908_v43
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
				_ = _nw151
				_908_v43 = _nw151
				var _909_v44 D10
				_ = _909_v44
				_909_v44 = Companion_D10_.Create_DC24_(_dafny.IntOfInt64(-90), _858_v1, _dafny.SeqOf(_dafny.CodePoint('q')))
				var _910_v45 _dafny.Set
				_ = _910_v45
				_910_v45 = _dafny.SetOf(_dafny.IntOfInt64(238), (_dafny.MultiSetOf(_857_v0)).Cardinality())
				var _911_v46 _dafny.Set
				_ = _911_v46
				_911_v46 = _dafny.SetOf(_857_v0, (_910_v45).Cardinality())
				var _912_v47 *C0
				_ = _912_v47
				var _nw152 *C0 = New_C0_()
				_ = _nw152
				_nw152.Ctor__()
				_912_v47 = _nw152
				var _913_v48 D13
				_ = _913_v48
				_913_v48 = Companion_D13_.Create_DC30_(_858_v1, (_909_v44).Dtor_cf32(), _911_v46, _912_v47)
				var _pat_let_tv41 = _858_v1
				_ = _pat_let_tv41
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_908_v43), 0))
				_ = _index181
				(_908_v43).ArraySet1((func(_pat_let34_0 D13) D13 {
					return func(_914_dt__update__tmp_h1 D13) D13 {
						return func(_pat_let35_0 bool) D13 {
							return func(_915_dt__update_hcf42_h0 bool) D13 {
								return Companion_D13_.Create_DC30_(_915_dt__update_hcf42_h0, (_914_dt__update__tmp_h1).Dtor_cf43(), (_914_dt__update__tmp_h1).Dtor_cf44(), (_914_dt__update__tmp_h1).Dtor_cf45())
							}(_pat_let35_0)
						}(_pat_let_tv41)
					}(_pat_let34_0)
				}(_913_v48)).Dtor_cf44(), (_index181).Int())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_908_v43), 0))
				_ = _index182
				(_908_v43).ArraySet1(_911_v46, (_index182).Int())
			}
		}
		var _916_v49 _dafny.Sequence
		_ = _916_v49
		_916_v49 = _dafny.SeqOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), true, (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool))
		var _917_v50 D4
		_ = _917_v50
		_917_v50 = Companion_D4_.Create_DC10_(_916_v49)
		var _pat_let_tv42 = _916_v49
		_ = _pat_let_tv42
		_917_v50 = func(_pat_let36_0 D4) D4 {
			return func(_918_dt__update__tmp_h2 D4) D4 {
				return func(_pat_let37_0 _dafny.Sequence) D4 {
					return func(_919_dt__update_hcf13_h0 _dafny.Sequence) D4 {
						return Companion_D4_.Create_DC10_(_919_dt__update_hcf13_h0)
					}(_pat_let37_0)
				}(_pat_let_tv42)
			}(_pat_let36_0)
		}(_917_v50)
		r0 = _857_v0
		r1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool) {
				return _916_v49
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool), (_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)), _dafny.SeqOf((_859_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_859_v2), 0))).Int()).(bool)))
		})()).Cardinality())
		var _920_v51 D0
		_ = _920_v51
		_920_v51 = Companion_D0_.Create_DC0_(_857_v0)
		var _pat_let_tv43 = _859_v2
		_ = _pat_let_tv43
		var _pat_let_tv44 = _859_v2
		_ = _pat_let_tv44
		var _pat_let_tv45 = _859_v2
		_ = _pat_let_tv45
		var _pat_let_tv46 = _859_v2
		_ = _pat_let_tv46
		var _pat_let_tv47 = _857_v0
		_ = _pat_let_tv47
		var _pat_let_tv48 = _857_v0
		_ = _pat_let_tv48
		var _pat_let_tv49 = _857_v0
		_ = _pat_let_tv49
		r2 = func(_source17 D0) _dafny.Int {
			if _source17.Is_DC1() {
				var _921___mcc_h1 _dafny.Int = _source17.Get_().(D0_DC1).Cf1
				_ = _921___mcc_h1
				var _922___mcc_h2 _dafny.Int = _source17.Get_().(D0_DC1).Cf2
				_ = _922___mcc_h2
				var _923_cf2 _dafny.Int = _922___mcc_h2
				_ = _923_cf2
				var _924_cf1 _dafny.Int = _921___mcc_h1
				_ = _924_cf1
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_pat_let_tv43), 0))).Int()).(bool), _923_cf2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_pat_let_tv45), 0))).Int()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_923_cf2, (_dafny.MultiSetOf(_923_cf2, _923_cf2)).Cardinality())).Cardinality()))).Cardinality()
			} else if _source17.Is_DC2() {
				return _pat_let_tv47
			} else if _source17.Is_DC0() {
				var _925___mcc_h3 _dafny.Int = _source17.Get_().(D0_DC0).Cf0
				_ = _925___mcc_h3
				var _926_cf0 _dafny.Int = _925___mcc_h3
				_ = _926_cf0
				return (_926_cf0).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-441)))
			} else {
				var _927___mcc_h4 D0 = _source17.Get_().(D0_DC3).Cf3
				_ = _927___mcc_h4
				var _928_cf3 D0 = _927___mcc_h4
				_ = _928_cf3
				return _pat_let_tv48
			}
		}(func(_pat_let38_0 D0) D0 {
			return func(_929_dt__update__tmp_h3 D0) D0 {
				return func(_pat_let39_0 _dafny.Int) D0 {
					return func(_930_dt__update_hcf0_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_930_dt__update_hcf0_h0)
					}(_pat_let39_0)
				}(_pat_let_tv49)
			}(_pat_let38_0)
		}(_920_v51))
		var _931_v52 _dafny.Array
		_ = _931_v52
		var _nw153 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
		_ = _nw153
		_931_v52 = _nw153
		var _932_v53 _dafny.Sequence
		_ = _932_v53
		_932_v53 = _dafny.SeqOf(_931_v52)
		r3 = _dafny.Companion_Sequence_.Contains(_932_v53, (_932_v53).Select((Companion_Default___.SafeIndex(_857_v0, _dafny.IntOfUint32((_932_v53).Cardinality()))).Uint32()).(_dafny.Array))
		return r0, r1, r2, r3
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	dummy byte
}

func New_C3_() *C3 {
	_this := C3{}

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

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-878), _dafny.IntOfInt64(740))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(663), _dafny.IntOfUint32((_dafny.SeqOf(!(true), false)).Cardinality())))).Merge(func() _dafny.Map {
			var _coll72 = _dafny.NewMapBuilder()
			_ = _coll72
			for _iter73 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll73 = _dafny.NewMapBuilder()
				_ = _coll73
				for _iter74 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), _dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(603)))))).Elements()); ; {
					_compr_73, _ok74 := _iter74()
					if !_ok74 {
						break
					}
					var _933_v1 _dafny.Sequence
					_933_v1 = interface{}(_compr_73).(_dafny.Sequence)
					if (_dafny.SetOf(_dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), _dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(603)))))).Contains(_933_v1) {
						_coll73.Add(_933_v1, false)
					}
				}
				return _coll73.ToMap()
			}()).Cardinality(), true)).Keys().Elements()); ; {
				_compr_72, _ok73 := _iter73()
				if !_ok73 {
					break
				}
				var _934_v0 _dafny.Int
				_934_v0 = interface{}(_compr_72).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll74 = _dafny.NewMapBuilder()
					_ = _coll74
					for _iter75 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), _dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(603)))))).Elements()); ; {
						_compr_74, _ok75 := _iter75()
						if !_ok75 {
							break
						}
						var _933_v1 _dafny.Sequence
						_933_v1 = interface{}(_compr_74).(_dafny.Sequence)
						if (_dafny.SetOf(_dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), _dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_())), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(603)))))).Contains(_933_v1) {
							_coll74.Add(_933_v1, false)
						}
					}
					return _coll74.ToMap()
				}()).Cardinality(), true)).Contains(_934_v0) {
					_coll72.Add(Companion_Default___.SafeDivisionInt(_934_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("idrigq")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-517))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_935_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					}))).Cardinality()))
				}
			}
			return _coll72.ToMap()
		}())
	}
}
func (_this *C3) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return !(!(false))
	}
}
func (_this *C3) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(401), _dafny.IntOfInt64(811))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ifkhv")).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, false)).Cardinality(), false)).Cardinality())))).Cardinality()
	}
}
func (_this *C3) Fm16(p0 bool, globalState *GlobalState) bool {
	{
		if true {
			return (true) == (!(true))
		} else {
			return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('k'), _dafny.CodePoint('j'), _dafny.CodePoint('u'), _dafny.CodePoint('x')))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())) != 0
		}
	}
}
func (_this *C3) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _936_v0 _dafny.Array
		_ = _936_v0
		var _len0_25 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_25
		var _nw154 _dafny.Array
		_ = _nw154
		if _len0_25.Cmp(_dafny.Zero) == 0 {
			_nw154 = _dafny.NewArray(_len0_25)
		} else {
			var _init25 func(_dafny.Int) bool = func(_937_i0 _dafny.Int) bool {
				return !(false)
			}
			_ = _init25
			var _element0_25 = _init25(_dafny.Zero)
			_ = _element0_25
			_nw154 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
			(_nw154).ArraySet1(_element0_25, 0)
			var _nativeLen0_25 = (_len0_25).Int()
			_ = _nativeLen0_25
			for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
				(_nw154).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
			}
		}
		_936_v0 = _nw154
		var _938_v1 bool
		_ = _938_v1
		_938_v1 = true
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
		_ = _index183
		(_936_v0).ArraySet1(_938_v1, (_index183).Int())
		var _939_v2 _dafny.Int
		_ = _939_v2
		_939_v2 = _dafny.IntOfInt64(370)
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
		_ = _index184
		(_936_v0).ArraySet1((_939_v2).Cmp((_dafny.Zero).Minus(_939_v2)) > 0, (_index184).Int())
		var _hi4 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-926))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_940_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality())
		_ = _hi4
		for _941_i1 := _939_v2; _941_i1.Cmp(_hi4) < 0; _941_i1 = _941_i1.Plus(_dafny.One) {
			var _942_v3 _dafny.Sequence
			_ = _942_v3
			_942_v3 = _dafny.UnicodeSeqOfUtf8Bytes("fm")
			var _943_v4 _dafny.Sequence
			_ = _943_v4
			_943_v4 = _dafny.SeqOf(_942_v3)
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((_943_v4).Select((Companion_Default___.SafeIndex(_939_v2, _dafny.IntOfUint32((_943_v4).Cardinality()))).Uint32()).(_dafny.Sequence), _942_v3), _942_v3) {
				var _944_v5 D0
				_ = _944_v5
				_944_v5 = Companion_D0_.Create_DC1_(_941_i1, _941_i1)
				(globalState).F8 = (_dafny.Zero).Minus((_944_v5).Dtor_cf1())
				var _945_v6 _dafny.CodePoint
				_ = _945_v6
				_945_v6 = _dafny.CodePoint('w')
				_945_v6 = _945_v6
				_942_v3 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_942_v3, _942_v3), (Companion_Default___.SafeIndex(_939_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_942_v3, _942_v3)).Cardinality()))).Uint32(), _945_v6)
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
				_ = _index185
				(_936_v0).ArraySet1(!(_938_v1), (_index185).Int())
				var _946_v7 *C0
				_ = _946_v7
				var _nw155 *C0 = New_C0_()
				_ = _nw155
				_nw155.Ctor__()
				_946_v7 = _nw155
			} else {
				var _947_v8 _dafny.Map
				_ = _947_v8
				_947_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_941_i1, _939_v2)
				var _948_v9 D10
				_ = _948_v9
				_948_v9 = Companion_D10_.Create_DC24_((func() _dafny.Int {
					if (_947_v8).Contains(_941_i1) {
						return (_947_v8).Get(_941_i1).(_dafny.Int)
					}
					return _939_v2
				})(), (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), _942_v3)
				r0 = (_948_v9).Dtor_cf32()
				(globalState).F8 = _939_v2
				(globalState).F8 = _939_v2
				var _949_v10 _dafny.Set
				_ = _949_v10
				_949_v10 = _dafny.SetOf(_941_i1, _dafny.IntOfInt64(903), (_dafny.Zero).Minus(_939_v2))
				var _950_v11 *C0
				_ = _950_v11
				var _nw156 *C0 = New_C0_()
				_ = _nw156
				_nw156.Ctor__()
				_950_v11 = _nw156
				var _951_v12 D13
				_ = _951_v12
				_951_v12 = Companion_D13_.Create_DC30_(true, _938_v1, _949_v10, _950_v11)
				r1 = (_951_v12).Dtor_cf42()
				var _952_v13 _dafny.Map
				_ = _952_v13
				_952_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_939_v2).Cmp(_939_v2) > 0, (_950_v11).Fm19(globalState))
				_952_v13 = ((_952_v13).Merge(_952_v13)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v1, _941_i1))
			}
			var _953_v14 _dafny.Map
			_ = _953_v14
			_953_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v2, _938_v1)
			var _954_v15 _dafny.Sequence
			_ = _954_v15
			_954_v15 = _dafny.SeqOf(_953_v14)
			var _955_v16 _dafny.Map
			_ = _955_v16
			_955_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_954_v15, _954_v15), _941_i1)
			_955_v16 = (_955_v16).Update(false, (_939_v2).Minus(_939_v2))
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
			_ = _index186
			var _rhs164 bool = !((_941_i1).Cmp(_941_i1) > 0)
			_ = _rhs164
			var _rhs165 _dafny.Int = _941_i1
			_ = _rhs165
			var _rhs166 bool = _938_v1
			_ = _rhs166
			var _rhs167 bool = _938_v1
			_ = _rhs167
			var _rhs168 bool = (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
			_ = _rhs168
			var _lhs134 _dafny.Array = _936_v0
			_ = _lhs134
			var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
			_ = _lhs135
			var _lhs136 *GlobalState = globalState
			_ = _lhs136
			var _lhs137 *GlobalState = globalState
			_ = _lhs137
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			(_lhs134).ArraySet1(_rhs164, (_lhs135).Int())
			_lhs136.F8 = _rhs165
			_lhs137.F4 = _rhs166
			r0 = _rhs167
			_lhs138.F4 = _rhs168
			if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
				var _956_v17 _dafny.Set
				_ = _956_v17
				_956_v17 = _dafny.SetOf(_942_v3, _dafny.UnicodeSeqOfUtf8Bytes("dfqsa"), _dafny.Companion_Sequence_.Concatenate(_942_v3, _942_v3), _942_v3, _942_v3)
				_956_v17 = (_dafny.SetOf(_942_v3)).Union(_956_v17)
				var _rhs169 _dafny.Int = _941_i1
				_ = _rhs169
				var _rhs170 bool = (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
				_ = _rhs170
				var _rhs171 _dafny.Int = _939_v2
				_ = _rhs171
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				_939_v2 = _rhs169
				_938_v1 = _rhs170
				_lhs139.F8 = _rhs171
				var _957_v18 _dafny.Map
				_ = _957_v18
				_957_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_941_i1, _941_i1)
				var _958_v20 _dafny.Array
				_ = _958_v20
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw157
				_958_v20 = _nw157
				var _959_v21 _dafny.Map
				_ = _959_v21
				_959_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_958_v20, false)
				var _960_v22 _dafny.Map
				_ = _960_v22
				_960_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
					var _coll75 = _dafny.NewBuilder()
					_ = _coll75
					for _iter76 := _dafny.Iterate((_957_v18).Keys().Elements()); ; {
						_compr_75, _ok76 := _iter76()
						if !_ok76 {
							break
						}
						var _961_v19 _dafny.Int
						_961_v19 = interface{}(_compr_75).(_dafny.Int)
						if (_957_v18).Contains(_961_v19) {
							_coll75.Add((_961_v19).Times(_dafny.IntOfInt64(664)))
						}
					}
					return _coll75.ToSet()
				}(), (func() bool {
					if (_959_v21).Contains(_958_v20) {
						return (_959_v21).Get(_958_v20).(bool)
					}
					return _938_v1
				})())
				_960_v22 = _960_v22
				var _962_v23 D3
				_ = _962_v23
				var _out21 D3
				_ = _out21
				_out21 = (_this).M10(_941_i1, (_953_v14).Merge(_953_v14), _dafny.IntOfInt64(329), globalState)
				_962_v23 = _out21
				var _963_v24 _dafny.Sequence
				_ = _963_v24
				_963_v24 = _dafny.SeqOf((_dafny.Zero).Minus(_939_v2), (_dafny.Zero).Minus((_dafny.Zero).Minus(_939_v2)), _dafny.IntOfUint32((_942_v3).Cardinality()))
				var _964_v25 _dafny.Array
				_ = _964_v25
				var _nwElement0_35 _dafny.Int = _939_v2
				_ = _nwElement0_35
				var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(6))
				_ = _nw158
				(_nw158).ArraySet1(_nwElement0_35, 0)
				(_nw158).ArraySet1((_963_v24).Select((Companion_Default___.SafeIndex(_941_i1, _dafny.IntOfUint32((_963_v24).Cardinality()))).Uint32()).(_dafny.Int), 1)
				(_nw158).ArraySet1(_941_i1, 2)
				(_nw158).ArraySet1(_941_i1, 3)
				(_nw158).ArraySet1(_939_v2, 4)
				(_nw158).ArraySet1(_939_v2, 5)
				_964_v25 = _nw158
				var _965_v26 D3
				_ = _965_v26
				_965_v26 = Companion_D3_.Create_DC8_(_957_v18)
				var _966_v27 _dafny.Sequence
				_ = _966_v27
				_966_v27 = _dafny.SeqOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool))
				var _967_v28 _dafny.MultiSet
				_ = _967_v28
				_967_v28 = _dafny.MultiSetOf(_966_v27)
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
				_ = _index187
				var _rhs172 _dafny.Int = Companion_Default___.SafeDivisionInt((_939_v2).Times(Companion_Default___.Fm26(_965_v26, _939_v2, _967_v28, globalState)), _dafny.IntOfInt64(853))
				_ = _rhs172
				var _rhs173 bool = (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
				_ = _rhs173
				var _rhs174 _dafny.Array = _964_v25
				_ = _rhs174
				var _rhs175 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_942_v3, _942_v3), _942_v3)
				_ = _rhs175
				var _rhs176 bool = false
				_ = _rhs176
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				var _lhs141 _dafny.Array = _936_v0
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
				_ = _lhs142
				_lhs140.F3 = _rhs172
				(_lhs141).ArraySet1(_rhs173, (_lhs142).Int())
				_964_v25 = _rhs174
				_942_v3 = _rhs175
				r0 = _rhs176
			} else {
				_938_v1 = _938_v1
				_939_v2 = Companion_Default___.Fm20((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), globalState)
				var _968_v29 _dafny.Array
				_ = _968_v29
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw159
				_968_v29 = _nw159
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_968_v29), 0))
				_ = _index188
				(_968_v29).ArraySet1(_941_i1, (_index188).Int())
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_968_v29), 0))
				_ = _index189
				(_968_v29).ArraySet1(_dafny.IntOfUint32((_942_v3).Cardinality()), (_index189).Int())
				_968_v29 = _968_v29
				var _969_v30 _dafny.CodePoint
				_ = _969_v30
				_969_v30 = _dafny.CodePoint('r')
				_969_v30 = _969_v30
			}
		}
		if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
			(globalState).F8 = ((_939_v2).Minus(_939_v2)).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(625), _dafny.IntOfInt64(995)))))
			var _970_v31 _dafny.MultiSet
			_ = _970_v31
			_970_v31 = _dafny.MultiSetOf(_939_v2, _939_v2)
			var _971_v32 _dafny.Set
			_ = _971_v32
			_971_v32 = _dafny.SetOf(_939_v2)
			var _972_v33 *C0
			_ = _972_v33
			var _nw160 *C0 = New_C0_()
			_ = _nw160
			_nw160.Ctor__()
			_972_v33 = _nw160
			var _973_v34 D13
			_ = _973_v34
			_973_v34 = Companion_D13_.Create_DC30_(_938_v1, true, _971_v32, _972_v33)
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
			_ = _index190
			var _rhs177 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
					return _939_v2
				}
				return _dafny.IntOfInt64(889)
			})()))
			_ = _rhs177
			var _rhs178 bool = (_970_v31).IsSubsetOf((_970_v31).Update(_939_v2, Companion_Default___.Abs(_939_v2)))
			_ = _rhs178
			var _rhs179 bool = (_973_v34).Dtor_cf42()
			_ = _rhs179
			var _rhs180 bool = (_939_v2).Cmp(_939_v2) == 0
			_ = _rhs180
			var _lhs143 *GlobalState = globalState
			_ = _lhs143
			var _lhs144 _dafny.Array = _936_v0
			_ = _lhs144
			var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
			_ = _lhs145
			_939_v2 = _rhs177
			_lhs143.F4 = _rhs178
			(_lhs144).ArraySet1(_rhs179, (_lhs145).Int())
			_938_v1 = _rhs180
			var _974_v35 _dafny.Array
			_ = _974_v35
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw161
			_974_v35 = _nw161
			var _975_v36 _dafny.Array
			_ = _975_v36
			var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw162
			_975_v36 = _nw162
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))
			_ = _index191
			(_974_v35).ArraySet1(_975_v36, (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))
			_ = _index192
			(_974_v35).ArraySet1(_975_v36, (_index192).Int())
			var _976_v37 _dafny.Set
			_ = _976_v37
			_976_v37 = _dafny.SetOf(_938_v1)
			_976_v37 = (_976_v37).Difference(_dafny.SetOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)))
			if (_938_v1) || (_938_v1) {
				(globalState).F4 = !((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool))
				r1 = (_this).Fm16(_938_v1, globalState)
				var _977_v38 *C1
				_ = _977_v38
				var _nw163 *C1 = New_C1_()
				_ = _nw163
				_nw163.Ctor__()
				_977_v38 = _nw163
				var _978_v39 D0
				_ = _978_v39
				_978_v39 = Companion_D0_.Create_DC2_()
				var _979_v40 D9
				_ = _979_v40
				_979_v40 = Companion_D9_.Create_DC21_((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), _dafny.SeqOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), _938_v1, (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)), _938_v1)
				var _rhs181 D0 = _978_v39
				_ = _rhs181
				var _rhs182 _dafny.Int = (_dafny.Zero).Minus(_939_v2)
				_ = _rhs182
				var _rhs183 bool = (func() bool {
					if _dafny.Companion_Sequence_.IsProperPrefixOf((_979_v40).Dtor_cf27(), _dafny.SeqOf(_938_v1, (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), false)) {
						return (_939_v2).Cmp(_939_v2) <= 0
					}
					return ((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)) && (false)
				})()
				_ = _rhs183
				var _rhs184 bool = _938_v1
				_ = _rhs184
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				_978_v39 = _rhs181
				_lhs146.F3 = _rhs182
				r0 = _rhs183
				r1 = _rhs184
				var _980_v41 _dafny.Array
				_ = _980_v41
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_26
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.CodePoint = func(_981_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('b')
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw164 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw164).ArraySet1CodePoint(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw164).ArraySet1CodePoint(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_980_v41 = _nw164
				var _982_v42 _dafny.Sequence
				_ = _982_v42
				_982_v42 = _dafny.SeqOf((func() _dafny.Array {
					if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
						return _980_v41
					}
					return _980_v41
				})(), _980_v41, _980_v41)
				var _rhs185 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_982_v42, _982_v42), _982_v42)
				_ = _rhs185
				var _rhs186 _dafny.Int = Companion_Default___.SafeModuloInt(_939_v2, _dafny.IntOfInt64(776))
				_ = _rhs186
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				_982_v42 = _rhs185
				_lhs147.F8 = _rhs186
			} else {
				var _983_v43 _dafny.Sequence
				_ = _983_v43
				_983_v43 = _dafny.UnicodeSeqOfUtf8Bytes("ae")
				var _984_v44 _dafny.Map
				_ = _984_v44
				_984_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_983_v43, _939_v2)
				_984_v44 = _984_v44
				var _985_v45 _dafny.CodePoint
				_ = _985_v45
				_985_v45 = _dafny.CodePoint('g')
				var _986_v46 D3
				_ = _986_v46
				_986_v46 = Companion_D3_.Create_DC9_()
				var _987_v47 _dafny.Map
				_ = _987_v47
				_987_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_985_v45, _986_v46)
				_987_v47 = (_987_v47).Update(_985_v45, _986_v46)
				var _988_v48 _dafny.Map
				_ = _988_v48
				_988_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_938_v1), _939_v2)
				var _989_v49 _dafny.Map
				_ = _989_v49
				_989_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v2, (_988_v48).Cardinality())
				var _990_v50 _dafny.Sequence
				_ = _990_v50
				_990_v50 = _dafny.SeqOf(((_988_v48).Update((_this).Fm16((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_983_v43).Cardinality()))).Cardinality(), (_989_v49).Cardinality())
				(globalState).F4 = (_this).Fm3(_990_v50, (_970_v31).Difference(_970_v31), _938_v1, globalState)
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_975_v36), 0))
				_ = _index193
				(_975_v36).ArraySet1((_989_v49).Cardinality(), (_index193).Int())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_975_v36), 0))
				_ = _index194
				(_975_v36).ArraySet1(Companion_Default___.SafeModuloInt(_939_v2, (_990_v50).Select((Companion_Default___.SafeIndex(_939_v2, _dafny.IntOfUint32((_990_v50).Cardinality()))).Uint32()).(_dafny.Int)), (_index194).Int())
				var _991_v51 _dafny.Sequence
				_ = _991_v51
				_991_v51 = _dafny.SeqOf(_dafny.ArrayCastTo((_974_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))).Int())), (func() _dafny.Array {
					if _938_v1 {
						return _dafny.ArrayCastTo((_974_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))).Int()))
					}
					return _dafny.ArrayCastTo((_974_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))).Int()))
				})(), _dafny.ArrayCastTo((_974_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))).Int())), _dafny.ArrayCastTo((_974_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_974_v35), 0))).Int())))
				_991_v51 = (func() _dafny.Sequence {
					if true {
						return _dafny.Companion_Sequence_.Concatenate(_991_v51, _991_v51)
					}
					return _991_v51
				})()
			}
		} else {
			var _992_v52 _dafny.Array
			_ = _992_v52
			var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
			_ = _nw165
			_992_v52 = _nw165
			var _993_v53 _dafny.Map
			_ = _993_v53
			_993_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v2, _992_v52)
			_993_v53 = (_993_v53).Update(_939_v2, _992_v52)
			var _994_v54 _dafny.Sequence
			_ = _994_v54
			_994_v54 = _dafny.UnicodeSeqOfUtf8Bytes("aeeyibtfx")
			var _995_v55 _dafny.MultiSet
			_ = _995_v55
			_995_v55 = _dafny.MultiSetOf(_994_v54)
			var _996_v56 D10
			_ = _996_v56
			_996_v56 = Companion_D10_.Create_DC23_(_995_v55)
			var _pat_let_tv50 = _995_v55
			_ = _pat_let_tv50
			_996_v56 = func(_pat_let40_0 D10) D10 {
				return func(_997_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let41_0 _dafny.MultiSet) D10 {
						return func(_998_dt__update_hcf30_h0 _dafny.MultiSet) D10 {
							return Companion_D10_.Create_DC23_(_998_dt__update_hcf30_h0)
						}(_pat_let41_0)
					}(_pat_let_tv50)
				}(_pat_let40_0)
			}((func() D10 {
				if _938_v1 {
					return Companion_Default___.Fm36(_dafny.IntOfUint32((_994_v54).Cardinality()), _dafny.SetOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)), _938_v1, globalState)
				}
				return Companion_D10_.Create_DC23_(_995_v55)
			})())
			(globalState).F2 = Companion_Default___.Fm27(_938_v1, globalState)
			var _999_v57 _dafny.Sequence
			_ = _999_v57
			_999_v57 = _dafny.SeqOf(_dafny.IntOfInt64(-589))
			var _1000_v58 _dafny.MultiSet
			_ = _1000_v58
			_1000_v58 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wc")).Cardinality()), _dafny.IntOfUint32((_999_v57).Cardinality()))
			var _1001_v59 _dafny.Sequence
			_ = _1001_v59
			_1001_v59 = _dafny.SeqOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool))
			var _1002_v60 D10
			_ = _1002_v60
			_1002_v60 = Companion_D10_.Create_DC24_(_939_v2, _938_v1, _994_v54)
			var _1003_v61 D9
			_ = _1003_v61
			_1003_v61 = Companion_D9_.Create_DC21_(_938_v1, _1001_v59, (_1002_v60).Dtor_cf32())
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))
			_ = _index195
			(_936_v0).ArraySet1(((func() bool {
				if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
					return (_this).Fm3(_999_v57, _1000_v58, _938_v1, globalState)
				}
				return (_1003_v61).Dtor_cf26()
			})()) || (false), (_index195).Int())
			var _1004_v62 _dafny.CodePoint
			_ = _1004_v62
			_1004_v62 = _dafny.CodePoint('o')
			(globalState).F3 = (_939_v2).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_994_v54, (Companion_Default___.SafeIndex(_939_v2, _dafny.IntOfUint32((_994_v54).Cardinality()))).Uint32(), _1004_v62)).Cardinality()))
		}
		(globalState).F3 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())
		var _hi5 _dafny.Int = _939_v2
		_ = _hi5
		for _1005_i4 := _939_v2; _1005_i4.Cmp(_hi5) < 0; _1005_i4 = _1005_i4.Plus(_dafny.One) {
			(globalState).F3 = Companion_Default___.SafeModuloInt(_1005_i4, (_dafny.MultiSetOf(_938_v1)).Cardinality())
			r0 = (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
			var _1006_v63 *C2
			_ = _1006_v63
			var _nw166 *C2 = New_C2_()
			_ = _nw166
			_nw166.Ctor__()
			_1006_v63 = _nw166
			(globalState).F3 = _1005_i4
		}
		var _1007_v64 D2
		_ = _1007_v64
		_1007_v64 = Companion_D2_.Create_DC7_(_dafny.CodePoint('l'), (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), Companion_Default___.Fm20((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), globalState))
		var _source18 D2 = _1007_v64
		_ = _source18
		if _source18.Is_DC6() {
			var _1008___mcc_h0 _dafny.Int = _source18.Get_().(D2_DC6).Cf6
			_ = _1008___mcc_h0
			var _1009___mcc_h1 bool = _source18.Get_().(D2_DC6).Cf7
			_ = _1009___mcc_h1
			var _1010___mcc_h2 bool = _source18.Get_().(D2_DC6).Cf8
			_ = _1010___mcc_h2
			var _1011_cf8 bool = _1010___mcc_h2
			_ = _1011_cf8
			var _1012_cf7 bool = _1009___mcc_h1
			_ = _1012_cf7
			var _1013_cf6 _dafny.Int = _1008___mcc_h0
			_ = _1013_cf6
			var _1014_v65 _dafny.Map
			_ = _1014_v65
			_1014_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v1, _938_v1)
			var _1015_v66 _dafny.MultiSet
			_ = _1015_v66
			_1015_v66 = _dafny.MultiSetOf(!(_938_v1), _1012_cf7, _1012_cf7, (func() bool {
				if (_1014_v65).Contains(_1012_cf7) {
					return (_1014_v65).Get(_1012_cf7).(bool)
				}
				return !(false)
			})())
			var _1016_v67 _dafny.Map
			_ = _1016_v67
			_1016_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1015_v66).Union(_1015_v66), _1013_cf6)
			var _1017_v68 D0
			_ = _1017_v68
			_1017_v68 = Companion_D0_.Create_DC1_(_1013_cf6, _1013_cf6)
			var _1018_v69 D0
			_ = _1018_v69
			_1018_v69 = Companion_D0_.Create_DC3_(_1017_v68)
			var _1019_v70 _dafny.Sequence
			_ = _1019_v70
			_1019_v70 = _dafny.SeqOf(_1018_v69, _1018_v69, Companion_D0_.Create_DC3_(_1017_v68), Companion_Default___.Fm38(_938_v1, globalState), Companion_D0_.Create_DC3_(_1017_v68))
			var _1020_v71 _dafny.CodePoint
			_ = _1020_v71
			_1020_v71 = _dafny.CodePoint('o')
			_1016_v67 = Companion_Default___.Fm37(_1019_v70, Companion_Default___.Fm17(_1020_v71, globalState), _939_v2, (_939_v2).Plus(_1013_cf6), globalState)
			var _1021_v73 _dafny.Map
			_ = _1021_v73
			_1021_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v2, _1013_cf6)
			var _1022_v74 _dafny.Set
			_ = _1022_v74
			_1022_v74 = _dafny.SetOf((func() _dafny.Map {
				var _coll76 = _dafny.NewMapBuilder()
				_ = _coll76
				for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(143), _dafny.IntOfInt64(-189))); ; {
					_compr_76, _ok77 := _iter77()
					if !_ok77 {
						break
					}
					var _1023_v72 _dafny.Int
					_1023_v72 = interface{}(_compr_76).(_dafny.Int)
					if ((_dafny.IntOfInt64(143)).Cmp(_1023_v72) <= 0) && ((_1023_v72).Cmp(_dafny.IntOfInt64(-189)) < 0) {
						_coll76.Add((_1023_v72).Times(_dafny.IntOfUint32((_dafny.SeqOf((_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool), !(_1012_cf7), !(_938_v1), _1011_cf8)).Cardinality())), _939_v2)
					}
				}
				return _coll76.ToMap()
			}()).Merge(_1021_v73))
			var _1024_v75 _dafny.Sequence
			_ = _1024_v75
			_1024_v75 = _dafny.SeqOf(_1022_v74)
			_1022_v74 = (_1024_v75).Select((Companion_Default___.SafeIndex(_1013_cf6, _dafny.IntOfUint32((_1024_v75).Cardinality()))).Uint32()).(_dafny.Set)
			var _1025_v76 _dafny.Array
			_ = _1025_v76
			var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw167
			_1025_v76 = _nw167
			_1025_v76 = _1025_v76
			var _1026_v77 _dafny.Sequence
			_ = _1026_v77
			_1026_v77 = _dafny.SeqOf((func() bool {
				if (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool) {
					return _1011_cf8
				}
				return (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
			})())
			var _1027_v78 _dafny.Array
			_ = _1027_v78
			var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
			_ = _nw168
			_1027_v78 = _nw168
			var _1028_v79 _dafny.Set
			_ = _1028_v79
			_1028_v79 = _dafny.SetOf(_1027_v78)
			var _1029_v80 _dafny.Set
			_ = _1029_v80
			_1029_v80 = _dafny.SetOf(_1012_cf7, (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool))
			_1026_v77 = Companion_Default___.Fm31((_1028_v79).IsProperSubsetOf(_dafny.SetOf(_1027_v78, _1027_v78, _1027_v78)), (_dafny.SetOf((_1026_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1013_cf6), _dafny.IntOfUint32((_1026_v77).Cardinality()))).Uint32()).(bool), _1011_cf8)).IsSubsetOf(_1029_v80), !((func() bool {
				if _1012_cf7 {
					return (_936_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_936_v0), 0))).Int()).(bool)
				}
				return _938_v1
			})()), globalState)
		} else if _source18.Is_DC7() {
			var _1030___mcc_h3 _dafny.CodePoint = _source18.Get_().(D2_DC7).Cf9
			_ = _1030___mcc_h3
			var _1031___mcc_h4 bool = _source18.Get_().(D2_DC7).Cf10
			_ = _1031___mcc_h4
			var _1032___mcc_h5 _dafny.Int = _source18.Get_().(D2_DC7).Cf11
			_ = _1032___mcc_h5
			var _1033_cf11 _dafny.Int = _1032___mcc_h5
			_ = _1033_cf11
			var _1034_cf10 bool = _1031___mcc_h4
			_ = _1034_cf10
			var _1035_cf9 _dafny.CodePoint = _1030___mcc_h3
			_ = _1035_cf9
			var _1036_v81 _dafny.Map
			_ = _1036_v81
			_1036_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1034_cf10, _1035_cf9)
			_1036_v81 = (_1036_v81).Update((_this).Fm16(_1034_cf10, globalState), _1035_cf9)
			(globalState).F4 = false
			var _1037_v82 *C1
			_ = _1037_v82
			var _nw169 *C1 = New_C1_()
			_ = _nw169
			_nw169.Ctor__()
			_1037_v82 = _nw169
			var _1038_v83 *C0
			_ = _1038_v83
			var _nw170 *C0 = New_C0_()
			_ = _nw170
			_nw170.Ctor__()
			_1038_v83 = _nw170
		} else {
			var _1039___mcc_h6 _dafny.Array = _source18.Get_().(D2_DC5).Cf5
			_ = _1039___mcc_h6
			var _1040_cf5 _dafny.Array = _1039___mcc_h6
			_ = _1040_cf5
			var _1041_v84 _dafny.Map
			_ = _1041_v84
			_1041_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v2, _939_v2)
			var _1042_v85 D3
			_ = _1042_v85
			_1042_v85 = Companion_D3_.Create_DC8_(_1041_v84)
			_1042_v85 = Companion_D3_.Create_DC8_(_1041_v84)
			var _1043_v86 _dafny.CodePoint
			_ = _1043_v86
			_1043_v86 = _dafny.CodePoint('t')
			_1043_v86 = _1043_v86
			var _1044_v87 _dafny.Sequence
			_ = _1044_v87
			_1044_v87 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("exys"))
			var _1045_v88 _dafny.Map
			_ = _1045_v88
			_1045_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(934), (_1044_v87).Select((Companion_Default___.SafeIndex(_939_v2, _dafny.IntOfUint32((_1044_v87).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1046_v89 _dafny.Set
			_ = _1046_v89
			_1046_v89 = _dafny.SetOf(_1043_v86)
			var _1047_v90 _dafny.Sequence
			_ = _1047_v90
			_1047_v90 = _dafny.UnicodeSeqOfUtf8Bytes("kjjgia")
			_1045_v88 = (_1045_v88).Update((_939_v2).Plus((_1046_v89).Cardinality()), _1047_v90)
			(globalState).F8 = _939_v2
		}
		var _1048_v91 _dafny.Set
		_ = _1048_v91
		_1048_v91 = _dafny.SetOf(_938_v1)
		r0 = (_1048_v91).IsDisjointFrom(_1048_v91)
		r1 = _938_v1
		return r0, r1
	}
}
func (_this *C3) M10(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) D3 {
	{
		var r0 D3 = Companion_D3_.Default()
		_ = r0
		var _1049_v0 bool
		_ = _1049_v0
		_1049_v0 = false
		var _1050_v1 _dafny.CodePoint
		_ = _1050_v1
		_1050_v1 = _dafny.CodePoint('w')
		var _1051_v2 _dafny.Sequence
		_ = _1051_v2
		_1051_v2 = _dafny.SeqOf(_1050_v1, _1050_v1)
		var _1052_v3 D10
		_ = _1052_v3
		_1052_v3 = Companion_D10_.Create_DC24_(p2, _1049_v0, _1051_v2)
		if (_1052_v3).Dtor_cf32() {
			if (_1049_v0) || (false) {
				var _1053_v4 _dafny.Array
				_ = _1053_v4
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw171
				_1053_v4 = _nw171
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))
				_ = _index196
				(_1053_v4).ArraySet1(_1049_v0, (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))
				_ = _index197
				(_1053_v4).ArraySet1(_1049_v0, (_index197).Int())
				var _1054_v5 _dafny.Array
				_ = _1054_v5
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw172
				_1054_v5 = _nw172
				var _1055_v6 _dafny.Array
				_ = _1055_v6
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw173
				_1055_v6 = _nw173
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1055_v6), 0))
				_ = _index198
				(_1055_v6).ArraySet1(p0, (_index198).Int())
				var _1056_v7 _dafny.Sequence
				_ = _1056_v7
				_1056_v7 = _dafny.SeqOf(p0)
				var _1057_v8 _dafny.MultiSet
				_ = _1057_v8
				_1057_v8 = _dafny.MultiSetOf(_dafny.IntOfInt64(-279))
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))
				_ = _index199
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1055_v6), 0))
				_ = _index200
				var _rhs187 bool = (_this).Fm3(_dafny.Companion_Sequence_.Concatenate(_1056_v7, _1056_v7), _1057_v8, _1049_v0, globalState)
				_ = _rhs187
				var _rhs188 _dafny.Array = _1054_v5
				_ = _rhs188
				var _rhs189 _dafny.Int = (_this).Fm15(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1051_v2, _dafny.UnicodeSeqOfUtf8Bytes("x"))).Cardinality()), globalState)
				_ = _rhs189
				var _rhs190 _dafny.Int = p0
				_ = _rhs190
				var _lhs148 _dafny.Array = _1053_v4
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))
				_ = _lhs149
				var _lhs150 _dafny.Array = _1055_v6
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1055_v6), 0))
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				(_lhs148).ArraySet1(_rhs187, (_lhs149).Int())
				_1054_v5 = _rhs188
				(_lhs150).ArraySet1(_rhs189, (_lhs151).Int())
				_lhs152.F3 = _rhs190
				(globalState).F0 = _1056_v7
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))
				_ = _index201
				(_1053_v4).ArraySet1((_1053_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))).Int()).(bool), (_index201).Int())
				var _1058_v9 _dafny.Map
				_ = _1058_v9
				_1058_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1059_v10 _dafny.Map
				_ = _1059_v10
				_1059_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1058_v9, (_1053_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1053_v4), 0))).Int()).(bool))
				(globalState).F0 = _dafny.Companion_Sequence_.Concatenate(_1056_v7, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1056_v7, _1056_v7), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1059_v10).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1056_v7, _1056_v7)).Cardinality()))).Uint32(), p0))
			} else {
				(globalState).F8 = Companion_Default___.SafeDivisionInt((p2).Minus(p0), p0)
				(globalState).F8 = p2
				var _1060_v11 _dafny.Sequence
				_ = _1060_v11
				_1060_v11 = _dafny.SeqOf(_1049_v0, _1049_v0, !(_1049_v0))
				var _1061_v12 D4
				_ = _1061_v12
				_1061_v12 = Companion_D4_.Create_DC10_(_1060_v11)
				var _1062_v13 _dafny.Sequence
				_ = _1062_v13
				_1062_v13 = _dafny.SeqOf(Companion_D4_.Create_DC12_(_1061_v12))
				var _1063_v15 D9
				_ = _1063_v15
				_1063_v15 = Companion_D9_.Create_DC22_(_1060_v11)
				var _1064_v16 _dafny.Map
				_ = _1064_v16
				_1064_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1063_v15)).Cardinality()), _dafny.IntOfUint32((_1051_v2).Cardinality()))
				var _1065_v17 _dafny.Map
				_ = _1065_v17
				_1065_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1064_v16, (_this).Fm16(_1049_v0, globalState))
				var _1066_v18 _dafny.Map
				_ = _1066_v18
				_1066_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v13, (func() _dafny.Map {
					var _coll77 = _dafny.NewMapBuilder()
					_ = _coll77
					for _iter78 := _dafny.Iterate((_1065_v17).Keys().Elements()); ; {
						_compr_77, _ok78 := _iter78()
						if !_ok78 {
							break
						}
						var _1067_v14 _dafny.Map
						_1067_v14 = interface{}(_compr_77).(_dafny.Map)
						if (_1065_v17).Contains(_1067_v14) {
							_coll77.Add(_1067_v14, p2)
						}
					}
					return _coll77.ToMap()
				}()).Update(_1064_v16, p0))
				var _1068_v19 _dafny.Map
				_ = _1068_v19
				_1068_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1064_v16, p0)
				_1066_v18 = (_1066_v18).Update(_1062_v13, _1068_v19)
				var _1069_v20 _dafny.MultiSet
				_ = _1069_v20
				_1069_v20 = _dafny.MultiSetOf(Companion_Default___.Fm20(_1049_v0, globalState))
				(globalState).F2 = (_1069_v20).Difference(_1069_v20)
				var _1070_v21 _dafny.Map
				_ = _1070_v21
				_1070_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _dafny.IntOfUint32((_1051_v2).Cardinality()))
				var _1071_v22 D12
				_ = _1071_v22
				_1071_v22 = Companion_D12_.Create_DC26_(_1070_v21)
				var _pat_let_tv51 = _1049_v0
				_ = _pat_let_tv51
				var _pat_let_tv52 = p2
				_ = _pat_let_tv52
				var _pat_let_tv53 = _1070_v21
				_ = _pat_let_tv53
				var _1072_v23 _dafny.Array
				_ = _1072_v23
				var _nwElement0_36 D12 = _1071_v22
				_ = _nwElement0_36
				var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(18))
				_ = _nw174
				(_nw174).ArraySet1(_nwElement0_36, 0)
				(_nw174).ArraySet1(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, p2)), 1)
				(_nw174).ArraySet1(_1071_v22, 2)
				(_nw174).ArraySet1(_1071_v22, 3)
				(_nw174).ArraySet1(_1071_v22, 4)
				(_nw174).ArraySet1(_1071_v22, 5)
				(_nw174).ArraySet1(_1071_v22, 6)
				(_nw174).ArraySet1((func() D12 {
					if _1049_v0 {
						return _1071_v22
					}
					return func(_pat_let42_0 D12) D12 {
						return func(_1073_dt__update__tmp_h0 D12) D12 {
							return func(_pat_let43_0 _dafny.Map) D12 {
								return func(_1074_dt__update_hcf35_h0 _dafny.Map) D12 {
									return Companion_D12_.Create_DC26_(_1074_dt__update_hcf35_h0)
								}(_pat_let43_0)
							}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv51, _pat_let_tv52))
						}(_pat_let42_0)
					}(_1071_v22)
				})(), 7)
				(_nw174).ArraySet1(func(_pat_let44_0 D12) D12 {
					return func(_1075_dt__update__tmp_h1 D12) D12 {
						return func(_pat_let45_0 _dafny.Map) D12 {
							return func(_1076_dt__update_hcf35_h1 _dafny.Map) D12 {
								return Companion_D12_.Create_DC26_(_1076_dt__update_hcf35_h1)
							}(_pat_let45_0)
						}(_pat_let_tv53)
					}(_pat_let44_0)
				}(Companion_D12_.Create_DC26_(_1070_v21)), 8)
				(_nw174).ArraySet1(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, p2)), 9)
				(_nw174).ArraySet1(Companion_Default___.Fm39(p2, _1049_v0, _1049_v0, globalState), 10)
				(_nw174).ArraySet1(_1071_v22, 11)
				(_nw174).ArraySet1(_1071_v22, 12)
				(_nw174).ArraySet1(_1071_v22, 13)
				(_nw174).ArraySet1(_1071_v22, 14)
				(_nw174).ArraySet1(Companion_Default___.Fm39(_dafny.IntOfInt64(616), _1049_v0, _1049_v0, globalState), 15)
				(_nw174).ArraySet1(_1071_v22, 16)
				(_nw174).ArraySet1(_1071_v22, 17)
				_1072_v23 = _nw174
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_1072_v23), 0))
				_ = _index202
				(_1072_v23).ArraySet1(_1071_v22, (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_1072_v23), 0))
				_ = _index203
				var _rhs191 _dafny.Int = (p0).Plus(_dafny.IntOfUint32((_1051_v2).Cardinality()))
				_ = _rhs191
				var _rhs192 D12 = Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, p2))
				_ = _rhs192
				var _rhs193 _dafny.Int = _dafny.IntOfUint32((_1051_v2).Cardinality())
				_ = _rhs193
				var _lhs153 *GlobalState = globalState
				_ = _lhs153
				var _lhs154 _dafny.Array = _1072_v23
				_ = _lhs154
				var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_1072_v23), 0))
				_ = _lhs155
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				_lhs153.F3 = _rhs191
				(_lhs154).ArraySet1(_rhs192, (_lhs155).Int())
				_lhs156.F8 = _rhs193
			}
			_1049_v0 = _1049_v0
			var _1077_v24 _dafny.MultiSet
			_ = _1077_v24
			_1077_v24 = _dafny.MultiSetOf(false, _1049_v0, _1049_v0, _1049_v0)
			var _1078_v25 _dafny.Map
			_ = _1078_v25
			_1078_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm28(p0, globalState), (_1077_v24).IsDisjointFrom(_1077_v24))
			var _1079_v26 _dafny.Map
			_ = _1079_v26
			_1079_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _dafny.IntOfInt64(373))
			var _1080_v27 _dafny.Array
			_ = _1080_v27
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw175
			_1080_v27 = _nw175
			var _1081_v28 _dafny.Set
			_ = _1081_v28
			_1081_v28 = _dafny.SetOf(_1080_v27, _1080_v27)
			var _1082_v29 _dafny.Set
			_ = _1082_v29
			_1082_v29 = _dafny.SetOf(_dafny.IntOfInt64(869), _dafny.IntOfInt64(951), p2, (_1079_v26).Cardinality(), (_1081_v28).Cardinality())
			if (func() bool {
				if (_1078_v25).Contains(_1082_v29) {
					return (_1078_v25).Get(_1082_v29).(bool)
				}
				return _1049_v0
			})() {
				var _1083_v30 *C2
				_ = _1083_v30
				var _nw176 *C2 = New_C2_()
				_ = _nw176
				_nw176.Ctor__()
				_1083_v30 = _nw176
				var _1084_v31 _dafny.Sequence
				_ = _1084_v31
				_1084_v31 = _dafny.SeqOf(_dafny.IntOfInt64(659), p0)
				var _1085_v32 _dafny.MultiSet
				_ = _1085_v32
				_1085_v32 = _dafny.MultiSetOf(p0)
				_1049_v0 = (_this).Fm3(_dafny.Companion_Sequence_.Concatenate(_1084_v31, _1084_v31), _1085_v32, _1049_v0, globalState)
				var _1086_v33 _dafny.Set
				_ = _1086_v33
				_1086_v33 = _dafny.SetOf(_1049_v0, _1049_v0)
				var _1087_v34 D9
				_ = _1087_v34
				_1087_v34 = Companion_D9_.Create_DC20_(_1086_v33)
				var _rhs194 bool = (_1049_v0) || (_1049_v0)
				_ = _rhs194
				var _rhs195 bool = ((_1087_v34).Dtor_cf25()).IsProperSubsetOf(_1086_v33)
				_ = _rhs195
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				_1049_v0 = _rhs194
				_lhs157.F4 = _rhs195
				var _1088_v35 _dafny.MultiSet
				_ = _1088_v35
				_1088_v35 = _1085_v32
				var _1089_v36 _dafny.Array
				_ = _1089_v36
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw177
				_1089_v36 = _nw177
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))
				_ = _index204
				(_1089_v36).ArraySet1((_1082_v29).IsSubsetOf(_1082_v29), (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))
				_ = _index205
				var _rhs196 _dafny.MultiSet = _1088_v35
				_ = _rhs196
				var _rhs197 bool = false
				_ = _rhs197
				var _rhs198 bool = _1049_v0
				_ = _rhs198
				var _rhs199 bool = _1049_v0
				_ = _rhs199
				var _lhs158 _dafny.Array = _1089_v36
				_ = _lhs158
				var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))
				_ = _lhs159
				_1088_v35 = _rhs196
				_1049_v0 = _rhs197
				_1049_v0 = _rhs198
				(_lhs158).ArraySet1(_rhs199, (_lhs159).Int())
				var _1090_v37 _dafny.Sequence
				_ = _1090_v37
				_1090_v37 = _dafny.SeqOf((_1089_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))).Int()).(bool))
				var _1091_v38 _dafny.Map
				_ = _1091_v38
				_1091_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				_1090_v37 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1049_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1090_v37, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm1(p0, _1050_v1, (_1089_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))).Int()).(bool), _1091_v38, globalState)).Cardinality()), _dafny.IntOfUint32((_1090_v37).Cardinality()))).Uint32(), (_1089_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1089_v36), 0))).Int()).(bool)), Companion_Default___.Fm31(true, false, _1049_v0, globalState)))
			} else {
				_1049_v0 = (_1077_v24).IsDisjointFrom((_1077_v24).Difference(_1077_v24))
				var _1092_v39 _dafny.Array
				_ = _1092_v39
				var _nw178 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw178
				_1092_v39 = _nw178
				var _1093_v40 _dafny.Map
				_ = _1093_v40
				_1093_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1092_v39)
				var _1094_v41 _dafny.Sequence
				_ = _1094_v41
				_1094_v41 = _dafny.SeqOf(_1092_v39)
				_1093_v40 = (_1093_v40).Update(_1049_v0, (_1094_v41).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1094_v41).Cardinality()))).Uint32()).(_dafny.Array))
				(globalState).F4 = _1049_v0
				_1049_v0 = (_dafny.IntOfInt64(66)).Cmp(p2) != 0
				var _1095_v42 _dafny.Set
				_ = _1095_v42
				_1095_v42 = _dafny.SetOf(_1049_v0)
				var _1096_v43 _dafny.Array
				_ = _1096_v43
				var _nwElement0_37 _dafny.Set = _1095_v42
				_ = _nwElement0_37
				var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(3))
				_ = _nw179
				(_nw179).ArraySet1(_nwElement0_37, 0)
				(_nw179).ArraySet1(_1095_v42, 1)
				(_nw179).ArraySet1(_dafny.SetOf(_1049_v0), 2)
				_1096_v43 = _nw179
				var _1097_v44 D4
				_ = _1097_v44
				_1097_v44 = Companion_D4_.Create_DC11_(true, _1049_v0, _dafny.IntOfInt64(-888), _1051_v2)
				var _1098_v45 _dafny.Map
				_ = _1098_v45
				_1098_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1096_v43, _dafny.IntOfUint32((_dafny.SeqOf(false, (_1097_v44).Dtor_cf14(), _1049_v0, _1049_v0)).Cardinality()))
				_1098_v45 = (_1098_v45).Update(_1096_v43, p2)
			}
			(globalState).F3 = p2
			var _1099_v46 D5
			_ = _1099_v46
			_1099_v46 = Companion_D5_.Create_DC15_()
			var _source19 D5 = _1099_v46
			_ = _source19
			if _source19.Is_DC14() {
				var _1100___mcc_h0 _dafny.Int = _source19.Get_().(D5_DC14).Cf20
				_ = _1100___mcc_h0
				var _1101___mcc_h1 bool = _source19.Get_().(D5_DC14).Cf21
				_ = _1101___mcc_h1
				var _1102_cf21 bool = _1101___mcc_h1
				_ = _1102_cf21
				var _1103_cf20 _dafny.Int = _1100___mcc_h0
				_ = _1103_cf20
				var _1104_v47 _dafny.Map
				_ = _1104_v47
				_1104_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1102_cf21, _1049_v0)
				var _1105_v48 _dafny.Sequence
				_ = _1105_v48
				_1105_v48 = _dafny.SeqOf(p2, (_dafny.Zero).Minus(p2))
				var _1106_v50 D12
				_ = _1106_v50
				_1106_v50 = Companion_D12_.Create_DC27_(_dafny.IntOfUint32((_1051_v2).Cardinality()), (_1104_v47).Cardinality(), p0, (func() _dafny.Set {
					var _coll78 = _dafny.NewBuilder()
					_ = _coll78
					for _iter79 := _dafny.Iterate((_1105_v48).Elements()); ; {
						_compr_78, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _1107_v49 _dafny.Int
						_1107_v49 = interface{}(_compr_78).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1105_v48, _1107_v49) {
							_coll78.Add((_1107_v49).Plus((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality()))
						}
					}
					return _coll78.ToSet()
				}()).Cardinality())
				var _1108_v51 D12
				_ = _1108_v51
				_1108_v51 = Companion_D12_.Create_DC28_(_1106_v50)
				var _1109_v52 _dafny.Set
				_ = _1109_v52
				_1109_v52 = _dafny.SetOf(_1102_cf21)
				var _rhs200 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, p0)).Plus(_1103_cf20)
				_ = _rhs200
				var _rhs201 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1105_v48).Cardinality()), Companion_Default___.SafeDivisionInt((_1109_v52).Cardinality(), (_1077_v24).Cardinality()))
				_ = _rhs201
				var _rhs202 _dafny.MultiSet = Companion_Default___.Fm0(p2, globalState)
				_ = _rhs202
				var _rhs203 _dafny.Sequence = (func() _dafny.Sequence {
					if ((_dafny.Zero).Minus((func() _dafny.Int {
						if (_1079_v26).Contains(_1049_v0) {
							return (_1079_v26).Get(_1049_v0).(_dafny.Int)
						}
						return _dafny.IntOfInt64(461)
					})())).Cmp(_dafny.IntOfInt64(15)) > 0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg47 _dafny.Int) interface{} {
								return coer47(arg47)
							}
						}((func(_1110_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1111_i0 _dafny.Int) _dafny.CodePoint {
								return _1110_v1
							}
						})(_1050_v1)))
					}
					return _1051_v2
				})()
				_ = _rhs203
				var _rhs204 D12 = _1108_v51
				_ = _rhs204
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				_lhs160.F3 = _rhs200
				_lhs161.F8 = _rhs201
				_1077_v24 = _rhs202
				_1051_v2 = _rhs203
				_1108_v51 = _rhs204
				_1104_v47 = (_1104_v47).Update(!(_1102_cf21) || (_1102_cf21), _1049_v0)
				(globalState).F8 = Companion_Default___.SafeModuloInt(_1103_cf20, Companion_Default___.SafeModuloInt(_1103_cf20, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()))))
				var _1112_v53 _dafny.Array
				_ = _1112_v53
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw180
				_1112_v53 = _nw180
				var _1113_v54 _dafny.Map
				_ = _1113_v54
				_1113_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _1112_v53)
				_1113_v54 = (_1113_v54).Update(_1049_v0, _1112_v53)
			} else if _source19.Is_DC15() {
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_1080_v27), 0))
				_ = _index206
				(_1080_v27).ArraySet1(p2, (_index206).Int())
				var _1114_v55 _dafny.Array
				_ = _1114_v55
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw181
				_1114_v55 = _nw181
				var _1115_v56 _dafny.Sequence
				_ = _1115_v56
				_1115_v56 = _dafny.SeqOf(_1114_v55)
				var _1116_v57 _dafny.Sequence
				_ = _1116_v57
				_1116_v57 = _dafny.SeqOf(_1114_v55, _1114_v55)
				var _1117_v58 D2
				_ = _1117_v58
				_1117_v58 = Companion_D2_.Create_DC7_(_1050_v1, (_1049_v0) || (_1049_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1115_v56, (_1116_v57))).Cardinality()))
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_1080_v27), 0))
				_ = _index207
				var _rhs205 _dafny.Int = p0
				_ = _rhs205
				var _rhs206 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1051_v2, _dafny.Companion_Sequence_.Concatenate(_1051_v2, _1051_v2))
				_ = _rhs206
				var _rhs207 D2 = (func() D2 {
					if _1049_v0 {
						return Companion_Default___.Fm40(!(_1049_v0), globalState)
					}
					return _1117_v58
				})()
				_ = _rhs207
				var _lhs162 _dafny.Array = _1080_v27
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_1080_v27), 0))
				_ = _lhs163
				(_lhs162).ArraySet1(_rhs205, (_lhs163).Int())
				_1051_v2 = _rhs206
				_1117_v58 = _rhs207
				var _1118_v59 _dafny.Sequence
				_ = _1118_v59
				_1118_v59 = _dafny.SeqOf((_1080_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_1080_v27), 0))).Int()).(_dafny.Int))
				var _1119_v60 _dafny.Array
				_ = _1119_v60
				var _nwElement0_38 _dafny.Sequence = _1118_v59
				_ = _nwElement0_38
				var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(3))
				_ = _nw182
				(_nw182).ArraySet1(_nwElement0_38, 0)
				(_nw182).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_1120_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1121_i1 _dafny.Int) _dafny.Int {
						return _1120_p0
					}
				})(p0))), 1)
				(_nw182).ArraySet1(_1118_v59, 2)
				_1119_v60 = _nw182
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_1119_v60), 0))
				_ = _index208
				(_1119_v60).ArraySet1(_1118_v59, (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_1119_v60), 0))
				_ = _index209
				(_1119_v60).ArraySet1(_1118_v59, (_index209).Int())
				var _1122_v61 _dafny.Map
				_ = _1122_v61
				_1122_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _1118_v59)
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_1119_v60), 0))
				_ = _index210
				var _rhs208 bool = (p0).Cmp((_1080_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_1080_v27), 0))).Int()).(_dafny.Int)) >= 0
				_ = _rhs208
				var _rhs209 _dafny.CodePoint = _1050_v1
				_ = _rhs209
				var _rhs210 _dafny.Int = p0
				_ = _rhs210
				var _rhs211 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1122_v61).Contains(_1049_v0) {
						return (_1122_v61).Get(_1049_v0).(_dafny.Sequence)
					}
					return (_1119_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_1119_v60), 0))).Int()).(_dafny.Sequence)
				})()
				_ = _rhs211
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				var _lhs166 _dafny.Array = _1119_v60
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_1119_v60), 0))
				_ = _lhs167
				_lhs164.F4 = _rhs208
				_1050_v1 = _rhs209
				_lhs165.F3 = _rhs210
				(_lhs166).ArraySet1(_rhs211, (_lhs167).Int())
				var _1123_v62 *C1
				_ = _1123_v62
				var _nw183 *C1 = New_C1_()
				_ = _nw183
				_nw183.Ctor__()
				_1123_v62 = _nw183
			} else {
				var _1124___mcc_h2 _dafny.Map = _source19.Get_().(D5_DC13).Cf19
				_ = _1124___mcc_h2
				var _1125_cf19 _dafny.Map = _1124___mcc_h2
				_ = _1125_cf19
				var _1126_v63 _dafny.Array
				_ = _1126_v63
				var _nw184 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(13))
				_ = _nw184
				_1126_v63 = _nw184
				var _1127_v64 D12
				_ = _1127_v64
				_1127_v64 = Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0))
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_1126_v63), 0))
				_ = _index211
				(_1126_v63).ArraySet1(_1127_v64, (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_1126_v63), 0))
				_ = _index212
				(_1126_v63).ArraySet1(_1127_v64, (_index212).Int())
				_1049_v0 = _1049_v0
				var _1128_v65 _dafny.Array
				_ = _1128_v65
				var _nw185 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw185
				_1128_v65 = _nw185
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_1128_v65), 0))
				_ = _index213
				(_1128_v65).ArraySet1(_1049_v0, (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_1128_v65), 0))
				_ = _index214
				(_1128_v65).ArraySet1((_1049_v0) && (_1049_v0), (_index214).Int())
				var _rhs212 _dafny.Int = p0
				_ = _rhs212
				var _rhs213 bool = (_1128_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_1128_v65), 0))).Int()).(bool)
				_ = _rhs213
				var _rhs214 bool = ((p1).Cardinality()).Cmp(p0) <= 0
				_ = _rhs214
				var _rhs215 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _rhs215
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				var _lhs169 *GlobalState = globalState
				_ = _lhs169
				_lhs168.F3 = _rhs212
				_1049_v0 = _rhs213
				_1049_v0 = _rhs214
				_lhs169.F8 = _rhs215
			}
		} else {
			var _1129_v66 _dafny.Map
			_ = _1129_v66
			_1129_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1049_v0), p2)
			_1129_v66 = (_1129_v66).Update(_1049_v0, p0)
			var _1130_v67 _dafny.Set
			_ = _1130_v67
			_1130_v67 = _dafny.SetOf(_1049_v0, !((func() bool {
				if (p1).Contains(p0) {
					return (p1).Get(p0).(bool)
				}
				return _1049_v0
			})()), false, _1049_v0)
			var _1131_v68 _dafny.Map
			_ = _1131_v68
			_1131_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1049_v0, _1049_v0, _1049_v0, _1049_v0), (_1130_v67).IsProperSubsetOf(_1130_v67))
			var _1132_v69 D9
			_ = _1132_v69
			_1132_v69 = Companion_D9_.Create_DC20_(_1130_v67)
			_1131_v68 = (_1131_v68).Update((_1132_v69).Dtor_cf25(), _1049_v0)
			var _1133_v70 _dafny.Map
			_ = _1133_v70
			_1133_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1051_v2).Cardinality()), (_this).Fm15(p0, globalState))
			var _1134_v71 _dafny.Sequence
			_ = _1134_v71
			_1134_v71 = _dafny.SeqOf(p0)
			var _1135_v72 _dafny.Array
			_ = _1135_v72
			var _nw186 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(14))
			_ = _nw186
			_1135_v72 = _nw186
			var _1136_v73 _dafny.Set
			_ = _1136_v73
			_1136_v73 = _dafny.SetOf(_1135_v72)
			var _1137_v74 _dafny.Array
			_ = _1137_v74
			var _nwElement0_39 _dafny.Int = (_dafny.SetOf(_1049_v0)).Cardinality()
			_ = _nwElement0_39
			var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(25))
			_ = _nw187
			(_nw187).ArraySet1(_nwElement0_39, 0)
			(_nw187).ArraySet1(p0, 1)
			(_nw187).ArraySet1(p0, 2)
			(_nw187).ArraySet1(((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)).Cardinality())).Cardinality()).Plus(p2), 3)
			(_nw187).ArraySet1((_dafny.IntOfInt64(-628)).Plus(p2), 4)
			(_nw187).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p2), p2)).Cardinality()), 5)
			(_nw187).ArraySet1(p0, 6)
			(_nw187).ArraySet1(p0, 7)
			(_nw187).ArraySet1(p0, 8)
			(_nw187).ArraySet1((_dafny.Zero).Minus(p2), 9)
			(_nw187).ArraySet1((_dafny.SetOf(_1050_v1, _1050_v1, _dafny.CodePoint('r'))).Cardinality(), 10)
			(_nw187).ArraySet1((func() _dafny.Int {
				if (_1133_v70).Contains(p0) {
					return (_1133_v70).Get(p0).(_dafny.Int)
				}
				return p2
			})(), 11)
			(_nw187).ArraySet1(_dafny.IntOfInt64(-132), 12)
			(_nw187).ArraySet1(p0, 13)
			(_nw187).ArraySet1(p2, 14)
			(_nw187).ArraySet1((p0).Times(p2), 15)
			(_nw187).ArraySet1(p0, 16)
			(_nw187).ArraySet1((func() _dafny.Int {
				if _1049_v0 {
					return p2
				}
				return _dafny.IntOfInt64(319)
			})(), 17)
			(_nw187).ArraySet1(_dafny.IntOfInt64(607), 18)
			(_nw187).ArraySet1((_dafny.Zero).Minus((p0).Plus(p0)), 19)
			(_nw187).ArraySet1(p2, 20)
			(_nw187).ArraySet1(p0, 21)
			(_nw187).ArraySet1(_dafny.IntOfUint32((_1134_v71).Cardinality()), 22)
			(_nw187).ArraySet1((_1136_v73).Cardinality(), 23)
			(_nw187).ArraySet1(p2, 24)
			_1137_v74 = _nw187
			_1137_v74 = _1137_v74
			var _1138_v75 _dafny.Array
			_ = _1138_v75
			var _nw188 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw188
			_1138_v75 = _nw188
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))
			_ = _index215
			(_1138_v75).ArraySet1(_1049_v0, (_index215).Int())
			var _1139_v76 _dafny.MultiSet
			_ = _1139_v76
			_1139_v76 = _dafny.MultiSetOf(_1135_v72)
			var _1140_v77 _dafny.Set
			_ = _1140_v77
			_1140_v77 = _dafny.SetOf(p0)
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))
			_ = _index216
			(_1138_v75).ArraySet1(!(_1140_v77).Contains(((_1139_v76).Intersection((_1139_v76))).Cardinality()), (_index216).Int())
			if _1049_v0 {
				var _1141_v78 *C0
				_ = _1141_v78
				var _nw189 *C0 = New_C0_()
				_ = _nw189
				_nw189.Ctor__()
				_1141_v78 = _nw189
				_1133_v70 = (_1133_v70).Update(((_1130_v67).Cardinality()).Plus(p2), p0)
				var _1142_v79 _dafny.Sequence
				_ = _1142_v79
				_1142_v79 = _dafny.SeqOf(false)
				_1049_v0 = (_1142_v79).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1142_v79).Cardinality()))).Uint32()).(bool)
				(globalState).F4 = (_1142_v79).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1142_v79).Cardinality()))).Uint32()).(bool)
				var _1143_v80 _dafny.Map
				_ = _1143_v80
				_1143_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool), (_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool))
				var _rhs216 _dafny.Int = p2
				_ = _rhs216
				var _rhs217 bool = (_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool)
				_ = _rhs217
				var _rhs218 _dafny.Set = (Companion_D13_.Create_DC30_((func() bool {
					if (_1143_v80).Contains(_1049_v0) {
						return (_1143_v80).Get(_1049_v0).(bool)
					}
					return (_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool)
				})(), _1049_v0, _1140_v77, _1141_v78)).Dtor_cf44()
				_ = _rhs218
				var _lhs170 *GlobalState = globalState
				_ = _lhs170
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				_lhs170.F3 = _rhs216
				_lhs171.F4 = _rhs217
				_1140_v77 = _rhs218
			} else {
				_1133_v70 = (_1133_v70).Merge(_1133_v70)
				var _1144_v81 _dafny.MultiSet
				_ = _1144_v81
				_1144_v81 = Companion_Default___.Fm0((_dafny.Zero).Minus(p2), globalState)
				var _1145_v82 _dafny.MultiSet
				_ = _1145_v82
				_1145_v82 = _dafny.MultiSetOf((_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool))
				_1144_v81 = (_1145_v82).Difference(_1145_v82)
				var _1146_v83 *C2
				_ = _1146_v83
				var _nw190 *C2 = New_C2_()
				_ = _nw190
				_nw190.Ctor__()
				_1146_v83 = _nw190
				var _1147_v84 _dafny.Sequence
				_ = _1147_v84
				_1147_v84 = _dafny.SeqOf((_this).Fm16(false, globalState))
				var _1148_v85 _dafny.Sequence
				_ = _1148_v85
				_1148_v85 = _dafny.SeqOf(_1049_v0, (func() bool {
					if (_1147_v84).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm20((_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_1147_v84).Cardinality()))).Uint32()).(bool) {
						return true
					}
					return (_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool)
				})(), (_1138_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1138_v75), 0))).Int()).(bool))
				_1049_v0 = (_1147_v84).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1147_v84).Cardinality()))).Uint32()).(bool)
				_1051_v2 = _1051_v2
			}
		}
		var _1149_v86 _dafny.Sequence
		_ = _1149_v86
		_1149_v86 = _dafny.SeqOf(_1049_v0)
		var _pat_let_tv54 = _1051_v2
		_ = _pat_let_tv54
		var _pat_let_tv55 = _1149_v86
		_ = _pat_let_tv55
		var _source20 D4 = func(_pat_let46_0 D4) D4 {
			return func(_1150_dt__update__tmp_h2 D4) D4 {
				return func(_pat_let47_0 _dafny.Sequence) D4 {
					return func(_1151_dt__update_hcf17_h0 _dafny.Sequence) D4 {
						return func(_pat_let48_0 _dafny.Int) D4 {
							return func(_1152_dt__update_hcf16_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC11_((_1150_dt__update__tmp_h2).Dtor_cf14(), (_1150_dt__update__tmp_h2).Dtor_cf15(), _1152_dt__update_hcf16_h0, _1151_dt__update_hcf17_h0)
							}(_pat_let48_0)
						}(_dafny.IntOfUint32((_pat_let_tv55).Cardinality()))
					}(_pat_let47_0)
				}(_pat_let_tv54)
			}(_pat_let46_0)
		}(Companion_D4_.Create_DC11_(_1049_v0, _1049_v0, p0, _1051_v2))
		_ = _source20
		if _source20.Is_DC11() {
			var _1153___mcc_h3 bool = _source20.Get_().(D4_DC11).Cf14
			_ = _1153___mcc_h3
			var _1154___mcc_h4 bool = _source20.Get_().(D4_DC11).Cf15
			_ = _1154___mcc_h4
			var _1155___mcc_h5 _dafny.Int = _source20.Get_().(D4_DC11).Cf16
			_ = _1155___mcc_h5
			var _1156___mcc_h6 _dafny.Sequence = _source20.Get_().(D4_DC11).Cf17
			_ = _1156___mcc_h6
			var _1157_cf17 _dafny.Sequence = _1156___mcc_h6
			_ = _1157_cf17
			var _1158_cf16 _dafny.Int = _1155___mcc_h5
			_ = _1158_cf16
			var _1159_cf15 bool = _1154___mcc_h4
			_ = _1159_cf15
			var _1160_cf14 bool = _1153___mcc_h3
			_ = _1160_cf14
			var _1161_v87 _dafny.Array
			_ = _1161_v87
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_27
			var _nw191 _dafny.Array
			_ = _nw191
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw191 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Sequence = (func(_1162_cf17 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1163_i2 _dafny.Int) _dafny.Sequence {
						return _1162_cf17
					}
				})(_1157_cf17)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw191 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw191).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw191).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1161_v87 = _nw191
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1161_v87), 0))
			_ = _index217
			(_1161_v87).ArraySet1(_1051_v2, (_index217).Int())
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1161_v87), 0))
			_ = _index218
			(_1161_v87).ArraySet1(Companion_Default___.Fm25(globalState), (_index218).Int())
			var _1164_v88 D9
			_ = _1164_v88
			_1164_v88 = Companion_D9_.Create_DC22_(_1149_v86)
			var _1165_v89 _dafny.Sequence
			_ = _1165_v89
			_1165_v89 = _dafny.SeqOf((_dafny.Zero).Minus(p0))
			var _1166_v90 _dafny.Map
			_ = _1166_v90
			_1166_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1158_cf16, p0)
			var _1167_v91 _dafny.Sequence
			_ = _1167_v91
			_1167_v91 = _dafny.SeqOf(_dafny.SeqOf(_1159_cf15))
			var _1168_v92 _dafny.MultiSet
			_ = _1168_v92
			_1168_v92 = _dafny.MultiSetOf(_1149_v86, (_1167_v91).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm20(_1160_cf14, globalState), _dafny.IntOfUint32((_1167_v91).Cardinality()))).Uint32()).(_dafny.Sequence), _1149_v86)
			var _pat_let_tv56 = _1166_v90
			_ = _pat_let_tv56
			var _rhs219 _dafny.Sequence = _1165_v89
			_ = _rhs219
			var _rhs220 _dafny.Int = Companion_Default___.Fm26(func(_pat_let49_0 D3) D3 {
				return func(_1169_dt__update__tmp_h3 D3) D3 {
					return func(_pat_let50_0 _dafny.Map) D3 {
						return func(_1170_dt__update_hcf12_h0 _dafny.Map) D3 {
							return Companion_D3_.Create_DC8_(_1170_dt__update_hcf12_h0)
						}(_pat_let50_0)
					}(_pat_let_tv56)
				}(_pat_let49_0)
			}(Companion_D3_.Create_DC8_(_1166_v90)), (_1158_cf16).Plus(p0), _1168_v92, globalState)
			_ = _rhs220
			var _rhs221 D9 = _1164_v88
			_ = _rhs221
			var _lhs172 *GlobalState = globalState
			_ = _lhs172
			var _lhs173 *GlobalState = globalState
			_ = _lhs173
			_lhs172.F0 = _rhs219
			_lhs173.F8 = _rhs220
			_1164_v88 = _rhs221
			var _1171_v93 _dafny.Array
			_ = _1171_v93
			var _nw192 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw192
			_1171_v93 = _nw192
			var _1172_v94 D4
			_ = _1172_v94
			_1172_v94 = Companion_D4_.Create_DC11_(_1160_cf14, _1159_cf15, _dafny.IntOfUint32((_1157_cf17).Cardinality()), _1051_v2)
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1171_v93), 0))
			_ = _index219
			(_1171_v93).ArraySet1((_1172_v94).Dtor_cf14(), (_index219).Int())
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1171_v93), 0))
			_ = _index220
			(_1171_v93).ArraySet1((func() bool {
				if _1049_v0 {
					return !(_1160_cf14)
				}
				return !(_dafny.Companion_Sequence_.IsPrefixOf(_1149_v86, _1149_v86))
			})(), (_index220).Int())
			var _1173_v95 _dafny.MultiSet
			_ = _1173_v95
			_1173_v95 = _dafny.MultiSetOf(_1049_v0)
			var _1174_v96 *C2
			_ = _1174_v96
			var _nw193 *C2 = New_C2_()
			_ = _nw193
			_nw193.Ctor__()
			_1174_v96 = _nw193
			var _1175_v97 D16
			_ = _1175_v97
			_1175_v97 = Companion_D16_.Create_DC34_(_1174_v96)
			var _1176_v98 _dafny.Map
			_ = _1176_v98
			_1176_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1175_v97).Dtor_cf49(), true)
			var _1177_v99 T0
			_ = _1177_v99
			var _nw194 *C1 = New_C1_()
			_ = _nw194
			_nw194.Ctor__()
			_1177_v99 = _nw194
			var _1178_v100 _dafny.MultiSet
			_ = _1178_v100
			_1178_v100 = _dafny.MultiSetOf(p0)
			var _1179_v101 _dafny.Map
			_ = _1179_v101
			_1179_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1177_v99, _1178_v100)
			var _1180_v102 _dafny.Array
			_ = _1180_v102
			var _nwElement0_40 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _nwElement0_40
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(27))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_40, 0)
			(_nw195).ArraySet1(p2, 1)
			(_nw195).ArraySet1((func() _dafny.Int {
				if (_1173_v95).Contains(false) {
					return (_1173_v95).Multiplicity(false)
				}
				return (_1176_v98).Cardinality()
			})(), 2)
			(_nw195).ArraySet1(_1158_cf16, 3)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality()), 4)
			(_nw195).ArraySet1(_dafny.IntOfInt64(141), 5)
			(_nw195).ArraySet1(_1158_cf16, 6)
			(_nw195).ArraySet1((p2).Times((_dafny.MultiSetOf(_1160_cf14)).Cardinality()), 7)
			(_nw195).ArraySet1((p2).Plus(p2), 8)
			(_nw195).ArraySet1(p0, 9)
			(_nw195).ArraySet1(p0, 10)
			(_nw195).ArraySet1((_1158_cf16).Plus(_1158_cf16), 11)
			(_nw195).ArraySet1(_dafny.IntOfInt64(-851), 12)
			(_nw195).ArraySet1(((func() _dafny.MultiSet {
				if (_1179_v101).Contains(_1177_v99) {
					return (_1179_v101).Get(_1177_v99).(_dafny.MultiSet)
				}
				return _1178_v100
			})()).Cardinality(), 13)
			(_nw195).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), 14)
			(_nw195).ArraySet1(_1158_cf16, 15)
			(_nw195).ArraySet1((_dafny.IntOfInt64(-421)).Minus(p2), 16)
			(_nw195).ArraySet1(p0, 17)
			(_nw195).ArraySet1(((_1166_v90).Cardinality()).Plus(_dafny.IntOfInt64(-24)), 18)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_1165_v89).Cardinality()), 19)
			(_nw195).ArraySet1((_dafny.Zero).Minus((_1158_cf16).Times(p0)), 20)
			(_nw195).ArraySet1(p0, 21)
			(_nw195).ArraySet1(p0, 22)
			(_nw195).ArraySet1(_1158_cf16, 23)
			(_nw195).ArraySet1(p2, 24)
			(_nw195).ArraySet1(p2, 25)
			(_nw195).ArraySet1(_1158_cf16, 26)
			_1180_v102 = _nw195
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1180_v102), 0))
			_ = _index221
			(_1180_v102).ArraySet1((_1165_v89).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1165_v89).Cardinality()))).Uint32()).(_dafny.Int), (_index221).Int())
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1180_v102), 0))
			_ = _index222
			(_1180_v102).ArraySet1((p2).Plus(p0), (_index222).Int())
		} else if _source20.Is_DC10() {
			var _1181___mcc_h7 _dafny.Sequence = _source20.Get_().(D4_DC10).Cf13
			_ = _1181___mcc_h7
			var _1182_cf13 _dafny.Sequence = _1181___mcc_h7
			_ = _1182_cf13
			var _1183_v103 *C0
			_ = _1183_v103
			var _nw196 *C0 = New_C0_()
			_ = _nw196
			_nw196.Ctor__()
			_1183_v103 = _nw196
			var _1184_v104 D0
			_ = _1184_v104
			_1184_v104 = Companion_D0_.Create_DC2_()
			var _source21 D0 = _1184_v104
			_ = _source21
			if _source21.Is_DC1() {
				var _1185___mcc_h9 _dafny.Int = _source21.Get_().(D0_DC1).Cf1
				_ = _1185___mcc_h9
				var _1186___mcc_h10 _dafny.Int = _source21.Get_().(D0_DC1).Cf2
				_ = _1186___mcc_h10
				var _1187_cf2 _dafny.Int = _1186___mcc_h10
				_ = _1187_cf2
				var _1188_cf1 _dafny.Int = _1185___mcc_h9
				_ = _1188_cf1
				var _1189_v105 _dafny.Array
				_ = _1189_v105
				var _nw197 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
				_ = _nw197
				_1189_v105 = _nw197
				var _1190_v106 _dafny.Set
				_ = _1190_v106
				_1190_v106 = _dafny.SetOf(_dafny.IntOfInt64(274), _1187_cf2, p2, _dafny.IntOfInt64(845), _1187_cf2)
				var _1191_v107 D13
				_ = _1191_v107
				_1191_v107 = Companion_D13_.Create_DC30_(false, _1049_v0, _1190_v106, _1183_v103)
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1189_v105), 0))
				_ = _index223
				(_1189_v105).ArraySet1(_1191_v107, (_index223).Int())
				var _1192_v108 _dafny.Sequence
				_ = _1192_v108
				_1192_v108 = _dafny.SeqOf(_1183_v103, _1183_v103)
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_1189_v105), 0))
				_ = _index224
				(_1189_v105).ArraySet1(Companion_D13_.Create_DC30_(_1049_v0, _1049_v0, _dafny.SetOf(_dafny.IntOfInt64(-260), p2), (_1192_v108).Select((Companion_Default___.SafeIndex(_1187_cf2, _dafny.IntOfUint32((_1192_v108).Cardinality()))).Uint32()).(*C0)), (_index224).Int())
				var _1193_v109 _dafny.Array
				_ = _1193_v109
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw198
				_1193_v109 = _nw198
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1193_v109), 0))
				_ = _index225
				(_1193_v109).ArraySet1(_1049_v0, (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1193_v109), 0))
				_ = _index226
				(_1193_v109).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1051_v2, _1050_v1), (_index226).Int())
				var _1194_v110 _dafny.Map
				_ = _1194_v110
				_1194_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(812))
				var _1195_v111 D3
				_ = _1195_v111
				_1195_v111 = Companion_D3_.Create_DC8_(_1194_v110)
				var _1196_v112 _dafny.MultiSet
				_ = _1196_v112
				_1196_v112 = _dafny.MultiSetOf(_1149_v86)
				var _pat_let_tv57 = _1194_v110
				_ = _pat_let_tv57
				(globalState).F8 = Companion_Default___.Fm26(func(_pat_let51_0 D3) D3 {
					return func(_1197_dt__update__tmp_h4 D3) D3 {
						return func(_pat_let52_0 _dafny.Map) D3 {
							return func(_1198_dt__update_hcf12_h1 _dafny.Map) D3 {
								return Companion_D3_.Create_DC8_(_1198_dt__update_hcf12_h1)
							}(_pat_let52_0)
						}(_pat_let_tv57)
					}(_pat_let51_0)
				}(_1195_v111), _dafny.IntOfUint32((_1051_v2).Cardinality()), _1196_v112, globalState)
				var _1199_v113 bool
				_ = _1199_v113
				_1199_v113 = !((_1193_v109).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1193_v109), 0))).Int()).(bool))
				var _1200_v114 _dafny.Map
				_ = _1200_v114
				_1200_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1199_v113, (_1193_v109).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1193_v109), 0))).Int()).(bool))
				var _1201_v115 _dafny.Sequence
				_ = _1201_v115
				_1201_v115 = _dafny.SeqOf(_dafny.IntOfUint32((_1051_v2).Cardinality()), p2, _1188_cf1)
				_1200_v114 = (_1200_v114).Update(_1199_v113, _dafny.Companion_Sequence_.IsPrefixOf(_1201_v115, _dafny.SeqOf(_dafny.IntOfUint32((_1182_cf13).Cardinality()), p2)))
			} else if _source21.Is_DC2() {
				var _1202_v116 _dafny.Array
				_ = _1202_v116
				var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw199
				_1202_v116 = _nw199
				var _1203_v117 _dafny.Array
				_ = _1203_v117
				var _nwElement0_41 _dafny.Array = _1202_v116
				_ = _nwElement0_41
				var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(5))
				_ = _nw200
				(_nw200).ArraySet1(_nwElement0_41, 0)
				(_nw200).ArraySet1(_1202_v116, 1)
				(_nw200).ArraySet1(_1202_v116, 2)
				(_nw200).ArraySet1(_1202_v116, 3)
				(_nw200).ArraySet1(_1202_v116, 4)
				_1203_v117 = _nw200
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1203_v117), 0))
				_ = _index227
				(_1203_v117).ArraySet1(_1202_v116, (_index227).Int())
				var _1204_v118 _dafny.Sequence
				_ = _1204_v118
				_1204_v118 = _dafny.SeqOf(p0)
				var _1205_v119 _dafny.Map
				_ = _1205_v119
				_1205_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p0).Times((_1204_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1051_v2).Cardinality()), _dafny.IntOfUint32((_1204_v118).Cardinality()))).Uint32()).(_dafny.Int)))
				var _1206_v120 _dafny.Array
				_ = _1206_v120
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_28
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Sequence = (func(_1207_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1208_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(855))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg49 _dafny.Int) interface{} {
									return coer49(arg49)
								}
							}((func(_1209_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1210_i4 _dafny.Int) _dafny.CodePoint {
									return _1209_v1
								}
							})(_1207_v1))), _dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('b'), _1207_v1, _1207_v1, _dafny.CodePoint('f')))
						}
					})(_1050_v1)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw201 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw201).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw201).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1206_v120 = _nw201
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1206_v120), 0))
				_ = _index228
				(_1206_v120).ArraySet1(_1051_v2, (_index228).Int())
				var _1211_v121 _dafny.Array
				_ = _1211_v121
				_1211_v121 = _1202_v116
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1203_v117), 0))
				_ = _index229
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1206_v120), 0))
				_ = _index230
				var _rhs222 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(p0, _1050_v1, _1049_v0, _1205_v119, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm1(p0, _1050_v1, _1049_v0, _1205_v119, globalState)).Cardinality()))).Uint32(), _dafny.IntOfInt64(973))).Cardinality()))
				_ = _rhs222
				var _rhs223 _dafny.Int = (_dafny.IntOfUint32((_1051_v2).Cardinality())).Minus(_dafny.IntOfUint32((_1149_v86).Cardinality()))
				_ = _rhs223
				var _rhs224 _dafny.Array = (_1211_v121)
				_ = _rhs224
				var _rhs225 _dafny.Map = _1205_v119
				_ = _rhs225
				var _rhs226 _dafny.Sequence = _1051_v2
				_ = _rhs226
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				var _lhs175 *GlobalState = globalState
				_ = _lhs175
				var _lhs176 _dafny.Array = _1203_v117
				_ = _lhs176
				var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1203_v117), 0))
				_ = _lhs177
				var _lhs178 _dafny.Array = _1206_v120
				_ = _lhs178
				var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1206_v120), 0))
				_ = _lhs179
				_lhs174.F3 = _rhs222
				_lhs175.F8 = _rhs223
				(_lhs176).ArraySet1(_rhs224, (_lhs177).Int())
				_1205_v119 = _rhs225
				(_lhs178).ArraySet1(_rhs226, (_lhs179).Int())
				var _1212_v122 _dafny.Array
				_ = _1212_v122
				var _nwElement0_42 bool = _1049_v0
				_ = _nwElement0_42
				var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(13))
				_ = _nw202
				(_nw202).ArraySet1(_nwElement0_42, 0)
				(_nw202).ArraySet1(_1049_v0, 1)
				(_nw202).ArraySet1(_1049_v0, 2)
				(_nw202).ArraySet1(_1049_v0, 3)
				(_nw202).ArraySet1(_1049_v0, 4)
				(_nw202).ArraySet1(_1049_v0, 5)
				(_nw202).ArraySet1(_1049_v0, 6)
				(_nw202).ArraySet1(_1049_v0, 7)
				(_nw202).ArraySet1(_1049_v0, 8)
				(_nw202).ArraySet1(_1049_v0, 9)
				(_nw202).ArraySet1((_this).Fm16(_1049_v0, globalState), 10)
				(_nw202).ArraySet1(_1049_v0, 11)
				(_nw202).ArraySet1(_1049_v0, 12)
				_1212_v122 = _nw202
				var _1213_v123 _dafny.Map
				_ = _1213_v123
				_1213_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1212_v122)
				_1213_v123 = (_1213_v123).Merge(_1213_v123)
				_1051_v2 = (_1206_v120).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1206_v120), 0))).Int()).(_dafny.Sequence)
				(globalState).F4 = _1049_v0
			} else if _source21.Is_DC0() {
				var _1214___mcc_h11 _dafny.Int = _source21.Get_().(D0_DC0).Cf0
				_ = _1214___mcc_h11
				var _1215_cf0 _dafny.Int = _1214___mcc_h11
				_ = _1215_cf0
				(globalState).F4 = _1049_v0
				var _1216_v124 _dafny.Array
				_ = _1216_v124
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_29
				var _nw203 _dafny.Array
				_ = _nw203
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw203 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Sequence = (func(_1217_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1218_i5 _dafny.Int) _dafny.Sequence {
							return _1217_v2
						}
					})(_1051_v2)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw203 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw203).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw203).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1216_v124 = _nw203
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1216_v124), 0))
				_ = _index231
				(_1216_v124).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rwhy"), (_index231).Int())
				var _1219_v125 _dafny.Map
				_ = _1219_v125
				_1219_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _1050_v1)
				var _1220_v126 _dafny.Map
				_ = _1220_v126
				_1220_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v2, !(_1219_v125).Contains(_1049_v0))
				var _1221_v127 _dafny.Array
				_ = _1221_v127
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw204
				_1221_v127 = _nw204
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))
				_ = _index232
				(_1221_v127).ArraySet1(p0, (_index232).Int())
				var _1222_v128 _dafny.Array
				_ = _1222_v128
				_1222_v128 = _1221_v127
				var _1223_v129 _dafny.Sequence
				_ = _1223_v129
				_1223_v129 = _dafny.SeqOf(_1222_v128)
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1216_v124), 0))
				_ = _index233
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))
				_ = _index234
				var _rhs227 _dafny.Sequence = _1051_v2
				_ = _rhs227
				var _rhs228 _dafny.Map = _1220_v126
				_ = _rhs228
				var _rhs229 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1149_v86, _dafny.Companion_Sequence_.Concatenate(_1182_cf13, Companion_Default___.Fm31(_1049_v0, _1049_v0, _1049_v0, globalState))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if _1049_v0 {
						return _1223_v129
					}
					return _1223_v129
				})()).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1149_v86, _dafny.Companion_Sequence_.Concatenate(_1182_cf13, Companion_Default___.Fm31(_1049_v0, _1049_v0, _1049_v0, globalState)))).Cardinality()))).Uint32(), _1049_v0)).Cardinality())
				_ = _rhs229
				var _rhs230 bool = _1049_v0
				_ = _rhs230
				var _lhs180 _dafny.Array = _1216_v124
				_ = _lhs180
				var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1216_v124), 0))
				_ = _lhs181
				var _lhs182 _dafny.Array = _1221_v127
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))
				_ = _lhs183
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				(_lhs180).ArraySet1(_rhs227, (_lhs181).Int())
				_1220_v126 = _rhs228
				(_lhs182).ArraySet1(_rhs229, (_lhs183).Int())
				_lhs184.F4 = _rhs230
				(globalState).F4 = _1049_v0
				var _1224_v130 _dafny.Sequence
				_ = _1224_v130
				_1224_v130 = _dafny.SeqOf(_1215_cf0, (_1221_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))).Int()).(_dafny.Int))
				var _1225_v131 _dafny.MultiSet
				_ = _1225_v131
				_1225_v131 = _dafny.MultiSetOf(_1215_cf0)
				var _1226_v132 _dafny.MultiSet
				_ = _1226_v132
				_1226_v132 = _dafny.MultiSetOf(_1225_v131)
				var _1227_v133 _dafny.Map
				_ = _1227_v133
				_1227_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-174))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_1228_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1229_i6 _dafny.Int) _dafny.Int {
						return _1228_p0
					}
				})(p0))), (_dafny.Zero).Minus(_1215_cf0))
				var _1230_v134 _dafny.Map
				_ = _1230_v134
				_1230_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1049_v0)
				var _1231_v135 _dafny.Map
				_ = _1231_v135
				_1231_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, (func() bool {
					if (_1230_v134).Contains(_1049_v0) {
						return (_1230_v134).Get(_1049_v0).(bool)
					}
					return (_this).Fm3(_1224_v130, _1225_v131, _1049_v0, globalState)
				})())
				var _1232_v136 _dafny.Map
				_ = _1232_v136
				_1232_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1231_v135).Cardinality())
				var _1233_v137 _dafny.Set
				_ = _1233_v137
				_1233_v137 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1215_cf0, p2)).Update((_1227_v133).Cardinality(), p0), _1232_v136)
				var _1234_v139 _dafny.Array
				_ = _1234_v139
				var _nwElement0_43 bool = _1049_v0
				_ = _nwElement0_43
				var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(20))
				_ = _nw205
				(_nw205).ArraySet1(_nwElement0_43, 0)
				(_nw205).ArraySet1(!(_1049_v0), 1)
				(_nw205).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1224_v130, _dafny.SeqOf((_1221_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))).Int()).(_dafny.Int))), 2)
				(_nw205).ArraySet1((_1226_v132).IsDisjointFrom(_1226_v132), 3)
				(_nw205).ArraySet1(_1049_v0, 4)
				(_nw205).ArraySet1(_1049_v0, 5)
				(_nw205).ArraySet1(_1049_v0, 6)
				(_nw205).ArraySet1((_dafny.SetOf(func() _dafny.Map {
					var _coll79 = _dafny.NewMapBuilder()
					_ = _coll79
					for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-765), _dafny.IntOfInt64(433))); ; {
						_compr_79, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _1235_v138 _dafny.Int
						_1235_v138 = interface{}(_compr_79).(_dafny.Int)
						if ((_dafny.IntOfInt64(-765)).Cmp(_1235_v138) <= 0) && ((_1235_v138).Cmp(_dafny.IntOfInt64(433)) < 0) {
							_coll79.Add((_1235_v138).Minus((_1221_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))).Int()).(_dafny.Int)), p0)
						}
					}
					return _coll79.ToMap()
				}(), _1232_v136)).IsSubsetOf(_1233_v137), 7)
				(_nw205).ArraySet1(true, 8)
				(_nw205).ArraySet1(!(!(_1049_v0)), 9)
				(_nw205).ArraySet1(_1049_v0, 10)
				(_nw205).ArraySet1(_1049_v0, 11)
				(_nw205).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_1215_cf0, p0)).Cardinality())).Cmp((_1221_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_1221_v127), 0))).Int()).(_dafny.Int)) <= 0, 12)
				(_nw205).ArraySet1(_1049_v0, 13)
				(_nw205).ArraySet1(true, 14)
				(_nw205).ArraySet1(!(!(_1049_v0)), 15)
				(_nw205).ArraySet1(_1049_v0, 16)
				(_nw205).ArraySet1(_1049_v0, 17)
				(_nw205).ArraySet1((_dafny.SetOf(_1215_cf0)).IsDisjointFrom(_dafny.SetOf(p2, p2, _dafny.IntOfUint32(((_1216_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1216_v124), 0))).Int()).(_dafny.Sequence)).Cardinality()))), 18)
				(_nw205).ArraySet1(_1049_v0, 19)
				_1234_v139 = _nw205
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1234_v139), 0))
				_ = _index235
				(_1234_v139).ArraySet1(_1049_v0, (_index235).Int())
				var _1236_v140 D12
				_ = _1236_v140
				_1236_v140 = Companion_D12_.Create_DC27_(p0, _1215_cf0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v0, _dafny.SetOf(p0, _1215_cf0, p2))).Cardinality(), p0)
				var _1237_v141 _dafny.Sequence
				_ = _1237_v141
				_1237_v141 = _dafny.SeqOf(_1236_v140)
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1234_v139), 0))
				_ = _index236
				(_1234_v139).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1237_v141, _dafny.Companion_Sequence_.Concatenate(_1237_v141, _1237_v141)), (_index236).Int())
			} else {
				var _1238___mcc_h12 D0 = _source21.Get_().(D0_DC3).Cf3
				_ = _1238___mcc_h12
				var _1239_cf3 D0 = _1238___mcc_h12
				_ = _1239_cf3
				var _1240_v142 _dafny.Array
				_ = _1240_v142
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw206
				_1240_v142 = _nw206
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))
				_ = _index237
				(_1240_v142).ArraySet1((Companion_D5_.Create_DC14_(p0, _1049_v0)).Dtor_cf21(), (_index237).Int())
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))
				_ = _index238
				(_1240_v142).ArraySet1(_1049_v0, (_index238).Int())
				_1049_v0 = (_1240_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))).Int()).(bool)
				var _1241_v143 _dafny.Set
				_ = _1241_v143
				_1241_v143 = _dafny.SetOf((_1240_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))).Int()).(bool), (_1240_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))).Int()).(bool))
				var _1242_v144 _dafny.Array
				_ = _1242_v144
				var _nwElement0_44 _dafny.Int = p0
				_ = _nwElement0_44
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(10))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_44, 0)
				(_nw207).ArraySet1(p2, 1)
				(_nw207).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(772))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_1243_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1244_i7 _dafny.Int) _dafny.CodePoint {
						return _1243_v1
					}
				})(_1050_v1)))).Cardinality()), 2)
				(_nw207).ArraySet1(p2, 3)
				(_nw207).ArraySet1((_dafny.SetOf(p0)).Cardinality(), 4)
				(_nw207).ArraySet1(p2, 5)
				(_nw207).ArraySet1(p2, 6)
				(_nw207).ArraySet1(_dafny.IntOfUint32((_1051_v2).Cardinality()), 7)
				(_nw207).ArraySet1(_dafny.IntOfUint32((_1182_cf13).Cardinality()), 8)
				(_nw207).ArraySet1((_1241_v143).Cardinality(), 9)
				_1242_v144 = _nw207
				var _1245_v145 _dafny.Map
				_ = _1245_v145
				_1245_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)
				var _1246_v146 _dafny.Sequence
				_ = _1246_v146
				_1246_v146 = _dafny.SeqOf(_1245_v145, _1245_v145)
				var _1247_v147 _dafny.Array
				_ = _1247_v147
				var _nwElement0_45 _dafny.Map = (func() _dafny.Map {
					if _1049_v0 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)
					}
					return _1245_v145
				})()
				_ = _nwElement0_45
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(26))
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_45, 0)
				(_nw208).ArraySet1(_1245_v145, 1)
				(_nw208).ArraySet1((_1245_v145).Merge((_1246_v146).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1246_v146).Cardinality()))).Uint32()).(_dafny.Map)), 2)
				(_nw208).ArraySet1(_1245_v145, 3)
				(_nw208).ArraySet1((func() _dafny.Map {
					if (_1240_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))).Int()).(bool) {
						return _1245_v145
					}
					return _1245_v145
				})(), 4)
				(_nw208).ArraySet1(_1245_v145, 5)
				(_nw208).ArraySet1(_1245_v145, 6)
				(_nw208).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)).Merge(_1245_v145), 7)
				(_nw208).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142), 8)
				(_nw208).ArraySet1(_1245_v145, 9)
				(_nw208).ArraySet1(((_1245_v145).Update(_1242_v144, _1240_v142)).Merge(_1245_v145), 10)
				(_nw208).ArraySet1(((_1245_v145).Update(_1242_v144, _1240_v142)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)), 11)
				(_nw208).ArraySet1((_1245_v145).Merge(_1245_v145), 12)
				(_nw208).ArraySet1(((_1245_v145).Update(_1242_v144, _1240_v142)).Merge(_1245_v145), 13)
				(_nw208).ArraySet1(_1245_v145, 14)
				(_nw208).ArraySet1(_1245_v145, 15)
				(_nw208).ArraySet1(_1245_v145, 16)
				(_nw208).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142), 17)
				(_nw208).ArraySet1(_1245_v145, 18)
				(_nw208).ArraySet1(_1245_v145, 19)
				(_nw208).ArraySet1(_1245_v145, 20)
				(_nw208).ArraySet1((_1245_v145).Merge(_1245_v145), 21)
				(_nw208).ArraySet1((_1245_v145).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)), 22)
				(_nw208).ArraySet1(_1245_v145, 23)
				(_nw208).ArraySet1(_1245_v145, 24)
				(_nw208).ArraySet1(_1245_v145, 25)
				_1247_v147 = _nw208
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1247_v147), 0))
				_ = _index239
				(_1247_v147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142), (_index239).Int())
				var _1248_v148 _dafny.MultiSet
				_ = _1248_v148
				_1248_v148 = _dafny.MultiSetOf((_1240_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))).Int()).(bool))
				var _1249_v149 _dafny.MultiSet
				_ = _1249_v149
				_1249_v149 = _dafny.MultiSetOf(p2, p2)
				var _1250_v150 _dafny.MultiSet
				_ = _1250_v150
				_1250_v150 = _1249_v149
				var _1251_v151 _dafny.Set
				_ = _1251_v151
				_1251_v151 = _dafny.SetOf(_1250_v150)
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))
				_ = _index240
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1247_v147), 0))
				_ = _index241
				var _rhs231 bool = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1248_v148).Cardinality(), (_1251_v151).Cardinality()))).Cmp(_dafny.IntOfInt64(632)) > 0
				_ = _rhs231
				var _rhs232 _dafny.Map = (_1245_v145).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v144, _1240_v142)))
				_ = _rhs232
				var _lhs185 _dafny.Array = _1240_v142
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1240_v142), 0))
				_ = _lhs186
				var _lhs187 _dafny.Array = _1247_v147
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1247_v147), 0))
				_ = _lhs188
				(_lhs185).ArraySet1(_rhs231, (_lhs186).Int())
				(_lhs187).ArraySet1(_rhs232, (_lhs188).Int())
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1242_v144), 0))
				_ = _index242
				(_1242_v144).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), (_index242).Int())
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1242_v144), 0))
				_ = _index243
				(_1242_v144).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1051_v2, _1051_v2), _1051_v2)).Cardinality()), (_index243).Int())
			}
			var _1252_v152 _dafny.MultiSet
			_ = _1252_v152
			_1252_v152 = _dafny.MultiSetOf((Companion_Default___.Fm27(_1049_v0, globalState)).Cardinality(), Companion_Default___.SafeModuloInt(p2, p2), ((_this).Fm15((p1).Cardinality(), globalState)).Times(p2), (p2).Minus(_dafny.IntOfInt64(603)), (p2).Times(_dafny.IntOfInt64(-108)))
			(globalState).F2 = _1252_v152
			var _1253_v153 _dafny.MultiSet
			_ = _1253_v153
			_1253_v153 = _dafny.MultiSetOf((_1252_v152).Difference(_1252_v152))
			var _1254_v154 _dafny.Array
			_ = _1254_v154
			var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw209
			_1254_v154 = _nw209
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1254_v154), 0))
			_ = _index244
			(_1254_v154).ArraySet1(p2, (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1254_v154), 0))
			_ = _index245
			var _rhs233 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if _1049_v0 {
					return p2
				}
				return (_this).Fm15(p2, globalState)
			})())
			_ = _rhs233
			var _rhs234 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1051_v2, _1051_v2)
			_ = _rhs234
			var _rhs235 _dafny.MultiSet = (_1253_v153).Update(_1252_v152, Companion_Default___.Abs(p2))
			_ = _rhs235
			var _rhs236 _dafny.Int = p2
			_ = _rhs236
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			var _lhs190 _dafny.Array = _1254_v154
			_ = _lhs190
			var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1254_v154), 0))
			_ = _lhs191
			_lhs189.F3 = _rhs233
			_1051_v2 = _rhs234
			_1253_v153 = _rhs235
			(_lhs190).ArraySet1(_rhs236, (_lhs191).Int())
		} else {
			var _1255___mcc_h8 D4 = _source20.Get_().(D4_DC12).Cf18
			_ = _1255___mcc_h8
			var _1256_cf18 D4 = _1255___mcc_h8
			_ = _1256_cf18
			(globalState).F3 = p0
			var _1257_v155 _dafny.Map
			_ = _1257_v155
			_1257_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p0)), _dafny.Companion_Sequence_.Concatenate(_1149_v86, _dafny.SeqOf(_1049_v0, _1049_v0)))
			_1257_v155 = (_1257_v155).Update(p2, _1149_v86)
			(globalState).F4 = _1049_v0
			var _1258_v156 *C2
			_ = _1258_v156
			var _nw210 *C2 = New_C2_()
			_ = _nw210
			_nw210.Ctor__()
			_1258_v156 = _nw210
		}
		var _1259_v157 _dafny.Array
		_ = _1259_v157
		var _nw211 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw211
		_1259_v157 = _nw211
		_1259_v157 = _1259_v157
		var _1260_i8 _dafny.Int
		_ = _1260_i8
		_1260_i8 = _dafny.Zero
		{
			for (func() bool {
				if (p1).Contains(p0) {
					return (p1).Get(p0).(bool)
				}
				return (_1149_v86).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1149_v86).Cardinality()))).Uint32()).(bool)
			})() {
				{
					if (_1260_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1260_i8 = (_1260_i8).Plus(_dafny.One)
					var _1261_v158 _dafny.Map
					_ = _1261_v158
					_1261_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v2, _dafny.IntOfUint32((_1051_v2).Cardinality()))
					_1261_v158 = (_1261_v158).Update(_1051_v2, (func() _dafny.Int {
						if _1049_v0 {
							return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("bly"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bly")).Cardinality()))).Uint32(), _1050_v1)).Cardinality())
						}
						return (_dafny.Zero).Minus(p2)
					})())
					var _1262_v159 _dafny.Set
					_ = _1262_v159
					_1262_v159 = _dafny.SetOf(p2, _dafny.IntOfInt64(-6), p0)
					var _1263_v160 _dafny.Map
					_ = _1263_v160
					_1263_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1050_v1, _1262_v159)
					_1263_v160 = (_1263_v160).Update(_1050_v1, (_1262_v159).Intersection(_dafny.SetOf(p2)))
					var _1264_v161 *C0
					_ = _1264_v161
					var _nw212 *C0 = New_C0_()
					_ = _nw212
					_nw212.Ctor__()
					_1264_v161 = _nw212
					(globalState).F4 = ((p1).Cardinality()).Cmp(p2) <= 0
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1265_v163 D13
		_ = _1265_v163
		_1265_v163 = Companion_D13_.Create_DC29_(func() _dafny.Set {
			var _coll80 = _dafny.NewBuilder()
			_ = _coll80
			for _iter81 := _dafny.Iterate((Companion_Default___.Fm41(_1149_v86, _1050_v1, p0, globalState)).Elements()); ; {
				_compr_80, _ok81 := _iter81()
				if !_ok81 {
					break
				}
				var _1266_v162 _dafny.Sequence
				_1266_v162 = interface{}(_compr_80).(_dafny.Sequence)
				if (Companion_Default___.Fm41(_1149_v86, _1050_v1, p0, globalState)).Contains(_1266_v162) {
					_coll80.Add(_1266_v162)
				}
			}
			return _coll80.ToSet()
		}())
		var _pat_let_tv58 = _1050_v1
		_ = _pat_let_tv58
		var _pat_let_tv59 = p2
		_ = _pat_let_tv59
		var _pat_let_tv60 = _1049_v0
		_ = _pat_let_tv60
		var _pat_let_tv61 = p0
		_ = _pat_let_tv61
		var _pat_let_tv62 = _1049_v0
		_ = _pat_let_tv62
		var _source22 D5 = func(_source23 D13) D5 {
			if _source23.Is_DC30() {
				var _1267___mcc_h16 bool = _source23.Get_().(D13_DC30).Cf42
				_ = _1267___mcc_h16
				var _1268___mcc_h17 bool = _source23.Get_().(D13_DC30).Cf43
				_ = _1268___mcc_h17
				var _1269___mcc_h18 _dafny.Set = _source23.Get_().(D13_DC30).Cf44
				_ = _1269___mcc_h18
				var _1270___mcc_h19 *C0 = _source23.Get_().(D13_DC30).Cf45
				_ = _1270___mcc_h19
				var _1271_cf45 *C0 = _1270___mcc_h19
				_ = _1271_cf45
				var _1272_cf44 _dafny.Set = _1269___mcc_h18
				_ = _1272_cf44
				var _1273_cf43 bool = _1268___mcc_h17
				_ = _1273_cf43
				var _1274_cf42 bool = _1267___mcc_h16
				_ = _1274_cf42
				return Companion_D5_.Create_DC14_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1275_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1276_i9 _dafny.Int) _dafny.CodePoint {
						return _1275_v1
					}
				})(_pat_let_tv58)))).Cardinality()), !(_1274_cf42))
			} else if _source23.Is_DC29() {
				var _1277___mcc_h20 _dafny.Set = _source23.Get_().(D13_DC29).Cf41
				_ = _1277___mcc_h20
				var _1278_cf41 _dafny.Set = _1277___mcc_h20
				_ = _1278_cf41
				return Companion_D5_.Create_DC14_(_pat_let_tv59, _pat_let_tv60)
			} else {
				var _1279___mcc_h21 D13 = _source23.Get_().(D13_DC31).Cf46
				_ = _1279___mcc_h21
				var _1280_cf46 D13 = _1279___mcc_h21
				_ = _1280_cf46
				return Companion_D5_.Create_DC14_(_pat_let_tv61, _pat_let_tv62)
			}
		}(_1265_v163)
		_ = _source22
		if _source22.Is_DC14() {
			var _1281___mcc_h13 _dafny.Int = _source22.Get_().(D5_DC14).Cf20
			_ = _1281___mcc_h13
			var _1282___mcc_h14 bool = _source22.Get_().(D5_DC14).Cf21
			_ = _1282___mcc_h14
			var _1283_cf21 bool = _1282___mcc_h14
			_ = _1283_cf21
			var _1284_cf20 _dafny.Int = _1281___mcc_h13
			_ = _1284_cf20
			_1049_v0 = _1283_cf21
			var _1285_v164 _dafny.Map
			_ = _1285_v164
			_1285_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(909), !(false) || (_1283_cf21))
			_1285_v164 = (_1285_v164).Update(_1284_cf20, false)
			var _1286_v165 _dafny.MultiSet
			_ = _1286_v165
			_1286_v165 = _dafny.MultiSetOf(_dafny.IntOfInt64(-720), p2, p0, _1284_cf20)
			(globalState).F2 = _1286_v165
			var _1287_v166 _dafny.Set
			_ = _1287_v166
			_1287_v166 = _dafny.SetOf((_this).Fm16(_1049_v0, globalState))
			var _1288_v167 _dafny.Sequence
			_ = _1288_v167
			_1288_v167 = _dafny.SeqOf(p2, p2, (_1287_v166).Cardinality())
			var _1289_v168 _dafny.MultiSet
			_ = _1289_v168
			_1289_v168 = _dafny.MultiSetOf((func() bool {
				if (_this).Fm3(_1288_v167, _1286_v165, _1049_v0, globalState) {
					return _1049_v0
				}
				return _1283_cf21
			})(), _1283_cf21, _1283_cf21, (_1149_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1149_v86).Cardinality()))).Uint32()).(bool), _1049_v0)
			var _1290_v169 D2
			_ = _1290_v169
			_1290_v169 = Companion_D2_.Create_DC7_(_1050_v1, false, p0)
			var _rhs237 _dafny.MultiSet = (_1289_v168).Difference(_1289_v168)
			_ = _rhs237
			var _rhs238 _dafny.Int = (_1290_v169).Dtor_cf11()
			_ = _rhs238
			var _rhs239 _dafny.Int = (_this).Fm15(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(824)), globalState)
			_ = _rhs239
			var _rhs240 _dafny.Set = (func() _dafny.Set {
				if _1049_v0 {
					return _1287_v166
				}
				return _1287_v166
			})()
			_ = _rhs240
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			_1289_v168 = _rhs237
			_lhs192.F3 = _rhs238
			_1284_cf20 = _rhs239
			_1287_v166 = _rhs240
		} else if _source22.Is_DC15() {
			var _1291_v170 *C0
			_ = _1291_v170
			var _nw213 *C0 = New_C0_()
			_ = _nw213
			_nw213.Ctor__()
			_1291_v170 = _nw213
			var _1292_v171 _dafny.Sequence
			_ = _1292_v171
			_1292_v171 = _dafny.SeqOf(p0, p0)
			var _1293_v172 _dafny.Map
			_ = _1293_v172
			_1293_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1149_v86)
			(globalState).F4 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1292_v171).Cardinality()), p2)).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((p0).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1293_v172).Contains((_1291_v170).Fm19(globalState)) {
					return (_1293_v172).Get((_1291_v170).Fm19(globalState)).(_dafny.Sequence)
				}
				return _1149_v86
			})()).Cardinality()))))) != 0
			var _1294_v173 D3
			_ = _1294_v173
			_1294_v173 = Companion_D3_.Create_DC8_((_this).Fm2(globalState))
			var _1295_v174 _dafny.MultiSet
			_ = _1295_v174
			_1295_v174 = _dafny.MultiSetOf(_1149_v86, _1149_v86)
			(globalState).F3 = (_dafny.Zero).Minus(Companion_Default___.Fm26(_1294_v173, p0, _1295_v174, globalState))
			var _1296_v175 *C2
			_ = _1296_v175
			var _nw214 *C2 = New_C2_()
			_ = _nw214
			_nw214.Ctor__()
			_1296_v175 = _nw214
		} else {
			var _1297___mcc_h15 _dafny.Map = _source22.Get_().(D5_DC13).Cf19
			_ = _1297___mcc_h15
			var _1298_cf19 _dafny.Map = _1297___mcc_h15
			_ = _1298_cf19
			(globalState).F4 = _dafny.Companion_Sequence_.Equal(_1051_v2, _dafny.UnicodeSeqOfUtf8Bytes("r"))
			var _1299_v176 _dafny.Array
			_ = _1299_v176
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(3))
			_ = _nw215
			_1299_v176 = _nw215
			var _1300_v177 _dafny.Map
			_ = _1300_v177
			_1300_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1299_v176, p0)
			_1300_v177 = ((_1300_v177).Merge(_1300_v177)).Merge(_1300_v177)
			var _1301_v178 _dafny.Array
			_ = _1301_v178
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_30
			var _nw216 _dafny.Array
			_ = _nw216
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw216 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Int = (func(_1302_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1303_i10 _dafny.Int) _dafny.Int {
						return (_1303_i10).Plus(_1302_p2)
					}
				})(p2)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw216 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw216).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw216).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1301_v178 = _nw216
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1301_v178), 0))
			_ = _index246
			(_1301_v178).ArraySet1(_dafny.IntOfInt64(899), (_index246).Int())
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1301_v178), 0))
			_ = _index247
			var _rhs241 bool = _1049_v0
			_ = _rhs241
			var _rhs242 _dafny.Array = _1301_v178
			_ = _rhs242
			var _rhs243 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if true {
					return p0
				}
				return p0
			})()))
			_ = _rhs243
			var _lhs193 *GlobalState = globalState
			_ = _lhs193
			var _lhs194 _dafny.Array = _1301_v178
			_ = _lhs194
			var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1301_v178), 0))
			_ = _lhs195
			_lhs193.F4 = _rhs241
			_1301_v178 = _rhs242
			(_lhs194).ArraySet1(_rhs243, (_lhs195).Int())
			var _1304_v179 _dafny.MultiSet
			_ = _1304_v179
			_1304_v179 = _dafny.MultiSetOf(_1049_v0, _1049_v0)
			var _1305_v180 _dafny.Map
			_ = _1305_v180
			_1305_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1050_v1, _1049_v0)
			var _1306_v181 _dafny.Sequence
			_ = _1306_v181
			_1306_v181 = _dafny.SeqOf(_1305_v180)
			var _1307_v182 _dafny.MultiSet
			_ = _1307_v182
			_1307_v182 = _dafny.MultiSetOf(_1049_v0, _1049_v0, _1049_v0, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1050_v1, _1049_v0)), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1304_v179).Contains(_1049_v0) {
					return (_1304_v179).Multiplicity(_1049_v0)
				}
				return (_1301_v178).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1301_v178), 0))).Int()).(_dafny.Int)
			})(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1050_v1, _1049_v0))).Cardinality()))).Uint32(), _1305_v180), _1306_v181))
			_1304_v179 = Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ldambf")).Cardinality()), (_1301_v178).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1301_v178), 0))).Int()).(_dafny.Int)), globalState)
		}
		var _1308_i11 _dafny.Int
		_ = _1308_i11
		_1308_i11 = _dafny.Zero
		{
			for !((Companion_Default___.Fm27(_1049_v0, globalState)).Contains(_dafny.IntOfInt64(387))) {
				{
					if (_1308_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1308_i11 = (_1308_i11).Plus(_dafny.One)
					(globalState).F3 = p2
					var _1309_v183 _dafny.MultiSet
					_ = _1309_v183
					_1309_v183 = _dafny.MultiSetOf(p1)
					(globalState).F8 = ((_1309_v183).Cardinality()).Plus(p0)
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1259_v157), 0))
					_ = _index248
					(_1259_v157).ArraySet1(true, (_index248).Int())
					var _1310_v184 D9
					_ = _1310_v184
					_1310_v184 = Companion_D9_.Create_DC21_(_1049_v0, _1149_v86, _1049_v0)
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1259_v157), 0))
					_ = _index249
					(_1259_v157).ArraySet1((_1310_v184).Dtor_cf28(), (_index249).Int())
					(globalState).F8 = p2
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1311_v185 D3
		_ = _1311_v185
		_1311_v185 = Companion_D3_.Create_DC9_()
		r0 = _1311_v185
		return r0
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f12 _dafny.Array
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C4) Ctor__(f12 _dafny.Array) {
	{
		(_this)._f12 = f12
	}
}
func (_this *C4) Fm2(globalState *GlobalState) _dafny.Map {
	{
		if !(_dafny.SetOf(false)).Equals(_dafny.SetOf(false, true, true)) {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_1312_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(206)
			}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l'))).Cardinality())).Merge(func() _dafny.Map {
				var _coll81 = _dafny.NewMapBuilder()
				_ = _coll81
				for _iter82 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(356), _dafny.IntOfInt64(2))); ; {
					_compr_81, _ok82 := _iter82()
					if !_ok82 {
						break
					}
					var _1313_v0 _dafny.Int
					_1313_v0 = interface{}(_compr_81).(_dafny.Int)
					if ((_dafny.IntOfInt64(356)).Cmp(_1313_v0) <= 0) && ((_1313_v0).Cmp(_dafny.IntOfInt64(2)) < 0) {
						_coll81.Add((_1313_v0).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}(func(_1314_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(742))).Cardinality())).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(759))).Cardinality())).Cardinality()))
					}
				}
				return _coll81.ToMap()
			}())
		} else {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-686), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(959), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), !(true))).Cardinality(), false)).Cardinality())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-167))).Cardinality()), _dafny.IntOfInt64(717))).Cardinality()))
		}
	}
}
func (_this *C4) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return ((func() bool {
			if true {
				return false
			}
			return true
		})()) || (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(264), true)).Cardinality()).Cmp(_dafny.IntOfInt64(-407)) < 0)
	}
}
func (_this *C4) Fm12(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(825)), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-928)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gpglxhpdq")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bw")).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true)).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("trm")).Cardinality()))
	}
}
func (_this *C4) Fm13(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool {
	{
		var _source24 D0 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_dafny.MultiSetOf(!(false), true)).Cardinality()))
		_ = _source24
		if _source24.Is_DC1() {
			var _1315___mcc_h0 _dafny.Int = _source24.Get_().(D0_DC1).Cf1
			_ = _1315___mcc_h0
			var _1316___mcc_h1 _dafny.Int = _source24.Get_().(D0_DC1).Cf2
			_ = _1316___mcc_h1
			var _1317_cf2 _dafny.Int = _1316___mcc_h1
			_ = _1317_cf2
			var _1318_cf1 _dafny.Int = _1315___mcc_h0
			_ = _1318_cf1
			return (_1318_cf1).Cmp((_dafny.SetOf(!(true), true, false, !(false))).Cardinality()) >= 0
		} else if _source24.Is_DC2() {
			return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kammnnur")).Cardinality())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true, !(!(true)), !(false)), false)).Cardinality()) < 0
		} else if _source24.Is_DC0() {
			var _1319___mcc_h2 _dafny.Int = _source24.Get_().(D0_DC0).Cf0
			_ = _1319___mcc_h2
			var _1320_cf0 _dafny.Int = _1319___mcc_h2
			_ = _1320_cf0
			return !((_dafny.MultiSetOf(_dafny.SeqOf(true))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.SeqOf((Companion_D2_.Create_DC6_(_1320_cf0, true, true)).Dtor_cf8(), false), _dafny.SeqOf(false, true, true, true))))
		} else {
			var _1321___mcc_h3 D0 = _source24.Get_().(D0_DC3).Cf3
			_ = _1321___mcc_h3
			var _1322_cf3 D0 = _1321___mcc_h3
			_ = _1322_cf3
			return false
		}
	}
}
func (_this *C4) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1323_v0 _dafny.Array
		_ = _1323_v0
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_31
		var _nw217 _dafny.Array
		_ = _nw217
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw217 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) bool = func(_1324_i0 _dafny.Int) bool {
				return (_dafny.MultiSetOf(true)).Equals(_dafny.MultiSetOf(true, true, false))
			}
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw217 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw217).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw217).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1323_v0 = _nw217
		var _1325_v1 bool
		_ = _1325_v1
		_1325_v1 = false
		var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))
		_ = _index250
		(_1323_v0).ArraySet1((_1325_v1) == (!(!(!(_1325_v1)))), (_index250).Int())
		var _1326_v2 _dafny.Sequence
		_ = _1326_v2
		_1326_v2 = _dafny.UnicodeSeqOfUtf8Bytes("upyxghc")
		var _1327_v3 _dafny.Int
		_ = _1327_v3
		_1327_v3 = _dafny.IntOfInt64(-263)
		var _1328_v4 _dafny.Sequence
		_ = _1328_v4
		_1328_v4 = _dafny.SeqOf(Companion_Default___.Fm14(_1327_v3, globalState), _1326_v2, _dafny.UnicodeSeqOfUtf8Bytes("vjftgkl"))
		var _1329_v5 _dafny.MultiSet
		_ = _1329_v5
		_1329_v5 = _dafny.MultiSetOf(_1326_v2)
		var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))
		_ = _index251
		(_1323_v0).ArraySet1(((_dafny.MultiSetOf(_1326_v2, _1326_v2)).Union(_dafny.MultiSetFromSeq(_1328_v4))).IsProperSubsetOf(_1329_v5), (_index251).Int())
		_1326_v2 = _1326_v2
		var _hi6 _dafny.Int = _1327_v3
		_ = _hi6
		for _1330_i1 := (_1327_v3).Minus(_dafny.IntOfInt64(-637)); _1330_i1.Cmp(_hi6) < 0; _1330_i1 = _1330_i1.Plus(_dafny.One) {
			(globalState).F4 = _1325_v1
			var _1331_v6 *C1
			_ = _1331_v6
			var _nw218 *C1 = New_C1_()
			_ = _nw218
			_nw218.Ctor__()
			_1331_v6 = _nw218
			(globalState).F4 = ((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)) || (_1325_v1)
			_1323_v0 = _1323_v0
		}
		var _1332_v7 *C3
		_ = _1332_v7
		var _nw219 *C3 = New_C3_()
		_ = _nw219
		_nw219.Ctor__()
		_1332_v7 = _nw219
		if (_1327_v3).Cmp((_dafny.MultiSetOf(_1327_v3, _dafny.IntOfInt64(663), (_dafny.Zero).Minus(_1327_v3))).Cardinality()) < 0 {
			var _1333_v8 *C1
			_ = _1333_v8
			var _nw220 *C1 = New_C1_()
			_ = _nw220
			_nw220.Ctor__()
			_1333_v8 = _nw220
			(globalState).F4 = _1325_v1
			var _1334_v9 _dafny.Map
			_ = _1334_v9
			_1334_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_1327_v3, Companion_Default___.Fm20(_1325_v1, globalState)), _1325_v1)
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))
			_ = _index252
			(_1323_v0).ArraySet1((func() bool {
				if (_1334_v9).Contains(_1327_v3) {
					return (_1334_v9).Get(_1327_v3).(bool)
				}
				return true
			})(), (_index252).Int())
			var _1335_v10 _dafny.MultiSet
			_ = _1335_v10
			_1335_v10 = _dafny.MultiSetOf((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool))
			var _1336_v11 _dafny.Map
			_ = _1336_v11
			_1336_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_v0, Companion_Default___.SafeDivisionInt((_1335_v10).Cardinality(), _1327_v3))
			var _1337_v12 D2
			_ = _1337_v12
			_1337_v12 = Companion_D2_.Create_DC5_(_1323_v0)
			var _1338_v13 _dafny.Map
			_ = _1338_v13
			_1338_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)), _1327_v3)
			var _1339_v14 _dafny.Array
			_ = _1339_v14
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_32
			var _nw221 _dafny.Array
			_ = _nw221
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw221 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Int = (func(_1340_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1341_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1341_i2, _1340_v3)
					}
				})(_1327_v3)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw221 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw221).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw221).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1339_v14 = _nw221
			var _1342_v15 _dafny.Map
			_ = _1342_v15
			_1342_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1338_v13).Contains((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)) {
					return (_1338_v13).Get((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)).(_dafny.Int)
				}
				return _1327_v3
			})(), _1339_v14)
			_1336_v11 = (_1336_v11).Update((_1337_v12).Dtor_cf5(), (_1342_v15).Cardinality())
			(globalState).F3 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1326_v2).Cardinality()), _1327_v3)
		} else {
			var _1343_v16 _dafny.Array
			_ = _1343_v16
			var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw222
			_1343_v16 = _nw222
			var _1344_v17 _dafny.Map
			_ = _1344_v17
			_1344_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1327_v3, (_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool))
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))
			_ = _index253
			(_1343_v16).ArraySet1((_1344_v17).Cardinality(), (_index253).Int())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))
			_ = _index254
			(_1343_v16).ArraySet1((_dafny.IntOfInt64(975)).Plus((_1327_v3).Times(_1327_v3)), (_index254).Int())
			(globalState).F4 = ((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)) == (false)
			(globalState).F3 = _1327_v3
			var _1345_v18 _dafny.Set
			_ = _1345_v18
			_1345_v18 = _dafny.SetOf(_dafny.IntOfInt64(207))
			var _1346_v19 *C0
			_ = _1346_v19
			var _nw223 *C0 = New_C0_()
			_ = _nw223
			_nw223.Ctor__()
			_1346_v19 = _nw223
			var _1347_v20 D13
			_ = _1347_v20
			_1347_v20 = Companion_D13_.Create_DC30_((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool), (_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool), _1345_v18, _1346_v19)
			var _1348_v21 D13
			_ = _1348_v21
			_1348_v21 = Companion_D13_.Create_DC31_(_1347_v20)
			var _1349_v23 D0
			_ = _1349_v23
			_1349_v23 = Companion_D0_.Create_DC1_(_1327_v3, _dafny.IntOfUint32((_1326_v2).Cardinality()))
			_1348_v21 = (func() D13 {
				if (func() _dafny.Set {
					var _coll82 = _dafny.NewBuilder()
					_ = _coll82
					for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(756), _dafny.IntOfInt64(690))); ; {
						_compr_82, _ok83 := _iter83()
						if !_ok83 {
							break
						}
						var _1350_v22 _dafny.Int
						_1350_v22 = interface{}(_compr_82).(_dafny.Int)
						if ((_dafny.IntOfInt64(756)).Cmp(_1350_v22) <= 0) && ((_1350_v22).Cmp(_dafny.IntOfInt64(690)) < 0) {
							_coll82.Add((_1350_v22).Times(_1327_v3))
						}
					}
					return _coll82.ToSet()
				}()).IsDisjointFrom(_dafny.SetOf((_1349_v23).Dtor_cf1())) {
					return (func() D13 {
						if _1325_v1 {
							return _1348_v21
						}
						return _1348_v21
					})()
				}
				return _1348_v21
			})()
			if true {
				var _1351_v24 _dafny.Sequence
				_ = _1351_v24
				_1351_v24 = _dafny.SeqOf((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool))
				var _rhs244 bool = (_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)
				_ = _rhs244
				var _rhs245 _dafny.Int = _1327_v3
				_ = _rhs245
				var _rhs246 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1351_v24, _1351_v24), _1351_v24)
				_ = _rhs246
				var _lhs196 *GlobalState = globalState
				_ = _lhs196
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				_lhs196.F4 = _rhs244
				_lhs197.F3 = _rhs245
				_1351_v24 = _rhs246
				_1325_v1 = (_dafny.IntOfUint32((_1351_v24).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1326_v2, Companion_Default___.Fm25(globalState))).Cardinality())) < 0
				_1325_v1 = _1325_v1
				var _1352_v25 _dafny.Sequence
				_ = _1352_v25
				_1352_v25 = _dafny.SeqOf(_1327_v3, (_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int))
				var _1353_v26 _dafny.Map
				_ = _1353_v26
				_1353_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_v0, _1352_v25)
				_1353_v26 = (_1353_v26).Update(_1323_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1354_v16 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1355_i3 _dafny.Int) _dafny.Int {
						return (_1354_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1354_v16), 0))).Int()).(_dafny.Int)
					}
				})(_1343_v16))), _1352_v25))
				var _1356_v27 _dafny.Array
				_ = _1356_v27
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_33
				var _nw224 _dafny.Array
				_ = _nw224
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw224 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.CodePoint = func(_1357_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('g')
					}
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw224 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw224).ArraySet1CodePoint(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw224).ArraySet1CodePoint(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1356_v27 = _nw224
				var _1358_v28 _dafny.CodePoint
				_ = _1358_v28
				_1358_v28 = _dafny.CodePoint('w')
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1356_v27), 0))
				_ = _index255
				(_1356_v27).ArraySet1CodePoint(_1358_v28, (_index255).Int())
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1356_v27), 0))
				_ = _index256
				(_1356_v27).ArraySet1CodePoint(_1358_v28, (_index256).Int())
			} else {
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))
				_ = _index257
				(_1343_v16).ArraySet1((Companion_Default___.SafeDivisionInt((_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int), (_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int))).Minus((_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int)), (_index257).Int())
				var _1359_v29 _dafny.Sequence
				_ = _1359_v29
				_1359_v29 = _dafny.SeqOf(!(_1325_v1))
				var _rhs247 _dafny.Array = _1323_v0
				_ = _rhs247
				var _rhs248 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1359_v29, _dafny.Companion_Sequence_.Update(_1359_v29, (Companion_Default___.SafeIndex(_1327_v3, _dafny.IntOfUint32((_1359_v29).Cardinality()))).Uint32(), _1325_v1)), _dafny.Companion_Sequence_.Concatenate(_1359_v29, _1359_v29))
				_ = _rhs248
				var _rhs249 _dafny.Sequence = _1326_v2
				_ = _rhs249
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				_1323_v0 = _rhs247
				_lhs198.F4 = _rhs248
				_1326_v2 = _rhs249
				var _1360_v30 bool
				_ = _1360_v30
				_1360_v30 = _1325_v1
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))
				_ = _index258
				(_1343_v16).ArraySet1((_this).Fm12((_dafny.IntOfUint32((_1326_v2).Cardinality())).Cmp((_dafny.Zero).Minus((_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int))) <= 0, _1327_v3, _1326_v2, _1360_v30, globalState), (_index258).Int())
				var _1361_v31 _dafny.Sequence
				_ = _1361_v31
				_1361_v31 = _dafny.SeqOf((_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int), _1327_v3)
				var _1362_v32 _dafny.Map
				_ = _1362_v32
				_1362_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1343_v16, _dafny.IntOfUint32((_1361_v31).Cardinality()))
				var _1363_v33 _dafny.Map
				_ = _1363_v33
				_1363_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1327_v3, (_1362_v32).Cardinality())
				(globalState).F8 = Companion_Default___.SafeDivisionInt((_1343_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1343_v16), 0))).Int()).(_dafny.Int), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1327_v3, _1327_v3)).Merge(_1363_v33)).Cardinality())
				var _1364_v34 *C2
				_ = _1364_v34
				var _nw225 *C2 = New_C2_()
				_ = _nw225
				_nw225.Ctor__()
				_1364_v34 = _nw225
			}
		}
		var _1365_v35 bool
		_ = _1365_v35
		_1365_v35 = _1325_v1
		var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))
		_ = _index259
		var _rhs250 bool = _1325_v1
		_ = _rhs250
		var _rhs251 _dafny.Sequence = (func() _dafny.Sequence {
			if (_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool) {
				return _1326_v2
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("vtakwmgb")
		})()
		_ = _rhs251
		var _rhs252 bool = (_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)
		_ = _rhs252
		var _rhs253 bool = !(_1365_v35)
		_ = _rhs253
		var _rhs254 bool = ((_1323_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))).Int()).(bool)) || (_1325_v1)
		_ = _rhs254
		var _lhs199 *GlobalState = globalState
		_ = _lhs199
		var _lhs200 _dafny.Array = _1323_v0
		_ = _lhs200
		var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_1323_v0), 0))
		_ = _lhs201
		var _lhs202 *GlobalState = globalState
		_ = _lhs202
		_lhs199.F4 = _rhs250
		_1326_v2 = _rhs251
		(_lhs200).ArraySet1(_rhs252, (_lhs201).Int())
		r1 = _rhs253
		_lhs202.F4 = _rhs254
		r0 = _1325_v1
		r1 = (_1327_v3).Cmp(_dafny.IntOfInt64(-913)) < 0
		return r0, r1
	}
}
func (_this *C4) M9(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _1366_v0 _dafny.Array
		_ = _1366_v0
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_34
		var _nw226 _dafny.Array
		_ = _nw226
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw226 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) bool = func(_1367_i0 _dafny.Int) bool {
				return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("fjxvcpd"))).Contains(false)
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw226 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw226).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw226).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1366_v0 = _nw226
		var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
		_ = _index260
		(_1366_v0).ArraySet1(!(false), (_index260).Int())
		var _1368_v1 bool
		_ = _1368_v1
		_1368_v1 = true
		var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
		_ = _index261
		(_1366_v0).ArraySet1(_1368_v1, (_index261).Int())
		(globalState).F8 = _dafny.IntOfUint32((Companion_Default___.Fm31(_1368_v1, (func() bool {
			if (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool) {
				return (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool)
			}
			return !(_1368_v1)
		})(), _1368_v1, globalState)).Cardinality())
		var _hi7 _dafny.Int = p1
		_ = _hi7
		for _1369_i1 := p0; _1369_i1.Cmp(_hi7) < 0; _1369_i1 = _1369_i1.Plus(_dafny.One) {
			var _1370_v2 _dafny.Sequence
			_ = _1370_v2
			_1370_v2 = _dafny.UnicodeSeqOfUtf8Bytes("koegpi")
			var _1371_v3 _dafny.CodePoint
			_ = _1371_v3
			_1371_v3 = _dafny.CodePoint('e')
			var _1372_v4 _dafny.Sequence
			_ = _1372_v4
			_1372_v4 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("smk"))
			var _1373_v5 *C0
			_ = _1373_v5
			var _nw227 *C0 = New_C0_()
			_ = _nw227
			_nw227.Ctor__()
			_1373_v5 = _nw227
			var _1374_v6 _dafny.Sequence
			_ = _1374_v6
			_1374_v6 = _dafny.SeqOf(_1373_v5)
			var _1375_v7 _dafny.Set
			_ = _1375_v7
			_1375_v7 = _dafny.SetOf(_1370_v2, (_1372_v4).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_1374_v6)).Cardinality(), _dafny.IntOfUint32((_1372_v4).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1376_v8 _dafny.Map
			_ = _1376_v8
			_1376_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1370_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1370_v2).Cardinality()))).Uint32(), _1371_v3), (Companion_Default___.SafeIndex((_dafny.SetOf(_1368_v1)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1370_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1370_v2).Cardinality()))).Uint32(), _1371_v3)).Cardinality()))).Uint32(), _1371_v3), _1370_v2), _1375_v7)
			var _1377_v9 _dafny.MultiSet
			_ = _1377_v9
			_1377_v9 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1370_v2).Cardinality()), _1369_i1, p0)
			var _1378_v10 _dafny.Set
			_ = _1378_v10
			_1378_v10 = _dafny.SetOf(_1377_v9)
			_1376_v8 = (_1376_v8).Update((_1378_v10).IsSubsetOf(_1378_v10), (_1375_v7).Difference(_1375_v7))
			var _1379_v11 _dafny.Map
			_ = _1379_v11
			_1379_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool))
			var _1380_v12 _dafny.Sequence
			_ = _1380_v12
			_1380_v12 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1368_v1), (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool)), _1379_v11)
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _index262
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _index263
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _index264
			var _rhs255 bool = !((_1379_v11).Equals((_1380_v12).Select((Companion_Default___.SafeIndex(_1369_i1, _dafny.IntOfUint32((_1380_v12).Cardinality()))).Uint32()).(_dafny.Map)))
			_ = _rhs255
			var _rhs256 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1370_v2, (Companion_Default___.SafeIndex(_1369_i1, _dafny.IntOfUint32((_1370_v2).Cardinality()))).Uint32(), _1371_v3)
			_ = _rhs256
			var _rhs257 bool = _1368_v1
			_ = _rhs257
			var _rhs258 bool = (_1377_v9).IsDisjointFrom((_1377_v9).Intersection(_1377_v9))
			_ = _rhs258
			var _rhs259 bool = (p1).Cmp(_dafny.IntOfUint32((_1370_v2).Cardinality())) > 0
			_ = _rhs259
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 _dafny.Array = _1366_v0
			_ = _lhs204
			var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _lhs205
			var _lhs206 _dafny.Array = _1366_v0
			_ = _lhs206
			var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _lhs207
			var _lhs208 _dafny.Array = _1366_v0
			_ = _lhs208
			var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))
			_ = _lhs209
			_lhs203.F4 = _rhs255
			_1370_v2 = _rhs256
			(_lhs204).ArraySet1(_rhs257, (_lhs205).Int())
			(_lhs206).ArraySet1(_rhs258, (_lhs207).Int())
			(_lhs208).ArraySet1(_rhs259, (_lhs209).Int())
			(globalState).F4 = (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool)
			_1371_v3 = _1371_v3
		}
		var _1381_v13 D12
		_ = _1381_v13
		_1381_v13 = Companion_D12_.Create_DC28_(Companion_D12_.Create_DC27_(p0, p1, _dafny.One, (_dafny.Zero).Minus(p1)))
		var _1382_v14 D12
		_ = _1382_v14
		_1382_v14 = Companion_D12_.Create_DC28_(_1381_v13)
		var _source25 D12 = Companion_D12_.Create_DC28_(_1381_v13)
		_ = _source25
		if _source25.Is_DC27() {
			var _1383___mcc_h0 _dafny.Int = _source25.Get_().(D12_DC27).Cf36
			_ = _1383___mcc_h0
			var _1384___mcc_h1 _dafny.Int = _source25.Get_().(D12_DC27).Cf37
			_ = _1384___mcc_h1
			var _1385___mcc_h2 _dafny.Int = _source25.Get_().(D12_DC27).Cf38
			_ = _1385___mcc_h2
			var _1386___mcc_h3 _dafny.Int = _source25.Get_().(D12_DC27).Cf39
			_ = _1386___mcc_h3
			var _1387_cf39 _dafny.Int = _1386___mcc_h3
			_ = _1387_cf39
			var _1388_cf38 _dafny.Int = _1385___mcc_h2
			_ = _1388_cf38
			var _1389_cf37 _dafny.Int = _1384___mcc_h1
			_ = _1389_cf37
			var _1390_cf36 _dafny.Int = _1383___mcc_h0
			_ = _1390_cf36
			_1387_cf39 = (_dafny.Zero).Minus(_1390_cf36)
			(globalState).F8 = Companion_Default___.Fm20((_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), globalState)
			var _1391_v15 _dafny.Sequence
			_ = _1391_v15
			_1391_v15 = _dafny.UnicodeSeqOfUtf8Bytes("io")
			_1388_cf38 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1387_cf39), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1391_v15, _1391_v15)).Cardinality()))
			var _1392_v16 *C3
			_ = _1392_v16
			var _nw228 *C3 = New_C3_()
			_ = _nw228
			_nw228.Ctor__()
			_1392_v16 = _nw228
		} else if _source25.Is_DC26() {
			var _1393___mcc_h4 _dafny.Map = _source25.Get_().(D12_DC26).Cf35
			_ = _1393___mcc_h4
			var _1394_cf35 _dafny.Map = _1393___mcc_h4
			_ = _1394_cf35
			var _1395_v17 _dafny.Array
			_ = _1395_v17
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_35
			var _nw229 _dafny.Array
			_ = _nw229
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw229 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = (func(_1396_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1397_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1397_i2, (_dafny.MultiSetOf((_1396_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1396_v0), 0))).Int()).(bool), (_1396_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1396_v0), 0))).Int()).(bool))).Cardinality())
					}
				})(_1366_v0)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw229 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw229).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw229).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1395_v17 = _nw229
			_1395_v17 = _1395_v17
			var _1398_v18 _dafny.Array
			_ = _1398_v18
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_36
			var _nw230 _dafny.Array
			_ = _nw230
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw230 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) D0 = (func(_1399_p1 _dafny.Int) func(_dafny.Int) D0 {
					return func(_1400_i3 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_1399_p1)
					}
				})(p1)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw230 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw230).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw230).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1398_v18 = _nw230
			var _1401_v19 _dafny.Sequence
			_ = _1401_v19
			_1401_v19 = _dafny.UnicodeSeqOfUtf8Bytes("ag")
			var _1402_v20 _dafny.Sequence
			_ = _1402_v20
			_1402_v20 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_1401_v19).Cardinality()))
			var _1403_v21 _dafny.Sequence
			_ = _1403_v21
			_1403_v21 = _dafny.SeqOf(p0, p0, p0, (_1402_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(555))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}(func(_1404_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality()), _dafny.IntOfUint32((_1402_v20).Cardinality()))).Uint32()).(_dafny.Int))
			var _1405_v22 D0
			_ = _1405_v22
			_1405_v22 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_1403_v21).Cardinality()))
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1398_v18), 0))
			_ = _index265
			(_1398_v18).ArraySet1(_1405_v22, (_index265).Int())
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1398_v18), 0))
			_ = _index266
			(_1398_v18).ArraySet1(_1405_v22, (_index266).Int())
			var _1406_v23 _dafny.MultiSet
			_ = _1406_v23
			_1406_v23 = _dafny.MultiSetOf((_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), _1368_v1)
			var _1407_v24 _dafny.Map
			_ = _1407_v24
			_1407_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1406_v23)
			(globalState).F8 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_1406_v23).Contains((_1406_v23).IsDisjointFrom((func() _dafny.MultiSet {
					if (_1407_v24).Contains(p1) {
						return (_1407_v24).Get(p1).(_dafny.MultiSet)
					}
					return _1406_v23
				})())) {
					return (_1406_v23).Multiplicity((_1406_v23).IsDisjointFrom((func() _dafny.MultiSet {
						if (_1407_v24).Contains(p1) {
							return (_1407_v24).Get(p1).(_dafny.MultiSet)
						}
						return _1406_v23
					})()))
				}
				return p0
			})())
			_1395_v17 = _1395_v17
		} else {
			var _1408___mcc_h5 D12 = _source25.Get_().(D12_DC28).Cf40
			_ = _1408___mcc_h5
			var _1409_cf40 D12 = _1408___mcc_h5
			_ = _1409_cf40
			(globalState).F3 = p0
			(globalState).F8 = _dafny.IntOfInt64(-323)
			(globalState).F8 = p0
			var _1410_v25 _dafny.Sequence
			_ = _1410_v25
			_1410_v25 = _dafny.UnicodeSeqOfUtf8Bytes("sdikjo")
			var _1411_v26 _dafny.CodePoint
			_ = _1411_v26
			_1411_v26 = _dafny.CodePoint('w')
			var _1412_v27 _dafny.Sequence
			_ = _1412_v27
			_1412_v27 = _dafny.SeqOf(_1410_v25)
			var _1413_v28 _dafny.Sequence
			_ = _1413_v28
			_1413_v28 = _dafny.SeqOf(p0)
			var _1414_v29 D10
			_ = _1414_v29
			_1414_v29 = Companion_D10_.Create_DC24_(p1, !(_1368_v1), _1410_v25)
			var _1415_v30 _dafny.Array
			_ = _1415_v30
			var _nwElement0_46 _dafny.Sequence = _1410_v25
			_ = _nwElement0_46
			var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(20))
			_ = _nw231
			(_nw231).ArraySet1(_nwElement0_46, 0)
			(_nw231).ArraySet1(_1410_v25, 1)
			(_nw231).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("i"), 2)
			(_nw231).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ceextbwbk"), 3)
			(_nw231).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1410_v25, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(451))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}(func(_1416_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))), 4)
			(_nw231).ArraySet1(_1410_v25, 5)
			(_nw231).ArraySet1(_1410_v25, 6)
			(_nw231).ArraySet1(_1410_v25, 7)
			(_nw231).ArraySet1(_1410_v25, 8)
			(_nw231).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lncxgmpd"), _1410_v25), 9)
			(_nw231).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1410_v25, _1410_v25), 10)
			(_nw231).ArraySet1(_1410_v25, 11)
			(_nw231).ArraySet1(_1410_v25, 12)
			(_nw231).ArraySet1(Companion_Default___.Fm17(_1411_v26, globalState), 13)
			(_nw231).ArraySet1((_1412_v27).Select((Companion_Default___.SafeIndex((_1413_v28).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1413_v28).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1412_v27).Cardinality()))).Uint32()).(_dafny.Sequence), 14)
			(_nw231).ArraySet1(_1410_v25, 15)
			(_nw231).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-408))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_1417_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1418_i6 _dafny.Int) _dafny.CodePoint {
					return _1417_v26
				}
			})(_1411_v26))), 16)
			(_nw231).ArraySet1((_1414_v29).Dtor_cf33(), 17)
			(_nw231).ArraySet1(_1410_v25, 18)
			(_nw231).ArraySet1(_1410_v25, 19)
			_1415_v30 = _nw231
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1415_v30), 0))
			_ = _index267
			(_1415_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1410_v25, _1410_v25), (_index267).Int())
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1415_v30), 0))
			_ = _index268
			(_1415_v30).ArraySet1(_1410_v25, (_index268).Int())
		}
		var _1419_v31 _dafny.Array
		_ = _1419_v31
		var _len0_37 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_37
		var _nw232 _dafny.Array
		_ = _nw232
		if _len0_37.Cmp(_dafny.Zero) == 0 {
			_nw232 = _dafny.NewArray(_len0_37)
		} else {
			var _init37 func(_dafny.Int) _dafny.Map = (func(_1420_p1 _dafny.Int, _1421_v1 bool) func(_dafny.Int) _dafny.Map {
				return func(_1422_i7 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1420_p1, _1421_v1)
				}
			})(p1, _1368_v1)
			_ = _init37
			var _element0_37 = _init37(_dafny.Zero)
			_ = _element0_37
			_nw232 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
			(_nw232).ArraySet1(_element0_37, 0)
			var _nativeLen0_37 = (_len0_37).Int()
			_ = _nativeLen0_37
			for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
				(_nw232).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
			}
		}
		_1419_v31 = _nw232
		var _1423_v32 _dafny.MultiSet
		_ = _1423_v32
		_1423_v32 = _dafny.MultiSetOf(p1, p0)
		var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1419_v31), 0))
		_ = _index269
		(_1419_v31).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm3(_dafny.SeqOf(p0), _1423_v32, false, globalState))).Merge(func() _dafny.Map {
			var _coll83 = _dafny.NewMapBuilder()
			_ = _coll83
			for _iter84 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(58), _dafny.IntOfInt64(318))); ; {
				_compr_83, _ok84 := _iter84()
				if !_ok84 {
					break
				}
				var _1424_v33 _dafny.Int
				_1424_v33 = interface{}(_compr_83).(_dafny.Int)
				if ((_dafny.IntOfInt64(58)).Cmp(_1424_v33) <= 0) && ((_1424_v33).Cmp(_dafny.IntOfInt64(318)) < 0) {
					_coll83.Add(Companion_Default___.SafeDivisionInt(_1424_v33, p1), _1368_v1)
				}
			}
			return _coll83.ToMap()
		}()), (_index269).Int())
		var _1425_v35 _dafny.Sequence
		_ = _1425_v35
		_1425_v35 = _dafny.SeqOf(_1368_v1)
		var _1426_v36 _dafny.Map
		_ = _1426_v36
		_1426_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-494), (_dafny.MultiSetFromSeq(_1425_v35)).Cardinality())
		var _1427_v37 *C3
		_ = _1427_v37
		var _nw233 *C3 = New_C3_()
		_ = _nw233
		_nw233.Ctor__()
		_1427_v37 = _nw233
		var _1428_v38 D0
		_ = _1428_v38
		_1428_v38 = Companion_D0_.Create_DC1_(p0, _dafny.IntOfUint32((_1425_v35).Cardinality()))
		var _1429_v39 _dafny.Map
		_ = _1429_v39
		_1429_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1427_v37, (_1428_v38).Dtor_cf1())
		var _1430_v40 *C3
		_ = _1430_v40
		_1430_v40 = _1427_v37
		var _1431_v41 _dafny.CodePoint
		_ = _1431_v41
		_1431_v41 = _dafny.CodePoint('k')
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1419_v31), 0))
		_ = _index270
		(_1419_v31).ArraySet1((func() _dafny.Map {
			var _coll84 = _dafny.NewMapBuilder()
			_ = _coll84
			for _iter85 := _dafny.Iterate((_1426_v36).Keys().Elements()); ; {
				_compr_84, _ok85 := _iter85()
				if !_ok85 {
					break
				}
				var _1432_v34 _dafny.Int
				_1432_v34 = interface{}(_compr_84).(_dafny.Int)
				if (_1426_v36).Contains(_1432_v34) {
					_coll84.Add(Companion_Default___.SafeModuloInt(_1432_v34, p0), _1368_v1)
				}
			}
			return _coll84.ToMap()
		}()).Update((func() _dafny.Int {
			if (_1429_v39).Contains((_1430_v40)) {
				return (_1429_v39).Get((_1430_v40)).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm1(p0, _1431_v41, (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), _1426_v36, globalState)).Cardinality())))
		})(), _1368_v1), (_index270).Int())
		var _1433_v42 _dafny.Sequence
		_ = _1433_v42
		_1433_v42 = _dafny.SeqOf(Companion_Default___.Fm18((_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), globalState), _dafny.CodePoint('a'), _1431_v41, _1431_v41)
		var _1434_v43 D10
		_ = _1434_v43
		_1434_v43 = Companion_D10_.Create_DC24_(p0, _1368_v1, _1433_v42)
		var _1435_v44 _dafny.Map
		_ = _1435_v44
		_1435_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), (_1434_v43).Dtor_cf33())
		var _1436_v45 D2
		_ = _1436_v45
		_1436_v45 = Companion_D2_.Create_DC7_(_1431_v41, (_1366_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_1366_v0), 0))).Int()).(bool), p1)
		_1435_v44 = (_1435_v44).Update((_1436_v45).Dtor_cf10(), _1433_v42)
	}
}
func (_this *C4) F12() _dafny.Array {
	{
		return _this._f12
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__() {
	{
	}
}
func (_this *C5) Fm2(globalState *GlobalState) _dafny.Map {
	{
		var _source26 D0 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq((Companion_D4_.Create_DC10_(_dafny.SeqOf(!(false), true))).Dtor_cf13())).Cardinality(), (func() _dafny.Set {
			var _coll85 = _dafny.NewBuilder()
			_ = _coll85
			for _iter86 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), true)).Keys().Elements()); ; {
				_compr_85, _ok86 := _iter86()
				if !_ok86 {
					break
				}
				var _1437_v0 _dafny.Set
				_1437_v0 = interface{}(_compr_85).(_dafny.Set)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), true)).Contains(_1437_v0) {
					_coll85.Add(_1437_v0)
				}
			}
			return _coll85.ToSet()
		}()).Cardinality(), (_dafny.SetOf(true, true)).Cardinality())).Cardinality()), _dafny.IntOfInt64(277))
		_ = _source26
		if _source26.Is_DC1() {
			var _1438___mcc_h0 _dafny.Int = _source26.Get_().(D0_DC1).Cf1
			_ = _1438___mcc_h0
			var _1439___mcc_h1 _dafny.Int = _source26.Get_().(D0_DC1).Cf2
			_ = _1439___mcc_h1
			var _1440_cf2 _dafny.Int = _1439___mcc_h1
			_ = _1440_cf2
			var _1441_cf1 _dafny.Int = _1438___mcc_h0
			_ = _1441_cf1
			if !(!(true)) {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1440_cf2, _1441_cf1)
			} else {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), _1441_cf1)
			}
		} else if _source26.Is_DC2() {
			return (func() _dafny.Map {
				var _coll86 = _dafny.NewMapBuilder()
				_ = _coll86
				for _iter87 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(532), _dafny.IntOfInt64(298))); ; {
					_compr_86, _ok87 := _iter87()
					if !_ok87 {
						break
					}
					var _1442_v1 _dafny.Int
					_1442_v1 = interface{}(_compr_86).(_dafny.Int)
					if ((_dafny.IntOfInt64(532)).Cmp(_1442_v1) <= 0) && ((_1442_v1).Cmp(_dafny.IntOfInt64(298)) < 0) {
						_coll86.Add((_1442_v1).Plus(_dafny.IntOfInt64(-161)), _dafny.IntOfInt64(-906))
					}
				}
				return _coll86.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(477), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}(func(_1443_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-366)
			}))).Cardinality())))
		} else if _source26.Is_DC0() {
			var _1444___mcc_h2 _dafny.Int = _source26.Get_().(D0_DC0).Cf0
			_ = _1444___mcc_h2
			var _1445_cf0 _dafny.Int = _1444___mcc_h2
			_ = _1445_cf0
			if true {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_cf0, _1445_cf0)
			} else {
				return func() _dafny.Map {
					var _coll87 = _dafny.NewMapBuilder()
					_ = _coll87
					for _iter88 := _dafny.Iterate((_dafny.SetOf(_1445_cf0, _dafny.IntOfInt64(771), _1445_cf0)).Elements()); ; {
						_compr_87, _ok88 := _iter88()
						if !_ok88 {
							break
						}
						var _1446_v2 _dafny.Int
						_1446_v2 = interface{}(_compr_87).(_dafny.Int)
						if (_dafny.SetOf(_1445_cf0, _dafny.IntOfInt64(771), _1445_cf0)).Contains(_1446_v2) {
							_coll87.Add((_1446_v2).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('k'))).Cardinality()), _1445_cf0)
						}
					}
					return _coll87.ToMap()
				}()
			}
		} else {
			var _1447___mcc_h3 D0 = _source26.Get_().(D0_DC3).Cf3
			_ = _1447___mcc_h3
			var _1448_cf3 D0 = _1447___mcc_h3
			_ = _1448_cf3
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), _dafny.IntOfInt64(-365))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()), _dafny.IntOfInt64(561)))
		}
	}
}
func (_this *C5) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qvoets")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-690))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}(func(_1449_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-419)
		}))), _dafny.SeqOf(_dafny.IntOfInt64(994), _dafny.IntOfInt64(515), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_1450_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(494)
		}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(897))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg62 _dafny.Int) interface{} {
				return coer62(arg62)
			}
		}(func(_1451_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality()))))
	}
}
func (_this *C5) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1452_v0 *C1
		_ = _1452_v0
		var _nw234 *C1 = New_C1_()
		_ = _nw234
		_nw234.Ctor__()
		_1452_v0 = _nw234
		var _1453_v1 _dafny.Array
		_ = _1453_v1
		var _nw235 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(22))
		_ = _nw235
		_1453_v1 = _nw235
		var _1454_v2 _dafny.Map
		_ = _1454_v2
		_1454_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1453_v1, false)
		var _1455_v3 _dafny.Sequence
		_ = _1455_v3
		_1455_v3 = _dafny.UnicodeSeqOfUtf8Bytes("dyu")
		var _1456_v4 _dafny.Set
		_ = _1456_v4
		_1456_v4 = _dafny.SetOf(_1455_v3)
		var _1457_v5 _dafny.Int
		_ = _1457_v5
		_1457_v5 = _dafny.IntOfInt64(112)
		var _1458_v6 _dafny.CodePoint
		_ = _1458_v6
		_1458_v6 = _dafny.CodePoint('o')
		var _1459_v7 bool
		_ = _1459_v7
		_1459_v7 = true
		_1454_v2 = (_1454_v2).Update(_1453_v1, !(_1456_v4).Equals(func() _dafny.Set {
			var _coll88 = _dafny.NewBuilder()
			_ = _coll88
			for _iter89 := _dafny.Iterate((Companion_Default___.Fm42(_1457_v5, _1457_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1458_v6, _1457_v5)).Cardinality(), _1459_v7, globalState)).Keys().Elements()); ; {
				_compr_88, _ok89 := _iter89()
				if !_ok89 {
					break
				}
				var _1460_v8 _dafny.Sequence
				_1460_v8 = interface{}(_compr_88).(_dafny.Sequence)
				if (Companion_Default___.Fm42(_1457_v5, _1457_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1458_v6, _1457_v5)).Cardinality(), _1459_v7, globalState)).Contains(_1460_v8) {
					_coll88.Add(_1460_v8)
				}
			}
			return _coll88.ToSet()
		}()))
		var _1461_v9 _dafny.MultiSet
		_ = _1461_v9
		_1461_v9 = _dafny.MultiSetOf(_1459_v7)
		(globalState).F3 = ((_dafny.IntOfInt64(-68)).Plus(_1457_v5)).Times(((_1461_v9).Difference(_1461_v9)).Cardinality())
		_1459_v7 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("lxdhgpiq"), _1455_v3)
		var _1462_v10 _dafny.Map
		_ = _1462_v10
		_1462_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1459_v7, _1459_v7)
		var _1463_v11 _dafny.Map
		_ = _1463_v11
		_1463_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(87), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1462_v10).Cardinality(), _1458_v6))
		var _1464_v12 _dafny.Map
		_ = _1464_v12
		_1464_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v5, _dafny.CodePoint('q'))
		var _1465_v13 _dafny.Sequence
		_ = _1465_v13
		_1465_v13 = _dafny.SeqOf(_1464_v12, _1464_v12, _1464_v12, _1464_v12)
		var _1466_v15 _dafny.Map
		_ = _1466_v15
		_1466_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-257), _dafny.IntOfInt64(869))
		var _1467_v16 _dafny.Sequence
		_ = _1467_v16
		_1467_v16 = _dafny.SeqOf(_1459_v7, !(_1459_v7))
		var _1468_v17 _dafny.Map
		_ = _1468_v17
		_1468_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1459_v7, (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm26(Companion_D3_.Create_DC8_(_1466_v15), _dafny.IntOfUint32((_1455_v3).Cardinality()), _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1459_v7, _1459_v7), (Companion_Default___.SafeIndex(_1457_v5, _dafny.IntOfUint32((_dafny.SeqOf(_1459_v7, _1459_v7)).Cardinality()))).Uint32(), false), _1467_v16), globalState))))
		var _1469_v18 _dafny.Sequence
		_ = _1469_v18
		_1469_v18 = _dafny.SeqOf(_dafny.IntOfUint32((_1455_v3).Cardinality()), _1457_v5, (_dafny.Zero).Minus((_1468_v17).Cardinality()), _1457_v5, _dafny.IntOfInt64(340))
		var _1470_v19 D19
		_ = _1470_v19
		_1470_v19 = Companion_D19_.Create_DC38_(_1464_v12)
		var _1471_v21 _dafny.Map
		_ = _1471_v21
		_1471_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1467_v16).Cardinality()), _1459_v7)
		var _1472_v22 _dafny.Array
		_ = _1472_v22
		var _nwElement0_47 _dafny.Map = (func() _dafny.Map {
			if (_1463_v11).Contains(_dafny.IntOfInt64(8)) {
				return (_1463_v11).Get(_dafny.IntOfInt64(8)).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v5, _1458_v6)
		})()
		_ = _nwElement0_47
		var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(20))
		_ = _nw236
		(_nw236).ArraySet1(_nwElement0_47, 0)
		(_nw236).ArraySet1(_1464_v12, 1)
		(_nw236).ArraySet1(_1464_v12, 2)
		(_nw236).ArraySet1((_1464_v12).Merge(_1464_v12), 3)
		(_nw236).ArraySet1(((_1465_v13).Select((Companion_Default___.SafeIndex(_1457_v5, _dafny.IntOfUint32((_1465_v13).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1464_v12), 4)
		(_nw236).ArraySet1(_1464_v12, 5)
		(_nw236).ArraySet1((_1464_v12).Merge(_1464_v12), 6)
		(_nw236).ArraySet1(_1464_v12, 7)
		(_nw236).ArraySet1(_1464_v12, 8)
		(_nw236).ArraySet1(_1464_v12, 9)
		(_nw236).ArraySet1(Companion_Default___.Fm43(!(_1459_v7), _1457_v5, _1458_v6, globalState), 10)
		(_nw236).ArraySet1(_1464_v12, 11)
		(_nw236).ArraySet1(func() _dafny.Map {
			var _coll89 = _dafny.NewMapBuilder()
			_ = _coll89
			for _iter90 := _dafny.Iterate((_1469_v18).Elements()); ; {
				_compr_89, _ok90 := _iter90()
				if !_ok90 {
					break
				}
				var _1473_v14 _dafny.Int
				_1473_v14 = interface{}(_compr_89).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_1469_v18, _1473_v14) {
					_coll89.Add((_1473_v14).Times(_1457_v5), _1458_v6)
				}
			}
			return _coll89.ToMap()
		}(), 12)
		(_nw236).ArraySet1(_1464_v12, 13)
		(_nw236).ArraySet1(_1464_v12, 14)
		(_nw236).ArraySet1(_1464_v12, 15)
		(_nw236).ArraySet1((func() _dafny.Map {
			if _1459_v7 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v5, _1458_v6)
			}
			return (_1470_v19).Dtor_cf57()
		})(), 16)
		(_nw236).ArraySet1((_1464_v12).Merge(_1464_v12), 17)
		(_nw236).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v5, _1458_v6), 18)
		(_nw236).ArraySet1(((func() _dafny.Map {
			var _coll90 = _dafny.NewMapBuilder()
			_ = _coll90
			for _iter91 := _dafny.Iterate((_1471_v21).Keys().Elements()); ; {
				_compr_90, _ok91 := _iter91()
				if !_ok91 {
					break
				}
				var _1474_v20 _dafny.Int
				_1474_v20 = interface{}(_compr_90).(_dafny.Int)
				if (_1471_v21).Contains(_1474_v20) {
					_coll90.Add((_1474_v20).Plus((_1464_v12).Cardinality()), _dafny.CodePoint('a'))
				}
			}
			return _coll90.ToMap()
		}()).Update(_1457_v5, _1458_v6)).Merge(_1464_v12), 19)
		_1472_v22 = _nw236
		_1472_v22 = _1472_v22
		var _1475_v23 D0
		_ = _1475_v23
		_1475_v23 = Companion_D0_.Create_DC1_(_1457_v5, _1457_v5)
		var _1476_v24 _dafny.Set
		_ = _1476_v24
		_1476_v24 = _dafny.SetOf(_1475_v23, _1475_v23)
		var _1477_v25 _dafny.MultiSet
		_ = _1477_v25
		_1477_v25 = _dafny.MultiSetOf(_1457_v5)
		var _1478_v26 D19
		_ = _1478_v26
		_1478_v26 = Companion_D19_.Create_DC39_(_1459_v7, _1457_v5)
		(_1452_v0).M13((_1476_v24).Cardinality(), (func() _dafny.MultiSet {
			if _1459_v7 {
				return (_1477_v25).Update(_dafny.IntOfUint32((_1455_v3).Cardinality()), Companion_Default___.Abs(_1457_v5))
			}
			return Companion_Default___.Fm27((_1478_v26).Dtor_cf58(), globalState)
		})(), _1459_v7, globalState)
		r0 = !(true)
		r1 = !((func() bool {
			if (_1457_v5).Cmp(_1457_v5) <= 0 {
				return _1459_v7
			}
			return (_1452_v0).Fm3(_1469_v18, _1477_v25, _1459_v7, globalState)
		})())
		return r0, r1
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f10 _dafny.Int
	_f11 _dafny.Sequence
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.EmptySeq
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

func (_this *C6) Ctor__(f10 _dafny.Int, f11 _dafny.Sequence) {
	{
		(_this)._f10 = f10
		(_this)._f11 = f11
	}
}
func (_this *C6) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((Companion_D0_.Create_DC1_((_this).F10(), (_this).F10())).Dtor_cf2(), _dafny.IntOfInt64(547)), (_this).F10())
	}
}
func (_this *C6) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(!(true)), _dafny.SeqOf(true, false))
	}
}
func (_this *C6) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1479_v0 _dafny.Array
		_ = _1479_v0
		var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw237
		_1479_v0 = _nw237
		var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))
		_ = _index271
		(_1479_v0).ArraySet1((_this).F10(), (_index271).Int())
		var _1480_v1 bool
		_ = _1480_v1
		_1480_v1 = false
		var _1481_v2 _dafny.Map
		_ = _1481_v2
		_1481_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1480_v1, (_this).F10())
		var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))
		_ = _index272
		(_1479_v0).ArraySet1((func() _dafny.Int {
			if _1480_v1 {
				return _dafny.IntOfUint32(((_this).F11()).Cardinality())
			}
			return (func() _dafny.Int {
				if (_1481_v2).Contains(_1480_v1) {
					return (_1481_v2).Get(_1480_v1).(_dafny.Int)
				}
				return _dafny.IntOfInt64(757)
			})()
		})(), (_index272).Int())
		var _hi8 _dafny.Int = _dafny.IntOfInt64(149)
		_ = _hi8
		for _1482_i0 := (_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int); _1482_i0.Cmp(_hi8) < 0; _1482_i0 = _1482_i0.Plus(_dafny.One) {
			_1479_v0 = (func() _dafny.Array {
				if _1480_v1 {
					return _1479_v0
				}
				return _1479_v0
			})()
			var _1483_v3 T0
			_ = _1483_v3
			var _nw238 *C3 = New_C3_()
			_ = _nw238
			_nw238.Ctor__()
			_1483_v3 = _nw238
			var _1484_v4 _dafny.Set
			_ = _1484_v4
			_1484_v4 = _dafny.SetOf(_1483_v3)
			_1484_v4 = (_1484_v4).Intersection(_1484_v4)
			var _1485_v5 _dafny.CodePoint
			_ = _1485_v5
			_1485_v5 = _dafny.CodePoint('d')
			var _1486_v6 _dafny.Map
			_ = _1486_v6
			_1486_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1482_i0)
			var _1487_v7 _dafny.MultiSet
			_ = _1487_v7
			_1487_v7 = _dafny.MultiSetOf((_this).F10())
			var _1488_v8 _dafny.Sequence
			_ = _1488_v8
			_1488_v8 = _dafny.SeqOf(_1480_v1, _1480_v1, _1480_v1, _1480_v1, _1480_v1)
			var _1489_v9 _dafny.Sequence
			_ = _1489_v9
			_1489_v9 = _dafny.SeqOf(_dafny.IntOfUint32(((_this).F11()).Cardinality()), (_dafny.Zero).Minus(_1482_i0))
			var _1490_v10 _dafny.Map
			_ = _1490_v10
			_1490_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1480_v1), _1480_v1)
			var _1491_v11 _dafny.Array
			_ = _1491_v11
			var _nwElement0_48 bool = _1480_v1
			_ = _nwElement0_48
			var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(23))
			_ = _nw239
			(_nw239).ArraySet1(_nwElement0_48, 0)
			(_nw239).ArraySet1(!(_1480_v1) || (_1480_v1), 1)
			(_nw239).ArraySet1(true, 2)
			(_nw239).ArraySet1(_1480_v1, 3)
			(_nw239).ArraySet1((func() bool {
				if _1480_v1 {
					return _1480_v1
				}
				return (_1483_v3).Fm3(Companion_Default___.Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), (_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int))).Cardinality(), _1485_v5, _1480_v1, _1486_v6, globalState), _1487_v7, _1480_v1, globalState)
			})(), 4)
			(_nw239).ArraySet1(!(_1480_v1), 5)
			(_nw239).ArraySet1(_1480_v1, 6)
			(_nw239).ArraySet1((_1488_v8).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1488_v8).Cardinality()))).Uint32()).(bool), 7)
			(_nw239).ArraySet1(_1480_v1, 8)
			(_nw239).ArraySet1(((_this).F10()).Cmp((_this).F10()) <= 0, 9)
			(_nw239).ArraySet1(_1480_v1, 10)
			(_nw239).ArraySet1(_1480_v1, 11)
			(_nw239).ArraySet1((_1483_v3).Fm3(_1489_v9, _1487_v7, _1480_v1, globalState), 12)
			(_nw239).ArraySet1(_1480_v1, 13)
			(_nw239).ArraySet1((_1480_v1) == (true), 14)
			(_nw239).ArraySet1(_1480_v1, 15)
			(_nw239).ArraySet1(_1480_v1, 16)
			(_nw239).ArraySet1(!(_1490_v10).Equals(_1490_v10), 17)
			(_nw239).ArraySet1((func() bool {
				if _1480_v1 {
					return _1480_v1
				}
				return _1480_v1
			})(), 18)
			(_nw239).ArraySet1(_1480_v1, 19)
			(_nw239).ArraySet1(_1480_v1, 20)
			(_nw239).ArraySet1(_1480_v1, 21)
			(_nw239).ArraySet1(_1480_v1, 22)
			_1491_v11 = _nw239
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))
			_ = _index273
			(_1491_v11).ArraySet1(_1480_v1, (_index273).Int())
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))
			_ = _index274
			(_1491_v11).ArraySet1(!((!(_1487_v7).Contains((_this).F10())) && (_1480_v1)), (_index274).Int())
			if ((_dafny.IntOfInt64(515)).Minus((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfInt64(-38)) == 0 {
				var _1492_v12 _dafny.Map
				_ = _1492_v12
				_1492_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1480_v1, _1485_v5)
				_1492_v12 = (_1492_v12).Update(_1480_v1, _1485_v5)
				(globalState).F4 = (_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool)
				var _1493_v13 _dafny.MultiSet
				_ = _1493_v13
				_1493_v13 = _dafny.MultiSetOf(_1480_v1, false)
				var _1494_v14 _dafny.Set
				_ = _1494_v14
				_1494_v14 = _dafny.SetOf((_1493_v13).Cardinality(), Companion_Default___.SafeDivisionInt((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int), (_this).F10()))
				var _1495_v15 _dafny.Sequence
				_ = _1495_v15
				_1495_v15 = _dafny.SeqOf(_1479_v0)
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))
				_ = _index275
				var _rhs260 _dafny.Set = _1494_v14
				_ = _rhs260
				var _rhs261 bool = !(_dafny.Companion_Sequence_.Equal(_1495_v15, _1495_v15))
				_ = _rhs261
				var _lhs210 _dafny.Array = _1491_v11
				_ = _lhs210
				var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))
				_ = _lhs211
				_1494_v14 = _rhs260
				(_lhs210).ArraySet1(_rhs261, (_lhs211).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))
				_ = _index276
				(_1491_v11).ArraySet1(((_this).F10()).Cmp(_1482_i0) < 0, (_index276).Int())
				var _1496_v16 _dafny.MultiSet
				_ = _1496_v16
				_1496_v16 = _dafny.MultiSetOf(_1488_v8)
				r1 = ((_1496_v16).Update(_1488_v8, Companion_Default___.Abs((_dafny.Zero).Minus((_this).F10())))).IsProperSubsetOf((_1496_v16).Intersection(_dafny.MultiSetOf(_1488_v8, _1488_v8)))
			} else {
				var _1497_v17 _dafny.Map
				_ = _1497_v17
				_1497_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1485_v5, _1490_v10)
				var _1498_v18 _dafny.Map
				_ = _1498_v18
				_1498_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					if (_1497_v17).Contains(_1485_v5) {
						return (_1497_v17).Get(_1485_v5).(_dafny.Map)
					}
					return _1490_v10
				})(), _1479_v0)
				(globalState).F4 = (_1498_v18).Contains(_1490_v10)
				r0 = !(_1480_v1)
				var _1499_v19 _dafny.Map
				_ = _1499_v19
				_1499_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), _dafny.SeqOf(_dafny.IntOfUint32((_1489_v9).Cardinality()), (_dafny.Zero).Minus((_this).F10())))
				var _rhs262 bool = (_1480_v1) && (_1480_v1)
				_ = _rhs262
				var _rhs263 _dafny.Map = (_1499_v19).Merge(_1499_v19)
				_ = _rhs263
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				_lhs212.F4 = _rhs262
				_1499_v19 = _rhs263
				var _1500_v20 _dafny.Set
				_ = _1500_v20
				_1500_v20 = _dafny.SetOf((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F11()).Cardinality())), (_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int), (_1482_i0).Minus((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int)))
				var _1501_v21 _dafny.MultiSet
				_ = _1501_v21
				_1501_v21 = _dafny.MultiSetOf(_1480_v1)
				var _1502_v22 _dafny.Array
				_ = _1502_v22
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_38
				var _nw240 _dafny.Array
				_ = _nw240
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw240 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Sequence = (func(_1503_v0 _dafny.Array, _1504_v1 bool, _1505_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1506_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate((Companion_D10_.Create_DC24_((_1503_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1503_v0), 0))).Int()).(_dafny.Int), _1504_v1, (_this).F11())).Dtor_cf33(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(834))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg63 _dafny.Int) interface{} {
									return coer63(arg63)
								}
							}((func(_1507_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1508_i2 _dafny.Int) _dafny.CodePoint {
									return _1507_v5
								}
							})(_1505_v5))))
						}
					})(_1479_v0, _1480_v1, _1485_v5)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw240 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw240).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw240).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1502_v22 = _nw240
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_1502_v22), 0))
				_ = _index277
				(_1502_v22).ArraySet1((_this).F11(), (_index277).Int())
				var _1509_v23 D19
				_ = _1509_v23
				_1509_v23 = Companion_D19_.Create_DC39_(true, (_this).F10())
				var _1510_v24 D19
				_ = _1510_v24
				_1510_v24 = Companion_D19_.Create_DC40_(_1509_v23)
				var _1511_v25 _dafny.Set
				_ = _1511_v25
				_1511_v25 = _dafny.SetOf(_1510_v24, _1510_v24, Companion_D19_.Create_DC40_(_1509_v23), _1510_v24)
				var _1512_v26 _dafny.Map
				_ = _1512_v26
				_1512_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1511_v25, (_dafny.MultiSetOf((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), _1480_v1, (_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), false)).Union(_dafny.MultiSetFromSeq(Companion_Default___.Fm31(_1480_v1, (_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), _1480_v1, globalState))))
				var _1513_v27 _dafny.Map
				_ = _1513_v27
				_1513_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), _1488_v8)
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_1502_v22), 0))
				_ = _index278
				var _rhs264 _dafny.Set = _1500_v20
				_ = _rhs264
				var _rhs265 _dafny.MultiSet = (func() _dafny.MultiSet {
					if (_1512_v26).Contains(_1511_v25) {
						return (_1512_v26).Get(_1511_v25).(_dafny.MultiSet)
					}
					return (_1501_v21).Union(_1501_v21)
				})()
				_ = _rhs265
				var _rhs266 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool) {
						return (func() _dafny.Sequence {
							if (_1513_v27).Contains((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool)) {
								return (_1513_v27).Get((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool)).(_dafny.Sequence)
							}
							return _1488_v8
						})()
					}
					return _1488_v8
				})()).Cardinality())
				_ = _rhs266
				var _rhs267 _dafny.Int = Companion_Default___.SafeModuloInt((_1487_v7).Cardinality(), _1482_i0)
				_ = _rhs267
				var _rhs268 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-313))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1514_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1515_i3 _dafny.Int) _dafny.CodePoint {
						return _1514_v5
					}
				})(_1485_v5))), (_this).F11())
				_ = _rhs268
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				var _lhs215 _dafny.Array = _1502_v22
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_1502_v22), 0))
				_ = _lhs216
				_1500_v20 = _rhs264
				_1501_v21 = _rhs265
				_lhs213.F3 = _rhs266
				_lhs214.F8 = _rhs267
				(_lhs215).ArraySet1(_rhs268, (_lhs216).Int())
				var _1516_v28 _dafny.Set
				_ = _1516_v28
				_1516_v28 = _dafny.SetOf((_1491_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1491_v11), 0))).Int()).(bool), true)
				_1516_v28 = ((_1516_v28).Intersection(_1516_v28)).Intersection(_dafny.SetOf(true))
			}
		}
		var _1517_v29 _dafny.Array
		_ = _1517_v29
		var _nwElement0_49 bool = false
		_ = _nwElement0_49
		var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(17))
		_ = _nw241
		(_nw241).ArraySet1(_nwElement0_49, 0)
		(_nw241).ArraySet1(_1480_v1, 1)
		(_nw241).ArraySet1(_1480_v1, 2)
		(_nw241).ArraySet1(_1480_v1, 3)
		(_nw241).ArraySet1(_1480_v1, 4)
		(_nw241).ArraySet1(_1480_v1, 5)
		(_nw241).ArraySet1(true, 6)
		(_nw241).ArraySet1(_1480_v1, 7)
		(_nw241).ArraySet1(_1480_v1, 8)
		(_nw241).ArraySet1(_1480_v1, 9)
		(_nw241).ArraySet1(_1480_v1, 10)
		(_nw241).ArraySet1(_1480_v1, 11)
		(_nw241).ArraySet1(_1480_v1, 12)
		(_nw241).ArraySet1(_1480_v1, 13)
		(_nw241).ArraySet1(_1480_v1, 14)
		(_nw241).ArraySet1(_1480_v1, 15)
		(_nw241).ArraySet1(_1480_v1, 16)
		_1517_v29 = _nw241
		var _1518_v30 _dafny.Set
		_ = _1518_v30
		_1518_v30 = _dafny.SetOf(_1517_v29)
		var _1519_v31 D2
		_ = _1519_v31
		_1519_v31 = Companion_D2_.Create_DC5_(_1517_v29)
		var _1520_v32 _dafny.Set
		_ = _1520_v32
		_1520_v32 = _dafny.SetOf(_1480_v1, _1480_v1)
		var _1521_v34 _dafny.Set
		_ = _1521_v34
		_1521_v34 = _dafny.SetOf((_this).F10(), (_this).F10())
		var _1522_v35 _dafny.Sequence
		_ = _1522_v35
		_1522_v35 = _dafny.SeqOf(_1480_v1)
		var _1523_v36 _dafny.Sequence
		_ = _1523_v36
		_1523_v36 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_1522_v35).Cardinality())))
		var _1524_v37 _dafny.CodePoint
		_ = _1524_v37
		_1524_v37 = _dafny.CodePoint('y')
		var _1525_v38 _dafny.Map
		_ = _1525_v38
		_1525_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wvptjdmb"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1523_v36).Select((Companion_Default___.SafeIndex((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1523_v36).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wvptjdmb")).Cardinality()))).Uint32(), _1524_v37), false)
		var _1526_v39 _dafny.Array
		_ = _1526_v39
		var _nwElement0_50 bool = ((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_this).F10())) <= 0
		_ = _nwElement0_50
		var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(9))
		_ = _nw242
		(_nw242).ArraySet1(_nwElement0_50, 0)
		(_nw242).ArraySet1((_1518_v30).IsDisjointFrom(_dafny.SetOf(_1517_v29, (_1519_v31).Dtor_cf5())), 1)
		(_nw242).ArraySet1((_1520_v32).IsDisjointFrom(_dafny.SetOf(_1480_v1)), 2)
		(_nw242).ArraySet1(_1480_v1, 3)
		(_nw242).ArraySet1(_1480_v1, 4)
		(_nw242).ArraySet1(_1480_v1, 5)
		(_nw242).ArraySet1(_1480_v1, 6)
		(_nw242).ArraySet1(!((_1521_v34).IsProperSubsetOf(func() _dafny.Set {
			var _coll91 = _dafny.NewBuilder()
			_ = _coll91
			for _iter92 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(964), _dafny.IntOfInt64(-933))); ; {
				_compr_91, _ok92 := _iter92()
				if !_ok92 {
					break
				}
				var _1527_v33 _dafny.Int
				_1527_v33 = interface{}(_compr_91).(_dafny.Int)
				if ((_dafny.IntOfInt64(964)).Cmp(_1527_v33) <= 0) && ((_1527_v33).Cmp(_dafny.IntOfInt64(-933)) < 0) {
					_coll91.Add((_1527_v33).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1480_v1)).Cardinality()))))
				}
			}
			return _coll91.ToSet()
		}())), 7)
		(_nw242).ArraySet1((func() bool {
			if _1480_v1 {
				return !(false)
			}
			return (func() bool {
				if (_1525_v38).Contains((_this).F11()) {
					return (_1525_v38).Get((_this).F11()).(bool)
				}
				return _1480_v1
			})()
		})(), 8)
		_1526_v39 = _nw242
		var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))
		_ = _index279
		(_1517_v29).ArraySet1(true, (_index279).Int())
		var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))
		_ = _index280
		(_1517_v29).ArraySet1(false, (_index280).Int())
		for _iter93 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1479_v0), 0))); ; {
			_guard_loop_1, _ok93 := _iter93()
			if !_ok93 {
				break
			}
			var _1528_i4 _dafny.Int
			_1528_i4 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_1528_i4).Sign() != -1) && ((_1528_i4).Cmp(_dafny.ArrayLen((_1479_v0), 0)) < 0)) {
				(_1479_v0).ArraySet1((_1528_i4).Times((_this).F10()), (_1528_i4).Int())
			}
		}
		var _hi9 _dafny.Int = (_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int)
		_ = _hi9
		for _1529_i5 := (_this).F10(); _1529_i5.Cmp(_hi9) < 0; _1529_i5 = _1529_i5.Plus(_dafny.One) {
			r0 = _1480_v1
			var _1530_v40 _dafny.Sequence
			_ = _1530_v40
			_1530_v40 = _dafny.SeqOf(_1479_v0)
			_1479_v0 = (func() _dafny.Array {
				if _1480_v1 {
					return (_1530_v40).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1530_v40).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _1479_v0
			})()
			var _1531_v41 *C0
			_ = _1531_v41
			var _nw243 *C0 = New_C0_()
			_ = _nw243
			_nw243.Ctor__()
			_1531_v41 = _nw243
			var _1532_v42 D7
			_ = _1532_v42
			_1532_v42 = Companion_D7_.Create_DC18_()
			var _source27 D7 = _1532_v42
			_ = _source27
			if _source27.Is_DC18() {
				(globalState).F3 = Companion_Default___.Fm20((_1517_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))).Int()).(bool), globalState)
				var _1533_v43 _dafny.Sequence
				_ = _1533_v43
				_1533_v43 = _dafny.SeqOf((_this).F10(), _1529_i5)
				var _1534_v44 _dafny.Map
				_ = _1534_v44
				_1534_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1480_v1, (_1517_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))).Int()).(bool))
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))
				_ = _index281
				(_1517_v29).ArraySet1((Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1529_i5), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1533_v43, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1522_v35).Cardinality()), _dafny.IntOfUint32((_1533_v43).Cardinality()))).Uint32(), (_1534_v44).Cardinality())).Cardinality()))).Cmp((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int)) <= 0, (_index281).Int())
				var _1535_v45 D9
				_ = _1535_v45
				_1535_v45 = Companion_D9_.Create_DC20_(_1520_v32)
				_1535_v45 = _1535_v45
				var _1536_v46 _dafny.Sequence
				_ = _1536_v46
				_1536_v46 = _dafny.UnicodeSeqOfUtf8Bytes("mn")
				_1536_v46 = _1536_v46
			} else {
				var _1537___mcc_h0 _dafny.Map = _source27.Get_().(D7_DC17).Cf23
				_ = _1537___mcc_h0
				var _1538_cf23 _dafny.Map = _1537___mcc_h0
				_ = _1538_cf23
				var _1539_v47 _dafny.MultiSet
				_ = _1539_v47
				_1539_v47 = _dafny.MultiSetOf((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(534))
				var _1540_v48 _dafny.MultiSet
				_ = _1540_v48
				_1540_v48 = _dafny.MultiSetOf((_this).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(267))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}(func(_1541_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(634)
				})), _1539_v47, (_1517_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))).Int()).(bool), globalState), _1480_v1, ((_this).F10()).Cmp(_dafny.IntOfInt64(790)) <= 0, (_1517_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))).Int()).(bool), _1480_v1)
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))
				_ = _index282
				(_1479_v0).ArraySet1((_1540_v48).Cardinality(), (_index282).Int())
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))
				_ = _index283
				(_1479_v0).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1517_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_1517_v29), 0))).Int()).(bool) {
						return _dafny.IntOfInt64(817)
					}
					return Companion_Default___.SafeDivisionInt((_this).F10(), _1529_i5)
				})()), (_index283).Int())
				(globalState).F8 = _1529_i5
				(globalState).F3 = (_dafny.Zero).Minus((_1479_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1479_v0), 0))).Int()).(_dafny.Int))
			}
		}
		_1524_v37 = _1524_v37
		r0 = !_dafny.Companion_Sequence_.Equal((_this).F11(), (_this).F11())
		r1 = _1480_v1
		return r0, r1
	}
}
func (_this *C6) M7(p0 bool, p1 _dafny.Sequence, p2 _dafny.Set, globalState *GlobalState) {
	{
		var _1542_i0 _dafny.Int
		_ = _1542_i0
		_1542_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_1542_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1542_i0 = (_1542_i0).Plus(_dafny.One)
					var _1543_v0 _dafny.Sequence
					_ = _1543_v0
					_1543_v0 = _dafny.SeqOf(p0, (func() bool {
						if p0 {
							return p0
						}
						return p0
					})(), p0, !(p0))
					var _1544_v1 _dafny.Array
					_ = _1544_v1
					var _len0_39 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_39
					var _nw244 _dafny.Array
					_ = _nw244
					if _len0_39.Cmp(_dafny.Zero) == 0 {
						_nw244 = _dafny.NewArray(_len0_39)
					} else {
						var _init39 func(_dafny.Int) _dafny.Sequence = func(_1545_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update((_this).F11(), (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32(), _dafny.CodePoint('l'))
						}
						_ = _init39
						var _element0_39 = _init39(_dafny.Zero)
						_ = _element0_39
						_nw244 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
						(_nw244).ArraySet1(_element0_39, 0)
						var _nativeLen0_39 = (_len0_39).Int()
						_ = _nativeLen0_39
						for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
							(_nw244).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
						}
					}
					_1544_v1 = _nw244
					var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1544_v1), 0))
					_ = _index284
					(_1544_v1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14((_this).F10(), globalState), (_this).F11()), (_index284).Int())
					var _1546_v2 *C2
					_ = _1546_v2
					var _nw245 *C2 = New_C2_()
					_ = _nw245
					_nw245.Ctor__()
					_1546_v2 = _nw245
					var _1547_v3 D16
					_ = _1547_v3
					_1547_v3 = Companion_D16_.Create_DC34_(_1546_v2)
					var _1548_v4 _dafny.CodePoint
					_ = _1548_v4
					_1548_v4 = _dafny.CodePoint('x')
					var _1549_v5 _dafny.MultiSet
					_ = _1549_v5
					_1549_v5 = _dafny.MultiSetOf((_this).F10())
					var _1550_v6 _dafny.Map
					_ = _1550_v6
					_1550_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm27(p0, globalState)).IsProperSubsetOf(_1549_v5), _1548_v4)
					var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1544_v1), 0))
					_ = _index285
					var _rhs269 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1543_v0, _1543_v0)
					_ = _rhs269
					var _rhs270 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F11(), (_this).F11())
					_ = _rhs270
					var _rhs271 D16 = _1547_v3
					_ = _rhs271
					var _rhs272 _dafny.CodePoint = (func() _dafny.CodePoint {
						if (_1550_v6).Contains((func() bool {
							if p0 {
								return p0
							}
							return p0
						})()) {
							return (_1550_v6).Get((func() bool {
								if p0 {
									return p0
								}
								return p0
							})()).(_dafny.CodePoint)
						}
						return _1548_v4
					})()
					_ = _rhs272
					var _rhs273 bool = p0
					_ = _rhs273
					var _lhs217 _dafny.Array = _1544_v1
					_ = _lhs217
					var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1544_v1), 0))
					_ = _lhs218
					var _lhs219 *GlobalState = globalState
					_ = _lhs219
					_1543_v0 = _rhs269
					(_lhs217).ArraySet1(_rhs270, (_lhs218).Int())
					_1547_v3 = _rhs271
					_1548_v4 = _rhs272
					_lhs219.F4 = _rhs273
					(globalState).F4 = p0
					var _1551_v7 _dafny.Sequence
					_ = _1551_v7
					_1551_v7 = _dafny.SeqOf((_this).F10())
					if (_1546_v2).Fm3(_1551_v7, _1549_v5, p0, globalState) {
						var _1552_v8 _dafny.Int
						_ = _1552_v8
						var _1553_v9 _dafny.Int
						_ = _1553_v9
						var _1554_v10 _dafny.Int
						_ = _1554_v10
						var _1555_v11 bool
						_ = _1555_v11
						var _out22 _dafny.Int
						_ = _out22
						var _out23 _dafny.Int
						_ = _out23
						var _out24 _dafny.Int
						_ = _out24
						var _out25 bool
						_ = _out25
						_out22, _out23, _out24, _out25 = (_1546_v2).M12(globalState)
						_1552_v8 = _out22
						_1553_v9 = _out23
						_1554_v10 = _out24
						_1555_v11 = _out25
						(globalState).F4 = _1555_v11
						var _1556_v12 _dafny.Set
						_ = _1556_v12
						_1556_v12 = _dafny.SetOf((_this).F10())
						_1556_v12 = (_1556_v12).Intersection(_1556_v12)
						(globalState).F8 = ((_dafny.Zero).Minus(_1553_v9)).Times(((_1551_v7).Select((Companion_Default___.SafeIndex(_1554_v10, _dafny.IntOfUint32((_1551_v7).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_1549_v5).Cardinality())))
						var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1544_v1), 0))
						_ = _index286
						(_1544_v1).ArraySet1((_this).F11(), (_index286).Int())
					} else {
						_1548_v4 = _1548_v4
						(globalState).F4 = p0
						var _1557_v13 _dafny.MultiSet
						_ = _1557_v13
						_1557_v13 = _dafny.MultiSetOf(!(p0), p0)
						var _1558_v14 D9
						_ = _1558_v14
						_1558_v14 = Companion_D9_.Create_DC22_(_dafny.Companion_Sequence_.Update(_1543_v0, (Companion_Default___.SafeIndex((_1557_v13).Cardinality(), _dafny.IntOfUint32((_1543_v0).Cardinality()))).Uint32(), p0))
						var _1559_v15 _dafny.Set
						_ = _1559_v15
						_1559_v15 = _dafny.SetOf(p0)
						var _1560_v16 _dafny.Map
						_ = _1560_v16
						_1560_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1558_v14, ((_this).F10()).Times((_1559_v15).Cardinality()))
						var _1561_v18 _dafny.Map
						_ = _1561_v18
						_1561_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(114), (func() _dafny.Map {
							var _coll92 = _dafny.NewMapBuilder()
							_ = _coll92
							for _iter94 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-921), _dafny.IntOfInt64(719))); ; {
								_compr_92, _ok94 := _iter94()
								if !_ok94 {
									break
								}
								var _1562_v17 _dafny.Int
								_1562_v17 = interface{}(_compr_92).(_dafny.Int)
								if ((_dafny.IntOfInt64(-921)).Cmp(_1562_v17) <= 0) && ((_1562_v17).Cmp(_dafny.IntOfInt64(719)) < 0) {
									_coll92.Add(Companion_Default___.SafeModuloInt(_1562_v17, (_1559_v15).Cardinality()), _1548_v4)
								}
							}
							return _coll92.ToMap()
						}()).Cardinality())
						var _1563_v19 _dafny.MultiSet
						_ = _1563_v19
						_1563_v19 = _dafny.MultiSetOf(_1543_v0)
						_1560_v16 = (_1560_v16).Update(_1558_v14, Companion_Default___.Fm26(Companion_D3_.Create_DC8_(_1561_v18), (_this).F10(), _1563_v19, globalState))
						var _1564_v20 _dafny.Set
						_ = _1564_v20
						_1564_v20 = _dafny.SetOf((_this).F10(), (_this).F10(), (_this).F10())
						(globalState).F8 = ((_this).F10()).Minus((_1564_v20).Cardinality())
						(globalState).F4 = ((_this).F10()).Cmp((func() _dafny.Int {
							if (_1561_v18).Contains((_this).F10()) {
								return (_1561_v18).Get((_this).F10()).(_dafny.Int)
							}
							return (_this).F10()
						})()) <= 0
					}
					var _1565_v21 *C3
					_ = _1565_v21
					var _nw246 *C3 = New_C3_()
					_ = _nw246
					_nw246.Ctor__()
					_1565_v21 = _nw246
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1566_v22 _dafny.Array
		_ = _1566_v22
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_40
		var _nw247 _dafny.Array
		_ = _nw247
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw247 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.Sequence = func(_1567_i3 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F11(), (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32(), _dafny.CodePoint('d')), _dafny.UnicodeSeqOfUtf8Bytes("kjr"))
			}
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw247 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw247).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw247).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1566_v22 = _nw247
		for _iter95 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1566_v22), 0))); ; {
			_guard_loop_2, _ok95 := _iter95()
			if !_ok95 {
				break
			}
			var _1568_i2 _dafny.Int
			_1568_i2 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1568_i2).Sign() != -1) && ((_1568_i2).Cmp(_dafny.ArrayLen((_1566_v22), 0)) < 0)) {
				(_1566_v22).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-239))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}(func(_1569_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hxtmbhm"), (_this).F11())), (_1568_i2).Int())
			}
		}
		var _1570_v23 _dafny.Array
		_ = _1570_v23
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_41
		var _nw248 _dafny.Array
		_ = _nw248
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw248 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) bool = (func(_1571_p0 bool) func(_dafny.Int) bool {
				return func(_1572_i5 _dafny.Int) bool {
					return (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1571_p0, _1571_p0)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg67 _dafny.Int) interface{} {
							return coer67(arg67)
						}
					}(func(_1573_i6 _dafny.Int) _dafny.Int {
						return (_this).F10()
					})))).Cardinality())).IsSubsetOf(_dafny.SetOf((_this).F10()))
				}
			})(p0)
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw248 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw248).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw248).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		_1570_v23 = _nw248
		var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
		_ = _index287
		(_1570_v23).ArraySet1(p0, (_index287).Int())
		var _1574_v24 _dafny.Sequence
		_ = _1574_v24
		_1574_v24 = _dafny.SeqOf((_this).F10(), (_this).F10())
		var _1575_v25 _dafny.Sequence
		_ = _1575_v25
		_1575_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_1574_v24).Cardinality()))
		var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
		_ = _index288
		var _rhs274 _dafny.Int = (_this).F10()
		_ = _rhs274
		var _rhs275 bool = (func() bool {
			if p0 {
				return p0
			}
			return (_this).Fm3(_1575_v25, _dafny.MultiSetOf((_this).F10()), p0, globalState)
		})()
		_ = _rhs275
		var _rhs276 bool = (p0) && (p0)
		_ = _rhs276
		var _rhs277 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F10(), (_this).F10())
		_ = _rhs277
		var _lhs220 *GlobalState = globalState
		_ = _lhs220
		var _lhs221 _dafny.Array = _1570_v23
		_ = _lhs221
		var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
		_ = _lhs222
		var _lhs223 *GlobalState = globalState
		_ = _lhs223
		var _lhs224 *GlobalState = globalState
		_ = _lhs224
		_lhs220.F8 = _rhs274
		(_lhs221).ArraySet1(_rhs275, (_lhs222).Int())
		_lhs223.F4 = _rhs276
		_lhs224.F8 = _rhs277
		var _1576_v26 _dafny.Array
		_ = _1576_v26
		var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw249
		_1576_v26 = _nw249
		var _1577_v27 _dafny.Array
		_ = _1577_v27
		var _nwElement0_51 _dafny.Array = _1576_v26
		_ = _nwElement0_51
		var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(18))
		_ = _nw250
		(_nw250).ArraySet1(_nwElement0_51, 0)
		(_nw250).ArraySet1(_1576_v26, 1)
		(_nw250).ArraySet1(_1576_v26, 2)
		(_nw250).ArraySet1(_1576_v26, 3)
		(_nw250).ArraySet1(_1576_v26, 4)
		(_nw250).ArraySet1(_1576_v26, 5)
		(_nw250).ArraySet1(_1576_v26, 6)
		(_nw250).ArraySet1(_1576_v26, 7)
		(_nw250).ArraySet1(_1576_v26, 8)
		(_nw250).ArraySet1(_1576_v26, 9)
		(_nw250).ArraySet1(_1576_v26, 10)
		(_nw250).ArraySet1(_1576_v26, 11)
		(_nw250).ArraySet1(_1576_v26, 12)
		(_nw250).ArraySet1(_1576_v26, 13)
		(_nw250).ArraySet1(_1576_v26, 14)
		(_nw250).ArraySet1(_1576_v26, 15)
		(_nw250).ArraySet1(_1576_v26, 16)
		(_nw250).ArraySet1(_1576_v26, 17)
		_1577_v27 = _nw250
		var _1578_v28 _dafny.Map
		_ = _1578_v28
		_1578_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1577_v27, _dafny.IntOfInt64(179))
		var _hi10 _dafny.Int = (_dafny.Zero).Minus(((_this).F10()).Plus((_this).F10()))
		_ = _hi10
		for _1579_i7 := (func() _dafny.Int {
			if (_1578_v28).Contains(_1577_v27) {
				return (_1578_v28).Get(_1577_v27).(_dafny.Int)
			}
			return Companion_Default___.Fm20((_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool), globalState)
		})(); _1579_i7.Cmp(_hi10) < 0; _1579_i7 = _1579_i7.Plus(_dafny.One) {
			var _1580_v29 _dafny.MultiSet
			_ = _1580_v29
			_1580_v29 = _dafny.MultiSetOf(p0)
			var _1581_v30 _dafny.Map
			_ = _1581_v30
			_1581_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool), (_1580_v29).IsSubsetOf(_1580_v29))
			var _1582_v31 _dafny.MultiSet
			_ = _1582_v31
			_1582_v31 = _dafny.MultiSetOf((_this).F10())
			var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
			_ = _index289
			var _rhs278 _dafny.Int = ((_1579_i7).Plus((_this).F10())).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(114), (_1582_v31).Cardinality()))
			_ = _rhs278
			var _rhs279 _dafny.Int = (_this).F10()
			_ = _rhs279
			var _rhs280 _dafny.Map = ((_1581_v30).Merge(_1581_v30)).Merge(_1581_v30)
			_ = _rhs280
			var _rhs281 bool = (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool)
			_ = _rhs281
			var _rhs282 bool = (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool)
			_ = _rhs282
			var _lhs225 *GlobalState = globalState
			_ = _lhs225
			var _lhs226 *GlobalState = globalState
			_ = _lhs226
			var _lhs227 *GlobalState = globalState
			_ = _lhs227
			var _lhs228 _dafny.Array = _1570_v23
			_ = _lhs228
			var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
			_ = _lhs229
			_lhs225.F3 = _rhs278
			_lhs226.F8 = _rhs279
			_1581_v30 = _rhs280
			_lhs227.F4 = _rhs281
			(_lhs228).ArraySet1(_rhs282, (_lhs229).Int())
			var _1583_v32 *C5
			_ = _1583_v32
			var _nw251 *C5 = New_C5_()
			_ = _nw251
			_nw251.Ctor__()
			_1583_v32 = _nw251
			var _1584_v33 _dafny.Map
			_ = _1584_v33
			_1584_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1581_v30, _1579_i7)
			var _1585_v34 _dafny.Map
			_ = _1585_v34
			_1585_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('v'))).Cardinality()), (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool))
			var _1586_v35 _dafny.Map
			_ = _1586_v35
			_1586_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1582_v31, !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0)).Equals(_1585_v34)))
			var _1587_v36 *C0
			_ = _1587_v36
			var _nw252 *C0 = New_C0_()
			_ = _nw252
			_nw252.Ctor__()
			_1587_v36 = _nw252
			var _1588_v37 _dafny.Set
			_ = _1588_v37
			_1588_v37 = _dafny.SetOf(_1587_v36, _1587_v36)
			var _rhs283 _dafny.Map = (_1584_v33).Merge((_1584_v33).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1581_v30, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1588_v37).Cardinality())).Cardinality())))
			_ = _rhs283
			var _rhs284 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1579_i7)).Intersection(_1582_v31), (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool))
			_ = _rhs284
			var _rhs285 _dafny.Array = _1576_v26
			_ = _rhs285
			var _rhs286 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1574_v24, _1575_v25)
			_ = _rhs286
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			_1584_v33 = _rhs283
			_1586_v35 = _rhs284
			_1576_v26 = _rhs285
			_lhs230.F0 = _rhs286
			var _1589_v38 _dafny.Sequence
			_ = _1589_v38
			_1589_v38 = _dafny.SeqOf(_1576_v26)
			var _1590_v39 _dafny.Map
			_ = _1590_v39
			_1590_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1589_v38, _1589_v38), (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool))
			var _1591_v40 _dafny.CodePoint
			_ = _1591_v40
			_1591_v40 = _dafny.CodePoint('l')
			(globalState).F4 = (func() bool {
				if (_1590_v39).Contains(_1589_v38) {
					return (_1590_v39).Get(_1589_v38).(bool)
				}
				return ((_this).F10()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F11(), (Companion_Default___.SafeIndex(_1579_i7, _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32(), _1591_v40)).Cardinality())) > 0
			})()
		}
		if true {
			var _1592_v41 _dafny.Sequence
			_ = _1592_v41
			_1592_v41 = _dafny.SeqOf((_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool))
			var _1593_v42 D4
			_ = _1593_v42
			_1593_v42 = Companion_D4_.Create_DC10_(_dafny.Companion_Sequence_.Update(_1592_v41, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_this).F10()), (_this).F10())).Cardinality()), _dafny.IntOfUint32((_1592_v41).Cardinality()))).Uint32(), p0))
			var _1594_v43 _dafny.Map
			_ = _1594_v43
			_1594_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).Times((_this).F10()), _1593_v42)
			_1594_v43 = (_1594_v43).Update((_this).F10(), _1593_v42)
			var _1595_v44 _dafny.Sequence
			_ = _1595_v44
			_1595_v44 = _dafny.UnicodeSeqOfUtf8Bytes("mjqmhkj")
			_1595_v44 = _1595_v44
			var _1596_v45 _dafny.MultiSet
			_ = _1596_v45
			_1596_v45 = _dafny.MultiSetOf((_this).F10(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1575_v25).Cardinality())), (_this).F10())
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))
			_ = _index290
			(_1570_v23).ArraySet1(((_this).F10()).Cmp((_1596_v45).Cardinality()) == 0, (_index290).Int())
			(globalState).F4 = p0
			(globalState).F3 = (_this).F10()
		} else {
			var _1597_v46 _dafny.CodePoint
			_ = _1597_v46
			_1597_v46 = _dafny.CodePoint('i')
			_1597_v46 = (func() _dafny.CodePoint {
				if (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool) {
					return _dafny.CodePoint('q')
				}
				return _1597_v46
			})()
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_1566_v22), 0))
			_ = _index291
			(_1566_v22).ArraySet1((_this).F11(), (_index291).Int())
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_1566_v22), 0))
			_ = _index292
			(_1566_v22).ArraySet1((_this).F11(), (_index292).Int())
			var _1598_v47 _dafny.Set
			_ = _1598_v47
			_1598_v47 = _dafny.SetOf((_this).F10())
			var _1599_v48 D12
			_ = _1599_v48
			_1599_v48 = Companion_D12_.Create_DC27_((_this).F10(), (_1598_v47).Cardinality(), (_this).F10(), (_this).F10())
			var _1600_v49 _dafny.Map
			_ = _1600_v49
			_1600_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F10())).Cardinality(), (_1599_v48).Dtor_cf39())
			var _1601_v50 D3
			_ = _1601_v50
			_1601_v50 = Companion_D3_.Create_DC8_(_1600_v49)
			(globalState).F8 = Companion_Default___.Fm26(_1601_v50, (_this).F10(), Companion_Default___.Fm44((_this).F10(), globalState), globalState)
			(globalState).F4 = (func() bool {
				if p0 {
					return (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool)
				}
				return (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool)
			})()
			var _1602_v51 _dafny.Sequence
			_ = _1602_v51
			_1602_v51 = _dafny.SeqOf((_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool), p0)
			var _1603_v52 _dafny.Sequence
			_ = _1603_v52
			_1603_v52 = _dafny.SeqOf(_1602_v51, _1602_v51, _1602_v51)
			var _1604_v53 _dafny.Set
			_ = _1604_v53
			_1604_v53 = _dafny.SetOf((_1603_v52).Select((Companion_Default___.SafeIndex((_dafny.SetOf(false, p0, p0, (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool), (_1570_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1570_v23), 0))).Int()).(bool))).Cardinality(), _dafny.IntOfUint32((_1603_v52).Cardinality()))).Uint32()).(_dafny.Sequence))
			(globalState).F3 = ((_1604_v53).Cardinality()).Times(_dafny.IntOfUint32(((_this).F11()).Cardinality()))
		}
		var _1605_v54 _dafny.CodePoint
		_ = _1605_v54
		_1605_v54 = _dafny.CodePoint('k')
		var _1606_v55 D2
		_ = _1606_v55
		_1606_v55 = Companion_D2_.Create_DC7_(_1605_v54, true, (_this).F10())
		var _1607_v56 _dafny.Set
		_ = _1607_v56
		_1607_v56 = _dafny.SetOf((_this).F10(), (_1606_v55).Dtor_cf11())
		(globalState).F8 = (_1607_v56).Cardinality()
	}
}
func (_this *C6) M8(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var _1608_v0 bool
		_ = _1608_v0
		_1608_v0 = true
		var _1609_v1 T0
		_ = _1609_v1
		var _nw253 *C1 = New_C1_()
		_ = _nw253
		_nw253.Ctor__()
		_1609_v1 = _nw253
		var _1610_v2 _dafny.Sequence
		_ = _1610_v2
		_1610_v2 = _dafny.SeqOf(_1609_v1)
		var _1611_v3 _dafny.Map
		_ = _1611_v3
		_1611_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1608_v0, _1610_v2)
		(globalState).F8 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F11()).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1611_v3).Contains(_1608_v0) {
				return (_1611_v3).Get(_1608_v0).(_dafny.Sequence)
			}
			return (_1610_v2)
		})()).Cardinality()))
		var _1612_v4 _dafny.Sequence
		_ = _1612_v4
		_1612_v4 = _dafny.SeqOf(p0)
		var _1613_v5 _dafny.Array
		_ = _1613_v5
		var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw254
		_1613_v5 = _nw254
		var _1614_v6 _dafny.Sequence
		_ = _1614_v6
		_1614_v6 = _dafny.SeqOf(_1613_v5, _1613_v5)
		var _1615_v7 D3
		_ = _1615_v7
		_1615_v7 = Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1614_v6).Cardinality()), _dafny.IntOfUint32(((_this).F11()).Cardinality())))
		var _1616_v8 _dafny.MultiSet
		_ = _1616_v8
		_1616_v8 = _dafny.MultiSetOf(Companion_Default___.Fm31(_1608_v0, _1608_v0, _1608_v0, globalState))
		(globalState).F3 = (_dafny.Zero).Minus((_1612_v4).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm26(_1615_v7, (_this).F10(), _1616_v8, globalState), _dafny.IntOfUint32((_1612_v4).Cardinality()))).Uint32()).(_dafny.Int))
		var _1617_v9 *C3
		_ = _1617_v9
		var _nw255 *C3 = New_C3_()
		_ = _nw255
		_nw255.Ctor__()
		_1617_v9 = _nw255
		var _1618_i0 _dafny.Int
		_ = _1618_i0
		_1618_i0 = _dafny.Zero
		{
			for ((_this).F10()).Cmp((_dafny.Zero).Minus((_dafny.IntOfInt64(389)).Times((_this).F10()))) == 0 {
				{
					if (_1618_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1618_i0 = (_1618_i0).Plus(_dafny.One)
					var _1619_v10 _dafny.Array
					_ = _1619_v10
					var _len0_42 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_42
					var _nw256 _dafny.Array
					_ = _nw256
					if _len0_42.Cmp(_dafny.Zero) == 0 {
						_nw256 = _dafny.NewArray(_len0_42)
					} else {
						var _init42 func(_dafny.Int) _dafny.Map = (func(_1620_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_1621_i1 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _1620_p0)
							}
						})(p0)
						_ = _init42
						var _element0_42 = _init42(_dafny.Zero)
						_ = _element0_42
						_nw256 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
						(_nw256).ArraySet1(_element0_42, 0)
						var _nativeLen0_42 = (_len0_42).Int()
						_ = _nativeLen0_42
						for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
							(_nw256).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
						}
					}
					_1619_v10 = _nw256
					var _1622_v11 *C4
					_ = _1622_v11
					var _nw257 *C4 = New_C4_()
					_ = _nw257
					_nw257.Ctor__(_1619_v10)
					_1622_v11 = _nw257
					var _1623_v12 _dafny.Map
					_ = _1623_v12
					_1623_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1608_v0, (_this).F10())
					(globalState).F3 = Companion_Default___.SafeDivisionInt(p0, (func() _dafny.Int {
						if (_1623_v12).Contains(_1608_v0) {
							return (_1623_v12).Get(_1608_v0).(_dafny.Int)
						}
						return (_this).F10()
					})())
					var _1624_v13 _dafny.CodePoint
					_ = _1624_v13
					_1624_v13 = _dafny.CodePoint('f')
					var _1625_v14 _dafny.Map
					_ = _1625_v14
					_1625_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1622_v11).Fm13(_1608_v0, (_this).F11(), _1608_v0, globalState), _1624_v13)
					_1625_v14 = (_1625_v14).Update(_1608_v0, _1624_v13)
					var _1626_v15 _dafny.MultiSet
					_ = _1626_v15
					_1626_v15 = _dafny.MultiSetOf((_this).F10())
					var _1627_v16 _dafny.Sequence
					_ = _1627_v16
					_1627_v16 = _dafny.SeqOf(!((_1622_v11).Fm3(_1612_v4, _1626_v15, _1608_v0, globalState)))
					var _1628_v17 bool
					_ = _1628_v17
					_1628_v17 = _1608_v0
					var _1629_v18 D0
					_ = _1629_v18
					_1629_v18 = Companion_D0_.Create_DC1_((_1617_v9).Fm15(_dafny.IntOfUint32((_1627_v16).Cardinality()), globalState), (_dafny.Zero).Minus(((_1622_v11).Fm12(_1608_v0, p0, (_this).F11(), _1628_v17, globalState)).Minus(p0)))
					var _1630_v19 _dafny.Map
					_ = _1630_v19
					_1630_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0)
					var _1631_v20 _dafny.Set
					_ = _1631_v20
					_1631_v20 = _dafny.SetOf(p0)
					var _1632_v21 _dafny.MultiSet
					_ = _1632_v21
					_1632_v21 = _dafny.MultiSetOf(_dafny.SetOf((func() _dafny.Int {
						if (_1626_v15).Contains(p0) {
							return (_1626_v15).Multiplicity(p0)
						}
						return p0
					})(), (_this).F10(), (_1630_v19).Cardinality()), _1631_v20, _1631_v20)
					var _1633_v22 _dafny.Map
					_ = _1633_v22
					_1633_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_1624_v13, _1608_v0, (_this).F10()), (_1632_v21).Cardinality())
					var _1634_v23 _dafny.Set
					_ = _1634_v23
					_1634_v23 = _dafny.SetOf(((_this).F10()).Minus((_this).F10()), (_1633_v22).Cardinality())
					var _pat_let_tv63 = _1626_v15
					_ = _pat_let_tv63
					var _pat_let_tv64 = _1626_v15
					_ = _pat_let_tv64
					var _pat_let_tv65 = p0
					_ = _pat_let_tv65
					var _rhs287 bool = _1608_v0
					_ = _rhs287
					var _rhs288 bool = !(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("wcumpsjw"), _1624_v13))
					_ = _rhs288
					var _rhs289 D0 = func(_pat_let53_0 D0) D0 {
						return func(_1635_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let54_0 _dafny.Int) D0 {
								return func(_1636_dt__update_hcf1_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC1_(_1636_dt__update_hcf1_h0, (_1635_dt__update__tmp_h0).Dtor_cf2())
								}(_pat_let54_0)
							}((_dafny.Zero).Minus((func() _dafny.Int {
								if (_pat_let_tv63).Contains((_this).F10()) {
									return (_pat_let_tv64).Multiplicity((_this).F10())
								}
								return _pat_let_tv65
							})()))
						}(_pat_let53_0)
					}(_1629_v18)
					_ = _rhs289
					var _rhs290 _dafny.Set = _1634_v23
					_ = _rhs290
					var _lhs231 *GlobalState = globalState
					_ = _lhs231
					_1608_v0 = _rhs287
					_lhs231.F4 = _rhs288
					_1629_v18 = _rhs289
					_1631_v20 = _rhs290
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1637_v24 _dafny.Set
		_ = _1637_v24
		_1637_v24 = _dafny.SetOf(_1608_v0, _1608_v0)
		var _1638_v25 _dafny.Sequence
		_ = _1638_v25
		_1638_v25 = _dafny.SeqOf(!((_1637_v24).IsProperSubsetOf(_1637_v24)))
		var _1639_v26 _dafny.Map
		_ = _1639_v26
		_1639_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1608_v0, _1608_v0)
		var _1640_i2 _dafny.Int
		_ = _1640_i2
		_1640_i2 = _dafny.Zero
		{
			for (_1638_v25).Select((Companion_Default___.SafeIndex((_1639_v26).Cardinality(), _dafny.IntOfUint32((_1638_v25).Cardinality()))).Uint32()).(bool) {
				{
					if (_1640_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1640_i2 = (_1640_i2).Plus(_dafny.One)
					var _1641_v27 _dafny.Array
					_ = _1641_v27
					var _len0_43 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_43
					var _nw258 _dafny.Array
					_ = _nw258
					if _len0_43.Cmp(_dafny.Zero) == 0 {
						_nw258 = _dafny.NewArray(_len0_43)
					} else {
						var _init43 func(_dafny.Int) bool = (func(_1642_v0 bool) func(_dafny.Int) bool {
							return func(_1643_i3 _dafny.Int) bool {
								return _1642_v0
							}
						})(_1608_v0)
						_ = _init43
						var _element0_43 = _init43(_dafny.Zero)
						_ = _element0_43
						_nw258 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
						(_nw258).ArraySet1(_element0_43, 0)
						var _nativeLen0_43 = (_len0_43).Int()
						_ = _nativeLen0_43
						for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
							(_nw258).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
						}
					}
					_1641_v27 = _nw258
					var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1641_v27), 0))
					_ = _index293
					(_1641_v27).ArraySet1(_1608_v0, (_index293).Int())
					var _1644_v28 _dafny.MultiSet
					_ = _1644_v28
					_1644_v28 = _dafny.MultiSetOf((_this).F10())
					var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1641_v27), 0))
					_ = _index294
					(_1641_v27).ArraySet1(((_1644_v28).Intersection(_1644_v28)).IsDisjointFrom(_1644_v28), (_index294).Int())
					(globalState).F8 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(463), _dafny.IntOfUint32(((_this).F11()).Cardinality()))
					var _1645_v29 _dafny.Sequence
					_ = _1645_v29
					_1645_v29 = _dafny.SeqOf((_dafny.Zero).Minus(p0))
					var _source28 _dafny.Sequence = _1645_v29
					_ = _source28
					var _1646___mcc_h0 _dafny.Sequence = _source28
					_ = _1646___mcc_h0
					var _1647_cf24 _dafny.Sequence = _1646___mcc_h0
					_ = _1647_cf24
					var _1648_v30 _dafny.Array
					_ = _1648_v30
					var _nw259 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(9))
					_ = _nw259
					_1648_v30 = _nw259
					var _1649_v31 _dafny.MultiSet
					_ = _1649_v31
					_1649_v31 = _dafny.MultiSetOf(_1648_v30)
					var _1650_v32 _dafny.MultiSet
					_ = _1650_v32
					_1650_v32 = _1649_v31
					_1650_v32 = _1650_v32
					r0 = _1644_v28
					var _1651_v33 _dafny.Sequence
					_ = _1651_v33
					_1651_v33 = _dafny.UnicodeSeqOfUtf8Bytes("uoxvu")
					var _1652_v34 _dafny.CodePoint
					_ = _1652_v34
					_1652_v34 = _dafny.CodePoint('e')
					var _1653_v35 D10
					_ = _1653_v35
					_1653_v35 = Companion_D10_.Create_DC24_(p0, (_1641_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1641_v27), 0))).Int()).(bool), _dafny.SeqOf(_1652_v34))
					_1651_v33 = (_1653_v35).Dtor_cf33()
					(globalState).F8 = (_dafny.Zero).Minus(((_this).F10()).Times(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F11()).Cardinality()), _dafny.IntOfUint32((_1651_v33).Cardinality()))))
					var _1654_v36 _dafny.Array
					_ = _1654_v36
					var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
					_ = _nw260
					_1654_v36 = _nw260
					var _1655_v37 _dafny.Array
					_ = _1655_v37
					var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw261
					_1655_v37 = _nw261
					var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_1654_v36), 0))
					_ = _index295
					(_1654_v36).ArraySet1(_1655_v37, (_index295).Int())
					var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_1654_v36), 0))
					_ = _index296
					(_1654_v36).ArraySet1(_1655_v37, (_index296).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1656_v38 _dafny.Sequence
		_ = _1656_v38
		_1656_v38 = _1610_v2
		var _source29 _dafny.Sequence = _1656_v38
		_ = _source29
		var _1657___mcc_h1 _dafny.Sequence = _source29
		_ = _1657___mcc_h1
		var _1658_cf61 _dafny.Sequence = _1657___mcc_h1
		_ = _1658_cf61
		if _1608_v0 {
			var _1659_v39 _dafny.CodePoint
			_ = _1659_v39
			_1659_v39 = _dafny.CodePoint('n')
			_1659_v39 = _1659_v39
			(globalState).F3 = p0
			_1612_v4 = _1612_v4
			(globalState).F8 = ((_dafny.Zero).Minus((_this).F10())).Plus((_this).F10())
			var _1660_v40 _dafny.Sequence
			_ = _1660_v40
			_1660_v40 = _dafny.UnicodeSeqOfUtf8Bytes("nujca")
			var _rhs291 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-550), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F11(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1660_v40).Cardinality())), _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32(), _1659_v39)).Cardinality()))
			_ = _rhs291
			var _rhs292 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_this).F11(), (Companion_Default___.SafeIndex(Companion_Default___.Fm26(_1615_v7, p0, Companion_Default___.Fm44(p0, globalState), globalState), _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32(), _1659_v39)
			_ = _rhs292
			var _lhs232 *GlobalState = globalState
			_ = _lhs232
			_lhs232.F3 = _rhs291
			_1660_v40 = _rhs292
		} else {
			var _1661_v41 _dafny.CodePoint
			_ = _1661_v41
			_1661_v41 = _dafny.CodePoint('a')
			var _1662_v42 _dafny.Map
			_ = _1662_v42
			_1662_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1661_v41, _1608_v0)
			(globalState).F0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(958))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}((func(_1663_v24 _dafny.Set) func(_dafny.Int) _dafny.Int {
				return func(_1664_i4 _dafny.Int) _dafny.Int {
					return (_1663_v24).Cardinality()
				}
			})(_1637_v24))), _dafny.SeqOf((_this).F10(), (_1662_v42).Cardinality(), _dafny.IntOfUint32((_1638_v25).Cardinality())))
			var _1665_v43 D5
			_ = _1665_v43
			_1665_v43 = Companion_D5_.Create_DC13_((_1639_v26).Merge(_1639_v26))
			_1665_v43 = _1665_v43
			var _1666_v44 _dafny.Map
			_ = _1666_v44
			_1666_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F11()).Cardinality()), _1608_v0)
			var _1667_v45 _dafny.Set
			_ = _1667_v45
			_1667_v45 = _dafny.SetOf(p0)
			var _1668_v46 _dafny.MultiSet
			_ = _1668_v46
			_1668_v46 = _dafny.MultiSetOf(_dafny.CodePoint('j'))
			_1666_v44 = (_1666_v44).Update(((_1667_v45).Intersection(_1667_v45)).Cardinality(), (_1668_v46).Equals(_dafny.MultiSetOf(_1661_v41)))
			var _1669_v47 _dafny.Map
			_ = _1669_v47
			_1669_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F11())
			var _1670_v48 _dafny.Map
			_ = _1670_v48
			_1670_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_1669_v47).Contains((_this).F10()) {
					return (_1669_v47).Get((_this).F10()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("t"), (Companion_Default___.SafeIndex((_dafny.SetOf(!(true), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))).Uint32(), _1661_v41)
			})(), (_dafny.Zero).Minus(((_this).F10()).Plus((_this).F10())))
			_1670_v48 = _1670_v48
			var _1671_v49 _dafny.MultiSet
			_ = _1671_v49
			_1671_v49 = _dafny.MultiSetOf((_this).F10(), (_this).F10(), p0)
			var _1672_v50 _dafny.Array
			_ = _1672_v50
			var _nwElement0_52 bool = _1608_v0
			_ = _nwElement0_52
			var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(2))
			_ = _nw262
			(_nw262).ArraySet1(_nwElement0_52, 0)
			(_nw262).ArraySet1((_1609_v1).Fm3(_dafny.Companion_Sequence_.Update(_1612_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1612_v4).Cardinality()))).Uint32(), (_this).F10()), _1671_v49, _1608_v0, globalState), 1)
			_1672_v50 = _nw262
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1672_v50), 0))
			_ = _index297
			(_1672_v50).ArraySet1((_1617_v9).Fm16(_1608_v0, globalState), (_index297).Int())
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1672_v50), 0))
			_ = _index298
			(_1672_v50).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1638_v25, _dafny.Companion_Sequence_.IsProperPrefixOf((_this).F11(), (_this).F11())), (_index298).Int())
		}
		_1608_v0 = !(_1608_v0)
		var _1673_v51 _dafny.Array
		_ = _1673_v51
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_44
		var _nw263 _dafny.Array
		_ = _nw263
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw263 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) bool = (func(_1674_v0 bool) func(_dafny.Int) bool {
				return func(_1675_i5 _dafny.Int) bool {
					return _1674_v0
				}
			})(_1608_v0)
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw263 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw263).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw263).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1673_v51 = _nw263
		var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1673_v51), 0))
		_ = _index299
		(_1673_v51).ArraySet1(_1608_v0, (_index299).Int())
		var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1673_v51), 0))
		_ = _index300
		(_1673_v51).ArraySet1(_1608_v0, (_index300).Int())
		var _1676_v52 *C1
		_ = _1676_v52
		var _nw264 *C1 = New_C1_()
		_ = _nw264
		_nw264.Ctor__()
		_1676_v52 = _nw264
		var _1677_v53 _dafny.Array
		_ = _1677_v53
		var _nw265 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw265
		_1677_v53 = _nw265
		var _1678_v54 _dafny.Sequence
		_ = _1678_v54
		_1678_v54 = _dafny.SeqOf(_1677_v53, _1677_v53)
		var _1679_v55 _dafny.MultiSet
		_ = _1679_v55
		_1679_v55 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt((_1617_v9).Fm15(p0, globalState), _dafny.IntOfUint32((_1678_v54).Cardinality())))
		r0 = _1679_v55
		return r0
	}
}
func (_this *C6) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *C6) F11() _dafny.Sequence {
	{
		return _this._f11
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) Ctor__() {
	{
	}
}
func (_this *C7) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_1680_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))).Cardinality()), _dafny.IntOfInt64(529))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ctfhl")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality())))).Merge((Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-955), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("awvtjnrl")).Cardinality()))))).Dtor_cf12())
	}
}
func (_this *C7) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("no")).Cardinality()), false)).Cardinality(), (Companion_D0_.Create_DC0_((func() _dafny.Set {
			var _coll93 = _dafny.NewBuilder()
			_ = _coll93
			for _iter96 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(385), _dafny.IntOfInt64(693))); ; {
				_compr_93, _ok96 := _iter96()
				if !_ok96 {
					break
				}
				var _1681_v0 _dafny.Int
				_1681_v0 = interface{}(_compr_93).(_dafny.Int)
				if ((_dafny.IntOfInt64(385)).Cmp(_1681_v0) <= 0) && ((_1681_v0).Cmp(_dafny.IntOfInt64(693)) < 0) {
					_coll93.Add((_1681_v0).Times(_dafny.IntOfInt64(-762)))
				}
			}
			return _coll93.ToSet()
		}()).Cardinality())).Dtor_cf0(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(501))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_1682_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}(func(_1683_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()))
		}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(300), _dafny.IntOfInt64(524))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(795))).Cardinality())))).IsDisjointFrom(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(470), true)).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality()), _dafny.IntOfInt64(-52), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(813), Companion_D3_.Create_DC8_(func() _dafny.Map {
			var _coll94 = _dafny.NewMapBuilder()
			_ = _coll94
			for _iter97 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(646), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fsmppmgyx")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(95))).Keys().Elements()); ; {
				_compr_94, _ok97 := _iter97()
				if !_ok97 {
					break
				}
				var _1684_v1 _dafny.Int
				_1684_v1 = interface{}(_compr_94).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(646), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fsmppmgyx")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(95))).Contains(_1684_v1) {
					_coll94.Add((_1684_v1).Minus(_dafny.IntOfInt64(-781)), _dafny.IntOfInt64(797))
				}
			}
			return _coll94.ToMap()
		}()))).Cardinality())).Cardinality())))
	}
}
func (_this *C7) Fm10(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(15))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_1685_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-113)
		})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality()))), _dafny.SeqOf(_dafny.IntOfInt64(993)))).Cardinality())
	}
}
func (_this *C7) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1686_v0 _dafny.Int
		_ = _1686_v0
		_1686_v0 = _dafny.IntOfInt64(905)
		var _1687_v1 D2
		_ = _1687_v1
		_1687_v1 = Companion_D2_.Create_DC6_(_1686_v0, false, true)
		var _1688_v2 bool
		_ = _1688_v2
		_1688_v2 = false
		var _1689_v3 _dafny.MultiSet
		_ = _1689_v3
		_1689_v3 = _dafny.MultiSetOf(_1687_v1, _1687_v1, _1687_v1, _1687_v1, Companion_D2_.Create_DC6_(_1686_v0, _1688_v2, _1688_v2))
		_1689_v3 = _1689_v3
		_1688_v2 = _1688_v2
		(globalState).F8 = _1686_v0
		var _1690_v4 bool
		_ = _1690_v4
		_1690_v4 = _1688_v2
		if func(_source30 bool) bool {
			var _1691___mcc_h6 bool = _source30
			_ = _1691___mcc_h6
			var _1692_cf4 bool = _1691___mcc_h6
			_ = _1692_cf4
			return _1692_cf4
		}(_1690_v4) {
			if (true) && (_1688_v2) {
				(globalState).F3 = (_dafny.Zero).Minus(_1686_v0)
				var _1693_v5 _dafny.Map
				_ = _1693_v5
				_1693_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1688_v2), _1686_v0)
				var _1694_v6 _dafny.Sequence
				_ = _1694_v6
				_1694_v6 = _dafny.SeqOf(_1686_v0, _1686_v0, _1686_v0)
				var _1695_v7 _dafny.MultiSet
				_ = _1695_v7
				_1695_v7 = _dafny.MultiSetOf(_1686_v0)
				var _1696_v8 _dafny.Map
				_ = _1696_v8
				_1696_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1693_v5, (_this).Fm3(_1694_v6, _1695_v7, _1688_v2, globalState))
				var _1697_v9 _dafny.Map
				_ = _1697_v9
				_1697_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1696_v8, _1686_v0)
				_1686_v0 = (func() _dafny.Int {
					if (func() bool {
						if _1688_v2 {
							return !(!(_1688_v2))
						}
						return !(_1688_v2)
					})() {
						return (Companion_Default___.Fm11(_1688_v2, _1688_v2, globalState)).Cardinality()
					}
					return (_dafny.Zero).Minus((func() _dafny.Int {
						if (_1697_v9).Contains(_1696_v8) {
							return (_1697_v9).Get(_1696_v8).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-971)
					})())
				})()
				var _1698_v10 _dafny.Array
				_ = _1698_v10
				var _nwElement0_53 _dafny.Map = _1693_v5
				_ = _nwElement0_53
				var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(5))
				_ = _nw266
				(_nw266).ArraySet1(_nwElement0_53, 0)
				(_nw266).ArraySet1(_1693_v5, 1)
				(_nw266).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _dafny.IntOfInt64(518)), 2)
				(_nw266).ArraySet1(_1693_v5, 3)
				(_nw266).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1686_v0), 4)
				_1698_v10 = _nw266
				var _1699_v11 *C4
				_ = _1699_v11
				var _nw267 *C4 = New_C4_()
				_ = _nw267
				_nw267.Ctor__(_1698_v10)
				_1699_v11 = _nw267
				var _1700_v12 T0
				_ = _1700_v12
				var _nw268 *C4 = New_C4_()
				_ = _nw268
				_nw268.Ctor__(_1698_v10)
				_1700_v12 = _nw268
				var _1701_v13 _dafny.Sequence
				_ = _1701_v13
				_1701_v13 = _dafny.UnicodeSeqOfUtf8Bytes("nph")
				var _1702_v14 _dafny.Map
				_ = _1702_v14
				_1702_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, (_1699_v11).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1703_v13 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1704_i0 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqOf(_1703_v13, _1703_v13, _dafny.UnicodeSeqOfUtf8Bytes("fmigiwsi"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(211))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg74 _dafny.Int) interface{} {
								return coer74(arg74)
							}
						}(func(_1705_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('t')
						})), _1703_v13)).Cardinality())
					}
				})(_1701_v13))), _1695_v7, _1688_v2, globalState))
				var _rhs293 T0 = _1700_v12
				_ = _rhs293
				var _rhs294 bool = !(_1688_v2)
				_ = _rhs294
				var _rhs295 _dafny.Int = (_dafny.Zero).Minus(((_1686_v0).Plus(_dafny.IntOfUint32((_1701_v13).Cardinality()))).Minus(((_1702_v14).Merge(Companion_Default___.Fm21(_1688_v2, globalState))).Cardinality()))
				_ = _rhs295
				var _rhs296 _dafny.Int = _1686_v0
				_ = _rhs296
				var _lhs233 *GlobalState = globalState
				_ = _lhs233
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				_1700_v12 = _rhs293
				_lhs233.F4 = _rhs294
				_lhs234.F8 = _rhs295
				_lhs235.F8 = _rhs296
				var _1706_v16 _dafny.Array
				_ = _1706_v16
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_45
				var _nw269 _dafny.Array
				_ = _nw269
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw269 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.Set = (func(_1707_v2 bool, _1708_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1709_i2 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(func() _dafny.Set {
								var _coll95 = _dafny.NewBuilder()
								_ = _coll95
								for _iter98 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1708_v0, (_dafny.Zero).Minus(_1708_v0))).Keys().Elements()); ; {
									_compr_95, _ok98 := _iter98()
									if !_ok98 {
										break
									}
									var _1710_v15 _dafny.Int
									_1710_v15 = interface{}(_compr_95).(_dafny.Int)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1708_v0, (_dafny.Zero).Minus(_1708_v0))).Contains(_1710_v15) {
										_coll95.Add(Companion_Default___.SafeModuloInt(_1710_v15, _dafny.IntOfInt64(612)))
									}
								}
								return _coll95.ToSet()
							}(), _dafny.SetOf(_1708_v0), _dafny.SetOf(_1708_v0, _dafny.IntOfUint32((_dafny.SeqOf(_1707_v2, _1707_v2)).Cardinality())))
						}
					})(_1688_v2, _1686_v0)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw269 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw269).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw269).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1706_v16 = _nw269
				var _1711_v17 _dafny.Set
				_ = _1711_v17
				_1711_v17 = _dafny.SetOf((_dafny.SetOf(true)).Cardinality(), (_dafny.Zero).Minus(_1686_v0), _dafny.IntOfUint32((_1701_v13).Cardinality()))
				var _1712_v18 _dafny.Set
				_ = _1712_v18
				_1712_v18 = _dafny.SetOf(_1711_v17)
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_1706_v16), 0))
				_ = _index301
				(_1706_v16).ArraySet1(_1712_v18, (_index301).Int())
				var _1713_v19 _dafny.Array
				_ = _1713_v19
				var _nw270 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw270
				_1713_v19 = _nw270
				var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1713_v19), 0))
				_ = _index302
				(_1713_v19).ArraySet1((_1700_v12).Fm3(_1694_v6, _dafny.MultiSetOf(_1686_v0, _1686_v0), _1688_v2, globalState), (_index302).Int())
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_1706_v16), 0))
				_ = _index303
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1713_v19), 0))
				_ = _index304
				var _rhs297 bool = (func() bool {
					if (_1702_v14).Contains(_1688_v2) {
						return (_1702_v14).Get(_1688_v2).(bool)
					}
					return !(_1688_v2)
				})()
				_ = _rhs297
				var _rhs298 T0 = _1700_v12
				_ = _rhs298
				var _rhs299 _dafny.Set = _1712_v18
				_ = _rhs299
				var _rhs300 bool = _1688_v2
				_ = _rhs300
				var _lhs236 _dafny.Array = _1706_v16
				_ = _lhs236
				var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_1706_v16), 0))
				_ = _lhs237
				var _lhs238 _dafny.Array = _1713_v19
				_ = _lhs238
				var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1713_v19), 0))
				_ = _lhs239
				r0 = _rhs297
				_1700_v12 = _rhs298
				(_lhs236).ArraySet1(_rhs299, (_lhs237).Int())
				(_lhs238).ArraySet1(_rhs300, (_lhs239).Int())
			} else {
				var _1714_v20 _dafny.Set
				_ = _1714_v20
				_1714_v20 = _dafny.SetOf(_1686_v0)
				var _1715_v21 _dafny.MultiSet
				_ = _1715_v21
				_1715_v21 = _dafny.MultiSetOf((_dafny.MultiSetOf(_1688_v2, _1688_v2)).Cardinality())
				var _1716_v22 _dafny.Sequence
				_ = _1716_v22
				_1716_v22 = _dafny.SeqOf((_1714_v20).IsSubsetOf(_1714_v20), !(!((_dafny.MultiSetFromSeq(_dafny.SeqOf(_1686_v0, _1686_v0))).IsSubsetOf(_1715_v21))), !(_1688_v2))
				_1716_v22 = _1716_v22
				var _1717_v23 *C1
				_ = _1717_v23
				var _nw271 *C1 = New_C1_()
				_ = _nw271
				_nw271.Ctor__()
				_1717_v23 = _nw271
				var _1718_v24 _dafny.Map
				_ = _1718_v24
				_1718_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1688_v2)
				var _1719_v25 _dafny.Map
				_ = _1719_v25
				_1719_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1717_v23, (func() bool {
					if (_1718_v24).Contains(_1688_v2) {
						return (_1718_v24).Get(_1688_v2).(bool)
					}
					return _1688_v2
				})())
				var _1720_v26 _dafny.Sequence
				_ = _1720_v26
				_1720_v26 = _dafny.SeqOf(_1717_v23, _1717_v23)
				_1719_v25 = (_1719_v25).Update((_1720_v26).Select((Companion_Default___.SafeIndex((_1714_v20).Cardinality(), _dafny.IntOfUint32((_1720_v26).Cardinality()))).Uint32()).(*C1), !(_1688_v2))
				var _1721_v27 _dafny.Map
				_ = _1721_v27
				_1721_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1686_v0)
				var _1722_v28 D12
				_ = _1722_v28
				_1722_v28 = Companion_D12_.Create_DC26_(_1721_v27)
				var _1723_v29 _dafny.Sequence
				_ = _1723_v29
				_1723_v29 = _dafny.SeqOf(_1718_v24, _1718_v24, _1718_v24)
				var _1724_v30 _dafny.Map
				_ = _1724_v30
				_1724_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1722_v28, Companion_Default___.Fm20(_1688_v2, globalState))
				var _1725_v31 _dafny.Set
				_ = _1725_v31
				_1725_v31 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1722_v28, ((_1723_v29).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1723_v29).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())).Update(_1722_v28, _1686_v0), (func() _dafny.Map {
					if _1688_v2 {
						return _1724_v30
					}
					return _1724_v30
				})())
				var _1726_v32 _dafny.Sequence
				_ = _1726_v32
				_1726_v32 = _dafny.UnicodeSeqOfUtf8Bytes("xhgx")
				var _1727_v33 _dafny.Sequence
				_ = _1727_v33
				_1727_v33 = _dafny.SeqOf(_1686_v0, _dafny.IntOfInt64(841))
				var _1728_v34 _dafny.Sequence
				_ = _1728_v34
				_1728_v34 = _1727_v33
				var _1729_v35 _dafny.Map
				_ = _1729_v35
				_1729_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(599))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}(func(_1730_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(940)
				})), _1686_v0)
				var _1731_v36 _dafny.Map
				_ = _1731_v36
				_1731_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1688_v2)
				_1725_v31 = Companion_Default___.Fm45(_1726_v32, ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1728_v34).Cardinality())))).Minus((_1729_v35).Cardinality()), (_1731_v36).Cardinality(), globalState)
				var _1732_v37 _dafny.Array
				_ = _1732_v37
				var _nwElement0_54 bool = false
				_ = _nwElement0_54
				var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(4))
				_ = _nw272
				(_nw272).ArraySet1(_nwElement0_54, 0)
				(_nw272).ArraySet1(_1688_v2, 1)
				(_nw272).ArraySet1(_1688_v2, 2)
				(_nw272).ArraySet1(_1688_v2, 3)
				_1732_v37 = _nw272
				var _1733_v38 D2
				_ = _1733_v38
				_1733_v38 = Companion_D2_.Create_DC5_(_1732_v37)
				var _1734_v39 _dafny.Map
				_ = _1734_v39
				_1734_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1733_v38).Dtor_cf5(), _1688_v2)
				var _1735_v40 _dafny.Array
				_ = _1735_v40
				var _nwElement0_55 bool = _1688_v2
				_ = _nwElement0_55
				var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(10))
				_ = _nw273
				(_nw273).ArraySet1(_nwElement0_55, 0)
				(_nw273).ArraySet1(!(_1688_v2), 1)
				(_nw273).ArraySet1(_1688_v2, 2)
				(_nw273).ArraySet1(_1688_v2, 3)
				(_nw273).ArraySet1(_1688_v2, 4)
				(_nw273).ArraySet1((func() bool {
					if (_1734_v39).Contains(_1732_v37) {
						return (_1734_v39).Get(_1732_v37).(bool)
					}
					return _1688_v2
				})(), 5)
				(_nw273).ArraySet1(_1688_v2, 6)
				(_nw273).ArraySet1(_1688_v2, 7)
				(_nw273).ArraySet1(_1688_v2, 8)
				(_nw273).ArraySet1(false, 9)
				_1735_v40 = _nw273
				var _1736_v41 _dafny.Map
				_ = _1736_v41
				_1736_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1735_v40, _1686_v0)
				_1736_v41 = (_1736_v41).Update(_1732_v37, _dafny.IntOfInt64(-116))
				var _1737_v42 *C5
				_ = _1737_v42
				var _nw274 *C5 = New_C5_()
				_ = _nw274
				_nw274.Ctor__()
				_1737_v42 = _nw274
			}
			var _1738_v43 _dafny.CodePoint
			_ = _1738_v43
			var _1739_v44 _dafny.Set
			_ = _1739_v44
			var _1740_v45 bool
			_ = _1740_v45
			var _out26 _dafny.CodePoint
			_ = _out26
			var _out27 _dafny.Set
			_ = _out27
			var _out28 bool
			_ = _out28
			_out26, _out27, _out28 = (_this).M5(_1688_v2, _1686_v0, (Companion_Default___.Fm46(globalState)).Cardinality(), _1686_v0, globalState)
			_1738_v43 = _out26
			_1739_v44 = _out27
			_1740_v45 = _out28
			_1740_v45 = _1688_v2
			var _1741_v46 _dafny.Sequence
			_ = _1741_v46
			_1741_v46 = _dafny.SeqOf(_1740_v45, _1688_v2, _1740_v45)
			var _1742_v47 D4
			_ = _1742_v47
			_1742_v47 = Companion_D4_.Create_DC10_(_1741_v46)
			var _1743_v48 _dafny.Sequence
			_ = _1743_v48
			_1743_v48 = _dafny.SeqOf(_1742_v47)
			r1 = (_dafny.IntOfUint32((_1743_v48).Cardinality())).Cmp(_1686_v0) <= 0
			var _1744_v49 D4
			_ = _1744_v49
			_1744_v49 = Companion_D4_.Create_DC10_(_1741_v46)
			var _1745_v50 D4
			_ = _1745_v50
			_1745_v50 = Companion_D4_.Create_DC12_(_1744_v49)
			var _1746_v51 D4
			_ = _1746_v51
			_1746_v51 = Companion_D4_.Create_DC12_(_1744_v49)
			var _source31 D4 = (func() D4 {
				if _1688_v2 {
					return Companion_D4_.Create_DC12_(Companion_D4_.Create_DC12_(_1744_v49))
				}
				return _1746_v51
			})()
			_ = _source31
			if _source31.Is_DC11() {
				var _1747___mcc_h0 bool = _source31.Get_().(D4_DC11).Cf14
				_ = _1747___mcc_h0
				var _1748___mcc_h1 bool = _source31.Get_().(D4_DC11).Cf15
				_ = _1748___mcc_h1
				var _1749___mcc_h2 _dafny.Int = _source31.Get_().(D4_DC11).Cf16
				_ = _1749___mcc_h2
				var _1750___mcc_h3 _dafny.Sequence = _source31.Get_().(D4_DC11).Cf17
				_ = _1750___mcc_h3
				var _1751_cf17 _dafny.Sequence = _1750___mcc_h3
				_ = _1751_cf17
				var _1752_cf16 _dafny.Int = _1749___mcc_h2
				_ = _1752_cf16
				var _1753_cf15 bool = _1748___mcc_h1
				_ = _1753_cf15
				var _1754_cf14 bool = _1747___mcc_h0
				_ = _1754_cf14
				var _1755_v52 *C5
				_ = _1755_v52
				var _nw275 *C5 = New_C5_()
				_ = _nw275
				_nw275.Ctor__()
				_1755_v52 = _nw275
				(globalState).F3 = (_1686_v0).Plus(_1686_v0)
				(globalState).F4 = _1688_v2
				var _1756_v53 _dafny.Map
				_ = _1756_v53
				_1756_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, (_1752_cf16).Plus(_1752_cf16))
				var _1757_v54 _dafny.MultiSet
				_ = _1757_v54
				_1757_v54 = _dafny.MultiSetOf(_1686_v0)
				_1756_v53 = (_1756_v53).Update((_1755_v52).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(259))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1758_v45 bool, _1759_cf17 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1760_i4 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1758_v45, _1759_cf17)).Cardinality()
					}
				})(_1740_v45, _1751_cf17))), _1757_v54, _1688_v2, globalState), _1686_v0)
			} else if _source31.Is_DC10() {
				var _1761___mcc_h4 _dafny.Sequence = _source31.Get_().(D4_DC10).Cf13
				_ = _1761___mcc_h4
				var _1762_cf13 _dafny.Sequence = _1761___mcc_h4
				_ = _1762_cf13
				var _1763_v55 _dafny.Array
				_ = _1763_v55
				var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw276
				_1763_v55 = _nw276
				var _1764_v56 _dafny.Array
				_ = _1764_v56
				var _nw277 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw277
				_1764_v56 = _nw277
				var _1765_v57 _dafny.Sequence
				_ = _1765_v57
				_1765_v57 = _dafny.SeqOf(_1764_v56, _1764_v56, _1764_v56)
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1763_v55), 0))
				_ = _index305
				(_1763_v55).ArraySet1(_1765_v57, (_index305).Int())
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_1763_v55), 0))
				_ = _index306
				(_1763_v55).ArraySet1(_dafny.Companion_Sequence_.Update(_1765_v57, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-642), _dafny.IntOfUint32((_1765_v57).Cardinality()))).Uint32(), _1764_v56), (_index306).Int())
				var _1766_v58 _dafny.Array
				_ = _1766_v58
				var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw278
				_1766_v58 = _nw278
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1766_v58), 0))
				_ = _index307
				(_1766_v58).ArraySet1(_1764_v56, (_index307).Int())
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1766_v58), 0))
				_ = _index308
				(_1766_v58).ArraySet1(_1764_v56, (_index308).Int())
				var _1767_v59 _dafny.Sequence
				_ = _1767_v59
				_1767_v59 = _dafny.UnicodeSeqOfUtf8Bytes("tkabsjr")
				var _1768_v60 _dafny.Map
				_ = _1768_v60
				_1768_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1686_v0).Cmp(_dafny.IntOfUint32((_1767_v59).Cardinality())) < 0, _1740_v45)
				_1768_v60 = (_1768_v60).Update(_1688_v2, _1688_v2)
				var _1769_v61 _dafny.Array
				_ = _1769_v61
				var _nwElement0_56 bool = _1688_v2
				_ = _nwElement0_56
				var _nw279 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(24))
				_ = _nw279
				(_nw279).ArraySet1(_nwElement0_56, 0)
				(_nw279).ArraySet1(_1740_v45, 1)
				(_nw279).ArraySet1(_1688_v2, 2)
				(_nw279).ArraySet1(_1688_v2, 3)
				(_nw279).ArraySet1(_1740_v45, 4)
				(_nw279).ArraySet1(false, 5)
				(_nw279).ArraySet1(true, 6)
				(_nw279).ArraySet1(_1740_v45, 7)
				(_nw279).ArraySet1(_1740_v45, 8)
				(_nw279).ArraySet1(_1740_v45, 9)
				(_nw279).ArraySet1(false, 10)
				(_nw279).ArraySet1(true, 11)
				(_nw279).ArraySet1(_1740_v45, 12)
				(_nw279).ArraySet1(!(_1740_v45), 13)
				(_nw279).ArraySet1(_1740_v45, 14)
				(_nw279).ArraySet1(_1688_v2, 15)
				(_nw279).ArraySet1(_1688_v2, 16)
				(_nw279).ArraySet1(_1688_v2, 17)
				(_nw279).ArraySet1(!(true), 18)
				(_nw279).ArraySet1(!(_1740_v45), 19)
				(_nw279).ArraySet1(_1740_v45, 20)
				(_nw279).ArraySet1(_1740_v45, 21)
				(_nw279).ArraySet1(_1688_v2, 22)
				(_nw279).ArraySet1(false, 23)
				_1769_v61 = _nw279
				(_this).M6((_this).Fm10(globalState), _1769_v61, (_1686_v0).Times(_dafny.IntOfInt64(454)), globalState)
			} else {
				var _1770___mcc_h5 D4 = _source31.Get_().(D4_DC12).Cf18
				_ = _1770___mcc_h5
				var _1771_cf18 D4 = _1770___mcc_h5
				_ = _1771_cf18
				var _1772_v62 _dafny.Sequence
				_ = _1772_v62
				_1772_v62 = _dafny.UnicodeSeqOfUtf8Bytes("fwbkil")
				var _1773_v63 D4
				_ = _1773_v63
				_1773_v63 = Companion_D4_.Create_DC11_(_1740_v45, _1740_v45, _1686_v0, _1772_v62)
				var _1774_v64 _dafny.MultiSet
				_ = _1774_v64
				_1774_v64 = _dafny.MultiSetOf(_1686_v0, _1686_v0)
				var _1775_v65 _dafny.Set
				_ = _1775_v65
				_1775_v65 = _dafny.SetOf(_1688_v2, false)
				var _1776_v66 _dafny.Map
				_ = _1776_v66
				_1776_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1774_v64, (_1775_v65).Difference(_1775_v65))
				var _1777_v67 D9
				_ = _1777_v67
				_1777_v67 = Companion_D9_.Create_DC21_(_1740_v45, _1741_v46, _1740_v45)
				var _rhs301 D4 = Companion_Default___.Fm30((_this).Fm10(globalState), _1777_v67, _1688_v2, _1686_v0, globalState)
				_ = _rhs301
				var _rhs302 _dafny.Map = _1776_v66
				_ = _rhs302
				var _rhs303 _dafny.Sequence = Companion_Default___.Fm25(globalState)
				_ = _rhs303
				var _rhs304 _dafny.Int = ((_1686_v0).Minus((_dafny.Zero).Minus(_1686_v0))).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1773_v63).Dtor_cf17(), (Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32(((_1773_v63).Dtor_cf17()).Cardinality()))).Uint32(), _1738_v43)).Cardinality()))
				_ = _rhs304
				var _rhs305 _dafny.Int = _1686_v0
				_ = _rhs305
				var _lhs240 *GlobalState = globalState
				_ = _lhs240
				var _lhs241 *GlobalState = globalState
				_ = _lhs241
				_1773_v63 = _rhs301
				_1776_v66 = _rhs302
				_1772_v62 = _rhs303
				_lhs240.F8 = _rhs304
				_lhs241.F8 = _rhs305
				var _1778_v68 *C2
				_ = _1778_v68
				var _nw280 *C2 = New_C2_()
				_ = _nw280
				_nw280.Ctor__()
				_1778_v68 = _nw280
				var _rhs306 bool = (_1741_v46).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1741_v46).Cardinality()))).Uint32()).(bool)
				_ = _rhs306
				var _rhs307 bool = _1688_v2
				_ = _rhs307
				var _rhs308 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1772_v62).Cardinality()), _1686_v0)
				_ = _rhs308
				var _rhs309 _dafny.Set = _1775_v65
				_ = _rhs309
				var _rhs310 bool = (func() bool {
					if false {
						return _1740_v45
					}
					return _1740_v45
				})()
				_ = _rhs310
				var _lhs242 *GlobalState = globalState
				_ = _lhs242
				var _lhs243 *GlobalState = globalState
				_ = _lhs243
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				_lhs242.F4 = _rhs306
				_1740_v45 = _rhs307
				_lhs243.F8 = _rhs308
				_1775_v65 = _rhs309
				_lhs244.F4 = _rhs310
				(globalState).F4 = (_dafny.SetOf(_1772_v62, _1772_v62)).Contains(_1772_v62)
			}
		} else {
			var _1779_v69 _dafny.Map
			_ = _1779_v69
			_1779_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1688_v2)
			var _1780_v70 _dafny.Sequence
			_ = _1780_v70
			_1780_v70 = _dafny.SeqOf(_1779_v69, _1779_v69)
			var _1781_v71 _dafny.Array
			_ = _1781_v71
			var _nwElement0_57 _dafny.Sequence = _dafny.SeqOf(_1779_v69, _1779_v69, _1779_v69)
			_ = _nwElement0_57
			var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(3))
			_ = _nw281
			(_nw281).ArraySet1(_nwElement0_57, 0)
			(_nw281).ArraySet1(_1780_v70, 1)
			(_nw281).ArraySet1(_1780_v70, 2)
			_1781_v71 = _nw281
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_1781_v71), 0))
			_ = _index309
			(_1781_v71).ArraySet1(_1780_v70, (_index309).Int())
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_1781_v71), 0))
			_ = _index310
			(_1781_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1780_v70, _1780_v70), _1780_v70), (_index310).Int())
			(globalState).F8 = _1686_v0
			var _1782_v72 *C1
			_ = _1782_v72
			var _nw282 *C1 = New_C1_()
			_ = _nw282
			_nw282.Ctor__()
			_1782_v72 = _nw282
			var _1783_v73 _dafny.Sequence
			_ = _1783_v73
			_1783_v73 = _dafny.SeqOf(_1688_v2, _1688_v2, false)
			var _1784_v74 _dafny.Sequence
			_ = _1784_v74
			_1784_v74 = _dafny.SeqOf(_1686_v0)
			var _1785_v75 _dafny.Sequence
			_ = _1785_v75
			_1785_v75 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			var _1786_v76 _dafny.MultiSet
			_ = _1786_v76
			_1786_v76 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1785_v75).Cardinality()))
			var _1787_v77 _dafny.Set
			_ = _1787_v77
			_1787_v77 = _dafny.SetOf(!((_1783_v73).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1783_v73).Cardinality()))).Uint32()).(bool)))
			var _1788_v78 _dafny.CodePoint
			_ = _1788_v78
			_1788_v78 = _dafny.CodePoint('v')
			var _1789_v79 _dafny.MultiSet
			_ = _1789_v79
			_1789_v79 = _dafny.MultiSetOf(_1788_v78)
			var _1790_v80 _dafny.Array
			_ = _1790_v80
			var _nwElement0_58 bool = (_1783_v73).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1783_v73).Cardinality()))).Uint32()).(bool)
			_ = _nwElement0_58
			var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(23))
			_ = _nw283
			(_nw283).ArraySet1(_nwElement0_58, 0)
			(_nw283).ArraySet1(_1688_v2, 1)
			(_nw283).ArraySet1(!(_1688_v2), 2)
			(_nw283).ArraySet1(_1688_v2, 3)
			(_nw283).ArraySet1(_1688_v2, 4)
			(_nw283).ArraySet1(_1688_v2, 5)
			(_nw283).ArraySet1((_1686_v0).Cmp((_this).Fm10(globalState)) >= 0, 6)
			(_nw283).ArraySet1(_1688_v2, 7)
			(_nw283).ArraySet1((_1782_v72).Fm3(_1784_v74, _1786_v76, _1688_v2, globalState), 8)
			(_nw283).ArraySet1(_1688_v2, 9)
			(_nw283).ArraySet1(_1688_v2, 10)
			(_nw283).ArraySet1((_1787_v77).IsSubsetOf(_1787_v77), 11)
			(_nw283).ArraySet1(true, 12)
			(_nw283).ArraySet1(_1688_v2, 13)
			(_nw283).ArraySet1(!(_1789_v79).Equals(_1789_v79), 14)
			(_nw283).ArraySet1(!(_1690_v4), 15)
			(_nw283).ArraySet1(_1688_v2, 16)
			(_nw283).ArraySet1(_1688_v2, 17)
			(_nw283).ArraySet1(!(_1688_v2), 18)
			(_nw283).ArraySet1(_1688_v2, 19)
			(_nw283).ArraySet1(_1688_v2, 20)
			(_nw283).ArraySet1(_1688_v2, 21)
			(_nw283).ArraySet1(_1688_v2, 22)
			_1790_v80 = _nw283
			var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_1790_v80), 0))
			_ = _index311
			(_1790_v80).ArraySet1((_1688_v2) == ((_1782_v72).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-537))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}((func(_1791_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1792_i5 _dafny.Int) _dafny.Int {
					return _1791_v0
				}
			})(_1686_v0))), _1786_v76, _1688_v2, globalState)), (_index311).Int())
			var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_1790_v80), 0))
			_ = _index312
			(_1790_v80).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1785_v75, _dafny.Companion_Sequence_.Update(_1785_v75, (Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1785_v75).Cardinality()))).Uint32(), _dafny.CodePoint('i'))), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm17(_1788_v78, globalState), _1785_v75)), (_index312).Int())
			_1688_v2 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1784_v74, _1784_v74))).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1784_v74, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-860), _dafny.IntOfUint32((_1784_v74).Cardinality()))).Uint32(), _1686_v0)))
		}
		if (func() bool {
			if _1688_v2 {
				return _1688_v2
			}
			return true
		})() {
			var _1793_v81 _dafny.Sequence
			_ = _1793_v81
			_1793_v81 = _dafny.UnicodeSeqOfUtf8Bytes("cfusokk")
			var _1794_v82 _dafny.CodePoint
			_ = _1794_v82
			_1794_v82 = _dafny.CodePoint('y')
			var _1795_v83 _dafny.Map
			_ = _1795_v83
			_1795_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v82, _1688_v2)
			var _1796_v84 _dafny.Sequence
			_ = _1796_v84
			_1796_v84 = _dafny.SeqOf((_1795_v83).Cardinality())
			var _1797_v85 _dafny.Map
			_ = _1797_v85
			_1797_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _1688_v2)
			var _1798_v86 _dafny.MultiSet
			_ = _1798_v86
			_1798_v86 = _dafny.MultiSetOf((_1797_v85).Cardinality())
			var _1799_v87 _dafny.Array
			_ = _1799_v87
			var _nwElement0_59 bool = _1688_v2
			_ = _nwElement0_59
			var _nw284 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(18))
			_ = _nw284
			(_nw284).ArraySet1(_nwElement0_59, 0)
			(_nw284).ArraySet1(true, 1)
			(_nw284).ArraySet1(_1688_v2, 2)
			(_nw284).ArraySet1(_1688_v2, 3)
			(_nw284).ArraySet1(_1688_v2, 4)
			(_nw284).ArraySet1(_1688_v2, 5)
			(_nw284).ArraySet1(_1688_v2, 6)
			(_nw284).ArraySet1(_1688_v2, 7)
			(_nw284).ArraySet1(_1688_v2, 8)
			(_nw284).ArraySet1(_1688_v2, 9)
			(_nw284).ArraySet1(false, 10)
			(_nw284).ArraySet1(_1688_v2, 11)
			(_nw284).ArraySet1(_1688_v2, 12)
			(_nw284).ArraySet1(false, 13)
			(_nw284).ArraySet1((_this).Fm3(_1796_v84, _1798_v86, _1688_v2, globalState), 14)
			(_nw284).ArraySet1(_1688_v2, 15)
			(_nw284).ArraySet1(_1688_v2, 16)
			(_nw284).ArraySet1(_1688_v2, 17)
			_1799_v87 = _nw284
			var _1800_v88 _dafny.Map
			_ = _1800_v88
			_1800_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1799_v87)
			var _1801_v89 _dafny.Map
			_ = _1801_v89
			_1801_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1793_v81, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(565))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}(func(_1802_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))), _1800_v88)
			_1801_v89 = (_1801_v89).Update(_1793_v81, _1800_v88)
			if _1688_v2 {
				_1793_v81 = _dafny.Companion_Sequence_.Concatenate(_1793_v81, _1793_v81)
				(globalState).F3 = _1686_v0
				_1793_v81 = _1793_v81
				var _1803_v90 _dafny.Array
				_ = _1803_v90
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_46
				var _nw285 _dafny.Array
				_ = _nw285
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw285 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) D10 = (func(_1804_v82 _dafny.CodePoint) func(_dafny.Int) D10 {
						return func(_1805_i7 _dafny.Int) D10 {
							return Companion_D10_.Create_DC23_(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(93))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg79 _dafny.Int) interface{} {
									return coer79(arg79)
								}
							}((func(_1806_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1807_i8 _dafny.Int) _dafny.CodePoint {
									return _1806_v82
								}
							})(_1804_v82))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-809))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg80 _dafny.Int) interface{} {
									return coer80(arg80)
								}
							}((func(_1808_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1809_i9 _dafny.Int) _dafny.CodePoint {
									return _1808_v82
								}
							})(_1804_v82)))))
						}
					})(_1794_v82)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw285 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw285).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw285).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1803_v90 = _nw285
				var _1810_v91 _dafny.MultiSet
				_ = _1810_v91
				_1810_v91 = _dafny.MultiSetOf(_1793_v81)
				var _1811_v92 D10
				_ = _1811_v92
				_1811_v92 = Companion_D10_.Create_DC23_(_1810_v91)
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1803_v90), 0))
				_ = _index313
				(_1803_v90).ArraySet1(_1811_v92, (_index313).Int())
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1803_v90), 0))
				_ = _index314
				(_1803_v90).ArraySet1(_1811_v92, (_index314).Int())
				(globalState).F4 = (_1686_v0).Cmp(Companion_Default___.SafeModuloInt(_1686_v0, _dafny.IntOfUint32((Companion_Default___.Fm25(globalState)).Cardinality()))) == 0
			} else {
				(globalState).F3 = _1686_v0
				(globalState).F8 = _1686_v0
				var _1812_v93 _dafny.Set
				_ = _1812_v93
				_1812_v93 = _dafny.SetOf(_1686_v0, _1686_v0, (_dafny.IntOfUint32((_1793_v81).Cardinality())).Plus(_1686_v0), _1686_v0)
				(globalState).F8 = (_dafny.Zero).Minus((_1812_v93).Cardinality())
				var _1813_v94 _dafny.Sequence
				_ = _1813_v94
				_1813_v94 = _dafny.SeqOf(_1688_v2)
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1799_v87), 0))
				_ = _index315
				(_1799_v87).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_1813_v94, (_1813_v94).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1813_v94).Cardinality()))).Uint32()).(bool))), (_index315).Int())
				var _1814_v95 _dafny.Set
				_ = _1814_v95
				_1814_v95 = _dafny.SetOf(_1688_v2, _1688_v2)
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1799_v87), 0))
				_ = _index316
				var _rhs311 bool = (func() bool {
					if _1688_v2 {
						return _1688_v2
					}
					return _1688_v2
				})()
				_ = _rhs311
				var _rhs312 _dafny.Set = _1814_v95
				_ = _rhs312
				var _lhs245 _dafny.Array = _1799_v87
				_ = _lhs245
				var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1799_v87), 0))
				_ = _lhs246
				(_lhs245).ArraySet1(_rhs311, (_lhs246).Int())
				_1814_v95 = _rhs312
				_1797_v85 = _1797_v85
			}
			var _rhs313 bool = true
			_ = _rhs313
			var _rhs314 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1686_v0, (_1686_v0).Minus((_dafny.Zero).Minus(_1686_v0))))
			_ = _rhs314
			var _lhs247 *GlobalState = globalState
			_ = _lhs247
			_1688_v2 = _rhs313
			_lhs247.F8 = _rhs314
			var _1815_v96 _dafny.Array
			_ = _1815_v96
			var _nw286 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw286
			_1815_v96 = _nw286
			var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1815_v96), 0))
			_ = _index317
			(_1815_v96).ArraySet1(_1686_v0, (_index317).Int())
			var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1815_v96), 0))
			_ = _index318
			(_1815_v96).ArraySet1(_1686_v0, (_index318).Int())
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1815_v96), 0))
			_ = _index319
			var _rhs315 _dafny.Int = (_1815_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1815_v96), 0))).Int()).(_dafny.Int)
			_ = _rhs315
			var _rhs316 _dafny.Array = _1815_v96
			_ = _rhs316
			var _lhs248 _dafny.Array = _1815_v96
			_ = _lhs248
			var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1815_v96), 0))
			_ = _lhs249
			(_lhs248).ArraySet1(_rhs315, (_lhs249).Int())
			_1815_v96 = _rhs316
		} else {
			var _1816_v97 _dafny.Array
			_ = _1816_v97
			var _nw287 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
			_ = _nw287
			_1816_v97 = _nw287
			var _1817_v98 _dafny.Set
			_ = _1817_v98
			_1817_v98 = _dafny.SetOf(_1686_v0)
			var _1818_v99 _dafny.Sequence
			_ = _1818_v99
			_1818_v99 = _dafny.UnicodeSeqOfUtf8Bytes("xxcadue")
			var _1819_v100 _dafny.Map
			_ = _1819_v100
			_1819_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1817_v98, _1818_v99)
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1816_v97), 0))
			_ = _index320
			(_1816_v97).ArraySet1(_1819_v100, (_index320).Int())
			var _1820_v101 _dafny.Array
			_ = _1820_v101
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_47
			var _nw288 _dafny.Array
			_ = _nw288
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw288 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) D7 = func(_1821_i10 _dafny.Int) D7 {
					return Companion_D7_.Create_DC18_()
				}
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw288 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw288).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw288).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1820_v101 = _nw288
			var _1822_v102 D7
			_ = _1822_v102
			_1822_v102 = Companion_D7_.Create_DC18_()
			var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1820_v101), 0))
			_ = _index321
			(_1820_v101).ArraySet1((func() D7 {
				if _1688_v2 {
					return _1822_v102
				}
				return _1822_v102
			})(), (_index321).Int())
			var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1816_v97), 0))
			_ = _index322
			var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1820_v101), 0))
			_ = _index323
			var _rhs317 _dafny.Map = ((_1819_v100).Merge(Companion_Default___.Fm47(_dafny.IntOfUint32((_1818_v99).Cardinality()), _1688_v2, _1688_v2, globalState))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1817_v98, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(53))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1823_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}))))
			_ = _rhs317
			var _rhs318 bool = !(_1688_v2) || (_1688_v2)
			_ = _rhs318
			var _rhs319 D7 = _1822_v102
			_ = _rhs319
			var _lhs250 _dafny.Array = _1816_v97
			_ = _lhs250
			var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1816_v97), 0))
			_ = _lhs251
			var _lhs252 _dafny.Array = _1820_v101
			_ = _lhs252
			var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1820_v101), 0))
			_ = _lhs253
			(_lhs250).ArraySet1(_rhs317, (_lhs251).Int())
			r0 = _rhs318
			(_lhs252).ArraySet1(_rhs319, (_lhs253).Int())
			var _1824_v103 *C1
			_ = _1824_v103
			var _nw289 *C1 = New_C1_()
			_ = _nw289
			_nw289.Ctor__()
			_1824_v103 = _nw289
			r1 = true
			var _1825_v104 _dafny.Map
			_ = _1825_v104
			_1825_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1688_v2) || (_1688_v2), _1686_v0)
			_1825_v104 = _1825_v104
			var _1826_v105 _dafny.Sequence
			_ = _1826_v105
			_1826_v105 = _dafny.SeqOf(_1825_v104)
			var _1827_v106 _dafny.Sequence
			_ = _1827_v106
			_1827_v106 = _dafny.SeqOf(_1686_v0, _1686_v0, _1686_v0, _1686_v0, _1686_v0)
			var _1828_v107 _dafny.Sequence
			_ = _1828_v107
			_1828_v107 = _dafny.SeqOf((_1826_v105).Select((Companion_Default___.SafeIndex(_1686_v0, _dafny.IntOfUint32((_1826_v105).Cardinality()))).Uint32()).(_dafny.Map), _1825_v104, _1825_v104, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1688_v2, _dafny.IntOfUint32((_1827_v106).Cardinality()))).Merge(_1825_v104))
			var _1829_v108 _dafny.Map
			_ = _1829_v108
			_1829_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v0, _1688_v2)
			_1828_v107 = _dafny.Companion_Sequence_.Update(_1826_v105, (Companion_Default___.SafeIndex(((_1829_v108).Update(_dafny.IntOfInt64(507), _1688_v2)).Cardinality(), _dafny.IntOfUint32((_1826_v105).Cardinality()))).Uint32(), Companion_Default___.Fm46(globalState))
		}
		var _1830_v109 _dafny.Array
		_ = _1830_v109
		var _len0_48 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_48
		var _nw290 _dafny.Array
		_ = _nw290
		if _len0_48.Cmp(_dafny.Zero) == 0 {
			_nw290 = _dafny.NewArray(_len0_48)
		} else {
			var _init48 func(_dafny.Int) _dafny.CodePoint = func(_1831_i13 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}
			_ = _init48
			var _element0_48 = _init48(_dafny.Zero)
			_ = _element0_48
			_nw290 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
			(_nw290).ArraySet1CodePoint(_element0_48, 0)
			var _nativeLen0_48 = (_len0_48).Int()
			_ = _nativeLen0_48
			for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
				(_nw290).ArraySet1CodePoint(_init48(_dafny.IntOf(_i0_48)), _i0_48)
			}
		}
		_1830_v109 = _nw290
		for _iter99 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1830_v109), 0))); ; {
			_guard_loop_3, _ok99 := _iter99()
			if !_ok99 {
				break
			}
			var _1832_i12 _dafny.Int
			_1832_i12 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1832_i12).Sign() != -1) && ((_1832_i12).Cmp(_dafny.ArrayLen((_1830_v109), 0)) < 0)) {
				(_1830_v109).ArraySet1CodePoint(_dafny.CodePoint('h'), (_1832_i12).Int())
			}
		}
		r0 = !(_1688_v2)
		var _1833_v110 _dafny.Map
		_ = _1833_v110
		_1833_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1687_v1, _1686_v0)
		var _1834_v111 _dafny.Set
		_ = _1834_v111
		_1834_v111 = _dafny.SetOf(_1833_v110)
		r1 = (_1834_v111).IsProperSubsetOf(_1834_v111)
		return r0, r1
	}
}
func (_this *C7) M5(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Set, bool) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var r2 bool = false
		_ = r2
		(globalState).F8 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1835_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality())
		var _1836_v0 _dafny.Set
		_ = _1836_v0
		_1836_v0 = _dafny.SetOf(p2, p2)
		var _1837_v1 _dafny.Array
		_ = _1837_v1
		var _len0_49 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_49
		var _nw291 _dafny.Array
		_ = _nw291
		if _len0_49.Cmp(_dafny.Zero) == 0 {
			_nw291 = _dafny.NewArray(_len0_49)
		} else {
			var _init49 func(_dafny.Int) bool = func(_1838_i2 _dafny.Int) bool {
				return true
			}
			_ = _init49
			var _element0_49 = _init49(_dafny.Zero)
			_ = _element0_49
			_nw291 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
			(_nw291).ArraySet1(_element0_49, 0)
			var _nativeLen0_49 = (_len0_49).Int()
			_ = _nativeLen0_49
			for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
				(_nw291).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
			}
		}
		_1837_v1 = _nw291
		var _1839_v2 _dafny.Map
		_ = _1839_v2
		_1839_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1837_v1, (_dafny.Zero).Minus(p2))
		var _1840_i1 _dafny.Int
		_ = _1840_i1
		_1840_i1 = _dafny.Zero
		{
			for (_dafny.SetOf((_1839_v2).Cardinality())).IsSubsetOf((_1836_v0).Difference(_1836_v0)) {
				{
					if (_1840_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1840_i1 = (_1840_i1).Plus(_dafny.One)
					var _1841_v3 _dafny.Map
					_ = _1841_v3
					_1841_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(749), p3)
					var _1842_v4 D3
					_ = _1842_v4
					_1842_v4 = Companion_D3_.Create_DC8_(_1841_v3)
					var _1843_v5 _dafny.Sequence
					_ = _1843_v5
					_1843_v5 = _dafny.SeqOf(p0, p0)
					var _1844_v6 _dafny.MultiSet
					_ = _1844_v6
					_1844_v6 = _dafny.MultiSetOf(_1843_v5)
					(globalState).F8 = (Companion_Default___.Fm26(_1842_v4, p1, _1844_v6, globalState)).Minus(p1)
					var _source32 bool = p0
					_ = _source32
					var _1845___mcc_h0 bool = _source32
					_ = _1845___mcc_h0
					var _1846_cf4 bool = _1845___mcc_h0
					_ = _1846_cf4
					var _1847_v7 _dafny.Sequence
					_ = _1847_v7
					_1847_v7 = _dafny.UnicodeSeqOfUtf8Bytes("cwtkmlf")
					var _1848_v8 _dafny.MultiSet
					_ = _1848_v8
					_1848_v8 = _dafny.MultiSetOf(_1847_v7)
					(globalState).F0 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(p1, (_1848_v8).Cardinality()), (func() _dafny.Int {
						if (_1841_v3).Contains(p1) {
							return (_1841_v3).Get(p1).(_dafny.Int)
						}
						return p2
					})(), Companion_Default___.Fm20(p0, globalState), p3)
					(globalState).F4 = p0
					var _nw292 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
					_ = _nw292
					_1837_v1 = _nw292
					_1847_v7 = Companion_Default___.Fm14(p3, globalState)
					var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1837_v1), 0))
					_ = _index324
					(_1837_v1).ArraySet1(p0, (_index324).Int())
					var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1837_v1), 0))
					_ = _index325
					(_1837_v1).ArraySet1(p0, (_index325).Int())
					var _1849_v9 _dafny.CodePoint
					_ = _1849_v9
					_1849_v9 = _dafny.CodePoint('e')
					var _1850_v10 _dafny.Sequence
					_ = _1850_v10
					_1850_v10 = _dafny.UnicodeSeqOfUtf8Bytes("jrrugumfh")
					var _1851_v11 _dafny.Map
					_ = _1851_v11
					_1851_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1849_v9, _1850_v10)
					var _1852_v12 _dafny.Sequence
					_ = _1852_v12
					_1852_v12 = _dafny.SeqOf(p2, p2, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_1851_v11).Contains(_1849_v9) {
							return (_1851_v11).Get(_1849_v9).(_dafny.Sequence)
						}
						return _1850_v10
					})()).Cardinality()), p2, _dafny.IntOfUint32((_1850_v10).Cardinality()))
					(globalState).F0 = _dafny.Companion_Sequence_.Concatenate(_1852_v12, _1852_v12)
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _1853_v13 _dafny.Sequence
		_ = _1853_v13
		_1853_v13 = _dafny.UnicodeSeqOfUtf8Bytes("uba")
		var _1854_v14 _dafny.Set
		_ = _1854_v14
		_1854_v14 = _dafny.SetOf(_1853_v13, _1853_v13)
		var _1855_v15 D13
		_ = _1855_v15
		_1855_v15 = Companion_D13_.Create_DC29_(_1854_v14)
		var _1856_v16 _dafny.Map
		_ = _1856_v16
		_1856_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1855_v15)
		var _1857_v17 D13
		_ = _1857_v17
		_1857_v17 = Companion_D13_.Create_DC31_((func() D13 {
			if (_1856_v16).Contains(p0) {
				return (_1856_v16).Get(p0).(D13)
			}
			return _1855_v15
		})())
		var _source33 D13 = _1857_v17
		_ = _source33
		if _source33.Is_DC30() {
			var _1858___mcc_h1 bool = _source33.Get_().(D13_DC30).Cf42
			_ = _1858___mcc_h1
			var _1859___mcc_h2 bool = _source33.Get_().(D13_DC30).Cf43
			_ = _1859___mcc_h2
			var _1860___mcc_h3 _dafny.Set = _source33.Get_().(D13_DC30).Cf44
			_ = _1860___mcc_h3
			var _1861___mcc_h4 *C0 = _source33.Get_().(D13_DC30).Cf45
			_ = _1861___mcc_h4
			var _1862_cf45 *C0 = _1861___mcc_h4
			_ = _1862_cf45
			var _1863_cf44 _dafny.Set = _1860___mcc_h3
			_ = _1863_cf44
			var _1864_cf43 bool = _1859___mcc_h2
			_ = _1864_cf43
			var _1865_cf42 bool = _1858___mcc_h1
			_ = _1865_cf42
			(globalState).F3 = ((p3).Times(p3)).Plus(p1)
			_1853_v13 = _1853_v13
			if p0 {
				(globalState).F4 = !(p0) || (p0)
				var _1866_v18 _dafny.MultiSet
				_ = _1866_v18
				_1866_v18 = _dafny.MultiSetOf(_dafny.IntOfInt64(638))
				var _1867_v19 _dafny.Map
				_ = _1867_v19
				_1867_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Times(p3), (_1866_v18).IsSubsetOf(_1866_v18))
				var _1868_v20 _dafny.Sequence
				_ = _1868_v20
				_1868_v20 = _dafny.SeqOf(p1)
				var _1869_v21 _dafny.Map
				_ = _1869_v21
				_1869_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1865_cf42, _dafny.MultiSetOf(_dafny.IntOfUint32((_1868_v20).Cardinality())))
				_1867_v19 = (_1867_v19).Update(Companion_Default___.SafeDivisionInt(p3, p3), ((func() _dafny.MultiSet {
					if (_1869_v21).Contains(_1864_cf43) {
						return (_1869_v21).Get(_1864_cf43).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_1868_v20)
				})()).IsProperSubsetOf(_1866_v18))
				var _1870_v22 _dafny.Map
				_ = _1870_v22
				_1870_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _1865_cf42)
				_1870_v22 = (_1870_v22).Update(_dafny.CodePoint('n'), _1865_cf42)
				var _1871_v23 T0
				_ = _1871_v23
				var _nw293 *C2 = New_C2_()
				_ = _nw293
				_nw293.Ctor__()
				_1871_v23 = _nw293
				var _1872_v24 _dafny.Map
				_ = _1872_v24
				_1872_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm3(_1868_v20, _dafny.MultiSetFromSeq(_1868_v20), _1865_cf42, globalState)), _dafny.CodePoint('s'))
				var _1873_v25 _dafny.Map
				_ = _1873_v25
				_1873_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1865_cf42, _1853_v13)
				var _1874_v26 D19
				_ = _1874_v26
				_1874_v26 = Companion_D19_.Create_DC39_(_1865_cf42, p1)
				var _1875_v27 _dafny.Array
				_ = _1875_v27
				var _nw294 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw294
				_1875_v27 = _nw294
				var _1876_v28 _dafny.Map
				_ = _1876_v28
				_1876_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1875_v27)
				var _1877_v29 _dafny.Map
				_ = _1877_v29
				_1877_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3)
				var _1878_v30 _dafny.Sequence
				_ = _1878_v30
				_1878_v30 = _dafny.SeqOf((func() _dafny.Array {
					if (_1876_v28).Contains((func() _dafny.Int {
						if (_1877_v29).Contains(_1864_cf43) {
							return (_1877_v29).Get(_1864_cf43).(_dafny.Int)
						}
						return p3
					})()) {
						return (_1876_v28).Get((func() _dafny.Int {
							if (_1877_v29).Contains(_1864_cf43) {
								return (_1877_v29).Get(_1864_cf43).(_dafny.Int)
							}
							return p3
						})()).(_dafny.Array)
					}
					return _1875_v27
				})())
				var _1879_v31 _dafny.Map
				_ = _1879_v31
				_1879_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1874_v26).Dtor_cf58(), (_1878_v30).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1865_cf42)).Cardinality(), _dafny.IntOfUint32((_1878_v30).Cardinality()))).Uint32()).(_dafny.Array))
				var _1880_v32 _dafny.Map
				_ = _1880_v32
				_1880_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1879_v31, (func() T0 {
					if p0 {
						return _1871_v23
					}
					return _1871_v23
				})())
				var _1881_v33 _dafny.CodePoint
				_ = _1881_v33
				_1881_v33 = _dafny.CodePoint('p')
				var _1882_v34 _dafny.Map
				_ = _1882_v34
				_1882_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1881_v33)
				var _rhs320 _dafny.Set = (func() _dafny.Set {
					if (p2).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_1873_v25).Contains(!(p0)) {
							return (_1873_v25).Get(!(p0)).(_dafny.Sequence)
						}
						return _1853_v13
					})()).Cardinality())) != 0 {
						return _1863_cf44
					}
					return _1863_cf44
				})()
				_ = _rhs320
				var _rhs321 T0 = (func() T0 {
					if (_1880_v32).Contains(_1879_v31) {
						return (_1880_v32).Get(_1879_v31).(T0)
					}
					return (func() T0 {
						if _1864_cf43 {
							return _1871_v23
						}
						return _1871_v23
					})()
				})()
				_ = _rhs321
				var _rhs322 bool = (Companion_Default___.Fm43(_1864_cf43, p3, _1881_v33, globalState)).Equals(_1882_v34)
				_ = _rhs322
				var _rhs323 _dafny.Map = (Companion_Default___.Fm11(_1865_cf42, _1864_cf43, globalState)).Merge(_1872_v24)
				_ = _rhs323
				var _rhs324 bool = p0
				_ = _rhs324
				var _lhs254 *GlobalState = globalState
				_ = _lhs254
				r1 = _rhs320
				_1871_v23 = _rhs321
				_lhs254.F4 = _rhs322
				_1872_v24 = _rhs323
				r2 = _rhs324
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_1875_v27), 0))
				_ = _index326
				(_1875_v27).ArraySet1(_dafny.IntOfInt64(462), (_index326).Int())
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_1875_v27), 0))
				_ = _index327
				(_1875_v27).ArraySet1((func() _dafny.Int {
					if false {
						return (_dafny.IntOfInt64(680)).Minus(p3)
					}
					return p2
				})(), (_index327).Int())
			} else {
				var _1883_v35 _dafny.Array
				_ = _1883_v35
				var _nw295 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw295
				_1883_v35 = _nw295
				var _1884_v36 _dafny.Array
				_ = _1884_v36
				_1884_v36 = _1883_v35
				_1884_v36 = _1884_v36
				var _1885_v37 *C2
				_ = _1885_v37
				var _nw296 *C2 = New_C2_()
				_ = _nw296
				_nw296.Ctor__()
				_1885_v37 = _nw296
				var _1886_v38 _dafny.Map
				_ = _1886_v38
				_1886_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1885_v37, _1853_v13)
				var _1887_v39 _dafny.MultiSet
				_ = _1887_v39
				_1887_v39 = _dafny.MultiSetOf(false)
				var _1888_v40 _dafny.Map
				_ = _1888_v40
				_1888_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(p1) > 0, (_1887_v39).IsProperSubsetOf(_1887_v39))
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))
				_ = _index328
				(_1883_v35).ArraySet1((p2).Times(_dafny.IntOfInt64(-434)), (_index328).Int())
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))
				_ = _index329
				var _rhs325 _dafny.Map = _1886_v38
				_ = _rhs325
				var _rhs326 bool = p0
				_ = _rhs326
				var _rhs327 _dafny.Map = ((_1888_v40).Merge(_1888_v40)).Update(!(_1864_cf43), (p1).Cmp(p2) > 0)
				_ = _rhs327
				var _rhs328 _dafny.Int = p2
				_ = _rhs328
				var _lhs255 *GlobalState = globalState
				_ = _lhs255
				var _lhs256 _dafny.Array = _1883_v35
				_ = _lhs256
				var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))
				_ = _lhs257
				_1886_v38 = _rhs325
				_lhs255.F4 = _rhs326
				_1888_v40 = _rhs327
				(_lhs256).ArraySet1(_rhs328, (_lhs257).Int())
				_1865_cf42 = _1864_cf43
				var _1889_v41 _dafny.Sequence
				_ = _1889_v41
				_1889_v41 = _dafny.SeqOf(_dafny.IntOfInt64(203), p2, p3)
				var _1890_v42 _dafny.MultiSet
				_ = _1890_v42
				_1890_v42 = _dafny.MultiSetOf(p2)
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))
				_ = _index330
				(_1883_v35).ArraySet1(((_1889_v41).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1890_v42).Contains((_1883_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))).Int()).(_dafny.Int)) {
						return (_1890_v42).Multiplicity((_1883_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))).Int()).(_dafny.Int))
					}
					return (_1883_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_1889_v41).Cardinality()))).Uint32()).(_dafny.Int)).Times((_dafny.IntOfInt64(-356)).Plus(p2)), (_index330).Int())
				var _1891_v43 _dafny.MultiSet
				_ = _1891_v43
				_1891_v43 = _1890_v42
				var _1892_v44 _dafny.MultiSet
				_ = _1892_v44
				_1892_v44 = _dafny.MultiSetOf(_1891_v43, _1891_v43, _1891_v43)
				var _1893_v45 _dafny.Map
				_ = _1893_v45
				_1893_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1887_v39, (_1892_v44).Cardinality())
				var _1894_v46 D12
				_ = _1894_v46
				_1894_v46 = Companion_D12_.Create_DC27_(p1, p2, _dafny.IntOfInt64(826), (_1893_v45).Cardinality())
				var _1895_v47 _dafny.Map
				_ = _1895_v47
				_1895_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _1894_v46)
				var _1896_v48 _dafny.CodePoint
				_ = _1896_v48
				_1896_v48 = _dafny.CodePoint('r')
				var _1897_v49 _dafny.Map
				_ = _1897_v49
				_1897_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm31(_1865_cf42, p0, _1864_cf43, globalState)).Cardinality()), p3)
				var _1898_v50 D3
				_ = _1898_v50
				_1898_v50 = Companion_D3_.Create_DC8_(_1897_v49)
				_1895_v47 = (_1895_v47).Update(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm1((_1883_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1883_v35), 0))).Int()).(_dafny.Int), _1896_v48, _1864_cf43, (_1898_v50).Dtor_cf12(), globalState), p2), _1894_v46)
			}
			(globalState).F8 = p2
		} else if _source33.Is_DC29() {
			var _1899___mcc_h5 _dafny.Set = _source33.Get_().(D13_DC29).Cf41
			_ = _1899___mcc_h5
			var _1900_cf41 _dafny.Set = _1899___mcc_h5
			_ = _1900_cf41
			(globalState).F8 = (_this).Fm10(globalState)
			var _1901_v51 _dafny.Array
			_ = _1901_v51
			var _nw297 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw297
			_1901_v51 = _nw297
			var _1902_v52 _dafny.Array
			_ = _1902_v52
			var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw298
			_1902_v52 = _nw298
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1901_v51), 0))
			_ = _index331
			(_1901_v51).ArraySet1(_1902_v52, (_index331).Int())
			var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1901_v51), 0))
			_ = _index332
			(_1901_v51).ArraySet1(_1902_v52, (_index332).Int())
			var _1903_v53 _dafny.Array
			_ = _1903_v53
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_50
			var _nw299 _dafny.Array
			_ = _nw299
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw299 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Int = (func(_1904_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_1905_i3 _dafny.Int) _dafny.Int {
						return (_1905_i3).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1904_p0)).Cardinality())
					}
				})(p0)
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw299 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw299).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw299).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1903_v53 = _nw299
			var _1906_v54 D21
			_ = _1906_v54
			_1906_v54 = Companion_D21_.Create_DC42_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1903_v53, _dafny.IntOfInt64(731)))
			(_this).M6(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(724), _dafny.IntOfUint32((_1853_v13).Cardinality())), _1837_v1, ((_1906_v54).Dtor_cf62()).Cardinality(), globalState)
			(globalState).F8 = p1
		} else {
			var _1907___mcc_h6 D13 = _source33.Get_().(D13_DC31).Cf46
			_ = _1907___mcc_h6
			var _1908_cf46 D13 = _1907___mcc_h6
			_ = _1908_cf46
			_1853_v13 = _1853_v13
			var _1909_v55 *C1
			_ = _1909_v55
			var _nw300 *C1 = New_C1_()
			_ = _nw300
			_nw300.Ctor__()
			_1909_v55 = _nw300
			var _1910_v56 D4
			_ = _1910_v56
			_1910_v56 = Companion_D4_.Create_DC11_(p0, true, _dafny.IntOfInt64(-284), _1853_v13)
			var _1911_v57 D4
			_ = _1911_v57
			_1911_v57 = Companion_D4_.Create_DC12_(_1910_v56)
			var _1912_v58 D4
			_ = _1912_v58
			_1912_v58 = Companion_D4_.Create_DC12_(_1910_v56)
			var _1913_v59 D4
			_ = _1913_v59
			_1913_v59 = Companion_D4_.Create_DC12_(_1910_v56)
			var _source34 D4 = (func() D4 {
				if (func() bool {
					if p0 {
						return false
					}
					return p0
				})() {
					return _1913_v59
				}
				return _1913_v59
			})()
			_ = _source34
			if _source34.Is_DC11() {
				var _1914___mcc_h7 bool = _source34.Get_().(D4_DC11).Cf14
				_ = _1914___mcc_h7
				var _1915___mcc_h8 bool = _source34.Get_().(D4_DC11).Cf15
				_ = _1915___mcc_h8
				var _1916___mcc_h9 _dafny.Int = _source34.Get_().(D4_DC11).Cf16
				_ = _1916___mcc_h9
				var _1917___mcc_h10 _dafny.Sequence = _source34.Get_().(D4_DC11).Cf17
				_ = _1917___mcc_h10
				var _1918_cf17 _dafny.Sequence = _1917___mcc_h10
				_ = _1918_cf17
				var _1919_cf16 _dafny.Int = _1916___mcc_h9
				_ = _1919_cf16
				var _1920_cf15 bool = _1915___mcc_h8
				_ = _1920_cf15
				var _1921_cf14 bool = _1914___mcc_h7
				_ = _1921_cf14
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1837_v1), 0))
				_ = _index333
				(_1837_v1).ArraySet1(false, (_index333).Int())
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1837_v1), 0))
				_ = _index334
				(_1837_v1).ArraySet1(true, (_index334).Int())
				var _1922_v60 _dafny.Array
				_ = _1922_v60
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_51
				var _nw301 _dafny.Array
				_ = _nw301
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw301 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Map = (func(_1923_cf14 bool, _1924_p3 _dafny.Int, _1925_p2 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1926_i4 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1923_cf14, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1924_p3, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_1925_p2)).Cardinality()))).Cardinality())).Cardinality())
						}
					})(_1921_cf14, p3, p2)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw301 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw301).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw301).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1922_v60 = _nw301
				var _1927_v61 *C4
				_ = _1927_v61
				var _nw302 *C4 = New_C4_()
				_ = _nw302
				_nw302.Ctor__(_1922_v60)
				_1927_v61 = _nw302
				_1922_v60 = (func() _dafny.Array {
					if (_1837_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1837_v1), 0))).Int()).(bool) {
						return _1922_v60
					}
					return (_1927_v61).F12()
				})()
				var _1928_v62 *C3
				_ = _1928_v62
				var _nw303 *C3 = New_C3_()
				_ = _nw303
				_nw303.Ctor__()
				_1928_v62 = _nw303
			} else if _source34.Is_DC10() {
				var _1929___mcc_h11 _dafny.Sequence = _source34.Get_().(D4_DC10).Cf13
				_ = _1929___mcc_h11
				var _1930_cf13 _dafny.Sequence = _1929___mcc_h11
				_ = _1930_cf13
				var _1931_v63 _dafny.CodePoint
				_ = _1931_v63
				_1931_v63 = _dafny.CodePoint('k')
				(globalState).F4 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1853_v13, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1853_v13).Cardinality()), _dafny.IntOfUint32((_1853_v13).Cardinality()))).Uint32(), _1931_v63)).Cardinality())).Cmp(p1) != 0
				var _1932_v64 *C3
				_ = _1932_v64
				var _nw304 *C3 = New_C3_()
				_ = _nw304
				_nw304.Ctor__()
				_1932_v64 = _nw304
				var _1933_v65 _dafny.MultiSet
				_ = _1933_v65
				_1933_v65 = _dafny.MultiSetOf(_1932_v64)
				var _1934_v66 _dafny.MultiSet
				_ = _1934_v66
				_1934_v66 = _dafny.MultiSetOf(p0)
				var _1935_v67 _dafny.MultiSet
				_ = _1935_v67
				_1935_v67 = _1934_v66
				var _1936_v68 _dafny.Map
				_ = _1936_v68
				_1936_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1935_v67)
				var _rhs329 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}(func(_1937_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))
				_ = _rhs329
				var _rhs330 _dafny.MultiSet = _1933_v65
				_ = _rhs330
				var _rhs331 _dafny.Int = ((_dafny.Zero).Minus(p3)).Times((_1936_v68).Cardinality())
				_ = _rhs331
				var _lhs258 *GlobalState = globalState
				_ = _lhs258
				_1853_v13 = _rhs329
				_1933_v65 = _rhs330
				_lhs258.F8 = _rhs331
				var _1938_v69 _dafny.Map
				_ = _1938_v69
				_1938_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
				var _1939_v70 _dafny.MultiSet
				_ = _1939_v70
				_1939_v70 = _dafny.MultiSetOf(_dafny.SeqOf(p0, p0))
				var _1940_v71 _dafny.MultiSet
				_ = _1940_v71
				_1940_v71 = _dafny.MultiSetOf(Companion_Default___.Fm26(Companion_D3_.Create_DC8_(_1938_v69), p2, _1939_v70, globalState), p1, p1)
				(_1909_v55).M13(_dafny.IntOfInt64(768), _1940_v71, p0, globalState)
				var _rhs332 bool = !(p0)
				_ = _rhs332
				var _rhs333 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, p3)
				_ = _rhs333
				var _rhs334 _dafny.CodePoint = _1931_v63
				_ = _rhs334
				var _rhs335 _dafny.Int = Companion_Default___.SafeDivisionInt(p2, p3)
				_ = _rhs335
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				var _lhs260 *GlobalState = globalState
				_ = _lhs260
				var _lhs261 *GlobalState = globalState
				_ = _lhs261
				_lhs259.F4 = _rhs332
				_lhs260.F8 = _rhs333
				r0 = _rhs334
				_lhs261.F3 = _rhs335
			} else {
				var _1941___mcc_h12 D4 = _source34.Get_().(D4_DC12).Cf18
				_ = _1941___mcc_h12
				var _1942_cf18 D4 = _1941___mcc_h12
				_ = _1942_cf18
				var _1943_v72 T0
				_ = _1943_v72
				var _nw305 *C1 = New_C1_()
				_ = _nw305
				_nw305.Ctor__()
				_1943_v72 = _nw305
				var _1944_v73 D2
				_ = _1944_v73
				_1944_v73 = Companion_D2_.Create_DC6_((_1836_v0).Cardinality(), p0, p0)
				var _1945_v74 _dafny.Map
				_ = _1945_v74
				_1945_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1944_v73).Dtor_cf7(), _1943_v72)
				_1943_v72 = (func() T0 {
					if p0 {
						return _1943_v72
					}
					return (func() T0 {
						if (_1945_v74).Contains(p0) {
							return (_1945_v74).Get(p0).(T0)
						}
						return _1943_v72
					})()
				})()
				var _1946_v75 *C3
				_ = _1946_v75
				var _nw306 *C3 = New_C3_()
				_ = _nw306
				_nw306.Ctor__()
				_1946_v75 = _nw306
				(globalState).F4 = (func() bool {
					if p0 {
						return p0
					}
					return p0
				})()
				var _1947_v76 _dafny.Map
				_ = _1947_v76
				_1947_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1909_v55, p1)
				var _1948_v77 _dafny.Map
				_ = _1948_v77
				_1948_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
				var _1949_v78 _dafny.MultiSet
				_ = _1949_v78
				_1949_v78 = _dafny.MultiSetOf(p3, (_dafny.Zero).Minus(p2), p2, _dafny.IntOfInt64(934), (_1948_v77).Cardinality())
				var _1950_v79 _dafny.Array
				_ = _1950_v79
				var _nwElement0_60 _dafny.Int = ((_1947_v76).Cardinality()).Plus(p2)
				_ = _nwElement0_60
				var _nw307 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(8))
				_ = _nw307
				(_nw307).ArraySet1(_nwElement0_60, 0)
				(_nw307).ArraySet1(p1, 1)
				(_nw307).ArraySet1(p2, 2)
				(_nw307).ArraySet1((p1).Minus(p3), 3)
				(_nw307).ArraySet1(_dafny.IntOfInt64(206), 4)
				(_nw307).ArraySet1((_dafny.IntOfInt64(833)).Minus(_dafny.IntOfUint32((_1853_v13).Cardinality())), 5)
				(_nw307).ArraySet1((_1949_v78).Cardinality(), 6)
				(_nw307).ArraySet1(p3, 7)
				_1950_v79 = _nw307
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1950_v79), 0))
				_ = _index335
				(_1950_v79).ArraySet1(p1, (_index335).Int())
				var _1951_v80 _dafny.Sequence
				_ = _1951_v80
				_1951_v80 = _dafny.SeqOf(p0)
				var _1952_v81 _dafny.Map
				_ = _1952_v81
				_1952_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_1951_v80).Cardinality()))
				var _1953_v82 _dafny.Sequence
				_ = _1953_v82
				_1953_v82 = _dafny.SeqOf((_1946_v75).Fm15(p2, globalState))
				var _1954_v83 _dafny.Sequence
				_ = _1954_v83
				_1954_v83 = _dafny.SeqOf((_1953_v82).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.IntOfUint32((_1953_v82).Cardinality()))).Uint32()).(_dafny.Int), p3, _dafny.IntOfInt64(-919), (_dafny.Zero).Minus(p1), p1)
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1950_v79), 0))
				_ = _index336
				(_1950_v79).ArraySet1((((_1952_v81).Update(_dafny.IntOfUint32((_1953_v82).Cardinality()), _dafny.IntOfInt64(-642))).Cardinality()).Minus((p1).Plus(_dafny.IntOfInt64(160))), (_index336).Int())
			}
			var _1955_v84 _dafny.Map
			_ = _1955_v84
			_1955_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p0))
			_1955_v84 = (_1955_v84).Update(true, p0)
		}
		var _1956_v85 _dafny.Sequence
		_ = _1956_v85
		_1956_v85 = _dafny.SeqOf(p1, _dafny.IntOfInt64(-829), p2, _dafny.IntOfUint32((_1853_v13).Cardinality()))
		(globalState).F3 = (_1956_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.IntOfUint32((_1956_v85).Cardinality()))).Uint32()).(_dafny.Int)
		(globalState).F3 = ((Companion_Default___.Fm28(p1, globalState)).Union(_1836_v0)).Cardinality()
		var _1957_v86 _dafny.Map
		_ = _1957_v86
		_1957_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
		var _1958_v87 D3
		_ = _1958_v87
		_1958_v87 = Companion_D3_.Create_DC8_(_1957_v86)
		var _1959_v89 _dafny.Sequence
		_ = _1959_v89
		_1959_v89 = _dafny.SeqOf(true)
		var _1960_v90 _dafny.MultiSet
		_ = _1960_v90
		_1960_v90 = _dafny.MultiSetOf(p1)
		var _1961_v91 _dafny.MultiSet
		_ = _1961_v91
		_1961_v91 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_1959_v89, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1959_v89).Cardinality()))).Uint32(), true), Companion_Default___.Fm31(p0, p0, (_this).Fm3(_1956_v85, _1960_v90, p0, globalState), globalState))
		var _1962_v92 _dafny.CodePoint
		_ = _1962_v92
		_1962_v92 = _dafny.CodePoint('v')
		var _pat_let_tv66 = _1957_v86
		_ = _pat_let_tv66
		var _1963_v93 *C6
		_ = _1963_v93
		var _nw308 *C6 = New_C6_()
		_ = _nw308
		_nw308.Ctor__(p1, _dafny.Companion_Sequence_.Update(_1853_v13, (Companion_Default___.SafeIndex(Companion_Default___.Fm26(func(_pat_let55_0 D3) D3 {
			return func(_1964_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let56_0 _dafny.Map) D3 {
					return func(_1965_dt__update_hcf12_h0 _dafny.Map) D3 {
						return Companion_D3_.Create_DC8_(_1965_dt__update_hcf12_h0)
					}(_pat_let56_0)
				}(_pat_let_tv66)
			}(_pat_let55_0)
		}(_1958_v87), (func() _dafny.Map {
			var _coll96 = _dafny.NewMapBuilder()
			_ = _coll96
			for _iter100 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(51), _dafny.IntOfInt64(-278))); ; {
				_compr_96, _ok100 := _iter100()
				if !_ok100 {
					break
				}
				var _1966_v88 _dafny.Int
				_1966_v88 = interface{}(_compr_96).(_dafny.Int)
				if ((_dafny.IntOfInt64(51)).Cmp(_1966_v88) <= 0) && ((_1966_v88).Cmp(_dafny.IntOfInt64(-278)) < 0) {
					_coll96.Add(Companion_Default___.SafeModuloInt(_1966_v88, p2), true)
				}
			}
			return _coll96.ToMap()
		}()).Cardinality(), _1961_v91, globalState), _dafny.IntOfUint32((_1853_v13).Cardinality()))).Uint32(), _1962_v92))
		_1963_v93 = _nw308
		r0 = _1962_v92
		r1 = _1836_v0
		r2 = (p3).Cmp(p1) < 0
		return r0, r1, r2
	}
}
func (_this *C7) M6(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _1967_v0 bool
		_ = _1967_v0
		_1967_v0 = true
		var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
		_ = _index337
		(p1).ArraySet1(_1967_v0, (_index337).Int())
		var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
		_ = _index338
		(p1).ArraySet1(false, (_index338).Int())
		var _1968_v1 _dafny.CodePoint
		_ = _1968_v1
		_1968_v1 = _dafny.CodePoint('m')
		var _1969_v2 _dafny.MultiSet
		_ = _1969_v2
		_1969_v2 = _dafny.MultiSetOf(_1968_v1, _1968_v1, _1968_v1, _1968_v1)
		var _1970_v3 _dafny.Sequence
		_ = _1970_v3
		_1970_v3 = _dafny.SeqOf(_1968_v1)
		var _1971_v4 _dafny.MultiSet
		_ = _1971_v4
		_1971_v4 = _1969_v2
		var _1972_v5 _dafny.Map
		_ = _1972_v5
		_1972_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1970_v3).Cardinality()), p2)
		var _1973_v6 D3
		_ = _1973_v6
		_1973_v6 = Companion_D3_.Create_DC8_(_1972_v5)
		var _1974_v7 _dafny.Sequence
		_ = _1974_v7
		_1974_v7 = _dafny.SeqOf(_1967_v0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1967_v0)
		var _1975_v8 _dafny.MultiSet
		_ = _1975_v8
		_1975_v8 = _dafny.MultiSetOf(_1974_v7, _1974_v7)
		var _1976_v9 D2
		_ = _1976_v9
		_1976_v9 = Companion_D2_.Create_DC7_(_1968_v1, false, p2)
		var _pat_let_tv67 = p1
		_ = _pat_let_tv67
		var _pat_let_tv68 = p1
		_ = _pat_let_tv68
		var _pat_let_tv69 = _1968_v1
		_ = _pat_let_tv69
		var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
		_ = _index339
		var _rhs336 _dafny.MultiSet = (_1969_v2).Difference((_dafny.MultiSetFromSeq(_1970_v3)).Difference((_1971_v4)))
		_ = _rhs336
		var _rhs337 bool = (func() bool {
			if (Companion_Default___.Fm26(_1973_v6, p0, _1975_v8, globalState)).Cmp(Companion_Default___.Fm20((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState)) == 0 {
				return (func(_pat_let57_0 D2) D2 {
					return func(_1977_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let58_0 bool) D2 {
							return func(_1978_dt__update_hcf10_h0 bool) D2 {
								return func(_pat_let59_0 _dafny.CodePoint) D2 {
									return func(_1979_dt__update_hcf9_h0 _dafny.CodePoint) D2 {
										return Companion_D2_.Create_DC7_(_1979_dt__update_hcf9_h0, _1978_dt__update_hcf10_h0, (_1977_dt__update__tmp_h0).Dtor_cf11())
									}(_pat_let59_0)
								}(_pat_let_tv69)
							}(_pat_let58_0)
						}((_pat_let_tv68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_pat_let_tv67), 0))).Int()).(bool))
					}(_pat_let57_0)
				}(_1976_v9)).Dtor_cf10()
			}
			return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
		})()
		_ = _rhs337
		var _rhs338 bool = _1967_v0
		_ = _rhs338
		var _lhs262 _dafny.Array = p1
		_ = _lhs262
		var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
		_ = _lhs263
		_1969_v2 = _rhs336
		_1967_v0 = _rhs337
		(_lhs262).ArraySet1(_rhs338, (_lhs263).Int())
		var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
		_ = _index340
		(p1).ArraySet1(!((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)), (_index340).Int())
		var _1980_v10 *C5
		_ = _1980_v10
		var _nw309 *C5 = New_C5_()
		_ = _nw309
		_nw309.Ctor__()
		_1980_v10 = _nw309
		var _1981_v11 _dafny.Map
		_ = _1981_v11
		_1981_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1980_v10)
		var _hi11 _dafny.Int = p2
		_ = _hi11
		for _1982_i0 := (_1981_v11).Cardinality(); _1982_i0.Cmp(_hi11) < 0; _1982_i0 = _1982_i0.Plus(_dafny.One) {
			var _1983_v12 _dafny.Array
			_ = _1983_v12
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_52
			var _nw310 _dafny.Array
			_ = _nw310
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw310 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Int = func(_1984_i1 _dafny.Int) _dafny.Int {
					return (_1984_i1).Minus(_dafny.IntOfInt64(310))
				}
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw310 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw310).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw310).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1983_v12 = _nw310
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1983_v12), 0))
			_ = _index341
			(_1983_v12).ArraySet1(_dafny.IntOfInt64(632), (_index341).Int())
			var _1985_v13 D10
			_ = _1985_v13
			_1985_v13 = Companion_D10_.Create_DC24_(p0, _1967_v0, _1970_v3)
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1983_v12), 0))
			_ = _index342
			(_1983_v12).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm26(_1973_v6, (_1985_v13).Dtor_cf31(), _1975_v8, globalState), p2), (_index342).Int())
			var _1986_v15 _dafny.Array
			_ = _1986_v15
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_53
			var _nw311 _dafny.Array
			_ = _nw311
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw311 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Map = (func(_1987_v1 _dafny.CodePoint, _1988_p1 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_1989_i2 _dafny.Int) _dafny.Map {
						return (func() _dafny.Map {
							var _coll97 = _dafny.NewMapBuilder()
							_ = _coll97
							for _iter101 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-720), _dafny.IntOfInt64(104))); ; {
								_compr_97, _ok101 := _iter101()
								if !_ok101 {
									break
								}
								var _1990_v14 _dafny.Int
								_1990_v14 = interface{}(_compr_97).(_dafny.Int)
								if ((_dafny.IntOfInt64(-720)).Cmp(_1990_v14) <= 0) && ((_1990_v14).Cmp(_dafny.IntOfInt64(104)) < 0) {
									_coll97.Add(Companion_Default___.SafeDivisionInt(_1990_v14, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(858))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg84 _dafny.Int) interface{} {
											return coer84(arg84)
										}
									}((func(_1991_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
										return func(_1992_i3 _dafny.Int) _dafny.CodePoint {
											return _1991_v1
										}
									})(_1987_v1)))).Cardinality())), (_1988_p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1988_p1), 0))).Int()).(bool))
								}
							}
							return _coll97.ToMap()
						}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(27), (_1988_p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1988_p1), 0))).Int()).(bool)))
					}
				})(_1968_v1, p1)
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw311 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw311).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw311).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_1986_v15 = _nw311
			var _1993_v16 _dafny.Map
			_ = _1993_v16
			_1993_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1983_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1983_v12), 0))).Int()).(_dafny.Int), _1967_v0)
			var _1994_v17 _dafny.Map
			_ = _1994_v17
			_1994_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-945), (func() bool {
				if (_1993_v16).Contains(_dafny.IntOfUint32((_1974_v7).Cardinality())) {
					return (_1993_v16).Get(_dafny.IntOfUint32((_1974_v7).Cardinality())).(bool)
				}
				return _1967_v0
			})())
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1986_v15), 0))
			_ = _index343
			(_1986_v15).ArraySet1(_1994_v17, (_index343).Int())
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1986_v15), 0))
			_ = _index344
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
			_ = _index345
			var _rhs339 _dafny.Map = _1993_v16
			_ = _rhs339
			var _rhs340 bool = _1967_v0
			_ = _rhs340
			var _rhs341 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}((func(_1995_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1996_i4 _dafny.Int) _dafny.CodePoint {
					return _1995_v1
				}
			})(_1968_v1))), (func() _dafny.Sequence {
				if _1967_v0 {
					return _dafny.UnicodeSeqOfUtf8Bytes("fpl")
				}
				return _1970_v3
			})())
			_ = _rhs341
			var _lhs264 _dafny.Array = _1986_v15
			_ = _lhs264
			var _lhs265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1986_v15), 0))
			_ = _lhs265
			var _lhs266 _dafny.Array = p1
			_ = _lhs266
			var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
			_ = _lhs267
			(_lhs264).ArraySet1(_rhs339, (_lhs265).Int())
			(_lhs266).ArraySet1(_rhs340, (_lhs267).Int())
			_1970_v3 = _rhs341
			_1967_v0 = _1967_v0
			var _1997_v18 *C3
			_ = _1997_v18
			var _nw312 *C3 = New_C3_()
			_ = _nw312
			_nw312.Ctor__()
			_1997_v18 = _nw312
		}
		var _1998_v19 _dafny.Set
		_ = _1998_v19
		_1998_v19 = _dafny.SetOf(_1970_v3)
		if (_1998_v19).IsDisjointFrom(_1998_v19) {
			var _1999_v20 _dafny.Sequence
			_ = _1999_v20
			_1999_v20 = _dafny.SeqOf(p0)
			var _2000_v21 _dafny.MultiSet
			_ = _2000_v21
			_2000_v21 = _dafny.MultiSetOf(p2)
			var _2001_v22 _dafny.Set
			_ = _2001_v22
			_2001_v22 = _dafny.SetOf(p2)
			var _2002_v23 *C0
			_ = _2002_v23
			var _nw313 *C0 = New_C0_()
			_ = _nw313
			_nw313.Ctor__()
			_2002_v23 = _nw313
			var _2003_v24 D13
			_ = _2003_v24
			_2003_v24 = Companion_D13_.Create_DC30_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_1980_v10).Fm3(_1999_v20, _2000_v21, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState), (_2001_v22).Difference(_2001_v22), _2002_v23)
			_2003_v24 = Companion_D13_.Create_DC30_((_2001_v22).IsSubsetOf(_2001_v22), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), _2001_v22, _2002_v23)
			var _2004_v25 _dafny.MultiSet
			_ = _2004_v25
			_2004_v25 = _dafny.MultiSetOf(_1970_v3, _1970_v3)
			var _2005_v26 D10
			_ = _2005_v26
			_2005_v26 = Companion_D10_.Create_DC23_(_2004_v25)
			var _source35 D10 = _2005_v26
			_ = _source35
			if _source35.Is_DC24() {
				var _2006___mcc_h0 _dafny.Int = _source35.Get_().(D10_DC24).Cf31
				_ = _2006___mcc_h0
				var _2007___mcc_h1 bool = _source35.Get_().(D10_DC24).Cf32
				_ = _2007___mcc_h1
				var _2008___mcc_h2 _dafny.Sequence = _source35.Get_().(D10_DC24).Cf33
				_ = _2008___mcc_h2
				var _2009_cf33 _dafny.Sequence = _2008___mcc_h2
				_ = _2009_cf33
				var _2010_cf32 bool = _2007___mcc_h1
				_ = _2010_cf32
				var _2011_cf31 _dafny.Int = _2006___mcc_h0
				_ = _2011_cf31
				var _2012_v27 _dafny.Array
				_ = _2012_v27
				var _len0_54 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_54
				var _nw314 _dafny.Array
				_ = _nw314
				if _len0_54.Cmp(_dafny.Zero) == 0 {
					_nw314 = _dafny.NewArray(_len0_54)
				} else {
					var _init54 func(_dafny.Int) _dafny.Int = (func(_2013_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2014_i5 _dafny.Int) _dafny.Int {
							return (_2014_i5).Times(_2013_p2)
						}
					})(p2)
					_ = _init54
					var _element0_54 = _init54(_dafny.Zero)
					_ = _element0_54
					_nw314 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
					(_nw314).ArraySet1(_element0_54, 0)
					var _nativeLen0_54 = (_len0_54).Int()
					_ = _nativeLen0_54
					for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
						(_nw314).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
					}
				}
				_2012_v27 = _nw314
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_2012_v27), 0))
				_ = _index346
				(_2012_v27).ArraySet1(_dafny.IntOfInt64(-106), (_index346).Int())
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_2012_v27), 0))
				_ = _index347
				(_2012_v27).ArraySet1(_2011_cf31, (_index347).Int())
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
				_ = _index348
				(p1).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_index348).Int())
				(globalState).F8 = p2
				var _2015_v28 _dafny.Array
				_ = _2015_v28
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_55
				var _nw315 _dafny.Array
				_ = _nw315
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw315 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) _dafny.CodePoint = (func(_2016_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2017_i6 _dafny.Int) _dafny.CodePoint {
							return _2016_v1
						}
					})(_1968_v1)
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw315 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw315).ArraySet1CodePoint(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw315).ArraySet1CodePoint(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_2015_v28 = _nw315
				var _2018_v29 _dafny.Map
				_ = _2018_v29
				_2018_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2015_v28, _dafny.IntOfInt64(459))
				_2018_v29 = _2018_v29
			} else {
				var _2019___mcc_h3 _dafny.MultiSet = _source35.Get_().(D10_DC23).Cf30
				_ = _2019___mcc_h3
				var _2020_cf30 _dafny.MultiSet = _2019___mcc_h3
				_ = _2020_cf30
				var _2021_v30 *C0
				_ = _2021_v30
				var _nw316 *C0 = New_C0_()
				_ = _nw316
				_nw316.Ctor__()
				_2021_v30 = _nw316
				_1967_v0 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				(globalState).F8 = p0
				var _2022_v31 D2
				_ = _2022_v31
				_2022_v31 = Companion_D2_.Create_DC5_(p1)
				_2022_v31 = _2022_v31
			}
			var _2023_v32 *C1
			_ = _2023_v32
			var _nw317 *C1 = New_C1_()
			_ = _nw317
			_nw317.Ctor__()
			_2023_v32 = _nw317
			var _2024_v33 D7
			_ = _2024_v33
			_2024_v33 = Companion_D7_.Create_DC18_()
			var _2025_v34 _dafny.Map
			_ = _2025_v34
			_2025_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2003_v24).Dtor_cf43(), _2024_v33)
			var _2026_v35 _dafny.Map
			_ = _2026_v35
			_2026_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2001_v22).Cardinality(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			if !((_2025_v34).Equals(Companion_Default___.Fm48(_1999_v20, _2026_v35, globalState))) {
				(globalState).F8 = p2
				var _2027_v36 _dafny.Map
				_ = _2027_v36
				_2027_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0, p0, p2, p2, _dafny.IntOfUint32((_1974_v7).Cardinality())), (_2000_v21).Intersection(_2000_v21))
				var _2028_v37 _dafny.Sequence
				_ = _2028_v37
				_2028_v37 = _dafny.SeqOf(_2000_v21)
				_2027_v36 = (_2027_v36).Update(_1999_v20, (_2028_v37).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2028_v37).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
				_ = _index349
				(p1).ArraySet1(((_2023_v32).Fm3(_1999_v20, _2000_v21, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState)), (_index349).Int())
				var _2029_v38 _dafny.Array
				_ = _2029_v38
				var _nw318 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw318
				_2029_v38 = _nw318
				_2029_v38 = _2029_v38
				var _2030_v39 _dafny.Array
				_ = _2030_v39
				var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw319
				_2030_v39 = _nw319
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2030_v39), 0))
				_ = _index350
				(_2030_v39).ArraySet1(_1974_v7, (_index350).Int())
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2030_v39), 0))
				_ = _index351
				(_2030_v39).ArraySet1(_1974_v7, (_index351).Int())
			} else {
				(globalState).F4 = true
				var _2031_v40 D19
				_ = _2031_v40
				_2031_v40 = Companion_D19_.Create_DC39_(false, _dafny.IntOfUint32((_1970_v3).Cardinality()))
				var _2032_v41 _dafny.Map
				_ = _2032_v41
				_2032_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1968_v1, p2)
				(globalState).F8 = ((_dafny.Zero).Minus((_2031_v40).Dtor_cf59())).Times((func() _dafny.Int {
					if (_2032_v41).Contains(_1968_v1) {
						return (_2032_v41).Get(_1968_v1).(_dafny.Int)
					}
					return p2
				})())
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
				_ = _index352
				(p1).ArraySet1(true, (_index352).Int())
				(globalState).F3 = (_dafny.IntOfInt64(39)).Times(Companion_Default___.Fm20((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState))
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
				_ = _index353
				(p1).ArraySet1(!(_2026_v35).Contains(p2), (_index353).Int())
			}
			var _2033_v42 _dafny.Array
			_ = _2033_v42
			var _len0_56 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_56
			var _nw320 _dafny.Array
			_ = _nw320
			if _len0_56.Cmp(_dafny.Zero) == 0 {
				_nw320 = _dafny.NewArray(_len0_56)
			} else {
				var _init56 func(_dafny.Int) _dafny.Map = (func(_2034_v0 bool, _2035_p2 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_2036_i7 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2034_v0, _2035_p2)
					}
				})(_1967_v0, p2)
				_ = _init56
				var _element0_56 = _init56(_dafny.Zero)
				_ = _element0_56
				_nw320 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
				(_nw320).ArraySet1(_element0_56, 0)
				var _nativeLen0_56 = (_len0_56).Int()
				_ = _nativeLen0_56
				for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
					(_nw320).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
				}
			}
			_2033_v42 = _nw320
			var _2037_v43 _dafny.Map
			_ = _2037_v43
			_2037_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_1970_v3).Cardinality()))
			var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_2033_v42), 0))
			_ = _index354
			(_2033_v42).ArraySet1((_2037_v43).Merge(_2037_v43), (_index354).Int())
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_2033_v42), 0))
			_ = _index355
			(_2033_v42).ArraySet1(_2037_v43, (_index355).Int())
		} else {
			var _2038_v44 _dafny.Set
			_ = _2038_v44
			_2038_v44 = _dafny.SetOf(_1967_v0, !(true), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			(globalState).F3 = ((_2038_v44).Union(_2038_v44)).Cardinality()
			var _2039_v45 _dafny.Array
			_ = _2039_v45
			var _nw321 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw321
			_2039_v45 = _nw321
			var _2040_v46 _dafny.MultiSet
			_ = _2040_v46
			_2040_v46 = _dafny.MultiSetOf(_dafny.IntOfInt64(244))
			var _2041_v47 _dafny.Array
			_ = _2041_v47
			var _nwElement0_61 _dafny.Array = _2039_v45
			_ = _nwElement0_61
			var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(21))
			_ = _nw322
			(_nw322).ArraySet1(_nwElement0_61, 0)
			(_nw322).ArraySet1(_2039_v45, 1)
			(_nw322).ArraySet1(_2039_v45, 2)
			(_nw322).ArraySet1((func() _dafny.Array {
				if (_this).Fm3(_dafny.SeqOf(p2, p0, p0), _2040_v46, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState) {
					return _2039_v45
				}
				return _2039_v45
			})(), 3)
			(_nw322).ArraySet1(_2039_v45, 4)
			(_nw322).ArraySet1((func() _dafny.Array {
				if !(_1967_v0) {
					return _2039_v45
				}
				return _2039_v45
			})(), 5)
			(_nw322).ArraySet1(_2039_v45, 6)
			(_nw322).ArraySet1(_2039_v45, 7)
			(_nw322).ArraySet1(_2039_v45, 8)
			(_nw322).ArraySet1(_2039_v45, 9)
			(_nw322).ArraySet1(_2039_v45, 10)
			(_nw322).ArraySet1(_2039_v45, 11)
			(_nw322).ArraySet1(_2039_v45, 12)
			(_nw322).ArraySet1(_2039_v45, 13)
			(_nw322).ArraySet1(_2039_v45, 14)
			(_nw322).ArraySet1(_2039_v45, 15)
			(_nw322).ArraySet1(_2039_v45, 16)
			(_nw322).ArraySet1(_2039_v45, 17)
			(_nw322).ArraySet1(_2039_v45, 18)
			(_nw322).ArraySet1(_2039_v45, 19)
			(_nw322).ArraySet1(_2039_v45, 20)
			_2041_v47 = _nw322
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_2041_v47), 0))
			_ = _index356
			(_2041_v47).ArraySet1(_2039_v45, (_index356).Int())
			var _2042_v48 _dafny.Array
			_ = _2042_v48
			var _nwElement0_62 _dafny.CodePoint = _1968_v1
			_ = _nwElement0_62
			var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(11))
			_ = _nw323
			(_nw323).ArraySet1CodePoint(_nwElement0_62, 0)
			(_nw323).ArraySet1CodePoint(_1968_v1, 1)
			(_nw323).ArraySet1CodePoint(_1968_v1, 2)
			(_nw323).ArraySet1CodePoint(_1968_v1, 3)
			(_nw323).ArraySet1CodePoint(_1968_v1, 4)
			(_nw323).ArraySet1CodePoint(_1968_v1, 5)
			(_nw323).ArraySet1CodePoint(_1968_v1, 6)
			(_nw323).ArraySet1CodePoint(_1968_v1, 7)
			(_nw323).ArraySet1CodePoint(_1968_v1, 8)
			(_nw323).ArraySet1CodePoint(_1968_v1, 9)
			(_nw323).ArraySet1CodePoint(_1968_v1, 10)
			_2042_v48 = _nw323
			var _2043_v49 _dafny.Set
			_ = _2043_v49
			_2043_v49 = _dafny.SetOf(p1)
			var _2044_v50 _dafny.Sequence
			_ = _2044_v50
			_2044_v50 = _dafny.SeqOf(_dafny.SetOf(p1, p1), _2043_v49)
			var _2045_v51 _dafny.Sequence
			_ = _2045_v51
			_2045_v51 = _dafny.SeqOf(((_2044_v50).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_1974_v7)).Cardinality(), _dafny.IntOfUint32((_2044_v50).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
			var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_2041_v47), 0))
			_ = _index357
			var _rhs342 _dafny.Array = _2039_v45
			_ = _rhs342
			var _rhs343 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1980_v10).Fm3(_2045_v51, _dafny.MultiSetOf(_dafny.IntOfInt64(436)), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(772))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg86 _dafny.Int) interface{} {
							return coer86(arg86)
						}
					}((func(_2046_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2047_i8 _dafny.Int) _dafny.CodePoint {
							return _2046_v1
						}
					})(_1968_v1)))
				}
				return _dafny.Companion_Sequence_.Concatenate(_1970_v3, _1970_v3)
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2045_v51).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1980_v10).Fm3(_2045_v51, _dafny.MultiSetOf(_dafny.IntOfInt64(436)), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(772))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg87 _dafny.Int) interface{} {
							return coer87(arg87)
						}
					}((func(_2048_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2049_i8 _dafny.Int) _dafny.CodePoint {
							return _2048_v1
						}
					})(_1968_v1)))
				}
				return _dafny.Companion_Sequence_.Concatenate(_1970_v3, _1970_v3)
			})()).Cardinality()))).Uint32(), _1968_v1)
			_ = _rhs343
			var _rhs344 _dafny.Array = (Companion_D23_.Create_DC45_(_2042_v48)).Dtor_cf66()
			_ = _rhs344
			var _lhs268 _dafny.Array = _2041_v47
			_ = _lhs268
			var _lhs269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_2041_v47), 0))
			_ = _lhs269
			(_lhs268).ArraySet1(_rhs342, (_lhs269).Int())
			_1970_v3 = _rhs343
			_2042_v48 = _rhs344
			var _2050_v52 _dafny.Array
			_ = _2050_v52
			var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw324
			_2050_v52 = _nw324
			var _2051_v53 _dafny.Array
			_ = _2051_v53
			_2051_v53 = _2050_v52
			var _source36 _dafny.Array = _2051_v53
			_ = _source36
			var _2052___mcc_h4 _dafny.Array = _source36
			_ = _2052___mcc_h4
			var _2053_cf55 _dafny.Array = _2052___mcc_h4
			_ = _2053_cf55
			(globalState).F4 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			var _2054_v54 _dafny.Sequence
			_ = _2054_v54
			_2054_v54 = _dafny.SeqOf(p1)
			var _2055_v55 _dafny.Array
			_ = _2055_v55
			var _nwElement0_63 _dafny.Array = p1
			_ = _nwElement0_63
			var _nw325 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(16))
			_ = _nw325
			(_nw325).ArraySet1(_nwElement0_63, 0)
			(_nw325).ArraySet1(p1, 1)
			(_nw325).ArraySet1(p1, 2)
			(_nw325).ArraySet1(p1, 3)
			(_nw325).ArraySet1(p1, 4)
			(_nw325).ArraySet1(p1, 5)
			(_nw325).ArraySet1(p1, 6)
			(_nw325).ArraySet1(p1, 7)
			(_nw325).ArraySet1((_2054_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.IntOfUint32((_2054_v54).Cardinality()))).Uint32()).(_dafny.Array), 8)
			(_nw325).ArraySet1(p1, 9)
			(_nw325).ArraySet1(p1, 10)
			(_nw325).ArraySet1(p1, 11)
			(_nw325).ArraySet1(p1, 12)
			(_nw325).ArraySet1(p1, 13)
			(_nw325).ArraySet1(p1, 14)
			(_nw325).ArraySet1(p1, 15)
			_2055_v55 = _nw325
			var _2056_v56 _dafny.Array
			_ = _2056_v56
			var _nw326 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
			_ = _nw326
			_2056_v56 = _nw326
			var _2057_v57 _dafny.Set
			_ = _2057_v57
			_2057_v57 = _dafny.SetOf(_2056_v56, _2056_v56)
			var _rhs345 _dafny.Sequence = _1970_v3
			_ = _rhs345
			var _rhs346 _dafny.Array = _2055_v55
			_ = _rhs346
			var _rhs347 _dafny.Int = ((_2057_v57).Cardinality()).Minus((func() _dafny.Int {
				if (_1972_v5).Contains(p2) {
					return (_1972_v5).Get(p2).(_dafny.Int)
				}
				return p2
			})())
			_ = _rhs347
			var _lhs270 *GlobalState = globalState
			_ = _lhs270
			_1970_v3 = _rhs345
			_2055_v55 = _rhs346
			_lhs270.F8 = _rhs347
			(globalState).F8 = p0
			var _2058_v58 _dafny.Map
			_ = _2058_v58
			_2058_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sl"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(711))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}((func(_2059_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2060_i9 _dafny.Int) _dafny.CodePoint {
					return _2059_v1
				}
			})(_1968_v1))))
			_1967_v0 = ((_1974_v7).Select((Companion_Default___.SafeIndex((_2058_v58).Cardinality(), _dafny.IntOfUint32((_1974_v7).Cardinality()))).Uint32()).(bool)) || (_1967_v0)
			(globalState).F4 = (_2040_v46).IsDisjointFrom(_2040_v46)
			var _2061_v59 _dafny.Map
			_ = _2061_v59
			_2061_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2038_v44)
			var _2062_v60 _dafny.Sequence
			_ = _2062_v60
			_2062_v60 = _dafny.SeqOf(p1)
			var _2063_v61 _dafny.Array
			_ = _2063_v61
			var _nwElement0_64 _dafny.Array = p1
			_ = _nwElement0_64
			var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(22))
			_ = _nw327
			(_nw327).ArraySet1(_nwElement0_64, 0)
			(_nw327).ArraySet1(p1, 1)
			(_nw327).ArraySet1(p1, 2)
			(_nw327).ArraySet1(p1, 3)
			(_nw327).ArraySet1(p1, 4)
			(_nw327).ArraySet1(p1, 5)
			(_nw327).ArraySet1((_2062_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.IntOfUint32((_2062_v60).Cardinality()))).Uint32()).(_dafny.Array), 6)
			(_nw327).ArraySet1(p1, 7)
			(_nw327).ArraySet1((_2062_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2045_v51).Cardinality()), _dafny.IntOfUint32((_2062_v60).Cardinality()))).Uint32()).(_dafny.Array), 8)
			(_nw327).ArraySet1(p1, 9)
			(_nw327).ArraySet1(p1, 10)
			(_nw327).ArraySet1(p1, 11)
			(_nw327).ArraySet1(p1, 12)
			(_nw327).ArraySet1(p1, 13)
			(_nw327).ArraySet1(p1, 14)
			(_nw327).ArraySet1(p1, 15)
			(_nw327).ArraySet1(p1, 16)
			(_nw327).ArraySet1(p1, 17)
			(_nw327).ArraySet1(p1, 18)
			(_nw327).ArraySet1(p1, 19)
			(_nw327).ArraySet1(p1, 20)
			(_nw327).ArraySet1(p1, 21)
			_2063_v61 = _nw327
			var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_2063_v61), 0))
			_ = _index358
			(_2063_v61).ArraySet1(p1, (_index358).Int())
			var _2064_v63 D12
			_ = _2064_v63
			_2064_v63 = Companion_D12_.Create_DC27_((func() _dafny.Map {
				var _coll98 = _dafny.NewMapBuilder()
				_ = _coll98
				for _iter102 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-216), _dafny.IntOfInt64(152))); ; {
					_compr_98, _ok102 := _iter102()
					if !_ok102 {
						break
					}
					var _2065_v62 _dafny.Int
					_2065_v62 = interface{}(_compr_98).(_dafny.Int)
					if ((_dafny.IntOfInt64(-216)).Cmp(_2065_v62) <= 0) && ((_2065_v62).Cmp(_dafny.IntOfInt64(152)) < 0) {
						_coll98.Add((_2065_v62).Minus(p2), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
					}
				}
				return _coll98.ToMap()
			}()).Cardinality(), p2, _dafny.IntOfUint32((_1970_v3).Cardinality()), (_dafny.Zero).Minus(p0))
			var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_2063_v61), 0))
			_ = _index359
			var _rhs348 _dafny.Int = (_2064_v63).Dtor_cf39()
			_ = _rhs348
			var _rhs349 _dafny.Map = ((_2061_v59).Update(p0, _2038_v44)).Merge(_2061_v59)
			_ = _rhs349
			var _rhs350 bool = false
			_ = _rhs350
			var _rhs351 _dafny.CodePoint = _1968_v1
			_ = _rhs351
			var _rhs352 _dafny.Array = p1
			_ = _rhs352
			var _lhs271 *GlobalState = globalState
			_ = _lhs271
			var _lhs272 *GlobalState = globalState
			_ = _lhs272
			var _lhs273 _dafny.Array = _2063_v61
			_ = _lhs273
			var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_2063_v61), 0))
			_ = _lhs274
			_lhs271.F3 = _rhs348
			_2061_v59 = _rhs349
			_lhs272.F4 = _rhs350
			_1968_v1 = _rhs351
			(_lhs273).ArraySet1(_rhs352, (_lhs274).Int())
		}
		var _2066_v65 _dafny.Map
		_ = _2066_v65
		_2066_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1968_v1, func() _dafny.Map {
			var _coll99 = _dafny.NewMapBuilder()
			_ = _coll99
			for _iter103 := _dafny.Iterate((_1970_v3).Elements()); ; {
				_compr_99, _ok103 := _iter103()
				if !_ok103 {
					break
				}
				var _2067_v64 _dafny.CodePoint
				_2067_v64 = interface{}(_compr_99).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_1970_v3, _2067_v64) {
					_coll99.Add(_2067_v64, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
				}
			}
			return _coll99.ToMap()
		}())
		var _2068_v66 _dafny.Sequence
		_ = _2068_v66
		_2068_v66 = _dafny.SeqOf(p2, p2, p2, (_2066_v65).Cardinality())
		if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(724))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}((func(_2069_v66 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_2070_i10 _dafny.Int) _dafny.Sequence {
				return _2069_v66
			}
		})(_2068_v66))), _2068_v66) {
			(globalState).F8 = (_dafny.Zero).Minus(p2)
			var _2071_v67 _dafny.Map
			_ = _2071_v67
			_2071_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), p0)
			var _2072_v68 _dafny.Map
			_ = _2072_v68
			_2072_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2068_v66, _1980_v10)
			(globalState).F8 = (Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_2071_v67).Contains(_1967_v0) {
					return (_2071_v67).Get(_1967_v0).(_dafny.Int)
				}
				return Companion_Default___.Fm26(_1973_v6, p2, _1975_v8, globalState)
			})(), _dafny.IntOfUint32((_1974_v7).Cardinality()))).Plus(Companion_Default___.SafeDivisionInt((_this).Fm10(globalState), (_2072_v68).Cardinality()))
			if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
				var _2073_v69 D12
				_ = _2073_v69
				_2073_v69 = Companion_D12_.Create_DC27_(_dafny.IntOfInt64(852), p0, _dafny.IntOfInt64(886), p0)
				var _pat_let_tv70 = p0
				_ = _pat_let_tv70
				var _2074_v70 _dafny.Array
				_ = _2074_v70
				var _nwElement0_65 D12 = _2073_v69
				_ = _nwElement0_65
				var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(18))
				_ = _nw328
				(_nw328).ArraySet1(_nwElement0_65, 0)
				(_nw328).ArraySet1(_2073_v69, 1)
				(_nw328).ArraySet1(_2073_v69, 2)
				(_nw328).ArraySet1(_2073_v69, 3)
				(_nw328).ArraySet1(_2073_v69, 4)
				(_nw328).ArraySet1(Companion_D12_.Create_DC27_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(201), p2)).Cardinality(), _dafny.IntOfUint32((_1970_v3).Cardinality()), Companion_Default___.Fm20(true, globalState), p0), 5)
				(_nw328).ArraySet1(Companion_D12_.Create_DC27_((_1972_v5).Cardinality(), _dafny.IntOfUint32((_1970_v3).Cardinality()), p2, p2), 6)
				(_nw328).ArraySet1(_2073_v69, 7)
				(_nw328).ArraySet1(_2073_v69, 8)
				(_nw328).ArraySet1(_2073_v69, 9)
				(_nw328).ArraySet1(_2073_v69, 10)
				(_nw328).ArraySet1(_2073_v69, 11)
				(_nw328).ArraySet1(_2073_v69, 12)
				(_nw328).ArraySet1(Companion_D12_.Create_DC27_(p2, (_2071_v67).Cardinality(), p2, p0), 13)
				(_nw328).ArraySet1(_2073_v69, 14)
				(_nw328).ArraySet1(_2073_v69, 15)
				(_nw328).ArraySet1(_2073_v69, 16)
				(_nw328).ArraySet1(func(_pat_let60_0 D12) D12 {
					return func(_2075_dt__update__tmp_h2 D12) D12 {
						return func(_pat_let61_0 _dafny.Int) D12 {
							return func(_2076_dt__update_hcf36_h0 _dafny.Int) D12 {
								return func(_pat_let62_0 _dafny.Int) D12 {
									return func(_2077_dt__update_hcf39_h0 _dafny.Int) D12 {
										return Companion_D12_.Create_DC27_(_2076_dt__update_hcf36_h0, (_2075_dt__update__tmp_h2).Dtor_cf37(), (_2075_dt__update__tmp_h2).Dtor_cf38(), _2077_dt__update_hcf39_h0)
									}(_pat_let62_0)
								}(_pat_let_tv70)
							}(_pat_let61_0)
						}(_dafny.IntOfInt64(183))
					}(_pat_let60_0)
				}(_2073_v69), 17)
				_2074_v70 = _nw328
				var _2078_v71 _dafny.Array
				_ = _2078_v71
				var _nwElement0_66 _dafny.Array = (func() _dafny.Array {
					if _1967_v0 {
						return _2074_v70
					}
					return _2074_v70
				})()
				_ = _nwElement0_66
				var _nw329 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(14))
				_ = _nw329
				(_nw329).ArraySet1(_nwElement0_66, 0)
				(_nw329).ArraySet1(_2074_v70, 1)
				(_nw329).ArraySet1(_2074_v70, 2)
				(_nw329).ArraySet1(_2074_v70, 3)
				(_nw329).ArraySet1(_2074_v70, 4)
				(_nw329).ArraySet1(_2074_v70, 5)
				(_nw329).ArraySet1(_2074_v70, 6)
				(_nw329).ArraySet1(_2074_v70, 7)
				(_nw329).ArraySet1(_2074_v70, 8)
				(_nw329).ArraySet1(_2074_v70, 9)
				(_nw329).ArraySet1(_2074_v70, 10)
				(_nw329).ArraySet1(_2074_v70, 11)
				(_nw329).ArraySet1((func() _dafny.Array {
					if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
						return _2074_v70
					}
					return _2074_v70
				})(), 12)
				(_nw329).ArraySet1(_2074_v70, 13)
				_2078_v71 = _nw329
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2078_v71), 0))
				_ = _index360
				(_2078_v71).ArraySet1(_2074_v70, (_index360).Int())
				var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2078_v71), 0))
				_ = _index361
				(_2078_v71).ArraySet1(_2074_v70, (_index361).Int())
				_2071_v67 = (_2071_v67).Update(_1967_v0, p0)
				_1967_v0 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				(globalState).F8 = ((_dafny.SetOf(_1968_v1)).Union(func() _dafny.Set {
					var _coll100 = _dafny.NewBuilder()
					_ = _coll100
					for _iter104 := _dafny.Iterate((_1970_v3).Elements()); ; {
						_compr_100, _ok104 := _iter104()
						if !_ok104 {
							break
						}
						var _2079_v72 _dafny.CodePoint
						_2079_v72 = interface{}(_compr_100).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_1970_v3, _2079_v72) {
							_coll100.Add(_2079_v72)
						}
					}
					return _coll100.ToSet()
				}())).Cardinality()
				var _2080_v73 _dafny.Map
				_ = _2080_v73
				_2080_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Times(_dafny.IntOfUint32((_2068_v66).Cardinality())), Companion_Default___.Fm46(globalState))
				var _2081_v74 _dafny.MultiSet
				_ = _2081_v74
				_2081_v74 = _dafny.MultiSetOf(_1967_v0)
				_2080_v73 = (_2080_v73).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(600), (_2081_v74).Cardinality()), (_2071_v67).Update(_1967_v0, p0))
			} else {
				var _2082_v75 _dafny.MultiSet
				_ = _2082_v75
				_2082_v75 = _dafny.MultiSetOf(p2)
				var _2083_v76 D4
				_ = _2083_v76
				_2083_v76 = Companion_D4_.Create_DC10_(_dafny.SeqOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_this).Fm3(_2068_v66, _2082_v75, false, globalState), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)))
				var _2084_v77 _dafny.Map
				_ = _2084_v77
				_2084_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1970_v3, _1970_v3), p0)
				var _rhs353 _dafny.Sequence = _1970_v3
				_ = _rhs353
				var _rhs354 D4 = Companion_D4_.Create_DC10_(_dafny.Companion_Sequence_.Update(_1974_v7, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1974_v7).Cardinality()))).Uint32(), !(_1967_v0)))
				_ = _rhs354
				var _rhs355 _dafny.Int = (func() _dafny.Int {
					if (_2084_v77).Contains(_1970_v3) {
						return (_2084_v77).Get(_1970_v3).(_dafny.Int)
					}
					return p2
				})()
				_ = _rhs355
				var _rhs356 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _1967_v0 {
						return _dafny.Companion_Sequence_.Concatenate(_1970_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg90 _dafny.Int) interface{} {
								return coer90(arg90)
							}
						}((func(_2085_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2086_i11 _dafny.Int) _dafny.CodePoint {
								return _2085_v1
							}
						})(_1968_v1))))
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg91 _dafny.Int) interface{} {
							return coer91(arg91)
						}
					}((func(_2087_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2088_i12 _dafny.Int) _dafny.CodePoint {
							return _2087_v1
						}
					})(_1968_v1)))
				})()).Cardinality())
				_ = _rhs356
				var _lhs275 *GlobalState = globalState
				_ = _lhs275
				var _lhs276 *GlobalState = globalState
				_ = _lhs276
				_1970_v3 = _rhs353
				_2083_v76 = _rhs354
				_lhs275.F8 = _rhs355
				_lhs276.F8 = _rhs356
				_1976_v9 = Companion_Default___.Fm40(_1967_v0, globalState)
				var _2089_v78 _dafny.Map
				_ = _2089_v78
				_2089_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).Fm10(globalState))
				_2089_v78 = _2089_v78
				var _2090_v79 _dafny.Array
				_ = _2090_v79
				var _nw330 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
				_ = _nw330
				_2090_v79 = _nw330
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_2090_v79), 0))
				_ = _index362
				(_2090_v79).ArraySet1CodePoint(_1968_v1, (_index362).Int())
				var _2091_v80 _dafny.Array
				_ = _2091_v80
				var _nw331 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw331
				_2091_v80 = _nw331
				var _2092_v81 _dafny.Set
				_ = _2092_v81
				_2092_v81 = _dafny.SetOf(p0, _dafny.IntOfUint32((_1970_v3).Cardinality()), p0)
				var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_2090_v79), 0))
				_ = _index363
				var _rhs357 _dafny.CodePoint = _dafny.CodePoint('t')
				_ = _rhs357
				var _rhs358 bool = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				_ = _rhs358
				var _rhs359 _dafny.Int = p0
				_ = _rhs359
				var _rhs360 _dafny.CodePoint = Companion_Default___.Fm18((_dafny.IntOfUint32((_1970_v3).Cardinality())).Cmp((_2092_v81).Cardinality()) <= 0, _1967_v0, globalState)
				_ = _rhs360
				var _rhs361 _dafny.Array = _2091_v80
				_ = _rhs361
				var _lhs277 *GlobalState = globalState
				_ = _lhs277
				var _lhs278 _dafny.Array = _2090_v79
				_ = _lhs278
				var _lhs279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_2090_v79), 0))
				_ = _lhs279
				_1968_v1 = _rhs357
				_1967_v0 = _rhs358
				_lhs277.F8 = _rhs359
				(_lhs278).ArraySet1CodePoint(_rhs360, (_lhs279).Int())
				_2091_v80 = _rhs361
				var _2093_v82 *C1
				_ = _2093_v82
				var _nw332 *C1 = New_C1_()
				_ = _nw332
				_nw332.Ctor__()
				_2093_v82 = _nw332
				_2093_v82 = _2093_v82
			}
			(globalState).F4 = !(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dgljestqn"), _1970_v3), _dafny.Companion_Sequence_.Concatenate(_1970_v3, _1970_v3)))
			var _2094_v83 _dafny.Set
			_ = _2094_v83
			_2094_v83 = _dafny.SetOf(!((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)), _1967_v0)
			var _2095_v84 _dafny.Map
			_ = _2095_v84
			_2095_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1967_v0)
			var _2096_v85 _dafny.Map
			_ = _2096_v85
			_2096_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			var _2097_v86 _dafny.Set
			_ = _2097_v86
			_2097_v86 = _dafny.SetOf(p2)
			var _2098_v87 *C0
			_ = _2098_v87
			var _nw333 *C0 = New_C0_()
			_ = _nw333
			_nw333.Ctor__()
			_2098_v87 = _nw333
			var _2099_v88 D13
			_ = _2099_v88
			_2099_v88 = Companion_D13_.Create_DC30_(_1967_v0, _1967_v0, _2097_v86, _2098_v87)
			var _2100_v89 D13
			_ = _2100_v89
			_2100_v89 = Companion_D13_.Create_DC31_(_2099_v88)
			var _2101_v90 _dafny.Set
			_ = _2101_v90
			_2101_v90 = _dafny.SetOf(_2100_v89)
			var _2102_v92 _dafny.Array
			_ = _2102_v92
			var _nwElement0_67 bool = (_1967_v0) && (_1967_v0)
			_ = _nwElement0_67
			var _nw334 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(20))
			_ = _nw334
			(_nw334).ArraySet1(_nwElement0_67, 0)
			(_nw334).ArraySet1((_dafny.IntOfInt64(-957)).Cmp(p2) <= 0, 1)
			(_nw334).ArraySet1(true, 2)
			(_nw334).ArraySet1(((_2094_v83).Cardinality()).Cmp(Companion_Default___.Fm20(_1967_v0, globalState)) <= 0, 3)
			(_nw334).ArraySet1((func() bool {
				if (_2095_v84).Contains(p1) {
					return (_2095_v84).Get(p1).(bool)
				}
				return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			})(), 4)
			(_nw334).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), 5)
			(_nw334).ArraySet1(_1967_v0, 6)
			(_nw334).ArraySet1((_dafny.IntOfInt64(695)).Cmp(p0) <= 0, 7)
			(_nw334).ArraySet1(true, 8)
			(_nw334).ArraySet1((func() bool {
				if (_2096_v85).Contains(_dafny.IntOfInt64(-954)) {
					return (_2096_v85).Get(_dafny.IntOfInt64(-954)).(bool)
				}
				return _1967_v0
			})(), 9)
			(_nw334).ArraySet1(false, 10)
			(_nw334).ArraySet1((p2).Cmp(p0) <= 0, 11)
			(_nw334).ArraySet1((_2101_v90).IsSubsetOf(_2101_v90), 12)
			(_nw334).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), 13)
			(_nw334).ArraySet1(true, 14)
			(_nw334).ArraySet1(_1967_v0, 15)
			(_nw334).ArraySet1(true, 16)
			(_nw334).ArraySet1(((func() _dafny.Set {
				var _coll101 = _dafny.NewBuilder()
				_ = _coll101
				for _iter105 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(321), _dafny.IntOfInt64(692))); ; {
					_compr_101, _ok105 := _iter105()
					if !_ok105 {
						break
					}
					var _2103_v91 _dafny.Int
					_2103_v91 = interface{}(_compr_101).(_dafny.Int)
					if ((_dafny.IntOfInt64(321)).Cmp(_2103_v91) <= 0) && ((_2103_v91).Cmp(_dafny.IntOfInt64(692)) < 0) {
						_coll101.Add((_2103_v91).Plus(p2))
					}
				}
				return _coll101.ToSet()
			}()).Cardinality()).Cmp((_2068_v66).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2068_v66).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, 17)
			(_nw334).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), 18)
			(_nw334).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool), 19)
			_2102_v92 = _nw334
			var _2104_v93 _dafny.Map
			_ = _2104_v93
			_2104_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1970_v3, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			var _2105_v94 _dafny.Map
			_ = _2105_v94
			_2105_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2104_v93, _1967_v0)
			var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
			_ = _index364
			var _rhs362 bool = (func() bool {
				if (_2105_v94).Contains(_2104_v93) {
					return (_2105_v94).Get(_2104_v93).(bool)
				}
				return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			})()
			_ = _rhs362
			var _rhs363 _dafny.Array = _2102_v92
			_ = _rhs363
			var _lhs280 _dafny.Array = p1
			_ = _lhs280
			var _lhs281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))
			_ = _lhs281
			(_lhs280).ArraySet1(_rhs362, (_lhs281).Int())
			_2102_v92 = _rhs363
		} else {
			_1967_v0 = _1967_v0
			var _2106_v95 *C6
			_ = _2106_v95
			var _nw335 *C6 = New_C6_()
			_ = _nw335
			_nw335.Ctor__(_dafny.IntOfInt64(966), _1970_v3)
			_2106_v95 = _nw335
			(globalState).F8 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_2106_v95).F10()))
			var _2107_v96 _dafny.Set
			_ = _2107_v96
			_2107_v96 = _dafny.SetOf(p0)
			(globalState).F3 = (p0).Minus(Companion_Default___.SafeDivisionInt((_2107_v96).Cardinality(), p2))
			var _2108_v97 _dafny.Array
			_ = _2108_v97
			var _nw336 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw336
			_2108_v97 = _nw336
			var _2109_v98 _dafny.MultiSet
			_ = _2109_v98
			_2109_v98 = _dafny.MultiSetOf(_2108_v97)
			var _rhs364 _dafny.MultiSet = _dafny.MultiSetOf(_2108_v97, _2108_v97, _2108_v97)
			_ = _rhs364
			var _rhs365 _dafny.Int = Companion_Default___.Fm20((func() bool {
				if true {
					return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				}
				return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			})(), globalState)
			_ = _rhs365
			var _rhs366 _dafny.Int = _dafny.IntOfInt64(-268)
			_ = _rhs366
			var _rhs367 _dafny.Int = p2
			_ = _rhs367
			var _lhs282 *GlobalState = globalState
			_ = _lhs282
			var _lhs283 *GlobalState = globalState
			_ = _lhs283
			var _lhs284 *GlobalState = globalState
			_ = _lhs284
			_2109_v98 = _rhs364
			_lhs282.F3 = _rhs365
			_lhs283.F3 = _rhs366
			_lhs284.F8 = _rhs367
		}
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
func (_this *C8) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wsbmpw"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-960))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}(func(_2110_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})))
	}
}
func (_this *C8) Fm9(p0 bool, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(579), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tspj")).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_2111_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("coabyxfb")
		}))).Cardinality()))).Cardinality()).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Cardinality())
	}
}
func (_this *C8) M4(p0 bool, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _hi12 _dafny.Int = (p2).Minus(p2)
		_ = _hi12
		for _2112_i0 := (p2).Plus(p2); _2112_i0.Cmp(_hi12) < 0; _2112_i0 = _2112_i0.Plus(_dafny.One) {
			var _2113_v0 D0
			_ = _2113_v0
			_2113_v0 = Companion_D0_.Create_DC2_()
			var _source37 D0 = _2113_v0
			_ = _source37
			if _source37.Is_DC1() {
				var _2114___mcc_h0 _dafny.Int = _source37.Get_().(D0_DC1).Cf1
				_ = _2114___mcc_h0
				var _2115___mcc_h1 _dafny.Int = _source37.Get_().(D0_DC1).Cf2
				_ = _2115___mcc_h1
				var _2116_cf2 _dafny.Int = _2115___mcc_h1
				_ = _2116_cf2
				var _2117_cf1 _dafny.Int = _2114___mcc_h0
				_ = _2117_cf1
				var _2118_v1 _dafny.Array
				_ = _2118_v1
				var _nw337 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(4))
				_ = _nw337
				_2118_v1 = _nw337
				var _2119_v2 D0
				_ = _2119_v2
				_2119_v2 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(179))
				var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2118_v1), 0))
				_ = _index365
				(_2118_v1).ArraySet1(_2119_v2, (_index365).Int())
				var _2120_v3 _dafny.CodePoint
				_ = _2120_v3
				_2120_v3 = _dafny.CodePoint('n')
				var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2118_v1), 0))
				_ = _index366
				var _rhs368 D0 = _2119_v2
				_ = _rhs368
				var _rhs369 _dafny.CodePoint = _2120_v3
				_ = _rhs369
				var _lhs285 _dafny.Array = _2118_v1
				_ = _lhs285
				var _lhs286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2118_v1), 0))
				_ = _lhs286
				(_lhs285).ArraySet1(_rhs368, (_lhs286).Int())
				_2120_v3 = _rhs369
				var _2121_v4 _dafny.Array
				_ = _2121_v4
				var _nw338 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw338
				_2121_v4 = _nw338
				var _2122_v5 _dafny.Sequence
				_ = _2122_v5
				_2122_v5 = _dafny.UnicodeSeqOfUtf8Bytes("vvute")
				var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_2121_v4), 0))
				_ = _index367
				(_2121_v4).ArraySet1(_2122_v5, (_index367).Int())
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_2121_v4), 0))
				_ = _index368
				(_2121_v4).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rjcj"), (_index368).Int())
				var _2123_v6 _dafny.Array
				_ = _2123_v6
				var _nw339 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw339
				_2123_v6 = _nw339
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_2123_v6), 0))
				_ = _index369
				(_2123_v6).ArraySet1CodePoint(_2120_v3, (_index369).Int())
				var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_2123_v6), 0))
				_ = _index370
				(_2123_v6).ArraySet1CodePoint(_2120_v3, (_index370).Int())
				var _2124_v7 _dafny.Map
				_ = _2124_v7
				_2124_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2116_cf2, p0)
				var _2125_v8 _dafny.Map
				_ = _2125_v8
				_2125_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v7, p0)
				(globalState).F4 = ((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_2126_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2127_i1 _dafny.Int) _dafny.CodePoint {
						return _2126_v3
					}
				})(_2120_v3)))).Cardinality())).Times((_2125_v8).Cardinality())).Cmp(p2) >= 0
			} else if _source37.Is_DC2() {
				(globalState).F4 = (p0) || (p0)
				var _2128_v9 _dafny.Map
				_ = _2128_v9
				_2128_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _2129_v10 _dafny.MultiSet
				_ = _2129_v10
				_2129_v10 = _dafny.MultiSetOf(_2128_v9)
				var _2130_v11 _dafny.Set
				_ = _2130_v11
				_2130_v11 = _dafny.SetOf(_2112_i0, _2112_i0)
				var _2131_v12 _dafny.MultiSet
				_ = _2131_v12
				_2131_v12 = _dafny.MultiSetOf(p0)
				var _2132_v13 _dafny.Sequence
				_ = _2132_v13
				_2132_v13 = _dafny.SeqOf(p2)
				var _2133_v14 _dafny.Array
				_ = _2133_v14
				var _nwElement0_68 bool = true
				_ = _nwElement0_68
				var _nw340 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(18))
				_ = _nw340
				(_nw340).ArraySet1(_nwElement0_68, 0)
				(_nw340).ArraySet1((false) && (p0), 1)
				(_nw340).ArraySet1(p0, 2)
				(_nw340).ArraySet1(p0, 3)
				(_nw340).ArraySet1(true, 4)
				(_nw340).ArraySet1(p0, 5)
				(_nw340).ArraySet1((_2129_v10).Contains(_2128_v9), 6)
				(_nw340).ArraySet1(p0, 7)
				(_nw340).ArraySet1(p0, 8)
				(_nw340).ArraySet1((_2130_v11).IsProperSubsetOf(_2130_v11), 9)
				(_nw340).ArraySet1(((_2130_v11).Cardinality()).Cmp(_2112_i0) == 0, 10)
				(_nw340).ArraySet1(((func() _dafny.Int {
					if (_2131_v12).Contains(false) {
						return (_2131_v12).Multiplicity(false)
					}
					return _dafny.IntOfUint32((_2132_v13).Cardinality())
				})()).Cmp(_2112_i0) <= 0, 11)
				(_nw340).ArraySet1(p0, 12)
				(_nw340).ArraySet1(!(p0) || (!(p0)), 13)
				(_nw340).ArraySet1((p0) == (p0), 14)
				(_nw340).ArraySet1(p0, 15)
				(_nw340).ArraySet1(p0, 16)
				(_nw340).ArraySet1(true, 17)
				_2133_v14 = _nw340
				var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_2133_v14), 0))
				_ = _index371
				(_2133_v14).ArraySet1(p0, (_index371).Int())
				var _2134_v15 bool
				_ = _2134_v15
				_2134_v15 = p0
				var _2135_v16 _dafny.Sequence
				_ = _2135_v16
				_2135_v16 = _dafny.SeqOf(p0, (_2134_v15))
				var _2136_v17 D2
				_ = _2136_v17
				_2136_v17 = Companion_D2_.Create_DC5_(_2133_v14)
				var _2137_v18 _dafny.MultiSet
				_ = _2137_v18
				_2137_v18 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2132_v13).Cardinality()), _dafny.IntOfUint32(((_this).Fm8(p2, p0, _2112_i0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}(func(_2138_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("iutevq")
				})), globalState)).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2133_v14, (_2136_v17).Dtor_cf5())).Update(_2133_v14, _2133_v14)).Cardinality())
				var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_2133_v14), 0))
				_ = _index372
				(_2133_v14).ArraySet1((_2135_v16).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_2137_v18).Cardinality()), _dafny.IntOfUint32((_2135_v16).Cardinality()))).Uint32()).(bool), (_index372).Int())
				var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_2133_v14), 0))
				_ = _index373
				(_2133_v14).ArraySet1(p0, (_index373).Int())
				(globalState).F3 = Companion_Default___.SafeDivisionInt(((func() _dafny.Int {
					if (_2137_v18).Contains(p2) {
						return (_2137_v18).Multiplicity(p2)
					}
					return _dafny.IntOfUint32((_2132_v13).Cardinality())
				})()).Times(_2112_i0), _2112_i0)
			} else if _source37.Is_DC0() {
				var _2139___mcc_h2 _dafny.Int = _source37.Get_().(D0_DC0).Cf0
				_ = _2139___mcc_h2
				var _2140_cf0 _dafny.Int = _2139___mcc_h2
				_ = _2140_cf0
				var _2141_v19 _dafny.Sequence
				_ = _2141_v19
				_2141_v19 = _dafny.UnicodeSeqOfUtf8Bytes("dlnuduto")
				var _2142_v20 D2
				_ = _2142_v20
				_2142_v20 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_2141_v19).Cardinality()), p0, p0)
				var _2143_v21 D0
				_ = _2143_v21
				_2143_v21 = Companion_D0_.Create_DC1_((_2142_v20).Dtor_cf6(), _dafny.IntOfInt64(-464))
				var _2144_v22 _dafny.Map
				_ = _2144_v22
				_2144_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2143_v21)
				_2144_v22 = _2144_v22
				(globalState).F3 = _2112_i0
				var _2145_v23 _dafny.Array
				_ = _2145_v23
				var _nw341 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw341
				_2145_v23 = _nw341
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_57
				var _nw342 _dafny.Array
				_ = _nw342
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw342 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) _dafny.Sequence = func(_2146_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("irpd")
					}
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw342 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw342).ArraySet1(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw342).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2145_v23 = _nw342
				var _2147_v24 _dafny.Set
				_ = _2147_v24
				_2147_v24 = _dafny.SetOf(p0, p0)
				var _2148_v25 _dafny.Map
				_ = _2148_v25
				_2148_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2147_v24, (_this).Fm9(p0, _2147_v24, globalState))
				_2148_v25 = (_2148_v25).Update(_2147_v24, _2140_cf0)
			} else {
				var _2149___mcc_h3 D0 = _source37.Get_().(D0_DC3).Cf3
				_ = _2149___mcc_h3
				var _2150_cf3 D0 = _2149___mcc_h3
				_ = _2150_cf3
				var _2151_v26 _dafny.Array
				_ = _2151_v26
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw343
				_2151_v26 = _nw343
				var _2152_v27 *C4
				_ = _2152_v27
				var _nw344 *C4 = New_C4_()
				_ = _nw344
				_nw344.Ctor__(_2151_v26)
				_2152_v27 = _nw344
				var _2153_v28 _dafny.MultiSet
				_ = _2153_v28
				_2153_v28 = _dafny.MultiSetOf(_2112_i0, _dafny.IntOfInt64(492))
				var _2154_v29 _dafny.Sequence
				_ = _2154_v29
				_2154_v29 = _dafny.UnicodeSeqOfUtf8Bytes("sjohkdo")
				var _2155_v30 _dafny.Sequence
				_ = _2155_v30
				_2155_v30 = _dafny.SeqOf(p0)
				var _2156_v31 _dafny.Sequence
				_ = _2156_v31
				_2156_v31 = _dafny.SeqOf(p2)
				var _2157_v32 _dafny.Map
				_ = _2157_v32
				_2157_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2153_v28).Cardinality(), _2154_v29)
				var _2158_v33 _dafny.Array
				_ = _2158_v33
				var _nwElement0_69 bool = (_2153_v28).Contains(_2112_i0)
				_ = _nwElement0_69
				var _nw345 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(17))
				_ = _nw345
				(_nw345).ArraySet1(_nwElement0_69, 0)
				(_nw345).ArraySet1(p0, 1)
				(_nw345).ArraySet1((_2152_v27).Fm13(p0, _2154_v29, p0, globalState), 2)
				(_nw345).ArraySet1(!(!(!(p0))), 3)
				(_nw345).ArraySet1(!((_2155_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2156_v31).Cardinality()), _dafny.IntOfUint32((_2155_v30).Cardinality()))).Uint32()).(bool)) || (p0), 4)
				(_nw345).ArraySet1((p2).Cmp(_2112_i0) == 0, 5)
				(_nw345).ArraySet1((_2152_v27).Fm13(p0, _2154_v29, p0, globalState), 6)
				(_nw345).ArraySet1(true, 7)
				(_nw345).ArraySet1(p0, 8)
				(_nw345).ArraySet1(p0, 9)
				(_nw345).ArraySet1((func() bool {
					if p0 {
						return p0
					}
					return p0
				})(), 10)
				(_nw345).ArraySet1(p0, 11)
				(_nw345).ArraySet1(p0, 12)
				(_nw345).ArraySet1(p0, 13)
				(_nw345).ArraySet1(p0, 14)
				(_nw345).ArraySet1(p0, 15)
				(_nw345).ArraySet1((_2152_v27).Fm13(p0, (func() _dafny.Sequence {
					if (_2157_v32).Contains(_dafny.IntOfInt64(823)) {
						return (_2157_v32).Get(_dafny.IntOfInt64(823)).(_dafny.Sequence)
					}
					return _2154_v29
				})(), p0, globalState), 16)
				_2158_v33 = _nw345
				var _2159_v34 _dafny.Set
				_ = _2159_v34
				_2159_v34 = _dafny.SetOf(p0, true)
				var _rhs370 _dafny.Int = _dafny.IntOfInt64(-61)
				_ = _rhs370
				var _rhs371 bool = (_dafny.SetOf(p0, false, !(p0), p0, p0)).IsSubsetOf((_2159_v34).Intersection(_dafny.SetOf((_2152_v27).Fm3(_2156_v31, _dafny.MultiSetOf(p2), p0, globalState))))
				_ = _rhs371
				var _rhs372 _dafny.Array = _2158_v33
				_ = _rhs372
				var _rhs373 _dafny.Int = _2112_i0
				_ = _rhs373
				var _lhs287 *GlobalState = globalState
				_ = _lhs287
				var _lhs288 *GlobalState = globalState
				_ = _lhs288
				var _lhs289 *GlobalState = globalState
				_ = _lhs289
				_lhs287.F8 = _rhs370
				_lhs288.F4 = _rhs371
				_2158_v33 = _rhs372
				_lhs289.F8 = _rhs373
				var _2160_v35 D4
				_ = _2160_v35
				_2160_v35 = Companion_D4_.Create_DC10_(_2155_v30)
				_2160_v35 = _2160_v35
				var _2161_v36 _dafny.Map
				_ = _2161_v36
				_2161_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p2, _2112_i0), p0)
				_2161_v36 = (_2161_v36).Update(_2112_i0, (p2).Cmp(_2112_i0) != 0)
			}
			var _2162_v37 _dafny.CodePoint
			_ = _2162_v37
			_2162_v37 = _dafny.CodePoint('s')
			_2162_v37 = _2162_v37
			_2162_v37 = _2162_v37
			var _2163_v38 _dafny.Array
			_ = _2163_v38
			var _nw346 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw346
			_2163_v38 = _nw346
			var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_2163_v38), 0))
			_ = _index374
			(_2163_v38).ArraySet1(Companion_Default___.Fm31(false, !(p0), p0, globalState), (_index374).Int())
			var _2164_v39 _dafny.Sequence
			_ = _2164_v39
			_2164_v39 = _dafny.SeqOf(p0)
			var _2165_v40 _dafny.Map
			_ = _2165_v40
			_2165_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_2163_v38), 0))
			_ = _index375
			(_2163_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2164_v39, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2164_v39, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-670), _dafny.IntOfUint32((_2164_v39).Cardinality()))).Uint32(), (func() bool {
				if (_2165_v40).Contains(p0) {
					return (_2165_v40).Get(p0).(bool)
				}
				return p0
			})()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2164_v39, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-670), _dafny.IntOfUint32((_2164_v39).Cardinality()))).Uint32(), (func() bool {
				if (_2165_v40).Contains(p0) {
					return (_2165_v40).Get(p0).(bool)
				}
				return p0
			})())).Cardinality()))).Uint32(), p0)), _dafny.Companion_Sequence_.Concatenate(_2164_v39, _2164_v39)), (_index375).Int())
		}
		var _hi13 _dafny.Int = (p2).Times(p2)
		_ = _hi13
		for _2166_i4 := p2; _2166_i4.Cmp(_hi13) < 0; _2166_i4 = _2166_i4.Plus(_dafny.One) {
			var _2167_v41 _dafny.Sequence
			_ = _2167_v41
			_2167_v41 = _dafny.SeqOf(p0)
			_2167_v41 = _2167_v41
			(globalState).F8 = p2
			var _2168_v42 _dafny.CodePoint
			_ = _2168_v42
			_2168_v42 = _dafny.CodePoint('a')
			var _2169_v43 _dafny.Sequence
			_ = _2169_v43
			_2169_v43 = _dafny.UnicodeSeqOfUtf8Bytes("vstrg")
			(globalState).F4 = !_dafny.Companion_Sequence_.Contains(_2169_v43, _2168_v42)
			var _2170_v44 *C6
			_ = _2170_v44
			var _nw347 *C6 = New_C6_()
			_ = _nw347
			_nw347.Ctor__(_2166_i4, _dafny.Companion_Sequence_.Concatenate(_2169_v43, _2169_v43))
			_2170_v44 = _nw347
		}
		(globalState).F4 = !(true)
		var _2171_v45 _dafny.Sequence
		_ = _2171_v45
		_2171_v45 = _dafny.SeqOf(p0, p0, p0, p0)
		var _2172_v46 _dafny.Sequence
		_ = _2172_v46
		_2172_v46 = _dafny.UnicodeSeqOfUtf8Bytes("am")
		var _2173_v47 _dafny.CodePoint
		_ = _2173_v47
		_2173_v47 = _dafny.CodePoint('b')
		_2171_v45 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, p0), _2171_v45), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (p1).Contains(p0) {
				return (p1).Get(p0).(_dafny.Sequence)
			}
			return _2172_v46
		})(), _2172_v46), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (p1).Contains(p0) {
				return (p1).Get(p0).(_dafny.Sequence)
			}
			return _2172_v46
		})(), _2172_v46)).Cardinality()))).Uint32(), _2173_v47)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, p0), _2171_v45)).Cardinality()))).Uint32(), false)
		if Companion_Default___.Fm49(p0, globalState) {
			var _2174_v48 _dafny.Array
			_ = _2174_v48
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_58
			var _nw348 _dafny.Array
			_ = _nw348
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw348 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) bool = func(_2175_i5 _dafny.Int) bool {
					return false
				}
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw348 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw348).ArraySet1(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw348).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_2174_v48 = _nw348
			var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2174_v48), 0))
			_ = _index376
			(_2174_v48).ArraySet1(p0, (_index376).Int())
			var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_2174_v48), 0))
			_ = _index377
			(_2174_v48).ArraySet1(p0, (_index377).Int())
			var _2176_v49 _dafny.Set
			_ = _2176_v49
			_2176_v49 = _dafny.SetOf(_2174_v48)
			var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2174_v48), 0))
			_ = _index378
			var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_2174_v48), 0))
			_ = _index379
			var _rhs374 _dafny.Int = ((_dafny.SetOf(_2174_v48)).Union(_2176_v49)).Cardinality()
			_ = _rhs374
			var _rhs375 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return p2
				}
				return (p2).Minus(p2)
			})()
			_ = _rhs375
			var _rhs376 bool = p0
			_ = _rhs376
			var _rhs377 bool = p0
			_ = _rhs377
			var _lhs290 *GlobalState = globalState
			_ = _lhs290
			var _lhs291 *GlobalState = globalState
			_ = _lhs291
			var _lhs292 _dafny.Array = _2174_v48
			_ = _lhs292
			var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2174_v48), 0))
			_ = _lhs293
			var _lhs294 _dafny.Array = _2174_v48
			_ = _lhs294
			var _lhs295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_2174_v48), 0))
			_ = _lhs295
			_lhs290.F8 = _rhs374
			_lhs291.F3 = _rhs375
			(_lhs292).ArraySet1(_rhs376, (_lhs293).Int())
			(_lhs294).ArraySet1(_rhs377, (_lhs295).Int())
			var _2177_v50 _dafny.Array
			_ = _2177_v50
			var _len0_59 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_59
			var _nw349 _dafny.Array
			_ = _nw349
			if _len0_59.Cmp(_dafny.Zero) == 0 {
				_nw349 = _dafny.NewArray(_len0_59)
			} else {
				var _init59 func(_dafny.Int) _dafny.Int = (func(_2178_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2179_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_2179_i6, _2178_p2)
					}
				})(p2)
				_ = _init59
				var _element0_59 = _init59(_dafny.Zero)
				_ = _element0_59
				_nw349 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
				(_nw349).ArraySet1(_element0_59, 0)
				var _nativeLen0_59 = (_len0_59).Int()
				_ = _nativeLen0_59
				for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
					(_nw349).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
				}
			}
			_2177_v50 = _nw349
			var _2180_v51 _dafny.Array
			_ = _2180_v51
			_2180_v51 = _2177_v50
			var _2181_v52 _dafny.Array
			_ = _2181_v52
			var _nwElement0_70 _dafny.Array = _2180_v51
			_ = _nwElement0_70
			var _nw350 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(9))
			_ = _nw350
			(_nw350).ArraySet1(_nwElement0_70, 0)
			(_nw350).ArraySet1(_2180_v51, 1)
			(_nw350).ArraySet1((func() _dafny.Array {
				if p0 {
					return _2177_v50
				}
				return _2180_v51
			})(), 2)
			(_nw350).ArraySet1(_2180_v51, 3)
			(_nw350).ArraySet1(_2177_v50, 4)
			(_nw350).ArraySet1(_2177_v50, 5)
			(_nw350).ArraySet1(_2180_v51, 6)
			(_nw350).ArraySet1(_2180_v51, 7)
			(_nw350).ArraySet1(_2180_v51, 8)
			_2181_v52 = _nw350
			var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_2181_v52), 0))
			_ = _index380
			(_2181_v52).ArraySet1(_2180_v51, (_index380).Int())
			var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_2181_v52), 0))
			_ = _index381
			(_2181_v52).ArraySet1(_2177_v50, (_index381).Int())
			var _2182_v53 _dafny.Map
			_ = _2182_v53
			_2182_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			_2182_v53 = (_2182_v53).Update(_dafny.IntOfUint32((_2171_v45).Cardinality()), p2)
			var _2183_v54 _dafny.Set
			_ = _2183_v54
			_2183_v54 = _dafny.SetOf(_dafny.IntOfInt64(-799), p2)
			var _2184_v55 D3
			_ = _2184_v55
			_2184_v55 = Companion_D3_.Create_DC9_()
			var _2185_v56 _dafny.Set
			_ = _2185_v56
			_2185_v56 = _dafny.SetOf(Companion_Default___.Fm50(globalState), _2184_v55, _2184_v55)
			var _2186_v57 _dafny.Sequence
			_ = _2186_v57
			_2186_v57 = _dafny.SeqOf(_2185_v56, _2185_v56)
			(globalState).F8 = (((_2183_v54).Union(_2183_v54)).Cardinality()).Minus((func() _dafny.Int {
				if p0 {
					return ((_2186_v57).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2186_v57).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
				}
				return _dafny.IntOfUint32((_2171_v45).Cardinality())
			})())
			var _2187_v58 _dafny.Array
			_ = _2187_v58
			var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw351
			_2187_v58 = _nw351
			var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2187_v58), 0))
			_ = _index382
			(_2187_v58).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2172_v46, _2172_v46), (_index382).Int())
			var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2187_v58), 0))
			_ = _index383
			(_2187_v58).ArraySet1(_dafny.Companion_Sequence_.Update(_2172_v46, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2172_v46).Cardinality()))).Uint32(), _2173_v47), (_index383).Int())
		} else {
			var _2188_v59 _dafny.Map
			_ = _2188_v59
			_2188_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2172_v46, p2)
			var _2189_v60 _dafny.Map
			_ = _2189_v60
			_2189_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			_2188_v59 = (_2188_v59).Update(_2172_v46, (_dafny.Zero).Minus((func() _dafny.Int {
				if (_2189_v60).Contains(p0) {
					return (_2189_v60).Get(p0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2, p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(800))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}((func(_2190_v47 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2191_i7 _dafny.Int) _dafny.CodePoint {
						return _2190_v47
					}
				})(_2173_v47)))).Cardinality())), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2172_v46).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p2, p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(800))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_2192_v47 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2193_i7 _dafny.Int) _dafny.CodePoint {
						return _2192_v47
					}
				})(_2173_v47)))).Cardinality()))).Cardinality()))).Uint32(), p2)).Cardinality()))
			})()))
			var _2194_v61 T0
			_ = _2194_v61
			var _nw352 *C2 = New_C2_()
			_ = _nw352
			_nw352.Ctor__()
			_2194_v61 = _nw352
			var _2195_v62 _dafny.Sequence
			_ = _2195_v62
			_2195_v62 = _dafny.SeqOf(_2194_v61)
			_2194_v61 = (_2195_v62).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2195_v62).Cardinality()))).Uint32()).(T0)
			_2172_v46 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tvyfq"), _2172_v46), _2172_v46)
			var _2196_v63 *C3
			_ = _2196_v63
			var _nw353 *C3 = New_C3_()
			_ = _nw353
			_nw353.Ctor__()
			_2196_v63 = _nw353
			var _2197_v64 _dafny.Array
			_ = _2197_v64
			var _nw354 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw354
			_2197_v64 = _nw354
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_2197_v64), 0))
			_ = _index384
			(_2197_v64).ArraySet1(_2172_v46, (_index384).Int())
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_2197_v64), 0))
			_ = _index385
			(_2197_v64).ArraySet1(_2172_v46, (_index385).Int())
		}
		var _2198_v65 _dafny.Array
		_ = _2198_v65
		var _nw355 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(26))
		_ = _nw355
		_2198_v65 = _nw355
		var _2199_v66 _dafny.Array
		_ = _2199_v66
		var _nwElement0_71 _dafny.CodePoint = _2173_v47
		_ = _nwElement0_71
		var _nw356 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(13))
		_ = _nw356
		(_nw356).ArraySet1CodePoint(_nwElement0_71, 0)
		(_nw356).ArraySet1CodePoint(_2173_v47, 1)
		(_nw356).ArraySet1CodePoint(_2173_v47, 2)
		(_nw356).ArraySet1CodePoint(_2173_v47, 3)
		(_nw356).ArraySet1CodePoint(_2173_v47, 4)
		(_nw356).ArraySet1CodePoint(_2173_v47, 5)
		(_nw356).ArraySet1CodePoint(_2173_v47, 6)
		(_nw356).ArraySet1CodePoint(_2173_v47, 7)
		(_nw356).ArraySet1CodePoint(_2173_v47, 8)
		(_nw356).ArraySet1CodePoint(_2173_v47, 9)
		(_nw356).ArraySet1CodePoint(_2173_v47, 10)
		(_nw356).ArraySet1CodePoint(_2173_v47, 11)
		(_nw356).ArraySet1CodePoint(_2173_v47, 12)
		_2199_v66 = _nw356
		var _2200_v67 _dafny.Set
		_ = _2200_v67
		_2200_v67 = _dafny.SetOf(_2199_v66, _2199_v66)
		var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_2198_v65), 0))
		_ = _index386
		(_2198_v65).ArraySet1(_2200_v67, (_index386).Int())
		var _2201_v68 D19
		_ = _2201_v68
		_2201_v68 = Companion_D19_.Create_DC39_(p0, _dafny.IntOfInt64(762))
		var _2202_v69 _dafny.Array
		_ = _2202_v69
		var _nwElement0_72 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("oxey"), _dafny.UnicodeSeqOfUtf8Bytes("brbh")))
		_ = _nwElement0_72
		var _nw357 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(6))
		_ = _nw357
		(_nw357).ArraySet1(_nwElement0_72, 0)
		(_nw357).ArraySet1(p0, 1)
		(_nw357).ArraySet1((_2201_v68).Dtor_cf58(), 2)
		(_nw357).ArraySet1(p0, 3)
		(_nw357).ArraySet1(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm14(p2, globalState), _2172_v46), 4)
		(_nw357).ArraySet1(false, 5)
		_2202_v69 = _nw357
		var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_2202_v69), 0))
		_ = _index387
		(_2202_v69).ArraySet1(p0, (_index387).Int())
		var _2203_v70 _dafny.Map
		_ = _2203_v70
		_2203_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2200_v67)
		var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_2198_v65), 0))
		_ = _index388
		var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_2202_v69), 0))
		_ = _index389
		var _rhs378 _dafny.Set = (_2200_v67).Difference((func() _dafny.Set {
			if (_2203_v70).Contains(p2) {
				return (_2203_v70).Get(p2).(_dafny.Set)
			}
			return _2200_v67
		})())
		_ = _rhs378
		var _rhs379 bool = p0
		_ = _rhs379
		var _rhs380 bool = (Companion_Default___.Fm51(p2, globalState))
		_ = _rhs380
		var _lhs296 _dafny.Array = _2198_v65
		_ = _lhs296
		var _lhs297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_2198_v65), 0))
		_ = _lhs297
		var _lhs298 *GlobalState = globalState
		_ = _lhs298
		var _lhs299 _dafny.Array = _2202_v69
		_ = _lhs299
		var _lhs300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_2202_v69), 0))
		_ = _lhs300
		(_lhs296).ArraySet1(_rhs378, (_lhs297).Int())
		_lhs298.F4 = _rhs379
		(_lhs299).ArraySet1(_rhs380, (_lhs300).Int())
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f9 bool
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f9 = false
	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) Ctor__(f9 bool) {
	{
		(_this)._f9 = f9
	}
}
func (_this *C9) M2(globalState *GlobalState) (bool, _dafny.Int, _dafny.Set) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var _source38 D0 = Companion_Default___.Fm5(globalState)
		_ = _source38
		if _source38.Is_DC1() {
			var _2204___mcc_h0 _dafny.Int = _source38.Get_().(D0_DC1).Cf1
			_ = _2204___mcc_h0
			var _2205___mcc_h1 _dafny.Int = _source38.Get_().(D0_DC1).Cf2
			_ = _2205___mcc_h1
			var _2206_cf2 _dafny.Int = _2205___mcc_h1
			_ = _2206_cf2
			var _2207_cf1 _dafny.Int = _2204___mcc_h0
			_ = _2207_cf1
			var _2208_v0 _dafny.Array
			_ = _2208_v0
			var _nw358 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw358
			_2208_v0 = _nw358
			_2208_v0 = _2208_v0
			var _2209_v1 _dafny.CodePoint
			_ = _2209_v1
			_2209_v1 = _dafny.CodePoint('h')
			var _2210_v2 _dafny.Map
			_ = _2210_v2
			_2210_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2209_v1)
			var _2211_v3 _dafny.Map
			_ = _2211_v3
			_2211_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2207_cf1).Plus(_2207_cf1), _2210_v2)
			var _2212_v4 _dafny.Map
			_ = _2212_v4
			_2212_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2207_cf1)
			var _2213_v5 _dafny.Sequence
			_ = _2213_v5
			_2213_v5 = _dafny.SeqOf((_this).F9())
			_2211_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6((func() _dafny.Int {
				if (_2212_v4).Contains((_this).F9()) {
					return (_2212_v4).Get((_this).F9()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_2213_v5).Cardinality())
			})(), globalState), Companion_Default___.Fm7(_2209_v1, globalState))
			r1 = _2206_cf2
			if (_this).F9() {
				var _2214_v6 *C7
				_ = _2214_v6
				var _nw359 *C7 = New_C7_()
				_ = _nw359
				_nw359.Ctor__()
				_2214_v6 = _nw359
				var _2215_v7 _dafny.Array
				_ = _2215_v7
				var _nw360 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw360
				_2215_v7 = _nw360
				var _2216_v8 *C0
				_ = _2216_v8
				var _nw361 *C0 = New_C0_()
				_ = _nw361
				_nw361.Ctor__()
				_2216_v8 = _nw361
				var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_2215_v7), 0))
				_ = _index390
				(_2215_v7).ArraySet1(_2216_v8, (_index390).Int())
				var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_2215_v7), 0))
				_ = _index391
				(_2215_v7).ArraySet1(_2216_v8, (_index391).Int())
				var _2217_v9 _dafny.Sequence
				_ = _2217_v9
				_2217_v9 = _dafny.SeqOf(_2207_cf1)
				(globalState).F4 = (_2214_v6).Fm3(_2217_v9, _dafny.MultiSetFromSeq(_2217_v9), (_this).F9(), globalState)
				var _2218_v10 *C8
				_ = _2218_v10
				var _nw362 *C8 = New_C8_()
				_ = _nw362
				_nw362.Ctor__()
				_2218_v10 = _nw362
				var _2219_v11 D2
				_ = _2219_v11
				_2219_v11 = Companion_D2_.Create_DC5_(_2208_v0)
				var _2220_v12 _dafny.Array
				_ = _2220_v12
				var _nwElement0_73 _dafny.Array = _2208_v0
				_ = _nwElement0_73
				var _nw363 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(11))
				_ = _nw363
				(_nw363).ArraySet1(_nwElement0_73, 0)
				(_nw363).ArraySet1(_2208_v0, 1)
				(_nw363).ArraySet1(_2208_v0, 2)
				(_nw363).ArraySet1(_2208_v0, 3)
				(_nw363).ArraySet1(_2208_v0, 4)
				(_nw363).ArraySet1(_2208_v0, 5)
				(_nw363).ArraySet1(_2208_v0, 6)
				(_nw363).ArraySet1(_2208_v0, 7)
				(_nw363).ArraySet1(_2208_v0, 8)
				(_nw363).ArraySet1((_2219_v11).Dtor_cf5(), 9)
				(_nw363).ArraySet1(_2208_v0, 10)
				_2220_v12 = _nw363
				var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_2220_v12), 0))
				_ = _index392
				(_2220_v12).ArraySet1(_2208_v0, (_index392).Int())
				var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_2220_v12), 0))
				_ = _index393
				(_2220_v12).ArraySet1(_2208_v0, (_index393).Int())
			} else {
				(globalState).F8 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(142), (_dafny.Zero).Minus(_2206_cf2))
				var _2221_v13 _dafny.Array
				_ = _2221_v13
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_60
				var _nw364 _dafny.Array
				_ = _nw364
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw364 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) _dafny.Int = (func(_2222_cf1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2223_i0 _dafny.Int) _dafny.Int {
							return (_2223_i0).Times(_2222_cf1)
						}
					})(_2207_cf1)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw364 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw364).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw364).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_2221_v13 = _nw364
				var _2224_v14 _dafny.Map
				_ = _2224_v14
				_2224_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), false)
				var _2225_v15 _dafny.Map
				_ = _2225_v15
				_2225_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2221_v13, _2224_v14)
				var _2226_v16 _dafny.Map
				_ = _2226_v16
				_2226_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2207_cf1).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_2225_v15).Cardinality()))), (_this).F9())
				var _2227_v18 _dafny.Set
				_ = _2227_v18
				_2227_v18 = _dafny.SetOf(_dafny.IntOfInt64(-25))
				_2226_v16 = (_2226_v16).Merge((func() _dafny.Map {
					var _coll102 = _dafny.NewMapBuilder()
					_ = _coll102
					for _iter106 := _dafny.Iterate((_2227_v18).Elements()); ; {
						_compr_102, _ok106 := _iter106()
						if !_ok106 {
							break
						}
						var _2228_v17 _dafny.Int
						_2228_v17 = interface{}(_compr_102).(_dafny.Int)
						if (_2227_v18).Contains(_2228_v17) {
							_coll102.Add((_2228_v17).Minus(_2207_cf1), (_this).F9())
						}
					}
					return _coll102.ToMap()
				}()).Update(_2207_cf1, (_this).F9()))
				var _2229_v19 _dafny.Sequence
				_ = _2229_v19
				_2229_v19 = _dafny.UnicodeSeqOfUtf8Bytes("fml")
				var _2230_v20 *C6
				_ = _2230_v20
				var _nw365 *C6 = New_C6_()
				_ = _nw365
				_nw365.Ctor__(_2206_cf2, _dafny.Companion_Sequence_.Concatenate(_2229_v19, _2229_v19))
				_2230_v20 = _nw365
				var _2231_v21 _dafny.Set
				_ = _2231_v21
				_2231_v21 = _dafny.SetOf((_this).F9(), (_this).F9(), (_this).F9(), (_2213_v5).Select((Companion_Default___.SafeIndex((_2230_v20).F10(), _dafny.IntOfUint32((_2213_v5).Cardinality()))).Uint32()).(bool), (_this).F9())
				(globalState).F3 = Companion_Default___.SafeDivisionInt((_2231_v21).Cardinality(), (_dafny.IntOfInt64(-741)).Times(Companion_Default___.Fm6(_2206_cf2, globalState)))
				var _2232_v22 D12
				_ = _2232_v22
				_2232_v22 = Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_2224_v14).Contains((_this).F9()) {
						return (_2224_v14).Get((_this).F9()).(bool)
					}
					return (_this).F9()
				})(), _2206_cf2))
				var _2233_v23 _dafny.Map
				_ = _2233_v23
				_2233_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2232_v22, true)
				var _2234_v24 _dafny.Map
				_ = _2234_v24
				_2234_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2206_cf2, _2233_v23)
				var _rhs381 _dafny.Int = ((_2234_v24).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_2230_v20).F11()).Cardinality()), _2233_v23))).Cardinality()
				_ = _rhs381
				var _rhs382 bool = (_this).F9()
				_ = _rhs382
				var _lhs301 *GlobalState = globalState
				_ = _lhs301
				_lhs301.F3 = _rhs381
				r0 = _rhs382
			}
		} else if _source38.Is_DC2() {
			var _2235_v25 _dafny.Sequence
			_ = _2235_v25
			_2235_v25 = _dafny.SeqOf((_this).F9(), !((_this).F9()))
			var _2236_v26 _dafny.Map
			_ = _2236_v26
			_2236_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_2235_v25).Cardinality()))
			var _2237_v27 _dafny.Int
			_ = _2237_v27
			_2237_v27 = _dafny.IntOfInt64(-901)
			var _2238_v29 _dafny.Sequence
			_ = _2238_v29
			_2238_v29 = _dafny.UnicodeSeqOfUtf8Bytes("ivrfg")
			var _2239_v30 _dafny.Sequence
			_ = _2239_v30
			_2239_v30 = _dafny.SeqOf(_dafny.IntOfUint32((_2238_v29).Cardinality()))
			var _2240_v31 _dafny.Array
			_ = _2240_v31
			var _nwElement0_74 _dafny.Map = _2236_v26
			_ = _nwElement0_74
			var _nw366 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(18))
			_ = _nw366
			(_nw366).ArraySet1(_nwElement0_74, 0)
			(_nw366).ArraySet1(_2236_v26, 1)
			(_nw366).ArraySet1(_2236_v26, 2)
			(_nw366).ArraySet1(_2236_v26, 3)
			(_nw366).ArraySet1(Companion_Default___.Fm46(globalState), 4)
			(_nw366).ArraySet1(_2236_v26, 5)
			(_nw366).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2237_v27), 6)
			(_nw366).ArraySet1(_2236_v26, 7)
			(_nw366).ArraySet1(_2236_v26, 8)
			(_nw366).ArraySet1(_2236_v26, 9)
			(_nw366).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2237_v27), 10)
			(_nw366).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2237_v27, (_this).F9())).Update((func() _dafny.Map {
				var _coll103 = _dafny.NewMapBuilder()
				_ = _coll103
				for _iter107 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(993), _dafny.IntOfInt64(539))); ; {
					_compr_103, _ok107 := _iter107()
					if !_ok107 {
						break
					}
					var _2241_v28 _dafny.Int
					_2241_v28 = interface{}(_compr_103).(_dafny.Int)
					if ((_dafny.IntOfInt64(993)).Cmp(_2241_v28) <= 0) && ((_2241_v28).Cmp(_dafny.IntOfInt64(539)) < 0) {
						_coll103.Add((_2241_v28).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jkmh")).Cardinality())), (_this).F9())
					}
				}
				return _coll103.ToMap()
			}()).Cardinality(), (_this).F9())).Cardinality()), 11)
			(_nw366).ArraySet1(_2236_v26, 12)
			(_nw366).ArraySet1(_2236_v26, 13)
			(_nw366).ArraySet1(_2236_v26, 14)
			(_nw366).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_2237_v27, (_dafny.MultiSetFromSeq(_2239_v30)).Cardinality(), _2237_v27, _2237_v27)).Cardinality())), 15)
			(_nw366).ArraySet1(_2236_v26, 16)
			(_nw366).ArraySet1(_2236_v26, 17)
			_2240_v31 = _nw366
			var _2242_v32 *C4
			_ = _2242_v32
			var _nw367 *C4 = New_C4_()
			_ = _nw367
			_nw367.Ctor__(_2240_v31)
			_2242_v32 = _nw367
			var _2243_v33 _dafny.MultiSet
			_ = _2243_v33
			_2243_v33 = _dafny.MultiSetOf((_this).F9())
			var _2244_v34 _dafny.Map
			_ = _2244_v34
			_2244_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.UnicodeSeqOfUtf8Bytes("fal"))
			var _2245_v35 D4
			_ = _2245_v35
			_2245_v35 = Companion_D4_.Create_DC11_((_dafny.IntOfInt64(417)).Cmp((((_2243_v33).Update(true, Companion_Default___.Abs(_dafny.IntOfInt64(836)))).Update((_this).F9(), Companion_Default___.Abs(_2237_v27))).Cardinality()) < 0, ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2244_v34).Contains((_this).F9()) {
					return (_2244_v34).Get((_this).F9()).(_dafny.Sequence)
				}
				return _2238_v29
			})()).Cardinality())))).Cmp(_2237_v27) >= 0, _2237_v27, _2238_v29)
			var _2246_v36 _dafny.Map
			_ = _2246_v36
			_2246_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2235_v25, (_this).F9())
			var _2247_v37 _dafny.Map
			_ = _2247_v37
			_2247_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			var _2248_v38 D12
			_ = _2248_v38
			_2248_v38 = Companion_D12_.Create_DC27_(_dafny.IntOfInt64(16), _2237_v27, _2237_v27, (_2247_v37).Cardinality())
			var _2249_v39 _dafny.Array
			_ = _2249_v39
			var _nwElement0_75 bool = (_this).F9()
			_ = _nwElement0_75
			var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(20))
			_ = _nw368
			(_nw368).ArraySet1(_nwElement0_75, 0)
			(_nw368).ArraySet1((_this).F9(), 1)
			(_nw368).ArraySet1((_this).F9(), 2)
			(_nw368).ArraySet1((_this).F9(), 3)
			(_nw368).ArraySet1((_this).F9(), 4)
			(_nw368).ArraySet1((_this).F9(), 5)
			(_nw368).ArraySet1((_this).F9(), 6)
			(_nw368).ArraySet1((_this).F9(), 7)
			(_nw368).ArraySet1((_this).F9(), 8)
			(_nw368).ArraySet1((_this).F9(), 9)
			(_nw368).ArraySet1((_this).F9(), 10)
			(_nw368).ArraySet1((_this).F9(), 11)
			(_nw368).ArraySet1((_this).F9(), 12)
			(_nw368).ArraySet1(false, 13)
			(_nw368).ArraySet1(false, 14)
			(_nw368).ArraySet1((_this).F9(), 15)
			(_nw368).ArraySet1((_this).F9(), 16)
			(_nw368).ArraySet1((_this).F9(), 17)
			(_nw368).ArraySet1((_this).F9(), 18)
			(_nw368).ArraySet1((_this).F9(), 19)
			_2249_v39 = _nw368
			var _2250_v40 _dafny.MultiSet
			_ = _2250_v40
			_2250_v40 = _dafny.MultiSetOf(_2249_v39, _2249_v39, _2249_v39)
			var _2251_v41 _dafny.Map
			_ = _2251_v41
			_2251_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2237_v27, _2237_v27)
			var _2252_v42 D3
			_ = _2252_v42
			_2252_v42 = Companion_D3_.Create_DC8_(_2251_v41)
			var _2253_v43 _dafny.MultiSet
			_ = _2253_v43
			_2253_v43 = _dafny.MultiSetOf(_2235_v25)
			var _2254_v44 *C6
			_ = _2254_v44
			var _nw369 *C6 = New_C6_()
			_ = _nw369
			_nw369.Ctor__(_2237_v27, (func() _dafny.Sequence {
				if (_2244_v34).Contains((_this).F9()) {
					return (_2244_v34).Get((_this).F9()).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(706))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_2255_v29 _dafny.Sequence, _2256_v27 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_2257_i1 _dafny.Int) _dafny.CodePoint {
						return (_2255_v29).Select((Companion_Default___.SafeIndex(_2256_v27, _dafny.IntOfUint32((_2255_v29).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_2238_v29, _2237_v27)))
			})())
			_2254_v44 = _nw369
			var _2258_v45 _dafny.Map
			_ = _2258_v45
			_2258_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2254_v44)
			var _2259_v46 _dafny.Array
			_ = _2259_v46
			var _nwElement0_76 _dafny.Int = (_dafny.Zero).Minus((_2246_v36).Cardinality())
			_ = _nwElement0_76
			var _nw370 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(23))
			_ = _nw370
			(_nw370).ArraySet1(_nwElement0_76, 0)
			(_nw370).ArraySet1(((_2248_v38).Dtor_cf39()).Minus(_2237_v27), 1)
			(_nw370).ArraySet1(_2237_v27, 2)
			(_nw370).ArraySet1(((_2250_v40).Update(_2249_v39, Companion_Default___.Abs(_2237_v27))).Cardinality(), 3)
			(_nw370).ArraySet1((func() _dafny.Int {
				if (_this).F9() {
					return _dafny.IntOfUint32((_2235_v25).Cardinality())
				}
				return Companion_Default___.Fm6(_2237_v27, globalState)
			})(), 4)
			(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt((_2251_v41).Cardinality(), _2237_v27), 5)
			(_nw370).ArraySet1(_2237_v27, 6)
			(_nw370).ArraySet1(_dafny.IntOfInt64(-926), 7)
			(_nw370).ArraySet1(_2237_v27, 8)
			(_nw370).ArraySet1((func() _dafny.Int {
				if (_2243_v33).Contains((_this).F9()) {
					return (_2243_v33).Multiplicity((_this).F9())
				}
				return _dafny.IntOfInt64(835)
			})(), 9)
			(_nw370).ArraySet1(Companion_Default___.SafeModuloInt(_2237_v27, _2237_v27), 10)
			(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm26(_2252_v42, _2237_v27, _2253_v43, globalState), _2237_v27), 11)
			(_nw370).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((((_2258_v45).Update((_this).F9(), _2254_v44)).Update((Companion_D9_.Create_DC21_((_this).F9(), _2235_v25, true)).Dtor_cf26(), _2254_v44)).Cardinality())), 12)
			(_nw370).ArraySet1(_2237_v27, 13)
			(_nw370).ArraySet1((_2254_v44).F10(), 14)
			(_nw370).ArraySet1((_dafny.Zero).Minus(_2237_v27), 15)
			(_nw370).ArraySet1((_dafny.Zero).Minus(_2237_v27), 16)
			(_nw370).ArraySet1((_2237_v27).Minus(_dafny.IntOfInt64(52)), 17)
			(_nw370).ArraySet1(_2237_v27, 18)
			(_nw370).ArraySet1(_2237_v27, 19)
			(_nw370).ArraySet1((_dafny.Zero).Minus(_2237_v27), 20)
			(_nw370).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_2254_v44)).Cardinality()), 21)
			(_nw370).ArraySet1(_2237_v27, 22)
			_2259_v46 = _nw370
			var _2260_v47 _dafny.Sequence
			_ = _2260_v47
			_2260_v47 = _dafny.SeqOf(_2259_v46, _2259_v46, _2259_v46, _2259_v46)
			var _2261_v48 _dafny.Sequence
			_ = _2261_v48
			_2261_v48 = _dafny.SeqOf(_2259_v46, _2259_v46, _2259_v46, (_2260_v47).Select((Companion_Default___.SafeIndex((_2254_v44).F10(), _dafny.IntOfUint32((_2260_v47).Cardinality()))).Uint32()).(_dafny.Array))
			var _rhs383 D4 = Companion_D4_.Create_DC11_((_this).F9(), true, _dafny.IntOfInt64(-501), _dafny.Companion_Sequence_.Concatenate((_2254_v44).F11(), (_2254_v44).F11()))
			_ = _rhs383
			var _rhs384 _dafny.Array = (_2260_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2235_v25).Cardinality()), _dafny.IntOfUint32((_2260_v47).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs384
			var _rhs385 bool = (_this).F9()
			_ = _rhs385
			_2245_v35 = _rhs383
			_2259_v46 = _rhs384
			r0 = _rhs385
			var _2262_v49 *C4
			_ = _2262_v49
			var _nw371 *C4 = New_C4_()
			_ = _nw371
			_nw371.Ctor__((_2242_v32).F12())
			_2262_v49 = _nw371
			r1 = Companion_Default___.Fm20((_this).F9(), globalState)
		} else if _source38.Is_DC0() {
			var _2263___mcc_h2 _dafny.Int = _source38.Get_().(D0_DC0).Cf0
			_ = _2263___mcc_h2
			var _2264_cf0 _dafny.Int = _2263___mcc_h2
			_ = _2264_cf0
			var _2265_v50 _dafny.Map
			_ = _2265_v50
			_2265_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_2264_cf0).Times(_dafny.IntOfInt64(944)))
			var _2266_v51 _dafny.Sequence
			_ = _2266_v51
			_2266_v51 = _dafny.SeqOf(_2264_cf0)
			_2265_v50 = (_2265_v50).Update((_this).F9(), (_2264_cf0).Plus(_dafny.IntOfUint32((_2266_v51).Cardinality())))
			(globalState).F8 = _2264_cf0
			(globalState).F4 = (_2264_cf0).Cmp((_dafny.Zero).Minus(_2264_cf0)) <= 0
			var _2267_v52 _dafny.Array
			_ = _2267_v52
			var _len0_61 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_61
			var _nw372 _dafny.Array
			_ = _nw372
			if _len0_61.Cmp(_dafny.Zero) == 0 {
				_nw372 = _dafny.NewArray(_len0_61)
			} else {
				var _init61 func(_dafny.Int) bool = func(_2268_i2 _dafny.Int) bool {
					return (_this).F9()
				}
				_ = _init61
				var _element0_61 = _init61(_dafny.Zero)
				_ = _element0_61
				_nw372 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
				(_nw372).ArraySet1(_element0_61, 0)
				var _nativeLen0_61 = (_len0_61).Int()
				_ = _nativeLen0_61
				for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
					(_nw372).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
				}
			}
			_2267_v52 = _nw372
			var _2269_v53 _dafny.Sequence
			_ = _2269_v53
			_2269_v53 = _dafny.UnicodeSeqOfUtf8Bytes("jwsx")
			var _2270_v54 _dafny.Map
			_ = _2270_v54
			_2270_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2269_v53).Cardinality()), _2267_v52)
			var _2271_v55 _dafny.MultiSet
			_ = _2271_v55
			_2271_v55 = _dafny.MultiSetOf(_2264_cf0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-984), _dafny.IntOfInt64(-528))).Cardinality()), _2264_cf0)
			var _2272_v56 _dafny.Array
			_ = _2272_v56
			var _nwElement0_77 _dafny.Array = _2267_v52
			_ = _nwElement0_77
			var _nw373 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(14))
			_ = _nw373
			(_nw373).ArraySet1(_nwElement0_77, 0)
			(_nw373).ArraySet1(_2267_v52, 1)
			(_nw373).ArraySet1((func() _dafny.Array {
				if (_2270_v54).Contains((_2271_v55).Cardinality()) {
					return (_2270_v54).Get((_2271_v55).Cardinality()).(_dafny.Array)
				}
				return _2267_v52
			})(), 2)
			(_nw373).ArraySet1(_2267_v52, 3)
			(_nw373).ArraySet1(_2267_v52, 4)
			(_nw373).ArraySet1(_2267_v52, 5)
			(_nw373).ArraySet1(_2267_v52, 6)
			(_nw373).ArraySet1(_2267_v52, 7)
			(_nw373).ArraySet1(_2267_v52, 8)
			(_nw373).ArraySet1(_2267_v52, 9)
			(_nw373).ArraySet1((func() _dafny.Array {
				if true {
					return _2267_v52
				}
				return _2267_v52
			})(), 10)
			(_nw373).ArraySet1((Companion_D2_.Create_DC5_(_2267_v52)).Dtor_cf5(), 11)
			(_nw373).ArraySet1(_2267_v52, 12)
			(_nw373).ArraySet1(_2267_v52, 13)
			_2272_v56 = _nw373
			var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_2272_v56), 0))
			_ = _index394
			(_2272_v56).ArraySet1(_2267_v52, (_index394).Int())
			var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_2272_v56), 0))
			_ = _index395
			(_2272_v56).ArraySet1(_2267_v52, (_index395).Int())
		} else {
			var _2273___mcc_h3 D0 = _source38.Get_().(D0_DC3).Cf3
			_ = _2273___mcc_h3
			var _2274_cf3 D0 = _2273___mcc_h3
			_ = _2274_cf3
			if (_this).F9() {
				var _2275_v57 _dafny.Array
				_ = _2275_v57
				var _nw374 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw374
				_2275_v57 = _nw374
				_2275_v57 = _2275_v57
				var _2276_v58 _dafny.Sequence
				_ = _2276_v58
				_2276_v58 = _dafny.UnicodeSeqOfUtf8Bytes("ljehk")
				var _2277_v59 _dafny.Sequence
				_ = _2277_v59
				_2277_v59 = _dafny.SeqOf(_2276_v58, _dafny.UnicodeSeqOfUtf8Bytes("pbuyasjcd"))
				var _2278_v60 _dafny.Sequence
				_ = _2278_v60
				_2278_v60 = _dafny.SeqOf(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_2276_v58, _2276_v58, _2276_v58, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}(func(_2279_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))), _2277_v59))
				_2278_v60 = _2278_v60
				var _2280_v61 *C5
				_ = _2280_v61
				var _nw375 *C5 = New_C5_()
				_ = _nw375
				_nw375.Ctor__()
				_2280_v61 = _nw375
				_2276_v58 = _dafny.UnicodeSeqOfUtf8Bytes("mrlflpdvu")
				var _2281_v62 _dafny.CodePoint
				_ = _2281_v62
				_2281_v62 = _dafny.CodePoint('m')
				var _2282_v63 _dafny.Map
				_ = _2282_v63
				_2282_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2281_v62)
				var _2283_v64 _dafny.Map
				_ = _2283_v64
				_2283_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2282_v63, _2276_v58)
				_2283_v64 = (_2283_v64).Update((_2282_v63).Merge(_2282_v63), _2276_v58)
			} else {
				var _2284_v65 _dafny.Sequence
				_ = _2284_v65
				_2284_v65 = _dafny.UnicodeSeqOfUtf8Bytes("kypi")
				var _2285_v66 _dafny.Map
				_ = _2285_v66
				_2285_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2284_v65).Cardinality()), false)
				var _2286_v67 _dafny.Int
				_ = _2286_v67
				_2286_v67 = _dafny.IntOfInt64(831)
				_2285_v66 = (_2285_v66).Update(Companion_Default___.Fm6(_2286_v67, globalState), false)
				(globalState).F3 = (_dafny.Zero).Minus(_2286_v67)
				(globalState).F3 = _2286_v67
				var _2287_v68 _dafny.Array
				_ = _2287_v68
				var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
				_ = _nw376
				_2287_v68 = _nw376
				var _2288_v69 _dafny.Sequence
				_ = _2288_v69
				_2288_v69 = _dafny.SeqOf((_this).F9())
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_2287_v68), 0))
				_ = _index396
				(_2287_v68).ArraySet1(_dafny.MultiSetFromSeq(_2288_v69), (_index396).Int())
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_2287_v68), 0))
				_ = _index397
				(_2287_v68).ArraySet1((_dafny.MultiSetOf((_this).F9(), (_this).F9(), !(!((_this).F9())))).Update((_this).F9(), Companion_Default___.Abs(_2286_v67)), (_index397).Int())
				var _2289_v70 _dafny.Set
				_ = _2289_v70
				_2289_v70 = _dafny.SetOf(!((_this).F9()))
				r2 = _dafny.SetOf(_2286_v67, Companion_Default___.SafeModuloInt(_2286_v67, _2286_v67), ((_dafny.SetOf(true, (_this).F9(), (_this).F9(), (_this).F9(), (_this).F9())).Intersection(_2289_v70)).Cardinality(), (_dafny.Zero).Minus((_2286_v67).Times(_dafny.IntOfInt64(293))), _2286_v67)
			}
			var _2290_v71 _dafny.Sequence
			_ = _2290_v71
			_2290_v71 = _dafny.SeqOf(_dafny.IntOfInt64(390))
			var _2291_v72 _dafny.Int
			_ = _2291_v72
			_2291_v72 = _dafny.IntOfInt64(603)
			var _2292_v73 _dafny.Set
			_ = _2292_v73
			_2292_v73 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_2290_v71, _2290_v71), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg100 _dafny.Int) interface{} {
					return coer100(arg100)
				}
			}(func(_2293_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-895)
			})), (Companion_Default___.SafeIndex(_2291_v72, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}(func(_2294_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-895)
			}))).Cardinality()))).Uint32(), _dafny.IntOfInt64(206)), _2290_v71)
			_2292_v73 = _2292_v73
			var _2295_v74 _dafny.CodePoint
			_ = _2295_v74
			_2295_v74 = _dafny.CodePoint('f')
			var _2296_v75 _dafny.Map
			_ = _2296_v75
			_2296_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2295_v74, _2295_v74)
			var _2297_v76 _dafny.Array
			_ = _2297_v76
			var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw377
			_2297_v76 = _nw377
			var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_2297_v76), 0))
			_ = _index398
			(_2297_v76).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2297_v76, (_this).F9())).Cardinality()), (_index398).Int())
			var _2298_v77 _dafny.Array
			_ = _2298_v77
			var _len0_62 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_62
			var _nw378 _dafny.Array
			_ = _nw378
			if _len0_62.Cmp(_dafny.Zero) == 0 {
				_nw378 = _dafny.NewArray(_len0_62)
			} else {
				var _init62 func(_dafny.Int) bool = func(_2299_i5 _dafny.Int) bool {
					return (_this).F9()
				}
				_ = _init62
				var _element0_62 = _init62(_dafny.Zero)
				_ = _element0_62
				_nw378 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
				(_nw378).ArraySet1(_element0_62, 0)
				var _nativeLen0_62 = (_len0_62).Int()
				_ = _nativeLen0_62
				for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
					(_nw378).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
				}
			}
			_2298_v77 = _nw378
			var _2300_v78 _dafny.Map
			_ = _2300_v78
			_2300_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2298_v77, _2296_v75)
			var _2301_v79 _dafny.Map
			_ = _2301_v79
			_2301_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2291_v72, _2291_v72)
			var _2302_v80 _dafny.Map
			_ = _2302_v80
			_2302_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2301_v79).Cardinality(), true)
			var _2303_v81 _dafny.MultiSet
			_ = _2303_v81
			_2303_v81 = _dafny.MultiSetOf((_this).F9())
			var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_2297_v76), 0))
			_ = _index399
			var _rhs386 _dafny.Map = (func() _dafny.Map {
				if (_2300_v78).Contains(_2298_v77) {
					return (_2300_v78).Get(_2298_v77).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2295_v74, Companion_Default___.Fm18((func() bool {
					if (_2302_v80).Contains(_2291_v72) {
						return (_2302_v80).Get(_2291_v72).(bool)
					}
					return (_this).F9()
				})(), (_this).F9(), globalState))
			})()
			_ = _rhs386
			var _rhs387 _dafny.Int = (func() _dafny.Int {
				if (_this).F9() {
					return (_2291_v72).Times((_2303_v81).Cardinality())
				}
				return (_2290_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.IntOfUint32((_2290_v71).Cardinality()))).Uint32()).(_dafny.Int)
			})()
			_ = _rhs387
			var _lhs302 _dafny.Array = _2297_v76
			_ = _lhs302
			var _lhs303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_2297_v76), 0))
			_ = _lhs303
			_2296_v75 = _rhs386
			(_lhs302).ArraySet1(_rhs387, (_lhs303).Int())
			var _2304_v82 _dafny.Sequence
			_ = _2304_v82
			_2304_v82 = _dafny.UnicodeSeqOfUtf8Bytes("ktuiokx")
			_2304_v82 = Companion_Default___.Fm25(globalState)
		}
		var _2305_v83 _dafny.Array
		_ = _2305_v83
		var _nw379 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw379
		_2305_v83 = _nw379
		var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))
		_ = _index400
		(_2305_v83).ArraySet1(Companion_Default___.Fm20((_this).F9(), globalState), (_index400).Int())
		var _2306_v84 _dafny.Int
		_ = _2306_v84
		_2306_v84 = _dafny.IntOfInt64(382)
		var _2307_v85 _dafny.Sequence
		_ = _2307_v85
		_2307_v85 = _dafny.SeqOf((_this).F9())
		var _2308_v86 _dafny.Sequence
		_ = _2308_v86
		_2308_v86 = _dafny.SeqOf(_2306_v84, _2306_v84, _2306_v84, _dafny.IntOfUint32((_2307_v85).Cardinality()))
		var _2309_v87 _dafny.Map
		_ = _2309_v87
		_2309_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
		var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))
		_ = _index401
		(_2305_v83).ArraySet1(Companion_Default___.Fm6(Companion_Default___.SafeModuloInt(_2306_v84, (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_2308_v86, (Companion_Default___.SafeIndex(Companion_Default___.Fm6((_2309_v87).Cardinality(), globalState), _dafny.IntOfUint32((_2308_v86).Cardinality()))).Uint32(), _2306_v84))).Cardinality()), globalState), (_index401).Int())
		var _2310_v88 _dafny.Sequence
		_ = _2310_v88
		_2310_v88 = _dafny.UnicodeSeqOfUtf8Bytes("ojpl")
		var _2311_v89 D4
		_ = _2311_v89
		_2311_v89 = Companion_D4_.Create_DC11_((_this).F9(), (_this).F9(), _dafny.IntOfUint32((_dafny.SeqOf(_2306_v84, (_2305_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))).Int()).(_dafny.Int))).Cardinality()), _2310_v88)
		var _hi14 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg102 _dafny.Int) interface{} {
				return coer102(arg102)
			}
		}(func(_2312_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))).Cardinality())
		_ = _hi14
		for _2313_i6 := (_2311_v89).Dtor_cf16(); _2313_i6.Cmp(_hi14) < 0; _2313_i6 = _2313_i6.Plus(_dafny.One) {
			var _2314_v90 D4
			_ = _2314_v90
			_2314_v90 = Companion_D4_.Create_DC10_(_dafny.SeqOf((_this).F9()))
			var _2315_v91 _dafny.Map
			_ = _2315_v91
			_2315_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm49((_this).F9(), globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F9()), (_2314_v90).Dtor_cf13()))
			r1 = (_2315_v91).Cardinality()
			r0 = (_this).F9()
			(globalState).F8 = _2313_i6
			_2310_v88 = _2310_v88
		}
		var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))
		_ = _index402
		var _rhs388 _dafny.Int = ((_2305_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))).Int()).(_dafny.Int)).Plus((_2305_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))).Int()).(_dafny.Int))
		_ = _rhs388
		var _rhs389 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2307_v85, _dafny.Companion_Sequence_.Update(_2307_v85, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.IntOfUint32((_2307_v85).Cardinality()))).Uint32(), (_this).F9())), _dafny.SeqOf((_this).F9()))).Cardinality())
		_ = _rhs389
		var _lhs304 *GlobalState = globalState
		_ = _lhs304
		var _lhs305 _dafny.Array = _2305_v83
		_ = _lhs305
		var _lhs306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))
		_ = _lhs306
		_lhs304.F8 = _rhs388
		(_lhs305).ArraySet1(_rhs389, (_lhs306).Int())
		var _2316_v92 T0
		_ = _2316_v92
		var _nw380 *C7 = New_C7_()
		_ = _nw380
		_nw380.Ctor__()
		_2316_v92 = _nw380
		var _2317_v93 _dafny.Sequence
		_ = _2317_v93
		_2317_v93 = _dafny.SeqOf(_2316_v92)
		var _2318_v94 _dafny.Sequence
		_ = _2318_v94
		_2318_v94 = _2317_v93
		_2318_v94 = _2317_v93
		(globalState).F3 = (_2305_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))).Int()).(_dafny.Int)
		var _2319_v95 D2
		_ = _2319_v95
		_2319_v95 = Companion_D2_.Create_DC6_((_2305_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_2305_v83), 0))).Int()).(_dafny.Int), (_this).F9(), (_this).F9())
		var _pat_let_tv71 = _2305_v83
		_ = _pat_let_tv71
		var _pat_let_tv72 = _2305_v83
		_ = _pat_let_tv72
		r0 = func(_source39 D2) bool {
			if _source39.Is_DC6() {
				var _2320___mcc_h4 _dafny.Int = _source39.Get_().(D2_DC6).Cf6
				_ = _2320___mcc_h4
				var _2321___mcc_h5 bool = _source39.Get_().(D2_DC6).Cf7
				_ = _2321___mcc_h5
				var _2322___mcc_h6 bool = _source39.Get_().(D2_DC6).Cf8
				_ = _2322___mcc_h6
				var _2323_cf8 bool = _2322___mcc_h6
				_ = _2323_cf8
				var _2324_cf7 bool = _2321___mcc_h5
				_ = _2324_cf7
				var _2325_cf6 _dafny.Int = _2320___mcc_h4
				_ = _2325_cf6
				if (_this).F9() {
					return _2323_cf8
				} else {
					return !(_2324_cf7)
				}
			} else if _source39.Is_DC7() {
				var _2326___mcc_h7 _dafny.CodePoint = _source39.Get_().(D2_DC7).Cf9
				_ = _2326___mcc_h7
				var _2327___mcc_h8 bool = _source39.Get_().(D2_DC7).Cf10
				_ = _2327___mcc_h8
				var _2328___mcc_h9 _dafny.Int = _source39.Get_().(D2_DC7).Cf11
				_ = _2328___mcc_h9
				var _2329_cf11 _dafny.Int = _2328___mcc_h9
				_ = _2329_cf11
				var _2330_cf10 bool = _2327___mcc_h8
				_ = _2330_cf10
				var _2331_cf9 _dafny.CodePoint = _2326___mcc_h7
				_ = _2331_cf9
				return (_this).F9()
			} else {
				var _2332___mcc_h10 _dafny.Array = _source39.Get_().(D2_DC5).Cf5
				_ = _2332___mcc_h10
				var _2333_cf5 _dafny.Array = _2332___mcc_h10
				_ = _2333_cf5
				return (_dafny.MultiSetOf(_dafny.IntOfInt64(529))).IsDisjointFrom(_dafny.MultiSetOf((_pat_let_tv72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_pat_let_tv71), 0))).Int()).(_dafny.Int)))
			}
		}(_2319_v95)
		r1 = _2306_v84
		var _2334_v96 _dafny.Set
		_ = _2334_v96
		_2334_v96 = _dafny.SetOf(_2306_v84)
		r2 = _2334_v96
		return r0, r1, r2
	}
}
func (_this *C9) M3(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _2335_v0 _dafny.Sequence
		_ = _2335_v0
		_2335_v0 = _dafny.SeqOf((_this).F9())
		var _2336_v1 _dafny.Sequence
		_ = _2336_v1
		_2336_v1 = _dafny.SeqOf(_2335_v0, _dafny.SeqOf(false), _2335_v0, _2335_v0, _2335_v0)
		var _2337_v2 _dafny.Map
		_ = _2337_v2
		_2337_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2335_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32(((_2336_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2336_v1).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
		var _2338_v3 _dafny.Sequence
		_ = _2338_v3
		_2338_v3 = _dafny.UnicodeSeqOfUtf8Bytes("julgoho")
		var _2339_v4 _dafny.CodePoint
		_ = _2339_v4
		_2339_v4 = _dafny.CodePoint('o')
		var _2340_v5 _dafny.MultiSet
		_ = _2340_v5
		_2340_v5 = _dafny.MultiSetOf(((func() _dafny.Int {
			if (_2337_v2).Contains(_2335_v0) {
				return (_2337_v2).Get(_2335_v0).(_dafny.Int)
			}
			return p0
		})()).Plus(p0), _dafny.IntOfInt64(-539), (func() _dafny.Int {
			if (_this).F9() {
				return _dafny.IntOfUint32((_2338_v3).Cardinality())
			}
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm17(_2339_v4, globalState)).Cardinality()))
		})(), (_dafny.Zero).Minus(p0), p0)
		(globalState).F3 = (func() _dafny.Int {
			if (_2340_v5).Contains(p0) {
				return (_2340_v5).Multiplicity(p0)
			}
			return p0
		})()
		var _2341_v6 _dafny.Set
		_ = _2341_v6
		_2341_v6 = _dafny.SetOf(p0)
		if !((_this).F9()) || ((_2341_v6).IsDisjointFrom(_2341_v6)) {
			var _2342_v7 _dafny.Sequence
			_ = _2342_v7
			_2342_v7 = _dafny.SeqOf(_dafny.IntOfInt64(989), p0)
			var _2343_v8 _dafny.Map
			_ = _2343_v8
			_2343_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(841))
			var _2344_v9 D3
			_ = _2344_v9
			_2344_v9 = Companion_D3_.Create_DC8_(_2343_v8)
			var _2345_v10 _dafny.MultiSet
			_ = _2345_v10
			_2345_v10 = _dafny.MultiSetOf(_2335_v0)
			(globalState).F0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2342_v7, _2342_v7), (Companion_Default___.SafeIndex((p1).Times(p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2342_v7, _2342_v7)).Cardinality()))).Uint32(), Companion_Default___.Fm26(_2344_v9, _dafny.IntOfInt64(687), _2345_v10, globalState))
			var _2346_v11 _dafny.Array
			_ = _2346_v11
			var _nw381 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw381
			_2346_v11 = _nw381
			var _2347_v12 _dafny.Set
			_ = _2347_v12
			_2347_v12 = _dafny.SetOf(_2346_v11, _2346_v11)
			var _2348_v13 _dafny.Array
			_ = _2348_v13
			var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw382
			_2348_v13 = _nw382
			var _2349_v14 T0
			_ = _2349_v14
			var _nw383 *C4 = New_C4_()
			_ = _nw383
			_nw383.Ctor__(_2348_v13)
			_2349_v14 = _nw383
			var _2350_v15 _dafny.Sequence
			_ = _2350_v15
			_2350_v15 = _dafny.SeqOf(_2346_v11)
			var _rhs390 bool = (_this).F9()
			_ = _rhs390
			var _rhs391 bool = (_this).F9()
			_ = _rhs391
			var _rhs392 _dafny.Set = _dafny.SetOf(_2346_v11, (func() _dafny.Array {
				if (_this).F9() {
					return _2346_v11
				}
				return _2346_v11
			})(), (_2350_v15).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2350_v15).Cardinality()))).Uint32()).(_dafny.Array))
			_ = _rhs392
			var _rhs393 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-777))
			_ = _rhs393
			var _rhs394 T0 = _2349_v14
			_ = _rhs394
			var _lhs307 *GlobalState = globalState
			_ = _lhs307
			var _lhs308 *GlobalState = globalState
			_ = _lhs308
			r3 = _rhs390
			_lhs307.F4 = _rhs391
			_2347_v12 = _rhs392
			_lhs308.F3 = _rhs393
			_2349_v14 = _rhs394
			if ((p0).Plus(p0)).Cmp(p0) < 0 {
				(globalState).F3 = Companion_Default___.SafeModuloInt(p1, (Companion_Default___.Fm6(Companion_Default___.Fm20((_this).F9(), globalState), globalState)).Minus(p0))
				var _2351_v16 _dafny.Set
				_ = _2351_v16
				_2351_v16 = _dafny.SetOf(_2339_v4, Companion_Default___.Fm18((_this).F9(), (_this).F9(), globalState), Companion_Default___.Fm18((_this).F9(), false, globalState), _2339_v4)
				_2351_v16 = (_2351_v16).Intersection((_dafny.SetOf(_2339_v4)).Difference(_dafny.SetOf((Companion_D16_.Create_DC35_(p0, _2339_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-613))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}((func(_2352_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2353_i0 _dafny.Int) _dafny.CodePoint {
						return _2352_v4
					}
				})(_2339_v4))), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-613))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg104 _dafny.Int) interface{} {
						return coer104(arg104)
					}
				}((func(_2354_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2355_i0 _dafny.Int) _dafny.CodePoint {
						return _2354_v4
					}
				})(_2339_v4)))).Cardinality()))).Uint32(), _2339_v4)).Cardinality()), p0, p1)).Dtor_cf51())))
				(globalState).F4 = (func() bool {
					if (_this).F9() {
						return (_this).F9()
					}
					return (_this).F9()
				})()
				var _2356_v17 _dafny.Array
				_ = _2356_v17
				var _len0_63 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_63
				var _nw384 _dafny.Array
				_ = _nw384
				if _len0_63.Cmp(_dafny.Zero) == 0 {
					_nw384 = _dafny.NewArray(_len0_63)
				} else {
					var _init63 func(_dafny.Int) bool = func(_2357_i1 _dafny.Int) bool {
						return (_this).F9()
					}
					_ = _init63
					var _element0_63 = _init63(_dafny.Zero)
					_ = _element0_63
					_nw384 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
					(_nw384).ArraySet1(_element0_63, 0)
					var _nativeLen0_63 = (_len0_63).Int()
					_ = _nativeLen0_63
					for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
						(_nw384).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
					}
				}
				_2356_v17 = _nw384
				var _2358_v18 _dafny.Sequence
				_ = _2358_v18
				_2358_v18 = _dafny.SeqOf(_2356_v17, _2356_v17, _2356_v17)
				_2358_v18 = (func() _dafny.Sequence {
					if (_this).F9() {
						return _2358_v18
					}
					return _2358_v18
				})()
				_2335_v0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2335_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2335_v0).Cardinality()))).Uint32(), (_this).F9()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2335_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2335_v0).Cardinality()))).Uint32(), (_this).F9())).Cardinality()))).Uint32(), (_this).F9()), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(Companion_Default___.Fm20((_this).F9(), globalState), p0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2335_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2335_v0).Cardinality()))).Uint32(), (_this).F9()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2335_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2335_v0).Cardinality()))).Uint32(), (_this).F9())).Cardinality()))).Uint32(), (_this).F9())).Cardinality()))).Uint32(), (_this).F9())
			} else {
				var _2359_v19 _dafny.Array
				_ = _2359_v19
				var _nw385 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
				_ = _nw385
				_2359_v19 = _nw385
				var _2360_v20 D21
				_ = _2360_v20
				_2360_v20 = Companion_D21_.Create_DC43_(_2342_v7, (_this).F9())
				var _2361_v21 _dafny.Array
				_ = _2361_v21
				var _nwElement0_78 bool = (_this).F9()
				_ = _nwElement0_78
				var _nw386 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(21))
				_ = _nw386
				(_nw386).ArraySet1(_nwElement0_78, 0)
				(_nw386).ArraySet1((_2349_v14).Fm3((_2360_v20).Dtor_cf63(), _2340_v5, (_this).F9(), globalState), 1)
				(_nw386).ArraySet1(false, 2)
				(_nw386).ArraySet1((_this).F9(), 3)
				(_nw386).ArraySet1((_this).F9(), 4)
				(_nw386).ArraySet1((_this).F9(), 5)
				(_nw386).ArraySet1((_this).F9(), 6)
				(_nw386).ArraySet1((_this).F9(), 7)
				(_nw386).ArraySet1((_this).F9(), 8)
				(_nw386).ArraySet1((_this).F9(), 9)
				(_nw386).ArraySet1(false, 10)
				(_nw386).ArraySet1((_this).F9(), 11)
				(_nw386).ArraySet1((_this).F9(), 12)
				(_nw386).ArraySet1((_this).F9(), 13)
				(_nw386).ArraySet1((_this).F9(), 14)
				(_nw386).ArraySet1((_this).F9(), 15)
				(_nw386).ArraySet1((_this).F9(), 16)
				(_nw386).ArraySet1((_this).F9(), 17)
				(_nw386).ArraySet1((_this).F9(), 18)
				(_nw386).ArraySet1((_this).F9(), 19)
				(_nw386).ArraySet1((_this).F9(), 20)
				_2361_v21 = _nw386
				var _2362_v22 _dafny.MultiSet
				_ = _2362_v22
				_2362_v22 = _dafny.MultiSetOf(_2361_v21)
				var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_2359_v19), 0))
				_ = _index403
				(_2359_v19).ArraySet1(_2362_v22, (_index403).Int())
				var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_2359_v19), 0))
				_ = _index404
				(_2359_v19).ArraySet1(_dafny.MultiSetOf(_2361_v21, _2361_v21, _2361_v21), (_index404).Int())
				var _2363_v23 _dafny.Map
				_ = _2363_v23
				_2363_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2335_v0, _2342_v7)
				var _2364_v24 _dafny.Sequence
				_ = _2364_v24
				_2364_v24 = _dafny.SeqOf(_2363_v23)
				_2363_v23 = (_2364_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-356), _dafny.IntOfUint32((_2364_v24).Cardinality()))).Uint32()).(_dafny.Map)
				var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_2361_v21), 0))
				_ = _index405
				(_2361_v21).ArraySet1(Companion_Default___.Fm49(Companion_Default___.Fm49((_this).F9(), globalState), globalState), (_index405).Int())
				var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_2361_v21), 0))
				_ = _index406
				(_2361_v21).ArraySet1((func() bool {
					if false {
						return (_this).F9()
					}
					return (_this).F9()
				})(), (_index406).Int())
				(globalState).F3 = p1
				(globalState).F4 = (_this).F9()
			}
			var _2365_v25 _dafny.Set
			_ = _2365_v25
			_2365_v25 = _dafny.SetOf((_this).F9(), (_this).F9())
			var _2366_v26 _dafny.Map
			_ = _2366_v26
			_2366_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2365_v25).IsDisjointFrom(_2365_v25))
			_2366_v26 = (_2366_v26).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2342_v7, _2346_v11)).Cardinality(), (_this).F9())
			var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_2346_v11), 0))
			_ = _index407
			(_2346_v11).ArraySet1(p0, (_index407).Int())
			var _2367_v27 _dafny.Sequence
			_ = _2367_v27
			_2367_v27 = _dafny.SeqOf(_2338_v3)
			var _2368_v28 D10
			_ = _2368_v28
			_2368_v28 = Companion_D10_.Create_DC24_(p1, !((_this).F9()), (_2367_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_2369_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2370_i2 _dafny.Int) _dafny.Sequence {
					return _2369_v3
				}
			})(_2338_v3)))).Cardinality()), _dafny.IntOfUint32((_2367_v27).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_2346_v11), 0))
			_ = _index408
			(_2346_v11).ArraySet1((_2368_v28).Dtor_cf31(), (_index408).Int())
		} else {
			_2338_v3 = _2338_v3
			var _2371_v29 _dafny.Array
			_ = _2371_v29
			var _nw387 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw387
			_2371_v29 = _nw387
			var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
			_ = _index409
			(_2371_v29).ArraySet1((_this).F9(), (_index409).Int())
			var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
			_ = _index410
			(_2371_v29).ArraySet1(!(_2341_v6).Equals(_dafny.SetOf(p0, p1, p1)), (_index410).Int())
			if (_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool) {
				var _2372_v30 _dafny.Map
				_ = _2372_v30
				_2372_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)))
				(globalState).F8 = (p1).Times((_dafny.Zero).Minus((_2372_v30).Cardinality()))
				(globalState).F4 = (_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)
				var _2373_v31 _dafny.Map
				_ = _2373_v31
				_2373_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.Concatenate(_2338_v3, _2338_v3))
				_2373_v31 = (_2373_v31).Update(((func() _dafny.MultiSet {
					if (_this).F9() {
						return _2340_v5
					}
					return _2340_v5
				})()).Cardinality(), _2338_v3)
				var _2374_v32 _dafny.Array
				_ = _2374_v32
				var _nw388 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw388
				_2374_v32 = _nw388
				var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_2374_v32), 0))
				_ = _index411
				(_2374_v32).ArraySet1((p0).Times(p1), (_index411).Int())
				var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_2374_v32), 0))
				_ = _index412
				(_2374_v32).ArraySet1(p0, (_index412).Int())
				var _2375_v33 *C6
				_ = _2375_v33
				var _nw389 *C6 = New_C6_()
				_ = _nw389
				_nw389.Ctor__(p1, _2338_v3)
				_2375_v33 = _nw389
				var _2376_v34 _dafny.Map
				_ = _2376_v34
				_2376_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2375_v33, !((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)))
				_2376_v34 = (_2376_v34).Update(_2375_v33, (_this).F9())
			} else {
				var _2377_v35 _dafny.Map
				_ = _2377_v35
				_2377_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _2378_v36 _dafny.Sequence
				_ = _2378_v36
				_2378_v36 = _dafny.SeqOf((_dafny.Zero).Minus(p1), p0)
				var _2379_v37 D21
				_ = _2379_v37
				_2379_v37 = Companion_D21_.Create_DC43_(_2378_v36, (_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool))
				var _2380_v38 _dafny.Sequence
				_ = _2380_v38
				_2380_v38 = _2378_v36
				var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
				_ = _index413
				var _rhs395 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1))
				_ = _rhs395
				var _rhs396 bool = (_this).F9()
				_ = _rhs396
				var _rhs397 _dafny.Array = _2371_v29
				_ = _rhs397
				var _rhs398 D21 = Companion_D21_.Create_DC43_(_dafny.Companion_Sequence_.Update(_2378_v36, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2378_v36).Cardinality()))).Uint32(), p0), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg106 _dafny.Int) interface{} {
						return coer106(arg106)
					}
				}((func(_2381_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2382_i3 _dafny.Int) _dafny.Int {
						return _2381_p0
					}
				})(p0))), _2380_v38, _2380_v38), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg107 _dafny.Int) interface{} {
						return coer107(arg107)
					}
				}((func(_2383_v38 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2384_i4 _dafny.Int) _dafny.Sequence {
						return _2383_v38
					}
				})(_2380_v38)))))
				_ = _rhs398
				var _rhs399 bool = (_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)
				_ = _rhs399
				var _lhs309 *GlobalState = globalState
				_ = _lhs309
				var _lhs310 _dafny.Array = _2371_v29
				_ = _lhs310
				var _lhs311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
				_ = _lhs311
				_2377_v35 = _rhs395
				_lhs309.F4 = _rhs396
				_2371_v29 = _rhs397
				_2379_v37 = _rhs398
				(_lhs310).ArraySet1(_rhs399, (_lhs311).Int())
				var _2385_v39 _dafny.Array
				_ = _2385_v39
				var _nw390 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
				_ = _nw390
				_2385_v39 = _nw390
				var _2386_v40 _dafny.Map
				_ = _2386_v40
				_2386_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool), p1)
				var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_2385_v39), 0))
				_ = _index414
				(_2385_v39).ArraySet1(_2386_v40, (_index414).Int())
				var _2387_v41 _dafny.Array
				_ = _2387_v41
				var _nw391 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
				_ = _nw391
				_2387_v41 = _nw391
				var _2388_v42 _dafny.MultiSet
				_ = _2388_v42
				_2388_v42 = _dafny.MultiSetOf((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool), (_this).F9())
				var _2389_v43 _dafny.Map
				_ = _2389_v43
				_2389_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool), (_this).F9())
				var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_2385_v39), 0))
				_ = _index415
				var _rhs400 bool = ((_2388_v42).Intersection(_2388_v42)).IsProperSubsetOf(_2388_v42)
				_ = _rhs400
				var _rhs401 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2378_v36, _dafny.Companion_Sequence_.Update(_2378_v36, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2378_v36).Cardinality()))).Uint32(), p1)), (Companion_Default___.SafeIndex((_2389_v43).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2378_v36, _dafny.Companion_Sequence_.Update(_2378_v36, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2378_v36).Cardinality()))).Uint32(), p1))).Cardinality()))).Uint32(), p1)
				_ = _rhs401
				var _rhs402 _dafny.Map = (_2386_v40).Merge((_2386_v40).Merge(_2386_v40))
				_ = _rhs402
				var _rhs403 _dafny.Array = _2387_v41
				_ = _rhs403
				var _lhs312 *GlobalState = globalState
				_ = _lhs312
				var _lhs313 *GlobalState = globalState
				_ = _lhs313
				var _lhs314 _dafny.Array = _2385_v39
				_ = _lhs314
				var _lhs315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_2385_v39), 0))
				_ = _lhs315
				_lhs312.F4 = _rhs400
				_lhs313.F0 = _rhs401
				(_lhs314).ArraySet1(_rhs402, (_lhs315).Int())
				_2387_v41 = _rhs403
				var _2390_v45 _dafny.Array
				_ = _2390_v45
				var _len0_64 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_64
				var _nw392 _dafny.Array
				_ = _nw392
				if _len0_64.Cmp(_dafny.Zero) == 0 {
					_nw392 = _dafny.NewArray(_len0_64)
				} else {
					var _init64 func(_dafny.Int) _dafny.Map = (func(_2391_v6 _dafny.Set, _2392_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_2393_i5 _dafny.Int) _dafny.Map {
							return func() _dafny.Map {
								var _coll104 = _dafny.NewMapBuilder()
								_ = _coll104
								for _iter108 := _dafny.Iterate((_2391_v6).Elements()); ; {
									_compr_104, _ok108 := _iter108()
									if !_ok108 {
										break
									}
									var _2394_v44 _dafny.Int
									_2394_v44 = interface{}(_compr_104).(_dafny.Int)
									if (_2391_v6).Contains(_2394_v44) {
										_coll104.Add((_2394_v44).Minus(_2392_p0), _2392_p0)
									}
								}
								return _coll104.ToMap()
							}()
						}
					})(_2341_v6, p0)
					_ = _init64
					var _element0_64 = _init64(_dafny.Zero)
					_ = _element0_64
					_nw392 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
					(_nw392).ArraySet1(_element0_64, 0)
					var _nativeLen0_64 = (_len0_64).Int()
					_ = _nativeLen0_64
					for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
						(_nw392).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
					}
				}
				_2390_v45 = _nw392
				var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_2390_v45), 0))
				_ = _index416
				(_2390_v45).ArraySet1(_2377_v35, (_index416).Int())
				var _2395_v46 D9
				_ = _2395_v46
				_2395_v46 = Companion_D9_.Create_DC21_((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool), _dafny.SeqOf((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)), (_this).F9())
				var _2396_v47 _dafny.Set
				_ = _2396_v47
				_2396_v47 = _dafny.SetOf(Companion_Default___.Fm30(p1, _2395_v46, (_this).F9(), p0, globalState))
				var _2397_v50 D3
				_ = _2397_v50
				_2397_v50 = Companion_D3_.Create_DC8_(func() _dafny.Map {
					var _coll105 = _dafny.NewMapBuilder()
					_ = _coll105
					for _iter109 := _dafny.Iterate((func() _dafny.Set {
						var _coll106 = _dafny.NewBuilder()
						_ = _coll106
						for _iter110 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(924), _dafny.IntOfInt64(161))); ; {
							_compr_106, _ok110 := _iter110()
							if !_ok110 {
								break
							}
							var _2398_v49 _dafny.Int
							_2398_v49 = interface{}(_compr_106).(_dafny.Int)
							if ((_dafny.IntOfInt64(924)).Cmp(_2398_v49) <= 0) && ((_2398_v49).Cmp(_dafny.IntOfInt64(161)) < 0) {
								_coll106.Add(Companion_Default___.SafeModuloInt(_2398_v49, _dafny.IntOfUint32((_2378_v36).Cardinality())))
							}
						}
						return _coll106.ToSet()
					}()).Elements()); ; {
						_compr_105, _ok109 := _iter109()
						if !_ok109 {
							break
						}
						var _2399_v48 _dafny.Int
						_2399_v48 = interface{}(_compr_105).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll107 = _dafny.NewBuilder()
							_ = _coll107
							for _iter111 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(924), _dafny.IntOfInt64(161))); ; {
								_compr_107, _ok111 := _iter111()
								if !_ok111 {
									break
								}
								var _2400_v49 _dafny.Int
								_2400_v49 = interface{}(_compr_107).(_dafny.Int)
								if ((_dafny.IntOfInt64(924)).Cmp(_2400_v49) <= 0) && ((_2400_v49).Cmp(_dafny.IntOfInt64(161)) < 0) {
									_coll107.Add(Companion_Default___.SafeModuloInt(_2400_v49, _dafny.IntOfUint32((_2378_v36).Cardinality())))
								}
							}
							return _coll107.ToSet()
						}()).Contains(_2399_v48) {
							_coll105.Add(Companion_Default___.SafeModuloInt(_2399_v48, p1), p1)
						}
					}
					return _coll105.ToMap()
				}())
				var _2401_v51 _dafny.MultiSet
				_ = _2401_v51
				_2401_v51 = _dafny.MultiSetOf(_dafny.SeqOf(!((_this).F9())))
				var _2402_v52 D24
				_ = _2402_v52
				_2402_v52 = Companion_D24_.Create_DC48_(_2396_v47)
				var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_2390_v45), 0))
				_ = _index417
				var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
				_ = _index418
				var _rhs404 _dafny.Map = _2377_v35
				_ = _rhs404
				var _rhs405 _dafny.Int = Companion_Default___.Fm26(_2397_v50, Companion_Default___.SafeDivisionInt(p1, p1), (_2401_v51).Difference(_2401_v51), globalState)
				_ = _rhs405
				var _rhs406 bool = Companion_Default___.Fm49((_this).F9(), globalState)
				_ = _rhs406
				var _rhs407 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p1))
				_ = _rhs407
				var _rhs408 _dafny.Set = ((_2396_v47).Difference((_2402_v52).Dtor_cf71())).Union((_2396_v47).Union(_2396_v47))
				_ = _rhs408
				var _lhs316 _dafny.Array = _2390_v45
				_ = _lhs316
				var _lhs317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_2390_v45), 0))
				_ = _lhs317
				var _lhs318 *GlobalState = globalState
				_ = _lhs318
				var _lhs319 _dafny.Array = _2371_v29
				_ = _lhs319
				var _lhs320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))
				_ = _lhs320
				var _lhs321 *GlobalState = globalState
				_ = _lhs321
				(_lhs316).ArraySet1(_rhs404, (_lhs317).Int())
				_lhs318.F3 = _rhs405
				(_lhs319).ArraySet1(_rhs406, (_lhs320).Int())
				_lhs321.F3 = _rhs407
				_2396_v47 = _rhs408
				(globalState).F4 = (_2335_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2335_v0).Cardinality()))).Uint32()).(bool)
				var _rhs409 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dhkfosxv")).Cardinality()), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9())).Cardinality()))).Plus(p1))
				_ = _rhs409
				var _rhs410 _dafny.Int = _dafny.IntOfInt64(-629)
				_ = _rhs410
				var _lhs322 *GlobalState = globalState
				_ = _lhs322
				_lhs322.F3 = _rhs409
				r2 = _rhs410
			}
			var _2403_v53 _dafny.Map
			_ = _2403_v53
			_2403_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool), _2341_v6)
			var _2404_v54 _dafny.Map
			_ = _2404_v54
			_2404_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _rhs411 bool = (_2371_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_2371_v29), 0))).Int()).(bool)
			_ = _rhs411
			var _rhs412 _dafny.Int = Companion_Default___.SafeModuloInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2403_v53).Cardinality())).Merge(_2404_v54)).Cardinality(), Companion_Default___.Fm20(Companion_Default___.Fm49(true, globalState), globalState))
			_ = _rhs412
			var _lhs323 *GlobalState = globalState
			_ = _lhs323
			r3 = _rhs411
			_lhs323.F3 = _rhs412
			var _2405_v55 _dafny.Map
			_ = _2405_v55
			_2405_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(630), p1)).Dtor_cf2(), _2340_v5)
			var _2406_v56 _dafny.Map
			_ = _2406_v56
			_2406_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2405_v55, (_dafny.Zero).Minus(p1))
			_2406_v56 = (_2406_v56).Update((_2405_v55).Merge(_2405_v55), _dafny.IntOfInt64(652))
		}
		var _2407_v57 _dafny.Array
		_ = _2407_v57
		var _nw393 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw393
		_2407_v57 = _nw393
		var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))
		_ = _index419
		(_2407_v57).ArraySet1((func() bool {
			if (_this).F9() {
				return (_this).F9()
			}
			return (_this).F9()
		})(), (_index419).Int())
		var _index420 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))
		_ = _index420
		(_2407_v57).ArraySet1(false, (_index420).Int())
		var _2408_v58 _dafny.Array
		_ = _2408_v58
		var _nwElement0_79 _dafny.CodePoint = _2339_v4
		_ = _nwElement0_79
		var _nw394 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(4))
		_ = _nw394
		(_nw394).ArraySet1CodePoint(_nwElement0_79, 0)
		(_nw394).ArraySet1CodePoint(_dafny.CodePoint('p'), 1)
		(_nw394).ArraySet1CodePoint(_2339_v4, 2)
		(_nw394).ArraySet1CodePoint(_2339_v4, 3)
		_2408_v58 = _nw394
		var _index421 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_2408_v58), 0))
		_ = _index421
		(_2408_v58).ArraySet1CodePoint(_2339_v4, (_index421).Int())
		var _index422 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_2408_v58), 0))
		_ = _index422
		(_2408_v58).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_2407_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))).Int()).(bool) {
				return _2339_v4
			}
			return (func() _dafny.CodePoint {
				if !((_2407_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))).Int()).(bool)) {
					return _2339_v4
				}
				return _dafny.CodePoint('r')
			})()
		})(), (_index422).Int())
		var _2409_v59 _dafny.Sequence
		_ = _2409_v59
		_2409_v59 = _dafny.SeqOf(_2407_v57)
		var _2410_v60 _dafny.Array
		_ = _2410_v60
		var _nwElement0_80 _dafny.Array = _2407_v57
		_ = _nwElement0_80
		var _nw395 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(22))
		_ = _nw395
		(_nw395).ArraySet1(_nwElement0_80, 0)
		(_nw395).ArraySet1(_2407_v57, 1)
		(_nw395).ArraySet1(_2407_v57, 2)
		(_nw395).ArraySet1(_2407_v57, 3)
		(_nw395).ArraySet1(_2407_v57, 4)
		(_nw395).ArraySet1(_2407_v57, 5)
		(_nw395).ArraySet1(_2407_v57, 6)
		(_nw395).ArraySet1(_2407_v57, 7)
		(_nw395).ArraySet1(_2407_v57, 8)
		(_nw395).ArraySet1((func() _dafny.Array {
			if (_2407_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))).Int()).(bool) {
				return _2407_v57
			}
			return _2407_v57
		})(), 9)
		(_nw395).ArraySet1(_2407_v57, 10)
		(_nw395).ArraySet1(_2407_v57, 11)
		(_nw395).ArraySet1(_2407_v57, 12)
		(_nw395).ArraySet1(_2407_v57, 13)
		(_nw395).ArraySet1(_2407_v57, 14)
		(_nw395).ArraySet1(_2407_v57, 15)
		(_nw395).ArraySet1(_2407_v57, 16)
		(_nw395).ArraySet1(_2407_v57, 17)
		(_nw395).ArraySet1((_2409_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2338_v3).Cardinality()), _dafny.IntOfUint32((_2409_v59).Cardinality()))).Uint32()).(_dafny.Array), 18)
		(_nw395).ArraySet1(_2407_v57, 19)
		(_nw395).ArraySet1(_2407_v57, 20)
		(_nw395).ArraySet1(_2407_v57, 21)
		_2410_v60 = _nw395
		_2410_v60 = _2410_v60
		var _2411_v61 _dafny.Array
		_ = _2411_v61
		var _nw396 _dafny.Array = _dafny.NewArrayWithValue(Companion_D19_.Default(), _dafny.IntOfInt64(10))
		_ = _nw396
		_2411_v61 = _nw396
		var _2412_v62 _dafny.Set
		_ = _2412_v62
		_2412_v62 = _dafny.SetOf(_2411_v61)
		var _2413_v63 _dafny.Map
		_ = _2413_v63
		_2413_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_2408_v58).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_2408_v58), 0))).Int()))
		var _2414_v64 _dafny.Array
		_ = _2414_v64
		var _nwElement0_81 _dafny.Int = p0
		_ = _nwElement0_81
		var _nw397 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(20))
		_ = _nw397
		(_nw397).ArraySet1(_nwElement0_81, 0)
		(_nw397).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), (_this).F9())).Cardinality())), 1)
		(_nw397).ArraySet1(p0, 2)
		(_nw397).ArraySet1(p0, 3)
		(_nw397).ArraySet1(_dafny.IntOfInt64(889), 4)
		(_nw397).ArraySet1(p1, 5)
		(_nw397).ArraySet1(p0, 6)
		(_nw397).ArraySet1(p1, 7)
		(_nw397).ArraySet1(Companion_Default___.Fm20((_2407_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_2407_v57), 0))).Int()).(bool), globalState), 8)
		(_nw397).ArraySet1(p0, 9)
		(_nw397).ArraySet1(p1, 10)
		(_nw397).ArraySet1((_dafny.Zero).Minus(p0), 11)
		(_nw397).ArraySet1(p1, 12)
		(_nw397).ArraySet1((_dafny.Zero).Minus(p0), 13)
		(_nw397).ArraySet1(p1, 14)
		(_nw397).ArraySet1((_dafny.Zero).Minus(p1), 15)
		(_nw397).ArraySet1(p1, 16)
		(_nw397).ArraySet1((_2413_v63).Cardinality(), 17)
		(_nw397).ArraySet1(_dafny.IntOfInt64(240), 18)
		(_nw397).ArraySet1(p0, 19)
		_2414_v64 = _nw397
		var _2415_v65 _dafny.Map
		_ = _2415_v65
		_2415_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2412_v62, _2414_v64)
		r0 = (func() _dafny.Array {
			if (_2415_v65).Contains((_2412_v62).Difference(_2412_v62)) {
				return (_2415_v65).Get((_2412_v62).Difference(_2412_v62)).(_dafny.Array)
			}
			return _2414_v64
		})()
		r0 = _2414_v64
		r1 = p0
		r2 = p1
		r3 = false
		return r0, r1, r2, r3
	}
}
func (_this *C9) F9() bool {
	{
		return _this._f9
	}
}

// End of class C9

// Definition of class C10
type C10 struct {
	dummy byte
}

func New_C10_() *C10 {
	_this := C10{}

	return &_this
}

type CompanionStruct_C10_ struct {
}

var Companion_C10_ = CompanionStruct_C10_{}

func (_this *C10) Equals(other *C10) bool {
	return _this == other
}

func (_this *C10) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C10)
	return ok && _this.Equals(other)
}

func (*C10) String() string {
	return "_module.C10"
}

func Type_C10_() _dafny.TypeDescriptor {
	return type_C10_{}
}

type type_C10_ struct {
}

func (_this type_C10_) Default() interface{} {
	return (*C10)(nil)
}

func (_this type_C10_) String() string {
	return "main.C10"
}
func (_this *C10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C10{}
var _ _dafny.TraitOffspring = &C10{}

func (_this *C10) Ctor__() {
	{
	}
}
func (_this *C10) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(645), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(663))).Cardinality()), (_dafny.SetOf((_dafny.MultiSetOf(true, true)).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(81), false)).Cardinality(), (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jsdxerq"))).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-368), _dafny.SetOf((_dafny.MultiSetOf(false, false)).Cardinality()))).Cardinality()), (func() _dafny.Map {
			var _coll108 = _dafny.NewMapBuilder()
			_ = _coll108
			for _iter112 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-85))).Elements()); ; {
				_compr_108, _ok112 := _iter112()
				if !_ok112 {
					break
				}
				var _2416_v0 _dafny.Int
				_2416_v0 = interface{}(_compr_108).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(-85))).Contains(_2416_v0) {
					_coll108.Add((_2416_v0).Minus(_dafny.IntOfInt64(898)), _dafny.CodePoint('t'))
				}
			}
			return _coll108.ToMap()
		}()).Cardinality())).Cardinality())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('r'))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg108 _dafny.Int) interface{} {
				return coer108(arg108)
			}
		}(func(_2417_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(23)
		}))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-29), (_dafny.SetOf(_dafny.IntOfInt64(491))).Cardinality()))
	}
}
func (_this *C10) Fm3(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C10) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (func() _dafny.Map {
			var _coll109 = _dafny.NewMapBuilder()
			_ = _coll109
			for _iter113 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC0_((_dafny.SetOf(true, true)).Cardinality()))).Elements()); ; {
				_compr_109, _ok113 := _iter113()
				if !_ok113 {
					break
				}
				var _2418_v0 D0
				_2418_v0 = interface{}(_compr_109).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC0_((_dafny.SetOf(true, true)).Cardinality())), _2418_v0) {
					_coll109.Add(_2418_v0, true)
				}
			}
			return _coll109.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll110 = _dafny.NewMapBuilder()
			_ = _coll110
			for _iter114 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-672)), _dafny.CodePoint('r'))).Keys().Elements()); ; {
				_compr_110, _ok114 := _iter114()
				if !_ok114 {
					break
				}
				var _2419_v1 D0
				_2419_v1 = interface{}(_compr_110).(D0)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-672)), _dafny.CodePoint('r'))).Contains(_2419_v1) {
					_coll110.Add(_2419_v1, false)
				}
			}
			return _coll110.ToMap()
		}())
	}
}
func (_this *C10) M0(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _2420_v0 *C0
		_ = _2420_v0
		var _nw398 *C0 = New_C0_()
		_ = _nw398
		_nw398.Ctor__()
		_2420_v0 = _nw398
		var _2421_v1 bool
		_ = _2421_v1
		_2421_v1 = true
		var _2422_v2 _dafny.MultiSet
		_ = _2422_v2
		_2422_v2 = _dafny.MultiSetOf(_2421_v1)
		_2422_v2 = _2422_v2
		var _2423_v3 _dafny.Int
		_ = _2423_v3
		_2423_v3 = _dafny.IntOfInt64(428)
		var _hi15 _dafny.Int = (_2423_v3).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(_2423_v3)))
		_ = _hi15
		for _2424_i0 := (_2423_v3).Plus(_2423_v3); _2424_i0.Cmp(_hi15) < 0; _2424_i0 = _2424_i0.Plus(_dafny.One) {
			_2423_v3 = _dafny.IntOfInt64(-567)
			var _2425_v4 _dafny.Sequence
			_ = _2425_v4
			_2425_v4 = _dafny.SeqOf(_2421_v1, _2421_v1)
			var _2426_v6 _dafny.Map
			_ = _2426_v6
			_2426_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_2424_i0)), _2421_v1)
			var _2427_v7 _dafny.Set
			_ = _2427_v7
			_2427_v7 = _dafny.SetOf((func() _dafny.Map {
				var _coll111 = _dafny.NewMapBuilder()
				_ = _coll111
				for _iter115 := _dafny.Iterate((_2426_v6).Keys().Elements()); ; {
					_compr_111, _ok115 := _iter115()
					if !_ok115 {
						break
					}
					var _2428_v5 _dafny.Int
					_2428_v5 = interface{}(_compr_111).(_dafny.Int)
					if (_2426_v6).Contains(_2428_v5) {
						_coll111.Add((_2428_v5).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(319))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg109 _dafny.Int) interface{} {
								return coer109(arg109)
							}
						}(func(_2429_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						}))).Cardinality())), _2421_v1)
					}
				}
				return _coll111.ToMap()
			}()).Cardinality())
			var _2430_v8 D13
			_ = _2430_v8
			_2430_v8 = Companion_D13_.Create_DC30_(_2421_v1, (_2425_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.IntOfUint32((_2425_v4).Cardinality()))).Uint32()).(bool), _2427_v7, _2420_v0)
			var _2431_v9 _dafny.Array
			_ = _2431_v9
			var _nwElement0_82 D13 = _2430_v8
			_ = _nwElement0_82
			var _nw399 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(3))
			_ = _nw399
			(_nw399).ArraySet1(_nwElement0_82, 0)
			(_nw399).ArraySet1(_2430_v8, 1)
			(_nw399).ArraySet1(_2430_v8, 2)
			_2431_v9 = _nw399
			var _index423 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_2431_v9), 0))
			_ = _index423
			(_2431_v9).ArraySet1(_2430_v8, (_index423).Int())
			var _2432_v10 _dafny.Array
			_ = _2432_v10
			var _nw400 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw400
			_2432_v10 = _nw400
			var _index424 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_2431_v9), 0))
			_ = _index424
			var _rhs413 D13 = _2430_v8
			_ = _rhs413
			var _rhs414 _dafny.Int = _2423_v3
			_ = _rhs414
			var _rhs415 _dafny.Array = (func() _dafny.Array {
				if _dafny.Companion_Sequence_.Equal(_2425_v4, _2425_v4) {
					return _2432_v10
				}
				return _2432_v10
			})()
			_ = _rhs415
			var _rhs416 _dafny.Int = _2423_v3
			_ = _rhs416
			var _lhs324 _dafny.Array = _2431_v9
			_ = _lhs324
			var _lhs325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_2431_v9), 0))
			_ = _lhs325
			var _lhs326 *GlobalState = globalState
			_ = _lhs326
			var _lhs327 *GlobalState = globalState
			_ = _lhs327
			(_lhs324).ArraySet1(_rhs413, (_lhs325).Int())
			_lhs326.F8 = _rhs414
			_2432_v10 = _rhs415
			_lhs327.F8 = _rhs416
			var _2433_v11 _dafny.Array
			_ = _2433_v11
			var _nw401 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
			_ = _nw401
			_2433_v11 = _nw401
			var _2434_v12 *C1
			_ = _2434_v12
			var _nw402 *C1 = New_C1_()
			_ = _nw402
			_nw402.Ctor__()
			_2434_v12 = _nw402
			var _index425 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2433_v11), 0))
			_ = _index425
			(_2433_v11).ArraySet1(_2434_v12, (_index425).Int())
			var _index426 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2433_v11), 0))
			_ = _index426
			var _rhs417 _dafny.Int = _2424_i0
			_ = _rhs417
			var _rhs418 *C1 = _2434_v12
			_ = _rhs418
			var _lhs328 *GlobalState = globalState
			_ = _lhs328
			var _lhs329 _dafny.Array = _2433_v11
			_ = _lhs329
			var _lhs330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2433_v11), 0))
			_ = _lhs330
			_lhs328.F3 = _rhs417
			(_lhs329).ArraySet1(_rhs418, (_lhs330).Int())
			var _2435_v13 _dafny.Sequence
			_ = _2435_v13
			_2435_v13 = _dafny.UnicodeSeqOfUtf8Bytes("qajy")
			var _2436_v14 _dafny.Sequence
			_ = _2436_v14
			_2436_v14 = _dafny.SeqOf(_dafny.IntOfInt64(63))
			var _2437_v15 _dafny.Map
			_ = _2437_v15
			_2437_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2435_v13).Cardinality()), _2424_i0), (_2436_v14).Select((Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2436_v14).Cardinality()))).Uint32()).(_dafny.Int))
			_2437_v15 = (_2437_v15).Update(_2423_v3, _2424_i0)
		}
		if ((func() _dafny.MultiSet {
			if _2421_v1 {
				return _2422_v2
			}
			return _2422_v2
		})()).IsProperSubsetOf(_2422_v2) {
			(globalState).F8 = _2423_v3
			if _2421_v1 {
				var _2438_v16 _dafny.Array
				_ = _2438_v16
				var _len0_65 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_65
				var _nw403 _dafny.Array
				_ = _nw403
				if _len0_65.Cmp(_dafny.Zero) == 0 {
					_nw403 = _dafny.NewArray(_len0_65)
				} else {
					var _init65 func(_dafny.Int) _dafny.Map = (func(_2439_v1 bool) func(_dafny.Int) _dafny.Map {
						return func(_2440_i2 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2439_v1, _dafny.CodePoint('h'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2439_v1, _dafny.CodePoint('b')))
						}
					})(_2421_v1)
					_ = _init65
					var _element0_65 = _init65(_dafny.Zero)
					_ = _element0_65
					_nw403 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
					(_nw403).ArraySet1(_element0_65, 0)
					var _nativeLen0_65 = (_len0_65).Int()
					_ = _nativeLen0_65
					for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
						(_nw403).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
					}
				}
				_2438_v16 = _nw403
				_2438_v16 = _2438_v16
				var _2441_v17 *C1
				_ = _2441_v17
				var _nw404 *C1 = New_C1_()
				_ = _nw404
				_nw404.Ctor__()
				_2441_v17 = _nw404
				var _2442_v18 _dafny.Map
				_ = _2442_v18
				_2442_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2441_v17, _dafny.IntOfInt64(761))
				(globalState).F3 = (_2423_v3).Times((func() _dafny.Int {
					if (_2442_v18).Contains(_2441_v17) {
						return (_2442_v18).Get(_2441_v17).(_dafny.Int)
					}
					return _2423_v3
				})())
				var _rhs419 bool = _2421_v1
				_ = _rhs419
				var _rhs420 bool = _2421_v1
				_ = _rhs420
				var _lhs331 *GlobalState = globalState
				_ = _lhs331
				r0 = _rhs419
				_lhs331.F4 = _rhs420
				var _2443_v19 _dafny.Array
				_ = _2443_v19
				var _nw405 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
				_ = _nw405
				_2443_v19 = _nw405
				var _2444_v20 T0
				_ = _2444_v20
				var _nw406 *C4 = New_C4_()
				_ = _nw406
				_nw406.Ctor__(_2443_v19)
				_2444_v20 = _nw406
				var _2445_v21 _dafny.Set
				_ = _2445_v21
				_2445_v21 = _dafny.SetOf(_2444_v20)
				var _2446_v22 _dafny.Map
				_ = _2446_v22
				_2446_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(898), _2445_v21)
				var _2447_v23 _dafny.Sequence
				_ = _2447_v23
				_2447_v23 = _dafny.SeqOf(_2445_v21)
				var _2448_v24 _dafny.Array
				_ = _2448_v24
				var _nwElement0_83 _dafny.Set = (_2445_v21).Difference(_dafny.SetOf(_2444_v20, _2444_v20, _2444_v20))
				_ = _nwElement0_83
				var _nw407 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(22))
				_ = _nw407
				(_nw407).ArraySet1(_nwElement0_83, 0)
				(_nw407).ArraySet1(_2445_v21, 1)
				(_nw407).ArraySet1((_2445_v21).Difference(_2445_v21), 2)
				(_nw407).ArraySet1((_2445_v21).Intersection(_2445_v21), 3)
				(_nw407).ArraySet1(_2445_v21, 4)
				(_nw407).ArraySet1((_2445_v21).Union(_2445_v21), 5)
				(_nw407).ArraySet1((_2445_v21).Difference(_2445_v21), 6)
				(_nw407).ArraySet1(_2445_v21, 7)
				(_nw407).ArraySet1((func() _dafny.Set {
					if (_2446_v22).Contains(_2423_v3) {
						return (_2446_v22).Get(_2423_v3).(_dafny.Set)
					}
					return _2445_v21
				})(), 8)
				(_nw407).ArraySet1((_2445_v21).Difference(_dafny.SetOf(_2444_v20, _2444_v20, _2444_v20)), 9)
				(_nw407).ArraySet1(_2445_v21, 10)
				(_nw407).ArraySet1((_2447_v23).Select((Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2447_v23).Cardinality()))).Uint32()).(_dafny.Set), 11)
				(_nw407).ArraySet1(_2445_v21, 12)
				(_nw407).ArraySet1(_2445_v21, 13)
				(_nw407).ArraySet1(_2445_v21, 14)
				(_nw407).ArraySet1(_2445_v21, 15)
				(_nw407).ArraySet1(_dafny.SetOf(_2444_v20), 16)
				(_nw407).ArraySet1(_2445_v21, 17)
				(_nw407).ArraySet1(_2445_v21, 18)
				(_nw407).ArraySet1(_2445_v21, 19)
				(_nw407).ArraySet1(_dafny.SetOf(_2444_v20), 20)
				(_nw407).ArraySet1((_2445_v21).Difference(_dafny.SetOf(_2444_v20)), 21)
				_2448_v24 = _nw407
				var _index427 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_2448_v24), 0))
				_ = _index427
				(_2448_v24).ArraySet1(_2445_v21, (_index427).Int())
				var _index428 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_2448_v24), 0))
				_ = _index428
				(_2448_v24).ArraySet1(_dafny.SetOf(_2444_v20), (_index428).Int())
				var _2449_v25 _dafny.CodePoint
				_ = _2449_v25
				_2449_v25 = _dafny.CodePoint('h')
				_2449_v25 = _2449_v25
			} else {
				var _2450_v26 _dafny.Array
				_ = _2450_v26
				var _nw408 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw408
				_2450_v26 = _nw408
				var _2451_v27 _dafny.Sequence
				_ = _2451_v27
				_2451_v27 = _dafny.UnicodeSeqOfUtf8Bytes("sdej")
				var _2452_v28 _dafny.CodePoint
				_ = _2452_v28
				_2452_v28 = _dafny.CodePoint('w')
				var _2453_v29 D10
				_ = _2453_v29
				_2453_v29 = Companion_D10_.Create_DC24_(_2423_v3, _2421_v1, _2451_v27)
				var _2454_v30 _dafny.Array
				_ = _2454_v30
				var _nwElement0_84 _dafny.Sequence = _2451_v27
				_ = _nwElement0_84
				var _nw409 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(28))
				_ = _nw409
				(_nw409).ArraySet1(_nwElement0_84, 0)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uvvaguomr"), 1)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vv"), 2)
				(_nw409).ArraySet1(_2451_v27, 3)
				(_nw409).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}(func(_2455_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}(func(_2456_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				}))).Cardinality()))).Uint32(), _2452_v28), 4)
				(_nw409).ArraySet1(_2451_v27, 5)
				(_nw409).ArraySet1(_2451_v27, 6)
				(_nw409).ArraySet1(_2451_v27, 7)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("r"), 8)
				(_nw409).ArraySet1(Companion_Default___.Fm25(globalState), 9)
				(_nw409).ArraySet1(_2451_v27, 10)
				(_nw409).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(319))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_2457_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2458_i4 _dafny.Int) _dafny.CodePoint {
						return _2457_v28
					}
				})(_2452_v28))), 11)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gwsellrim"), 12)
				(_nw409).ArraySet1(_dafny.Companion_Sequence_.Update(_2451_v27, (Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2451_v27).Cardinality()))).Uint32(), _2452_v28), 13)
				(_nw409).ArraySet1(_2451_v27, 14)
				(_nw409).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-182))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_2459_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2460_i5 _dafny.Int) _dafny.CodePoint {
						return _2459_v28
					}
				})(_2452_v28))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-182))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}((func(_2461_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2462_i5 _dafny.Int) _dafny.CodePoint {
						return _2461_v28
					}
				})(_2452_v28)))).Cardinality()))).Uint32(), _2452_v28), 15)
				(_nw409).ArraySet1(_2451_v27, 16)
				(_nw409).ArraySet1(_2451_v27, 17)
				(_nw409).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}((func(_2463_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2464_i6 _dafny.Int) _dafny.CodePoint {
						return _2463_v28
					}
				})(_2452_v28))), 18)
				(_nw409).ArraySet1((_2453_v29).Dtor_cf33(), 19)
				(_nw409).ArraySet1(_2451_v27, 20)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("p"), 21)
				(_nw409).ArraySet1(_2451_v27, 22)
				(_nw409).ArraySet1(_2451_v27, 23)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fuxk"), 24)
				(_nw409).ArraySet1(_2451_v27, 25)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lmxp"), 26)
				(_nw409).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gsdb"), 27)
				_2454_v30 = _nw409
				var _index429 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2450_v26), 0))
				_ = _index429
				(_2450_v26).ArraySet1(_2454_v30, (_index429).Int())
				var _index430 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2450_v26), 0))
				_ = _index430
				var _nw410 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw410
				(_2450_v26).ArraySet1(_nw410, (_index430).Int())
				var _2465_v31 _dafny.Map
				_ = _2465_v31
				_2465_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, true)
				var _2466_v32 _dafny.Sequence
				_ = _2466_v32
				_2466_v32 = _dafny.SeqOf(_2451_v27, _dafny.UnicodeSeqOfUtf8Bytes("tgmejkpo"), _2451_v27, _2451_v27, _2451_v27)
				var _2467_v33 _dafny.Set
				_ = _2467_v33
				_2467_v33 = _dafny.SetOf(_2421_v1, _2421_v1)
				var _2468_v34 _dafny.Array
				_ = _2468_v34
				var _nwElement0_85 _dafny.Int = (_2422_v2).Cardinality()
				_ = _nwElement0_85
				var _nw411 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(26))
				_ = _nw411
				(_nw411).ArraySet1(_nwElement0_85, 0)
				(_nw411).ArraySet1(_2423_v3, 1)
				(_nw411).ArraySet1((_dafny.Zero).Minus((_2465_v31).Cardinality()), 2)
				(_nw411).ArraySet1(_2423_v3, 3)
				(_nw411).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_2466_v32).Cardinality())), 4)
				(_nw411).ArraySet1(_2423_v3, 5)
				(_nw411).ArraySet1(_2423_v3, 6)
				(_nw411).ArraySet1(_2423_v3, 7)
				(_nw411).ArraySet1(_2423_v3, 8)
				(_nw411).ArraySet1((_dafny.Zero).Minus(_2423_v3), 9)
				(_nw411).ArraySet1(_2423_v3, 10)
				(_nw411).ArraySet1((_2467_v33).Cardinality(), 11)
				(_nw411).ArraySet1(_2423_v3, 12)
				(_nw411).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer116 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}((func(_2469_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2470_i7 _dafny.Int) _dafny.Int {
						return _2469_v3
					}
				})(_2423_v3)))).Cardinality()), 13)
				(_nw411).ArraySet1(_2423_v3, 14)
				(_nw411).ArraySet1(_2423_v3, 15)
				(_nw411).ArraySet1(_2423_v3, 16)
				(_nw411).ArraySet1(_2423_v3, 17)
				(_nw411).ArraySet1(_2423_v3, 18)
				(_nw411).ArraySet1(_2423_v3, 19)
				(_nw411).ArraySet1(_2423_v3, 20)
				(_nw411).ArraySet1((_2420_v0).Fm19(globalState), 21)
				(_nw411).ArraySet1(_2423_v3, 22)
				(_nw411).ArraySet1(_2423_v3, 23)
				(_nw411).ArraySet1(_2423_v3, 24)
				(_nw411).ArraySet1(_dafny.IntOfUint32((_2451_v27).Cardinality()), 25)
				_2468_v34 = _nw411
				var _2471_v35 _dafny.Array
				_ = _2471_v35
				var _nwElement0_86 _dafny.Array = _2468_v34
				_ = _nwElement0_86
				var _nw412 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(10))
				_ = _nw412
				(_nw412).ArraySet1(_nwElement0_86, 0)
				(_nw412).ArraySet1(_2468_v34, 1)
				(_nw412).ArraySet1(_2468_v34, 2)
				(_nw412).ArraySet1(_2468_v34, 3)
				(_nw412).ArraySet1(_2468_v34, 4)
				(_nw412).ArraySet1(_2468_v34, 5)
				(_nw412).ArraySet1(_2468_v34, 6)
				(_nw412).ArraySet1(_2468_v34, 7)
				(_nw412).ArraySet1(_2468_v34, 8)
				(_nw412).ArraySet1(_2468_v34, 9)
				_2471_v35 = _nw412
				var _index431 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_2471_v35), 0))
				_ = _index431
				(_2471_v35).ArraySet1(_2468_v34, (_index431).Int())
				var _index432 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_2471_v35), 0))
				_ = _index432
				var _rhs421 _dafny.Array = _2468_v34
				_ = _rhs421
				var _rhs422 bool = _2421_v1
				_ = _rhs422
				var _rhs423 bool = (_2423_v3).Cmp(_2423_v3) < 0
				_ = _rhs423
				var _lhs332 _dafny.Array = _2471_v35
				_ = _lhs332
				var _lhs333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_2471_v35), 0))
				_ = _lhs333
				(_lhs332).ArraySet1(_rhs421, (_lhs333).Int())
				r0 = _rhs422
				_2421_v1 = _rhs423
				var _2472_v36 _dafny.Array
				_ = _2472_v36
				var _nw413 _dafny.Array = _dafny.NewArrayWithValue(Companion_D23_.Default(), _dafny.IntOfInt64(12))
				_ = _nw413
				_2472_v36 = _nw413
				var _2473_v37 _dafny.Array
				_ = _2473_v37
				var _nwElement0_87 _dafny.CodePoint = _2452_v28
				_ = _nwElement0_87
				var _nw414 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(4))
				_ = _nw414
				(_nw414).ArraySet1CodePoint(_nwElement0_87, 0)
				(_nw414).ArraySet1CodePoint(_2452_v28, 1)
				(_nw414).ArraySet1CodePoint(_2452_v28, 2)
				(_nw414).ArraySet1CodePoint(_2452_v28, 3)
				_2473_v37 = _nw414
				var _2474_v38 D23
				_ = _2474_v38
				_2474_v38 = Companion_D23_.Create_DC45_(_2473_v37)
				var _index433 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_2472_v36), 0))
				_ = _index433
				(_2472_v36).ArraySet1(_2474_v38, (_index433).Int())
				var _pat_let_tv73 = _2473_v37
				_ = _pat_let_tv73
				var _index434 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_2472_v36), 0))
				_ = _index434
				(_2472_v36).ArraySet1(func(_pat_let63_0 D23) D23 {
					return func(_2475_dt__update__tmp_h0 D23) D23 {
						return func(_pat_let64_0 _dafny.Array) D23 {
							return func(_2476_dt__update_hcf66_h0 _dafny.Array) D23 {
								return Companion_D23_.Create_DC45_(_2476_dt__update_hcf66_h0)
							}(_pat_let64_0)
						}(_pat_let_tv73)
					}(_pat_let63_0)
				}(_2474_v38), (_index434).Int())
				var _2477_v39 _dafny.Map
				_ = _2477_v39
				_2477_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, _2423_v3)
				var _2478_v40 D2
				_ = _2478_v40
				_2478_v40 = Companion_D2_.Create_DC7_(_2452_v28, _2421_v1, _dafny.IntOfInt64(786))
				var _2479_v41 _dafny.Sequence
				_ = _2479_v41
				_2479_v41 = _dafny.SeqOf(_2467_v33)
				_2477_v39 = (_2477_v39).Update((_2478_v40).Dtor_cf11(), ((_2467_v33).Intersection((_2479_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("guqrkvlrc")).Cardinality()), _dafny.IntOfUint32((_2479_v41).Cardinality()))).Uint32()).(_dafny.Set))).Cardinality())
				(globalState).F3 = (_dafny.Zero).Minus(_2423_v3)
			}
			var _2480_v42 _dafny.Array
			_ = _2480_v42
			var _len0_66 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_66
			var _nw415 _dafny.Array
			_ = _nw415
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw415 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) bool = (func(_2481_v1 bool) func(_dafny.Int) bool {
					return func(_2482_i8 _dafny.Int) bool {
						return _2481_v1
					}
				})(_2421_v1)
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw415 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw415).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw415).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2480_v42 = _nw415
			_2480_v42 = _2480_v42
			var _2483_v43 _dafny.Array
			_ = _2483_v43
			var _nw416 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw416
			_2483_v43 = _nw416
			var _index435 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2483_v43), 0))
			_ = _index435
			(_2483_v43).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-180)), (_index435).Int())
			var _2484_v44 _dafny.Map
			_ = _2484_v44
			_2484_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, _2423_v3)
			var _2485_v45 _dafny.Map
			_ = _2485_v45
			_2485_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2484_v44, _2423_v3)
			var _index436 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2483_v43), 0))
			_ = _index436
			(_2483_v43).ArraySet1((func() _dafny.Int {
				if (_2485_v45).Contains(_2484_v44) {
					return (_2485_v45).Get(_2484_v44).(_dafny.Int)
				}
				return _2423_v3
			})(), (_index436).Int())
			var _2486_v46 _dafny.CodePoint
			_ = _2486_v46
			_2486_v46 = _dafny.CodePoint('d')
			var _2487_v47 _dafny.Sequence
			_ = _2487_v47
			_2487_v47 = _dafny.SeqOf(_2486_v46, _2486_v46, _2486_v46)
			var _2488_v48 _dafny.Sequence
			_ = _2488_v48
			_2488_v48 = _dafny.SeqOf(_2486_v46, _2486_v46, _2486_v46, (_2487_v47).Select((Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2487_v47).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_2487_v47 = _2487_v47
		} else {
			if false {
				var _2489_v49 _dafny.Array
				_ = _2489_v49
				var _len0_67 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_67
				var _nw417 _dafny.Array
				_ = _nw417
				if _len0_67.Cmp(_dafny.Zero) == 0 {
					_nw417 = _dafny.NewArray(_len0_67)
				} else {
					var _init67 func(_dafny.Int) _dafny.Int = (func(_2490_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2491_i9 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_2491_i9, _2490_v3)
						}
					})(_2423_v3)
					_ = _init67
					var _element0_67 = _init67(_dafny.Zero)
					_ = _element0_67
					_nw417 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
					(_nw417).ArraySet1(_element0_67, 0)
					var _nativeLen0_67 = (_len0_67).Int()
					_ = _nativeLen0_67
					for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
						(_nw417).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
					}
				}
				_2489_v49 = _nw417
				var _index437 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))
				_ = _index437
				(_2489_v49).ArraySet1(_2423_v3, (_index437).Int())
				var _index438 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))
				_ = _index438
				(_2489_v49).ArraySet1(_2423_v3, (_index438).Int())
				var _2492_v50 _dafny.Array
				_ = _2492_v50
				var _len0_68 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_68
				var _nw418 _dafny.Array
				_ = _nw418
				if _len0_68.Cmp(_dafny.Zero) == 0 {
					_nw418 = _dafny.NewArray(_len0_68)
				} else {
					var _init68 func(_dafny.Int) bool = (func(_2493_v1 bool) func(_dafny.Int) bool {
						return func(_2494_i10 _dafny.Int) bool {
							return _2493_v1
						}
					})(_2421_v1)
					_ = _init68
					var _element0_68 = _init68(_dafny.Zero)
					_ = _element0_68
					_nw418 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
					(_nw418).ArraySet1(_element0_68, 0)
					var _nativeLen0_68 = (_len0_68).Int()
					_ = _nativeLen0_68
					for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
						(_nw418).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
					}
				}
				_2492_v50 = _nw418
				var _index439 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_2492_v50), 0))
				_ = _index439
				(_2492_v50).ArraySet1(_2421_v1, (_index439).Int())
				var _index440 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_2492_v50), 0))
				_ = _index440
				(_2492_v50).ArraySet1(_2421_v1, (_index440).Int())
				var _2495_v51 _dafny.Map
				_ = _2495_v51
				_2495_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), _2421_v1)
				var _2496_v52 _dafny.Sequence
				_ = _2496_v52
				_2496_v52 = _dafny.SeqOf((func() bool {
					if (_2495_v51).Contains(_2423_v3) {
						return (_2495_v51).Get(_2423_v3).(bool)
					}
					return _2421_v1
				})())
				var _2497_v53 _dafny.Map
				_ = _2497_v53
				_2497_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v1, (_2420_v0).Fm19(globalState))
				var _2498_v54 _dafny.MultiSet
				_ = _2498_v54
				_2498_v54 = _dafny.MultiSetOf((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), _2423_v3, (_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), (_2497_v53).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("youkw")).Cardinality()))
				var _2499_v55 _dafny.MultiSet
				_ = _2499_v55
				_2499_v55 = _dafny.MultiSetOf((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), (_2498_v54).Cardinality(), _2423_v3)
				var _2500_v56 _dafny.MultiSet
				_ = _2500_v56
				_2500_v56 = _dafny.MultiSetOf(_2499_v55, _2498_v54)
				var _2501_v57 _dafny.Map
				_ = _2501_v57
				_2501_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), _2500_v56)
				var _2502_v58 _dafny.Sequence
				_ = _2502_v58
				_2502_v58 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}(func(_2503_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})))
				var _2504_v59 _dafny.Sequence
				_ = _2504_v59
				_2504_v59 = _dafny.SeqOf(((func() _dafny.MultiSet {
					if (_2501_v57).Contains((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)) {
						return (_2501_v57).Get((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_dafny.SeqOf(_2498_v54, _2498_v54))
				})()).Cardinality(), (_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), (_dafny.IntOfUint32((_2502_v58).Cardinality())).Plus((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)))
				var _index441 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_2492_v50), 0))
				_ = _index441
				var _index442 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_2492_v50), 0))
				_ = _index442
				var _rhs424 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-617)))
				_ = _rhs424
				var _rhs425 bool = _2421_v1
				_ = _rhs425
				var _rhs426 bool = (_2496_v52).Select((Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2496_v52).Cardinality()))).Uint32()).(bool)
				_ = _rhs426
				var _rhs427 bool = _2421_v1
				_ = _rhs427
				var _rhs428 _dafny.Int = (_2504_v59).Select((Companion_Default___.SafeIndex((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2504_v59).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs428
				var _lhs334 _dafny.Array = _2492_v50
				_ = _lhs334
				var _lhs335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_2492_v50), 0))
				_ = _lhs335
				var _lhs336 _dafny.Array = _2492_v50
				_ = _lhs336
				var _lhs337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_2492_v50), 0))
				_ = _lhs337
				var _lhs338 *GlobalState = globalState
				_ = _lhs338
				_2423_v3 = _rhs424
				(_lhs334).ArraySet1(_rhs425, (_lhs335).Int())
				(_lhs336).ArraySet1(_rhs426, (_lhs337).Int())
				_2421_v1 = _rhs427
				_lhs338.F3 = _rhs428
				var _2505_v60 _dafny.CodePoint
				_ = _2505_v60
				_2505_v60 = _dafny.CodePoint('n')
				var _2506_v61 _dafny.Map
				_ = _2506_v61
				_2506_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v1, _2505_v60)
				var _2507_v62 _dafny.Set
				_ = _2507_v62
				_2507_v62 = _dafny.SetOf(_2423_v3, (_2506_v61).Cardinality())
				var _2508_v63 _dafny.Set
				_ = _2508_v63
				_2508_v63 = _dafny.SetOf((_2507_v62).Cardinality(), (_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int))
				var _2509_v64 _dafny.Set
				_ = _2509_v64
				_2509_v64 = _dafny.SetOf(_2507_v62, (_2508_v63).Union(_dafny.SetOf(_dafny.IntOfUint32((_2496_v52).Cardinality()))), _2507_v62)
				var _2510_v65 _dafny.Array
				_ = _2510_v65
				var _len0_69 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_69
				var _nw419 _dafny.Array
				_ = _nw419
				if _len0_69.Cmp(_dafny.Zero) == 0 {
					_nw419 = _dafny.NewArray(_len0_69)
				} else {
					var _init69 func(_dafny.Int) _dafny.Map = (func(_2511_v51 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_2512_i12 _dafny.Int) _dafny.Map {
							return _2511_v51
						}
					})(_2495_v51)
					_ = _init69
					var _element0_69 = _init69(_dafny.Zero)
					_ = _element0_69
					_nw419 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
					(_nw419).ArraySet1(_element0_69, 0)
					var _nativeLen0_69 = (_len0_69).Int()
					_ = _nativeLen0_69
					for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
						(_nw419).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
					}
				}
				_2510_v65 = _nw419
				var _index443 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2510_v65), 0))
				_ = _index443
				(_2510_v65).ArraySet1(func() _dafny.Map {
					var _coll112 = _dafny.NewMapBuilder()
					_ = _coll112
					for _iter116 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(137), _dafny.IntOfInt64(239))); ; {
						_compr_112, _ok116 := _iter116()
						if !_ok116 {
							break
						}
						var _2513_v66 _dafny.Int
						_2513_v66 = interface{}(_compr_112).(_dafny.Int)
						if ((_dafny.IntOfInt64(137)).Cmp(_2513_v66) <= 0) && ((_2513_v66).Cmp(_dafny.IntOfInt64(239)) < 0) {
							_coll112.Add((_2513_v66).Plus(_dafny.IntOfInt64(520)), _2421_v1)
						}
					}
					return _coll112.ToMap()
				}(), (_index443).Int())
				var _index444 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2510_v65), 0))
				_ = _index444
				var _rhs429 _dafny.Set = _2509_v64
				_ = _rhs429
				var _rhs430 _dafny.Map = (_2495_v51).Merge(_2495_v51)
				_ = _rhs430
				var _rhs431 _dafny.Int = (_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)
				_ = _rhs431
				var _lhs339 _dafny.Array = _2510_v65
				_ = _lhs339
				var _lhs340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2510_v65), 0))
				_ = _lhs340
				var _lhs341 *GlobalState = globalState
				_ = _lhs341
				_2509_v64 = _rhs429
				(_lhs339).ArraySet1(_rhs430, (_lhs340).Int())
				_lhs341.F8 = _rhs431
				var _2514_v67 _dafny.Sequence
				_ = _2514_v67
				_2514_v67 = _dafny.UnicodeSeqOfUtf8Bytes("catkgef")
				var _2515_v68 _dafny.Map
				_ = _2515_v68
				_2515_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int), _2514_v67)
				var _index445 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))
				_ = _index445
				(_2489_v49).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ds"), (func() _dafny.Sequence {
					if (_2515_v68).Contains((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)) {
						return (_2515_v68).Get((_2489_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_2489_v49), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
					}
					return _2514_v67
				})())).Cardinality()), (_index445).Int())
				var _2516_v69 D2
				_ = _2516_v69
				_2516_v69 = Companion_D2_.Create_DC5_(_2492_v50)
				var _2517_v70 _dafny.Array
				_ = _2517_v70
				var _nwElement0_88 _dafny.Array = _2492_v50
				_ = _nwElement0_88
				var _nw420 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(13))
				_ = _nw420
				(_nw420).ArraySet1(_nwElement0_88, 0)
				(_nw420).ArraySet1((func() _dafny.Array {
					if (_2492_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_2492_v50), 0))).Int()).(bool) {
						return _2492_v50
					}
					return _2492_v50
				})(), 1)
				(_nw420).ArraySet1(_2492_v50, 2)
				(_nw420).ArraySet1(_2492_v50, 3)
				(_nw420).ArraySet1(_2492_v50, 4)
				(_nw420).ArraySet1(_2492_v50, 5)
				(_nw420).ArraySet1(_2492_v50, 6)
				(_nw420).ArraySet1(_2492_v50, 7)
				(_nw420).ArraySet1(_2492_v50, 8)
				(_nw420).ArraySet1((_2516_v69).Dtor_cf5(), 9)
				(_nw420).ArraySet1(_2492_v50, 10)
				(_nw420).ArraySet1(_2492_v50, 11)
				(_nw420).ArraySet1(_2492_v50, 12)
				_2517_v70 = _nw420
				var _index446 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_2517_v70), 0))
				_ = _index446
				(_2517_v70).ArraySet1(_2492_v50, (_index446).Int())
				var _index447 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_2517_v70), 0))
				_ = _index447
				(_2517_v70).ArraySet1(_2492_v50, (_index447).Int())
			} else {
				var _2518_v71 _dafny.CodePoint
				_ = _2518_v71
				_2518_v71 = _dafny.CodePoint('t')
				var _2519_v72 _dafny.Sequence
				_ = _2519_v72
				_2519_v72 = _dafny.UnicodeSeqOfUtf8Bytes("gi")
				var _2520_v73 _dafny.Array
				_ = _2520_v73
				var _nwElement0_89 _dafny.CodePoint = _2518_v71
				_ = _nwElement0_89
				var _nw421 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(12))
				_ = _nw421
				(_nw421).ArraySet1CodePoint(_nwElement0_89, 0)
				(_nw421).ArraySet1CodePoint(_2518_v71, 1)
				(_nw421).ArraySet1CodePoint(_2518_v71, 2)
				(_nw421).ArraySet1CodePoint(_2518_v71, 3)
				(_nw421).ArraySet1CodePoint(_2518_v71, 4)
				(_nw421).ArraySet1CodePoint(_2518_v71, 5)
				(_nw421).ArraySet1CodePoint(_2518_v71, 6)
				(_nw421).ArraySet1CodePoint(_dafny.CodePoint('p'), 7)
				(_nw421).ArraySet1CodePoint(_2518_v71, 8)
				(_nw421).ArraySet1CodePoint((_2519_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.IntOfUint32((_2519_v72).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
				(_nw421).ArraySet1CodePoint(_2518_v71, 10)
				(_nw421).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _2421_v1 {
						return _2518_v71
					}
					return _2518_v71
				})(), 11)
				_2520_v73 = _nw421
				var _index448 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2520_v73), 0))
				_ = _index448
				(_2520_v73).ArraySet1CodePoint(_dafny.CodePoint('x'), (_index448).Int())
				var _2521_v74 _dafny.Map
				_ = _2521_v74
				_2521_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, (_2519_v72).Select((Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2519_v72).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _index449 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2520_v73), 0))
				_ = _index449
				(_2520_v73).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_2521_v74).Contains(_2423_v3) {
						return (_2521_v74).Get(_2423_v3).(_dafny.CodePoint)
					}
					return _2518_v71
				})(), (_index449).Int())
				var _2522_v75 _dafny.Array
				_ = _2522_v75
				var _len0_70 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_70
				var _nw422 _dafny.Array
				_ = _nw422
				if _len0_70.Cmp(_dafny.Zero) == 0 {
					_nw422 = _dafny.NewArray(_len0_70)
				} else {
					var _init70 func(_dafny.Int) _dafny.Int = (func(_2523_v3 _dafny.Int, _2524_v2 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
						return func(_2525_i13 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_2525_i13, (_dafny.MultiSetOf((_dafny.Zero).Minus(_2523_v3), _dafny.IntOfInt64(966), (_2524_v2).Cardinality(), _2523_v3)).Cardinality())
						}
					})(_2423_v3, _2422_v2)
					_ = _init70
					var _element0_70 = _init70(_dafny.Zero)
					_ = _element0_70
					_nw422 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
					(_nw422).ArraySet1(_element0_70, 0)
					var _nativeLen0_70 = (_len0_70).Int()
					_ = _nativeLen0_70
					for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
						(_nw422).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
					}
				}
				_2522_v75 = _nw422
				var _index450 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2522_v75), 0))
				_ = _index450
				(_2522_v75).ArraySet1(_2423_v3, (_index450).Int())
				var _index451 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2522_v75), 0))
				_ = _index451
				(_2522_v75).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm6(_2423_v3, globalState), (_dafny.Zero).Minus(_dafny.IntOfUint32((_2519_v72).Cardinality()))), (_index451).Int())
				var _2526_v76 *C3
				_ = _2526_v76
				var _nw423 *C3 = New_C3_()
				_ = _nw423
				_nw423.Ctor__()
				_2526_v76 = _nw423
				var _2527_v77 *C3
				_ = _2527_v77
				_2527_v77 = _2526_v76
				var _2528_v78 _dafny.Sequence
				_ = _2528_v78
				_2528_v78 = _dafny.SeqOf(_2527_v77, _2526_v76, _2527_v77, _2527_v77)
				_2528_v78 = _2528_v78
				(globalState).F3 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(436))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_2529_v73 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_2530_i14 _dafny.Int) _dafny.CodePoint {
						return (_2529_v73).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2529_v73), 0))).Int())
					}
				})(_2520_v73)))).Cardinality())).Minus(_2423_v3)
				(globalState).F4 = _2421_v1
			}
			(globalState).F3 = (_2423_v3).Times((_2423_v3).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, _2421_v1)).Cardinality()))
			var _2531_v79 _dafny.Map
			_ = _2531_v79
			_2531_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm49(_2421_v1, globalState), _2421_v1)
			(globalState).F4 = !(_2531_v79).Contains(_2421_v1)
			var _2532_v80 _dafny.Array
			_ = _2532_v80
			var _nw424 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
			_ = _nw424
			_2532_v80 = _nw424
			var _2533_v81 _dafny.Sequence
			_ = _2533_v81
			_2533_v81 = _dafny.SeqOf(_2421_v1)
			var _2534_v82 _dafny.Sequence
			_ = _2534_v82
			_2534_v82 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_2533_v81).Cardinality())), _2423_v3)
			var _2535_v83 _dafny.Map
			_ = _2535_v83
			_2535_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v1, _2534_v82)
			var _index452 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_2532_v80), 0))
			_ = _index452
			(_2532_v80).ArraySet1((_2535_v83).Update(_2421_v1, _2534_v82), (_index452).Int())
			var _index453 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_2532_v80), 0))
			_ = _index453
			(_2532_v80).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2421_v1) && (_2421_v1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-944))).Uint32(), func(coer119 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg119 _dafny.Int) interface{} {
					return coer119(arg119)
				}
			}((func(_2536_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2537_i15 _dafny.Int) _dafny.Int {
					return _2536_v3
				}
			})(_2423_v3)))), (_index453).Int())
			var _2538_v84 _dafny.CodePoint
			_ = _2538_v84
			_2538_v84 = _dafny.CodePoint('j')
			var _2539_v85 _dafny.Sequence
			_ = _2539_v85
			_2539_v85 = _dafny.UnicodeSeqOfUtf8Bytes("pudx")
			r1 = _dafny.Companion_Sequence_.Contains(_2539_v85, _2538_v84)
		}
		var _2540_v86 _dafny.CodePoint
		_ = _2540_v86
		_2540_v86 = _dafny.CodePoint('b')
		var _2541_v87 _dafny.Map
		_ = _2541_v87
		_2541_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-713), _2540_v86)
		var _2542_v88 D19
		_ = _2542_v88
		_2542_v88 = Companion_D19_.Create_DC38_(_2541_v87)
		var _2543_v89 D19
		_ = _2543_v89
		_2543_v89 = Companion_D19_.Create_DC40_(_2542_v88)
		var _2544_v90 D19
		_ = _2544_v90
		_2544_v90 = Companion_D19_.Create_DC40_(_2542_v88)
		var _source40 D19 = _2544_v90
		_ = _source40
		if _source40.Is_DC39() {
			var _2545___mcc_h0 bool = _source40.Get_().(D19_DC39).Cf58
			_ = _2545___mcc_h0
			var _2546___mcc_h1 _dafny.Int = _source40.Get_().(D19_DC39).Cf59
			_ = _2546___mcc_h1
			var _2547_cf59 _dafny.Int = _2546___mcc_h1
			_ = _2547_cf59
			var _2548_cf58 bool = _2545___mcc_h0
			_ = _2548_cf58
			var _2549_v91 _dafny.Sequence
			_ = _2549_v91
			_2549_v91 = _dafny.UnicodeSeqOfUtf8Bytes("imuchiwj")
			var _2550_v92 _dafny.Array
			_ = _2550_v92
			var _nw425 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw425
			_2550_v92 = _nw425
			var _2551_v93 _dafny.Map
			_ = _2551_v93
			_2551_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, _2550_v92)
			var _2552_v94 _dafny.Map
			_ = _2552_v94
			_2552_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2423_v3, _2423_v3)
			_2423_v3 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2549_v91, (Companion_Default___.SafeIndex((_2551_v93).Cardinality(), _dafny.IntOfUint32((_2549_v91).Cardinality()))).Uint32(), _2540_v86)).Cardinality())).Times(((func() _dafny.Int {
				if (_2552_v94).Contains(Companion_Default___.Fm20(_2548_cf58, globalState)) {
					return (_2552_v94).Get(Companion_Default___.Fm20(_2548_cf58, globalState)).(_dafny.Int)
				}
				return _2423_v3
			})()).Minus(_dafny.IntOfInt64(890)))
			r0 = _2548_cf58
			(globalState).F4 = _2421_v1
			_2550_v92 = _2550_v92
		} else if _source40.Is_DC38() {
			var _2553___mcc_h2 _dafny.Map = _source40.Get_().(D19_DC38).Cf57
			_ = _2553___mcc_h2
			var _2554_cf57 _dafny.Map = _2553___mcc_h2
			_ = _2554_cf57
			var _2555_v95 _dafny.MultiSet
			_ = _2555_v95
			_2555_v95 = _dafny.MultiSetOf((_dafny.Zero).Minus(_2423_v3))
			var _2556_v96 _dafny.MultiSet
			_ = _2556_v96
			_2556_v96 = _2555_v95
			_2556_v96 = _2556_v96
			var _2557_v97 _dafny.Array
			_ = _2557_v97
			var _nw426 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw426
			_2557_v97 = _nw426
			var _2558_v98 _dafny.Map
			_ = _2558_v98
			_2558_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2557_v97, _2423_v3)
			var _2559_v99 D21
			_ = _2559_v99
			_2559_v99 = Companion_D21_.Create_DC42_(_2558_v98)
			var _2560_v100 _dafny.Set
			_ = _2560_v100
			_2560_v100 = _dafny.SetOf(_2421_v1)
			var _2561_v101 _dafny.Map
			_ = _2561_v101
			_2561_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v1, false)
			var _2562_v102 D0
			_ = _2562_v102
			_2562_v102 = Companion_D0_.Create_DC0_((_2561_v101).Cardinality())
			var _rhs432 _dafny.MultiSet = ((func() _dafny.MultiSet {
				if _2421_v1 {
					return _2555_v95
				}
				return _2555_v95
			})()).Difference((_2555_v95).Difference(_dafny.MultiSetOf((func() _dafny.Int {
				if (_2422_v2).Contains(_2421_v1) {
					return (_2422_v2).Multiplicity(_2421_v1)
				}
				return _2423_v3
			})(), _2423_v3, (_2560_v100).Cardinality(), _2423_v3, _dafny.IntOfInt64(309))))
			_ = _rhs432
			var _rhs433 D21 = _2559_v99
			_ = _rhs433
			var _rhs434 _dafny.Int = (_2562_v102).Dtor_cf0()
			_ = _rhs434
			var _lhs342 *GlobalState = globalState
			_ = _lhs342
			var _lhs343 *GlobalState = globalState
			_ = _lhs343
			_lhs342.F2 = _rhs432
			_2559_v99 = _rhs433
			_lhs343.F3 = _rhs434
			var _2563_v103 _dafny.Sequence
			_ = _2563_v103
			_2563_v103 = _dafny.UnicodeSeqOfUtf8Bytes("jo")
			var _2564_v104 _dafny.Map
			_ = _2564_v104
			_2564_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_2563_v103).Cardinality())).Cmp(_2423_v3) != 0, _2423_v3)
			(globalState).F3 = (_2564_v104).Cardinality()
			(globalState).F8 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_2423_v3, (_2561_v101).Cardinality()), ((_2422_v2).Update(_2421_v1, Companion_Default___.Abs(_2423_v3))).Cardinality()))
		} else {
			var _2565___mcc_h3 D19 = _source40.Get_().(D19_DC40).Cf60
			_ = _2565___mcc_h3
			var _2566_cf60 D19 = _2565___mcc_h3
			_ = _2566_cf60
			var _2567_v105 _dafny.Sequence
			_ = _2567_v105
			_2567_v105 = _dafny.UnicodeSeqOfUtf8Bytes("fgp")
			var _rhs435 bool = _2421_v1
			_ = _rhs435
			var _rhs436 bool = _2421_v1
			_ = _rhs436
			var _rhs437 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_2567_v105, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2567_v105, (Companion_Default___.SafeIndex(_2423_v3, _dafny.IntOfUint32((_2567_v105).Cardinality()))).Uint32(), _2540_v86), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg120 _dafny.Int) interface{} {
					return coer120(arg120)
				}
			}((func(_2568_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2569_i16 _dafny.Int) _dafny.CodePoint {
					return _2568_v86
				}
			})(_2540_v86)))))
			_ = _rhs437
			var _rhs438 bool = _2421_v1
			_ = _rhs438
			var _lhs344 *GlobalState = globalState
			_ = _lhs344
			var _lhs345 *GlobalState = globalState
			_ = _lhs345
			_2421_v1 = _rhs435
			_lhs344.F4 = _rhs436
			r0 = _rhs437
			_lhs345.F4 = _rhs438
			var _2570_v106 _dafny.Array
			_ = _2570_v106
			var _nw427 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw427
			_2570_v106 = _nw427
			_2570_v106 = _2570_v106
			var _2571_v107 _dafny.Set
			_ = _2571_v107
			_2571_v107 = _dafny.SetOf(_2540_v86)
			var _2572_v108 *C3
			_ = _2572_v108
			var _nw428 *C3 = New_C3_()
			_ = _nw428
			_nw428.Ctor__()
			_2572_v108 = _nw428
			var _2573_v109 _dafny.Map
			_ = _2573_v109
			_2573_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2571_v107).Intersection(_2571_v107), _2572_v108)
			_2573_v109 = (_2573_v109).Update(_2571_v107, _2572_v108)
			(globalState).F4 = _2421_v1
		}
		var _2574_v110 _dafny.Map
		_ = _2574_v110
		_2574_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v1, _2540_v86)
		(globalState).F4 = !((_2574_v110).Merge(_2574_v110)).Contains(!(_2421_v1))
		r0 = _2421_v1
		r1 = _2421_v1
		return r0, r1
	}
}
func (_this *C10) M1(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _2575_v0 _dafny.Map
		_ = _2575_v0
		_2575_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _2576_v1 _dafny.MultiSet
		_ = _2576_v1
		_2576_v1 = _dafny.MultiSetOf((func() bool {
			if (_2575_v0).Contains(p0) {
				return (_2575_v0).Get(p0).(bool)
			}
			return p1
		})(), p1, p1)
		var _source41 D2 = Companion_Default___.Fm40((_2576_v1).IsSubsetOf(_dafny.MultiSetOf(p1)), globalState)
		_ = _source41
		if _source41.Is_DC6() {
			var _2577___mcc_h0 _dafny.Int = _source41.Get_().(D2_DC6).Cf6
			_ = _2577___mcc_h0
			var _2578___mcc_h1 bool = _source41.Get_().(D2_DC6).Cf7
			_ = _2578___mcc_h1
			var _2579___mcc_h2 bool = _source41.Get_().(D2_DC6).Cf8
			_ = _2579___mcc_h2
			var _2580_cf8 bool = _2579___mcc_h2
			_ = _2580_cf8
			var _2581_cf7 bool = _2578___mcc_h1
			_ = _2581_cf7
			var _2582_cf6 _dafny.Int = _2577___mcc_h0
			_ = _2582_cf6
			var _2583_v2 *C3
			_ = _2583_v2
			var _nw429 *C3 = New_C3_()
			_ = _nw429
			_nw429.Ctor__()
			_2583_v2 = _nw429
			var _2584_v3 _dafny.MultiSet
			_ = _2584_v3
			_2584_v3 = _dafny.MultiSetOf(_2583_v2)
			var _2585_v4 _dafny.Set
			_ = _2585_v4
			_2585_v4 = _dafny.SetOf((_2582_cf6).Minus(p0), (_2582_cf6).Plus((_2584_v3).Cardinality()), p0, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(293), _2582_cf6), _2582_cf6)
			var _2586_v5 _dafny.Array
			_ = _2586_v5
			var _nw430 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw430
			_2586_v5 = _nw430
			var _index454 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))
			_ = _index454
			(_2586_v5).ArraySet1(!(p1), (_index454).Int())
			var _2587_v6 _dafny.Array
			_ = _2587_v6
			var _nw431 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw431
			_2587_v6 = _nw431
			var _2588_v7 _dafny.Sequence
			_ = _2588_v7
			_2588_v7 = _dafny.UnicodeSeqOfUtf8Bytes("wbp")
			var _index455 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))
			_ = _index455
			(_2587_v6).ArraySet1(_2588_v7, (_index455).Int())
			var _index456 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))
			_ = _index456
			var _index457 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))
			_ = _index457
			var _rhs439 _dafny.Set = (_2585_v4).Difference(_2585_v4)
			_ = _rhs439
			var _rhs440 bool = ((p0).Minus(p0)).Cmp((_dafny.Zero).Minus((p0).Minus(_2582_cf6))) != 0
			_ = _rhs440
			var _rhs441 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2588_v7, _2588_v7)
			_ = _rhs441
			var _rhs442 bool = _2580_cf8
			_ = _rhs442
			var _lhs346 _dafny.Array = _2586_v5
			_ = _lhs346
			var _lhs347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))
			_ = _lhs347
			var _lhs348 _dafny.Array = _2587_v6
			_ = _lhs348
			var _lhs349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))
			_ = _lhs349
			var _lhs350 *GlobalState = globalState
			_ = _lhs350
			_2585_v4 = _rhs439
			(_lhs346).ArraySet1(_rhs440, (_lhs347).Int())
			(_lhs348).ArraySet1(_rhs441, (_lhs349).Int())
			_lhs350.F4 = _rhs442
			var _2589_v8 _dafny.CodePoint
			_ = _2589_v8
			_2589_v8 = _dafny.CodePoint('y')
			var _2590_v9 _dafny.Map
			_ = _2590_v9
			_2590_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Contains((_2587_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))).Int()).(_dafny.Sequence), _2589_v8))
			_2590_v9 = Companion_Default___.Fm21(!((_2586_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))).Int()).(bool)), globalState)
			var _index458 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))
			_ = _index458
			(_2586_v5).ArraySet1(!((_2586_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_2586_v5), 0))).Int()).(bool)), (_index458).Int())
			var _2591_v11 _dafny.Sequence
			_ = _2591_v11
			_2591_v11 = _dafny.SeqOf((func() _dafny.Map {
				var _coll113 = _dafny.NewMapBuilder()
				_ = _coll113
				for _iter117 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-461), _dafny.IntOfInt64(451))); ; {
					_compr_113, _ok117 := _iter117()
					if !_ok117 {
						break
					}
					var _2592_v10 _dafny.Int
					_2592_v10 = interface{}(_compr_113).(_dafny.Int)
					if ((_dafny.IntOfInt64(-461)).Cmp(_2592_v10) <= 0) && ((_2592_v10).Cmp(_dafny.IntOfInt64(451)) < 0) {
						_coll113.Add((_2592_v10).Minus(p0), _dafny.Companion_Sequence_.Update((_2587_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_2587_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2587_v6), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _2589_v8))
					}
				}
				return _coll113.ToMap()
			}()).Cardinality())
			(globalState).F0 = _dafny.Companion_Sequence_.Update(_2591_v11, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.IntOfUint32((_2591_v11).Cardinality()))).Uint32(), p0)
		} else if _source41.Is_DC7() {
			var _2593___mcc_h3 _dafny.CodePoint = _source41.Get_().(D2_DC7).Cf9
			_ = _2593___mcc_h3
			var _2594___mcc_h4 bool = _source41.Get_().(D2_DC7).Cf10
			_ = _2594___mcc_h4
			var _2595___mcc_h5 _dafny.Int = _source41.Get_().(D2_DC7).Cf11
			_ = _2595___mcc_h5
			var _2596_cf11 _dafny.Int = _2595___mcc_h5
			_ = _2596_cf11
			var _2597_cf10 bool = _2594___mcc_h4
			_ = _2597_cf10
			var _2598_cf9 _dafny.CodePoint = _2593___mcc_h3
			_ = _2598_cf9
			var _2599_v12 *C5
			_ = _2599_v12
			var _nw432 *C5 = New_C5_()
			_ = _nw432
			_nw432.Ctor__()
			_2599_v12 = _nw432
			var _2600_v13 *C5
			_ = _2600_v13
			var _nw433 *C5 = New_C5_()
			_ = _nw433
			_nw433.Ctor__()
			_2600_v13 = _nw433
			(globalState).F8 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-857))
			if !(_2576_v1).Equals(_2576_v1) {
				var _2601_v14 _dafny.Array
				_ = _2601_v14
				var _len0_71 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_71
				var _nw434 _dafny.Array
				_ = _nw434
				if _len0_71.Cmp(_dafny.Zero) == 0 {
					_nw434 = _dafny.NewArray(_len0_71)
				} else {
					var _init71 func(_dafny.Int) _dafny.Int = (func(_2602_cf11 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2603_i0 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_2603_i0, _2602_cf11)
						}
					})(_2596_cf11)
					_ = _init71
					var _element0_71 = _init71(_dafny.Zero)
					_ = _element0_71
					_nw434 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
					(_nw434).ArraySet1(_element0_71, 0)
					var _nativeLen0_71 = (_len0_71).Int()
					_ = _nativeLen0_71
					for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
						(_nw434).ArraySet1(_init71(_dafny.IntOf(_i0_71)), _i0_71)
					}
				}
				_2601_v14 = _nw434
				_2601_v14 = _2601_v14
				var _2604_v15 _dafny.Array
				_ = _2604_v15
				var _nw435 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
				_ = _nw435
				_2604_v15 = _nw435
				var _2605_v16 *C1
				_ = _2605_v16
				var _nw436 *C1 = New_C1_()
				_ = _nw436
				_nw436.Ctor__()
				_2605_v16 = _nw436
				var _index459 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_2604_v15), 0))
				_ = _index459
				(_2604_v15).ArraySet1(_2605_v16, (_index459).Int())
				var _2606_v17 _dafny.Map
				_ = _2606_v17
				_2606_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2601_v14)
				var _2607_v18 T0
				_ = _2607_v18
				var _nw437 *C3 = New_C3_()
				_ = _nw437
				_nw437.Ctor__()
				_2607_v18 = _nw437
				var _2608_v19 _dafny.Sequence
				_ = _2608_v19
				_2608_v19 = _dafny.SeqOf(_2606_v17, (func() _dafny.Map {
					if _2597_cf10 {
						return _2606_v17
					}
					return _2606_v17
				})(), _2606_v17)
				var _2609_v20 _dafny.Sequence
				_ = _2609_v20
				_2609_v20 = _dafny.SeqOf(_2601_v14, (func() _dafny.Array {
					if (_2606_v17).Contains(p1) {
						return (_2606_v17).Get(p1).(_dafny.Array)
					}
					return _2601_v14
				})())
				var _index460 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_2604_v15), 0))
				_ = _index460
				var _rhs443 *C1 = _2605_v16
				_ = _rhs443
				var _rhs444 _dafny.Int = _2596_cf11
				_ = _rhs444
				var _rhs445 _dafny.Map = (_2608_v19).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, _2596_cf11), _dafny.IntOfUint32((_2608_v19).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _rhs445
				var _rhs446 T0 = _2607_v18
				_ = _rhs446
				var _rhs447 bool = !_dafny.Companion_Sequence_.Contains(_2609_v20, _2601_v14)
				_ = _rhs447
				var _lhs351 _dafny.Array = _2604_v15
				_ = _lhs351
				var _lhs352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_2604_v15), 0))
				_ = _lhs352
				var _lhs353 *GlobalState = globalState
				_ = _lhs353
				(_lhs351).ArraySet1(_rhs443, (_lhs352).Int())
				_lhs353.F3 = _rhs444
				_2606_v17 = _rhs445
				_2607_v18 = _rhs446
				_2597_cf10 = _rhs447
				var _2610_v21 _dafny.Map
				_ = _2610_v21
				_2610_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2597_cf10) || (p1), (_dafny.Zero).Minus(p0))
				_2610_v21 = (_2610_v21).Update(_2597_cf10, p0)
				var _index461 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_2601_v14), 0))
				_ = _index461
				(_2601_v14).ArraySet1((_dafny.Zero).Minus(p0), (_index461).Int())
				var _index462 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_2601_v14), 0))
				_ = _index462
				(_2601_v14).ArraySet1(Companion_Default___.Fm6(p0, globalState), (_index462).Int())
				r0 = _dafny.IntOfInt64(321)
			} else {
				var _2611_v22 _dafny.Sequence
				_ = _2611_v22
				_2611_v22 = _dafny.UnicodeSeqOfUtf8Bytes("k")
				_2611_v22 = _2611_v22
				(globalState).F3 = p0
				(globalState).F8 = (Companion_Default___.Fm20(_2597_cf10, globalState)).Plus((_dafny.IntOfInt64(212)).Minus(_2596_cf11))
				(globalState).F4 = _2597_cf10
				var _2612_v23 D2
				_ = _2612_v23
				_2612_v23 = Companion_D2_.Create_DC6_(_2596_cf11, false, p1)
				_2612_v23 = _2612_v23
			}
		} else {
			var _2613___mcc_h6 _dafny.Array = _source41.Get_().(D2_DC5).Cf5
			_ = _2613___mcc_h6
			var _2614_cf5 _dafny.Array = _2613___mcc_h6
			_ = _2614_cf5
			var _2615_v24 _dafny.Array
			_ = _2615_v24
			var _nw438 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
			_ = _nw438
			_2615_v24 = _nw438
			var _2616_v25 _dafny.Map
			_ = _2616_v25
			_2616_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2615_v24)
			var _2617_v26 *C4
			_ = _2617_v26
			var _nw439 *C4 = New_C4_()
			_ = _nw439
			_nw439.Ctor__((func() _dafny.Array {
				if (_2616_v25).Contains(p1) {
					return (_2616_v25).Get(p1).(_dafny.Array)
				}
				return _2615_v24
			})())
			_2617_v26 = _nw439
			var _2618_v27 _dafny.Array
			_ = _2618_v27
			var _len0_72 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_72
			var _nw440 _dafny.Array
			_ = _nw440
			if _len0_72.Cmp(_dafny.Zero) == 0 {
				_nw440 = _dafny.NewArray(_len0_72)
			} else {
				var _init72 func(_dafny.Int) _dafny.Int = (func(_2619_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2620_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_2620_i1, _2619_p0)
					}
				})(p0)
				_ = _init72
				var _element0_72 = _init72(_dafny.Zero)
				_ = _element0_72
				_nw440 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
				(_nw440).ArraySet1(_element0_72, 0)
				var _nativeLen0_72 = (_len0_72).Int()
				_ = _nativeLen0_72
				for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
					(_nw440).ArraySet1(_init72(_dafny.IntOf(_i0_72)), _i0_72)
				}
			}
			_2618_v27 = _nw440
			var _index463 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_2618_v27), 0))
			_ = _index463
			(_2618_v27).ArraySet1((p0).Minus(p0), (_index463).Int())
			var _2621_v28 _dafny.Array
			_ = _2621_v28
			var _nwElement0_90 _dafny.Array = _2618_v27
			_ = _nwElement0_90
			var _nw441 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(11))
			_ = _nw441
			(_nw441).ArraySet1(_nwElement0_90, 0)
			(_nw441).ArraySet1(_2618_v27, 1)
			(_nw441).ArraySet1(_2618_v27, 2)
			(_nw441).ArraySet1(_2618_v27, 3)
			(_nw441).ArraySet1(_2618_v27, 4)
			(_nw441).ArraySet1(_2618_v27, 5)
			(_nw441).ArraySet1(_2618_v27, 6)
			(_nw441).ArraySet1(_2618_v27, 7)
			(_nw441).ArraySet1(_2618_v27, 8)
			(_nw441).ArraySet1(_2618_v27, 9)
			(_nw441).ArraySet1(_2618_v27, 10)
			_2621_v28 = _nw441
			var _index464 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_2621_v28), 0))
			_ = _index464
			(_2621_v28).ArraySet1(_2618_v27, (_index464).Int())
			var _index465 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_2618_v27), 0))
			_ = _index465
			var _index466 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_2621_v28), 0))
			_ = _index466
			var _rhs448 _dafny.Int = p0
			_ = _rhs448
			var _rhs449 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(-463))
			_ = _rhs449
			var _rhs450 _dafny.Int = (p0).Minus(p0)
			_ = _rhs450
			var _rhs451 _dafny.Array = _2618_v27
			_ = _rhs451
			var _lhs354 _dafny.Array = _2618_v27
			_ = _lhs354
			var _lhs355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_2618_v27), 0))
			_ = _lhs355
			var _lhs356 *GlobalState = globalState
			_ = _lhs356
			var _lhs357 *GlobalState = globalState
			_ = _lhs357
			var _lhs358 _dafny.Array = _2621_v28
			_ = _lhs358
			var _lhs359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_2621_v28), 0))
			_ = _lhs359
			(_lhs354).ArraySet1(_rhs448, (_lhs355).Int())
			_lhs356.F3 = _rhs449
			_lhs357.F8 = _rhs450
			(_lhs358).ArraySet1(_rhs451, (_lhs359).Int())
			(globalState).F4 = p1
			var _2622_v29 bool
			_ = _2622_v29
			var _2623_v30 bool
			_ = _2623_v30
			var _out29 bool
			_ = _out29
			var _out30 bool
			_ = _out30
			_out29, _out30 = (_2617_v26).M0(globalState)
			_2622_v29 = _out29
			_2623_v30 = _out30
		}
		var _hi16 _dafny.Int = (p0).Plus(Companion_Default___.Fm6(p0, globalState))
		_ = _hi16
		for _2624_i2 := p0; _2624_i2.Cmp(_hi16) < 0; _2624_i2 = _2624_i2.Plus(_dafny.One) {
			var _2625_v31 _dafny.Array
			_ = _2625_v31
			var _nw442 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw442
			_2625_v31 = _nw442
			var _nw443 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw443
			_2625_v31 = _nw443
			var _2626_v32 _dafny.Array
			_ = _2626_v32
			var _len0_73 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_73
			var _nw444 _dafny.Array
			_ = _nw444
			if _len0_73.Cmp(_dafny.Zero) == 0 {
				_nw444 = _dafny.NewArray(_len0_73)
			} else {
				var _init73 func(_dafny.Int) bool = func(_2627_i3 _dafny.Int) bool {
					return true
				}
				_ = _init73
				var _element0_73 = _init73(_dafny.Zero)
				_ = _element0_73
				_nw444 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
				(_nw444).ArraySet1(_element0_73, 0)
				var _nativeLen0_73 = (_len0_73).Int()
				_ = _nativeLen0_73
				for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
					(_nw444).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
				}
			}
			_2626_v32 = _nw444
			_2626_v32 = (func() _dafny.Array {
				if !(p1) {
					return _2626_v32
				}
				return _2626_v32
			})()
			if p1 {
				(globalState).F4 = p1
				(globalState).F3 = p0
				var _2628_v33 _dafny.Array
				_ = _2628_v33
				var _len0_74 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_74
				var _nw445 _dafny.Array
				_ = _nw445
				if _len0_74.Cmp(_dafny.Zero) == 0 {
					_nw445 = _dafny.NewArray(_len0_74)
				} else {
					var _init74 func(_dafny.Int) _dafny.Int = (func(_2629_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2630_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_2630_i4, (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer121 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg121 _dafny.Int) interface{} {
									return coer121(arg121)
								}
							}((func(_2631_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_2632_i5 _dafny.Int) _dafny.Int {
									return _2631_i2
								}
							})(_2629_i2))))).Cardinality())
						}
					})(_2624_i2)
					_ = _init74
					var _element0_74 = _init74(_dafny.Zero)
					_ = _element0_74
					_nw445 = _dafny.NewArrayFromExample(_element0_74, nil, _len0_74)
					(_nw445).ArraySet1(_element0_74, 0)
					var _nativeLen0_74 = (_len0_74).Int()
					_ = _nativeLen0_74
					for _i0_74 := 1; _i0_74 < _nativeLen0_74; _i0_74++ {
						(_nw445).ArraySet1(_init74(_dafny.IntOf(_i0_74)), _i0_74)
					}
				}
				_2628_v33 = _nw445
				var _2633_v34 _dafny.Sequence
				_ = _2633_v34
				_2633_v34 = _dafny.UnicodeSeqOfUtf8Bytes("cqbcks")
				var _2634_v35 _dafny.Sequence
				_ = _2634_v35
				_2634_v35 = _dafny.SeqOf(_2633_v34)
				var _2635_v36 _dafny.Array
				_ = _2635_v36
				_2635_v36 = _2628_v33
				var _rhs452 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_2633_v34), _2634_v35)
				_ = _rhs452
				var _rhs453 _dafny.Array = (_2635_v36)
				_ = _rhs453
				var _lhs360 *GlobalState = globalState
				_ = _lhs360
				_lhs360.F4 = _rhs452
				_2628_v33 = _rhs453
				var _2636_v37 _dafny.Array
				_ = _2636_v37
				var _nw446 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(13))
				_ = _nw446
				_2636_v37 = _nw446
				_2636_v37 = _2636_v37
				(globalState).F4 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(586))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}((func(_2637_v34 _dafny.Sequence, _2638_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_2639_i6 _dafny.Int) _dafny.CodePoint {
						return (_2637_v34).Select((Companion_Default___.SafeIndex(_2638_p0, _dafny.IntOfUint32((_2637_v34).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_2633_v34, p0))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(738))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}(func(_2640_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})))
			} else {
				var _2641_v38 _dafny.CodePoint
				_ = _2641_v38
				_2641_v38 = _dafny.CodePoint('l')
				var _2642_v39 _dafny.Sequence
				_ = _2642_v39
				_2642_v39 = _dafny.SeqOf(_2641_v38, _2641_v38)
				var _2643_v40 D10
				_ = _2643_v40
				_2643_v40 = Companion_D10_.Create_DC24_(p0, p1, _2642_v39)
				var _2644_v41 _dafny.Sequence
				_ = _2644_v41
				_2644_v41 = _dafny.SeqOf(_2642_v39, _dafny.Companion_Sequence_.Update((_2643_v40).Dtor_cf33(), (Companion_Default___.SafeIndex(_2624_i2, _dafny.IntOfUint32(((_2643_v40).Dtor_cf33()).Cardinality()))).Uint32(), _2641_v38))
				var _2645_v42 _dafny.Map
				_ = _2645_v42
				_2645_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm20(p1, globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2641_v38, _2641_v38, _2641_v38), (_2644_v41).Select((Companion_Default___.SafeIndex(_2624_i2, _dafny.IntOfUint32((_2644_v41).Cardinality()))).Uint32()).(_dafny.Sequence)))
				_2645_v42 = (_2645_v42).Update((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, p0)), _2642_v39)
				(globalState).F4 = p1
				_2641_v38 = (func() _dafny.CodePoint {
					if p1 {
						return _2641_v38
					}
					return _2641_v38
				})()
				var _2646_v43 D25
				_ = _2646_v43
				_2646_v43 = Companion_D25_.Create_DC52_(_2625_v31)
				var _2647_v44 _dafny.Array
				_ = _2647_v44
				var _nwElement0_91 _dafny.Array = (_2646_v43).Dtor_cf73()
				_ = _nwElement0_91
				var _nw447 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(27))
				_ = _nw447
				(_nw447).ArraySet1(_nwElement0_91, 0)
				(_nw447).ArraySet1(_2625_v31, 1)
				(_nw447).ArraySet1(_2625_v31, 2)
				(_nw447).ArraySet1(_2625_v31, 3)
				(_nw447).ArraySet1(_2625_v31, 4)
				(_nw447).ArraySet1(_2625_v31, 5)
				(_nw447).ArraySet1(_2625_v31, 6)
				(_nw447).ArraySet1(_2625_v31, 7)
				(_nw447).ArraySet1((func() _dafny.Array {
					if p1 {
						return _2625_v31
					}
					return _2625_v31
				})(), 8)
				(_nw447).ArraySet1(_2625_v31, 9)
				(_nw447).ArraySet1(_2625_v31, 10)
				(_nw447).ArraySet1(_2625_v31, 11)
				(_nw447).ArraySet1(_2625_v31, 12)
				(_nw447).ArraySet1(_2625_v31, 13)
				(_nw447).ArraySet1((func() _dafny.Array {
					if p1 {
						return _2625_v31
					}
					return _2625_v31
				})(), 14)
				(_nw447).ArraySet1(_2625_v31, 15)
				(_nw447).ArraySet1(_2625_v31, 16)
				(_nw447).ArraySet1(_2625_v31, 17)
				(_nw447).ArraySet1(_2625_v31, 18)
				(_nw447).ArraySet1(_2625_v31, 19)
				(_nw447).ArraySet1(_2625_v31, 20)
				(_nw447).ArraySet1(_2625_v31, 21)
				(_nw447).ArraySet1(_2625_v31, 22)
				(_nw447).ArraySet1(_2625_v31, 23)
				(_nw447).ArraySet1((_2646_v43).Dtor_cf73(), 24)
				(_nw447).ArraySet1(_2625_v31, 25)
				(_nw447).ArraySet1(_2625_v31, 26)
				_2647_v44 = _nw447
				var _index467 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_2647_v44), 0))
				_ = _index467
				(_2647_v44).ArraySet1(_2625_v31, (_index467).Int())
				var _index468 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_2647_v44), 0))
				_ = _index468
				(_2647_v44).ArraySet1(_2625_v31, (_index468).Int())
				(globalState).F8 = (p0).Plus(_2624_i2)
			}
			var _2648_v45 _dafny.Set
			_ = _2648_v45
			_2648_v45 = _dafny.SetOf(p1)
			var _index469 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2626_v32), 0))
			_ = _index469
			(_2626_v32).ArraySet1((_2648_v45).IsDisjointFrom(_dafny.SetOf(p1, p1, false)), (_index469).Int())
			var _index470 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2626_v32), 0))
			_ = _index470
			(_2626_v32).ArraySet1(p1, (_index470).Int())
		}
		r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(357), (_dafny.IntOfInt64(663)).Times(p0))
		var _2649_v46 _dafny.CodePoint
		_ = _2649_v46
		_2649_v46 = _dafny.CodePoint('h')
		var _2650_v47 _dafny.Map
		_ = _2650_v47
		_2650_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Plus(p0), _2649_v46)
		_2650_v47 = (_2650_v47).Update(p0, (func() _dafny.CodePoint {
			if p1 {
				return _dafny.CodePoint('a')
			}
			return _2649_v46
		})())
		var _2651_v48 _dafny.Array
		_ = _2651_v48
		var _nw448 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw448
		_2651_v48 = _nw448
		var _2652_v49 _dafny.Sequence
		_ = _2652_v49
		_2652_v49 = _dafny.UnicodeSeqOfUtf8Bytes("msd")
		var _2653_v50 _dafny.Sequence
		_ = _2653_v50
		_2653_v50 = _dafny.SeqOf(_2652_v49)
		var _2654_v51 D4
		_ = _2654_v51
		_2654_v51 = Companion_D4_.Create_DC11_(p1, p1, p0, _2652_v49)
		var _2655_v52 D4
		_ = _2655_v52
		_2655_v52 = Companion_D4_.Create_DC12_(_2654_v51)
		var _2656_v53 _dafny.Sequence
		_ = _2656_v53
		_2656_v53 = _dafny.SeqOf(_2655_v52, Companion_D4_.Create_DC12_(_2654_v51))
		var _2657_v54 _dafny.Map
		_ = _2657_v54
		_2657_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2651_v48, (_2653_v50).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2656_v53).Cardinality()), _dafny.IntOfUint32((_2653_v50).Cardinality()))).Uint32()).(_dafny.Sequence))
		_2657_v54 = (_2657_v54).Merge((_2657_v54).Merge(_2657_v54))
		(globalState).F4 = Companion_Default___.Fm49(p1, globalState)
		var _2658_v55 _dafny.Set
		_ = _2658_v55
		_2658_v55 = _dafny.SetOf(_dafny.IntOfInt64(167), p0)
		var _2659_v56 _dafny.Map
		_ = _2659_v56
		_2659_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm49(p1, globalState), (_2658_v55).Cardinality())
		var _2660_v57 _dafny.MultiSet
		_ = _2660_v57
		_2660_v57 = _dafny.MultiSetOf(_2659_v56, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)).Update(p1, p0))
		var _2661_v58 _dafny.Sequence
		_ = _2661_v58
		_2661_v58 = _dafny.SeqOf((_2660_v57).Cardinality())
		r0 = (p0).Plus((_2661_v58).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2661_v58).Cardinality()))).Uint32()).(_dafny.Int))
		return r0
	}
}

// End of class C10
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
