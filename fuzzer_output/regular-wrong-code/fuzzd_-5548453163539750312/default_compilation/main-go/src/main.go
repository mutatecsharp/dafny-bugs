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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(874), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(773), !(true))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(431), false), func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(160), _dafny.IntOfInt64(497))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(160)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(497)) < 0) {
				_coll0.Add((_0_v0).Plus((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, false, true, true))).Cardinality()), false)
			}
		}
		return _coll0.ToMap()
	}()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(416), true), func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))).Cardinality()), !(true))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			}))).Cardinality()), !(true))).Contains(_2_v1) {
				_coll1.Add(Companion_Default___.SafeModuloInt(_2_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(334), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-137), _dafny.IntOfInt64(864), _dafny.IntOfInt64(242))).Cardinality())), true)
			}
		}
		return _coll1.ToMap()
	}())))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return (((_dafny.SetOf(true)).Union(_dafny.SetOf(!(true)))).Cardinality()).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(583))).Cardinality()), false)).Cardinality()).Times(_dafny.IntOfInt64(-101)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wlhgm")).Cardinality()), _dafny.IntOfInt64(865))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), (_dafny.SetOf(!(false))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-747), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfpaxux")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) bool {
	return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xtwhavygd"), _dafny.IntOfInt64(606))).Cardinality(), !(false)))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(794), !(true))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))).Merge((func() _dafny.Map {
		if !(false) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
	})())
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(!(false), ((_dafny.Zero).Minus(_dafny.IntOfInt64(-641))).Cmp(_dafny.IntOfInt64(912)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('a')
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 D3, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("yathjppk"), _dafny.UnicodeSeqOfUtf8Bytes("g")), (_dafny.SetOf(!(false))).IsSubsetOf(_dafny.SetOf(false, false)))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((Companion_D3_.Create_DC8_(_dafny.UnicodeSeqOfUtf8Bytes("gdigkgpmo"))).Dtor_cf14(), _dafny.UnicodeSeqOfUtf8Bytes("ybuwlsavw"))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Union(_dafny.MultiSetOf(true, false, !(false)))).Union(_dafny.MultiSetOf(true, true))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if (false) == (!(!(true))) {
		return _dafny.MultiSetOf(false)
	} else {
		return (_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(true, false, true, !(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(true)), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.SeqOf(true, true, false, true, !(!(!(false))))).Cardinality()))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(!(true), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dqkjudxne")).Cardinality()))).Cardinality())), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(478))))).Union(_dafny.SetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lcy")).Cardinality()))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(!(true), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))).Cardinality()))))).Union(_dafny.SetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(false))))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(!(true), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(650))).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true, (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hjukjem"), _dafny.UnicodeSeqOfUtf8Bytes("wuyw")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("epjwidlg"), _dafny.UnicodeSeqOfUtf8Bytes("ieovtcopc")))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D9 = Companion_D9_.Create_DC27_(Companion_D9_.Create_DC26_())
	_ = _source0
	if _source0.Is_DC26() {
		return _dafny.CodePoint('l')
	} else if _source0.Is_DC25() {
		var _3___mcc_h0 _dafny.Sequence = _source0.Get_().(D9_DC25).Cf50
		_ = _3___mcc_h0
		var _4_cf50 _dafny.Sequence = _3___mcc_h0
		_ = _4_cf50
		return _dafny.CodePoint('u')
	} else {
		var _5___mcc_h1 D9 = _source0.Get_().(D9_DC27).Cf51
		_ = _5___mcc_h1
		var _6_cf51 D9 = _5___mcc_h1
		_ = _6_cf51
		if false {
			return _dafny.CodePoint('j')
		} else {
			return _dafny.CodePoint('i')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D1_.Create_DC5_(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(57))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(true)).Cardinality())).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(859))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), Companion_D1_.Create_DC5_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("np")).Cardinality()), _dafny.IntOfInt64(692), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(!(false))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D1_.Create_DC5_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), (func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(354), _dafny.IntOfInt64(741))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(354)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(741)) < 0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_7_v0, _dafny.IntOfInt64(168)))
			}
		}
		return _coll2.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(-782), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(652))))))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	return !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eaasr")).Cardinality())).Cmp((Companion_D15_.Create_DC37_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-626))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(218)
	}))).Cardinality()))).Dtor_cf67()) > 0) || (false)
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(437), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-733))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality())), (Companion_D10_.Create_DC28_(_dafny.SeqOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Dtor_cf52())
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Set, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true, !(_dafny.SetOf(true, false)).Equals(_dafny.SetOf(true, false, true)))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(261))).Uint32(), func(coer4 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC4_(false, false)
	}))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(938), true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(495), false)))
}
func (_static *CompanionStruct_Default___) Fm35(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, true)).Intersection(_dafny.SetOf((Companion_D3_.Create_DC10_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Dtor_cf16(), true, true)), (_dafny.IntOfInt64(536)).Cmp(_dafny.IntOfInt64(-239)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(30)).Cmp(_dafny.IntOfInt64(650)) > 0, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.SetOf(false, !(true))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Cardinality())), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(676), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("crt")).Cardinality()))))).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(-349), _dafny.IntOfInt64(-137), _dafny.IntOfInt64(200))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(633))))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SetOf(false, true))).Intersection((_dafny.SetOf(_dafny.SetOf(!(true)))).Intersection(_dafny.SetOf(_dafny.SetOf(false, true, true))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Set, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-876))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(Companion_D6_.Create_DC16_(_dafny.SeqOf(false, true))))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _12_v0 _dafny.Sequence
			_12_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(Companion_D6_.Create_DC16_(_dafny.SeqOf(false, true)))), _12_v0) {
				_coll3.Add(_12_v0, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-294), (_dafny.MultiSetOf(_dafny.IntOfInt64(238), _dafny.IntOfInt64(653), _dafny.IntOfInt64(98))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(-670))))
			}
		}
		return _coll3.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm41(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(((Companion_D7_.Create_DC19_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false)))).Dtor_cf34()).Cardinality(), _dafny.IntOfInt64(-751), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(146))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_13_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_14_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32(((Companion_D10_.Create_DC28_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(254), false)).Cardinality()))).Dtor_cf52()).Cardinality())
	})))).Cardinality(), _dafny.IntOfInt64(192))).Union(((Companion_D8_.Create_DC23_(_dafny.IntOfInt64(319), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hm")).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tcggd")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(581), _dafny.IntOfInt64(-312))).Dtor_cf41()).Union(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-842), _dafny.IntOfInt64(321))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _15_v0 _dafny.Int
			_15_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-842)).Cmp(_15_v0) <= 0) && ((_15_v0).Cmp(_dafny.IntOfInt64(321)) < 0) {
				_coll4.Add(Companion_Default___.SafeModuloInt(_15_v0, _dafny.IntOfInt64(646)))
			}
		}
		return _coll4.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _16_v0 bool
	_ = _16_v0
	_16_v0 = true
	var _17_v1 _dafny.Map
	_ = _17_v1
	_17_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-313)))
	var _18_v2 _dafny.Sequence
	_ = _18_v2
	_18_v2 = _dafny.SeqOf(_dafny.SetOf(_16_v0, _16_v0))
	var _19_v3 _dafny.Array
	_ = _19_v3
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
	_ = _nw0
	_19_v3 = _nw0
	var _20_v4 _dafny.Int
	_ = _20_v4
	_20_v4 = _dafny.IntOfInt64(333)
	var _21_v5 _dafny.Map
	_ = _21_v5
	_21_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_20_v4, _20_v4)
	var _22_v6 _dafny.Map
	_ = _22_v6
	_22_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v3, _21_v5)
	var _23_v7 _dafny.Sequence
	_ = _23_v7
	_23_v7 = _dafny.UnicodeSeqOfUtf8Bytes("hllu")
	var _24_v8 _dafny.Array
	_ = _24_v8
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.MultiSet = (func(_25_v4 _dafny.Int, _26_v0 bool) func(_dafny.Int) _dafny.MultiSet {
			return func(_27_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_25_v4, (_dafny.Zero).Minus(_25_v4), (_dafny.MultiSetOf(_26_v0, _26_v0)).Cardinality(), _dafny.IntOfInt64(837))
			}
		})(_20_v4, _16_v0)
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
	_24_v8 = _nw1
	var _28_globalState *GlobalState
	_ = _28_globalState
	var _nw2 *GlobalState = New_GlobalState_()
	_ = _nw2
	_nw2.Ctor__(_17_v1, true, _dafny.IntOfInt64(-406), false, _18_v2, _dafny.IntOfInt64(145), false, _dafny.IntOfInt64(873), _dafny.IntOfInt64(506), true, _22_v6, _23_v7, _dafny.IntOfInt64(895), _dafny.IntOfInt64(755), _dafny.IntOfInt64(670), true, _dafny.IntOfInt64(550), _dafny.IntOfInt64(75), _24_v8, true, false, true, _dafny.IntOfInt64(-889))
	_28_globalState = _nw2
	var _29_v9 _dafny.Sequence
	_ = _29_v9
	_29_v9 = _dafny.SeqOf(_dafny.IntOfUint32((_23_v7).Cardinality()))
	var _hi0 _dafny.Int = (func() _dafny.Int {
		if _16_v0 {
			return (_29_v9).Select((Companion_Default___.SafeIndex(_20_v4, _dafny.IntOfUint32((_29_v9).Cardinality()))).Uint32()).(_dafny.Int)
		}
		return _20_v4
	})()
	_ = _hi0
	for _30_i1 := (_20_v4).Minus(_dafny.IntOfUint32((_23_v7).Cardinality())); _30_i1.Cmp(_hi0) < 0; _30_i1 = _30_i1.Plus(_dafny.One) {
		var _31_v10 _dafny.CodePoint
		_ = _31_v10
		_31_v10 = _dafny.CodePoint('v')
		var _32_v11 _dafny.Map
		_ = _32_v11
		_32_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-829), _16_v0))
		var _33_v13 _dafny.Map
		_ = _33_v13
		_33_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_20_v4, _16_v0)
		var _34_v14 _dafny.Sequence
		_ = _34_v14
		_34_v14 = _dafny.SeqOf(func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_33_v13).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _35_v12 _dafny.Int
				_35_v12 = interface{}(_compr_5).(_dafny.Int)
				if (_33_v13).Contains(_35_v12) {
					_coll5.Add(Companion_Default___.SafeDivisionInt(_35_v12, _20_v4), _16_v0)
				}
			}
			return _coll5.ToMap()
		}())
		var _36_v16 _dafny.Sequence
		_ = _36_v16
		_36_v16 = _dafny.SeqOf((func() _dafny.Map {
			if (_32_v11).Contains(_31_v10) {
				return (_32_v11).Get(_31_v10).(_dafny.Map)
			}
			return (_34_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_29_v9).Cardinality()), _dafny.IntOfUint32((_34_v14).Cardinality()))).Uint32()).(_dafny.Map)
		})(), func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(901), _dafny.IntOfInt64(437))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _37_v15 _dafny.Int
				_37_v15 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(901)).Cmp(_37_v15) <= 0) && ((_37_v15).Cmp(_dafny.IntOfInt64(437)) < 0) {
					_coll6.Add((_37_v15).Minus((_dafny.Zero).Minus((_dafny.SetOf(_16_v0, _16_v0)).Cardinality())), _16_v0)
				}
			}
			return _coll6.ToMap()
		}())
		var _38_v17 _dafny.Set
		_ = _38_v17
		_38_v17 = _dafny.SetOf(!(_16_v0))
		_36_v16 = Companion_Default___.Fm0((func() _dafny.Int {
			if _16_v0 {
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-98))
			}
			return Companion_Default___.Fm1(Companion_Default___.Fm1(_20_v4, _16_v0, _28_globalState), !(_16_v0), _28_globalState)
		})(), _20_v4, _23_v7, (_38_v17).Cardinality(), _28_globalState)
		(_28_globalState).F5 = _20_v4
		var _39_v18 _dafny.Set
		_ = _39_v18
		_39_v18 = _dafny.SetOf(_33_v13)
		_39_v18 = (_39_v18).Difference((_39_v18).Intersection(_39_v18))
		var _rhs0 bool = _16_v0
		_ = _rhs0
		var _rhs1 _dafny.CodePoint = _31_v10
		_ = _rhs1
		var _lhs0 *GlobalState = _28_globalState
		_ = _lhs0
		_lhs0.F9 = _rhs0
		_31_v10 = _rhs1
	}
	(_28_globalState).F9 = _16_v0
	if _16_v0 {
		var _40_v19 _dafny.Set
		_ = _40_v19
		_40_v19 = _dafny.SetOf(_16_v0, _16_v0)
		var _41_v20 _dafny.Map
		_ = _41_v20
		_41_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v0, (_40_v19).IsProperSubsetOf(_40_v19))
		var _42_v21 _dafny.Sequence
		_ = _42_v21
		_42_v21 = _dafny.SeqOf(_16_v0)
		(_28_globalState).F6 = (func() bool {
			if (_41_v20).Contains((_42_v21).Select((Companion_Default___.SafeIndex(_20_v4, _dafny.IntOfUint32((_42_v21).Cardinality()))).Uint32()).(bool)) {
				return (_41_v20).Get((_42_v21).Select((Companion_Default___.SafeIndex(_20_v4, _dafny.IntOfUint32((_42_v21).Cardinality()))).Uint32()).(bool)).(bool)
			}
			return !(_16_v0)
		})()
		(_28_globalState).F14 = _20_v4
		var _43_v22 _dafny.Set
		_ = _43_v22
		_43_v22 = _dafny.SetOf(_23_v7)
		if !((_43_v22).IsSubsetOf((_43_v22).Difference(_43_v22))) {
			var _44_v23 _dafny.Array
			_ = _44_v23
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_1
			var _nw3 _dafny.Array
			_ = _nw3
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw3 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.MultiSet = (func(_45_v0 bool) func(_dafny.Int) _dafny.MultiSet {
					return func(_46_i2 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(_dafny.MultiSetOf(false, _45_v0))
					}
				})(_16_v0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw3 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw3).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw3).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_44_v23 = _nw3
			var _47_v24 _dafny.MultiSet
			_ = _47_v24
			_47_v24 = _dafny.MultiSetOf(_16_v0)
			var _48_v25 _dafny.MultiSet
			_ = _48_v25
			_48_v25 = _dafny.MultiSetOf(_47_v24)
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_44_v23), 0))
			_ = _index0
			(_44_v23).ArraySet1(_48_v25, (_index0).Int())
			var _49_v26 _dafny.Map
			_ = _49_v26
			_49_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v0, _48_v25)
			var _50_v27 _dafny.Array
			_ = _50_v27
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw4
			_50_v27 = _nw4
			var _51_v28 _dafny.Map
			_ = _51_v28
			_51_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v27, _16_v0)
			var _52_v30 _dafny.CodePoint
			_ = _52_v30
			_52_v30 = _dafny.CodePoint('t')
			var _53_v31 _dafny.MultiSet
			_ = _53_v31
			_53_v31 = _dafny.MultiSetOf(_52_v30)
			var _54_v32 _dafny.Set
			_ = _54_v32
			_54_v32 = _dafny.SetOf(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_53_v31).Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _55_v29 _dafny.CodePoint
					_55_v29 = interface{}(_compr_7).(_dafny.CodePoint)
					if (_53_v31).Contains(_55_v29) {
						_coll7.Add(_55_v29, _16_v0)
					}
				}
				return _coll7.ToMap()
			}())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_44_v23), 0))
			_ = _index1
			var _rhs2 _dafny.MultiSet = (_48_v25).Difference((func() _dafny.MultiSet {
				if (_49_v26).Contains((func() bool {
					if (_51_v28).Contains(_50_v27) {
						return (_51_v28).Get(_50_v27).(bool)
					}
					return _16_v0
				})()) {
					return (_49_v26).Get((func() bool {
						if (_51_v28).Contains(_50_v27) {
							return (_51_v28).Get(_50_v27).(bool)
						}
						return _16_v0
					})()).(_dafny.MultiSet)
				}
				return _48_v25
			})())
			_ = _rhs2
			var _rhs3 _dafny.Map = _17_v1
			_ = _rhs3
			var _rhs4 _dafny.Int = (((_54_v32).Intersection(_54_v32)).Union(_54_v32)).Cardinality()
			_ = _rhs4
			var _lhs1 _dafny.Array = _44_v23
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_44_v23), 0))
			_ = _lhs2
			var _lhs3 *GlobalState = _28_globalState
			_ = _lhs3
			(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
			_17_v1 = _rhs3
			_lhs3.F5 = _rhs4
			var _56_v33 _dafny.Array
			_ = _56_v33
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw5
			_56_v33 = _nw5
			var _rhs5 _dafny.Array = (func() _dafny.Array {
				if (_42_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.IntOfUint32((_42_v21).Cardinality()))).Uint32()).(bool) {
					return _56_v33
				}
				return _56_v33
			})()
			_ = _rhs5
			var _rhs6 _dafny.Int = (_20_v4).Minus(_20_v4)
			_ = _rhs6
			var _lhs4 *GlobalState = _28_globalState
			_ = _lhs4
			_56_v33 = _rhs5
			_lhs4.F17 = _rhs6
			var _57_v34 *C2
			_ = _57_v34
			var _nw6 *C2 = New_C2_()
			_ = _nw6
			_nw6.Ctor__(_20_v4)
			_57_v34 = _nw6
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_56_v33), 0))
			_ = _index2
			(_56_v33).ArraySet1((func() _dafny.Array {
				if _16_v0 {
					return _50_v27
				}
				return _50_v27
			})(), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_56_v33), 0))
			_ = _index3
			var _rhs7 bool = !(_16_v0)
			_ = _rhs7
			var _rhs8 bool = _16_v0
			_ = _rhs8
			var _rhs9 _dafny.Map = _17_v1
			_ = _rhs9
			var _rhs10 _dafny.Array = _50_v27
			_ = _rhs10
			var _lhs5 *GlobalState = _28_globalState
			_ = _lhs5
			var _lhs6 *GlobalState = _28_globalState
			_ = _lhs6
			var _lhs7 *GlobalState = _28_globalState
			_ = _lhs7
			var _lhs8 _dafny.Array = _56_v33
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_56_v33), 0))
			_ = _lhs9
			_lhs5.F1 = _rhs7
			_lhs6.F1 = _rhs8
			_lhs7.F0 = _rhs9
			(_lhs8).ArraySet1(_rhs10, (_lhs9).Int())
			var _58_v35 _dafny.Array
			_ = _58_v35
			var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw7
			_58_v35 = _nw7
			var _59_v36 _dafny.Set
			_ = _59_v36
			_59_v36 = _dafny.SetOf((_dafny.MultiSetOf(_58_v35)).Cardinality(), (_21_v5).Cardinality())
			var _60_v37 bool
			_ = _60_v37
			var _out0 bool
			_ = _out0
			_out0 = (_57_v34).M7(_19_v3, _16_v0, _16_v0, Companion_Default___.Fm29((_59_v36).Cardinality(), _20_v4, _16_v0, _28_globalState), _28_globalState)
			_60_v37 = _out0
		} else {
			var _61_v38 _dafny.CodePoint
			_ = _61_v38
			_61_v38 = _dafny.CodePoint('j')
			var _rhs11 _dafny.CodePoint = _61_v38
			_ = _rhs11
			var _rhs12 _dafny.Int = _20_v4
			_ = _rhs12
			var _lhs10 *GlobalState = _28_globalState
			_ = _lhs10
			_61_v38 = _rhs11
			_lhs10.F14 = _rhs12
			var _62_v39 _dafny.Set
			_ = _62_v39
			_62_v39 = _dafny.SetOf(_19_v3)
			var _63_v40 *C0
			_ = _63_v40
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__((_62_v39).Union(_62_v39), _20_v4, _20_v4)
			_63_v40 = _nw8
			var _64_v41 *C8
			_ = _64_v41
			var _nw9 *C8 = New_C8_()
			_ = _nw9
			_nw9.Ctor__(_20_v4, _16_v0, _16_v0)
			_64_v41 = _nw9
			_64_v41 = _64_v41
			var _65_v42 _dafny.Array
			_ = _65_v42
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
			_ = _nw10
			_65_v42 = _nw10
			var _66_v43 T1
			_ = _66_v43
			var _nw11 *C4 = New_C4_()
			_ = _nw11
			_nw11.Ctor__(_16_v0, _16_v0, _16_v0)
			_66_v43 = _nw11
			var _67_v44 _dafny.Map
			_ = _67_v44
			_67_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v43, _dafny.IntOfUint32((_42_v21).Cardinality()))
			var _68_v45 _dafny.Map
			_ = _68_v45
			_68_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v40, _67_v44)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_65_v42), 0))
			_ = _index4
			(_65_v42).ArraySet1((func() _dafny.Map {
				if _16_v0 {
					return _68_v45
				}
				return _68_v45
			})(), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_65_v42), 0))
			_ = _index5
			(_65_v42).ArraySet1((_68_v45).Merge((_68_v45).Merge((_68_v45).Update(_63_v40, _67_v44))), (_index5).Int())
			var _69_v46 _dafny.Array
			_ = _69_v46
			var _nw12 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw12
			_69_v46 = _nw12
			var _70_v47 _dafny.Set
			_ = _70_v47
			_70_v47 = _dafny.SetOf(_69_v46)
			var _71_v48 _dafny.Array
			_ = _71_v48
			var _nw13 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
			_ = _nw13
			_71_v48 = _nw13
			var _72_v49 *C4
			_ = _72_v49
			var _nw14 *C4 = New_C4_()
			_ = _nw14
			_nw14.Ctor__(_66_v43.F24(), _66_v43.F24(), _16_v0)
			_72_v49 = _nw14
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_71_v48), 0))
			_ = _index6
			(_71_v48).ArraySet1(_72_v49, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_71_v48), 0))
			_ = _index7
			var _rhs13 bool = ((func() _dafny.Int {
				if (_66_v43).F25() {
					return (_dafny.SetOf(_63_v40.F31)).Cardinality()
				}
				return _63_v40.F31
			})()).Cmp(_20_v4) < 0
			_ = _rhs13
			var _rhs14 _dafny.Set = _70_v47
			_ = _rhs14
			var _rhs15 _dafny.Int = _63_v40.F31
			_ = _rhs15
			var _rhs16 *C4 = _72_v49
			_ = _rhs16
			var _lhs11 *GlobalState = _28_globalState
			_ = _lhs11
			var _lhs12 *C0 = _63_v40
			_ = _lhs12
			var _lhs13 _dafny.Array = _71_v48
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_71_v48), 0))
			_ = _lhs14
			_lhs11.F15 = _rhs13
			_70_v47 = _rhs14
			_lhs12.F31 = _rhs15
			(_lhs13).ArraySet1(_rhs16, (_lhs14).Int())
		}
		var _73_v50 _dafny.Set
		_ = _73_v50
		_73_v50 = _dafny.SetOf(_19_v3, _19_v3, _19_v3)
		var _74_v51 *C0
		_ = _74_v51
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__(_73_v50, _20_v4, Companion_Default___.SafeModuloInt(_20_v4, _20_v4))
		_74_v51 = _nw15
		var _75_v52 _dafny.Map
		_ = _75_v52
		_75_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v7).Select((Companion_Default___.SafeIndex(_74_v51.F31, _dafny.IntOfUint32((_23_v7).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.SeqOf(false, _16_v0))
		(_28_globalState).F8 = (_dafny.Zero).Minus((_74_v51.F31).Times(((_dafny.Zero).Minus((_75_v52).Cardinality())).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ja")).Cardinality()))))
	} else {
		var _76_v53 D3
		_ = _76_v53
		_76_v53 = Companion_D3_.Create_DC10_(_16_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v0, _16_v0))
		var _77_v54 _dafny.Sequence
		_ = _77_v54
		_77_v54 = _dafny.SeqOf(!((_20_v4).Cmp(_20_v4) <= 0), (_76_v53).Dtor_cf16(), _16_v0)
		_77_v54 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_77_v54, (Companion_Default___.SafeIndex(_20_v4, _dafny.IntOfUint32((_77_v54).Cardinality()))).Uint32(), _16_v0), _77_v54)
		var _78_v55 _dafny.Set
		_ = _78_v55
		_78_v55 = _dafny.SetOf((_dafny.IntOfUint32((_23_v7).Cardinality())).Cmp(_20_v4) == 0, !(_16_v0))
		var _79_v56 _dafny.MultiSet
		_ = _79_v56
		_79_v56 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(626))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_80_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))
		var _rhs17 _dafny.Array = _19_v3
		_ = _rhs17
		var _rhs18 _dafny.Set = _78_v55
		_ = _rhs18
		var _rhs19 _dafny.Int = ((((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(689))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_81_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})))).Update(_23_v7, Companion_Default___.Abs(_20_v4))).Union(_79_v56)).Difference(_79_v56)).Cardinality()
		_ = _rhs19
		_19_v3 = _rhs17
		_78_v55 = _rhs18
		_20_v4 = _rhs19
		var _82_v57 *C3
		_ = _82_v57
		var _nw16 *C3 = New_C3_()
		_ = _nw16
		_nw16.Ctor__(_16_v0, (func() _dafny.Int {
			if true {
				return _dafny.IntOfInt64(-487)
			}
			return _20_v4
		})(), !(!(_16_v0)), !((_20_v4).Cmp(_dafny.IntOfInt64(-148)) > 0), _20_v4)
		_82_v57 = _nw16
		var _83_v58 _dafny.Set
		_ = _83_v58
		_83_v58 = _dafny.SetOf(_19_v3)
		var _84_v59 _dafny.MultiSet
		_ = _84_v59
		_84_v59 = _dafny.MultiSetOf(!(_16_v0), _16_v0)
		var _85_v60 *C0
		_ = _85_v60
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__((_dafny.SetOf(_19_v3)).Union(_83_v58), _20_v4, (_20_v4).Plus((_84_v59).Cardinality()))
		_85_v60 = _nw17
		var _86_v61 *C0
		_ = _86_v61
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(_dafny.SetOf(_19_v3, _19_v3), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v3, (_82_v57).F28())).Cardinality(), (_82_v57).F29())
		_86_v61 = _nw18
	}
	var _87_v62 T1
	_ = _87_v62
	var _nw19 *C4 = New_C4_()
	_ = _nw19
	_nw19.Ctor__(_16_v0, _16_v0, _16_v0)
	_87_v62 = _nw19
	var _88_v63 _dafny.Array
	_ = _88_v63
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(24))
	_ = _nw20
	_88_v63 = _nw20
	var _89_v64 _dafny.Map
	_ = _89_v64
	_89_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v62, _88_v63)
	var _90_v65 _dafny.Map
	_ = _90_v65
	_90_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
		if (_89_v64).Contains(_87_v62) {
			return (_89_v64).Get(_87_v62).(_dafny.Array)
		}
		return _88_v63
	})(), (_dafny.IntOfInt64(22)).Cmp(_20_v4) == 0)
	_90_v65 = (_90_v65).Update(_88_v63, _16_v0)
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _index8
	(_19_v3).ArraySet1(_16_v0, (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _index9
	(_19_v3).ArraySet1(_87_v62.F24(), (_index9).Int())
	(_28_globalState).F1 = !(Companion_Default___.Fm34(_20_v4, _28_globalState)).Contains((_20_v4).Minus(_dafny.IntOfInt64(660)))
	var _91_v66 _dafny.Array
	_ = _91_v66
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
	_ = _nw21
	_91_v66 = _nw21
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
	_ = _index10
	(_91_v66).ArraySet1(_20_v4, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
	_ = _index11
	var _rhs20 _dafny.Int = _20_v4
	_ = _rhs20
	var _rhs21 bool = !((_87_v62).F25())
	_ = _rhs21
	var _rhs22 bool = _16_v0
	_ = _rhs22
	var _lhs15 _dafny.Array = _91_v66
	_ = _lhs15
	var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
	_ = _lhs16
	var _lhs17 *GlobalState = _28_globalState
	_ = _lhs17
	var _lhs18 *GlobalState = _28_globalState
	_ = _lhs18
	(_lhs15).ArraySet1(_rhs20, (_lhs16).Int())
	_lhs17.F9 = _rhs21
	_lhs18.F1 = _rhs22
	var _92_v67 _dafny.Array
	_ = _92_v67
	var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
	_ = _nw22
	_92_v67 = _nw22
	var _93_v68 _dafny.Array
	_ = _93_v68
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
	_ = _nw23
	_93_v68 = _nw23
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_92_v67), 0))
	_ = _index12
	(_92_v67).ArraySet1(_93_v68, (_index12).Int())
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_92_v67), 0))
	_ = _index13
	(_92_v67).ArraySet1(_93_v68, (_index13).Int())
	var _94_v69 _dafny.Set
	_ = _94_v69
	_94_v69 = _dafny.SetOf(_19_v3, _19_v3, _19_v3)
	var _95_v70 T0
	_ = _95_v70
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__(_94_v69, (_20_v4).Times((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)), _20_v4)
	_95_v70 = _nw24
	var _96_v71 _dafny.MultiSet
	_ = _96_v71
	_96_v71 = _dafny.MultiSetOf(_87_v62.F24())
	var _97_v72 _dafny.Map
	_ = _97_v72
	_97_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v71, _19_v3)
	var _rhs23 T0 = _95_v70
	_ = _rhs23
	var _rhs24 _dafny.Int = (_dafny.Zero).Minus((_95_v70).F23())
	_ = _rhs24
	var _rhs25 _dafny.Map = _97_v72
	_ = _rhs25
	var _lhs19 *GlobalState = _28_globalState
	_ = _lhs19
	_95_v70 = _rhs23
	_lhs19.F8 = _rhs24
	_97_v72 = _rhs25
	var _98_v73 _dafny.Map
	_ = _98_v73
	_98_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _16_v0)
	_98_v73 = (_98_v73).Update((_95_v70).F23(), (func() bool {
		if (_87_v62).F25() {
			return !(_87_v62.F24())
		}
		return (_87_v62).F25()
	})())
	var _99_v74 _dafny.Sequence
	_ = _99_v74
	_99_v74 = _dafny.SeqOf(_87_v62.F24())
	var _100_v75 _dafny.Map
	_ = _100_v75
	_100_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_99_v74).Cardinality()), _23_v7)
	_23_v7 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_100_v75).Contains(_dafny.IntOfInt64(456)) {
			return (_100_v75).Get(_dafny.IntOfInt64(456)).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("gstyka")
	})(), _dafny.Companion_Sequence_.Concatenate(_23_v7, _23_v7))
	var _101_v76 _dafny.Map
	_ = _101_v76
	_101_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool), (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool))
	var _102_v77 _dafny.Map
	_ = _102_v77
	_102_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v76, (_21_v5).Cardinality())
	var _103_v78 _dafny.Map
	_ = _103_v78
	_103_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v77, (_87_v62).F25())
	var _104_v79 *C1
	_ = _104_v79
	var _nw25 *C1 = New_C1_()
	_ = _nw25
	_nw25.Ctor__(true, (func() bool {
		if (_103_v78).Contains(_102_v77) {
			return (_103_v78).Get(_102_v77).(bool)
		}
		return (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool)
	})())
	_104_v79 = _nw25
	var _105_i5 _dafny.Int
	_ = _105_i5
	_105_i5 = _dafny.Zero
	{
		for (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool) {
			{
				if (_105_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_105_i5 = (_105_i5).Plus(_dafny.One)
				var _106_v80 *C0
				_ = _106_v80
				var _nw26 *C0 = New_C0_()
				_ = _nw26
				_nw26.Ctor__(_94_v69, (_95_v70).F23(), (_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int))
				_106_v80 = _nw26
				var _107_v81 _dafny.Map
				_ = _107_v81
				_107_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_95_v70).F23(), (_dafny.MultiSetFromSeq(_99_v74)).Cardinality()), _106_v80)
				_107_v81 = (_107_v81).Update(Companion_Default___.SafeDivisionInt((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _20_v4), _106_v80)
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
				_ = _index14
				(_19_v3).ArraySet1(false, (_index14).Int())
				var _108_v82 _dafny.MultiSet
				_ = _108_v82
				_108_v82 = _dafny.MultiSetOf((_95_v70).F23(), _106_v80.F31, _dafny.IntOfUint32((_23_v7).Cardinality()))
				var _109_v83 *C7
				_ = _109_v83
				var _nw27 *C7 = New_C7_()
				_ = _nw27
				_nw27.Ctor__()
				_109_v83 = _nw27
				var _110_v84 _dafny.Map
				_ = _110_v84
				_110_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v83, (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool))
				(_28_globalState).F9 = !(_108_v82).Equals(_dafny.MultiSetOf((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_29_v9, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_21_v5).Contains(_106_v80.F31) {
						return (_21_v5).Get(_106_v80.F31).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_21_v5).Contains((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)) {
							return (_21_v5).Get((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)).(_dafny.Int)
						}
						return _106_v80.F31
					})()
				})(), _dafny.IntOfUint32((_29_v9).Cardinality()))).Uint32(), (_110_v84).Cardinality())).Cardinality())))
				if ((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)).Cmp((_95_v70).F23()) >= 0 {
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
					_ = _index15
					(_91_v66).ArraySet1(((func() _dafny.Int {
						if (_87_v62).F25() {
							return _20_v4
						}
						return (_95_v70).F23()
					})()).Minus((_95_v70).F23()), (_index15).Int())
					var _111_v85 _dafny.CodePoint
					_ = _111_v85
					_111_v85 = _dafny.CodePoint('a')
					_111_v85 = _111_v85
					(_28_globalState).F17 = _dafny.IntOfInt64(297)
					var _112_v86 _dafny.Map
					_ = _112_v86
					_112_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v76, _87_v62)
					var _113_v87 _dafny.Map
					_ = _113_v87
					_113_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_87_v62).F25(), _87_v62)
					var _114_v88 _dafny.Array
					_ = _114_v88
					var _nwElement0_0 T1 = (func() T1 {
						if !((_87_v62).F25()) {
							return _87_v62
						}
						return _87_v62
					})()
					_ = _nwElement0_0
					var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(23))
					_ = _nw28
					(_nw28).ArraySet1(_nwElement0_0, 0)
					(_nw28).ArraySet1(_87_v62, 1)
					(_nw28).ArraySet1(_87_v62, 2)
					(_nw28).ArraySet1(_87_v62, 3)
					(_nw28).ArraySet1(_87_v62, 4)
					(_nw28).ArraySet1(_87_v62, 5)
					(_nw28).ArraySet1(_87_v62, 6)
					(_nw28).ArraySet1(_87_v62, 7)
					(_nw28).ArraySet1(_87_v62, 8)
					(_nw28).ArraySet1(_87_v62, 9)
					(_nw28).ArraySet1(_87_v62, 10)
					(_nw28).ArraySet1(_87_v62, 11)
					(_nw28).ArraySet1(_87_v62, 12)
					(_nw28).ArraySet1(_87_v62, 13)
					(_nw28).ArraySet1((func() T1 {
						if (_87_v62).F25() {
							return _87_v62
						}
						return (func() T1 {
							if (_112_v86).Contains(_101_v76) {
								return (_112_v86).Get(_101_v76).(T1)
							}
							return (func() T1 {
								if (_113_v87).Contains(_16_v0) {
									return (_113_v87).Get(_16_v0).(T1)
								}
								return _87_v62
							})()
						})()
					})(), 14)
					(_nw28).ArraySet1((func() T1 {
						if _87_v62.F24() {
							return _87_v62
						}
						return _87_v62
					})(), 15)
					(_nw28).ArraySet1(_87_v62, 16)
					(_nw28).ArraySet1(_87_v62, 17)
					(_nw28).ArraySet1(_87_v62, 18)
					(_nw28).ArraySet1(_87_v62, 19)
					(_nw28).ArraySet1(_87_v62, 20)
					(_nw28).ArraySet1(_87_v62, 21)
					(_nw28).ArraySet1(_87_v62, 22)
					_114_v88 = _nw28
					_114_v88 = _114_v88
					var _115_v91 _dafny.MultiSet
					_ = _115_v91
					_115_v91 = _dafny.MultiSetOf(_98_v73, (func() _dafny.Map {
						var _coll8 = _dafny.NewMapBuilder()
						_ = _coll8
						for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(391), _dafny.IntOfInt64(463))); ; {
							_compr_8, _ok8 := _iter8()
							if !_ok8 {
								break
							}
							var _116_v89 _dafny.Int
							_116_v89 = interface{}(_compr_8).(_dafny.Int)
							if ((_dafny.IntOfInt64(391)).Cmp(_116_v89) <= 0) && ((_116_v89).Cmp(_dafny.IntOfInt64(463)) < 0) {
								_coll8.Add(Companion_Default___.SafeDivisionInt(_116_v89, (_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)), (_87_v62).F25())
							}
						}
						return _coll8.ToMap()
					}()).Merge(func() _dafny.Map {
						var _coll9 = _dafny.NewMapBuilder()
						_ = _coll9
						for _iter9 := _dafny.Iterate((_98_v73).Keys().Elements()); ; {
							_compr_9, _ok9 := _iter9()
							if !_ok9 {
								break
							}
							var _117_v90 _dafny.Int
							_117_v90 = interface{}(_compr_9).(_dafny.Int)
							if (_98_v73).Contains(_117_v90) {
								_coll9.Add(Companion_Default___.SafeDivisionInt(_117_v90, _20_v4), true)
							}
						}
						return _coll9.ToMap()
					}()), _98_v73, _98_v73, _98_v73)
					var _rhs26 _dafny.MultiSet = _115_v91
					_ = _rhs26
					var _rhs27 bool = _16_v0
					_ = _rhs27
					var _lhs20 *GlobalState = _28_globalState
					_ = _lhs20
					_115_v91 = _rhs26
					_lhs20.F1 = _rhs27
				} else {
					var _118_v92 *C3
					_ = _118_v92
					var _nw29 *C3 = New_C3_()
					_ = _nw29
					_nw29.Ctor__((_87_v62).F25(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg10 _dafny.Int) interface{} {
							return coer10(arg10)
						}
					}((func(_119_v62 T1) func(_dafny.Int) _dafny.Map {
						return func(_120_i6 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (_119_v62).F25())
						}
					})(_87_v62)))).Cardinality()), (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool), _87_v62.F24(), _106_v80.F31)
					_118_v92 = _nw29
					var _121_v93 _dafny.Sequence
					_ = _121_v93
					_121_v93 = _dafny.SeqOf(_118_v92, _118_v92)
					_121_v93 = _dafny.Companion_Sequence_.Concatenate(_121_v93, _dafny.Companion_Sequence_.Concatenate(_121_v93, _121_v93))
					var _122_v94 _dafny.Map
					_ = _122_v94
					_122_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v80, _96_v71)
					var _rhs28 bool = _87_v62.F24()
					_ = _rhs28
					var _rhs29 _dafny.Map = ((_122_v94).Merge(_122_v94)).Merge(_122_v94)
					_ = _rhs29
					var _lhs21 T1 = _87_v62
					_ = _lhs21
					_lhs21.F24_set_(_rhs28)
					_122_v94 = _rhs29
					var _123_v95 *C3
					_ = _123_v95
					_123_v95 = _118_v92
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
					_ = _index16
					var _rhs30 bool = (_87_v62).F25()
					_ = _rhs30
					var _rhs31 bool = false
					_ = _rhs31
					var _rhs32 *C3 = (_123_v95)
					_ = _rhs32
					var _rhs33 _dafny.Array = _19_v3
					_ = _rhs33
					var _lhs22 _dafny.Array = _19_v3
					_ = _lhs22
					var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
					_ = _lhs23
					var _lhs24 *GlobalState = _28_globalState
					_ = _lhs24
					(_lhs22).ArraySet1(_rhs30, (_lhs23).Int())
					_lhs24.F6 = _rhs31
					_118_v92 = _rhs32
					_19_v3 = _rhs33
					var _124_v96 bool
					_ = _124_v96
					var _125_v97 T0
					_ = _125_v97
					var _126_v98 bool
					_ = _126_v98
					var _127_v99 _dafny.Int
					_ = _127_v99
					var _out1 bool
					_ = _out1
					var _out2 T0
					_ = _out2
					var _out3 bool
					_ = _out3
					var _out4 _dafny.Int
					_ = _out4
					_out1, _out2, _out3, _out4 = (_104_v79).M1(_28_globalState)
					_124_v96 = _out1
					_125_v97 = _out2
					_126_v98 = _out3
					_127_v99 = _out4
					_19_v3 = _19_v3
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _128_v100 _dafny.Array
	_ = _128_v100
	var _nw30 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
	_ = _nw30
	_128_v100 = _nw30
	var _129_v101 _dafny.Sequence
	_ = _129_v101
	_129_v101 = _dafny.SeqOf(_128_v100)
	var _130_v102 _dafny.Sequence
	_ = _130_v102
	_130_v102 = _dafny.SeqOf((_129_v101).Select((Companion_Default___.SafeIndex((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_129_v101).Cardinality()))).Uint32()).(_dafny.Array))
	var _131_v103 D3
	_ = _131_v103
	_131_v103 = Companion_D3_.Create_DC10_((_87_v62).F25(), _101_v76)
	var _pat_let_tv0 = _101_v76
	_ = _pat_let_tv0
	var _132_v104 _dafny.Sequence
	_ = _132_v104
	_132_v104 = _dafny.SeqOf(func(_pat_let0_0 D3) D3 {
		return func(_133_dt__update__tmp_h0 D3) D3 {
			return func(_pat_let1_0 _dafny.Map) D3 {
				return func(_134_dt__update_hcf17_h0 _dafny.Map) D3 {
					return Companion_D3_.Create_DC10_((_133_dt__update__tmp_h0).Dtor_cf16(), _134_dt__update_hcf17_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(Companion_D3_.Create_DC10_((_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool), _101_v76)), _131_v103)
	var _135_v105 _dafny.Sequence
	_ = _135_v105
	_135_v105 = _dafny.SeqOf(_87_v62, _87_v62, _87_v62)
	var _136_v106 _dafny.Set
	_ = _136_v106
	_136_v106 = _dafny.SetOf(_135_v105)
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _index17
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
	_ = _index18
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _index19
	var _rhs34 bool = ((_95_v70).F23()).Cmp(_dafny.IntOfUint32((_132_v104).Cardinality())) > 0
	_ = _rhs34
	var _rhs35 _dafny.Int = ((func() _dafny.Set {
		if (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool) {
			return _136_v106
		}
		return _136_v106
	})()).Cardinality()
	_ = _rhs35
	var _rhs36 bool = _16_v0
	_ = _rhs36
	var _rhs37 _dafny.Int = _20_v4
	_ = _rhs37
	var _rhs38 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_130_v102, _129_v101)
	_ = _rhs38
	var _lhs25 _dafny.Array = _19_v3
	_ = _lhs25
	var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _lhs26
	var _lhs27 _dafny.Array = _91_v66
	_ = _lhs27
	var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
	_ = _lhs28
	var _lhs29 _dafny.Array = _19_v3
	_ = _lhs29
	var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
	_ = _lhs30
	var _lhs31 *GlobalState = _28_globalState
	_ = _lhs31
	(_lhs25).ArraySet1(_rhs34, (_lhs26).Int())
	(_lhs27).ArraySet1(_rhs35, (_lhs28).Int())
	(_lhs29).ArraySet1(_rhs36, (_lhs30).Int())
	_lhs31.F8 = _rhs37
	_129_v101 = _rhs38
	var _137_v107 _dafny.CodePoint
	_ = _137_v107
	_137_v107 = _dafny.CodePoint('d')
	var _138_v108 _dafny.Map
	_ = _138_v108
	_138_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool), _137_v107)
	var _139_v109 _dafny.Map
	_ = _139_v109
	_139_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v108, _16_v0)
	(_104_v79).M8(!((func() bool {
		if (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool) {
			return (_87_v62).F25()
		}
		return !(_16_v0)
	})()), (_95_v70).Fm2(_28_globalState), _139_v109, false, _28_globalState)
	if _16_v0 {
		if !_dafny.Companion_Sequence_.Contains(_23_v7, _137_v107) {
			var _140_v110 bool
			_ = _140_v110
			var _141_v111 T0
			_ = _141_v111
			var _142_v112 bool
			_ = _142_v112
			var _143_v113 _dafny.Int
			_ = _143_v113
			var _out5 bool
			_ = _out5
			var _out6 T0
			_ = _out6
			var _out7 bool
			_ = _out7
			var _out8 _dafny.Int
			_ = _out8
			_out5, _out6, _out7, _out8 = (_87_v62).M1(_28_globalState)
			_140_v110 = _out5
			_141_v111 = _out6
			_142_v112 = _out7
			_143_v113 = _out8
			var _144_v114 *C5
			_ = _144_v114
			var _nw31 *C5 = New_C5_()
			_ = _nw31
			_nw31.Ctor__(_dafny.Companion_Sequence_.Equal(_23_v7, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm25((_87_v62).F25(), (_87_v62).F25(), _137_v107, _28_globalState), (Companion_Default___.SafeIndex((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm25((_87_v62).F25(), (_87_v62).F25(), _137_v107, _28_globalState)).Cardinality()))).Uint32(), _dafny.CodePoint('t'))), _140_v110)
			_144_v114 = _nw31
			var _145_v115 *C0
			_ = _145_v115
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(_94_v69, (_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), (_95_v70).F23())
			_145_v115 = _nw32
			_145_v115 = _145_v115
			var _146_v116 _dafny.Sequence
			_ = _146_v116
			_146_v116 = _dafny.SeqOf(_91_v66)
			_91_v66 = (_146_v116).Select((Companion_Default___.SafeIndex((_20_v4).Times((_29_v9).Select((Companion_Default___.SafeIndex((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_29_v9).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_146_v116).Cardinality()))).Uint32()).(_dafny.Array)
			var _147_v117 _dafny.Sequence
			_ = _147_v117
			_147_v117 = _dafny.SeqOf(_135_v105)
			var _148_v118 _dafny.MultiSet
			_ = _148_v118
			_148_v118 = _dafny.MultiSetOf((_141_v111).F23(), _145_v115.F31)
			var _149_v119 _dafny.Map
			_ = _149_v119
			_149_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_143_v113, _20_v4), (_147_v117).Select((Companion_Default___.SafeIndex((_148_v118).Cardinality(), _dafny.IntOfUint32((_147_v117).Cardinality()))).Uint32()).(_dafny.Sequence))
			_149_v119 = (_149_v119).Update(_20_v4, _dafny.Companion_Sequence_.Update(_135_v105, (Companion_Default___.SafeIndex((_141_v111).F23(), _dafny.IntOfUint32((_135_v105).Cardinality()))).Uint32(), _87_v62))
		} else {
			_23_v7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_23_v7, _23_v7), _23_v7)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
			_ = _index20
			(_19_v3).ArraySet1((_87_v62).F25(), (_index20).Int())
			var _150_v120 *C0
			_ = _150_v120
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__(_94_v69, (_95_v70).F23(), (_95_v70).F23())
			_150_v120 = _nw33
			var _rhs39 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_95_v70).F23(), (_95_v70).F23()), (_95_v70).F23())
			_ = _rhs39
			var _rhs40 *C0 = _150_v120
			_ = _rhs40
			var _lhs32 *GlobalState = _28_globalState
			_ = _lhs32
			_lhs32.F8 = _rhs39
			_150_v120 = _rhs40
			var _151_v121 _dafny.Set
			_ = _151_v121
			_151_v121 = _dafny.SetOf(!(_87_v62.F24()), (func() bool {
				if _87_v62.F24() {
					return (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool)
				}
				return !(!(!((_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool))))
			})())
			_151_v121 = _151_v121
			var _152_v122 _dafny.MultiSet
			_ = _152_v122
			_152_v122 = _dafny.MultiSetOf(_23_v7, _dafny.UnicodeSeqOfUtf8Bytes("guaqmlf"))
			_16_v0 = !((((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_23_v7).Cardinality()))).Cmp((func() _dafny.Int {
				if (_152_v122).Contains(_23_v7) {
					return (_152_v122).Multiplicity(_23_v7)
				}
				return (_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)
			})()) == 0)
		}
		_104_v79 = _104_v79
		var _153_v123 _dafny.Map
		_ = _153_v123
		_153_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_((_87_v62).F25(), _101_v76), _17_v1)
		_153_v123 = (func() _dafny.Map {
			if _87_v62.F24() {
				return (_153_v123).Merge(_153_v123)
			}
			return _153_v123
		})()
		var _154_v124 _dafny.Set
		_ = _154_v124
		_154_v124 = _dafny.SetOf((_87_v62).F25(), (_87_v62).F25())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))
		_ = _index21
		(_91_v66).ArraySet1((_dafny.Zero).Minus((_154_v124).Cardinality()), (_index21).Int())
		_16_v0 = (_154_v124).IsProperSubsetOf(_154_v124)
	} else {
		(_28_globalState).F5 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_23_v7).Cardinality()))
		var _155_v125 _dafny.Array
		_ = _155_v125
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw34
		_155_v125 = _nw34
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_155_v125), 0))
		_ = _index22
		(_155_v125).ArraySet1(_19_v3, (_index22).Int())
		var _156_v126 _dafny.Sequence
		_ = _156_v126
		_156_v126 = _dafny.SeqOf(_19_v3)
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_155_v125), 0))
		_ = _index23
		(_155_v125).ArraySet1((func() _dafny.Array {
			if true {
				return (func() _dafny.Array {
					if _16_v0 {
						return _19_v3
					}
					return _19_v3
				})()
			}
			return (_156_v126).Select((Companion_Default___.SafeIndex((_95_v70).F23(), _dafny.IntOfUint32((_156_v126).Cardinality()))).Uint32()).(_dafny.Array)
		})(), (_index23).Int())
		var _157_v127 *C7
		_ = _157_v127
		var _nw35 *C7 = New_C7_()
		_ = _nw35
		_nw35.Ctor__()
		_157_v127 = _nw35
		var _158_v128 _dafny.Sequence
		_ = _158_v128
		_158_v128 = _dafny.SeqOf(_157_v127)
		_157_v127 = (_158_v128).Select((Companion_Default___.SafeIndex((_91_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_91_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_158_v128).Cardinality()))).Uint32()).(*C7)
		var _159_v129 _dafny.Map
		_ = _159_v129
		_159_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool), _94_v69)
		var _160_v130 *C0
		_ = _160_v130
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__((func() _dafny.Set {
			if (_159_v129).Contains((_87_v62).F25()) {
				return (_159_v129).Get((_87_v62).F25()).(_dafny.Set)
			}
			return _94_v69
		})(), (_95_v70).Fm2(_28_globalState), (_95_v70).F23())
		_160_v130 = _nw36
		_160_v130 = _160_v130
		var _161_v131 _dafny.Sequence
		_ = _161_v131
		_161_v131 = _dafny.SeqOf(_91_v66, _91_v66)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
		_ = _index24
		var _rhs41 _dafny.Int = (Companion_Default___.SafeModuloInt((_95_v70).F23(), _160_v130.F31)).Plus((_95_v70).F23())
		_ = _rhs41
		var _rhs42 bool = (_19_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))).Int()).(bool)
		_ = _rhs42
		var _rhs43 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_161_v131, _dafny.SeqOf(_91_v66)), _161_v131)
		_ = _rhs43
		var _lhs33 *GlobalState = _28_globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = _19_v3
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_19_v3), 0))
		_ = _lhs35
		var _lhs36 *GlobalState = _28_globalState
		_ = _lhs36
		_lhs33.F8 = _rhs41
		(_lhs34).ArraySet1(_rhs42, (_lhs35).Int())
		_lhs36.F9 = _rhs43
	}
	_dafny.Print(_16_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_17_v1).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(313))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_18_v2, _dafny.SeqOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v3).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v3).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_20_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_21_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(333), _dafny.IntOfInt64(333))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_22_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_23_v7.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_24_v8).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState.F0).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(313))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_28_globalState.F4, _dafny.SeqOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_28_globalState).F10()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F11().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_28_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_28_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(333), _dafny.IntOfInt64(-333), _dafny.IntOfInt64(2), _dafny.IntOfInt64(837))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_28_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_29_v9, _dafny.SeqOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_v62.F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_v62).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_89_v64).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v65).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v66).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v69).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_95_v70).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v71).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v72).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v73).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(333), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_99_v74, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v75).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.UnicodeSeqOfUtf8Bytes("hllu"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v76).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v77).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v78).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.One), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_129_v101).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_130_v102).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v103).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_v103).Dtor_cf17()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_132_v104, _dafny.SeqOf(Companion_D3_.Create_DC10_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), Companion_D3_.Create_DC10_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_135_v105).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v106).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_v107)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v108).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v109).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d')), true)))
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
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
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

type D0_DC2 struct {
	Cf3 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf3 D0) D0 {
	return D0{D0_DC2{Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
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
	Cf5 bool
	Cf6 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 bool, Cf6 bool) D1 {
	return D1{D1_DC4{Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf7  bool
	Cf8  _dafny.Map
	Cf9  _dafny.Int
	Cf10 _dafny.Int
	Cf11 _dafny.Map
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 bool, Cf8 _dafny.Map, Cf9 _dafny.Int, Cf10 _dafny.Int, Cf11 _dafny.Map) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC6 struct {
	Cf12 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf12 D1) D1 {
	return D1{D1_DC6{Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, false)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf12() D1 {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Equals(data2.Cf11)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf13 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 _dafny.Array) D2 {
	return D2{D2_DC7{Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf13 == data2.Cf13
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
	Cf15 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.Int) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf16 bool
	Cf17 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 bool, Cf17 _dafny.Map) D3 {
	return D3{D3_DC10{Cf16, Cf17}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf14 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero)
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + data.Cf14.VerbatimString(true) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Equals(data2.Cf17)
		}
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

type D4_DC12 struct {
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf19 _dafny.Int, Cf20 _dafny.Int) D4 {
	return D4{D4_DC12{Cf19, Cf20}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf18 _dafny.Map
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf18 _dafny.Map) D4 {
	return D4{D4_DC11{Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(_dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
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
	Cf22 _dafny.Int
	Cf23 bool
	Cf24 _dafny.Set
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf22 _dafny.Int, Cf23 bool, Cf24 _dafny.Set) D5 {
	return D5{D5_DC14{Cf22, Cf23, Cf24}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC15 struct {
	Cf25 _dafny.Int
	Cf26 _dafny.CodePoint
	Cf27 _dafny.Map
	Cf28 _dafny.Map
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf25 _dafny.Int, Cf26 _dafny.CodePoint, Cf27 _dafny.Map, Cf28 _dafny.Map) D5 {
	return D5{D5_DC15{Cf25, Cf26, Cf27, Cf28}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC13 struct {
	Cf21 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf21 _dafny.Map) D5 {
	return D5{D5_DC13{Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.Zero, false, _dafny.EmptySet)
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC14).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Set {
	return _this.Get_().(D5_DC14).Cf24
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC15).Cf25
}

func (_this D5) Dtor_cf26() _dafny.CodePoint {
	return _this.Get_().(D5_DC15).Cf26
}

func (_this D5) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D5_DC15).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) Dtor_cf21() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23 && data1.Cf24.Equals(data2.Cf24)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26 && data1.Cf27.Equals(data2.Cf27) && data1.Cf28.Equals(data2.Cf28)
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
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

type D6_DC17 struct {
	Cf30 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 bool) D6 {
	return D6{D6_DC17{Cf30}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf31 bool
	Cf32 bool
	Cf33 _dafny.CodePoint
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf31 bool, Cf32 bool, Cf33 _dafny.CodePoint) D6 {
	return D6{D6_DC18{Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC16 struct {
	Cf29 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf29 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf29}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false)
}

func (_this D6) Dtor_cf30() bool {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf31() bool {
	return _this.Get_().(D6_DC18).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC18).Cf32
}

func (_this D6) Dtor_cf33() _dafny.CodePoint {
	return _this.Get_().(D6_DC18).Cf33
}

func (_this D6) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30 == data2.Cf30
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	Cf35 bool
	Cf36 _dafny.Map
	Cf37 _dafny.Array
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf35 bool, Cf36 _dafny.Map, Cf37 _dafny.Array) D7 {
	return D7{D7_DC20{Cf35, Cf36, Cf37}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC19 struct {
	Cf34 _dafny.MultiSet
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf34 _dafny.MultiSet) D7 {
	return D7{D7_DC19{Cf34}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC21 struct {
	Cf38 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf38 D7) D7 {
	return D7{D7_DC21{Cf38}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(false, _dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D7) Dtor_cf35() bool {
	return _this.Get_().(D7_DC20).Cf35
}

func (_this D7) Dtor_cf36() _dafny.Map {
	return _this.Get_().(D7_DC20).Cf36
}

func (_this D7) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D7_DC20).Cf37
}

func (_this D7) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D7_DC19).Cf34
}

func (_this D7) Dtor_cf38() D7 {
	return _this.Get_().(D7_DC21).Cf38
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf38) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37 == data2.Cf37
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
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

type D8_DC23 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Set
	Cf42 _dafny.Int
	Cf43 _dafny.Int
	Cf44 _dafny.Int
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf40 _dafny.Int, Cf41 _dafny.Set, Cf42 _dafny.Int, Cf43 _dafny.Int, Cf44 _dafny.Int) D8 {
	return D8{D8_DC23{Cf40, Cf41, Cf42, Cf43, Cf44}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf45 _dafny.Int
	Cf46 _dafny.CodePoint
	Cf47 _dafny.MultiSet
	Cf48 bool
	Cf49 *C0
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf45 _dafny.Int, Cf46 _dafny.CodePoint, Cf47 _dafny.MultiSet, Cf48 bool, Cf49 *C0) D8 {
	return D8{D8_DC24{Cf45, Cf46, Cf47, Cf48, Cf49}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf39 _dafny.Array
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf39 _dafny.Array) D8 {
	return D8{D8_DC22{Cf39}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.Zero, _dafny.EmptySet, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf40
}

func (_this D8) Dtor_cf41() _dafny.Set {
	return _this.Get_().(D8_DC23).Cf41
}

func (_this D8) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf42
}

func (_this D8) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf43
}

func (_this D8) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf44
}

func (_this D8) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf45
}

func (_this D8) Dtor_cf46() _dafny.CodePoint {
	return _this.Get_().(D8_DC24).Cf46
}

func (_this D8) Dtor_cf47() _dafny.MultiSet {
	return _this.Get_().(D8_DC24).Cf47
}

func (_this D8) Dtor_cf48() bool {
	return _this.Get_().(D8_DC24).Cf48
}

func (_this D8) Dtor_cf49() *C0 {
	return _this.Get_().(D8_DC24).Cf49
}

func (_this D8) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf39
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Equals(data2.Cf41) && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47) && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf39 == data2.Cf39
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
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_() D9 {
	return D9{D9_DC26{}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC25 struct {
	Cf50 _dafny.Sequence
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf50 _dafny.Sequence) D9 {
	return D9{D9_DC25{Cf50}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC27 struct {
	Cf51 D9
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf51 D9) D9 {
	return D9{D9_DC27{Cf51}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_()
}

func (_this D9) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D9_DC25).Cf50
}

func (_this D9) Dtor_cf51() D9 {
	return _this.Get_().(D9_DC27).Cf51
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf51) + ")"
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
			_, ok := other.Get_().(D9_DC26)
			return ok
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf50.Equals(data2.Cf50)
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf51.Equals(data2.Cf51)
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
	Cf53 bool
	Cf54 bool
	Cf55 _dafny.Int
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf53 bool, Cf54 bool, Cf55 _dafny.Int) D10 {
	return D10{D10_DC29{Cf53, Cf54, Cf55}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf52 _dafny.Sequence
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf52 _dafny.Sequence) D10 {
	return D10{D10_DC28{Cf52}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(false, false, _dafny.Zero)
}

func (_this D10) Dtor_cf53() bool {
	return _this.Get_().(D10_DC29).Cf53
}

func (_this D10) Dtor_cf54() bool {
	return _this.Get_().(D10_DC29).Cf54
}

func (_this D10) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf55
}

func (_this D10) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D10_DC28).Cf52
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf52) + ")"
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
			return ok && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54 && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf52.Equals(data2.Cf52)
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

type D11_DC31 struct {
	Cf57 _dafny.Array
	Cf58 _dafny.Int
	Cf59 bool
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf57 _dafny.Array, Cf58 _dafny.Int, Cf59 bool) D11 {
	return D11{D11_DC31{Cf57, Cf58, Cf59}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC30 struct {
	Cf56 T0
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf56 T0) D11 {
	return D11{D11_DC30{Cf56}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC31_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, false)
}

func (_this D11) Dtor_cf57() _dafny.Array {
	return _this.Get_().(D11_DC31).Cf57
}

func (_this D11) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf58
}

func (_this D11) Dtor_cf59() bool {
	return _this.Get_().(D11_DC31).Cf59
}

func (_this D11) Dtor_cf56() T0 {
	return _this.Get_().(D11_DC30).Cf56
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0 && data1.Cf59 == data2.Cf59
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && _dafny.AreEqual(data1.Cf56, data2.Cf56)
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
	Cf61 _dafny.Int
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf61 _dafny.Int) D12 {
	return D12{D12_DC33{Cf61}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf60 _dafny.MultiSet
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf60 _dafny.MultiSet) D12 {
	return D12{D12_DC32{Cf60}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_(_dafny.Zero)
}

func (_this D12) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D12_DC33).Cf61
}

func (_this D12) Dtor_cf60() _dafny.MultiSet {
	return _this.Get_().(D12_DC32).Cf60
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf61) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf60) + ")"
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
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf60.Equals(data2.Cf60)
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
	Cf62 _dafny.Set
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf62 _dafny.Set) D13 {
	return D13{D13_DC34{Cf62}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D13) Dtor_cf62() _dafny.Set {
	return _this.Get_().(D13_DC34).Cf62
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf62) + ")"
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
			return ok && data1.Cf62.Equals(data2.Cf62)
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

type D14_DC35 struct {
	Cf63 _dafny.Map
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf63 _dafny.Map) D14 {
	return D14{D14_DC35{Cf63}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D14) Dtor_cf63() _dafny.Map {
	return _this.Get_().(D14_DC35).Cf63
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf63.Equals(data2.Cf63)
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

type D15_DC37 struct {
	Cf65 bool
	Cf66 _dafny.Map
	Cf67 _dafny.Int
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf65 bool, Cf66 _dafny.Map, Cf67 _dafny.Int) D15 {
	return D15{D15_DC37{Cf65, Cf66, Cf67}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

type D15_DC38 struct {
	Cf68 bool
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf68 bool) D15 {
	return D15{D15_DC38{Cf68}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC36 struct {
	Cf64 _dafny.Map
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf64 _dafny.Map) D15 {
	return D15{D15_DC36{Cf64}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC37_(false, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D15) Dtor_cf65() bool {
	return _this.Get_().(D15_DC37).Cf65
}

func (_this D15) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D15_DC37).Cf66
}

func (_this D15) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D15_DC37).Cf67
}

func (_this D15) Dtor_cf68() bool {
	return _this.Get_().(D15_DC38).Cf68
}

func (_this D15) Dtor_cf64() _dafny.Map {
	return _this.Get_().(D15_DC36).Cf64
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf65 == data2.Cf65 && data1.Cf66.Equals(data2.Cf66) && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf68 == data2.Cf68
		}
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf64.Equals(data2.Cf64)
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
	Cf70 _dafny.Int
	Cf71 _dafny.Array
	Cf72 _dafny.Map
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf70 _dafny.Int, Cf71 _dafny.Array, Cf72 _dafny.Map) D16 {
	return D16{D16_DC40{Cf70, Cf71, Cf72}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC39 struct {
	Cf69 _dafny.Array
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf69 _dafny.Array) D16 {
	return D16{D16_DC39{Cf69}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC40_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap)
}

func (_this D16) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf70
}

func (_this D16) Dtor_cf71() _dafny.Array {
	return _this.Get_().(D16_DC40).Cf71
}

func (_this D16) Dtor_cf72() _dafny.Map {
	return _this.Get_().(D16_DC40).Cf72
}

func (_this D16) Dtor_cf69() _dafny.Array {
	return _this.Get_().(D16_DC39).Cf69
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf69) + ")"
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
			return ok && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71 == data2.Cf71 && data1.Cf72.Equals(data2.Cf72)
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf69 == data2.Cf69
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

type D17_DC42 struct {
}

func (D17_DC42) isD17() {}

func (CompanionStruct_D17_) Create_DC42_() D17 {
	return D17{D17_DC42{}}
}

func (_this D17) Is_DC42() bool {
	_, ok := _this.Get_().(D17_DC42)
	return ok
}

type D17_DC41 struct {
	Cf73 *C4
}

func (D17_DC41) isD17() {}

func (CompanionStruct_D17_) Create_DC41_(Cf73 *C4) D17 {
	return D17{D17_DC41{Cf73}}
}

func (_this D17) Is_DC41() bool {
	_, ok := _this.Get_().(D17_DC41)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC42_()
}

func (_this D17) Dtor_cf73() *C4 {
	return _this.Get_().(D17_DC41).Cf73
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC42:
		{
			return "D17.DC42"
		}
	case D17_DC41:
		{
			return "D17.DC41" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC42:
		{
			_, ok := other.Get_().(D17_DC42)
			return ok
		}
	case D17_DC41:
		{
			data2, ok := other.Get_().(D17_DC41)
			return ok && data1.Cf73 == data2.Cf73
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

type D18_DC43 struct {
	Cf74 *C3
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_(Cf74 *C3) D18 {
	return D18{D18_DC43{Cf74}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

func (CompanionStruct_D18_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D18) Dtor_cf74() *C3 {
	return _this.Get_().(D18_DC43).Cf74
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC43:
		{
			return "D18.DC43" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC43:
		{
			data2, ok := other.Get_().(D18_DC43)
			return ok && data1.Cf74 == data2.Cf74
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm2(globalState *GlobalState) _dafny.Int
	Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool
	M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool)
	F23() _dafny.Int
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
	F24() bool
	F24_set_(value bool)
	Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map
	Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int
	M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int)
	F25() bool
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
	F0   _dafny.Map
	F1   bool
	F4   _dafny.Sequence
	F5   _dafny.Int
	F6   bool
	F8   _dafny.Int
	F9   bool
	F14  _dafny.Int
	F15  bool
	F17  _dafny.Int
	_f2  _dafny.Int
	_f3  bool
	_f7  _dafny.Int
	_f10 _dafny.Map
	_f11 _dafny.Sequence
	_f12 _dafny.Int
	_f13 _dafny.Int
	_f16 _dafny.Int
	_f18 _dafny.Array
	_f19 bool
	_f20 bool
	_f21 bool
	_f22 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptyMap
	_this.F1 = false
	_this.F4 = _dafny.EmptySeq
	_this.F5 = _dafny.Zero
	_this.F6 = false
	_this.F8 = _dafny.Zero
	_this.F9 = false
	_this.F14 = _dafny.Zero
	_this.F15 = false
	_this.F17 = _dafny.Zero
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f7 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
	_this._f11 = _dafny.EmptySeq
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.Zero
	_this._f16 = _dafny.Zero
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f19 = false
	_this._f20 = false
	_this._f21 = false
	_this._f22 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Sequence, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 _dafny.Int, f9 bool, f10 _dafny.Map, f11 _dafny.Sequence, f12 _dafny.Int, f13 _dafny.Int, f14 _dafny.Int, f15 bool, f16 _dafny.Int, f17 _dafny.Int, f18 _dafny.Array, f19 bool, f20 bool, f21 bool, f22 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F10() _dafny.Map {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Sequence {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
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
func (_this *GlobalState) F22() _dafny.Int {
	{
		return _this._f22
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f23 _dafny.Int
	F31  _dafny.Int
	_f30 _dafny.Set
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f23 = _dafny.Zero
	_this.F31 = _dafny.Zero
	_this._f30 = _dafny.EmptySet
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F23() _dafny.Int {
	return _this._f23
}
func (_this *C0) Ctor__(f30 _dafny.Set, f31 _dafny.Int, f23 _dafny.Int) {
	{
		(_this)._f30 = f30
		(_this).F31 = f31
		(_this)._f23 = f23
	}
}
func (_this *C0) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(867)).Minus(_this.F31)
	}
}
func (_this *C0) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		var _source1 D0 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jwdbhn")).Cardinality()))))
		_ = _source1
		if _source1.Is_DC1() {
			var _162___mcc_h0 bool = _source1.Get_().(D0_DC1).Cf1
			_ = _162___mcc_h0
			var _163___mcc_h1 _dafny.Int = _source1.Get_().(D0_DC1).Cf2
			_ = _163___mcc_h1
			var _164_cf2 _dafny.Int = _163___mcc_h1
			_ = _164_cf2
			var _165_cf1 bool = _162___mcc_h0
			_ = _165_cf1
			return !(_165_cf1)
		} else if _source1.Is_DC0() {
			var _166___mcc_h2 bool = _source1.Get_().(D0_DC0).Cf0
			_ = _166___mcc_h2
			var _167_cf0 bool = _166___mcc_h2
			_ = _167_cf0
			return _167_cf0
		} else {
			var _168___mcc_h3 D0 = _source1.Get_().(D0_DC2).Cf3
			_ = _168___mcc_h3
			var _169_cf3 D0 = _168___mcc_h3
			_ = _169_cf3
			return false
		}
	}
}
func (_this *C0) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		r2 = p0
		var _170_v0 _dafny.CodePoint
		_ = _170_v0
		_170_v0 = _dafny.CodePoint('h')
		var _171_v1 _dafny.Sequence
		_ = _171_v1
		_171_v1 = _dafny.SeqOf(_170_v0, _170_v0)
		_171_v1 = _171_v1
		var _172_v2 _dafny.Array
		_ = _172_v2
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw37
		_172_v2 = _nw37
		var _173_v3 _dafny.Array
		_ = _173_v3
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw38
		_173_v3 = _nw38
		var _174_v4 _dafny.Array
		_ = _174_v4
		_174_v4 = _173_v3
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_172_v2), 0))
		_ = _index25
		(_172_v2).ArraySet1(_173_v3, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_172_v2), 0))
		_ = _index26
		(_172_v2).ArraySet1(_174_v4, (_index26).Int())
		var _175_v5 D0
		_ = _175_v5
		_175_v5 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(p1))
		var _176_v6 _dafny.Set
		_ = _176_v6
		_176_v6 = _dafny.SetOf(Companion_Default___.Fm24((_dafny.Zero).Minus((_this).F23()), p0, p0, globalState), _175_v5)
		(globalState).F1 = (func() bool {
			if !(!(p1)) {
				return !(Companion_Default___.Fm23(_this.F31, globalState)).Equals(_176_v6)
			}
			return true
		})()
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_173_v3), 0))
		_ = _index27
		(_173_v3).ArraySet1(p1, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_173_v3), 0))
		_ = _index28
		(_173_v3).ArraySet1(p1, (_index28).Int())
		var _177_v7 _dafny.MultiSet
		_ = _177_v7
		_177_v7 = _dafny.MultiSetOf(p1, p1)
		var _178_v8 _dafny.Sequence
		_ = _178_v8
		_178_v8 = _dafny.SeqOf(_dafny.IntOfInt64(44), _this.F31, (_this).F23(), _this.F31)
		var _179_v9 _dafny.Map
		_ = _179_v9
		_179_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_173_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_173_v3), 0))).Int()).(bool), _171_v1)
		var _180_v10 _dafny.Map
		_ = _180_v10
		_180_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(_171_v1, p0, _this.F31, _178_v8, globalState), (_this).F23())
		var _181_v11 _dafny.Sequence
		_ = _181_v11
		_181_v11 = _dafny.SeqOf(p0)
		var _182_v12 D4
		_ = _182_v12
		_182_v12 = Companion_D4_.Create_DC12_((_this).F23(), _this.F31)
		var _183_v13 _dafny.Array
		_ = _183_v13
		var _nwElement0_1 _dafny.Int = (_this).F23()
		_ = _nwElement0_1
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_1, 0)
		(_nw39).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F31)), 1)
		(_nw39).ArraySet1(_this.F31, 2)
		(_nw39).ArraySet1(_this.F31, 3)
		(_nw39).ArraySet1((_this).F23(), 4)
		(_nw39).ArraySet1((_this).F23(), 5)
		(_nw39).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_178_v8).Cardinality())), 6)
		(_nw39).ArraySet1((_179_v9).Cardinality(), 7)
		(_nw39).ArraySet1((_this).F23(), 8)
		(_nw39).ArraySet1((_178_v8).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_178_v8).Cardinality()))).Uint32()).(_dafny.Int), 9)
		(_nw39).ArraySet1((_this).F23(), 10)
		(_nw39).ArraySet1((_this).F23(), 11)
		(_nw39).ArraySet1(_dafny.IntOfInt64(656), 12)
		(_nw39).ArraySet1((_this).F23(), 13)
		(_nw39).ArraySet1((func() _dafny.Int {
			if (_180_v10).Contains(p0) {
				return (_180_v10).Get(p0).(_dafny.Int)
			}
			return _this.F31
		})(), 14)
		(_nw39).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_171_v1).Cardinality()), p0, globalState), 15)
		(_nw39).ArraySet1(_dafny.IntOfInt64(678), 16)
		(_nw39).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ft")).Cardinality()), 17)
		(_nw39).ArraySet1(_this.F31, 18)
		(_nw39).ArraySet1((_this).F23(), 19)
		(_nw39).ArraySet1((_dafny.Zero).Minus(_this.F31), 20)
		(_nw39).ArraySet1(_this.F31, 21)
		(_nw39).ArraySet1(_this.F31, 22)
		(_nw39).ArraySet1(_dafny.IntOfUint32((_181_v11).Cardinality()), 23)
		(_nw39).ArraySet1(_dafny.IntOfInt64(970), 24)
		(_nw39).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_182_v12).Dtor_cf19(), (_177_v7).Cardinality())).Cardinality(), 25)
		_183_v13 = _nw39
		var _184_v14 _dafny.Sequence
		_ = _184_v14
		_184_v14 = _dafny.SeqOf(_183_v13, _183_v13)
		var _rhs44 _dafny.Array = (_184_v14).Select((Companion_Default___.SafeIndex(_this.F31, _dafny.IntOfUint32((_184_v14).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs44
		var _rhs45 _dafny.MultiSet = (_177_v7).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _dafny.SeqOf((_173_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_173_v3), 0))).Int()).(bool)))))
		_ = _rhs45
		var _rhs46 _dafny.Int = ((func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(61), _dafny.IntOfInt64(671))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _185_v15 _dafny.Int
				_185_v15 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(61)).Cmp(_185_v15) <= 0) && ((_185_v15).Cmp(_dafny.IntOfInt64(671)) < 0) {
					_coll10.Add((_185_v15).Minus(_dafny.IntOfInt64(758)), _dafny.MultiSetOf(_170_v0, _170_v0))
				}
			}
			return _coll10.ToMap()
		}()).Cardinality()).Minus((_this.F31).Plus((_this).F23()))
		_ = _rhs46
		var _rhs47 _dafny.Int = Companion_Default___.SafeModuloInt(((func() _dafny.Map {
			if p1 {
				return _180_v10
			}
			return _180_v10
		})()).Cardinality(), (_this.F31).Times(_this.F31))
		_ = _rhs47
		var _rhs48 _dafny.Int = (_this).F23()
		_ = _rhs48
		var _lhs37 *GlobalState = globalState
		_ = _lhs37
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		var _lhs39 *GlobalState = globalState
		_ = _lhs39
		r1 = _rhs44
		_177_v7 = _rhs45
		_lhs37.F17 = _rhs46
		_lhs38.F14 = _rhs47
		_lhs39.F14 = _rhs48
		r0 = ((_dafny.Zero).Minus(_this.F31)).Plus((_this).F23())
		r1 = _183_v13
		var _186_v16 _dafny.Map
		_ = _186_v16
		_186_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _dafny.IntOfInt64(151))
		var _187_v17 _dafny.Map
		_ = _187_v17
		_187_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_173_v3, (_this).Fm3(_171_v1, true, (_this).F23(), _dafny.SeqOf((_186_v16).Cardinality(), _this.F31), globalState))
		r2 = !((_187_v17).Merge(_187_v17)).Contains(_173_v3)
		return r0, r1, r2
	}
}
func (_this *C0) M10(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.MultiSet, _dafny.Array, _dafny.Int, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		r2 = p2
		var _188_v0 _dafny.Sequence
		_ = _188_v0
		_188_v0 = _dafny.UnicodeSeqOfUtf8Bytes("mevk")
		_188_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(717))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_189_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), _dafny.Companion_Sequence_.Concatenate(_188_v0, _188_v0))
		var _190_v1 _dafny.CodePoint
		_ = _190_v1
		_190_v1 = _dafny.CodePoint('w')
		var _191_v2 _dafny.Sequence
		_ = _191_v2
		_191_v2 = _dafny.SeqOf(p0)
		var _192_v3 _dafny.Map
		_ = _192_v3
		_192_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
		var _193_v4 _dafny.Map
		_ = _193_v4
		_193_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _194_v5 D3
		_ = _194_v5
		_194_v5 = Companion_D3_.Create_DC10_(p0, _193_v4)
		var _195_v6 _dafny.Array
		_ = _195_v6
		var _nwElement0_2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_188_v0, _188_v0)
		_ = _nwElement0_2
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_2, 0)
		(_nw40).ArraySet1(_188_v0, 1)
		(_nw40).ArraySet1(_dafny.Companion_Sequence_.Update(_188_v0, (Companion_Default___.SafeIndex(_this.F31, _dafny.IntOfUint32((_188_v0).Cardinality()))).Uint32(), _190_v1), 2)
		(_nw40).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rkd"), 3)
		(_nw40).ArraySet1(_188_v0, 4)
		(_nw40).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("aai"), 5)
		(_nw40).ArraySet1(_188_v0, 6)
		(_nw40).ArraySet1(_188_v0, 7)
		(_nw40).ArraySet1(_188_v0, 8)
		(_nw40).ArraySet1(Companion_Default___.Fm25(p0, p0, _190_v1, globalState), 9)
		(_nw40).ArraySet1(_188_v0, 10)
		(_nw40).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(58))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_196_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_197_i1 _dafny.Int) _dafny.CodePoint {
				return _196_v1
			}
		})(_190_v1))), 11)
		(_nw40).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v0, _188_v0), 12)
		(_nw40).ArraySet1(_188_v0, 13)
		(_nw40).ArraySet1(_188_v0, 14)
		(_nw40).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_198_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_199_i2 _dafny.Int) _dafny.CodePoint {
				return _198_v1
			}
		})(_190_v1))), _188_v0), 15)
		(_nw40).ArraySet1(_188_v0, 16)
		(_nw40).ArraySet1(Companion_Default___.Fm25((_191_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_191_v2).Cardinality()))).Uint32()).(bool), p0, Companion_Default___.Fm26((_192_v3).Cardinality(), p0, globalState), globalState), 17)
		(_nw40).ArraySet1(_188_v0, 18)
		(_nw40).ArraySet1(_188_v0, 19)
		(_nw40).ArraySet1(Companion_Default___.Fm25((func() bool {
			if (_193_v4).Contains(p0) {
				return (_193_v4).Get(p0).(bool)
			}
			return !(p0)
		})(), !((_194_v5).Dtor_cf16()), _190_v1, globalState), 20)
		(_nw40).ArraySet1(_188_v0, 21)
		(_nw40).ArraySet1(_188_v0, 22)
		(_nw40).ArraySet1(_188_v0, 23)
		(_nw40).ArraySet1(_188_v0, 24)
		(_nw40).ArraySet1(_188_v0, 25)
		_195_v6 = _nw40
		var _rhs49 bool = !(((_this).F23()).Cmp(p2) != 0)
		_ = _rhs49
		var _rhs50 _dafny.Array = _195_v6
		_ = _rhs50
		var _rhs51 bool = p0
		_ = _rhs51
		var _lhs40 *GlobalState = globalState
		_ = _lhs40
		_lhs40.F1 = _rhs49
		_195_v6 = _rhs50
		r3 = _rhs51
		var _200_v7 _dafny.Map
		_ = _200_v7
		_200_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_188_v0, _188_v0), _193_v4)
		var _201_v8 _dafny.MultiSet
		_ = _201_v8
		_201_v8 = _dafny.MultiSetOf(p0, p0, p0)
		var _202_v9 _dafny.MultiSet
		_ = _202_v9
		_202_v9 = _dafny.MultiSetOf(Companion_Default___.Fm1(p2, p0, globalState))
		_200_v7 = (Companion_Default___.Fm27(_this.F31, _201_v8, (_202_v9).Cardinality(), p2, globalState)).Merge(_200_v7)
		var _203_v10 _dafny.Map
		_ = _203_v10
		_203_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F23())
		var _204_v11 _dafny.Sequence
		_ = _204_v11
		_204_v11 = _dafny.SeqOf(_this.F31, p3)
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_204_v11).Cardinality())
		_ = _hi1
		for _205_i3 := (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hk")).Cardinality())).Plus((func() _dafny.Int {
			if (_203_v10).Contains((func() bool {
				if (_193_v4).Contains(false) {
					return (_193_v4).Get(false).(bool)
				}
				return !(p0)
			})()) {
				return (_203_v10).Get((func() bool {
					if (_193_v4).Contains(false) {
						return (_193_v4).Get(false).(bool)
					}
					return !(p0)
				})()).(_dafny.Int)
			}
			return (_this).F23()
		})()); _205_i3.Cmp(_hi1) < 0; _205_i3 = _205_i3.Plus(_dafny.One) {
			var _206_v12 _dafny.Array
			_ = _206_v12
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw41
			_206_v12 = _nw41
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _index29
			(_206_v12).ArraySet1(p0, (_index29).Int())
			var _207_v13 _dafny.Array
			_ = _207_v13
			var _nwElement0_3 _dafny.CodePoint = _190_v1
			_ = _nwElement0_3
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
			_ = _nw42
			(_nw42).ArraySet1CodePoint(_nwElement0_3, 0)
			(_nw42).ArraySet1CodePoint(_190_v1, 1)
			(_nw42).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _190_v1
				}
				return _190_v1
			})(), 2)
			(_nw42).ArraySet1CodePoint(_190_v1, 3)
			(_nw42).ArraySet1CodePoint(_190_v1, 4)
			(_nw42).ArraySet1CodePoint(_190_v1, 5)
			_207_v13 = _nw42
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _index30
			var _rhs52 _dafny.Sequence = _191_v2
			_ = _rhs52
			var _rhs53 bool = p0
			_ = _rhs53
			var _rhs54 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_188_v0, (func() _dafny.Sequence {
				if p0 {
					return _188_v0
				}
				return _188_v0
			})())
			_ = _rhs54
			var _rhs55 _dafny.Array = _207_v13
			_ = _rhs55
			var _lhs41 _dafny.Array = _206_v12
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _lhs42
			_191_v2 = _rhs52
			(_lhs41).ArraySet1(_rhs53, (_lhs42).Int())
			_188_v0 = _rhs54
			_207_v13 = _rhs55
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _index31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _index32
			var _rhs56 bool = !(!_dafny.Companion_Sequence_.Contains(_191_v2, (_191_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahkl")).Cardinality()), _dafny.IntOfUint32((_191_v2).Cardinality()))).Uint32()).(bool)))
			_ = _rhs56
			var _rhs57 bool = p0
			_ = _rhs57
			var _rhs58 bool = (_206_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))).Int()).(bool)
			_ = _rhs58
			var _rhs59 bool = (!((_206_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))).Int()).(bool)) || (false)) && ((_206_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))).Int()).(bool))
			_ = _rhs59
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			var _lhs44 _dafny.Array = _206_v12
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _lhs45
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			var _lhs47 _dafny.Array = _206_v12
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_206_v12), 0))
			_ = _lhs48
			_lhs43.F15 = _rhs56
			(_lhs44).ArraySet1(_rhs57, (_lhs45).Int())
			_lhs46.F1 = _rhs58
			(_lhs47).ArraySet1(_rhs59, (_lhs48).Int())
			(globalState).F8 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F31, _205_i3))
			(globalState).F17 = (_this).Fm2(globalState)
		}
		(_this).F31 = Companion_Default___.SafeDivisionInt((_this.F31).Minus((_dafny.Zero).Minus((_this).F23())), (_201_v8).Cardinality())
		r0 = _202_v9
		var _208_v14 _dafny.Array
		_ = _208_v14
		var _nwElement0_4 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Contains(p0)
		_ = _nwElement0_4
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(3))
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_4, 0)
		(_nw43).ArraySet1(p0, 1)
		(_nw43).ArraySet1(p0, 2)
		_208_v14 = _nw43
		r1 = _208_v14
		var _209_v15 _dafny.Map
		_ = _209_v15
		_209_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_188_v0).Cardinality()), _dafny.IntOfInt64(497))
		r2 = (func() _dafny.Int {
			if (_202_v9).Contains((_dafny.Zero).Minus(p3)) {
				return (_202_v9).Multiplicity((_dafny.Zero).Minus(p3))
			}
			return (func() _dafny.Int {
				if (_209_v15).Contains((_this).Fm2(globalState)) {
					return (_209_v15).Get((_this).Fm2(globalState)).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_this.F31)
			})()
		})()
		r3 = p0
		return r0, r1, r2, r3
	}
}
func (_this *C0) F30() _dafny.Set {
	{
		return _this._f30
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f24 bool
	_f25 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f24 = false
	_this._f25 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F24() bool {
	return _this._f24
}
func (_this *C1) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C1) F25() bool {
	return _this._f25
}
func (_this *C1) Ctor__(f24 bool, f25 bool) {
	{
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C1) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), (_this).F25()))))
	}
}
func (_this *C1) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfUint32((_dafny.SeqOf(true, (_this).F25())).Cardinality())).Times(_dafny.IntOfInt64(953))).Plus((func() _dafny.Int {
			if _this.F24() {
				return _dafny.IntOfInt64(954)
			}
			return (_dafny.MultiSetOf(_dafny.IntOfInt64(36), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lidbxgdo")).Cardinality())))).Cardinality()
		})())
	}
}
func (_this *C1) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _210_v0 _dafny.Sequence
		_ = _210_v0
		_210_v0 = _dafny.UnicodeSeqOfUtf8Bytes("pqbcphfi")
		var _211_v1 D3
		_ = _211_v1
		_211_v1 = Companion_D3_.Create_DC8_(_210_v0)
		var _source2 D3 = _211_v1
		_ = _source2
		if _source2.Is_DC9() {
			var _212___mcc_h0 _dafny.Int = _source2.Get_().(D3_DC9).Cf15
			_ = _212___mcc_h0
			var _213_cf15 _dafny.Int = _212___mcc_h0
			_ = _213_cf15
			var _214_v2 _dafny.Array
			_ = _214_v2
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw44
			_214_v2 = _nw44
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))
			_ = _index33
			(_214_v2).ArraySet1((_this).F25(), (_index33).Int())
			var _215_v3 _dafny.Map
			_ = _215_v3
			_215_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _this.F24())
			var _216_v4 _dafny.Sequence
			_ = _216_v4
			_216_v4 = _dafny.SeqOf(_213_cf15, (_215_v3).Cardinality(), _213_cf15)
			var _217_v5 _dafny.Set
			_ = _217_v5
			_217_v5 = _dafny.SetOf(_this.F24())
			var _218_v6 _dafny.CodePoint
			_ = _218_v6
			_218_v6 = _dafny.CodePoint('l')
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))
			_ = _index34
			var _rhs60 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_219_cf15 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_220_i0 _dafny.Int) _dafny.Int {
					return _219_cf15
				}
			})(_213_cf15))), _216_v4), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_216_v4).Cardinality()), _213_cf15, _213_cf15), _216_v4))
			_ = _rhs60
			var _rhs61 bool = ((_217_v5).Union(_217_v5)).IsProperSubsetOf(_217_v5)
			_ = _rhs61
			var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_210_v0, (Companion_Default___.SafeIndex(_213_cf15, _dafny.IntOfUint32((_210_v0).Cardinality()))).Uint32(), _218_v6)
			_ = _rhs62
			var _rhs63 bool = !((_this).F25()) || (_this.F24())
			_ = _rhs63
			var _rhs64 _dafny.Sequence = _210_v0
			_ = _rhs64
			var _lhs49 *GlobalState = globalState
			_ = _lhs49
			var _lhs50 _dafny.Array = _214_v2
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))
			_ = _lhs51
			var _lhs52 *GlobalState = globalState
			_ = _lhs52
			_lhs49.F6 = _rhs60
			(_lhs50).ArraySet1(_rhs61, (_lhs51).Int())
			_210_v0 = _rhs62
			_lhs52.F1 = _rhs63
			_210_v0 = _rhs64
			var _221_v7 _dafny.Array
			_ = _221_v7
			var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw45
			_221_v7 = _nw45
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_221_v7), 0))
			_ = _index35
			(_221_v7).ArraySet1(_dafny.MultiSetOf(_213_cf15), (_index35).Int())
			var _222_v8 _dafny.MultiSet
			_ = _222_v8
			_222_v8 = _dafny.MultiSetOf(_213_cf15)
			var _223_v9 _dafny.Map
			_ = _223_v9
			_223_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _213_cf15)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_221_v7), 0))
			_ = _index36
			(_221_v7).ArraySet1(((_222_v8).Union(_222_v8)).Difference((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_223_v9).Cardinality(), _this.F24())).Cardinality())).Union(_222_v8)), (_index36).Int())
			if false {
				(globalState).F8 = _213_cf15
				var _224_v10 _dafny.Array
				_ = _224_v10
				var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
				_ = _nw46
				_224_v10 = _nw46
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_224_v10), 0))
				_ = _index37
				(_224_v10).ArraySet1(_213_cf15, (_index37).Int())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_224_v10), 0))
				_ = _index38
				(_224_v10).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_213_cf15)), (_index38).Int())
				(globalState).F17 = Companion_Default___.Fm1(((_224_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_224_v10), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(582)), ((_this).F25()) == ((_this).F25()), globalState)
				var _225_v11 _dafny.MultiSet
				_ = _225_v11
				_225_v11 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_216_v4, _216_v4))
				_225_v11 = _225_v11
				var _226_v12 _dafny.MultiSet
				_ = _226_v12
				_226_v12 = _dafny.MultiSetOf((_214_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))).Int()).(bool), _this.F24())
				r2 = ((_226_v12).Intersection(Companion_Default___.Fm22(_213_cf15, (_214_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))).Int()).(bool), _this.F24(), (_224_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_224_v10), 0))).Int()).(_dafny.Int), globalState))).IsDisjointFrom(_dafny.MultiSetOf((_214_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))).Int()).(bool), _this.F24(), (_this).F25(), true))
			} else {
				(globalState).F1 = _this.F24()
				(_this).F24_set_((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(615), _213_cf15)).Cmp((_213_cf15).Plus(_213_cf15)) > 0)
				r2 = (_214_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))).Int()).(bool)
				_210_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(517))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_227_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_228_i1 _dafny.Int) _dafny.CodePoint {
						return _227_v6
					}
				})(_218_v6))), _210_v0)
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_214_v2), 0))
				_ = _index39
				(_214_v2).ArraySet1((false) || ((_this).F25()), (_index39).Int())
			}
			(globalState).F5 = (_dafny.Zero).Minus(_213_cf15)
		} else if _source2.Is_DC10() {
			var _229___mcc_h1 bool = _source2.Get_().(D3_DC10).Cf16
			_ = _229___mcc_h1
			var _230___mcc_h2 _dafny.Map = _source2.Get_().(D3_DC10).Cf17
			_ = _230___mcc_h2
			var _231_cf17 _dafny.Map = _230___mcc_h2
			_ = _231_cf17
			var _232_cf16 bool = _229___mcc_h1
			_ = _232_cf16
			var _233_v13 _dafny.Array
			_ = _233_v13
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw47
			_233_v13 = _nw47
			var _234_v14 _dafny.Array
			_ = _234_v14
			_234_v14 = _233_v13
			_234_v14 = _234_v14
			var _235_v16 _dafny.Array
			_ = _235_v16
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_2
			var _nw48 _dafny.Array
			_ = _nw48
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw48 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Set = (func(_236_cf16 bool) func(_dafny.Int) _dafny.Set {
					return func(_237_i2 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll11 = _dafny.NewBuilder()
							_ = _coll11
							for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(380), _dafny.IntOfInt64(719))); ; {
								_compr_11, _ok11 := _iter11()
								if !_ok11 {
									break
								}
								var _238_v15 _dafny.Int
								_238_v15 = interface{}(_compr_11).(_dafny.Int)
								if ((_dafny.IntOfInt64(380)).Cmp(_238_v15) <= 0) && ((_238_v15).Cmp(_dafny.IntOfInt64(719)) < 0) {
									_coll11.Add(Companion_Default___.SafeDivisionInt(_238_v15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_cf16, _dafny.MultiSetOf(_this.F24()))).Cardinality()))
								}
							}
							return _coll11.ToSet()
						}()
					}
				})(_232_cf16)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw48 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw48).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw48).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_235_v16 = _nw48
			var _239_v17 _dafny.Int
			_ = _239_v17
			_239_v17 = _dafny.IntOfInt64(721)
			var _240_v18 _dafny.Map
			_ = _240_v18
			_240_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_232_cf16, _239_v17)
			var _241_v19 _dafny.Set
			_ = _241_v19
			_241_v19 = _dafny.SetOf((_240_v18).Cardinality(), (_dafny.Zero).Minus(_239_v17), (_dafny.SetOf(_232_cf16)).Cardinality(), _dafny.IntOfInt64(63))
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_235_v16), 0))
			_ = _index40
			(_235_v16).ArraySet1(_241_v19, (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_235_v16), 0))
			_ = _index41
			(_235_v16).ArraySet1(_241_v19, (_index41).Int())
			var _242_v20 _dafny.Array
			_ = _242_v20
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
			_ = _nw49
			_242_v20 = _nw49
			var _243_v21 _dafny.Sequence
			_ = _243_v21
			_243_v21 = _dafny.SeqOf(_242_v20)
			var _244_v23 _dafny.Sequence
			_ = _244_v23
			_244_v23 = _dafny.SeqOf(_242_v20, _242_v20, (_243_v21).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(628), _dafny.IntOfInt64(146))); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _245_v22 _dafny.Int
					_245_v22 = interface{}(_compr_12).(_dafny.Int)
					if ((_dafny.IntOfInt64(628)).Cmp(_245_v22) <= 0) && ((_245_v22).Cmp(_dafny.IntOfInt64(146)) < 0) {
						_coll12.Add(Companion_Default___.SafeDivisionInt(_245_v22, _dafny.IntOfUint32((_dafny.SeqOf(_232_cf16)).Cardinality())), _232_cf16)
					}
				}
				return _coll12.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_243_v21).Cardinality()))).Uint32()).(_dafny.Array))
			_242_v20 = (_244_v23).Select((Companion_Default___.SafeIndex((_239_v17).Plus(_dafny.IntOfInt64(155)), _dafny.IntOfUint32((_244_v23).Cardinality()))).Uint32()).(_dafny.Array)
			var _246_v24 _dafny.MultiSet
			_ = _246_v24
			_246_v24 = _dafny.MultiSetOf((_this).F25())
			var _rhs65 _dafny.Int = _239_v17
			_ = _rhs65
			var _rhs66 bool = false
			_ = _rhs66
			var _rhs67 bool = (_dafny.MultiSetOf(_232_cf16, _232_cf16)).IsSubsetOf((Companion_D7_.Create_DC19_(_246_v24)).Dtor_cf34())
			_ = _rhs67
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			_lhs53.F5 = _rhs65
			_lhs54.F15 = _rhs66
			r2 = _rhs67
		} else {
			var _247___mcc_h3 _dafny.Sequence = _source2.Get_().(D3_DC8).Cf14
			_ = _247___mcc_h3
			var _248_cf14 _dafny.Sequence = _247___mcc_h3
			_ = _248_cf14
			var _249_v25 _dafny.CodePoint
			_ = _249_v25
			_249_v25 = _dafny.CodePoint('l')
			_249_v25 = _249_v25
			var _250_v26 _dafny.Array
			_ = _250_v26
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_3
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = func(_251_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(552)), _dafny.SeqOf(_dafny.IntOfInt64(-777), _dafny.IntOfUint32((_dafny.SeqOf(_this.F24(), _this.F24())).Cardinality())))
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
			_250_v26 = _nw50
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_250_v26), 0))
			_ = _index42
			(_250_v26).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}(func(_252_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-996)
			})), (_index42).Int())
			var _253_v27 _dafny.Int
			_ = _253_v27
			_253_v27 = _dafny.IntOfInt64(40)
			var _254_v28 _dafny.Sequence
			_ = _254_v28
			_254_v28 = _dafny.SeqOf(_253_v27)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_250_v26), 0))
			_ = _index43
			(_250_v26).ArraySet1(_254_v28, (_index43).Int())
			var _255_v29 _dafny.Sequence
			_ = _255_v29
			_255_v29 = _dafny.SeqOf(_this.F24(), _this.F24(), !((_this).F25()))
			var _256_v30 _dafny.MultiSet
			_ = _256_v30
			_256_v30 = _dafny.MultiSetOf(_this.F24(), (_this).F25())
			var _257_v31 _dafny.Set
			_ = _257_v31
			_257_v31 = _dafny.SetOf(_253_v27)
			(_this).M9(_249_v25, !((_this).F25()) || (!(_this.F24())), (_255_v29).Select((Companion_Default___.SafeIndex(_253_v27, _dafny.IntOfUint32((_255_v29).Cardinality()))).Uint32()).(bool), Companion_Default___.SafeModuloInt((_256_v30).Cardinality(), (_257_v31).Cardinality()), globalState)
			var _258_v32 _dafny.Array
			_ = _258_v32
			var _nwElement0_5 bool = _this.F24()
			_ = _nwElement0_5
			var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
			_ = _nw51
			(_nw51).ArraySet1(_nwElement0_5, 0)
			(_nw51).ArraySet1(true, 1)
			(_nw51).ArraySet1(_this.F24(), 2)
			(_nw51).ArraySet1(_this.F24(), 3)
			(_nw51).ArraySet1((_this).F25(), 4)
			(_nw51).ArraySet1(_this.F24(), 5)
			(_nw51).ArraySet1((_this).F25(), 6)
			(_nw51).ArraySet1((_this).F25(), 7)
			(_nw51).ArraySet1(_this.F24(), 8)
			(_nw51).ArraySet1(!(_this.F24()), 9)
			(_nw51).ArraySet1(!((_this).F25()), 10)
			(_nw51).ArraySet1(_this.F24(), 11)
			(_nw51).ArraySet1((_this).F25(), 12)
			(_nw51).ArraySet1(true, 13)
			(_nw51).ArraySet1((_255_v29).Select((Companion_Default___.SafeIndex(_253_v27, _dafny.IntOfUint32((_255_v29).Cardinality()))).Uint32()).(bool), 14)
			(_nw51).ArraySet1(_this.F24(), 15)
			(_nw51).ArraySet1((_this).F25(), 16)
			(_nw51).ArraySet1((_this).F25(), 17)
			(_nw51).ArraySet1((_this).F25(), 18)
			(_nw51).ArraySet1((_this).F25(), 19)
			(_nw51).ArraySet1(_this.F24(), 20)
			(_nw51).ArraySet1((_this).F25(), 21)
			(_nw51).ArraySet1((_this).F25(), 22)
			(_nw51).ArraySet1(_this.F24(), 23)
			(_nw51).ArraySet1((_this).F25(), 24)
			_258_v32 = _nw51
			var _259_v33 _dafny.Set
			_ = _259_v33
			_259_v33 = _dafny.SetOf(_258_v32)
			var _260_v34 T0
			_ = _260_v34
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__((_259_v33).Intersection(_259_v33), Companion_Default___.Fm1(_253_v27, (_this).F25(), globalState), _253_v27)
			_260_v34 = _nw52
			r1 = _260_v34
		}
		var _261_v35 _dafny.Sequence
		_ = _261_v35
		_261_v35 = _dafny.SeqOf((_this).F25())
		var _262_v36 _dafny.Int
		_ = _262_v36
		_262_v36 = _dafny.IntOfInt64(-850)
		var _263_v37 _dafny.Map
		_ = _263_v37
		_263_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
		var _264_v38 _dafny.Map
		_ = _264_v38
		_264_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _262_v36)
		var _265_v39 D1
		_ = _265_v39
		_265_v39 = Companion_D1_.Create_DC5_((_261_v35).Select((Companion_Default___.SafeIndex(_262_v36, _dafny.IntOfUint32((_261_v35).Cardinality()))).Uint32()).(bool), _263_v37, _262_v36, _262_v36, _264_v38)
		var _266_v40 _dafny.Map
		_ = _266_v40
		_266_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _265_v39)
		var _267_v41 _dafny.MultiSet
		_ = _267_v41
		_267_v41 = _dafny.MultiSetOf(Companion_Default___.Fm28(_262_v36, _210_v0, _262_v36, globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _265_v39), _266_v40)
		var _268_i5 _dafny.Int
		_ = _268_i5
		_268_i5 = _dafny.Zero
		{
			for ((_dafny.MultiSetOf(_266_v40, _266_v40)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _265_v39))))).IsProperSubsetOf((_dafny.MultiSetOf(_266_v40)).Difference(_267_v41)) {
				{
					if (_268_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_268_i5 = (_268_i5).Plus(_dafny.One)
					var _269_v42 _dafny.Array
					_ = _269_v42
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_4
					var _nw53 _dafny.Array
					_ = _nw53
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw53 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) bool = func(_270_i6 _dafny.Int) bool {
							return _this.F24()
						}
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw53 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw53).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw53).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_269_v42 = _nw53
					var _271_v43 _dafny.Set
					_ = _271_v43
					_271_v43 = _dafny.SetOf(_269_v42, _269_v42)
					var _272_v44 _dafny.Map
					_ = _272_v44
					_272_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v36, _dafny.UnicodeSeqOfUtf8Bytes("wvrqm"))
					var _273_v45 *C0
					_ = _273_v45
					var _nw54 *C0 = New_C0_()
					_ = _nw54
					_nw54.Ctor__(_271_v43, _262_v36, Companion_Default___.SafeDivisionInt(_262_v36, (_272_v44).Cardinality()))
					_273_v45 = _nw54
					var _274_v46 _dafny.Sequence
					_ = _274_v46
					_274_v46 = _dafny.SeqOf(_262_v36)
					var _275_v47 _dafny.Set
					_ = _275_v47
					_275_v47 = _dafny.SetOf(_dafny.IntOfUint32((_274_v46).Cardinality()))
					var _276_v48 _dafny.MultiSet
					_ = _276_v48
					_276_v48 = _dafny.MultiSetOf(_275_v47)
					var _277_v49 _dafny.Map
					_ = _277_v49
					_277_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.SetOf(_262_v36))).Difference(_276_v48), (_273_v45.F31).Minus((_this).Fm5((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(292), _269_v42)).Cardinality(), _273_v45.F31, globalState)))
					var _278_v50 _dafny.Sequence
					_ = _278_v50
					_278_v50 = _dafny.SeqOf(_275_v47, _275_v47, _275_v47)
					var _279_v51 _dafny.Map
					_ = _279_v51
					_279_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F24()), _269_v42)
					var _280_v52 _dafny.MultiSet
					_ = _280_v52
					_280_v52 = _dafny.MultiSetOf(_269_v42, (func() _dafny.Array {
						if (_279_v51).Contains(_this.F24()) {
							return (_279_v51).Get(_this.F24()).(_dafny.Array)
						}
						return _269_v42
					})(), _269_v42, _269_v42)
					_277_v49 = (_277_v49).Update(_dafny.MultiSetFromSeq(_278_v50), (func() _dafny.Int {
						if (_280_v52).Contains(_269_v42) {
							return (_280_v52).Multiplicity(_269_v42)
						}
						return (_dafny.Zero).Minus((_this).Fm5(_273_v45.F31, _273_v45.F31, globalState))
					})())
					r3 = _273_v45.F31
					var _281_v53 D0
					_ = _281_v53
					_281_v53 = Companion_D0_.Create_DC1_(_this.F24(), _262_v36)
					var _282_v54 _dafny.Map
					_ = _282_v54
					_282_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_274_v46, (Companion_Default___.SafeIndex((_281_v53).Dtor_cf2(), _dafny.IntOfUint32((_274_v46).Cardinality()))).Uint32(), _273_v45.F31), (_this).F25())
					var _283_v55 _dafny.MultiSet
					_ = _283_v55
					_283_v55 = _dafny.MultiSetOf(Companion_Default___.Fm22(_273_v45.F31, (_this).F25(), !(true), _273_v45.F31, globalState), Companion_Default___.Fm22(_273_v45.F31, false, (_this).F25(), (_282_v54).Cardinality(), globalState))
					_283_v55 = (_283_v55).Intersection((_283_v55).Union(_283_v55))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _284_v56 _dafny.Array
		_ = _284_v56
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw55
		_284_v56 = _nw55
		var _285_v57 _dafny.Map
		_ = _285_v57
		_285_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v36, (_this).F25())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_284_v56), 0))
		_ = _index44
		(_284_v56).ArraySet1(_285_v57, (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_284_v56), 0))
		_ = _index45
		(_284_v56).ArraySet1(_285_v57, (_index45).Int())
		if _this.F24() {
			var _286_v58 _dafny.Array
			_ = _286_v58
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_5
			var _nw56 _dafny.Array
			_ = _nw56
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw56 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = func(_287_i7 _dafny.Int) _dafny.CodePoint {
					return (func() _dafny.CodePoint {
						if (_this).F25() {
							return _dafny.CodePoint('g')
						}
						return _dafny.CodePoint('c')
					})()
				}
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
			_286_v58 = _nw56
			var _288_v59 _dafny.CodePoint
			_ = _288_v59
			_288_v59 = _dafny.CodePoint('a')
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_286_v58), 0))
			_ = _index46
			(_286_v58).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_this).F25() {
					return _288_v59
				}
				return _288_v59
			})(), (_index46).Int())
			var _289_v60 _dafny.Sequence
			_ = _289_v60
			_289_v60 = _dafny.SeqOf(_dafny.IntOfInt64(973))
			var _290_v61 _dafny.MultiSet
			_ = _290_v61
			_290_v61 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_289_v60, _289_v60))
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_286_v58), 0))
			_ = _index47
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_284_v56), 0))
			_ = _index48
			var _rhs68 _dafny.Int = (func() _dafny.Int {
				if (_290_v61).Contains(_dafny.Companion_Sequence_.Concatenate(_289_v60, _289_v60)) {
					return (_290_v61).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_289_v60, _289_v60))
				}
				return _262_v36
			})()
			_ = _rhs68
			var _rhs69 _dafny.CodePoint = _288_v59
			_ = _rhs69
			var _rhs70 _dafny.Map = (_284_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_284_v56), 0))).Int()).(_dafny.Map)
			_ = _rhs70
			var _lhs55 *GlobalState = globalState
			_ = _lhs55
			var _lhs56 _dafny.Array = _286_v58
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_286_v58), 0))
			_ = _lhs57
			var _lhs58 _dafny.Array = _284_v56
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_284_v56), 0))
			_ = _lhs59
			_lhs55.F8 = _rhs68
			(_lhs56).ArraySet1CodePoint(_rhs69, (_lhs57).Int())
			(_lhs58).ArraySet1(_rhs70, (_lhs59).Int())
			var _291_v62 _dafny.Array
			_ = _291_v62
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw57
			_291_v62 = _nw57
			var _rhs71 _dafny.Int = (_dafny.IntOfUint32((_261_v35).Cardinality())).Times((_262_v36).Plus(_262_v36))
			_ = _rhs71
			var _rhs72 _dafny.Array = _291_v62
			_ = _rhs72
			var _rhs73 bool = (_261_v35).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_262_v36, (_this).F25(), globalState), _dafny.IntOfUint32((_261_v35).Cardinality()))).Uint32()).(bool)
			_ = _rhs73
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			_lhs60.F14 = _rhs71
			_291_v62 = _rhs72
			_lhs61.F9 = _rhs73
			(globalState).F8 = _262_v36
			var _292_v63 D1
			_ = _292_v63
			_292_v63 = Companion_D1_.Create_DC4_((_this).F25(), _this.F24())
			var _293_v64 _dafny.Array
			_ = _293_v64
			var _nwElement0_6 bool = (_292_v63).Dtor_cf5()
			_ = _nwElement0_6
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_6, 0)
			(_nw58).ArraySet1(_this.F24(), 1)
			(_nw58).ArraySet1(_this.F24(), 2)
			(_nw58).ArraySet1(!((_this).F25()) || ((_this).F25()), 3)
			_293_v64 = _nw58
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_293_v64), 0))
			_ = _index49
			(_293_v64).ArraySet1(_this.F24(), (_index49).Int())
			var _294_v65 _dafny.Set
			_ = _294_v65
			_294_v65 = _dafny.SetOf(_this.F24(), (_this).F25())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_293_v64), 0))
			_ = _index50
			(_293_v64).ArraySet1((_294_v65).IsSubsetOf(_294_v65), (_index50).Int())
			(globalState).F15 = (_this).F25()
		} else {
			var _295_v66 _dafny.CodePoint
			_ = _295_v66
			_295_v66 = _dafny.CodePoint('p')
			_295_v66 = _295_v66
			var _296_v67 _dafny.Array
			_ = _296_v67
			var _nwElement0_7 bool = _this.F24()
			_ = _nwElement0_7
			var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(25))
			_ = _nw59
			(_nw59).ArraySet1(_nwElement0_7, 0)
			(_nw59).ArraySet1(!(false), 1)
			(_nw59).ArraySet1(true, 2)
			(_nw59).ArraySet1((_this).F25(), 3)
			(_nw59).ArraySet1((_this).F25(), 4)
			(_nw59).ArraySet1(true, 5)
			(_nw59).ArraySet1(true, 6)
			(_nw59).ArraySet1(_this.F24(), 7)
			(_nw59).ArraySet1(true, 8)
			(_nw59).ArraySet1((_this).F25(), 9)
			(_nw59).ArraySet1(false, 10)
			(_nw59).ArraySet1((_this).F25(), 11)
			(_nw59).ArraySet1(Companion_Default___.Fm29(_262_v36, _262_v36, (_this).F25(), globalState), 12)
			(_nw59).ArraySet1(_this.F24(), 13)
			(_nw59).ArraySet1((_this).F25(), 14)
			(_nw59).ArraySet1((_this).F25(), 15)
			(_nw59).ArraySet1(_this.F24(), 16)
			(_nw59).ArraySet1(_this.F24(), 17)
			(_nw59).ArraySet1((_this).F25(), 18)
			(_nw59).ArraySet1(_this.F24(), 19)
			(_nw59).ArraySet1((_this).F25(), 20)
			(_nw59).ArraySet1((_this).F25(), 21)
			(_nw59).ArraySet1((_this).F25(), 22)
			(_nw59).ArraySet1((_this).F25(), 23)
			(_nw59).ArraySet1(!(_this.F24()), 24)
			_296_v67 = _nw59
			var _297_v68 _dafny.Set
			_ = _297_v68
			_297_v68 = _dafny.SetOf(_296_v67, _296_v67)
			var _298_v69 *C0
			_ = _298_v69
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__(_297_v68, _dafny.IntOfUint32((_261_v35).Cardinality()), (_262_v36).Times(_262_v36))
			_298_v69 = _nw60
			var _299_v70 _dafny.Array
			_ = _299_v70
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_6
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) D0 = func(_300_i8 _dafny.Int) D0 {
					return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_((_this).F25()))
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw61 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw61).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw61).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_299_v70 = _nw61
			var _301_v71 _dafny.Map
			_ = _301_v71
			_301_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_299_v70, _295_v66)
			var _302_v72 T0
			_ = _302_v72
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__((_298_v69).F30(), _262_v36, _298_v69.F31)
			_302_v72 = _nw62
			var _303_v73 D3
			_ = _303_v73
			_303_v73 = Companion_D3_.Create_DC9_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_302_v72, _298_v69.F31)).Cardinality())
			var _304_v74 _dafny.Map
			_ = _304_v74
			_304_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v36, _210_v0)
			var _305_v75 _dafny.MultiSet
			_ = _305_v75
			_305_v75 = _dafny.MultiSetOf(_262_v36, _dafny.IntOfInt64(-199))
			var _306_v76 _dafny.MultiSet
			_ = _306_v76
			_306_v76 = _dafny.MultiSetOf(_this.F24())
			var _307_v77 _dafny.Sequence
			_ = _307_v77
			_307_v77 = _dafny.SeqOf(_262_v36, _262_v36)
			var _308_v79 _dafny.Set
			_ = _308_v79
			_308_v79 = _dafny.SetOf(_dafny.IntOfInt64(978))
			var _309_v80 _dafny.Array
			_ = _309_v80
			var _nwElement0_8 _dafny.Int = (_302_v72).F23()
			_ = _nwElement0_8
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_8, 0)
			(_nw63).ArraySet1(_298_v69.F31, 1)
			(_nw63).ArraySet1(_dafny.IntOfInt64(966), 2)
			(_nw63).ArraySet1(_298_v69.F31, 3)
			(_nw63).ArraySet1((_302_v72).F23(), 4)
			(_nw63).ArraySet1(_dafny.IntOfUint32((_261_v35).Cardinality()), 5)
			(_nw63).ArraySet1(_262_v36, 6)
			(_nw63).ArraySet1(_262_v36, 7)
			(_nw63).ArraySet1(_298_v69.F31, 8)
			(_nw63).ArraySet1((_304_v74).Cardinality(), 9)
			(_nw63).ArraySet1(_262_v36, 10)
			(_nw63).ArraySet1((_305_v75).Cardinality(), 11)
			(_nw63).ArraySet1((_302_v72).F23(), 12)
			(_nw63).ArraySet1(_262_v36, 13)
			(_nw63).ArraySet1((func() _dafny.Int {
				if (_306_v76).Contains((_this).F25()) {
					return (_306_v76).Multiplicity((_this).F25())
				}
				return _dafny.IntOfUint32((_307_v77).Cardinality())
			})(), 14)
			(_nw63).ArraySet1((_302_v72).F23(), 15)
			(_nw63).ArraySet1(_dafny.IntOfUint32((_307_v77).Cardinality()), 16)
			(_nw63).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_307_v77).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _310_v78 _dafny.Int
					_310_v78 = interface{}(_compr_13).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_307_v77, _310_v78) {
						_coll13.Add((_310_v78).Minus(_dafny.IntOfInt64(421)))
					}
				}
				return _coll13.ToSet()
			}(), _308_v79)).Cardinality(), _this.F24())).Cardinality(), 17)
			(_nw63).ArraySet1((_307_v77).Select((Companion_Default___.SafeIndex(_262_v36, _dafny.IntOfUint32((_307_v77).Cardinality()))).Uint32()).(_dafny.Int), 18)
			(_nw63).ArraySet1(_262_v36, 19)
			(_nw63).ArraySet1(_298_v69.F31, 20)
			(_nw63).ArraySet1(_dafny.IntOfUint32((_210_v0).Cardinality()), 21)
			(_nw63).ArraySet1(_262_v36, 22)
			(_nw63).ArraySet1(_262_v36, 23)
			(_nw63).ArraySet1(_262_v36, 24)
			_309_v80 = _nw63
			var _311_v81 D8
			_ = _311_v81
			_311_v81 = Companion_D8_.Create_DC22_(_309_v80)
			var _312_v82 _dafny.Map
			_ = _312_v82
			_312_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v73, (_311_v81).Dtor_cf39())
			var _313_v83 D5
			_ = _313_v83
			_313_v83 = Companion_D5_.Create_DC15_((_302_v72).F23(), _295_v66, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_299_v70, _295_v66), _312_v82)
			var _314_v84 D5
			_ = _314_v84
			_314_v84 = Companion_D5_.Create_DC15_(_dafny.IntOfInt64(589), _295_v66, ((_301_v71).Update(_299_v70, _295_v66)).Merge(_301_v71), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v73, _309_v80)).Merge((_313_v83).Dtor_cf28()))
			_314_v84 = _314_v84
			if (Companion_Default___.Fm22((_302_v72).F23(), (_this).F25(), !((_this).F25()), (_302_v72).F23(), globalState)).IsProperSubsetOf(_306_v76) {
				var _315_v85 _dafny.Array
				_ = _315_v85
				var _len0_7 _dafny.Int = _dafny.One
				_ = _len0_7
				var _nw64 _dafny.Array
				_ = _nw64
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw64 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.MultiSet = (func(_316_v76 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_317_i9 _dafny.Int) _dafny.MultiSet {
							return _316_v76
						}
					})(_306_v76)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw64 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw64).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw64).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_315_v85 = _nw64
				var _318_v86 _dafny.Sequence
				_ = _318_v86
				_318_v86 = _dafny.SeqOf(_315_v85)
				_315_v85 = (_318_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-367), _dafny.IntOfUint32((_318_v86).Cardinality()))).Uint32()).(_dafny.Array)
				(globalState).F5 = (_302_v72).F23()
				var _319_v87 _dafny.Map
				_ = _319_v87
				_319_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_298_v69.F31, _dafny.IntOfUint32((_307_v77).Cardinality()))
				(globalState).F1 = (_308_v79).IsDisjointFrom(_dafny.SetOf((func() _dafny.Int {
					if (_319_v87).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_319_v87, (_302_v72).F23())).Cardinality()) {
						return (_319_v87).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_319_v87, (_302_v72).F23())).Cardinality()).(_dafny.Int)
					}
					return (_302_v72).F23()
				})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pdjrt")).Cardinality()), (_this).F25())).Cardinality(), _298_v69.F31))
				var _320_v88 _dafny.Array
				_ = _320_v88
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(9))
				_ = _nw65
				_320_v88 = _nw65
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_320_v88), 0))
				_ = _index51
				(_320_v88).ArraySet1((func() D5 {
					if _this.F24() {
						return Companion_D5_.Create_DC15_(_262_v36, _295_v66, _301_v71, ((_314_v84).Dtor_cf28()).Update(Companion_D3_.Create_DC9_(_262_v36), _309_v80))
					}
					return _314_v84
				})(), (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_320_v88), 0))
				_ = _index52
				(_320_v88).ArraySet1(_313_v83, (_index52).Int())
				_319_v87 = (_319_v87).Update(_dafny.IntOfInt64(446), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_261_v35).Cardinality()), (_302_v72).F23()))
			} else {
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_309_v80), 0))
				_ = _index53
				(_309_v80).ArraySet1(_298_v69.F31, (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_309_v80), 0))
				_ = _index54
				(_309_v80).ArraySet1((_dafny.Zero).Minus(((_298_v69).F30()).Cardinality()), (_index54).Int())
				var _321_v89 _dafny.Array
				_ = _321_v89
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw66
				_321_v89 = _nw66
				var _322_v90 _dafny.Sequence
				_ = _322_v90
				_322_v90 = _dafny.SeqOf(_321_v89)
				var _323_v91 _dafny.Array
				_ = _323_v91
				var _nwElement0_9 _dafny.Array = _321_v89
				_ = _nwElement0_9
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(9))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_9, 0)
				(_nw67).ArraySet1(_321_v89, 1)
				(_nw67).ArraySet1(_321_v89, 2)
				(_nw67).ArraySet1(_321_v89, 3)
				(_nw67).ArraySet1(_321_v89, 4)
				(_nw67).ArraySet1(_321_v89, 5)
				(_nw67).ArraySet1((_322_v90).Select((Companion_Default___.SafeIndex((_307_v77).Select((Companion_Default___.SafeIndex((_302_v72).F23(), _dafny.IntOfUint32((_307_v77).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_322_v90).Cardinality()))).Uint32()).(_dafny.Array), 6)
				(_nw67).ArraySet1(_321_v89, 7)
				(_nw67).ArraySet1(_321_v89, 8)
				_323_v91 = _nw67
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_323_v91), 0))
				_ = _index55
				(_323_v91).ArraySet1(_321_v89, (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_323_v91), 0))
				_ = _index56
				(_323_v91).ArraySet1(_321_v89, (_index56).Int())
				r3 = (_dafny.Zero).Minus(_298_v69.F31)
				var _324_v92 *C0
				_ = _324_v92
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__(_297_v68, (_302_v72).F23(), (_302_v72).F23())
				_324_v92 = _nw68
				(globalState).F6 = _this.F24()
			}
			var _325_v93 _dafny.Array
			_ = _325_v93
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_8
			var _nw69 _dafny.Array
			_ = _nw69
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw69 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_326_i10 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("uqbkxbwhd")
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw69 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw69).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw69).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_325_v93 = _nw69
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_325_v93), 0))
			_ = _index57
			(_325_v93).ArraySet1(_210_v0, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_325_v93), 0))
			_ = _index58
			(_325_v93).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("easmd"), (_index58).Int())
		}
		r2 = (_this).F25()
		var _327_v94 _dafny.CodePoint
		_ = _327_v94
		_327_v94 = _dafny.CodePoint('x')
		_327_v94 = _327_v94
		r0 = _this.F24()
		var _328_v95 _dafny.Array
		_ = _328_v95
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw70
		_328_v95 = _nw70
		var _329_v96 _dafny.Set
		_ = _329_v96
		_329_v96 = _dafny.SetOf(_328_v95, _328_v95, _328_v95)
		var _330_v97 T0
		_ = _330_v97
		var _nw71 *C0 = New_C0_()
		_ = _nw71
		_nw71.Ctor__((_dafny.SetOf(_328_v95, _328_v95)).Union(_329_v96), _262_v36, (_263_v37).Cardinality())
		_330_v97 = _nw71
		r1 = _330_v97
		r2 = !(_this.F24())
		var _331_v98 _dafny.Map
		_ = _331_v98
		_331_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-174), Companion_Default___.SafeDivisionInt(_262_v36, _dafny.IntOfInt64(-599)))
		var _332_v99 _dafny.Map
		_ = _332_v99
		_332_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v35, _dafny.SetOf((_this).F25(), (_this).F25()))
		r3 = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if (_331_v98).Contains(_262_v36) {
				return (_331_v98).Get(_262_v36).(_dafny.Int)
			}
			return Companion_Default___.SafeModuloInt((_332_v99).Cardinality(), _262_v36)
		})()))
		return r0, r1, r2, r3
	}
}
func (_this *C1) M8(p0 bool, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) {
	{
		(globalState).F8 = p1
		(globalState).F6 = Companion_Default___.Fm29(_dafny.IntOfInt64(856), p1, (_this).F25(), globalState)
		var _333_v0 _dafny.Array
		_ = _333_v0
		var _nwElement0_10 bool = p3
		_ = _nwElement0_10
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(6))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_10, 0)
		(_nw72).ArraySet1(_this.F24(), 1)
		(_nw72).ArraySet1(p3, 2)
		(_nw72).ArraySet1(_this.F24(), 3)
		(_nw72).ArraySet1((_this).F25(), 4)
		(_nw72).ArraySet1((_this).F25(), 5)
		_333_v0 = _nw72
		var _334_v1 _dafny.Set
		_ = _334_v1
		_334_v1 = _dafny.SetOf(_333_v0, _333_v0)
		var _335_v2 *C0
		_ = _335_v2
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(_334_v1, (_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), _dafny.IntOfInt64(372))
		_335_v2 = _nw73
		var _336_v3 _dafny.Set
		_ = _336_v3
		_336_v3 = _dafny.SetOf(p0)
		var _337_i0 _dafny.Int
		_ = _337_i0
		_337_i0 = _dafny.Zero
		{
			for (_336_v3).IsProperSubsetOf((_dafny.SetOf((_this).F25())).Union(_336_v3)) {
				{
					if (_337_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_337_i0 = (_337_i0).Plus(_dafny.One)
					var _338_v4 _dafny.Sequence
					_ = _338_v4
					_338_v4 = _dafny.UnicodeSeqOfUtf8Bytes("helwwm")
					var _339_v5 _dafny.Sequence
					_ = _339_v5
					_339_v5 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-671))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_340_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('n')
					})))
					var _341_v6 _dafny.Map
					_ = _341_v6
					_341_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), p1)
					var _342_v7 _dafny.Sequence
					_ = _342_v7
					_342_v7 = _dafny.SeqOf((_341_v6).Cardinality())
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_333_v0), 0))
					_ = _index59
					(_333_v0).ArraySet1(_dafny.Companion_Sequence_.Equal(_338_v4, (_339_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_342_v7).Cardinality()), _dafny.IntOfUint32((_339_v5).Cardinality()))).Uint32()).(_dafny.Sequence)), (_index59).Int())
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_333_v0), 0))
					_ = _index60
					(_333_v0).ArraySet1(true, (_index60).Int())
					(globalState).F9 = !(!(_this.F24()))
					(globalState).F15 = (_333_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_333_v0), 0))).Int()).(bool)
					var _343_v8 _dafny.Array
					_ = _343_v8
					var _len0_9 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_9
					var _nw74 _dafny.Array
					_ = _nw74
					if _len0_9.Cmp(_dafny.Zero) == 0 {
						_nw74 = _dafny.NewArray(_len0_9)
					} else {
						var _init9 func(_dafny.Int) bool = (func(_344_p0 bool) func(_dafny.Int) bool {
							return func(_345_i2 _dafny.Int) bool {
								return !(_344_p0)
							}
						})(p0)
						_ = _init9
						var _element0_9 = _init9(_dafny.Zero)
						_ = _element0_9
						_nw74 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
						(_nw74).ArraySet1(_element0_9, 0)
						var _nativeLen0_9 = (_len0_9).Int()
						_ = _nativeLen0_9
						for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
							(_nw74).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
						}
					}
					_343_v8 = _nw74
					var _346_v9 _dafny.Map
					_ = _346_v9
					_346_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pfq"), _343_v8)
					_346_v9 = ((_346_v9).Merge(_346_v9)).Merge(_346_v9)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		(globalState).F9 = p3
		var _347_v10 _dafny.Map
		_ = _347_v10
		_347_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _335_v2.F31)
		var _source3 D0 = Companion_D0_.Create_DC1_(!((_347_v10).Equals(_347_v10)), p1)
		_ = _source3
		if _source3.Is_DC1() {
			var _348___mcc_h0 bool = _source3.Get_().(D0_DC1).Cf1
			_ = _348___mcc_h0
			var _349___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
			_ = _349___mcc_h1
			var _350_cf2 _dafny.Int = _349___mcc_h1
			_ = _350_cf2
			var _351_cf1 bool = _348___mcc_h0
			_ = _351_cf1
			(globalState).F1 = p3
			var _352_v11 _dafny.Map
			_ = _352_v11
			_352_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_335_v2.F31, _335_v2.F31, _dafny.IntOfInt64(483), _350_cf2), false)
			_352_v11 = _352_v11
			var _353_v12 _dafny.Sequence
			_ = _353_v12
			_353_v12 = _dafny.UnicodeSeqOfUtf8Bytes("ovboiepi")
			(globalState).F5 = (_dafny.Zero).Minus(((_this).Fm5(_350_cf2, p1, globalState)).Times(Companion_Default___.Fm1(_dafny.IntOfUint32((_353_v12).Cardinality()), (_this).F25(), globalState)))
			(_335_v2).F31 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gqn")).Cardinality()), _335_v2.F31)).Plus(_335_v2.F31)
		} else if _source3.Is_DC0() {
			var _354___mcc_h2 bool = _source3.Get_().(D0_DC0).Cf0
			_ = _354___mcc_h2
			var _355_cf0 bool = _354___mcc_h2
			_ = _355_cf0
			var _356_v13 D4
			_ = _356_v13
			_356_v13 = Companion_D4_.Create_DC12_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality()), p1)
			var _357_v14 _dafny.Array
			_ = _357_v14
			var _nwElement0_11 _dafny.Int = _dafny.IntOfInt64(985)
			_ = _nwElement0_11
			var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(3))
			_ = _nw75
			(_nw75).ArraySet1(_nwElement0_11, 0)
			(_nw75).ArraySet1((_this).Fm5((_356_v13).Dtor_cf19(), (func() _dafny.Int {
				if (_347_v10).Contains(_this.F24()) {
					return (_347_v10).Get(_this.F24()).(_dafny.Int)
				}
				return _335_v2.F31
			})(), globalState), 1)
			(_nw75).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p3, !((_this).F25()))).Cardinality()), 2)
			_357_v14 = _nw75
			var _358_v15 _dafny.Map
			_ = _358_v15
			_358_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(304), _357_v14)
			var _359_v16 _dafny.Sequence
			_ = _359_v16
			_359_v16 = _dafny.SeqOf(_357_v14, _357_v14, _357_v14, _357_v14, (func() _dafny.Array {
				if (_358_v15).Contains(p1) {
					return (_358_v15).Get(p1).(_dafny.Array)
				}
				return _357_v14
			})())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_333_v0), 0))
			_ = _index61
			(_333_v0).ArraySet1(Companion_Default___.Fm29(_335_v2.F31, (_347_v10).Cardinality(), !(_355_cf0), globalState), (_index61).Int())
			var _360_v17 _dafny.Set
			_ = _360_v17
			_360_v17 = _dafny.SetOf(_335_v2.F31, _335_v2.F31, (_dafny.Zero).Minus(p1))
			var _361_v18 _dafny.Map
			_ = _361_v18
			_361_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v2.F31, (_360_v17).Cardinality())
			var _362_v19 _dafny.Sequence
			_ = _362_v19
			_362_v19 = _dafny.SeqOf((_this).F25())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_333_v0), 0))
			_ = _index62
			var _rhs74 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_359_v16, _359_v16)
			_ = _rhs74
			var _rhs75 bool = _this.F24()
			_ = _rhs75
			var _rhs76 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, _335_v2.F31)
			_ = _rhs76
			var _rhs77 bool = ((func() bool {
				if _this.F24() {
					return p0
				}
				return false
			})()) == ((_dafny.IntOfInt64(335)).Cmp((func() _dafny.Int {
				if (_361_v18).Contains(_dafny.IntOfUint32((_362_v19).Cardinality())) {
					return (_361_v18).Get(_dafny.IntOfUint32((_362_v19).Cardinality())).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-467)
			})()) >= 0)
			_ = _rhs77
			var _rhs78 _dafny.Int = _335_v2.F31
			_ = _rhs78
			var _lhs62 *GlobalState = globalState
			_ = _lhs62
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			var _lhs64 _dafny.Array = _333_v0
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_333_v0), 0))
			_ = _lhs65
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			_359_v16 = _rhs74
			_lhs62.F6 = _rhs75
			_lhs63.F14 = _rhs76
			(_lhs64).ArraySet1(_rhs77, (_lhs65).Int())
			_lhs66.F8 = _rhs78
			var _363_v20 _dafny.Array
			_ = _363_v20
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw76
			_363_v20 = _nw76
			var _364_v21 _dafny.CodePoint
			_ = _364_v21
			_364_v21 = _dafny.CodePoint('d')
			var _365_v22 *C0
			_ = _365_v22
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(_dafny.SetOf(_363_v20), _dafny.IntOfUint32((Companion_Default___.Fm25((_362_v19).Select((Companion_Default___.SafeIndex((_361_v18).Cardinality(), _dafny.IntOfUint32((_362_v19).Cardinality()))).Uint32()).(bool), _this.F24(), _364_v21, globalState)).Cardinality()), p1)
			_365_v22 = _nw77
			(globalState).F6 = (_362_v19).Select((Companion_Default___.SafeIndex((p1).Minus(_365_v22.F31), _dafny.IntOfUint32((_362_v19).Cardinality()))).Uint32()).(bool)
			var _366_v24 _dafny.Array
			_ = _366_v24
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_10
			var _nw78 _dafny.Array
			_ = _nw78
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw78 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Sequence = (func(_367_v21 _dafny.CodePoint, _368_v22 *C0) func(_dafny.Int) _dafny.Sequence {
					return func(_369_i3 _dafny.Int) _dafny.Sequence {
						return (Companion_D9_.Create_DC25_(_dafny.SeqOf(func() _dafny.Map {
							var _coll14 = _dafny.NewMapBuilder()
							_ = _coll14
							for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v21, _367_v21)).Keys().Elements()); ; {
								_compr_14, _ok14 := _iter14()
								if !_ok14 {
									break
								}
								var _370_v23 _dafny.CodePoint
								_370_v23 = interface{}(_compr_14).(_dafny.CodePoint)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v21, _367_v21)).Contains(_370_v23) {
									_coll14.Add(_370_v23, _368_v22.F31)
								}
							}
							return _coll14.ToMap()
						}()))).Dtor_cf50()
					}
				})(_364_v21, _365_v22)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw78 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw78).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw78).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_366_v24 = _nw78
			var _371_v25 _dafny.Map
			_ = _371_v25
			_371_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v21, (_dafny.Zero).Minus(_365_v22.F31))
			var _372_v27 _dafny.Sequence
			_ = _372_v27
			_372_v27 = _dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('f'))
			var _373_v28 _dafny.Sequence
			_ = _373_v28
			_373_v28 = _dafny.SeqOf(_371_v25, _371_v25, _371_v25, func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_372_v27).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _374_v26 _dafny.CodePoint
					_374_v26 = interface{}(_compr_15).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_372_v27, _374_v26) {
						_coll15.Add(_374_v26, _335_v2.F31)
					}
				}
				return _coll15.ToMap()
			}())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_366_v24), 0))
			_ = _index63
			(_366_v24).ArraySet1(_373_v28, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_366_v24), 0))
			_ = _index64
			(_366_v24).ArraySet1(_373_v28, (_index64).Int())
		} else {
			var _375___mcc_h3 D0 = _source3.Get_().(D0_DC2).Cf3
			_ = _375___mcc_h3
			var _376_cf3 D0 = _375___mcc_h3
			_ = _376_cf3
			var _377_v29 _dafny.Sequence
			_ = _377_v29
			_377_v29 = _dafny.SeqOf(_335_v2.F31)
			var _378_v30 *C0
			_ = _378_v30
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__((_335_v2).F30(), _335_v2.F31, _dafny.IntOfUint32((_377_v29).Cardinality()))
			_378_v30 = _nw79
			var _379_v31 _dafny.MultiSet
			_ = _379_v31
			_379_v31 = _dafny.MultiSetOf(p0, p0, p3)
			var _380_v32 _dafny.Map
			_ = _380_v32
			_380_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _379_v31)
			var _381_v33 _dafny.Sequence
			_ = _381_v33
			_381_v33 = _dafny.UnicodeSeqOfUtf8Bytes("vb")
			_380_v32 = (_380_v32).Update((_335_v2).Fm3(_381_v33, (_378_v30).Fm3(_381_v33, (_this).F25(), p1, _dafny.SeqOf(_dafny.IntOfUint32((_381_v33).Cardinality())), globalState), _378_v30.F31, _377_v29, globalState), _379_v31)
			var _382_v34 _dafny.MultiSet
			_ = _382_v34
			_382_v34 = _dafny.MultiSetOf(_336_v3)
			var _383_v35 _dafny.MultiSet
			_ = _383_v35
			_383_v35 = _dafny.MultiSetOf(_378_v30.F31)
			var _384_v36 _dafny.MultiSet
			_ = _384_v36
			_384_v36 = _dafny.MultiSetOf((_382_v34).Cardinality(), (func() _dafny.Int {
				if (_383_v35).Contains(_dafny.IntOfUint32((_377_v29).Cardinality())) {
					return (_383_v35).Multiplicity(_dafny.IntOfUint32((_377_v29).Cardinality()))
				}
				return _335_v2.F31
			})())
			var _385_v37 _dafny.Map
			_ = _385_v37
			_385_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_381_v33).Cardinality()), p0)
			var _386_v38 _dafny.Array
			_ = _386_v38
			var _nwElement0_12 _dafny.Int = _335_v2.F31
			_ = _nwElement0_12
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(10))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_12, 0)
			(_nw80).ArraySet1((_dafny.Zero).Minus(_335_v2.F31), 1)
			(_nw80).ArraySet1(Companion_Default___.SafeDivisionInt(_335_v2.F31, _335_v2.F31), 2)
			(_nw80).ArraySet1((func() _dafny.Int {
				if (_383_v35).Contains(_378_v30.F31) {
					return (_383_v35).Multiplicity(_378_v30.F31)
				}
				return (_dafny.Zero).Minus(_378_v30.F31)
			})(), 3)
			(_nw80).ArraySet1((_378_v30.F31).Plus(_335_v2.F31), 4)
			(_nw80).ArraySet1((_335_v2.F31).Plus((_dafny.Zero).Minus(_335_v2.F31)), 5)
			(_nw80).ArraySet1((p1).Times(p1), 6)
			(_nw80).ArraySet1((_335_v2).Fm2(globalState), 7)
			(_nw80).ArraySet1((_336_v3).Cardinality(), 8)
			(_nw80).ArraySet1(((_385_v37).Merge(_385_v37)).Cardinality(), 9)
			_386_v38 = _nw80
			_386_v38 = _386_v38
			_381_v33 = _381_v33
		}
	}
}
func (_this *C1) M9(p0 _dafny.CodePoint, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _387_v0 _dafny.Sequence
		_ = _387_v0
		_387_v0 = _dafny.UnicodeSeqOfUtf8Bytes("nyt")
		var _388_v1 _dafny.Map
		_ = _388_v1
		_388_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), p3)
		var _389_v2 D3
		_ = _389_v2
		_389_v2 = Companion_D3_.Create_DC9_((_388_v1).Cardinality())
		var _hi2 _dafny.Int = (_389_v2).Dtor_cf15()
		_ = _hi2
		for _390_i0 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_387_v0, _387_v0)).Cardinality()); _390_i0.Cmp(_hi2) < 0; _390_i0 = _390_i0.Plus(_dafny.One) {
			var _391_v3 _dafny.Sequence
			_ = _391_v3
			_391_v3 = _dafny.SeqOf((_this).F25())
			(globalState).F9 = Companion_Default___.Fm29(_dafny.IntOfInt64(444), _390_i0, (_391_v3).Select((Companion_Default___.SafeIndex(_390_i0, _dafny.IntOfUint32((_391_v3).Cardinality()))).Uint32()).(bool), globalState)
			var _392_v4 _dafny.Map
			_ = _392_v4
			_392_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _393_v5 _dafny.Map
			_ = _393_v5
			_393_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v4, (_this).F25())
			(_this).M8(_this.F24(), p3, (_393_v5).Update(_392_v4, _this.F24()), Companion_Default___.Fm29(_390_i0, (_dafny.Zero).Minus(_390_i0), p2, globalState), globalState)
			var _394_v6 _dafny.Map
			_ = _394_v6
			_394_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24())
			var _395_v7 _dafny.Map
			_ = _395_v7
			_395_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_394_v6, p3)
			(_this).F24_set_(Companion_Default___.Fm29((_dafny.Zero).Minus(_390_i0), (func() _dafny.Int {
				if (_395_v7).Contains(_394_v6) {
					return (_395_v7).Get(_394_v6).(_dafny.Int)
				}
				return _390_i0
			})(), (_this).F25(), globalState))
			var _396_v8 _dafny.Array
			_ = _396_v8
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw81
			_396_v8 = _nw81
			var _397_v9 _dafny.Set
			_ = _397_v9
			_397_v9 = _dafny.SetOf(_396_v8)
			var _398_v10 *C0
			_ = _398_v10
			var _nw82 *C0 = New_C0_()
			_ = _nw82
			_nw82.Ctor__((func() _dafny.Set {
				if p1 {
					return _397_v9
				}
				return _397_v9
			})(), _390_i0, (_dafny.IntOfUint32((_387_v0).Cardinality())).Minus(_dafny.IntOfUint32((_387_v0).Cardinality())))
			_398_v10 = _nw82
		}
		var _399_v11 _dafny.Map
		_ = _399_v11
		_399_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _400_v12 _dafny.Array
		_ = _400_v12
		var _nwElement0_13 _dafny.Int = p3
		_ = _nwElement0_13
		var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
		_ = _nw83
		(_nw83).ArraySet1(_nwElement0_13, 0)
		(_nw83).ArraySet1(p3, 1)
		(_nw83).ArraySet1(_dafny.IntOfUint32((_387_v0).Cardinality()), 2)
		(_nw83).ArraySet1(_dafny.IntOfInt64(-13), 3)
		(_nw83).ArraySet1(p3, 4)
		(_nw83).ArraySet1(p3, 5)
		(_nw83).ArraySet1(p3, 6)
		(_nw83).ArraySet1(p3, 7)
		(_nw83).ArraySet1(_dafny.IntOfUint32((_387_v0).Cardinality()), 8)
		(_nw83).ArraySet1(((_399_v11).Cardinality()).Times(p3), 9)
		_400_v12 = _nw83
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))
		_ = _index65
		(_400_v12).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_387_v0).Cardinality())), (_index65).Int())
		var _401_v13 _dafny.Sequence
		_ = _401_v13
		_401_v13 = _dafny.SeqOf(_this.F24(), p2)
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))
		_ = _index66
		(_400_v12).ArraySet1((p3).Minus(_dafny.IntOfUint32((_401_v13).Cardinality())), (_index66).Int())
		(globalState).F14 = (_dafny.Zero).Minus(p3)
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))
		_ = _index67
		(_400_v12).ArraySet1((_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int), (_index67).Int())
		var _402_v14 _dafny.Sequence
		_ = _402_v14
		_402_v14 = _dafny.SeqOf(p3)
		var _403_v15 _dafny.MultiSet
		_ = _403_v15
		_403_v15 = _dafny.MultiSetOf(p3, (_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int), (_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int))
		var _404_v16 _dafny.Array
		_ = _404_v16
		var _nwElement0_14 bool = p2
		_ = _nwElement0_14
		var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(23))
		_ = _nw84
		(_nw84).ArraySet1(_nwElement0_14, 0)
		(_nw84).ArraySet1(false, 1)
		(_nw84).ArraySet1(p1, 2)
		(_nw84).ArraySet1(_this.F24(), 3)
		(_nw84).ArraySet1((_this).F25(), 4)
		(_nw84).ArraySet1(!(p1), 5)
		(_nw84).ArraySet1((_this).F25(), 6)
		(_nw84).ArraySet1(p2, 7)
		(_nw84).ArraySet1(_this.F24(), 8)
		(_nw84).ArraySet1((_this).F25(), 9)
		(_nw84).ArraySet1((_this).F25(), 10)
		(_nw84).ArraySet1(p1, 11)
		(_nw84).ArraySet1(p1, 12)
		(_nw84).ArraySet1(p2, 13)
		(_nw84).ArraySet1((_this).F25(), 14)
		(_nw84).ArraySet1((_this).F25(), 15)
		(_nw84).ArraySet1(_this.F24(), 16)
		(_nw84).ArraySet1((_this).F25(), 17)
		(_nw84).ArraySet1(p2, 18)
		(_nw84).ArraySet1(p2, 19)
		(_nw84).ArraySet1(_this.F24(), 20)
		(_nw84).ArraySet1(p1, 21)
		(_nw84).ArraySet1((_this).F25(), 22)
		_404_v16 = _nw84
		var _405_v17 _dafny.Sequence
		_ = _405_v17
		_405_v17 = _dafny.SeqOf(_404_v16, _404_v16)
		var _406_v18 _dafny.MultiSet
		_ = _406_v18
		_406_v18 = _dafny.MultiSetOf(p2)
		var _407_v19 _dafny.Sequence
		_ = _407_v19
		_407_v19 = _dafny.SeqOf(_dafny.IntOfUint32((_402_v14).Cardinality()), (func() _dafny.Int {
			if (_403_v15).Contains(_dafny.IntOfUint32((_405_v17).Cardinality())) {
				return (_403_v15).Multiplicity(_dafny.IntOfUint32((_405_v17).Cardinality()))
			}
			return (_406_v18).Cardinality()
		})())
		var _408_v20 _dafny.Sequence
		_ = _408_v20
		_408_v20 = _dafny.SeqOf((_402_v14).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_402_v14).Cardinality()))).Uint32()).(_dafny.Int), (_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int))
		var _rhs79 _dafny.Int = (_dafny.Zero).Minus(p3)
		_ = _rhs79
		var _rhs80 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int)), p3)
		_ = _rhs80
		var _rhs81 bool = p1
		_ = _rhs81
		var _rhs82 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(812))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_409_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_410_i1 _dafny.Int) _dafny.CodePoint {
				return _409_p0
			}
		})(p0)))
		_ = _rhs82
		var _rhs83 bool = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
			if !(Companion_Default___.Fm29(_dafny.IntOfInt64(724), p3, (_this).F25(), globalState)) {
				return _408_v20
			}
			return _dafny.Companion_Sequence_.Update((Companion_D10_.Create_DC28_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_411_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality())
			})))).Dtor_cf52(), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32(((Companion_D10_.Create_DC28_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_412_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality())
			})))).Dtor_cf52()).Cardinality()))).Uint32(), p3)
		})(), _402_v14)
		_ = _rhs83
		var _lhs67 *GlobalState = globalState
		_ = _lhs67
		var _lhs68 *GlobalState = globalState
		_ = _lhs68
		var _lhs69 *C1 = _this
		_ = _lhs69
		var _lhs70 *GlobalState = globalState
		_ = _lhs70
		_lhs67.F14 = _rhs79
		_lhs68.F8 = _rhs80
		_lhs69.F24_set_(_rhs81)
		_387_v0 = _rhs82
		_lhs70.F15 = _rhs83
		(globalState).F15 = Companion_Default___.Fm29(p3, (p3).Times((_400_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_400_v12), 0))).Int()).(_dafny.Int)), (_this).F25(), globalState)
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f23 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f23 = _dafny.Zero
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

func (_this *C2) F23() _dafny.Int {
	return _this._f23
}
func (_this *C2) Ctor__(f23 _dafny.Int) {
	{
		(_this)._f23 = f23
	}
}
func (_this *C2) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F23()
	}
}
func (_this *C2) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(!(false), false), _dafny.SeqOf(!(false)))) || ((_dafny.MultiSetOf((_this).F23())).IsProperSubsetOf(_dafny.MultiSetOf((_this).F23(), (_this).F23())))
	}
}
func (_this *C2) Fm19(p0 _dafny.Map, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(_dafny.SeqOf(true, true, true), _dafny.SeqOf(false, true), _dafny.SeqOf(true))).Difference(_dafny.MultiSetOf(_dafny.SeqOf(false, false, true), _dafny.SeqOf(false)))
	}
}
func (_this *C2) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		r0 = (_this).F23()
		if (p0) || (p0) {
			var _413_v0 _dafny.Array
			_ = _413_v0
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_11
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_414_p1 bool) func(_dafny.Int) bool {
					return func(_415_i0 _dafny.Int) bool {
						return (func() bool {
							if _414_p1 {
								return _414_p1
							}
							return !(_414_p1)
						})()
					}
				})(p1)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw85 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw85).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw85).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_413_v0 = _nw85
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))
			_ = _index68
			(_413_v0).ArraySet1(p0, (_index68).Int())
			var _416_v1 _dafny.Sequence
			_ = _416_v1
			_416_v1 = _dafny.UnicodeSeqOfUtf8Bytes("qbhsiypk")
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))
			_ = _index69
			var _rhs84 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F23(), _dafny.IntOfInt64(982))
			_ = _rhs84
			var _rhs85 bool = p0
			_ = _rhs85
			var _rhs86 bool = false
			_ = _rhs86
			var _rhs87 _dafny.Sequence = _416_v1
			_ = _rhs87
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			var _lhs72 _dafny.Array = _413_v0
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))
			_ = _lhs73
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			_lhs71.F8 = _rhs84
			(_lhs72).ArraySet1(_rhs85, (_lhs73).Int())
			_lhs74.F1 = _rhs86
			_416_v1 = _rhs87
			var _417_v2 D3
			_ = _417_v2
			_417_v2 = Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-517))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_418_i1 _dafny.Int) _dafny.CodePoint {
				return (Companion_D1_.Create_DC3_(_dafny.CodePoint('m'))).Dtor_cf4()
			})))
			var _419_v3 _dafny.Map
			_ = _419_v3
			_419_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_413_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))).Int()).(bool))
			var _420_v4 _dafny.Set
			_ = _420_v4
			_420_v4 = _dafny.SetOf(p0, p0)
			var _rhs88 _dafny.Int = (((_419_v3).Merge(_419_v3)).Cardinality()).Minus((_this).F23())
			_ = _rhs88
			var _rhs89 D3 = _417_v2
			_ = _rhs89
			var _rhs90 _dafny.Int = ((_this).F23()).Minus((_this).F23())
			_ = _rhs90
			var _rhs91 bool = (p1) && (p0)
			_ = _rhs91
			var _rhs92 _dafny.Int = Companion_Default___.SafeDivisionInt(((_420_v4).Difference(_dafny.SetOf(false))).Cardinality(), (_this).F23())
			_ = _rhs92
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			_lhs75.F17 = _rhs88
			_417_v2 = _rhs89
			_lhs76.F14 = _rhs90
			_lhs77.F6 = _rhs91
			_lhs78.F17 = _rhs92
			(globalState).F14 = (_this).F23()
			var _421_v5 _dafny.Map
			_ = _421_v5
			_421_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_416_v1, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oqfaosxj")).Cardinality())).Cmp((_this).F23()) < 0)
			var _422_v6 _dafny.Set
			_ = _422_v6
			_422_v6 = _dafny.SetOf((_this).F23())
			var _423_v7 _dafny.Sequence
			_ = _423_v7
			_423_v7 = _dafny.SeqOf((_413_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))).Int()).(bool), (_413_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_413_v0), 0))).Int()).(bool))
			var _rhs93 _dafny.Map = _421_v5
			_ = _rhs93
			var _rhs94 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F23(), ((_this).F23()).Plus((_this).F23()))
			_ = _rhs94
			var _rhs95 _dafny.Int = (func() _dafny.Int {
				if (_422_v6).IsProperSubsetOf(_422_v6) {
					return (_this).F23()
				}
				return _dafny.IntOfUint32((_423_v7).Cardinality())
			})()
			_ = _rhs95
			var _rhs96 _dafny.Int = (_this).F23()
			_ = _rhs96
			var _lhs79 *GlobalState = globalState
			_ = _lhs79
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_421_v5 = _rhs93
			_lhs79.F17 = _rhs94
			_lhs80.F14 = _rhs95
			_lhs81.F17 = _rhs96
			var _424_v8 _dafny.Array
			_ = _424_v8
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
			_ = _nw86
			_424_v8 = _nw86
			var _425_v9 _dafny.Map
			_ = _425_v9
			_425_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_413_v0, p1)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_424_v8), 0))
			_ = _index70
			(_424_v8).ArraySet1(_425_v9, (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_424_v8), 0))
			_ = _index71
			(_424_v8).ArraySet1((Companion_D4_.Create_DC11_(_425_v9)).Dtor_cf18(), (_index71).Int())
		} else {
			var _426_v10 _dafny.Array
			_ = _426_v10
			var _nwElement0_15 bool = !(p1)
			_ = _nwElement0_15
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(23))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_15, 0)
			(_nw87).ArraySet1(p1, 1)
			(_nw87).ArraySet1(p0, 2)
			(_nw87).ArraySet1(p1, 3)
			(_nw87).ArraySet1(p1, 4)
			(_nw87).ArraySet1(false, 5)
			(_nw87).ArraySet1(p0, 6)
			(_nw87).ArraySet1(!(p1), 7)
			(_nw87).ArraySet1(p0, 8)
			(_nw87).ArraySet1(p0, 9)
			(_nw87).ArraySet1(!(false), 10)
			(_nw87).ArraySet1(p0, 11)
			(_nw87).ArraySet1(p0, 12)
			(_nw87).ArraySet1(p0, 13)
			(_nw87).ArraySet1(p1, 14)
			(_nw87).ArraySet1(p1, 15)
			(_nw87).ArraySet1(p1, 16)
			(_nw87).ArraySet1(p1, 17)
			(_nw87).ArraySet1(p1, 18)
			(_nw87).ArraySet1(p0, 19)
			(_nw87).ArraySet1(p1, 20)
			(_nw87).ArraySet1(p0, 21)
			(_nw87).ArraySet1(p0, 22)
			_426_v10 = _nw87
			var _427_v11 _dafny.Sequence
			_ = _427_v11
			_427_v11 = _dafny.UnicodeSeqOfUtf8Bytes("wt")
			var _428_v12 _dafny.Map
			_ = _428_v12
			_428_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _427_v11)
			var _429_v13 _dafny.MultiSet
			_ = _429_v13
			_429_v13 = _dafny.MultiSetOf(_427_v11, (func() _dafny.Sequence {
				if (_428_v12).Contains(p1) {
					return (_428_v12).Get(p1).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-505))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_430_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))
			})(), _427_v11)
			var _431_v14 _dafny.Sequence
			_ = _431_v14
			_431_v14 = _dafny.SeqOf(_429_v13, _429_v13, _429_v13)
			var _432_v15 _dafny.Map
			_ = _432_v15
			_432_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(341), ((_431_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(861))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_433_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()), _dafny.IntOfUint32((_431_v14).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
			var _434_v16 _dafny.Map
			_ = _434_v16
			_434_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_v10, _432_v15)
			_434_v16 = (_434_v16).Merge(_434_v16)
			(globalState).F1 = false
			var _435_v17 _dafny.Array
			_ = _435_v17
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw88
			_435_v17 = _nw88
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_435_v17), 0))
			_ = _index72
			(_435_v17).ArraySet1(((_this).F23()).Minus((_this).F23()), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_435_v17), 0))
			_ = _index73
			(_435_v17).ArraySet1(Companion_Default___.SafeModuloInt((_this).F23(), ((_this).F23()).Times((_this).F23())), (_index73).Int())
			(globalState).F15 = p1
			var _436_v18 _dafny.Map
			_ = _436_v18
			_436_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
			var _437_v19 _dafny.Map
			_ = _437_v19
			_437_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(431))
			var _438_v20 D1
			_ = _438_v20
			_438_v20 = Companion_D1_.Create_DC5_(p0, _436_v18, (_this).F23(), (func() _dafny.Int {
				if (_432_v15).Contains((_this).F23()) {
					return (_432_v15).Get((_this).F23()).(_dafny.Int)
				}
				return (_435_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_435_v17), 0))).Int()).(_dafny.Int)
			})(), _437_v19)
			var _439_v21 _dafny.Map
			_ = _439_v21
			_439_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_435_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_435_v17), 0))).Int()).(_dafny.Int), _438_v20)
			_439_v21 = (_439_v21).Update((_435_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_435_v17), 0))).Int()).(_dafny.Int), _438_v20)
		}
		var _440_v22 _dafny.Map
		_ = _440_v22
		_440_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p0)), p0)
		var _441_v23 D3
		_ = _441_v23
		_441_v23 = Companion_D3_.Create_DC10_(!(p1), (_440_v22).Merge(_440_v22))
		var _442_v24 _dafny.CodePoint
		_ = _442_v24
		_442_v24 = _dafny.CodePoint('g')
		var _443_v26 _dafny.Map
		_ = _443_v26
		_443_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(240), _dafny.IntOfInt64(994))); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _444_v25 _dafny.Int
				_444_v25 = interface{}(_compr_16).(_dafny.Int)
				if ((_dafny.IntOfInt64(240)).Cmp(_444_v25) <= 0) && ((_444_v25).Cmp(_dafny.IntOfInt64(994)) < 0) {
					_coll16.Add((_444_v25).Minus((_this).F23()), (_this).F23())
				}
			}
			return _coll16.ToMap()
		}(), (_this).F23())
		var _445_v27 D5
		_ = _445_v27
		_445_v27 = Companion_D5_.Create_DC13_(_443_v26)
		var _446_v28 _dafny.Sequence
		_ = _446_v28
		_446_v28 = _dafny.SeqOf(false)
		var _447_v29 _dafny.Sequence
		_ = _447_v29
		_447_v29 = _dafny.UnicodeSeqOfUtf8Bytes("ngg")
		var _rhs97 _dafny.Int = ((_445_v27).Dtor_cf21()).Cardinality()
		_ = _rhs97
		var _rhs98 bool = ((_446_v28).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_446_v28).Cardinality()))).Uint32()).(bool)) && (!(_dafny.Companion_Sequence_.IsProperPrefixOf(_447_v29, _447_v29)))
		_ = _rhs98
		var _rhs99 D3 = _441_v23
		_ = _rhs99
		var _rhs100 _dafny.CodePoint = _442_v24
		_ = _rhs100
		var _lhs82 *GlobalState = globalState
		_ = _lhs82
		var _lhs83 *GlobalState = globalState
		_ = _lhs83
		_lhs82.F8 = _rhs97
		_lhs83.F15 = _rhs98
		_441_v23 = _rhs99
		_442_v24 = _rhs100
		_447_v29 = _dafny.Companion_Sequence_.Concatenate(_447_v29, Companion_Default___.Fm20(p0, globalState))
		var _448_v30 _dafny.Map
		_ = _448_v30
		_448_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F23())
		var _449_v31 D1
		_ = _449_v31
		_449_v31 = Companion_D1_.Create_DC5_(true, _440_v22, (_this).F23(), (_this).F23(), _448_v30)
		var _450_v32 _dafny.Map
		_ = _450_v32
		_450_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (((_449_v31).Dtor_cf8()).Merge(_440_v22)).Cardinality())
		var _451_v33 _dafny.Sequence
		_ = _451_v33
		_451_v33 = _dafny.SeqOf((_this).F23(), (_this).F23(), (_this).F23())
		_450_v32 = (_450_v32).Update((_this).Fm3(_447_v29, p1, (_this).F23(), _451_v33, globalState), ((_this).F23()).Plus(Companion_Default___.Fm1((_this).F23(), p1, globalState)))
		var _452_v34 _dafny.MultiSet
		_ = _452_v34
		_452_v34 = _dafny.MultiSetOf(p0)
		var _453_v35 _dafny.Sequence
		_ = _453_v35
		_453_v35 = _dafny.SeqOf(_451_v33)
		var _454_v36 _dafny.MultiSet
		_ = _454_v36
		_454_v36 = _dafny.MultiSetOf(_dafny.IntOfInt64(-143))
		var _455_v37 _dafny.Set
		_ = _455_v37
		_455_v37 = _dafny.SetOf((_454_v36).Cardinality())
		var _rhs101 bool = (_this).Fm3((func() _dafny.Sequence {
			if p1 {
				return _447_v29
			}
			return _447_v29
		})(), true, (_this).F23(), (_453_v35).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_453_v35).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)
		_ = _rhs101
		var _rhs102 bool = (_455_v37).IsSubsetOf(_455_v37)
		_ = _rhs102
		var _rhs103 _dafny.MultiSet = (_dafny.MultiSetOf(p0, p1)).Union((_dafny.MultiSetOf(p1, p0)).Update(!(p0), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))))
		_ = _rhs103
		var _lhs84 *GlobalState = globalState
		_ = _lhs84
		r2 = _rhs101
		_lhs84.F6 = _rhs102
		_452_v34 = _rhs103
		r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(30), (_this).F23())
		var _456_v38 _dafny.Array
		_ = _456_v38
		var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw89
		_456_v38 = _nw89
		r1 = _456_v38
		r2 = (_dafny.IntOfInt64(384)).Cmp((_this).F23()) < 0
		return r0, r1, r2
	}
}
func (_this *C2) M6(p0 bool, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) (D0, _dafny.Int, _dafny.Int) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		(globalState).F9 = p2
		var _457_v0 _dafny.Sequence
		_ = _457_v0
		_457_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wqyjmxlhd")
		var _hi3 _dafny.Int = (_this).Fm2(globalState)
		_ = _hi3
		for _458_i0 := ((_dafny.Zero).Minus(_dafny.IntOfUint32((_457_v0).Cardinality()))).Plus((_dafny.MultiSetOf(Companion_D3_.Create_DC9_(p3))).Cardinality()); _458_i0.Cmp(_hi3) < 0; _458_i0 = _458_i0.Plus(_dafny.One) {
			var _459_v1 _dafny.Map
			_ = _459_v1
			_459_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_458_i0, p3)).Cardinality())
			var _460_v2 D1
			_ = _460_v2
			_460_v2 = Companion_D1_.Create_DC5_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2), _458_i0, p3, _459_v1)
			var _461_v3 _dafny.Map
			_ = _461_v3
			_461_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_460_v2).Dtor_cf7(), p1)
			var _462_v4 _dafny.Map
			_ = _462_v4
			_462_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_460_v2, _458_i0)
			var _463_v5 _dafny.Sequence
			_ = _463_v5
			_463_v5 = _dafny.SeqOf(_462_v4, _462_v4, _462_v4)
			var _464_v6 _dafny.Sequence
			_ = _464_v6
			_464_v6 = _dafny.SeqOf(_463_v5, _463_v5, _463_v5)
			var _465_v7 _dafny.Map
			_ = _465_v7
			_465_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v3).Cardinality(), _dafny.Companion_Sequence_.Concatenate((_464_v6).Select((Companion_Default___.SafeIndex(_458_i0, _dafny.IntOfUint32((_464_v6).Cardinality()))).Uint32()).(_dafny.Sequence), (_464_v6).Select((Companion_Default___.SafeIndex(_458_i0, _dafny.IntOfUint32((_464_v6).Cardinality()))).Uint32()).(_dafny.Sequence)))
			_465_v7 = (_465_v7).Update((_this).F23(), _463_v5)
			_459_v1 = (_459_v1).Update(!(p2), _458_i0)
			var _466_v8 D3
			_ = _466_v8
			_466_v8 = Companion_D3_.Create_DC8_(_457_v0)
			var _467_v9 _dafny.CodePoint
			_ = _467_v9
			_467_v9 = _dafny.CodePoint('i')
			var _468_v10 _dafny.Sequence
			_ = _468_v10
			_468_v10 = _dafny.SeqOf(p2, p0)
			var _469_v11 _dafny.Set
			_ = _469_v11
			_469_v11 = _dafny.SetOf(p0)
			var _470_v12 _dafny.MultiSet
			_ = _470_v12
			_470_v12 = _dafny.MultiSetOf(_458_i0)
			var _471_v13 _dafny.Array
			_ = _471_v13
			var _nwElement0_16 bool = p2
			_ = _nwElement0_16
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(29))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_16, 0)
			(_nw90).ArraySet1((_this).Fm3((_466_v8).Dtor_cf14(), p0, _458_i0, p1, globalState), 1)
			(_nw90).ArraySet1(p2, 2)
			(_nw90).ArraySet1(p2, 3)
			(_nw90).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_457_v0, _dafny.Companion_Sequence_.Update(_457_v0, (Companion_Default___.SafeIndex(_458_i0, _dafny.IntOfUint32((_457_v0).Cardinality()))).Uint32(), _467_v9)), 4)
			(_nw90).ArraySet1(p0, 5)
			(_nw90).ArraySet1(p0, 6)
			(_nw90).ArraySet1(p0, 7)
			(_nw90).ArraySet1(p0, 8)
			(_nw90).ArraySet1(!(p0), 9)
			(_nw90).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_468_v10, _dafny.SeqOf((_468_v10).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1((_this).F23(), !(p0), globalState), _dafny.IntOfUint32((_468_v10).Cardinality()))).Uint32()).(bool), p2)), 10)
			(_nw90).ArraySet1(p0, 11)
			(_nw90).ArraySet1(!(p2), 12)
			(_nw90).ArraySet1(!((p3).Cmp((_this).F23()) < 0), 13)
			(_nw90).ArraySet1((_469_v11).IsSubsetOf(_469_v11), 14)
			(_nw90).ArraySet1((func() bool {
				if p2 {
					return !((_468_v10).Select((Companion_Default___.SafeIndex(_458_i0, _dafny.IntOfUint32((_468_v10).Cardinality()))).Uint32()).(bool))
				}
				return !(!(p0))
			})(), 15)
			(_nw90).ArraySet1(p2, 16)
			(_nw90).ArraySet1(p0, 17)
			(_nw90).ArraySet1(p2, 18)
			(_nw90).ArraySet1(p2, 19)
			(_nw90).ArraySet1((_this).Fm3(_457_v0, p0, _458_i0, p1, globalState), 20)
			(_nw90).ArraySet1(p0, 21)
			(_nw90).ArraySet1(true, 22)
			(_nw90).ArraySet1(p2, 23)
			(_nw90).ArraySet1(p0, 24)
			(_nw90).ArraySet1(p0, 25)
			(_nw90).ArraySet1(p0, 26)
			(_nw90).ArraySet1((_470_v12).IsDisjointFrom(_470_v12), 27)
			(_nw90).ArraySet1(p0, 28)
			_471_v13 = _nw90
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))
			_ = _index74
			(_471_v13).ArraySet1(true, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))
			_ = _index75
			(_471_v13).ArraySet1(p2, (_index75).Int())
			var _472_v14 _dafny.Set
			_ = _472_v14
			_472_v14 = _dafny.SetOf(_dafny.IntOfInt64(-582), _458_i0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(366))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_473_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_474_i1 _dafny.Int) _dafny.Int {
					return _473_i0
				}
			})(_458_i0)))).Cardinality()), p3)
			var _475_v15 _dafny.Map
			_ = _475_v15
			_475_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(_dafny.SetOf((_this).F23(), p3)).Equals(_472_v14))
			if (func() bool {
				if (_475_v15).Contains((_dafny.IntOfUint32((_457_v0).Cardinality())).Cmp(_dafny.IntOfUint32((_457_v0).Cardinality())) <= 0) {
					return (_475_v15).Get((_dafny.IntOfUint32((_457_v0).Cardinality())).Cmp(_dafny.IntOfUint32((_457_v0).Cardinality())) <= 0).(bool)
				}
				return p0
			})() {
				var _476_v16 _dafny.Map
				_ = _476_v16
				_476_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_v9, (_this).F23())
				var _477_v17 _dafny.Sequence
				_ = _477_v17
				_477_v17 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_v9, Companion_Default___.Fm1((_this).F23(), true, globalState)), _476_v16)
				_477_v17 = _dafny.Companion_Sequence_.Concatenate(_477_v17, _477_v17)
				var _478_v18 _dafny.Array
				_ = _478_v18
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(14))
				_ = _nw91
				_478_v18 = _nw91
				var _479_v19 _dafny.Map
				_ = _479_v19
				_479_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v18, _467_v9)
				var _480_v20 D3
				_ = _480_v20
				_480_v20 = Companion_D3_.Create_DC9_(_458_i0)
				var _481_v21 _dafny.Array
				_ = _481_v21
				var _nwElement0_17 _dafny.Int = _dafny.IntOfInt64(-441)
				_ = _nwElement0_17
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_17, 0)
				(_nw92).ArraySet1((_this).Fm2(globalState), 1)
				(_nw92).ArraySet1(p3, 2)
				(_nw92).ArraySet1((_this).F23(), 3)
				(_nw92).ArraySet1(p3, 4)
				(_nw92).ArraySet1(_458_i0, 5)
				(_nw92).ArraySet1((_this).F23(), 6)
				(_nw92).ArraySet1((_this).F23(), 7)
				(_nw92).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 8)
				(_nw92).ArraySet1((_this).F23(), 9)
				(_nw92).ArraySet1((_this).F23(), 10)
				(_nw92).ArraySet1(_dafny.IntOfInt64(-676), 11)
				(_nw92).ArraySet1((_this).F23(), 12)
				(_nw92).ArraySet1(_458_i0, 13)
				(_nw92).ArraySet1(_458_i0, 14)
				(_nw92).ArraySet1((_this).F23(), 15)
				(_nw92).ArraySet1(p3, 16)
				(_nw92).ArraySet1((_dafny.Zero).Minus(_458_i0), 17)
				(_nw92).ArraySet1((_this).F23(), 18)
				(_nw92).ArraySet1((_this).F23(), 19)
				(_nw92).ArraySet1((_469_v11).Cardinality(), 20)
				(_nw92).ArraySet1((_this).F23(), 21)
				(_nw92).ArraySet1((_this).F23(), 22)
				(_nw92).ArraySet1((_this).F23(), 23)
				(_nw92).ArraySet1((_this).F23(), 24)
				(_nw92).ArraySet1(p3, 25)
				(_nw92).ArraySet1(p3, 26)
				(_nw92).ArraySet1(p3, 27)
				(_nw92).ArraySet1((_this).F23(), 28)
				_481_v21 = _nw92
				var _482_v22 _dafny.Map
				_ = _482_v22
				_482_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_480_v20, _481_v21)
				var _483_v23 D5
				_ = _483_v23
				_483_v23 = Companion_D5_.Create_DC15_(_dafny.IntOfInt64(292), (func() _dafny.CodePoint {
					if (_471_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))).Int()).(bool) {
						return (_457_v0).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_457_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _467_v9
				})(), _479_v19, (_482_v22).Merge(_482_v22))
				var _484_v24 _dafny.Array
				_ = _484_v24
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_12
				var _nw93 _dafny.Array
				_ = _nw93
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw93 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Sequence = (func(_485_p0 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_486_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_485_p0)
						}
					})(p0)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw93 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw93).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw93).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_484_v24 = _nw93
				var _487_v25 D6
				_ = _487_v25
				_487_v25 = Companion_D6_.Create_DC16_(_468_v10)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_484_v24), 0))
				_ = _index76
				(_484_v24).ArraySet1((_487_v25).Dtor_cf29(), (_index76).Int())
				var _488_v26 _dafny.Map
				_ = _488_v26
				_488_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if p2 {
						return _dafny.IntOfInt64(-978)
					}
					return (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-663), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int)
				})(), _467_v9)
				var _pat_let_tv1 = _479_v19
				_ = _pat_let_tv1
				var _pat_let_tv2 = _479_v19
				_ = _pat_let_tv2
				var _pat_let_tv3 = p3
				_ = _pat_let_tv3
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_484_v24), 0))
				_ = _index77
				var _rhs104 D5 = func(_pat_let2_0 D5) D5 {
					return func(_489_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let3_0 _dafny.Map) D5 {
							return func(_490_dt__update_hcf27_h0 _dafny.Map) D5 {
								return func(_pat_let4_0 _dafny.Int) D5 {
									return func(_491_dt__update_hcf25_h0 _dafny.Int) D5 {
										return Companion_D5_.Create_DC15_(_491_dt__update_hcf25_h0, (_489_dt__update__tmp_h0).Dtor_cf26(), _490_dt__update_hcf27_h0, (_489_dt__update__tmp_h0).Dtor_cf28())
									}(_pat_let4_0)
								}((_pat_let_tv3).Minus((_this).F23()))
							}(_pat_let3_0)
						}((_pat_let_tv1).Merge(_pat_let_tv2))
					}(_pat_let2_0)
				}(_483_v23)
				_ = _rhs104
				var _rhs105 _dafny.Array = _471_v13
				_ = _rhs105
				var _rhs106 _dafny.Array = _471_v13
				_ = _rhs106
				var _rhs107 _dafny.Sequence = _468_v10
				_ = _rhs107
				var _rhs108 _dafny.Map = (_488_v26).Merge(_488_v26)
				_ = _rhs108
				var _lhs85 _dafny.Array = _484_v24
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_484_v24), 0))
				_ = _lhs86
				_483_v23 = _rhs104
				_471_v13 = _rhs105
				_471_v13 = _rhs106
				(_lhs85).ArraySet1(_rhs107, (_lhs86).Int())
				_488_v26 = _rhs108
				(globalState).F5 = Companion_Default___.Fm1(p3, true, globalState)
				(globalState).F1 = (_469_v11).IsSubsetOf(_469_v11)
				(globalState).F9 = !((_471_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))).Int()).(bool))
			} else {
				var _492_v27 D6
				_ = _492_v27
				_492_v27 = Companion_D6_.Create_DC16_(_468_v10)
				_492_v27 = _492_v27
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))
				_ = _index78
				(_471_v13).ArraySet1(p2, (_index78).Int())
				var _493_v28 _dafny.Map
				_ = _493_v28
				_493_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.MultiSetOf((_468_v10).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p3), _dafny.IntOfUint32((_468_v10).Cardinality()))).Uint32()).(bool)))
				_493_v28 = (_493_v28).Update((_471_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))).Int()).(bool), Companion_Default___.Fm21(p2, true, (_this).F23(), globalState))
				var _494_v29 _dafny.Map
				_ = _494_v29
				_494_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_458_i0).Minus((_this).F23()), (func() _dafny.Int {
					if (_471_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_471_v13), 0))).Int()).(bool) {
						return p3
					}
					return _458_i0
				})())
				_494_v29 = (_494_v29).Update(p3, (_this).F23())
				(globalState).F6 = (!(p0)) == (false)
			}
		}
		var _495_v30 _dafny.Map
		_ = _495_v30
		_495_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (p3).Minus(p3))
		(globalState).F0 = _495_v30
		var _496_v31 _dafny.Array
		_ = _496_v31
		var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(9))
		_ = _nw94
		_496_v31 = _nw94
		var _497_v32 _dafny.Set
		_ = _497_v32
		_497_v32 = _dafny.SetOf(!(p0))
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_496_v31), 0))
		_ = _index79
		(_496_v31).ArraySet1(_497_v32, (_index79).Int())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_496_v31), 0))
		_ = _index80
		(_496_v31).ArraySet1(_497_v32, (_index80).Int())
		_495_v30 = (_495_v30).Merge(_495_v30)
		(globalState).F17 = (_this).F23()
		var _498_v33 D0
		_ = _498_v33
		_498_v33 = Companion_D0_.Create_DC1_(p2, p3)
		r0 = _498_v33
		r1 = p3
		var _499_v34 D0
		_ = _499_v34
		_499_v34 = Companion_D0_.Create_DC0_(p0)
		var _pat_let_tv4 = _457_v0
		_ = _pat_let_tv4
		r2 = func(_source4 D0) _dafny.Int {
			if _source4.Is_DC1() {
				var _500___mcc_h0 bool = _source4.Get_().(D0_DC1).Cf1
				_ = _500___mcc_h0
				var _501___mcc_h1 _dafny.Int = _source4.Get_().(D0_DC1).Cf2
				_ = _501___mcc_h1
				var _502_cf2 _dafny.Int = _501___mcc_h1
				_ = _502_cf2
				var _503_cf1 bool = _500___mcc_h0
				_ = _503_cf1
				return _dafny.IntOfUint32((_pat_let_tv4).Cardinality())
			} else if _source4.Is_DC0() {
				var _504___mcc_h2 bool = _source4.Get_().(D0_DC0).Cf0
				_ = _504___mcc_h2
				var _505_cf0 bool = _504___mcc_h2
				_ = _505_cf0
				return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-577))).Cardinality())
			} else {
				var _506___mcc_h3 D0 = _source4.Get_().(D0_DC2).Cf3
				_ = _506___mcc_h3
				var _507_cf3 D0 = _506___mcc_h3
				_ = _507_cf3
				return (_this).F23()
			}
		}(Companion_D0_.Create_DC2_(_499_v34))
		return r0, r1, r2
	}
}
func (_this *C2) M7(p0 _dafny.Array, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F14 = (_this).F23()
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))
		_ = _index81
		(p0).ArraySet1(p2, (_index81).Int())
		var _508_v0 _dafny.Sequence
		_ = _508_v0
		_508_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qiu")
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))
		_ = _index82
		(p0).ArraySet1((func() bool {
			if true {
				return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("mseqsel"), _508_v0)
			}
			return p1
		})(), (_index82).Int())
		(globalState).F17 = (_this).F23()
		var _509_v1 _dafny.MultiSet
		_ = _509_v1
		_509_v1 = _dafny.MultiSetOf(p1, p1, p1, p2)
		var _510_v2 _dafny.MultiSet
		_ = _510_v2
		_510_v2 = _dafny.MultiSetOf((_this).F23())
		var _511_v3 _dafny.Sequence
		_ = _511_v3
		_511_v3 = _dafny.SeqOf(p2, p2, p1)
		var _512_v4 _dafny.Map
		_ = _512_v4
		_512_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_510_v2).Cardinality(), p2)
		var _513_v5 _dafny.Map
		_ = _513_v5
		_513_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_512_v4).Cardinality(), _508_v0)
		var _514_v6 _dafny.Map
		_ = _514_v6
		_514_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
			if (_513_v5).Contains((_this).F23()) {
				return (_513_v5).Get((_this).F23()).(_dafny.Sequence)
			}
			return _508_v0
		})(), _511_v3)
		var _515_v7 _dafny.CodePoint
		_ = _515_v7
		_515_v7 = _dafny.CodePoint('j')
		var _516_v8 _dafny.Set
		_ = _516_v8
		_516_v8 = _dafny.SetOf(_515_v7)
		var _517_v9 _dafny.Array
		_ = _517_v9
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_13
		var _nw95 _dafny.Array
		_ = _nw95
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw95 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) D0 = (func(_518_p1 bool) func(_dafny.Int) D0 {
				return func(_519_i1 _dafny.Int) D0 {
					return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_518_p1, (_this).F23()))
				}
			})(p1)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw95 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw95).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw95).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_517_v9 = _nw95
		var _520_v10 _dafny.Map
		_ = _520_v10
		_520_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_517_v9, _515_v7)
		var _521_v11 D3
		_ = _521_v11
		_521_v11 = Companion_D3_.Create_DC9_((_this).F23())
		var _522_v12 _dafny.Array
		_ = _522_v12
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_14
		var _nw96 _dafny.Array
		_ = _nw96
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw96 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Int = func(_523_i2 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_523_i2, (_this).F23())
			}
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw96 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw96).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw96).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_522_v12 = _nw96
		var _524_v13 D5
		_ = _524_v13
		_524_v13 = Companion_D5_.Create_DC15_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_508_v0, (Companion_Default___.SafeIndex((_516_v8).Cardinality(), _dafny.IntOfUint32((_508_v0).Cardinality()))).Uint32(), _515_v7)).Cardinality()), _515_v7, _520_v10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_521_v11, _522_v12))
		var _525_v14 _dafny.Array
		_ = _525_v14
		var _nwElement0_18 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(223))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_526_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality())
		_ = _nwElement0_18
		var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
		_ = _nw97
		(_nw97).ArraySet1(_nwElement0_18, 0)
		(_nw97).ArraySet1((_this).F23(), 1)
		(_nw97).ArraySet1((_this).F23(), 2)
		(_nw97).ArraySet1((func() _dafny.Int {
			if (_509_v1).Contains(p2) {
				return (_509_v1).Multiplicity(p2)
			}
			return (_this).F23()
		})(), 3)
		(_nw97).ArraySet1((_510_v2).Cardinality(), 4)
		(_nw97).ArraySet1((Companion_Default___.Fm21(true, p2, (_this).F23(), globalState)).Cardinality(), 5)
		(_nw97).ArraySet1((_this).F23(), 6)
		(_nw97).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm20((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)).Cardinality())), 7)
		(_nw97).ArraySet1((_dafny.IntOfUint32((_508_v0).Cardinality())).Times((_this).F23()), 8)
		(_nw97).ArraySet1((_this).F23(), 9)
		(_nw97).ArraySet1((_this).F23(), 10)
		(_nw97).ArraySet1((_this).Fm2(globalState), 11)
		(_nw97).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_511_v3, (func() _dafny.Sequence {
			if (_514_v6).Contains(_508_v0) {
				return (_514_v6).Get(_508_v0).(_dafny.Sequence)
			}
			return _511_v3
		})())).Cardinality()), 12)
		(_nw97).ArraySet1((_524_v13).Dtor_cf25(), 13)
		(_nw97).ArraySet1((_this).F23(), 14)
		_525_v14 = _nw97
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))
		_ = _index83
		(_522_v12).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), (_index83).Int())
		var _527_v15 _dafny.Map
		_ = _527_v15
		_527_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(globalState), p0)
		var _528_v16 _dafny.Set
		_ = _528_v16
		_528_v16 = _dafny.SetOf((func() _dafny.Array {
			if (_527_v15).Contains((_this).F23()) {
				return (_527_v15).Get((_this).F23()).(_dafny.Array)
			}
			return p0
		})())
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))
		_ = _index84
		(_522_v12).ArraySet1((_dafny.Zero).Minus((((func() _dafny.Set {
			if p3 {
				return _dafny.SetOf(p0)
			}
			return _528_v16
		})()).Cardinality()).Minus((_this).F23())), (_index84).Int())
		var _529_i3 _dafny.Int
		_ = _529_i3
		_529_i3 = _dafny.Zero
		{
			for !_dafny.Companion_Sequence_.Contains(_508_v0, _515_v7) {
				{
					if (_529_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_529_i3 = (_529_i3).Plus(_dafny.One)
					if !(p2) || (p2) {
						var _530_v17 *C1
						_ = _530_v17
						var _nw98 *C1 = New_C1_()
						_ = _nw98
						_nw98.Ctor__(p3, false)
						_530_v17 = _nw98
						var _531_v18 *C0
						_ = _531_v18
						var _nw99 *C0 = New_C0_()
						_ = _nw99
						_nw99.Ctor__((_528_v16).Difference(_528_v16), (_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(262))
						_531_v18 = _nw99
						var _532_v19 _dafny.Map
						_ = _532_v19
						_532_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
						(globalState).F1 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29(_531_v18.F31, (_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), p3)).Merge(_532_v19)).Cardinality()).Cmp((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int)) <= 0
						var _533_v20 _dafny.Array
						_ = _533_v20
						var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
						_ = _nw100
						_533_v20 = _nw100
						var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_533_v20), 0))
						_ = _index85
						(_533_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_508_v0, _508_v0), (_index85).Int())
						var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_533_v20), 0))
						_ = _index86
						(_533_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_508_v0, Companion_Default___.Fm20((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)), (_index86).Int())
						var _534_v21 _dafny.Sequence
						_ = _534_v21
						_534_v21 = _dafny.SeqOf(_531_v18.F31, _dafny.IntOfInt64(715))
						var _535_v22 _dafny.Set
						_ = _535_v22
						_535_v22 = _dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool))
						var _536_v23 _dafny.Map
						_ = _536_v23
						_536_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_534_v21, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_534_v21).Cardinality()))).Uint32(), (_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int)), (_535_v22).IsDisjointFrom(_535_v22))
						var _537_v25 _dafny.Sequence
						_ = _537_v25
						_537_v25 = _dafny.SeqOf(_534_v21, _534_v21)
						_536_v23 = ((func() _dafny.Map {
							var _coll17 = _dafny.NewMapBuilder()
							_ = _coll17
							for _iter17 := _dafny.Iterate((_537_v25).Elements()); ; {
								_compr_17, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _538_v24 _dafny.Sequence
								_538_v24 = interface{}(_compr_17).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_537_v25, _538_v24) {
									_coll17.Add(_538_v24, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool))
								}
							}
							return _coll17.ToMap()
						}()).Merge(func() _dafny.Map {
							var _coll18 = _dafny.NewMapBuilder()
							_ = _coll18
							for _iter18 := _dafny.Iterate((_536_v23).Keys().Elements()); ; {
								_compr_18, _ok18 := _iter18()
								if !_ok18 {
									break
								}
								var _539_v26 _dafny.Sequence
								_539_v26 = interface{}(_compr_18).(_dafny.Sequence)
								if (_536_v23).Contains(_539_v26) {
									_coll18.Add(_539_v26, p1)
								}
							}
							return _coll18.ToMap()
						}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v21, p1))
					} else {
						var _540_v27 _dafny.Map
						_ = _540_v27
						_540_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-120))), (_dafny.MultiSetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), p3, p1, true, p3)).Cardinality())
						_540_v27 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
							if (_510_v2).Contains((_this).F23()) {
								return (_510_v2).Multiplicity((_this).F23())
							}
							return _dafny.IntOfInt64(-265)
						})(), _dafny.CodePoint('p'))).Cardinality(), Companion_Default___.Fm1(Companion_Default___.Fm1((_dafny.MultiSetFromSeq(_511_v3)).Cardinality(), p2, globalState), p2, globalState))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23()))
						r0 = !(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_508_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(511))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg26 _dafny.Int) interface{} {
								return coer26(arg26)
							}
						}(func(_541_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('m')
						}))), _515_v7))
						(globalState).F17 = Companion_Default___.SafeDivisionInt(((_this).F23()).Plus((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int)), (_this).F23())
						(globalState).F17 = _dafny.IntOfInt64(368)
						var _542_v30 _dafny.Array
						_ = _542_v30
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(18)
						_ = _len0_15
						var _nw101 _dafny.Array
						_ = _nw101
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw101 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) _dafny.Map = (func(_543_v12 _dafny.Array, _544_v7 _dafny.CodePoint, _545_v4 _dafny.Map) func(_dafny.Int) _dafny.Map {
								return func(_546_i5 _dafny.Int) _dafny.Map {
									return func() _dafny.Map {
										var _coll19 = _dafny.NewMapBuilder()
										_ = _coll19
										for _iter19 := _dafny.Iterate((func() _dafny.Set {
											var _coll20 = _dafny.NewBuilder()
											_ = _coll20
											for _iter20 := _dafny.Iterate((_545_v4).Keys().Elements()); ; {
												_compr_20, _ok20 := _iter20()
												if !_ok20 {
													break
												}
												var _547_v29 _dafny.Int
												_547_v29 = interface{}(_compr_20).(_dafny.Int)
												if (_545_v4).Contains(_547_v29) {
													_coll20.Add(Companion_Default___.SafeModuloInt(_547_v29, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qnqjr")).Cardinality())))
												}
											}
											return _coll20.ToSet()
										}()).Elements()); ; {
											_compr_19, _ok19 := _iter19()
											if !_ok19 {
												break
											}
											var _548_v28 _dafny.Int
											_548_v28 = interface{}(_compr_19).(_dafny.Int)
											if (func() _dafny.Set {
												var _coll21 = _dafny.NewBuilder()
												_ = _coll21
												for _iter21 := _dafny.Iterate((_545_v4).Keys().Elements()); ; {
													_compr_21, _ok21 := _iter21()
													if !_ok21 {
														break
													}
													var _549_v29 _dafny.Int
													_549_v29 = interface{}(_compr_21).(_dafny.Int)
													if (_545_v4).Contains(_549_v29) {
														_coll21.Add(Companion_Default___.SafeModuloInt(_549_v29, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qnqjr")).Cardinality())))
													}
												}
												return _coll21.ToSet()
											}()).Contains(_548_v28) {
												_coll19.Add(Companion_Default___.SafeDivisionInt(_548_v28, (_543_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_543_v12), 0))).Int()).(_dafny.Int)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v7, _544_v7)).Cardinality())
											}
										}
										return _coll19.ToMap()
									}()
								}
							})(_522_v12, _515_v7, _512_v4)
							_ = _init15
							var _element0_15 = _init15(_dafny.Zero)
							_ = _element0_15
							_nw101 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
							(_nw101).ArraySet1(_element0_15, 0)
							var _nativeLen0_15 = (_len0_15).Int()
							_ = _nativeLen0_15
							for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
								(_nw101).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
							}
						}
						_542_v30 = _nw101
						var _len0_16 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_16
						var _nw102 _dafny.Array
						_ = _nw102
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw102 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) _dafny.Map = (func(_550_v27 _dafny.Map) func(_dafny.Int) _dafny.Map {
								return func(_551_i6 _dafny.Int) _dafny.Map {
									return _550_v27
								}
							})(_540_v27)
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw102 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw102).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw102).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_542_v30 = _nw102
					}
					var _552_v31 _dafny.Map
					_ = _552_v31
					_552_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool))
					var _553_v32 _dafny.Set
					_ = _553_v32
					_553_v32 = _dafny.SetOf((_this).F23())
					var _554_v33 _dafny.Map
					_ = _554_v33
					_554_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_553_v32).Cardinality())
					var _555_v34 D1
					_ = _555_v34
					_555_v34 = Companion_D1_.Create_DC5_((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), _552_v31, _dafny.IntOfInt64(-897), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm30((func() _dafny.Int {
						if (_509_v1).Contains(p2) {
							return (_509_v1).Multiplicity(p2)
						}
						return _dafny.IntOfInt64(-629)
					})(), _508_v0, p3, (func() _dafny.Int {
						if (_509_v1).Contains(true) {
							return (_509_v1).Multiplicity(true)
						}
						return (func() _dafny.Int {
							if (_554_v33).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
								return (_554_v33).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_508_v0).Cardinality())
						})()
					})(), globalState), (Companion_Default___.SafeIndex((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm30((func() _dafny.Int {
						if (_509_v1).Contains(p2) {
							return (_509_v1).Multiplicity(p2)
						}
						return _dafny.IntOfInt64(-629)
					})(), _508_v0, p3, (func() _dafny.Int {
						if (_509_v1).Contains(true) {
							return (_509_v1).Multiplicity(true)
						}
						return (func() _dafny.Int {
							if (_554_v33).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
								return (_554_v33).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_508_v0).Cardinality())
						})()
					})(), globalState)).Cardinality()))).Uint32(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool))).Cardinality()), (_554_v33).Merge(_554_v33))
					_555_v34 = _555_v34
					(globalState).F8 = (_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int)
					(globalState).F8 = (_this).F23()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _556_v35 _dafny.Sequence
		_ = _556_v35
		_556_v35 = _dafny.SeqOf(_510_v2)
		var _557_v36 _dafny.Sequence
		_ = _557_v36
		_557_v36 = _dafny.SeqOf(_556_v35, _556_v35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-187))).Uint32(), func(coer27 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_558_v2 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_559_i7 _dafny.Int) _dafny.MultiSet {
				return _558_v2
			}
		})(_510_v2))))
		(globalState).F9 = _dafny.Companion_Sequence_.IsProperPrefixOf(_556_v35, (_557_v36).Select((Companion_Default___.SafeIndex((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_557_v36).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _560_v37 *C0
		_ = _560_v37
		var _nw103 *C0 = New_C0_()
		_ = _nw103
		_nw103.Ctor__(_528_v16, (_this).F23(), (_dafny.Zero).Minus((_522_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_522_v12), 0))).Int()).(_dafny.Int)))
		_560_v37 = _nw103
		var _561_v38 _dafny.Sequence
		_ = _561_v38
		_561_v38 = _dafny.SeqOf(_560_v37, _560_v37)
		r0 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _561_v38)).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((p0), 0))).Int()).(bool), _561_v38)).Contains(p3)
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f24 bool
	_f23 _dafny.Int
	_f25 bool
	_f28 bool
	_f29 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f24 = false
	_this._f23 = _dafny.Zero
	_this._f25 = false
	_this._f28 = false
	_this._f29 = _dafny.Zero
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

func (_this *C3) F24() bool {
	return _this._f24
}
func (_this *C3) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C3) F23() _dafny.Int {
	return _this._f23
}
func (_this *C3) F25() bool {
	return _this._f25
}
func (_this *C3) Ctor__(f28 bool, f29 _dafny.Int, f24 bool, f25 bool, f23 _dafny.Int) {
	{
		(_this)._f28 = f28
		(_this)._f29 = f29
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f23 = f23
	}
}
func (_this *C3) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F28())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())))
	}
}
func (_this *C3) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_this).F29())).Times((_this).F29())
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('u'), _dafny.CodePoint('m'))).Cardinality())).Plus((_this).F29())
	}
}
func (_this *C3) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Set {
			if (_this).F28() {
				return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("njpng"))
			}
			return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("arqurnxwi"))
		})()).IsProperSubsetOf((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("pfvvyx"), _dafny.UnicodeSeqOfUtf8Bytes("htve"))).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("exnwakoad"), _dafny.UnicodeSeqOfUtf8Bytes("jfhvlkukr"))))
	}
}
func (_this *C3) Fm16(p0 bool, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return !(!(!(((_this).F29()).Cmp((_this).F23()) != 0)) || ((_dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality())).Cmp((_this).F23()) <= 0))
	}
}
func (_this *C3) Fm17(p0 D1, globalState *GlobalState) bool {
	{
		return (_this).F25()
	}
}
func (_this *C3) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _562_v0 _dafny.Array
		_ = _562_v0
		var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
		_ = _nw104
		_562_v0 = _nw104
		var _563_v1 _dafny.CodePoint
		_ = _563_v1
		_563_v1 = _dafny.CodePoint('d')
		var _564_v2 D1
		_ = _564_v2
		_564_v2 = Companion_D1_.Create_DC3_(_563_v1)
		var _565_v3 D1
		_ = _565_v3
		_565_v3 = Companion_D1_.Create_DC6_(_564_v2)
		var _566_v4 _dafny.Array
		_ = _566_v4
		var _nwElement0_19 bool = (_this).F25()
		_ = _nwElement0_19
		var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(6))
		_ = _nw105
		(_nw105).ArraySet1(_nwElement0_19, 0)
		(_nw105).ArraySet1(_this.F24(), 1)
		(_nw105).ArraySet1((_this).F28(), 2)
		(_nw105).ArraySet1((_this).Fm17(_565_v3, globalState), 3)
		(_nw105).ArraySet1((_this).F25(), 4)
		(_nw105).ArraySet1(_this.F24(), 5)
		_566_v4 = _nw105
		var _567_v5 _dafny.Set
		_ = _567_v5
		_567_v5 = _dafny.SetOf(_566_v4)
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_562_v0), 0))
		_ = _index87
		(_562_v0).ArraySet1(_567_v5, (_index87).Int())
		var _568_v6 _dafny.Sequence
		_ = _568_v6
		_568_v6 = _dafny.SeqOf(_567_v5)
		var _569_v7 _dafny.Map
		_ = _569_v7
		_569_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _567_v5)
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_562_v0), 0))
		_ = _index88
		(_562_v0).ArraySet1(((_568_v6).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_568_v6).Cardinality()))).Uint32()).(_dafny.Set)).Difference((func() _dafny.Set {
			if (_569_v7).Contains(_dafny.IntOfInt64(20)) {
				return (_569_v7).Get(_dafny.IntOfInt64(20)).(_dafny.Set)
			}
			return _567_v5
		})()), (_index88).Int())
		var _570_v8 _dafny.MultiSet
		_ = _570_v8
		_570_v8 = _dafny.MultiSetOf((_this).F25(), ((_this).F28()) || (true))
		_570_v8 = _570_v8
		(globalState).F15 = (_this).F28()
		var _571_v9 _dafny.Sequence
		_ = _571_v9
		_571_v9 = _dafny.SeqOf((_this).F28())
		var _572_v10 _dafny.Sequence
		_ = _572_v10
		_572_v10 = _dafny.SeqOf((_this).F23())
		var _573_v11 _dafny.Sequence
		_ = _573_v11
		_573_v11 = _dafny.SeqOf(_572_v10)
		var _574_v12 _dafny.MultiSet
		_ = _574_v12
		_574_v12 = _dafny.MultiSetOf(_dafny.IntOfUint32((_573_v11).Cardinality()), (_this).F29())
		var _575_v13 _dafny.Map
		_ = _575_v13
		_575_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (func() _dafny.Int {
			if (_574_v12).Contains(_dafny.IntOfInt64(197)) {
				return (_574_v12).Multiplicity(_dafny.IntOfInt64(197))
			}
			return _dafny.IntOfInt64(410)
		})())
		var _576_v14 D1
		_ = _576_v14
		_576_v14 = Companion_D1_.Create_DC5_(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhhqehb")).Cardinality()), _dafny.IntOfUint32((_571_v9).Cardinality()), _575_v13)
		var _577_v15 _dafny.Sequence
		_ = _577_v15
		_577_v15 = _dafny.SeqOf((_576_v14).Dtor_cf10(), (_this).F23())
		(globalState).F8 = (_572_v10).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_this).F29()).Times((_this).F29())), _dafny.IntOfUint32((_572_v10).Cardinality()))).Uint32()).(_dafny.Int)
		var _578_v16 _dafny.Sequence
		_ = _578_v16
		_578_v16 = _dafny.SeqOf(_576_v14)
		r3 = (func() _dafny.Int {
			if (_this).F25() {
				return (func() _dafny.Int {
					if true {
						return _dafny.IntOfUint32((_578_v16).Cardinality())
					}
					return (_this).F23()
				})()
			}
			return (_this).F23()
		})()
		var _579_v17 _dafny.Array
		_ = _579_v17
		var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw106
		_579_v17 = _nw106
		var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_579_v17), 0))
		_ = _index89
		(_579_v17).ArraySet1(_566_v4, (_index89).Int())
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_566_v4), 0))
		_ = _index90
		(_566_v4).ArraySet1(((_this).F25()) || ((_571_v9).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_571_v9).Cardinality()))).Uint32()).(bool)), (_index90).Int())
		var _580_v18 D3
		_ = _580_v18
		_580_v18 = Companion_D3_.Create_DC9_((_this).F23())
		var _581_v19 _dafny.Array
		_ = _581_v19
		var _nwElement0_20 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm18((_this).F23(), (_this).F29(), _571_v9, _580_v18, globalState)).Cardinality())
		_ = _nwElement0_20
		var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(9))
		_ = _nw107
		(_nw107).ArraySet1(_nwElement0_20, 0)
		(_nw107).ArraySet1(((_this).F23()).Minus((_this).F23()), 1)
		(_nw107).ArraySet1((_this).F29(), 2)
		(_nw107).ArraySet1((_this).F23(), 3)
		(_nw107).ArraySet1((Companion_Default___.Fm1((_this).F29(), true, globalState)).Minus((_this).F29()), 4)
		(_nw107).ArraySet1((_this).F29(), 5)
		(_nw107).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F28(), (_this).F25())).Cardinality()), 6)
		(_nw107).ArraySet1((_this).F29(), 7)
		(_nw107).ArraySet1((_this).F29(), 8)
		_581_v19 = _nw107
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_581_v19), 0))
		_ = _index91
		(_581_v19).ArraySet1((_this).F23(), (_index91).Int())
		var _582_v20 _dafny.Set
		_ = _582_v20
		_582_v20 = _dafny.SetOf(_dafny.IntOfInt64(181), (_this).F23())
		var _583_v21 _dafny.Set
		_ = _583_v21
		_583_v21 = _dafny.SetOf(_582_v20)
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_579_v17), 0))
		_ = _index92
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_566_v4), 0))
		_ = _index93
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_581_v19), 0))
		_ = _index94
		var _rhs109 _dafny.Array = _566_v4
		_ = _rhs109
		var _rhs110 bool = (_this).Fm16((_this).F25(), (_this).F28(), _this.F24(), globalState)
		_ = _rhs110
		var _rhs111 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm1((_dafny.Zero).Minus((_this).F23()), _this.F24(), globalState), ((_this).F29()).Minus((_583_v21).Cardinality()))
		_ = _rhs111
		var _rhs112 _dafny.Int = (_dafny.Zero).Minus((_this).F29())
		_ = _rhs112
		var _lhs87 _dafny.Array = _579_v17
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_579_v17), 0))
		_ = _lhs88
		var _lhs89 _dafny.Array = _566_v4
		_ = _lhs89
		var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_566_v4), 0))
		_ = _lhs90
		var _lhs91 _dafny.Array = _581_v19
		_ = _lhs91
		var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_581_v19), 0))
		_ = _lhs92
		var _lhs93 *GlobalState = globalState
		_ = _lhs93
		(_lhs87).ArraySet1(_rhs109, (_lhs88).Int())
		(_lhs89).ArraySet1(_rhs110, (_lhs90).Int())
		(_lhs91).ArraySet1(_rhs111, (_lhs92).Int())
		_lhs93.F8 = _rhs112
		var _584_v22 _dafny.Map
		_ = _584_v22
		_584_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v1, _dafny.CodePoint('h'))
		var _pat_let_tv5 = _571_v9
		_ = _pat_let_tv5
		var _pat_let_tv6 = _581_v19
		_ = _pat_let_tv6
		var _pat_let_tv7 = _581_v19
		_ = _pat_let_tv7
		var _pat_let_tv8 = _571_v9
		_ = _pat_let_tv8
		var _pat_let_tv9 = _566_v4
		_ = _pat_let_tv9
		var _pat_let_tv10 = _566_v4
		_ = _pat_let_tv10
		r0 = func(_source5 D1) bool {
			if _source5.Is_DC4() {
				var _585___mcc_h0 bool = _source5.Get_().(D1_DC4).Cf5
				_ = _585___mcc_h0
				var _586___mcc_h1 bool = _source5.Get_().(D1_DC4).Cf6
				_ = _586___mcc_h1
				var _587_cf6 bool = _586___mcc_h1
				_ = _587_cf6
				var _588_cf5 bool = _585___mcc_h0
				_ = _588_cf5
				return _587_cf6
			} else if _source5.Is_DC5() {
				var _589___mcc_h2 bool = _source5.Get_().(D1_DC5).Cf7
				_ = _589___mcc_h2
				var _590___mcc_h3 _dafny.Map = _source5.Get_().(D1_DC5).Cf8
				_ = _590___mcc_h3
				var _591___mcc_h4 _dafny.Int = _source5.Get_().(D1_DC5).Cf9
				_ = _591___mcc_h4
				var _592___mcc_h5 _dafny.Int = _source5.Get_().(D1_DC5).Cf10
				_ = _592___mcc_h5
				var _593___mcc_h6 _dafny.Map = _source5.Get_().(D1_DC5).Cf11
				_ = _593___mcc_h6
				var _594_cf11 _dafny.Map = _593___mcc_h6
				_ = _594_cf11
				var _595_cf10 _dafny.Int = _592___mcc_h5
				_ = _595_cf10
				var _596_cf9 _dafny.Int = _591___mcc_h4
				_ = _596_cf9
				var _597_cf8 _dafny.Map = _590___mcc_h3
				_ = _597_cf8
				var _598_cf7 bool = _589___mcc_h2
				_ = _598_cf7
				return _598_cf7
			} else if _source5.Is_DC3() {
				var _599___mcc_h7 _dafny.CodePoint = _source5.Get_().(D1_DC3).Cf4
				_ = _599___mcc_h7
				var _600_cf4 _dafny.CodePoint = _599___mcc_h7
				_ = _600_cf4
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv5).Select((Companion_Default___.SafeIndex((_pat_let_tv7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_pat_let_tv6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_pat_let_tv8).Cardinality()))).Uint32()).(bool), (_this).F28())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_pat_let_tv10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_pat_let_tv9), 0))).Int()).(bool)))
			} else {
				var _601___mcc_h8 D1 = _source5.Get_().(D1_DC6).Cf12
				_ = _601___mcc_h8
				var _602_cf12 D1 = _601___mcc_h8
				_ = _602_cf12
				return (_this).F28()
			}
		}(Companion_D1_.Create_DC3_((func() _dafny.CodePoint {
			if (_584_v22).Contains(_563_v1) {
				return (_584_v22).Get(_563_v1).(_dafny.CodePoint)
			}
			return _dafny.CodePoint('k')
		})()))
		var _603_v23 T0
		_ = _603_v23
		var _nw108 *C2 = New_C2_()
		_ = _nw108
		_nw108.Ctor__((_dafny.Zero).Minus((_this).F29()))
		_603_v23 = _nw108
		r1 = _603_v23
		r2 = !((_this).F25()) || ((_this).F28())
		r3 = (_this).F23()
		return r0, r1, r2, r3
	}
}
func (_this *C3) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var _604_v0 _dafny.Set
		_ = _604_v0
		_604_v0 = _dafny.SetOf((_this).F29())
		var _605_v1 D8
		_ = _605_v1
		_605_v1 = Companion_D8_.Create_DC23_((_this).F23(), _604_v0, (_this).F23(), (_dafny.Zero).Minus(Companion_Default___.Fm1((_this).F29(), p1, globalState)), _dafny.IntOfInt64(850))
		var _606_v2 _dafny.Map
		_ = _606_v2
		_606_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v1, _dafny.UnicodeSeqOfUtf8Bytes("h"))
		var _607_v3 _dafny.Sequence
		_ = _607_v3
		_607_v3 = _dafny.UnicodeSeqOfUtf8Bytes("gnilurti")
		_606_v2 = (_606_v2).Update((func() D8 {
			if p0 {
				return _605_v1
			}
			return _605_v1
		})(), _607_v3)
		var _608_v4 _dafny.Sequence
		_ = _608_v4
		_608_v4 = _dafny.SeqOf(Companion_Default___.Fm31(p0, (_this).F28(), globalState))
		var _609_v5 _dafny.Sequence
		_ = _609_v5
		_609_v5 = _dafny.SeqOf(_dafny.IntOfInt64(692))
		var _610_i0 _dafny.Int
		_ = _610_i0
		_610_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update((_608_v4).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_608_v4).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F29()), _dafny.IntOfUint32(((_608_v4).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_608_v4).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_this).F29())), _609_v5) {
				{
					if (_610_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_610_i0 = (_610_i0).Plus(_dafny.One)
					var _611_v6 _dafny.Array
					_ = _611_v6
					var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
					_ = _nw109
					_611_v6 = _nw109
					var _612_v7 _dafny.Set
					_ = _612_v7
					_612_v7 = _dafny.SetOf(_611_v6, _611_v6, _611_v6, _611_v6)
					var _613_v8 *C0
					_ = _613_v8
					var _nw110 *C0 = New_C0_()
					_ = _nw110
					_nw110.Ctor__(_612_v7, (_this).F23(), Companion_Default___.SafeDivisionInt((_this).F23(), (_dafny.Zero).Minus((_this).F23())))
					_613_v8 = _nw110
					var _614_v9 _dafny.Array
					_ = _614_v9
					var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
					_ = _nw111
					_614_v9 = _nw111
					var _615_v10 _dafny.Set
					_ = _615_v10
					_615_v10 = _dafny.SetOf(_614_v9)
					var _616_v11 _dafny.Set
					_ = _616_v11
					_616_v11 = _dafny.SetOf((_this).F28())
					var _617_v12 _dafny.Map
					_ = _617_v12
					_617_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_this).F23(), (_this).F29(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}(func(_618_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('r')
					}))).Cardinality()), (_615_v10).Cardinality(), (_616_v11).Cardinality())).Cardinality(), (_this).F23())
					var _619_v13 *C1
					_ = _619_v13
					var _nw112 *C1 = New_C1_()
					_ = _nw112
					_nw112.Ctor__(false, (_617_v12).Equals(_617_v12))
					_619_v13 = _nw112
					var _620_v14 _dafny.Sequence
					_ = _620_v14
					_620_v14 = _dafny.SeqOf(_this.F24(), p1)
					var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_611_v6), 0))
					_ = _index95
					(_611_v6).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_620_v14, _620_v14), (_index95).Int())
					var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_611_v6), 0))
					_ = _index96
					(_611_v6).ArraySet1(true, (_index96).Int())
					var _621_v15 _dafny.Map
					_ = _621_v15
					_621_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(848))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg29 _dafny.Int) interface{} {
							return coer29(arg29)
						}
					}(func(_622_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					}))).Cardinality()))
					(globalState).F6 = !(!(_621_v15).Contains(_607_v3))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _623_v16 _dafny.MultiSet
		_ = _623_v16
		_623_v16 = _dafny.MultiSetOf((_this).F23())
		var _624_v17 _dafny.MultiSet
		_ = _624_v17
		_624_v17 = _dafny.MultiSetOf(!(p1))
		var _625_v18 _dafny.Map
		_ = _625_v18
		_625_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_624_v17).Cardinality(), _623_v16)
		var _626_v19 D4
		_ = _626_v19
		_626_v19 = Companion_D4_.Create_DC12_(((_623_v16).Difference((func() _dafny.MultiSet {
			if (_625_v18).Contains((_this).F29()) {
				return (_625_v18).Get((_this).F29()).(_dafny.MultiSet)
			}
			return _623_v16
		})())).Cardinality(), (func() _dafny.Int {
			if (_this).F28() {
				return (_this).F23()
			}
			return (_this).F23()
		})())
		_626_v19 = Companion_D4_.Create_DC12_((_this).F29(), (_dafny.Zero).Minus((_this).F23()))
		var _hi4 _dafny.Int = _dafny.IntOfInt64(-402)
		_ = _hi4
		for _627_i3 := (_this).F29(); _627_i3.Cmp(_hi4) < 0; _627_i3 = _627_i3.Plus(_dafny.One) {
			var _628_v20 _dafny.Set
			_ = _628_v20
			_628_v20 = _dafny.SetOf(p0)
			var _629_v21 _dafny.Sequence
			_ = _629_v21
			_629_v21 = _dafny.SeqOf(p1)
			var _630_v22 _dafny.Sequence
			_ = _630_v22
			_630_v22 = _dafny.SeqOf(_629_v21, _629_v21, _629_v21, _629_v21)
			var _631_v23 D1
			_ = _631_v23
			_631_v23 = Companion_D1_.Create_DC4_(p1, true)
			var _632_v24 _dafny.Sequence
			_ = _632_v24
			_632_v24 = _dafny.SeqOf(_631_v23, _631_v23, _631_v23, _631_v23, _631_v23)
			var _633_v26 _dafny.Array
			_ = _633_v26
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_17
			var _nw113 _dafny.Array
			_ = _nw113
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw113 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = (func(_634_p1 bool) func(_dafny.Int) bool {
					return func(_635_i4 _dafny.Int) bool {
						return _634_p1
					}
				})(p1)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw113 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw113).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw113).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_633_v26 = _nw113
			var _636_v27 _dafny.Map
			_ = _636_v27
			_636_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v26, (_this).F28())
			var _637_v28 _dafny.Map
			_ = _637_v28
			_637_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _633_v26)
			var _638_v29 _dafny.Set
			_ = _638_v29
			_638_v29 = _dafny.SetOf(_631_v23)
			var _639_v31 _dafny.Map
			_ = _639_v31
			_639_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-446), _628_v20)
			var _640_v32 _dafny.Sequence
			_ = _640_v32
			_640_v32 = _dafny.SeqOf(_628_v20, _628_v20, _628_v20, _628_v20, _628_v20)
			var _641_v33 _dafny.Map
			_ = _641_v33
			_641_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p1)
			var _642_v34 D1
			_ = _642_v34
			_642_v34 = Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_((_this).F28(), _641_v33, _627_i3, (_this).F23(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(_627_i3))))
			var _643_v35 _dafny.Array
			_ = _643_v35
			var _nwElement0_21 _dafny.Set = _628_v20
			_ = _nwElement0_21
			var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(22))
			_ = _nw114
			(_nw114).ArraySet1(_nwElement0_21, 0)
			(_nw114).ArraySet1(_628_v20, 1)
			(_nw114).ArraySet1(_dafny.SetOf(true), 2)
			(_nw114).ArraySet1(_628_v20, 3)
			(_nw114).ArraySet1(_628_v20, 4)
			(_nw114).ArraySet1(Companion_Default___.Fm32(_this.F24(), _609_v5, (_630_v22).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_630_v22).Cardinality()))).Uint32()).(_dafny.Sequence), func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate((_632_v24).Elements()); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _644_v25 D1
					_644_v25 = interface{}(_compr_22).(D1)
					if _dafny.Companion_Sequence_.Contains(_632_v24, _644_v25) {
						_coll22.Add(_644_v25)
					}
				}
				return _coll22.ToSet()
			}(), globalState), 5)
			(_nw114).ArraySet1(_628_v20, 6)
			(_nw114).ArraySet1(_628_v20, 7)
			(_nw114).ArraySet1(Companion_Default___.Fm32((func() bool {
				if (_636_v27).Contains((func() _dafny.Array {
					if (_637_v28).Contains(!(p0)) {
						return (_637_v28).Get(!(p0)).(_dafny.Array)
					}
					return _633_v26
				})()) {
					return (_636_v27).Get((func() _dafny.Array {
						if (_637_v28).Contains(!(p0)) {
							return (_637_v28).Get(!(p0)).(_dafny.Array)
						}
						return _633_v26
					})()).(bool)
				}
				return p1
			})(), _609_v5, _629_v21, _638_v29, globalState), 8)
			(_nw114).ArraySet1(Companion_Default___.Fm32((_this).F28(), _609_v5, _629_v21, func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate((Companion_Default___.Fm33((_this).F28(), _607_v3, _dafny.IntOfInt64(846), (_this).F28(), globalState)).Elements()); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _645_v30 D1
					_645_v30 = interface{}(_compr_23).(D1)
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm33((_this).F28(), _607_v3, _dafny.IntOfInt64(846), (_this).F28(), globalState), _645_v30) {
						_coll23.Add(_645_v30)
					}
				}
				return _coll23.ToSet()
			}(), globalState), 9)
			(_nw114).ArraySet1(_628_v20, 10)
			(_nw114).ArraySet1((func() _dafny.Set {
				if (_639_v31).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_629_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_607_v3).Cardinality()), _dafny.IntOfUint32((_629_v21).Cardinality()))).Uint32()).(bool), (_this).F28())).Cardinality()) {
					return (_639_v31).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_629_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_607_v3).Cardinality()), _dafny.IntOfUint32((_629_v21).Cardinality()))).Uint32()).(bool), (_this).F28())).Cardinality()).(_dafny.Set)
				}
				return _628_v20
			})(), 11)
			(_nw114).ArraySet1(((_640_v32).Select((Companion_Default___.SafeIndex((_641_v33).Cardinality(), _dafny.IntOfUint32((_640_v32).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_628_v20), 12)
			(_nw114).ArraySet1(_dafny.SetOf(_this.F24(), _this.F24(), (_629_v21).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F29()), _dafny.IntOfUint32((_629_v21).Cardinality()))).Uint32()).(bool)), 13)
			(_nw114).ArraySet1((func() _dafny.Set {
				if true {
					return _628_v20
				}
				return _dafny.SetOf(_this.F24())
			})(), 14)
			(_nw114).ArraySet1(_dafny.SetOf(!(p1)), 15)
			(_nw114).ArraySet1(_628_v20, 16)
			(_nw114).ArraySet1(_628_v20, 17)
			(_nw114).ArraySet1(_628_v20, 18)
			(_nw114).ArraySet1(_dafny.SetOf(!((_this).Fm17(_642_v34, globalState)), (_this).F25(), !(p0)), 19)
			(_nw114).ArraySet1(_628_v20, 20)
			(_nw114).ArraySet1(Companion_Default___.Fm32(_this.F24(), _609_v5, _629_v21, _638_v29, globalState), 21)
			_643_v35 = _nw114
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_643_v35), 0))
			_ = _index97
			(_643_v35).ArraySet1(_628_v20, (_index97).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_633_v26), 0))
			_ = _index98
			(_633_v26).ArraySet1(_this.F24(), (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_643_v35), 0))
			_ = _index99
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_633_v26), 0))
			_ = _index100
			var _rhs113 _dafny.Set = (_628_v20).Union(_628_v20)
			_ = _rhs113
			var _rhs114 bool = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_607_v3).Cardinality()), (_this).F23())).Cmp(_dafny.IntOfInt64(386)) <= 0
			_ = _rhs114
			var _rhs115 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
			_ = _rhs115
			var _rhs116 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F23(), (_this).F23())))
			_ = _rhs116
			var _lhs94 _dafny.Array = _643_v35
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_643_v35), 0))
			_ = _lhs95
			var _lhs96 _dafny.Array = _633_v26
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_633_v26), 0))
			_ = _lhs97
			var _lhs98 *GlobalState = globalState
			_ = _lhs98
			var _lhs99 *GlobalState = globalState
			_ = _lhs99
			(_lhs94).ArraySet1(_rhs113, (_lhs95).Int())
			(_lhs96).ArraySet1(_rhs114, (_lhs97).Int())
			_lhs98.F17 = _rhs115
			_lhs99.F14 = _rhs116
			(globalState).F5 = (func() _dafny.Int {
				if (_624_v17).Contains(((_633_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_633_v26), 0))).Int()).(bool)) == (_this.F24())) {
					return (_624_v17).Multiplicity(((_633_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_633_v26), 0))).Int()).(bool)) == (_this.F24()))
				}
				return ((_this).F29()).Minus(_dafny.IntOfUint32((_607_v3).Cardinality()))
			})()
			var _646_v36 _dafny.MultiSet
			_ = _646_v36
			_646_v36 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(553))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_647_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})))
			var _648_v37 *C1
			_ = _648_v37
			var _nw115 *C1 = New_C1_()
			_ = _nw115
			_nw115.Ctor__((_this).F28(), (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("riafjk"))).IsDisjointFrom(_646_v36))
			_648_v37 = _nw115
			(globalState).F14 = (_this).F29()
		}
		if p0 {
			if (_this).Fm16(true, (_this).F28(), _this.F24(), globalState) {
				var _649_v38 *C2
				_ = _649_v38
				var _nw116 *C2 = New_C2_()
				_ = _nw116
				_nw116.Ctor__((Companion_D10_.Create_DC29_(_this.F24(), _this.F24(), (_this).F23())).Dtor_cf55())
				_649_v38 = _nw116
				(globalState).F15 = ((_this).F25()) && (_this.F24())
				var _650_v39 _dafny.CodePoint
				_ = _650_v39
				_650_v39 = _dafny.CodePoint('g')
				_607_v3 = _dafny.Companion_Sequence_.Concatenate(_607_v3, _dafny.Companion_Sequence_.Update((Companion_D3_.Create_DC8_(_dafny.Companion_Sequence_.Update(_607_v3, (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_607_v3).Cardinality()))).Uint32(), _650_v39))).Dtor_cf14(), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32(((Companion_D3_.Create_DC8_(_dafny.Companion_Sequence_.Update(_607_v3, (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_607_v3).Cardinality()))).Uint32(), _650_v39))).Dtor_cf14()).Cardinality()))).Uint32(), _650_v39))
				var _651_v40 *C1
				_ = _651_v40
				var _nw117 *C1 = New_C1_()
				_ = _nw117
				_nw117.Ctor__((_this).F25(), p1)
				_651_v40 = _nw117
				var _652_v41 _dafny.Array
				_ = _652_v41
				var _nwElement0_22 *C1 = _651_v40
				_ = _nwElement0_22
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_22, 0)
				(_nw118).ArraySet1(_651_v40, 1)
				(_nw118).ArraySet1(_651_v40, 2)
				(_nw118).ArraySet1(_651_v40, 3)
				(_nw118).ArraySet1(_651_v40, 4)
				_652_v41 = _nw118
				var _653_v42 _dafny.MultiSet
				_ = _653_v42
				_653_v42 = _dafny.MultiSetOf(_652_v41)
				var _654_v43 _dafny.Array
				_ = _654_v43
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw119
				_654_v43 = _nw119
				var _655_v44 _dafny.Map
				_ = _655_v44
				_655_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v43, (func() _dafny.Int {
					if (_624_v17).Contains((_this).F28()) {
						return (_624_v17).Multiplicity((_this).F28())
					}
					return (_this).F29()
				})())
				var _rhs117 bool = ((_653_v42).Update(_652_v41, Companion_Default___.Abs((_this).F29()))).IsDisjointFrom(_653_v42)
				_ = _rhs117
				var _rhs118 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-779), ((_655_v44).Merge(_655_v44)).Cardinality())
				_ = _rhs118
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				_lhs100.F1 = _rhs117
				_lhs101.F17 = _rhs118
				var _656_v45 _dafny.Map
				_ = _656_v45
				_656_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), true)
				var _657_v46 _dafny.Map
				_ = _657_v46
				_657_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_624_v17, p0)
				var _658_v47 _dafny.Map
				_ = _658_v47
				_658_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(65))
				var _659_v48 D1
				_ = _659_v48
				_659_v48 = Companion_D1_.Create_DC5_(p1, _656_v45, _dafny.IntOfInt64(-776), (_657_v46).Cardinality(), _658_v47)
				var _660_v49 _dafny.Sequence
				_ = _660_v49
				_660_v49 = _dafny.SeqOf((_this).F28(), _this.F24())
				var _661_v50 _dafny.Array
				_ = _661_v50
				var _nwElement0_23 _dafny.Int = (_this).F23()
				_ = _nwElement0_23
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(8))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_23, 0)
				(_nw120).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_662_i6 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(858), true)).Cardinality()
				}))).Cardinality()), 1)
				(_nw120).ArraySet1((_this).F23(), 2)
				(_nw120).ArraySet1((_659_v48).Dtor_cf10(), 3)
				(_nw120).ArraySet1((_this).F23(), 4)
				(_nw120).ArraySet1((_this).F29(), 5)
				(_nw120).ArraySet1(_dafny.IntOfUint32((_660_v49).Cardinality()), 6)
				(_nw120).ArraySet1((_this).F23(), 7)
				_661_v50 = _nw120
				var _663_v51 _dafny.Array
				_ = _663_v51
				var _nwElement0_24 _dafny.Array = _661_v50
				_ = _nwElement0_24
				var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(14))
				_ = _nw121
				(_nw121).ArraySet1(_nwElement0_24, 0)
				(_nw121).ArraySet1(_661_v50, 1)
				(_nw121).ArraySet1(_661_v50, 2)
				(_nw121).ArraySet1(_661_v50, 3)
				(_nw121).ArraySet1(_661_v50, 4)
				(_nw121).ArraySet1(_661_v50, 5)
				(_nw121).ArraySet1(_661_v50, 6)
				(_nw121).ArraySet1(_661_v50, 7)
				(_nw121).ArraySet1((func() _dafny.Array {
					if p1 {
						return _661_v50
					}
					return _661_v50
				})(), 8)
				(_nw121).ArraySet1(_661_v50, 9)
				(_nw121).ArraySet1(_661_v50, 10)
				(_nw121).ArraySet1(_661_v50, 11)
				(_nw121).ArraySet1(_661_v50, 12)
				(_nw121).ArraySet1(_661_v50, 13)
				_663_v51 = _nw121
				var _rhs119 _dafny.Int = (_this).F29()
				_ = _rhs119
				var _rhs120 _dafny.Int = (_this).F23()
				_ = _rhs120
				var _rhs121 bool = p0
				_ = _rhs121
				var _rhs122 _dafny.CodePoint = _dafny.CodePoint('t')
				_ = _rhs122
				var _rhs123 _dafny.Array = _663_v51
				_ = _rhs123
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				var _lhs104 *GlobalState = globalState
				_ = _lhs104
				_lhs102.F8 = _rhs119
				_lhs103.F8 = _rhs120
				_lhs104.F9 = _rhs121
				_650_v39 = _rhs122
				_663_v51 = _rhs123
			} else {
				var _664_v52 _dafny.Map
				_ = _664_v52
				_664_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), true)
				var _665_v53 _dafny.Map
				_ = _665_v53
				_665_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-899))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_666_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality())).Plus((_664_v52).Cardinality()), (_this).F25())
				_665_v53 = (_665_v53).Update((_this).F29(), p1)
				(globalState).F5 = Companion_Default___.SafeDivisionInt((_this).F29(), (_this).F29())
				var _667_v54 _dafny.CodePoint
				_ = _667_v54
				_667_v54 = _dafny.CodePoint('m')
				_667_v54 = _667_v54
				var _668_v55 _dafny.Map
				_ = _668_v55
				_668_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F23())
				(globalState).F8 = (((_this).F29()).Times((_668_v55).Cardinality())).Times(((_this).F29()).Minus((_this).F29()))
				(globalState).F5 = (_this).Fm2(globalState)
			}
			var _669_v56 _dafny.Map
			_ = _669_v56
			_669_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_604_v0).Cardinality(), p0)
			_669_v56 = _669_v56
			var _670_v57 _dafny.Array
			_ = _670_v57
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw122
			_670_v57 = _nw122
			r1 = _670_v57
			var _671_v58 _dafny.CodePoint
			_ = _671_v58
			_671_v58 = _dafny.CodePoint('n')
			_671_v58 = _671_v58
			(globalState).F6 = false
		} else {
			(globalState).F1 = (_624_v17).IsDisjointFrom(_624_v17)
			var _672_v59 _dafny.Array
			_ = _672_v59
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_18
			var _nw123 _dafny.Array
			_ = _nw123
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw123 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = func(_673_i8 _dafny.Int) bool {
					return (_this).F25()
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw123 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw123).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw123).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_672_v59 = _nw123
			_672_v59 = _672_v59
			var _rhs124 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_607_v3, _607_v3)
			_ = _rhs124
			var _rhs125 _dafny.Int = (_this).F29()
			_ = _rhs125
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			_607_v3 = _rhs124
			_lhs105.F5 = _rhs125
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_672_v59), 0))
			_ = _index101
			(_672_v59).ArraySet1((_this).F28(), (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_672_v59), 0))
			_ = _index102
			(_672_v59).ArraySet1(_this.F24(), (_index102).Int())
			var _674_v60 T1
			_ = _674_v60
			var _nw124 *C1 = New_C1_()
			_ = _nw124
			_nw124.Ctor__(p0, (_this).F28())
			_674_v60 = _nw124
			var _675_v61 _dafny.Set
			_ = _675_v61
			_675_v61 = _dafny.SetOf(_674_v60)
			(_this).F24_set_((p0) == (!((_675_v61).Equals(_675_v61))))
		}
		var _676_v62 _dafny.Array
		_ = _676_v62
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_19
		var _nw125 _dafny.Array
		_ = _nw125
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw125 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Int = func(_677_i9 _dafny.Int) _dafny.Int {
				return (_677_i9).Plus((_this).F29())
			}
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw125 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw125).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw125).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_676_v62 = _nw125
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_676_v62), 0))
		_ = _index103
		(_676_v62).ArraySet1((_this).F23(), (_index103).Int())
		var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_676_v62), 0))
		_ = _index104
		(_676_v62).ArraySet1((_this).F23(), (_index104).Int())
		r0 = (_this).F29()
		r1 = _676_v62
		r2 = !(_this.F24())
		return r0, r1, r2
	}
}
func (_this *C3) M5(p0 bool, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _678_v0 _dafny.Array
		_ = _678_v0
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_20
		var _nw126 _dafny.Array
		_ = _nw126
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw126 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) bool = func(_679_i1 _dafny.Int) bool {
				return _this.F24()
			}
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw126 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw126).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw126).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_678_v0 = _nw126
		for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_678_v0), 0))); ; {
			_guard_loop_0, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _680_i0 _dafny.Int
			_680_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_680_i0).Sign() != -1) && ((_680_i0).Cmp(_dafny.ArrayLen((_678_v0), 0)) < 0)) {
				(_678_v0).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_this).F23())), (_this).F23()), (_680_i0).Int())
			}
		}
		var _681_v1 _dafny.Array
		_ = _681_v1
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_21
		var _nw127 _dafny.Array
		_ = _nw127
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw127 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) _dafny.Int = func(_682_i2 _dafny.Int) _dafny.Int {
				return (_682_i2).Times((_this).F29())
			}
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw127 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw127).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw127).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_681_v1 = _nw127
		var _683_v2 _dafny.Map
		_ = _683_v2
		_683_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_681_v1, _678_v0)
		_683_v2 = (_683_v2).Update(_681_v1, _678_v0)
		var _684_i3 _dafny.Int
		_ = _684_i3
		_684_i3 = _dafny.Zero
		{
			for ((_this).F28()) && ((_this).F25()) {
				{
					if (_684_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_684_i3 = (_684_i3).Plus(_dafny.One)
					var _685_v3 _dafny.MultiSet
					_ = _685_v3
					_685_v3 = _dafny.MultiSetOf(!(p0))
					var _686_v4 D10
					_ = _686_v4
					_686_v4 = Companion_D10_.Create_DC29_(_this.F24(), (_685_v3).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F25(), (_this).F25(), (_this).F25()))), p1)
					var _687_v5 _dafny.Sequence
					_ = _687_v5
					_687_v5 = _dafny.SeqOf(_this.F24(), p0)
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_678_v0), 0))
					_ = _index105
					(_678_v0).ArraySet1(!_dafny.Companion_Sequence_.Equal(_687_v5, _dafny.SeqOf(_this.F24(), (_this).F28(), _this.F24(), (_this).F25(), _this.F24())), (_index105).Int())
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_678_v0), 0))
					_ = _index106
					var _rhs126 D10 = _686_v4
					_ = _rhs126
					var _rhs127 bool = _this.F24()
					_ = _rhs127
					var _rhs128 _dafny.Int = (_this).F29()
					_ = _rhs128
					var _lhs106 _dafny.Array = _678_v0
					_ = _lhs106
					var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_678_v0), 0))
					_ = _lhs107
					var _lhs108 *GlobalState = globalState
					_ = _lhs108
					_686_v4 = _rhs126
					(_lhs106).ArraySet1(_rhs127, (_lhs107).Int())
					_lhs108.F14 = _rhs128
					var _688_v6 _dafny.Sequence
					_ = _688_v6
					_688_v6 = _dafny.SeqOf((_this).F29())
					var _689_v7 D10
					_ = _689_v7
					_689_v7 = Companion_D10_.Create_DC28_(_688_v6)
					var _690_v8 _dafny.Map
					_ = _690_v8
					_690_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F23())
					var _691_v9 _dafny.Map
					_ = _691_v9
					_691_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_689_v7).Dtor_cf52(), (_690_v8).Cardinality())
					var _692_v10 _dafny.Map
					_ = _692_v10
					_692_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v9, (_this).F29())
					var _693_v11 _dafny.Sequence
					_ = _693_v11
					_693_v11 = _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F29(), (_dafny.Zero).Minus((_this).F29()))), _dafny.IntOfInt64(68), _dafny.IntOfUint32((_688_v6).Cardinality()), _dafny.IntOfInt64(800), (func() _dafny.Int {
						if (_692_v10).Contains(_691_v9) {
							return (_692_v10).Get(_691_v9).(_dafny.Int)
						}
						return (_this).F29()
					})())
					_693_v11 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}(func(_694_i4 _dafny.Int) _dafny.Int {
						return (_this).F23()
					}))
					var _695_v12 *C1
					_ = _695_v12
					var _nw128 *C1 = New_C1_()
					_ = _nw128
					_nw128.Ctor__((_this.F24()) || (_this.F24()), p0)
					_695_v12 = _nw128
					var _696_v13 _dafny.Sequence
					_ = _696_v13
					_696_v13 = _dafny.SeqOf(_681_v1, _681_v1)
					var _697_v14 _dafny.Map
					_ = _697_v14
					_697_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _681_v1)
					var _698_v15 _dafny.Sequence
					_ = _698_v15
					_698_v15 = _dafny.SeqOf((func() _dafny.Array {
						if (_678_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_678_v0), 0))).Int()).(bool) {
							return (_696_v13).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_696_v13).Cardinality()))).Uint32()).(_dafny.Array)
						}
						return _681_v1
					})(), (func() _dafny.Array {
						if (_this).F25() {
							return _681_v1
						}
						return (func() _dafny.Array {
							if (_697_v14).Contains(_this.F24()) {
								return (_697_v14).Get(_this.F24()).(_dafny.Array)
							}
							return _681_v1
						})()
					})(), _681_v1, _681_v1)
					var _699_v16 D0
					_ = _699_v16
					_699_v16 = Companion_D0_.Create_DC0_(p0)
					var _700_v17 _dafny.CodePoint
					_ = _700_v17
					_700_v17 = _dafny.CodePoint('e')
					var _701_v18 _dafny.Set
					_ = _701_v18
					_701_v18 = _dafny.SetOf(_this.F24(), (_678_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_678_v0), 0))).Int()).(bool), (_this).F28())
					var _rhs129 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25((_699_v16).Dtor_cf0(), (_this).F28(), _700_v17, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg34 _dafny.Int) interface{} {
							return coer34(arg34)
						}
					}((func(_702_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_703_i5 _dafny.Int) _dafny.CodePoint {
							return _702_v17
						}
					})(_700_v17))))).Cardinality()))
					_ = _rhs129
					var _rhs130 bool = (_701_v18).IsProperSubsetOf(_701_v18)
					_ = _rhs130
					var _rhs131 _dafny.Int = (_this).F29()
					_ = _rhs131
					var _rhs132 _dafny.Sequence = _696_v13
					_ = _rhs132
					var _lhs109 *GlobalState = globalState
					_ = _lhs109
					var _lhs110 *GlobalState = globalState
					_ = _lhs110
					var _lhs111 *GlobalState = globalState
					_ = _lhs111
					_lhs109.F17 = _rhs129
					_lhs110.F6 = _rhs130
					_lhs111.F8 = _rhs131
					_698_v15 = _rhs132
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _704_v19 _dafny.CodePoint
		_ = _704_v19
		_704_v19 = _dafny.CodePoint('v')
		var _705_v20 _dafny.Map
		_ = _705_v20
		_705_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _704_v19)
		var _706_v21 _dafny.Sequence
		_ = _706_v21
		_706_v21 = _dafny.UnicodeSeqOfUtf8Bytes("vvejejg")
		(globalState).F6 = (((_705_v20).Merge(_705_v20)).Cardinality()).Cmp(_dafny.IntOfUint32((_706_v21).Cardinality())) != 0
		var _707_v22 _dafny.Map
		_ = _707_v22
		_707_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.IntOfInt64(643)).Minus((_this).F23()))
		_707_v22 = (_707_v22).Update(p1, (_this).F29())
		var _708_v23 _dafny.Map
		_ = _708_v23
		_708_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
		var _709_v25 _dafny.Sequence
		_ = _709_v25
		_709_v25 = _dafny.SeqOf((_this).F23())
		var _710_v26 _dafny.Array
		_ = _710_v26
		var _nwElement0_25 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _this.F24())).Merge(_708_v23)
		_ = _nwElement0_25
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(8))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_25, 0)
		(_nw129).ArraySet1((_708_v23).Merge(Companion_Default___.Fm34((_dafny.Zero).Minus(_dafny.IntOfUint32((_706_v21).Cardinality())), globalState)), 1)
		(_nw129).ArraySet1(_708_v23, 2)
		(_nw129).ArraySet1(_708_v23, 3)
		(_nw129).ArraySet1(_708_v23, 4)
		(_nw129).ArraySet1(func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter25 := _dafny.Iterate((_709_v25).Elements()); ; {
				_compr_24, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _711_v24 _dafny.Int
				_711_v24 = interface{}(_compr_24).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_709_v25, _711_v24) {
					_coll24.Add((_711_v24).Plus((_this).F29()), (_this).F25())
				}
			}
			return _coll24.ToMap()
		}(), 5)
		(_nw129).ArraySet1(_708_v23, 6)
		(_nw129).ArraySet1(_708_v23, 7)
		_710_v26 = _nw129
		for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_710_v26), 0))); ; {
			_guard_loop_1, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _712_i6 _dafny.Int
			_712_i6 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_712_i6).Sign() != -1) && ((_712_i6).Cmp(_dafny.ArrayLen((_710_v26), 0)) < 0)) {
				(_710_v26).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-59)), _this.F24())).Merge(_708_v23), (_712_i6).Int())
			}
		}
	}
}
func (_this *C3) F28() bool {
	{
		return _this._f28
	}
}
func (_this *C3) F29() _dafny.Int {
	{
		return _this._f29
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f24 bool
	_f25 bool
	_f27 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f24 = false
	_this._f25 = false
	_this._f27 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F24() bool {
	return _this._f24
}
func (_this *C4) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C4) F25() bool {
	return _this._f25
}
func (_this *C4) Ctor__(f27 bool, f24 bool, f25 bool) {
	{
		(_this)._f27 = f27
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C4) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _this.F24()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), (Companion_D0_.Create_DC1_(!((_this).F25()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aeynwfee")).Cardinality()))).Dtor_cf1())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F27()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F24()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _this.F24()))))
	}
}
func (_this *C4) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-771)), _dafny.MultiSetOf(_dafny.IntOfInt64(233)))).Merge(func() _dafny.Map {
			var _coll25 = _dafny.NewMapBuilder()
			_ = _coll25
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(923), _dafny.IntOfInt64(571))); ; {
				_compr_25, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _713_v0 _dafny.Int
				_713_v0 = interface{}(_compr_25).(_dafny.Int)
				if ((_dafny.IntOfInt64(923)).Cmp(_713_v0) <= 0) && ((_713_v0).Cmp(_dafny.IntOfInt64(571)) < 0) {
					_coll25.Add((_713_v0).Minus(_dafny.IntOfInt64(248)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mpuhcuieu")).Cardinality()), (func() _dafny.Set {
						var _coll26 = _dafny.NewBuilder()
						_ = _coll26
						for _iter28 := _dafny.Iterate((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(362))).Cardinality())).Elements()); ; {
							_compr_26, _ok28 := _iter28()
							if !_ok28 {
								break
							}
							var _714_v1 _dafny.Int
							_714_v1 = interface{}(_compr_26).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(362))).Cardinality()), _714_v1) {
								_coll26.Add(Companion_Default___.SafeDivisionInt(_714_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality())))
							}
						}
						return _coll26.ToSet()
					}()).Cardinality(), _dafny.IntOfInt64(220)))
				}
			}
			return _coll25.ToMap()
		}())).Cardinality()
	}
}
func (_this *C4) Fm14(globalState *GlobalState) bool {
	{
		return (_this).F27()
	}
}
func (_this *C4) Fm15(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var _source6 D1 = Companion_D1_.Create_DC4_((_this).F27(), _this.F24())
		_ = _source6
		if _source6.Is_DC4() {
			var _715___mcc_h0 bool = _source6.Get_().(D1_DC4).Cf5
			_ = _715___mcc_h0
			var _716___mcc_h1 bool = _source6.Get_().(D1_DC4).Cf6
			_ = _716___mcc_h1
			var _717_cf6 bool = _716___mcc_h1
			_ = _717_cf6
			var _718_cf5 bool = _715___mcc_h0
			_ = _718_cf5
			return _dafny.IntOfInt64(981)
		} else if _source6.Is_DC5() {
			var _719___mcc_h2 bool = _source6.Get_().(D1_DC5).Cf7
			_ = _719___mcc_h2
			var _720___mcc_h3 _dafny.Map = _source6.Get_().(D1_DC5).Cf8
			_ = _720___mcc_h3
			var _721___mcc_h4 _dafny.Int = _source6.Get_().(D1_DC5).Cf9
			_ = _721___mcc_h4
			var _722___mcc_h5 _dafny.Int = _source6.Get_().(D1_DC5).Cf10
			_ = _722___mcc_h5
			var _723___mcc_h6 _dafny.Map = _source6.Get_().(D1_DC5).Cf11
			_ = _723___mcc_h6
			var _724_cf11 _dafny.Map = _723___mcc_h6
			_ = _724_cf11
			var _725_cf10 _dafny.Int = _722___mcc_h5
			_ = _725_cf10
			var _726_cf9 _dafny.Int = _721___mcc_h4
			_ = _726_cf9
			var _727_cf8 _dafny.Map = _720___mcc_h3
			_ = _727_cf8
			var _728_cf7 bool = _719___mcc_h2
			_ = _728_cf7
			return _725_cf10
		} else if _source6.Is_DC3() {
			var _729___mcc_h7 _dafny.CodePoint = _source6.Get_().(D1_DC3).Cf4
			_ = _729___mcc_h7
			var _730_cf4 _dafny.CodePoint = _729___mcc_h7
			_ = _730_cf4
			return _dafny.IntOfInt64(810)
		} else {
			var _731___mcc_h8 D1 = _source6.Get_().(D1_DC6).Cf12
			_ = _731___mcc_h8
			var _732_cf12 D1 = _731___mcc_h8
			_ = _732_cf12
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("caeqeee"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_733_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})))).Cardinality())
		}
	}
}
func (_this *C4) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _734_v0 _dafny.Int
		_ = _734_v0
		_734_v0 = _dafny.IntOfInt64(428)
		var _hi5 _dafny.Int = _734_v0
		_ = _hi5
		for _735_i0 := (_734_v0).Times(_734_v0); _735_i0.Cmp(_hi5) < 0; _735_i0 = _735_i0.Plus(_dafny.One) {
			(globalState).F15 = !((_this).F27()) || (_this.F24())
			var _736_v1 _dafny.Array
			_ = _736_v1
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw130
			_736_v1 = _nw130
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_736_v1), 0))
			_ = _index107
			(_736_v1).ArraySet1((_this).F27(), (_index107).Int())
			var _737_v2 _dafny.Map
			_ = _737_v2
			_737_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F24()), (_this).F27())
			var _738_v3 _dafny.Map
			_ = _738_v3
			_738_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_737_v2).Cardinality())
			var _739_v4 D1
			_ = _739_v4
			_739_v4 = Companion_D1_.Create_DC5_((_this).F27(), _737_v2, _735_i0, _dafny.IntOfInt64(996), _738_v3)
			var _740_v5 D1
			_ = _740_v5
			_740_v5 = Companion_D1_.Create_DC6_(_739_v4)
			var _pat_let_tv11 = _739_v4
			_ = _pat_let_tv11
			var _741_v6 _dafny.MultiSet
			_ = _741_v6
			_741_v6 = _dafny.MultiSetOf(func(_pat_let5_0 D1) D1 {
				return func(_742_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let6_0 D1) D1 {
						return func(_743_dt__update_hcf12_h0 D1) D1 {
							return Companion_D1_.Create_DC6_(_743_dt__update_hcf12_h0)
						}(_pat_let6_0)
					}(_pat_let_tv11)
				}(_pat_let5_0)
			}(_740_v5), _740_v5, _740_v5)
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_736_v1), 0))
			_ = _index108
			(_736_v1).ArraySet1(((_741_v6).IsProperSubsetOf(_741_v6)) || (false), (_index108).Int())
			var _744_v7 D1
			_ = _744_v7
			_744_v7 = Companion_D1_.Create_DC5_(false, (_737_v2).Merge(_737_v2), _734_v0, _735_i0, _738_v3)
			_744_v7 = _744_v7
			var _745_v8 _dafny.Array
			_ = _745_v8
			var _nwElement0_26 _dafny.Array = _736_v1
			_ = _nwElement0_26
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.One)
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_26, 0)
			_745_v8 = _nw131
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_745_v8), 0))
			_ = _index109
			(_745_v8).ArraySet1(_736_v1, (_index109).Int())
			var _746_v9 _dafny.Array
			_ = _746_v9
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
			_ = _nw132
			_746_v9 = _nw132
			var _747_v10 _dafny.Array
			_ = _747_v10
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_22
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_748_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_749_i1 _dafny.Int) _dafny.Int {
						return (_749_i1).Minus(_748_i0)
					}
				})(_735_i0)
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
			_747_v10 = _nw133
			var _750_v11 _dafny.Map
			_ = _750_v11
			_750_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v0, _747_v10)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_746_v9), 0))
			_ = _index110
			(_746_v9).ArraySet1(_750_v11, (_index110).Int())
			var _751_v12 _dafny.Array
			_ = _751_v12
			_751_v12 = _736_v1
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_745_v8), 0))
			_ = _index111
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_746_v9), 0))
			_ = _index112
			var _rhs133 _dafny.Array = (_736_v1)
			_ = _rhs133
			var _rhs134 _dafny.Int = _735_i0
			_ = _rhs134
			var _rhs135 _dafny.Map = _750_v11
			_ = _rhs135
			var _rhs136 _dafny.Int = _734_v0
			_ = _rhs136
			var _lhs112 _dafny.Array = _745_v8
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_745_v8), 0))
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			var _lhs115 _dafny.Array = _746_v9
			_ = _lhs115
			var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_746_v9), 0))
			_ = _lhs116
			var _lhs117 *GlobalState = globalState
			_ = _lhs117
			(_lhs112).ArraySet1(_rhs133, (_lhs113).Int())
			_lhs114.F14 = _rhs134
			(_lhs115).ArraySet1(_rhs135, (_lhs116).Int())
			_lhs117.F14 = _rhs136
		}
		var _752_v13 _dafny.Array
		_ = _752_v13
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_23
		var _nw134 _dafny.Array
		_ = _nw134
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw134 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) bool = func(_753_i2 _dafny.Int) bool {
				return _this.F24()
			}
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw134 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw134).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw134).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_752_v13 = _nw134
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))
		_ = _index113
		(_752_v13).ArraySet1((_this).F25(), (_index113).Int())
		var _754_v14 _dafny.Map
		_ = _754_v14
		_754_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v0, _734_v0)
		var _755_v16 _dafny.MultiSet
		_ = _755_v16
		_755_v16 = _dafny.MultiSetOf(_754_v14, func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-595), _dafny.IntOfInt64(441))); ; {
				_compr_27, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _756_v15 _dafny.Int
				_756_v15 = interface{}(_compr_27).(_dafny.Int)
				if ((_dafny.IntOfInt64(-595)).Cmp(_756_v15) <= 0) && ((_756_v15).Cmp(_dafny.IntOfInt64(441)) < 0) {
					_coll27.Add(Companion_Default___.SafeModuloInt(_756_v15, _734_v0), _734_v0)
				}
			}
			return _coll27.ToMap()
		}(), _754_v14)
		var _757_v17 _dafny.Map
		_ = _757_v17
		_757_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_755_v16, _734_v0)
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))
		_ = _index114
		var _rhs137 bool = !((_734_v0).Cmp(_734_v0) < 0) || (_this.F24())
		_ = _rhs137
		var _rhs138 bool = _this.F24()
		_ = _rhs138
		var _rhs139 bool = _this.F24()
		_ = _rhs139
		var _rhs140 _dafny.Int = (func() _dafny.Int {
			if (_757_v17).Contains(_755_v16) {
				return (_757_v17).Get(_755_v16).(_dafny.Int)
			}
			return (_734_v0).Minus(_734_v0)
		})()
		_ = _rhs140
		var _lhs118 *GlobalState = globalState
		_ = _lhs118
		var _lhs119 *C4 = _this
		_ = _lhs119
		var _lhs120 _dafny.Array = _752_v13
		_ = _lhs120
		var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))
		_ = _lhs121
		var _lhs122 *GlobalState = globalState
		_ = _lhs122
		_lhs118.F9 = _rhs137
		_lhs119.F24_set_(_rhs138)
		(_lhs120).ArraySet1(_rhs139, (_lhs121).Int())
		_lhs122.F5 = _rhs140
		var _758_v18 _dafny.Sequence
		_ = _758_v18
		_758_v18 = _dafny.SeqOf((_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool), false)
		var _759_v19 _dafny.Map
		_ = _759_v19
		_759_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _this.F24())
		var _760_v20 _dafny.MultiSet
		_ = _760_v20
		_760_v20 = _dafny.MultiSetOf((func() bool {
			if (_759_v19).Contains(_this.F24()) {
				return (_759_v19).Get(_this.F24()).(bool)
			}
			return (_this).F25()
		})())
		var _761_v21 _dafny.Sequence
		_ = _761_v21
		_761_v21 = _dafny.UnicodeSeqOfUtf8Bytes("lwbojnlo")
		var _762_v22 _dafny.MultiSet
		_ = _762_v22
		_762_v22 = _dafny.MultiSetOf(_734_v0, _734_v0)
		var _763_v23 _dafny.Set
		_ = _763_v23
		_763_v23 = _dafny.SetOf(_734_v0)
		var _764_v24 _dafny.Array
		_ = _764_v24
		var _nwElement0_27 _dafny.Int = _734_v0
		_ = _nwElement0_27
		var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(19))
		_ = _nw135
		(_nw135).ArraySet1(_nwElement0_27, 0)
		(_nw135).ArraySet1(_734_v0, 1)
		(_nw135).ArraySet1(_734_v0, 2)
		(_nw135).ArraySet1(_734_v0, 3)
		(_nw135).ArraySet1(_734_v0, 4)
		(_nw135).ArraySet1(_734_v0, 5)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_758_v18, (Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_758_v18).Cardinality()))).Uint32(), _this.F24())).Cardinality()), 6)
		(_nw135).ArraySet1(_734_v0, 7)
		(_nw135).ArraySet1(_734_v0, 8)
		(_nw135).ArraySet1((_760_v20).Cardinality(), 9)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality()), 10)
		(_nw135).ArraySet1(_734_v0, 11)
		(_nw135).ArraySet1(_734_v0, 12)
		(_nw135).ArraySet1(_dafny.IntOfUint32((_761_v21).Cardinality()), 13)
		(_nw135).ArraySet1((func() _dafny.Int {
			if (_762_v22).Contains(_734_v0) {
				return (_762_v22).Multiplicity(_734_v0)
			}
			return (_dafny.Zero).Minus((_dafny.Zero).Minus(_734_v0))
		})(), 14)
		(_nw135).ArraySet1(_734_v0, 15)
		(_nw135).ArraySet1((_763_v23).Cardinality(), 16)
		(_nw135).ArraySet1(_734_v0, 17)
		(_nw135).ArraySet1(_734_v0, 18)
		_764_v24 = _nw135
		var _765_v25 _dafny.Sequence
		_ = _765_v25
		_765_v25 = _dafny.SeqOf(_764_v24)
		var _766_v26 _dafny.Sequence
		_ = _766_v26
		_766_v26 = _dafny.SeqOf(_dafny.IntOfInt64(662), _734_v0, _dafny.IntOfUint32((_761_v21).Cardinality()))
		var _767_v27 _dafny.Sequence
		_ = _767_v27
		_767_v27 = _dafny.SeqOf((_766_v26).Select((Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_766_v26).Cardinality()))).Uint32()).(_dafny.Int))
		var _hi6 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_765_v25, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_766_v26).Cardinality()), _dafny.IntOfUint32((_765_v25).Cardinality()))).Uint32(), (_765_v25).Select((Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_765_v25).Cardinality()))).Uint32()).(_dafny.Array))).Cardinality())).Minus(_dafny.IntOfInt64(845))
		_ = _hi6
		for _768_i3 := _734_v0; _768_i3.Cmp(_hi6) < 0; _768_i3 = _768_i3.Plus(_dafny.One) {
			if (_this).F25() {
				_761_v21 = _761_v21
				_758_v18 = _758_v18
				var _769_v28 _dafny.Array
				_ = _769_v28
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_24
				var _nw136 _dafny.Array
				_ = _nw136
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw136 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Map = (func(_770_i3 _dafny.Int, _771_v21 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_772_i4 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _770_i3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfUint32(((Companion_D3_.Create_DC8_(_771_v21)).Dtor_cf14()).Cardinality())))
						}
					})(_768_i3, _761_v21)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw136 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw136).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw136).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_769_v28 = _nw136
				var _773_v29 _dafny.CodePoint
				_ = _773_v29
				_773_v29 = _dafny.CodePoint('n')
				var _774_v30 _dafny.Map
				_ = _774_v30
				_774_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v29, _768_i3)
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_769_v28), 0))
				_ = _index115
				(_769_v28).ArraySet1(_774_v30, (_index115).Int())
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_769_v28), 0))
				_ = _index116
				(_769_v28).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v29, _734_v0)).Update(_773_v29, _768_i3), (_index116).Int())
				_773_v29 = (func() _dafny.CodePoint {
					if (_768_i3).Cmp(_734_v0) < 0 {
						return (_761_v21).Select((Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_761_v21).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _773_v29
				})()
				var _775_v31 D0
				_ = _775_v31
				_775_v31 = Companion_D0_.Create_DC1_((_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool), _734_v0)
				var _776_v32 D0
				_ = _776_v32
				_776_v32 = Companion_D0_.Create_DC2_(_775_v31)
				var _777_v33 _dafny.Map
				_ = _777_v33
				_777_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _776_v32)
				_777_v33 = (_777_v33).Update(_this.F24(), _776_v32)
			} else {
				var _778_v34 _dafny.Map
				_ = _778_v34
				_778_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm14(globalState), _752_v13)
				_778_v34 = _778_v34
				_759_v19 = (_759_v19).Merge((_759_v19).Merge(_759_v19))
				(globalState).F6 = (Companion_D0_.Create_DC0_(_this.F24())).Dtor_cf0()
				(globalState).F9 = (_734_v0).Cmp((_768_i3).Minus(_dafny.IntOfInt64(468))) <= 0
				var _779_v35 _dafny.MultiSet
				_ = _779_v35
				_779_v35 = _dafny.MultiSetOf(_762_v22)
				_779_v35 = _779_v35
			}
			var _780_v36 _dafny.Array
			_ = _780_v36
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_25
			var _nw137 _dafny.Array
			_ = _nw137
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw137 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Sequence = (func(_781_v18 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_782_i5 _dafny.Int) _dafny.Sequence {
						return _781_v18
					}
				})(_758_v18)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw137 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw137).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw137).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_780_v36 = _nw137
			var _783_v37 D3
			_ = _783_v37
			_783_v37 = Companion_D3_.Create_DC9_(_dafny.IntOfUint32((_761_v21).Cardinality()))
			var _784_v38 _dafny.Map
			_ = _784_v38
			_784_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v37, _dafny.IntOfInt64(666))
			var _785_v39 _dafny.Map
			_ = _785_v39
			_785_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_784_v38, _780_v36)
			var _786_v40 _dafny.Array
			_ = _786_v40
			var _nwElement0_28 _dafny.Array = _780_v36
			_ = _nwElement0_28
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(19))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_28, 0)
			(_nw138).ArraySet1(_780_v36, 1)
			(_nw138).ArraySet1(_780_v36, 2)
			(_nw138).ArraySet1(_780_v36, 3)
			(_nw138).ArraySet1((func() _dafny.Array {
				if (_785_v39).Contains(_784_v38) {
					return (_785_v39).Get(_784_v38).(_dafny.Array)
				}
				return _780_v36
			})(), 4)
			(_nw138).ArraySet1(_780_v36, 5)
			(_nw138).ArraySet1(_780_v36, 6)
			(_nw138).ArraySet1(_780_v36, 7)
			(_nw138).ArraySet1(_780_v36, 8)
			(_nw138).ArraySet1(_780_v36, 9)
			(_nw138).ArraySet1(_780_v36, 10)
			(_nw138).ArraySet1(_780_v36, 11)
			(_nw138).ArraySet1(_780_v36, 12)
			(_nw138).ArraySet1(_780_v36, 13)
			(_nw138).ArraySet1(_780_v36, 14)
			(_nw138).ArraySet1(_780_v36, 15)
			(_nw138).ArraySet1(_780_v36, 16)
			(_nw138).ArraySet1(_780_v36, 17)
			(_nw138).ArraySet1(_780_v36, 18)
			_786_v40 = _nw138
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_786_v40), 0))
			_ = _index117
			(_786_v40).ArraySet1(_780_v36, (_index117).Int())
			var _787_v41 _dafny.Set
			_ = _787_v41
			_787_v41 = _dafny.SetOf(_766_v26)
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_786_v40), 0))
			_ = _index118
			var _rhs141 _dafny.Int = (_this).Fm15(((_787_v41).Cardinality()).Plus(_734_v0), _dafny.IntOfInt64(375), (_this).Fm14(globalState), globalState)
			_ = _rhs141
			var _rhs142 _dafny.Array = _780_v36
			_ = _rhs142
			var _lhs123 *GlobalState = globalState
			_ = _lhs123
			var _lhs124 _dafny.Array = _786_v40
			_ = _lhs124
			var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_786_v40), 0))
			_ = _lhs125
			_lhs123.F14 = _rhs141
			(_lhs124).ArraySet1(_rhs142, (_lhs125).Int())
			if (_this).F27() {
				r0 = (_this).F25()
				var _788_v42 _dafny.CodePoint
				_ = _788_v42
				_788_v42 = _dafny.CodePoint('o')
				var _789_v43 _dafny.Map
				_ = _789_v43
				_789_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _788_v42)
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_764_v24), 0))
				_ = _index119
				(_764_v24).ArraySet1((_789_v43).Cardinality(), (_index119).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_764_v24), 0))
				_ = _index120
				(_764_v24).ArraySet1(_768_i3, (_index120).Int())
				var _790_v44 _dafny.Map
				_ = _790_v44
				_790_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_761_v21).Cardinality()), _752_v13)
				_790_v44 = (_790_v44).Update(_dafny.IntOfInt64(94), _752_v13)
				var _791_v45 _dafny.Map
				_ = _791_v45
				_791_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_752_v13, _762_v22)
				var _792_v46 _dafny.Sequence
				_ = _792_v46
				_792_v46 = _dafny.SeqOf(_762_v22, _dafny.MultiSetFromSeq(_767_v27))
				_791_v45 = (_791_v45).Update(_752_v13, (_792_v46).Select((Companion_Default___.SafeIndex(_768_i3, _dafny.IntOfUint32((_792_v46).Cardinality()))).Uint32()).(_dafny.MultiSet))
				(globalState).F5 = _734_v0
			} else {
				(globalState).F8 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_768_i3, (_dafny.Zero).Minus((_768_i3).Minus(_768_i3))))
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))
				_ = _index121
				(_752_v13).ArraySet1((_this).Fm14(globalState), (_index121).Int())
				r0 = (_this).F27()
				var _793_v47 *C1
				_ = _793_v47
				var _nw139 *C1 = New_C1_()
				_ = _nw139
				_nw139.Ctor__(_this.F24(), true)
				_793_v47 = _nw139
				var _794_v48 _dafny.CodePoint
				_ = _794_v48
				_794_v48 = _dafny.CodePoint('c')
				var _795_v49 _dafny.Array
				_ = _795_v49
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw140
				_795_v49 = _nw140
				var _796_v50 _dafny.Array
				_ = _796_v50
				var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
				_ = _nw141
				_796_v50 = _nw141
				var _rhs143 _dafny.CodePoint = (_761_v21).Select((Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_761_v21).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs143
				var _rhs144 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_768_i3, _dafny.IntOfUint32((_758_v18).Cardinality())))
				_ = _rhs144
				var _rhs145 _dafny.Array = _795_v49
				_ = _rhs145
				var _rhs146 _dafny.Sequence = _761_v21
				_ = _rhs146
				var _rhs147 _dafny.Array = (func() _dafny.Array {
					if (_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool) {
						return _796_v50
					}
					return _796_v50
				})()
				_ = _rhs147
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				_794_v48 = _rhs143
				_lhs126.F8 = _rhs144
				_795_v49 = _rhs145
				_761_v21 = _rhs146
				_796_v50 = _rhs147
			}
			var _797_v51 *C3
			_ = _797_v51
			var _nw142 *C3 = New_C3_()
			_ = _nw142
			_nw142.Ctor__(!((_762_v22).IsDisjointFrom(_dafny.MultiSetOf(_734_v0))), (_dafny.Zero).Minus(_768_i3), (_this).F27(), _this.F24(), _734_v0)
			_797_v51 = _nw142
		}
		var _798_v52 _dafny.Map
		_ = _798_v52
		_798_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _764_v24)
		_798_v52 = _798_v52
		(globalState).F5 = _734_v0
		var _799_i6 _dafny.Int
		_ = _799_i6
		_799_i6 = _dafny.Zero
		{
			for _this.F24() {
				{
					if (_799_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_799_i6 = (_799_i6).Plus(_dafny.One)
					var _800_v53 _dafny.CodePoint
					_ = _800_v53
					_800_v53 = _dafny.CodePoint('d')
					var _801_v54 _dafny.Sequence
					_ = _801_v54
					_801_v54 = _dafny.SeqOf(_761_v21, _dafny.Companion_Sequence_.Update(_761_v21, (Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_761_v21).Cardinality()))).Uint32(), _800_v53))
					_801_v54 = _801_v54
					var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
					_ = _nw143
					_764_v24 = _nw143
					var _802_v55 _dafny.Set
					_ = _802_v55
					_802_v55 = _dafny.SetOf(!(_this.F24()), true)
					var _803_v56 _dafny.Map
					_ = _803_v56
					_803_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v55, (_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool))
					var _804_v58 _dafny.Sequence
					_ = _804_v58
					_804_v58 = _dafny.SeqOf(_803_v56)
					var _805_v59 _dafny.Map
					_ = _805_v59
					_805_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v0, _803_v56)
					var _806_v61 _dafny.Array
					_ = _806_v61
					var _nwElement0_29 _dafny.Map = _803_v56
					_ = _nwElement0_29
					var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(14))
					_ = _nw144
					(_nw144).ArraySet1(_nwElement0_29, 0)
					(_nw144).ArraySet1(_803_v56, 1)
					(_nw144).ArraySet1((func() _dafny.Map {
						if (_this).F27() {
							return func() _dafny.Map {
								var _coll28 = _dafny.NewMapBuilder()
								_ = _coll28
								for _iter30 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
									return func(arg36 _dafny.Int) interface{} {
										return coer36(arg36)
									}
								}((func(_807_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
									return func(_808_i7 _dafny.Int) _dafny.Set {
										return _807_v55
									}
								})(_802_v55))))).Elements()); ; {
									_compr_28, _ok30 := _iter30()
									if !_ok30 {
										break
									}
									var _809_v57 _dafny.Set
									_809_v57 = interface{}(_compr_28).(_dafny.Set)
									if (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
										return func(arg37 _dafny.Int) interface{} {
											return coer37(arg37)
										}
									}((func(_810_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
										return func(_808_i7 _dafny.Int) _dafny.Set {
											return _810_v55
										}
									})(_802_v55))))).Contains(_809_v57) {
										_coll28.Add(_809_v57, (_this).F27())
									}
								}
								return _coll28.ToMap()
							}()
						}
						return _803_v56
					})(), 2)
					(_nw144).ArraySet1((_803_v56).Merge((_804_v58).Select((Companion_Default___.SafeIndex(_734_v0, _dafny.IntOfUint32((_804_v58).Cardinality()))).Uint32()).(_dafny.Map)), 3)
					(_nw144).ArraySet1((func() _dafny.Map {
						if (_805_v59).Contains(_734_v0) {
							return (_805_v59).Get(_734_v0).(_dafny.Map)
						}
						return Companion_Default___.Fm35(globalState)
					})(), 4)
					(_nw144).ArraySet1(_803_v56, 5)
					(_nw144).ArraySet1((_803_v56).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool), (_this).F25()), _this.F24())), 6)
					(_nw144).ArraySet1(_803_v56, 7)
					(_nw144).ArraySet1((_803_v56).Merge(_803_v56), 8)
					(_nw144).ArraySet1(_803_v56, 9)
					(_nw144).ArraySet1(_803_v56, 10)
					(_nw144).ArraySet1((func() _dafny.Map {
						if (_752_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_752_v13), 0))).Int()).(bool) {
							return func() _dafny.Map {
								var _coll29 = _dafny.NewMapBuilder()
								_ = _coll29
								for _iter31 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
									return func(arg38 _dafny.Int) interface{} {
										return coer38(arg38)
									}
								}((func(_811_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
									return func(_812_i8 _dafny.Int) _dafny.Set {
										return _811_v55
									}
								})(_802_v55)))).Elements()); ; {
									_compr_29, _ok31 := _iter31()
									if !_ok31 {
										break
									}
									var _813_v60 _dafny.Set
									_813_v60 = interface{}(_compr_29).(_dafny.Set)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
										return func(arg39 _dafny.Int) interface{} {
											return coer39(arg39)
										}
									}((func(_814_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
										return func(_812_i8 _dafny.Int) _dafny.Set {
											return _814_v55
										}
									})(_802_v55))), _813_v60) {
										_coll29.Add(_813_v60, (_this).F27())
									}
								}
								return _coll29.ToMap()
							}()
						}
						return _803_v56
					})(), 11)
					(_nw144).ArraySet1(_803_v56, 12)
					(_nw144).ArraySet1(_803_v56, 13)
					_806_v61 = _nw144
					var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_806_v61), 0))
					_ = _index122
					(_806_v61).ArraySet1(_803_v56, (_index122).Int())
					var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_806_v61), 0))
					_ = _index123
					(_806_v61).ArraySet1(_803_v56, (_index123).Int())
					(globalState).F1 = _this.F24()
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = (func() bool {
			if ((_this).F25()) == ((_this).F25()) {
				return !((_this).F27())
			}
			return (_dafny.IntOfInt64(235)).Cmp(_734_v0) == 0
		})()
		var _815_v62 T0
		_ = _815_v62
		var _nw145 *C2 = New_C2_()
		_ = _nw145
		_nw145.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(_this.F24(), (_this).F27())).Cardinality()))
		_815_v62 = _nw145
		var _816_v63 D11
		_ = _816_v63
		_816_v63 = Companion_D11_.Create_DC30_(_815_v62)
		r1 = (_816_v63).Dtor_cf56()
		r2 = (_this.F24()) == (_this.F24())
		var _817_v65 _dafny.MultiSet
		_ = _817_v65
		_817_v65 = _dafny.MultiSetOf(_758_v18)
		r3 = (func() _dafny.Int {
			if (_this).F25() {
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_758_v18, false)).Merge(func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter32 := _dafny.Iterate((_817_v65).Elements()); ; {
						_compr_30, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _818_v64 _dafny.Sequence
						_818_v64 = interface{}(_compr_30).(_dafny.Sequence)
						if (_817_v65).Contains(_818_v64) {
							_coll30.Add(_818_v64, _this.F24())
						}
					}
					return _coll30.ToMap()
				}())).Cardinality()
			}
			return _734_v0
		})()
		return r0, r1, r2, r3
	}
}
func (_this *C4) F27() bool {
	{
		return _this._f27
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f24 bool
	_f25 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f24 = false
	_this._f25 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F24() bool {
	return _this._f24
}
func (_this *C5) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C5) F25() bool {
	return _this._f25
}
func (_this *C5) Ctor__(f24 bool, f25 bool) {
	{
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C5) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var _source7 D0 = Companion_D0_.Create_DC0_((_this).F25())
		_ = _source7
		if _source7.Is_DC1() {
			var _819___mcc_h0 bool = _source7.Get_().(D0_DC1).Cf1
			_ = _819___mcc_h0
			var _820___mcc_h1 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
			_ = _820___mcc_h1
			var _821_cf2 _dafny.Int = _820___mcc_h1
			_ = _821_cf2
			var _822_cf1 bool = _819___mcc_h0
			_ = _822_cf1
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_822_cf1, !(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_822_cf1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _822_cf1)))
		} else if _source7.Is_DC0() {
			var _823___mcc_h2 bool = _source7.Get_().(D0_DC0).Cf0
			_ = _823___mcc_h2
			var _824_cf0 bool = _823___mcc_h2
			_ = _824_cf0
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25()))
		} else {
			var _825___mcc_h3 D0 = _source7.Get_().(D0_DC2).Cf3
			_ = _825___mcc_h3
			var _826_cf3 D0 = _825___mcc_h3
			_ = _826_cf3
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F24()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F24()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())))
		}
	}
}
func (_this *C5) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}(func(_827_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(86), _dafny.IntOfInt64(469), _dafny.IntOfUint32((_dafny.SeqOf(_this.F24())).Cardinality()))
		})), _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(224))))).Cardinality())
	}
}
func (_this *C5) Fm13(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(939))).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(634))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality()))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}(func(_828_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))).Cardinality()))
	}
}
func (_this *C5) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _829_v0 _dafny.Int
		_ = _829_v0
		_829_v0 = _dafny.IntOfInt64(82)
		var _830_v1 _dafny.CodePoint
		_ = _830_v1
		_830_v1 = _dafny.CodePoint('v')
		var _831_v2 _dafny.Sequence
		_ = _831_v2
		_831_v2 = _dafny.SeqOf(_830_v1, _830_v1)
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_831_v2, _dafny.SeqOf(_830_v1, _830_v1))).Cardinality())
		_ = _hi7
		for _832_i0 := _829_v0; _832_i0.Cmp(_hi7) < 0; _832_i0 = _832_i0.Plus(_dafny.One) {
			var _833_v3 *C2
			_ = _833_v3
			var _nw146 *C2 = New_C2_()
			_ = _nw146
			_nw146.Ctor__(_829_v0)
			_833_v3 = _nw146
			if !(_this.F24()) {
				var _834_v4 _dafny.Map
				_ = _834_v4
				_834_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_832_i0, _dafny.IntOfInt64(134))
				_834_v4 = (_834_v4).Update(_832_i0, _832_i0)
				(globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("radkjif"), _831_v2)
				var _835_v5 _dafny.Array
				_ = _835_v5
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_26
				var _nw147 _dafny.Array
				_ = _nw147
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw147 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = func(_836_i1 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw147 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw147).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw147).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_835_v5 = _nw147
				var _837_v6 *C0
				_ = _837_v6
				var _nw148 *C0 = New_C0_()
				_ = _nw148
				_nw148.Ctor__(_dafny.SetOf(_835_v5, _835_v5, _835_v5, _835_v5), (_this).Fm13(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("arwqolr")).Cardinality()), (_this).F25(), globalState), _dafny.IntOfInt64(-476))
				_837_v6 = _nw148
				var _838_v7 _dafny.Map
				_ = _838_v7
				_838_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v6, _this.F24())
				_838_v7 = (_838_v7).Update(_837_v6, (_this).F25())
				var _839_v8 _dafny.Map
				_ = _839_v8
				_839_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_v5, _837_v6.F31)
				var _840_v9 _dafny.Set
				_ = _840_v9
				_840_v9 = _dafny.SetOf(_this.F24(), false, false)
				var _841_v10 _dafny.Array
				_ = _841_v10
				var _nwElement0_30 _dafny.Int = (_840_v9).Cardinality()
				_ = _nwElement0_30
				var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(2))
				_ = _nw149
				(_nw149).ArraySet1(_nwElement0_30, 0)
				(_nw149).ArraySet1(_829_v0, 1)
				_841_v10 = _nw149
				var _842_v11 D11
				_ = _842_v11
				_842_v11 = Companion_D11_.Create_DC31_(_841_v10, _832_i0, _this.F24())
				var _843_v12 _dafny.Array
				_ = _843_v12
				var _nwElement0_31 _dafny.CodePoint = _830_v1
				_ = _nwElement0_31
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
				_ = _nw150
				(_nw150).ArraySet1CodePoint(_nwElement0_31, 0)
				(_nw150).ArraySet1CodePoint(Companion_Default___.Fm26((_842_v11).Dtor_cf58(), (_this).F25(), globalState), 1)
				(_nw150).ArraySet1CodePoint(_830_v1, 2)
				(_nw150).ArraySet1CodePoint(_830_v1, 3)
				(_nw150).ArraySet1CodePoint(_dafny.CodePoint('y'), 4)
				(_nw150).ArraySet1CodePoint(_830_v1, 5)
				(_nw150).ArraySet1CodePoint(_830_v1, 6)
				(_nw150).ArraySet1CodePoint(_830_v1, 7)
				_843_v12 = _nw150
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_843_v12), 0))
				_ = _index124
				(_843_v12).ArraySet1CodePoint(_830_v1, (_index124).Int())
				var _844_v13 _dafny.Array
				_ = _844_v13
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw151
				_844_v13 = _nw151
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_844_v13), 0))
				_ = _index125
				(_844_v13).ArraySet1(_843_v12, (_index125).Int())
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_835_v5), 0))
				_ = _index126
				(_835_v5).ArraySet1((_this).F25(), (_index126).Int())
				var _845_v14 _dafny.Sequence
				_ = _845_v14
				_845_v14 = _dafny.SeqOf((_this).F25(), _this.F24())
				var _846_v15 _dafny.MultiSet
				_ = _846_v15
				_846_v15 = _dafny.MultiSetOf((_this).F25(), _this.F24(), ((_this).F25()) && (!(_this.F24())), (_845_v14).Select((Companion_Default___.SafeIndex((_833_v3).Fm2(globalState), _dafny.IntOfUint32((_845_v14).Cardinality()))).Uint32()).(bool), (_this).F25())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_843_v12), 0))
				_ = _index127
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_844_v13), 0))
				_ = _index128
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_835_v5), 0))
				_ = _index129
				var _rhs148 _dafny.Map = _839_v8
				_ = _rhs148
				var _rhs149 _dafny.CodePoint = _830_v1
				_ = _rhs149
				var _rhs150 _dafny.Int = (func() _dafny.Int {
					if (_846_v15).Contains(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm25(true, _this.F24(), _830_v1, globalState), _dafny.CodePoint('n'))) {
						return (_846_v15).Multiplicity(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm25(true, _this.F24(), _830_v1, globalState), _dafny.CodePoint('n')))
					}
					return _dafny.IntOfInt64(191)
				})()
				_ = _rhs150
				var _rhs151 _dafny.Array = _843_v12
				_ = _rhs151
				var _rhs152 bool = Companion_Default___.Fm29(_dafny.IntOfUint32((_831_v2).Cardinality()), _837_v6.F31, Companion_Default___.Fm29(_837_v6.F31, _837_v6.F31, _this.F24(), globalState), globalState)
				_ = _rhs152
				var _lhs127 _dafny.Array = _843_v12
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_843_v12), 0))
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 _dafny.Array = _844_v13
				_ = _lhs130
				var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_844_v13), 0))
				_ = _lhs131
				var _lhs132 _dafny.Array = _835_v5
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_835_v5), 0))
				_ = _lhs133
				_839_v8 = _rhs148
				(_lhs127).ArraySet1CodePoint(_rhs149, (_lhs128).Int())
				_lhs129.F17 = _rhs150
				(_lhs130).ArraySet1(_rhs151, (_lhs131).Int())
				(_lhs132).ArraySet1(_rhs152, (_lhs133).Int())
				r0 = (_835_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_835_v5), 0))).Int()).(bool)
			} else {
				var _847_v16 _dafny.MultiSet
				_ = _847_v16
				_847_v16 = _dafny.MultiSetOf(!(_this.F24()))
				var _848_v17 _dafny.Sequence
				_ = _848_v17
				_848_v17 = _dafny.SeqOf((_847_v16).IsDisjointFrom(_847_v16), (_this).F25())
				var _849_v18 _dafny.Map
				_ = _849_v18
				_849_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_830_v1, _829_v0)
				(globalState).F15 = (_848_v17).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_849_v18).Cardinality(), _829_v0), _dafny.IntOfUint32((_848_v17).Cardinality()))).Uint32()).(bool)
				(globalState).F14 = _829_v0
				var _850_v19 _dafny.Array
				_ = _850_v19
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw152
				_850_v19 = _nw152
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_850_v19), 0))
				_ = _index130
				(_850_v19).ArraySet1((func() _dafny.Sequence {
					if (_848_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.IntOfUint32((_848_v17).Cardinality()))).Uint32()).(bool) {
						return _831_v2
					}
					return _831_v2
				})(), (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_850_v19), 0))
				_ = _index131
				(_850_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_831_v2, _831_v2), _831_v2), (_index131).Int())
				var _851_v20 _dafny.Map
				_ = _851_v20
				_851_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_848_v17).Cardinality()), _dafny.IntOfUint32(((_850_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_850_v19), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				var _852_v21 T0
				_ = _852_v21
				var _nw153 *C2 = New_C2_()
				_ = _nw153
				_nw153.Ctor__((_851_v20).Cardinality())
				_852_v21 = _nw153
				var _853_v22 D11
				_ = _853_v22
				_853_v22 = Companion_D11_.Create_DC30_(_852_v21)
				_853_v22 = _853_v22
				var _854_v23 _dafny.Array
				_ = _854_v23
				var _nwElement0_32 bool = (_this).F25()
				_ = _nwElement0_32
				var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.One)
				_ = _nw154
				(_nw154).ArraySet1(_nwElement0_32, 0)
				_854_v23 = _nw154
				var _855_v24 _dafny.Array
				_ = _855_v24
				var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw155
				_855_v24 = _nw155
				var _856_v25 D11
				_ = _856_v25
				_856_v25 = Companion_D11_.Create_DC31_(_855_v24, _829_v0, (_848_v17).Select((Companion_Default___.SafeIndex((_852_v21).F23(), _dafny.IntOfUint32((_848_v17).Cardinality()))).Uint32()).(bool))
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_854_v23), 0))
				_ = _index132
				(_854_v23).ArraySet1((_856_v25).Dtor_cf59(), (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_854_v23), 0))
				_ = _index133
				(_854_v23).ArraySet1((_this).F25(), (_index133).Int())
			}
			if (_this).F25() {
				var _857_v26 _dafny.Array
				_ = _857_v26
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw156
				_857_v26 = _nw156
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))
				_ = _index134
				(_857_v26).ArraySet1(_829_v0, (_index134).Int())
				var _858_v27 _dafny.Sequence
				_ = _858_v27
				_858_v27 = _dafny.SeqOf(_857_v26)
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))
				_ = _index135
				var _rhs153 _dafny.Int = _829_v0
				_ = _rhs153
				var _rhs154 bool = (_832_i0).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-158), _dafny.IntOfInt64(526)))) >= 0
				_ = _rhs154
				var _rhs155 _dafny.Int = _dafny.IntOfUint32((_858_v27).Cardinality())
				_ = _rhs155
				var _rhs156 bool = (_this).F25()
				_ = _rhs156
				var _rhs157 _dafny.Int = _832_i0
				_ = _rhs157
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				var _lhs136 _dafny.Array = _857_v26
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))
				_ = _lhs137
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				_lhs134.F17 = _rhs153
				_lhs135.F9 = _rhs154
				(_lhs136).ArraySet1(_rhs155, (_lhs137).Int())
				_lhs138.F1 = _rhs156
				_lhs139.F14 = _rhs157
				var _859_v28 _dafny.Map
				_ = _859_v28
				_859_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _830_v1)
				var _860_v29 _dafny.MultiSet
				_ = _860_v29
				_860_v29 = _dafny.MultiSetOf(_859_v28, _859_v28)
				(globalState).F15 = (_860_v29).IsProperSubsetOf(_860_v29)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))
				_ = _index136
				var _rhs158 _dafny.Sequence = _858_v27
				_ = _rhs158
				var _rhs159 _dafny.Int = _832_i0
				_ = _rhs159
				var _rhs160 _dafny.Int = _829_v0
				_ = _rhs160
				var _lhs140 _dafny.Array = _857_v26
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))
				_ = _lhs141
				_858_v27 = _rhs158
				(_lhs140).ArraySet1(_rhs159, (_lhs141).Int())
				_829_v0 = _rhs160
				(globalState).F8 = (_829_v0).Plus(_832_i0)
				var _rhs161 bool = true
				_ = _rhs161
				var _rhs162 bool = !((_this).F25())
				_ = _rhs162
				var _rhs163 _dafny.Int = (_857_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_857_v26), 0))).Int()).(_dafny.Int)
				_ = _rhs163
				var _rhs164 _dafny.Int = _832_i0
				_ = _rhs164
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				_lhs142.F1 = _rhs161
				_lhs143.F1 = _rhs162
				_lhs144.F14 = _rhs163
				_lhs145.F17 = _rhs164
			} else {
				(globalState).F1 = _this.F24()
				(globalState).F17 = _829_v0
				var _861_v30 _dafny.Array
				_ = _861_v30
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw157
				_861_v30 = _nw157
				var _862_v31 _dafny.Sequence
				_ = _862_v31
				_862_v31 = _dafny.SeqOf(_861_v30, _861_v30, (func() _dafny.Array {
					if (_this).F25() {
						return _861_v30
					}
					return _861_v30
				})())
				_862_v31 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_862_v31, _862_v31), (Companion_Default___.SafeIndex(_832_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_862_v31, _862_v31)).Cardinality()))).Uint32(), _861_v30), _862_v31)
				(globalState).F8 = (Companion_Default___.SafeDivisionInt(_829_v0, (_dafny.Zero).Minus(_832_i0))).Minus(_832_i0)
				var _863_v32 _dafny.Set
				_ = _863_v32
				_863_v32 = _dafny.SetOf(!(_this.F24()))
				var _864_v33 _dafny.Array
				_ = _864_v33
				var _nwElement0_33 _dafny.Set = _863_v32
				_ = _nwElement0_33
				var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(3))
				_ = _nw158
				(_nw158).ArraySet1(_nwElement0_33, 0)
				(_nw158).ArraySet1(_863_v32, 1)
				(_nw158).ArraySet1(_863_v32, 2)
				_864_v33 = _nw158
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_864_v33), 0))
				_ = _index137
				(_864_v33).ArraySet1(_863_v32, (_index137).Int())
				var _865_v34 _dafny.Map
				_ = _865_v34
				_865_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _829_v0)
				var _866_v35 _dafny.Array
				_ = _866_v35
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw159
				_866_v35 = _nw159
				var _867_v36 D7
				_ = _867_v36
				_867_v36 = Companion_D7_.Create_DC20_((_this).F25(), _865_v34, _866_v35)
				var _868_v37 D7
				_ = _868_v37
				_868_v37 = Companion_D7_.Create_DC21_(_867_v36)
				var _869_v38 D7
				_ = _869_v38
				_869_v38 = Companion_D7_.Create_DC21_(_868_v37)
				var _pat_let_tv12 = _865_v34
				_ = _pat_let_tv12
				var _pat_let_tv13 = _866_v35
				_ = _pat_let_tv13
				var _870_v39 _dafny.MultiSet
				_ = _870_v39
				_870_v39 = _dafny.MultiSetOf(func(_pat_let7_0 D7) D7 {
					return func(_871_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let8_0 D7) D7 {
							return func(_872_dt__update_hcf38_h0 D7) D7 {
								return Companion_D7_.Create_DC21_(_872_dt__update_hcf38_h0)
							}(_pat_let8_0)
						}(Companion_D7_.Create_DC20_((_this).F25(), _pat_let_tv12, _pat_let_tv13))
					}(_pat_let7_0)
				}(_869_v38), _869_v38)
				var _873_v40 _dafny.Sequence
				_ = _873_v40
				_873_v40 = _dafny.SeqOf((_870_v39).Cardinality())
				var _874_v41 _dafny.Sequence
				_ = _874_v41
				_874_v41 = _dafny.SeqOf((_this).F25(), _this.F24())
				var _875_v42 D1
				_ = _875_v42
				_875_v42 = Companion_D1_.Create_DC4_((_this).F25(), _this.F24())
				var _876_v43 _dafny.Set
				_ = _876_v43
				_876_v43 = _dafny.SetOf(_875_v42)
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_864_v33), 0))
				_ = _index138
				(_864_v33).ArraySet1((_863_v32).Intersection((Companion_Default___.Fm32((_this).F25(), _873_v40, _874_v41, _876_v43, globalState)).Union(_863_v32)), (_index138).Int())
			}
			var _877_v44 _dafny.Array
			_ = _877_v44
			var _nw160 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw160
			_877_v44 = _nw160
			var _878_v45 _dafny.Map
			_ = _878_v45
			_878_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
			var _879_v46 _dafny.Map
			_ = _879_v46
			_879_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _832_i0)
			var _880_v47 D1
			_ = _880_v47
			_880_v47 = Companion_D1_.Create_DC5_((_this).F25(), _878_v45, _832_i0, _832_i0, _879_v46)
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_877_v44), 0))
			_ = _index139
			(_877_v44).ArraySet1(!((_880_v47).Dtor_cf7()), (_index139).Int())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_877_v44), 0))
			_ = _index140
			(_877_v44).ArraySet1(!(((_832_i0).Minus(_829_v0)).Cmp(_832_i0) == 0), (_index140).Int())
		}
		var _881_v48 _dafny.Sequence
		_ = _881_v48
		_881_v48 = _dafny.SeqOf((_this).F25())
		var _882_v49 _dafny.MultiSet
		_ = _882_v49
		_882_v49 = _dafny.MultiSetOf(false)
		var _883_v50 _dafny.Sequence
		_ = _883_v50
		_883_v50 = _dafny.SeqOf(_829_v0, _829_v0, (func() _dafny.Int {
			if (_882_v49).Contains(_this.F24()) {
				return (_882_v49).Multiplicity(_this.F24())
			}
			return _829_v0
		})())
		r0 = (_881_v48).Select((Companion_Default___.SafeIndex((_883_v50).Select((Companion_Default___.SafeIndex(_829_v0, _dafny.IntOfUint32((_883_v50).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_881_v48).Cardinality()))).Uint32()).(bool)
		var _884_v51 _dafny.Map
		_ = _884_v51
		_884_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _829_v0)
		var _885_v52 _dafny.Map
		_ = _885_v52
		_885_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _829_v0))
		var _886_v53 _dafny.Array
		_ = _886_v53
		var _nwElement0_34 _dafny.Int = (_829_v0).Minus((_dafny.Zero).Minus(_829_v0))
		_ = _nwElement0_34
		var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(18))
		_ = _nw161
		(_nw161).ArraySet1(_nwElement0_34, 0)
		(_nw161).ArraySet1((_884_v51).Cardinality(), 1)
		(_nw161).ArraySet1(_829_v0, 2)
		(_nw161).ArraySet1(_829_v0, 3)
		(_nw161).ArraySet1(_829_v0, 4)
		(_nw161).ArraySet1((_885_v52).Cardinality(), 5)
		(_nw161).ArraySet1(_829_v0, 6)
		(_nw161).ArraySet1(_829_v0, 7)
		(_nw161).ArraySet1(_829_v0, 8)
		(_nw161).ArraySet1(_dafny.IntOfUint32((_831_v2).Cardinality()), 9)
		(_nw161).ArraySet1(_829_v0, 10)
		(_nw161).ArraySet1((_dafny.Zero).Minus(_829_v0), 11)
		(_nw161).ArraySet1(_829_v0, 12)
		(_nw161).ArraySet1(Companion_Default___.SafeModuloInt(_829_v0, _829_v0), 13)
		(_nw161).ArraySet1((_dafny.Zero).Minus(_829_v0), 14)
		(_nw161).ArraySet1(_829_v0, 15)
		(_nw161).ArraySet1(_dafny.IntOfUint32((_883_v50).Cardinality()), 16)
		(_nw161).ArraySet1((func() _dafny.Int {
			if _this.F24() {
				return _829_v0
			}
			return _829_v0
		})(), 17)
		_886_v53 = _nw161
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_886_v53), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _887_i2 _dafny.Int
			_887_i2 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_887_i2).Sign() != -1) && ((_887_i2).Cmp(_dafny.ArrayLen((_886_v53), 0)) < 0)) {
				(_886_v53).ArraySet1((_887_i2).Plus((func() _dafny.Int {
					if (_this).F25() {
						return _829_v0
					}
					return _829_v0
				})()), (_887_i2).Int())
			}
		}
		(globalState).F5 = _829_v0
		var _888_v54 _dafny.Map
		_ = _888_v54
		_888_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_884_v51, (_this).F25())
		var _889_v57 _dafny.Map
		_ = _889_v57
		_889_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _829_v0)
		var _890_v58 _dafny.Array
		_ = _890_v58
		var _nwElement0_35 bool = (_this).F25()
		_ = _nwElement0_35
		var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(28))
		_ = _nw162
		(_nw162).ArraySet1(_nwElement0_35, 0)
		(_nw162).ArraySet1(_this.F24(), 1)
		(_nw162).ArraySet1(((_881_v48).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_883_v50).Select((Companion_Default___.SafeIndex(_829_v0, _dafny.IntOfUint32((_883_v50).Cardinality()))).Uint32()).(_dafny.Int))), _dafny.IntOfUint32((_881_v48).Cardinality()))).Uint32()).(bool)) == (!(_this.F24())), 2)
		(_nw162).ArraySet1(_this.F24(), 3)
		(_nw162).ArraySet1((_888_v54).Equals(_888_v54), 4)
		(_nw162).ArraySet1(((_this).F25()) || (true), 5)
		(_nw162).ArraySet1(_this.F24(), 6)
		(_nw162).ArraySet1((_this).F25(), 7)
		(_nw162).ArraySet1(_this.F24(), 8)
		(_nw162).ArraySet1((_this).F25(), 9)
		(_nw162).ArraySet1(true, 10)
		(_nw162).ArraySet1((_this).F25(), 11)
		(_nw162).ArraySet1((_this).F25(), 12)
		(_nw162).ArraySet1((_this).F25(), 13)
		(_nw162).ArraySet1(Companion_Default___.Fm29((func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(322), _dafny.IntOfInt64(256))); ; {
				_compr_31, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _891_v55 _dafny.Int
				_891_v55 = interface{}(_compr_31).(_dafny.Int)
				if ((_dafny.IntOfInt64(322)).Cmp(_891_v55) <= 0) && ((_891_v55).Cmp(_dafny.IntOfInt64(256)) < 0) {
					_coll31.Add((_891_v55).Plus(_829_v0), _this.F24())
				}
			}
			return _coll31.ToMap()
		}()).Cardinality(), (func() _dafny.Map {
			var _coll32 = _dafny.NewMapBuilder()
			_ = _coll32
			for _iter35 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_883_v50, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.IntOfUint32((_883_v50).Cardinality()))).Uint32(), (_882_v49).Cardinality())).Elements()); ; {
				_compr_32, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _892_v56 _dafny.Int
				_892_v56 = interface{}(_compr_32).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_883_v50, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.IntOfUint32((_883_v50).Cardinality()))).Uint32(), (_882_v49).Cardinality()), _892_v56) {
					_coll32.Add((_892_v56).Minus(_829_v0), (_this).F25())
				}
			}
			return _coll32.ToMap()
		}()).Cardinality(), (_this).F25(), globalState), 14)
		(_nw162).ArraySet1(_this.F24(), 15)
		(_nw162).ArraySet1(_this.F24(), 16)
		(_nw162).ArraySet1(((func() _dafny.Int {
			if (_889_v57).Contains(_829_v0) {
				return (_889_v57).Get(_829_v0).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(_829_v0)
		})()).Cmp((_dafny.Zero).Minus(_829_v0)) != 0, 17)
		(_nw162).ArraySet1((_882_v49).IsProperSubsetOf(_882_v49), 18)
		(_nw162).ArraySet1(_this.F24(), 19)
		(_nw162).ArraySet1(!((_this).F25()), 20)
		(_nw162).ArraySet1((_this).F25(), 21)
		(_nw162).ArraySet1(Companion_Default___.Fm29(_829_v0, _829_v0, _this.F24(), globalState), 22)
		(_nw162).ArraySet1(!(_this.F24()), 23)
		(_nw162).ArraySet1(!((_829_v0).Cmp(_829_v0) <= 0), 24)
		(_nw162).ArraySet1((_dafny.IntOfUint32((_883_v50).Cardinality())).Cmp(_829_v0) < 0, 25)
		(_nw162).ArraySet1(false, 26)
		(_nw162).ArraySet1((_this).F25(), 27)
		_890_v58 = _nw162
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_890_v58), 0))
		_ = _index141
		(_890_v58).ArraySet1(true, (_index141).Int())
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_890_v58), 0))
		_ = _index142
		(_890_v58).ArraySet1(_this.F24(), (_index142).Int())
		var _893_v59 *C1
		_ = _893_v59
		var _nw163 *C1 = New_C1_()
		_ = _nw163
		_nw163.Ctor__(_this.F24(), true)
		_893_v59 = _nw163
		var _894_v60 _dafny.Map
		_ = _894_v60
		_894_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _893_v59)
		var _rhs165 _dafny.MultiSet = _dafny.MultiSetOf((_890_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_890_v58), 0))).Int()).(bool), (_this).F25(), (_881_v48).Select((Companion_Default___.SafeIndex(_829_v0, _dafny.IntOfUint32((_881_v48).Cardinality()))).Uint32()).(bool))
		_ = _rhs165
		var _rhs166 bool = Companion_Default___.Fm29(_829_v0, _829_v0, (_829_v0).Cmp((_894_v60).Cardinality()) == 0, globalState)
		_ = _rhs166
		var _lhs146 *GlobalState = globalState
		_ = _lhs146
		_882_v49 = _rhs165
		_lhs146.F9 = _rhs166
		r0 = (_829_v0).Cmp((_dafny.Zero).Minus(_829_v0)) >= 0
		var _nw164 *C3 = New_C3_()
		_ = _nw164
		_nw164.Ctor__(_this.F24(), _829_v0, (_890_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_890_v58), 0))).Int()).(bool), (_890_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_890_v58), 0))).Int()).(bool), (_dafny.Zero).Minus((_dafny.Zero).Minus(_829_v0)))
		r1 = _nw164
		r2 = _this.F24()
		r3 = _829_v0
		return r0, r1, r2, r3
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f24 bool
	_f23 _dafny.Int
	_f25 bool
	F26  D0
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f24 = false
	_this._f23 = _dafny.Zero
	_this._f25 = false
	_this.F26 = Companion_D0_.Default()
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

func (_this *C6) F24() bool {
	return _this._f24
}
func (_this *C6) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C6) F23() _dafny.Int {
	return _this._f23
}
func (_this *C6) F25() bool {
	return _this._f25
}
func (_this *C6) Ctor__(f26 D0, f24 bool, f25 bool, f23 _dafny.Int) {
	{
		(_this).F26 = f26
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f23 = f23
	}
}
func (_this *C6) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24()))))
	}
}
func (_this *C6) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F23()
	}
}
func (_this *C6) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F23()
	}
}
func (_this *C6) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F25()
	}
}
func (_this *C6) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _895_v0 _dafny.Sequence
		_ = _895_v0
		_895_v0 = _dafny.UnicodeSeqOfUtf8Bytes("koa")
		var _896_v1 _dafny.Map
		_ = _896_v1
		_896_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v0, _this.F24())
		_896_v1 = (_896_v1).Update(_dafny.UnicodeSeqOfUtf8Bytes("xx"), !(_this.F24()))
		var _897_v2 _dafny.Array
		_ = _897_v2
		var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
		_ = _nw165
		_897_v2 = _nw165
		var _898_v3 _dafny.CodePoint
		_ = _898_v3
		_898_v3 = _dafny.CodePoint('r')
		var _899_v4 _dafny.Array
		_ = _899_v4
		var _out9 _dafny.Array
		_ = _out9
		_out9 = (_this).M4(_897_v2, (_this).F23(), !(!(_this.F24())), _898_v3, globalState)
		_899_v4 = _out9
		var _900_v5 _dafny.Map
		_ = _900_v5
		_900_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F23())
		var _901_i0 _dafny.Int
		_ = _901_i0
		_901_i0 = _dafny.Zero
		{
			for !((func() bool {
				if (Companion_D1_.Create_DC5_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25()), (_this).F23(), (_this).F23(), _900_v5)).Dtor_cf7() {
					return ((_this).F23()).Cmp(_dafny.IntOfInt64(-781)) == 0
				}
				return (_this).F25()
			})()) {
				{
					if (_901_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_901_i0 = (_901_i0).Plus(_dafny.One)
					r2 = (_this).F25()
					var _902_v6 _dafny.Sequence
					_ = _902_v6
					_902_v6 = _dafny.SeqOf((_this).F25())
					var _903_v7 _dafny.Array
					_ = _903_v7
					var _len0_27 _dafny.Int = _dafny.IntOfInt64(11)
					_ = _len0_27
					var _nw166 _dafny.Array
					_ = _nw166
					if _len0_27.Cmp(_dafny.Zero) == 0 {
						_nw166 = _dafny.NewArray(_len0_27)
					} else {
						var _init27 func(_dafny.Int) _dafny.Int = func(_904_i1 _dafny.Int) _dafny.Int {
							return (_904_i1).Minus((_this).F23())
						}
						_ = _init27
						var _element0_27 = _init27(_dafny.Zero)
						_ = _element0_27
						_nw166 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
						(_nw166).ArraySet1(_element0_27, 0)
						var _nativeLen0_27 = (_len0_27).Int()
						_ = _nativeLen0_27
						for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
							(_nw166).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
						}
					}
					_903_v7 = _nw166
					var _905_v8 _dafny.Set
					_ = _905_v8
					_905_v8 = _dafny.SetOf(_903_v7)
					var _906_v9 _dafny.Sequence
					_ = _906_v9
					_906_v9 = _dafny.SeqOf((_this).F23())
					var _907_v10 _dafny.Map
					_ = _907_v10
					_907_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), !((_this).F25()))
					var _908_v11 _dafny.Array
					_ = _908_v11
					var _nwElement0_36 _dafny.Int = (_this).F23()
					_ = _nwElement0_36
					var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(26))
					_ = _nw167
					(_nw167).ArraySet1(_nwElement0_36, 0)
					(_nw167).ArraySet1((_this).F23(), 1)
					(_nw167).ArraySet1((func() _dafny.Int {
						if (_this).F25() {
							return (_this).F23()
						}
						return _dafny.IntOfUint32((_902_v6).Cardinality())
					})(), 2)
					(_nw167).ArraySet1(_dafny.IntOfUint32((_902_v6).Cardinality()), 3)
					(_nw167).ArraySet1(_dafny.IntOfUint32((_895_v0).Cardinality()), 4)
					(_nw167).ArraySet1((_this).F23(), 5)
					(_nw167).ArraySet1(_dafny.IntOfInt64(142), 6)
					(_nw167).ArraySet1((_this).F23(), 7)
					(_nw167).ArraySet1((_this).F23(), 8)
					(_nw167).ArraySet1((_dafny.Zero).Minus((_this).F23()), 9)
					(_nw167).ArraySet1((_this).F23(), 10)
					(_nw167).ArraySet1(((_this).F23()).Minus((_this).F23()), 11)
					(_nw167).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm11((_this).F23(), _this.F24(), _this.F24(), _dafny.CodePoint('a'), globalState))).Cardinality()), 12)
					(_nw167).ArraySet1(_dafny.IntOfInt64(89), 13)
					(_nw167).ArraySet1((func() _dafny.Int {
						if _this.F24() {
							return (_this).F23()
						}
						return (_this).F23()
					})(), 14)
					(_nw167).ArraySet1(_dafny.IntOfInt64(-763), 15)
					(_nw167).ArraySet1((_this).F23(), 16)
					(_nw167).ArraySet1((_this).F23(), 17)
					(_nw167).ArraySet1((_this).F23(), 18)
					(_nw167).ArraySet1((_dafny.Zero).Minus(((_this).F23()).Plus((_905_v8).Cardinality())), 19)
					(_nw167).ArraySet1((_this).F23(), 20)
					(_nw167).ArraySet1((_dafny.Zero).Minus((_this).F23()), 21)
					(_nw167).ArraySet1((_this).F23(), 22)
					(_nw167).ArraySet1((_dafny.IntOfUint32((_906_v9).Cardinality())).Plus((_this).F23()), 23)
					(_nw167).ArraySet1(Companion_Default___.Fm1((Companion_D1_.Create_DC5_(!(false), _907_v10, (_this).F23(), (_this).F23(), _900_v5)).Dtor_cf10(), (_this).F25(), globalState), 24)
					(_nw167).ArraySet1((_this).Fm2(globalState), 25)
					_908_v11 = _nw167
					var _909_v12 _dafny.MultiSet
					_ = _909_v12
					_909_v12 = _dafny.MultiSetOf(_908_v11, _908_v11, _903_v7, _908_v11)
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_903_v7), 0))
					_ = _index143
					(_903_v7).ArraySet1((func() _dafny.Int {
						if (_909_v12).Contains(_908_v11) {
							return (_909_v12).Multiplicity(_908_v11)
						}
						return (_this).F23()
					})(), (_index143).Int())
					var _910_v13 D1
					_ = _910_v13
					_910_v13 = Companion_D1_.Create_DC5_(_this.F24(), _907_v10, _dafny.IntOfUint32((_906_v9).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F25())).Cardinality(), _900_v5)
					var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_903_v7), 0))
					_ = _index144
					(_903_v7).ArraySet1((_910_v13).Dtor_cf9(), (_index144).Int())
					_898_v3 = Companion_Default___.Fm12((_this).F25(), globalState)
					var _911_v14 T1
					_ = _911_v14
					var _nw168 *C5 = New_C5_()
					_ = _nw168
					_nw168.Ctor__(_this.F24(), (_this).F25())
					_911_v14 = _nw168
					var _912_v15 _dafny.Array
					_ = _912_v15
					var _out10 _dafny.Array
					_ = _out10
					_out10 = (_this).M4(_897_v2, (_this).F23(), (_this).Fm3(_895_v0, _this.F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_911_v14, _907_v10)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(930))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}(func(_913_i2 _dafny.Int) _dafny.Int {
						return (_this).F23()
					})), globalState), _dafny.CodePoint('a'), globalState)
					_912_v15 = _out10
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))
		_ = _index145
		(_899_v4).ArraySet1(_this.F24(), (_index145).Int())
		var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))
		_ = _index146
		(_899_v4).ArraySet1(false, (_index146).Int())
		var _914_v16 _dafny.Sequence
		_ = _914_v16
		_914_v16 = _dafny.SeqOf((_this).F23())
		var _915_v18 _dafny.Set
		_ = _915_v18
		_915_v18 = _dafny.SetOf((_this).F23(), (_dafny.Zero).Minus((_this).F23()), (_this).F23(), (_this).F23(), (_this).F23())
		var _916_v19 D0
		_ = _916_v19
		_916_v19 = Companion_D0_.Create_DC0_((_this).F25())
		var _917_v20 D0
		_ = _917_v20
		_917_v20 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(_916_v19))
		var _pat_let_tv14 = _917_v20
		_ = _pat_let_tv14
		var _918_v21 _dafny.Array
		_ = _918_v21
		var _nwElement0_37 D0 = _this.F26
		_ = _nwElement0_37
		var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(21))
		_ = _nw169
		(_nw169).ArraySet1(_nwElement0_37, 0)
		(_nw169).ArraySet1(_this.F26, 1)
		(_nw169).ArraySet1(_this.F26, 2)
		(_nw169).ArraySet1(_this.F26, 3)
		(_nw169).ArraySet1(_this.F26, 4)
		(_nw169).ArraySet1(_this.F26, 5)
		(_nw169).ArraySet1(_this.F26, 6)
		(_nw169).ArraySet1(_this.F26, 7)
		(_nw169).ArraySet1(Companion_D0_.Create_DC2_(_916_v19), 8)
		(_nw169).ArraySet1(_this.F26, 9)
		(_nw169).ArraySet1(_this.F26, 10)
		(_nw169).ArraySet1(_this.F26, 11)
		(_nw169).ArraySet1(func(_pat_let9_0 D0) D0 {
			return func(_919_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let10_0 D0) D0 {
					return func(_920_dt__update_hcf3_h0 D0) D0 {
						return Companion_D0_.Create_DC2_(_920_dt__update_hcf3_h0)
					}(_pat_let10_0)
				}(_pat_let_tv14)
			}(_pat_let9_0)
		}(_this.F26), 12)
		(_nw169).ArraySet1(_this.F26, 13)
		(_nw169).ArraySet1(_this.F26, 14)
		(_nw169).ArraySet1(_this.F26, 15)
		(_nw169).ArraySet1(_this.F26, 16)
		(_nw169).ArraySet1(_this.F26, 17)
		(_nw169).ArraySet1(_this.F26, 18)
		(_nw169).ArraySet1(_this.F26, 19)
		(_nw169).ArraySet1(_this.F26, 20)
		_918_v21 = _nw169
		var _921_v22 _dafny.Array
		_ = _921_v22
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_28
		var _nw170 _dafny.Array
		_ = _nw170
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw170 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) _dafny.Int = func(_922_i3 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_922_i3, (_this).F23())
			}
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw170 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw170).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw170).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_921_v22 = _nw170
		var _923_v23 _dafny.Map
		_ = _923_v23
		_923_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC9_((_this).F23()), _921_v22)
		var _924_v24 _dafny.Array
		_ = _924_v24
		var _nwElement0_38 _dafny.Int = _dafny.IntOfUint32((_914_v16).Cardinality())
		_ = _nwElement0_38
		var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(13))
		_ = _nw171
		(_nw171).ArraySet1(_nwElement0_38, 0)
		(_nw171).ArraySet1(((_this).F23()).Plus(_dafny.IntOfInt64(968)), 1)
		(_nw171).ArraySet1(Companion_Default___.Fm1((_dafny.Zero).Minus((func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter36 := _dafny.Iterate((_915_v18).Elements()); ; {
				_compr_33, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _925_v17 _dafny.Int
				_925_v17 = interface{}(_compr_33).(_dafny.Int)
				if (_915_v18).Contains(_925_v17) {
					_coll33.Add(Companion_Default___.SafeModuloInt(_925_v17, _dafny.IntOfInt64(529)), _this.F24())
				}
			}
			return _coll33.ToMap()
		}()).Cardinality()), (_this).F25(), globalState), 2)
		(_nw171).ArraySet1((_this).F23(), 3)
		(_nw171).ArraySet1((_dafny.Zero).Minus(((_this).F23()).Plus((_this).F23())), 4)
		(_nw171).ArraySet1(Companion_Default___.SafeModuloInt((_this).F23(), (Companion_D5_.Create_DC15_((_this).F23(), _898_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_918_v21, _898_v3), _923_v23)).Dtor_cf25()), 5)
		(_nw171).ArraySet1((_this).F23(), 6)
		(_nw171).ArraySet1((_this).F23(), 7)
		(_nw171).ArraySet1((_this).F23(), 8)
		(_nw171).ArraySet1((_this).F23(), 9)
		(_nw171).ArraySet1(((_this).F23()).Minus((_this).F23()), 10)
		(_nw171).ArraySet1((_this).F23(), 11)
		(_nw171).ArraySet1((_this).F23(), 12)
		_924_v24 = _nw171
		var _926_v25 _dafny.Map
		_ = _926_v25
		_926_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_924_v24, _this.F24())
		var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))
		_ = _index147
		var _rhs167 bool = (_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool)
		_ = _rhs167
		var _rhs168 _dafny.Array = (func() _dafny.Array {
			if (_dafny.IntOfInt64(140)).Cmp(_dafny.IntOfInt64(926)) != 0 {
				return _924_v24
			}
			return _921_v22
		})()
		_ = _rhs168
		var _rhs169 _dafny.Map = _926_v25
		_ = _rhs169
		var _lhs147 _dafny.Array = _899_v4
		_ = _lhs147
		var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))
		_ = _lhs148
		(_lhs147).ArraySet1(_rhs167, (_lhs148).Int())
		_921_v22 = _rhs168
		_926_v25 = _rhs169
		var _hi8 _dafny.Int = (_this).F23()
		_ = _hi8
		for _927_i4 := _dafny.IntOfInt64(825); _927_i4.Cmp(_hi8) < 0; _927_i4 = _927_i4.Plus(_dafny.One) {
			_900_v5 = ((Companion_Default___.Fm36(globalState)).Update((_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _this.F24())).Cardinality())).Update((_this).F25(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F25()), _dafny.IntOfInt64(285))).Merge(_900_v5)).Cardinality())
			var _928_v26 *C4
			_ = _928_v26
			var _nw172 *C4 = New_C4_()
			_ = _nw172
			_nw172.Ctor__(!((_this).F25()), _this.F24(), !((_this).F25()))
			_928_v26 = _nw172
			r3 = (_this).F23()
			var _929_v27 *C4
			_ = _929_v27
			var _nw173 *C4 = New_C4_()
			_ = _nw173
			_nw173.Ctor__(((_this).F25()) && (!((_this).F25())), (_this).F25(), !((_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool)) || (_this.F24()))
			_929_v27 = _nw173
			var _930_v28 _dafny.Sequence
			_ = _930_v28
			_930_v28 = _dafny.SeqOf((_929_v27).F27(), (_928_v26).F27(), (_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool))
			var _rhs170 bool = _this.F24()
			_ = _rhs170
			var _rhs171 bool = (_930_v28).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_927_i4, (_this).F23()), _dafny.IntOfUint32((_930_v28).Cardinality()))).Uint32()).(bool)
			_ = _rhs171
			var _rhs172 *C4 = _928_v26
			_ = _rhs172
			var _rhs173 bool = (_929_v27).F27()
			_ = _rhs173
			var _rhs174 _dafny.Array = _899_v4
			_ = _rhs174
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			var _lhs150 *C6 = _this
			_ = _lhs150
			r0 = _rhs170
			_lhs149.F6 = _rhs171
			_929_v27 = _rhs172
			_lhs150.F24_set_(_rhs173)
			_899_v4 = _rhs174
		}
		var _931_v29 _dafny.Sequence
		_ = _931_v29
		_931_v29 = _dafny.SeqOf((_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool), (_this).F25(), _this.F24())
		r0 = _dafny.Companion_Sequence_.Contains(_931_v29, (_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool))
		var _932_v30 T0
		_ = _932_v30
		var _nw174 *C2 = New_C2_()
		_ = _nw174
		_nw174.Ctor__((_this).F23())
		_932_v30 = _nw174
		r1 = _932_v30
		r2 = (_this).F25()
		var _933_v31 _dafny.Map
		_ = _933_v31
		_933_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_899_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_899_v4), 0))).Int()).(bool))
		var _934_v32 D1
		_ = _934_v32
		_934_v32 = Companion_D1_.Create_DC5_(_this.F24(), _933_v31, (_this).F23(), (_this).F23(), _900_v5)
		var _935_v33 _dafny.Map
		_ = _935_v33
		_935_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_932_v30).F23(), (_934_v32).Dtor_cf9())
		var _936_v34 _dafny.Map
		_ = _936_v34
		_936_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_932_v30).F23(), (_935_v33).Cardinality())
		r3 = Companion_Default___.SafeDivisionInt((_932_v30).F23(), ((_936_v34).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_932_v30).F23(), (_this).F23()))).Cardinality())
		return r0, r1, r2, r3
	}
}
func (_this *C6) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var _937_v0 _dafny.Sequence
		_ = _937_v0
		_937_v0 = _dafny.SeqOf((_this).F23())
		var _938_v1 D10
		_ = _938_v1
		_938_v1 = Companion_D10_.Create_DC28_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F23(), (_this).F23()), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F23(), (_this).F23())).Cardinality()))).Uint32(), _dafny.IntOfInt64(952)))
		var _939_v2 _dafny.Sequence
		_ = _939_v2
		_939_v2 = _dafny.UnicodeSeqOfUtf8Bytes("itfta")
		var _940_v3 _dafny.Array
		_ = _940_v3
		var _nwElement0_39 _dafny.Sequence = _937_v0
		_ = _nwElement0_39
		var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(18))
		_ = _nw175
		(_nw175).ArraySet1(_nwElement0_39, 0)
		(_nw175).ArraySet1(_937_v0, 1)
		(_nw175).ArraySet1(_937_v0, 2)
		(_nw175).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(-698)), 3)
		(_nw175).ArraySet1(_937_v0, 4)
		(_nw175).ArraySet1(_937_v0, 5)
		(_nw175).ArraySet1((_938_v1).Dtor_cf52(), 6)
		(_nw175).ArraySet1(_937_v0, 7)
		(_nw175).ArraySet1(_dafny.Companion_Sequence_.Update(_937_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_939_v2).Cardinality()), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32(), (_this).F23()), 8)
		(_nw175).ArraySet1(Companion_Default___.Fm31(p1, true, globalState), 9)
		(_nw175).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg43 _dafny.Int) interface{} {
				return coer43(arg43)
			}
		}(func(_941_i1 _dafny.Int) _dafny.Int {
			return (_this).F23()
		})), 10)
		(_nw175).ArraySet1(_937_v0, 11)
		(_nw175).ArraySet1(_937_v0, 12)
		(_nw175).ArraySet1(_937_v0, 13)
		(_nw175).ArraySet1(_937_v0, 14)
		(_nw175).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_937_v0, _937_v0), 15)
		(_nw175).ArraySet1(_937_v0, 16)
		(_nw175).ArraySet1(_dafny.SeqOf((_this).F23()), 17)
		_940_v3 = _nw175
		for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_940_v3), 0))); ; {
			_guard_loop_3, _ok37 := _iter37()
			if !_ok37 {
				break
			}
			var _942_i0 _dafny.Int
			_942_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_942_i0).Sign() != -1) && ((_942_i0).Cmp(_dafny.ArrayLen((_940_v3), 0)) < 0)) {
				(_940_v3).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_937_v0, _dafny.Companion_Sequence_.Update(_937_v0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32(), (_this).F23())), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_937_v0, _dafny.Companion_Sequence_.Update(_937_v0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32(), (_this).F23()))).Cardinality()))).Uint32(), (_this).F23()), (_942_i0).Int())
			}
		}
		var _943_v4 _dafny.Array
		_ = _943_v4
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_29
		var _nw176 _dafny.Array
		_ = _nw176
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw176 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) bool = func(_944_i2 _dafny.Int) bool {
				return true
			}
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw176 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw176).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw176).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_943_v4 = _nw176
		_943_v4 = _943_v4
		(globalState).F15 = (((_this).F23()).Plus((_this).F23())).Cmp((_this).F23()) >= 0
		_937_v0 = _937_v0
		var _945_v5 _dafny.Map
		_ = _945_v5
		_945_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), p1)
		var _946_v6 _dafny.Set
		_ = _946_v6
		_946_v6 = _dafny.SetOf((_this).F25())
		var _947_v7 _dafny.Sequence
		_ = _947_v7
		_947_v7 = _dafny.SeqOf(_946_v6)
		var _948_v8 _dafny.Sequence
		_ = _948_v8
		_948_v8 = _dafny.SeqOf(_939_v2)
		var _949_v9 _dafny.Sequence
		_ = _949_v9
		_949_v9 = _dafny.SeqOf((_this).F25())
		var _950_v10 _dafny.Array
		_ = _950_v10
		var _len0_30 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_30
		var _nw177 _dafny.Array
		_ = _nw177
		if _len0_30.Cmp(_dafny.Zero) == 0 {
			_nw177 = _dafny.NewArray(_len0_30)
		} else {
			var _init30 func(_dafny.Int) _dafny.Int = func(_951_i4 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_951_i4, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sl")).Cardinality()))
			}
			_ = _init30
			var _element0_30 = _init30(_dafny.Zero)
			_ = _element0_30
			_nw177 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
			(_nw177).ArraySet1(_element0_30, 0)
			var _nativeLen0_30 = (_len0_30).Int()
			_ = _nativeLen0_30
			for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
				(_nw177).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
			}
		}
		_950_v10 = _nw177
		var _952_v11 _dafny.MultiSet
		_ = _952_v11
		_952_v11 = _dafny.MultiSetOf(_950_v10)
		var _953_v12 _dafny.MultiSet
		_ = _953_v12
		_953_v12 = _dafny.MultiSetOf(_dafny.SeqOf(false, _this.F24(), (_this).F25(), _this.F24(), (_this).F25()))
		var _954_v13 _dafny.MultiSet
		_ = _954_v13
		_954_v13 = _dafny.MultiSetOf(_this.F24())
		var _955_v14 _dafny.Set
		_ = _955_v14
		_955_v14 = _dafny.SetOf(_954_v13)
		var _956_v15 _dafny.Array
		_ = _956_v15
		var _nwElement0_40 _dafny.Int = (_dafny.Zero).Minus((_937_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_939_v2).Cardinality()), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _nwElement0_40
		var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(21))
		_ = _nw178
		(_nw178).ArraySet1(_nwElement0_40, 0)
		(_nw178).ArraySet1((_this).F23(), 1)
		(_nw178).ArraySet1((_this).F23(), 2)
		(_nw178).ArraySet1((_this).F23(), 3)
		(_nw178).ArraySet1((_this).F23(), 4)
		(_nw178).ArraySet1((_this).F23(), 5)
		(_nw178).ArraySet1((_dafny.Zero).Minus((_937_v0).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _945_v5)).Cardinality(), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32()).(_dafny.Int)), 6)
		(_nw178).ArraySet1((_this).F23(), 7)
		(_nw178).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_947_v7).Cardinality())), 8)
		(_nw178).ArraySet1((_this).F23(), 9)
		(_nw178).ArraySet1(Companion_Default___.SafeModuloInt((_this).F23(), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_948_v8).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_948_v8).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))), 10)
		(_nw178).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), 11)
		(_nw178).ArraySet1(Companion_Default___.Fm1((_this).F23(), (_this).F25(), globalState), 12)
		(_nw178).ArraySet1(((func() _dafny.MultiSet {
			if (_949_v9).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_949_v9).Cardinality()))).Uint32()).(bool) {
				return _952_v11
			}
			return _952_v11
		})()).Cardinality(), 13)
		(_nw178).ArraySet1(((_this).F23()).Minus((_this).F23()), 14)
		(_nw178).ArraySet1((_this).F23(), 15)
		(_nw178).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_939_v2).Cardinality()), _dafny.IntOfInt64(996)), 16)
		(_nw178).ArraySet1((func() _dafny.Int {
			if (_953_v12).Contains(_949_v9) {
				return (_953_v12).Multiplicity(_949_v9)
			}
			return (_this).F23()
		})(), 17)
		(_nw178).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), 18)
		(_nw178).ArraySet1((_955_v14).Cardinality(), 19)
		(_nw178).ArraySet1((_this).F23(), 20)
		_956_v15 = _nw178
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_950_v10), 0))); ; {
			_guard_loop_4, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _957_i3 _dafny.Int
			_957_i3 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_957_i3).Sign() != -1) && ((_957_i3).Cmp(_dafny.ArrayLen((_950_v10), 0)) < 0)) {
				(_950_v10).ArraySet1((_957_i3).Times(((_this).F23()).Plus(_dafny.IntOfUint32((_937_v0).Cardinality()))), (_957_i3).Int())
			}
		}
		var _hi9 _dafny.Int = (_937_v0).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_937_v0).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi9
		for _958_i5 := ((_946_v6).Cardinality()).Plus((_this).F23()); _958_i5.Cmp(_hi9) < 0; _958_i5 = _958_i5.Plus(_dafny.One) {
			(globalState).F14 = (Companion_Default___.SafeDivisionInt(_958_i5, _958_i5)).Plus((_this).F23())
			var _959_v16 _dafny.Map
			_ = _959_v16
			_959_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), p0)
			var _960_v17 _dafny.MultiSet
			_ = _960_v17
			_960_v17 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(757))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}(func(_961_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			}))).Cardinality()), (Companion_Default___.Fm37((_this).F25(), (_959_v16).Cardinality(), globalState)).Cardinality(), _958_i5, _958_i5, (_this).F23())
			var _962_v18 _dafny.Set
			_ = _962_v18
			_962_v18 = _dafny.SetOf((_960_v17).Update((_this).F23(), Companion_Default___.Abs((_this).F23())))
			var _963_v19 _dafny.Map
			_ = _963_v19
			_963_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(false, !(p1)))
			var _964_v20 _dafny.Array
			_ = _964_v20
			var _nwElement0_41 _dafny.Set = (_946_v6).Difference(_946_v6)
			_ = _nwElement0_41
			var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(20))
			_ = _nw179
			(_nw179).ArraySet1(_nwElement0_41, 0)
			(_nw179).ArraySet1(_946_v6, 1)
			(_nw179).ArraySet1(_946_v6, 2)
			(_nw179).ArraySet1(_dafny.SetOf(p1, p1, (_this).F25(), p1), 3)
			(_nw179).ArraySet1((_dafny.SetOf(_this.F24())).Difference(_946_v6), 4)
			(_nw179).ArraySet1((func() _dafny.Set {
				if _this.F24() {
					return _946_v6
				}
				return _946_v6
			})(), 5)
			(_nw179).ArraySet1((_946_v6).Union(_946_v6), 6)
			(_nw179).ArraySet1((_946_v6).Union(_dafny.SetOf(_this.F24())), 7)
			(_nw179).ArraySet1(_946_v6, 8)
			(_nw179).ArraySet1(_946_v6, 9)
			(_nw179).ArraySet1(_946_v6, 10)
			(_nw179).ArraySet1(_946_v6, 11)
			(_nw179).ArraySet1(_946_v6, 12)
			(_nw179).ArraySet1((_946_v6).Intersection(_946_v6), 13)
			(_nw179).ArraySet1(_946_v6, 14)
			(_nw179).ArraySet1((_946_v6).Union(_946_v6), 15)
			(_nw179).ArraySet1((_dafny.SetOf(p1, (_this).F25(), _this.F24(), _this.F24())).Intersection((func() _dafny.Set {
				if (_963_v19).Contains((_this).F25()) {
					return (_963_v19).Get((_this).F25()).(_dafny.Set)
				}
				return _946_v6
			})()), 16)
			(_nw179).ArraySet1(_dafny.SetOf(p0, p0, p0, _this.F24(), (_this).Fm3(_939_v2, (_this).F25(), _958_i5, _937_v0, globalState)), 17)
			(_nw179).ArraySet1(_946_v6, 18)
			(_nw179).ArraySet1(_946_v6, 19)
			_964_v20 = _nw179
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_964_v20), 0))
			_ = _index148
			(_964_v20).ArraySet1(_946_v6, (_index148).Int())
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_964_v20), 0))
			_ = _index149
			var _rhs175 _dafny.Set = _962_v18
			_ = _rhs175
			var _rhs176 _dafny.Int = Companion_Default___.SafeDivisionInt(_958_i5, (_dafny.Zero).Minus(Companion_Default___.Fm1((_dafny.Zero).Minus(_958_i5), p0, globalState)))
			_ = _rhs176
			var _rhs177 _dafny.Set = _dafny.SetOf(p1, (_958_i5).Cmp(_958_i5) < 0, !(false))
			_ = _rhs177
			var _lhs151 *GlobalState = globalState
			_ = _lhs151
			var _lhs152 _dafny.Array = _964_v20
			_ = _lhs152
			var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_964_v20), 0))
			_ = _lhs153
			_962_v18 = _rhs175
			_lhs151.F8 = _rhs176
			(_lhs152).ArraySet1(_rhs177, (_lhs153).Int())
			if p0 {
				var _965_v21 _dafny.CodePoint
				_ = _965_v21
				_965_v21 = _dafny.CodePoint('y')
				var _966_v22 _dafny.MultiSet
				_ = _966_v22
				_966_v22 = _dafny.MultiSetOf(_960_v17)
				var _967_v23 _dafny.Map
				_ = _967_v23
				_967_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.SeqOf(_965_v21)).Cardinality())).Times((_this).F23()), _966_v22)
				_967_v23 = (_967_v23).Update(((func() _dafny.Int {
					if (_960_v17).Contains(_dafny.IntOfUint32((_949_v9).Cardinality())) {
						return (_960_v17).Multiplicity(_dafny.IntOfUint32((_949_v9).Cardinality()))
					}
					return (_this).F23()
				})()).Plus(_dafny.IntOfInt64(895)), _966_v22)
				var _968_v24 _dafny.Set
				_ = _968_v24
				_968_v24 = _dafny.SetOf(_943_v4)
				var _969_v25 *C0
				_ = _969_v25
				var _nw180 *C0 = New_C0_()
				_ = _nw180
				_nw180.Ctor__((_968_v24).Difference(_968_v24), (_this).F23(), _958_i5)
				_969_v25 = _nw180
				(globalState).F17 = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(p0, _this.F24()), _dafny.Companion_Sequence_.Concatenate(_949_v9, _949_v9), _949_v9)).Cardinality())
				var _970_v26 _dafny.Array
				_ = _970_v26
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw181
				_970_v26 = _nw181
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_970_v26), 0))
				_ = _index150
				(_970_v26).ArraySet1(_950_v10, (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_970_v26), 0))
				_ = _index151
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_31
				var _nw182 _dafny.Array
				_ = _nw182
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw182 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Int = (func(_971_v25 *C0) func(_dafny.Int) _dafny.Int {
						return func(_972_i7 _dafny.Int) _dafny.Int {
							return (_972_i7).Times((_dafny.Zero).Minus(_971_v25.F31))
						}
					})(_969_v25)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw182 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw182).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw182).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				(_970_v26).ArraySet1(_nw182, (_index151).Int())
				_943_v4 = _943_v4
			} else {
				_939_v2 = _dafny.Companion_Sequence_.Concatenate(_939_v2, _939_v2)
				var _973_v27 _dafny.Set
				_ = _973_v27
				_973_v27 = _dafny.SetOf(_946_v6, _946_v6)
				(globalState).F9 = ((func() _dafny.Set {
					if p1 {
						return _dafny.SetOf((_964_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_964_v20), 0))).Int()).(_dafny.Set), (_964_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_964_v20), 0))).Int()).(_dafny.Set))
					}
					return Companion_Default___.Fm38((_this).F23(), (_dafny.Zero).Minus(_958_i5), globalState)
				})()).IsDisjointFrom(_973_v27)
				(globalState).F6 = (_949_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.IntOfUint32((_949_v9).Cardinality()))).Uint32()).(bool)
				(globalState).F15 = (_this).Fm3(_dafny.UnicodeSeqOfUtf8Bytes("gvwh"), p0, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_937_v0).Cardinality()), _958_i5), _937_v0, globalState)
				var _974_v28 _dafny.Map
				_ = _974_v28
				_974_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_949_v9, _958_i5)
				var _975_v29 *C3
				_ = _975_v29
				var _nw183 *C3 = New_C3_()
				_ = _nw183
				_nw183.Ctor__(p0, (_974_v28).Cardinality(), p0, (_958_i5).Cmp(_958_i5) > 0, _958_i5)
				_975_v29 = _nw183
			}
			(globalState).F14 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_949_v9, (func() _dafny.Sequence {
				if _this.F24() {
					return _949_v9
				}
				return _dafny.SeqOf(_this.F24(), _this.F24())
			})())).Cardinality())
		}
		r0 = (_this).F23()
		r1 = _956_v15
		r2 = (_this).F25()
		return r0, r1, r2
	}
}
func (_this *C6) M4(p0 _dafny.Array, p1 _dafny.Int, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _976_v1 _dafny.Sequence
		_ = _976_v1
		_976_v1 = _dafny.SeqOf(p3, _dafny.CodePoint('u'))
		var _rhs178 _dafny.Int = (_this).Fm2(globalState)
		_ = _rhs178
		var _rhs179 bool = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality())).Cmp(((func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-489), _dafny.IntOfInt64(-153))); ; {
				_compr_34, _ok39 := _iter39()
				if !_ok39 {
					break
				}
				var _977_v0 _dafny.Int
				_977_v0 = interface{}(_compr_34).(_dafny.Int)
				if ((_dafny.IntOfInt64(-489)).Cmp(_977_v0) <= 0) && ((_977_v0).Cmp(_dafny.IntOfInt64(-153)) < 0) {
					_coll34.Add((_977_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F23(), p1, p1)).Cardinality())), (_this).F25())
				}
			}
			return _coll34.ToMap()
		}()).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_976_v1, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_976_v1).Cardinality()))).Uint32(), _dafny.CodePoint('q'))).Cardinality()))) <= 0
		_ = _rhs179
		var _lhs154 *GlobalState = globalState
		_ = _lhs154
		var _lhs155 *GlobalState = globalState
		_ = _lhs155
		_lhs154.F8 = _rhs178
		_lhs155.F9 = _rhs179
		var _978_v2 _dafny.Map
		_ = _978_v2
		_978_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bhkyjg")).Cardinality()), (_this).F23()), (_this).F23())
		var _979_v3 _dafny.MultiSet
		_ = _979_v3
		_979_v3 = _dafny.MultiSetOf((_this).F23(), p1)
		var _980_v4 _dafny.Map
		_ = _980_v4
		_980_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F23())
		var _981_v5 _dafny.Array
		_ = _981_v5
		var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw184
		_981_v5 = _nw184
		var _982_v6 D7
		_ = _982_v6
		_982_v6 = Companion_D7_.Create_DC21_(Companion_D7_.Create_DC20_(_this.F24(), _980_v4, _981_v5))
		var _983_v7 D7
		_ = _983_v7
		_983_v7 = Companion_D7_.Create_DC21_(_982_v6)
		var _984_v8 _dafny.Sequence
		_ = _984_v8
		_984_v8 = _dafny.SeqOf(p1)
		var _985_v9 _dafny.Map
		_ = _985_v9
		_985_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_983_v7, (_984_v8).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_984_v8).Cardinality()))).Uint32()).(_dafny.Int))
		var _986_v10 _dafny.Sequence
		_ = _986_v10
		_986_v10 = _dafny.SeqOf((func() _dafny.Int {
			if (_985_v9).Contains(_983_v7) {
				return (_985_v9).Get(_983_v7).(_dafny.Int)
			}
			return (_this).F23()
		})())
		(globalState).F1 = (_this).Fm3(_dafny.Companion_Sequence_.Concatenate(_976_v1, _dafny.Companion_Sequence_.Update(_976_v1, (Companion_Default___.SafeIndex((_this).Fm5(p1, (_978_v2).Cardinality(), globalState), _dafny.IntOfUint32((_976_v1).Cardinality()))).Uint32(), p3)), (_this).F25(), (_979_v3).Cardinality(), _986_v10, globalState)
		if (_dafny.IntOfInt64(754)).Cmp((_this).F23()) == 0 {
			var _987_v11 *C5
			_ = _987_v11
			var _nw185 *C5 = New_C5_()
			_ = _nw185
			_nw185.Ctor__((_this).F25(), (_this).F25())
			_987_v11 = _nw185
			_987_v11 = _987_v11
			var _988_v12 _dafny.Set
			_ = _988_v12
			_988_v12 = _dafny.SetOf((_this).F25())
			var _source8 D3 = Companion_Default___.Fm39(p2, (p1).Minus((_this).F23()), true, (_dafny.SetOf(p2, _this.F24(), (_this).F25(), p2, true)).Difference(_988_v12), globalState)
			_ = _source8
			if _source8.Is_DC9() {
				var _989___mcc_h0 _dafny.Int = _source8.Get_().(D3_DC9).Cf15
				_ = _989___mcc_h0
				var _990_cf15 _dafny.Int = _989___mcc_h0
				_ = _990_cf15
				(globalState).F8 = (((_dafny.MultiSetOf((_this).F23(), _990_cf15)).Intersection(_979_v3)).Cardinality()).Times(_990_cf15)
				var _991_v13 _dafny.Map
				_ = _991_v13
				_991_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), (_dafny.Zero).Minus((_this).F23()))
				var _992_v14 _dafny.Array
				_ = _992_v14
				var _nwElement0_42 _dafny.CodePoint = p3
				_ = _nwElement0_42
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(28))
				_ = _nw186
				(_nw186).ArraySet1CodePoint(_nwElement0_42, 0)
				(_nw186).ArraySet1CodePoint(p3, 1)
				(_nw186).ArraySet1CodePoint(p3, 2)
				(_nw186).ArraySet1CodePoint(p3, 3)
				(_nw186).ArraySet1CodePoint(p3, 4)
				(_nw186).ArraySet1CodePoint(p3, 5)
				(_nw186).ArraySet1CodePoint(p3, 6)
				(_nw186).ArraySet1CodePoint(p3, 7)
				(_nw186).ArraySet1CodePoint(p3, 8)
				(_nw186).ArraySet1CodePoint(p3, 9)
				(_nw186).ArraySet1CodePoint(p3, 10)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('t'), 11)
				(_nw186).ArraySet1CodePoint(p3, 12)
				(_nw186).ArraySet1CodePoint(p3, 13)
				(_nw186).ArraySet1CodePoint(p3, 14)
				(_nw186).ArraySet1CodePoint(p3, 15)
				(_nw186).ArraySet1CodePoint(p3, 16)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('n'), 17)
				(_nw186).ArraySet1CodePoint(p3, 18)
				(_nw186).ArraySet1CodePoint(p3, 19)
				(_nw186).ArraySet1CodePoint(p3, 20)
				(_nw186).ArraySet1CodePoint(p3, 21)
				(_nw186).ArraySet1CodePoint(p3, 22)
				(_nw186).ArraySet1CodePoint(p3, 23)
				(_nw186).ArraySet1CodePoint(p3, 24)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('g'), 25)
				(_nw186).ArraySet1CodePoint(p3, 26)
				(_nw186).ArraySet1CodePoint(p3, 27)
				_992_v14 = _nw186
				var _993_v15 _dafny.Set
				_ = _993_v15
				_993_v15 = _dafny.SetOf(_992_v14)
				var _994_v16 D5
				_ = _994_v16
				_994_v16 = Companion_D5_.Create_DC14_((_this).F23(), p2, _993_v15)
				var _995_v17 _dafny.Map
				_ = _995_v17
				_995_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)
				var _996_v18 _dafny.Array
				_ = _996_v18
				var _nwElement0_43 bool = _this.F24()
				_ = _nwElement0_43
				var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(16))
				_ = _nw187
				(_nw187).ArraySet1(_nwElement0_43, 0)
				(_nw187).ArraySet1((_994_v16).Dtor_cf23(), 1)
				(_nw187).ArraySet1((_this).F25(), 2)
				(_nw187).ArraySet1((_this).Fm3(_976_v1, _this.F24(), (_980_v4).Cardinality(), _984_v8, globalState), 3)
				(_nw187).ArraySet1(false, 4)
				(_nw187).ArraySet1((func() bool {
					if (_995_v17).Contains((_this).F25()) {
						return (_995_v17).Get((_this).F25()).(bool)
					}
					return _this.F24()
				})(), 5)
				(_nw187).ArraySet1(true, 6)
				(_nw187).ArraySet1(p2, 7)
				(_nw187).ArraySet1((_this).F25(), 8)
				(_nw187).ArraySet1(p2, 9)
				(_nw187).ArraySet1(_this.F24(), 10)
				(_nw187).ArraySet1(_this.F24(), 11)
				(_nw187).ArraySet1((_this).Fm3(_976_v1, _this.F24(), (_this).F23(), _986_v10, globalState), 12)
				(_nw187).ArraySet1(_this.F24(), 13)
				(_nw187).ArraySet1((_this).F25(), 14)
				(_nw187).ArraySet1((_this).F25(), 15)
				_996_v18 = _nw187
				var _997_v19 _dafny.Array
				_ = _997_v19
				var _nwElement0_44 _dafny.Int = (_this).F23()
				_ = _nwElement0_44
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(26))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_44, 0)
				(_nw188).ArraySet1((_dafny.MultiSetOf(p1)).Cardinality(), 1)
				(_nw188).ArraySet1(_990_cf15, 2)
				(_nw188).ArraySet1(_990_cf15, 3)
				(_nw188).ArraySet1(p1, 4)
				(_nw188).ArraySet1(_dafny.IntOfInt64(283), 5)
				(_nw188).ArraySet1((_this).F23(), 6)
				(_nw188).ArraySet1(p1, 7)
				(_nw188).ArraySet1((_this).F23(), 8)
				(_nw188).ArraySet1(_990_cf15, 9)
				(_nw188).ArraySet1((_this).F23(), 10)
				(_nw188).ArraySet1(_dafny.IntOfInt64(406), 11)
				(_nw188).ArraySet1((_this).F23(), 12)
				(_nw188).ArraySet1(p1, 13)
				(_nw188).ArraySet1((_991_v13).Cardinality(), 14)
				(_nw188).ArraySet1(_990_cf15, 15)
				(_nw188).ArraySet1(p1, 16)
				(_nw188).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), p2)).Cardinality(), 17)
				(_nw188).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("huhwbkh")).Cardinality()), 18)
				(_nw188).ArraySet1(_dafny.IntOfUint32((_976_v1).Cardinality()), 19)
				(_nw188).ArraySet1(_990_cf15, 20)
				(_nw188).ArraySet1(_dafny.IntOfUint32((_986_v10).Cardinality()), 21)
				(_nw188).ArraySet1((_this).F23(), 22)
				(_nw188).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_996_v18)).Cardinality()), 23)
				(_nw188).ArraySet1((_this).F23(), 24)
				(_nw188).ArraySet1(p1, 25)
				_997_v19 = _nw188
				var _998_v20 _dafny.MultiSet
				_ = _998_v20
				_998_v20 = _dafny.MultiSetOf(_997_v19)
				var _999_v21 _dafny.Map
				_ = _999_v21
				_999_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v20, (_987_v11).Fm5((_this).F23(), _990_cf15, globalState))
				_999_v21 = (_999_v21).Update((Companion_D12_.Create_DC32_(_998_v20)).Dtor_cf60(), ((_this).F23()).Plus(_990_cf15))
				(globalState).F9 = true
				(globalState).F1 = (_990_cf15).Cmp((_this).F23()) != 0
			} else if _source8.Is_DC10() {
				var _1000___mcc_h1 bool = _source8.Get_().(D3_DC10).Cf16
				_ = _1000___mcc_h1
				var _1001___mcc_h2 _dafny.Map = _source8.Get_().(D3_DC10).Cf17
				_ = _1001___mcc_h2
				var _1002_cf17 _dafny.Map = _1001___mcc_h2
				_ = _1002_cf17
				var _1003_cf16 bool = _1000___mcc_h1
				_ = _1003_cf16
				(globalState).F8 = Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23())
				(globalState).F14 = p1
				var _rhs180 _dafny.Int = p1
				_ = _rhs180
				var _rhs181 _dafny.Int = (_this).Fm2(globalState)
				_ = _rhs181
				var _rhs182 bool = (false) || (_this.F24())
				_ = _rhs182
				var _rhs183 _dafny.Int = p1
				_ = _rhs183
				var _rhs184 bool = p2
				_ = _rhs184
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *C6 = _this
				_ = _lhs158
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				_lhs156.F5 = _rhs180
				_lhs157.F5 = _rhs181
				_lhs158.F24_set_(_rhs182)
				_lhs159.F17 = _rhs183
				_lhs160.F15 = _rhs184
				var _1004_v22 *C1
				_ = _1004_v22
				var _nw189 *C1 = New_C1_()
				_ = _nw189
				_nw189.Ctor__(p2, false)
				_1004_v22 = _nw189
			} else {
				var _1005___mcc_h3 _dafny.Sequence = _source8.Get_().(D3_DC8).Cf14
				_ = _1005___mcc_h3
				var _1006_cf14 _dafny.Sequence = _1005___mcc_h3
				_ = _1006_cf14
				var _1007_v23 _dafny.Array
				_ = _1007_v23
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_32
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.MultiSet = func(_1008_i0 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F24()))).Intersection(_dafny.MultiSetOf((_this).F25()))
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw190 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw190).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw190).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1007_v23 = _nw190
				var _1009_v24 _dafny.MultiSet
				_ = _1009_v24
				_1009_v24 = _dafny.MultiSetOf(p2)
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1007_v23), 0))
				_ = _index152
				(_1007_v23).ArraySet1(_1009_v24, (_index152).Int())
				var _1010_v25 D7
				_ = _1010_v25
				_1010_v25 = Companion_D7_.Create_DC19_(_1009_v24)
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1007_v23), 0))
				_ = _index153
				(_1007_v23).ArraySet1((_1010_v25).Dtor_cf34(), (_index153).Int())
				(globalState).F8 = (_dafny.MultiSetOf(p2, p2)).Cardinality()
				var _1011_v26 _dafny.Sequence
				_ = _1011_v26
				_1011_v26 = _dafny.SeqOf(_981_v5, _981_v5, _981_v5, _981_v5)
				_981_v5 = (_1011_v26).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1011_v26).Cardinality()))).Uint32()).(_dafny.Array)
				(globalState).F17 = p1
			}
			var _1012_v27 _dafny.Sequence
			_ = _1012_v27
			_1012_v27 = _dafny.SeqOf(_979_v3)
			var _rhs185 _dafny.MultiSet = ((_1012_v27).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1012_v27).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference(_979_v3)
			_ = _rhs185
			var _rhs186 bool = p2
			_ = _rhs186
			var _rhs187 bool = p2
			_ = _rhs187
			var _lhs161 *GlobalState = globalState
			_ = _lhs161
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			_979_v3 = _rhs185
			_lhs161.F1 = _rhs186
			_lhs162.F1 = _rhs187
			(globalState).F17 = (_this).F23()
			var _1013_v28 _dafny.CodePoint
			_ = _1013_v28
			_1013_v28 = _dafny.CodePoint('m')
			_1013_v28 = _1013_v28
		} else {
			(globalState).F15 = p2
			(globalState).F17 = (_dafny.Zero).Minus(p1)
			(globalState).F8 = (_this).F23()
			(globalState).F8 = _dafny.IntOfInt64(129)
			(globalState).F8 = ((_dafny.MultiSetOf((_this).F25(), (_this).F25(), !(true))).Update((_this).F25(), Companion_Default___.Abs(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23())))).Cardinality()
		}
		(globalState).F15 = _this.F24()
		(globalState).F9 = (func() bool {
			if false {
				return _this.F24()
			}
			return (func() bool {
				if (_this).F25() {
					return (_this).Fm3(_976_v1, p2, (_this).F23(), _984_v8, globalState)
				}
				return p2
			})()
		})()
		var _1014_v29 _dafny.Array
		_ = _1014_v29
		var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(16))
		_ = _nw191
		_1014_v29 = _nw191
		var _1015_v30 _dafny.Set
		_ = _1015_v30
		_1015_v30 = _dafny.SetOf(_1014_v29, _1014_v29)
		var _1016_v31 D5
		_ = _1016_v31
		_1016_v31 = Companion_D5_.Create_DC14_((_this).F23(), _this.F24(), _1015_v30)
		var _1017_v32 _dafny.MultiSet
		_ = _1017_v32
		_1017_v32 = _dafny.MultiSetOf(_1016_v31, _1016_v31, _1016_v31, _1016_v31, _1016_v31)
		(globalState).F14 = (func() _dafny.Int {
			if (_this).F25() {
				return (_1017_v32).Cardinality()
			}
			return p1
		})()
		var _1018_v33 _dafny.Array
		_ = _1018_v33
		var _nwElement0_45 bool = _this.F24()
		_ = _nwElement0_45
		var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(2))
		_ = _nw192
		(_nw192).ArraySet1(_nwElement0_45, 0)
		(_nw192).ArraySet1(_this.F24(), 1)
		_1018_v33 = _nw192
		r0 = _1018_v33
		return r0
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) Ctor__() {
	{
	}
}
func (_this *C7) Fm7(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetFromSeq(_dafny.SeqOf(false, !(true), false, !(false), !(true))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(false, false, true, true)))
	}
}
func (_this *C7) M2(p0 _dafny.Int, p1 bool, globalState *GlobalState) (bool, bool, D0) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 D0 = Companion_D0_.Default()
		_ = r2
		r0 = p1
		var _1019_v0 _dafny.Sequence
		_ = _1019_v0
		_1019_v0 = _dafny.UnicodeSeqOfUtf8Bytes("hrd")
		var _1020_v1 D0
		_ = _1020_v1
		_1020_v1 = Companion_D0_.Create_DC1_(p1, p0)
		var _pat_let_tv15 = _1020_v1
		_ = _pat_let_tv15
		var _pat_let_tv16 = p0
		_ = _pat_let_tv16
		var _pat_let_tv17 = p0
		_ = _pat_let_tv17
		var _rhs188 _dafny.Int = (_dafny.IntOfUint32((_1019_v0).Cardinality())).Minus(p0)
		_ = _rhs188
		var _rhs189 _dafny.Int = (_dafny.Zero).Minus(func(_source9 D0) _dafny.Int {
			if _source9.Is_DC1() {
				var _1021___mcc_h0 bool = _source9.Get_().(D0_DC1).Cf1
				_ = _1021___mcc_h0
				var _1022___mcc_h1 _dafny.Int = _source9.Get_().(D0_DC1).Cf2
				_ = _1022___mcc_h1
				var _1023_cf2 _dafny.Int = _1022___mcc_h1
				_ = _1023_cf2
				var _1024_cf1 bool = _1021___mcc_h0
				_ = _1024_cf1
				return (_1023_cf2).Minus(_1023_cf2)
			} else if _source9.Is_DC0() {
				var _1025___mcc_h2 bool = _source9.Get_().(D0_DC0).Cf0
				_ = _1025___mcc_h2
				var _1026_cf0 bool = _1025___mcc_h2
				_ = _1026_cf0
				return (_pat_let_tv15).Dtor_cf2()
			} else {
				var _1027___mcc_h3 D0 = _source9.Get_().(D0_DC2).Cf3
				_ = _1027___mcc_h3
				var _1028_cf3 D0 = _1027___mcc_h3
				_ = _1028_cf3
				return Companion_Default___.SafeDivisionInt(_pat_let_tv16, _pat_let_tv17)
			}
		}(_1020_v1))
		_ = _rhs189
		var _lhs163 *GlobalState = globalState
		_ = _lhs163
		var _lhs164 *GlobalState = globalState
		_ = _lhs164
		_lhs163.F8 = _rhs188
		_lhs164.F8 = _rhs189
		var _1029_v2 _dafny.Array
		_ = _1029_v2
		var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw193
		_1029_v2 = _nw193
		var _1030_v3 _dafny.Array
		_ = _1030_v3
		var _nw194 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw194
		_1030_v3 = _nw194
		var _1031_v4 _dafny.Sequence
		_ = _1031_v4
		_1031_v4 = _dafny.SeqOf(_1030_v3, _1030_v3, _1030_v3)
		var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_1029_v2), 0))
		_ = _index154
		(_1029_v2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1031_v4, _1031_v4), (_index154).Int())
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_1029_v2), 0))
		_ = _index155
		(_1029_v2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1031_v4, _1031_v4), (_index155).Int())
		var _1032_v5 _dafny.Sequence
		_ = _1032_v5
		_1032_v5 = _dafny.SeqOf(p0)
		(globalState).F8 = Companion_Default___.Fm1((Companion_Default___.Fm8(_dafny.IntOfUint32((_1032_v5).Cardinality()), globalState)).Cardinality(), (_dafny.IntOfUint32((_1019_v0).Cardinality())).Cmp(_dafny.IntOfInt64(639)) == 0, globalState)
		var _1033_v7 _dafny.Array
		_ = _1033_v7
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_33
		var _nw195 _dafny.Array
		_ = _nw195
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw195 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.Map = (func(_1034_v5 _dafny.Sequence, _1035_p1 bool) func(_dafny.Int) _dafny.Map {
				return func(_1036_i0 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1034_v5, _dafny.SeqOf(_1035_p1, _1035_p1))).Merge(func() _dafny.Map {
						var _coll35 = _dafny.NewMapBuilder()
						_ = _coll35
						for _iter40 := _dafny.Iterate((_dafny.MultiSetOf(_1034_v5, _1034_v5)).Elements()); ; {
							_compr_35, _ok40 := _iter40()
							if !_ok40 {
								break
							}
							var _1037_v6 _dafny.Sequence
							_1037_v6 = interface{}(_compr_35).(_dafny.Sequence)
							if (_dafny.MultiSetOf(_1034_v5, _1034_v5)).Contains(_1037_v6) {
								_coll35.Add(_1037_v6, _dafny.SeqOf(_1035_p1, _1035_p1))
							}
						}
						return _coll35.ToMap()
					}())
				}
			})(_1032_v5, p1)
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw195 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw195).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw195).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1033_v7 = _nw195
		var _1038_v8 _dafny.Sequence
		_ = _1038_v8
		_1038_v8 = _dafny.SeqOf(p1)
		var _1039_v9 _dafny.Map
		_ = _1039_v9
		_1039_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0, p0, p0), _1038_v8)
		var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_1033_v7), 0))
		_ = _index156
		(_1033_v7).ArraySet1((_1039_v9).Merge(_1039_v9), (_index156).Int())
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_1033_v7), 0))
		_ = _index157
		(_1033_v7).ArraySet1(_1039_v9, (_index157).Int())
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1030_v3), 0))
		_ = _index158
		(_1030_v3).ArraySet1(p1, (_index158).Int())
		var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1030_v3), 0))
		_ = _index159
		(_1030_v3).ArraySet1((func() bool {
			if false {
				return p1
			}
			return !(Companion_Default___.Fm9(p1, globalState)) || (Companion_Default___.Fm9(true, globalState))
		})(), (_index159).Int())
		r0 = p1
		r1 = Companion_Default___.Fm9(p1, globalState)
		r2 = _1020_v1
		return r0, r1, r2
	}
}
func (_this *C7) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Array, globalState *GlobalState) (_dafny.Map, _dafny.Array) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1040_v0 bool
		_ = _1040_v0
		_1040_v0 = false
		(globalState).F15 = _1040_v0
		var _1041_v1 _dafny.CodePoint
		_ = _1041_v1
		_1041_v1 = _dafny.CodePoint('l')
		var _1042_v2 _dafny.Map
		_ = _1042_v2
		_1042_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _dafny.CodePoint('i'))
		var _1043_v3 D1
		_ = _1043_v3
		_1043_v3 = Companion_D1_.Create_DC3_(_1041_v1)
		var _1044_v4 _dafny.Array
		_ = _1044_v4
		var _nwElement0_46 _dafny.CodePoint = _1041_v1
		_ = _nwElement0_46
		var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(14))
		_ = _nw196
		(_nw196).ArraySet1CodePoint(_nwElement0_46, 0)
		(_nw196).ArraySet1CodePoint(_1041_v1, 1)
		(_nw196).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_1042_v2).Contains(_1040_v0) {
				return (_1042_v2).Get(_1040_v0).(_dafny.CodePoint)
			}
			return _1041_v1
		})(), 2)
		(_nw196).ArraySet1CodePoint(_1041_v1, 3)
		(_nw196).ArraySet1CodePoint(_1041_v1, 4)
		(_nw196).ArraySet1CodePoint(_1041_v1, 5)
		(_nw196).ArraySet1CodePoint(_1041_v1, 6)
		(_nw196).ArraySet1CodePoint(_dafny.CodePoint('o'), 7)
		(_nw196).ArraySet1CodePoint((_1043_v3).Dtor_cf4(), 8)
		(_nw196).ArraySet1CodePoint(_1041_v1, 9)
		(_nw196).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _1040_v0 {
				return _1041_v1
			}
			return _1041_v1
		})(), 10)
		(_nw196).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_1042_v2).Contains(_1040_v0) {
				return (_1042_v2).Get(_1040_v0).(_dafny.CodePoint)
			}
			return _1041_v1
		})(), 11)
		(_nw196).ArraySet1CodePoint(_1041_v1, 12)
		(_nw196).ArraySet1CodePoint(_dafny.CodePoint('g'), 13)
		_1044_v4 = _nw196
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))
		_ = _index160
		(_1044_v4).ArraySet1CodePoint(_1041_v1, (_index160).Int())
		var _1045_v5 _dafny.Map
		_ = _1045_v5
		_1045_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1040_v0)
		var _pat_let_tv18 = _1040_v0
		_ = _pat_let_tv18
		var _pat_let_tv19 = _1040_v0
		_ = _pat_let_tv19
		var _pat_let_tv20 = _1040_v0
		_ = _pat_let_tv20
		var _pat_let_tv21 = _1040_v0
		_ = _pat_let_tv21
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))
		_ = _index161
		var _rhs190 _dafny.CodePoint = _1041_v1
		_ = _rhs190
		var _rhs191 bool = func(_source10 D1) bool {
			if _source10.Is_DC4() {
				var _1046___mcc_h0 bool = _source10.Get_().(D1_DC4).Cf5
				_ = _1046___mcc_h0
				var _1047___mcc_h1 bool = _source10.Get_().(D1_DC4).Cf6
				_ = _1047___mcc_h1
				var _1048_cf6 bool = _1047___mcc_h1
				_ = _1048_cf6
				var _1049_cf5 bool = _1046___mcc_h0
				_ = _1049_cf5
				return _pat_let_tv18
			} else if _source10.Is_DC5() {
				var _1050___mcc_h2 bool = _source10.Get_().(D1_DC5).Cf7
				_ = _1050___mcc_h2
				var _1051___mcc_h3 _dafny.Map = _source10.Get_().(D1_DC5).Cf8
				_ = _1051___mcc_h3
				var _1052___mcc_h4 _dafny.Int = _source10.Get_().(D1_DC5).Cf9
				_ = _1052___mcc_h4
				var _1053___mcc_h5 _dafny.Int = _source10.Get_().(D1_DC5).Cf10
				_ = _1053___mcc_h5
				var _1054___mcc_h6 _dafny.Map = _source10.Get_().(D1_DC5).Cf11
				_ = _1054___mcc_h6
				var _1055_cf11 _dafny.Map = _1054___mcc_h6
				_ = _1055_cf11
				var _1056_cf10 _dafny.Int = _1053___mcc_h5
				_ = _1056_cf10
				var _1057_cf9 _dafny.Int = _1052___mcc_h4
				_ = _1057_cf9
				var _1058_cf8 _dafny.Map = _1051___mcc_h3
				_ = _1058_cf8
				var _1059_cf7 bool = _1050___mcc_h2
				_ = _1059_cf7
				return _pat_let_tv19
			} else if _source10.Is_DC3() {
				var _1060___mcc_h7 _dafny.CodePoint = _source10.Get_().(D1_DC3).Cf4
				_ = _1060___mcc_h7
				var _1061_cf4 _dafny.CodePoint = _1060___mcc_h7
				_ = _1061_cf4
				return _pat_let_tv20
			} else {
				var _1062___mcc_h8 D1 = _source10.Get_().(D1_DC6).Cf12
				_ = _1062___mcc_h8
				var _1063_cf12 D1 = _1062___mcc_h8
				_ = _1063_cf12
				return _pat_let_tv21
			}
		}(Companion_D1_.Create_DC3_(_1041_v1))
		_ = _rhs191
		var _rhs192 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1040_v0)).Merge(_1045_v5)).Cardinality()
		_ = _rhs192
		var _rhs193 bool = false
		_ = _rhs193
		var _lhs165 _dafny.Array = _1044_v4
		_ = _lhs165
		var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))
		_ = _lhs166
		var _lhs167 *GlobalState = globalState
		_ = _lhs167
		var _lhs168 *GlobalState = globalState
		_ = _lhs168
		var _lhs169 *GlobalState = globalState
		_ = _lhs169
		(_lhs165).ArraySet1CodePoint(_rhs190, (_lhs166).Int())
		_lhs167.F15 = _rhs191
		_lhs168.F8 = _rhs192
		_lhs169.F1 = _rhs193
		var _1064_i0 _dafny.Int
		_ = _1064_i0
		_1064_i0 = _dafny.Zero
		{
			for _1040_v0 {
				{
					if (_1064_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1064_i0 = (_1064_i0).Plus(_dafny.One)
					var _1065_v6 _dafny.Sequence
					_ = _1065_v6
					_1065_v6 = _dafny.UnicodeSeqOfUtf8Bytes("asg")
					_1065_v6 = _1065_v6
					(globalState).F8 = Companion_Default___.SafeDivisionInt(p2, p0)
					var _1066_v7 _dafny.Map
					_ = _1066_v7
					_1066_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, p0)
					var _1067_v8 D1
					_ = _1067_v8
					_1067_v8 = Companion_D1_.Create_DC5_(_1040_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1040_v0)).Merge(_1045_v5), (p0).Plus(p1), _dafny.IntOfInt64(357), _1066_v7)
					var _source11 D1 = _1067_v8
					_ = _source11
					if _source11.Is_DC4() {
						var _1068___mcc_h9 bool = _source11.Get_().(D1_DC4).Cf5
						_ = _1068___mcc_h9
						var _1069___mcc_h10 bool = _source11.Get_().(D1_DC4).Cf6
						_ = _1069___mcc_h10
						var _1070_cf6 bool = _1069___mcc_h10
						_ = _1070_cf6
						var _1071_cf5 bool = _1068___mcc_h9
						_ = _1071_cf5
						var _1072_v9 _dafny.Sequence
						_ = _1072_v9
						_1072_v9 = _dafny.SeqOf(_1071_cf5)
						var _1073_v10 _dafny.Sequence
						_ = _1073_v10
						_1073_v10 = _dafny.SeqOf(Companion_Default___.Fm1(p0, _1040_v0, globalState), p0)
						var _1074_v11 _dafny.Sequence
						_ = _1074_v11
						_1074_v11 = _dafny.SeqOf(_dafny.IntOfUint32((_1073_v10).Cardinality()), p1, _dafny.IntOfUint32((_1072_v9).Cardinality()), p1)
						var _1075_v12 _dafny.Sequence
						_ = _1075_v12
						_1075_v12 = _dafny.SeqOf((_1072_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1074_v11).Cardinality()), _dafny.IntOfUint32((_1072_v9).Cardinality()))).Uint32()).(bool), _1071_cf5)
						_1075_v12 = _1075_v12
						_1041_v1 = (func() _dafny.CodePoint {
							if !((_1071_cf5) && (_1040_v0)) {
								return _1041_v1
							}
							return _dafny.CodePoint('b')
						})()
						var _1076_v13 _dafny.Map
						_ = _1076_v13
						_1076_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
						_1076_v13 = (_1076_v13).Update(_dafny.IntOfUint32((_1065_v6).Cardinality()), p1)
						(globalState).F14 = (func() _dafny.Int {
							if _1070_cf6 {
								return (Companion_Default___.Fm10(_1040_v0, _dafny.IntOfUint32((_1075_v12).Cardinality()), _1041_v1, (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()), globalState)).Cardinality()
							}
							return p1
						})()
					} else if _source11.Is_DC5() {
						var _1077___mcc_h11 bool = _source11.Get_().(D1_DC5).Cf7
						_ = _1077___mcc_h11
						var _1078___mcc_h12 _dafny.Map = _source11.Get_().(D1_DC5).Cf8
						_ = _1078___mcc_h12
						var _1079___mcc_h13 _dafny.Int = _source11.Get_().(D1_DC5).Cf9
						_ = _1079___mcc_h13
						var _1080___mcc_h14 _dafny.Int = _source11.Get_().(D1_DC5).Cf10
						_ = _1080___mcc_h14
						var _1081___mcc_h15 _dafny.Map = _source11.Get_().(D1_DC5).Cf11
						_ = _1081___mcc_h15
						var _1082_cf11 _dafny.Map = _1081___mcc_h15
						_ = _1082_cf11
						var _1083_cf10 _dafny.Int = _1080___mcc_h14
						_ = _1083_cf10
						var _1084_cf9 _dafny.Int = _1079___mcc_h13
						_ = _1084_cf9
						var _1085_cf8 _dafny.Map = _1078___mcc_h12
						_ = _1085_cf8
						var _1086_cf7 bool = _1077___mcc_h11
						_ = _1086_cf7
						var _1087_v14 _dafny.Array
						_ = _1087_v14
						var _len0_34 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_34
						var _nw197 _dafny.Array
						_ = _nw197
						if _len0_34.Cmp(_dafny.Zero) == 0 {
							_nw197 = _dafny.NewArray(_len0_34)
						} else {
							var _init34 func(_dafny.Int) bool = (func(_1088_v0 bool) func(_dafny.Int) bool {
								return func(_1089_i1 _dafny.Int) bool {
									return _1088_v0
								}
							})(_1040_v0)
							_ = _init34
							var _element0_34 = _init34(_dafny.Zero)
							_ = _element0_34
							_nw197 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
							(_nw197).ArraySet1(_element0_34, 0)
							var _nativeLen0_34 = (_len0_34).Int()
							_ = _nativeLen0_34
							for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
								(_nw197).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
							}
						}
						_1087_v14 = _nw197
						var _1090_v15 _dafny.Set
						_ = _1090_v15
						_1090_v15 = _dafny.SetOf(_1087_v14)
						var _1091_v16 _dafny.Set
						_ = _1091_v16
						_1091_v16 = _1090_v15
						var _1092_v17 *C0
						_ = _1092_v17
						var _nw198 *C0 = New_C0_()
						_ = _nw198
						_nw198.Ctor__((_1090_v15).Intersection((_1091_v16)), _dafny.IntOfInt64(675), (_dafny.Zero).Minus(_1084_cf9))
						_1092_v17 = _nw198
						var _1093_v18 _dafny.Set
						_ = _1093_v18
						_1093_v18 = _dafny.SetOf(_1084_cf9)
						var _1094_v19 _dafny.Sequence
						_ = _1094_v19
						_1094_v19 = _dafny.SeqOf(_1093_v18)
						var _rhs194 _dafny.Int = _dafny.IntOfUint32((_1094_v19).Cardinality())
						_ = _rhs194
						var _rhs195 bool = _1040_v0
						_ = _rhs195
						var _rhs196 bool = !(_1086_cf7)
						_ = _rhs196
						var _lhs170 *GlobalState = globalState
						_ = _lhs170
						var _lhs171 *GlobalState = globalState
						_ = _lhs171
						var _lhs172 *GlobalState = globalState
						_ = _lhs172
						_lhs170.F8 = _rhs194
						_lhs171.F1 = _rhs195
						_lhs172.F6 = _rhs196
						var _rhs197 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(-321)).Plus(p2), Companion_Default___.SafeModuloInt(_1084_cf9, _dafny.IntOfInt64(319)))
						_ = _rhs197
						var _rhs198 _dafny.CodePoint = _1041_v1
						_ = _rhs198
						var _lhs173 *GlobalState = globalState
						_ = _lhs173
						_lhs173.F14 = _rhs197
						_1041_v1 = _rhs198
						var _1095_v20 _dafny.Sequence
						_ = _1095_v20
						_1095_v20 = _dafny.SeqOf(_1084_cf9)
						var _1096_v21 *C5
						_ = _1096_v21
						var _nw199 *C5 = New_C5_()
						_ = _nw199
						_nw199.Ctor__((_1092_v17).Fm3(_1065_v6, _1040_v0, p1, _1095_v20, globalState), !(_1040_v0))
						_1096_v21 = _nw199
						var _1097_v22 _dafny.MultiSet
						_ = _1097_v22
						_1097_v22 = _dafny.MultiSetOf(_1096_v21)
						var _1098_v23 _dafny.Map
						_ = _1098_v23
						_1098_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((p1).Plus(_1083_cf10)), ((_1097_v22).Difference(_1097_v22)).Cardinality())
						var _1099_v25 _dafny.MultiSet
						_ = _1099_v25
						_1099_v25 = _dafny.MultiSetOf(p2, _1092_v17.F31)
						_1098_v23 = func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter41 := _dafny.Iterate(((_1099_v25).Update((_1096_v21).Fm13(_1083_cf10, _1040_v0, globalState), Companion_Default___.Abs(_dafny.IntOfUint32((_1065_v6).Cardinality())))).Elements()); ; {
								_compr_36, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _1100_v24 _dafny.Int
								_1100_v24 = interface{}(_compr_36).(_dafny.Int)
								if ((_1099_v25).Update((_1096_v21).Fm13(_1083_cf10, _1040_v0, globalState), Companion_Default___.Abs(_dafny.IntOfUint32((_1065_v6).Cardinality())))).Contains(_1100_v24) {
									_coll36.Add((_1100_v24).Plus(_1084_cf9), Companion_Default___.SafeModuloInt(p0, p0))
								}
							}
							return _coll36.ToMap()
						}()
					} else if _source11.Is_DC3() {
						var _1101___mcc_h16 _dafny.CodePoint = _source11.Get_().(D1_DC3).Cf4
						_ = _1101___mcc_h16
						var _1102_cf4 _dafny.CodePoint = _1101___mcc_h16
						_ = _1102_cf4
						(globalState).F17 = Companion_Default___.Fm1(p1, true, globalState)
						var _1103_v26 _dafny.Set
						_ = _1103_v26
						_1103_v26 = _dafny.SetOf(p0, p0)
						var _1104_v27 _dafny.Map
						_ = _1104_v27
						_1104_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1103_v26).Cardinality(), Companion_Default___.Fm9(_1040_v0, globalState))
						var _1105_v28 _dafny.MultiSet
						_ = _1105_v28
						_1105_v28 = _dafny.MultiSetOf(Companion_Default___.Fm1((_1103_v26).Cardinality(), _1040_v0, globalState))
						var _1106_v29 _dafny.Map
						_ = _1106_v29
						_1106_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1104_v27, (_1105_v28).IsSubsetOf(_1105_v28))
						_1106_v29 = ((_1106_v29).Merge(_1106_v29)).Merge(_1106_v29)
						(globalState).F8 = (_dafny.Zero).Minus(p0)
						_1045_v5 = (_1045_v5).Merge(Companion_Default___.Fm10(_1040_v0, p2, _1102_cf4, _1041_v1, globalState))
					} else {
						var _1107___mcc_h17 D1 = _source11.Get_().(D1_DC6).Cf12
						_ = _1107___mcc_h17
						var _1108_cf12 D1 = _1107___mcc_h17
						_ = _1108_cf12
						var _1109_v30 _dafny.Array
						_ = _1109_v30
						var _len0_35 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_35
						var _nw200 _dafny.Array
						_ = _nw200
						if _len0_35.Cmp(_dafny.Zero) == 0 {
							_nw200 = _dafny.NewArray(_len0_35)
						} else {
							var _init35 func(_dafny.Int) _dafny.Int = (func(_1110_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1111_i2 _dafny.Int) _dafny.Int {
									return (_1111_i2).Minus(_1110_p1)
								}
							})(p1)
							_ = _init35
							var _element0_35 = _init35(_dafny.Zero)
							_ = _element0_35
							_nw200 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
							(_nw200).ArraySet1(_element0_35, 0)
							var _nativeLen0_35 = (_len0_35).Int()
							_ = _nativeLen0_35
							for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
								(_nw200).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
							}
						}
						_1109_v30 = _nw200
						_1109_v30 = p3
						(globalState).F17 = p0
						var _1112_v31 _dafny.Map
						_ = _1112_v31
						_1112_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1065_v6)
						var _1113_v32 _dafny.Sequence
						_ = _1113_v32
						_1113_v32 = _dafny.SeqOf((func() _dafny.Sequence {
							if (_1112_v31).Contains(_1040_v0) {
								return (_1112_v31).Get(_1040_v0).(_dafny.Sequence)
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("buqwaquu")
						})(), _1065_v6)
						var _1114_v34 _dafny.Sequence
						_ = _1114_v34
						_1114_v34 = _dafny.SeqOf(_1040_v0, _1040_v0, _1040_v0)
						var _1115_v35 _dafny.Sequence
						_ = _1115_v35
						_1115_v35 = _dafny.SeqOf(p1)
						var _1116_v36 _dafny.Set
						_ = _1116_v36
						_1116_v36 = _dafny.SetOf(_1115_v35)
						var _1117_v37 _dafny.Array
						_ = _1117_v37
						var _nwElement0_47 bool = _1040_v0
						_ = _nwElement0_47
						var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(25))
						_ = _nw201
						(_nw201).ArraySet1(_nwElement0_47, 0)
						(_nw201).ArraySet1((func() _dafny.Set {
							var _coll37 = _dafny.NewBuilder()
							_ = _coll37
							for _iter42 := _dafny.Iterate((_dafny.MultiSetFromSeq(_1113_v32)).Elements()); ; {
								_compr_37, _ok42 := _iter42()
								if !_ok42 {
									break
								}
								var _1118_v33 _dafny.Sequence
								_1118_v33 = interface{}(_compr_37).(_dafny.Sequence)
								if (_dafny.MultiSetFromSeq(_1113_v32)).Contains(_1118_v33) {
									_coll37.Add(_1118_v33)
								}
							}
							return _coll37.ToSet()
						}()).IsDisjointFrom(_dafny.SetOf(_1065_v6)), 1)
						(_nw201).ArraySet1(!(true), 2)
						(_nw201).ArraySet1(_1040_v0, 3)
						(_nw201).ArraySet1(true, 4)
						(_nw201).ArraySet1(_1040_v0, 5)
						(_nw201).ArraySet1((_1040_v0) || (true), 6)
						(_nw201).ArraySet1(false, 7)
						(_nw201).ArraySet1(_1040_v0, 8)
						(_nw201).ArraySet1((_1114_v34).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1114_v34).Cardinality()))).Uint32()).(bool), 9)
						(_nw201).ArraySet1(_1040_v0, 10)
						(_nw201).ArraySet1(_1040_v0, 11)
						(_nw201).ArraySet1(_1040_v0, 12)
						(_nw201).ArraySet1(_1040_v0, 13)
						(_nw201).ArraySet1(true, 14)
						(_nw201).ArraySet1(_1040_v0, 15)
						(_nw201).ArraySet1(_1040_v0, 16)
						(_nw201).ArraySet1(_1040_v0, 17)
						(_nw201).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1065_v6, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1065_v6).Cardinality()))).Uint32(), (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()))).Cardinality()), _1040_v0)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1040_v0)), 18)
						(_nw201).ArraySet1(!(_1040_v0), 19)
						(_nw201).ArraySet1(_1040_v0, 20)
						(_nw201).ArraySet1(_1040_v0, 21)
						(_nw201).ArraySet1(_1040_v0, 22)
						(_nw201).ArraySet1(!(_1040_v0), 23)
						(_nw201).ArraySet1(!(_1116_v36).Contains(_1115_v35), 24)
						_1117_v37 = _nw201
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_1117_v37), 0))
						_ = _index162
						(_1117_v37).ArraySet1(_1040_v0, (_index162).Int())
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_1117_v37), 0))
						_ = _index163
						(_1117_v37).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1114_v34).Cardinality()), p2)).Cmp((p1).Minus(p0)) != 0, (_index163).Int())
						var _1119_v38 D0
						_ = _1119_v38
						_1119_v38 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_1040_v0, p2))
						var _1120_v39 D0
						_ = _1120_v39
						_1120_v39 = Companion_D0_.Create_DC2_(_1119_v38)
						var _1121_v40 D0
						_ = _1121_v40
						_1121_v40 = Companion_D0_.Create_DC2_(_1119_v38)
						var _1122_v41 D0
						_ = _1122_v41
						_1122_v41 = Companion_D0_.Create_DC2_(_1119_v38)
						var _1123_v42 D0
						_ = _1123_v42
						_1123_v42 = Companion_D0_.Create_DC2_((_1122_v41).Dtor_cf3())
						var _1124_v43 D0
						_ = _1124_v43
						_1124_v43 = Companion_D0_.Create_DC2_(_1119_v38)
						var _1125_v44 D0
						_ = _1125_v44
						_1125_v44 = Companion_D0_.Create_DC2_(_1121_v40)
						var _1126_v45 _dafny.Sequence
						_ = _1126_v45
						_1126_v45 = _dafny.SeqOf(_1120_v39)
						var _1127_v46 D0
						_ = _1127_v46
						_1127_v46 = Companion_D0_.Create_DC2_((_1126_v45).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1126_v45).Cardinality()))).Uint32()).(D0))
						var _1128_v47 D0
						_ = _1128_v47
						_1128_v47 = Companion_D0_.Create_DC2_(_1125_v44)
						var _1129_v48 *C6
						_ = _1129_v48
						var _nw202 *C6 = New_C6_()
						_ = _nw202
						_nw202.Ctor__(_1128_v47, (_1117_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_1117_v37), 0))).Int()).(bool), (_1117_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_1117_v37), 0))).Int()).(bool), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), p1))
						_1129_v48 = _nw202
					}
					var _1130_v49 _dafny.Array
					_ = _1130_v49
					var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
					_ = _nw203
					_1130_v49 = _nw203
					var _1131_v50 _dafny.Map
					_ = _1131_v50
					_1131_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1130_v49)
					var _1132_v51 _dafny.Sequence
					_ = _1132_v51
					_1132_v51 = _dafny.SeqOf(_1131_v50, _1131_v50)
					_1131_v50 = (((_1132_v51).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1132_v51).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1131_v50)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, _1130_v49))
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1133_v52 _dafny.Array
		_ = _1133_v52
		var _len0_36 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_36
		var _nw204 _dafny.Array
		_ = _nw204
		if _len0_36.Cmp(_dafny.Zero) == 0 {
			_nw204 = _dafny.NewArray(_len0_36)
		} else {
			var _init36 func(_dafny.Int) _dafny.Int = (func(_1134_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1135_i4 _dafny.Int) _dafny.Int {
					return (_1135_i4).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1134_p1)).Cardinality()))))).Cardinality())).Cardinality()))
				}
			})(p1)
			_ = _init36
			var _element0_36 = _init36(_dafny.Zero)
			_ = _element0_36
			_nw204 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
			(_nw204).ArraySet1(_element0_36, 0)
			var _nativeLen0_36 = (_len0_36).Int()
			_ = _nativeLen0_36
			for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
				(_nw204).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
			}
		}
		_1133_v52 = _nw204
		for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1133_v52), 0))); ; {
			_guard_loop_5, _ok43 := _iter43()
			if !_ok43 {
				break
			}
			var _1136_i3 _dafny.Int
			_1136_i3 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1136_i3).Sign() != -1) && ((_1136_i3).Cmp(_dafny.ArrayLen((_1133_v52), 0)) < 0)) {
				(_1133_v52).ArraySet1(Companion_Default___.SafeModuloInt(_1136_i3, (p2).Minus(p2)), (_1136_i3).Int())
			}
		}
		var _1137_v53 _dafny.Map
		_ = _1137_v53
		_1137_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, p3)
		var _1138_v54 _dafny.Sequence
		_ = _1138_v54
		_1138_v54 = _dafny.UnicodeSeqOfUtf8Bytes("etgsrnnx")
		_1137_v53 = (_1137_v53).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_1138_v54, _1138_v54), _1133_v52)
		var _1139_v56 _dafny.Sequence
		_ = _1139_v56
		_1139_v56 = _dafny.SeqOf(p2)
		var _1140_v57 D10
		_ = _1140_v57
		_1140_v57 = Companion_D10_.Create_DC28_(_1139_v56)
		if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(778))).Equals((func() _dafny.Map {
			var _coll38 = _dafny.NewMapBuilder()
			_ = _coll38
			for _iter44 := _dafny.Iterate(((_1140_v57).Dtor_cf52()).Elements()); ; {
				_compr_38, _ok44 := _iter44()
				if !_ok44 {
					break
				}
				var _1141_v55 _dafny.Int
				_1141_v55 = interface{}(_compr_38).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains((_1140_v57).Dtor_cf52(), _1141_v55) {
					_coll38.Add(Companion_Default___.SafeModuloInt(_1141_v55, p2), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1040_v0)).Cardinality())
				}
			}
			return _coll38.ToMap()
		}())) {
			var _1142_v58 *C2
			_ = _1142_v58
			var _nw205 *C2 = New_C2_()
			_ = _nw205
			_nw205.Ctor__(p1)
			_1142_v58 = _nw205
			(globalState).F17 = ((_dafny.Zero).Minus(p2)).Minus((p1).Minus(_dafny.IntOfInt64(518)))
			var _1143_v59 _dafny.MultiSet
			_ = _1143_v59
			_1143_v59 = _dafny.MultiSetOf(p1)
			if (_1143_v59).IsProperSubsetOf(_1143_v59) {
				(globalState).F6 = _1040_v0
				(globalState).F14 = _dafny.IntOfInt64(-370)
				(globalState).F9 = _1040_v0
				var _1144_v60 _dafny.Sequence
				_ = _1144_v60
				_1144_v60 = _dafny.SeqOf((_1143_v59).Union(_1143_v59), (_1143_v59).Intersection(_1143_v59))
				_1144_v60 = _dafny.SeqOf(_1143_v59)
				var _1145_v61 _dafny.Int
				_ = _1145_v61
				var _1146_v62 _dafny.Array
				_ = _1146_v62
				var _1147_v63 bool
				_ = _1147_v63
				var _out11 _dafny.Int
				_ = _out11
				var _out12 _dafny.Array
				_ = _out12
				var _out13 bool
				_ = _out13
				_out11, _out12, _out13 = (_1142_v58).M0(_1040_v0, _1040_v0, globalState)
				_1145_v61 = _out11
				_1146_v62 = _out12
				_1147_v63 = _out13
			} else {
				var _1148_v64 _dafny.MultiSet
				_ = _1148_v64
				_1148_v64 = _dafny.MultiSetOf(_1040_v0, (_1040_v0) == (_1040_v0), !(false), !(_1045_v5).Equals(_1045_v5))
				_1148_v64 = _1148_v64
				(globalState).F15 = _1040_v0
				var _1149_v65 _dafny.Sequence
				_ = _1149_v65
				_1149_v65 = _dafny.SeqOf(_1040_v0)
				var _1150_v66 D6
				_ = _1150_v66
				_1150_v66 = Companion_D6_.Create_DC16_(_1149_v65)
				var _1151_v67 _dafny.Sequence
				_ = _1151_v67
				_1151_v67 = _dafny.SeqOf(_1150_v66, Companion_D6_.Create_DC16_(_1149_v65))
				var _1152_v68 _dafny.Map
				_ = _1152_v68
				_1152_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1151_v67, _1040_v0)
				_1152_v68 = Companion_Default___.Fm40(_1040_v0, globalState)
				var _1153_v69 _dafny.Array
				_ = _1153_v69
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw206
				_1153_v69 = _nw206
				var _1154_v70 _dafny.Map
				_ = _1154_v70
				_1154_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1153_v69, (_1040_v0) || (_1040_v0))
				_1154_v70 = (_1154_v70).Update(_1153_v69, _1040_v0)
				var _1155_v71 _dafny.Map
				_ = _1155_v71
				_1155_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Plus(p0), _1153_v69)
				_1155_v71 = (_1155_v71).Update((p2).Minus(p2), _1153_v69)
			}
			(globalState).F8 = p1
			var _rhs199 bool = !(_1040_v0)
			_ = _rhs199
			var _rhs200 bool = (_1040_v0) == ((_1040_v0) || (_1040_v0))
			_ = _rhs200
			var _rhs201 bool = _1040_v0
			_ = _rhs201
			var _rhs202 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}((func(_1156_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1157_i5 _dafny.Int) _dafny.CodePoint {
					return _1156_v1
				}
			})(_1041_v1)))
			_ = _rhs202
			var _rhs203 _dafny.Int = _dafny.IntOfInt64(264)
			_ = _rhs203
			var _lhs174 *GlobalState = globalState
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			var _lhs177 *GlobalState = globalState
			_ = _lhs177
			_lhs174.F1 = _rhs199
			_lhs175.F6 = _rhs200
			_lhs176.F6 = _rhs201
			_1138_v54 = _rhs202
			_lhs177.F5 = _rhs203
		} else {
			if true {
				_1138_v54 = _dafny.SeqOf(_1041_v1)
				var _1158_v72 *C4
				_ = _1158_v72
				var _nw207 *C4 = New_C4_()
				_ = _nw207
				_nw207.Ctor__(!(_1040_v0), !(true), (p2).Cmp(_dafny.IntOfInt64(456)) > 0)
				_1158_v72 = _nw207
				_1158_v72 = _1158_v72
				_1138_v54 = _dafny.Companion_Sequence_.Concatenate(_1138_v54, _1138_v54)
				var _1159_v73 _dafny.Map
				_ = _1159_v73
				_1159_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1158_v72).F27(), p1)
				_1159_v73 = (_1159_v73).Update(!(!((_1158_v72).F27())), p0)
				var _1160_v74 _dafny.Sequence
				_ = _1160_v74
				_1160_v74 = _dafny.SeqOf(_1040_v0)
				var _1161_v75 _dafny.Array
				_ = _1161_v75
				var _nwElement0_48 _dafny.Sequence = _1139_v56
				_ = _nwElement0_48
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(8))
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_48, 0)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_1162_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1163_i6 _dafny.Int) _dafny.Int {
						return _1162_p2
					}
				})(p2))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(787))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_1164_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1165_i7 _dafny.Int) _dafny.Int {
						return _1164_p1
					}
				})(p1)))), 1)
				(_nw208).ArraySet1(_1139_v56, 2)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_1160_v74).Cardinality()), p1), _1139_v56), 3)
				(_nw208).ArraySet1((_1140_v57).Dtor_cf52(), 4)
				(_nw208).ArraySet1(_1139_v56, 5)
				(_nw208).ArraySet1(_1139_v56, 6)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1139_v56, _1139_v56), 7)
				_1161_v75 = _nw208
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1161_v75), 0))
				_ = _index164
				(_1161_v75).ArraySet1(_1139_v56, (_index164).Int())
				var _1166_v76 _dafny.Set
				_ = _1166_v76
				_1166_v76 = _dafny.SetOf(_dafny.IntOfUint32((_1139_v56).Cardinality()), p2, p1, _dafny.IntOfUint32((_1160_v74).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("luqkew")).Cardinality()))
				var _1167_v77 _dafny.MultiSet
				_ = _1167_v77
				_1167_v77 = _dafny.MultiSetOf(p1)
				var _1168_v80 D8
				_ = _1168_v80
				_1168_v80 = Companion_D8_.Create_DC23_(p0, _1166_v76, p2, p0, p2)
				var _1169_v81 D0
				_ = _1169_v81
				_1169_v81 = Companion_D0_.Create_DC1_(_1040_v0, _dafny.IntOfInt64(244))
				var _1170_v82 _dafny.MultiSet
				_ = _1170_v82
				_1170_v82 = _dafny.MultiSetOf(_1169_v81, _1169_v81)
				var _1171_v83 _dafny.Array
				_ = _1171_v83
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw209
				_1171_v83 = _nw209
				var _1172_v84 _dafny.Set
				_ = _1172_v84
				_1172_v84 = _dafny.SetOf(_1171_v83)
				var _1173_v85 *C0
				_ = _1173_v85
				var _nw210 *C0 = New_C0_()
				_ = _nw210
				_nw210.Ctor__(_1172_v84, _dafny.IntOfUint32((_1138_v54).Cardinality()), p0)
				_1173_v85 = _nw210
				var _1174_v86 D8
				_ = _1174_v86
				_1174_v86 = Companion_D8_.Create_DC24_(p2, (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()), (_1170_v82).Update(_1169_v81, Companion_Default___.Abs((Companion_Default___.Fm32(_1040_v0, _dafny.SeqOf((_1042_v2).Cardinality()), _1160_v74, _dafny.SetOf(Companion_D1_.Create_DC4_((_1158_v72).F27(), (_1158_v72).F27())), globalState)).Cardinality())), _1040_v0, _1173_v85)
				var _1175_v87 _dafny.Map
				_ = _1175_v87
				_1175_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1173_v85.F31, _1173_v85)
				var _1176_v90 _dafny.Array
				_ = _1176_v90
				var _nwElement0_49 _dafny.Set = _1166_v76
				_ = _nwElement0_49
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(21))
				_ = _nw211
				(_nw211).ArraySet1(_nwElement0_49, 0)
				(_nw211).ArraySet1(_1166_v76, 1)
				(_nw211).ArraySet1(_1166_v76, 2)
				(_nw211).ArraySet1(_1166_v76, 3)
				(_nw211).ArraySet1(func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter45 := _dafny.Iterate((_1167_v77).Elements()); ; {
						_compr_39, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _1177_v78 _dafny.Int
						_1177_v78 = interface{}(_compr_39).(_dafny.Int)
						if (_1167_v77).Contains(_1177_v78) {
							_coll39.Add((_1177_v78).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality())))
						}
					}
					return _coll39.ToSet()
				}(), 4)
				(_nw211).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(-874), (_1045_v5).Cardinality())).Difference(_1166_v76), 5)
				(_nw211).ArraySet1(_1166_v76, 6)
				(_nw211).ArraySet1(_1166_v76, 7)
				(_nw211).ArraySet1(_1166_v76, 8)
				(_nw211).ArraySet1((_1166_v76).Union(func() _dafny.Set {
					var _coll40 = _dafny.NewBuilder()
					_ = _coll40
					for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(628), _dafny.IntOfInt64(436))); ; {
						_compr_40, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _1178_v79 _dafny.Int
						_1178_v79 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(628)).Cmp(_1178_v79) <= 0) && ((_1178_v79).Cmp(_dafny.IntOfInt64(436)) < 0) {
							_coll40.Add((_1178_v79).Plus(p1))
						}
					}
					return _coll40.ToSet()
				}()), 9)
				(_nw211).ArraySet1(_dafny.SetOf(p2, p1, (_1045_v5).Cardinality()), 10)
				(_nw211).ArraySet1((_1168_v80).Dtor_cf41(), 11)
				(_nw211).ArraySet1((_1166_v76).Difference(_1166_v76), 12)
				(_nw211).ArraySet1(_1166_v76, 13)
				(_nw211).ArraySet1(_dafny.SetOf(p2, p2, (_1174_v86).Dtor_cf45(), (_1175_v87).Cardinality()), 14)
				(_nw211).ArraySet1(_1166_v76, 15)
				(_nw211).ArraySet1(_1166_v76, 16)
				(_nw211).ArraySet1((func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(51), _dafny.IntOfInt64(172))); ; {
						_compr_41, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1179_v88 _dafny.Int
						_1179_v88 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(51)).Cmp(_1179_v88) <= 0) && ((_1179_v88).Cmp(_dafny.IntOfInt64(172)) < 0) {
							_coll41.Add(Companion_Default___.SafeDivisionInt(_1179_v88, _dafny.IntOfUint32((_1138_v54).Cardinality())))
						}
					}
					return _coll41.ToSet()
				}()).Union(_1166_v76), 17)
				(_nw211).ArraySet1(Companion_Default___.Fm41(globalState), 18)
				(_nw211).ArraySet1(_1166_v76, 19)
				(_nw211).ArraySet1(func() _dafny.Set {
					var _coll42 = _dafny.NewBuilder()
					_ = _coll42
					for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(806), _dafny.IntOfInt64(-107))); ; {
						_compr_42, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1180_v89 _dafny.Int
						_1180_v89 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(806)).Cmp(_1180_v89) <= 0) && ((_1180_v89).Cmp(_dafny.IntOfInt64(-107)) < 0) {
							_coll42.Add((_1180_v89).Minus(_1173_v85.F31))
						}
					}
					return _coll42.ToSet()
				}(), 20)
				_1176_v90 = _nw211
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1161_v75), 0))
				_ = _index165
				var _rhs204 _dafny.Sequence = _1139_v56
				_ = _rhs204
				var _rhs205 _dafny.Array = _1176_v90
				_ = _rhs205
				var _rhs206 bool = (_1158_v72).F27()
				_ = _rhs206
				var _lhs178 _dafny.Array = _1161_v75
				_ = _lhs178
				var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1161_v75), 0))
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				(_lhs178).ArraySet1(_rhs204, (_lhs179).Int())
				_1176_v90 = _rhs205
				_lhs180.F1 = _rhs206
			} else {
				var _1181_v91 _dafny.MultiSet
				_ = _1181_v91
				_1181_v91 = _dafny.MultiSetOf(_1040_v0, _1040_v0)
				_1040_v0 = ((_1181_v91).Difference(_1181_v91)).Contains(false)
				var _1182_v92 *C5
				_ = _1182_v92
				var _nw212 *C5 = New_C5_()
				_ = _nw212
				_nw212.Ctor__(Companion_Default___.Fm29((_1181_v91).Cardinality(), p2, _1040_v0, globalState), _1040_v0)
				_1182_v92 = _nw212
				var _nw213 *C5 = New_C5_()
				_ = _nw213
				_nw213.Ctor__(_1040_v0, _1040_v0)
				_1182_v92 = _nw213
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))
				_ = _index166
				(_1044_v4).ArraySet1CodePoint(_1041_v1, (_index166).Int())
				var _rhs207 _dafny.Int = (func() _dafny.Int {
					if false {
						return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(698), p2)
					}
					return _dafny.IntOfUint32((_1138_v54).Cardinality())
				})()
				_ = _rhs207
				var _rhs208 _dafny.Sequence = _1138_v54
				_ = _rhs208
				var _lhs181 *GlobalState = globalState
				_ = _lhs181
				_lhs181.F17 = _rhs207
				_1138_v54 = _rhs208
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))
				_ = _index167
				(_1044_v4).ArraySet1CodePoint(_1041_v1, (_index167).Int())
			}
			(globalState).F14 = Companion_Default___.Fm1(p0, true, globalState)
			var _1183_v93 _dafny.Map
			_ = _1183_v93
			_1183_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1040_v0)
			var _1184_v94 D0
			_ = _1184_v94
			_1184_v94 = Companion_D0_.Create_DC0_((func() bool {
				if (_1183_v93).Contains(p0) {
					return (_1183_v93).Get(p0).(bool)
				}
				return _1040_v0
			})())
			var _1185_v95 D0
			_ = _1185_v95
			_1185_v95 = Companion_D0_.Create_DC2_(_1184_v94)
			var _source12 D0 = _1185_v95
			_ = _source12
			if _source12.Is_DC1() {
				var _1186___mcc_h18 bool = _source12.Get_().(D0_DC1).Cf1
				_ = _1186___mcc_h18
				var _1187___mcc_h19 _dafny.Int = _source12.Get_().(D0_DC1).Cf2
				_ = _1187___mcc_h19
				var _1188_cf2 _dafny.Int = _1187___mcc_h19
				_ = _1188_cf2
				var _1189_cf1 bool = _1186___mcc_h18
				_ = _1189_cf1
				var _1190_v96 *C6
				_ = _1190_v96
				var _nw214 *C6 = New_C6_()
				_ = _nw214
				_nw214.Ctor__(_1185_v95, _1189_cf1, _1189_cf1, Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ohhj")).Cardinality()), _1189_cf1, globalState))
				_1190_v96 = _nw214
				var _1191_v97 _dafny.Map
				_ = _1191_v97
				_1191_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("urgll"), _1190_v96)
				_1191_v97 = (_1191_v97).Update(_1138_v54, _1190_v96)
				var _1192_v98 *C2
				_ = _1192_v98
				var _nw215 *C2 = New_C2_()
				_ = _nw215
				_nw215.Ctor__((_1190_v96).Fm5(p0, _dafny.IntOfUint32((Companion_Default___.Fm31(_1189_cf1, _1040_v0, globalState)).Cardinality()), globalState))
				_1192_v98 = _nw215
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1133_v52), 0))
				_ = _index168
				(_1133_v52).ArraySet1(_1188_cf2, (_index168).Int())
				var _1193_v99 _dafny.Sequence
				_ = _1193_v99
				_1193_v99 = _dafny.SeqOf(_1189_cf1, _1189_cf1)
				var _1194_v100 _dafny.Map
				_ = _1194_v100
				_1194_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1193_v99, p2)
				var _1195_v101 _dafny.Map
				_ = _1195_v101
				_1195_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1190_v96).Fm2(globalState), _1192_v98)
				var _1196_v102 _dafny.Sequence
				_ = _1196_v102
				_1196_v102 = _dafny.SeqOf((_1195_v101).Update(p1, _1192_v98), _1195_v101, _1195_v101)
				var _1197_v103 _dafny.Map
				_ = _1197_v103
				_1197_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1196_v102).Cardinality()), p1)
				var _1198_v104 _dafny.MultiSet
				_ = _1198_v104
				_1198_v104 = _dafny.MultiSetOf(_1041_v1, (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()), _1041_v1)
				var _1199_v105 D0
				_ = _1199_v105
				_1199_v105 = Companion_D0_.Create_DC0_((_1198_v104).IsProperSubsetOf(_1198_v104))
				var _1200_v106 D15
				_ = _1200_v106
				_1200_v106 = Companion_D15_.Create_DC36_(_1194_v100)
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1133_v52), 0))
				_ = _index169
				var _rhs209 _dafny.Int = _1188_cf2
				_ = _rhs209
				var _rhs210 _dafny.Map = (_1194_v100).Merge((_1200_v106).Dtor_cf64())
				_ = _rhs210
				var _rhs211 _dafny.Sequence = _1138_v54
				_ = _rhs211
				var _rhs212 _dafny.Map = (Companion_Default___.Fm8((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1188_cf2, ((_1183_v93).Update((func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(109), _dafny.IntOfInt64(642))); ; {
						_compr_43, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1201_v107 _dafny.Int
						_1201_v107 = interface{}(_compr_43).(_dafny.Int)
						if ((_dafny.IntOfInt64(109)).Cmp(_1201_v107) <= 0) && ((_1201_v107).Cmp(_dafny.IntOfInt64(642)) < 0) {
							_coll43.Add((_1201_v107).Times(p1), _1189_cf1)
						}
					}
					return _coll43.ToMap()
				}()).Cardinality(), _1189_cf1)).Cardinality())).Cardinality(), globalState)).Merge((_1197_v103).Merge(_1197_v103))
				_ = _rhs212
				var _rhs213 D0 = _1199_v105
				_ = _rhs213
				var _lhs182 _dafny.Array = _1133_v52
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1133_v52), 0))
				_ = _lhs183
				(_lhs182).ArraySet1(_rhs209, (_lhs183).Int())
				_1194_v100 = _rhs210
				_1138_v54 = _rhs211
				_1197_v103 = _rhs212
				_1199_v105 = _rhs213
				_1188_cf2 = (_1133_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1133_v52), 0))).Int()).(_dafny.Int)
			} else if _source12.Is_DC0() {
				var _1202___mcc_h20 bool = _source12.Get_().(D0_DC0).Cf0
				_ = _1202___mcc_h20
				var _1203_cf0 bool = _1202___mcc_h20
				_ = _1203_cf0
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p3), 0))
				_ = _index170
				(p3).ArraySet1(p0, (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p3), 0))
				_ = _index171
				(p3).ArraySet1(Companion_Default___.Fm1(p0, _1203_cf0, globalState), (_index171).Int())
				(globalState).F9 = false
				var _1204_v108 *C3
				_ = _1204_v108
				var _nw216 *C3 = New_C3_()
				_ = _nw216
				_nw216.Ctor__(_1040_v0, p1, _1040_v0, _1203_cf0, p1)
				_1204_v108 = _nw216
				var _1205_v109 _dafny.Array
				_ = _1205_v109
				var _nwElement0_50 *C3 = _1204_v108
				_ = _nwElement0_50
				var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(18))
				_ = _nw217
				(_nw217).ArraySet1(_nwElement0_50, 0)
				(_nw217).ArraySet1(_1204_v108, 1)
				(_nw217).ArraySet1(_1204_v108, 2)
				(_nw217).ArraySet1(_1204_v108, 3)
				(_nw217).ArraySet1(_1204_v108, 4)
				(_nw217).ArraySet1(_1204_v108, 5)
				(_nw217).ArraySet1(_1204_v108, 6)
				(_nw217).ArraySet1(_1204_v108, 7)
				(_nw217).ArraySet1(_1204_v108, 8)
				(_nw217).ArraySet1(_1204_v108, 9)
				(_nw217).ArraySet1(_1204_v108, 10)
				(_nw217).ArraySet1(_1204_v108, 11)
				(_nw217).ArraySet1(_1204_v108, 12)
				(_nw217).ArraySet1(_1204_v108, 13)
				(_nw217).ArraySet1(_1204_v108, 14)
				(_nw217).ArraySet1(_1204_v108, 15)
				(_nw217).ArraySet1(_1204_v108, 16)
				(_nw217).ArraySet1(_1204_v108, 17)
				_1205_v109 = _nw217
				var _1206_v110 D16
				_ = _1206_v110
				_1206_v110 = Companion_D16_.Create_DC39_(_1205_v109)
				_1205_v109 = (_1206_v110).Dtor_cf69()
				var _1207_v111 T0
				_ = _1207_v111
				var _nw218 *C3 = New_C3_()
				_ = _nw218
				_nw218.Ctor__(_1203_cf0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1203_cf0)).Cardinality(), _1203_cf0, (_1204_v108).F28(), _dafny.IntOfInt64(150))
				_1207_v111 = _nw218
				_1207_v111 = _1207_v111
			} else {
				var _1208___mcc_h21 D0 = _source12.Get_().(D0_DC2).Cf3
				_ = _1208___mcc_h21
				var _1209_cf3 D0 = _1208___mcc_h21
				_ = _1209_cf3
				var _1210_v112 _dafny.Array
				_ = _1210_v112
				var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
				_ = _nw219
				_1210_v112 = _nw219
				var _1211_v113 _dafny.MultiSet
				_ = _1211_v113
				_1211_v113 = _dafny.MultiSetOf(_1040_v0)
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_1210_v112), 0))
				_ = _index172
				(_1210_v112).ArraySet1(_1211_v113, (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_1210_v112), 0))
				_ = _index173
				(_1210_v112).ArraySet1(_1211_v113, (_index173).Int())
				var _1212_v114 _dafny.Sequence
				_ = _1212_v114
				_1212_v114 = _dafny.SeqOf(_1040_v0)
				var _1213_v115 _dafny.Set
				_ = _1213_v115
				_1213_v115 = _dafny.SetOf((_1212_v114).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1138_v54, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1138_v54).Cardinality()))).Uint32(), (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()))).Cardinality()), _dafny.IntOfUint32((_1212_v114).Cardinality()))).Uint32()).(bool))
				var _1214_v116 _dafny.MultiSet
				_ = _1214_v116
				_1214_v116 = _dafny.MultiSetOf(_1213_v115, _1213_v115, _1213_v115, _dafny.SetOf(_1040_v0, _1040_v0, _1040_v0, !(_1040_v0)))
				_1214_v116 = _dafny.MultiSetOf(_1213_v115)
				var _1215_v117 _dafny.Map
				_ = _1215_v117
				_1215_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1212_v114, _1212_v114)
				_1215_v117 = (_1215_v117).Update(_1212_v114, _1212_v114)
				_1138_v54 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqOf(_dafny.CodePoint('f'))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqOf(_dafny.CodePoint('f')))).Cardinality()))).Uint32(), _1041_v1), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqOf(_dafny.CodePoint('f'))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqOf(_dafny.CodePoint('f')))).Cardinality()))).Uint32(), _1041_v1)).Cardinality()))).Uint32(), _1041_v1), _dafny.Companion_Sequence_.Concatenate(_1138_v54, _1138_v54))
			}
			var _1216_v118 T1
			_ = _1216_v118
			var _nw220 *C5 = New_C5_()
			_ = _nw220
			_nw220.Ctor__(!(_1040_v0), _1040_v0)
			_1216_v118 = _nw220
			var _1217_v119 _dafny.Map
			_ = _1217_v119
			_1217_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1216_v118, (func() bool {
				if (_1183_v93).Contains(p1) {
					return (_1183_v93).Get(p1).(bool)
				}
				return (_1216_v118).F25()
			})())
			if Companion_Default___.Fm29((_1217_v119).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(p2)), !((_1216_v118).F25()) || (_1216_v118.F24()), globalState) {
				var _1218_v120 _dafny.Map
				_ = _1218_v120
				_1218_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				(globalState).F5 = (func() _dafny.Int {
					if (_1218_v120).Contains(p2) {
						return (_1218_v120).Get(p2).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(p1)
				})()
				var _1219_v121 D3
				_ = _1219_v121
				_1219_v121 = Companion_D3_.Create_DC8_(_1138_v54)
				var _1220_v122 D8
				_ = _1220_v122
				_1220_v122 = Companion_D8_.Create_DC22_(p3)
				var _1221_v123 _dafny.MultiSet
				_ = _1221_v123
				_1221_v123 = _dafny.MultiSetOf(_1133_v52, (_1220_v122).Dtor_cf39())
				var _1222_v124 _dafny.Map
				_ = _1222_v124
				_1222_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25(_1040_v0, (_1216_v118).F25(), _dafny.CodePoint('j'), globalState), (_1219_v121).Dtor_cf14()), _1221_v123)
				_1222_v124 = (_1222_v124).Update(_1138_v54, (_1221_v123).Update(_1133_v52, Companion_Default___.Abs(p0)))
				(globalState).F5 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1138_v54).Cardinality()), (_1218_v120).Cardinality())
				_1041_v1 = (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int())
				_1138_v54 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(151))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_1223_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1224_i8 _dafny.Int) _dafny.CodePoint {
						return _1223_v1
					}
				})(_1041_v1)))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1138_v54, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(151))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_1225_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1226_i8 _dafny.Int) _dafny.CodePoint {
						return _1225_v1
					}
				})(_1041_v1))))).Cardinality()))).Uint32(), (_1044_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1044_v4), 0))).Int()))
			} else {
				var _1227_v125 _dafny.Map
				_ = _1227_v125
				_1227_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-822)).Plus(p0), _1133_v52)
				var _1228_v126 D15
				_ = _1228_v126
				_1228_v126 = Companion_D15_.Create_DC37_((_1216_v118).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, (_1216_v118).F25()), p2)
				var _1229_v127 _dafny.Map
				_ = _1229_v127
				_1229_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1228_v126).Dtor_cf65(), p0)
				var _1230_v128 _dafny.MultiSet
				_ = _1230_v128
				_1230_v128 = _dafny.MultiSetOf(_1229_v127)
				_1227_v125 = (_1227_v125).Update(Companion_Default___.SafeDivisionInt((_1230_v128).Cardinality(), p2), p3)
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))
				_ = _index174
				(_1133_v52).ArraySet1(_dafny.IntOfUint32((_1138_v54).Cardinality()), (_index174).Int())
				var _1231_v129 *C6
				_ = _1231_v129
				var _nw221 *C6 = New_C6_()
				_ = _nw221
				_nw221.Ctor__(_1185_v95, (_1216_v118).F25(), (_1216_v118).F25(), p0)
				_1231_v129 = _nw221
				var _1232_v130 _dafny.Map
				_ = _1232_v130
				_1232_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1231_v129, (p2).Cmp(_dafny.IntOfInt64(151)) < 0)
				var _1233_v131 _dafny.Array
				_ = _1233_v131
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_37
				var _nw222 _dafny.Array
				_ = _nw222
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw222 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) bool = (func(_1234_v118 T1) func(_dafny.Int) bool {
						return func(_1235_i9 _dafny.Int) bool {
							return _1234_v118.F24()
						}
					})(_1216_v118)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw222 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw222).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw222).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1233_v131 = _nw222
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))
				_ = _index175
				var _rhs214 bool = (func() bool {
					if (_1232_v130).Contains(_1231_v129) {
						return (_1232_v130).Get(_1231_v129).(bool)
					}
					return (_1216_v118).F25()
				})()
				_ = _rhs214
				var _rhs215 bool = _1040_v0
				_ = _rhs215
				var _rhs216 _dafny.Int = (p0).Times((_1139_v56).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1139_v56).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs216
				var _rhs217 _dafny.Array = _1233_v131
				_ = _rhs217
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				var _lhs185 _dafny.Array = _1133_v52
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))
				_ = _lhs186
				_1040_v0 = _rhs214
				_lhs184.F6 = _rhs215
				(_lhs185).ArraySet1(_rhs216, (_lhs186).Int())
				r1 = _rhs217
				(globalState).F15 = _1040_v0
				var _1236_v132 _dafny.Array
				_ = _1236_v132
				_1236_v132 = _1233_v131
				var _1237_v133 _dafny.Map
				_ = _1237_v133
				_1237_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1040_v0, (_1236_v132))
				var _1238_v134 _dafny.Array
				_ = _1238_v134
				var _nwElement0_51 _dafny.Array = _1233_v131
				_ = _nwElement0_51
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(28))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_51, 0)
				(_nw223).ArraySet1(_1233_v131, 1)
				(_nw223).ArraySet1(_1233_v131, 2)
				(_nw223).ArraySet1(_1233_v131, 3)
				(_nw223).ArraySet1(_1233_v131, 4)
				(_nw223).ArraySet1((_1236_v132), 5)
				(_nw223).ArraySet1(_1233_v131, 6)
				(_nw223).ArraySet1(_1233_v131, 7)
				(_nw223).ArraySet1(_1233_v131, 8)
				(_nw223).ArraySet1(_1233_v131, 9)
				(_nw223).ArraySet1(_1233_v131, 10)
				(_nw223).ArraySet1(_1233_v131, 11)
				(_nw223).ArraySet1(_1233_v131, 12)
				(_nw223).ArraySet1(_1233_v131, 13)
				(_nw223).ArraySet1(_1233_v131, 14)
				(_nw223).ArraySet1((func() _dafny.Array {
					if (_1237_v133).Contains(_1216_v118.F24()) {
						return (_1237_v133).Get(_1216_v118.F24()).(_dafny.Array)
					}
					return _1233_v131
				})(), 15)
				(_nw223).ArraySet1(_1233_v131, 16)
				(_nw223).ArraySet1(_1233_v131, 17)
				(_nw223).ArraySet1(_1233_v131, 18)
				(_nw223).ArraySet1(_1233_v131, 19)
				(_nw223).ArraySet1(_1233_v131, 20)
				(_nw223).ArraySet1(_1233_v131, 21)
				(_nw223).ArraySet1(_1233_v131, 22)
				(_nw223).ArraySet1(_1233_v131, 23)
				(_nw223).ArraySet1(_1233_v131, 24)
				(_nw223).ArraySet1(_1233_v131, 25)
				(_nw223).ArraySet1(_1233_v131, 26)
				(_nw223).ArraySet1(_1233_v131, 27)
				_1238_v134 = _nw223
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1238_v134), 0))
				_ = _index176
				(_1238_v134).ArraySet1((func() _dafny.Array {
					if (_1237_v133).Contains((_1216_v118).F25()) {
						return (_1237_v133).Get((_1216_v118).F25()).(_dafny.Array)
					}
					return _1233_v131
				})(), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1238_v134), 0))
				_ = _index177
				var _rhs218 _dafny.Array = _1233_v131
				_ = _rhs218
				var _rhs219 _dafny.Int = (_1133_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))).Int()).(_dafny.Int)
				_ = _rhs219
				var _rhs220 bool = !(!(!(Companion_Default___.Fm29((_1133_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))).Int()).(_dafny.Int), p2, (_1216_v118).F25(), globalState))))
				_ = _rhs220
				var _lhs187 _dafny.Array = _1238_v134
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1238_v134), 0))
				_ = _lhs188
				var _lhs189 *GlobalState = globalState
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				(_lhs187).ArraySet1(_rhs218, (_lhs188).Int())
				_lhs189.F17 = _rhs219
				_lhs190.F15 = _rhs220
				var _1239_v135 _dafny.Array
				_ = _1239_v135
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_38
				var _nw224 _dafny.Array
				_ = _nw224
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw224 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Map = (func(_1240_v93 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1241_i10 _dafny.Int) _dafny.Map {
							return (_1240_v93).Merge(_1240_v93)
						}
					})(_1183_v93)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw224 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw224).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw224).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1239_v135 = _nw224
				var _1242_v136 _dafny.MultiSet
				_ = _1242_v136
				_1242_v136 = _dafny.MultiSetOf(p2)
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))
				_ = _index178
				var _rhs221 _dafny.Int = _dafny.IntOfInt64(481)
				_ = _rhs221
				var _rhs222 _dafny.Array = _1239_v135
				_ = _rhs222
				var _rhs223 bool = (_1216_v118).F25()
				_ = _rhs223
				var _rhs224 _dafny.Int = ((_1133_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))).Int()).(_dafny.Int)).Times(((_1242_v136).Intersection(_1242_v136)).Cardinality())
				_ = _rhs224
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				var _lhs193 _dafny.Array = _1133_v52
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_1133_v52), 0))
				_ = _lhs194
				_lhs191.F5 = _rhs221
				_1239_v135 = _rhs222
				_lhs192.F1 = _rhs223
				(_lhs193).ArraySet1(_rhs224, (_lhs194).Int())
			}
			var _1243_v137 _dafny.Map
			_ = _1243_v137
			_1243_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1138_v54, p0)
			_1243_v137 = (_1243_v137).Update(_1138_v54, p2)
		}
		var _1244_v138 T1
		_ = _1244_v138
		var _nw225 *C3 = New_C3_()
		_ = _nw225
		_nw225.Ctor__(_1040_v0, _dafny.IntOfUint32((_1139_v56).Cardinality()), _1040_v0, _1040_v0, p2)
		_1244_v138 = _nw225
		var _1245_v139 _dafny.Map
		_ = _1245_v139
		_1245_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1138_v54).Cardinality()), _1244_v138)
		var _1246_v140 _dafny.Map
		_ = _1246_v140
		_1246_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1245_v139)
		var _1247_v141 _dafny.Sequence
		_ = _1247_v141
		_1247_v141 = _dafny.SeqOf(Companion_Default___.Fm9(_1244_v138.F24(), globalState), Companion_Default___.Fm29(p1, p2, _1040_v0, globalState))
		r0 = (func() _dafny.Map {
			if (_1246_v140).Contains(_dafny.IntOfUint32((_1247_v141).Cardinality())) {
				return (_1246_v140).Get(_dafny.IntOfUint32((_1247_v141).Cardinality())).(_dafny.Map)
			}
			return (_1245_v139).Merge(_1245_v139)
		})()
		var _1248_v142 _dafny.Array
		_ = _1248_v142
		var _nw226 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw226
		_1248_v142 = _nw226
		r1 = _1248_v142
		return r0, r1
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f24 bool
	_f23 _dafny.Int
	_f25 bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f24 = false
	_this._f23 = _dafny.Zero
	_this._f25 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C8{}
var _ T1 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F24() bool {
	return _this._f24
}
func (_this *C8) F24_set_(value bool) {
	_this._f24 = value
}
func (_this *C8) F23() _dafny.Int {
	return _this._f23
}
func (_this *C8) F25() bool {
	return _this._f25
}
func (_this *C8) Ctor__(f23 _dafny.Int, f24 bool, f25 bool) {
	{
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C8) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F23()
	}
}
func (_this *C8) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("motvhqyi"), _dafny.UnicodeSeqOfUtf8Bytes("hcsjvtutw")))
	}
}
func (_this *C8) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _this.F24()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F25())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), (_this).F25())))
	}
}
func (_this *C8) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F23()
	}
}
func (_this *C8) Fm6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((_this).F25())
	}
}
func (_this *C8) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		(_this).F24_set_(false)
		var _1249_i0 _dafny.Int
		_ = _1249_i0
		_1249_i0 = _dafny.Zero
		{
			for !(!(!((_this).F25()) || (p1)) || (_this.F24())) {
				{
					if (_1249_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1249_i0 = (_1249_i0).Plus(_dafny.One)
					var _1250_v0 _dafny.Sequence
					_ = _1250_v0
					_1250_v0 = _dafny.UnicodeSeqOfUtf8Bytes("luxyg")
					var _1251_v1 _dafny.Sequence
					_ = _1251_v1
					_1251_v1 = _dafny.SeqOf(_dafny.IntOfUint32((_1250_v0).Cardinality()))
					var _1252_v2 D0
					_ = _1252_v2
					_1252_v2 = Companion_D0_.Create_DC0_(p0)
					var _1253_v3 _dafny.Array
					_ = _1253_v3
					var _nwElement0_52 bool = (_this).Fm6(_dafny.IntOfUint32((_dafny.SeqOf((_this).F23())).Cardinality()), _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(44))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_1254_p1 bool) func(_dafny.Int) _dafny.Int {
						return func(_1255_i2 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqOf(_1254_p1, true)).Cardinality())
						}
					})(p1))), _1251_v1), (_this).F25(), (_this).F23(), globalState)
					_ = _nwElement0_52
					var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(22))
					_ = _nw227
					(_nw227).ArraySet1(_nwElement0_52, 0)
					(_nw227).ArraySet1((_this).F25(), 1)
					(_nw227).ArraySet1((_this).F25(), 2)
					(_nw227).ArraySet1(true, 3)
					(_nw227).ArraySet1(p1, 4)
					(_nw227).ArraySet1((_this).F25(), 5)
					(_nw227).ArraySet1(true, 6)
					(_nw227).ArraySet1(false, 7)
					(_nw227).ArraySet1((_this).F25(), 8)
					(_nw227).ArraySet1(p1, 9)
					(_nw227).ArraySet1((_1252_v2).Dtor_cf0(), 10)
					(_nw227).ArraySet1(false, 11)
					(_nw227).ArraySet1(_this.F24(), 12)
					(_nw227).ArraySet1((_this).F25(), 13)
					(_nw227).ArraySet1((_1252_v2).Dtor_cf0(), 14)
					(_nw227).ArraySet1(true, 15)
					(_nw227).ArraySet1(p0, 16)
					(_nw227).ArraySet1(p0, 17)
					(_nw227).ArraySet1(p1, 18)
					(_nw227).ArraySet1(p0, 19)
					(_nw227).ArraySet1((_this).F25(), 20)
					(_nw227).ArraySet1(true, 21)
					_1253_v3 = _nw227
					var _1256_v4 _dafny.Map
					_ = _1256_v4
					_1256_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_v3, (_this).F23())
					(globalState).F8 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(600))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg51 _dafny.Int) interface{} {
							return coer51(arg51)
						}
					}(func(_1257_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					}))).Cardinality()))).Minus((_1256_v4).Cardinality())
					var _1258_v5 *C6
					_ = _1258_v5
					var _nw228 *C6 = New_C6_()
					_ = _nw228
					_nw228.Ctor__(Companion_Default___.Fm24((_this).F23(), (_this).F25(), (_this).F25(), globalState), _this.F24(), _dafny.Companion_Sequence_.IsPrefixOf(_1250_v0, _dafny.UnicodeSeqOfUtf8Bytes("s")), Companion_Default___.Fm1((_this).F23(), true, globalState))
					_1258_v5 = _nw228
					var _1259_v6 _dafny.Map
					_ = _1259_v6
					_1259_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _this.F24())
					_1259_v6 = (_1259_v6).Update((_this).F23(), p0)
					var _1260_v7 *C7
					_ = _1260_v7
					var _nw229 *C7 = New_C7_()
					_ = _nw229
					_nw229.Ctor__()
					_1260_v7 = _nw229
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		(globalState).F14 = _dafny.IntOfInt64(846)
		var _1261_v8 *C5
		_ = _1261_v8
		var _nw230 *C5 = New_C5_()
		_ = _nw230
		_nw230.Ctor__(false, (_this).F25())
		_1261_v8 = _nw230
		var _1262_v9 _dafny.Sequence
		_ = _1262_v9
		_1262_v9 = _dafny.UnicodeSeqOfUtf8Bytes("xwvnhkl")
		var _1263_v10 _dafny.Set
		_ = _1263_v10
		_1263_v10 = _dafny.SetOf(_dafny.IntOfUint32((_1262_v9).Cardinality()), (_this).F23())
		var _1264_v11 _dafny.Set
		_ = _1264_v11
		_1264_v11 = _dafny.SetOf(_1263_v10, _1263_v10, _dafny.SetOf((_this).F23()))
		(globalState).F5 = (_1264_v11).Cardinality()
		var _1265_v12 D0
		_ = _1265_v12
		_1265_v12 = Companion_D0_.Create_DC0_(_this.F24())
		var _1266_v13 D0
		_ = _1266_v13
		_1266_v13 = Companion_D0_.Create_DC2_(_1265_v12)
		_1266_v13 = _1266_v13
		r0 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iasboupi")).Cardinality())
		var _1267_v14 _dafny.Array
		_ = _1267_v14
		var _nw231 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw231
		_1267_v14 = _nw231
		r1 = (func() _dafny.Array {
			if (_this).F25() {
				return _1267_v14
			}
			return _1267_v14
		})()
		r2 = p1
		return r0, r1, r2
	}
}
func (_this *C8) M1(globalState *GlobalState) (bool, T0, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 T0 = (T0)(nil)
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1268_v0 _dafny.Array
		_ = _1268_v0
		var _nw232 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw232
		_1268_v0 = _nw232
		var _1269_v1 _dafny.Set
		_ = _1269_v1
		_1269_v1 = _dafny.SetOf(_1268_v0)
		var _1270_v2 _dafny.Map
		_ = _1270_v2
		_1270_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(122), _dafny.IntOfInt64(-910))
		var _1271_v3 *C0
		_ = _1271_v3
		var _nw233 *C0 = New_C0_()
		_ = _nw233
		_nw233.Ctor__(_1269_v1, ((_this).F23()).Plus((_1270_v2).Cardinality()), (_this).F23())
		_1271_v3 = _nw233
		var _1272_i0 _dafny.Int
		_ = _1272_i0
		_1272_i0 = _dafny.Zero
		{
			for !(_this.F24()) {
				{
					if (_1272_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1272_i0 = (_1272_i0).Plus(_dafny.One)
					var _1273_v4 _dafny.Array
					_ = _1273_v4
					var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
					_ = _nw234
					_1273_v4 = _nw234
					_1273_v4 = _1273_v4
					var _1274_v5 *C7
					_ = _1274_v5
					var _nw235 *C7 = New_C7_()
					_ = _nw235
					_nw235.Ctor__()
					_1274_v5 = _nw235
					var _1275_v6 _dafny.Sequence
					_ = _1275_v6
					_1275_v6 = _dafny.SeqOf(_1271_v3.F31)
					var _1276_v7 _dafny.Map
					_ = _1276_v7
					_1276_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1274_v5, _1275_v6)
					var _1277_v8 _dafny.MultiSet
					_ = _1277_v8
					_1277_v8 = _dafny.MultiSetOf((_this).F25())
					(globalState).F14 = (func() _dafny.Int {
						if (_this).F25() {
							return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
								if (_1276_v7).Contains(_1274_v5) {
									return (_1276_v7).Get(_1274_v5).(_dafny.Sequence)
								}
								return _1275_v6
							})(), _1275_v6)).Cardinality())
						}
						return (_1277_v8).Cardinality()
					})()
					var _1278_v9 _dafny.Sequence
					_ = _1278_v9
					_1278_v9 = _dafny.SeqOf((_this).F25())
					var _1279_v10 *C5
					_ = _1279_v10
					var _nw236 *C5 = New_C5_()
					_ = _nw236
					_nw236.Ctor__(!((_this).F25()), (_1278_v9).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1278_v9).Cardinality()))).Uint32()).(bool))
					_1279_v10 = _nw236
					_1279_v10 = _1279_v10
					(globalState).F8 = _1271_v3.F31
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1280_v11 _dafny.MultiSet
		_ = _1280_v11
		_1280_v11 = _dafny.MultiSetOf((_this).F23())
		var _1281_v12 _dafny.Sequence
		_ = _1281_v12
		_1281_v12 = _dafny.SeqOf((_this).F25(), (_this).F25(), _this.F24())
		var _1282_v13 _dafny.Sequence
		_ = _1282_v13
		_1282_v13 = _dafny.SeqOf(true, (_1280_v11).Equals(_1280_v11), (_1281_v12).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1281_v12).Cardinality()))).Uint32()).(bool), !(Companion_Default___.Fm29(_1271_v3.F31, (func() _dafny.Int {
			if (_1280_v11).Contains((_this).F23()) {
				return (_1280_v11).Multiplicity((_this).F23())
			}
			return (_this).F23()
		})(), (_this).F25(), globalState)))
		var _rhs225 _dafny.Int = (_this).F23()
		_ = _rhs225
		var _rhs226 bool = (_1282_v13).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F23(), _1271_v3.F31), _dafny.IntOfUint32((_1282_v13).Cardinality()))).Uint32()).(bool)
		_ = _rhs226
		var _lhs195 *GlobalState = globalState
		_ = _lhs195
		var _lhs196 *C8 = _this
		_ = _lhs196
		_lhs195.F5 = _rhs225
		_lhs196.F24_set_(_rhs226)
		var _1283_v14 _dafny.Array
		_ = _1283_v14
		var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw237
		_1283_v14 = _nw237
		var _1284_v15 D8
		_ = _1284_v15
		_1284_v15 = Companion_D8_.Create_DC22_(_1283_v14)
		var _source13 D8 = _1284_v15
		_ = _source13
		if _source13.Is_DC23() {
			var _1285___mcc_h0 _dafny.Int = _source13.Get_().(D8_DC23).Cf40
			_ = _1285___mcc_h0
			var _1286___mcc_h1 _dafny.Set = _source13.Get_().(D8_DC23).Cf41
			_ = _1286___mcc_h1
			var _1287___mcc_h2 _dafny.Int = _source13.Get_().(D8_DC23).Cf42
			_ = _1287___mcc_h2
			var _1288___mcc_h3 _dafny.Int = _source13.Get_().(D8_DC23).Cf43
			_ = _1288___mcc_h3
			var _1289___mcc_h4 _dafny.Int = _source13.Get_().(D8_DC23).Cf44
			_ = _1289___mcc_h4
			var _1290_cf44 _dafny.Int = _1289___mcc_h4
			_ = _1290_cf44
			var _1291_cf43 _dafny.Int = _1288___mcc_h3
			_ = _1291_cf43
			var _1292_cf42 _dafny.Int = _1287___mcc_h2
			_ = _1292_cf42
			var _1293_cf41 _dafny.Set = _1286___mcc_h1
			_ = _1293_cf41
			var _1294_cf40 _dafny.Int = _1285___mcc_h0
			_ = _1294_cf40
			var _1295_v16 _dafny.Set
			_ = _1295_v16
			_1295_v16 = (_1271_v3).F30()
			var _1296_v17 _dafny.Sequence
			_ = _1296_v17
			_1296_v17 = _dafny.SeqOf(_1295_v16)
			(globalState).F5 = (_dafny.IntOfInt64(904)).Minus(_dafny.IntOfUint32((_1296_v17).Cardinality()))
			(globalState).F1 = _this.F24()
			if (_this.F24()) || (_this.F24()) {
				_1281_v12 = _1281_v12
				var _1297_v19 _dafny.Sequence
				_ = _1297_v19
				_1297_v19 = _dafny.SeqOf(_1292_cf42)
				(globalState).F1 = (Companion_D10_.Create_DC29_((_this).F25(), (_this).F25(), (func() _dafny.Map {
					var _coll44 = _dafny.NewMapBuilder()
					_ = _coll44
					for _iter50 := _dafny.Iterate((_1297_v19).Elements()); ; {
						_compr_44, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _1298_v18 _dafny.Int
						_1298_v18 = interface{}(_compr_44).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1297_v19, _1298_v18) {
							_coll44.Add(Companion_Default___.SafeDivisionInt(_1298_v18, _1290_cf44), (_this).F25())
						}
					}
					return _coll44.ToMap()
				}()).Cardinality())).Dtor_cf54()
				var _1299_v20 _dafny.Sequence
				_ = _1299_v20
				_1299_v20 = _dafny.UnicodeSeqOfUtf8Bytes("wrebo")
				var _1300_v21 *C4
				_ = _1300_v21
				var _nw238 *C4 = New_C4_()
				_ = _nw238
				_nw238.Ctor__(!(_this.F24()), !(_this.F24()), (_this).F25())
				_1300_v21 = _nw238
				var _1301_v22 D17
				_ = _1301_v22
				_1301_v22 = Companion_D17_.Create_DC41_(_1300_v21)
				var _1302_v23 _dafny.MultiSet
				_ = _1302_v23
				_1302_v23 = _dafny.MultiSetOf(_this.F24(), _this.F24())
				var _rhs227 _dafny.Set = _1293_cf41
				_ = _rhs227
				var _rhs228 _dafny.Sequence = _1299_v20
				_ = _rhs228
				var _rhs229 *C4 = (func() *C4 {
					if (_this).F25() {
						return _1300_v21
					}
					return (_1301_v22).Dtor_cf73()
				})()
				_ = _rhs229
				var _rhs230 bool = (_1302_v23).IsSubsetOf(_1302_v23)
				_ = _rhs230
				_1293_cf41 = _rhs227
				_1299_v20 = _rhs228
				_1300_v21 = _rhs229
				r0 = _rhs230
				var _1303_v24 _dafny.Map
				_ = _1303_v24
				_1303_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1280_v11, !((_this).F25()))
				var _1304_v25 _dafny.Sequence
				_ = _1304_v25
				_1304_v25 = _dafny.SeqOf(_1297_v19)
				(globalState).F15 = (func() bool {
					if (func() bool {
						if (_1300_v21).F27() {
							return _this.F24()
						}
						return (_this).F25()
					})() {
						return (_1300_v21).F27()
					}
					return (func() bool {
						if (_1303_v24).Contains(((_dafny.MultiSetFromSeq((_1304_v25).Select((Companion_Default___.SafeIndex(_1290_cf44, _dafny.IntOfUint32((_1304_v25).Cardinality()))).Uint32()).(_dafny.Sequence))).Update(_1292_cf42, Companion_Default___.Abs((_this).F23()))).Update(_1294_cf40, Companion_Default___.Abs(_1271_v3.F31))) {
							return (_1303_v24).Get(((_dafny.MultiSetFromSeq((_1304_v25).Select((Companion_Default___.SafeIndex(_1290_cf44, _dafny.IntOfUint32((_1304_v25).Cardinality()))).Uint32()).(_dafny.Sequence))).Update(_1292_cf42, Companion_Default___.Abs((_this).F23()))).Update(_1294_cf40, Companion_Default___.Abs(_1271_v3.F31))).(bool)
						}
						return _this.F24()
					})()
				})()
				var _1305_v26 _dafny.Array
				_ = _1305_v26
				var _nw239 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
				_ = _nw239
				_1305_v26 = _nw239
				var _1306_v27 T1
				_ = _1306_v27
				var _nw240 *C4 = New_C4_()
				_ = _nw240
				_nw240.Ctor__((_1300_v21).F27(), (_this).F25(), (_1300_v21).F27())
				_1306_v27 = _nw240
				var _1307_v28 _dafny.Sequence
				_ = _1307_v28
				_1307_v28 = _dafny.SeqOf(_1306_v27)
				var _nwElement0_53 T1 = _1306_v27
				_ = _nwElement0_53
				var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(21))
				_ = _nw241
				(_nw241).ArraySet1(_nwElement0_53, 0)
				(_nw241).ArraySet1(_1306_v27, 1)
				(_nw241).ArraySet1(_1306_v27, 2)
				(_nw241).ArraySet1(_1306_v27, 3)
				(_nw241).ArraySet1((func() T1 {
					if (_this).Fm3(_1299_v20, (_1306_v27).F25(), _1271_v3.F31, _1297_v19, globalState) {
						return _1306_v27
					}
					return _1306_v27
				})(), 4)
				(_nw241).ArraySet1(_1306_v27, 5)
				(_nw241).ArraySet1(_1306_v27, 6)
				(_nw241).ArraySet1(_1306_v27, 7)
				(_nw241).ArraySet1(_1306_v27, 8)
				(_nw241).ArraySet1(_1306_v27, 9)
				(_nw241).ArraySet1(_1306_v27, 10)
				(_nw241).ArraySet1(_1306_v27, 11)
				(_nw241).ArraySet1(_1306_v27, 12)
				(_nw241).ArraySet1(_1306_v27, 13)
				(_nw241).ArraySet1(_1306_v27, 14)
				(_nw241).ArraySet1(_1306_v27, 15)
				(_nw241).ArraySet1(_1306_v27, 16)
				(_nw241).ArraySet1(_1306_v27, 17)
				(_nw241).ArraySet1(_1306_v27, 18)
				(_nw241).ArraySet1((_1307_v28).Select((Companion_Default___.SafeIndex((_1293_cf41).Cardinality(), _dafny.IntOfUint32((_1307_v28).Cardinality()))).Uint32()).(T1), 19)
				(_nw241).ArraySet1(_1306_v27, 20)
				_1305_v26 = _nw241
			} else {
				var _1308_v29 _dafny.Sequence
				_ = _1308_v29
				_1308_v29 = _dafny.UnicodeSeqOfUtf8Bytes("gigjdhnxs")
				var _1309_v30 _dafny.CodePoint
				_ = _1309_v30
				_1309_v30 = _dafny.CodePoint('p')
				_1308_v29 = _dafny.Companion_Sequence_.Concatenate(_1308_v29, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1308_v29, _1308_v29), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1281_v12).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1308_v29, _1308_v29)).Cardinality()))).Uint32(), _1309_v30))
				_1293_cf41 = (func() _dafny.Set {
					var _coll45 = _dafny.NewBuilder()
					_ = _coll45
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-720), _dafny.IntOfInt64(674))); ; {
						_compr_45, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1310_v31 _dafny.Int
						_1310_v31 = interface{}(_compr_45).(_dafny.Int)
						if ((_dafny.IntOfInt64(-720)).Cmp(_1310_v31) <= 0) && ((_1310_v31).Cmp(_dafny.IntOfInt64(674)) < 0) {
							_coll45.Add((_1310_v31).Plus(_1271_v3.F31))
						}
					}
					return _coll45.ToSet()
				}()).Difference((_1293_cf41).Intersection(_1293_cf41))
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index179
				(_1268_v0).ArraySet1((_this).F25(), (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index180
				(_1268_v0).ArraySet1(!((_this).F25()), (_index180).Int())
				var _1311_v32 _dafny.Map
				_ = _1311_v32
				_1311_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1271_v3.F31, (_this).F25())
				var _1312_v33 _dafny.Sequence
				_ = _1312_v33
				_1312_v33 = _dafny.SeqOf(_1271_v3.F31, _dafny.IntOfInt64(702), _1294_cf40, _1294_cf40)
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index181
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index182
				var _rhs231 bool = _this.F24()
				_ = _rhs231
				var _rhs232 bool = !(true)
				_ = _rhs232
				var _rhs233 bool = (_1293_cf41).IsSubsetOf(_1293_cf41)
				_ = _rhs233
				var _rhs234 bool = (_this).Fm3(_1308_v29, !(_this.F24()), (_1311_v32).Cardinality(), _1312_v33, globalState)
				_ = _rhs234
				var _lhs197 _dafny.Array = _1268_v0
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1268_v0), 0))
				_ = _lhs198
				var _lhs199 _dafny.Array = _1268_v0
				_ = _lhs199
				var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1268_v0), 0))
				_ = _lhs200
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				(_lhs197).ArraySet1(_rhs231, (_lhs198).Int())
				(_lhs199).ArraySet1(_rhs232, (_lhs200).Int())
				_lhs201.F15 = _rhs233
				_lhs202.F15 = _rhs234
				var _1313_v34 _dafny.Array
				_ = _1313_v34
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_39
				var _nw242 _dafny.Array
				_ = _nw242
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw242 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.Sequence = (func(_1314_v29 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1315_i1 _dafny.Int) _dafny.Sequence {
							return _1314_v29
						}
					})(_1308_v29)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw242 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw242).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw242).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1313_v34 = _nw242
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1313_v34), 0))
				_ = _index183
				(_1313_v34).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1308_v29, _1308_v29), (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1313_v34), 0))
				_ = _index184
				(_1313_v34).ArraySet1(_1308_v29, (_index184).Int())
				var _1316_v35 _dafny.Map
				_ = _1316_v35
				_1316_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1283_v14, (_1313_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1313_v34), 0))).Int()).(_dafny.Sequence))
				(globalState).F9 = (_1316_v35).Contains(_1283_v14)
			}
			var _1317_v36 _dafny.Array
			_ = _1317_v36
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_40
			var _nw243 _dafny.Array
			_ = _nw243
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw243 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) _dafny.CodePoint = func(_1318_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw243 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw243).ArraySet1CodePoint(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw243).ArraySet1CodePoint(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1317_v36 = _nw243
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1317_v36), 0))
			_ = _index185
			(_1317_v36).ArraySet1CodePoint(_dafny.CodePoint('l'), (_index185).Int())
			var _1319_v37 _dafny.CodePoint
			_ = _1319_v37
			_1319_v37 = _dafny.CodePoint('m')
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_1317_v36), 0))
			_ = _index186
			(_1317_v36).ArraySet1CodePoint((func() _dafny.CodePoint {
				if false {
					return _1319_v37
				}
				return _1319_v37
			})(), (_index186).Int())
		} else if _source13.Is_DC24() {
			var _1320___mcc_h5 _dafny.Int = _source13.Get_().(D8_DC24).Cf45
			_ = _1320___mcc_h5
			var _1321___mcc_h6 _dafny.CodePoint = _source13.Get_().(D8_DC24).Cf46
			_ = _1321___mcc_h6
			var _1322___mcc_h7 _dafny.MultiSet = _source13.Get_().(D8_DC24).Cf47
			_ = _1322___mcc_h7
			var _1323___mcc_h8 bool = _source13.Get_().(D8_DC24).Cf48
			_ = _1323___mcc_h8
			var _1324___mcc_h9 *C0 = _source13.Get_().(D8_DC24).Cf49
			_ = _1324___mcc_h9
			var _1325_cf49 *C0 = _1324___mcc_h9
			_ = _1325_cf49
			var _1326_cf48 bool = _1323___mcc_h8
			_ = _1326_cf48
			var _1327_cf47 _dafny.MultiSet = _1322___mcc_h7
			_ = _1327_cf47
			var _1328_cf46 _dafny.CodePoint = _1321___mcc_h6
			_ = _1328_cf46
			var _1329_cf45 _dafny.Int = _1320___mcc_h5
			_ = _1329_cf45
			var _1330_v38 _dafny.Sequence
			_ = _1330_v38
			_1330_v38 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			_1330_v38 = _1330_v38
			var _1331_v39 _dafny.Array
			_ = _1331_v39
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_41
			var _nw244 _dafny.Array
			_ = _nw244
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw244 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) D17 = func(_1332_i3 _dafny.Int) D17 {
					return Companion_D17_.Create_DC42_()
				}
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw244 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw244).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw244).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1331_v39 = _nw244
			_1331_v39 = _1331_v39
			var _1333_v40 _dafny.Sequence
			_ = _1333_v40
			_1333_v40 = _dafny.SeqOf((_this).F23())
			var _1334_v41 D11
			_ = _1334_v41
			_1334_v41 = Companion_D11_.Create_DC31_(_1283_v14, _1325_cf49.F31, (_this).F25())
			var _rhs235 _dafny.Sequence = _1333_v40
			_ = _rhs235
			var _rhs236 _dafny.Sequence = _1281_v12
			_ = _rhs236
			var _rhs237 _dafny.Sequence = Companion_Default___.Fm25(!(_1326_cf48), (_1325_cf49.F31).Cmp((_dafny.Zero).Minus(_1329_cf45)) > 0, _1328_cf46, globalState)
			_ = _rhs237
			var _rhs238 bool = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm31((_this).F25(), (_1334_v41).Dtor_cf59(), globalState), (func() _dafny.Sequence {
				if _1326_cf48 {
					return _dafny.Companion_Sequence_.Update(_1333_v40, (Companion_Default___.SafeIndex(_1329_cf45, _dafny.IntOfUint32((_1333_v40).Cardinality()))).Uint32(), _1271_v3.F31)
				}
				return _1333_v40
			})())
			_ = _rhs238
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			_1333_v40 = _rhs235
			_1281_v12 = _rhs236
			_1330_v38 = _rhs237
			_lhs203.F1 = _rhs238
			(globalState).F8 = _1325_cf49.F31
		} else {
			var _1335___mcc_h10 _dafny.Array = _source13.Get_().(D8_DC22).Cf39
			_ = _1335___mcc_h10
			var _1336_cf39 _dafny.Array = _1335___mcc_h10
			_ = _1336_cf39
			if _this.F24() {
				var _1337_v42 _dafny.Map
				_ = _1337_v42
				_1337_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _1280_v11)
				var _1338_v43 _dafny.Set
				_ = _1338_v43
				_1338_v43 = _dafny.SetOf((_this).F25())
				var _1339_v44 _dafny.Sequence
				_ = _1339_v44
				_1339_v44 = _dafny.SeqOf(_1271_v3.F31)
				var _1340_v45 D1
				_ = _1340_v45
				_1340_v45 = Companion_D1_.Create_DC4_(_this.F24(), _this.F24())
				var _1341_v46 _dafny.Set
				_ = _1341_v46
				_1341_v46 = _dafny.SetOf(_1340_v45, _1340_v45, _1340_v45)
				var _1342_v47 _dafny.Map
				_ = _1342_v47
				_1342_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1338_v43).Union(Companion_Default___.Fm32(_this.F24(), _1339_v44, _1281_v12, _1341_v46, globalState)), _1280_v11)
				var _1343_v48 _dafny.Sequence
				_ = _1343_v48
				_1343_v48 = _dafny.SeqOf(_1338_v43)
				var _1344_v49 _dafny.Map
				_ = _1344_v49
				_1344_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _dafny.IntOfInt64(-195))
				var _rhs239 _dafny.MultiSet = (func() _dafny.MultiSet {
					if (_1342_v47).Contains((_1343_v48).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1343_v48).Cardinality()))).Uint32()).(_dafny.Set)) {
						return (_1342_v47).Get((_1343_v48).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1343_v48).Cardinality()))).Uint32()).(_dafny.Set)).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf((_dafny.Zero).Minus((_1344_v49).Cardinality()))
				})()
				_ = _rhs239
				var _rhs240 _dafny.Int = (_1339_v44).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1339_v44).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs240
				var _rhs241 _dafny.Int = (_1271_v3.F31).Minus(_1271_v3.F31)
				_ = _rhs241
				var _rhs242 _dafny.Map = _1337_v42
				_ = _rhs242
				var _lhs204 *GlobalState = globalState
				_ = _lhs204
				_1280_v11 = _rhs239
				_lhs204.F5 = _rhs240
				r3 = _rhs241
				_1337_v42 = _rhs242
				var _1345_v50 _dafny.CodePoint
				_ = _1345_v50
				_1345_v50 = _dafny.CodePoint('v')
				var _1346_v51 D0
				_ = _1346_v51
				_1346_v51 = Companion_D0_.Create_DC1_(_this.F24(), _1271_v3.F31)
				var _1347_v52 _dafny.Map
				_ = _1347_v52
				_1347_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1345_v50, !((_1346_v51).Dtor_cf1()))
				var _1348_v53 D16
				_ = _1348_v53
				_1348_v53 = Companion_D16_.Create_DC40_((_this).F23(), _1336_cf39, _1347_v52)
				var _1349_v54 D0
				_ = _1349_v54
				_1349_v54 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_((_this).F25()))
				var _1350_v55 D0
				_ = _1350_v55
				_1350_v55 = Companion_D0_.Create_DC2_(_1349_v54)
				var _1351_v56 *C6
				_ = _1351_v56
				var _nw245 *C6 = New_C6_()
				_ = _nw245
				_nw245.Ctor__(_1350_v55, true, _this.F24(), _1271_v3.F31)
				_1351_v56 = _nw245
				var _1352_v57 _dafny.MultiSet
				_ = _1352_v57
				_1352_v57 = _dafny.MultiSetOf(_1351_v56, _1351_v56, _1351_v56, _1351_v56)
				var _1353_v58 _dafny.Map
				_ = _1353_v58
				_1353_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1348_v53).Dtor_cf71(), _dafny.IntOfUint32((_dafny.SeqOf(_1352_v57, _1352_v57, _1352_v57, _1352_v57)).Cardinality()))
				var _1354_v59 _dafny.Sequence
				_ = _1354_v59
				_1354_v59 = _dafny.UnicodeSeqOfUtf8Bytes("qqj")
				(globalState).F5 = Companion_Default___.SafeModuloInt(((_1270_v2).Update(_1271_v3.F31, (_this).F23())).Cardinality(), (func() _dafny.Int {
					if (_1353_v58).Contains(_1336_cf39) {
						return (_1353_v58).Get(_1336_cf39).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1354_v59).Cardinality())
				})())
				var _1355_v60 *C1
				_ = _1355_v60
				var _nw246 *C1 = New_C1_()
				_ = _nw246
				_nw246.Ctor__(_this.F24(), _this.F24())
				_1355_v60 = _nw246
				var _1356_v61 _dafny.Sequence
				_ = _1356_v61
				_1356_v61 = _dafny.SeqOf(_1268_v0)
				var _1357_v62 _dafny.Map
				_ = _1357_v62
				_1357_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1356_v61).Cardinality()), (_this).F25())
				var _1358_v63 _dafny.Sequence
				_ = _1358_v63
				_1358_v63 = _dafny.SeqOf(_1357_v62, _1357_v62)
				var _1359_v64 _dafny.Map
				_ = _1359_v64
				_1359_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1355_v60, _dafny.IntOfUint32((_1358_v63).Cardinality()))
				var _1360_v65 _dafny.Set
				_ = _1360_v65
				_1360_v65 = _dafny.SetOf((_1359_v64).Cardinality())
				var _1361_v66 _dafny.Map
				_ = _1361_v66
				_1361_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1360_v65, (_dafny.Zero).Minus((_this).F23()))
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index187
				(_1268_v0).ArraySet1((_1271_v3.F31).Cmp((func() _dafny.Int {
					if (_1361_v66).Contains(_1360_v65) {
						return (_1361_v66).Get(_1360_v65).(_dafny.Int)
					}
					return (_this).F23()
				})()) > 0, (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index188
				(_1268_v0).ArraySet1(_this.F24(), (_index188).Int())
				_1345_v50 = _1345_v50
				(globalState).F15 = !(!((_this).F25()))
			} else {
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1336_cf39), 0))
				_ = _index189
				(_1336_cf39).ArraySet1((_1271_v3.F31).Times(_dafny.IntOfInt64(272)), (_index189).Int())
				var _1362_v67 _dafny.CodePoint
				_ = _1362_v67
				_1362_v67 = _dafny.CodePoint('g')
				var _1363_v68 _dafny.Map
				_ = _1363_v68
				_1363_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1336_cf39), 0))
				_ = _index190
				(_1336_cf39).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), false)).Merge((Companion_Default___.Fm10(_this.F24(), _1271_v3.F31, Companion_Default___.Fm12(false, globalState), _1362_v67, globalState)).Merge(_1363_v68))).Cardinality(), (_index190).Int())
				var _1364_v69 *C5
				_ = _1364_v69
				var _nw247 *C5 = New_C5_()
				_ = _nw247
				_nw247.Ctor__((_dafny.IntOfInt64(6)).Cmp(_1271_v3.F31) == 0, (_this).F25())
				_1364_v69 = _nw247
				var _1365_v70 _dafny.Set
				_ = _1365_v70
				_1365_v70 = _dafny.SetOf((_1281_v12).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_1271_v3.F31, (_this).F25(), globalState), _dafny.IntOfUint32((_1281_v12).Cardinality()))).Uint32()).(bool), !(_this.F24()), _this.F24(), (_this).F25())
				_1365_v70 = (_1365_v70).Intersection(_dafny.SetOf((_this).F25(), _this.F24(), !(_this.F24())))
				var _1366_v71 _dafny.Set
				_ = _1366_v71
				_1366_v71 = _dafny.SetOf(_1271_v3.F31, (_this).F23())
				(globalState).F15 = !(!((_1366_v71).Intersection(_1366_v71)).Equals(_1366_v71))
				var _1367_v72 _dafny.Array
				_ = _1367_v72
				var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
				_ = _nw248
				_1367_v72 = _nw248
				var _nwElement0_54 _dafny.Set = _1366_v71
				_ = _nwElement0_54
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(4))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_54, 0)
				(_nw249).ArraySet1(func() _dafny.Set {
					var _coll46 = _dafny.NewBuilder()
					_ = _coll46
					for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(988), _dafny.IntOfInt64(324))); ; {
						_compr_46, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1368_v73 _dafny.Int
						_1368_v73 = interface{}(_compr_46).(_dafny.Int)
						if ((_dafny.IntOfInt64(988)).Cmp(_1368_v73) <= 0) && ((_1368_v73).Cmp(_dafny.IntOfInt64(324)) < 0) {
							_coll46.Add(Companion_Default___.SafeDivisionInt(_1368_v73, (_1365_v70).Cardinality()))
						}
					}
					return _coll46.ToSet()
				}(), 1)
				(_nw249).ArraySet1((_1366_v71).Difference(_1366_v71), 2)
				(_nw249).ArraySet1((_1366_v71).Union(_1366_v71), 3)
				_1367_v72 = _nw249
			}
			var _1369_v74 T1
			_ = _1369_v74
			var _nw250 *C3 = New_C3_()
			_ = _nw250
			_nw250.Ctor__((_this).F25(), _1271_v3.F31, (_this).F25(), _this.F24(), _1271_v3.F31)
			_1369_v74 = _nw250
			var _1370_v75 _dafny.Sequence
			_ = _1370_v75
			_1370_v75 = _dafny.SeqOf(_1369_v74, _1369_v74)
			var _1371_v76 _dafny.Map
			_ = _1371_v76
			_1371_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_Default___.Fm29(_dafny.IntOfUint32((_1370_v75).Cardinality()), _1271_v3.F31, _1369_v74.F24(), globalState)), _1282_v13)
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))
			_ = _index191
			(_1268_v0).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
				if (_1371_v76).Contains(_1281_v12) {
					return (_1371_v76).Get(_1281_v12).(_dafny.Sequence)
				}
				return _1281_v12
			})(), _1281_v12), (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))
			_ = _index192
			(_1268_v0).ArraySet1((_this).F25(), (_index192).Int())
			(globalState).F6 = _this.F24()
			if (_1369_v74).F25() {
				var _1372_v77 _dafny.Array
				_ = _1372_v77
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_42
				var _nw251 _dafny.Array
				_ = _nw251
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw251 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) D10 = (func(_1373_v74 T1, _1374_v3 *C0) func(_dafny.Int) D10 {
						return func(_1375_i4 _dafny.Int) D10 {
							return (func() D10 {
								if (_1373_v74).F25() {
									return Companion_D10_.Create_DC29_((_this).F25(), !((_1373_v74).F25()), _1374_v3.F31)
								}
								return Companion_D10_.Create_DC29_(_1373_v74.F24(), _1373_v74.F24(), (_this).F23())
							})()
						}
					})(_1369_v74, _1271_v3)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw251 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw251).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw251).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1372_v77 = _nw251
				var _1376_v78 _dafny.Sequence
				_ = _1376_v78
				_1376_v78 = _dafny.SeqOf(_1271_v3.F31)
				var _1377_v79 _dafny.Sequence
				_ = _1377_v79
				_1377_v79 = _dafny.UnicodeSeqOfUtf8Bytes("aawnfu")
				var _1378_v80 _dafny.MultiSet
				_ = _1378_v80
				_1378_v80 = _dafny.MultiSetOf(_1376_v78, _dafny.SeqOf(_dafny.IntOfUint32((_1377_v79).Cardinality())))
				var _1379_v81 _dafny.Set
				_ = _1379_v81
				_1379_v81 = _dafny.SetOf(true, true, !((_this).F25()))
				var _1380_v82 D10
				_ = _1380_v82
				_1380_v82 = Companion_D10_.Create_DC29_((_this).Fm6(_1271_v3.F31, _1378_v80, _this.F24(), (_1379_v81).Cardinality(), globalState), (_1369_v74).F25(), (_this).F23())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1372_v77), 0))
				_ = _index193
				(_1372_v77).ArraySet1(_1380_v82, (_index193).Int())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index194
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1372_v77), 0))
				_ = _index195
				var _rhs243 bool = (_1380_v82).Dtor_cf54()
				_ = _rhs243
				var _rhs244 D10 = _1380_v82
				_ = _rhs244
				var _lhs205 _dafny.Array = _1268_v0
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))
				_ = _lhs206
				var _lhs207 _dafny.Array = _1372_v77
				_ = _lhs207
				var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1372_v77), 0))
				_ = _lhs208
				(_lhs205).ArraySet1(_rhs243, (_lhs206).Int())
				(_lhs207).ArraySet1(_rhs244, (_lhs208).Int())
				var _1381_v83 _dafny.MultiSet
				_ = _1381_v83
				_1381_v83 = _dafny.MultiSetOf((_1369_v74).F25())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1283_v14), 0))
				_ = _index196
				(_1283_v14).ArraySet1((_1381_v83).Cardinality(), (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1283_v14), 0))
				_ = _index197
				var _rhs245 _dafny.Sequence = _1377_v79
				_ = _rhs245
				var _rhs246 bool = ((_this).F23()).Cmp((_this).F23()) == 0
				_ = _rhs246
				var _rhs247 _dafny.Int = (func() _dafny.Int {
					if (_1381_v83).Contains(!((_1369_v74).F25()) || ((_this).F25())) {
						return (_1381_v83).Multiplicity(!((_1369_v74).F25()) || ((_this).F25()))
					}
					return _1271_v3.F31
				})()
				_ = _rhs247
				var _rhs248 _dafny.Int = _1271_v3.F31
				_ = _rhs248
				var _lhs209 *GlobalState = globalState
				_ = _lhs209
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				var _lhs211 _dafny.Array = _1283_v14
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1283_v14), 0))
				_ = _lhs212
				_1377_v79 = _rhs245
				_lhs209.F1 = _rhs246
				_lhs210.F5 = _rhs247
				(_lhs211).ArraySet1(_rhs248, (_lhs212).Int())
				var _1382_v84 *C4
				_ = _1382_v84
				var _nw252 *C4 = New_C4_()
				_ = _nw252
				_nw252.Ctor__((_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), _1369_v74.F24(), _this.F24())
				_1382_v84 = _nw252
				var _rhs249 bool = (func() bool {
					if !((_1381_v83).IsSubsetOf(_1381_v83)) {
						return (_1280_v11).IsSubsetOf(_1280_v11)
					}
					return (_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool)
				})()
				_ = _rhs249
				var _rhs250 bool = (_1382_v84).F27()
				_ = _rhs250
				var _rhs251 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1281_v12, _1281_v12)
				_ = _rhs251
				var _rhs252 _dafny.Int = _1271_v3.F31
				_ = _rhs252
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				var _lhs215 *GlobalState = globalState
				_ = _lhs215
				_lhs213.F15 = _rhs249
				_lhs214.F1 = _rhs250
				_1281_v12 = _rhs251
				_lhs215.F8 = _rhs252
				var _rhs253 _dafny.Array = _1336_cf39
				_ = _rhs253
				var _rhs254 bool = ((_this).F25()) || ((_this).F25())
				_ = _rhs254
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				_1336_cf39 = _rhs253
				_lhs216.F1 = _rhs254
			} else {
				var _1383_v85 _dafny.Array
				_ = _1383_v85
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_43
				var _nw253 _dafny.Array
				_ = _nw253
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw253 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.MultiSet = (func(_1384_v11 _dafny.MultiSet, _1385_v74 T1) func(_dafny.Int) _dafny.MultiSet {
						return func(_1386_i5 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_1384_v11, _1384_v11)).Difference(_dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _1385_v74.F24())).Cardinality(), _dafny.IntOfInt64(981))))
						}
					})(_1280_v11, _1369_v74)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw253 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw253).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw253).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1383_v85 = _nw253
				var _1387_v86 _dafny.MultiSet
				_ = _1387_v86
				_1387_v86 = _dafny.MultiSetOf(_1280_v11, _1280_v11)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1383_v85), 0))
				_ = _index198
				(_1383_v85).ArraySet1((func() _dafny.MultiSet {
					if _1369_v74.F24() {
						return _1387_v86
					}
					return _1387_v86
				})(), (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1383_v85), 0))
				_ = _index199
				(_1383_v85).ArraySet1(_1387_v86, (_index199).Int())
				_1270_v2 = _1270_v2
				var _1388_v87 _dafny.CodePoint
				_ = _1388_v87
				_1388_v87 = _dafny.CodePoint('f')
				var _1389_v88 _dafny.Map
				_ = _1389_v88
				_1389_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v87, (_this).F23())
				var _1390_v89 _dafny.Sequence
				_ = _1390_v89
				_1390_v89 = _dafny.SeqOf(_1389_v88)
				var _1391_v90 D9
				_ = _1391_v90
				_1391_v90 = Companion_D9_.Create_DC27_(Companion_D9_.Create_DC25_(_1390_v89))
				var _1392_v91 D9
				_ = _1392_v91
				_1392_v91 = Companion_D9_.Create_DC25_(_1390_v89)
				_1391_v90 = Companion_D9_.Create_DC27_(_1392_v91)
				var _1393_v92 _dafny.Array
				_ = _1393_v92
				var _nw254 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
				_ = _nw254
				_1393_v92 = _nw254
				_1393_v92 = _1393_v92
				var _1394_v93 _dafny.Array
				_ = _1394_v93
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(2))
				_ = _nw255
				_1394_v93 = _nw255
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_44
				var _nw256 _dafny.Array
				_ = _nw256
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw256 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) D6 = (func(_1395_v74 T1) func(_dafny.Int) D6 {
						return func(_1396_i6 _dafny.Int) D6 {
							return (func() D6 {
								if !(false) {
									return Companion_D6_.Create_DC17_((_1395_v74).F25())
								}
								return Companion_D6_.Create_DC17_(_1395_v74.F24())
							})()
						}
					})(_1369_v74)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw256 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw256).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw256).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1394_v93 = _nw256
			}
		}
		var _1397_v94 _dafny.Sequence
		_ = _1397_v94
		_1397_v94 = _dafny.UnicodeSeqOfUtf8Bytes("erevcsyu")
		_1397_v94 = _dafny.UnicodeSeqOfUtf8Bytes("s")
		if _this.F24() {
			var _1398_v95 _dafny.Map
			_ = _1398_v95
			_1398_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(83), _1268_v0)
			var _1399_v96 *C3
			_ = _1399_v96
			var _nw257 *C3 = New_C3_()
			_ = _nw257
			_nw257.Ctor__(_this.F24(), _dafny.IntOfUint32((_1397_v94).Cardinality()), _this.F24(), (_this).F25(), (_1398_v95).Cardinality())
			_1399_v96 = _nw257
			var _1400_v97 _dafny.Map
			_ = _1400_v97
			_1400_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1399_v96).F28(), _1399_v96)
			var _rhs255 bool = _this.F24()
			_ = _rhs255
			var _rhs256 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24(), _1399_v96)).Merge(_1400_v97)).Cardinality()
			_ = _rhs256
			var _rhs257 bool = !(_this.F24())
			_ = _rhs257
			var _lhs217 *GlobalState = globalState
			_ = _lhs217
			var _lhs218 *GlobalState = globalState
			_ = _lhs218
			var _lhs219 *GlobalState = globalState
			_ = _lhs219
			_lhs217.F9 = _rhs255
			_lhs218.F8 = _rhs256
			_lhs219.F1 = _rhs257
			var _1401_v98 _dafny.CodePoint
			_ = _1401_v98
			_1401_v98 = _dafny.CodePoint('k')
			var _1402_v99 D0
			_ = _1402_v99
			_1402_v99 = Companion_D0_.Create_DC1_((_1399_v96).F28(), _dafny.IntOfUint32((Companion_Default___.Fm20((_1399_v96).F28(), globalState)).Cardinality()))
			var _1403_v100 D0
			_ = _1403_v100
			_1403_v100 = Companion_D0_.Create_DC2_(_1402_v99)
			var _1404_v101 *C6
			_ = _1404_v101
			var _nw258 *C6 = New_C6_()
			_ = _nw258
			_nw258.Ctor__(_1403_v100, (_1399_v96).F28(), true, _dafny.IntOfInt64(219))
			_1404_v101 = _nw258
			var _1405_v102 _dafny.MultiSet
			_ = _1405_v102
			_1405_v102 = _dafny.MultiSetOf(_1404_v101)
			var _pat_let_tv22 = _1402_v99
			_ = _pat_let_tv22
			var _1406_v103 _dafny.Array
			_ = _1406_v103
			var _nwElement0_55 D0 = Companion_D0_.Create_DC2_(_1402_v99)
			_ = _nwElement0_55
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(6))
			_ = _nw259
			(_nw259).ArraySet1(_nwElement0_55, 0)
			(_nw259).ArraySet1(_1403_v100, 1)
			(_nw259).ArraySet1(_1403_v100, 2)
			(_nw259).ArraySet1(_1403_v100, 3)
			(_nw259).ArraySet1(_1403_v100, 4)
			(_nw259).ArraySet1(func(_pat_let11_0 D0) D0 {
				return func(_1407_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let12_0 D0) D0 {
						return func(_1408_dt__update_hcf3_h0 D0) D0 {
							return Companion_D0_.Create_DC2_(_1408_dt__update_hcf3_h0)
						}(_pat_let12_0)
					}(_pat_let_tv22)
				}(_pat_let11_0)
			}(_1404_v101.F26), 5)
			_1406_v103 = _nw259
			var _1409_v104 _dafny.Map
			_ = _1409_v104
			_1409_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1406_v103, _1401_v98)
			var _1410_v105 D3
			_ = _1410_v105
			_1410_v105 = Companion_D3_.Create_DC9_(_dafny.IntOfInt64(101))
			var _1411_v106 _dafny.Map
			_ = _1411_v106
			_1411_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1410_v105, _1283_v14)
			var _1412_v107 D5
			_ = _1412_v107
			_1412_v107 = Companion_D5_.Create_DC15_((func() _dafny.Int {
				if (_1270_v2).Contains((_1405_v102).Cardinality()) {
					return (_1270_v2).Get((_1405_v102).Cardinality()).(_dafny.Int)
				}
				return (_this).F23()
			})(), _1401_v98, _1409_v104, _1411_v106)
			var _1413_v108 _dafny.Array
			_ = _1413_v108
			var _nwElement0_56 _dafny.CodePoint = _1401_v98
			_ = _nwElement0_56
			var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(18))
			_ = _nw260
			(_nw260).ArraySet1CodePoint(_nwElement0_56, 0)
			(_nw260).ArraySet1CodePoint(_1401_v98, 1)
			(_nw260).ArraySet1CodePoint(_1401_v98, 2)
			(_nw260).ArraySet1CodePoint(_1401_v98, 3)
			(_nw260).ArraySet1CodePoint(_1401_v98, 4)
			(_nw260).ArraySet1CodePoint(_1401_v98, 5)
			(_nw260).ArraySet1CodePoint(_1401_v98, 6)
			(_nw260).ArraySet1CodePoint(_1401_v98, 7)
			(_nw260).ArraySet1CodePoint(_dafny.CodePoint('g'), 8)
			(_nw260).ArraySet1CodePoint(_1401_v98, 9)
			(_nw260).ArraySet1CodePoint(_1401_v98, 10)
			(_nw260).ArraySet1CodePoint(_1401_v98, 11)
			(_nw260).ArraySet1CodePoint(_dafny.CodePoint('t'), 12)
			(_nw260).ArraySet1CodePoint(_1401_v98, 13)
			(_nw260).ArraySet1CodePoint((_1412_v107).Dtor_cf26(), 14)
			(_nw260).ArraySet1CodePoint(_1401_v98, 15)
			(_nw260).ArraySet1CodePoint(_dafny.CodePoint('y'), 16)
			(_nw260).ArraySet1CodePoint(_1401_v98, 17)
			_1413_v108 = _nw260
			var _1414_v109 _dafny.Set
			_ = _1414_v109
			_1414_v109 = _dafny.SetOf(_1413_v108, _1413_v108)
			var _1415_v110 D5
			_ = _1415_v110
			_1415_v110 = Companion_D5_.Create_DC14_(_1271_v3.F31, ((_this).F23()).Cmp(_dafny.IntOfUint32((_1397_v94).Cardinality())) <= 0, _1414_v109)
			var _1416_v111 _dafny.Sequence
			_ = _1416_v111
			_1416_v111 = _dafny.SeqOf((_this).F23(), (_this).F23())
			var _pat_let_tv23 = _1416_v111
			_ = _pat_let_tv23
			var _pat_let_tv24 = _1271_v3
			_ = _pat_let_tv24
			var _pat_let_tv25 = _1416_v111
			_ = _pat_let_tv25
			_1415_v110 = func(_pat_let13_0 D5) D5 {
				return func(_1417_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let14_0 bool) D5 {
						return func(_1418_dt__update_hcf23_h0 bool) D5 {
							return func(_pat_let15_0 _dafny.Int) D5 {
								return func(_1419_dt__update_hcf22_h0 _dafny.Int) D5 {
									return Companion_D5_.Create_DC14_(_1419_dt__update_hcf22_h0, _1418_dt__update_hcf23_h0, (_1417_dt__update__tmp_h1).Dtor_cf24())
								}(_pat_let15_0)
							}((_pat_let_tv23).Select((Companion_Default___.SafeIndex(_pat_let_tv24.F31, _dafny.IntOfUint32((_pat_let_tv25).Cardinality()))).Uint32()).(_dafny.Int))
						}(_pat_let14_0)
					}(_this.F24())
				}(_pat_let13_0)
			}(_1415_v110)
			var _1420_v113 _dafny.Map
			_ = _1420_v113
			_1420_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(999), (_1399_v96).F28())
			var _1421_v114 _dafny.Sequence
			_ = _1421_v114
			_1421_v114 = _dafny.SeqOf(_1270_v2, func() _dafny.Map {
				var _coll47 = _dafny.NewMapBuilder()
				_ = _coll47
				for _iter53 := _dafny.Iterate((_1420_v113).Keys().Elements()); ; {
					_compr_47, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _1422_v112 _dafny.Int
					_1422_v112 = interface{}(_compr_47).(_dafny.Int)
					if (_1420_v113).Contains(_1422_v112) {
						_coll47.Add(Companion_Default___.SafeModuloInt(_1422_v112, (_1399_v96).F29()), _1271_v3.F31)
					}
				}
				return _coll47.ToMap()
			}(), _1270_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1271_v3.F31, (_1399_v96).F29()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _1271_v3.F31))
			_1421_v114 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1270_v2), (Companion_Default___.SafeIndex((_1399_v96).F29(), _dafny.IntOfUint32((_dafny.SeqOf(_1270_v2)).Cardinality()))).Uint32(), _1270_v2), _dafny.Companion_Sequence_.Concatenate(_1421_v114, _1421_v114))
			var _1423_v115 _dafny.Array
			_ = _1423_v115
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw261
			_1423_v115 = _nw261
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_1423_v115), 0))
			_ = _index200
			(_1423_v115).ArraySet1(_1281_v12, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_1423_v115), 0))
			_ = _index201
			(_1423_v115).ArraySet1(_dafny.SeqOf((_this).F25(), !((_1281_v12).Select((Companion_Default___.SafeIndex((_1399_v96).F29(), _dafny.IntOfUint32((_1281_v12).Cardinality()))).Uint32()).(bool)), (func() bool {
				if (_1399_v96).F28() {
					return (_1399_v96).F28()
				}
				return (_1399_v96).F28()
			})(), (_1399_v96).F28()), (_index201).Int())
			var _1424_v116 _dafny.Array
			_ = _1424_v116
			var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw262
			_1424_v116 = _nw262
			var _1425_v117 *C1
			_ = _1425_v117
			var _nw263 *C1 = New_C1_()
			_ = _nw263
			_nw263.Ctor__((_1399_v96).F28(), true)
			_1425_v117 = _nw263
			var _1426_v118 _dafny.Map
			_ = _1426_v118
			_1426_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1425_v117, _1401_v98)
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_1283_v14), 0))
			_ = _index202
			(_1283_v14).ArraySet1(((_1399_v96).F29()).Times(_1271_v3.F31), (_index202).Int())
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_1283_v14), 0))
			_ = _index203
			var _rhs258 _dafny.Int = (_1399_v96).Fm5(_1271_v3.F31, (_this).F23(), globalState)
			_ = _rhs258
			var _rhs259 _dafny.Array = _1424_v116
			_ = _rhs259
			var _rhs260 _dafny.Int = _1271_v3.F31
			_ = _rhs260
			var _rhs261 _dafny.Map = _1426_v118
			_ = _rhs261
			var _rhs262 _dafny.Int = _1271_v3.F31
			_ = _rhs262
			var _lhs220 *GlobalState = globalState
			_ = _lhs220
			var _lhs221 *GlobalState = globalState
			_ = _lhs221
			var _lhs222 _dafny.Array = _1283_v14
			_ = _lhs222
			var _lhs223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_1283_v14), 0))
			_ = _lhs223
			_lhs220.F14 = _rhs258
			_1424_v116 = _rhs259
			_lhs221.F8 = _rhs260
			_1426_v118 = _rhs261
			(_lhs222).ArraySet1(_rhs262, (_lhs223).Int())
		} else {
			var _1427_v119 _dafny.CodePoint
			_ = _1427_v119
			_1427_v119 = _dafny.CodePoint('d')
			var _1428_v120 _dafny.Map
			_ = _1428_v120
			_1428_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1271_v3.F31, Companion_Default___.Fm11(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fdnmeq")).Cardinality()), _this.F24(), _this.F24(), _1427_v119, globalState))
			var _1429_v121 D1
			_ = _1429_v121
			_1429_v121 = Companion_D1_.Create_DC4_((_this).F25(), (_this).F25())
			(_this).F24_set_(!((_1428_v120).Update(_1271_v3.F31, _1429_v121)).Equals(_1428_v120))
			var _1430_v122 _dafny.Sequence
			_ = _1430_v122
			_1430_v122 = _dafny.SeqOf(_1397_v94)
			var _1431_v123 _dafny.Sequence
			_ = _1431_v123
			_1431_v123 = _dafny.SeqOf(_1430_v122)
			var _rhs263 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1430_v122, (_1431_v123).Select((Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1431_v123).Cardinality()))).Uint32()).(_dafny.Sequence))
			_ = _rhs263
			var _rhs264 _dafny.CodePoint = _1427_v119
			_ = _rhs264
			_1430_v122 = _rhs263
			_1427_v119 = _rhs264
			var _1432_v124 *C7
			_ = _1432_v124
			var _nw264 *C7 = New_C7_()
			_ = _nw264
			_nw264.Ctor__()
			_1432_v124 = _nw264
			var _1433_v125 _dafny.MultiSet
			_ = _1433_v125
			_1433_v125 = _dafny.MultiSetOf(_1432_v124, _1432_v124)
			var _1434_v126 _dafny.Sequence
			_ = _1434_v126
			_1434_v126 = _dafny.SeqOf((_this).F23(), (_1271_v3.F31).Times((_1433_v125).Cardinality()), (func() _dafny.Int {
				if !(_this.F24()) {
					return _1271_v3.F31
				}
				return _1271_v3.F31
			})(), _1271_v3.F31, (_1271_v3.F31).Plus((_this).F23()))
			_1434_v126 = _dafny.SeqOf((_this).F23())
			if (_1271_v3.F31).Cmp((_1280_v11).Cardinality()) < 0 {
				var _1435_v127 bool
				_ = _1435_v127
				var _1436_v128 bool
				_ = _1436_v128
				var _1437_v129 D0
				_ = _1437_v129
				var _out14 bool
				_ = _out14
				var _out15 bool
				_ = _out15
				var _out16 D0
				_ = _out16
				_out14, _out15, _out16 = (_1432_v124).M2((_this).F23(), ((_1280_v11).Update(_dafny.IntOfInt64(745), Companion_Default___.Abs((_1270_v2).Cardinality()))).IsProperSubsetOf(_1280_v11), globalState)
				_1435_v127 = _out14
				_1436_v128 = _out15
				_1437_v129 = _out16
				var _1438_v130 *C7
				_ = _1438_v130
				var _nw265 *C7 = New_C7_()
				_ = _nw265
				_nw265.Ctor__()
				_1438_v130 = _nw265
				(globalState).F1 = _this.F24()
				_1435_v127 = _this.F24()
				_1436_v128 = _this.F24()
			} else {
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1283_v14), 0))
				_ = _index204
				(_1283_v14).ArraySet1(((func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(421), _dafny.IntOfInt64(182))); ; {
						_compr_48, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1439_v131 _dafny.Int
						_1439_v131 = interface{}(_compr_48).(_dafny.Int)
						if ((_dafny.IntOfInt64(421)).Cmp(_1439_v131) <= 0) && ((_1439_v131).Cmp(_dafny.IntOfInt64(182)) < 0) {
							_coll48.Add(Companion_Default___.SafeDivisionInt(_1439_v131, _dafny.IntOfUint32((_1397_v94).Cardinality())), _this.F24())
						}
					}
					return _coll48.ToMap()
				}()).Update(_1271_v3.F31, (_this).F25())).Cardinality(), (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1283_v14), 0))
				_ = _index205
				(_1283_v14).ArraySet1(Companion_Default___.SafeModuloInt((_this).F23(), (_this).F23()), (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index206
				(_1268_v0).ArraySet1((_this).F25(), (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))
				_ = _index207
				(_1268_v0).ArraySet1(_this.F24(), (_index207).Int())
				var _1440_v132 *C3
				_ = _1440_v132
				var _nw266 *C3 = New_C3_()
				_ = _nw266
				_nw266.Ctor__((_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), _1271_v3.F31, (_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), (_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), _1271_v3.F31)
				_1440_v132 = _nw266
				var _1441_v133 _dafny.Sequence
				_ = _1441_v133
				_1441_v133 = _dafny.SeqOf(_1440_v132)
				var _1442_v134 _dafny.Sequence
				_ = _1442_v134
				_1442_v134 = _dafny.SeqOf(_1441_v133, _1441_v133)
				var _1443_v135 _dafny.Array
				_ = _1443_v135
				var _nwElement0_57 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1441_v133, _1441_v133)
				_ = _nwElement0_57
				var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(12))
				_ = _nw267
				(_nw267).ArraySet1(_nwElement0_57, 0)
				(_nw267).ArraySet1(_1441_v133, 1)
				(_nw267).ArraySet1((_1442_v134).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1442_v134).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
				(_nw267).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1441_v133, _dafny.SeqOf(_1440_v132)), 3)
				(_nw267).ArraySet1(_1441_v133, 4)
				(_nw267).ArraySet1(_1441_v133, 5)
				(_nw267).ArraySet1(_1441_v133, 6)
				(_nw267).ArraySet1((func() _dafny.Sequence {
					if (_1440_v132).F28() {
						return _1441_v133
					}
					return _dafny.SeqOf(_1440_v132)
				})(), 7)
				(_nw267).ArraySet1(_1441_v133, 8)
				(_nw267).ArraySet1(_1441_v133, 9)
				(_nw267).ArraySet1(_1441_v133, 10)
				(_nw267).ArraySet1(_dafny.Companion_Sequence_.Update(_1441_v133, (Companion_Default___.SafeIndex(_1271_v3.F31, _dafny.IntOfUint32((_1441_v133).Cardinality()))).Uint32(), _1440_v132), 11)
				_1443_v135 = _nw267
				var _1444_v136 *C1
				_ = _1444_v136
				var _nw268 *C1 = New_C1_()
				_ = _nw268
				_nw268.Ctor__((_1440_v132).F28(), !((_1440_v132).F28()))
				_1444_v136 = _nw268
				var _rhs265 _dafny.Int = _1271_v3.F31
				_ = _rhs265
				var _rhs266 _dafny.Array = _1443_v135
				_ = _rhs266
				var _rhs267 _dafny.Int = ((_1271_v3.F31).Minus(_1271_v3.F31)).Times(_1271_v3.F31)
				_ = _rhs267
				var _rhs268 *C1 = _1444_v136
				_ = _rhs268
				var _lhs224 *GlobalState = globalState
				_ = _lhs224
				var _lhs225 *GlobalState = globalState
				_ = _lhs225
				_lhs224.F5 = _rhs265
				_1443_v135 = _rhs266
				_lhs225.F5 = _rhs267
				_1444_v136 = _rhs268
				var _1445_v137 T1
				_ = _1445_v137
				var _nw269 *C5 = New_C5_()
				_ = _nw269
				_nw269.Ctor__((_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), (_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool))
				_1445_v137 = _nw269
				var _1446_v138 _dafny.Map
				_ = _1446_v138
				_1446_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_v137, _1445_v137.F24())
				var _1447_v139 _dafny.Map
				_ = _1447_v139
				_1447_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool), _1446_v138)
				var _1448_v140 _dafny.Array
				_ = _1448_v140
				var _nw270 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(3))
				_ = _nw270
				_1448_v140 = _nw270
				var _1449_v141 _dafny.Map
				_ = _1449_v141
				_1449_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1448_v140, _1427_v119)
				var _1450_v142 D3
				_ = _1450_v142
				_1450_v142 = Companion_D3_.Create_DC9_((_1283_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1283_v14), 0))).Int()).(_dafny.Int))
				var _1451_v143 _dafny.Array
				_ = _1451_v143
				var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw271
				_1451_v143 = _nw271
				var _1452_v144 _dafny.Map
				_ = _1452_v144
				_1452_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1450_v142, _1451_v143)
				var _1453_v145 D5
				_ = _1453_v145
				_1453_v145 = Companion_D5_.Create_DC15_(((_1447_v139).Update(false, _1446_v138)).Cardinality(), _1427_v119, _1449_v141, (_1452_v144).Merge(_1452_v144))
				var _1454_v146 _dafny.Set
				_ = _1454_v146
				_1454_v146 = _dafny.SetOf(_1427_v119, _1427_v119)
				_1453_v145 = (func() D5 {
					if (_1281_v12).Select((Companion_Default___.SafeIndex((_1454_v146).Cardinality(), _dafny.IntOfUint32((_1281_v12).Cardinality()))).Uint32()).(bool) {
						return _1453_v145
					}
					return _1453_v145
				})()
				var _1455_v147 D0
				_ = _1455_v147
				_1455_v147 = Companion_D0_.Create_DC0_((_1440_v132).F28())
				var _1456_v148 D0
				_ = _1456_v148
				_1456_v148 = Companion_D0_.Create_DC2_(_1455_v147)
				var _1457_v149 _dafny.Map
				_ = _1457_v149
				_1457_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1440_v132).F29(), _this.F24())
				var _1458_v150 *C6
				_ = _1458_v150
				var _nw272 *C6 = New_C6_()
				_ = _nw272
				_nw272.Ctor__((func() D0 {
					if (_1268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1268_v0), 0))).Int()).(bool) {
						return _1456_v148
					}
					return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_((func() bool {
						if (_1457_v149).Contains((_1440_v132).F29()) {
							return (_1457_v149).Get((_1440_v132).F29()).(bool)
						}
						return _this.F24()
					})()))
				})(), (_1440_v132).F28(), (_1445_v137).F25(), _1271_v3.F31)
				_1458_v150 = _nw272
			}
			(globalState).F15 = (_this).F25()
		}
		var _1459_v151 _dafny.Sequence
		_ = _1459_v151
		_1459_v151 = _dafny.SeqOf(_dafny.IntOfInt64(171), (_dafny.Zero).Minus((_dafny.Zero).Minus(_1271_v3.F31)), _1271_v3.F31, _1271_v3.F31, _dafny.IntOfInt64(443))
		r0 = (_1271_v3).Fm3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-291))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_1460_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})), _this.F24(), _dafny.IntOfUint32((_1459_v151).Cardinality()), (func() _dafny.Sequence {
			if !((_this).F25()) {
				return _1459_v151
			}
			return _dafny.Companion_Sequence_.Update(_1459_v151, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-625)), _dafny.IntOfUint32((_1459_v151).Cardinality()))).Uint32(), (_this).F23())
		})(), globalState)
		var _1461_v152 T0
		_ = _1461_v152
		var _nw273 *C2 = New_C2_()
		_ = _nw273
		_nw273.Ctor__((_dafny.Zero).Minus(_1271_v3.F31))
		_1461_v152 = _nw273
		r1 = _1461_v152
		r2 = (_this).F25()
		r3 = Companion_Default___.Fm1((_this).Fm2(globalState), false, globalState)
		return r0, r1, r2, r3
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
